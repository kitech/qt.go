package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qlayout.h
// dst-file: /src/widgets/qlayout.go
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
  // proto:  void QLayout::invalidate();
extern void C_ZN7QLayout10invalidateEv(void* qthis); // 4
  // proto:  int QLayout::totalHeightForWidth(int w);
extern int32_t C_ZNK7QLayout19totalHeightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QLayout::totalMinimumSize();
extern void* C_ZNK7QLayout16totalMinimumSizeEv(void* qthis); // 4
  // proto:  void QLayout::~QLayout();
extern void C_ZN7QLayoutD2Ev(void* qthis); // 4
  // proto:  QMargins QLayout::contentsMargins();
extern void* C_ZNK7QLayout15contentsMarginsEv(void* qthis); // 4
  // proto:  QRect QLayout::contentsRect();
extern void* C_ZNK7QLayout12contentsRectEv(void* qthis); // 4
  // proto:  void QLayout::setMenuBar(QWidget * w);
extern void C_ZN7QLayout10setMenuBarEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QLayout::SizeConstraint QLayout::sizeConstraint();
extern void C_ZNK7QLayout14sizeConstraintEv(void* qthis); // 4
  // proto:  Qt::Orientations QLayout::expandingDirections();
extern void C_ZNK7QLayout19expandingDirectionsEv(void* qthis); // 4
  // proto:  QLayout * QLayout::layout();
extern void* C_ZN7QLayout6layoutEv(void* qthis); // 4
  // proto:  void QLayout::removeWidget(QWidget * w);
extern void C_ZN7QLayout12removeWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  bool QLayout::isEmpty();
extern bool C_ZNK7QLayout7isEmptyEv(void* qthis); // 4
  // proto:  void QLayout::setContentsMargins(int left, int top, int right, int bottom);
extern void C_ZN7QLayout18setContentsMarginsEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QLayout::setContentsMargins(const QMargins & margins);
extern void C_ZN7QLayout18setContentsMarginsERK8QMargins(void* qthis, void* arg0); // 4
  // proto:  void QLayout::setEnabled(bool );
extern void C_ZN7QLayout10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QSizePolicy::ControlTypes QLayout::controlTypes();
extern void C_ZNK7QLayout12controlTypesEv(void* qthis); // 4
  // proto:  void QLayout::setGeometry(const QRect & );
extern void C_ZN7QLayout11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  int QLayout::indexOf(QWidget * );
extern int32_t C_ZNK7QLayout7indexOfEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QLayout::setSpacing(int );
extern void C_ZN7QLayout10setSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  int QLayout::spacing();
extern int32_t C_ZNK7QLayout7spacingEv(void* qthis); // 4
  // proto:  void QLayout::update();
extern void C_ZN7QLayout6updateEv(void* qthis); // 4
  // proto:  void QLayout::setMargin(int );
extern void C_ZN7QLayout9setMarginEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLayout::QLayout(QWidget * parent);
extern void* C_ZN7QLayoutC2EP7QWidget(void* arg0); // 3
  // proto:  void QLayout::QLayout();
extern void* C_ZN7QLayoutC2Ev(); // 3
  // proto:  void QLayout::removeItem(QLayoutItem * );
extern void C_ZN7QLayout10removeItemEP11QLayoutItem(void* qthis, void* arg0); // 4
  // proto:  void QLayout::addWidget(QWidget * w);
extern void C_ZN7QLayout9addWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  bool QLayout::activate();
extern bool C_ZN7QLayout8activateEv(void* qthis); // 4
  // proto:  QWidget * QLayout::parentWidget();
extern void* C_ZNK7QLayout12parentWidgetEv(void* qthis); // 4
  // proto:  QSize QLayout::totalSizeHint();
extern void* C_ZNK7QLayout13totalSizeHintEv(void* qthis); // 4
  // proto:  const QMetaObject * QLayout::metaObject();
extern void C_ZNK7QLayout10metaObjectEv(void* qthis); // 4
  // proto:  QSize QLayout::totalMaximumSize();
