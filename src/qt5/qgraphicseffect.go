package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qgraphicseffect.h
// dst-file: /src/widgets/qgraphicseffect.go
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
  // proto:  void QGraphicsColorizeEffect::setColor(const QColor & c);
extern void _ZN23QGraphicsColorizeEffect8setColorERK6QColor(void* qthis, void* arg0);
  // proto:  void QGraphicsColorizeEffect::QGraphicsColorizeEffect(const QGraphicsColorizeEffect & );
extern void* dector_ZN23QGraphicsColorizeEffectC1ERKS_(void* arg0);
extern void _ZN23QGraphicsColorizeEffectC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsColorizeEffect::setStrength(qreal strength);
extern void _ZN23QGraphicsColorizeEffect11setStrengthEd(void* qthis, double arg0);
  // proto:  void QGraphicsColorizeEffect::QGraphicsColorizeEffect(QObject * parent);
extern void* dector_ZN23QGraphicsColorizeEffectC1EP7QObject(void* arg0);
extern void _ZN23QGraphicsColorizeEffectC1EP7QObject(void* qthis, void* arg0);
  // proto:  qreal QGraphicsColorizeEffect::strength();
extern void _ZNK23QGraphicsColorizeEffect8strengthEv(void* qthis);
  // proto:  void QGraphicsColorizeEffect::~QGraphicsColorizeEffect();
extern void _ZN23QGraphicsColorizeEffectD0Ev(void* qthis);
  // proto:  QColor QGraphicsColorizeEffect::color();
extern void _ZNK23QGraphicsColorizeEffect5colorEv(void* qthis);
  // proto:  const QMetaObject * QGraphicsColorizeEffect::metaObject();
extern void _ZNK23QGraphicsColorizeEffect10metaObjectEv(void* qthis);
  // proto:  QRectF QGraphicsEffect::boundingRectFor(const QRectF & sourceRect);
extern void _ZNK15QGraphicsEffect15boundingRectForERK6QRectF(void* qthis, void* arg0);
  // proto:  QGraphicsEffectSource * QGraphicsEffect::source();
extern void _ZNK15QGraphicsEffect6sourceEv(void* qthis);
  // proto:  void QGraphicsEffect::update();
extern void _ZN15QGraphicsEffect6updateEv(void* qthis);
  // proto:  void QGraphicsEffect::setEnabled(bool enable);
extern void _ZN15QGraphicsEffect10setEnabledEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QGraphicsEffect::metaObject();
extern void _ZNK15QGraphicsEffect10metaObjectEv(void* qthis);
  // proto:  bool QGraphicsEffect::isEnabled();
extern void _ZNK15QGraphicsEffect9isEnabledEv(void* qthis);
  // proto:  QRectF QGraphicsEffect::boundingRect();
extern void _ZNK15QGraphicsEffect12boundingRectEv(void* qthis);
  // proto:  void QGraphicsEffect::~QGraphicsEffect();
extern void _ZN15QGraphicsEffectD0Ev(void* qthis);
  // proto:  void QGraphicsEffect::QGraphicsEffect(const QGraphicsEffect & );
extern void* dector_ZN15QGraphicsEffectC1ERKS_(void* arg0);
extern void _ZN15QGraphicsEffectC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsEffect::QGraphicsEffect(QObject * parent);
extern void* dector_ZN15QGraphicsEffectC1EP7QObject(void* arg0);
extern void _ZN15QGraphicsEffectC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QGraphicsDropShadowEffect::metaObject();
extern void _ZNK25QGraphicsDropShadowEffect10metaObjectEv(void* qthis);
  // proto:  QRectF QGraphicsDropShadowEffect::boundingRectFor(const QRectF & rect);
extern void _ZNK25QGraphicsDropShadowEffect15boundingRectForERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsDropShadowEffect::QGraphicsDropShadowEffect(QObject * parent);
extern void* dector_ZN25QGraphicsDropShadowEffectC1EP7QObject(void* arg0);
extern void _ZN25QGraphicsDropShadowEffectC1EP7QObject(void* qthis, void* arg0);
  // proto:  QPointF QGraphicsDropShadowEffect::offset();
