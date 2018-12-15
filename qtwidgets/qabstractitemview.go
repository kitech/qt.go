// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractitemview.h
// #include <qabstractitemview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 44
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

// void rowsInserted(const QModelIndex &, int, int)
func (this *QAbstractItemView) InheritRowsInserted(f func(parent *qtcore.QModelIndex, start int, end_ int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsInserted", f)
}

// void rowsAboutToBeRemoved(const QModelIndex &, int, int)
func (this *QAbstractItemView) InheritRowsAboutToBeRemoved(f func(parent *qtcore.QModelIndex, start int, end_ int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsAboutToBeRemoved", f)
}

// void selectionChanged(const QItemSelection &, const QItemSelection &)
func (this *QAbstractItemView) InheritSelectionChanged(f func(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) /*void*/) {
	qtrt.SetAllInheritCallback(this, "selectionChanged", f)
}

// void currentChanged(const QModelIndex &, const QModelIndex &)
func (this *QAbstractItemView) InheritCurrentChanged(f func(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "currentChanged", f)
}

// void updateEditorData()
func (this *QAbstractItemView) InheritUpdateEditorData(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateEditorData", f)
}

// void updateEditorGeometries()
func (this *QAbstractItemView) InheritUpdateEditorGeometries(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateEditorGeometries", f)
}

// void updateGeometries()
func (this *QAbstractItemView) InheritUpdateGeometries(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateGeometries", f)
}

// void verticalScrollbarAction(int)
func (this *QAbstractItemView) InheritVerticalScrollbarAction(f func(action int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "verticalScrollbarAction", f)
}

// void horizontalScrollbarAction(int)
func (this *QAbstractItemView) InheritHorizontalScrollbarAction(f func(action int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "horizontalScrollbarAction", f)
}

// void verticalScrollbarValueChanged(int)
func (this *QAbstractItemView) InheritVerticalScrollbarValueChanged(f func(value int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "verticalScrollbarValueChanged", f)
}

// void horizontalScrollbarValueChanged(int)
func (this *QAbstractItemView) InheritHorizontalScrollbarValueChanged(f func(value int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "horizontalScrollbarValueChanged", f)
}

// void closeEditor(QWidget *, QAbstractItemDelegate::EndEditHint)
func (this *QAbstractItemView) InheritCloseEditor(f func(editor *QWidget /*777 QWidget **/, hint int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEditor", f)
}

// void commitData(QWidget *)
func (this *QAbstractItemView) InheritCommitData(f func(editor *QWidget /*777 QWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "commitData", f)
}

// void editorDestroyed(QObject *)
func (this *QAbstractItemView) InheritEditorDestroyed(f func(editor *qtcore.QObject /*777 QObject **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "editorDestroyed", f)
}

// void setHorizontalStepsPerItem(int)
func (this *QAbstractItemView) InheritSetHorizontalStepsPerItem(f func(steps int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setHorizontalStepsPerItem", f)
}

// int horizontalStepsPerItem()
func (this *QAbstractItemView) InheritHorizontalStepsPerItem(f func() int) {
	qtrt.SetAllInheritCallback(this, "horizontalStepsPerItem", f)
}

// void setVerticalStepsPerItem(int)
func (this *QAbstractItemView) InheritSetVerticalStepsPerItem(f func(steps int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setVerticalStepsPerItem", f)
}

// int verticalStepsPerItem()
func (this *QAbstractItemView) InheritVerticalStepsPerItem(f func() int) {
	qtrt.SetAllInheritCallback(this, "verticalStepsPerItem", f)
}

// QModelIndex moveCursor(QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QAbstractItemView) InheritMoveCursor(f func(cursorAction int, modifiers int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "moveCursor", f)
}

// int horizontalOffset()
func (this *QAbstractItemView) InheritHorizontalOffset(f func() int) {
	qtrt.SetAllInheritCallback(this, "horizontalOffset", f)
}

// int verticalOffset()
func (this *QAbstractItemView) InheritVerticalOffset(f func() int) {
	qtrt.SetAllInheritCallback(this, "verticalOffset", f)
}

// bool isIndexHidden(const QModelIndex &)
func (this *QAbstractItemView) InheritIsIndexHidden(f func(index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "isIndexHidden", f)
}

// void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)
func (this *QAbstractItemView) InheritSetSelection(f func(rect *qtcore.QRect, command int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSelection", f)
}

// QRegion visualRegionForSelection(const QItemSelection &)
func (this *QAbstractItemView) InheritVisualRegionForSelection(f func(selection *qtcore.QItemSelection) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "visualRegionForSelection", f)
}

// QModelIndexList selectedIndexes()
func (this *QAbstractItemView) InheritSelectedIndexes(f func() *qtcore.QModelIndexList /*9999*/) {
	qtrt.SetAllInheritCallback(this, "selectedIndexes", f)
}

// bool edit(const QModelIndex &, QAbstractItemView::EditTrigger, QEvent *)
func (this *QAbstractItemView) InheritEdit(f func(index *qtcore.QModelIndex, trigger int, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "edit", f)
}

// QItemSelectionModel::SelectionFlags selectionCommand(const QModelIndex &, const QEvent *)
func (this *QAbstractItemView) InheritSelectionCommand(f func(index *qtcore.QModelIndex, event *qtcore.QEvent /*777 const QEvent **/) int) {
	qtrt.SetAllInheritCallback(this, "selectionCommand", f)
}

// void startDrag(Qt::DropActions)
func (this *QAbstractItemView) InheritStartDrag(f func(supportedActions int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "startDrag", f)
}

// QStyleOptionViewItem viewOptions()
func (this *QAbstractItemView) InheritViewOptions(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewOptions", f)
}

// QAbstractItemView::State state()
func (this *QAbstractItemView) InheritState(f func() int) {
	qtrt.SetAllInheritCallback(this, "state", f)
}

// void setState(QAbstractItemView::State)
func (this *QAbstractItemView) InheritSetState(f func(state int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setState", f)
}

// void scheduleDelayedItemsLayout()
func (this *QAbstractItemView) InheritScheduleDelayedItemsLayout(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "scheduleDelayedItemsLayout", f)
}

// void executeDelayedItemsLayout()
func (this *QAbstractItemView) InheritExecuteDelayedItemsLayout(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "executeDelayedItemsLayout", f)
}

// void setDirtyRegion(const QRegion &)
func (this *QAbstractItemView) InheritSetDirtyRegion(f func(region *qtgui.QRegion) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setDirtyRegion", f)
}

// void scrollDirtyRegion(int, int)
func (this *QAbstractItemView) InheritScrollDirtyRegion(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollDirtyRegion", f)
}

// QPoint dirtyRegionOffset()
func (this *QAbstractItemView) InheritDirtyRegionOffset(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "dirtyRegionOffset", f)
}

// void startAutoScroll()
func (this *QAbstractItemView) InheritStartAutoScroll(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "startAutoScroll", f)
}

// void stopAutoScroll()
func (this *QAbstractItemView) InheritStopAutoScroll(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "stopAutoScroll", f)
}

// void doAutoScroll()
func (this *QAbstractItemView) InheritDoAutoScroll(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "doAutoScroll", f)
}

