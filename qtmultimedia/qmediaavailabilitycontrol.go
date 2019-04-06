package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaavailabilitycontrol.h
// #include <qmediaavailabilitycontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QMediaAvailabilityControl struct {
	*QMediaControl
}
type QMediaAvailabilityControl_ITF interface {
	QMediaControl_ITF
	QMediaAvailabilityControl_PTR() *QMediaAvailabilityControl
}

func (ptr *QMediaAvailabilityControl) QMediaAvailabilityControl_PTR() *QMediaAvailabilityControl {
	return ptr
}

func (this *QMediaAvailabilityControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMediaAvailabilityControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMediaAvailabilityControlFromPointer(cthis unsafe.Pointer) *QMediaAvailabilityControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMediaAvailabilityControl{bcthis0}
}
func (*QMediaAvailabilityControl) NewFromPointer(cthis unsafe.Pointer) *QMediaAvailabilityControl {
	return NewQMediaAvailabilityControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaavailabilitycontrol.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaAvailabilityControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaAvailabilityControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaavailabilitycontrol.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaAvailabilityControl()

/*

 */
func DeleteQMediaAvailabilityControl(this *QMediaAvailabilityControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaAvailabilityControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaavailabilitycontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QMultimedia::AvailabilityStatus availability() const

/*
Returns the current availability of this instance of the media service. If the availability changes at run time (for example, some other media client takes all media resources) the availabilityChanges() signal should be emitted.
*/
func (this *QMediaAvailabilityControl) Availability() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaAvailabilityControl12availabilityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaavailabilitycontrol.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void availabilityChanged(QMultimedia::AvailabilityStatus)

/*
Signal emitted when the current availability value changed.
*/
func (this *QMediaAvailabilityControl) AvailabilityChanged(availability int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaAvailabilityControl19availabilityChangedEN11QMultimedia18AvailabilityStatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), availability)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaavailabilitycontrol.h:65
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaAvailabilityControl(QObject *)

/*
Constructs an availability control object with parent.
*/
func (*QMediaAvailabilityControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaAvailabilityControl {
	return NewQMediaAvailabilityControl(parent)
}
func NewQMediaAvailabilityControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaAvailabilityControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaAvailabilityControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaAvailabilityControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaAvailabilityControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaavailabilitycontrol.h:65
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaAvailabilityControl(QObject *)

/*
Constructs an availability control object with parent.
*/
func (*QMediaAvailabilityControl) NewForInheritp() *QMediaAvailabilityControl {
	return NewQMediaAvailabilityControlp()
}
func NewQMediaAvailabilityControlp() *QMediaAvailabilityControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaAvailabilityControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaAvailabilityControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaAvailabilityControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11849() {
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
