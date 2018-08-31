package qtgui

// /usr/include/qt/QtGui/qopenglversionfunctions.h
// #include <qopenglversionfunctions.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// bool isInitialized()
func (this *QAbstractOpenGLFunctions) InheritIsInitialized(f func() bool) {
	qtrt.SetAllInheritCallback(this, "isInitialized", f)
}

/*

 */
type QAbstractOpenGLFunctions struct {
	*qtrt.CObject
}
type QAbstractOpenGLFunctions_ITF interface {
	QAbstractOpenGLFunctions_PTR() *QAbstractOpenGLFunctions
}

func (ptr *QAbstractOpenGLFunctions) QAbstractOpenGLFunctions_PTR() *QAbstractOpenGLFunctions {
	return ptr
}

func (this *QAbstractOpenGLFunctions) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAbstractOpenGLFunctions) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAbstractOpenGLFunctionsFromPointer(cthis unsafe.Pointer) *QAbstractOpenGLFunctions {
	return &QAbstractOpenGLFunctions{&qtrt.CObject{cthis}}
}
func (*QAbstractOpenGLFunctions) NewFromPointer(cthis unsafe.Pointer) *QAbstractOpenGLFunctions {
	return NewQAbstractOpenGLFunctionsFromPointer(cthis)
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:213
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractOpenGLFunctions()

/*

 */
func DeleteQAbstractOpenGLFunctions(this *QAbstractOpenGLFunctions) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractOpenGLFunctionsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:215
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool initializeOpenGLFunctions()

/*

 */
func (this *QAbstractOpenGLFunctions) InitializeOpenGLFunctions() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:220
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAbstractOpenGLFunctions()

/*

 */
func (*QAbstractOpenGLFunctions) NewForInherit() *QAbstractOpenGLFunctions {
	return NewQAbstractOpenGLFunctions()
}
func NewQAbstractOpenGLFunctions() *QAbstractOpenGLFunctions {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractOpenGLFunctionsC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractOpenGLFunctionsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAbstractOpenGLFunctions)
	return gothis
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:223
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool isInitialized() const

/*

 */
func (this *QAbstractOpenGLFunctions) IsInitialized() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAbstractOpenGLFunctions13isInitializedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
}

//  keep block end
