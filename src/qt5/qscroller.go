package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qscroller.h
// dst-file: /src/widgets/qscroller.go
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
  // proto: static QScroller * QScroller::scroller(QObject * target);
extern void* C_ZN9QScroller8scrollerEP7QObject(void* arg0); // 4
  // proto: static const QScroller * QScroller::scroller(const QObject * target);
extern void* C_ZN9QScroller8scrollerEPK7QObject(void* arg0); // 4
  // proto:  void QScroller::scrollTo(const QPointF & pos, int scrollTime);
extern void C_ZN9QScroller8scrollToERK7QPointFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QScroller::scrollTo(const QPointF & pos);
extern void C_ZN9QScroller8scrollToERK7QPointF(void* qthis, void* arg0); // 4
  // proto: static QList<QScroller *> QScroller::activeScrollers();
extern void C_ZN9QScroller15activeScrollersEv(); // 4
  // proto: static bool QScroller::hasScroller(QObject * target);
extern bool C_ZN9QScroller11hasScrollerEP7QObject(void* arg0); // 4
  // proto:  void QScroller::setSnapPositionsY(qreal first, qreal interval);
extern void C_ZN9QScroller17setSnapPositionsYEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  void QScroller::setSnapPositionsX(qreal first, qreal interval);
extern void C_ZN9QScroller17setSnapPositionsXEdd(void* qthis, double arg0, double arg1); // 4
  // proto: static void QScroller::ungrabGesture(QObject * target);
extern void C_ZN9QScroller13ungrabGestureEP7QObject(void* arg0); // 4
  // proto:  QScroller::State QScroller::state();
extern void C_ZNK9QScroller5stateEv(void* qthis); // 4
  // proto:  void QScroller::resendPrepareEvent();
extern void C_ZN9QScroller18resendPrepareEventEv(void* qthis); // 4
  // proto:  void QScroller::stop();
extern void C_ZN9QScroller4stopEv(void* qthis); // 4
  // proto:  QScrollerProperties QScroller::scrollerProperties();
extern void* C_ZNK9QScroller18scrollerPropertiesEv(void* qthis); // 4
  // proto:  QPointF QScroller::finalPosition();
extern void* C_ZNK9QScroller13finalPositionEv(void* qthis); // 4
  // proto:  void QScroller::ensureVisible(const QRectF & rect, qreal xmargin, qreal ymargin, int scrollTime);
extern void C_ZN9QScroller13ensureVisibleERK6QRectFddi(void* qthis, void* arg0, double arg1, double arg2, int32_t arg3); // 4
  // proto:  void QScroller::ensureVisible(const QRectF & rect, qreal xmargin, qreal ymargin);
extern void C_ZN9QScroller13ensureVisibleERK6QRectFdd(void* qthis, void* arg0, double arg1, double arg2); // 4
  // proto:  QPointF QScroller::pixelPerMeter();
extern void* C_ZNK9QScroller13pixelPerMeterEv(void* qthis); // 4
  // proto:  const QMetaObject * QScroller::metaObject();
extern void C_ZNK9QScroller10metaObjectEv(void* qthis); // 4
  // proto:  QObject * QScroller::target();
extern void* C_ZNK9QScroller6targetEv(void* qthis); // 4
  // proto: static Qt::GestureType QScroller::grabbedGesture(QObject * target);
extern void C_ZN9QScroller14grabbedGestureEP7QObject(void* arg0); // 4
  // proto:  void QScroller::setScrollerProperties(const QScrollerProperties & prop);
extern void C_ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties(void* qthis, void* arg0); // 4
  // proto:  QPointF QScroller::velocity();
extern void* C_ZNK9QScroller8velocityEv(void* qthis); // 4
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

// class sizeof(QScroller)=1
type QScroller struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _stateChanged QScroller_stateChanged_signal;
//  _scrollerPropertiesChanged QScroller_scrollerPropertiesChanged_signal;
}

// scroller(class QObject *)
func (this *QScroller) Scroller_S(args ...interface{}) (ret interface{}) {
  // scroller(class QObject *)
  // scroller(const class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "const QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller8scrollerEP7QObject
    // invoke: QScroller * scroller(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QScroller8scrollerEP7QObject(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QScroller{}) // "QScroller *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN9QScroller8scrollerEPK7QObject
    // invoke: const QScroller * scroller(const class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QScroller8scrollerEPK7QObject(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QScroller{}) // "const QScroller *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScroller", "scroller", args)
  }

  return
}

// scrollTo(const class QPointF &, int)
func (this *QScroller) Scrollto(args ...interface{}) () {
  // scrollTo(const class QPointF &, int)
  // scrollTo(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller8scrollToERK7QPointFi
    // invoke: void scrollTo(const class QPointF &, int)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QScroller8scrollToERK7QPointFi(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN9QScroller8scrollToERK7QPointF
    // invoke: void scrollTo(const class QPointF &)
    var arg0 = args[0].(*QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QScroller8scrollToERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScroller", "scrollTo", args)
  }

  return
}

// activeScrollers()
func (this *QScroller) Activescrollers_S(args ...interface{}) () {
  // activeScrollers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller15activeScrollersEv
    // invoke: QList<QScroller *> activeScrollers()
    C.C_ZN9QScroller15activeScrollersEv()
  default:
    qtrt.ErrorResolve("QScroller", "activeScrollers", args)
  }

  return
}

