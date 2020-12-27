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
type QCloseEvent struct {
	*qtcore.QEvent
}
type QCloseEvent_ITF interface {
	qtcore.QEvent_ITF
	QCloseEvent_PTR() *QCloseEvent
}

func (ptr *QCloseEvent) QCloseEvent_PTR() *QCloseEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QCloseEventFromptr(cthis Voidptr) *QCloseEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QCloseEvent{bcthis0}
}
func (*QCloseEvent) Fromptr(cthis Voidptr) *QCloseEvent {
	return QCloseEventFromptr(cthis)
}

func DeleteQCloseEvent(this *QCloseEvent) {
	rv, err := qtrt.Qtcc3(1300068505, "_ZN11QCloseEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10095() {
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
