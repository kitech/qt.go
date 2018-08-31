package qtcore

// /usr/include/qt/QtCore/qsequentialanimationgroup.h
// #include <qsequentialanimationgroup.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
func (this *QSequentialAnimationGroup) InheritEvent(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void updateCurrentTime(int)
func (this *QSequentialAnimationGroup) InheritUpdateCurrentTime(f func(arg0 int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCurrentTime", f)
}

// void updateState(QAbstractAnimation::State, QAbstractAnimation::State)
func (this *QSequentialAnimationGroup) InheritUpdateState(f func(newState int, oldState int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateState", f)
}

// void updateDirection(QAbstractAnimation::Direction)
func (this *QSequentialAnimationGroup) InheritUpdateDirection(f func(direction int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateDirection", f)
}

/*

 */
type QSequentialAnimationGroup struct {
	*QAnimationGroup
}
type QSequentialAnimationGroup_ITF interface {
	QAnimationGroup_ITF
	QSequentialAnimationGroup_PTR() *QSequentialAnimationGroup
}

func (ptr *QSequentialAnimationGroup) QSequentialAnimationGroup_PTR() *QSequentialAnimationGroup {
	return ptr
}

func (this *QSequentialAnimationGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAnimationGroup.GetCthis()
	}
}
func (this *QSequentialAnimationGroup) SetCthis(cthis unsafe.Pointer) {
	this.QAnimationGroup = NewQAnimationGroupFromPointer(cthis)
}
func NewQSequentialAnimationGroupFromPointer(cthis unsafe.Pointer) *QSequentialAnimationGroup {
	bcthis0 := NewQAnimationGroupFromPointer(cthis)
	return &QSequentialAnimationGroup{bcthis0}
}
func (*QSequentialAnimationGroup) NewFromPointer(cthis unsafe.Pointer) *QSequentialAnimationGroup {
	return NewQSequentialAnimationGroupFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSequentialAnimationGroup) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSequentialAnimationGroup(QObject *)

/*
Constructs a QSequentialAnimationGroup. parent is passed to QObject's constructor.
*/
func (*QSequentialAnimationGroup) NewForInherit(parent QObject_ITF /*777 QObject **/) *QSequentialAnimationGroup {
	return NewQSequentialAnimationGroup(parent)
}
func NewQSequentialAnimationGroup(parent QObject_ITF /*777 QObject **/) *QSequentialAnimationGroup {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSequentialAnimationGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSequentialAnimationGroup")
	return gothis
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSequentialAnimationGroup(QObject *)

/*
Constructs a QSequentialAnimationGroup. parent is passed to QObject's constructor.
*/
func (*QSequentialAnimationGroup) NewForInherit__() *QSequentialAnimationGroup {
	return NewQSequentialAnimationGroup__()
}
func NewQSequentialAnimationGroup__() *QSequentialAnimationGroup {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSequentialAnimationGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSequentialAnimationGroup")
	return gothis
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSequentialAnimationGroup()

/*

 */
func DeleteQSequentialAnimationGroup(this *QSequentialAnimationGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QPauseAnimation * addPause(int)

/*
Adds a pause of msecs to this animation group. The pause is considered as a special type of animation, thus animationCount will be increased by one.

See also insertPause() and QAnimationGroup::addAnimation().
*/
func (this *QSequentialAnimationGroup) AddPause(msecs int) *QPauseAnimation /*777 QPauseAnimation **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup8addPauseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPauseAnimationFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QPauseAnimation * insertPause(int, int)

/*
Inserts a pause of msecs milliseconds at index in this animation group.

See also addPause() and QAnimationGroup::insertAnimation().
*/
func (this *QSequentialAnimationGroup) InsertPause(index int, msecs int) *QPauseAnimation /*777 QPauseAnimation **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup11insertPauseEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, msecs)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPauseAnimationFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractAnimation * currentAnimation() const

/*

 */
func (this *QSequentialAnimationGroup) CurrentAnimation() *QAbstractAnimation /*777 QAbstractAnimation **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup16currentAnimationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractAnimationFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int duration() const

/*
Reimplemented from QAbstractAnimation::duration().
*/
func (this *QSequentialAnimationGroup) Duration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentAnimationChanged(QAbstractAnimation *)

/*
QSequentialAnimationGroup emits this signal when currentAnimation has been changed. current is the current animation.

Note: Notifier signal for property currentAnimation.

See also currentAnimation().
*/
func (this *QSequentialAnimationGroup) CurrentAnimationChanged(current QAbstractAnimation_ITF /*777 QAbstractAnimation **/) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QAbstractAnimation_PTR() != nil {
		convArg0 = current.QAbstractAnimation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup23currentAnimationChangedEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:73
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QSequentialAnimationGroup) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:75
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateCurrentTime(int)

/*
Reimplemented from QAbstractAnimation::updateCurrentTime().
*/
func (this *QSequentialAnimationGroup) UpdateCurrentTime(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup17updateCurrentTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:76
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateState(QAbstractAnimation::State, QAbstractAnimation::State)

/*
Reimplemented from QAbstractAnimation::updateState().
*/
func (this *QSequentialAnimationGroup) UpdateState(newState int, oldState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:77
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateDirection(QAbstractAnimation::Direction)

/*
Reimplemented from QAbstractAnimation::updateDirection().
*/
func (this *QSequentialAnimationGroup) UpdateDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

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
