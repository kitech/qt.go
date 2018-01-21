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
func NewQXmlStreamNamespaceDeclarationFromPointer(cthis unsafe.Pointer) *QXmlStreamNamespaceDeclaration {
	return &QXmlStreamNamespaceDeclaration{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qxmlstream.h:199
// index:0
// Public
// void QXmlStreamNamespaceDeclaration()
func NewQXmlStreamNamespaceDeclaration() *QXmlStreamNamespaceDeclaration {
	cthis := qtrt.Calloc(1, 256) // 40
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamNamespaceDeclarationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:216
// index:1
// Public
// void QXmlStreamNamespaceDeclaration(const class QString &, const class QString &)
func NewQXmlStreamNamespaceDeclaration_1(prefix *QString, namespaceUri *QString) *QXmlStreamNamespaceDeclaration {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = prefix.GetCthis()
	var convArg1 = namespaceUri.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationC2ERK7QStringS2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamNamespaceDeclarationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:217
// index:0
// Public
// void ~QXmlStreamNamespaceDeclaration()
func DeleteQXmlStreamNamespaceDeclaration(*QXmlStreamNamespaceDeclaration) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:221
// index:0
// Public inline
// QStringRef prefix()
func (this *QXmlStreamNamespaceDeclaration) Prefix() *QStringRef /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclaration6prefixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:222
// index:0
// Public inline
// QStringRef namespaceUri()
func (this *QXmlStreamNamespaceDeclaration) NamespaceUri() *QStringRef /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclaration12namespaceUriEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