extern void* C_ZNK7QLayout16totalMaximumSizeEv(void* qthis); // 4
  // proto:  QRect QLayout::geometry();
extern void* C_ZNK7QLayout8geometryEv(void* qthis); // 4
  // proto:  void QLayout::getContentsMargins(int * left, int * top, int * right, int * bottom);
extern void C_ZNK7QLayout18getContentsMarginsEPiS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  QSize QLayout::minimumSize();
extern void* C_ZNK7QLayout11minimumSizeEv(void* qthis); // 4
  // proto: static QSize QLayout::closestAcceptableSize(const QWidget * w, const QSize & s);
extern void* C_ZN7QLayout21closestAcceptableSizeEPK7QWidgetRK5QSize(void* arg0, void* arg1); // 4
  // proto:  int QLayout::margin();
extern int32_t C_ZNK7QLayout6marginEv(void* qthis); // 4
  // proto:  bool QLayout::isEnabled();
extern bool C_ZNK7QLayout9isEnabledEv(void* qthis); // 4
  // proto:  QSize QLayout::maximumSize();
extern void* C_ZNK7QLayout11maximumSizeEv(void* qthis); // 4
  // proto:  QWidget * QLayout::menuBar();
extern void* C_ZNK7QLayout7menuBarEv(void* qthis); // 4
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

// class sizeof(QLayout)=1
type QLayout struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// invalidate()
func (this *QLayout) Invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout10invalidateEv
    // invoke: void invalidate()
    C.C_ZN7QLayout10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "invalidate", args)
  }

  return
}

// totalHeightForWidth(int)
func (this *QLayout) Totalheightforwidth(args ...interface{}) (ret interface{}) {
  // totalHeightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout19totalHeightForWidthEi
    // invoke: int totalHeightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QLayout19totalHeightForWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "totalHeightForWidth", args)
  }

  return
}

// totalMinimumSize()
func (this *QLayout) Totalminimumsize(args ...interface{}) (ret interface{}) {
  // totalMinimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout16totalMinimumSizeEv
    // invoke: QSize totalMinimumSize()
    var ret0 = C.C_ZNK7QLayout16totalMinimumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "totalMinimumSize", args)
  }

  return
}

// ~QLayout()
func (this *QLayout) Freeqlayout(args ...interface{}) () {
  // ~QLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayoutD0Ev
    // invoke: void ~QLayout()
    C.C_ZN7QLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "~QLayout", args)
  }

  return
}

// contentsMargins()
func (this *QLayout) Contentsmargins(args ...interface{}) (ret interface{}) {
  // contentsMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout15contentsMarginsEv
    // invoke: QMargins contentsMargins()
    var ret0 = C.C_ZNK7QLayout15contentsMarginsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMargins{}) // "QMargins"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "contentsMargins", args)
  }

  return
}

// contentsRect()
func (this *QLayout) Contentsrect(args ...interface{}) (ret interface{}) {
  // contentsRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout12contentsRectEv
    // invoke: QRect contentsRect()
    var ret0 = C.C_ZNK7QLayout12contentsRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "contentsRect", args)
  }

  return
}

// setMenuBar(class QWidget *)
func (this *QLayout) Setmenubar(args ...interface{}) () {
  // setMenuBar(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout10setMenuBarEP7QWidget
    // invoke: void setMenuBar(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QLayout10setMenuBarEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setMenuBar", args)
  }

  return
}

// sizeConstraint()
func (this *QLayout) Sizeconstraint(args ...interface{}) () {
  // sizeConstraint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout14sizeConstraintEv
    // invoke: QLayout::SizeConstraint sizeConstraint()
    C.C_ZNK7QLayout14sizeConstraintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "sizeConstraint", args)
  }

  return
}

// expandingDirections()
func (this *QLayout) Expandingdirections(args ...interface{}) () {
  // expandingDirections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout19expandingDirectionsEv
    // invoke: Qt::Orientations expandingDirections()
    C.C_ZNK7QLayout19expandingDirectionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "expandingDirections", args)
  }

  return
}

