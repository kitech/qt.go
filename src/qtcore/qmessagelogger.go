//  header block begin
// /usr/include/qt/QtCore/qlogging.h
// #include <qlogging.h>
// #include <QtCore>
package qtcore

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
	return this.Cthis
}

// /usr/include/qt/QtCore/qlogging.h:90
// index:0
// inline
// void QMessageLogger()
func NewQMessageLogger() *QMessageLogger {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMessageLoggerC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(cthis)
	return gothis
}
func NewQMessageLoggerFromPointer(cthis unsafe.Pointer) *QMessageLogger {
	return &QMessageLogger{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qlogging.h:91
// index:1
// inline
// void QMessageLogger(const char *, int, const char *)
func NewQMessageLogger_1(file unsafe.Pointer, line int, function unsafe.Pointer) *QMessageLogger {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMessageLoggerC2EPKciS1_", ffiqt.FFI_TYPE_VOID, cthis, file, &line, function)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlogging.h:93
// index:2
// inline
// void QMessageLogger(const char *, int, const char *, const char *)
func NewQMessageLogger_2(file unsafe.Pointer, line int, function unsafe.Pointer, category unsafe.Pointer) *QMessageLogger {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMessageLoggerC2EPKciS1_S1_", ffiqt.FFI_TYPE_VOID, cthis, file, &line, function, category)
	gopp.ErrPrint(err, rv)
	gothis := NewQMessageLoggerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlogging.h:120
// index:0
// QDebug debug()
func (this *QMessageLogger) Debug() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMessageLogger5debugEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:121
// index:1
// QDebug debug(const class QLoggingCategory &)
func (this *QMessageLogger) Debug_1(cat unsafe.Pointer) {
	// 1: (, cat const QLoggingCategory &), (cat)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMessageLogger5debugERK16QLoggingCategory", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:123
// index:0
// QDebug info()
func (this *QMessageLogger) Info() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMessageLogger4infoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:124
// index:1
// QDebug info(const class QLoggingCategory &)
func (this *QMessageLogger) Info_1(cat unsafe.Pointer) {
	// 1: (, cat const QLoggingCategory &), (cat)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMessageLogger4infoERK16QLoggingCategory", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:126
// index:0
// QDebug warning()
func (this *QMessageLogger) Warning() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMessageLogger7warningEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:127
// index:1
// QDebug warning(const class QLoggingCategory &)
func (this *QMessageLogger) Warning_1(cat unsafe.Pointer) {
	// 1: (, cat const QLoggingCategory &), (cat)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMessageLogger7warningERK16QLoggingCategory", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:129
// index:0
// QDebug critical()
func (this *QMessageLogger) Critical() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMessageLogger8criticalEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:130
// index:1
// QDebug critical(const class QLoggingCategory &)
func (this *QMessageLogger) Critical_1(cat unsafe.Pointer) {
	// 1: (, cat const QLoggingCategory &), (cat)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMessageLogger8criticalERK16QLoggingCategory", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:133
// index:0
// QNoDebug noDebug()
func (this *QMessageLogger) NoDebug() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMessageLogger7noDebugEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
