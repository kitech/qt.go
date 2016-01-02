package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qicon.h
// dst-file: /src/gui/qicon.go
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
  // proto:  void QIcon::QIcon(const QIcon & other);
extern void* dector_ZN5QIconC1ERKS_(void* arg0);
extern void _ZN5QIconC1ERKS_(void* qthis, void* arg0);
  // proto: static QStringList QIcon::themeSearchPaths();
extern void _ZN5QIcon16themeSearchPathsEv();
  // proto:  void QIcon::detach();
extern void _ZN5QIcon6detachEv(void* qthis);
  // proto:  bool QIcon::isNull();
extern void _ZNK5QIcon6isNullEv(void* qthis);
  // proto: static void QIcon::setThemeSearchPaths(const QStringList & searchpath);
extern void _ZN5QIcon19setThemeSearchPathsERK11QStringList(void* arg0);
  // proto: static bool QIcon::hasThemeIcon(const QString & name);
extern void _ZN5QIcon12hasThemeIconERK7QString(void* arg0);
  // proto:  void QIcon::QIcon(const QPixmap & pixmap);
extern void* dector_ZN5QIconC1ERK7QPixmap(void* arg0);
extern void _ZN5QIconC1ERK7QPixmap(void* qthis, void* arg0);
  // proto: static QIcon QIcon::fromTheme(const QString & name, const QIcon & fallback);
extern void _ZN5QIcon9fromThemeERK7QStringRKS_(void* arg0, void* arg1);
  // proto: static QString QIcon::themeName();
extern void _ZN5QIcon9themeNameEv();
  // proto:  QString QIcon::name();
extern void _ZNK5QIcon4nameEv(void* qthis);
  // proto:  void QIcon::QIcon();
extern void* dector_ZN5QIconC1Ev();
extern void _ZN5QIconC1Ev(void* qthis);
  // proto:  void QIcon::QIcon(QIconEngine * engine);
extern void* dector_ZN5QIconC1EP11QIconEngine(void* arg0);
extern void _ZN5QIconC1EP11QIconEngine(void* qthis, void* arg0);
  // proto:  void QIcon::~QIcon();
extern void _ZN5QIconD0Ev(void* qthis);
  // proto:  bool QIcon::isDetached();
extern void _ZNK5QIcon10isDetachedEv(void* qthis);
  // proto:  void QIcon::QIcon(const QString & fileName);
extern void* dector_ZN5QIconC1ERK7QString(void* arg0);
extern void _ZN5QIconC1ERK7QString(void* qthis, void* arg0);
  // proto:  qint64 QIcon::cacheKey();
extern void _ZNK5QIcon8cacheKeyEv(void* qthis);
  // proto:  void QIcon::swap(QIcon & other);
extern void demth_ZN5QIcon4swapERS_(void* qthis, void* arg0);
  // proto: static void QIcon::setThemeName(const QString & path);
extern void _ZN5QIcon12setThemeNameERK7QString(void* arg0);
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

// class sizeof(QIcon)=8
type QIcon struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QIcon::QIcon(const QIcon & other);
func NewQIcon(args ...interface{}) QIcon {
  return QIcon{}
}

  // proto: static QStringList QIcon::themeSearchPaths();
func (this *QIcon) themeSearchPaths_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIcon", "themeSearchPaths", args)
  }

}

  // proto:  void QIcon::detach();
func (this *QIcon) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon6detachEv
    // invoke: void detach()
    C._ZN5QIcon6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "detach", args)
  }

}

  // proto:  bool QIcon::isNull();
func (this *QIcon) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QIcon6isNullEv
    // invoke: bool isNull()
    C._ZNK5QIcon6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "isNull", args)
  }

}

  // proto: static void QIcon::setThemeSearchPaths(const QStringList & searchpath);
func (this *QIcon) setThemeSearchPaths_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIcon", "setThemeSearchPaths", args)
  }

}

  // proto: static bool QIcon::hasThemeIcon(const QString & name);
func (this *QIcon) hasThemeIcon_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIcon", "hasThemeIcon", args)
  }

}

  // proto: static QIcon QIcon::fromTheme(const QString & name, const QIcon & fallback);
func (this *QIcon) fromTheme_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIcon", "fromTheme", args)
  }

}

  // proto: static QString QIcon::themeName();
func (this *QIcon) themeName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIcon", "themeName", args)
  }

}

  // proto:  QString QIcon::name();
func (this *QIcon) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QIcon4nameEv
    // invoke: QString name()
    C._ZNK5QIcon4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "name", args)
  }

}

  // proto:  void QIcon::~QIcon();
func (this *QIcon) FreeQIcon(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIcon", "~QIcon", args)
  }

}

  // proto:  bool QIcon::isDetached();
func (this *QIcon) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QIcon10isDetachedEv
    // invoke: bool isDetached()
    C._ZNK5QIcon10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "isDetached", args)
  }

}

  // proto:  qint64 QIcon::cacheKey();
func (this *QIcon) cacheKey(args ...interface{}) () {
  // cacheKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QIcon8cacheKeyEv
    // invoke: qint64 cacheKey()
    C._ZNK5QIcon8cacheKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "cacheKey", args)
  }

}

  // proto:  void QIcon::swap(QIcon & other);
func (this *QIcon) swap(args ...interface{}) () {
  // swap(class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon4swapERS_
    // invoke: void swap(class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN5QIcon4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIcon", "swap", args)
  }

}

  // proto: static void QIcon::setThemeName(const QString & path);
func (this *QIcon) setThemeName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIcon", "setThemeName", args)
  }

}

// <= body block end

