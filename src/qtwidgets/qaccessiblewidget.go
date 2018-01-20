//  header block begin
// /usr/include/qt/QtWidgets/qaccessiblewidget.h
// #include <qaccessiblewidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 134
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QAccessibleWidget struct {
	*qtrt.CObject
}

func (this *QAccessibleWidget) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQAccessibleWidgetFromPointer(cthis unsafe.Pointer) *QAccessibleWidget {
	return &QAccessibleWidget{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:56
// index:0
// Public
// void QAccessibleWidget(class QWidget *, class QAccessible::Role, const class QString &)
func NewQAccessibleWidget(o unsafe.Pointer, r int, name *qtcore.QString) *QAccessibleWidget {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg2 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidgetC2EP7QWidgetN11QAccessible4RoleERK7QString", ffiqt.FFI_TYPE_VOID, cthis, o, &r, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:57
// index:0
// Public virtual
// bool isValid()
func (this *QAccessibleWidget) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:59
// index:0
// Public virtual
// QWindow * window()
func (this *QAccessibleWidget) Window() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:60
// index:0
// Public virtual
// int childCount()
func (this *QAccessibleWidget) ChildCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget10childCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:61
// index:0
// Public virtual
// int indexOfChild(const class QAccessibleInterface *)
func (this *QAccessibleWidget) IndexOfChild(child unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), child)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:63
// index:0
// Public virtual
// QAccessibleInterface * focusChild()
func (this *QAccessibleWidget) FocusChild() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget10focusChildEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:65
// index:0
// Public virtual
// QRect rect()
func (this *QAccessibleWidget) Rect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4rectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:67
// index:0
// Public virtual
// QAccessibleInterface * parent()
func (this *QAccessibleWidget) Parent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:68
// index:0
// Public virtual
// QAccessibleInterface * child(int)
func (this *QAccessibleWidget) Child(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget5childEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:70
// index:0
// Public virtual
// QString text(class QAccessible::Text)
func (this *QAccessibleWidget) Text(t int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4textEN11QAccessible4TextE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:71
// index:0
// Public virtual
// QAccessible::Role role()
func (this *QAccessibleWidget) Role() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4roleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:72
// index:0
// Public virtual
// QAccessible::State state()
func (this *QAccessibleWidget) State() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:74
// index:0
// Public virtual
// QColor foregroundColor()
func (this *QAccessibleWidget) ForegroundColor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget15foregroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:75
// index:0
// Public virtual
// QColor backgroundColor()
func (this *QAccessibleWidget) BackgroundColor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget15backgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:77
// index:0
// Public virtual
// void * interface_cast(class QAccessible::InterfaceType)
func (this *QAccessibleWidget) Interface_cast(t int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget14interface_castEN11QAccessible13InterfaceTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:80
// index:0
// Public virtual
// QStringList actionNames()
func (this *QAccessibleWidget) ActionNames() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget11actionNamesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:81
// index:0
// Public virtual
// void doAction(const class QString &)
func (this *QAccessibleWidget) DoAction(actionName *qtcore.QString) {
	var convArg0 = actionName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget8doActionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:82
// index:0
// Public virtual
// QStringList keyBindingsForAction(const class QString &)
func (this *QAccessibleWidget) KeyBindingsForAction(actionName *qtcore.QString) interface{} {
	var convArg0 = actionName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget20keyBindingsForActionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:84
// index:0
// Protected virtual
// void ~QAccessibleWidget()
func DeleteQAccessibleWidget(*QAccessibleWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:85
// index:0
// Protected
// QWidget * widget()
func (this *QAccessibleWidget) Widget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:86
// index:0
// Protected
// QObject * parentObject()
func (this *QAccessibleWidget) ParentObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget12parentObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:88
// index:0
// Protected
// void addControllingSignal(const class QString &)
func (this *QAccessibleWidget) AddControllingSignal(signal *qtcore.QString) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget20addControllingSignalERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
