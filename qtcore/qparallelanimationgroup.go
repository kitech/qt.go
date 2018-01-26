package qtcore

// /usr/include/qt/QtCore/qparallelanimationgroup.h
// #include <qparallelanimationgroup.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QParallelAnimationGroup struct {
	*QAnimationGroup
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QParallelAnimationGroup) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QParallelAnimationGroup10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:56
// index:0
// Public
// void QParallelAnimationGroup(class QObject *)
func NewQParallelAnimationGroup(parent *QObject /*777 QObject **/) *QParallelAnimationGroup {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQParallelAnimationGroupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:57
// index:0
// Public virtual
// void ~QParallelAnimationGroup()
func DeleteQParallelAnimationGroup(*QParallelAnimationGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:59
// index:0
// Public virtual
// int duration()
func (this *QParallelAnimationGroup) Duration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QParallelAnimationGroup8durationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:63
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QParallelAnimationGroup) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroup5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:65
// index:0
// Protected virtual
// void updateCurrentTime(int)
func (this *QParallelAnimationGroup) UpdateCurrentTime(currentTime int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroup17updateCurrentTimeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), currentTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:66
// index:0
// Protected virtual
// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QParallelAnimationGroup) UpdateState(newState int, oldState int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:67
// index:0
// Protected virtual
// void updateDirection(class QAbstractAnimation::Direction)
func (this *QParallelAnimationGroup) UpdateDirection(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

//  body block end
