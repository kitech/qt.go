package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QString QIcon::themeName();
extern void C_ZN5QIcon9themeNameEv(); // 4
  // proto: static void QIcon::setThemeName(const QString & path);
extern void C_ZN5QIcon12setThemeNameERK7QString(void* arg0); // 4
  // proto:  void QIcon::swap(QIcon & other);
extern void C_ZN5QIcon4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QIcon::isDetached();
extern void C_ZNK5QIcon10isDetachedEv(void* qthis); // 4
  // proto:  qint64 QIcon::cacheKey();
extern void C_ZNK5QIcon8cacheKeyEv(void* qthis); // 4
  // proto: static void QIcon::setThemeSearchPaths(const QStringList & searchpath);
extern void C_ZN5QIcon19setThemeSearchPathsERK11QStringList(void* arg0); // 4
  // proto:  void QIcon::detach();
extern void C_ZN5QIcon6detachEv(void* qthis); // 4
  // proto:  QString QIcon::name();
extern void C_ZNK5QIcon4nameEv(void* qthis); // 4
  // proto: static QStringList QIcon::themeSearchPaths();
extern void C_ZN5QIcon16themeSearchPathsEv(); // 4
  // proto:  void QIcon::~QIcon();
extern void C_ZN5QIconD2Ev(void* qthis); // 4
  // proto:  void QIcon::QIcon(const QString & fileName);
extern void C_ZN5QIconC2ERK7QString(void* qthis, void* arg0); // 3
  // proto:  void QIcon::QIcon(const QPixmap & pixmap);
extern void C_ZN5QIconC2ERK7QPixmap(void* qthis, void* arg0); // 3
  // proto:  void QIcon::QIcon(const QIcon & other);
extern void C_ZN5QIconC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QIcon::QIcon();
extern void C_ZN5QIconC2Ev(void* qthis); // 3
  // proto:  void QIcon::QIcon(QIconEngine * engine);
extern void C_ZN5QIconC2EP11QIconEngine(void* qthis, void* arg0); // 3
  // proto: static bool QIcon::hasThemeIcon(const QString & name);
extern void C_ZN5QIcon12hasThemeIconERK7QString(void* arg0); // 4
  // proto:  bool QIcon::isNull();
extern void C_ZNK5QIcon6isNullEv(void* qthis); // 4
  // proto: static QIcon QIcon::fromTheme(const QString & name, const QIcon & fallback);
extern void C_ZN5QIcon9fromThemeERK7QStringRKS_(void* arg0, void* arg1); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// themeName()
func (this *QIcon) themeName_s(args ...interface{}) () {
  // themeName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon9themeNameEv
    // invoke: QString themeName()
    C.C_ZN5QIcon9themeNameEv()
  default:
    qtrt.ErrorResolve("QIcon", "themeName", args)
  }

}

// setThemeName(const class QString &)
func (this *QIcon) setThemeName_s(args ...interface{}) () {
  // setThemeName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon12setThemeNameERK7QString
    // invoke: void setThemeName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QIcon12setThemeNameERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QIcon", "setThemeName", args)
  }

}

// swap(class QIcon &)
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
    C.C_ZN5QIcon4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIcon", "swap", args)
  }

}

// isDetached()
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
    C.C_ZNK5QIcon10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "isDetached", args)
  }

}

// cacheKey()
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
    C.C_ZNK5QIcon8cacheKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "cacheKey", args)
  }

}

// setThemeSearchPaths(const class QStringList &)
func (this *QIcon) setThemeSearchPaths_s(args ...interface{}) () {
  // setThemeSearchPaths(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon19setThemeSearchPathsERK11QStringList
    // invoke: void setThemeSearchPaths(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QIcon19setThemeSearchPathsERK11QStringList(arg0)
  default:
    qtrt.ErrorResolve("QIcon", "setThemeSearchPaths", args)
  }

}

// detach()
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
    C.C_ZN5QIcon6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "detach", args)
  }

}

// name()
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
    C.C_ZNK5QIcon4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "name", args)
  }

}

// themeSearchPaths()
func (this *QIcon) themeSearchPaths_s(args ...interface{}) () {
  // themeSearchPaths()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon16themeSearchPathsEv
    // invoke: QStringList themeSearchPaths()
    C.C_ZN5QIcon16themeSearchPathsEv()
  default:
    qtrt.ErrorResolve("QIcon", "themeSearchPaths", args)
  }

}

// ~QIcon()
func (this *QIcon) FreeQIcon(args ...interface{}) () {
  // ~QIcon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIconD0Ev
    // invoke: void ~QIcon()
    C.C_ZN5QIconD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "~QIcon", args)
  }

}

// QIcon(const class QString &)
func NewQIcon(args ...interface{}) QIcon {
  // QIcon(const class QString &)
  // QIcon(const class QPixmap &)
  // QIcon(const class QIcon &)
  // QIcon()
  // QIcon(class QIconEngine *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QIconEngine{}) // "QIconEngine *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIconC1ERK7QString
    // invoke: void QIcon(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QIconC2ERK7QString(qthis, arg0)
  case 1:
    // invoke: _ZN5QIconC1ERK7QPixmap
    // invoke: void QIcon(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QIconC2ERK7QPixmap(qthis, arg0)
  case 2:
    // invoke: _ZN5QIconC1ERKS_
    // invoke: void QIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QIconC2ERKS_(qthis, arg0)
  case 3:
    // invoke: _ZN5QIconC1Ev
    // invoke: void QIcon()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QIconC2Ev(qthis)
  case 4:
    // invoke: _ZN5QIconC1EP11QIconEngine
    // invoke: void QIcon(class QIconEngine *)
    var arg0 = args[0].(QIconEngine).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN5QIconC2EP11QIconEngine(qthis, arg0)
  default:
    qtrt.ErrorResolve("QIcon", "QIcon", args)
  }

  return QIcon{}
}

// hasThemeIcon(const class QString &)
func (this *QIcon) hasThemeIcon_s(args ...interface{}) () {
  // hasThemeIcon(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon12hasThemeIconERK7QString
    // invoke: bool hasThemeIcon(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QIcon12hasThemeIconERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QIcon", "hasThemeIcon", args)
  }

}

// isNull()
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
    C.C_ZNK5QIcon6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "isNull", args)
  }

}

// fromTheme(const class QString &, const class QIcon &)
func (this *QIcon) fromTheme_s(args ...interface{}) () {
  // fromTheme(const class QString &, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon9fromThemeERK7QStringRKS_
    // invoke: QIcon fromTheme(const class QString &, const class QIcon &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN5QIcon9fromThemeERK7QStringRKS_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QIcon", "fromTheme", args)
  }

}

// <= body block end

