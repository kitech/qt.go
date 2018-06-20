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
type QMediaServiceDefaultDeviceInterface struct {
	*qtrt.CObject
}
type QMediaServiceDefaultDeviceInterface_ITF interface {
	QMediaServiceDefaultDeviceInterface_PTR() *QMediaServiceDefaultDeviceInterface
}

func (ptr *QMediaServiceDefaultDeviceInterface) QMediaServiceDefaultDeviceInterface_PTR() *QMediaServiceDefaultDeviceInterface {
	return ptr
}

func (this *QMediaServiceDefaultDeviceInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaServiceDefaultDeviceInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaServiceDefaultDeviceInterfaceFromPointer(cthis unsafe.Pointer) *QMediaServiceDefaultDeviceInterface {
	return &QMediaServiceDefaultDeviceInterface{&qtrt.CObject{cthis}}
}
func (*QMediaServiceDefaultDeviceInterface) NewFromPointer(cthis unsafe.Pointer) *QMediaServiceDefaultDeviceInterface {
	return NewQMediaServiceDefaultDeviceInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:156
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QMediaServiceDefaultDeviceInterface()

/*

 */
func DeleteQMediaServiceDefaultDeviceInterface(this *QMediaServiceDefaultDeviceInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN35QMediaServiceDefaultDeviceInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:157
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QByteArray defaultDevice(const QByteArray &) const

/*

 */
func (this *QMediaServiceDefaultDeviceInterface) DefaultDevice(service qtcore.QByteArray_ITF) *qtcore.QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if service != nil && service.QByteArray_PTR() != nil {
		convArg0 = service.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK35QMediaServiceDefaultDeviceInterface13defaultDeviceERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
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
