package qtgui

// /usr/include/qt/QtGui/qpixmapcache.h
// #include <qpixmapcache.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QPixmapCache struct {
	*qtrt.CObject
}
type QPixmapCache_ITF interface {
	QPixmapCache_PTR() *QPixmapCache
}

func (ptr *QPixmapCache) QPixmapCache_PTR() *QPixmapCache { return ptr }

func (this *QPixmapCache) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPixmapCache) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPixmapCacheFromPointer(cthis unsafe.Pointer) *QPixmapCache {
	return &QPixmapCache{&qtrt.CObject{cthis}}
}
func (*QPixmapCache) NewFromPointer(cthis unsafe.Pointer) *QPixmapCache {
	return NewQPixmapCacheFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpixmapcache.h:77
// index:0
// Public static Visibility=Default Availability=Available
// [4] int cacheLimit()

/*
Returns the cache limit (in kilobytes).

The default cache limit is 10240 KB.

See also setCacheLimit().
*/
func (this *QPixmapCache) CacheLimit() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixmapCache10cacheLimitEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QPixmapCache_CacheLimit() int {
	var nilthis *QPixmapCache
	rv := nilthis.CacheLimit()
	return rv
}

// /usr/include/qt/QtGui/qpixmapcache.h:78
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setCacheLimit(int)

/*
Sets the cache limit to n kilobytes.

The default setting is 10240 KB.

See also cacheLimit().
*/
func (this *QPixmapCache) SetCacheLimit(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixmapCache13setCacheLimitEi", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QPixmapCache_SetCacheLimit(arg0 int) {
	var nilthis *QPixmapCache
	nilthis.SetCacheLimit(arg0)
}

// /usr/include/qt/QtGui/qpixmapcache.h:79
// index:0
// Public static Visibility=Default Availability=Available
// [8] QPixmap * find(const QString &)

/*
Looks for a cached pixmap associated with the given key in the cache. If the pixmap is found, the function sets pixmap to that pixmap and returns true; otherwise it leaves pixmap alone and returns false.

Example:


  QPixmap pm;
  if (!QPixmapCache::find("my_big_image", &pm)) {
      pm.load("bigimage.png");
      QPixmapCache::insert("my_big_image", pm);
  }
  painter->drawPixmap(0, 0, pm);



This function was introduced in  Qt 4.6.
*/
func (this *QPixmapCache) Find(key string) *QPixmap /*777 QPixmap **/ {
	var tmpArg0 = qtcore.NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QPixmapCache_Find(key string) *QPixmap /*777 QPixmap **/ {
	var nilthis *QPixmapCache
	rv := nilthis.Find(key)
	return rv
}

// /usr/include/qt/QtGui/qpixmapcache.h:80
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool find(const QString &, QPixmap &)

/*
Looks for a cached pixmap associated with the given key in the cache. If the pixmap is found, the function sets pixmap to that pixmap and returns true; otherwise it leaves pixmap alone and returns false.

Example:


  QPixmap pm;
  if (!QPixmapCache::find("my_big_image", &pm)) {
      pm.load("bigimage.png");
      QPixmapCache::insert("my_big_image", pm);
  }
  painter->drawPixmap(0, 0, pm);



This function was introduced in  Qt 4.6.
*/
func (this *QPixmapCache) Find1(key string, pixmap QPixmap_ITF) bool {
	var tmpArg0 = qtcore.NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QStringR7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QPixmapCache_Find1(key string, pixmap QPixmap_ITF) bool {
	var nilthis *QPixmapCache
	rv := nilthis.Find1(key, pixmap)
	return rv
}

// /usr/include/qt/QtGui/qpixmapcache.h:81
// index:2
// Public static Visibility=Default Availability=Available
// [1] bool find(const QString &, QPixmap *)

/*
Looks for a cached pixmap associated with the given key in the cache. If the pixmap is found, the function sets pixmap to that pixmap and returns true; otherwise it leaves pixmap alone and returns false.

Example:


  QPixmap pm;
  if (!QPixmapCache::find("my_big_image", &pm)) {
      pm.load("bigimage.png");
      QPixmapCache::insert("my_big_image", pm);
  }
  painter->drawPixmap(0, 0, pm);



This function was introduced in  Qt 4.6.
*/
func (this *QPixmapCache) Find2(key string, pixmap QPixmap_ITF /*777 QPixmap **/) bool {
	var tmpArg0 = qtcore.NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixmapCache4findERK7QStringP7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QPixmapCache_Find2(key string, pixmap QPixmap_ITF /*777 QPixmap **/) bool {
	var nilthis *QPixmapCache
	rv := nilthis.Find2(key, pixmap)
	return rv
}

// /usr/include/qt/QtGui/qpixmapcache.h:86
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void remove(const QString &)

/*
Removes the pixmap associated with key from the cache.
*/
func (this *QPixmapCache) Remove(key string) {
	var tmpArg0 = qtcore.NewQString5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixmapCache6removeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QPixmapCache_Remove(key string) {
	var nilthis *QPixmapCache
	nilthis.Remove(key)
}

// /usr/include/qt/QtGui/qpixmapcache.h:88
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all pixmaps from the cache.
*/
func (this *QPixmapCache) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixmapCache5clearEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QPixmapCache_Clear() {
	var nilthis *QPixmapCache
	nilthis.Clear()
}

func DeleteQPixmapCache(this *QPixmapCache) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QPixmapCacheD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10927() {
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
