package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
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
extern void* C_ZNK11QPageLayout8pageSizeEv(void* qthis); // 4
  // proto:  void QPageLayout::setPageSize(const QPageSize & pageSize, const QMarginsF & minMargins);
extern void C_ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QRect QPageLayout::paintRectPoints();
extern void* C_ZNK11QPageLayout15paintRectPointsEv(void* qthis); // 4
  // proto:  QMarginsF QPageLayout::margins();
extern void* C_ZNK11QPageLayout7marginsEv(void* qthis); // 4
  // proto:  QMarginsF QPageLayout::minimumMargins();
extern void* C_ZNK11QPageLayout14minimumMarginsEv(void* qthis); // 4
  // proto:  void QPageLayout::setMinimumMargins(const QMarginsF & minMargins);
extern void C_ZN11QPageLayout17setMinimumMarginsERK9QMarginsF(void* qthis, void* arg0); // 4
  // proto:  void QPageLayout::QPageLayout(const QPageLayout & other);
extern void* C_ZN11QPageLayoutC2ERKS_(void* arg0); // 3
  // proto:  void QPageLayout::QPageLayout();
extern void* C_ZN11QPageLayoutC2Ev(); // 3
  // proto:  QRect QPageLayout::fullRectPoints();
extern void* C_ZNK11QPageLayout14fullRectPointsEv(void* qthis); // 4
  // proto:  QRectF QPageLayout::fullRect();
extern void* C_ZNK11QPageLayout8fullRectEv(void* qthis); // 4
  // proto:  QMargins QPageLayout::marginsPoints();
extern void* C_ZNK11QPageLayout13marginsPointsEv(void* qthis); // 4
  // proto:  void QPageLayout::swap(QPageLayout & other);
extern void C_ZN11QPageLayout4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QPageLayout::~QPageLayout();
extern void C_ZN11QPageLayoutD2Ev(void* qthis); // 4
  // proto:  QMarginsF QPageLayout::maximumMargins();
extern void* C_ZNK11QPageLayout14maximumMarginsEv(void* qthis); // 4
  // proto:  QRect QPageLayout::fullRectPixels(int resolution);
extern void* C_ZNK11QPageLayout14fullRectPixelsEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QPageLayout::setTopMargin(qreal topMargin);
extern bool C_ZN11QPageLayout12setTopMarginEd(void* qthis, double arg0); // 4
  // proto:  bool QPageLayout::isValid();
extern bool C_ZNK11QPageLayout7isValidEv(void* qthis); // 4
  // proto:  QPageLayout::Unit QPageLayout::units();
extern void C_ZNK11QPageLayout5unitsEv(void* qthis); // 4
  // proto:  QRect QPageLayout::paintRectPixels(int resolution);
extern void* C_ZNK11QPageLayout15paintRectPixelsEi(void* qthis, int32_t arg0); // 4
  // proto:  QRectF QPageLayout::paintRect();
extern void* C_ZNK11QPageLayout9paintRectEv(void* qthis); // 4
  // proto:  bool QPageLayout::setLeftMargin(qreal leftMargin);
extern bool C_ZN11QPageLayout13setLeftMarginEd(void* qthis, double arg0); // 4
  // proto:  bool QPageLayout::setRightMargin(qreal rightMargin);
extern bool C_ZN11QPageLayout14setRightMarginEd(void* qthis, double arg0); // 4
  // proto:  bool QPageLayout::isEquivalentTo(const QPageLayout & other);
extern bool C_ZNK11QPageLayout14isEquivalentToERKS_(void* qthis, void* arg0); // 4
  // proto:  QMargins QPageLayout::marginsPixels(int resolution);
extern void* C_ZNK11QPageLayout13marginsPixelsEi(void* qthis, int32_t arg0); // 4
  // proto:  QPageLayout::Mode QPageLayout::mode();
extern void C_ZNK11QPageLayout4modeEv(void* qthis); // 4
  // proto:  bool QPageLayout::setMargins(const QMarginsF & margins);
extern bool C_ZN11QPageLayout10setMarginsERK9QMarginsF(void* qthis, void* arg0); // 4
  // proto:  bool QPageLayout::setBottomMargin(qreal bottomMargin);
extern bool C_ZN11QPageLayout15setBottomMarginEd(void* qthis, double arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QPageLayout)=1
type QPageLayout struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// orientation()
func (this *QPageLayout) Orientation(args ...interface{}) () {
  // orientation()
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
    // invoke: _ZNK11QPageLayout11orientationEv
    // invoke: QPageLayout::Orientation orientation()
    C.C_ZNK11QPageLayout11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "orientation", args)
  }

  return
}

// pageSize()
func (this *QPageLayout) Pagesize(args ...interface{}) (ret interface{}) {
  // pageSize()
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
    // invoke: _ZNK11QPageLayout8pageSizeEv
    // invoke: QPageSize pageSize()
    var ret0 = C.C_ZNK11QPageLayout8pageSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPageSize{}) // "QPageSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "pageSize", args)
  }

  return
}

