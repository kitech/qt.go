package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QPalette::isCopyOf(const QPalette & p);
extern void _ZNK8QPalette8isCopyOfERKS_(void* qthis, void* arg0); // 4
  // proto:  const QBrush & QPalette::highlightedText();
extern void _ZNK8QPalette15highlightedTextEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::mid();
extern void _ZNK8QPalette3midEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::windowText();
extern void _ZNK8QPalette10windowTextEv(void* qthis); // 2
  // proto:  void QPalette::QPalette(const QColor & windowText, const QColor & window, const QColor & light, const QColor & dark, const QColor & mid, const QColor & text, const QColor & base);
extern void _ZN8QPaletteC2ERK6QColorS2_S2_S2_S2_S2_S2_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6); // 3
  // proto:  void QPalette::QPalette(const QColor & button);
extern void _ZN8QPaletteC2ERK6QColor(void* qthis, void* arg0); // 3
  // proto:  void QPalette::QPalette(const QPalette & palette);
extern void _ZN8QPaletteC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QPalette::QPalette(const QColor & button, const QColor & window);
extern void _ZN8QPaletteC2ERK6QColorS2_(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QPalette::QPalette(const QBrush & windowText, const QBrush & button, const QBrush & light, const QBrush & dark, const QBrush & mid, const QBrush & text, const QBrush & bright_text, const QBrush & base, const QBrush & window);
extern void _ZN8QPaletteC2ERK6QBrushS2_S2_S2_S2_S2_S2_S2_S2_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8); // 3
  // proto:  void QPalette::QPalette();
extern void _ZN8QPaletteC2Ev(void* qthis); // 3
  // proto:  const QBrush & QPalette::toolTipText();
extern void _ZNK8QPalette11toolTipTextEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::midlight();
extern void _ZNK8QPalette8midlightEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::window();
extern void _ZNK8QPalette6windowEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::brightText();
extern void _ZNK8QPalette10brightTextEv(void* qthis); // 2
  // proto:  void QPalette::swap(QPalette & other);
extern void _ZN8QPalette4swapERS_(void* qthis, void* arg0); // 2
  // proto:  const QBrush & QPalette::text();
extern void _ZNK8QPalette4textEv(void* qthis); // 2
  // proto:  QPalette::ColorGroup QPalette::currentColorGroup();
extern void _ZNK8QPalette17currentColorGroupEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::foreground();
extern void _ZNK8QPalette10foregroundEv(void* qthis); // 2
  // proto:  qint64 QPalette::cacheKey();
extern void _ZNK8QPalette8cacheKeyEv(void* qthis); // 4
  // proto:  const QBrush & QPalette::dark();
extern void _ZNK8QPalette4darkEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::base();
extern void _ZNK8QPalette4baseEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::link();
extern void _ZNK8QPalette4linkEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::background();
extern void _ZNK8QPalette10backgroundEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::shadow();
extern void _ZNK8QPalette6shadowEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::toolTipBase();
extern void _ZNK8QPalette11toolTipBaseEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::linkVisited();
extern void _ZNK8QPalette11linkVisitedEv(void* qthis); // 2
  // proto:  void QPalette::resolve(uint mask);
extern void _ZN8QPalette7resolveEj(void* qthis, int32_t arg0); // 2
  // proto:  QPalette QPalette::resolve(const QPalette & );
extern void _ZNK8QPalette7resolveERKS_(void* qthis, void* arg0); // 4
  // proto:  uint QPalette::resolve();
extern void _ZNK8QPalette7resolveEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::light();
extern void _ZNK8QPalette5lightEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::button();
extern void _ZNK8QPalette6buttonEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::buttonText();
extern void _ZNK8QPalette10buttonTextEv(void* qthis); // 2
  // proto:  void QPalette::~QPalette();
extern void _ZN8QPaletteD2Ev(void* qthis); // 4
  // proto:  const QBrush & QPalette::alternateBase();
extern void _ZNK8QPalette13alternateBaseEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::highlight();
extern void _ZNK8QPalette9highlightEv(void* qthis); // 2
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

// isCopyOf(const class QPalette &)
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

// highlightedText()
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
    C._ZNK8QPalette15highlightedTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "highlightedText", args)
  }

}

// mid()
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
    C._ZNK8QPalette3midEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "mid", args)
  }

}

// windowText()
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
    C._ZNK8QPalette10windowTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "windowText", args)
  }

}

