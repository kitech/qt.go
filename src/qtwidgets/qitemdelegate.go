//  header block begin
// /usr/include/qt/QtWidgets/qitemdelegate.h
// #include <qitemdelegate.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 52
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
func NewQItemDelegateFromPointer(cthis unsafe.Pointer) *QItemDelegate {
	bcthis0 := NewQAbstractItemDelegateFromPointer(cthis)
	return &QItemDelegate{bcthis0}
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QItemDelegate) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:62
// index:0
// Public
// void QItemDelegate(class QObject *)
func NewQItemDelegate(parent unsafe.Pointer) *QItemDelegate {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegateC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemDelegateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:63
// index:0
// Public virtual
// void ~QItemDelegate()
func DeleteQItemDelegate(*QItemDelegate) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:65
// index:0
// Public
// bool hasClipping()
func (this *QItemDelegate) HasClipping() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate11hasClippingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:66
// index:0
// Public
// void setClipping(_Bool)
func (this *QItemDelegate) SetClipping(clip bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate11setClippingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &clip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:69
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) Paint(painter unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:72
// index:0
// Public virtual
// QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) SizeHint(option *QStyleOptionViewItem, index *qtcore.QModelIndex) interface{} {
	var convArg0 = option.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:76
// index:0
// Public virtual
// QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) CreateEditor(parent unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) interface{} {
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), parent, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:80
// index:0
// Public virtual
// void setEditorData(class QWidget *, const class QModelIndex &)
func (this *QItemDelegate) SetEditorData(editor unsafe.Pointer, index *qtcore.QModelIndex) {
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:81
// index:0
// Public virtual
// void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
func (this *QItemDelegate) SetModelData(editor unsafe.Pointer, model unsafe.Pointer, index *qtcore.QModelIndex) {
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, model, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:83
// index:0
// Public virtual
// void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) UpdateEditorGeometry(editor unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:88
// index:0
// Public
// QItemEditorFactory * itemEditorFactory()
func (this *QItemDelegate) ItemEditorFactory() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate17itemEditorFactoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:89
// index:0
// Public
// void setItemEditorFactory(class QItemEditorFactory *)
func (this *QItemDelegate) SetItemEditorFactory(factory unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), factory)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:92
// index:0
// Protected virtual
// void drawDisplay(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, const class QString &)
func (this *QItemDelegate) DrawDisplay(painter unsafe.Pointer, option *QStyleOptionViewItem, rect *qtcore.QRect, text *qtcore.QString) {
	var convArg1 = option.GetCthis()
	var convArg2 = rect.GetCthis()
	var convArg3 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate11drawDisplayEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:94
// index:0
// Protected virtual
// void drawDecoration(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, const class QPixmap &)
func (this *QItemDelegate) DrawDecoration(painter unsafe.Pointer, option *QStyleOptionViewItem, rect *qtcore.QRect, pixmap *qtgui.QPixmap) {
	var convArg1 = option.GetCthis()
	var convArg2 = rect.GetCthis()
	var convArg3 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate14drawDecorationEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:96
// index:0
// Protected virtual
// void drawFocus(class QPainter *, const class QStyleOptionViewItem &, const class QRect &)
func (this *QItemDelegate) DrawFocus(painter unsafe.Pointer, option *QStyleOptionViewItem, rect *qtcore.QRect) {
	var convArg1 = option.GetCthis()
	var convArg2 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate9drawFocusEP8QPainterRK20QStyleOptionViewItemRK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:98
// index:0
// Protected virtual
// void drawCheck(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, Qt::CheckState)
func (this *QItemDelegate) DrawCheck(painter unsafe.Pointer, option *QStyleOptionViewItem, rect *qtcore.QRect, state int) {
	var convArg1 = option.GetCthis()
	var convArg2 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate9drawCheckEP8QPainterRK20QStyleOptionViewItemRK5QRectN2Qt10CheckStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:100
// index:0
// Protected
// void drawBackground(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) DrawBackground(painter unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate14drawBackgroundEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:103
// index:0
// Protected
// void doLayout(const class QStyleOptionViewItem &, class QRect *, class QRect *, class QRect *, _Bool)
func (this *QItemDelegate) DoLayout(option *QStyleOptionViewItem, checkRect unsafe.Pointer, iconRect unsafe.Pointer, textRect unsafe.Pointer, hint bool) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate8doLayoutERK20QStyleOptionViewItemP5QRectS4_S4_b", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, checkRect, iconRect, textRect, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:106
// index:0
// Protected
// QRect rect(const class QStyleOptionViewItem &, const class QModelIndex &, int)
func (this *QItemDelegate) Rect(option *QStyleOptionViewItem, index *qtcore.QModelIndex, role int) interface{} {
	var convArg0 = option.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate4rectERK20QStyleOptionViewItemRK11QModelIndexi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:108
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QItemDelegate) EventFilter(object unsafe.Pointer, event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object, event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:109
// index:0
// Protected virtual
// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) EditorEvent(event unsafe.Pointer, model unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) interface{} {
	var convArg2 = option.GetCthis()
	var convArg3 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event, model, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:112
// index:0
// Protected
// QStyleOptionViewItem setOptions(const class QModelIndex &, const class QStyleOptionViewItem &)
func (this *QItemDelegate) SetOptions(index *qtcore.QModelIndex, option *QStyleOptionViewItem) interface{} {
	var convArg0 = index.GetCthis()
	var convArg1 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate10setOptionsERK11QModelIndexRK20QStyleOptionViewItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:115
// index:0
// Protected
// QPixmap decoration(const class QStyleOptionViewItem &, const class QVariant &)
func (this *QItemDelegate) Decoration(option *QStyleOptionViewItem, variant *qtcore.QVariant) interface{} {
	var convArg0 = option.GetCthis()
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate10decorationERK20QStyleOptionViewItemRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:116
// index:0
// Protected
// QPixmap * selected(const class QPixmap &, const class QPalette &, _Bool)
func (this *QItemDelegate) Selected(pixmap *qtgui.QPixmap, palette *qtgui.QPalette, enabled bool) interface{} {
	var convArg0 = pixmap.GetCthis()
	var convArg1 = palette.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate8selectedERK7QPixmapRK8QPaletteb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &enabled)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:118
// index:0
// Protected
// QRect doCheck(const class QStyleOptionViewItem &, const class QRect &, const class QVariant &)
func (this *QItemDelegate) DoCheck(option *QStyleOptionViewItem, bounding *qtcore.QRect, variant *qtcore.QVariant) interface{} {
	var convArg0 = option.GetCthis()
	var convArg1 = bounding.GetCthis()
	var convArg2 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate7doCheckERK20QStyleOptionViewItemRK5QRectRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:120
// index:0
// Protected
// QRect textRectangle(class QPainter *, const class QRect &, const class QFont &, const class QString &)
func (this *QItemDelegate) TextRectangle(painter unsafe.Pointer, rect *qtcore.QRect, font *qtgui.QFont, text *qtcore.QString) interface{} {
	var convArg1 = rect.GetCthis()
	var convArg2 = font.GetCthis()
	var convArg3 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate13textRectangleEP8QPainterRK5QRectRK5QFontRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
