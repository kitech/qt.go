package qtnetwork

// /usr/include/qt/QtNetwork/qsslsocket.h
// #include <qsslsocket.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

// long long readData(char *, qint64)
func (this *QSslSocket) InheritReadData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readData", f)
}

// long long writeData(const char *, qint64)
func (this *QSslSocket) InheritWriteData(f func(data string, len int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

type QSslSocket struct {
	*QTcpSocket
}
type QSslSocket_ITF interface {
	QTcpSocket_ITF
	QSslSocket_PTR() *QSslSocket
}

func (ptr *QSslSocket) QSslSocket_PTR() *QSslSocket { return ptr }

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
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qsslsocket.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslSocket(QObject *)
func NewQSslSocket(parent qtcore.QObject_ITF /*777 QObject **/) *QSslSocket {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSslSocket")
	return gothis
}

// /usr/include/qt/QtNetwork/qsslsocket.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSslSocket()
func DeleteQSslSocket(this *QSslSocket) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocketD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void resume()
func (this *QSslSocket) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, QIODevice::OpenMode, enum QAbstractSocket::NetworkLayerProtocol)
func (this *QSslSocket) ConnectToHostEncrypted(hostName string, port uint16, mode int, protocol int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket22connectToHostEncryptedERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:88
// index:1
// Public Visibility=Default Availability=Available
// [-2] void connectToHostEncrypted(const QString &, quint16, const QString &, QIODevice::OpenMode, enum QAbstractSocket::NetworkLayerProtocol)
func (this *QSslSocket) ConnectToHostEncrypted_1(hostName string, port uint16, sslPeerName string, mode int, protocol int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(sslPeerName)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket22connectToHostEncryptedERK7QStringtS2_6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2, mode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setSocketDescriptor(qintptr, enum QAbstractSocket::SocketState, QIODevice::OpenMode)
func (this *QSslSocket) SetSocketDescriptor(socketDescriptor int64, state int, openMode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setSocketDescriptorExN15QAbstractSocket11SocketStateE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), socketDescriptor, state, openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void connectToHost(const QString &, quint16, QIODevice::OpenMode, enum QAbstractSocket::NetworkLayerProtocol)
func (this *QSslSocket) ConnectToHost(hostName string, port uint16, openMode int, protocol int) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13connectToHostERK7QStringt6QFlagsIN9QIODevice12OpenModeFlagEEN15QAbstractSocket20NetworkLayerProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, openMode, protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void disconnectFromHost()
func (this *QSslSocket) DisconnectFromHost() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket18disconnectFromHostEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSocketOption(QAbstractSocket::SocketOption, const QVariant &)
func (this *QSslSocket) SetSocketOption(option int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket15setSocketOptionEN15QAbstractSocket12SocketOptionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant socketOption(QAbstractSocket::SocketOption)
func (this *QSslSocket) SocketOption(option int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket12socketOptionEN15QAbstractSocket12SocketOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEncrypted()
func (this *QSslSocket) IsEncrypted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket11isEncryptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol protocol()
func (this *QSslSocket) Protocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket8protocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProtocol(QSsl::SslProtocol)
func (this *QSslSocket) SetProtocol(protocol int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket11setProtocolEN4QSsl11SslProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslSocket::PeerVerifyMode peerVerifyMode()
func (this *QSslSocket) PeerVerifyMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket14peerVerifyModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyMode(QSslSocket::PeerVerifyMode)
func (this *QSslSocket) SetPeerVerifyMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17setPeerVerifyModeENS_14PeerVerifyModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] int peerVerifyDepth()
func (this *QSslSocket) PeerVerifyDepth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket15peerVerifyDepthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslsocket.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyDepth(int)
func (this *QSslSocket) SetPeerVerifyDepth(depth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket18setPeerVerifyDepthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QString peerVerifyName()
func (this *QSslSocket) PeerVerifyName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket14peerVerifyNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslsocket.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyName(const QString &)
func (this *QSslSocket) SetPeerVerifyName(hostName string) {
	var tmpArg0 = qtcore.NewQString_5(hostName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17setPeerVerifyNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable()
func (this *QSslSocket) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:116
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite()
func (this *QSslSocket) BytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket12bytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine()
func (this *QSslSocket) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()
func (this *QSslSocket) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd()
func (this *QSslSocket) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool flush()
func (this *QSslSocket) Flush() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void abort()
func (this *QSslSocket) Abort() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket5abortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setReadBufferSize(qint64)
func (this *QSslSocket) SetReadBufferSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17setReadBufferSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 encryptedBytesAvailable()
func (this *QSslSocket) EncryptedBytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket23encryptedBytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 encryptedBytesToWrite()
func (this *QSslSocket) EncryptedBytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket21encryptedBytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslConfiguration sslConfiguration()
func (this *QSslSocket) SslConfiguration() *QSslConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket16sslConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSslConfiguration(const QSslConfiguration &)
func (this *QSslSocket) SetSslConfiguration(config QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if config != nil && config.QSslConfiguration_PTR() != nil {
		convArg0 = config.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setSslConfigurationERK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocalCertificate(const QSslCertificate &)
func (this *QSslSocket) SetLocalCertificate(certificate QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if certificate != nil && certificate.QSslCertificate_PTR() != nil {
		convArg0 = certificate.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setLocalCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:139
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setLocalCertificate(const QString &, QSsl::EncodingFormat)
func (this *QSslSocket) SetLocalCertificate_1(fileName string, format int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19setLocalCertificateERK7QStringN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:140
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCertificate localCertificate()
func (this *QSslSocket) LocalCertificate() *QSslCertificate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket16localCertificateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QSslKey &)
func (this *QSslSocket) SetPrivateKey(key QSslKey_ITF) {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSslKey_PTR() != nil {
		convArg0 = key.QSslKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QSslKey", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:148
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QString &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, const QByteArray &)
func (this *QSslSocket) SetPrivateKey_1(fileName string, algorithm int, format int, passPhrase qtcore.QByteArray_ITF) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg3 unsafe.Pointer
	if passPhrase != nil && passPhrase.QByteArray_PTR() != nil {
		convArg3 = passPhrase.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QStringN4QSsl12KeyAlgorithmENS3_14EncodingFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, algorithm, format, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:151
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslKey privateKey()
func (this *QSslSocket) PrivateKey() *QSslKey /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslSocket10privateKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCiphers(const QString &)
func (this *QSslSocket) SetCiphers(ciphers string) {
	var tmpArg0 = qtcore.NewQString_5(ciphers)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket10setCiphersERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addCaCertificates(const QString &, QSsl::EncodingFormat, QRegExp::PatternSyntax)
func (this *QSslSocket) AddCaCertificates(path string, format int, syntax int) bool {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket17addCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format, syntax)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addCaCertificate(const QSslCertificate &)
func (this *QSslSocket) AddCaCertificate(certificate QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if certificate != nil && certificate.QSslCertificate_PTR() != nil {
		convArg0 = certificate.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16addCaCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:172
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool addDefaultCaCertificates(const QString &, QSsl::EncodingFormat, QRegExp::PatternSyntax)
func (this *QSslSocket) AddDefaultCaCertificates(path string, format int, syntax int) bool {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket24addDefaultCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", qtrt.FFI_TYPE_POINTER, convArg0, format, syntax)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QSslSocket_AddDefaultCaCertificates(path string, format int, syntax int) bool {
	var nilthis *QSslSocket
	rv := nilthis.AddDefaultCaCertificates(path, format, syntax)
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:174
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addDefaultCaCertificate(const QSslCertificate &)
func (this *QSslSocket) AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if certificate != nil && certificate.QSslCertificate_PTR() != nil {
		convArg0 = certificate.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket23addDefaultCaCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QSslSocket_AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	var nilthis *QSslSocket
	nilthis.AddDefaultCaCertificate(certificate)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:182
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForConnected(int)
func (this *QSslSocket) WaitForConnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForConnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:183
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForEncrypted(int)
func (this *QSslSocket) WaitForEncrypted(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForEncryptedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:184
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)
func (this *QSslSocket) WaitForReadyRead(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:185
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)
func (this *QSslSocket) WaitForBytesWritten(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:186
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForDisconnected(int)
func (this *QSslSocket) WaitForDisconnected(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket19waitForDisconnectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:190
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool supportsSsl()
func (this *QSslSocket) SupportsSsl() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket11supportsSslEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
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
func (this *QSslSocket) SslLibraryVersionString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket23sslLibraryVersionStringEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QSslSocket_SslLibraryVersionString() string {
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
	qtrt.ErrPrint(err, rv)
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
func (this *QSslSocket) SslLibraryBuildVersionString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket28sslLibraryBuildVersionStringEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QSslSocket_SslLibraryBuildVersionString() string {
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
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startServerEncryption()
func (this *QSslSocket) StartServerEncryption() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket21startServerEncryptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ignoreSslErrors()
func (this *QSslSocket) IgnoreSslErrors() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket15ignoreSslErrorsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void encrypted()
func (this *QSslSocket) Encrypted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket9encryptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void peerVerifyError(const QSslError &)
func (this *QSslSocket) PeerVerifyError(error QSslError_ITF) {
	var convArg0 unsafe.Pointer
	if error != nil && error.QSslError_PTR() != nil {
		convArg0 = error.QSslError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket15peerVerifyErrorERK9QSslError", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void modeChanged(QSslSocket::SslMode)
func (this *QSslSocket) ModeChanged(newMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket11modeChangedENS_7SslModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void encryptedBytesWritten(qint64)
func (this *QSslSocket) EncryptedBytesWritten(totalBytes int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket21encryptedBytesWrittenEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), totalBytes)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:209
// index:0
// Public Visibility=Default Availability=Available
// [-2] void preSharedKeyAuthenticationRequired(QSslPreSharedKeyAuthenticator *)
func (this *QSslSocket) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 unsafe.Pointer
	if authenticator != nil && authenticator.QSslPreSharedKeyAuthenticator_PTR() != nil {
		convArg0 = authenticator.QSslPreSharedKeyAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket34preSharedKeyAuthenticationRequiredEP29QSslPreSharedKeyAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:212
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)
func (this *QSslSocket) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslSocket8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
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
