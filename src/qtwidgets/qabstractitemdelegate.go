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

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:60
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractItemDelegate) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:72
// index:0
// void QAbstractItemDelegate(class QObject *)
func NewQAbstractItemDelegate(parent unsafe.Pointer) *QAbstractItemDelegate {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegateC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractItemDelegateFromPointer(cthis)
	return gothis
}
func NewQAbstractItemDelegateFromPointer(cthis unsafe.Pointer) *QAbstractItemDelegate {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractItemDelegate{bcthis0}
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:122
// index:1
// void QAbstractItemDelegate(class QObjectPrivate &, class QObject *)
func NewQAbstractItemDelegate_1(arg0 unsafe.Pointer, parent unsafe.Pointer) *QAbstractItemDelegate {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegateC1ER14QObjectPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractItemDelegateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:73
// index:0
// virtual
// void ~QAbstractItemDelegate()
func DeleteQAbstractItemDelegate(*QAbstractItemDelegate) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:76
// index:0
// pure virtual
// void paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) Paint(painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, painter QPainter *, option const QStyleOptionViewItem &, index const QModelIndex &), (painter, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:80
// index:0
// pure virtual
// QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) SizeHint(option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, option const QStyleOptionViewItem &, index const QModelIndex &), (option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:84
// index:0
// virtual
// QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) CreateEditor(parent unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, parent QWidget *, option const QStyleOptionViewItem &, index const QModelIndex &), (parent, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:88
// index:0
// virtual
// void destroyEditor(class QWidget *, const class QModelIndex &)
func (this *QAbstractItemDelegate) DestroyEditor(editor unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, editor QWidget *, index const QModelIndex &), (editor, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13destroyEditorEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:90
// index:0
// virtual
// void setEditorData(class QWidget *, const class QModelIndex &)
func (this *QAbstractItemDelegate) SetEditorData(editor unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, editor QWidget *, index const QModelIndex &), (editor, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:92
// index:0
// virtual
// void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
func (this *QAbstractItemDelegate) SetModelData(editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, editor QWidget *, model QAbstractItemModel *, index const QModelIndex &), (editor, model, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, model, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:96
// index:0
// virtual
// void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) UpdateEditorGeometry(editor unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, editor QWidget *, option const QStyleOptionViewItem &, index const QModelIndex &), (editor, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:101
// index:0
// virtual
// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) EditorEvent(event unsafe.Pointer, model unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, event QEvent *, model QAbstractItemModel *, option const QStyleOptionViewItem &, index const QModelIndex &), (event, model, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event, model, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:106
// index:0
// static
// QString elidedText(const class QFontMetrics &, int, Qt::TextElideMode, const class QString &)
func (this *QAbstractItemDelegate) ElidedText(fontMetrics unsafe.Pointer, width int, mode int, text unsafe.Pointer) {
	// 0: (fontMetrics const QFontMetrics &, width int, mode Qt::TextElideMode, text const QString &), (fontMetrics, width, mode, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate10elidedTextERK12QFontMetricsiN2Qt13TextElideModeERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAbstractItemDelegate_ElidedText(fontMetrics unsafe.Pointer, width int, mode int, text unsafe.Pointer) {
	// 0: (fontMetrics const QFontMetrics &, width int, mode Qt::TextElideMode, text const QString &), (fontMetrics, width, mode, text)
	var nilthis *QAbstractItemDelegate
	nilthis.ElidedText(fontMetrics, width, mode, text)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:109
// index:0
// virtual
// bool helpEvent(class QHelpEvent *, class QAbstractItemView *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QAbstractItemDelegate) HelpEvent(event unsafe.Pointer, view unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, event QHelpEvent *, view QAbstractItemView *, option const QStyleOptionViewItem &, index const QModelIndex &), (event, view, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate9helpEventEP10QHelpEventP17QAbstractItemViewRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event, view, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:114
// index:0
// virtual
// QVector<int> paintingRoles()
func (this *QAbstractItemDelegate) PaintingRoles() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13paintingRolesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:117
// index:0
// void commitData(class QWidget *)
func (this *QAbstractItemDelegate) CommitData(editor unsafe.Pointer) {
	// 0: (, editor QWidget *), (editor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate10commitDataEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:118
// index:0
// void closeEditor(class QWidget *, class QAbstractItemDelegate::EndEditHint)
func (this *QAbstractItemDelegate) CloseEditor(editor unsafe.Pointer, hint int) {
	// 0: (, editor QWidget *, hint QAbstractItemDelegate::EndEditHint), (editor, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11closeEditorEP7QWidgetNS_11EndEditHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:119
// index:0
// void sizeHintChanged(const class QModelIndex &)
func (this *QAbstractItemDelegate) SizeHintChanged(arg0 unsafe.Pointer) {
	// 0: (, const QModelIndex & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QAbstractItemDelegate15sizeHintChangedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
