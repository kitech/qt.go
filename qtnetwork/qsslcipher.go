package qtnetwork

// /usr/include/qt/QtNetwork/qsslcipher.h
// #include <qsslcipher.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QSslCipher struct {
	*qtrt.CObject
}
type QSslCipher_ITF interface {
	QSslCipher_PTR() *QSslCipher
}

func (ptr *QSslCipher) QSslCipher_PTR() *QSslCipher { return ptr }

func (this *QSslCipher) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslCipher) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslCipherFromPointer(cthis unsafe.Pointer) *QSslCipher {
	return &QSslCipher{&qtrt.CObject{cthis}}
}
func (*QSslCipher) NewFromPointer(cthis unsafe.Pointer) *QSslCipher {
	return NewQSslCipherFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslcipher.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslCipher()

/*
Constructs an empty QSslCipher object.
*/
func (*QSslCipher) NewForInherit() *QSslCipher {
	return NewQSslCipher()
}
func NewQSslCipher() *QSslCipher {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipherC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCipher)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcipher.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslCipher(const QString &)

/*
Constructs an empty QSslCipher object.
*/
func (*QSslCipher) NewForInherit1(name string) *QSslCipher {
	return NewQSslCipher1(name)
}
func NewQSslCipher1(name string) *QSslCipher {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipherC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCipher)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcipher.h:60
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSslCipher(const QString &, QSsl::SslProtocol)

/*
Constructs an empty QSslCipher object.
*/
func (*QSslCipher) NewForInherit2(name string, protocol int) *QSslCipher {
	return NewQSslCipher2(name, protocol)
}
func NewQSslCipher2(name string, protocol int) *QSslCipher {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipherC2ERK7QStringN4QSsl11SslProtocolE", qtrt.FFI_TYPE_POINTER, convArg0, protocol)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCipher)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcipher.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSslCipher & operator=(QSslCipher &&)

/*

 */
func (this *QSslCipher) Operator_equal(other unsafe.Pointer /*333*/) *QSslCipher {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipheraSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCipher)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcipher.h:65
// index:1
// Public Visibility=Default Availability=Available
// [8] QSslCipher & operator=(const QSslCipher &)

/*

 */
func (this *QSslCipher) Operator_equal1(other QSslCipher_ITF) *QSslCipher {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslCipher_PTR() != nil {
		convArg0 = other.QSslCipher_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipheraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCipher)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcipher.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslCipher()

/*

 */
func DeleteQSslCipher(this *QSslCipher) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipherD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslcipher.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslCipher &)

/*
Swaps this cipher instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QSslCipher) Swap(other QSslCipher_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslCipher_PTR() != nil {
		convArg0 = other.QSslCipher_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipher4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcipher.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QSslCipher &) const

/*

 */
func (this *QSslCipher) Operator_equal_equal(other QSslCipher_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslCipher_PTR() != nil {
		convArg0 = other.QSslCipher_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCiphereqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcipher.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QSslCipher &) const

/*

 */
func (this *QSslCipher) Operator_not_equal(other QSslCipher_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslCipher_PTR() != nil {
		convArg0 = other.QSslCipher_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipherneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcipher.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this is a null cipher; otherwise returns false.
*/
func (this *QSslCipher) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcipher.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns the name of the cipher, or an empty QString if this is a null cipher.

See also isNull().
*/
func (this *QSslCipher) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcipher.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int supportedBits() const

/*
Returns the number of bits supported by the cipher.

See also usedBits().
*/
func (this *QSslCipher) SupportedBits() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher13supportedBitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslcipher.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int usedBits() const

/*
Returns the number of bits used by the cipher.

See also supportedBits().
*/
func (this *QSslCipher) UsedBits() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher8usedBitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslcipher.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString keyExchangeMethod() const

/*
Returns the cipher's key exchange method as a QString.
*/
func (this *QSslCipher) KeyExchangeMethod() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher17keyExchangeMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcipher.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString authenticationMethod() const

/*
Returns the cipher's authentication method as a QString.
*/
func (this *QSslCipher) AuthenticationMethod() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher20authenticationMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcipher.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QString encryptionMethod() const

/*
Returns the cipher's encryption method as a QString.
*/
func (this *QSslCipher) EncryptionMethod() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher16encryptionMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcipher.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString protocolString() const

/*
Returns the cipher's protocol as a QString.

See also protocol().
*/
func (this *QSslCipher) ProtocolString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher14protocolStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcipher.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol protocol() const

/*
Returns the cipher's protocol type, or QSsl::UnknownProtocol if QSslCipher is unable to determine the protocol (protocolString() may contain more information).

See also protocolString().
*/
func (this *QSslCipher) Protocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher8protocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
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
