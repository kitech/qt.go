package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QGraphicsItemAnimation::setPosAt(qreal step, const QPointF & pos);
extern void _ZN22QGraphicsItemAnimation8setPosAtEdRK7QPointF(void* qthis, double arg0, void* arg1);
  // proto:  void QGraphicsItemAnimation::QGraphicsItemAnimation(const QGraphicsItemAnimation & );
extern void* dector_ZN22QGraphicsItemAnimationC1ERKS_(void* arg0);
extern void _ZN22QGraphicsItemAnimationC1ERKS_(void* qthis, void* arg0);
  // proto:  qreal QGraphicsItemAnimation::xTranslationAt(qreal step);
extern void _ZNK22QGraphicsItemAnimation14xTranslationAtEd(void* qthis, double arg0);
  // proto:  void QGraphicsItemAnimation::setRotationAt(qreal step, qreal angle);
extern void _ZN22QGraphicsItemAnimation13setRotationAtEdd(void* qthis, double arg0, double arg1);
  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::posList();
extern void _ZNK22QGraphicsItemAnimation7posListEv(void* qthis);
  // proto:  qreal QGraphicsItemAnimation::verticalScaleAt(qreal step);
extern void _ZNK22QGraphicsItemAnimation15verticalScaleAtEd(void* qthis, double arg0);
  // proto:  QPointF QGraphicsItemAnimation::posAt(qreal step);
extern void _ZNK22QGraphicsItemAnimation5posAtEd(void* qthis, double arg0);
  // proto:  qreal QGraphicsItemAnimation::horizontalShearAt(qreal step);
extern void _ZNK22QGraphicsItemAnimation17horizontalShearAtEd(void* qthis, double arg0);
  // proto:  qreal QGraphicsItemAnimation::yTranslationAt(qreal step);
extern void _ZNK22QGraphicsItemAnimation14yTranslationAtEd(void* qthis, double arg0);
  // proto:  QMatrix QGraphicsItemAnimation::matrixAt(qreal step);
extern void _ZNK22QGraphicsItemAnimation8matrixAtEd(void* qthis, double arg0);
  // proto:  QGraphicsItem * QGraphicsItemAnimation::item();
extern void _ZNK22QGraphicsItemAnimation4itemEv(void* qthis);
  // proto:  void QGraphicsItemAnimation::QGraphicsItemAnimation(QObject * parent);
extern void* dector_ZN22QGraphicsItemAnimationC1EP7QObject(void* arg0);
extern void _ZN22QGraphicsItemAnimationC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QGraphicsItemAnimation::~QGraphicsItemAnimation();
extern void _ZN22QGraphicsItemAnimationD0Ev(void* qthis);
  // proto:  void QGraphicsItemAnimation::setScaleAt(qreal step, qreal sx, qreal sy);
extern void _ZN22QGraphicsItemAnimation10setScaleAtEddd(void* qthis, double arg0, double arg1, double arg2);
  // proto:  void QGraphicsItemAnimation::setTranslationAt(qreal step, qreal dx, qreal dy);
extern void _ZN22QGraphicsItemAnimation16setTranslationAtEddd(void* qthis, double arg0, double arg1, double arg2);
  // proto:  void QGraphicsItemAnimation::setShearAt(qreal step, qreal sh, qreal sv);
extern void _ZN22QGraphicsItemAnimation10setShearAtEddd(void* qthis, double arg0, double arg1, double arg2);
  // proto:  qreal QGraphicsItemAnimation::rotationAt(qreal step);
extern void _ZNK22QGraphicsItemAnimation10rotationAtEd(void* qthis, double arg0);
  // proto:  const QMetaObject * QGraphicsItemAnimation::metaObject();
extern void _ZNK22QGraphicsItemAnimation10metaObjectEv(void* qthis);
  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::scaleList();
extern void _ZNK22QGraphicsItemAnimation9scaleListEv(void* qthis);
  // proto:  QList<QPair<qreal, qreal> > QGraphicsItemAnimation::rotationList();
extern void _ZNK22QGraphicsItemAnimation12rotationListEv(void* qthis);
  // proto:  void QGraphicsItemAnimation::reset();
extern void _ZN22QGraphicsItemAnimation5resetEv(void* qthis);
  // proto:  void QGraphicsItemAnimation::setTimeLine(QTimeLine * timeLine);
