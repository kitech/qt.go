package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qcoreevent.h
// dst-file: /src/core/qcoreevent.go
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
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QDeferredDeleteEvent::QDeferredDeleteEvent();
extern void* C_ZN20QDeferredDeleteEventC2Ev(); // 3
  // proto:  void QDeferredDeleteEvent::~QDeferredDeleteEvent();
extern void C_ZN20QDeferredDeleteEventD2Ev(void* qthis); // 4
  // proto:  int QDeferredDeleteEvent::loopLevel();
extern int32_t C_ZNK20QDeferredDeleteEvent9loopLevelEv(void* qthis); // 2
  // proto:  void QDynamicPropertyChangeEvent::QDynamicPropertyChangeEvent(const QByteArray & name);
extern void* C_ZN27QDynamicPropertyChangeEventC2ERK10QByteArray(void* arg0); // 3
  // proto:  QByteArray QDynamicPropertyChangeEvent::propertyName();
extern void* C_ZNK27QDynamicPropertyChangeEvent12propertyNameEv(void* qthis); // 2
  // proto:  void QDynamicPropertyChangeEvent::~QDynamicPropertyChangeEvent();
extern void C_ZN27QDynamicPropertyChangeEventD2Ev(void* qthis); // 4
  // proto:  int QTimerEvent::timerId();
extern int32_t C_ZNK11QTimerEvent7timerIdEv(void* qthis); // 2
  // proto:  void QTimerEvent::~QTimerEvent();
extern void C_ZN11QTimerEventD2Ev(void* qthis); // 4
  // proto:  void QTimerEvent::QTimerEvent(int timerId);
extern void* C_ZN11QTimerEventC2Ei(int32_t arg0); // 3
  // proto:  bool QChildEvent::added();
extern bool C_ZNK11QChildEvent5addedEv(void* qthis); // 2
  // proto:  bool QChildEvent::polished();
extern bool C_ZNK11QChildEvent8polishedEv(void* qthis); // 2
  // proto:  void QChildEvent::~QChildEvent();
extern void C_ZN11QChildEventD2Ev(void* qthis); // 4
  // proto:  QObject * QChildEvent::child();
extern void* C_ZNK11QChildEvent5childEv(void* qthis); // 2
  // proto:  bool QChildEvent::removed();
extern bool C_ZNK11QChildEvent7removedEv(void* qthis); // 2
  // proto: static int QEvent::registerEventType(int hint);
extern int32_t C_ZN6QEvent17registerEventTypeEi(int32_t arg0); // 4
  // proto:  bool QEvent::spontaneous();
extern bool C_ZNK6QEvent11spontaneousEv(void* qthis); // 2
  // proto:  void QEvent::accept();
extern void C_ZN6QEvent6acceptEv(void* qthis); // 2
  // proto:  void QEvent::ignore();
extern void C_ZN6QEvent6ignoreEv(void* qthis); // 2
  // proto:  bool QEvent::isAccepted();
extern bool C_ZNK6QEvent10isAcceptedEv(void* qthis); // 2
  // proto:  void QEvent::setAccepted(bool accepted);
extern void C_ZN6QEvent11setAcceptedEb(void* qthis, bool arg0); // 2
  // proto:  QEvent::Type QEvent::type();
extern void C_ZNK6QEvent4typeEv(void* qthis); // 2
  // proto:  void QEvent::QEvent(const QEvent & other);
extern void* C_ZN6QEventC2ERKS_(void* arg0); // 3
  // proto:  void QEvent::~QEvent();
extern void C_ZN6QEventD2Ev(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QDeferredDeleteEvent)=24
type QDeferredDeleteEvent struct {
  /*qbase*/ QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDynamicPropertyChangeEvent)=32
type QDynamicPropertyChangeEvent struct {
  /*qbase*/ QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTimerEvent)=24
type QTimerEvent struct {
  /*qbase*/ QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QChildEvent)=32
type QChildEvent struct {
  /*qbase*/ QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QEvent)=24
