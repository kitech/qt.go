package main

import (
	"bytes"
	"fmt"
	"gopp"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/ianlancetaylor/demangle"
)

var lang = "go" // v

// generate symbols for qt.go/qtrel
// list symbols from libQt5Inline.so
func main() {
	if len(os.Args) > 1 {
		lang = os.Args[1]
	}

	sofile := "/usr/lib/libQt5Inline.so"
	cmdo := exec.Command("objdump", "-T", sofile)
	bcc, err := cmdo.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	lines := strings.Split(string(bcc), "\n")
	// log.Println(lines)
	var lines2 []string
	var mthoverloads = make(map[string]int)

	for _, line := range lines {
		cols := strings.Fields(line)
		if len(cols) < 3 {
			log.Println("wtt", cols)
			continue
		}
		symname := cols[len(cols)-1]
		if !strings.HasPrefix(symname, "C_") {
			continue
		}
		// log.Println(len(cols), symname)
		_, err := demangle.ToString(symname[1:])
		if err != nil {
			log.Println("ignore", symname, err)
			continue
		}
		if true {
			lines2 = append(lines2, symname)
		}
	}

	log.Println(len(lines2))
	lines3 := sort.StringSlice(lines2)
	sort.Stable(lines3)
	var clsmths = make(map[string][]string) // class => mths
	for _, line := range lines3 {
		// log.Println(len(line), line)
		symproto, err := demangle.ToString(line[1:])
		if err != nil {
			log.Panicln("wtt", line)
		}
		lparenpos := strings.Index(symproto, "(")
		colonpos := strings.Index(symproto, ":")
		if false {
			log.Println(line, symproto, lparenpos, colonpos, len(symproto))
		}
		fullname := symproto[:lparenpos]
		olcnt, ok := mthoverloads[fullname]
		if ok {
			mthoverloads[fullname] = olcnt + 1
		} else {
			mthoverloads[fullname] = 1
		}
		if strings.HasPrefix(fullname, "operator") {
			continue // ignore now
		}
		if strings.ContainsAny(fullname, "<]") {
			continue
		}
		log.Println(symproto, fullname, olcnt)
		pair := strings.Split(fullname, "::")
		if len(pair) != 2 {
			log.Println("wtt", symproto, fullname)
			continue
		}
		if strings.HasPrefix(pair[1], "operator") {
			continue
		}
		if pair[1] == "qt_metacall" || pair[1] == "qt_metacall" {
			continue
		}

		if pair[0] == pair[1] {
			pair[1] = "new"
		}
		if strings.HasPrefix(pair[1], "~") {
			pair[1] = "dtor"
		}
		ovsfx := fmt.Sprintf("%d", olcnt)
		if olcnt == 0 {
			ovsfx = ""
		}

		newname := fmt.Sprintf("%s%s%s", pair[0], strings.Title(pair[1]), ovsfx)
		if newname == "QWidgetActionEvent" {
			continue
		}
		clsmths[pair[0]] = append(clsmths[pair[0]],
			fmt.Sprintf("%s %s%s", line, pair[1], ovsfx))
	}

	modhdr := gopp.IfElseStr(lang == "go", "package", "module") + " qtrel\n\n"
	cbuf := bytes.NewBuffer([]byte(modhdr))
	//genwith_general(cbuf, clsmths) // go compiler like this
	genwith_bigstr(cbuf, clsmths) // v compiler like this

	if true {
		ioutil.WriteFile("ccsymlst3."+lang, cbuf.Bytes(), 0644)
	}
}

func genwith_general(cbuf *bytes.Buffer, clsmths map[string][]string) {
	for cls, mths := range clsmths {
		line := ""
		for _, mth := range mths {
			part := strings.Split(mth, " ")
			if !strings.HasPrefix(part[0], "C_ZN") {
				log.Panicln("wtt", part)
			}
			if lang == "go" {
				line += fmt.Sprintf("const (%s%s = \"%s\")\n",
					cls, strings.Title(part[1]), part[0][4:])
			} else if lang == "v" {
				line += fmt.Sprintf("const (%s%s = \"%s\")\n",
					cls, strings.Title(part[1]), part[0][4:])
				//line += fmt.Sprintf("const (%s%s = \"%s\")\n",
				//	strings.ToLower(cls), strings.ToLower(part[1]), part[0][4:])
			}
		}
		cbuf.Write([]byte(line))
	}
}

func genwith_dotns(cbuf *bytes.Buffer, clsmths map[string][]string) {
	const dotns = false
	for cls, mths := range clsmths {
		line := ""
		for _, mth := range mths {
			part := strings.Split(mth, " ")
			if !strings.HasPrefix(part[0], "C_ZN") {
				log.Panicln("wtt", part)
			}
			if dotns {
				line += fmt.Sprintf("type T%s byte\n", part[0][1:])
			} else {
				line += fmt.Sprintf("const %s%s = \"%s\"\n",
					cls, strings.Title(part[1]), part[0][4:])
			}
		}
		if !dotns {
			cbuf.Write([]byte(line))
			continue
		}
		line += fmt.Sprintf("var %s = struct {\n", cls)
		for _, mth := range mths {
			part := strings.Split(mth, " ")
			line += fmt.Sprintf("  %s T%s\n", strings.Title(part[1]), part[0][1:])
		}
		line += "} {\n"
		line += "}\n\n"
		cbuf.Write([]byte(line))
		// newline := fmt.Sprintf("var %s  = \"%s\"\n", newname, line[1:])
		// cbuf.Write([]byte(newline))
	}
}
func genwith_typealias(cbuf *bytes.Buffer, clsmths map[string][]string) {
}
func genwith_bigstr(cbuf *bytes.Buffer, clsmths map[string][]string) {
	bigstr := ""
	for cls, mths := range clsmths {
		line := ""
		for _, mth := range mths {
			part := strings.Split(mth, " ")
			line += fmt.Sprintf("const (%s%s = %d)\n",
				cls, strings.Title(part[1]), len(bigstr))
			bigstr += part[0][4:] + "\n"
		}
		cbuf.Write([]byte(line))
	}
	cbuf.Write([]byte("const (symdat = \""))
	cbuf.Write([]byte(bigstr))
	cbuf.Write([]byte("\")"))
}

func symcrc32(s string) uint32 { return crc32.ChecksumIEEE([]byte(s)) }
