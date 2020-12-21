package qtwidgets

import (
	"os"

	"github.com/kitech/qt.go/qtrt"
)

var appargc int32 // cannot free, qt一直持有其引用
func NewQApplication() *QApplication {
	appargc = int32(len(os.Args))
	var convArg0 = Voidptr(&appargc)
	var convArg1 = qtrt.StringSliceToCCharPP(os.Args)
	var arg2 = 0
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(3158502489, "_ZN12QApplicationC2ERiPPci", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&arg2))
	qtrt.ErrPrint(err, rv)
	gothis := QApplicationFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QApplication")
	return gothis
}

// 阻塞的调用不能用asmcgocall,否则回调崩溃
func (this *QApplication) Exec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QApplication4execEv", qtrt.FFITY_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
	return 0
}
