package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qdatetimeedit.h
// dst-file: /src/widgets/qdatetimeedit.go
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
  // proto:  void QTimeEdit::QTimeEdit(QWidget * parent);
extern void* dector_ZN9QTimeEditC1EP7QWidget(void* arg0);
extern void _ZN9QTimeEditC1EP7QWidget(void* qthis, void* arg0);
  // proto:  const QMetaObject * QTimeEdit::metaObject();
extern void _ZNK9QTimeEdit10metaObjectEv(void* qthis);
  // proto:  void QTimeEdit::QTimeEdit(const QTime & time, QWidget * parent);
extern void* dector_ZN9QTimeEditC1ERK5QTimeP7QWidget(void* arg0, void* arg1);
extern void _ZN9QTimeEditC1ERK5QTimeP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QTimeEdit::~QTimeEdit();
extern void _ZN9QTimeEditD0Ev(void* qthis);
  // proto:  void QDateEdit::QDateEdit(const QDate & date, QWidget * parent);
extern void* dector_ZN9QDateEditC1ERK5QDateP7QWidget(void* arg0, void* arg1);
extern void _ZN9QDateEditC1ERK5QDateP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  const QMetaObject * QDateEdit::metaObject();
extern void _ZNK9QDateEdit10metaObjectEv(void* qthis);
  // proto:  void QDateEdit::QDateEdit(QWidget * parent);
extern void* dector_ZN9QDateEditC1EP7QWidget(void* arg0);
extern void _ZN9QDateEditC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QDateEdit::~QDateEdit();
extern void _ZN9QDateEditD0Ev(void* qthis);
  // proto:  void QDateTimeEdit::QDateTimeEdit(const QDateTimeEdit & );
extern void* dector_ZN13QDateTimeEditC1ERKS_(void* arg0);
extern void _ZN13QDateTimeEditC1ERKS_(void* qthis, void* arg0);
  // proto:  QDate QDateTimeEdit::date();
extern void _ZNK13QDateTimeEdit4dateEv(void* qthis);
  // proto:  void QDateTimeEdit::clearMinimumDateTime();
extern void _ZN13QDateTimeEdit20clearMinimumDateTimeEv(void* qthis);
  // proto:  void QDateTimeEdit::clear();
extern void _ZN13QDateTimeEdit5clearEv(void* qthis);
  // proto:  void QDateTimeEdit::QDateTimeEdit(const QTime & t, QWidget * parent);
extern void* dector_ZN13QDateTimeEditC1ERK5QTimeP7QWidget(void* arg0, void* arg1);
extern void _ZN13QDateTimeEditC1ERK5QTimeP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QDateTimeEdit::setDate(const QDate & date);
extern void _ZN13QDateTimeEdit7setDateERK5QDate(void* qthis, void* arg0);
  // proto:  QDateTime QDateTimeEdit::maximumDateTime();
extern void _ZNK13QDateTimeEdit15maximumDateTimeEv(void* qthis);
  // proto:  void QDateTimeEdit::setMinimumDate(const QDate & min);
extern void _ZN13QDateTimeEdit14setMinimumDateERK5QDate(void* qthis, void* arg0);
  // proto:  void QDateTimeEdit::QDateTimeEdit(const QDate & d, QWidget * parent);
extern void* dector_ZN13QDateTimeEditC1ERK5QDateP7QWidget(void* arg0, void* arg1);
extern void _ZN13QDateTimeEditC1ERK5QDateP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QDateTimeEdit::setMaximumDate(const QDate & max);
extern void _ZN13QDateTimeEdit14setMaximumDateERK5QDate(void* qthis, void* arg0);
  // proto:  QTime QDateTimeEdit::maximumTime();
extern void _ZNK13QDateTimeEdit11maximumTimeEv(void* qthis);
  // proto:  const QMetaObject * QDateTimeEdit::metaObject();
extern void _ZNK13QDateTimeEdit10metaObjectEv(void* qthis);
  // proto:  void QDateTimeEdit::~QDateTimeEdit();
extern void _ZN13QDateTimeEditD0Ev(void* qthis);
  // proto:  void QDateTimeEdit::clearMinimumDate();
extern void _ZN13QDateTimeEdit16clearMinimumDateEv(void* qthis);
  // proto:  void QDateTimeEdit::setDateRange(const QDate & min, const QDate & max);
