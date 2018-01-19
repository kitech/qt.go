//  header block begin
// /usr/include/qt/QtCore/qmimetype.h
// #include <qmimetype.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qmimetype.h:78
// index:0
// void QMimeType()
func NewQMimeType() *QMimeType {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeTypeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QMimeType{cthis}
}

// /usr/include/qt/QtCore/qmimetype.h:88
// index:1
// void QMimeType(const class QMimeTypePrivate &)
func NewQMimeType_1(dd unsafe.Pointer) *QMimeType {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeTypeC2ERK16QMimeTypePrivate", ffiqt.FFI_TYPE_VOID, cthis, dd)
	gopp.ErrPrint(err, rv)
	return &QMimeType{cthis}
}

// /usr/include/qt/QtCore/qmimetype.h:84
// index:0
// inline
// void swap(class QMimeType &)
func (this *QMimeType) Swap(other unsafe.Pointer) {
	// 0: (, QMimeType & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeType4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:89
// index:0
// void ~QMimeType()
func DeleteQMimeType(*QMimeType) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeTypeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:98
// index:0
// bool isValid()
func (this *QMimeType) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:100
// index:0
// bool isDefault()
func (this *QMimeType) IsDefault() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType9isDefaultEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:102
// index:0
// QString name()
func (this *QMimeType) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:103
// index:0
// QString comment()
func (this *QMimeType) Comment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType7commentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:104
// index:0
// QString genericIconName()
func (this *QMimeType) GenericIconName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType15genericIconNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:105
// index:0
// QString iconName()
func (this *QMimeType) IconName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType8iconNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:106
// index:0
// QStringList globPatterns()
func (this *QMimeType) GlobPatterns() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType12globPatternsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:107
// index:0
// QStringList parentMimeTypes()
func (this *QMimeType) ParentMimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType15parentMimeTypesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:108
// index:0
// QStringList allAncestors()
func (this *QMimeType) AllAncestors() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType12allAncestorsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:109
// index:0
// QStringList aliases()
func (this *QMimeType) Aliases() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType7aliasesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:110
// index:0
// QStringList suffixes()
func (this *QMimeType) Suffixes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType8suffixesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:111
// index:0
// QString preferredSuffix()
func (this *QMimeType) PreferredSuffix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType15preferredSuffixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:113
// index:0
// bool inherits(const class QString &)
func (this *QMimeType) Inherits(mimeTypeName unsafe.Pointer) {
	// 0: (, const QString & mimeTypeName), (mimeTypeName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType8inheritsERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, mimeTypeName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimetype.h:115
// index:0
// QString filterString()
func (this *QMimeType) FilterString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeType12filterStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
