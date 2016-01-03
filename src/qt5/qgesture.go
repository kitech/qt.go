package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qgesture.h
// dst-file: /src/widgets/qgesture.go
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
  // proto:  void QSwipeGesture::~QSwipeGesture();
extern void _ZN13QSwipeGestureD0Ev(void* qthis);
  // proto:  const QMetaObject * QSwipeGesture::metaObject();
extern void _ZNK13QSwipeGesture10metaObjectEv(void* qthis);
  // proto:  void QSwipeGesture::setSwipeAngle(qreal value);
extern void _ZN13QSwipeGesture13setSwipeAngleEd(void* qthis, double arg0);
  // proto:  void QSwipeGesture::QSwipeGesture(QObject * parent);
extern void* dector_ZN13QSwipeGestureC1EP7QObject(void* arg0);
extern void _ZN13QSwipeGestureC1EP7QObject(void* qthis, void* arg0);
  // proto:  qreal QSwipeGesture::swipeAngle();
extern void _ZNK13QSwipeGesture10swipeAngleEv(void* qthis);
  // proto:  QPointF QGesture::hotSpot();
extern void _ZNK8QGesture7hotSpotEv(void* qthis);
  // proto:  bool QGesture::hasHotSpot();
extern void _ZNK8QGesture10hasHotSpotEv(void* qthis);
  // proto:  void QGesture::unsetHotSpot();
extern void _ZN8QGesture12unsetHotSpotEv(void* qthis);
  // proto:  const QMetaObject * QGesture::metaObject();
extern void _ZNK8QGesture10metaObjectEv(void* qthis);
  // proto:  void QGesture::QGesture(QObject * parent);
extern void* dector_ZN8QGestureC1EP7QObject(void* arg0);
extern void _ZN8QGestureC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QGesture::setHotSpot(const QPointF & value);
extern void _ZN8QGesture10setHotSpotERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGesture::~QGesture();
extern void _ZN8QGestureD0Ev(void* qthis);
  // proto:  bool QGestureEvent::isAccepted(QGesture * );
extern void _ZNK13QGestureEvent10isAcceptedEP8QGesture(void* qthis, void* arg0);
  // proto:  QWidget * QGestureEvent::widget();
extern void _ZNK13QGestureEvent6widgetEv(void* qthis);
  // proto:  void QGestureEvent::ignore(QGesture * );
extern void _ZN13QGestureEvent6ignoreEP8QGesture(void* qthis, void* arg0);
  // proto:  void QGestureEvent::accept(QGesture * );
extern void _ZN13QGestureEvent6acceptEP8QGesture(void* qthis, void* arg0);
  // proto:  QList<QGesture *> QGestureEvent::activeGestures();
extern void _ZNK13QGestureEvent14activeGesturesEv(void* qthis);
  // proto:  QList<QGesture *> QGestureEvent::gestures();
extern void _ZNK13QGestureEvent8gesturesEv(void* qthis);
  // proto:  void QGestureEvent::setAccepted(QGesture * , bool );
extern void _ZN13QGestureEvent11setAcceptedEP8QGestureb(void* qthis, void* arg0, bool arg1);
  // proto:  void QGestureEvent::setWidget(QWidget * widget);
extern void _ZN13QGestureEvent9setWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QGestureEvent::~QGestureEvent();
extern void _ZN13QGestureEventD0Ev(void* qthis);
  // proto:  QList<QGesture *> QGestureEvent::canceledGestures();
extern void _ZNK13QGestureEvent16canceledGesturesEv(void* qthis);
  // proto:  QPointF QGestureEvent::mapToGraphicsScene(const QPointF & gesturePoint);
extern void _ZNK13QGestureEvent18mapToGraphicsSceneERK7QPointF(void* qthis, void* arg0);
  // proto:  QPointF QPanGesture::offset();
extern void _ZNK11QPanGesture6offsetEv(void* qthis);
  // proto:  QPointF QPanGesture::delta();
extern void _ZNK11QPanGesture5deltaEv(void* qthis);
  // proto:  void QPanGesture::setOffset(const QPointF & value);
extern void _ZN11QPanGesture9setOffsetERK7QPointF(void* qthis, void* arg0);
  // proto:  qreal QPanGesture::acceleration();
extern void _ZNK11QPanGesture12accelerationEv(void* qthis);
  // proto:  void QPanGesture::~QPanGesture();
