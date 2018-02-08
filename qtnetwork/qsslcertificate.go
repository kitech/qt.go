package qtnetwork

// /usr/include/qt/QtNetwork/qsslcertificate.h
// #include <qsslcertificate.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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

type QSslCertificate struct {
	*qtrt.CObject
}

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
func NewQSslCertificate(device *qtcore.QIODevice /*777 QIODevice **/, format int) *QSslCertificate {
	var convArg0 = device.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateC2EP9QIODeviceN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCertificate)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:90
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslCertificate(const QByteArray &, QSsl::EncodingFormat)
func NewQSslCertificate_1(data *qtcore.QByteArray, format int) *QSslCertificate {
	var convArg0 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateC2ERK10QByteArrayN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslCertificateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCertificate)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslCertificate()
func DeleteQSslCertificate(this *QSslCertificate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslCertificate &)
func (this *QSslCertificate) Swap(other *QSslCertificate) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificate4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QSslCertificate) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBlacklisted()
func (this *QSslCertificate) IsBlacklisted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate13isBlacklistedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSelfSigned()
func (this *QSslCertificate) IsSelfSigned() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate12isSelfSignedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QSslCertificate) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSslCertificate5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray version()
func (this *QSslCertificate) Version() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate7versionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray serialNumber()
func (this *QSslCertificate) SerialNumber() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate12serialNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray digest(QCryptographicHash::Algorithm)
func (this *QSslCertificate) Digest(algorithm int) *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate6digestEN18QCryptographicHash9AlgorithmE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), algorithm)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime effectiveDate()
func (this *QSslCertificate) EffectiveDate() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate13effectiveDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime expiryDate()
func (this *QSslCertificate) ExpiryDate() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate10expiryDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslKey publicKey()
func (this *QSslCertificate) PublicKey() *QSslKey /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate9publicKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:137
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toPem()
func (this *QSslCertificate) ToPem() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate5toPemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toDer()
func (this *QSslCertificate) ToDer() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate5toDerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toText()
func (this *QSslCertificate) ToText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate6toTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] Qt::HANDLE handle()
func (this *QSslCertificate) Handle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSslCertificate6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

type QSslCertificate__SubjectInfo = int

const QSslCertificate__Organization QSslCertificate__SubjectInfo = 0
const QSslCertificate__CommonName QSslCertificate__SubjectInfo = 1
const QSslCertificate__LocalityName QSslCertificate__SubjectInfo = 2
const QSslCertificate__OrganizationalUnitName QSslCertificate__SubjectInfo = 3
const QSslCertificate__CountryName QSslCertificate__SubjectInfo = 4
const QSslCertificate__StateOrProvinceName QSslCertificate__SubjectInfo = 5
const QSslCertificate__DistinguishedNameQualifier QSslCertificate__SubjectInfo = 6
const QSslCertificate__SerialNumber QSslCertificate__SubjectInfo = 7
const QSslCertificate__EmailAddress QSslCertificate__SubjectInfo = 8

//  body block end
