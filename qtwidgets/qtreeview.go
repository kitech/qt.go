package qtwidgets

// /usr/include/qt/QtWidgets/qtreeview.h
// #include <qtreeview.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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

// void columnResized(int, int, int)
func (this *QTreeView) InheritColumnResized(f func(column int, oldSize int, newSize int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "columnResized", f)
}

// void columnCountChanged(int, int)
func (this *QTreeView) InheritColumnCountChanged(f func(oldCount int, newCount int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "columnCountChanged", f)
}

// void columnMoved()
func (this *QTreeView) InheritColumnMoved(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "columnMoved", f)
}

// void reexpand()
func (this *QTreeView) InheritReexpand(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "reexpand", f)
}

// void rowsRemoved(const class QModelIndex &, int, int)
func (this *QTreeView) InheritRowsRemoved(f func(parent *qtcore.QModelIndex, first int, last int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsRemoved", f)
}

// void scrollContentsBy(int, int)
func (this *QTreeView) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// void rowsInserted(const class QModelIndex &, int, int)
func (this *QTreeView) InheritRowsInserted(f func(parent *qtcore.QModelIndex, start int, end int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsInserted", f)
}

// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QTreeView) InheritRowsAboutToBeRemoved(f func(parent *qtcore.QModelIndex, start int, end int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsAboutToBeRemoved", f)
}

// QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QTreeView) InheritMoveCursor(f func(cursorAction int, modifiers int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "moveCursor", f)
}

// int horizontalOffset()
func (this *QTreeView) InheritHorizontalOffset(f func() int) {
	qtrt.SetAllInheritCallback(this, "horizontalOffset", f)
}

// int verticalOffset()
func (this *QTreeView) InheritVerticalOffset(f func() int) {
	qtrt.SetAllInheritCallback(this, "verticalOffset", f)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QTreeView) InheritSetSelection(f func(rect *qtcore.QRect, command int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSelection", f)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QTreeView) InheritVisualRegionForSelection(f func(selection *qtcore.QItemSelection) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "visualRegionForSelection", f)
}

// QModelIndexList selectedIndexes()
func (this *QTreeView) InheritSelectedIndexes(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "selectedIndexes", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QTreeView) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QTreeView) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void drawTree(class QPainter *, const class QRegion &)
func (this *QTreeView) InheritDrawTree(f func(painter *qtgui.QPainter /*777 QPainter **/, region *qtgui.QRegion) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawTree", f)
}

// void drawRow(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QTreeView) InheritDrawRow(f func(painter *qtgui.QPainter /*777 QPainter **/, options *QStyleOptionViewItem, index *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawRow", f)
}

// void drawBranches(class QPainter *, const class QRect &, const class QModelIndex &)
func (this *QTreeView) InheritDrawBranches(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, index *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawBranches", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QTreeView) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QTreeView) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QTreeView) InheritMouseDoubleClickEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QTreeView) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QTreeView) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void dragMoveEvent(class QDragMoveEvent *)
func (this *QTreeView) InheritDragMoveEvent(f func(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// bool viewportEvent(class QEvent *)
func (this *QTreeView) InheritViewportEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "viewportEvent", f)
}

// void updateGeometries()
func (this *QTreeView) InheritUpdateGeometries(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateGeometries", f)
}

// QSize viewportSizeHint()
func (this *QTreeView) InheritViewportSizeHint(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "viewportSizeHint", f)
}

// int sizeHintForColumn(int)
func (this *QTreeView) InheritSizeHintForColumn(f func(column int) int) {
	qtrt.SetAllInheritCallback(this, "sizeHintForColumn", f)
}

// int indexRowSizeHint(const class QModelIndex &)
func (this *QTreeView) InheritIndexRowSizeHint(f func(index *qtcore.QModelIndex) int) {
	qtrt.SetAllInheritCallback(this, "indexRowSizeHint", f)
}

// int rowHeight(const class QModelIndex &)
func (this *QTreeView) InheritRowHeight(f func(index *qtcore.QModelIndex) int) {
	qtrt.SetAllInheritCallback(this, "rowHeight", f)
}

// void horizontalScrollbarAction(int)
func (this *QTreeView) InheritHorizontalScrollbarAction(f func(action int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "horizontalScrollbarAction", f)
}

