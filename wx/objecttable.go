package wx

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"unsafe"
)

var globalObjectTable = NewObjectTable()

type objectTableEntry struct {
	// held, 1 bit.
	// ref, reference count, 15 bits.
	// id, 16 bits.
	held_ref_id uint32
	// How to delete this object.
	del func(unsafe.Pointer)
}

func (entry *objectTableEntry) Held() bool {
	return (entry.held_ref_id & 0x80000000) != 0
}

func (entry *objectTableEntry) SetHeld(h bool) {
	if h {
		entry.held_ref_id |= 0x80000000
	} else {
		entry.held_ref_id &= 0x7FFFFFFF
	}
}

func (entry *objectTableEntry) Id() uint16 {
	return uint16(entry.held_ref_id & 0x0000FFFF)
}

func (entry *objectTableEntry) SetId(id uint16) {
	entry.held_ref_id &= 0xFFFF0000
	entry.held_ref_id |= uint32(id)
}

func (entry *objectTableEntry) Ref() uint16 {
	return uint16((entry.held_ref_id >> 16) & 0x7FFF)
}

func (entry *objectTableEntry) SetRef(ref uint16) {
	if ref > 0x7FFF {
		panic("Object reference count overflow")
	}
	entry.held_ref_id &= 0x8000FFFF
	entry.held_ref_id |= (uint32(ref) << 16)
}

func (entry *objectTableEntry) String() string {
	return fmt.Sprintf("Id: %v, Ref: %v Held: %v", entry.Id(), entry.Ref(), entry.Held())
}

// The map key is object pointer.
type objectTable struct {
	nextSeq uint16
	entries map[unsafe.Pointer]*objectTableEntry
	l       sync.RWMutex
}

func NewObjectTable() *objectTable {
	return &objectTable{
		nextSeq: 1,
		entries: make(map[unsafe.Pointer]*objectTableEntry),
	}
}

func (t *objectTable) Remove(ptr unsafe.Pointer) {
	t.l.Lock()
	defer t.l.Unlock()
	if DebugObjectBind {
		if entry, exists := t.entries[ptr]; exists {
			log.Printf("[B] Remove  0x%x \t\t%s\n", ptr, entry)
		}
	}
	delete(t.entries, ptr)
}

func (t *objectTable) lookup(ptr unsafe.Pointer, id uint16) *objectTableEntry {
	if entry, exists := t.entries[ptr]; exists {
		if entry.Id() == id {
			return entry
		}
	}
	return nil
}

func (t *objectTable) IsValid(ptr unsafe.Pointer, id uint16) bool {
	t.l.RLock()
	defer t.l.RUnlock()
	return t.lookup(ptr, id) != nil
}

func (t *objectTable) NextSeq() (id uint16) {
	id = t.nextSeq
	if id == 0xFFFF {
		t.nextSeq = 1
	} else {
		t.nextSeq++
	}
	return
}

func (t *objectTable) Bind(ptr unsafe.Pointer, del func(unsafe.Pointer), hold bool) (id uint16) {
	var entryForDebug *objectTableEntry
	func() {
		t.l.Lock()
		defer t.l.Unlock()
		entry, exists := t.entries[ptr]
		if exists {
			id = entry.Id()
			entry.SetRef(entry.Ref() + 1)
			if DebugObjectBind {
				entryCopy := *entry
				entryForDebug = &entryCopy
			}
			if hold {
				entry.SetHeld(true)
			}
		} else {
			entry = &objectTableEntry{
				del: del,
			}
			entry.SetRef(1)
			entry.SetHeld(hold)
			id = t.NextSeq()
			entry.SetId(id)
			t.entries[ptr] = entry
			if DebugObjectBind {
				entryCopy := *entry
				entryForDebug = &entryCopy
			}
		}
	}()
	if entryForDebug != nil {
		log.Printf("[B] Bind 0x%x(%v)  \t%s Want to hold: %v\n", ptr, wxPtrTypeName(ptr), entryForDebug, hold)
	}
	return
}

func (t *objectTable) Release(ptr unsafe.Pointer, id uint16) {
	var del func(unsafe.Pointer)
	var entryForDebug *objectTableEntry
	func() {
		t.l.Lock()
		defer t.l.Unlock()
		entry := t.lookup(ptr, id)
		if entry == nil {
			return
		}
		entry.SetRef(entry.Ref() - 1)
		if entry.Ref() == 0 {
			delete(t.entries, ptr)
			if entry.Held() {
				del = entry.del
			}
		}
		if DebugObjectBind {
			var entryCopy = *entry
			entryForDebug = &entryCopy
		}
	}()
	if DebugObjectBind {
		result := ""
		if entryForDebug.Ref() == 0 {
			result = " [Removed]"
			if del != nil {
				result = " [Deleted]"
			}
		}
		log.Printf("[B] Release 0x%x  \t%s%v\n", ptr, entryForDebug, result)
	}
	if del != nil {
		del(ptr)
	}
}

func (t *objectTable) Unhold(ptr unsafe.Pointer, id uint16) {
	t.l.Lock()
	defer t.l.Unlock()
	entry := t.lookup(ptr, id)
	if entry == nil {
		return
	}
	if DebugObjectBind {
		log.Printf("[B] Unhold 0x%x  \t\t%s\n", ptr, entry)
	}
	t.entries[ptr].SetHeld(false)
}

func (t *objectTable) DumpToWriter(w io.Writer) {
	t.l.RLock()
	defer t.l.RUnlock()

	fmt.Fprintln(w, "+++++ Begin Object table +++++")
	fmt.Fprintf(w, "  Next id: %v\n", t.nextSeq)
	fmt.Fprintf(w, "  Entry count: %v\n", len(t.entries))
	var i = 0
	for ptr, entry := range t.entries {
		fmt.Fprintf(w, "    %v) 0x%x  \t\t%s\n", i, ptr, entry)
		i++
	}
	fmt.Fprintln(w, "----- End Object table -----")
}

func DumpObjectTable() {
	globalObjectTable.DumpToWriter(os.Stderr)
}
