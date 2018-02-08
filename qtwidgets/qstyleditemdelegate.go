package qtwidgets

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h
// #include <qstyleditemdelegate.h>
// #include <QtWidgets>

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
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// void initStyleOption(class QStyleOptionViewItem *, const class QModelIndex &)
func (this *QStyledItemDelegate) InheritInitStyleOption(f func(option *QStyleOptionViewItem /*777 QStyleOptionViewItem **/, index *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QStyledItemDelegate) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QStyledItemDelegate) InheritEditorEvent(f func(event *qtcore.QEvent /*777 QEvent **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "editorEvent", f)
}

type QStyledItemDelegate struct {
	*QAbstractItemDelegate
}

func (this *QStyledItemDelegate) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemDelegate.GetCthis()
	}
}
func (this *QStyledItemDelegate) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemDelegate = NewQAbstractItemDelegateFromPointer(cthis)
}
func NewQStyledItemDelegateFromPointer(cthis unsafe.Pointer) *QStyledItemDelegate {
	bcthis0 := NewQAbstractItemDelegateFromPointer(cthis)
	return &QStyledItemDelegate{bcthis0}
}
func (*QStyledItemDelegate) NewFromPointer(cthis unsafe.Pointer) *QStyledItemDelegate {
	return NewQStyledItemDelegateFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QStyledItemDelegate) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyledItemDelegate(QObject *)
func NewQStyledItemDelegate(parent *qtcore.QObject /*777 QObject **/) *QStyledItemDelegate {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyledItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStyledItemDelegate()
func DeleteQStyledItemDelegate(this *QStyledItemDelegate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QStyledItemDelegate) Paint(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint(const QStyleOptionViewItem &, const QModelIndex &)
func (this *QStyledItemDelegate) SizeHint(option *QStyleOptionViewItem, index *qtcore.QModelIndex) *qtcore.QSize /*123*/ {
	var convArg0 = option.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(QWidget *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QStyledItemDelegate) CreateEditor(parent *QWidget /*777 QWidget **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) *QWidget /*777 QWidget **/ {
	var convArg0 = parent.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setEditorData(QWidget *, const QModelIndex &)
func (this *QStyledItemDelegate) SetEditorData(editor *QWidget /*777 QWidget **/, index *qtcore.QModelIndex) {
	var convArg0 = editor.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModelData(QWidget *, QAbstractItemModel *, const QModelIndex &)
func (this *QStyledItemDelegate) SetModelData(editor *QWidget /*777 QWidget **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, index *qtcore.QModelIndex) {
	var convArg0 = editor.GetCthis()
	var convArg1 = model.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometry(QWidget *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QStyledItemDelegate) UpdateEditorGeometry(editor *QWidget /*777 QWidget **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg0 = editor.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QItemEditorFactory * itemEditorFactory()
func (this *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory /*777 QItemEditorFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate17itemEditorFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQItemEditorFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemEditorFactory(QItemEditorFactory *)
func (this *QStyledItemDelegate) SetItemEditorFactory(factory *QItemEditorFactory /*777 QItemEditorFactory **/) {
	var convArg0 = factory.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString displayText(const QVariant &, const QLocale &)
func (this *QStyledItemDelegate) DisplayText(value *qtcore.QVariant, locale *qtcore.QLocale) string {
	var convArg0 = value.GetCthis()
	var convArg1 = locale.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionViewItem *, const QModelIndex &)
func (this *QStyledItemDelegate) InitStyleOption(option *QStyleOptionViewItem /*777 QStyleOptionViewItem **/, index *qtcore.QModelIndex) {
	var convArg0 = option.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate15initStyleOptionEP20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QStyledItemDelegate) EventFilter(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = object.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegate11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QStyledItemDelegate) EditorEvent(event *qtcore.QEvent /*777 QEvent **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) bool {
	var convArg0 = event.GetCthis()
	var convArg1 = model.GetCthis()
	var convArg2 = option.GetCthis()
	var convArg3 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

//  body block end
