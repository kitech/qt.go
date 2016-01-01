package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qbackingstore.h
// dst-file: /src/gui/qbackingstore.go
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
// class sizeof(QBackingStore)=1
type QBackingStore struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QBackingStore) paintDevice(args ...interface{}) () {
  // paintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore11paintDeviceEv
  default:
    qtrt.ErrorResolve("QBackingStore", "paintDevice", args)
  }

}


func (this *QBackingStore) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QBackingStore6windowEv
  default:
    qtrt.ErrorResolve("QBackingStore", "window", args)
  }

}


func (this *QBackingStore) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QBackingStore4sizeEv
  default:
    qtrt.ErrorResolve("QBackingStore", "size", args)
  }

}


func (this *QBackingStore) staticContents(args ...interface{}) () {
  // staticContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QBackingStore14staticContentsEv
  default:
    qtrt.ErrorResolve("QBackingStore", "staticContents", args)
  }

}


func (this *QBackingStore) FreeQBackingStore(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QBackingStore", "~QBackingStore", args)
  }

}


func (this *QBackingStore) setStaticContents(args ...interface{}) () {
  // setStaticContents(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore17setStaticContentsERK7QRegion
  default:
    qtrt.ErrorResolve("QBackingStore", "setStaticContents", args)
  }

}


func (this *QBackingStore) resize(args ...interface{}) () {
  // resize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore6resizeERK5QSize
  default:
    qtrt.ErrorResolve("QBackingStore", "resize", args)
  }

}


func (this *QBackingStore) flush(args ...interface{}) () {
  // flush(const class QRegion &, class QWindow *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[0][1] = reflect.TypeOf(QWindow{}) // "QWindow *"
  vtys[0][2] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint
  default:
    qtrt.ErrorResolve("QBackingStore", "flush", args)
  }

}


func (this *QBackingStore) beginPaint(args ...interface{}) () {
  // beginPaint(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore10beginPaintERK7QRegion
  default:
    qtrt.ErrorResolve("QBackingStore", "beginPaint", args)
  }

}


func (this *QBackingStore) hasStaticContents(args ...interface{}) () {
  // hasStaticContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QBackingStore17hasStaticContentsEv
  default:
    qtrt.ErrorResolve("QBackingStore", "hasStaticContents", args)
  }

}


func (this *QBackingStore) endPaint(args ...interface{}) () {
  // endPaint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore8endPaintEv
  default:
    qtrt.ErrorResolve("QBackingStore", "endPaint", args)
  }

}


func (this *QBackingStore) scroll(args ...interface{}) () {
  // scroll(const class QRegion &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QBackingStore6scrollERK7QRegionii
  default:
    qtrt.ErrorResolve("QBackingStore", "scroll", args)
  }

}


func (this *QBackingStore) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QBackingStore6handleEv
  default:
    qtrt.ErrorResolve("QBackingStore", "handle", args)
  }

}


func NewQBackingStore(args ...interface{}) QBackingStore {
  return QBackingStore{}
}

// <= body block end

