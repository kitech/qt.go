//  header block begin
// /usr/include/qt/QtCore/qmimetype.h
// #include <qmimetype.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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
type QMimeType struct {
	*qtrt.CObject
}

func (this *QMimeType) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQMimeTypeFromPointer(cthis unsafe.Pointer) *QMimeType {
	return &QMimeType{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qmimetype.h:78
// index:0
// Public
// void QMimeType()
func NewQMimeType() *QMimeType {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeTypeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMimeTypeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmimetype.h:84
// index:0
// Public inline
// void swap(class QMimeType &)
func (this *QMimeType) Swap(other *QMimeType) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeType4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:89
// index:0
// Public
// void ~QMimeType()
func DeleteQMimeType(*QMimeType) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeTypeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:98
// index:0
// Public
// bool isValid()
func (this *QMimeType) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:100
// index:0
// Public
// bool isDefault()
func (this *QMimeType) IsDefault() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType9isDefaultEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:102
// index:0
// Public
// QString name()
func (this *QMimeType) Name() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:103
// index:0
// Public
// QString comment()
func (this *QMimeType) Comment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType7commentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:104
// index:0
// Public
// QString genericIconName()
func (this *QMimeType) GenericIconName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType15genericIconNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:105
// index:0
// Public
// QString iconName()
func (this *QMimeType) IconName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType8iconNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:106
// index:0
// Public
// QStringList globPatterns()
func (this *QMimeType) GlobPatterns() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType12globPatternsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:107
// index:0
// Public
// QStringList parentMimeTypes()
func (this *QMimeType) ParentMimeTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType15parentMimeTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:108
// index:0
// Public
// QStringList allAncestors()
func (this *QMimeType) AllAncestors() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType12allAncestorsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:109
// index:0
// Public
// QStringList aliases()
func (this *QMimeType) Aliases() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType7aliasesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:110
// index:0
// Public
// QStringList suffixes()
func (this *QMimeType) Suffixes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType8suffixesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:111
// index:0
// Public
// QString preferredSuffix()
func (this *QMimeType) PreferredSuffix() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType15preferredSuffixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:113
// index:0
// Public
// bool inherits(const class QString &)
func (this *QMimeType) Inherits(mimeTypeName *QString) interface{} {
	var convArg0 = mimeTypeName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType8inheritsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimetype.h:115
// index:0
// Public
// QString filterString()
func (this *QMimeType) FilterString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType12filterStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