extern void _ZN22QGraphicsItemAnimation11setTimeLineEP9QTimeLine(void* qthis, void* arg0);
  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::shearList();
extern void _ZNK22QGraphicsItemAnimation9shearListEv(void* qthis);
  // proto:  void QGraphicsItemAnimation::clear();
extern void _ZN22QGraphicsItemAnimation5clearEv(void* qthis);
  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::translationList();
extern void _ZNK22QGraphicsItemAnimation15translationListEv(void* qthis);
  // proto:  void QGraphicsItemAnimation::setItem(QGraphicsItem * item);
extern void _ZN22QGraphicsItemAnimation7setItemEP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsItemAnimation::setStep(qreal x);
extern void _ZN22QGraphicsItemAnimation7setStepEd(void* qthis, double arg0);
  // proto:  QTimeLine * QGraphicsItemAnimation::timeLine();
extern void _ZNK22QGraphicsItemAnimation8timeLineEv(void* qthis);
  // proto:  qreal QGraphicsItemAnimation::horizontalScaleAt(qreal step);
extern void _ZNK22QGraphicsItemAnimation17horizontalScaleAtEd(void* qthis, double arg0);
  // proto:  qreal QGraphicsItemAnimation::verticalShearAt(qreal step);
extern void _ZNK22QGraphicsItemAnimation15verticalShearAtEd(void* qthis, double arg0);
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

// class sizeof(QGraphicsItemAnimation)=1
type QGraphicsItemAnimation struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QGraphicsItemAnimation::setPosAt(qreal step, const QPointF & pos);
func (this *QGraphicsItemAnimation) setPosAt(args ...interface{}) () {
  // setPosAt(qreal, const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation8setPosAtEdRK7QPointF
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setPosAt", args)
  }

}

  // proto:  void QGraphicsItemAnimation::QGraphicsItemAnimation(const QGraphicsItemAnimation & );
func NewQGraphicsItemAnimation(args ...interface{}) QGraphicsItemAnimation {
  return QGraphicsItemAnimation{}
}

  // proto:  qreal QGraphicsItemAnimation::xTranslationAt(qreal step);
func (this *QGraphicsItemAnimation) xTranslationAt(args ...interface{}) () {
  // xTranslationAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation14xTranslationAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "xTranslationAt", args)
  }

}

  // proto:  void QGraphicsItemAnimation::setRotationAt(qreal step, qreal angle);
func (this *QGraphicsItemAnimation) setRotationAt(args ...interface{}) () {
  // setRotationAt(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation13setRotationAtEdd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setRotationAt", args)
  }

}

  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::posList();
func (this *QGraphicsItemAnimation) posList(args ...interface{}) () {
  // posList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation7posListEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "posList", args)
  }

}

  // proto:  qreal QGraphicsItemAnimation::verticalScaleAt(qreal step);
func (this *QGraphicsItemAnimation) verticalScaleAt(args ...interface{}) () {
  // verticalScaleAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation15verticalScaleAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "verticalScaleAt", args)
  }

}

  // proto:  QPointF QGraphicsItemAnimation::posAt(qreal step);
func (this *QGraphicsItemAnimation) posAt(args ...interface{}) () {
  // posAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation5posAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "posAt", args)
  }

}

  // proto:  qreal QGraphicsItemAnimation::horizontalShearAt(qreal step);
func (this *QGraphicsItemAnimation) horizontalShearAt(args ...interface{}) () {
  // horizontalShearAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation17horizontalShearAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "horizontalShearAt", args)
  }

}

  // proto:  qreal QGraphicsItemAnimation::yTranslationAt(qreal step);
func (this *QGraphicsItemAnimation) yTranslationAt(args ...interface{}) () {
  // yTranslationAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation14yTranslationAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "yTranslationAt", args)
  }

}

  // proto:  QMatrix QGraphicsItemAnimation::matrixAt(qreal step);
func (this *QGraphicsItemAnimation) matrixAt(args ...interface{}) () {
  // matrixAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation8matrixAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "matrixAt", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsItemAnimation::item();
func (this *QGraphicsItemAnimation) item(args ...interface{}) () {
  // item()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation4itemEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "item", args)
  }

}

  // proto:  void QGraphicsItemAnimation::~QGraphicsItemAnimation();
