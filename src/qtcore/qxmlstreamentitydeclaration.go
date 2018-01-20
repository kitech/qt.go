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
type QXmlStreamEntityDeclaration struct {
	*qtrt.CObject
}

func (this *QXmlStreamEntityDeclaration) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qxmlstream.h:286
// index:0
// void QXmlStreamEntityDeclaration()
func NewQXmlStreamEntityDeclaration() *QXmlStreamEntityDeclaration {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QXmlStreamEntityDeclarationC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamEntityDeclarationFromPointer(cthis)
	return gothis
}
func NewQXmlStreamEntityDeclarationFromPointer(cthis unsafe.Pointer) *QXmlStreamEntityDeclaration {
	return &QXmlStreamEntityDeclaration{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qxmlstream.h:288
// index:0
// void ~QXmlStreamEntityDeclaration()
func DeleteQXmlStreamEntityDeclaration(*QXmlStreamEntityDeclaration) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QXmlStreamEntityDeclarationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:313
// index:0
// inline
// QStringRef name()
func (this *QXmlStreamEntityDeclaration) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration4nameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:314
// index:0
// inline
// QStringRef notationName()
func (this *QXmlStreamEntityDeclaration) NotationName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration12notationNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:315
// index:0
// inline
// QStringRef systemId()
func (this *QXmlStreamEntityDeclaration) SystemId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration8systemIdEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:316
// index:0
// inline
// QStringRef publicId()
func (this *QXmlStreamEntityDeclaration) PublicId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration8publicIdEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:317
// index:0
// inline
// QStringRef value()
func (this *QXmlStreamEntityDeclaration) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration5valueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
