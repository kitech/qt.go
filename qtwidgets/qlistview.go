package qtwidgets

// /usr/include/qt/QtWidgets/qlistview.h
// #include <qlistview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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

// bool event(class QEvent *)
func (this *QListView) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void scrollContentsBy(int, int)
func (this *QListView) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// void resizeContents(int, int)
func (this *QListView) InheritResizeContents(f func(width int, height int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeContents", f)
}

// QSize contentsSize()
func (this *QListView) InheritContentsSize(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "contentsSize", f)
}

// void rowsInserted(const class QModelIndex &, int, int)
func (this *QListView) InheritRowsInserted(f func(parent *qtcore.QModelIndex, start int, end int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsInserted", f)
}

// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QListView) InheritRowsAboutToBeRemoved(f func(parent *qtcore.QModelIndex, start int, end int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsAboutToBeRemoved", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QListView) InheritMouseMoveEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QListView) InheritMouseReleaseEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QListView) InheritWheelEvent(f func(e *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QListView) InheritTimerEvent(f func(e *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QListView) InheritResizeEvent(f func(e *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void dragMoveEvent(class QDragMoveEvent *)
func (this *QListView) InheritDragMoveEvent(f func(e *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QListView) InheritDragLeaveEvent(f func(e *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(class QDropEvent *)
func (this *QListView) InheritDropEvent(f func(e *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void startDrag(Qt::DropActions)
func (this *QListView) InheritStartDrag(f func(supportedActions int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "startDrag", f)
}

// QStyleOptionViewItem viewOptions()
func (this *QListView) InheritViewOptions(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewOptions", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QListView) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// int horizontalOffset()
func (this *QListView) InheritHorizontalOffset(f func() int) {
	qtrt.SetAllInheritCallback(this, "horizontalOffset", f)
}

// int verticalOffset()
func (this *QListView) InheritVerticalOffset(f func() int) {
	qtrt.SetAllInheritCallback(this, "verticalOffset", f)
}

// QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QListView) InheritMoveCursor(f func(cursorAction int, modifiers int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "moveCursor", f)
}

// QRect rectForIndex(const class QModelIndex &)
func (this *QListView) InheritRectForIndex(f func(index *qtcore.QModelIndex) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "rectForIndex", f)
}

// void setPositionForIndex(const class QPoint &, const class QModelIndex &)
func (this *QListView) InheritSetPositionForIndex(f func(position *qtcore.QPoint, index *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setPositionForIndex", f)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QListView) InheritSetSelection(f func(rect *qtcore.QRect, command int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSelection", f)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QListView) InheritVisualRegionForSelection(f func(selection *qtcore.QItemSelection) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "visualRegionForSelection", f)
}

// QModelIndexList selectedIndexes()
func (this *QListView) InheritSelectedIndexes(f func() *qtcore.QModelIndexList /*9999*/) {
	qtrt.SetAllInheritCallback(this, "selectedIndexes", f)
}

// void updateGeometries()
func (this *QListView) InheritUpdateGeometries(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateGeometries", f)
}

// bool isIndexHidden(const class QModelIndex &)
func (this *QListView) InheritIsIndexHidden(f func(index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "isIndexHidden", f)
}

// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QListView) InheritSelectionChanged(f func(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) /*void*/) {
	qtrt.SetAllInheritCallback(this, "selectionChanged", f)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QListView) InheritCurrentChanged(f func(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "currentChanged", f)
}

// QSize viewportSizeHint()
func (this *QListView) InheritViewportSizeHint(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewportSizeHint", f)
}

/*

 */
type QListView struct {
	*QAbstractItemView
}
type QListView_ITF interface {
	QAbstractItemView_ITF
	QListView_PTR() *QListView
}

func (ptr *QListView) QListView_PTR() *QListView { return ptr }

func (this *QListView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemView.GetCthis()
	}
}
func (this *QListView) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemView = NewQAbstractItemViewFromPointer(cthis)
}
func NewQListViewFromPointer(cthis unsafe.Pointer) *QListView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QListView{bcthis0}
}
func (*QListView) NewFromPointer(cthis unsafe.Pointer) *QListView {
	return NewQListViewFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlistview.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QListView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlistview.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListView(QWidget *)

/*
Creates a new QListView with the given parent to view a model. Use setModel() to set the model.
*/
func NewQListView(parent QWidget_ITF /*777 QWidget **/) *QListView {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QListView")
	return gothis
}

// /usr/include/qt/QtWidgets/qlistview.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListView(QWidget *)

/*
Creates a new QListView with the given parent to view a model. Use setModel() to set the model.
*/
func NewQListView__() *QListView {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQListViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QListView")
	return gothis
}

// /usr/include/qt/QtWidgets/qlistview.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QListView()

/*

 */
func DeleteQListView(this *QListView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlistview.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovement(enum QListView::Movement)

/*

 */
func (this *QListView) SetMovement(movement int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView11setMovementENS_8MovementE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), movement)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] QListView::Movement movement() const

