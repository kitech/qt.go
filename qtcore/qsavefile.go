package qtcore

// /usr/include/qt/QtCore/qsavefile.h
// #include <qsavefile.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// long long writeData(const char *, qint64)
func (this *QSaveFile) InheritWriteData(f func(data string, len_ int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

/*

 */
type QSaveFile struct {
	*QFileDevice
}
type QSaveFile_ITF interface {
	QFileDevice_ITF
	QSaveFile_PTR() *QSaveFile
}

func (ptr *QSaveFile) QSaveFile_PTR() *QSaveFile { return ptr }

func (this *QSaveFile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFileDevice.GetCthis()
	}
}
func (this *QSaveFile) SetCthis(cthis unsafe.Pointer) {
	this.QFileDevice = NewQFileDeviceFromPointer(cthis)
}
func NewQSaveFileFromPointer(cthis unsafe.Pointer) *QSaveFile {
	bcthis0 := NewQFileDeviceFromPointer(cthis)
	return &QSaveFile{bcthis0}
}
func (*QSaveFile) NewFromPointer(cthis unsafe.Pointer) *QSaveFile {
	return NewQSaveFileFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsavefile.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSaveFile) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSaveFile10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsavefile.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSaveFile(const QString &)

/*
Constructs a new file object to represent the file with the given name.
*/
func NewQSaveFile(name string) *QSaveFile {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFileC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSaveFile")
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSaveFile(QObject *)

/*
Constructs a new file object to represent the file with the given name.
*/
func NewQSaveFile_1(parent QObject_ITF /*777 QObject **/) *QSaveFile {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFileC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSaveFile")
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSaveFile(QObject *)

/*
Constructs a new file object to represent the file with the given name.
*/
func NewQSaveFile_1_() *QSaveFile {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFileC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSaveFile")
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:71
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSaveFile(const QString &, QObject *)

/*
Constructs a new file object to represent the file with the given name.
*/
func NewQSaveFile_2(name string, parent QObject_ITF /*777 QObject **/) *QSaveFile {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFileC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSaveFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSaveFile")
	return gothis
}

// /usr/include/qt/QtCore/qsavefile.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSaveFile()

/*

 */
func DeleteQSaveFile(this *QSaveFile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsavefile.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString fileName() const

/*
Reimplemented from QFileDevice::fileName().

Returns the name set by setFileName() or to the QSaveFile constructor.

See also setFileName().
*/
func (this *QSaveFile) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSaveFile8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsavefile.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)

/*
Sets the name of the file. The name can have no path, a relative path, or an absolute path.

See also QFile::setFileName() and fileName().
*/
func (this *QSaveFile) SetFileName(name string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)

/*
Reimplemented from QIODevice::open().

Opens the file using OpenMode mode, returning true if successful; otherwise false.

Important: the mode must include QIODevice::WriteOnly. It may also have additional flags, such as QIODevice::Text and QIODevice::Unbuffered.

QIODevice::ReadWrite and QIODevice::Append are not supported at the moment.

See also QIODevice::OpenMode and setFileName().
*/
func (this *QSaveFile) Open(flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsavefile.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool commit()

/*
Commits the changes to disk, if all previous writes were successful.

It is mandatory to call this at the end of the saving operation, otherwise the file will be discarded.

If an error happened during writing, deletes the temporary file and returns false. Otherwise, renames it to the final fileName and returns true on success. Finally, closes the device.

See also cancelWriting().
*/
func (this *QSaveFile) Commit() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile6commitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsavefile.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancelWriting()

/*
Cancels writing the new file.

If the application changes its mind while saving, it can call cancelWriting(), which sets an error code so that commit() will discard the temporary file.

Alternatively, it can simply make sure not to call commit().

Further write operations are possible after calling this method, but none of it will have any effect, the written file will be discarded.

This method has no effect when direct write fallback is used. This is the case when saving over an existing file in a readonly directory: no temporary file can be created, so the existing file is overwritten no matter what, and cancelWriting() cannot do anything about that, the contents of the existing file will be lost.

See also commit().
*/
func (this *QSaveFile) CancelWriting() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile13cancelWritingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDirectWriteFallback(bool)

/*
Allows writing over the existing file if necessary.

QSaveFile creates a temporary file in the same directory as the final file and atomically renames it. However this is not possible if the directory permissions do not allow creating new files. In order to preserve atomicity guarantees, open() fails when it cannot create the temporary file.

In order to allow users to edit files with write permissions in a directory with restricted permissions, call setDirectWriteFallback() with enabled set to true, and the following calls to open() will fallback to opening the existing file directly and writing into it, without the use of a temporary file. This does not have atomicity guarantees, i.e. an application crash or for instance a power failure could lead to a partially-written file on disk. It also means cancelWriting() has no effect, in such a case.

Typically, to save documents edited by the user, call setDirectWriteFallback(true), and to save application internal files (configuration files, data files, ...), keep the default setting which ensures atomicity.

See also directWriteFallback().
*/
func (this *QSaveFile) SetDirectWriteFallback(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile22setDirectWriteFallbackEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsavefile.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool directWriteFallback() const

/*
Returns true if the fallback solution for saving files in read-only directories is enabled.

See also setDirectWriteFallback().
*/
func (this *QSaveFile) DirectWriteFallback() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSaveFile19directWriteFallbackEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsavefile.h:87
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)

/*
Reimplemented from QIODevice::writeData().
*/
func (this *QSaveFile) WriteData(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSaveFile9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

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
