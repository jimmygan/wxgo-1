package wx

///*  ---------------------------------------------------------------------------- */
///*  Geometric flags */
///*  ---------------------------------------------------------------------------- */
type GeometryCentre int

const (
	CENTRE           GeometryCentre = 0x0001
	CENTER_FRAME                    = 0x0000
	CENTRE_ON_SCREEN                = 0x0002
)

type Orientation int

const (
	/* don't change the values of these elements, they are used elsewhere */
	HORIZONTAL Orientation = 0x0004
	VERTICAL               = 0x0008
	BOTH                   = VERTICAL | HORIZONTAL
	/*  a mask to extract orientation from the combination of flags */
	ORIENTATION_MASK = BOTH
)

type Direction int

const (
	LEFT  Direction = 0x0010
	RIGHT           = 0x0020
	UP              = 0x0040
	DOWN            = 0x0080

	TOP    = UP
	BOTTOM = DOWN

	NORTH = UP
	SOUTH = DOWN
	WEST  = LEFT
	EAST  = RIGHT

	ALL = (UP | DOWN | RIGHT | LEFT)

	/*  a mask to extract direction from the combination of flags */
	DIRECTION_MASK = ALL
)

type Alignment int

const (
	/*
	   0 is a valid wxAlignment value (both wxALIGN_LEFT and wxALIGN_TOP
	   use it) so define a symbolic name for an invalid alignment value
	   which can be assumed to be different from anything else
	*/
	ALIGN_INVALID Alignment = -1

	ALIGN_NOT               = 0x0000
	ALIGN_CENTER_HORIZONTAL = 0x0100
	ALIGN_CENTRE_HORIZONTAL = ALIGN_CENTER_HORIZONTAL
	ALIGN_LEFT              = ALIGN_NOT
	ALIGN_TOP               = ALIGN_NOT
	ALIGN_RIGHT             = 0x0200
	ALIGN_BOTTOM            = 0x0400
	ALIGN_CENTER_VERTICAL   = 0x0800
	ALIGN_CENTRE_VERTICAL   = ALIGN_CENTER_VERTICAL

	ALIGN_CENTER = (ALIGN_CENTER_HORIZONTAL | ALIGN_CENTER_VERTICAL)
	ALIGN_CENTRE = ALIGN_CENTER

	/*  a mask to extract alignment from the combination of flags */
	ALIGN_MASK = 0x0f00
)

type SizerFlagBits int

const (
	/*
	   wxADJUST_MINSIZE doesn't do anything any more but we still define
	   it for compatibility. Notice that it may be also predefined (as 0,
	   hopefully) in the user code in order to use it even in
	   !WXWIN_COMPATIBILITY_2_8 builds so don't redefine it in such case.
	*/
	ADJUST_MINSIZE               SizerFlagBits = 0
	FIXED_MINSIZE                              = 0x8000
	RESERVE_SPACE_EVEN_IF_HIDDEN               = 0x0002

	/*  a mask to extract wxSizerFlagBits from combination of flags */
	SIZER_FLAG_BITS_MASK = 0x8002
)

type Stretch int

const (
	STRETCH_NOT Stretch = 0x0000
	SHRINK              = 0x1000
	GROW                = 0x2000
	EXPAND              = GROW
	SHAPED              = 0x4000
	TILE                = SHAPED | FIXED_MINSIZE

	/*  a mask to extract stretch from the combination of flags */
	STRETCH_MASK = 0x7000 /* sans wxTILE */
)

/*  border flags: the values are chosen for backwards compatibility */
type Border int

const (
	/*  this is different from wxBORDER_NONE as by default the controls do have */
	/*  border */
	BORDER_DEFAULT Border = 0

	BORDER_NONE   = 0x00200000
	BORDER_STATIC = 0x01000000
	BORDER_SIMPLE = 0x02000000
	BORDER_RAISED = 0x04000000
	BORDER_SUNKEN = 0x08000000
	BORDER_DOUBLE = 0x10000000 /* deprecated */
	BORDER_THEME  = BORDER_DOUBLE

	/*  a mask to extract border style from the combination of flags */
	BORDER_MASK = 0x1f200000
)

///* This makes it easier to specify a 'normal' border for a control */
//#if defined(__SMARTPHONE__) || defined(__POCKETPC__)
//#define wxDEFAULT_CONTROL_BORDER    wxBORDER_SIMPLE
//#else
//#define wxDEFAULT_CONTROL_BORDER    wxBORDER_SUNKEN
//#endif

///*  ---------------------------------------------------------------------------- */
///*  Window style flags */
///*  ---------------------------------------------------------------------------- */

///*
// * Values are chosen so they can be |'ed in a bit list.
// * Some styles are used across more than one group,
// * so the values mustn't clash with others in the group.
// * Otherwise, numbers can be reused across groups.
// */

///*
//    Summary of the bits used by various styles.

//    High word, containing styles which can be used with many windows:

//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |31|30|29|28|27|26|25|24|23|22|21|20|19|18|17|16|
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//      |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
//      |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  \_ wxFULL_REPAINT_ON_RESIZE
//      |  |  |  |  |  |  |  |  |  |  |  |  |  |  \____ wxPOPUP_WINDOW
//      |  |  |  |  |  |  |  |  |  |  |  |  |  \_______ wxWANTS_CHARS
//      |  |  |  |  |  |  |  |  |  |  |  |  \__________ wxTAB_TRAVERSAL
//      |  |  |  |  |  |  |  |  |  |  |  \_____________ wxTRANSPARENT_WINDOW
//      |  |  |  |  |  |  |  |  |  |  \________________ wxBORDER_NONE
//      |  |  |  |  |  |  |  |  |  \___________________ wxCLIP_CHILDREN
//      |  |  |  |  |  |  |  |  \______________________ wxALWAYS_SHOW_SB
//      |  |  |  |  |  |  |  \_________________________ wxBORDER_STATIC
//      |  |  |  |  |  |  \____________________________ wxBORDER_SIMPLE
//      |  |  |  |  |  \_______________________________ wxBORDER_RAISED
//      |  |  |  |  \__________________________________ wxBORDER_SUNKEN
//      |  |  |  \_____________________________________ wxBORDER_{DOUBLE,THEME}
//      |  |  \________________________________________ wxCAPTION/wxCLIP_SIBLINGS
//      |  \___________________________________________ wxHSCROLL
//      \______________________________________________ wxVSCROLL

//    Low word style bits is class-specific meaning that the same bit can have
//    different meanings for different controls (e.g. 0x10 is wxCB_READONLY
//    meaning that the control can't be modified for wxComboBox but wxLB_SORT
//    meaning that the control should be kept sorted for wxListBox, while
//    wxLB_SORT has a different value -- and this is just fine).
// */

///*
// * Window (Frame/dialog/subwindow/panel item) style flags
// */
//#define wxVSCROLL               0x80000000
//#define wxHSCROLL               0x40000000
//#define wxCAPTION               0x20000000

///*  New styles (border styles are now in their own enum) */
//#define wxDOUBLE_BORDER         wxBORDER_DOUBLE
//#define wxSUNKEN_BORDER         wxBORDER_SUNKEN
//#define wxRAISED_BORDER         wxBORDER_RAISED
//#define wxBORDER                wxBORDER_SIMPLE
//#define wxSIMPLE_BORDER         wxBORDER_SIMPLE
//#define wxSTATIC_BORDER         wxBORDER_STATIC
//#define wxNO_BORDER             wxBORDER_NONE

///*  wxALWAYS_SHOW_SB: instead of hiding the scrollbar when it is not needed, */
///*  disable it - but still show (see also wxLB_ALWAYS_SB style) */
///*  */
///*  NB: as this style is only supported by wxUniversal and wxMSW so far */
//#define wxALWAYS_SHOW_SB        0x00800000

///*  Clip children when painting, which reduces flicker in e.g. frames and */
///*  splitter windows, but can't be used in a panel where a static box must be */
///*  'transparent' (panel paints the background for it) */
//#define wxCLIP_CHILDREN         0x00400000

///*  Note we're reusing the wxCAPTION style because we won't need captions */
///*  for subwindows/controls */
//#define wxCLIP_SIBLINGS         0x20000000

//#define wxTRANSPARENT_WINDOW    0x00100000

///*  Add this style to a panel to get tab traversal working outside of dialogs */
///*  (on by default for wxPanel, wxDialog, wxScrolledWindow) */
//#define wxTAB_TRAVERSAL         0x00080000

///*  Add this style if the control wants to get all keyboard messages (under */
///*  Windows, it won't normally get the dialog navigation key events) */
//#define wxWANTS_CHARS           0x00040000

///*  Make window retained (Motif only, see src/generic/scrolwing.cpp)
// *  This is non-zero only under wxMotif, to avoid a clash with wxPOPUP_WINDOW
// *  on other platforms
// */

//#ifdef __WXMOTIF__
//#define wxRETAINED              0x00020000
//#else
//#define wxRETAINED              0x00000000
//#endif
//#define wxBACKINGSTORE          wxRETAINED

///*  set this flag to create a special popup window: it will be always shown on */
///*  top of other windows, will capture the mouse and will be dismissed when the */
///*  mouse is clicked outside of it or if it loses focus in any other way */
//#define wxPOPUP_WINDOW          0x00020000

///*  force a full repaint when the window is resized (instead of repainting just */
///*  the invalidated area) */
//#define wxFULL_REPAINT_ON_RESIZE 0x00010000

///*  obsolete: now this is the default behaviour */
///*  */
///*  don't invalidate the whole window (resulting in a PAINT event) when the */
///*  window is resized (currently, makes sense for wxMSW only) */
//#define wxNO_FULL_REPAINT_ON_RESIZE 0

///* A mask which can be used to filter (out) all wxWindow-specific styles.
// */
//#define wxWINDOW_STYLE_MASK     \
//    (wxVSCROLL|wxHSCROLL|wxBORDER_MASK|wxALWAYS_SHOW_SB|wxCLIP_CHILDREN| \
//     wxCLIP_SIBLINGS|wxTRANSPARENT_WINDOW|wxTAB_TRAVERSAL|wxWANTS_CHARS| \
//     wxRETAINED|wxPOPUP_WINDOW|wxFULL_REPAINT_ON_RESIZE)

///*
// * Extra window style flags (use wxWS_EX prefix to make it clear that they
// * should be passed to wxWindow::SetExtraStyle(), not SetWindowStyle())
// */

