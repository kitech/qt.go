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
// extern C begin: 16
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
type QMediaServiceProviderFactoryInterface struct {
	*qtrt.CObject
}
type QMediaServiceProviderFactoryInterface_ITF interface {
	QMediaServiceProviderFactoryInterface_PTR() *QMediaServiceProviderFactoryInterface
}

func (ptr *QMediaServiceProviderFactoryInterface) QMediaServiceProviderFactoryInterface_PTR() *QMediaServiceProviderFactoryInterface {
	return ptr
}

func (this *QMediaServiceProviderFactoryInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaServiceProviderFactoryInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaServiceProviderFactoryInterfaceFromPointer(cthis unsafe.Pointer) *QMediaServiceProviderFactoryInterface {
	return &QMediaServiceProviderFactoryInterface{&qtrt.CObject{cthis}}
}
func (*QMediaServiceProviderFactoryInterface) NewFromPointer(cthis unsafe.Pointer) *QMediaServiceProviderFactoryInterface {
	return NewQMediaServiceProviderFactoryInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:112
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QMediaService * create(const QString &)

/*
Constructs a new instance of the QMediaService identified by key.

The QMediaService returned must be destroyed with release().
*/
func (this *QMediaServiceProviderFactoryInterface) Create(key string) *QMediaService /*777 QMediaService **/ {
	var tmpArg0 = qtcore.NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN37QMediaServiceProviderFactoryInterface6createERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaServiceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:113
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void release(QMediaService *)

/*
Destroys a media service constructed with create().
*/
func (this *QMediaServiceProviderFactoryInterface) Release(service QMediaService_ITF /*777 QMediaService **/) {
	var convArg0 unsafe.Pointer
	if service != nil && service.QMediaService_PTR() != nil {
		convArg0 = service.QMediaService_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN37QMediaServiceProviderFactoryInterface7releaseEP13QMediaService", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaServiceProviderFactoryInterface()

/*

 */
func DeleteQMediaServiceProviderFactoryInterface(this *QMediaServiceProviderFactoryInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN37QMediaServiceProviderFactoryInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11875() {
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
