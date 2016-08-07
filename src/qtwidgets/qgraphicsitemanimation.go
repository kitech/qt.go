package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtWidgets/qgraphicsitemanimation.h
// dst-file: /src/widgets/qgraphicsitemanimation.go
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
  // proto:  qreal QGraphicsItemAnimation::verticalScaleAt(qreal step);
extern double C_ZNK22QGraphicsItemAnimation15verticalScaleAtEd(void* qthis, double arg0); // 4
  // proto:  qreal QGraphicsItemAnimation::horizontalScaleAt(qreal step);
extern double C_ZNK22QGraphicsItemAnimation17horizontalScaleAtEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItemAnimation::setShearAt(qreal step, qreal sh, qreal sv);
extern void C_ZN22QGraphicsItemAnimation10setShearAtEddd(void* qthis, double arg0, double arg1, double arg2); // 4
  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::posList();
extern void C_ZNK22QGraphicsItemAnimation7posListEv(void* qthis); // 4
  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::translationList();
extern void C_ZNK22QGraphicsItemAnimation15translationListEv(void* qthis); // 4
  // proto:  void QGraphicsItemAnimation::QGraphicsItemAnimation(QObject * parent);
extern void* C_ZN22QGraphicsItemAnimationC2EP7QObject(void* arg0); // 3
  // proto:  qreal QGraphicsItemAnimation::rotationAt(qreal step);
extern double C_ZNK22QGraphicsItemAnimation10rotationAtEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItemAnimation::setPosAt(qreal step, const QPointF & pos);
extern void C_ZN22QGraphicsItemAnimation8setPosAtEdRK7QPointF(void* qthis, double arg0, void* arg1); // 4
  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::scaleList();
extern void C_ZNK22QGraphicsItemAnimation9scaleListEv(void* qthis); // 4
  // proto:  QTimeLine * QGraphicsItemAnimation::timeLine();
extern void* C_ZNK22QGraphicsItemAnimation8timeLineEv(void* qthis); // 4
  // proto:  qreal QGraphicsItemAnimation::yTranslationAt(qreal step);
extern double C_ZNK22QGraphicsItemAnimation14yTranslationAtEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItemAnimation::setStep(qreal x);
extern void C_ZN22QGraphicsItemAnimation7setStepEd(void* qthis, double arg0); // 4
  // proto:  qreal QGraphicsItemAnimation::xTranslationAt(qreal step);
extern double C_ZNK22QGraphicsItemAnimation14xTranslationAtEd(void* qthis, double arg0); // 4
  // proto:  void QGraphicsItemAnimation::setRotationAt(qreal step, qreal angle);
extern void C_ZN22QGraphicsItemAnimation13setRotationAtEdd(void* qthis, double arg0, double arg1); // 4
  // proto:  void QGraphicsItemAnimation::setTimeLine(QTimeLine * timeLine);
extern void C_ZN22QGraphicsItemAnimation11setTimeLineEP9QTimeLine(void* qthis, void* arg0); // 4
  // proto:  QMatrix QGraphicsItemAnimation::matrixAt(qreal step);
extern void* C_ZNK22QGraphicsItemAnimation8matrixAtEd(void* qthis, double arg0); // 4
  // proto:  qreal QGraphicsItemAnimation::horizontalShearAt(qreal step);
extern double C_ZNK22QGraphicsItemAnimation17horizontalShearAtEd(void* qthis, double arg0); // 4
  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::shearList();
extern void C_ZNK22QGraphicsItemAnimation9shearListEv(void* qthis); // 4
  // proto:  void QGraphicsItemAnimation::setScaleAt(qreal step, qreal sx, qreal sy);
extern void C_ZN22QGraphicsItemAnimation10setScaleAtEddd(void* qthis, double arg0, double arg1, double arg2); // 4
  // proto:  QList<QPair<qreal, qreal> > QGraphicsItemAnimation::rotationList();
