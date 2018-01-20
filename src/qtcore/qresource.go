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
	*qtrt.CObject
}

func (this *QResource) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQResourceFromPointer(cthis unsafe.Pointer) *QResource {
	return &QResource{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qresource.h:57
// index:0
// Public
// void QResource(const class QString &, const class QLocale &)
func NewQResource(file *QString, locale *QLocale) *QResource {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = file.GetCthis()
	var convArg1 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResourceC2ERK7QStringRK7QLocale", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQResourceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qresource.h:58
// index:0
// Public
// void ~QResource()
func DeleteQResource(*QResource) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResourceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:60
// index:0
// Public
// void setFileName(const class QString &)
func (this *QResource) SetFileName(file *QString) {
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:61
// index:0
// Public
// QString fileName()
func (this *QResource) FileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:62
// index:0
// Public
// QString absoluteFilePath()
func (this *QResource) AbsoluteFilePath() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource16absoluteFilePathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:64
// index:0
// Public
// void setLocale(const class QLocale &)
func (this *QResource) SetLocale(locale *QLocale) {
	var convArg0 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource9setLocaleERK7QLocale", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:65
// index:0
// Public
// QLocale locale()
func (this *QResource) Locale() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource6localeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:67
// index:0
// Public
// bool isValid()
func (this *QResource) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:69
// index:0
// Public
// bool isCompressed()
func (this *QResource) IsCompressed() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource12isCompressedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:70
// index:0
// Public
// qint64 size()
func (this *QResource) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:71
// index:0
// Public
// const uchar * data()
func (this *QResource) Data() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:72
// index:0
// Public
// QDateTime lastModified()
func (this *QResource) LastModified() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource12lastModifiedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:74
// index:0
// Public static
// void addSearchPath(const class QString &)
func (this *QResource) AddSearchPath(path *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource13addSearchPathERK7QString", ffiqt.FFI_TYPE_POINTER, path)
	gopp.ErrPrint(err, rv)
}
func QResource_AddSearchPath(path *QString) {
	var nilthis *QResource
	nilthis.AddSearchPath(path)
}

// /usr/include/qt/QtCore/qresource.h:75
// index:0
// Public static
// QStringList searchPaths()
func (this *QResource) SearchPaths() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource11searchPathsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QResource_SearchPaths() {
	var nilthis *QResource
	nilthis.SearchPaths()
}

// /usr/include/qt/QtCore/qresource.h:77
// index:0
// Public static
// bool registerResource(const class QString &, const class QString &)
func (this *QResource) RegisterResource(rccFilename *QString, resourceRoot *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource16registerResourceERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, rccFilename, resourceRoot)
	gopp.ErrPrint(err, rv)
	return rv
}
func QResource_RegisterResource(rccFilename *QString, resourceRoot *QString) {
	var nilthis *QResource
	nilthis.RegisterResource(rccFilename, resourceRoot)
}

// /usr/include/qt/QtCore/qresource.h:80
// index:1
// Public static
// bool registerResource(const uchar *, const class QString &)
func (this *QResource) RegisterResource_1(rccData unsafe.Pointer, resourceRoot *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource16registerResourceEPKhRK7QString", ffiqt.FFI_TYPE_POINTER, rccData, resourceRoot)
	gopp.ErrPrint(err, rv)
	return rv
}
func QResource_RegisterResource_1(rccData unsafe.Pointer, resourceRoot *QString) {
	var nilthis *QResource
	nilthis.RegisterResource_1(rccData, resourceRoot)
}

// /usr/include/qt/QtCore/qresource.h:78
// index:0
// Public static
// bool unregisterResource(const class QString &, const class QString &)
func (this *QResource) UnregisterResource(rccFilename *QString, resourceRoot *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource18unregisterResourceERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, rccFilename, resourceRoot)
	gopp.ErrPrint(err, rv)
	return rv
}
func QResource_UnregisterResource(rccFilename *QString, resourceRoot *QString) {
	var nilthis *QResource
	nilthis.UnregisterResource(rccFilename, resourceRoot)
}

// /usr/include/qt/QtCore/qresource.h:81
// index:1
// Public static
// bool unregisterResource(const uchar *, const class QString &)
func (this *QResource) UnregisterResource_1(rccData unsafe.Pointer, resourceRoot *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource18unregisterResourceEPKhRK7QString", ffiqt.FFI_TYPE_POINTER, rccData, resourceRoot)
	gopp.ErrPrint(err, rv)
	return rv
}
func QResource_UnregisterResource_1(rccData unsafe.Pointer, resourceRoot *QString) {
	var nilthis *QResource
	nilthis.UnregisterResource_1(rccData, resourceRoot)
}

// /usr/include/qt/QtCore/qresource.h:86
// index:0
// Protected
// bool isDir()
func (this *QResource) IsDir() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource5isDirEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:87
// index:0
// Protected inline
// bool isFile()
func (this *QResource) IsFile() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource6isFileEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:88
// index:0
// Protected
// QStringList children()
func (this *QResource) Children() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource8childrenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
