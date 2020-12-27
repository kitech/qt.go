package qtcore

// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>

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

//  ext block end

//  body block begin

/*
 */
// size 24
type QDeferredDeleteEvent struct {
	*QEvent
}
type QDeferredDeleteEvent_ITF interface {
	QEvent_ITF
	QDeferredDeleteEvent_PTR() *QDeferredDeleteEvent
}

func (ptr *QDeferredDeleteEvent) QDeferredDeleteEvent_PTR() *QDeferredDeleteEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QDeferredDeleteEventFromptr(cthis Voidptr) *QDeferredDeleteEvent {
	bcthis0 := QEventFromptr(cthis)
	return &QDeferredDeleteEvent{bcthis0}
}
func (*QDeferredDeleteEvent) Fromptr(cthis Voidptr) *QDeferredDeleteEvent {
	return QDeferredDeleteEventFromptr(cthis)
}

func DeleteQDeferredDeleteEvent(this *QDeferredDeleteEvent) {
	rv, err := qtrt.Qtcc3(1062849289, "_ZN20QDeferredDeleteEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10035() {
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
}

//  keep block end
