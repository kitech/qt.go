package main

/*
 */
import "C"
import (
	"log"

	"qt.go/cffiqt"
	"qt.go/qtcore"
	// "github.com/therecipe/qt/core"
)

func main() {
	// rs := qtrt.CString("hehehg呵呵")
	txt := qtcore.NewQString_5("hehehg呵呵")
	log.Println(txt.Length(), txt.IsEmpty())
	log.Println(txt.GetCthis())
	log.Println(txt.ToLatin1())
	log.Println(txt.ToLocal8Bit().Data())

	argv := []string{"./guiapp", "-v", "-x"}
	app := qtcore.NewQCoreApplication(len(argv), argv, 0)
	log.Println(app)

	tmer := qtcore.NewQTimer(nil)
	log.Println(tmer)
	tmer.SetInterval(3000)
	tmer.Start_1()

	dyslot := ffiqt.NewQDynSlotObject("abc", 123)
	if false {
		dyslot.Connect(tmer, "timeout()", func(a int, b byte, d string, c *qtcore.QString) {
			log.Println("hehehhe")
		})

		dyslot.Connect(tmer, "timeout()", func(a []string) {
			log.Println("hehehhe222")
		})
	}

	dyslot.Connect(tmer, "objectNameChanged(const QString&)", func(s *qtcore.QString) {
		log.Println("hehehhe333", s, s.Length(), s.ToLocal8Bit().Data())
	})

	newname := qtcore.NewQString_5("testt禾tttttttt")
	tmer.SetObjectName(newname)

	app.Exec()
}
