package qtcore

// /usr/include/qt/QtCore/qdiriterator.h
// #include <qdiriterator.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 65
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QDirIterator struct {
	*qtrt.CObject
}

func (this *QDirIterator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDirIterator) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDirIteratorFromPointer(cthis unsafe.Pointer) *QDirIterator {
	return &QDirIterator{&qtrt.CObject{cthis}}
}
func (*QDirIterator) NewFromPointer(cthis unsafe.Pointer) *QDirIterator {
	return NewQDirIteratorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qdiriterator.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDirIterator(const QDir &, QDirIterator::IteratorFlags)
func NewQDirIterator(dir *QDir, flags int) *QDirIterator {
	var convArg0 = dir.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK4QDir6QFlagsINS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDirIterator)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDirIterator(const QString &, QDirIterator::IteratorFlags)
func NewQDirIterator_1(path string, flags int) *QDirIterator {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QString6QFlagsINS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDirIterator)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:61
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDirIterator(const QString &, QDir::Filters, QDirIterator::IteratorFlags)
func NewQDirIterator_2(path string, filter int, flags int) *QDirIterator {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QString6QFlagsIN4QDir6FilterEES3_INS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, filter, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDirIterator)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:64
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDirIterator(const QString &, const QStringList &, QDir::Filters, QDirIterator::IteratorFlags)
func NewQDirIterator_3(path string, nameFilters *QStringList, filters int, flags int) *QDirIterator {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = nameFilters.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QStringRK11QStringList6QFlagsIN4QDir6FilterEES6_INS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, filters, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDirIterator)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDirIterator()
func DeleteQDirIterator(this *QDirIterator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIteratorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qdiriterator.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString next()
func (this *QDirIterator) Next() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIterator4nextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdiriterator.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasNext()
func (this *QDirIterator) HasNext() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QDirIterator7hasNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdiriterator.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QDirIterator) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QDirIterator8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdiriterator.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath()
func (this *QDirIterator) FilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QDirIterator8filePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qdiriterator.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QFileInfo fileInfo()
func (this *QDirIterator) FileInfo() *QFileInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QDirIterator8fileInfoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFileInfo)
	return rv2
}

// /usr/include/qt/QtCore/qdiriterator.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path()
func (this *QDirIterator) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QDirIterator4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

type QDirIterator__IteratorFlag = int

const QDirIterator__NoIteratorFlags QDirIterator__IteratorFlag = 0
const QDirIterator__FollowSymlinks QDirIterator__IteratorFlag = 1
const QDirIterator__Subdirectories QDirIterator__IteratorFlag = 2

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
