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
// extern C begin: 18
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

/*
Constructs an invalid (expired) policy with empty host name and subdomains not included.
*/
func (*QHstsPolicy) NewForInherit() *QHstsPolicy {
	return NewQHstsPolicy()
}
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

/*
Constructs an invalid (expired) policy with empty host name and subdomains not included.
*/
func (*QHstsPolicy) NewForInherit1(expiry qtcore.QDateTime_ITF, flags int, host string, mode int) *QHstsPolicy {
	return NewQHstsPolicy1(expiry, flags, host, mode)
}
func NewQHstsPolicy1(expiry qtcore.QDateTime_ITF, flags int, host string, mode int) *QHstsPolicy {
	var convArg0 unsafe.Pointer
	if expiry != nil && expiry.QDateTime_PTR() != nil {
		convArg0 = expiry.QDateTime_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(host)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicyC2ERK9QDateTime6QFlagsINS_10PolicyFlagEERK7QStringN4QUrl11ParsingModeE", qtrt.FFI_TYPE_POINTER, convArg0, flags, convArg2, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHstsPolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHstsPolicy)
	return gothis
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHstsPolicy(const QDateTime &, QHstsPolicy::PolicyFlags, const QString &, QUrl::ParsingMode)

/*
Constructs an invalid (expired) policy with empty host name and subdomains not included.
*/
func (*QHstsPolicy) NewForInherit1p(expiry qtcore.QDateTime_ITF, flags int, host string) *QHstsPolicy {
	return NewQHstsPolicy1p(expiry, flags, host)
}
func NewQHstsPolicy1p(expiry qtcore.QDateTime_ITF, flags int, host string) *QHstsPolicy {
	var convArg0 unsafe.Pointer
	if expiry != nil && expiry.QDateTime_PTR() != nil {
		convArg0 = expiry.QDateTime_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(host)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QUrl::ParsingMode=Elaborated, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicyC2ERK9QDateTime6QFlagsINS_10PolicyFlagEERK7QStringN4QUrl11ParsingModeE", qtrt.FFI_TYPE_POINTER, convArg0, flags, convArg2, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHstsPolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHstsPolicy)
	return gothis
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QHstsPolicy & operator=(const QHstsPolicy &)

/*

 */
func (this *QHstsPolicy) Operator_equal(rhs QHstsPolicy_ITF) *QHstsPolicy {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QHstsPolicy_PTR() != nil {
		convArg0 = rhs.QHstsPolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicyaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHstsPolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHstsPolicy)
	return rv2
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:68
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QHstsPolicy & operator=(QHstsPolicy &&)

/*

 */
func (this *QHstsPolicy) Operator_equal1(other unsafe.Pointer /*333*/) *QHstsPolicy {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicyaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHstsPolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHstsPolicy)
	return rv2
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QHstsPolicy()

/*

 */
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

/*
Swaps this policy with the other policy.
*/
func (this *QHstsPolicy) Swap(other QHstsPolicy_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QHstsPolicy_PTR() != nil {
		convArg0 = other.QHstsPolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicy4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHost(const QString &, QUrl::ParsingMode)

/*
Sets a host, host data is interpreted according to mode parameter.

See also host(), QUrl::setHost(), and QUrl::ParsingMode.
*/
func (this *QHstsPolicy) SetHost(host string, mode int) {
	var tmpArg0 = qtcore.NewQString5(host)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicy7setHostERK7QStringN4QUrl11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHost(const QString &, QUrl::ParsingMode)

/*
Sets a host, host data is interpreted according to mode parameter.

See also host(), QUrl::setHost(), and QUrl::ParsingMode.
*/
func (this *QHstsPolicy) SetHostp(host string) {
	var tmpArg0 = qtcore.NewQString5(host)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Elaborated, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicy7setHostERK7QStringN4QUrl11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpiry(const QDateTime &)

/*
Sets the expiration date for the policy (in UTC) to expiry.

See also expiry().
*/
func (this *QHstsPolicy) SetExpiry(expiry qtcore.QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if expiry != nil && expiry.QDateTime_PTR() != nil {
		convArg0 = expiry.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicy9setExpiryERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime expiry() const

/*
Returns the expiration date for the policy (in UTC).

See also setExpiry().
*/
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
// [-2] void setIncludesSubDomains(bool)

/*
Sets whether subdomains are included for this policy to include.

See also includesSubDomains().
*/
func (this *QHstsPolicy) SetIncludesSubDomains(include_ bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHstsPolicy21setIncludesSubDomainsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), include_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool includesSubDomains() const

/*
Returns true if this policy also includes subdomains.

See also setIncludesSubDomains().
*/
func (this *QHstsPolicy) IncludesSubDomains() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHstsPolicy18includesSubDomainsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qhstspolicy.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExpired() const

/*
Return true if this policy has a valid expiration date and this date is greater than QDateTime::currentGetDateTimeUtc().

See also setExpiry() and expiry().
*/
func (this *QHstsPolicy) IsExpired() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHstsPolicy9isExpiredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QHstsPolicy__PolicyFlag = int

//
const QHstsPolicy__IncludeSubDomains QHstsPolicy__PolicyFlag = 1

func (this *QHstsPolicy) PolicyFlagItemName(val int) string {
	switch val {
	case QHstsPolicy__IncludeSubDomains: // 1
		return "IncludeSubDomains"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QHstsPolicy_PolicyFlagItemName(val int) string {
	var nilthis *QHstsPolicy
	return nilthis.PolicyFlagItemName(val)
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
