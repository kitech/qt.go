package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qundoview.h
// dst-file: /src/widgets/qundoview.go
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
  // proto:  void QUndoView::setCleanIcon(const QIcon & icon);
extern void C_ZN9QUndoView12setCleanIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  void QUndoView::~QUndoView();
extern void C_ZN9QUndoViewD2Ev(void* qthis); // 4
  // proto:  QIcon QUndoView::cleanIcon();
extern void C_ZNK9QUndoView9cleanIconEv(void* qthis); // 4
  // proto:  void QUndoView::QUndoView(QWidget * parent);
extern void* C_ZN9QUndoViewC2EP7QWidget(void* arg0); // 3
  // proto:  void QUndoView::QUndoView(QUndoGroup * group, QWidget * parent);
extern void* C_ZN9QUndoViewC2EP10QUndoGroupP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QUndoView::QUndoView(QUndoStack * stack, QWidget * parent);
extern void* C_ZN9QUndoViewC2EP10QUndoStackP7QWidget(void* arg0, void* arg1); // 3
  // proto:  QString QUndoView::emptyLabel();
extern void C_ZNK9QUndoView10emptyLabelEv(void* qthis); // 4
  // proto:  void QUndoView::setEmptyLabel(const QString & label);
extern void C_ZN9QUndoView13setEmptyLabelERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QUndoView::setStack(QUndoStack * stack);
extern void C_ZN9QUndoView8setStackEP10QUndoStack(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QUndoView::metaObject();
extern void C_ZNK9QUndoView10metaObjectEv(void* qthis); // 4
  // proto:  void QUndoView::setGroup(QUndoGroup * group);
extern void C_ZN9QUndoView8setGroupEP10QUndoGroup(void* qthis, void* arg0); // 4
  // proto:  QUndoGroup * QUndoView::group();
extern void C_ZNK9QUndoView5groupEv(void* qthis); // 4
  // proto:  QUndoStack * QUndoView::stack();
extern void C_ZNK9QUndoView5stackEv(void* qthis); // 4
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

// class sizeof(QUndoView)=1
type QUndoView struct {
  /*qbase*/ QListView;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setCleanIcon(const class QIcon &)
func (this *QUndoView) setCleanIcon(args ...interface{}) () {
  // setCleanIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoView12setCleanIconERK5QIcon
    // invoke: void setCleanIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUndoView12setCleanIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setCleanIcon", args)
  }

}

// ~QUndoView()
func (this *QUndoView) FreeQUndoView(args ...interface{}) () {
  // ~QUndoView()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoViewD0Ev
    // invoke: void ~QUndoView()
    C.C_ZN9QUndoViewD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoView", "~QUndoView", args)
  }

}

// cleanIcon()
func (this *QUndoView) cleanIcon(args ...interface{}) () {
  // cleanIcon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUndoView9cleanIconEv
    // invoke: QIcon cleanIcon()
    var ret = C.C_ZNK9QUndoView9cleanIconEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoView", "cleanIcon", args)
  }

}

// QUndoView(class QWidget *)
func NewQUndoView(args ...interface{}) *QUndoView {
  // QUndoView(class QWidget *)
  // QUndoView(class QUndoGroup *, class QWidget *)
  // QUndoView(class QUndoStack *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QUndoGroup{}) // "QUndoGroup *"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QUndoStack{}) // "QUndoStack *"
  vtys[2][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoViewC1EP7QWidget
    // invoke: void QUndoView(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QUndoViewC2EP7QWidget(arg0)
    return &QUndoView{qclsinst:qthis}
  case 1:
    // invoke: _ZN9QUndoViewC1EP10QUndoGroupP7QWidget
    // invoke: void QUndoView(class QUndoGroup *, class QWidget *)
    var arg0 = args[0].(QUndoGroup).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QUndoViewC2EP10QUndoGroupP7QWidget(arg0, arg1)
    return &QUndoView{qclsinst:qthis}
  case 2:
    // invoke: _ZN9QUndoViewC1EP10QUndoStackP7QWidget
    // invoke: void QUndoView(class QUndoStack *, class QWidget *)
    var arg0 = args[0].(QUndoStack).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QUndoViewC2EP10QUndoStackP7QWidget(arg0, arg1)
    return &QUndoView{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QUndoView", "QUndoView", args)
  }

  return nil // QUndoView{qclsinst:qthis}
}

// emptyLabel()
func (this *QUndoView) emptyLabel(args ...interface{}) () {
  // emptyLabel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUndoView10emptyLabelEv
    // invoke: QString emptyLabel()
    var ret = C.C_ZNK9QUndoView10emptyLabelEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoView", "emptyLabel", args)
  }

}

// setEmptyLabel(const class QString &)
func (this *QUndoView) setEmptyLabel(args ...interface{}) () {
  // setEmptyLabel(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoView13setEmptyLabelERK7QString
    // invoke: void setEmptyLabel(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUndoView13setEmptyLabelERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setEmptyLabel", args)
  }

}

// setStack(class QUndoStack *)
func (this *QUndoView) setStack(args ...interface{}) () {
  // setStack(class QUndoStack *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUndoStack{}) // "QUndoStack *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoView8setStackEP10QUndoStack
    // invoke: void setStack(class QUndoStack *)
    var arg0 = args[0].(QUndoStack).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUndoView8setStackEP10QUndoStack(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setStack", args)
  }

}

// metaObject()
func (this *QUndoView) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUndoView10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QUndoView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoView", "metaObject", args)
  }

}

// setGroup(class QUndoGroup *)
func (this *QUndoView) setGroup(args ...interface{}) () {
  // setGroup(class QUndoGroup *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUndoGroup{}) // "QUndoGroup *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoView8setGroupEP10QUndoGroup
    // invoke: void setGroup(class QUndoGroup *)
    var arg0 = args[0].(QUndoGroup).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUndoView8setGroupEP10QUndoGroup(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setGroup", args)
  }

}

// group()
func (this *QUndoView) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUndoView5groupEv
    // invoke: QUndoGroup * group()
    var ret = C.C_ZNK9QUndoView5groupEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoView", "group", args)
  }

}

// stack()
func (this *QUndoView) stack(args ...interface{}) () {
  // stack()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUndoView5stackEv
    // invoke: QUndoStack * stack()
    var ret = C.C_ZNK9QUndoView5stackEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUndoView", "stack", args)
  }

}

// <= body block end

