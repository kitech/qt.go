//  header block begin
// /usr/include/qt/QtCore/qtranslator.h
// #include <qtranslator.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 35
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
type QTranslator struct {
	*QObject
}

func (this *QTranslator) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtCore/qtranslator.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTranslator) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTranslator10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtranslator.h:58
// index:0
// void QTranslator(class QObject *)
func NewQTranslator(parent unsafe.Pointer) *QTranslator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslatorC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTranslatorFromPointer(cthis)
	return gothis
}
func NewQTranslatorFromPointer(cthis unsafe.Pointer) *QTranslator {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QTranslator{bcthis0}
}

// /usr/include/qt/QtCore/qtranslator.h:59
// index:0
// virtual
// void ~QTranslator()
func DeleteQTranslator(*QTranslator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtranslator.h:61
// index:0
// virtual
// QString translate(const char *, const char *, const char *, int)
func (this *QTranslator) Translate(context unsafe.Pointer, sourceText unsafe.Pointer, disambiguation unsafe.Pointer, n int) {
	// 0: (, context const char *, sourceText const char *, disambiguation const char *, n int), (context, sourceText, disambiguation, &n)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTranslator9translateEPKcS1_S1_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), context, sourceText, disambiguation, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtranslator.h:64
// index:0
// virtual
// bool isEmpty()
func (this *QTranslator) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTranslator7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtranslator.h:66
// index:0
// bool load(const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QTranslator) Load(filename unsafe.Pointer, directory unsafe.Pointer, search_delimiters unsafe.Pointer, suffix unsafe.Pointer) {
	// 0: (, filename const QString &, directory const QString &, search_delimiters const QString &, suffix const QString &), (filename, directory, search_delimiters, suffix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QStringS2_S2_S2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filename, directory, search_delimiters, suffix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtranslator.h:70
// index:1
// bool load(const class QLocale &, const class QString &, const class QString &, const class QString &, const class QString &)
func (this *QTranslator) Load_1(locale unsafe.Pointer, filename unsafe.Pointer, prefix unsafe.Pointer, directory unsafe.Pointer, suffix unsafe.Pointer) {
	// 1: (, locale const QLocale &, filename const QString &, prefix const QString &, directory const QString &, suffix const QString &), (locale, filename, prefix, directory, suffix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), locale, filename, prefix, directory, suffix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtranslator.h:75
// index:2
// bool load(const uchar *, int, const class QString &)
func (this *QTranslator) Load_2(data unsafe.Pointer, len int, directory unsafe.Pointer) {
	// 2: (, data const uchar *, len int, directory const QString &), (data, &len, directory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTranslator4loadEPKhiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &len, directory)
	gopp.ErrPrint(err, rv)
}

//  body block end
