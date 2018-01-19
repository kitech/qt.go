//  header block begin
// /usr/include/qt/QtCore/qpropertyanimation.h
// #include <qpropertyanimation.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QPropertyAnimation struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qpropertyanimation.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPropertyAnimation) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPropertyAnimation10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:58
// index:0
// void QPropertyAnimation(class QObject *)
func NewQPropertyAnimation(parent unsafe.Pointer) *QPropertyAnimation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QPropertyAnimation{cthis}
}

// /usr/include/qt/QtCore/qpropertyanimation.h:59
// index:1
// void QPropertyAnimation(class QObject *, const class QByteArray &, class QObject *)
func NewQPropertyAnimation_1(target unsafe.Pointer, propertyName unsafe.Pointer, parent unsafe.Pointer) *QPropertyAnimation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimationC2EP7QObjectRK10QByteArrayS1_", ffiqt.FFI_TYPE_VOID, cthis, target, propertyName, parent)
	gopp.ErrPrint(err, rv)
	return &QPropertyAnimation{cthis}
}

// /usr/include/qt/QtCore/qpropertyanimation.h:60
// index:0
// virtual
// void ~QPropertyAnimation()
func DeleteQPropertyAnimation(*QPropertyAnimation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:62
// index:0
// QObject * targetObject()
func (this *QPropertyAnimation) TargetObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPropertyAnimation12targetObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:63
// index:0
// void setTargetObject(class QObject *)
func (this *QPropertyAnimation) SetTargetObject(target unsafe.Pointer) {
	// 0: (, QObject * target), (target)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimation15setTargetObjectEP7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, target)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:65
// index:0
// QByteArray propertyName()
func (this *QPropertyAnimation) PropertyName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QPropertyAnimation12propertyNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpropertyanimation.h:66
// index:0
// void setPropertyName(const class QByteArray &)
func (this *QPropertyAnimation) SetPropertyName(propertyName unsafe.Pointer) {
	// 0: (, const QByteArray & propertyName), (propertyName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QPropertyAnimation15setPropertyNameERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, propertyName)
	gopp.ErrPrint(err, rv)
}

//  body block end
