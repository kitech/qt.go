// +build !minimal

package qtnetwork

// /usr/include/qt/QtNetwork/qdtls.h
// #include <qdtls.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QDtls struct {
	*qtcore.QObject
}
type QDtls_ITF interface {
	qtcore.QObject_ITF
	QDtls_PTR() *QDtls
}

func (ptr *QDtls) QDtls_PTR() *QDtls { return ptr }

func (this *QDtls) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QDtls) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQDtlsFromPointer(cthis unsafe.Pointer) *QDtls {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QDtls{bcthis0}
}
func (*QDtls) NewFromPointer(cthis unsafe.Pointer) *QDtls {
	return NewQDtlsFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdtls.h:116
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDtls) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qdtls.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDtls(QSslSocket::SslMode, QObject *)

/*
Creates a QDtls object, parent is passed to the QObject constructor. mode is QSslSocket::SslServerMode for a server-side DTLS connection or QSslSocket::SslClientMode for a client.

See also sslMode() and QSslSocket::SslMode.
*/
func (*QDtls) NewForInherit(mode int, parent qtcore.QObject_ITF /*777 QObject **/) *QDtls {
	return NewQDtls(mode, parent)
}
func NewQDtls(mode int, parent qtcore.QObject_ITF /*777 QObject **/) *QDtls {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtlsC2EN10QSslSocket7SslModeEP7QObject", qtrt.FFI_TYPE_POINTER, mode, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDtlsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDtls")
	return gothis
}

// /usr/include/qt/QtNetwork/qdtls.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDtls(QSslSocket::SslMode, QObject *)

/*
Creates a QDtls object, parent is passed to the QObject constructor. mode is QSslSocket::SslServerMode for a server-side DTLS connection or QSslSocket::SslClientMode for a client.

See also sslMode() and QSslSocket::SslMode.
*/
func (*QDtls) NewForInheritp(mode int) *QDtls {
	return NewQDtlsp(mode)
}
func NewQDtlsp(mode int) *QDtls {
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtlsC2EN10QSslSocket7SslModeEP7QObject", qtrt.FFI_TYPE_POINTER, mode, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDtlsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDtls")
	return gothis
}

// /usr/include/qt/QtNetwork/qdtls.h:129
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDtls()

/*

 */
func DeleteQDtls(this *QDtls) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtlsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qdtls.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPeer(const QHostAddress &, quint16, const QString &)

/*
Sets the peer's address, port, and host name and returns true if successful. address must not be null, multicast, or broadcast. verificationName is the host name used for the certificate validation.

See also peerAddress(), peerPort(), and peerVerificationName().
*/
func (this *QDtls) SetPeer(address QHostAddress_ITF, port uint16, verificationName string) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(verificationName)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls7setPeerERK12QHostAddresstRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPeer(const QHostAddress &, quint16, const QString &)

/*
Sets the peer's address, port, and host name and returns true if successful. address must not be null, multicast, or broadcast. verificationName is the host name used for the certificate validation.

See also peerAddress(), peerPort(), and peerVerificationName().
*/
func (this *QDtls) SetPeerp(address QHostAddress_ITF, port uint16) bool {
	var convArg0 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg0 = address.QHostAddress_PTR().GetCthis()
	}
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls7setPeerERK12QHostAddresstRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, port, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPeerVerificationName(const QString &)

/*
Sets the host name that will be used for the certificate validation and returns true if successful.

Note: This function must be called before the handshake starts.

See also peerVerificationName() and setPeer().
*/
func (this *QDtls) SetPeerVerificationName(name string) bool {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls23setPeerVerificationNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostAddress peerAddress() const

/*
Returns the peer's address, set by setPeer(), or QHostAddress::Null.

See also setPeer().
*/
func (this *QDtls) PeerAddress() *QHostAddress /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls11peerAddressEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostAddressFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostAddress)
	return rv2
}

// /usr/include/qt/QtNetwork/qdtls.h:135
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 peerPort() const

/*
Returns the peer's port number, set by setPeer(), or 0.

See also setPeer().
*/
func (this *QDtls) PeerPort() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls8peerPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qdtls.h:136
// index:0
// Public Visibility=Default Availability=Available
// [8] QString peerVerificationName() const

/*
Returns the host name set by setPeer() or setPeerVerificationName(). The default value is an empty string.

See also setPeerVerificationName() and setPeer().
*/
func (this *QDtls) PeerVerificationName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls20peerVerificationNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdtls.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslSocket::SslMode sslMode() const

/*
Returns QSslSocket::SslServerMode for a server-side connection and QSslSocket::SslClientMode for a client.

See also QDtls() and QSslSocket::SslMode.
*/
func (this *QDtls) SslMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls7sslModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qdtls.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMtuHint(quint16)

