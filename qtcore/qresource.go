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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool isDir()
func (this *QResource) InheritIsDir(f func() bool) {
	qtrt.SetAllInheritCallback(this, "isDir", f)
}

// bool isFile()
func (this *QResource) InheritIsFile(f func() bool) {
	qtrt.SetAllInheritCallback(this, "isFile", f)
}

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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQResourceFromPointer(cthis unsafe.Pointer) *QResource {
	return &QResource{&qtrt.CObject{cthis}}
}
func (*QResource) NewFromPointer(cthis unsafe.Pointer) *QResource {
	return NewQResourceFromPointer(cthis)
}

// /usr/include/qt/QtCore/qresource.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QResource(const QString &, const QLocale &)
func NewQResource(file string, locale *QLocale) *QResource {
	var tmpArg0 = NewQString_5(file)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = locale.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QResourceC2ERK7QStringRK7QLocale", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQResourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQResource)
	return gothis
}

// /usr/include/qt/QtCore/qresource.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QResource()
func DeleteQResource(this *QResource) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QResourceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qresource.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QResource) SetFileName(file string) {
	var tmpArg0 = NewQString_5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QResource11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QResource) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QResource8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qresource.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QString absoluteFilePath()
func (this *QResource) AbsoluteFilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QResource16absoluteFilePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qresource.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocale(const QLocale &)
func (this *QResource) SetLocale(locale *QLocale) {
	var convArg0 = locale.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QResource9setLocaleERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qresource.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale()
func (this *QResource) Locale() *QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QResource6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtCore/qresource.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QResource) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QResource7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qresource.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCompressed()
func (this *QResource) IsCompressed() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QResource12isCompressedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qresource.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 size()
func (this *QResource) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QResource4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qresource.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] const uchar * data()
func (this *QResource) Data() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QResource4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qresource.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime lastModified()
func (this *QResource) LastModified() *QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QResource12lastModifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtCore/qresource.h:74
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addSearchPath(const QString &)
func (this *QResource) AddSearchPath(path string) {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QResource13addSearchPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QResource_AddSearchPath(path string) {
	var nilthis *QResource
	nilthis.AddSearchPath(path)
}

// /usr/include/qt/QtCore/qresource.h:77
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool registerResource(const QString &, const QString &)
func (this *QResource) RegisterResource(rccFilename string, resourceRoot string) bool {
	var tmpArg0 = NewQString_5(rccFilename)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(resourceRoot)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QResource16registerResourceERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QResource_RegisterResource(rccFilename string, resourceRoot string) bool {
	var nilthis *QResource
	rv := nilthis.RegisterResource(rccFilename, resourceRoot)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:80
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool registerResource(const uchar *, const QString &)
func (this *QResource) RegisterResource_1(rccData unsafe.Pointer /*666*/, resourceRoot string) bool {
	var tmpArg1 = NewQString_5(resourceRoot)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QResource16registerResourceEPKhRK7QString", qtrt.FFI_TYPE_POINTER, &rccData, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QResource_RegisterResource_1(rccData unsafe.Pointer /*666*/, resourceRoot string) bool {
	var nilthis *QResource
	rv := nilthis.RegisterResource_1(rccData, resourceRoot)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:78
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool unregisterResource(const QString &, const QString &)
func (this *QResource) UnregisterResource(rccFilename string, resourceRoot string) bool {
	var tmpArg0 = NewQString_5(rccFilename)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(resourceRoot)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QResource18unregisterResourceERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QResource_UnregisterResource(rccFilename string, resourceRoot string) bool {
	var nilthis *QResource
	rv := nilthis.UnregisterResource(rccFilename, resourceRoot)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:81
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool unregisterResource(const uchar *, const QString &)
func (this *QResource) UnregisterResource_1(rccData unsafe.Pointer /*666*/, resourceRoot string) bool {
	var tmpArg1 = NewQString_5(resourceRoot)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QResource18unregisterResourceEPKhRK7QString", qtrt.FFI_TYPE_POINTER, &rccData, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}
func QResource_UnregisterResource_1(rccData unsafe.Pointer /*666*/, resourceRoot string) bool {
	var nilthis *QResource
	rv := nilthis.UnregisterResource_1(rccData, resourceRoot)
	return rv
}

// /usr/include/qt/QtCore/qresource.h:86
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool isDir()
func (this *QResource) IsDir() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QResource5isDirEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qresource.h:87
// index:0
// Protected inline Visibility=Default Availability=Available
// [1] bool isFile()
func (this *QResource) IsFile() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QResource6isFileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

//  body block end
