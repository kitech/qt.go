package qtnetwork

import "unsafe"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
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
func QHash_36(key *QSslError, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK9QSslErrorj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:59
// index:37
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslDiffieHellmanParameters &, uint)
func QHash_37(dhparam *QSslDiffieHellmanParameters, seed uint) uint {
	var convArg0 = dhparam.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK27QSslDiffieHellmanParametersj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:71
// index:38
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslCertificate &, uint)
func QHash_38(key *QSslCertificate, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK15QSslCertificatej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qhostaddress.h:69
// index:39
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QHostAddress &, uint)
func QHash_39(key *QHostAddress, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK12QHostAddressj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:55
// index:40
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSslEllipticCurve, uint)
func QHash_40(curve *QSslEllipticCurve /*123*/, seed uint) uint {
	var convArg0 = curve.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash17QSslEllipticCurvej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:90
// index:41
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSslEllipticCurve, uint)
func QHash_41(curve *QSslEllipticCurve /*123*/, seed uint) uint {
	var convArg0 = curve.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash17QSslEllipticCurvej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	return uint(rv) // 222
}

//  body block end
