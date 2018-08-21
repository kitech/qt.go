package qtnetwork

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h
// #include <qssldiffiehellmanparameters.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QSslDiffieHellmanParameters struct {
	*qtrt.CObject
}
type QSslDiffieHellmanParameters_ITF interface {
	QSslDiffieHellmanParameters_PTR() *QSslDiffieHellmanParameters
}

func (ptr *QSslDiffieHellmanParameters) QSslDiffieHellmanParameters_PTR() *QSslDiffieHellmanParameters {
	return ptr
}

func (this *QSslDiffieHellmanParameters) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslDiffieHellmanParameters) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslDiffieHellmanParametersFromPointer(cthis unsafe.Pointer) *QSslDiffieHellmanParameters {
	return &QSslDiffieHellmanParameters{&qtrt.CObject{cthis}}
}
func (*QSslDiffieHellmanParameters) NewFromPointer(cthis unsafe.Pointer) *QSslDiffieHellmanParameters {
	return NewQSslDiffieHellmanParametersFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:82
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSslDiffieHellmanParameters defaultParameters()

/*
Returns the default QSslDiffieHellmanParameters used by QSslSocket.

This is currently the 1024-bit MODP group from RFC 2459, also known as the Second Oakley Group.
*/
func (this *QSslDiffieHellmanParameters) DefaultParameters() *QSslDiffieHellmanParameters /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParameters17defaultParametersEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslDiffieHellmanParameters)
	return rv2
}
func QSslDiffieHellmanParameters_DefaultParameters() *QSslDiffieHellmanParameters /*123*/ {
	var nilthis *QSslDiffieHellmanParameters
	rv := nilthis.DefaultParameters()
	return rv
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslDiffieHellmanParameters()

/*
Constructs an empty QSslDiffieHellmanParameters instance.

If an empty QSslDiffieHellmanParameters instance is set on a QSslConfiguration object, Diffie-Hellman negotiation will be disabled.

See also isValid() and QSslConfiguration.
*/
func NewQSslDiffieHellmanParameters() *QSslDiffieHellmanParameters {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParametersC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslDiffieHellmanParameters)
	return gothis
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslDiffieHellmanParameters()

/*

 */
func DeleteQSslDiffieHellmanParameters(this *QSslDiffieHellmanParameters) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParametersD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslDiffieHellmanParameters & operator=(const QSslDiffieHellmanParameters &)

/*

 */
func (this *QSslDiffieHellmanParameters) Operator_equal(other QSslDiffieHellmanParameters_ITF) *QSslDiffieHellmanParameters {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslDiffieHellmanParameters_PTR() != nil {
		convArg0 = other.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParametersaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslDiffieHellmanParameters)
	return rv2
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:90
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QSslDiffieHellmanParameters & operator=(QSslDiffieHellmanParameters &&)

/*

 */
func (this *QSslDiffieHellmanParameters) Operator_equal_1(other unsafe.Pointer /*333*/) *QSslDiffieHellmanParameters {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParametersaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslDiffieHellmanParameters)
	return rv2
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslDiffieHellmanParameters &)

/*
Swaps this QSslDiffieHellmanParameters with other. This function is very fast and never fails.
*/
func (this *QSslDiffieHellmanParameters) Swap(other QSslDiffieHellmanParameters_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslDiffieHellmanParameters_PTR() != nil {
		convArg0 = other.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParameters4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSslDiffieHellmanParameters fromEncoded(const QByteArray &, QSsl::EncodingFormat)

/*
Constructs a QSslDiffieHellmanParameters object using the byte array encoded in either PEM or DER form as specified by encoding.

Use the isValid() method on the returned object to check whether the Diffie-Hellman parameters were valid and loaded correctly.

See also isValid() and QSslConfiguration.
*/
func (this *QSslDiffieHellmanParameters) FromEncoded(encoded qtcore.QByteArray_ITF, format int) *QSslDiffieHellmanParameters /*123*/ {
	var convArg0 unsafe.Pointer
	if encoded != nil && encoded.QByteArray_PTR() != nil {
		convArg0 = encoded.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParameters11fromEncodedERK10QByteArrayN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslDiffieHellmanParameters)
	return rv2
}
func QSslDiffieHellmanParameters_FromEncoded(encoded qtcore.QByteArray_ITF, format int) *QSslDiffieHellmanParameters /*123*/ {
	var nilthis *QSslDiffieHellmanParameters
	rv := nilthis.FromEncoded(encoded, format)
	return rv
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSslDiffieHellmanParameters fromEncoded(const QByteArray &, QSsl::EncodingFormat)

/*
Constructs a QSslDiffieHellmanParameters object using the byte array encoded in either PEM or DER form as specified by encoding.

Use the isValid() method on the returned object to check whether the Diffie-Hellman parameters were valid and loaded correctly.

See also isValid() and QSslConfiguration.
*/
func (this *QSslDiffieHellmanParameters) FromEncoded__(encoded qtcore.QByteArray_ITF) *QSslDiffieHellmanParameters /*123*/ {
	var convArg0 unsafe.Pointer
	if encoded != nil && encoded.QByteArray_PTR() != nil {
		convArg0 = encoded.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParameters11fromEncodedERK10QByteArrayN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslDiffieHellmanParameters)
	return rv2
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:95
// index:1
// Public static Visibility=Default Availability=Available
// [8] QSslDiffieHellmanParameters fromEncoded(QIODevice *, QSsl::EncodingFormat)

/*
Constructs a QSslDiffieHellmanParameters object using the byte array encoded in either PEM or DER form as specified by encoding.

Use the isValid() method on the returned object to check whether the Diffie-Hellman parameters were valid and loaded correctly.

See also isValid() and QSslConfiguration.
*/
func (this *QSslDiffieHellmanParameters) FromEncoded_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format int) *QSslDiffieHellmanParameters /*123*/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParameters11fromEncodedEP9QIODeviceN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslDiffieHellmanParameters)
	return rv2
}
func QSslDiffieHellmanParameters_FromEncoded_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format int) *QSslDiffieHellmanParameters /*123*/ {
	var nilthis *QSslDiffieHellmanParameters
	rv := nilthis.FromEncoded_1(device, format)
	return rv
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:95
// index:1
// Public static Visibility=Default Availability=Available
// [8] QSslDiffieHellmanParameters fromEncoded(QIODevice *, QSsl::EncodingFormat)

/*
Constructs a QSslDiffieHellmanParameters object using the byte array encoded in either PEM or DER form as specified by encoding.

Use the isValid() method on the returned object to check whether the Diffie-Hellman parameters were valid and loaded correctly.

See also isValid() and QSslConfiguration.
*/
func (this *QSslDiffieHellmanParameters) FromEncoded_1_(device qtcore.QIODevice_ITF /*777 QIODevice **/) *QSslDiffieHellmanParameters /*123*/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 1, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParameters11fromEncodedEP9QIODeviceN4QSsl14EncodingFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslDiffieHellmanParameters)
	return rv2
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if this is a an empty QSslDiffieHellmanParameters instance.

Setting an empty QSslDiffieHellmanParameters instance on a QSslSocket-based server will disable Diffie-Hellman key exchange.
*/
func (this *QSslDiffieHellmanParameters) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QSslDiffieHellmanParameters7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this is a valid QSslDiffieHellmanParameters; otherwise false.

This method should be used after constructing a QSslDiffieHellmanParameters object to determine its validity.

If a QSslDiffieHellmanParameters object is not valid, you can use the error() method to determine what error prevented the object from being constructed.

See also error().
*/
func (this *QSslDiffieHellmanParameters) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QSslDiffieHellmanParameters7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QSslDiffieHellmanParameters::Error error() const

/*
Returns the error that caused the QSslDiffieHellmanParameters object to be invalid.
*/
func (this *QSslDiffieHellmanParameters) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QSslDiffieHellmanParameters5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a human-readable description of the error that caused the QSslDiffieHellmanParameters object to be invalid.
*/
func (this *QSslDiffieHellmanParameters) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QSslDiffieHellmanParameters11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

/*
Describes a QSslDiffieHellmanParameters error.


*/
type QSslDiffieHellmanParameters__Error = int

// No error occurred.
const QSslDiffieHellmanParameters__NoError QSslDiffieHellmanParameters__Error = 0

// The given input data could not be used to construct a QSslDiffieHellmanParameters object.
const QSslDiffieHellmanParameters__InvalidInputDataError QSslDiffieHellmanParameters__Error = 1

// The Diffie-Hellman parameters are unsafe and should not be used.
const QSslDiffieHellmanParameters__UnsafeParametersError QSslDiffieHellmanParameters__Error = 2

func (this *QSslDiffieHellmanParameters) ErrorItemName(val int) string {
	switch val {
	case QSslDiffieHellmanParameters__NoError: // 0
		return "NoError"
	case QSslDiffieHellmanParameters__InvalidInputDataError: // 1
		return "InvalidInputDataError"
	case QSslDiffieHellmanParameters__UnsafeParametersError: // 2
		return "UnsafeParametersError"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSslDiffieHellmanParameters_ErrorItemName(val int) string {
	var nilthis *QSslDiffieHellmanParameters
	return nilthis.ErrorItemName(val)
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
