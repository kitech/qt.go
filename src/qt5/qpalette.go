package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern bool C_ZNK8QPalette8isCopyOfERKS_(void* qthis, void* arg0); // 4
  // proto:  const QBrush & QPalette::highlightedText();
extern void* C_ZNK8QPalette15highlightedTextEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::mid();
extern void* C_ZNK8QPalette3midEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::windowText();
extern void* C_ZNK8QPalette10windowTextEv(void* qthis); // 2
  // proto:  void QPalette::QPalette(const QColor & windowText, const QColor & window, const QColor & light, const QColor & dark, const QColor & mid, const QColor & text, const QColor & base);
extern void* C_ZN8QPaletteC2ERK6QColorS2_S2_S2_S2_S2_S2_(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6); // 3
  // proto:  void QPalette::QPalette(const QColor & button);
extern void* C_ZN8QPaletteC2ERK6QColor(void* arg0); // 3
  // proto:  void QPalette::QPalette(const QPalette & palette);
extern void* C_ZN8QPaletteC2ERKS_(void* arg0); // 3
  // proto:  void QPalette::QPalette(const QColor & button, const QColor & window);
extern void* C_ZN8QPaletteC2ERK6QColorS2_(void* arg0, void* arg1); // 3
  // proto:  void QPalette::QPalette(const QBrush & windowText, const QBrush & button, const QBrush & light, const QBrush & dark, const QBrush & mid, const QBrush & text, const QBrush & bright_text, const QBrush & base, const QBrush & window);
extern void* C_ZN8QPaletteC2ERK6QBrushS2_S2_S2_S2_S2_S2_S2_S2_(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8); // 3
  // proto:  void QPalette::QPalette();
extern void* C_ZN8QPaletteC2Ev(); // 3
  // proto:  const QBrush & QPalette::toolTipText();
extern void* C_ZNK8QPalette11toolTipTextEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::midlight();
extern void* C_ZNK8QPalette8midlightEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::window();
extern void* C_ZNK8QPalette6windowEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::brightText();
extern void* C_ZNK8QPalette10brightTextEv(void* qthis); // 2
  // proto:  void QPalette::swap(QPalette & other);
extern void C_ZN8QPalette4swapERS_(void* qthis, void* arg0); // 2
  // proto:  const QBrush & QPalette::text();
extern void* C_ZNK8QPalette4textEv(void* qthis); // 2
  // proto:  QPalette::ColorGroup QPalette::currentColorGroup();
extern void C_ZNK8QPalette17currentColorGroupEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::foreground();
extern void* C_ZNK8QPalette10foregroundEv(void* qthis); // 2
  // proto:  qint64 QPalette::cacheKey();
extern int64_t C_ZNK8QPalette8cacheKeyEv(void* qthis); // 4
  // proto:  const QBrush & QPalette::dark();
extern void* C_ZNK8QPalette4darkEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::base();
extern void* C_ZNK8QPalette4baseEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::link();
extern void* C_ZNK8QPalette4linkEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::background();
extern void* C_ZNK8QPalette10backgroundEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::shadow();
extern void* C_ZNK8QPalette6shadowEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::toolTipBase();
extern void* C_ZNK8QPalette11toolTipBaseEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::linkVisited();
extern void* C_ZNK8QPalette11linkVisitedEv(void* qthis); // 2
  // proto:  void QPalette::resolve(uint mask);
extern void C_ZN8QPalette7resolveEj(void* qthis, int32_t arg0); // 2
  // proto:  QPalette QPalette::resolve(const QPalette & );
extern void* C_ZNK8QPalette7resolveERKS_(void* qthis, void* arg0); // 4
  // proto:  uint QPalette::resolve();
extern int32_t C_ZNK8QPalette7resolveEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::light();
extern void* C_ZNK8QPalette5lightEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::button();
extern void* C_ZNK8QPalette6buttonEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::buttonText();
extern void* C_ZNK8QPalette10buttonTextEv(void* qthis); // 2
  // proto:  void QPalette::~QPalette();
extern void C_ZN8QPaletteD2Ev(void* qthis); // 4
  // proto:  const QBrush & QPalette::alternateBase();
extern void* C_ZNK8QPalette13alternateBaseEv(void* qthis); // 2
  // proto:  const QBrush & QPalette::highlight();
extern void* C_ZNK8QPalette9highlightEv(void* qthis); // 2
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
func (this *QPalette) Iscopyof(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette8isCopyOfERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "isCopyOf", args)
  }

  return
}

// highlightedText()
func (this *QPalette) Highlightedtext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette15highlightedTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "highlightedText", args)
  }

  return
}

// mid()
func (this *QPalette) Mid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette3midEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "mid", args)
  }

  return
}

