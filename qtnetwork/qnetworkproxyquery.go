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
// extern C begin: 15
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
type QNetworkProxyQuery struct {
	*qtrt.CObject
}
type QNetworkProxyQuery_ITF interface {
	QNetworkProxyQuery_PTR() *QNetworkProxyQuery
}

func (ptr *QNetworkProxyQuery) QNetworkProxyQuery_PTR() *QNetworkProxyQuery { return ptr }

func (this *QNetworkProxyQuery) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkProxyQuery) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQNetworkProxyQueryFromPointer(cthis unsafe.Pointer) *QNetworkProxyQuery {
	return &QNetworkProxyQuery{&qtrt.CObject{cthis}}
}
func (*QNetworkProxyQuery) NewFromPointer(cthis unsafe.Pointer) *QNetworkProxyQuery {
	return NewQNetworkProxyQueryFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxyQuery()

/*

 */
func (*QNetworkProxyQuery) NewForInherit() *QNetworkProxyQuery {
	return NewQNetworkProxyQuery()
}
func NewQNetworkProxyQuery() *QNetworkProxyQuery {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:73
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxyQuery(const QUrl &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_1(requestUrl qtcore.QUrl_ITF, queryType int) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_1(requestUrl, queryType)
}
func NewQNetworkProxyQuery_1(requestUrl qtcore.QUrl_ITF, queryType int) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if requestUrl != nil && requestUrl.QUrl_PTR() != nil {
		convArg0 = requestUrl.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK4QUrlNS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:73
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxyQuery(const QUrl &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_1_(requestUrl qtcore.QUrl_ITF) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_1_(requestUrl)
}
func NewQNetworkProxyQuery_1_(requestUrl qtcore.QUrl_ITF) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if requestUrl != nil && requestUrl.QUrl_PTR() != nil {
		convArg0 = requestUrl.QUrl_PTR().GetCthis()
	}
	// arg: 1, QNetworkProxyQuery::QueryType=Enum, QNetworkProxyQuery::QueryType=Enum, , Invalid
	queryType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK4QUrlNS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:74
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxyQuery(const QString &, int, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_2(hostname string, port int, protocolTag string, queryType int) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_2(hostname, port, protocolTag, queryType)
}
func NewQNetworkProxyQuery_2(hostname string, port int, protocolTag string, queryType int) *QNetworkProxyQuery {
	var tmpArg0 = qtcore.NewQString_5(hostname)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(protocolTag)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK7QStringiS2_NS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, port, convArg2, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:74
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxyQuery(const QString &, int, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_2_(hostname string, port int) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_2_(hostname, port)
}
func NewQNetworkProxyQuery_2_(hostname string, port int) *QNetworkProxyQuery {
	var tmpArg0 = qtcore.NewQString_5(hostname)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, QNetworkProxyQuery::QueryType=Enum, QNetworkProxyQuery::QueryType=Enum, , Invalid
	queryType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK7QStringiS2_NS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, port, convArg2, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:74
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxyQuery(const QString &, int, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_2_1(hostname string, port int, protocolTag string) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_2_1(hostname, port, protocolTag)
}
func NewQNetworkProxyQuery_2_1(hostname string, port int, protocolTag string) *QNetworkProxyQuery {
	var tmpArg0 = qtcore.NewQString_5(hostname)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(protocolTag)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QNetworkProxyQuery::QueryType=Enum, QNetworkProxyQuery::QueryType=Enum, , Invalid
	queryType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK7QStringiS2_NS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, port, convArg2, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:76
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxyQuery(quint16, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_3(bindPort uint16, protocolTag string, queryType int) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_3(bindPort, protocolTag, queryType)
}
func NewQNetworkProxyQuery_3(bindPort uint16, protocolTag string, queryType int) *QNetworkProxyQuery {
	var tmpArg1 = qtcore.NewQString_5(protocolTag)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2EtRK7QStringNS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, bindPort, convArg1, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:76
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxyQuery(quint16, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_3_(bindPort uint16) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_3_(bindPort)
}
func NewQNetworkProxyQuery_3_(bindPort uint16) *QNetworkProxyQuery {
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, QNetworkProxyQuery::QueryType=Enum, QNetworkProxyQuery::QueryType=Enum, , Invalid
	queryType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2EtRK7QStringNS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, bindPort, convArg1, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:76
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QNetworkProxyQuery(quint16, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_3_1(bindPort uint16, protocolTag string) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_3_1(bindPort, protocolTag)
}
func NewQNetworkProxyQuery_3_1(bindPort uint16, protocolTag string) *QNetworkProxyQuery {
	var tmpArg1 = qtcore.NewQString_5(protocolTag)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QNetworkProxyQuery::QueryType=Enum, QNetworkProxyQuery::QueryType=Enum, , Invalid
	queryType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2EtRK7QStringNS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, bindPort, convArg1, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:80
// index:4
// Public Visibility=Default Availability=Deprecated
// [-2] void QNetworkProxyQuery(const QNetworkConfiguration &, const QUrl &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_4(networkConfiguration QNetworkConfiguration_ITF, requestUrl qtcore.QUrl_ITF, queryType int) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_4(networkConfiguration, requestUrl, queryType)
}
func NewQNetworkProxyQuery_4(networkConfiguration QNetworkConfiguration_ITF, requestUrl qtcore.QUrl_ITF, queryType int) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if networkConfiguration != nil && networkConfiguration.QNetworkConfiguration_PTR() != nil {
		convArg0 = networkConfiguration.QNetworkConfiguration_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if requestUrl != nil && requestUrl.QUrl_PTR() != nil {
		convArg1 = requestUrl.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationRK4QUrlNS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:80
// index:4
// Public Visibility=Default Availability=Deprecated
// [-2] void QNetworkProxyQuery(const QNetworkConfiguration &, const QUrl &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_4_(networkConfiguration QNetworkConfiguration_ITF, requestUrl qtcore.QUrl_ITF) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_4_(networkConfiguration, requestUrl)
}
func NewQNetworkProxyQuery_4_(networkConfiguration QNetworkConfiguration_ITF, requestUrl qtcore.QUrl_ITF) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if networkConfiguration != nil && networkConfiguration.QNetworkConfiguration_PTR() != nil {
		convArg0 = networkConfiguration.QNetworkConfiguration_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if requestUrl != nil && requestUrl.QUrl_PTR() != nil {
		convArg1 = requestUrl.QUrl_PTR().GetCthis()
	}
	// arg: 2, QNetworkProxyQuery::QueryType=Enum, QNetworkProxyQuery::QueryType=Enum, , Invalid
	queryType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationRK4QUrlNS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:83
// index:5
// Public Visibility=Default Availability=Deprecated
// [-2] void QNetworkProxyQuery(const QNetworkConfiguration &, const QString &, int, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_5(networkConfiguration QNetworkConfiguration_ITF, hostname string, port int, protocolTag string, queryType int) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_5(networkConfiguration, hostname, port, protocolTag, queryType)
}
func NewQNetworkProxyQuery_5(networkConfiguration QNetworkConfiguration_ITF, hostname string, port int, protocolTag string, queryType int) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if networkConfiguration != nil && networkConfiguration.QNetworkConfiguration_PTR() != nil {
		convArg0 = networkConfiguration.QNetworkConfiguration_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(hostname)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(protocolTag)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationRK7QStringiS5_NS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, port, convArg3, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:83
// index:5
// Public Visibility=Default Availability=Deprecated
// [-2] void QNetworkProxyQuery(const QNetworkConfiguration &, const QString &, int, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_5_(networkConfiguration QNetworkConfiguration_ITF, hostname string, port int) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_5_(networkConfiguration, hostname, port)
}
func NewQNetworkProxyQuery_5_(networkConfiguration QNetworkConfiguration_ITF, hostname string, port int) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if networkConfiguration != nil && networkConfiguration.QNetworkConfiguration_PTR() != nil {
		convArg0 = networkConfiguration.QNetworkConfiguration_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(hostname)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QNetworkProxyQuery::QueryType=Enum, QNetworkProxyQuery::QueryType=Enum, , Invalid
	queryType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationRK7QStringiS5_NS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, port, convArg3, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:83
// index:5
// Public Visibility=Default Availability=Deprecated
// [-2] void QNetworkProxyQuery(const QNetworkConfiguration &, const QString &, int, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_5_1(networkConfiguration QNetworkConfiguration_ITF, hostname string, port int, protocolTag string) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_5_1(networkConfiguration, hostname, port, protocolTag)
}
func NewQNetworkProxyQuery_5_1(networkConfiguration QNetworkConfiguration_ITF, hostname string, port int, protocolTag string) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if networkConfiguration != nil && networkConfiguration.QNetworkConfiguration_PTR() != nil {
		convArg0 = networkConfiguration.QNetworkConfiguration_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(hostname)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(protocolTag)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 4, QNetworkProxyQuery::QueryType=Enum, QNetworkProxyQuery::QueryType=Enum, , Invalid
	queryType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationRK7QStringiS5_NS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, port, convArg3, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:87
// index:6
// Public Visibility=Default Availability=Deprecated
// [-2] void QNetworkProxyQuery(const QNetworkConfiguration &, quint16, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_6(networkConfiguration QNetworkConfiguration_ITF, bindPort uint16, protocolTag string, queryType int) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_6(networkConfiguration, bindPort, protocolTag, queryType)
}
func NewQNetworkProxyQuery_6(networkConfiguration QNetworkConfiguration_ITF, bindPort uint16, protocolTag string, queryType int) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if networkConfiguration != nil && networkConfiguration.QNetworkConfiguration_PTR() != nil {
		convArg0 = networkConfiguration.QNetworkConfiguration_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(protocolTag)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationtRK7QStringNS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, bindPort, convArg2, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:87
// index:6
// Public Visibility=Default Availability=Deprecated
// [-2] void QNetworkProxyQuery(const QNetworkConfiguration &, quint16, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_6_(networkConfiguration QNetworkConfiguration_ITF, bindPort uint16) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_6_(networkConfiguration, bindPort)
}
func NewQNetworkProxyQuery_6_(networkConfiguration QNetworkConfiguration_ITF, bindPort uint16) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if networkConfiguration != nil && networkConfiguration.QNetworkConfiguration_PTR() != nil {
		convArg0 = networkConfiguration.QNetworkConfiguration_PTR().GetCthis()
	}
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, QNetworkProxyQuery::QueryType=Enum, QNetworkProxyQuery::QueryType=Enum, , Invalid
	queryType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationtRK7QStringNS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, bindPort, convArg2, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:87
// index:6
// Public Visibility=Default Availability=Deprecated
// [-2] void QNetworkProxyQuery(const QNetworkConfiguration &, quint16, const QString &, QNetworkProxyQuery::QueryType)

/*

 */
func (*QNetworkProxyQuery) NewForInherit_6_1(networkConfiguration QNetworkConfiguration_ITF, bindPort uint16, protocolTag string) *QNetworkProxyQuery {
	return NewQNetworkProxyQuery_6_1(networkConfiguration, bindPort, protocolTag)
}
func NewQNetworkProxyQuery_6_1(networkConfiguration QNetworkConfiguration_ITF, bindPort uint16, protocolTag string) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if networkConfiguration != nil && networkConfiguration.QNetworkConfiguration_PTR() != nil {
		convArg0 = networkConfiguration.QNetworkConfiguration_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(protocolTag)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QNetworkProxyQuery::QueryType=Enum, QNetworkProxyQuery::QueryType=Enum, , Invalid
	queryType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryC2ERK21QNetworkConfigurationtRK7QStringNS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, convArg0, bindPort, convArg2, queryType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkProxyQuery)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QNetworkProxyQuery & operator=(QNetworkProxyQuery &&)

/*

 */
func (this *QNetworkProxyQuery) Operator_equal(other unsafe.Pointer /*333*/) *QNetworkProxyQuery {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkProxyQuery)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:95
// index:1
// Public Visibility=Default Availability=Available
// [8] QNetworkProxyQuery & operator=(const QNetworkProxyQuery &)

/*

 */
func (this *QNetworkProxyQuery) Operator_equal_1(other QNetworkProxyQuery_ITF) *QNetworkProxyQuery {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkProxyQuery_PTR() != nil {
		convArg0 = other.QNetworkProxyQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkProxyQueryFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkProxyQuery)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QNetworkProxyQuery()

/*

 */
func DeleteQNetworkProxyQuery(this *QNetworkProxyQuery) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQueryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkProxyQuery &)

/*
Swaps this network proxy instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QNetworkProxyQuery) Swap(other QNetworkProxyQuery_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkProxyQuery_PTR() != nil {
		convArg0 = other.QNetworkProxyQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQuery4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QNetworkProxyQuery &) const

/*

 */
func (this *QNetworkProxyQuery) Operator_equal_equal(other QNetworkProxyQuery_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkProxyQuery_PTR() != nil {
		convArg0 = other.QNetworkProxyQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QNetworkProxyQueryeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QNetworkProxyQuery &) const

/*

 */
func (this *QNetworkProxyQuery) Operator_not_equal(other QNetworkProxyQuery_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QNetworkProxyQuery_PTR() != nil {
		convArg0 = other.QNetworkProxyQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QNetworkProxyQueryneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] QNetworkProxyQuery::QueryType queryType() const

/*

 */
func (this *QNetworkProxyQuery) QueryType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery9queryTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQueryType(QNetworkProxyQuery::QueryType)

/*

 */
func (this *QNetworkProxyQuery) SetQueryType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQuery12setQueryTypeENS_9QueryTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int peerPort() const

/*

 */
func (this *QNetworkProxyQuery) PeerPort() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery8peerPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerPort(int)

/*

 */
func (this *QNetworkProxyQuery) SetPeerPort(port int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQuery11setPeerPortEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QString peerHostName() const

/*

 */
func (this *QNetworkProxyQuery) PeerHostName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery12peerHostNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerHostName(const QString &)

/*

 */
func (this *QNetworkProxyQuery) SetPeerHostName(hostname string) {
	var tmpArg0 = qtcore.NewQString_5(hostname)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQuery15setPeerHostNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] int localPort() const

/*

 */
func (this *QNetworkProxyQuery) LocalPort() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery9localPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocalPort(int)

/*

 */
func (this *QNetworkProxyQuery) SetLocalPort(port int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQuery12setLocalPortEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QString protocolTag() const

/*

 */
func (this *QNetworkProxyQuery) ProtocolTag() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery11protocolTagEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProtocolTag(const QString &)

/*

 */
func (this *QNetworkProxyQuery) SetProtocolTag(protocolTag string) {
	var tmpArg0 = qtcore.NewQString_5(protocolTag)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQuery14setProtocolTagERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*

 */
func (this *QNetworkProxyQuery) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QUrl &)

/*

 */
func (this *QNetworkProxyQuery) SetUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQuery6setUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:124
// index:0
// Public Visibility=Default Availability=Deprecated
// [8] QNetworkConfiguration networkConfiguration() const

/*

 */
func (this *QNetworkProxyQuery) NetworkConfiguration() *QNetworkConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QNetworkProxyQuery20networkConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:126
// index:0
// Public Visibility=Default Availability=Deprecated
// [-2] void setNetworkConfiguration(const QNetworkConfiguration &)

/*

 */
func (this *QNetworkProxyQuery) SetNetworkConfiguration(networkConfiguration QNetworkConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if networkConfiguration != nil && networkConfiguration.QNetworkConfiguration_PTR() != nil {
		convArg0 = networkConfiguration.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QNetworkProxyQuery23setNetworkConfigurationERK21QNetworkConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QNetworkProxyQuery__QueryType = int

//
const QNetworkProxyQuery__TcpSocket QNetworkProxyQuery__QueryType = 0

//
const QNetworkProxyQuery__UdpSocket QNetworkProxyQuery__QueryType = 1

//
const QNetworkProxyQuery__SctpSocket QNetworkProxyQuery__QueryType = 2

//
const QNetworkProxyQuery__TcpServer QNetworkProxyQuery__QueryType = 100

//
const QNetworkProxyQuery__UrlRequest QNetworkProxyQuery__QueryType = 101

//
const QNetworkProxyQuery__SctpServer QNetworkProxyQuery__QueryType = 102

func (this *QNetworkProxyQuery) QueryTypeItemName(val int) string {
	switch val {
	case QNetworkProxyQuery__TcpSocket: // 0
		return "TcpSocket"
	case QNetworkProxyQuery__UdpSocket: // 1
		return "UdpSocket"
	case QNetworkProxyQuery__SctpSocket: // 2
		return "SctpSocket"
	case QNetworkProxyQuery__TcpServer: // 100
		return "TcpServer"
	case QNetworkProxyQuery__UrlRequest: // 101
		return "UrlRequest"
	case QNetworkProxyQuery__SctpServer: // 102
		return "SctpServer"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QNetworkProxyQuery_QueryTypeItemName(val int) string {
	var nilthis *QNetworkProxyQuery
	return nilthis.QueryTypeItemName(val)
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