///*  by default, TransferDataTo/FromWindow() only work on direct children of the */
///*  window (compatible behaviour), set this flag to make them recursively */
///*  descend into all subwindows */
//#define wxWS_EX_VALIDATE_RECURSIVELY    0x00000001

///*  wxCommandEvents and the objects of the derived classes are forwarded to the */
///*  parent window and so on recursively by default. Using this flag for the */
///*  given window allows to block this propagation at this window, i.e. prevent */
///*  the events from being propagated further upwards. The dialogs have this */
///*  flag on by default. */
//#define wxWS_EX_BLOCK_EVENTS            0x00000002

///*  don't use this window as an implicit parent for the other windows: this must */
///*  be used with transient windows as otherwise there is the risk of creating a */
///*  dialog/frame with this window as a parent which would lead to a crash if the */
///*  parent is destroyed before the child */
//#define wxWS_EX_TRANSIENT               0x00000004

///*  don't paint the window background, we'll assume it will */
///*  be done by a theming engine. This is not yet used but could */
///*  possibly be made to work in the future, at least on Windows */
//#define wxWS_EX_THEMED_BACKGROUND       0x00000008

///*  this window should always process idle events */
//#define wxWS_EX_PROCESS_IDLE            0x00000010

///*  this window should always process UI update events */
//#define wxWS_EX_PROCESS_UI_UPDATES      0x00000020

///*  Draw the window in a metal theme on Mac */
//#define wxFRAME_EX_METAL                0x00000040
//#define wxDIALOG_EX_METAL               0x00000040

///*  Use this style to add a context-sensitive help to the window (currently for */
///*  Win32 only and it doesn't work if wxMINIMIZE_BOX or wxMAXIMIZE_BOX are used) */
//#define wxWS_EX_CONTEXTHELP             0x00000080

///* synonyms for wxWS_EX_CONTEXTHELP for compatibility */
//#define wxFRAME_EX_CONTEXTHELP          wxWS_EX_CONTEXTHELP
//#define wxDIALOG_EX_CONTEXTHELP         wxWS_EX_CONTEXTHELP

///*  Create a window which is attachable to another top level window */
//#define wxFRAME_DRAWER          0x0020

///*
// * MDI parent frame style flags
// * Can overlap with some of the above.
// */

//#define wxFRAME_NO_WINDOW_MENU  0x0100

///*
// * wxMenuBar style flags
// */
///*  use native docking */
//#define wxMB_DOCKABLE       0x0001

///*
// * wxMenu style flags
// */
//#define wxMENU_TEAROFF      0x0001

///*
// * Apply to all panel items
// */
//#define wxCOLOURED          0x0800
//#define wxFIXED_LENGTH      0x0400

///*
// * Styles for wxListBox
// */
//#define wxLB_SORT           0x0010
//#define wxLB_SINGLE         0x0020
//#define wxLB_MULTIPLE       0x0040
//#define wxLB_EXTENDED       0x0080
///*  wxLB_OWNERDRAW is Windows-only */
//#define wxLB_NEEDED_SB      0x0000
//#define wxLB_OWNERDRAW      0x0100
//#define wxLB_ALWAYS_SB      0x0200
//#define wxLB_NO_SB          0x0400
//#define wxLB_HSCROLL        wxHSCROLL
///*  always show an entire number of rows */
//#define wxLB_INT_HEIGHT     0x0800

//#if WXWIN_COMPATIBILITY_2_6
//    /*  deprecated synonyms */
//    #define wxPROCESS_ENTER   0x0400  /*  wxTE_PROCESS_ENTER */
//    #define wxPASSWORD        0x0800  /*  wxTE_PASSWORD */
//#endif

///*
// * wxComboBox style flags
// */
//#define wxCB_SIMPLE         0x0004
//#define wxCB_SORT           0x0008
//#define wxCB_READONLY       0x0010
//#define wxCB_DROPDOWN       0x0020

///*
// * wxRadioBox style flags
// */
///*  should we number the items from left to right or from top to bottom in a 2d */
///*  radiobox? */
//#define wxRA_LEFTTORIGHT    0x0001
//#define wxRA_TOPTOBOTTOM    0x0002

///*  New, more intuitive names to specify majorDim argument */
//#define wxRA_SPECIFY_COLS   wxHORIZONTAL
//#define wxRA_SPECIFY_ROWS   wxVERTICAL

///*  Old names for compatibility */
//#define wxRA_HORIZONTAL     wxHORIZONTAL
//#define wxRA_VERTICAL       wxVERTICAL

///*
// * wxRadioButton style flag
// */
//#define wxRB_GROUP          0x0004
//#define wxRB_SINGLE         0x0008

///*
// * wxScrollBar flags
// */
//#define wxSB_HORIZONTAL      wxHORIZONTAL
//#define wxSB_VERTICAL        wxVERTICAL

///*
// * wxSpinButton flags.
// * Note that a wxSpinCtrl is sometimes defined as a wxTextCtrl, and so the
// * flags shouldn't overlap with wxTextCtrl flags that can be used for a single
// * line controls (currently we reuse wxTE_CHARWRAP and wxTE_RICH2 neither of
// * which makes sense for them).
// */
//#define wxSP_HORIZONTAL       wxHORIZONTAL /*  4 */
//#define wxSP_VERTICAL         wxVERTICAL   /*  8 */
//#define wxSP_ARROW_KEYS       0x4000
//#define wxSP_WRAP             0x8000

///*
// * wxTabCtrl flags
// */
//#define wxTC_RIGHTJUSTIFY     0x0010
//#define wxTC_FIXEDWIDTH       0x0020
//#define wxTC_TOP              0x0000    /*  default */
//#define wxTC_LEFT             0x0020
//#define wxTC_RIGHT            0x0040
//#define wxTC_BOTTOM           0x0080
//#define wxTC_MULTILINE        0x0200    /* == wxNB_MULTILINE */
//#define wxTC_OWNERDRAW        0x0400

///*
// * wxStaticBitmap flags
// */
//#define wxBI_EXPAND           wxEXPAND

///*
// * wxStaticLine flags
// */
//#define wxLI_HORIZONTAL         wxHORIZONTAL
//#define wxLI_VERTICAL           wxVERTICAL

///*
// * extended dialog specifiers. these values are stored in a different
// * flag and thus do not overlap with other style flags. note that these
// * values do not correspond to the return values of the dialogs (for
// * those values, look at the wxID_XXX defines).
// */

///*  wxCENTRE already defined as  0x00000001 */
//#define wxYES                   0x00000002
//#define wxOK                    0x00000004
//#define wxNO                    0x00000008
//#define wxYES_NO                (wxYES | wxNO)
//#define wxCANCEL                0x00000010
//#define wxAPPLY                 0x00000020
//#define wxCLOSE                 0x00000040

//#define wxOK_DEFAULT            0x00000000  /* has no effect (default) */
//#define wxYES_DEFAULT           0x00000000  /* has no effect (default) */
//#define wxNO_DEFAULT            0x00000080  /* only valid with wxYES_NO */
//#define wxCANCEL_DEFAULT        0x80000000  /* only valid with wxCANCEL */

//#define wxICON_EXCLAMATION      0x00000100
//#define wxICON_HAND             0x00000200
//#define wxICON_WARNING          wxICON_EXCLAMATION
//#define wxICON_ERROR            wxICON_HAND
//#define wxICON_QUESTION         0x00000400
//#define wxICON_INFORMATION      0x00000800
//#define wxICON_STOP             wxICON_HAND
//#define wxICON_ASTERISK         wxICON_INFORMATION

//#define wxHELP                  0x00001000
//#define wxFORWARD               0x00002000
//#define wxBACKWARD              0x00004000
//#define wxRESET                 0x00008000
//#define wxMORE                  0x00010000
//#define wxSETUP                 0x00020000
//#define wxICON_NONE             0x00040000
//#define wxICON_AUTH_NEEDED      0x00080000

//#define wxICON_MASK \
//    (wxICON_EXCLAMATION|wxICON_HAND|wxICON_QUESTION|wxICON_INFORMATION|wxICON_NONE|wxICON_AUTH_NEEDED)

///*
// * Background styles. See wxWindow::SetBackgroundStyle
// */
//enum wxBackgroundStyle
//{
//    /*
//        background is erased in the EVT_ERASE_BACKGROUND handler or using
//        the system default background if no such handler is defined (this
//        is the default style)
//     */
//    wxBG_STYLE_ERASE,

//    /*
//        background is erased by the system, no EVT_ERASE_BACKGROUND event
//        is generated at all
//     */
//    wxBG_STYLE_SYSTEM,

//    /*
//        background is erased in EVT_PAINT handler and not erased at all
//        before it, this should be used if the paint handler paints over
//        the entire window to avoid flicker
//     */
//    wxBG_STYLE_PAINT,

//    /* this is a Mac-only style, don't use in portable code */
//    wxBG_STYLE_TRANSPARENT,

//    /* this style is deprecated and doesn't do anything, don't use */
//    wxBG_STYLE_COLOUR,

//    /*
//        this style is deprecated and is synonymous with
//        wxBG_STYLE_PAINT, use the new name
//     */
//    wxBG_STYLE_CUSTOM = wxBG_STYLE_PAINT
//};

///*
// * Key types used by (old style) lists and hashes.
// */
//enum wxKeyType
//{
//    wxKEY_NONE,
//    wxKEY_INTEGER,
//    wxKEY_STRING
//};

///*  ---------------------------------------------------------------------------- */
///*  standard IDs */
///*  ---------------------------------------------------------------------------- */

