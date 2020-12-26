package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*
 */
// size 24
type QIconDragEvent struct {
	*qtcore.QEvent
}
type QIconDragEvent_ITF interface {
	qtcore.QEvent_ITF
	QIconDragEvent_PTR() *QIconDragEvent
}

func (ptr *QIconDragEvent) QIconDragEvent_PTR() *QIconDragEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QIconDragEventFromptr(cthis Voidptr) *QIconDragEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QIconDragEvent{bcthis0}
}
func (*QIconDragEvent) Fromptr(cthis Voidptr) *QIconDragEvent {
	return QIconDragEventFromptr(cthis)
}

func DeleteQIconDragEvent(this *QIconDragEvent) {
	rv, err := qtrt.Qtcc3(2673315635, "_ZN14QIconDragEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10089() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
