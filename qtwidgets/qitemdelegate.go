package qtwidgets

// /usr/include/qt/QtWidgets/qitemdelegate.h
// #include <qitemdelegate.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 60
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

// void drawDisplay(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, const class QString &)
func (this *QItemDelegate) InheritDrawDisplay(f func(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect, text string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawDisplay", f)
}

// void drawDecoration(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, const class QPixmap &)
func (this *QItemDelegate) InheritDrawDecoration(f func(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect, pixmap *qtgui.QPixmap) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawDecoration", f)
}

// void drawFocus(class QPainter *, const class QStyleOptionViewItem &, const class QRect &)
func (this *QItemDelegate) InheritDrawFocus(f func(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawFocus", f)
}

// void drawCheck(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, Qt::CheckState)
func (this *QItemDelegate) InheritDrawCheck(f func(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect, state int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawCheck", f)
}

// void drawBackground(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) InheritDrawBackground(f func(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawBackground", f)
}

// void doLayout(const class QStyleOptionViewItem &, class QRect *, class QRect *, class QRect *, _Bool)
func (this *QItemDelegate) InheritDoLayout(f func(option *QStyleOptionViewItem, checkRect *qtcore.QRect /*777 QRect **/, iconRect *qtcore.QRect /*777 QRect **/, textRect *qtcore.QRect /*777 QRect **/, hint bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "doLayout", f)
}

// QRect rect(const class QStyleOptionViewItem &, const class QModelIndex &, int)
func (this *QItemDelegate) InheritRect(f func(option *QStyleOptionViewItem, index *qtcore.QModelIndex, role int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "rect", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QItemDelegate) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) InheritEditorEvent(f func(event *qtcore.QEvent /*777 QEvent **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "editorEvent", f)
}

// QStyleOptionViewItem setOptions(const class QModelIndex &, const class QStyleOptionViewItem &)
func (this *QItemDelegate) InheritSetOptions(f func(index *qtcore.QModelIndex, option *QStyleOptionViewItem) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "setOptions", f)
}

// QPixmap decoration(const class QStyleOptionViewItem &, const class QVariant &)
func (this *QItemDelegate) InheritDecoration(f func(option *QStyleOptionViewItem, variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "decoration", f)
}

// QPixmap * selected(const class QPixmap &, const class QPalette &, _Bool)
func (this *QItemDelegate) InheritSelected(f func(pixmap *qtgui.QPixmap, palette *qtgui.QPalette, enabled bool) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "selected", f)
}

// QRect doCheck(const class QStyleOptionViewItem &, const class QRect &, const class QVariant &)
func (this *QItemDelegate) InheritDoCheck(f func(option *QStyleOptionViewItem, bounding *qtcore.QRect, variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "doCheck", f)
}

// QRect textRectangle(class QPainter *, const class QRect &, const class QFont &, const class QString &)
func (this *QItemDelegate) InheritTextRectangle(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, font *qtgui.QFont, text string) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "textRectangle", f)
}

type QItemDelegate struct {
	*QAbstractItemDelegate
}
type QItemDelegate_ITF interface {
	QAbstractItemDelegate_ITF
	QItemDelegate_PTR() *QItemDelegate
}

func (ptr *QItemDelegate) QItemDelegate_PTR() *QItemDelegate { return ptr }

func (this *QItemDelegate) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemDelegate.GetCthis()
	}
}
func (this *QItemDelegate) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemDelegate = NewQAbstractItemDelegateFromPointer(cthis)
}
func NewQItemDelegateFromPointer(cthis unsafe.Pointer) *QItemDelegate {
	bcthis0 := NewQAbstractItemDelegateFromPointer(cthis)
	return &QItemDelegate{bcthis0}
}
func (*QItemDelegate) NewFromPointer(cthis unsafe.Pointer) *QItemDelegate {
	return NewQItemDelegateFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QItemDelegate) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QItemDelegate(QObject *)
func NewQItemDelegate(parent qtcore.QObject_ITF /*777 QObject **/) *QItemDelegate {
	var convArg0 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QItemDelegate()
func DeleteQItemDelegate(this *QItemDelegate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasClipping()
func (this *QItemDelegate) HasClipping() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate11hasClippingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipping(_Bool)
func (this *QItemDelegate) SetClipping(clip bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegate11setClippingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clip)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 = painter.QPainter_PTR().GetCthis()
	var convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg2 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint(const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) SizeHint(option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *qtcore.QSize /*123*/ {
	var convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg1 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(QWidget *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) CreateEditor(parent QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *QWidget /*777 QWidget **/ {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	var convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg2 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setEditorData(QWidget *, const QModelIndex &)
func (this *QItemDelegate) SetEditorData(editor QWidget_ITF /*777 QWidget **/, index qtcore.QModelIndex_ITF) {
	var convArg0 = editor.QWidget_PTR().GetCthis()
	var convArg1 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModelData(QWidget *, QAbstractItemModel *, const QModelIndex &)
func (this *QItemDelegate) SetModelData(editor QWidget_ITF /*777 QWidget **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, index qtcore.QModelIndex_ITF) {
	var convArg0 = editor.QWidget_PTR().GetCthis()
	var convArg1 = model.QAbstractItemModel_PTR().GetCthis()
	var convArg2 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometry(QWidget *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) UpdateEditorGeometry(editor QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 = editor.QWidget_PTR().GetCthis()
	var convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg2 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QItemEditorFactory * itemEditorFactory()
func (this *QItemDelegate) ItemEditorFactory() *QItemEditorFactory /*777 QItemEditorFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate17itemEditorFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQItemEditorFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemEditorFactory(QItemEditorFactory *)
func (this *QItemDelegate) SetItemEditorFactory(factory QItemEditorFactory_ITF /*777 QItemEditorFactory **/) {
	var convArg0 = factory.QItemEditorFactory_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:92
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawDisplay(QPainter *, const QStyleOptionViewItem &, const QRect &, const QString &)
func (this *QItemDelegate) DrawDisplay(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, rect qtcore.QRect_ITF, text string) {
	var convArg0 = painter.QPainter_PTR().GetCthis()
	var convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg2 = rect.QRect_PTR().GetCthis()
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate11drawDisplayEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawDecoration(QPainter *, const QStyleOptionViewItem &, const QRect &, const QPixmap &)
func (this *QItemDelegate) DrawDecoration(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, rect qtcore.QRect_ITF, pixmap qtgui.QPixmap_ITF) {
	var convArg0 = painter.QPainter_PTR().GetCthis()
	var convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg2 = rect.QRect_PTR().GetCthis()
	var convArg3 = pixmap.QPixmap_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate14drawDecorationEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawFocus(QPainter *, const QStyleOptionViewItem &, const QRect &)
func (this *QItemDelegate) DrawFocus(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, rect qtcore.QRect_ITF) {
	var convArg0 = painter.QPainter_PTR().GetCthis()
	var convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg2 = rect.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate9drawFocusEP8QPainterRK20QStyleOptionViewItemRK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:98
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawCheck(QPainter *, const QStyleOptionViewItem &, const QRect &, Qt::CheckState)
func (this *QItemDelegate) DrawCheck(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, rect qtcore.QRect_ITF, state int) {
	var convArg0 = painter.QPainter_PTR().GetCthis()
	var convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg2 = rect.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate9drawCheckEP8QPainterRK20QStyleOptionViewItemRK5QRectN2Qt10CheckStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:100
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawBackground(QPainter *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) DrawBackground(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 = painter.QPainter_PTR().GetCthis()
	var convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg2 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate14drawBackgroundEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:103
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void doLayout(const QStyleOptionViewItem &, QRect *, QRect *, QRect *, _Bool)
func (this *QItemDelegate) DoLayout(option QStyleOptionViewItem_ITF, checkRect qtcore.QRect_ITF /*777 QRect **/, iconRect qtcore.QRect_ITF /*777 QRect **/, textRect qtcore.QRect_ITF /*777 QRect **/, hint bool) {
	var convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg1 = checkRect.QRect_PTR().GetCthis()
	var convArg2 = iconRect.QRect_PTR().GetCthis()
	var convArg3 = textRect.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate8doLayoutERK20QStyleOptionViewItemP5QRectS4_S4_b", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect rect(const QStyleOptionViewItem &, const QModelIndex &, int)
func (this *QItemDelegate) Rect(option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF, role int) *qtcore.QRect /*123*/ {
	var convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg1 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate4rectERK20QStyleOptionViewItemRK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:108
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QItemDelegate) EventFilter(object qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 = object.QObject_PTR().GetCthis()
	var convArg1 = event.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegate11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) EditorEvent(event qtcore.QEvent_ITF /*777 QEvent **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) bool {
	var convArg0 = event.QEvent_PTR().GetCthis()
	var convArg1 = model.QAbstractItemModel_PTR().GetCthis()
	var convArg2 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg3 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:112
// index:0
// Protected Visibility=Default Availability=Available
// [192] QStyleOptionViewItem setOptions(const QModelIndex &, const QStyleOptionViewItem &)
func (this *QItemDelegate) SetOptions(index qtcore.QModelIndex_ITF, option QStyleOptionViewItem_ITF) *QStyleOptionViewItem /*123*/ {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	var convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate10setOptionsERK11QModelIndexRK20QStyleOptionViewItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStyleOptionViewItem)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:115
// index:0
// Protected Visibility=Default Availability=Available
// [32] QPixmap decoration(const QStyleOptionViewItem &, const QVariant &)
func (this *QItemDelegate) Decoration(option QStyleOptionViewItem_ITF, variant qtcore.QVariant_ITF) *qtgui.QPixmap /*123*/ {
	var convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg1 = variant.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate10decorationERK20QStyleOptionViewItemRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:116
// index:0
// Protected Visibility=Default Availability=Available
// [8] QPixmap * selected(const QPixmap &, const QPalette &, _Bool)
func (this *QItemDelegate) Selected(pixmap qtgui.QPixmap_ITF, palette qtgui.QPalette_ITF, enabled bool) *qtgui.QPixmap /*777 QPixmap **/ {
	var convArg0 = pixmap.QPixmap_PTR().GetCthis()
	var convArg1 = palette.QPalette_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate8selectedERK7QPixmapRK8QPaletteb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, enabled)
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:118
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect doCheck(const QStyleOptionViewItem &, const QRect &, const QVariant &)
func (this *QItemDelegate) DoCheck(option QStyleOptionViewItem_ITF, bounding qtcore.QRect_ITF, variant qtcore.QVariant_ITF) *qtcore.QRect /*123*/ {
	var convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	var convArg1 = bounding.QRect_PTR().GetCthis()
	var convArg2 = variant.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate7doCheckERK20QStyleOptionViewItemRK5QRectRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:120
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect textRectangle(QPainter *, const QRect &, const QFont &, const QString &)
func (this *QItemDelegate) TextRectangle(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, font qtgui.QFont_ITF, text string) *qtcore.QRect /*123*/ {
	var convArg0 = painter.QPainter_PTR().GetCthis()
	var convArg1 = rect.QRect_PTR().GetCthis()
	var convArg2 = font.QFont_PTR().GetCthis()
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate13textRectangleEP8QPainterRK5QRectRK5QFontRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
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
