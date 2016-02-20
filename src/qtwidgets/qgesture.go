package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QSwipeGesture::QSwipeGesture(QObject * parent);
extern void* C_ZN13QSwipeGestureC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QSwipeGesture::metaObject();
extern void C_ZNK13QSwipeGesture10metaObjectEv(void* qthis); // 4
  // proto:  QSwipeGesture::SwipeDirection QSwipeGesture::horizontalDirection();
extern void C_ZNK13QSwipeGesture19horizontalDirectionEv(void* qthis); // 4
  // proto:  QSwipeGesture::SwipeDirection QSwipeGesture::verticalDirection();
extern void C_ZNK13QSwipeGesture17verticalDirectionEv(void* qthis); // 4
  // proto:  qreal QSwipeGesture::swipeAngle();
extern double C_ZNK13QSwipeGesture10swipeAngleEv(void* qthis); // 4
  // proto:  void QSwipeGesture::~QSwipeGesture();
extern void C_ZN13QSwipeGestureD2Ev(void* qthis); // 4
  // proto:  void QSwipeGesture::setSwipeAngle(qreal value);
extern void C_ZN13QSwipeGesture13setSwipeAngleEd(void* qthis, double arg0); // 4
  // proto:  const QMetaObject * QGesture::metaObject();
extern void C_ZNK8QGesture10metaObjectEv(void* qthis); // 4
  // proto:  bool QGesture::hasHotSpot();
extern bool C_ZNK8QGesture10hasHotSpotEv(void* qthis); // 4
  // proto:  QGesture::GestureCancelPolicy QGesture::gestureCancelPolicy();
extern void C_ZNK8QGesture19gestureCancelPolicyEv(void* qthis); // 4
  // proto:  Qt::GestureState QGesture::state();
extern void C_ZNK8QGesture5stateEv(void* qthis); // 4
  // proto:  void QGesture::setHotSpot(const QPointF & value);
extern void C_ZN8QGesture10setHotSpotERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPointF QGesture::hotSpot();
extern void* C_ZNK8QGesture7hotSpotEv(void* qthis); // 4
  // proto:  void QGesture::~QGesture();
extern void C_ZN8QGestureD2Ev(void* qthis); // 4
  // proto:  Qt::GestureType QGesture::gestureType();
extern void C_ZNK8QGesture11gestureTypeEv(void* qthis); // 4
  // proto:  void QGesture::unsetHotSpot();
extern void C_ZN8QGesture12unsetHotSpotEv(void* qthis); // 4
  // proto:  void QGesture::QGesture(QObject * parent);
extern void* C_ZN8QGestureC2EP7QObject(void* arg0); // 3
  // proto:  QList<QGesture *> QGestureEvent::activeGestures();
extern void C_ZNK13QGestureEvent14activeGesturesEv(void* qthis); // 4
  // proto:  QList<QGesture *> QGestureEvent::gestures();
extern void C_ZNK13QGestureEvent8gesturesEv(void* qthis); // 4
  // proto:  QWidget * QGestureEvent::widget();
extern void* C_ZNK13QGestureEvent6widgetEv(void* qthis); // 4
  // proto:  QList<QGesture *> QGestureEvent::canceledGestures();
extern void C_ZNK13QGestureEvent16canceledGesturesEv(void* qthis); // 4
  // proto:  void QGestureEvent::setWidget(QWidget * widget);
extern void C_ZN13QGestureEvent9setWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QGestureEvent::accept(QGesture * );
extern void C_ZN13QGestureEvent6acceptEP8QGesture(void* qthis, void* arg0); // 4
  // proto:  void QGestureEvent::ignore(QGesture * );
extern void C_ZN13QGestureEvent6ignoreEP8QGesture(void* qthis, void* arg0); // 4
  // proto:  bool QGestureEvent::isAccepted(QGesture * );
extern bool C_ZNK13QGestureEvent10isAcceptedEP8QGesture(void* qthis, void* arg0); // 4
  // proto:  void QGestureEvent::~QGestureEvent();
