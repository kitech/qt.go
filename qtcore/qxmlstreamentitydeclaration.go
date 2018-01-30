package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

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
type QXmlStreamEntityDeclaration struct {
	*qtrt.CObject
}

func (this *QXmlStreamEntityDeclaration) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamEntityDeclaration) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQXmlStreamEntityDeclarationFromPointer(cthis unsafe.Pointer) *QXmlStreamEntityDeclaration {
	return &QXmlStreamEntityDeclaration{&qtrt.CObject{cthis}}
}
func (*QXmlStreamEntityDeclaration) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamEntityDeclaration {
	return NewQXmlStreamEntityDeclarationFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:286
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamEntityDeclaration()
func NewQXmlStreamEntityDeclaration() *QXmlStreamEntityDeclaration {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QXmlStreamEntityDeclarationC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamEntityDeclarationFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QXmlStreamEntityDeclaration()
func DeleteQXmlStreamEntityDeclaration(*QXmlStreamEntityDeclaration) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QXmlStreamEntityDeclarationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:313
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef name()
func (this *QXmlStreamEntityDeclaration) Name() *QStringRef /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:314
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef notationName()
func (this *QXmlStreamEntityDeclaration) NotationName() *QStringRef /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration12notationNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:315
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef systemId()
func (this *QXmlStreamEntityDeclaration) SystemId() *QStringRef /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration8systemIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:316
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef publicId()
func (this *QXmlStreamEntityDeclaration) PublicId() *QStringRef /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration8publicIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:317
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef value()
func (this *QXmlStreamEntityDeclaration) Value() *QStringRef /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
