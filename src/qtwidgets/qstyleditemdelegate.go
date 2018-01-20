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
func NewQStyledItemDelegateFromPointer(cthis unsafe.Pointer) *QStyledItemDelegate {
	bcthis0 := NewQAbstractItemDelegateFromPointer(cthis)
	return &QStyledItemDelegate{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QStyledItemDelegate) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:61
// index:0
// Public
// void QStyledItemDelegate(class QObject *)
func NewQStyledItemDelegate(parent unsafe.Pointer) *QStyledItemDelegate {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegateC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyledItemDelegateFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:62
// index:0
// Public virtual
// void ~QStyledItemDelegate()
func DeleteQStyledItemDelegate(*QStyledItemDelegate) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegateD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:65
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) Paint(painter unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:67
// index:0
// Public virtual
// QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) SizeHint(option *QStyleOptionViewItem, index *qtcore.QModelIndex) interface{} {
	var convArg0 = option.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:71
// index:0
// Public virtual
// QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) CreateEditor(parent unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) interface{} {
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), parent, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:75
// index:0
// Public virtual
// void setEditorData(class QWidget *, const class QModelIndex &)
func (this *QStyledItemDelegate) SetEditorData(editor unsafe.Pointer, index *qtcore.QModelIndex) {
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:76
// index:0
// Public virtual
// void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
func (this *QStyledItemDelegate) SetModelData(editor unsafe.Pointer, model unsafe.Pointer, index *qtcore.QModelIndex) {
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, model, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:80
// index:0
// Public virtual
// void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) UpdateEditorGeometry(editor unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:85
// index:0
// Public
// QItemEditorFactory * itemEditorFactory()
func (this *QStyledItemDelegate) ItemEditorFactory() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate17itemEditorFactoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:86
// index:0
// Public
// void setItemEditorFactory(class QItemEditorFactory *)
func (this *QStyledItemDelegate) SetItemEditorFactory(factory unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), factory)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:88
// index:0
// Public virtual
// QString displayText(const class QVariant &, const class QLocale &)
func (this *QStyledItemDelegate) DisplayText(value *qtcore.QVariant, locale *qtcore.QLocale) interface{} {
	var convArg0 = value.GetCthis()
	var convArg1 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:91
// index:0
// Protected virtual
// void initStyleOption(class QStyleOptionViewItem *, const class QModelIndex &)
func (this *QStyledItemDelegate) InitStyleOption(option unsafe.Pointer, index *qtcore.QModelIndex) {
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate15initStyleOptionEP20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:94
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QStyledItemDelegate) EventFilter(object unsafe.Pointer, event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegate11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object, event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:95
// index:0
// Protected virtual
// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) EditorEvent(event unsafe.Pointer, model unsafe.Pointer, option *QStyleOptionViewItem, index *qtcore.QModelIndex) interface{} {
	var convArg2 = option.GetCthis()
	var convArg3 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event, model, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