extern void C_ZN13QGestureEventD2Ev(void* qthis); // 4
  // proto:  void QGestureEvent::setAccepted(QGesture * , bool );
extern void C_ZN13QGestureEvent11setAcceptedEP8QGestureb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  QPointF QGestureEvent::mapToGraphicsScene(const QPointF & gesturePoint);
extern void* C_ZNK13QGestureEvent18mapToGraphicsSceneERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  qreal QPanGesture::acceleration();
extern double C_ZNK11QPanGesture12accelerationEv(void* qthis); // 4
  // proto:  void QPanGesture::setAcceleration(qreal value);
extern void C_ZN11QPanGesture15setAccelerationEd(void* qthis, double arg0); // 4
  // proto:  const QMetaObject * QPanGesture::metaObject();
extern void C_ZNK11QPanGesture10metaObjectEv(void* qthis); // 4
  // proto:  void QPanGesture::QPanGesture(QObject * parent);
extern void* C_ZN11QPanGestureC2EP7QObject(void* arg0); // 3
  // proto:  QPointF QPanGesture::offset();
extern void* C_ZNK11QPanGesture6offsetEv(void* qthis); // 4
  // proto:  void QPanGesture::~QPanGesture();
extern void C_ZN11QPanGestureD2Ev(void* qthis); // 4
  // proto:  QPointF QPanGesture::delta();
extern void* C_ZNK11QPanGesture5deltaEv(void* qthis); // 4
  // proto:  void QPanGesture::setOffset(const QPointF & value);
extern void C_ZN11QPanGesture9setOffsetERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  QPointF QPanGesture::lastOffset();
extern void* C_ZNK11QPanGesture10lastOffsetEv(void* qthis); // 4
  // proto:  void QPanGesture::setLastOffset(const QPointF & value);
extern void C_ZN11QPanGesture13setLastOffsetERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QTapAndHoldGesture::QTapAndHoldGesture(QObject * parent);
extern void* C_ZN18QTapAndHoldGestureC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QTapAndHoldGesture::metaObject();
extern void C_ZNK18QTapAndHoldGesture10metaObjectEv(void* qthis); // 4
  // proto:  void QTapAndHoldGesture::~QTapAndHoldGesture();
extern void C_ZN18QTapAndHoldGestureD2Ev(void* qthis); // 4
  // proto: static int QTapAndHoldGesture::timeout();
extern int32_t C_ZN18QTapAndHoldGesture7timeoutEv(); // 4
  // proto:  QPointF QTapAndHoldGesture::position();
extern void* C_ZNK18QTapAndHoldGesture8positionEv(void* qthis); // 4
  // proto:  void QTapAndHoldGesture::setPosition(const QPointF & pos);
extern void C_ZN18QTapAndHoldGesture11setPositionERK7QPointF(void* qthis, void* arg0); // 4
  // proto: static void QTapAndHoldGesture::setTimeout(int msecs);
extern void C_ZN18QTapAndHoldGesture10setTimeoutEi(int32_t arg0); // 4
  // proto:  const QMetaObject * QTapGesture::metaObject();
extern void C_ZNK11QTapGesture10metaObjectEv(void* qthis); // 4
  // proto:  void QTapGesture::QTapGesture(QObject * parent);
extern void* C_ZN11QTapGestureC2EP7QObject(void* arg0); // 3
  // proto:  void QTapGesture::~QTapGesture();
extern void C_ZN11QTapGestureD2Ev(void* qthis); // 4
  // proto:  QPointF QTapGesture::position();
extern void* C_ZNK11QTapGesture8positionEv(void* qthis); // 4
  // proto:  void QTapGesture::setPosition(const QPointF & pos);
extern void C_ZN11QTapGesture11setPositionERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  ChangeFlags QPinchGesture::changeFlags();
extern void C_ZNK13QPinchGesture11changeFlagsEv(void* qthis); // 4
  // proto:  QPointF QPinchGesture::startCenterPoint();
