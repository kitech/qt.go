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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QAbstractItemDelegate struct {
	*qtcore.QObject
}
type QAbstractItemDelegate_ITF interface {
	qtcore.QObject_ITF
	QAbstractItemDelegate_PTR() *QAbstractItemDelegate
}

func (ptr *QAbstractItemDelegate) QAbstractItemDelegate_PTR() *QAbstractItemDelegate { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractItemDelegate) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemDelegate(QObject *)

/*
Creates a new abstract item delegate with the given parent.
*/
func NewQAbstractItemDelegate(parent qtcore.QObject_ITF /*777 QObject **/) *QAbstractItemDelegate {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractItemDelegate")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemDelegate(QObject *)

/*
Creates a new abstract item delegate with the given parent.
*/
func NewQAbstractItemDelegate__() *QAbstractItemDelegate {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractItemDelegate")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractItemDelegate()

/*

 */
func DeleteQAbstractItemDelegate(this *QAbstractItemDelegate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
This pure abstract function must be reimplemented if you want to provide custom rendering. Use the painter and style option to render the item specified by the item index.

If you reimplement this you must also reimplement sizeHint().
*/
func (this *QAbstractItemDelegate) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize sizeHint(const QStyleOptionViewItem &, const QModelIndex &) const

/*
This pure abstract function must be reimplemented if you want to provide custom rendering. The options are specified by option and the model item by index.

If you reimplement this you must also reimplement paint().
*/
func (this *QAbstractItemDelegate) SizeHint(option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(QWidget *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
Returns the editor to be used for editing the data item with the given index. Note that the index contains information about the model being used. The editor's parent widget is specified by parent, and the item options by option.

The base implementation returns 0. If you want custom editing you will need to reimplement this function.

The returned editor widget should have Qt::StrongFocus; otherwise, QMouseEvents received by the widget will propagate to the view. The view's background will shine through unless the editor paints its own background (e.g., with setAutoFillBackground()).

See also destroyEditor(), setModelData(), and setEditorData().
*/
func (this *QAbstractItemDelegate) CreateEditor(parent QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *QWidget /*777 QWidget **/ {
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void destroyEditor(QWidget *, const QModelIndex &) const

/*
Called when the editor is no longer needed for editing the data item with the given index and should be destroyed. The default behavior is a call to deleteLater on the editor. It is possible e.g. to avoid this delete by reimplementing this function.

This function was introduced in  Qt 5.0.

See also createEditor().
*/
func (this *QAbstractItemDelegate) DestroyEditor(editor QWidget_ITF /*777 QWidget **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13destroyEditorEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setEditorData(QWidget *, const QModelIndex &) const

/*
Sets the contents of the given editor to the data for the item at the given index. Note that the index contains information about the model being used.

The base implementation does nothing. If you want custom editing you will need to reimplement this function.

See also setModelData().
*/
func (this *QAbstractItemDelegate) SetEditorData(editor QWidget_ITF /*777 QWidget **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModelData(QWidget *, QAbstractItemModel *, const QModelIndex &) const

/*
Sets the data for the item at the given index in the model to the contents of the given editor.

The base implementation does nothing. If you want custom editing you will need to reimplement this function.

See also setEditorData().
*/
func (this *QAbstractItemDelegate) SetModelData(editor QWidget_ITF /*777 QWidget **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, index qtcore.QModelIndex_ITF) {
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometry(QWidget *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
Updates the geometry of the editor for the item with the given index, according to the rectangle specified in the option. If the item has an internal layout, the editor will be laid out accordingly. Note that the index contains information about the model being used.

The base implementation does nothing. If you want custom editing you must reimplement this function.
*/
func (this *QAbstractItemDelegate) UpdateEditorGeometry(editor QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)

/*
When editing of an item starts, this function is called with the event that triggered the editing, the model, the index of the item, and the option used for rendering the item.

Mouse events are sent to editorEvent() even if they don't start editing of the item. This can, for instance, be useful if you wish to open a context menu when the right mouse button is pressed on an item.

The base implementation returns false (indicating that it has not handled the event).
*/
func (this *QAbstractItemDelegate) EditorEvent(event qtcore.QEvent_ITF /*777 QEvent **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) bool {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:106
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString elidedText(const QFontMetrics &, int, Qt::TextElideMode, const QString &)

/*

 */
func (this *QAbstractItemDelegate) ElidedText(fontMetrics qtgui.QFontMetrics_ITF, width int, mode int, text string) string {
	var convArg0 unsafe.Pointer
	if fontMetrics != nil && fontMetrics.QFontMetrics_PTR() != nil {
		convArg0 = fontMetrics.QFontMetrics_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate10elidedTextERK12QFontMetricsiN2Qt13TextElideModeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0, width, mode, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QAbstractItemDelegate_ElidedText(fontMetrics qtgui.QFontMetrics_ITF, width int, mode int, text string) string {
	var nilthis *QAbstractItemDelegate
	rv := nilthis.ElidedText(fontMetrics, width, mode, text)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:109
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool helpEvent(QHelpEvent *, QAbstractItemView *, const QStyleOptionViewItem &, const QModelIndex &)

/*
Whenever a help event occurs, this function is called with the event view option and the index that corresponds to the item where the event occurs.

Returns true if the delegate can handle the event; otherwise returns false. A return value of true indicates that the data obtained using the index had the required role.

For QEvent::ToolTip and QEvent::WhatsThis events that were handled successfully, the relevant popup may be shown depending on the user's system configuration.

This function was introduced in  Qt 4.3.

See also QHelpEvent.
*/
func (this *QAbstractItemDelegate) HelpEvent(event qtgui.QHelpEvent_ITF /*777 QHelpEvent **/, view QAbstractItemView_ITF /*777 QAbstractItemView **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QHelpEvent_PTR() != nil {
		convArg0 = event.QHelpEvent_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if view != nil && view.QAbstractItemView_PTR() != nil {
		convArg1 = view.QAbstractItemView_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg2 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg3 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate9helpEventEP10QHelpEventP17QAbstractItemViewRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void commitData(QWidget *)

/*
This signal must be emitted when the editor widget has completed editing the data, and wants to write it back into the model.
*/
func (this *QAbstractItemDelegate) CommitData(editor QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate10commitDataEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeEditor(QWidget *, QAbstractItemDelegate::EndEditHint)

/*
This signal is emitted when the user has finished editing an item using the specified editor.

The hint provides a way for the delegate to influence how the model and view behave after editing is completed. It indicates to these components what action should be performed next to provide a comfortable editing experience for the user. For example, if EditNextItem is specified, the view should use a delegate to open an editor on the next item in the model.

See also EndEditHint.
*/
func (this *QAbstractItemDelegate) CloseEditor(editor QWidget_ITF /*777 QWidget **/, hint int) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11closeEditorEP7QWidgetNS_11EndEditHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeEditor(QWidget *, QAbstractItemDelegate::EndEditHint)

/*
This signal is emitted when the user has finished editing an item using the specified editor.

The hint provides a way for the delegate to influence how the model and view behave after editing is completed. It indicates to these components what action should be performed next to provide a comfortable editing experience for the user. For example, if EditNextItem is specified, the view should use a delegate to open an editor on the next item in the model.

See also EndEditHint.
*/
func (this *QAbstractItemDelegate) CloseEditor__(editor QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	// arg: 1, QAbstractItemDelegate::EndEditHint=Elaborated, QAbstractItemDelegate::EndEditHint=Enum, , Invalid
	hint := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11closeEditorEP7QWidgetNS_11EndEditHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sizeHintChanged(const QModelIndex &)

/*
This signal must be emitted when the sizeHint() of index changed.

Views automatically connect to this signal and relayout items as necessary.

This function was introduced in  Qt 4.4.
*/
func (this *QAbstractItemDelegate) SizeHintChanged(arg0 qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QModelIndex_PTR() != nil {
		convArg0 = arg0.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate15sizeHintChangedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the different hints that the delegate can give to the model and view components to make editing data in a model a comfortable experience for the user.



These hints let the delegate influence the behavior of the view:



Note that custom views may interpret the concepts of next and previous differently.

The following hints are most useful when models are used that cache data, such as those that manipulate data locally in order to increase performance or conserve network bandwidth.



Although models and views should respond to these hints in appropriate ways, custom components may ignore any or all of them if they are not relevant.

*/
type QAbstractItemDelegate__EndEditHint = int

// There is no recommended action to be performed.
const QAbstractItemDelegate__NoHint QAbstractItemDelegate__EndEditHint = 0

// The view should use the delegate to open an editor on the next item in the view.
const QAbstractItemDelegate__EditNextItem QAbstractItemDelegate__EndEditHint = 1

// The view should use the delegate to open an editor on the previous item in the view.
const QAbstractItemDelegate__EditPreviousItem QAbstractItemDelegate__EndEditHint = 2

// If the model caches data, it should write out cached data to the underlying data store.
const QAbstractItemDelegate__SubmitModelCache QAbstractItemDelegate__EndEditHint = 3

// If the model caches data, it should discard cached data and replace it with data from the underlying data store.
const QAbstractItemDelegate__RevertModelCache QAbstractItemDelegate__EndEditHint = 4

func (this *QAbstractItemDelegate) EndEditHintItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemDelegate_EndEditHintItemName(val int) string {
	var nilthis *QAbstractItemDelegate
	return nilthis.EndEditHintItemName(val)
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