extern void _ZN11QPanGestureD0Ev(void* qthis);
  // proto:  void QPanGesture::QPanGesture(QObject * parent);
extern void* dector_ZN11QPanGestureC1EP7QObject(void* arg0);
extern void _ZN11QPanGestureC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QPanGesture::metaObject();
extern void _ZNK11QPanGesture10metaObjectEv(void* qthis);
  // proto:  void QPanGesture::setAcceleration(qreal value);
extern void _ZN11QPanGesture15setAccelerationEd(void* qthis, double arg0);
  // proto:  QPointF QPanGesture::lastOffset();
extern void _ZNK11QPanGesture10lastOffsetEv(void* qthis);
  // proto:  void QPanGesture::setLastOffset(const QPointF & value);
extern void _ZN11QPanGesture13setLastOffsetERK7QPointF(void* qthis, void* arg0);
  // proto:  void QTapAndHoldGesture::QTapAndHoldGesture(QObject * parent);
extern void* dector_ZN18QTapAndHoldGestureC1EP7QObject(void* arg0);
extern void _ZN18QTapAndHoldGestureC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QTapAndHoldGesture::~QTapAndHoldGesture();
extern void _ZN18QTapAndHoldGestureD0Ev(void* qthis);
  // proto:  QPointF QTapAndHoldGesture::position();
extern void _ZNK18QTapAndHoldGesture8positionEv(void* qthis);
  // proto: static void QTapAndHoldGesture::setTimeout(int msecs);
extern void _ZN18QTapAndHoldGesture10setTimeoutEi(int32_t arg0);
  // proto: static int QTapAndHoldGesture::timeout();
extern void _ZN18QTapAndHoldGesture7timeoutEv();
  // proto:  const QMetaObject * QTapAndHoldGesture::metaObject();
extern void _ZNK18QTapAndHoldGesture10metaObjectEv(void* qthis);
  // proto:  void QTapAndHoldGesture::setPosition(const QPointF & pos);
extern void _ZN18QTapAndHoldGesture11setPositionERK7QPointF(void* qthis, void* arg0);
  // proto:  QPointF QTapGesture::position();
extern void _ZNK11QTapGesture8positionEv(void* qthis);
  // proto:  void QTapGesture::setPosition(const QPointF & pos);
extern void _ZN11QTapGesture11setPositionERK7QPointF(void* qthis, void* arg0);
  // proto:  void QTapGesture::QTapGesture(QObject * parent);
extern void* dector_ZN11QTapGestureC1EP7QObject(void* arg0);
extern void _ZN11QTapGestureC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QTapGesture::metaObject();
extern void _ZNK11QTapGesture10metaObjectEv(void* qthis);
  // proto:  void QTapGesture::~QTapGesture();
extern void _ZN11QTapGestureD0Ev(void* qthis);
  // proto:  void QPinchGesture::setRotationAngle(qreal value);
extern void _ZN13QPinchGesture16setRotationAngleEd(void* qthis, double arg0);
  // proto:  qreal QPinchGesture::lastScaleFactor();
extern void _ZNK13QPinchGesture15lastScaleFactorEv(void* qthis);
  // proto:  qreal QPinchGesture::lastRotationAngle();
extern void _ZNK13QPinchGesture17lastRotationAngleEv(void* qthis);
  // proto:  QPointF QPinchGesture::startCenterPoint();
extern void _ZNK13QPinchGesture16startCenterPointEv(void* qthis);
  // proto:  qreal QPinchGesture::rotationAngle();
extern void _ZNK13QPinchGesture13rotationAngleEv(void* qthis);
  // proto:  QPointF QPinchGesture::lastCenterPoint();
extern void _ZNK13QPinchGesture15lastCenterPointEv(void* qthis);
  // proto:  void QPinchGesture::QPinchGesture(QObject * parent);
extern void* dector_ZN13QPinchGestureC1EP7QObject(void* arg0);
extern void _ZN13QPinchGestureC1EP7QObject(void* qthis, void* arg0);
  // proto:  qreal QPinchGesture::totalScaleFactor();
extern void _ZNK13QPinchGesture16totalScaleFactorEv(void* qthis);
  // proto:  void QPinchGesture::setTotalScaleFactor(qreal value);
extern void _ZN13QPinchGesture19setTotalScaleFactorEd(void* qthis, double arg0);
  // proto:  qreal QPinchGesture::totalRotationAngle();
extern void _ZNK13QPinchGesture18totalRotationAngleEv(void* qthis);
  // proto:  void QPinchGesture::setLastScaleFactor(qreal value);