// QPalette(const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &)
func NewQPalette(args ...interface{}) QPalette {
  // QPalette(const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &)
  // QPalette(const class QColor &)
  // QPalette(const class QPalette &)
  // QPalette(const class QColor &, const class QColor &)
  // QPalette(const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &)
  // QPalette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[0][2] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[0][3] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[0][4] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[0][5] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[0][6] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[3][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[4][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[4][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[4][3] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[4][4] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[4][5] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[4][6] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[4][7] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[4][8] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[5] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPaletteC1ERK6QColorS2_S2_S2_S2_S2_S2_
    // invoke: void QPalette(const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QColor).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QColor).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QColor).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QColor).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(QColor).qclsinst
    if false {fmt.Println(arg6)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QPaletteC2ERK6QColorS2_S2_S2_S2_S2_S2_(qthis, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  case 1:
    // invoke: _ZN8QPaletteC1ERK6QColor
    // invoke: void QPalette(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QPaletteC2ERK6QColor(qthis, arg0)
  case 2:
    // invoke: _ZN8QPaletteC1ERKS_
    // invoke: void QPalette(const class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QPaletteC2ERKS_(qthis, arg0)
  case 3:
    // invoke: _ZN8QPaletteC1ERK6QColorS2_
    // invoke: void QPalette(const class QColor &, const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QPaletteC2ERK6QColorS2_(qthis, arg0, arg1)
  case 4:
    // invoke: _ZN8QPaletteC1ERK6QBrushS2_S2_S2_S2_S2_S2_S2_S2_
    // invoke: void QPalette(const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QBrush).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QBrush).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QBrush).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QBrush).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QBrush).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(QBrush).qclsinst
    if false {fmt.Println(arg6)}
    var arg7 = args[7].(QBrush).qclsinst
    if false {fmt.Println(arg7)}
    var arg8 = args[8].(QBrush).qclsinst
    if false {fmt.Println(arg8)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QPaletteC2ERK6QBrushS2_S2_S2_S2_S2_S2_S2_S2_(qthis, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
  case 5:
    // invoke: _ZN8QPaletteC1Ev
    // invoke: void QPalette()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN8QPaletteC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QPalette", "QPalette", args)
  }

  return QPalette{}
}

// toolTipText()
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
    C._ZNK8QPalette11toolTipTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "toolTipText", args)
  }

}

// midlight()
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
    C._ZNK8QPalette8midlightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "midlight", args)
  }

}

// window()
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
    C._ZNK8QPalette6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "window", args)
  }

}

// brightText()
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
    C._ZNK8QPalette10brightTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "brightText", args)
  }

}

// swap(class QPalette &)
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
    C._ZN8QPalette4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPalette", "swap", args)
  }

}

// text()
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
    C._ZNK8QPalette4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "text", args)
  }

}

// currentColorGroup()
func (this *QPalette) currentColorGroup(args ...interface{}) () {
  // currentColorGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPalette17currentColorGroupEv
    // invoke: QPalette::ColorGroup currentColorGroup()
    C._ZNK8QPalette17currentColorGroupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "currentColorGroup", args)
  }

}

// foreground()
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
    C._ZNK8QPalette10foregroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "foreground", args)
  }

}

// cacheKey()
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

// dark()
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
    C._ZNK8QPalette4darkEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "dark", args)
  }

}

// base()
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
    C._ZNK8QPalette4baseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "base", args)
  }

}

// link()
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
    C._ZNK8QPalette4linkEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "link", args)
  }

}

// background()
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
    C._ZNK8QPalette10backgroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "background", args)
  }

}

// shadow()
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
    C._ZNK8QPalette6shadowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "shadow", args)
  }

}

// toolTipBase()
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
    C._ZNK8QPalette11toolTipBaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "toolTipBase", args)
  }

}

// linkVisited()
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
    C._ZNK8QPalette11linkVisitedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "linkVisited", args)
  }

}

// resolve(uint)
func (this *QPalette) resolve(args ...interface{}) () {
  // resolve(uint)
  // resolve(const class QPalette &)
  // resolve()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPalette7resolveEj
    // invoke: void resolve(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN8QPalette7resolveEj(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK8QPalette7resolveERKS_
    // invoke: QPalette resolve(const class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK8QPalette7resolveERKS_(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK8QPalette7resolveEv
    // invoke: uint resolve()
    C._ZNK8QPalette7resolveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "resolve", args)
  }

}

// light()
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
    C._ZNK8QPalette5lightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "light", args)
  }

}

// button()
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
    C._ZNK8QPalette6buttonEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "button", args)
  }

}

// buttonText()
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
    C._ZNK8QPalette10buttonTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "buttonText", args)
  }

}

// ~QPalette()
func (this *QPalette) FreeQPalette(args ...interface{}) () {
  // ~QPalette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPaletteD0Ev
    // invoke: void ~QPalette()
    C._ZN8QPaletteD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "~QPalette", args)
  }

}

// alternateBase()
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
    C._ZNK8QPalette13alternateBaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "alternateBase", args)
  }

}

// highlight()
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
    C._ZNK8QPalette9highlightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "highlight", args)
  }

}

// <= body block end

