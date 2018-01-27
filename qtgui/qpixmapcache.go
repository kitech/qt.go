package qtgui

// /usr/include/qt/QtGui/qpixmapcache.h
// #include <qpixmapcache.h>
// #include <QtGui>

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
type QPixmapCache struct {
	*qtrt.CObject
}

func (this *QPixmapCache) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPixmapCache) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPixmapCacheFromPointer(cthis unsafe.Pointer) *QPixmapCache {
	return &QPixmapCache{&qtrt.CObject{cthis}}
}
func (*QPixmapCache) NewFromPointer(cthis unsafe.Pointer) *QPixmapCache {
	return NewQPixmapCacheFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpixmapcache.h:77
// index:0
// Public static
// int cacheLimit()
func (this *QPixmapCache) CacheLimit() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache10cacheLimitEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QPixmapCache_CacheLimit() int {
	var nilthis *QPixmapCache
	rv := nilthis.CacheLimit()
	return rv
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
// QPixmap * find(const QString &)
func (this *QPixmapCache) Find(key *qtcore.QString) *QPixmap /*777 QPixmap **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QString", ffiqt.FFI_TYPE_POINTER, key)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QPixmapCache_Find(key *qtcore.QString) *QPixmap /*777 QPixmap **/ {
	var nilthis *QPixmapCache
	rv := nilthis.Find(key)
	return rv
}

// /usr/include/qt/QtGui/qpixmapcache.h:80
// index:1
// Public static
// bool find(const QString &, QPixmap &)
func (this *QPixmapCache) Find_1(key *qtcore.QString, pixmap *QPixmap) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QStringR7QPixmap", ffiqt.FFI_TYPE_POINTER, key, pixmap)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QPixmapCache_Find_1(key *qtcore.QString, pixmap *QPixmap) bool {
	var nilthis *QPixmapCache
	rv := nilthis.Find_1(key, pixmap)
	return rv
}

// /usr/include/qt/QtGui/qpixmapcache.h:81
// index:2
// Public static
// bool find(const QString &, QPixmap *)
func (this *QPixmapCache) Find_2(key *qtcore.QString, pixmap *QPixmap /*777 QPixmap **/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QStringP7QPixmap", ffiqt.FFI_TYPE_POINTER, key, pixmap)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QPixmapCache_Find_2(key *qtcore.QString, pixmap *QPixmap /*777 QPixmap **/) bool {
	var nilthis *QPixmapCache
	rv := nilthis.Find_2(key, pixmap)
	return rv
}

// /usr/include/qt/QtGui/qpixmapcache.h:86
// index:0
// Public static
// void remove(const QString &)
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
