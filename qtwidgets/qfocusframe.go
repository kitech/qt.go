package qtwidgets

// /usr/include/qt/QtWidgets/qfocusframe.h
// #include <qfocusframe.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 53
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

// bool event(QEvent *)
func (this *QFocusFrame) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QFocusFrame) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// void paintEvent(QPaintEvent *)
func (this *QFocusFrame) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void initStyleOption(QStyleOption *)
func (this *QFocusFrame) InheritInitStyleOption(f func(option *QStyleOption /*777 QStyleOption **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QFocusFrame struct {
	*QWidget
}
type QFocusFrame_ITF interface {
	QWidget_ITF
	QFocusFrame_PTR() *QFocusFrame
}

func (ptr *QFocusFrame) QFocusFrame_PTR() *QFocusFrame { return ptr }

func (this *QFocusFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QFocusFrame) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQFocusFrameFromPointer(cthis unsafe.Pointer) *QFocusFrame {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QFocusFrame{bcthis0}
}
func (*QFocusFrame) NewFromPointer(cthis unsafe.Pointer) *QFocusFrame {
	return NewQFocusFrameFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QFocusFrame) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFocusFrame10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfocusframe.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFocusFrame(QWidget *)

/*
Constructs a QFocusFrame.

The focus frame will not monitor parent for updates but rather can be placed manually or by using QFocusFrame::setWidget. A QFocusFrame sets Qt::WA_NoChildEventsForParent attribute; as a result the parent will not receive a QEvent::ChildAdded event, this will make it possible to manually set the geometry of the QFocusFrame inside of a QSplitter or other child event monitoring widget.

See also QFocusFrame::setWidget().
*/
func NewQFocusFrame(parent QWidget_ITF /*777 QWidget **/) *QFocusFrame {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrameC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFocusFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFocusFrame")
	return gothis
}

// /usr/include/qt/QtWidgets/qfocusframe.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFocusFrame(QWidget *)

/*
Constructs a QFocusFrame.

The focus frame will not monitor parent for updates but rather can be placed manually or by using QFocusFrame::setWidget. A QFocusFrame sets Qt::WA_NoChildEventsForParent attribute; as a result the parent will not receive a QEvent::ChildAdded event, this will make it possible to manually set the geometry of the QFocusFrame inside of a QSplitter or other child event monitoring widget.

See also QFocusFrame::setWidget().
*/
func NewQFocusFrame__() *QFocusFrame {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrameC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFocusFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFocusFrame")
	return gothis
}

// /usr/include/qt/QtWidgets/qfocusframe.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFocusFrame()

/*

 */
func DeleteQFocusFrame(this *QFocusFrame) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrameD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)

/*
QFocusFrame will track changes to widget and resize itself automatically. If the monitored widget's parent changes, QFocusFrame will follow the widget and place itself around the widget automatically. If the monitored widget is deleted, QFocusFrame will set it to zero.

See also QFocusFrame::widget().
*/
func (this *QFocusFrame) SetWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrame9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget() const

/*
Returns the currently monitored widget for automatically resize and update.

See also QFocusFrame::setWidget().
*/
func (this *QFocusFrame) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFocusFrame6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfocusframe.h:63
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QFocusFrame) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrame5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfocusframe.h:65
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().
*/
func (this *QFocusFrame) EventFilter(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrame11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfocusframe.h:66
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QFocusFrame) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFocusFrame10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfocusframe.h:67
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOption *) const

/*
Initialize option with the values from this QFocusFrame. This method is useful for subclasses when they need a QStyleOption, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QFocusFrame) InitStyleOption(option QStyleOption_ITF /*777 QStyleOption **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOption_PTR() != nil {
		convArg0 = option.QStyleOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFocusFrame15initStyleOptionEP12QStyleOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
