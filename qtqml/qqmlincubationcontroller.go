package qtqml

// /usr/include/qt/QtQml/qqmlincubator.h
// #include <qqmlincubator.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQmlIncubationController struct {
	*qtrt.CObject
}

func (this *QQmlIncubationController) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlIncubationController) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlIncubationControllerFromPointer(cthis unsafe.Pointer) *QQmlIncubationController {
	return &QQmlIncubationController{&qtrt.CObject{cthis}}
}
func (*QQmlIncubationController) NewFromPointer(cthis unsafe.Pointer) *QQmlIncubationController {
	return NewQQmlIncubationControllerFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlincubator.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlIncubationController()
func NewQQmlIncubationController() *QQmlIncubationController {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QQmlIncubationControllerC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlIncubationControllerFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlincubator.h:104
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlIncubationController()
func DeleteQQmlIncubationController(*QQmlIncubationController) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QQmlIncubationControllerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlEngine * engine()
func (this *QQmlIncubationController) Engine() *QQmlEngine /*777 QQmlEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QQmlIncubationController6engineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlincubator.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int incubatingObjectCount()
func (this *QQmlIncubationController) IncubatingObjectCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QQmlIncubationController21incubatingObjectCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlincubator.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void incubateFor(int)
func (this *QQmlIncubationController) IncubateFor(msecs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QQmlIncubationController11incubateForEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void incubateWhile(volatile _Bool *, int)
func (this *QQmlIncubationController) IncubateWhile(flag unsafe.Pointer /*666*/, msecs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QQmlIncubationController13incubateWhileEPVbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &flag, msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void incubatingObjectCountChanged(int)
func (this *QQmlIncubationController) IncubatingObjectCountChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QQmlIncubationController28incubatingObjectCountChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