extern void _ZN13QDateTimeEdit12setDateRangeERK5QDateS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QDateTimeEdit::setMinimumTime(const QTime & min);
extern void _ZN13QDateTimeEdit14setMinimumTimeERK5QTime(void* qthis, void* arg0);
  // proto:  QTime QDateTimeEdit::time();
extern void _ZNK13QDateTimeEdit4timeEv(void* qthis);
  // proto:  int QDateTimeEdit::currentSectionIndex();
extern void _ZNK13QDateTimeEdit19currentSectionIndexEv(void* qthis);
  // proto:  bool QDateTimeEdit::event(QEvent * event);
extern void _ZN13QDateTimeEdit5eventEP6QEvent(void* qthis, void* arg0);
  // proto:  void QDateTimeEdit::setDateTime(const QDateTime & dateTime);
extern void _ZN13QDateTimeEdit11setDateTimeERK9QDateTime(void* qthis, void* arg0);
  // proto:  void QDateTimeEdit::setCalendarWidget(QCalendarWidget * calendarWidget);
extern void _ZN13QDateTimeEdit17setCalendarWidgetEP15QCalendarWidget(void* qthis, void* arg0);
  // proto:  void QDateTimeEdit::setDateTimeRange(const QDateTime & min, const QDateTime & max);
extern void _ZN13QDateTimeEdit16setDateTimeRangeERK9QDateTimeS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QDateTimeEdit::setTime(const QTime & time);
extern void _ZN13QDateTimeEdit7setTimeERK5QTime(void* qthis, void* arg0);
  // proto:  QDateTime QDateTimeEdit::dateTime();
extern void _ZNK13QDateTimeEdit8dateTimeEv(void* qthis);
  // proto:  void QDateTimeEdit::setMaximumTime(const QTime & max);
extern void _ZN13QDateTimeEdit14setMaximumTimeERK5QTime(void* qthis, void* arg0);
  // proto:  void QDateTimeEdit::setMinimumDateTime(const QDateTime & dt);
extern void _ZN13QDateTimeEdit18setMinimumDateTimeERK9QDateTime(void* qthis, void* arg0);
  // proto:  void QDateTimeEdit::clearMaximumTime();
extern void _ZN13QDateTimeEdit16clearMaximumTimeEv(void* qthis);
  // proto:  bool QDateTimeEdit::calendarPopup();
extern void _ZNK13QDateTimeEdit13calendarPopupEv(void* qthis);
  // proto:  void QDateTimeEdit::setMaximumDateTime(const QDateTime & dt);
extern void _ZN13QDateTimeEdit18setMaximumDateTimeERK9QDateTime(void* qthis, void* arg0);
  // proto:  void QDateTimeEdit::clearMaximumDateTime();
extern void _ZN13QDateTimeEdit20clearMaximumDateTimeEv(void* qthis);
  // proto:  void QDateTimeEdit::QDateTimeEdit(QWidget * parent);
extern void* dector_ZN13QDateTimeEditC1EP7QWidget(void* arg0);
extern void _ZN13QDateTimeEditC1EP7QWidget(void* qthis, void* arg0);
  // proto:  QDateTime QDateTimeEdit::minimumDateTime();
extern void _ZNK13QDateTimeEdit15minimumDateTimeEv(void* qthis);
  // proto:  QCalendarWidget * QDateTimeEdit::calendarWidget();
extern void _ZNK13QDateTimeEdit14calendarWidgetEv(void* qthis);
  // proto:  void QDateTimeEdit::setDisplayFormat(const QString & format);
extern void _ZN13QDateTimeEdit16setDisplayFormatERK7QString(void* qthis, void* arg0);
  // proto:  void QDateTimeEdit::clearMaximumDate();
extern void _ZN13QDateTimeEdit16clearMaximumDateEv(void* qthis);
  // proto:  void QDateTimeEdit::setCalendarPopup(bool enable);
extern void _ZN13QDateTimeEdit16setCalendarPopupEb(void* qthis, bool arg0);
  // proto:  void QDateTimeEdit::stepBy(int steps);
extern void _ZN13QDateTimeEdit6stepByEi(void* qthis, int32_t arg0);
  // proto:  void QDateTimeEdit::QDateTimeEdit(const QDateTime & dt, QWidget * parent);
