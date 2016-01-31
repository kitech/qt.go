package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern void* C_ZNK15QCalendarWidget16headerTextFormatEv(void* qthis); // 4
  // proto:  void QCalendarWidget::showSelectedDate();
extern void C_ZN15QCalendarWidget16showSelectedDateEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setGridVisible(bool show);
extern void C_ZN15QCalendarWidget14setGridVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QCalendarWidget::setDateEditAcceptDelay(int delay);
extern void C_ZN15QCalendarWidget22setDateEditAcceptDelayEi(void* qthis, int32_t arg0); // 4
  // proto:  void QCalendarWidget::setDateRange(const QDate & min, const QDate & max);
extern void C_ZN15QCalendarWidget12setDateRangeERK5QDateS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QCalendarWidget::isDateEditEnabled();
extern bool C_ZNK15QCalendarWidget17isDateEditEnabledEv(void* qthis); // 4
  // proto:  void QCalendarWidget::~QCalendarWidget();
extern void C_ZN15QCalendarWidgetD2Ev(void* qthis); // 4
  // proto:  QTextCharFormat QCalendarWidget::dateTextFormat(const QDate & date);
extern void* C_ZNK15QCalendarWidget14dateTextFormatERK5QDate(void* qthis, void* arg0); // 4
  // proto:  QMap<QDate, QTextCharFormat> QCalendarWidget::dateTextFormat();
extern void C_ZNK15QCalendarWidget14dateTextFormatEv(void* qthis); // 4
  // proto:  bool QCalendarWidget::isNavigationBarVisible();
extern bool C_ZNK15QCalendarWidget22isNavigationBarVisibleEv(void* qthis); // 4
  // proto:  Qt::DayOfWeek QCalendarWidget::firstDayOfWeek();
extern void C_ZNK15QCalendarWidget14firstDayOfWeekEv(void* qthis); // 4
  // proto:  int QCalendarWidget::monthShown();
extern int32_t C_ZNK15QCalendarWidget10monthShownEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setNavigationBarVisible(bool visible);
extern void C_ZN15QCalendarWidget23setNavigationBarVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QCalendarWidget::showNextYear();
extern void C_ZN15QCalendarWidget12showNextYearEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setMinimumDate(const QDate & date);
extern void C_ZN15QCalendarWidget14setMinimumDateERK5QDate(void* qthis, void* arg0); // 4
  // proto:  int QCalendarWidget::yearShown();
extern int32_t C_ZNK15QCalendarWidget9yearShownEv(void* qthis); // 4
  // proto:  void QCalendarWidget::setHeaderTextFormat(const QTextCharFormat & format);
extern void C_ZN15QCalendarWidget19setHeaderTextFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  QCalendarWidget::VerticalHeaderFormat QCalendarWidget::verticalHeaderFormat();
extern void C_ZNK15QCalendarWidget20verticalHeaderFormatEv(void* qthis); // 4
  // proto:  QDate QCalendarWidget::minimumDate();
extern void* C_ZNK15QCalendarWidget11minimumDateEv(void* qthis); // 4
  // proto:  QSize QCalendarWidget::sizeHint();
extern void* C_ZNK15QCalendarWidget8sizeHintEv(void* qthis); // 4
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
extern void* C_ZNK15QCalendarWidget12selectedDateEv(void* qthis); // 4
  // proto:  QDate QCalendarWidget::maximumDate();
extern void* C_ZNK15QCalendarWidget11maximumDateEv(void* qthis); // 4
  // proto:  QSize QCalendarWidget::minimumSizeHint();
extern void* C_ZNK15QCalendarWidget15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QCalendarWidget::showPreviousMonth();
extern void C_ZN15QCalendarWidget17showPreviousMonthEv(void* qthis); // 4
  // proto:  int QCalendarWidget::dateEditAcceptDelay();
extern int32_t C_ZNK15QCalendarWidget19dateEditAcceptDelayEv(void* qthis); // 4
  // proto:  void QCalendarWidget::showPreviousYear();
extern void C_ZN15QCalendarWidget16showPreviousYearEv(void* qthis); // 4
  // proto:  bool QCalendarWidget::isGridVisible();
extern bool C_ZNK15QCalendarWidget13isGridVisibleEv(void* qthis); // 4
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
func (this *QCalendarWidget) Horizontalheaderformat(args ...interface{}) () {
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

  return
}

// setMaximumDate(const class QDate &)
func (this *QCalendarWidget) Setmaximumdate(args ...interface{}) () {
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

  return
}

// headerTextFormat()
func (this *QCalendarWidget) Headertextformat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget16headerTextFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCharFormat{}) // "QTextCharFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "headerTextFormat", args)
  }

  return
}

// showSelectedDate()
func (this *QCalendarWidget) Showselecteddate(args ...interface{}) () {
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

  return
}

