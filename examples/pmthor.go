package main

/*
 */
import "C"
import (
	"log"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

// protected method override demo
func main() {
	qtrt.SetDebugDynSlot(true)
	argv := []string{"./guiapp", "-v", "-x"}
	app := qtcore.NewQCoreApplication(len(argv), argv, 0)
	log.Println(app)

	tmer := qtcore.NewQTimer(nil)
	log.Println(tmer)
	tmer.SetInterval(3000)
	tmer.Start1()

	tmer.InheritTimerEvent(func(arg0 *qtcore.QTimerEvent) {
		log.Println("protected method called:", arg0, arg0.Type(), arg0.TimerId())
	})

	app.Exec()
}
