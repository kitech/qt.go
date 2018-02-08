package qtcore

// /usr/include/qt/QtCore/qtimeline.h
// #include <qtimeline.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
// void timerEvent(class QTimerEvent *)
func (this *QTimeLine) InheritTimerEvent(f func(event *QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

type QTimeLine struct {
	*QObject
}

func (this *QTimeLine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QTimeLine) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQTimeLineFromPointer(cthis unsafe.Pointer) *QTimeLine {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QTimeLine{bcthis0}
}
func (*QTimeLine) NewFromPointer(cthis unsafe.Pointer) *QTimeLine {
	return NewQTimeLineFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtimeline.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTimeLine) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtimeline.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTimeLine(int, QObject *)
func NewQTimeLine(duration int, parent *QObject /*777 QObject **/) *QTimeLine {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLineC2EiP7QObject", qtrt.FFI_TYPE_POINTER, duration, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeLineFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qtimeline.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTimeLine()
func DeleteQTimeLine(this *QTimeLine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtimeline.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QTimeLine::State state()
func (this *QTimeLine) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtimeline.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopCount()
func (this *QTimeLine) LoopCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine9loopCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimeline.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLoopCount(int)
func (this *QTimeLine) SetLoopCount(count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine12setLoopCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] QTimeLine::Direction direction()
func (this *QTimeLine) Direction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine9directionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtimeline.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDirection(enum QTimeLine::Direction)
func (this *QTimeLine) SetDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine12setDirectionENS_9DirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int duration()
func (this *QTimeLine) Duration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimeline.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDuration(int)
func (this *QTimeLine) SetDuration(duration int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine11setDurationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), duration)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int startFrame()
func (this *QTimeLine) StartFrame() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine10startFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimeline.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartFrame(int)
func (this *QTimeLine) SetStartFrame(frame int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine13setStartFrameEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frame)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] int endFrame()
func (this *QTimeLine) EndFrame() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine8endFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimeline.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEndFrame(int)
func (this *QTimeLine) SetEndFrame(frame int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine11setEndFrameEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frame)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrameRange(int, int)
func (this *QTimeLine) SetFrameRange(startFrame int, endFrame int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine13setFrameRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startFrame, endFrame)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int updateInterval()
func (this *QTimeLine) UpdateInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine14updateIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimeline.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUpdateInterval(int)
func (this *QTimeLine) SetUpdateInterval(interval int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine17setUpdateIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] QTimeLine::CurveShape curveShape()
func (this *QTimeLine) CurveShape() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine10curveShapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qtimeline.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurveShape(enum QTimeLine::CurveShape)
func (this *QTimeLine) SetCurveShape(shape int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine13setCurveShapeENS_10CurveShapeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), shape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QEasingCurve easingCurve()
func (this *QTimeLine) EasingCurve() *QEasingCurve /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine11easingCurveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQEasingCurve)
	return rv2
}

// /usr/include/qt/QtCore/qtimeline.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEasingCurve(const QEasingCurve &)
func (this *QTimeLine) SetEasingCurve(curve *QEasingCurve) {
	var convArg0 = curve.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine14setEasingCurveERK12QEasingCurve", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentTime()
func (this *QTimeLine) CurrentTime() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine11currentTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimeline.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentFrame()
func (this *QTimeLine) CurrentFrame() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine12currentFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimeline.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal currentValue()
func (this *QTimeLine) CurrentValue() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine12currentValueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qtimeline.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int frameForTime(int)
func (this *QTimeLine) FrameForTime(msec int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine12frameForTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimeline.h:113
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qreal valueForTime(int)
func (this *QTimeLine) ValueForTime(msec int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeLine12valueForTimeEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qtimeline.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start()
func (this *QTimeLine) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resume()
func (this *QTimeLine) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()
func (this *QTimeLine) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaused(_Bool)
func (this *QTimeLine) SetPaused(paused bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine9setPausedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentTime(int)
func (this *QTimeLine) SetCurrentTime(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine14setCurrentTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggleDirection()
func (this *QTimeLine) ToggleDirection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine15toggleDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:130
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QTimeLine) TimerEvent(event *QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeLine10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QTimeLine__State = int

const QTimeLine__NotRunning QTimeLine__State = 0
const QTimeLine__Paused QTimeLine__State = 1
const QTimeLine__Running QTimeLine__State = 2

type QTimeLine__Direction = int

const QTimeLine__Forward QTimeLine__Direction = 0
const QTimeLine__Backward QTimeLine__Direction = 1

type QTimeLine__CurveShape = int

const QTimeLine__EaseInCurve QTimeLine__CurveShape = 0
const QTimeLine__EaseOutCurve QTimeLine__CurveShape = 1
const QTimeLine__EaseInOutCurve QTimeLine__CurveShape = 2
const QTimeLine__LinearCurve QTimeLine__CurveShape = 3
const QTimeLine__SineCurve QTimeLine__CurveShape = 4
const QTimeLine__CosineCurve QTimeLine__CurveShape = 5

//  body block end