extern void _ZNK25QGraphicsDropShadowEffect6offsetEv(void* qthis);
  // proto:  void QGraphicsDropShadowEffect::setYOffset(qreal dy);
extern void demth_ZN25QGraphicsDropShadowEffect10setYOffsetEd(void* qthis, double arg0);
  // proto:  qreal QGraphicsDropShadowEffect::xOffset();
extern void demth_ZNK25QGraphicsDropShadowEffect7xOffsetEv(void* qthis);
  // proto:  qreal QGraphicsDropShadowEffect::blurRadius();
extern void _ZNK25QGraphicsDropShadowEffect10blurRadiusEv(void* qthis);
  // proto:  QColor QGraphicsDropShadowEffect::color();
extern void _ZNK25QGraphicsDropShadowEffect5colorEv(void* qthis);
  // proto:  void QGraphicsDropShadowEffect::setColor(const QColor & color);
extern void _ZN25QGraphicsDropShadowEffect8setColorERK6QColor(void* qthis, void* arg0);
  // proto:  void QGraphicsDropShadowEffect::setOffset(const QPointF & ofs);
extern void _ZN25QGraphicsDropShadowEffect9setOffsetERK7QPointF(void* qthis, void* arg0);
  // proto:  void QGraphicsDropShadowEffect::setOffset(qreal dx, qreal dy);
extern void demth_ZN25QGraphicsDropShadowEffect9setOffsetEdd(void* qthis, double arg0, double arg1);
  // proto:  void QGraphicsDropShadowEffect::setOffset(qreal d);
extern void demth_ZN25QGraphicsDropShadowEffect9setOffsetEd(void* qthis, double arg0);
  // proto:  void QGraphicsDropShadowEffect::QGraphicsDropShadowEffect(const QGraphicsDropShadowEffect & );
extern void* dector_ZN25QGraphicsDropShadowEffectC1ERKS_(void* arg0);
extern void _ZN25QGraphicsDropShadowEffectC1ERKS_(void* qthis, void* arg0);
  // proto:  qreal QGraphicsDropShadowEffect::yOffset();
extern void demth_ZNK25QGraphicsDropShadowEffect7yOffsetEv(void* qthis);
  // proto:  void QGraphicsDropShadowEffect::setXOffset(qreal dx);
extern void demth_ZN25QGraphicsDropShadowEffect10setXOffsetEd(void* qthis, double arg0);
  // proto:  void QGraphicsDropShadowEffect::setBlurRadius(qreal blurRadius);
extern void _ZN25QGraphicsDropShadowEffect13setBlurRadiusEd(void* qthis, double arg0);
  // proto:  void QGraphicsDropShadowEffect::~QGraphicsDropShadowEffect();
extern void _ZN25QGraphicsDropShadowEffectD0Ev(void* qthis);
  // proto:  void QGraphicsOpacityEffect::QGraphicsOpacityEffect(QObject * parent);
extern void* dector_ZN22QGraphicsOpacityEffectC1EP7QObject(void* arg0);
extern void _ZN22QGraphicsOpacityEffectC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QGraphicsOpacityEffect::~QGraphicsOpacityEffect();
extern void _ZN22QGraphicsOpacityEffectD0Ev(void* qthis);
  // proto:  void QGraphicsOpacityEffect::setOpacityMask(const QBrush & mask);
extern void _ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush(void* qthis, void* arg0);
  // proto:  void QGraphicsOpacityEffect::QGraphicsOpacityEffect(const QGraphicsOpacityEffect & );
extern void* dector_ZN22QGraphicsOpacityEffectC1ERKS_(void* arg0);
extern void _ZN22QGraphicsOpacityEffectC1ERKS_(void* qthis, void* arg0);
  // proto:  QBrush QGraphicsOpacityEffect::opacityMask();
extern void _ZNK22QGraphicsOpacityEffect11opacityMaskEv(void* qthis);
  // proto:  const QMetaObject * QGraphicsOpacityEffect::metaObject();
extern void _ZNK22QGraphicsOpacityEffect10metaObjectEv(void* qthis);
  // proto:  qreal QGraphicsOpacityEffect::opacity();
