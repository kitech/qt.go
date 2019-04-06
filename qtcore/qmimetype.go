// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qmimetype.h
// #include <qmimetype.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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
type QMimeType struct {
	*qtrt.CObject
}
type QMimeType_ITF interface {
	QMimeType_PTR() *QMimeType
}

func (ptr *QMimeType) QMimeType_PTR() *QMimeType { return ptr }

func (this *QMimeType) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMimeType) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMimeTypeFromPointer(cthis unsafe.Pointer) *QMimeType {
	return &QMimeType{&qtrt.CObject{cthis}}
}
func (*QMimeType) NewFromPointer(cthis unsafe.Pointer) *QMimeType {
	return NewQMimeTypeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmimetype.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMimeType()

/*
Constructs this QMimeType object initialized with default property values that indicate an invalid MIME type.
*/
func (*QMimeType) NewForInherit() *QMimeType {
	return NewQMimeType()
}
func NewQMimeType() *QMimeType {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeTypeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMimeType)
	return gothis
}

// /usr/include/qt/QtCore/qmimetype.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QMimeType & operator=(const QMimeType &)

/*

 */
func (this *QMimeType) Operator_equal(other QMimeType_ITF) *QMimeType {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMimeType_PTR() != nil {
		convArg0 = other.QMimeType_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeTypeaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QMimeType & operator=(QMimeType &&)

/*

 */
func (this *QMimeType) Operator_equal1(other unsafe.Pointer /*333*/) *QMimeType {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeTypeaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMimeType)
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QMimeType &)

/*
Swaps QMimeType other with this QMimeType object.

This operation is very fast and never fails.

The swap() method helps with the implementation of assignment operators in an exception-safe way. For more information consult More C++ Idioms - Copy-and-swap.
*/
func (this *QMimeType) Swap(other QMimeType_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMimeType_PTR() != nil {
		convArg0 = other.QMimeType_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeType4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMimeType()

/*

 */
func DeleteQMimeType(this *QMimeType) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeTypeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmimetype.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QMimeType &) const

/*

 */
func (this *QMimeType) Operator_equal_equal(other QMimeType_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMimeType_PTR() != nil {
		convArg0 = other.QMimeType_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeTypeeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QMimeType &) const

/*

 */
func (this *QMimeType) Operator_not_equal(other QMimeType_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMimeType_PTR() != nil {
		convArg0 = other.QMimeType_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeTypeneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QMimeType) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDefault() const

/*

 */
func (this *QMimeType) IsDefault() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType9isDefaultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*

 */
func (this *QMimeType) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimetype.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QString comment() const

/*

 */
func (this *QMimeType) Comment() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType7commentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimetype.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QString genericIconName() const

/*

 */
func (this *QMimeType) GenericIconName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType15genericIconNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimetype.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QString iconName() const

/*

 */
func (this *QMimeType) IconName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType8iconNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimetype.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList globPatterns() const

/*

 */
func (this *QMimeType) GlobPatterns() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType12globPatternsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList parentMimeTypes() const

/*

 */
func (this *QMimeType) ParentMimeTypes() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType15parentMimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList allAncestors() const

/*

 */
func (this *QMimeType) AllAncestors() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType12allAncestorsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList aliases() const

/*

 */
func (this *QMimeType) Aliases() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType7aliasesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList suffixes() const

/*

 */
func (this *QMimeType) Suffixes() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType8suffixesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QString preferredSuffix() const

/*

 */
func (this *QMimeType) PreferredSuffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType15preferredSuffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimetype.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inherits(const QString &) const

/*
Returns true if this mimetype is mimeTypeName, or inherits mimeTypeName (see parentMimeTypes()), or mimeTypeName is an alias for this mimetype.

This method has been made invokable from QML since 5.10.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.
*/
func (this *QMimeType) Inherits(mimeTypeName string) bool {
	var tmpArg0 = NewQString5(mimeTypeName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType8inheritsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filterString() const

/*

 */
func (this *QMimeType) FilterString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType12filterStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

//  body block end

//  keep block begin

func init_unused_10485() {
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
