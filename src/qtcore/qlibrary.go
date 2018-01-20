//  header block begin
// /usr/include/qt/QtCore/qlibrary.h
// #include <qlibrary.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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

// /usr/include/qt/QtCore/qlibrary.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QLibrary) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:68
// index:0
// void QLibrary(class QObject *)
func NewQLibrary(parent unsafe.Pointer) *QLibrary {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(cthis)
	return gothis
}
func NewQLibraryFromPointer(cthis unsafe.Pointer) *QLibrary {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QLibrary{bcthis0}
}

// /usr/include/qt/QtCore/qlibrary.h:69
// index:1
// void QLibrary(const class QString &, class QObject *)
func NewQLibrary_1(fileName unsafe.Pointer, parent unsafe.Pointer) *QLibrary {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, fileName, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:70
// index:2
// void QLibrary(const class QString &, int, class QObject *)
func NewQLibrary_2(fileName unsafe.Pointer, verNum int, parent unsafe.Pointer) *QLibrary {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, fileName, &verNum, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:71
// index:3
// void QLibrary(const class QString &, const class QString &, class QObject *)
func NewQLibrary_3(fileName unsafe.Pointer, version unsafe.Pointer, parent unsafe.Pointer) *QLibrary {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringS2_P7QObject", ffiqt.FFI_TYPE_VOID, cthis, fileName, version, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:72
// index:0
// virtual
// void ~QLibrary()
func DeleteQLibrary(*QLibrary) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibraryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:74
// index:0
// QFunctionPointer resolve(const char *)
func (this *QLibrary) Resolve(symbol unsafe.Pointer) {
	// 0: (, symbol const char *), (symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), symbol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:75
// index:1
// static
// QFunctionPointer resolve(const class QString &, const char *)
func (this *QLibrary) Resolve_1(fileName unsafe.Pointer, symbol unsafe.Pointer) {
	// 1: (fileName const QString &, symbol const char *), (fileName, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringPKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLibrary_Resolve_1(fileName unsafe.Pointer, symbol unsafe.Pointer) {
	// 1: (fileName const QString &, symbol const char *), (fileName, symbol)
	var nilthis *QLibrary
	nilthis.Resolve_1(fileName, symbol)
}

// /usr/include/qt/QtCore/qlibrary.h:76
// index:2
// static
// QFunctionPointer resolve(const class QString &, int, const char *)
func (this *QLibrary) Resolve_2(fileName unsafe.Pointer, verNum int, symbol unsafe.Pointer) {
	// 2: (fileName const QString &, verNum int, symbol const char *), (fileName, verNum, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringiPKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLibrary_Resolve_2(fileName unsafe.Pointer, verNum int, symbol unsafe.Pointer) {
	// 2: (fileName const QString &, verNum int, symbol const char *), (fileName, verNum, symbol)
	var nilthis *QLibrary
	nilthis.Resolve_2(fileName, verNum, symbol)
}

// /usr/include/qt/QtCore/qlibrary.h:77
// index:3
// static
// QFunctionPointer resolve(const class QString &, const class QString &, const char *)
func (this *QLibrary) Resolve_3(fileName unsafe.Pointer, version unsafe.Pointer, symbol unsafe.Pointer) {
	// 3: (fileName const QString &, version const QString &, symbol const char *), (fileName, version, symbol)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringS2_PKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLibrary_Resolve_3(fileName unsafe.Pointer, version unsafe.Pointer, symbol unsafe.Pointer) {
	// 3: (fileName const QString &, version const QString &, symbol const char *), (fileName, version, symbol)
	var nilthis *QLibrary
	nilthis.Resolve_3(fileName, version, symbol)
}

// /usr/include/qt/QtCore/qlibrary.h:79
// index:0
// bool load()
func (this *QLibrary) Load() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary4loadEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:80
// index:0
// bool unload()
func (this *QLibrary) Unload() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary6unloadEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:81
// index:0
// bool isLoaded()
func (this *QLibrary) IsLoaded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary8isLoadedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:83
// index:0
// static
// bool isLibrary(const class QString &)
func (this *QLibrary) IsLibrary(fileName unsafe.Pointer) {
	// 0: (fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary9isLibraryERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLibrary_IsLibrary(fileName unsafe.Pointer) {
	// 0: (fileName const QString &), (fileName)
	var nilthis *QLibrary
	nilthis.IsLibrary(fileName)
}

// /usr/include/qt/QtCore/qlibrary.h:85
// index:0
// void setFileName(const class QString &)
func (this *QLibrary) SetFileName(fileName unsafe.Pointer) {
	// 0: (, fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary11setFileNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:86
// index:0
// QString fileName()
func (this *QLibrary) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary8fileNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:88
// index:0
// void setFileNameAndVersion(const class QString &, int)
func (this *QLibrary) SetFileNameAndVersion(fileName unsafe.Pointer, verNum int) {
	// 0: (, fileName const QString &, verNum int), (fileName, &verNum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary21setFileNameAndVersionERK7QStringi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, &verNum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:89
// index:1
// void setFileNameAndVersion(const class QString &, const class QString &)
func (this *QLibrary) SetFileNameAndVersion_1(fileName unsafe.Pointer, version unsafe.Pointer) {
	// 1: (, fileName const QString &, version const QString &), (fileName, version)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary21setFileNameAndVersionERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, version)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:90
// index:0
// QString errorString()
func (this *QLibrary) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary11errorStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:92
// index:0
// void setLoadHints(QLibrary::LoadHints)
func (this *QLibrary) SetLoadHints(hints int) {
	// 0: (, QFlags<QLibrary::LoadHint> hints), (hints)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QLibrary12setLoadHintsE6QFlagsINS_8LoadHintEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), hints)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:93
// index:0
// QLibrary::LoadHints loadHints()
func (this *QLibrary) LoadHints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QLibrary9loadHintsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
