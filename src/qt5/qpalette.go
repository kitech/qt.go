package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtGui/qpalette.h
// dst-file: /src/gui/qpalette.go
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
  // proto:  void QPalette::QPalette(const QColor & windowText, const QColor & window, const QColor & light, const QColor & dark, const QColor & mid, const QColor & text, const QColor & base);
extern void* dector_ZN8QPaletteC1ERK6QColorS2_S2_S2_S2_S2_S2_(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6);
extern void _ZN8QPaletteC1ERK6QColorS2_S2_S2_S2_S2_S2_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6);
  // proto:  void QPalette::~QPalette();
extern void _ZN8QPaletteD0Ev(void* qthis);
  // proto:  const QBrush & QPalette::button();
extern void demth_ZNK8QPalette6buttonEv(void* qthis);
  // proto:  const QBrush & QPalette::foreground();
extern void demth_ZNK8QPalette10foregroundEv(void* qthis);
  // proto:  const QBrush & QPalette::background();
extern void demth_ZNK8QPalette10backgroundEv(void* qthis);
  // proto:  void QPalette::resolve(uint mask);
extern void demth_ZN8QPalette7resolveEj(void* qthis, int32_t arg0);
  // proto:  void QPalette::QPalette();
extern void* dector_ZN8QPaletteC1Ev();
extern void _ZN8QPaletteC1Ev(void* qthis);
  // proto:  void QPalette::QPalette(const QColor & button);
extern void* dector_ZN8QPaletteC1ERK6QColor(void* arg0);
extern void _ZN8QPaletteC1ERK6QColor(void* qthis, void* arg0);
  // proto:  bool QPalette::isCopyOf(const QPalette & p);
extern void _ZNK8QPalette8isCopyOfERKS_(void* qthis, void* arg0);
  // proto:  void QPalette::swap(QPalette & other);
extern void demth_ZN8QPalette4swapERS_(void* qthis, void* arg0);
  // proto:  uint QPalette::resolve();
extern void demth_ZNK8QPalette7resolveEv(void* qthis);
  // proto:  const QBrush & QPalette::window();
extern void demth_ZNK8QPalette6windowEv(void* qthis);
  // proto:  const QBrush & QPalette::highlightedText();
extern void demth_ZNK8QPalette15highlightedTextEv(void* qthis);
  // proto:  void QPalette::QPalette(const QColor & button, const QColor & window);
extern void* dector_ZN8QPaletteC1ERK6QColorS2_(void* arg0, void* arg1);
extern void _ZN8QPaletteC1ERK6QColorS2_(void* qthis, void* arg0, void* arg1);
  // proto:  const QBrush & QPalette::text();
extern void demth_ZNK8QPalette4textEv(void* qthis);
  // proto:  const QBrush & QPalette::light();
extern void demth_ZNK8QPalette5lightEv(void* qthis);
  // proto:  QPalette QPalette::resolve(const QPalette & );
extern void _ZNK8QPalette7resolveERKS_(void* qthis, void* arg0);
  // proto:  const QBrush & QPalette::link();
extern void demth_ZNK8QPalette4linkEv(void* qthis);
  // proto:  qint64 QPalette::cacheKey();
extern void _ZNK8QPalette8cacheKeyEv(void* qthis);
  // proto:  const QBrush & QPalette::base();
extern void demth_ZNK8QPalette4baseEv(void* qthis);
  // proto:  const QBrush & QPalette::dark();
extern void demth_ZNK8QPalette4darkEv(void* qthis);
  // proto:  const QBrush & QPalette::highlight();
extern void demth_ZNK8QPalette9highlightEv(void* qthis);
  // proto:  const QBrush & QPalette::mid();
extern void demth_ZNK8QPalette3midEv(void* qthis);
  // proto:  void QPalette::QPalette(const QPalette & palette);
extern void* dector_ZN8QPaletteC1ERKS_(void* arg0);
extern void _ZN8QPaletteC1ERKS_(void* qthis, void* arg0);
  // proto:  const QBrush & QPalette::shadow();
extern void demth_ZNK8QPalette6shadowEv(void* qthis);
  // proto:  const QBrush & QPalette::buttonText();
extern void demth_ZNK8QPalette10buttonTextEv(void* qthis);
  // proto:  const QBrush & QPalette::toolTipBase();
extern void demth_ZNK8QPalette11toolTipBaseEv(void* qthis);
  // proto:  const QBrush & QPalette::midlight();
extern void demth_ZNK8QPalette8midlightEv(void* qthis);
  // proto:  const QBrush & QPalette::brightText();
extern void demth_ZNK8QPalette10brightTextEv(void* qthis);
  // proto:  const QBrush & QPalette::linkVisited();
extern void demth_ZNK8QPalette11linkVisitedEv(void* qthis);
  // proto:  const QBrush & QPalette::alternateBase();
extern void demth_ZNK8QPalette13alternateBaseEv(void* qthis);
  // proto:  void QPalette::QPalette(const QBrush & windowText, const QBrush & button, const QBrush & light, const QBrush & dark, const QBrush & mid, const QBrush & text, const QBrush & bright_text, const QBrush & base, const QBrush & window);
extern void* dector_ZN8QPaletteC1ERK6QBrushS2_S2_S2_S2_S2_S2_S2_S2_(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8);
extern void _ZN8QPaletteC1ERK6QBrushS2_S2_S2_S2_S2_S2_S2_S2_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8);
  // proto:  const QBrush & QPalette::windowText();