//*  Standard menu IDs */
const (
	//    /*
	//       These ids delimit the range used by automatically-generated ids
	//       (i.e. those used when wxID_ANY is specified during construction).
	//     */
	//#if defined(__WXMSW__) || wxUSE_AUTOID_MANAGEMENT
	//    /*
	//       On MSW the range is always restricted no matter if id management
	//       is used or not because the native window ids are limited to short
	//       range.  On other platforms the range is only restricted if id
	//       management is used so the reference count buffer won't be so big.
	//     */
	//    wxID_AUTO_LOWEST = -32000,
	//    wxID_AUTO_HIGHEST = -2000,
	//#else
	//    wxID_AUTO_LOWEST = -1000000,
	//    wxID_AUTO_HIGHEST = -2000,
	//#endif

	/* no id matches this one when compared to it */
	ID_NONE int = -3

	/*  id for a separator line in the menu (invalid for normal item) */
	ID_SEPARATOR = -2

	/* any id: means that we don't care about the id, whether when installing
	 * an event handler or when creating a new window */
	ID_ANY = -1

	/* all predefined ids are between ID_LOWEST and ID_HIGHEST */
	ID_LOWEST = 4999

	ID_OPEN
	ID_CLOSE
	ID_NEW
	ID_SAVE
	ID_SAVEAS
	ID_REVERT
	ID_EXIT
	ID_UNDO
	ID_REDO
	ID_HELP
	ID_PRINT
	ID_PRINT_SETUP
	ID_PAGE_SETUP
	ID_PREVIEW
	ID_ABOUT
	ID_HELP_CONTENTS
	ID_HELP_INDEX
	ID_HELP_SEARCH
	ID_HELP_COMMANDS
	ID_HELP_PROCEDURES
	ID_HELP_CONTEXT
	ID_CLOSE_ALL
	ID_PREFERENCES

	ID_EDIT = 5030
	ID_CUT
	ID_COPY
	ID_PASTE
	ID_CLEAR
	ID_FIND
	ID_DUPLICATE
	ID_SELECTALL
	ID_DELETE
	ID_REPLACE
	ID_REPLACE_ALL
	ID_PROPERTIES

	ID_VIEW_DETAILS
	ID_VIEW_LARGEICONS
	ID_VIEW_SMALLICONS
	ID_VIEW_LIST
	ID_VIEW_SORTDATE
	ID_VIEW_SORTNAME
	ID_VIEW_SORTSIZE
	ID_VIEW_SORTTYPE

	ID_FILE = 5050
	ID_FILE1
	ID_FILE2
	ID_FILE3
	ID_FILE4
	ID_FILE5
	ID_FILE6
	ID_FILE7
	ID_FILE8
	ID_FILE9

	/*  Standard button and menu IDs */
	ID_OK = 5100
	ID_CANCEL
	ID_APPLY
	ID_YES
	ID_NO
	ID_STATIC
	ID_FORWARD
	ID_BACKWARD
	ID_DEFAULT
	ID_MORE
	ID_SETUP
	ID_RESET
	ID_CONTEXT_HELP
	ID_YESTOALL
	ID_NOTOALL
	ID_ABORT
	ID_RETRY
	ID_IGNORE
	ID_ADD
	ID_REMOVE

	ID_UP
	ID_DOWN
	ID_HOME
	ID_REFRESH
	ID_STOP
	ID_INDEX

	ID_BOLD
	ID_ITALIC
	ID_JUSTIFY_CENTER
	ID_JUSTIFY_FILL
	ID_JUSTIFY_RIGHT
	ID_JUSTIFY_LEFT
	ID_UNDERLINE
	ID_INDENT
	ID_UNINDENT
	ID_ZOOM_100
	ID_ZOOM_FIT
	ID_ZOOM_IN
	ID_ZOOM_OUT
	ID_UNDELETE
	ID_REVERT_TO_SAVED
	ID_CDROM
	ID_CONVERT
	ID_EXECUTE
	ID_FLOPPY
	ID_HARDDISK
	ID_BOTTOM
	ID_FIRST
	ID_LAST
	ID_TOP
	ID_INFO
	ID_JUMP_TO
	ID_NETWORK
	ID_SELECT_COLOR
	ID_SELECT_FONT
	ID_SORT_ASCENDING
	ID_SORT_DESCENDING
	ID_SPELL_CHECK
	ID_STRIKETHROUGH

	/*  System menu IDs (used by wxUniv): */
	ID_SYSTEM_MENU = 5200
	ID_CLOSE_FRAME
	ID_MOVE_FRAME
	ID_RESIZE_FRAME
	ID_MAXIMIZE_FRAME
	ID_ICONIZE_FRAME
	ID_RESTORE_FRAME

	/* MDI window menu ids */
	ID_MDI_WINDOW_FIRST   = 5230
	ID_MDI_WINDOW_CASCADE = ID_MDI_WINDOW_FIRST
	ID_MDI_WINDOW_TILE_HORZ
	ID_MDI_WINDOW_TILE_VERT
	ID_MDI_WINDOW_ARRANGE_ICONS
	ID_MDI_WINDOW_PREV
	ID_MDI_WINDOW_NEXT
	ID_MDI_WINDOW_LAST = ID_MDI_WINDOW_NEXT

	/* OS X system menu ids */
	ID_OSX_MENU_FIRST = 5250
	ID_OSX_HIDE       = ID_OSX_MENU_FIRST
	ID_OSX_HIDEOTHERS
	ID_OSX_SHOWALL
	ID_OSX_MENU_LAST = ID_OSX_SHOWALL

	/*  IDs used by generic file dialog (13 consecutive starting from this value) */
	ID_FILEDLGG = 5900

	/*  IDs used by generic file ctrl (4 consecutive starting from this value) */
	ID_FILECTRL = 5950

	ID_HIGHEST = 5999
)

///*  ---------------------------------------------------------------------------- */
///*  wxWindowID type (after wxID_XYZ enum, platform detection, and dlimpexp.h)    */
///*  ---------------------------------------------------------------------------- */

///*  special care should be taken with this type under Windows where the real */
///*  window id is unsigned, so we must always do the cast before comparing them */
///*  (or else they would be always different!). Using wxGetWindowId() which does */
///*  the cast itself is recommended. Note that this type can't be unsigned */
///*  because wxID_ANY == -1 is a valid (and largely used) value for window id. */
//#if defined(__cplusplus) && wxUSE_GUI
//    #include "wx/windowid.h"
//#endif

///*  ---------------------------------------------------------------------------- */
///*  other constants */
///*  ---------------------------------------------------------------------------- */

//*  menu and toolbar item kinds */
type ItemKind int

const (
	ITEM_SEPARATOR ItemKind = -1
	ITEM_NORMAL    ItemKind = iota
	ITEM_CHECK
	ITEM_RADIO
	ITEM_DROPDOWN
	ITEM_MAX
)

///*
// * The possible states of a 3-state checkbox (Compatible
// * with the 2-state checkbox).
// */
//enum wxCheckBoxState
//{
//    wxCHK_UNCHECKED,
//    wxCHK_CHECKED,
//    wxCHK_UNDETERMINED /* 3-state checkbox only */
//};

///*  hit test results */
//enum wxHitTest
//{
//    wxHT_NOWHERE,

//    /*  scrollbar */
//    wxHT_SCROLLBAR_FIRST = wxHT_NOWHERE,
//    wxHT_SCROLLBAR_ARROW_LINE_1,    /*  left or upper arrow to scroll by line */
//    wxHT_SCROLLBAR_ARROW_LINE_2,    /*  right or down */
//    wxHT_SCROLLBAR_ARROW_PAGE_1,    /*  left or upper arrow to scroll by page */
//    wxHT_SCROLLBAR_ARROW_PAGE_2,    /*  right or down */
//    wxHT_SCROLLBAR_THUMB,           /*  on the thumb */
//    wxHT_SCROLLBAR_BAR_1,           /*  bar to the left/above the thumb */
//    wxHT_SCROLLBAR_BAR_2,           /*  bar to the right/below the thumb */
//    wxHT_SCROLLBAR_LAST,

//    /*  window */
//    wxHT_WINDOW_OUTSIDE,            /*  not in this window at all */
//    wxHT_WINDOW_INSIDE,             /*  in the client area */
//    wxHT_WINDOW_VERT_SCROLLBAR,     /*  on the vertical scrollbar */
//    wxHT_WINDOW_HORZ_SCROLLBAR,     /*  on the horizontal scrollbar */
//    wxHT_WINDOW_CORNER,             /*  on the corner between 2 scrollbars */

//    wxHT_MAX
//};

///*  ---------------------------------------------------------------------------- */
///*  Possible SetSize flags */
///*  ---------------------------------------------------------------------------- */

///*  Use internally-calculated width if -1 */
//#define wxSIZE_AUTO_WIDTH       0x0001
///*  Use internally-calculated height if -1 */
//#define wxSIZE_AUTO_HEIGHT      0x0002
///*  Use internally-calculated width and height if each is -1 */
//#define wxSIZE_AUTO             (wxSIZE_AUTO_WIDTH|wxSIZE_AUTO_HEIGHT)
///*  Ignore missing (-1) dimensions (use existing). */
///*  For readability only: test for wxSIZE_AUTO_WIDTH/HEIGHT in code. */
//#define wxSIZE_USE_EXISTING     0x0000
///*  Allow -1 as a valid position */
//#define wxSIZE_ALLOW_MINUS_ONE  0x0004
///*  Don't do parent client adjustments (for implementation only) */
//#define wxSIZE_NO_ADJUSTMENTS   0x0008
///*  Change the window position even if it seems to be already correct */
//#define wxSIZE_FORCE            0x0010
///*  Emit size event even if size didn't change */
//#define wxSIZE_FORCE_EVENT      0x0020

///*  ---------------------------------------------------------------------------- */
///*  GDI descriptions */
///*  ---------------------------------------------------------------------------- */

//// Hatch styles used by both pen and brush styles.
////
//// NB: Do not use these constants directly, they're for internal use only, use
////     wxBRUSHSTYLE_XXX_HATCH and wxPENSTYLE_XXX_HATCH instead.
//enum wxHatchStyle
//{
//    wxHATCHSTYLE_INVALID = -1,

//    /*
//        The value of the first style is chosen to fit with
//        wxDeprecatedGUIConstants values below, don't change it.
//     */
//    wxHATCHSTYLE_FIRST = 111,
//    wxHATCHSTYLE_BDIAGONAL = wxHATCHSTYLE_FIRST,
//    wxHATCHSTYLE_CROSSDIAG,
//    wxHATCHSTYLE_FDIAGONAL,
//    wxHATCHSTYLE_CROSS,
//    wxHATCHSTYLE_HORIZONTAL,
//    wxHATCHSTYLE_VERTICAL,
//    wxHATCHSTYLE_LAST = wxHATCHSTYLE_VERTICAL
//};

///*
//    WARNING: the following styles are deprecated; use the
//             wxFontFamily, wxFontStyle, wxFontWeight, wxBrushStyle,
//             wxPenStyle, wxPenCap, wxPenJoin enum values instead!
//*/

//#if FUTURE_WXWIN_COMPATIBILITY_3_0

