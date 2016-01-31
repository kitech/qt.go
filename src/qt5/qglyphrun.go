package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QRectF QGlyphRun::boundingRect();
extern void C_ZNK9QGlyphRun12boundingRectEv(void* qthis); // 4
  // proto:  QRawFont QGlyphRun::rawFont();
extern void C_ZNK9QGlyphRun7rawFontEv(void* qthis); // 4
  // proto:  bool QGlyphRun::underline();
extern void C_ZNK9QGlyphRun9underlineEv(void* qthis); // 4
  // proto:  void QGlyphRun::setBoundingRect(const QRectF & boundingRect);
extern void C_ZN9QGlyphRun15setBoundingRectERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGlyphRun::setRawFont(const QRawFont & rawFont);
extern void C_ZN9QGlyphRun10setRawFontERK8QRawFont(void* qthis, void* arg0); // 4
  // proto:  bool QGlyphRun::isEmpty();
extern void C_ZNK9QGlyphRun7isEmptyEv(void* qthis); // 4
  // proto:  void QGlyphRun::swap(QGlyphRun & other);
extern void C_ZN9QGlyphRun4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QGlyphRun::setRawData(const quint32 * glyphIndexArray, const QPointF * glyphPositionArray, int size);
extern void C_ZN9QGlyphRun10setRawDataEPKjPK7QPointFi(void* qthis, int32_t* arg0, void* arg1, int32_t arg2); // 4
  // proto:  bool QGlyphRun::isRightToLeft();
extern void C_ZNK9QGlyphRun13isRightToLeftEv(void* qthis); // 4
  // proto:  bool QGlyphRun::strikeOut();
extern void C_ZNK9QGlyphRun9strikeOutEv(void* qthis); // 4
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
extern void C_ZN9QGlyphRunC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QGlyphRun::QGlyphRun();
extern void C_ZN9QGlyphRunC2Ev(void* qthis); // 3
  // proto:  bool QGlyphRun::overline();
extern void C_ZNK9QGlyphRun8overlineEv(void* qthis); // 4
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
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QGlyphRun)=1
type QGlyphRun struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// boundingRect()
func (this *QGlyphRun) boundingRect(args ...interface{}) () {
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
    C.C_ZNK9QGlyphRun12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "boundingRect", args)
  }

}

// rawFont()
func (this *QGlyphRun) rawFont(args ...interface{}) () {
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
    C.C_ZNK9QGlyphRun7rawFontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "rawFont", args)
  }

}

// underline()
func (this *QGlyphRun) underline(args ...interface{}) () {
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
    C.C_ZNK9QGlyphRun9underlineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "underline", args)
  }

}

// setBoundingRect(const class QRectF &)
func (this *QGlyphRun) setBoundingRect(args ...interface{}) () {
  // setBoundingRect(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun15setBoundingRectERK6QRectF
    // invoke: void setBoundingRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QGlyphRun15setBoundingRectERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setBoundingRect", args)
  }

}

// setRawFont(const class QRawFont &)
func (this *QGlyphRun) setRawFont(args ...interface{}) () {
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
    var arg0 = args[0].(QRawFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QGlyphRun10setRawFontERK8QRawFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRawFont", args)
  }

}

// isEmpty()
func (this *QGlyphRun) isEmpty(args ...interface{}) () {
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
    C.C_ZNK9QGlyphRun7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "isEmpty", args)
  }

}

// swap(class QGlyphRun &)
func (this *QGlyphRun) swap(args ...interface{}) () {
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
    var arg0 = args[0].(QGlyphRun).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QGlyphRun4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "swap", args)
  }

}

// setRawData(const quint32 *, const class QPointF *, int)
func (this *QGlyphRun) setRawData(args ...interface{}) () {
  // setRawData(const quint32 *, const class QPointF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "const quint32 *"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun10setRawDataEPKjPK7QPointFi
    // invoke: void setRawData(const quint32 *, const class QPointF *, int)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN9QGlyphRun10setRawDataEPKjPK7QPointFi(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRawData", args)
  }

}

// isRightToLeft()
func (this *QGlyphRun) isRightToLeft(args ...interface{}) () {
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
    C.C_ZNK9QGlyphRun13isRightToLeftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "isRightToLeft", args)
  }

}

// strikeOut()
func (this *QGlyphRun) strikeOut(args ...interface{}) () {
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
    C.C_ZNK9QGlyphRun9strikeOutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "strikeOut", args)
  }

}

// ~QGlyphRun()
func (this *QGlyphRun) FreeQGlyphRun(args ...interface{}) () {
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
    C.C_ZN9QGlyphRunD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "~QGlyphRun", args)
  }

}

// setOverline(_Bool)
func (this *QGlyphRun) setOverline(args ...interface{}) () {
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
    C.C_ZN9QGlyphRun11setOverlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setOverline", args)
  }

}

// setRightToLeft(_Bool)
func (this *QGlyphRun) setRightToLeft(args ...interface{}) () {
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
    C.C_ZN9QGlyphRun14setRightToLeftEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRightToLeft", args)
  }

}

// setStrikeOut(_Bool)
func (this *QGlyphRun) setStrikeOut(args ...interface{}) () {
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
    C.C_ZN9QGlyphRun12setStrikeOutEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setStrikeOut", args)
  }

}

// positions()
func (this *QGlyphRun) positions(args ...interface{}) () {
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
    C.C_ZNK9QGlyphRun9positionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "positions", args)
  }

}

// clear()
func (this *QGlyphRun) clear(args ...interface{}) () {
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
    C.C_ZN9QGlyphRun5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "clear", args)
  }

}

// QGlyphRun(const class QGlyphRun &)
func NewQGlyphRun(args ...interface{}) QGlyphRun {
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
    var arg0 = args[0].(QGlyphRun).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QGlyphRunC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN9QGlyphRunC1Ev
    // invoke: void QGlyphRun()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QGlyphRunC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QGlyphRun", "QGlyphRun", args)
  }

  return QGlyphRun{}
}

// overline()
func (this *QGlyphRun) overline(args ...interface{}) () {
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
    C.C_ZNK9QGlyphRun8overlineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "overline", args)
  }

}

// flags()
func (this *QGlyphRun) flags(args ...interface{}) () {
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
    C.C_ZNK9QGlyphRun5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "flags", args)
  }

}

// setUnderline(_Bool)
func (this *QGlyphRun) setUnderline(args ...interface{}) () {
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
    C.C_ZN9QGlyphRun12setUnderlineEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGlyphRun", "setUnderline", args)
  }

}

// glyphIndexes()
func (this *QGlyphRun) glyphIndexes(args ...interface{}) () {
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
    C.C_ZNK9QGlyphRun12glyphIndexesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGlyphRun", "glyphIndexes", args)
  }

}

// <= body block end