extern void demth_ZNK8QPalette10windowTextEv(void* qthis);
  // proto:  const QBrush & QPalette::toolTipText();
extern void demth_ZNK8QPalette11toolTipTextEv(void* qthis);
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

// class sizeof(QPalette)=16
type QPalette struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QPalette::QPalette(const QColor & windowText, const QColor & window, const QColor & light, const QColor & dark, const QColor & mid, const QColor & text, const QColor & base);
func NewQPalette(args ...interface{}) QPalette {
  return QPalette{}
}

  // proto:  void QPalette::~QPalette();
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

  // proto:  const QBrush & QPalette::button();
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
    // invoke: const QBrush & button()
    C.demth_ZNK8QPalette6buttonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "button", args)
  }

}

  // proto:  const QBrush & QPalette::foreground();
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
    // invoke: const QBrush & foreground()
    C.demth_ZNK8QPalette10foregroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "foreground", args)
  }

}

  // proto:  const QBrush & QPalette::background();
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
    // invoke: const QBrush & background()
    C.demth_ZNK8QPalette10backgroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "background", args)
  }

}

  // proto:  void QPalette::resolve(uint mask);
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
    // invoke: void resolve(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPalette7resolveEj(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK8QPalette7resolveEv
    // invoke: uint resolve()
    C.demth_ZNK8QPalette7resolveEv(this.qclsinst)
  case 2:
    // invoke: _ZNK8QPalette7resolveERKS_
    // invoke: QPalette resolve(const class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QPalette7resolveERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPalette", "resolve", args)
  }

}

  // proto:  bool QPalette::isCopyOf(const QPalette & p);
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
    // invoke: bool isCopyOf(const class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QPalette8isCopyOfERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPalette", "isCopyOf", args)
  }

}

  // proto:  void QPalette::swap(QPalette & other);
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
    // invoke: void swap(class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPalette4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPalette", "swap", args)
  }

}

  // proto:  const QBrush & QPalette::window();
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
    // invoke: const QBrush & window()
    C.demth_ZNK8QPalette6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "window", args)
  }

}

  // proto:  const QBrush & QPalette::highlightedText();
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
    // invoke: const QBrush & highlightedText()
    C.demth_ZNK8QPalette15highlightedTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "highlightedText", args)
  }

}

  // proto:  const QBrush & QPalette::text();
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
    // invoke: const QBrush & text()
    C.demth_ZNK8QPalette4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "text", args)
  }

}

  // proto:  const QBrush & QPalette::light();
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
    // invoke: const QBrush & light()
    C.demth_ZNK8QPalette5lightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "light", args)
  }

}

  // proto:  const QBrush & QPalette::link();
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
    // invoke: const QBrush & link()
    C.demth_ZNK8QPalette4linkEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "link", args)
  }

}

  // proto:  qint64 QPalette::cacheKey();
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
    // invoke: qint64 cacheKey()
    C._ZNK8QPalette8cacheKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "cacheKey", args)
  }

}

  // proto:  const QBrush & QPalette::base();
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
    // invoke: const QBrush & base()
    C.demth_ZNK8QPalette4baseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "base", args)
  }

}

  // proto:  const QBrush & QPalette::dark();
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
    // invoke: const QBrush & dark()
    C.demth_ZNK8QPalette4darkEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "dark", args)
  }

}

  // proto:  const QBrush & QPalette::highlight();
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
    // invoke: const QBrush & highlight()
    C.demth_ZNK8QPalette9highlightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "highlight", args)
  }

}

  // proto:  const QBrush & QPalette::mid();
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
    // invoke: const QBrush & mid()
    C.demth_ZNK8QPalette3midEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "mid", args)
  }

}

  // proto:  const QBrush & QPalette::shadow();
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
    // invoke: const QBrush & shadow()
    C.demth_ZNK8QPalette6shadowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "shadow", args)
  }

}

  // proto:  const QBrush & QPalette::buttonText();
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
    // invoke: const QBrush & buttonText()
    C.demth_ZNK8QPalette10buttonTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "buttonText", args)
  }

}

  // proto:  const QBrush & QPalette::toolTipBase();
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
    // invoke: const QBrush & toolTipBase()
    C.demth_ZNK8QPalette11toolTipBaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "toolTipBase", args)
  }

}

  // proto:  const QBrush & QPalette::midlight();
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
    // invoke: const QBrush & midlight()
    C.demth_ZNK8QPalette8midlightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "midlight", args)
  }

}

  // proto:  const QBrush & QPalette::brightText();
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
    // invoke: const QBrush & brightText()
    C.demth_ZNK8QPalette10brightTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "brightText", args)
  }

}

  // proto:  const QBrush & QPalette::linkVisited();
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
    // invoke: const QBrush & linkVisited()
    C.demth_ZNK8QPalette11linkVisitedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "linkVisited", args)
  }

}

  // proto:  const QBrush & QPalette::alternateBase();
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
    // invoke: const QBrush & alternateBase()
    C.demth_ZNK8QPalette13alternateBaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "alternateBase", args)
  }

}

  // proto:  const QBrush & QPalette::windowText();
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
    // invoke: const QBrush & windowText()
    C.demth_ZNK8QPalette10windowTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "windowText", args)
  }

}

  // proto:  const QBrush & QPalette::toolTipText();
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
    // invoke: const QBrush & toolTipText()
    C.demth_ZNK8QPalette11toolTipTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "toolTipText", args)
  }

}

// <= body block end

