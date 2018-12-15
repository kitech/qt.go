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
import "log"
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

// void rowsRemoved(const QModelIndex &, int, int)
func (this *QTreeView) InheritRowsRemoved(f func(parent *qtcore.QModelIndex, first int, last int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsRemoved", f)
}

// void scrollContentsBy(int, int)
func (this *QTreeView) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// void rowsInserted(const QModelIndex &, int, int)
func (this *QTreeView) InheritRowsInserted(f func(parent *qtcore.QModelIndex, start int, end_ int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsInserted", f)
}

// void rowsAboutToBeRemoved(const QModelIndex &, int, int)
func (this *QTreeView) InheritRowsAboutToBeRemoved(f func(parent *qtcore.QModelIndex, start int, end_ int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "rowsAboutToBeRemoved", f)
}

// QModelIndex moveCursor(QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
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

// void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)
func (this *QTreeView) InheritSetSelection(f func(rect *qtcore.QRect, command int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setSelection", f)
}

// QRegion visualRegionForSelection(const QItemSelection &)
func (this *QTreeView) InheritVisualRegionForSelection(f func(selection *qtcore.QItemSelection) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "visualRegionForSelection", f)
}

// QModelIndexList selectedIndexes()
func (this *QTreeView) InheritSelectedIndexes(f func() *qtcore.QModelIndexList /*9999*/) {
	qtrt.SetAllInheritCallback(this, "selectedIndexes", f)
}

// void timerEvent(QTimerEvent *)
func (this *QTreeView) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QTreeView) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void drawTree(QPainter *, const QRegion &)
func (this *QTreeView) InheritDrawTree(f func(painter *qtgui.QPainter /*777 QPainter **/, region *qtgui.QRegion) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawTree", f)
}

// void drawRow(QPainter *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QTreeView) InheritDrawRow(f func(painter *qtgui.QPainter /*777 QPainter **/, options *QStyleOptionViewItem, index *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawRow", f)
}

