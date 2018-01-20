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
	*QAbstractItemView
}

func (this *QTreeView) GetCthis() unsafe.Pointer {
	return this.QAbstractItemView.GetCthis()
}

// /usr/include/qt/QtWidgets/qtreeview.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTreeView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:71
// index:0
// void QTreeView(class QWidget *)
func NewQTreeView(parent unsafe.Pointer) *QTreeView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeViewFromPointer(cthis)
	return gothis
}
func NewQTreeViewFromPointer(cthis unsafe.Pointer) *QTreeView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QTreeView{bcthis0}
}

// /usr/include/qt/QtWidgets/qtreeview.h:174
// index:1
// void QTreeView(class QTreeViewPrivate &, class QWidget *)
func NewQTreeView_1(dd unsafe.Pointer, parent unsafe.Pointer) *QTreeView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeViewC2ER16QTreeViewPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeViewFromPointer(cthis)
	return gothis
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
	// 0: (, model QAbstractItemModel *), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:75
// index:0
// virtual
// void setRootIndex(const class QModelIndex &)
func (this *QTreeView) SetRootIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:76
// index:0
// virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QTreeView) SetSelectionModel(selectionModel unsafe.Pointer) {
	// 0: (, selectionModel QItemSelectionModel *), (selectionModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:78