extern void* C_ZNK13QPinchGesture16startCenterPointEv(void* qthis); // 4
  // proto:  void QPinchGesture::setLastCenterPoint(const QPointF & value);
extern void C_ZN13QPinchGesture18setLastCenterPointERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPinchGesture::setStartCenterPoint(const QPointF & value);
extern void C_ZN13QPinchGesture19setStartCenterPointERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QPinchGesture::setTotalRotationAngle(qreal value);
extern void C_ZN13QPinchGesture21setTotalRotationAngleEd(void* qthis, double arg0); // 4
  // proto:  qreal QPinchGesture::lastRotationAngle();
extern double C_ZNK13QPinchGesture17lastRotationAngleEv(void* qthis); // 4
  // proto:  qreal QPinchGesture::rotationAngle();
extern double C_ZNK13QPinchGesture13rotationAngleEv(void* qthis); // 4
  // proto:  ChangeFlags QPinchGesture::totalChangeFlags();
extern void C_ZNK13QPinchGesture16totalChangeFlagsEv(void* qthis); // 4
  // proto:  void QPinchGesture::setTotalScaleFactor(qreal value);
extern void C_ZN13QPinchGesture19setTotalScaleFactorEd(void* qthis, double arg0); // 4
  // proto:  void QPinchGesture::setLastScaleFactor(qreal value);
extern void C_ZN13QPinchGesture18setLastScaleFactorEd(void* qthis, double arg0); // 4
  // proto:  void QPinchGesture::~QPinchGesture();
extern void C_ZN13QPinchGestureD2Ev(void* qthis); // 4
  // proto:  qreal QPinchGesture::totalRotationAngle();
extern double C_ZNK13QPinchGesture18totalRotationAngleEv(void* qthis); // 4
  // proto:  void QPinchGesture::setRotationAngle(qreal value);
extern void C_ZN13QPinchGesture16setRotationAngleEd(void* qthis, double arg0); // 4
  // proto:  void QPinchGesture::QPinchGesture(QObject * parent);
extern void* C_ZN13QPinchGestureC2EP7QObject(void* arg0); // 3
  // proto:  QPointF QPinchGesture::lastCenterPoint();
extern void* C_ZNK13QPinchGesture15lastCenterPointEv(void* qthis); // 4
  // proto:  QPointF QPinchGesture::centerPoint();
extern void* C_ZNK13QPinchGesture11centerPointEv(void* qthis); // 4
  // proto:  qreal QPinchGesture::totalScaleFactor();
extern double C_ZNK13QPinchGesture16totalScaleFactorEv(void* qthis); // 4
  // proto:  qreal QPinchGesture::lastScaleFactor();
extern double C_ZNK13QPinchGesture15lastScaleFactorEv(void* qthis); // 4
  // proto:  void QPinchGesture::setCenterPoint(const QPointF & value);
extern void C_ZN13QPinchGesture14setCenterPointERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  qreal QPinchGesture::scaleFactor();
extern double C_ZNK13QPinchGesture11scaleFactorEv(void* qthis); // 4
  // proto:  const QMetaObject * QPinchGesture::metaObject();
extern void C_ZNK13QPinchGesture10metaObjectEv(void* qthis); // 4
  // proto:  void QPinchGesture::setLastRotationAngle(qreal value);
extern void C_ZN13QPinchGesture20setLastRotationAngleEd(void* qthis, double arg0); // 4
  // proto:  void QPinchGesture::setScaleFactor(qreal value);
extern void C_ZN13QPinchGesture14setScaleFactorEd(void* qthis, double arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QSwipeGesture)=1
type QSwipeGesture struct {
  /*qbase*/ QGesture;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGesture)=1
type QGesture struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QGestureEvent)=1
type QGestureEvent struct {
  /*qbase*/ qtcore.QEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPanGesture)=1
type QPanGesture struct {
  /*qbase*/ QGesture;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTapAndHoldGesture)=1