// windowText()
func (this *QPalette) Windowtext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette10windowTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "windowText", args)
  }

  return
}

// QPalette(const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &)
func NewQPalette(args ...interface{}) *QPalette {
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
    qthis = C.C_ZN8QPaletteC2ERK6QColorS2_S2_S2_S2_S2_S2_(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
    return &QPalette{qclsinst:qthis}
  case 1:
    // invoke: _ZN8QPaletteC1ERK6QColor
    // invoke: void QPalette(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPaletteC2ERK6QColor(arg0)
    return &QPalette{qclsinst:qthis}
  case 2:
    // invoke: _ZN8QPaletteC1ERKS_
    // invoke: void QPalette(const class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPaletteC2ERKS_(arg0)
    return &QPalette{qclsinst:qthis}
  case 3:
    // invoke: _ZN8QPaletteC1ERK6QColorS2_
    // invoke: void QPalette(const class QColor &, const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPaletteC2ERK6QColorS2_(arg0, arg1)
    return &QPalette{qclsinst:qthis}
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
    qthis = C.C_ZN8QPaletteC2ERK6QBrushS2_S2_S2_S2_S2_S2_S2_S2_(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return &QPalette{qclsinst:qthis}
  case 5:
    // invoke: _ZN8QPaletteC1Ev
    // invoke: void QPalette()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPaletteC2Ev()
    return &QPalette{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPalette", "QPalette", args)
  }

  return nil // QPalette{qclsinst:qthis}
}

// toolTipText()
func (this *QPalette) Tooltiptext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette11toolTipTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "toolTipText", args)
  }

  return
}

// midlight()
func (this *QPalette) Midlight(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette8midlightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "midlight", args)
  }

  return
}

// window()
func (this *QPalette) Window(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette6windowEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "window", args)
  }

  return
}

// brightText()
func (this *QPalette) Brighttext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette10brightTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "brightText", args)
  }

  return
}

// swap(class QPalette &)
func (this *QPalette) Swap(args ...interface{}) () {
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
    C.C_ZN8QPalette4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPalette", "swap", args)
  }

  return
}

// text()
func (this *QPalette) Text(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette4textEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "text", args)
  }

  return
}

// currentColorGroup()
func (this *QPalette) Currentcolorgroup(args ...interface{}) () {
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
    C.C_ZNK8QPalette17currentColorGroupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "currentColorGroup", args)
  }

  return
}

// foreground()
func (this *QPalette) Foreground(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette10foregroundEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "foreground", args)
  }

  return
}

// cacheKey()
func (this *QPalette) Cachekey(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette8cacheKeyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "cacheKey", args)
  }

  return
}

// dark()
func (this *QPalette) Dark(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette4darkEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "dark", args)
  }

  return
}

// base()
func (this *QPalette) Base(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette4baseEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "base", args)
  }

  return
}

// link()
func (this *QPalette) Link(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette4linkEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "link", args)
  }

  return
}

// background()
func (this *QPalette) Background(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette10backgroundEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "background", args)
  }

  return
}

// shadow()
func (this *QPalette) Shadow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette6shadowEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "shadow", args)
  }

  return
}

// toolTipBase()
func (this *QPalette) Tooltipbase(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette11toolTipBaseEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "toolTipBase", args)
  }

  return
}

// linkVisited()
func (this *QPalette) Linkvisited(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette11linkVisitedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "linkVisited", args)
  }

  return
}

// resolve(uint)
func (this *QPalette) Resolve(args ...interface{}) () {
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
    C.C_ZN8QPalette7resolveEj(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK8QPalette7resolveERKS_
    // invoke: QPalette resolve(const class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QPalette7resolveERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
  case 2:
    // invoke: _ZNK8QPalette7resolveEv
    // invoke: uint resolve()
    var ret0 = C.C_ZNK8QPalette7resolveEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
  default:
    qtrt.ErrorResolve("QPalette", "resolve", args)
  }

  return
}

// light()
func (this *QPalette) Light(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette5lightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "light", args)
  }

  return
}

// button()
func (this *QPalette) Button(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette6buttonEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "button", args)
  }

  return
}

// buttonText()
func (this *QPalette) Buttontext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette10buttonTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "buttonText", args)
  }

  return
}

// ~QPalette()
func (this *QPalette) Freeqpalette(args ...interface{}) () {
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
    C.C_ZN8QPaletteD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPalette", "~QPalette", args)
  }

  return
}

// alternateBase()
func (this *QPalette) Alternatebase(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette13alternateBaseEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "alternateBase", args)
  }

  return
}

// highlight()
func (this *QPalette) Highlight(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QPalette9highlightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "const QBrush &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPalette", "highlight", args)
  }

  return
}

// <= body block end

