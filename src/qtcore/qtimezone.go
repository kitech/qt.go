package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
extern void* C_ZNK9QTimeZone7commentEv(void* qthis); // 4
  // proto: static QByteArray QTimeZone::systemTimeZoneId();
extern void* C_ZN9QTimeZone16systemTimeZoneIdEv(); // 4
  // proto: static bool QTimeZone::isTimeZoneIdAvailable(const QByteArray & ianaId);
extern bool C_ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray(void* arg0); // 4
  // proto:  QString QTimeZone::abbreviation(const QDateTime & atDateTime);
extern void* C_ZNK9QTimeZone12abbreviationERK9QDateTime(void* qthis, void* arg0); // 4
  // proto: static QByteArray QTimeZone::windowsIdToDefaultIanaId(const QByteArray & windowsId);
extern void* C_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray(void* arg0); // 4
  // proto:  void QTimeZone::QTimeZone(const QTimeZone & other);
extern void* C_ZN9QTimeZoneC2ERKS_(void* arg0); // 3
  // proto:  void QTimeZone::QTimeZone(int offsetSeconds);
extern void* C_ZN9QTimeZoneC2Ei(int32_t arg0); // 3
  // proto:  void QTimeZone::QTimeZone();
extern void* C_ZN9QTimeZoneC2Ev(); // 3
  // proto:  void QTimeZone::QTimeZone(const QByteArray & ianaId);
extern void* C_ZN9QTimeZoneC2ERK10QByteArray(void* arg0); // 3
  // proto:  OffsetDataList QTimeZone::transitions(const QDateTime & fromDateTime, const QDateTime & toDateTime);
extern void C_ZNK9QTimeZone11transitionsERK9QDateTimeS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QTimeZone::OffsetData QTimeZone::nextTransition(const QDateTime & afterDateTime);
extern void C_ZNK9QTimeZone14nextTransitionERK9QDateTime(void* qthis, void* arg0); // 4
  // proto:  QByteArray QTimeZone::id();
extern void* C_ZNK9QTimeZone2idEv(void* qthis); // 4
  // proto:  int QTimeZone::offsetFromUtc(const QDateTime & atDateTime);
extern int32_t C_ZNK9QTimeZone13offsetFromUtcERK9QDateTime(void* qthis, void* arg0); // 4
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
extern void* C_ZN9QTimeZone14systemTimeZoneEv(); // 4
  // proto:  void QTimeZone::swap(QTimeZone & other);
extern void C_ZN9QTimeZone4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QTimeZone::OffsetData QTimeZone::offsetData(const QDateTime & forDateTime);
extern void C_ZNK9QTimeZone10offsetDataERK9QDateTime(void* qthis, void* arg0); // 4
  // proto:  int QTimeZone::standardTimeOffset(const QDateTime & atDateTime);
extern int32_t C_ZNK9QTimeZone18standardTimeOffsetERK9QDateTime(void* qthis, void* arg0); // 4
  // proto:  bool QTimeZone::isValid();
extern bool C_ZNK9QTimeZone7isValidEv(void* qthis); // 4
  // proto:  bool QTimeZone::hasDaylightTime();
extern bool C_ZNK9QTimeZone15hasDaylightTimeEv(void* qthis); // 4
  // proto:  int QTimeZone::daylightTimeOffset(const QDateTime & atDateTime);
extern int32_t C_ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime(void* qthis, void* arg0); // 4
  // proto: static QTimeZone QTimeZone::utc();
extern void* C_ZN9QTimeZone3utcEv(); // 4
  // proto:  bool QTimeZone::hasTransitions();
extern bool C_ZNK9QTimeZone14hasTransitionsEv(void* qthis); // 4
  // proto: static QByteArray QTimeZone::ianaIdToWindowsId(const QByteArray & ianaId);
extern void* C_ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray(void* arg0); // 4
  // proto:  QLocale::Country QTimeZone::country();
extern void C_ZNK9QTimeZone7countryEv(void* qthis); // 4
  // proto:  bool QTimeZone::isDaylightTime(const QDateTime & atDateTime);
extern bool C_ZNK9QTimeZone14isDaylightTimeERK9QDateTime(void* qthis, void* arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QTimeZone)=1
type QTimeZone struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// comment()
func (this *QTimeZone) Comment(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTimeZone7commentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "comment", args)
  }

  return
}

// systemTimeZoneId()
func (this *QTimeZone) SystemTimeZoneId_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN9QTimeZone16systemTimeZoneIdEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "systemTimeZoneId", args)
  }

  return
}

// isTimeZoneIdAvailable(const class QByteArray &)
func (this *QTimeZone) IsTimeZoneIdAvailable_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "isTimeZoneIdAvailable", args)
  }

  return
}

// abbreviation(const class QDateTime &)
func (this *QTimeZone) Abbreviation(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTimeZone12abbreviationERK9QDateTime(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "abbreviation", args)
  }

  return
}

// windowsIdToDefaultIanaId(const class QByteArray &)
func (this *QTimeZone) WindowsIdToDefaultIanaId_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "windowsIdToDefaultIanaId", args)
  }

  return
}

