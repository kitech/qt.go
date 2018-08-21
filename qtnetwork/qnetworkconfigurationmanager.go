package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h
// #include <qnetworkconfigmanager.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QNetworkConfigurationManager struct {
	*qtcore.QObject
}
type QNetworkConfigurationManager_ITF interface {
	qtcore.QObject_ITF
	QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager
}

func (ptr *QNetworkConfigurationManager) QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager {
	return ptr
}

func (this *QNetworkConfigurationManager) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QNetworkConfigurationManager) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQNetworkConfigurationManagerFromPointer(cthis unsafe.Pointer) *QNetworkConfigurationManager {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QNetworkConfigurationManager{bcthis0}
}
func (*QNetworkConfigurationManager) NewFromPointer(cthis unsafe.Pointer) *QNetworkConfigurationManager {
	return NewQNetworkConfigurationManagerFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QNetworkConfigurationManager) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QNetworkConfigurationManager10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkConfigurationManager(QObject *)

/*

 */
func NewQNetworkConfigurationManager(parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkConfigurationManager {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QNetworkConfigurationManagerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkConfigurationManagerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkConfigurationManager")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkConfigurationManager(QObject *)

/*

 */
func NewQNetworkConfigurationManager__() *QNetworkConfigurationManager {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN28QNetworkConfigurationManagerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkConfigurationManagerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkConfigurationManager")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNetworkConfigurationManager()

/*

 */
func DeleteQNetworkConfigurationManager(this *QNetworkConfigurationManager) {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QNetworkConfigurationManagerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:72
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfigurationManager::Capabilities capabilities() const

/*

 */
func (this *QNetworkConfigurationManager) Capabilities() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QNetworkConfigurationManager12capabilitiesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkConfiguration defaultConfiguration() const

/*

 */
func (this *QNetworkConfigurationManager) DefaultConfiguration() *QNetworkConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QNetworkConfigurationManager20defaultConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkConfiguration configurationFromIdentifier(const QString &) const

/*

 */
func (this *QNetworkConfigurationManager) ConfigurationFromIdentifier(identifier string) *QNetworkConfiguration /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(identifier)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QNetworkConfigurationManager27configurationFromIdentifierERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isOnline() const

/*

 */
func (this *QNetworkConfigurationManager) IsOnline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK28QNetworkConfigurationManager8isOnlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateConfigurations()

/*

 */
func (this *QNetworkConfigurationManager) UpdateConfigurations() {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager20updateConfigurationsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void configurationAdded(const QNetworkConfiguration &)

/*

 */
func (this *QNetworkConfigurationManager) ConfigurationAdded(config QNetworkConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if config != nil && config.QNetworkConfiguration_PTR() != nil {
		convArg0 = config.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager18configurationAddedERK21QNetworkConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void configurationRemoved(const QNetworkConfiguration &)

/*

 */
func (this *QNetworkConfigurationManager) ConfigurationRemoved(config QNetworkConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if config != nil && config.QNetworkConfiguration_PTR() != nil {
		convArg0 = config.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager20configurationRemovedERK21QNetworkConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void configurationChanged(const QNetworkConfiguration &)

/*

 */
func (this *QNetworkConfigurationManager) ConfigurationChanged(config QNetworkConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if config != nil && config.QNetworkConfiguration_PTR() != nil {
		convArg0 = config.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager20configurationChangedERK21QNetworkConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void onlineStateChanged(bool)

/*

 */
func (this *QNetworkConfigurationManager) OnlineStateChanged(isOnline bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager18onlineStateChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), isOnline)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateCompleted()

/*

 */
func (this *QNetworkConfigurationManager) UpdateCompleted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager15updateCompletedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QNetworkConfigurationManager__Capability = int

//
const QNetworkConfigurationManager__CanStartAndStopInterfaces QNetworkConfigurationManager__Capability = 1

//
const QNetworkConfigurationManager__DirectConnectionRouting QNetworkConfigurationManager__Capability = 2

//
const QNetworkConfigurationManager__SystemSessionSupport QNetworkConfigurationManager__Capability = 4

//
const QNetworkConfigurationManager__ApplicationLevelRoaming QNetworkConfigurationManager__Capability = 8

//
const QNetworkConfigurationManager__ForcedRoaming QNetworkConfigurationManager__Capability = 16

//
const QNetworkConfigurationManager__DataStatistics QNetworkConfigurationManager__Capability = 32

//
const QNetworkConfigurationManager__NetworkSessionRequired QNetworkConfigurationManager__Capability = 64

func (this *QNetworkConfigurationManager) CapabilityItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QNetworkConfigurationManager_CapabilityItemName(val int) string {
	var nilthis *QNetworkConfigurationManager
	return nilthis.CapabilityItemName(val)
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
