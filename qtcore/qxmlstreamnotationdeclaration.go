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
type QXmlStreamNotationDeclaration struct {
	*qtrt.CObject
}
type QXmlStreamNotationDeclaration_ITF interface {
	QXmlStreamNotationDeclaration_PTR() *QXmlStreamNotationDeclaration
}

func (ptr *QXmlStreamNotationDeclaration) QXmlStreamNotationDeclaration_PTR() *QXmlStreamNotationDeclaration {
	return ptr
}

func (this *QXmlStreamNotationDeclaration) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamNotationDeclaration) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQXmlStreamNotationDeclarationFromPointer(cthis unsafe.Pointer) *QXmlStreamNotationDeclaration {
	return &QXmlStreamNotationDeclaration{&qtrt.CObject{cthis}}
}
func (*QXmlStreamNotationDeclaration) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamNotationDeclaration {
	return NewQXmlStreamNotationDeclarationFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamNotationDeclaration()

/*

 */
func NewQXmlStreamNotationDeclaration() *QXmlStreamNotationDeclaration {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QXmlStreamNotationDeclarationC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamNotationDeclarationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamNotationDeclaration)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QXmlStreamNotationDeclaration()

/*

 */
func DeleteQXmlStreamNotationDeclaration(this *QXmlStreamNotationDeclaration) {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QXmlStreamNotationDeclarationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 56)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:253
// index:0
// Public Visibility=Default Availability=Available
// [56] QXmlStreamNotationDeclaration & operator=(const QXmlStreamNotationDeclaration &)

/*

 */
func (this *QXmlStreamNotationDeclaration) Operator_equal(arg0 QXmlStreamNotationDeclaration_ITF) *QXmlStreamNotationDeclaration {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QXmlStreamNotationDeclaration_PTR() != nil {
		convArg0 = arg0.QXmlStreamNotationDeclaration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QXmlStreamNotationDeclarationaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQXmlStreamNotationDeclarationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQXmlStreamNotationDeclaration)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:254
// index:1
// Public inline Visibility=Default Availability=Available
// [56] QXmlStreamNotationDeclaration & operator=(QXmlStreamNotationDeclaration &&)

/*

 */
func (this *QXmlStreamNotationDeclaration) Operator_equal_1(other unsafe.Pointer /*333*/) *QXmlStreamNotationDeclaration {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QXmlStreamNotationDeclarationaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQXmlStreamNotationDeclarationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQXmlStreamNotationDeclaration)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:264
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef name() const

/*

 */
func (this *QXmlStreamNotationDeclaration) Name() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QXmlStreamNotationDeclaration4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:265
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef systemId() const

/*

 */
func (this *QXmlStreamNotationDeclaration) SystemId() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QXmlStreamNotationDeclaration8systemIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:266
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef publicId() const

/*

 */
func (this *QXmlStreamNotationDeclaration) PublicId() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QXmlStreamNotationDeclaration8publicIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:267
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QXmlStreamNotationDeclaration &) const

/*

 */
func (this *QXmlStreamNotationDeclaration) Operator_equal_equal(other QXmlStreamNotationDeclaration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QXmlStreamNotationDeclaration_PTR() != nil {
		convArg0 = other.QXmlStreamNotationDeclaration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QXmlStreamNotationDeclarationeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:271
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QXmlStreamNotationDeclaration &) const

/*

 */
func (this *QXmlStreamNotationDeclaration) Operator_not_equal(other QXmlStreamNotationDeclaration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QXmlStreamNotationDeclaration_PTR() != nil {
		convArg0 = other.QXmlStreamNotationDeclaration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QXmlStreamNotationDeclarationneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
