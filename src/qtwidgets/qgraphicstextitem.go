//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QGraphicsTextItem struct {
	*qtrt.CObject
}

func (this *QGraphicsTextItem) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:872
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsTextItem) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:877
// index:0
// void QGraphicsTextItem(class QGraphicsItem *)
func NewQGraphicsTextItem(parent unsafe.Pointer) *QGraphicsTextItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsTextItemFromPointer(cthis)
	return gothis
}
func NewQGraphicsTextItemFromPointer(cthis unsafe.Pointer) *QGraphicsTextItem {
	return &QGraphicsTextItem{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:878
// index:1
// void QGraphicsTextItem(const class QString &, class QGraphicsItem *)
func NewQGraphicsTextItem_1(text unsafe.Pointer, parent unsafe.Pointer) *QGraphicsTextItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItemC2ERK7QStringP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsTextItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:879
// index:0
// virtual
// void ~QGraphicsTextItem()
func DeleteQGraphicsTextItem(*QGraphicsTextItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:881
// index:0
// QString toHtml()
func (this *QGraphicsTextItem) ToHtml() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem6toHtmlEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:882
// index:0
// void setHtml(const class QString &)
func (this *QGraphicsTextItem) SetHtml(html unsafe.Pointer) {
	// 0: (, html const QString &), (html)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem7setHtmlERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), html)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:884
// index:0
// QString toPlainText()
func (this *QGraphicsTextItem) ToPlainText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem11toPlainTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:885
// index:0
// void setPlainText(const class QString &)
func (this *QGraphicsTextItem) SetPlainText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:887
// index:0
// QFont font()
func (this *QGraphicsTextItem) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem4fontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:888
// index:0
// void setFont(const class QFont &)
func (this *QGraphicsTextItem) SetFont(font unsafe.Pointer) {
	// 0: (, font const QFont &), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:890
// index:0
// void setDefaultTextColor(const class QColor &)
func (this *QGraphicsTextItem) SetDefaultTextColor(c unsafe.Pointer) {
	// 0: (, c const QColor &), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:891
// index:0
// QColor defaultTextColor()
func (this *QGraphicsTextItem) DefaultTextColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem16defaultTextColorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:893
// index:0
// virtual
// QRectF boundingRect()
func (this *QGraphicsTextItem) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:894
// index:0
// virtual
// QPainterPath shape()
func (this *QGraphicsTextItem) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem5shapeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:895
// index:0
// virtual
// bool contains(const class QPointF &)
func (this *QGraphicsTextItem) Contains(point unsafe.Pointer) {
	// 0: (, point const QPointF &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem8containsERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:897
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsTextItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, painter QPainter *, option const QStyleOptionGraphicsItem *, widget QWidget *), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:899
// index:0
// virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsTextItem) IsObscuredBy(item unsafe.Pointer) {
	// 0: (, item const QGraphicsItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:900
// index:0
// virtual
// QPainterPath opaqueArea()
func (this *QGraphicsTextItem) OpaqueArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10opaqueAreaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:903
// index:0
// virtual
// int type()
func (this *QGraphicsTextItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:905
// index:0
// void setTextWidth(qreal)
func (this *QGraphicsTextItem) SetTextWidth(width float64) {
	// 0: (, width qreal), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setTextWidthEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:906
// index:0
// qreal textWidth()
func (this *QGraphicsTextItem) TextWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem9textWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:908
// index:0
// void adjustSize()
func (this *QGraphicsTextItem) AdjustSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem10adjustSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:910
// index:0
// void setDocument(class QTextDocument *)
func (this *QGraphicsTextItem) SetDocument(document unsafe.Pointer) {
	// 0: (, document QTextDocument *), (document)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem11setDocumentEP13QTextDocument", ffiqt.FFI_TYPE_VOID, this.GetCthis(), document)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:911
// index:0
// QTextDocument * document()
func (this *QGraphicsTextItem) Document() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem8documentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:913
// index:0
// void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QGraphicsTextItem) SetTextInteractionFlags(flags int) {
	// 0: (, QFlags<Qt::TextInteractionFlag> flags), (&flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:914
// index:0
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QGraphicsTextItem) TextInteractionFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem20textInteractionFlagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:916
// index:0
// void setTabChangesFocus(_Bool)
func (this *QGraphicsTextItem) SetTabChangesFocus(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem18setTabChangesFocusEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:917
// index:0
// bool tabChangesFocus()
func (this *QGraphicsTextItem) TabChangesFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem15tabChangesFocusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:919
// index:0
// void setOpenExternalLinks(_Bool)
func (this *QGraphicsTextItem) SetOpenExternalLinks(open bool) {
	// 0: (, open bool), (&open)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem20setOpenExternalLinksEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &open)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:920
// index:0
// bool openExternalLinks()
func (this *QGraphicsTextItem) OpenExternalLinks() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem17openExternalLinksEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:922
// index:0
// void setTextCursor(const class QTextCursor &)
func (this *QGraphicsTextItem) SetTextCursor(cursor unsafe.Pointer) {
	// 0: (, cursor const QTextCursor &), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:923
// index:0
// QTextCursor textCursor()
func (this *QGraphicsTextItem) TextCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10textCursorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:926
// index:0
// void linkActivated(const class QString &)
func (this *QGraphicsTextItem) LinkActivated(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem13linkActivatedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:927
// index:0
// void linkHovered(const class QString &)
func (this *QGraphicsTextItem) LinkHovered(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem11linkHoveredERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:930
// index:0
// virtual
// bool sceneEvent(class QEvent *)
func (this *QGraphicsTextItem) SceneEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem10sceneEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:931
// index:0
// virtual
// void mousePressEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) MousePressEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem15mousePressEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:932
// index:0
// virtual
// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) MouseMoveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem14mouseMoveEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:933
// index:0
// virtual
// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) MouseReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:934
// index:0
// virtual
// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsTextItem) MouseDoubleClickEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:935
// index:0
// virtual
// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsTextItem) ContextMenuEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneContextMenuEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:936
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsTextItem) KeyPressEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:937
// index:0
// virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsTextItem) KeyReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:938
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsTextItem) FocusInEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:939
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsTextItem) FocusOutEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:940
// index:0
// virtual
// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) DragEnterEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem14dragEnterEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:941
// index:0
// virtual
// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) DragLeaveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:942
// index:0
// virtual
// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) DragMoveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem13dragMoveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:943
// index:0
// virtual
// void dropEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsTextItem) DropEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem9dropEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:944
// index:0
// virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsTextItem) InputMethodEvent(event unsafe.Pointer) {
	// 0: (, event QInputMethodEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:945
// index:0
// virtual
// void hoverEnterEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) HoverEnterEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneHoverEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem15hoverEnterEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:946
// index:0
// virtual
// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) HoverMoveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneHoverEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem14hoverMoveEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:947
// index:0
// virtual
// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsTextItem) HoverLeaveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneHoverEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:949
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsTextItem) InputMethodQuery(query int) {
	// 0: (, query Qt::InputMethodQuery), (&query)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &query)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:951
// index:0
// virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsTextItem) SupportsExtension(extension int) {
	// 0: (, extension QGraphicsItem::Extension), (&extension)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &extension)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:952
// index:0
// virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsTextItem) SetExtension(extension int, variant unsafe.Pointer) {
	// 0: (, extension QGraphicsItem::Extension, variant const QVariant &), (&extension, variant)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &extension, variant)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:953
// index:0
// virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsTextItem) Extension(variant unsafe.Pointer) {
	// 0: (, variant const QVariant &), (variant)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem9extensionERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), variant)
	gopp.ErrPrint(err, rv)
}

//  body block end
