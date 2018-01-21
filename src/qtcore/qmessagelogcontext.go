package qtcore

// /usr/include/qt/QtCore/qlogging.h
// #include <qlogging.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QMessageLogContext struct {
	*qtrt.CObject
}

func (this *QMessageLogContext) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQMessageLogContextFromPointer(cthis unsafe.Pointer) *QMessageLogContext {
	return &QMessageLogContext{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qlogging.h:66
// index:0
// Public inline
// void QMessageLogContext()
func NewQMessageLogContext() *QMessageLogContext {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QMessageLogContextC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLogContextFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlogging.h:68
// index:1
// Public inline
// void QMessageLogContext(const char *, int, const char *, const char *)
func NewQMessageLogContext_1(fileName string, lineNumber int, functionName string, categoryName string) *QMessageLogContext {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = qtrt.CString(fileName)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = qtrt.CString(functionName)
	defer qtrt.FreeMem(convArg2)
	var convArg3 = qtrt.CString(categoryName)
	defer qtrt.FreeMem(convArg3)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QMessageLogContextC2EPKciS1_S1_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &lineNumber, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLogContextFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlogging.h:71
// index:0
// Public
// void copy(const class QMessageLogContext &)
func (this *QMessageLogContext) Copy(logContext *QMessageLogContext) {
	var convArg0 = logContext.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QMessageLogContext4copyERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
