//  header block begin

// +build !minimal

package qtnetwork

// /usr/include/qt/QtNetwork/qsslconfiguration.h
// #include <qsslconfiguration.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// /usr/include/qt/QtNetwork/qsslconfiguration.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dtlsCookieVerificationEnabled() const

/*
This function returns true if DTLS cookie verification was enabled on a server-side socket.

See also setDtlsCookieVerificationEnabled().
*/
func (this *QSslConfiguration) DtlsCookieVerificationEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSslConfiguration29dtlsCookieVerificationEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDtlsCookieVerificationEnabled(bool)

/*
This function enables DTLS cookie verification when enable is true.

See also dtlsCookieVerificationEnabled().
*/
func (this *QSslConfiguration) SetDtlsCookieVerificationEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration32setDtlsCookieVerificationEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:169
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSslConfiguration defaultDtlsConfiguration()

/*
Returns the default DTLS configuration to be used in new DTLS connections.

The default DTLS configuration consists of:


no local certificate and no private key
protocol DtlsV1_2OrLater
the system's default CA certificate list
the cipher list equal to the list of the SSL libraries' supported TLS 1.2 ciphers that use 128 or more secret bits for the cipher.


See also setDefaultDtlsConfiguration().
*/
func (this *QSslConfiguration) DefaultDtlsConfiguration() *QSslConfiguration /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration24defaultDtlsConfigurationEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslConfigurationFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslConfiguration)
	return rv2
}
func QSslConfiguration_DefaultDtlsConfiguration() *QSslConfiguration /*123*/ {
	var nilthis *QSslConfiguration
	rv := nilthis.DefaultDtlsConfiguration()
	return rv
}

// /usr/include/qt/QtNetwork/qsslconfiguration.h:170
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultDtlsConfiguration(const QSslConfiguration &)

/*
Sets the default DTLS configuration to be used in new DTLS connections to be configuration. Existing connections are not affected by this call.

See also defaultDtlsConfiguration().
*/
func (this *QSslConfiguration) SetDefaultDtlsConfiguration(configuration QSslConfiguration_ITF) {
	var convArg0 unsafe.Pointer
	if configuration != nil && configuration.QSslConfiguration_PTR() != nil {
		convArg0 = configuration.QSslConfiguration_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSslConfiguration27setDefaultDtlsConfigurationERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QSslConfiguration_SetDefaultDtlsConfiguration(configuration QSslConfiguration_ITF) {
	var nilthis *QSslConfiguration
	nilthis.SetDefaultDtlsConfiguration(configuration)
}

//  body block end

//  keep block begin

func init_unused_11430() {
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
		qtcore.KeepMe()
	}
}

//  keep block end