type QEvent struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QDeferredDeleteEvent()
func GcfreeQDeferredDeleteEvent(this *QDeferredDeleteEvent) {
  qtrt.UniverseFree(this)
}
func NewQDeferredDeleteEvent(args ...interface{}) *QDeferredDeleteEvent {
  // QDeferredDeleteEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QDeferredDeleteEventC1Ev
    // invoke: void QDeferredDeleteEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QDeferredDeleteEventC2Ev()
    this := &QDeferredDeleteEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQDeferredDeleteEvent)
    return this
  default:
    qtrt.ErrorResolve("QDeferredDeleteEvent", "QDeferredDeleteEvent", args)
  }

  return nil // QDeferredDeleteEvent{Qclsinst:qthis}
}

// ~QDeferredDeleteEvent()
func (this *QDeferredDeleteEvent) Free(args ...interface{}) () {
  // ~QDeferredDeleteEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QDeferredDeleteEventD0Ev
    // invoke: void ~QDeferredDeleteEvent()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN20QDeferredDeleteEventD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QDeferredDeleteEvent", "~QDeferredDeleteEvent", args)
  }

  return
}

// loopLevel()
func (this *QDeferredDeleteEvent) LoopLevel(args ...interface{}) (ret interface{}) {
  // loopLevel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QDeferredDeleteEvent9loopLevelEv
    // invoke: int loopLevel()
    var ret0 = C.C_ZNK20QDeferredDeleteEvent9loopLevelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDeferredDeleteEvent", "loopLevel", args)
  }

  return
}

// QDynamicPropertyChangeEvent(const class QByteArray &)
func GcfreeQDynamicPropertyChangeEvent(this *QDynamicPropertyChangeEvent) {
  qtrt.UniverseFree(this)
}
func NewQDynamicPropertyChangeEvent(args ...interface{}) *QDynamicPropertyChangeEvent {
  // QDynamicPropertyChangeEvent(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QDynamicPropertyChangeEventC1ERK10QByteArray
    // invoke: void QDynamicPropertyChangeEvent(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QDynamicPropertyChangeEventC2ERK10QByteArray(arg0)
    this := &QDynamicPropertyChangeEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQDynamicPropertyChangeEvent)
    return this
  default:
    qtrt.ErrorResolve("QDynamicPropertyChangeEvent", "QDynamicPropertyChangeEvent", args)
  }

  return nil // QDynamicPropertyChangeEvent{Qclsinst:qthis}
}

// propertyName()
func (this *QDynamicPropertyChangeEvent) PropertyName(args ...interface{}) (ret interface{}) {
  // propertyName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QDynamicPropertyChangeEvent12propertyNameEv
    // invoke: QByteArray propertyName()
    var ret0 = C.C_ZNK27QDynamicPropertyChangeEvent12propertyNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDynamicPropertyChangeEvent", "propertyName", args)
  }

  return
}

// ~QDynamicPropertyChangeEvent()
func (this *QDynamicPropertyChangeEvent) Free(args ...interface{}) () {
  // ~QDynamicPropertyChangeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QDynamicPropertyChangeEventD0Ev
    // invoke: void ~QDynamicPropertyChangeEvent()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN27QDynamicPropertyChangeEventD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QDynamicPropertyChangeEvent", "~QDynamicPropertyChangeEvent", args)
  }

  return
}

// timerId()
func (this *QTimerEvent) TimerId(args ...interface{}) (ret interface{}) {
  // timerId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTimerEvent7timerIdEv
    // invoke: int timerId()
    var ret0 = C.C_ZNK11QTimerEvent7timerIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimerEvent", "timerId", args)
  }

  return
}

// ~QTimerEvent()
func (this *QTimerEvent) Free(args ...interface{}) () {
  // ~QTimerEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTimerEventD0Ev
    // invoke: void ~QTimerEvent()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QTimerEventD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTimerEvent", "~QTimerEvent", args)
  }

  return
}

// QTimerEvent(int)
func GcfreeQTimerEvent(this *QTimerEvent) {
  qtrt.UniverseFree(this)
}
func NewQTimerEvent(args ...interface{}) *QTimerEvent {
  // QTimerEvent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTimerEventC1Ei
    // invoke: void QTimerEvent(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTimerEventC2Ei(arg0)
    this := &QTimerEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTimerEvent)
    return this
  default:
    qtrt.ErrorResolve("QTimerEvent", "QTimerEvent", args)
  }

  return nil // QTimerEvent{Qclsinst:qthis}
}

