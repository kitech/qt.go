package qtgui

// /usr/include/qt/QtGui/qaccessibleobject.h
// #include <qaccessibleobject.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleApplication) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQAccessibleApplicationFromPointer(cthis unsafe.Pointer) *QAccessibleApplication {
	return &QAccessibleApplication{&qtrt.CObject{cthis}}
}
func (*QAccessibleApplication) NewFromPointer(cthis unsafe.Pointer) *QAccessibleApplication {
	return NewQAccessibleApplicationFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAccessibleApplication()
func NewQAccessibleApplication() *QAccessibleApplication {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QAccessibleApplicationC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleApplicationFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qaccessibleobject.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWindow * window()
func (this *QAccessibleApplication) Window() *QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qaccessibleobject.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int childCount()
func (this *QAccessibleApplication) ChildCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication10childCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessibleobject.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int indexOfChild(const QAccessibleInterface *)
func (this *QAccessibleApplication) IndexOfChild(arg0 *QAccessibleInterface /*777 const QAccessibleInterface **/) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessibleobject.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * focusChild()
func (this *QAccessibleApplication) FocusChild() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication10focusChildEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qaccessibleobject.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * parent()
func (this *QAccessibleApplication) Parent() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qaccessibleobject.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * child(int)
func (this *QAccessibleApplication) Child(index int) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication5childEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qaccessibleobject.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString text(QAccessible::Text)
func (this *QAccessibleApplication) Text(t int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication4textEN11QAccessible4TextE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), t)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qaccessibleobject.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QAccessible::Role role()
func (this *QAccessibleApplication) Role() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication4roleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessible::State state()
func (this *QAccessibleApplication) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QAccessibleApplication5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
