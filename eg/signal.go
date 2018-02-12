package main

/*
 */
import "C"
import (
	"fmt"
	"log"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

func main() {
	qtrt.SetDebugDynSlot(true)
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
	// dyslot := qtrt.NewQDynSlotObject("abc", 123)
	if true {
		/*
			qtrt.Connect(tmer, "timeout()", func(a int, b byte, d string, c *qtcore.QString) {
				log.Println("hehehhe")
			})

			qtrt.Connect(tmer, "timeout()", func(a []string) {
				log.Println("hehehhe222")
			})
		*/
		qtrt.Connect(tmer, "timeout()", func() {
			cnter++
			log.Println("hehehhe222")
			tmer.SetObjectName(fmt.Sprintf("testt禾tttttttt-%d", cnter))
		})
	}

	tmer.InheritTimerEvent(func(arg0 *qtcore.QTimerEvent) {
		log.Println(arg0)
	})
	qtrt.Connect(tmer, "objectNameChanged(const QString&)", func(s *qtcore.QString) {
		log.Println("hehehhe333", s, s.Length(), s.ToLocal8Bit().Data())
	})

	tmer.SetObjectName("testt禾tttttttt")

	app.Exec()
}
