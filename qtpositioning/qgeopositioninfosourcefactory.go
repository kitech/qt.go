package qtpositioning

// /usr/include/qt/QtPositioning/qgeopositioninfosourcefactory.h
// #include <qgeopositioninfosourcefactory.h>
// #include <QtPositioning>

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

//  ext block end

//  body block begin

/*

 */
type QGeoPositionInfoSourceFactory struct {
	*qtrt.CObject
}
type QGeoPositionInfoSourceFactory_ITF interface {
	QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory
}

func (ptr *QGeoPositionInfoSourceFactory) QGeoPositionInfoSourceFactory_PTR() *QGeoPositionInfoSourceFactory {
	return ptr
}

func (this *QGeoPositionInfoSourceFactory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGeoPositionInfoSourceFactory) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGeoPositionInfoSourceFactoryFromPointer(cthis unsafe.Pointer) *QGeoPositionInfoSourceFactory {
	return &QGeoPositionInfoSourceFactory{&qtrt.CObject{cthis}}
}
func (*QGeoPositionInfoSourceFactory) NewFromPointer(cthis unsafe.Pointer) *QGeoPositionInfoSourceFactory {
	return NewQGeoPositionInfoSourceFactoryFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosourcefactory.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGeoPositionInfoSourceFactory()

/*

 */
func DeleteQGeoPositionInfoSourceFactory(this *QGeoPositionInfoSourceFactory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QGeoPositionInfoSourceFactoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosourcefactory.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QGeoPositionInfoSource * positionInfoSource(QObject *)

/*
Returns a new QGeoPositionInfoSource associated with this plugin with parent parent. Can also return 0, in which case the plugin loader will use the factory with the next highest priority.
*/
func (this *QGeoPositionInfoSourceFactory) PositionInfoSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoPositionInfoSource /*777 QGeoPositionInfoSource **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QGeoPositionInfoSourceFactory18positionInfoSourceEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGeoPositionInfoSourceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtPositioning/qgeopositioninfosourcefactory.h:56
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QGeoSatelliteInfoSource * satelliteInfoSource(QObject *)

/*
Returns a new QGeoSatelliteInfoSource associated with this plugin with parent parent. Can also return 0, in which case the plugin loader will use the factory with the next highest priority.
*/
func (this *QGeoPositionInfoSourceFactory) SatelliteInfoSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoSatelliteInfoSource /*777 QGeoSatelliteInfoSource **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QGeoPositionInfoSourceFactory19satelliteInfoSourceEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGeoSatelliteInfoSourceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtPositioning/qgeopositioninfosourcefactory.h:57
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QGeoAreaMonitorSource * areaMonitor(QObject *)

/*
Returns a new QGeoAreaMonitorSource associated with this plugin with parent parent. Can also return 0, in which case the plugin loader will use the factory with the next highest priority.
*/
func (this *QGeoPositionInfoSourceFactory) AreaMonitor(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoAreaMonitorSource /*777 QGeoAreaMonitorSource **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QGeoPositionInfoSourceFactory11areaMonitorEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGeoAreaMonitorSourceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_11659() {
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
