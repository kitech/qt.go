package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtGui/qglyphrun.h
// dst-file: /src/gui/qglyphrun.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
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
  // proto:  QRectF QGlyphRun::boundingRect();
extern void* C_ZNK9QGlyphRun12boundingRectEv(void* qthis); // 4
  // proto:  QRawFont QGlyphRun::rawFont();
extern void* C_ZNK9QGlyphRun7rawFontEv(void* qthis); // 4
  // proto:  bool QGlyphRun::underline();
extern bool C_ZNK9QGlyphRun9underlineEv(void* qthis); // 4
  // proto:  void QGlyphRun::setBoundingRect(const QRectF & boundingRect);
extern void C_ZN9QGlyphRun15setBoundingRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGlyphRun::setRawFont(const QRawFont & rawFont);
extern void C_ZN9QGlyphRun10setRawFontERK8QRawFont(void* qthis, void* arg0); // 4
  // proto:  bool QGlyphRun::isEmpty();
extern bool C_ZNK9QGlyphRun7isEmptyEv(void* qthis); // 4
  // proto:  void QGlyphRun::swap(QGlyphRun & other);
extern void C_ZN9QGlyphRun4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QGlyphRun::setRawData(const quint32 * glyphIndexArray, const QPointF * glyphPositionArray, int size);
extern void C_ZN9QGlyphRun10setRawDataEPKjPK7QPointFi(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  bool QGlyphRun::isRightToLeft();
extern bool C_ZNK9QGlyphRun13isRightToLeftEv(void* qthis); // 4
  // proto:  bool QGlyphRun::strikeOut();
extern bool C_ZNK9QGlyphRun9strikeOutEv(void* qthis); // 4
  // proto:  void QGlyphRun::~QGlyphRun();
extern void C_ZN9QGlyphRunD2Ev(void* qthis); // 4
  // proto:  void QGlyphRun::setOverline(bool overline);
extern void C_ZN9QGlyphRun11setOverlineEb(void* qthis, bool arg0); // 4
  // proto:  void QGlyphRun::setRightToLeft(bool on);
extern void C_ZN9QGlyphRun14setRightToLeftEb(void* qthis, bool arg0); // 4
  // proto:  void QGlyphRun::setStrikeOut(bool strikeOut);
extern void C_ZN9QGlyphRun12setStrikeOutEb(void* qthis, bool arg0); // 4
  // proto:  QVector<QPointF> QGlyphRun::positions();
extern void C_ZNK9QGlyphRun9positionsEv(void* qthis); // 4
  // proto:  void QGlyphRun::clear();
extern void C_ZN9QGlyphRun5clearEv(void* qthis); // 4
  // proto:  void QGlyphRun::QGlyphRun(const QGlyphRun & other);
extern void* C_ZN9QGlyphRunC2ERKS_(void* arg0); // 3
  // proto:  void QGlyphRun::QGlyphRun();
extern void* C_ZN9QGlyphRunC2Ev(); // 3
  // proto:  bool QGlyphRun::overline();
extern bool C_ZNK9QGlyphRun8overlineEv(void* qthis); // 4
  // proto:  GlyphRunFlags QGlyphRun::flags();
extern void C_ZNK9QGlyphRun5flagsEv(void* qthis); // 4
  // proto:  void QGlyphRun::setUnderline(bool underline);
extern void C_ZN9QGlyphRun12setUnderlineEb(void* qthis, bool arg0); // 4
  // proto:  QVector<quint32> QGlyphRun::glyphIndexes();
extern void C_ZNK9QGlyphRun12glyphIndexesEv(void* qthis); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QGlyphRun)=1
type QGlyphRun struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// boundingRect()
func (this *QGlyphRun) BoundingRect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK9QGlyphRun12boundingRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGlyphRun", "boundingRect", args)
  }

  return
}

// rawFont()
func (this *QGlyphRun) RawFont(args ...interface{}) (ret interface{}) {
  // rawFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun7rawFontEv
    // invoke: QRawFont rawFont()
    var ret0 = C.C_ZNK9QGlyphRun7rawFontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRawFont{}) // "QRawFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGlyphRun", "rawFont", args)
  }

  return
}

// underline()
func (this *QGlyphRun) Underline(args ...interface{}) (ret interface{}) {
  // underline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun9underlineEv
    // invoke: bool underline()
    var ret0 = C.C_ZNK9QGlyphRun9underlineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGlyphRun", "underline", args)
  }

  return
}

// setBoundingRect(const class QRectF &)
func (this *QGlyphRun) SetBoundingRect(args ...interface{}) () {
  // setBoundingRect(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun15setBoundingRectERK6QRectF
    // invoke: void setBoundingRect(const class QRectF &)
    var arg0 = args[0].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QGlyphRun15setBoundingRectERK6QRectF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setBoundingRect", args)
  }

  return
}

// setRawFont(const class QRawFont &)
func (this *QGlyphRun) SetRawFont(args ...interface{}) () {
  // setRawFont(const class QRawFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRawFont{}) // "const QRawFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun10setRawFontERK8QRawFont
    // invoke: void setRawFont(const class QRawFont &)
    var arg0 = args[0].(*QRawFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QGlyphRun10setRawFontERK8QRawFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRawFont", args)
  }

  return
}

// isEmpty()
func (this *QGlyphRun) IsEmpty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK9QGlyphRun7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGlyphRun", "isEmpty", args)
  }

  return
}

