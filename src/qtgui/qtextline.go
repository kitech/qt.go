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
func NewQTextLineFromPointer(cthis unsafe.Pointer) *QTextLine {
	return &QTextLine{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextlayout.h:213
// index:0
// Public inline
// void QTextLine()
func NewQTextLine() *QTextLine {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLineC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:214
// index:0
// Public inline
// bool isValid()
func (this *QTextLine) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:216
// index:0
// Public
// QRectF rect()
func (this *QTextLine) Rect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine4rectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:217
// index:0
// Public
// qreal x()
func (this *QTextLine) X() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:218
// index:0
// Public
// qreal y()
func (this *QTextLine) Y() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:219
// index:0
// Public
// qreal width()
func (this *QTextLine) Width() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:220
// index:0
// Public
// qreal ascent()
func (this *QTextLine) Ascent() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine6ascentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:221
// index:0
// Public
// qreal descent()
func (this *QTextLine) Descent() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine7descentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:222
// index:0
// Public
// qreal height()
func (this *QTextLine) Height() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:223
// index:0
// Public
// qreal leading()
func (this *QTextLine) Leading() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine7leadingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:225
// index:0
// Public
// void setLeadingIncluded(_Bool)
func (this *QTextLine) SetLeadingIncluded(included bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLine18setLeadingIncludedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &included)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:226
// index:0
// Public
// bool leadingIncluded()
func (this *QTextLine) LeadingIncluded() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine15leadingIncludedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:228
// index:0
// Public
// qreal naturalTextWidth()
func (this *QTextLine) NaturalTextWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine16naturalTextWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:229
// index:0
// Public
// qreal horizontalAdvance()
func (this *QTextLine) HorizontalAdvance() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine17horizontalAdvanceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:230
// index:0
// Public
// QRectF naturalTextRect()
func (this *QTextLine) NaturalTextRect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine15naturalTextRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:242
// index:0
// Public
// qreal cursorToX(int *, enum QTextLine::Edge)
func (this *QTextLine) CursorToX(cursorPos unsafe.Pointer /*666*/, edge int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine9cursorToXEPiNS_4EdgeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cursorPos, &edge)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:243
// index:1
// Public inline
// qreal cursorToX(int, enum QTextLine::Edge)
func (this *QTextLine) CursorToX_1(cursorPos int, edge int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine9cursorToXEiNS_4EdgeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cursorPos, &edge)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextlayout.h:244
// index:0
// Public
// int xToCursor(qreal, enum QTextLine::CursorPosition)
func (this *QTextLine) XToCursor(x float64, arg1 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine9xToCursorEdNS_14CursorPositionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextlayout.h:246
// index:0
// Public
// void setLineWidth(qreal)
func (this *QTextLine) SetLineWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLine12setLineWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:247
// index:0
// Public
// void setNumColumns(int)
func (this *QTextLine) SetNumColumns(columns int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLine13setNumColumnsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:248
// index:1
// Public
// void setNumColumns(int, qreal)
func (this *QTextLine) SetNumColumns_1(columns int, alignmentWidth float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLine13setNumColumnsEid", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &columns, &alignmentWidth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:250
// index:0
// Public
// void setPosition(const class QPointF &)
func (this *QTextLine) SetPosition(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextLine11setPositionERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:251
// index:0
// Public
// QPointF position()
func (this *QTextLine) Position() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:253
// index:0
// Public
// int textStart()
func (this *QTextLine) TextStart() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine9textStartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextlayout.h:254
// index:0
// Public
// int textLength()
func (this *QTextLine) TextLength() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine10textLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextlayout.h:256
// index:0
// Public inline
// int lineNumber()
func (this *QTextLine) LineNumber() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextLine10lineNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
