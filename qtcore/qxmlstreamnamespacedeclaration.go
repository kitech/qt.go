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
type QXmlStreamNamespaceDeclaration struct {
	*qtrt.CObject
}
type QXmlStreamNamespaceDeclaration_ITF interface {
	QXmlStreamNamespaceDeclaration_PTR() *QXmlStreamNamespaceDeclaration
}

func (ptr *QXmlStreamNamespaceDeclaration) QXmlStreamNamespaceDeclaration_PTR() *QXmlStreamNamespaceDeclaration {
	return ptr
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

/*

 */
func (*QXmlStreamNamespaceDeclaration) NewForInherit() *QXmlStreamNamespaceDeclaration {
	return NewQXmlStreamNamespaceDeclaration()
}
func NewQXmlStreamNamespaceDeclaration() *QXmlStreamNamespaceDeclaration {
	rv, err := qtrt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamNamespaceDeclarationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamNamespaceDeclaration)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:216
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamNamespaceDeclaration(const QString &, const QString &)

/*

 */
func (*QXmlStreamNamespaceDeclaration) NewForInherit1(prefix string, namespaceUri string) *QXmlStreamNamespaceDeclaration {
	return NewQXmlStreamNamespaceDeclaration1(prefix, namespaceUri)
}
func NewQXmlStreamNamespaceDeclaration1(prefix string, namespaceUri string) *QXmlStreamNamespaceDeclaration {
	var tmpArg0 = NewQString5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(namespaceUri)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationC2ERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamNamespaceDeclarationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamNamespaceDeclaration)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:209
// index:0
// Public inline Visibility=Default Availability=Available
// [40] QXmlStreamNamespaceDeclaration & operator=(QXmlStreamNamespaceDeclaration &&)

/*

 */
func (this *QXmlStreamNamespaceDeclaration) Operator_equal(other unsafe.Pointer /*333*/) *QXmlStreamNamespaceDeclaration {
	rv, err := qtrt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQXmlStreamNamespaceDeclarationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQXmlStreamNamespaceDeclaration)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:218
// index:1
// Public Visibility=Default Availability=Available
// [40] QXmlStreamNamespaceDeclaration & operator=(const QXmlStreamNamespaceDeclaration &)

/*

 */
func (this *QXmlStreamNamespaceDeclaration) Operator_equal1(arg0 QXmlStreamNamespaceDeclaration_ITF) *QXmlStreamNamespaceDeclaration {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QXmlStreamNamespaceDeclaration_PTR() != nil {
		convArg0 = arg0.QXmlStreamNamespaceDeclaration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQXmlStreamNamespaceDeclarationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQXmlStreamNamespaceDeclaration)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QXmlStreamNamespaceDeclaration()

/*

 */
func DeleteQXmlStreamNamespaceDeclaration(this *QXmlStreamNamespaceDeclaration) {
	rv, err := qtrt.InvokeQtFunc6("_ZN30QXmlStreamNamespaceDeclarationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef prefix() const

/*

 */
func (this *QXmlStreamNamespaceDeclaration) Prefix() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclaration6prefixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef namespaceUri() const

/*

 */
func (this *QXmlStreamNamespaceDeclaration) NamespaceUri() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclaration12namespaceUriEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:223
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QXmlStreamNamespaceDeclaration &) const

/*

 */
func (this *QXmlStreamNamespaceDeclaration) Operator_equal_equal(other QXmlStreamNamespaceDeclaration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QXmlStreamNamespaceDeclaration_PTR() != nil {
		convArg0 = other.QXmlStreamNamespaceDeclaration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclarationeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qxmlstream.h:226
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QXmlStreamNamespaceDeclaration &) const

/*

 */
func (this *QXmlStreamNamespaceDeclaration) Operator_not_equal(other QXmlStreamNamespaceDeclaration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QXmlStreamNamespaceDeclaration_PTR() != nil {
		convArg0 = other.QXmlStreamNamespaceDeclaration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK30QXmlStreamNamespaceDeclarationneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
