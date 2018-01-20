//  header block begin
// /usr/include/qt/QtGui/qopenglversionfunctions.h
// #include <qopenglversionfunctions.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QAbstractOpenGLFunctions struct {
	*qtrt.CObject
}

func (this *QAbstractOpenGLFunctions) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQAbstractOpenGLFunctionsFromPointer(cthis unsafe.Pointer) *QAbstractOpenGLFunctions {
	return &QAbstractOpenGLFunctions{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:213
// index:0
// Public virtual
// void ~QAbstractOpenGLFunctions()
func DeleteQAbstractOpenGLFunctions(*QAbstractOpenGLFunctions) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractOpenGLFunctionsD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:215
// index:0
// Public virtual
// bool initializeOpenGLFunctions()
func (this *QAbstractOpenGLFunctions) InitializeOpenGLFunctions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:217
// index:0
// Public inline
// QAbstractOpenGLFunctionsPrivate * d_func()
func (this *QAbstractOpenGLFunctions) D_func() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractOpenGLFunctions6d_funcEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:217
// index:1
// Public inline
// const QAbstractOpenGLFunctionsPrivate * d_func()
func (this *QAbstractOpenGLFunctions) D_func_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAbstractOpenGLFunctions6d_funcEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:220
// index:0
// Protected
// void QAbstractOpenGLFunctions()
func NewQAbstractOpenGLFunctions() *QAbstractOpenGLFunctions {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractOpenGLFunctionsC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractOpenGLFunctionsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:223
// index:0
// Protected
// bool isInitialized()
func (this *QAbstractOpenGLFunctions) IsInitialized() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAbstractOpenGLFunctions13isInitializedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:225
// index:0
// Protected
// void setOwningContext(const class QOpenGLContext *)
func (this *QAbstractOpenGLFunctions) SetOwningContext(context unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractOpenGLFunctions16setOwningContextEPK14QOpenGLContext", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), context)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qopenglversionfunctions.h:226
// index:0
// Protected
// QOpenGLContext * owningContext()
func (this *QAbstractOpenGLFunctions) OwningContext() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAbstractOpenGLFunctions13owningContextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
