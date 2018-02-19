package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QXmlStreamAttributes struct {
	*qtrt.CObject
}
type QXmlStreamAttributes_ITF interface {
	QXmlStreamAttributes_PTR() *QXmlStreamAttributes
}

func (ptr *QXmlStreamAttributes) QXmlStreamAttributes_PTR() *QXmlStreamAttributes { return ptr }

func (this *QXmlStreamAttributes) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamAttributes) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQXmlStreamAttributesFromPointer(cthis unsafe.Pointer) *QXmlStreamAttributes {
	return &QXmlStreamAttributes{&qtrt.CObject{cthis}}
}
func (*QXmlStreamAttributes) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamAttributes {
	return NewQXmlStreamAttributesFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:164
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QXmlStreamAttributes()
func NewQXmlStreamAttributes() *QXmlStreamAttributes {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QXmlStreamAttributesC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamAttributesFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamAttributes)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:165
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef value(const QString &, const QString &) const
func (this *QXmlStreamAttributes) Value(namespaceUri string, name string) *QStringRef /*123*/ {
	var tmpArg0 = NewQString_5(namespaceUri)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:166
// index:1
// Public Visibility=Default Availability=Available
// [16] QStringRef value(const QString &, QLatin1String) const
func (this *QXmlStreamAttributes) Value_1(namespaceUri string, name QLatin1String_ITF /*123*/) *QStringRef /*123*/ {
	var tmpArg0 = NewQString_5(namespaceUri)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if name != nil && name.QLatin1String_PTR() != nil {
		convArg1 = name.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueERK7QString13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:167
// index:2
// Public Visibility=Default Availability=Available
// [16] QStringRef value(QLatin1String, QLatin1String) const
func (this *QXmlStreamAttributes) Value_2(namespaceUri QLatin1String_ITF /*123*/, name QLatin1String_ITF /*123*/) *QStringRef /*123*/ {
	var convArg0 unsafe.Pointer
	if namespaceUri != nil && namespaceUri.QLatin1String_PTR() != nil {
		convArg0 = namespaceUri.QLatin1String_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if name != nil && name.QLatin1String_PTR() != nil {
		convArg1 = name.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueE13QLatin1StringS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:168
// index:3
// Public Visibility=Default Availability=Available
// [16] QStringRef value(const QString &) const
func (this *QXmlStreamAttributes) Value_3(qualifiedName string) *QStringRef /*123*/ {
	var tmpArg0 = NewQString_5(qualifiedName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:169
// index:4
// Public Visibility=Default Availability=Available
// [16] QStringRef value(QLatin1String) const
func (this *QXmlStreamAttributes) Value_4(qualifiedName QLatin1String_ITF /*123*/) *QStringRef /*123*/ {
	var convArg0 unsafe.Pointer
	if qualifiedName != nil && qualifiedName.QLatin1String_PTR() != nil {
		convArg0 = qualifiedName.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:173
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasAttribute(const QString &) const
func (this *QXmlStreamAttributes) HasAttribute(qualifiedName string) bool {
	var tmpArg0 = NewQString_5(qualifiedName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes12hasAttributeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:178
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool hasAttribute(QLatin1String) const
func (this *QXmlStreamAttributes) HasAttribute_1(qualifiedName QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if qualifiedName != nil && qualifiedName.QLatin1String_PTR() != nil {
		convArg0 = qualifiedName.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes12hasAttributeE13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:183
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool hasAttribute(const QString &, const QString &) const
func (this *QXmlStreamAttributes) HasAttribute_2(namespaceUri string, name string) bool {
	var tmpArg0 = NewQString_5(namespaceUri)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes12hasAttributeERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQXmlStreamAttributes(this *QXmlStreamAttributes) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QXmlStreamAttributesD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
