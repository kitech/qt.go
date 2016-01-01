package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qheaderview.h
// dst-file: /src/widgets/qheaderview.go
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
// class sizeof(QHeaderView)=1
type QHeaderView struct {
  /*qbase*/ QAbstractItemView;
  qclsinst uint64 /* *mut c_void*/;
//  _sectionHandleDoubleClicked QHeaderView_sectionHandleDoubleClicked_signal;
//  _sectionEntered QHeaderView_sectionEntered_signal;
//  _sortIndicatorChanged QHeaderView_sortIndicatorChanged_signal;
//  _sectionClicked QHeaderView_sectionClicked_signal;
//  _sectionCountChanged QHeaderView_sectionCountChanged_signal;
//  _geometriesChanged QHeaderView_geometriesChanged_signal;
//  _sectionMoved QHeaderView_sectionMoved_signal;
//  _sectionPressed QHeaderView_sectionPressed_signal;
//  _sectionDoubleClicked QHeaderView_sectionDoubleClicked_signal;
//  _sectionResized QHeaderView_sectionResized_signal;
}


func (this *QHeaderView) maximumSectionSize(args ...interface{}) () {
  // maximumSectionSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18maximumSectionSizeEv
  default:
    qtrt.ErrorResolve("QHeaderView", "maximumSectionSize", args)
  }

}


func (this *QHeaderView) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView8sizeHintEv
  default:
    qtrt.ErrorResolve("QHeaderView", "sizeHint", args)
  }

}


func (this *QHeaderView) sectionPosition(args ...interface{}) () {
  // sectionPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView15sectionPositionEi
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionPosition", args)
  }

}


func (this *QHeaderView) sectionSize(args ...interface{}) () {
  // sectionSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView11sectionSizeEi
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionSize", args)
  }

}


func NewQHeaderView(args ...interface{}) QHeaderView {
  return QHeaderView{}
}


func (this *QHeaderView) setStretchLastSection(args ...interface{}) () {
  // setStretchLastSection(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView21setStretchLastSectionEb
  default:
    qtrt.ErrorResolve("QHeaderView", "setStretchLastSection", args)
  }

}


func (this *QHeaderView) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView5resetEv
  default:
    qtrt.ErrorResolve("QHeaderView", "reset", args)
  }

}


func (this *QHeaderView) resetDefaultSectionSize(args ...interface{}) () {
  // resetDefaultSectionSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView23resetDefaultSectionSizeEv
  default:
    qtrt.ErrorResolve("QHeaderView", "resetDefaultSectionSize", args)
  }

}


func (this *QHeaderView) saveState(args ...interface{}) () {
  // saveState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView9saveStateEv
  default:
    qtrt.ErrorResolve("QHeaderView", "saveState", args)
  }

}


func (this *QHeaderView) sectionsClickable(args ...interface{}) () {
  // sectionsClickable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView17sectionsClickableEv
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsClickable", args)
  }

}


func (this *QHeaderView) resizeContentsPrecision(args ...interface{}) () {
  // resizeContentsPrecision()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView23resizeContentsPrecisionEv
  default:
    qtrt.ErrorResolve("QHeaderView", "resizeContentsPrecision", args)
  }

}


func (this *QHeaderView) setOffsetToSectionPosition(args ...interface{}) () {
  // setOffsetToSectionPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView26setOffsetToSectionPositionEi
  default:
    qtrt.ErrorResolve("QHeaderView", "setOffsetToSectionPosition", args)
  }

}


func (this *QHeaderView) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView6lengthEv
  default:
    qtrt.ErrorResolve("QHeaderView", "length", args)
  }

}


func (this *QHeaderView) hideSection(args ...interface{}) () {
  // hideSection(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView11hideSectionEi
  default:
    qtrt.ErrorResolve("QHeaderView", "hideSection", args)
  }

}


func (this *QHeaderView) sortIndicatorSection(args ...interface{}) () {
  // sortIndicatorSection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView20sortIndicatorSectionEv
  default:
    qtrt.ErrorResolve("QHeaderView", "sortIndicatorSection", args)
  }

}


func (this *QHeaderView) cascadingSectionResizes(args ...interface{}) () {
  // cascadingSectionResizes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView23cascadingSectionResizesEv
  default:
    qtrt.ErrorResolve("QHeaderView", "cascadingSectionResizes", args)
  }

}


func (this *QHeaderView) setMinimumSectionSize(args ...interface{}) () {
  // setMinimumSectionSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView21setMinimumSectionSizeEi
  default:
    qtrt.ErrorResolve("QHeaderView", "setMinimumSectionSize", args)
  }

}


