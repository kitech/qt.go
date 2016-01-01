package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qdatetimeedit.h
// dst-file: /src/widgets/qdatetimeedit.go
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
// class sizeof(QTimeEdit)=1
type QTimeEdit struct {
  /*qbase*/ QDateTimeEdit;
  qclsinst uint64 /* *mut c_void*/;
//  _userTimeChanged QTimeEdit_userTimeChanged_signal;
}

// class sizeof(QDateEdit)=1
type QDateEdit struct {
  /*qbase*/ QDateTimeEdit;
  qclsinst uint64 /* *mut c_void*/;
//  _userDateChanged QDateEdit_userDateChanged_signal;
}

// class sizeof(QDateTimeEdit)=1
type QDateTimeEdit struct {
  /*qbase*/ QAbstractSpinBox;
  qclsinst uint64 /* *mut c_void*/;
//  _dateChanged QDateTimeEdit_dateChanged_signal;
//  _timeChanged QDateTimeEdit_timeChanged_signal;
//  _dateTimeChanged QDateTimeEdit_dateTimeChanged_signal;
}


func NewQTimeEdit(args ...interface{}) QTimeEdit {
  return QTimeEdit{}
}


func (this *QTimeEdit) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTimeEdit10metaObjectEv
  default:
    qtrt.ErrorResolve("QTimeEdit", "metaObject", args)
  }

}


func (this *QTimeEdit) FreeQTimeEdit(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTimeEdit", "~QTimeEdit", args)
  }

}


func NewQDateEdit(args ...interface{}) QDateEdit {
  return QDateEdit{}
}


func (this *QDateEdit) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDateEdit10metaObjectEv
  default:
    qtrt.ErrorResolve("QDateEdit", "metaObject", args)
  }

}


func (this *QDateEdit) FreeQDateEdit(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateEdit", "~QDateEdit", args)
  }

}


func NewQDateTimeEdit(args ...interface{}) QDateTimeEdit {
  return QDateTimeEdit{}
}


func (this *QDateTimeEdit) date(args ...interface{}) () {
  // date()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit4dateEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "date", args)
  }

}


func (this *QDateTimeEdit) clearMinimumDateTime(args ...interface{}) () {
  // clearMinimumDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit20clearMinimumDateTimeEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMinimumDateTime", args)
  }

}


func (this *QDateTimeEdit) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit5clearEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clear", args)
  }

}


func (this *QDateTimeEdit) setDate(args ...interface{}) () {
  // setDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit7setDateERK5QDate
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDate", args)
  }

}


func (this *QDateTimeEdit) maximumDateTime(args ...interface{}) () {
  // maximumDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit15maximumDateTimeEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "maximumDateTime", args)
  }

}


func (this *QDateTimeEdit) setMinimumDate(args ...interface{}) () {
  // setMinimumDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit14setMinimumDateERK5QDate
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMinimumDate", args)
  }

}


func (this *QDateTimeEdit) setMaximumDate(args ...interface{}) () {
  // setMaximumDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit14setMaximumDateERK5QDate
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMaximumDate", args)
  }

}


func (this *QDateTimeEdit) maximumTime(args ...interface{}) () {
  // maximumTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit11maximumTimeEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "maximumTime", args)
  }

}


func (this *QDateTimeEdit) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit10metaObjectEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "metaObject", args)
  }

}


func (this *QDateTimeEdit) FreeQDateTimeEdit(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "~QDateTimeEdit", args)
  }

}


func (this *QDateTimeEdit) clearMinimumDate(args ...interface{}) () {
  // clearMinimumDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16clearMinimumDateEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMinimumDate", args)
  }

}


func (this *QDateTimeEdit) setDateRange(args ...interface{}) () {
  // setDateRange(const class QDate &, const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[0][1] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit12setDateRangeERK5QDateS2_
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDateRange", args)
  }

}


func (this *QDateTimeEdit) setMinimumTime(args ...interface{}) () {
  // setMinimumTime(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit14setMinimumTimeERK5QTime
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMinimumTime", args)
  }

}


func (this *QDateTimeEdit) time(args ...interface{}) () {
  // time()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit4timeEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "time", args)
  }

}


func (this *QDateTimeEdit) currentSectionIndex(args ...interface{}) () {
  // currentSectionIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit19currentSectionIndexEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "currentSectionIndex", args)
  }

}


func (this *QDateTimeEdit) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit5eventEP6QEvent
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "event", args)
  }

}


func (this *QDateTimeEdit) setDateTime(args ...interface{}) () {
  // setDateTime(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit11setDateTimeERK9QDateTime
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDateTime", args)
  }

}


