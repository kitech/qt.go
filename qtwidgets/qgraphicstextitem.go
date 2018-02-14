package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool sceneEvent(class QEvent *)
func (this *QGraphicsTextItem) InheritSceneEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "sceneEvent", f)
}

// void mousePressEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) InheritMousePressEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) InheritMouseMoveEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) InheritMouseReleaseEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) InheritMouseDoubleClickEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsTextItem) InheritContextMenuEvent(f func(event *QGraphicsSceneContextMenuEvent /*777 QGraphicsSceneContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsTextItem) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsTextItem) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsTextItem) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsTextItem) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) InheritDragEnterEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) InheritDragLeaveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) InheritDragMoveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dropEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) InheritDropEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsTextItem) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void hoverEnterEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) InheritHoverEnterEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverEnterEvent", f)
}

// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) InheritHoverMoveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverMoveEvent", f)
}

// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) InheritHoverLeaveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverLeaveEvent", f)
}

// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsTextItem) InheritInputMethodQuery(f func(query int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "inputMethodQuery", f)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsTextItem) InheritSupportsExtension(f func(extension int) bool) {
	qtrt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsTextItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const class QVariant &)
func (this *QGraphicsTextItem) InheritExtension(f func(variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "extension", f)
}

type QGraphicsTextItem struct {
	*qtrt.CObject
}
type QGraphicsTextItem_ITF interface {
	QGraphicsTextItem_PTR() *QGraphicsTextItem
}

func (ptr *QGraphicsTextItem) QGraphicsTextItem_PTR() *QGraphicsTextItem { return ptr }