func (this *QHeaderView) visualIndexAt(args ...interface{}) () {
  // visualIndexAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView13visualIndexAtEi
  default:
    qtrt.ErrorResolve("QHeaderView", "visualIndexAt", args)
  }

}


func (this *QHeaderView) setOffset(args ...interface{}) () {
  // setOffset(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView9setOffsetEi
  default:
    qtrt.ErrorResolve("QHeaderView", "setOffset", args)
  }

}


func (this *QHeaderView) logicalIndexAt(args ...interface{}) () {
  // logicalIndexAt(const class QPoint &)
  // logicalIndexAt(int, int)
  // logicalIndexAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView14logicalIndexAtERK6QPoint
  case 1:
    // invoke: _ZNK11QHeaderView14logicalIndexAtEii
  case 2:
    // invoke: _ZNK11QHeaderView14logicalIndexAtEi
  default:
    qtrt.ErrorResolve("QHeaderView", "logicalIndexAt", args)
  }

}


func (this *QHeaderView) FreeQHeaderView(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QHeaderView", "~QHeaderView", args)
  }

}


func (this *QHeaderView) sectionViewportPosition(args ...interface{}) () {
  // sectionViewportPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView23sectionViewportPositionEi
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionViewportPosition", args)
  }

}


func (this *QHeaderView) highlightSections(args ...interface{}) () {
  // highlightSections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView17highlightSectionsEv
  default:
    qtrt.ErrorResolve("QHeaderView", "highlightSections", args)
  }

}


func (this *QHeaderView) offset(args ...interface{}) () {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView6offsetEv
  default:
    qtrt.ErrorResolve("QHeaderView", "offset", args)
  }

}


func (this *QHeaderView) setSortIndicatorShown(args ...interface{}) () {
  // setSortIndicatorShown(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView21setSortIndicatorShownEb
  default:
    qtrt.ErrorResolve("QHeaderView", "setSortIndicatorShown", args)
  }

}


func (this *QHeaderView) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView10metaObjectEv
  default:
    qtrt.ErrorResolve("QHeaderView", "metaObject", args)
  }

}


func (this *QHeaderView) showSection(args ...interface{}) () {
  // showSection(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView11showSectionEi
  default:
    qtrt.ErrorResolve("QHeaderView", "showSection", args)
  }

}


func (this *QHeaderView) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView10setVisibleEb
  default:
    qtrt.ErrorResolve("QHeaderView", "setVisible", args)
  }

}


func (this *QHeaderView) hiddenSectionCount(args ...interface{}) () {
  // hiddenSectionCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18hiddenSectionCountEv
  default:
    qtrt.ErrorResolve("QHeaderView", "hiddenSectionCount", args)
  }

}


func (this *QHeaderView) setSectionsClickable(args ...interface{}) () {
  // setSectionsClickable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView20setSectionsClickableEb
  default:
    qtrt.ErrorResolve("QHeaderView", "setSectionsClickable", args)
  }

}


func (this *QHeaderView) setResizeContentsPrecision(args ...interface{}) () {
  // setResizeContentsPrecision(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView26setResizeContentsPrecisionEi
  default:
    qtrt.ErrorResolve("QHeaderView", "setResizeContentsPrecision", args)
  }

}


func (this *QHeaderView) defaultSectionSize(args ...interface{}) () {
  // defaultSectionSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18defaultSectionSizeEv
  default:
    qtrt.ErrorResolve("QHeaderView", "defaultSectionSize", args)
  }

}


func (this *QHeaderView) setOffsetToLastSection(args ...interface{}) () {
  // setOffsetToLastSection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView22setOffsetToLastSectionEv
  default:
    qtrt.ErrorResolve("QHeaderView", "setOffsetToLastSection", args)
  }

}


func (this *QHeaderView) swapSections(args ...interface{}) () {
  // swapSections(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView12swapSectionsEii
  default:
    qtrt.ErrorResolve("QHeaderView", "swapSections", args)
  }

}


func (this *QHeaderView) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView5countEv
  default:
    qtrt.ErrorResolve("QHeaderView", "count", args)
  }

}


func (this *QHeaderView) visualIndex(args ...interface{}) () {
  // visualIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView11visualIndexEi
  default:
    qtrt.ErrorResolve("QHeaderView", "visualIndex", args)
  }

}


func (this *QHeaderView) sectionsMoved(args ...interface{}) () {
  // sectionsMoved()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView13sectionsMovedEv
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsMoved", args)
  }

}


