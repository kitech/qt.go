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
// extern C begin: 27
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
type QTextListFormat struct {
	*QTextFormat
}
type QTextListFormat_ITF interface {
	QTextFormat_ITF
	QTextListFormat_PTR() *QTextListFormat
}

func (ptr *QTextListFormat) QTextListFormat_PTR() *QTextListFormat { return ptr }

func (this *QTextListFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextFormat.GetCthis()
	}
}
func (this *QTextListFormat) SetCthis(cthis unsafe.Pointer) {
	this.QTextFormat = NewQTextFormatFromPointer(cthis)
}
func NewQTextListFormatFromPointer(cthis unsafe.Pointer) *QTextListFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextListFormat{bcthis0}
}
func (*QTextListFormat) NewFromPointer(cthis unsafe.Pointer) *QTextListFormat {
	return NewQTextListFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:688
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextListFormat()

/*

 */
func (*QTextListFormat) NewForInherit() *QTextListFormat {
	return NewQTextListFormat()
}
func NewQTextListFormat() *QTextListFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextListFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextListFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:721
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextListFormat(const QTextFormat &)

/*

 */
func (*QTextListFormat) NewForInherit1(fmt_ QTextFormat_ITF) *QTextListFormat {
	return NewQTextListFormat1(fmt_)
}
func NewQTextListFormat1(fmt_ QTextFormat_ITF) *QTextListFormat {
	var convArg0 unsafe.Pointer
	if fmt_ != nil && fmt_.QTextFormat_PTR() != nil {
		convArg0 = fmt_.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormatC2ERK11QTextFormat", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextListFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextListFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:690
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the format is valid (i.e. is not InvalidFormat); otherwise returns false.
*/
func (this *QTextListFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextListFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:704
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setStyle(QTextListFormat::Style)

/*

 */
func (this *QTextListFormat) SetStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormat8setStyleENS_5StyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:705
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextListFormat::Style style() const

/*

 */
func (this *QTextListFormat) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextListFormat5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:708
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIndent(int)

/*

 */
func (this *QTextListFormat) SetIndent(indent int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormat9setIndentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), indent)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:709
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indent() const

/*

 */
func (this *QTextListFormat) Indent() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextListFormat6indentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:712
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setNumberPrefix(const QString &)

/*

 */
func (this *QTextListFormat) SetNumberPrefix(numberPrefix string) {
	var tmpArg0 = qtcore.NewQString5(numberPrefix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormat15setNumberPrefixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:713
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString numberPrefix() const

/*

 */
func (this *QTextListFormat) NumberPrefix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextListFormat12numberPrefixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextformat.h:716
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setNumberSuffix(const QString &)

/*

 */
func (this *QTextListFormat) SetNumberSuffix(numberSuffix string) {
	var tmpArg0 = qtcore.NewQString5(numberSuffix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormat15setNumberSuffixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:717
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString numberSuffix() const

/*

 */
func (this *QTextListFormat) NumberSuffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextListFormat12numberSuffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

func DeleteQTextListFormat(this *QTextListFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QTextListFormat__Style = int

//
const QTextListFormat__ListDisc QTextListFormat__Style = -1

//
const QTextListFormat__ListCircle QTextListFormat__Style = -2

//
const QTextListFormat__ListSquare QTextListFormat__Style = -3

//
const QTextListFormat__ListDecimal QTextListFormat__Style = -4

//
const QTextListFormat__ListLowerAlpha QTextListFormat__Style = -5

//
const QTextListFormat__ListUpperAlpha QTextListFormat__Style = -6

//
const QTextListFormat__ListLowerRoman QTextListFormat__Style = -7

//
const QTextListFormat__ListUpperRoman QTextListFormat__Style = -8

//
const QTextListFormat__ListStyleUndefined QTextListFormat__Style = 0

func (this *QTextListFormat) StyleItemName(val int) string {
	switch val {
	case QTextListFormat__ListDisc: // -1
		return "ListDisc"
	case QTextListFormat__ListCircle: // -2
		return "ListCircle"
	case QTextListFormat__ListSquare: // -3
		return "ListSquare"
	case QTextListFormat__ListDecimal: // -4
		return "ListDecimal"
	case QTextListFormat__ListLowerAlpha: // -5
		return "ListLowerAlpha"
	case QTextListFormat__ListUpperAlpha: // -6
		return "ListUpperAlpha"
	case QTextListFormat__ListLowerRoman: // -7
		return "ListLowerRoman"
	case QTextListFormat__ListUpperRoman: // -8
		return "ListUpperRoman"
	case QTextListFormat__ListStyleUndefined: // 0
		return "ListStyleUndefined"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextListFormat_StyleItemName(val int) string {
	var nilthis *QTextListFormat
	return nilthis.StyleItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10767() {
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