extern void _ZNK22QGraphicsOpacityEffect7opacityEv(void* qthis);
  // proto:  void QGraphicsOpacityEffect::setOpacity(qreal opacity);
extern void _ZN22QGraphicsOpacityEffect10setOpacityEd(void* qthis, double arg0);
  // proto:  qreal QGraphicsBlurEffect::blurRadius();
extern void _ZNK19QGraphicsBlurEffect10blurRadiusEv(void* qthis);
  // proto:  void QGraphicsBlurEffect::setBlurRadius(qreal blurRadius);
extern void _ZN19QGraphicsBlurEffect13setBlurRadiusEd(void* qthis, double arg0);
  // proto:  void QGraphicsBlurEffect::~QGraphicsBlurEffect();
extern void _ZN19QGraphicsBlurEffectD0Ev(void* qthis);
  // proto:  void QGraphicsBlurEffect::QGraphicsBlurEffect(const QGraphicsBlurEffect & );
extern void* dector_ZN19QGraphicsBlurEffectC1ERKS_(void* arg0);
extern void _ZN19QGraphicsBlurEffectC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QGraphicsBlurEffect::metaObject();
extern void _ZNK19QGraphicsBlurEffect10metaObjectEv(void* qthis);
  // proto:  QRectF QGraphicsBlurEffect::boundingRectFor(const QRectF & rect);
extern void _ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsBlurEffect::QGraphicsBlurEffect(QObject * parent);
extern void* dector_ZN19QGraphicsBlurEffectC1EP7QObject(void* arg0);
extern void _ZN19QGraphicsBlurEffectC1EP7QObject(void* qthis, void* arg0);
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

// class sizeof(QGraphicsColorizeEffect)=1
type QGraphicsColorizeEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst unsafe.Pointer /* *C.void */;
//  _strengthChanged QGraphicsColorizeEffect_strengthChanged_signal;
//  _colorChanged QGraphicsColorizeEffect_colorChanged_signal;
}

// class sizeof(QGraphicsEffect)=1
type QGraphicsEffect struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _enabledChanged QGraphicsEffect_enabledChanged_signal;
}

// class sizeof(QGraphicsDropShadowEffect)=1
type QGraphicsDropShadowEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst unsafe.Pointer /* *C.void */;
//  _colorChanged QGraphicsDropShadowEffect_colorChanged_signal;
//  _offsetChanged QGraphicsDropShadowEffect_offsetChanged_signal;
//  _blurRadiusChanged QGraphicsDropShadowEffect_blurRadiusChanged_signal;
}

// class sizeof(QGraphicsOpacityEffect)=1
type QGraphicsOpacityEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst unsafe.Pointer /* *C.void */;
//  _opacityMaskChanged QGraphicsOpacityEffect_opacityMaskChanged_signal;
//  _opacityChanged QGraphicsOpacityEffect_opacityChanged_signal;
}

// class sizeof(QGraphicsBlurEffect)=1
type QGraphicsBlurEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst unsafe.Pointer /* *C.void */;
//  _blurHintsChanged QGraphicsBlurEffect_blurHintsChanged_signal;
//  _blurRadiusChanged QGraphicsBlurEffect_blurRadiusChanged_signal;
}

  // proto:  void QGraphicsColorizeEffect::setColor(const QColor & c);
func (this *QGraphicsColorizeEffect) setColor(args ...interface{}) () {
  // setColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsColorizeEffect8setColorERK6QColor
    // invoke: void setColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN23QGraphicsColorizeEffect8setColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "setColor", args)
  }

}

  // proto:  void QGraphicsColorizeEffect::QGraphicsColorizeEffect(const QGraphicsColorizeEffect & );
func NewQGraphicsColorizeEffect(args ...interface{}) QGraphicsColorizeEffect {
  return QGraphicsColorizeEffect{}
}

  // proto:  void QGraphicsColorizeEffect::setStrength(qreal strength);
func (this *QGraphicsColorizeEffect) setStrength(args ...interface{}) () {
  // setStrength(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsColorizeEffect11setStrengthEd
    // invoke: void setStrength(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN23QGraphicsColorizeEffect11setStrengthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "setStrength", args)
  }

}

  // proto:  qreal QGraphicsColorizeEffect::strength();