///* don't use any elements of this enum in the new code */
//enum wxDeprecatedGUIConstants
//{
//    /*  Text font families */
//    wxDEFAULT    = 70,
//    wxDECORATIVE,
//    wxROMAN,
//    wxSCRIPT,
//    wxSWISS,
//    wxMODERN,
//    wxTELETYPE,  /* @@@@ */

//    /*  Proportional or Fixed width fonts (not yet used) */
//    wxVARIABLE   = 80,
//    wxFIXED,

//    wxNORMAL     = 90,
//    wxLIGHT,
//    wxBOLD,
//    /*  Also wxNORMAL for normal (non-italic text) */
//    wxITALIC,
//    wxSLANT,

//    /*  Pen styles */
//    wxSOLID      =   100,
//    wxDOT,
//    wxLONG_DASH,
//    wxSHORT_DASH,
//    wxDOT_DASH,
//    wxUSER_DASH,

//    wxTRANSPARENT,

//    /*  Brush & Pen Stippling. Note that a stippled pen cannot be dashed!! */
//    /*  Note also that stippling a Pen IS meaningful, because a Line is */
//    wxSTIPPLE_MASK_OPAQUE, /* mask is used for blitting monochrome using text fore and back ground colors */
//    wxSTIPPLE_MASK,        /* mask is used for masking areas in the stipple bitmap (TO DO) */
//    /*  drawn with a Pen, and without any Brush -- and it can be stippled. */
//    wxSTIPPLE =          110,

//    wxBDIAGONAL_HATCH = wxHATCHSTYLE_BDIAGONAL,
//    wxCROSSDIAG_HATCH = wxHATCHSTYLE_CROSSDIAG,
//    wxFDIAGONAL_HATCH = wxHATCHSTYLE_FDIAGONAL,
//    wxCROSS_HATCH = wxHATCHSTYLE_CROSS,
//    wxHORIZONTAL_HATCH = wxHATCHSTYLE_HORIZONTAL,
//    wxVERTICAL_HATCH = wxHATCHSTYLE_VERTICAL,
//    wxFIRST_HATCH = wxHATCHSTYLE_FIRST,
//    wxLAST_HATCH = wxHATCHSTYLE_LAST
//};
//#endif

///*  ToolPanel in wxFrame (VZ: unused?) */
//enum
//{
//    wxTOOL_TOP = 1,
//    wxTOOL_BOTTOM,
//    wxTOOL_LEFT,
//    wxTOOL_RIGHT
//};

///*  the values of the format constants should be the same as corresponding */
///*  CF_XXX constants in Windows API */
//enum wxDataFormatId
//{
//    wxDF_INVALID =          0,
//    wxDF_TEXT =             1,  /* CF_TEXT */
//    wxDF_BITMAP =           2,  /* CF_BITMAP */
//    wxDF_METAFILE =         3,  /* CF_METAFILEPICT */
//    wxDF_SYLK =             4,
//    wxDF_DIF =              5,
//    wxDF_TIFF =             6,
//    wxDF_OEMTEXT =          7,  /* CF_OEMTEXT */
//    wxDF_DIB =              8,  /* CF_DIB */
//    wxDF_PALETTE =          9,
//    wxDF_PENDATA =          10,
//    wxDF_RIFF =             11,
//    wxDF_WAVE =             12,
//    wxDF_UNICODETEXT =      13,
//    wxDF_ENHMETAFILE =      14,
//    wxDF_FILENAME =         15, /* CF_HDROP */
//    wxDF_LOCALE =           16,
//    wxDF_PRIVATE =          20,
//    wxDF_HTML =             30, /* Note: does not correspond to CF_ constant */
//    wxDF_MAX
//};

///* Key codes */
//enum wxKeyCode
//{
//    WXK_NONE    =    0,

//    WXK_CONTROL_A = 1,
//    WXK_CONTROL_B,
//    WXK_CONTROL_C,
//    WXK_CONTROL_D,
//    WXK_CONTROL_E,
//    WXK_CONTROL_F,
//    WXK_CONTROL_G,
//    WXK_CONTROL_H,
//    WXK_CONTROL_I,
//    WXK_CONTROL_J,
//    WXK_CONTROL_K,
//    WXK_CONTROL_L,
//    WXK_CONTROL_M,
//    WXK_CONTROL_N,
//    WXK_CONTROL_O,
//    WXK_CONTROL_P,
//    WXK_CONTROL_Q,
//    WXK_CONTROL_R,
//    WXK_CONTROL_S,
//    WXK_CONTROL_T,
//    WXK_CONTROL_U,
//    WXK_CONTROL_V,
//    WXK_CONTROL_W,
//    WXK_CONTROL_X,
//    WXK_CONTROL_Y,
//    WXK_CONTROL_Z,

//    WXK_BACK    =    8, /* backspace */
//    WXK_TAB     =    9,
//    WXK_RETURN  =    13,
//    WXK_ESCAPE  =    27,

//    /* values from 33 to 126 are reserved for the standard ASCII characters */

//    WXK_SPACE   =    32,
//    WXK_DELETE  =    127,

//    /* values from 128 to 255 are reserved for ASCII extended characters
//       (note that there isn't a single fixed standard for the meaning
//       of these values; avoid them in portable apps!) */

//    /* These are not compatible with unicode characters.
//       If you want to get a unicode character from a key event, use
//       wxKeyEvent::GetUnicodeKey                                    */
//    WXK_START   = 300,
//    WXK_LBUTTON,
//    WXK_RBUTTON,
//    WXK_CANCEL,
//    WXK_MBUTTON,
//    WXK_CLEAR,
//    WXK_SHIFT,
//    WXK_ALT,
//    WXK_CONTROL,
//    WXK_MENU,
//    WXK_PAUSE,
//    WXK_CAPITAL,
//    WXK_END,
//    WXK_HOME,
//    WXK_LEFT,
//    WXK_UP,
//    WXK_RIGHT,
//    WXK_DOWN,
//    WXK_SELECT,
//    WXK_PRINT,
//    WXK_EXECUTE,
//    WXK_SNAPSHOT,
//    WXK_INSERT,
//    WXK_HELP,
//    WXK_NUMPAD0,
//    WXK_NUMPAD1,
//    WXK_NUMPAD2,
//    WXK_NUMPAD3,
//    WXK_NUMPAD4,
//    WXK_NUMPAD5,
//    WXK_NUMPAD6,
//    WXK_NUMPAD7,
//    WXK_NUMPAD8,
//    WXK_NUMPAD9,
//    WXK_MULTIPLY,
//    WXK_ADD,
//    WXK_SEPARATOR,
//    WXK_SUBTRACT,
//    WXK_DECIMAL,
//    WXK_DIVIDE,
//    WXK_F1,
//    WXK_F2,
//    WXK_F3,
//    WXK_F4,
//    WXK_F5,
//    WXK_F6,
//    WXK_F7,
//    WXK_F8,
//    WXK_F9,
//    WXK_F10,
//    WXK_F11,
//    WXK_F12,
//    WXK_F13,
//    WXK_F14,
//    WXK_F15,
//    WXK_F16,
//    WXK_F17,
//    WXK_F18,
//    WXK_F19,
//    WXK_F20,
//    WXK_F21,
//    WXK_F22,
//    WXK_F23,
//    WXK_F24,
//    WXK_NUMLOCK,
//    WXK_SCROLL,
//    WXK_PAGEUP,
//    WXK_PAGEDOWN,
//#if WXWIN_COMPATIBILITY_2_6
//    WXK_PRIOR = WXK_PAGEUP,
//    WXK_NEXT  = WXK_PAGEDOWN,
//#endif

//    WXK_NUMPAD_SPACE,
//    WXK_NUMPAD_TAB,
//    WXK_NUMPAD_ENTER,
//    WXK_NUMPAD_F1,
//    WXK_NUMPAD_F2,
//    WXK_NUMPAD_F3,
//    WXK_NUMPAD_F4,
//    WXK_NUMPAD_HOME,
//    WXK_NUMPAD_LEFT,
//    WXK_NUMPAD_UP,
//    WXK_NUMPAD_RIGHT,
//    WXK_NUMPAD_DOWN,
//    WXK_NUMPAD_PAGEUP,
//    WXK_NUMPAD_PAGEDOWN,
//#if WXWIN_COMPATIBILITY_2_6
//    WXK_NUMPAD_PRIOR = WXK_NUMPAD_PAGEUP,
//    WXK_NUMPAD_NEXT  = WXK_NUMPAD_PAGEDOWN,
//#endif
//    WXK_NUMPAD_END,
//    WXK_NUMPAD_BEGIN,
//    WXK_NUMPAD_INSERT,
//    WXK_NUMPAD_DELETE,
//    WXK_NUMPAD_EQUAL,
//    WXK_NUMPAD_MULTIPLY,
//    WXK_NUMPAD_ADD,
//    WXK_NUMPAD_SEPARATOR,
//    WXK_NUMPAD_SUBTRACT,
//    WXK_NUMPAD_DECIMAL,
//    WXK_NUMPAD_DIVIDE,

//    WXK_WINDOWS_LEFT,
//    WXK_WINDOWS_RIGHT,
//    WXK_WINDOWS_MENU ,
//#ifdef __WXOSX__
//    WXK_RAW_CONTROL,
//#else
//    WXK_RAW_CONTROL = WXK_CONTROL,
//#endif
//    WXK_COMMAND = WXK_CONTROL,

//    /* Hardware-specific buttons */
//    WXK_SPECIAL1 = 193,
//    WXK_SPECIAL2,
//    WXK_SPECIAL3,
//    WXK_SPECIAL4,
//    WXK_SPECIAL5,
//    WXK_SPECIAL6,
//    WXK_SPECIAL7,
//    WXK_SPECIAL8,
//    WXK_SPECIAL9,
//    WXK_SPECIAL10,
//    WXK_SPECIAL11,
//    WXK_SPECIAL12,
//    WXK_SPECIAL13,
//    WXK_SPECIAL14,
//    WXK_SPECIAL15,
//    WXK_SPECIAL16,
//    WXK_SPECIAL17,
//    WXK_SPECIAL18,
//    WXK_SPECIAL19,
//    WXK_SPECIAL20
//};

