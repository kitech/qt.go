package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qgraphicsgridlayout.h
// dst-file: /src/widgets/qgraphicsgridlayout.go
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
// class sizeof(QGraphicsGridLayout)=1
type QGraphicsGridLayout struct {
  /*qbase*/ QGraphicsLayout;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QGraphicsGridLayout) setRowPreferredHeight(args ...interface{}) () {
  // setRowPreferredHeight(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout21setRowPreferredHeightEid
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowPreferredHeight", args)
 }

}


func (this *QGraphicsGridLayout) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout11columnCountEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnCount", args)
 }

}


func (this *QGraphicsGridLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  // itemAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout6itemAtEi
  case 1:
    // invoke: _ZNK19QGraphicsGridLayout6itemAtEii
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "itemAt", args)
 }

}


func (this *QGraphicsGridLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout5countEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "count", args)
 }

}


func (this *QGraphicsGridLayout) setColumnFixedWidth(args ...interface{}) () {
  // setColumnFixedWidth(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout19setColumnFixedWidthEid
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnFixedWidth", args)
 }

}


func (this *QGraphicsGridLayout) setColumnMaximumWidth(args ...interface{}) () {
  // setColumnMaximumWidth(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout21setColumnMaximumWidthEid
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnMaximumWidth", args)
 }

}


func (this *QGraphicsGridLayout) rowStretchFactor(args ...interface{}) () {
  // rowStretchFactor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout16rowStretchFactorEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowStretchFactor", args)
 }

}


func (this *QGraphicsGridLayout) verticalSpacing(args ...interface{}) () {
  // verticalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout15verticalSpacingEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "verticalSpacing", args)
 }

}


func (this *QGraphicsGridLayout) columnStretchFactor(args ...interface{}) () {
  // columnStretchFactor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout19columnStretchFactorEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnStretchFactor", args)
 }

}


func (this *QGraphicsGridLayout) setRowMaximumHeight(args ...interface{}) () {
  // setRowMaximumHeight(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout19setRowMaximumHeightEid
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowMaximumHeight", args)
 }

}


func (this *QGraphicsGridLayout) removeItem(args ...interface{}) () {
  // removeItem(class QGraphicsLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayoutItem{}) // "QGraphicsLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "removeItem", args)
 }

}


func (this *QGraphicsGridLayout) FreeQGraphicsGridLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "~QGraphicsGridLayout", args)
 }

}


func (this *QGraphicsGridLayout) rowMinimumHeight(args ...interface{}) () {
  // rowMinimumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout16rowMinimumHeightEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowMinimumHeight", args)
 }

}


func (this *QGraphicsGridLayout) rowMaximumHeight(args ...interface{}) () {
  // rowMaximumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout16rowMaximumHeightEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowMaximumHeight", args)
 }

}


func NewQGraphicsGridLayout(args ...interface{})() {
}


func (this *QGraphicsGridLayout) setColumnSpacing(args ...interface{}) () {
  // setColumnSpacing(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout16setColumnSpacingEid
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnSpacing", args)
 }

}


func (this *QGraphicsGridLayout) rowSpacing(args ...interface{}) () {
  // rowSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout10rowSpacingEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowSpacing", args)
 }

}


func (this *QGraphicsGridLayout) columnMaximumWidth(args ...interface{}) () {
  // columnMaximumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout18columnMaximumWidthEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnMaximumWidth", args)
 }

}


func (this *QGraphicsGridLayout) setRowFixedHeight(args ...interface{}) () {
  // setRowFixedHeight(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout17setRowFixedHeightEid
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowFixedHeight", args)
 }

}


func (this *QGraphicsGridLayout) rowPreferredHeight(args ...interface{}) () {
  // rowPreferredHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout18rowPreferredHeightEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowPreferredHeight", args)
 }

}


func (this *QGraphicsGridLayout) setVerticalSpacing(args ...interface{}) () {
  // setVerticalSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout18setVerticalSpacingEd
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setVerticalSpacing", args)
 }

}


func (this *QGraphicsGridLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout11setGeometryERK6QRectF
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setGeometry", args)
 }

}


func (this *QGraphicsGridLayout) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout8rowCountEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "rowCount", args)
 }

}


func (this *QGraphicsGridLayout) setSpacing(args ...interface{}) () {
  // setSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout10setSpacingEd
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setSpacing", args)
 }

}


func (this *QGraphicsGridLayout) setRowStretchFactor(args ...interface{}) () {
  // setRowStretchFactor(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout19setRowStretchFactorEii
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowStretchFactor", args)
 }

}


func (this *QGraphicsGridLayout) columnMinimumWidth(args ...interface{}) () {
  // columnMinimumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout18columnMinimumWidthEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnMinimumWidth", args)
 }

}


func (this *QGraphicsGridLayout) setColumnMinimumWidth(args ...interface{}) () {
  // setColumnMinimumWidth(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout21setColumnMinimumWidthEid
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnMinimumWidth", args)
 }

}


func (this *QGraphicsGridLayout) setRowMinimumHeight(args ...interface{}) () {
  // setRowMinimumHeight(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout19setRowMinimumHeightEid
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowMinimumHeight", args)
 }

}


func (this *QGraphicsGridLayout) setHorizontalSpacing(args ...interface{}) () {
  // setHorizontalSpacing(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout20setHorizontalSpacingEd
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setHorizontalSpacing", args)
 }

}


func (this *QGraphicsGridLayout) horizontalSpacing(args ...interface{}) () {
  // horizontalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout17horizontalSpacingEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "horizontalSpacing", args)
 }

}


func (this *QGraphicsGridLayout) setColumnStretchFactor(args ...interface{}) () {
  // setColumnStretchFactor(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout22setColumnStretchFactorEii
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnStretchFactor", args)
 }

}


func (this *QGraphicsGridLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout10invalidateEv
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "invalidate", args)
 }

}


func (this *QGraphicsGridLayout) columnPreferredWidth(args ...interface{}) () {
  // columnPreferredWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout20columnPreferredWidthEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnPreferredWidth", args)
 }

}


func (this *QGraphicsGridLayout) setColumnPreferredWidth(args ...interface{}) () {
  // setColumnPreferredWidth(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout23setColumnPreferredWidthEid
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setColumnPreferredWidth", args)
 }

}


func (this *QGraphicsGridLayout) columnSpacing(args ...interface{}) () {
  // columnSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsGridLayout13columnSpacingEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "columnSpacing", args)
 }

}


func (this *QGraphicsGridLayout) setRowSpacing(args ...interface{}) () {
  // setRowSpacing(int, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout13setRowSpacingEid
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "setRowSpacing", args)
 }

}


func (this *QGraphicsGridLayout) removeAt(args ...interface{}) () {
  // removeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsGridLayout8removeAtEi
  default:
    qtrt.ErrorResolve("QGraphicsGridLayout", "removeAt", args)
 }

}

// <= body block end

