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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSslKeyFromPointer(cthis unsafe.Pointer) *QSslKey {
	return &QSslKey{&qtrt.CObject{cthis}}
}
func (*QSslKey) NewFromPointer(cthis unsafe.Pointer) *QSslKey {
	return NewQSslKeyFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslkey.h:63
// index:0
// Public
// void QSslKey()
func NewQSslKey() *QSslKey {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSslKeyC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:64
// index:1
// Public
// void QSslKey(const QByteArray &, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)
func NewQSslKey_1(encoded *qtcore.QByteArray, algorithm int, format int, type_ int, passPhrase *qtcore.QByteArray) *QSslKey {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = encoded.GetCthis()
	var convArg4 = passPhrase.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSslKeyC2ERK10QByteArrayN4QSsl12KeyAlgorithmENS3_14EncodingFormatENS3_7KeyTypeES2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, algorithm, format, type_, convArg4)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:68
// index:2
// Public
// void QSslKey(QIODevice *, QSsl::KeyAlgorithm, QSsl::EncodingFormat, QSsl::KeyType, const QByteArray &)
func NewQSslKey_2(device *qtcore.QIODevice /*777 QIODevice **/, algorithm int, format int, type_ int, passPhrase *qtcore.QByteArray) *QSslKey {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = device.GetCthis()
	var convArg4 = passPhrase.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSslKeyC2EP9QIODeviceN4QSsl12KeyAlgorithmENS2_14EncodingFormatENS2_7KeyTypeERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0, algorithm, format, type_, convArg4)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslKeyFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslkey.h:78
// index:0
// Public
// void ~QSslKey()
func DeleteQSslKey(*QSslKey) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSslKeyD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:80
// index:0
// Public inline
// void swap(QSslKey &)
func (this *QSslKey) Swap(other *QSslKey) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSslKey4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:82
// index:0
// Public
// bool isNull()
func (this *QSslKey) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSslKey6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslkey.h:83
// index:0
// Public
// void clear()
func (this *QSslKey) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSslKey5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:85
// index:0
// Public
// int length()
func (this *QSslKey) Length() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSslKey6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtNetwork/qsslkey.h:86
// index:0
// Public
// QSsl::KeyType type()
func (this *QSslKey) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSslKey4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:87
// index:0
// Public
// QSsl::KeyAlgorithm algorithm()
func (this *QSslKey) Algorithm() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSslKey9algorithmEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:89
// index:0
// Public
// QByteArray toPem(const QByteArray &)
func (this *QSslKey) ToPem(passPhrase *qtcore.QByteArray) *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = passPhrase.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSslKey5toPemERK10QByteArray", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslkey.h:90
// index:0
// Public
// QByteArray toDer(const QByteArray &)
func (this *QSslKey) ToDer(passPhrase *qtcore.QByteArray) *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = passPhrase.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSslKey5toDerERK10QByteArray", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslkey.h:92
// index:0
// Public
// Qt::HANDLE handle()
func (this *QSslKey) Handle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSslKey6handleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
