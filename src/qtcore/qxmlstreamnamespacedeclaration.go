//  header block begin
// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>
package qtcore

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qxmlstream.h:199
// index:0
// void QXmlStreamNamespaceDeclaration()
func NewQXmlStreamNamespaceDeclaration() *QXmlStreamNamespaceDeclaration {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QXmlStreamNamespaceDeclaration{cthis}
}

// /usr/include/qt/QtCore/qxmlstream.h:216
// index:1
// void QXmlStreamNamespaceDeclaration(const class QString &, const class QString &)
func NewQXmlStreamNamespaceDeclaration_1(prefix unsafe.Pointer, namespaceUri unsafe.Pointer) *QXmlStreamNamespaceDeclaration {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationC2ERK7QStringS2_", ffiqt.FFI_TYPE_VOID, cthis, prefix, namespaceUri)
	gopp.ErrPrint(err, rv)
	return &QXmlStreamNamespaceDeclaration{cthis}
}

// /usr/include/qt/QtCore/qxmlstream.h:217
// index:0
// void ~QXmlStreamNamespaceDeclaration()
func DeleteQXmlStreamNamespaceDeclaration(*QXmlStreamNamespaceDeclaration) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:221
// index:0
// inline
// QStringRef prefix()
func (this *QXmlStreamNamespaceDeclaration) Prefix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclaration6prefixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:222
// index:0
// inline
// QStringRef namespaceUri()
func (this *QXmlStreamNamespaceDeclaration) NamespaceUri() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclaration12namespaceUriEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
