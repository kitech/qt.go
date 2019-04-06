//  header block begin

// +build !minimal

package qtquick

// /usr/include/qt/QtQuick/qquickitem.h
// #include <qquickitem.h>
// #include <QtQuick>

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
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

// /usr/include/qt/QtQuick/qquickitem.h:302
// index:0
// Public Visibility=Default Availability=Available
// [8] QCursor cursor() const

/*
Returns the cursor shape for this item.

The mouse cursor will assume this shape when it is over this item, unless an override cursor is set. See the list of predefined cursor objects for a range of useful shapes.

If no cursor shape has been set this returns a cursor with the Qt::ArrowCursor shape, however another cursor shape may be displayed if an overlapping item has a valid cursor.

See also setCursor() and unsetCursor().
*/
func (this *QQuickItem) Cursor() *qtgui.QCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem6cursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQCursor)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:303
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursor(const QCursor &)

/*
Sets the cursor shape for this item.

See also cursor() and unsetCursor().
*/
func (this *QQuickItem) SetCursor(cursor qtgui.QCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QCursor_PTR() != nil {
		convArg0 = cursor.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9setCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:304
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetCursor()

/*
Clears the cursor shape for this item.

See also cursor() and setCursor().
*/
func (this *QQuickItem) UnsetCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11unsetCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:352
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
This method is only relevant for input items.

If this item is an input item, this method should be reimplemented to return the relevant input method flags for the given query.

See also QWidget::inputMethodQuery().
*/
func (this *QQuickItem) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:406
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateInputMethod(Qt::InputMethodQueries)

/*
Notify input method on updated query values if needed. queries indicates the changed attributes.
*/
func (this *QQuickItem) UpdateInputMethod(queries int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17updateInputMethodE6QFlagsIN2Qt16InputMethodQueryEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), queries)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:406
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateInputMethod(Qt::InputMethodQueries)

/*
Notify input method on updated query values if needed. queries indicates the changed attributes.
*/
func (this *QQuickItem) UpdateInputMethodp() {
	// arg: 0, Qt::InputMethodQueries=Elaborated, Qt::InputMethodQueries=Typedef, QFlags<Qt::InputMethodQuery>, Unexposed
	queries := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17updateInputMethodE6QFlagsIN2Qt16InputMethodQueryEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), queries)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:419
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
This event handler can be reimplemented in a subclass to receive input method events for an item. The event information is provided by the event parameter.

The event is accepted by default, so it is not necessary to explicitly accept the event if you reimplement this function. If you don't accept the event, call event->ignore().
*/
func (this *QQuickItem) InputMethodEvent(arg0 qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QInputMethodEvent_PTR() != nil {
		convArg0 = arg0.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:430
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
This event handler can be reimplemented in a subclass to receive wheel events for an item. The event information is provided by the event parameter.

The event is accepted by default, so it is not necessary to explicitly accept the event if you reimplement this function. If you don't accept the event, call event->ignore().
*/
func (this *QQuickItem) WheelEvent(event qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QWheelEvent_PTR() != nil {
		convArg0 = event.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:437
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
This event handler can be reimplemented in a subclass to receive drag-enter events for an item. The event information is provided by the event parameter.

Drag and drop events are only provided if the ItemAcceptsDrops flag has been set for this item.

The event is accepted by default, so it is not necessary to explicitly accept the event if you reimplement this function. If you don't accept the event, call event->ignore().

See also Drag and Drag and Drop.
*/
func (this *QQuickItem) DragEnterEvent(arg0 qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragEnterEvent_PTR() != nil {
		convArg0 = arg0.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:438
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
This event handler can be reimplemented in a subclass to receive drag-move events for an item. The event information is provided by the event parameter.

Drag and drop events are only provided if the ItemAcceptsDrops flag has been set for this item.

The event is accepted by default, so it is not necessary to explicitly accept the event if you reimplement this function. If you don't accept the event, call event->ignore().

See also Drag and Drag and Drop.
*/
func (this *QQuickItem) DragMoveEvent(arg0 qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragMoveEvent_PTR() != nil {
		convArg0 = arg0.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:439
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
This event handler can be reimplemented in a subclass to receive drag-leave events for an item. The event information is provided by the event parameter.

Drag and drop events are only provided if the ItemAcceptsDrops flag has been set for this item.

The event is accepted by default, so it is not necessary to explicitly accept the event if you reimplement this function. If you don't accept the event, call event->ignore().

See also Drag and Drag and Drop.
*/
func (this *QQuickItem) DragLeaveEvent(arg0 qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragLeaveEvent_PTR() != nil {
		convArg0 = arg0.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:440
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
This event handler can be reimplemented in a subclass to receive drop events for an item. The event information is provided by the event parameter.

Drag and drop events are only provided if the ItemAcceptsDrops flag has been set for this item.

The event is accepted by default, so it is not necessary to explicitly accept the event if you reimplement this function. If you don't accept the event, call event->ignore().

See also Drag and Drag and Drop.
*/
func (this *QQuickItem) DropEvent(arg0 qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDropEvent_PTR() != nil {
		convArg0 = arg0.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11548() {
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
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