// layout()
func (this *QLayout) Layout(args ...interface{}) (ret interface{}) {
  // layout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout6layoutEv
    // invoke: QLayout * layout()
    var ret0 = C.C_ZN7QLayout6layoutEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLayout{}) // "QLayout *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "layout", args)
  }

  return
}

// removeWidget(class QWidget *)
func (this *QLayout) Removewidget(args ...interface{}) () {
  // removeWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout12removeWidgetEP7QWidget
    // invoke: void removeWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QLayout12removeWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "removeWidget", args)
  }

  return
}

// isEmpty()
func (this *QLayout) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK7QLayout7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "isEmpty", args)
  }

  return
}

// setContentsMargins(int, int, int, int)
func (this *QLayout) Setcontentsmargins(args ...interface{}) () {
  // setContentsMargins(int, int, int, int)
  // setContentsMargins(const class QMargins &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMargins{}) // "const QMargins &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout18setContentsMarginsEiiii
    // invoke: void setContentsMargins(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN7QLayout18setContentsMarginsEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN7QLayout18setContentsMarginsERK8QMargins
    // invoke: void setContentsMargins(const class QMargins &)
    var arg0 = args[0].(QMargins).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QLayout18setContentsMarginsERK8QMargins(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setContentsMargins", args)
  }

  return
}

// setEnabled(_Bool)
func (this *QLayout) Setenabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout10setEnabledEb
    // invoke: void setEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QLayout10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setEnabled", args)
  }

  return
}

// controlTypes()
func (this *QLayout) Controltypes(args ...interface{}) () {
  // controlTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout12controlTypesEv
    // invoke: QSizePolicy::ControlTypes controlTypes()
    C.C_ZNK7QLayout12controlTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "controlTypes", args)
  }

  return
}

// setGeometry(const class QRect &)
func (this *QLayout) Setgeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QLayout11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setGeometry", args)
  }

  return
}

// indexOf(class QWidget *)
func (this *QLayout) Indexof(args ...interface{}) (ret interface{}) {
  // indexOf(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout7indexOfEP7QWidget
    // invoke: int indexOf(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QLayout7indexOfEP7QWidget(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "indexOf", args)
  }

  return
}

// setSpacing(int)
func (this *QLayout) Setspacing(args ...interface{}) () {
  // setSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout10setSpacingEi
    // invoke: void setSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QLayout10setSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setSpacing", args)
  }

  return
}

// spacing()
func (this *QLayout) Spacing(args ...interface{}) (ret interface{}) {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout7spacingEv
    // invoke: int spacing()
    var ret0 = C.C_ZNK7QLayout7spacingEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "spacing", args)
  }

  return
}

// update()
func (this *QLayout) Update(args ...interface{}) () {
  // update()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout6updateEv
    // invoke: void update()
    C.C_ZN7QLayout6updateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "update", args)
  }

  return
}

// setMargin(int)
func (this *QLayout) Setmargin(args ...interface{}) () {
  // setMargin(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout9setMarginEi
    // invoke: void setMargin(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QLayout9setMarginEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "setMargin", args)
  }

  return
}

// QLayout(class QWidget *)
func NewQLayout(args ...interface{}) *QLayout {
  // QLayout(class QWidget *)
  // QLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayoutC1EP7QWidget
    // invoke: void QLayout(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QLayoutC2EP7QWidget(arg0)
    return &QLayout{qclsinst:qthis}
  case 1:
    // invoke: _ZN7QLayoutC1Ev
    // invoke: void QLayout()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QLayoutC2Ev()
    return &QLayout{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLayout", "QLayout", args)
  }

  return nil // QLayout{qclsinst:qthis}
}

// removeItem(class QLayoutItem *)
func (this *QLayout) Removeitem(args ...interface{}) () {
  // removeItem(class QLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout10removeItemEP11QLayoutItem
    // invoke: void removeItem(class QLayoutItem *)
    var arg0 = args[0].(QLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QLayout10removeItemEP11QLayoutItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "removeItem", args)
  }

  return
}

