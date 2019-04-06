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

// void connectNotify(const QMetaMethod &)
func (this *QNetworkSession) InheritConnectNotify(f func(signal *qtcore.QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "connectNotify", f)
}

// void disconnectNotify(const QMetaMethod &)
func (this *QNetworkSession) InheritDisconnectNotify(f func(signal *qtcore.QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "disconnectNotify", f)
}

/*

 */
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

/*

 */
func (this *QNetworkSession) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworksession.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkSession(const QNetworkConfiguration &, QObject *)

/*
Constructs a session based on connectionConfig with the given parent.

See also QNetworkConfiguration.
*/
func (*QNetworkSession) NewForInherit(connConfig QNetworkConfiguration_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkSession {
	return NewQNetworkSession(connConfig, parent)
}
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

/*
Constructs a session based on connectionConfig with the given parent.

See also QNetworkConfiguration.
*/
func (*QNetworkSession) NewForInheritp(connConfig QNetworkConfiguration_ITF) *QNetworkSession {
	return NewQNetworkSessionp(connConfig)
}
func NewQNetworkSessionp(connConfig QNetworkConfiguration_ITF) *QNetworkSession {
	var convArg0 unsafe.Pointer
	if connConfig != nil && connConfig.QNetworkConfiguration_PTR() != nil {
		convArg0 = connConfig.QNetworkConfiguration_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
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

/*

 */
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

/*
Returns true if this session is open. If the number of all open sessions is greater than zero the underlying network interface will remain connected/up.

The session can be controlled via open() and close().
*/
func (this *QNetworkSession) IsOpen() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession6isOpenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworksession.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkConfiguration configuration() const

/*
Returns the QNetworkConfiguration that this network session object is based on.

See also QNetworkConfiguration.
*/
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

/*
Returns the network interface that is used by this session.

This function only returns a valid QNetworkInterface when this session is Connected.

The returned interface may change as a result of a roaming process.

See also state().
*/
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

/*
Returns the state of the session.

If the session is based on a single access point configuration the state of the session is the same as the state of the associated network interface. Therefore a network session object can be used to monitor network interfaces.

A QNetworkConfiguration::ServiceNetwork based session summarizes the state of all its children and therefore returns the Connected state if at least one of the service network's children() configurations is active.

Note that it is not required to hold an open session in order to obtain the network interface state. A connected but closed session may be used to monitor network interfaces whereas an open and connected session object may prevent the network interface from being shut down.

See also error() and stateChanged().
*/
func (this *QNetworkSession) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkSession::SessionError error() const

/*
Returns the type of error that last occurred.

See also state() and errorString().
*/
func (this *QNetworkSession) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:128
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QNetworkSession::SessionError)

/*
Returns the type of error that last occurred.

See also state() and errorString().
*/
func (this *QNetworkSession) Error1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession5errorENS_12SessionErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a human-readable description of the last device error that occurred.

See also error().
*/
func (this *QNetworkSession) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworksession.h:102
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant sessionProperty(const QString &) const

/*
Returns the value for property key.

A network session can have properties attached which may describe the session in more details. This function can be used to gain access to those properties.

The following property keys are guaranteed to be specified on all platforms:


 KeyDescription
ActiveConfigurationIf the session isOpen() this property returns the identifier of the QNetworkConfiguration that is used by this session; otherwise an empty string.The main purpose of this key is to determine which Internet access point is used if the session is based on a ServiceNetwork. The following code snippet highlights the difference:

      QNetworkConfigurationManager mgr;
      QNetworkConfiguration ap = mgr.defaultConfiguration();
      QNetworkSession *session = new QNetworkSession(ap);
      ... //code activates session

      QString ident = session->sessionProperty("ActiveConfiguration").toString();
      if ( ap.type() == QNetworkConfiguration::ServiceNetwork ) {
          Q_ASSERT( ap.identifier() != ident );
          Q_ASSERT( ap.children().contains( mgr.configurationFromIdentifier(ident) ) );
      } else if ( ap.type() == QNetworkConfiguration::InternetAccessPoint ) {
          Q_ASSERT( ap.identifier() == ident );
      }
                  \endcode



UserChoiceConfigurationIf the session isOpen() and is bound to a QNetworkConfiguration of type UserChoice, this property returns the identifier of the QNetworkConfiguration that the configuration resolved to when open() was called; otherwise an empty string.The purpose of this key is to determine the real QNetworkConfiguration that the session is using. This key is different from ActiveConfiguration in that this key may return an identifier for either a service network or a Internet access points configurations, whereas ActiveConfiguration always returns identifiers to Internet access points configurations.

ConnectInBackgroundSetting this property to true before calling open() implies that the connection attempt is made but if no connection can be established, the user is not connsulted and asked to select a suitable connection. This property is not set by default and support for it depends on the platform.
AutoCloseSessionTimeoutIf the session requires polling to keep its state up to date, this property holds the timeout in milliseconds before the session will automatically close. If the value of this property is -1 the session will not automatically close. This property is set to -1 by default.The purpose of this property is to minimize resource use on platforms that use polling to update the state of the session. Applications can set the value of this property to the desired timeout before the session is closed. In response to the closed() signal the network session should be deleted to ensure that all polling is stopped. The session can then be recreated once it is required again. This property has no effect for sessions that do not require polling.



See also setSessionProperty().
*/
func (this *QNetworkSession) SessionProperty(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString5(key)
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

/*
Sets the property value on the session. The property is identified using key. Removing an already set property can be achieved by passing an invalid QVariant.

Note that the UserChoiceConfiguration and ActiveConfiguration properties are read only and cannot be changed using this method.

See also sessionProperty().
*/
func (this *QNetworkSession) SetSessionProperty(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString5(key)
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

/*
Returns the amount of data sent in bytes; otherwise 0.

This field value includes the usage across all open network sessions which use the same network interface.

If the session is based on a service network configuration the number of sent bytes across all active member configurations are returned.

This function may not always be supported on all platforms and returns 0. The platform capability can be detected via QNetworkConfigurationManager::DataStatistics.

Note: On some platforms this function may run the main event loop.
*/
func (this *QNetworkSession) BytesWritten() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession12bytesWrittenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworksession.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] quint64 bytesReceived() const

/*
Returns the amount of data received in bytes; otherwise 0.

This field value includes the usage across all open network sessions which use the same network interface.

If the session is based on a service network configuration the number of sent bytes across all active member configurations are returned.

This function may not always be supported on all platforms and returns 0. The platform capability can be detected via QNetworkConfigurationManager::DataStatistics.

Note: On some platforms this function may run the main event loop.
*/
func (this *QNetworkSession) BytesReceived() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession13bytesReceivedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworksession.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] quint64 activeTime() const

/*
Returns the number of seconds that the session has been active.
*/
func (this *QNetworkSession) ActiveTime() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession10activeTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworksession.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkSession::UsagePolicies usagePolicies() const

/*
Returns the network usage policies currently in force by the system.
*/
func (this *QNetworkSession) UsagePolicies() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QNetworkSession13usagePoliciesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForOpened(int)

/*
Waits until the session has been opened, up to msecs milliseconds. If the session has been opened, this function returns true; otherwise it returns false. In the case where it returns false, you can call error() to determine the cause of the error.

The following example waits up to one second for the session to be opened:


      session->open();
      if (session->waitForOpened(1000))
          qDebug("Open!");



If msecs is -1, this function will not time out.

See also open() and error().
*/
func (this *QNetworkSession) WaitForOpened(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession13waitForOpenedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworksession.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForOpened(int)

/*
Waits until the session has been opened, up to msecs milliseconds. If the session has been opened, this function returns true; otherwise it returns false. In the case where it returns false, you can call error() to determine the cause of the error.

The following example waits up to one second for the session to be opened:


      session->open();
      if (session->waitForOpened(1000))
          qDebug("Open!");



If msecs is -1, this function will not time out.

See also open() and error().
*/
func (this *QNetworkSession) WaitForOpenedp() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession13waitForOpenedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworksession.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open()

/*
Creates an open session which increases the session counter on the underlying network interface. The system will not terminate a network interface until the session reference counter reaches zero. Therefore an open session allows an application to register its use of the interface.

As a result of calling open() the interface will be started if it is not connected/up yet. Some platforms may not provide support for out-of-process sessions. On such platforms the session counter ignores any sessions held by another process. The platform capabilities can be detected via QNetworkConfigurationManager::capabilities().

Note that this call is asynchronous. Depending on the outcome of this call the results can be enquired by connecting to the stateChanged(), opened() or error() signals.

It is not a requirement to open a session in order to monitor the underlying network interface.

See also close(), stop(), and isOpen().
*/
func (this *QNetworkSession) Open() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession4openEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void close()

/*
Decreases the session counter on the associated network configuration. If the session counter reaches zero the active network interface is shut down. This also means that state() will only change from Connected to Disconnected if the current session was the last open session.

If the platform does not support out-of-process sessions calling this function does not stop the interface. In this case stop() has to be used to force a shut down. The platform capabilities can be detected via QNetworkConfigurationManager::capabilities().

Note that this call is asynchronous. Depending on the outcome of this call the results can be enquired by connecting to the stateChanged(), opened() or error() signals.

See also open(), stop(), and isOpen().
*/
func (this *QNetworkSession) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Invalidates all open sessions against the network interface and therefore stops the underlying network interface. This function always changes the session's state() flag to Disconnected.

See also open() and close().
*/
func (this *QNetworkSession) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void migrate()

/*
Instructs the session to roam to the new access point. The old access point remains active until the application calls accept().

The newConfigurationActivated() signal is emitted once roaming has been completed.

See also accept().
*/
func (this *QNetworkSession) Migrate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession7migrateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ignore()

/*
This function indicates that the application does not wish to roam the session.

See also migrate().
*/
func (this *QNetworkSession) Ignore() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession6ignoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accept()

/*
Instructs the session to permanently accept the new access point. Once this function has been called the session may not return to the old access point.

The old access point may be closed in the process if there are no other network sessions for it. Therefore any open socket that still uses the old access point may become unusable and should be closed before completing the migration.
*/
func (this *QNetworkSession) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reject()

/*
The new access point is not suitable for the application. By calling this function the session returns to the previous access point/configuration. This action may invalidate any socket that has been created via the not desired access point.

See also accept().
*/
func (this *QNetworkSession) Reject() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession6rejectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QNetworkSession::State)

/*
This signal is emitted whenever the state of the network session changes. The state parameter is the new state.

See also state().
*/
func (this *QNetworkSession) StateChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession12stateChangedENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opened()

/*
This signal is emitted when the network session has been opened.

The underlying network interface will not be shut down as long as the session remains open. Note that this feature is dependent on system wide session support.
*/
func (this *QNetworkSession) Opened() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession6openedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closed()

/*
This signal is emitted when the network session has been closed.
*/
func (this *QNetworkSession) Closed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession6closedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void preferredConfigurationChanged(const QNetworkConfiguration &, bool)

/*
This signal is emitted when the preferred configuration/access point for the session changes. Only sessions which are based on service network configurations may emit this signal. config can be used to determine access point specific details such as proxy settings and isSeamless indicates whether roaming will break the sessions IP address.

As a consequence to this signal the application must either start the roaming process by calling migrate() or choose to ignore() the new access point.

If the roaming process is non-seamless the IP address will change which means that a socket becomes invalid. However seamless mobility can ensure that the local IP address does not change. This is achieved by using a virtual IP address which is bound to the actual link address. During the roaming process the virtual address is attached to the new link address.

Some platforms may support the concept of Forced Roaming and Application Level Roaming (ALR). Forced roaming implies that the platform may simply roam to a new configuration without consulting applications. It is up to the application to detect the link layer loss and reestablish its sockets. In contrast ALR provides the opportunity to prevent the system from roaming. If this session is based on a configuration that supports roaming the application can choose whether it wants to be consulted (ALR use case) by connecting to this signal. For as long as this signal connection remains the session remains registered as a roaming stakeholder; otherwise roaming will be enforced by the platform.

See also migrate(), ignore(), and QNetworkConfiguration::isRoamingAvailable().
*/
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

/*
This signal is emitted once the session has roamed to the new access point. The application may reopen its socket and test the suitability of the new network link. Subsequently it must either accept() or reject() the new access point.

See also accept() and reject().
*/
func (this *QNetworkSession) NewConfigurationActivated() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession25newConfigurationActivatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void usagePoliciesChanged(QNetworkSession::UsagePolicies)

/*
This signal is emitted when the usagePolicies in force are changed by the system.

This function was introduced in  Qt 5.0.
*/
func (this *QNetworkSession) UsagePoliciesChanged(usagePolicies int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession20usagePoliciesChangedE6QFlagsINS_11UsagePolicyEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), usagePolicies)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworksession.h:134
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void connectNotify(const QMetaMethod &)

/*

 */
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

/*

 */
func (this *QNetworkSession) DisconnectNotify(signal qtcore.QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QNetworkSession16disconnectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the connectivity state of the session. If the session is based on a single access point configuration the state of the session is the same as the state of the associated network interface.


*/
type QNetworkSession__State = int

// The session is invalid due to an invalid configuration. This may happen due to a removed access point or a configuration that was invalid to begin with.
const QNetworkSession__Invalid QNetworkSession__State = 0

// The session is based on a defined but not yet discovered QNetworkConfiguration (see QNetworkConfiguration::StateFlag).
const QNetworkSession__NotAvailable QNetworkSession__State = 1

// The network session is being established.
const QNetworkSession__Connecting QNetworkSession__State = 2

// The network session is connected. If the current process wishes to use this session it has to register its interest by calling open(). A network session is considered to be ready for socket operations if it isOpen() and connected.
const QNetworkSession__Connected QNetworkSession__State = 3

// The network session is in the process of being shut down.
const QNetworkSession__Closing QNetworkSession__State = 4

// The network session is not connected. The associated QNetworkConfiguration has the state QNetworkConfiguration::Discovered.
const QNetworkSession__Disconnected QNetworkSession__State = 5

// The network session is roaming from one access point to another access point.
const QNetworkSession__Roaming QNetworkSession__State = 6

func (this *QNetworkSession) StateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QNetworkSession_StateItemName(val int) string {
	var nilthis *QNetworkSession
	return nilthis.StateItemName(val)
}

/*
This enum describes the session errors that can occur.


*/
type QNetworkSession__SessionError = int

// An unidentified error occurred.
const QNetworkSession__UnknownSessionError QNetworkSession__SessionError = 0

// The session was aborted by the user or system.
const QNetworkSession__SessionAbortedError QNetworkSession__SessionError = 1

// The session cannot roam to a new configuration.
const QNetworkSession__RoamingError QNetworkSession__SessionError = 2

// The operation is not supported for current configuration.
const QNetworkSession__OperationNotSupportedError QNetworkSession__SessionError = 3

// The operation cannot currently be performed for the current configuration.
const QNetworkSession__InvalidConfigurationError QNetworkSession__SessionError = 4

func (this *QNetworkSession) SessionErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QNetworkSession_SessionErrorItemName(val int) string {
	var nilthis *QNetworkSession
	return nilthis.SessionErrorItemName(val)
}

/*


 */
type QNetworkSession__UsagePolicy = int

//
const QNetworkSession__NoPolicy QNetworkSession__UsagePolicy = 0

//
const QNetworkSession__NoBackgroundTrafficPolicy QNetworkSession__UsagePolicy = 1

func (this *QNetworkSession) UsagePolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QNetworkSession_UsagePolicyItemName(val int) string {
	var nilthis *QNetworkSession
	return nilthis.UsagePolicyItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11459() {
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
