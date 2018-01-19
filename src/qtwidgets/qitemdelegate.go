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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QItemDelegate) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:62
// index:0
// void QItemDelegate(class QObject *)
func NewQItemDelegate(parent unsafe.Pointer) *QItemDelegate {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegateC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QItemDelegate{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate11hasClippingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:66
// index:0
// void setClipping(_Bool)
func (this *QItemDelegate) SetClipping(clip bool) {
	// 0: (, bool clip), (&clip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate11setClippingEb", ffiqt.FFI_TYPE_VOID, this.cthis, &clip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:69
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) Paint(painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index), (painter, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:72
// index:0
// virtual
// QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) SizeHint(option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, const QStyleOptionViewItem & option, const QModelIndex & index), (option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:76
// index:0
// virtual
// QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) CreateEditor(parent unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index), (parent, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:80
// index:0
// virtual
// void setEditorData(class QWidget *, const class QModelIndex &)
func (this *QItemDelegate) SetEditorData(editor unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, QWidget * editor, const QModelIndex & index), (editor, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, editor, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:81
// index:0
// virtual
// void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
func (this *QItemDelegate) SetModelData(editor unsafe.Pointer, model unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, QWidget * editor, QAbstractItemModel * model, const QModelIndex & index), (editor, model, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, editor, model, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:83
// index:0
// virtual
// void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QItemDelegate) UpdateEditorGeometry(editor unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index), (editor, option, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, editor, option, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:88
// index:0
// QItemEditorFactory * itemEditorFactory()
func (this *QItemDelegate) ItemEditorFactory() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QItemDelegate17itemEditorFactoryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:89
// index:0
// void setItemEditorFactory(class QItemEditorFactory *)
func (this *QItemDelegate) SetItemEditorFactory(factory unsafe.Pointer) {
	// 0: (, QItemEditorFactory * factory), (factory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", ffiqt.FFI_TYPE_VOID, this.cthis, factory)
	gopp.ErrPrint(err, rv)
}

//  body block end
