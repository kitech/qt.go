package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qpagelayout.h
// dst-file: /src/gui/qpagelayout.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QPageLayout)=1
type QPageLayout struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "setRightMargin", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "swap", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "marginsPoints", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "marginsPixels", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "isValid", args)
  }

}


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
  case 1:
    // invoke: _ZNK11QPageLayout8fullRectENS_4UnitE
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRect", args)
  }

}


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
  case 1:
    // invoke: _ZNK11QPageLayout9paintRectENS_4UnitE
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRect", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "setMinimumMargins", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "setLeftMargin", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "setBottomMargin", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRectPoints", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "minimumMargins", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "pageSize", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "setTopMargin", args)
  }

}


func NewQPageLayout(args ...interface{}) QPageLayout {
  return QPageLayout{}
}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "fullRectPixels", args)
  }

}


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
    // invoke: _ZNK11QPageLayout7marginsENS_4UnitE
  case 1:
    // invoke: _ZNK11QPageLayout7marginsEv
  default:
    qtrt.ErrorResolve("QPageLayout", "margins", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRectPoints", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "paintRectPixels", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "setPageSize", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "maximumMargins", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "isEquivalentTo", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QPageLayout", "setMargins", args)
  }

}

// <= body block end

