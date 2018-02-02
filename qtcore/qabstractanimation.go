package qtcore

// /usr/include/qt/QtCore/qabstractanimation.h
// #include <qabstractanimation.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/cffiqt"
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
// bool event(class QEvent *)
func (this *QAbstractAnimation) InheritEvent(f func(event *QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

// void updateCurrentTime(int)
func (this *QAbstractAnimation) InheritUpdateCurrentTime(f func(currentTime int)) {
	ffiqt.SetAllInheritCallback(this, "updateCurrentTime", f)
}

// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QAbstractAnimation) InheritUpdateState(f func(newState int, oldState int)) {
	ffiqt.SetAllInheritCallback(this, "updateState", f)
}

// void updateDirection(class QAbstractAnimation::Direction)
func (this *QAbstractAnimation) InheritUpdateDirection(f func(direction int)) {
	ffiqt.SetAllInheritCallback(this, "updateDirection", f)
}

type QAbstractAnimation struct {
	*QObject
}

func (this *QAbstractAnimation) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractAnimation) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQAbstractAnimationFromPointer(cthis unsafe.Pointer) *QAbstractAnimation {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractAnimation{bcthis0}
}
func (*QAbstractAnimation) NewFromPointer(cthis unsafe.Pointer) *QAbstractAnimation {
	return NewQAbstractAnimationFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractanimation.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAbstractAnimation) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractanimation.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractAnimation(QObject *)
func NewQAbstractAnimation(parent *QObject /*777 QObject **/) *QAbstractAnimation {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimationC1EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qabstractanimation.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractAnimation()
func DeleteQAbstractAnimation(this *QAbstractAnimation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimationD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractanimation.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractAnimation::State state()
func (this *QAbstractAnimation) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QAnimationGroup * group()
func (this *QAbstractAnimation) Group() *QAnimationGroup /*777 QAnimationGroup **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation5groupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAnimationGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractanimation.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractAnimation::Direction direction()
func (this *QAbstractAnimation) Direction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation9directionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDirection(enum QAbstractAnimation::Direction)
func (this *QAbstractAnimation) SetDirection(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation12setDirectionENS_9DirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentTime()
func (this *QAbstractAnimation) CurrentTime() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation11currentTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentLoopTime()
func (this *QAbstractAnimation) CurrentLoopTime() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation15currentLoopTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopCount()
func (this *QAbstractAnimation) LoopCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation9loopCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLoopCount(int)
func (this *QAbstractAnimation) SetLoopCount(loopCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation12setLoopCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), loopCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentLoop()
func (this *QAbstractAnimation) CurrentLoop() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation11currentLoopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:102
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int duration()
func (this *QAbstractAnimation) Duration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation8durationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int totalDuration()
func (this *QAbstractAnimation) TotalDuration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractAnimation13totalDurationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()
func (this *QAbstractAnimation) Finished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation8finishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QAbstractAnimation::State, QAbstractAnimation::State)
func (this *QAbstractAnimation) StateChanged(newState int, oldState int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation12stateChangedENS_5StateES0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentLoopChanged(int)
func (this *QAbstractAnimation) CurrentLoopChanged(currentLoop int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation18currentLoopChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), currentLoop)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void directionChanged(QAbstractAnimation::Direction)
func (this *QAbstractAnimation) DirectionChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation16directionChangedENS_9DirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(QAbstractAnimation::DeletionPolicy)
func (this *QAbstractAnimation) Start(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation5startENS_14DeletionPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pause()
func (this *QAbstractAnimation) Pause() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation5pauseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resume()
func (this *QAbstractAnimation) Resume() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation6resumeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaused(_Bool)
func (this *QAbstractAnimation) SetPaused(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation9setPausedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()
func (this *QAbstractAnimation) Stop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation4stopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentTime(int)
func (this *QAbstractAnimation) SetCurrentTime(msecs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation14setCurrentTimeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:121
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAbstractAnimation) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractanimation.h:123
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void updateCurrentTime(int)
func (this *QAbstractAnimation) UpdateCurrentTime(currentTime int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation17updateCurrentTimeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), currentTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:124
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateState(QAbstractAnimation::State, QAbstractAnimation::State)
func (this *QAbstractAnimation) UpdateState(newState int, oldState int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation11updateStateENS_5StateES0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:125
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateDirection(QAbstractAnimation::Direction)
func (this *QAbstractAnimation) UpdateDirection(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractAnimation15updateDirectionENS_9DirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

type QAbstractAnimation__Direction = int

const QAbstractAnimation__Forward QAbstractAnimation__Direction = 0
const QAbstractAnimation__Backward QAbstractAnimation__Direction = 1

type QAbstractAnimation__State = int

const QAbstractAnimation__Stopped QAbstractAnimation__State = 0
const QAbstractAnimation__Paused QAbstractAnimation__State = 1
const QAbstractAnimation__Running QAbstractAnimation__State = 2

type QAbstractAnimation__DeletionPolicy = int

const QAbstractAnimation__KeepWhenStopped QAbstractAnimation__DeletionPolicy = 0
const QAbstractAnimation__DeleteWhenStopped QAbstractAnimation__DeletionPolicy = 1

//  body block end
