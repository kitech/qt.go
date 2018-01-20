//  header block begin
// /usr/include/qt/QtCore/qtimeline.h
// #include <qtimeline.h>
// #include <QtCore>
package qtcore

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
type QTimeLine struct {
	*QObject
}

func (this *QTimeLine) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtCore/qtimeline.h:52
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTimeLine) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:79
// index:0
// void QTimeLine(int, class QObject *)
func NewQTimeLine(duration int, parent unsafe.Pointer) *QTimeLine {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLineC2EiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &duration, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeLineFromPointer(cthis)
	return gothis
}
func NewQTimeLineFromPointer(cthis unsafe.Pointer) *QTimeLine {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QTimeLine{bcthis0}
}

// /usr/include/qt/QtCore/qtimeline.h:80
// index:0
// virtual
// void ~QTimeLine()
func DeleteQTimeLine(*QTimeLine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:82
// index:0
// QTimeLine::State state()
func (this *QTimeLine) State() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine5stateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:84
// index:0
// int loopCount()
func (this *QTimeLine) LoopCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine9loopCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:85
// index:0
// void setLoopCount(int)
func (this *QTimeLine) SetLoopCount(count int) {
	// 0: (, count int), (&count)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine12setLoopCountEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:87
// index:0
// QTimeLine::Direction direction()
func (this *QTimeLine) Direction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine9directionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:88
// index:0
// void setDirection(enum QTimeLine::Direction)
func (this *QTimeLine) SetDirection(direction int) {
	// 0: (, direction QTimeLine::Direction), (&direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine12setDirectionENS_9DirectionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:90
// index:0
// int duration()
func (this *QTimeLine) Duration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine8durationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:91
// index:0
// void setDuration(int)
func (this *QTimeLine) SetDuration(duration int) {
	// 0: (, duration int), (&duration)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine11setDurationEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &duration)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:93
// index:0
// int startFrame()
func (this *QTimeLine) StartFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine10startFrameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:94
// index:0
// void setStartFrame(int)
func (this *QTimeLine) SetStartFrame(frame int) {
	// 0: (, frame int), (&frame)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine13setStartFrameEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &frame)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:95
// index:0
// int endFrame()
func (this *QTimeLine) EndFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine8endFrameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:96
// index:0
// void setEndFrame(int)
func (this *QTimeLine) SetEndFrame(frame int) {
	// 0: (, frame int), (&frame)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine11setEndFrameEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &frame)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:97
// index:0
// void setFrameRange(int, int)
func (this *QTimeLine) SetFrameRange(startFrame int, endFrame int) {
	// 0: (, startFrame int, endFrame int), (&startFrame, &endFrame)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine13setFrameRangeEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &startFrame, &endFrame)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:99
// index:0
// int updateInterval()
func (this *QTimeLine) UpdateInterval() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine14updateIntervalEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:100
// index:0
// void setUpdateInterval(int)
func (this *QTimeLine) SetUpdateInterval(interval int) {
	// 0: (, interval int), (&interval)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine17setUpdateIntervalEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:102
// index:0
// QTimeLine::CurveShape curveShape()
func (this *QTimeLine) CurveShape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine10curveShapeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:103
// index:0
// void setCurveShape(enum QTimeLine::CurveShape)
func (this *QTimeLine) SetCurveShape(shape int) {
	// 0: (, shape QTimeLine::CurveShape), (&shape)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine13setCurveShapeENS_10CurveShapeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &shape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:105
// index:0
// QEasingCurve easingCurve()
func (this *QTimeLine) EasingCurve() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine11easingCurveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:106
// index:0
// void setEasingCurve(const class QEasingCurve &)
func (this *QTimeLine) SetEasingCurve(curve unsafe.Pointer) {
	// 0: (, curve const QEasingCurve &), (curve)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine14setEasingCurveERK12QEasingCurve", ffiqt.FFI_TYPE_VOID, this.GetCthis(), curve)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:108
// index:0
// int currentTime()
func (this *QTimeLine) CurrentTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine11currentTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:109
// index:0
// int currentFrame()
func (this *QTimeLine) CurrentFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine12currentFrameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:110
// index:0
// qreal currentValue()
func (this *QTimeLine) CurrentValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine12currentValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:112
// index:0
// int frameForTime(int)
func (this *QTimeLine) FrameForTime(msec int) {
	// 0: (, msec int), (&msec)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine12frameForTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:113
// index:0
// virtual
// qreal valueForTime(int)
func (this *QTimeLine) ValueForTime(msec int) {
	// 0: (, msec int), (&msec)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeLine12valueForTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:116
// index:0
// void start()
func (this *QTimeLine) Start() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine5startEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:117
// index:0
// void resume()
func (this *QTimeLine) Resume() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine6resumeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:118
// index:0
// void stop()
func (this *QTimeLine) Stop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine4stopEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:119
// index:0
// void setPaused(_Bool)
func (this *QTimeLine) SetPaused(paused bool) {
	// 0: (, paused bool), (&paused)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine9setPausedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:120
// index:0
// void setCurrentTime(int)
func (this *QTimeLine) SetCurrentTime(msec int) {
	// 0: (, msec int), (&msec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine14setCurrentTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:121
// index:0
// void toggleDirection()
func (this *QTimeLine) ToggleDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine15toggleDirectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimeline.h:130
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QTimeLine) TimerEvent(event unsafe.Pointer) {
	// 0: (, event QTimerEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeLine10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

//  body block end
