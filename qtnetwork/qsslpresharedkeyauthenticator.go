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
// extern C begin: 35
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

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslPreSharedKeyAuthenticator()
func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticatorC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslPreSharedKeyAuthenticatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslPreSharedKeyAuthenticator)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslPreSharedKeyAuthenticator()
func DeleteQSslPreSharedKeyAuthenticator(this *QSslPreSharedKeyAuthenticator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:58
// index:0
// Public Visibility=Default Availability=Available
// [8] QSslPreSharedKeyAuthenticator & operator=(const QSslPreSharedKeyAuthenticator &)
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

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QSslPreSharedKeyAuthenticator & operator=(QSslPreSharedKeyAuthenticator &&)
func (this *QSslPreSharedKeyAuthenticator) Operator_equal_1(other unsafe.Pointer /*333*/) *QSslPreSharedKeyAuthenticator {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticatoraSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslPreSharedKeyAuthenticatorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslPreSharedKeyAuthenticator)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslPreSharedKeyAuthenticator &)
func (this *QSslPreSharedKeyAuthenticator) Swap(other QSslPreSharedKeyAuthenticator_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslPreSharedKeyAuthenticator_PTR() != nil {
		convArg0 = other.QSslPreSharedKeyAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray identityHint() const
func (this *QSslPreSharedKeyAuthenticator) IdentityHint() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator12identityHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIdentity(const QByteArray &)
func (this *QSslPreSharedKeyAuthenticator) SetIdentity(identity qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if identity != nil && identity.QByteArray_PTR() != nil {
		convArg0 = identity.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator11setIdentityERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray identity() const
func (this *QSslPreSharedKeyAuthenticator) Identity() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator8identityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumIdentityLength() const
func (this *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator21maximumIdentityLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreSharedKey(const QByteArray &)
func (this *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if preSharedKey != nil && preSharedKey.QByteArray_PTR() != nil {
		convArg0 = preSharedKey.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator15setPreSharedKeyERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray preSharedKey() const
func (this *QSslPreSharedKeyAuthenticator) PreSharedKey() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator12preSharedKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumPreSharedKeyLength() const
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
