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
// extern C begin: 34
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
type QHostInfo struct {
	*qtrt.CObject
}
type QHostInfo_ITF interface {
	QHostInfo_PTR() *QHostInfo
}

func (ptr *QHostInfo) QHostInfo_PTR() *QHostInfo { return ptr }

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

/*
Constructs an empty host info object with lookup ID id.

See also lookupId().
*/
func NewQHostInfo(lookupId int) *QHostInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfoC2Ei", qtrt.FFI_TYPE_POINTER, lookupId)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHostInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHostInfo)
	return gothis
}

// /usr/include/qt/QtNetwork/qhostinfo.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHostInfo(int)

/*
Constructs an empty host info object with lookup ID id.

See also lookupId().
*/
func NewQHostInfo__() *QHostInfo {
	// arg: 0, int=Int, =Invalid,
	lookupId := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfoC2Ei", qtrt.FFI_TYPE_POINTER, lookupId)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHostInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHostInfo)
	return gothis
}

// /usr/include/qt/QtNetwork/qhostinfo.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QHostInfo & operator=(const QHostInfo &)

/*

 */
func (this *QHostInfo) Operator_equal(d QHostInfo_ITF) *QHostInfo {
	var convArg0 unsafe.Pointer
	if d != nil && d.QHostInfo_PTR() != nil {
		convArg0 = d.QHostInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfoaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostInfo)
	return rv2
}

// /usr/include/qt/QtNetwork/qhostinfo.h:66
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QHostInfo & operator=(QHostInfo &&)

/*

 */
func (this *QHostInfo) Operator_equal_1(other unsafe.Pointer /*333*/) *QHostInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfoaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQHostInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQHostInfo)
	return rv2
}

// /usr/include/qt/QtNetwork/qhostinfo.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QHostInfo()

/*

 */
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

/*
Swaps host-info other with this host-info. This operation is very fast and never fails.

This function was introduced in  Qt 5.10.
*/
func (this *QHostInfo) Swap(other QHostInfo_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QHostInfo_PTR() != nil {
		convArg0 = other.QHostInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QString hostName() const

/*
Returns the name of the host whose IP addresses were looked up.

See also setHostName() and localHostName().
*/
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

/*
Sets the host name of this QHostInfo to hostName.

See also hostName().
*/
func (this *QHostInfo) SetHostName(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo11setHostNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] QHostInfo::HostInfoError error() const

/*
Returns the type of error that occurred if the host name lookup failed; otherwise returns NoError.

See also setError() and errorString().
*/
func (this *QHostInfo) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QHostInfo5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setError(enum QHostInfo::HostInfoError)

/*
Sets the error type of this QHostInfo to error.

See also error() and errorString().
*/
func (this *QHostInfo) SetError(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo8setErrorENS_13HostInfoErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
If the lookup failed, this function returns a human readable description of the error; otherwise "Unknown error" is returned.

See also setErrorString() and error().
*/
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

/*
Sets the human readable description of the error that occurred to str if the lookup failed.

See also errorString() and setError().
*/
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

/*
Sets the ID of this lookup to id.

See also lookupId() and lookupHost().
*/
func (this *QHostInfo) SetLookupId(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo11setLookupIdEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qhostinfo.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int lookupId() const

/*
Returns the ID of this lookup.

See also setLookupId(), abortHostLookup(), and hostName().
*/
func (this *QHostInfo) LookupId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QHostInfo8lookupIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qhostinfo.h:86
// index:0
// Public static Visibility=Default Availability=Available
// [4] int lookupHost(const QString &, QObject *, const char *)

/*
Looks up the IP address(es) associated with host name name, and returns an ID for the lookup. When the result of the lookup is ready, the slot or signal member in receiver is called with a QHostInfo argument. The QHostInfo object can then be inspected to get the results of the lookup.

The lookup is performed by a single function call, for example:


  QHostInfo::lookupHost("www.kde.org",
                        this, SLOT(lookedUp(QHostInfo)));



The implementation of the slot prints basic information about the addresses returned by the lookup, or reports an error if it failed:


  void MyWidget::lookedUp(const QHostInfo &host)
  {
      if (host.error() != QHostInfo::NoError) {
          qDebug() << "Lookup failed:" << host.errorString();
          return;
      }

      const auto addresses = host.addresses();
      for (const QHostAddress &address : addresses)
          qDebug() << "Found address:" << address.toString();
  }



If you pass a literal IP address to name instead of a host name, QHostInfo will search for the domain name for the IP (i.e., QHostInfo will perform a reverse lookup). On success, the resulting QHostInfo will contain both the resolved domain name and IP addresses for the host name. Example:


  QHostInfo::lookupHost("4.2.2.1",
                        this, SLOT(lookedUp(QHostInfo)));



Note: There is no guarantee on the order the signals will be emitted if you start multiple requests with lookupHost().

See also abortHostLookup(), addresses(), error(), and fromName().
*/
func (this *QHostInfo) LookupHost(name string, receiver qtcore.QObject_ITF /*777 QObject **/, member string) int {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QHostInfo10lookupHostERK7QStringP7QObjectPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QHostInfo_LookupHost(name string, receiver qtcore.QObject_ITF /*777 QObject **/, member string) int {
	var nilthis *QHostInfo
	rv := nilthis.LookupHost(name, receiver, member)
	return rv
}

// /usr/include/qt/QtNetwork/qhostinfo.h:87
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void abortHostLookup(int)

/*
Aborts the host lookup with the ID id, as returned by lookupHost().

See also lookupHost() and lookupId().
*/
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

/*
Looks up the IP address(es) for the given host name. The function blocks during the lookup which means that execution of the program is suspended until the results of the lookup are ready. Returns the result of the lookup in a QHostInfo object.

If you pass a literal IP address to name instead of a host name, QHostInfo will search for the domain name for the IP (i.e., QHostInfo will perform a reverse lookup). On success, the returned QHostInfo will contain both the resolved domain name and IP addresses for the host name.

See also lookupHost().
*/
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

/*
Returns this machine's host name, if one is configured. Note that hostnames are not guaranteed to be globally unique, especially if they were configured automatically.

This function does not guarantee the returned host name is a Fully Qualified Domain Name (FQDN). For that, use fromName() to resolve the returned name to an FQDN.

This function returns the same as QSysInfo::machineHostName().

See also hostName() and localDomainName().
*/
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

/*
Returns the DNS domain of this machine.

Note: DNS domains are not related to domain names found in Windows networks.

See also hostName().
*/
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

/*
This enum describes the various errors that can occur when trying to resolve a host name.



See also error() and setError().

*/
type QHostInfo__HostInfoError = int

// The lookup was successful.
const QHostInfo__NoError QHostInfo__HostInfoError = 0

// No IP addresses were found for the host.
const QHostInfo__HostNotFound QHostInfo__HostInfoError = 1

// An unknown error occurred.
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
