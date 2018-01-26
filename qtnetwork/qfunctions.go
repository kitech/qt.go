package qtnetwork

import "unsafe"
import "mkuse/cffiqt"
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
// index:37
// Invalid
// uint qHash(const class QHostAddress &, uint)
func QHash_37(key *QHostAddress, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK12QHostAddressj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:71
// index:38
// Invalid
// uint qHash(const class QSslCertificate &, uint)
func QHash_38(key *QSslCertificate, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK15QSslCertificatej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslerror.h:115
// index:39
// Invalid
// uint qHash(const class QSslError &, uint)
func QHash_39(key *QSslError, seed uint) uint {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK9QSslErrorj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:59
// index:40
// Invalid
// uint qHash(const class QSslDiffieHellmanParameters &, uint)
func QHash_40(dhparam *QSslDiffieHellmanParameters, seed uint) uint {
	var convArg0 = dhparam.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHashRK27QSslDiffieHellmanParametersj", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:55
// index:41
// Invalid inline
// uint qHash(class QSslEllipticCurve, uint)
func QHash_41(curve *QSslEllipticCurve /*123*/, seed uint) uint {
	var convArg0 = curve.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHash17QSslEllipticCurvej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:90
// index:42
// Invalid inline
// uint qHash(class QSslEllipticCurve, uint)
func QHash_42(curve *QSslEllipticCurve /*123*/, seed uint) uint {
	var convArg0 = curve.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_Z5qHash17QSslEllipticCurvej", ffiqt.FFI_TYPE_POINTER, convArg0, seed)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}

//  body block end
