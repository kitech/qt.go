package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qtimezone.h
// dst-file: /src/core/qtimezone.go
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
  // proto:  QString QTimeZone::comment();
extern void C_ZNK9QTimeZone7commentEv(void* qthis); // 4
  // proto: static QByteArray QTimeZone::systemTimeZoneId();
extern void C_ZN9QTimeZone16systemTimeZoneIdEv(); // 4
  // proto: static bool QTimeZone::isTimeZoneIdAvailable(const QByteArray & ianaId);
extern void C_ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray(void* arg0); // 4
  // proto:  QString QTimeZone::abbreviation(const QDateTime & atDateTime);
extern void C_ZNK9QTimeZone12abbreviationERK9QDateTime(void* qthis, void* arg0); // 4
  // proto: static QByteArray QTimeZone::windowsIdToDefaultIanaId(const QByteArray & windowsId);
extern void C_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray(void* arg0); // 4
  // proto:  void QTimeZone::QTimeZone(const QTimeZone & other);
extern void C_ZN9QTimeZoneC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QTimeZone::QTimeZone(int offsetSeconds);
extern void C_ZN9QTimeZoneC2Ei(void* qthis, int32_t arg0); // 3
  // proto:  void QTimeZone::QTimeZone();
extern void C_ZN9QTimeZoneC2Ev(void* qthis); // 3
  // proto:  void QTimeZone::QTimeZone(const QByteArray & ianaId);
extern void C_ZN9QTimeZoneC2ERK10QByteArray(void* qthis, void* arg0); // 3
  // proto:  OffsetDataList QTimeZone::transitions(const QDateTime & fromDateTime, const QDateTime & toDateTime);
extern void C_ZNK9QTimeZone11transitionsERK9QDateTimeS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QTimeZone::OffsetData QTimeZone::nextTransition(const QDateTime & afterDateTime);
extern void C_ZNK9QTimeZone14nextTransitionERK9QDateTime(void* qthis, void* arg0); // 4
  // proto:  QByteArray QTimeZone::id();
extern void C_ZNK9QTimeZone2idEv(void* qthis); // 4
  // proto:  int QTimeZone::offsetFromUtc(const QDateTime & atDateTime);
extern void C_ZNK9QTimeZone13offsetFromUtcERK9QDateTime(void* qthis, void* arg0); // 4
  // proto:  QTimeZone::OffsetData QTimeZone::previousTransition(const QDateTime & beforeDateTime);
extern void C_ZNK9QTimeZone18previousTransitionERK9QDateTime(void* qthis, void* arg0); // 4
  // proto: static QList<QByteArray> QTimeZone::availableTimeZoneIds();
extern void C_ZN9QTimeZone20availableTimeZoneIdsEv(); // 4
  // proto: static QList<QByteArray> QTimeZone::availableTimeZoneIds(int offsetSeconds);
extern void C_ZN9QTimeZone20availableTimeZoneIdsEi(int32_t arg0); // 4
  // proto:  void QTimeZone::~QTimeZone();
extern void C_ZN9QTimeZoneD2Ev(void* qthis); // 4
  // proto: static QList<QByteArray> QTimeZone::windowsIdToIanaIds(const QByteArray & windowsId);
extern void C_ZN9QTimeZone18windowsIdToIanaIdsERK10QByteArray(void* arg0); // 4
  // proto: static QTimeZone QTimeZone::systemTimeZone();
extern void C_ZN9QTimeZone14systemTimeZoneEv(); // 4
  // proto:  void QTimeZone::swap(QTimeZone & other);
extern void C_ZN9QTimeZone4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QTimeZone::OffsetData QTimeZone::offsetData(const QDateTime & forDateTime);
extern void C_ZNK9QTimeZone10offsetDataERK9QDateTime(void* qthis, void* arg0); // 4
  // proto:  int QTimeZone::standardTimeOffset(const QDateTime & atDateTime);
extern void C_ZNK9QTimeZone18standardTimeOffsetERK9QDateTime(void* qthis, void* arg0); // 4
  // proto:  bool QTimeZone::isValid();
extern void C_ZNK9QTimeZone7isValidEv(void* qthis); // 4
  // proto:  bool QTimeZone::hasDaylightTime();
extern void C_ZNK9QTimeZone15hasDaylightTimeEv(void* qthis); // 4
  // proto:  int QTimeZone::daylightTimeOffset(const QDateTime & atDateTime);
