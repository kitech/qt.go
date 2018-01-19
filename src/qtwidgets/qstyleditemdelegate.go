//  header block begin
// /usr/include/qt/QtWidgets/qstyleditemdelegate.h
// #include <qstyleditemdelegate.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QStyledItemDelegate) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:61
// index:0
// void QStyledItemDelegate(class QObject *)
func NewQStyledItemDelegate(parent unsafe.Pointer) *QStyledItemDelegate {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegateC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QStyledItemDelegate{cthis}
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
	// 0: (, QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index), (painter, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:67
// index:0
// virtual
// QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) SizeHint(option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, const QStyleOptionViewItem & option, const QModelIndex & index), (option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:71
// index:0
// virtual
// QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) CreateEditor(parent unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index), (parent, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:75
// index:0
// virtual
// void setEditorData(class QWidget *, const class QModelIndex &)
func (this *QStyledItemDelegate) SetEditorData(editor unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, QWidget * editor, const QModelIndex & index), (editor, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, editor, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:76
// index:0
// virtual
// void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
func (this *QStyledItemDelegate) SetModelData(editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, QWidget * editor, QAbstractItemModel * model, const QModelIndex & index), (editor, model, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, editor, model, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:80
// index:0
// virtual
// void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) UpdateEditorGeometry(editor unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index), (editor, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, editor, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:85
// index:0
// QItemEditorFactory * itemEditorFactory()
func (this *QStyledItemDelegate) ItemEditorFactory() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate17itemEditorFactoryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:86
// index:0
// void setItemEditorFactory(class QItemEditorFactory *)
func (this *QStyledItemDelegate) SetItemEditorFactory(factory unsafe.Pointer) {
	// 0: (, QItemEditorFactory * factory), (factory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", ffiqt.FFI_TYPE_VOID, this.cthis, factory)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:88
// index:0
// virtual
// QString displayText(const class QVariant &, const class QLocale &)
func (this *QStyledItemDelegate) DisplayText(value unsafe.Pointer, locale unsafe.Pointer) {
	// 0: (, const QVariant & value, const QLocale & locale), (value, locale)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale", ffiqt.FFI_TYPE_VOID, this.cthis, value, locale)
	gopp.ErrPrint(err, rv)
}

//  body block end
