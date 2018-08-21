package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QTextLength struct {
	*qtrt.CObject
}
type QTextLength_ITF interface {
	QTextLength_PTR() *QTextLength
}

func (ptr *QTextLength) QTextLength_PTR() *QTextLength { return ptr }

func (this *QTextLength) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextLength) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextLengthFromPointer(cthis unsafe.Pointer) *QTextLength {
	return &QTextLength{&qtrt.CObject{cthis}}
}
func (*QTextLength) NewFromPointer(cthis unsafe.Pointer) *QTextLength {
	return NewQTextLengthFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextLength()

/*

 */
func NewQTextLength() *QTextLength {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLengthC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLengthFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLength)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:91
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QTextLength(QTextLength::Type, qreal)

/*

 */
func NewQTextLength_1(type_ int, value float64) *QTextLength {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLengthC2ENS_4TypeEd", qtrt.FFI_TYPE_POINTER, type_, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLengthFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLength)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextLength::Type type() const

/*
Returns the type of this format.

See also FormatType.
*/
func (this *QTextLength) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLength4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal value(qreal) const

/*

 */
func (this *QTextLength) Value(maximumLength float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLength5valueEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), maximumLength)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:104
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal rawValue() const

/*

 */
func (this *QTextLength) RawValue() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLength8rawValueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QTextLength &) const

/*

 */
func (this *QTextLength) Operator_equal_equal(other QTextLength_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextLength_PTR() != nil {
		convArg0 = other.QTextLength_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLengtheqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QTextLength &) const

/*

 */
func (this *QTextLength) Operator_not_equal(other QTextLength_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextLength_PTR() != nil {
		convArg0 = other.QTextLength_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLengthneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQTextLength(this *QTextLength) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLengthD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QTextLength__Type = int

//
const QTextLength__VariableLength QTextLength__Type = 0

//
const QTextLength__FixedLength QTextLength__Type = 1

//
const QTextLength__PercentageLength QTextLength__Type = 2

func (this *QTextLength) TypeItemName(val int) string {
	switch val {
	case QTextLength__VariableLength: // 0
		return "VariableLength"
	case QTextLength__FixedLength: // 1
		return "FixedLength"
	case QTextLength__PercentageLength: // 2
		return "PercentageLength"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextLength_TypeItemName(val int) string {
	var nilthis *QTextLength
	return nilthis.TypeItemName(val)
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
