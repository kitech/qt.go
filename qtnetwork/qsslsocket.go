package qtnetwork

// /usr/include/qt/QtNetwork/qsslsocket.h
// #include <qsslsocket.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
// qint64 readData(char *, qint64)
func (this *QSslSocket) InheritReadData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readData", f)
}

// qint64 writeData(const char *, qint64)
func (this *QSslSocket) InheritWriteData(f func(data string, len int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

type QSslSocket struct {
	*QTcpSocket
}

func (this *QSslSocket) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTcpSocket.GetCthis()
	}
}
func (this *QSslSocket) SetCthis(cthis unsafe.Pointer) {
	this.QTcpSocket = NewQTcpSocketFromPointer(cthis)
}
func NewQSslSocketFromPointer(cthis unsafe.Pointer) *QSslSocket {
	bcthis0 := NewQTcpSocketFromPointer(cthis)
	return &QSslSocket{bcthis0}
}
func (*QSslSocket) NewFromPointer(cthis unsafe.Pointer) *QSslSocket {
	return NewQSslSocketFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSslSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslSocket(QObject *)
func NewQSslSocket(parent *qtcore.QObject /*777 QObject **/) *QSslSocket {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qsslsocket.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSslSocket()
func DeleteQSslSocket(this *QSslSocket) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocketD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void resume()
func (this *QSslSocket) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, QIODevice::OpenMode, enum QAbstractSocket::NetworkLayerProtocol)
func (this *QSslSocket) ConnectToHostEncrypted(hostName *qtcore.QString, port uint16, mode int, protocol int) {
	var convArg0 = hostName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket22connectToHostEncryptedERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:88
// index:1
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, const QString &, QIODevice::OpenMode, enum QAbstractSocket::NetworkLayerProtocol)
func (this *QSslSocket) ConnectToHostEncrypted_1(hostName *qtcore.QString, port uint16, sslPeerName *qtcore.QString, mode int, protocol int) {
	var convArg0 = hostName.GetCthis()
	var convArg2 = sslPeerName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket22connectToHostEncryptedERK7QStringtS2_6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2, mode, protocol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, enum QAbstractSocket::SocketState, QIODevice::OpenMode)
func (this *QSslSocket) SetSocketDescriptor(socketDescriptor int64, state int, openMode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setSocketDescriptorExN15QAbstractSocket11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, enum QAbstractSocket::NetworkLayerProtocol)
func (this *QSslSocket) ConnectToHost(hostName *qtcore.QString, port uint16, openMode int, protocol int) {
	var convArg0 = hostName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, openMode, protocol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void disconnectFromHost()
func (this *QSslSocket) DisconnectFromHost() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket18disconnectFromHostEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSocketOption(QAbstractSocket::SocketOption, const QVariant &)
func (this *QSslSocket) SetSocketOption(option int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket15setSocketOptionEN15QAbstractSocket12SocketOptionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant socketOption(QAbstractSocket::SocketOption)
func (this *QSslSocket) SocketOption(option int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket12socketOptionEN15QAbstractSocket12SocketOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslSocket::SslMode mode()
func (this *QSslSocket) Mode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket4modeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEncrypted()
func (this *QSslSocket) IsEncrypted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket11isEncryptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol protocol()
func (this *QSslSocket) Protocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket8protocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProtocol(QSsl::SslProtocol)
func (this *QSslSocket) SetProtocol(protocol int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket11setProtocolEN4QSsl11SslProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), protocol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslSocket::PeerVerifyMode peerVerifyMode()
func (this *QSslSocket) PeerVerifyMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket14peerVerifyModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyMode(QSslSocket::PeerVerifyMode)
func (this *QSslSocket) SetPeerVerifyMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17setPeerVerifyModeENS_14PeerVerifyModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] int peerVerifyDepth()
func (this *QSslSocket) PeerVerifyDepth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket15peerVerifyDepthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslsocket.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyDepth(int)
func (this *QSslSocket) SetPeerVerifyDepth(depth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket18setPeerVerifyDepthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QString peerVerifyName()
func (this *QSslSocket) PeerVerifyName() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket14peerVerifyNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyName(const QString &)
func (this *QSslSocket) SetPeerVerifyName(hostName *qtcore.QString) {
	var convArg0 = hostName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17setPeerVerifyNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable()
func (this *QSslSocket) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:116
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite()
func (this *QSslSocket) BytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket12bytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine()
func (this *QSslSocket) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()
func (this *QSslSocket) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd()
func (this *QSslSocket) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool flush()
func (this *QSslSocket) Flush() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void abort()
func (this *QSslSocket) Abort() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket5abortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setReadBufferSize(qint64)
func (this *QSslSocket) SetReadBufferSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17setReadBufferSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 encryptedBytesAvailable()
func (this *QSslSocket) EncryptedBytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket23encryptedBytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 encryptedBytesToWrite()
func (this *QSslSocket) EncryptedBytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket21encryptedBytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslConfiguration sslConfiguration()
func (this *QSslSocket) SslConfiguration() *QSslConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket16sslConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSslConfiguration(const QSslConfiguration &)
func (this *QSslSocket) SetSslConfiguration(config *QSslConfiguration) {
	var convArg0 = config.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setSslConfigurationERK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocalCertificate(const QSslCertificate &)
func (this *QSslSocket) SetLocalCertificate(certificate *QSslCertificate) {
	var convArg0 = certificate.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setLocalCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:139
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setLocalCertificate(const QString &, QSsl::EncodingFormat)
func (this *QSslSocket) SetLocalCertificate_1(fileName *qtcore.QString, format int) {
	var convArg0 = fileName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setLocalCertificateERK7QStringN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:140
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCertificate localCertificate()
func (this *QSslSocket) LocalCertificate() *QSslCertificate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket16localCertificateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCertificate peerCertificate()
func (this *QSslSocket) PeerCertificate() *QSslCertificate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket15peerCertificateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCipher sessionCipher()
func (this *QSslSocket) SessionCipher() *QSslCipher /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket13sessionCipherEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCipher)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:144
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol sessionProtocol()
func (this *QSslSocket) SessionProtocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket15sessionProtocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QSslKey &)
func (this *QSslSocket) SetPrivateKey(key *QSslKey) {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QSslKey", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:148
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QString &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, const QByteArray &)
func (this *QSslSocket) SetPrivateKey_1(fileName *qtcore.QString, algorithm int, format int, passPhrase *qtcore.QByteArray) {
	var convArg0 = fileName.GetCthis()
	var convArg3 = passPhrase.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QStringN4QSsl12KeyAlgorithmENS3_14EncodingFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, algorithm, format, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:151
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslKey privateKey()
func (this *QSslSocket) PrivateKey() *QSslKey /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket10privateKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCiphers(const QString &)
func (this *QSslSocket) SetCiphers(ciphers *qtcore.QString) {
	var convArg0 = ciphers.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket10setCiphersERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addCaCertificates(const QString &, QSsl::EncodingFormat, QRegExp::PatternSyntax)
func (this *QSslSocket) AddCaCertificates(path *qtcore.QString, format int, syntax int) bool {
	var convArg0 = path.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17addCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format, syntax)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addCaCertificate(const QSslCertificate &)
func (this *QSslSocket) AddCaCertificate(certificate *QSslCertificate) {
	var convArg0 = certificate.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16addCaCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:172
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool addDefaultCaCertificates(const QString &, QSsl::EncodingFormat, QRegExp::PatternSyntax)
func (this *QSslSocket) AddDefaultCaCertificates(path *qtcore.QString, format int, syntax int) bool {
	var convArg0 = path.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket24addDefaultCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, convArg0, format, syntax)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QSslSocket_AddDefaultCaCertificates(path *qtcore.QString, format int, syntax int) bool {
	var nilthis *QSslSocket
	rv := nilthis.AddDefaultCaCertificates(path, format, syntax)
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:174
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addDefaultCaCertificate(const QSslCertificate &)
func (this *QSslSocket) AddDefaultCaCertificate(certificate *QSslCertificate) {
	var convArg0 = certificate.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket23addDefaultCaCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QSslSocket_AddDefaultCaCertificate(certificate *QSslCertificate) {
	var nilthis *QSslSocket
	nilthis.AddDefaultCaCertificate(certificate)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:182
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForConnected(int)
func (this *QSslSocket) WaitForConnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:183
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForEncrypted(int)
func (this *QSslSocket) WaitForEncrypted(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForEncryptedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:184
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)
func (this *QSslSocket) WaitForReadyRead(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:185
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)
func (this *QSslSocket) WaitForBytesWritten(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:186
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)
func (this *QSslSocket) WaitForDisconnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:190
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool supportsSsl()
func (this *QSslSocket) SupportsSsl() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket11supportsSslEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QSslSocket_SupportsSsl() bool {
	var nilthis *QSslSocket
	rv := nilthis.SupportsSsl()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:191
// index:0
// Public static Visibility=Default Availability=Available
// [8] long sslLibraryVersionNumber()
func (this *QSslSocket) SslLibraryVersionNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket23sslLibraryVersionNumberEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QSslSocket_SslLibraryVersionNumber() int {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryVersionNumber()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:192
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString sslLibraryVersionString()
func (this *QSslSocket) SslLibraryVersionString() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket23sslLibraryVersionStringEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}
func QSslSocket_SslLibraryVersionString() *qtcore.QString /*123*/ {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryVersionString()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:193
// index:0
// Public static Visibility=Default Availability=Available
// [8] long sslLibraryBuildVersionNumber()
func (this *QSslSocket) SslLibraryBuildVersionNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket28sslLibraryBuildVersionNumberEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QSslSocket_SslLibraryBuildVersionNumber() int {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryBuildVersionNumber()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:194
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString sslLibraryBuildVersionString()
func (this *QSslSocket) SslLibraryBuildVersionString() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket28sslLibraryBuildVersionStringEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}
func QSslSocket_SslLibraryBuildVersionString() *qtcore.QString /*123*/ {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryBuildVersionString()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startClientEncryption()
func (this *QSslSocket) StartClientEncryption() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket21startClientEncryptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startServerEncryption()
func (this *QSslSocket) StartServerEncryption() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket21startServerEncryptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ignoreSslErrors()
func (this *QSslSocket) IgnoreSslErrors() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket15ignoreSslErrorsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void encrypted()
func (this *QSslSocket) Encrypted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket9encryptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void peerVerifyError(const QSslError &)
func (this *QSslSocket) PeerVerifyError(error *QSslError) {
	var convArg0 = error.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket15peerVerifyErrorERK9QSslError", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void modeChanged(QSslSocket::SslMode)
func (this *QSslSocket) ModeChanged(newMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket11modeChangedENS_7SslModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void encryptedBytesWritten(qint64)
func (this *QSslSocket) EncryptedBytesWritten(totalBytes int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket21encryptedBytesWrittenEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), totalBytes)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:209
// index:0
// Public Visibility=Default Availability=Available
// [-2] void preSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator *)
func (this *QSslSocket) PreSharedKeyAuthenticationRequired(authenticator *QSslPreSharedKeyAuthenticator /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 = authenticator.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket34preSharedKeyAuthenticationRequiredEP29QSslPreSharedKeyAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:212
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)
func (this *QSslSocket) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:213
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)
func (this *QSslSocket) WriteData(data string, len int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

type QSslSocket__SslMode = int

const QSslSocket__UnencryptedMode QSslSocket__SslMode = 0
const QSslSocket__SslClientMode QSslSocket__SslMode = 1
const QSslSocket__SslServerMode QSslSocket__SslMode = 2

type QSslSocket__PeerVerifyMode = int

const QSslSocket__VerifyNone QSslSocket__PeerVerifyMode = 0
const QSslSocket__QueryPeer QSslSocket__PeerVerifyMode = 1
const QSslSocket__VerifyPeer QSslSocket__PeerVerifyMode = 2
const QSslSocket__AutoVerifyPeer QSslSocket__PeerVerifyMode = 3

//  body block end
