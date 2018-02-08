package qtnetwork

// /usr/include/qt/QtNetwork/qsslkey.h
// #include <qsslkey.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

type QSslKey struct {
	*qtrt.CObject
}

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
func NewQSslKey() *QSslKey {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslKey(const QByteArray &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)
func NewQSslKey_1(encoded *qtcore.QByteArray, algorithm int, format int, type_ int, passPhrase *qtcore.QByteArray) *QSslKey {
	var convArg0 = encoded.GetCthis()
	var convArg4 = passPhrase.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2ERK10QByteArrayN4QSsl12KeyAlgorithmENS3_14EncodingFormatENS3_7KeyTypeES2_", qtrt.FFI_TYPE_POINTER, convArg0, algorithm, format, type_, convArg4)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSslKey(QIODevice *, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)
func NewQSslKey_2(device *qtcore.QIODevice /*777 QIODevice **/, algorithm int, format int, type_ int, passPhrase *qtcore.QByteArray) *QSslKey {
	var convArg0 = device.GetCthis()
	var convArg4 = passPhrase.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyC2EP9QIODeviceN4QSsl12KeyAlgorithmENS2_14EncodingFormatENS2_7KeyTypeERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, algorithm, format, type_, convArg4)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslKey)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslKey()
func DeleteQSslKey(this *QSslKey) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKeyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslkey.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslKey &)
func (this *QSslKey) Swap(other *QSslKey) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKey4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QSslKey) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslkey.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QSslKey) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSslKey5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int length()
func (this *QSslKey) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslkey.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::KeyType type()
func (this *QSslKey) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::KeyAlgorithm algorithm()
func (this *QSslKey) Algorithm() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey9algorithmEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toPem(const QByteArray &)
func (this *QSslKey) ToPem(passPhrase *qtcore.QByteArray) *qtcore.QByteArray /*123*/ {
	var convArg0 = passPhrase.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey5toPemERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslkey.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray toDer(const QByteArray &)
func (this *QSslKey) ToDer(passPhrase *qtcore.QByteArray) *qtcore.QByteArray /*123*/ {
	var convArg0 = passPhrase.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey5toDerERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslkey.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] Qt::HANDLE handle()
func (this *QSslKey) Handle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSslKey6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

//  body block end
