package qtnetwork

// /usr/include/qt/QtNetwork/qsslellipticcurve.h
// #include <qsslellipticcurve.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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

type QSslEllipticCurve struct {
	*qtrt.CObject
}

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
func NewQSslEllipticCurve() *QSslEllipticCurve {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslEllipticCurveC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslEllipticCurveFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslEllipticCurve)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:64
// index:0
// Public static Visibility=Default Availability=Available
// [4] QSslEllipticCurve fromShortName(const QString &)
func (this *QSslEllipticCurve) FromShortName(name string) *QSslEllipticCurve /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslEllipticCurve13fromShortNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
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
func (this *QSslEllipticCurve) FromLongName(name string) *QSslEllipticCurve /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslEllipticCurve12fromLongNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
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
// [8] QString shortName()
func (this *QSslEllipticCurve) ShortName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslEllipticCurve9shortNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString longName()
func (this *QSslEllipticCurve) LongName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslEllipticCurve8longNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QSslEllipticCurve) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslEllipticCurve7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTlsNamedCurve()
func (this *QSslEllipticCurve) IsTlsNamedCurve() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslEllipticCurve15isTlsNamedCurveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQSslEllipticCurve(this *QSslEllipticCurve) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslEllipticCurveD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