func (this *QGraphicsItemAnimation) FreeQGraphicsItemAnimation(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "~QGraphicsItemAnimation", args)
  }

}

  // proto:  void QGraphicsItemAnimation::setScaleAt(qreal step, qreal sx, qreal sy);
func (this *QGraphicsItemAnimation) setScaleAt(args ...interface{}) () {
  // setScaleAt(qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation10setScaleAtEddd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setScaleAt", args)
  }

}

  // proto:  void QGraphicsItemAnimation::setTranslationAt(qreal step, qreal dx, qreal dy);
func (this *QGraphicsItemAnimation) setTranslationAt(args ...interface{}) () {
  // setTranslationAt(qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation16setTranslationAtEddd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setTranslationAt", args)
  }

}

  // proto:  void QGraphicsItemAnimation::setShearAt(qreal step, qreal sh, qreal sv);
func (this *QGraphicsItemAnimation) setShearAt(args ...interface{}) () {
  // setShearAt(qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation10setShearAtEddd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setShearAt", args)
  }

}

  // proto:  qreal QGraphicsItemAnimation::rotationAt(qreal step);
func (this *QGraphicsItemAnimation) rotationAt(args ...interface{}) () {
  // rotationAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation10rotationAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "rotationAt", args)
  }

}

  // proto:  const QMetaObject * QGraphicsItemAnimation::metaObject();
func (this *QGraphicsItemAnimation) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation10metaObjectEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "metaObject", args)
  }

}

  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::scaleList();
func (this *QGraphicsItemAnimation) scaleList(args ...interface{}) () {
  // scaleList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation9scaleListEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "scaleList", args)
  }

}

  // proto:  QList<QPair<qreal, qreal> > QGraphicsItemAnimation::rotationList();
func (this *QGraphicsItemAnimation) rotationList(args ...interface{}) () {
  // rotationList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation12rotationListEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "rotationList", args)
  }

}

  // proto:  void QGraphicsItemAnimation::reset();
func (this *QGraphicsItemAnimation) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation5resetEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "reset", args)
  }

}

  // proto:  void QGraphicsItemAnimation::setTimeLine(QTimeLine * timeLine);
func (this *QGraphicsItemAnimation) setTimeLine(args ...interface{}) () {
  // setTimeLine(class QTimeLine *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTimeLine{}) // "QTimeLine *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation11setTimeLineEP9QTimeLine
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setTimeLine", args)
  }

}

  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::shearList();
func (this *QGraphicsItemAnimation) shearList(args ...interface{}) () {
  // shearList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation9shearListEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "shearList", args)
  }

}

  // proto:  void QGraphicsItemAnimation::clear();
func (this *QGraphicsItemAnimation) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation5clearEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "clear", args)
  }

}

  // proto:  QList<QPair<qreal, QPointF> > QGraphicsItemAnimation::translationList();
func (this *QGraphicsItemAnimation) translationList(args ...interface{}) () {
  // translationList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation15translationListEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "translationList", args)
  }

}

  // proto:  void QGraphicsItemAnimation::setItem(QGraphicsItem * item);
func (this *QGraphicsItemAnimation) setItem(args ...interface{}) () {
  // setItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation7setItemEP13QGraphicsItem
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setItem", args)
  }

}

  // proto:  void QGraphicsItemAnimation::setStep(qreal x);
func (this *QGraphicsItemAnimation) setStep(args ...interface{}) () {
  // setStep(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsItemAnimation7setStepEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "setStep", args)
  }

}

  // proto:  QTimeLine * QGraphicsItemAnimation::timeLine();
func (this *QGraphicsItemAnimation) timeLine(args ...interface{}) () {
  // timeLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation8timeLineEv
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "timeLine", args)
  }

}

  // proto:  qreal QGraphicsItemAnimation::horizontalScaleAt(qreal step);
func (this *QGraphicsItemAnimation) horizontalScaleAt(args ...interface{}) () {
  // horizontalScaleAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation17horizontalScaleAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "horizontalScaleAt", args)
  }

}

  // proto:  qreal QGraphicsItemAnimation::verticalShearAt(qreal step);
func (this *QGraphicsItemAnimation) verticalShearAt(args ...interface{}) () {
  // verticalShearAt(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsItemAnimation15verticalShearAtEd
  default:
    qtrt.ErrorResolve("QGraphicsItemAnimation", "verticalShearAt", args)
  }

}

// <= body block end

