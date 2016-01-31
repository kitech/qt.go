package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtGui/qpagelayout.h
// dst-file: /src/gui/qpagelayout.go
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
  // proto:  QPageLayout::Orientation QPageLayout::orientation();
extern void C_ZNK11QPageLayout11orientationEv(void* qthis); // 4
  // proto:  QPageSize QPageLayout::pageSize();
extern void C_ZNK11QPageLayout8pageSizeEv(void* qthis); // 4
  // proto:  void QPageLayout::setPageSize(const QPageSize & pageSize, const QMarginsF & minMargins);
extern void C_ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QRect QPageLayout::paintRectPoints();
extern void C_ZNK11QPageLayout15paintRectPointsEv(void* qthis); // 4
  // proto:  QMarginsF QPageLayout::margins();
extern void C_ZNK11QPageLayout7marginsEv(void* qthis); // 4
  // proto:  QMarginsF QPageLayout::minimumMargins();
extern void C_ZNK11QPageLayout14minimumMarginsEv(void* qthis); // 4
  // proto:  void QPageLayout::setMinimumMargins(const QMarginsF & minMargins);
extern void C_ZN11QPageLayout17setMinimumMarginsERK9QMarginsF(void* qthis, void* arg0); // 4
  // proto:  void QPageLayout::QPageLayout(const QPageLayout & other);
extern void C_ZN11QPageLayoutC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QPageLayout::QPageLayout();
extern void C_ZN11QPageLayoutC2Ev(void* qthis); // 3
  // proto:  QRect QPageLayout::fullRectPoints();
extern void C_ZNK11QPageLayout14fullRectPointsEv(void* qthis); // 4
  // proto:  QRectF QPageLayout::fullRect();
extern void C_ZNK11QPageLayout8fullRectEv(void* qthis); // 4
  // proto:  QMargins QPageLayout::marginsPoints();
extern void C_ZNK11QPageLayout13marginsPointsEv(void* qthis); // 4
  // proto:  void QPageLayout::swap(QPageLayout & other);
extern void C_ZN11QPageLayout4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QPageLayout::~QPageLayout();
extern void C_ZN11QPageLayoutD2Ev(void* qthis); // 4
  // proto:  QMarginsF QPageLayout::maximumMargins();
extern void C_ZNK11QPageLayout14maximumMarginsEv(void* qthis); // 4
  // proto:  QRect QPageLayout::fullRectPixels(int resolution);
extern void C_ZNK11QPageLayout14fullRectPixelsEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QPageLayout::setTopMargin(qreal topMargin);
extern void C_ZN11QPageLayout12setTopMarginEd(void* qthis, double arg0); // 4
  // proto:  bool QPageLayout::isValid();
extern void C_ZNK11QPageLayout7isValidEv(void* qthis); // 4
  // proto:  QPageLayout::Unit QPageLayout::units();
extern void C_ZNK11QPageLayout5unitsEv(void* qthis); // 4
  // proto:  QRect QPageLayout::paintRectPixels(int resolution);
extern void C_ZNK11QPageLayout15paintRectPixelsEi(void* qthis, int32_t arg0); // 4
  // proto:  QRectF QPageLayout::paintRect();
extern void C_ZNK11QPageLayout9paintRectEv(void* qthis); // 4
  // proto:  bool QPageLayout::setLeftMargin(qreal leftMargin);
extern void C_ZN11QPageLayout13setLeftMarginEd(void* qthis, double arg0); // 4
  // proto:  bool QPageLayout::setRightMargin(qreal rightMargin);
extern void C_ZN11QPageLayout14setRightMarginEd(void* qthis, double arg0); // 4
  // proto:  bool QPageLayout::isEquivalentTo(const QPageLayout & other);
extern void C_ZNK11QPageLayout14isEquivalentToERKS_(void* qthis, void* arg0); // 4
  // proto:  QMargins QPageLayout::marginsPixels(int resolution);
extern void C_ZNK11QPageLayout13marginsPixelsEi(void* qthis, int32_t arg0); // 4
  // proto:  QPageLayout::Mode QPageLayout::mode();
extern void C_ZNK11QPageLayout4modeEv(void* qthis); // 4
  // proto:  bool QPageLayout::setMargins(const QMarginsF & margins);
