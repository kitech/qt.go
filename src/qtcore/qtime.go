//  header block begin
// /usr/include/qt/QtCore/qdatetime.h
// #include <qdatetime.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QTime struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qdatetime.h:159
// index:0
// inline
// void QTime()
func NewQTime() *QTime {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTimeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTime{cthis}
}

// /usr/include/qt/QtCore/qdatetime.h:161
// index:1
// void QTime(int, int, int, int)
func NewQTime_1(h int, m int, s int, ms int) *QTime {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTimeC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, &h, &m, &s, &ms)
	gopp.ErrPrint(err, rv)
	return &QTime{cthis}
}

// /usr/include/qt/QtCore/qdatetime.h:163
// index:0
// inline
// bool isNull()
func (this *QTime) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:164
// index:0
// bool isValid()
func (this *QTime) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:199
// index:1
// static
// bool isValid(int, int, int, int)
func (this *QTime) IsValid_1(h int, m int, s int, ms int) {
	// 1: (int h, int m, int s, int ms), (h, m, s, ms)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime7isValidEiiii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTime_IsValid_1(h int, m int, s int, ms int) {
	// 1: (int h, int m, int s, int ms), (h, m, s, ms)
	var nilthis *QTime
	nilthis.IsValid_1(h, m, s, ms)
}

// /usr/include/qt/QtCore/qdatetime.h:166
// index:0
// int hour()
func (this *QTime) Hour() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime4hourEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:167
// index:0
// int minute()
func (this *QTime) Minute() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime6minuteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:168
// index:0
// int second()
func (this *QTime) Second() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime6secondEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:169
// index:0
// int msec()
func (this *QTime) Msec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime4msecEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:171
// index:0
// QString toString(Qt::DateFormat)
func (this *QTime) ToString(f int) {
	// 0: (, Qt::DateFormat f), (&f)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime8toStringEN2Qt10DateFormatE", ffiqt.FFI_TYPE_VOID, this.cthis, &f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:173
// index:1
// QString toString(const class QString &)
func (this *QTime) ToString_1(format unsafe.Pointer) {
	// 1: (, const QString & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime8toStringERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:175
// index:2
// QString toString(class QStringView)
func (this *QTime) ToString_2(format unsafe.Pointer) {
	// 2: (, QStringView format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime8toStringE11QStringView", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:177
// index:0
// bool setHMS(int, int, int, int)
func (this *QTime) SetHMS(h int, m int, s int, ms int) {
	// 0: (, int h, int m, int s, int ms), (&h, &m, &s, &ms)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime6setHMSEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &h, &m, &s, &ms)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:179
// index:0
// QTime addSecs(int)
func (this *QTime) AddSecs(secs int) {
	// 0: (, int secs), (&secs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime7addSecsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &secs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:180
// index:0
// int secsTo(const class QTime &)
func (this *QTime) SecsTo(arg0 unsafe.Pointer) {
	// 0: (, const QTime & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime6secsToERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:181
// index:0
// QTime addMSecs(int)
func (this *QTime) AddMSecs(ms int) {
	// 0: (, int ms), (&ms)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime8addMSecsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &ms)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:182
// index:0
// int msecsTo(const class QTime &)
func (this *QTime) MsecsTo(arg0 unsafe.Pointer) {
	// 0: (, const QTime & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime7msecsToERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:191
// index:0
// static inline
// QTime fromMSecsSinceStartOfDay(int)
func (this *QTime) FromMSecsSinceStartOfDay(msecs int) {
	// 0: (int msecs), (msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime24fromMSecsSinceStartOfDayEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTime_FromMSecsSinceStartOfDay(msecs int) {
	// 0: (int msecs), (msecs)
	var nilthis *QTime
	nilthis.FromMSecsSinceStartOfDay(msecs)
}

// /usr/include/qt/QtCore/qdatetime.h:192
// index:0
// inline
// int msecsSinceStartOfDay()
func (this *QTime) MsecsSinceStartOfDay() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime20msecsSinceStartOfDayEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:194
// index:0
// static
// QTime currentTime()
func (this *QTime) CurrentTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime11currentTimeEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTime_CurrentTime() {
	// 0: (), ()
	var nilthis *QTime
	nilthis.CurrentTime()
}

// /usr/include/qt/QtCore/qdatetime.h:196
// index:0
// static
// QTime fromString(const class QString &, Qt::DateFormat)
func (this *QTime) FromString(s unsafe.Pointer, f int) {
	// 0: (const QString & s, Qt::DateFormat f), (s, f)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime10fromStringERK7QStringN2Qt10DateFormatE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTime_FromString(s unsafe.Pointer, f int) {
	// 0: (const QString & s, Qt::DateFormat f), (s, f)
	var nilthis *QTime
	nilthis.FromString(s, f)
}

// /usr/include/qt/QtCore/qdatetime.h:197
// index:1
// static
// QTime fromString(const class QString &, const class QString &)
func (this *QTime) FromString_1(s unsafe.Pointer, format unsafe.Pointer) {
	// 1: (const QString & s, const QString & format), (s, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime10fromStringERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTime_FromString_1(s unsafe.Pointer, format unsafe.Pointer) {
	// 1: (const QString & s, const QString & format), (s, format)
	var nilthis *QTime
	nilthis.FromString_1(s, format)
}

// /usr/include/qt/QtCore/qdatetime.h:201
// index:0
// void start()
func (this *QTime) Start() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime5startEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:202
// index:0
// int restart()
func (this *QTime) Restart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime7restartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:203
// index:0
// int elapsed()
func (this *QTime) Elapsed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime7elapsedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