extern void _ZN13QPinchGesture18setLastScaleFactorEd(void* qthis, double arg0);
  // proto:  void QPinchGesture::setLastCenterPoint(const QPointF & value);
extern void _ZN13QPinchGesture18setLastCenterPointERK7QPointF(void* qthis, void* arg0);
  // proto:  const QMetaObject * QPinchGesture::metaObject();
extern void _ZNK13QPinchGesture10metaObjectEv(void* qthis);
  // proto:  void QPinchGesture::setLastRotationAngle(qreal value);
extern void _ZN13QPinchGesture20setLastRotationAngleEd(void* qthis, double arg0);
  // proto:  QPointF QPinchGesture::centerPoint();
extern void _ZNK13QPinchGesture11centerPointEv(void* qthis);
  // proto:  void QPinchGesture::setCenterPoint(const QPointF & value);
extern void _ZN13QPinchGesture14setCenterPointERK7QPointF(void* qthis, void* arg0);
  // proto:  void QPinchGesture::setTotalRotationAngle(qreal value);
extern void _ZN13QPinchGesture21setTotalRotationAngleEd(void* qthis, double arg0);
  // proto:  void QPinchGesture::setScaleFactor(qreal value);
extern void _ZN13QPinchGesture14setScaleFactorEd(void* qthis, double arg0);
  // proto:  void QPinchGesture::~QPinchGesture();
extern void _ZN13QPinchGestureD0Ev(void* qthis);
  // proto:  void QPinchGesture::setStartCenterPoint(const QPointF & value);
extern void _ZN13QPinchGesture19setStartCenterPointERK7QPointF(void* qthis, void* arg0);
  // proto:  qreal QPinchGesture::scaleFactor();
extern void _ZNK13QPinchGesture11scaleFactorEv(void* qthis);
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

// class sizeof(QSwipeGesture)=1
type QSwipeGesture struct {
  /*qbase*/ QGesture;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGesture)=1
type QGesture struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGestureEvent)=1
type QGestureEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPanGesture)=1
type QPanGesture struct {
  /*qbase*/ QGesture;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTapAndHoldGesture)=1
type QTapAndHoldGesture struct {
  /*qbase*/ QGesture;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTapGesture)=1
type QTapGesture struct {
  /*qbase*/ QGesture;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPinchGesture)=1
type QPinchGesture struct {
  /*qbase*/ QGesture;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QSwipeGesture::~QSwipeGesture();
func (this *QSwipeGesture) FreeQSwipeGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSwipeGesture", "~QSwipeGesture", args)
  }

}

  // proto:  const QMetaObject * QSwipeGesture::metaObject();
func (this *QSwipeGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSwipeGesture10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QSwipeGesture10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSwipeGesture", "metaObject", args)
  }

}

  // proto:  void QSwipeGesture::setSwipeAngle(qreal value);
func (this *QSwipeGesture) setSwipeAngle(args ...interface{}) () {
  // setSwipeAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSwipeGesture13setSwipeAngleEd
    // invoke: void setSwipeAngle(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QSwipeGesture13setSwipeAngleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSwipeGesture", "setSwipeAngle", args)
  }

}

  // proto:  void QSwipeGesture::QSwipeGesture(QObject * parent);
func NewQSwipeGesture(args ...interface{}) QSwipeGesture {
  return QSwipeGesture{}
}

  // proto:  qreal QSwipeGesture::swipeAngle();
func (this *QSwipeGesture) swipeAngle(args ...interface{}) () {
  // swipeAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSwipeGesture10swipeAngleEv
    // invoke: qreal swipeAngle()
    C._ZNK13QSwipeGesture10swipeAngleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSwipeGesture", "swipeAngle", args)
  }

}

  // proto:  QPointF QGesture::hotSpot();
func (this *QGesture) hotSpot(args ...interface{}) () {
  // hotSpot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QGesture7hotSpotEv
    // invoke: QPointF hotSpot()
    C._ZNK8QGesture7hotSpotEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGesture", "hotSpot", args)
  }

}

  // proto:  bool QGesture::hasHotSpot();
func (this *QGesture) hasHotSpot(args ...interface{}) () {
  // hasHotSpot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QGesture10hasHotSpotEv
    // invoke: bool hasHotSpot()
    C._ZNK8QGesture10hasHotSpotEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGesture", "hasHotSpot", args)
  }

}

  // proto:  void QGesture::unsetHotSpot();
