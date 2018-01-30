package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h
// #include <qnetworkconfiguration.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 46
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QNetworkConfiguration struct {
	*qtrt.CObject
}

func (this *QNetworkConfiguration) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkConfiguration) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
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
func NewQNetworkConfiguration() *QNetworkConfiguration {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkConfigurationC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QNetworkConfiguration()
func DeleteQNetworkConfiguration(*QNetworkConfiguration) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkConfigurationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkConfiguration &)
func (this *QNetworkConfiguration) Swap(other *QNetworkConfiguration) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkConfiguration4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfiguration::StateFlags state()
func (this *QNetworkConfiguration) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfiguration::Type type()
func (this *QNetworkConfiguration) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:110
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfiguration::Purpose purpose()
func (this *QNetworkConfiguration) Purpose() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration7purposeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfiguration::BearerType bearerType()
func (this *QNetworkConfiguration) BearerType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration10bearerTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkConfiguration::BearerType bearerTypeFamily()
func (this *QNetworkConfiguration) BearerTypeFamily() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration16bearerTypeFamilyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QString bearerTypeName()
func (this *QNetworkConfiguration) BearerTypeName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration14bearerTypeNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QString identifier()
func (this *QNetworkConfiguration) Identifier() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration10identifierEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRoamingAvailable()
func (this *QNetworkConfiguration) IsRoamingAvailable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration18isRoamingAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QNetworkConfiguration) Name() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QNetworkConfiguration) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:123
// index:0
// Public Visibility=Default Availability=Available
// [4] int connectTimeout()
func (this *QNetworkConfiguration) ConnectTimeout() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkConfiguration14connectTimeoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:124
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setConnectTimeout(int)
func (this *QNetworkConfiguration) SetConnectTimeout(timeout int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkConfiguration17setConnectTimeoutEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QNetworkConfiguration__Type = int

const QNetworkConfiguration__InternetAccessPoint QNetworkConfiguration__Type = 0
const QNetworkConfiguration__ServiceNetwork QNetworkConfiguration__Type = 1
const QNetworkConfiguration__UserChoice QNetworkConfiguration__Type = 2
const QNetworkConfiguration__Invalid QNetworkConfiguration__Type = 3

type QNetworkConfiguration__Purpose = int

const QNetworkConfiguration__UnknownPurpose QNetworkConfiguration__Purpose = 0
const QNetworkConfiguration__PublicPurpose QNetworkConfiguration__Purpose = 1
const QNetworkConfiguration__PrivatePurpose QNetworkConfiguration__Purpose = 2
const QNetworkConfiguration__ServiceSpecificPurpose QNetworkConfiguration__Purpose = 3

type QNetworkConfiguration__StateFlag = int

const QNetworkConfiguration__Undefined QNetworkConfiguration__StateFlag = 1
const QNetworkConfiguration__Defined QNetworkConfiguration__StateFlag = 2
const QNetworkConfiguration__Discovered QNetworkConfiguration__StateFlag = 6
const QNetworkConfiguration__Active QNetworkConfiguration__StateFlag = 14

type QNetworkConfiguration__BearerType = int

const QNetworkConfiguration__BearerUnknown QNetworkConfiguration__BearerType = 0
const QNetworkConfiguration__BearerEthernet QNetworkConfiguration__BearerType = 1
const QNetworkConfiguration__BearerWLAN QNetworkConfiguration__BearerType = 2
const QNetworkConfiguration__Bearer2G QNetworkConfiguration__BearerType = 3
const QNetworkConfiguration__BearerCDMA2000 QNetworkConfiguration__BearerType = 4
const QNetworkConfiguration__BearerWCDMA QNetworkConfiguration__BearerType = 5
const QNetworkConfiguration__BearerHSPA QNetworkConfiguration__BearerType = 6
const QNetworkConfiguration__BearerBluetooth QNetworkConfiguration__BearerType = 7
const QNetworkConfiguration__BearerWiMAX QNetworkConfiguration__BearerType = 8
const QNetworkConfiguration__BearerEVDO QNetworkConfiguration__BearerType = 9
const QNetworkConfiguration__BearerLTE QNetworkConfiguration__BearerType = 10
const QNetworkConfiguration__Bearer3G QNetworkConfiguration__BearerType = 11
const QNetworkConfiguration__Bearer4G QNetworkConfiguration__BearerType = 12

//  body block end
