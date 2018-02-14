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
// extern C begin: 23
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

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
func NewQMimeType() *QMimeType {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeTypeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMimeType)
	return gothis
}

// /usr/include/qt/QtCore/qmimetype.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QMimeType &)
func (this *QMimeType) Swap(other QMimeType_ITF) {
	var convArg0 = other.QMimeType_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeType4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMimeType()
func DeleteQMimeType(this *QMimeType) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMimeTypeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmimetype.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QMimeType) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDefault()
func (this *QMimeType) IsDefault() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType9isDefaultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QMimeType) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimetype.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QString comment()
func (this *QMimeType) Comment() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType7commentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimetype.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QString genericIconName()
func (this *QMimeType) GenericIconName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType15genericIconNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimetype.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QString iconName()
func (this *QMimeType) IconName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType8iconNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimetype.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList globPatterns()
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
// [8] QStringList parentMimeTypes()
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
// [8] QStringList allAncestors()
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
// [8] QStringList aliases()
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
// [8] QStringList suffixes()
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
// [8] QString preferredSuffix()
func (this *QMimeType) PreferredSuffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType15preferredSuffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmimetype.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inherits(const QString &)
func (this *QMimeType) Inherits(mimeTypeName string) bool {
	var tmpArg0 = NewQString_5(mimeTypeName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType8inheritsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filterString()
func (this *QMimeType) FilterString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMimeType12filterStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
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
