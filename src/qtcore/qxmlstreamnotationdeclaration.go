//  header block begin
// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QXmlStreamNotationDeclaration struct {
	*qtrt.CObject
}

func (this *QXmlStreamNotationDeclaration) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQXmlStreamNotationDeclarationFromPointer(cthis unsafe.Pointer) *QXmlStreamNotationDeclaration {
	return &QXmlStreamNotationDeclaration{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qxmlstream.h:241
// index:0
// Public
// void QXmlStreamNotationDeclaration()
func NewQXmlStreamNotationDeclaration() *QXmlStreamNotationDeclaration {
	cthis := qtrt.Calloc(1, 256) // 56
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QXmlStreamNotationDeclarationC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamNotationDeclarationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:243
// index:0
// Public
// void ~QXmlStreamNotationDeclaration()
func DeleteQXmlStreamNotationDeclaration(*QXmlStreamNotationDeclaration) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QXmlStreamNotationDeclarationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:264
// index:0
// Public inline
// QStringRef name()
func (this *QXmlStreamNotationDeclaration) Name() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QXmlStreamNotationDeclaration4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qxmlstream.h:265
// index:0
// Public inline
// QStringRef systemId()
func (this *QXmlStreamNotationDeclaration) SystemId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QXmlStreamNotationDeclaration8systemIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qxmlstream.h:266
// index:0
// Public inline
// QStringRef publicId()
func (this *QXmlStreamNotationDeclaration) PublicId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QXmlStreamNotationDeclaration8publicIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
