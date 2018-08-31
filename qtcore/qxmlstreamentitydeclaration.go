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
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
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

/*

 */
func (*QXmlStreamEntityDeclaration) NewForInherit() *QXmlStreamEntityDeclaration {
	return NewQXmlStreamEntityDeclaration()
}
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

/*

 */
func DeleteQXmlStreamEntityDeclaration(this *QXmlStreamEntityDeclaration) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QXmlStreamEntityDeclarationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 88)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:300
// index:0
// Public Visibility=Default Availability=Available
// [88] QXmlStreamEntityDeclaration & operator=(const QXmlStreamEntityDeclaration &)

/*

 */
func (this *QXmlStreamEntityDeclaration) Operator_equal(arg0 QXmlStreamEntityDeclaration_ITF) *QXmlStreamEntityDeclaration {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QXmlStreamEntityDeclaration_PTR() != nil {
		convArg0 = arg0.QXmlStreamEntityDeclaration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QXmlStreamEntityDeclarationaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQXmlStreamEntityDeclarationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQXmlStreamEntityDeclaration)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:301
// index:1
// Public inline Visibility=Default Availability=Available
// [88] QXmlStreamEntityDeclaration & operator=(QXmlStreamEntityDeclaration &&)

/*

 */
func (this *QXmlStreamEntityDeclaration) Operator_equal_1(other unsafe.Pointer /*333*/) *QXmlStreamEntityDeclaration {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QXmlStreamEntityDeclarationaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQXmlStreamEntityDeclarationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQXmlStreamEntityDeclaration)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:313
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef name() const

/*

 */
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
// [16] QStringRef notationName() const

/*

 */
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
// [16] QStringRef systemId() const

/*

 */
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
// [16] QStringRef publicId() const

/*

 */
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
// [16] QStringRef value() const

/*

 */
func (this *QXmlStreamEntityDeclaration) Value() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclaration5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:318
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QXmlStreamEntityDeclaration &) const

/*

 */
func (this *QXmlStreamEntityDeclaration) Operator_equal_equal(other QXmlStreamEntityDeclaration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QXmlStreamEntityDeclaration_PTR() != nil {
		convArg0 = other.QXmlStreamEntityDeclaration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclarationeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:325
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QXmlStreamEntityDeclaration &) const

/*

 */
func (this *QXmlStreamEntityDeclaration) Operator_not_equal(other QXmlStreamEntityDeclaration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QXmlStreamEntityDeclaration_PTR() != nil {
		convArg0 = other.QXmlStreamEntityDeclaration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QXmlStreamEntityDeclarationneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
