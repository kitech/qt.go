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
// extern C begin: 2
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
type QMediaServiceCameraInfoInterface struct {
	*qtrt.CObject
}
type QMediaServiceCameraInfoInterface_ITF interface {
	QMediaServiceCameraInfoInterface_PTR() *QMediaServiceCameraInfoInterface
}

func (ptr *QMediaServiceCameraInfoInterface) QMediaServiceCameraInfoInterface_PTR() *QMediaServiceCameraInfoInterface {
	return ptr
}

func (this *QMediaServiceCameraInfoInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaServiceCameraInfoInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaServiceCameraInfoInterfaceFromPointer(cthis unsafe.Pointer) *QMediaServiceCameraInfoInterface {
	return &QMediaServiceCameraInfoInterface{&qtrt.CObject{cthis}}
}
func (*QMediaServiceCameraInfoInterface) NewFromPointer(cthis unsafe.Pointer) *QMediaServiceCameraInfoInterface {
	return NewQMediaServiceCameraInfoInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:166
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QMediaServiceCameraInfoInterface()

/*

 */
func DeleteQMediaServiceCameraInfoInterface(this *QMediaServiceCameraInfoInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QMediaServiceCameraInfoInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:167
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCamera::Position cameraPosition(const QByteArray &) const

/*

 */
func (this *QMediaServiceCameraInfoInterface) CameraPosition(device qtcore.QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg0 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QMediaServiceCameraInfoInterface14cameraPositionERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:168
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int cameraOrientation(const QByteArray &) const

/*

 */
func (this *QMediaServiceCameraInfoInterface) CameraOrientation(device qtcore.QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg0 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QMediaServiceCameraInfoInterface17cameraOrientationERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end

//  keep block begin

func init_unused_11883() {
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