///* This enum contains bit mask constants used in wxKeyEvent */
//enum wxKeyModifier
//{
//    wxMOD_NONE      = 0x0000,
//    wxMOD_ALT       = 0x0001,
//    wxMOD_CONTROL   = 0x0002,
//    wxMOD_ALTGR     = wxMOD_ALT | wxMOD_CONTROL,
//    wxMOD_SHIFT     = 0x0004,
//    wxMOD_META      = 0x0008,
//    wxMOD_WIN       = wxMOD_META,
//#if defined(__WXMAC__) || defined(__WXCOCOA__)
//    wxMOD_RAW_CONTROL = 0x0010,
//#else
//    wxMOD_RAW_CONTROL = wxMOD_CONTROL,
//#endif
//    wxMOD_CMD       = wxMOD_CONTROL,
//    wxMOD_ALL       = 0xffff
//};

///* Shortcut for easier dialog-unit-to-pixel conversion */
//#define wxDLG_UNIT(parent, pt) parent->ConvertDialogToPixels(pt)

///* Paper types */
//typedef enum
//{
//    wxPAPER_NONE,               /*  Use specific dimensions */
//    wxPAPER_LETTER,             /*  Letter, 8 1/2 by 11 inches */
//    wxPAPER_LEGAL,              /*  Legal, 8 1/2 by 14 inches */
//    wxPAPER_A4,                 /*  A4 Sheet, 210 by 297 millimeters */
//    wxPAPER_CSHEET,             /*  C Sheet, 17 by 22 inches */
//    wxPAPER_DSHEET,             /*  D Sheet, 22 by 34 inches */
//    wxPAPER_ESHEET,             /*  E Sheet, 34 by 44 inches */
//    wxPAPER_LETTERSMALL,        /*  Letter Small, 8 1/2 by 11 inches */
//    wxPAPER_TABLOID,            /*  Tabloid, 11 by 17 inches */
//    wxPAPER_LEDGER,             /*  Ledger, 17 by 11 inches */
//    wxPAPER_STATEMENT,          /*  Statement, 5 1/2 by 8 1/2 inches */
//    wxPAPER_EXECUTIVE,          /*  Executive, 7 1/4 by 10 1/2 inches */
//    wxPAPER_A3,                 /*  A3 sheet, 297 by 420 millimeters */
//    wxPAPER_A4SMALL,            /*  A4 small sheet, 210 by 297 millimeters */
//    wxPAPER_A5,                 /*  A5 sheet, 148 by 210 millimeters */
//    wxPAPER_B4,                 /*  B4 sheet, 250 by 354 millimeters */
//    wxPAPER_B5,                 /*  B5 sheet, 182-by-257-millimeter paper */
//    wxPAPER_FOLIO,              /*  Folio, 8-1/2-by-13-inch paper */
//    wxPAPER_QUARTO,             /*  Quarto, 215-by-275-millimeter paper */
//    wxPAPER_10X14,              /*  10-by-14-inch sheet */
//    wxPAPER_11X17,              /*  11-by-17-inch sheet */
//    wxPAPER_NOTE,               /*  Note, 8 1/2 by 11 inches */
//    wxPAPER_ENV_9,              /*  #9 Envelope, 3 7/8 by 8 7/8 inches */
//    wxPAPER_ENV_10,             /*  #10 Envelope, 4 1/8 by 9 1/2 inches */
//    wxPAPER_ENV_11,             /*  #11 Envelope, 4 1/2 by 10 3/8 inches */
//    wxPAPER_ENV_12,             /*  #12 Envelope, 4 3/4 by 11 inches */
//    wxPAPER_ENV_14,             /*  #14 Envelope, 5 by 11 1/2 inches */
//    wxPAPER_ENV_DL,             /*  DL Envelope, 110 by 220 millimeters */
//    wxPAPER_ENV_C5,             /*  C5 Envelope, 162 by 229 millimeters */
//    wxPAPER_ENV_C3,             /*  C3 Envelope, 324 by 458 millimeters */
//    wxPAPER_ENV_C4,             /*  C4 Envelope, 229 by 324 millimeters */
//    wxPAPER_ENV_C6,             /*  C6 Envelope, 114 by 162 millimeters */
//    wxPAPER_ENV_C65,            /*  C65 Envelope, 114 by 229 millimeters */
//    wxPAPER_ENV_B4,             /*  B4 Envelope, 250 by 353 millimeters */
//    wxPAPER_ENV_B5,             /*  B5 Envelope, 176 by 250 millimeters */
//    wxPAPER_ENV_B6,             /*  B6 Envelope, 176 by 125 millimeters */
//    wxPAPER_ENV_ITALY,          /*  Italy Envelope, 110 by 230 millimeters */
//    wxPAPER_ENV_MONARCH,        /*  Monarch Envelope, 3 7/8 by 7 1/2 inches */
//    wxPAPER_ENV_PERSONAL,       /*  6 3/4 Envelope, 3 5/8 by 6 1/2 inches */
//    wxPAPER_FANFOLD_US,         /*  US Std Fanfold, 14 7/8 by 11 inches */
//    wxPAPER_FANFOLD_STD_GERMAN, /*  German Std Fanfold, 8 1/2 by 12 inches */
//    wxPAPER_FANFOLD_LGL_GERMAN, /*  German Legal Fanfold, 8 1/2 by 13 inches */

//    wxPAPER_ISO_B4,             /*  B4 (ISO) 250 x 353 mm */
//    wxPAPER_JAPANESE_POSTCARD,  /*  Japanese Postcard 100 x 148 mm */
//    wxPAPER_9X11,               /*  9 x 11 in */
//    wxPAPER_10X11,              /*  10 x 11 in */
//    wxPAPER_15X11,              /*  15 x 11 in */
//    wxPAPER_ENV_INVITE,         /*  Envelope Invite 220 x 220 mm */
//    wxPAPER_LETTER_EXTRA,       /*  Letter Extra 9 \275 x 12 in */
//    wxPAPER_LEGAL_EXTRA,        /*  Legal Extra 9 \275 x 15 in */
//    wxPAPER_TABLOID_EXTRA,      /*  Tabloid Extra 11.69 x 18 in */
//    wxPAPER_A4_EXTRA,           /*  A4 Extra 9.27 x 12.69 in */
//    wxPAPER_LETTER_TRANSVERSE,  /*  Letter Transverse 8 \275 x 11 in */
//    wxPAPER_A4_TRANSVERSE,      /*  A4 Transverse 210 x 297 mm */
//    wxPAPER_LETTER_EXTRA_TRANSVERSE, /*  Letter Extra Transverse 9\275 x 12 in */
//    wxPAPER_A_PLUS,             /*  SuperA/SuperA/A4 227 x 356 mm */
//    wxPAPER_B_PLUS,             /*  SuperB/SuperB/A3 305 x 487 mm */
//    wxPAPER_LETTER_PLUS,        /*  Letter Plus 8.5 x 12.69 in */
//    wxPAPER_A4_PLUS,            /*  A4 Plus 210 x 330 mm */
//    wxPAPER_A5_TRANSVERSE,      /*  A5 Transverse 148 x 210 mm */
//    wxPAPER_B5_TRANSVERSE,      /*  B5 (JIS) Transverse 182 x 257 mm */
//    wxPAPER_A3_EXTRA,           /*  A3 Extra 322 x 445 mm */
//    wxPAPER_A5_EXTRA,           /*  A5 Extra 174 x 235 mm */
//    wxPAPER_B5_EXTRA,           /*  B5 (ISO) Extra 201 x 276 mm */
//    wxPAPER_A2,                 /*  A2 420 x 594 mm */
//    wxPAPER_A3_TRANSVERSE,      /*  A3 Transverse 297 x 420 mm */
//    wxPAPER_A3_EXTRA_TRANSVERSE, /*  A3 Extra Transverse 322 x 445 mm */

//    wxPAPER_DBL_JAPANESE_POSTCARD,/* Japanese Double Postcard 200 x 148 mm */
//    wxPAPER_A6,                 /* A6 105 x 148 mm */
//    wxPAPER_JENV_KAKU2,         /* Japanese Envelope Kaku #2 */
//    wxPAPER_JENV_KAKU3,         /* Japanese Envelope Kaku #3 */
//    wxPAPER_JENV_CHOU3,         /* Japanese Envelope Chou #3 */
//    wxPAPER_JENV_CHOU4,         /* Japanese Envelope Chou #4 */
//    wxPAPER_LETTER_ROTATED,     /* Letter Rotated 11 x 8 1/2 in */
//    wxPAPER_A3_ROTATED,         /* A3 Rotated 420 x 297 mm */
//    wxPAPER_A4_ROTATED,         /* A4 Rotated 297 x 210 mm */
//    wxPAPER_A5_ROTATED,         /* A5 Rotated 210 x 148 mm */
//    wxPAPER_B4_JIS_ROTATED,     /* B4 (JIS) Rotated 364 x 257 mm */
//    wxPAPER_B5_JIS_ROTATED,     /* B5 (JIS) Rotated 257 x 182 mm */
//    wxPAPER_JAPANESE_POSTCARD_ROTATED,/* Japanese Postcard Rotated 148 x 100 mm */
//    wxPAPER_DBL_JAPANESE_POSTCARD_ROTATED,/* Double Japanese Postcard Rotated 148 x 200 mm */
//    wxPAPER_A6_ROTATED,         /* A6 Rotated 148 x 105 mm */
//    wxPAPER_JENV_KAKU2_ROTATED, /* Japanese Envelope Kaku #2 Rotated */
//    wxPAPER_JENV_KAKU3_ROTATED, /* Japanese Envelope Kaku #3 Rotated */
//    wxPAPER_JENV_CHOU3_ROTATED, /* Japanese Envelope Chou #3 Rotated */
//    wxPAPER_JENV_CHOU4_ROTATED, /* Japanese Envelope Chou #4 Rotated */
//    wxPAPER_B6_JIS,             /* B6 (JIS) 128 x 182 mm */
//    wxPAPER_B6_JIS_ROTATED,     /* B6 (JIS) Rotated 182 x 128 mm */
//    wxPAPER_12X11,              /* 12 x 11 in */
//    wxPAPER_JENV_YOU4,          /* Japanese Envelope You #4 */
//    wxPAPER_JENV_YOU4_ROTATED,  /* Japanese Envelope You #4 Rotated */
//    wxPAPER_P16K,               /* PRC 16K 146 x 215 mm */
//    wxPAPER_P32K,               /* PRC 32K 97 x 151 mm */
//    wxPAPER_P32KBIG,            /* PRC 32K(Big) 97 x 151 mm */
//    wxPAPER_PENV_1,             /* PRC Envelope #1 102 x 165 mm */
//    wxPAPER_PENV_2,             /* PRC Envelope #2 102 x 176 mm */
//    wxPAPER_PENV_3,             /* PRC Envelope #3 125 x 176 mm */
//    wxPAPER_PENV_4,             /* PRC Envelope #4 110 x 208 mm */
//    wxPAPER_PENV_5,             /* PRC Envelope #5 110 x 220 mm */
//    wxPAPER_PENV_6,             /* PRC Envelope #6 120 x 230 mm */
//    wxPAPER_PENV_7,             /* PRC Envelope #7 160 x 230 mm */
//    wxPAPER_PENV_8,             /* PRC Envelope #8 120 x 309 mm */
//    wxPAPER_PENV_9,             /* PRC Envelope #9 229 x 324 mm */
//    wxPAPER_PENV_10,            /* PRC Envelope #10 324 x 458 mm */
//    wxPAPER_P16K_ROTATED,       /* PRC 16K Rotated */
//    wxPAPER_P32K_ROTATED,       /* PRC 32K Rotated */
//    wxPAPER_P32KBIG_ROTATED,    /* PRC 32K(Big) Rotated */
//    wxPAPER_PENV_1_ROTATED,     /* PRC Envelope #1 Rotated 165 x 102 mm */
//    wxPAPER_PENV_2_ROTATED,     /* PRC Envelope #2 Rotated 176 x 102 mm */
//    wxPAPER_PENV_3_ROTATED,     /* PRC Envelope #3 Rotated 176 x 125 mm */
//    wxPAPER_PENV_4_ROTATED,     /* PRC Envelope #4 Rotated 208 x 110 mm */
//    wxPAPER_PENV_5_ROTATED,     /* PRC Envelope #5 Rotated 220 x 110 mm */
//    wxPAPER_PENV_6_ROTATED,     /* PRC Envelope #6 Rotated 230 x 120 mm */
//    wxPAPER_PENV_7_ROTATED,     /* PRC Envelope #7 Rotated 230 x 160 mm */
//    wxPAPER_PENV_8_ROTATED,     /* PRC Envelope #8 Rotated 309 x 120 mm */
//    wxPAPER_PENV_9_ROTATED,     /* PRC Envelope #9 Rotated 324 x 229 mm */
//    wxPAPER_PENV_10_ROTATED,    /* PRC Envelope #10 Rotated 458 x 324 m */
//    wxPAPER_A0,                 /* A0 Sheet 841 x 1189 mm */
//    wxPAPER_A1                  /* A1 Sheet 594 x 841 mm */
//} wxPaperSize;

