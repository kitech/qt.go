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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSslEllipticCurveFromPointer(cthis unsafe.Pointer) *QSslEllipticCurve {
	return &QSslEllipticCurve{&qtrt.CObject{cthis}}
}
func (*QSslEllipticCurve) NewFromPointer(cthis unsafe.Pointer) *QSslEllipticCurve {
	return NewQSslEllipticCurveFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:59
// index:0
// Public inline
// void QSslEllipticCurve()
func NewQSslEllipticCurve() *QSslEllipticCurve {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslEllipticCurveC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslEllipticCurveFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:64
// index:0
// Public static
// QSslEllipticCurve fromShortName(const class QString &)
func (this *QSslEllipticCurve) FromShortName(name *qtcore.QString) *QSslEllipticCurve /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslEllipticCurve13fromShortNameERK7QString", ffiqt.FFI_TYPE_POINTER, name)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQSslEllipticCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSslEllipticCurve_FromShortName(name *qtcore.QString) *QSslEllipticCurve /*123*/ {
	var nilthis *QSslEllipticCurve
	rv := nilthis.FromShortName(name)
	return rv
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:65
// index:0
// Public static
// QSslEllipticCurve fromLongName(const class QString &)
func (this *QSslEllipticCurve) FromLongName(name *qtcore.QString) *QSslEllipticCurve /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSslEllipticCurve12fromLongNameERK7QString", ffiqt.FFI_TYPE_POINTER, name)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQSslEllipticCurveFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSslEllipticCurve_FromLongName(name *qtcore.QString) *QSslEllipticCurve /*123*/ {
	var nilthis *QSslEllipticCurve
	rv := nilthis.FromLongName(name)
	return rv
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:67
// index:0
// Public
// QString shortName()
func (this *QSslEllipticCurve) ShortName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslEllipticCurve9shortNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:68
// index:0
// Public
// QString longName()
func (this *QSslEllipticCurve) LongName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslEllipticCurve8longNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:70
// index:0
// Public inline
// bool isValid()
func (this *QSslEllipticCurve) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslEllipticCurve7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:75
// index:0
// Public
// bool isTlsNamedCurve()
func (this *QSslEllipticCurve) IsTlsNamedCurve() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSslEllipticCurve15isTlsNamedCurveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
