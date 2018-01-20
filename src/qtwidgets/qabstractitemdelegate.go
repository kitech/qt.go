//  header block begin
// /usr/include/qt/QtWidgets/qabstractitemdelegate.h
// #include <qabstractitemdelegate.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QAbstractItemDelegate struct {
	*qtcore.QObject
}

func (this *QAbstractItemDelegate) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQAbstractItemDelegateFromPointer(cthis unsafe.Pointer) *QAbstractItemDelegate {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractItemDelegate{bcthis0}
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:60
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractItemDelegate) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:72
// index:0
// Public
// void QAbstractItemDelegate(class QObject *)
func NewQAbstractItemDelegate(parent unsafe.Pointer) *QAbstractItemDelegate {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegateC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractItemDelegateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:73
// index:0
// Public virtual
// void ~QAbstractItemDelegate()
func DeleteQAbstractItemDelegate(*QAbstractItemDelegate) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:76
// index:0
// Public pure virtual
// void paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) Paint(painter unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:80
// index:0
// Public pure virtual
// QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) SizeHint(option *QStyleOptionViewItem, index *qtcore.QModelIndex) interface{} {
	var convArg0 = option.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:84
// index:0
// Public virtual
// QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) CreateEditor(parent unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) interface{} {
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), parent, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:88
// index:0
// Public virtual
// void destroyEditor(class QWidget *, const class QModelIndex &)
func (this *QAbstractItemDelegate) DestroyEditor(editor unsafe.Pointer, index *qtcore.QModelIndex) {
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13destroyEditorEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:90
// index:0
// Public virtual
// void setEditorData(class QWidget *, const class QModelIndex &)
func (this *QAbstractItemDelegate) SetEditorData(editor unsafe.Pointer, index *qtcore.QModelIndex) {
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:92
// index:0
// Public virtual
// void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
func (this *QAbstractItemDelegate) SetModelData(editor unsafe.Pointer, model unsafe.Pointer, index *qtcore.QModelIndex) {
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, model, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:96
// index:0
// Public virtual
// void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) UpdateEditorGeometry(editor unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:101
// index:0
// Public virtual
// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) EditorEvent(event unsafe.Pointer, model unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) interface{} {
	var convArg2 = option.GetCthis()
	var convArg3 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event, model, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:106
// index:0
// Public static
// QString elidedText(const class QFontMetrics &, int, Qt::TextElideMode, const class QString &)
func (this *QAbstractItemDelegate) ElidedText(fontMetrics *qtgui.QFontMetrics, width int, mode int, text *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate10elidedTextERK12QFontMetricsiN2Qt13TextElideModeERK7QString", ffiqt.FFI_TYPE_POINTER, fontMetrics, width, mode, text)
	gopp.ErrPrint(err, rv)
	return rv
}
func QAbstractItemDelegate_ElidedText(fontMetrics *qtgui.QFontMetrics, width int, mode int, text *qtcore.QString) {
	var nilthis *QAbstractItemDelegate
	nilthis.ElidedText(fontMetrics, width, mode, text)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:109
// index:0
// Public virtual
// bool helpEvent(class QHelpEvent *, class QAbstractItemView *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) HelpEvent(event unsafe.Pointer, view unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) interface{} {
	var convArg2 = option.GetCthis()
	var convArg3 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate9helpEventEP10QHelpEventP17QAbstractItemViewRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event, view, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:114
// index:0
// Public virtual
// QVector<int> paintingRoles()
func (this *QAbstractItemDelegate) PaintingRoles() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13paintingRolesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:117
// index:0
// Public
// void commitData(class QWidget *)
func (this *QAbstractItemDelegate) CommitData(editor unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate10commitDataEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:118
// index:0
// Public
// void closeEditor(class QWidget *, class QAbstractItemDelegate::EndEditHint)
func (this *QAbstractItemDelegate) CloseEditor(editor unsafe.Pointer, hint int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11closeEditorEP7QWidgetNS_11EndEditHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:119
// index:0
// Public
// void sizeHintChanged(const class QModelIndex &)
func (this *QAbstractItemDelegate) SizeHintChanged(arg0 *qtcore.QModelIndex) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate15sizeHintChangedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
