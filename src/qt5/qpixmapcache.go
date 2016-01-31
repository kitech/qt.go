package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static bool QPixmapCache::insert(const QString & key, const QPixmap & pixmap);
extern void C_ZN12QPixmapCache6insertERK7QStringRK7QPixmap(void* arg0, void* arg1); // 4
  // proto: static QPixmapCache::Key QPixmapCache::insert(const QPixmap & pixmap);
extern void C_ZN12QPixmapCache6insertERK7QPixmap(void* arg0); // 4
  // proto: static int QPixmapCache::cacheLimit();
extern void C_ZN12QPixmapCache10cacheLimitEv(); // 4
  // proto: static void QPixmapCache::clear();
extern void C_ZN12QPixmapCache5clearEv(); // 4
  // proto: static void QPixmapCache::remove(const QString & key);
extern void C_ZN12QPixmapCache6removeERK7QString(void* arg0); // 4
  // proto: static void QPixmapCache::setCacheLimit(int );
extern void C_ZN12QPixmapCache13setCacheLimitEi(int32_t arg0); // 4
  // proto: static QPixmap * QPixmapCache::find(const QString & key);
extern void C_ZN12QPixmapCache4findERK7QString(void* arg0); // 4
  // proto: static bool QPixmapCache::find(const QString & key, QPixmap & pixmap);
extern void C_ZN12QPixmapCache4findERK7QStringR7QPixmap(void* arg0, void* arg1); // 4
  // proto: static bool QPixmapCache::find(const QString & key, QPixmap * pixmap);
extern void C_ZN12QPixmapCache4findERK7QStringP7QPixmap(void* arg0, void* arg1); // 4
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

// insert(const class QString &, const class QPixmap &)
func (this *QPixmapCache) insert_s(args ...interface{}) () {
  // insert(const class QString &, const class QPixmap &)
  // insert(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPixmapCache6insertERK7QStringRK7QPixmap
    // invoke: bool insert(const class QString &, const class QPixmap &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QPixmapCache6insertERK7QStringRK7QPixmap(arg0, arg1)
  case 1:
    // invoke: _ZN12QPixmapCache6insertERK7QPixmap
    // invoke: QPixmapCache::Key insert(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPixmapCache6insertERK7QPixmap(arg0)
  default:
    qtrt.ErrorResolve("QPixmapCache", "insert", args)
  }

}

// cacheLimit()
func (this *QPixmapCache) cacheLimit_s(args ...interface{}) () {
  // cacheLimit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPixmapCache10cacheLimitEv
    // invoke: int cacheLimit()
    C.C_ZN12QPixmapCache10cacheLimitEv()
  default:
    qtrt.ErrorResolve("QPixmapCache", "cacheLimit", args)
  }

}

// clear()
func (this *QPixmapCache) clear_s(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPixmapCache5clearEv
    // invoke: void clear()
    C.C_ZN12QPixmapCache5clearEv()
  default:
    qtrt.ErrorResolve("QPixmapCache", "clear", args)
  }

}

// remove(const class QString &)
func (this *QPixmapCache) remove_s(args ...interface{}) () {
  // remove(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPixmapCache6removeERK7QString
    // invoke: void remove(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPixmapCache6removeERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QPixmapCache", "remove", args)
  }

}

// setCacheLimit(int)
func (this *QPixmapCache) setCacheLimit_s(args ...interface{}) () {
  // setCacheLimit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPixmapCache13setCacheLimitEi
    // invoke: void setCacheLimit(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QPixmapCache13setCacheLimitEi(arg0)
  default:
    qtrt.ErrorResolve("QPixmapCache", "setCacheLimit", args)
  }

}

// find(const class QString &)
func (this *QPixmapCache) find_s(args ...interface{}) () {
  // find(const class QString &)
  // find(const class QString &, class QPixmap &)
  // find(const class QString &, class QPixmap *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QPixmap{}) // "QPixmap &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QPixmap{}) // "QPixmap *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QPixmapCache4findERK7QString
    // invoke: QPixmap * find(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QPixmapCache4findERK7QString(arg0)
  case 1:
    // invoke: _ZN12QPixmapCache4findERK7QStringR7QPixmap
    // invoke: bool find(const class QString &, class QPixmap &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QPixmapCache4findERK7QStringR7QPixmap(arg0, arg1)
  case 2:
    // invoke: _ZN12QPixmapCache4findERK7QStringP7QPixmap
    // invoke: bool find(const class QString &, class QPixmap *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QPixmapCache4findERK7QStringP7QPixmap(arg0, arg1)
  default:
    qtrt.ErrorResolve("QPixmapCache", "find", args)
  }

}

// <= body block end

