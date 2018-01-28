package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkproxy.h
// #include <qnetworkproxy.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QNetworkProxy struct {
	*qtrt.CObject
}

func (this *QNetworkProxy) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkProxy) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQNetworkProxyFromPointer(cthis unsafe.Pointer) *QNetworkProxy {
	return &QNetworkProxy{&qtrt.CObject{cthis}}
}
func (*QNetworkProxy) NewFromPointer(cthis unsafe.Pointer) *QNetworkProxy {
	return NewQNetworkProxyFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:160
// index:0
// Public
// void QNetworkProxy()
func NewQNetworkProxy() *QNetworkProxy {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxyC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkProxyFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:161
// index:1
// Public
// void QNetworkProxy(enum QNetworkProxy::ProxyType, const QString &, quint16, const QString &, const QString &)
func NewQNetworkProxy_1(type_ int, hostName *qtcore.QString, port uint16, user *qtcore.QString, password *qtcore.QString) *QNetworkProxy {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg1 = hostName.GetCthis()
	var convArg3 = user.GetCthis()
	var convArg4 = password.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxyC2ENS_9ProxyTypeERK7QStringtS3_S3_", ffiqt.FFI_TYPE_VOID, cthis, type_, convArg1, port, convArg3, convArg4)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkProxyFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:168
// index:0
// Public
// void ~QNetworkProxy()
func DeleteQNetworkProxy(*QNetworkProxy) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxyD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:170
// index:0
// Public inline
// void swap(QNetworkProxy &)
func (this *QNetworkProxy) Swap(other *QNetworkProxy) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxy4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:176
// index:0
// Public
// void setType(QNetworkProxy::ProxyType)
func (this *QNetworkProxy) SetType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxy7setTypeENS_9ProxyTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:177
// index:0
// Public
// QNetworkProxy::ProxyType type()
func (this *QNetworkProxy) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:180
// index:0
// Public
// QNetworkProxy::Capabilities capabilities()
func (this *QNetworkProxy) Capabilities() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy12capabilitiesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:181
// index:0
// Public
// bool isCachingProxy()
func (this *QNetworkProxy) IsCachingProxy() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy14isCachingProxyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:182
// index:0
// Public
// bool isTransparentProxy()
func (this *QNetworkProxy) IsTransparentProxy() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy18isTransparentProxyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:184
// index:0
// Public
// void setUser(const QString &)
func (this *QNetworkProxy) SetUser(userName *qtcore.QString) {
	var convArg0 = userName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxy7setUserERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:185
// index:0
// Public
// QString user()
func (this *QNetworkProxy) User() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy4userEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:187
// index:0
// Public
// void setPassword(const QString &)
func (this *QNetworkProxy) SetPassword(password *qtcore.QString) {
	var convArg0 = password.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxy11setPasswordERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:188
// index:0
// Public
// QString password()
func (this *QNetworkProxy) Password() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy8passwordEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:190
// index:0
// Public
// void setHostName(const QString &)
func (this *QNetworkProxy) SetHostName(hostName *qtcore.QString) {
	var convArg0 = hostName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxy11setHostNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:191
// index:0
// Public
// QString hostName()
func (this *QNetworkProxy) HostName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy8hostNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:193
// index:0
// Public
// void setPort(quint16)
func (this *QNetworkProxy) SetPort(port uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxy7setPortEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), port)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:194
// index:0
// Public
// quint16 port()
func (this *QNetworkProxy) Port() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy4portEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:196
// index:0
// Public static
// void setApplicationProxy(const QNetworkProxy &)
func (this *QNetworkProxy) SetApplicationProxy(proxy *QNetworkProxy) {
	var convArg0 = proxy.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxy19setApplicationProxyERKS_", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QNetworkProxy_SetApplicationProxy(proxy *QNetworkProxy) {
	var nilthis *QNetworkProxy
	nilthis.SetApplicationProxy(proxy)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:197
// index:0
// Public static
// QNetworkProxy applicationProxy()
func (this *QNetworkProxy) ApplicationProxy() *QNetworkProxy /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxy16applicationProxyEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QNetworkProxy_ApplicationProxy() *QNetworkProxy /*123*/ {
	var nilthis *QNetworkProxy
	rv := nilthis.ApplicationProxy()
	return rv
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:200
// index:0
// Public
// QVariant header(QNetworkRequest::KnownHeaders)
func (this *QNetworkProxy) Header(header int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy6headerEN15QNetworkRequest12KnownHeadersE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), header)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:201
// index:0
// Public
// void setHeader(QNetworkRequest::KnownHeaders, const QVariant &)
func (this *QNetworkProxy) SetHeader(header int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxy9setHeaderEN15QNetworkRequest12KnownHeadersERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), header, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:204
// index:0
// Public
// bool hasRawHeader(const QByteArray &)
func (this *QNetworkProxy) HasRawHeader(headerName *qtcore.QByteArray) bool {
	var convArg0 = headerName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy12hasRawHeaderERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:206
// index:0
// Public
// QByteArray rawHeader(const QByteArray &)
func (this *QNetworkProxy) RawHeader(headerName *qtcore.QByteArray) *qtcore.QByteArray /*123*/ {
	var convArg0 = headerName.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QNetworkProxy9rawHeaderERK10QByteArray", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:207
// index:0
// Public
// void setRawHeader(const QByteArray &, const QByteArray &)
func (this *QNetworkProxy) SetRawHeader(headerName *qtcore.QByteArray, value *qtcore.QByteArray) {
	var convArg0 = headerName.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QNetworkProxy12setRawHeaderERK10QByteArrayS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
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