func (this *QGraphicsColorizeEffect) strength(args ...interface{}) () {
  // strength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsColorizeEffect8strengthEv
    // invoke: qreal strength()
    C._ZNK23QGraphicsColorizeEffect8strengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "strength", args)
  }

}

  // proto:  void QGraphicsColorizeEffect::~QGraphicsColorizeEffect();
func (this *QGraphicsColorizeEffect) FreeQGraphicsColorizeEffect(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "~QGraphicsColorizeEffect", args)
  }

}

  // proto:  QColor QGraphicsColorizeEffect::color();
func (this *QGraphicsColorizeEffect) color(args ...interface{}) () {
  // color()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsColorizeEffect5colorEv
    // invoke: QColor color()
    C._ZNK23QGraphicsColorizeEffect5colorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "color", args)
  }

}

  // proto:  const QMetaObject * QGraphicsColorizeEffect::metaObject();
func (this *QGraphicsColorizeEffect) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsColorizeEffect10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK23QGraphicsColorizeEffect10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "metaObject", args)
  }

}

  // proto:  QRectF QGraphicsEffect::boundingRectFor(const QRectF & sourceRect);
func (this *QGraphicsEffect) boundingRectFor(args ...interface{}) () {
  // boundingRectFor(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect15boundingRectForERK6QRectF
    // invoke: QRectF boundingRectFor(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK15QGraphicsEffect15boundingRectForERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "boundingRectFor", args)
  }

}

  // proto:  QGraphicsEffectSource * QGraphicsEffect::source();
func (this *QGraphicsEffect) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect6sourceEv
    // invoke: QGraphicsEffectSource * source()
    C._ZNK15QGraphicsEffect6sourceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "source", args)
  }

}

  // proto:  void QGraphicsEffect::update();
func (this *QGraphicsEffect) update(args ...interface{}) () {
  // update()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsEffect6updateEv
    // invoke: void update()
    C._ZN15QGraphicsEffect6updateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "update", args)
  }

}

  // proto:  void QGraphicsEffect::setEnabled(bool enable);
func (this *QGraphicsEffect) setEnabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsEffect10setEnabledEb
    // invoke: void setEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsEffect10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "setEnabled", args)
  }

}

  // proto:  const QMetaObject * QGraphicsEffect::metaObject();
func (this *QGraphicsEffect) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK15QGraphicsEffect10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "metaObject", args)
  }

}

  // proto:  bool QGraphicsEffect::isEnabled();
func (this *QGraphicsEffect) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect9isEnabledEv
    // invoke: bool isEnabled()
    C._ZNK15QGraphicsEffect9isEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "isEnabled", args)
  }

}

  // proto:  QRectF QGraphicsEffect::boundingRect();
func (this *QGraphicsEffect) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect12boundingRectEv
    // invoke: QRectF boundingRect()
    C._ZNK15QGraphicsEffect12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "boundingRect", args)
  }

}

  // proto:  void QGraphicsEffect::~QGraphicsEffect();
func (this *QGraphicsEffect) FreeQGraphicsEffect(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "~QGraphicsEffect", args)
  }

}

  // proto:  void QGraphicsEffect::QGraphicsEffect(const QGraphicsEffect & );
func NewQGraphicsEffect(args ...interface{}) QGraphicsEffect {
  return QGraphicsEffect{}
}

  // proto:  const QMetaObject * QGraphicsDropShadowEffect::metaObject();
func (this *QGraphicsDropShadowEffect) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK25QGraphicsDropShadowEffect10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "metaObject", args)
  }

}

  // proto:  QRectF QGraphicsDropShadowEffect::boundingRectFor(const QRectF & rect);
func (this *QGraphicsDropShadowEffect) boundingRectFor(args ...interface{}) () {
  // boundingRectFor(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect15boundingRectForERK6QRectF
    // invoke: QRectF boundingRectFor(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK25QGraphicsDropShadowEffect15boundingRectForERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "boundingRectFor", args)
  }

}

  // proto:  void QGraphicsDropShadowEffect::QGraphicsDropShadowEffect(QObject * parent);
