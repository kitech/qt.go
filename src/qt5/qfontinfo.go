package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  bool QFontInfo::rawMode();
extern void _ZNK9QFontInfo7rawModeEv(void* qthis);
  // proto:  bool QFontInfo::exactMatch();
extern void _ZNK9QFontInfo10exactMatchEv(void* qthis);
  // proto:  int QFontInfo::pointSize();
extern void _ZNK9QFontInfo9pointSizeEv(void* qthis);
  // proto:  void QFontInfo::QFontInfo(const QFontInfo & );
extern void* dector_ZN9QFontInfoC1ERKS_(void* arg0);
extern void _ZN9QFontInfoC1ERKS_(void* qthis, void* arg0);
  // proto:  QString QFontInfo::family();
extern void _ZNK9QFontInfo6familyEv(void* qthis);
  // proto:  bool QFontInfo::bold();
extern void demth_ZNK9QFontInfo4boldEv(void* qthis);
  // proto:  qreal QFontInfo::pointSizeF();
extern void _ZNK9QFontInfo10pointSizeFEv(void* qthis);
  // proto:  bool QFontInfo::fixedPitch();
extern void _ZNK9QFontInfo10fixedPitchEv(void* qthis);
  // proto:  bool QFontInfo::overline();
extern void _ZNK9QFontInfo8overlineEv(void* qthis);
  // proto:  void QFontInfo::swap(QFontInfo & other);
extern void _ZN9QFontInfo4swapERS_(void* qthis, void* arg0);
  // proto:  void QFontInfo::QFontInfo(const QFont & );
extern void* dector_ZN9QFontInfoC1ERK5QFont(void* arg0);
extern void _ZN9QFontInfoC1ERK5QFont(void* qthis, void* arg0);
  // proto:  int QFontInfo::pixelSize();
extern void _ZNK9QFontInfo9pixelSizeEv(void* qthis);
  // proto:  bool QFontInfo::strikeOut();
extern void _ZNK9QFontInfo9strikeOutEv(void* qthis);
  // proto:  void QFontInfo::~QFontInfo();
extern void _ZN9QFontInfoD0Ev(void* qthis);
  // proto:  bool QFontInfo::italic();
extern void _ZNK9QFontInfo6italicEv(void* qthis);
  // proto:  bool QFontInfo::underline();
extern void _ZNK9QFontInfo9underlineEv(void* qthis);
  // proto:  QString QFontInfo::styleName();
extern void _ZNK9QFontInfo9styleNameEv(void* qthis);
  // proto:  int QFontInfo::weight();
extern void _ZNK9QFontInfo6weightEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QFontInfo::rawMode();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "rawMode", args)
  }

}

  // proto:  bool QFontInfo::exactMatch();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "exactMatch", args)
  }

}

  // proto:  int QFontInfo::pointSize();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "pointSize", args)
  }

}

  // proto:  void QFontInfo::QFontInfo(const QFontInfo & );
func NewQFontInfo(args ...interface{}) QFontInfo {
  return QFontInfo{}
}

  // proto:  QString QFontInfo::family();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "family", args)
  }

}

  // proto:  bool QFontInfo::bold();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "bold", args)
  }

}

  // proto:  qreal QFontInfo::pointSizeF();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "pointSizeF", args)
  }

}

  // proto:  bool QFontInfo::fixedPitch();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "fixedPitch", args)
  }

}

  // proto:  bool QFontInfo::overline();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "overline", args)
  }

}

  // proto:  void QFontInfo::swap(QFontInfo & other);
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
  default:
    qtrt.ErrorResolve("QFontInfo", "swap", args)
  }

}

  // proto:  int QFontInfo::pixelSize();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "pixelSize", args)
  }

}

  // proto:  bool QFontInfo::strikeOut();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "strikeOut", args)
  }

}

  // proto:  void QFontInfo::~QFontInfo();
func (this *QFontInfo) FreeQFontInfo(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontInfo", "~QFontInfo", args)
  }

}

  // proto:  bool QFontInfo::italic();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "italic", args)
  }

}

  // proto:  bool QFontInfo::underline();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "underline", args)
  }

}

  // proto:  QString QFontInfo::styleName();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "styleName", args)
  }

}

  // proto:  int QFontInfo::weight();
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
  default:
    qtrt.ErrorResolve("QFontInfo", "weight", args)
  }

}

// <= body block end