extern void* dector_ZN13QDateTimeEditC1ERK9QDateTimeP7QWidget(void* arg0, void* arg1);
extern void _ZN13QDateTimeEditC1ERK9QDateTimeP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  QString QDateTimeEdit::displayFormat();
extern void _ZNK13QDateTimeEdit13displayFormatEv(void* qthis);
  // proto:  QTime QDateTimeEdit::minimumTime();
extern void _ZNK13QDateTimeEdit11minimumTimeEv(void* qthis);
  // proto:  QSize QDateTimeEdit::sizeHint();
extern void _ZNK13QDateTimeEdit8sizeHintEv(void* qthis);
  // proto:  int QDateTimeEdit::sectionCount();
extern void _ZNK13QDateTimeEdit12sectionCountEv(void* qthis);
  // proto:  void QDateTimeEdit::setCurrentSectionIndex(int index);
extern void _ZN13QDateTimeEdit22setCurrentSectionIndexEi(void* qthis, int32_t arg0);
  // proto:  void QDateTimeEdit::clearMinimumTime();
extern void _ZN13QDateTimeEdit16clearMinimumTimeEv(void* qthis);
  // proto:  void QDateTimeEdit::setTimeRange(const QTime & min, const QTime & max);
extern void _ZN13QDateTimeEdit12setTimeRangeERK5QTimeS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QDate QDateTimeEdit::minimumDate();
extern void _ZNK13QDateTimeEdit11minimumDateEv(void* qthis);
  // proto:  QDate QDateTimeEdit::maximumDate();
extern void _ZNK13QDateTimeEdit11maximumDateEv(void* qthis);
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

// class sizeof(QTimeEdit)=1
type QTimeEdit struct {
  /*qbase*/ QDateTimeEdit;
  qclsinst unsafe.Pointer /* *C.void */;
//  _userTimeChanged QTimeEdit_userTimeChanged_signal;
}

// class sizeof(QDateEdit)=1
type QDateEdit struct {
  /*qbase*/ QDateTimeEdit;
  qclsinst unsafe.Pointer /* *C.void */;
//  _userDateChanged QDateEdit_userDateChanged_signal;
}

// class sizeof(QDateTimeEdit)=1
type QDateTimeEdit struct {
  /*qbase*/ QAbstractSpinBox;
  qclsinst unsafe.Pointer /* *C.void */;
//  _dateChanged QDateTimeEdit_dateChanged_signal;
//  _timeChanged QDateTimeEdit_timeChanged_signal;
//  _dateTimeChanged QDateTimeEdit_dateTimeChanged_signal;
}

  // proto:  void QTimeEdit::QTimeEdit(QWidget * parent);
func NewQTimeEdit(args ...interface{}) QTimeEdit {
  return QTimeEdit{}
}

  // proto:  const QMetaObject * QTimeEdit::metaObject();
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QTimeEdit10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimeEdit", "metaObject", args)
  }

}

  // proto:  void QTimeEdit::~QTimeEdit();
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

  // proto:  void QDateEdit::QDateEdit(const QDate & date, QWidget * parent);
func NewQDateEdit(args ...interface{}) QDateEdit {
  return QDateEdit{}
}

  // proto:  const QMetaObject * QDateEdit::metaObject();
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QDateEdit10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateEdit", "metaObject", args)
  }

}

  // proto:  void QDateEdit::~QDateEdit();
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

  // proto:  void QDateTimeEdit::QDateTimeEdit(const QDateTimeEdit & );