func (this *QGraphicsTextItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGraphicsTextItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGraphicsTextItemFromPointer(cthis unsafe.Pointer) *QGraphicsTextItem {
	return &QGraphicsTextItem{&qtrt.CObject{cthis}}
}
func (*QGraphicsTextItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsTextItem {
	return NewQGraphicsTextItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:872
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGraphicsTextItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:877
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsTextItem(QGraphicsItem *)
func NewQGraphicsTextItem(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsTextItem {
	var convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsTextItemFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:878
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsTextItem(const QString &, QGraphicsItem *)
func NewQGraphicsTextItem_1(text string, parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsTextItem {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.QGraphicsItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItemC2ERK7QStringP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsTextItemFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:879
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsTextItem()
func DeleteQGraphicsTextItem(this *QGraphicsTextItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:881
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toHtml()
func (this *QGraphicsTextItem) ToHtml() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem6toHtmlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:882
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &)
func (this *QGraphicsTextItem) SetHtml(html string) {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem7setHtmlERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:884
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toPlainText()
func (this *QGraphicsTextItem) ToPlainText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem11toPlainTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:885
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlainText(const QString &)
func (this *QGraphicsTextItem) SetPlainText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setPlainTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:887
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font()
func (this *QGraphicsTextItem) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:888
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QGraphicsTextItem) SetFont(font qtgui.QFont_ITF) {
	var convArg0 = font.QFont_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:890
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultTextColor(const QColor &)
func (this *QGraphicsTextItem) SetDefaultTextColor(c qtgui.QColor_ITF) {
	var convArg0 = c.QColor_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:891
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor defaultTextColor()
func (this *QGraphicsTextItem) DefaultTextColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem16defaultTextColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:893
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QGraphicsTextItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:894
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath shape()
func (this *QGraphicsTextItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:895
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &)
func (this *QGraphicsTextItem) Contains(point qtcore.QPointF_ITF) bool {
	var convArg0 = point.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:897
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsTextItem) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 = painter.QPainter_PTR().GetCthis()
	var convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	var convArg2 = widget.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:899
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *)
func (this *QGraphicsTextItem) IsObscuredBy(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 = item.QGraphicsItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:900
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea()
func (this *QGraphicsTextItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10opaqueAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:903
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type()
func (this *QGraphicsTextItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:905
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextWidth(qreal)
func (this *QGraphicsTextItem) SetTextWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setTextWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:906
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal textWidth()
func (this *QGraphicsTextItem) TextWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem9textWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:908
// index:0
// Public Visibility=Default Availability=Available
// [-2] void adjustSize()
func (this *QGraphicsTextItem) AdjustSize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem10adjustSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:910
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocument(QTextDocument *)
func (this *QGraphicsTextItem) SetDocument(document qtgui.QTextDocument_ITF /*777 QTextDocument **/) {
	var convArg0 = document.QTextDocument_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem11setDocumentEP13QTextDocument", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:911
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document()
func (this *QGraphicsTextItem) Document() *qtgui.QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:913
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QGraphicsTextItem) SetTextInteractionFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:914
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextInteractionFlags textInteractionFlags()
func (this *QGraphicsTextItem) TextInteractionFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem20textInteractionFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:916
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabChangesFocus(_Bool)
func (this *QGraphicsTextItem) SetTabChangesFocus(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem18setTabChangesFocusEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:917
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabChangesFocus()
func (this *QGraphicsTextItem) TabChangesFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem15tabChangesFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:919
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpenExternalLinks(_Bool)
func (this *QGraphicsTextItem) SetOpenExternalLinks(open bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem20setOpenExternalLinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), open)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:920
// index:0
// Public Visibility=Default Availability=Available
// [1] bool openExternalLinks()
func (this *QGraphicsTextItem) OpenExternalLinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem17openExternalLinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:922
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextCursor(const QTextCursor &)
func (this *QGraphicsTextItem) SetTextCursor(cursor qtgui.QTextCursor_ITF) {
	var convArg0 = cursor.QTextCursor_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:923
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor textCursor()
func (this *QGraphicsTextItem) TextCursor() *qtgui.QTextCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10textCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:926
// index:0
// Public Visibility=Default Availability=Available
// [-2] void linkActivated(const QString &)
func (this *QGraphicsTextItem) LinkActivated(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem13linkActivatedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:927
// index:0
// Public Visibility=Default Availability=Available
// [-2] void linkHovered(const QString &)
func (this *QGraphicsTextItem) LinkHovered(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem11linkHoveredERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:930
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool sceneEvent(QEvent *)
func (this *QGraphicsTextItem) SceneEvent(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 = event.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem10sceneEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:931
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) MousePressEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem15mousePressEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:932
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) MouseMoveEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem14mouseMoveEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:933
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) MouseReleaseEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:934
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) MouseDoubleClickEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:935
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsTextItem) ContextMenuEvent(event QGraphicsSceneContextMenuEvent_ITF /*777 QGraphicsSceneContextMenuEvent **/) {
	var convArg0 = event.QGraphicsSceneContextMenuEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:936
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QGraphicsTextItem) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 = event.QKeyEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:937
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QGraphicsTextItem) KeyReleaseEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 = event.QKeyEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:938
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QGraphicsTextItem) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 = event.QFocusEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:939
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QGraphicsTextItem) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 = event.QFocusEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:940
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) DragEnterEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem14dragEnterEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:941
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) DragLeaveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:942
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) DragMoveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem13dragMoveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:943
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) DropEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem9dropEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:944
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)
func (this *QGraphicsTextItem) InputMethodEvent(event qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 = event.QInputMethodEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:945
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverEnterEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) HoverEnterEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem15hoverEnterEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:946
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverMoveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) HoverMoveEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem14hoverMoveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:947
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverLeaveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) HoverLeaveEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:949
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsTextItem) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:951
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsTextItem) SupportsExtension(extension int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:952
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(enum QGraphicsItem::Extension, const QVariant &)
func (this *QGraphicsTextItem) SetExtension(extension int, variant qtcore.QVariant_ITF) {
	var convArg1 = variant.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:953
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant extension(const QVariant &)
func (this *QGraphicsTextItem) Extension(variant qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg0 = variant.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsTextItem9extensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

type QGraphicsTextItem__ = int

const QGraphicsTextItem__Type QGraphicsTextItem__ = 8

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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
