package qtnetwork

// /usr/include/qt/QtNetwork/qhstspolicy.h
// #include <qhstspolicy.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QHstsPolicy struct {
	*qtrt.CObject
}
type QHstsPolicy_ITF interface {
	QHstsPolicy_PTR() *QHstsPolicy
}

func (ptr *QHstsPolicy) QHstsPolicy_PTR() *QHstsPolicy { return ptr }

func (this *QHstsPolicy) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHstsPolicy) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQHstsPolicyFromPointer(cthis unsafe.Pointer) *QHstsPolicy {
	return &QHstsPolicy{&qtrt.CObject{cthis}}
}
func (*QHstsPolicy) NewFromPointer(cthis unsafe.Pointer) *QHstsPolicy {
	return NewQHstsPolicyFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHstsPolicy()
func NewQHstsPolicy() *QHstsPolicy {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicyC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHstsPolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHstsPolicy)
	return gothis
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHstsPolicy(const QDateTime &, QHstsPolicy::PolicyFlags, const QString &, QUrl::ParsingMode)
func NewQHstsPolicy_1(expiry *qtcore.QDateTime, flags int, host string, mode int) *QHstsPolicy {
	var convArg0 = expiry.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(host)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicyC2ERK9QDateTime6QFlagsINS_10PolicyFlagEERK7QStringN4QUrl11ParsingModeE", qtrt.FFI_TYPE_POINTER, convArg0, flags, convArg2, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHstsPolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHstsPolicy)
	return gothis
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QHstsPolicy()
func DeleteQHstsPolicy(this *QHstsPolicy) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QHstsPolicy &)
func (this *QHstsPolicy) Swap(other *QHstsPolicy) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicy4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHost(const QString &, QUrl::ParsingMode)
func (this *QHstsPolicy) SetHost(host string, mode int) {
	var tmpArg0 = qtcore.NewQString_5(host)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicy7setHostERK7QStringN4QUrl11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpiry(const QDateTime &)
func (this *QHstsPolicy) SetExpiry(expiry *qtcore.QDateTime) {
	var convArg0 = expiry.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicy9setExpiryERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime expiry()
func (this *QHstsPolicy) Expiry() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHstsPolicy6expiryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIncludesSubDomains(_Bool)
func (this *QHstsPolicy) SetIncludesSubDomains(include bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicy21setIncludesSubDomainsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), include)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool includesSubDomains()
func (this *QHstsPolicy) IncludesSubDomains() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHstsPolicy18includesSubDomainsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExpired()
func (this *QHstsPolicy) IsExpired() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHstsPolicy9isExpiredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QHstsPolicy__PolicyFlag = int

const QHstsPolicy__IncludeSubDomains QHstsPolicy__PolicyFlag = 1

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
