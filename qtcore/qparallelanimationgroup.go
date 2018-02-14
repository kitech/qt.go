package qtcore

// /usr/include/qt/QtCore/qparallelanimationgroup.h
// #include <qparallelanimationgroup.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QParallelAnimationGroup) InheritEvent(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void updateCurrentTime(int)
func (this *QParallelAnimationGroup) InheritUpdateCurrentTime(f func(currentTime int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCurrentTime", f)
}

// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QParallelAnimationGroup) InheritUpdateState(f func(newState int, oldState int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateState", f)
}

// void updateDirection(class QAbstractAnimation::Direction)
func (this *QParallelAnimationGroup) InheritUpdateDirection(f func(direction int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateDirection", f)
}

type QParallelAnimationGroup struct {
	*QAnimationGroup
}
type QParallelAnimationGroup_ITF interface {
	QAnimationGroup_ITF
	QParallelAnimationGroup_PTR() *QParallelAnimationGroup
}

func (ptr *QParallelAnimationGroup) QParallelAnimationGroup_PTR() *QParallelAnimationGroup { return ptr }

func (this *QParallelAnimationGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAnimationGroup.GetCthis()
	}
}
func (this *QParallelAnimationGroup) SetCthis(cthis unsafe.Pointer) {
	this.QAnimationGroup = NewQAnimationGroupFromPointer(cthis)
}
func NewQParallelAnimationGroupFromPointer(cthis unsafe.Pointer) *QParallelAnimationGroup {
	bcthis0 := NewQAnimationGroupFromPointer(cthis)
	return &QParallelAnimationGroup{bcthis0}
}
func (*QParallelAnimationGroup) NewFromPointer(cthis unsafe.Pointer) *QParallelAnimationGroup {
	return NewQParallelAnimationGroupFromPointer(cthis)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QParallelAnimationGroup) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QParallelAnimationGroup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QParallelAnimationGroup(QObject *)
func NewQParallelAnimationGroup(parent QObject_ITF /*777 QObject **/) *QParallelAnimationGroup {
	var convArg0 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QParallelAnimationGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQParallelAnimationGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QParallelAnimationGroup()
func DeleteQParallelAnimationGroup(this *QParallelAnimationGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QParallelAnimationGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int duration()
func (this *QParallelAnimationGroup) Duration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QParallelAnimationGroup8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:63
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QParallelAnimationGroup) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 = event.QEvent_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QParallelAnimationGroup5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:65
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateCurrentTime(int)
func (this *QParallelAnimationGroup) UpdateCurrentTime(currentTime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QParallelAnimationGroup17updateCurrentTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), currentTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:66
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateState(QAbstractAnimation::State, QAbstractAnimation::State)
func (this *QParallelAnimationGroup) UpdateState(newState int, oldState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QParallelAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:67
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateDirection(QAbstractAnimation::Direction)
func (this *QParallelAnimationGroup) UpdateDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QParallelAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
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
		qtrt.KeepMe()
	}
}

//  keep block end
