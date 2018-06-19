package qtcore

// /usr/include/qt/QtCore/qfile.h
// #include <qfile.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QFile struct {
	*QFileDevice
}
type QFile_ITF interface {
	QFileDevice_ITF
	QFile_PTR() *QFile
}

func (ptr *QFile) QFile_PTR() *QFile { return ptr }

func (this *QFile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFileDevice.GetCthis()
	}
}
func (this *QFile) SetCthis(cthis unsafe.Pointer) {
	this.QFileDevice = NewQFileDeviceFromPointer(cthis)
}
func NewQFileFromPointer(cthis unsafe.Pointer) *QFile {
	bcthis0 := NewQFileDeviceFromPointer(cthis)
	return &QFile{bcthis0}
}
func (*QFile) NewFromPointer(cthis unsafe.Pointer) *QFile {
	return NewQFileFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfile.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QFile) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qfile.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFile()

/*
Constructs a QFile object.
*/
func NewQFile() *QFile {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFileC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFile")
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:66
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFile(const QString &)

/*
Constructs a QFile object.
*/
func NewQFile_1(name string) *QFile {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFileC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFile")
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QFile(QObject *)

/*
Constructs a QFile object.
*/
func NewQFile_2(parent QObject_ITF /*777 QObject **/) *QFile {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFileC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFile")
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:69
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QFile(const QString &, QObject *)

/*
Constructs a QFile object.
*/
func NewQFile_3(name string, parent QObject_ITF /*777 QObject **/) *QFile {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFileC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFile")
	return gothis
}

// /usr/include/qt/QtCore/qfile.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFile()

/*

 */
func DeleteQFile(this *QFile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfile.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString fileName() const

/*
Reimplemented from QFileDevice::fileName().

Returns the name set by setFileName() or to the QFile constructors.

See also setFileName() and QFileInfo::fileName().
*/
func (this *QFile) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfile.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)

/*
Sets the name of the file. The name can have no path, a relative path, or an absolute path.

Do not call this function if the file has already been opened.

If the file name has no path or a relative path, the path used will be the application's current directory path at the time of the open() call.

Example:


  QFile file;
  QDir::setCurrent("/tmp");
  file.setFileName("readme.txt");
  QDir::setCurrent("/home");
  file.open(QIODevice::ReadOnly);      // opens "/home/readme.txt" under Unix



Note that the directory separator "/" works for all operating systems supported by Qt.

See also fileName(), QFileInfo, and QDir.
*/
func (this *QFile) SetFileName(name string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfile.h:88
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QByteArray encodeName(const QString &)

/*
Converts fileName to the local 8-bit encoding determined by the user's locale. This is sufficient for file names that the user chooses. File names hard-coded into the application should only use 7-bit ASCII filename characters.

See also decodeName().
*/
func (this *QFile) EncodeName(fileName string) *QByteArray /*123*/ {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile10encodeNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QFile_EncodeName(fileName string) *QByteArray /*123*/ {
	var nilthis *QFile
	rv := nilthis.EncodeName(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:92
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QString decodeName(const QByteArray &)

/*
This does the reverse of QFile::encodeName() using localFileName.

See also encodeName().
*/
func (this *QFile) DecodeName(localFileName QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if localFileName != nil && localFileName.QByteArray_PTR() != nil {
		convArg0 = localFileName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile10decodeNameERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QFile_DecodeName(localFileName QByteArray_ITF) string {
	var nilthis *QFile
	rv := nilthis.DecodeName(localFileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:97
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QString decodeName(const char *)

/*
This does the reverse of QFile::encodeName() using localFileName.

See also encodeName().
*/
func (this *QFile) DecodeName_1(localFileName string) string {
	var convArg0 = qtrt.CString(localFileName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile10decodeNameEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QFile_DecodeName_1(localFileName string) string {
	var nilthis *QFile
	rv := nilthis.DecodeName_1(localFileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exists() const

/*
Returns true if the file specified by fileName exists; otherwise returns false.

Note: If fileName is a symlink that points to a non-existing file, false is returned.
*/
func (this *QFile) Exists() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile6existsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:108
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool exists(const QString &)

/*
Returns true if the file specified by fileName exists; otherwise returns false.

Note: If fileName is a symlink that points to a non-existing file, false is returned.
*/
func (this *QFile) Exists_1(fileName string) bool {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6existsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Exists_1(fileName string) bool {
	var nilthis *QFile
	rv := nilthis.Exists_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readLink() const

/*

 */
func (this *QFile) ReadLink() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile8readLinkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfile.h:111
// index:1
// Public static Visibility=Default Availability=Available
// [8] QString readLink(const QString &)

/*

 */
func (this *QFile) ReadLink_1(fileName string) string {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile8readLinkERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QFile_ReadLink_1(fileName string) string {
	var nilthis *QFile
	rv := nilthis.ReadLink_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString symLinkTarget() const

/*
Returns the absolute path of the file or directory referred to by the symlink (or shortcut on Windows) specified by fileName, or returns an empty string if the fileName does not correspond to a symbolic link.

This name may not represent an existing file; it is only a string. QFile::exists() returns true if the symlink points to an existing file.

This function was introduced in  Qt 4.2.
*/
func (this *QFile) SymLinkTarget() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile13symLinkTargetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfile.h:113
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QString symLinkTarget(const QString &)

/*
Returns the absolute path of the file or directory referred to by the symlink (or shortcut on Windows) specified by fileName, or returns an empty string if the fileName does not correspond to a symbolic link.

This name may not represent an existing file; it is only a string. QFile::exists() returns true if the symlink points to an existing file.

This function was introduced in  Qt 4.2.
*/
func (this *QFile) SymLinkTarget_1(fileName string) string {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile13symLinkTargetERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QFile_SymLinkTarget_1(fileName string) string {
	var nilthis *QFile
	rv := nilthis.SymLinkTarget_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove()

/*
Removes the file specified by fileName(). Returns true if successful; otherwise returns false.

The file is closed before it is removed.

See also setFileName().
*/
func (this *QFile) Remove() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6removeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:116
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool remove(const QString &)

/*
Removes the file specified by fileName(). Returns true if successful; otherwise returns false.

The file is closed before it is removed.

See also setFileName().
*/
func (this *QFile) Remove_1(fileName string) bool {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6removeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Remove_1(fileName string) bool {
	var nilthis *QFile
	rv := nilthis.Remove_1(fileName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rename(const QString &)

/*
Renames the file currently specified by fileName() to newName. Returns true if successful; otherwise returns false.

If a file with the name newName already exists, rename() returns false (i.e., QFile will not overwrite it).

The file is closed before it is renamed.

If the rename operation fails, Qt will attempt to copy this file's contents to newName, and then remove this file, keeping only newName. If that copy operation fails or this file can't be removed, the destination file newName is removed to restore the old state.

See also setFileName().
*/
func (this *QFile) Rename(newName string) bool {
	var tmpArg0 = NewQString_5(newName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6renameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:119
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool rename(const QString &, const QString &)

/*
Renames the file currently specified by fileName() to newName. Returns true if successful; otherwise returns false.

If a file with the name newName already exists, rename() returns false (i.e., QFile will not overwrite it).

The file is closed before it is renamed.

If the rename operation fails, Qt will attempt to copy this file's contents to newName, and then remove this file, keeping only newName. If that copy operation fails or this file can't be removed, the destination file newName is removed to restore the old state.

See also setFileName().
*/
func (this *QFile) Rename_1(oldName string, newName string) bool {
	var tmpArg0 = NewQString_5(oldName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(newName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6renameERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Rename_1(oldName string, newName string) bool {
	var nilthis *QFile
	rv := nilthis.Rename_1(oldName, newName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool link(const QString &)

/*
Creates a link named linkName that points to the file currently specified by fileName(). What a link is depends on the underlying filesystem (be it a shortcut on Windows or a symbolic link on Unix). Returns true if successful; otherwise returns false.

This function will not overwrite an already existing entity in the file system; in this case, link() will return false and set error() to return RenameError.

Note: To create a valid link on Windows, linkName must have a .lnk file extension.

See also setFileName().
*/
func (this *QFile) Link(newName string) bool {
	var tmpArg0 = NewQString_5(newName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4linkERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:122
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool link(const QString &, const QString &)

/*
Creates a link named linkName that points to the file currently specified by fileName(). What a link is depends on the underlying filesystem (be it a shortcut on Windows or a symbolic link on Unix). Returns true if successful; otherwise returns false.

This function will not overwrite an already existing entity in the file system; in this case, link() will return false and set error() to return RenameError.

Note: To create a valid link on Windows, linkName must have a .lnk file extension.

See also setFileName().
*/
func (this *QFile) Link_1(oldname string, newName string) bool {
	var tmpArg0 = NewQString_5(oldname)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(newName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4linkERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Link_1(oldname string, newName string) bool {
	var nilthis *QFile
	rv := nilthis.Link_1(oldname, newName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:124
// index:0
// Public Visibility=Default Availability=Available
// [1] bool copy(const QString &)

/*
Copies the file currently specified by fileName() to a file called newName. Returns true if successful; otherwise returns false.

Note that if a file with the name newName already exists, copy() returns false (i.e. QFile will not overwrite it).

The source file is closed before it is copied.

See also setFileName().
*/
func (this *QFile) Copy(newName string) bool {
	var tmpArg0 = NewQString_5(newName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4copyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:125
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool copy(const QString &, const QString &)

/*
Copies the file currently specified by fileName() to a file called newName. Returns true if successful; otherwise returns false.

Note that if a file with the name newName already exists, copy() returns false (i.e. QFile will not overwrite it).

The source file is closed before it is copied.

See also setFileName().
*/
func (this *QFile) Copy_1(fileName string, newName string) bool {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(newName)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4copyERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Copy_1(fileName string, newName string) bool {
	var nilthis *QFile
	rv := nilthis.Copy_1(fileName, newName)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)

/*
Reimplemented from QIODevice::open().

Opens the file using OpenMode mode, returning true if successful; otherwise false.

The mode must be QIODevice::ReadOnly, QIODevice::WriteOnly, or QIODevice::ReadWrite. It may also have additional flags, such as QIODevice::Text and QIODevice::Unbuffered.

Note: In WriteOnly or ReadWrite mode, if the relevant file does not already exist, this function will try to create a new file before opening it.

See also QIODevice::OpenMode and setFileName().
*/
func (this *QFile) Open(flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:129
// index:1
// Public Visibility=Default Availability=Available
// [1] bool open(int, QIODevice::OpenMode, QFileDevice::FileHandleFlags)

/*
Reimplemented from QIODevice::open().

Opens the file using OpenMode mode, returning true if successful; otherwise false.

The mode must be QIODevice::ReadOnly, QIODevice::WriteOnly, or QIODevice::ReadWrite. It may also have additional flags, such as QIODevice::Text and QIODevice::Unbuffered.

Note: In WriteOnly or ReadWrite mode, if the relevant file does not already exist, this function will try to create a new file before opening it.

See also QIODevice::OpenMode and setFileName().
*/
func (this *QFile) Open_1(fd int, ioFlags int, handleFlags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4openEi6QFlagsIN9QIODevice12OpenModeFlagEES0_IN11QFileDevice14FileHandleFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fd, ioFlags, handleFlags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:129
// index:1
// Public Visibility=Default Availability=Available
// [1] bool open(int, QIODevice::OpenMode, QFileDevice::FileHandleFlags)

/*
Reimplemented from QIODevice::open().

Opens the file using OpenMode mode, returning true if successful; otherwise false.

The mode must be QIODevice::ReadOnly, QIODevice::WriteOnly, or QIODevice::ReadWrite. It may also have additional flags, such as QIODevice::Text and QIODevice::Unbuffered.

Note: In WriteOnly or ReadWrite mode, if the relevant file does not already exist, this function will try to create a new file before opening it.

See also QIODevice::OpenMode and setFileName().
*/
func (this *QFile) Open_1_(fd int, ioFlags int) bool {
	// arg: 2, QFileDevice::FileHandleFlags=Typedef, QFileDevice::FileHandleFlags=Typedef, QFlags<QFileDevice::FileHandleFlag>, Unexposed
	handleFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile4openEi6QFlagsIN9QIODevice12OpenModeFlagEES0_IN11QFileDevice14FileHandleFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fd, ioFlags, handleFlags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:131
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 size() const

/*
Reimplemented from QIODevice::size().
*/
func (this *QFile) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qfile.h:133
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool resize(qint64)

/*
Reimplemented from QFileDevice::resize().
*/
func (this *QFile) Resize(sz int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6resizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sz)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:134
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool resize(const QString &, qint64)

/*
Reimplemented from QFileDevice::resize().
*/
func (this *QFile) Resize_1(filename string, sz int64) bool {
	var tmpArg0 = NewQString_5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile6resizeERK7QStringx", qtrt.FFI_TYPE_POINTER, convArg0, sz)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFile_Resize_1(filename string, sz int64) bool {
	var nilthis *QFile
	rv := nilthis.Resize_1(filename, sz)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:136
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QFileDevice::Permissions permissions() const

/*
Reimplemented from QFileDevice::permissions().

See also setPermissions().
*/
func (this *QFile) Permissions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFile11permissionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qfile.h:137
// index:1
// Public static Visibility=Default Availability=Available
// [4] QFileDevice::Permissions permissions(const QString &)

/*
Reimplemented from QFileDevice::permissions().

See also setPermissions().
*/
func (this *QFile) Permissions_1(filename string) int {
	var tmpArg0 = NewQString_5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile11permissionsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QFile_Permissions_1(filename string) int {
	var nilthis *QFile
	rv := nilthis.Permissions_1(filename)
	return rv
}

// /usr/include/qt/QtCore/qfile.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setPermissions(QFileDevice::Permissions)

/*
Reimplemented from QFileDevice::setPermissions().

Sets the permissions for the file to the permissions specified. Returns true if successful, or false if the permissions cannot be modified.

Warning: This function does not manipulate ACLs, which may limit its effectiveness.

See also permissions() and setFileName().
*/
func (this *QFile) SetPermissions(permissionSpec int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile14setPermissionsE6QFlagsIN11QFileDevice10PermissionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), permissionSpec)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfile.h:139
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool setPermissions(const QString &, QFileDevice::Permissions)

/*
Reimplemented from QFileDevice::setPermissions().

Sets the permissions for the file to the permissions specified. Returns true if successful, or false if the permissions cannot be modified.

Warning: This function does not manipulate ACLs, which may limit its effectiveness.

See also permissions() and setFileName().
*/
func (this *QFile) SetPermissions_1(filename string, permissionSpec int) bool {
	var tmpArg0 = NewQString_5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFile14setPermissionsERK7QString6QFlagsIN11QFileDevice10PermissionEE", qtrt.FFI_TYPE_POINTER, convArg0, permissionSpec)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFile_SetPermissions_1(filename string, permissionSpec int) bool {
	var nilthis *QFile
	rv := nilthis.SetPermissions_1(filename, permissionSpec)
	return rv
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
