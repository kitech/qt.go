package qtcore

// /usr/include/qt/QtCore/qlogging.h
// #include <qlogging.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QMessageLogger struct {
	*qtrt.CObject
}

func (this *QMessageLogger) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMessageLogger) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMessageLoggerFromPointer(cthis unsafe.Pointer) *QMessageLogger {
	return &QMessageLogger{&qtrt.CObject{cthis}}
}
func (*QMessageLogger) NewFromPointer(cthis unsafe.Pointer) *QMessageLogger {
	return NewQMessageLoggerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qlogging.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QMessageLogger()
func NewQMessageLogger() *QMessageLogger {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMessageLoggerC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMessageLogger)
	return gothis
}

// /usr/include/qt/QtCore/qlogging.h:91
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QMessageLogger(const char *, int, const char *)
func NewQMessageLogger_1(file string, line int, function string) *QMessageLogger {
	var convArg0 = qtrt.CString(file)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = qtrt.CString(function)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMessageLoggerC2EPKciS1_", qtrt.FFI_TYPE_POINTER, convArg0, line, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMessageLogger)
	return gothis
}

// /usr/include/qt/QtCore/qlogging.h:93
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QMessageLogger(const char *, int, const char *, const char *)
func NewQMessageLogger_2(file string, line int, function string, category string) *QMessageLogger {
	var convArg0 = qtrt.CString(file)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = qtrt.CString(function)
	defer qtrt.FreeMem(convArg2)
	var convArg3 = qtrt.CString(category)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMessageLoggerC2EPKciS1_S1_", qtrt.FFI_TYPE_POINTER, convArg0, line, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMessageLogger)
	return gothis
}

func DeleteQMessageLogger(this *QMessageLogger) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QMessageLoggerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
}

//  keep block end
