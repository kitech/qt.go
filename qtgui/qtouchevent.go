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
// extern C begin: 4
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
type QTouchEvent struct {
	*QInputEvent
}
type QTouchEvent_ITF interface {
	QInputEvent_ITF
	QTouchEvent_PTR() *QTouchEvent
}

func (ptr *QTouchEvent) QTouchEvent_PTR() *QTouchEvent { return ptr }

func (this *QTouchEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QTouchEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQTouchEventFromPointer(cthis unsafe.Pointer) *QTouchEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QTouchEvent{bcthis0}
}
func (*QTouchEvent) NewFromPointer(cthis unsafe.Pointer) *QTouchEvent {
	return NewQTouchEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:940
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTouchEvent()

/*

 */
func DeleteQTouchEvent(this *QTouchEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTouchEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 72)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:942
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QWindow * window() const

/*

 */
func (this *QTouchEvent) Window() *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTouchEvent6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qevent.h:943
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QObject * target() const

/*

 */
func (this *QTouchEvent) Target() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTouchEvent6targetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qevent.h:947
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::TouchPointStates touchPointStates() const

/*

 */
func (this *QTouchEvent) TouchPointStates() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTouchEvent16touchPointStatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:949
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QTouchDevice * device() const

/*

 */
func (this *QTouchEvent) Device() *QTouchDevice /*777 QTouchDevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTouchEvent6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTouchDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qevent.h:952
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWindow(QWindow *)

/*

 */
func (this *QTouchEvent) SetWindow(awindow QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if awindow != nil && awindow.QWindow_PTR() != nil {
		convArg0 = awindow.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTouchEvent9setWindowEP7QWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:953
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTarget(QObject *)

/*

 */
func (this *QTouchEvent) SetTarget(atarget qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if atarget != nil && atarget.QObject_PTR() != nil {
		convArg0 = atarget.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTouchEvent9setTargetEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:954
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTouchPointStates(Qt::TouchPointStates)

/*

 */
func (this *QTouchEvent) SetTouchPointStates(aTouchPointStates int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTouchEvent19setTouchPointStatesE6QFlagsIN2Qt15TouchPointStateEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), aTouchPointStates)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:956
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setDevice(QTouchDevice *)

/*

 */
func (this *QTouchEvent) SetDevice(adevice QTouchDevice_ITF /*777 QTouchDevice **/) {
	var convArg0 unsafe.Pointer
	if adevice != nil && adevice.QTouchDevice_PTR() != nil {
		convArg0 = adevice.QTouchDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTouchEvent9setDeviceEP12QTouchDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