/*
mtuHint is the maximum transmission unit (MTU), either discovered or guessed by the application. The application is not required to set this value.

See also mtuHint() and QAbstractSocket::PathMtuSocketOption.
*/
func (this *QDtls) SetMtuHint(mtuHint uint16) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls10setMtuHintEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mtuHint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdtls.h:140
// index:0
// Public Visibility=Default Availability=Available
// [2] quint16 mtuHint() const

/*
Returns the value previously set by setMtuHint(). The default value is 0.

See also setMtuHint().
*/
func (this *QDtls) MtuHint() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls7mtuHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtNetwork/qdtls.h:144
// index:0
// Public Visibility=Default Availability=Available
// [16] QDtls::GeneratorParameters cookieGeneratorParameters() const

/*
Returns the current hash algorithm and secret, either default ones or previously set by a call to setCookieGeneratorParameters().

The default hash algorithm is QCryptographicHash::Sha256 if Qt was configured to support it, QCryptographicHash::Sha1 otherwise. The default secret is obtained from the backend-specific cryptographically strong pseudorandom number generator.

See also setCookieGeneratorParameters(), QDtlsClientVerifier, and cookieGeneratorParameters().
*/
func (this *QDtls) CookieGeneratorParameters() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls25cookieGeneratorParametersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtNetwork/qdtls.h:146
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setDtlsConfiguration(const QSslConfiguration &)

