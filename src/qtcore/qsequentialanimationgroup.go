//  header block begin
// /usr/include/qt/QtCore/qsequentialanimationgroup.h
// #include <qsequentialanimationgroup.h>
// #include <QtCore>
package qtcore

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
type QSequentialAnimationGroup struct {
	*qtrt.CObject
}

func (this *QSequentialAnimationGroup) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSequentialAnimationGroup) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:59
// index:0
// void QSequentialAnimationGroup(class QObject *)
func NewQSequentialAnimationGroup(parent unsafe.Pointer) *QSequentialAnimationGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSequentialAnimationGroupFromPointer(cthis)
	return gothis
}
func NewQSequentialAnimationGroupFromPointer(cthis unsafe.Pointer) *QSequentialAnimationGroup {
	return &QSequentialAnimationGroup{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:72
// index:1
// void QSequentialAnimationGroup(class QSequentialAnimationGroupPrivate &, class QObject *)
func NewQSequentialAnimationGroup_1(dd unsafe.Pointer, parent unsafe.Pointer) *QSequentialAnimationGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroupC2ER32QSequentialAnimationGroupPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSequentialAnimationGroupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:60
// index:0
// virtual
// void ~QSequentialAnimationGroup()
func DeleteQSequentialAnimationGroup(*QSequentialAnimationGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:62
// index:0
// QPauseAnimation * addPause(int)
func (this *QSequentialAnimationGroup) AddPause(msecs int) {
	// 0: (, msecs int), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup8addPauseEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:63
// index:0
// QPauseAnimation * insertPause(int, int)
func (this *QSequentialAnimationGroup) InsertPause(index int, msecs int) {
	// 0: (, index int, msecs int), (&index, &msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup11insertPauseEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:65
// index:0
// QAbstractAnimation * currentAnimation()
func (this *QSequentialAnimationGroup) CurrentAnimation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup16currentAnimationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:66
// index:0
// virtual
// int duration()
func (this *QSequentialAnimationGroup) Duration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup8durationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:69
// index:0
// void currentAnimationChanged(class QAbstractAnimation *)
func (this *QSequentialAnimationGroup) CurrentAnimationChanged(current unsafe.Pointer) {
	// 0: (, current QAbstractAnimation *), (current)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup23currentAnimationChangedEP18QAbstractAnimation", ffiqt.FFI_TYPE_VOID, this.GetCthis(), current)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:73
// index:0
// virtual
// bool event(class QEvent *)
func (this *QSequentialAnimationGroup) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:75
// index:0
// virtual
// void updateCurrentTime(int)
func (this *QSequentialAnimationGroup) UpdateCurrentTime(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup17updateCurrentTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:76
// index:0
// virtual
// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QSequentialAnimationGroup) UpdateState(newState int, oldState int) {
	// 0: (, newState QAbstractAnimation::State, oldState QAbstractAnimation::State), (&newState, &oldState)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &newState, &oldState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:77
// index:0
// virtual
// void updateDirection(class QAbstractAnimation::Direction)
func (this *QSequentialAnimationGroup) UpdateDirection(direction int) {
	// 0: (, direction QAbstractAnimation::Direction), (&direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &direction)
	gopp.ErrPrint(err, rv)
}

//  body block end
