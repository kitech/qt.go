package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qcalendarwidget.h
// dst-file: /src/widgets/qcalendarwidget.go
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
  // proto:  void QCalendarWidget::showPreviousYear();
extern void _ZN15QCalendarWidget16showPreviousYearEv(void* qthis);
  // proto:  QDate QCalendarWidget::maximumDate();
extern void _ZNK15QCalendarWidget11maximumDateEv(void* qthis);
  // proto:  void QCalendarWidget::showPreviousMonth();
extern void _ZN15QCalendarWidget17showPreviousMonthEv(void* qthis);
  // proto:  void QCalendarWidget::showSelectedDate();
extern void _ZN15QCalendarWidget16showSelectedDateEv(void* qthis);
  // proto:  QSize QCalendarWidget::minimumSizeHint();
extern void _ZNK15QCalendarWidget15minimumSizeHintEv(void* qthis);
  // proto:  void QCalendarWidget::setDateEditAcceptDelay(int delay);
extern void _ZN15QCalendarWidget22setDateEditAcceptDelayEi(void* qthis, int arg0);
  // proto:  void QCalendarWidget::setGridVisible(bool show);
extern void _ZN15QCalendarWidget14setGridVisibleEb(void* qthis, bool arg0);
  // proto:  void QCalendarWidget::~QCalendarWidget();
extern void _ZN15QCalendarWidgetD0Ev(void* qthis);
  // proto:  QSize QCalendarWidget::sizeHint();
extern void _ZNK15QCalendarWidget8sizeHintEv(void* qthis);
  // proto:  int QCalendarWidget::monthShown();
extern void _ZNK15QCalendarWidget10monthShownEv(void* qthis);
  // proto:  void QCalendarWidget::setSelectedDate(const QDate & date);
extern void _ZN15QCalendarWidget15setSelectedDateERK5QDate(void* qthis, void* arg0);
  // proto:  void QCalendarWidget::QCalendarWidget(QWidget * parent);
extern void* dector_ZN15QCalendarWidgetC1EP7QWidget(void* arg0);
extern void _ZN15QCalendarWidgetC1EP7QWidget(void* qthis, void* arg0);
  // proto:  const QMetaObject * QCalendarWidget::metaObject();
extern void _ZNK15QCalendarWidget10metaObjectEv(void* qthis);
  // proto:  void QCalendarWidget::setNavigationBarVisible(bool visible);
extern void _ZN15QCalendarWidget23setNavigationBarVisibleEb(void* qthis, bool arg0);
  // proto:  bool QCalendarWidget::isNavigationBarVisible();
extern void _ZNK15QCalendarWidget22isNavigationBarVisibleEv(void* qthis);
  // proto:  QTextCharFormat QCalendarWidget::dateTextFormat(const QDate & date);
extern void _ZNK15QCalendarWidget14dateTextFormatERK5QDate(void* qthis, void* arg0);
  // proto:  void QCalendarWidget::setMinimumDate(const QDate & date);
extern void _ZN15QCalendarWidget14setMinimumDateERK5QDate(void* qthis, void* arg0);
  // proto:  int QCalendarWidget::dateEditAcceptDelay();
extern void _ZNK15QCalendarWidget19dateEditAcceptDelayEv(void* qthis);
  // proto:  QDate QCalendarWidget::minimumDate();
extern void _ZNK15QCalendarWidget11minimumDateEv(void* qthis);
  // proto:  void QCalendarWidget::QCalendarWidget(const QCalendarWidget & );
extern void* dector_ZN15QCalendarWidgetC1ERKS_(void* arg0);
extern void _ZN15QCalendarWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QCalendarWidget::isDateEditEnabled();
extern void _ZNK15QCalendarWidget17isDateEditEnabledEv(void* qthis);
  // proto:  QMap<QDate, QTextCharFormat> QCalendarWidget::dateTextFormat();
extern void _ZNK15QCalendarWidget14dateTextFormatEv(void* qthis);
  // proto:  void QCalendarWidget::setDateEditEnabled(bool enable);
extern void _ZN15QCalendarWidget18setDateEditEnabledEb(void* qthis, bool arg0);
  // proto:  void QCalendarWidget::setDateTextFormat(const QDate & date, const QTextCharFormat & format);
extern void _ZN15QCalendarWidget17setDateTextFormatERK5QDateRK15QTextCharFormat(void* qthis, void* arg0, void* arg1);
  // proto:  void QCalendarWidget::showNextMonth();
extern void _ZN15QCalendarWidget13showNextMonthEv(void* qthis);
  // proto:  void QCalendarWidget::setDateRange(const QDate & min, const QDate & max);
extern void _ZN15QCalendarWidget12setDateRangeERK5QDateS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QDate QCalendarWidget::selectedDate();
extern void _ZNK15QCalendarWidget12selectedDateEv(void* qthis);
  // proto:  void QCalendarWidget::setHeaderTextFormat(const QTextCharFormat & format);
