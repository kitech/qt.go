package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtGui/qpixmapcache.h
// dst-file: /src/gui/qpixmapcache.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static int QPixmapCache::cacheLimit();
extern void _ZN12QPixmapCache10cacheLimitEv();
  // proto: static void QPixmapCache::clear();
extern void _ZN12QPixmapCache5clearEv();
  // proto: static bool QPixmapCache::insert(const QString & key, const QPixmap & pixmap);
extern void _ZN12QPixmapCache6insertERK7QStringRK7QPixmap(void* arg0, void* arg1);
  // proto: static bool QPixmapCache::find(const QString & key, QPixmap & pixmap);
extern void _ZN12QPixmapCache4findERK7QStringR7QPixmap(void* arg0, void* arg1);
  // proto: static QPixmap * QPixmapCache::find(const QString & key);
extern void _ZN12QPixmapCache4findERK7QString(void* arg0);
  // proto: static bool QPixmapCache::find(const QString & key, QPixmap * pixmap);
extern void _ZN12QPixmapCache4findERK7QStringP7QPixmap(void* arg0, void* arg1);
  // proto: static void QPixmapCache::remove(const QString & key);
extern void _ZN12QPixmapCache6removeERK7QString(void* arg0);
  // proto: static void QPixmapCache::setCacheLimit(int );
extern void _ZN12QPixmapCache13setCacheLimitEi(int32_t arg0);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QPixmapCache)=1
type QPixmapCache struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto: static int QPixmapCache::cacheLimit();
func (this *QPixmapCache) cacheLimit_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmapCache", "cacheLimit", args)
  }

}

  // proto: static void QPixmapCache::clear();
func (this *QPixmapCache) clear_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmapCache", "clear", args)
  }

}

  // proto: static bool QPixmapCache::insert(const QString & key, const QPixmap & pixmap);
func (this *QPixmapCache) insert_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmapCache", "insert", args)
  }

}

  // proto: static bool QPixmapCache::find(const QString & key, QPixmap & pixmap);
func (this *QPixmapCache) find_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmapCache", "find", args)
  }

}

  // proto: static void QPixmapCache::remove(const QString & key);
func (this *QPixmapCache) remove_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmapCache", "remove", args)
  }

}

  // proto: static void QPixmapCache::setCacheLimit(int );
func (this *QPixmapCache) setCacheLimit_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmapCache", "setCacheLimit", args)
  }

}

// <= body block end

