package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  bool QPageLayout::setRightMargin(qreal rightMargin);
extern void _ZN11QPageLayout14setRightMarginEd(void* qthis, double arg0);
  // proto:  void QPageLayout::swap(QPageLayout & other);
extern void demth_ZN11QPageLayout4swapERS_(void* qthis, void* arg0);
  // proto:  QMargins QPageLayout::marginsPoints();
extern void _ZNK11QPageLayout13marginsPointsEv(void* qthis);
  // proto:  QMargins QPageLayout::marginsPixels(int resolution);
extern void _ZNK11QPageLayout13marginsPixelsEi(void* qthis, int32_t arg0);
  // proto:  bool QPageLayout::isValid();
extern void _ZNK11QPageLayout7isValidEv(void* qthis);
  // proto:  QRectF QPageLayout::fullRect();
extern void _ZNK11QPageLayout8fullRectEv(void* qthis);
  // proto:  QRectF QPageLayout::paintRect();
extern void _ZNK11QPageLayout9paintRectEv(void* qthis);
  // proto:  void QPageLayout::setMinimumMargins(const QMarginsF & minMargins);
extern void _ZN11QPageLayout17setMinimumMarginsERK9QMarginsF(void* qthis, void* arg0);
  // proto:  bool QPageLayout::setLeftMargin(qreal leftMargin);
extern void _ZN11QPageLayout13setLeftMarginEd(void* qthis, double arg0);
  // proto:  bool QPageLayout::setBottomMargin(qreal bottomMargin);
extern void _ZN11QPageLayout15setBottomMarginEd(void* qthis, double arg0);
  // proto:  QRect QPageLayout::fullRectPoints();
extern void _ZNK11QPageLayout14fullRectPointsEv(void* qthis);
  // proto:  QMarginsF QPageLayout::minimumMargins();
extern void _ZNK11QPageLayout14minimumMarginsEv(void* qthis);
  // proto:  QPageSize QPageLayout::pageSize();
extern void _ZNK11QPageLayout8pageSizeEv(void* qthis);
  // proto:  bool QPageLayout::setTopMargin(qreal topMargin);
extern void _ZN11QPageLayout12setTopMarginEd(void* qthis, double arg0);
  // proto:  void QPageLayout::QPageLayout();
extern void* dector_ZN11QPageLayoutC1Ev();
extern void _ZN11QPageLayoutC1Ev(void* qthis);
  // proto:  void QPageLayout::QPageLayout(const QPageLayout & other);
extern void* dector_ZN11QPageLayoutC1ERKS_(void* arg0);
extern void _ZN11QPageLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  void QPageLayout::~QPageLayout();
extern void _ZN11QPageLayoutD0Ev(void* qthis);
  // proto:  QRect QPageLayout::fullRectPixels(int resolution);
extern void _ZNK11QPageLayout14fullRectPixelsEi(void* qthis, int32_t arg0);
  // proto:  QMarginsF QPageLayout::margins();
extern void _ZNK11QPageLayout7marginsEv(void* qthis);
  // proto:  QRect QPageLayout::paintRectPoints();
extern void _ZNK11QPageLayout15paintRectPointsEv(void* qthis);
  // proto:  QRect QPageLayout::paintRectPixels(int resolution);
extern void _ZNK11QPageLayout15paintRectPixelsEi(void* qthis, int32_t arg0);
  // proto:  void QPageLayout::setPageSize(const QPageSize & pageSize, const QMarginsF & minMargins);
extern void _ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF(void* qthis, void* arg0, void* arg1);
  // proto:  QMarginsF QPageLayout::maximumMargins();
extern void _ZNK11QPageLayout14maximumMarginsEv(void* qthis);
  // proto:  bool QPageLayout::isEquivalentTo(const QPageLayout & other);
extern void _ZNK11QPageLayout14isEquivalentToERKS_(void* qthis, void* arg0);
  // proto:  bool QPageLayout::setMargins(const QMarginsF & margins);