type QTapAndHoldGesture struct {
  /*qbase*/ QGesture;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTapGesture)=1
type QTapGesture struct {
  /*qbase*/ QGesture;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPinchGesture)=1
type QPinchGesture struct {
  /*qbase*/ QGesture;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QSwipeGesture(class QObject *)
func NewQSwipeGesture(args ...interface{}) *QSwipeGesture {
  // QSwipeGesture(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSwipeGestureC1EP7QObject
    // invoke: void QSwipeGesture(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QSwipeGestureC2EP7QObject(arg0)
    return &QSwipeGesture{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSwipeGesture", "QSwipeGesture", args)
  }

  return nil // QSwipeGesture{Qclsinst:qthis}
}

// metaObject()
func (this *QSwipeGesture) Metaobject(args ...interface{}) () {
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
    C.C_ZNK13QSwipeGesture10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSwipeGesture", "metaObject", args)
  }

  return
}

// horizontalDirection()
func (this *QSwipeGesture) Horizontaldirection(args ...interface{}) () {
  // horizontalDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSwipeGesture19horizontalDirectionEv
    // invoke: QSwipeGesture::SwipeDirection horizontalDirection()
    C.C_ZNK13QSwipeGesture19horizontalDirectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSwipeGesture", "horizontalDirection", args)
  }

  return
}

// verticalDirection()
func (this *QSwipeGesture) Verticaldirection(args ...interface{}) () {
  // verticalDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSwipeGesture17verticalDirectionEv
    // invoke: QSwipeGesture::SwipeDirection verticalDirection()
    C.C_ZNK13QSwipeGesture17verticalDirectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSwipeGesture", "verticalDirection", args)
  }

  return
}

// swipeAngle()
func (this *QSwipeGesture) Swipeangle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QSwipeGesture10swipeAngleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSwipeGesture", "swipeAngle", args)
  }

  return
}

// ~QSwipeGesture()
func (this *QSwipeGesture) Freeqswipegesture(args ...interface{}) () {
  // ~QSwipeGesture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSwipeGestureD0Ev
    // invoke: void ~QSwipeGesture()
    C.C_ZN13QSwipeGestureD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSwipeGesture", "~QSwipeGesture", args)
  }

  return
}

// setSwipeAngle(qreal)
func (this *QSwipeGesture) Setswipeangle(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QSwipeGesture13setSwipeAngleEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSwipeGesture", "setSwipeAngle", args)
  }

  return
}

// metaObject()
func (this *QGesture) Metaobject(args ...interface{}) () {
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
    C.C_ZNK8QGesture10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGesture", "metaObject", args)
  }

  return
}

// hasHotSpot()
func (this *QGesture) Hashotspot(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QGesture10hasHotSpotEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGesture", "hasHotSpot", args)
  }

  return
}

// gestureCancelPolicy()
func (this *QGesture) Gesturecancelpolicy(args ...interface{}) () {
  // gestureCancelPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QGesture19gestureCancelPolicyEv
    // invoke: QGesture::GestureCancelPolicy gestureCancelPolicy()
    C.C_ZNK8QGesture19gestureCancelPolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGesture", "gestureCancelPolicy", args)
  }

  return
}

// state()
func (this *QGesture) State(args ...interface{}) () {
  // state()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QGesture5stateEv
    // invoke: Qt::GestureState state()
    C.C_ZNK8QGesture5stateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGesture", "state", args)
  }

  return
}

// setHotSpot(const class QPointF &)
func (this *QGesture) Sethotspot(args ...interface{}) () {
  // setHotSpot(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QGesture10setHotSpotERK7QPointF
    // invoke: void setHotSpot(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QGesture10setHotSpotERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGesture", "setHotSpot", args)
  }

  return
}

// hotSpot()
func (this *QGesture) Hotspot(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QGesture7hotSpotEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGesture", "hotSpot", args)
  }

  return
}

