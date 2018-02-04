package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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

type QXmlStreamNamespaceDeclaration struct {
	*qtrt.CObject
}

func (this *QXmlStreamNamespaceDeclaration) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamNamespaceDeclaration) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQXmlStreamNamespaceDeclarationFromPointer(cthis unsafe.Pointer) *QXmlStreamNamespaceDeclaration {
	return &QXmlStreamNamespaceDeclaration{&qtrt.CObject{cthis}}
}
func (*QXmlStreamNamespaceDeclaration) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamNamespaceDeclaration {
	return NewQXmlStreamNamespaceDeclarationFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamNamespaceDeclaration()
func NewQXmlStreamNamespaceDeclaration() *QXmlStreamNamespaceDeclaration {
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamNamespaceDeclarationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamNamespaceDeclaration)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:216
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamNamespaceDeclaration(const QString &, const QString &)
func NewQXmlStreamNamespaceDeclaration_1(prefix *QString, namespaceUri *QString) *QXmlStreamNamespaceDeclaration {
	var convArg0 = prefix.GetCthis()
	var convArg1 = namespaceUri.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationC2ERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamNamespaceDeclarationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamNamespaceDeclaration)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QXmlStreamNamespaceDeclaration()
func DeleteQXmlStreamNamespaceDeclaration(this *QXmlStreamNamespaceDeclaration) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef prefix()
func (this *QXmlStreamNamespaceDeclaration) Prefix() *QStringRef /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclaration6prefixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef namespaceUri()
func (this *QXmlStreamNamespaceDeclaration) NamespaceUri() *QStringRef /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclaration12namespaceUriEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

//  body block end
