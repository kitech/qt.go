package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qdatetime.h
// dst-file: /src/core/qdatetime.go
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
// class sizeof(QTime)=4
type QTime struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDateTime)=1
type QDateTime struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QDate)=8
type QDate struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTime) addMSecs(args ...interface{}) () {
  // addMSecs(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime8addMSecsEi
  default:
    qtrt.ErrorResolve("QTime", "addMSecs", args)
  }

}


func (this *QTime) fromMSecsSinceStartOfDay_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTime", "fromMSecsSinceStartOfDay", args)
  }

}


func (this *QTime) currentTime_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTime", "currentTime", args)
  }

}


func (this *QTime) second(args ...interface{}) () {
  // second()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime6secondEv
  default:
    qtrt.ErrorResolve("QTime", "second", args)
  }

}


func (this *QTime) restart(args ...interface{}) () {
  // restart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTime7restartEv
  default:
    qtrt.ErrorResolve("QTime", "restart", args)
  }

}


func (this *QTime) start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTime5startEv
  default:
    qtrt.ErrorResolve("QTime", "start", args)
  }

}


func (this *QTime) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime6isNullEv
  default:
    qtrt.ErrorResolve("QTime", "isNull", args)
  }

}


func (this *QTime) msecsSinceStartOfDay(args ...interface{}) () {
  // msecsSinceStartOfDay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime20msecsSinceStartOfDayEv
  default:
    qtrt.ErrorResolve("QTime", "msecsSinceStartOfDay", args)
  }

}


func (this *QTime) hour(args ...interface{}) () {
  // hour()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime4hourEv
  default:
    qtrt.ErrorResolve("QTime", "hour", args)
  }

}


func (this *QTime) elapsed(args ...interface{}) () {
  // elapsed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime7elapsedEv
  default:
    qtrt.ErrorResolve("QTime", "elapsed", args)
  }

}


func (this *QTime) addSecs(args ...interface{}) () {
  // addSecs(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime7addSecsEi
  default:
    qtrt.ErrorResolve("QTime", "addSecs", args)
  }

}


func (this *QTime) isValid(args ...interface{}) () {
  // isValid()
  // isValid(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime7isValidEv
  case 1:
    // invoke: _ZN5QTime7isValidEiiii
  default:
    qtrt.ErrorResolve("QTime", "isValid", args)
  }

}


func NewQTime(args ...interface{}) QTime {
  return QTime{}
}


func (this *QTime) msec(args ...interface{}) () {
  // msec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime4msecEv
  default:
    qtrt.ErrorResolve("QTime", "msec", args)
  }

}


func (this *QTime) secsTo(args ...interface{}) () {
  // secsTo(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime6secsToERKS_
  default:
    qtrt.ErrorResolve("QTime", "secsTo", args)
  }

}


func (this *QTime) setHMS(args ...interface{}) () {
  // setHMS(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QTime6setHMSEiiii
  default:
    qtrt.ErrorResolve("QTime", "setHMS", args)
  }

}


func (this *QTime) toString(args ...interface{}) () {
  // toString(const class QString &)
  // toString(Qt::DateFormat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::DateFormat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime8toStringERK7QString
  case 1:
    // invoke: _ZNK5QTime8toStringEN2Qt10DateFormatE
  default:
    qtrt.ErrorResolve("QTime", "toString", args)
  }

}


func (this *QTime) msecsTo(args ...interface{}) () {
  // msecsTo(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime7msecsToERKS_
  default:
    qtrt.ErrorResolve("QTime", "msecsTo", args)
  }

}


func (this *QTime) minute(args ...interface{}) () {
  // minute()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QTime6minuteEv
  default:
    qtrt.ErrorResolve("QTime", "minute", args)
  }

}


func (this *QTime) isValid_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTime", "isValid", args)
  }

}


func (this *QTime) fromString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTime", "fromString", args)
  }

}


func (this *QDateTime) toLocalTime(args ...interface{}) () {
  // toLocalTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime11toLocalTimeEv
  default:
    qtrt.ErrorResolve("QDateTime", "toLocalTime", args)
  }

}


