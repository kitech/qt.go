package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  int QDeferredDeleteEvent::loopLevel();
extern void _ZNK20QDeferredDeleteEvent9loopLevelEv(void* qthis);
  // proto:  void QDeferredDeleteEvent::~QDeferredDeleteEvent();
extern void _ZN20QDeferredDeleteEventD0Ev(void* qthis);
  // proto:  void QDeferredDeleteEvent::QDeferredDeleteEvent();
extern void* dector_ZN20QDeferredDeleteEventC1Ev();
extern void _ZN20QDeferredDeleteEventC1Ev(void* qthis);
  // proto:  void QDynamicPropertyChangeEvent::~QDynamicPropertyChangeEvent();
extern void _ZN27QDynamicPropertyChangeEventD0Ev(void* qthis);
  // proto:  void QDynamicPropertyChangeEvent::QDynamicPropertyChangeEvent(const QByteArray & name);
extern void* dector_ZN27QDynamicPropertyChangeEventC1ERK10QByteArray(void* arg0);
extern void _ZN27QDynamicPropertyChangeEventC1ERK10QByteArray(void* qthis, void* arg0);
  // proto:  QByteArray QDynamicPropertyChangeEvent::propertyName();
extern void demth_ZNK27QDynamicPropertyChangeEvent12propertyNameEv(void* qthis);
  // proto:  void QTimerEvent::QTimerEvent(int timerId);
extern void* dector_ZN11QTimerEventC1Ei(int arg0);
extern void _ZN11QTimerEventC1Ei(void* qthis, int arg0);
  // proto:  void QTimerEvent::~QTimerEvent();
extern void _ZN11QTimerEventD0Ev(void* qthis);
  // proto:  int QTimerEvent::timerId();
extern void _ZNK11QTimerEvent7timerIdEv(void* qthis);
  // proto:  bool QChildEvent::added();
extern void _ZNK11QChildEvent5addedEv(void* qthis);
  // proto:  bool QChildEvent::polished();
extern void _ZNK11QChildEvent8polishedEv(void* qthis);
  // proto:  void QChildEvent::~QChildEvent();
extern void _ZN11QChildEventD0Ev(void* qthis);
  // proto:  bool QChildEvent::removed();
extern void _ZNK11QChildEvent7removedEv(void* qthis);
  // proto:  QObject * QChildEvent::child();
extern void _ZNK11QChildEvent5childEv(void* qthis);
  // proto:  void QEvent::setAccepted(bool accepted);
extern void demth_ZN6QEvent11setAcceptedEb(void* qthis, bool arg0);
  // proto:  void QEvent::ignore();
extern void demth_ZN6QEvent6ignoreEv(void* qthis);
  // proto:  bool QEvent::isAccepted();
extern void demth_ZNK6QEvent10isAcceptedEv(void* qthis);
  // proto:  void QEvent::~QEvent();
extern void _ZN6QEventD0Ev(void* qthis);
  // proto:  void QEvent::QEvent(const QEvent & other);
extern void* dector_ZN6QEventC1ERKS_(void* arg0);
extern void _ZN6QEventC1ERKS_(void* qthis, void* arg0);
  // proto:  void QEvent::accept();
extern void demth_ZN6QEvent6acceptEv(void* qthis);
  // proto: static int QEvent::registerEventType(int hint);
extern void _ZN6QEvent17registerEventTypeEi(int arg0);
  // proto:  bool QEvent::spontaneous();
extern void demth_ZNK6QEvent11spontaneousEv(void* qthis);
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

// class sizeof(QDeferredDeleteEvent)=24
type QDeferredDeleteEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDynamicPropertyChangeEvent)=32
type QDynamicPropertyChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTimerEvent)=24
type QTimerEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QChildEvent)=32
type QChildEvent struct {
  /*qbase*/ QEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QEvent)=24
type QEvent struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int QDeferredDeleteEvent::loopLevel();
func (this *QDeferredDeleteEvent) loopLevel(args ...interface{}) () {
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
    C._ZNK20QDeferredDeleteEvent9loopLevelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDeferredDeleteEvent", "loopLevel", args)
  }

}

  // proto:  void QDeferredDeleteEvent::~QDeferredDeleteEvent();
