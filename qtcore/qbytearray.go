package qtcore

// /usr/include/qt/QtCore/qbytearray.h
// #include <qbytearray.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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
type QByteArray struct {
	*qtrt.CObject
}
type QByteArray_ITF interface {
	QByteArray_PTR() *QByteArray
}

func (ptr *QByteArray) QByteArray_PTR() *QByteArray { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QByteArrayFromptr(cthis Voidptr) *QByteArray {
	return &QByteArray{&qtrt.CObject{cthis}}
}
func (*QByteArray) Fromptr(cthis Voidptr) *QByteArray {
	return QByteArrayFromptr(cthis)
}

// /usr/include/qt/QtCore/qbytearray.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QByteArray(const char *, int)

/*
 */
func (*QByteArray) NewForInherit(arg0 string, size int) *QByteArray {
	return NewQByteArray(arg0, size)
}
func NewQByteArray(arg0 string, size int) *QByteArray {
	var convArg0 = qtrt.CStringRef(&arg0)
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc1(1185669528, "_ZN10QByteArrayC2EPKci", qtrt.FFITY_POINTER, cthis, convArg0, size)
	qtrt.ErrPrint(err, rv)
	gothis := QByteArrayFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQByteArray)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QByteArray(const char *, int)

/*
 */
func (*QByteArray) NewForInheritp(arg0 string) *QByteArray {
	return NewQByteArrayp(arg0)
}
func NewQByteArrayp(arg0 string) *QByteArray {
	var convArg0 = qtrt.CStringRef(&arg0)
	// arg: 1, int=Int, =Invalid, , Invalid
	size := int(-1)
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc1(1185669528, "_ZN10QByteArrayC2EPKci", qtrt.FFITY_POINTER, cthis, convArg0, size)
	qtrt.ErrPrint(err, rv)
	gothis := QByteArrayFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQByteArray)
	return gothis
}

// /usr/include/qt/QtCore/qbytearray.h:196
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int size() const

/*
 */
func (this *QByteArray) Size() int {
	rv, err := qtrt.Qtcc1(2381602376, "_ZNK10QByteArray4sizeEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:210
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] char * data()

/*
 */
func (this *QByteArray) Data() string {
	rv, err := qtrt.Qtcc1(3584956573, "_ZN10QByteArray4dataEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

func DeleteQByteArray(this *QByteArray) {
	rv, err := qtrt.Qtcc1(0, "_ZN10QByteArrayD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QByteArray__Base64Option = int

//
const QByteArray__Base64Encoding QByteArray__Base64Option = 0

//
const QByteArray__Base64UrlEncoding QByteArray__Base64Option = 1

//
const QByteArray__KeepTrailingEquals QByteArray__Base64Option = 0

//
const QByteArray__OmitTrailingEquals QByteArray__Base64Option = 2

//
const QByteArray__IgnoreBase64DecodingErrors QByteArray__Base64Option = 0

//
const QByteArray__AbortOnBase64DecodingErrors QByteArray__Base64Option = 4

func (this *QByteArray) Base64OptionItemName(val int) string {
	switch val {
	case QByteArray__Base64Encoding: // 0
		return "Base64Encoding,KeepTrailingEquals,IgnoreBase64DecodingErrors"
	case QByteArray__Base64UrlEncoding: // 1
		return "Base64UrlEncoding"
		// case QByteArray__KeepTrailingEquals: // 0
		// return ""
	case QByteArray__OmitTrailingEquals: // 2
		return "OmitTrailingEquals"
		// case QByteArray__IgnoreBase64DecodingErrors: // 0
		// return ""
	case QByteArray__AbortOnBase64DecodingErrors: // 4
		return "AbortOnBase64DecodingErrors"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QByteArray_Base64OptionItemName(val int) string {
	var nilthis *QByteArray
	return nilthis.Base64OptionItemName(val)
}

/*


 */
type QByteArray__Base64DecodingStatus = int

//
const QByteArray__Ok QByteArray__Base64DecodingStatus = 0

//
const QByteArray__IllegalInputLength QByteArray__Base64DecodingStatus = 1

//
const QByteArray__IllegalCharacter QByteArray__Base64DecodingStatus = 2

//
const QByteArray__IllegalPadding QByteArray__Base64DecodingStatus = 3

func (this *QByteArray) Base64DecodingStatusItemName(val int) string {
	switch val {
	case QByteArray__Ok: // 0
		return "Ok"
	case QByteArray__IllegalInputLength: // 1
		return "IllegalInputLength"
	case QByteArray__IllegalCharacter: // 2
		return "IllegalCharacter"
	case QByteArray__IllegalPadding: // 3
		return "IllegalPadding"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QByteArray_Base64DecodingStatusItemName(val int) string {
	var nilthis *QByteArray
	return nilthis.Base64DecodingStatusItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10007() {
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
