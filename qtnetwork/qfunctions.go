package qtnetwork

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

func init_unused_10011() {
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
// /usr/include/qt/QtNetwork/qssl.h:118
// index:78
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSsl::SslOptions::enum_type, int)

/*

 */
func Operator_or78(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN4QSsl9SslOptionEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkconfigmanager.h:94
// index:79
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QNetworkConfigurationManager::Capabilities::enum_type, int)

/*

 */
func Operator_or79(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN28QNetworkConfigurationManager10CapabilityEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:178
// index:80
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QNetworkInterface::InterfaceFlags::enum_type, int)

/*

 */
func Operator_or80(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN17QNetworkInterface13InterfaceFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:239
// index:81
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QAbstractSocket::PauseModes::enum_type, int)

/*

 */
func Operator_or81(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN15QAbstractSocket9PauseModeEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractsocket.h:238
// index:82
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QAbstractSocket::BindMode::enum_type, int)

/*

 */
func Operator_or82(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN15QAbstractSocket8BindFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:214
// index:83
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QNetworkProxy::Capabilities::enum_type, int)

/*

 */
func Operator_or83(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN13QNetworkProxy10CapabilityEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtNetwork/qhostaddress.h:165
// index:85
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QHostAddress::ConversionMode::enum_type, int)

/*

 */
func Operator_or85(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN12QHostAddress18ConversionModeFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:90
// index:86
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QHstsPolicy::PolicyFlags::enum_type, int)

/*

 */