///* Printing orientation */
//enum wxPrintOrientation
//{
//   wxPORTRAIT = 1,
//   wxLANDSCAPE
//};

///* Duplex printing modes
// */

//enum wxDuplexMode
//{
//    wxDUPLEX_SIMPLEX, /*  Non-duplex */
//    wxDUPLEX_HORIZONTAL,
//    wxDUPLEX_VERTICAL
//};

///* Print quality.
// */

//#define wxPRINT_QUALITY_HIGH    -1
//#define wxPRINT_QUALITY_MEDIUM  -2
//#define wxPRINT_QUALITY_LOW     -3
//#define wxPRINT_QUALITY_DRAFT   -4

//typedef int wxPrintQuality;

///* Print mode (currently PostScript only)
// */

//enum wxPrintMode
//{
//    wxPRINT_MODE_NONE =    0,
//    wxPRINT_MODE_PREVIEW = 1,   /*  Preview in external application */
//    wxPRINT_MODE_FILE =    2,   /*  Print to file */
//    wxPRINT_MODE_PRINTER = 3,   /*  Send to printer */
//    wxPRINT_MODE_STREAM =  4    /*  Send postscript data into a stream */
//};

///*  ---------------------------------------------------------------------------- */
///*  UpdateWindowUI flags */
///*  ---------------------------------------------------------------------------- */

//enum wxUpdateUI
//{
//    wxUPDATE_UI_NONE          = 0x0000,
//    wxUPDATE_UI_RECURSE       = 0x0001,
//    wxUPDATE_UI_FROMIDLE      = 0x0002 /*  Invoked from On(Internal)Idle */
//};

///* ---------------------------------------------------------------------------- */
///* wxList types */
///* ---------------------------------------------------------------------------- */

///* type of compare function for list sort operation (as in 'qsort'): it should
//   return a negative value, 0 or positive value if the first element is less
//   than, equal or greater than the second */

//typedef int (* LINKAGEMODE wxSortCompareFunction)(const void *elem1, const void *elem2);

///* wxList iterator function */
//typedef int (* LINKAGEMODE wxListIterateFunction)(void *current);

///*  ---------------------------------------------------------------------------- */
///*  miscellaneous */
///*  ---------------------------------------------------------------------------- */

///*  define this macro if font handling is done using the X font names */
//#if (defined(__WXGTK__) && !defined(__WXGTK20__)) || defined(__X__)
//    #define _WX_X_FONTLIKE
//#endif

///*  macro to specify "All Files" on different platforms */
//#if defined(__WXMSW__) || defined(__WXPM__)
//#   define wxALL_FILES_PATTERN   wxT("*.*")
//#   define wxALL_FILES           gettext_noop("All files (*.*)|*.*")
//#else
//#   define wxALL_FILES_PATTERN   wxT("*")
//#   define wxALL_FILES           gettext_noop("All files (*)|*")
//#endif

//#if defined(__CYGWIN__) && defined(__WXMSW__)
//#   if wxUSE_STD_CONTAINERS || defined(wxUSE_STD_STRING)
//         /*
//            NASTY HACK because the gethostname in sys/unistd.h which the gnu
//            stl includes and wx builds with by default clash with each other
//            (windows version 2nd param is int, sys/unistd.h version is unsigned
//            int).
//          */
//#        define gethostname gethostnameHACK
//#        include <unistd.h>
//#        undef gethostname
//#   endif
//#endif

///*  --------------------------------------------------------------------------- */
///*  macros that enable wxWidgets apps to be compiled in absence of the */
///*  system headers, although some platform specific types are used in the */
///*  platform specific (implementation) parts of the headers */
///*  --------------------------------------------------------------------------- */

//#ifdef __DARWIN__
//#define DECLARE_WXOSX_OPAQUE_CFREF( name ) typedef struct __##name* name##Ref;
//#define DECLARE_WXOSX_OPAQUE_CONST_CFREF( name ) typedef const struct __##name* name##Ref;
//#endif

//#ifdef __WXMAC__

//#define WX_OPAQUE_TYPE( name ) struct wxOpaque##name

//typedef void*       WXHBITMAP;
//typedef void*       WXHCURSOR;
//typedef void*       WXRECTPTR;
//typedef void*       WXPOINTPTR;
//typedef void*       WXHWND;
//typedef void*       WXEVENTREF;
//typedef void*       WXEVENTHANDLERREF;
//typedef void*       WXEVENTHANDLERCALLREF;
//typedef void*       WXAPPLEEVENTREF;

//typedef unsigned int    WXUINT;
//typedef unsigned long   WXDWORD;
//typedef unsigned short  WXWORD;

//typedef WX_OPAQUE_TYPE(PicHandle ) * WXHMETAFILE ;
//#if wxOSX_USE_CARBON
//typedef struct OpaqueControlRef* WXWidget ;
//typedef struct OpaqueWindowPtr* WXWindow ;
//typedef struct __AGLPixelFormatRec   *WXGLPixelFormat;
//typedef struct __AGLContextRec       *WXGLContext;
//#endif

//typedef void*       WXDisplay;

///*
// * core frameworks
// */

//typedef const void * CFTypeRef;

///* typedef const struct __CFString * CFStringRef; */

//DECLARE_WXOSX_OPAQUE_CONST_CFREF( CFString )
//typedef struct __CFString * CFMutableStringRef;

//DECLARE_WXOSX_OPAQUE_CFREF( CFRunLoopSource )
//DECLARE_WXOSX_OPAQUE_CONST_CFREF( CTFont )
//DECLARE_WXOSX_OPAQUE_CONST_CFREF( CTFontDescriptor )

//#define DECLARE_WXOSX_OPAQUE_CGREF( name ) typedef struct name* name##Ref;

//DECLARE_WXOSX_OPAQUE_CGREF( CGColor )
//DECLARE_WXOSX_OPAQUE_CGREF( CGImage )
//DECLARE_WXOSX_OPAQUE_CGREF( CGContext )
//DECLARE_WXOSX_OPAQUE_CGREF( CGFont )

//typedef CGColorRef    WXCOLORREF;
//typedef CGImageRef    WXCGIMAGEREF;
//typedef CGContextRef  WXHDC;

///*
// * carbon
// */

//typedef const struct __HIShape * HIShapeRef;
//typedef struct __HIShape * HIMutableShapeRef;

//#define DECLARE_WXMAC_OPAQUE_REF( name ) typedef struct Opaque##name* name;

//DECLARE_WXMAC_OPAQUE_REF( PasteboardRef )
//DECLARE_WXMAC_OPAQUE_REF( IconRef )
//DECLARE_WXMAC_OPAQUE_REF( MenuRef )

//typedef IconRef WXHICON ;
//typedef HIShapeRef WXHRGN;
//#if wxOSX_USE_CARBON
//typedef MenuRef WXHMENU;
//#endif

//#endif

//#if defined( __WXCOCOA__ ) || defined(__WXMAC__)

///* Definitions of 32-bit/64-bit types
// * These are typedef'd exactly the same way in newer OS X headers so
// * redefinition when real headers are included should not be a problem.  If
// * it is, the types are being defined wrongly here.
// * The purpose of these types is so they can be used from public wx headers.
// * and also because the older (pre-Leopard) headers don't define them.
// */

///* NOTE: We don't pollute namespace with CGFLOAT_MIN/MAX/IS_DOUBLE macros
// * since they are unlikely to be needed in a public header.
// */
//#if defined(__LP64__) && __LP64__
//    typedef double CGFloat;
//#else
//    typedef float CGFloat;
//#endif

//#if (defined(__LP64__) && __LP64__) || (defined(NS_BUILD_32_LIKE_64) && NS_BUILD_32_LIKE_64)
//typedef long NSInteger;
//typedef unsigned long NSUInteger;
//#else
//typedef int NSInteger;
//typedef unsigned int NSUInteger;
//#endif

///* Objective-C type declarations.
// * These are to be used in public headers in lieu of NSSomething* because
// * Objective-C class names are not available in C/C++ code.
// */

