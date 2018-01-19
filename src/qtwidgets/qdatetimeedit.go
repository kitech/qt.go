//  header block begin
// /usr/include/qt/QtWidgets/qdatetimeedit.h
// #include <qdatetimeedit.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
type QDateTimeEdit struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDateTimeEdit) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:95
// index:0
// void QDateTimeEdit(class QWidget *)
func NewQDateTimeEdit(parent unsafe.Pointer) *QDateTimeEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QDateTimeEdit{cthis}
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:96
// index:1
// void QDateTimeEdit(const class QDateTime &, class QWidget *)
func NewQDateTimeEdit_1(dt unsafe.Pointer, parent unsafe.Pointer) *QDateTimeEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK9QDateTimeP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dt, parent)
	gopp.ErrPrint(err, rv)
	return &QDateTimeEdit{cthis}
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:97
// index:2
// void QDateTimeEdit(const class QDate &, class QWidget *)
func NewQDateTimeEdit_2(d unsafe.Pointer, parent unsafe.Pointer) *QDateTimeEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK5QDateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, d, parent)
	gopp.ErrPrint(err, rv)
	return &QDateTimeEdit{cthis}
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:98
// index:3
// void QDateTimeEdit(const class QTime &, class QWidget *)
func NewQDateTimeEdit_3(t unsafe.Pointer, parent unsafe.Pointer) *QDateTimeEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK5QTimeP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, t, parent)
	gopp.ErrPrint(err, rv)
	return &QDateTimeEdit{cthis}
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:99
// index:0
// virtual
// void ~QDateTimeEdit()
func DeleteQDateTimeEdit(*QDateTimeEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:101
// index:0
// QDateTime dateTime()
func (this *QDateTimeEdit) DateTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit8dateTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:102
// index:0
// QDate date()
func (this *QDateTimeEdit) Date() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit4dateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:103
// index:0
// QTime time()
func (this *QDateTimeEdit) Time() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit4timeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:105
// index:0
// QDateTime minimumDateTime()
func (this *QDateTimeEdit) MinimumDateTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit15minimumDateTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:106
// index:0
// void clearMinimumDateTime()
func (this *QDateTimeEdit) ClearMinimumDateTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit20clearMinimumDateTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:107
// index:0
// void setMinimumDateTime(const class QDateTime &)
func (this *QDateTimeEdit) SetMinimumDateTime(dt unsafe.Pointer) {
	// 0: (, const QDateTime & dt), (dt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit18setMinimumDateTimeERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.cthis, dt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:109
// index:0
// QDateTime maximumDateTime()
func (this *QDateTimeEdit) MaximumDateTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit15maximumDateTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:110
// index:0
// void clearMaximumDateTime()
func (this *QDateTimeEdit) ClearMaximumDateTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit20clearMaximumDateTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:111
// index:0
// void setMaximumDateTime(const class QDateTime &)
func (this *QDateTimeEdit) SetMaximumDateTime(dt unsafe.Pointer) {
	// 0: (, const QDateTime & dt), (dt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit18setMaximumDateTimeERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.cthis, dt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:113
// index:0
// void setDateTimeRange(const class QDateTime &, const class QDateTime &)
func (this *QDateTimeEdit) SetDateTimeRange(min unsafe.Pointer, max unsafe.Pointer) {
	// 0: (, const QDateTime & min, const QDateTime & max), (min, max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16setDateTimeRangeERK9QDateTimeS2_", ffiqt.FFI_TYPE_VOID, this.cthis, min, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:115
// index:0
// QDate minimumDate()
func (this *QDateTimeEdit) MinimumDate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11minimumDateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:116
// index:0
// void setMinimumDate(const class QDate &)
func (this *QDateTimeEdit) SetMinimumDate(min unsafe.Pointer) {
	// 0: (, const QDate & min), (min)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMinimumDateERK5QDate", ffiqt.FFI_TYPE_VOID, this.cthis, min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:117
// index:0
// void clearMinimumDate()
func (this *QDateTimeEdit) ClearMinimumDate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMinimumDateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:119
// index:0
// QDate maximumDate()
func (this *QDateTimeEdit) MaximumDate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11maximumDateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:120
// index:0
// void setMaximumDate(const class QDate &)
func (this *QDateTimeEdit) SetMaximumDate(max unsafe.Pointer) {
	// 0: (, const QDate & max), (max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMaximumDateERK5QDate", ffiqt.FFI_TYPE_VOID, this.cthis, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:121
// index:0
// void clearMaximumDate()
func (this *QDateTimeEdit) ClearMaximumDate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMaximumDateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:123
// index:0
// void setDateRange(const class QDate &, const class QDate &)
func (this *QDateTimeEdit) SetDateRange(min unsafe.Pointer, max unsafe.Pointer) {
	// 0: (, const QDate & min, const QDate & max), (min, max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit12setDateRangeERK5QDateS2_", ffiqt.FFI_TYPE_VOID, this.cthis, min, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:125
// index:0
// QTime minimumTime()
func (this *QDateTimeEdit) MinimumTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11minimumTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:126
// index:0
// void setMinimumTime(const class QTime &)
func (this *QDateTimeEdit) SetMinimumTime(min unsafe.Pointer) {
	// 0: (, const QTime & min), (min)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMinimumTimeERK5QTime", ffiqt.FFI_TYPE_VOID, this.cthis, min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:127
// index:0
// void clearMinimumTime()
func (this *QDateTimeEdit) ClearMinimumTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMinimumTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:129
// index:0
// QTime maximumTime()
func (this *QDateTimeEdit) MaximumTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11maximumTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:130
// index:0
// void setMaximumTime(const class QTime &)
func (this *QDateTimeEdit) SetMaximumTime(max unsafe.Pointer) {
	// 0: (, const QTime & max), (max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMaximumTimeERK5QTime", ffiqt.FFI_TYPE_VOID, this.cthis, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:131
// index:0
// void clearMaximumTime()
func (this *QDateTimeEdit) ClearMaximumTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMaximumTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:133
// index:0
// void setTimeRange(const class QTime &, const class QTime &)
func (this *QDateTimeEdit) SetTimeRange(min unsafe.Pointer, max unsafe.Pointer) {
	// 0: (, const QTime & min, const QTime & max), (min, max)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit12setTimeRangeERK5QTimeS2_", ffiqt.FFI_TYPE_VOID, this.cthis, min, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:135
// index:0
// QDateTimeEdit::Sections displayedSections()
func (this *QDateTimeEdit) DisplayedSections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit17displayedSectionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:136
// index:0
// QDateTimeEdit::Section currentSection()
func (this *QDateTimeEdit) CurrentSection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit14currentSectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:137
// index:0
// QDateTimeEdit::Section sectionAt(int)
func (this *QDateTimeEdit) SectionAt(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit9sectionAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:138
// index:0
// void setCurrentSection(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SetCurrentSection(section int) {
	// 0: (, QDateTimeEdit::Section section), (&section)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit17setCurrentSectionENS_7SectionE", ffiqt.FFI_TYPE_VOID, this.cthis, &section)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:140
// index:0
// int currentSectionIndex()
func (this *QDateTimeEdit) CurrentSectionIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit19currentSectionIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:141
// index:0
// void setCurrentSectionIndex(int)
func (this *QDateTimeEdit) SetCurrentSectionIndex(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit22setCurrentSectionIndexEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:143
// index:0
// QCalendarWidget * calendarWidget()
func (this *QDateTimeEdit) CalendarWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit14calendarWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:144
// index:0
// void setCalendarWidget(class QCalendarWidget *)
func (this *QDateTimeEdit) SetCalendarWidget(calendarWidget unsafe.Pointer) {
	// 0: (, QCalendarWidget * calendarWidget), (calendarWidget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit17setCalendarWidgetEP15QCalendarWidget", ffiqt.FFI_TYPE_VOID, this.cthis, calendarWidget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:146
// index:0
// int sectionCount()
func (this *QDateTimeEdit) SectionCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit12sectionCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:148
// index:0
// void setSelectedSection(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SetSelectedSection(section int) {
	// 0: (, QDateTimeEdit::Section section), (&section)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit18setSelectedSectionENS_7SectionE", ffiqt.FFI_TYPE_VOID, this.cthis, &section)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:150
// index:0
// QString sectionText(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SectionText(section int) {
	// 0: (, QDateTimeEdit::Section section), (&section)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit11sectionTextENS_7SectionE", ffiqt.FFI_TYPE_VOID, this.cthis, &section)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:152
// index:0
// QString displayFormat()
func (this *QDateTimeEdit) DisplayFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit13displayFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:153
// index:0
// void setDisplayFormat(const class QString &)
func (this *QDateTimeEdit) SetDisplayFormat(format unsafe.Pointer) {
	// 0: (, const QString & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16setDisplayFormatERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:155
// index:0
// bool calendarPopup()
func (this *QDateTimeEdit) CalendarPopup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit13calendarPopupEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:156
// index:0
// void setCalendarPopup(_Bool)
func (this *QDateTimeEdit) SetCalendarPopup(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit16setCalendarPopupEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:158
// index:0
// Qt::TimeSpec timeSpec()
func (this *QDateTimeEdit) TimeSpec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit8timeSpecEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:159
// index:0
// void setTimeSpec(Qt::TimeSpec)
func (this *QDateTimeEdit) SetTimeSpec(spec int) {
	// 0: (, Qt::TimeSpec spec), (&spec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit11setTimeSpecEN2Qt8TimeSpecE", ffiqt.FFI_TYPE_VOID, this.cthis, &spec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:161
// index:0
// virtual
// QSize sizeHint()
func (this *QDateTimeEdit) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QDateTimeEdit8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:163
// index:0
// virtual
// void clear()
func (this *QDateTimeEdit) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:164
// index:0
// virtual
// void stepBy(int)
func (this *QDateTimeEdit) StepBy(steps int) {
	// 0: (, int steps), (&steps)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit6stepByEi", ffiqt.FFI_TYPE_VOID, this.cthis, &steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:166
// index:0
// virtual
// bool event(class QEvent *)
func (this *QDateTimeEdit) Event(event unsafe.Pointer) {
	// 0: (, QEvent * event), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.cthis, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:168
// index:0
// void dateTimeChanged(const class QDateTime &)
func (this *QDateTimeEdit) DateTimeChanged(dateTime unsafe.Pointer) {
	// 0: (, const QDateTime & dateTime), (dateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit15dateTimeChangedERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.cthis, dateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:169
// index:0
// void timeChanged(const class QTime &)
func (this *QDateTimeEdit) TimeChanged(time unsafe.Pointer) {
	// 0: (, const QTime & time), (time)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit11timeChangedERK5QTime", ffiqt.FFI_TYPE_VOID, this.cthis, time)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:170
// index:0
// void dateChanged(const class QDate &)
func (this *QDateTimeEdit) DateChanged(date unsafe.Pointer) {
	// 0: (, const QDate & date), (date)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit11dateChangedERK5QDate", ffiqt.FFI_TYPE_VOID, this.cthis, date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:173
// index:0
// void setDateTime(const class QDateTime &)
func (this *QDateTimeEdit) SetDateTime(dateTime unsafe.Pointer) {
	// 0: (, const QDateTime & dateTime), (dateTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit11setDateTimeERK9QDateTime", ffiqt.FFI_TYPE_VOID, this.cthis, dateTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:174
// index:0
// void setDate(const class QDate &)
func (this *QDateTimeEdit) SetDate(date unsafe.Pointer) {
	// 0: (, const QDate & date), (date)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit7setDateERK5QDate", ffiqt.FFI_TYPE_VOID, this.cthis, date)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:175
// index:0
// void setTime(const class QTime &)
func (this *QDateTimeEdit) SetTime(time unsafe.Pointer) {
	// 0: (, const QTime & time), (time)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QDateTimeEdit7setTimeERK5QTime", ffiqt.FFI_TYPE_VOID, this.cthis, time)
	gopp.ErrPrint(err, rv)
}

//  body block end
