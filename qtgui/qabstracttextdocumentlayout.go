package qtgui

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h
// #include <qabstracttextdocumentlayout.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 51
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// void documentChanged(int, int, int)
func (this *QAbstractTextDocumentLayout) InheritDocumentChanged(f func(from int, charsRemoved int, charsAdded int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "documentChanged", f)
}

// void resizeInlineObject(QTextInlineObject, int, const QTextFormat &)
func (this *QAbstractTextDocumentLayout) InheritResizeInlineObject(f func(item *QTextInlineObject /*123*/, posInDocument int, format *QTextFormat) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeInlineObject", f)
}

// void positionInlineObject(QTextInlineObject, int, const QTextFormat &)
func (this *QAbstractTextDocumentLayout) InheritPositionInlineObject(f func(item *QTextInlineObject /*123*/, posInDocument int, format *QTextFormat) /*void*/) {
	qtrt.SetAllInheritCallback(this, "positionInlineObject", f)
}

// void drawInlineObject(QPainter *, const QRectF &, QTextInlineObject, int, const QTextFormat &)
func (this *QAbstractTextDocumentLayout) InheritDrawInlineObject(f func(painter *QPainter /*777 QPainter **/, rect *qtcore.QRectF, object *QTextInlineObject /*123*/, posInDocument int, format *QTextFormat) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawInlineObject", f)
}

// int formatIndex(int)
func (this *QAbstractTextDocumentLayout) InheritFormatIndex(f func(pos int) int) {
	qtrt.SetAllInheritCallback(this, "formatIndex", f)
}

// QTextCharFormat format(int)
func (this *QAbstractTextDocumentLayout) InheritFormat(f func(pos int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "format", f)
}

/*

 */
type QAbstractTextDocumentLayout struct {
	*qtcore.QObject
}
type QAbstractTextDocumentLayout_ITF interface {
	qtcore.QObject_ITF
	QAbstractTextDocumentLayout_PTR() *QAbstractTextDocumentLayout
}