///*  NOTE: This ought to work with other compilers too, but I'm being cautious */
//#if (defined(__GNUC__) && defined(__APPLE__))
///* It's desirable to have type safety for Objective-C(++) code as it does
//at least catch typos of method names among other things.  However, it
//is not possible to declare an Objective-C class from plain old C or C++
//code.  Furthermore, because of C++ name mangling, the type name must
//be the same for both C++ and Objective-C++ code.  Therefore, we define
//what should be a pointer to an Objective-C class as a pointer to a plain
//old C struct with the same name.  Unfortunately, because the compiler
//does not see a struct as an Objective-C class we cannot declare it
//as a struct in Objective-C(++) mode.
//*/
//#if defined(__OBJC__)
//#define DECLARE_WXCOCOA_OBJC_CLASS(klass) \
//@class klass; \
//typedef klass *WX_##klass
//#else /*  not defined(__OBJC__) */
//#define DECLARE_WXCOCOA_OBJC_CLASS(klass) \
//typedef struct klass *WX_##klass
//#endif /*  defined(__OBJC__) */

//#else /*  not Apple's gcc */
//#warning "Objective-C types will not be checked by the compiler."
///*  NOTE: typedef struct objc_object *id; */
///*  IOW, we're declaring these using the id type without using that name, */
///*  since "id" is used extensively not only within wxWidgets itself, but */
///*  also in wxWidgets application code.  The following works fine when */
///*  compiling C(++) code, and works without typesafety for Obj-C(++) code */
//#define DECLARE_WXCOCOA_OBJC_CLASS(klass) \
//typedef struct objc_object *WX_##klass

//#endif /*  (defined(__GNUC__) && defined(__APPLE__)) */

//DECLARE_WXCOCOA_OBJC_CLASS(NSApplication);
//DECLARE_WXCOCOA_OBJC_CLASS(NSBitmapImageRep);
//DECLARE_WXCOCOA_OBJC_CLASS(NSBox);
//DECLARE_WXCOCOA_OBJC_CLASS(NSButton);
//DECLARE_WXCOCOA_OBJC_CLASS(NSColor);
//DECLARE_WXCOCOA_OBJC_CLASS(NSColorPanel);
//DECLARE_WXCOCOA_OBJC_CLASS(NSControl);
//DECLARE_WXCOCOA_OBJC_CLASS(NSCursor);
//DECLARE_WXCOCOA_OBJC_CLASS(NSEvent);
//DECLARE_WXCOCOA_OBJC_CLASS(NSFont);
//DECLARE_WXCOCOA_OBJC_CLASS(NSFontDescriptor);
//DECLARE_WXCOCOA_OBJC_CLASS(NSFontPanel);
//DECLARE_WXCOCOA_OBJC_CLASS(NSImage);
//DECLARE_WXCOCOA_OBJC_CLASS(NSLayoutManager);
//DECLARE_WXCOCOA_OBJC_CLASS(NSMenu);
//DECLARE_WXCOCOA_OBJC_CLASS(NSMenuExtra);
//DECLARE_WXCOCOA_OBJC_CLASS(NSMenuItem);
//DECLARE_WXCOCOA_OBJC_CLASS(NSMutableArray);
//DECLARE_WXCOCOA_OBJC_CLASS(NSNotification);
//DECLARE_WXCOCOA_OBJC_CLASS(NSObject);
//DECLARE_WXCOCOA_OBJC_CLASS(NSPanel);
//DECLARE_WXCOCOA_OBJC_CLASS(NSResponder);
//DECLARE_WXCOCOA_OBJC_CLASS(NSScrollView);
//DECLARE_WXCOCOA_OBJC_CLASS(NSSound);
//DECLARE_WXCOCOA_OBJC_CLASS(NSStatusItem);
//DECLARE_WXCOCOA_OBJC_CLASS(NSTableColumn);
//DECLARE_WXCOCOA_OBJC_CLASS(NSTableView);
//DECLARE_WXCOCOA_OBJC_CLASS(NSTextContainer);
//DECLARE_WXCOCOA_OBJC_CLASS(NSTextField);
//DECLARE_WXCOCOA_OBJC_CLASS(NSTextStorage);
//DECLARE_WXCOCOA_OBJC_CLASS(NSThread);
//DECLARE_WXCOCOA_OBJC_CLASS(NSWindow);
//DECLARE_WXCOCOA_OBJC_CLASS(NSView);
//DECLARE_WXCOCOA_OBJC_CLASS(NSOpenGLContext);
//DECLARE_WXCOCOA_OBJC_CLASS(NSOpenGLPixelFormat);
//DECLARE_WXCOCOA_OBJC_CLASS( NSPrintInfo );
//#ifndef __WXMAC__
//typedef WX_NSView WXWidget; /*  wxWidgets BASE definition */
//#endif
//#endif /*  __WXCOCOA__  || ( __WXMAC__ &__DARWIN__)*/

//#ifdef __WXMAC__

//DECLARE_WXCOCOA_OBJC_CLASS(NSString);

//#if wxOSX_USE_COCOA

//typedef WX_NSWindow WXWindow;
//typedef WX_NSView WXWidget;
//typedef WX_NSMenu WXHMENU;
//typedef WX_NSOpenGLPixelFormat WXGLPixelFormat;
//typedef WX_NSOpenGLContext WXGLContext;

//#elif wxOSX_USE_IPHONE

//DECLARE_WXCOCOA_OBJC_CLASS(UIWindow);
//DECLARE_WXCOCOA_OBJC_CLASS(UIView);
//DECLARE_WXCOCOA_OBJC_CLASS(UIFont);
//DECLARE_WXCOCOA_OBJC_CLASS(UIImage);
//DECLARE_WXCOCOA_OBJC_CLASS(UIEvent);
//DECLARE_WXCOCOA_OBJC_CLASS(NSSet);
//DECLARE_WXCOCOA_OBJC_CLASS(EAGLContext);

//typedef WX_UIWindow WXWindow;
//typedef WX_UIView WXWidget;
//typedef WX_EAGLContext WXGLContext;
//typedef WX_NSString* WXGLPixelFormat;

//#endif

//#endif /* __WXMAC__ */

///* ABX: check __WIN32__ instead of __WXMSW__ for the same MSWBase in any Win32 port */
//#if defined(__WIN32__)

///*  Stand-ins for Windows types to avoid #including all of windows.h */

//#ifndef NO_STRICT
//    #define WX_MSW_DECLARE_HANDLE(type) typedef struct type##__ * WX##type
//#else
//    #define WX_MSW_DECLARE_HANDLE(type) typedef void * WX##type
//#endif

//typedef void* WXHANDLE;
//WX_MSW_DECLARE_HANDLE(HWND);
//WX_MSW_DECLARE_HANDLE(HICON);
//WX_MSW_DECLARE_HANDLE(HFONT);
//WX_MSW_DECLARE_HANDLE(HMENU);
//WX_MSW_DECLARE_HANDLE(HPEN);
//WX_MSW_DECLARE_HANDLE(HBRUSH);
//WX_MSW_DECLARE_HANDLE(HPALETTE);
//WX_MSW_DECLARE_HANDLE(HCURSOR);
//WX_MSW_DECLARE_HANDLE(HRGN);
//WX_MSW_DECLARE_HANDLE(RECTPTR);
//WX_MSW_DECLARE_HANDLE(HACCEL);
//WX_MSW_DECLARE_HANDLE(HINSTANCE);
//WX_MSW_DECLARE_HANDLE(HBITMAP);
//WX_MSW_DECLARE_HANDLE(HIMAGELIST);
//WX_MSW_DECLARE_HANDLE(HGLOBAL);
//WX_MSW_DECLARE_HANDLE(HDC);
//typedef WXHINSTANCE WXHMODULE;

//#undef WX_MSW_DECLARE_HANDLE

//typedef unsigned int    WXUINT;
//typedef unsigned long   WXDWORD;
//typedef unsigned short  WXWORD;

//typedef unsigned long   WXCOLORREF;
//typedef void *          WXRGNDATA;
//typedef struct tagMSG   WXMSG;
//typedef void *          WXHCONV;
//typedef void *          WXHKEY;
//typedef void *          WXHTREEITEM;

//typedef void *          WXDRAWITEMSTRUCT;
//typedef void *          WXMEASUREITEMSTRUCT;
//typedef void *          WXLPCREATESTRUCT;

//#ifdef __WXMSW__
//typedef WXHWND          WXWidget;
//#endif

//#ifdef __WIN64__
//typedef unsigned __int64   WXWPARAM;
//typedef __int64            WXLPARAM;
//typedef __int64            WXLRESULT;
//#else
//typedef wxW64 unsigned int WXWPARAM;
//typedef wxW64 long         WXLPARAM;
//typedef wxW64 long         WXLRESULT;
//#endif

//#if defined(__GNUWIN32__) || defined(__WXMICROWIN__)
//typedef int             (*WXFARPROC)();
//#else
//typedef int             (__stdcall *WXFARPROC)();
//#endif
//#endif /*  __WIN32__ */

//#if defined(__OS2__)
//typedef unsigned long   DWORD;
//typedef unsigned short  WORD;
//#endif

//#if defined(__WXPM__) || defined(__EMX__)
//#ifdef __WXPM__
///*  Stand-ins for OS/2 types, to avoid #including all of os2.h */
//typedef unsigned long   WXHWND;
//typedef unsigned long   WXHANDLE;
//typedef unsigned long   WXHICON;
//typedef unsigned long   WXHFONT;
//typedef unsigned long   WXHMENU;
//typedef unsigned long   WXHPEN;
//typedef unsigned long   WXHBRUSH;
//typedef unsigned long   WXHPALETTE;
//typedef unsigned long   WXHCURSOR;
//typedef unsigned long   WXHRGN;
//typedef unsigned long   WXHACCEL;
//typedef unsigned long   WXHINSTANCE;
//typedef unsigned long   WXHMODULE;
//typedef unsigned long   WXHBITMAP;
//typedef unsigned long   WXHDC;
//typedef unsigned int    WXUINT;
//typedef unsigned long   WXDWORD;
//typedef unsigned short  WXWORD;

//typedef unsigned long   WXCOLORREF;
//typedef void *          WXMSG;
//typedef unsigned long   WXHTREEITEM;

//typedef void *          WXDRAWITEMSTRUCT;
//typedef void *          WXMEASUREITEMSTRUCT;
//typedef void *          WXLPCREATESTRUCT;

