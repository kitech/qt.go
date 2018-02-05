package qtnetwork

// /usr/include/qt/QtNetwork/qauthenticator.h
// #include <qauthenticator.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
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

type QAuthenticator struct {
	*qtrt.CObject
}

func (this *QAuthenticator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAuthenticator) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAuthenticatorFromPointer(cthis unsafe.Pointer) *QAuthenticator {
	return &QAuthenticator{&qtrt.CObject{cthis}}
}
func (*QAuthenticator) NewFromPointer(cthis unsafe.Pointer) *QAuthenticator {
	return NewQAuthenticatorFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qauthenticator.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAuthenticator()
func NewQAuthenticator() *QAuthenticator {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAuthenticatorC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQAuthenticatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAuthenticator)
	return gothis
}

// /usr/include/qt/QtNetwork/qauthenticator.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAuthenticator()
func DeleteQAuthenticator(this *QAuthenticator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAuthenticatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qauthenticator.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QString user()
func (this *QAuthenticator) User() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAuthenticator4userEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qauthenticator.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUser(const QString &)
func (this *QAuthenticator) SetUser(user *qtcore.QString) {
	var convArg0 = user.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAuthenticator7setUserERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qauthenticator.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString password()
func (this *QAuthenticator) Password() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAuthenticator8passwordEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qauthenticator.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPassword(const QString &)
func (this *QAuthenticator) SetPassword(password *qtcore.QString) {
	var convArg0 = password.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAuthenticator11setPasswordERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qauthenticator.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString realm()
func (this *QAuthenticator) Realm() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAuthenticator5realmEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qauthenticator.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRealm(const QString &)
func (this *QAuthenticator) SetRealm(realm *qtcore.QString) {
	var convArg0 = realm.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAuthenticator8setRealmERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qauthenticator.h:74
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant option(const QString &)
func (this *QAuthenticator) Option(opt *qtcore.QString) *qtcore.QVariant /*123*/ {
	var convArg0 = opt.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAuthenticator6optionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qauthenticator.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(const QString &, const QVariant &)
func (this *QAuthenticator) SetOption(opt *qtcore.QString, value *qtcore.QVariant) {
	var convArg0 = opt.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAuthenticator9setOptionERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qauthenticator.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QAuthenticator) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAuthenticator6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qauthenticator.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()
func (this *QAuthenticator) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAuthenticator6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
