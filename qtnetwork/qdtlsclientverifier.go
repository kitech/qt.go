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
type QDtlsClientVerifier struct {
	*qtcore.QObject
}
type QDtlsClientVerifier_ITF interface {
	qtcore.QObject_ITF
	QDtlsClientVerifier_PTR() *QDtlsClientVerifier
}

func (ptr *QDtlsClientVerifier) QDtlsClientVerifier_PTR() *QDtlsClientVerifier { return ptr }

func (this *QDtlsClientVerifier) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QDtlsClientVerifier) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQDtlsClientVerifierFromPointer(cthis unsafe.Pointer) *QDtlsClientVerifier {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QDtlsClientVerifier{bcthis0}
}
func (*QDtlsClientVerifier) NewFromPointer(cthis unsafe.Pointer) *QDtlsClientVerifier {
	return NewQDtlsClientVerifierFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qdtls.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDtlsClientVerifier) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QDtlsClientVerifier10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qdtls.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDtlsClientVerifier(QObject *)

/*

 */
func (*QDtlsClientVerifier) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QDtlsClientVerifier {
	return NewQDtlsClientVerifier(parent)
}
func NewQDtlsClientVerifier(parent qtcore.QObject_ITF /*777 QObject **/) *QDtlsClientVerifier {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QDtlsClientVerifierC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDtlsClientVerifierFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDtlsClientVerifier")
	return gothis
}

// /usr/include/qt/QtNetwork/qdtls.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDtlsClientVerifier(QObject *)

/*

 */
func (*QDtlsClientVerifier) NewForInheritp() *QDtlsClientVerifier {
	return NewQDtlsClientVerifierp()
}
func NewQDtlsClientVerifierp() *QDtlsClientVerifier {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QDtlsClientVerifierC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDtlsClientVerifierFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDtlsClientVerifier")
	return gothis
}

// /usr/include/qt/QtNetwork/qdtls.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDtlsClientVerifier()

/*

 */
func DeleteQDtlsClientVerifier(this *QDtlsClientVerifier) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QDtlsClientVerifierD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qdtls.h:92
// index:0
// Public Visibility=Default Availability=Available
// [16] QDtlsClientVerifier::GeneratorParameters cookieGeneratorParameters() const

/*
Returns the current hash algorithm and secret, either default ones or previously set by a call to setCookieGeneratorParameters().

The default hash algorithm is QCryptographicHash::Sha256 if Qt was configured to support it, QCryptographicHash::Sha1 otherwise. The default secret is obtained from the backend-specific cryptographically strong pseudorandom number generator.

See also setCookieGeneratorParameters(), QDtlsClientVerifier, and cookieGeneratorParameters().
*/
func (this *QDtlsClientVerifier) CookieGeneratorParameters() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QDtlsClientVerifier25cookieGeneratorParametersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtNetwork/qdtls.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool verifyClient(QUdpSocket *, const QByteArray &, const QHostAddress &, quint16)

/*

 */
func (this *QDtlsClientVerifier) VerifyClient(socket QUdpSocket_ITF /*777 QUdpSocket **/, dgram qtcore.QByteArray_ITF, address QHostAddress_ITF, port uint16) bool {
	var convArg0 unsafe.Pointer
	if socket != nil && socket.QUdpSocket_PTR() != nil {
		convArg0 = socket.QUdpSocket_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if dgram != nil && dgram.QByteArray_PTR() != nil {
		convArg1 = dgram.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if address != nil && address.QHostAddress_PTR() != nil {
		convArg2 = address.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QDtlsClientVerifier12verifyClientEP10QUdpSocketRK10QByteArrayRK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, port)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qdtls.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray verifiedHello() const

/*

 */
func (this *QDtlsClientVerifier) VerifiedHello() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QDtlsClientVerifier13verifiedHelloEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qdtls.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] QDtlsError dtlsError() const

/*
Returns the last error encountered by the connection or QDtlsError::NoError.

See also dtlsErrorString() and QDtlsError.
*/
func (this *QDtlsClientVerifier) DtlsError() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QDtlsClientVerifier9dtlsErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qdtls.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QString dtlsErrorString() const

/*
Returns a textual description for the last error encountered by the connection or empty string.

See also dtlsError().
*/
func (this *QDtlsClientVerifier) DtlsErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QDtlsClientVerifier15dtlsErrorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

//  body block end

//  keep block begin

func init_unused_11409() {
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
