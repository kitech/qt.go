package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 61
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QTextBlockFormat struct {
	*QTextFormat
}

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
func NewQTextBlockFormat() *QTextBlockFormat {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormatC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBlockFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:649
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextBlockFormat(const QTextFormat &)
func NewQTextBlockFormat_1(fmt *QTextFormat) *QTextBlockFormat {
	var convArg0 = fmt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextBlockFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:592
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextBlockFormat) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:595
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Alignment alignment()
func (this *QTextBlockFormat) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:598
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopMargin(qreal)
func (this *QTextBlockFormat) SetTopMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat12setTopMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:600
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal topMargin()
func (this *QTextBlockFormat) TopMargin() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat9topMarginEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:603
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomMargin(qreal)
func (this *QTextBlockFormat) SetBottomMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat15setBottomMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:605
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal bottomMargin()
func (this *QTextBlockFormat) BottomMargin() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat12bottomMarginEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:608
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLeftMargin(qreal)
func (this *QTextBlockFormat) SetLeftMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat13setLeftMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:610
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal leftMargin()
func (this *QTextBlockFormat) LeftMargin() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat10leftMarginEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:613
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRightMargin(qreal)
func (this *QTextBlockFormat) SetRightMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat14setRightMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:615
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal rightMargin()
func (this *QTextBlockFormat) RightMargin() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat11rightMarginEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:618
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextIndent(qreal)
func (this *QTextBlockFormat) SetTextIndent(aindent float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat13setTextIndentEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), aindent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:620
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal textIndent()
func (this *QTextBlockFormat) TextIndent() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat10textIndentEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:623
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIndent(int)
func (this *QTextBlockFormat) SetIndent(indent int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat9setIndentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), indent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:624
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indent()
func (this *QTextBlockFormat) Indent() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat6indentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:627
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLineHeight(qreal, int)
func (this *QTextBlockFormat) SetLineHeight(height float64, heightType int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat13setLineHeightEdi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), height, heightType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:629
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal lineHeight(qreal, qreal)
func (this *QTextBlockFormat) LineHeight(scriptLineHeight float64, scaling float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat10lineHeightEdd", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), scriptLineHeight, scaling)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:630
// index:1
// Public inline Visibility=Default Availability=Available
// [8] qreal lineHeight()
func (this *QTextBlockFormat) LineHeight_1() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat10lineHeightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:632
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lineHeightType()
func (this *QTextBlockFormat) LineHeightType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat14lineHeightTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:635
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setNonBreakableLines(_Bool)
func (this *QTextBlockFormat) SetNonBreakableLines(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat20setNonBreakableLinesEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:637
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool nonBreakableLines()
func (this *QTextBlockFormat) NonBreakableLines() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat17nonBreakableLinesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:640
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setPageBreakPolicy(QTextFormat::PageBreakFlags)
func (this *QTextBlockFormat) SetPageBreakPolicy(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat18setPageBreakPolicyE6QFlagsIN11QTextFormat13PageBreakFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:642
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextFormat::PageBreakFlags pageBreakPolicy()
func (this *QTextBlockFormat) PageBreakPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat15pageBreakPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

func DeleteQTextBlockFormat(this *QTextBlockFormat) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormatD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QTextBlockFormat__LineHeightTypes = int

const QTextBlockFormat__SingleHeight QTextBlockFormat__LineHeightTypes = 0
const QTextBlockFormat__ProportionalHeight QTextBlockFormat__LineHeightTypes = 1
const QTextBlockFormat__FixedHeight QTextBlockFormat__LineHeightTypes = 2
const QTextBlockFormat__MinimumHeight QTextBlockFormat__LineHeightTypes = 3
const QTextBlockFormat__LineDistanceHeight QTextBlockFormat__LineHeightTypes = 4

//  body block end
