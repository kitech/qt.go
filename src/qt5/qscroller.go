package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  void QScroller::~QScroller();
extern void _ZN9QScrollerD0Ev(void* qthis);
  // proto:  void QScroller::setSnapPositionsY(qreal first, qreal interval);
extern void _ZN9QScroller17setSnapPositionsYEdd(void* qthis, double arg0, double arg1);
  // proto:  QPointF QScroller::finalPosition();
extern void _ZNK9QScroller13finalPositionEv(void* qthis);
  // proto: static QList<QScroller *> QScroller::activeScrollers();
extern void _ZN9QScroller15activeScrollersEv();
  // proto:  void QScroller::setScrollerProperties(const QScrollerProperties & prop);
extern void _ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties(void* qthis, void* arg0);
  // proto:  const QMetaObject * QScroller::metaObject();
extern void _ZNK9QScroller10metaObjectEv(void* qthis);
  // proto:  QPointF QScroller::velocity();
extern void _ZNK9QScroller8velocityEv(void* qthis);
  // proto:  void QScroller::resendPrepareEvent();
extern void _ZN9QScroller18resendPrepareEventEv(void* qthis);
  // proto:  void QScroller::ensureVisible(const QRectF & rect, qreal xmargin, qreal ymargin);
extern void _ZN9QScroller13ensureVisibleERK6QRectFdd(void* qthis, void* arg0, double arg1, double arg2);
  // proto:  void QScroller::setSnapPositionsX(qreal first, qreal interval);
extern void _ZN9QScroller17setSnapPositionsXEdd(void* qthis, double arg0, double arg1);
  // proto: static bool QScroller::hasScroller(QObject * target);
extern void _ZN9QScroller11hasScrollerEP7QObject(void* arg0);
  // proto:  void QScroller::QScroller(const QScroller & );
extern void* dector_ZN9QScrollerC1ERKS_(void* arg0);
extern void _ZN9QScrollerC1ERKS_(void* qthis, void* arg0);
  // proto:  void QScroller::scrollTo(const QPointF & pos, int scrollTime);
extern void _ZN9QScroller8scrollToERK7QPointFi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QScroller::stop();
extern void _ZN9QScroller4stopEv(void* qthis);
  // proto:  void QScroller::ensureVisible(const QRectF & rect, qreal xmargin, qreal ymargin, int scrollTime);
extern void _ZN9QScroller13ensureVisibleERK6QRectFddi(void* qthis, void* arg0, double arg1, double arg2, int32_t arg3);
  // proto:  void QScroller::QScroller(QObject * target);
extern void* dector_ZN9QScrollerC1EP7QObject(void* arg0);
extern void _ZN9QScrollerC1EP7QObject(void* qthis, void* arg0);
  // proto: static void QScroller::ungrabGesture(QObject * target);
extern void _ZN9QScroller13ungrabGestureEP7QObject(void* arg0);
  // proto:  QScrollerProperties QScroller::scrollerProperties();
extern void _ZNK9QScroller18scrollerPropertiesEv(void* qthis);
  // proto: static QScroller * QScroller::scroller(QObject * target);
extern void _ZN9QScroller8scrollerEP7QObject(void* arg0);
  // proto: static const QScroller * QScroller::scroller(const QObject * target);
extern void _ZN9QScroller8scrollerEPK7QObject(void* arg0);
  // proto:  void QScroller::scrollTo(const QPointF & pos);
extern void _ZN9QScroller8scrollToERK7QPointF(void* qthis, void* arg0);
  // proto:  QObject * QScroller::target();
extern void _ZNK9QScroller6targetEv(void* qthis);
  // proto:  QPointF QScroller::pixelPerMeter();
extern void _ZNK9QScroller13pixelPerMeterEv(void* qthis);
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

  // proto:  void QScroller::~QScroller();
func (this *QScroller) FreeQScroller(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScroller", "~QScroller", args)
  }

}

  // proto:  void QScroller::setSnapPositionsY(qreal first, qreal interval);
