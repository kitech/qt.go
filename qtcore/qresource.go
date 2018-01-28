package qtcore

// /usr/include/qt/QtCore/qresource.h
// #include <qresource.h>
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
type QResource struct {
	*qtrt.CObject
}

func (this *QResource) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QResource) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQResourceFromPointer(cthis unsafe.Pointer) *QResource {
	return &QResource{&qtrt.CObject{cthis}}
}
func (*QResource) NewFromPointer(cthis unsafe.Pointer) *QResource {
	return NewQResourceFromPointer(cthis)
}

// /usr/include/qt/QtCore/qresource.h:57
// index:0
// Public
// void QResource(const QString &, const QLocale &)
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
// void setFileName(const QString &)
func (this *QResource) SetFileName(file *QString) {
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:61
// index:0
// Public
// QString fileName()
func (this *QResource) FileName() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource8fileNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qresource.h:62
// index:0
// Public
// QString absoluteFilePath()
func (this *QResource) AbsoluteFilePath() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource16absoluteFilePathEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qresource.h:64
// index:0
// Public
// void setLocale(const QLocale &)
func (this *QResource) SetLocale(locale *QLocale) {
	var convArg0 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource9setLocaleERK7QLocale", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:65
// index:0
// Public
// QLocale locale()
func (this *QResource) Locale() *QLocale /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource6localeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qresource.h:67
// index:0
// Public
// bool isValid()
func (this *QResource) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qresource.h:69
// index:0
// Public
// bool isCompressed()
func (this *QResource) IsCompressed() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource12isCompressedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qresource.h:70
// index:0
// Public
// qint64 size()
func (this *QResource) Size() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qresource.h:71
// index:0
// Public
// const uchar * data()
func (this *QResource) Data() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qresource.h:72
// index:0
// Public
// QDateTime lastModified()
func (this *QResource) LastModified() *QDateTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource12lastModifiedEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qresource.h:74
// index:0
// Public static
// void addSearchPath(const QString &)
func (this *QResource) AddSearchPath(path *QString) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource13addSearchPathERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QResource_AddSearchPath(path *QString) {
	var nilthis *QResource
	nilthis.AddSearchPath(path)
}

// /usr/include/qt/QtCore/qresource.h:77
// index:0
// Public static
// bool registerResource(const QString &, const QString &)
func (this *QResource) RegisterResource(rccFilename *QString, resourceRoot *QString) bool {
	var convArg0 = rccFilename.GetCthis()
	var convArg1 = resourceRoot.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource16registerResourceERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QResource_RegisterResource(rccFilename *QString, resourceRoot *QString) bool {
	var nilthis *QResource
	rv := nilthis.RegisterResource(rccFilename, resourceRoot)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:80
// index:1
// Public static
// bool registerResource(const uchar *, const QString &)
func (this *QResource) RegisterResource_1(rccData unsafe.Pointer /*666*/, resourceRoot *QString) bool {
	var convArg1 = resourceRoot.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource16registerResourceEPKhRK7QString", ffiqt.FFI_TYPE_POINTER, &rccData, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QResource_RegisterResource_1(rccData unsafe.Pointer /*666*/, resourceRoot *QString) bool {
	var nilthis *QResource
	rv := nilthis.RegisterResource_1(rccData, resourceRoot)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:78
// index:0
// Public static
// bool unregisterResource(const QString &, const QString &)
func (this *QResource) UnregisterResource(rccFilename *QString, resourceRoot *QString) bool {
	var convArg0 = rccFilename.GetCthis()
	var convArg1 = resourceRoot.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource18unregisterResourceERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QResource_UnregisterResource(rccFilename *QString, resourceRoot *QString) bool {
	var nilthis *QResource
	rv := nilthis.UnregisterResource(rccFilename, resourceRoot)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:81
// index:1
// Public static
// bool unregisterResource(const uchar *, const QString &)
func (this *QResource) UnregisterResource_1(rccData unsafe.Pointer /*666*/, resourceRoot *QString) bool {
	var convArg1 = resourceRoot.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QResource18unregisterResourceEPKhRK7QString", ffiqt.FFI_TYPE_POINTER, &rccData, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QResource_UnregisterResource_1(rccData unsafe.Pointer /*666*/, resourceRoot *QString) bool {
	var nilthis *QResource
	rv := nilthis.UnregisterResource_1(rccData, resourceRoot)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:86
// index:0
// Protected
// bool isDir()
func (this *QResource) IsDir() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource5isDirEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qresource.h:87
// index:0
// Protected inline
// bool isFile()
func (this *QResource) IsFile() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QResource6isFileEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
