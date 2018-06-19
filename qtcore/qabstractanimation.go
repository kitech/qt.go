package qtcore

// /usr/include/qt/QtCore/qabstractanimation.h
// #include <qabstractanimation.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QAbstractAnimation) InheritEvent(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void updateCurrentTime(int)
func (this *QAbstractAnimation) InheritUpdateCurrentTime(f func(currentTime int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCurrentTime", f)
}

// void updateState(QAbstractAnimation::State, QAbstractAnimation::State)
func (this *QAbstractAnimation) InheritUpdateState(f func(newState int, oldState int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateState", f)
}

// void updateDirection(QAbstractAnimation::Direction)
func (this *QAbstractAnimation) InheritUpdateDirection(f func(direction int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateDirection", f)
}

/*

 */
type QAbstractAnimation struct {
	*QObject
}
type QAbstractAnimation_ITF interface {
	QObject_ITF
	QAbstractAnimation_PTR() *QAbstractAnimation
}

func (ptr *QAbstractAnimation) QAbstractAnimation_PTR() *QAbstractAnimation { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractAnimation) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractAnimation10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractanimation.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractAnimation(QObject *)

/*
Constructs the QAbstractAnimation base class, and passes parent to QObject's constructor.

See also QVariantAnimation and QAnimationGroup.
*/
func NewQAbstractAnimation(parent QObject_ITF /*777 QObject **/) *QAbstractAnimation {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qabstractanimation.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractAnimation(QObject *)

/*
Constructs the QAbstractAnimation base class, and passes parent to QObject's constructor.

See also QVariantAnimation and QAnimationGroup.
*/
func NewQAbstractAnimation__() *QAbstractAnimation {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qabstractanimation.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractAnimation()

/*

 */
func DeleteQAbstractAnimation(this *QAbstractAnimation) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractanimation.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractAnimation::State state() const

/*

 */
func (this *QAbstractAnimation) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractAnimation5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QAnimationGroup * group() const

/*
If this animation is part of a QAnimationGroup, this function returns a pointer to the group; otherwise, it returns 0.

See also QAnimationGroup::addAnimation().
*/
func (this *QAbstractAnimation) Group() *QAnimationGroup /*777 QAnimationGroup **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractAnimation5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAnimationGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractanimation.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractAnimation::Direction direction() const

/*

 */
func (this *QAbstractAnimation) Direction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractAnimation9directionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDirection(QAbstractAnimation::Direction)

/*

 */
func (this *QAbstractAnimation) SetDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation12setDirectionENS_9DirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentTime() const

/*

 */
func (this *QAbstractAnimation) CurrentTime() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractAnimation11currentTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentLoopTime() const

/*
Returns the current time inside the current loop. It can go from 0 to duration().

See also duration() and currentTime.
*/
func (this *QAbstractAnimation) CurrentLoopTime() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractAnimation15currentLoopTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopCount() const

/*

 */
func (this *QAbstractAnimation) LoopCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractAnimation9loopCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLoopCount(int)

/*

 */
func (this *QAbstractAnimation) SetLoopCount(loopCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation12setLoopCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), loopCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentLoop() const

/*

 */
func (this *QAbstractAnimation) CurrentLoop() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractAnimation11currentLoopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:102
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int duration() const

/*
This pure virtual function returns the duration of the animation, and defines for how long QAbstractAnimation should update the current time. This duration is local, and does not include the loop count.

A return value of -1 indicates that the animation has no defined duration; the animation should run forever until stopped. This is useful for animations that are not time driven, or where you cannot easily predict its duration (e.g., event driven audio playback in a game).

If the animation is a parallel QAnimationGroup, the duration will be the longest duration of all its animations. If the animation is a sequential QAnimationGroup, the duration will be the sum of the duration of all its animations.

Note: Getter function for property duration.

See also loopCount.
*/
func (this *QAbstractAnimation) Duration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractAnimation8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int totalDuration() const

/*
Returns the total and effective duration of the animation, including the loop count.

See also duration() and currentTime.
*/
func (this *QAbstractAnimation) TotalDuration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractAnimation13totalDurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractanimation.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()

/*
QAbstractAnimation emits this signal after the animation has stopped and has reached the end.

This signal is emitted after stateChanged().

See also stateChanged().
*/
func (this *QAbstractAnimation) Finished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation8finishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QAbstractAnimation::State, QAbstractAnimation::State)

/*
QAbstractAnimation emits this signal whenever the state of the animation has changed from oldState to newState. This signal is emitted after the virtual updateState() function is called.

Note: Notifier signal for property state.

See also updateState().
*/
func (this *QAbstractAnimation) StateChanged(newState int, oldState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation12stateChangedENS_5StateES0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentLoopChanged(int)

/*
QAbstractAnimation emits this signal whenever the current loop changes. currentLoop is the current loop.

Note: Notifier signal for property currentLoop.

See also currentLoop() and loopCount().
*/
func (this *QAbstractAnimation) CurrentLoopChanged(currentLoop int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation18currentLoopChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), currentLoop)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void directionChanged(QAbstractAnimation::Direction)

/*
QAbstractAnimation emits this signal whenever the direction has been changed. newDirection is the new direction.

Note: Notifier signal for property direction.

See also direction.
*/
func (this *QAbstractAnimation) DirectionChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation16directionChangedENS_9DirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(QAbstractAnimation::DeletionPolicy)

/*
Starts the animation. The policy argument says whether or not the animation should be deleted when it's done. When the animation starts, the stateChanged() signal is emitted, and state() returns Running. When control reaches the event loop, the animation will run by itself, periodically calling updateCurrentTime() as the animation progresses.

If the animation is currently stopped or has already reached the end, calling start() will rewind the animation and start again from the beginning. When the animation reaches the end, the animation will either stop, or if the loop level is more than 1, it will rewind and continue from the beginning.

If the animation is already running, this function does nothing.

See also stop() and state().
*/
func (this *QAbstractAnimation) Start(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation5startENS_14DeletionPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(QAbstractAnimation::DeletionPolicy)

/*
Starts the animation. The policy argument says whether or not the animation should be deleted when it's done. When the animation starts, the stateChanged() signal is emitted, and state() returns Running. When control reaches the event loop, the animation will run by itself, periodically calling updateCurrentTime() as the animation progresses.

If the animation is currently stopped or has already reached the end, calling start() will rewind the animation and start again from the beginning. When the animation reaches the end, the animation will either stop, or if the loop level is more than 1, it will rewind and continue from the beginning.

If the animation is already running, this function does nothing.

See also stop() and state().
*/
func (this *QAbstractAnimation) Start__() {
	// arg: 0, QAbstractAnimation::DeletionPolicy=Elaborated, QAbstractAnimation::DeletionPolicy=Enum, , Invalid
	policy := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation5startENS_14DeletionPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pause()

/*
Pauses the animation. When the animation is paused, state() returns Paused. The value of currentTime will remain unchanged until resume() or start() is called. If you want to continue from the current time, call resume().

See also start(), state(), and resume().
*/
func (this *QAbstractAnimation) Pause() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation5pauseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resume()

/*
Resumes the animation after it was paused. When the animation is resumed, it emits the resumed() and stateChanged() signals. The currenttime is not changed.

See also start(), pause(), and state().
*/
func (this *QAbstractAnimation) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaused(bool)

/*
If paused is true, the animation is paused. If paused is false, the animation is resumed.

See also state(), pause(), and resume().
*/
func (this *QAbstractAnimation) SetPaused(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation9setPausedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops the animation. When the animation is stopped, it emits the stateChanged() signal, and state() returns Stopped. The current time is not changed.

If the animation stops by itself after reaching the end (i.e., currentLoopTime() == duration() and currentLoop() > loopCount() - 1), the finished() signal is emitted.

See also start() and state().
*/
func (this *QAbstractAnimation) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentTime(int)

/*

 */
func (this *QAbstractAnimation) SetCurrentTime(msecs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation14setCurrentTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:121
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QAbstractAnimation) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractanimation.h:123
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void updateCurrentTime(int)

/*
This pure virtual function is called every time the animation's currentTime changes.

See also updateState().
*/
func (this *QAbstractAnimation) UpdateCurrentTime(currentTime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation17updateCurrentTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), currentTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:124
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateState(QAbstractAnimation::State, QAbstractAnimation::State)

/*
This virtual function is called by QAbstractAnimation when the state of the animation is changed from oldState to newState.

See also start(), stop(), pause(), and resume().
*/
func (this *QAbstractAnimation) UpdateState(newState int, oldState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation11updateStateENS_5StateES0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:125
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateDirection(QAbstractAnimation::Direction)

/*
This virtual function is called by QAbstractAnimation when the direction of the animation is changed. The direction argument is the new direction.

See also setDirection() and direction().
*/
func (this *QAbstractAnimation) UpdateDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractAnimation15updateDirectionENS_9DirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the direction of the animation when in Running state.



See also direction.

*/
type QAbstractAnimation__Direction = int

//
const QAbstractAnimation__Forward QAbstractAnimation__Direction = 0

//
const QAbstractAnimation__Backward QAbstractAnimation__Direction = 1

/*
This enum describes the state of the animation.



See also state() and stateChanged().

*/
type QAbstractAnimation__State = int

// The animation is not running. This is the initial state of QAbstractAnimation, and the state QAbstractAnimation reenters when finished. The current time remain unchanged until either setCurrentTime() is called, or the animation is started by calling start().
const QAbstractAnimation__Stopped QAbstractAnimation__State = 0

// The animation is paused (i.e., temporarily suspended). Calling resume() will resume animation activity.
const QAbstractAnimation__Paused QAbstractAnimation__State = 1

// The animation is running. While control is in the event loop, QAbstractAnimation will update its current time at regular intervals, calling updateCurrentTime() when appropriate.
const QAbstractAnimation__Running QAbstractAnimation__State = 2

/*

 */
type QAbstractAnimation__DeletionPolicy = int

// The animation will not be deleted when stopped.
const QAbstractAnimation__KeepWhenStopped QAbstractAnimation__DeletionPolicy = 0

// The animation will be automatically deleted when stopped.
const QAbstractAnimation__DeleteWhenStopped QAbstractAnimation__DeletionPolicy = 1

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
}

//  keep block end
