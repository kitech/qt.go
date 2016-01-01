package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qcalendarwidget.h
// dst-file: /src/widgets/qcalendarwidget.go
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
// class sizeof(QCalendarWidget)=1
type QCalendarWidget struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _activated QCalendarWidget_activated_signal;
//  _clicked QCalendarWidget_clicked_signal;
//  _currentPageChanged QCalendarWidget_currentPageChanged_signal;
//  _selectionChanged QCalendarWidget_selectionChanged_signal;
}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateEditAcceptDelay", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setGridVisible", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setSelectedDate", args)
 }

}


func NewQCalendarWidget(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setNavigationBarVisible", args)
 }

}


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
  case 1:
    // invoke: _ZNK15QCalendarWidget14dateTextFormatEv
  default:
    qtrt.ErrorResolve("QCalendarWidget", "dateTextFormat", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setMinimumDate", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateEditEnabled", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateTextFormat", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setDateRange", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setHeaderTextFormat", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setMaximumDate", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QCalendarWidget", "setCurrentPage", args)
 }

}


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

