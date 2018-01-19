//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:872
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsTextItem) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:877
// index:0
// void QGraphicsTextItem(class QGraphicsItem *)
func NewQGraphicsTextItem(parent unsafe.Pointer) *QGraphicsTextItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsTextItem{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:878
// index:1
// void QGraphicsTextItem(const class QString &, class QGraphicsItem *)
func NewQGraphicsTextItem_1(text unsafe.Pointer, parent unsafe.Pointer) *QGraphicsTextItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItemC2ERK7QStringP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsTextItem{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem6toHtmlEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:882
// index:0
// void setHtml(const class QString &)
func (this *QGraphicsTextItem) SetHtml(html unsafe.Pointer) {
	// 0: (, const QString & html), (html)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem7setHtmlERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, html)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:884
// index:0
// QString toPlainText()
func (this *QGraphicsTextItem) ToPlainText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem11toPlainTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:885
// index:0
// void setPlainText(const class QString &)
func (this *QGraphicsTextItem) SetPlainText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:887
// index:0
// QFont font()
func (this *QGraphicsTextItem) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem4fontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:888
// index:0
// void setFont(const class QFont &)
func (this *QGraphicsTextItem) SetFont(font unsafe.Pointer) {
	// 0: (, const QFont & font), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:890
// index:0
// void setDefaultTextColor(const class QColor &)
func (this *QGraphicsTextItem) SetDefaultTextColor(c unsafe.Pointer) {
	// 0: (, const QColor & c), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem19setDefaultTextColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:891
// index:0
// QColor defaultTextColor()
func (this *QGraphicsTextItem) DefaultTextColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem16defaultTextColorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:893
// index:0
// virtual
// QRectF boundingRect()
func (this *QGraphicsTextItem) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:894
// index:0
// virtual
// QPainterPath shape()
func (this *QGraphicsTextItem) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem5shapeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:895
// index:0
// virtual
// bool contains(const class QPointF &)
func (this *QGraphicsTextItem) Contains(point unsafe.Pointer) {
	// 0: (, const QPointF & point), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem8containsERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:897
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsTextItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:899
// index:0
// virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsTextItem) IsObscuredBy(item unsafe.Pointer) {
	// 0: (, const QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:900
// index:0
// virtual
// QPainterPath opaqueArea()
func (this *QGraphicsTextItem) OpaqueArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10opaqueAreaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:903
// index:0
// virtual
// int type()
func (this *QGraphicsTextItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:905
// index:0
// void setTextWidth(qreal)
func (this *QGraphicsTextItem) SetTextWidth(width float64) {
	// 0: (, qreal width), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem12setTextWidthEd", ffiqt.FFI_TYPE_VOID, this.cthis, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:906
// index:0
// qreal textWidth()
func (this *QGraphicsTextItem) TextWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem9textWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:908
// index:0
// void adjustSize()
func (this *QGraphicsTextItem) AdjustSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem10adjustSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:910
// index:0
// void setDocument(class QTextDocument *)
func (this *QGraphicsTextItem) SetDocument(document unsafe.Pointer) {
	// 0: (, QTextDocument * document), (document)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem11setDocumentEP13QTextDocument", ffiqt.FFI_TYPE_VOID, this.cthis, document)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:911
// index:0
// QTextDocument * document()
func (this *QGraphicsTextItem) Document() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem8documentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:914
// index:0
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QGraphicsTextItem) TextInteractionFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem20textInteractionFlagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:916
// index:0
// void setTabChangesFocus(_Bool)
func (this *QGraphicsTextItem) SetTabChangesFocus(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem18setTabChangesFocusEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:917
// index:0
// bool tabChangesFocus()
func (this *QGraphicsTextItem) TabChangesFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem15tabChangesFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:919
// index:0
// void setOpenExternalLinks(_Bool)
func (this *QGraphicsTextItem) SetOpenExternalLinks(open bool) {
	// 0: (, bool open), (&open)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem20setOpenExternalLinksEb", ffiqt.FFI_TYPE_VOID, this.cthis, &open)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:920
// index:0
// bool openExternalLinks()
func (this *QGraphicsTextItem) OpenExternalLinks() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem17openExternalLinksEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:922
// index:0
// void setTextCursor(const class QTextCursor &)
func (this *QGraphicsTextItem) SetTextCursor(cursor unsafe.Pointer) {
	// 0: (, const QTextCursor & cursor), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem13setTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.cthis, cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:923
// index:0
// QTextCursor textCursor()
func (this *QGraphicsTextItem) TextCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsTextItem10textCursorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:926
// index:0
// void linkActivated(const class QString &)
func (this *QGraphicsTextItem) LinkActivated(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem13linkActivatedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:927
// index:0
// void linkHovered(const class QString &)
func (this *QGraphicsTextItem) LinkHovered(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsTextItem11linkHoveredERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