func NewQDateTimeEdit(args ...interface{}) QDateTimeEdit {
  return QDateTimeEdit{}
}

  // proto:  QDate QDateTimeEdit::date();
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
    // invoke: QDate date()
    C._ZNK13QDateTimeEdit4dateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "date", args)
  }

}

  // proto:  void QDateTimeEdit::clearMinimumDateTime();
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
    // invoke: void clearMinimumDateTime()
    C._ZN13QDateTimeEdit20clearMinimumDateTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMinimumDateTime", args)
  }

}

  // proto:  void QDateTimeEdit::clear();
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
    // invoke: void clear()
    C._ZN13QDateTimeEdit5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clear", args)
  }

}

  // proto:  void QDateTimeEdit::setDate(const QDate & date);
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
    // invoke: void setDate(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit7setDateERK5QDate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDate", args)
  }

}

  // proto:  QDateTime QDateTimeEdit::maximumDateTime();
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
    // invoke: QDateTime maximumDateTime()
    C._ZNK13QDateTimeEdit15maximumDateTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "maximumDateTime", args)
  }

}

  // proto:  void QDateTimeEdit::setMinimumDate(const QDate & min);
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
    // invoke: void setMinimumDate(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit14setMinimumDateERK5QDate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMinimumDate", args)
  }

}

  // proto:  void QDateTimeEdit::setMaximumDate(const QDate & max);
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
    // invoke: void setMaximumDate(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit14setMaximumDateERK5QDate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMaximumDate", args)
  }

}

  // proto:  QTime QDateTimeEdit::maximumTime();
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
    // invoke: QTime maximumTime()
    C._ZNK13QDateTimeEdit11maximumTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "maximumTime", args)
  }

}

  // proto:  const QMetaObject * QDateTimeEdit::metaObject();
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QDateTimeEdit10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "metaObject", args)
  }

}

  // proto:  void QDateTimeEdit::~QDateTimeEdit();
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

  // proto:  void QDateTimeEdit::clearMinimumDate();
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
    // invoke: void clearMinimumDate()
    C._ZN13QDateTimeEdit16clearMinimumDateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMinimumDate", args)
  }

}

  // proto:  void QDateTimeEdit::setDateRange(const QDate & min, const QDate & max);
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
    // invoke: void setDateRange(const class QDate &, const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QDate).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN13QDateTimeEdit12setDateRangeERK5QDateS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDateRange", args)
  }

}

  // proto:  void QDateTimeEdit::setMinimumTime(const QTime & min);
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
    // invoke: void setMinimumTime(const class QTime &)
    var arg0 = args[0].(QTime).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit14setMinimumTimeERK5QTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMinimumTime", args)
  }

}

  // proto:  QTime QDateTimeEdit::time();
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
    // invoke: QTime time()
    C._ZNK13QDateTimeEdit4timeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "time", args)
  }

}

  // proto:  int QDateTimeEdit::currentSectionIndex();
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
    // invoke: int currentSectionIndex()
    C._ZNK13QDateTimeEdit19currentSectionIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "currentSectionIndex", args)
  }

}

  // proto:  bool QDateTimeEdit::event(QEvent * event);
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
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "event", args)
  }

}

  // proto:  void QDateTimeEdit::setDateTime(const QDateTime & dateTime);
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
    // invoke: void setDateTime(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit11setDateTimeERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDateTime", args)
  }

}

  // proto:  void QDateTimeEdit::setCalendarWidget(QCalendarWidget * calendarWidget);
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
    // invoke: void setCalendarWidget(class QCalendarWidget *)
    var arg0 = args[0].(QCalendarWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit17setCalendarWidgetEP15QCalendarWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setCalendarWidget", args)
  }

}

  // proto:  void QDateTimeEdit::setDateTimeRange(const QDateTime & min, const QDateTime & max);
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
    // invoke: void setDateTimeRange(const class QDateTime &, const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QDateTime).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN13QDateTimeEdit16setDateTimeRangeERK9QDateTimeS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDateTimeRange", args)
  }

}

  // proto:  void QDateTimeEdit::setTime(const QTime & time);
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
    // invoke: void setTime(const class QTime &)
    var arg0 = args[0].(QTime).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit7setTimeERK5QTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setTime", args)
  }

}

  // proto:  QDateTime QDateTimeEdit::dateTime();
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
    // invoke: QDateTime dateTime()
    C._ZNK13QDateTimeEdit8dateTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "dateTime", args)
  }

}

  // proto:  void QDateTimeEdit::setMaximumTime(const QTime & max);
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
    // invoke: void setMaximumTime(const class QTime &)
    var arg0 = args[0].(QTime).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit14setMaximumTimeERK5QTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMaximumTime", args)
  }

}

  // proto:  void QDateTimeEdit::setMinimumDateTime(const QDateTime & dt);
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
    // invoke: void setMinimumDateTime(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit18setMinimumDateTimeERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMinimumDateTime", args)
  }

}

  // proto:  void QDateTimeEdit::clearMaximumTime();
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
    // invoke: void clearMaximumTime()
    C._ZN13QDateTimeEdit16clearMaximumTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMaximumTime", args)
  }

}

  // proto:  bool QDateTimeEdit::calendarPopup();
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
    // invoke: bool calendarPopup()
    C._ZNK13QDateTimeEdit13calendarPopupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "calendarPopup", args)
  }

}

  // proto:  void QDateTimeEdit::setMaximumDateTime(const QDateTime & dt);
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
    // invoke: void setMaximumDateTime(const class QDateTime &)
    var arg0 = args[0].(QDateTime).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit18setMaximumDateTimeERK9QDateTime(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setMaximumDateTime", args)
  }

}

  // proto:  void QDateTimeEdit::clearMaximumDateTime();
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
    // invoke: void clearMaximumDateTime()
    C._ZN13QDateTimeEdit20clearMaximumDateTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMaximumDateTime", args)
  }

}

  // proto:  QDateTime QDateTimeEdit::minimumDateTime();
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
    // invoke: QDateTime minimumDateTime()
    C._ZNK13QDateTimeEdit15minimumDateTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "minimumDateTime", args)
  }

}

  // proto:  QCalendarWidget * QDateTimeEdit::calendarWidget();
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
    // invoke: QCalendarWidget * calendarWidget()
    C._ZNK13QDateTimeEdit14calendarWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "calendarWidget", args)
  }

}

  // proto:  void QDateTimeEdit::setDisplayFormat(const QString & format);
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
    // invoke: void setDisplayFormat(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit16setDisplayFormatERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setDisplayFormat", args)
  }

}

  // proto:  void QDateTimeEdit::clearMaximumDate();
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
    // invoke: void clearMaximumDate()
    C._ZN13QDateTimeEdit16clearMaximumDateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMaximumDate", args)
  }

}

  // proto:  void QDateTimeEdit::setCalendarPopup(bool enable);
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
    // invoke: void setCalendarPopup(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit16setCalendarPopupEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setCalendarPopup", args)
  }

}

  // proto:  void QDateTimeEdit::stepBy(int steps);
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
    // invoke: void stepBy(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit6stepByEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "stepBy", args)
  }

}

  // proto:  QString QDateTimeEdit::displayFormat();
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
    // invoke: QString displayFormat()
    C._ZNK13QDateTimeEdit13displayFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "displayFormat", args)
  }

}

  // proto:  QTime QDateTimeEdit::minimumTime();
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
    // invoke: QTime minimumTime()
    C._ZNK13QDateTimeEdit11minimumTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "minimumTime", args)
  }

}

  // proto:  QSize QDateTimeEdit::sizeHint();
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
    // invoke: QSize sizeHint()
    C._ZNK13QDateTimeEdit8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "sizeHint", args)
  }

}

  // proto:  int QDateTimeEdit::sectionCount();
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
    // invoke: int sectionCount()
    C._ZNK13QDateTimeEdit12sectionCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "sectionCount", args)
  }

}

  // proto:  void QDateTimeEdit::setCurrentSectionIndex(int index);
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
    // invoke: void setCurrentSectionIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN13QDateTimeEdit22setCurrentSectionIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setCurrentSectionIndex", args)
  }

}

  // proto:  void QDateTimeEdit::clearMinimumTime();
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
    // invoke: void clearMinimumTime()
    C._ZN13QDateTimeEdit16clearMinimumTimeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "clearMinimumTime", args)
  }

}

  // proto:  void QDateTimeEdit::setTimeRange(const QTime & min, const QTime & max);
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
    // invoke: void setTimeRange(const class QTime &, const class QTime &)
    var arg0 = args[0].(QTime).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTime).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN13QDateTimeEdit12setTimeRangeERK5QTimeS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "setTimeRange", args)
  }

}

  // proto:  QDate QDateTimeEdit::minimumDate();
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
    // invoke: QDate minimumDate()
    C._ZNK13QDateTimeEdit11minimumDateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "minimumDate", args)
  }

}

  // proto:  QDate QDateTimeEdit::maximumDate();
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
    // invoke: QDate maximumDate()
    C._ZNK13QDateTimeEdit11maximumDateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDateTimeEdit", "maximumDate", args)
  }

}

// <= body block end