// setPageSize(const class QPageSize &, const class QMarginsF &)
func (this *QPageLayout) Setpagesize(args ...interface{}) () {
  // setPageSize(const class QPageSize &, const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageSize{}) // "const QPageSize &"
  vtys[0][1] = reflect.TypeOf(qtcore.QMarginsF{}) // "const QMarginsF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF
    // invoke: void setPageSize(const class QPageSize &, const class QMarginsF &)
    var arg0 = args[0].(*QPageSize).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QMarginsF).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPageLayout", "setPageSize", args)
  }

  return
}

// paintRectPoints()
func (this *QPageLayout) Paintrectpoints(args ...interface{}) (ret interface{}) {
  // paintRectPoints()
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
    // invoke: _ZNK11QPageLayout15paintRectPointsEv
    // invoke: QRect paintRectPoints()
    var ret0 = C.C_ZNK11QPageLayout15paintRectPointsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRectPoints", args)
  }

  return
}

// margins()
func (this *QPageLayout) Margins(args ...interface{}) (ret interface{}) {
  // margins()
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
    // invoke: _ZNK11QPageLayout7marginsEv
    // invoke: QMarginsF margins()
    var ret0 = C.C_ZNK11QPageLayout7marginsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QMarginsF{}) // "QMarginsF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "margins", args)
  }

  return
}

// minimumMargins()
func (this *QPageLayout) Minimummargins(args ...interface{}) (ret interface{}) {
  // minimumMargins()
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
    // invoke: _ZNK11QPageLayout14minimumMarginsEv
    // invoke: QMarginsF minimumMargins()
    var ret0 = C.C_ZNK11QPageLayout14minimumMarginsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QMarginsF{}) // "QMarginsF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "minimumMargins", args)
  }

  return
}

// setMinimumMargins(const class QMarginsF &)
func (this *QPageLayout) Setminimummargins(args ...interface{}) () {
  // setMinimumMargins(const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QMarginsF{}) // "const QMarginsF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout17setMinimumMarginsERK9QMarginsF
    // invoke: void setMinimumMargins(const class QMarginsF &)
    var arg0 = args[0].(*qtcore.QMarginsF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QPageLayout17setMinimumMarginsERK9QMarginsF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "setMinimumMargins", args)
  }

  return
}

// QPageLayout(const class QPageLayout &)
func NewQPageLayout(args ...interface{}) *QPageLayout {
  // QPageLayout(const class QPageLayout &)
  // QPageLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageLayout{}) // "const QPageLayout &"
  vtys[1] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayoutC1ERKS_
    // invoke: void QPageLayout(const class QPageLayout &)
    var arg0 = args[0].(*QPageLayout).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPageLayoutC2ERKS_(arg0)
    return &QPageLayout{Qclsinst:qthis}
  case 1:
    // invoke: _ZN11QPageLayoutC1Ev
    // invoke: void QPageLayout()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QPageLayoutC2Ev()
    return &QPageLayout{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPageLayout", "QPageLayout", args)
  }

  return nil // QPageLayout{Qclsinst:qthis}
}

// fullRectPoints()
func (this *QPageLayout) Fullrectpoints(args ...interface{}) (ret interface{}) {
  // fullRectPoints()
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
    // invoke: _ZNK11QPageLayout14fullRectPointsEv
    // invoke: QRect fullRectPoints()
    var ret0 = C.C_ZNK11QPageLayout14fullRectPointsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRectPoints", args)
  }

  return
}

// fullRect()
func (this *QPageLayout) Fullrect(args ...interface{}) (ret interface{}) {
  // fullRect()
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
    // invoke: _ZNK11QPageLayout8fullRectEv
    // invoke: QRectF fullRect()
    var ret0 = C.C_ZNK11QPageLayout8fullRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRect", args)
  }

  return
}

// marginsPoints()
func (this *QPageLayout) Marginspoints(args ...interface{}) (ret interface{}) {
  // marginsPoints()
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
    // invoke: _ZNK11QPageLayout13marginsPointsEv
    // invoke: QMargins marginsPoints()
    var ret0 = C.C_ZNK11QPageLayout13marginsPointsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QMargins{}) // "QMargins"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "marginsPoints", args)
  }

  return
}

// swap(class QPageLayout &)
func (this *QPageLayout) Swap(args ...interface{}) () {
  // swap(class QPageLayout &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageLayout{}) // "QPageLayout &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout4swapERS_
    // invoke: void swap(class QPageLayout &)
    var arg0 = args[0].(*QPageLayout).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QPageLayout4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageLayout", "swap", args)
  }

  return
}

// ~QPageLayout()
func (this *QPageLayout) Freeqpagelayout(args ...interface{}) () {
  // ~QPageLayout()
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
    // invoke: _ZN11QPageLayoutD0Ev
    // invoke: void ~QPageLayout()
    C.C_ZN11QPageLayoutD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "~QPageLayout", args)
  }

  return
}

