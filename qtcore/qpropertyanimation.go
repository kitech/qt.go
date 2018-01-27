package qtcore

// /usr/include/qt/QtCore/qpropertyanimation.h
// #include <qpropertyanimation.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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
type QPropertyAnimation struct {
	*QVariantAnimation
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QPropertyAnimation) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPropertyAnimation10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qpropertyanimation.h:58
// index:0
// Public
// void QPropertyAnimation(QObject *)
func NewQPropertyAnimation(parent *QObject /*777 QObject **/) *QPropertyAnimation {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPropertyAnimationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpropertyanimation.h:59
// index:1
// Public
// void QPropertyAnimation(QObject *, const QByteArray &, QObject *)
func NewQPropertyAnimation_1(target *QObject /*777 QObject **/, propertyName *QByteArray, parent *QObject /*777 QObject **/) *QPropertyAnimation {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = target.GetCthis()
	var convArg1 = propertyName.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimationC2EP7QObjectRK10QByteArrayS1_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQPropertyAnimationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpropertyanimation.h:60
// index:0
// Public virtual
// void ~QPropertyAnimation()
func DeleteQPropertyAnimation(*QPropertyAnimation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:62
// index:0
// Public
// QObject * targetObject()
func (this *QPropertyAnimation) TargetObject() *QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPropertyAnimation12targetObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qpropertyanimation.h:63
// index:0
// Public
// void setTargetObject(QObject *)
func (this *QPropertyAnimation) SetTargetObject(target *QObject /*777 QObject **/) {
	var convArg0 = target.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimation15setTargetObjectEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:65
// index:0
// Public
// QByteArray propertyName()
func (this *QPropertyAnimation) PropertyName() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPropertyAnimation12propertyNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qpropertyanimation.h:66
// index:0
// Public
// void setPropertyName(const QByteArray &)
func (this *QPropertyAnimation) SetPropertyName(propertyName *QByteArray) {
	var convArg0 = propertyName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimation15setPropertyNameERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:69
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QPropertyAnimation) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimation5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qpropertyanimation.h:70
// index:0
// Protected virtual
// void updateCurrentValue(const QVariant &)
func (this *QPropertyAnimation) UpdateCurrentValue(value *QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimation18updateCurrentValueERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:71
// index:0
// Protected virtual
// void updateState(QAbstractAnimation::State, QAbstractAnimation::State)
func (this *QPropertyAnimation) UpdateState(newState int, oldState int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimation11updateStateEN18QAbstractAnimation5StateES1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newState, oldState)
	gopp.ErrPrint(err, rv)
}

//  body block end