func (this *QDeferredDeleteEvent) FreeQDeferredDeleteEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDeferredDeleteEvent", "~QDeferredDeleteEvent", args)
  }

}

  // proto:  void QDeferredDeleteEvent::QDeferredDeleteEvent();
func NewQDeferredDeleteEvent(args ...interface{}) QDeferredDeleteEvent {
  return QDeferredDeleteEvent{}
}

  // proto:  void QDynamicPropertyChangeEvent::~QDynamicPropertyChangeEvent();
func (this *QDynamicPropertyChangeEvent) FreeQDynamicPropertyChangeEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDynamicPropertyChangeEvent", "~QDynamicPropertyChangeEvent", args)
  }

}

  // proto:  void QDynamicPropertyChangeEvent::QDynamicPropertyChangeEvent(const QByteArray & name);
func NewQDynamicPropertyChangeEvent(args ...interface{}) QDynamicPropertyChangeEvent {
  return QDynamicPropertyChangeEvent{}
}

  // proto:  QByteArray QDynamicPropertyChangeEvent::propertyName();
func (this *QDynamicPropertyChangeEvent) propertyName(args ...interface{}) () {
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
    C.demth_ZNK27QDynamicPropertyChangeEvent12propertyNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDynamicPropertyChangeEvent", "propertyName", args)
  }

}

  // proto:  void QTimerEvent::QTimerEvent(int timerId);
func NewQTimerEvent(args ...interface{}) QTimerEvent {
  return QTimerEvent{}
}

  // proto:  void QTimerEvent::~QTimerEvent();
func (this *QTimerEvent) FreeQTimerEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTimerEvent", "~QTimerEvent", args)
  }

}

  // proto:  int QTimerEvent::timerId();
func (this *QTimerEvent) timerId(args ...interface{}) () {
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
    C._ZNK11QTimerEvent7timerIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimerEvent", "timerId", args)
  }

}

  // proto:  bool QChildEvent::added();
func (this *QChildEvent) added(args ...interface{}) () {
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
    C._ZNK11QChildEvent5addedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChildEvent", "added", args)
  }

}

  // proto:  bool QChildEvent::polished();
func (this *QChildEvent) polished(args ...interface{}) () {
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
    C._ZNK11QChildEvent8polishedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChildEvent", "polished", args)
  }

}

  // proto:  void QChildEvent::~QChildEvent();
func (this *QChildEvent) FreeQChildEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QChildEvent", "~QChildEvent", args)
  }

}

  // proto:  bool QChildEvent::removed();
func (this *QChildEvent) removed(args ...interface{}) () {
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
    C._ZNK11QChildEvent7removedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChildEvent", "removed", args)
  }

}

  // proto:  QObject * QChildEvent::child();
func (this *QChildEvent) child(args ...interface{}) () {
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
    C._ZNK11QChildEvent5childEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChildEvent", "child", args)
  }

}

  // proto:  void QEvent::setAccepted(bool accepted);
func (this *QEvent) setAccepted(args ...interface{}) () {
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C.demth_ZN6QEvent11setAcceptedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEvent", "setAccepted", args)
  }

}

  // proto:  void QEvent::ignore();
func (this *QEvent) ignore(args ...interface{}) () {
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
    C.demth_ZN6QEvent6ignoreEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "ignore", args)
  }

}

  // proto:  bool QEvent::isAccepted();
func (this *QEvent) isAccepted(args ...interface{}) () {
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
    C.demth_ZNK6QEvent10isAcceptedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "isAccepted", args)
  }

}

  // proto:  void QEvent::~QEvent();
func (this *QEvent) FreeQEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEvent", "~QEvent", args)
  }

}

  // proto:  void QEvent::QEvent(const QEvent & other);
func NewQEvent(args ...interface{}) QEvent {
  return QEvent{}
}

  // proto:  void QEvent::accept();
func (this *QEvent) accept(args ...interface{}) () {
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
    C.demth_ZN6QEvent6acceptEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "accept", args)
  }

}

  // proto: static int QEvent::registerEventType(int hint);
func (this *QEvent) registerEventType_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QEvent", "registerEventType", args)
  }

}

  // proto:  bool QEvent::spontaneous();
func (this *QEvent) spontaneous(args ...interface{}) () {
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
    C.demth_ZNK6QEvent11spontaneousEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "spontaneous", args)
  }

}

// <= body block end

