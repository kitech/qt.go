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
// extern C begin: 71
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QDirIterator struct {
	*qtrt.CObject
}
type QDirIterator_ITF interface {
	QDirIterator_PTR() *QDirIterator
}

func (ptr *QDirIterator) QDirIterator_PTR() *QDirIterator { return ptr }

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

/*
Constructs a QDirIterator that can iterate over dir's entrylist, using dir's name filters and regular filters. You can pass options via flags to decide how the directory should be iterated.

By default, flags is NoIteratorFlags, which provides the same behavior as in QDir::entryList().

The sorting in dir is ignored.

Note: To list symlinks that point to non existing files, QDir::System must be passed to the flags.

See also hasNext(), next(), and IteratorFlags.
*/
func NewQDirIterator(dir QDir_ITF, flags int) *QDirIterator {
	var convArg0 unsafe.Pointer
	if dir != nil && dir.QDir_PTR() != nil {
		convArg0 = dir.QDir_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK4QDir6QFlagsINS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDirIterator)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDirIterator(const QDir &, QDirIterator::IteratorFlags)

/*
Constructs a QDirIterator that can iterate over dir's entrylist, using dir's name filters and regular filters. You can pass options via flags to decide how the directory should be iterated.

By default, flags is NoIteratorFlags, which provides the same behavior as in QDir::entryList().

The sorting in dir is ignored.

Note: To list symlinks that point to non existing files, QDir::System must be passed to the flags.

See also hasNext(), next(), and IteratorFlags.
*/
func NewQDirIterator__(dir QDir_ITF) *QDirIterator {
	var convArg0 unsafe.Pointer
	if dir != nil && dir.QDir_PTR() != nil {
		convArg0 = dir.QDir_PTR().GetCthis()
	}
	// arg: 1, QDirIterator::IteratorFlags=Typedef, QDirIterator::IteratorFlags=Typedef, QFlags<QDirIterator::IteratorFlag>, Unexposed
	flags := 0
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

/*
Constructs a QDirIterator that can iterate over dir's entrylist, using dir's name filters and regular filters. You can pass options via flags to decide how the directory should be iterated.

By default, flags is NoIteratorFlags, which provides the same behavior as in QDir::entryList().

The sorting in dir is ignored.

Note: To list symlinks that point to non existing files, QDir::System must be passed to the flags.

See also hasNext(), next(), and IteratorFlags.
*/
func NewQDirIterator_1(path string, flags int) *QDirIterator {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QString6QFlagsINS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDirIterator)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDirIterator(const QString &, QDirIterator::IteratorFlags)

/*
Constructs a QDirIterator that can iterate over dir's entrylist, using dir's name filters and regular filters. You can pass options via flags to decide how the directory should be iterated.

By default, flags is NoIteratorFlags, which provides the same behavior as in QDir::entryList().

The sorting in dir is ignored.

Note: To list symlinks that point to non existing files, QDir::System must be passed to the flags.

See also hasNext(), next(), and IteratorFlags.
*/
func NewQDirIterator_1_(path string) *QDirIterator {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QDirIterator::IteratorFlags=Typedef, QDirIterator::IteratorFlags=Typedef, QFlags<QDirIterator::IteratorFlag>, Unexposed
	flags := 0
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

/*
Constructs a QDirIterator that can iterate over dir's entrylist, using dir's name filters and regular filters. You can pass options via flags to decide how the directory should be iterated.

By default, flags is NoIteratorFlags, which provides the same behavior as in QDir::entryList().

The sorting in dir is ignored.

Note: To list symlinks that point to non existing files, QDir::System must be passed to the flags.

See also hasNext(), next(), and IteratorFlags.
*/
func NewQDirIterator_2(path string, filter int, flags int) *QDirIterator {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QString6QFlagsIN4QDir6FilterEES3_INS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, filter, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDirIterator)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:61
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDirIterator(const QString &, QDir::Filters, QDirIterator::IteratorFlags)

/*
Constructs a QDirIterator that can iterate over dir's entrylist, using dir's name filters and regular filters. You can pass options via flags to decide how the directory should be iterated.

By default, flags is NoIteratorFlags, which provides the same behavior as in QDir::entryList().

The sorting in dir is ignored.

Note: To list symlinks that point to non existing files, QDir::System must be passed to the flags.

See also hasNext(), next(), and IteratorFlags.
*/
func NewQDirIterator_2_(path string, filter int) *QDirIterator {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QDirIterator::IteratorFlags=Typedef, QDirIterator::IteratorFlags=Typedef, QFlags<QDirIterator::IteratorFlag>, Unexposed
	flags := 0
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

/*
Constructs a QDirIterator that can iterate over dir's entrylist, using dir's name filters and regular filters. You can pass options via flags to decide how the directory should be iterated.

By default, flags is NoIteratorFlags, which provides the same behavior as in QDir::entryList().

The sorting in dir is ignored.

Note: To list symlinks that point to non existing files, QDir::System must be passed to the flags.

See also hasNext(), next(), and IteratorFlags.
*/
func NewQDirIterator_3(path string, nameFilters QStringList_ITF, filters int, flags int) *QDirIterator {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg1 = nameFilters.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QStringRK11QStringList6QFlagsIN4QDir6FilterEES6_INS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, filters, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDirIterator)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:64
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDirIterator(const QString &, const QStringList &, QDir::Filters, QDirIterator::IteratorFlags)

/*
Constructs a QDirIterator that can iterate over dir's entrylist, using dir's name filters and regular filters. You can pass options via flags to decide how the directory should be iterated.

By default, flags is NoIteratorFlags, which provides the same behavior as in QDir::entryList().

The sorting in dir is ignored.

Note: To list symlinks that point to non existing files, QDir::System must be passed to the flags.

See also hasNext(), next(), and IteratorFlags.
*/
func NewQDirIterator_3_(path string, nameFilters QStringList_ITF) *QDirIterator {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg1 = nameFilters.QStringList_PTR().GetCthis()
	}
	// arg: 2, QDir::Filters=Elaborated, QDir::Filters=Typedef, QFlags<QDir::Filter>, Unexposed
	filters := 0
	// arg: 3, QDirIterator::IteratorFlags=Typedef, QDirIterator::IteratorFlags=Typedef, QFlags<QDirIterator::IteratorFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QDirIteratorC2ERK7QStringRK11QStringList6QFlagsIN4QDir6FilterEES6_INS_12IteratorFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, filters, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDirIterator)
	return gothis
}

// /usr/include/qt/QtCore/qdiriterator.h:64
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDirIterator(const QString &, const QStringList &, QDir::Filters, QDirIterator::IteratorFlags)

/*
Constructs a QDirIterator that can iterate over dir's entrylist, using dir's name filters and regular filters. You can pass options via flags to decide how the directory should be iterated.

By default, flags is NoIteratorFlags, which provides the same behavior as in QDir::entryList().

The sorting in dir is ignored.

Note: To list symlinks that point to non existing files, QDir::System must be passed to the flags.

See also hasNext(), next(), and IteratorFlags.
*/
func NewQDirIterator_3_1(path string, nameFilters QStringList_ITF, filters int) *QDirIterator {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg1 = nameFilters.QStringList_PTR().GetCthis()
	}
	// arg: 3, QDirIterator::IteratorFlags=Typedef, QDirIterator::IteratorFlags=Typedef, QFlags<QDirIterator::IteratorFlag>, Unexposed
	flags := 0
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

/*

 */
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

/*
Advances the iterator to the next entry, and returns the file path of this new entry. If hasNext() returns false, this function does nothing, and returns an empty QString.

You can call fileName() or filePath() to get the current entry file name or path, or fileInfo() to get a QFileInfo for the current entry.

See also hasNext(), fileName(), filePath(), and fileInfo().
*/
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
// [1] bool hasNext() const

/*
Returns true if there is at least one more entry in the directory; otherwise, false is returned.

See also next(), fileName(), filePath(), and fileInfo().
*/
func (this *QDirIterator) HasNext() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QDirIterator7hasNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdiriterator.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*
Returns the file name for the current directory entry, without the path prepended.

This function is convenient when iterating a single directory. When using the QDirIterator::Subdirectories flag, you can use filePath() to get the full path.

See also filePath() and fileInfo().
*/
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
// [8] QString filePath() const

/*
Returns the full file path for the current directory entry.

See also fileInfo() and fileName().
*/
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
// [8] QFileInfo fileInfo() const

/*
Returns a QFileInfo for the current directory entry.

See also filePath() and fileName().
*/
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
// [8] QString path() const

/*
Returns the base directory of the iterator.
*/
func (this *QDirIterator) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QDirIterator4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

/*


 */
type QDirIterator__IteratorFlag = int

//
const QDirIterator__NoIteratorFlags QDirIterator__IteratorFlag = 0

//
const QDirIterator__FollowSymlinks QDirIterator__IteratorFlag = 1

//
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
