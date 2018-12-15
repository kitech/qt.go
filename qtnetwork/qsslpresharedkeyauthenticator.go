// +build !minimal

package qtnetwork

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h
// #include <qsslpresharedkeyauthenticator.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
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

/*

 */
type QSslPreSharedKeyAuthenticator struct {
	*qtrt.CObject
}
type QSslPreSharedKeyAuthenticator_ITF interface {
	QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator
}

func (ptr *QSslPreSharedKeyAuthenticator) QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator {
	return ptr
}

func (this *QSslPreSharedKeyAuthenticator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslPreSharedKeyAuthenticator) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslPreSharedKeyAuthenticatorFromPointer(cthis unsafe.Pointer) *QSslPreSharedKeyAuthenticator {
	return &QSslPreSharedKeyAuthenticator{&qtrt.CObject{cthis}}
}
func (*QSslPreSharedKeyAuthenticator) NewFromPointer(cthis unsafe.Pointer) *QSslPreSharedKeyAuthenticator {
	return NewQSslPreSharedKeyAuthenticatorFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslPreSharedKeyAuthenticator()

/*
Constructs a default QSslPreSharedKeyAuthenticator object.

The identity hint, the identity and the key will be initialized to empty byte arrays; the maximum length for both the identity and the key will be initialized to 0.
*/
func (*QSslPreSharedKeyAuthenticator) NewForInherit() *QSslPreSharedKeyAuthenticator {
	return NewQSslPreSharedKeyAuthenticator()
}
func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticatorC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslPreSharedKeyAuthenticatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslPreSharedKeyAuthenticator)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslPreSharedKeyAuthenticator()

/*

 */
func DeleteQSslPreSharedKeyAuthenticator(this *QSslPreSharedKeyAuthenticator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslPreSharedKeyAuthenticator & operator=(const QSslPreSharedKeyAuthenticator &)

/*

 */
func (this *QSslPreSharedKeyAuthenticator) Operator_equal(authenticator QSslPreSharedKeyAuthenticator_ITF) *QSslPreSharedKeyAuthenticator {
	var convArg0 unsafe.Pointer
	if authenticator != nil && authenticator.QSslPreSharedKeyAuthenticator_PTR() != nil {
		convArg0 = authenticator.QSslPreSharedKeyAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticatoraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslPreSharedKeyAuthenticatorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslPreSharedKeyAuthenticator)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:63
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QSslPreSharedKeyAuthenticator & operator=(QSslPreSharedKeyAuthenticator &&)

/*

 */
func (this *QSslPreSharedKeyAuthenticator) Operator_equal1(other unsafe.Pointer /*333*/) *QSslPreSharedKeyAuthenticator {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticatoraSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslPreSharedKeyAuthenticatorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslPreSharedKeyAuthenticator)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslPreSharedKeyAuthenticator &)

/*
Swaps the QSslPreSharedKeyAuthenticator object authenticator with this object. This operation is very fast and never fails.
*/
func (this *QSslPreSharedKeyAuthenticator) Swap(other QSslPreSharedKeyAuthenticator_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslPreSharedKeyAuthenticator_PTR() != nil {
		convArg0 = other.QSslPreSharedKeyAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray identityHint() const

/*
Returns the PSK identity hint as provided by the server. The interpretation of this hint is left to the application.
*/
func (this *QSslPreSharedKeyAuthenticator) IdentityHint() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator12identityHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIdentity(const QByteArray &)

/*
Sets the PSK client identity (to be advised to the server) to identity.

Note: it is possible to set an identity whose length is greater than maximumIdentityLength(); in this case, only the first maximumIdentityLength() bytes will be actually sent to the server.

See also identity() and maximumIdentityLength().
*/
func (this *QSslPreSharedKeyAuthenticator) SetIdentity(identity qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if identity != nil && identity.QByteArray_PTR() != nil {
		convArg0 = identity.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator11setIdentityERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray identity() const

/*
Returns the PSK client identity.

See also setIdentity().
*/
func (this *QSslPreSharedKeyAuthenticator) Identity() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator8identityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:72
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumIdentityLength() const

/*
Returns the maximum length, in bytes, of the PSK client identity.

Note: it is possible to set an identity whose length is greater than maximumIdentityLength(); in this case, only the first maximumIdentityLength() bytes will be actually sent to the server.

See also setIdentity().
*/
func (this *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator21maximumIdentityLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreSharedKey(const QByteArray &)

/*
Sets the pre shared key to preSharedKey.

Note: it is possible to set a key whose length is greater than the maximumPreSharedKeyLength(); in this case, only the first maximumPreSharedKeyLength() bytes will be actually sent to the server.

See also preSharedKey(), maximumPreSharedKeyLength(), and QByteArray::fromHex().
*/
func (this *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if preSharedKey != nil && preSharedKey.QByteArray_PTR() != nil {
		convArg0 = preSharedKey.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator15setPreSharedKeyERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray preSharedKey() const

/*
Returns the pre shared key.

See also setPreSharedKey().
*/
func (this *QSslPreSharedKeyAuthenticator) PreSharedKey() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator12preSharedKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumPreSharedKeyLength() const

/*
Returns the maximum length, in bytes, of the pre shared key.

Note: it is possible to set a key whose length is greater than the maximumPreSharedKeyLength(); in this case, only the first maximumPreSharedKeyLength() bytes will be actually sent to the server.

See also setPreSharedKey().
*/
func (this *QSslPreSharedKeyAuthenticator) MaximumPreSharedKeyLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator25maximumPreSharedKeyLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end

//  keep block begin

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