//typedef WXHWND          WXWidget;
//#endif
//#ifdef __EMX__
///* Need a well-known type for WXFARPROC
//   below. MPARAM is typedef'ed too late. */
//#define WXWPARAM        void *
//#define WXLPARAM        void *
//#else
//#define WXWPARAM        MPARAM
//#define WXLPARAM        MPARAM
//#endif
//#define RECT            RECTL
//#define LOGFONT         FATTRS
//#define LOWORD          SHORT1FROMMP
//#define HIWORD          SHORT2FROMMP

//typedef unsigned long   WXMPARAM;
//typedef unsigned long   WXMSGID;
//typedef void*           WXRESULT;
///* typedef int             (*WXFARPROC)(); */
///*  some windows handles not defined by PM */
//typedef unsigned long   HANDLE;
//typedef unsigned long   HICON;
//typedef unsigned long   HFONT;
//typedef unsigned long   HMENU;
//typedef unsigned long   HPEN;
//typedef unsigned long   HBRUSH;
//typedef unsigned long   HPALETTE;
//typedef unsigned long   HCURSOR;
//typedef unsigned long   HINSTANCE;
//typedef unsigned long   HIMAGELIST;
//typedef unsigned long   HGLOBAL;
//#endif /*  WXPM || EMX */

//#if defined (__WXPM__)
///*  WIN32 graphics types for OS/2 GPI */

///*  RGB under OS2 is more like a PALETTEENTRY struct under Windows so we need a real RGB def */
//#define OS2RGB(r,g,b) ((DWORD)((unsigned char)(b) | ((unsigned char)(g) << 8)) | ((unsigned char)(r) << 16))

//typedef unsigned long COLORREF;
//#define GetRValue(rgb) ((unsigned char)((rgb) >> 16))
//#define GetGValue(rgb) ((unsigned char)(((unsigned short)(rgb)) >> 8))
//#define GetBValue(rgb) ((unsigned char)(rgb))
//#define PALETTEINDEX(i) ((COLORREF)(0x01000000 | (DWORD)(WORD)(i)))
//#define PALETTERGB(r,g,b) (0x02000000 | OS2RGB(r,g,b))
///*  OS2's RGB/RGB2 is backwards from this */
//typedef struct tagPALETTEENTRY
//{
//    char bRed;
//    char bGreen;
//    char bBlue;
//    char bFlags;
//} PALETTEENTRY;
//typedef struct tagLOGPALETTE
//{
//    WORD palVersion;
//    WORD palNumentries;
//    WORD PALETTEENTRY[1];
//} LOGPALETTE;

//#if (defined(__VISAGECPP__) && (__IBMCPP__ < 400)) || defined (__WATCOMC__)
//    /*  VA 3.0 for some reason needs base data types when typedefing a proc proto??? */
//typedef void* (_System *WXFARPROC)(unsigned long, unsigned long, void*, void*);
//#else
//#if defined(__EMX__) && !defined(_System)
//#define _System
//#endif
//typedef WXRESULT (_System *WXFARPROC)(WXHWND, WXMSGID, WXWPARAM, WXLPARAM);
//#endif

//#endif /* __WXPM__ */

//#if defined(__WXMOTIF__) || defined(__WXX11__)
///* Stand-ins for X/Xt/Motif types */
//typedef void*           WXWindow;
//typedef void*           WXWidget;
//typedef void*           WXAppContext;
//typedef void*           WXColormap;
//typedef void*           WXColor;
//typedef void            WXDisplay;
//typedef void            WXEvent;
//typedef void*           WXCursor;
//typedef void*           WXPixmap;
//typedef void*           WXFontStructPtr;
//typedef void*           WXGC;
//typedef void*           WXRegion;
//typedef void*           WXFont;
//typedef void*           WXImage;
//typedef void*           WXFontList;
//typedef void*           WXFontSet;
//typedef void*           WXRendition;
//typedef void*           WXRenderTable;
//typedef void*           WXFontType; /* either a XmFontList or XmRenderTable */
//typedef void*           WXString;

//typedef unsigned long   Atom;  /* this might fail on a few architectures */
//typedef long            WXPixel; /* safety catch in src/motif/colour.cpp */

//#endif /*  Motif */

//#ifdef __WXGTK__

///* Stand-ins for GLIB types */
//typedef struct _GSList GSList;

///* Stand-ins for GDK types */
//typedef struct _GdkColor        GdkColor;
//typedef struct _GdkCursor       GdkCursor;
//typedef struct _GdkDragContext  GdkDragContext;

//#if defined(__WXGTK20__)
//    typedef struct _GdkAtom* GdkAtom;
//#else
//    typedef unsigned long GdkAtom;
//#endif

//#if !defined(__WXGTK3__)
//    typedef struct _GdkColormap GdkColormap;
//    typedef struct _GdkFont GdkFont;
//    typedef struct _GdkGC GdkGC;
//    typedef struct _GdkRegion GdkRegion;
//#endif

//#if defined(__WXGTK3__)
//    typedef struct _GdkWindow GdkWindow;
//#elif defined(__WXGTK20__)
//    typedef struct _GdkDrawable GdkWindow;
//    typedef struct _GdkDrawable GdkPixmap;
//#else
//    typedef struct _GdkWindow GdkWindow;
//    typedef struct _GdkWindow GdkBitmap;
//    typedef struct _GdkWindow GdkPixmap;
//#endif

///* Stand-ins for GTK types */
//typedef struct _GtkWidget         GtkWidget;
//typedef struct _GtkRcStyle        GtkRcStyle;
//typedef struct _GtkAdjustment     GtkAdjustment;
//typedef struct _GtkToolbar        GtkToolbar;
//typedef struct _GtkNotebook       GtkNotebook;
//typedef struct _GtkNotebookPage   GtkNotebookPage;
//typedef struct _GtkAccelGroup     GtkAccelGroup;
//typedef struct _GtkSelectionData  GtkSelectionData;
//typedef struct _GtkTextBuffer     GtkTextBuffer;
//typedef struct _GtkRange          GtkRange;
//typedef struct _GtkCellRenderer   GtkCellRenderer;

//typedef GtkWidget *WXWidget;

//#ifndef __WXGTK20__
//#define GTK_OBJECT_GET_CLASS(object) (GTK_OBJECT(object)->klass)
//#define GTK_CLASS_TYPE(klass) ((klass)->type)
//#endif

//#endif /*  __WXGTK__ */

//#if defined(__WXGTK20__) || (defined(__WXX11__) && wxUSE_UNICODE)
//#define wxUSE_PANGO 1
//#else
//#define wxUSE_PANGO 0
//#endif

//#if wxUSE_PANGO
///* Stand-ins for Pango types */
//typedef struct _PangoContext         PangoContext;
//typedef struct _PangoLayout          PangoLayout;
//typedef struct _PangoFontDescription PangoFontDescription;
//#endif

//#ifdef __WXDFB__
///* DirectFB doesn't have the concept of non-TLW window, so use
//   something arbitrary */
//typedef const void* WXWidget;
//#endif /*  DFB */

///*  This is required because of clashing macros in windows.h, which may be */
///*  included before or after wxWidgets classes, and therefore must be */
///*  disabled here before any significant wxWidgets headers are included. */
//#ifdef __cplusplus
//#ifdef __WINDOWS__
//#include "wx/msw/winundef.h"
//#endif /* __WINDOWS__ */
//#endif /* __cplusplus */

///*  include the feature test macros */
//#include "wx/features.h"

///*  --------------------------------------------------------------------------- */
///*  macros to define a class without copy ctor nor assignment operator */
///*  --------------------------------------------------------------------------- */

//#define wxDECLARE_NO_COPY_CLASS(classname)      \
//    private:                                    \
//        classname(const classname&);            \
//        classname& operator=(const classname&)

//#define wxDECLARE_NO_COPY_TEMPLATE_CLASS(classname, arg)  \
//    private:                                              \
//        classname(const classname<arg>&);                 \
//        classname& operator=(const classname<arg>&)

//#define wxDECLARE_NO_COPY_TEMPLATE_CLASS_2(classname, arg1, arg2) \
//    private:                                                      \
//        classname(const classname<arg1, arg2>&);                  \
//        classname& operator=(const classname<arg1, arg2>&)

//#define wxDECLARE_NO_ASSIGN_CLASS(classname)    \
//    private:                                    \
//        classname& operator=(const classname&)

///* deprecated variants _not_ requiring a semicolon after them */
//#define DECLARE_NO_COPY_CLASS(classname) \
//    wxDECLARE_NO_COPY_CLASS(classname);
//#define DECLARE_NO_COPY_TEMPLATE_CLASS(classname, arg) \
//    wxDECLARE_NO_COPY_TEMPLATE_CLASS(classname, arg);
//#define DECLARE_NO_ASSIGN_CLASS(classname) \
//    wxDECLARE_NO_ASSIGN_CLASS(classname);

///*  --------------------------------------------------------------------------- */
///*  If a manifest is being automatically generated, add common controls 6 to it */
///*  --------------------------------------------------------------------------- */

//#if wxUSE_GUI && \
//    (!defined wxUSE_NO_MANIFEST || wxUSE_NO_MANIFEST == 0 ) && \
//    ( defined _MSC_FULL_VER && _MSC_FULL_VER >= 140040130 )

//#define WX_CC_MANIFEST(cpu)                     \
//    "/manifestdependency:\"type='win32'         \
//     name='Microsoft.Windows.Common-Controls'   \
//     version='6.0.0.0'                          \
//     processorArchitecture='" cpu "'            \
//     publicKeyToken='6595b64144ccf1df'          \
//     language='*'\""

//#if defined _M_IX86
//    #pragma comment(linker, WX_CC_MANIFEST("x86"))
//#elif defined _M_X64
//    #pragma comment(linker, WX_CC_MANIFEST("amd64"))
//#elif defined _M_IA64
//    #pragma comment(linker, WX_CC_MANIFEST("ia64"))
//#else
//    #pragma comment(linker, WX_CC_MANIFEST("*"))
//#endif

//#endif /* !wxUSE_NO_MANIFEST && _MSC_FULL_VER >= 140040130 */

///* wxThread and wxProcess priorities */
//enum
//{
//    wxPRIORITY_MIN     = 0u,   /* lowest possible priority */
//    wxPRIORITY_DEFAULT = 50u,  /* normal priority */
//    wxPRIORITY_MAX     = 100u  /* highest possible priority */
//};

//#endif
//    /*  _WX_DEFS_H_ */