func (this *QDateTimeEdit) setCalendarWidget(args ...interface{}) () {
  // setCalendarWidget(class QCalendarWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCalendarWidget{}) // "QCalendarWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit17setCalendarWidgetEP15QCalendarWidget
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setCalendarWidget", args)
  }

}


func (this *QDateTimeEdit) setDateTimeRange(args ...interface{}) () {
  // setDateTimeRange(const class QDateTime &, const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"
  vtys[0][1] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16setDateTimeRangeERK9QDateTimeS2_
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDateTimeRange", args)
  }

}


func (this *QDateTimeEdit) setTime(args ...interface{}) () {
  // setTime(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit7setTimeERK5QTime
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setTime", args)
  }

}


func (this *QDateTimeEdit) dateTime(args ...interface{}) () {
  // dateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit8dateTimeEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "dateTime", args)
  }

}


func (this *QDateTimeEdit) setMaximumTime(args ...interface{}) () {
  // setMaximumTime(const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit14setMaximumTimeERK5QTime
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMaximumTime", args)
  }

}


func (this *QDateTimeEdit) setMinimumDateTime(args ...interface{}) () {
  // setMinimumDateTime(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit18setMinimumDateTimeERK9QDateTime
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMinimumDateTime", args)
  }

}


func (this *QDateTimeEdit) clearMaximumTime(args ...interface{}) () {
  // clearMaximumTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16clearMaximumTimeEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMaximumTime", args)
  }

}


func (this *QDateTimeEdit) calendarPopup(args ...interface{}) () {
  // calendarPopup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit13calendarPopupEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "calendarPopup", args)
  }

}


func (this *QDateTimeEdit) setMaximumDateTime(args ...interface{}) () {
  // setMaximumDateTime(const class QDateTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDateTime{}) // "const QDateTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit18setMaximumDateTimeERK9QDateTime
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMaximumDateTime", args)
  }

}


func (this *QDateTimeEdit) clearMaximumDateTime(args ...interface{}) () {
  // clearMaximumDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit20clearMaximumDateTimeEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMaximumDateTime", args)
  }

}


func (this *QDateTimeEdit) minimumDateTime(args ...interface{}) () {
  // minimumDateTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit15minimumDateTimeEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "minimumDateTime", args)
  }

}


func (this *QDateTimeEdit) calendarWidget(args ...interface{}) () {
  // calendarWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit14calendarWidgetEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "calendarWidget", args)
  }

}


func (this *QDateTimeEdit) setDisplayFormat(args ...interface{}) () {
  // setDisplayFormat(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16setDisplayFormatERK7QString
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDisplayFormat", args)
  }

}


func (this *QDateTimeEdit) clearMaximumDate(args ...interface{}) () {
  // clearMaximumDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16clearMaximumDateEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMaximumDate", args)
  }

}


func (this *QDateTimeEdit) setCalendarPopup(args ...interface{}) () {
  // setCalendarPopup(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16setCalendarPopupEb
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setCalendarPopup", args)
  }

}


func (this *QDateTimeEdit) stepBy(args ...interface{}) () {
  // stepBy(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit6stepByEi
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "stepBy", args)
  }

}


func (this *QDateTimeEdit) displayFormat(args ...interface{}) () {
  // displayFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit13displayFormatEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "displayFormat", args)
  }

}


func (this *QDateTimeEdit) minimumTime(args ...interface{}) () {
  // minimumTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit11minimumTimeEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "minimumTime", args)
  }

}


func (this *QDateTimeEdit) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit8sizeHintEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "sizeHint", args)
  }

}


func (this *QDateTimeEdit) sectionCount(args ...interface{}) () {
  // sectionCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit12sectionCountEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "sectionCount", args)
  }

}


func (this *QDateTimeEdit) setCurrentSectionIndex(args ...interface{}) () {
  // setCurrentSectionIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit22setCurrentSectionIndexEi
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setCurrentSectionIndex", args)
  }

}


func (this *QDateTimeEdit) clearMinimumTime(args ...interface{}) () {
  // clearMinimumTime()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit16clearMinimumTimeEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMinimumTime", args)
  }

}


func (this *QDateTimeEdit) setTimeRange(args ...interface{}) () {
  // setTimeRange(const class QTime &, const class QTime &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTime{}) // "const QTime &"
  vtys[0][1] = reflect.TypeOf(QTime{}) // "const QTime &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QDateTimeEdit12setTimeRangeERK5QTimeS2_
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setTimeRange", args)
  }

}


func (this *QDateTimeEdit) minimumDate(args ...interface{}) () {
  // minimumDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit11minimumDateEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "minimumDate", args)
  }

}


func (this *QDateTimeEdit) maximumDate(args ...interface{}) () {
  // maximumDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QDateTimeEdit11maximumDateEv
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "maximumDate", args)
  }

}

// <= body block end

