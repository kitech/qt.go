//  header block begin
// /usr/include/qt/QtCore/qabstractanimation.h
// #include <qabstractanimation.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QAbstractAnimation struct {
	*QObject
}

func (this *QAbstractAnimation) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtCore/qabstractanimation.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractAnimation) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:85
// index:0
// void QAbstractAnimation(class QObject *)
func NewQAbstractAnimation(parent unsafe.Pointer) *QAbstractAnimation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimationC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractAnimationFromPointer(cthis)
	return gothis
}
func NewQAbstractAnimationFromPointer(cthis unsafe.Pointer) *QAbstractAnimation {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractAnimation{bcthis0}
}

// /usr/include/qt/QtCore/qabstractanimation.h:120
// index:1
// void QAbstractAnimation(class QAbstractAnimationPrivate &, class QObject *)
func NewQAbstractAnimation_1(dd unsafe.Pointer, parent unsafe.Pointer) *QAbstractAnimation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimationC1ER25QAbstractAnimationPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractAnimationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractanimation.h:86
// index:0
// virtual
// void ~QAbstractAnimation()
func DeleteQAbstractAnimation(*QAbstractAnimation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:88
// index:0
// QAbstractAnimation::State state()
func (this *QAbstractAnimation) State() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation5stateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:90
// index:0
// QAnimationGroup * group()
func (this *QAbstractAnimation) Group() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation5groupEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:92
// index:0
// QAbstractAnimation::Direction direction()
func (this *QAbstractAnimation) Direction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation9directionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:93
// index:0
// void setDirection(enum QAbstractAnimation::Direction)
func (this *QAbstractAnimation) SetDirection(direction int) {
	// 0: (, direction QAbstractAnimation::Direction), (&direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation12setDirectionENS_9DirectionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:95
// index:0
// int currentTime()
func (this *QAbstractAnimation) CurrentTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation11currentTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:96
// index:0
// int currentLoopTime()
func (this *QAbstractAnimation) CurrentLoopTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation15currentLoopTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:98
// index:0
// int loopCount()
func (this *QAbstractAnimation) LoopCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation9loopCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:99
// index:0
// void setLoopCount(int)
func (this *QAbstractAnimation) SetLoopCount(loopCount int) {
	// 0: (, loopCount int), (&loopCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation12setLoopCountEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &loopCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:100
// index:0
// int currentLoop()
func (this *QAbstractAnimation) CurrentLoop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation11currentLoopEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:102
// index:0
// pure virtual
// int duration()
func (this *QAbstractAnimation) Duration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation8durationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:103
// index:0
// int totalDuration()
func (this *QAbstractAnimation) TotalDuration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation13totalDurationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:106
// index:0
// void finished()
func (this *QAbstractAnimation) Finished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation8finishedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:107
// index:0
// void stateChanged(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QAbstractAnimation) StateChanged(newState int, oldState int) {
	// 0: (, newState QAbstractAnimation::State, oldState QAbstractAnimation::State), (&newState, &oldState)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation12stateChangedENS_5StateES0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &newState, &oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:108
// index:0
// void currentLoopChanged(int)
func (this *QAbstractAnimation) CurrentLoopChanged(currentLoop int) {
	// 0: (, currentLoop int), (&currentLoop)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation18currentLoopChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &currentLoop)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:109
// index:0
// void directionChanged(class QAbstractAnimation::Direction)
func (this *QAbstractAnimation) DirectionChanged(arg0 int) {
	// 0: (, QAbstractAnimation::Direction arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation16directionChangedENS_9DirectionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:112
// index:0
// void start(class QAbstractAnimation::DeletionPolicy)
func (this *QAbstractAnimation) Start(policy int) {
	// 0: (, policy QAbstractAnimation::DeletionPolicy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation5startENS_14DeletionPolicyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:113
// index:0
// void pause()
func (this *QAbstractAnimation) Pause() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation5pauseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:114
// index:0
// void resume()
func (this *QAbstractAnimation) Resume() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation6resumeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:115
// index:0
// void setPaused(_Bool)
func (this *QAbstractAnimation) SetPaused(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation9setPausedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:116
// index:0
// void stop()
func (this *QAbstractAnimation) Stop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation4stopEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:117
// index:0
// void setCurrentTime(int)
func (this *QAbstractAnimation) SetCurrentTime(msecs int) {
	// 0: (, msecs int), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation14setCurrentTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:121
// index:0
// virtual
// bool event(class QEvent *)
func (this *QAbstractAnimation) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:123
// index:0
// pure virtual
// void updateCurrentTime(int)
func (this *QAbstractAnimation) UpdateCurrentTime(currentTime int) {
	// 0: (, currentTime int), (&currentTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation17updateCurrentTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &currentTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:124
// index:0
// virtual
// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QAbstractAnimation) UpdateState(newState int, oldState int) {
	// 0: (, newState QAbstractAnimation::State, oldState QAbstractAnimation::State), (&newState, &oldState)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation11updateStateENS_5StateES0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &newState, &oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:125
// index:0
// virtual
// void updateDirection(class QAbstractAnimation::Direction)
func (this *QAbstractAnimation) UpdateDirection(direction int) {
	// 0: (, direction QAbstractAnimation::Direction), (&direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation15updateDirectionENS_9DirectionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &direction)
	gopp.ErrPrint(err, rv)
}

//  body block end