extern void _ZN15QCalendarWidget19setHeaderTextFormatERK15QTextCharFormat(void* qthis, void* arg0);
  // proto:  bool QCalendarWidget::isGridVisible();
extern void _ZNK15QCalendarWidget13isGridVisibleEv(void* qthis);
  // proto:  int QCalendarWidget::yearShown();
extern void _ZNK15QCalendarWidget9yearShownEv(void* qthis);
  // proto:  void QCalendarWidget::setMaximumDate(const QDate & date);
extern void _ZN15QCalendarWidget14setMaximumDateERK5QDate(void* qthis, void* arg0);
  // proto:  QTextCharFormat QCalendarWidget::headerTextFormat();
extern void _ZNK15QCalendarWidget16headerTextFormatEv(void* qthis);
  // proto:  void QCalendarWidget::setCurrentPage(int year, int month);
extern void _ZN15QCalendarWidget14setCurrentPageEii(void* qthis, int arg0, int arg1);
  // proto:  void QCalendarWidget::showToday();
extern void _ZN15QCalendarWidget9showTodayEv(void* qthis);
  // proto:  void QCalendarWidget::showNextYear();
extern void _ZN15QCalendarWidget12showNextYearEv(void* qthis);
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

// class sizeof(QCalendarWidget)=1
type QCalendarWidget struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _activated QCalendarWidget_activated_signal;
//  _clicked QCalendarWidget_clicked_signal;
//  _currentPageChanged QCalendarWidget_currentPageChanged_signal;
//  _selectionChanged QCalendarWidget_selectionChanged_signal;
}

  // proto:  void QCalendarWidget::showPreviousYear();
func (this *QCalendarWidget) showPreviousYear(args ...interface{}) () {
  // showPreviousYear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget16showPreviousYearEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showPreviousYear", args)
  }

}

  // proto:  QDate QCalendarWidget::maximumDate();
func (this *QCalendarWidget) maximumDate(args ...interface{}) () {
  // maximumDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget11maximumDateEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "maximumDate", args)
  }

}

  // proto:  void QCalendarWidget::showPreviousMonth();
func (this *QCalendarWidget) showPreviousMonth(args ...interface{}) () {
  // showPreviousMonth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget17showPreviousMonthEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showPreviousMonth", args)
  }

}

  // proto:  void QCalendarWidget::showSelectedDate();
func (this *QCalendarWidget) showSelectedDate(args ...interface{}) () {
  // showSelectedDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget16showSelectedDateEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showSelectedDate", args)
  }

}

  // proto:  QSize QCalendarWidget::minimumSizeHint();
func (this *QCalendarWidget) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "minimumSizeHint", args)
  }

}

  // proto:  void QCalendarWidget::setDateEditAcceptDelay(int delay);
func (this *QCalendarWidget) setDateEditAcceptDelay(args ...interface{}) () {
  // setDateEditAcceptDelay(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget22setDateEditAcceptDelayEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateEditAcceptDelay", args)
  }

}

  // proto:  void QCalendarWidget::setGridVisible(bool show);
func (this *QCalendarWidget) setGridVisible(args ...interface{}) () {
  // setGridVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget14setGridVisibleEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setGridVisible", args)
  }

}

  // proto:  void QCalendarWidget::~QCalendarWidget();
func (this *QCalendarWidget) FreeQCalendarWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCalendarWidget", "~QCalendarWidget", args)
  }

}

  // proto:  QSize QCalendarWidget::sizeHint();
func (this *QCalendarWidget) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget8sizeHintEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "sizeHint", args)
  }

}

  // proto:  int QCalendarWidget::monthShown();
func (this *QCalendarWidget) monthShown(args ...interface{}) () {
  // monthShown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget10monthShownEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "monthShown", args)
  }

}

  // proto:  void QCalendarWidget::setSelectedDate(const QDate & date);
func (this *QCalendarWidget) setSelectedDate(args ...interface{}) () {
  // setSelectedDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget15setSelectedDateERK5QDate
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setSelectedDate", args)
  }

}

  // proto:  void QCalendarWidget::QCalendarWidget(QWidget * parent);
func NewQCalendarWidget(args ...interface{}) QCalendarWidget {
  return QCalendarWidget{}
}

  // proto:  const QMetaObject * QCalendarWidget::metaObject();
func (this *QCalendarWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget10metaObjectEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "metaObject", args)
  }

}

  // proto:  void QCalendarWidget::setNavigationBarVisible(bool visible);
func (this *QCalendarWidget) setNavigationBarVisible(args ...interface{}) () {
  // setNavigationBarVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget23setNavigationBarVisibleEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setNavigationBarVisible", args)
  }

}

  // proto:  bool QCalendarWidget::isNavigationBarVisible();
