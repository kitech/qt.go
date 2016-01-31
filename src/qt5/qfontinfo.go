package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtGui/qfontinfo.h
// dst-file: /src/gui/qfontinfo.go
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
  // proto:  QFont::StyleHint QFontInfo::styleHint();
extern void C_ZNK9QFontInfo9styleHintEv(void* qthis); // 4
  // proto:  QFont::Style QFontInfo::style();
extern void C_ZNK9QFontInfo5styleEv(void* qthis); // 4
  // proto:  bool QFontInfo::strikeOut();
extern void C_ZNK9QFontInfo9strikeOutEv(void* qthis); // 4
  // proto:  int QFontInfo::pointSize();
extern void C_ZNK9QFontInfo9pointSizeEv(void* qthis); // 4
  // proto:  bool QFontInfo::bold();
extern void C_ZNK9QFontInfo4boldEv(void* qthis); // 2
  // proto:  QString QFontInfo::family();
extern void C_ZNK9QFontInfo6familyEv(void* qthis); // 4
  // proto:  int QFontInfo::weight();
extern void C_ZNK9QFontInfo6weightEv(void* qthis); // 4
  // proto:  qreal QFontInfo::pointSizeF();
extern void C_ZNK9QFontInfo10pointSizeFEv(void* qthis); // 4
  // proto:  bool QFontInfo::exactMatch();
extern void C_ZNK9QFontInfo10exactMatchEv(void* qthis); // 4
  // proto:  bool QFontInfo::underline();
extern void C_ZNK9QFontInfo9underlineEv(void* qthis); // 4
  // proto:  bool QFontInfo::overline();
extern void C_ZNK9QFontInfo8overlineEv(void* qthis); // 4
  // proto:  bool QFontInfo::fixedPitch();
extern void C_ZNK9QFontInfo10fixedPitchEv(void* qthis); // 4
  // proto:  void QFontInfo::swap(QFontInfo & other);
extern void C_ZN9QFontInfo4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QFontInfo::italic();
extern void C_ZNK9QFontInfo6italicEv(void* qthis); // 4
  // proto:  bool QFontInfo::rawMode();
extern void C_ZNK9QFontInfo7rawModeEv(void* qthis); // 4
  // proto:  int QFontInfo::pixelSize();
extern void C_ZNK9QFontInfo9pixelSizeEv(void* qthis); // 4
  // proto:  QString QFontInfo::styleName();
extern void C_ZNK9QFontInfo9styleNameEv(void* qthis); // 4
  // proto:  void QFontInfo::~QFontInfo();
extern void C_ZN9QFontInfoD2Ev(void* qthis); // 4
  // proto:  void QFontInfo::QFontInfo(const QFontInfo & );
extern void C_ZN9QFontInfoC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QFontInfo::QFontInfo(const QFont & );
extern void C_ZN9QFontInfoC2ERK5QFont(void* qthis, void* arg0); // 3
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

// class sizeof(QFontInfo)=1
type QFontInfo struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// styleHint()
func (this *QFontInfo) styleHint(args ...interface{}) () {
  // styleHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9styleHintEv
    // invoke: QFont::StyleHint styleHint()
    C.C_ZNK9QFontInfo9styleHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "styleHint", args)
  }

}

// style()
func (this *QFontInfo) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo5styleEv
    // invoke: QFont::Style style()
    C.C_ZNK9QFontInfo5styleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "style", args)
  }

}

// strikeOut()
func (this *QFontInfo) strikeOut(args ...interface{}) () {
  // strikeOut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9strikeOutEv
    // invoke: bool strikeOut()
    C.C_ZNK9QFontInfo9strikeOutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "strikeOut", args)
  }

}

// pointSize()
func (this *QFontInfo) pointSize(args ...interface{}) () {
  // pointSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9pointSizeEv
    // invoke: int pointSize()
    C.C_ZNK9QFontInfo9pointSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "pointSize", args)
  }

}

// bold()
func (this *QFontInfo) bold(args ...interface{}) () {
  // bold()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo4boldEv
    // invoke: bool bold()
    C.C_ZNK9QFontInfo4boldEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "bold", args)
  }

}