// bool focusNextPrevChild(bool)
func (this *QAbstractItemView) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// bool event(QEvent *)
func (this *QAbstractItemView) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool viewportEvent(QEvent *)
func (this *QAbstractItemView) InheritViewportEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "viewportEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QAbstractItemView) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QAbstractItemView) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QAbstractItemView) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QAbstractItemView) InheritMouseDoubleClickEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void dragEnterEvent(QDragEnterEvent *)
func (this *QAbstractItemView) InheritDragEnterEvent(f func(event *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(QDragMoveEvent *)
func (this *QAbstractItemView) InheritDragMoveEvent(f func(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QAbstractItemView) InheritDragLeaveEvent(f func(event *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(QDropEvent *)
func (this *QAbstractItemView) InheritDropEvent(f func(event *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QAbstractItemView) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QAbstractItemView) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QAbstractItemView) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QAbstractItemView) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QAbstractItemView) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QAbstractItemView) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QAbstractItemView) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// QAbstractItemView::DropIndicatorPosition dropIndicatorPosition()
func (this *QAbstractItemView) InheritDropIndicatorPosition(f func() int) {
	qtrt.SetAllInheritCallback(this, "dropIndicatorPosition", f)
}

// QSize viewportSizeHint()
func (this *QAbstractItemView) InheritViewportSizeHint(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewportSizeHint", f)
}

/*

 */
type QAbstractItemView struct {
	*QAbstractScrollArea
}
type QAbstractItemView_ITF interface {
	QAbstractScrollArea_ITF
	QAbstractItemView_PTR() *QAbstractItemView
}

func (ptr *QAbstractItemView) QAbstractItemView_PTR() *QAbstractItemView { return ptr }