extern void C_ZNK22QGraphicsItemAnimation12rotationListEv(void* qthis); // 4
  // proto:  void QGraphicsItemAnimation::setTranslationAt(qreal step, qreal dx, qreal dy);
extern void C_ZN22QGraphicsItemAnimation16setTranslationAtEddd(void* qthis, double arg0, double arg1, double arg2); // 4
  // proto:  void QGraphicsItemAnimation::reset();
extern void C_ZN22QGraphicsItemAnimation5resetEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsItemAnimation::metaObject();
extern void C_ZNK22QGraphicsItemAnimation10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsItemAnimation::clear();
extern void C_ZN22QGraphicsItemAnimation5clearEv(void* qthis); // 4
  // proto:  void QGraphicsItemAnimation::setItem(QGraphicsItem * item);
extern void C_ZN22QGraphicsItemAnimation7setItemEP13QGraphicsItem(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsItemAnimation::~QGraphicsItemAnimation();
extern void C_ZN22QGraphicsItemAnimationD2Ev(void* qthis); // 4
  // proto:  QPointF QGraphicsItemAnimation::posAt(qreal step);
extern void* C_ZNK22QGraphicsItemAnimation5posAtEd(void* qthis, double arg0); // 4
  // proto:  QGraphicsItem * QGraphicsItemAnimation::item();
extern void C_ZNK22QGraphicsItemAnimation4itemEv(void* qthis); // 4
  // proto:  qreal QGraphicsItemAnimation::verticalShearAt(qreal step);
extern double C_ZNK22QGraphicsItemAnimation15verticalShearAtEd(void* qthis, double arg0); // 4
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

// class sizeof(QGraphicsItemAnimation)=1
type QGraphicsItemAnimation struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// verticalScaleAt(qreal)
func (this *QGraphicsItemAnimation) Verticalscaleat(args ...interface{}) (ret interface{}) {
  // verticalScaleAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation15verticalScaleAtEd
    // invoke: qreal verticalScaleAt(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QGraphicsItemAnimation15verticalScaleAtEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "verticalScaleAt", args)
  }

  return
}

// horizontalScaleAt(qreal)
func (this *QGraphicsItemAnimation) Horizontalscaleat(args ...interface{}) (ret interface{}) {
  // horizontalScaleAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation17horizontalScaleAtEd
    // invoke: qreal horizontalScaleAt(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QGraphicsItemAnimation17horizontalScaleAtEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "horizontalScaleAt", args)
  }

  return
}

// setShearAt(qreal, qreal, qreal)
func (this *QGraphicsItemAnimation) Setshearat(args ...interface{}) () {
  // setShearAt(qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation10setShearAtEddd
    // invoke: void setShearAt(qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    C.C_ZN22QGraphicsItemAnimation10setShearAtEddd(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setShearAt", args)
  }

  return
}

// posList()
func (this *QGraphicsItemAnimation) Poslist(args ...interface{}) () {
  // posList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation7posListEv
    // invoke: QList<QPair<qreal, QPointF> > posList()
    C.C_ZNK22QGraphicsItemAnimation7posListEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "posList", args)
  }

  return
}

// translationList()
func (this *QGraphicsItemAnimation) Translationlist(args ...interface{}) () {
  // translationList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation15translationListEv
    // invoke: QList<QPair<qreal, QPointF> > translationList()
    C.C_ZNK22QGraphicsItemAnimation15translationListEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "translationList", args)
  }

  return
}

// QGraphicsItemAnimation(class QObject *)
func NewQGraphicsItemAnimation(args ...interface{}) *QGraphicsItemAnimation {
  // QGraphicsItemAnimation(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimationC1EP7QObject
    // invoke: void QGraphicsItemAnimation(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QGraphicsItemAnimationC2EP7QObject(arg0)
    return &QGraphicsItemAnimation{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "QGraphicsItemAnimation", args)
  }

  return nil // QGraphicsItemAnimation{Qclsinst:qthis}
}