// setGridVisible(_Bool)
func (this *QCalendarWidget) Setgridvisible(args ...interface{}) () {
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

  return
}

// setDateEditAcceptDelay(int)
func (this *QCalendarWidget) Setdateeditacceptdelay(args ...interface{}) () {
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

  return
}

// setDateRange(const class QDate &, const class QDate &)
func (this *QCalendarWidget) Setdaterange(args ...interface{}) () {
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

  return
}

// isDateEditEnabled()
func (this *QCalendarWidget) Isdateeditenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget17isDateEditEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "isDateEditEnabled", args)
  }

  return
}

// ~QCalendarWidget()
func (this *QCalendarWidget) Freeqcalendarwidget(args ...interface{}) () {
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

  return
}

// dateTextFormat(const class QDate &)
func (this *QCalendarWidget) Datetextformat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget14dateTextFormatERK5QDate(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCharFormat{}) // "QTextCharFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK15QCalendarWidget14dateTextFormatEv
    // invoke: QMap<QDate, QTextCharFormat> dateTextFormat()
    C.C_ZNK15QCalendarWidget14dateTextFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCalendarWidget", "dateTextFormat", args)
  }

  return
}

// isNavigationBarVisible()
func (this *QCalendarWidget) Isnavigationbarvisible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget22isNavigationBarVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "isNavigationBarVisible", args)
  }

  return
}

// firstDayOfWeek()
func (this *QCalendarWidget) Firstdayofweek(args ...interface{}) () {
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

  return
}

// monthShown()
func (this *QCalendarWidget) Monthshown(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget10monthShownEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "monthShown", args)
  }

  return
}

// setNavigationBarVisible(_Bool)
func (this *QCalendarWidget) Setnavigationbarvisible(args ...interface{}) () {
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

  return
}

// showNextYear()
func (this *QCalendarWidget) Shownextyear(args ...interface{}) () {
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

  return
}

// setMinimumDate(const class QDate &)
func (this *QCalendarWidget) Setminimumdate(args ...interface{}) () {
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

  return
}

// yearShown()
func (this *QCalendarWidget) Yearshown(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget9yearShownEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "yearShown", args)
  }

  return
}

// setHeaderTextFormat(const class QTextCharFormat &)
func (this *QCalendarWidget) Setheadertextformat(args ...interface{}) () {
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

  return
}

// verticalHeaderFormat()
func (this *QCalendarWidget) Verticalheaderformat(args ...interface{}) () {
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

  return
}

// minimumDate()
func (this *QCalendarWidget) Minimumdate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget11minimumDateEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "minimumDate", args)
  }

  return
}

// sizeHint()
func (this *QCalendarWidget) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "sizeHint", args)
  }

  return
}

// setSelectedDate(const class QDate &)
func (this *QCalendarWidget) Setselecteddate(args ...interface{}) () {
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

  return
}

// metaObject()
func (this *QCalendarWidget) Metaobject(args ...interface{}) () {
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

  return
}

// selectionMode()
func (this *QCalendarWidget) Selectionmode(args ...interface{}) () {
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

  return
}

// setCurrentPage(int, int)
func (this *QCalendarWidget) Setcurrentpage(args ...interface{}) () {
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

  return
}

// setDateTextFormat(const class QDate &, const class QTextCharFormat &)
func (this *QCalendarWidget) Setdatetextformat(args ...interface{}) () {
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

  return
}

// selectedDate()
func (this *QCalendarWidget) Selecteddate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget12selectedDateEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "selectedDate", args)
  }

  return
}

// maximumDate()
func (this *QCalendarWidget) Maximumdate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget11maximumDateEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "maximumDate", args)
  }

  return
}

// minimumSizeHint()
func (this *QCalendarWidget) Minimumsizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget15minimumSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "minimumSizeHint", args)
  }

  return
}

// showPreviousMonth()
func (this *QCalendarWidget) Showpreviousmonth(args ...interface{}) () {
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

  return
}

// dateEditAcceptDelay()
func (this *QCalendarWidget) Dateeditacceptdelay(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget19dateEditAcceptDelayEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "dateEditAcceptDelay", args)
  }

  return
}

// showPreviousYear()
func (this *QCalendarWidget) Showpreviousyear(args ...interface{}) () {
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

  return
}

// isGridVisible()
func (this *QCalendarWidget) Isgridvisible(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK15QCalendarWidget13isGridVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QCalendarWidget", "isGridVisible", args)
  }

  return
}

// setDateEditEnabled(_Bool)
func (this *QCalendarWidget) Setdateeditenabled(args ...interface{}) () {
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

  return
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
func (this *QCalendarWidget) Shownextmonth(args ...interface{}) () {
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

  return
}

// showToday()
func (this *QCalendarWidget) Showtoday(args ...interface{}) () {
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

  return
}

// <= body block end

