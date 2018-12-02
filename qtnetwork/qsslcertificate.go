package qtnetwork

// /usr/include/qt/QtNetwork/qsslcertificate.h
// #include <qsslcertificate.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QSslCertificate struct {
	*qtrt.CObject
}
type QSslCertificate_ITF interface {
	QSslCertificate_PTR() *QSslCertificate
}

func (ptr *QSslCertificate) QSslCertificate_PTR() *QSslCertificate { return ptr }

func (this *QSslCertificate) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslCertificate) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslCertificateFromPointer(cthis unsafe.Pointer) *QSslCertificate {
	return &QSslCertificate{&qtrt.CObject{cthis}}
}
func (*QSslCertificate) NewFromPointer(cthis unsafe.Pointer) *QSslCertificate {
	return NewQSslCertificateFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslCertificate(QIODevice *, QSsl::EncodingFormat)

/*
Constructs a QSslCertificate by reading format encoded data from device and using the first certificate found. You can later call isNull() to see if device contained a certificate, and if this certificate was loaded successfully.
*/
func (*QSslCertificate) NewForInherit(device qtcore.QIODevice_ITF /*777 QIODevice **/, format int) *QSslCertificate {
	return NewQSslCertificate(device, format)
}
func NewQSslCertificate(device qtcore.QIODevice_ITF /*777 QIODevice **/, format int) *QSslCertificate {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateC2EP9QIODeviceN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCertificate)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslCertificate(QIODevice *, QSsl::EncodingFormat)

/*
Constructs a QSslCertificate by reading format encoded data from device and using the first certificate found. You can later call isNull() to see if device contained a certificate, and if this certificate was loaded successfully.
*/
func (*QSslCertificate) NewForInherit__(device qtcore.QIODevice_ITF /*777 QIODevice **/) *QSslCertificate {
	return NewQSslCertificate__(device)
}
func NewQSslCertificate__(device qtcore.QIODevice_ITF /*777 QIODevice **/) *QSslCertificate {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 1, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateC2EP9QIODeviceN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCertificate)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:90
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslCertificate(const QByteArray &, QSsl::EncodingFormat)

/*
Constructs a QSslCertificate by reading format encoded data from device and using the first certificate found. You can later call isNull() to see if device contained a certificate, and if this certificate was loaded successfully.
*/
func (*QSslCertificate) NewForInherit_1(data qtcore.QByteArray_ITF, format int) *QSslCertificate {
	return NewQSslCertificate_1(data, format)
}
func NewQSslCertificate_1(data qtcore.QByteArray_ITF, format int) *QSslCertificate {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateC2ERK10QByteArrayN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCertificate)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:90
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslCertificate(const QByteArray &, QSsl::EncodingFormat)

/*
Constructs a QSslCertificate by reading format encoded data from device and using the first certificate found. You can later call isNull() to see if device contained a certificate, and if this certificate was loaded successfully.
*/
func (*QSslCertificate) NewForInherit_1_() *QSslCertificate {
	return NewQSslCertificate_1_()
}
func NewQSslCertificate_1_() *QSslCertificate {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg0 = qtcore.NewQByteArray()
	// arg: 1, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateC2ERK10QByteArrayN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCertificate)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:90
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslCertificate(const QByteArray &, QSsl::EncodingFormat)

/*
Constructs a QSslCertificate by reading format encoded data from device and using the first certificate found. You can later call isNull() to see if device contained a certificate, and if this certificate was loaded successfully.
*/
func (*QSslCertificate) NewForInherit_1_1(data qtcore.QByteArray_ITF) *QSslCertificate {
	return NewQSslCertificate_1_1(data)
}
func NewQSslCertificate_1_1(data qtcore.QByteArray_ITF) *QSslCertificate {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateC2ERK10QByteArrayN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCertificate)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslCertificate()

/*

 */
func DeleteQSslCertificate(this *QSslCertificate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSslCertificate & operator=(QSslCertificate &&)

/*

 */
func (this *QSslCertificate) Operator_equal(other unsafe.Pointer /*333*/) *QSslCertificate {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:96
// index:1
// Public Visibility=Default Availability=Available
// [8] QSslCertificate & operator=(const QSslCertificate &)

/*

 */
func (this *QSslCertificate) Operator_equal_1(other QSslCertificate_ITF) *QSslCertificate {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslCertificate_PTR() != nil {
		convArg0 = other.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificate)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslCertificate &)

/*
Swaps this certificate instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QSslCertificate) Swap(other QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslCertificate_PTR() != nil {
		convArg0 = other.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificate4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:101
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QSslCertificate &) const

/*

 */
func (this *QSslCertificate) Operator_equal_equal(other QSslCertificate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslCertificate_PTR() != nil {
		convArg0 = other.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificateeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QSslCertificate &) const

/*

 */
func (this *QSslCertificate) Operator_not_equal(other QSslCertificate_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslCertificate_PTR() != nil {
		convArg0 = other.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificateneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this is a null certificate (i.e., a certificate with no contents); otherwise returns false.

By default, QSslCertificate constructs a null certificate.

See also clear().
*/
func (this *QSslCertificate) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBlacklisted() const

/*
Returns true if this certificate is blacklisted; otherwise returns false.

See also isNull().
*/
func (this *QSslCertificate) IsBlacklisted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate13isBlacklistedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSelfSigned() const

/*
Returns true if this certificate is self signed; otherwise returns false.

A certificate is considered self-signed its issuer and subject are identical.

This function was introduced in  Qt 5.4.
*/
func (this *QSslCertificate) IsSelfSigned() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate12isSelfSignedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the contents of this certificate, making it a null certificate.

See also isNull().
*/
func (this *QSslCertificate) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificate5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray version() const

/*
Returns the certificate's version string.
*/
func (this *QSslCertificate) Version() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate7versionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray serialNumber() const

/*
Returns the certificate's serial number string in hexadecimal format.
*/
func (this *QSslCertificate) SerialNumber() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate12serialNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray digest(QCryptographicHash::Algorithm) const

/*
Returns a cryptographic digest of this certificate. By default, an MD5 digest will be generated, but you can also specify a custom algorithm.
*/
func (this *QSslCertificate) Digest(algorithm int) *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate6digestEN18QCryptographicHash9AlgorithmE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), algorithm)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray digest(QCryptographicHash::Algorithm) const

/*
Returns a cryptographic digest of this certificate. By default, an MD5 digest will be generated, but you can also specify a custom algorithm.
*/
func (this *QSslCertificate) Digest__() *qtcore.QByteArray /*123*/ {
	// arg: 0, QCryptographicHash::Algorithm=Elaborated, QCryptographicHash::Algorithm=Enum, , Invalid
	algorithm := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate6digestEN18QCryptographicHash9AlgorithmE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), algorithm)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList issuerInfo(QSslCertificate::SubjectInfo) const

/*
Returns the issuer information for the subject from the certificate, or an empty list if there is no information for subject in the certificate. There can be more than one entry of each type.

See also subjectInfo().
*/
func (this *QSslCertificate) IssuerInfo(info int) *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate10issuerInfoENS_11SubjectInfoE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), info)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:122
// index:1
// Public Visibility=Default Availability=Available
// [8] QStringList issuerInfo(const QByteArray &) const

/*
Returns the issuer information for the subject from the certificate, or an empty list if there is no information for subject in the certificate. There can be more than one entry of each type.

See also subjectInfo().
*/
func (this *QSslCertificate) IssuerInfo_1(attribute qtcore.QByteArray_ITF) *qtcore.QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if attribute != nil && attribute.QByteArray_PTR() != nil {
		convArg0 = attribute.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate10issuerInfoERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList subjectInfo(QSslCertificate::SubjectInfo) const

/*
Returns the information for the subject, or an empty list if there is no information for subject in the certificate. There can be more than one entry of each type.

See also issuerInfo().
*/
func (this *QSslCertificate) SubjectInfo(info int) *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate11subjectInfoENS_11SubjectInfoE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), info)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:124
// index:1
// Public Visibility=Default Availability=Available
// [8] QStringList subjectInfo(const QByteArray &) const

/*
Returns the information for the subject, or an empty list if there is no information for subject in the certificate. There can be more than one entry of each type.

See also issuerInfo().
*/
func (this *QSslCertificate) SubjectInfo_1(attribute qtcore.QByteArray_ITF) *qtcore.QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if attribute != nil && attribute.QByteArray_PTR() != nil {
		convArg0 = attribute.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate11subjectInfoERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime effectiveDate() const

/*
Returns the date-time that the certificate becomes valid, or an empty QDateTime if this is a null certificate.

See also expiryDate().
*/
func (this *QSslCertificate) EffectiveDate() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate13effectiveDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime expiryDate() const

/*
Returns the date-time that the certificate expires, or an empty QDateTime if this is a null certificate.

See also effectiveDate().
*/
func (this *QSslCertificate) ExpiryDate() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate10expiryDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslKey publicKey() const

/*
Returns the certificate subject's public key.
*/
func (this *QSslCertificate) PublicKey() *QSslKey /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate9publicKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toPem() const

/*
Returns this certificate converted to a PEM (Base64) encoded representation.
*/
func (this *QSslCertificate) ToPem() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate5toPemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toDer() const

/*
Returns this certificate converted to a DER (binary) encoded representation.
*/
func (this *QSslCertificate) ToDer() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate5toDerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toText() const

/*
Returns this certificate converted to a human-readable text representation.

This function was introduced in  Qt 5.0.
*/
func (this *QSslCertificate) ToText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate6toTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] Qt::HANDLE handle() const

/*
Returns a pointer to the native certificate handle, if there is one, or a null pointer otherwise.

You can use this handle, together with the native API, to access extended information about the certificate.

Warning: Use of this function has a high probability of being non-portable, and its return value may vary from platform to platform or change from minor release to minor release.
*/
func (this *QSslCertificate) Handle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*
Describes keys that you can pass to QSslCertificate::issuerInfo() or QSslCertificate::subjectInfo() to get information about the certificate issuer or subject.


*/
type QSslCertificate__SubjectInfo = int

// "O" The name of the organization.
const QSslCertificate__Organization QSslCertificate__SubjectInfo = 0

// "CN" The common name; most often this is used to store the host name.
const QSslCertificate__CommonName QSslCertificate__SubjectInfo = 1

// "L" The locality.
const QSslCertificate__LocalityName QSslCertificate__SubjectInfo = 2

// "OU" The organizational unit name.
const QSslCertificate__OrganizationalUnitName QSslCertificate__SubjectInfo = 3

// "C" The country.
const QSslCertificate__CountryName QSslCertificate__SubjectInfo = 4

// "ST" The state or province.
const QSslCertificate__StateOrProvinceName QSslCertificate__SubjectInfo = 5

// The distinguished name qualifier
const QSslCertificate__DistinguishedNameQualifier QSslCertificate__SubjectInfo = 6

// The certificate's serial number
const QSslCertificate__SerialNumber QSslCertificate__SubjectInfo = 7

// The email address associated with the certificate
const QSslCertificate__EmailAddress QSslCertificate__SubjectInfo = 8

func (this *QSslCertificate) SubjectInfoItemName(val int) string {
	switch val {
	case QSslCertificate__Organization: // 0
		return "Organization"
	case QSslCertificate__CommonName: // 1
		return "CommonName"
	case QSslCertificate__LocalityName: // 2
		return "LocalityName"
	case QSslCertificate__OrganizationalUnitName: // 3
		return "OrganizationalUnitName"
	case QSslCertificate__CountryName: // 4
		return "CountryName"
	case QSslCertificate__StateOrProvinceName: // 5
		return "StateOrProvinceName"
	case QSslCertificate__DistinguishedNameQualifier: // 6
		return "DistinguishedNameQualifier"
	case QSslCertificate__SerialNumber: // 7
		return "SerialNumber"
	case QSslCertificate__EmailAddress: // 8
		return "EmailAddress"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSslCertificate_SubjectInfoItemName(val int) string {
	var nilthis *QSslCertificate
	return nilthis.SubjectInfoItemName(val)
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