func (ptr *QAbstractTextDocumentLayout) QAbstractTextDocumentLayout_PTR() *QAbstractTextDocumentLayout {
	return ptr
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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractTextDocumentLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractTextDocumentLayout(QTextDocument *)

/*
Creates a new text document layout for the given document.
*/
func (*QAbstractTextDocumentLayout) NewForInherit(doc QTextDocument_ITF /*777 QTextDocument **/) *QAbstractTextDocumentLayout {
	return NewQAbstractTextDocumentLayout(doc)
}
func NewQAbstractTextDocumentLayout(doc QTextDocument_ITF /*777 QTextDocument **/) *QAbstractTextDocumentLayout {
	var convArg0 unsafe.Pointer
	if doc != nil && doc.QTextDocument_PTR() != nil {
		convArg0 = doc.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayoutC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractTextDocumentLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractTextDocumentLayout")
	return gothis
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractTextDocumentLayout()

/*

 */
func DeleteQAbstractTextDocumentLayout(this *QAbstractTextDocumentLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:85
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int hitTest(const QPointF &, Qt::HitTestAccuracy) const

/*
Returns the cursor position for the given point with the specified accuracy. Returns -1 if no valid cursor position was found.
*/
func (this *QAbstractTextDocumentLayout) HitTest(point qtcore.QPointF_ITF, accuracy int) int {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout7hitTestERK7QPointFN2Qt15HitTestAccuracyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, accuracy)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString anchorAt(const QPointF &) const

/*
Returns the reference of the anchor the given position, or an empty string if no anchor exists at that point.
*/
func (this *QAbstractTextDocumentLayout) AnchorAt(pos qtcore.QPointF_ITF) string {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout8anchorAtERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QString imageAt(const QPointF &) const

/*
Returns the source of the image at the given position pos, or an empty string if no image exists at that point.

This function was introduced in  Qt 5.8.
*/
func (this *QAbstractTextDocumentLayout) ImageAt(pos qtcore.QPointF_ITF) string {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout7imageAtERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:89
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextFormat formatAt(const QPointF &) const

/*
Returns the text format at the given position pos.

This function was introduced in  Qt 5.8.
*/
func (this *QAbstractTextDocumentLayout) FormatAt(pos qtcore.QPointF_ITF) *QTextFormat /*123*/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout8formatAtERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextFormat)
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:91
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int pageCount() const

/*
Returns the number of pages contained in the layout.

See also pageCountChanged().
*/
func (this *QAbstractTextDocumentLayout) PageCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout9pageCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:92
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QSizeF documentSize() const

/*
Returns the total size of the document's layout.

This information can be used by display widgets to update their scroll bars correctly.

See also documentSizeChanged() and QTextDocument::pageSize.
*/
func (this *QAbstractTextDocumentLayout) DocumentSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout12documentSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:94
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [32] QRectF frameBoundingRect(QTextFrame *) const

/*
Returns the bounding rectangle of frame.
*/
func (this *QAbstractTextDocumentLayout) FrameBoundingRect(frame QTextFrame_ITF /*777 QTextFrame **/) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if frame != nil && frame.QTextFrame_PTR() != nil {
		convArg0 = frame.QTextFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout17frameBoundingRectEP10QTextFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:95
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [32] QRectF blockBoundingRect(const QTextBlock &) const

/*
Returns the bounding rectangle of block.
*/
func (this *QAbstractTextDocumentLayout) BlockBoundingRect(block QTextBlock_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout17blockBoundingRectERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaintDevice(QPaintDevice *)

/*
Sets the paint device used for rendering the document's layout to the given device.

See also paintDevice().
*/
func (this *QAbstractTextDocumentLayout) SetPaintDevice(device QPaintDevice_ITF /*777 QPaintDevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QPaintDevice_PTR() != nil {
		convArg0 = device.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout14setPaintDeviceEP12QPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QPaintDevice * paintDevice() const

/*
Returns the paint device used to render the document's layout.

See also setPaintDevice().
*/
func (this *QAbstractTextDocumentLayout) PaintDevice() *QPaintDevice /*777 QPaintDevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout11paintDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document() const

/*
Returns the text document that this layout is operating on.
*/
func (this *QAbstractTextDocumentLayout) Document() *QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void registerHandler(int, QObject *)

/*
Registers the given component as a handler for items of the given objectType.

Note: registerHandler() has to be called once for each object type. This means that there is only one handler for multiple replacement characters of the same object type.

The text document layout does not take ownership of component.
*/
func (this *QAbstractTextDocumentLayout) RegisterHandler(objectType int, component qtcore.QObject_ITF /*777 QObject **/) {
	var convArg1 unsafe.Pointer
	if component != nil && component.QObject_PTR() != nil {
		convArg1 = component.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout15registerHandlerEiP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), objectType, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unregisterHandler(int, QObject *)

/*
Unregisters the given component as a handler for items of the given objectType, or any handler if the component is not specified.

This function was introduced in  Qt 5.2.
*/
func (this *QAbstractTextDocumentLayout) UnregisterHandler(objectType int, component qtcore.QObject_ITF /*777 QObject **/) {
	var convArg1 unsafe.Pointer
	if component != nil && component.QObject_PTR() != nil {
		convArg1 = component.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout17unregisterHandlerEiP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), objectType, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unregisterHandler(int, QObject *)

/*
Unregisters the given component as a handler for items of the given objectType, or any handler if the component is not specified.

This function was introduced in  Qt 5.2.
*/
func (this *QAbstractTextDocumentLayout) UnregisterHandlerp(objectType int) {
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout17unregisterHandlerEiP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), objectType, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextObjectInterface * handlerForObject(int) const

/*
Returns a handler for objects of the given objectType.
*/
func (this *QAbstractTextDocumentLayout) HandlerForObject(objectType int) *QTextObjectInterface /*777 QTextObjectInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QAbstractTextDocumentLayout16handlerForObjectEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), objectType)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextObjectInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update(const QRectF &)

/*
This signal is emitted when the rectangle rect has been updated.

Subclasses of QAbstractTextDocumentLayout should emit this signal when the layout of the contents change in order to repaint.
*/
func (this *QAbstractTextDocumentLayout) Update(arg0 qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRectF_PTR() != nil {
		convArg0 = arg0.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout6updateERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update(const QRectF &)

/*
This signal is emitted when the rectangle rect has been updated.

Subclasses of QAbstractTextDocumentLayout should emit this signal when the layout of the contents change in order to repaint.
*/
func (this *QAbstractTextDocumentLayout) Updatep() {
	// arg: 0, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout6updateERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateBlock(const QTextBlock &)

/*
This signal is emitted when the specified block has been updated.

Subclasses of QAbstractTextDocumentLayout should emit this signal when the layout of block has changed in order to repaint.

This function was introduced in  Qt 4.4.
*/
func (this *QAbstractTextDocumentLayout) UpdateBlock(block QTextBlock_ITF) {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout11updateBlockERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void documentSizeChanged(const QSizeF &)

/*
This signal is emitted when the size of the document layout changes to newSize.

Subclasses of QAbstractTextDocumentLayout should emit this signal when the document's entire layout size changes. This signal is useful for widgets that display text documents since it enables them to update their scroll bars correctly.

See also documentSize().
*/
func (this *QAbstractTextDocumentLayout) DocumentSizeChanged(newSize qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if newSize != nil && newSize.QSizeF_PTR() != nil {
		convArg0 = newSize.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout19documentSizeChangedERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pageCountChanged(int)

/*
This signal is emitted when the number of pages in the layout changes; newPages is the updated page count.

Subclasses of QAbstractTextDocumentLayout should emit this signal when the number of pages in the layout has changed. Changes to the page count are caused by changes to the layout or the document content itself.

See also pageCount().
*/
func (this *QAbstractTextDocumentLayout) PageCountChanged(newPages int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout16pageCountChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newPages)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:115
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void documentChanged(int, int, int)

/*
This function is called whenever the contents of the document change. A change occurs when text is inserted, removed, or a combination of these two. The change is specified by position, charsRemoved, and charsAdded corresponding to the starting character position of the change, the number of characters removed from the document, and the number of characters added.

For example, when inserting the text "Hello" into an empty document, charsRemoved would be 0 and charsAdded would be 5 (the length of the string).

Replacing text is a combination of removing and inserting. For example, if the text "Hello" gets replaced by "Hi", charsRemoved would be 5 and charsAdded would be 2.

For subclasses of QAbstractTextDocumentLayout, this is the central function where a large portion of the work to lay out and position document contents is done.

For example, in a subclass that only arranges blocks of text, an implementation of this function would have to do the following:


Determine the list of changed QTextBlock(s) using the parameters provided.
Each QTextBlock object's corresponding QTextLayout object needs to be processed. You can access the QTextBlock's layout using the QTextBlock::layout() function. This processing should take the document's page size into consideration.
If the total number of pages changed, the pageCountChanged() signal should be emitted.
If the total size changed, the documentSizeChanged() signal should be emitted.
The update() signal should be emitted to schedule a repaint of areas in the layout that require repainting.


See also QTextLayout.
*/
func (this *QAbstractTextDocumentLayout) DocumentChanged(from int, charsRemoved int, charsAdded int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout15documentChangedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, charsRemoved, charsAdded)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeInlineObject(QTextInlineObject, int, const QTextFormat &)

/*
Sets the size of the inline object item corresponding to the text format.

posInDocument specifies the position of the object within the document.

The default implementation resizes the item to the size returned by the object handler's intrinsicSize() function. This function is called only within Qt. Subclasses can reimplement this function to customize the resizing of inline objects.
*/
func (this *QAbstractTextDocumentLayout) ResizeInlineObject(item QTextInlineObject_ITF /*123*/, posInDocument int, format QTextFormat_ITF) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTextInlineObject_PTR() != nil {
		convArg0 = item.QTextInlineObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if format != nil && format.QTextFormat_PTR() != nil {
		convArg2 = format.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout18resizeInlineObjectE17QTextInlineObjectiRK11QTextFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, posInDocument, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:118
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void positionInlineObject(QTextInlineObject, int, const QTextFormat &)

/*
Lays out the inline object item using the given text format.

posInDocument specifies the position of the object within the document.

The default implementation does nothing. This function is called only within Qt. Subclasses can reimplement this function to customize the position of inline objects.

See also drawInlineObject().
*/
func (this *QAbstractTextDocumentLayout) PositionInlineObject(item QTextInlineObject_ITF /*123*/, posInDocument int, format QTextFormat_ITF) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTextInlineObject_PTR() != nil {
		convArg0 = item.QTextInlineObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if format != nil && format.QTextFormat_PTR() != nil {
		convArg2 = format.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout20positionInlineObjectE17QTextInlineObjectiRK11QTextFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, posInDocument, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawInlineObject(QPainter *, const QRectF &, QTextInlineObject, int, const QTextFormat &)

/*
This function is called to draw the inline object, object, with the given painter within the rectangle specified by rect using the specified text format.

posInDocument specifies the position of the object within the document.

The default implementation calls drawObject() on the object handlers. This function is called only within Qt. Subclasses can reimplement this function to customize the drawing of inline objects.

See also draw().
*/
func (this *QAbstractTextDocumentLayout) DrawInlineObject(painter QPainter_ITF /*777 QPainter **/, rect qtcore.QRectF_ITF, object QTextInlineObject_ITF /*123*/, posInDocument int, format QTextFormat_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if object != nil && object.QTextInlineObject_PTR() != nil {
		convArg2 = object.QTextInlineObject_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if format != nil && format.QTextFormat_PTR() != nil {
		convArg4 = format.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout16drawInlineObjectEP8QPainterRK6QRectF17QTextInlineObjectiRK11QTextFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, posInDocument, convArg4)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:121
// index:0
// Protected Visibility=Default Availability=Available
// [4] int formatIndex(int)

/*

 */
func (this *QAbstractTextDocumentLayout) FormatIndex(pos int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout11formatIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:122
// index:0
// Protected Visibility=Default Availability=Available
// [16] QTextCharFormat format(int)

/*
Returns the character format that is applicable at the given position.
*/
func (this *QAbstractTextDocumentLayout) Format(pos int) *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAbstractTextDocumentLayout6formatEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
