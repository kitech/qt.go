package qtcore

// /usr/include/qt/QtCore/qpropertyanimation.h
// #include <qpropertyanimation.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QPropertyAnimation) InheritEvent(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void updateCurrentValue(const class QVariant &)
func (this *QPropertyAnimation) InheritUpdateCurrentValue(f func(value *QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCurrentValue", f)
}

// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
func (this *QPropertyAnimation) InheritUpdateState(f func(newState int, oldState int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateState", f)
}

type QPropertyAnimation struct {
	*QVariantAnimation
}
type QPropertyAnimation_ITF interface {
	QVariantAnimation_ITF
	QPropertyAnimation_PTR() *QPropertyAnimation
}

func (ptr *QPropertyAnimation) QPropertyAnimation_PTR() *QPropertyAnimation { return ptr }

func (this *QPropertyAnimation) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QVariantAnimation.GetCthis()
	}
}
func (this *QPropertyAnimation) SetCthis(cthis unsafe.Pointer) {
	this.QVariantAnimation = NewQVariantAnimationFromPointer(cthis)
}
func NewQPropertyAnimationFromPointer(cthis unsafe.Pointer) *QPropertyAnimation {
	bcthis0 := NewQVariantAnimationFromPointer(cthis)
	return &QPropertyAnimation{bcthis0}
}
func (*QPropertyAnimation) NewFromPointer(cthis unsafe.Pointer) *QPropertyAnimation {
	return NewQPropertyAnimationFromPointer(cthis)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QPropertyAnimation) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QPropertyAnimation10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qpropertyanimation.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPropertyAnimation(QObject *)
func NewQPropertyAnimation(parent QObject_ITF /*777 QObject **/) *QPropertyAnimation {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPropertyAnimationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPropertyAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPropertyAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qpropertyanimation.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPropertyAnimation(QObject *)
func NewQPropertyAnimation__() *QPropertyAnimation {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPropertyAnimationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPropertyAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPropertyAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qpropertyanimation.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPropertyAnimation(QObject *, const QByteArray &, QObject *)
func NewQPropertyAnimation_1(target QObject_ITF /*777 QObject **/, propertyName QByteArray_ITF, parent QObject_ITF /*777 QObject **/) *QPropertyAnimation {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if propertyName != nil && propertyName.QByteArray_PTR() != nil {
		convArg1 = propertyName.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPropertyAnimationC2EP7QObjectRK10QByteArrayS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPropertyAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPropertyAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qpropertyanimation.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPropertyAnimation(QObject *, const QByteArray &, QObject *)
func NewQPropertyAnimation_1_(target QObject_ITF /*777 QObject **/, propertyName QByteArray_ITF) *QPropertyAnimation {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if propertyName != nil && propertyName.QByteArray_PTR() != nil {
		convArg1 = propertyName.QByteArray_PTR().GetCthis()
	}
	// arg: 2, QObject *=Pointer, QObject=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPropertyAnimationC2EP7QObjectRK10QByteArrayS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPropertyAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPropertyAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qpropertyanimation.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPropertyAnimation()
func DeleteQPropertyAnimation(this *QPropertyAnimation) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPropertyAnimationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * targetObject() const
func (this *QPropertyAnimation) TargetObject() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QPropertyAnimation12targetObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qpropertyanimation.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTargetObject(QObject *)
func (this *QPropertyAnimation) SetTargetObject(target QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if target != nil && target.QObject_PTR() != nil {
		convArg0 = target.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPropertyAnimation15setTargetObjectEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray propertyName() const
func (this *QPropertyAnimation) PropertyName() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QPropertyAnimation12propertyNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qpropertyanimation.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPropertyName(const QByteArray &)
func (this *QPropertyAnimation) SetPropertyName(propertyName QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if propertyName != nil && propertyName.QByteArray_PTR() != nil {
		convArg0 = propertyName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPropertyAnimation15setPropertyNameERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QPropertyAnimation) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPropertyAnimation5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpropertyanimation.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateCurrentValue(const QVariant &)
func (this *QPropertyAnimation) UpdateCurrentValue(value QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPropertyAnimation18updateCurrentValueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:71
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateState(QAbstractAnimation::State, QAbstractAnimation::State)
func (this *QPropertyAnimation) UpdateState(newState int, oldState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QPropertyAnimation11updateStateEN18QAbstractAnimation5StateES1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
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
