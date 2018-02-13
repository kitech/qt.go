package qtgui

// /usr/include/qt/QtGui/qaccessibleobject.h
// #include <qaccessibleobject.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QAccessibleApplication struct {
	*qtrt.CObject
}
type QAccessibleApplication_ITF interface {
	QAccessibleApplication_PTR() *QAccessibleApplication
}

func (ptr *QAccessibleApplication) QAccessibleApplication_PTR() *QAccessibleApplication { return ptr }

func (this *QAccessibleApplication) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleApplication) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZN22QAccessibleApplicationC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleApplicationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleApplication)
	return gothis
}

// /usr/include/qt/QtGui/qaccessibleobject.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWindow * window()
func (this *QAccessibleApplication) Window() *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QAccessibleApplication6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessibleobject.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int childCount()
func (this *QAccessibleApplication) ChildCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QAccessibleApplication10childCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessibleobject.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int indexOfChild(const QAccessibleInterface *)
func (this *QAccessibleApplication) IndexOfChild(arg0 *QAccessibleInterface /*777 const QAccessibleInterface **/) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QAccessibleApplication12indexOfChildEPK20QAccessibleInterface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessibleobject.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * focusChild()
func (this *QAccessibleApplication) FocusChild() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QAccessibleApplication10focusChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessibleobject.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * parent()
func (this *QAccessibleApplication) Parent() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QAccessibleApplication6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessibleobject.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * child(int)
func (this *QAccessibleApplication) Child(index int) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QAccessibleApplication5childEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessibleobject.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString text(QAccessible::Text)
func (this *QAccessibleApplication) Text(t int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QAccessibleApplication4textEN11QAccessible4TextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessibleobject.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QAccessible::Role role()
func (this *QAccessibleApplication) Role() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QAccessibleApplication4roleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessibleobject.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessible::State state()
func (this *QAccessibleApplication) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QAccessibleApplication5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

func DeleteQAccessibleApplication(this *QAccessibleApplication) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QAccessibleApplicationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
