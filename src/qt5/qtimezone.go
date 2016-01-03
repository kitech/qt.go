package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto: static QList<QByteArray> QTimeZone::availableTimeZoneIds();
extern void _ZN9QTimeZone20availableTimeZoneIdsEv();
  // proto:  void QTimeZone::swap(QTimeZone & other);
extern void demth_ZN9QTimeZone4swapERS_(void* qthis, void* arg0);
  // proto:  bool QTimeZone::isValid();
extern void _ZNK9QTimeZone7isValidEv(void* qthis);
  // proto:  bool QTimeZone::hasDaylightTime();
extern void _ZNK9QTimeZone15hasDaylightTimeEv(void* qthis);
  // proto: static QTimeZone QTimeZone::utc();
extern void _ZN9QTimeZone3utcEv();
  // proto: static QList<QByteArray> QTimeZone::availableTimeZoneIds(int offsetSeconds);
extern void _ZN9QTimeZone20availableTimeZoneIdsEi(int32_t arg0);
  // proto:  void QTimeZone::QTimeZone(int offsetSeconds);
extern void* dector_ZN9QTimeZoneC1Ei(int32_t arg0);
extern void _ZN9QTimeZoneC1Ei(void* qthis, int32_t arg0);
  // proto:  QString QTimeZone::abbreviation(const QDateTime & atDateTime);
extern void _ZNK9QTimeZone12abbreviationERK9QDateTime(void* qthis, void* arg0);
  // proto:  void QTimeZone::QTimeZone();
extern void* dector_ZN9QTimeZoneC1Ev();
extern void _ZN9QTimeZoneC1Ev(void* qthis);
  // proto: static QByteArray QTimeZone::ianaIdToWindowsId(const QByteArray & ianaId);
extern void _ZN9QTimeZone17ianaIdToWindowsIdERK10QByteArray(void* arg0);
  // proto: static QByteArray QTimeZone::systemTimeZoneId();
extern void _ZN9QTimeZone16systemTimeZoneIdEv();
  // proto:  bool QTimeZone::isDaylightTime(const QDateTime & atDateTime);
extern void _ZNK9QTimeZone14isDaylightTimeERK9QDateTime(void* qthis, void* arg0);
  // proto: static bool QTimeZone::isTimeZoneIdAvailable(const QByteArray & ianaId);
extern void _ZN9QTimeZone21isTimeZoneIdAvailableERK10QByteArray(void* arg0);
  // proto:  QString QTimeZone::comment();
extern void _ZNK9QTimeZone7commentEv(void* qthis);
  // proto: static QByteArray QTimeZone::windowsIdToDefaultIanaId(const QByteArray & windowsId);
extern void _ZN9QTimeZone24windowsIdToDefaultIanaIdERK10QByteArray(void* arg0);
  // proto:  bool QTimeZone::hasTransitions();
extern void _ZNK9QTimeZone14hasTransitionsEv(void* qthis);
  // proto:  int QTimeZone::daylightTimeOffset(const QDateTime & atDateTime);
extern void _ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime(void* qthis, void* arg0);
  // proto: static QTimeZone QTimeZone::systemTimeZone();
extern void _ZN9QTimeZone14systemTimeZoneEv();
  // proto:  void QTimeZone::QTimeZone(const QByteArray & ianaId);
extern void* dector_ZN9QTimeZoneC1ERK10QByteArray(void* arg0);
extern void _ZN9QTimeZoneC1ERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QTimeZone::QTimeZone(const QTimeZone & other);
extern void* dector_ZN9QTimeZoneC1ERKS_(void* arg0);
extern void _ZN9QTimeZoneC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTimeZone::~QTimeZone();
extern void _ZN9QTimeZoneD0Ev(void* qthis);
  // proto:  int QTimeZone::standardTimeOffset(const QDateTime & atDateTime);
extern void _ZNK9QTimeZone18standardTimeOffsetERK9QDateTime(void* qthis, void* arg0);
  // proto:  QByteArray QTimeZone::id();
extern void _ZNK9QTimeZone2idEv(void* qthis);
  // proto:  int QTimeZone::offsetFromUtc(const QDateTime & atDateTime);
extern void _ZNK9QTimeZone13offsetFromUtcERK9QDateTime(void* qthis, void* arg0);
  // proto: static QList<QByteArray> QTimeZone::windowsIdToIanaIds(const QByteArray & windowsId);
extern void _ZN9QTimeZone18windowsIdToIanaIdsERK10QByteArray(void* arg0);
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

  // proto: static QList<QByteArray> QTimeZone::availableTimeZoneIds();
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

  // proto:  void QTimeZone::swap(QTimeZone & other);
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
    C.demth_ZN9QTimeZone4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "swap", args)
  }

}

  // proto:  bool QTimeZone::isValid();
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
    C._ZNK9QTimeZone7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "isValid", args)
  }

}

  // proto:  bool QTimeZone::hasDaylightTime();
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
    C._ZNK9QTimeZone15hasDaylightTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "hasDaylightTime", args)
  }

}

  // proto: static QTimeZone QTimeZone::utc();
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

  // proto:  void QTimeZone::QTimeZone(int offsetSeconds);
func NewQTimeZone(args ...interface{}) QTimeZone {
  return QTimeZone{}
}

  // proto:  QString QTimeZone::abbreviation(const QDateTime & atDateTime);
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
    C._ZNK9QTimeZone12abbreviationERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "abbreviation", args)
  }

}

  // proto: static QByteArray QTimeZone::ianaIdToWindowsId(const QByteArray & ianaId);
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

  // proto: static QByteArray QTimeZone::systemTimeZoneId();
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

  // proto:  bool QTimeZone::isDaylightTime(const QDateTime & atDateTime);
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
    C._ZNK9QTimeZone14isDaylightTimeERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "isDaylightTime", args)
  }

}

  // proto: static bool QTimeZone::isTimeZoneIdAvailable(const QByteArray & ianaId);
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

  // proto:  QString QTimeZone::comment();
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
    C._ZNK9QTimeZone7commentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "comment", args)
  }

}

  // proto: static QByteArray QTimeZone::windowsIdToDefaultIanaId(const QByteArray & windowsId);
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

  // proto:  bool QTimeZone::hasTransitions();
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
    C._ZNK9QTimeZone14hasTransitionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "hasTransitions", args)
  }

}

  // proto:  int QTimeZone::daylightTimeOffset(const QDateTime & atDateTime);
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
    C._ZNK9QTimeZone18daylightTimeOffsetERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "daylightTimeOffset", args)
  }

}

  // proto: static QTimeZone QTimeZone::systemTimeZone();
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

  // proto:  void QTimeZone::~QTimeZone();
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

  // proto:  int QTimeZone::standardTimeOffset(const QDateTime & atDateTime);
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
    C._ZNK9QTimeZone18standardTimeOffsetERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "standardTimeOffset", args)
  }

}

  // proto:  QByteArray QTimeZone::id();
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
    C._ZNK9QTimeZone2idEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeZone", "id", args)
  }

}

  // proto:  int QTimeZone::offsetFromUtc(const QDateTime & atDateTime);
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
    C._ZNK9QTimeZone13offsetFromUtcERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTimeZone", "offsetFromUtc", args)
  }

}

  // proto: static QList<QByteArray> QTimeZone::windowsIdToIanaIds(const QByteArray & windowsId);
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

