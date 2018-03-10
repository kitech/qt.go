package qtcore

// /usr/include/qt/QtCore/qtemporaryfile.h
// #include <qtemporaryfile.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool open(QIODevice::OpenMode)
func (this *QTemporaryFile) InheritOpen(f func(flags int) bool) {
	qtrt.SetAllInheritCallback(this, "open", f)
}

/*

 */
type QTemporaryFile struct {
	*QFile
}
type QTemporaryFile_ITF interface {
	QFile_ITF
	QTemporaryFile_PTR() *QTemporaryFile
}

func (ptr *QTemporaryFile) QTemporaryFile_PTR() *QTemporaryFile { return ptr }

func (this *QTemporaryFile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFile.GetCthis()
	}
}
func (this *QTemporaryFile) SetCthis(cthis unsafe.Pointer) {
	this.QFile = NewQFileFromPointer(cthis)
}
func NewQTemporaryFileFromPointer(cthis unsafe.Pointer) *QTemporaryFile {
	bcthis0 := NewQFileFromPointer(cthis)
	return &QTemporaryFile{bcthis0}
}
func (*QTemporaryFile) NewFromPointer(cthis unsafe.Pointer) *QTemporaryFile {
	return NewQTemporaryFileFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTemporaryFile) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTemporaryFile10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtemporaryfile.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryFile()

