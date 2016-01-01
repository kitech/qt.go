package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qglyphrun.h
// dst-file: /src/gui/qglyphrun.go
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
// class sizeof(QGlyphRun)=1
type QGlyphRun struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QGlyphRun) FreeQGlyphRun(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGlyphRun", "~QGlyphRun", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "setBoundingRect", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "overline", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRawData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "setOverline", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "swap", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "setUnderline", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "positions", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "clear", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "strikeOut", args)
 }

}


func NewQGlyphRun(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "rawFont", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRawFont", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "isRightToLeft", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "glyphIndexes", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "boundingRect", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRightToLeft", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "underline", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "setStrikeOut", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QGlyphRun", "isEmpty", args)
 }

}

// <= body block end

