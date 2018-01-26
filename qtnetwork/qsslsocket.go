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
import "mkuse/cffiqt"
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QSslSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:82
// index:0
// Public
// void QSslSocket(class QObject *)
func NewQSslSocket(parent *qtcore.QObject /*777 QObject **/) *QSslSocket {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocketC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslSocketFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslsocket.h:83
// index:0
// Public virtual
// void ~QSslSocket()
func DeleteQSslSocket(*QSslSocket) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocketD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:84
// index:0
// Public virtual
// void resume()
func (this *QSslSocket) Resume() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket6resumeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:94
// index:0
// Public virtual
// void disconnectFromHost()
func (this *QSslSocket) DisconnectFromHost() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket18disconnectFromHostEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:96
// index:0
// Public virtual
// void setSocketOption(class QAbstractSocket::SocketOption, const class QVariant &)
func (this *QSslSocket) SetSocketOption(option int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket15setSocketOptionEN15QAbstractSocket12SocketOptionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:97
// index:0
// Public virtual
// QVariant socketOption(class QAbstractSocket::SocketOption)
func (this *QSslSocket) SocketOption(option int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket12socketOptionEN15QAbstractSocket12SocketOptionE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:99
// index:0
// Public
// QSslSocket::SslMode mode()
func (this *QSslSocket) Mode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket4modeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:100
// index:0
// Public
// bool isEncrypted()
func (this *QSslSocket) IsEncrypted() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket11isEncryptedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:102
// index:0
// Public
// QSsl::SslProtocol protocol()
func (this *QSslSocket) Protocol() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket8protocolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:103
// index:0
// Public
// void setProtocol(QSsl::SslProtocol)
func (this *QSslSocket) SetProtocol(protocol int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket11setProtocolEN4QSsl11SslProtocolE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), protocol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:105
// index:0
// Public
// QSslSocket::PeerVerifyMode peerVerifyMode()
func (this *QSslSocket) PeerVerifyMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket14peerVerifyModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:106
// index:0
// Public
// void setPeerVerifyMode(class QSslSocket::PeerVerifyMode)
func (this *QSslSocket) SetPeerVerifyMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket17setPeerVerifyModeENS_14PeerVerifyModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:108
// index:0
// Public
// int peerVerifyDepth()
func (this *QSslSocket) PeerVerifyDepth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket15peerVerifyDepthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtNetwork/qsslsocket.h:109
// index:0
// Public
// void setPeerVerifyDepth(int)
func (this *QSslSocket) SetPeerVerifyDepth(depth int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket18setPeerVerifyDepthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:111
// index:0
// Public
// QString peerVerifyName()
func (this *QSslSocket) PeerVerifyName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket14peerVerifyNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:112
// index:0
// Public
// void setPeerVerifyName(const class QString &)
func (this *QSslSocket) SetPeerVerifyName(hostName *qtcore.QString) {
	var convArg0 = hostName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket17setPeerVerifyNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:115
// index:0
// Public virtual
// qint64 bytesAvailable()
func (this *QSslSocket) BytesAvailable() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket14bytesAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:116
// index:0
// Public virtual
// qint64 bytesToWrite()
func (this *QSslSocket) BytesToWrite() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket12bytesToWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:117
// index:0
// Public virtual
// bool canReadLine()
func (this *QSslSocket) CanReadLine() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket11canReadLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:118
// index:0
// Public virtual
// void close()
func (this *QSslSocket) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:119
// index:0
// Public virtual
// bool atEnd()
func (this *QSslSocket) AtEnd() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:120
// index:0
// Public
// bool flush()
func (this *QSslSocket) Flush() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket5flushEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:121
// index:0
// Public
// void abort()
func (this *QSslSocket) Abort() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket5abortEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:124
// index:0
// Public virtual
// void setReadBufferSize(qint64)
func (this *QSslSocket) SetReadBufferSize(size int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket17setReadBufferSizeEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:127
// index:0
// Public
// qint64 encryptedBytesAvailable()
func (this *QSslSocket) EncryptedBytesAvailable() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket23encryptedBytesAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:128
// index:0
// Public
// qint64 encryptedBytesToWrite()
func (this *QSslSocket) EncryptedBytesToWrite() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket21encryptedBytesToWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:131
// index:0
// Public
// QSslConfiguration sslConfiguration()
func (this *QSslSocket) SslConfiguration() *QSslConfiguration /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket16sslConfigurationEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:132
// index:0
// Public
// void setSslConfiguration(const class QSslConfiguration &)
func (this *QSslSocket) SetSslConfiguration(config *QSslConfiguration) {
	var convArg0 = config.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket19setSslConfigurationERK17QSslConfiguration", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:138
// index:0
// Public
// void setLocalCertificate(const class QSslCertificate &)
func (this *QSslSocket) SetLocalCertificate(certificate *QSslCertificate) {
	var convArg0 = certificate.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket19setLocalCertificateERK15QSslCertificate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:139
// index:1
// Public
// void setLocalCertificate(const class QString &, QSsl::EncodingFormat)
func (this *QSslSocket) SetLocalCertificate_1(fileName *qtcore.QString, format int) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket19setLocalCertificateERK7QStringN4QSsl14EncodingFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:140
// index:0
// Public
// QSslCertificate localCertificate()
func (this *QSslSocket) LocalCertificate() *QSslCertificate /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket16localCertificateEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:141
// index:0
// Public
// QSslCertificate peerCertificate()
func (this *QSslSocket) PeerCertificate() *QSslCertificate /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket15peerCertificateEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:143
// index:0
// Public
// QSslCipher sessionCipher()
func (this *QSslSocket) SessionCipher() *QSslCipher /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket13sessionCipherEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:144
// index:0
// Public
// QSsl::SslProtocol sessionProtocol()
func (this *QSslSocket) SessionProtocol() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket15sessionProtocolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:147
// index:0
// Public
// void setPrivateKey(const class QSslKey &)
func (this *QSslSocket) SetPrivateKey(key *QSslKey) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QSslKey", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:148
// index:1
// Public
// void setPrivateKey(const class QString &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, const class QByteArray &)
func (this *QSslSocket) SetPrivateKey_1(fileName *qtcore.QString, algorithm int, format int, passPhrase *qtcore.QByteArray) {
	var convArg0 = fileName.GetCthis()
	var convArg3 = passPhrase.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket13setPrivateKeyERK7QStringN4QSsl12KeyAlgorithmENS3_14EncodingFormatERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, algorithm, format, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:151
// index:0
// Public
// QSslKey privateKey()
func (this *QSslSocket) PrivateKey() *QSslKey /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSslSocket10privateKeyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslsocket.h:157
// index:0
// Public
// void setCiphers(const class QString &)
func (this *QSslSocket) SetCiphers(ciphers *qtcore.QString) {
	var convArg0 = ciphers.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket10setCiphersERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:164
// index:0
// Public
// bool addCaCertificates(const class QString &, QSsl::EncodingFormat, class QRegExp::PatternSyntax)
func (this *QSslSocket) AddCaCertificates(path *qtcore.QString, format int, syntax int) bool {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket17addCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format, syntax)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:166
// index:0
// Public
// void addCaCertificate(const class QSslCertificate &)
func (this *QSslSocket) AddCaCertificate(certificate *QSslCertificate) {
	var convArg0 = certificate.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket16addCaCertificateERK15QSslCertificate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:172
// index:0
// Public static
// bool addDefaultCaCertificates(const class QString &, QSsl::EncodingFormat, class QRegExp::PatternSyntax)
func (this *QSslSocket) AddDefaultCaCertificates(path *qtcore.QString, format int, syntax int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket24addDefaultCaCertificatesERK7QStringN4QSsl14EncodingFormatEN7QRegExp13PatternSyntaxE", ffiqt.FFI_TYPE_POINTER, path, format, syntax)
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
// Public static
// void addDefaultCaCertificate(const class QSslCertificate &)
func (this *QSslSocket) AddDefaultCaCertificate(certificate *QSslCertificate) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket23addDefaultCaCertificateERK15QSslCertificate", ffiqt.FFI_TYPE_POINTER, certificate)
	gopp.ErrPrint(err, rv)
}
func QSslSocket_AddDefaultCaCertificate(certificate *QSslCertificate) {
	var nilthis *QSslSocket
	nilthis.AddDefaultCaCertificate(certificate)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:182
// index:0
// Public virtual
// bool waitForConnected(int)
func (this *QSslSocket) WaitForConnected(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket16waitForConnectedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:183
// index:0
// Public
// bool waitForEncrypted(int)
func (this *QSslSocket) WaitForEncrypted(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket16waitForEncryptedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:184
// index:0
// Public virtual
// bool waitForReadyRead(int)
func (this *QSslSocket) WaitForReadyRead(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket16waitForReadyReadEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:185
// index:0
// Public virtual
// bool waitForBytesWritten(int)
func (this *QSslSocket) WaitForBytesWritten(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket19waitForBytesWrittenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:186
// index:0
// Public virtual
// bool waitForDisconnected(int)
func (this *QSslSocket) WaitForDisconnected(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket19waitForDisconnectedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslsocket.h:190
// index:0
// Public static
// bool supportsSsl()
func (this *QSslSocket) SupportsSsl() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket11supportsSslEv", ffiqt.FFI_TYPE_POINTER)
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
// Public static
// long sslLibraryVersionNumber()
func (this *QSslSocket) SslLibraryVersionNumber() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket23sslLibraryVersionNumberEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QSslSocket_SslLibraryVersionNumber() int {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryVersionNumber()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:192
// index:0
// Public static
// QString sslLibraryVersionString()
func (this *QSslSocket) SslLibraryVersionString() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket23sslLibraryVersionStringEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSslSocket_SslLibraryVersionString() *qtcore.QString /*123*/ {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryVersionString()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:193
// index:0
// Public static
// long sslLibraryBuildVersionNumber()
func (this *QSslSocket) SslLibraryBuildVersionNumber() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket28sslLibraryBuildVersionNumberEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QSslSocket_SslLibraryBuildVersionNumber() int {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryBuildVersionNumber()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:194
// index:0
// Public static
// QString sslLibraryBuildVersionString()
func (this *QSslSocket) SslLibraryBuildVersionString() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket28sslLibraryBuildVersionStringEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSslSocket_SslLibraryBuildVersionString() *qtcore.QString /*123*/ {
	var nilthis *QSslSocket
	rv := nilthis.SslLibraryBuildVersionString()
	return rv
}

// /usr/include/qt/QtNetwork/qsslsocket.h:199
// index:0
// Public
// void startClientEncryption()
func (this *QSslSocket) StartClientEncryption() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket21startClientEncryptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:200
// index:0
// Public
// void startServerEncryption()
func (this *QSslSocket) StartServerEncryption() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket21startServerEncryptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:201
// index:0
// Public
// void ignoreSslErrors()
func (this *QSslSocket) IgnoreSslErrors() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket15ignoreSslErrorsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:204
// index:0
// Public
// void encrypted()
func (this *QSslSocket) Encrypted() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket9encryptedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:205
// index:0
// Public
// void peerVerifyError(const class QSslError &)
func (this *QSslSocket) PeerVerifyError(error *QSslError) {
	var convArg0 = error.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket15peerVerifyErrorERK9QSslError", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:207
// index:0
// Public
// void modeChanged(class QSslSocket::SslMode)
func (this *QSslSocket) ModeChanged(newMode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket11modeChangedENS_7SslModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:208
// index:0
// Public
// void encryptedBytesWritten(qint64)
func (this *QSslSocket) EncryptedBytesWritten(totalBytes int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket21encryptedBytesWrittenEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), totalBytes)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:209
// index:0
// Public
// void preSharedKeyAuthenticationRequired(class QSslPreSharedKeyAuthenticator *)
func (this *QSslSocket) PreSharedKeyAuthenticationRequired(authenticator *QSslPreSharedKeyAuthenticator /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 = authenticator.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket34preSharedKeyAuthenticationRequiredEP29QSslPreSharedKeyAuthenticator", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslsocket.h:212
// index:0
// Protected virtual
// qint64 readData(char *, qint64)
func (this *QSslSocket) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket8readDataEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslsocket.h:213
// index:0
// Protected virtual
// qint64 writeData(const char *, qint64)
func (this *QSslSocket) WriteData(data string, len int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSslSocket9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
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
