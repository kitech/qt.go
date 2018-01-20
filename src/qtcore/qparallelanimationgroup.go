//  header block begin
// /usr/include/qt/QtCore/qparallelanimationgroup.h
// #include <qparallelanimationgroup.h>
// #include <QtCore>
package qtcore

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
	*qtrt.CObject
}

func (this *QParallelAnimationGroup) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QParallelAnimationGroup) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QParallelAnimationGroup10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:56
// index:0
// void QParallelAnimationGroup(class QObject *)
func NewQParallelAnimationGroup(parent unsafe.Pointer) *QParallelAnimationGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQParallelAnimationGroupFromPointer(cthis)
	return gothis
}
func NewQParallelAnimationGroupFromPointer(cthis unsafe.Pointer) *QParallelAnimationGroup {
	return &QParallelAnimationGroup{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:62
// index:1
// void QParallelAnimationGroup(class QParallelAnimationGroupPrivate &, class QObject *)
func NewQParallelAnimationGroup_1(dd unsafe.Pointer, parent unsafe.Pointer) *QParallelAnimationGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroupC2ER30QParallelAnimationGroupPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQParallelAnimationGroupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:57
// index:0
// virtual
// void ~QParallelAnimationGroup()
func DeleteQParallelAnimationGroup(*QParallelAnimationGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:59
// index:0
// virtual
// int duration()
func (this *QParallelAnimationGroup) Duration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QParallelAnimationGroup8durationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:63
// index:0
// virtual
// bool event(class QEvent *)
func (this *QParallelAnimationGroup) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroup5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:65
// index:0
// virtual
// void updateCurrentTime(int)
func (this *QParallelAnimationGroup) UpdateCurrentTime(currentTime int) {
	// 0: (, currentTime int), (&currentTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroup17updateCurrentTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &currentTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:66
// index:0
// virtual
// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QParallelAnimationGroup) UpdateState(newState int, oldState int) {
	// 0: (, newState QAbstractAnimation::State, oldState QAbstractAnimation::State), (&newState, &oldState)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &newState, &oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:67
// index:0
// virtual
// void updateDirection(class QAbstractAnimation::Direction)
func (this *QParallelAnimationGroup) UpdateDirection(direction int) {
	// 0: (, direction QAbstractAnimation::Direction), (&direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &direction)
	gopp.ErrPrint(err, rv)
}

//  body block end
