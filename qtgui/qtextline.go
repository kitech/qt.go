package qtgui

// /usr/include/qt/QtGui/qtextlayout.h
// #include <qtextlayout.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 41
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QTextLine struct {
	*qtrt.CObject
}

func (this *QTextLine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextLine) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextLineFromPointer(cthis unsafe.Pointer) *QTextLine {
	return &QTextLine{&qtrt.CObject{cthis}}
}
func (*QTextLine) NewFromPointer(cthis unsafe.Pointer) *QTextLine {
	return NewQTextLineFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextlayout.h:213
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextLine()
func NewQTextLine() *QTextLine {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLineC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLine)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextLine) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:216
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect()
func (this *QTextLine) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:217
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal x()
func (this *QTextLine) X() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:218
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal y()
func (this *QTextLine) Y() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:219
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width()
func (this *QTextLine) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:220
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal ascent()
func (this *QTextLine) Ascent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine6ascentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:221
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal descent()
func (this *QTextLine) Descent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine7descentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal height()
func (this *QTextLine) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:223
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal leading()
func (this *QTextLine) Leading() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine7leadingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLeadingIncluded(_Bool)
func (this *QTextLine) SetLeadingIncluded(included bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLine18setLeadingIncludedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), included)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:226
// index:0
// Public Visibility=Default Availability=Available
// [1] bool leadingIncluded()
func (this *QTextLine) LeadingIncluded() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine15leadingIncludedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal naturalTextWidth()
func (this *QTextLine) NaturalTextWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine16naturalTextWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:229
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal horizontalAdvance()
func (this *QTextLine) HorizontalAdvance() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine17horizontalAdvanceEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:230
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF naturalTextRect()
func (this *QTextLine) NaturalTextRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine15naturalTextRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:242
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal cursorToX(int *, enum QTextLine::Edge)
func (this *QTextLine) CursorToX(cursorPos unsafe.Pointer /*666*/, edge int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9cursorToXEPiNS_4EdgeE", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), &cursorPos, edge)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:243
// index:1
// Public inline Visibility=Default Availability=Available
// [8] qreal cursorToX(int, enum QTextLine::Edge)
func (this *QTextLine) CursorToX_1(cursorPos int, edge int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9cursorToXEiNS_4EdgeE", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), cursorPos, edge)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:244
// index:0
// Public Visibility=Default Availability=Available
// [4] int xToCursor(qreal, enum QTextLine::CursorPosition)
func (this *QTextLine) XToCursor(x float64, arg1 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9xToCursorEdNS_14CursorPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:246
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineWidth(qreal)
func (this *QTextLine) SetLineWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLine12setLineWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNumColumns(int)
func (this *QTextLine) SetNumColumns(columns int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLine13setNumColumnsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:248
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setNumColumns(int, qreal)
func (this *QTextLine) SetNumColumns_1(columns int, alignmentWidth float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLine13setNumColumnsEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns, alignmentWidth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPointF &)
func (this *QTextLine) SetPosition(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLine11setPositionERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:251
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF position()
func (this *QTextLine) Position() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:253
// index:0
// Public Visibility=Default Availability=Available
// [4] int textStart()
func (this *QTextLine) TextStart() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9textStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:254
// index:0
// Public Visibility=Default Availability=Available
// [4] int textLength()
func (this *QTextLine) TextLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine10textLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:256
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lineNumber()
func (this *QTextLine) LineNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine10lineNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQTextLine(this *QTextLine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QTextLine__Edge = int

const QTextLine__Leading QTextLine__Edge = 0
const QTextLine__Trailing QTextLine__Edge = 1

type QTextLine__CursorPosition = int

const QTextLine__CursorBetweenCharacters QTextLine__CursorPosition = 0
const QTextLine__CursorOnCharacter QTextLine__CursorPosition = 1

//  body block end
