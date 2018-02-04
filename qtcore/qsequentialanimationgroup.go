package qtcore

// /usr/include/qt/QtCore/qsequentialanimationgroup.h
// #include <qsequentialanimationgroup.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
// bool event(class QEvent *)
func (this *QSequentialAnimationGroup) InheritEvent(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void updateCurrentTime(int)
func (this *QSequentialAnimationGroup) InheritUpdateCurrentTime(f func(arg0 int)) {
	qtrt.SetAllInheritCallback(this, "updateCurrentTime", f)
}

// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QSequentialAnimationGroup) InheritUpdateState(f func(newState int, oldState int)) {
	qtrt.SetAllInheritCallback(this, "updateState", f)
}

// void updateDirection(class QAbstractAnimation::Direction)
func (this *QSequentialAnimationGroup) InheritUpdateDirection(f func(direction int)) {
	qtrt.SetAllInheritCallback(this, "updateDirection", f)
}

type QSequentialAnimationGroup struct {
	*QAnimationGroup
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
// [8] const QMetaObject * metaObject()
func (this *QSequentialAnimationGroup) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSequentialAnimationGroup(QObject *)
func NewQSequentialAnimationGroup(parent *QObject /*777 QObject **/) *QSequentialAnimationGroup {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSequentialAnimationGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSequentialAnimationGroup()
func DeleteQSequentialAnimationGroup(this *QSequentialAnimationGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QPauseAnimation * addPause(int)
func (this *QSequentialAnimationGroup) AddPause(msecs int) *QPauseAnimation /*777 QPauseAnimation **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup8addPauseEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPauseAnimationFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QPauseAnimation * insertPause(int, int)
func (this *QSequentialAnimationGroup) InsertPause(index int, msecs int) *QPauseAnimation /*777 QPauseAnimation **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup11insertPauseEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPauseAnimationFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractAnimation * currentAnimation()
func (this *QSequentialAnimationGroup) CurrentAnimation() *QAbstractAnimation /*777 QAbstractAnimation **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup16currentAnimationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractAnimationFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int duration()
func (this *QSequentialAnimationGroup) Duration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentAnimationChanged(QAbstractAnimation *)
func (this *QSequentialAnimationGroup) CurrentAnimationChanged(current *QAbstractAnimation /*777 QAbstractAnimation **/) {
	var convArg0 = current.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup23currentAnimationChangedEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:73
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QSequentialAnimationGroup) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:75
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateCurrentTime(int)
func (this *QSequentialAnimationGroup) UpdateCurrentTime(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup17updateCurrentTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:76
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateState(QAbstractAnimation::State, QAbstractAnimation::State)
func (this *QSequentialAnimationGroup) UpdateState(newState int, oldState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:77
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateDirection(QAbstractAnimation::Direction)
func (this *QSequentialAnimationGroup) UpdateDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

//  body block end
