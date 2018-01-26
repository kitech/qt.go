package qtcore

// /usr/include/qt/QtCore/qdatetime.h
// #include <qdatetime.h>
// #include <QtCore>

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
	*qtrt.CObject
}

func (this *QTime) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTime) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTimeFromPointer(cthis unsafe.Pointer) *QTime {
	return &QTime{&qtrt.CObject{cthis}}
}
func (*QTime) NewFromPointer(cthis unsafe.Pointer) *QTime {
	return NewQTimeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qdatetime.h:159
// index:0
// Public inline
// void QTime()
func NewQTime() *QTime {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTimeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:161
// index:1
// Public
// void QTime(int, int, int, int)
func NewQTime_1(h int, m int, s int, ms int) *QTime {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTimeC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, h, m, s, ms)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatetime.h:163
// index:0
// Public inline
// bool isNull()
func (this *QTime) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:164
// index:0
// Public
// bool isValid()
func (this *QTime) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:199
// index:1
// Public static
// bool isValid(int, int, int, int)
func (this *QTime) IsValid_1(h int, m int, s int, ms int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime7isValidEiiii", ffiqt.FFI_TYPE_POINTER, h, m, s, ms)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QTime_IsValid_1(h int, m int, s int, ms int) bool {
	var nilthis *QTime
	rv := nilthis.IsValid_1(h, m, s, ms)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:166
// index:0
// Public
// int hour()
func (this *QTime) Hour() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime4hourEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatetime.h:167
// index:0
// Public
// int minute()
func (this *QTime) Minute() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime6minuteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatetime.h:168
// index:0
// Public
// int second()
func (this *QTime) Second() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime6secondEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatetime.h:169
// index:0
// Public
// int msec()
func (this *QTime) Msec() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime4msecEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatetime.h:171
// index:0
// Public
// QString toString(Qt::DateFormat)
func (this *QTime) ToString(f int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime8toStringEN2Qt10DateFormatE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), f)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:173
// index:1
// Public
// QString toString(const class QString &)
func (this *QTime) ToString_1(format *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime8toStringERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:175
// index:2
// Public
// QString toString(class QStringView)
func (this *QTime) ToString_2(format *QStringView /*123*/) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime8toStringE11QStringView", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:177
// index:0
// Public
// bool setHMS(int, int, int, int)
func (this *QTime) SetHMS(h int, m int, s int, ms int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime6setHMSEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), h, m, s, ms)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdatetime.h:179
// index:0
// Public
// QTime addSecs(int)
func (this *QTime) AddSecs(secs int) *QTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime7addSecsEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), secs)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:180
// index:0
// Public
// int secsTo(const class QTime &)
func (this *QTime) SecsTo(arg0 *QTime) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime6secsToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatetime.h:181
// index:0
// Public
// QTime addMSecs(int)
func (this *QTime) AddMSecs(ms int) *QTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime8addMSecsEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), ms)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qdatetime.h:182
// index:0
// Public
// int msecsTo(const class QTime &)
func (this *QTime) MsecsTo(arg0 *QTime) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime7msecsToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatetime.h:191
// index:0
// Public static inline
// QTime fromMSecsSinceStartOfDay(int)
func (this *QTime) FromMSecsSinceStartOfDay(msecs int) *QTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime24fromMSecsSinceStartOfDayEi", ffiqt.FFI_TYPE_POINTER, msecs)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QTime_FromMSecsSinceStartOfDay(msecs int) *QTime /*123*/ {
	var nilthis *QTime
	rv := nilthis.FromMSecsSinceStartOfDay(msecs)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:192
// index:0
// Public inline
// int msecsSinceStartOfDay()
func (this *QTime) MsecsSinceStartOfDay() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime20msecsSinceStartOfDayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatetime.h:194
// index:0
// Public static
// QTime currentTime()
func (this *QTime) CurrentTime() *QTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime11currentTimeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QTime_CurrentTime() *QTime /*123*/ {
	var nilthis *QTime
	rv := nilthis.CurrentTime()
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:196
// index:0
// Public static
// QTime fromString(const class QString &, Qt::DateFormat)
func (this *QTime) FromString(s *QString, f int) *QTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime10fromStringERK7QStringN2Qt10DateFormatE", ffiqt.FFI_TYPE_POINTER, s, f)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QTime_FromString(s *QString, f int) *QTime /*123*/ {
	var nilthis *QTime
	rv := nilthis.FromString(s, f)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:197
// index:1
// Public static
// QTime fromString(const class QString &, const class QString &)
func (this *QTime) FromString_1(s *QString, format *QString) *QTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime10fromStringERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, s, format)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QTime_FromString_1(s *QString, format *QString) *QTime /*123*/ {
	var nilthis *QTime
	rv := nilthis.FromString_1(s, format)
	return rv
}

// /usr/include/qt/QtCore/qdatetime.h:201
// index:0
// Public
// void start()
func (this *QTime) Start() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime5startEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatetime.h:202
// index:0
// Public
// int restart()
func (this *QTime) Restart() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QTime7restartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatetime.h:203
// index:0
// Public
// int elapsed()
func (this *QTime) Elapsed() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QTime7elapsedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

type QTime__TimeFlag = int

const QTime__NullTime QTime__TimeFlag = 4294967295

//  body block end
