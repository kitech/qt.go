package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h
// #include <qnetworkconfiguration.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 46
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
type QNetworkConfiguration struct {
	*qtrt.CObject
}
type QNetworkConfiguration_ITF interface {
	QNetworkConfiguration_PTR() *QNetworkConfiguration
}

func (ptr *QNetworkConfiguration) QNetworkConfiguration_PTR() *QNetworkConfiguration { return ptr }

func (this *QNetworkConfiguration) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkConfiguration) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQNetworkConfigurationFromPointer(cthis unsafe.Pointer) *QNetworkConfiguration {
	return &QNetworkConfiguration{&qtrt.CObject{cthis}}
}
func (*QNetworkConfiguration) NewFromPointer(cthis unsafe.Pointer) *QNetworkConfiguration {
	return NewQNetworkConfigurationFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkConfiguration()

/*
Constructs an invalid configuration object.

See also isValid().
*/
func (*QNetworkConfiguration) NewForInherit() *QNetworkConfiguration {
	return NewQNetworkConfiguration()
}
func NewQNetworkConfiguration() *QNetworkConfiguration {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkConfigurationC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkConfiguration)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QNetworkConfiguration & operator=(QNetworkConfiguration &&)

/*

 */
func (this *QNetworkConfiguration) Operator_equal(other unsafe.Pointer /*333*/) *QNetworkConfiguration {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkConfigurationaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:61
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkConfiguration & operator=(const QNetworkConfiguration &)

/*

 */
func (this *QNetworkConfiguration) Operator_equal1(other QNetworkConfiguration_ITF) *QNetworkConfiguration {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkConfiguration_PTR() != nil {
		convArg0 = other.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkConfigurationaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QNetworkConfiguration()

/*

 */
func DeleteQNetworkConfiguration(this *QNetworkConfiguration) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkConfigurationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkConfiguration &)

/*
Swaps this network configuration with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QNetworkConfiguration) Swap(other QNetworkConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkConfiguration_PTR() != nil {
		convArg0 = other.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkConfiguration4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QNetworkConfiguration &) const

/*

 */
func (this *QNetworkConfiguration) Operator_equal_equal(other QNetworkConfiguration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkConfiguration_PTR() != nil {
		convArg0 = other.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfigurationeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QNetworkConfiguration &) const

/*

 */
func (this *QNetworkConfiguration) Operator_not_equal(other QNetworkConfiguration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkConfiguration_PTR() != nil {
		convArg0 = other.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfigurationneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfiguration::StateFlags state() const

/*
Returns the current state of the configuration.
*/
func (this *QNetworkConfiguration) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfiguration::Type type() const

/*
Returns the type of the configuration.

A configuration can represent a single access point configuration or a set of access point configurations. Such a set is called service network. A configuration that is based on a service network can potentially support roaming of network sessions.
*/
func (this *QNetworkConfiguration) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:110
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfiguration::Purpose purpose() const

/*
Returns the purpose of this configuration.

The purpose field may be used to programmatically determine the purpose of a configuration. Such information is usually part of the access point or service network meta data.
*/
func (this *QNetworkConfiguration) Purpose() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration7purposeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfiguration::BearerType bearerType() const

/*
Returns the type of bearer used by this network configuration.

If the bearer type is unknown the bearerTypeName() function can be used to retrieve a textural type name for the bearer.

An invalid network configuration always returns the BearerUnknown value.

See also bearerTypeName() and bearerTypeFamily().
*/
func (this *QNetworkConfiguration) BearerType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration10bearerTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfiguration::BearerType bearerTypeFamily() const

/*
Returns the bearer type family used by this network configuration. The following table lists how bearerType() values map to bearerTypeFamily() values:


 bearer typebearer type family
BearerUnknown, Bearer2G, BearerEthernet, BearerWLAN, BearerBluetooth(same type)
BearerCDMA2000, BearerEVDO, BearerWCDMA, BearerHSPA, Bearer3GBearer3G
BearerWiMAX, BearerLTE, Bearer4GBearer4G


An invalid network configuration always returns the BearerUnknown value.

This function was introduced in  Qt 5.2.

See also bearerType() and bearerTypeName().
*/
func (this *QNetworkConfiguration) BearerTypeFamily() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration16bearerTypeFamilyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QString bearerTypeName() const

/*
Returns the type of bearer used by this network configuration as a string.

The string is not translated and therefore can not be shown to the user. The subsequent table shows the fixed mappings between BearerType and the bearer type name for known types. If the BearerType is unknown this function may return additional information if it is available; otherwise an empty string will be returned.


 BearerTypeValue
BearerUnknownThe session is based on an unknown or unspecified bearer type. The value of the string returned describes the bearer type.
BearerEthernetEthernet
BearerWLANWLAN
Bearer2G2G
Bearer3G3G
Bearer4G4G
BearerCDMA2000CDMA2000
BearerWCDMAWCDMA
BearerHSPAHSPA
BearerBluetoothBluetooth
BearerWiMAXWiMAX
BearerEVDOEVDO
BearerLTELTE


This function returns an empty string if this is an invalid configuration, a network configuration of type QNetworkConfiguration::ServiceNetwork or QNetworkConfiguration::UserChoice.

See also bearerType() and bearerTypeFamily().
*/
func (this *QNetworkConfiguration) BearerTypeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration14bearerTypeNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QString identifier() const

/*
Returns the unique and platform specific identifier for this network configuration; otherwise an empty string.
*/
func (this *QNetworkConfiguration) Identifier() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration10identifierEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRoamingAvailable() const

/*
Returns true if this configuration supports roaming; otherwise false.
*/
func (this *QNetworkConfiguration) IsRoamingAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration18isRoamingAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns the user visible name of this configuration.

The name may either be the name of the underlying access point or the name for service network that this configuration represents.
*/
func (this *QNetworkConfiguration) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this QNetworkConfiguration object is valid. A configuration may become invalid if the user deletes the configuration or the configuration was default-constructed.

The addition and removal of configurations can be monitored via the QNetworkConfigurationManager.

See also QNetworkConfigurationManager.
*/
func (this *QNetworkConfiguration) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:123
// index:0
// Public Visibility=Default Availability=Available
// [4] int connectTimeout() const

/*
Returns the connect timeout of this configuration.

This function was introduced in  Qt 5.9.

See also setConnectTimeout.
*/
func (this *QNetworkConfiguration) ConnectTimeout() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkConfiguration14connectTimeoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:124
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setConnectTimeout(int)

/*
Sets the connect timeout of this configuration to timeout. This allows control of the timeout used by QAbstractSocket to establish a connection.

Warning: This will have no effect if the bearer plugin doesn't have the CanStartAndStopInterfaces capability.

Returns true if succeeded.

This function was introduced in  Qt 5.9.

See also connectTimeout.
*/
func (this *QNetworkConfiguration) SetConnectTimeout(timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkConfiguration17setConnectTimeoutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum describes the type of configuration.


*/
type QNetworkConfiguration__Type = int

// The configuration specifies the details for a single access point. Note that configurations of type InternetAccessPoint may be part of other QNetworkConfigurations of type ServiceNetwork.
const QNetworkConfiguration__InternetAccessPoint QNetworkConfiguration__Type = 0

// The configuration is based on a group of QNetworkConfigurations of type InternetAccessPoint. All group members can reach the same target network. This type of configuration is a mandatory requirement for roaming enabled network sessions. On some platforms this form of configuration may also be called Service Network Access Point (SNAP).
const QNetworkConfiguration__ServiceNetwork QNetworkConfiguration__Type = 1

// The configuration is a placeholder which will be resolved to an actual configuration by the platform when a session is opened. Depending on the platform the selection may generate a popup dialog asking the user for his preferred choice.
const QNetworkConfiguration__UserChoice QNetworkConfiguration__Type = 2

// The configuration is invalid.
const QNetworkConfiguration__Invalid QNetworkConfiguration__Type = 3

func (this *QNetworkConfiguration) TypeItemName(val int) string {
	switch val {
	case QNetworkConfiguration__InternetAccessPoint: // 0
		return "InternetAccessPoint"
	case QNetworkConfiguration__ServiceNetwork: // 1
		return "ServiceNetwork"
	case QNetworkConfiguration__UserChoice: // 2
		return "UserChoice"
	case QNetworkConfiguration__Invalid: // 3
		return "Invalid"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkConfiguration_TypeItemName(val int) string {
	var nilthis *QNetworkConfiguration
	return nilthis.TypeItemName(val)
}

/*
Specifies the purpose of the configuration.


*/
type QNetworkConfiguration__Purpose = int

// The configuration doesn't specify any purpose. This is the default value.
const QNetworkConfiguration__UnknownPurpose QNetworkConfiguration__Purpose = 0

// The configuration can be used for general purpose internet access.
const QNetworkConfiguration__PublicPurpose QNetworkConfiguration__Purpose = 1

// The configuration is suitable to access a private network such as an office Intranet.
const QNetworkConfiguration__PrivatePurpose QNetworkConfiguration__Purpose = 2

// The configuration can be used for operator specific services (e.g. receiving MMS messages or content streaming).
const QNetworkConfiguration__ServiceSpecificPurpose QNetworkConfiguration__Purpose = 3

func (this *QNetworkConfiguration) PurposeItemName(val int) string {
	switch val {
	case QNetworkConfiguration__UnknownPurpose: // 0
		return "UnknownPurpose"
	case QNetworkConfiguration__PublicPurpose: // 1
		return "PublicPurpose"
	case QNetworkConfiguration__PrivatePurpose: // 2
		return "PrivatePurpose"
	case QNetworkConfiguration__ServiceSpecificPurpose: // 3
		return "ServiceSpecificPurpose"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkConfiguration_PurposeItemName(val int) string {
	var nilthis *QNetworkConfiguration
	return nilthis.PurposeItemName(val)
}

/*


 */
type QNetworkConfiguration__StateFlag = int

//
const QNetworkConfiguration__Undefined QNetworkConfiguration__StateFlag = 1

//
const QNetworkConfiguration__Defined QNetworkConfiguration__StateFlag = 2

//
const QNetworkConfiguration__Discovered QNetworkConfiguration__StateFlag = 6

//
const QNetworkConfiguration__Active QNetworkConfiguration__StateFlag = 14

func (this *QNetworkConfiguration) StateFlagItemName(val int) string {
	switch val {
	case QNetworkConfiguration__Undefined: // 1
		return "Undefined"
	case QNetworkConfiguration__Defined: // 2
		return "Defined"
	case QNetworkConfiguration__Discovered: // 6
		return "Discovered"
	case QNetworkConfiguration__Active: // 14
		return "Active"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkConfiguration_StateFlagItemName(val int) string {
	var nilthis *QNetworkConfiguration
	return nilthis.StateFlagItemName(val)
}

/*
Specifies the type of bearer used by a configuration.


*/
type QNetworkConfiguration__BearerType = int

// The type of bearer is unknown or unspecified. The bearerTypeName() function may return additional information.
const QNetworkConfiguration__BearerUnknown QNetworkConfiguration__BearerType = 0

// The configuration is for an Ethernet interfaces.
const QNetworkConfiguration__BearerEthernet QNetworkConfiguration__BearerType = 1

// The configuration is for a Wireless LAN interface.
const QNetworkConfiguration__BearerWLAN QNetworkConfiguration__BearerType = 2

// The configuration is for a CSD, GPRS, HSCSD, EDGE or cdmaOne interface.
const QNetworkConfiguration__Bearer2G QNetworkConfiguration__BearerType = 3

// The configuration is for CDMA interface.
const QNetworkConfiguration__BearerCDMA2000 QNetworkConfiguration__BearerType = 4

// The configuration is for W-CDMA/UMTS interface.
const QNetworkConfiguration__BearerWCDMA QNetworkConfiguration__BearerType = 5

// The configuration is for High Speed Packet Access (HSPA) interface.
const QNetworkConfiguration__BearerHSPA QNetworkConfiguration__BearerType = 6

// The configuration is for a Bluetooth interface.
const QNetworkConfiguration__BearerBluetooth QNetworkConfiguration__BearerType = 7

// The configuration is for a WiMAX interface.
const QNetworkConfiguration__BearerWiMAX QNetworkConfiguration__BearerType = 8

//
const QNetworkConfiguration__BearerEVDO QNetworkConfiguration__BearerType = 9

//
const QNetworkConfiguration__BearerLTE QNetworkConfiguration__BearerType = 10

//
const QNetworkConfiguration__Bearer3G QNetworkConfiguration__BearerType = 11

//
const QNetworkConfiguration__Bearer4G QNetworkConfiguration__BearerType = 12

func (this *QNetworkConfiguration) BearerTypeItemName(val int) string {
	switch val {
	case QNetworkConfiguration__BearerUnknown: // 0
		return "BearerUnknown"
	case QNetworkConfiguration__BearerEthernet: // 1
		return "BearerEthernet"
	case QNetworkConfiguration__BearerWLAN: // 2
		return "BearerWLAN"
	case QNetworkConfiguration__Bearer2G: // 3
		return "Bearer2G"
	case QNetworkConfiguration__BearerCDMA2000: // 4
		return "BearerCDMA2000"
	case QNetworkConfiguration__BearerWCDMA: // 5
		return "BearerWCDMA"
	case QNetworkConfiguration__BearerHSPA: // 6
		return "BearerHSPA"
	case QNetworkConfiguration__BearerBluetooth: // 7
		return "BearerBluetooth"
	case QNetworkConfiguration__BearerWiMAX: // 8
		return "BearerWiMAX"
	case QNetworkConfiguration__BearerEVDO: // 9
		return "BearerEVDO"
	case QNetworkConfiguration__BearerLTE: // 10
		return "BearerLTE"
	case QNetworkConfiguration__Bearer3G: // 11
		return "Bearer3G"
	case QNetworkConfiguration__Bearer4G: // 12
		return "Bearer4G"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkConfiguration_BearerTypeItemName(val int) string {
	var nilthis *QNetworkConfiguration
	return nilthis.BearerTypeItemName(val)
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
