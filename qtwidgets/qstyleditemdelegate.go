package qtwidgets

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h
// #include <qstyleditemdelegate.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void initStyleOption(QStyleOptionViewItem *, const QModelIndex &)
func (this *QStyledItemDelegate) InheritInitStyleOption(f func(option *QStyleOptionViewItem /*777 QStyleOptionViewItem **/, index *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QStyledItemDelegate) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QStyledItemDelegate) InheritEditorEvent(f func(event *qtcore.QEvent /*777 QEvent **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "editorEvent", f)
}

/*

 */
type QStyledItemDelegate struct {
	*QAbstractItemDelegate
}
type QStyledItemDelegate_ITF interface {
	QAbstractItemDelegate_ITF
	QStyledItemDelegate_PTR() *QStyledItemDelegate
}

func (ptr *QStyledItemDelegate) QStyledItemDelegate_PTR() *QStyledItemDelegate { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QStyledItemDelegate) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyledItemDelegate(QObject *)

/*
Constructs an item delegate with the given parent.
*/
func NewQStyledItemDelegate(parent qtcore.QObject_ITF /*777 QObject **/) *QStyledItemDelegate {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyledItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStyledItemDelegate")
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyledItemDelegate(QObject *)

/*
Constructs an item delegate with the given parent.
*/
func NewQStyledItemDelegate__() *QStyledItemDelegate {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyledItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStyledItemDelegate")
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStyledItemDelegate()

/*

 */
func DeleteQStyledItemDelegate(this *QStyledItemDelegate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::paint().

Renders the delegate using the given painter and style option for the item specified by index.

This function paints the item using the view's QStyle.

When reimplementing paint in a subclass. Use the initStyleOption() to set up the option in the same way as the QStyledItemDelegate.

Whenever possible, use the option while painting. Especially its rect variable to decide where to draw and its state to determine if it is enabled or selected.

After painting, you should ensure that the painter is returned to the state it was supplied in when this function was called. For example, it may be useful to call QPainter::save() before painting and QPainter::restore() afterwards.

See also QItemDelegate::paint(), QStyle::drawControl(), and QStyle::CE_ItemViewItem.
*/
func (this *QStyledItemDelegate) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint(const QStyleOptionViewItem &, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::sizeHint().

Returns the size needed by the delegate to display the item specified by index, taking into account the style information provided by option.

This function uses the view's QStyle to determine the size of the item.

See also QStyle::sizeFromContents() and QStyle::CT_ItemViewItem.
*/
func (this *QStyledItemDelegate) SizeHint(option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(QWidget *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::createEditor().

Returns the widget used to edit the item specified by index for editing. The parent widget and style option are used to control how the editor widget appears.

See also QAbstractItemDelegate::createEditor().
*/
func (this *QStyledItemDelegate) CreateEditor(parent QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setEditorData(QWidget *, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::setEditorData().

Sets the data to be displayed and edited by the editor from the data model item specified by the model index.

The default implementation stores the data in the editor widget's user property.

See also QMetaProperty::isUser().
*/
func (this *QStyledItemDelegate) SetEditorData(editor QWidget_ITF /*777 QWidget **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModelData(QWidget *, QAbstractItemModel *, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::setModelData().

Gets data from the editor widget and stores it in the specified model at the item index.

The default implementation gets the value to be stored in the data model from the editor widget's user property.

See also QMetaProperty::isUser().
*/
func (this *QStyledItemDelegate) SetModelData(editor QWidget_ITF /*777 QWidget **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg1 = model.QAbstractItemModel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometry(QWidget *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::updateEditorGeometry().

Updates the editor for the item specified by index according to the style option given.
*/
func (this *QStyledItemDelegate) UpdateEditorGeometry(editor QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QItemEditorFactory * itemEditorFactory() const

/*
Returns the editor factory used by the item delegate. If no editor factory is set, the function will return null.

See also setItemEditorFactory().
*/
func (this *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory /*777 QItemEditorFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate17itemEditorFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQItemEditorFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemEditorFactory(QItemEditorFactory *)

/*
Sets the editor factory to be used by the item delegate to be the factory specified. If no editor factory is set, the item delegate will use the default editor factory.

See also itemEditorFactory().
*/
func (this *QStyledItemDelegate) SetItemEditorFactory(factory QItemEditorFactory_ITF /*777 QItemEditorFactory **/) {
	var convArg0 unsafe.Pointer
	if factory != nil && factory.QItemEditorFactory_PTR() != nil {
		convArg0 = factory.QItemEditorFactory_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString displayText(const QVariant &, const QLocale &) const

/*
This function returns the string that the delegate will use to display the Qt::DisplayRole of the model in locale. value is the value of the Qt::DisplayRole provided by the model.

The default implementation uses the QLocale::toString to convert value into a QString.

This function is not called for empty model indices, i.e., indices for which the model returns an invalid QVariant.

See also QAbstractItemModel::data().
*/
func (this *QStyledItemDelegate) DisplayText(value qtcore.QVariant_ITF, locale qtcore.QLocale_ITF) string {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg1 = locale.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionViewItem *, const QModelIndex &) const

/*
Initialize option with the values using the index index. This method is useful for subclasses when they need a QStyleOptionViewItem, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QStyledItemDelegate) InitStyleOption(option QStyleOptionViewItem_ITF /*777 QStyleOptionViewItem **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QStyledItemDelegate15initStyleOptionEP20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().

Returns true if the given editor is a valid QWidget and the given event is handled; otherwise returns false. The following key press events are handled by default:


Tab
Backtab
Enter
Return
Esc


If the editor's type is QTextEdit or QPlainTextEdit then Enter and Return keys are not handled.

In the case of Tab, Backtab, Enter and Return key press events, the editor's data is comitted to the model and the editor is closed. If the event is a Tab key press the view will open an editor on the next item in the view. Likewise, if the event is a Backtab key press the view will open an editor on the previous item in the view.

If the event is a Esc key press event, the editor is closed without committing its data.

See also commitData() and closeEditor().
*/
func (this *QStyledItemDelegate) EventFilter(object qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegate11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstyleditemdelegate.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)

/*
Reimplemented from QAbstractItemDelegate::editorEvent().
*/
func (this *QStyledItemDelegate) EditorEvent(event qtcore.QEvent_ITF /*777 QEvent **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg1 = model.QAbstractItemModel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg2 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg3 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyledItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
		log.Println(123)
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
