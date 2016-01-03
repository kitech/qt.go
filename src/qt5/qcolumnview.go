package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qcolumnview.h
// dst-file: /src/widgets/qcolumnview.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QColumnView::QColumnView(QWidget * parent);
extern void* dector_ZN11QColumnViewC1EP7QWidget(void* arg0);
extern void _ZN11QColumnViewC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QColumnView::selectAll();
extern void _ZN11QColumnView9selectAllEv(void* qthis);
  // proto:  void QColumnView::setPreviewWidget(QWidget * widget);
extern void _ZN11QColumnView16setPreviewWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  QModelIndex QColumnView::indexAt(const QPoint & point);
extern void _ZNK11QColumnView7indexAtERK6QPoint(void* qthis, void* arg0);
  // proto:  const QMetaObject * QColumnView::metaObject();
extern void _ZNK11QColumnView10metaObjectEv(void* qthis);
  // proto:  QSize QColumnView::sizeHint();
extern void _ZNK11QColumnView8sizeHintEv(void* qthis);
  // proto:  QList<int> QColumnView::columnWidths();
extern void _ZNK11QColumnView12columnWidthsEv(void* qthis);
  // proto:  void QColumnView::setResizeGripsVisible(bool visible);
extern void _ZN11QColumnView21setResizeGripsVisibleEb(void* qthis, bool arg0);
  // proto:  bool QColumnView::resizeGripsVisible();
extern void _ZNK11QColumnView18resizeGripsVisibleEv(void* qthis);
  // proto:  void QColumnView::QColumnView(const QColumnView & );
extern void* dector_ZN11QColumnViewC1ERKS_(void* arg0);
extern void _ZN11QColumnViewC1ERKS_(void* qthis, void* arg0);
  // proto:  void QColumnView::setModel(QAbstractItemModel * model);
extern void _ZN11QColumnView8setModelEP18QAbstractItemModel(void* qthis, void* arg0);
  // proto:  void QColumnView::setRootIndex(const QModelIndex & index);
extern void _ZN11QColumnView12setRootIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QWidget * QColumnView::previewWidget();
extern void _ZNK11QColumnView13previewWidgetEv(void* qthis);
  // proto:  void QColumnView::setSelectionModel(QItemSelectionModel * selectionModel);
extern void _ZN11QColumnView17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0);
  // proto:  QRect QColumnView::visualRect(const QModelIndex & index);
extern void _ZNK11QColumnView10visualRectERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QColumnView::~QColumnView();
extern void _ZN11QColumnViewD0Ev(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QColumnView)=1
type QColumnView struct {
  /*qbase*/ QAbstractItemView;
  qclsinst unsafe.Pointer /* *C.void */;
//  _updatePreviewWidget QColumnView_updatePreviewWidget_signal;
}

  // proto:  void QColumnView::QColumnView(QWidget * parent);
func NewQColumnView(args ...interface{}) QColumnView {
  return QColumnView{}
}

  // proto:  void QColumnView::selectAll();
func (this *QColumnView) selectAll(args ...interface{}) () {
  // selectAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnView9selectAllEv
    // invoke: void selectAll()
    C._ZN11QColumnView9selectAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "selectAll", args)
  }

}

  // proto:  void QColumnView::setPreviewWidget(QWidget * widget);
func (this *QColumnView) setPreviewWidget(args ...interface{}) () {
  // setPreviewWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnView16setPreviewWidgetEP7QWidget
    // invoke: void setPreviewWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QColumnView16setPreviewWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setPreviewWidget", args)
  }

}

  // proto:  QModelIndex QColumnView::indexAt(const QPoint & point);
func (this *QColumnView) indexAt(args ...interface{}) () {
  // indexAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QColumnView7indexAtERK6QPoint
    // invoke: QModelIndex indexAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK11QColumnView7indexAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "indexAt", args)
  }

}

  // proto:  const QMetaObject * QColumnView::metaObject();
func (this *QColumnView) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QColumnView10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QColumnView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "metaObject", args)
  }

}

  // proto:  QSize QColumnView::sizeHint();
func (this *QColumnView) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QColumnView8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK11QColumnView8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "sizeHint", args)
  }

}

  // proto:  QList<int> QColumnView::columnWidths();
func (this *QColumnView) columnWidths(args ...interface{}) () {
  // columnWidths()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QColumnView12columnWidthsEv
    // invoke: QList<int> columnWidths()
    C._ZNK11QColumnView12columnWidthsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "columnWidths", args)
  }

}

  // proto:  void QColumnView::setResizeGripsVisible(bool visible);
func (this *QColumnView) setResizeGripsVisible(args ...interface{}) () {
  // setResizeGripsVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnView21setResizeGripsVisibleEb
    // invoke: void setResizeGripsVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QColumnView21setResizeGripsVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setResizeGripsVisible", args)
  }

}

  // proto:  bool QColumnView::resizeGripsVisible();
func (this *QColumnView) resizeGripsVisible(args ...interface{}) () {
  // resizeGripsVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QColumnView18resizeGripsVisibleEv
    // invoke: bool resizeGripsVisible()
    C._ZNK11QColumnView18resizeGripsVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "resizeGripsVisible", args)
  }

}

  // proto:  void QColumnView::setModel(QAbstractItemModel * model);
func (this *QColumnView) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnView8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QColumnView8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setModel", args)
  }

}

  // proto:  void QColumnView::setRootIndex(const QModelIndex & index);
func (this *QColumnView) setRootIndex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnView12setRootIndexERK11QModelIndex
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QColumnView12setRootIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setRootIndex", args)
  }

}

  // proto:  QWidget * QColumnView::previewWidget();
func (this *QColumnView) previewWidget(args ...interface{}) () {
  // previewWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QColumnView13previewWidgetEv
    // invoke: QWidget * previewWidget()
    C._ZNK11QColumnView13previewWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "previewWidget", args)
  }

}

  // proto:  void QColumnView::setSelectionModel(QItemSelectionModel * selectionModel);
func (this *QColumnView) setSelectionModel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnView17setSelectionModelEP19QItemSelectionModel
    // invoke: void setSelectionModel(class QItemSelectionModel *)
    var arg0 = args[0].(QItemSelectionModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QColumnView17setSelectionModelEP19QItemSelectionModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setSelectionModel", args)
  }

}

  // proto:  QRect QColumnView::visualRect(const QModelIndex & index);
func (this *QColumnView) visualRect(args ...interface{}) () {
  // visualRect(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QColumnView10visualRectERK11QModelIndex
    // invoke: QRect visualRect(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK11QColumnView10visualRectERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "visualRect", args)
  }

}

  // proto:  void QColumnView::~QColumnView();
func (this *QColumnView) FreeQColumnView(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColumnView", "~QColumnView", args)
  }

}

// <= body block end

