//  header block begin
// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 62
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.QTextFormat.GetCthis()
}

// /usr/include/qt/QtGui/qtextformat.h:590
// index:0
// void QTextBlockFormat()
func NewQTextBlockFormat() *QTextBlockFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBlockFormatFromPointer(cthis)
	return gothis
}
func NewQTextBlockFormatFromPointer(cthis unsafe.Pointer) *QTextBlockFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextBlockFormat{bcthis0}
}

// /usr/include/qt/QtGui/qtextformat.h:649
// index:1
// void QTextBlockFormat(const class QTextFormat &)
func NewQTextBlockFormat_1(fmt unsafe.Pointer) *QTextBlockFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_VOID, cthis, fmt)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBlockFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:592
// index:0
// inline
// bool isValid()
func (this *QTextBlockFormat) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:594
// index:0
// inline
// void setAlignment(Qt::Alignment)
func (this *QTextBlockFormat) SetAlignment(alignment int) {
	// 0: (, QFlags<Qt::AlignmentFlag> alignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:595
// index:0
// inline
// Qt::Alignment alignment()
func (this *QTextBlockFormat) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat9alignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:598
// index:0
// inline
// void setTopMargin(qreal)
func (this *QTextBlockFormat) SetTopMargin(margin float64) {
	// 0: (, margin qreal), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat12setTopMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:600
// index:0
// inline
// qreal topMargin()
func (this *QTextBlockFormat) TopMargin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat9topMarginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:603
// index:0
// inline
// void setBottomMargin(qreal)
func (this *QTextBlockFormat) SetBottomMargin(margin float64) {
	// 0: (, margin qreal), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat15setBottomMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:605
// index:0
// inline
// qreal bottomMargin()
func (this *QTextBlockFormat) BottomMargin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat12bottomMarginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:608
// index:0
// inline
// void setLeftMargin(qreal)
func (this *QTextBlockFormat) SetLeftMargin(margin float64) {
	// 0: (, margin qreal), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat13setLeftMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:610
// index:0
// inline
// qreal leftMargin()
func (this *QTextBlockFormat) LeftMargin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat10leftMarginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:613
// index:0
// inline
// void setRightMargin(qreal)
func (this *QTextBlockFormat) SetRightMargin(margin float64) {
	// 0: (, margin qreal), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat14setRightMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:615
// index:0
// inline
// qreal rightMargin()
func (this *QTextBlockFormat) RightMargin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat11rightMarginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:618
// index:0
// inline
// void setTextIndent(qreal)
func (this *QTextBlockFormat) SetTextIndent(aindent float64) {
	// 0: (, aindent qreal), (&aindent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat13setTextIndentEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &aindent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:620
// index:0
// inline
// qreal textIndent()
func (this *QTextBlockFormat) TextIndent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat10textIndentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:623
// index:0
// inline
// void setIndent(int)
func (this *QTextBlockFormat) SetIndent(indent int) {
	// 0: (, indent int), (&indent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat9setIndentEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &indent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:624
// index:0
// inline
// int indent()
func (this *QTextBlockFormat) Indent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat6indentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:627
// index:0
// inline
// void setLineHeight(qreal, int)
func (this *QTextBlockFormat) SetLineHeight(height float64, heightType int) {
	// 0: (, height qreal, heightType int), (&height, &heightType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat13setLineHeightEdi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &height, &heightType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:629
// index:0
// inline
// qreal lineHeight(qreal, qreal)
func (this *QTextBlockFormat) LineHeight(scriptLineHeight float64, scaling float64) {
	// 0: (, scriptLineHeight qreal, scaling qreal), (&scriptLineHeight, &scaling)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat10lineHeightEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &scriptLineHeight, &scaling)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:630
// index:1
// inline
// qreal lineHeight()
func (this *QTextBlockFormat) LineHeight_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat10lineHeightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:632
// index:0
// inline
// int lineHeightType()
func (this *QTextBlockFormat) LineHeightType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat14lineHeightTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:635
// index:0
// inline
// void setNonBreakableLines(_Bool)
func (this *QTextBlockFormat) SetNonBreakableLines(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat20setNonBreakableLinesEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:637
// index:0
// inline
// bool nonBreakableLines()
func (this *QTextBlockFormat) NonBreakableLines() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat17nonBreakableLinesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:640
// index:0
// inline
// void setPageBreakPolicy(QTextFormat::PageBreakFlags)
func (this *QTextBlockFormat) SetPageBreakPolicy(flags int) {
	// 0: (, QFlags<QTextFormat::PageBreakFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextBlockFormat18setPageBreakPolicyE6QFlagsIN11QTextFormat13PageBreakFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:642
// index:0
// inline
// QTextFormat::PageBreakFlags pageBreakPolicy()
func (this *QTextBlockFormat) PageBreakPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat15pageBreakPolicyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:646
// index:0
// QList<QTextOption::Tab> tabPositions()
func (this *QTextBlockFormat) TabPositions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextBlockFormat12tabPositionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
