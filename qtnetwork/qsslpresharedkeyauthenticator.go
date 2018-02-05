package qtnetwork

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h
// #include <qsslpresharedkeyauthenticator.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QSslPreSharedKeyAuthenticator struct {
	*qtrt.CObject
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
	gopp.ErrPrint(err, rv)
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
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslPreSharedKeyAuthenticator &)
func (this *QSslPreSharedKeyAuthenticator) Swap(other *QSslPreSharedKeyAuthenticator) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray identityHint()
func (this *QSslPreSharedKeyAuthenticator) IdentityHint() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator12identityHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIdentity(const QByteArray &)
func (this *QSslPreSharedKeyAuthenticator) SetIdentity(identity *qtcore.QByteArray) {
	var convArg0 = identity.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator11setIdentityERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray identity()
func (this *QSslPreSharedKeyAuthenticator) Identity() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator8identityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumIdentityLength()
func (this *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator21maximumIdentityLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreSharedKey(const QByteArray &)
func (this *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey *qtcore.QByteArray) {
	var convArg0 = preSharedKey.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator15setPreSharedKeyERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray preSharedKey()
func (this *QSslPreSharedKeyAuthenticator) PreSharedKey() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator12preSharedKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumPreSharedKeyLength()
func (this *QSslPreSharedKeyAuthenticator) MaximumPreSharedKeyLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator25maximumPreSharedKeyLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
