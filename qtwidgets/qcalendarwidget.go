package qtwidgets

// /usr/include/qt/QtWidgets/qcalendarwidget.h
// #include <qcalendarwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QCalendarWidget) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QCalendarWidget) InheritEventFilter(f func(watched *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QCalendarWidget) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QCalendarWidget) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QCalendarWidget) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void paintCell(QPainter *, const QRect &, const QDate &)
func (this *QCalendarWidget) InheritPaintCell(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, date *qtcore.QDate) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintCell", f)
}

// void updateCell(const QDate &)
func (this *QCalendarWidget) InheritUpdateCell(f func(date *qtcore.QDate) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCell", f)
}

// void updateCells()
func (this *QCalendarWidget) InheritUpdateCells(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCells", f)
}

/*

 */
type QCalendarWidget struct {
	*QWidget
}
type QCalendarWidget_ITF interface {
	QWidget_ITF
	QCalendarWidget_PTR() *QCalendarWidget
}

func (ptr *QCalendarWidget) QCalendarWidget_PTR() *QCalendarWidget { return ptr }

func (this *QCalendarWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QCalendarWidget) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQCalendarWidgetFromPointer(cthis unsafe.Pointer) *QCalendarWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QCalendarWidget{bcthis0}
}
func (*QCalendarWidget) NewFromPointer(cthis unsafe.Pointer) *QCalendarWidget {
	return NewQCalendarWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCalendarWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCalendarWidget(QWidget *)

/*
Constructs a calendar widget with the given parent.

The widget is initialized with the current month and year, and the currently selected date is today.

See also setCurrentPage().
*/
func NewQCalendarWidget(parent QWidget_ITF /*777 QWidget **/) *QCalendarWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCalendarWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCalendarWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCalendarWidget(QWidget *)

/*
Constructs a calendar widget with the given parent.

The widget is initialized with the current month and year, and the currently selected date is today.

See also setCurrentPage().
*/
func NewQCalendarWidget__() *QCalendarWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCalendarWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCalendarWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCalendarWidget()

/*

 */
func DeleteQCalendarWidget(this *QCalendarWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QCalendarWidget) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QCalendarWidget) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate selectedDate() const

/*

 */
func (this *QCalendarWidget) SelectedDate() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget12selectedDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int yearShown() const

/*
Returns the year of the currently displayed month. Months are numbered from 1 to 12.

See also monthShown() and setCurrentPage().
*/
func (this *QCalendarWidget) YearShown() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget9yearShownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int monthShown() const

/*
Returns the currently displayed month. Months are numbered from 1 to 12.

See also yearShown() and setCurrentPage().
*/
func (this *QCalendarWidget) MonthShown() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget10monthShownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate minimumDate() const

/*

 */
func (this *QCalendarWidget) MinimumDate() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget11minimumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumDate(const QDate &)

/*

 */
