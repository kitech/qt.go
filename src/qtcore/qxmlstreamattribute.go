//  header block begin
// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
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
type QXmlStreamAttribute struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qxmlstream.h:109
// index:0
// void QXmlStreamAttribute()
func NewQXmlStreamAttribute() *QXmlStreamAttribute {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamAttributeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QXmlStreamAttribute{cthis}
}

// /usr/include/qt/QtCore/qxmlstream.h:110
// index:1
// void QXmlStreamAttribute(const class QString &, const class QString &)
func NewQXmlStreamAttribute_1(qualifiedName unsafe.Pointer, value unsafe.Pointer) *QXmlStreamAttribute {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamAttributeC2ERK7QStringS2_", ffiqt.FFI_TYPE_VOID, cthis, qualifiedName, value)
	gopp.ErrPrint(err, rv)
	return &QXmlStreamAttribute{cthis}
}

// /usr/include/qt/QtCore/qxmlstream.h:112
// index:2
// void QXmlStreamAttribute(const class QString &, const class QString &, const class QString &)
func NewQXmlStreamAttribute_2(namespaceUri unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) *QXmlStreamAttribute {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamAttributeC2ERK7QStringS2_S2_", ffiqt.FFI_TYPE_VOID, cthis, namespaceUri, name, value)
	gopp.ErrPrint(err, rv)
	return &QXmlStreamAttribute{cthis}
}

// /usr/include/qt/QtCore/qxmlstream.h:137
// index:0
// void ~QXmlStreamAttribute()
func DeleteQXmlStreamAttribute(*QXmlStreamAttribute) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamAttributeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:140
// index:0
// inline
// QStringRef namespaceUri()
func (this *QXmlStreamAttribute) NamespaceUri() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute12namespaceUriEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:141
// index:0
// inline
// QStringRef name()
func (this *QXmlStreamAttribute) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:142
// index:0
// inline
// QStringRef qualifiedName()
func (this *QXmlStreamAttribute) QualifiedName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute13qualifiedNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:143
// index:0
// inline
// QStringRef prefix()
func (this *QXmlStreamAttribute) Prefix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute6prefixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:148
// index:0
// inline
// QStringRef value()
func (this *QXmlStreamAttribute) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute5valueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:149
// index:0
// inline
// bool isDefault()
func (this *QXmlStreamAttribute) IsDefault() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute9isDefaultEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
