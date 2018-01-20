//  header block begin
// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h
// #include <qabstracttextdocumentlayout.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 47
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
type QAbstractTextDocumentLayout struct {
	*qtcore.QObject
}

func (this *QAbstractTextDocumentLayout) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:60
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractTextDocumentLayout) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:64
// index:0
// void QAbstractTextDocumentLayout(class QTextDocument *)
func NewQAbstractTextDocumentLayout(doc unsafe.Pointer) *QAbstractTextDocumentLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayoutC1EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, doc)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractTextDocumentLayoutFromPointer(cthis)
	return gothis
}
func NewQAbstractTextDocumentLayoutFromPointer(cthis unsafe.Pointer) *QAbstractTextDocumentLayout {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractTextDocumentLayout{bcthis0}
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:113
// index:1
// void QAbstractTextDocumentLayout(class QAbstractTextDocumentLayoutPrivate &, class QTextDocument *)
func NewQAbstractTextDocumentLayout_1(arg0 unsafe.Pointer, arg1 unsafe.Pointer) *QAbstractTextDocumentLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayoutC1ER34QAbstractTextDocumentLayoutPrivateP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, arg0, arg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractTextDocumentLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:65
// index:0
// virtual
// void ~QAbstractTextDocumentLayout()
func DeleteQAbstractTextDocumentLayout(*QAbstractTextDocumentLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:84
// index:0
// pure virtual
// void draw(class QPainter *, const struct QAbstractTextDocumentLayout::PaintContext &)
func (this *QAbstractTextDocumentLayout) Draw(painter unsafe.Pointer, context unsafe.Pointer) {
	// 0: (, painter QPainter *, context const QAbstractTextDocumentLayout::PaintContext &), (painter, context)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout4drawEP8QPainterRKNS_12PaintContextE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, context)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:85
// index:0
// pure virtual
// int hitTest(const class QPointF &, Qt::HitTestAccuracy)
func (this *QAbstractTextDocumentLayout) HitTest(point unsafe.Pointer, accuracy int) {
	// 0: (, point const QPointF &, accuracy Qt::HitTestAccuracy), (point, &accuracy)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout7hitTestERK7QPointFN2Qt15HitTestAccuracyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point, &accuracy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:87
// index:0
// QString anchorAt(const class QPointF &)
func (this *QAbstractTextDocumentLayout) AnchorAt(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout8anchorAtERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:88
// index:0
// QString imageAt(const class QPointF &)
func (this *QAbstractTextDocumentLayout) ImageAt(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout7imageAtERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:89
// index:0
// QTextFormat formatAt(const class QPointF &)
func (this *QAbstractTextDocumentLayout) FormatAt(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout8formatAtERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:91
// index:0
// pure virtual
// int pageCount()
func (this *QAbstractTextDocumentLayout) PageCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout9pageCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:92
// index:0
// pure virtual
// QSizeF documentSize()
func (this *QAbstractTextDocumentLayout) DocumentSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout12documentSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:94
// index:0
// pure virtual
// QRectF frameBoundingRect(class QTextFrame *)
func (this *QAbstractTextDocumentLayout) FrameBoundingRect(frame unsafe.Pointer) {
	// 0: (, frame QTextFrame *), (frame)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout17frameBoundingRectEP10QTextFrame", ffiqt.FFI_TYPE_VOID, this.GetCthis(), frame)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:95
// index:0
// pure virtual
// QRectF blockBoundingRect(const class QTextBlock &)
func (this *QAbstractTextDocumentLayout) BlockBoundingRect(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout17blockBoundingRectERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:97
// index:0
// void setPaintDevice(class QPaintDevice *)
func (this *QAbstractTextDocumentLayout) SetPaintDevice(device unsafe.Pointer) {
	// 0: (, device QPaintDevice *), (device)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout14setPaintDeviceEP12QPaintDevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:98
// index:0
// QPaintDevice * paintDevice()
func (this *QAbstractTextDocumentLayout) PaintDevice() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout11paintDeviceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:100
// index:0
// QTextDocument * document()
func (this *QAbstractTextDocumentLayout) Document() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout8documentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:102
// index:0
// void registerHandler(int, class QObject *)
func (this *QAbstractTextDocumentLayout) RegisterHandler(objectType int, component unsafe.Pointer) {
	// 0: (, objectType int, component QObject *), (&objectType, component)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout15registerHandlerEiP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &objectType, component)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:103
// index:0
// void unregisterHandler(int, class QObject *)
func (this *QAbstractTextDocumentLayout) UnregisterHandler(objectType int, component unsafe.Pointer) {
	// 0: (, objectType int, component QObject *), (&objectType, component)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout17unregisterHandlerEiP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &objectType, component)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:104
// index:0
// QTextObjectInterface * handlerForObject(int)
func (this *QAbstractTextDocumentLayout) HandlerForObject(objectType int) {
	// 0: (, objectType int), (&objectType)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout16handlerForObjectEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &objectType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:107
// index:0
// void update(const class QRectF &)
func (this *QAbstractTextDocumentLayout) Update(arg0 unsafe.Pointer) {
	// 0: (, const QRectF & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout6updateERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:108
// index:0
// void updateBlock(const class QTextBlock &)
func (this *QAbstractTextDocumentLayout) UpdateBlock(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout11updateBlockERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:109
// index:0
// void documentSizeChanged(const class QSizeF &)
func (this *QAbstractTextDocumentLayout) DocumentSizeChanged(newSize unsafe.Pointer) {
	// 0: (, newSize const QSizeF &), (newSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout19documentSizeChangedERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), newSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:110
// index:0
// void pageCountChanged(int)
func (this *QAbstractTextDocumentLayout) PageCountChanged(newPages int) {
	// 0: (, newPages int), (&newPages)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout16pageCountChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &newPages)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:115
// index:0
// pure virtual
// void documentChanged(int, int, int)
func (this *QAbstractTextDocumentLayout) DocumentChanged(from int, charsRemoved int, charsAdded int) {
	// 0: (, from int, charsRemoved int, charsAdded int), (&from, &charsRemoved, &charsAdded)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout15documentChangedEiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &from, &charsRemoved, &charsAdded)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:117
// index:0
// virtual
// void resizeInlineObject(class QTextInlineObject, int, const class QTextFormat &)
func (this *QAbstractTextDocumentLayout) ResizeInlineObject(item unsafe.Pointer, posInDocument int, format unsafe.Pointer) {
	// 0: (, item QTextInlineObject, posInDocument int, format const QTextFormat &), (item, &posInDocument, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout18resizeInlineObjectE17QTextInlineObjectiRK11QTextFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &posInDocument, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:118
// index:0
// virtual
// void positionInlineObject(class QTextInlineObject, int, const class QTextFormat &)
func (this *QAbstractTextDocumentLayout) PositionInlineObject(item unsafe.Pointer, posInDocument int, format unsafe.Pointer) {
	// 0: (, item QTextInlineObject, posInDocument int, format const QTextFormat &), (item, &posInDocument, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout20positionInlineObjectE17QTextInlineObjectiRK11QTextFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &posInDocument, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:119
// index:0
// virtual
// void drawInlineObject(class QPainter *, const class QRectF &, class QTextInlineObject, int, const class QTextFormat &)
func (this *QAbstractTextDocumentLayout) DrawInlineObject(painter unsafe.Pointer, rect unsafe.Pointer, object unsafe.Pointer, posInDocument int, format unsafe.Pointer) {
	// 0: (, painter QPainter *, rect const QRectF &, object QTextInlineObject, posInDocument int, format const QTextFormat &), (painter, rect, object, &posInDocument, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout16drawInlineObjectEP8QPainterRK6QRectF17QTextInlineObjectiRK11QTextFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect, object, &posInDocument, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:121
// index:0
// int formatIndex(int)
func (this *QAbstractTextDocumentLayout) FormatIndex(pos int) {
	// 0: (, pos int), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout11formatIndexEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:122
// index:0
// QTextCharFormat format(int)
func (this *QAbstractTextDocumentLayout) Format(pos int) {
	// 0: (, pos int), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout6formatEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

//  body block end
