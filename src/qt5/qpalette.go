package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qpalette.h
// dst-file: /src/gui/qpalette.go
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
// class sizeof(QPalette)=16
type QPalette struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQPalette(args ...interface{})() {
}


func (this *QPalette) FreeQPalette(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QPalette", "~QPalette", args)
 }

}


func (this *QPalette) button(args ...interface{}) () {
  // button()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette6buttonEv
  default:
    qtrt.ErrorResolve("QPalette", "button", args)
 }

}


func (this *QPalette) foreground(args ...interface{}) () {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette10foregroundEv
  default:
    qtrt.ErrorResolve("QPalette", "foreground", args)
 }

}


func (this *QPalette) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette10backgroundEv
  default:
    qtrt.ErrorResolve("QPalette", "background", args)
 }

}


func (this *QPalette) resolve(args ...interface{}) () {
  // resolve(uint)
  // resolve()
  // resolve(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QPalette7resolveEj
  case 1:
    // invoke: _ZNK8QPalette7resolveEv
  case 2:
    // invoke: _ZNK8QPalette7resolveERKS_
  default:
    qtrt.ErrorResolve("QPalette", "resolve", args)
 }

}


func (this *QPalette) isCopyOf(args ...interface{}) () {
  // isCopyOf(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette8isCopyOfERKS_
  default:
    qtrt.ErrorResolve("QPalette", "isCopyOf", args)
 }

}


func (this *QPalette) swap(args ...interface{}) () {
  // swap(class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QPalette4swapERS_
  default:
    qtrt.ErrorResolve("QPalette", "swap", args)
 }

}


func (this *QPalette) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette6windowEv
  default:
    qtrt.ErrorResolve("QPalette", "window", args)
 }

}


func (this *QPalette) highlightedText(args ...interface{}) () {
  // highlightedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette15highlightedTextEv
  default:
    qtrt.ErrorResolve("QPalette", "highlightedText", args)
 }

}


func (this *QPalette) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette4textEv
  default:
    qtrt.ErrorResolve("QPalette", "text", args)
 }

}


func (this *QPalette) light(args ...interface{}) () {
  // light()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette5lightEv
  default:
    qtrt.ErrorResolve("QPalette", "light", args)
 }

}


func (this *QPalette) link(args ...interface{}) () {
  // link()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette4linkEv
  default:
    qtrt.ErrorResolve("QPalette", "link", args)
 }

}


func (this *QPalette) cacheKey(args ...interface{}) () {
  // cacheKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette8cacheKeyEv
  default:
    qtrt.ErrorResolve("QPalette", "cacheKey", args)
 }

}


func (this *QPalette) base(args ...interface{}) () {
  // base()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette4baseEv
  default:
    qtrt.ErrorResolve("QPalette", "base", args)
 }

}


func (this *QPalette) dark(args ...interface{}) () {
  // dark()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette4darkEv
  default:
    qtrt.ErrorResolve("QPalette", "dark", args)
 }

}


func (this *QPalette) highlight(args ...interface{}) () {
  // highlight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette9highlightEv
  default:
    qtrt.ErrorResolve("QPalette", "highlight", args)
 }

}


func (this *QPalette) mid(args ...interface{}) () {
  // mid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette3midEv
  default:
    qtrt.ErrorResolve("QPalette", "mid", args)
 }

}


func (this *QPalette) shadow(args ...interface{}) () {
  // shadow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette6shadowEv
  default:
    qtrt.ErrorResolve("QPalette", "shadow", args)
 }

}


func (this *QPalette) buttonText(args ...interface{}) () {
  // buttonText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette10buttonTextEv
  default:
    qtrt.ErrorResolve("QPalette", "buttonText", args)
 }

}


func (this *QPalette) toolTipBase(args ...interface{}) () {
  // toolTipBase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette11toolTipBaseEv
  default:
    qtrt.ErrorResolve("QPalette", "toolTipBase", args)
 }

}


func (this *QPalette) midlight(args ...interface{}) () {
  // midlight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette8midlightEv
  default:
    qtrt.ErrorResolve("QPalette", "midlight", args)
 }

}


func (this *QPalette) brightText(args ...interface{}) () {
  // brightText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette10brightTextEv
  default:
    qtrt.ErrorResolve("QPalette", "brightText", args)
 }

}


func (this *QPalette) linkVisited(args ...interface{}) () {
  // linkVisited()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette11linkVisitedEv
  default:
    qtrt.ErrorResolve("QPalette", "linkVisited", args)
 }

}


func (this *QPalette) alternateBase(args ...interface{}) () {
  // alternateBase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette13alternateBaseEv
  default:
    qtrt.ErrorResolve("QPalette", "alternateBase", args)
 }

}


func (this *QPalette) windowText(args ...interface{}) () {
  // windowText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette10windowTextEv
  default:
    qtrt.ErrorResolve("QPalette", "windowText", args)
 }

}


func (this *QPalette) toolTipText(args ...interface{}) () {
  // toolTipText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette11toolTipTextEv
  default:
    qtrt.ErrorResolve("QPalette", "toolTipText", args)
 }

}

// <= body block end

