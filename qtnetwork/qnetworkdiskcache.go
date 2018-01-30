package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h
// #include <qnetworkdiskcache.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
type QNetworkDiskCache struct {
	*QAbstractNetworkCache
}

func (this *QNetworkDiskCache) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractNetworkCache.GetCthis()
	}
}
func (this *QNetworkDiskCache) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractNetworkCache = NewQAbstractNetworkCacheFromPointer(cthis)
}
func NewQNetworkDiskCacheFromPointer(cthis unsafe.Pointer) *QNetworkDiskCache {
	bcthis0 := NewQAbstractNetworkCacheFromPointer(cthis)
	return &QNetworkDiskCache{bcthis0}
}
func (*QNetworkDiskCache) NewFromPointer(cthis unsafe.Pointer) *QNetworkDiskCache {
	return NewQNetworkDiskCacheFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QNetworkDiskCache) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QNetworkDiskCache10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkDiskCache(QObject *)
func NewQNetworkDiskCache(parent *qtcore.QObject /*777 QObject **/) *QNetworkDiskCache {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCacheC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkDiskCacheFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNetworkDiskCache()
func DeleteQNetworkDiskCache(*QNetworkDiskCache) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCacheD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:59
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cacheDirectory()
func (this *QNetworkDiskCache) CacheDirectory() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QNetworkDiskCache14cacheDirectoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheDirectory(const QString &)
func (this *QNetworkDiskCache) SetCacheDirectory(cacheDir *qtcore.QString) {
	var convArg0 = cacheDir.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCache17setCacheDirectoryERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 maximumCacheSize()
func (this *QNetworkDiskCache) MaximumCacheSize() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QNetworkDiskCache16maximumCacheSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumCacheSize(qint64)
func (this *QNetworkDiskCache) SetMaximumCacheSize(size int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCache19setMaximumCacheSizeEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 cacheSize()
func (this *QNetworkDiskCache) CacheSize() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QNetworkDiskCache9cacheSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QNetworkCacheMetaData metaData(const QUrl &)
func (this *QNetworkDiskCache) MetaData(url *qtcore.QUrl) *QNetworkCacheMetaData /*123*/ {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCache8metaDataERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkCacheMetaDataFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateMetaData(const QNetworkCacheMetaData &)
func (this *QNetworkDiskCache) UpdateMetaData(metaData *QNetworkCacheMetaData) {
	var convArg0 = metaData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCache14updateMetaDataERK21QNetworkCacheMetaData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIODevice * data(const QUrl &)
func (this *QNetworkDiskCache) Data(url *qtcore.QUrl) *qtcore.QIODevice /*777 QIODevice **/ {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCache4dataERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool remove(const QUrl &)
func (this *QNetworkDiskCache) Remove(url *qtcore.QUrl) bool {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCache6removeERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIODevice * prepare(const QNetworkCacheMetaData &)
func (this *QNetworkDiskCache) Prepare(metaData *QNetworkCacheMetaData) *qtcore.QIODevice /*777 QIODevice **/ {
	var convArg0 = metaData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCache7prepareERK21QNetworkCacheMetaData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkCacheMetaData fileMetaData(const QString &)
func (this *QNetworkDiskCache) FileMetaData(fileName *qtcore.QString) *QNetworkCacheMetaData /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QNetworkDiskCache12fileMetaDataERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQNetworkCacheMetaDataFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clear()
func (this *QNetworkDiskCache) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCache5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:79
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 expire()
func (this *QNetworkDiskCache) Expire() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QNetworkDiskCache6expireEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

//  body block end
