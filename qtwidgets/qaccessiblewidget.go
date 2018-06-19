package qtwidgets

// /usr/include/qt/QtWidgets/qaccessiblewidget.h
// #include <qaccessiblewidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 138
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

// QWidget * widget()
func (this *QAccessibleWidget) InheritWidget(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "widget", f)
}

// QObject * parentObject()
func (this *QAccessibleWidget) InheritParentObject(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "parentObject", f)
}

// void addControllingSignal(const QString &)
func (this *QAccessibleWidget) InheritAddControllingSignal(f func(signal string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "addControllingSignal", f)
}

/*

 */
type QAccessibleWidget struct {
	*qtrt.CObject
}
type QAccessibleWidget_ITF interface {
	QAccessibleWidget_PTR() *QAccessibleWidget
}

func (ptr *QAccessibleWidget) QAccessibleWidget_PTR() *QAccessibleWidget { return ptr }

func (this *QAccessibleWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleWidget) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleWidgetFromPointer(cthis unsafe.Pointer) *QAccessibleWidget {
	return &QAccessibleWidget{&qtrt.CObject{cthis}}
}
func (*QAccessibleWidget) NewFromPointer(cthis unsafe.Pointer) *QAccessibleWidget {
	return NewQAccessibleWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAccessibleWidget(QWidget *, QAccessible::Role, const QString &)

/*
Creates a QAccessibleWidget object for widget w. role and name are optional parameters that set the object's role and name properties.
*/
func NewQAccessibleWidget(o QWidget_ITF /*777 QWidget **/, r int, name string) *QAccessibleWidget {
	var convArg0 unsafe.Pointer
	if o != nil && o.QWidget_PTR() != nil {
		convArg0 = o.QWidget_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(name)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessibleWidgetC2EP7QWidgetN11QAccessible4RoleERK7QString", qtrt.FFI_TYPE_POINTER, convArg0, r, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleWidget)
	return gothis
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAccessibleWidget(QWidget *, QAccessible::Role, const QString &)

/*
Creates a QAccessibleWidget object for widget w. role and name are optional parameters that set the object's role and name properties.
*/
func NewQAccessibleWidget__(o QWidget_ITF /*777 QWidget **/) *QAccessibleWidget {
	var convArg0 unsafe.Pointer
	if o != nil && o.QWidget_PTR() != nil {
		convArg0 = o.QWidget_PTR().GetCthis()
	}
	// arg: 1, QAccessible::Role=Elaborated, QAccessible::Role=Enum, , Invalid
	r := 0
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessibleWidgetC2EP7QWidgetN11QAccessible4RoleERK7QString", qtrt.FFI_TYPE_POINTER, convArg0, r, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleWidget)
	return gothis
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAccessibleWidget(QWidget *, QAccessible::Role, const QString &)

/*
Creates a QAccessibleWidget object for widget w. role and name are optional parameters that set the object's role and name properties.
*/
func NewQAccessibleWidget__1(o QWidget_ITF /*777 QWidget **/, r int) *QAccessibleWidget {
	var convArg0 unsafe.Pointer
	if o != nil && o.QWidget_PTR() != nil {
		convArg0 = o.QWidget_PTR().GetCthis()
	}
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessibleWidgetC2EP7QWidgetN11QAccessible4RoleERK7QString", qtrt.FFI_TYPE_POINTER, convArg0, r, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleWidget)
	return gothis
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Reimplemented from QAccessibleInterface::isValid().
*/
func (this *QAccessibleWidget) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWindow * window() const

/*
Reimplemented from QAccessibleInterface::window().
*/
func (this *QAccessibleWidget) Window() *qtgui.QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int childCount() const

/*
Reimplemented from QAccessibleInterface::childCount().
*/
func (this *QAccessibleWidget) ChildCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget10childCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int indexOfChild(const QAccessibleInterface *) const

/*
Reimplemented from QAccessibleInterface::indexOfChild().
*/
func (this *QAccessibleWidget) IndexOfChild(child qtgui.QAccessibleInterface_ITF /*777 const QAccessibleInterface **/) int {
	var convArg0 unsafe.Pointer
	if child != nil && child.QAccessibleInterface_PTR() != nil {
		convArg0 = child.QAccessibleInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * focusChild() const

/*
Reimplemented from QAccessibleInterface::focusChild().
*/
func (this *QAccessibleWidget) FocusChild() *qtgui.QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget10focusChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect rect() const

/*
Reimplemented from QAccessibleInterface::rect().
*/
func (this *QAccessibleWidget) Rect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * parent() const

/*
Reimplemented from QAccessibleInterface::parent().
*/
func (this *QAccessibleWidget) Parent() *qtgui.QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * child(int) const

/*
Reimplemented from QAccessibleInterface::child().
*/
func (this *QAccessibleWidget) Child(index int) *qtgui.QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget5childEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString text(QAccessible::Text) const

/*
Reimplemented from QAccessibleInterface::text().
*/
func (this *QAccessibleWidget) Text(t int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget4textEN11QAccessible4TextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QAccessible::Role role() const

/*
Reimplemented from QAccessibleInterface::role().
*/
func (this *QAccessibleWidget) Role() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget4roleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessible::State state() const

/*
Reimplemented from QAccessibleInterface::state().
*/
func (this *QAccessibleWidget) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QColor foregroundColor() const

/*
Reimplemented from QAccessibleInterface::foregroundColor().
*/
func (this *QAccessibleWidget) ForegroundColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget15foregroundColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QColor backgroundColor() const

/*
Reimplemented from QAccessibleInterface::backgroundColor().
*/
func (this *QAccessibleWidget) BackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget15backgroundColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] void * interface_cast(QAccessible::InterfaceType)

/*
Reimplemented from QAccessibleInterface::interface_cast().
*/
func (this *QAccessibleWidget) Interface_cast(t int) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessibleWidget14interface_castEN11QAccessible13InterfaceTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList actionNames() const

/*
Reimplemented from QAccessibleActionInterface::actionNames().
*/
func (this *QAccessibleWidget) ActionNames() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget11actionNamesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doAction(const QString &)

/*
Reimplemented from QAccessibleActionInterface::doAction().
*/
func (this *QAccessibleWidget) DoAction(actionName string) {
	var tmpArg0 = qtcore.NewQString_5(actionName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessibleWidget8doActionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList keyBindingsForAction(const QString &) const

/*
Reimplemented from QAccessibleActionInterface::keyBindingsForAction().
*/
func (this *QAccessibleWidget) KeyBindingsForAction(actionName string) *qtcore.QStringList /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(actionName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget20keyBindingsForActionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:84
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleWidget()

/*

 */
func DeleteQAccessibleWidget(this *QAccessibleWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessibleWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:85
// index:0
// Protected Visibility=Default Availability=Available
// [8] QWidget * widget() const

/*
Returns the associated widget.
*/
func (this *QAccessibleWidget) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:86
// index:0
// Protected Visibility=Default Availability=Available
// [8] QObject * parentObject() const

/*
Returns the associated widget's parent object, which is either the parent widget, or qApp for top-level widgets.
*/
func (this *QAccessibleWidget) ParentObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAccessibleWidget12parentObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:88
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addControllingSignal(const QString &)

/*
Registers signal as a controlling signal.

An object is a Controller to any other object connected to a controlling signal.
*/
func (this *QAccessibleWidget) AddControllingSignal(signal string) {
	var tmpArg0 = qtcore.NewQString_5(signal)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessibleWidget20addControllingSignalERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