func (this *QGesture) unsetHotSpot(args ...interface{}) () {
  // unsetHotSpot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QGesture12unsetHotSpotEv
    // invoke: void unsetHotSpot()
    C._ZN8QGesture12unsetHotSpotEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGesture", "unsetHotSpot", args)
  }

}

  // proto:  const QMetaObject * QGesture::metaObject();
func (this *QGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QGesture10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK8QGesture10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGesture", "metaObject", args)
  }

}

  // proto:  void QGesture::QGesture(QObject * parent);
func NewQGesture(args ...interface{}) QGesture {
  return QGesture{}
}

  // proto:  void QGesture::setHotSpot(const QPointF & value);
func (this *QGesture) setHotSpot(args ...interface{}) () {
  // setHotSpot(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QGesture10setHotSpotERK7QPointF
    // invoke: void setHotSpot(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QGesture10setHotSpotERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGesture", "setHotSpot", args)
  }

}

  // proto:  void QGesture::~QGesture();
func (this *QGesture) FreeQGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGesture", "~QGesture", args)
  }

}

  // proto:  bool QGestureEvent::isAccepted(QGesture * );
func (this *QGestureEvent) isAccepted(args ...interface{}) () {
  // isAccepted(class QGesture *)
  // isAccepted(Qt::GestureType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::GestureType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent10isAcceptedEP8QGesture
    // invoke: bool isAccepted(class QGesture *)
    var arg0 = args[0].(QGesture).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGestureEvent10isAcceptedEP8QGesture(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureEvent", "isAccepted", args)
  }

}

  // proto:  QWidget * QGestureEvent::widget();
func (this *QGestureEvent) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent6widgetEv
    // invoke: QWidget * widget()
    C._ZNK13QGestureEvent6widgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGestureEvent", "widget", args)
  }

}

  // proto:  void QGestureEvent::ignore(QGesture * );
func (this *QGestureEvent) ignore(args ...interface{}) () {
  // ignore(class QGesture *)
  // ignore(Qt::GestureType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::GestureType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent6ignoreEP8QGesture
    // invoke: void ignore(class QGesture *)
    var arg0 = args[0].(QGesture).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QGestureEvent6ignoreEP8QGesture(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureEvent", "ignore", args)
  }

}

  // proto:  void QGestureEvent::accept(QGesture * );
func (this *QGestureEvent) accept(args ...interface{}) () {
  // accept(class QGesture *)
  // accept(Qt::GestureType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::GestureType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent6acceptEP8QGesture
    // invoke: void accept(class QGesture *)
    var arg0 = args[0].(QGesture).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QGestureEvent6acceptEP8QGesture(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureEvent", "accept", args)
  }

}

  // proto:  QList<QGesture *> QGestureEvent::activeGestures();
func (this *QGestureEvent) activeGestures(args ...interface{}) () {
  // activeGestures()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent14activeGesturesEv
    // invoke: QList<QGesture *> activeGestures()
    C._ZNK13QGestureEvent14activeGesturesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGestureEvent", "activeGestures", args)
  }

}

  // proto:  QList<QGesture *> QGestureEvent::gestures();
func (this *QGestureEvent) gestures(args ...interface{}) () {
  // gestures()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent8gesturesEv
    // invoke: QList<QGesture *> gestures()
    C._ZNK13QGestureEvent8gesturesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGestureEvent", "gestures", args)
  }

}

  // proto:  void QGestureEvent::setAccepted(QGesture * , bool );
func (this *QGestureEvent) setAccepted(args ...interface{}) () {
  // setAccepted(Qt::GestureType, _Bool)
  // setAccepted(class QGesture *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "Qt::GestureType"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QGesture{}) // "QGesture *"
  vtys[1][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent11setAcceptedEP8QGestureb
    // invoke: void setAccepted(class QGesture *, _Bool)
    var arg0 = args[0].(QGesture).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN13QGestureEvent11setAcceptedEP8QGestureb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGestureEvent", "setAccepted", args)
  }

}

  // proto:  void QGestureEvent::setWidget(QWidget * widget);
func (this *QGestureEvent) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent9setWidgetEP7QWidget
    // invoke: void setWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QGestureEvent9setWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureEvent", "setWidget", args)
  }

}

  // proto:  void QGestureEvent::~QGestureEvent();
