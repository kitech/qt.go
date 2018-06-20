package qtmultimedia

// /usr/include/qt/QtMultimedia/qmedianetworkaccesscontrol.h
// #include <qmedianetworkaccesscontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QMediaNetworkAccessControl struct {
	*QMediaControl
}
type QMediaNetworkAccessControl_ITF interface {
	QMediaControl_ITF
	QMediaNetworkAccessControl_PTR() *QMediaNetworkAccessControl
}

func (ptr *QMediaNetworkAccessControl) QMediaNetworkAccessControl_PTR() *QMediaNetworkAccessControl {
	return ptr
}

func (this *QMediaNetworkAccessControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QMediaNetworkAccessControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQMediaNetworkAccessControlFromPointer(cthis unsafe.Pointer) *QMediaNetworkAccessControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QMediaNetworkAccessControl{bcthis0}
}
func (*QMediaNetworkAccessControl) NewFromPointer(cthis unsafe.Pointer) *QMediaNetworkAccessControl {
	return NewQMediaNetworkAccessControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmedianetworkaccesscontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaNetworkAccessControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QMediaNetworkAccessControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmedianetworkaccesscontrol.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaNetworkAccessControl()

/*

 */
func DeleteQMediaNetworkAccessControl(this *QMediaNetworkAccessControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMediaNetworkAccessControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmedianetworkaccesscontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QNetworkConfiguration currentConfiguration() const

/*
Returns the current active configuration in use. A default constructed QNetworkConfigration is returned if no user supplied configuration are in use.
*/
func (this *QMediaNetworkAccessControl) CurrentConfiguration() *qtnetwork.QNetworkConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QMediaNetworkAccessControl20currentConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtnetwork.NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtnetwork.DeleteQNetworkConfiguration)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmedianetworkaccesscontrol.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void configurationChanged(const QNetworkConfiguration &)

/*
This signal is emitted when the current active network configuration changes to configuration.
*/
func (this *QMediaNetworkAccessControl) ConfigurationChanged(configuration qtnetwork.QNetworkConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if configuration != nil && configuration.QNetworkConfiguration_PTR() != nil {
		convArg0 = configuration.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMediaNetworkAccessControl20configurationChangedERK21QNetworkConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmedianetworkaccesscontrol.h:68
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaNetworkAccessControl(QObject *)

/*

 */
func NewQMediaNetworkAccessControl(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaNetworkAccessControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMediaNetworkAccessControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaNetworkAccessControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaNetworkAccessControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qmedianetworkaccesscontrol.h:68
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaNetworkAccessControl(QObject *)

/*

 */
func NewQMediaNetworkAccessControl__() *QMediaNetworkAccessControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN26QMediaNetworkAccessControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaNetworkAccessControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaNetworkAccessControl")
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
