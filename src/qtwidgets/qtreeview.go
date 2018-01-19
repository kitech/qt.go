//  header block begin
// /usr/include/qt/QtWidgets/qtreeview.h
// #include <qtreeview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QTreeView struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qtreeview.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTreeView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:71
// index:0
// void QTreeView(class QWidget *)
func NewQTreeView(parent unsafe.Pointer) *QTreeView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QTreeView{cthis}
}

// /usr/include/qt/QtWidgets/qtreeview.h:72
// index:0
// virtual
// void ~QTreeView()
func DeleteQTreeView(*QTreeView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:74
// index:0
// virtual
// void setModel(class QAbstractItemModel *)
func (this *QTreeView) SetModel(model unsafe.Pointer) {
	// 0: (, QAbstractItemModel * model), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.cthis, model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:75
// index:0
// virtual
// void setRootIndex(const class QModelIndex &)
func (this *QTreeView) SetRootIndex(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:76
// index:0
// virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QTreeView) SetSelectionModel(selectionModel unsafe.Pointer) {
	// 0: (, QItemSelectionModel * selectionModel), (selectionModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_VOID, this.cthis, selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:78
// index:0
// QHeaderView * header()
func (this *QTreeView) Header() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView6headerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:79
// index:0
// void setHeader(class QHeaderView *)
func (this *QTreeView) SetHeader(header unsafe.Pointer) {
	// 0: (, QHeaderView * header), (header)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9setHeaderEP11QHeaderView", ffiqt.FFI_TYPE_VOID, this.cthis, header)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:81
// index:0
// int autoExpandDelay()
func (this *QTreeView) AutoExpandDelay() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15autoExpandDelayEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:82
// index:0
// void setAutoExpandDelay(int)
func (this *QTreeView) SetAutoExpandDelay(delay int) {
	// 0: (, int delay), (&delay)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setAutoExpandDelayEi", ffiqt.FFI_TYPE_VOID, this.cthis, &delay)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:84
// index:0
// int indentation()
func (this *QTreeView) Indentation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11indentationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:85
// index:0
// void setIndentation(int)
func (this *QTreeView) SetIndentation(i int) {
	// 0: (, int i), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14setIndentationEi", ffiqt.FFI_TYPE_VOID, this.cthis, &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:86
// index:0
// void resetIndentation()
func (this *QTreeView) ResetIndentation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16resetIndentationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:88
// index:0
// bool rootIsDecorated()
func (this *QTreeView) RootIsDecorated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15rootIsDecoratedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:89
// index:0
// void setRootIsDecorated(_Bool)
func (this *QTreeView) SetRootIsDecorated(show bool) {
	// 0: (, bool show), (&show)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setRootIsDecoratedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:91
// index:0
// bool uniformRowHeights()
func (this *QTreeView) UniformRowHeights() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView17uniformRowHeightsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:92
// index:0
// void setUniformRowHeights(_Bool)
func (this *QTreeView) SetUniformRowHeights(uniform bool) {
	// 0: (, bool uniform), (&uniform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView20setUniformRowHeightsEb", ffiqt.FFI_TYPE_VOID, this.cthis, &uniform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:94
// index:0
// bool itemsExpandable()
func (this *QTreeView) ItemsExpandable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15itemsExpandableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:95
// index:0
// void setItemsExpandable(_Bool)
func (this *QTreeView) SetItemsExpandable(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setItemsExpandableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:97
// index:0
// bool expandsOnDoubleClick()
func (this *QTreeView) ExpandsOnDoubleClick() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView20expandsOnDoubleClickEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:98
// index:0
// void setExpandsOnDoubleClick(_Bool)
func (this *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView23setExpandsOnDoubleClickEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:100
// index:0
// int columnViewportPosition(int)
func (this *QTreeView) ColumnViewportPosition(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView22columnViewportPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:101
// index:0
// int columnWidth(int)
func (this *QTreeView) ColumnWidth(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11columnWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:102
// index:0
// void setColumnWidth(int, int)
func (this *QTreeView) SetColumnWidth(column int, width int) {
	// 0: (, int column, int width), (&column, &width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14setColumnWidthEii", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:103
// index:0
// int columnAt(int)
func (this *QTreeView) ColumnAt(x int) {
	// 0: (, int x), (&x)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8columnAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:105
// index:0
// bool isColumnHidden(int)
func (this *QTreeView) IsColumnHidden(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14isColumnHiddenEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:106
// index:0
// void setColumnHidden(int, _Bool)
func (this *QTreeView) SetColumnHidden(column int, hide bool) {
	// 0: (, int column, bool hide), (&column, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setColumnHiddenEib", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:108
// index:0
// bool isHeaderHidden()
func (this *QTreeView) IsHeaderHidden() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14isHeaderHiddenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:109
// index:0
// void setHeaderHidden(_Bool)
func (this *QTreeView) SetHeaderHidden(hide bool) {
	// 0: (, bool hide), (&hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setHeaderHiddenEb", ffiqt.FFI_TYPE_VOID, this.cthis, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:111
// index:0
// bool isRowHidden(int, const class QModelIndex &)
func (this *QTreeView) IsRowHidden(row int, parent unsafe.Pointer) {
	// 0: (, int row, const QModelIndex & parent), (&row, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11isRowHiddenEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:112
// index:0
// void setRowHidden(int, const class QModelIndex &, _Bool)
func (this *QTreeView) SetRowHidden(row int, parent unsafe.Pointer, hide bool) {
	// 0: (, int row, const QModelIndex & parent, bool hide), (&row, parent, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12setRowHiddenEiRK11QModelIndexb", ffiqt.FFI_TYPE_VOID, this.cthis, &row, parent, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:114
// index:0
// bool isFirstColumnSpanned(int, const class QModelIndex &)
func (this *QTreeView) IsFirstColumnSpanned(row int, parent unsafe.Pointer) {
	// 0: (, int row, const QModelIndex & parent), (&row, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:115
// index:0
// void setFirstColumnSpanned(int, const class QModelIndex &, _Bool)
func (this *QTreeView) SetFirstColumnSpanned(row int, parent unsafe.Pointer, span bool) {
	// 0: (, int row, const QModelIndex & parent, bool span), (&row, parent, &span)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb", ffiqt.FFI_TYPE_VOID, this.cthis, &row, parent, &span)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:117
// index:0
// bool isExpanded(const class QModelIndex &)
func (this *QTreeView) IsExpanded(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10isExpandedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:118
// index:0
// void setExpanded(const class QModelIndex &, _Bool)
func (this *QTreeView) SetExpanded(index unsafe.Pointer, expand bool) {
	// 0: (, const QModelIndex & index, bool expand), (index, &expand)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setExpandedERK11QModelIndexb", ffiqt.FFI_TYPE_VOID, this.cthis, index, &expand)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:120
// index:0
// void setSortingEnabled(_Bool)
func (this *QTreeView) SetSortingEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17setSortingEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:121
// index:0
// bool isSortingEnabled()
func (this *QTreeView) IsSortingEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16isSortingEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:123
// index:0
// void setAnimated(_Bool)
func (this *QTreeView) SetAnimated(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setAnimatedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:124
// index:0
// bool isAnimated()
func (this *QTreeView) IsAnimated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10isAnimatedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:126
// index:0
// void setAllColumnsShowFocus(_Bool)
func (this *QTreeView) SetAllColumnsShowFocus(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView22setAllColumnsShowFocusEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:127
// index:0
// bool allColumnsShowFocus()
func (this *QTreeView) AllColumnsShowFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView19allColumnsShowFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:129
// index:0
// void setWordWrap(_Bool)
func (this *QTreeView) SetWordWrap(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setWordWrapEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:130
// index:0
// bool wordWrap()
func (this *QTreeView) WordWrap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8wordWrapEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:132
// index:0
// void setTreePosition(int)
func (this *QTreeView) SetTreePosition(logicalIndex int) {
	// 0: (, int logicalIndex), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setTreePositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:133
// index:0
// int treePosition()
func (this *QTreeView) TreePosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView12treePositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:135
// index:0
// virtual
// void keyboardSearch(const class QString &)
func (this *QTreeView) KeyboardSearch(search unsafe.Pointer) {
	// 0: (, const QString & search), (search)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14keyboardSearchERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, search)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:137
// index:0
// virtual
// QRect visualRect(const class QModelIndex &)
func (this *QTreeView) VisualRect(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:138
// index:0
// virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QTreeView) ScrollTo(index unsafe.Pointer, hint int) {
	// 0: (, const QModelIndex & index, QAbstractItemView::ScrollHint hint), (index, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.cthis, index, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:139
// index:0
// virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QTreeView) IndexAt(p unsafe.Pointer) {
	// 0: (, const QPoint & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView7indexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:140
// index:0
// QModelIndex indexAbove(const class QModelIndex &)
func (this *QTreeView) IndexAbove(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10indexAboveERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:141
// index:0
// QModelIndex indexBelow(const class QModelIndex &)
func (this *QTreeView) IndexBelow(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10indexBelowERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:143
// index:0
// virtual
// void doItemsLayout()
func (this *QTreeView) DoItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13doItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:144
// index:0
// virtual
// void reset()
func (this *QTreeView) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView5resetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:146
// index:0
// void sortByColumn(int, Qt::SortOrder)
func (this *QTreeView) SortByColumn(column int, order int) {
	// 0: (, int column, Qt::SortOrder order), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:161
// index:1
// void sortByColumn(int)
func (this *QTreeView) SortByColumn_1(column int) {
	// 1: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:149
// index:0
// virtual
// void selectAll()
func (this *QTreeView) SelectAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9selectAllEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:152
// index:0
// void expanded(const class QModelIndex &)
func (this *QTreeView) Expanded(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8expandedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:153
// index:0
// void collapsed(const class QModelIndex &)
func (this *QTreeView) Collapsed(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9collapsedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:156
// index:0
// void hideColumn(int)
func (this *QTreeView) HideColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10hideColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:157
// index:0
// void showColumn(int)
func (this *QTreeView) ShowColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10showColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:158
// index:0
// void expand(const class QModelIndex &)
func (this *QTreeView) Expand(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView6expandERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:159
// index:0
// void collapse(const class QModelIndex &)
func (this *QTreeView) Collapse(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8collapseERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:160
// index:0
// void resizeColumnToContents(int)
func (this *QTreeView) ResizeColumnToContents(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView22resizeColumnToContentsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:162
// index:0
// void expandAll()
func (this *QTreeView) ExpandAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9expandAllEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:163
// index:0
// void collapseAll()
func (this *QTreeView) CollapseAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11collapseAllEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:164
// index:0
// void expandToDepth(int)
func (this *QTreeView) ExpandToDepth(depth int) {
	// 0: (, int depth), (&depth)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13expandToDepthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &depth)
	gopp.ErrPrint(err, rv)
}

//  body block end
