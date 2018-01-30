package qtnetwork

// /usr/include/qt/QtNetwork/qhostinfo.h
// #include <qhostinfo.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QHostInfo struct {
	*qtrt.CObject
}

func (this *QHostInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHostInfo) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQHostInfoFromPointer(cthis unsafe.Pointer) *QHostInfo {
	return &QHostInfo{&qtrt.CObject{cthis}}
}
func (*QHostInfo) NewFromPointer(cthis unsafe.Pointer) *QHostInfo {
	return NewQHostInfoFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHostInfo(int)
func NewQHostInfo(lookupId int) *QHostInfo {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfoC2Ei", ffiqt.FFI_TYPE_POINTER, lookupId)
	gopp.ErrPrint(err, rv)
	gothis := NewQHostInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qhostinfo.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QHostInfo()
func DeleteQHostInfo(*QHostInfo) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfoD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QHostInfo &)
func (this *QHostInfo) Swap(other *QHostInfo) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfo4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString hostName()
func (this *QHostInfo) HostName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QHostInfo8hostNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qhostinfo.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHostName(const QString &)
func (this *QHostInfo) SetHostName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfo11setHostNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] QHostInfo::HostInfoError error()
func (this *QHostInfo) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QHostInfo5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setError(enum QHostInfo::HostInfoError)
func (this *QHostInfo) SetError(error int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfo8setErrorENS_13HostInfoErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), error)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QHostInfo) ErrorString() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QHostInfo11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qhostinfo.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setErrorString(const QString &)
func (this *QHostInfo) SetErrorString(errorString *qtcore.QString) {
	var convArg0 = errorString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfo14setErrorStringERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLookupId(int)
func (this *QHostInfo) SetLookupId(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfo11setLookupIdEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int lookupId()
func (this *QHostInfo) LookupId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QHostInfo8lookupIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qhostinfo.h:86
// index:0
// Public static Visibility=Default Availability=Available
// [4] int lookupHost(const QString &, QObject *, const char *)
func (this *QHostInfo) LookupHost(name *qtcore.QString, receiver *qtcore.QObject /*777 QObject **/, member string) int {
	var convArg0 = name.GetCthis()
	var convArg1 = receiver.GetCthis()
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfo10lookupHostERK7QStringP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QHostInfo_LookupHost(name *qtcore.QString, receiver *qtcore.QObject /*777 QObject **/, member string) int {
	var nilthis *QHostInfo
	rv := nilthis.LookupHost(name, receiver, member)
	return rv
}

// /usr/include/qt/QtNetwork/qhostinfo.h:87
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void abortHostLookup(int)
func (this *QHostInfo) AbortHostLookup(lookupId int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfo15abortHostLookupEi", ffiqt.FFI_TYPE_POINTER, lookupId)
	gopp.ErrPrint(err, rv)
}
func QHostInfo_AbortHostLookup(lookupId int) {
	var nilthis *QHostInfo
	nilthis.AbortHostLookup(lookupId)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:89
// index:0
// Public static Visibility=Default Availability=Available
// [8] QHostInfo fromName(const QString &)
func (this *QHostInfo) FromName(name *qtcore.QString) *QHostInfo /*123*/ {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfo8fromNameERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQHostInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QHostInfo_FromName(name *qtcore.QString) *QHostInfo /*123*/ {
	var nilthis *QHostInfo
	rv := nilthis.FromName(name)
	return rv
}

// /usr/include/qt/QtNetwork/qhostinfo.h:90
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString localHostName()
func (this *QHostInfo) LocalHostName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfo13localHostNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QHostInfo_LocalHostName() *qtcore.QString /*123*/ {
	var nilthis *QHostInfo
	rv := nilthis.LocalHostName()
	return rv
}

// /usr/include/qt/QtNetwork/qhostinfo.h:91
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString localDomainName()
func (this *QHostInfo) LocalDomainName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QHostInfo15localDomainNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QHostInfo_LocalDomainName() *qtcore.QString /*123*/ {
	var nilthis *QHostInfo
	rv := nilthis.LocalDomainName()
	return rv
}

type QHostInfo__HostInfoError = int

const QHostInfo__NoError QHostInfo__HostInfoError = 0
const QHostInfo__HostNotFound QHostInfo__HostInfoError = 1
const QHostInfo__UnknownError QHostInfo__HostInfoError = 2

//  body block end
