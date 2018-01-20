//  header block begin
// /usr/include/qt/QtWidgets/qplaintextedit.h
// #include <qplaintextedit.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 116
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
	return this.QAbstractTextDocumentLayout.GetCthis()
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:297
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPlainTextDocumentLayout) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:302
// index:0
// void QPlainTextDocumentLayout(class QTextDocument *)
func NewQPlainTextDocumentLayout(document unsafe.Pointer) *QPlainTextDocumentLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayoutC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, document)
	gopp.ErrPrint(err, rv)
	gothis := NewQPlainTextDocumentLayoutFromPointer(cthis)
	return gothis
}
func NewQPlainTextDocumentLayoutFromPointer(cthis unsafe.Pointer) *QPlainTextDocumentLayout {
	bcthis0 := qtgui.NewQAbstractTextDocumentLayoutFromPointer(cthis)
	return &QPlainTextDocumentLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:303
// index:0
// virtual
// void ~QPlainTextDocumentLayout()
func DeleteQPlainTextDocumentLayout(*QPlainTextDocumentLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:305
// index:0
// virtual
// void draw(class QPainter *, const struct QAbstractTextDocumentLayout::PaintContext &)
func (this *QPlainTextDocumentLayout) Draw(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, QPainter * arg0, const QAbstractTextDocumentLayout::PaintContext & arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout4drawEP8QPainterRKN27QAbstractTextDocumentLayout12PaintContextE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:306
// index:0
// virtual
// int hitTest(const class QPointF &, Qt::HitTestAccuracy)
func (this *QPlainTextDocumentLayout) HitTest(arg0 unsafe.Pointer, arg1 int) {
	// 0: (, const QPointF & arg0, Qt::HitTestAccuracy arg1), (arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout7hitTestERK7QPointFN2Qt15HitTestAccuracyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:308
// index:0
// virtual
// int pageCount()
func (this *QPlainTextDocumentLayout) PageCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout9pageCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:309
// index:0
// virtual
// QSizeF documentSize()
func (this *QPlainTextDocumentLayout) DocumentSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout12documentSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:311
// index:0
// virtual
// QRectF frameBoundingRect(class QTextFrame *)
func (this *QPlainTextDocumentLayout) FrameBoundingRect(arg0 unsafe.Pointer) {
	// 0: (, QTextFrame * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:312
// index:0
// virtual
// QRectF blockBoundingRect(const class QTextBlock &)
func (this *QPlainTextDocumentLayout) BlockBoundingRect(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:314
// index:0
// void ensureBlockLayout(const class QTextBlock &)
func (this *QPlainTextDocumentLayout) EnsureBlockLayout(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:316
// index:0
// void setCursorWidth(int)
func (this *QPlainTextDocumentLayout) SetCursorWidth(width int) {
	// 0: (, width int), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout14setCursorWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:317
// index:0
// int cursorWidth()
func (this *QPlainTextDocumentLayout) CursorWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QPlainTextDocumentLayout11cursorWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:319
// index:0
// void requestUpdate()
func (this *QPlainTextDocumentLayout) RequestUpdate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout13requestUpdateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:322
// index:0
// virtual
// void documentChanged(int, int, int)
func (this *QPlainTextDocumentLayout) DocumentChanged(from int, arg1 int, charsAdded int) {
	// 0: (, from int, int arg1, charsAdded int), (&from, &arg1, &charsAdded)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QPlainTextDocumentLayout15documentChangedEiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &from, &arg1, &charsAdded)
	gopp.ErrPrint(err, rv)
}

//  body block end