extern void C_ZN11QPageLayout10setMarginsERK9QMarginsF(void* qthis, void* arg0); // 4
  // proto:  bool QPageLayout::setBottomMargin(qreal bottomMargin);
extern void C_ZN11QPageLayout15setBottomMarginEd(void* qthis, double arg0); // 4
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

// class sizeof(QPageLayout)=1
type QPageLayout struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// orientation()
func (this *QPageLayout) orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout11orientationEv
    // invoke: QPageLayout::Orientation orientation()
    C.C_ZNK11QPageLayout11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "orientation", args)
  }

}

// pageSize()
func (this *QPageLayout) pageSize(args ...interface{}) () {
  // pageSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout8pageSizeEv
    // invoke: QPageSize pageSize()
    C.C_ZNK11QPageLayout8pageSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "pageSize", args)
  }

}

// setPageSize(const class QPageSize &, const class QMarginsF &)
func (this *QPageLayout) setPageSize(args ...interface{}) () {
  // setPageSize(const class QPageSize &, const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageSize{}) // "const QPageSize &"
  vtys[0][1] = reflect.TypeOf(QMarginsF{}) // "const QMarginsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF
    // invoke: void setPageSize(const class QPageSize &, const class QMarginsF &)
    var arg0 = args[0].(QPageSize).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QMarginsF).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPageLayout", "setPageSize", args)
  }

}

// paintRectPoints()
func (this *QPageLayout) paintRectPoints(args ...interface{}) () {
  // paintRectPoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout15paintRectPointsEv
    // invoke: QRect paintRectPoints()
    C.C_ZNK11QPageLayout15paintRectPointsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRectPoints", args)
  }

}

// margins()
func (this *QPageLayout) margins(args ...interface{}) () {
  // margins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout7marginsEv
    // invoke: QMarginsF margins()
    C.C_ZNK11QPageLayout7marginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "margins", args)
  }

}

// minimumMargins()
func (this *QPageLayout) minimumMargins(args ...interface{}) () {
  // minimumMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout14minimumMarginsEv
    // invoke: QMarginsF minimumMargins()
    C.C_ZNK11QPageLayout14minimumMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "minimumMargins", args)
  }

}

// setMinimumMargins(const class QMarginsF &)
func (this *QPageLayout) setMinimumMargins(args ...interface{}) () {
  // setMinimumMargins(const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMarginsF{}) // "const QMarginsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout17setMinimumMarginsERK9QMarginsF
    // invoke: void setMinimumMargins(const class QMarginsF &)
    var arg0 = args[0].(QMarginsF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QPageLayout17setMinimumMarginsERK9QMarginsF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setMinimumMargins", args)
  }

}

// QPageLayout(const class QPageLayout &)
func NewQPageLayout(args ...interface{}) QPageLayout {
  // QPageLayout(const class QPageLayout &)
  // QPageLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageLayout{}) // "const QPageLayout &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayoutC1ERKS_
    // invoke: void QPageLayout(const class QPageLayout &)
    var arg0 = args[0].(QPageLayout).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QPageLayoutC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN11QPageLayoutC1Ev
    // invoke: void QPageLayout()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QPageLayoutC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QPageLayout", "QPageLayout", args)
  }

  return QPageLayout{}
}

// fullRectPoints()
func (this *QPageLayout) fullRectPoints(args ...interface{}) () {
  // fullRectPoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout14fullRectPointsEv
    // invoke: QRect fullRectPoints()
    C.C_ZNK11QPageLayout14fullRectPointsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRectPoints", args)
  }

}

// fullRect()
func (this *QPageLayout) fullRect(args ...interface{}) () {
  // fullRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout8fullRectEv
    // invoke: QRectF fullRect()
    C.C_ZNK11QPageLayout8fullRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRect", args)
  }

}

// marginsPoints()
func (this *QPageLayout) marginsPoints(args ...interface{}) () {
  // marginsPoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout13marginsPointsEv
    // invoke: QMargins marginsPoints()
    C.C_ZNK11QPageLayout13marginsPointsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "marginsPoints", args)
  }

}

// swap(class QPageLayout &)
func (this *QPageLayout) swap(args ...interface{}) () {
  // swap(class QPageLayout &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageLayout{}) // "QPageLayout &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout4swapERS_
    // invoke: void swap(class QPageLayout &)
    var arg0 = args[0].(QPageLayout).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QPageLayout4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "swap", args)
  }

}