func (this *QCalendarWidget) isNavigationBarVisible(args ...interface{}) () {
  // isNavigationBarVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget22isNavigationBarVisibleEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "isNavigationBarVisible", args)
  }

}

  // proto:  QTextCharFormat QCalendarWidget::dateTextFormat(const QDate & date);
func (this *QCalendarWidget) dateTextFormat(args ...interface{}) () {
  // dateTextFormat(const class QDate &)
  // dateTextFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget14dateTextFormatERK5QDate
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK15QCalendarWidget14dateTextFormatEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "dateTextFormat", args)
  }

}

  // proto:  void QCalendarWidget::setMinimumDate(const QDate & date);
func (this *QCalendarWidget) setMinimumDate(args ...interface{}) () {
  // setMinimumDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget14setMinimumDateERK5QDate
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setMinimumDate", args)
  }

}

  // proto:  int QCalendarWidget::dateEditAcceptDelay();
func (this *QCalendarWidget) dateEditAcceptDelay(args ...interface{}) () {
  // dateEditAcceptDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget19dateEditAcceptDelayEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "dateEditAcceptDelay", args)
  }

}

  // proto:  QDate QCalendarWidget::minimumDate();
func (this *QCalendarWidget) minimumDate(args ...interface{}) () {
  // minimumDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget11minimumDateEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "minimumDate", args)
  }

}

  // proto:  bool QCalendarWidget::isDateEditEnabled();
func (this *QCalendarWidget) isDateEditEnabled(args ...interface{}) () {
  // isDateEditEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget17isDateEditEnabledEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "isDateEditEnabled", args)
  }

}

  // proto:  void QCalendarWidget::setDateEditEnabled(bool enable);
func (this *QCalendarWidget) setDateEditEnabled(args ...interface{}) () {
  // setDateEditEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget18setDateEditEnabledEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateEditEnabled", args)
  }

}

  // proto:  void QCalendarWidget::setDateTextFormat(const QDate & date, const QTextCharFormat & format);
func (this *QCalendarWidget) setDateTextFormat(args ...interface{}) () {
  // setDateTextFormat(const class QDate &, const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"
  vtys[0][1] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget17setDateTextFormatERK5QDateRK15QTextCharFormat
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateTextFormat", args)
  }

}

  // proto:  void QCalendarWidget::showNextMonth();
func (this *QCalendarWidget) showNextMonth(args ...interface{}) () {
  // showNextMonth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget13showNextMonthEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showNextMonth", args)
  }

}

  // proto:  void QCalendarWidget::setDateRange(const QDate & min, const QDate & max);
func (this *QCalendarWidget) setDateRange(args ...interface{}) () {
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
    // invoke: _ZN15QCalendarWidget12setDateRangeERK5QDateS2_
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QDate).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateRange", args)
  }

}

  // proto:  QDate QCalendarWidget::selectedDate();
func (this *QCalendarWidget) selectedDate(args ...interface{}) () {
  // selectedDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget12selectedDateEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "selectedDate", args)
  }

}

  // proto:  void QCalendarWidget::setHeaderTextFormat(const QTextCharFormat & format);
func (this *QCalendarWidget) setHeaderTextFormat(args ...interface{}) () {
  // setHeaderTextFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget19setHeaderTextFormatERK15QTextCharFormat
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setHeaderTextFormat", args)
  }

}

  // proto:  bool QCalendarWidget::isGridVisible();
func (this *QCalendarWidget) isGridVisible(args ...interface{}) () {
  // isGridVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget13isGridVisibleEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "isGridVisible", args)
  }

}

  // proto:  int QCalendarWidget::yearShown();
func (this *QCalendarWidget) yearShown(args ...interface{}) () {
  // yearShown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget9yearShownEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "yearShown", args)
  }

}

  // proto:  void QCalendarWidget::setMaximumDate(const QDate & date);
func (this *QCalendarWidget) setMaximumDate(args ...interface{}) () {
  // setMaximumDate(const class QDate &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDate{}) // "const QDate &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget14setMaximumDateERK5QDate
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setMaximumDate", args)
  }

}

  // proto:  QTextCharFormat QCalendarWidget::headerTextFormat();
func (this *QCalendarWidget) headerTextFormat(args ...interface{}) () {
  // headerTextFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget16headerTextFormatEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "headerTextFormat", args)
  }

}

  // proto:  void QCalendarWidget::setCurrentPage(int year, int month);
func (this *QCalendarWidget) setCurrentPage(args ...interface{}) () {
  // setCurrentPage(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget14setCurrentPageEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setCurrentPage", args)
  }

}

  // proto:  void QCalendarWidget::showToday();
func (this *QCalendarWidget) showToday(args ...interface{}) () {
  // showToday()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget9showTodayEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showToday", args)
  }

}

  // proto:  void QCalendarWidget::showNextYear();
func (this *QCalendarWidget) showNextYear(args ...interface{}) () {
  // showNextYear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidget12showNextYearEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showNextYear", args)
  }

}

// <= body block end

