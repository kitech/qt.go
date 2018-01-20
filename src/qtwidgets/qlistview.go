//  header block begin
// /usr/include/qt/QtWidgets/qlistview.h
// #include <qlistview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
type QListView struct {
	*QAbstractItemView
}

func (this *QListView) GetCthis() unsafe.Pointer {
	return this.QAbstractItemView.GetCthis()
}

// /usr/include/qt/QtWidgets/qlistview.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QListView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:81
// index:0
// void QListView(class QWidget *)
func NewQListView(parent unsafe.Pointer) *QListView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQListViewFromPointer(cthis)
	return gothis
}
func NewQListViewFromPointer(cthis unsafe.Pointer) *QListView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QListView{bcthis0}
}

// /usr/include/qt/QtWidgets/qlistview.h:140
// index:1
// void QListView(class QListViewPrivate &, class QWidget *)
func NewQListView_1(arg0 unsafe.Pointer, parent unsafe.Pointer) *QListView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListViewC2ER16QListViewPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQListViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistview.h:82
// index:0
// virtual
// void ~QListView()
func DeleteQListView(*QListView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:84
// index:0
// void setMovement(enum QListView::Movement)
func (this *QListView) SetMovement(movement int) {
	// 0: (, movement QListView::Movement), (&movement)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setMovementENS_8MovementE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &movement)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:85
// index:0
// QListView::Movement movement()
func (this *QListView) Movement() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8movementEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:87
// index:0
// void setFlow(enum QListView::Flow)
func (this *QListView) SetFlow(flow int) {
	// 0: (, flow QListView::Flow), (&flow)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView7setFlowENS_4FlowE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:88
// index:0
// QListView::Flow flow()
func (this *QListView) Flow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView4flowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:90
// index:0
// void setWrapping(_Bool)
func (this *QListView) SetWrapping(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setWrappingEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:91
// index:0
// bool isWrapping()
func (this *QListView) IsWrapping() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10isWrappingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:93
// index:0
// void setResizeMode(enum QListView::ResizeMode)
func (this *QListView) SetResizeMode(mode int) {
	// 0: (, mode QListView::ResizeMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13setResizeModeENS_10ResizeModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:94
// index:0
// QListView::ResizeMode resizeMode()
func (this *QListView) ResizeMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10resizeModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:96
// index:0
// void setLayoutMode(enum QListView::LayoutMode)
func (this *QListView) SetLayoutMode(mode int) {
	// 0: (, mode QListView::LayoutMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13setLayoutModeENS_10LayoutModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:97
// index:0
// QListView::LayoutMode layoutMode()
func (this *QListView) LayoutMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10layoutModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:99
// index:0
// void setSpacing(int)
func (this *QListView) SetSpacing(space int) {
	// 0: (, space int), (&space)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView10setSpacingEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &space)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:100
// index:0
// int spacing()
func (this *QListView) Spacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView7spacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:102
// index:0
// void setBatchSize(int)
func (this *QListView) SetBatchSize(batchSize int) {
	// 0: (, batchSize int), (&batchSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setBatchSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &batchSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:103
// index:0
// int batchSize()
func (this *QListView) BatchSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView9batchSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:105
// index:0
// void setGridSize(const class QSize &)
func (this *QListView) SetGridSize(size unsafe.Pointer) {
	// 0: (, size const QSize &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setGridSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:106
// index:0
// QSize gridSize()
func (this *QListView) GridSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8gridSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:108
// index:0
// void setViewMode(enum QListView::ViewMode)
func (this *QListView) SetViewMode(mode int) {
	// 0: (, mode QListView::ViewMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setViewModeENS_8ViewModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:109
// index:0
// QListView::ViewMode viewMode()
func (this *QListView) ViewMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8viewModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:111
// index:0
// void clearPropertyFlags()
func (this *QListView) ClearPropertyFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView18clearPropertyFlagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:113
// index:0
// bool isRowHidden(int)
func (this *QListView) IsRowHidden(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView11isRowHiddenEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:114
// index:0
// void setRowHidden(int, _Bool)
func (this *QListView) SetRowHidden(row int, hide bool) {
	// 0: (, row int, hide bool), (&row, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setRowHiddenEib", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:116
// index:0
// void setModelColumn(int)
func (this *QListView) SetModelColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14setModelColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:117
// index:0
// int modelColumn()
func (this *QListView) ModelColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView11modelColumnEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:119
// index:0
// void setUniformItemSizes(_Bool)
func (this *QListView) SetUniformItemSizes(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView19setUniformItemSizesEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:120
// index:0
// bool uniformItemSizes()
func (this *QListView) UniformItemSizes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView16uniformItemSizesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:122
// index:0
// void setWordWrap(_Bool)
func (this *QListView) SetWordWrap(on bool) {
	// 0: (, on bool), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setWordWrapEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:123
// index:0
// bool wordWrap()
func (this *QListView) WordWrap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8wordWrapEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:125
// index:0
// void setSelectionRectVisible(_Bool)
func (this *QListView) SetSelectionRectVisible(show bool) {
	// 0: (, show bool), (&show)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView23setSelectionRectVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:126
// index:0
// bool isSelectionRectVisible()
func (this *QListView) IsSelectionRectVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView22isSelectionRectVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:128
// index:0
// virtual
// QRect visualRect(const class QModelIndex &)
func (this *QListView) VisualRect(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:129
// index:0
// virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QListView) ScrollTo(index unsafe.Pointer, hint int) {
	// 0: (, index const QModelIndex &, hint QAbstractItemView::ScrollHint), (index, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:130
// index:0
// virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QListView) IndexAt(p unsafe.Pointer) {
	// 0: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView7indexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:132
// index:0
// virtual
// void doItemsLayout()
func (this *QListView) DoItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13doItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:133
// index:0
// virtual
// void reset()
func (this *QListView) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView5resetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:134
// index:0
// virtual
// void setRootIndex(const class QModelIndex &)
func (this *QListView) SetRootIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:142
// index:0
// virtual
// bool event(class QEvent *)
func (this *QListView) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:144
// index:0
// virtual
// void scrollContentsBy(int, int)
func (this *QListView) ScrollContentsBy(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView16scrollContentsByEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:146
// index:0
// void resizeContents(int, int)
func (this *QListView) ResizeContents(width int, height int) {
	// 0: (, width int, height int), (&width, &height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14resizeContentsEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width, &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:147
// index:0
// QSize contentsSize()
func (this *QListView) ContentsSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView12contentsSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:150
// index:0
// virtual
// void rowsInserted(const class QModelIndex &, int, int)
func (this *QListView) RowsInserted(parent unsafe.Pointer, start int, end int) {
	// 0: (, parent const QModelIndex &, start int, end int), (parent, &start, &end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:151
// index:0
// virtual
// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QListView) RowsAboutToBeRemoved(parent unsafe.Pointer, start int, end int) {
	// 0: (, parent const QModelIndex &, start int, end int), (parent, &start, &end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView20rowsAboutToBeRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:153
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QListView) MouseMoveEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:154
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QListView) MouseReleaseEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:156
// index:0
// virtual
// void wheelEvent(class QWheelEvent *)
func (this *QListView) WheelEvent(e unsafe.Pointer) {
	// 0: (, e QWheelEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:159
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QListView) TimerEvent(e unsafe.Pointer) {
	// 0: (, e QTimerEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:160
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QListView) ResizeEvent(e unsafe.Pointer) {
	// 0: (, e QResizeEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:162
// index:0
// virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QListView) DragMoveEvent(e unsafe.Pointer) {
	// 0: (, e QDragMoveEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:163
// index:0
// virtual
// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QListView) DragLeaveEvent(e unsafe.Pointer) {
	// 0: (, e QDragLeaveEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:164
// index:0
// virtual
// void dropEvent(class QDropEvent *)
func (this *QListView) DropEvent(e unsafe.Pointer) {
	// 0: (, e QDropEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:165
// index:0
// virtual
// void startDrag(Qt::DropActions)
func (this *QListView) StartDrag(supportedActions int) {
	// 0: (, QFlags<Qt::DropAction> supportedActions), (&supportedActions)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView9startDragE6QFlagsIN2Qt10DropActionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &supportedActions)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:168
// index:0
// virtual
// QStyleOptionViewItem viewOptions()
func (this *QListView) ViewOptions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView11viewOptionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:169
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QListView) PaintEvent(e unsafe.Pointer) {
	// 0: (, e QPaintEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:171
// index:0
// virtual
// int horizontalOffset()
func (this *QListView) HorizontalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView16horizontalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:172
// index:0
// virtual
// int verticalOffset()
func (this *QListView) VerticalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView14verticalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:173
// index:0
// virtual
// QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QListView) MoveCursor(cursorAction int, modifiers int) {
	// 0: (, cursorAction QAbstractItemView::CursorAction, QFlags<Qt::KeyboardModifier> modifiers), (&cursorAction, &modifiers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cursorAction, &modifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:174
// index:0
// QRect rectForIndex(const class QModelIndex &)
func (this *QListView) RectForIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView12rectForIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:175
// index:0
// void setPositionForIndex(const class QPoint &, const class QModelIndex &)
func (this *QListView) SetPositionForIndex(position unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, position const QPoint &, index const QModelIndex &), (position, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView19setPositionForIndexERK6QPointRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), position, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:177
// index:0
// virtual
// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QListView) SetSelection(rect unsafe.Pointer, command int) {
	// 0: (, rect const QRect &, QFlags<QItemSelectionModel::SelectionFlag> command), (rect, &command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:178
// index:0
// virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QListView) VisualRegionForSelection(selection unsafe.Pointer) {
	// 0: (, selection const QItemSelection &), (selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:179
// index:0
// virtual
// QModelIndexList selectedIndexes()
func (this *QListView) SelectedIndexes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView15selectedIndexesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:181
// index:0
// virtual
// void updateGeometries()
func (this *QListView) UpdateGeometries() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView16updateGeometriesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:183
// index:0
// virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QListView) IsIndexHidden(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:185
// index:0
// virtual
// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QListView) SelectionChanged(selected unsafe.Pointer, deselected unsafe.Pointer) {
	// 0: (, selected const QItemSelection &, deselected const QItemSelection &), (selected, deselected)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selected, deselected)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:186
// index:0
// virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QListView) CurrentChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, current const QModelIndex &, previous const QModelIndex &), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:188
// index:0
// virtual
// QSize viewportSizeHint()
func (this *QListView) ViewportSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView16viewportSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