func NewQGraphicsDropShadowEffect(args ...interface{}) QGraphicsDropShadowEffect {
  return QGraphicsDropShadowEffect{}
}

  // proto:  QPointF QGraphicsDropShadowEffect::offset();
func (this *QGraphicsDropShadowEffect) offset(args ...interface{}) () {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect6offsetEv
    // invoke: QPointF offset()
    C._ZNK25QGraphicsDropShadowEffect6offsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "offset", args)
  }

}

  // proto:  void QGraphicsDropShadowEffect::setYOffset(qreal dy);
func (this *QGraphicsDropShadowEffect) setYOffset(args ...interface{}) () {
  // setYOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect10setYOffsetEd
    // invoke: void setYOffset(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.demth_ZN25QGraphicsDropShadowEffect10setYOffsetEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setYOffset", args)
  }

}

  // proto:  qreal QGraphicsDropShadowEffect::xOffset();
func (this *QGraphicsDropShadowEffect) xOffset(args ...interface{}) () {
  // xOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect7xOffsetEv
    // invoke: qreal xOffset()
    C.demth_ZNK25QGraphicsDropShadowEffect7xOffsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "xOffset", args)
  }

}

  // proto:  qreal QGraphicsDropShadowEffect::blurRadius();
func (this *QGraphicsDropShadowEffect) blurRadius(args ...interface{}) () {
  // blurRadius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect10blurRadiusEv
    // invoke: qreal blurRadius()
    C._ZNK25QGraphicsDropShadowEffect10blurRadiusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "blurRadius", args)
  }

}

  // proto:  QColor QGraphicsDropShadowEffect::color();
func (this *QGraphicsDropShadowEffect) color(args ...interface{}) () {
  // color()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect5colorEv
    // invoke: QColor color()
    C._ZNK25QGraphicsDropShadowEffect5colorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "color", args)
  }

}

  // proto:  void QGraphicsDropShadowEffect::setColor(const QColor & color);
func (this *QGraphicsDropShadowEffect) setColor(args ...interface{}) () {
  // setColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect8setColorERK6QColor
    // invoke: void setColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN25QGraphicsDropShadowEffect8setColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setColor", args)
  }

}

  // proto:  void QGraphicsDropShadowEffect::setOffset(const QPointF & ofs);
func (this *QGraphicsDropShadowEffect) setOffset(args ...interface{}) () {
  // setOffset(const class QPointF &)
  // setOffset(qreal, qreal)
  // setOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect9setOffsetERK7QPointF
    // invoke: void setOffset(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN25QGraphicsDropShadowEffect9setOffsetERK7QPointF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN25QGraphicsDropShadowEffect9setOffsetEdd
    // invoke: void setOffset(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.demth_ZN25QGraphicsDropShadowEffect9setOffsetEdd(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN25QGraphicsDropShadowEffect9setOffsetEd
    // invoke: void setOffset(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.demth_ZN25QGraphicsDropShadowEffect9setOffsetEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setOffset", args)
  }

}

  // proto:  qreal QGraphicsDropShadowEffect::yOffset();
func (this *QGraphicsDropShadowEffect) yOffset(args ...interface{}) () {
  // yOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect7yOffsetEv
    // invoke: qreal yOffset()
    C.demth_ZNK25QGraphicsDropShadowEffect7yOffsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "yOffset", args)
  }

}

  // proto:  void QGraphicsDropShadowEffect::setXOffset(qreal dx);
func (this *QGraphicsDropShadowEffect) setXOffset(args ...interface{}) () {
  // setXOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect10setXOffsetEd
    // invoke: void setXOffset(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.demth_ZN25QGraphicsDropShadowEffect10setXOffsetEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setXOffset", args)
  }

}

  // proto:  void QGraphicsDropShadowEffect::setBlurRadius(qreal blurRadius);
func (this *QGraphicsDropShadowEffect) setBlurRadius(args ...interface{}) () {
  // setBlurRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect13setBlurRadiusEd
    // invoke: void setBlurRadius(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN25QGraphicsDropShadowEffect13setBlurRadiusEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setBlurRadius", args)
  }

}

  // proto:  void QGraphicsDropShadowEffect::~QGraphicsDropShadowEffect();
