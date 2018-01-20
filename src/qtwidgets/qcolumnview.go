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
func NewQColumnViewFromPointer(cthis unsafe.Pointer) *QColumnView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QColumnView{bcthis0}
}

// /usr/include/qt/QtWidgets/qcolumnview.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QColumnView) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:58
// index:0
// Public
// void updatePreviewWidget(const class QModelIndex &)
func (this *QColumnView) UpdatePreviewWidget(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView19updatePreviewWidgetERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:61
// index:0
// Public
// void QColumnView(class QWidget *)
func NewQColumnView(parent unsafe.Pointer) *QColumnView {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQColumnViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcolumnview.h:62
// index:0
// Public virtual
// void ~QColumnView()
func DeleteQColumnView(*QColumnView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:65
// index:0
// Public virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QColumnView) IndexAt(point *qtcore.QPoint) interface{} {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView7indexAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:66
// index:0
// Public virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QColumnView) ScrollTo(index *qtcore.QModelIndex, hint int) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:67
// index:0
// Public virtual
// QSize sizeHint()
func (this *QColumnView) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:68
// index:0
// Public virtual
// QRect visualRect(const class QModelIndex &)
func (this *QColumnView) VisualRect(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:69
// index:0
// Public virtual
// void setModel(class QAbstractItemModel *)
func (this *QColumnView) SetModel(model unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:70
// index:0
// Public virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QColumnView) SetSelectionModel(selectionModel unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:71
// index:0
// Public virtual
// void setRootIndex(const class QModelIndex &)
func (this *QColumnView) SetRootIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:72
// index:0
// Public virtual
// void selectAll()
func (this *QColumnView) SelectAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView9selectAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:75
// index:0
// Public
// void setResizeGripsVisible(_Bool)
func (this *QColumnView) SetResizeGripsVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView21setResizeGripsVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:76
// index:0
// Public
// bool resizeGripsVisible()
func (this *QColumnView) ResizeGripsVisible() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView18resizeGripsVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:78
// index:0
// Public
// QWidget * previewWidget()
func (this *QColumnView) PreviewWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView13previewWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:79
// index:0
// Public
// void setPreviewWidget(class QWidget *)
func (this *QColumnView) SetPreviewWidget(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView16setPreviewWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:82
// index:0
// Public
// QList<int> columnWidths()
func (this *QColumnView) ColumnWidths() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView12columnWidthsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:88
// index:0
// Protected virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QColumnView) IsIndexHidden(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:90
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QColumnView) ResizeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:92
// index:0
// Protected virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QColumnView) VisualRegionForSelection(selection *qtcore.QItemSelection) interface{} {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:93
// index:0
// Protected virtual
// int horizontalOffset()
func (this *QColumnView) HorizontalOffset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView16horizontalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:94
// index:0
// Protected virtual
// int verticalOffset()
func (this *QColumnView) VerticalOffset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView14verticalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:95
// index:0
// Protected virtual
// void rowsInserted(const class QModelIndex &, int, int)
func (this *QColumnView) RowsInserted(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:96
// index:0
// Protected virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QColumnView) CurrentChanged(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:99
// index:0
// Protected virtual
// void scrollContentsBy(int, int)
func (this *QColumnView) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:100
// index:0
// Protected virtual
// QAbstractItemView * createColumn(const class QModelIndex &)
func (this *QColumnView) CreateColumn(rootIndex *qtcore.QModelIndex) interface{} {
	var convArg0 = rootIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView12createColumnERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcolumnview.h:101
// index:0
// Protected
// void initializeColumn(class QAbstractItemView *)
func (this *QColumnView) InitializeColumn(column unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView16initializeColumnEP17QAbstractItemView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

//  body block end
