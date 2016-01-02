package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qmap.h
// dst-file: /src/core/qmap.go
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
  // proto:  void QMapDataBase::rebalance(QMapNodeBase * x);
extern void _ZN12QMapDataBase9rebalanceEP12QMapNodeBase(void* qthis, void* arg0);
  // proto: static QMapDataBase * QMapDataBase::createData();
extern void _ZN12QMapDataBase10createDataEv();
  // proto:  void QMapDataBase::rotateRight(QMapNodeBase * x);
extern void _ZN12QMapDataBase11rotateRightEP12QMapNodeBase(void* qthis, void* arg0);
  // proto:  void QMapDataBase::freeTree(QMapNodeBase * root, int alignment);
extern void _ZN12QMapDataBase8freeTreeEP12QMapNodeBasei(void* qthis, void* arg0, int arg1);
  // proto: static void QMapDataBase::freeData(QMapDataBase * d);
extern void _ZN12QMapDataBase8freeDataEPS_(void* arg0);
  // proto:  void QMapDataBase::rotateLeft(QMapNodeBase * x);
extern void _ZN12QMapDataBase10rotateLeftEP12QMapNodeBase(void* qthis, void* arg0);
  // proto:  void QMapDataBase::recalcMostLeftNode();
extern void _ZN12QMapDataBase18recalcMostLeftNodeEv(void* qthis);
  // proto:  QMapNodeBase * QMapDataBase::createNode(int size, int alignment, QMapNodeBase * parent, bool left);
extern void _ZN12QMapDataBase10createNodeEiiP12QMapNodeBaseb(void* qthis, int arg0, int arg1, void* arg2, bool arg3);
  // proto:  void QMapDataBase::freeNodeAndRebalance(QMapNodeBase * z);
extern void _ZN12QMapDataBase20freeNodeAndRebalanceEP12QMapNodeBase(void* qthis, void* arg0);
  // proto:  void QMapNodeBase::setParent(QMapNodeBase * pp);
extern void _ZN12QMapNodeBase9setParentEPS_(void* qthis, void* arg0);
  // proto:  QMapNodeBase * QMapNodeBase::previousNode();
extern void _ZN12QMapNodeBase12previousNodeEv(void* qthis);
  // proto:  QMapNodeBase * QMapNodeBase::nextNode();
extern void _ZN12QMapNodeBase8nextNodeEv(void* qthis);
  // proto:  QMapNodeBase * QMapNodeBase::parent();
extern void _ZNK12QMapNodeBase6parentEv(void* qthis);
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

// class sizeof(QMapDataBase)=1
type QMapDataBase struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMapNodeBase)=24
type QMapNodeBase struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QMapDataBase::rebalance(QMapNodeBase * x);
func (this *QMapDataBase) rebalance(args ...interface{}) () {
  // rebalance(struct QMapNodeBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase9rebalanceEP12QMapNodeBase
  default:
    qtrt.ErrorResolve("QMapDataBase", "rebalance", args)
  }

}

  // proto: static QMapDataBase * QMapDataBase::createData();
func (this *QMapDataBase) createData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMapDataBase", "createData", args)
  }

}

  // proto:  void QMapDataBase::rotateRight(QMapNodeBase * x);
func (this *QMapDataBase) rotateRight(args ...interface{}) () {
  // rotateRight(struct QMapNodeBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase11rotateRightEP12QMapNodeBase
  default:
    qtrt.ErrorResolve("QMapDataBase", "rotateRight", args)
  }

}

  // proto:  void QMapDataBase::freeTree(QMapNodeBase * root, int alignment);
func (this *QMapDataBase) freeTree(args ...interface{}) () {
  // freeTree(struct QMapNodeBase *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase8freeTreeEP12QMapNodeBasei
  default:
    qtrt.ErrorResolve("QMapDataBase", "freeTree", args)
  }

}

  // proto: static void QMapDataBase::freeData(QMapDataBase * d);
func (this *QMapDataBase) freeData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMapDataBase", "freeData", args)
  }

}

  // proto:  void QMapDataBase::rotateLeft(QMapNodeBase * x);
func (this *QMapDataBase) rotateLeft(args ...interface{}) () {
  // rotateLeft(struct QMapNodeBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase10rotateLeftEP12QMapNodeBase
  default:
    qtrt.ErrorResolve("QMapDataBase", "rotateLeft", args)
  }

}

  // proto:  void QMapDataBase::recalcMostLeftNode();
func (this *QMapDataBase) recalcMostLeftNode(args ...interface{}) () {
  // recalcMostLeftNode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase18recalcMostLeftNodeEv
  default:
    qtrt.ErrorResolve("QMapDataBase", "recalcMostLeftNode", args)
  }

}

  // proto:  QMapNodeBase * QMapDataBase::createNode(int size, int alignment, QMapNodeBase * parent, bool left);
func (this *QMapDataBase) createNode(args ...interface{}) () {
  // createNode(int, int, struct QMapNodeBase *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"
  vtys[0][3] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase10createNodeEiiP12QMapNodeBaseb
  default:
    qtrt.ErrorResolve("QMapDataBase", "createNode", args)
  }

}

  // proto:  void QMapDataBase::freeNodeAndRebalance(QMapNodeBase * z);
func (this *QMapDataBase) freeNodeAndRebalance(args ...interface{}) () {
  // freeNodeAndRebalance(struct QMapNodeBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase20freeNodeAndRebalanceEP12QMapNodeBase
  default:
    qtrt.ErrorResolve("QMapDataBase", "freeNodeAndRebalance", args)
  }

}

  // proto:  void QMapNodeBase::setParent(QMapNodeBase * pp);
func (this *QMapNodeBase) setParent(args ...interface{}) () {
  // setParent(struct QMapNodeBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapNodeBase9setParentEPS_
  default:
    qtrt.ErrorResolve("QMapNodeBase", "setParent", args)
  }

}

  // proto:  QMapNodeBase * QMapNodeBase::previousNode();
func (this *QMapNodeBase) previousNode(args ...interface{}) () {
  // previousNode()
  // previousNode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QMapNodeBase12previousNodeEv
  case 1:
    // invoke: _ZN12QMapNodeBase12previousNodeEv
  default:
    qtrt.ErrorResolve("QMapNodeBase", "previousNode", args)
  }

}

  // proto:  QMapNodeBase * QMapNodeBase::nextNode();
func (this *QMapNodeBase) nextNode(args ...interface{}) () {
  // nextNode()
  // nextNode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QMapNodeBase8nextNodeEv
  case 1:
    // invoke: _ZN12QMapNodeBase8nextNodeEv
  default:
    qtrt.ErrorResolve("QMapNodeBase", "nextNode", args)
  }

}

  // proto:  QMapNodeBase * QMapNodeBase::parent();
func (this *QMapNodeBase) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QMapNodeBase6parentEv
  default:
    qtrt.ErrorResolve("QMapNodeBase", "parent", args)
  }

}

// <= body block end

