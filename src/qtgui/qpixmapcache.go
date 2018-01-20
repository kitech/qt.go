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

// /usr/include/qt/QtGui/qpixmapcache.h:77
// index:0
// static
// int cacheLimit()
func (this *QPixmapCache) CacheLimit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache10cacheLimitEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_CacheLimit() {
	// 0: (), ()
	var nilthis *QPixmapCache
	nilthis.CacheLimit()
}

// /usr/include/qt/QtGui/qpixmapcache.h:78
// index:0
// static
// void setCacheLimit(int)
func (this *QPixmapCache) SetCacheLimit(arg0 int) {
	// 0: (int arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache13setCacheLimitEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_SetCacheLimit(arg0 int) {
	// 0: (int arg0), (arg0)
	var nilthis *QPixmapCache
	nilthis.SetCacheLimit(arg0)
}

// /usr/include/qt/QtGui/qpixmapcache.h:79
// index:0
// static
// QPixmap * find(const class QString &)
func (this *QPixmapCache) Find(key unsafe.Pointer) {
	// 0: (key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_Find(key unsafe.Pointer) {
	// 0: (key const QString &), (key)
	var nilthis *QPixmapCache
	nilthis.Find(key)
}

// /usr/include/qt/QtGui/qpixmapcache.h:80
// index:1
// static
// bool find(const class QString &, class QPixmap &)
func (this *QPixmapCache) Find_1(key unsafe.Pointer, pixmap unsafe.Pointer) {
	// 1: (key const QString &, pixmap QPixmap &), (key, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QStringR7QPixmap", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_Find_1(key unsafe.Pointer, pixmap unsafe.Pointer) {
	// 1: (key const QString &, pixmap QPixmap &), (key, pixmap)
	var nilthis *QPixmapCache
	nilthis.Find_1(key, pixmap)
}

// /usr/include/qt/QtGui/qpixmapcache.h:81
// index:2
// static
// bool find(const class QString &, class QPixmap *)
func (this *QPixmapCache) Find_2(key unsafe.Pointer, pixmap unsafe.Pointer) {
	// 2: (key const QString &, pixmap QPixmap *), (key, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QStringP7QPixmap", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_Find_2(key unsafe.Pointer, pixmap unsafe.Pointer) {
	// 2: (key const QString &, pixmap QPixmap *), (key, pixmap)
	var nilthis *QPixmapCache
	nilthis.Find_2(key, pixmap)
}

// /usr/include/qt/QtGui/qpixmapcache.h:82
// index:3
// static
// bool find(const class QPixmapCache::Key &, class QPixmap *)
func (this *QPixmapCache) Find_3(key unsafe.Pointer, pixmap unsafe.Pointer) {
	// 3: (key const QPixmapCache::Key &, pixmap QPixmap *), (key, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache4findERKNS_3KeyEP7QPixmap", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_Find_3(key unsafe.Pointer, pixmap unsafe.Pointer) {
	// 3: (key const QPixmapCache::Key &, pixmap QPixmap *), (key, pixmap)
	var nilthis *QPixmapCache
	nilthis.Find_3(key, pixmap)
}

// /usr/include/qt/QtGui/qpixmapcache.h:85
// index:0
// static
// bool replace(const class QPixmapCache::Key &, const class QPixmap &)
func (this *QPixmapCache) Replace(key unsafe.Pointer, pixmap unsafe.Pointer) {
	// 0: (key const QPixmapCache::Key &, pixmap const QPixmap &), (key, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache7replaceERKNS_3KeyERK7QPixmap", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_Replace(key unsafe.Pointer, pixmap unsafe.Pointer) {
	// 0: (key const QPixmapCache::Key &, pixmap const QPixmap &), (key, pixmap)
	var nilthis *QPixmapCache
	nilthis.Replace(key, pixmap)
}

// /usr/include/qt/QtGui/qpixmapcache.h:86
// index:0
// static
// void remove(const class QString &)
func (this *QPixmapCache) Remove(key unsafe.Pointer) {
	// 0: (key const QString &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache6removeERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_Remove(key unsafe.Pointer) {
	// 0: (key const QString &), (key)
	var nilthis *QPixmapCache
	nilthis.Remove(key)
}

// /usr/include/qt/QtGui/qpixmapcache.h:87
// index:1
// static
// void remove(const class QPixmapCache::Key &)
func (this *QPixmapCache) Remove_1(key unsafe.Pointer) {
	// 1: (key const QPixmapCache::Key &), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache6removeERKNS_3KeyE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_Remove_1(key unsafe.Pointer) {
	// 1: (key const QPixmapCache::Key &), (key)
	var nilthis *QPixmapCache
	nilthis.Remove_1(key)
}

// /usr/include/qt/QtGui/qpixmapcache.h:88
// index:0
// static
// void clear()
func (this *QPixmapCache) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache5clearEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPixmapCache_Clear() {
	// 0: (), ()
	var nilthis *QPixmapCache
	nilthis.Clear()
}

//  body block end
