package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qregion.h
// dst-file: /src/gui/qregion.go
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
// class sizeof(QRegion)=8
type QRegion struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QRegion) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion6isNullEv
  default:
    qtrt.ErrorResolve("QRegion", "isNull", args)
  }

}


func (this *QRegion) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion12boundingRectEv
  default:
    qtrt.ErrorResolve("QRegion", "boundingRect", args)
  }

}


func NewQRegion(args ...interface{}) QRegion {
  return QRegion{}
}


func (this *QRegion) rectCount(args ...interface{}) () {
  // rectCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion9rectCountEv
  default:
    qtrt.ErrorResolve("QRegion", "rectCount", args)
  }

}


func (this *QRegion) translate(args ...interface{}) () {
  // translate(int, int)
  // translate(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegion9translateEii
  case 1:
    // invoke: _ZN7QRegion9translateERK6QPoint
  default:
    qtrt.ErrorResolve("QRegion", "translate", args)
  }

}


func (this *QRegion) united(args ...interface{}) () {
  // united(const class QRegion &)
  // united(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion6unitedERKS_
  case 1:
    // invoke: _ZNK7QRegion6unitedERK5QRect
  default:
    qtrt.ErrorResolve("QRegion", "united", args)
  }

}


func (this *QRegion) translated(args ...interface{}) () {
  // translated(const class QPoint &)
  // translated(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion10translatedERK6QPoint
  case 1:
    // invoke: _ZNK7QRegion10translatedEii
  default:
    qtrt.ErrorResolve("QRegion", "translated", args)
  }

}


func (this *QRegion) swap(args ...interface{}) () {
  // swap(class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegion4swapERS_
  default:
    qtrt.ErrorResolve("QRegion", "swap", args)
  }

}


func (this *QRegion) FreeQRegion(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegion", "~QRegion", args)
  }

}


func (this *QRegion) contains(args ...interface{}) () {
  // contains(const class QRect &)
  // contains(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion8containsERK5QRect
  case 1:
    // invoke: _ZNK7QRegion8containsERK6QPoint
  default:
    qtrt.ErrorResolve("QRegion", "contains", args)
  }

}


func (this *QRegion) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion7isEmptyEv
  default:
    qtrt.ErrorResolve("QRegion", "isEmpty", args)
  }

}


func (this *QRegion) intersected(args ...interface{}) () {
  // intersected(const class QRect &)
  // intersected(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion11intersectedERK5QRect
  case 1:
    // invoke: _ZNK7QRegion11intersectedERKS_
  default:
    qtrt.ErrorResolve("QRegion", "intersected", args)
  }

}


func (this *QRegion) setRects(args ...interface{}) () {
  // setRects(const class QRect *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegion8setRectsEPK5QRecti
  default:
    qtrt.ErrorResolve("QRegion", "setRects", args)
  }

}


func (this *QRegion) rects(args ...interface{}) () {
  // rects()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion5rectsEv
  default:
    qtrt.ErrorResolve("QRegion", "rects", args)
  }

}


func (this *QRegion) subtracted(args ...interface{}) () {
  // subtracted(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion10subtractedERKS_
  default:
    qtrt.ErrorResolve("QRegion", "subtracted", args)
  }

}


func (this *QRegion) intersects(args ...interface{}) () {
  // intersects(const class QRect &)
  // intersects(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion10intersectsERK5QRect
  case 1:
    // invoke: _ZNK7QRegion10intersectsERKS_
  default:
    qtrt.ErrorResolve("QRegion", "intersects", args)
  }

}


func (this *QRegion) xored(args ...interface{}) () {
  // xored(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion5xoredERKS_
  default:
    qtrt.ErrorResolve("QRegion", "xored", args)
  }

}

// <= body block end