// hasScroller(class QObject *)
func (this *QScroller) Hasscroller_S(args ...interface{}) (ret interface{}) {
  // hasScroller(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller11hasScrollerEP7QObject
    // invoke: bool hasScroller(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QScroller11hasScrollerEP7QObject(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScroller", "hasScroller", args)
  }

  return
}

// setSnapPositionsY(qreal, qreal)
func (this *QScroller) Setsnappositionsy(args ...interface{}) () {
  // setSnapPositionsY(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller17setSnapPositionsYEdd
    // invoke: void setSnapPositionsY(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN9QScroller17setSnapPositionsYEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QScroller", "setSnapPositionsY", args)
  }

  return
}

// setSnapPositionsX(qreal, qreal)
func (this *QScroller) Setsnappositionsx(args ...interface{}) () {
  // setSnapPositionsX(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller17setSnapPositionsXEdd
    // invoke: void setSnapPositionsX(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN9QScroller17setSnapPositionsXEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QScroller", "setSnapPositionsX", args)
  }

  return
}

// ungrabGesture(class QObject *)
func (this *QScroller) Ungrabgesture_S(args ...interface{}) () {
  // ungrabGesture(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller13ungrabGestureEP7QObject
    // invoke: void ungrabGesture(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QScroller13ungrabGestureEP7QObject(arg0)
  default:
    qtrt.ErrorResolve("QScroller", "ungrabGesture", args)
  }

  return
}

// state()
func (this *QScroller) State(args ...interface{}) () {
  // state()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller5stateEv
    // invoke: QScroller::State state()
    C.C_ZNK9QScroller5stateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "state", args)
  }

  return
}

// resendPrepareEvent()
func (this *QScroller) Resendprepareevent(args ...interface{}) () {
  // resendPrepareEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller18resendPrepareEventEv
    // invoke: void resendPrepareEvent()
    C.C_ZN9QScroller18resendPrepareEventEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "resendPrepareEvent", args)
  }

  return
}

// stop()
func (this *QScroller) Stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller4stopEv
    // invoke: void stop()
    C.C_ZN9QScroller4stopEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "stop", args)
  }

  return
}

// scrollerProperties()
func (this *QScroller) Scrollerproperties(args ...interface{}) (ret interface{}) {
  // scrollerProperties()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller18scrollerPropertiesEv
    // invoke: QScrollerProperties scrollerProperties()
    var ret0 = C.C_ZNK9QScroller18scrollerPropertiesEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QScrollerProperties{}) // "QScrollerProperties"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScroller", "scrollerProperties", args)
  }

  return
}

// finalPosition()
func (this *QScroller) Finalposition(args ...interface{}) (ret interface{}) {
  // finalPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller13finalPositionEv
    // invoke: QPointF finalPosition()
    var ret0 = C.C_ZNK9QScroller13finalPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScroller", "finalPosition", args)
  }

  return
}

// ensureVisible(const class QRectF &, qreal, qreal, int)
func (this *QScroller) Ensurevisible(args ...interface{}) () {
  // ensureVisible(const class QRectF &, qreal, qreal, int)
  // ensureVisible(const class QRectF &, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller13ensureVisibleERK6QRectFddi
    // invoke: void ensureVisible(const class QRectF &, qreal, qreal, int)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN9QScroller13ensureVisibleERK6QRectFddi(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN9QScroller13ensureVisibleERK6QRectFdd
    // invoke: void ensureVisible(const class QRectF &, qreal, qreal)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    C.C_ZN9QScroller13ensureVisibleERK6QRectFdd(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QScroller", "ensureVisible", args)
  }

  return
}

// pixelPerMeter()
func (this *QScroller) Pixelpermeter(args ...interface{}) (ret interface{}) {
  // pixelPerMeter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller13pixelPerMeterEv
    // invoke: QPointF pixelPerMeter()
    var ret0 = C.C_ZNK9QScroller13pixelPerMeterEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScroller", "pixelPerMeter", args)
  }

  return
}

// metaObject()
func (this *QScroller) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QScroller10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "metaObject", args)
  }

  return
}

// target()
func (this *QScroller) Target(args ...interface{}) (ret interface{}) {
  // target()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller6targetEv
    // invoke: QObject * target()
    var ret0 = C.C_ZNK9QScroller6targetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScroller", "target", args)
  }

  return
}

// grabbedGesture(class QObject *)
func (this *QScroller) Grabbedgesture_S(args ...interface{}) () {
  // grabbedGesture(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller14grabbedGestureEP7QObject
    // invoke: Qt::GestureType grabbedGesture(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QScroller14grabbedGestureEP7QObject(arg0)
  default:
    qtrt.ErrorResolve("QScroller", "grabbedGesture", args)
  }

  return
}

// setScrollerProperties(const class QScrollerProperties &)
func (this *QScroller) Setscrollerproperties(args ...interface{}) () {
  // setScrollerProperties(const class QScrollerProperties &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScrollerProperties{}) // "const QScrollerProperties &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties
    // invoke: void setScrollerProperties(const class QScrollerProperties &)
    var arg0 = args[0].(*QScrollerProperties).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScroller", "setScrollerProperties", args)
  }

  return
}

// velocity()
func (this *QScroller) Velocity(args ...interface{}) (ret interface{}) {
  // velocity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QScroller8velocityEv
    // invoke: QPointF velocity()
    var ret0 = C.C_ZNK9QScroller8velocityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScroller", "velocity", args)
  }

  return
}

// <= body block end