// family()
func (this *QFontInfo) family(args ...interface{}) () {
  // family()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo6familyEv
    // invoke: QString family()
    C.C_ZNK9QFontInfo6familyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "family", args)
  }

}

// weight()
func (this *QFontInfo) weight(args ...interface{}) () {
  // weight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo6weightEv
    // invoke: int weight()
    C.C_ZNK9QFontInfo6weightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "weight", args)
  }

}

// pointSizeF()
func (this *QFontInfo) pointSizeF(args ...interface{}) () {
  // pointSizeF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo10pointSizeFEv
    // invoke: qreal pointSizeF()
    C.C_ZNK9QFontInfo10pointSizeFEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "pointSizeF", args)
  }

}

// exactMatch()
func (this *QFontInfo) exactMatch(args ...interface{}) () {
  // exactMatch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo10exactMatchEv
    // invoke: bool exactMatch()
    C.C_ZNK9QFontInfo10exactMatchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "exactMatch", args)
  }

}

// underline()
func (this *QFontInfo) underline(args ...interface{}) () {
  // underline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9underlineEv
    // invoke: bool underline()
    C.C_ZNK9QFontInfo9underlineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "underline", args)
  }

}

// overline()
func (this *QFontInfo) overline(args ...interface{}) () {
  // overline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo8overlineEv
    // invoke: bool overline()
    C.C_ZNK9QFontInfo8overlineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "overline", args)
  }

}

// fixedPitch()
func (this *QFontInfo) fixedPitch(args ...interface{}) () {
  // fixedPitch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo10fixedPitchEv
    // invoke: bool fixedPitch()
    C.C_ZNK9QFontInfo10fixedPitchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "fixedPitch", args)
  }

}

// swap(class QFontInfo &)
func (this *QFontInfo) swap(args ...interface{}) () {
  // swap(class QFontInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontInfo{}) // "QFontInfo &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFontInfo4swapERS_
    // invoke: void swap(class QFontInfo &)
    var arg0 = args[0].(QFontInfo).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QFontInfo4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontInfo", "swap", args)
  }

}

// italic()
func (this *QFontInfo) italic(args ...interface{}) () {
  // italic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo6italicEv
    // invoke: bool italic()
    C.C_ZNK9QFontInfo6italicEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "italic", args)
  }

}

// rawMode()
func (this *QFontInfo) rawMode(args ...interface{}) () {
  // rawMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo7rawModeEv
    // invoke: bool rawMode()
    C.C_ZNK9QFontInfo7rawModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "rawMode", args)
  }

}

// pixelSize()
func (this *QFontInfo) pixelSize(args ...interface{}) () {
  // pixelSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9pixelSizeEv
    // invoke: int pixelSize()
    C.C_ZNK9QFontInfo9pixelSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "pixelSize", args)
  }

}

// styleName()
func (this *QFontInfo) styleName(args ...interface{}) () {
  // styleName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9styleNameEv
    // invoke: QString styleName()
    C.C_ZNK9QFontInfo9styleNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "styleName", args)
  }

}

// ~QFontInfo()
func (this *QFontInfo) FreeQFontInfo(args ...interface{}) () {
  // ~QFontInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFontInfoD0Ev
    // invoke: void ~QFontInfo()
    C.C_ZN9QFontInfoD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "~QFontInfo", args)
  }

}

// QFontInfo(const class QFontInfo &)
func NewQFontInfo(args ...interface{}) QFontInfo {
  // QFontInfo(const class QFontInfo &)
  // QFontInfo(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontInfo{}) // "const QFontInfo &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFontInfoC1ERKS_
    // invoke: void QFontInfo(const class QFontInfo &)
    var arg0 = args[0].(QFontInfo).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QFontInfoC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN9QFontInfoC1ERK5QFont
    // invoke: void QFontInfo(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QFontInfoC2ERK5QFont(qthis, arg0)
  default:
    qtrt.ErrorResolve("QFontInfo", "QFontInfo", args)
  }

  return QFontInfo{}
}

// <= body block end

