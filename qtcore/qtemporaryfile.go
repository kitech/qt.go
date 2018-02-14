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
// [8] const QMetaObject * metaObject()
func (this *QTemporaryFile) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTemporaryFile10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtemporaryfile.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTemporaryFile()
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
func DeleteQTemporaryFile(this *QTemporaryFile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRemove()
func (this *QTemporaryFile) AutoRemove() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTemporaryFile10autoRemoveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporaryfile.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRemove(_Bool)
func (this *QTemporaryFile) SetAutoRemove(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile13setAutoRemoveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool open()
func (this *QTemporaryFile) Open() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile4openEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporaryfile.h:98
// index:1
// Protected virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)
func (this *QTemporaryFile) Open_1(flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTemporaryFile4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporaryfile.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString fileName()
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
// [8] QString fileTemplate()
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