func (this *QGestureEvent) FreeQGestureEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGestureEvent", "~QGestureEvent", args)
  }

}

  // proto:  QList<QGesture *> QGestureEvent::canceledGestures();
func (this *QGestureEvent) canceledGestures(args ...interface{}) () {
  // canceledGestures()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent16canceledGesturesEv
    // invoke: QList<QGesture *> canceledGestures()
    C._ZNK13QGestureEvent16canceledGesturesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGestureEvent", "canceledGestures", args)
  }

}

  // proto:  QPointF QGestureEvent::mapToGraphicsScene(const QPointF & gesturePoint);
func (this *QGestureEvent) mapToGraphicsScene(args ...interface{}) () {
  // mapToGraphicsScene(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent18mapToGraphicsSceneERK7QPointF
    // invoke: QPointF mapToGraphicsScene(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QGestureEvent18mapToGraphicsSceneERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureEvent", "mapToGraphicsScene", args)
  }

}

  // proto:  QPointF QPanGesture::offset();
func (this *QPanGesture) offset(args ...interface{}) () {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPanGesture6offsetEv
    // invoke: QPointF offset()
    C._ZNK11QPanGesture6offsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPanGesture", "offset", args)
  }

}

  // proto:  QPointF QPanGesture::delta();
func (this *QPanGesture) delta(args ...interface{}) () {
  // delta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPanGesture5deltaEv
    // invoke: QPointF delta()
    C._ZNK11QPanGesture5deltaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPanGesture", "delta", args)
  }

}

  // proto:  void QPanGesture::setOffset(const QPointF & value);
func (this *QPanGesture) setOffset(args ...interface{}) () {
  // setOffset(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPanGesture9setOffsetERK7QPointF
    // invoke: void setOffset(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QPanGesture9setOffsetERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPanGesture", "setOffset", args)
  }

}

  // proto:  qreal QPanGesture::acceleration();
func (this *QPanGesture) acceleration(args ...interface{}) () {
  // acceleration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPanGesture12accelerationEv
    // invoke: qreal acceleration()
    C._ZNK11QPanGesture12accelerationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPanGesture", "acceleration", args)
  }

}

  // proto:  void QPanGesture::~QPanGesture();
func (this *QPanGesture) FreeQPanGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPanGesture", "~QPanGesture", args)
  }

}

  // proto:  void QPanGesture::QPanGesture(QObject * parent);
func NewQPanGesture(args ...interface{}) QPanGesture {
  return QPanGesture{}
}

  // proto:  const QMetaObject * QPanGesture::metaObject();
func (this *QPanGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPanGesture10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QPanGesture10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPanGesture", "metaObject", args)
  }

}

  // proto:  void QPanGesture::setAcceleration(qreal value);
func (this *QPanGesture) setAcceleration(args ...interface{}) () {
  // setAcceleration(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPanGesture15setAccelerationEd
    // invoke: void setAcceleration(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN11QPanGesture15setAccelerationEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPanGesture", "setAcceleration", args)
  }

}

  // proto:  QPointF QPanGesture::lastOffset();
func (this *QPanGesture) lastOffset(args ...interface{}) () {
  // lastOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPanGesture10lastOffsetEv
    // invoke: QPointF lastOffset()
    C._ZNK11QPanGesture10lastOffsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPanGesture", "lastOffset", args)
  }

}

  // proto:  void QPanGesture::setLastOffset(const QPointF & value);
func (this *QPanGesture) setLastOffset(args ...interface{}) () {
  // setLastOffset(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPanGesture13setLastOffsetERK7QPointF
    // invoke: void setLastOffset(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QPanGesture13setLastOffsetERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPanGesture", "setLastOffset", args)
  }

}

  // proto:  void QTapAndHoldGesture::QTapAndHoldGesture(QObject * parent);
func NewQTapAndHoldGesture(args ...interface{}) QTapAndHoldGesture {
  return QTapAndHoldGesture{}
}

  // proto:  void QTapAndHoldGesture::~QTapAndHoldGesture();
func (this *QTapAndHoldGesture) FreeQTapAndHoldGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "~QTapAndHoldGesture", args)
  }

}

  // proto:  QPointF QTapAndHoldGesture::position();
func (this *QTapAndHoldGesture) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QTapAndHoldGesture8positionEv
    // invoke: QPointF position()
    C._ZNK18QTapAndHoldGesture8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "position", args)
  }

}

  // proto: static void QTapAndHoldGesture::setTimeout(int msecs);
