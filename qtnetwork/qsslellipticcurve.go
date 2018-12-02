package qtnetwork

// /usr/include/qt/QtNetwork/qsslellipticcurve.h
// #include <qsslellipticcurve.h>
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

/*

 */
type QSslEllipticCurve struct {
	*qtrt.CObject
}
type QSslEllipticCurve_ITF interface {
	QSslEllipticCurve_PTR() *QSslEllipticCurve
}

func (ptr *QSslEllipticCurve) QSslEllipticCurve_PTR() *QSslEllipticCurve { return ptr }

func (this *QSslEllipticCurve) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslEllipticCurve) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslEllipticCurveFromPointer(cthis unsafe.Pointer) *QSslEllipticCurve {
	return &QSslEllipticCurve{&qtrt.CObject{cthis}}
}
func (*QSslEllipticCurve) NewFromPointer(cthis unsafe.Pointer) *QSslEllipticCurve {
	return NewQSslEllipticCurveFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSslEllipticCurve()

/*
Constructs an invalid elliptic curve.

See also isValid() and QSslConfiguration::supportedEllipticCurves().
*/
func (*QSslEllipticCurve) NewForInherit() *QSslEllipticCurve {
	return NewQSslEllipticCurve()
}
func NewQSslEllipticCurve() *QSslEllipticCurve {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslEllipticCurveC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslEllipticCurveFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslEllipticCurve)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:64
// index:0
// Public static Visibility=Default Availability=Available
// [4] QSslEllipticCurve fromShortName(const QString &)

/*
Returns an QSslEllipticCurve instance representing the named curve name. The name is the conventional short name for the curve, as represented by RFC 4492 (for instance secp521r1), or as NIST short names (for instance P-256). The actual set of recognized names depends on the SSL implementation.

If the given name is not supported, returns an invalid QSslEllipticCurve instance.

Note: The OpenSSL implementation of this function treats the name case-sensitively.

See also shortName().
*/
func (this *QSslEllipticCurve) FromShortName(name string) *QSslEllipticCurve /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslEllipticCurve13fromShortNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslEllipticCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslEllipticCurve)
	return rv2
}
func QSslEllipticCurve_FromShortName(name string) *QSslEllipticCurve /*123*/ {
	var nilthis *QSslEllipticCurve
	rv := nilthis.FromShortName(name)
	return rv
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:65
// index:0
// Public static Visibility=Default Availability=Available
// [4] QSslEllipticCurve fromLongName(const QString &)

/*
Returns an QSslEllipticCurve instance representing the named curve name. The name is a long name for the curve, whose exact spelling depends on the SSL implementation.

If the given name is not supported, returns an invalid QSslEllipticCurve instance.

Note: The OpenSSL implementation of this function treats the name case-sensitively.

See also longName().
*/
func (this *QSslEllipticCurve) FromLongName(name string) *QSslEllipticCurve /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslEllipticCurve12fromLongNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslEllipticCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslEllipticCurve)
	return rv2
}
func QSslEllipticCurve_FromLongName(name string) *QSslEllipticCurve /*123*/ {
	var nilthis *QSslEllipticCurve
	rv := nilthis.FromLongName(name)
	return rv
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString shortName() const

/*
Returns the conventional short name for this curve. If this curve is invalid, returns an empty string.

See also longName().
*/
func (this *QSslEllipticCurve) ShortName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslEllipticCurve9shortNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString longName() const

/*
Returns the conventional long name for this curve. If this curve is invalid, returns an empty string.

See also shortName().
*/
func (this *QSslEllipticCurve) LongName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslEllipticCurve8longNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this elliptic curve is a valid curve, false otherwise.
*/
func (this *QSslEllipticCurve) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslEllipticCurve7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTlsNamedCurve() const

/*
Returns true if this elliptic curve is one of the named curves that can be used in the key exchange when using an elliptic curve cipher with TLS; false otherwise.
*/
func (this *QSslEllipticCurve) IsTlsNamedCurve() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslEllipticCurve15isTlsNamedCurveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQSslEllipticCurve(this *QSslEllipticCurve) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslEllipticCurveD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
