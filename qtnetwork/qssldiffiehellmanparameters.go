package qtnetwork

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h
// #include <qssldiffiehellmanparameters.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
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
type QSslDiffieHellmanParameters struct {
	*qtrt.CObject
}

func (this *QSslDiffieHellmanParameters) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslDiffieHellmanParameters) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSslDiffieHellmanParametersFromPointer(cthis unsafe.Pointer) *QSslDiffieHellmanParameters {
	return &QSslDiffieHellmanParameters{&qtrt.CObject{cthis}}
}
func (*QSslDiffieHellmanParameters) NewFromPointer(cthis unsafe.Pointer) *QSslDiffieHellmanParameters {
	return NewQSslDiffieHellmanParametersFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:82
// index:0
// Public static
// QSslDiffieHellmanParameters defaultParameters()
func (this *QSslDiffieHellmanParameters) DefaultParameters() *QSslDiffieHellmanParameters /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParameters17defaultParametersEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSslDiffieHellmanParameters_DefaultParameters() *QSslDiffieHellmanParameters /*123*/ {
	var nilthis *QSslDiffieHellmanParameters
	rv := nilthis.DefaultParameters()
	return rv
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:84
// index:0
// Public
// void QSslDiffieHellmanParameters()
func NewQSslDiffieHellmanParameters() *QSslDiffieHellmanParameters {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParametersC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslDiffieHellmanParametersFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:87
// index:0
// Public
// void ~QSslDiffieHellmanParameters()
func DeleteQSslDiffieHellmanParameters(*QSslDiffieHellmanParameters) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParametersD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:92
// index:0
// Public inline
// void swap(class QSslDiffieHellmanParameters &)
func (this *QSslDiffieHellmanParameters) Swap(other *QSslDiffieHellmanParameters) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParameters4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:94
// index:0
// Public static
// QSslDiffieHellmanParameters fromEncoded(const class QByteArray &, QSsl::EncodingFormat)
func (this *QSslDiffieHellmanParameters) FromEncoded(encoded *qtcore.QByteArray, format int) *QSslDiffieHellmanParameters /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParameters11fromEncodedERK10QByteArrayN4QSsl14EncodingFormatE", ffiqt.FFI_TYPE_POINTER, encoded, format)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSslDiffieHellmanParameters_FromEncoded(encoded *qtcore.QByteArray, format int) *QSslDiffieHellmanParameters /*123*/ {
	var nilthis *QSslDiffieHellmanParameters
	rv := nilthis.FromEncoded(encoded, format)
	return rv
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:95
// index:1
// Public static
// QSslDiffieHellmanParameters fromEncoded(class QIODevice *, QSsl::EncodingFormat)
func (this *QSslDiffieHellmanParameters) FromEncoded_1(device *qtcore.QIODevice /*777 QIODevice **/, format int) *QSslDiffieHellmanParameters /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QSslDiffieHellmanParameters11fromEncodedEP9QIODeviceN4QSsl14EncodingFormatE", ffiqt.FFI_TYPE_POINTER, device, format)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQSslDiffieHellmanParametersFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSslDiffieHellmanParameters_FromEncoded_1(device *qtcore.QIODevice /*777 QIODevice **/, format int) *QSslDiffieHellmanParameters /*123*/ {
	var nilthis *QSslDiffieHellmanParameters
	rv := nilthis.FromEncoded_1(device, format)
	return rv
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:97
// index:0
// Public
// bool isEmpty()
func (this *QSslDiffieHellmanParameters) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QSslDiffieHellmanParameters7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:98
// index:0
// Public
// bool isValid()
func (this *QSslDiffieHellmanParameters) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QSslDiffieHellmanParameters7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:99
// index:0
// Public
// QSslDiffieHellmanParameters::Error error()
func (this *QSslDiffieHellmanParameters) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QSslDiffieHellmanParameters5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:100
// index:0
// Public
// QString errorString()
func (this *QSslDiffieHellmanParameters) ErrorString() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QSslDiffieHellmanParameters11errorStringEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QSslDiffieHellmanParameters__Error = int

const QSslDiffieHellmanParameters__NoError QSslDiffieHellmanParameters__Error = 0
const QSslDiffieHellmanParameters__InvalidInputDataError QSslDiffieHellmanParameters__Error = 1
const QSslDiffieHellmanParameters__UnsafeParametersError QSslDiffieHellmanParameters__Error = 2

//  body block end
