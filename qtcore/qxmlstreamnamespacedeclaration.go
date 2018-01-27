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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQXmlStreamNamespaceDeclarationFromPointer(cthis unsafe.Pointer) *QXmlStreamNamespaceDeclaration {
	return &QXmlStreamNamespaceDeclaration{&qtrt.CObject{cthis}}
}
func (*QXmlStreamNamespaceDeclaration) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamNamespaceDeclaration {
	return NewQXmlStreamNamespaceDeclarationFromPointer(cthis)
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
// void QXmlStreamNamespaceDeclaration(const QString &, const QString &)
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
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclaration6prefixEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:222
// index:0
// Public inline
// QStringRef namespaceUri()
func (this *QXmlStreamNamespaceDeclaration) NamespaceUri() *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclaration12namespaceUriEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
