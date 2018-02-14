package qtnetwork

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtNetwork/qsslerror.h:115
// index:36
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslError &, uint)
func QHash_36(key QSslError_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSslError_PTR() != nil {
		convArg0 = key.QSslError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK9QSslErrorj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:59
// index:37
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslDiffieHellmanParameters &, uint)
func QHash_37(dhparam QSslDiffieHellmanParameters_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if dhparam != nil && dhparam.QSslDiffieHellmanParameters_PTR() != nil {
		convArg0 = dhparam.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK27QSslDiffieHellmanParametersj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:71
// index:38
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslCertificate &, uint)
func QHash_38(key QSslCertificate_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSslCertificate_PTR() != nil {
		convArg0 = key.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK15QSslCertificatej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qhostaddress.h:69
// index:39
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QHostAddress &, uint)
func QHash_39(key QHostAddress_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QHostAddress_PTR() != nil {
		convArg0 = key.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK12QHostAddressj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:55
// index:40
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSslEllipticCurve, uint)
func QHash_40(curve QSslEllipticCurve_ITF /*123*/, seed uint) uint {
	var convArg0 unsafe.Pointer
	if curve != nil && curve.QSslEllipticCurve_PTR() != nil {
		convArg0 = curve.QSslEllipticCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash17QSslEllipticCurvej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:90
// index:41
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSslEllipticCurve, uint)
func QHash_41(curve QSslEllipticCurve_ITF /*123*/, seed uint) uint {
	var convArg0 unsafe.Pointer
	if curve != nil && curve.QSslEllipticCurve_PTR() != nil {
		convArg0 = curve.QSslEllipticCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash17QSslEllipticCurvej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

//  body block end
