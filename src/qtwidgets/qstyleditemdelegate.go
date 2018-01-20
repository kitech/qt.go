//  header block begin
// /usr/include/qt/QtWidgets/qstyleditemdelegate.h
// #include <qstyleditemdelegate.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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
type QStyledItemDelegate struct {
	*QAbstractItemDelegate
}

func (this *QStyledItemDelegate) GetCthis() unsafe.Pointer {
	return this.QAbstractItemDelegate.GetCthis()
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QStyledItemDelegate) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:61
// index:0
// void QStyledItemDelegate(class QObject *)
func NewQStyledItemDelegate(parent unsafe.Pointer) *QStyledItemDelegate {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegateC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyledItemDelegateFromPointer(cthis)
	return gothis
}
func NewQStyledItemDelegateFromPointer(cthis unsafe.Pointer) *QStyledItemDelegate {
	bcthis0 := NewQAbstractItemDelegateFromPointer(cthis)
	return &QStyledItemDelegate{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:62
// index:0
// virtual
// void ~QStyledItemDelegate()
func DeleteQStyledItemDelegate(*QStyledItemDelegate) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:65
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) Paint(painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, painter QPainter *, option const QStyleOptionViewItem &, index const QModelIndex &), (painter, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:67
// index:0
// virtual
// QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) SizeHint(option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, option const QStyleOptionViewItem &, index const QModelIndex &), (option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:71
// index:0
// virtual
// QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) CreateEditor(parent unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, parent QWidget *, option const QStyleOptionViewItem &, index const QModelIndex &), (parent, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:75
// index:0
// virtual
// void setEditorData(class QWidget *, const class QModelIndex &)
func (this *QStyledItemDelegate) SetEditorData(editor unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, editor QWidget *, index const QModelIndex &), (editor, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:76
// index:0
// virtual
// void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
func (this *QStyledItemDelegate) SetModelData(editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, editor QWidget *, model QAbstractItemModel *, index const QModelIndex &), (editor, model, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, model, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:80
// index:0
// virtual
// void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) UpdateEditorGeometry(editor unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, editor QWidget *, option const QStyleOptionViewItem &, index const QModelIndex &), (editor, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:85
// index:0
// QItemEditorFactory * itemEditorFactory()
func (this *QStyledItemDelegate) ItemEditorFactory() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate17itemEditorFactoryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:86
// index:0
// void setItemEditorFactory(class QItemEditorFactory *)
func (this *QStyledItemDelegate) SetItemEditorFactory(factory unsafe.Pointer) {
	// 0: (, factory QItemEditorFactory *), (factory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", ffiqt.FFI_TYPE_VOID, this.GetCthis(), factory)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:88
// index:0
// virtual
// QString displayText(const class QVariant &, const class QLocale &)
func (this *QStyledItemDelegate) DisplayText(value unsafe.Pointer, locale unsafe.Pointer) {
	// 0: (, value const QVariant &, locale const QLocale &), (value, locale)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value, locale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:91
// index:0
// virtual
// void initStyleOption(class QStyleOptionViewItem *, const class QModelIndex &)
func (this *QStyledItemDelegate) InitStyleOption(option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, option QStyleOptionViewItem *, index const QModelIndex &), (option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate15initStyleOptionEP20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:94
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QStyledItemDelegate) EventFilter(object unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, object QObject *, event QEvent *), (object, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegate11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), object, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:95
// index:0
// virtual
// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) EditorEvent(event unsafe.Pointer, model unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, event QEvent *, model QAbstractItemModel *, option const QStyleOptionViewItem &, index const QModelIndex &), (event, model, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event, model, option, index)
	gopp.ErrPrint(err, rv)
}

//  body block end