func Operator_or86(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN11QHstsPolicy10PolicyFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:68
// index:65
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(const QSslDiffieHellmanParameters &, const QSslDiffieHellmanParameters &)

/*

 */
func Operator_not_equal65(lhs QSslDiffieHellmanParameters_ITF, rhs QSslDiffieHellmanParameters_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QSslDiffieHellmanParameters_PTR() != nil {
		convArg0 = lhs.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QSslDiffieHellmanParameters_PTR() != nil {
		convArg1 = rhs.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK27QSslDiffieHellmanParametersS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:94
// index:66
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(const QHstsPolicy &, const QHstsPolicy &)

/*

 */
func Operator_not_equal66(lhs QHstsPolicy_ITF, rhs QHstsPolicy_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QHstsPolicy_PTR() != nil {
		convArg0 = lhs.QHstsPolicy_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QHstsPolicy_PTR() != nil {
		convArg1 = rhs.QHstsPolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK11QHstsPolicyS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:170
// index:67
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(QHostAddress::SpecialAddress, const QHostAddress &)

/*

 */
func Operator_not_equal67(lhs int, rhs QHostAddress_ITF) bool {
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QHostAddress_PTR() != nil {
		convArg1 = rhs.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneN12QHostAddress14SpecialAddressERKS_", qtrt.FFI_TYPE_POINTER, lhs, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:96
// index:68
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(QSslEllipticCurve, QSslEllipticCurve)

/*

 */
func Operator_not_equal68(lhs QSslEllipticCurve_ITF /*123*/, rhs QSslEllipticCurve_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QSslEllipticCurve_PTR() != nil {
		convArg0 = lhs.QSslEllipticCurve_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QSslEllipticCurve_PTR() != nil {
		convArg1 = rhs.QSslEllipticCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Zne17QSslEllipticCurveS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:66
// index:65
// Invalid Visibility=Default Availability=Available
// [1] bool operator==(const QSslDiffieHellmanParameters &, const QSslDiffieHellmanParameters &)

/*

 */
func Operator_equal_equal65(lhs QSslDiffieHellmanParameters_ITF, rhs QSslDiffieHellmanParameters_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QSslDiffieHellmanParameters_PTR() != nil {
		convArg0 = lhs.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QSslDiffieHellmanParameters_PTR() != nil {
		convArg1 = rhs.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqRK27QSslDiffieHellmanParametersS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:92
// index:66
// Invalid Visibility=Default Availability=Available
// [1] bool operator==(const QHstsPolicy &, const QHstsPolicy &)

/*

 */
func Operator_equal_equal66(lhs QHstsPolicy_ITF, rhs QHstsPolicy_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QHstsPolicy_PTR() != nil {
		convArg0 = lhs.QHstsPolicy_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QHstsPolicy_PTR() != nil {
		convArg1 = rhs.QHstsPolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqRK11QHstsPolicyS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhostaddress.h:168
// index:67
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator==(QHostAddress::SpecialAddress, const QHostAddress &)

/*

 */
func Operator_equal_equal67(address1 int, address2 QHostAddress_ITF) bool {
	var convArg1 unsafe.Pointer
	if address2 != nil && address2.QHostAddress_PTR() != nil {
		convArg1 = address2.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqN12QHostAddress14SpecialAddressERKS_", qtrt.FFI_TYPE_POINTER, address1, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:93
// index:68
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator==(QSslEllipticCurve, QSslEllipticCurve)

/*

 */
func Operator_equal_equal68(lhs QSslEllipticCurve_ITF /*123*/, rhs QSslEllipticCurve_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QSslEllipticCurve_PTR() != nil {
		convArg0 = lhs.QSslEllipticCurve_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QSslEllipticCurve_PTR() != nil {
		convArg1 = rhs.QSslEllipticCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Zeq17QSslEllipticCurveS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qpassworddigestor.h:53
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QByteArray deriveKeyPbkdf2(QCryptographicHash::Algorithm, const QByteArray &, const QByteArray &, int, quint64)

/*
Derive a key using the PBKDF2-algorithm as defined in RFC 8018.

This function takes the data and salt, and then applies HMAC-X, where the X is algorithm, repeatedly. It internally concatenates intermediate results to the final output until at least dkLen amount of bytes have been computed and it will execute HMAC-X iterations times each time a concatenation is required. The total number of times it will execute HMAC-X depends on iterations, dkLen and algorithm and can be calculated as iterations * ceil(dkLen / QCryptographicHash::hashLength(algorithm)).

This function was introduced in  Qt 5.12.

See also deriveKeyPbkdf1, QMessageAuthenticationCode, and QCryptographicHash.
*/
func DeriveKeyPbkdf2(algorithm int, password qtcore.QByteArray_ITF, salt qtcore.QByteArray_ITF, iterations int, dkLen uint64) *qtcore.QByteArray /*123*/ {
	var convArg1 unsafe.Pointer
	if password != nil && password.QByteArray_PTR() != nil {
		convArg1 = password.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if salt != nil && salt.QByteArray_PTR() != nil {
		convArg2 = salt.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPasswordDigestor15deriveKeyPbkdf2EN18QCryptographicHash9AlgorithmERK10QByteArrayS4_iy", qtrt.FFI_TYPE_POINTER, algorithm, convArg1, convArg2, iterations, dkLen)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qpassworddigestor.h:50
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QByteArray deriveKeyPbkdf1(QCryptographicHash::Algorithm, const QByteArray &, const QByteArray &, int, quint64)

/*
Returns a hash computed using the PBKDF1-algorithm as defined in RFC 8018.

The function takes the data and salt, and then hashes it repeatedly for iterations iterations using the specified hash algorithm. If the resulting hash is longer than dkLen then it is truncated before it is returned.

This function only supports SHA-1 and MD5! The max output size is 160 bits (20 bytes) when using SHA-1, or 128 bits (16 bytes) when using MD5. Specifying a value for dkLen which is greater than this results in a warning and an empty QByteArray is returned. To programmatically check this limit you can use QCryptographicHash::hashLength. Furthermore: the salt must always be 8 bytes long!

Note: This function is provided for use with legacy applications and all new applications are recommended to use PBKDF2.

This function was introduced in  Qt 5.12.

See also deriveKeyPbkdf2, QCryptographicHash, and QCryptographicHash::hashLength.
*/
func DeriveKeyPbkdf1(algorithm int, password qtcore.QByteArray_ITF, salt qtcore.QByteArray_ITF, iterations int, dkLen uint64) *qtcore.QByteArray /*123*/ {
	var convArg1 unsafe.Pointer
	if password != nil && password.QByteArray_PTR() != nil {
		convArg1 = password.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if salt != nil && salt.QByteArray_PTR() != nil {
		convArg2 = salt.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPasswordDigestor15deriveKeyPbkdf1EN18QCryptographicHash9AlgorithmERK10QByteArrayS4_iy", qtrt.FFI_TYPE_POINTER, algorithm, convArg1, convArg2, iterations, dkLen)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslerror.h:115
// index:45
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslError &, uint)

/*

 */
func QHash45(key QSslError_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSslError_PTR() != nil {
		convArg0 = key.QSslError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK9QSslErrorj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:59
// index:46
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslDiffieHellmanParameters &, uint)

/*
Returns an hash value for dhparam, using seed to seed the calculation.

This function was introduced in  Qt 5.8.
*/
func QHash46(dhparam QSslDiffieHellmanParameters_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if dhparam != nil && dhparam.QSslDiffieHellmanParameters_PTR() != nil {
		convArg0 = dhparam.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK27QSslDiffieHellmanParametersj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:69
// index:47
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QSslCertificate &, uint)

/*

 */
func QHash47(key QSslCertificate_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QSslCertificate_PTR() != nil {
		convArg0 = key.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK15QSslCertificatej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qhostaddress.h:69
// index:48
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QHostAddress &, uint)

/*
Returns a hash of the host address key, using seed to seed the calculation.

This function was introduced in  Qt 5.0.
*/
func QHash48(key QHostAddress_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QHostAddress_PTR() != nil {
		convArg0 = key.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK12QHostAddressj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:55
// index:49
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSslEllipticCurve, uint)

/*

 */
func QHash49(curve QSslEllipticCurve_ITF /*123*/, seed uint) uint {
	var convArg0 unsafe.Pointer
	if curve != nil && curve.QSslEllipticCurve_PTR() != nil {
		convArg0 = curve.QSslEllipticCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash17QSslEllipticCurvej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslellipticcurve.h:90
// index:50
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(QSslEllipticCurve, uint)

/*

 */
func QHash50(curve QSslEllipticCurve_ITF /*123*/, seed uint) uint {
	var convArg0 unsafe.Pointer
	if curve != nil && curve.QSslEllipticCurve_PTR() != nil {
		convArg0 = curve.QSslEllipticCurve_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash17QSslEllipticCurvej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtNetwork/qsslerror.h:113
// index:62
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QSslError &, QSslError &)

/*
Swaps this error instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap62(value1 QSslError_ITF, value2 QSslError_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QSslError_PTR() != nil {
		convArg0 = value1.QSslError_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QSslError_PTR() != nil {
		convArg1 = value2.QSslError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR9QSslErrorS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:160
// index:64
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QHostInfo &, QHostInfo &)

/*
Swaps host-info other with this host-info. This operation is very fast and never fails.

This function was introduced in  Qt 5.10.
*/
func Swap64(value1 QHostInfo_ITF, value2 QHostInfo_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QHostInfo_PTR() != nil {
		convArg0 = value1.QHostInfo_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QHostInfo_PTR() != nil {
		convArg1 = value2.QHostInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR9QHostInfoS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslkey.h:103
// index:65
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QSslKey &, QSslKey &)

/*
Swaps this ssl key with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap65(value1 QSslKey_ITF, value2 QSslKey_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QSslKey_PTR() != nil {
		convArg0 = value1.QSslKey_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QSslKey_PTR() != nil {
		convArg1 = value2.QSslKey_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR7QSslKeyS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qssldiffiehellmanparameters.h:112
// index:67
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QSslDiffieHellmanParameters &, QSslDiffieHellmanParameters &)

/*
Swaps this QSslDiffieHellmanParameters with other. This function is very fast and never fails.
*/
func Swap67(value1 QSslDiffieHellmanParameters_ITF, value2 QSslDiffieHellmanParameters_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QSslDiffieHellmanParameters_PTR() != nil {
		convArg0 = value1.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QSslDiffieHellmanParameters_PTR() != nil {
		convArg1 = value2.QSslDiffieHellmanParameters_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR27QSslDiffieHellmanParametersS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:78
// index:68
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QSslCertificateExtension &, QSslCertificateExtension &)

/*
Swaps this certificate extension instance with other. This function is very fast and never fails.
*/
func Swap68(value1 QSslCertificateExtension_ITF, value2 QSslCertificateExtension_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QSslCertificateExtension_PTR() != nil {
		convArg0 = value1.QSslCertificateExtension_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QSslCertificateExtension_PTR() != nil {
		convArg1 = value2.QSslCertificateExtension_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR24QSslCertificateExtensionS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h:134
// index:70
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkConfiguration &, QNetworkConfiguration &)

/*
Swaps this network configuration with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap70(value1 QNetworkConfiguration_ITF, value2 QNetworkConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QNetworkConfiguration_PTR() != nil {
		convArg0 = value1.QNetworkConfiguration_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QNetworkConfiguration_PTR() != nil {
		convArg1 = value2.QNetworkConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR21QNetworkConfigurationS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:107
// index:71
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkCacheMetaData &, QNetworkCacheMetaData &)

/*

 */
func Swap71(value1 QNetworkCacheMetaData_ITF, value2 QNetworkCacheMetaData_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QNetworkCacheMetaData_PTR() != nil {
		convArg0 = value1.QNetworkCacheMetaData_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QNetworkCacheMetaData_PTR() != nil {
		convArg1 = value2.QNetworkCacheMetaData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR21QNetworkCacheMetaDataS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:105
// index:73
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkAddressEntry &, QNetworkAddressEntry &)

/*
Swaps this network interface instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap73(value1 QNetworkAddressEntry_ITF, value2 QNetworkAddressEntry_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QNetworkAddressEntry_PTR() != nil {
		convArg0 = value1.QNetworkAddressEntry_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QNetworkAddressEntry_PTR() != nil {
		convArg1 = value2.QNetworkAddressEntry_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR20QNetworkAddressEntryS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:133
// index:75
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkProxyQuery &, QNetworkProxyQuery &)

/*
Swaps this network proxy instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap75(value1 QNetworkProxyQuery_ITF, value2 QNetworkProxyQuery_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QNetworkProxyQuery_PTR() != nil {
		convArg0 = value1.QNetworkProxyQuery_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QNetworkProxyQuery_PTR() != nil {
		convArg1 = value2.QNetworkProxyQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR18QNetworkProxyQueryS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:204
// index:76
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QSslConfiguration &, QSslConfiguration &)

/*
Swaps this SSL configuration instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap76(value1 QSslConfiguration_ITF, value2 QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QSslConfiguration_PTR() != nil {
		convArg0 = value1.QSslConfiguration_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QSslConfiguration_PTR() != nil {
		convArg1 = value2.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR17QSslConfigurationS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkinterface.h:176
// index:77
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkInterface &, QNetworkInterface &)

/*
Swaps this network interface instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap77(value1 QNetworkInterface_ITF, value2 QNetworkInterface_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QNetworkInterface_PTR() != nil {
		convArg0 = value1.QNetworkInterface_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QNetworkInterface_PTR() != nil {
		convArg1 = value2.QNetworkInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR17QNetworkInterfaceS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdatagram.h:115
// index:79
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkDatagram &, QNetworkDatagram &)

/*
Swaps this instance with other.
*/
func Swap79(value1 QNetworkDatagram_ITF, value2 QNetworkDatagram_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QNetworkDatagram_PTR() != nil {
		convArg0 = value1.QNetworkDatagram_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QNetworkDatagram_PTR() != nil {
		convArg1 = value2.QNetworkDatagram_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR16QNetworkDatagramS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcertificate.h:174
// index:80
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QSslCertificate &, QSslCertificate &)

/*
Swaps this certificate instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap80(value1 QSslCertificate_ITF, value2 QSslCertificate_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QSslCertificate_PTR() != nil {
		convArg0 = value1.QSslCertificate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QSslCertificate_PTR() != nil {
		convArg1 = value2.QSslCertificate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR15QSslCertificateS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkrequest.h:181
// index:81
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkRequest &, QNetworkRequest &)

/*
Swaps this network request with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap81(value1 QNetworkRequest_ITF, value2 QNetworkRequest_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QNetworkRequest_PTR() != nil {
		convArg0 = value1.QNetworkRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QNetworkRequest_PTR() != nil {
		convArg1 = value2.QNetworkRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR15QNetworkRequestS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkcookie.h:113
// index:82
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkCookie &, QNetworkCookie &)

/*
Swaps this cookie with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap82(value1 QNetworkCookie_ITF, value2 QNetworkCookie_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QNetworkCookie_PTR() != nil {
		convArg0 = value1.QNetworkCookie_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QNetworkCookie_PTR() != nil {
		convArg1 = value2.QNetworkCookie_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR14QNetworkCookieS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:213
// index:84
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkProxy &, QNetworkProxy &)

/*
Swaps this network proxy instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap84(value1 QNetworkProxy_ITF, value2 QNetworkProxy_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QNetworkProxy_PTR() != nil {
		convArg0 = value1.QNetworkProxy_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QNetworkProxy_PTR() != nil {
		convArg1 = value2.QNetworkProxy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR13QNetworkProxyS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostaddress.h:166
// index:85
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QHostAddress &, QHostAddress &)

/*
Swaps this host address with other. This operation is very fast and never fails.

This function was introduced in  Qt 5.6.
*/
func Swap85(value1 QHostAddress_ITF, value2 QHostAddress_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QHostAddress_PTR() != nil {
		convArg0 = value1.QHostAddress_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QHostAddress_PTR() != nil {
		convArg1 = value2.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR12QHostAddressS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:89
// index:86
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QHstsPolicy &, QHstsPolicy &)

/*
Swaps this policy with the other policy.
*/
func Swap86(value1 QHstsPolicy_ITF, value2 QHstsPolicy_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QHstsPolicy_PTR() != nil {
		convArg0 = value1.QHstsPolicy_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QHstsPolicy_PTR() != nil {
		convArg1 = value2.QHstsPolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR11QHstsPolicyS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcipher.h:90
// index:87
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QSslCipher &, QSslCipher &)

/*
Swaps this cipher instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap87(value1 QSslCipher_ITF, value2 QSslCipher_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QSslCipher_PTR() != nil {
		convArg0 = value1.QSslCipher_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QSslCipher_PTR() != nil {
		convArg1 = value2.QSslCipher_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR10QSslCipherS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end
