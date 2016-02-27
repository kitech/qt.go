package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  void QUndoView::setCleanIcon(const QIcon & icon);
extern void C_ZN9QUndoView12setCleanIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  void QUndoView::~QUndoView();
extern void C_ZN9QUndoViewD2Ev(void* qthis); // 4
  // proto:  QIcon QUndoView::cleanIcon();
extern void* C_ZNK9QUndoView9cleanIconEv(void* qthis); // 4
  // proto:  void QUndoView::QUndoView(QWidget * parent);
extern void* C_ZN9QUndoViewC2EP7QWidget(void* arg0); // 3
  // proto:  void QUndoView::QUndoView(QUndoGroup * group, QWidget * parent);
extern void* C_ZN9QUndoViewC2EP10QUndoGroupP7QWidget(void* arg0, void* arg1); // 3
  // proto:  void QUndoView::QUndoView(QUndoStack * stack, QWidget * parent);
extern void* C_ZN9QUndoViewC2EP10QUndoStackP7QWidget(void* arg0, void* arg1); // 3
  // proto:  QString QUndoView::emptyLabel();
extern void* C_ZNK9QUndoView10emptyLabelEv(void* qthis); // 4
  // proto:  void QUndoView::setEmptyLabel(const QString & label);
extern void C_ZN9QUndoView13setEmptyLabelERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QUndoView::setStack(QUndoStack * stack);
extern void C_ZN9QUndoView8setStackEP10QUndoStack(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QUndoView::metaObject();
extern void C_ZNK9QUndoView10metaObjectEv(void* qthis); // 4
  // proto:  void QUndoView::setGroup(QUndoGroup * group);
extern void C_ZN9QUndoView8setGroupEP10QUndoGroup(void* qthis, void* arg0); // 4
  // proto:  QUndoGroup * QUndoView::group();
extern void* C_ZNK9QUndoView5groupEv(void* qthis); // 4
  // proto:  QUndoStack * QUndoView::stack();
extern void* C_ZNK9QUndoView5stackEv(void* qthis); // 4
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

// class sizeof(QUndoView)=1
type QUndoView struct {
  /*qbase*/ QListView;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setCleanIcon(const class QIcon &)
func (this *QUndoView) SetCleanIcon(args ...interface{}) () {
  // setCleanIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoView12setCleanIconERK5QIcon
    // invoke: void setCleanIcon(const class QIcon &)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUndoView12setCleanIconERK5QIcon(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setCleanIcon", args)
  }

  return
}

// ~QUndoView()
func (this *QUndoView) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN9QUndoViewD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QUndoView", "~QUndoView", args)
  }

  return
}

// cleanIcon()
func (this *QUndoView) CleanIcon(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QUndoView9cleanIconEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoView", "cleanIcon", args)
  }

  return
}

// QUndoView(class QWidget *)
func GcfreeQUndoView(this *QUndoView) {
  qtrt.UniverseFree(this)
}
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QUndoViewC2EP7QWidget(arg0)
    this := &QUndoView{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQUndoView)
    return this
  case 1:
    // invoke: _ZN9QUndoViewC1EP10QUndoGroupP7QWidget
    // invoke: void QUndoView(class QUndoGroup *, class QWidget *)
    var arg0 = args[0].(*QUndoGroup).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QUndoViewC2EP10QUndoGroupP7QWidget(arg0, arg1)
    this := &QUndoView{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQUndoView)
    return this
  case 2:
    // invoke: _ZN9QUndoViewC1EP10QUndoStackP7QWidget
    // invoke: void QUndoView(class QUndoStack *, class QWidget *)
    var arg0 = args[0].(*QUndoStack).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QUndoViewC2EP10QUndoStackP7QWidget(arg0, arg1)
    this := &QUndoView{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQUndoView)
    return this
  default:
    qtrt.ErrorResolve("QUndoView", "QUndoView", args)
  }

  return nil // QUndoView{Qclsinst:qthis}
}

// emptyLabel()
func (this *QUndoView) EmptyLabel(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QUndoView10emptyLabelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoView", "emptyLabel", args)
  }

  return
}

// setEmptyLabel(const class QString &)
func (this *QUndoView) SetEmptyLabel(args ...interface{}) () {
  // setEmptyLabel(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUndoView13setEmptyLabelERK7QString
    // invoke: void setEmptyLabel(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUndoView13setEmptyLabelERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setEmptyLabel", args)
  }

  return
}

// setStack(class QUndoStack *)
func (this *QUndoView) SetStack(args ...interface{}) () {
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
    var arg0 = args[0].(*QUndoStack).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUndoView8setStackEP10QUndoStack(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setStack", args)
  }

  return
}

// metaObject()
func (this *QUndoView) MetaObject(args ...interface{}) () {
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
    C.C_ZNK9QUndoView10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUndoView", "metaObject", args)
  }

  return
}

// setGroup(class QUndoGroup *)
func (this *QUndoView) SetGroup(args ...interface{}) () {
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
    var arg0 = args[0].(*QUndoGroup).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUndoView8setGroupEP10QUndoGroup(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setGroup", args)
  }

  return
}

// group()
func (this *QUndoView) Group(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QUndoView5groupEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUndoGroup{}) // "QUndoGroup *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoView", "group", args)
  }

  return
}

// stack()
func (this *QUndoView) Stack(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QUndoView5stackEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUndoStack{}) // "QUndoStack *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUndoView", "stack", args)
  }

  return
}

// <= body block end

