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
type QMediaServiceProviderPlugin struct {
	*qtcore.QObject
	*QMediaServiceProviderFactoryInterface
}
type QMediaServiceProviderPlugin_ITF interface {
	qtcore.QObject_ITF
	QMediaServiceProviderFactoryInterface_ITF
	QMediaServiceProviderPlugin_PTR() *QMediaServiceProviderPlugin
}

func (ptr *QMediaServiceProviderPlugin) QMediaServiceProviderPlugin_PTR() *QMediaServiceProviderPlugin {
	return ptr
}

func (this *QMediaServiceProviderPlugin) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMediaServiceProviderPlugin) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QMediaServiceProviderFactoryInterface = NewQMediaServiceProviderFactoryInterfaceFromPointer(cthis)
}
func NewQMediaServiceProviderPluginFromPointer(cthis unsafe.Pointer) *QMediaServiceProviderPlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQMediaServiceProviderFactoryInterfaceFromPointer(cthis)
	return &QMediaServiceProviderPlugin{bcthis0, bcthis1}
}
func (*QMediaServiceProviderPlugin) NewFromPointer(cthis unsafe.Pointer) *QMediaServiceProviderPlugin {
	return NewQMediaServiceProviderPluginFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:194
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaServiceProviderPlugin) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QMediaServiceProviderPlugin10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:198
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QMediaService * create(const QString &)

/*
Constructs a new instance of the QMediaService identified by key.

The QMediaService returned must be destroyed with release().
*/
func (this *QMediaServiceProviderPlugin) Create(key string) *QMediaService /*777 QMediaService **/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN27QMediaServiceProviderPlugin6createERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaServiceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:199
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void release(QMediaService *)

/*
Destroys a media service constructed with create().
*/
func (this *QMediaServiceProviderPlugin) Release(service QMediaService_ITF /*777 QMediaService **/) {
	var convArg0 unsafe.Pointer
	if service != nil && service.QMediaService_PTR() != nil {
		convArg0 = service.QMediaService_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QMediaServiceProviderPlugin7releaseEP13QMediaService", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQMediaServiceProviderPlugin(this *QMediaServiceProviderPlugin) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QMediaServiceProviderPluginD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