func (this *QDateTime) setOffsetFromUtc(args ...interface{}) () {
  // setOffsetFromUtc(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime16setOffsetFromUtcEi
  default:
    qtrt.ErrorResolve("QDateTime", "setOffsetFromUtc", args)
  }

}


func (this *QDateTime) timeZone(args ...interface{}) () {
  // timeZone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8timeZoneEv
  default:
    qtrt.ErrorResolve("QDateTime", "timeZone", args)
  }

}


func (this *QDateTime) setTime(args ...interface{}) () {
  // setTime(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime7setTimeERK5QTime
  default:
    qtrt.ErrorResolve("QDateTime", "setTime", args)
  }

}


func (this *QDateTime) toMSecsSinceEpoch(args ...interface{}) () {
  // toMSecsSinceEpoch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime17toMSecsSinceEpochEv
  default:
    qtrt.ErrorResolve("QDateTime", "toMSecsSinceEpoch", args)
  }

}


func (this *QDateTime) setTime_t(args ...interface{}) () {
  // setTime_t(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime9setTime_tEj
  default:
    qtrt.ErrorResolve("QDateTime", "setTime_t", args)
  }

}


func NewQDateTime(args ...interface{}) QDateTime {
  return QDateTime{}
}


func (this *QDateTime) isDaylightTime(args ...interface{}) () {
  // isDaylightTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime14isDaylightTimeEv
  default:
    qtrt.ErrorResolve("QDateTime", "isDaylightTime", args)
  }

}


func (this *QDateTime) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime7isValidEv
  default:
    qtrt.ErrorResolve("QDateTime", "isValid", args)
  }

}


func (this *QDateTime) toString(args ...interface{}) () {
  // toString(Qt::DateFormat)
  // toString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "Qt::DateFormat"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8toStringEN2Qt10DateFormatE
  case 1:
    // invoke: _ZNK9QDateTime8toStringERK7QString
  default:
    qtrt.ErrorResolve("QDateTime", "toString", args)
  }

}


func (this *QDateTime) addYears(args ...interface{}) () {
  // addYears(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8addYearsEi
  default:
    qtrt.ErrorResolve("QDateTime", "addYears", args)
  }

}


func (this *QDateTime) setMSecsSinceEpoch(args ...interface{}) () {
  // setMSecsSinceEpoch(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime18setMSecsSinceEpochEx
  default:
    qtrt.ErrorResolve("QDateTime", "setMSecsSinceEpoch", args)
  }

}


func (this *QDateTime) toOffsetFromUtc(args ...interface{}) () {
  // toOffsetFromUtc(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime15toOffsetFromUtcEi
  default:
    qtrt.ErrorResolve("QDateTime", "toOffsetFromUtc", args)
  }

}


func (this *QDateTime) setUtcOffset(args ...interface{}) () {
  // setUtcOffset(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime12setUtcOffsetEi
  default:
    qtrt.ErrorResolve("QDateTime", "setUtcOffset", args)
  }

}


func (this *QDateTime) addSecs(args ...interface{}) () {
  // addSecs(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime7addSecsEx
  default:
    qtrt.ErrorResolve("QDateTime", "addSecs", args)
  }

}


func (this *QDateTime) fromMSecsSinceEpoch_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "fromMSecsSinceEpoch", args)
  }

}


func (this *QDateTime) fromString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "fromString", args)
  }

}


func (this *QDateTime) swap(args ...interface{}) () {
  // swap(class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime4swapERS_
  default:
    qtrt.ErrorResolve("QDateTime", "swap", args)
  }

}


func (this *QDateTime) toTime_t(args ...interface{}) () {
  // toTime_t()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8toTime_tEv
  default:
    qtrt.ErrorResolve("QDateTime", "toTime_t", args)
  }

}


func (this *QDateTime) timeZoneAbbreviation(args ...interface{}) () {
  // timeZoneAbbreviation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime20timeZoneAbbreviationEv
  default:
    qtrt.ErrorResolve("QDateTime", "timeZoneAbbreviation", args)
  }

}


