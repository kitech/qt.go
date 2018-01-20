//  header block begin
// /usr/include/qt/QtWidgets/qcalendarwidget.h
// #include <qcalendarwidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QCalendarWidget struct {
	*QWidget
}

func (this *QCalendarWidget) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QCalendarWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:92
// index:0
// void QCalendarWidget(class QWidget *)
func NewQCalendarWidget(parent unsafe.Pointer) *QCalendarWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQCalendarWidgetFromPointer(cthis)
	return gothis
}
func NewQCalendarWidgetFromPointer(cthis unsafe.Pointer) *QCalendarWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QCalendarWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:93
// index:0
// virtual
// void ~QCalendarWidget()
func DeleteQCalendarWidget(*QCalendarWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:95
// index:0
// virtual
// QSize sizeHint()
func (this *QCalendarWidget) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:96
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QCalendarWidget) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:98
// index:0
// QDate selectedDate()
func (this *QCalendarWidget) SelectedDate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget12selectedDateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:100
// index:0
// int yearShown()
func (this *QCalendarWidget) YearShown() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget9yearShownEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:101
// index:0
// int monthShown()
func (this *QCalendarWidget) MonthShown() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget10monthShownEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:103
// index:0
// QDate minimumDate()
func (this *QCalendarWidget) MinimumDate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget11minimumDateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:104
// index:0
// void setMinimumDate(const class QDate &)
func (this *QCalendarWidget) SetMinimumDate(date unsafe.Pointer) {
	// 0: (, date const QDate &), (date)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget14setMinimumDateERK5QDate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:106
// index:0
// QDate maximumDate()
func (this *QCalendarWidget) MaximumDate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget11maximumDateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:107
// index:0
// void setMaximumDate(const class QDate &)
func (this *QCalendarWidget) SetMaximumDate(date unsafe.Pointer) {
	// 0: (, date const QDate &), (date)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget14setMaximumDateERK5QDate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:109
// index:0
// Qt::DayOfWeek firstDayOfWeek()
func (this *QCalendarWidget) FirstDayOfWeek() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget14firstDayOfWeekEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:110
// index:0
// void setFirstDayOfWeek(Qt::DayOfWeek)
func (this *QCalendarWidget) SetFirstDayOfWeek(dayOfWeek int) {
	// 0: (, dayOfWeek Qt::DayOfWeek), (&dayOfWeek)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget17setFirstDayOfWeekEN2Qt9DayOfWeekE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dayOfWeek)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:112
// index:0
// bool isNavigationBarVisible()
func (this *QCalendarWidget) IsNavigationBarVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget22isNavigationBarVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:113
// index:0
// bool isGridVisible()
func (this *QCalendarWidget) IsGridVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget13isGridVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:115
// index:0
// QCalendarWidget::SelectionMode selectionMode()
func (this *QCalendarWidget) SelectionMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget13selectionModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:116
// index:0
// void setSelectionMode(enum QCalendarWidget::SelectionMode)
func (this *QCalendarWidget) SetSelectionMode(mode int) {
	// 0: (, mode QCalendarWidget::SelectionMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget16setSelectionModeENS_13SelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:118
// index:0
// QCalendarWidget::HorizontalHeaderFormat horizontalHeaderFormat()
func (this *QCalendarWidget) HorizontalHeaderFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget22horizontalHeaderFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:119
// index:0
// void setHorizontalHeaderFormat(enum QCalendarWidget::HorizontalHeaderFormat)
func (this *QCalendarWidget) SetHorizontalHeaderFormat(format int) {
	// 0: (, format QCalendarWidget::HorizontalHeaderFormat), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget25setHorizontalHeaderFormatENS_22HorizontalHeaderFormatE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:121
// index:0
// QCalendarWidget::VerticalHeaderFormat verticalHeaderFormat()
func (this *QCalendarWidget) VerticalHeaderFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget20verticalHeaderFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:122
// index:0
// void setVerticalHeaderFormat(enum QCalendarWidget::VerticalHeaderFormat)
func (this *QCalendarWidget) SetVerticalHeaderFormat(format int) {
	// 0: (, format QCalendarWidget::VerticalHeaderFormat), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget23setVerticalHeaderFormatENS_20VerticalHeaderFormatE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:124
// index:0
// QTextCharFormat headerTextFormat()
func (this *QCalendarWidget) HeaderTextFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget16headerTextFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:125
// index:0
// void setHeaderTextFormat(const class QTextCharFormat &)
func (this *QCalendarWidget) SetHeaderTextFormat(format unsafe.Pointer) {
	// 0: (, format const QTextCharFormat &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget19setHeaderTextFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:127
// index:0
// QTextCharFormat weekdayTextFormat(Qt::DayOfWeek)
func (this *QCalendarWidget) WeekdayTextFormat(dayOfWeek int) {
	// 0: (, dayOfWeek Qt::DayOfWeek), (&dayOfWeek)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget17weekdayTextFormatEN2Qt9DayOfWeekE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dayOfWeek)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:128
// index:0
// void setWeekdayTextFormat(Qt::DayOfWeek, const class QTextCharFormat &)
func (this *QCalendarWidget) SetWeekdayTextFormat(dayOfWeek int, format unsafe.Pointer) {
	// 0: (, dayOfWeek Qt::DayOfWeek, format const QTextCharFormat &), (&dayOfWeek, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget20setWeekdayTextFormatEN2Qt9DayOfWeekERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dayOfWeek, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:130
// index:0
// QMap<QDate, QTextCharFormat> dateTextFormat()
func (this *QCalendarWidget) DateTextFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget14dateTextFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:131
// index:1
// QTextCharFormat dateTextFormat(const class QDate &)
func (this *QCalendarWidget) DateTextFormat_1(date unsafe.Pointer) {
	// 1: (, date const QDate &), (date)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget14dateTextFormatERK5QDate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:132
// index:0
// void setDateTextFormat(const class QDate &, const class QTextCharFormat &)
func (this *QCalendarWidget) SetDateTextFormat(date unsafe.Pointer, format unsafe.Pointer) {
	// 0: (, date const QDate &, format const QTextCharFormat &), (date, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget17setDateTextFormatERK5QDateRK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:134
// index:0
// bool isDateEditEnabled()
func (this *QCalendarWidget) IsDateEditEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget17isDateEditEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:135
// index:0
// void setDateEditEnabled(_Bool)
func (this *QCalendarWidget) SetDateEditEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget18setDateEditEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:137
// index:0
// int dateEditAcceptDelay()
func (this *QCalendarWidget) DateEditAcceptDelay() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget19dateEditAcceptDelayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:138
// index:0
// void setDateEditAcceptDelay(int)
func (this *QCalendarWidget) SetDateEditAcceptDelay(delay int) {
	// 0: (, delay int), (&delay)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget22setDateEditAcceptDelayEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &delay)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:141
// index:0
// virtual
// bool event(class QEvent *)
func (this *QCalendarWidget) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:142
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QCalendarWidget) EventFilter(watched unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, watched QObject *, event QEvent *), (watched, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), watched, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:143
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QCalendarWidget) MousePressEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:144
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QCalendarWidget) ResizeEvent(event unsafe.Pointer) {
	// 0: (, event QResizeEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:145
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QCalendarWidget) KeyPressEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:147
// index:0
// virtual
// void paintCell(class QPainter *, const class QRect &, const class QDate &)
func (this *QCalendarWidget) PaintCell(painter unsafe.Pointer, rect unsafe.Pointer, date unsafe.Pointer) {
	// 0: (, painter QPainter *, rect const QRect &, date const QDate &), (painter, rect, date)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QCalendarWidget9paintCellEP8QPainterRK5QRectRK5QDate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect, date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:148
// index:0
// void updateCell(const class QDate &)
func (this *QCalendarWidget) UpdateCell(date unsafe.Pointer) {
	// 0: (, date const QDate &), (date)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget10updateCellERK5QDate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:149
// index:0
// void updateCells()
func (this *QCalendarWidget) UpdateCells() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget11updateCellsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:152
// index:0
// void setSelectedDate(const class QDate &)
func (this *QCalendarWidget) SetSelectedDate(date unsafe.Pointer) {
	// 0: (, date const QDate &), (date)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget15setSelectedDateERK5QDate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:153
// index:0
// void setDateRange(const class QDate &, const class QDate &)
func (this *QCalendarWidget) SetDateRange(min unsafe.Pointer, max unsafe.Pointer) {
	// 0: (, min const QDate &, max const QDate &), (min, max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget12setDateRangeERK5QDateS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), min, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:154
// index:0
// void setCurrentPage(int, int)
func (this *QCalendarWidget) SetCurrentPage(year int, month int) {
	// 0: (, year int, month int), (&year, &month)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget14setCurrentPageEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &year, &month)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:155
// index:0
// void setGridVisible(_Bool)
func (this *QCalendarWidget) SetGridVisible(show bool) {
	// 0: (, show bool), (&show)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget14setGridVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:156
// index:0
// void setNavigationBarVisible(_Bool)
func (this *QCalendarWidget) SetNavigationBarVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget23setNavigationBarVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:157
// index:0
// void showNextMonth()
func (this *QCalendarWidget) ShowNextMonth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget13showNextMonthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:158
// index:0
// void showPreviousMonth()
func (this *QCalendarWidget) ShowPreviousMonth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget17showPreviousMonthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:159
// index:0
// void showNextYear()
func (this *QCalendarWidget) ShowNextYear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget12showNextYearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:160
// index:0
// void showPreviousYear()
func (this *QCalendarWidget) ShowPreviousYear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget16showPreviousYearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:161
// index:0
// void showSelectedDate()
func (this *QCalendarWidget) ShowSelectedDate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget16showSelectedDateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:162
// index:0
// void showToday()
func (this *QCalendarWidget) ShowToday() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget9showTodayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:165
// index:0
// void selectionChanged()
func (this *QCalendarWidget) SelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget16selectionChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:166
// index:0
// void clicked(const class QDate &)
func (this *QCalendarWidget) Clicked(date unsafe.Pointer) {
	// 0: (, date const QDate &), (date)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget7clickedERK5QDate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:167
// index:0
// void activated(const class QDate &)
func (this *QCalendarWidget) Activated(date unsafe.Pointer) {
	// 0: (, date const QDate &), (date)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget9activatedERK5QDate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:168
// index:0
// void currentPageChanged(int, int)
func (this *QCalendarWidget) CurrentPageChanged(year int, month int) {
	// 0: (, year int, month int), (&year, &month)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QCalendarWidget18currentPageChangedEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &year, &month)
	gopp.ErrPrint(err, rv)
}

//  body block end
