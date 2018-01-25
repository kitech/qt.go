package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
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
	*qtrt.CObject
}

func (this *QXmlStreamAttribute) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamAttribute) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQXmlStreamAttributeFromPointer(cthis unsafe.Pointer) *QXmlStreamAttribute {
	return &QXmlStreamAttribute{&qtrt.CObject{cthis}}
}
func (*QXmlStreamAttribute) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamAttribute {
	return NewQXmlStreamAttributeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:109
// index:0
// Public
// void QXmlStreamAttribute()
func NewQXmlStreamAttribute() *QXmlStreamAttribute {
	cthis := qtrt.Calloc(1, 256) // 80
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamAttributeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamAttributeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:110
// index:1
// Public
// void QXmlStreamAttribute(const class QString &, const class QString &)
func NewQXmlStreamAttribute_1(qualifiedName *QString, value *QString) *QXmlStreamAttribute {
	cthis := qtrt.Calloc(1, 256) // 80
	var convArg0 = qualifiedName.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamAttributeC2ERK7QStringS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamAttributeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:112
// index:2
// Public
// void QXmlStreamAttribute(const class QString &, const class QString &, const class QString &)
func NewQXmlStreamAttribute_2(namespaceUri *QString, name *QString, value *QString) *QXmlStreamAttribute {
	cthis := qtrt.Calloc(1, 256) // 80
	var convArg0 = namespaceUri.GetCthis()
	var convArg1 = name.GetCthis()
	var convArg2 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamAttributeC2ERK7QStringS2_S2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamAttributeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:137
// index:0
// Public
// void ~QXmlStreamAttribute()
func DeleteQXmlStreamAttribute(*QXmlStreamAttribute) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamAttributeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:140
// index:0
// Public inline
// QStringRef namespaceUri()
func (this *QXmlStreamAttribute) NamespaceUri() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute12namespaceUriEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:141
// index:0
// Public inline
// QStringRef name()
func (this *QXmlStreamAttribute) Name() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:142
// index:0
// Public inline
// QStringRef qualifiedName()
func (this *QXmlStreamAttribute) QualifiedName() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute13qualifiedNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:143
// index:0
// Public inline
// QStringRef prefix()
func (this *QXmlStreamAttribute) Prefix() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute6prefixEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:148
// index:0
// Public inline
// QStringRef value()
func (this *QXmlStreamAttribute) Value() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute5valueEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:149
// index:0
// Public inline
// bool isDefault()
func (this *QXmlStreamAttribute) IsDefault() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute9isDefaultEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