// added()
func (this *QChildEvent) Added(args ...interface{}) (ret interface{}) {
  // added()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent5addedEv
    // invoke: bool added()
    var ret0 = C.C_ZNK11QChildEvent5addedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChildEvent", "added", args)
  }

  return
}

// polished()
func (this *QChildEvent) Polished(args ...interface{}) (ret interface{}) {
  // polished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent8polishedEv
    // invoke: bool polished()
    var ret0 = C.C_ZNK11QChildEvent8polishedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChildEvent", "polished", args)
  }

  return
}

// ~QChildEvent()
func (this *QChildEvent) Free(args ...interface{}) () {
  // ~QChildEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QChildEventD0Ev
    // invoke: void ~QChildEvent()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QChildEventD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QChildEvent", "~QChildEvent", args)
  }

  return
}

// child()
func (this *QChildEvent) Child(args ...interface{}) (ret interface{}) {
  // child()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent5childEv
    // invoke: QObject * child()
    var ret0 = C.C_ZNK11QChildEvent5childEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChildEvent", "child", args)
  }

  return
}

// removed()
func (this *QChildEvent) Removed(args ...interface{}) (ret interface{}) {
  // removed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent7removedEv
    // invoke: bool removed()
    var ret0 = C.C_ZNK11QChildEvent7removedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QChildEvent", "removed", args)
  }

  return
}

// registerEventType(int)
func (this *QEvent) RegisterEventType_s(args ...interface{}) (ret interface{}) {
  // registerEventType(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent17registerEventTypeEi
    // invoke: int registerEventType(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN6QEvent17registerEventTypeEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEvent", "registerEventType", args)
  }

  return
}

// spontaneous()
func (this *QEvent) Spontaneous(args ...interface{}) (ret interface{}) {
  // spontaneous()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QEvent11spontaneousEv
    // invoke: bool spontaneous()
    var ret0 = C.C_ZNK6QEvent11spontaneousEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEvent", "spontaneous", args)
  }

  return
}

// accept()
func (this *QEvent) Accept(args ...interface{}) () {
  // accept()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent6acceptEv
    // invoke: void accept()
    C.C_ZN6QEvent6acceptEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "accept", args)
  }

  return
}

// ignore()
func (this *QEvent) Ignore(args ...interface{}) () {
  // ignore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent6ignoreEv
    // invoke: void ignore()
    C.C_ZN6QEvent6ignoreEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "ignore", args)
  }

  return
}

// isAccepted()
func (this *QEvent) IsAccepted(args ...interface{}) (ret interface{}) {
  // isAccepted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QEvent10isAcceptedEv
    // invoke: bool isAccepted()
    var ret0 = C.C_ZNK6QEvent10isAcceptedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QEvent", "isAccepted", args)
  }

  return
}

// setAccepted(_Bool)
func (this *QEvent) SetAccepted(args ...interface{}) () {
  // setAccepted(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent11setAcceptedEb
    // invoke: void setAccepted(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN6QEvent11setAcceptedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEvent", "setAccepted", args)
  }

  return
}

// type()
func (this *QEvent) Type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QEvent4typeEv
    // invoke: QEvent::Type type()
    C.C_ZNK6QEvent4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "type", args)
  }

  return
}

// QEvent(const class QEvent &)
func GcfreeQEvent(this *QEvent) {
  qtrt.UniverseFree(this)
}
func NewQEvent(args ...interface{}) *QEvent {
  // QEvent(const class QEvent &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "const QEvent &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEventC1ERKS_
    // invoke: void QEvent(const class QEvent &)
    var arg0 = args[0].(*QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QEventC2ERKS_(arg0)
    this := &QEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQEvent)
    return this
  default:
    qtrt.ErrorResolve("QEvent", "QEvent", args)
  }

  return nil // QEvent{Qclsinst:qthis}
}

// ~QEvent()
func (this *QEvent) Free(args ...interface{}) () {
  // ~QEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEventD0Ev
    // invoke: void ~QEvent()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN6QEventD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QEvent", "~QEvent", args)
  }

  return
}

// <= body block end

