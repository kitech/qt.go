//  header block begin
// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QTextFrameFormat struct {
	*QTextFormat
}

func (this *QTextFrameFormat) GetCthis() unsafe.Pointer {
	return this.QTextFormat.GetCthis()
}

// /usr/include/qt/QtGui/qtextformat.h:770
// index:0
// void QTextFrameFormat()
func NewQTextFrameFormat() *QTextFrameFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFrameFormatFromPointer(cthis)
	return gothis
}
func NewQTextFrameFormatFromPointer(cthis unsafe.Pointer) *QTextFrameFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextFrameFormat{bcthis0}
}

// /usr/include/qt/QtGui/qtextformat.h:852
// index:1
// void QTextFrameFormat(const class QTextFormat &)
func NewQTextFrameFormat_1(fmt unsafe.Pointer) *QTextFrameFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_VOID, cthis, fmt)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFrameFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:772
// index:0
// inline
// bool isValid()
func (this *QTextFrameFormat) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:796
// index:0
// inline
// void setPosition(enum QTextFrameFormat::Position)
func (this *QTextFrameFormat) SetPosition(f int) {
	// 0: (, f QTextFrameFormat::Position), (&f)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat11setPositionENS_8PositionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:798
// index:0
// inline
// QTextFrameFormat::Position position()
func (this *QTextFrameFormat) Position() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat8positionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:801
// index:0
// inline
// void setBorder(qreal)
func (this *QTextFrameFormat) SetBorder(border float64) {
	// 0: (, border qreal), (&border)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat9setBorderEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &border)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:802
// index:0
// inline
// qreal border()
func (this *QTextFrameFormat) Border() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat6borderEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:805
// index:0
// inline
// void setBorderBrush(const class QBrush &)
func (this *QTextFrameFormat) SetBorderBrush(brush unsafe.Pointer) {
	// 0: (, brush const QBrush &), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat14setBorderBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:807
// index:0
// inline
// QBrush borderBrush()
func (this *QTextFrameFormat) BorderBrush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat11borderBrushEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:810
// index:0
// inline
// void setBorderStyle(enum QTextFrameFormat::BorderStyle)
func (this *QTextFrameFormat) SetBorderStyle(style int) {
	// 0: (, style QTextFrameFormat::BorderStyle), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat14setBorderStyleENS_11BorderStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:812
// index:0
// inline
// QTextFrameFormat::BorderStyle borderStyle()
func (this *QTextFrameFormat) BorderStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat11borderStyleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:815
// index:0
// void setMargin(qreal)
func (this *QTextFrameFormat) SetMargin(margin float64) {
	// 0: (, margin qreal), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat9setMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:816
// index:0
// inline
// qreal margin()
func (this *QTextFrameFormat) Margin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat6marginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:819
// index:0
// inline
// void setTopMargin(qreal)
func (this *QTextFrameFormat) SetTopMargin(margin float64) {
	// 0: (, margin qreal), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat12setTopMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:820
// index:0
// qreal topMargin()
func (this *QTextFrameFormat) TopMargin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat9topMarginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:822
// index:0
// inline
// void setBottomMargin(qreal)
func (this *QTextFrameFormat) SetBottomMargin(margin float64) {
	// 0: (, margin qreal), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat15setBottomMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:823
// index:0
// qreal bottomMargin()
func (this *QTextFrameFormat) BottomMargin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat12bottomMarginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:825
// index:0
// inline
// void setLeftMargin(qreal)
func (this *QTextFrameFormat) SetLeftMargin(margin float64) {
	// 0: (, margin qreal), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat13setLeftMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:826
// index:0
// qreal leftMargin()
func (this *QTextFrameFormat) LeftMargin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat10leftMarginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:828
// index:0
// inline
// void setRightMargin(qreal)
func (this *QTextFrameFormat) SetRightMargin(margin float64) {
	// 0: (, margin qreal), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat14setRightMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:829
// index:0
// qreal rightMargin()
func (this *QTextFrameFormat) RightMargin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat11rightMarginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:831
// index:0
// inline
// void setPadding(qreal)
func (this *QTextFrameFormat) SetPadding(padding float64) {
	// 0: (, padding qreal), (&padding)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat10setPaddingEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:832
// index:0
// inline
// qreal padding()
func (this *QTextFrameFormat) Padding() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat7paddingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:835
// index:0
// inline
// void setWidth(qreal)
func (this *QTextFrameFormat) SetWidth(width float64) {
	// 0: (, width qreal), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat8setWidthEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:836
// index:1
// inline
// void setWidth(const class QTextLength &)
func (this *QTextFrameFormat) SetWidth_1(length unsafe.Pointer) {
	// 1: (, length const QTextLength &), (length)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat8setWidthERK11QTextLength", ffiqt.FFI_TYPE_VOID, this.GetCthis(), length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:838
// index:0
// inline
// QTextLength width()
func (this *QTextFrameFormat) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat5widthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:841
// index:0
// inline
// void setHeight(qreal)
func (this *QTextFrameFormat) SetHeight(height float64) {
	// 0: (, height qreal), (&height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat9setHeightEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:842
// index:1
// inline
// void setHeight(const class QTextLength &)
func (this *QTextFrameFormat) SetHeight_1(height unsafe.Pointer) {
	// 1: (, height const QTextLength &), (height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat9setHeightERK11QTextLength", ffiqt.FFI_TYPE_VOID, this.GetCthis(), height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:843
// index:0
// inline
// QTextLength height()
func (this *QTextFrameFormat) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat6heightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:846
// index:0
// inline
// void setPageBreakPolicy(QTextFormat::PageBreakFlags)
func (this *QTextFrameFormat) SetPageBreakPolicy(flags int) {
	// 0: (, QFlags<QTextFormat::PageBreakFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextFrameFormat18setPageBreakPolicyE6QFlagsIN11QTextFormat13PageBreakFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:848
// index:0
// inline
// QTextFormat::PageBreakFlags pageBreakPolicy()
func (this *QTextFrameFormat) PageBreakPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextFrameFormat15pageBreakPolicyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