func (this *QGraphicsDropShadowEffect) FreeQGraphicsDropShadowEffect(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "~QGraphicsDropShadowEffect", args)
  }

}

  // proto:  void QGraphicsOpacityEffect::QGraphicsOpacityEffect(QObject * parent);
func NewQGraphicsOpacityEffect(args ...interface{}) QGraphicsOpacityEffect {
  return QGraphicsOpacityEffect{}
}

  // proto:  void QGraphicsOpacityEffect::~QGraphicsOpacityEffect();
func (this *QGraphicsOpacityEffect) FreeQGraphicsOpacityEffect(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "~QGraphicsOpacityEffect", args)
  }

}

  // proto:  void QGraphicsOpacityEffect::setOpacityMask(const QBrush & mask);
func (this *QGraphicsOpacityEffect) setOpacityMask(args ...interface{}) () {
  // setOpacityMask(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush
    // invoke: void setOpacityMask(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "setOpacityMask", args)
  }

}

  // proto:  QBrush QGraphicsOpacityEffect::opacityMask();
func (this *QGraphicsOpacityEffect) opacityMask(args ...interface{}) () {
  // opacityMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsOpacityEffect11opacityMaskEv
    // invoke: QBrush opacityMask()
    C._ZNK22QGraphicsOpacityEffect11opacityMaskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "opacityMask", args)
  }

}

  // proto:  const QMetaObject * QGraphicsOpacityEffect::metaObject();
func (this *QGraphicsOpacityEffect) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsOpacityEffect10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK22QGraphicsOpacityEffect10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "metaObject", args)
  }

}

  // proto:  qreal QGraphicsOpacityEffect::opacity();
func (this *QGraphicsOpacityEffect) opacity(args ...interface{}) () {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsOpacityEffect7opacityEv
    // invoke: qreal opacity()
    C._ZNK22QGraphicsOpacityEffect7opacityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "opacity", args)
  }

}

  // proto:  void QGraphicsOpacityEffect::setOpacity(qreal opacity);
func (this *QGraphicsOpacityEffect) setOpacity(args ...interface{}) () {
  // setOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsOpacityEffect10setOpacityEd
    // invoke: void setOpacity(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN22QGraphicsOpacityEffect10setOpacityEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "setOpacity", args)
  }

}

  // proto:  qreal QGraphicsBlurEffect::blurRadius();
func (this *QGraphicsBlurEffect) blurRadius(args ...interface{}) () {
  // blurRadius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsBlurEffect10blurRadiusEv
    // invoke: qreal blurRadius()
    C._ZNK19QGraphicsBlurEffect10blurRadiusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "blurRadius", args)
  }

}

  // proto:  void QGraphicsBlurEffect::setBlurRadius(qreal blurRadius);
func (this *QGraphicsBlurEffect) setBlurRadius(args ...interface{}) () {
  // setBlurRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsBlurEffect13setBlurRadiusEd
    // invoke: void setBlurRadius(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN19QGraphicsBlurEffect13setBlurRadiusEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "setBlurRadius", args)
  }

}

  // proto:  void QGraphicsBlurEffect::~QGraphicsBlurEffect();
func (this *QGraphicsBlurEffect) FreeQGraphicsBlurEffect(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "~QGraphicsBlurEffect", args)
  }

}

  // proto:  void QGraphicsBlurEffect::QGraphicsBlurEffect(const QGraphicsBlurEffect & );
func NewQGraphicsBlurEffect(args ...interface{}) QGraphicsBlurEffect {
  return QGraphicsBlurEffect{}
}

  // proto:  const QMetaObject * QGraphicsBlurEffect::metaObject();
func (this *QGraphicsBlurEffect) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsBlurEffect10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK19QGraphicsBlurEffect10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "metaObject", args)
  }

}

  // proto:  QRectF QGraphicsBlurEffect::boundingRectFor(const QRectF & rect);
func (this *QGraphicsBlurEffect) boundingRectFor(args ...interface{}) () {
  // boundingRectFor(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF
    // invoke: QRectF boundingRectFor(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "boundingRectFor", args)
  }

}

// <= body block end

