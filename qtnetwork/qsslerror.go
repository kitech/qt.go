package qtnetwork

// /usr/include/qt/QtNetwork/qsslerror.h
// #include <qsslerror.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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
type QSslError struct {
	*qtrt.CObject
}
type QSslError_ITF interface {
	QSslError_PTR() *QSslError
}

func (ptr *QSslError) QSslError_PTR() *QSslError { return ptr }

func (this *QSslError) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslError) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslErrorFromPointer(cthis unsafe.Pointer) *QSslError {
	return &QSslError{&qtrt.CObject{cthis}}
}
func (*QSslError) NewFromPointer(cthis unsafe.Pointer) *QSslError {
	return NewQSslErrorFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslerror.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslError()

/*
Constructs a QSslError object with no error and default certificate.
*/
func NewQSslError() *QSslError {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSslErrorC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslErrorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslError)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslerror.h:89
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslError(QSslError::SslError)

/*
Constructs a QSslError object with no error and default certificate.
*/
func NewQSslError_1(error int) *QSslError {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSslErrorC2ENS_8SslErrorE", qtrt.FFI_TYPE_POINTER, error)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslErrorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslError)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslerror.h:90
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSslError(QSslError::SslError, const QSslCertificate &)

/*
Constructs a QSslError object with no error and default certificate.
*/
func NewQSslError_2(error int, certificate QSslCertificate_ITF) *QSslError {
	var convArg1 unsafe.Pointer
	if certificate != nil && certificate.QSslCertificate_PTR() != nil {
		convArg1 = certificate.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSslErrorC2ENS_8SslErrorERK15QSslCertificate", qtrt.FFI_TYPE_POINTER, error, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslErrorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslError)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslerror.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslError &)

/*
Swaps this error instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QSslError) Swap(other QSslError_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslError_PTR() != nil {
		convArg0 = other.QSslError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSslError4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslerror.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslError()

/*

 */
func DeleteQSslError(this *QSslError) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSslErrorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslerror.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSslError & operator=(QSslError &&)

/*

 */
func (this *QSslError) Operator_equal(other unsafe.Pointer /*333*/) *QSslError {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSslErroraSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslErrorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslError)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslerror.h:101
// index:1
// Public Visibility=Default Availability=Available
// [8] QSslError & operator=(const QSslError &)

/*

 */
func (this *QSslError) Operator_equal_1(other QSslError_ITF) *QSslError {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslError_PTR() != nil {
		convArg0 = other.QSslError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSslErroraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslErrorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslError)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslerror.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QSslError &) const

/*

 */
func (this *QSslError) Operator_equal_equal(other QSslError_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslError_PTR() != nil {
		convArg0 = other.QSslError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSslErroreqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslerror.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QSslError &) const

/*

 */
func (this *QSslError) Operator_not_equal(other QSslError_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslError_PTR() != nil {
		convArg0 = other.QSslError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSslErrorneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslerror.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslError::SslError error() const

/*
Returns the type of the error.

See also errorString() and certificate().
*/
func (this *QSslError) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSslError5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslerror.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a short localized human-readable description of the error.

See also error() and certificate().
*/
func (this *QSslError) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSslError11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslerror.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslCertificate certificate() const

/*
Returns the certificate associated with this error, or a null certificate if the error does not relate to any certificate.

See also error() and errorString().
*/
func (this *QSslError) Certificate() *QSslCertificate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSslError11certificateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

/*
Describes all recognized errors that can occur during an SSL handshake.

ConstantValue
QSslError::NoError0
QSslError::UnableToGetIssuerCertificate1
QSslError::UnableToDecryptCertificateSignature2
QSslError::UnableToDecodeIssuerPublicKey3
QSslError::CertificateSignatureFailed4
QSslError::CertificateNotYetValid5
QSslError::CertificateExpired6
QSslError::InvalidNotBeforeField7
QSslError::InvalidNotAfterField8
QSslError::SelfSignedCertificate9
QSslError::UnspecifiedError-1


See also QSslError::errorString().

*/
type QSslError__SslError = int

//
const QSslError__NoError QSslError__SslError = 0

//
const QSslError__UnableToGetIssuerCertificate QSslError__SslError = 1

//
const QSslError__UnableToDecryptCertificateSignature QSslError__SslError = 2

//
const QSslError__UnableToDecodeIssuerPublicKey QSslError__SslError = 3

//
const QSslError__CertificateSignatureFailed QSslError__SslError = 4

//
const QSslError__CertificateNotYetValid QSslError__SslError = 5

//
const QSslError__CertificateExpired QSslError__SslError = 6

//
const QSslError__InvalidNotBeforeField QSslError__SslError = 7

//
const QSslError__InvalidNotAfterField QSslError__SslError = 8

//
const QSslError__SelfSignedCertificate QSslError__SslError = 9

// 0
const QSslError__SelfSignedCertificateInChain QSslError__SslError = 10

// 1
const QSslError__UnableToGetLocalIssuerCertificate QSslError__SslError = 11

// 2
const QSslError__UnableToVerifyFirstCertificate QSslError__SslError = 12

// 3
const QSslError__CertificateRevoked QSslError__SslError = 13

// 4
const QSslError__InvalidCaCertificate QSslError__SslError = 14

// 5
const QSslError__PathLengthExceeded QSslError__SslError = 15

// 6
const QSslError__InvalidPurpose QSslError__SslError = 16

// 7
const QSslError__CertificateUntrusted QSslError__SslError = 17

// 8
const QSslError__CertificateRejected QSslError__SslError = 18

// 9
const QSslError__SubjectIssuerMismatch QSslError__SslError = 19

// 0
const QSslError__AuthorityIssuerSerialNumberMismatch QSslError__SslError = 20

// 1
const QSslError__NoPeerCertificate QSslError__SslError = 21

// 2
const QSslError__HostNameMismatch QSslError__SslError = 22

// 3
const QSslError__NoSslSupport QSslError__SslError = 23

// 4
const QSslError__CertificateBlacklisted QSslError__SslError = 24

//
const QSslError__UnspecifiedError QSslError__SslError = -1

func (this *QSslError) SslErrorItemName(val int) string {
	switch val {
	case QSslError__NoError: // 0
		return "NoError"
	case QSslError__UnableToGetIssuerCertificate: // 1
		return "UnableToGetIssuerCertificate"
	case QSslError__UnableToDecryptCertificateSignature: // 2
		return "UnableToDecryptCertificateSignature"
	case QSslError__UnableToDecodeIssuerPublicKey: // 3
		return "UnableToDecodeIssuerPublicKey"
	case QSslError__CertificateSignatureFailed: // 4
		return "CertificateSignatureFailed"
	case QSslError__CertificateNotYetValid: // 5
		return "CertificateNotYetValid"
	case QSslError__CertificateExpired: // 6
		return "CertificateExpired"
	case QSslError__InvalidNotBeforeField: // 7
		return "InvalidNotBeforeField"
	case QSslError__InvalidNotAfterField: // 8
		return "InvalidNotAfterField"
	case QSslError__SelfSignedCertificate: // 9
		return "SelfSignedCertificate"
	case QSslError__SelfSignedCertificateInChain: // 10
		return "SelfSignedCertificateInChain"
	case QSslError__UnableToGetLocalIssuerCertificate: // 11
		return "UnableToGetLocalIssuerCertificate"
	case QSslError__UnableToVerifyFirstCertificate: // 12
		return "UnableToVerifyFirstCertificate"
	case QSslError__CertificateRevoked: // 13
		return "CertificateRevoked"
	case QSslError__InvalidCaCertificate: // 14
		return "InvalidCaCertificate"
	case QSslError__PathLengthExceeded: // 15
		return "PathLengthExceeded"
	case QSslError__InvalidPurpose: // 16
		return "InvalidPurpose"
	case QSslError__CertificateUntrusted: // 17
		return "CertificateUntrusted"
	case QSslError__CertificateRejected: // 18
		return "CertificateRejected"
	case QSslError__SubjectIssuerMismatch: // 19
		return "SubjectIssuerMismatch"
	case QSslError__AuthorityIssuerSerialNumberMismatch: // 20
		return "AuthorityIssuerSerialNumberMismatch"
	case QSslError__NoPeerCertificate: // 21
		return "NoPeerCertificate"
	case QSslError__HostNameMismatch: // 22
		return "HostNameMismatch"
	case QSslError__NoSslSupport: // 23
		return "NoSslSupport"
	case QSslError__CertificateBlacklisted: // 24
		return "CertificateBlacklisted"
	case QSslError__UnspecifiedError: // -1
		return "UnspecifiedError"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSslError_SslErrorItemName(val int) string {
	var nilthis *QSslError
	return nilthis.SslErrorItemName(val)
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
