package qtgui

// /usr/include/qt/QtGui/qtextlayout.h
// #include <qtextlayout.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QTextLayout struct {
	*qtrt.CObject
}
type QTextLayout_ITF interface {
	QTextLayout_PTR() *QTextLayout
}

func (ptr *QTextLayout) QTextLayout_PTR() *QTextLayout { return ptr }

func (this *QTextLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextLayout) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextLayoutFromPointer(cthis unsafe.Pointer) *QTextLayout {
	return &QTextLayout{&qtrt.CObject{cthis}}
}
func (*QTextLayout) NewFromPointer(cthis unsafe.Pointer) *QTextLayout {
	return NewQTextLayoutFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextlayout.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout()
func NewQTextLayout() *QTextLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLayout)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:109
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout(const QString &)
func NewQTextLayout_1(text string) *QTextLayout {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLayout)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:110
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout(const QString &, const QFont &, QPaintDevice *)
func NewQTextLayout_2(text string, font QFont_ITF, paintdevice QPaintDevice_ITF /*777 QPaintDevice **/) *QTextLayout {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = font.QFont_PTR().GetCthis()
	var convArg2 = paintdevice.QPaintDevice_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QStringRK5QFontP12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLayout)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:111
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout(const QTextBlock &)
func NewQTextLayout_3(b QTextBlock_ITF) *QTextLayout {
	var convArg0 = b.QTextBlock_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK10QTextBlock", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLayout)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextLayout()
func DeleteQTextLayout(this *QTextLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextlayout.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QTextLayout) SetFont(f QFont_ITF) {
	var convArg0 = f.QFont_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:115
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font()
func (this *QTextLayout) Font() *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawFont(const QRawFont &)
func (this *QTextLayout) SetRawFont(rawFont QRawFont_ITF) {
	var convArg0 = rawFont.QRawFont_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout10setRawFontERK8QRawFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QTextLayout) SetText(string string) {
	var tmpArg0 = qtcore.NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QTextLayout) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextlayout.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextOption(const QTextOption &)
func (this *QTextLayout) SetTextOption(option QTextOption_ITF) {
	var convArg0 = option.QTextOption_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout13setTextOptionERK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:125
// index:0
// Public Visibility=Default Availability=Available
// [32] const QTextOption & textOption()
func (this *QTextLayout) TextOption() *QTextOption {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout10textOptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextOption)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreeditArea(int, const QString &)
func (this *QTextLayout) SetPreeditArea(position int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout14setPreeditAreaEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] int preeditAreaPosition()
func (this *QTextLayout) PreeditAreaPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout19preeditAreaPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QString preeditAreaText()
func (this *QTextLayout) PreeditAreaText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout15preeditAreaTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextlayout.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearAdditionalFormats()
func (this *QTextLayout) ClearAdditionalFormats() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout22clearAdditionalFormatsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearFormats()
func (this *QTextLayout) ClearFormats() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout12clearFormatsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheEnabled(_Bool)
func (this *QTextLayout) SetCacheEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout15setCacheEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cacheEnabled()
func (this *QTextLayout) CacheEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout12cacheEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorMoveStyle(Qt::CursorMoveStyle)
func (this *QTextLayout) SetCursorMoveStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout18setCursorMoveStyleEN2Qt15CursorMoveStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:154
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CursorMoveStyle cursorMoveStyle()
func (this *QTextLayout) CursorMoveStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout15cursorMoveStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginLayout()
func (this *QTextLayout) BeginLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout11beginLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endLayout()
func (this *QTextLayout) EndLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout9endLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearLayout()
func (this *QTextLayout) ClearLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout11clearLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:160
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLine createLine()
func (this *QTextLayout) CreateLine() *QTextLine /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout10createLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLine)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:162
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineCount()
func (this *QTextLayout) LineCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout9lineCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:163
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLine lineAt(int)
func (this *QTextLayout) LineAt(i int) *QTextLine /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout6lineAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLine)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:164
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLine lineForTextPosition(int)
func (this *QTextLayout) LineForTextPosition(pos int) *QTextLine /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout19lineForTextPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLine)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValidCursorPosition(int)
func (this *QTextLayout) IsValidCursorPosition(pos int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout21isValidCursorPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:171
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextCursorPosition(int, enum QTextLayout::CursorMode)
func (this *QTextLayout) NextCursorPosition(oldPos int, mode int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout18nextCursorPositionEiNS_10CursorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldPos, mode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:172
// index:0
// Public Visibility=Default Availability=Available
// [4] int previousCursorPosition(int, enum QTextLayout::CursorMode)
func (this *QTextLayout) PreviousCursorPosition(oldPos int, mode int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout22previousCursorPositionEiNS_10CursorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldPos, mode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:173
// index:0
// Public Visibility=Default Availability=Available
// [4] int leftCursorPosition(int)
func (this *QTextLayout) LeftCursorPosition(oldPos int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout18leftCursorPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldPos)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:174
// index:0
// Public Visibility=Default Availability=Available
// [4] int rightCursorPosition(int)
func (this *QTextLayout) RightCursorPosition(oldPos int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout19rightCursorPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldPos)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawCursor(QPainter *, const QPointF &, int)
func (this *QTextLayout) DrawCursor(p QPainter_ITF /*777 QPainter **/, pos qtcore.QPointF_ITF, cursorPosition int) {
	var convArg0 = p.QPainter_PTR().GetCthis()
	var convArg1 = pos.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cursorPosition)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:179
// index:1
// Public Visibility=Default Availability=Available
// [-2] void drawCursor(QPainter *, const QPointF &, int, int)
func (this *QTextLayout) DrawCursor_1(p QPainter_ITF /*777 QPainter **/, pos qtcore.QPointF_ITF, cursorPosition int, width int) {
	var convArg0 = p.QPainter_PTR().GetCthis()
	var convArg1 = pos.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cursorPosition, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:181
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF position()
func (this *QTextLayout) Position() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPointF &)
func (this *QTextLayout) SetPosition(p qtcore.QPointF_ITF) {
	var convArg0 = p.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout11setPositionERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:184
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QTextLayout) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:186
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal minimumWidth()
func (this *QTextLayout) MinimumWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout12minimumWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:187
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maximumWidth()
func (this *QTextLayout) MaximumWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout12maximumWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(int)
func (this *QTextLayout) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout8setFlagsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

type QTextLayout__CursorMode = int

const QTextLayout__SkipCharacters QTextLayout__CursorMode = 0
const QTextLayout__SkipWords QTextLayout__CursorMode = 1

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
