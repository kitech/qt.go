package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebenginecertificateerror.h
// #include <qwebenginecertificateerror.h>
// #include <QtWebEngineWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtprintsupport"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

//  ext block end

//  body block begin

/*

 */
type QWebEngineCertificateError struct {
	*qtrt.CObject
}
type QWebEngineCertificateError_ITF interface {
	QWebEngineCertificateError_PTR() *QWebEngineCertificateError
}

func (ptr *QWebEngineCertificateError) QWebEngineCertificateError_PTR() *QWebEngineCertificateError {
	return ptr
}

func (this *QWebEngineCertificateError) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineCertificateError) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineCertificateErrorFromPointer(cthis unsafe.Pointer) *QWebEngineCertificateError {
	return &QWebEngineCertificateError{&qtrt.CObject{cthis}}
}
func (*QWebEngineCertificateError) NewFromPointer(cthis unsafe.Pointer) *QWebEngineCertificateError {
	return NewQWebEngineCertificateErrorFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecertificateerror.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineCertificateError(int, QUrl, bool, QString)

/*

 */
func (*QWebEngineCertificateError) NewForInherit(error int, url qtcore.QUrl_ITF /*123*/, overridable bool, errorDescription string) *QWebEngineCertificateError {
	return NewQWebEngineCertificateError(error, url, overridable, errorDescription)
}
func NewQWebEngineCertificateError(error int, url qtcore.QUrl_ITF /*123*/, overridable bool, errorDescription string) *QWebEngineCertificateError {
	var convArg1 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg1 = url.QUrl_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString5(errorDescription)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWebEngineCertificateErrorC2Ei4QUrlb7QString", qtrt.FFI_TYPE_POINTER, error, convArg1, overridable, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineCertificateErrorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWebEngineCertificateError)
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecertificateerror.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWebEngineCertificateError()

/*

 */
func DeleteQWebEngineCertificateError(this *QWebEngineCertificateError) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWebEngineCertificateErrorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecertificateerror.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineCertificateError::Error error() const

/*

 */
func (this *QWebEngineCertificateError) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWebEngineCertificateError5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecertificateerror.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*

 */
func (this *QWebEngineCertificateError) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWebEngineCertificateError3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecertificateerror.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isOverridable() const

/*

 */
func (this *QWebEngineCertificateError) IsOverridable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWebEngineCertificateError13isOverridableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecertificateerror.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorDescription() const

/*

 */
func (this *QWebEngineCertificateError) ErrorDescription() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWebEngineCertificateError16errorDescriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

/*


 */
type QWebEngineCertificateError__Error = int

//
const QWebEngineCertificateError__SslPinnedKeyNotInCertificateChain QWebEngineCertificateError__Error = -150

//
const QWebEngineCertificateError__CertificateCommonNameInvalid QWebEngineCertificateError__Error = -200

//
const QWebEngineCertificateError__CertificateDateInvalid QWebEngineCertificateError__Error = -201

//
const QWebEngineCertificateError__CertificateAuthorityInvalid QWebEngineCertificateError__Error = -202

//
const QWebEngineCertificateError__CertificateContainsErrors QWebEngineCertificateError__Error = -203

//
const QWebEngineCertificateError__CertificateNoRevocationMechanism QWebEngineCertificateError__Error = -204

//
const QWebEngineCertificateError__CertificateUnableToCheckRevocation QWebEngineCertificateError__Error = -205

//
const QWebEngineCertificateError__CertificateRevoked QWebEngineCertificateError__Error = -206

//
const QWebEngineCertificateError__CertificateInvalid QWebEngineCertificateError__Error = -207

//
const QWebEngineCertificateError__CertificateWeakSignatureAlgorithm QWebEngineCertificateError__Error = -208

//
const QWebEngineCertificateError__CertificateNonUniqueName QWebEngineCertificateError__Error = -210

//
const QWebEngineCertificateError__CertificateWeakKey QWebEngineCertificateError__Error = -211

//
const QWebEngineCertificateError__CertificateNameConstraintViolation QWebEngineCertificateError__Error = -212

//
const QWebEngineCertificateError__CertificateValidityTooLong QWebEngineCertificateError__Error = -213

//
const QWebEngineCertificateError__CertificateTransparencyRequired QWebEngineCertificateError__Error = -214

func (this *QWebEngineCertificateError) ErrorItemName(val int) string {
	switch val {
	case QWebEngineCertificateError__SslPinnedKeyNotInCertificateChain: // -150
		return "SslPinnedKeyNotInCertificateChain"
	case QWebEngineCertificateError__CertificateCommonNameInvalid: // -200
		return "CertificateCommonNameInvalid"
	case QWebEngineCertificateError__CertificateDateInvalid: // -201
		return "CertificateDateInvalid"
	case QWebEngineCertificateError__CertificateAuthorityInvalid: // -202
		return "CertificateAuthorityInvalid"
	case QWebEngineCertificateError__CertificateContainsErrors: // -203
		return "CertificateContainsErrors"
	case QWebEngineCertificateError__CertificateNoRevocationMechanism: // -204
		return "CertificateNoRevocationMechanism"
	case QWebEngineCertificateError__CertificateUnableToCheckRevocation: // -205
		return "CertificateUnableToCheckRevocation"
	case QWebEngineCertificateError__CertificateRevoked: // -206
		return "CertificateRevoked"
	case QWebEngineCertificateError__CertificateInvalid: // -207
		return "CertificateInvalid"
	case QWebEngineCertificateError__CertificateWeakSignatureAlgorithm: // -208
		return "CertificateWeakSignatureAlgorithm"
	case QWebEngineCertificateError__CertificateNonUniqueName: // -210
		return "CertificateNonUniqueName"
	case QWebEngineCertificateError__CertificateWeakKey: // -211
		return "CertificateWeakKey"
	case QWebEngineCertificateError__CertificateNameConstraintViolation: // -212
		return "CertificateNameConstraintViolation"
	case QWebEngineCertificateError__CertificateValidityTooLong: // -213
		return "CertificateValidityTooLong"
	case QWebEngineCertificateError__CertificateTransparencyRequired: // -214
		return "CertificateTransparencyRequired"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineCertificateError_ErrorItemName(val int) string {
	var nilthis *QWebEngineCertificateError
	return nilthis.ErrorItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11707() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtquickwidgets.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtwidgets.KeepMe()
	}
	if false {
		qtprintsupport.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
	if false {
		qtwebenginecore.KeepMe()
	}
}

//  keep block end
