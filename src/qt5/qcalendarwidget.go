package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QCalendarWidget::HorizontalHeaderFormat QCalendarWidget::horizontalHeaderFormat();
extern void C_ZNK15QCalendarWidget22horizontalHeaderFormatEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setMaximumDate(const QDate & date);
extern void C_ZN15QCalendarWidget14setMaximumDateERK5QDate(void* qthis, void* arg0); // 4
  // proto:  QTextCharFormat QCalendarWidget::headerTextFormat();
extern void C_ZNK15QCalendarWidget16headerTextFormatEv(void* qthis); // 4
  // proto:  void QCalendarWidget::showSelectedDate();
extern void C_ZN15QCalendarWidget16showSelectedDateEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setGridVisible(bool show);
extern void C_ZN15QCalendarWidget14setGridVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QCalendarWidget::setDateEditAcceptDelay(int delay);
extern void C_ZN15QCalendarWidget22setDateEditAcceptDelayEi(void* qthis, int32_t arg0); // 4
  // proto:  void QCalendarWidget::setDateRange(const QDate & min, const QDate & max);
extern void C_ZN15QCalendarWidget12setDateRangeERK5QDateS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QCalendarWidget::isDateEditEnabled();
extern void C_ZNK15QCalendarWidget17isDateEditEnabledEv(void* qthis); // 4
  // proto:  void QCalendarWidget::~QCalendarWidget();
extern void C_ZN15QCalendarWidgetD2Ev(void* qthis); // 4
  // proto:  QTextCharFormat QCalendarWidget::dateTextFormat(const QDate & date);
extern void C_ZNK15QCalendarWidget14dateTextFormatERK5QDate(void* qthis, void* arg0); // 4
  // proto:  QMap<QDate, QTextCharFormat> QCalendarWidget::dateTextFormat();
extern void C_ZNK15QCalendarWidget14dateTextFormatEv(void* qthis); // 4
  // proto:  bool QCalendarWidget::isNavigationBarVisible();
extern void C_ZNK15QCalendarWidget22isNavigationBarVisibleEv(void* qthis); // 4
  // proto:  Qt::DayOfWeek QCalendarWidget::firstDayOfWeek();
extern void C_ZNK15QCalendarWidget14firstDayOfWeekEv(void* qthis); // 4
  // proto:  int QCalendarWidget::monthShown();
extern void C_ZNK15QCalendarWidget10monthShownEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setNavigationBarVisible(bool visible);
extern void C_ZN15QCalendarWidget23setNavigationBarVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QCalendarWidget::showNextYear();
extern void C_ZN15QCalendarWidget12showNextYearEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setMinimumDate(const QDate & date);
extern void C_ZN15QCalendarWidget14setMinimumDateERK5QDate(void* qthis, void* arg0); // 4
  // proto:  int QCalendarWidget::yearShown();
extern void C_ZNK15QCalendarWidget9yearShownEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setHeaderTextFormat(const QTextCharFormat & format);
extern void C_ZN15QCalendarWidget19setHeaderTextFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  QCalendarWidget::VerticalHeaderFormat QCalendarWidget::verticalHeaderFormat();
extern void C_ZNK15QCalendarWidget20verticalHeaderFormatEv(void* qthis); // 4
  // proto:  QDate QCalendarWidget::minimumDate();
extern void C_ZNK15QCalendarWidget11minimumDateEv(void* qthis); // 4
  // proto:  QSize QCalendarWidget::sizeHint();
