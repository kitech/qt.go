package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkproxy.h
// #include <qnetworkproxy.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QNetworkProxyQuery struct {
	*qtrt.CObject
}

func (this *QNetworkProxyQuery) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkProxyQuery) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQNetworkProxyQueryFromPointer(cthis unsafe.Pointer) *QNetworkProxyQuery {
	return &QNetworkProxyQuery{&qtrt.CObject{cthis}}
}
func (*QNetworkProxyQuery) NewFromPointer(cthis unsafe.Pointer) *QNetworkProxyQuery {
	return NewQNetworkProxyQueryFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:72
// index:0
// Public
// void QNetworkProxyQuery()
func NewQNetworkProxyQuery() *QNetworkProxyQuery {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:73
// index:1
// Public
// void QNetworkProxyQuery(const QUrl &, QNetworkProxyQuery::QueryType)
func NewQNetworkProxyQuery_1(requestUrl *qtcore.QUrl, queryType int) *QNetworkProxyQuery {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = requestUrl.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK4QUrlNS_9QueryTypeE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, queryType)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:74
// index:2
// Public
// void QNetworkProxyQuery(const QString &, int, const QString &, QNetworkProxyQuery::QueryType)
func NewQNetworkProxyQuery_2(hostname *qtcore.QString, port int, protocolTag *qtcore.QString, queryType int) *QNetworkProxyQuery {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = hostname.GetCthis()
	var convArg2 = protocolTag.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK7QStringiS2_NS_9QueryTypeE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, port, convArg2, queryType)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:76
// index:3
// Public
// void QNetworkProxyQuery(quint16, const QString &, QNetworkProxyQuery::QueryType)
func NewQNetworkProxyQuery_3(bindPort uint16, protocolTag *qtcore.QString, queryType int) *QNetworkProxyQuery {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg1 = protocolTag.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2EtRK7QStringNS_9QueryTypeE", ffiqt.FFI_TYPE_VOID, cthis, bindPort, convArg1, queryType)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:80
// index:4
// Public
// void QNetworkProxyQuery(const QNetworkConfiguration &, const QUrl &, QNetworkProxyQuery::QueryType)
func NewQNetworkProxyQuery_4(networkConfiguration *QNetworkConfiguration, requestUrl *qtcore.QUrl, queryType int) *QNetworkProxyQuery {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = networkConfiguration.GetCthis()
	var convArg1 = requestUrl.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationRK4QUrlNS_9QueryTypeE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, queryType)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:83
// index:5
// Public
// void QNetworkProxyQuery(const QNetworkConfiguration &, const QString &, int, const QString &, QNetworkProxyQuery::QueryType)
func NewQNetworkProxyQuery_5(networkConfiguration *QNetworkConfiguration, hostname *qtcore.QString, port int, protocolTag *qtcore.QString, queryType int) *QNetworkProxyQuery {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = networkConfiguration.GetCthis()
	var convArg1 = hostname.GetCthis()
	var convArg3 = protocolTag.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationRK7QStringiS5_NS_9QueryTypeE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, port, convArg3, queryType)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:87
// index:6
// Public
// void QNetworkProxyQuery(const QNetworkConfiguration &, quint16, const QString &, QNetworkProxyQuery::QueryType)
func NewQNetworkProxyQuery_6(networkConfiguration *QNetworkConfiguration, bindPort uint16, protocolTag *qtcore.QString, queryType int) *QNetworkProxyQuery {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = networkConfiguration.GetCthis()
	var convArg2 = protocolTag.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationtRK7QStringNS_9QueryTypeE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, bindPort, convArg2, queryType)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:96
// index:0
// Public
// void ~QNetworkProxyQuery()
func DeleteQNetworkProxyQuery(*QNetworkProxyQuery) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQueryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:98
// index:0
// Public inline
// void swap(QNetworkProxyQuery &)
func (this *QNetworkProxyQuery) Swap(other *QNetworkProxyQuery) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQuery4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:104
// index:0
// Public
// QNetworkProxyQuery::QueryType queryType()
func (this *QNetworkProxyQuery) QueryType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery9queryTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:105
// index:0
// Public
// void setQueryType(QNetworkProxyQuery::QueryType)
func (this *QNetworkProxyQuery) SetQueryType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQuery12setQueryTypeENS_9QueryTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:107
// index:0
// Public
// int peerPort()
func (this *QNetworkProxyQuery) PeerPort() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery8peerPortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:108
// index:0
// Public
// void setPeerPort(int)
func (this *QNetworkProxyQuery) SetPeerPort(port int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQuery11setPeerPortEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), port)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:110
// index:0
// Public
// QString peerHostName()
func (this *QNetworkProxyQuery) PeerHostName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery12peerHostNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:111
// index:0
// Public
// void setPeerHostName(const QString &)
func (this *QNetworkProxyQuery) SetPeerHostName(hostname *qtcore.QString) {
	var convArg0 = hostname.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQuery15setPeerHostNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:113
// index:0
// Public
// int localPort()
func (this *QNetworkProxyQuery) LocalPort() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery9localPortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:114
// index:0
// Public
// void setLocalPort(int)
func (this *QNetworkProxyQuery) SetLocalPort(port int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQuery12setLocalPortEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), port)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:116
// index:0
// Public
// QString protocolTag()
func (this *QNetworkProxyQuery) ProtocolTag() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery11protocolTagEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:117
// index:0
// Public
// void setProtocolTag(const QString &)
func (this *QNetworkProxyQuery) SetProtocolTag(protocolTag *qtcore.QString) {
	var convArg0 = protocolTag.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQuery14setProtocolTagERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:119
// index:0
// Public
// QUrl url()
func (this *QNetworkProxyQuery) Url() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery3urlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:120
// index:0
// Public
// void setUrl(const QUrl &)
func (this *QNetworkProxyQuery) SetUrl(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQuery6setUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:124
// index:0
// Public
// QNetworkConfiguration networkConfiguration()
func (this *QNetworkProxyQuery) NetworkConfiguration() *QNetworkConfiguration /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery20networkConfigurationEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:126
// index:0
// Public
// void setNetworkConfiguration(const QNetworkConfiguration &)
func (this *QNetworkProxyQuery) SetNetworkConfiguration(networkConfiguration *QNetworkConfiguration) {
	var convArg0 = networkConfiguration.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QNetworkProxyQuery23setNetworkConfigurationERK21QNetworkConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QNetworkProxyQuery__QueryType = int

const QNetworkProxyQuery__TcpSocket QNetworkProxyQuery__QueryType = 0
const QNetworkProxyQuery__UdpSocket QNetworkProxyQuery__QueryType = 1
const QNetworkProxyQuery__SctpSocket QNetworkProxyQuery__QueryType = 2
const QNetworkProxyQuery__TcpServer QNetworkProxyQuery__QueryType = 100
const QNetworkProxyQuery__UrlRequest QNetworkProxyQuery__QueryType = 101
const QNetworkProxyQuery__SctpServer QNetworkProxyQuery__QueryType = 102

//  body block end