func (this *QTapAndHoldGesture) setTimeout_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "setTimeout", args)
  }

}

  // proto: static int QTapAndHoldGesture::timeout();
func (this *QTapAndHoldGesture) timeout_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "timeout", args)
  }

}

  // proto:  const QMetaObject * QTapAndHoldGesture::metaObject();
func (this *QTapAndHoldGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QTapAndHoldGesture10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK18QTapAndHoldGesture10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "metaObject", args)
  }

}

  // proto:  void QTapAndHoldGesture::setPosition(const QPointF & pos);
func (this *QTapAndHoldGesture) setPosition(args ...interface{}) () {
  // setPosition(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QTapAndHoldGesture11setPositionERK7QPointF
    // invoke: void setPosition(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QTapAndHoldGesture11setPositionERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "setPosition", args)
  }

}

  // proto:  QPointF QTapGesture::position();
func (this *QTapGesture) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTapGesture8positionEv
    // invoke: QPointF position()
    C._ZNK11QTapGesture8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTapGesture", "position", args)
  }

}

  // proto:  void QTapGesture::setPosition(const QPointF & pos);
func (this *QTapGesture) setPosition(args ...interface{}) () {
  // setPosition(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTapGesture11setPositionERK7QPointF
    // invoke: void setPosition(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QTapGesture11setPositionERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTapGesture", "setPosition", args)
  }

}

  // proto:  void QTapGesture::QTapGesture(QObject * parent);
func NewQTapGesture(args ...interface{}) QTapGesture {
  return QTapGesture{}
}

  // proto:  const QMetaObject * QTapGesture::metaObject();
func (this *QTapGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTapGesture10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QTapGesture10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTapGesture", "metaObject", args)
  }

}

  // proto:  void QTapGesture::~QTapGesture();
func (this *QTapGesture) FreeQTapGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTapGesture", "~QTapGesture", args)
  }

}

  // proto:  void QPinchGesture::setRotationAngle(qreal value);
func (this *QPinchGesture) setRotationAngle(args ...interface{}) () {
  // setRotationAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture16setRotationAngleEd
    // invoke: void setRotationAngle(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QPinchGesture16setRotationAngleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setRotationAngle", args)
  }

}

  // proto:  qreal QPinchGesture::lastScaleFactor();
func (this *QPinchGesture) lastScaleFactor(args ...interface{}) () {
  // lastScaleFactor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture15lastScaleFactorEv
    // invoke: qreal lastScaleFactor()
    C._ZNK13QPinchGesture15lastScaleFactorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "lastScaleFactor", args)
  }

}

  // proto:  qreal QPinchGesture::lastRotationAngle();
func (this *QPinchGesture) lastRotationAngle(args ...interface{}) () {
  // lastRotationAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture17lastRotationAngleEv
    // invoke: qreal lastRotationAngle()
    C._ZNK13QPinchGesture17lastRotationAngleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "lastRotationAngle", args)
  }

}

  // proto:  QPointF QPinchGesture::startCenterPoint();
func (this *QPinchGesture) startCenterPoint(args ...interface{}) () {
  // startCenterPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture16startCenterPointEv
    // invoke: QPointF startCenterPoint()
    C._ZNK13QPinchGesture16startCenterPointEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "startCenterPoint", args)
  }

}

  // proto:  qreal QPinchGesture::rotationAngle();
func (this *QPinchGesture) rotationAngle(args ...interface{}) () {
  // rotationAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture13rotationAngleEv
    // invoke: qreal rotationAngle()
    C._ZNK13QPinchGesture13rotationAngleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "rotationAngle", args)
  }

}

  // proto:  QPointF QPinchGesture::lastCenterPoint();
func (this *QPinchGesture) lastCenterPoint(args ...interface{}) () {
  // lastCenterPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture15lastCenterPointEv
    // invoke: QPointF lastCenterPoint()
    C._ZNK13QPinchGesture15lastCenterPointEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "lastCenterPoint", args)
  }

}

  // proto:  void QPinchGesture::QPinchGesture(QObject * parent);
func NewQPinchGesture(args ...interface{}) QPinchGesture {
  return QPinchGesture{}
}

  // proto:  qreal QPinchGesture::totalScaleFactor();
