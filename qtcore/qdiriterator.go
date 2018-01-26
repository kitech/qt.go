package qtcore

// /usr/include/qt/QtCore/qdiriterator.h
// #include <qdiriterator.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 54
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQDirIteratorFromPointer(cthis unsafe.Pointer) *QDirIterator {
	return &QDirIterator{&qtrt.CObject{cthis}}
}
func (*QDirIterator) NewFromPointer(cthis unsafe.Pointer) *QDirIterator {
	return NewQDirIteratorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qdiriterator.h:58
// index:0
// Public
// void QDirIterator(const class QDir &, QDirIterator::IteratorFlags)
func NewQDirIterator(dir *QDir, flags int) *QDirIterator {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = dir.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK4QDir6QFlagsINS_12IteratorFlagEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:59
// index:1
// Public
// void QDirIterator(const class QString &, QDirIterator::IteratorFlags)
func NewQDirIterator_1(path *QString, flags int) *QDirIterator {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QString6QFlagsINS_12IteratorFlagEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:69
// index:0
// Public
// void ~QDirIterator()
func DeleteQDirIterator(*QDirIterator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QDirIteratorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdiriterator.h:71
// index:0
// Public
// QString next()
func (this *QDirIterator) Next() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QDirIterator4nextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdiriterator.h:72
// index:0
// Public
// bool hasNext()
func (this *QDirIterator) HasNext() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QDirIterator7hasNextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdiriterator.h:74
// index:0
// Public
// QString fileName()
func (this *QDirIterator) FileName() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QDirIterator8fileNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdiriterator.h:75
// index:0
// Public
// QString filePath()
func (this *QDirIterator) FilePath() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QDirIterator8filePathEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdiriterator.h:76
// index:0
// Public
// QFileInfo fileInfo()
func (this *QDirIterator) FileInfo() *QFileInfo /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QDirIterator8fileInfoEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdiriterator.h:77
// index:0
// Public
// QString path()
func (this *QDirIterator) Path() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QDirIterator4pathEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QDirIterator__IteratorFlag = int

const QDirIterator__NoIteratorFlags QDirIterator__IteratorFlag = 0
const QDirIterator__FollowSymlinks QDirIterator__IteratorFlag = 1
const QDirIterator__Subdirectories QDirIterator__IteratorFlag = 2

//  body block end
