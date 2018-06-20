package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h
// #include <qcameraviewfindersettingscontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QCameraViewfinderSettingsControl2 struct {
	*QMediaControl
}
type QCameraViewfinderSettingsControl2_ITF interface {
	QMediaControl_ITF
	QCameraViewfinderSettingsControl2_PTR() *QCameraViewfinderSettingsControl2
}

func (ptr *QCameraViewfinderSettingsControl2) QCameraViewfinderSettingsControl2_PTR() *QCameraViewfinderSettingsControl2 {
	return ptr
}

func (this *QCameraViewfinderSettingsControl2) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraViewfinderSettingsControl2) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraViewfinderSettingsControl2FromPointer(cthis unsafe.Pointer) *QCameraViewfinderSettingsControl2 {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraViewfinderSettingsControl2{bcthis0}
}
func (*QCameraViewfinderSettingsControl2) NewFromPointer(cthis unsafe.Pointer) *QCameraViewfinderSettingsControl2 {
	return NewQCameraViewfinderSettingsControl2FromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraViewfinderSettingsControl2) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK33QCameraViewfinderSettingsControl210metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraViewfinderSettingsControl2()

/*

 */
func DeleteQCameraViewfinderSettingsControl2(this *QCameraViewfinderSettingsControl2) {
	rv, err := qtrt.InvokeQtFunc6("_ZN33QCameraViewfinderSettingsControl2D2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:91
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QCameraViewfinderSettings viewfinderSettings() const

/*

 */
func (this *QCameraViewfinderSettingsControl2) ViewfinderSettings() *QCameraViewfinderSettings /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK33QCameraViewfinderSettingsControl218viewfinderSettingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraViewfinderSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCameraViewfinderSettings)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:92
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setViewfinderSettings(const QCameraViewfinderSettings &)

/*

 */
func (this *QCameraViewfinderSettingsControl2) SetViewfinderSettings(settings QCameraViewfinderSettings_ITF) {
	var convArg0 unsafe.Pointer
	if settings != nil && settings.QCameraViewfinderSettings_PTR() != nil {
		convArg0 = settings.QCameraViewfinderSettings_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN33QCameraViewfinderSettingsControl221setViewfinderSettingsERK25QCameraViewfinderSettings", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:95
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraViewfinderSettingsControl2(QObject *)

/*

 */
func NewQCameraViewfinderSettingsControl2(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraViewfinderSettingsControl2 {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN33QCameraViewfinderSettingsControl2C2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraViewfinderSettingsControl2FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraViewfinderSettingsControl2")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcameraviewfindersettingscontrol.h:95
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraViewfinderSettingsControl2(QObject *)

/*

 */
func NewQCameraViewfinderSettingsControl2__() *QCameraViewfinderSettingsControl2 {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN33QCameraViewfinderSettingsControl2C2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraViewfinderSettingsControl2FromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraViewfinderSettingsControl2")
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
