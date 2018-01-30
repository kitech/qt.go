package qtcore

// /usr/include/qt/QtCore/qanimationgroup.h
// #include <qanimationgroup.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QAnimationGroup struct {
	*QAbstractAnimation
}

func (this *QAnimationGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractAnimation.GetCthis()
	}
}
func (this *QAnimationGroup) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractAnimation = NewQAbstractAnimationFromPointer(cthis)
}
func NewQAnimationGroupFromPointer(cthis unsafe.Pointer) *QAnimationGroup {
	bcthis0 := NewQAbstractAnimationFromPointer(cthis)
	return &QAnimationGroup{bcthis0}
}
func (*QAnimationGroup) NewFromPointer(cthis unsafe.Pointer) *QAnimationGroup {
	return NewQAnimationGroupFromPointer(cthis)
}

// /usr/include/qt/QtCore/qanimationgroup.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAnimationGroup) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAnimationGroup10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qanimationgroup.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAnimationGroup(QObject *)
func NewQAnimationGroup(parent *QObject /*777 QObject **/) *QAnimationGroup {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAnimationGroupC1EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAnimationGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qanimationgroup.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAnimationGroup()
func DeleteQAnimationGroup(*QAnimationGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAnimationGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qanimationgroup.h:59
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractAnimation * animationAt(int)
func (this *QAnimationGroup) AnimationAt(index int) *QAbstractAnimation /*777 QAbstractAnimation **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAnimationGroup11animationAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractAnimationFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qanimationgroup.h:60
// index:0
// Public Visibility=Default Availability=Available
// [4] int animationCount()
func (this *QAnimationGroup) AnimationCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAnimationGroup14animationCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qanimationgroup.h:61
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOfAnimation(QAbstractAnimation *)
func (this *QAnimationGroup) IndexOfAnimation(animation *QAbstractAnimation /*777 QAbstractAnimation **/) int {
	var convArg0 = animation.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAnimationGroup16indexOfAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qanimationgroup.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAnimation(QAbstractAnimation *)
func (this *QAnimationGroup) AddAnimation(animation *QAbstractAnimation /*777 QAbstractAnimation **/) {
	var convArg0 = animation.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAnimationGroup12addAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qanimationgroup.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertAnimation(int, QAbstractAnimation *)
func (this *QAnimationGroup) InsertAnimation(index int, animation *QAbstractAnimation /*777 QAbstractAnimation **/) {
	var convArg1 = animation.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAnimationGroup15insertAnimationEiP18QAbstractAnimation", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qanimationgroup.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAnimation(QAbstractAnimation *)
func (this *QAnimationGroup) RemoveAnimation(animation *QAbstractAnimation /*777 QAbstractAnimation **/) {
	var convArg0 = animation.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAnimationGroup15removeAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qanimationgroup.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractAnimation * takeAnimation(int)
func (this *QAnimationGroup) TakeAnimation(index int) *QAbstractAnimation /*777 QAbstractAnimation **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAnimationGroup13takeAnimationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractAnimationFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qanimationgroup.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QAnimationGroup) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAnimationGroup5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qanimationgroup.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAnimationGroup) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAnimationGroup5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