// bool isIndexHidden(const class QModelIndex &)
func (this *QTreeView) InheritIsIndexHidden(f func(index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "isIndexHidden", f)
}

// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QTreeView) InheritSelectionChanged(f func(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) /*void*/) {
	qtrt.SetAllInheritCallback(this, "selectionChanged", f)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QTreeView) InheritCurrentChanged(f func(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "currentChanged", f)
}

type QTreeView struct {
	*QAbstractItemView
}

func (this *QTreeView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemView.GetCthis()
	}
}
func (this *QTreeView) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemView = NewQAbstractItemViewFromPointer(cthis)
}
func NewQTreeViewFromPointer(cthis unsafe.Pointer) *QTreeView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QTreeView{bcthis0}
}
func (*QTreeView) NewFromPointer(cthis unsafe.Pointer) *QTreeView {
	return NewQTreeViewFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtreeview.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTreeView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreeview.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeView(QWidget *)
func NewQTreeView(parent *QWidget /*777 QWidget **/) *QTreeView {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qtreeview.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTreeView()
func DeleteQTreeView(this *QTreeView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtreeview.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)
func (this *QTreeView) SetModel(model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/) {
	var convArg0 = model.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRootIndex(const QModelIndex &)
func (this *QTreeView) SetRootIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12setRootIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSelectionModel(QItemSelectionModel *)
func (this *QTreeView) SetSelectionModel(selectionModel *qtcore.QItemSelectionModel /*777 QItemSelectionModel **/) {
	var convArg0 = selectionModel.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView17setSelectionModelEP19QItemSelectionModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QHeaderView * header()
func (this *QTreeView) Header() *QHeaderView /*777 QHeaderView **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView6headerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQHeaderViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreeview.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeader(QHeaderView *)
func (this *QTreeView) SetHeader(header *QHeaderView /*777 QHeaderView **/) {
	var convArg0 = header.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView9setHeaderEP11QHeaderView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int autoExpandDelay()
func (this *QTreeView) AutoExpandDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView15autoExpandDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoExpandDelay(int)
func (this *QTreeView) SetAutoExpandDelay(delay int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView18setAutoExpandDelayEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), delay)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int indentation()
func (this *QTreeView) Indentation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView11indentationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndentation(int)
func (this *QTreeView) SetIndentation(i int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView14setIndentationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetIndentation()
func (this *QTreeView) ResetIndentation() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView16resetIndentationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rootIsDecorated()
func (this *QTreeView) RootIsDecorated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView15rootIsDecoratedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRootIsDecorated(_Bool)
func (this *QTreeView) SetRootIsDecorated(show bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView18setRootIsDecoratedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), show)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool uniformRowHeights()
func (this *QTreeView) UniformRowHeights() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView17uniformRowHeightsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUniformRowHeights(_Bool)
func (this *QTreeView) SetUniformRowHeights(uniform bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView20setUniformRowHeightsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), uniform)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool itemsExpandable()
func (this *QTreeView) ItemsExpandable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView15itemsExpandableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemsExpandable(_Bool)
func (this *QTreeView) SetItemsExpandable(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView18setItemsExpandableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool expandsOnDoubleClick()
func (this *QTreeView) ExpandsOnDoubleClick() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView20expandsOnDoubleClickEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpandsOnDoubleClick(_Bool)
func (this *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView23setExpandsOnDoubleClickEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnViewportPosition(int)
func (this *QTreeView) ColumnViewportPosition(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView22columnViewportPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnWidth(int)
func (this *QTreeView) ColumnWidth(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView11columnWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnWidth(int, int)
func (this *QTreeView) SetColumnWidth(column int, width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView14setColumnWidthEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnAt(int)
func (this *QTreeView) ColumnAt(x int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView8columnAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isColumnHidden(int)
func (this *QTreeView) IsColumnHidden(column int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView14isColumnHiddenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnHidden(int, _Bool)
func (this *QTreeView) SetColumnHidden(column int, hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView15setColumnHiddenEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isHeaderHidden()
func (this *QTreeView) IsHeaderHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView14isHeaderHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeaderHidden(_Bool)
func (this *QTreeView) SetHeaderHidden(hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView15setHeaderHiddenEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRowHidden(int, const QModelIndex &)
func (this *QTreeView) IsRowHidden(row int, parent *qtcore.QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView11isRowHiddenEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowHidden(int, const QModelIndex &, _Bool)
func (this *QTreeView) SetRowHidden(row int, parent *qtcore.QModelIndex, hide bool) {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12setRowHiddenEiRK11QModelIndexb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFirstColumnSpanned(int, const QModelIndex &)
func (this *QTreeView) IsFirstColumnSpanned(row int, parent *qtcore.QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFirstColumnSpanned(int, const QModelIndex &, _Bool)
func (this *QTreeView) SetFirstColumnSpanned(row int, parent *qtcore.QModelIndex, span bool) {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, span)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExpanded(const QModelIndex &)
func (this *QTreeView) IsExpanded(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10isExpandedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpanded(const QModelIndex &, _Bool)
func (this *QTreeView) SetExpanded(index *qtcore.QModelIndex, expand bool) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11setExpandedERK11QModelIndexb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, expand)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortingEnabled(_Bool)
func (this *QTreeView) SetSortingEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView17setSortingEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortingEnabled()
func (this *QTreeView) IsSortingEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView16isSortingEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAnimated(_Bool)
func (this *QTreeView) SetAnimated(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11setAnimatedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:124
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAnimated()
func (this *QTreeView) IsAnimated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10isAnimatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAllColumnsShowFocus(_Bool)
func (this *QTreeView) SetAllColumnsShowFocus(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView22setAllColumnsShowFocusEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool allColumnsShowFocus()
func (this *QTreeView) AllColumnsShowFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView19allColumnsShowFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordWrap(_Bool)
func (this *QTreeView) SetWordWrap(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11setWordWrapEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wordWrap()
func (this *QTreeView) WordWrap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView8wordWrapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTreePosition(int)
func (this *QTreeView) SetTreePosition(logicalIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView15setTreePositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:133
// index:0
// Public Visibility=Default Availability=Available
// [4] int treePosition()
func (this *QTreeView) TreePosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView12treePositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:135
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void keyboardSearch(const QString &)
func (this *QTreeView) KeyboardSearch(search string) {
	var tmpArg0 = qtcore.NewQString_5(search)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView14keyboardSearchERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:137
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect visualRect(const QModelIndex &)
func (this *QTreeView) VisualRect(index *qtcore.QModelIndex) *qtcore.QRect /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10visualRectERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QTreeView) ScrollTo(index *qtcore.QModelIndex, hint int) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex indexAt(const QPoint &)
func (this *QTreeView) IndexAt(p *qtcore.QPoint) *qtcore.QModelIndex /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView7indexAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:140
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex indexAbove(const QModelIndex &)
func (this *QTreeView) IndexAbove(index *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10indexAboveERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:141
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex indexBelow(const QModelIndex &)
func (this *QTreeView) IndexBelow(index *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10indexBelowERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:143
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doItemsLayout()
func (this *QTreeView) DoItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13doItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reset()
func (this *QTreeView) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortByColumn(int, Qt::SortOrder)
func (this *QTreeView) SortByColumn(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:161
// index:1
// Public Visibility=Default Availability=Available
// [-2] void sortByColumn(int)
func (this *QTreeView) SortByColumn_1(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:149
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void selectAll()
func (this *QTreeView) SelectAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView9selectAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expanded(const QModelIndex &)
func (this *QTreeView) Expanded(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8expandedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapsed(const QModelIndex &)
func (this *QTreeView) Collapsed(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView9collapsedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hideColumn(int)
func (this *QTreeView) HideColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView10hideColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showColumn(int)
func (this *QTreeView) ShowColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView10showColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expand(const QModelIndex &)
func (this *QTreeView) Expand(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView6expandERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapse(const QModelIndex &)
func (this *QTreeView) Collapse(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8collapseERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeColumnToContents(int)
func (this *QTreeView) ResizeColumnToContents(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView22resizeColumnToContentsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expandAll()
func (this *QTreeView) ExpandAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView9expandAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapseAll()
func (this *QTreeView) CollapseAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11collapseAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expandToDepth(int)
func (this *QTreeView) ExpandToDepth(depth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13expandToDepthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:167
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnResized(int, int, int)
func (this *QTreeView) ColumnResized(column int, oldSize int, newSize int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13columnResizedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, oldSize, newSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:168
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnCountChanged(int, int)
func (this *QTreeView) ColumnCountChanged(oldCount int, newCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView18columnCountChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldCount, newCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:169
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnMoved()
func (this *QTreeView) ColumnMoved() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11columnMovedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:170
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void reexpand()
func (this *QTreeView) Reexpand() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8reexpandEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:171
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void rowsRemoved(const QModelIndex &, int, int)
func (this *QTreeView) RowsRemoved(parent *qtcore.QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11rowsRemovedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:175
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)
func (this *QTreeView) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:176
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsInserted(const QModelIndex &, int, int)
func (this *QTreeView) RowsInserted(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12rowsInsertedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:177
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsAboutToBeRemoved(const QModelIndex &, int, int)
func (this *QTreeView) RowsAboutToBeRemoved(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView20rowsAboutToBeRemovedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:179
// index:0
// Protected virtual Visibility=Default Availability=Available
// [24] QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QTreeView) MoveCursor(cursorAction int, modifiers int) *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cursorAction, modifiers)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:180
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int horizontalOffset()
func (this *QTreeView) HorizontalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView16horizontalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:181
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int verticalOffset()
func (this *QTreeView) VerticalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView14verticalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)
func (this *QTreeView) SetSelection(rect *qtcore.QRect, command int) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:184
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QRegion visualRegionForSelection(const QItemSelection &)
func (this *QTreeView) VisualRegionForSelection(selection *qtcore.QItemSelection) *qtgui.QRegion /*123*/ {
	var convArg0 = selection.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView24visualRegionForSelectionERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:185
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QModelIndexList selectedIndexes()
func (this *QTreeView) SelectedIndexes() *qtcore.QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView15selectedIndexesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:187
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QTreeView) TimerEvent(event *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:188
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QTreeView) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:190
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawTree(QPainter *, const QRegion &)
func (this *QTreeView) DrawTree(painter *qtgui.QPainter /*777 QPainter **/, region *qtgui.QRegion) {
	var convArg0 = painter.GetCthis()
	var convArg1 = region.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView8drawTreeEP8QPainterRK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:191
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawRow(QPainter *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QTreeView) DrawRow(painter *qtgui.QPainter /*777 QPainter **/, options *QStyleOptionViewItem, index *qtcore.QModelIndex) {
	var convArg0 = painter.GetCthis()
	var convArg1 = options.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView7drawRowEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:194
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawBranches(QPainter *, const QRect &, const QModelIndex &)
func (this *QTreeView) DrawBranches(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, index *qtcore.QModelIndex) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg2 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView12drawBranchesEP8QPainterRK5QRectRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:198
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QTreeView) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:199
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QTreeView) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:200
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QTreeView) MouseDoubleClickEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:201
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QTreeView) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:202
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QTreeView) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:204
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)
func (this *QTreeView) DragMoveEvent(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:206
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)
func (this *QTreeView) ViewportEvent(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13viewportEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometries()
func (this *QTreeView) UpdateGeometries() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView16updateGeometriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:210
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint()
func (this *QTreeView) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView16viewportSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:212
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int sizeHintForColumn(int)
func (this *QTreeView) SizeHintForColumn(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView17sizeHintForColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:213
// index:0
// Protected Visibility=Default Availability=Available
// [4] int indexRowSizeHint(const QModelIndex &)
func (this *QTreeView) IndexRowSizeHint(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView16indexRowSizeHintERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:214
// index:0
// Protected Visibility=Default Availability=Available
// [4] int rowHeight(const QModelIndex &)
func (this *QTreeView) RowHeight(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView9rowHeightERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:216
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void horizontalScrollbarAction(int)
func (this *QTreeView) HorizontalScrollbarAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView25horizontalScrollbarActionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool isIndexHidden(const QModelIndex &)
func (this *QTreeView) IsIndexHidden(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView13isIndexHiddenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void selectionChanged(const QItemSelection &, const QItemSelection &)
func (this *QTreeView) SelectionChanged(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) {
	var convArg0 = selected.GetCthis()
	var convArg1 = deselected.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView16selectionChangedERK14QItemSelectionS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)
func (this *QTreeView) CurrentChanged(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView14currentChangedERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
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
