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
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSslPreSharedKeyAuthenticatorFromPointer(cthis unsafe.Pointer) *QSslPreSharedKeyAuthenticator {
	return &QSslPreSharedKeyAuthenticator{&qtrt.CObject{cthis}}
}
func (*QSslPreSharedKeyAuthenticator) NewFromPointer(cthis unsafe.Pointer) *QSslPreSharedKeyAuthenticator {
	return NewQSslPreSharedKeyAuthenticatorFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:55
// index:0
// Public
// void QSslPreSharedKeyAuthenticator()
func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticatorC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslPreSharedKeyAuthenticatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:56
// index:0
// Public
// void ~QSslPreSharedKeyAuthenticator()
func DeleteQSslPreSharedKeyAuthenticator(*QSslPreSharedKeyAuthenticator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:64
// index:0
// Public inline
// void swap(QSslPreSharedKeyAuthenticator &)
func (this *QSslPreSharedKeyAuthenticator) Swap(other *QSslPreSharedKeyAuthenticator) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:66
// index:0
// Public
// QByteArray identityHint()
func (this *QSslPreSharedKeyAuthenticator) IdentityHint() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator12identityHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:68
// index:0
// Public
// void setIdentity(const QByteArray &)
func (this *QSslPreSharedKeyAuthenticator) SetIdentity(identity *qtcore.QByteArray) {
	var convArg0 = identity.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator11setIdentityERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:69
// index:0
// Public
// QByteArray identity()
func (this *QSslPreSharedKeyAuthenticator) Identity() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator8identityEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:70
// index:0
// Public
// int maximumIdentityLength()
func (this *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator21maximumIdentityLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:72
// index:0
// Public
// void setPreSharedKey(const QByteArray &)
func (this *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey *qtcore.QByteArray) {
	var convArg0 = preSharedKey.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QSslPreSharedKeyAuthenticator15setPreSharedKeyERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:73
// index:0
// Public
// QByteArray preSharedKey()
func (this *QSslPreSharedKeyAuthenticator) PreSharedKey() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator12preSharedKeyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qsslpresharedkeyauthenticator.h:74
// index:0
// Public
// int maximumPreSharedKeyLength()
func (this *QSslPreSharedKeyAuthenticator) MaximumPreSharedKeyLength() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QSslPreSharedKeyAuthenticator25maximumPreSharedKeyLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