// swap(class QGlyphRun &)
func (this *QGlyphRun) Swap(args ...interface{}) () {
  // swap(class QGlyphRun &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGlyphRun{}) // "QGlyphRun &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun4swapERS_
    // invoke: void swap(class QGlyphRun &)
    var arg0 = args[0].(*QGlyphRun).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QGlyphRun4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "swap", args)
  }

  return
}

// setRawData(const quint32 *, const class QPointF *, int)
func (this *QGlyphRun) SetRawData(args ...interface{}) () {
  // setRawData(const quint32 *, const class QPointF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "const quint32 *"
  vtys[0][1] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun10setRawDataEPKjPK7QPointFi
    // invoke: void setRawData(const quint32 *, const class QPointF *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN9QGlyphRun10setRawDataEPKjPK7QPointFi(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRawData", args)
  }

  return
}

// isRightToLeft()
func (this *QGlyphRun) IsRightToLeft(args ...interface{}) (ret interface{}) {
  // isRightToLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun13isRightToLeftEv
    // invoke: bool isRightToLeft()
    var ret0 = C.C_ZNK9QGlyphRun13isRightToLeftEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGlyphRun", "isRightToLeft", args)
  }

  return
}

// strikeOut()
func (this *QGlyphRun) StrikeOut(args ...interface{}) (ret interface{}) {
  // strikeOut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun9strikeOutEv
    // invoke: bool strikeOut()
    var ret0 = C.C_ZNK9QGlyphRun9strikeOutEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGlyphRun", "strikeOut", args)
  }

  return
}

// ~QGlyphRun()
func (this *QGlyphRun) Free(args ...interface{}) () {
  // ~QGlyphRun()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRunD0Ev
    // invoke: void ~QGlyphRun()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN9QGlyphRunD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QGlyphRun", "~QGlyphRun", args)
  }

  return
}

// setOverline(_Bool)
func (this *QGlyphRun) SetOverline(args ...interface{}) () {
  // setOverline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun11setOverlineEb
    // invoke: void setOverline(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QGlyphRun11setOverlineEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setOverline", args)
  }

  return
}

// setRightToLeft(_Bool)
func (this *QGlyphRun) SetRightToLeft(args ...interface{}) () {
  // setRightToLeft(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun14setRightToLeftEb
    // invoke: void setRightToLeft(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QGlyphRun14setRightToLeftEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRightToLeft", args)
  }

  return
}

// setStrikeOut(_Bool)
func (this *QGlyphRun) SetStrikeOut(args ...interface{}) () {
  // setStrikeOut(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun12setStrikeOutEb
    // invoke: void setStrikeOut(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QGlyphRun12setStrikeOutEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setStrikeOut", args)
  }

  return
}

// positions()
func (this *QGlyphRun) Positions(args ...interface{}) () {
  // positions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun9positionsEv
    // invoke: QVector<QPointF> positions()
    C.C_ZNK9QGlyphRun9positionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "positions", args)
  }

  return
}

// clear()
func (this *QGlyphRun) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun5clearEv
    // invoke: void clear()
    C.C_ZN9QGlyphRun5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "clear", args)
  }

  return
}

// QGlyphRun(const class QGlyphRun &)
func GcfreeQGlyphRun(this *QGlyphRun) {
  qtrt.UniverseFree(this)
}
func NewQGlyphRun(args ...interface{}) *QGlyphRun {
  // QGlyphRun(const class QGlyphRun &)
  // QGlyphRun()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGlyphRun{}) // "const QGlyphRun &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRunC1ERKS_
    // invoke: void QGlyphRun(const class QGlyphRun &)
    var arg0 = args[0].(*QGlyphRun).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QGlyphRunC2ERKS_(arg0)
    this := &QGlyphRun{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQGlyphRun)
    return this
  case 1:
    // invoke: _ZN9QGlyphRunC1Ev
    // invoke: void QGlyphRun()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QGlyphRunC2Ev()
    this := &QGlyphRun{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQGlyphRun)
    return this
  default:
    qtrt.ErrorResolve("QGlyphRun", "QGlyphRun", args)
  }

  return nil // QGlyphRun{Qclsinst:qthis}
}

// overline()
func (this *QGlyphRun) Overline(args ...interface{}) (ret interface{}) {
  // overline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun8overlineEv
    // invoke: bool overline()
    var ret0 = C.C_ZNK9QGlyphRun8overlineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGlyphRun", "overline", args)
  }

  return
}

// flags()
func (this *QGlyphRun) Flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun5flagsEv
    // invoke: GlyphRunFlags flags()
    C.C_ZNK9QGlyphRun5flagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "flags", args)
  }

  return
}

// setUnderline(_Bool)
func (this *QGlyphRun) SetUnderline(args ...interface{}) () {
  // setUnderline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun12setUnderlineEb
    // invoke: void setUnderline(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QGlyphRun12setUnderlineEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setUnderline", args)
  }

  return
}

// glyphIndexes()
func (this *QGlyphRun) GlyphIndexes(args ...interface{}) () {
  // glyphIndexes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun12glyphIndexesEv
    // invoke: QVector<quint32> glyphIndexes()
    C.C_ZNK9QGlyphRun12glyphIndexesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "glyphIndexes", args)
  }

  return
}

// <= body block end

