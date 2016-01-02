package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qfontdatabase.h
// dst-file: /src/gui/qfontdatabase.go
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
  // proto:  QList<int> QFontDatabase::pointSizes(const QString & family, const QString & style);
extern void _ZN13QFontDatabase10pointSizesERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QString QFontDatabase::styleString(const QFont & font);
extern void _ZN13QFontDatabase11styleStringERK5QFont(void* qthis, void* arg0);
  // proto:  QList<int> QFontDatabase::smoothSizes(const QString & family, const QString & style);
extern void _ZN13QFontDatabase11smoothSizesERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QStringList QFontDatabase::styles(const QString & family);
extern void _ZNK13QFontDatabase6stylesERK7QString(void* qthis, void* arg0);
  // proto:  bool QFontDatabase::italic(const QString & family, const QString & style);
extern void _ZNK13QFontDatabase6italicERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QFontDatabase::QFontDatabase();
extern void* dector_ZN13QFontDatabaseC1Ev();
extern void _ZN13QFontDatabaseC1Ev(void* qthis);
  // proto: static QStringList QFontDatabase::applicationFontFamilies(int id);
extern void _ZN13QFontDatabase23applicationFontFamiliesEi(int arg0);
  // proto:  bool QFontDatabase::hasFamily(const QString & family);
extern void _ZNK13QFontDatabase9hasFamilyERK7QString(void* qthis, void* arg0);
  // proto:  bool QFontDatabase::isFixedPitch(const QString & family, const QString & style);
extern void _ZNK13QFontDatabase12isFixedPitchERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QFont QFontDatabase::font(const QString & family, const QString & style, int pointSize);
extern void _ZNK13QFontDatabase4fontERK7QStringS2_i(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  int QFontDatabase::weight(const QString & family, const QString & style);
extern void _ZNK13QFontDatabase6weightERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto: static bool QFontDatabase::removeAllApplicationFonts();
extern void _ZN13QFontDatabase25removeAllApplicationFontsEv();
  // proto: static int QFontDatabase::addApplicationFontFromData(const QByteArray & fontData);
extern void _ZN13QFontDatabase26addApplicationFontFromDataERK10QByteArray(void* arg0);
  // proto: static bool QFontDatabase::supportsThreadedFontRendering();
extern void _ZN13QFontDatabase29supportsThreadedFontRenderingEv();
  // proto:  bool QFontDatabase::isPrivateFamily(const QString & family);
extern void _ZNK13QFontDatabase15isPrivateFamilyERK7QString(void* qthis, void* arg0);
  // proto:  bool QFontDatabase::isScalable(const QString & family, const QString & style);
extern void _ZNK13QFontDatabase10isScalableERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto: static bool QFontDatabase::removeApplicationFont(int id);
extern void _ZN13QFontDatabase21removeApplicationFontEi(int arg0);
  // proto:  QString QFontDatabase::styleString(const QFontInfo & fontInfo);
extern void _ZN13QFontDatabase11styleStringERK9QFontInfo(void* qthis, void* arg0);
  // proto:  bool QFontDatabase::isBitmapScalable(const QString & family, const QString & style);
extern void _ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  bool QFontDatabase::isSmoothlyScalable(const QString & family, const QString & style);
extern void _ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  bool QFontDatabase::bold(const QString & family, const QString & style);
extern void _ZNK13QFontDatabase4boldERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto: static int QFontDatabase::addApplicationFont(const QString & fileName);
extern void _ZN13QFontDatabase18addApplicationFontERK7QString(void* arg0);
  // proto: static QList<int> QFontDatabase::standardSizes();
extern void _ZN13QFontDatabase13standardSizesEv();
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

// class sizeof(QFontDatabase)=8
type QFontDatabase struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QList<int> QFontDatabase::pointSizes(const QString & family, const QString & style);
func (this *QFontDatabase) pointSizes(args ...interface{}) () {
  // pointSizes(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase10pointSizesERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "pointSizes", args)
  }

}

  // proto:  QString QFontDatabase::styleString(const QFont & font);
func (this *QFontDatabase) styleString(args ...interface{}) () {
  // styleString(const class QFont &)
  // styleString(const class QFontInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFontInfo{}) // "const QFontInfo &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase11styleStringERK5QFont
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN13QFontDatabase11styleStringERK9QFontInfo
    var arg0 = args[0].(QFontInfo).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "styleString", args)
  }

}

  // proto:  QList<int> QFontDatabase::smoothSizes(const QString & family, const QString & style);
