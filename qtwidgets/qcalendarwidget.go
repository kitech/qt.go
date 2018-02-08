package qtwidgets

// /usr/include/qt/QtWidgets/qcalendarwidget.h
// #include <qcalendarwidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// bool event(class QEvent *)
func (this *QCalendarWidget) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QCalendarWidget) InheritEventFilter(f func(watched *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QCalendarWidget) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QCalendarWidget) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QCalendarWidget) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void paintCell(class QPainter *, const class QRect &, const class QDate &)
func (this *QCalendarWidget) InheritPaintCell(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, date *qtcore.QDate) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintCell", f)
}

// void updateCell(const class QDate &)
func (this *QCalendarWidget) InheritUpdateCell(f func(date *qtcore.QDate) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCell", f)
}

// void updateCells()
func (this *QCalendarWidget) InheritUpdateCells(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCells", f)
}

type QCalendarWidget struct {
	*QWidget
}

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
// [8] const QMetaObject * metaObject()
func (this *QCalendarWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCalendarWidget(QWidget *)
func NewQCalendarWidget(parent *QWidget /*777 QWidget **/) *QCalendarWidget {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCalendarWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCalendarWidget()
func DeleteQCalendarWidget(this *QCalendarWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QCalendarWidget) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QCalendarWidget) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate selectedDate()
func (this *QCalendarWidget) SelectedDate() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget12selectedDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int yearShown()
func (this *QCalendarWidget) YearShown() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget9yearShownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int monthShown()
func (this *QCalendarWidget) MonthShown() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget10monthShownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate minimumDate()
func (this *QCalendarWidget) MinimumDate() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget11minimumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumDate(const QDate &)
func (this *QCalendarWidget) SetMinimumDate(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget14setMinimumDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate maximumDate()
func (this *QCalendarWidget) MaximumDate() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget11maximumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumDate(const QDate &)
func (this *QCalendarWidget) SetMaximumDate(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget14setMaximumDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::DayOfWeek firstDayOfWeek()
func (this *QCalendarWidget) FirstDayOfWeek() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget14firstDayOfWeekEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFirstDayOfWeek(Qt::DayOfWeek)
func (this *QCalendarWidget) SetFirstDayOfWeek(dayOfWeek int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget17setFirstDayOfWeekEN2Qt9DayOfWeekE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dayOfWeek)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNavigationBarVisible()
func (this *QCalendarWidget) IsNavigationBarVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget22isNavigationBarVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isGridVisible()
func (this *QCalendarWidget) IsGridVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget13isGridVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QCalendarWidget::SelectionMode selectionMode()
func (this *QCalendarWidget) SelectionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget13selectionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionMode(enum QCalendarWidget::SelectionMode)
func (this *QCalendarWidget) SetSelectionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget16setSelectionModeENS_13SelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:118
// index:0
// Public Visibility=Default Availability=Available
// [4] QCalendarWidget::HorizontalHeaderFormat horizontalHeaderFormat()
func (this *QCalendarWidget) HorizontalHeaderFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget22horizontalHeaderFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeaderFormat(enum QCalendarWidget::HorizontalHeaderFormat)
func (this *QCalendarWidget) SetHorizontalHeaderFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget25setHorizontalHeaderFormatENS_22HorizontalHeaderFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] QCalendarWidget::VerticalHeaderFormat verticalHeaderFormat()
func (this *QCalendarWidget) VerticalHeaderFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget20verticalHeaderFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeaderFormat(enum QCalendarWidget::VerticalHeaderFormat)
func (this *QCalendarWidget) SetVerticalHeaderFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget23setVerticalHeaderFormatENS_20VerticalHeaderFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:124
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat headerTextFormat()
func (this *QCalendarWidget) HeaderTextFormat() *qtgui.QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget16headerTextFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeaderTextFormat(const QTextCharFormat &)
func (this *QCalendarWidget) SetHeaderTextFormat(format *qtgui.QTextCharFormat) {
	var convArg0 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget19setHeaderTextFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:127
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat weekdayTextFormat(Qt::DayOfWeek)
func (this *QCalendarWidget) WeekdayTextFormat(dayOfWeek int) *qtgui.QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget17weekdayTextFormatEN2Qt9DayOfWeekE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dayOfWeek)
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWeekdayTextFormat(Qt::DayOfWeek, const QTextCharFormat &)
func (this *QCalendarWidget) SetWeekdayTextFormat(dayOfWeek int, format *qtgui.QTextCharFormat) {
	var convArg1 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget20setWeekdayTextFormatEN2Qt9DayOfWeekERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dayOfWeek, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:131
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat dateTextFormat(const QDate &)
func (this *QCalendarWidget) DateTextFormat(date *qtcore.QDate) *qtgui.QTextCharFormat /*123*/ {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget14dateTextFormatERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateTextFormat(const QDate &, const QTextCharFormat &)
func (this *QCalendarWidget) SetDateTextFormat(date *qtcore.QDate, format *qtgui.QTextCharFormat) {
	var convArg0 = date.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget17setDateTextFormatERK5QDateRK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:134
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDateEditEnabled()
func (this *QCalendarWidget) IsDateEditEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget17isDateEditEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateEditEnabled(_Bool)
func (this *QCalendarWidget) SetDateEditEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget18setDateEditEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int dateEditAcceptDelay()
func (this *QCalendarWidget) DateEditAcceptDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget19dateEditAcceptDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateEditAcceptDelay(int)
func (this *QCalendarWidget) SetDateEditAcceptDelay(delay int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget22setDateEditAcceptDelayEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), delay)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:141
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QCalendarWidget) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QCalendarWidget) EventFilter(watched *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = watched.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:143
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QCalendarWidget) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:144
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QCalendarWidget) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QCalendarWidget) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:147
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintCell(QPainter *, const QRect &, const QDate &)
func (this *QCalendarWidget) PaintCell(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, date *qtcore.QDate) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg2 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCalendarWidget9paintCellEP8QPainterRK5QRectRK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:148
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateCell(const QDate &)
func (this *QCalendarWidget) UpdateCell(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget10updateCellERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:149
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateCells()
func (this *QCalendarWidget) UpdateCells() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget11updateCellsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectedDate(const QDate &)
func (this *QCalendarWidget) SetSelectedDate(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget15setSelectedDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateRange(const QDate &, const QDate &)
func (this *QCalendarWidget) SetDateRange(min *qtcore.QDate, max *qtcore.QDate) {
	var convArg0 = min.GetCthis()
	var convArg1 = max.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget12setDateRangeERK5QDateS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentPage(int, int)
func (this *QCalendarWidget) SetCurrentPage(year int, month int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget14setCurrentPageEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), year, month)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGridVisible(_Bool)
func (this *QCalendarWidget) SetGridVisible(show bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget14setGridVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNavigationBarVisible(_Bool)
func (this *QCalendarWidget) SetNavigationBarVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget23setNavigationBarVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showNextMonth()
func (this *QCalendarWidget) ShowNextMonth() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget13showNextMonthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showPreviousMonth()
func (this *QCalendarWidget) ShowPreviousMonth() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget17showPreviousMonthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showNextYear()
func (this *QCalendarWidget) ShowNextYear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget12showNextYearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showPreviousYear()
func (this *QCalendarWidget) ShowPreviousYear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget16showPreviousYearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showSelectedDate()
func (this *QCalendarWidget) ShowSelectedDate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget16showSelectedDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showToday()
func (this *QCalendarWidget) ShowToday() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget9showTodayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()
func (this *QCalendarWidget) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(const QDate &)
func (this *QCalendarWidget) Clicked(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget7clickedERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated(const QDate &)
func (this *QCalendarWidget) Activated(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget9activatedERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcalendarwidget.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentPageChanged(int, int)
func (this *QCalendarWidget) CurrentPageChanged(year int, month int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCalendarWidget18currentPageChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), year, month)
	gopp.ErrPrint(err, rv)
}

type QCalendarWidget__HorizontalHeaderFormat = int

const QCalendarWidget__NoHorizontalHeader QCalendarWidget__HorizontalHeaderFormat = 0
const QCalendarWidget__SingleLetterDayNames QCalendarWidget__HorizontalHeaderFormat = 1
const QCalendarWidget__ShortDayNames QCalendarWidget__HorizontalHeaderFormat = 2
const QCalendarWidget__LongDayNames QCalendarWidget__HorizontalHeaderFormat = 3

type QCalendarWidget__VerticalHeaderFormat = int

const QCalendarWidget__NoVerticalHeader QCalendarWidget__VerticalHeaderFormat = 0
const QCalendarWidget__ISOWeekNumbers QCalendarWidget__VerticalHeaderFormat = 1

type QCalendarWidget__SelectionMode = int

const QCalendarWidget__NoSelection QCalendarWidget__SelectionMode = 0
const QCalendarWidget__SingleSelection QCalendarWidget__SelectionMode = 1

//  body block end
