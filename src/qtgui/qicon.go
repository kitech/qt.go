package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QString QIcon::themeName();
extern void* C_ZN5QIcon9themeNameEv(); // 4
  // proto: static void QIcon::setThemeName(const QString & path);
extern void C_ZN5QIcon12setThemeNameERK7QString(void* arg0); // 4
  // proto:  void QIcon::swap(QIcon & other);
extern void C_ZN5QIcon4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QIcon::isDetached();
extern bool C_ZNK5QIcon10isDetachedEv(void* qthis); // 4
  // proto:  qint64 QIcon::cacheKey();
extern int64_t C_ZNK5QIcon8cacheKeyEv(void* qthis); // 4
  // proto: static void QIcon::setThemeSearchPaths(const QStringList & searchpath);
extern void C_ZN5QIcon19setThemeSearchPathsERK11QStringList(void* arg0); // 4
  // proto:  void QIcon::detach();
extern void C_ZN5QIcon6detachEv(void* qthis); // 4
  // proto:  QString QIcon::name();
extern void* C_ZNK5QIcon4nameEv(void* qthis); // 4
  // proto: static QStringList QIcon::themeSearchPaths();
extern void C_ZN5QIcon16themeSearchPathsEv(); // 4
  // proto:  void QIcon::~QIcon();
extern void C_ZN5QIconD2Ev(void* qthis); // 4
  // proto:  void QIcon::QIcon(const QString & fileName);
extern void* C_ZN5QIconC2ERK7QString(void* arg0); // 3
  // proto:  void QIcon::QIcon(const QPixmap & pixmap);
extern void* C_ZN5QIconC2ERK7QPixmap(void* arg0); // 3
  // proto:  void QIcon::QIcon(const QIcon & other);
extern void* C_ZN5QIconC2ERKS_(void* arg0); // 3
  // proto:  void QIcon::QIcon();
extern void* C_ZN5QIconC2Ev(); // 3
  // proto:  void QIcon::QIcon(QIconEngine * engine);
extern void* C_ZN5QIconC2EP11QIconEngine(void* arg0); // 3
  // proto: static bool QIcon::hasThemeIcon(const QString & name);
extern bool C_ZN5QIcon12hasThemeIconERK7QString(void* arg0); // 4
  // proto:  bool QIcon::isNull();
extern bool C_ZNK5QIcon6isNullEv(void* qthis); // 4
  // proto: static QIcon QIcon::fromTheme(const QString & name, const QIcon & fallback);
extern void* C_ZN5QIcon9fromThemeERK7QStringRKS_(void* arg0, void* arg1); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QIcon)=8
type QIcon struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// themeName()
func (this *QIcon) Themename_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN5QIcon9themeNameEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIcon", "themeName", args)
  }

  return
}

// setThemeName(const class QString &)
func (this *QIcon) Setthemename_S(args ...interface{}) () {
  // setThemeName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon12setThemeNameERK7QString
    // invoke: void setThemeName(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QIcon12setThemeNameERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QIcon", "setThemeName", args)
  }

  return
}

// swap(class QIcon &)
func (this *QIcon) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QIcon4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QIcon", "swap", args)
  }

  return
}

// isDetached()
func (this *QIcon) Isdetached(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QIcon10isDetachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIcon", "isDetached", args)
  }

  return
}

// cacheKey()
func (this *QIcon) Cachekey(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QIcon8cacheKeyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIcon", "cacheKey", args)
  }

  return
}

// setThemeSearchPaths(const class QStringList &)
func (this *QIcon) Setthemesearchpaths_S(args ...interface{}) () {
  // setThemeSearchPaths(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon19setThemeSearchPathsERK11QStringList
    // invoke: void setThemeSearchPaths(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN5QIcon19setThemeSearchPathsERK11QStringList(arg0)
  default:
    qtrt.ErrorResolve("QIcon", "setThemeSearchPaths", args)
  }

  return
}

// detach()
func (this *QIcon) Detach(args ...interface{}) () {
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
    C.C_ZN5QIcon6detachEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "detach", args)
  }

  return
}

// name()
func (this *QIcon) Name(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QIcon4nameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIcon", "name", args)
  }

  return
}

// themeSearchPaths()
func (this *QIcon) Themesearchpaths_S(args ...interface{}) () {
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

  return
}

// ~QIcon()
func (this *QIcon) Freeqicon(args ...interface{}) () {
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
    C.C_ZN5QIconD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QIcon", "~QIcon", args)
  }

  return
}

// QIcon(const class QString &)
func NewQIcon(args ...interface{}) *QIcon {
  // QIcon(const class QString &)
  // QIcon(const class QPixmap &)
  // QIcon(const class QIcon &)
  // QIcon()
  // QIcon(class QIconEngine *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
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
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QIconC2ERK7QString(arg0)
    return &QIcon{Qclsinst:qthis}
  case 1:
    // invoke: _ZN5QIconC1ERK7QPixmap
    // invoke: void QIcon(const class QPixmap &)
    var arg0 = args[0].(*QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QIconC2ERK7QPixmap(arg0)
    return &QIcon{Qclsinst:qthis}
  case 2:
    // invoke: _ZN5QIconC1ERKS_
    // invoke: void QIcon(const class QIcon &)
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QIconC2ERKS_(arg0)
    return &QIcon{Qclsinst:qthis}
  case 3:
    // invoke: _ZN5QIconC1Ev
    // invoke: void QIcon()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QIconC2Ev()
    return &QIcon{Qclsinst:qthis}
  case 4:
    // invoke: _ZN5QIconC1EP11QIconEngine
    // invoke: void QIcon(class QIconEngine *)
    var arg0 = args[0].(*QIconEngine).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QIconC2EP11QIconEngine(arg0)
    return &QIcon{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QIcon", "QIcon", args)
  }

  return nil // QIcon{Qclsinst:qthis}
}

// hasThemeIcon(const class QString &)
func (this *QIcon) Hasthemeicon_S(args ...interface{}) (ret interface{}) {
  // hasThemeIcon(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon12hasThemeIconERK7QString
    // invoke: bool hasThemeIcon(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QIcon12hasThemeIconERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIcon", "hasThemeIcon", args)
  }

  return
}

// isNull()
func (this *QIcon) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QIcon6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIcon", "isNull", args)
  }

  return
}

// fromTheme(const class QString &, const class QIcon &)
func (this *QIcon) Fromtheme_S(args ...interface{}) (ret interface{}) {
  // fromTheme(const class QString &, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QIcon9fromThemeERK7QStringRKS_
    // invoke: QIcon fromTheme(const class QString &, const class QIcon &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QIcon).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QIcon9fromThemeERK7QStringRKS_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QIcon", "fromTheme", args)
  }

  return
}

// <= body block end