extern void _ZN11QPageLayout10setMarginsERK9QMarginsF(void* qthis, void* arg0);
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

  // proto:  bool QPageLayout::setRightMargin(qreal rightMargin);
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
    C._ZN11QPageLayout14setRightMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setRightMargin", args)
  }

}

  // proto:  void QPageLayout::swap(QPageLayout & other);
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
    C.demth_ZN11QPageLayout4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "swap", args)
  }

}

  // proto:  QMargins QPageLayout::marginsPoints();
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
    C._ZNK11QPageLayout13marginsPointsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "marginsPoints", args)
  }

}

  // proto:  QMargins QPageLayout::marginsPixels(int resolution);
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
    C._ZNK11QPageLayout13marginsPixelsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "marginsPixels", args)
  }

}

  // proto:  bool QPageLayout::isValid();
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
    C._ZNK11QPageLayout7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "isValid", args)
  }

}

  // proto:  QRectF QPageLayout::fullRect();
func (this *QPageLayout) fullRect(args ...interface{}) () {
  // fullRect()
  // fullRect(enum QPageLayout::Unit)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QPageLayout::Unit"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout8fullRectEv
    // invoke: QRectF fullRect()
    C._ZNK11QPageLayout8fullRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRect", args)
  }

}

  // proto:  QRectF QPageLayout::paintRect();
func (this *QPageLayout) paintRect(args ...interface{}) () {
  // paintRect()
  // paintRect(enum QPageLayout::Unit)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QPageLayout::Unit"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout9paintRectEv
    // invoke: QRectF paintRect()
    C._ZNK11QPageLayout9paintRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRect", args)
  }

}

  // proto:  void QPageLayout::setMinimumMargins(const QMarginsF & minMargins);
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
    C._ZN11QPageLayout17setMinimumMarginsERK9QMarginsF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setMinimumMargins", args)
  }

}

  // proto:  bool QPageLayout::setLeftMargin(qreal leftMargin);
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
    C._ZN11QPageLayout13setLeftMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setLeftMargin", args)
  }

}

  // proto:  bool QPageLayout::setBottomMargin(qreal bottomMargin);
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
    C._ZN11QPageLayout15setBottomMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setBottomMargin", args)
  }

}

  // proto:  QRect QPageLayout::fullRectPoints();
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
    C._ZNK11QPageLayout14fullRectPointsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRectPoints", args)
  }

}

  // proto:  QMarginsF QPageLayout::minimumMargins();
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
    C._ZNK11QPageLayout14minimumMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "minimumMargins", args)
  }

}

  // proto:  QPageSize QPageLayout::pageSize();
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
    C._ZNK11QPageLayout8pageSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "pageSize", args)
  }

}

  // proto:  bool QPageLayout::setTopMargin(qreal topMargin);
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
    C._ZN11QPageLayout12setTopMarginEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setTopMargin", args)
  }

}

  // proto:  void QPageLayout::QPageLayout();
func NewQPageLayout(args ...interface{}) QPageLayout {
  return QPageLayout{}
}

  // proto:  void QPageLayout::~QPageLayout();
func (this *QPageLayout) FreeQPageLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPageLayout", "~QPageLayout", args)
  }

}

  // proto:  QRect QPageLayout::fullRectPixels(int resolution);
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
    C._ZNK11QPageLayout14fullRectPixelsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRectPixels", args)
  }

}

  // proto:  QMarginsF QPageLayout::margins();
func (this *QPageLayout) margins(args ...interface{}) () {
  // margins(enum QPageLayout::Unit)
  // margins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QPageLayout::Unit"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout7marginsEv
    // invoke: QMarginsF margins()
    C._ZNK11QPageLayout7marginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "margins", args)
  }

}

  // proto:  QRect QPageLayout::paintRectPoints();
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
    C._ZNK11QPageLayout15paintRectPointsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRectPoints", args)
  }

}

  // proto:  QRect QPageLayout::paintRectPixels(int resolution);
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
    C._ZNK11QPageLayout15paintRectPixelsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRectPixels", args)
  }

}

  // proto:  void QPageLayout::setPageSize(const QPageSize & pageSize, const QMarginsF & minMargins);
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
    C._ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPageLayout", "setPageSize", args)
  }

}

  // proto:  QMarginsF QPageLayout::maximumMargins();
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
    C._ZNK11QPageLayout14maximumMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "maximumMargins", args)
  }

}

  // proto:  bool QPageLayout::isEquivalentTo(const QPageLayout & other);
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
    C._ZNK11QPageLayout14isEquivalentToERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "isEquivalentTo", args)
  }

}

  // proto:  bool QPageLayout::setMargins(const QMarginsF & margins);
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
    C._ZN11QPageLayout10setMarginsERK9QMarginsF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setMargins", args)
  }

}

// <= body block end

