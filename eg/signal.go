package main

/*
 */
import "C"
import (
	"fmt"
	"log"

	ffiqt "qt.go/cffiqt"
	"qt.go/qtcore"
)

func main() {
	ffiqt.SetDebugDynSlot(true)
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

	cnter := 3
	// dyslot := ffiqt.NewQDynSlotObject("abc", 123)
	if true {
		/*
			ffiqt.Connect(tmer, "timeout()", func(a int, b byte, d string, c *qtcore.QString) {
				log.Println("hehehhe")
			})

			ffiqt.Connect(tmer, "timeout()", func(a []string) {
				log.Println("hehehhe222")
			})
		*/
		ffiqt.Connect(tmer, "timeout()", func() {
			newname := qtcore.NewQString_5(fmt.Sprintf("testt禾tttttttt-%d", cnter))
			cnter++
			log.Println("hehehhe222")
			tmer.SetObjectName(newname)
		})
	}

	tmer.InheritTimerEvent(func(arg0 *qtcore.QTimerEvent) {
		log.Println(arg0)
	})
	ffiqt.Connect(tmer, "objectNameChanged(const QString&)", func(s *qtcore.QString) {
		log.Println("hehehhe333", s, s.Length(), s.ToLocal8Bit().Data())
	})

	newname := qtcore.NewQString_5("testt禾tttttttt")
	tmer.SetObjectName(newname)

	app.Exec()
}