// void drawBranches(QPainter *, const QRect &, const QModelIndex &)
func (this *QTreeView) InheritDrawBranches(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, index *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawBranches", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QTreeView) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QTreeView) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QTreeView) InheritMouseDoubleClickEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QTreeView) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QTreeView) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void dragMoveEvent(QDragMoveEvent *)
func (this *QTreeView) InheritDragMoveEvent(f func(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// bool viewportEvent(QEvent *)
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

// int indexRowSizeHint(const QModelIndex &)
func (this *QTreeView) InheritIndexRowSizeHint(f func(index *qtcore.QModelIndex) int) {
	qtrt.SetAllInheritCallback(this, "indexRowSizeHint", f)
}

// int rowHeight(const QModelIndex &)
func (this *QTreeView) InheritRowHeight(f func(index *qtcore.QModelIndex) int) {
	qtrt.SetAllInheritCallback(this, "rowHeight", f)
}

// void horizontalScrollbarAction(int)
func (this *QTreeView) InheritHorizontalScrollbarAction(f func(action int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "horizontalScrollbarAction", f)
}

// bool isIndexHidden(const QModelIndex &)
func (this *QTreeView) InheritIsIndexHidden(f func(index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "isIndexHidden", f)
}

// void selectionChanged(const QItemSelection &, const QItemSelection &)
func (this *QTreeView) InheritSelectionChanged(f func(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) /*void*/) {
	qtrt.SetAllInheritCallback(this, "selectionChanged", f)
}

// void currentChanged(const QModelIndex &, const QModelIndex &)
func (this *QTreeView) InheritCurrentChanged(f func(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "currentChanged", f)
}

/*

 */
type QTreeView struct {
	*QAbstractItemView
}
type QTreeView_ITF interface {
	QAbstractItemView_ITF
	QTreeView_PTR() *QTreeView
}

func (ptr *QTreeView) QTreeView_PTR() *QTreeView { return ptr }

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

// /usr/include/qt/QtWidgets/qtreeview.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTreeView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreeview.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeView(QWidget *)

/*
Constructs a tree view with a parent to represent a model's data. Use setModel() to set the model.

See also QAbstractItemModel.
*/
func (*QTreeView) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QTreeView {
	return NewQTreeView(parent)
}
func NewQTreeView(parent QWidget_ITF /*777 QWidget **/) *QTreeView {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTreeView")
	return gothis
}

// /usr/include/qt/QtWidgets/qtreeview.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeView(QWidget *)

/*
Constructs a tree view with a parent to represent a model's data. Use setModel() to set the model.

See also QAbstractItemModel.
*/
func (*QTreeView) NewForInheritp() *QTreeView {
	return NewQTreeViewp()
}
func NewQTreeViewp() *QTreeView {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTreeView")
	return gothis
}

// /usr/include/qt/QtWidgets/qtreeview.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTreeView()

/*

 */
func DeleteQTreeView(this *QTreeView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtreeview.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)

/*
Reimplemented from QAbstractItemView::setModel().
*/
func (this *QTreeView) SetModel(model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRootIndex(const QModelIndex &)

/*
Reimplemented from QAbstractItemView::setRootIndex().
*/
func (this *QTreeView) SetRootIndex(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12setRootIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSelectionModel(QItemSelectionModel *)

/*
Reimplemented from QAbstractItemView::setSelectionModel().
*/
func (this *QTreeView) SetSelectionModel(selectionModel qtcore.QItemSelectionModel_ITF /*777 QItemSelectionModel **/) {
	var convArg0 unsafe.Pointer
	if selectionModel != nil && selectionModel.QItemSelectionModel_PTR() != nil {
		convArg0 = selectionModel.QItemSelectionModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView17setSelectionModelEP19QItemSelectionModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QHeaderView * header() const

/*
Returns the header for the tree view.

See also setHeader() and QAbstractItemModel::headerData().
*/
func (this *QTreeView) Header() *QHeaderView /*777 QHeaderView **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView6headerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQHeaderViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreeview.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeader(QHeaderView *)

/*
Sets the header for the tree view, to the given header.

The view takes ownership over the given header and deletes it when a new header is set.

See also QAbstractItemModel::headerData().
*/
func (this *QTreeView) SetHeader(header QHeaderView_ITF /*777 QHeaderView **/) {
	var convArg0 unsafe.Pointer
	if header != nil && header.QHeaderView_PTR() != nil {
		convArg0 = header.QHeaderView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView9setHeaderEP11QHeaderView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int autoExpandDelay() const

/*

 */
func (this *QTreeView) AutoExpandDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView15autoExpandDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoExpandDelay(int)

/*

 */
func (this *QTreeView) SetAutoExpandDelay(delay int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView18setAutoExpandDelayEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), delay)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int indentation() const

/*

 */
func (this *QTreeView) Indentation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView11indentationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndentation(int)

/*

 */
func (this *QTreeView) SetIndentation(i int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView14setIndentationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetIndentation()

/*

 */
func (this *QTreeView) ResetIndentation() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView16resetIndentationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rootIsDecorated() const

/*

 */
func (this *QTreeView) RootIsDecorated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView15rootIsDecoratedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRootIsDecorated(bool)

/*

 */
func (this *QTreeView) SetRootIsDecorated(show bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView18setRootIsDecoratedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), show)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool uniformRowHeights() const

/*

 */
func (this *QTreeView) UniformRowHeights() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView17uniformRowHeightsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUniformRowHeights(bool)

/*

 */
func (this *QTreeView) SetUniformRowHeights(uniform bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView20setUniformRowHeightsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), uniform)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool itemsExpandable() const

/*

 */
func (this *QTreeView) ItemsExpandable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView15itemsExpandableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemsExpandable(bool)

/*

 */
func (this *QTreeView) SetItemsExpandable(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView18setItemsExpandableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool expandsOnDoubleClick() const

/*

 */
func (this *QTreeView) ExpandsOnDoubleClick() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView20expandsOnDoubleClickEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpandsOnDoubleClick(bool)

/*

 */
func (this *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView23setExpandsOnDoubleClickEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnViewportPosition(int) const

/*
Returns the horizontal position of the column in the viewport.
*/
func (this *QTreeView) ColumnViewportPosition(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView22columnViewportPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnWidth(int) const

/*
Returns the width of the column.

See also resizeColumnToContents() and setColumnWidth().
*/
func (this *QTreeView) ColumnWidth(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView11columnWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnWidth(int, int)

/*
Sets the width of the given column to the width specified.

This function was introduced in  Qt 4.2.

See also columnWidth() and resizeColumnToContents().
*/
func (this *QTreeView) SetColumnWidth(column int, width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView14setColumnWidthEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnAt(int) const

/*
Returns the column in the tree view whose header covers the x coordinate given.
*/
func (this *QTreeView) ColumnAt(x int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView8columnAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isColumnHidden(int) const

/*
Returns true if the column is hidden; otherwise returns false.

See also hideColumn() and isRowHidden().
*/
func (this *QTreeView) IsColumnHidden(column int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView14isColumnHiddenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnHidden(int, bool)

/*
If hide is true the column is hidden, otherwise the column is shown.

See also isColumnHidden(), hideColumn(), and setRowHidden().
*/
func (this *QTreeView) SetColumnHidden(column int, hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView15setColumnHiddenEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isHeaderHidden() const

/*

 */
func (this *QTreeView) IsHeaderHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView14isHeaderHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeaderHidden(bool)

/*

 */
func (this *QTreeView) SetHeaderHidden(hide bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView15setHeaderHiddenEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRowHidden(int, const QModelIndex &) const

/*
Returns true if the item in the given row of the parent is hidden; otherwise returns false.

See also setRowHidden() and isColumnHidden().
*/
func (this *QTreeView) IsRowHidden(row int, parent qtcore.QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView11isRowHiddenEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowHidden(int, const QModelIndex &, bool)

/*
If hide is true the row with the given parent is hidden, otherwise the row is shown.

See also isRowHidden() and setColumnHidden().
*/
func (this *QTreeView) SetRowHidden(row int, parent qtcore.QModelIndex_ITF, hide bool) {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12setRowHiddenEiRK11QModelIndexb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFirstColumnSpanned(int, const QModelIndex &) const

/*
Returns true if the item in first column in the given row of the parent is spanning all the columns; otherwise returns false.

This function was introduced in  Qt 4.3.

See also setFirstColumnSpanned().
*/
func (this *QTreeView) IsFirstColumnSpanned(row int, parent qtcore.QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFirstColumnSpanned(int, const QModelIndex &, bool)

/*
If span is true the item in the first column in the row with the given parent is set to span all columns, otherwise all items on the row are shown.

This function was introduced in  Qt 4.3.

See also isFirstColumnSpanned().
*/
func (this *QTreeView) SetFirstColumnSpanned(row int, parent qtcore.QModelIndex_ITF, span bool) {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1, span)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExpanded(const QModelIndex &) const

/*
Returns true if the model item index is expanded; otherwise returns false.

See also expand(), expanded(), and setExpanded().
*/
func (this *QTreeView) IsExpanded(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10isExpandedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpanded(const QModelIndex &, bool)

/*
Sets the item referred to by index to either collapse or expanded, depending on the value of expanded.

See also expanded(), expand(), and isExpanded().
*/
func (this *QTreeView) SetExpanded(index qtcore.QModelIndex_ITF, expand bool) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11setExpandedERK11QModelIndexb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, expand)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortingEnabled(bool)

/*

 */
func (this *QTreeView) SetSortingEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView17setSortingEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:122
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortingEnabled() const

/*

 */
func (this *QTreeView) IsSortingEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView16isSortingEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAnimated(bool)

/*

 */
func (this *QTreeView) SetAnimated(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11setAnimatedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:125
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAnimated() const

/*

 */
func (this *QTreeView) IsAnimated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10isAnimatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAllColumnsShowFocus(bool)

/*

 */
func (this *QTreeView) SetAllColumnsShowFocus(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView22setAllColumnsShowFocusEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:128
// index:0
// Public Visibility=Default Availability=Available
// [1] bool allColumnsShowFocus() const

/*

 */
func (this *QTreeView) AllColumnsShowFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView19allColumnsShowFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordWrap(bool)

/*

 */
func (this *QTreeView) SetWordWrap(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11setWordWrapEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wordWrap() const

/*

 */
func (this *QTreeView) WordWrap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView8wordWrapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTreePosition(int)

/*
This specifies that the tree structure should be placed at logical index index. If  set to -1 then the tree will always follow visual index 0.

This function was introduced in  Qt 5.2.

See also treePosition(), QHeaderView::swapSections(), and QHeaderView::moveSection().
*/
func (this *QTreeView) SetTreePosition(logicalIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView15setTreePositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), logicalIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int treePosition() const

/*
Return the logical index the tree is set on. If the return value is -1 then the tree is placed on the visual index 0.

This function was introduced in  Qt 5.2.

See also setTreePosition().
*/
func (this *QTreeView) TreePosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView12treePositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:136
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void keyboardSearch(const QString &)

/*
Reimplemented from QAbstractItemView::keyboardSearch().
*/
func (this *QTreeView) KeyboardSearch(search string) {
	var tmpArg0 = qtcore.NewQString5(search)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView14keyboardSearchERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect visualRect(const QModelIndex &) const

/*
Reimplemented from QAbstractItemView::visualRect().

Returns the rectangle on the viewport occupied by the item at index. If the index is not visible or explicitly hidden, the returned rectangle is invalid.
*/
func (this *QTreeView) VisualRect(index qtcore.QModelIndex_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10visualRectERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, QAbstractItemView::ScrollHint)

/*
Reimplemented from QAbstractItemView::scrollTo().

Scroll the contents of the tree view until the given model item index is visible. The hint parameter specifies more precisely where the item should be located after the operation. If any of the parents of the model item are collapsed, they will be expanded to ensure that the model item is visible.
*/
func (this *QTreeView) ScrollTo(index qtcore.QModelIndex_ITF, hint int) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, QAbstractItemView::ScrollHint)

/*
Reimplemented from QAbstractItemView::scrollTo().

Scroll the contents of the tree view until the given model item index is visible. The hint parameter specifies more precisely where the item should be located after the operation. If any of the parents of the model item are collapsed, they will be expanded to ensure that the model item is visible.
*/
func (this *QTreeView) ScrollTop(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, QAbstractItemView::ScrollHint=Enum, QAbstractItemView::ScrollHint=Enum, , Invalid
	hint := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:140
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex indexAt(const QPoint &) const

/*
Reimplemented from QAbstractItemView::indexAt().
*/
func (this *QTreeView) IndexAt(p qtcore.QPoint_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView7indexAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:141
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex indexAbove(const QModelIndex &) const

/*
Returns the model index of the item above index.
*/
func (this *QTreeView) IndexAbove(index qtcore.QModelIndex_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10indexAboveERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:142
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex indexBelow(const QModelIndex &) const

/*
Returns the model index of the item below index.
*/
func (this *QTreeView) IndexBelow(index qtcore.QModelIndex_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView10indexBelowERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doItemsLayout()

/*

 */
func (this *QTreeView) DoItemsLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13doItemsLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:145
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reset()

/*
Reimplemented from QAbstractItemView::reset().
*/
func (this *QTreeView) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortByColumn(int, Qt::SortOrder)

/*
Sets the model up for sorting by the values in the given column and order.

column may be -1, in which case no sort indicator will be shown and the model will return to its natural, unsorted order. Note that not all models support this and may even crash in this case.

This function was introduced in  Qt 4.2.

See also sortingEnabled.
*/
func (this *QTreeView) SortByColumn(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:162
// index:1
// Public Visibility=Default Availability=Available
// [-2] void sortByColumn(int)

/*
Sets the model up for sorting by the values in the given column and order.

column may be -1, in which case no sort indicator will be shown and the model will return to its natural, unsorted order. Note that not all models support this and may even crash in this case.

This function was introduced in  Qt 4.2.

See also sortingEnabled.
*/
func (this *QTreeView) SortByColumn1(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:150
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void selectAll()

/*
Reimplemented from QAbstractItemView::selectAll().
*/
func (this *QTreeView) SelectAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView9selectAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expanded(const QModelIndex &)

/*
This signal is emitted when the item specified by index is expanded.

See also setExpanded().
*/
func (this *QTreeView) Expanded(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8expandedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapsed(const QModelIndex &)

/*
This signal is emitted when the item specified by index is collapsed.
*/
func (this *QTreeView) Collapsed(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView9collapsedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hideColumn(int)

/*
Hides the column given.

Note: This function should only be called after the model has been initialized, as the view needs to know the number of columns in order to hide column.

See also showColumn() and setColumnHidden().
*/
func (this *QTreeView) HideColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView10hideColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showColumn(int)

/*
Shows the given column in the tree view.

See also hideColumn() and setColumnHidden().
*/
func (this *QTreeView) ShowColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView10showColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expand(const QModelIndex &)

/*
Expands the model item specified by the index.

See also expanded().
*/
func (this *QTreeView) Expand(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView6expandERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapse(const QModelIndex &)

/*
Collapses the model item specified by the index.

See also collapsed().
*/
func (this *QTreeView) Collapse(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8collapseERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resizeColumnToContents(int)

/*
Resizes the column given to the size of its contents.

See also columnWidth(), setColumnWidth(), sizeHintForColumn(), and QHeaderView::resizeContentsPrecision().
*/
func (this *QTreeView) ResizeColumnToContents(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView22resizeColumnToContentsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expandAll()

/*
Expands all expandable items.

Warning: if the model contains a large number of items, this function will take some time to execute.

This function was introduced in  Qt 4.2.

See also collapseAll(), expand(), collapse(), and setExpanded().
*/
func (this *QTreeView) ExpandAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView9expandAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapseAll()

/*
Collapses all expanded items.

This function was introduced in  Qt 4.2.

See also expandAll(), expand(), collapse(), and setExpanded().
*/
func (this *QTreeView) CollapseAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11collapseAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expandToDepth(int)

/*
Expands all expandable items to the given depth.

This function was introduced in  Qt 4.3.

See also expandAll(), collapseAll(), expand(), collapse(), and setExpanded().
*/
func (this *QTreeView) ExpandToDepth(depth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13expandToDepthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:168
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnResized(int, int, int)

/*
This function is called whenever column's size is changed in the header. oldSize and newSize give the previous size and the new size in pixels.

See also setColumnWidth().
*/
func (this *QTreeView) ColumnResized(column int, oldSize int, newSize int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13columnResizedEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, oldSize, newSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:169
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnCountChanged(int, int)

/*
Informs the tree view that the number of columns in the tree view has changed from oldCount to newCount.
*/
func (this *QTreeView) ColumnCountChanged(oldCount int, newCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView18columnCountChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldCount, newCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:170
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void columnMoved()

/*
This slot is called whenever a column has been moved.
*/
func (this *QTreeView) ColumnMoved() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11columnMovedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:171
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void reexpand()

/*

 */
func (this *QTreeView) Reexpand() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView8reexpandEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:172
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void rowsRemoved(const QModelIndex &, int, int)

/*
Informs the view that the rows from the start row to the end row inclusive have been removed from the given parent model item.

This function was introduced in  Qt 4.1.
*/
func (this *QTreeView) RowsRemoved(parent qtcore.QModelIndex_ITF, first int, last int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView11rowsRemovedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:176
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)

/*
Reimplemented from QAbstractScrollArea::scrollContentsBy().

Scrolls the contents of the tree view by (dx, dy).
*/
func (this *QTreeView) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:177
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsInserted(const QModelIndex &, int, int)

/*
Reimplemented from QAbstractItemView::rowsInserted().

Informs the view that the rows from the start row to the end row inclusive have been inserted into the parent model item.
*/
func (this *QTreeView) RowsInserted(parent qtcore.QModelIndex_ITF, start int, end_ int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12rowsInsertedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:178
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsAboutToBeRemoved(const QModelIndex &, int, int)

/*
Reimplemented from QAbstractItemView::rowsAboutToBeRemoved().

Informs the view that the rows from the start row to the end row inclusive are about to removed from the given parent model item.
*/
func (this *QTreeView) RowsAboutToBeRemoved(parent qtcore.QModelIndex_ITF, start int, end_ int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView20rowsAboutToBeRemovedERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:180
// index:0
// Protected virtual Visibility=Default Availability=Available
// [24] QModelIndex moveCursor(QAbstractItemView::CursorAction, Qt::KeyboardModifiers)

/*
Reimplemented from QAbstractItemView::moveCursor().

Move the cursor in the way described by cursorAction, using the information provided by the button modifiers.
*/
func (this *QTreeView) MoveCursor(cursorAction int, modifiers int) *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cursorAction, modifiers)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:181
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int horizontalOffset() const

/*
Reimplemented from QAbstractItemView::horizontalOffset().

Returns the horizontal offset of the items in the treeview.

Note that the tree view uses the horizontal header section positions to determine the positions of columns in the view.

See also verticalOffset().
*/
func (this *QTreeView) HorizontalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView16horizontalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:182
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int verticalOffset() const

/*
Reimplemented from QAbstractItemView::verticalOffset().

Returns the vertical offset of the items in the tree view.

See also horizontalOffset().
*/
func (this *QTreeView) VerticalOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView14verticalOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:184
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)

/*
Reimplemented from QAbstractItemView::setSelection().

Applies the selection command to the items in or touched by the rectangle, rect.

See also selectionCommand().
*/
func (this *QTreeView) SetSelection(rect qtcore.QRect_ITF, command int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:185
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QRegion visualRegionForSelection(const QItemSelection &) const

/*
Reimplemented from QAbstractItemView::visualRegionForSelection().

Returns the rectangle from the viewport of the items in the given selection.

Since 4.7, the returned region only contains rectangles intersecting (or included in) the viewport.
*/
func (this *QTreeView) VisualRegionForSelection(selection qtcore.QItemSelection_ITF) *qtgui.QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if selection != nil && selection.QItemSelection_PTR() != nil {
		convArg0 = selection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView24visualRegionForSelectionERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:186
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QModelIndexList selectedIndexes() const

/*
Reimplemented from QAbstractItemView::selectedIndexes().
*/
func (this *QTreeView) SelectedIndexes() *qtcore.QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView15selectedIndexesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:188
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QTreeView) TimerEvent(event qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QTimerEvent_PTR() != nil {
		convArg0 = event.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:189
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QTreeView) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:191
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawTree(QPainter *, const QRegion &) const

/*
Draws the part of the tree intersecting the given region using the specified painter.

This function was introduced in  Qt 4.2.

See also paintEvent().
*/
func (this *QTreeView) DrawTree(painter qtgui.QPainter_ITF /*777 QPainter **/, region qtgui.QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg1 = region.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView8drawTreeEP8QPainterRK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:192
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawRow(QPainter *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
Draws the row in the tree view that contains the model item index, using the painter given. The option controls how the item is displayed.

See also setAlternatingRowColors().
*/
func (this *QTreeView) DrawRow(painter qtgui.QPainter_ITF /*777 QPainter **/, options QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if options != nil && options.QStyleOptionViewItem_PTR() != nil {
		convArg1 = options.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView7drawRowEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:195
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawBranches(QPainter *, const QRect &, const QModelIndex &) const

/*
Draws the branches in the tree view on the same row as the model item index, using the painter given. The branches are drawn in the rectangle specified by rect.
*/
func (this *QTreeView) DrawBranches(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView12drawBranchesEP8QPainterRK5QRectRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:199
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QTreeView) MousePressEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:200
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QTreeView) MouseReleaseEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:201
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseDoubleClickEvent().
*/
func (this *QTreeView) MouseDoubleClickEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:202
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QTreeView) MouseMoveEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:203
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QTreeView) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:205
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QWidget::dragMoveEvent().
*/
func (this *QTreeView) DragMoveEvent(event qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDragMoveEvent_PTR() != nil {
		convArg0 = event.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:207
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)

/*
Reimplemented from QAbstractScrollArea::viewportEvent().
*/
func (this *QTreeView) ViewportEvent(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView13viewportEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:209
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometries()

/*
Reimplemented from QAbstractItemView::updateGeometries().
*/
func (this *QTreeView) UpdateGeometries() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView16updateGeometriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:211
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint() const

/*
Reimplemented from QAbstractScrollArea::viewportSizeHint().
*/
func (this *QTreeView) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView16viewportSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreeview.h:213
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int sizeHintForColumn(int) const

/*
Reimplemented from QAbstractItemView::sizeHintForColumn().

Returns the size hint for the column's width or -1 if there is no model.

If you need to set the width of a given column to a fixed value, call QHeaderView::resizeSection() on the view's header.

If you reimplement this function in a subclass, note that the value you return is only used when resizeColumnToContents() is called. In that case, if a larger column width is required by either the view's header or the item delegate, that width will be used instead.

See also QWidget::sizeHint, header(), and QHeaderView::resizeContentsPrecision().
*/
func (this *QTreeView) SizeHintForColumn(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView17sizeHintForColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:214
// index:0
// Protected Visibility=Default Availability=Available
// [4] int indexRowSizeHint(const QModelIndex &) const

/*
Returns the size hint for the row indicated by index.

See also sizeHintForColumn() and uniformRowHeights().
*/
func (this *QTreeView) IndexRowSizeHint(index qtcore.QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView16indexRowSizeHintERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:215
// index:0
// Protected Visibility=Default Availability=Available
// [4] int rowHeight(const QModelIndex &) const

/*
Returns the height of the row indicated by the given index.

This function was introduced in  Qt 4.3.

See also indexRowSizeHint().
*/
func (this *QTreeView) RowHeight(index qtcore.QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView9rowHeightERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreeview.h:217
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void horizontalScrollbarAction(int)

/*

 */
func (this *QTreeView) HorizontalScrollbarAction(action int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView25horizontalScrollbarActionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool isIndexHidden(const QModelIndex &) const

/*
Reimplemented from QAbstractItemView::isIndexHidden().
*/
func (this *QTreeView) IsIndexHidden(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTreeView13isIndexHiddenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreeview.h:220
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void selectionChanged(const QItemSelection &, const QItemSelection &)

/*
Reimplemented from QAbstractItemView::selectionChanged().
*/
func (this *QTreeView) SelectionChanged(selected qtcore.QItemSelection_ITF, deselected qtcore.QItemSelection_ITF) {
	var convArg0 unsafe.Pointer
	if selected != nil && selected.QItemSelection_PTR() != nil {
		convArg0 = selected.QItemSelection_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if deselected != nil && deselected.QItemSelection_PTR() != nil {
		convArg1 = deselected.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTreeView16selectionChangedERK14QItemSelectionS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)

/*
Reimplemented from QAbstractItemView::currentChanged().
*/
func (this *QTreeView) CurrentChanged(current qtcore.QModelIndex_ITF, previous qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QModelIndex_PTR() != nil {
		convArg0 = current.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QModelIndex_PTR() != nil {
		convArg1 = previous.QModelIndex_PTR().GetCthis()
	}
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
