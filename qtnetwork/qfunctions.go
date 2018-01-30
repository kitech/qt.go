package qtnetwork

import "unsafe"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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
// /usr/include/qt/QtNetwork/qhostaddress.h:69
// index:0
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QHostAddress &, uint)
func QHash(key *QHostAddress, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK12QHostAddressj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:71
// index:1
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslCertificate &, uint)
func QHash_1(key *QSslCertificate, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK15QSslCertificatej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslerror.h:115
// index:2
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslError &, uint)
func QHash_2(key *QSslError, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK9QSslErrorj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:59
// index:3
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslDiffieHellmanParameters &, uint)
func QHash_3(dhparam *QSslDiffieHellmanParameters, seed uint) uint {
	var convArg0 = dhparam.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK27QSslDiffieHellmanParametersj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:55
// index:4
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSslEllipticCurve, uint)
func QHash_4(curve *QSslEllipticCurve /*123*/, seed uint) uint {
	var convArg0 = curve.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHash17QSslEllipticCurvej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:90
// index:5
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSslEllipticCurve, uint)
func QHash_5(curve *QSslEllipticCurve /*123*/, seed uint) uint {
	var convArg0 = curve.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHash17QSslEllipticCurvej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

//  body block end
