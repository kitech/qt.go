//  header block begin
// /usr/include/qt/QtCore/qsavefile.h
// #include <qsavefile.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
	*qtrt.CObject
}

func (this *QSaveFile) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQSaveFileFromPointer(cthis unsafe.Pointer) *QSaveFile {
	return &QSaveFile{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qsavefile.h:62
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSaveFile) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSaveFile10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsavefile.h:68
// index:0
// Public
// void QSaveFile(const class QString &)
func NewQSaveFile(name *QString) *QSaveFile {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:70
// index:1
// Public
// void QSaveFile(class QObject *)
func NewQSaveFile_1(parent unsafe.Pointer) *QSaveFile {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:71
// index:2
// Public
// void QSaveFile(const class QString &, class QObject *)
func NewQSaveFile_2(name *QString, parent unsafe.Pointer) *QSaveFile {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:73
// index:0
// Public virtual
// void ~QSaveFile()
func DeleteQSaveFile(*QSaveFile) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:75
// index:0
// Public virtual
// QString fileName()
func (this *QSaveFile) FileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSaveFile8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsavefile.h:76
// index:0
// Public
// void setFileName(const class QString &)
func (this *QSaveFile) SetFileName(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:79
// index:0
// Public
// bool commit()
func (this *QSaveFile) Commit() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile6commitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsavefile.h:81
// index:0
// Public
// void cancelWriting()
func (this *QSaveFile) CancelWriting() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile13cancelWritingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:83
// index:0
// Public
// void setDirectWriteFallback(_Bool)
func (this *QSaveFile) SetDirectWriteFallback(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile22setDirectWriteFallbackEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:84
// index:0
// Public
// bool directWriteFallback()
func (this *QSaveFile) DirectWriteFallback() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSaveFile19directWriteFallbackEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsavefile.h:87
// index:0
// Protected virtual
// qint64 writeData(const char *, qint64)
func (this *QSaveFile) WriteData(data string, len int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
