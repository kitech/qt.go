package qtcore

// /usr/include/qt/QtCore/qtranslator.h
// #include <qtranslator.h>
// #include <QtCore>

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
import "mkuse/cffiqt"
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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTranslator struct {
	*QObject
}

func (this *QTranslator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QTranslator) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQTranslatorFromPointer(cthis unsafe.Pointer) *QTranslator {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QTranslator{bcthis0}
}
func (*QTranslator) NewFromPointer(cthis unsafe.Pointer) *QTranslator {
	return NewQTranslatorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtranslator.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTranslator) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTranslator10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qtranslator.h:58
// index:0
// Public
// void QTranslator(class QObject *)
func NewQTranslator(parent *QObject /*777 QObject **/) *QTranslator {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslatorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTranslatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtranslator.h:59
// index:0
// Public virtual
// void ~QTranslator()
func DeleteQTranslator(*QTranslator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtranslator.h:61
// index:0
// Public virtual
// QString translate(const char *, const char *, const char *, int)
func (this *QTranslator) Translate(context string, sourceText string, disambiguation string, n int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(sourceText)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(disambiguation)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTranslator9translateEPKcS1_S1_i", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1, convArg2, n)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtranslator.h:64
// index:0
// Public virtual
// bool isEmpty()
func (this *QTranslator) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTranslator7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:66
// index:0
// Public
// bool load(const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QTranslator) Load(filename *QString, directory *QString, search_delimiters *QString, suffix *QString) bool {
	var convArg0 = filename.GetCthis()
	var convArg1 = directory.GetCthis()
	var convArg2 = search_delimiters.GetCthis()
	var convArg3 = suffix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QStringS2_S2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:70
// index:1
// Public
// bool load(const class QLocale &, const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QTranslator) Load_1(locale *QLocale, filename *QString, prefix *QString, directory *QString, suffix *QString) bool {
	var convArg0 = locale.GetCthis()
	var convArg1 = filename.GetCthis()
	var convArg2 = prefix.GetCthis()
	var convArg3 = directory.GetCthis()
	var convArg4 = suffix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:75
// index:2
// Public
// bool load(const uchar *, int, const class QString &)
func (this *QTranslator) Load_2(data unsafe.Pointer /*666*/, len int, directory *QString) bool {
	var convArg2 = directory.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslator4loadEPKhiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &data, len, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
