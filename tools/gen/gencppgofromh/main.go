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
var CType2GoType = map[string]string{
	"Size":   "Size",
	"Rect":   "Rect",
	"Point":  "Point",
	"String": "string",
	"float":  "float32",
	"long":   "int",
	"BOOL":   "bool",
}
var CType2CGoCallTypeConvFormat = map[string]string{
	"String":          "cString(&%v)",
	"BOOL":            "cBool(%v)",
	"Size*":           "(*C.Size)(%v)",
	"Point*":          "(*C.Point)(%v)",
	"Rect*":           "(*C.Rect)(%v)",
	CType_WxObjectPtr: "%v.wxPtr()",
}
var CType2GoReturnTypeConvFormat = map[string]string{
	"String": "return goString(%v)",
	"BOOL":   "return goBool(%v)",
	"Size":   "return Size(%v)",
	"Point":  "return Point(%v)",
	"Rect":   "return Rect(%v)",
	CType_WxObjectPtr: `if obj := NewObjectFromPtr(%v, false); obj != nil {
		return obj.(%[1]v)
	}
	return nil`,
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
	var goMethodBuf bytes.Buffer
	var goCtorBuf bytes.Buffer
	var goInterfaceFuncBuf bytes.Buffer

	hFileScanner := bufio.NewScanner(hFile)
	var lineNumber = 1
	for hFileScanner.Scan() {
		line := strings.TrimSpace(hFileScanner.Text())
		ok, f := parseCFuncDecl(line, lineNumber)
		if !ok {
			continue
		}
		cFuncImplBuf.WriteString(genCFuncImpl(f))
		ctor, mtd, face := genGoMethod(&f)
		if ctor {
			goCtorBuf.WriteString(mtd)
		} else {
			goMethodBuf.WriteString(mtd)
			goInterfaceFuncBuf.WriteString("\t" + face + "\n")
		}
		lineNumber++
	}
	if err = hFileScanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Read file %v error: %v\n", err)
		os.Exit(3)
	}
	// Write cpp file
	hExtLen := len(filepath.Ext(hFilePath))
	hFileName := filepath.Base(hFilePath)
	hFileBaseName := hFileName[:len(hFileName)-hExtLen]
	cppFileName := hFileBaseName + ".cpp"
	fmt.Printf("Writing %v ...\n", cppFileName)
	cppFile, err := os.OpenFile(cppFileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open file to write: %v %v\n", cppFileName, err)
		os.Exit(100)
	}
	defer cppFile.Close()
	if _, err := cppFile.Write(cFuncImplBuf.Bytes()); err != nil {
		fmt.Fprintf(os.Stderr, "Can't write to file: %v %v\n", cppFileName, err)
		os.Exit(101)
	}
	fmt.Println("Done!")
	// Write go file
	goFileName := hFileBaseName + ".go"
	fmt.Printf("Writing %v ...\n", goFileName)
	goFile, err := os.OpenFile(goFileName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open file to write: %v %v\n", goFileName, err)
		os.Exit(100)
	}
	defer goFile.Close()

	var goBuf bytes.Buffer
	var typeName = hFileBaseName
	var interfaceName = strings.ToUpper(typeName[0:1]) + typeName[1:]
	fmt.Fprintf(&goBuf, `package wx

import (
	"reflect"
)

//#include "%s.h"
import "C"

func init() {
	mapType("wx%s", reflect.TypeOf(%[1]s{}))
}

type %[1]s struct {
	object
}

%[3]s

type %s interface {
	Object
%s
}

%s

`, typeName, interfaceName, goMethodBuf.String(), interfaceName, goInterfaceFuncBuf.String(), goCtorBuf.String())

	if _, err := goFile.Write(goBuf.Bytes()); err != nil {
		fmt.Fprintf(os.Stderr, "Can't write to file: %v %v\n", cppFileName, err)
		os.Exit(101)
	}
	fmt.Println("Done!")
}

