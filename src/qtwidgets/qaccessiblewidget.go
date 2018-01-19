//  header block begin
// /usr/include/qt/QtWidgets/qaccessiblewidget.h
// #include <qaccessiblewidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:56
// index:0
// void QAccessibleWidget(class QWidget *, class QAccessible::Role, const class QString &)
func NewQAccessibleWidget(o unsafe.Pointer, r int, name unsafe.Pointer) *QAccessibleWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidgetC2EP7QWidgetN11QAccessible4RoleERK7QString", ffiqt.FFI_TYPE_VOID, cthis, o, &r, name)
	gopp.ErrPrint(err, rv)
	return &QAccessibleWidget{cthis}
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:57
// index:0
// virtual
// bool isValid()
func (this *QAccessibleWidget) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:59
// index:0
// virtual
// QWindow * window()
func (this *QAccessibleWidget) Window() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6windowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:60
// index:0
// virtual
// int childCount()
func (this *QAccessibleWidget) ChildCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget10childCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:61
// index:0
// virtual
// int indexOfChild(const class QAccessibleInterface *)
func (this *QAccessibleWidget) IndexOfChild(child unsafe.Pointer) {
	// 0: (, const QAccessibleInterface * child), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget12indexOfChildEPK20QAccessibleInterface", ffiqt.FFI_TYPE_VOID, this.cthis, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:63
// index:0
// virtual
// QAccessibleInterface * focusChild()
func (this *QAccessibleWidget) FocusChild() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget10focusChildEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:65
// index:0
// virtual
// QRect rect()
func (this *QAccessibleWidget) Rect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4rectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:67
// index:0
// virtual
// QAccessibleInterface * parent()
func (this *QAccessibleWidget) Parent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget6parentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:68
// index:0
// virtual
// QAccessibleInterface * child(int)
func (this *QAccessibleWidget) Child(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget5childEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:70
// index:0
// virtual
// QString text(class QAccessible::Text)
func (this *QAccessibleWidget) Text(t int) {
	// 0: (, QAccessible::Text t), (&t)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4textEN11QAccessible4TextE", ffiqt.FFI_TYPE_VOID, this.cthis, &t)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:71
// index:0
// virtual
// QAccessible::Role role()
func (this *QAccessibleWidget) Role() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget4roleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:72
// index:0
// virtual
// QAccessible::State state()
func (this *QAccessibleWidget) State() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget5stateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:74
// index:0
// virtual
// QColor foregroundColor()
func (this *QAccessibleWidget) ForegroundColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget15foregroundColorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:75
// index:0
// virtual
// QColor backgroundColor()
func (this *QAccessibleWidget) BackgroundColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget15backgroundColorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:77
// index:0
// virtual
// void * interface_cast(class QAccessible::InterfaceType)
func (this *QAccessibleWidget) Interface_cast(t int) {
	// 0: (, QAccessible::InterfaceType t), (&t)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget14interface_castEN11QAccessible13InterfaceTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &t)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:80
// index:0
// virtual
// QStringList actionNames()
func (this *QAccessibleWidget) ActionNames() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget11actionNamesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:81
// index:0
// virtual
// void doAction(const class QString &)
func (this *QAccessibleWidget) DoAction(actionName unsafe.Pointer) {
	// 0: (, const QString & actionName), (actionName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleWidget8doActionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, actionName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qaccessiblewidget.h:82
// index:0
// virtual
// QStringList keyBindingsForAction(const class QString &)
func (this *QAccessibleWidget) KeyBindingsForAction(actionName unsafe.Pointer) {
	// 0: (, const QString & actionName), (actionName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessibleWidget20keyBindingsForActionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, actionName)
	gopp.ErrPrint(err, rv)
}

//  body block end
