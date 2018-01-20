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
func NewQAccessibleApplicationFromPointer(cthis unsafe.Pointer) *QAccessibleApplication {
	return &QAccessibleApplication{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qaccessibleobject.h:78
// index:0
// Public
// void QAccessibleApplication()
func NewQAccessibleApplication() *QAccessibleApplication {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QAccessibleApplicationC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleApplicationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessibleobject.h:80
// index:0
// Public virtual
// QWindow * window()
func (this *QAccessibleApplication) Window() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessibleobject.h:82
// index:0
// Public virtual
// int childCount()
func (this *QAccessibleApplication) ChildCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication10childCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessibleobject.h:83
// index:0
// Public virtual
// int indexOfChild(const class QAccessibleInterface *)
func (this *QAccessibleApplication) IndexOfChild(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessibleobject.h:84
// index:0
// Public virtual
// QAccessibleInterface * focusChild()
func (this *QAccessibleApplication) FocusChild() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication10focusChildEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessibleobject.h:87
// index:0
// Public virtual
// QAccessibleInterface * parent()
func (this *QAccessibleApplication) Parent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessibleobject.h:88
// index:0
// Public virtual
// QAccessibleInterface * child(int)
func (this *QAccessibleApplication) Child(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication5childEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessibleobject.h:91
// index:0
// Public virtual
// QString text(class QAccessible::Text)
func (this *QAccessibleApplication) Text(t int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication4textEN11QAccessible4TextE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &t)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessibleobject.h:92
// index:0
// Public virtual
// QAccessible::Role role()
func (this *QAccessibleApplication) Role() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication4roleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessibleobject.h:93
// index:0
// Public virtual
// QAccessible::State state()
func (this *QAccessibleApplication) State() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