// ~QPageLayout()
func (this *QPageLayout) FreeQPageLayout(args ...interface{}) () {
  // ~QPageLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayoutD0Ev
    // invoke: void ~QPageLayout()
    C.C_ZN11QPageLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "~QPageLayout", args)
  }

}

// maximumMargins()
func (this *QPageLayout) maximumMargins(args ...interface{}) () {
  // maximumMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout14maximumMarginsEv
    // invoke: QMarginsF maximumMargins()
    C.C_ZNK11QPageLayout14maximumMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "maximumMargins", args)
  }

}

// fullRectPixels(int)
func (this *QPageLayout) fullRectPixels(args ...interface{}) () {
  // fullRectPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout14fullRectPixelsEi
    // invoke: QRect fullRectPixels(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QPageLayout14fullRectPixelsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRectPixels", args)
  }

}

// setTopMargin(qreal)
func (this *QPageLayout) setTopMargin(args ...interface{}) () {
  // setTopMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout12setTopMarginEd
    // invoke: bool setTopMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QPageLayout12setTopMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setTopMargin", args)
  }

}

// isValid()
func (this *QPageLayout) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout7isValidEv
    // invoke: bool isValid()
    C.C_ZNK11QPageLayout7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "isValid", args)
  }

}

// units()
func (this *QPageLayout) units(args ...interface{}) () {
  // units()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout5unitsEv
    // invoke: QPageLayout::Unit units()
    C.C_ZNK11QPageLayout5unitsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "units", args)
  }

}

// paintRectPixels(int)
func (this *QPageLayout) paintRectPixels(args ...interface{}) () {
  // paintRectPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout15paintRectPixelsEi
    // invoke: QRect paintRectPixels(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QPageLayout15paintRectPixelsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRectPixels", args)
  }

}

// paintRect()
func (this *QPageLayout) paintRect(args ...interface{}) () {
  // paintRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout9paintRectEv
    // invoke: QRectF paintRect()
    C.C_ZNK11QPageLayout9paintRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRect", args)
  }

}

// setLeftMargin(qreal)
func (this *QPageLayout) setLeftMargin(args ...interface{}) () {
  // setLeftMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout13setLeftMarginEd
    // invoke: bool setLeftMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QPageLayout13setLeftMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setLeftMargin", args)
  }

}

// setRightMargin(qreal)
func (this *QPageLayout) setRightMargin(args ...interface{}) () {
  // setRightMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout14setRightMarginEd
    // invoke: bool setRightMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QPageLayout14setRightMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setRightMargin", args)
  }

}

// isEquivalentTo(const class QPageLayout &)
func (this *QPageLayout) isEquivalentTo(args ...interface{}) () {
  // isEquivalentTo(const class QPageLayout &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageLayout{}) // "const QPageLayout &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout14isEquivalentToERKS_
    // invoke: bool isEquivalentTo(const class QPageLayout &)
    var arg0 = args[0].(QPageLayout).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QPageLayout14isEquivalentToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "isEquivalentTo", args)
  }

}

// marginsPixels(int)
func (this *QPageLayout) marginsPixels(args ...interface{}) () {
  // marginsPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout13marginsPixelsEi
    // invoke: QMargins marginsPixels(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QPageLayout13marginsPixelsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "marginsPixels", args)
  }

}

// mode()
func (this *QPageLayout) mode(args ...interface{}) () {
  // mode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout4modeEv
    // invoke: QPageLayout::Mode mode()
    C.C_ZNK11QPageLayout4modeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "mode", args)
  }

}

// setMargins(const class QMarginsF &)
func (this *QPageLayout) setMargins(args ...interface{}) () {
  // setMargins(const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMarginsF{}) // "const QMarginsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout10setMarginsERK9QMarginsF
    // invoke: bool setMargins(const class QMarginsF &)
    var arg0 = args[0].(QMarginsF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QPageLayout10setMarginsERK9QMarginsF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setMargins", args)
  }

}

// setBottomMargin(qreal)
func (this *QPageLayout) setBottomMargin(args ...interface{}) () {
  // setBottomMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout15setBottomMarginEd
    // invoke: bool setBottomMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QPageLayout15setBottomMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setBottomMargin", args)
  }

}

// <= body block end