/*

 */
func (this *QListView) Movement() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView8movementEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlow(enum QListView::Flow)

/*

 */
func (this *QListView) SetFlow(flow int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView7setFlowENS_4FlowE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flow)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] QListView::Flow flow() const

/*

 */
func (this *QListView) Flow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView4flowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWrapping(_Bool)

/*

 */
func (this *QListView) SetWrapping(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView11setWrappingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWrapping() const

/*

 */
func (this *QListView) IsWrapping() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView10isWrappingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResizeMode(enum QListView::ResizeMode)

/*

 */
func (this *QListView) SetResizeMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView13setResizeModeENS_10ResizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QListView::ResizeMode resizeMode() const

/*

 */
func (this *QListView) ResizeMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView10resizeModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayoutMode(enum QListView::LayoutMode)

/*

 */
func (this *QListView) SetLayoutMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView13setLayoutModeENS_10LayoutModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QListView::LayoutMode layoutMode() const

/*

 */
func (this *QListView) LayoutMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView10layoutModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(int)

/*

 */
func (this *QListView) SetSpacing(space int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView10setSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), space)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int spacing() const

/*

 */
func (this *QListView) Spacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView7spacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistview.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBatchSize(int)

/*

 */
func (this *QListView) SetBatchSize(batchSize int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView12setBatchSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), batchSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int batchSize() const

/*

 */
func (this *QListView) BatchSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView9batchSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistview.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGridSize(const QSize &)

/*

 */
func (this *QListView) SetGridSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView11setGridSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize gridSize() const

/*

 */
func (this *QListView) GridSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView8gridSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewMode(enum QListView::ViewMode)

/*

 */
func (this *QListView) SetViewMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView11setViewModeENS_8ViewModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QListView::ViewMode viewMode() const

/*

 */
func (this *QListView) ViewMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView8viewModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearPropertyFlags()

/*
Clears the QListView-specific property flags. See viewMode.

Properties inherited from QAbstractItemView are not covered by the property flags. Specifically, dragEnabled and acceptsDrops are computed by QListView when calling setMovement() or setViewMode().
*/
func (this *QListView) ClearPropertyFlags() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView18clearPropertyFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRowHidden(int) const