// ~QGesture()
func (this *QGesture) Freeqgesture(args ...interface{}) () {
  // ~QGesture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QGestureD0Ev
    // invoke: void ~QGesture()
    C.C_ZN8QGestureD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGesture", "~QGesture", args)
  }

  return
}

// gestureType()
func (this *QGesture) Gesturetype(args ...interface{}) () {
  // gestureType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QGesture11gestureTypeEv
    // invoke: Qt::GestureType gestureType()
    C.C_ZNK8QGesture11gestureTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGesture", "gestureType", args)
  }

  return
}

// unsetHotSpot()
func (this *QGesture) Unsethotspot(args ...interface{}) () {
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
    C.C_ZN8QGesture12unsetHotSpotEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGesture", "unsetHotSpot", args)
  }

  return
}

// QGesture(class QObject *)
func NewQGesture(args ...interface{}) *QGesture {
  // QGesture(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QGestureC1EP7QObject
    // invoke: void QGesture(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QGestureC2EP7QObject(arg0)
    return &QGesture{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGesture", "QGesture", args)
  }

  return nil // QGesture{Qclsinst:qthis}
}

// activeGestures()
func (this *QGestureEvent) Activegestures(args ...interface{}) () {
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
    C.C_ZNK13QGestureEvent14activeGesturesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGestureEvent", "activeGestures", args)
  }

  return
}

// gestures()
func (this *QGestureEvent) Gestures(args ...interface{}) () {
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
    C.C_ZNK13QGestureEvent8gesturesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGestureEvent", "gestures", args)
  }

  return
}

// widget()
func (this *QGestureEvent) Widget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QGestureEvent6widgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGestureEvent", "widget", args)
  }

  return
}

// canceledGestures()
func (this *QGestureEvent) Canceledgestures(args ...interface{}) () {
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
    C.C_ZNK13QGestureEvent16canceledGesturesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGestureEvent", "canceledGestures", args)
  }

  return
}

// setWidget(class QWidget *)
func (this *QGestureEvent) Setwidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGestureEvent9setWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureEvent", "setWidget", args)
  }

  return
}

// accept(class QGesture *)
func (this *QGestureEvent) Accept(args ...interface{}) () {
  // accept(class QGesture *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent6acceptEP8QGesture
    // invoke: void accept(class QGesture *)
    var arg0 = args[0].(*QGesture).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGestureEvent6acceptEP8QGesture(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureEvent", "accept", args)
  }

  return
}

// ignore(class QGesture *)
func (this *QGestureEvent) Ignore(args ...interface{}) () {
  // ignore(class QGesture *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent6ignoreEP8QGesture
    // invoke: void ignore(class QGesture *)
    var arg0 = args[0].(*QGesture).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QGestureEvent6ignoreEP8QGesture(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGestureEvent", "ignore", args)
  }

  return
}

// isAccepted(class QGesture *)
func (this *QGestureEvent) Isaccepted(args ...interface{}) (ret interface{}) {
  // isAccepted(class QGesture *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent10isAcceptedEP8QGesture
    // invoke: bool isAccepted(class QGesture *)
    var arg0 = args[0].(*QGesture).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGestureEvent10isAcceptedEP8QGesture(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGestureEvent", "isAccepted", args)
  }

  return
}

// ~QGestureEvent()
func (this *QGestureEvent) Freeqgestureevent(args ...interface{}) () {
  // ~QGestureEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEventD0Ev
    // invoke: void ~QGestureEvent()
    C.C_ZN13QGestureEventD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGestureEvent", "~QGestureEvent", args)
  }

  return
}

// setAccepted(class QGesture *, _Bool)
func (this *QGestureEvent) Setaccepted(args ...interface{}) () {
  // setAccepted(class QGesture *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QGestureEvent11setAcceptedEP8QGestureb
    // invoke: void setAccepted(class QGesture *, _Bool)
    var arg0 = args[0].(*QGesture).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN13QGestureEvent11setAcceptedEP8QGestureb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGestureEvent", "setAccepted", args)
  }

  return
}

