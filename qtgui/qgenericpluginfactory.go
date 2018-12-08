package qtgui

// /usr/include/qt/QtGui/qgenericpluginfactory.h
// #include <qgenericpluginfactory.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QGenericPluginFactory struct {
	*qtrt.CObject
}
type QGenericPluginFactory_ITF interface {
	QGenericPluginFactory_PTR() *QGenericPluginFactory
}

func (ptr *QGenericPluginFactory) QGenericPluginFactory_PTR() *QGenericPluginFactory { return ptr }

func (this *QGenericPluginFactory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGenericPluginFactory) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGenericPluginFactoryFromPointer(cthis unsafe.Pointer) *QGenericPluginFactory {
	return &QGenericPluginFactory{&qtrt.CObject{cthis}}
}
func (*QGenericPluginFactory) NewFromPointer(cthis unsafe.Pointer) *QGenericPluginFactory {
	return NewQGenericPluginFactoryFromPointer(cthis)
}

// /usr/include/qt/QtGui/qgenericpluginfactory.h:55
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList keys()

/*
Returns the list of valid keys, i.e. the available mouse drivers.

See also create().
*/
func (this *QGenericPluginFactory) Keys() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGenericPluginFactory4keysEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QGenericPluginFactory_Keys() *qtcore.QStringList /*123*/ {
	var nilthis *QGenericPluginFactory
	rv := nilthis.Keys()
	return rv
}

// /usr/include/qt/QtGui/qgenericpluginfactory.h:56
// index:0
// Public static Visibility=Default Availability=Available
// [8] QObject * create(const QString &, const QString &)

/*
Creates the driver specified by key, using the given specification.

Note that the keys are case-insensitive.

See also keys().
*/
func (this *QGenericPluginFactory) Create(arg0 string, arg1 string) *qtcore.QObject /*777 QObject **/ {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGenericPluginFactory6createERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGenericPluginFactory_Create(arg0 string, arg1 string) *qtcore.QObject /*777 QObject **/ {
	var nilthis *QGenericPluginFactory
	rv := nilthis.Create(arg0, arg1)
	return rv
}

func DeleteQGenericPluginFactory(this *QGenericPluginFactory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGenericPluginFactoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
}

//  keep block end
