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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
func NewQMessageLoggerFromPointer(cthis unsafe.Pointer) *QMessageLogger {
	return &QMessageLogger{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qlogging.h:90
// index:0
// Public inline
// void QMessageLogger()
func NewQMessageLogger() *QMessageLogger {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMessageLoggerC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlogging.h:91
// index:1
// Public inline
// void QMessageLogger(const char *, int, const char *)
func NewQMessageLogger_1(file string, line int, function string) *QMessageLogger {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = qtrt.CString(file)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = qtrt.CString(function)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMessageLoggerC2EPKciS1_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &line, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlogging.h:93
// index:2
// Public inline
// void QMessageLogger(const char *, int, const char *, const char *)
func NewQMessageLogger_2(file string, line int, function string, category string) *QMessageLogger {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = qtrt.CString(file)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = qtrt.CString(function)
	defer qtrt.FreeMem(convArg2)
	var convArg3 = qtrt.CString(category)
	defer qtrt.FreeMem(convArg3)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMessageLoggerC2EPKciS1_S1_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &line, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(cthis)
	return gothis
}

//  body block end
