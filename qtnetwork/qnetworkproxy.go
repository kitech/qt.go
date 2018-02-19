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
// [-2] void QNetworkProxy(enum QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)
func NewQNetworkProxy_1(type_ int, hostName string, port uint16, user string, password string) *QNetworkProxy {
	var tmpArg1 = qtcore.NewQString_5(hostName)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(user)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(password)
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
// [-2] void QNetworkProxy(enum QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)
func NewQNetworkProxy_1_(type_ int) *QNetworkProxy {
	// arg: 1, const QString &=LValueReference, QString=Record,
	var convArg1 = qtcore.NewQString()
	// arg: 2, quint16=Typedef, quint16=Typedef, unsigned short
	port := uint16(0)
	// arg: 3, const QString &=LValueReference, QString=Record,
	var convArg3 = qtcore.NewQString()
	// arg: 4, const QString &=LValueReference, QString=Record,
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
// [-2] void QNetworkProxy(enum QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)
func NewQNetworkProxy_1_1(type_ int, hostName string) *QNetworkProxy {
	var tmpArg1 = qtcore.NewQString_5(hostName)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, quint16=Typedef, quint16=Typedef, unsigned short
	port := uint16(0)
	// arg: 3, const QString &=LValueReference, QString=Record,
	var convArg3 = qtcore.NewQString()
	// arg: 4, const QString &=LValueReference, QString=Record,
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
// [-2] void QNetworkProxy(enum QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)
func NewQNetworkProxy_1_2(type_ int, hostName string, port uint16) *QNetworkProxy {
	var tmpArg1 = qtcore.NewQString_5(hostName)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record,
	var convArg3 = qtcore.NewQString()
	// arg: 4, const QString &=LValueReference, QString=Record,
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
// [-2] void QNetworkProxy(enum QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)
func NewQNetworkProxy_1_3(type_ int, hostName string, port uint16, user string) *QNetworkProxy {
	var tmpArg1 = qtcore.NewQString_5(hostName)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(user)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 4, const QString &=LValueReference, QString=Record,
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
func (this *QNetworkProxy) Operator_equal_1(other QNetworkProxy_ITF) *QNetworkProxy {
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
func (this *QNetworkProxy) SetType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy7setTypeENS_9ProxyTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:177
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkProxy::ProxyType type() const
func (this *QNetworkProxy) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCapabilities(QNetworkProxy::Capabilities)
func (this *QNetworkProxy) SetCapabilities(capab int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy15setCapabilitiesE6QFlagsINS_10CapabilityEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), capab)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:180
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkProxy::Capabilities capabilities() const
func (this *QNetworkProxy) Capabilities() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy12capabilitiesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCachingProxy() const
func (this *QNetworkProxy) IsCachingProxy() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy14isCachingProxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:182
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTransparentProxy() const
func (this *QNetworkProxy) IsTransparentProxy() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy18isTransparentProxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUser(const QString &)
func (this *QNetworkProxy) SetUser(userName string) {
	var tmpArg0 = qtcore.NewQString_5(userName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy7setUserERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QString user() const
func (this *QNetworkProxy) User() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy4userEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPassword(const QString &)
func (this *QNetworkProxy) SetPassword(password string) {
	var tmpArg0 = qtcore.NewQString_5(password)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy11setPasswordERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] QString password() const
func (this *QNetworkProxy) Password() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy8passwordEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHostName(const QString &)
func (this *QNetworkProxy) SetHostName(hostName string) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy11setHostNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:191
// index:0
// Public Visibility=Default Availability=Available
// [8] QString hostName() const
func (this *QNetworkProxy) HostName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy8hostNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPort(quint16)
func (this *QNetworkProxy) SetPort(port uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QNetworkProxy7setPortEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:194
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 port() const
func (this *QNetworkProxy) Port() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QNetworkProxy4portEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:196
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setApplicationProxy(const QNetworkProxy &)
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

type QNetworkProxy__ProxyType = int

const QNetworkProxy__DefaultProxy QNetworkProxy__ProxyType = 0
const QNetworkProxy__Socks5Proxy QNetworkProxy__ProxyType = 1
const QNetworkProxy__NoProxy QNetworkProxy__ProxyType = 2
const QNetworkProxy__HttpProxy QNetworkProxy__ProxyType = 3
const QNetworkProxy__HttpCachingProxy QNetworkProxy__ProxyType = 4
const QNetworkProxy__FtpCachingProxy QNetworkProxy__ProxyType = 5

type QNetworkProxy__Capability = int

const QNetworkProxy__TunnelingCapability QNetworkProxy__Capability = 1
const QNetworkProxy__ListeningCapability QNetworkProxy__Capability = 2
const QNetworkProxy__UdpTunnelingCapability QNetworkProxy__Capability = 4
const QNetworkProxy__CachingCapability QNetworkProxy__Capability = 8
const QNetworkProxy__HostNameLookupCapability QNetworkProxy__Capability = 16
const QNetworkProxy__SctpTunnelingCapability QNetworkProxy__Capability = 32
const QNetworkProxy__SctpListeningCapability QNetworkProxy__Capability = 64

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
