package qtgui

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h
// #include <qabstracttextdocumentlayout.h>
// #include <QtGui>

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
import "mkuse/cffiqt"
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
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractTextDocumentLayout) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAbstractTextDocumentLayoutFromPointer(cthis unsafe.Pointer) *QAbstractTextDocumentLayout {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractTextDocumentLayout{bcthis0}
}
func (*QAbstractTextDocumentLayout) NewFromPointer(cthis unsafe.Pointer) *QAbstractTextDocumentLayout {
	return NewQAbstractTextDocumentLayoutFromPointer(cthis)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:60
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractTextDocumentLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:64
// index:0
// Public
// void QAbstractTextDocumentLayout(class QTextDocument *)
func NewQAbstractTextDocumentLayout(doc *QTextDocument /*777 QTextDocument **/) *QAbstractTextDocumentLayout {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = doc.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayoutC1EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractTextDocumentLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:65
// index:0
// Public virtual
// void ~QAbstractTextDocumentLayout()
func DeleteQAbstractTextDocumentLayout(*QAbstractTextDocumentLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:85
// index:0
// Public pure virtual
// int hitTest(const class QPointF &, Qt::HitTestAccuracy)
func (this *QAbstractTextDocumentLayout) HitTest(point *qtcore.QPointF, accuracy int) int {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout7hitTestERK7QPointFN2Qt15HitTestAccuracyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, accuracy)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:87
// index:0
// Public
// QString anchorAt(const class QPointF &)
func (this *QAbstractTextDocumentLayout) AnchorAt(pos *qtcore.QPointF) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout8anchorAtERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:88
// index:0
// Public
// QString imageAt(const class QPointF &)
func (this *QAbstractTextDocumentLayout) ImageAt(pos *qtcore.QPointF) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout7imageAtERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:89
// index:0
// Public
// QTextFormat formatAt(const class QPointF &)
func (this *QAbstractTextDocumentLayout) FormatAt(pos *qtcore.QPointF) *QTextFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout8formatAtERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:91
// index:0
// Public pure virtual
// int pageCount()
func (this *QAbstractTextDocumentLayout) PageCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout9pageCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:92
// index:0
// Public pure virtual
// QSizeF documentSize()
func (this *QAbstractTextDocumentLayout) DocumentSize() *qtcore.QSizeF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout12documentSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:94
// index:0
// Public pure virtual
// QRectF frameBoundingRect(class QTextFrame *)
func (this *QAbstractTextDocumentLayout) FrameBoundingRect(frame *QTextFrame /*777 QTextFrame **/) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = frame.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout17frameBoundingRectEP10QTextFrame", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:95
// index:0
// Public pure virtual
// QRectF blockBoundingRect(const class QTextBlock &)
func (this *QAbstractTextDocumentLayout) BlockBoundingRect(block *QTextBlock) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout17blockBoundingRectERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:97
// index:0
// Public
// void setPaintDevice(class QPaintDevice *)
func (this *QAbstractTextDocumentLayout) SetPaintDevice(device *QPaintDevice /*777 QPaintDevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout14setPaintDeviceEP12QPaintDevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:98
// index:0
// Public
// QPaintDevice * paintDevice()
func (this *QAbstractTextDocumentLayout) PaintDevice() *QPaintDevice /*777 QPaintDevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout11paintDeviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:100
// index:0
// Public
// QTextDocument * document()
func (this *QAbstractTextDocumentLayout) Document() *QTextDocument /*777 QTextDocument **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout8documentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:102
// index:0
// Public
// void registerHandler(int, class QObject *)
func (this *QAbstractTextDocumentLayout) RegisterHandler(objectType int, component *qtcore.QObject /*777 QObject **/) {
	var convArg1 = component.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout15registerHandlerEiP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), objectType, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:103
// index:0
// Public
// void unregisterHandler(int, class QObject *)
func (this *QAbstractTextDocumentLayout) UnregisterHandler(objectType int, component *qtcore.QObject /*777 QObject **/) {
	var convArg1 = component.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout17unregisterHandlerEiP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), objectType, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:104
// index:0
// Public
// QTextObjectInterface * handlerForObject(int)
func (this *QAbstractTextDocumentLayout) HandlerForObject(objectType int) *QTextObjectInterface /*777 QTextObjectInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout16handlerForObjectEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), objectType)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextObjectInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:107
// index:0
// Public
// void update(const class QRectF &)
func (this *QAbstractTextDocumentLayout) Update(arg0 *qtcore.QRectF) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout6updateERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:108
// index:0
// Public
// void updateBlock(const class QTextBlock &)
func (this *QAbstractTextDocumentLayout) UpdateBlock(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout11updateBlockERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:109
// index:0
// Public
// void documentSizeChanged(const class QSizeF &)
func (this *QAbstractTextDocumentLayout) DocumentSizeChanged(newSize *qtcore.QSizeF) {
	var convArg0 = newSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout19documentSizeChangedERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:110
// index:0
// Public
// void pageCountChanged(int)
func (this *QAbstractTextDocumentLayout) PageCountChanged(newPages int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout16pageCountChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newPages)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:115
// index:0
// Protected pure virtual
// void documentChanged(int, int, int)
func (this *QAbstractTextDocumentLayout) DocumentChanged(from int, charsRemoved int, charsAdded int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout15documentChangedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), from, charsRemoved, charsAdded)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:117
// index:0
// Protected virtual
// void resizeInlineObject(class QTextInlineObject, int, const class QTextFormat &)
func (this *QAbstractTextDocumentLayout) ResizeInlineObject(item *QTextInlineObject /*123*/, posInDocument int, format *QTextFormat) {
	var convArg0 = item.GetCthis()
	var convArg2 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout18resizeInlineObjectE17QTextInlineObjectiRK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, posInDocument, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:118
// index:0
// Protected virtual
// void positionInlineObject(class QTextInlineObject, int, const class QTextFormat &)
func (this *QAbstractTextDocumentLayout) PositionInlineObject(item *QTextInlineObject /*123*/, posInDocument int, format *QTextFormat) {
	var convArg0 = item.GetCthis()
	var convArg2 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout20positionInlineObjectE17QTextInlineObjectiRK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, posInDocument, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:119
// index:0
// Protected virtual
// void drawInlineObject(class QPainter *, const class QRectF &, class QTextInlineObject, int, const class QTextFormat &)
func (this *QAbstractTextDocumentLayout) DrawInlineObject(painter *QPainter /*777 QPainter **/, rect *qtcore.QRectF, object *QTextInlineObject /*123*/, posInDocument int, format *QTextFormat) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg2 = object.GetCthis()
	var convArg4 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout16drawInlineObjectEP8QPainterRK6QRectF17QTextInlineObjectiRK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, posInDocument, convArg4)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:121
// index:0
// Protected
// int formatIndex(int)
func (this *QAbstractTextDocumentLayout) FormatIndex(pos int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout11formatIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:122
// index:0
// Protected
// QTextCharFormat format(int)
func (this *QAbstractTextDocumentLayout) Format(pos int) *QTextCharFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout6formatEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
