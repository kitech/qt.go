//  header block begin
// /usr/include/qt/QtCore/qresource.h
// #include <qresource.h>
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
type QResource struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qresource.h:57
// index:0
// void QResource(const class QString &, const class QLocale &)
func NewQResource(file unsafe.Pointer, locale unsafe.Pointer) *QResource {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResourceC2ERK7QStringRK7QLocale", ffiqt.FFI_TYPE_VOID, cthis, file, locale)
	gopp.ErrPrint(err, rv)
	return &QResource{cthis}
}

// /usr/include/qt/QtCore/qresource.h:58
// index:0
// void ~QResource()
func DeleteQResource(*QResource) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResourceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:60
// index:0
// void setFileName(const class QString &)
func (this *QResource) SetFileName(file unsafe.Pointer) {
	// 0: (, const QString & file), (file)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource11setFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, file)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:61
// index:0
// QString fileName()
func (this *QResource) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource8fileNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:62
// index:0
// QString absoluteFilePath()
func (this *QResource) AbsoluteFilePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource16absoluteFilePathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:64
// index:0
// void setLocale(const class QLocale &)
func (this *QResource) SetLocale(locale unsafe.Pointer) {
	// 0: (, const QLocale & locale), (locale)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource9setLocaleERK7QLocale", ffiqt.FFI_TYPE_VOID, this.cthis, locale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:65
// index:0
// QLocale locale()
func (this *QResource) Locale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource6localeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:67
// index:0
// bool isValid()
func (this *QResource) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:69
// index:0
// bool isCompressed()
func (this *QResource) IsCompressed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource12isCompressedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:70
// index:0
// qint64 size()
func (this *QResource) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:71
// index:0
// const uchar * data()
func (this *QResource) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource4dataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:72
// index:0
// QDateTime lastModified()
func (this *QResource) LastModified() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource12lastModifiedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:74
// index:0
// static
// void addSearchPath(const class QString &)
func (this *QResource) AddSearchPath(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource13addSearchPathERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QResource_AddSearchPath(path unsafe.Pointer) {
	// 0: (const QString & path), (path)
	var nilthis *QResource
	nilthis.AddSearchPath(path)
}

// /usr/include/qt/QtCore/qresource.h:75
// index:0
// static
// QStringList searchPaths()
func (this *QResource) SearchPaths() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource11searchPathsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QResource_SearchPaths() {
	// 0: (), ()
	var nilthis *QResource
	nilthis.SearchPaths()
}

// /usr/include/qt/QtCore/qresource.h:77
// index:0
// static
// bool registerResource(const class QString &, const class QString &)
func (this *QResource) RegisterResource(rccFilename unsafe.Pointer, resourceRoot unsafe.Pointer) {
	// 0: (const QString & rccFilename, const QString & resourceRoot), (rccFilename, resourceRoot)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource16registerResourceERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QResource_RegisterResource(rccFilename unsafe.Pointer, resourceRoot unsafe.Pointer) {
	// 0: (const QString & rccFilename, const QString & resourceRoot), (rccFilename, resourceRoot)
	var nilthis *QResource
	nilthis.RegisterResource(rccFilename, resourceRoot)
}

// /usr/include/qt/QtCore/qresource.h:80
// index:1
// static
// bool registerResource(const uchar *, const class QString &)
func (this *QResource) RegisterResource_1(rccData unsafe.Pointer, resourceRoot unsafe.Pointer) {
	// 1: (const uchar * rccData, const QString & resourceRoot), (rccData, resourceRoot)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource16registerResourceEPKhRK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QResource_RegisterResource_1(rccData unsafe.Pointer, resourceRoot unsafe.Pointer) {
	// 1: (const uchar * rccData, const QString & resourceRoot), (rccData, resourceRoot)
	var nilthis *QResource
	nilthis.RegisterResource_1(rccData, resourceRoot)
}

// /usr/include/qt/QtCore/qresource.h:78
// index:0
// static
// bool unregisterResource(const class QString &, const class QString &)
func (this *QResource) UnregisterResource(rccFilename unsafe.Pointer, resourceRoot unsafe.Pointer) {
	// 0: (const QString & rccFilename, const QString & resourceRoot), (rccFilename, resourceRoot)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource18unregisterResourceERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QResource_UnregisterResource(rccFilename unsafe.Pointer, resourceRoot unsafe.Pointer) {
	// 0: (const QString & rccFilename, const QString & resourceRoot), (rccFilename, resourceRoot)
	var nilthis *QResource
	nilthis.UnregisterResource(rccFilename, resourceRoot)
}

// /usr/include/qt/QtCore/qresource.h:81
// index:1
// static
// bool unregisterResource(const uchar *, const class QString &)
func (this *QResource) UnregisterResource_1(rccData unsafe.Pointer, resourceRoot unsafe.Pointer) {
	// 1: (const uchar * rccData, const QString & resourceRoot), (rccData, resourceRoot)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource18unregisterResourceEPKhRK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QResource_UnregisterResource_1(rccData unsafe.Pointer, resourceRoot unsafe.Pointer) {
	// 1: (const uchar * rccData, const QString & resourceRoot), (rccData, resourceRoot)
	var nilthis *QResource
	nilthis.UnregisterResource_1(rccData, resourceRoot)
}

//  body block end
