package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
// size 8
type QString struct {
	*qtrt.CObject
}
type QString_ITF interface {
	QString_PTR() *QString
}

func (ptr *QString) QString_PTR() *QString { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QStringFromptr(cthis Voidptr) *QString {
	return &QString{&qtrt.CObject{cthis}}
}
func (*QString) Fromptr(cthis Voidptr) *QString {
	return QStringFromptr(cthis)
}

// /usr/include/qt/QtCore/qstring.h:279
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int length() const

/*
 */
func (this *QString) Length() int {
	rv, err := qtrt.Qtcc3(1867525835, "_ZNK7QString6lengthEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:683
// index:0
// Public inline Indirect Visibility=Default Availability=Available
// [8] QByteArray toUtf8() const

/*
 */
func (this *QString) ToUtf8() *QByteArray /*123*/ {
	sretobj := qtrt.Malloc(8) // QByteArray
	rv, err := qtrt.Qtcc3(637795128, "_ZNKR7QString6toUtf8Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QByteArrayFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:704
// index:0
// Public static inline Indirect Visibility=Default Availability=Available
// [8] QString fromUtf8(const char *, int)

/*
 */
func (this *QString) FromUtf8(str string, size int) string {
	var convArg0 = qtrt.CStringRef(&str)
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2549669065, "_ZN7QString8fromUtf8EPKci", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), Voidptr(&convArg0), Voidptr(&size))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QString_FromUtf8(str string, size int) string {
	var nilthis *QString
	rv := nilthis.FromUtf8(str, size)
	return rv
}

// /usr/include/qt/QtCore/qstring.h:704
// index:0
// Public static inline Indirect Visibility=Default Availability=Available
// [8] QString fromUtf8(const char *, int)

/*
 */
func (this *QString) FromUtf8p(str string) string {
	var convArg0 = qtrt.CStringRef(&str)
	// arg: 1, int=Int, =Invalid, , Invalid
	size := int(-1)
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2549669065, "_ZN7QString8fromUtf8EPKci", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), Voidptr(&convArg0), Voidptr(&size))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:835
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QString(const char *)

/*
 */
func (*QString) NewForInherit(ch string) *QString {
	return NewQString(ch)
}
func NewQString(ch string) *QString {
	var convArg0 = qtrt.CStringRef(&ch)
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc3(2640353187, "_ZN7QStringC2EPKc", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QStringFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQString)
	return gothis
}

func DeleteQString(this *QString) {
	rv, err := qtrt.Qtcc3(2000556505, "_ZN7QStringD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QString__SectionFlag = int

//
const QString__SectionDefault QString__SectionFlag = 0

//
const QString__SectionSkipEmpty QString__SectionFlag = 1

//
const QString__SectionIncludeLeadingSep QString__SectionFlag = 2

//
const QString__SectionIncludeTrailingSep QString__SectionFlag = 4

//
const QString__SectionCaseInsensitiveSeps QString__SectionFlag = 8

func (this *QString) SectionFlagItemName(val int) string {
	switch val {
	case QString__SectionDefault: // 0
		return "SectionDefault"
	case QString__SectionSkipEmpty: // 1
		return "SectionSkipEmpty"
	case QString__SectionIncludeLeadingSep: // 2
		return "SectionIncludeLeadingSep"
	case QString__SectionIncludeTrailingSep: // 4
		return "SectionIncludeTrailingSep"
	case QString__SectionCaseInsensitiveSeps: // 8
		return "SectionCaseInsensitiveSeps"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QString_SectionFlagItemName(val int) string {
	var nilthis *QString
	return nilthis.SectionFlagItemName(val)
}

/*


 */
type QString__SplitBehavior = int

//
const QString__KeepEmptyParts QString__SplitBehavior = 0

//
const QString__SkipEmptyParts QString__SplitBehavior = 1

func (this *QString) SplitBehaviorItemName(val int) string {
	switch val {
	case QString__KeepEmptyParts: // 0
		return "KeepEmptyParts"
	case QString__SkipEmptyParts: // 1
		return "SkipEmptyParts"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QString_SplitBehaviorItemName(val int) string {
	var nilthis *QString
	return nilthis.SplitBehaviorItemName(val)
}

/*


 */
type QString__NormalizationForm = int

//
const QString__NormalizationForm_D QString__NormalizationForm = 0

//
const QString__NormalizationForm_C QString__NormalizationForm = 1

//
const QString__NormalizationForm_KD QString__NormalizationForm = 2

//
const QString__NormalizationForm_KC QString__NormalizationForm = 3

func (this *QString) NormalizationFormItemName(val int) string {
	switch val {
	case QString__NormalizationForm_D: // 0
		return "NormalizationForm_D"
	case QString__NormalizationForm_C: // 1
		return "NormalizationForm_C"
	case QString__NormalizationForm_KD: // 2
		return "NormalizationForm_KD"
	case QString__NormalizationForm_KC: // 3
		return "NormalizationForm_KC"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QString_NormalizationFormItemName(val int) string {
	var nilthis *QString
	return nilthis.NormalizationFormItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10009() {
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
