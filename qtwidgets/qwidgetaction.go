package qtwidgets

// /usr/include/qt/QtWidgets/qwidgetaction.h
// #include <qwidgetaction.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
func (this *QWidgetAction) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QWidgetAction) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// QWidget * createWidget(QWidget *)
func (this *QWidgetAction) InheritCreateWidget(f func(parent *QWidget /*777 QWidget **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createWidget", f)
}

// void deleteWidget(QWidget *)
func (this *QWidgetAction) InheritDeleteWidget(f func(widget *QWidget /*777 QWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "deleteWidget", f)
}

/*

 */
type QWidgetAction struct {
	*QAction
}
type QWidgetAction_ITF interface {
	QAction_ITF
	QWidgetAction_PTR() *QWidgetAction
}

func (ptr *QWidgetAction) QWidgetAction_PTR() *QWidgetAction { return ptr }

func (this *QWidgetAction) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAction.GetCthis()
	}
}
func (this *QWidgetAction) SetCthis(cthis unsafe.Pointer) {
	this.QAction = NewQActionFromPointer(cthis)
}
func NewQWidgetActionFromPointer(cthis unsafe.Pointer) *QWidgetAction {
	bcthis0 := NewQActionFromPointer(cthis)
	return &QWidgetAction{bcthis0}
}
func (*QWidgetAction) NewFromPointer(cthis unsafe.Pointer) *QWidgetAction {
	return NewQWidgetActionFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWidgetAction) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QWidgetAction10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidgetAction(QObject *)

/*
Constructs an action with parent.
*/
func NewQWidgetAction(parent qtcore.QObject_ITF /*777 QObject **/) *QWidgetAction {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetActionC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWidgetActionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWidgetAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWidgetAction()

/*

 */
func DeleteQWidgetAction(this *QWidgetAction) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetActionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultWidget(QWidget *)

/*
Sets widget to be the default widget. The ownership is transferred to QWidgetAction. Unless createWidget() is reimplemented by a subclass to return a new widget the default widget is used when a container widget requests a widget through requestWidget().

See also defaultWidget().
*/
func (this *QWidgetAction) SetDefaultWidget(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction16setDefaultWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * defaultWidget() const

/*
Returns the default widget.

See also setDefaultWidget().
*/
func (this *QWidgetAction) DefaultWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QWidgetAction13defaultWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * requestWidget(QWidget *)

/*
Returns a widget that represents the action, with the given parent.

Container widgets that support actions can call this function to request a widget as visual representation of the action.

See also releaseWidget(), createWidget(), and defaultWidget().
*/
func (this *QWidgetAction) RequestWidget(parent QWidget_ITF /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction13requestWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseWidget(QWidget *)

/*
Releases the specified widget.

Container widgets that support actions call this function when a widget action is removed.

See also requestWidget(), deleteWidget(), and defaultWidget().
*/
func (this *QWidgetAction) ReleaseWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction13releaseWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QWidgetAction) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().
*/
func (this *QWidgetAction) EventFilter(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:71
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QWidget * createWidget(QWidget *)

/*
This function is called whenever the action is added to a container widget that supports custom widgets. If you don't want a custom widget to be used as representation of the action in the specified parent widget then 0 should be returned.

See also deleteWidget().
*/
func (this *QWidgetAction) CreateWidget(parent QWidget_ITF /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction12createWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:72
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void deleteWidget(QWidget *)

/*
This function is called whenever the action is removed from a container widget that displays the action using a custom widget previously created using createWidget(). The default implementation hides the widget and schedules it for deletion using QObject::deleteLater().

See also createWidget().
*/
func (this *QWidgetAction) DeleteWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QWidgetAction12deleteWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
