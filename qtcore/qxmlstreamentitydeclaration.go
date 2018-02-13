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
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QXmlStreamEntityDeclaration struct {
	*qtrt.CObject
}
type QXmlStreamEntityDeclaration_ITF interface {
	QXmlStreamEntityDeclaration_PTR() *QXmlStreamEntityDeclaration
}

func (ptr *QXmlStreamEntityDeclaration) QXmlStreamEntityDeclaration_PTR() *QXmlStreamEntityDeclaration {
	return ptr
}

func (this *QXmlStreamEntityDeclaration) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamEntityDeclaration) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZN27QXmlStreamEntityDeclarationC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamEntityDeclarationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamEntityDeclaration)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QXmlStreamEntityDeclaration()
func DeleteQXmlStreamEntityDeclaration(this *QXmlStreamEntityDeclaration) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QXmlStreamEntityDeclarationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 88)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:313
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef name()
func (this *QXmlStreamEntityDeclaration) Name() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:314
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef notationName()
func (this *QXmlStreamEntityDeclaration) NotationName() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration12notationNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:315
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef systemId()
func (this *QXmlStreamEntityDeclaration) SystemId() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration8systemIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:316
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef publicId()
func (this *QXmlStreamEntityDeclaration) PublicId() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration8publicIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:317
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef value()
func (this *QXmlStreamEntityDeclaration) Value() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
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
		qtrt.KeepMe()
	}
}

//  keep block end
