package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h
// #include <qnetworkconfigmanager.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QNetworkConfigurationManager struct {
	*qtcore.QObject
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QNetworkConfigurationManager) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK28QNetworkConfigurationManager10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:69
// index:0
// Public
// void QNetworkConfigurationManager(QObject *)
func NewQNetworkConfigurationManager(parent *qtcore.QObject /*777 QObject **/) *QNetworkConfigurationManager {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN28QNetworkConfigurationManagerC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkConfigurationManagerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:70
// index:0
// Public virtual
// void ~QNetworkConfigurationManager()
func DeleteQNetworkConfigurationManager(*QNetworkConfigurationManager) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN28QNetworkConfigurationManagerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:72
// index:0
// Public
// QNetworkConfigurationManager::Capabilities capabilities()
func (this *QNetworkConfigurationManager) Capabilities() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK28QNetworkConfigurationManager12capabilitiesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:74
// index:0
// Public
// QNetworkConfiguration defaultConfiguration()
func (this *QNetworkConfigurationManager) DefaultConfiguration() *QNetworkConfiguration /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK28QNetworkConfigurationManager20defaultConfigurationEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:76
// index:0
// Public
// QNetworkConfiguration configurationFromIdentifier(const QString &)
func (this *QNetworkConfigurationManager) ConfigurationFromIdentifier(identifier *qtcore.QString) *QNetworkConfiguration /*123*/ {
	var convArg0 = identifier.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK28QNetworkConfigurationManager27configurationFromIdentifierERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:78
// index:0
// Public
// bool isOnline()
func (this *QNetworkConfigurationManager) IsOnline() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK28QNetworkConfigurationManager8isOnlineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:81
// index:0
// Public
// void updateConfigurations()
func (this *QNetworkConfigurationManager) UpdateConfigurations() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager20updateConfigurationsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:84
// index:0
// Public
// void configurationAdded(const QNetworkConfiguration &)
func (this *QNetworkConfigurationManager) ConfigurationAdded(config *QNetworkConfiguration) {
	var convArg0 = config.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager18configurationAddedERK21QNetworkConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:85
// index:0
// Public
// void configurationRemoved(const QNetworkConfiguration &)
func (this *QNetworkConfigurationManager) ConfigurationRemoved(config *QNetworkConfiguration) {
	var convArg0 = config.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager20configurationRemovedERK21QNetworkConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:86
// index:0
// Public
// void configurationChanged(const QNetworkConfiguration &)
func (this *QNetworkConfigurationManager) ConfigurationChanged(config *QNetworkConfiguration) {
	var convArg0 = config.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager20configurationChangedERK21QNetworkConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:87
// index:0
// Public
// void onlineStateChanged(_Bool)
func (this *QNetworkConfigurationManager) OnlineStateChanged(isOnline bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager18onlineStateChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), isOnline)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:88
// index:0
// Public
// void updateCompleted()
func (this *QNetworkConfigurationManager) UpdateCompleted() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN28QNetworkConfigurationManager15updateCompletedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QNetworkConfigurationManager__Capability = int

const QNetworkConfigurationManager__CanStartAndStopInterfaces QNetworkConfigurationManager__Capability = 1
const QNetworkConfigurationManager__DirectConnectionRouting QNetworkConfigurationManager__Capability = 2
const QNetworkConfigurationManager__SystemSessionSupport QNetworkConfigurationManager__Capability = 4
const QNetworkConfigurationManager__ApplicationLevelRoaming QNetworkConfigurationManager__Capability = 8
const QNetworkConfigurationManager__ForcedRoaming QNetworkConfigurationManager__Capability = 16
const QNetworkConfigurationManager__DataStatistics QNetworkConfigurationManager__Capability = 32
const QNetworkConfigurationManager__NetworkSessionRequired QNetworkConfigurationManager__Capability = 64

//  body block end
