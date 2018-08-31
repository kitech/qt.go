package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h
// #include <qnetworkdiskcache.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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

// long long expire()
func (this *QNetworkDiskCache) InheritExpire(f func() int64) {
	qtrt.SetAllInheritCallback(this, "expire", f)
}

/*

 */
type QNetworkDiskCache struct {
	*QAbstractNetworkCache
}
type QNetworkDiskCache_ITF interface {
	QAbstractNetworkCache_ITF
	QNetworkDiskCache_PTR() *QNetworkDiskCache
}

func (ptr *QNetworkDiskCache) QNetworkDiskCache_PTR() *QNetworkDiskCache { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QNetworkDiskCache) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkDiskCache10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkDiskCache(QObject *)

/*
Creates a new disk cache. The parent argument is passed to QAbstractNetworkCache's constructor.
*/
func (*QNetworkDiskCache) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkDiskCache {
	return NewQNetworkDiskCache(parent)
}
func NewQNetworkDiskCache(parent qtcore.QObject_ITF /*777 QObject **/) *QNetworkDiskCache {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCacheC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkDiskCacheFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkDiskCache")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNetworkDiskCache(QObject *)

/*
Creates a new disk cache. The parent argument is passed to QAbstractNetworkCache's constructor.
*/
func (*QNetworkDiskCache) NewForInherit__() *QNetworkDiskCache {
	return NewQNetworkDiskCache__()
}
func NewQNetworkDiskCache__() *QNetworkDiskCache {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCacheC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNetworkDiskCacheFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNetworkDiskCache")
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNetworkDiskCache()

/*

 */
func DeleteQNetworkDiskCache(this *QNetworkDiskCache) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCacheD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:59
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cacheDirectory() const

/*
Returns the location where cached files will be stored.

See also setCacheDirectory().
*/
func (this *QNetworkDiskCache) CacheDirectory() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkDiskCache14cacheDirectoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheDirectory(const QString &)

/*
Sets the directory where cached files will be stored to cacheDir

QNetworkDiskCache will create this directory if it does not exists.

Prepared cache items will be stored in the new cache directory when they are inserted.

See also cacheDirectory() and QDesktopServices::CacheLocation.
*/
func (this *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {
	var tmpArg0 = qtcore.NewQString_5(cacheDir)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCache17setCacheDirectoryERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 maximumCacheSize() const

/*
Returns the current maximum size for the disk cache.

See also setMaximumCacheSize().
*/
func (this *QNetworkDiskCache) MaximumCacheSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkDiskCache16maximumCacheSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumCacheSize(qint64)

/*
Sets the maximum size of the disk cache to be size.

If the new size is smaller then the current cache size then the cache will call expire().

See also maximumCacheSize().
*/
func (this *QNetworkDiskCache) SetMaximumCacheSize(size int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCache19setMaximumCacheSizeEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 cacheSize() const

/*
Reimplemented from QAbstractNetworkCache::cacheSize().
*/
func (this *QNetworkDiskCache) CacheSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkDiskCache9cacheSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QNetworkCacheMetaData metaData(const QUrl &)

/*
Reimplemented from QAbstractNetworkCache::metaData().
*/
func (this *QNetworkDiskCache) MetaData(url qtcore.QUrl_ITF) *QNetworkCacheMetaData /*123*/ {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCache8metaDataERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkCacheMetaDataFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkCacheMetaData)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateMetaData(const QNetworkCacheMetaData &)

/*
Reimplemented from QAbstractNetworkCache::updateMetaData().
*/
func (this *QNetworkDiskCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	var convArg0 unsafe.Pointer
	if metaData != nil && metaData.QNetworkCacheMetaData_PTR() != nil {
		convArg0 = metaData.QNetworkCacheMetaData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCache14updateMetaDataERK21QNetworkCacheMetaData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIODevice * data(const QUrl &)

/*
Reimplemented from QAbstractNetworkCache::data().
*/
func (this *QNetworkDiskCache) Data(url qtcore.QUrl_ITF) *qtcore.QIODevice /*777 QIODevice **/ {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCache4dataERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool remove(const QUrl &)

/*
Reimplemented from QAbstractNetworkCache::remove().
*/
func (this *QNetworkDiskCache) Remove(url qtcore.QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCache6removeERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIODevice * prepare(const QNetworkCacheMetaData &)

/*
Reimplemented from QAbstractNetworkCache::prepare().
*/
func (this *QNetworkDiskCache) Prepare(metaData QNetworkCacheMetaData_ITF) *qtcore.QIODevice /*777 QIODevice **/ {
	var convArg0 unsafe.Pointer
	if metaData != nil && metaData.QNetworkCacheMetaData_PTR() != nil {
		convArg0 = metaData.QNetworkCacheMetaData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCache7prepareERK21QNetworkCacheMetaData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkCacheMetaData fileMetaData(const QString &) const

/*
Returns the QNetworkCacheMetaData for the cache file fileName.

If fileName is not a cache file QNetworkCacheMetaData will be invalid.
*/
func (this *QNetworkDiskCache) FileMetaData(fileName string) *QNetworkCacheMetaData /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QNetworkDiskCache12fileMetaDataERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkCacheMetaDataFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkCacheMetaData)
	return rv2
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clear()

/*
Reimplemented from QAbstractNetworkCache::clear().
*/
func (this *QNetworkDiskCache) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCache5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkdiskcache.h:79
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 expire()

/*
Cleans the cache so that its size is under the maximum cache size. Returns the current size of the cache.

When the current size of the cache is greater than the maximumCacheSize() older cache files are removed until the total size is less then 90% of maximumCacheSize() starting with the oldest ones first using the file creation date to determine how old a cache file is.

Subclasses can reimplement this function to change the order that cache files are removed taking into account information in the application knows about that QNetworkDiskCache does not, for example the number of times a cache is accessed.

Note: cacheSize() calls expire if the current cache size is unknown.

See also maximumCacheSize() and fileMetaData().
*/
func (this *QNetworkDiskCache) Expire() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QNetworkDiskCache6expireEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
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
