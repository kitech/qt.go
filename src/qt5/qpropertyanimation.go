package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qpropertyanimation.h
// dst-file: /src/core/qpropertyanimation.go
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
  // proto:  void QPropertyAnimation::~QPropertyAnimation();
extern void C_ZN18QPropertyAnimationD2Ev(void* qthis); // 4
  // proto:  QByteArray QPropertyAnimation::propertyName();
extern void C_ZNK18QPropertyAnimation12propertyNameEv(void* qthis); // 4
  // proto:  void QPropertyAnimation::QPropertyAnimation(QObject * parent);
extern void C_ZN18QPropertyAnimationC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QPropertyAnimation::QPropertyAnimation(QObject * target, const QByteArray & propertyName, QObject * parent);
extern void C_ZN18QPropertyAnimationC2EP7QObjectRK10QByteArrayS1_(void* qthis, void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QPropertyAnimation::setTargetObject(QObject * target);
extern void C_ZN18QPropertyAnimation15setTargetObjectEP7QObject(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QPropertyAnimation::metaObject();
extern void C_ZNK18QPropertyAnimation10metaObjectEv(void* qthis); // 4
  // proto:  void QPropertyAnimation::setPropertyName(const QByteArray & propertyName);
extern void C_ZN18QPropertyAnimation15setPropertyNameERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QObject * QPropertyAnimation::targetObject();
extern void C_ZNK18QPropertyAnimation12targetObjectEv(void* qthis); // 4
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

// class sizeof(QPropertyAnimation)=1
type QPropertyAnimation struct {
  /*qbase*/ QVariantAnimation;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QPropertyAnimation()
func (this *QPropertyAnimation) FreeQPropertyAnimation(args ...interface{}) () {
  // ~QPropertyAnimation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QPropertyAnimationD0Ev
    // invoke: void ~QPropertyAnimation()
    C.C_ZN18QPropertyAnimationD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "~QPropertyAnimation", args)
  }

}

// propertyName()
func (this *QPropertyAnimation) propertyName(args ...interface{}) () {
  // propertyName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QPropertyAnimation12propertyNameEv
    // invoke: QByteArray propertyName()
    C.C_ZNK18QPropertyAnimation12propertyNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "propertyName", args)
  }

}

// QPropertyAnimation(class QObject *)
func NewQPropertyAnimation(args ...interface{}) QPropertyAnimation {
  // QPropertyAnimation(class QObject *)
  // QPropertyAnimation(class QObject *, const class QByteArray &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][2] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QPropertyAnimationC1EP7QObject
    // invoke: void QPropertyAnimation(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN18QPropertyAnimationC2EP7QObject(qthis, arg0)
  case 1:
    // invoke: _ZN18QPropertyAnimationC1EP7QObjectRK10QByteArrayS1_
    // invoke: void QPropertyAnimation(class QObject *, const class QByteArray &, class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN18QPropertyAnimationC2EP7QObjectRK10QByteArrayS1_(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "QPropertyAnimation", args)
  }

  return QPropertyAnimation{}
}

// setTargetObject(class QObject *)
func (this *QPropertyAnimation) setTargetObject(args ...interface{}) () {
  // setTargetObject(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QPropertyAnimation15setTargetObjectEP7QObject
    // invoke: void setTargetObject(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QPropertyAnimation15setTargetObjectEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "setTargetObject", args)
  }

}

// metaObject()
func (this *QPropertyAnimation) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QPropertyAnimation10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QPropertyAnimation10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "metaObject", args)
  }

}

// setPropertyName(const class QByteArray &)
func (this *QPropertyAnimation) setPropertyName(args ...interface{}) () {
  // setPropertyName(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QPropertyAnimation15setPropertyNameERK10QByteArray
    // invoke: void setPropertyName(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QPropertyAnimation15setPropertyNameERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "setPropertyName", args)
  }

}

// targetObject()
func (this *QPropertyAnimation) targetObject(args ...interface{}) () {
  // targetObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QPropertyAnimation12targetObjectEv
    // invoke: QObject * targetObject()
    C.C_ZNK18QPropertyAnimation12targetObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "targetObject", args)
  }

}

// <= body block end

