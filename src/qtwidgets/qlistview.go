//  header block begin
// /usr/include/qt/QtWidgets/qlistview.h
// #include <qlistview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qlistview.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QListView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:81
// index:0
// void QListView(class QWidget *)
func NewQListView(parent unsafe.Pointer) *QListView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QListView{cthis}
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
	// 0: (, QListView::Movement movement), (&movement)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setMovementENS_8MovementE", ffiqt.FFI_TYPE_VOID, this.cthis, &movement)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:85
// index:0
// QListView::Movement movement()
func (this *QListView) Movement() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8movementEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:87
// index:0
// void setFlow(enum QListView::Flow)
func (this *QListView) SetFlow(flow int) {
	// 0: (, QListView::Flow flow), (&flow)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView7setFlowENS_4FlowE", ffiqt.FFI_TYPE_VOID, this.cthis, &flow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:88
// index:0
// QListView::Flow flow()
func (this *QListView) Flow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView4flowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:90
// index:0
// void setWrapping(_Bool)
func (this *QListView) SetWrapping(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setWrappingEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:91
// index:0
// bool isWrapping()
func (this *QListView) IsWrapping() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10isWrappingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:93
// index:0
// void setResizeMode(enum QListView::ResizeMode)
func (this *QListView) SetResizeMode(mode int) {
	// 0: (, QListView::ResizeMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13setResizeModeENS_10ResizeModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:94
// index:0
// QListView::ResizeMode resizeMode()
func (this *QListView) ResizeMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10resizeModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:96
// index:0
// void setLayoutMode(enum QListView::LayoutMode)
func (this *QListView) SetLayoutMode(mode int) {
	// 0: (, QListView::LayoutMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13setLayoutModeENS_10LayoutModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:97
// index:0
// QListView::LayoutMode layoutMode()
func (this *QListView) LayoutMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10layoutModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:99
// index:0
// void setSpacing(int)
func (this *QListView) SetSpacing(space int) {
	// 0: (, int space), (&space)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView10setSpacingEi", ffiqt.FFI_TYPE_VOID, this.cthis, &space)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:100
// index:0
// int spacing()
func (this *QListView) Spacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView7spacingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:102
// index:0
// void setBatchSize(int)
func (this *QListView) SetBatchSize(batchSize int) {
	// 0: (, int batchSize), (&batchSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setBatchSizeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &batchSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:103
// index:0
// int batchSize()
func (this *QListView) BatchSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView9batchSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:105
// index:0
// void setGridSize(const class QSize &)
func (this *QListView) SetGridSize(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setGridSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:106
// index:0
// QSize gridSize()
func (this *QListView) GridSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8gridSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:108
// index:0
// void setViewMode(enum QListView::ViewMode)
func (this *QListView) SetViewMode(mode int) {
	// 0: (, QListView::ViewMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setViewModeENS_8ViewModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:109
// index:0
// QListView::ViewMode viewMode()
func (this *QListView) ViewMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8viewModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:111
// index:0
// void clearPropertyFlags()
func (this *QListView) ClearPropertyFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView18clearPropertyFlagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:113
// index:0
// bool isRowHidden(int)
func (this *QListView) IsRowHidden(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView11isRowHiddenEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:114
// index:0
// void setRowHidden(int, _Bool)
func (this *QListView) SetRowHidden(row int, hide bool) {
	// 0: (, int row, bool hide), (&row, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setRowHiddenEib", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:116
// index:0
// void setModelColumn(int)
func (this *QListView) SetModelColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14setModelColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:117
// index:0
// int modelColumn()
func (this *QListView) ModelColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView11modelColumnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:119
// index:0
// void setUniformItemSizes(_Bool)
func (this *QListView) SetUniformItemSizes(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView19setUniformItemSizesEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:120
// index:0
// bool uniformItemSizes()
func (this *QListView) UniformItemSizes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView16uniformItemSizesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:122
// index:0
// void setWordWrap(_Bool)
func (this *QListView) SetWordWrap(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setWordWrapEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:123
// index:0
// bool wordWrap()
func (this *QListView) WordWrap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8wordWrapEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:125
// index:0
// void setSelectionRectVisible(_Bool)
func (this *QListView) SetSelectionRectVisible(show bool) {
	// 0: (, bool show), (&show)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView23setSelectionRectVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:126
// index:0
// bool isSelectionRectVisible()
func (this *QListView) IsSelectionRectVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView22isSelectionRectVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:128
// index:0
// virtual
// QRect visualRect(const class QModelIndex &)
func (this *QListView) VisualRect(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:129
// index:0
// virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QListView) ScrollTo(index unsafe.Pointer, hint int) {
	// 0: (, const QModelIndex & index, QAbstractItemView::ScrollHint hint), (index, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.cthis, index, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:130
// index:0
// virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QListView) IndexAt(p unsafe.Pointer) {
	// 0: (, const QPoint & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView7indexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:132
// index:0
// virtual
// void doItemsLayout()
func (this *QListView) DoItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13doItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:133
// index:0
// virtual
// void reset()
func (this *QListView) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView5resetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:134
// index:0
// virtual
// void setRootIndex(const class QModelIndex &)
func (this *QListView) SetRootIndex(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

//  body block end
