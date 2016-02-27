package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
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
extern bool C_ZNK11QColumnView18resizeGripsVisibleEv(void* qthis); // 4
  // proto:  QModelIndex QColumnView::indexAt(const QPoint & point);
extern void* C_ZNK11QColumnView7indexAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QWidget * QColumnView::previewWidget();
extern void* C_ZNK11QColumnView13previewWidgetEv(void* qthis); // 4
  // proto:  QRect QColumnView::visualRect(const QModelIndex & index);
extern void* C_ZNK11QColumnView10visualRectERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QColumnView::setPreviewWidget(QWidget * widget);
extern void C_ZN11QColumnView16setPreviewWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QColumnView::setRootIndex(const QModelIndex & index);
extern void C_ZN11QColumnView12setRootIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QSize QColumnView::sizeHint();
extern void* C_ZNK11QColumnView8sizeHintEv(void* qthis); // 4
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
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QColumnView)=1
type QColumnView struct {
  /*qbase*/ QAbstractItemView;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _updatePreviewWidget QColumnView_updatePreviewWidget_signal;
}

// selectAll()
func (this *QColumnView) SelectAll(args ...interface{}) () {
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
    C.C_ZN11QColumnView9selectAllEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "selectAll", args)
  }

  return
}

// setResizeGripsVisible(_Bool)
func (this *QColumnView) SetResizeGripsVisible(args ...interface{}) () {
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
    C.C_ZN11QColumnView21setResizeGripsVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setResizeGripsVisible", args)
  }

  return
}

// ~QColumnView()
func (this *QColumnView) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QColumnViewD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QColumnView", "~QColumnView", args)
  }

  return
}

// resizeGripsVisible()
func (this *QColumnView) ResizeGripsVisible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QColumnView18resizeGripsVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColumnView", "resizeGripsVisible", args)
  }

  return
}

// indexAt(const class QPoint &)
func (this *QColumnView) IndexAt(args ...interface{}) (ret interface{}) {
  // indexAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QColumnView7indexAtERK6QPoint
    // invoke: QModelIndex indexAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QColumnView7indexAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColumnView", "indexAt", args)
  }

  return
}

// previewWidget()
func (this *QColumnView) PreviewWidget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QColumnView13previewWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColumnView", "previewWidget", args)
  }

  return
}

// visualRect(const class QModelIndex &)
func (this *QColumnView) VisualRect(args ...interface{}) (ret interface{}) {
  // visualRect(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QColumnView10visualRectERK11QModelIndex
    // invoke: QRect visualRect(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QColumnView10visualRectERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColumnView", "visualRect", args)
  }

  return
}

// setPreviewWidget(class QWidget *)
func (this *QColumnView) SetPreviewWidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QColumnView16setPreviewWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setPreviewWidget", args)
  }

  return
}

// setRootIndex(const class QModelIndex &)
func (this *QColumnView) SetRootIndex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnView12setRootIndexERK11QModelIndex
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QColumnView12setRootIndexERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setRootIndex", args)
  }

  return
}

// sizeHint()
func (this *QColumnView) SizeHint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QColumnView8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColumnView", "sizeHint", args)
  }

  return
}

// setModel(class QAbstractItemModel *)
func (this *QColumnView) SetModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnView8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(*qtcore.QAbstractItemModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QColumnView8setModelEP18QAbstractItemModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setModel", args)
  }

  return
}

// setSelectionModel(class QItemSelectionModel *)
func (this *QColumnView) SetSelectionModel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QColumnView17setSelectionModelEP19QItemSelectionModel
    // invoke: void setSelectionModel(class QItemSelectionModel *)
    var arg0 = args[0].(*qtcore.QItemSelectionModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QColumnView17setSelectionModelEP19QItemSelectionModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColumnView", "setSelectionModel", args)
  }

  return
}

// columnWidths()
func (this *QColumnView) ColumnWidths(args ...interface{}) () {
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
    C.C_ZNK11QColumnView12columnWidthsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "columnWidths", args)
  }

  return
}

// metaObject()
func (this *QColumnView) MetaObject(args ...interface{}) () {
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
    C.C_ZNK11QColumnView10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QColumnView", "metaObject", args)
  }

  return
}

// QColumnView(class QWidget *)
func GcfreeQColumnView(this *QColumnView) {
  qtrt.UniverseFree(this)
}
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QColumnViewC2EP7QWidget(arg0)
    this := &QColumnView{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQColumnView)
    return this
  default:
    qtrt.ErrorResolve("QColumnView", "QColumnView", args)
  }

  return nil // QColumnView{Qclsinst:qthis}
}

// <= body block end

