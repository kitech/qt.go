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
// extern C begin: 62
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
type QTextBlockFormat struct {
	*QTextFormat
}
type QTextBlockFormat_ITF interface {
	QTextFormat_ITF
	QTextBlockFormat_PTR() *QTextBlockFormat
}

func (ptr *QTextBlockFormat) QTextBlockFormat_PTR() *QTextBlockFormat { return ptr }

func (this *QTextBlockFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextFormat.GetCthis()
	}
}
func (this *QTextBlockFormat) SetCthis(cthis unsafe.Pointer) {
	this.QTextFormat = NewQTextFormatFromPointer(cthis)
}
func NewQTextBlockFormatFromPointer(cthis unsafe.Pointer) *QTextBlockFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextBlockFormat{bcthis0}
}
func (*QTextBlockFormat) NewFromPointer(cthis unsafe.Pointer) *QTextBlockFormat {
	return NewQTextBlockFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:590
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextBlockFormat()

/*

 */
func NewQTextBlockFormat() *QTextBlockFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBlockFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:649
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextBlockFormat(const QTextFormat &)

/*

 */
func NewQTextBlockFormat_1(fmt_ QTextFormat_ITF) *QTextBlockFormat {
	var convArg0 unsafe.Pointer
	if fmt_ != nil && fmt_.QTextFormat_PTR() != nil {
		convArg0 = fmt_.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormatC2ERK11QTextFormat", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBlockFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:592
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the format is valid (i.e. is not InvalidFormat); otherwise returns false.
*/
func (this *QTextBlockFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:594
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)

/*

 */
func (this *QTextBlockFormat) SetAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormat12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:595
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const

/*

 */
func (this *QTextBlockFormat) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:598
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopMargin(qreal)

/*

 */
func (this *QTextBlockFormat) SetTopMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormat12setTopMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:600
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal topMargin() const

/*

 */
func (this *QTextBlockFormat) TopMargin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat9topMarginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:603
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomMargin(qreal)

/*

 */
func (this *QTextBlockFormat) SetBottomMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormat15setBottomMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:605
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal bottomMargin() const

/*

 */
func (this *QTextBlockFormat) BottomMargin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat12bottomMarginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:608
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLeftMargin(qreal)

/*

 */
func (this *QTextBlockFormat) SetLeftMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormat13setLeftMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:610
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal leftMargin() const

/*

 */
func (this *QTextBlockFormat) LeftMargin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat10leftMarginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:613
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRightMargin(qreal)

/*

 */
func (this *QTextBlockFormat) SetRightMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormat14setRightMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:615
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal rightMargin() const

/*

 */
func (this *QTextBlockFormat) RightMargin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat11rightMarginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:618
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextIndent(qreal)

/*

 */
func (this *QTextBlockFormat) SetTextIndent(aindent float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormat13setTextIndentEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), aindent)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:620
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal textIndent() const

/*

 */
func (this *QTextBlockFormat) TextIndent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat10textIndentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:623
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIndent(int)

/*

 */
func (this *QTextBlockFormat) SetIndent(indent int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormat9setIndentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), indent)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:624
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indent() const

/*

 */
func (this *QTextBlockFormat) Indent() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat6indentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:627
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLineHeight(qreal, int)

/*

 */
func (this *QTextBlockFormat) SetLineHeight(height float64, heightType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormat13setLineHeightEdi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), height, heightType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:629
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal lineHeight(qreal, qreal) const

/*

 */
func (this *QTextBlockFormat) LineHeight(scriptLineHeight float64, scaling float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat10lineHeightEdd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), scriptLineHeight, scaling)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:630
// index:1
// Public inline Visibility=Default Availability=Available
// [8] qreal lineHeight() const

/*

 */
func (this *QTextBlockFormat) LineHeight_1() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat10lineHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:632
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lineHeightType() const

/*

 */
func (this *QTextBlockFormat) LineHeightType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat14lineHeightTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:635
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setNonBreakableLines(bool)

/*

 */
func (this *QTextBlockFormat) SetNonBreakableLines(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormat20setNonBreakableLinesEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:637
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool nonBreakableLines() const

/*

 */
func (this *QTextBlockFormat) NonBreakableLines() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat17nonBreakableLinesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:640
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setPageBreakPolicy(QTextFormat::PageBreakFlags)

/*

 */
func (this *QTextBlockFormat) SetPageBreakPolicy(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormat18setPageBreakPolicyE6QFlagsIN11QTextFormat13PageBreakFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:642
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextFormat::PageBreakFlags pageBreakPolicy() const

/*

 */
func (this *QTextBlockFormat) PageBreakPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextBlockFormat15pageBreakPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

func DeleteQTextBlockFormat(this *QTextBlockFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextBlockFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QTextBlockFormat__LineHeightTypes = int

//
const QTextBlockFormat__SingleHeight QTextBlockFormat__LineHeightTypes = 0

//
const QTextBlockFormat__ProportionalHeight QTextBlockFormat__LineHeightTypes = 1

//
const QTextBlockFormat__FixedHeight QTextBlockFormat__LineHeightTypes = 2

//
const QTextBlockFormat__MinimumHeight QTextBlockFormat__LineHeightTypes = 3

//
const QTextBlockFormat__LineDistanceHeight QTextBlockFormat__LineHeightTypes = 4

func (this *QTextBlockFormat) LineHeightTypesItemName(val int) string {
	switch val {
	case QTextBlockFormat__SingleHeight: // 0
		return "SingleHeight"
	case QTextBlockFormat__ProportionalHeight: // 1
		return "ProportionalHeight"
	case QTextBlockFormat__FixedHeight: // 2
		return "FixedHeight"
	case QTextBlockFormat__MinimumHeight: // 3
		return "MinimumHeight"
	case QTextBlockFormat__LineDistanceHeight: // 4
		return "LineDistanceHeight"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextBlockFormat_LineHeightTypesItemName(val int) string {
	var nilthis *QTextBlockFormat
	return nilthis.LineHeightTypesItemName(val)
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
