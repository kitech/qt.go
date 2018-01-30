package qtcore

// /usr/include/qt/QtCore/qsavefile.h
// #include <qsavefile.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QSaveFile struct {
	*qtrt.CObject
}

func (this *QSaveFile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSaveFile) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSaveFileFromPointer(cthis unsafe.Pointer) *QSaveFile {
	return &QSaveFile{&qtrt.CObject{cthis}}
}
func (*QSaveFile) NewFromPointer(cthis unsafe.Pointer) *QSaveFile {
	return NewQSaveFileFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsavefile.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSaveFile) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSaveFile10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsavefile.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSaveFile(const QString &)
func NewQSaveFile(name *QString) *QSaveFile {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSaveFile(QObject *)
func NewQSaveFile_1(parent *QObject /*777 QObject **/) *QSaveFile {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:71
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSaveFile(const QString &, QObject *)
func NewQSaveFile_2(name *QString, parent *QObject /*777 QObject **/) *QSaveFile {
	var convArg0 = name.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSaveFile()
func DeleteQSaveFile(*QSaveFile) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFileD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QSaveFile) FileName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSaveFile8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qsavefile.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QSaveFile) SetFileName(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool commit()
func (this *QSaveFile) Commit() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile6commitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsavefile.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancelWriting()
func (this *QSaveFile) CancelWriting() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile13cancelWritingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDirectWriteFallback(_Bool)
func (this *QSaveFile) SetDirectWriteFallback(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile22setDirectWriteFallbackEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool directWriteFallback()
func (this *QSaveFile) DirectWriteFallback() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSaveFile19directWriteFallbackEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsavefile.h:87
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)
func (this *QSaveFile) WriteData(data string, len int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSaveFile9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

//  body block end