/*
Returns true if the row is hidden; otherwise returns false.
*/
func (this *QListView) IsRowHidden(row int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView11isRowHiddenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowHidden(int, _Bool)

/*
If hide is true, the given row will be hidden; otherwise the row will be shown.

See also isRowHidden().
*/
func (this *QListView) SetRowHidden(row int, hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView12setRowHiddenEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModelColumn(int)

/*

 */
func (this *QListView) SetModelColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView14setModelColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int modelColumn() const

/*

 */
func (this *QListView) ModelColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView11modelColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistview.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUniformItemSizes(_Bool)

/*

 */
func (this *QListView) SetUniformItemSizes(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView19setUniformItemSizesEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool uniformItemSizes() const

/*

 */
func (this *QListView) UniformItemSizes() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView16uniformItemSizesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordWrap(_Bool)

/*

 */
func (this *QListView) SetWordWrap(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView11setWordWrapEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wordWrap() const

/*

 */
func (this *QListView) WordWrap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView8wordWrapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionRectVisible(_Bool)

/*

 */
func (this *QListView) SetSelectionRectVisible(show bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView23setSelectionRectVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), show)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:126
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSelectionRectVisible() const

/*

 */
func (this *QListView) IsSelectionRectVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView22isSelectionRectVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:128
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect visualRect(const QModelIndex &) const

/*
Reimplemented from QAbstractItemView::visualRect().
*/
func (this *QListView) VisualRect(index qtcore.QModelIndex_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView10visualRectERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:129
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, enum QAbstractItemView::ScrollHint)

/*
Reimplemented from QAbstractItemView::scrollTo().
*/
func (this *QListView) ScrollTo(index qtcore.QModelIndex_ITF, hint int) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:129
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, enum QAbstractItemView::ScrollHint)

/*
Reimplemented from QAbstractItemView::scrollTo().
*/
func (this *QListView) ScrollTo__(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, QAbstractItemView::ScrollHint=Enum, QAbstractItemView::ScrollHint=Enum,
	hint := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:130
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex indexAt(const QPoint &) const

/*
Reimplemented from QAbstractItemView::indexAt().
*/
func (this *QListView) IndexAt(p qtcore.QPoint_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView7indexAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:132
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doItemsLayout()

/*

 */
func (this *QListView) DoItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView13doItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:133
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reset()

/*

 */
func (this *QListView) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:134
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRootIndex(const QModelIndex &)

/*

 */
func (this *QListView) SetRootIndex(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView12setRootIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QListView) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:144
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)

/*

 */
func (this *QListView) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:146
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void resizeContents(int, int)

/*

 */
func (this *QListView) ResizeContents(width int, height int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView14resizeContentsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:147
// index:0
// Protected Visibility=Default Availability=Available
// [8] QSize contentsSize() const

/*

 */
func (this *QListView) ContentsSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView12contentsSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:150
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsInserted(const QModelIndex &, int, int)

/*
Reimplemented from QAbstractItemView::rowsInserted().
*/
func (this *QListView) RowsInserted(parent qtcore.QModelIndex_ITF, start int, end int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView12rowsInsertedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:151
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsAboutToBeRemoved(const QModelIndex &, int, int)

/*
Reimplemented from QAbstractItemView::rowsAboutToBeRemoved().
*/
func (this *QListView) RowsAboutToBeRemoved(parent qtcore.QModelIndex_ITF, start int, end int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView20rowsAboutToBeRemovedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:153
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QListView) MouseMoveEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:154
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QListView) MouseReleaseEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:156
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QListView) WheelEvent(e qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QWheelEvent_PTR() != nil {
		convArg0 = e.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:159
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QListView) TimerEvent(e qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QTimerEvent_PTR() != nil {
		convArg0 = e.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:160
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QListView) ResizeEvent(e qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QResizeEvent_PTR() != nil {
		convArg0 = e.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:162
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QWidget::dragMoveEvent().
*/
func (this *QListView) DragMoveEvent(e qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragMoveEvent_PTR() != nil {
		convArg0 = e.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:163
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
Reimplemented from QWidget::dragLeaveEvent().
*/
func (this *QListView) DragLeaveEvent(e qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragLeaveEvent_PTR() != nil {
		convArg0 = e.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:164
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QWidget::dropEvent().
*/
func (this *QListView) DropEvent(e qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDropEvent_PTR() != nil {
		convArg0 = e.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:165
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void startDrag(Qt::DropActions)

/*
Reimplemented from QAbstractItemView::startDrag().
*/
func (this *QListView) StartDrag(supportedActions int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView9startDragE6QFlagsIN2Qt10DropActionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), supportedActions)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:168
// index:0
// Protected virtual Visibility=Default Availability=Available
// [192] QStyleOptionViewItem viewOptions() const

/*
Reimplemented from QAbstractItemView::viewOptions().
*/
func (this *QListView) ViewOptions() *QStyleOptionViewItem /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView11viewOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStyleOptionViewItem)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:169
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QListView) PaintEvent(e qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QPaintEvent_PTR() != nil {
		convArg0 = e.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:171
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int horizontalOffset() const

/*
Reimplemented from QAbstractItemView::horizontalOffset().
*/
func (this *QListView) HorizontalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView16horizontalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistview.h:172
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int verticalOffset() const

/*
Reimplemented from QAbstractItemView::verticalOffset().
*/
func (this *QListView) VerticalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView14verticalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistview.h:173
// index:0
// Protected virtual Visibility=Default Availability=Available
// [24] QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)

/*
Reimplemented from QAbstractItemView::moveCursor().
*/
func (this *QListView) MoveCursor(cursorAction int, modifiers int) *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cursorAction, modifiers)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:174
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect rectForIndex(const QModelIndex &) const

/*
Returns the rectangle of the item at position index in the model. The rectangle is in contents coordinates.

See also visualRect().
*/
func (this *QListView) RectForIndex(index qtcore.QModelIndex_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView12rectForIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:175
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setPositionForIndex(const QPoint &, const QModelIndex &)

/*
Sets the contents position of the item at index in the model to the given position. If the list view's movement mode is Static or its view mode is ListView, this function will have no effect.

This function was introduced in  Qt 4.1.
*/
func (this *QListView) SetPositionForIndex(position qtcore.QPoint_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if position != nil && position.QPoint_PTR() != nil {
		convArg0 = position.QPoint_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView19setPositionForIndexERK6QPointRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:177
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)

/*
Reimplemented from QAbstractItemView::setSelection().
*/
func (this *QListView) SetSelection(rect qtcore.QRect_ITF, command int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:178
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QRegion visualRegionForSelection(const QItemSelection &) const

/*
Reimplemented from QAbstractItemView::visualRegionForSelection().

Since 4.7, the returned region only contains rectangles intersecting (or included in) the viewport.
*/
func (this *QListView) VisualRegionForSelection(selection qtcore.QItemSelection_ITF) *qtgui.QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if selection != nil && selection.QItemSelection_PTR() != nil {
		convArg0 = selection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView24visualRegionForSelectionERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:179
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QModelIndexList selectedIndexes() const

/*
Reimplemented from QAbstractItemView::selectedIndexes().
*/
func (this *QListView) SelectedIndexes() *qtcore.QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView15selectedIndexesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:181
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometries()

/*
Reimplemented from QAbstractItemView::updateGeometries().
*/
func (this *QListView) UpdateGeometries() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView16updateGeometriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool isIndexHidden(const QModelIndex &) const

/*
Reimplemented from QAbstractItemView::isIndexHidden().
*/
func (this *QListView) IsIndexHidden(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView13isIndexHiddenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:185
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void selectionChanged(const QItemSelection &, const QItemSelection &)

/*
Reimplemented from QAbstractItemView::selectionChanged().
*/
func (this *QListView) SelectionChanged(selected qtcore.QItemSelection_ITF, deselected qtcore.QItemSelection_ITF) {
	var convArg0 unsafe.Pointer
	if selected != nil && selected.QItemSelection_PTR() != nil {
		convArg0 = selected.QItemSelection_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if deselected != nil && deselected.QItemSelection_PTR() != nil {
		convArg1 = deselected.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView16selectionChangedERK14QItemSelectionS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:186
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)

/*
Reimplemented from QAbstractItemView::currentChanged().
*/
func (this *QListView) CurrentChanged(current qtcore.QModelIndex_ITF, previous qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QModelIndex_PTR() != nil {
		convArg0 = current.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QModelIndex_PTR() != nil {
		convArg1 = previous.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QListView14currentChangedERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:188
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint() const

/*
Reimplemented from QAbstractScrollArea::viewportSizeHint().

This function was introduced in  Qt 5.2.
*/
func (this *QListView) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QListView16viewportSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

/*

 */
type QListView__Movement = int

// The items cannot be moved by the user.
const QListView__Static QListView__Movement = 0

// The items can be moved freely by the user.
const QListView__Free QListView__Movement = 1

// The items snap to the specified grid when moved; see setGridSize().
const QListView__Snap QListView__Movement = 2

/*

 */
type QListView__Flow = int

// The items are laid out in the view from the left to the right.
const QListView__LeftToRight QListView__Flow = 0

// The items are laid out in the view from the top to the bottom.
const QListView__TopToBottom QListView__Flow = 1

/*

 */
type QListView__ResizeMode = int

// The items will only be laid out the first time the view is shown.
const QListView__Fixed QListView__ResizeMode = 0

// The items will be laid out every time the view is resized.
const QListView__Adjust QListView__ResizeMode = 1

/*


See also batchSize.

*/
type QListView__LayoutMode = int

// The items are laid out all at once.
const QListView__SinglePass QListView__LayoutMode = 0

// The items are laid out in batches of batchSize items.
const QListView__Batched QListView__LayoutMode = 1

/*

 */
type QListView__ViewMode = int

// The items are laid out using TopToBottom flow, with Small size and Static movement
const QListView__ListMode QListView__ViewMode = 0

// The items are laid out using LeftToRight flow, with Large size and Free movement
const QListView__IconMode QListView__ViewMode = 1

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
