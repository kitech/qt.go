package qtnetwork

// /usr/include/qt/QtNetwork/qabstractsocket.h
// #include <qabstractsocket.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QAbstractSocket struct {
	*qtcore.QIODevice
}

func (this *QAbstractSocket) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QIODevice.GetCthis()
	}
}
func (this *QAbstractSocket) SetCthis(cthis unsafe.Pointer) {
	this.QIODevice = qtcore.NewQIODeviceFromPointer(cthis)
}
func NewQAbstractSocketFromPointer(cthis unsafe.Pointer) *QAbstractSocket {
	bcthis0 := qtcore.NewQIODeviceFromPointer(cthis)
	return &QAbstractSocket{bcthis0}
}
func (*QAbstractSocket) NewFromPointer(cthis unsafe.Pointer) *QAbstractSocket {
	return NewQAbstractSocketFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:62
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:139
// index:0
// Public
// void QAbstractSocket(QAbstractSocket::SocketType, QObject *)
func NewQAbstractSocket(socketType int, parent *qtcore.QObject /*777 QObject **/) *QAbstractSocket {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocketC2ENS_10SocketTypeEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, socketType, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractSocketFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:140
// index:0
// Public virtual
// void ~QAbstractSocket()
func DeleteQAbstractSocket(*QAbstractSocket) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocketD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:142
// index:0
// Public virtual
// void resume()
func (this *QAbstractSocket) Resume() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket6resumeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:143
// index:0
// Public
// QAbstractSocket::PauseModes pauseMode()
func (this *QAbstractSocket) PauseMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket9pauseModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:153
// index:0
// Public virtual
// void disconnectFromHost()
func (this *QAbstractSocket) DisconnectFromHost() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket18disconnectFromHostEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:155
// index:0
// Public
// bool isValid()
func (this *QAbstractSocket) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:157
// index:0
// Public virtual
// qint64 bytesAvailable()
func (this *QAbstractSocket) BytesAvailable() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket14bytesAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:158
// index:0
// Public virtual
// qint64 bytesToWrite()
func (this *QAbstractSocket) BytesToWrite() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket12bytesToWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:160
// index:0
// Public virtual
// bool canReadLine()
func (this *QAbstractSocket) CanReadLine() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket11canReadLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:162
// index:0
// Public
// quint16 localPort()
func (this *QAbstractSocket) LocalPort() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket9localPortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:163
// index:0
// Public
// QHostAddress localAddress()
func (this *QAbstractSocket) LocalAddress() *QHostAddress /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket12localAddressEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:164
// index:0
// Public
// quint16 peerPort()
func (this *QAbstractSocket) PeerPort() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket8peerPortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:165
// index:0
// Public
// QHostAddress peerAddress()
func (this *QAbstractSocket) PeerAddress() *QHostAddress /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket11peerAddressEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:166
// index:0
// Public
// QString peerName()
func (this *QAbstractSocket) PeerName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket8peerNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:168
// index:0
// Public
// qint64 readBufferSize()
func (this *QAbstractSocket) ReadBufferSize() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket14readBufferSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:169
// index:0
// Public virtual
// void setReadBufferSize(qint64)
func (this *QAbstractSocket) SetReadBufferSize(size int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket17setReadBufferSizeEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:171
// index:0
// Public
// void abort()
func (this *QAbstractSocket) Abort() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket5abortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:173
// index:0
// Public virtual
// qintptr socketDescriptor()
func (this *QAbstractSocket) SocketDescriptor() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket16socketDescriptorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:177
// index:0
// Public virtual
// void setSocketOption(QAbstractSocket::SocketOption, const QVariant &)
func (this *QAbstractSocket) SetSocketOption(option int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket15setSocketOptionENS_12SocketOptionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:178
// index:0
// Public virtual
// QVariant socketOption(QAbstractSocket::SocketOption)
func (this *QAbstractSocket) SocketOption(option int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket12socketOptionENS_12SocketOptionE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:180
// index:0
// Public
// QAbstractSocket::SocketType socketType()
func (this *QAbstractSocket) SocketType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket10socketTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:181
// index:0
// Public
// QAbstractSocket::SocketState state()
func (this *QAbstractSocket) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:182
// index:0
// Public
// QAbstractSocket::SocketError error()
func (this *QAbstractSocket) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:206
// index:1
// Public
// void error(QAbstractSocket::SocketError)
func (this *QAbstractSocket) Error_1(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket5errorENS_11SocketErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:185
// index:0
// Public virtual
// void close()
func (this *QAbstractSocket) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:186
// index:0
// Public virtual
// bool isSequential()
func (this *QAbstractSocket) IsSequential() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket12isSequentialEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:187
// index:0
// Public virtual
// bool atEnd()
func (this *QAbstractSocket) AtEnd() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:188
// index:0
// Public
// bool flush()
func (this *QAbstractSocket) Flush() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket5flushEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:191
// index:0
// Public virtual
// bool waitForConnected(int)
func (this *QAbstractSocket) WaitForConnected(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket16waitForConnectedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:192
// index:0
// Public virtual
// bool waitForReadyRead(int)
func (this *QAbstractSocket) WaitForReadyRead(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket16waitForReadyReadEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:193
// index:0
// Public virtual
// bool waitForBytesWritten(int)
func (this *QAbstractSocket) WaitForBytesWritten(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket19waitForBytesWrittenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:194
// index:0
// Public virtual
// bool waitForDisconnected(int)
func (this *QAbstractSocket) WaitForDisconnected(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket19waitForDisconnectedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:197
// index:0
// Public
// void setProxy(const QNetworkProxy &)
func (this *QAbstractSocket) SetProxy(networkProxy *QNetworkProxy) {
	var convArg0 = networkProxy.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket8setProxyERK13QNetworkProxy", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:198
// index:0
// Public
// QNetworkProxy proxy()
func (this *QAbstractSocket) Proxy() *QNetworkProxy /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QAbstractSocket5proxyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQNetworkProxyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:202
// index:0
// Public
// void hostFound()
func (this *QAbstractSocket) HostFound() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket9hostFoundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:203
// index:0
// Public
// void connected()
func (this *QAbstractSocket) Connected() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket9connectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:204
// index:0
// Public
// void disconnected()
func (this *QAbstractSocket) Disconnected() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket12disconnectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:205
// index:0
// Public
// void stateChanged(QAbstractSocket::SocketState)
func (this *QAbstractSocket) StateChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket12stateChangedENS_11SocketStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:208
// index:0
// Public
// void proxyAuthenticationRequired(const QNetworkProxy &, QAuthenticator *)
func (this *QAbstractSocket) ProxyAuthenticationRequired(proxy *QNetworkProxy, authenticator *QAuthenticator /*777 QAuthenticator **/) {
	var convArg0 = proxy.GetCthis()
	var convArg1 = authenticator.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket27proxyAuthenticationRequiredERK13QNetworkProxyP14QAuthenticator", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:212
// index:0
// Protected virtual
// qint64 readData(char *, qint64)
func (this *QAbstractSocket) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket8readDataEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:213
// index:0
// Protected virtual
// qint64 readLineData(char *, qint64)
func (this *QAbstractSocket) ReadLineData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket12readLineDataEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:214
// index:0
// Protected virtual
// qint64 writeData(const char *, qint64)
func (this *QAbstractSocket) WriteData(data string, len int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:216
// index:0
// Protected
// void setSocketState(QAbstractSocket::SocketState)
func (this *QAbstractSocket) SetSocketState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket14setSocketStateENS_11SocketStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:217
// index:0
// Protected
// void setSocketError(QAbstractSocket::SocketError)
func (this *QAbstractSocket) SetSocketError(socketError int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket14setSocketErrorENS_11SocketErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), socketError)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:218
// index:0
// Protected
// void setLocalPort(quint16)
func (this *QAbstractSocket) SetLocalPort(port uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket12setLocalPortEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), port)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:219
// index:0
// Protected
// void setLocalAddress(const QHostAddress &)
func (this *QAbstractSocket) SetLocalAddress(address *QHostAddress) {
	var convArg0 = address.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket15setLocalAddressERK12QHostAddress", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:220
// index:0
// Protected
// void setPeerPort(quint16)
func (this *QAbstractSocket) SetPeerPort(port uint16) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket11setPeerPortEt", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), port)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:221
// index:0
// Protected
// void setPeerAddress(const QHostAddress &)
func (this *QAbstractSocket) SetPeerAddress(address *QHostAddress) {
	var convArg0 = address.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket14setPeerAddressERK12QHostAddress", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:222
// index:0
// Protected
// void setPeerName(const QString &)
func (this *QAbstractSocket) SetPeerName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QAbstractSocket11setPeerNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QAbstractSocket__SocketType = int

const QAbstractSocket__TcpSocket QAbstractSocket__SocketType = 0
const QAbstractSocket__UdpSocket QAbstractSocket__SocketType = 1
const QAbstractSocket__SctpSocket QAbstractSocket__SocketType = 2
const QAbstractSocket__UnknownSocketType QAbstractSocket__SocketType = 4294967295

type QAbstractSocket__NetworkLayerProtocol = int

const QAbstractSocket__IPv4Protocol QAbstractSocket__NetworkLayerProtocol = 0
const QAbstractSocket__IPv6Protocol QAbstractSocket__NetworkLayerProtocol = 1
const QAbstractSocket__AnyIPProtocol QAbstractSocket__NetworkLayerProtocol = 2
const QAbstractSocket__UnknownNetworkLayerProtocol QAbstractSocket__NetworkLayerProtocol = 4294967295

type QAbstractSocket__SocketError = int

const QAbstractSocket__ConnectionRefusedError QAbstractSocket__SocketError = 0
const QAbstractSocket__RemoteHostClosedError QAbstractSocket__SocketError = 1
const QAbstractSocket__HostNotFoundError QAbstractSocket__SocketError = 2
const QAbstractSocket__SocketAccessError QAbstractSocket__SocketError = 3
const QAbstractSocket__SocketResourceError QAbstractSocket__SocketError = 4
const QAbstractSocket__SocketTimeoutError QAbstractSocket__SocketError = 5
const QAbstractSocket__DatagramTooLargeError QAbstractSocket__SocketError = 6
const QAbstractSocket__NetworkError QAbstractSocket__SocketError = 7
const QAbstractSocket__AddressInUseError QAbstractSocket__SocketError = 8
const QAbstractSocket__SocketAddressNotAvailableError QAbstractSocket__SocketError = 9
const QAbstractSocket__UnsupportedSocketOperationError QAbstractSocket__SocketError = 10
const QAbstractSocket__UnfinishedSocketOperationError QAbstractSocket__SocketError = 11
const QAbstractSocket__ProxyAuthenticationRequiredError QAbstractSocket__SocketError = 12
const QAbstractSocket__SslHandshakeFailedError QAbstractSocket__SocketError = 13
const QAbstractSocket__ProxyConnectionRefusedError QAbstractSocket__SocketError = 14
const QAbstractSocket__ProxyConnectionClosedError QAbstractSocket__SocketError = 15
const QAbstractSocket__ProxyConnectionTimeoutError QAbstractSocket__SocketError = 16
const QAbstractSocket__ProxyNotFoundError QAbstractSocket__SocketError = 17
const QAbstractSocket__ProxyProtocolError QAbstractSocket__SocketError = 18
const QAbstractSocket__OperationError QAbstractSocket__SocketError = 19
const QAbstractSocket__SslInternalError QAbstractSocket__SocketError = 20
const QAbstractSocket__SslInvalidUserDataError QAbstractSocket__SocketError = 21
const QAbstractSocket__TemporaryError QAbstractSocket__SocketError = 22
const QAbstractSocket__UnknownSocketError QAbstractSocket__SocketError = 4294967295

type QAbstractSocket__SocketState = int

const QAbstractSocket__UnconnectedState QAbstractSocket__SocketState = 0
const QAbstractSocket__HostLookupState QAbstractSocket__SocketState = 1
const QAbstractSocket__ConnectingState QAbstractSocket__SocketState = 2
const QAbstractSocket__ConnectedState QAbstractSocket__SocketState = 3
const QAbstractSocket__BoundState QAbstractSocket__SocketState = 4
const QAbstractSocket__ListeningState QAbstractSocket__SocketState = 5
const QAbstractSocket__ClosingState QAbstractSocket__SocketState = 6

type QAbstractSocket__SocketOption = int

const QAbstractSocket__LowDelayOption QAbstractSocket__SocketOption = 0
const QAbstractSocket__KeepAliveOption QAbstractSocket__SocketOption = 1
const QAbstractSocket__MulticastTtlOption QAbstractSocket__SocketOption = 2
const QAbstractSocket__MulticastLoopbackOption QAbstractSocket__SocketOption = 3
const QAbstractSocket__TypeOfServiceOption QAbstractSocket__SocketOption = 4
const QAbstractSocket__SendBufferSizeSocketOption QAbstractSocket__SocketOption = 5
const QAbstractSocket__ReceiveBufferSizeSocketOption QAbstractSocket__SocketOption = 6

type QAbstractSocket__BindFlag = int

const QAbstractSocket__DefaultForPlatform QAbstractSocket__BindFlag = 0
const QAbstractSocket__ShareAddress QAbstractSocket__BindFlag = 1
const QAbstractSocket__DontShareAddress QAbstractSocket__BindFlag = 2
const QAbstractSocket__ReuseAddressHint QAbstractSocket__BindFlag = 4

type QAbstractSocket__PauseMode = int

const QAbstractSocket__PauseNever QAbstractSocket__PauseMode = 0
const QAbstractSocket__PauseOnSslErrors QAbstractSocket__PauseMode = 1

//  body block end
