package qtnetwork

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h
// #include <qabstractnetworkcache.h>
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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QAbstractNetworkCache struct {
	*qtcore.QObject
}
type QAbstractNetworkCache_ITF interface {
	qtcore.QObject_ITF
	QAbstractNetworkCache_PTR() *QAbstractNetworkCache
}

func (ptr *QAbstractNetworkCache) QAbstractNetworkCache_PTR() *QAbstractNetworkCache { return ptr }

func (this *QAbstractNetworkCache) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractNetworkCache) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAbstractNetworkCacheFromPointer(cthis unsafe.Pointer) *QAbstractNetworkCache {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractNetworkCache{bcthis0}
}
func (*QAbstractNetworkCache) NewFromPointer(cthis unsafe.Pointer) *QAbstractNetworkCache {
	return NewQAbstractNetworkCacheFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:116
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractNetworkCache) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractNetworkCache10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractNetworkCache()

/*

 */
func DeleteQAbstractNetworkCache(this *QAbstractNetworkCache) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractNetworkCacheD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:121
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QNetworkCacheMetaData metaData(const QUrl &)

/*
Returns the meta data for the url url.

If the url is valid and the cache contains the data for url, a valid QNetworkCacheMetaData is returned.

In the base class this is a pure virtual function.

See also updateMetaData() and data().
*/
func (this *QAbstractNetworkCache) MetaData(url qtcore.QUrl_ITF) *QNetworkCacheMetaData /*123*/ {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractNetworkCache8metaDataERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkCacheMetaDataFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkCacheMetaData)
	return rv2
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:122
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void updateMetaData(const QNetworkCacheMetaData &)

/*
Updates the cache meta date for the metaData's url to metaData

If the cache does not contains a cache item for the url then no action is taken.

In the base class this is a pure virtual function.

See also metaData() and prepare().
*/
func (this *QAbstractNetworkCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	var convArg0 unsafe.Pointer
	if metaData != nil && metaData.QNetworkCacheMetaData_PTR() != nil {
		convArg0 = metaData.QNetworkCacheMetaData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractNetworkCache14updateMetaDataERK21QNetworkCacheMetaData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:123
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIODevice * data(const QUrl &)

/*
Returns the data associated with url.

It is up to the application that requests the data to delete the QIODevice when done with it.

If there is no cache for url, the url is invalid, or if there is an internal cache error 0 is returned.

In the base class this is a pure virtual function.

See also metaData() and prepare().
*/
func (this *QAbstractNetworkCache) Data(url qtcore.QUrl_ITF) *qtcore.QIODevice /*777 QIODevice **/ {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractNetworkCache4dataERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:124
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool remove(const QUrl &)

/*
Removes the cache entry for url, returning true if success otherwise false.

In the base class this is a pure virtual function.

See also clear() and prepare().
*/
func (this *QAbstractNetworkCache) Remove(url qtcore.QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractNetworkCache6removeERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:125
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] qint64 cacheSize() const

/*
Returns the current size taken up by the cache. Depending upon the cache implementation this might be disk or memory size.

In the base class this is a pure virtual function.

See also clear().
*/
func (this *QAbstractNetworkCache) CacheSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractNetworkCache9cacheSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:127
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIODevice * prepare(const QNetworkCacheMetaData &)

/*
Returns the device that should be populated with the data for the cache item metaData. When all of the data has been written insert() should be called. If metaData is invalid or the url in the metadata is invalid 0 is returned.

The cache owns the device and will take care of deleting it when it is inserted or removed.

To cancel a prepared inserted call remove() on the metadata's url.

In the base class this is a pure virtual function.

See also remove(), updateMetaData(), and insert().
*/
func (this *QAbstractNetworkCache) Prepare(metaData QNetworkCacheMetaData_ITF) *qtcore.QIODevice /*777 QIODevice **/ {
	var convArg0 unsafe.Pointer
	if metaData != nil && metaData.QNetworkCacheMetaData_PTR() != nil {
		convArg0 = metaData.QNetworkCacheMetaData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractNetworkCache7prepareERK21QNetworkCacheMetaData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:131
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all items from the cache. Unless there was failures clearing the cache cacheSize() should return 0 after a call to clear.

In the base class this is a pure virtual function.

See also cacheSize() and remove().
*/
func (this *QAbstractNetworkCache) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractNetworkCache5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:134
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAbstractNetworkCache(QObject *)

/*
Constructs an abstract network cache with the given parent.
*/
func (*QAbstractNetworkCache) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAbstractNetworkCache {
	return NewQAbstractNetworkCache(parent)
}
func NewQAbstractNetworkCache(parent qtcore.QObject_ITF /*777 QObject **/) *QAbstractNetworkCache {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractNetworkCacheC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractNetworkCacheFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractNetworkCache")
	return gothis
}

// /usr/include/qt/QtNetwork/qabstractnetworkcache.h:134
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QAbstractNetworkCache(QObject *)

/*
Constructs an abstract network cache with the given parent.
*/
func (*QAbstractNetworkCache) NewForInheritp() *QAbstractNetworkCache {
	return NewQAbstractNetworkCachep()
}
func NewQAbstractNetworkCachep() *QAbstractNetworkCache {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractNetworkCacheC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractNetworkCacheFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractNetworkCache")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11383() {
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
