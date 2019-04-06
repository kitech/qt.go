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
type QMediaServiceFeaturesInterface struct {
	*qtrt.CObject
}
type QMediaServiceFeaturesInterface_ITF interface {
	QMediaServiceFeaturesInterface_PTR() *QMediaServiceFeaturesInterface
}

func (ptr *QMediaServiceFeaturesInterface) QMediaServiceFeaturesInterface_PTR() *QMediaServiceFeaturesInterface {
	return ptr
}

func (this *QMediaServiceFeaturesInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaServiceFeaturesInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaServiceFeaturesInterfaceFromPointer(cthis unsafe.Pointer) *QMediaServiceFeaturesInterface {
	return &QMediaServiceFeaturesInterface{&qtrt.CObject{cthis}}
}
func (*QMediaServiceFeaturesInterface) NewFromPointer(cthis unsafe.Pointer) *QMediaServiceFeaturesInterface {
	return NewQMediaServiceFeaturesInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:180
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QMediaServiceFeaturesInterface()

/*

 */
func DeleteQMediaServiceFeaturesInterface(this *QMediaServiceFeaturesInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN30QMediaServiceFeaturesInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:181
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QMediaServiceProviderHint::Features supportedFeatures(const QByteArray &) const

/*

 */
func (this *QMediaServiceFeaturesInterface) SupportedFeatures(service qtcore.QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if service != nil && service.QByteArray_PTR() != nil {
		convArg0 = service.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK30QMediaServiceFeaturesInterface17supportedFeaturesERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

//  body block end

//  keep block begin

func init_unused_11885() {
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
