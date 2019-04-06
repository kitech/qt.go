package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

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
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QGestureEvent struct {
	*qtcore.QEvent
}
type QGestureEvent_ITF interface {
	qtcore.QEvent_ITF
	QGestureEvent_PTR() *QGestureEvent
}

func (ptr *QGestureEvent) QGestureEvent_PTR() *QGestureEvent { return ptr }

func (this *QGestureEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QGestureEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQGestureEventFromPointer(cthis unsafe.Pointer) *QGestureEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QGestureEvent{bcthis0}
}
func (*QGestureEvent) NewFromPointer(cthis unsafe.Pointer) *QGestureEvent {
	return NewQGestureEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesture.h:278
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGestureEvent()

/*

 */
func DeleteQGestureEvent(this *QGestureEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 56)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesture.h:281
// index:0
// Public Visibility=Default Availability=Available
// [8] QGesture * gesture(Qt::GestureType) const

/*

 */
func (this *QGestureEvent) Gesture(type_ int) *QGesture /*777 QGesture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGestureEvent7gestureEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGestureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgesture.h:291
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccepted(QGesture *, bool)

/*

 */
func (this *QGestureEvent) SetAccepted(arg0 QGesture_ITF /*777 QGesture **/, arg1 bool) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QGesture_PTR() != nil {
		convArg0 = arg0.QGesture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent11setAcceptedEP8QGestureb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:296
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setAccepted(Qt::GestureType, bool)

/*

 */
func (this *QGestureEvent) SetAccepted1(arg0 int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent11setAcceptedEN2Qt11GestureTypeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:292
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accept(QGesture *)

/*

 */
func (this *QGestureEvent) Accept(arg0 QGesture_ITF /*777 QGesture **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QGesture_PTR() != nil {
		convArg0 = arg0.QGesture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent6acceptEP8QGesture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:297
// index:1
// Public Visibility=Default Availability=Available
// [-2] void accept(Qt::GestureType)

/*

 */
func (this *QGestureEvent) Accept1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent6acceptEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ignore(QGesture *)

/*

 */
func (this *QGestureEvent) Ignore(arg0 QGesture_ITF /*777 QGesture **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QGesture_PTR() != nil {
		convArg0 = arg0.QGesture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent6ignoreEP8QGesture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:298
// index:1
// Public Visibility=Default Availability=Available
// [-2] void ignore(Qt::GestureType)

/*

 */
func (this *QGestureEvent) Ignore1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent6ignoreEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:294
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAccepted(QGesture *) const

/*

 */
func (this *QGestureEvent) IsAccepted(arg0 QGesture_ITF /*777 QGesture **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QGesture_PTR() != nil {
		convArg0 = arg0.QGesture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGestureEvent10isAcceptedEP8QGesture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgesture.h:299
// index:1
// Public Visibility=Default Availability=Available
// [1] bool isAccepted(Qt::GestureType) const

/*

 */
func (this *QGestureEvent) IsAccepted1(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGestureEvent10isAcceptedEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgesture.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)

/*

 */
func (this *QGestureEvent) SetWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGestureEvent9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:302
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget() const

/*

 */
func (this *QGestureEvent) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGestureEvent6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_11161() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
