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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void rowsInserted(const class QModelIndex &, int, int)
func (this *QAbstractItemView) InheritRowsInserted(f func(parent *qtcore.QModelIndex, start int, end int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsInserted", f)
}

// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QAbstractItemView) InheritRowsAboutToBeRemoved(f func(parent *qtcore.QModelIndex, start int, end int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsAboutToBeRemoved", f)
}

// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QAbstractItemView) InheritSelectionChanged(f func(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) /*void*/) {
	qtrt.SetAllInheritCallback(this, "selectionChanged", f)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
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

// void closeEditor(class QWidget *, class QAbstractItemDelegate::EndEditHint)
func (this *QAbstractItemView) InheritCloseEditor(f func(editor *QWidget /*777 QWidget **/, hint int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEditor", f)
}

// void commitData(class QWidget *)
func (this *QAbstractItemView) InheritCommitData(f func(editor *QWidget /*777 QWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "commitData", f)
}

// void editorDestroyed(class QObject *)
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

// QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
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

// bool isIndexHidden(const class QModelIndex &)
func (this *QAbstractItemView) InheritIsIndexHidden(f func(index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "isIndexHidden", f)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QAbstractItemView) InheritSetSelection(f func(rect *qtcore.QRect, command int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSelection", f)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QAbstractItemView) InheritVisualRegionForSelection(f func(selection *qtcore.QItemSelection) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "visualRegionForSelection", f)
}

// QModelIndexList selectedIndexes()
func (this *QAbstractItemView) InheritSelectedIndexes(f func() *qtcore.QModelIndexList /*9999*/) {
	qtrt.SetAllInheritCallback(this, "selectedIndexes", f)
}

// bool edit(const class QModelIndex &, enum QAbstractItemView::EditTrigger, class QEvent *)
func (this *QAbstractItemView) InheritEdit(f func(index *qtcore.QModelIndex, trigger int, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "edit", f)
}

// QItemSelectionModel::SelectionFlags selectionCommand(const class QModelIndex &, const class QEvent *)
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

// void setState(enum QAbstractItemView::State)
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

// void setDirtyRegion(const class QRegion &)
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

// bool focusNextPrevChild(_Bool)
func (this *QAbstractItemView) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// bool event(class QEvent *)
func (this *QAbstractItemView) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool viewportEvent(class QEvent *)
func (this *QAbstractItemView) InheritViewportEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "viewportEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractItemView) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractItemView) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractItemView) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QAbstractItemView) InheritMouseDoubleClickEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void dragEnterEvent(class QDragEnterEvent *)
func (this *QAbstractItemView) InheritDragEnterEvent(f func(event *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(class QDragMoveEvent *)
func (this *QAbstractItemView) InheritDragMoveEvent(f func(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QAbstractItemView) InheritDragLeaveEvent(f func(event *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(class QDropEvent *)
func (this *QAbstractItemView) InheritDropEvent(f func(event *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QAbstractItemView) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractItemView) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractItemView) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QAbstractItemView) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QAbstractItemView) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void inputMethodEvent(class QInputMethodEvent *)
func (this *QAbstractItemView) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// QAbstractItemView::DropIndicatorPosition dropIndicatorPosition()
func (this *QAbstractItemView) InheritDropIndicatorPosition(f func() int) {
	qtrt.SetAllInheritCallback(this, "dropIndicatorPosition", f)
}

// QSize viewportSizeHint()
func (this *QAbstractItemView) InheritViewportSizeHint(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewportSizeHint", f)
}

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
// [8] const QMetaObject * metaObject()
func (this *QAbstractItemView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemView(QWidget *)
func NewQAbstractItemView(parent QWidget_ITF /*777 QWidget **/) *QAbstractItemView {
	var convArg0 = parent.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractItemViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:128
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractItemView()
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
func (this *QAbstractItemView) SetModel(model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * model()
func (this *QAbstractItemView) Model() *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:133
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSelectionModel(QItemSelectionModel *)
func (this *QAbstractItemView) SetSelectionModel(selectionModel qtcore.QItemSelectionModel_ITF /*777 QItemSelectionModel **/) {
	var convArg0 = selectionModel.QItemSelectionModel_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView17setSelectionModelEP19QItemSelectionModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QItemSelectionModel * selectionModel()
func (this *QAbstractItemView) SelectionModel() *qtcore.QItemSelectionModel /*777 QItemSelectionModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView14selectionModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQItemSelectionModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemDelegate(QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegate(delegate QAbstractItemDelegate_ITF /*777 QAbstractItemDelegate **/) {
	var convArg0 = delegate.QAbstractItemDelegate_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15setItemDelegateEP21QAbstractItemDelegate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegate()
func (this *QAbstractItemView) ItemDelegate() *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView12itemDelegateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:226
// index:1
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegate(const QModelIndex &)
func (this *QAbstractItemView) ItemDelegate_1(index qtcore.QModelIndex_ITF) *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView12itemDelegateERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionMode(QAbstractItemView::SelectionMode)
func (this *QAbstractItemView) SetSelectionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16setSelectionModeENS_13SelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::SelectionMode selectionMode()
func (this *QAbstractItemView) SelectionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView13selectionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionBehavior(QAbstractItemView::SelectionBehavior)
func (this *QAbstractItemView) SetSelectionBehavior(behavior int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView20setSelectionBehaviorENS_17SelectionBehaviorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), behavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::SelectionBehavior selectionBehavior()
func (this *QAbstractItemView) SelectionBehavior() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17selectionBehaviorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:145
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex currentIndex()
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
// [24] QModelIndex rootIndex()
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
func (this *QAbstractItemView) SetEditTriggers(triggers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15setEditTriggersE6QFlagsINS_11EditTriggerEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), triggers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:149
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::EditTriggers editTriggers()
func (this *QAbstractItemView) EditTriggers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView12editTriggersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalScrollMode(enum QAbstractItemView::ScrollMode)
func (this *QAbstractItemView) SetVerticalScrollMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21setVerticalScrollModeENS_10ScrollModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:152
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::ScrollMode verticalScrollMode()
func (this *QAbstractItemView) VerticalScrollMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView18verticalScrollModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetVerticalScrollMode()
func (this *QAbstractItemView) ResetVerticalScrollMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView23resetVerticalScrollModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalScrollMode(enum QAbstractItemView::ScrollMode)
func (this *QAbstractItemView) SetHorizontalScrollMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView23setHorizontalScrollModeENS_10ScrollModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:156
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::ScrollMode horizontalScrollMode()
func (this *QAbstractItemView) HorizontalScrollMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView20horizontalScrollModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetHorizontalScrollMode()
func (this *QAbstractItemView) ResetHorizontalScrollMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView25resetHorizontalScrollModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoScroll(_Bool)
func (this *QAbstractItemView) SetAutoScroll(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13setAutoScrollEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAutoScroll()
func (this *QAbstractItemView) HasAutoScroll() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView13hasAutoScrollEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoScrollMargin(int)
func (this *QAbstractItemView) SetAutoScrollMargin(margin int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView19setAutoScrollMarginEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:163
// index:0
// Public Visibility=Default Availability=Available
// [4] int autoScrollMargin()
func (this *QAbstractItemView) AutoScrollMargin() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16autoScrollMarginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabKeyNavigation(_Bool)
func (this *QAbstractItemView) SetTabKeyNavigation(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView19setTabKeyNavigationEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabKeyNavigation()
func (this *QAbstractItemView) TabKeyNavigation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16tabKeyNavigationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDropIndicatorShown(_Bool)
func (this *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21setDropIndicatorShownEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool showDropIndicator()
func (this *QAbstractItemView) ShowDropIndicator() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17showDropIndicatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragEnabled(_Bool)
func (this *QAbstractItemView) SetDragEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14setDragEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dragEnabled()
func (this *QAbstractItemView) DragEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView11dragEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragDropOverwriteMode(_Bool)
func (this *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView24setDragDropOverwriteModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), overwrite)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:176
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dragDropOverwriteMode()
func (this *QAbstractItemView) DragDropOverwriteMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView21dragDropOverwriteModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragDropMode(enum QAbstractItemView::DragDropMode)
func (this *QAbstractItemView) SetDragDropMode(behavior int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15setDragDropModeENS_12DragDropModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), behavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:188
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractItemView::DragDropMode dragDropMode()
func (this *QAbstractItemView) DragDropMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView12dragDropModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultDropAction(Qt::DropAction)
func (this *QAbstractItemView) SetDefaultDropAction(dropAction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView20setDefaultDropActionEN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dropAction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:191
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DropAction defaultDropAction()
func (this *QAbstractItemView) DefaultDropAction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17defaultDropActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlternatingRowColors(_Bool)
func (this *QAbstractItemView) SetAlternatingRowColors(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView23setAlternatingRowColorsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:195
// index:0
// Public Visibility=Default Availability=Available
// [1] bool alternatingRowColors()
func (this *QAbstractItemView) AlternatingRowColors() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView20alternatingRowColorsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)
func (this *QAbstractItemView) SetIconSize(size qtcore.QSize_ITF) {
	var convArg0 = size.QSize_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView11setIconSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:198
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize()
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
func (this *QAbstractItemView) SetTextElideMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16setTextElideModeEN2Qt13TextElideModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:201
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextElideMode textElideMode()
func (this *QAbstractItemView) TextElideMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView13textElideModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:203
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void keyboardSearch(const QString &)
func (this *QAbstractItemView) KeyboardSearch(search string) {
	var tmpArg0 = qtcore.NewQString_5(search)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14keyboardSearchERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:205
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QRect visualRect(const QModelIndex &)
func (this *QAbstractItemView) VisualRect(index qtcore.QModelIndex_ITF) *qtcore.QRect /*123*/ {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView10visualRectERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:206
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QAbstractItemView) ScrollTo(index qtcore.QModelIndex_ITF, hint int) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView8scrollToERK11QModelIndexNS_10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:207
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex indexAt(const QPoint &)
func (this *QAbstractItemView) IndexAt(point qtcore.QPoint_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg0 = point.QPoint_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView7indexAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:209
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizeHintForIndex(const QModelIndex &)
func (this *QAbstractItemView) SizeHintForIndex(index qtcore.QModelIndex_ITF) *qtcore.QSize /*123*/ {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16sizeHintForIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:210
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int sizeHintForRow(int)
func (this *QAbstractItemView) SizeHintForRow(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView14sizeHintForRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:211
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int sizeHintForColumn(int)
func (this *QAbstractItemView) SizeHintForColumn(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView17sizeHintForColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void openPersistentEditor(const QModelIndex &)
func (this *QAbstractItemView) OpenPersistentEditor(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView20openPersistentEditorERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closePersistentEditor(const QModelIndex &)
func (this *QAbstractItemView) ClosePersistentEditor(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21closePersistentEditorERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:215
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentEditorOpen(const QModelIndex &)
func (this *QAbstractItemView) IsPersistentEditorOpen(index qtcore.QModelIndex_ITF) bool {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView22isPersistentEditorOpenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndexWidget(const QModelIndex &, QWidget *)
func (this *QAbstractItemView) SetIndexWidget(index qtcore.QModelIndex_ITF, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	var convArg1 = widget.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14setIndexWidgetERK11QModelIndexP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:218
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * indexWidget(const QModelIndex &)
func (this *QAbstractItemView) IndexWidget(index qtcore.QModelIndex_ITF) *QWidget /*777 QWidget **/ {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView11indexWidgetERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemDelegateForRow(int, QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegateForRow(row int, delegate QAbstractItemDelegate_ITF /*777 QAbstractItemDelegate **/) {
	var convArg1 = delegate.QAbstractItemDelegate_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21setItemDelegateForRowEiP21QAbstractItemDelegate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:221
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegateForRow(int)
func (this *QAbstractItemView) ItemDelegateForRow(row int) *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView18itemDelegateForRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemDelegateForColumn(int, QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegateForColumn(column int, delegate QAbstractItemDelegate_ITF /*777 QAbstractItemDelegate **/) {
	var convArg1 = delegate.QAbstractItemDelegate_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView24setItemDelegateForColumnEiP21QAbstractItemDelegate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegateForColumn(int)
func (this *QAbstractItemView) ItemDelegateForColumn(column int) *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView21itemDelegateForColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:228
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery)
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
func (this *QAbstractItemView) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:234
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRootIndex(const QModelIndex &)
func (this *QAbstractItemView) SetRootIndex(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView12setRootIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:235
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doItemsLayout()
func (this *QAbstractItemView) DoItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13doItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:236
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void selectAll()
func (this *QAbstractItemView) SelectAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView9selectAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void edit(const QModelIndex &)
func (this *QAbstractItemView) Edit(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView4editERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:295
// index:1
// Protected virtual Visibility=Default Availability=Available
// [1] bool edit(const QModelIndex &, enum QAbstractItemView::EditTrigger, QEvent *)
func (this *QAbstractItemView) Edit_1(index qtcore.QModelIndex_ITF, trigger int, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	var convArg2 = event.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView4editERK11QModelIndexNS_11EditTriggerEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, trigger, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:238
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearSelection()
func (this *QAbstractItemView) ClearSelection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14clearSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(const QModelIndex &)
func (this *QAbstractItemView) SetCurrentIndex(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15setCurrentIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToTop()
func (this *QAbstractItemView) ScrollToTop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView11scrollToTopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToBottom()
func (this *QAbstractItemView) ScrollToBottom() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14scrollToBottomEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update(const QModelIndex &)
func (this *QAbstractItemView) Update(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView6updateERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:246
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsInserted(const QModelIndex &, int, int)
func (this *QAbstractItemView) RowsInserted(parent qtcore.QModelIndex_ITF, start int, end int) {
	var convArg0 = parent.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView12rowsInsertedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:247
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsAboutToBeRemoved(const QModelIndex &, int, int)
func (this *QAbstractItemView) RowsAboutToBeRemoved(parent qtcore.QModelIndex_ITF, start int, end int) {
	var convArg0 = parent.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView20rowsAboutToBeRemovedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:248
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void selectionChanged(const QItemSelection &, const QItemSelection &)
func (this *QAbstractItemView) SelectionChanged(selected qtcore.QItemSelection_ITF, deselected qtcore.QItemSelection_ITF) {
	var convArg0 = selected.QItemSelection_PTR().GetCthis()
	var convArg1 = deselected.QItemSelection_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16selectionChangedERK14QItemSelectionS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:249
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)
func (this *QAbstractItemView) CurrentChanged(current qtcore.QModelIndex_ITF, previous qtcore.QModelIndex_ITF) {
	var convArg0 = current.QModelIndex_PTR().GetCthis()
	var convArg1 = previous.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14currentChangedERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:250
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateEditorData()
func (this *QAbstractItemView) UpdateEditorData() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16updateEditorDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:251
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometries()
func (this *QAbstractItemView) UpdateEditorGeometries() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView22updateEditorGeometriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:252
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometries()
func (this *QAbstractItemView) UpdateGeometries() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16updateGeometriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:253
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void verticalScrollbarAction(int)
func (this *QAbstractItemView) VerticalScrollbarAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView23verticalScrollbarActionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:254
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void horizontalScrollbarAction(int)
func (this *QAbstractItemView) HorizontalScrollbarAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView25horizontalScrollbarActionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:255
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void verticalScrollbarValueChanged(int)
func (this *QAbstractItemView) VerticalScrollbarValueChanged(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView29verticalScrollbarValueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:256
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void horizontalScrollbarValueChanged(int)
func (this *QAbstractItemView) HorizontalScrollbarValueChanged(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView31horizontalScrollbarValueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:257
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEditor(QWidget *, QAbstractItemDelegate::EndEditHint)
func (this *QAbstractItemView) CloseEditor(editor QWidget_ITF /*777 QWidget **/, hint int) {
	var convArg0 = editor.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView11closeEditorEP7QWidgetN21QAbstractItemDelegate11EndEditHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:258
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void commitData(QWidget *)
func (this *QAbstractItemView) CommitData(editor QWidget_ITF /*777 QWidget **/) {
	var convArg0 = editor.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView10commitDataEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:259
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void editorDestroyed(QObject *)
func (this *QAbstractItemView) EditorDestroyed(editor qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 = editor.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15editorDestroyedEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pressed(const QModelIndex &)
func (this *QAbstractItemView) Pressed(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView7pressedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:263
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(const QModelIndex &)
func (this *QAbstractItemView) Clicked(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView7clickedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void doubleClicked(const QModelIndex &)
func (this *QAbstractItemView) DoubleClicked(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13doubleClickedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:266
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated(const QModelIndex &)
func (this *QAbstractItemView) Activated(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView9activatedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void entered(const QModelIndex &)
func (this *QAbstractItemView) Entered(index qtcore.QModelIndex_ITF) {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView7enteredERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void viewportEntered()
func (this *QAbstractItemView) ViewportEntered() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15viewportEnteredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:270
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconSizeChanged(const QSize &)
func (this *QAbstractItemView) IconSizeChanged(size qtcore.QSize_ITF) {
	var convArg0 = size.QSize_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15iconSizeChangedERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:275
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setHorizontalStepsPerItem(int)
func (this *QAbstractItemView) SetHorizontalStepsPerItem(steps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView25setHorizontalStepsPerItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:276
// index:0
// Protected Visibility=Default Availability=Available
// [4] int horizontalStepsPerItem()
func (this *QAbstractItemView) HorizontalStepsPerItem() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView22horizontalStepsPerItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:277
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setVerticalStepsPerItem(int)
func (this *QAbstractItemView) SetVerticalStepsPerItem(steps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView23setVerticalStepsPerItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:278
// index:0
// Protected Visibility=Default Availability=Available
// [4] int verticalStepsPerItem()
func (this *QAbstractItemView) VerticalStepsPerItem() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView20verticalStepsPerItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:283
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
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
// [4] int horizontalOffset()
func (this *QAbstractItemView) HorizontalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16horizontalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:287
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [4] int verticalOffset()
func (this *QAbstractItemView) VerticalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView14verticalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:289
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [1] bool isIndexHidden(const QModelIndex &)
func (this *QAbstractItemView) IsIndexHidden(index qtcore.QModelIndex_ITF) bool {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView13isIndexHiddenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:291
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)
func (this *QAbstractItemView) SetSelection(rect qtcore.QRect_ITF, command int) {
	var convArg0 = rect.QRect_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:292
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [8] QRegion visualRegionForSelection(const QItemSelection &)
func (this *QAbstractItemView) VisualRegionForSelection(selection qtcore.QItemSelection_ITF) *qtgui.QRegion /*123*/ {
	var convArg0 = selection.QItemSelection_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView24visualRegionForSelectionERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:293
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QModelIndexList selectedIndexes()
func (this *QAbstractItemView) SelectedIndexes() *qtcore.QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView15selectedIndexesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:297
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QItemSelectionModel::SelectionFlags selectionCommand(const QModelIndex &, const QEvent *)
func (this *QAbstractItemView) SelectionCommand(index qtcore.QModelIndex_ITF, event qtcore.QEvent_ITF /*777 const QEvent **/) int {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	var convArg1 = event.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16selectionCommandERK11QModelIndexPK6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:301
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void startDrag(Qt::DropActions)
func (this *QAbstractItemView) StartDrag(supportedActions int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView9startDragE6QFlagsIN2Qt10DropActionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), supportedActions)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:304
// index:0
// Protected virtual Visibility=Default Availability=Available
// [192] QStyleOptionViewItem viewOptions()
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
// [4] QAbstractItemView::State state()
func (this *QAbstractItemView) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:317
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setState(enum QAbstractItemView::State)
func (this *QAbstractItemView) SetState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView8setStateENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:319
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void scheduleDelayedItemsLayout()
func (this *QAbstractItemView) ScheduleDelayedItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView26scheduleDelayedItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:320
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void executeDelayedItemsLayout()
func (this *QAbstractItemView) ExecuteDelayedItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView25executeDelayedItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:322
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setDirtyRegion(const QRegion &)
func (this *QAbstractItemView) SetDirtyRegion(region qtgui.QRegion_ITF) {
	var convArg0 = region.QRegion_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14setDirtyRegionERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:323
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void scrollDirtyRegion(int, int)
func (this *QAbstractItemView) ScrollDirtyRegion(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView17scrollDirtyRegionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:324
// index:0
// Protected Visibility=Default Availability=Available
// [8] QPoint dirtyRegionOffset()
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
func (this *QAbstractItemView) StartAutoScroll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15startAutoScrollEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:327
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void stopAutoScroll()
func (this *QAbstractItemView) StopAutoScroll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14stopAutoScrollEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:328
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void doAutoScroll()
func (this *QAbstractItemView) DoAutoScroll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView12doAutoScrollEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:330
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(_Bool)
func (this *QAbstractItemView) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:331
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAbstractItemView) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 = event.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:332
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)
func (this *QAbstractItemView) ViewportEvent(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 = event.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13viewportEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:333
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QAbstractItemView) MousePressEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 = event.QMouseEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:334
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QAbstractItemView) MouseMoveEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 = event.QMouseEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:335
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QAbstractItemView) MouseReleaseEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 = event.QMouseEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:336
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QAbstractItemView) MouseDoubleClickEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 = event.QMouseEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:338
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)
func (this *QAbstractItemView) DragEnterEvent(event qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 = event.QDragEnterEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:339
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)
func (this *QAbstractItemView) DragMoveEvent(event qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 = event.QDragMoveEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:340
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)
func (this *QAbstractItemView) DragLeaveEvent(event qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 = event.QDragLeaveEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:341
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)
func (this *QAbstractItemView) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 = event.QDropEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:343
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QAbstractItemView) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 = event.QFocusEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:344
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QAbstractItemView) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 = event.QFocusEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:345
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QAbstractItemView) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 = event.QKeyEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:346
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QAbstractItemView) ResizeEvent(event qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 = event.QResizeEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:347
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QAbstractItemView) TimerEvent(event qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 = event.QTimerEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:348
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)
func (this *QAbstractItemView) InputMethodEvent(event qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 = event.QInputMethodEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractItemView16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:352
// index:0
// Protected Visibility=Default Availability=Available
// [4] QAbstractItemView::DropIndicatorPosition dropIndicatorPosition()
func (this *QAbstractItemView) DropIndicatorPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView21dropIndicatorPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:355
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint()
func (this *QAbstractItemView) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAbstractItemView16viewportSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

type QAbstractItemView__SelectionMode = int

const QAbstractItemView__NoSelection QAbstractItemView__SelectionMode = 0
const QAbstractItemView__SingleSelection QAbstractItemView__SelectionMode = 1
const QAbstractItemView__MultiSelection QAbstractItemView__SelectionMode = 2
const QAbstractItemView__ExtendedSelection QAbstractItemView__SelectionMode = 3
const QAbstractItemView__ContiguousSelection QAbstractItemView__SelectionMode = 4

type QAbstractItemView__SelectionBehavior = int

const QAbstractItemView__SelectItems QAbstractItemView__SelectionBehavior = 0
const QAbstractItemView__SelectRows QAbstractItemView__SelectionBehavior = 1
const QAbstractItemView__SelectColumns QAbstractItemView__SelectionBehavior = 2

type QAbstractItemView__ScrollHint = int

const QAbstractItemView__EnsureVisible QAbstractItemView__ScrollHint = 0
const QAbstractItemView__PositionAtTop QAbstractItemView__ScrollHint = 1
const QAbstractItemView__PositionAtBottom QAbstractItemView__ScrollHint = 2
const QAbstractItemView__PositionAtCenter QAbstractItemView__ScrollHint = 3

type QAbstractItemView__EditTrigger = int

const QAbstractItemView__NoEditTriggers QAbstractItemView__EditTrigger = 0
const QAbstractItemView__CurrentChanged QAbstractItemView__EditTrigger = 1
const QAbstractItemView__DoubleClicked QAbstractItemView__EditTrigger = 2
const QAbstractItemView__SelectedClicked QAbstractItemView__EditTrigger = 4
const QAbstractItemView__EditKeyPressed QAbstractItemView__EditTrigger = 8
const QAbstractItemView__AnyKeyPressed QAbstractItemView__EditTrigger = 16
const QAbstractItemView__AllEditTriggers QAbstractItemView__EditTrigger = 31

type QAbstractItemView__ScrollMode = int

const QAbstractItemView__ScrollPerItem QAbstractItemView__ScrollMode = 0
const QAbstractItemView__ScrollPerPixel QAbstractItemView__ScrollMode = 1

type QAbstractItemView__DragDropMode = int

const QAbstractItemView__NoDragDrop QAbstractItemView__DragDropMode = 0
const QAbstractItemView__DragOnly QAbstractItemView__DragDropMode = 1
const QAbstractItemView__DropOnly QAbstractItemView__DragDropMode = 2
const QAbstractItemView__DragDrop QAbstractItemView__DragDropMode = 3
const QAbstractItemView__InternalMove QAbstractItemView__DragDropMode = 4

type QAbstractItemView__CursorAction = int

const QAbstractItemView__MoveUp QAbstractItemView__CursorAction = 0
const QAbstractItemView__MoveDown QAbstractItemView__CursorAction = 1
const QAbstractItemView__MoveLeft QAbstractItemView__CursorAction = 2
const QAbstractItemView__MoveRight QAbstractItemView__CursorAction = 3
const QAbstractItemView__MoveHome QAbstractItemView__CursorAction = 4
const QAbstractItemView__MoveEnd QAbstractItemView__CursorAction = 5
const QAbstractItemView__MovePageUp QAbstractItemView__CursorAction = 6
const QAbstractItemView__MovePageDown QAbstractItemView__CursorAction = 7
const QAbstractItemView__MoveNext QAbstractItemView__CursorAction = 8
const QAbstractItemView__MovePrevious QAbstractItemView__CursorAction = 9

type QAbstractItemView__State = int

const QAbstractItemView__NoState QAbstractItemView__State = 0
const QAbstractItemView__DraggingState QAbstractItemView__State = 1
const QAbstractItemView__DragSelectingState QAbstractItemView__State = 2
const QAbstractItemView__EditingState QAbstractItemView__State = 3
const QAbstractItemView__ExpandingState QAbstractItemView__State = 4
const QAbstractItemView__CollapsingState QAbstractItemView__State = 5
const QAbstractItemView__AnimatingState QAbstractItemView__State = 6

type QAbstractItemView__DropIndicatorPosition = int

const QAbstractItemView__OnItem QAbstractItemView__DropIndicatorPosition = 0
const QAbstractItemView__AboveItem QAbstractItemView__DropIndicatorPosition = 1
const QAbstractItemView__BelowItem QAbstractItemView__DropIndicatorPosition = 2
const QAbstractItemView__OnViewport QAbstractItemView__DropIndicatorPosition = 3

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
