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
// size 32
type QChildEvent struct {
	*QEvent
}
type QChildEvent_ITF interface {
	QEvent_ITF
	QChildEvent_PTR() *QChildEvent
}

func (ptr *QChildEvent) QChildEvent_PTR() *QChildEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QChildEventFromptr(cthis Voidptr) *QChildEvent {
	bcthis0 := QEventFromptr(cthis)
	return &QChildEvent{bcthis0}
}
func (*QChildEvent) Fromptr(cthis Voidptr) *QChildEvent {
	return QChildEventFromptr(cthis)
}

func DeleteQChildEvent(this *QChildEvent) {
	rv, err := qtrt.Qtcc3(2956324060, "_ZN11QChildEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10031() {
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
