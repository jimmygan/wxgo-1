package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const CType_WxObjectPtr = "WxObjectPtr"
const CFunc_Ptr = "__PTR"

var CFuncDeclType2CallTypeConvFormat = map[string]string{
	"Size*":  "*((wxSize*)%v)",
	"Rect*":  "*((wxSize*)%v)",
	"Point*": "*((wxPoint*)%v)",
	"String": "NewWxString(%v)",
}
var CFuncReturnType2GoTypeConvFormat = map[string]string{
	"Size":   "wxSize sizeVar = %v;\n\treturn (*(Size*)&sizeVar);",
	"Rect":   "wxRect rectVar = %v;\n\treturn (*(Rect*)&rectVar);",
	"Point*": "wxPoint ptVar = %v;\n\treturn (*(Point*)&ptVar)",
	"String": "return NewGoString(%v);",
}

func init() {
	flag.Parse()
}

var regexCFuncDecl = regexp.MustCompile(`(\w+\*?)\s+(\w+)\s*\((.*?)\)\s*(;?)`)
var regexCFuncDeclParams = regexp.MustCompile(`(\w+\*?)\s+(\w+)(?:\,(\w+\*?)\s+(\w+))?`)

func main() {
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Wrong number of args. 1 arg needed: the .h file path.")
		os.Exit(1)
	}
	hFilePath := flag.Arg(0)
	hFile, err := os.Open(hFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open file: %v\n", hFilePath)
		os.Exit(2)
	}
	defer hFile.Close()

	var cFuncImplBuf bytes.Buffer

	hFileScanner := bufio.NewScanner(hFile)
	var lineNumber = 1
	for hFileScanner.Scan() {
		line := strings.TrimSpace(hFileScanner.Text())
		ok, f := parseCFuncDecl(line, lineNumber)
		if !ok {
			continue
		}
		cFuncImplBuf.WriteString(genCFuncImpl(f))
		lineNumber++
	}
	if err = hFileScanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Read file %v error: %v\n", err)
		os.Exit(3)
	}
	// Write cpp file
	hExtLen := len(filepath.Ext(hFilePath))
	hFileName := filepath.Base(hFilePath)
	cppFileName := hFileName[:len(hFileName)-hExtLen] + ".cpp"
	fmt.Printf("Writing %v ...\n", cppFileName)
	cppFile, err := os.OpenFile(cppFileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open file to write: %v %v\n", cppFileName, err)
		os.Exit(100)
	}
	defer cppFile.Close()
	if _, err := cppFile.Write(cFuncImplBuf.Bytes()); err != nil {
		fmt.Fprintf(os.Stderr, "Can't write to file: %v %v\n", cppFileName, err)
	}
	fmt.Println("Done!")
}

// ""(wxWindow*)param"
func getCFuncCallParam(param *Param) string {
	convFmt := CFuncDeclType2CallTypeConvFormat[param.Type]
	if convFmt == "" {
		// Determine the type conversion by guess work.
		switch param.Name {
		case "window", "parent":
			convFmt = "(wxWindow*)%v"
		case "validator":
			convFmt = "*(wxValidator*)%v"
		case "event":
			convFmt = "*(wxEvent*)%v"
		}
	}
	if convFmt != "" {
		return fmt.Sprintf(convFmt, param.Name)
	}
	return param.Name
}

// "(wxWindow*)parent, id, flag, NewWxString(name)"
func genCFuncCallParams(params []*Param) string {
	var funcCallParams []string
	for _, p := range params {
		funcCallParams = append(funcCallParams, getCFuncCallParam(p))
	}
	return strings.Join(funcCallParams, ", ")
}

// "__PTR->MethodName((conv1)p1, p2, p3)""
func genCMethodCallExpr(f FuncDecl) string {
	typeMethod := strings.Split(f.Name, "_")
	if len(typeMethod) != 2 {
		fmt.Fprintf(os.Stderr, "Warning: Invalid function name: %v. Underscore missing.\n", f.Name)
	}

	var callParams []*Param
	if len(f.Params) < 1 || f.Params[0].Name != "p" || f.Params[0].Type != CType_WxObjectPtr {
		fmt.Fprintf(os.Stderr, "Warning: Invalid signature: %v. parameter 'p' is missing.\n", f.Name)
	} else {
		callParams = f.Params[1:]
	}
	return fmt.Sprintf("%v->%v(%v)", CFunc_Ptr, typeMethod[1], genCFuncCallParams(callParams))
}

// "new wxYYY(...)"
func genCFuncCtorCallExpr(f FuncDecl, typeName string) string {
	return fmt.Sprintf("new %v(%v)", typeName, genCFuncCallParams(f.Params))
}

// "return ...;"
func genCFuncCallBody(f FuncDecl) string {
	needReturn := f.ReturnType != "void"
	var expr string
	isConstructor := f.ReturnType == CType_WxObjectPtr && strings.HasSuffix(f.Name, "_New")
	if isConstructor {
		typeName := f.Name[:len(f.Name)-4]
		expr = genCFuncCtorCallExpr(f, typeName)
	} else {
		expr = genCMethodCallExpr(f)
	}

	if needReturn {
		format := CFuncReturnType2GoTypeConvFormat[f.ReturnType]
		if format == "" {
			format = "return %v;"
		}
		return fmt.Sprintf(format, expr)
	} else {
		return expr + ";"
	}
}

func genCFuncImpl(f FuncDecl) string {
	// Param: "type name"
	var params []string
	for _, p := range f.Params {
		params = append(params, p.Type+" "+p.Name)
	}

	return fmt.Sprintf(`%v %v(%v) {
	%v
}
`, f.ReturnType, f.Name, strings.Join(params, ", "), genCFuncCallBody(f))
}

type Param struct {
	Type, Name string
}

type FuncDecl struct {
	Name, ReturnType string
	Params           []*Param
}

func parseCFuncDecl(line string, lineNumber int) (ok bool, f FuncDecl) {
	if ok = regexCFuncDecl.MatchString(line); !ok {
		return
	}
	if m := regexCFuncDecl.FindStringSubmatch(line); len(m) == 0 {
		return // Not a C function declaration.
	} else {
		if m[4] == "" { // Warning missing semicolon.
			fmt.Printf("Warning: line %v: \"%v\" No semicolon before end of line!\n", lineNumber, line)
		}
		f.ReturnType = m[1]
		f.Name = m[2]
		ok, f.Params = parseCFuncDeclParam(strings.TrimSpace(m[3]))
		if !ok {
			return
		}

	}
	return
}

func parseCFuncDeclParam(paramStr string) (ok bool, params []*Param) {
	if len(paramStr) == 0 {
		ok = true // No param is not an error.
		return
	}
	if ok = regexCFuncDeclParams.MatchString(paramStr); !ok {
		return
	}
	if m := regexCFuncDeclParams.FindAllStringSubmatch(paramStr, -1); len(m) == 0 {
		return
	} else {
		for _, p := range m { // One param
			params = append(params, &Param{Type: p[1], Name: p[2]})
		}
	}
	ok = true
	return
}
