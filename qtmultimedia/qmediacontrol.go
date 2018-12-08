package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediacontrol.h
// #include <qmediacontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QMediaControl struct {
	*qtcore.QObject
}
type QMediaControl_ITF interface {
	qtcore.QObject_ITF
	QMediaControl_PTR() *QMediaControl
}

func (ptr *QMediaControl) QMediaControl_PTR() *QMediaControl { return ptr }

func (this *QMediaControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMediaControl) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQMediaControlFromPointer(cthis unsafe.Pointer) *QMediaControl {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QMediaControl{bcthis0}
}
func (*QMediaControl) NewFromPointer(cthis unsafe.Pointer) *QMediaControl {
	return NewQMediaControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediacontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMediaControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediacontrol.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaControl()

/*

 */
func DeleteQMediaControl(this *QMediaControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediacontrol.h:62
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaControl(QObject *)

/*
Constructs a media control with the given parent.
*/
func (*QMediaControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaControl {
	return NewQMediaControl(parent)
}
func NewQMediaControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediacontrol.h:62
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaControl(QObject *)

/*
Constructs a media control with the given parent.
*/
func (*QMediaControl) NewForInheritp() *QMediaControl {
	return NewQMediaControlp()
}
func NewQMediaControlp() *QMediaControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaControl")
	return gothis
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
