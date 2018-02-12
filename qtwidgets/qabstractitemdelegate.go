package qtwidgets

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h
// #include <qabstractitemdelegate.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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

type QAbstractItemDelegate struct {
	*qtcore.QObject
}

func (this *QAbstractItemDelegate) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractItemDelegate) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAbstractItemDelegateFromPointer(cthis unsafe.Pointer) *QAbstractItemDelegate {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractItemDelegate{bcthis0}
}
func (*QAbstractItemDelegate) NewFromPointer(cthis unsafe.Pointer) *QAbstractItemDelegate {
	return NewQAbstractItemDelegateFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAbstractItemDelegate) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemDelegate(QObject *)
func NewQAbstractItemDelegate(parent *qtcore.QObject /*777 QObject **/) *QAbstractItemDelegate {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractItemDelegate()
func DeleteQAbstractItemDelegate(this *QAbstractItemDelegate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QAbstractItemDelegate) Paint(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize sizeHint(const QStyleOptionViewItem &, const QModelIndex &)
func (this *QAbstractItemDelegate) SizeHint(option *QStyleOptionViewItem, index *qtcore.QModelIndex) *qtcore.QSize /*123*/ {
	var convArg0 = option.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(QWidget *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QAbstractItemDelegate) CreateEditor(parent *QWidget /*777 QWidget **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) *QWidget /*777 QWidget **/ {
	var convArg0 = parent.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void destroyEditor(QWidget *, const QModelIndex &)
func (this *QAbstractItemDelegate) DestroyEditor(editor *QWidget /*777 QWidget **/, index *qtcore.QModelIndex) {
	var convArg0 = editor.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13destroyEditorEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setEditorData(QWidget *, const QModelIndex &)
func (this *QAbstractItemDelegate) SetEditorData(editor *QWidget /*777 QWidget **/, index *qtcore.QModelIndex) {
	var convArg0 = editor.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModelData(QWidget *, QAbstractItemModel *, const QModelIndex &)
func (this *QAbstractItemDelegate) SetModelData(editor *QWidget /*777 QWidget **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, index *qtcore.QModelIndex) {
	var convArg0 = editor.GetCthis()
	var convArg1 = model.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometry(QWidget *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QAbstractItemDelegate) UpdateEditorGeometry(editor *QWidget /*777 QWidget **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg0 = editor.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QAbstractItemDelegate) EditorEvent(event *qtcore.QEvent /*777 QEvent **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) bool {
	var convArg0 = event.GetCthis()
	var convArg1 = model.GetCthis()
	var convArg2 = option.GetCthis()
	var convArg3 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:106
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString elidedText(const QFontMetrics &, int, Qt::TextElideMode, const QString &)
func (this *QAbstractItemDelegate) ElidedText(fontMetrics *qtgui.QFontMetrics, width int, mode int, text string) string {
	var convArg0 = fontMetrics.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate10elidedTextERK12QFontMetricsiN2Qt13TextElideModeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0, width, mode, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QAbstractItemDelegate_ElidedText(fontMetrics *qtgui.QFontMetrics, width int, mode int, text string) string {
	var nilthis *QAbstractItemDelegate
	rv := nilthis.ElidedText(fontMetrics, width, mode, text)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:109
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool helpEvent(QHelpEvent *, QAbstractItemView *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QAbstractItemDelegate) HelpEvent(event *qtgui.QHelpEvent /*777 QHelpEvent **/, view *QAbstractItemView /*777 QAbstractItemView **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) bool {
	var convArg0 = event.GetCthis()
	var convArg1 = view.GetCthis()
	var convArg2 = option.GetCthis()
	var convArg3 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate9helpEventEP10QHelpEventP17QAbstractItemViewRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void commitData(QWidget *)
func (this *QAbstractItemDelegate) CommitData(editor *QWidget /*777 QWidget **/) {
	var convArg0 = editor.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate10commitDataEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeEditor(QWidget *, QAbstractItemDelegate::EndEditHint)
func (this *QAbstractItemDelegate) CloseEditor(editor *QWidget /*777 QWidget **/, hint int) {
	var convArg0 = editor.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11closeEditorEP7QWidgetNS_11EndEditHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sizeHintChanged(const QModelIndex &)
func (this *QAbstractItemDelegate) SizeHintChanged(arg0 *qtcore.QModelIndex) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate15sizeHintChangedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QAbstractItemDelegate__EndEditHint = int

const QAbstractItemDelegate__NoHint QAbstractItemDelegate__EndEditHint = 0
const QAbstractItemDelegate__EditNextItem QAbstractItemDelegate__EndEditHint = 1
const QAbstractItemDelegate__EditPreviousItem QAbstractItemDelegate__EndEditHint = 2
const QAbstractItemDelegate__SubmitModelCache QAbstractItemDelegate__EndEditHint = 3
const QAbstractItemDelegate__RevertModelCache QAbstractItemDelegate__EndEditHint = 4

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
