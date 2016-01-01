package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qtimezone.h
// dst-file: /src/core/qtimezone.go
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
// class sizeof(QTimeZone)=1
type QTimeZone struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTimeZone) availableTimeZoneIds_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeZone", "availableTimeZoneIds", args)
 }

}


func (this *QTimeZone) swap(args ...interface{}) () {
  // swap(class QTimeZone &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTimeZone{}) // "QTimeZone &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZone4swapERS_
  default:
    qtrt.ErrorResolve("QTimeZone", "swap", args)
 }

}


func (this *QTimeZone) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone7isValidEv
  default:
    qtrt.ErrorResolve("QTimeZone", "isValid", args)
 }

}


func (this *QTimeZone) hasDaylightTime(args ...interface{}) () {
  // hasDaylightTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone15hasDaylightTimeEv
  default:
    qtrt.ErrorResolve("QTimeZone", "hasDaylightTime", args)
 }

}


func (this *QTimeZone) utc_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeZone", "utc", args)
 }

}


func NewQTimeZone(args ...interface{})() {
}


func (this *QTimeZone) abbreviation(args ...interface{}) () {
  // abbreviation(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone12abbreviationERK9QDateTime
  default:
    qtrt.ErrorResolve("QTimeZone", "abbreviation", args)
 }

}


func (this *QTimeZone) ianaIdToWindowsId_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeZone", "ianaIdToWindowsId", args)
 }

}


func (this *QTimeZone) systemTimeZoneId_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeZone", "systemTimeZoneId", args)
 }

}


func (this *QTimeZone) isDaylightTime(args ...interface{}) () {
  // isDaylightTime(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone14isDaylightTimeERK9QDateTime
  default:
    qtrt.ErrorResolve("QTimeZone", "isDaylightTime", args)
 }

}


func (this *QTimeZone) isTimeZoneIdAvailable_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeZone", "isTimeZoneIdAvailable", args)
 }

}


func (this *QTimeZone) comment(args ...interface{}) () {
  // comment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone7commentEv
  default:
    qtrt.ErrorResolve("QTimeZone", "comment", args)
 }

}


func (this *QTimeZone) windowsIdToDefaultIanaId_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeZone", "windowsIdToDefaultIanaId", args)
 }

}


func (this *QTimeZone) hasTransitions(args ...interface{}) () {
  // hasTransitions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone14hasTransitionsEv
  default:
    qtrt.ErrorResolve("QTimeZone", "hasTransitions", args)
 }

}


func (this *QTimeZone) daylightTimeOffset(args ...interface{}) () {
  // daylightTimeOffset(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime
  default:
    qtrt.ErrorResolve("QTimeZone", "daylightTimeOffset", args)
 }

}


func (this *QTimeZone) systemTimeZone_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeZone", "systemTimeZone", args)
 }

}


func (this *QTimeZone) FreeQTimeZone(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeZone", "~QTimeZone", args)
 }

}


func (this *QTimeZone) standardTimeOffset(args ...interface{}) () {
  // standardTimeOffset(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone18standardTimeOffsetERK9QDateTime
  default:
    qtrt.ErrorResolve("QTimeZone", "standardTimeOffset", args)
 }

}


func (this *QTimeZone) id(args ...interface{}) () {
  // id()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone2idEv
  default:
    qtrt.ErrorResolve("QTimeZone", "id", args)
 }

}


func (this *QTimeZone) offsetFromUtc(args ...interface{}) () {
  // offsetFromUtc(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone13offsetFromUtcERK9QDateTime
  default:
    qtrt.ErrorResolve("QTimeZone", "offsetFromUtc", args)
 }

}


func (this *QTimeZone) windowsIdToIanaIds_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeZone", "windowsIdToIanaIds", args)
 }

}

// <= body block end

