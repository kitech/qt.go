package qtnetwork

// /usr/include/qt/QtNetwork/qsslconfiguration.h
// #include <qsslconfiguration.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 67
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
type QSslConfiguration struct {
	*qtrt.CObject
}
type QSslConfiguration_ITF interface {
	QSslConfiguration_PTR() *QSslConfiguration
}

func (ptr *QSslConfiguration) QSslConfiguration_PTR() *QSslConfiguration { return ptr }

func (this *QSslConfiguration) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslConfiguration) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslConfigurationFromPointer(cthis unsafe.Pointer) *QSslConfiguration {
	return &QSslConfiguration{&qtrt.CObject{cthis}}
}
func (*QSslConfiguration) NewFromPointer(cthis unsafe.Pointer) *QSslConfiguration {
	return NewQSslConfigurationFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslConfiguration()

/*
Constructs an empty SSL configuration. This configuration contains no valid settings and the state will be empty. isNull() will return true after this constructor is called.

Once any setter methods are called, isNull() will return false.
*/
func NewQSslConfiguration() *QSslConfiguration {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfigurationC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslConfiguration)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslConfiguration()

/*

 */
func DeleteQSslConfiguration(this *QSslConfiguration) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfigurationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSslConfiguration & operator=(QSslConfiguration &&)

/*

 */
func (this *QSslConfiguration) Operator_equal(other unsafe.Pointer /*333*/) *QSslConfiguration {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfigurationaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:85
// index:1
// Public Visibility=Default Availability=Available
// [8] QSslConfiguration & operator=(const QSslConfiguration &)

/*

 */
func (this *QSslConfiguration) Operator_equal_1(other QSslConfiguration_ITF) *QSslConfiguration {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslConfiguration_PTR() != nil {
		convArg0 = other.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfigurationaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:87
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslConfiguration &)

/*
Swaps this SSL configuration instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QSslConfiguration) Swap(other QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslConfiguration_PTR() != nil {
		convArg0 = other.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QSslConfiguration &) const

/*

 */
func (this *QSslConfiguration) Operator_equal_equal(other QSslConfiguration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslConfiguration_PTR() != nil {
		convArg0 = other.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfigurationeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QSslConfiguration &) const

/*

 */
func (this *QSslConfiguration) Operator_not_equal(other QSslConfiguration_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslConfiguration_PTR() != nil {
		convArg0 = other.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfigurationneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this is a null QSslConfiguration object.

A QSslConfiguration object is null if it has been default-constructed and no setter methods have been called.

See also setProtocol(), setLocalCertificate(), setPrivateKey(), setCiphers(), and setCaCertificates().
*/
func (this *QSslConfiguration) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol protocol() const

/*
Returns the protocol setting for this SSL configuration.

See also setProtocol().
*/
func (this *QSslConfiguration) Protocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration8protocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProtocol(QSsl::SslProtocol)

/*
Sets the protocol setting for this configuration to be protocol.

Setting the protocol once the connection has already been established has no effect.

See also protocol().
*/
func (this *QSslConfiguration) SetProtocol(protocol int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration11setProtocolEN4QSsl11SslProtocolE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), protocol)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslSocket::PeerVerifyMode peerVerifyMode() const

/*
Returns the verify mode. This mode decides whether QSslSocket should request a certificate from the peer (i.e., the client requests a certificate from the server, or a server requesting a certificate from the client), and whether it should require that this certificate is valid.

The default mode is AutoVerifyPeer, which tells QSslSocket to use VerifyPeer for clients, QueryPeer for servers.

See also setPeerVerifyMode().
*/
func (this *QSslConfiguration) PeerVerifyMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration14peerVerifyModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyMode(QSslSocket::PeerVerifyMode)

/*
Sets the verify mode to mode. This mode decides whether QSslSocket should request a certificate from the peer (i.e., the client requests a certificate from the server, or a server requesting a certificate from the client), and whether it should require that this certificate is valid.

The default mode is AutoVerifyPeer, which tells QSslSocket to use VerifyPeer for clients, QueryPeer for servers.

See also peerVerifyMode().
*/
func (this *QSslConfiguration) SetPeerVerifyMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration17setPeerVerifyModeEN10QSslSocket14PeerVerifyModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int peerVerifyDepth() const

/*
Returns the maximum number of certificates in the peer's certificate chain to be checked during the SSL handshake phase, or 0 (the default) if no maximum depth has been set, indicating that the whole certificate chain should be checked.

The certificates are checked in issuing order, starting with the peer's own certificate, then its issuer's certificate, and so on.

See also setPeerVerifyDepth() and peerVerifyMode().
*/
func (this *QSslConfiguration) PeerVerifyDepth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration15peerVerifyDepthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPeerVerifyDepth(int)

/*
Sets the maximum number of certificates in the peer's certificate chain to be checked during the SSL handshake phase, to depth. Setting a depth of 0 means that no maximum depth is set, indicating that the whole certificate chain should be checked.

The certificates are checked in issuing order, starting with the peer's own certificate, then its issuer's certificate, and so on.

See also peerVerifyDepth() and setPeerVerifyMode().
*/
func (this *QSslConfiguration) SetPeerVerifyDepth(depth int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration18setPeerVerifyDepthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCertificate localCertificate() const

/*
Returns the certificate to be presented to the peer during the SSL handshake process.

See also setLocalCertificate().
*/
func (this *QSslConfiguration) LocalCertificate() *QSslCertificate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration16localCertificateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocalCertificate(const QSslCertificate &)

/*
Sets the certificate to be presented to the peer during SSL handshake to be certificate.

Setting the certificate once the connection has been established has no effect.

A certificate is the means of identification used in the SSL process. The local certificate is used by the remote end to verify the local user's identity against its list of Certification Authorities. In most cases, such as in HTTP web browsing, only servers identify to the clients, so the client does not send a certificate.

See also localCertificate().
*/
func (this *QSslConfiguration) SetLocalCertificate(certificate QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if certificate != nil && certificate.QSslCertificate_PTR() != nil {
		convArg0 = certificate.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration19setLocalCertificateERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCertificate peerCertificate() const

/*
Returns the peer's digital certificate (i.e., the immediate certificate of the host you are connected to), or a null certificate, if the peer has not assigned a certificate.

The peer certificate is checked automatically during the handshake phase, so this function is normally used to fetch the certificate for display or for connection diagnostic purposes. It contains information about the peer, including its host name, the certificate issuer, and the peer's public key.

Because the peer certificate is set during the handshake phase, it is safe to access the peer certificate from a slot connected to the QSslSocket::sslErrors() signal, QNetworkReply::sslErrors() signal, or the QSslSocket::encrypted() signal.

If a null certificate is returned, it can mean the SSL handshake failed, or it can mean the host you are connected to doesn't have a certificate, or it can mean there is no connection.

If you want to check the peer's complete chain of certificates, use peerCertificateChain() to get them all at once.

See also peerCertificateChain(), QSslSocket::sslErrors(), QSslSocket::ignoreSslErrors(), QNetworkReply::sslErrors(), and QNetworkReply::ignoreSslErrors().
*/
func (this *QSslConfiguration) PeerCertificate() *QSslCertificate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration15peerCertificateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCipher sessionCipher() const

/*
Returns the socket's cryptographic cipher, or a null cipher if the connection isn't encrypted. The socket's cipher for the session is set during the handshake phase. The cipher is used to encrypt and decrypt data transmitted through the socket.

The SSL infrastructure also provides functions for setting the ordered list of ciphers from which the handshake phase will eventually select the session cipher. This ordered list must be in place before the handshake phase begins.

See also ciphers(), setCiphers(), and QSslSocket::supportedCiphers().
*/
func (this *QSslConfiguration) SessionCipher() *QSslCipher /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration13sessionCipherEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCipher)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol sessionProtocol() const

/*
Returns the socket's SSL/TLS protocol or UnknownProtocol if the connection isn't encrypted. The socket's protocol for the session is set during the handshake phase.

This function was introduced in  Qt 5.4.

See also protocol() and setProtocol().
*/
func (this *QSslConfiguration) SessionProtocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration15sessionProtocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslKey privateKey() const

/*
Returns the SSL key assigned to this connection or a null key if none has been assigned yet.

See also setPrivateKey() and localCertificate().
*/
func (this *QSslConfiguration) PrivateKey() *QSslKey /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration10privateKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrivateKey(const QSslKey &)

/*
Sets the connection's private key to key. The private key and the local certificate are used by clients and servers that must prove their identity to SSL peers.

Both the key and the local certificate are required if you are creating an SSL server socket. If you are creating an SSL client socket, the key and local certificate are required if your client must identify itself to an SSL server.

See also privateKey() and setLocalCertificate().
*/
func (this *QSslConfiguration) SetPrivateKey(key QSslKey_ITF) {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSslKey_PTR() != nil {
		convArg0 = key.QSslKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration13setPrivateKeyERK7QSslKey", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSslOption(QSsl::SslOption, _Bool)

/*
Enables or disables an SSL compatibility option. If on is true, the option is enabled. If on is false, the option is disabled.

See also testSslOption().
*/
func (this *QSslConfiguration) SetSslOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration12setSslOptionEN4QSsl9SslOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testSslOption(QSsl::SslOption) const

/*
Returns true if the specified SSL compatibility option is enabled.

This function was introduced in  Qt 4.8.

See also setSslOption().
*/
func (this *QSslConfiguration) TestSslOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration13testSslOptionEN4QSsl9SslOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray sessionTicket() const

/*
If QSsl::SslOptionDisableSessionPersistence was turned off, this function returns the session ticket used in the SSL handshake in ASN.1 format, suitable to e.g. be persisted to disk. If no session ticket was used or QSsl::SslOptionDisableSessionPersistence was not turned off, this function returns an empty QByteArray.

Note: When persisting the session ticket to disk or similar, be careful not to expose the session to a potential attacker, as knowledge of the session allows for eavesdropping on data encrypted with the session parameters.

This function was introduced in  Qt 5.2.

See also setSessionTicket(), QSsl::SslOptionDisableSessionPersistence, and setSslOption().
*/
func (this *QSslConfiguration) SessionTicket() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration13sessionTicketEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSessionTicket(const QByteArray &)

/*
Sets the session ticket to be used in an SSL handshake. QSsl::SslOptionDisableSessionPersistence must be turned off for this to work, and sessionTicket must be in ASN.1 format as returned by sessionTicket().

This function was introduced in  Qt 5.2.

See also sessionTicket(), QSsl::SslOptionDisableSessionPersistence, and setSslOption().
*/
func (this *QSslConfiguration) SetSessionTicket(sessionTicket qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if sessionTicket != nil && sessionTicket.QByteArray_PTR() != nil {
		convArg0 = sessionTicket.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration16setSessionTicketERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int sessionTicketLifeTimeHint() const

/*
If QSsl::SslOptionDisableSessionPersistence was turned off, this function returns the session ticket life time hint sent by the server (which might be 0). If the server did not send a session ticket (e.g. when resuming a session or when the server does not support it) or QSsl::SslOptionDisableSessionPersistence was not turned off, this function returns -1.

This function was introduced in  Qt 5.2.

See also sessionTicket(), QSsl::SslOptionDisableSessionPersistence, and setSslOption().
*/
func (this *QSslConfiguration) SessionTicketLifeTimeHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration25sessionTicketLifeTimeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslKey ephemeralServerKey() const

/*
Returns the ephemeral server key used for cipher algorithms with forward secrecy, e.g. DHE-RSA-AES128-SHA.

The ephemeral key is only available when running in client mode, i.e. QSslSocket::SslClientMode. When running in server mode or using a cipher algorithm without forward secrecy a null key is returned. The ephemeral server key will be set before emitting the encrypted() signal.

This function was introduced in  Qt 5.7.
*/
func (this *QSslConfiguration) EphemeralServerKey() *QSslKey /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration18ephemeralServerKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:146
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray preSharedKeyIdentityHint() const

/*
Returns the identity hint.

This function was introduced in  Qt 5.8.

See also setPreSharedKeyIdentityHint().
*/
func (this *QSslConfiguration) PreSharedKeyIdentityHint() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration24preSharedKeyIdentityHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreSharedKeyIdentityHint(const QByteArray &)

/*
Sets the identity hint for a preshared key authentication to hint. This will affect the next initiated handshake; calling this function on an already-encrypted socket will not affect the socket's identity hint.

The identity hint is used in QSslSocket::SslServerMode only!

This function was introduced in  Qt 5.8.

See also preSharedKeyIdentityHint().
*/
func (this *QSslConfiguration) SetPreSharedKeyIdentityHint(hint qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if hint != nil && hint.QByteArray_PTR() != nil {
		convArg0 = hint.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration27setPreSharedKeyIdentityHintERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:149
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslDiffieHellmanParameters diffieHellmanParameters() const

/*
Retrieves the current set of Diffie-Hellman parameters.

If no Diffie-Hellman parameters have been set, the QSslConfiguration object defaults to using the 1024-bit MODP group from RFC 2409.

This function was introduced in  Qt 5.8.

See also setDiffieHellmanParameters().
*/
func (this *QSslConfiguration) DiffieHellmanParameters() *QSslDiffieHellmanParameters /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration23diffieHellmanParametersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslDiffieHellmanParameters)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDiffieHellmanParameters(const QSslDiffieHellmanParameters &)

/*
Sets a custom set of Diffie-Hellman parameters to be used by this socket when functioning as a server to dhparams.

If no Diffie-Hellman parameters have been set, the QSslConfiguration object defaults to using the 1024-bit MODP group from RFC 2409.

This function was introduced in  Qt 5.8.

See also diffieHellmanParameters().
*/
func (this *QSslConfiguration) SetDiffieHellmanParameters(dhparams QSslDiffieHellmanParameters_ITF) {
	var convArg0 unsafe.Pointer
	if dhparams != nil && dhparams.QSslDiffieHellmanParameters_PTR() != nil {
		convArg0 = dhparams.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration26setDiffieHellmanParametersERK27QSslDiffieHellmanParameters", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:152
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSslConfiguration defaultConfiguration()

/*
Returns the default SSL configuration to be used in new SSL connections.

The default SSL configuration consists of:


no local certificate and no private key
protocol SecureProtocols (meaning either TLS 1.0 or SSL 3 will be used)
the system's default CA certificate list
the cipher list equal to the list of the SSL libraries' supported SSL ciphers that are 128 bits or more


See also QSslSocket::supportedCiphers() and setDefaultConfiguration().
*/
func (this *QSslConfiguration) DefaultConfiguration() *QSslConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration20defaultConfigurationEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}
func QSslConfiguration_DefaultConfiguration() *QSslConfiguration /*123*/ {
	var nilthis *QSslConfiguration
	rv := nilthis.DefaultConfiguration()
	return rv
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:153
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultConfiguration(const QSslConfiguration &)

/*
Sets the default SSL configuration to be used in new SSL connections to be configuration. Existing connections are not affected by this call.

See also QSslSocket::supportedCiphers() and defaultConfiguration().
*/
func (this *QSslConfiguration) SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if configuration != nil && configuration.QSslConfiguration_PTR() != nil {
		convArg0 = configuration.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration23setDefaultConfigurationERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QSslConfiguration_SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	var nilthis *QSslConfiguration
	nilthis.SetDefaultConfiguration(configuration)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray nextNegotiatedProtocol() const

/*
This function returns the protocol negotiated with the server if the Next Protocol Negotiation (NPN) or Application-Layer Protocol Negotiation (ALPN) TLS extension was enabled. In order for the NPN/ALPN extension to be enabled, setAllowedNextProtocols() needs to be called explicitly before connecting to the server.

If no protocol could be negotiated or the extension was not enabled, this function returns a QByteArray which is null.

This function was introduced in  Qt 5.3.

See also setAllowedNextProtocols() and nextProtocolNegotiationStatus().
*/
func (this *QSslConfiguration) NextNegotiatedProtocol() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration22nextNegotiatedProtocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:169
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslConfiguration::NextProtocolNegotiationStatus nextProtocolNegotiationStatus() const

/*
This function returns the status of the Next Protocol Negotiation (NPN) or Application-Layer Protocol Negotiation (ALPN). If the feature has not been enabled through setAllowedNextProtocols(), this function returns NextProtocolNegotiationNone. The status will be set before emitting the encrypted() signal.

This function was introduced in  Qt 5.3.

See also setAllowedNextProtocols(), allowedNextProtocols(), nextNegotiatedProtocol(), and QSslConfiguration::NextProtocolNegotiationStatus.
*/
func (this *QSslConfiguration) NextProtocolNegotiationStatus() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration29nextProtocolNegotiationStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*
Describes the status of the Next Protocol Negotiation (NPN) or Application-Layer Protocol Negotiation (ALPN).


*/
type QSslConfiguration__NextProtocolNegotiationStatus = int

// No application protocol has been negotiated (yet).
const QSslConfiguration__NextProtocolNegotiationNone QSslConfiguration__NextProtocolNegotiationStatus = 0

// A next protocol has been negotiated (see nextNegotiatedProtocol()).
const QSslConfiguration__NextProtocolNegotiationNegotiated QSslConfiguration__NextProtocolNegotiationStatus = 1

// The client and server could not agree on a common next application protocol.
const QSslConfiguration__NextProtocolNegotiationUnsupported QSslConfiguration__NextProtocolNegotiationStatus = 2

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
