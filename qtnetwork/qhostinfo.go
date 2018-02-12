package qtnetwork

// /usr/include/qt/QtNetwork/qhostinfo.h
// #include <qhostinfo.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfoC2Ei", qtrt.FFI_TYPE_POINTER, lookupId)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHostInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHostInfo)
	return gothis
}

// /usr/include/qt/QtNetwork/qhostinfo.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QHostInfo()
func DeleteQHostInfo(this *QHostInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QHostInfo &)
func (this *QHostInfo) Swap(other *QHostInfo) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString hostName()
func (this *QHostInfo) HostName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QHostInfo8hostNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qhostinfo.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHostName(const QString &)
func (this *QHostInfo) SetHostName(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo11setHostNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] QHostInfo::HostInfoError error()
func (this *QHostInfo) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QHostInfo5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setError(enum QHostInfo::HostInfoError)
func (this *QHostInfo) SetError(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo8setErrorENS_13HostInfoErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QHostInfo) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QHostInfo11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qhostinfo.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setErrorString(const QString &)
func (this *QHostInfo) SetErrorString(errorString string) {
	var tmpArg0 = qtcore.NewQString_5(errorString)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo14setErrorStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLookupId(int)
func (this *QHostInfo) SetLookupId(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo11setLookupIdEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int lookupId()
func (this *QHostInfo) LookupId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QHostInfo8lookupIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qhostinfo.h:86
// index:0
// Public static Visibility=Default Availability=Available
// [4] int lookupHost(const QString &, QObject *, const char *)
func (this *QHostInfo) LookupHost(name string, receiver *qtcore.QObject /*777 QObject **/, member string) int {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = receiver.GetCthis()
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo10lookupHostERK7QStringP7QObjectPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QHostInfo_LookupHost(name string, receiver *qtcore.QObject /*777 QObject **/, member string) int {
	var nilthis *QHostInfo
	rv := nilthis.LookupHost(name, receiver, member)
	return rv
}

// /usr/include/qt/QtNetwork/qhostinfo.h:87
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void abortHostLookup(int)
func (this *QHostInfo) AbortHostLookup(lookupId int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo15abortHostLookupEi", qtrt.FFI_TYPE_POINTER, lookupId)
	qtrt.ErrPrint(err, rv)
}
func QHostInfo_AbortHostLookup(lookupId int) {
	var nilthis *QHostInfo
	nilthis.AbortHostLookup(lookupId)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:89
// index:0
// Public static Visibility=Default Availability=Available
// [8] QHostInfo fromName(const QString &)
func (this *QHostInfo) FromName(name string) *QHostInfo /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo8fromNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostInfo)
	return rv2
}
func QHostInfo_FromName(name string) *QHostInfo /*123*/ {
	var nilthis *QHostInfo
	rv := nilthis.FromName(name)
	return rv
}

// /usr/include/qt/QtNetwork/qhostinfo.h:90
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString localHostName()
func (this *QHostInfo) LocalHostName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo13localHostNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QHostInfo_LocalHostName() string {
	var nilthis *QHostInfo
	rv := nilthis.LocalHostName()
	return rv
}

// /usr/include/qt/QtNetwork/qhostinfo.h:91
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString localDomainName()
func (this *QHostInfo) LocalDomainName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo15localDomainNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QHostInfo_LocalDomainName() string {
	var nilthis *QHostInfo
	rv := nilthis.LocalDomainName()
	return rv
}

type QHostInfo__HostInfoError = int

const QHostInfo__NoError QHostInfo__HostInfoError = 0
const QHostInfo__HostNotFound QHostInfo__HostInfoError = 1
const QHostInfo__UnknownError QHostInfo__HostInfoError = 2

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
