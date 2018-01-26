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
import "mkuse/cffiqt"
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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QTimeLine) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qtimeline.h:79
// index:0
// Public
// void QTimeLine(int, class QObject *)
func NewQTimeLine(duration int, parent *QObject /*777 QObject **/) *QTimeLine {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLineC2EiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, duration, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeLineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtimeline.h:80
// index:0
// Public virtual
// void ~QTimeLine()
func DeleteQTimeLine(*QTimeLine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:82
// index:0
// Public
// QTimeLine::State state()
func (this *QTimeLine) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtimeline.h:84
// index:0
// Public
// int loopCount()
func (this *QTimeLine) LoopCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine9loopCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimeline.h:85
// index:0
// Public
// void setLoopCount(int)
func (this *QTimeLine) SetLoopCount(count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine12setLoopCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:87
// index:0
// Public
// QTimeLine::Direction direction()
func (this *QTimeLine) Direction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine9directionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtimeline.h:88
// index:0
// Public
// void setDirection(enum QTimeLine::Direction)
func (this *QTimeLine) SetDirection(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine12setDirectionENS_9DirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:90
// index:0
// Public
// int duration()
func (this *QTimeLine) Duration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine8durationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimeline.h:91
// index:0
// Public
// void setDuration(int)
func (this *QTimeLine) SetDuration(duration int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine11setDurationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), duration)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:93
// index:0
// Public
// int startFrame()
func (this *QTimeLine) StartFrame() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine10startFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimeline.h:94
// index:0
// Public
// void setStartFrame(int)
func (this *QTimeLine) SetStartFrame(frame int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine13setStartFrameEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), frame)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:95
// index:0
// Public
// int endFrame()
func (this *QTimeLine) EndFrame() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine8endFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimeline.h:96
// index:0
// Public
// void setEndFrame(int)
func (this *QTimeLine) SetEndFrame(frame int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine11setEndFrameEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), frame)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:97
// index:0
// Public
// void setFrameRange(int, int)
func (this *QTimeLine) SetFrameRange(startFrame int, endFrame int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine13setFrameRangeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), startFrame, endFrame)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:99
// index:0
// Public
// int updateInterval()
func (this *QTimeLine) UpdateInterval() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine14updateIntervalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimeline.h:100
// index:0
// Public
// void setUpdateInterval(int)
func (this *QTimeLine) SetUpdateInterval(interval int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine17setUpdateIntervalEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:102
// index:0
// Public
// QTimeLine::CurveShape curveShape()
func (this *QTimeLine) CurveShape() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine10curveShapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtimeline.h:103
// index:0
// Public
// void setCurveShape(enum QTimeLine::CurveShape)
func (this *QTimeLine) SetCurveShape(shape int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine13setCurveShapeENS_10CurveShapeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), shape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:105
// index:0
// Public
// QEasingCurve easingCurve()
func (this *QTimeLine) EasingCurve() *QEasingCurve /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine11easingCurveEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQEasingCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtimeline.h:106
// index:0
// Public
// void setEasingCurve(const class QEasingCurve &)
func (this *QTimeLine) SetEasingCurve(curve *QEasingCurve) {
	var convArg0 = curve.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine14setEasingCurveERK12QEasingCurve", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:108
// index:0
// Public
// int currentTime()
func (this *QTimeLine) CurrentTime() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine11currentTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimeline.h:109
// index:0
// Public
// int currentFrame()
func (this *QTimeLine) CurrentFrame() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine12currentFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimeline.h:110
// index:0
// Public
// qreal currentValue()
func (this *QTimeLine) CurrentValue() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine12currentValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qtimeline.h:112
// index:0
// Public
// int frameForTime(int)
func (this *QTimeLine) FrameForTime(msec int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine12frameForTimeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimeline.h:113
// index:0
// Public virtual
// qreal valueForTime(int)
func (this *QTimeLine) ValueForTime(msec int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine12valueForTimeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtCore/qtimeline.h:116
// index:0
// Public
// void start()
func (this *QTimeLine) Start() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine5startEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:117
// index:0
// Public
// void resume()
func (this *QTimeLine) Resume() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine6resumeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:118
// index:0
// Public
// void stop()
func (this *QTimeLine) Stop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine4stopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:119
// index:0
// Public
// void setPaused(_Bool)
func (this *QTimeLine) SetPaused(paused bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine9setPausedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:120
// index:0
// Public
// void setCurrentTime(int)
func (this *QTimeLine) SetCurrentTime(msec int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine14setCurrentTimeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:121
// index:0
// Public
// void toggleDirection()
func (this *QTimeLine) ToggleDirection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine15toggleDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:130
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QTimeLine) TimerEvent(event *QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