func (this *QDateTime) toUTC(args ...interface{}) () {
  // toUTC()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime5toUTCEv
  default:
    qtrt.ErrorResolve("QDateTime", "toUTC", args)
  }

}


func (this *QDateTime) date(args ...interface{}) () {
  // date()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime4dateEv
  default:
    qtrt.ErrorResolve("QDateTime", "date", args)
  }

}


func (this *QDateTime) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime6isNullEv
  default:
    qtrt.ErrorResolve("QDateTime", "isNull", args)
  }

}


func (this *QDateTime) currentMSecsSinceEpoch_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "currentMSecsSinceEpoch", args)
  }

}


func (this *QDateTime) offsetFromUtc(args ...interface{}) () {
  // offsetFromUtc()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime13offsetFromUtcEv
  default:
    qtrt.ErrorResolve("QDateTime", "offsetFromUtc", args)
  }

}


func (this *QDateTime) addMSecs(args ...interface{}) () {
  // addMSecs(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime8addMSecsEx
  default:
    qtrt.ErrorResolve("QDateTime", "addMSecs", args)
  }

}


func (this *QDateTime) secsTo(args ...interface{}) () {
  // secsTo(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime6secsToERKS_
  default:
    qtrt.ErrorResolve("QDateTime", "secsTo", args)
  }

}


func (this *QDateTime) FreeQDateTime(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "~QDateTime", args)
  }

}


func (this *QDateTime) addMonths(args ...interface{}) () {
  // addMonths(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime9addMonthsEi
  default:
    qtrt.ErrorResolve("QDateTime", "addMonths", args)
  }

}


func (this *QDateTime) currentDateTime_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "currentDateTime", args)
  }

}


func (this *QDateTime) toTimeZone(args ...interface{}) () {
  // toTimeZone(const class QTimeZone &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTimeZone{}) // "const QTimeZone &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime10toTimeZoneERK9QTimeZone
  default:
    qtrt.ErrorResolve("QDateTime", "toTimeZone", args)
  }

}


func (this *QDateTime) msecsTo(args ...interface{}) () {
  // msecsTo(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime7msecsToERKS_
  default:
    qtrt.ErrorResolve("QDateTime", "msecsTo", args)
  }

}


func (this *QDateTime) fromTime_t_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "fromTime_t", args)
  }

}


func (this *QDateTime) setDate(args ...interface{}) () {
  // setDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime7setDateERK5QDate
  default:
    qtrt.ErrorResolve("QDateTime", "setDate", args)
  }

}


func (this *QDateTime) utcOffset(args ...interface{}) () {
  // utcOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime9utcOffsetEv
  default:
    qtrt.ErrorResolve("QDateTime", "utcOffset", args)
  }

}


func (this *QDateTime) time(args ...interface{}) () {
  // time()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime4timeEv
  default:
    qtrt.ErrorResolve("QDateTime", "time", args)
  }

}


func (this *QDateTime) daysTo(args ...interface{}) () {
  // daysTo(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime6daysToERKS_
  default:
    qtrt.ErrorResolve("QDateTime", "daysTo", args)
  }

}


func (this *QDateTime) addDays(args ...interface{}) () {
  // addDays(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateTime7addDaysEx
  default:
    qtrt.ErrorResolve("QDateTime", "addDays", args)
  }

}


func (this *QDateTime) setTimeZone(args ...interface{}) () {
  // setTimeZone(const class QTimeZone &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTimeZone{}) // "const QTimeZone &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDateTime11setTimeZoneERK9QTimeZone
  default:
    qtrt.ErrorResolve("QDateTime", "setTimeZone", args)
  }

}


func (this *QDateTime) currentDateTimeUtc_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTime", "currentDateTimeUtc", args)
  }

}


func (this *QDate) daysTo(args ...interface{}) () {
  // daysTo(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate6daysToERKS_
  default:
    qtrt.ErrorResolve("QDate", "daysTo", args)
  }

}