extern void C_ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime(void* qthis, void* arg0); // 4
  // proto: static QTimeZone QTimeZone::utc();
extern void C_ZN9QTimeZone3utcEv(); // 4
  // proto:  bool QTimeZone::hasTransitions();
extern void C_ZNK9QTimeZone14hasTransitionsEv(void* qthis); // 4
  // proto: static QByteArray QTimeZone::ianaIdToWindowsId(const QByteArray & ianaId);
extern void C_ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray(void* arg0); // 4
  // proto:  QLocale::Country QTimeZone::country();
extern void C_ZNK9QTimeZone7countryEv(void* qthis); // 4
  // proto:  bool QTimeZone::isDaylightTime(const QDateTime & atDateTime);
extern void C_ZNK9QTimeZone14isDaylightTimeERK9QDateTime(void* qthis, void* arg0); // 4
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

// class sizeof(QTimeZone)=1
type QTimeZone struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// comment()
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
    // invoke: QString comment()
    C.C_ZNK9QTimeZone7commentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "comment", args)
  }

}

// systemTimeZoneId()
func (this *QTimeZone) systemTimeZoneId_s(args ...interface{}) () {
  // systemTimeZoneId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZone16systemTimeZoneIdEv
    // invoke: QByteArray systemTimeZoneId()
    C.C_ZN9QTimeZone16systemTimeZoneIdEv()
  default:
    qtrt.ErrorResolve("QTimeZone", "systemTimeZoneId", args)
  }

}

// isTimeZoneIdAvailable(const class QByteArray &)
func (this *QTimeZone) isTimeZoneIdAvailable_s(args ...interface{}) () {
  // isTimeZoneIdAvailable(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray
    // invoke: bool isTimeZoneIdAvailable(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "isTimeZoneIdAvailable", args)
  }

}

// abbreviation(const class QDateTime &)
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
    // invoke: QString abbreviation(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone12abbreviationERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "abbreviation", args)
  }

}

// windowsIdToDefaultIanaId(const class QByteArray &)
func (this *QTimeZone) windowsIdToDefaultIanaId_s(args ...interface{}) () {
  // windowsIdToDefaultIanaId(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray
    // invoke: QByteArray windowsIdToDefaultIanaId(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "windowsIdToDefaultIanaId", args)
  }

}

// QTimeZone(const class QTimeZone &)
func NewQTimeZone(args ...interface{}) QTimeZone {
  // QTimeZone(const class QTimeZone &)
  // QTimeZone(int)
  // QTimeZone()
  // QTimeZone(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTimeZone{}) // "const QTimeZone &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZoneC1ERKS_
    // invoke: void QTimeZone(const class QTimeZone &)
    var arg0 = args[0].(QTimeZone).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QTimeZoneC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN9QTimeZoneC1Ei
    // invoke: void QTimeZone(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QTimeZoneC2Ei(qthis, arg0)
  case 2:
    // invoke: _ZN9QTimeZoneC1Ev
    // invoke: void QTimeZone()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QTimeZoneC2Ev(qthis)
  case 3:
    // invoke: _ZN9QTimeZoneC1ERK10QByteArray
    // invoke: void QTimeZone(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QTimeZoneC2ERK10QByteArray(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "QTimeZone", args)
  }

  return QTimeZone{}
}

// transitions(const class QDateTime &, const class QDateTime &)
func (this *QTimeZone) transitions(args ...interface{}) () {
  // transitions(const class QDateTime &, const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"
  vtys[0][1] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone11transitionsERK9QDateTimeS2_
    // invoke: OffsetDataList transitions(const class QDateTime &, const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QDateTime).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK9QTimeZone11transitionsERK9QDateTimeS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTimeZone", "transitions", args)
  }

}

// nextTransition(const class QDateTime &)
func (this *QTimeZone) nextTransition(args ...interface{}) () {
  // nextTransition(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone14nextTransitionERK9QDateTime
    // invoke: QTimeZone::OffsetData nextTransition(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone14nextTransitionERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "nextTransition", args)
  }

}

// id()
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
    // invoke: QByteArray id()
    C.C_ZNK9QTimeZone2idEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "id", args)
  }

}

// offsetFromUtc(const class QDateTime &)
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
    // invoke: int offsetFromUtc(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone13offsetFromUtcERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "offsetFromUtc", args)
  }

}

