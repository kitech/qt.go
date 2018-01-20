//  header block begin
// /usr/include/qt/QtCore/qdiriterator.h
// #include <qdiriterator.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 66
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
type QDirIterator struct {
	*qtrt.CObject
}

func (this *QDirIterator) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qdiriterator.h:58
// index:0
// void QDirIterator(const class QDir &, QDirIterator::IteratorFlags)
func NewQDirIterator(dir unsafe.Pointer, flags int) *QDirIterator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK4QDir6QFlagsINS_12IteratorFlagEE", ffiqt.FFI_TYPE_VOID, cthis, dir, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(cthis)
	return gothis
}
func NewQDirIteratorFromPointer(cthis unsafe.Pointer) *QDirIterator {
	return &QDirIterator{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qdiriterator.h:59
// index:1
// void QDirIterator(const class QString &, QDirIterator::IteratorFlags)
func NewQDirIterator_1(path unsafe.Pointer, flags int) *QDirIterator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QString6QFlagsINS_12IteratorFlagEE", ffiqt.FFI_TYPE_VOID, cthis, path, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:61
// index:2
// void QDirIterator(const class QString &, class QDir::Filters, QDirIterator::IteratorFlags)
func NewQDirIterator_2(path unsafe.Pointer, filter int, flags int) *QDirIterator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QString6QFlagsIN4QDir6FilterEES3_INS_12IteratorFlagEE", ffiqt.FFI_TYPE_VOID, cthis, path, &filter, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:64
// index:3
// void QDirIterator(const class QString &, const class QStringList &, class QDir::Filters, QDirIterator::IteratorFlags)
func NewQDirIterator_3(path unsafe.Pointer, nameFilters unsafe.Pointer, filters int, flags int) *QDirIterator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QStringRK11QStringList6QFlagsIN4QDir6FilterEES6_INS_12IteratorFlagEE", ffiqt.FFI_TYPE_VOID, cthis, path, nameFilters, &filters, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:69
// index:0
// void ~QDirIterator()
func DeleteQDirIterator(*QDirIterator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QDirIteratorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdiriterator.h:71
// index:0
// QString next()
func (this *QDirIterator) Next() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QDirIterator4nextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdiriterator.h:72
// index:0
// bool hasNext()
func (this *QDirIterator) HasNext() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QDirIterator7hasNextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdiriterator.h:74
// index:0
// QString fileName()
func (this *QDirIterator) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QDirIterator8fileNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdiriterator.h:75
// index:0
// QString filePath()
func (this *QDirIterator) FilePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QDirIterator8filePathEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdiriterator.h:76
// index:0
// QFileInfo fileInfo()
func (this *QDirIterator) FileInfo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QDirIterator8fileInfoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdiriterator.h:77
// index:0
// QString path()
func (this *QDirIterator) Path() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QDirIterator4pathEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
