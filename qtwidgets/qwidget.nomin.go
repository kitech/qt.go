//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qwidget.h
// #include <qwidget.h>
// #include <QtWidgets>

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
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// /usr/include/qt/QtWidgets/qwidget.h:351
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsEffect * graphicsEffect() const

/*
The graphicsEffect function returns a pointer to the widget's graphics effect.

If the widget has no graphics effect, 0 is returned.

This function was introduced in  Qt 4.6.

See also setGraphicsEffect().
*/
func (this *QWidget) GraphicsEffect() *QGraphicsEffect /*777 QGraphicsEffect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14graphicsEffectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:352
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGraphicsEffect(QGraphicsEffect *)

/*
The setGraphicsEffect function is for setting the widget's graphics effect.

Sets effect as the widget's effect. If there already is an effect installed on this widget, QWidget will delete the existing effect before installing the new effect.

If effect is the installed effect on a different widget, setGraphicsEffect() will remove the effect from the widget and install it on this widget.

QWidget takes ownership of effect.

Note: This function will apply the effect on itself and all its children.

Note: Graphics effects are not supported for OpenGL-based widgets, such as QGLWidget, QOpenGLWidget and QQuickWidget.

This function was introduced in  Qt 4.6.

See also graphicsEffect().
*/
func (this *QWidget) SetGraphicsEffect(effect QGraphicsEffect_ITF /*777 QGraphicsEffect **/) {
	var convArg0 unsafe.Pointer
	if effect != nil && effect.QGraphicsEffect_PTR() != nil {
		convArg0 = effect.QGraphicsEffect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:390
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)

/*

 */
func (this *QWidget) SetStatusTip(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setStatusTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:391
// index:0
// Public Visibility=Default Availability=Available
// [8] QString statusTip() const

/*

 */
func (this *QWidget) StatusTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget9statusTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:394
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)

/*

 */
func (this *QWidget) SetWhatsThis(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setWhatsThisERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:395
// index:0
// Public Visibility=Default Availability=Available
// [8] QString whatsThis() const

/*

 */
func (this *QWidget) WhatsThis() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget9whatsThisEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:455
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsProxyWidget * graphicsProxyWidget() const

/*
Returns the proxy widget for the corresponding embedded widget in a graphics view; otherwise returns 0.

This function was introduced in  Qt 4.5.

See also QGraphicsProxyWidget::createProxyForChildWidget() and QGraphicsScene::addWidget().
*/
func (this *QWidget) GraphicsProxyWidget() *QGraphicsProxyWidget /*777 QGraphicsProxyWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget19graphicsProxyWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:617
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive wheel events for the widget.

If you reimplement this handler, it is very important that you ignore() the event if you do not handle it, so that the widget's parent can interpret it.

The default implementation ignores the event.

See also QEvent::ignore(), QEvent::accept(), event(), and QWheelEvent.
*/
func (this *QWidget) WheelEvent(event qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QWheelEvent_PTR() != nil {
		convArg0 = event.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:633
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabletEvent(QTabletEvent *)

/*
This event handler, for event event, can be reimplemented in a subclass to receive tablet events for the widget.

If you reimplement this handler, it is very important that you ignore() the event if you do not handle it, so that the widget's parent can interpret it.

The default implementation ignores the event.

If tablet tracking is switched off, tablet move events only occur if the stylus is in contact with the tablet, or at least one stylus button is pressed, while the stylus is being moved. If tablet tracking is switched on, tablet move events occur even while the stylus is hovering in proximity of the tablet, with no buttons pressed.

See also QEvent::ignore(), QEvent::accept(), event(), setTabletTracking(), and QTabletEvent.
*/
func (this *QWidget) TabletEvent(event qtgui.QTabletEvent_ITF /*777 QTabletEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QTabletEvent_PTR() != nil {
		convArg0 = event.QTabletEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11tabletEventEP12QTabletEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:640
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
This event handler is called when a drag is in progress and the mouse enters this widget. The event is passed in the event parameter.

If the event is ignored, the widget won't receive any drag move events.

See the Drag-and-drop documentation for an overview of how to provide drag-and-drop in your application.

See also QDrag and QDragEnterEvent.
*/
func (this *QWidget) DragEnterEvent(event qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragEnterEvent_PTR() != nil {
		convArg0 = event.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:641
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
This event handler is called if a drag is in progress, and when any of the following conditions occur: the cursor enters this widget, the cursor moves within this widget, or a modifier key is pressed on the keyboard while this widget has the focus. The event is passed in the event parameter.

See the Drag-and-drop documentation for an overview of how to provide drag-and-drop in your application.

See also QDrag and QDragMoveEvent.
*/
func (this *QWidget) DragMoveEvent(event qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragMoveEvent_PTR() != nil {
		convArg0 = event.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:642
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
This event handler is called when a drag is in progress and the mouse leaves this widget. The event is passed in the event parameter.

See the Drag-and-drop documentation for an overview of how to provide drag-and-drop in your application.

See also QDrag and QDragLeaveEvent.
*/
func (this *QWidget) DragLeaveEvent(event qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragLeaveEvent_PTR() != nil {
		convArg0 = event.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:643
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
This event handler is called when the drag is dropped on this widget. The event is passed in the event parameter.

See the Drag-and-drop documentation for an overview of how to provide drag-and-drop in your application.

See also QDrag and QDropEvent.
*/
func (this *QWidget) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDropEvent_PTR() != nil {
		convArg0 = event.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10984() {
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
