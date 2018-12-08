package qtmultimedia

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h
// #include <qvideodeviceselectorcontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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
type QVideoDeviceSelectorControl struct {
	*QMediaControl
}
type QVideoDeviceSelectorControl_ITF interface {
	QMediaControl_ITF
	QVideoDeviceSelectorControl_PTR() *QVideoDeviceSelectorControl
}

func (ptr *QVideoDeviceSelectorControl) QVideoDeviceSelectorControl_PTR() *QVideoDeviceSelectorControl {
	return ptr
}

func (this *QVideoDeviceSelectorControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QVideoDeviceSelectorControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQVideoDeviceSelectorControlFromPointer(cthis unsafe.Pointer) *QVideoDeviceSelectorControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QVideoDeviceSelectorControl{bcthis0}
}
func (*QVideoDeviceSelectorControl) NewFromPointer(cthis unsafe.Pointer) *QVideoDeviceSelectorControl {
	return NewQVideoDeviceSelectorControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QVideoDeviceSelectorControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QVideoDeviceSelectorControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QVideoDeviceSelectorControl()

/*

 */
func DeleteQVideoDeviceSelectorControl(this *QVideoDeviceSelectorControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QVideoDeviceSelectorControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:57
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int deviceCount() const

/*
Returns the number of available video devices;
*/
func (this *QVideoDeviceSelectorControl) DeviceCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QVideoDeviceSelectorControl11deviceCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString deviceName(int) const

/*
Returns the name of the video device at index.
*/
func (this *QVideoDeviceSelectorControl) DeviceName(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QVideoDeviceSelectorControl10deviceNameEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString deviceDescription(int) const

/*
Returns a description of the video device at index.
*/
func (this *QVideoDeviceSelectorControl) DeviceDescription(index int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QVideoDeviceSelectorControl17deviceDescriptionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int defaultDevice() const

/*
Returns the index of the default video device.
*/
func (this *QVideoDeviceSelectorControl) DefaultDevice() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QVideoDeviceSelectorControl13defaultDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int selectedDevice() const

/*
Returns the index of the selected video device.

See also setSelectedDevice().
*/
func (this *QVideoDeviceSelectorControl) SelectedDevice() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QVideoDeviceSelectorControl14selectedDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setSelectedDevice(int)

/*
Sets the selected video device index.

See also selectedDevice().
*/
func (this *QVideoDeviceSelectorControl) SetSelectedDevice(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QVideoDeviceSelectorControl17setSelectedDeviceEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectedDeviceChanged(int)

/*
Signals that the selected video device index has changed.

Note: Signal selectedDeviceChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(videoDeviceSelectorControl, QOverload<int>::of(&QVideoDeviceSelectorControl::selectedDeviceChanged),
      [=](int index){ /-* ... *-/ });
*/
func (this *QVideoDeviceSelectorControl) SelectedDeviceChanged(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QVideoDeviceSelectorControl21selectedDeviceChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void selectedDeviceChanged(const QString &)

/*
Signals that the selected video device index has changed.

Note: Signal selectedDeviceChanged is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(videoDeviceSelectorControl, QOverload<int>::of(&QVideoDeviceSelectorControl::selectedDeviceChanged),
      [=](int index){ /-* ... *-/ });
*/
func (this *QVideoDeviceSelectorControl) SelectedDeviceChanged1(deviceName string) {
	var tmpArg0 = qtcore.NewQString5(deviceName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN27QVideoDeviceSelectorControl21selectedDeviceChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void devicesChanged()

/*
Signals that the list of available video devices has changed.
*/
func (this *QVideoDeviceSelectorControl) DevicesChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QVideoDeviceSelectorControl14devicesChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:74
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QVideoDeviceSelectorControl(QObject *)

/*
Constructs a video device selector control with the given parent.
*/
func (*QVideoDeviceSelectorControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QVideoDeviceSelectorControl {
	return NewQVideoDeviceSelectorControl(parent)
}
func NewQVideoDeviceSelectorControl(parent qtcore.QObject_ITF /*777 QObject **/) *QVideoDeviceSelectorControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QVideoDeviceSelectorControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoDeviceSelectorControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVideoDeviceSelectorControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideodeviceselectorcontrol.h:74
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QVideoDeviceSelectorControl(QObject *)

/*
Constructs a video device selector control with the given parent.
*/
func (*QVideoDeviceSelectorControl) NewForInheritp() *QVideoDeviceSelectorControl {
	return NewQVideoDeviceSelectorControlp()
}
func NewQVideoDeviceSelectorControlp() *QVideoDeviceSelectorControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN27QVideoDeviceSelectorControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoDeviceSelectorControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVideoDeviceSelectorControl")
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
