package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
extern void C_ZN9QScroller8scrollerEP7QObject(void* arg0); // 4
  // proto: static const QScroller * QScroller::scroller(const QObject * target);
extern void C_ZN9QScroller8scrollerEPK7QObject(void* arg0); // 4
  // proto:  void QScroller::scrollTo(const QPointF & pos, int scrollTime);
extern void C_ZN9QScroller8scrollToERK7QPointFi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QScroller::scrollTo(const QPointF & pos);
extern void C_ZN9QScroller8scrollToERK7QPointF(void* qthis, void* arg0); // 4
  // proto: static QList<QScroller *> QScroller::activeScrollers();
extern void C_ZN9QScroller15activeScrollersEv(); // 4
  // proto: static bool QScroller::hasScroller(QObject * target);
extern void C_ZN9QScroller11hasScrollerEP7QObject(void* arg0); // 4
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
extern void C_ZNK9QScroller18scrollerPropertiesEv(void* qthis); // 4
  // proto:  QPointF QScroller::finalPosition();
extern void C_ZNK9QScroller13finalPositionEv(void* qthis); // 4
  // proto:  void QScroller::ensureVisible(const QRectF & rect, qreal xmargin, qreal ymargin, int scrollTime);
extern void C_ZN9QScroller13ensureVisibleERK6QRectFddi(void* qthis, void* arg0, double arg1, double arg2, int32_t arg3); // 4
  // proto:  void QScroller::ensureVisible(const QRectF & rect, qreal xmargin, qreal ymargin);
extern void C_ZN9QScroller13ensureVisibleERK6QRectFdd(void* qthis, void* arg0, double arg1, double arg2); // 4
  // proto:  QPointF QScroller::pixelPerMeter();
extern void C_ZNK9QScroller13pixelPerMeterEv(void* qthis); // 4
  // proto:  const QMetaObject * QScroller::metaObject();
extern void C_ZNK9QScroller10metaObjectEv(void* qthis); // 4
  // proto:  QObject * QScroller::target();
extern void C_ZNK9QScroller6targetEv(void* qthis); // 4
  // proto: static Qt::GestureType QScroller::grabbedGesture(QObject * target);
extern void C_ZN9QScroller14grabbedGestureEP7QObject(void* arg0); // 4
  // proto:  void QScroller::setScrollerProperties(const QScrollerProperties & prop);
extern void C_ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties(void* qthis, void* arg0); // 4
  // proto:  QPointF QScroller::velocity();
extern void C_ZNK9QScroller8velocityEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _stateChanged QScroller_stateChanged_signal;
//  _scrollerPropertiesChanged QScroller_scrollerPropertiesChanged_signal;
}

// scroller(class QObject *)
func (this *QScroller) scroller_s(args ...interface{}) () {
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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QScroller8scrollerEP7QObject(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN9QScroller8scrollerEPK7QObject
    // invoke: const QScroller * scroller(const class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QScroller8scrollerEPK7QObject(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QScroller", "scroller", args)
  }

}

// scrollTo(const class QPointF &, int)
func (this *QScroller) scrollTo(args ...interface{}) () {
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
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QScroller8scrollToERK7QPointFi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN9QScroller8scrollToERK7QPointF
    // invoke: void scrollTo(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QScroller8scrollToERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScroller", "scrollTo", args)
  }

}

// activeScrollers()
func (this *QScroller) activeScrollers_s(args ...interface{}) () {
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

}

// hasScroller(class QObject *)
func (this *QScroller) hasScroller_s(args ...interface{}) () {
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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN9QScroller11hasScrollerEP7QObject(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QScroller", "hasScroller", args)
  }

}

// setSnapPositionsY(qreal, qreal)
func (this *QScroller) setSnapPositionsY(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN9QScroller17setSnapPositionsYEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QScroller", "setSnapPositionsY", args)
  }

}

// setSnapPositionsX(qreal, qreal)
func (this *QScroller) setSnapPositionsX(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN9QScroller17setSnapPositionsXEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QScroller", "setSnapPositionsX", args)
  }

}

// ungrabGesture(class QObject *)
func (this *QScroller) ungrabGesture_s(args ...interface{}) () {
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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QScroller13ungrabGestureEP7QObject(arg0)
  default:
    qtrt.ErrorResolve("QScroller", "ungrabGesture", args)
  }

}

// state()
func (this *QScroller) state(args ...interface{}) () {
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
    C.C_ZNK9QScroller5stateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "state", args)
  }

}

// resendPrepareEvent()
func (this *QScroller) resendPrepareEvent(args ...interface{}) () {
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
    C.C_ZN9QScroller18resendPrepareEventEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "resendPrepareEvent", args)
  }

}

// stop()
func (this *QScroller) stop(args ...interface{}) () {
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
    C.C_ZN9QScroller4stopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "stop", args)
  }

}

// scrollerProperties()
func (this *QScroller) scrollerProperties(args ...interface{}) () {
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
    var ret = C.C_ZNK9QScroller18scrollerPropertiesEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QScroller", "scrollerProperties", args)
  }

}

// finalPosition()
func (this *QScroller) finalPosition(args ...interface{}) () {
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
    var ret = C.C_ZNK9QScroller13finalPositionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QScroller", "finalPosition", args)
  }

}

// ensureVisible(const class QRectF &, qreal, qreal, int)
func (this *QScroller) ensureVisible(args ...interface{}) () {
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
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN9QScroller13ensureVisibleERK6QRectFddi(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN9QScroller13ensureVisibleERK6QRectFdd
    // invoke: void ensureVisible(const class QRectF &, qreal, qreal)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    C.C_ZN9QScroller13ensureVisibleERK6QRectFdd(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QScroller", "ensureVisible", args)
  }

}

// pixelPerMeter()
func (this *QScroller) pixelPerMeter(args ...interface{}) () {
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
    var ret = C.C_ZNK9QScroller13pixelPerMeterEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QScroller", "pixelPerMeter", args)
  }

}

// metaObject()
func (this *QScroller) metaObject(args ...interface{}) () {
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
    C.C_ZNK9QScroller10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "metaObject", args)
  }

}

// target()
func (this *QScroller) target(args ...interface{}) () {
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
    var ret = C.C_ZNK9QScroller6targetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QScroller", "target", args)
  }

}

// grabbedGesture(class QObject *)
func (this *QScroller) grabbedGesture_s(args ...interface{}) () {
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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QScroller14grabbedGestureEP7QObject(arg0)
  default:
    qtrt.ErrorResolve("QScroller", "grabbedGesture", args)
  }

}

// setScrollerProperties(const class QScrollerProperties &)
func (this *QScroller) setScrollerProperties(args ...interface{}) () {
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
    var arg0 = args[0].(QScrollerProperties).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScroller", "setScrollerProperties", args)
  }

}

// velocity()
func (this *QScroller) velocity(args ...interface{}) () {
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
    var ret = C.C_ZNK9QScroller8velocityEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QScroller", "velocity", args)
  }

}

// <= body block end