func (this *QDate) addYears(args ...interface{}) () {
  // addYears(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate8addYearsEi
  default:
    qtrt.ErrorResolve("QDate", "addYears", args)
  }

}


func (this *QDate) month(args ...interface{}) () {
  // month()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate5monthEv
  default:
    qtrt.ErrorResolve("QDate", "month", args)
  }

}


func (this *QDate) toJulianDay(args ...interface{}) () {
  // toJulianDay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate11toJulianDayEv
  default:
    qtrt.ErrorResolve("QDate", "toJulianDay", args)
  }

}


func NewQDate(args ...interface{}) QDate {
  return QDate{}
}


func (this *QDate) getDate(args ...interface{}) () {
  // getDate(int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDate7getDateEPiS0_S0_
  default:
    qtrt.ErrorResolve("QDate", "getDate", args)
  }

}


func (this *QDate) currentDate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDate", "currentDate", args)
  }

}


func (this *QDate) weekNumber(args ...interface{}) () {
  // weekNumber(int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate10weekNumberEPi
  default:
    qtrt.ErrorResolve("QDate", "weekNumber", args)
  }

}


func (this *QDate) toString(args ...interface{}) () {
  // toString(const class QString &)
  // toString(Qt::DateFormat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::DateFormat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate8toStringERK7QString
  case 1:
    // invoke: _ZNK5QDate8toStringEN2Qt10DateFormatE
  default:
    qtrt.ErrorResolve("QDate", "toString", args)
  }

}


func (this *QDate) dayOfYear(args ...interface{}) () {
  // dayOfYear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate9dayOfYearEv
  default:
    qtrt.ErrorResolve("QDate", "dayOfYear", args)
  }

}


func (this *QDate) day(args ...interface{}) () {
  // day()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate3dayEv
  default:
    qtrt.ErrorResolve("QDate", "day", args)
  }

}


func (this *QDate) setDate(args ...interface{}) () {
  // setDate(int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDate7setDateEiii
  default:
    qtrt.ErrorResolve("QDate", "setDate", args)
  }

}


func (this *QDate) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate6isNullEv
  default:
    qtrt.ErrorResolve("QDate", "isNull", args)
  }

}


func (this *QDate) fromJulianDay_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDate", "fromJulianDay", args)
  }

}


func (this *QDate) isValid(args ...interface{}) () {
  // isValid()
  // isValid(int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate7isValidEv
  case 1:
    // invoke: _ZN5QDate7isValidEiii
  default:
    qtrt.ErrorResolve("QDate", "isValid", args)
  }

}


func (this *QDate) addDays(args ...interface{}) () {
  // addDays(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate7addDaysEx
  default:
    qtrt.ErrorResolve("QDate", "addDays", args)
  }

}


func (this *QDate) isValid_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDate", "isValid", args)
  }

}


func (this *QDate) daysInMonth(args ...interface{}) () {
  // daysInMonth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate11daysInMonthEv
  default:
    qtrt.ErrorResolve("QDate", "daysInMonth", args)
  }

}


func (this *QDate) fromString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDate", "fromString", args)
  }

}


func (this *QDate) isLeapYear_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDate", "isLeapYear", args)
  }

}


func (this *QDate) daysInYear(args ...interface{}) () {
  // daysInYear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate10daysInYearEv
  default:
    qtrt.ErrorResolve("QDate", "daysInYear", args)
  }

}


func (this *QDate) dayOfWeek(args ...interface{}) () {
  // dayOfWeek()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate9dayOfWeekEv
  default:
    qtrt.ErrorResolve("QDate", "dayOfWeek", args)
  }

}


func (this *QDate) addMonths(args ...interface{}) () {
  // addMonths(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate9addMonthsEi
  default:
    qtrt.ErrorResolve("QDate", "addMonths", args)
  }

}


func (this *QDate) year(args ...interface{}) () {
  // year()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDate4yearEv
  default:
    qtrt.ErrorResolve("QDate", "year", args)
  }

}

// <= body block end

