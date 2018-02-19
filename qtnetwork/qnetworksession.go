package qtnetwork

// /usr/include/qt/QtNetwork/qnetworksession.h
// #include <qnetworksession.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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

// void connectNotify(const class QMetaMethod &)
func (this *QNetworkSession) InheritConnectNotify(f func(signal *qtcore.QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "connectNotify", f)
}

// void disconnectNotify(const class QMetaMethod &)
func (this *QNetworkSession) InheritDisconnectNotify(f func(signal *qtcore.QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "disconnectNotify", f)
}

type QNetworkSession struct {
	*qtcore.QObject
}
type QNetworkSession_ITF interface {
	qtcore.QObject_ITF
	QNetworkSession_PTR() *QNetworkSession
}

func (ptr *QNetworkSession) QNetworkSession_PTR() *QNetworkSession { return ptr }

func (this *QNetworkSession) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QNetworkSession) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQNetworkSessionFromPointer(cthis unsafe.Pointer) *QNetworkSession {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QNetworkSession{bcthis0}
}
func (*QNetworkSession) NewFromPointer(cthis unsafe.Pointer) *QNetworkSession {
	return NewQNetworkSessionFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QNetworkSession) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworksession.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkSession(const QNetworkConfiguration &, QObject *)
func NewQNetworkSession(connConfig QNetworkConfiguration_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkSession {
	var convArg0 unsafe.Pointer
	if connConfig != nil && connConfig.QNetworkConfiguration_PTR() != nil {
		convArg0 = connConfig.QNetworkConfiguration_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSessionC2ERK21QNetworkConfigurationP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkSessionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkSession")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworksession.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkSession(const QNetworkConfiguration &, QObject *)
func NewQNetworkSession__(connConfig QNetworkConfiguration_ITF) *QNetworkSession {
	var convArg0 unsafe.Pointer
	if connConfig != nil && connConfig.QNetworkConfiguration_PTR() != nil {
		convArg0 = connConfig.QNetworkConfiguration_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSessionC2ERK21QNetworkConfigurationP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkSessionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkSession")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworksession.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNetworkSession()
func DeleteQNetworkSession(this *QNetworkSession) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSessionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isOpen() const
func (this *QNetworkSession) IsOpen() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession6isOpenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworksession.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkConfiguration configuration() const
func (this *QNetworkSession) Configuration() *QNetworkConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession13configurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworksession.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkInterface interface() const
func (this *QNetworkSession) Interface() *QNetworkInterface /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession9interfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkInterface)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworksession.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkSession::State state() const
func (this *QNetworkSession) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkSession::SessionError error() const
func (this *QNetworkSession) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:128
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QNetworkSession::SessionError)
func (this *QNetworkSession) Error_1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession5errorENS_12SessionErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const
func (this *QNetworkSession) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworksession.h:102
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant sessionProperty(const QString &) const
func (this *QNetworkSession) SessionProperty(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession15sessionPropertyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworksession.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSessionProperty(const QString &, const QVariant &)
func (this *QNetworkSession) SetSessionProperty(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession18setSessionPropertyERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] quint64 bytesWritten() const
func (this *QNetworkSession) BytesWritten() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession12bytesWrittenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworksession.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] quint64 bytesReceived() const
func (this *QNetworkSession) BytesReceived() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession13bytesReceivedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworksession.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] quint64 activeTime() const
func (this *QNetworkSession) ActiveTime() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession10activeTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworksession.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkSession::UsagePolicies usagePolicies() const
func (this *QNetworkSession) UsagePolicies() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession13usagePoliciesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForOpened(int)
func (this *QNetworkSession) WaitForOpened(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession13waitForOpenedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworksession.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForOpened(int)
func (this *QNetworkSession) WaitForOpened__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession13waitForOpenedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworksession.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open()
func (this *QNetworkSession) Open() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession4openEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void close()
func (this *QNetworkSession) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()
func (this *QNetworkSession) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void migrate()
func (this *QNetworkSession) Migrate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession7migrateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ignore()
func (this *QNetworkSession) Ignore() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession6ignoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accept()
func (this *QNetworkSession) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reject()
func (this *QNetworkSession) Reject() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession6rejectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QNetworkSession::State)
func (this *QNetworkSession) StateChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession12stateChangedENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opened()
func (this *QNetworkSession) Opened() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession6openedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closed()
func (this *QNetworkSession) Closed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession6closedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void preferredConfigurationChanged(const QNetworkConfiguration &, _Bool)
func (this *QNetworkSession) PreferredConfigurationChanged(config QNetworkConfiguration_ITF, isSeamless bool) {
	var convArg0 unsafe.Pointer
	if config != nil && config.QNetworkConfiguration_PTR() != nil {
		convArg0 = config.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession29preferredConfigurationChangedERK21QNetworkConfigurationb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, isSeamless)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void newConfigurationActivated()
func (this *QNetworkSession) NewConfigurationActivated() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession25newConfigurationActivatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void usagePoliciesChanged(QNetworkSession::UsagePolicies)
func (this *QNetworkSession) UsagePoliciesChanged(usagePolicies int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession20usagePoliciesChangedE6QFlagsINS_11UsagePolicyEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), usagePolicies)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:134
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void connectNotify(const QMetaMethod &)
func (this *QNetworkSession) ConnectNotify(signal qtcore.QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession13connectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:135
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void disconnectNotify(const QMetaMethod &)
func (this *QNetworkSession) DisconnectNotify(signal qtcore.QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession16disconnectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QNetworkSession__State = int

const QNetworkSession__Invalid QNetworkSession__State = 0
const QNetworkSession__NotAvailable QNetworkSession__State = 1
const QNetworkSession__Connecting QNetworkSession__State = 2
const QNetworkSession__Connected QNetworkSession__State = 3
const QNetworkSession__Closing QNetworkSession__State = 4
const QNetworkSession__Disconnected QNetworkSession__State = 5
const QNetworkSession__Roaming QNetworkSession__State = 6

type QNetworkSession__SessionError = int

const QNetworkSession__UnknownSessionError QNetworkSession__SessionError = 0
const QNetworkSession__SessionAbortedError QNetworkSession__SessionError = 1
const QNetworkSession__RoamingError QNetworkSession__SessionError = 2
const QNetworkSession__OperationNotSupportedError QNetworkSession__SessionError = 3
const QNetworkSession__InvalidConfigurationError QNetworkSession__SessionError = 4

type QNetworkSession__UsagePolicy = int

const QNetworkSession__NoPolicy QNetworkSession__UsagePolicy = 0
const QNetworkSession__NoBackgroundTrafficPolicy QNetworkSession__UsagePolicy = 1

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