extern void C_ZNK15QCalendarWidget8sizeHintEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setSelectedDate(const QDate & date);
extern void C_ZN15QCalendarWidget15setSelectedDateERK5QDate(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QCalendarWidget::metaObject();
extern void C_ZNK15QCalendarWidget10metaObjectEv(void* qthis); // 4
  // proto:  QCalendarWidget::SelectionMode QCalendarWidget::selectionMode();
extern void C_ZNK15QCalendarWidget13selectionModeEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setCurrentPage(int year, int month);
extern void C_ZN15QCalendarWidget14setCurrentPageEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QCalendarWidget::setDateTextFormat(const QDate & date, const QTextCharFormat & format);
extern void C_ZN15QCalendarWidget17setDateTextFormatERK5QDateRK15QTextCharFormat(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QDate QCalendarWidget::selectedDate();
extern void C_ZNK15QCalendarWidget12selectedDateEv(void* qthis); // 4
  // proto:  QDate QCalendarWidget::maximumDate();
extern void C_ZNK15QCalendarWidget11maximumDateEv(void* qthis); // 4
  // proto:  QSize QCalendarWidget::minimumSizeHint();
extern void C_ZNK15QCalendarWidget15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QCalendarWidget::showPreviousMonth();
extern void C_ZN15QCalendarWidget17showPreviousMonthEv(void* qthis); // 4
  // proto:  int QCalendarWidget::dateEditAcceptDelay();
extern void C_ZNK15QCalendarWidget19dateEditAcceptDelayEv(void* qthis); // 4
  // proto:  void QCalendarWidget::showPreviousYear();
extern void C_ZN15QCalendarWidget16showPreviousYearEv(void* qthis); // 4
  // proto:  bool QCalendarWidget::isGridVisible();
extern void C_ZNK15QCalendarWidget13isGridVisibleEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setDateEditEnabled(bool enable);
extern void C_ZN15QCalendarWidget18setDateEditEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QCalendarWidget::QCalendarWidget(QWidget * parent);
extern void* C_ZN15QCalendarWidgetC2EP7QWidget(void* arg0); // 3
  // proto:  void QCalendarWidget::showNextMonth();
extern void C_ZN15QCalendarWidget13showNextMonthEv(void* qthis); // 4
  // proto:  void QCalendarWidget::showToday();
extern void C_ZN15QCalendarWidget9showTodayEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _activated QCalendarWidget_activated_signal;
//  _clicked QCalendarWidget_clicked_signal;
//  _currentPageChanged QCalendarWidget_currentPageChanged_signal;
//  _selectionChanged QCalendarWidget_selectionChanged_signal;
}

// horizontalHeaderFormat()
func (this *QCalendarWidget) horizontalHeaderFormat(args ...interface{}) () {
  // horizontalHeaderFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget22horizontalHeaderFormatEv
    // invoke: QCalendarWidget::HorizontalHeaderFormat horizontalHeaderFormat()
    C.C_ZNK15QCalendarWidget22horizontalHeaderFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "horizontalHeaderFormat", args)
  }

}

// setMaximumDate(const class QDate &)
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
    // invoke: void setMaximumDate(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QCalendarWidget14setMaximumDateERK5QDate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setMaximumDate", args)
  }

}

// headerTextFormat()
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
    // invoke: QTextCharFormat headerTextFormat()
    var ret = C.C_ZNK15QCalendarWidget16headerTextFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "headerTextFormat", args)
  }

}

// showSelectedDate()
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
    // invoke: void showSelectedDate()
    C.C_ZN15QCalendarWidget16showSelectedDateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showSelectedDate", args)
  }

}

// setGridVisible(_Bool)
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
    // invoke: void setGridVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QCalendarWidget14setGridVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setGridVisible", args)
  }

}

// setDateEditAcceptDelay(int)
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
    // invoke: void setDateEditAcceptDelay(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QCalendarWidget22setDateEditAcceptDelayEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateEditAcceptDelay", args)
  }

}

// setDateRange(const class QDate &, const class QDate &)
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
    // invoke: void setDateRange(const class QDate &, const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QDate).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QCalendarWidget12setDateRangeERK5QDateS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateRange", args)
  }

}

// isDateEditEnabled()
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
    // invoke: bool isDateEditEnabled()
    var ret = C.C_ZNK15QCalendarWidget17isDateEditEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "isDateEditEnabled", args)
  }

}

// ~QCalendarWidget()
func (this *QCalendarWidget) FreeQCalendarWidget(args ...interface{}) () {
  // ~QCalendarWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidgetD0Ev
    // invoke: void ~QCalendarWidget()
    C.C_ZN15QCalendarWidgetD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "~QCalendarWidget", args)
  }

}

// dateTextFormat(const class QDate &)
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
    // invoke: QTextCharFormat dateTextFormat(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK15QCalendarWidget14dateTextFormatERK5QDate(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK15QCalendarWidget14dateTextFormatEv
    // invoke: QMap<QDate, QTextCharFormat> dateTextFormat()
    C.C_ZNK15QCalendarWidget14dateTextFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "dateTextFormat", args)
  }

}

// isNavigationBarVisible()
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
    // invoke: bool isNavigationBarVisible()
    var ret = C.C_ZNK15QCalendarWidget22isNavigationBarVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "isNavigationBarVisible", args)
  }

}

// firstDayOfWeek()
func (this *QCalendarWidget) firstDayOfWeek(args ...interface{}) () {
  // firstDayOfWeek()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget14firstDayOfWeekEv
    // invoke: Qt::DayOfWeek firstDayOfWeek()
    C.C_ZNK15QCalendarWidget14firstDayOfWeekEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "firstDayOfWeek", args)
  }

}

// monthShown()
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
    // invoke: int monthShown()
    var ret = C.C_ZNK15QCalendarWidget10monthShownEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "monthShown", args)
  }

}

// setNavigationBarVisible(_Bool)
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
    // invoke: void setNavigationBarVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QCalendarWidget23setNavigationBarVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setNavigationBarVisible", args)
  }

}

