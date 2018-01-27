package qtcore

// /usr/include/qt/QtCore/qtemporaryfile.h
// #include <qtemporaryfile.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
type QTemporaryFile struct {
	*QFile
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QTemporaryFile) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTemporaryFile10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qtemporaryfile.h:66
// index:0
// Public
// void QTemporaryFile()
func NewQTemporaryFile() *QTemporaryFile {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFileC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:67
// index:1
// Public
// void QTemporaryFile(const QString &)
func NewQTemporaryFile_1(templateName *QString) *QTemporaryFile {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = templateName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFileC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:69
// index:2
// Public
// void QTemporaryFile(QObject *)
func NewQTemporaryFile_2(parent *QObject /*777 QObject **/) *QTemporaryFile {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFileC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:70
// index:3
// Public
// void QTemporaryFile(const QString &, QObject *)
func NewQTemporaryFile_3(templateName *QString, parent *QObject /*777 QObject **/) *QTemporaryFile {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = templateName.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFileC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:72
// index:0
// Public virtual
// void ~QTemporaryFile()
func DeleteQTemporaryFile(*QTemporaryFile) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFileD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:74
// index:0
// Public
// bool autoRemove()
func (this *QTemporaryFile) AutoRemove() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTemporaryFile10autoRemoveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporaryfile.h:75
// index:0
// Public
// void setAutoRemove(bool)
func (this *QTemporaryFile) SetAutoRemove(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile13setAutoRemoveEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:78
// index:0
// Public inline
// bool open()
func (this *QTemporaryFile) Open() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile4openEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporaryfile.h:80
// index:0
// Public virtual
// QString fileName()
func (this *QTemporaryFile) FileName() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTemporaryFile8fileNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtemporaryfile.h:81
// index:0
// Public
// QString fileTemplate()
func (this *QTemporaryFile) FileTemplate() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTemporaryFile12fileTemplateEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtemporaryfile.h:82
// index:0
// Public
// void setFileTemplate(const QString &)
func (this *QTemporaryFile) SetFileTemplate(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile15setFileTemplateERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:85
// index:0
// Public
// bool rename(const QString &)
func (this *QTemporaryFile) Rename(newName *QString) bool {
	var convArg0 = newName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile6renameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtemporaryfile.h:88
// index:0
// Public static inline
// QTemporaryFile * createLocalFile(const QString &)
func (this *QTemporaryFile) CreateLocalFile(fileName *QString) *QTemporaryFile /*777 QTemporaryFile **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile15createLocalFileERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTemporaryFile_CreateLocalFile(fileName *QString) *QTemporaryFile /*777 QTemporaryFile **/ {
	var nilthis *QTemporaryFile
	rv := nilthis.CreateLocalFile(fileName)
	return rv
}

// /usr/include/qt/QtCore/qtemporaryfile.h:90
// index:1
// Public static inline
// QTemporaryFile * createLocalFile(QFile &)
func (this *QTemporaryFile) CreateLocalFile_1(file *QFile) *QTemporaryFile /*777 QTemporaryFile **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile15createLocalFileER5QFile", ffiqt.FFI_TYPE_POINTER, file)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTemporaryFile_CreateLocalFile_1(file *QFile) *QTemporaryFile /*777 QTemporaryFile **/ {
	var nilthis *QTemporaryFile
	rv := nilthis.CreateLocalFile_1(file)
	return rv
}

// /usr/include/qt/QtCore/qtemporaryfile.h:93
// index:0
// Public static inline
// QTemporaryFile * createNativeFile(const QString &)
func (this *QTemporaryFile) CreateNativeFile(fileName *QString) *QTemporaryFile /*777 QTemporaryFile **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile16createNativeFileERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTemporaryFile_CreateNativeFile(fileName *QString) *QTemporaryFile /*777 QTemporaryFile **/ {
	var nilthis *QTemporaryFile
	rv := nilthis.CreateNativeFile(fileName)
	return rv
}

// /usr/include/qt/QtCore/qtemporaryfile.h:95
// index:1
// Public static
// QTemporaryFile * createNativeFile(QFile &)
func (this *QTemporaryFile) CreateNativeFile_1(file *QFile) *QTemporaryFile /*777 QTemporaryFile **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile16createNativeFileER5QFile", ffiqt.FFI_TYPE_POINTER, file)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTemporaryFileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QTemporaryFile_CreateNativeFile_1(file *QFile) *QTemporaryFile /*777 QTemporaryFile **/ {
	var nilthis *QTemporaryFile
	rv := nilthis.CreateNativeFile_1(file)
	return rv
}

//  body block end
