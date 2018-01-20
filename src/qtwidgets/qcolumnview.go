//  header block begin
// /usr/include/qt/QtWidgets/qcolumnview.h
// #include <qcolumnview.h>
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
type QColumnView struct {
	*QAbstractItemView
}

func (this *QColumnView) GetCthis() unsafe.Pointer {
	return this.QAbstractItemView.GetCthis()
}

// /usr/include/qt/QtWidgets/qcolumnview.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QColumnView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:58
// index:0
// void updatePreviewWidget(const class QModelIndex &)
func (this *QColumnView) UpdatePreviewWidget(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView19updatePreviewWidgetERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:61
// index:0
// void QColumnView(class QWidget *)
func NewQColumnView(parent unsafe.Pointer) *QColumnView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQColumnViewFromPointer(cthis)
	return gothis
}
func NewQColumnViewFromPointer(cthis unsafe.Pointer) *QColumnView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QColumnView{bcthis0}
}

// /usr/include/qt/QtWidgets/qcolumnview.h:85
// index:1
// void QColumnView(class QColumnViewPrivate &, class QWidget *)
func NewQColumnView_1(dd unsafe.Pointer, parent unsafe.Pointer) *QColumnView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnViewC2ER18QColumnViewPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQColumnViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcolumnview.h:62
// index:0
// virtual
// void ~QColumnView()
func DeleteQColumnView(*QColumnView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:65
// index:0
// virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QColumnView) IndexAt(point unsafe.Pointer) {
	// 0: (, point const QPoint &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView7indexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:66
// index:0
// virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QColumnView) ScrollTo(index unsafe.Pointer, hint int) {
	// 0: (, index const QModelIndex &, hint QAbstractItemView::ScrollHint), (index, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:67
// index:0
// virtual
// QSize sizeHint()
func (this *QColumnView) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:68
// index:0
// virtual
// QRect visualRect(const class QModelIndex &)
func (this *QColumnView) VisualRect(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:69
// index:0
// virtual
// void setModel(class QAbstractItemModel *)
func (this *QColumnView) SetModel(model unsafe.Pointer) {
	// 0: (, model QAbstractItemModel *), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:70
// index:0
// virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QColumnView) SetSelectionModel(selectionModel unsafe.Pointer) {
	// 0: (, selectionModel QItemSelectionModel *), (selectionModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:71
// index:0
// virtual
// void setRootIndex(const class QModelIndex &)
func (this *QColumnView) SetRootIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:72
// index:0
// virtual
// void selectAll()
func (this *QColumnView) SelectAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView9selectAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:75
// index:0
// void setResizeGripsVisible(_Bool)
func (this *QColumnView) SetResizeGripsVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView21setResizeGripsVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:76
// index:0
// bool resizeGripsVisible()
func (this *QColumnView) ResizeGripsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView18resizeGripsVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:78
// index:0
// QWidget * previewWidget()
func (this *QColumnView) PreviewWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView13previewWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:79
// index:0
// void setPreviewWidget(class QWidget *)
func (this *QColumnView) SetPreviewWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView16setPreviewWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:82
// index:0
// QList<int> columnWidths()
func (this *QColumnView) ColumnWidths() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView12columnWidthsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:88
// index:0
// virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QColumnView) IsIndexHidden(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:89
// index:0
// virtual
// QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QColumnView) MoveCursor(cursorAction int, modifiers int) {
	// 0: (, cursorAction QAbstractItemView::CursorAction, QFlags<Qt::KeyboardModifier> modifiers), (&cursorAction, &modifiers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView10moveCursorEN17QAbstractItemView12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cursorAction, &modifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:90
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QColumnView) ResizeEvent(event unsafe.Pointer) {
	// 0: (, event QResizeEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:91
// index:0
// virtual
// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QColumnView) SetSelection(rect unsafe.Pointer, command int) {
	// 0: (, rect const QRect &, QFlags<QItemSelectionModel::SelectionFlag> command), (rect, &command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:92
// index:0
// virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QColumnView) VisualRegionForSelection(selection unsafe.Pointer) {
	// 0: (, selection const QItemSelection &), (selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:93
// index:0
// virtual
// int horizontalOffset()
func (this *QColumnView) HorizontalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView16horizontalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:94
// index:0
// virtual
// int verticalOffset()
func (this *QColumnView) VerticalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView14verticalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:95
// index:0
// virtual
// void rowsInserted(const class QModelIndex &, int, int)
func (this *QColumnView) RowsInserted(parent unsafe.Pointer, start int, end int) {
	// 0: (, parent const QModelIndex &, start int, end int), (parent, &start, &end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:96
// index:0
// virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QColumnView) CurrentChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, current const QModelIndex &, previous const QModelIndex &), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:99
// index:0
// virtual
// void scrollContentsBy(int, int)
func (this *QColumnView) ScrollContentsBy(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView16scrollContentsByEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:100
// index:0
// virtual
// QAbstractItemView * createColumn(const class QModelIndex &)
func (this *QColumnView) CreateColumn(rootIndex unsafe.Pointer) {
	// 0: (, rootIndex const QModelIndex &), (rootIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView12createColumnERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rootIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:101
// index:0
// void initializeColumn(class QAbstractItemView *)
func (this *QColumnView) InitializeColumn(column unsafe.Pointer) {
	// 0: (, column QAbstractItemView *), (column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView16initializeColumnEP17QAbstractItemView", ffiqt.FFI_TYPE_VOID, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

//  body block end
