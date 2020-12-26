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
// extern C begin: 7
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
type QTimerEvent struct {
	*QEvent
}
type QTimerEvent_ITF interface {
	QEvent_ITF
	QTimerEvent_PTR() *QTimerEvent
}

func (ptr *QTimerEvent) QTimerEvent_PTR() *QTimerEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QTimerEventFromptr(cthis Voidptr) *QTimerEvent {
	bcthis0 := QEventFromptr(cthis)
	return &QTimerEvent{bcthis0}
}
func (*QTimerEvent) Fromptr(cthis Voidptr) *QTimerEvent {
	return QTimerEventFromptr(cthis)
}

func DeleteQTimerEvent(this *QTimerEvent) {
	rv, err := qtrt.Qtcc1(0, "_ZN11QTimerEventD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10029() {
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
