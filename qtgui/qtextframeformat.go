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
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QTextFrameFormat struct {
	*QTextFormat
}
type QTextFrameFormat_ITF interface {
	QTextFormat_ITF
	QTextFrameFormat_PTR() *QTextFrameFormat
}

func (ptr *QTextFrameFormat) QTextFrameFormat_PTR() *QTextFrameFormat { return ptr }

func (this *QTextFrameFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextFormat.GetCthis()
	}
}
func (this *QTextFrameFormat) SetCthis(cthis unsafe.Pointer) {
	this.QTextFormat = NewQTextFormatFromPointer(cthis)
}
func NewQTextFrameFormatFromPointer(cthis unsafe.Pointer) *QTextFrameFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextFrameFormat{bcthis0}
}
func (*QTextFrameFormat) NewFromPointer(cthis unsafe.Pointer) *QTextFrameFormat {
	return NewQTextFrameFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:770
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextFrameFormat()
func NewQTextFrameFormat() *QTextFrameFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextFrameFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextFrameFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:852
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextFrameFormat(const QTextFormat &)
func NewQTextFrameFormat_1(fmt QTextFormat_ITF) *QTextFrameFormat {
	var convArg0 = fmt.QTextFormat_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormatC2ERK11QTextFormat", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextFrameFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextFrameFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:772
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextFrameFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:796
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setPosition(enum QTextFrameFormat::Position)
func (this *QTextFrameFormat) SetPosition(f int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat11setPositionENS_8PositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:798
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextFrameFormat::Position position()
func (this *QTextFrameFormat) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:801
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBorder(qreal)
func (this *QTextFrameFormat) SetBorder(border float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat9setBorderEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), border)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:802
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal border()
func (this *QTextFrameFormat) Border() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat6borderEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:805
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBorderBrush(const QBrush &)
func (this *QTextFrameFormat) SetBorderBrush(brush QBrush_ITF) {
	var convArg0 = brush.QBrush_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat14setBorderBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:807
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush borderBrush()
func (this *QTextFrameFormat) BorderBrush() *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat11borderBrushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:810
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBorderStyle(enum QTextFrameFormat::BorderStyle)
func (this *QTextFrameFormat) SetBorderStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat14setBorderStyleENS_11BorderStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:812
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextFrameFormat::BorderStyle borderStyle()
func (this *QTextFrameFormat) BorderStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat11borderStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:815
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMargin(qreal)
func (this *QTextFrameFormat) SetMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat9setMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:816
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal margin()
func (this *QTextFrameFormat) Margin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat6marginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:819
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopMargin(qreal)
func (this *QTextFrameFormat) SetTopMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat12setTopMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:820
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal topMargin()
func (this *QTextFrameFormat) TopMargin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat9topMarginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:822
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomMargin(qreal)
func (this *QTextFrameFormat) SetBottomMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat15setBottomMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:823
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal bottomMargin()
func (this *QTextFrameFormat) BottomMargin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat12bottomMarginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:825
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLeftMargin(qreal)
func (this *QTextFrameFormat) SetLeftMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat13setLeftMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:826
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal leftMargin()
func (this *QTextFrameFormat) LeftMargin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat10leftMarginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:828
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRightMargin(qreal)
func (this *QTextFrameFormat) SetRightMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat14setRightMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:829
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rightMargin()
func (this *QTextFrameFormat) RightMargin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat11rightMarginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:831
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setPadding(qreal)
func (this *QTextFrameFormat) SetPadding(padding float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat10setPaddingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:832
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal padding()
func (this *QTextFrameFormat) Padding() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat7paddingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:835
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(qreal)
func (this *QTextFrameFormat) SetWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:836
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setWidth(const QTextLength &)
func (this *QTextFrameFormat) SetWidth_1(length QTextLength_ITF) {
	var convArg0 = length.QTextLength_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat8setWidthERK11QTextLength", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:838
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QTextLength width()
func (this *QTextFrameFormat) Width() *QTextLength /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextLengthFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLength)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:841
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(qreal)
func (this *QTextFrameFormat) SetHeight(height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat9setHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:842
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setHeight(const QTextLength &)
func (this *QTextFrameFormat) SetHeight_1(height QTextLength_ITF) {
	var convArg0 = height.QTextLength_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat9setHeightERK11QTextLength", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:843
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QTextLength height()
func (this *QTextFrameFormat) Height() *QTextLength /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextLengthFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLength)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:846
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setPageBreakPolicy(QTextFormat::PageBreakFlags)
func (this *QTextFrameFormat) SetPageBreakPolicy(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormat18setPageBreakPolicyE6QFlagsIN11QTextFormat13PageBreakFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:848
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextFormat::PageBreakFlags pageBreakPolicy()
func (this *QTextFrameFormat) PageBreakPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextFrameFormat15pageBreakPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

func DeleteQTextFrameFormat(this *QTextFrameFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextFrameFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QTextFrameFormat__Position = int

const QTextFrameFormat__InFlow QTextFrameFormat__Position = 0
const QTextFrameFormat__FloatLeft QTextFrameFormat__Position = 1
const QTextFrameFormat__FloatRight QTextFrameFormat__Position = 2

type QTextFrameFormat__BorderStyle = int

const QTextFrameFormat__BorderStyle_None QTextFrameFormat__BorderStyle = 0
const QTextFrameFormat__BorderStyle_Dotted QTextFrameFormat__BorderStyle = 1
const QTextFrameFormat__BorderStyle_Dashed QTextFrameFormat__BorderStyle = 2
const QTextFrameFormat__BorderStyle_Solid QTextFrameFormat__BorderStyle = 3
const QTextFrameFormat__BorderStyle_Double QTextFrameFormat__BorderStyle = 4
const QTextFrameFormat__BorderStyle_DotDash QTextFrameFormat__BorderStyle = 5
const QTextFrameFormat__BorderStyle_DotDotDash QTextFrameFormat__BorderStyle = 6
const QTextFrameFormat__BorderStyle_Groove QTextFrameFormat__BorderStyle = 7
const QTextFrameFormat__BorderStyle_Ridge QTextFrameFormat__BorderStyle = 8
const QTextFrameFormat__BorderStyle_Inset QTextFrameFormat__BorderStyle = 9
const QTextFrameFormat__BorderStyle_Outset QTextFrameFormat__BorderStyle = 10

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
