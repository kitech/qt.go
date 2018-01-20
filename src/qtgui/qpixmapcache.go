//  header block begin
// /usr/include/qt/QtGui/qpixmapcache.h
// #include <qpixmapcache.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
type QPixmapCache struct {
	*qtrt.CObject
}

func (this *QPixmapCache) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQPixmapCacheFromPointer(cthis unsafe.Pointer) *QPixmapCache {
	return &QPixmapCache{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpixmapcache.h:77
// index:0
// Public static
// int cacheLimit()
func (this *QPixmapCache) CacheLimit() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache10cacheLimitEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QPixmapCache_CacheLimit() {
	var nilthis *QPixmapCache
	nilthis.CacheLimit()
}

// /usr/include/qt/QtGui/qpixmapcache.h:78
// index:0
// Public static
// void setCacheLimit(int)
func (this *QPixmapCache) SetCacheLimit(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache13setCacheLimitEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_SetCacheLimit(arg0 int) {
	var nilthis *QPixmapCache
	nilthis.SetCacheLimit(arg0)
}

// /usr/include/qt/QtGui/qpixmapcache.h:79
// index:0
// Public static
// QPixmap * find(const class QString &)
func (this *QPixmapCache) Find(key *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QString", ffiqt.FFI_TYPE_POINTER, key)
	gopp.ErrPrint(err, rv)
	return rv
}
func QPixmapCache_Find(key *qtcore.QString) {
	var nilthis *QPixmapCache
	nilthis.Find(key)
}

// /usr/include/qt/QtGui/qpixmapcache.h:80
// index:1
// Public static
// bool find(const class QString &, class QPixmap &)
func (this *QPixmapCache) Find_1(key *qtcore.QString, pixmap *QPixmap) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QStringR7QPixmap", ffiqt.FFI_TYPE_POINTER, key, pixmap)
	gopp.ErrPrint(err, rv)
	return rv
}
func QPixmapCache_Find_1(key *qtcore.QString, pixmap *QPixmap) {
	var nilthis *QPixmapCache
	nilthis.Find_1(key, pixmap)
}

// /usr/include/qt/QtGui/qpixmapcache.h:81
// index:2
// Public static
// bool find(const class QString &, class QPixmap *)
func (this *QPixmapCache) Find_2(key *qtcore.QString, pixmap unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QStringP7QPixmap", ffiqt.FFI_TYPE_POINTER, key, pixmap)
	gopp.ErrPrint(err, rv)
	return rv
}
func QPixmapCache_Find_2(key *qtcore.QString, pixmap unsafe.Pointer) {
	var nilthis *QPixmapCache
	nilthis.Find_2(key, pixmap)
}

// /usr/include/qt/QtGui/qpixmapcache.h:86
// index:0
// Public static
// void remove(const class QString &)
func (this *QPixmapCache) Remove(key *qtcore.QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache6removeERK7QString", ffiqt.FFI_TYPE_POINTER, key)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_Remove(key *qtcore.QString) {
	var nilthis *QPixmapCache
	nilthis.Remove(key)
}

// /usr/include/qt/QtGui/qpixmapcache.h:88
// index:0
// Public static
// void clear()
func (this *QPixmapCache) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache5clearEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_Clear() {
	var nilthis *QPixmapCache
	nilthis.Clear()
}

//  body block end