// QTimeZone(const class QTimeZone &)
func GcfreeQTimeZone(this *QTimeZone) {
  qtrt.UniverseFree(this)
}
func NewQTimeZone(args ...interface{}) *QTimeZone {
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
    var arg0 = args[0].(*QTimeZone).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTimeZoneC2ERKS_(arg0)
    this := &QTimeZone{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTimeZone)
    return this
  case 1:
    // invoke: _ZN9QTimeZoneC1Ei
    // invoke: void QTimeZone(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTimeZoneC2Ei(arg0)
    this := &QTimeZone{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTimeZone)
    return this
  case 2:
    // invoke: _ZN9QTimeZoneC1Ev
    // invoke: void QTimeZone()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTimeZoneC2Ev()
    this := &QTimeZone{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTimeZone)
    return this
  case 3:
    // invoke: _ZN9QTimeZoneC1ERK10QByteArray
    // invoke: void QTimeZone(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTimeZoneC2ERK10QByteArray(arg0)
    this := &QTimeZone{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTimeZone)
    return this
  default:
    qtrt.ErrorResolve("QTimeZone", "QTimeZone", args)
  }

  return nil // QTimeZone{Qclsinst:qthis}
}

// transitions(const class QDateTime &, const class QDateTime &)
func (this *QTimeZone) Transitions(args ...interface{}) () {
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
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QDateTime).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK9QTimeZone11transitionsERK9QDateTimeS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTimeZone", "transitions", args)
  }

  return
}

// nextTransition(const class QDateTime &)
func (this *QTimeZone) NextTransition(args ...interface{}) () {
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
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone14nextTransitionERK9QDateTime(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "nextTransition", args)
  }

  return
}

// id()
func (this *QTimeZone) Id(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTimeZone2idEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "id", args)
  }

  return
}

// offsetFromUtc(const class QDateTime &)
func (this *QTimeZone) OffsetFromUtc(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTimeZone13offsetFromUtcERK9QDateTime(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "offsetFromUtc", args)
  }

  return
}

// previousTransition(const class QDateTime &)
func (this *QTimeZone) PreviousTransition(args ...interface{}) () {
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
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone18previousTransitionERK9QDateTime(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "previousTransition", args)
  }

  return
}

// availableTimeZoneIds()
func (this *QTimeZone) AvailableTimeZoneIds_s(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeZone20availableTimeZoneIdsEi(arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "availableTimeZoneIds", args)
  }

  return
}

// ~QTimeZone()
func (this *QTimeZone) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN9QTimeZoneD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "~QTimeZone", args)
  }

  return
}

// windowsIdToIanaIds(const class QByteArray &)
func (this *QTimeZone) WindowsIdToIanaIds_s(args ...interface{}) () {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeZone18windowsIdToIanaIdsERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "windowsIdToIanaIds", args)
  }

  return
}

// systemTimeZone()
func (this *QTimeZone) SystemTimeZone_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN9QTimeZone14systemTimeZoneEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTimeZone{}) // "QTimeZone"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "systemTimeZone", args)
  }

  return
}

// swap(class QTimeZone &)
func (this *QTimeZone) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QTimeZone).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTimeZone4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "swap", args)
  }

  return
}

// offsetData(const class QDateTime &)
func (this *QTimeZone) OffsetData(args ...interface{}) () {
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
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QTimeZone10offsetDataERK9QDateTime(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "offsetData", args)
  }

  return
}

// standardTimeOffset(const class QDateTime &)
func (this *QTimeZone) StandardTimeOffset(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTimeZone18standardTimeOffsetERK9QDateTime(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "standardTimeOffset", args)
  }

  return
}

// isValid()
func (this *QTimeZone) IsValid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTimeZone7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "isValid", args)
  }

  return
}

// hasDaylightTime()
func (this *QTimeZone) HasDaylightTime(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTimeZone15hasDaylightTimeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "hasDaylightTime", args)
  }

  return
}

// daylightTimeOffset(const class QDateTime &)
func (this *QTimeZone) DaylightTimeOffset(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "daylightTimeOffset", args)
  }

  return
}

// utc()
func (this *QTimeZone) Utc_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN9QTimeZone3utcEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTimeZone{}) // "QTimeZone"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "utc", args)
  }

  return
}

// hasTransitions()
func (this *QTimeZone) HasTransitions(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QTimeZone14hasTransitionsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "hasTransitions", args)
  }

  return
}

// ianaIdToWindowsId(const class QByteArray &)
func (this *QTimeZone) IanaIdToWindowsId_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "ianaIdToWindowsId", args)
  }

  return
}

// country()
func (this *QTimeZone) Country(args ...interface{}) () {
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
    C.C_ZNK9QTimeZone7countryEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "country", args)
  }

  return
}

// isDaylightTime(const class QDateTime &)
func (this *QTimeZone) IsDaylightTime(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QDateTime).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTimeZone14isDaylightTimeERK9QDateTime(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTimeZone", "isDaylightTime", args)
  }

  return
}

// <= body block end

