package qtwidgets

// /usr/include/qt/QtWidgets/qplaintextedit.h
// #include <qplaintextedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 111
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QPlainTextDocumentLayout struct {
	*qtgui.QAbstractTextDocumentLayout
}

func (this *QPlainTextDocumentLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractTextDocumentLayout.GetCthis()
	}
}
func NewQPlainTextDocumentLayoutFromPointer(cthis unsafe.Pointer) *QPlainTextDocumentLayout {
	bcthis0 := qtgui.NewQAbstractTextDocumentLayoutFromPointer(cthis)
	return &QPlainTextDocumentLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:297
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPlainTextDocumentLayout) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:302
// index:0
// Public
// void QPlainTextDocumentLayout(class QTextDocument *)
func NewQPlainTextDocumentLayout(document *qtgui.QTextDocument /*444 QTextDocument **/) *QPlainTextDocumentLayout {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = document.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayoutC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPlainTextDocumentLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:303
// index:0
// Public virtual
// void ~QPlainTextDocumentLayout()
func DeleteQPlainTextDocumentLayout(*QPlainTextDocumentLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:306
// index:0
// Public virtual
// int hitTest(const class QPointF &, Qt::HitTestAccuracy)
func (this *QPlainTextDocumentLayout) HitTest(arg0 *qtcore.QPointF, arg1 int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout7hitTestERK7QPointFN2Qt15HitTestAccuracyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:308
// index:0
// Public virtual
// int pageCount()
func (this *QPlainTextDocumentLayout) PageCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout9pageCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:309
// index:0
// Public virtual
// QSizeF documentSize()
func (this *QPlainTextDocumentLayout) DocumentSize() *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout12documentSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:311
// index:0
// Public virtual
// QRectF frameBoundingRect(class QTextFrame *)
func (this *QPlainTextDocumentLayout) FrameBoundingRect(arg0 *qtgui.QTextFrame /*444 QTextFrame **/) *qtcore.QRectF /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:312
// index:0
// Public virtual
// QRectF blockBoundingRect(const class QTextBlock &)
func (this *QPlainTextDocumentLayout) BlockBoundingRect(block *qtgui.QTextBlock) *qtcore.QRectF /*123*/ {
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:314
// index:0
// Public
// void ensureBlockLayout(const class QTextBlock &)
func (this *QPlainTextDocumentLayout) EnsureBlockLayout(block *qtgui.QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:316
// index:0
// Public
// void setCursorWidth(int)
func (this *QPlainTextDocumentLayout) SetCursorWidth(width int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout14setCursorWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:317
// index:0
// Public
// int cursorWidth()
func (this *QPlainTextDocumentLayout) CursorWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout11cursorWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:319
// index:0
// Public
// void requestUpdate()
func (this *QPlainTextDocumentLayout) RequestUpdate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout13requestUpdateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:322
// index:0
// Protected virtual
// void documentChanged(int, int, int)
func (this *QPlainTextDocumentLayout) DocumentChanged(from int, arg1 int, charsAdded int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout15documentChangedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &from, &arg1, &charsAdded)
	gopp.ErrPrint(err, rv)
}

//  body block end