/*
Sets the connection's TLS configuration from configuration and returns true if successful.

Note: This function must be called before the handshake starts.

See also dtlsConfiguration() and doHandshake().
*/
func (this *QDtls) SetDtlsConfiguration(configuration QSslConfiguration_ITF) bool {
	var convArg0 unsafe.Pointer
	if configuration != nil && configuration.QSslConfiguration_PTR() != nil {
		convArg0 = configuration.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls20setDtlsConfigurationERK17QSslConfiguration", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:147
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslConfiguration dtlsConfiguration() const

/*
Returns either the default DTLS configuration or the configuration set by an earlier call to setDtlsConfiguration().

See also setDtlsConfiguration() and QSslConfiguration::defaultDtlsConfiguration().
*/
func (this *QDtls) DtlsConfiguration() *QSslConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls17dtlsConfigurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qdtls.h:149
// index:0
// Public Visibility=Default Availability=Available
// [4] QDtls::HandshakeState handshakeState() const

/*
Returns the current handshake state for this QDtls.

See also doHandshake() and QDtls::HandshakeState.
*/
func (this *QDtls) HandshakeState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls14handshakeStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qdtls.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool doHandshake(QUdpSocket *, const QByteArray &)

/*
Starts or continues a DTLS handshake. socket must be a valid pointer. When starting a server-side DTLS handshake, dgram must contain the initial ClientHello message read from QUdpSocket. This function returns true if no error was found. Handshake state can be tested using handshakeState(). false return means some error occurred, use dtlsError() for more detailed information.

Note: If the identity of the peer can't be established, the error is set to QDtlsError::PeerVerificationError. If you want to ignore verification errors and continue connecting, you must call ignoreVerificationErrors() and then resumeHandshake(). If the errors cannot be ignored, you must call abortHandshake().


  if (!dtls.doHandshake(&socket, dgram)) {
      if (dtls.dtlsError() == QDtlsError::PeerVerificationError)
          dtls.abortAfterError(&socket);
  }



See also handshakeState(), dtlsError(), ignoreVerificationErrors(), resumeHandshake(), and abortHandshake().
*/
func (this *QDtls) DoHandshake(socket QUdpSocket_ITF /*777 QUdpSocket **/, dgram qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if socket != nil && socket.QUdpSocket_PTR() != nil {
		convArg0 = socket.QUdpSocket_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if dgram != nil && dgram.QByteArray_PTR() != nil {
		convArg1 = dgram.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls11doHandshakeEP10QUdpSocketRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool doHandshake(QUdpSocket *, const QByteArray &)

/*
Starts or continues a DTLS handshake. socket must be a valid pointer. When starting a server-side DTLS handshake, dgram must contain the initial ClientHello message read from QUdpSocket. This function returns true if no error was found. Handshake state can be tested using handshakeState(). false return means some error occurred, use dtlsError() for more detailed information.

Note: If the identity of the peer can't be established, the error is set to QDtlsError::PeerVerificationError. If you want to ignore verification errors and continue connecting, you must call ignoreVerificationErrors() and then resumeHandshake(). If the errors cannot be ignored, you must call abortHandshake().


  if (!dtls.doHandshake(&socket, dgram)) {
      if (dtls.dtlsError() == QDtlsError::PeerVerificationError)
          dtls.abortAfterError(&socket);
  }



See also handshakeState(), dtlsError(), ignoreVerificationErrors(), resumeHandshake(), and abortHandshake().
*/
func (this *QDtls) DoHandshakep(socket QUdpSocket_ITF /*777 QUdpSocket **/) bool {
	var convArg0 unsafe.Pointer
	if socket != nil && socket.QUdpSocket_PTR() != nil {
		convArg0 = socket.QUdpSocket_PTR().GetCthis()
	}
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls11doHandshakeEP10QUdpSocketRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:152
// index:0
// Public Visibility=Default Availability=Available
// [1] bool handleTimeout(QUdpSocket *)

/*
If a timeout occures during the handshake, the handshakeTimeout() signal is emitted. The application must call handleTimeout() to retransmit handshake messages; handleTimeout() returns true if a timeout has occurred, false otherwise. socket must be a valid pointer.

See also handshakeTimeout().
*/
func (this *QDtls) HandleTimeout(socket QUdpSocket_ITF /*777 QUdpSocket **/) bool {
	var convArg0 unsafe.Pointer
	if socket != nil && socket.QUdpSocket_PTR() != nil {
		convArg0 = socket.QUdpSocket_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls13handleTimeoutEP10QUdpSocket", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:153
// index:0
// Public Visibility=Default Availability=Available
// [1] bool resumeHandshake(QUdpSocket *)

/*
If peer verification errors were ignored during the handshake, resumeHandshake() resumes and completes the handshake and returns true. socket must be a valid pointer. Returns false if the handshake could not be resumed.

See also doHandshake(), abortHandshake(), peerVerificationErrors(), and ignoreVerificationErrors().
*/
func (this *QDtls) ResumeHandshake(socket QUdpSocket_ITF /*777 QUdpSocket **/) bool {
	var convArg0 unsafe.Pointer
	if socket != nil && socket.QUdpSocket_PTR() != nil {
		convArg0 = socket.QUdpSocket_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls15resumeHandshakeEP10QUdpSocket", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:154
// index:0
// Public Visibility=Default Availability=Available
// [1] bool abortHandshake(QUdpSocket *)

/*
Aborts the ongoing handshake. Returns true if one was on-going on socket; otherwise, sets a suitable error and returns false.

See also doHandshake() and resumeHandshake().
*/
func (this *QDtls) AbortHandshake(socket QUdpSocket_ITF /*777 QUdpSocket **/) bool {
	var convArg0 unsafe.Pointer
	if socket != nil && socket.QUdpSocket_PTR() != nil {
		convArg0 = socket.QUdpSocket_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls14abortHandshakeEP10QUdpSocket", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:155
// index:0
// Public Visibility=Default Availability=Available
// [1] bool shutdown(QUdpSocket *)

/*
Sends an encrypted shutdown alert message and closes the DTLS connection. Handshake state changes to QDtls::HandshakeNotStarted. socket must be a valid pointer. This function returns true on success.

See also doHandshake().
*/
func (this *QDtls) Shutdown(socket QUdpSocket_ITF /*777 QUdpSocket **/) bool {
	var convArg0 unsafe.Pointer
	if socket != nil && socket.QUdpSocket_PTR() != nil {
		convArg0 = socket.QUdpSocket_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls8shutdownEP10QUdpSocket", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:157
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isConnectionEncrypted() const

/*
Returns true if DTLS handshake completed successfully.

See also doHandshake() and handshakeState().
*/
func (this *QDtls) IsConnectionEncrypted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls21isConnectionEncryptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCipher sessionCipher() const

/*
Returns the cryptographic cipher used by this connection, or a null cipher if the connection isn't encrypted. The cipher for the session is selected during the handshake phase. The cipher is used to encrypt and decrypt data.

QSslConfiguration provides functions for setting the ordered list of ciphers from which the handshake phase will eventually select the session cipher. This ordered list must be in place before the handshake phase begins.

See also QSslConfiguration, setDtlsConfiguration(), and dtlsConfiguration().
*/
func (this *QDtls) SessionCipher() *QSslCipher /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls13sessionCipherEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCipher)
	return rv2
}

// /usr/include/qt/QtNetwork/qdtls.h:159
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol sessionProtocol() const

/*
Returns the DTLS protocol version used by this connection, or UnknownProtocol if the connection isn't encrypted yet. The protocol for the connection is selected during the handshake phase.

setDtlsConfiguration() can set the preferred version before the handshake starts.

See also setDtlsConfiguration(), QSslConfiguration, QSslConfiguration::defaultDtlsConfiguration(), and QSslConfiguration::setProtocol().
*/
func (this *QDtls) SessionProtocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls15sessionProtocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qdtls.h:161
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 writeDatagramEncrypted(QUdpSocket *, const QByteArray &)

/*
Encrypts dgram and writes the encrypted data into socket. Returns the number of bytes written, or -1 in case of error. The handshake must be completed before writing encrypted data. socket must be a valid pointer.

See also doHandshake(), handshakeState(), isConnectionEncrypted(), and dtlsError().
*/
func (this *QDtls) WriteDatagramEncrypted(socket QUdpSocket_ITF /*777 QUdpSocket **/, dgram qtcore.QByteArray_ITF) int64 {
	var convArg0 unsafe.Pointer
	if socket != nil && socket.QUdpSocket_PTR() != nil {
		convArg0 = socket.QUdpSocket_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if dgram != nil && dgram.QByteArray_PTR() != nil {
		convArg1 = dgram.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls22writeDatagramEncryptedEP10QUdpSocketRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qdtls.h:162
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray decryptDatagram(QUdpSocket *, const QByteArray &)

/*
Decrypts dgram and returns its contents as plain text. The handshake must be completed before datagrams can be decrypted. Depending on the type of the TLS message the connection may write into socket, which must be a valid pointer.
*/
func (this *QDtls) DecryptDatagram(socket QUdpSocket_ITF /*777 QUdpSocket **/, dgram qtcore.QByteArray_ITF) *qtcore.QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if socket != nil && socket.QUdpSocket_PTR() != nil {
		convArg0 = socket.QUdpSocket_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if dgram != nil && dgram.QByteArray_PTR() != nil {
		convArg1 = dgram.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls15decryptDatagramEP10QUdpSocketRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qdtls.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] QDtlsError dtlsError() const

/*
Returns the last error encountered by the connection or QDtlsError::NoError.

See also dtlsErrorString() and QDtlsError.
*/
func (this *QDtls) DtlsError() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls9dtlsErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qdtls.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dtlsErrorString() const

/*
Returns a textual description for the last error encountered by the connection or empty string.

See also dtlsError().
*/
func (this *QDtls) DtlsErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QDtls15dtlsErrorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qdtls.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pskRequired(QSslPreSharedKeyAuthenticator *)

/*
QDtls emits this signal when it negotiates a PSK ciphersuite, and therefore a PSK authentication is then required.

When using PSK, the client must send to the server a valid identity and a valid pre shared key, in order for the TLS handshake to continue. Applications can provide this information in a slot connected to this signal, by filling in the passed authenticator object according to their needs.

Note: Ignoring this signal, or failing to provide the required credentials, will cause the handshake to fail, and therefore the connection to be aborted.

Note: The authenticator object is owned by QDtls and must not be deleted by the application.

See also QSslPreSharedKeyAuthenticator.
*/
func (this *QDtls) PskRequired(authenticator QSslPreSharedKeyAuthenticator_ITF /*777 QSslPreSharedKeyAuthenticator **/) {
	var convArg0 unsafe.Pointer
	if authenticator != nil && authenticator.QSslPreSharedKeyAuthenticator_PTR() != nil {
		convArg0 = authenticator.QSslPreSharedKeyAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls11pskRequiredEP29QSslPreSharedKeyAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qdtls.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void handshakeTimeout()

/*
Packet loss can result in timeouts during the handshake phase. In this case QDtls emits a handshakeTimeout() signal. Call handleTimeout() to retransmit the handshake messages:


  DtlsClient::DtlsClient()
  {
      // Some initialization code here ...
      connect(&clientDtls, &QDtls::handshakeTimeout, this, &DtlsClient::handleTimeout);
  }

  void DtlsClient::handleTimeout()
  {
      clientDtls.handleTimeout(&clientSocket);
  }



See also handleTimeout().
*/
func (this *QDtls) HandshakeTimeout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QDtls16handshakeTimeoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*
Describes the current state of DTLS handshake.

This enum describes the current state of DTLS handshake for a QDtls connection.



This enum was introduced or modified in  Qt 5.12.

See also QDtls::doHandshake() and QDtls::handshakeState().

*/
type QDtls__HandshakeState = int

// Nothing done yet.
const QDtls__HandshakeNotStarted QDtls__HandshakeState = 0

// Handshake was initiated and no errors were found so far.
const QDtls__HandshakeInProgress QDtls__HandshakeState = 1

// The identity of the peer can't be established.
const QDtls__PeerVerificationFailed QDtls__HandshakeState = 2

// Handshake completed successfully and encrypted connection was established.
const QDtls__HandshakeComplete QDtls__HandshakeState = 3

func (this *QDtls) HandshakeStateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QDtls_HandshakeStateItemName(val int) string {
	var nilthis *QDtls
	return nilthis.HandshakeStateItemName(val)
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
