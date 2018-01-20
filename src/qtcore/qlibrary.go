//  header block begin
// /usr/include/qt/QtCore/qlibrary.h
// #include <qlibrary.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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
type QLibrary struct {
	*QObject
}

func (this *QLibrary) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQLibraryFromPointer(cthis unsafe.Pointer) *QLibrary {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QLibrary{bcthis0}
}

// /usr/include/qt/QtCore/qlibrary.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QLibrary) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:68
// index:0
// Public
// void QLibrary(class QObject *)
func NewQLibrary(parent unsafe.Pointer) *QLibrary {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:69
// index:1
// Public
// void QLibrary(const class QString &, class QObject *)
func NewQLibrary_1(fileName *QString, parent unsafe.Pointer) *QLibrary {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:70
// index:2
// Public
// void QLibrary(const class QString &, int, class QObject *)
func NewQLibrary_2(fileName *QString, verNum int, parent unsafe.Pointer) *QLibrary {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &verNum, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:71
// index:3
// Public
// void QLibrary(const class QString &, const class QString &, class QObject *)
func NewQLibrary_3(fileName *QString, version *QString, parent unsafe.Pointer) *QLibrary {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = fileName.GetCthis()
	var convArg1 = version.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringS2_P7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:72
// index:0
// Public virtual
// void ~QLibrary()
func DeleteQLibrary(*QLibrary) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:74
// index:0
// Public
// QFunctionPointer resolve(const char *)
func (this *QLibrary) Resolve(symbol string) interface{} {
	var convArg0 = qtrt.CString(symbol)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:75
// index:1
// Public static
// QFunctionPointer resolve(const class QString &, const char *)
func (this *QLibrary) Resolve_1(fileName *QString, symbol string) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringPKc", ffiqt.FFI_TYPE_POINTER, fileName, symbol)
	gopp.ErrPrint(err, rv)
	return rv
}
func QLibrary_Resolve_1(fileName *QString, symbol string) {
	var nilthis *QLibrary
	nilthis.Resolve_1(fileName, symbol)
}

// /usr/include/qt/QtCore/qlibrary.h:76
// index:2
// Public static
// QFunctionPointer resolve(const class QString &, int, const char *)
func (this *QLibrary) Resolve_2(fileName *QString, verNum int, symbol string) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringiPKc", ffiqt.FFI_TYPE_POINTER, fileName, verNum, symbol)
	gopp.ErrPrint(err, rv)
	return rv
}
func QLibrary_Resolve_2(fileName *QString, verNum int, symbol string) {
	var nilthis *QLibrary
	nilthis.Resolve_2(fileName, verNum, symbol)
}

// /usr/include/qt/QtCore/qlibrary.h:77
// index:3
// Public static
// QFunctionPointer resolve(const class QString &, const class QString &, const char *)
func (this *QLibrary) Resolve_3(fileName *QString, version *QString, symbol string) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringS2_PKc", ffiqt.FFI_TYPE_POINTER, fileName, version, symbol)
	gopp.ErrPrint(err, rv)
	return rv
}
func QLibrary_Resolve_3(fileName *QString, version *QString, symbol string) {
	var nilthis *QLibrary
	nilthis.Resolve_3(fileName, version, symbol)
}

// /usr/include/qt/QtCore/qlibrary.h:79
// index:0
// Public
// bool load()
func (this *QLibrary) Load() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary4loadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:80
// index:0
// Public
// bool unload()
func (this *QLibrary) Unload() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary6unloadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:81
// index:0
// Public
// bool isLoaded()
func (this *QLibrary) IsLoaded() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary8isLoadedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:83
// index:0
// Public static
// bool isLibrary(const class QString &)
func (this *QLibrary) IsLibrary(fileName *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary9isLibraryERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QLibrary_IsLibrary(fileName *QString) {
	var nilthis *QLibrary
	nilthis.IsLibrary(fileName)
}

// /usr/include/qt/QtCore/qlibrary.h:85
// index:0
// Public
// void setFileName(const class QString &)
func (this *QLibrary) SetFileName(fileName *QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:86
// index:0
// Public
// QString fileName()
func (this *QLibrary) FileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:88
// index:0
// Public
// void setFileNameAndVersion(const class QString &, int)
func (this *QLibrary) SetFileNameAndVersion(fileName *QString, verNum int) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary21setFileNameAndVersionERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &verNum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:89
// index:1
// Public
// void setFileNameAndVersion(const class QString &, const class QString &)
func (this *QLibrary) SetFileNameAndVersion_1(fileName *QString, version *QString) {
	var convArg0 = fileName.GetCthis()
	var convArg1 = version.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary21setFileNameAndVersionERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:90
// index:0
// Public
// QString errorString()
func (this *QLibrary) ErrorString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:93
// index:0
// Public
// QLibrary::LoadHints loadHints()
func (this *QLibrary) LoadHints() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary9loadHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
