package qtcore

// /usr/include/qt/QtCore/qtemporarydir.h
// #include <qtemporarydir.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QTemporaryDir struct {
	*qtrt.CObject
}
type QTemporaryDir_ITF interface {
	QTemporaryDir_PTR() *QTemporaryDir
}

func (ptr *QTemporaryDir) QTemporaryDir_PTR() *QTemporaryDir { return ptr }

func (this *QTemporaryDir) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTemporaryDir) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTemporaryDirFromPointer(cthis unsafe.Pointer) *QTemporaryDir {
	return &QTemporaryDir{&qtrt.CObject{cthis}}
}
func (*QTemporaryDir) NewFromPointer(cthis unsafe.Pointer) *QTemporaryDir {
	return NewQTemporaryDirFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtemporarydir.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryDir()

/*
Constructs a QTemporaryDir using as template the application name returned by QCoreApplication::applicationName() (otherwise qt_temp). The directory is stored in the system's temporary directory, QDir::tempPath().

See also QDir::tempPath().
*/
func (*QTemporaryDir) NewForInherit() *QTemporaryDir {
	return NewQTemporaryDir()
}
func NewQTemporaryDir() *QTemporaryDir {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDirC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTemporaryDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTemporaryDir)
	return gothis
}

// /usr/include/qt/QtCore/qtemporarydir.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryDir(const QString &)

/*
Constructs a QTemporaryDir using as template the application name returned by QCoreApplication::applicationName() (otherwise qt_temp). The directory is stored in the system's temporary directory, QDir::tempPath().

See also QDir::tempPath().
*/
func (*QTemporaryDir) NewForInherit1(templateName string) *QTemporaryDir {
	return NewQTemporaryDir1(templateName)
}
func NewQTemporaryDir1(templateName string) *QTemporaryDir {
	var tmpArg0 = NewQString5(templateName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDirC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTemporaryDirFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTemporaryDir)
	return gothis
}

// /usr/include/qt/QtCore/qtemporarydir.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTemporaryDir()

/*

 */
func DeleteQTemporaryDir(this *QTemporaryDir) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDirD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtemporarydir.h:60
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the QTemporaryDir was created successfully.
*/
func (this *QTemporaryDir) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporarydir.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
If isValid() returns false, this function returns the error string that explains why the creation of the temporary directory failed. Otherwise, this function return an empty string.

This function was introduced in  Qt 5.6.
*/
func (this *QTemporaryDir) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtemporarydir.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRemove() const

/*
Returns true if the QTemporaryDir is in auto remove mode. Auto-remove mode will automatically delete the directory from disk upon destruction. This makes it very easy to create your QTemporaryDir object on the stack, fill it with files, do something with the files, and finally on function return it will automatically clean up after itself.

Auto-remove is on by default.

See also setAutoRemove() and remove().
*/
func (this *QTemporaryDir) AutoRemove() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir10autoRemoveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporarydir.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRemove(bool)

/*
Sets the QTemporaryDir into auto-remove mode if b is true.

Auto-remove is on by default.

See also autoRemove() and remove().
*/
func (this *QTemporaryDir) SetAutoRemove(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDir13setAutoRemoveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporarydir.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove()

/*
Removes the temporary directory, including all its contents.

Returns true if removing was successful.
*/
func (this *QTemporaryDir) Remove() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTemporaryDir6removeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporarydir.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString path() const

/*
Returns the path to the temporary directory. Empty if the QTemporaryDir could not be created.
*/
func (this *QTemporaryDir) Path() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtemporarydir.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath(const QString &) const

/*
Returns the path name of a file in the temporary directory. Does not check if the file actually exists in the directory. Redundant multiple separators or "." and ".." directories in fileName are not removed (see QDir::cleanPath()). Absolute paths are not allowed.

This function was introduced in  Qt 5.9.
*/
func (this *QTemporaryDir) FilePath(fileName string) string {
	var tmpArg0 = NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTemporaryDir8filePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
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