func (this *QCalendarWidget) SetMinimumDate(date qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget14setMinimumDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate maximumDate() const

/*

 */
func (this *QCalendarWidget) MaximumDate() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget11maximumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumDate(const QDate &)

/*

 */
func (this *QCalendarWidget) SetMaximumDate(date qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget14setMaximumDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DayOfWeek firstDayOfWeek() const

/*

 */
func (this *QCalendarWidget) FirstDayOfWeek() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget14firstDayOfWeekEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFirstDayOfWeek(Qt::DayOfWeek)

/*

 */
func (this *QCalendarWidget) SetFirstDayOfWeek(dayOfWeek int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget17setFirstDayOfWeekEN2Qt9DayOfWeekE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dayOfWeek)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNavigationBarVisible() const

/*

 */
func (this *QCalendarWidget) IsNavigationBarVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget22isNavigationBarVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isGridVisible() const

/*

 */
func (this *QCalendarWidget) IsGridVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget13isGridVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QCalendarWidget::SelectionMode selectionMode() const

/*

 */
func (this *QCalendarWidget) SelectionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget13selectionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionMode(QCalendarWidget::SelectionMode)

/*

 */
func (this *QCalendarWidget) SetSelectionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget16setSelectionModeENS_13SelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:118
// index:0
// Public Visibility=Default Availability=Available
// [4] QCalendarWidget::HorizontalHeaderFormat horizontalHeaderFormat() const

/*

 */
func (this *QCalendarWidget) HorizontalHeaderFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget22horizontalHeaderFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeaderFormat(QCalendarWidget::HorizontalHeaderFormat)

/*

 */
func (this *QCalendarWidget) SetHorizontalHeaderFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget25setHorizontalHeaderFormatENS_22HorizontalHeaderFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] QCalendarWidget::VerticalHeaderFormat verticalHeaderFormat() const

/*

 */
func (this *QCalendarWidget) VerticalHeaderFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget20verticalHeaderFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeaderFormat(QCalendarWidget::VerticalHeaderFormat)

/*

 */
func (this *QCalendarWidget) SetVerticalHeaderFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget23setVerticalHeaderFormatENS_20VerticalHeaderFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:124
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat headerTextFormat() const

/*
Returns the text char format for rendering the header.

See also setHeaderTextFormat().
*/
func (this *QCalendarWidget) HeaderTextFormat() *qtgui.QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget16headerTextFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeaderTextFormat(const QTextCharFormat &)

/*
Sets the text char format for rendering the header to format. If you also set a weekday text format, this format's foreground and background color will take precedence over the header's format. The other formatting information will still be decided by the header's format.

See also headerTextFormat().
*/
func (this *QCalendarWidget) SetHeaderTextFormat(format qtgui.QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg0 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget19setHeaderTextFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:127
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat weekdayTextFormat(Qt::DayOfWeek) const

/*
Returns the text char format for rendering of day in the week dayOfWeek.

See also setWeekdayTextFormat() and headerTextFormat().
*/
func (this *QCalendarWidget) WeekdayTextFormat(dayOfWeek int) *qtgui.QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget17weekdayTextFormatEN2Qt9DayOfWeekE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dayOfWeek)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWeekdayTextFormat(Qt::DayOfWeek, const QTextCharFormat &)

/*
Sets the text char format for rendering of day in the week dayOfWeek to format. The format will take precedence over the header format in case of foreground and background color. Other text formatting information is taken from the headers format.

See also weekdayTextFormat() and setHeaderTextFormat().
*/
func (this *QCalendarWidget) SetWeekdayTextFormat(dayOfWeek int, format qtgui.QTextCharFormat_ITF) {
	var convArg1 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg1 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget20setWeekdayTextFormatEN2Qt9DayOfWeekERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dayOfWeek, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:131
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat dateTextFormat(const QDate &) const

/*
Returns a QMap from QDate to QTextCharFormat showing all dates that use a special format that alters their rendering.

See also setDateTextFormat().
*/
func (this *QCalendarWidget) DateTextFormat(date qtcore.QDate_ITF) *qtgui.QTextCharFormat /*123*/ {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget14dateTextFormatERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateTextFormat(const QDate &, const QTextCharFormat &)

/*
Sets the format used to render the given date to that specified by format.

If date is null, all date formats are cleared.

See also dateTextFormat().
*/
func (this *QCalendarWidget) SetDateTextFormat(date qtcore.QDate_ITF, format qtgui.QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg1 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget17setDateTextFormatERK5QDateRK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:134
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDateEditEnabled() const

/*

 */
func (this *QCalendarWidget) IsDateEditEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget17isDateEditEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateEditEnabled(bool)

/*

 */
func (this *QCalendarWidget) SetDateEditEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget18setDateEditEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int dateEditAcceptDelay() const

/*

 */
func (this *QCalendarWidget) DateEditAcceptDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget19dateEditAcceptDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateEditAcceptDelay(int)

/*

 */
func (this *QCalendarWidget) SetDateEditAcceptDelay(delay int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget22setDateEditAcceptDelayEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), delay)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:141
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QCalendarWidget) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().
*/
func (this *QCalendarWidget) EventFilter(watched qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if watched != nil && watched.QObject_PTR() != nil {
		convArg0 = watched.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:143
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QCalendarWidget) MousePressEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:144
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QCalendarWidget) ResizeEvent(event qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QResizeEvent_PTR() != nil {
		convArg0 = event.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QCalendarWidget) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:147
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintCell(QPainter *, const QRect &, const QDate &) const

/*
Paints the cell specified by the given date, using the given painter and rect.
*/
func (this *QCalendarWidget) PaintCell(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, date qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg2 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget9paintCellEP8QPainterRK5QRectRK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:148
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateCell(const QDate &)

/*
Updates the cell specified by the given date unless updates are disabled or the cell is hidden.

This function was introduced in  Qt 4.4.

See also updateCells(), yearShown(), and monthShown().
*/
func (this *QCalendarWidget) UpdateCell(date qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget10updateCellERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:149
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateCells()

/*
Updates all visible cells unless updates are disabled.

This function was introduced in  Qt 4.4.

See also updateCell().
*/
func (this *QCalendarWidget) UpdateCells() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget11updateCellsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectedDate(const QDate &)

/*

 */
func (this *QCalendarWidget) SetSelectedDate(date qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget15setSelectedDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateRange(const QDate &, const QDate &)

/*
Defines a date range by setting the minimumDate and maximumDate properties.

The date range restricts the user selection, i.e. the user can only select dates within the specified date range. Note that


  QCalendarWidget *calendar;

  calendar->setDateRange(min, max);



is analogous to


  QCalendarWidget *calendar;

  calendar->setMinimumDate(min);
  calendar->setMaximumDate(max);



If either the min or max parameters are not valid QDate objects, this function does nothing.

See also setMinimumDate() and setMaximumDate().
*/
func (this *QCalendarWidget) SetDateRange(min qtcore.QDate_ITF, max qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if min != nil && min.QDate_PTR() != nil {
		convArg0 = min.QDate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if max != nil && max.QDate_PTR() != nil {
		convArg1 = max.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget12setDateRangeERK5QDateS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentPage(int, int)

/*
Displays the given month of the given year without changing the selected date. Use the setSelectedDate() function to alter the selected date.

The currently displayed month and year can be retrieved using the monthShown() and yearShown() functions respectively.

See also yearShown(), monthShown(), showPreviousMonth(), showNextMonth(), showPreviousYear(), and showNextYear().
*/
func (this *QCalendarWidget) SetCurrentPage(year int, month int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget14setCurrentPageEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), year, month)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGridVisible(bool)

/*

 */
func (this *QCalendarWidget) SetGridVisible(show bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget14setGridVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), show)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNavigationBarVisible(bool)

/*

 */
func (this *QCalendarWidget) SetNavigationBarVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget23setNavigationBarVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showNextMonth()

/*
Shows the next month relative to the currently displayed month. Note that the selected date is not changed.

See also showPreviousMonth(), setCurrentPage(), and setSelectedDate().
*/
func (this *QCalendarWidget) ShowNextMonth() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget13showNextMonthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showPreviousMonth()

/*
Shows the previous month relative to the currently displayed month. Note that the selected date is not changed.

See also showNextMonth(), setCurrentPage(), and setSelectedDate().
*/
func (this *QCalendarWidget) ShowPreviousMonth() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget17showPreviousMonthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showNextYear()

/*
Shows the currently displayed month in the next year relative to the currently displayed year. Note that the selected date is not changed.

See also showPreviousYear(), setCurrentPage(), and setSelectedDate().
*/
func (this *QCalendarWidget) ShowNextYear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget12showNextYearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showPreviousYear()

/*
Shows the currently displayed month in the previous year relative to the currently displayed year. Note that the selected date is not changed.

See also showNextYear(), setCurrentPage(), and setSelectedDate().
*/
func (this *QCalendarWidget) ShowPreviousYear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget16showPreviousYearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showSelectedDate()

/*
Shows the month of the selected date.

See also selectedDate() and setCurrentPage().
*/
func (this *QCalendarWidget) ShowSelectedDate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget16showSelectedDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showToday()

/*
Shows the month of the today's date.

See also selectedDate() and setCurrentPage().
*/
func (this *QCalendarWidget) ShowToday() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget9showTodayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()

/*
This signal is emitted when the currently selected date is changed.

The currently selected date can be changed by the user using the mouse or keyboard, or by the programmer using setSelectedDate().

See also selectedDate().
*/
func (this *QCalendarWidget) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(const QDate &)

/*
This signal is emitted when a mouse button is clicked. The date the mouse was clicked on is specified by date. The signal is only emitted when clicked on a valid date, e.g., dates are not outside the minimumDate() and maximumDate(). If the selection mode is NoSelection, this signal will not be emitted.
*/
func (this *QCalendarWidget) Clicked(date qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget7clickedERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated(const QDate &)

/*
This signal is emitted whenever the user presses the Return or Enter key or double-clicks a date in the calendar widget.
*/
func (this *QCalendarWidget) Activated(date qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget9activatedERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentPageChanged(int, int)

/*
This signal is emitted when the currently shown month is changed. The new year and month are passed as parameters.

See also setCurrentPage().
*/
func (this *QCalendarWidget) CurrentPageChanged(year int, month int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget18currentPageChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), year, month)
	qtrt.ErrPrint(err, rv)
}

/*
This enum type defines the various formats the horizontal header can display.



See also horizontalHeaderFormat() and VerticalHeaderFormat.

*/
type QCalendarWidget__HorizontalHeaderFormat = int

// The header is hidden.
const QCalendarWidget__NoHorizontalHeader QCalendarWidget__HorizontalHeaderFormat = 0

// The header displays a single letter abbreviation for day names (e.g. M for Monday).
const QCalendarWidget__SingleLetterDayNames QCalendarWidget__HorizontalHeaderFormat = 1

// The header displays a short abbreviation for day names (e.g. Mon for Monday).
const QCalendarWidget__ShortDayNames QCalendarWidget__HorizontalHeaderFormat = 2

// The header displays complete day names (e.g. Monday).
const QCalendarWidget__LongDayNames QCalendarWidget__HorizontalHeaderFormat = 3

/*
This enum type defines the various formats the vertical header can display.



See also verticalHeaderFormat() and HorizontalHeaderFormat.

*/
type QCalendarWidget__VerticalHeaderFormat = int

// The header is hidden.
const QCalendarWidget__NoVerticalHeader QCalendarWidget__VerticalHeaderFormat = 0

// The header displays ISO week numbers as described by QDate::weekNumber().
const QCalendarWidget__ISOWeekNumbers QCalendarWidget__VerticalHeaderFormat = 1

/*
This enum describes the types of selection offered to the user for selecting dates in the calendar.



See also selectionMode.

*/
type QCalendarWidget__SelectionMode = int

// Dates cannot be selected.
const QCalendarWidget__NoSelection QCalendarWidget__SelectionMode = 0

// Single dates can be selected.
const QCalendarWidget__SingleSelection QCalendarWidget__SelectionMode = 1

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
