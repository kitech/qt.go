package qtnetwork

// /usr/include/qt/QtNetwork/qnetworksession.h
// #include <qnetworksession.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
type QNetworkSession struct {
	*qtcore.QObject
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QNetworkSession) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworksession.h:90
// index:0
// Public
// void QNetworkSession(const QNetworkConfiguration &, QObject *)
func NewQNetworkSession(connConfig *QNetworkConfiguration, parent *qtcore.QObject /*777 QObject **/) *QNetworkSession {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = connConfig.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSessionC2ERK21QNetworkConfigurationP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkSessionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworksession.h:91
// index:0
// Public virtual
// void ~QNetworkSession()
func DeleteQNetworkSession(*QNetworkSession) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSessionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:93
// index:0
// Public
// bool isOpen()
func (this *QNetworkSession) IsOpen() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession6isOpenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworksession.h:94
// index:0
// Public
// QNetworkConfiguration configuration()
func (this *QNetworkSession) Configuration() *QNetworkConfiguration /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession13configurationEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworksession.h:96
// index:0
// Public
// QNetworkInterface interface()
func (this *QNetworkSession) Interface() *QNetworkInterface /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession9interfaceEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworksession.h:99
// index:0
// Public
// QNetworkSession::State state()
func (this *QNetworkSession) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:100
// index:0
// Public
// QNetworkSession::SessionError error()
func (this *QNetworkSession) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:128
// index:1
// Public
// void error(QNetworkSession::SessionError)
func (this *QNetworkSession) Error_1(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession5errorENS_12SessionErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:101
// index:0
// Public
// QString errorString()
func (this *QNetworkSession) ErrorString() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworksession.h:102
// index:0
// Public
// QVariant sessionProperty(const QString &)
func (this *QNetworkSession) SessionProperty(key *qtcore.QString) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession15sessionPropertyERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworksession.h:103
// index:0
// Public
// void setSessionProperty(const QString &, const QVariant &)
func (this *QNetworkSession) SetSessionProperty(key *qtcore.QString, value *qtcore.QVariant) {
	var convArg0 = key.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession18setSessionPropertyERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:105
// index:0
// Public
// quint64 bytesWritten()
func (this *QNetworkSession) BytesWritten() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession12bytesWrittenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworksession.h:106
// index:0
// Public
// quint64 bytesReceived()
func (this *QNetworkSession) BytesReceived() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession13bytesReceivedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworksession.h:107
// index:0
// Public
// quint64 activeTime()
func (this *QNetworkSession) ActiveTime() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession10activeTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworksession.h:109
// index:0
// Public
// QNetworkSession::UsagePolicies usagePolicies()
func (this *QNetworkSession) UsagePolicies() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QNetworkSession13usagePoliciesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:111
// index:0
// Public
// bool waitForOpened(int)
func (this *QNetworkSession) WaitForOpened(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession13waitForOpenedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworksession.h:114
// index:0
// Public
// void open()
func (this *QNetworkSession) Open() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession4openEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:115
// index:0
// Public
// void close()
func (this *QNetworkSession) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:116
// index:0
// Public
// void stop()
func (this *QNetworkSession) Stop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession4stopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:119
// index:0
// Public
// void migrate()
func (this *QNetworkSession) Migrate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession7migrateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:120
// index:0
// Public
// void ignore()
func (this *QNetworkSession) Ignore() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession6ignoreEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:121
// index:0
// Public
// void accept()
func (this *QNetworkSession) Accept() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession6acceptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:122
// index:0
// Public
// void reject()
func (this *QNetworkSession) Reject() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession6rejectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:125
// index:0
// Public
// void stateChanged(QNetworkSession::State)
func (this *QNetworkSession) StateChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession12stateChangedENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:126
// index:0
// Public
// void opened()
func (this *QNetworkSession) Opened() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession6openedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:127
// index:0
// Public
// void closed()
func (this *QNetworkSession) Closed() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession6closedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:129
// index:0
// Public
// void preferredConfigurationChanged(const QNetworkConfiguration &, bool)
func (this *QNetworkSession) PreferredConfigurationChanged(config *QNetworkConfiguration, isSeamless bool) {
	var convArg0 = config.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession29preferredConfigurationChangedERK21QNetworkConfigurationb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, isSeamless)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:130
// index:0
// Public
// void newConfigurationActivated()
func (this *QNetworkSession) NewConfigurationActivated() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession25newConfigurationActivatedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:134
// index:0
// Protected virtual
// void connectNotify(const QMetaMethod &)
func (this *QNetworkSession) ConnectNotify(signal *qtcore.QMetaMethod) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession13connectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:135
// index:0
// Protected virtual
// void disconnectNotify(const QMetaMethod &)
func (this *QNetworkSession) DisconnectNotify(signal *qtcore.QMetaMethod) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QNetworkSession16disconnectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
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