/*
Constructs a QTemporaryFile using as file template the application name returned by QCoreApplication::applicationName() (otherwise qt_temp) followed by ".XXXXXX". The file is stored in the system's temporary directory, QDir::tempPath().

See also setFileTemplate() and QDir::tempPath().
*/
func NewQTemporaryFile() *QTemporaryFile {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFileC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTemporaryFile")
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:67
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryFile(const QString &)

/*
Constructs a QTemporaryFile using as file template the application name returned by QCoreApplication::applicationName() (otherwise qt_temp) followed by ".XXXXXX". The file is stored in the system's temporary directory, QDir::tempPath().

See also setFileTemplate() and QDir::tempPath().
*/
func NewQTemporaryFile_1(templateName string) *QTemporaryFile {
	var tmpArg0 = NewQString_5(templateName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFileC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTemporaryFile")
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryFile(QObject *)

/*
Constructs a QTemporaryFile using as file template the application name returned by QCoreApplication::applicationName() (otherwise qt_temp) followed by ".XXXXXX". The file is stored in the system's temporary directory, QDir::tempPath().

See also setFileTemplate() and QDir::tempPath().
*/
func NewQTemporaryFile_2(parent QObject_ITF /*777 QObject **/) *QTemporaryFile {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFileC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTemporaryFile")
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:70
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryFile(const QString &, QObject *)

/*
Constructs a QTemporaryFile using as file template the application name returned by QCoreApplication::applicationName() (otherwise qt_temp) followed by ".XXXXXX". The file is stored in the system's temporary directory, QDir::tempPath().

See also setFileTemplate() and QDir::tempPath().
*/
func NewQTemporaryFile_3(templateName string, parent QObject_ITF /*777 QObject **/) *QTemporaryFile {
	var tmpArg0 = NewQString_5(templateName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFileC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTemporaryFile")
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTemporaryFile()

/*

 */
func DeleteQTemporaryFile(this *QTemporaryFile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRemove() const

/*
Returns true if the QTemporaryFile is in auto remove mode. Auto-remove mode will automatically delete the filename from disk upon destruction. This makes it very easy to create your QTemporaryFile object on the stack, fill it with data, read from it, and finally on function return it will automatically clean up after itself.

Auto-remove is on by default.

See also setAutoRemove() and remove().
*/
func (this *QTemporaryFile) AutoRemove() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTemporaryFile10autoRemoveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporaryfile.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRemove(_Bool)

/*
Sets the QTemporaryFile into auto-remove mode if b is true.

Auto-remove is on by default.

If you set this property to false, ensure the application provides a way to remove the file once it is no longer needed, including passing the responsibility on to another process. Always use the fileName() function to obtain the name and never try to guess the name that QTemporaryFile has generated.

On some systems, if fileName() is not called before closing the file, the temporary file may be removed regardless of the state of this property. This behavior should not be relied upon, so application code should either call fileName() or leave the auto removal functionality enabled.

See also autoRemove() and remove().
*/
func (this *QTemporaryFile) SetAutoRemove(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile13setAutoRemoveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool open()

/*
A QTemporaryFile will always be opened in QIODevice::ReadWrite mode, this allows easy access to the data in the file. This function will return true upon success and will set the fileName() to the unique filename used.

See also fileName().
*/
func (this *QTemporaryFile) Open() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile4openEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporaryfile.h:98
// index:1
// Protected virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)

/*
A QTemporaryFile will always be opened in QIODevice::ReadWrite mode, this allows easy access to the data in the file. This function will return true upon success and will set the fileName() to the unique filename used.

See also fileName().
*/
func (this *QTemporaryFile) Open_1(flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporaryfile.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString fileName() const

/*
Reimplemented from QFileDevice::fileName().

Returns the complete unique filename backing the QTemporaryFile object. This string is null before the QTemporaryFile is opened, afterwards it will contain the fileTemplate() plus additional characters to make it unique.

See also fileTemplate().
*/
func (this *QTemporaryFile) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTemporaryFile8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtemporaryfile.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileTemplate() const

/*
Returns the set file template. The default file template will be called qcoreappname.XXXXXX and be placed in QDir::tempPath().

See also setFileTemplate().
*/
func (this *QTemporaryFile) FileTemplate() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTemporaryFile12fileTemplateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtemporaryfile.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileTemplate(const QString &)

/*
Sets the static portion of the file name to name. If the file template contains XXXXXX that will automatically be replaced with the unique part of the filename, otherwise a filename will be determined automatically based on the static portion specified.

If name contains a relative file path, the path will be relative to the current working directory. You can use QDir::tempPath() to construct name if you want use the system's temporary directory.

See also fileTemplate().
*/
func (this *QTemporaryFile) SetFileTemplate(name string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile15setFileTemplateERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rename(const QString &)

/*

 */
func (this *QTemporaryFile) Rename(newName string) bool {
	var tmpArg0 = NewQString_5(newName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile6renameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporaryfile.h:88
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QTemporaryFile * createLocalFile(const QString &)

/*

 */
func (this *QTemporaryFile) CreateLocalFile(fileName string) *QTemporaryFile /*777 QTemporaryFile **/ {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile15createLocalFileERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTemporaryFile_CreateLocalFile(fileName string) *QTemporaryFile /*777 QTemporaryFile **/ {
	var nilthis *QTemporaryFile
	rv := nilthis.CreateLocalFile(fileName)
	return rv
}

// /usr/include/qt/QtCore/qtemporaryfile.h:90
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QTemporaryFile * createLocalFile(QFile &)

/*

 */
func (this *QTemporaryFile) CreateLocalFile_1(file QFile_ITF) *QTemporaryFile /*777 QTemporaryFile **/ {
	var convArg0 unsafe.Pointer
	if file != nil && file.QFile_PTR() != nil {
		convArg0 = file.QFile_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile15createLocalFileER5QFile", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTemporaryFile_CreateLocalFile_1(file QFile_ITF) *QTemporaryFile /*777 QTemporaryFile **/ {
	var nilthis *QTemporaryFile
	rv := nilthis.CreateLocalFile_1(file)
	return rv
}

// /usr/include/qt/QtCore/qtemporaryfile.h:93
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QTemporaryFile * createNativeFile(const QString &)

/*
If file is not already a native file, then a QTemporaryFile is created in QDir::tempPath(), the contents of file is copied into it, and a pointer to the temporary file is returned. Does nothing and returns 0 if file is already a native file.

For example:


  QFile f(":/resources/file.txt");
  QTemporaryFile::createNativeFile(f); // Returns a pointer to a temporary file

  QFile f("/users/qt/file.txt");
  QTemporaryFile::createNativeFile(f); // Returns 0



See also QFileInfo::isNativePath().
*/
func (this *QTemporaryFile) CreateNativeFile(fileName string) *QTemporaryFile /*777 QTemporaryFile **/ {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile16createNativeFileERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTemporaryFile_CreateNativeFile(fileName string) *QTemporaryFile /*777 QTemporaryFile **/ {
	var nilthis *QTemporaryFile
	rv := nilthis.CreateNativeFile(fileName)
	return rv
}

// /usr/include/qt/QtCore/qtemporaryfile.h:95
// index:1
// Public static Visibility=Default Availability=Available
// [8] QTemporaryFile * createNativeFile(QFile &)

/*
If file is not already a native file, then a QTemporaryFile is created in QDir::tempPath(), the contents of file is copied into it, and a pointer to the temporary file is returned. Does nothing and returns 0 if file is already a native file.

For example:


  QFile f(":/resources/file.txt");
  QTemporaryFile::createNativeFile(f); // Returns a pointer to a temporary file

  QFile f("/users/qt/file.txt");
  QTemporaryFile::createNativeFile(f); // Returns 0



See also QFileInfo::isNativePath().
*/
func (this *QTemporaryFile) CreateNativeFile_1(file QFile_ITF) *QTemporaryFile /*777 QTemporaryFile **/ {
	var convArg0 unsafe.Pointer
	if file != nil && file.QFile_PTR() != nil {
		convArg0 = file.QFile_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile16createNativeFileER5QFile", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTemporaryFile_CreateNativeFile_1(file QFile_ITF) *QTemporaryFile /*777 QTemporaryFile **/ {
	var nilthis *QTemporaryFile
	rv := nilthis.CreateNativeFile_1(file)
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
