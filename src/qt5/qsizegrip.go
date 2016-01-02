package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qsizegrip.h
// dst-file: /src/widgets/qsizegrip.go
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
  // proto:  void QSizeGrip::QSizeGrip(const QSizeGrip & );
extern void* dector_ZN9QSizeGripC1ERKS_(void* arg0);
extern void _ZN9QSizeGripC1ERKS_(void* qthis, void* arg0);
  // proto:  void QSizeGrip::QSizeGrip(QWidget * parent);
extern void* dector_ZN9QSizeGripC1EP7QWidget(void* arg0);
extern void _ZN9QSizeGripC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QSizeGrip::~QSizeGrip();
extern void _ZN9QSizeGripD0Ev(void* qthis);
  // proto:  void QSizeGrip::setVisible(bool );
extern void _ZN9QSizeGrip10setVisibleEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QSizeGrip::metaObject();
extern void _ZNK9QSizeGrip10metaObjectEv(void* qthis);
  // proto:  QSize QSizeGrip::sizeHint();
extern void _ZNK9QSizeGrip8sizeHintEv(void* qthis);
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

// class sizeof(QSizeGrip)=1
type QSizeGrip struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QSizeGrip::QSizeGrip(const QSizeGrip & );
func NewQSizeGrip(args ...interface{}) QSizeGrip {
  return QSizeGrip{}
}

  // proto:  void QSizeGrip::~QSizeGrip();
func (this *QSizeGrip) FreeQSizeGrip(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSizeGrip", "~QSizeGrip", args)
  }

}

  // proto:  void QSizeGrip::setVisible(bool );
func (this *QSizeGrip) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSizeGrip10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QSizeGrip10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizeGrip", "setVisible", args)
  }

}

  // proto:  const QMetaObject * QSizeGrip::metaObject();
func (this *QSizeGrip) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSizeGrip10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QSizeGrip10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizeGrip", "metaObject", args)
  }

}

  // proto:  QSize QSizeGrip::sizeHint();
func (this *QSizeGrip) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSizeGrip8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK9QSizeGrip8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizeGrip", "sizeHint", args)
  }

}

// <= body block end