// rotationAt(qreal)
func (this *QGraphicsItemAnimation) Rotationat(args ...interface{}) (ret interface{}) {
  // rotationAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation10rotationAtEd
    // invoke: qreal rotationAt(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QGraphicsItemAnimation10rotationAtEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "rotationAt", args)
  }

  return
}

// setPosAt(qreal, const class QPointF &)
func (this *QGraphicsItemAnimation) Setposat(args ...interface{}) () {
  // setPosAt(qreal, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation8setPosAtEdRK7QPointF
    // invoke: void setPosAt(qreal, const class QPointF &)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN22QGraphicsItemAnimation8setPosAtEdRK7QPointF(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setPosAt", args)
  }

  return
}

// scaleList()
func (this *QGraphicsItemAnimation) Scalelist(args ...interface{}) () {
  // scaleList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation9scaleListEv
    // invoke: QList<QPair<qreal, QPointF> > scaleList()
    C.C_ZNK22QGraphicsItemAnimation9scaleListEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "scaleList", args)
  }

  return
}

// timeLine()
func (this *QGraphicsItemAnimation) Timeline(args ...interface{}) (ret interface{}) {
  // timeLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation8timeLineEv
    // invoke: QTimeLine * timeLine()
    var ret0 = C.C_ZNK22QGraphicsItemAnimation8timeLineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QTimeLine{}) // "QTimeLine *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "timeLine", args)
  }

  return
}

// yTranslationAt(qreal)
func (this *QGraphicsItemAnimation) Ytranslationat(args ...interface{}) (ret interface{}) {
  // yTranslationAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation14yTranslationAtEd
    // invoke: qreal yTranslationAt(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QGraphicsItemAnimation14yTranslationAtEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "yTranslationAt", args)
  }

  return
}

// setStep(qreal)
func (this *QGraphicsItemAnimation) Setstep(args ...interface{}) () {
  // setStep(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation7setStepEd
    // invoke: void setStep(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN22QGraphicsItemAnimation7setStepEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setStep", args)
  }

  return
}

// xTranslationAt(qreal)
func (this *QGraphicsItemAnimation) Xtranslationat(args ...interface{}) (ret interface{}) {
  // xTranslationAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation14xTranslationAtEd
    // invoke: qreal xTranslationAt(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QGraphicsItemAnimation14xTranslationAtEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "xTranslationAt", args)
  }

  return
}

// setRotationAt(qreal, qreal)
func (this *QGraphicsItemAnimation) Setrotationat(args ...interface{}) () {
  // setRotationAt(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation13setRotationAtEdd
    // invoke: void setRotationAt(qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN22QGraphicsItemAnimation13setRotationAtEdd(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setRotationAt", args)
  }

  return
}

// setTimeLine(class QTimeLine *)
func (this *QGraphicsItemAnimation) Settimeline(args ...interface{}) () {
  // setTimeLine(class QTimeLine *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QTimeLine{}) // "QTimeLine *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation11setTimeLineEP9QTimeLine
    // invoke: void setTimeLine(class QTimeLine *)
    var arg0 = args[0].(*qtcore.QTimeLine).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN22QGraphicsItemAnimation11setTimeLineEP9QTimeLine(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setTimeLine", args)
  }

  return
}

// matrixAt(qreal)
func (this *QGraphicsItemAnimation) Matrixat(args ...interface{}) (ret interface{}) {
  // matrixAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation8matrixAtEd
    // invoke: QMatrix matrixAt(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QGraphicsItemAnimation8matrixAtEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QMatrix{}) // "QMatrix"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "matrixAt", args)
  }

  return
}

// horizontalShearAt(qreal)
func (this *QGraphicsItemAnimation) Horizontalshearat(args ...interface{}) (ret interface{}) {
  // horizontalShearAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation17horizontalShearAtEd
    // invoke: qreal horizontalShearAt(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QGraphicsItemAnimation17horizontalShearAtEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "horizontalShearAt", args)
  }

  return
}

