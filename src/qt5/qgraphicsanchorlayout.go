package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qgraphicsanchorlayout.h
// dst-file: /src/widgets/qgraphicsanchorlayout.go
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
  // proto:  void QGraphicsAnchorLayout::QGraphicsAnchorLayout(QGraphicsLayoutItem * parent);
extern void* dector_ZN21QGraphicsAnchorLayoutC1EP19QGraphicsLayoutItem(void* arg0);
extern void _ZN21QGraphicsAnchorLayoutC1EP19QGraphicsLayoutItem(void* qthis, void* arg0);
  // proto:  void QGraphicsAnchorLayout::QGraphicsAnchorLayout(const QGraphicsAnchorLayout & );
extern void* dector_ZN21QGraphicsAnchorLayoutC1ERKS_(void* arg0);
extern void _ZN21QGraphicsAnchorLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  qreal QGraphicsAnchorLayout::verticalSpacing();
extern void _ZNK21QGraphicsAnchorLayout15verticalSpacingEv(void* qthis);
  // proto:  void QGraphicsAnchorLayout::setSpacing(qreal spacing);
extern void _ZN21QGraphicsAnchorLayout10setSpacingEd(void* qthis, double arg0);
  // proto:  int QGraphicsAnchorLayout::count();
extern void _ZNK21QGraphicsAnchorLayout5countEv(void* qthis);
  // proto:  qreal QGraphicsAnchorLayout::horizontalSpacing();
extern void _ZNK21QGraphicsAnchorLayout17horizontalSpacingEv(void* qthis);
  // proto:  void QGraphicsAnchorLayout::invalidate();
extern void _ZN21QGraphicsAnchorLayout10invalidateEv(void* qthis);
  // proto:  QGraphicsLayoutItem * QGraphicsAnchorLayout::itemAt(int index);
extern void _ZNK21QGraphicsAnchorLayout6itemAtEi(void* qthis, int arg0);
  // proto:  void QGraphicsAnchorLayout::setVerticalSpacing(qreal spacing);
extern void _ZN21QGraphicsAnchorLayout18setVerticalSpacingEd(void* qthis, double arg0);
  // proto:  void QGraphicsAnchorLayout::setGeometry(const QRectF & rect);
extern void _ZN21QGraphicsAnchorLayout11setGeometryERK6QRectF(void* qthis, void* arg0);
  // proto:  void QGraphicsAnchorLayout::setHorizontalSpacing(qreal spacing);
extern void _ZN21QGraphicsAnchorLayout20setHorizontalSpacingEd(void* qthis, double arg0);
  // proto:  void QGraphicsAnchorLayout::~QGraphicsAnchorLayout();
extern void _ZN21QGraphicsAnchorLayoutD0Ev(void* qthis);
  // proto:  void QGraphicsAnchorLayout::removeAt(int index);
extern void _ZN21QGraphicsAnchorLayout8removeAtEi(void* qthis, int arg0);
  // proto:  void QGraphicsAnchor::~QGraphicsAnchor();
extern void _ZN15QGraphicsAnchorD0Ev(void* qthis);
  // proto:  void QGraphicsAnchor::unsetSpacing();
extern void _ZN15QGraphicsAnchor12unsetSpacingEv(void* qthis);
  // proto:  void QGraphicsAnchor::setSpacing(qreal spacing);
extern void _ZN15QGraphicsAnchor10setSpacingEd(void* qthis, double arg0);
  // proto:  void QGraphicsAnchor::QGraphicsAnchor(QGraphicsAnchorLayout * parent);
extern void* dector_ZN15QGraphicsAnchorC1EP21QGraphicsAnchorLayout(void* arg0);
extern void _ZN15QGraphicsAnchorC1EP21QGraphicsAnchorLayout(void* qthis, void* arg0);
  // proto:  const QMetaObject * QGraphicsAnchor::metaObject();
extern void _ZNK15QGraphicsAnchor10metaObjectEv(void* qthis);
  // proto:  qreal QGraphicsAnchor::spacing();
extern void _ZNK15QGraphicsAnchor7spacingEv(void* qthis);
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

// class sizeof(QGraphicsAnchorLayout)=1
type QGraphicsAnchorLayout struct {
  /*qbase*/ QGraphicsLayout;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QGraphicsAnchor)=1
type QGraphicsAnchor struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QGraphicsAnchorLayout::QGraphicsAnchorLayout(QGraphicsLayoutItem * parent);
func NewQGraphicsAnchorLayout(args ...interface{}) QGraphicsAnchorLayout {
  return QGraphicsAnchorLayout{}
}

  // proto:  qreal QGraphicsAnchorLayout::verticalSpacing();
func (this *QGraphicsAnchorLayout) verticalSpacing(args ...interface{}) () {
  // verticalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsAnchorLayout15verticalSpacingEv
    // invoke: qreal verticalSpacing()
    C._ZNK21QGraphicsAnchorLayout15verticalSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "verticalSpacing", args)
  }

}

  // proto:  void QGraphicsAnchorLayout::setSpacing(qreal spacing);
