//  header block begin
// /usr/include/qt/QtGui/qaccessibleobject.h
// #include <qaccessibleobject.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QAccessibleApplication struct {
	*qtrt.CObject
}

func (this *QAccessibleApplication) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qaccessibleobject.h:78
// index:0
// void QAccessibleApplication()
func NewQAccessibleApplication() *QAccessibleApplication {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QAccessibleApplicationC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleApplicationFromPointer(cthis)
	return gothis
}
func NewQAccessibleApplicationFromPointer(cthis unsafe.Pointer) *QAccessibleApplication {
	return &QAccessibleApplication{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qaccessibleobject.h:80
// index:0
// virtual
// QWindow * window()
func (this *QAccessibleApplication) Window() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication6windowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:82
// index:0
// virtual
// int childCount()
func (this *QAccessibleApplication) ChildCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication10childCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:83
// index:0
// virtual
// int indexOfChild(const class QAccessibleInterface *)
func (this *QAccessibleApplication) IndexOfChild(arg0 unsafe.Pointer) {
	// 0: (, const QAccessibleInterface * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:84
// index:0
// virtual
// QAccessibleInterface * focusChild()
func (this *QAccessibleApplication) FocusChild() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication10focusChildEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:87
// index:0
// virtual
// QAccessibleInterface * parent()
func (this *QAccessibleApplication) Parent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication6parentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:88
// index:0
// virtual
// QAccessibleInterface * child(int)
func (this *QAccessibleApplication) Child(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication5childEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:91
// index:0
// virtual
// QString text(class QAccessible::Text)
func (this *QAccessibleApplication) Text(t int) {
	// 0: (, t QAccessible::Text), (&t)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication4textEN11QAccessible4TextE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:92
// index:0
// virtual
// QAccessible::Role role()
func (this *QAccessibleApplication) Role() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication4roleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:93
// index:0
// virtual
// QAccessible::State state()
func (this *QAccessibleApplication) State() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication5stateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
