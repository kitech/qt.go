package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h
// #include <qmediaserviceproviderplugin.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QMediaServiceSupportedDevicesInterface struct {
	*qtrt.CObject
}
type QMediaServiceSupportedDevicesInterface_ITF interface {
	QMediaServiceSupportedDevicesInterface_PTR() *QMediaServiceSupportedDevicesInterface
}

func (ptr *QMediaServiceSupportedDevicesInterface) QMediaServiceSupportedDevicesInterface_PTR() *QMediaServiceSupportedDevicesInterface {
	return ptr
}

func (this *QMediaServiceSupportedDevicesInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaServiceSupportedDevicesInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaServiceSupportedDevicesInterfaceFromPointer(cthis unsafe.Pointer) *QMediaServiceSupportedDevicesInterface {
	return &QMediaServiceSupportedDevicesInterface{&qtrt.CObject{cthis}}
}
func (*QMediaServiceSupportedDevicesInterface) NewFromPointer(cthis unsafe.Pointer) *QMediaServiceSupportedDevicesInterface {
	return NewQMediaServiceSupportedDevicesInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:140
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QMediaServiceSupportedDevicesInterface()

/*

 */
func DeleteQMediaServiceSupportedDevicesInterface(this *QMediaServiceSupportedDevicesInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN38QMediaServiceSupportedDevicesInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:142
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString deviceDescription(const QByteArray &, const QByteArray &)

/*

 */
func (this *QMediaServiceSupportedDevicesInterface) DeviceDescription(service qtcore.QByteArray_ITF, device qtcore.QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if service != nil && service.QByteArray_PTR() != nil {
		convArg0 = service.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg1 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN38QMediaServiceSupportedDevicesInterface17deviceDescriptionERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

//  body block end

//  keep block begin

func init_unused_11879() {
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
