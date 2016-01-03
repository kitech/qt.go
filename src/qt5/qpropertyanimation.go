package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  QByteArray QPropertyAnimation::propertyName();
extern void _ZNK18QPropertyAnimation12propertyNameEv(void* qthis);
  // proto:  void QPropertyAnimation::setTargetObject(QObject * target);
extern void _ZN18QPropertyAnimation15setTargetObjectEP7QObject(void* qthis, void* arg0);
  // proto:  void QPropertyAnimation::QPropertyAnimation(QObject * parent);
extern void* dector_ZN18QPropertyAnimationC1EP7QObject(void* arg0);
extern void _ZN18QPropertyAnimationC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QPropertyAnimation::~QPropertyAnimation();
extern void _ZN18QPropertyAnimationD0Ev(void* qthis);
  // proto:  QObject * QPropertyAnimation::targetObject();
extern void _ZNK18QPropertyAnimation12targetObjectEv(void* qthis);
  // proto:  const QMetaObject * QPropertyAnimation::metaObject();
extern void _ZNK18QPropertyAnimation10metaObjectEv(void* qthis);
  // proto:  void QPropertyAnimation::QPropertyAnimation(const QPropertyAnimation & );
extern void* dector_ZN18QPropertyAnimationC1ERKS_(void* arg0);
extern void _ZN18QPropertyAnimationC1ERKS_(void* qthis, void* arg0);
  // proto:  void QPropertyAnimation::QPropertyAnimation(QObject * target, const QByteArray & propertyName, QObject * parent);
extern void* dector_ZN18QPropertyAnimationC1EP7QObjectRK10QByteArrayS1_(void* arg0, void* arg1, void* arg2);
extern void _ZN18QPropertyAnimationC1EP7QObjectRK10QByteArrayS1_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPropertyAnimation::setPropertyName(const QByteArray & propertyName);
extern void _ZN18QPropertyAnimation15setPropertyNameERK10QByteArray(void* qthis, void* arg0);
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

  // proto:  QByteArray QPropertyAnimation::propertyName();
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
    C._ZNK18QPropertyAnimation12propertyNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "propertyName", args)
  }

}

  // proto:  void QPropertyAnimation::setTargetObject(QObject * target);
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
    C._ZN18QPropertyAnimation15setTargetObjectEP7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "setTargetObject", args)
  }

}

  // proto:  void QPropertyAnimation::QPropertyAnimation(QObject * parent);
func NewQPropertyAnimation(args ...interface{}) QPropertyAnimation {
  return QPropertyAnimation{}
}

  // proto:  void QPropertyAnimation::~QPropertyAnimation();
func (this *QPropertyAnimation) FreeQPropertyAnimation(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "~QPropertyAnimation", args)
  }

}

  // proto:  QObject * QPropertyAnimation::targetObject();
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
    C._ZNK18QPropertyAnimation12targetObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "targetObject", args)
  }

}

  // proto:  const QMetaObject * QPropertyAnimation::metaObject();
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
    C._ZNK18QPropertyAnimation10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "metaObject", args)
  }

}

  // proto:  void QPropertyAnimation::setPropertyName(const QByteArray & propertyName);
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
    C._ZN18QPropertyAnimation15setPropertyNameERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "setPropertyName", args)
  }

}

// <= body block end

