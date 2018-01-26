package qtgui

// /usr/include/qt/QtGui/qopenglversionfunctions.h
// #include <qopenglversionfunctions.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAbstractOpenGLFunctions) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQAbstractOpenGLFunctionsFromPointer(cthis unsafe.Pointer) *QAbstractOpenGLFunctions {
	return &QAbstractOpenGLFunctions{&qtrt.CObject{cthis}}
}
func (*QAbstractOpenGLFunctions) NewFromPointer(cthis unsafe.Pointer) *QAbstractOpenGLFunctions {
	return NewQAbstractOpenGLFunctionsFromPointer(cthis)
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
func (this *QAbstractOpenGLFunctions) InitializeOpenGLFunctions() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractOpenGLFunctions25initializeOpenGLFunctionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
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
func (this *QAbstractOpenGLFunctions) IsInitialized() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAbstractOpenGLFunctions13isInitializedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
