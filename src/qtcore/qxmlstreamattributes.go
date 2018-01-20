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
type QXmlStreamAttributes struct {
	*qtrt.CObject
}

func (this *QXmlStreamAttributes) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQXmlStreamAttributesFromPointer(cthis unsafe.Pointer) *QXmlStreamAttributes {
	return &QXmlStreamAttributes{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qxmlstream.h:164
// index:0
// Public inline
// void QXmlStreamAttributes()
func NewQXmlStreamAttributes() *QXmlStreamAttributes {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QXmlStreamAttributesC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamAttributesFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:165
// index:0
// Public
// QStringRef value(const class QString &, const class QString &)
func (this *QXmlStreamAttributes) Value(namespaceUri *QString, name *QString) interface{} {
	var convArg0 = namespaceUri.GetCthis()
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qxmlstream.h:166
// index:1
// Public
// QStringRef value(const class QString &, class QLatin1String)
func (this *QXmlStreamAttributes) Value_1(namespaceUri *QString, name *QLatin1String) interface{} {
	var convArg0 = namespaceUri.GetCthis()
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueERK7QString13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qxmlstream.h:167
// index:2
// Public
// QStringRef value(class QLatin1String, class QLatin1String)
func (this *QXmlStreamAttributes) Value_2(namespaceUri *QLatin1String, name *QLatin1String) interface{} {
	var convArg0 = namespaceUri.GetCthis()
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueE13QLatin1StringS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qxmlstream.h:168
// index:3
// Public
// QStringRef value(const class QString &)
func (this *QXmlStreamAttributes) Value_3(qualifiedName *QString) interface{} {
	var convArg0 = qualifiedName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qxmlstream.h:169
// index:4
// Public
// QStringRef value(class QLatin1String)
func (this *QXmlStreamAttributes) Value_4(qualifiedName *QLatin1String) interface{} {
	var convArg0 = qualifiedName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qxmlstream.h:173
// index:0
// Public inline
// bool hasAttribute(const class QString &)
func (this *QXmlStreamAttributes) HasAttribute(qualifiedName *QString) interface{} {
	var convArg0 = qualifiedName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes12hasAttributeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qxmlstream.h:178
// index:1
// Public inline
// bool hasAttribute(class QLatin1String)
func (this *QXmlStreamAttributes) HasAttribute_1(qualifiedName *QLatin1String) interface{} {
	var convArg0 = qualifiedName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes12hasAttributeE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qxmlstream.h:183
// index:2
// Public inline
// bool hasAttribute(const class QString &, const class QString &)
func (this *QXmlStreamAttributes) HasAttribute_2(namespaceUri *QString, name *QString) interface{} {
	var convArg0 = namespaceUri.GetCthis()
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes12hasAttributeERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