func (this *QPinchGesture) totalScaleFactor(args ...interface{}) () {
  // totalScaleFactor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture16totalScaleFactorEv
    // invoke: qreal totalScaleFactor()
    C._ZNK13QPinchGesture16totalScaleFactorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "totalScaleFactor", args)
  }

}

  // proto:  void QPinchGesture::setTotalScaleFactor(qreal value);
func (this *QPinchGesture) setTotalScaleFactor(args ...interface{}) () {
  // setTotalScaleFactor(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture19setTotalScaleFactorEd
    // invoke: void setTotalScaleFactor(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QPinchGesture19setTotalScaleFactorEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setTotalScaleFactor", args)
  }

}

  // proto:  qreal QPinchGesture::totalRotationAngle();
func (this *QPinchGesture) totalRotationAngle(args ...interface{}) () {
  // totalRotationAngle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture18totalRotationAngleEv
    // invoke: qreal totalRotationAngle()
    C._ZNK13QPinchGesture18totalRotationAngleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "totalRotationAngle", args)
  }

}

  // proto:  void QPinchGesture::setLastScaleFactor(qreal value);
func (this *QPinchGesture) setLastScaleFactor(args ...interface{}) () {
  // setLastScaleFactor(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture18setLastScaleFactorEd
    // invoke: void setLastScaleFactor(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QPinchGesture18setLastScaleFactorEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setLastScaleFactor", args)
  }

}

  // proto:  void QPinchGesture::setLastCenterPoint(const QPointF & value);
func (this *QPinchGesture) setLastCenterPoint(args ...interface{}) () {
  // setLastCenterPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture18setLastCenterPointERK7QPointF
    // invoke: void setLastCenterPoint(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QPinchGesture18setLastCenterPointERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setLastCenterPoint", args)
  }

}

  // proto:  const QMetaObject * QPinchGesture::metaObject();
func (this *QPinchGesture) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QPinchGesture10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "metaObject", args)
  }

}

  // proto:  void QPinchGesture::setLastRotationAngle(qreal value);
func (this *QPinchGesture) setLastRotationAngle(args ...interface{}) () {
  // setLastRotationAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture20setLastRotationAngleEd
    // invoke: void setLastRotationAngle(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QPinchGesture20setLastRotationAngleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setLastRotationAngle", args)
  }

}

  // proto:  QPointF QPinchGesture::centerPoint();
func (this *QPinchGesture) centerPoint(args ...interface{}) () {
  // centerPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture11centerPointEv
    // invoke: QPointF centerPoint()
    C._ZNK13QPinchGesture11centerPointEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "centerPoint", args)
  }

}

  // proto:  void QPinchGesture::setCenterPoint(const QPointF & value);
func (this *QPinchGesture) setCenterPoint(args ...interface{}) () {
  // setCenterPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture14setCenterPointERK7QPointF
    // invoke: void setCenterPoint(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QPinchGesture14setCenterPointERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setCenterPoint", args)
  }

}

  // proto:  void QPinchGesture::setTotalRotationAngle(qreal value);
func (this *QPinchGesture) setTotalRotationAngle(args ...interface{}) () {
  // setTotalRotationAngle(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture21setTotalRotationAngleEd
    // invoke: void setTotalRotationAngle(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QPinchGesture21setTotalRotationAngleEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setTotalRotationAngle", args)
  }

}

  // proto:  void QPinchGesture::setScaleFactor(qreal value);
func (this *QPinchGesture) setScaleFactor(args ...interface{}) () {
  // setScaleFactor(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture14setScaleFactorEd
    // invoke: void setScaleFactor(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QPinchGesture14setScaleFactorEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setScaleFactor", args)
  }

}

  // proto:  void QPinchGesture::~QPinchGesture();
func (this *QPinchGesture) FreeQPinchGesture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPinchGesture", "~QPinchGesture", args)
  }

}

  // proto:  void QPinchGesture::setStartCenterPoint(const QPointF & value);
func (this *QPinchGesture) setStartCenterPoint(args ...interface{}) () {
  // setStartCenterPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture19setStartCenterPointERK7QPointF
    // invoke: void setStartCenterPoint(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QPinchGesture19setStartCenterPointERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setStartCenterPoint", args)
  }

}

  // proto:  qreal QPinchGesture::scaleFactor();
func (this *QPinchGesture) scaleFactor(args ...interface{}) () {
  // scaleFactor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture11scaleFactorEv
    // invoke: qreal scaleFactor()
    C._ZNK13QPinchGesture11scaleFactorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "scaleFactor", args)
  }

}

// <= body block end

