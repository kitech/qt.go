package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qgraphicsanchorlayout.h
// dst-file: /src/widgets/qgraphicsanchorlayout.go
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


func NewQGraphicsAnchorLayout(args ...interface{}) QGraphicsAnchorLayout {
  return QGraphicsAnchorLayout{}
}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "verticalSpacing", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setSpacing", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "count", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "horizontalSpacing", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "invalidate", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "itemAt", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setVerticalSpacing", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setGeometry", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "setHorizontalSpacing", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchorLayout", "removeAt", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "unsetSpacing", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "setSpacing", args)
  }

}


func NewQGraphicsAnchor(args ...interface{}) QGraphicsAnchor {
  return QGraphicsAnchor{}
}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "metaObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QGraphicsAnchor", "spacing", args)
  }

}

// <= body block end

