package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
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
extern bool C_ZNK9QFontInfo9strikeOutEv(void* qthis); // 4
  // proto:  int QFontInfo::pointSize();
extern int32_t C_ZNK9QFontInfo9pointSizeEv(void* qthis); // 4
  // proto:  bool QFontInfo::bold();
extern bool C_ZNK9QFontInfo4boldEv(void* qthis); // 2
  // proto:  QString QFontInfo::family();
extern void* C_ZNK9QFontInfo6familyEv(void* qthis); // 4
  // proto:  int QFontInfo::weight();
extern int32_t C_ZNK9QFontInfo6weightEv(void* qthis); // 4
  // proto:  qreal QFontInfo::pointSizeF();
extern double C_ZNK9QFontInfo10pointSizeFEv(void* qthis); // 4
  // proto:  bool QFontInfo::exactMatch();
extern bool C_ZNK9QFontInfo10exactMatchEv(void* qthis); // 4
  // proto:  bool QFontInfo::underline();
extern bool C_ZNK9QFontInfo9underlineEv(void* qthis); // 4
  // proto:  bool QFontInfo::overline();
extern bool C_ZNK9QFontInfo8overlineEv(void* qthis); // 4
  // proto:  bool QFontInfo::fixedPitch();
extern bool C_ZNK9QFontInfo10fixedPitchEv(void* qthis); // 4
  // proto:  void QFontInfo::swap(QFontInfo & other);
extern void C_ZN9QFontInfo4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QFontInfo::italic();
extern bool C_ZNK9QFontInfo6italicEv(void* qthis); // 4
  // proto:  bool QFontInfo::rawMode();
extern bool C_ZNK9QFontInfo7rawModeEv(void* qthis); // 4
  // proto:  int QFontInfo::pixelSize();
extern int32_t C_ZNK9QFontInfo9pixelSizeEv(void* qthis); // 4
  // proto:  QString QFontInfo::styleName();
extern void* C_ZNK9QFontInfo9styleNameEv(void* qthis); // 4
  // proto:  void QFontInfo::~QFontInfo();
extern void C_ZN9QFontInfoD2Ev(void* qthis); // 4
  // proto:  void QFontInfo::QFontInfo(const QFontInfo & );
extern void* C_ZN9QFontInfoC2ERKS_(void* arg0); // 3
  // proto:  void QFontInfo::QFontInfo(const QFont & );
extern void* C_ZN9QFontInfoC2ERK5QFont(void* arg0); // 3
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QFontInfo)=1
type QFontInfo struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// styleHint()
func (this *QFontInfo) Stylehint(args ...interface{}) () {
  // styleHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9styleHintEv
    // invoke: QFont::StyleHint styleHint()
    C.C_ZNK9QFontInfo9styleHintEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "styleHint", args)
  }

  return
}

// style()
func (this *QFontInfo) Style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo5styleEv
    // invoke: QFont::Style style()
    C.C_ZNK9QFontInfo5styleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "style", args)
  }

  return
}

// strikeOut()
func (this *QFontInfo) Strikeout(args ...interface{}) (ret interface{}) {
  // strikeOut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9strikeOutEv
    // invoke: bool strikeOut()
    var ret0 = C.C_ZNK9QFontInfo9strikeOutEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "strikeOut", args)
  }

  return
}

// pointSize()
func (this *QFontInfo) Pointsize(args ...interface{}) (ret interface{}) {
  // pointSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9pointSizeEv
    // invoke: int pointSize()
    var ret0 = C.C_ZNK9QFontInfo9pointSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "pointSize", args)
  }

  return
}

// bold()
func (this *QFontInfo) Bold(args ...interface{}) (ret interface{}) {
  // bold()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo4boldEv
    // invoke: bool bold()
    var ret0 = C.C_ZNK9QFontInfo4boldEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "bold", args)
  }

  return
}

// family()
func (this *QFontInfo) Family(args ...interface{}) (ret interface{}) {
  // family()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo6familyEv
    // invoke: QString family()
    var ret0 = C.C_ZNK9QFontInfo6familyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "family", args)
  }

  return
}

// weight()
func (this *QFontInfo) Weight(args ...interface{}) (ret interface{}) {
  // weight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo6weightEv
    // invoke: int weight()
    var ret0 = C.C_ZNK9QFontInfo6weightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "weight", args)
  }

  return
}