// mapToGraphicsScene(const class QPointF &)
func (this *QGestureEvent) Maptographicsscene(args ...interface{}) (ret interface{}) {
  // mapToGraphicsScene(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QGestureEvent18mapToGraphicsSceneERK7QPointF
    // invoke: QPointF mapToGraphicsScene(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QGestureEvent18mapToGraphicsSceneERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGestureEvent", "mapToGraphicsScene", args)
  }

  return
}

// acceleration()
func (this *QPanGesture) Acceleration(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QPanGesture12accelerationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPanGesture", "acceleration", args)
  }

  return
}

// setAcceleration(qreal)
func (this *QPanGesture) Setacceleration(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QPanGesture15setAccelerationEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPanGesture", "setAcceleration", args)
  }

  return
}

// metaObject()
func (this *QPanGesture) Metaobject(args ...interface{}) () {
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
    C.C_ZNK11QPanGesture10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPanGesture", "metaObject", args)
  }

  return
}

// QPanGesture(class QObject *)
func NewQPanGesture(args ...interface{}) *QPanGesture {
  // QPanGesture(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPanGestureC1EP7QObject
    // invoke: void QPanGesture(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPanGestureC2EP7QObject(arg0)
    return &QPanGesture{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPanGesture", "QPanGesture", args)
  }

  return nil // QPanGesture{Qclsinst:qthis}
}

// offset()
func (this *QPanGesture) Offset(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QPanGesture6offsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPanGesture", "offset", args)
  }

  return
}

// ~QPanGesture()
func (this *QPanGesture) Freeqpangesture(args ...interface{}) () {
  // ~QPanGesture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPanGestureD0Ev
    // invoke: void ~QPanGesture()
    C.C_ZN11QPanGestureD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPanGesture", "~QPanGesture", args)
  }

  return
}

// delta()
func (this *QPanGesture) Delta(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QPanGesture5deltaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPanGesture", "delta", args)
  }

  return
}

// setOffset(const class QPointF &)
func (this *QPanGesture) Setoffset(args ...interface{}) () {
  // setOffset(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPanGesture9setOffsetERK7QPointF
    // invoke: void setOffset(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QPanGesture9setOffsetERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPanGesture", "setOffset", args)
  }

  return
}

// lastOffset()
func (this *QPanGesture) Lastoffset(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QPanGesture10lastOffsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPanGesture", "lastOffset", args)
  }

  return
}

// setLastOffset(const class QPointF &)
func (this *QPanGesture) Setlastoffset(args ...interface{}) () {
  // setLastOffset(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPanGesture13setLastOffsetERK7QPointF
    // invoke: void setLastOffset(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QPanGesture13setLastOffsetERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPanGesture", "setLastOffset", args)
  }

  return
}

// QTapAndHoldGesture(class QObject *)
func NewQTapAndHoldGesture(args ...interface{}) *QTapAndHoldGesture {
  // QTapAndHoldGesture(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QTapAndHoldGestureC1EP7QObject
    // invoke: void QTapAndHoldGesture(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QTapAndHoldGestureC2EP7QObject(arg0)
    return &QTapAndHoldGesture{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "QTapAndHoldGesture", args)
  }

  return nil // QTapAndHoldGesture{Qclsinst:qthis}
}

// metaObject()
func (this *QTapAndHoldGesture) Metaobject(args ...interface{}) () {
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
    C.C_ZNK18QTapAndHoldGesture10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "metaObject", args)
  }

  return
}

// ~QTapAndHoldGesture()
func (this *QTapAndHoldGesture) Freeqtapandholdgesture(args ...interface{}) () {
  // ~QTapAndHoldGesture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QTapAndHoldGestureD0Ev
    // invoke: void ~QTapAndHoldGesture()
    C.C_ZN18QTapAndHoldGestureD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "~QTapAndHoldGesture", args)
  }

  return
}

