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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQNetworkCacheMetaDataFromPointer(cthis unsafe.Pointer) *QNetworkCacheMetaData {
	return &QNetworkCacheMetaData{&qtrt.CObject{cthis}}
}
func (*QNetworkCacheMetaData) NewFromPointer(cthis unsafe.Pointer) *QNetworkCacheMetaData {
	return NewQNetworkCacheMetaDataFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:66
// index:0
// Public
// void QNetworkCacheMetaData()
func NewQNetworkCacheMetaData() *QNetworkCacheMetaData {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkCacheMetaDataC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkCacheMetaDataFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:68
// index:0
// Public
// void ~QNetworkCacheMetaData()
func DeleteQNetworkCacheMetaData(*QNetworkCacheMetaData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkCacheMetaDataD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:75
// index:0
// Public inline
// void swap(QNetworkCacheMetaData &)
func (this *QNetworkCacheMetaData) Swap(other *QNetworkCacheMetaData) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkCacheMetaData4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:82
// index:0
// Public
// bool isValid()
func (this *QNetworkCacheMetaData) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkCacheMetaData7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:84
// index:0
// Public
// QUrl url()
func (this *QNetworkCacheMetaData) Url() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkCacheMetaData3urlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:85
// index:0
// Public
// void setUrl(const QUrl &)
func (this *QNetworkCacheMetaData) SetUrl(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkCacheMetaData6setUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:90
// index:0
// Public
// QDateTime lastModified()
func (this *QNetworkCacheMetaData) LastModified() *qtcore.QDateTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkCacheMetaData12lastModifiedEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:91
// index:0
// Public
// void setLastModified(const QDateTime &)
func (this *QNetworkCacheMetaData) SetLastModified(dateTime *qtcore.QDateTime) {
	var convArg0 = dateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkCacheMetaData15setLastModifiedERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:93
// index:0
// Public
// QDateTime expirationDate()
func (this *QNetworkCacheMetaData) ExpirationDate() *qtcore.QDateTime /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkCacheMetaData14expirationDateEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:94
// index:0
// Public
// void setExpirationDate(const QDateTime &)
func (this *QNetworkCacheMetaData) SetExpirationDate(dateTime *qtcore.QDateTime) {
	var convArg0 = dateTime.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkCacheMetaData17setExpirationDateERK9QDateTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:96
// index:0
// Public
// bool saveToDisk()
func (this *QNetworkCacheMetaData) SaveToDisk() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QNetworkCacheMetaData10saveToDiskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:97
// index:0
// Public
// void setSaveToDisk(bool)
func (this *QNetworkCacheMetaData) SetSaveToDisk(allow bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QNetworkCacheMetaData13setSaveToDiskEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), allow)
	gopp.ErrPrint(err, rv)
}

//  body block end
