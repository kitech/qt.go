package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QColumnView::selectAll();
extern void C_ZN11QColumnView9selectAllEv(void* qthis); // 4
  // proto:  void QColumnView::setResizeGripsVisible(bool visible);
extern void C_ZN11QColumnView21setResizeGripsVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QColumnView::~QColumnView();
extern void C_ZN11QColumnViewD2Ev(void* qthis); // 4
  // proto:  bool QColumnView::resizeGripsVisible();
extern void C_ZNK11QColumnView18resizeGripsVisibleEv(void* qthis); // 4
  // proto:  QModelIndex QColumnView::indexAt(const QPoint & point);
extern void C_ZNK11QColumnView7indexAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QWidget * QColumnView::previewWidget();
extern void C_ZNK11QColumnView13previewWidgetEv(void* qthis); // 4
  // proto:  QRect QColumnView::visualRect(const QModelIndex & index);
extern void C_ZNK11QColumnView10visualRectERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QColumnView::setPreviewWidget(QWidget * widget);
extern void C_ZN11QColumnView16setPreviewWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QColumnView::setRootIndex(const QModelIndex & index);
extern void C_ZN11QColumnView12setRootIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QSize QColumnView::sizeHint();
extern void C_ZNK11QColumnView8sizeHintEv(void* qthis); // 4
  // proto:  void QColumnView::setModel(QAbstractItemModel * model);
extern void C_ZN11QColumnView8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  void QColumnView::setSelectionModel(QItemSelectionModel * selectionModel);
extern void C_ZN11QColumnView17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0); // 4
  // proto:  QList<int> QColumnView::columnWidths();
extern void C_ZNK11QColumnView12columnWidthsEv(void* qthis); // 4
  // proto:  const QMetaObject * QColumnView::metaObject();
extern void C_ZNK11QColumnView10metaObjectEv(void* qthis); // 4
  // proto:  void QColumnView::QColumnView(QWidget * parent);
extern void* C_ZN11QColumnViewC2EP7QWidget(void* arg0); // 3
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

// selectAll()
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
    C.C_ZN11QColumnView9selectAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "selectAll", args)
  }

}

// setResizeGripsVisible(_Bool)
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
    C.C_ZN11QColumnView21setResizeGripsVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setResizeGripsVisible", args)
  }

}

// ~QColumnView()
func (this *QColumnView) FreeQColumnView(args ...interface{}) () {
  // ~QColumnView()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnViewD0Ev
    // invoke: void ~QColumnView()
    C.C_ZN11QColumnViewD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "~QColumnView", args)
  }

}

// resizeGripsVisible()
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
    var ret = C.C_ZNK11QColumnView18resizeGripsVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColumnView", "resizeGripsVisible", args)
  }

}

// indexAt(const class QPoint &)
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
    var ret = C.C_ZNK11QColumnView7indexAtERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColumnView", "indexAt", args)
  }

}

// previewWidget()
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
    var ret = C.C_ZNK11QColumnView13previewWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColumnView", "previewWidget", args)
  }

}

// visualRect(const class QModelIndex &)
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
    var ret = C.C_ZNK11QColumnView10visualRectERK11QModelIndex(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColumnView", "visualRect", args)
  }

}

// setPreviewWidget(class QWidget *)
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
    C.C_ZN11QColumnView16setPreviewWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setPreviewWidget", args)
  }

}

// setRootIndex(const class QModelIndex &)
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
    C.C_ZN11QColumnView12setRootIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setRootIndex", args)
  }

}

// sizeHint()
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
    var ret = C.C_ZNK11QColumnView8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColumnView", "sizeHint", args)
  }

}

// setModel(class QAbstractItemModel *)
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
    C.C_ZN11QColumnView8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setModel", args)
  }

}

// setSelectionModel(class QItemSelectionModel *)
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
    C.C_ZN11QColumnView17setSelectionModelEP19QItemSelectionModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setSelectionModel", args)
  }

}

// columnWidths()
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
    C.C_ZNK11QColumnView12columnWidthsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "columnWidths", args)
  }

}

// metaObject()
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
    C.C_ZNK11QColumnView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "metaObject", args)
  }

}

// QColumnView(class QWidget *)
func NewQColumnView(args ...interface{}) *QColumnView {
  // QColumnView(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnViewC1EP7QWidget
    // invoke: void QColumnView(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QColumnViewC2EP7QWidget(arg0)
    return &QColumnView{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QColumnView", "QColumnView", args)
  }

  return nil // QColumnView{qclsinst:qthis}
}

// <= body block end

