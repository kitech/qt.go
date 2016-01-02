package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QUndoView::QUndoView(QUndoGroup * group, QWidget * parent);
extern void* dector_ZN9QUndoViewC1EP10QUndoGroupP7QWidget(void* arg0, void* arg1);
extern void _ZN9QUndoViewC1EP10QUndoGroupP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QUndoView::setStack(QUndoStack * stack);
extern void _ZN9QUndoView8setStackEP10QUndoStack(void* qthis, void* arg0);
  // proto:  void QUndoView::setEmptyLabel(const QString & label);
extern void _ZN9QUndoView13setEmptyLabelERK7QString(void* qthis, void* arg0);
  // proto:  void QUndoView::setCleanIcon(const QIcon & icon);
extern void _ZN9QUndoView12setCleanIconERK5QIcon(void* qthis, void* arg0);
  // proto:  void QUndoView::setGroup(QUndoGroup * group);
extern void _ZN9QUndoView8setGroupEP10QUndoGroup(void* qthis, void* arg0);
  // proto:  QUndoGroup * QUndoView::group();
extern void _ZNK9QUndoView5groupEv(void* qthis);
  // proto:  const QMetaObject * QUndoView::metaObject();
extern void _ZNK9QUndoView10metaObjectEv(void* qthis);
  // proto:  QUndoStack * QUndoView::stack();
extern void _ZNK9QUndoView5stackEv(void* qthis);
  // proto:  QIcon QUndoView::cleanIcon();
extern void _ZNK9QUndoView9cleanIconEv(void* qthis);
  // proto:  QString QUndoView::emptyLabel();
extern void _ZNK9QUndoView10emptyLabelEv(void* qthis);
  // proto:  void QUndoView::QUndoView(const QUndoView & );
extern void* dector_ZN9QUndoViewC1ERKS_(void* arg0);
extern void _ZN9QUndoViewC1ERKS_(void* qthis, void* arg0);
  // proto:  void QUndoView::QUndoView(QWidget * parent);
extern void* dector_ZN9QUndoViewC1EP7QWidget(void* arg0);
extern void _ZN9QUndoViewC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QUndoView::~QUndoView();
extern void _ZN9QUndoViewD0Ev(void* qthis);
  // proto:  void QUndoView::QUndoView(QUndoStack * stack, QWidget * parent);
extern void* dector_ZN9QUndoViewC1EP10QUndoStackP7QWidget(void* arg0, void* arg1);
extern void _ZN9QUndoViewC1EP10QUndoStackP7QWidget(void* qthis, void* arg0, void* arg1);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QUndoView::QUndoView(QUndoGroup * group, QWidget * parent);
func NewQUndoView(args ...interface{}) QUndoView {
  return QUndoView{}
}

  // proto:  void QUndoView::setStack(QUndoStack * stack);
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
    C._ZN9QUndoView8setStackEP10QUndoStack(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setStack", args)
  }

}

  // proto:  void QUndoView::setEmptyLabel(const QString & label);
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
    C._ZN9QUndoView13setEmptyLabelERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setEmptyLabel", args)
  }

}

  // proto:  void QUndoView::setCleanIcon(const QIcon & icon);
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
    C._ZN9QUndoView12setCleanIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setCleanIcon", args)
  }

}

  // proto:  void QUndoView::setGroup(QUndoGroup * group);
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
    C._ZN9QUndoView8setGroupEP10QUndoGroup(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUndoView", "setGroup", args)
  }

}

  // proto:  QUndoGroup * QUndoView::group();
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
    C._ZNK9QUndoView5groupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoView", "group", args)
  }

}

  // proto:  const QMetaObject * QUndoView::metaObject();
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
    C._ZNK9QUndoView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoView", "metaObject", args)
  }

}

  // proto:  QUndoStack * QUndoView::stack();
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
    C._ZNK9QUndoView5stackEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoView", "stack", args)
  }

}

  // proto:  QIcon QUndoView::cleanIcon();
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
    C._ZNK9QUndoView9cleanIconEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoView", "cleanIcon", args)
  }

}

  // proto:  QString QUndoView::emptyLabel();
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
    C._ZNK9QUndoView10emptyLabelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUndoView", "emptyLabel", args)
  }

}

  // proto:  void QUndoView::~QUndoView();
func (this *QUndoView) FreeQUndoView(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUndoView", "~QUndoView", args)
  }

}

// <= body block end

