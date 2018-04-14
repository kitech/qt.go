package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kitech/qt.go/toolutil"
)

var rcdir string = "."
var rcpref string = ""
var rcout string = "rcc"

func main() {
	flag.StringVar(&rcdir, "dir", ".", "dir of resources")
	flag.StringVar(&rcpref, "prefix", "", "prefix of resources")
	flag.StringVar(&rcout, "out", "rcc", "output file of resources")
	flag.Parse()

	cp := toolutil.NewCodePagerNoComment()
	cp.AddPointer("main")
	cp.AP("main", "<RCC><qresource>")

	if len(rcpref) > 1 {
		rcpref = strings.Trim(rcpref, "/")
	}
	orcdir := rcdir
	rcdir, _ = filepath.Abs(rcdir)
	log.Println(orcdir, "=>", rcdir)
	filepath.Walk(rcdir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		rcpath, err := filepath.Rel(rcdir, path)
		if err != nil {
			log.Println(err, path)
		}
		if rcpath == "." {
			return nil
		}

		frcpath := rcpath
		if rcpref != "" {
			frcpath = fmt.Sprintf("%s/%s", rcpref, rcpath)
		}
		cp.APf("main", "<file>%s</file>", frcpath)
		return nil
	})

	cp.AP("main", "</qresource></RCC>")
	// log.Println(cp.ExportAll())
	outFile := fmt.Sprintf("%s.qrc", rcout)
	ioutil.WriteFile(outFile, []byte(cp.ExportAll()), 0644)
}
