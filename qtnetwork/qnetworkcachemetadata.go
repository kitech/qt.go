package qtnetwork

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h
// #include <qabstractnetworkcache.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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

type QNetworkCacheMetaData struct {
	*qtrt.CObject
}

func (this *QNetworkCacheMetaData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkCacheMetaData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQNetworkCacheMetaDataFromPointer(cthis unsafe.Pointer) *QNetworkCacheMetaData {
	return &QNetworkCacheMetaData{&qtrt.CObject{cthis}}
}
func (*QNetworkCacheMetaData) NewFromPointer(cthis unsafe.Pointer) *QNetworkCacheMetaData {
	return NewQNetworkCacheMetaDataFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkCacheMetaData()
func NewQNetworkCacheMetaData() *QNetworkCacheMetaData {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkCacheMetaDataC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkCacheMetaDataFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQNetworkCacheMetaData)
	return gothis
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QNetworkCacheMetaData()
func DeleteQNetworkCacheMetaData(this *QNetworkCacheMetaData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkCacheMetaDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QNetworkCacheMetaData &)
func (this *QNetworkCacheMetaData) Swap(other *QNetworkCacheMetaData) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkCacheMetaData4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QNetworkCacheMetaData) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkCacheMetaData7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url()
func (this *QNetworkCacheMetaData) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkCacheMetaData3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QUrl &)
func (this *QNetworkCacheMetaData) SetUrl(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkCacheMetaData6setUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime lastModified()
func (this *QNetworkCacheMetaData) LastModified() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkCacheMetaData12lastModifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastModified(const QDateTime &)
func (this *QNetworkCacheMetaData) SetLastModified(dateTime *qtcore.QDateTime) {
	var convArg0 = dateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkCacheMetaData15setLastModifiedERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime expirationDate()
func (this *QNetworkCacheMetaData) ExpirationDate() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkCacheMetaData14expirationDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpirationDate(const QDateTime &)
func (this *QNetworkCacheMetaData) SetExpirationDate(dateTime *qtcore.QDateTime) {
	var convArg0 = dateTime.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkCacheMetaData17setExpirationDateERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool saveToDisk()
func (this *QNetworkCacheMetaData) SaveToDisk() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QNetworkCacheMetaData10saveToDiskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSaveToDisk(_Bool)
func (this *QNetworkCacheMetaData) SetSaveToDisk(allow bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QNetworkCacheMetaData13setSaveToDiskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), allow)
	gopp.ErrPrint(err, rv)
}

//  body block end