// shearList()
func (this *QGraphicsItemAnimation) Shearlist(args ...interface{}) () {
  // shearList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation9shearListEv
    // invoke: QList<QPair<qreal, QPointF> > shearList()
    C.C_ZNK22QGraphicsItemAnimation9shearListEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "shearList", args)
  }

  return
}

// setScaleAt(qreal, qreal, qreal)
func (this *QGraphicsItemAnimation) Setscaleat(args ...interface{}) () {
  // setScaleAt(qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation10setScaleAtEddd
    // invoke: void setScaleAt(qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    C.C_ZN22QGraphicsItemAnimation10setScaleAtEddd(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setScaleAt", args)
  }

  return
}

// rotationList()
func (this *QGraphicsItemAnimation) Rotationlist(args ...interface{}) () {
  // rotationList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation12rotationListEv
    // invoke: QList<QPair<qreal, qreal> > rotationList()
    C.C_ZNK22QGraphicsItemAnimation12rotationListEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "rotationList", args)
  }

  return
}

// setTranslationAt(qreal, qreal, qreal)
func (this *QGraphicsItemAnimation) Settranslationat(args ...interface{}) () {
  // setTranslationAt(qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation16setTranslationAtEddd
    // invoke: void setTranslationAt(qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    C.C_ZN22QGraphicsItemAnimation16setTranslationAtEddd(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setTranslationAt", args)
  }

  return
}

// reset()
func (this *QGraphicsItemAnimation) Reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation5resetEv
    // invoke: void reset()
    C.C_ZN22QGraphicsItemAnimation5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "reset", args)
  }

  return
}

// metaObject()
func (this *QGraphicsItemAnimation) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK22QGraphicsItemAnimation10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "metaObject", args)
  }

  return
}

// clear()
func (this *QGraphicsItemAnimation) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation5clearEv
    // invoke: void clear()
    C.C_ZN22QGraphicsItemAnimation5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "clear", args)
  }

  return
}

// setItem(class QGraphicsItem *)
func (this *QGraphicsItemAnimation) Setitem(args ...interface{}) () {
  // setItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation7setItemEP13QGraphicsItem
    // invoke: void setItem(class QGraphicsItem *)
    var arg0 = args[0].(*QGraphicsItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN22QGraphicsItemAnimation7setItemEP13QGraphicsItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setItem", args)
  }

  return
}

// ~QGraphicsItemAnimation()
func (this *QGraphicsItemAnimation) Freeqgraphicsitemanimation(args ...interface{}) () {
  // ~QGraphicsItemAnimation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimationD0Ev
    // invoke: void ~QGraphicsItemAnimation()
    C.C_ZN22QGraphicsItemAnimationD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "~QGraphicsItemAnimation", args)
  }

  return
}

// posAt(qreal)
func (this *QGraphicsItemAnimation) Posat(args ...interface{}) (ret interface{}) {
  // posAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation5posAtEd
    // invoke: QPointF posAt(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QGraphicsItemAnimation5posAtEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPointF{}) // "QPointF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "posAt", args)
  }

  return
}

// item()
func (this *QGraphicsItemAnimation) Item(args ...interface{}) () {
  // item()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation4itemEv
    // invoke: QGraphicsItem * item()
    C.C_ZNK22QGraphicsItemAnimation4itemEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "item", args)
  }

  return
}

// verticalShearAt(qreal)
func (this *QGraphicsItemAnimation) Verticalshearat(args ...interface{}) (ret interface{}) {
  // verticalShearAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation15verticalShearAtEd
    // invoke: qreal verticalShearAt(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK22QGraphicsItemAnimation15verticalShearAtEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "verticalShearAt", args)
  }

  return
}

// <= body block end

