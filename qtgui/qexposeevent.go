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
// size 32
type QExposeEvent struct {
	*qtcore.QEvent
}
type QExposeEvent_ITF interface {
	qtcore.QEvent_ITF
	QExposeEvent_PTR() *QExposeEvent
}

func (ptr *QExposeEvent) QExposeEvent_PTR() *QExposeEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QExposeEventFromptr(cthis Voidptr) *QExposeEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QExposeEvent{bcthis0}
}
func (*QExposeEvent) Fromptr(cthis Voidptr) *QExposeEvent {
	return QExposeEventFromptr(cthis)
}

func DeleteQExposeEvent(this *QExposeEvent) {
	rv, err := qtrt.Qtcc3(1889365036, "_ZN12QExposeEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10081() {
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