func (this *QFontDatabase) smoothSizes(args ...interface{}) () {
  // smoothSizes(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase11smoothSizesERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "smoothSizes", args)
  }

}

  // proto:  QStringList QFontDatabase::styles(const QString & family);
func (this *QFontDatabase) styles(args ...interface{}) () {
  // styles(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase6stylesERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "styles", args)
  }

}

  // proto:  bool QFontDatabase::italic(const QString & family, const QString & style);
func (this *QFontDatabase) italic(args ...interface{}) () {
  // italic(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase6italicERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "italic", args)
  }

}

  // proto:  void QFontDatabase::QFontDatabase();
func NewQFontDatabase(args ...interface{}) QFontDatabase {
  return QFontDatabase{}
}

  // proto: static QStringList QFontDatabase::applicationFontFamilies(int id);
func (this *QFontDatabase) applicationFontFamilies_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontDatabase", "applicationFontFamilies", args)
  }

}

  // proto:  bool QFontDatabase::hasFamily(const QString & family);
func (this *QFontDatabase) hasFamily(args ...interface{}) () {
  // hasFamily(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase9hasFamilyERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "hasFamily", args)
  }

}

  // proto:  bool QFontDatabase::isFixedPitch(const QString & family, const QString & style);
func (this *QFontDatabase) isFixedPitch(args ...interface{}) () {
  // isFixedPitch(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase12isFixedPitchERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "isFixedPitch", args)
  }

}

  // proto:  QFont QFontDatabase::font(const QString & family, const QString & style, int pointSize);
func (this *QFontDatabase) font(args ...interface{}) () {
  // font(const class QString &, const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase4fontERK7QStringS2_i
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "font", args)
  }

}

  // proto:  int QFontDatabase::weight(const QString & family, const QString & style);
func (this *QFontDatabase) weight(args ...interface{}) () {
  // weight(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase6weightERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "weight", args)
  }

}

  // proto: static bool QFontDatabase::removeAllApplicationFonts();
func (this *QFontDatabase) removeAllApplicationFonts_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontDatabase", "removeAllApplicationFonts", args)
  }

}

  // proto: static int QFontDatabase::addApplicationFontFromData(const QByteArray & fontData);
func (this *QFontDatabase) addApplicationFontFromData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontDatabase", "addApplicationFontFromData", args)
  }

}

  // proto: static bool QFontDatabase::supportsThreadedFontRendering();
func (this *QFontDatabase) supportsThreadedFontRendering_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontDatabase", "supportsThreadedFontRendering", args)
  }

}

  // proto:  bool QFontDatabase::isPrivateFamily(const QString & family);
func (this *QFontDatabase) isPrivateFamily(args ...interface{}) () {
  // isPrivateFamily(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase15isPrivateFamilyERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "isPrivateFamily", args)
  }

}

  // proto:  bool QFontDatabase::isScalable(const QString & family, const QString & style);
func (this *QFontDatabase) isScalable(args ...interface{}) () {
  // isScalable(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase10isScalableERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "isScalable", args)
  }

}

  // proto: static bool QFontDatabase::removeApplicationFont(int id);
func (this *QFontDatabase) removeApplicationFont_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontDatabase", "removeApplicationFont", args)
  }

}

  // proto:  bool QFontDatabase::isBitmapScalable(const QString & family, const QString & style);
func (this *QFontDatabase) isBitmapScalable(args ...interface{}) () {
  // isBitmapScalable(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "isBitmapScalable", args)
  }

}

  // proto:  bool QFontDatabase::isSmoothlyScalable(const QString & family, const QString & style);
func (this *QFontDatabase) isSmoothlyScalable(args ...interface{}) () {
  // isSmoothlyScalable(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "isSmoothlyScalable", args)
  }

}

  // proto:  bool QFontDatabase::bold(const QString & family, const QString & style);
func (this *QFontDatabase) bold(args ...interface{}) () {
  // bold(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase4boldERK7QStringS2_
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFontDatabase", "bold", args)
  }

}

  // proto: static int QFontDatabase::addApplicationFont(const QString & fileName);
func (this *QFontDatabase) addApplicationFont_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontDatabase", "addApplicationFont", args)
  }

}

  // proto: static QList<int> QFontDatabase::standardSizes();
func (this *QFontDatabase) standardSizes_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontDatabase", "standardSizes", args)
  }

}

// <= body block end

