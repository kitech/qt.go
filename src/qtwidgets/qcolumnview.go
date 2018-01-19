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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qcolumnview.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QColumnView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:58
// index:0
// void updatePreviewWidget(const class QModelIndex &)
func (this *QColumnView) UpdatePreviewWidget(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView19updatePreviewWidgetERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:61
// index:0
// void QColumnView(class QWidget *)
func NewQColumnView(parent unsafe.Pointer) *QColumnView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnViewC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QColumnView{cthis}
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
	// 0: (, const QPoint & point), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView7indexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:66
// index:0
// virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QColumnView) ScrollTo(index unsafe.Pointer, hint int) {
	// 0: (, const QModelIndex & index, QAbstractItemView::ScrollHint hint), (index, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.cthis, index, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:67
// index:0
// virtual
// QSize sizeHint()
func (this *QColumnView) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:68
// index:0
// virtual
// QRect visualRect(const class QModelIndex &)
func (this *QColumnView) VisualRect(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:69
// index:0
// virtual
// void setModel(class QAbstractItemModel *)
func (this *QColumnView) SetModel(model unsafe.Pointer) {
	// 0: (, QAbstractItemModel * model), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.cthis, model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:70
// index:0
// virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QColumnView) SetSelectionModel(selectionModel unsafe.Pointer) {
	// 0: (, QItemSelectionModel * selectionModel), (selectionModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_VOID, this.cthis, selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:71
// index:0
// virtual
// void setRootIndex(const class QModelIndex &)
func (this *QColumnView) SetRootIndex(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:72
// index:0
// virtual
// void selectAll()
func (this *QColumnView) SelectAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView9selectAllEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:75
// index:0
// void setResizeGripsVisible(_Bool)
func (this *QColumnView) SetResizeGripsVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView21setResizeGripsVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:76
// index:0
// bool resizeGripsVisible()
func (this *QColumnView) ResizeGripsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView18resizeGripsVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:78
// index:0
// QWidget * previewWidget()
func (this *QColumnView) PreviewWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView13previewWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:79
// index:0
// void setPreviewWidget(class QWidget *)
func (this *QColumnView) SetPreviewWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QColumnView16setPreviewWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolumnview.h:82
// index:0
// QList<int> columnWidths()
func (this *QColumnView) ColumnWidths() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QColumnView12columnWidthsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
