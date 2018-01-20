//  header block begin
// /usr/include/qt/QtCore/qtemporaryfile.h
// #include <qtemporaryfile.h>
// #include <QtCore>
package qtcore

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
type QTemporaryFile struct {
	*QFile
}

func (this *QTemporaryFile) GetCthis() unsafe.Pointer {
	return this.QFile.GetCthis()
}

// /usr/include/qt/QtCore/qtemporaryfile.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTemporaryFile) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTemporaryFile10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:66
// index:0
// void QTemporaryFile()
func NewQTemporaryFile() *QTemporaryFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFileC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(cthis)
	return gothis
}
func NewQTemporaryFileFromPointer(cthis unsafe.Pointer) *QTemporaryFile {
	bcthis0 := NewQFileFromPointer(cthis)
	return &QTemporaryFile{bcthis0}
}

// /usr/include/qt/QtCore/qtemporaryfile.h:67
// index:1
// void QTemporaryFile(const class QString &)
func NewQTemporaryFile_1(templateName unsafe.Pointer) *QTemporaryFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFileC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, templateName)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:69
// index:2
// void QTemporaryFile(class QObject *)
func NewQTemporaryFile_2(parent unsafe.Pointer) *QTemporaryFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFileC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:70
// index:3
// void QTemporaryFile(const class QString &, class QObject *)
func NewQTemporaryFile_3(templateName unsafe.Pointer, parent unsafe.Pointer) *QTemporaryFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFileC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, templateName, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTemporaryFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtemporaryfile.h:72
// index:0
// virtual
// void ~QTemporaryFile()
func DeleteQTemporaryFile(*QTemporaryFile) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFileD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:74
// index:0
// bool autoRemove()
func (this *QTemporaryFile) AutoRemove() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTemporaryFile10autoRemoveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:75
// index:0
// void setAutoRemove(_Bool)
func (this *QTemporaryFile) SetAutoRemove(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile13setAutoRemoveEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:78
// index:0
// inline
// bool open()
func (this *QTemporaryFile) Open() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile4openEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:98
// index:1
// virtual
// bool open(QIODevice::OpenMode)
func (this *QTemporaryFile) Open_1(flags int) {
	// 1: (, QFlags<QIODevice::OpenModeFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile4openE6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:80
// index:0
// virtual
// QString fileName()
func (this *QTemporaryFile) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTemporaryFile8fileNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:81
// index:0
// QString fileTemplate()
func (this *QTemporaryFile) FileTemplate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTemporaryFile12fileTemplateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:82
// index:0
// void setFileTemplate(const class QString &)
func (this *QTemporaryFile) SetFileTemplate(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile15setFileTemplateERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:85
// index:0
// bool rename(const class QString &)
func (this *QTemporaryFile) Rename(newName unsafe.Pointer) {
	// 0: (, newName const QString &), (newName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile6renameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), newName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:88
// index:0
// static inline
// QTemporaryFile * createLocalFile(const class QString &)
func (this *QTemporaryFile) CreateLocalFile(fileName unsafe.Pointer) {
	// 0: (fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile15createLocalFileERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTemporaryFile_CreateLocalFile(fileName unsafe.Pointer) {
	// 0: (fileName const QString &), (fileName)
	var nilthis *QTemporaryFile
	nilthis.CreateLocalFile(fileName)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:90
// index:1
// static inline
// QTemporaryFile * createLocalFile(class QFile &)
func (this *QTemporaryFile) CreateLocalFile_1(file unsafe.Pointer) {
	// 1: (file QFile &), (file)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile15createLocalFileER5QFile", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTemporaryFile_CreateLocalFile_1(file unsafe.Pointer) {
	// 1: (file QFile &), (file)
	var nilthis *QTemporaryFile
	nilthis.CreateLocalFile_1(file)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:93
// index:0
// static inline
// QTemporaryFile * createNativeFile(const class QString &)
func (this *QTemporaryFile) CreateNativeFile(fileName unsafe.Pointer) {
	// 0: (fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile16createNativeFileERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTemporaryFile_CreateNativeFile(fileName unsafe.Pointer) {
	// 0: (fileName const QString &), (fileName)
	var nilthis *QTemporaryFile
	nilthis.CreateNativeFile(fileName)
}

// /usr/include/qt/QtCore/qtemporaryfile.h:95
// index:1
// static
// QTemporaryFile * createNativeFile(class QFile &)
func (this *QTemporaryFile) CreateNativeFile_1(file unsafe.Pointer) {
	// 1: (file QFile &), (file)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTemporaryFile16createNativeFileER5QFile", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTemporaryFile_CreateNativeFile_1(file unsafe.Pointer) {
	// 1: (file QFile &), (file)
	var nilthis *QTemporaryFile
	nilthis.CreateNativeFile_1(file)
}

//  body block end
