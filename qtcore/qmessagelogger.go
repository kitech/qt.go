package qtcore

// /usr/include/qt/QtCore/qlogging.h
// #include <qlogging.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
}

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
	this.CObject = &qtrt.CObject{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMessageLoggerC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(unsafe.Pointer(uintptr(rv)))
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMessageLoggerC2EPKciS1_", ffiqt.FFI_TYPE_POINTER, convArg0, line, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(unsafe.Pointer(uintptr(rv)))
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMessageLoggerC2EPKciS1_S1_", ffiqt.FFI_TYPE_POINTER, convArg0, line, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

//  body block end