// previousTransition(const class QDateTime &)
func (this *QTimeZone) previousTransition(args ...interface{}) () {
  // previousTransition(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone18previousTransitionERK9QDateTime
    // invoke: QTimeZone::OffsetData previousTransition(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone18previousTransitionERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "previousTransition", args)
  }

}

// availableTimeZoneIds()
func (this *QTimeZone) availableTimeZoneIds_s(args ...interface{}) () {
  // availableTimeZoneIds()
  // availableTimeZoneIds(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZone20availableTimeZoneIdsEv
    // invoke: QList<QByteArray> availableTimeZoneIds()
    C.C_ZN9QTimeZone20availableTimeZoneIdsEv()
  case 1:
    // invoke: _ZN9QTimeZone20availableTimeZoneIdsEi
    // invoke: QList<QByteArray> availableTimeZoneIds(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeZone20availableTimeZoneIdsEi(arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "availableTimeZoneIds", args)
  }

}

// ~QTimeZone()
func (this *QTimeZone) FreeQTimeZone(args ...interface{}) () {
  // ~QTimeZone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZoneD0Ev
    // invoke: void ~QTimeZone()
    C.C_ZN9QTimeZoneD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "~QTimeZone", args)
  }

}

// windowsIdToIanaIds(const class QByteArray &)
func (this *QTimeZone) windowsIdToIanaIds_s(args ...interface{}) () {
  // windowsIdToIanaIds(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZone18windowsIdToIanaIdsERK10QByteArray
    // invoke: QList<QByteArray> windowsIdToIanaIds(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeZone18windowsIdToIanaIdsERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "windowsIdToIanaIds", args)
  }

}

// systemTimeZone()
func (this *QTimeZone) systemTimeZone_s(args ...interface{}) () {
  // systemTimeZone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZone14systemTimeZoneEv
    // invoke: QTimeZone systemTimeZone()
    C.C_ZN9QTimeZone14systemTimeZoneEv()
  default:
    qtrt.ErrorResolve("QTimeZone", "systemTimeZone", args)
  }

}

// swap(class QTimeZone &)
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
    // invoke: void swap(class QTimeZone &)
    var arg0 = args[0].(QTimeZone).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeZone4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "swap", args)
  }

}

// offsetData(const class QDateTime &)
func (this *QTimeZone) offsetData(args ...interface{}) () {
  // offsetData(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone10offsetDataERK9QDateTime
    // invoke: QTimeZone::OffsetData offsetData(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone10offsetDataERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "offsetData", args)
  }

}

// standardTimeOffset(const class QDateTime &)
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
    // invoke: int standardTimeOffset(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone18standardTimeOffsetERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "standardTimeOffset", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C.C_ZNK9QTimeZone7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "isValid", args)
  }

}

// hasDaylightTime()
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
    // invoke: bool hasDaylightTime()
    C.C_ZNK9QTimeZone15hasDaylightTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "hasDaylightTime", args)
  }

}

// daylightTimeOffset(const class QDateTime &)
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
    // invoke: int daylightTimeOffset(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "daylightTimeOffset", args)
  }

}

// utc()
func (this *QTimeZone) utc_s(args ...interface{}) () {
  // utc()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZone3utcEv
    // invoke: QTimeZone utc()
    C.C_ZN9QTimeZone3utcEv()
  default:
    qtrt.ErrorResolve("QTimeZone", "utc", args)
  }

}

// hasTransitions()
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
    // invoke: bool hasTransitions()
    C.C_ZNK9QTimeZone14hasTransitionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "hasTransitions", args)
  }

}

// ianaIdToWindowsId(const class QByteArray &)
func (this *QTimeZone) ianaIdToWindowsId_s(args ...interface{}) () {
  // ianaIdToWindowsId(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray
    // invoke: QByteArray ianaIdToWindowsId(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "ianaIdToWindowsId", args)
  }

}

// country()
func (this *QTimeZone) country(args ...interface{}) () {
  // country()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeZone7countryEv
    // invoke: QLocale::Country country()
    C.C_ZNK9QTimeZone7countryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "country", args)
  }

}

// isDaylightTime(const class QDateTime &)
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
    // invoke: bool isDaylightTime(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone14isDaylightTimeERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "isDaylightTime", args)
  }

}

// <= body block end

