package qtnetwork

// /usr/include/qt/QtNetwork/qsslconfiguration.h
// #include <qsslconfiguration.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 63
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
type QSslConfiguration struct {
	*qtrt.CObject
}

func (this *QSslConfiguration) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslConfiguration) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSslConfigurationFromPointer(cthis unsafe.Pointer) *QSslConfiguration {
	return &QSslConfiguration{&qtrt.CObject{cthis}}
}
func (*QSslConfiguration) NewFromPointer(cthis unsafe.Pointer) *QSslConfiguration {
	return NewQSslConfigurationFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:79
// index:0
// Public
// void QSslConfiguration()
func NewQSslConfiguration() *QSslConfiguration {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfigurationC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslConfigurationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:81
// index:0
// Public
// void ~QSslConfiguration()
func DeleteQSslConfiguration(*QSslConfiguration) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfigurationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:87
// index:0
// Public inline
// void swap(QSslConfiguration &)
func (this *QSslConfiguration) Swap(other *QSslConfiguration) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:94
// index:0
// Public
// bool isNull()
func (this *QSslConfiguration) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:96
// index:0
// Public
// QSsl::SslProtocol protocol()
func (this *QSslConfiguration) Protocol() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration8protocolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:97
// index:0
// Public
// void setProtocol(QSsl::SslProtocol)
func (this *QSslConfiguration) SetProtocol(protocol int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration11setProtocolEN4QSsl11SslProtocolE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), protocol)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:100
// index:0
// Public
// QSslSocket::PeerVerifyMode peerVerifyMode()
func (this *QSslConfiguration) PeerVerifyMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration14peerVerifyModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:101
// index:0
// Public
// void setPeerVerifyMode(QSslSocket::PeerVerifyMode)
func (this *QSslConfiguration) SetPeerVerifyMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration17setPeerVerifyModeEN10QSslSocket14PeerVerifyModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:103
// index:0
// Public
// int peerVerifyDepth()
func (this *QSslConfiguration) PeerVerifyDepth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration15peerVerifyDepthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:104
// index:0
// Public
// void setPeerVerifyDepth(int)
func (this *QSslConfiguration) SetPeerVerifyDepth(depth int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration18setPeerVerifyDepthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), depth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:110
// index:0
// Public
// QSslCertificate localCertificate()
func (this *QSslConfiguration) LocalCertificate() *QSslCertificate /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration16localCertificateEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:111
// index:0
// Public
// void setLocalCertificate(const QSslCertificate &)
func (this *QSslConfiguration) SetLocalCertificate(certificate *QSslCertificate) {
	var convArg0 = certificate.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration19setLocalCertificateERK15QSslCertificate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:113
// index:0
// Public
// QSslCertificate peerCertificate()
func (this *QSslConfiguration) PeerCertificate() *QSslCertificate /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration15peerCertificateEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:115
// index:0
// Public
// QSslCipher sessionCipher()
func (this *QSslConfiguration) SessionCipher() *QSslCipher /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration13sessionCipherEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:116
// index:0
// Public
// QSsl::SslProtocol sessionProtocol()
func (this *QSslConfiguration) SessionProtocol() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration15sessionProtocolEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:119
// index:0
// Public
// QSslKey privateKey()
func (this *QSslConfiguration) PrivateKey() *QSslKey /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration10privateKeyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:120
// index:0
// Public
// void setPrivateKey(const QSslKey &)
func (this *QSslConfiguration) SetPrivateKey(key *QSslKey) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration13setPrivateKeyERK7QSslKey", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:132
// index:0
// Public
// void setSslOption(QSsl::SslOption, bool)
func (this *QSslConfiguration) SetSslOption(option int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration12setSslOptionEN4QSsl9SslOptionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:133
// index:0
// Public
// bool testSslOption(QSsl::SslOption)
func (this *QSslConfiguration) TestSslOption(option int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration13testSslOptionEN4QSsl9SslOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:135
// index:0
// Public
// QByteArray sessionTicket()
func (this *QSslConfiguration) SessionTicket() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration13sessionTicketEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:136
// index:0
// Public
// void setSessionTicket(const QByteArray &)
func (this *QSslConfiguration) SetSessionTicket(sessionTicket *qtcore.QByteArray) {
	var convArg0 = sessionTicket.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration16setSessionTicketERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:137
// index:0
// Public
// int sessionTicketLifeTimeHint()
func (this *QSslConfiguration) SessionTicketLifeTimeHint() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration25sessionTicketLifeTimeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:139
// index:0
// Public
// QSslKey ephemeralServerKey()
func (this *QSslConfiguration) EphemeralServerKey() *QSslKey /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration18ephemeralServerKeyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:146
// index:0
// Public
// QByteArray preSharedKeyIdentityHint()
func (this *QSslConfiguration) PreSharedKeyIdentityHint() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration24preSharedKeyIdentityHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:147
// index:0
// Public
// void setPreSharedKeyIdentityHint(const QByteArray &)
func (this *QSslConfiguration) SetPreSharedKeyIdentityHint(hint *qtcore.QByteArray) {
	var convArg0 = hint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration27setPreSharedKeyIdentityHintERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:149
// index:0
// Public
// QSslDiffieHellmanParameters diffieHellmanParameters()
func (this *QSslConfiguration) DiffieHellmanParameters() *QSslDiffieHellmanParameters /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration23diffieHellmanParametersEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:150
// index:0
// Public
// void setDiffieHellmanParameters(const QSslDiffieHellmanParameters &)
func (this *QSslConfiguration) SetDiffieHellmanParameters(dhparams *QSslDiffieHellmanParameters) {
	var convArg0 = dhparams.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration26setDiffieHellmanParametersERK27QSslDiffieHellmanParameters", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:152
// index:0
// Public static
// QSslConfiguration defaultConfiguration()
func (this *QSslConfiguration) DefaultConfiguration() *QSslConfiguration /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration20defaultConfigurationEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSslConfiguration_DefaultConfiguration() *QSslConfiguration /*123*/ {
	var nilthis *QSslConfiguration
	rv := nilthis.DefaultConfiguration()
	return rv
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:153
// index:0
// Public static
// void setDefaultConfiguration(const QSslConfiguration &)
func (this *QSslConfiguration) SetDefaultConfiguration(configuration *QSslConfiguration) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslConfiguration23setDefaultConfigurationERKS_", ffiqt.FFI_TYPE_POINTER, configuration)
	gopp.ErrPrint(err, rv)
}
func QSslConfiguration_SetDefaultConfiguration(configuration *QSslConfiguration) {
	var nilthis *QSslConfiguration
	nilthis.SetDefaultConfiguration(configuration)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:168
// index:0
// Public
// QByteArray nextNegotiatedProtocol()
func (this *QSslConfiguration) NextNegotiatedProtocol() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration22nextNegotiatedProtocolEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:169
// index:0
// Public
// QSslConfiguration::NextProtocolNegotiationStatus nextProtocolNegotiationStatus()
func (this *QSslConfiguration) NextProtocolNegotiationStatus() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslConfiguration29nextProtocolNegotiationStatusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

type QSslConfiguration__NextProtocolNegotiationStatus = int

const QSslConfiguration__NextProtocolNegotiationNone QSslConfiguration__NextProtocolNegotiationStatus = 0
const QSslConfiguration__NextProtocolNegotiationNegotiated QSslConfiguration__NextProtocolNegotiationStatus = 1
const QSslConfiguration__NextProtocolNegotiationUnsupported QSslConfiguration__NextProtocolNegotiationStatus = 2

//  body block end