// timeout()
func (this *QTapAndHoldGesture) Timeout_S(args ...interface{}) (ret interface{}) {
  // timeout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QTapAndHoldGesture7timeoutEv
    // invoke: int timeout()
    var ret0 = C.C_ZN18QTapAndHoldGesture7timeoutEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "timeout", args)
  }

  return
}

// position()
func (this *QTapAndHoldGesture) Position(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QTapAndHoldGesture8positionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "position", args)
  }

  return
}

// setPosition(const class QPointF &)
func (this *QTapAndHoldGesture) Setposition(args ...interface{}) () {
  // setPosition(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QTapAndHoldGesture11setPositionERK7QPointF
    // invoke: void setPosition(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QTapAndHoldGesture11setPositionERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "setPosition", args)
  }

  return
}

// setTimeout(int)
func (this *QTapAndHoldGesture) Settimeout_S(args ...interface{}) () {
  // setTimeout(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QTapAndHoldGesture10setTimeoutEi
    // invoke: void setTimeout(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QTapAndHoldGesture10setTimeoutEi(arg0)
  default:
    qtrt.ErrorResolve("QTapAndHoldGesture", "setTimeout", args)
  }

  return
}

// metaObject()
func (this *QTapGesture) Metaobject(args ...interface{}) () {
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
    C.C_ZNK11QTapGesture10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTapGesture", "metaObject", args)
  }

  return
}

// QTapGesture(class QObject *)
func NewQTapGesture(args ...interface{}) *QTapGesture {
  // QTapGesture(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTapGestureC1EP7QObject
    // invoke: void QTapGesture(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTapGestureC2EP7QObject(arg0)
    return &QTapGesture{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTapGesture", "QTapGesture", args)
  }

  return nil // QTapGesture{Qclsinst:qthis}
}

// ~QTapGesture()
func (this *QTapGesture) Freeqtapgesture(args ...interface{}) () {
  // ~QTapGesture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTapGestureD0Ev
    // invoke: void ~QTapGesture()
    C.C_ZN11QTapGestureD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTapGesture", "~QTapGesture", args)
  }

  return
}

// position()
func (this *QTapGesture) Position(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTapGesture8positionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTapGesture", "position", args)
  }

  return
}

// setPosition(const class QPointF &)
func (this *QTapGesture) Setposition(args ...interface{}) () {
  // setPosition(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTapGesture11setPositionERK7QPointF
    // invoke: void setPosition(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTapGesture11setPositionERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTapGesture", "setPosition", args)
  }

  return
}

// changeFlags()
func (this *QPinchGesture) Changeflags(args ...interface{}) () {
  // changeFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture11changeFlagsEv
    // invoke: ChangeFlags changeFlags()
    C.C_ZNK13QPinchGesture11changeFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "changeFlags", args)
  }

  return
}

// startCenterPoint()
func (this *QPinchGesture) Startcenterpoint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPinchGesture16startCenterPointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPinchGesture", "startCenterPoint", args)
  }

  return
}

// setLastCenterPoint(const class QPointF &)
func (this *QPinchGesture) Setlastcenterpoint(args ...interface{}) () {
  // setLastCenterPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture18setLastCenterPointERK7QPointF
    // invoke: void setLastCenterPoint(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QPinchGesture18setLastCenterPointERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setLastCenterPoint", args)
  }

  return
}

// setStartCenterPoint(const class QPointF &)
func (this *QPinchGesture) Setstartcenterpoint(args ...interface{}) () {
  // setStartCenterPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture19setStartCenterPointERK7QPointF
    // invoke: void setStartCenterPoint(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QPinchGesture19setStartCenterPointERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setStartCenterPoint", args)
  }

  return
}

// setTotalRotationAngle(qreal)
func (this *QPinchGesture) Settotalrotationangle(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QPinchGesture21setTotalRotationAngleEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setTotalRotationAngle", args)
  }

  return
}

// lastRotationAngle()
func (this *QPinchGesture) Lastrotationangle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPinchGesture17lastRotationAngleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPinchGesture", "lastRotationAngle", args)
  }

  return
}

