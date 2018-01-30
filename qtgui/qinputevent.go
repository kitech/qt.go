package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

func init() {
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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QInputEvent struct {
	*qtcore.QEvent
}

func (this *QInputEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QInputEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQInputEventFromPointer(cthis unsafe.Pointer) *QInputEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QInputEvent{bcthis0}
}
func (*QInputEvent) NewFromPointer(cthis unsafe.Pointer) *QInputEvent {
	return NewQInputEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QInputEvent()
func DeleteQInputEvent(*QInputEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QInputEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers()
func (this *QInputEvent) Modifiers() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QInputEvent9modifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [8] ulong timestamp()
func (this *QInputEvent) Timestamp() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QInputEvent9timestampEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qevent.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTimestamp(ulong)
func (this *QInputEvent) SetTimestamp(atimestamp uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QInputEvent12setTimestampEm", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), atimestamp)
	gopp.ErrPrint(err, rv)
}

//  body block end
