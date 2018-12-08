package qtnetwork

// /usr/include/qt/QtNetwork/qsslkey.h
// #include <qsslkey.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QSslKey struct {
	*qtrt.CObject
}
type QSslKey_ITF interface {
	QSslKey_PTR() *QSslKey
}

func (ptr *QSslKey) QSslKey_PTR() *QSslKey { return ptr }

func (this *QSslKey) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslKey) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslKeyFromPointer(cthis unsafe.Pointer) *QSslKey {
	return &QSslKey{&qtrt.CObject{cthis}}
}
func (*QSslKey) NewFromPointer(cthis unsafe.Pointer) *QSslKey {
	return NewQSslKeyFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslkey.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslKey()

/*
Constructs a null key.

See also isNull().
*/
func (*QSslKey) NewForInherit() *QSslKey {
	return NewQSslKey()
}
func NewQSslKey() *QSslKey {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslKey(const QByteArray &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)

/*
Constructs a null key.

See also isNull().
*/
func (*QSslKey) NewForInherit1(encoded qtcore.QByteArray_ITF, algorithm int, format int, type_ int, passPhrase qtcore.QByteArray_ITF) *QSslKey {
	return NewQSslKey1(encoded, algorithm, format, type_, passPhrase)
}
func NewQSslKey1(encoded qtcore.QByteArray_ITF, algorithm int, format int, type_ int, passPhrase qtcore.QByteArray_ITF) *QSslKey {
	var convArg0 unsafe.Pointer
	if encoded != nil && encoded.QByteArray_PTR() != nil {
		convArg0 = encoded.QByteArray_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if passPhrase != nil && passPhrase.QByteArray_PTR() != nil {
		convArg4 = passPhrase.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2ERK10QByteArrayN4QSsl12KeyAlgorithmENS3_14EncodingFormatENS3_7KeyTypeES2_", qtrt.FFI_TYPE_POINTER, convArg0, algorithm, format, type_, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslKey(const QByteArray &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)

/*
Constructs a null key.

See also isNull().
*/
func (*QSslKey) NewForInherit1p(encoded qtcore.QByteArray_ITF, algorithm int) *QSslKey {
	return NewQSslKey1p(encoded, algorithm)
}
func NewQSslKey1p(encoded qtcore.QByteArray_ITF, algorithm int) *QSslKey {
	var convArg0 unsafe.Pointer
	if encoded != nil && encoded.QByteArray_PTR() != nil {
		convArg0 = encoded.QByteArray_PTR().GetCthis()
	}
	// arg: 2, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum, , Invalid
	format := 0
	// arg: 3, QSsl::KeyType=Elaborated, QSsl::KeyType=Enum, , Invalid
	type_ := 0
	// arg: 4, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg4 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2ERK10QByteArrayN4QSsl12KeyAlgorithmENS3_14EncodingFormatENS3_7KeyTypeES2_", qtrt.FFI_TYPE_POINTER, convArg0, algorithm, format, type_, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslKey(const QByteArray &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)

/*
Constructs a null key.

See also isNull().
*/
func (*QSslKey) NewForInherit1p1(encoded qtcore.QByteArray_ITF, algorithm int, format int) *QSslKey {
	return NewQSslKey1p1(encoded, algorithm, format)
}
func NewQSslKey1p1(encoded qtcore.QByteArray_ITF, algorithm int, format int) *QSslKey {
	var convArg0 unsafe.Pointer
	if encoded != nil && encoded.QByteArray_PTR() != nil {
		convArg0 = encoded.QByteArray_PTR().GetCthis()
	}
	// arg: 3, QSsl::KeyType=Elaborated, QSsl::KeyType=Enum, , Invalid
	type_ := 0
	// arg: 4, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg4 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2ERK10QByteArrayN4QSsl12KeyAlgorithmENS3_14EncodingFormatENS3_7KeyTypeES2_", qtrt.FFI_TYPE_POINTER, convArg0, algorithm, format, type_, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslKey(const QByteArray &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)

/*
Constructs a null key.

See also isNull().
*/
func (*QSslKey) NewForInherit1p2(encoded qtcore.QByteArray_ITF, algorithm int, format int, type_ int) *QSslKey {
	return NewQSslKey1p2(encoded, algorithm, format, type_)
}
func NewQSslKey1p2(encoded qtcore.QByteArray_ITF, algorithm int, format int, type_ int) *QSslKey {
	var convArg0 unsafe.Pointer
	if encoded != nil && encoded.QByteArray_PTR() != nil {
		convArg0 = encoded.QByteArray_PTR().GetCthis()
	}
	// arg: 4, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg4 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2ERK10QByteArrayN4QSsl12KeyAlgorithmENS3_14EncodingFormatENS3_7KeyTypeES2_", qtrt.FFI_TYPE_POINTER, convArg0, algorithm, format, type_, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSslKey(QIODevice *, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)

/*
Constructs a null key.

See also isNull().
*/
func (*QSslKey) NewForInherit2(device qtcore.QIODevice_ITF /*777 QIODevice **/, algorithm int, format int, type_ int, passPhrase qtcore.QByteArray_ITF) *QSslKey {
	return NewQSslKey2(device, algorithm, format, type_, passPhrase)
}
func NewQSslKey2(device qtcore.QIODevice_ITF /*777 QIODevice **/, algorithm int, format int, type_ int, passPhrase qtcore.QByteArray_ITF) *QSslKey {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if passPhrase != nil && passPhrase.QByteArray_PTR() != nil {
		convArg4 = passPhrase.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2EP9QIODeviceN4QSsl12KeyAlgorithmENS2_14EncodingFormatENS2_7KeyTypeERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, algorithm, format, type_, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSslKey(QIODevice *, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)

/*
Constructs a null key.

See also isNull().
*/
func (*QSslKey) NewForInherit2p(device qtcore.QIODevice_ITF /*777 QIODevice **/, algorithm int) *QSslKey {
	return NewQSslKey2p(device, algorithm)
}
func NewQSslKey2p(device qtcore.QIODevice_ITF /*777 QIODevice **/, algorithm int) *QSslKey {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 2, QSsl::EncodingFormat=Elaborated, QSsl::EncodingFormat=Enum, , Invalid
	format := 0
	// arg: 3, QSsl::KeyType=Elaborated, QSsl::KeyType=Enum, , Invalid
	type_ := 0
	// arg: 4, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg4 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2EP9QIODeviceN4QSsl12KeyAlgorithmENS2_14EncodingFormatENS2_7KeyTypeERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, algorithm, format, type_, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSslKey(QIODevice *, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)

/*
Constructs a null key.

See also isNull().
*/
func (*QSslKey) NewForInherit2p1(device qtcore.QIODevice_ITF /*777 QIODevice **/, algorithm int, format int) *QSslKey {
	return NewQSslKey2p1(device, algorithm, format)
}
func NewQSslKey2p1(device qtcore.QIODevice_ITF /*777 QIODevice **/, algorithm int, format int) *QSslKey {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 3, QSsl::KeyType=Elaborated, QSsl::KeyType=Enum, , Invalid
	type_ := 0
	// arg: 4, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg4 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2EP9QIODeviceN4QSsl12KeyAlgorithmENS2_14EncodingFormatENS2_7KeyTypeERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, algorithm, format, type_, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSslKey(QIODevice *, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)

/*
Constructs a null key.

See also isNull().
*/
func (*QSslKey) NewForInherit2p2(device qtcore.QIODevice_ITF /*777 QIODevice **/, algorithm int, format int, type_ int) *QSslKey {
	return NewQSslKey2p2(device, algorithm, format, type_)
}
func NewQSslKey2p2(device qtcore.QIODevice_ITF /*777 QIODevice **/, algorithm int, format int, type_ int) *QSslKey {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	// arg: 4, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg4 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2EP9QIODeviceN4QSsl12KeyAlgorithmENS2_14EncodingFormatENS2_7KeyTypeERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, algorithm, format, type_, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSslKey & operator=(QSslKey &&)

/*

 */
func (this *QSslKey) Operator_equal(other unsafe.Pointer /*333*/) *QSslKey {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslkey.h:77
// index:1
// Public Visibility=Default Availability=Available
// [8] QSslKey & operator=(const QSslKey &)

/*

 */
func (this *QSslKey) Operator_equal1(other QSslKey_ITF) *QSslKey {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslKey_PTR() != nil {
		convArg0 = other.QSslKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslKey)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslkey.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslKey()

/*

 */
func DeleteQSslKey(this *QSslKey) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslkey.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslKey &)

/*
Swaps this ssl key with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QSslKey) Swap(other QSslKey_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslKey_PTR() != nil {
		convArg0 = other.QSslKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKey4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this is a null key; otherwise false.

See also clear().
*/
func (this *QSslKey) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslkey.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the contents of this key, making it a null key.

See also isNull().
*/
func (this *QSslKey) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKey5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int length() const

/*
Returns the length of the key in bits, or -1 if the key is null.
*/
func (this *QSslKey) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslkey.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::KeyType type() const

/*
Returns the type of the key (i.e., PublicKey or PrivateKey).
*/
func (this *QSslKey) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::KeyAlgorithm algorithm() const

/*
Returns the key algorithm.
*/
func (this *QSslKey) Algorithm() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey9algorithmEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toPem(const QByteArray &) const

/*
Returns the key in PEM encoding. The result is encrypted with passPhrase if the key is a private key and passPhrase is non-empty.
*/
func (this *QSslKey) ToPem(passPhrase qtcore.QByteArray_ITF) *qtcore.QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if passPhrase != nil && passPhrase.QByteArray_PTR() != nil {
		convArg0 = passPhrase.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey5toPemERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslkey.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toPem(const QByteArray &) const

/*
Returns the key in PEM encoding. The result is encrypted with passPhrase if the key is a private key and passPhrase is non-empty.
*/
func (this *QSslKey) ToPemp() *qtcore.QByteArray /*123*/ {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg0 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey5toPemERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslkey.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toDer(const QByteArray &) const

/*
Returns the key in DER encoding.

The passPhrase argument should be omitted as DER cannot be encrypted. It will be removed in a future version of Qt.
*/
func (this *QSslKey) ToDer(passPhrase qtcore.QByteArray_ITF) *qtcore.QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if passPhrase != nil && passPhrase.QByteArray_PTR() != nil {
		convArg0 = passPhrase.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey5toDerERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslkey.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toDer(const QByteArray &) const

/*
Returns the key in DER encoding.

The passPhrase argument should be omitted as DER cannot be encrypted. It will be removed in a future version of Qt.
*/
func (this *QSslKey) ToDerp() *qtcore.QByteArray /*123*/ {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg0 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey5toDerERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslkey.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] Qt::HANDLE handle() const

/*
Returns a pointer to the native key handle, if it is available; otherwise a null pointer is returned.

You can use this handle together with the native API to access extended information about the key.

Warning: Use of this function has a high probability of being non-portable, and its return value may vary across platforms, and between minor Qt releases.
*/
func (this *QSslKey) Handle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QSslKey &) const

/*

 */
func (this *QSslKey) Operator_equal_equal(key QSslKey_ITF) bool {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSslKey_PTR() != nil {
		convArg0 = key.QSslKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKeyeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslkey.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QSslKey &) const

/*

 */
func (this *QSslKey) Operator_not_equal(key QSslKey_ITF) bool {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSslKey_PTR() != nil {
		convArg0 = key.QSslKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKeyneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