func (this *QScroller) setSnapPositionsY(args ...interface{}) () {
  // setSnapPositionsY(qreal, qreal)
  // setSnapPositionsY(const QList<qreal> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  // vtys[1][0] = reflect.TypeOf(QList<double>{}) // "const QList<qreal> &"

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
    C._ZN9QScroller17setSnapPositionsYEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QScroller", "setSnapPositionsY", args)
  }

}

  // proto:  QPointF QScroller::finalPosition();
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
    C._ZNK9QScroller13finalPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "finalPosition", args)
  }

}

  // proto: static QList<QScroller *> QScroller::activeScrollers();
func (this *QScroller) activeScrollers_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScroller", "activeScrollers", args)
  }

}

  // proto:  void QScroller::setScrollerProperties(const QScrollerProperties & prop);
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
    C._ZN9QScroller21setScrollerPropertiesERK19QScrollerProperties(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScroller", "setScrollerProperties", args)
  }

}

  // proto:  const QMetaObject * QScroller::metaObject();
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
    C._ZNK9QScroller10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "metaObject", args)
  }

}

  // proto:  QPointF QScroller::velocity();
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
    C._ZNK9QScroller8velocityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "velocity", args)
  }

}

  // proto:  void QScroller::resendPrepareEvent();
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
    C._ZN9QScroller18resendPrepareEventEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "resendPrepareEvent", args)
  }

}

  // proto:  void QScroller::ensureVisible(const QRectF & rect, qreal xmargin, qreal ymargin);
func (this *QScroller) ensureVisible(args ...interface{}) () {
  // ensureVisible(const class QRectF &, qreal, qreal)
  // ensureVisible(const class QRectF &, qreal, qreal, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QScroller13ensureVisibleERK6QRectFdd
    // invoke: void ensureVisible(const class QRectF &, qreal, qreal)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    C._ZN9QScroller13ensureVisibleERK6QRectFdd(this.qclsinst, arg0, arg1, arg2)
  case 1:
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
    C._ZN9QScroller13ensureVisibleERK6QRectFddi(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QScroller", "ensureVisible", args)
  }

}

  // proto:  void QScroller::setSnapPositionsX(qreal first, qreal interval);
func (this *QScroller) setSnapPositionsX(args ...interface{}) () {
  // setSnapPositionsX(qreal, qreal)
  // setSnapPositionsX(const QList<qreal> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  // vtys[1][0] = reflect.TypeOf(QList<double>{}) // "const QList<qreal> &"

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
    C._ZN9QScroller17setSnapPositionsXEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QScroller", "setSnapPositionsX", args)
  }

}

  // proto: static bool QScroller::hasScroller(QObject * target);
func (this *QScroller) hasScroller_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScroller", "hasScroller", args)
  }

}

  // proto:  void QScroller::QScroller(const QScroller & );
func NewQScroller(args ...interface{}) QScroller {
  return QScroller{}
}

  // proto:  void QScroller::scrollTo(const QPointF & pos, int scrollTime);
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
    C._ZN9QScroller8scrollToERK7QPointFi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN9QScroller8scrollToERK7QPointF
    // invoke: void scrollTo(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QScroller8scrollToERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScroller", "scrollTo", args)
  }

}

  // proto:  void QScroller::stop();
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
    C._ZN9QScroller4stopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "stop", args)
  }

}

  // proto: static void QScroller::ungrabGesture(QObject * target);
func (this *QScroller) ungrabGesture_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScroller", "ungrabGesture", args)
  }

}

  // proto:  QScrollerProperties QScroller::scrollerProperties();
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
    C._ZNK9QScroller18scrollerPropertiesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "scrollerProperties", args)
  }

}

  // proto: static QScroller * QScroller::scroller(QObject * target);
func (this *QScroller) scroller_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScroller", "scroller", args)
  }

}

  // proto:  QObject * QScroller::target();
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
    C._ZNK9QScroller6targetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "target", args)
  }

}

  // proto:  QPointF QScroller::pixelPerMeter();
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
    C._ZNK9QScroller13pixelPerMeterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScroller", "pixelPerMeter", args)
  }

}

// <= body block end

