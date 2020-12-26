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
type QInputEvent struct {
	*qtcore.QEvent
}
type QInputEvent_ITF interface {
	qtcore.QEvent_ITF
	QInputEvent_PTR() *QInputEvent
}

func (ptr *QInputEvent) QInputEvent_PTR() *QInputEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QInputEventFromptr(cthis Voidptr) *QInputEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QInputEvent{bcthis0}
}
func (*QInputEvent) Fromptr(cthis Voidptr) *QInputEvent {
	return QInputEventFromptr(cthis)
}

// /usr/include/qt/QtGui/qevent.h:75
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] ulong timestamp() const

/*
 */
func (this *QInputEvent) Timestamp() uint {
	rv, err := qtrt.Qtcc3(3682385401, "_ZNK11QInputEvent9timestampEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qevent.h:76
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setTimestamp(ulong)

/*
 */
func (this *QInputEvent) SetTimestamp(atimestamp uint) {
	rv, err := qtrt.Qtcc3(1066347960, "_ZN11QInputEvent12setTimestampEm", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&atimestamp))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQInputEvent(this *QInputEvent) {
	rv, err := qtrt.Qtcc1(0, "_ZN11QInputEventD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10059() {
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
