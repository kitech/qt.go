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
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

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
func NewQSslCipher_1(name string) *QSslCipher {
	var tmpArg0 = qtcore.NewQString_5(name)
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
func NewQSslCipher_2(name string, protocol int) *QSslCipher {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipherC2ERK7QStringN4QSsl11SslProtocolE", qtrt.FFI_TYPE_POINTER, convArg0, protocol)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCipher)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcipher.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslCipher()
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
func (this *QSslCipher) Swap(other QSslCipher_ITF) {
	var convArg0 = other.QSslCipher_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipher4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcipher.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QSslCipher) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcipher.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QSslCipher) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcipher.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int supportedBits()
func (this *QSslCipher) SupportedBits() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher13supportedBitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslcipher.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int usedBits()
func (this *QSslCipher) UsedBits() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher8usedBitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslcipher.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString keyExchangeMethod()
func (this *QSslCipher) KeyExchangeMethod() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher17keyExchangeMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcipher.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString authenticationMethod()
func (this *QSslCipher) AuthenticationMethod() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher20authenticationMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcipher.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QString encryptionMethod()
func (this *QSslCipher) EncryptionMethod() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher16encryptionMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcipher.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString protocolString()
func (this *QSslCipher) ProtocolString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher14protocolStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcipher.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol protocol()
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