// addWidget(class QWidget *)
func (this *QLayout) Addwidget(args ...interface{}) () {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout9addWidgetEP7QWidget
    // invoke: void addWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QLayout9addWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLayout", "addWidget", args)
  }

  return
}

// activate()
func (this *QLayout) Activate(args ...interface{}) (ret interface{}) {
  // activate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout8activateEv
    // invoke: bool activate()
    var ret0 = C.C_ZN7QLayout8activateEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "activate", args)
  }

  return
}

// parentWidget()
func (this *QLayout) Parentwidget(args ...interface{}) (ret interface{}) {
  // parentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout12parentWidgetEv
    // invoke: QWidget * parentWidget()
    var ret0 = C.C_ZNK7QLayout12parentWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "parentWidget", args)
  }

  return
}

// totalSizeHint()
func (this *QLayout) Totalsizehint(args ...interface{}) (ret interface{}) {
  // totalSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout13totalSizeHintEv
    // invoke: QSize totalSizeHint()
    var ret0 = C.C_ZNK7QLayout13totalSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "totalSizeHint", args)
  }

  return
}

// metaObject()
func (this *QLayout) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLayout", "metaObject", args)
  }

  return
}

// totalMaximumSize()
func (this *QLayout) Totalmaximumsize(args ...interface{}) (ret interface{}) {
  // totalMaximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout16totalMaximumSizeEv
    // invoke: QSize totalMaximumSize()
    var ret0 = C.C_ZNK7QLayout16totalMaximumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "totalMaximumSize", args)
  }

  return
}

// geometry()
func (this *QLayout) Geometry(args ...interface{}) (ret interface{}) {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout8geometryEv
    // invoke: QRect geometry()
    var ret0 = C.C_ZNK7QLayout8geometryEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "geometry", args)
  }

  return
}

// getContentsMargins(int *, int *, int *, int *)
func (this *QLayout) Getcontentsmargins(args ...interface{}) () {
  // getContentsMargins(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout18getContentsMarginsEPiS0_S0_S0_
    // invoke: void getContentsMargins(int *, int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK7QLayout18getContentsMarginsEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QLayout", "getContentsMargins", args)
  }

  return
}

// minimumSize()
func (this *QLayout) Minimumsize(args ...interface{}) (ret interface{}) {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout11minimumSizeEv
    // invoke: QSize minimumSize()
    var ret0 = C.C_ZNK7QLayout11minimumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "minimumSize", args)
  }

  return
}

// closestAcceptableSize(const class QWidget *, const class QSize &)
func (this *QLayout) Closestacceptablesize_S(args ...interface{}) (ret interface{}) {
  // closestAcceptableSize(const class QWidget *, const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[0][1] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QLayout21closestAcceptableSizeEPK7QWidgetRK5QSize
    // invoke: QSize closestAcceptableSize(const class QWidget *, const class QSize &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSize).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QLayout21closestAcceptableSizeEPK7QWidgetRK5QSize(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "closestAcceptableSize", args)
  }

  return
}

// margin()
func (this *QLayout) Margin(args ...interface{}) (ret interface{}) {
  // margin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout6marginEv
    // invoke: int margin()
    var ret0 = C.C_ZNK7QLayout6marginEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "margin", args)
  }

  return
}

// isEnabled()
func (this *QLayout) Isenabled(args ...interface{}) (ret interface{}) {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout9isEnabledEv
    // invoke: bool isEnabled()
    var ret0 = C.C_ZNK7QLayout9isEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "isEnabled", args)
  }

  return
}

// maximumSize()
func (this *QLayout) Maximumsize(args ...interface{}) (ret interface{}) {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout11maximumSizeEv
    // invoke: QSize maximumSize()
    var ret0 = C.C_ZNK7QLayout11maximumSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "maximumSize", args)
  }

  return
}

// menuBar()
func (this *QLayout) Menubar(args ...interface{}) (ret interface{}) {
  // menuBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QLayout7menuBarEv
    // invoke: QWidget * menuBar()
    var ret0 = C.C_ZNK7QLayout7menuBarEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QLayout", "menuBar", args)
  }

  return
}

// <= body block end