// maximumMargins()
func (this *QPageLayout) Maximummargins(args ...interface{}) (ret interface{}) {
  // maximumMargins()
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
    // invoke: _ZNK11QPageLayout14maximumMarginsEv
    // invoke: QMarginsF maximumMargins()
    var ret0 = C.C_ZNK11QPageLayout14maximumMarginsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QMarginsF{}) // "QMarginsF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "maximumMargins", args)
  }

  return
}

// fullRectPixels(int)
func (this *QPageLayout) Fullrectpixels(args ...interface{}) (ret interface{}) {
  // fullRectPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout14fullRectPixelsEi
    // invoke: QRect fullRectPixels(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QPageLayout14fullRectPixelsEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRectPixels", args)
  }

  return
}

// setTopMargin(qreal)
func (this *QPageLayout) Settopmargin(args ...interface{}) (ret interface{}) {
  // setTopMargin(qreal)
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
    // invoke: _ZN11QPageLayout12setTopMarginEd
    // invoke: bool setTopMargin(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QPageLayout12setTopMarginEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "setTopMargin", args)
  }

  return
}

// isValid()
func (this *QPageLayout) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
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
    // invoke: _ZNK11QPageLayout7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK11QPageLayout7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "isValid", args)
  }

  return
}

// units()
func (this *QPageLayout) Units(args ...interface{}) () {
  // units()
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
    // invoke: _ZNK11QPageLayout5unitsEv
    // invoke: QPageLayout::Unit units()
    C.C_ZNK11QPageLayout5unitsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "units", args)
  }

  return
}

// paintRectPixels(int)
func (this *QPageLayout) Paintrectpixels(args ...interface{}) (ret interface{}) {
  // paintRectPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout15paintRectPixelsEi
    // invoke: QRect paintRectPixels(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QPageLayout15paintRectPixelsEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRectPixels", args)
  }

  return
}

// paintRect()
func (this *QPageLayout) Paintrect(args ...interface{}) (ret interface{}) {
  // paintRect()
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
    // invoke: _ZNK11QPageLayout9paintRectEv
    // invoke: QRectF paintRect()
    var ret0 = C.C_ZNK11QPageLayout9paintRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRect", args)
  }

  return
}

// setLeftMargin(qreal)
func (this *QPageLayout) Setleftmargin(args ...interface{}) (ret interface{}) {
  // setLeftMargin(qreal)
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
    // invoke: _ZN11QPageLayout13setLeftMarginEd
    // invoke: bool setLeftMargin(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QPageLayout13setLeftMarginEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "setLeftMargin", args)
  }

  return
}

// setRightMargin(qreal)
func (this *QPageLayout) Setrightmargin(args ...interface{}) (ret interface{}) {
  // setRightMargin(qreal)
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
    // invoke: _ZN11QPageLayout14setRightMarginEd
    // invoke: bool setRightMargin(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QPageLayout14setRightMarginEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "setRightMargin", args)
  }

  return
}

// isEquivalentTo(const class QPageLayout &)
func (this *QPageLayout) Isequivalentto(args ...interface{}) (ret interface{}) {
  // isEquivalentTo(const class QPageLayout &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageLayout{}) // "const QPageLayout &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout14isEquivalentToERKS_
    // invoke: bool isEquivalentTo(const class QPageLayout &)
    var arg0 = args[0].(*QPageLayout).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QPageLayout14isEquivalentToERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "isEquivalentTo", args)
  }

  return
}

// marginsPixels(int)
func (this *QPageLayout) Marginspixels(args ...interface{}) (ret interface{}) {
  // marginsPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QPageLayout13marginsPixelsEi
    // invoke: QMargins marginsPixels(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QPageLayout13marginsPixelsEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QMargins{}) // "QMargins"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "marginsPixels", args)
  }

  return
}

// mode()
func (this *QPageLayout) Mode(args ...interface{}) () {
  // mode()
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
    // invoke: _ZNK11QPageLayout4modeEv
    // invoke: QPageLayout::Mode mode()
    C.C_ZNK11QPageLayout4modeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPageLayout", "mode", args)
  }

  return
}

// setMargins(const class QMarginsF &)
func (this *QPageLayout) Setmargins(args ...interface{}) (ret interface{}) {
  // setMargins(const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QMarginsF{}) // "const QMarginsF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QPageLayout10setMarginsERK9QMarginsF
    // invoke: bool setMargins(const class QMarginsF &)
    var arg0 = args[0].(*qtcore.QMarginsF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QPageLayout10setMarginsERK9QMarginsF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "setMargins", args)
  }

  return
}

// setBottomMargin(qreal)
func (this *QPageLayout) Setbottommargin(args ...interface{}) (ret interface{}) {
  // setBottomMargin(qreal)
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
    // invoke: _ZN11QPageLayout15setBottomMarginEd
    // invoke: bool setBottomMargin(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QPageLayout15setBottomMarginEd(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPageLayout", "setBottomMargin", args)
  }

  return
}

// <= body block end

