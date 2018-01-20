//  header block begin
// /usr/include/qt/QtWidgets/qitemdelegate.h
// #include <qitemdelegate.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 60
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
type QItemDelegate struct {
	*QAbstractItemDelegate
}

func (this *QItemDelegate) GetCthis() unsafe.Pointer {
	return this.QAbstractItemDelegate.GetCthis()
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QItemDelegate) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:62
// index:0
// void QItemDelegate(class QObject *)
func NewQItemDelegate(parent unsafe.Pointer) *QItemDelegate {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegateC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemDelegateFromPointer(cthis)
	return gothis
}
func NewQItemDelegateFromPointer(cthis unsafe.Pointer) *QItemDelegate {
	bcthis0 := NewQAbstractItemDelegateFromPointer(cthis)
	return &QItemDelegate{bcthis0}
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:63
// index:0
// virtual
// void ~QItemDelegate()
func DeleteQItemDelegate(*QItemDelegate) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:65
// index:0
// bool hasClipping()
func (this *QItemDelegate) HasClipping() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate11hasClippingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:66
// index:0
// void setClipping(_Bool)
func (this *QItemDelegate) SetClipping(clip bool) {
	// 0: (, clip bool), (&clip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate11setClippingEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &clip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:69
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) Paint(painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, painter QPainter *, option const QStyleOptionViewItem &, index const QModelIndex &), (painter, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:72
// index:0
// virtual
// QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) SizeHint(option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, option const QStyleOptionViewItem &, index const QModelIndex &), (option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:76
// index:0
// virtual
// QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) CreateEditor(parent unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, parent QWidget *, option const QStyleOptionViewItem &, index const QModelIndex &), (parent, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:80
// index:0
// virtual
// void setEditorData(class QWidget *, const class QModelIndex &)
func (this *QItemDelegate) SetEditorData(editor unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, editor QWidget *, index const QModelIndex &), (editor, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:81
// index:0
// virtual
// void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
func (this *QItemDelegate) SetModelData(editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, editor QWidget *, model QAbstractItemModel *, index const QModelIndex &), (editor, model, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, model, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:83
// index:0
// virtual
// void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) UpdateEditorGeometry(editor unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, editor QWidget *, option const QStyleOptionViewItem &, index const QModelIndex &), (editor, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:88
// index:0
// QItemEditorFactory * itemEditorFactory()
func (this *QItemDelegate) ItemEditorFactory() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate17itemEditorFactoryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:89
// index:0
// void setItemEditorFactory(class QItemEditorFactory *)
func (this *QItemDelegate) SetItemEditorFactory(factory unsafe.Pointer) {
	// 0: (, factory QItemEditorFactory *), (factory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", ffiqt.FFI_TYPE_VOID, this.GetCthis(), factory)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:92
// index:0
// virtual
// void drawDisplay(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, const class QString &)
func (this *QItemDelegate) DrawDisplay(painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer, text unsafe.Pointer) {
	// 0: (, painter QPainter *, option const QStyleOptionViewItem &, rect const QRect &, text const QString &), (painter, option, rect, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate11drawDisplayEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, rect, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:94
// index:0
// virtual
// void drawDecoration(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, const class QPixmap &)
func (this *QItemDelegate) DrawDecoration(painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer, pixmap unsafe.Pointer) {
	// 0: (, painter QPainter *, option const QStyleOptionViewItem &, rect const QRect &, pixmap const QPixmap &), (painter, option, rect, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate14drawDecorationEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, rect, pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:96
// index:0
// virtual
// void drawFocus(class QPainter *, const class QStyleOptionViewItem &, const class QRect &)
func (this *QItemDelegate) DrawFocus(painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer) {
	// 0: (, painter QPainter *, option const QStyleOptionViewItem &, rect const QRect &), (painter, option, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate9drawFocusEP8QPainterRK20QStyleOptionViewItemRK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:98
// index:0
// virtual
// void drawCheck(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, Qt::CheckState)
func (this *QItemDelegate) DrawCheck(painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer, state int) {
	// 0: (, painter QPainter *, option const QStyleOptionViewItem &, rect const QRect &, state Qt::CheckState), (painter, option, rect, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate9drawCheckEP8QPainterRK20QStyleOptionViewItemRK5QRectN2Qt10CheckStateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, rect, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:100
// index:0
// void drawBackground(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) DrawBackground(painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, painter QPainter *, option const QStyleOptionViewItem &, index const QModelIndex &), (painter, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate14drawBackgroundEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:103
// index:0
// void doLayout(const class QStyleOptionViewItem &, class QRect *, class QRect *, class QRect *, _Bool)
func (this *QItemDelegate) DoLayout(option unsafe.Pointer, checkRect unsafe.Pointer, iconRect unsafe.Pointer, textRect unsafe.Pointer, hint bool) {
	// 0: (, option const QStyleOptionViewItem &, checkRect QRect *, iconRect QRect *, textRect QRect *, hint bool), (option, checkRect, iconRect, textRect, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate8doLayoutERK20QStyleOptionViewItemP5QRectS4_S4_b", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option, checkRect, iconRect, textRect, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:106
// index:0
// QRect rect(const class QStyleOptionViewItem &, const class QModelIndex &, int)
func (this *QItemDelegate) Rect(option unsafe.Pointer, index unsafe.Pointer, role int) {
	// 0: (, option const QStyleOptionViewItem &, index const QModelIndex &, role int), (option, index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate4rectERK20QStyleOptionViewItemRK11QModelIndexi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option, index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:108
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QItemDelegate) EventFilter(object unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, object QObject *, event QEvent *), (object, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), object, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:109
// index:0
// virtual
// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) EditorEvent(event unsafe.Pointer, model unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, event QEvent *, model QAbstractItemModel *, option const QStyleOptionViewItem &, index const QModelIndex &), (event, model, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event, model, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:112
// index:0
// QStyleOptionViewItem setOptions(const class QModelIndex &, const class QStyleOptionViewItem &)
func (this *QItemDelegate) SetOptions(index unsafe.Pointer, option unsafe.Pointer) {
	// 0: (, index const QModelIndex &, option const QStyleOptionViewItem &), (index, option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate10setOptionsERK11QModelIndexRK20QStyleOptionViewItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:115
// index:0
// QPixmap decoration(const class QStyleOptionViewItem &, const class QVariant &)
func (this *QItemDelegate) Decoration(option unsafe.Pointer, variant unsafe.Pointer) {
	// 0: (, option const QStyleOptionViewItem &, variant const QVariant &), (option, variant)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate10decorationERK20QStyleOptionViewItemRK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option, variant)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:116
// index:0
// QPixmap * selected(const class QPixmap &, const class QPalette &, _Bool)
func (this *QItemDelegate) Selected(pixmap unsafe.Pointer, palette unsafe.Pointer, enabled bool) {
	// 0: (, pixmap const QPixmap &, palette const QPalette &, enabled bool), (pixmap, palette, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate8selectedERK7QPixmapRK8QPaletteb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pixmap, palette, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:118
// index:0
// QRect doCheck(const class QStyleOptionViewItem &, const class QRect &, const class QVariant &)
func (this *QItemDelegate) DoCheck(option unsafe.Pointer, bounding unsafe.Pointer, variant unsafe.Pointer) {
	// 0: (, option const QStyleOptionViewItem &, bounding const QRect &, variant const QVariant &), (option, bounding, variant)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate7doCheckERK20QStyleOptionViewItemRK5QRectRK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option, bounding, variant)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:120
// index:0
// QRect textRectangle(class QPainter *, const class QRect &, const class QFont &, const class QString &)
func (this *QItemDelegate) TextRectangle(painter unsafe.Pointer, rect unsafe.Pointer, font unsafe.Pointer, text unsafe.Pointer) {
	// 0: (, painter QPainter *, rect const QRect &, font const QFont &, text const QString &), (painter, rect, font, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate13textRectangleEP8QPainterRK5QRectRK5QFontRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect, font, text)
	gopp.ErrPrint(err, rv)
}

//  body block end
