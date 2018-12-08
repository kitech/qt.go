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
// extern C begin: 8
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
type QStatusTipEvent struct {
	*qtcore.QEvent
}
type QStatusTipEvent_ITF interface {
	qtcore.QEvent_ITF
	QStatusTipEvent_PTR() *QStatusTipEvent
}

func (ptr *QStatusTipEvent) QStatusTipEvent_PTR() *QStatusTipEvent { return ptr }

func (this *QStatusTipEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QStatusTipEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQStatusTipEventFromPointer(cthis unsafe.Pointer) *QStatusTipEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QStatusTipEvent{bcthis0}
}
func (*QStatusTipEvent) NewFromPointer(cthis unsafe.Pointer) *QStatusTipEvent {
	return NewQStatusTipEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:700
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStatusTipEvent(const QString &)

/*

 */
func (*QStatusTipEvent) NewForInherit(tip string) *QStatusTipEvent {
	return NewQStatusTipEvent(tip)
}
func NewQStatusTipEvent(tip string) *QStatusTipEvent {
	var tmpArg0 = qtcore.NewQString5(tip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStatusTipEventC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStatusTipEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStatusTipEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:701
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStatusTipEvent()

/*

 */
func DeleteQStatusTipEvent(this *QStatusTipEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStatusTipEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:703
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString tip() const

/*

 */
func (this *QStatusTipEvent) Tip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QStatusTipEvent3tipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

//  body block end

//  keep block begin

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
