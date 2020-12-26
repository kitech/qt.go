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
type QDynamicPropertyChangeEvent struct {
	*QEvent
}
type QDynamicPropertyChangeEvent_ITF interface {
	QEvent_ITF
	QDynamicPropertyChangeEvent_PTR() *QDynamicPropertyChangeEvent
}

func (ptr *QDynamicPropertyChangeEvent) QDynamicPropertyChangeEvent_PTR() *QDynamicPropertyChangeEvent {
	return ptr
}

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QDynamicPropertyChangeEventFromptr(cthis Voidptr) *QDynamicPropertyChangeEvent {
	bcthis0 := QEventFromptr(cthis)
	return &QDynamicPropertyChangeEvent{bcthis0}
}
func (*QDynamicPropertyChangeEvent) Fromptr(cthis Voidptr) *QDynamicPropertyChangeEvent {
	return QDynamicPropertyChangeEventFromptr(cthis)
}

func DeleteQDynamicPropertyChangeEvent(this *QDynamicPropertyChangeEvent) {
	rv, err := qtrt.Qtcc1(0, "_ZN27QDynamicPropertyChangeEventD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10033() {
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