// pointSizeF()
func (this *QFontInfo) Pointsizef(args ...interface{}) (ret interface{}) {
  // pointSizeF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo10pointSizeFEv
    // invoke: qreal pointSizeF()
    var ret0 = C.C_ZNK9QFontInfo10pointSizeFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "pointSizeF", args)
  }

  return
}

// exactMatch()
func (this *QFontInfo) Exactmatch(args ...interface{}) (ret interface{}) {
  // exactMatch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo10exactMatchEv
    // invoke: bool exactMatch()
    var ret0 = C.C_ZNK9QFontInfo10exactMatchEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "exactMatch", args)
  }

  return
}

// underline()
func (this *QFontInfo) Underline(args ...interface{}) (ret interface{}) {
  // underline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9underlineEv
    // invoke: bool underline()
    var ret0 = C.C_ZNK9QFontInfo9underlineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "underline", args)
  }

  return
}

// overline()
func (this *QFontInfo) Overline(args ...interface{}) (ret interface{}) {
  // overline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo8overlineEv
    // invoke: bool overline()
    var ret0 = C.C_ZNK9QFontInfo8overlineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "overline", args)
  }

  return
}

// fixedPitch()
func (this *QFontInfo) Fixedpitch(args ...interface{}) (ret interface{}) {
  // fixedPitch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo10fixedPitchEv
    // invoke: bool fixedPitch()
    var ret0 = C.C_ZNK9QFontInfo10fixedPitchEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "fixedPitch", args)
  }

  return
}

// swap(class QFontInfo &)
func (this *QFontInfo) Swap(args ...interface{}) () {
  // swap(class QFontInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontInfo{}) // "QFontInfo &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFontInfo4swapERS_
    // invoke: void swap(class QFontInfo &)
    var arg0 = args[0].(*QFontInfo).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QFontInfo4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontInfo", "swap", args)
  }

  return
}

// italic()
func (this *QFontInfo) Italic(args ...interface{}) (ret interface{}) {
  // italic()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo6italicEv
    // invoke: bool italic()
    var ret0 = C.C_ZNK9QFontInfo6italicEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "italic", args)
  }

  return
}

// rawMode()
func (this *QFontInfo) Rawmode(args ...interface{}) (ret interface{}) {
  // rawMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo7rawModeEv
    // invoke: bool rawMode()
    var ret0 = C.C_ZNK9QFontInfo7rawModeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "rawMode", args)
  }

  return
}

// pixelSize()
func (this *QFontInfo) Pixelsize(args ...interface{}) (ret interface{}) {
  // pixelSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9pixelSizeEv
    // invoke: int pixelSize()
    var ret0 = C.C_ZNK9QFontInfo9pixelSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "pixelSize", args)
  }

  return
}

// styleName()
func (this *QFontInfo) Stylename(args ...interface{}) (ret interface{}) {
  // styleName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QFontInfo9styleNameEv
    // invoke: QString styleName()
    var ret0 = C.C_ZNK9QFontInfo9styleNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontInfo", "styleName", args)
  }

  return
}

// ~QFontInfo()
func (this *QFontInfo) Freeqfontinfo(args ...interface{}) () {
  // ~QFontInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFontInfoD0Ev
    // invoke: void ~QFontInfo()
    C.C_ZN9QFontInfoD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontInfo", "~QFontInfo", args)
  }

  return
}

// QFontInfo(const class QFontInfo &)
func NewQFontInfo(args ...interface{}) *QFontInfo {
  // QFontInfo(const class QFontInfo &)
  // QFontInfo(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontInfo{}) // "const QFontInfo &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QFontInfoC1ERKS_
    // invoke: void QFontInfo(const class QFontInfo &)
    var arg0 = args[0].(*QFontInfo).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFontInfoC2ERKS_(arg0)
    return &QFontInfo{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QFontInfoC1ERK5QFont
    // invoke: void QFontInfo(const class QFont &)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QFontInfoC2ERK5QFont(arg0)
    return &QFontInfo{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFontInfo", "QFontInfo", args)
  }

  return nil // QFontInfo{Qclsinst:qthis}
}

// <= body block end