func (this *QAbstractItemView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractScrollArea.GetCthis()
	}
}
func (this *QAbstractItemView) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractScrollArea = NewQAbstractScrollAreaFromPointer(cthis)
}
func NewQAbstractItemViewFromPointer(cthis unsafe.Pointer) *QAbstractItemView {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QAbstractItemView{bcthis0}
}
func (*QAbstractItemView) NewFromPointer(cthis unsafe.Pointer) *QAbstractItemView {
	return NewQAbstractItemViewFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractItemView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemView(QWidget *)

/*
Constructs an abstract item view with the given parent.
*/
func (*QAbstractItemView) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QAbstractItemView {
	return NewQAbstractItemView(parent)
}
func NewQAbstractItemView(parent QWidget_ITF /*777 QWidget **/) *QAbstractItemView {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractItemViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractItemView")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemView(QWidget *)

/*
Constructs an abstract item view with the given parent.
*/
func (*QAbstractItemView) NewForInheritp() *QAbstractItemView {
	return NewQAbstractItemViewp()
}
func NewQAbstractItemViewp() *QAbstractItemView {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractItemViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractItemView")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:128
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractItemView()

/*

 */
func DeleteQAbstractItemView(this *QAbstractItemView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:130
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)

/*
Sets the model for the view to present.

This function will create and set a new selection model, replacing any model that was previously set with setSelectionModel(). However, the old selection model will not be deleted as it may be shared between several views. We recommend that you delete the old selection model if it is no longer required. This is done with the following code:


  QItemSelectionModel *m = view->selectionModel();
  view->setModel(new model);
  delete m;



If both the old model and the old selection model do not have parents, or if their parents are long-lived objects, it may be preferable to call their deleteLater() functions to explicitly delete them.

The view does not take ownership of the model unless it is the model's parent object because the model may be shared between many different views.

See also model(), selectionModel(), and setSelectionModel().
*/
func (this *QAbstractItemView) SetModel(model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * model() const

/*
Returns the model that this view is presenting.

See also setModel().
*/
func (this *QAbstractItemView) Model() *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:133
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSelectionModel(QItemSelectionModel *)

/*
Sets the current selection model to the given selectionModel.

Note that, if you call setModel() after this function, the given selectionModel will be replaced by one created by the view.

Note: It is up to the application to delete the old selection model if it is no longer needed; i.e., if it is not being used by other views. This will happen automatically when its parent object is deleted. However, if it does not have a parent, or if the parent is a long-lived object, it may be preferable to call its deleteLater() function to explicitly delete it.

See also selectionModel(), setModel(), and clearSelection().
*/
func (this *QAbstractItemView) SetSelectionModel(selectionModel qtcore.QItemSelectionModel_ITF /*777 QItemSelectionModel **/) {
	var convArg0 unsafe.Pointer
	if selectionModel != nil && selectionModel.QItemSelectionModel_PTR() != nil {
		convArg0 = selectionModel.QItemSelectionModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView17setSelectionModelEP19QItemSelectionModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QItemSelectionModel * selectionModel() const

/*
Returns the current selection model.

See also setSelectionModel() and selectedIndexes().
*/
func (this *QAbstractItemView) SelectionModel() *qtcore.QItemSelectionModel /*777 QItemSelectionModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView14selectionModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQItemSelectionModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemDelegate(QAbstractItemDelegate *)

/*
Sets the item delegate for this view and its model to delegate. This is useful if you want complete control over the editing and display of items.

Any existing delegate will be removed, but not deleted. QAbstractItemView does not take ownership of delegate.

Warning: You should not share the same instance of a delegate between views. Doing so can cause incorrect or unintuitive editing behavior since each view connected to a given delegate may receive the closeEditor() signal, and attempt to access, modify or close an editor that has already been closed.

See also itemDelegate().
*/
func (this *QAbstractItemView) SetItemDelegate(delegate QAbstractItemDelegate_ITF /*777 QAbstractItemDelegate **/) {
	var convArg0 unsafe.Pointer
	if delegate != nil && delegate.QAbstractItemDelegate_PTR() != nil {
		convArg0 = delegate.QAbstractItemDelegate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15setItemDelegateEP21QAbstractItemDelegate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegate() const

/*
Returns the item delegate used by this view and model. This is either one set with setItemDelegate(), or the default one.

See also setItemDelegate().
*/
func (this *QAbstractItemView) ItemDelegate() *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView12itemDelegateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:226
// index:1
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegate(const QModelIndex &) const

/*
Returns the item delegate used by this view and model. This is either one set with setItemDelegate(), or the default one.

See also setItemDelegate().
*/
func (this *QAbstractItemView) ItemDelegate1(index qtcore.QModelIndex_ITF) *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView12itemDelegateERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionMode(QAbstractItemView::SelectionMode)

/*

 */
func (this *QAbstractItemView) SetSelectionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16setSelectionModeENS_13SelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::SelectionMode selectionMode() const

/*

 */
func (this *QAbstractItemView) SelectionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView13selectionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionBehavior(QAbstractItemView::SelectionBehavior)

/*

 */
func (this *QAbstractItemView) SetSelectionBehavior(behavior int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView20setSelectionBehaviorENS_17SelectionBehaviorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), behavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::SelectionBehavior selectionBehavior() const

/*

 */
func (this *QAbstractItemView) SelectionBehavior() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17selectionBehaviorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:145
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex currentIndex() const

/*
Returns the model index of the current item.

See also setCurrentIndex().
*/
func (this *QAbstractItemView) CurrentIndex() *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:146
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex rootIndex() const

/*
Returns the model index of the model's root item. The root item is the parent item to the view's toplevel items. The root can be invalid.

See also setRootIndex().
*/
func (this *QAbstractItemView) RootIndex() *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView9rootIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEditTriggers(QAbstractItemView::EditTriggers)

/*

 */
func (this *QAbstractItemView) SetEditTriggers(triggers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15setEditTriggersE6QFlagsINS_11EditTriggerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), triggers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:149
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::EditTriggers editTriggers() const

/*

 */
func (this *QAbstractItemView) EditTriggers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView12editTriggersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalScrollMode(QAbstractItemView::ScrollMode)

/*

 */
func (this *QAbstractItemView) SetVerticalScrollMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21setVerticalScrollModeENS_10ScrollModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:152
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::ScrollMode verticalScrollMode() const

/*

 */
func (this *QAbstractItemView) VerticalScrollMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView18verticalScrollModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetVerticalScrollMode()

/*

 */
func (this *QAbstractItemView) ResetVerticalScrollMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView23resetVerticalScrollModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalScrollMode(QAbstractItemView::ScrollMode)

/*

 */
func (this *QAbstractItemView) SetHorizontalScrollMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView23setHorizontalScrollModeENS_10ScrollModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:156
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::ScrollMode horizontalScrollMode() const

/*

 */
func (this *QAbstractItemView) HorizontalScrollMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView20horizontalScrollModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetHorizontalScrollMode()

/*

 */
func (this *QAbstractItemView) ResetHorizontalScrollMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView25resetHorizontalScrollModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoScroll(bool)

/*

 */
func (this *QAbstractItemView) SetAutoScroll(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13setAutoScrollEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAutoScroll() const

/*

 */
func (this *QAbstractItemView) HasAutoScroll() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView13hasAutoScrollEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoScrollMargin(int)

/*

 */
func (this *QAbstractItemView) SetAutoScrollMargin(margin int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView19setAutoScrollMarginEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:163
// index:0
// Public Visibility=Default Availability=Available
// [4] int autoScrollMargin() const

/*

 */
func (this *QAbstractItemView) AutoScrollMargin() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16autoScrollMarginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabKeyNavigation(bool)

/*

 */
func (this *QAbstractItemView) SetTabKeyNavigation(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView19setTabKeyNavigationEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabKeyNavigation() const

/*

 */
func (this *QAbstractItemView) TabKeyNavigation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16tabKeyNavigationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDropIndicatorShown(bool)

/*

 */
func (this *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21setDropIndicatorShownEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool showDropIndicator() const

/*

 */
func (this *QAbstractItemView) ShowDropIndicator() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17showDropIndicatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragEnabled(bool)

/*

 */
func (this *QAbstractItemView) SetDragEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14setDragEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dragEnabled() const

/*

 */
func (this *QAbstractItemView) DragEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView11dragEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragDropOverwriteMode(bool)

/*

 */
func (this *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView24setDragDropOverwriteModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), overwrite)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:176
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dragDropOverwriteMode() const

/*

 */
func (this *QAbstractItemView) DragDropOverwriteMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView21dragDropOverwriteModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragDropMode(QAbstractItemView::DragDropMode)

/*

 */
func (this *QAbstractItemView) SetDragDropMode(behavior int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15setDragDropModeENS_12DragDropModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), behavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:188
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::DragDropMode dragDropMode() const

/*

 */
func (this *QAbstractItemView) DragDropMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView12dragDropModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultDropAction(Qt::DropAction)

/*

 */
func (this *QAbstractItemView) SetDefaultDropAction(dropAction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView20setDefaultDropActionEN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dropAction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:191
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction defaultDropAction() const

/*

 */
func (this *QAbstractItemView) DefaultDropAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17defaultDropActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlternatingRowColors(bool)

/*

 */
func (this *QAbstractItemView) SetAlternatingRowColors(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView23setAlternatingRowColorsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:195
// index:0
// Public Visibility=Default Availability=Available
// [1] bool alternatingRowColors() const

/*

 */
func (this *QAbstractItemView) AlternatingRowColors() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView20alternatingRowColorsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)

/*

 */
func (this *QAbstractItemView) SetIconSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:198
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize() const

/*

 */
func (this *QAbstractItemView) IconSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView8iconSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextElideMode(Qt::TextElideMode)

/*

 */
func (this *QAbstractItemView) SetTextElideMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16setTextElideModeEN2Qt13TextElideModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:201
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextElideMode textElideMode() const

/*

 */
func (this *QAbstractItemView) TextElideMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView13textElideModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:203
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void keyboardSearch(const QString &)

/*
Moves to and selects the item best matching the string search. If no item is found nothing happens.

In the default implementation, the search is reset if search is empty, or the time interval since the last search has exceeded QApplication::keyboardInputInterval().
*/
func (this *QAbstractItemView) KeyboardSearch(search string) {
	var tmpArg0 = qtcore.NewQString5(search)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14keyboardSearchERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:205
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QRect visualRect(const QModelIndex &) const

/*
Returns the rectangle on the viewport occupied by the item at index.

If your item is displayed in several areas then visualRect should return the primary area that contains index and not the complete area that index might encompasses, touch or cause drawing.

In the base class this is a pure virtual function.

See also indexAt() and visualRegionForSelection().
*/
func (this *QAbstractItemView) VisualRect(index qtcore.QModelIndex_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView10visualRectERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:206
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, QAbstractItemView::ScrollHint)

/*
Scrolls the view if necessary to ensure that the item at index is visible. The view will try to position the item according to the given hint.

In the base class this is a pure virtual function.
*/
func (this *QAbstractItemView) ScrollTo(index qtcore.QModelIndex_ITF, hint int) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView8scrollToERK11QModelIndexNS_10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:206
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, QAbstractItemView::ScrollHint)

/*
Scrolls the view if necessary to ensure that the item at index is visible. The view will try to position the item according to the given hint.

In the base class this is a pure virtual function.
*/
func (this *QAbstractItemView) ScrollTop(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, QAbstractItemView::ScrollHint=Enum, QAbstractItemView::ScrollHint=Enum, , Invalid
	hint := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView8scrollToERK11QModelIndexNS_10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:207
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex indexAt(const QPoint &) const

/*
Returns the model index of the item at the viewport coordinates point.

In the base class this is a pure virtual function.

See also visualRect().
*/
func (this *QAbstractItemView) IndexAt(point qtcore.QPoint_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPoint_PTR() != nil {
		convArg0 = point.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView7indexAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:209
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizeHintForIndex(const QModelIndex &) const

/*
Returns the size hint for the item with the specified index or an invalid size for invalid indexes.

See also sizeHintForRow() and sizeHintForColumn().
*/
func (this *QAbstractItemView) SizeHintForIndex(index qtcore.QModelIndex_ITF) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16sizeHintForIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:210
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int sizeHintForRow(int) const

/*
Returns the height size hint for the specified row or -1 if there is no model.

The returned height is calculated using the size hints of the given row's items, i.e. the returned value is the maximum height among the items. Note that to control the height of a row, you must reimplement the QAbstractItemDelegate::sizeHint() function.

This function is used in views with a vertical header to find the size hint for a header section based on the contents of the given row.

See also sizeHintForColumn().
*/
func (this *QAbstractItemView) SizeHintForRow(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView14sizeHintForRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:211
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int sizeHintForColumn(int) const

/*
Returns the width size hint for the specified column or -1 if there is no model.

This function is used in views with a horizontal header to find the size hint for a header section based on the contents of the given column.

See also sizeHintForRow().
*/
func (this *QAbstractItemView) SizeHintForColumn(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17sizeHintForColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void openPersistentEditor(const QModelIndex &)

/*
Opens a persistent editor on the item at the given index. If no editor exists, the delegate will create a new editor.

See also closePersistentEditor() and isPersistentEditorOpen().
*/
func (this *QAbstractItemView) OpenPersistentEditor(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView20openPersistentEditorERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closePersistentEditor(const QModelIndex &)

/*
Closes the persistent editor for the item at the given index.

See also openPersistentEditor() and isPersistentEditorOpen().
*/
func (this *QAbstractItemView) ClosePersistentEditor(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21closePersistentEditorERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:215
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentEditorOpen(const QModelIndex &) const

/*
Returns whether a persistent editor is open for the item at index index.

This function was introduced in  Qt 5.10.

See also openPersistentEditor() and closePersistentEditor().
*/
func (this *QAbstractItemView) IsPersistentEditorOpen(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView22isPersistentEditorOpenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndexWidget(const QModelIndex &, QWidget *)

/*
Sets the given widget on the item at the given index, passing the ownership of the widget to the viewport.

If index is invalid (e.g., if you pass the root index), this function will do nothing.

The given widget's autoFillBackground property must be set to true, otherwise the widget's background will be transparent, showing both the model data and the item at the given index.

If index widget A is replaced with index widget B, index widget A will be deleted. For example, in the code snippet below, the QLineEdit object will be deleted.


  setIndexWidget(index, new QLineEdit);
  ...
  setIndexWidget(index, new QTextEdit);



This function should only be used to display static content within the visible area corresponding to an item of data. If you want to display custom dynamic content or implement a custom editor widget, subclass QItemDelegate instead.

This function was introduced in  Qt 4.1.

See also indexWidget() and Delegate Classes.
*/
func (this *QAbstractItemView) SetIndexWidget(index qtcore.QModelIndex_ITF, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14setIndexWidgetERK11QModelIndexP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:218
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * indexWidget(const QModelIndex &) const

/*
Returns the widget for the item at the given index.

This function was introduced in  Qt 4.1.

See also setIndexWidget().
*/
func (this *QAbstractItemView) IndexWidget(index qtcore.QModelIndex_ITF) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView11indexWidgetERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemDelegateForRow(int, QAbstractItemDelegate *)

/*
Sets the given item delegate used by this view and model for the given row. All items on row will be drawn and managed by delegate instead of using the default delegate (i.e., itemDelegate()).

Any existing row delegate for row will be removed, but not deleted. QAbstractItemView does not take ownership of delegate.

Note: If a delegate has been assigned to both a row and a column, the row delegate (i.e., this delegate) will take precedence and manage the intersecting cell index.

Warning: You should not share the same instance of a delegate between views. Doing so can cause incorrect or unintuitive editing behavior since each view connected to a given delegate may receive the closeEditor() signal, and attempt to access, modify or close an editor that has already been closed.

This function was introduced in  Qt 4.2.

See also itemDelegateForRow(), setItemDelegateForColumn(), and itemDelegate().
*/
func (this *QAbstractItemView) SetItemDelegateForRow(row int, delegate QAbstractItemDelegate_ITF /*777 QAbstractItemDelegate **/) {
	var convArg1 unsafe.Pointer
	if delegate != nil && delegate.QAbstractItemDelegate_PTR() != nil {
		convArg1 = delegate.QAbstractItemDelegate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21setItemDelegateForRowEiP21QAbstractItemDelegate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:221
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegateForRow(int) const

/*
Returns the item delegate used by this view and model for the given row, or 0 if no delegate has been assigned. You can call itemDelegate() to get a pointer to the current delegate for a given index.

This function was introduced in  Qt 4.2.

See also setItemDelegateForRow(), itemDelegateForColumn(), and setItemDelegate().
*/
func (this *QAbstractItemView) ItemDelegateForRow(row int) *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView18itemDelegateForRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemDelegateForColumn(int, QAbstractItemDelegate *)

/*
Sets the given item delegate used by this view and model for the given column. All items on column will be drawn and managed by delegate instead of using the default delegate (i.e., itemDelegate()).

Any existing column delegate for column will be removed, but not deleted. QAbstractItemView does not take ownership of delegate.

Note: If a delegate has been assigned to both a row and a column, the row delegate will take precedence and manage the intersecting cell index.

Warning: You should not share the same instance of a delegate between views. Doing so can cause incorrect or unintuitive editing behavior since each view connected to a given delegate may receive the closeEditor() signal, and attempt to access, modify or close an editor that has already been closed.

This function was introduced in  Qt 4.2.

See also itemDelegateForColumn(), setItemDelegateForRow(), and itemDelegate().
*/
func (this *QAbstractItemView) SetItemDelegateForColumn(column int, delegate QAbstractItemDelegate_ITF /*777 QAbstractItemDelegate **/) {
	var convArg1 unsafe.Pointer
	if delegate != nil && delegate.QAbstractItemDelegate_PTR() != nil {
		convArg1 = delegate.QAbstractItemDelegate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView24setItemDelegateForColumnEiP21QAbstractItemDelegate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegateForColumn(int) const

/*
Returns the item delegate used by this view and model for the given column. You can call itemDelegate() to get a pointer to the current delegate for a given index.

This function was introduced in  Qt 4.2.

See also setItemDelegateForColumn(), itemDelegateForRow(), and itemDelegate().
*/
func (this *QAbstractItemView) ItemDelegateForColumn(column int) *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView21itemDelegateForColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:228
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
Reimplemented from QWidget::inputMethodQuery().
*/
func (this *QAbstractItemView) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:233
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reset()

/*
Reset the internal state of the view.

Warning: This function will reset open editors, scroll bar positions, selections, etc. Existing changes will not be committed. If you would like to save your changes when resetting the view, you can reimplement this function, commit your changes, and then call the superclass' implementation.
*/
func (this *QAbstractItemView) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:234
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRootIndex(const QModelIndex &)

/*
Sets the root item to the item at the given index.

See also rootIndex().
*/
func (this *QAbstractItemView) SetRootIndex(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView12setRootIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:235
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doItemsLayout()

/*

 */
func (this *QAbstractItemView) DoItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13doItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:236
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void selectAll()

/*
Selects all items in the view. This function will use the selection behavior set on the view when selecting.

See also setSelection(), selectedIndexes(), and clearSelection().
*/
func (this *QAbstractItemView) SelectAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView9selectAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void edit(const QModelIndex &)

/*
Starts editing the item corresponding to the given index if it is editable.

Note that this function does not change the current index. Since the current index defines the next and previous items to edit, users may find that keyboard navigation does not work as expected. To provide consistent navigation behavior, call setCurrentIndex() before this function with the same model index.

See also QModelIndex::flags().
*/
func (this *QAbstractItemView) Edit(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView4editERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:295
// index:1
// Protected virtual Visibility=Default Availability=Available
// [1] bool edit(const QModelIndex &, QAbstractItemView::EditTrigger, QEvent *)

/*
Starts editing the item corresponding to the given index if it is editable.

Note that this function does not change the current index. Since the current index defines the next and previous items to edit, users may find that keyboard navigation does not work as expected. To provide consistent navigation behavior, call setCurrentIndex() before this function with the same model index.

See also QModelIndex::flags().
*/
func (this *QAbstractItemView) Edit1(index qtcore.QModelIndex_ITF, trigger int, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg2 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView4editERK11QModelIndexNS_11EditTriggerEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, trigger, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:238
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearSelection()

/*
Deselects all selected items. The current index will not be changed.

See also setSelection() and selectAll().
*/
func (this *QAbstractItemView) ClearSelection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14clearSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(const QModelIndex &)

/*
Sets the current item to be the item at index.

Unless the current selection mode is NoSelection, the item is also selected. Note that this function also updates the starting position for any new selections the user performs.

To set an item as the current item without selecting it, call

selectionModel()->setCurrentIndex(index, QItemSelectionModel::NoUpdate);

See also currentIndex(), currentChanged(), and selectionMode.
*/
func (this *QAbstractItemView) SetCurrentIndex(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15setCurrentIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToTop()

/*
Scrolls the view to the top.

This function was introduced in  Qt 4.1.

See also scrollTo() and scrollToBottom().
*/
func (this *QAbstractItemView) ScrollToTop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView11scrollToTopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToBottom()

/*
Scrolls the view to the bottom.

This function was introduced in  Qt 4.1.

See also scrollTo() and scrollToTop().
*/
func (this *QAbstractItemView) ScrollToBottom() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14scrollToBottomEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update(const QModelIndex &)

/*
Updates the area occupied by the given index.

This function was introduced in  Qt 4.3.
*/
func (this *QAbstractItemView) Update(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView6updateERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:246
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsInserted(const QModelIndex &, int, int)

/*
This slot is called when rows are inserted. The new rows are those under the given parent from start to end inclusive. The base class implementation calls fetchMore() on the model to check for more data.

See also rowsAboutToBeRemoved().
*/
func (this *QAbstractItemView) RowsInserted(parent qtcore.QModelIndex_ITF, start int, end_ int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView12rowsInsertedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:247
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsAboutToBeRemoved(const QModelIndex &, int, int)

/*
This slot is called when rows are about to be removed. The deleted rows are those under the given parent from start to end inclusive.

See also rowsInserted().
*/
func (this *QAbstractItemView) RowsAboutToBeRemoved(parent qtcore.QModelIndex_ITF, start int, end_ int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView20rowsAboutToBeRemovedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:248
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void selectionChanged(const QItemSelection &, const QItemSelection &)

/*
This slot is called when the selection is changed. The previous selection (which may be empty), is specified by deselected, and the new selection by selected.

See also setSelection().
*/
func (this *QAbstractItemView) SelectionChanged(selected qtcore.QItemSelection_ITF, deselected qtcore.QItemSelection_ITF) {
	var convArg0 unsafe.Pointer
	if selected != nil && selected.QItemSelection_PTR() != nil {
		convArg0 = selected.QItemSelection_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if deselected != nil && deselected.QItemSelection_PTR() != nil {
		convArg1 = deselected.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16selectionChangedERK14QItemSelectionS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:249
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)

/*
This slot is called when a new item becomes the current item. The previous current item is specified by the previous index, and the new item by the current index.

If you want to know about changes to items see the dataChanged() signal.
*/
func (this *QAbstractItemView) CurrentChanged(current qtcore.QModelIndex_ITF, previous qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QModelIndex_PTR() != nil {
		convArg0 = current.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QModelIndex_PTR() != nil {
		convArg1 = previous.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14currentChangedERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:250
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateEditorData()

/*

 */
func (this *QAbstractItemView) UpdateEditorData() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16updateEditorDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:251
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometries()

/*

 */
func (this *QAbstractItemView) UpdateEditorGeometries() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView22updateEditorGeometriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:252
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometries()

/*
Updates the geometry of the child widgets of the view.

This function was introduced in  Qt 4.4.
*/
func (this *QAbstractItemView) UpdateGeometries() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16updateGeometriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:253
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void verticalScrollbarAction(int)

/*

 */
func (this *QAbstractItemView) VerticalScrollbarAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView23verticalScrollbarActionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:254
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void horizontalScrollbarAction(int)

/*

 */
func (this *QAbstractItemView) HorizontalScrollbarAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView25horizontalScrollbarActionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:255
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void verticalScrollbarValueChanged(int)

/*

 */
func (this *QAbstractItemView) VerticalScrollbarValueChanged(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView29verticalScrollbarValueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:256
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void horizontalScrollbarValueChanged(int)

/*

 */
func (this *QAbstractItemView) HorizontalScrollbarValueChanged(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView31horizontalScrollbarValueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:257
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEditor(QWidget *, QAbstractItemDelegate::EndEditHint)

/*
Closes the given editor, and releases it. The hint is used to specify how the view should respond to the end of the editing operation. For example, the hint may indicate that the next item in the view should be opened for editing.

See also edit() and commitData().
*/
func (this *QAbstractItemView) CloseEditor(editor QWidget_ITF /*777 QWidget **/, hint int) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView11closeEditorEP7QWidgetN21QAbstractItemDelegate11EndEditHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:258
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void commitData(QWidget *)

/*
Commit the data in the editor to the model.

See also closeEditor().
*/
func (this *QAbstractItemView) CommitData(editor QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView10commitDataEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:259
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void editorDestroyed(QObject *)

/*
This function is called when the given editor has been destroyed.

See also closeEditor().
*/
func (this *QAbstractItemView) EditorDestroyed(editor qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QObject_PTR() != nil {
		convArg0 = editor.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15editorDestroyedEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pressed(const QModelIndex &)

/*
This signal is emitted when a mouse button is pressed. The item the mouse was pressed on is specified by index. The signal is only emitted when the index is valid.

Use the QApplication::mouseButtons() function to get the state of the mouse buttons.

See also activated(), clicked(), doubleClicked(), and entered().
*/
func (this *QAbstractItemView) Pressed(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView7pressedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:263
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(const QModelIndex &)

/*
This signal is emitted when a mouse button is left-clicked. The item the mouse was clicked on is specified by index. The signal is only emitted when the index is valid.

See also activated(), doubleClicked(), entered(), and pressed().
*/
func (this *QAbstractItemView) Clicked(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView7clickedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void doubleClicked(const QModelIndex &)

/*
This signal is emitted when a mouse button is double-clicked. The item the mouse was double-clicked on is specified by index. The signal is only emitted when the index is valid.

See also clicked() and activated().
*/
func (this *QAbstractItemView) DoubleClicked(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13doubleClickedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:266
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated(const QModelIndex &)

/*
This signal is emitted when the item specified by index is activated by the user. How to activate items depends on the platform; e.g., by single- or double-clicking the item, or by pressing the Return or Enter key when the item is current.

See also clicked(), doubleClicked(), entered(), and pressed().
*/
func (this *QAbstractItemView) Activated(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView9activatedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void entered(const QModelIndex &)

/*
This signal is emitted when the mouse cursor enters the item specified by index. Mouse tracking needs to be enabled for this feature to work.

See also viewportEntered(), activated(), clicked(), doubleClicked(), and pressed().
*/
func (this *QAbstractItemView) Entered(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView7enteredERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void viewportEntered()

/*
This signal is emitted when the mouse cursor enters the viewport. Mouse tracking needs to be enabled for this feature to work.

See also entered().
*/
func (this *QAbstractItemView) ViewportEntered() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15viewportEnteredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:270
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconSizeChanged(const QSize &)

/*

 */
func (this *QAbstractItemView) IconSizeChanged(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15iconSizeChangedERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:275
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setHorizontalStepsPerItem(int)

/*

 */
func (this *QAbstractItemView) SetHorizontalStepsPerItem(steps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView25setHorizontalStepsPerItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:276
// index:0
// Protected Visibility=Default Availability=Available
// [4] int horizontalStepsPerItem() const

/*

 */
func (this *QAbstractItemView) HorizontalStepsPerItem() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView22horizontalStepsPerItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:277
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setVerticalStepsPerItem(int)

/*

 */
func (this *QAbstractItemView) SetVerticalStepsPerItem(steps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView23setVerticalStepsPerItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:278
// index:0
// Protected Visibility=Default Availability=Available
// [4] int verticalStepsPerItem() const

/*

 */
func (this *QAbstractItemView) VerticalStepsPerItem() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView20verticalStepsPerItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:283
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex moveCursor(QAbstractItemView::CursorAction, Qt::KeyboardModifiers)

/*
Returns a QModelIndex object pointing to the next object in the view, based on the given cursorAction and keyboard modifiers specified by modifiers.

In the base class this is a pure virtual function.
*/
func (this *QAbstractItemView) MoveCursor(cursorAction int, modifiers int) *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView10moveCursorENS_12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cursorAction, modifiers)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:286
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [4] int horizontalOffset() const

/*
Returns the horizontal offset of the view.

In the base class this is a pure virtual function.

See also verticalOffset().
*/
func (this *QAbstractItemView) HorizontalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16horizontalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:287
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [4] int verticalOffset() const

/*
Returns the vertical offset of the view.

In the base class this is a pure virtual function.

See also horizontalOffset().
*/
func (this *QAbstractItemView) VerticalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView14verticalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:289
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [1] bool isIndexHidden(const QModelIndex &) const

/*
Returns true if the item referred to by the given index is hidden in the view, otherwise returns false.

Hiding is a view specific feature. For example in TableView a column can be marked as hidden or a row in the TreeView.

In the base class this is a pure virtual function.
*/
func (this *QAbstractItemView) IsIndexHidden(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView13isIndexHiddenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:291
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)

/*
Applies the selection flags to the items in or touched by the rectangle, rect.

When implementing your own itemview setSelection should call selectionModel()->select(selection, flags) where selection is either an empty QModelIndex or a QItemSelection that contains all items that are contained in rect.

See also selectionCommand() and selectedIndexes().
*/
func (this *QAbstractItemView) SetSelection(rect qtcore.QRect_ITF, command int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:292
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [8] QRegion visualRegionForSelection(const QItemSelection &) const

/*
Returns the region from the viewport of the items in the given selection.

In the base class this is a pure virtual function.

See also visualRect() and selectedIndexes().
*/
func (this *QAbstractItemView) VisualRegionForSelection(selection qtcore.QItemSelection_ITF) *qtgui.QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if selection != nil && selection.QItemSelection_PTR() != nil {
		convArg0 = selection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView24visualRegionForSelectionERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:293
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QModelIndexList selectedIndexes() const

/*
This convenience function returns a list of all selected and non-hidden item indexes in the view. The list contains no duplicates, and is not sorted.

See also QItemSelectionModel::selectedIndexes().
*/
func (this *QAbstractItemView) SelectedIndexes() *qtcore.QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView15selectedIndexesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:297
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QItemSelectionModel::SelectionFlags selectionCommand(const QModelIndex &, const QEvent *) const

/*
Returns the SelectionFlags to be used when updating a selection with to include the index specified. The event is a user input event, such as a mouse or keyboard event.

Reimplement this function to define your own selection behavior.

See also setSelection().
*/
func (this *QAbstractItemView) SelectionCommand(index qtcore.QModelIndex_ITF, event qtcore.QEvent_ITF /*777 const QEvent **/) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16selectionCommandERK11QModelIndexPK6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:297
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QItemSelectionModel::SelectionFlags selectionCommand(const QModelIndex &, const QEvent *) const

/*
Returns the SelectionFlags to be used when updating a selection with to include the index specified. The event is a user input event, such as a mouse or keyboard event.

Reimplement this function to define your own selection behavior.

See also setSelection().
*/
func (this *QAbstractItemView) SelectionCommandp(index qtcore.QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, const QEvent *=Pointer, QEvent=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16selectionCommandERK11QModelIndexPK6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:301
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void startDrag(Qt::DropActions)

/*
Starts a drag by calling drag->exec() using the given supportedActions.
*/
func (this *QAbstractItemView) StartDrag(supportedActions int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView9startDragE6QFlagsIN2Qt10DropActionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), supportedActions)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:304
// index:0
// Protected virtual Visibility=Default Availability=Available
// [192] QStyleOptionViewItem viewOptions() const

/*
Returns a QStyleOptionViewItem structure populated with the view's palette, font, state, alignments etc.
*/
func (this *QAbstractItemView) ViewOptions() *QStyleOptionViewItem /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView11viewOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStyleOptionViewItem)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:316
// index:0
// Protected Visibility=Default Availability=Available
// [4] QAbstractItemView::State state() const

/*
Returns the item view's state.

See also setState().
*/
func (this *QAbstractItemView) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:317
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setState(QAbstractItemView::State)

/*
Sets the item view's state to the given state.

See also state().
*/
func (this *QAbstractItemView) SetState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView8setStateENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:319
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void scheduleDelayedItemsLayout()

/*
Schedules a layout of the items in the view to be executed when the event processing starts.

Even if scheduleDelayedItemsLayout() is called multiple times before events are processed, the view will only do the layout once.

See also executeDelayedItemsLayout().
*/
func (this *QAbstractItemView) ScheduleDelayedItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView26scheduleDelayedItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:320
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void executeDelayedItemsLayout()

/*
Executes the scheduled layouts without waiting for the event processing to begin.

See also scheduleDelayedItemsLayout().
*/
func (this *QAbstractItemView) ExecuteDelayedItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView25executeDelayedItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:322
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setDirtyRegion(const QRegion &)

/*
Marks the given region as dirty and schedules it to be updated. You only need to call this function if you are implementing your own view subclass.

This function was introduced in  Qt 4.1.

See also scrollDirtyRegion() and dirtyRegionOffset().
*/
func (this *QAbstractItemView) SetDirtyRegion(region qtgui.QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg0 = region.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14setDirtyRegionERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:323
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void scrollDirtyRegion(int, int)

/*
Prepares the view for scrolling by (dx,dy) pixels by moving the dirty regions in the opposite direction. You only need to call this function if you are implementing a scrolling viewport in your view subclass.

If you implement scrollContentsBy() in a subclass of QAbstractItemView, call this function before you call QWidget::scroll() on the viewport. Alternatively, just call update().

See also scrollContentsBy(), dirtyRegionOffset(), and setDirtyRegion().
*/
func (this *QAbstractItemView) ScrollDirtyRegion(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView17scrollDirtyRegionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:324
// index:0
// Protected Visibility=Default Availability=Available
// [8] QPoint dirtyRegionOffset() const

/*
Returns the offset of the dirty regions in the view.

If you use scrollDirtyRegion() and implement a paintEvent() in a subclass of QAbstractItemView, you should translate the area given by the paint event with the offset returned from this function.

See also scrollDirtyRegion() and setDirtyRegion().
*/
func (this *QAbstractItemView) DirtyRegionOffset() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17dirtyRegionOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:326
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void startAutoScroll()

/*

 */
func (this *QAbstractItemView) StartAutoScroll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15startAutoScrollEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:327
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void stopAutoScroll()

/*

 */
func (this *QAbstractItemView) StopAutoScroll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14stopAutoScrollEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:328
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void doAutoScroll()

/*

 */
func (this *QAbstractItemView) DoAutoScroll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView12doAutoScrollEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:330
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Reimplemented from QWidget::focusNextPrevChild().
*/
func (this *QAbstractItemView) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:331
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QAbstractScrollArea::event().
*/
func (this *QAbstractItemView) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:332
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)

/*
Reimplemented from QAbstractScrollArea::viewportEvent().

This function is used to handle tool tips, and What's This? mode, if the given event is a QEvent::ToolTip,or a QEvent::WhatsThis. It passes all other events on to its base class viewportEvent() handler.
*/
func (this *QAbstractItemView) ViewportEvent(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13viewportEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:333
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QAbstractScrollArea::mousePressEvent().

This function is called with the given event when a mouse button is pressed while the cursor is inside the widget. If a valid item is pressed on it is made into the current item. This function emits the pressed() signal.
*/
func (this *QAbstractItemView) MousePressEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:334
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QAbstractScrollArea::mouseMoveEvent().

This function is called with the given event when a mouse move event is sent to the widget. If a selection is in progress and new items are moved over the selection is extended; if a drag is in progress it is continued.
*/
func (this *QAbstractItemView) MouseMoveEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:335
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QAbstractScrollArea::mouseReleaseEvent().

This function is called with the given event when a mouse button is released, after a mouse press event on the widget. If a user presses the mouse inside your widget and then drags the mouse to another location before releasing the mouse button, your widget receives the release event. The function will emit the clicked() signal if an item was being pressed.
*/
func (this *QAbstractItemView) MouseReleaseEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:336
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
Reimplemented from QAbstractScrollArea::mouseDoubleClickEvent().

This function is called with the given event when a mouse button is double clicked inside the widget. If the double-click is on a valid item it emits the doubleClicked() signal and calls edit() on the item.
*/
func (this *QAbstractItemView) MouseDoubleClickEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:338
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
Reimplemented from QAbstractScrollArea::dragEnterEvent().

This function is called with the given event when a drag and drop operation enters the widget. If the drag is over a valid dropping place (e.g. over an item that accepts drops), the event is accepted; otherwise it is ignored.

See also dropEvent() and startDrag().
*/
func (this *QAbstractItemView) DragEnterEvent(event qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragEnterEvent_PTR() != nil {
		convArg0 = event.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:339
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QAbstractScrollArea::dragMoveEvent().

This function is called continuously with the given event during a drag and drop operation over the widget. It can cause the view to scroll if, for example, the user drags a selection to view's right or bottom edge. In this case, the event will be accepted; otherwise it will be ignored.

See also dropEvent() and startDrag().
*/
func (this *QAbstractItemView) DragMoveEvent(event qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragMoveEvent_PTR() != nil {
		convArg0 = event.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:340
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
Reimplemented from QAbstractScrollArea::dragLeaveEvent().

This function is called when the item being dragged leaves the view. The event describes the state of the drag and drop operation.
*/
func (this *QAbstractItemView) DragLeaveEvent(event qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragLeaveEvent_PTR() != nil {
		convArg0 = event.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:341
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QAbstractScrollArea::dropEvent().

This function is called with the given event when a drop event occurs over the widget. If the model accepts the even position the drop event is accepted; otherwise it is ignored.

See also startDrag().
*/
func (this *QAbstractItemView) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDropEvent_PTR() != nil {
		convArg0 = event.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:343
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().

This function is called with the given event when the widget obtains the focus. By default, the event is ignored.

See also setFocus() and focusOutEvent().
*/
func (this *QAbstractItemView) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:344
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().

This function is called with the given event when the widget loses the focus. By default, the event is ignored.

See also clearFocus() and focusInEvent().
*/
func (this *QAbstractItemView) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:345
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QAbstractScrollArea::keyPressEvent().

This function is called with the given event when a key event is sent to the widget. The default implementation handles basic cursor movement, e.g. Up, Down, Left, Right, Home, PageUp, and PageDown; the activated() signal is emitted if the current index is valid and the activation key is pressed (e.g. Enter or Return, depending on the platform). This function is where editing is initiated by key press, e.g. if F2 is pressed.

See also edit(), moveCursor(), keyboardSearch(), and tabKeyNavigation.
*/
func (this *QAbstractItemView) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:346
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QAbstractScrollArea::resizeEvent().

This function is called with the given event when a resize event is sent to the widget.

See also QWidget::resizeEvent().
*/
func (this *QAbstractItemView) ResizeEvent(event qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QResizeEvent_PTR() != nil {
		convArg0 = event.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:347
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().

This function is called with the given event when a timer event is sent to the widget.

See also QObject::timerEvent().
*/
func (this *QAbstractItemView) TimerEvent(event qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QTimerEvent_PTR() != nil {
		convArg0 = event.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:348
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
Reimplemented from QWidget::inputMethodEvent().
*/
func (this *QAbstractItemView) InputMethodEvent(event qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QInputMethodEvent_PTR() != nil {
		convArg0 = event.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:349
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*

 */
func (this *QAbstractItemView) EventFilter(object qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:353
// index:0
// Protected Visibility=Default Availability=Available
// [4] QAbstractItemView::DropIndicatorPosition dropIndicatorPosition() const

/*
Returns the position of the drop indicator in relation to the closest item.

This function was introduced in  Qt 4.1.
*/
func (this *QAbstractItemView) DropIndicatorPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView21dropIndicatorPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:356
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint() const

/*
Reimplemented from QAbstractScrollArea::viewportSizeHint().

This function was introduced in  Qt 5.2.
*/
func (this *QAbstractItemView) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16viewportSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

/*
This enum indicates how the view responds to user selections:



The most commonly used modes are SingleSelection and ExtendedSelection.

*/
type QAbstractItemView__SelectionMode = int

// Items cannot be selected.
const QAbstractItemView__NoSelection QAbstractItemView__SelectionMode = 0

// When the user selects an item, any already-selected item becomes unselected. It is possible for the user to deselect the selected item.
const QAbstractItemView__SingleSelection QAbstractItemView__SelectionMode = 1

// When the user selects an item in the usual way, the selection status of that item is toggled and the other items are left alone. Multiple items can be toggled by dragging the mouse over them.
const QAbstractItemView__MultiSelection QAbstractItemView__SelectionMode = 2

// When the user selects an item in the usual way, the selection is cleared and the new item selected. However, if the user presses the Ctrl key when clicking on an item, the clicked item gets toggled and all other items are left untouched. If the user presses the Shift key while clicking on an item, all items between the current item and the clicked item are selected or unselected, depending on the state of the clicked item. Multiple items can be selected by dragging the mouse over them.
const QAbstractItemView__ExtendedSelection QAbstractItemView__SelectionMode = 3

// When the user selects an item in the usual way, the selection is cleared and the new item selected. However, if the user presses the Shift key while clicking on an item, all items between the current item and the clicked item are selected or unselected, depending on the state of the clicked item.
const QAbstractItemView__ContiguousSelection QAbstractItemView__SelectionMode = 4

func (this *QAbstractItemView) SelectionModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_SelectionModeItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.SelectionModeItemName(val)
}

/*

 */
type QAbstractItemView__SelectionBehavior = int

// Selecting single items.
const QAbstractItemView__SelectItems QAbstractItemView__SelectionBehavior = 0

// Selecting only rows.
const QAbstractItemView__SelectRows QAbstractItemView__SelectionBehavior = 1

// Selecting only columns.
const QAbstractItemView__SelectColumns QAbstractItemView__SelectionBehavior = 2

func (this *QAbstractItemView) SelectionBehaviorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_SelectionBehaviorItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.SelectionBehaviorItemName(val)
}

/*

 */
type QAbstractItemView__ScrollHint = int

// Scroll to ensure that the item is visible.
const QAbstractItemView__EnsureVisible QAbstractItemView__ScrollHint = 0

// Scroll to position the item at the top of the viewport.
const QAbstractItemView__PositionAtTop QAbstractItemView__ScrollHint = 1

// Scroll to position the item at the bottom of the viewport.
const QAbstractItemView__PositionAtBottom QAbstractItemView__ScrollHint = 2

// Scroll to position the item at the center of the viewport.
const QAbstractItemView__PositionAtCenter QAbstractItemView__ScrollHint = 3

func (this *QAbstractItemView) ScrollHintItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_ScrollHintItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.ScrollHintItemName(val)
}

/*


 */
type QAbstractItemView__EditTrigger = int

//
const QAbstractItemView__NoEditTriggers QAbstractItemView__EditTrigger = 0

//
const QAbstractItemView__CurrentChanged QAbstractItemView__EditTrigger = 1

//
const QAbstractItemView__DoubleClicked QAbstractItemView__EditTrigger = 2

//
const QAbstractItemView__SelectedClicked QAbstractItemView__EditTrigger = 4

//
const QAbstractItemView__EditKeyPressed QAbstractItemView__EditTrigger = 8

//
const QAbstractItemView__AnyKeyPressed QAbstractItemView__EditTrigger = 16

//
const QAbstractItemView__AllEditTriggers QAbstractItemView__EditTrigger = 31

func (this *QAbstractItemView) EditTriggerItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_EditTriggerItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.EditTriggerItemName(val)
}

/*
Describes how the scrollbar should behave. When setting the scroll mode to ScrollPerPixel the single step size will adjust automatically unless it was set explicitly using setSingleStep(). The automatic adjustment can be restored by setting the single step size to -1.



This enum was introduced or modified in  Qt 4.2.

*/
type QAbstractItemView__ScrollMode = int

// The view will scroll the contents one item at a time.
const QAbstractItemView__ScrollPerItem QAbstractItemView__ScrollMode = 0

// The view will scroll the contents one pixel at a time.
const QAbstractItemView__ScrollPerPixel QAbstractItemView__ScrollMode = 1

func (this *QAbstractItemView) ScrollModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_ScrollModeItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.ScrollModeItemName(val)
}

/*
Describes the various drag and drop events the view can act upon. By default the view does not support dragging or dropping (NoDragDrop).



Note that the model used needs to provide support for drag and drop operations.

This enum was introduced or modified in  Qt 4.2.

See also setDragDropMode() and Using drag and drop with item views.

*/
type QAbstractItemView__DragDropMode = int

// Does not support dragging or dropping.
const QAbstractItemView__NoDragDrop QAbstractItemView__DragDropMode = 0

// The view supports dragging of its own items
const QAbstractItemView__DragOnly QAbstractItemView__DragDropMode = 1

// The view accepts drops
const QAbstractItemView__DropOnly QAbstractItemView__DragDropMode = 2

// The view supports both dragging and dropping
const QAbstractItemView__DragDrop QAbstractItemView__DragDropMode = 3

// The view accepts move (not copy) operations only from itself.
const QAbstractItemView__InternalMove QAbstractItemView__DragDropMode = 4

func (this *QAbstractItemView) DragDropModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_DragDropModeItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.DragDropModeItemName(val)
}

/*
This enum describes the different ways to navigate between items,



See also moveCursor().

*/
type QAbstractItemView__CursorAction = int

// Move to the item above the current item.
const QAbstractItemView__MoveUp QAbstractItemView__CursorAction = 0

// Move to the item below the current item.
const QAbstractItemView__MoveDown QAbstractItemView__CursorAction = 1

// Move to the item left of the current item.
const QAbstractItemView__MoveLeft QAbstractItemView__CursorAction = 2

// Move to the item right of the current item.
const QAbstractItemView__MoveRight QAbstractItemView__CursorAction = 3

// Move to the top-left corner item.
const QAbstractItemView__MoveHome QAbstractItemView__CursorAction = 4

// Move to the bottom-right corner item.
const QAbstractItemView__MoveEnd QAbstractItemView__CursorAction = 5

// Move one page up above the current item.
const QAbstractItemView__MovePageUp QAbstractItemView__CursorAction = 6

// Move one page down below the current item.
const QAbstractItemView__MovePageDown QAbstractItemView__CursorAction = 7

// Move to the item after the current item.
const QAbstractItemView__MoveNext QAbstractItemView__CursorAction = 8

// Move to the item before the current item.
const QAbstractItemView__MovePrevious QAbstractItemView__CursorAction = 9

func (this *QAbstractItemView) CursorActionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_CursorActionItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.CursorActionItemName(val)
}

/*
Describes the different states the view can be in. This is usually only interesting when reimplementing your own view.


*/
type QAbstractItemView__State = int

// The is the default state.
const QAbstractItemView__NoState QAbstractItemView__State = 0

// The user is dragging items.
const QAbstractItemView__DraggingState QAbstractItemView__State = 1

// The user is selecting items.
const QAbstractItemView__DragSelectingState QAbstractItemView__State = 2

// The user is editing an item in a widget editor.
const QAbstractItemView__EditingState QAbstractItemView__State = 3

// The user is opening a branch of items.
const QAbstractItemView__ExpandingState QAbstractItemView__State = 4

// The user is closing a branch of items.
const QAbstractItemView__CollapsingState QAbstractItemView__State = 5

// The item view is performing an animation.
const QAbstractItemView__AnimatingState QAbstractItemView__State = 6

func (this *QAbstractItemView) StateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_StateItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.StateItemName(val)
}

/*
This enum indicates the position of the drop indicator in relation to the index at the current mouse position:


*/
type QAbstractItemView__DropIndicatorPosition = int

// The item will be dropped on the index.
const QAbstractItemView__OnItem QAbstractItemView__DropIndicatorPosition = 0

// The item will be dropped above the index.
const QAbstractItemView__AboveItem QAbstractItemView__DropIndicatorPosition = 1

// The item will be dropped below the index.
const QAbstractItemView__BelowItem QAbstractItemView__DropIndicatorPosition = 2

// The item will be dropped onto a region of the viewport with no items. The way each view handles items dropped onto the viewport depends on the behavior of the underlying model in use.
const QAbstractItemView__OnViewport QAbstractItemView__DropIndicatorPosition = 3

func (this *QAbstractItemView) DropIndicatorPositionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemView_DropIndicatorPositionItemName(val int) string {
	var nilthis *QAbstractItemView
	return nilthis.DropIndicatorPositionItemName(val)
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
