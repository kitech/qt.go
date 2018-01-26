package main

import (
	"fmt"
	"gopp"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var cpplines = make([]string, 0)
var hdrlines = make([]string, 0)
var golines = make([]string, 0)

var class string
var wclass string // setupUi中用到的widget class
var file string
var filep string // file name without externsion

// usage: go-uic <ui_foo.h>
func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	file = os.Args[1]
	filep = file[:len(file)-2]
	log.Println(file, filep)

	bcc, err := ioutil.ReadFile(file)
	gopp.ErrPrint(err, file)
	lines := strings.Split(string(bcc), "\n")
	log.Println(len(lines))

	inmem := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "class ") {
			class = strings.Split(line, " ")[1]
			continue
		}
		if strings.HasPrefix(line, "public:") {
			inmem = true
			continue
		}
		if strings.HasPrefix(line, "void setupUi") {
			wclass = strings.Split(strings.Split(line, "(")[1], " *")[0]
			onClass(class)
			break
		}

		if inmem {
			// do smth...
			// log.Println(line)
			transformMember(line, class)
		}
	}

	saveCode()
}

func transformMember(line string, class string) {
	parts := strings.Split(line, "*")
	mname := parts[1][:len(parts[1])-1]
	log.Println("transform member...", mname, line)

	var reline string

	reline = fmt.Sprintf("void* %s_Get_%s(void*p);", class, mname)
	hdrlines = append(hdrlines, reline)

	reline = fmt.Sprintf("void* %s_Get_%s(void*p){return ((%s*)p)->%s;}",
		class, mname, class, mname)
	cpplines = append(cpplines, reline)

	reline = fmt.Sprintf("func %s_Get_%s(p unsafe.Pointer) unsafe.Pointer {return C.%s_Get_%s(p)}",
		class, mname, class, mname)
	golines = append(golines, reline)
}

func onClass(class string) {
	var reline string

	log.Printf("generating %s_new...\n", class)
	reline = fmt.Sprintf("void* %s_new();", class)
	hdrlines = append(hdrlines, reline)
	reline = fmt.Sprintf("void* %s_new(){return new %s();}", class, class)
	log.Println(reline)
	cpplines = append(cpplines, reline)
	reline = fmt.Sprintf("func %s_new() unsafe.Pointer{return C.%s_new()}", class, class)
	golines = append(golines, reline)

	//
	log.Printf("generating %s_setupUi...\n", class)
	reline = fmt.Sprintf("void %s_setupUi(void* p, void* w);", class)
	hdrlines = append(hdrlines, reline)
	reline = fmt.Sprintf("void %s_setupUi(void* p, void* w){((%s*)p)->setupUi((%s*)w);}",
		class, class, wclass)
	cpplines = append(cpplines, reline)
	reline = fmt.Sprintf("func %s_setupUi(p unsafe.Pointer, w unsafe.Pointer) {C.%s_setupUi(p, w)}", class, class)
	golines = append(golines, reline)

	//
	log.Printf("generating %s_retranslateUi...\n", class)
	reline = fmt.Sprintf("void %s_retranslateUi(void* p, void* w);", class)
	hdrlines = append(hdrlines, reline)
	reline = fmt.Sprintf("void %s_retranslateUi(void* p, void* w){((%s*)p)->retranslateUi((%s*)w);}",
		class, class, wclass)
	cpplines = append(cpplines, reline)
	reline = fmt.Sprintf("func %s_retranslateUi(p unsafe.Pointer, w unsafe.Pointer) {C.%s_retranslateUi(p, w)}", class, class)
	golines = append(golines, reline)
}

func saveCode() {
	var code string
	var mod os.FileMode = 0644

	log.Printf("saving... %s_c.h...\n", filep)
	code = fmt.Sprintf("#ifndef _%s_C_H_\n", class)
	code += fmt.Sprintf("#define _%s_C_H_\n", class)
	code += fmt.Sprintf("#ifdef __cplusplus\n")
	code += fmt.Sprintf("extern \"C\" {\n")
	code += fmt.Sprintf("#endif\n\n")
	code += strings.Join(hdrlines, "\n")
	code += fmt.Sprintf("\n\n#ifdef __cplusplus\n")
	code += fmt.Sprintf("}\n")
	code += fmt.Sprintf("#endif\n")
	code += fmt.Sprintf("\n#endif // %s", class)
	ioutil.WriteFile(fmt.Sprintf("%s_c.h", filep), []byte(code), mod)

	log.Printf("saving... %s.cpp...\n", filep)
	code = fmt.Sprintf("#include \"%s_c.h\"\n", filep)
	code += fmt.Sprintf("#include \"%s.h\"\n", filep)
	code += strings.Join(cpplines, "\n")
	ioutil.WriteFile(fmt.Sprintf("%s.cpp", filep), []byte(code), mod)

	log.Printf("saving... %s.go...\n", filep)
	code = "package main\n"
	code += "/*\n"
	code += fmt.Sprintf("#include \"%s_c.h\"\n", filep)
	code += "*/\n"
	code += "import \"C\"\n"
	code += "import \"unsafe\"\n"
	code += strings.Join(golines, "\n")
	ioutil.WriteFile(fmt.Sprintf("%s.go", filep), []byte(code), mod)
}
