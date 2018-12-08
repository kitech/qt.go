package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkproxy.h
// #include <qnetworkproxy.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
type QNetworkProxy struct {
	*qtrt.CObject
}
type QNetworkProxy_ITF interface {
	QNetworkProxy_PTR() *QNetworkProxy
}

func (ptr *QNetworkProxy) QNetworkProxy_PTR() *QNetworkProxy { return ptr }

func (this *QNetworkProxy) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkProxy) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQNetworkProxyFromPointer(cthis unsafe.Pointer) *QNetworkProxy {
	return &QNetworkProxy{&qtrt.CObject{cthis}}
}
func (*QNetworkProxy) NewFromPointer(cthis unsafe.Pointer) *QNetworkProxy {
	return NewQNetworkProxyFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxy()

/*
Constructs a QNetworkProxy with DefaultProxy type; the proxy type is determined by applicationProxy(), which defaults to NoProxy.

See also setType() and setApplicationProxy().
*/
func (*QNetworkProxy) NewForInherit() *QNetworkProxy {
	return NewQNetworkProxy()
}
func NewQNetworkProxy() *QNetworkProxy {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxyC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxy)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:161
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxy(QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)

/*
Constructs a QNetworkProxy with DefaultProxy type; the proxy type is determined by applicationProxy(), which defaults to NoProxy.

See also setType() and setApplicationProxy().
*/
func (*QNetworkProxy) NewForInherit1(type_ int, hostName string, port uint16, user string, password string) *QNetworkProxy {
	return NewQNetworkProxy1(type_, hostName, port, user, password)
}
func NewQNetworkProxy1(type_ int, hostName string, port uint16, user string, password string) *QNetworkProxy {
	var tmpArg1 = qtcore.NewQString5(hostName)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg3 = qtcore.NewQString5(user)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString5(password)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxyC2ENS_9ProxyTypeERK7QStringtS3_S3_", qtrt.FFI_TYPE_POINTER, type_, convArg1, port, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxy)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:161
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxy(QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)

/*
Constructs a QNetworkProxy with DefaultProxy type; the proxy type is determined by applicationProxy(), which defaults to NoProxy.

See also setType() and setApplicationProxy().
*/
func (*QNetworkProxy) NewForInherit1p(type_ int) *QNetworkProxy {
	return NewQNetworkProxy1p(type_)
}
func NewQNetworkProxy1p(type_ int) *QNetworkProxy {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(0)
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, const QString &=LValueReference, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxyC2ENS_9ProxyTypeERK7QStringtS3_S3_", qtrt.FFI_TYPE_POINTER, type_, convArg1, port, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxy)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:161
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxy(QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)

/*
Constructs a QNetworkProxy with DefaultProxy type; the proxy type is determined by applicationProxy(), which defaults to NoProxy.

See also setType() and setApplicationProxy().
*/
func (*QNetworkProxy) NewForInherit1p1(type_ int, hostName string) *QNetworkProxy {
	return NewQNetworkProxy1p1(type_, hostName)
}
func NewQNetworkProxy1p1(type_ int, hostName string) *QNetworkProxy {
	var tmpArg1 = qtcore.NewQString5(hostName)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, quint16=Typedef, quint16=Typedef, unsigned short, UShort
	port := uint16(0)
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, const QString &=LValueReference, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxyC2ENS_9ProxyTypeERK7QStringtS3_S3_", qtrt.FFI_TYPE_POINTER, type_, convArg1, port, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxy)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:161
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxy(QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)

/*
Constructs a QNetworkProxy with DefaultProxy type; the proxy type is determined by applicationProxy(), which defaults to NoProxy.

See also setType() and setApplicationProxy().
*/
func (*QNetworkProxy) NewForInherit1p2(type_ int, hostName string, port uint16) *QNetworkProxy {
	return NewQNetworkProxy1p2(type_, hostName, port)
}
func NewQNetworkProxy1p2(type_ int, hostName string, port uint16) *QNetworkProxy {
	var tmpArg1 = qtcore.NewQString5(hostName)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, const QString &=LValueReference, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxyC2ENS_9ProxyTypeERK7QStringtS3_S3_", qtrt.FFI_TYPE_POINTER, type_, convArg1, port, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxy)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:161
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxy(QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)

/*
Constructs a QNetworkProxy with DefaultProxy type; the proxy type is determined by applicationProxy(), which defaults to NoProxy.

See also setType() and setApplicationProxy().
*/
func (*QNetworkProxy) NewForInherit1p3(type_ int, hostName string, port uint16, user string) *QNetworkProxy {
	return NewQNetworkProxy1p3(type_, hostName, port, user)
}
func NewQNetworkProxy1p3(type_ int, hostName string, port uint16, user string) *QNetworkProxy {
	var tmpArg1 = qtcore.NewQString5(hostName)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg3 = qtcore.NewQString5(user)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 4, const QString &=LValueReference, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxyC2ENS_9ProxyTypeERK7QStringtS3_S3_", qtrt.FFI_TYPE_POINTER, type_, convArg1, port, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxy)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:165
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QNetworkProxy & operator=(QNetworkProxy &&)

/*

 */
func (this *QNetworkProxy) Operator_equal(other unsafe.Pointer /*333*/) *QNetworkProxy {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxyaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkProxy)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:167
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkProxy & operator=(const QNetworkProxy &)

/*

 */
func (this *QNetworkProxy) Operator_equal1(other QNetworkProxy_ITF) *QNetworkProxy {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkProxy_PTR() != nil {
		convArg0 = other.QNetworkProxy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxyaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkProxy)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QNetworkProxy()

/*

 */
func DeleteQNetworkProxy(this *QNetworkProxy) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:170
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkProxy &)

/*
Swaps this network proxy instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QNetworkProxy) Swap(other QNetworkProxy_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkProxy_PTR() != nil {
		convArg0 = other.QNetworkProxy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:172
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QNetworkProxy &) const

/*

 */
func (this *QNetworkProxy) Operator_equal_equal(other QNetworkProxy_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkProxy_PTR() != nil {
		convArg0 = other.QNetworkProxy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxyeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:173
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QNetworkProxy &) const

/*

 */
func (this *QNetworkProxy) Operator_not_equal(other QNetworkProxy_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkProxy_PTR() != nil {
		convArg0 = other.QNetworkProxy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxyneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setType(QNetworkProxy::ProxyType)

/*
Sets the proxy type for this instance to be type.

Note that changing the type of a proxy does not change the set of capabilities this QNetworkProxy object holds if any capabilities have been set with setCapabilities().

See also type() and setCapabilities().
*/
func (this *QNetworkProxy) SetType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy7setTypeENS_9ProxyTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:177
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkProxy::ProxyType type() const

/*
Returns the proxy type for this instance.

See also setType().
*/
func (this *QNetworkProxy) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCapabilities(QNetworkProxy::Capabilities)

/*
Sets the capabilities of this proxy to capabilities.

This function was introduced in  Qt 4.5.

See also setType() and capabilities().
*/
func (this *QNetworkProxy) SetCapabilities(capab int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy15setCapabilitiesE6QFlagsINS_10CapabilityEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), capab)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:180
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkProxy::Capabilities capabilities() const

/*
Returns the capabilities of this proxy server.

This function was introduced in  Qt 4.5.

See also setCapabilities() and type().
*/
func (this *QNetworkProxy) Capabilities() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy12capabilitiesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCachingProxy() const

/*
Returns true if this proxy supports the QNetworkProxy::CachingCapability capability.

In Qt 4.4, the capability was tied to the proxy type, but since Qt 4.5 it is possible to remove the capability of caching from a proxy by calling setCapabilities().

This function was introduced in  Qt 4.4.

See also capabilities(), type(), and isTransparentProxy().
*/
func (this *QNetworkProxy) IsCachingProxy() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy14isCachingProxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTransparentProxy() const

/*
Returns true if this proxy supports transparent tunneling of TCP connections. This matches the QNetworkProxy::TunnelingCapability capability.

In Qt 4.4, the capability was tied to the proxy type, but since Qt 4.5 it is possible to remove the capability of caching from a proxy by calling setCapabilities().

This function was introduced in  Qt 4.4.

See also capabilities(), type(), and isCachingProxy().
*/
func (this *QNetworkProxy) IsTransparentProxy() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy18isTransparentProxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUser(const QString &)

/*
Sets the user name for proxy authentication to be user.

See also user(), setPassword(), and password().
*/
func (this *QNetworkProxy) SetUser(userName string) {
	var tmpArg0 = qtcore.NewQString5(userName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy7setUserERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QString user() const

/*
Returns the user name used for authentication.

See also setUser(), setPassword(), and password().
*/
func (this *QNetworkProxy) User() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy4userEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPassword(const QString &)

/*
Sets the password for proxy authentication to be password.

See also user(), setUser(), and password().
*/
func (this *QNetworkProxy) SetPassword(password string) {
	var tmpArg0 = qtcore.NewQString5(password)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy11setPasswordERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] QString password() const

/*
Returns the password used for authentication.

See also user(), setPassword(), and setUser().
*/
func (this *QNetworkProxy) Password() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy8passwordEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHostName(const QString &)

/*
Sets the host name of the proxy host to be hostName.

See also hostName(), setPort(), and port().
*/
func (this *QNetworkProxy) SetHostName(hostName string) {
	var tmpArg0 = qtcore.NewQString5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy11setHostNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:191
// index:0
// Public Visibility=Default Availability=Available
// [8] QString hostName() const

/*
Returns the host name of the proxy host.

See also setHostName(), setPort(), and port().
*/
func (this *QNetworkProxy) HostName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy8hostNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPort(quint16)

/*
Sets the port of the proxy host to be port.

See also hostName(), setHostName(), and port().
*/
func (this *QNetworkProxy) SetPort(port uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy7setPortEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:194
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 port() const

/*
Returns the port of the proxy host.

See also setHostName(), setPort(), and hostName().
*/
func (this *QNetworkProxy) Port() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy4portEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:196
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setApplicationProxy(const QNetworkProxy &)

/*
Sets the application level network proxying to be networkProxy.

If a QAbstractSocket or QTcpSocket has the QNetworkProxy::DefaultProxy type, then the QNetworkProxy set with this function is used. If you want more flexibility in determining which proxy is used, use the QNetworkProxyFactory class.

Setting a default proxy value with this function will override the application proxy factory set with QNetworkProxyFactory::setApplicationProxyFactory, and disable the use of a system proxy.

See also QNetworkProxyFactory, applicationProxy(), QAbstractSocket::setProxy(), and QTcpServer::setProxy().
*/
func (this *QNetworkProxy) SetApplicationProxy(proxy QNetworkProxy_ITF) {
	var convArg0 unsafe.Pointer
	if proxy != nil && proxy.QNetworkProxy_PTR() != nil {
		convArg0 = proxy.QNetworkProxy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy19setApplicationProxyERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QNetworkProxy_SetApplicationProxy(proxy QNetworkProxy_ITF) {
	var nilthis *QNetworkProxy
	nilthis.SetApplicationProxy(proxy)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:197
// index:0
// Public static Visibility=Default Availability=Available
// [8] QNetworkProxy applicationProxy()

/*
Returns the application level network proxying.

If a QAbstractSocket or QTcpSocket has the QNetworkProxy::DefaultProxy type, then the QNetworkProxy returned by this function is used.

See also QNetworkProxyFactory, setApplicationProxy(), QAbstractSocket::proxy(), and QTcpServer::proxy().
*/
func (this *QNetworkProxy) ApplicationProxy() *QNetworkProxy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy16applicationProxyEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkProxy)
	return rv2
}
func QNetworkProxy_ApplicationProxy() *QNetworkProxy /*123*/ {
	var nilthis *QNetworkProxy
	rv := nilthis.ApplicationProxy()
	return rv
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:200
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant header(QNetworkRequest::KnownHeaders) const

/*
Returns the value of the known network header header if it is in use for this proxy. If it is not present, returns QVariant() (i.e., an invalid variant).

This function was introduced in  Qt 5.0.

See also QNetworkRequest::KnownHeaders, rawHeader(), and setHeader().
*/
func (this *QNetworkProxy) Header(header int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy6headerEN15QNetworkRequest12KnownHeadersE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), header)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeader(QNetworkRequest::KnownHeaders, const QVariant &)

/*
Sets the value of the known header header to be value, overriding any previously set headers. This operation also sets the equivalent raw HTTP header.

If the proxy is not of type HttpProxy or HttpCachingProxy this has no effect.

This function was introduced in  Qt 5.0.

See also QNetworkRequest::KnownHeaders, setRawHeader(), and header().
*/
func (this *QNetworkProxy) SetHeader(header int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy9setHeaderEN15QNetworkRequest12KnownHeadersERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:204
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasRawHeader(const QByteArray &) const

/*
Returns true if the raw header headerName is in use for this proxy. Returns false if the proxy is not of type HttpProxy or HttpCachingProxy.

This function was introduced in  Qt 5.0.

See also rawHeader() and setRawHeader().
*/
func (this *QNetworkProxy) HasRawHeader(headerName qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy12hasRawHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:206
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray rawHeader(const QByteArray &) const

/*
Returns the raw form of header headerName. If no such header is present or the proxy is not of type HttpProxy or HttpCachingProxy, an empty QByteArray is returned, which may be indistinguishable from a header that is present but has no content (use hasRawHeader() to find out if the header exists or not).

Raw headers can be set with setRawHeader() or with setHeader().

This function was introduced in  Qt 5.0.

See also header() and setRawHeader().
*/
func (this *QNetworkProxy) RawHeader(headerName qtcore.QByteArray_ITF) *qtcore.QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy9rawHeaderERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawHeader(const QByteArray &, const QByteArray &)

/*
Sets the header headerName to be of value headerValue. If headerName corresponds to a known header (see QNetworkRequest::KnownHeaders), the raw format will be parsed and the corresponding "cooked" header will be set as well.

For example:


  request.setRawHeader(QByteArray("Last-Modified"), QByteArray("Sun, 06 Nov 1994 08:49:37 GMT"));



will also set the known header LastModifiedHeader to be the QDateTime object of the parsed date.

Note: Setting the same header twice overrides the previous setting. To accomplish the behaviour of multiple HTTP headers of the same name, you should concatenate the two values, separating them with a comma (",") and set one single raw header.

If the proxy is not of type HttpProxy or HttpCachingProxy this has no effect.

This function was introduced in  Qt 5.0.

See also QNetworkRequest::KnownHeaders, setHeader(), hasRawHeader(), and rawHeader().
*/
func (this *QNetworkProxy) SetRawHeader(headerName qtcore.QByteArray_ITF, value qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if headerName != nil && headerName.QByteArray_PTR() != nil {
		convArg0 = headerName.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QByteArray_PTR() != nil {
		convArg1 = value.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy12setRawHeaderERK10QByteArrayS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the types of network proxying provided in Qt.

There are two types of proxies that Qt understands: transparent proxies and caching proxies. The first group consists of proxies that can handle any arbitrary data transfer, while the second can only handle specific requests. The caching proxies only make sense for the specific classes where they can be used.



The table below lists different proxy types and their capabilities. Since each proxy type has different capabilities, it is important to understand them before choosing a proxy type.


 Proxy typeDescriptionDefault capabilities
SOCKS 5Generic proxy for any kind of connection. Supports TCP, UDP, binding to a port (incoming connections) and authentication.TunnelingCapability, ListeningCapability, UdpTunnelingCapability, HostNameLookupCapability
HTTPImplemented using the "CONNECT" command, supports only outgoing TCP connections; supports authentication.TunnelingCapability, CachingCapability, HostNameLookupCapability
Caching-only HTTPImplemented using normal HTTP commands, it is useful only in the context of HTTP requests (see QNetworkAccessManager)CachingCapability, HostNameLookupCapability
Caching FTPImplemented using an FTP proxy, it is useful only in the context of FTP requests (see QNetworkAccessManager)CachingCapability, HostNameLookupCapability


Also note that you shouldn't set the application default proxy (setApplicationProxy()) to a proxy that doesn't have the TunnelingCapability capability. If you do, QTcpSocket will not know how to open connections.

See also setType(), type(), capabilities(), and setCapabilities().

*/
type QNetworkProxy__ProxyType = int

// Proxy is determined based on the application proxy set using setApplicationProxy()
const QNetworkProxy__DefaultProxy QNetworkProxy__ProxyType = 0

//
const QNetworkProxy__Socks5Proxy QNetworkProxy__ProxyType = 1

// No proxying is used
const QNetworkProxy__NoProxy QNetworkProxy__ProxyType = 2

// HTTP transparent proxying is used
const QNetworkProxy__HttpProxy QNetworkProxy__ProxyType = 3

// Proxying for HTTP requests only
const QNetworkProxy__HttpCachingProxy QNetworkProxy__ProxyType = 4

// Proxying for FTP requests only
const QNetworkProxy__FtpCachingProxy QNetworkProxy__ProxyType = 5

func (this *QNetworkProxy) ProxyTypeItemName(val int) string {
	switch val {
	case QNetworkProxy__DefaultProxy: // 0
		return "DefaultProxy"
	case QNetworkProxy__Socks5Proxy: // 1
		return "Socks5Proxy"
	case QNetworkProxy__NoProxy: // 2
		return "NoProxy"
	case QNetworkProxy__HttpProxy: // 3
		return "HttpProxy"
	case QNetworkProxy__HttpCachingProxy: // 4
		return "HttpCachingProxy"
	case QNetworkProxy__FtpCachingProxy: // 5
		return "FtpCachingProxy"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkProxy_ProxyTypeItemName(val int) string {
	var nilthis *QNetworkProxy
	return nilthis.ProxyTypeItemName(val)
}

/*


 */
type QNetworkProxy__Capability = int

//
const QNetworkProxy__TunnelingCapability QNetworkProxy__Capability = 1

//
const QNetworkProxy__ListeningCapability QNetworkProxy__Capability = 2

//
const QNetworkProxy__UdpTunnelingCapability QNetworkProxy__Capability = 4

//
const QNetworkProxy__CachingCapability QNetworkProxy__Capability = 8

//
const QNetworkProxy__HostNameLookupCapability QNetworkProxy__Capability = 16

//
const QNetworkProxy__SctpTunnelingCapability QNetworkProxy__Capability = 32

//
const QNetworkProxy__SctpListeningCapability QNetworkProxy__Capability = 64

func (this *QNetworkProxy) CapabilityItemName(val int) string {
	switch val {
	case QNetworkProxy__TunnelingCapability: // 1
		return "TunnelingCapability"
	case QNetworkProxy__ListeningCapability: // 2
		return "ListeningCapability"
	case QNetworkProxy__UdpTunnelingCapability: // 4
		return "UdpTunnelingCapability"
	case QNetworkProxy__CachingCapability: // 8
		return "CachingCapability"
	case QNetworkProxy__HostNameLookupCapability: // 16
		return "HostNameLookupCapability"
	case QNetworkProxy__SctpTunnelingCapability: // 32
		return "SctpTunnelingCapability"
	case QNetworkProxy__SctpListeningCapability: // 64
		return "SctpListeningCapability"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkProxy_CapabilityItemName(val int) string {
	var nilthis *QNetworkProxy
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
