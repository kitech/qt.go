package qtcore

// /usr/include/qt/QtCore/qfileinfo.h
// #include <qfileinfo.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
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
type QFileInfo struct {
	*qtrt.CObject
}
type QFileInfo_ITF interface {
	QFileInfo_PTR() *QFileInfo
}

func (ptr *QFileInfo) QFileInfo_PTR() *QFileInfo { return ptr }

func (this *QFileInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFileInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFileInfoFromPointer(cthis unsafe.Pointer) *QFileInfo {
	return &QFileInfo{&qtrt.CObject{cthis}}
}
func (*QFileInfo) NewFromPointer(cthis unsafe.Pointer) *QFileInfo {
	return NewQFileInfoFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfileinfo.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileInfo()

/*
Constructs an empty QFileInfo object.

Note that an empty QFileInfo object contain no file reference.

See also setFile().
*/
func (*QFileInfo) NewForInherit() *QFileInfo {
	return NewQFileInfo()
}
func NewQFileInfo() *QFileInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileInfo)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileInfo(const QString &)

/*
Constructs an empty QFileInfo object.

Note that an empty QFileInfo object contain no file reference.

See also setFile().
*/
func (*QFileInfo) NewForInherit1(file string) *QFileInfo {
	return NewQFileInfo1(file)
}
func NewQFileInfo1(file string) *QFileInfo {
	var tmpArg0 = NewQString5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileInfo)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:64
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QFileInfo(const QFile &)

/*
Constructs an empty QFileInfo object.

Note that an empty QFileInfo object contain no file reference.

See also setFile().
*/
func (*QFileInfo) NewForInherit2(file QFile_ITF) *QFileInfo {
	return NewQFileInfo2(file)
}
func NewQFileInfo2(file QFile_ITF) *QFileInfo {
	var convArg0 unsafe.Pointer
	if file != nil && file.QFile_PTR() != nil {
		convArg0 = file.QFile_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoC2ERK5QFile", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileInfo)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:65
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QFileInfo(const QDir &, const QString &)

/*
Constructs an empty QFileInfo object.

Note that an empty QFileInfo object contain no file reference.

See also setFile().
*/
func (*QFileInfo) NewForInherit3(dir QDir_ITF, file string) *QFileInfo {
	return NewQFileInfo3(dir, file)
}
func NewQFileInfo3(dir QDir_ITF, file string) *QFileInfo {
	var convArg0 unsafe.Pointer
	if dir != nil && dir.QDir_PTR() != nil {
		convArg0 = dir.QDir_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(file)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoC2ERK4QDirRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileInfo)
	return gothis
}

// /usr/include/qt/QtCore/qfileinfo.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QFileInfo()

/*

 */
func DeleteQFileInfo(this *QFileInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfileinfo.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QFileInfo & operator=(const QFileInfo &)

/*

 */
func (this *QFileInfo) Operator_equal(fileinfo QFileInfo_ITF) *QFileInfo {
	var convArg0 unsafe.Pointer
	if fileinfo != nil && fileinfo.QFileInfo_PTR() != nil {
		convArg0 = fileinfo.QFileInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFileInfo)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:71
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QFileInfo & operator=(QFileInfo &&)

/*

 */
func (this *QFileInfo) Operator_equal1(other unsafe.Pointer /*333*/) *QFileInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfoaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFileInfo)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QFileInfo &)

/*
Swaps this file info with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QFileInfo) Swap(other QFileInfo_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFileInfo_PTR() != nil {
		convArg0 = other.QFileInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:77
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QFileInfo &) const

/*

 */
func (this *QFileInfo) Operator_equal_equal(fileinfo QFileInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if fileinfo != nil && fileinfo.QFileInfo_PTR() != nil {
		convArg0 = fileinfo.QFileInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfoeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QFileInfo &) const

/*

 */
func (this *QFileInfo) Operator_not_equal(fileinfo QFileInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if fileinfo != nil && fileinfo.QFileInfo_PTR() != nil {
		convArg0 = fileinfo.QFileInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfoneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFile(const QString &)

/*
Sets the file that the QFileInfo provides information about to file.

The file can also include an absolute or relative file path. Absolute paths begin with the directory separator (e.g. "/" under Unix) or a drive specification (under Windows). Relative file names begin with a directory name or a file name and specify a path relative to the current directory.

Example:


  QString absolute = "/local/bin";
  QString relative = "local/bin";
  QFileInfo absFile(absolute);
  QFileInfo relFile(relative);

  QDir::setCurrent(QDir::rootPath());
  // absFile and relFile now point to the same file

  QDir::setCurrent("/tmp");
  // absFile now points to "/local/bin",
  // while relFile points to "/tmp/local/bin"



See also isFile(), isRelative(), QDir::setCurrent(), and QDir::isRelativePath().
*/
func (this *QFileInfo) SetFile(file string) {
	var tmpArg0 = NewQString5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFile(const QFile &)

/*
Sets the file that the QFileInfo provides information about to file.

The file can also include an absolute or relative file path. Absolute paths begin with the directory separator (e.g. "/" under Unix) or a drive specification (under Windows). Relative file names begin with a directory name or a file name and specify a path relative to the current directory.

Example:


  QString absolute = "/local/bin";
  QString relative = "local/bin";
  QFileInfo absFile(absolute);
  QFileInfo relFile(relative);

  QDir::setCurrent(QDir::rootPath());
  // absFile and relFile now point to the same file

  QDir::setCurrent("/tmp");
  // absFile now points to "/local/bin",
  // while relFile points to "/tmp/local/bin"



See also isFile(), isRelative(), QDir::setCurrent(), and QDir::isRelativePath().
*/
func (this *QFileInfo) SetFile1(file QFile_ITF) {
	var convArg0 unsafe.Pointer
	if file != nil && file.QFile_PTR() != nil {
		convArg0 = file.QFile_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK5QFile", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:82
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setFile(const QDir &, const QString &)

/*
Sets the file that the QFileInfo provides information about to file.

The file can also include an absolute or relative file path. Absolute paths begin with the directory separator (e.g. "/" under Unix) or a drive specification (under Windows). Relative file names begin with a directory name or a file name and specify a path relative to the current directory.

Example:


  QString absolute = "/local/bin";
  QString relative = "local/bin";
  QFileInfo absFile(absolute);
  QFileInfo relFile(relative);

  QDir::setCurrent(QDir::rootPath());
  // absFile and relFile now point to the same file

  QDir::setCurrent("/tmp");
  // absFile now points to "/local/bin",
  // while relFile points to "/tmp/local/bin"



See also isFile(), isRelative(), QDir::setCurrent(), and QDir::isRelativePath().
*/
func (this *QFileInfo) SetFile2(dir QDir_ITF, file string) {
	var convArg0 unsafe.Pointer
	if dir != nil && dir.QDir_PTR() != nil {
		convArg0 = dir.QDir_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(file)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo7setFileERK4QDirRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exists() const

/*
Returns true if the file exists; otherwise returns false.

Note: If the file is a symlink that points to a non-existing file, false is returned.
*/
func (this *QFileInfo) Exists() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo6existsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:84
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool exists(const QString &)

/*
Returns true if the file exists; otherwise returns false.

Note: If the file is a symlink that points to a non-existing file, false is returned.
*/
func (this *QFileInfo) Exists1(file string) bool {
	var tmpArg0 = NewQString5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo6existsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFileInfo_Exists1(file string) bool {
	var nilthis *QFileInfo
	rv := nilthis.Exists1(file)
	return rv
}

// /usr/include/qt/QtCore/qfileinfo.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh()

/*
Refreshes the information about the file, i.e. reads in information from the file system the next time a cached property is fetched.
*/
func (this *QFileInfo) Refresh() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo7refreshEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath() const

/*
Returns the file name, including the path (which may be absolute or relative).

See also absoluteFilePath(), canonicalFilePath(), and isRelative().
*/
func (this *QFileInfo) FilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8filePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absoluteFilePath() const

/*
Returns an absolute path including the file name.

The absolute path name consists of the full path and the file name. On Unix this will always begin with the root, '/', directory. On Windows this will always begin 'D:/' where D is a drive letter, except for network shares that are not mapped to a drive letter, in which case the path will begin '//sharename/'. QFileInfo will uppercase drive letters. Note that QDir does not do this. The code snippet below shows this.


      QFileInfo fi("c:/temp/foo"); => fi.absoluteFilePath() => "C:/temp/foo"



This function returns the same as filePath(), unless isRelative() is true. In contrast to canonicalFilePath(), symbolic links or redundant "." or ".." elements are not necessarily removed.

Warning: If filePath() is empty the behavior of this function is undefined.

See also filePath(), canonicalFilePath(), and isRelative().
*/
func (this *QFileInfo) AbsoluteFilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo16absoluteFilePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString canonicalFilePath() const

/*
Returns the canonical path including the file name, i.e. an absolute path without symbolic links or redundant "." or ".." elements.

If the file does not exist, canonicalFilePath() returns an empty string.

See also filePath(), absoluteFilePath(), and dir().
*/
func (this *QFileInfo) CanonicalFilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo17canonicalFilePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*
Returns the name of the file, excluding the path.

Example:


  QFileInfo fi("/tmp/archive.tar.gz");
  QString name = fi.fileName();                // name = "archive.tar.gz"



Note that, if this QFileInfo object is given a path ending in a slash, the name of the file is considered empty.

See also isRelative(), filePath(), baseName(), and suffix().
*/
func (this *QFileInfo) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QString baseName() const

/*
Returns the base name of the file without the path.

The base name consists of all characters in the file up to (but not including) the first '.' character.

Example:


  QFileInfo fi("/tmp/archive.tar.gz");
  QString base = fi.baseName();  // base = "archive"



The base name of a file is computed equally on all platforms, independent of file naming conventions (e.g., ".bashrc" on Unix has an empty base name, and the suffix is "bashrc").

See also fileName(), suffix(), completeSuffix(), and completeBaseName().
*/
func (this *QFileInfo) BaseName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8baseNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QString completeBaseName() const

/*
Returns the complete base name of the file without the path.

The complete base name consists of all characters in the file up to (but not including) the last '.' character.

Example:


  QFileInfo fi("/tmp/archive.tar.gz");
  QString base = fi.completeBaseName();  // base = "archive.tar"



See also fileName(), suffix(), completeSuffix(), and baseName().
*/
func (this *QFileInfo) CompleteBaseName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo16completeBaseNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString suffix() const

/*
Returns the suffix (extension) of the file.

The suffix consists of all characters in the file after (but not including) the last '.'.

Example:


  QFileInfo fi("/tmp/archive.tar.gz");
  QString ext = fi.suffix();  // ext = "gz"



The suffix of a file is computed equally on all platforms, independent of file naming conventions (e.g., ".bashrc" on Unix has an empty base name, and the suffix is "bashrc").

See also fileName(), completeSuffix(), baseName(), and completeBaseName().
*/
func (this *QFileInfo) Suffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo6suffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QString bundleName() const

/*
Returns the name of the bundle.

On macOS and iOS this returns the proper localized name for a bundle if the path isBundle(). On all other platforms an empty QString is returned.

Example:


  QFileInfo fi("/Applications/Safari.app");
  QString bundle = fi.bundleName();                // name = "Safari"



This function was introduced in  Qt 4.3.

See also isBundle(), filePath(), baseName(), and suffix().
*/
func (this *QFileInfo) BundleName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10bundleNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QString completeSuffix() const

/*
Returns the complete suffix (extension) of the file.

The complete suffix consists of all characters in the file after (but not including) the first '.'.

Example:


  QFileInfo fi("/tmp/archive.tar.gz");
  QString ext = fi.completeSuffix();  // ext = "tar.gz"



See also fileName(), suffix(), baseName(), and completeBaseName().
*/
func (this *QFileInfo) CompleteSuffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo14completeSuffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path() const

/*
Returns the file's path. This doesn't include the file name.

Note that, if this QFileInfo object is given a path ending in a slash, the name of the file is considered empty and this function will return the entire path.

See also filePath(), absolutePath(), canonicalPath(), dir(), fileName(), and isRelative().
*/
func (this *QFileInfo) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absolutePath() const

/*
Returns a file's path absolute path. This doesn't include the file name.

On Unix the absolute path will always begin with the root, '/', directory. On Windows this will always begin 'D:/' where D is a drive letter, except for network shares that are not mapped to a drive letter, in which case the path will begin '//sharename/'.

In contrast to canonicalPath() symbolic links or redundant "." or ".." elements are not necessarily removed.

Warning: If filePath() is empty the behavior of this function is undefined.

See also absoluteFilePath(), path(), canonicalPath(), fileName(), and isRelative().
*/
func (this *QFileInfo) AbsolutePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo12absolutePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QString canonicalPath() const

/*
Returns the file's path canonical path (excluding the file name), i.e. an absolute path without symbolic links or redundant "." or ".." elements.

If the file does not exist, canonicalPath() returns an empty string.

See also path() and absolutePath().
*/
func (this *QFileInfo) CanonicalPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo13canonicalPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QDir dir() const

/*
Returns the path of the object's parent directory as a QDir object.

Note: The QDir returned always corresponds to the object's parent directory, even if the QFileInfo represents a directory.

For each of the following, dir() returns a QDir for "~/examples/191697".


      QFileInfo fileInfo1("~/examples/191697/.");
      QFileInfo fileInfo2("~/examples/191697/..");
      QFileInfo fileInfo3("~/examples/191697/main.cpp");



For each of the following, dir() returns a QDir for ".".


      QFileInfo fileInfo4(".");
      QFileInfo fileInfo5("..");
      QFileInfo fileInfo6("main.cpp");



See also absolutePath(), filePath(), fileName(), isRelative(), and absoluteDir().
*/
func (this *QFileInfo) Dir() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo3dirEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QDir absoluteDir() const

/*
Returns the file's absolute path as a QDir object.

See also dir(), filePath(), fileName(), and isRelative().
*/
func (this *QFileInfo) AbsoluteDir() *QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo11absoluteDirEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDir)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadable() const

/*
Returns true if the user can read the file; otherwise returns false.

Note: If the NTFS permissions check has not been enabled, the result on Windows will merely reflect whether the file exists.

See also isWritable(), isExecutable(), and permission().
*/
func (this *QFileInfo) IsReadable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10isReadableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWritable() const

/*
Returns true if the user can write to the file; otherwise returns false.

Note: If the NTFS permissions check has not been enabled, the result on Windows will merely reflect whether the file is marked as Read Only.

See also isReadable(), isExecutable(), and permission().
*/
func (this *QFileInfo) IsWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10isWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExecutable() const

/*
Returns true if the file is executable; otherwise returns false.

See also isReadable(), isWritable(), and permission().
*/
func (this *QFileInfo) IsExecutable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo12isExecutableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isHidden() const

/*
Returns true if this is a `hidden' file; otherwise returns false.

Note: This function returns true for the special entries "." and ".." on Unix, even though QDir::entryList threats them as shown.
*/
func (this *QFileInfo) IsHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8isHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNativePath() const

/*
Returns true if the file path can be used directly with native APIs. Returns false if the file is otherwise supported by a virtual file system inside Qt, such as the Qt Resource System.

Note: Native paths may still require conversion of path separators and character encoding, depending on platform and input requirements of the native API.

This function was introduced in  Qt 5.0.

See also QDir::toNativeSeparators(), QFile::encodeName(), filePath(), absoluteFilePath(), and canonicalFilePath().
*/
func (this *QFileInfo) IsNativePath() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo12isNativePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRelative() const

/*
Returns true if the file path name is relative, otherwise returns false if the path is absolute (e.g. under Unix a path is absolute if it begins with a "/").

See also isAbsolute().
*/
func (this *QFileInfo) IsRelative() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10isRelativeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAbsolute() const

/*
Returns true if the file path name is absolute, otherwise returns false if the path is relative.

See also isRelative().
*/
func (this *QFileInfo) IsAbsolute() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10isAbsoluteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool makeAbsolute()

/*
Converts the file's path to an absolute path if it is not already in that form. Returns true to indicate that the path was converted; otherwise returns false to indicate that the path was already absolute.

See also filePath() and isRelative().
*/
func (this *QFileInfo) MakeAbsolute() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo12makeAbsoluteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFile() const

/*
Returns true if this object points to a file or to a symbolic link to a file. Returns false if the object points to something which isn't a file, such as a directory.

See also isDir(), isSymLink(), and isBundle().
*/
func (this *QFileInfo) IsFile() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo6isFileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDir() const

/*
Returns true if this object points to a directory or to a symbolic link to a directory; otherwise returns false.

See also isFile(), isSymLink(), and isBundle().
*/
func (this *QFileInfo) IsDir() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo5isDirEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSymLink() const

/*
Returns true if this object points to a symbolic link (or to a shortcut on Windows); otherwise returns false.

On Unix (including macOS and iOS), opening a symlink effectively opens the link's target. On Windows, it opens the .lnk file itself.

Example:


  QFileInfo info(fileName);
  if (info.isSymLink())
      fileName = info.symLinkTarget();



Note: If the symlink points to a non existing file, exists() returns false.

See also isFile(), isDir(), and symLinkTarget().
*/
func (this *QFileInfo) IsSymLink() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo9isSymLinkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRoot() const

/*
Returns true if the object points to a directory or to a symbolic link to a directory, and that directory is the root directory; otherwise returns false.
*/
func (this *QFileInfo) IsRoot() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo6isRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBundle() const

/*
Returns true if this object points to a bundle or to a symbolic link to a bundle on macOS and iOS; otherwise returns false.

This function was introduced in  Qt 4.3.

See also isDir(), isSymLink(), and isFile().
*/
func (this *QFileInfo) IsBundle() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8isBundleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QString readLink() const

/*

 */
func (this *QFileInfo) ReadLink() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8readLinkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString symLinkTarget() const

/*
Returns the absolute path to the file or directory a symlink (or shortcut on Windows) points to, or a an empty string if the object isn't a symbolic link.

This name may not represent an existing file; it is only a string. QFileInfo::exists() returns true if the symlink points to an existing file.

This function was introduced in  Qt 4.2.

See also exists(), isSymLink(), isDir(), and isFile().
*/
func (this *QFileInfo) SymLinkTarget() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo13symLinkTargetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QString owner() const

/*
Returns the owner of the file. On systems where files do not have owners, or if an error occurs, an empty string is returned.

This function can be time consuming under Unix (in the order of milliseconds). On Windows, it will return an empty string unless the NTFS permissions check has been enabled.

See also ownerId(), group(), and groupId().
*/
func (this *QFileInfo) Owner() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo5ownerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:123
// index:0
// Public Visibility=Default Availability=Available
// [4] uint ownerId() const

/*
Returns the id of the owner of the file.

On Windows and on systems where files do not have owners this function returns ((uint) -2).

See also owner(), group(), and groupId().
*/
func (this *QFileInfo) OwnerId() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo7ownerIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qfileinfo.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QString group() const

/*
Returns the group of the file. On Windows, on systems where files do not have groups, or if an error occurs, an empty string is returned.

This function can be time consuming under Unix (in the order of milliseconds).

See also groupId(), owner(), and ownerId().
*/
func (this *QFileInfo) Group() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfileinfo.h:125
// index:0
// Public Visibility=Default Availability=Available
// [4] uint groupId() const

/*
Returns the id of the group the file belongs to.

On Windows and on systems where files do not have groups this function always returns (uint) -2.

See also group(), owner(), and ownerId().
*/
func (this *QFileInfo) GroupId() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo7groupIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qfileinfo.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool permission(QFile::Permissions) const

/*
Tests for file permissions. The permissions argument can be several flags of type QFile::Permissions OR-ed together to check for permission combinations.

On systems where files do not have permissions this function always returns true.

Note: The result might be inaccurate on Windows if the NTFS permissions check has not been enabled.

Example:


  QFileInfo fi("/tmp/archive.tar.gz");
  if (fi.permission(QFile::WriteUser | QFile::ReadGroup))
      qWarning("I can change the file; my group can read the file");
  if (fi.permission(QFile::WriteGroup | QFile::WriteOther))
      qWarning("The group or others can change the file");



See also isReadable(), isWritable(), and isExecutable().
*/
func (this *QFileInfo) Permission(permissions int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo10permissionE6QFlagsIN11QFileDevice10PermissionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), permissions)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] QFile::Permissions permissions() const

/*
Returns the complete OR-ed together combination of QFile::Permissions for the file.

Note: The result might be inaccurate on Windows if the NTFS permissions check has not been enabled.
*/
func (this *QFileInfo) Permissions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo11permissionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qfileinfo.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 size() const

/*
Returns the file size in bytes. If the file does not exist or cannot be fetched, 0 is returned.

See also exists().
*/
func (this *QFileInfo) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qfileinfo.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime created() const

/*
Returns the date and local time when the file was created.

On most Unix systems, this function returns the time of the last status change. A status change occurs when the file is created, but it also occurs whenever the user writes or sets inode information (for example, changing the file permissions).

If neither creation time nor "last status change" time are not available, returns the same as lastModified().

See also lastModified() and lastRead().
*/
func (this *QFileInfo) Created() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo7createdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime lastModified() const

/*
Returns the date and local time when the file was last modified.

See also created() and lastRead().
*/
func (this *QFileInfo) LastModified() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo12lastModifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime lastRead() const

/*
Returns the date and local time when the file was last read (accessed).

On platforms where this information is not available, returns the same as lastModified().

See also created() and lastModified().
*/
func (this *QFileInfo) LastRead() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo8lastReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qfileinfo.h:136
// index:0
// Public Visibility=Default Availability=Available
// [1] bool caching() const

/*
Returns true if caching is enabled; otherwise returns false.

See also setCaching() and refresh().
*/
func (this *QFileInfo) Caching() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QFileInfo7cachingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfileinfo.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaching(bool)

/*
If enable is true, enables caching of file information. If enable is false caching is disabled.

When caching is enabled, QFileInfo reads the file information from the file system the first time it's needed, but generally not later.

Caching is enabled by default.

See also refresh() and caching().
*/
func (this *QFileInfo) SetCaching(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QFileInfo10setCachingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
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
