//  header block begin
// /usr/include/qt/QtCore/qlogging.h
// #include <qlogging.h>
// #include <QtCore>
package qtcore

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qlogging.h:66
// index:0
// inline
// void QMessageLogContext()
func NewQMessageLogContext() *QMessageLogContext {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QMessageLogContextC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QMessageLogContext{cthis}
}

// /usr/include/qt/QtCore/qlogging.h:68
// index:1
// inline
// void QMessageLogContext(const char *, int, const char *, const char *)
func NewQMessageLogContext_1(fileName unsafe.Pointer, lineNumber int, functionName unsafe.Pointer, categoryName unsafe.Pointer) *QMessageLogContext {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QMessageLogContextC2EPKciS1_S1_", ffiqt.FFI_TYPE_VOID, cthis, fileName, &lineNumber, functionName, categoryName)
	gopp.ErrPrint(err, rv)
	return &QMessageLogContext{cthis}
}

// /usr/include/qt/QtCore/qlogging.h:71
// index:0
// void copy(const class QMessageLogContext &)
func (this *QMessageLogContext) Copy(logContext unsafe.Pointer) {
	// 0: (, const QMessageLogContext & logContext), (logContext)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QMessageLogContext4copyERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, logContext)
	gopp.ErrPrint(err, rv)
}

//  body block end
