package qtcore

// /usr/include/qt/QtCore/qmimetype.h
// #include <qmimetype.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QMimeType struct {
	*qtrt.CObject
}

func (this *QMimeType) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMimeType) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeTypeC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQMimeTypeFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qmimetype.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QMimeType &)
func (this *QMimeType) Swap(other *QMimeType) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeType4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMimeType()
func DeleteQMimeType(*QMimeType) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeTypeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QMimeType) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDefault()
func (this *QMimeType) IsDefault() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType9isDefaultEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QMimeType) Name() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QString comment()
func (this *QMimeType) Comment() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType7commentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QString genericIconName()
func (this *QMimeType) GenericIconName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType15genericIconNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QString iconName()
func (this *QMimeType) IconName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType8iconNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QString preferredSuffix()
func (this *QMimeType) PreferredSuffix() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType15preferredSuffixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmimetype.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inherits(const QString &)
func (this *QMimeType) Inherits(mimeTypeName *QString) bool {
	var convArg0 = mimeTypeName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType8inheritsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmimetype.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filterString()
func (this *QMimeType) FilterString() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType12filterStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