func (this *QGraphicsAnchorLayout) setSpacing(args ...interface{}) () {
  // setSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout10setSpacingEd
    // invoke: void setSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN21QGraphicsAnchorLayout10setSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setSpacing", args)
  }

}

  // proto:  int QGraphicsAnchorLayout::count();
func (this *QGraphicsAnchorLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsAnchorLayout5countEv
    // invoke: int count()
    C._ZNK21QGraphicsAnchorLayout5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "count", args)
  }

}

  // proto:  qreal QGraphicsAnchorLayout::horizontalSpacing();
func (this *QGraphicsAnchorLayout) horizontalSpacing(args ...interface{}) () {
  // horizontalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsAnchorLayout17horizontalSpacingEv
    // invoke: qreal horizontalSpacing()
    C._ZNK21QGraphicsAnchorLayout17horizontalSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "horizontalSpacing", args)
  }

}

  // proto:  void QGraphicsAnchorLayout::invalidate();
func (this *QGraphicsAnchorLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout10invalidateEv
    // invoke: void invalidate()
    C._ZN21QGraphicsAnchorLayout10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "invalidate", args)
  }

}

  // proto:  QGraphicsLayoutItem * QGraphicsAnchorLayout::itemAt(int index);
func (this *QGraphicsAnchorLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QGraphicsAnchorLayout6itemAtEi
    // invoke: QGraphicsLayoutItem * itemAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK21QGraphicsAnchorLayout6itemAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "itemAt", args)
  }

}

  // proto:  void QGraphicsAnchorLayout::setVerticalSpacing(qreal spacing);
func (this *QGraphicsAnchorLayout) setVerticalSpacing(args ...interface{}) () {
  // setVerticalSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout18setVerticalSpacingEd
    // invoke: void setVerticalSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN21QGraphicsAnchorLayout18setVerticalSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setVerticalSpacing", args)
  }

}

  // proto:  void QGraphicsAnchorLayout::setGeometry(const QRectF & rect);
func (this *QGraphicsAnchorLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout11setGeometryERK6QRectF
    // invoke: void setGeometry(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN21QGraphicsAnchorLayout11setGeometryERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setGeometry", args)
  }

}

  // proto:  void QGraphicsAnchorLayout::setHorizontalSpacing(qreal spacing);
func (this *QGraphicsAnchorLayout) setHorizontalSpacing(args ...interface{}) () {
  // setHorizontalSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout20setHorizontalSpacingEd
    // invoke: void setHorizontalSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN21QGraphicsAnchorLayout20setHorizontalSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setHorizontalSpacing", args)
  }

}

  // proto:  void QGraphicsAnchorLayout::~QGraphicsAnchorLayout();
func (this *QGraphicsAnchorLayout) FreeQGraphicsAnchorLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "~QGraphicsAnchorLayout", args)
  }

}

  // proto:  void QGraphicsAnchorLayout::removeAt(int index);
func (this *QGraphicsAnchorLayout) removeAt(args ...interface{}) () {
  // removeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGraphicsAnchorLayout8removeAtEi
    // invoke: void removeAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN21QGraphicsAnchorLayout8removeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "removeAt", args)
  }

}

  // proto:  void QGraphicsAnchor::~QGraphicsAnchor();
func (this *QGraphicsAnchor) FreeQGraphicsAnchor(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "~QGraphicsAnchor", args)
  }

}

  // proto:  void QGraphicsAnchor::unsetSpacing();
func (this *QGraphicsAnchor) unsetSpacing(args ...interface{}) () {
  // unsetSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsAnchor12unsetSpacingEv
    // invoke: void unsetSpacing()
    C._ZN15QGraphicsAnchor12unsetSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "unsetSpacing", args)
  }

}

  // proto:  void QGraphicsAnchor::setSpacing(qreal spacing);
func (this *QGraphicsAnchor) setSpacing(args ...interface{}) () {
  // setSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsAnchor10setSpacingEd
    // invoke: void setSpacing(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsAnchor10setSpacingEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "setSpacing", args)
  }

}

  // proto:  void QGraphicsAnchor::QGraphicsAnchor(QGraphicsAnchorLayout * parent);
func NewQGraphicsAnchor(args ...interface{}) QGraphicsAnchor {
  return QGraphicsAnchor{}
}

  // proto:  const QMetaObject * QGraphicsAnchor::metaObject();
func (this *QGraphicsAnchor) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsAnchor10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK15QGraphicsAnchor10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "metaObject", args)
  }

}

  // proto:  qreal QGraphicsAnchor::spacing();
func (this *QGraphicsAnchor) spacing(args ...interface{}) () {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsAnchor7spacingEv
    // invoke: qreal spacing()
    C._ZNK15QGraphicsAnchor7spacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "spacing", args)
  }

}

// <= body block end