func (this *QHeaderView) stretchSectionCount(args ...interface{}) () {
  // stretchSectionCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView19stretchSectionCountEv
  default:
    qtrt.ErrorResolve("QHeaderView", "stretchSectionCount", args)
  }

}


func (this *QHeaderView) doItemsLayout(args ...interface{}) () {
  // doItemsLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView13doItemsLayoutEv
  default:
    qtrt.ErrorResolve("QHeaderView", "doItemsLayout", args)
  }

}


func (this *QHeaderView) setSectionsMovable(args ...interface{}) () {
  // setSectionsMovable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView18setSectionsMovableEb
  default:
    qtrt.ErrorResolve("QHeaderView", "setSectionsMovable", args)
  }

}


func (this *QHeaderView) sectionsHidden(args ...interface{}) () {
  // sectionsHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView14sectionsHiddenEv
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsHidden", args)
  }

}


func (this *QHeaderView) minimumSectionSize(args ...interface{}) () {
  // minimumSectionSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18minimumSectionSizeEv
  default:
    qtrt.ErrorResolve("QHeaderView", "minimumSectionSize", args)
  }

}


func (this *QHeaderView) setCascadingSectionResizes(args ...interface{}) () {
  // setCascadingSectionResizes(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView26setCascadingSectionResizesEb
  default:
    qtrt.ErrorResolve("QHeaderView", "setCascadingSectionResizes", args)
  }

}


func (this *QHeaderView) setDefaultSectionSize(args ...interface{}) () {
  // setDefaultSectionSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView21setDefaultSectionSizeEi
  default:
    qtrt.ErrorResolve("QHeaderView", "setDefaultSectionSize", args)
  }

}


func (this *QHeaderView) moveSection(args ...interface{}) () {
  // moveSection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView11moveSectionEii
  default:
    qtrt.ErrorResolve("QHeaderView", "moveSection", args)
  }

}


func (this *QHeaderView) stretchLastSection(args ...interface{}) () {
  // stretchLastSection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView18stretchLastSectionEv
  default:
    qtrt.ErrorResolve("QHeaderView", "stretchLastSection", args)
  }

}


func (this *QHeaderView) sectionSizeHint(args ...interface{}) () {
  // sectionSizeHint(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView15sectionSizeHintEi
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionSizeHint", args)
  }

}


func (this *QHeaderView) sectionsMovable(args ...interface{}) () {
  // sectionsMovable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView15sectionsMovableEv
  default:
    qtrt.ErrorResolve("QHeaderView", "sectionsMovable", args)
  }

}


func (this *QHeaderView) isSectionHidden(args ...interface{}) () {
  // isSectionHidden(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView15isSectionHiddenEi
  default:
    qtrt.ErrorResolve("QHeaderView", "isSectionHidden", args)
  }

}


func (this *QHeaderView) logicalIndex(args ...interface{}) () {
  // logicalIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView12logicalIndexEi
  default:
    qtrt.ErrorResolve("QHeaderView", "logicalIndex", args)
  }

}


func (this *QHeaderView) setMaximumSectionSize(args ...interface{}) () {
  // setMaximumSectionSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView21setMaximumSectionSizeEi
  default:
    qtrt.ErrorResolve("QHeaderView", "setMaximumSectionSize", args)
  }

}


func (this *QHeaderView) setHighlightSections(args ...interface{}) () {
  // setHighlightSections(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView20setHighlightSectionsEb
  default:
    qtrt.ErrorResolve("QHeaderView", "setHighlightSections", args)
  }

}


func (this *QHeaderView) setSectionHidden(args ...interface{}) () {
  // setSectionHidden(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView16setSectionHiddenEib
  default:
    qtrt.ErrorResolve("QHeaderView", "setSectionHidden", args)
  }

}


func (this *QHeaderView) resizeSection(args ...interface{}) () {
  // resizeSection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView13resizeSectionEii
  default:
    qtrt.ErrorResolve("QHeaderView", "resizeSection", args)
  }

}


func (this *QHeaderView) restoreState(args ...interface{}) () {
  // restoreState(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView12restoreStateERK10QByteArray
  default:
    qtrt.ErrorResolve("QHeaderView", "restoreState", args)
  }

}


func (this *QHeaderView) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QHeaderView8setModelEP18QAbstractItemModel
  default:
    qtrt.ErrorResolve("QHeaderView", "setModel", args)
  }

}


func (this *QHeaderView) isSortIndicatorShown(args ...interface{}) () {
  // isSortIndicatorShown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QHeaderView20isSortIndicatorShownEv
  default:
    qtrt.ErrorResolve("QHeaderView", "isSortIndicatorShown", args)
  }

}

// <= body block end

