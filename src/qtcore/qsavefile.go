//  header block begin
// /usr/include/qt/QtCore/qsavefile.h
// #include <qsavefile.h>
// #include <QtCore>
package qtcore

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
type QSaveFile struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qsavefile.h:62
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSaveFile) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSaveFile10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:68
// index:0
// void QSaveFile(const class QString &)
func NewQSaveFile(name unsafe.Pointer) *QSaveFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, name)
	gopp.ErrPrint(err, rv)
	return &QSaveFile{cthis}
}

// /usr/include/qt/QtCore/qsavefile.h:70
// index:1
// void QSaveFile(class QObject *)
func NewQSaveFile_1(parent unsafe.Pointer) *QSaveFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QSaveFile{cthis}
}

// /usr/include/qt/QtCore/qsavefile.h:71
// index:2
// void QSaveFile(const class QString &, class QObject *)
func NewQSaveFile_2(name unsafe.Pointer, parent unsafe.Pointer) *QSaveFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, name, parent)
	gopp.ErrPrint(err, rv)
	return &QSaveFile{cthis}
}

// /usr/include/qt/QtCore/qsavefile.h:73
// index:0
// virtual
// void ~QSaveFile()
func DeleteQSaveFile(*QSaveFile) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:75
// index:0
// virtual
// QString fileName()
func (this *QSaveFile) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSaveFile8fileNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:76
// index:0
// void setFileName(const class QString &)
func (this *QSaveFile) SetFileName(name unsafe.Pointer) {
	// 0: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile11setFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:79
// index:0
// bool commit()
func (this *QSaveFile) Commit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile6commitEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:81
// index:0
// void cancelWriting()
func (this *QSaveFile) CancelWriting() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile13cancelWritingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:83
// index:0
// void setDirectWriteFallback(_Bool)
func (this *QSaveFile) SetDirectWriteFallback(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile22setDirectWriteFallbackEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:84
// index:0
// bool directWriteFallback()
func (this *QSaveFile) DirectWriteFallback() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSaveFile19directWriteFallbackEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
