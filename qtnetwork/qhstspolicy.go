package qtnetwork

// /usr/include/qt/QtNetwork/qhstspolicy.h
// #include <qhstspolicy.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QHstsPolicy struct {
	*qtrt.CObject
}

func (this *QHstsPolicy) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHstsPolicy) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHstsPolicyC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQHstsPolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHstsPolicy(const QDateTime &, QHstsPolicy::PolicyFlags, const QString &, QUrl::ParsingMode)
func NewQHstsPolicy_1(expiry *qtcore.QDateTime, flags int, host *qtcore.QString, mode int) *QHstsPolicy {
	var convArg0 = expiry.GetCthis()
	var convArg2 = host.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHstsPolicyC2ERK9QDateTime6QFlagsINS_10PolicyFlagEERK7QStringN4QUrl11ParsingModeE", ffiqt.FFI_TYPE_POINTER, convArg0, flags, convArg2, mode)
	gopp.ErrPrint(err, rv)
	gothis := NewQHstsPolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QHstsPolicy()
func DeleteQHstsPolicy(*QHstsPolicy) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHstsPolicyD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QHstsPolicy &)
func (this *QHstsPolicy) Swap(other *QHstsPolicy) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHstsPolicy4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHost(const QString &, QUrl::ParsingMode)
func (this *QHstsPolicy) SetHost(host *qtcore.QString, mode int) {
	var convArg0 = host.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHstsPolicy7setHostERK7QStringN4QUrl11ParsingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpiry(const QDateTime &)
func (this *QHstsPolicy) SetExpiry(expiry *qtcore.QDateTime) {
	var convArg0 = expiry.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHstsPolicy9setExpiryERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime expiry()
func (this *QHstsPolicy) Expiry() *qtcore.QDateTime /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHstsPolicy6expiryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIncludesSubDomains(_Bool)
func (this *QHstsPolicy) SetIncludesSubDomains(include bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHstsPolicy21setIncludesSubDomainsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), include)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool includesSubDomains()
func (this *QHstsPolicy) IncludesSubDomains() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHstsPolicy18includesSubDomainsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExpired()
func (this *QHstsPolicy) IsExpired() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHstsPolicy9isExpiredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QHstsPolicy__PolicyFlag = int

const QHstsPolicy__IncludeSubDomains QHstsPolicy__PolicyFlag = 1

//  body block end