// showNextYear()
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
    // invoke: void showNextYear()
    C.C_ZN15QCalendarWidget12showNextYearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showNextYear", args)
  }

}

// setMinimumDate(const class QDate &)
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
    // invoke: void setMinimumDate(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QCalendarWidget14setMinimumDateERK5QDate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setMinimumDate", args)
  }

}

// yearShown()
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
    // invoke: int yearShown()
    var ret = C.C_ZNK15QCalendarWidget9yearShownEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "yearShown", args)
  }

}

// setHeaderTextFormat(const class QTextCharFormat &)
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
    // invoke: void setHeaderTextFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QCalendarWidget19setHeaderTextFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setHeaderTextFormat", args)
  }

}

// verticalHeaderFormat()
func (this *QCalendarWidget) verticalHeaderFormat(args ...interface{}) () {
  // verticalHeaderFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget20verticalHeaderFormatEv
    // invoke: QCalendarWidget::VerticalHeaderFormat verticalHeaderFormat()
    C.C_ZNK15QCalendarWidget20verticalHeaderFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "verticalHeaderFormat", args)
  }

}

// minimumDate()
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
    // invoke: QDate minimumDate()
    var ret = C.C_ZNK15QCalendarWidget11minimumDateEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "minimumDate", args)
  }

}

// sizeHint()
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
    // invoke: QSize sizeHint()
    var ret = C.C_ZNK15QCalendarWidget8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "sizeHint", args)
  }

}

// setSelectedDate(const class QDate &)
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
    // invoke: void setSelectedDate(const class QDate &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QCalendarWidget15setSelectedDateERK5QDate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setSelectedDate", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QCalendarWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "metaObject", args)
  }

}

// selectionMode()
func (this *QCalendarWidget) selectionMode(args ...interface{}) () {
  // selectionMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QCalendarWidget13selectionModeEv
    // invoke: QCalendarWidget::SelectionMode selectionMode()
    C.C_ZNK15QCalendarWidget13selectionModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "selectionMode", args)
  }

}

// setCurrentPage(int, int)
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
    // invoke: void setCurrentPage(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN15QCalendarWidget14setCurrentPageEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setCurrentPage", args)
  }

}

// setDateTextFormat(const class QDate &, const class QTextCharFormat &)
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
    // invoke: void setDateTextFormat(const class QDate &, const class QTextCharFormat &)
    var arg0 = args[0].(QDate).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QCalendarWidget17setDateTextFormatERK5QDateRK15QTextCharFormat(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateTextFormat", args)
  }

}

// selectedDate()
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
    // invoke: QDate selectedDate()
    var ret = C.C_ZNK15QCalendarWidget12selectedDateEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "selectedDate", args)
  }

}

// maximumDate()
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
    // invoke: QDate maximumDate()
    var ret = C.C_ZNK15QCalendarWidget11maximumDateEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "maximumDate", args)
  }

}

// minimumSizeHint()
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
    // invoke: QSize minimumSizeHint()
    var ret = C.C_ZNK15QCalendarWidget15minimumSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "minimumSizeHint", args)
  }

}

// showPreviousMonth()
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
    // invoke: void showPreviousMonth()
    C.C_ZN15QCalendarWidget17showPreviousMonthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showPreviousMonth", args)
  }

}

// dateEditAcceptDelay()
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
    // invoke: int dateEditAcceptDelay()
    var ret = C.C_ZNK15QCalendarWidget19dateEditAcceptDelayEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "dateEditAcceptDelay", args)
  }

}

// showPreviousYear()
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
    // invoke: void showPreviousYear()
    C.C_ZN15QCalendarWidget16showPreviousYearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showPreviousYear", args)
  }

}

// isGridVisible()
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
    // invoke: bool isGridVisible()
    var ret = C.C_ZNK15QCalendarWidget13isGridVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "isGridVisible", args)
  }

}

// setDateEditEnabled(_Bool)
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
    // invoke: void setDateEditEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QCalendarWidget18setDateEditEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateEditEnabled", args)
  }

}

// QCalendarWidget(class QWidget *)
func NewQCalendarWidget(args ...interface{}) *QCalendarWidget {
  // QCalendarWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QCalendarWidgetC1EP7QWidget
    // invoke: void QCalendarWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QCalendarWidgetC2EP7QWidget(arg0)
    return &QCalendarWidget{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QCalendarWidget", "QCalendarWidget", args)
  }

  return nil // QCalendarWidget{qclsinst:qthis}
}

// showNextMonth()
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
    // invoke: void showNextMonth()
    C.C_ZN15QCalendarWidget13showNextMonthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showNextMonth", args)
  }

}

// showToday()
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
    // invoke: void showToday()
    C.C_ZN15QCalendarWidget9showTodayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "showToday", args)
  }

}

// <= body block end