// index:0
// QHeaderView * header()
func (this *QTreeView) Header() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView6headerEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:79
// index:0
// void setHeader(class QHeaderView *)
func (this *QTreeView) SetHeader(header unsafe.Pointer) {
	// 0: (, header QHeaderView *), (header)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9setHeaderEP11QHeaderView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), header)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:81
// index:0
// int autoExpandDelay()
func (this *QTreeView) AutoExpandDelay() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15autoExpandDelayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:82
// index:0
// void setAutoExpandDelay(int)
func (this *QTreeView) SetAutoExpandDelay(delay int) {
	// 0: (, delay int), (&delay)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setAutoExpandDelayEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &delay)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:84
// index:0
// int indentation()
func (this *QTreeView) Indentation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11indentationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:85
// index:0
// void setIndentation(int)
func (this *QTreeView) SetIndentation(i int) {
	// 0: (, i int), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14setIndentationEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:86
// index:0
// void resetIndentation()
func (this *QTreeView) ResetIndentation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16resetIndentationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:88
// index:0
// bool rootIsDecorated()
func (this *QTreeView) RootIsDecorated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15rootIsDecoratedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:89
// index:0
// void setRootIsDecorated(_Bool)
func (this *QTreeView) SetRootIsDecorated(show bool) {
	// 0: (, show bool), (&show)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setRootIsDecoratedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:91
// index:0
// bool uniformRowHeights()
func (this *QTreeView) UniformRowHeights() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView17uniformRowHeightsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:92
// index:0
// void setUniformRowHeights(_Bool)
func (this *QTreeView) SetUniformRowHeights(uniform bool) {
	// 0: (, uniform bool), (&uniform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView20setUniformRowHeightsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &uniform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:94
// index:0
// bool itemsExpandable()
func (this *QTreeView) ItemsExpandable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15itemsExpandableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:95
// index:0
// void setItemsExpandable(_Bool)
func (this *QTreeView) SetItemsExpandable(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18setItemsExpandableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:97
// index:0
// bool expandsOnDoubleClick()
func (this *QTreeView) ExpandsOnDoubleClick() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView20expandsOnDoubleClickEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:98
// index:0
// void setExpandsOnDoubleClick(_Bool)
func (this *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView23setExpandsOnDoubleClickEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:100
// index:0
// int columnViewportPosition(int)
func (this *QTreeView) ColumnViewportPosition(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView22columnViewportPositionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:101
// index:0
// int columnWidth(int)
func (this *QTreeView) ColumnWidth(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11columnWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:102
// index:0
// void setColumnWidth(int, int)
func (this *QTreeView) SetColumnWidth(column int, width int) {
	// 0: (, column int, width int), (&column, &width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14setColumnWidthEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:103
// index:0
// int columnAt(int)
func (this *QTreeView) ColumnAt(x int) {
	// 0: (, x int), (&x)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8columnAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:105
// index:0
// bool isColumnHidden(int)
func (this *QTreeView) IsColumnHidden(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14isColumnHiddenEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:106
// index:0
// void setColumnHidden(int, _Bool)
func (this *QTreeView) SetColumnHidden(column int, hide bool) {
	// 0: (, column int, hide bool), (&column, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setColumnHiddenEib", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:108
// index:0
// bool isHeaderHidden()
func (this *QTreeView) IsHeaderHidden() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14isHeaderHiddenEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:109
// index:0
// void setHeaderHidden(_Bool)
func (this *QTreeView) SetHeaderHidden(hide bool) {
	// 0: (, hide bool), (&hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setHeaderHiddenEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:111
// index:0
// bool isRowHidden(int, const class QModelIndex &)
func (this *QTreeView) IsRowHidden(row int, parent unsafe.Pointer) {
	// 0: (, row int, parent const QModelIndex &), (&row, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView11isRowHiddenEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:112
// index:0
// void setRowHidden(int, const class QModelIndex &, _Bool)
func (this *QTreeView) SetRowHidden(row int, parent unsafe.Pointer, hide bool) {
	// 0: (, row int, parent const QModelIndex &, hide bool), (&row, parent, &hide)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12setRowHiddenEiRK11QModelIndexb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, parent, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:114
// index:0
// bool isFirstColumnSpanned(int, const class QModelIndex &)
func (this *QTreeView) IsFirstColumnSpanned(row int, parent unsafe.Pointer) {
	// 0: (, row int, parent const QModelIndex &), (&row, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView20isFirstColumnSpannedEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:115
// index:0
// void setFirstColumnSpanned(int, const class QModelIndex &, _Bool)
func (this *QTreeView) SetFirstColumnSpanned(row int, parent unsafe.Pointer, span bool) {
	// 0: (, row int, parent const QModelIndex &, span bool), (&row, parent, &span)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView21setFirstColumnSpannedEiRK11QModelIndexb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, parent, &span)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:117
// index:0
// bool isExpanded(const class QModelIndex &)
func (this *QTreeView) IsExpanded(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10isExpandedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:118
// index:0
// void setExpanded(const class QModelIndex &, _Bool)
func (this *QTreeView) SetExpanded(index unsafe.Pointer, expand bool) {
	// 0: (, index const QModelIndex &, expand bool), (index, &expand)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setExpandedERK11QModelIndexb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &expand)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:120
// index:0
// void setSortingEnabled(_Bool)
func (this *QTreeView) SetSortingEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17setSortingEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:121
// index:0
// bool isSortingEnabled()
func (this *QTreeView) IsSortingEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16isSortingEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:123
// index:0
// void setAnimated(_Bool)
func (this *QTreeView) SetAnimated(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setAnimatedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:124
// index:0
// bool isAnimated()
func (this *QTreeView) IsAnimated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10isAnimatedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:126
// index:0
// void setAllColumnsShowFocus(_Bool)
func (this *QTreeView) SetAllColumnsShowFocus(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView22setAllColumnsShowFocusEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:127
// index:0
// bool allColumnsShowFocus()
func (this *QTreeView) AllColumnsShowFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView19allColumnsShowFocusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:129
// index:0
// void setWordWrap(_Bool)
func (this *QTreeView) SetWordWrap(on bool) {
	// 0: (, on bool), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11setWordWrapEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:130
// index:0
// bool wordWrap()
func (this *QTreeView) WordWrap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8wordWrapEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:132
// index:0
// void setTreePosition(int)
func (this *QTreeView) SetTreePosition(logicalIndex int) {
	// 0: (, logicalIndex int), (&logicalIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15setTreePositionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:133
// index:0
// int treePosition()
func (this *QTreeView) TreePosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView12treePositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:135
// index:0
// virtual
// void keyboardSearch(const class QString &)
func (this *QTreeView) KeyboardSearch(search unsafe.Pointer) {
	// 0: (, search const QString &), (search)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14keyboardSearchERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), search)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:137
// index:0
// virtual
// QRect visualRect(const class QModelIndex &)
func (this *QTreeView) VisualRect(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:138
// index:0
// virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QTreeView) ScrollTo(index unsafe.Pointer, hint int) {
	// 0: (, index const QModelIndex &, hint QAbstractItemView::ScrollHint), (index, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:139
// index:0
// virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QTreeView) IndexAt(p unsafe.Pointer) {
	// 0: (, p const QPoint &), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView7indexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:140
// index:0
// QModelIndex indexAbove(const class QModelIndex &)
func (this *QTreeView) IndexAbove(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10indexAboveERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:141
// index:0
// QModelIndex indexBelow(const class QModelIndex &)
func (this *QTreeView) IndexBelow(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView10indexBelowERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:143
// index:0
// virtual
// void doItemsLayout()
func (this *QTreeView) DoItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13doItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:144
// index:0
// virtual
// void reset()
func (this *QTreeView) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView5resetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:146
// index:0
// void sortByColumn(int, Qt::SortOrder)
func (this *QTreeView) SortByColumn(column int, order int) {
	// 0: (, column int, order Qt::SortOrder), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:161
// index:1
// void sortByColumn(int)
func (this *QTreeView) SortByColumn_1(column int) {
	// 1: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12sortByColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:149
// index:0
// virtual
// void selectAll()
func (this *QTreeView) SelectAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9selectAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:152
// index:0
// void expanded(const class QModelIndex &)
func (this *QTreeView) Expanded(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8expandedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:153
// index:0
// void collapsed(const class QModelIndex &)
func (this *QTreeView) Collapsed(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9collapsedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:156
// index:0
// void hideColumn(int)
func (this *QTreeView) HideColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10hideColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:157
// index:0
// void showColumn(int)
func (this *QTreeView) ShowColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10showColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:158
// index:0
// void expand(const class QModelIndex &)
func (this *QTreeView) Expand(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView6expandERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:159
// index:0
// void collapse(const class QModelIndex &)
func (this *QTreeView) Collapse(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8collapseERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:160
// index:0
// void resizeColumnToContents(int)
func (this *QTreeView) ResizeColumnToContents(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView22resizeColumnToContentsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:162
// index:0
// void expandAll()
func (this *QTreeView) ExpandAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView9expandAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:163
// index:0
// void collapseAll()
func (this *QTreeView) CollapseAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11collapseAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:164
// index:0
// void expandToDepth(int)
func (this *QTreeView) ExpandToDepth(depth int) {
	// 0: (, depth int), (&depth)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13expandToDepthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &depth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:167
// index:0
// void columnResized(int, int, int)
func (this *QTreeView) ColumnResized(column int, oldSize int, newSize int) {
	// 0: (, column int, oldSize int, newSize int), (&column, &oldSize, &newSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13columnResizedEiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &oldSize, &newSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:168
// index:0
// void columnCountChanged(int, int)
func (this *QTreeView) ColumnCountChanged(oldCount int, newCount int) {
	// 0: (, oldCount int, newCount int), (&oldCount, &newCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView18columnCountChangedEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &oldCount, &newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:169
// index:0
// void columnMoved()
func (this *QTreeView) ColumnMoved() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11columnMovedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:170
// index:0
// void reexpand()
func (this *QTreeView) Reexpand() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView8reexpandEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:171
// index:0
// void rowsRemoved(const class QModelIndex &, int, int)
func (this *QTreeView) RowsRemoved(parent unsafe.Pointer, first int, last int) {
	// 0: (, parent const QModelIndex &, first int, last int), (parent, &first, &last)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView11rowsRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &first, &last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:175
// index:0
// virtual
// void scrollContentsBy(int, int)
func (this *QTreeView) ScrollContentsBy(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16scrollContentsByEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:176
// index:0
// virtual
// void rowsInserted(const class QModelIndex &, int, int)
func (this *QTreeView) RowsInserted(parent unsafe.Pointer, start int, end int) {
	// 0: (, parent const QModelIndex &, start int, end int), (parent, &start, &end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:177
// index:0
// virtual
// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QTreeView) RowsAboutToBeRemoved(parent unsafe.Pointer, start int, end int) {
	// 0: (, parent const QModelIndex &, start int, end int), (parent, &start, &end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView20rowsAboutToBeRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:179
// index:0
// virtual
// QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QTreeView) MoveCursor(cursorAction int, modifiers int) {
	// 0: (, cursorAction QAbstractItemView::CursorAction, QFlags<Qt::KeyboardModifier> modifiers), (&cursorAction, &modifiers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cursorAction, &modifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:180
// index:0
// virtual
// int horizontalOffset()
func (this *QTreeView) HorizontalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16horizontalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:181
// index:0
// virtual
// int verticalOffset()
func (this *QTreeView) VerticalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView14verticalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:183
// index:0
// virtual
// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QTreeView) SetSelection(rect unsafe.Pointer, command int) {
	// 0: (, rect const QRect &, QFlags<QItemSelectionModel::SelectionFlag> command), (rect, &command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:184
// index:0
// virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QTreeView) VisualRegionForSelection(selection unsafe.Pointer) {
	// 0: (, selection const QItemSelection &), (selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:185
// index:0
// virtual
// QModelIndexList selectedIndexes()
func (this *QTreeView) SelectedIndexes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView15selectedIndexesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:187
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QTreeView) TimerEvent(event unsafe.Pointer) {
	// 0: (, event QTimerEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:188
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QTreeView) PaintEvent(event unsafe.Pointer) {
	// 0: (, event QPaintEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:190
// index:0
// void drawTree(class QPainter *, const class QRegion &)
func (this *QTreeView) DrawTree(painter unsafe.Pointer, region unsafe.Pointer) {
	// 0: (, painter QPainter *, region const QRegion &), (painter, region)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView8drawTreeEP8QPainterRK7QRegion", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, region)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:191
// index:0
// virtual
// void drawRow(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
func (this *QTreeView) DrawRow(painter unsafe.Pointer, options unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, painter QPainter *, options const QStyleOptionViewItem &, index const QModelIndex &), (painter, options, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView7drawRowEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, options, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:194
// index:0
// virtual
// void drawBranches(class QPainter *, const class QRect &, const class QModelIndex &)
func (this *QTreeView) DrawBranches(painter unsafe.Pointer, rect unsafe.Pointer, index unsafe.Pointer) {
	// 0: (, painter QPainter *, rect const QRect &, index const QModelIndex &), (painter, rect, index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView12drawBranchesEP8QPainterRK5QRectRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:198
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QTreeView) MousePressEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:199
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QTreeView) MouseReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:200
// index:0
// virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QTreeView) MouseDoubleClickEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:201
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QTreeView) MouseMoveEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:202
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QTreeView) KeyPressEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:204
// index:0
// virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QTreeView) DragMoveEvent(event unsafe.Pointer) {
	// 0: (, event QDragMoveEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:206
// index:0
// virtual
// bool viewportEvent(class QEvent *)
func (this *QTreeView) ViewportEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView13viewportEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:208
// index:0
// virtual
// void updateGeometries()
func (this *QTreeView) UpdateGeometries() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16updateGeometriesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:210
// index:0
// virtual
// QSize viewportSizeHint()
func (this *QTreeView) ViewportSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16viewportSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:212
// index:0
// virtual
// int sizeHintForColumn(int)
func (this *QTreeView) SizeHintForColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView17sizeHintForColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:213
// index:0
// int indexRowSizeHint(const class QModelIndex &)
func (this *QTreeView) IndexRowSizeHint(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView16indexRowSizeHintERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:214
// index:0
// int rowHeight(const class QModelIndex &)
func (this *QTreeView) RowHeight(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView9rowHeightERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:216
// index:0
// virtual
// void horizontalScrollbarAction(int)
func (this *QTreeView) HorizontalScrollbarAction(action int) {
	// 0: (, action int), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView25horizontalScrollbarActionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:218
// index:0
// virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QTreeView) IsIndexHidden(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTreeView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:219
// index:0
// virtual
// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QTreeView) SelectionChanged(selected unsafe.Pointer, deselected unsafe.Pointer) {
	// 0: (, selected const QItemSelection &, deselected const QItemSelection &), (selected, deselected)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selected, deselected)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreeview.h:221
// index:0
// virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QTreeView) CurrentChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, current const QModelIndex &, previous const QModelIndex &), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTreeView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), current, previous)
	gopp.ErrPrint(err, rv)
}

//  body block end