// rotationAngle()
func (this *QPinchGesture) Rotationangle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPinchGesture13rotationAngleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPinchGesture", "rotationAngle", args)
  }

  return
}

// totalChangeFlags()
func (this *QPinchGesture) Totalchangeflags(args ...interface{}) () {
  // totalChangeFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QPinchGesture16totalChangeFlagsEv
    // invoke: ChangeFlags totalChangeFlags()
    C.C_ZNK13QPinchGesture16totalChangeFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "totalChangeFlags", args)
  }

  return
}

// setTotalScaleFactor(qreal)
func (this *QPinchGesture) Settotalscalefactor(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QPinchGesture19setTotalScaleFactorEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setTotalScaleFactor", args)
  }

  return
}

// setLastScaleFactor(qreal)
func (this *QPinchGesture) Setlastscalefactor(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QPinchGesture18setLastScaleFactorEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setLastScaleFactor", args)
  }

  return
}

// ~QPinchGesture()
func (this *QPinchGesture) Freeqpinchgesture(args ...interface{}) () {
  // ~QPinchGesture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGestureD0Ev
    // invoke: void ~QPinchGesture()
    C.C_ZN13QPinchGestureD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "~QPinchGesture", args)
  }

  return
}

// totalRotationAngle()
func (this *QPinchGesture) Totalrotationangle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPinchGesture18totalRotationAngleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPinchGesture", "totalRotationAngle", args)
  }

  return
}

// setRotationAngle(qreal)
func (this *QPinchGesture) Setrotationangle(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QPinchGesture16setRotationAngleEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setRotationAngle", args)
  }

  return
}

// QPinchGesture(class QObject *)
func NewQPinchGesture(args ...interface{}) *QPinchGesture {
  // QPinchGesture(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGestureC1EP7QObject
    // invoke: void QPinchGesture(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QPinchGestureC2EP7QObject(arg0)
    return &QPinchGesture{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPinchGesture", "QPinchGesture", args)
  }

  return nil // QPinchGesture{Qclsinst:qthis}
}

// lastCenterPoint()
func (this *QPinchGesture) Lastcenterpoint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPinchGesture15lastCenterPointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPinchGesture", "lastCenterPoint", args)
  }

  return
}

// centerPoint()
func (this *QPinchGesture) Centerpoint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPinchGesture11centerPointEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPinchGesture", "centerPoint", args)
  }

  return
}

// totalScaleFactor()
func (this *QPinchGesture) Totalscalefactor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPinchGesture16totalScaleFactorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPinchGesture", "totalScaleFactor", args)
  }

  return
}

// lastScaleFactor()
func (this *QPinchGesture) Lastscalefactor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPinchGesture15lastScaleFactorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPinchGesture", "lastScaleFactor", args)
  }

  return
}

// setCenterPoint(const class QPointF &)
func (this *QPinchGesture) Setcenterpoint(args ...interface{}) () {
  // setCenterPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QPinchGesture14setCenterPointERK7QPointF
    // invoke: void setCenterPoint(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QPinchGesture14setCenterPointERK7QPointF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setCenterPoint", args)
  }

  return
}

// scaleFactor()
func (this *QPinchGesture) Scalefactor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QPinchGesture11scaleFactorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPinchGesture", "scaleFactor", args)
  }

  return
}

// metaObject()
func (this *QPinchGesture) Metaobject(args ...interface{}) () {
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
    C.C_ZNK13QPinchGesture10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPinchGesture", "metaObject", args)
  }

  return
}

// setLastRotationAngle(qreal)
func (this *QPinchGesture) Setlastrotationangle(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QPinchGesture20setLastRotationAngleEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setLastRotationAngle", args)
  }

  return
}

// setScaleFactor(qreal)
func (this *QPinchGesture) Setscalefactor(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QPinchGesture14setScaleFactorEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPinchGesture", "setScaleFactor", args)
  }

  return
}

// <= body block end