func genGoMethod(f *FuncDecl) (ctor bool, method, interfaceFunc string) {
	tf := strings.Split(f.Name, "_")
	wxType := tf[0]
	interfaceName := wxType[2:]
	typeName := strings.ToLower(interfaceName[0:1]) + interfaceName[1:] // Lower case the first letter.
	methodName := tf[1]
	if strings.HasPrefix(methodName, "Get") {
		methodName = methodName[3:] // Trim off "Get"
	}
	var returnType string
	var params []*Param = f.Params
	if strings.HasSuffix(f.Name, "_New") {
		ctor = true
		methodName = "New" + interfaceName
		returnType = interfaceName
	} else {
		returnType = getGoMethodReturnType(f.ReturnType)
		params = params[1:]
	}
	interfaceFunc = fmt.Sprintf("%v(%v)%v", methodName, genGoMethodDeclParams(params), returnType)
	receiver := typeName[0:1]
	if !ctor { // Do not generate a method for ctor.
		method = fmt.Sprintf(`func (%v *%v) %v {
	p := wxPtr(%v)
	if p == nil {
		return %v
	}
	%v
}
`, receiver, typeName, interfaceFunc, receiver, getGoDefReturnValue(returnType), genGoMethodCGoCall(f.ReturnType, f.Name, params))
	} else {
		method = fmt.Sprintf(`func New_%v(%v) %v {
	%v := &%v{}
	%[4]v.bindWxPtr(C.%[6]v(%v), true)
	return %[4]v
}
`, interfaceName, genGoMethodDeclParams(params), interfaceName, receiver, typeName, f.Name, genCGoCallParams(f.Params, false))
	}
	return
}

func genGoMethodCGoCall(creturnType string, funcName string, params []*Param) string {
	expr := genCGoCallExpr(funcName, params)
	if creturnType == "void" {
		return expr
	} else {
		format := CType2GoReturnTypeConvFormat[creturnType]
		if format == "" {
			format = "%v"
		}
		return fmt.Sprintf(format, expr)
	}
}

func genCGoCallExpr(funcName string, params []*Param) string {
	return fmt.Sprintf("C.%v(%v)", funcName, genCGoCallParams(params, true))
}

// includeP Whether add "p" as the first parameter to cgo call. Set it to false
// to generate a ctor(NewXXX()XXX{}).
func genCGoCallParams(params []*Param, includeP bool) string {
	var ps []string
	if includeP {
		ps = append(ps, "p")
	}
	for _, p := range params {
		ps = append(ps, genCGoCallParam(p))
	}
	return strings.Join(ps, ", ")
}

func genCGoCallParam(param *Param) string {
	var format = CType2CGoCallTypeConvFormat[param.Type]
	if format == "" {
		format = fmt.Sprintf("C.%v(%%v)", param.Type)
	}
	return fmt.Sprintf(format, param.Name)
}

func getGoDefReturnValue(typeName string) string {
	switch typeName {
	case "":
		return ""
	case "int":
		return "0"
	case "bool":
		return "false"
	case "string":
		return `""`
	case "Size":
		return "DefaultSize"
	case "Point":
		return "DefaultPosition"
	case "Rect":
		return "DefaultRect"
	default:
		return "nil"
	}
}

func cTypeToGoType(ctype string) string {
	var ptr bool
	var baseType string = ctype
	if strings.HasSuffix(ctype, "*") {
		ptr = true
		baseType = ctype[:len(ctype)-1]
	}
	t := CType2GoType[baseType]
	if t == "" {
		return ""
	} else {
		if ptr {
			return "*" + t
		}
		return t
	}
}

func getGoMethodReturnType(returnType string) string {
	if returnType == "void" {
		return ""
	}
	goType := cTypeToGoType(returnType)
	if goType == "" {
		if strings.HasSuffix(returnType, "*") { // Pointer type
			goType = "*" + returnType[:len(returnType)-1]
		} else if goType == CType_WxObjectPtr {
			goType = "Object"
		} else {
			goType = returnType
		}
	}
	return goType
}

func getGoMethodDeclParam(param *Param) string {
	goType := cTypeToGoType(param.Type)
	if goType == "" {
		switch param.Type {
		case "WxObjectPtr":
			goType = strings.ToUpper(param.Name[0:1]) + param.Name[1:]
		default:
			if strings.HasSuffix(param.Type, "*") { // Pointer type
				goType = "*" + param.Type[:len(param.Type)-1]
			}
		}
	}
	if goType == "" {
		goType = param.Type
	}
	return param.Name + " " + goType
}

func genGoMethodDeclParams(params []*Param) string {
	var declParams []string
	for _, p := range params {
		declParams = append(declParams, getGoMethodDeclParam(p))
	}
	return strings.Join(declParams, ", ")
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
