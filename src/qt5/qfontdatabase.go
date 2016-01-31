package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static int QFontDatabase::addApplicationFontFromData(const QByteArray & fontData);
extern void C_ZN13QFontDatabase26addApplicationFontFromDataERK10QByteArray(void* arg0); // 4
  // proto:  int QFontDatabase::weight(const QString & family, const QString & style);
extern void C_ZNK13QFontDatabase6weightERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QFontDatabase::isFixedPitch(const QString & family, const QString & style);
extern void C_ZNK13QFontDatabase12isFixedPitchERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QFont QFontDatabase::font(const QString & family, const QString & style, int pointSize);
extern void C_ZNK13QFontDatabase4fontERK7QStringS2_i(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  QList<int> QFontDatabase::pointSizes(const QString & family, const QString & style);
extern void C_ZN13QFontDatabase10pointSizesERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QList<QFontDatabase::WritingSystem> QFontDatabase::writingSystems();
extern void C_ZNK13QFontDatabase14writingSystemsEv(void* qthis); // 4
  // proto:  QList<QFontDatabase::WritingSystem> QFontDatabase::writingSystems(const QString & family);
extern void C_ZNK13QFontDatabase14writingSystemsERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QFontDatabase::isPrivateFamily(const QString & family);
extern void C_ZNK13QFontDatabase15isPrivateFamilyERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFontDatabase::QFontDatabase();
extern void C_ZN13QFontDatabaseC2Ev(void* qthis); // 3
  // proto:  QString QFontDatabase::styleString(const QFontInfo & fontInfo);
extern void C_ZN13QFontDatabase11styleStringERK9QFontInfo(void* qthis, void* arg0); // 4
  // proto:  QString QFontDatabase::styleString(const QFont & font);
extern void C_ZN13QFontDatabase11styleStringERK5QFont(void* qthis, void* arg0); // 4
  // proto:  bool QFontDatabase::italic(const QString & family, const QString & style);
extern void C_ZNK13QFontDatabase6italicERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QStringList QFontDatabase::styles(const QString & family);
extern void C_ZNK13QFontDatabase6stylesERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QFontDatabase::isSmoothlyScalable(const QString & family, const QString & style);
extern void C_ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QFontDatabase::bold(const QString & family, const QString & style);
extern void C_ZNK13QFontDatabase4boldERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QFontDatabase::isScalable(const QString & family, const QString & style);
extern void C_ZNK13QFontDatabase10isScalableERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto: static int QFontDatabase::addApplicationFont(const QString & fileName);
extern void C_ZN13QFontDatabase18addApplicationFontERK7QString(void* arg0); // 4
  // proto:  bool QFontDatabase::isBitmapScalable(const QString & family, const QString & style);
extern void C_ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto: static QStringList QFontDatabase::applicationFontFamilies(int id);
extern void C_ZN13QFontDatabase23applicationFontFamiliesEi(int32_t arg0); // 4
  // proto: static bool QFontDatabase::supportsThreadedFontRendering();
extern void C_ZN13QFontDatabase29supportsThreadedFontRenderingEv(); // 4
  // proto:  QList<int> QFontDatabase::smoothSizes(const QString & family, const QString & style);
extern void C_ZN13QFontDatabase11smoothSizesERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto: static bool QFontDatabase::removeApplicationFont(int id);
extern void C_ZN13QFontDatabase21removeApplicationFontEi(int32_t arg0); // 4
  // proto:  bool QFontDatabase::hasFamily(const QString & family);
extern void C_ZNK13QFontDatabase9hasFamilyERK7QString(void* qthis, void* arg0); // 4
  // proto: static QList<int> QFontDatabase::standardSizes();
extern void C_ZN13QFontDatabase13standardSizesEv(); // 4
  // proto: static bool QFontDatabase::removeAllApplicationFonts();
extern void C_ZN13QFontDatabase25removeAllApplicationFontsEv(); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// addApplicationFontFromData(const class QByteArray &)
func (this *QFontDatabase) addApplicationFontFromData_s(args ...interface{}) () {
  // addApplicationFontFromData(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase26addApplicationFontFromDataERK10QByteArray
    // invoke: int addApplicationFontFromData(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QFontDatabase26addApplicationFontFromDataERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "addApplicationFontFromData", args)
  }

}

// weight(const class QString &, const class QString &)
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
    // invoke: int weight(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK13QFontDatabase6weightERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "weight", args)
  }

}

// isFixedPitch(const class QString &, const class QString &)
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
    // invoke: bool isFixedPitch(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK13QFontDatabase12isFixedPitchERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "isFixedPitch", args)
  }

}

// font(const class QString &, const class QString &, int)
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
    // invoke: QFont font(const class QString &, const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZNK13QFontDatabase4fontERK7QStringS2_i(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QFontDatabase", "font", args)
  }

}

// pointSizes(const class QString &, const class QString &)
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
    // invoke: QList<int> pointSizes(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QFontDatabase10pointSizesERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "pointSizes", args)
  }

}

// writingSystems()
func (this *QFontDatabase) writingSystems(args ...interface{}) () {
  // writingSystems()
  // writingSystems(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontDatabase14writingSystemsEv
    // invoke: QList<QFontDatabase::WritingSystem> writingSystems()
    C.C_ZNK13QFontDatabase14writingSystemsEv(this.qclsinst)
  case 1:
    // invoke: _ZNK13QFontDatabase14writingSystemsERK7QString
    // invoke: QList<QFontDatabase::WritingSystem> writingSystems(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QFontDatabase14writingSystemsERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "writingSystems", args)
  }

}

// isPrivateFamily(const class QString &)
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
    // invoke: bool isPrivateFamily(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QFontDatabase15isPrivateFamilyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "isPrivateFamily", args)
  }

}

// QFontDatabase()
func NewQFontDatabase(args ...interface{}) QFontDatabase {
  // QFontDatabase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabaseC1Ev
    // invoke: void QFontDatabase()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN13QFontDatabaseC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QFontDatabase", "QFontDatabase", args)
  }

  return QFontDatabase{}
}

// styleString(const class QFontInfo &)
func (this *QFontDatabase) styleString(args ...interface{}) () {
  // styleString(const class QFontInfo &)
  // styleString(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontInfo{}) // "const QFontInfo &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase11styleStringERK9QFontInfo
    // invoke: QString styleString(const class QFontInfo &)
    var arg0 = args[0].(QFontInfo).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QFontDatabase11styleStringERK9QFontInfo(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN13QFontDatabase11styleStringERK5QFont
    // invoke: QString styleString(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QFontDatabase11styleStringERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "styleString", args)
  }

}

// italic(const class QString &, const class QString &)
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
    // invoke: bool italic(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK13QFontDatabase6italicERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "italic", args)
  }

}

// styles(const class QString &)
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
    // invoke: QStringList styles(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QFontDatabase6stylesERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "styles", args)
  }

}

// isSmoothlyScalable(const class QString &, const class QString &)
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
    // invoke: bool isSmoothlyScalable(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "isSmoothlyScalable", args)
  }

}

// bold(const class QString &, const class QString &)
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
    // invoke: bool bold(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK13QFontDatabase4boldERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "bold", args)
  }

}

// isScalable(const class QString &, const class QString &)
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
    // invoke: bool isScalable(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK13QFontDatabase10isScalableERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "isScalable", args)
  }

}

// addApplicationFont(const class QString &)
func (this *QFontDatabase) addApplicationFont_s(args ...interface{}) () {
  // addApplicationFont(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase18addApplicationFontERK7QString
    // invoke: int addApplicationFont(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QFontDatabase18addApplicationFontERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "addApplicationFont", args)
  }

}

// isBitmapScalable(const class QString &, const class QString &)
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
    // invoke: bool isBitmapScalable(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "isBitmapScalable", args)
  }

}

// applicationFontFamilies(int)
func (this *QFontDatabase) applicationFontFamilies_s(args ...interface{}) () {
  // applicationFontFamilies(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase23applicationFontFamiliesEi
    // invoke: QStringList applicationFontFamilies(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QFontDatabase23applicationFontFamiliesEi(arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "applicationFontFamilies", args)
  }

}

// supportsThreadedFontRendering()
func (this *QFontDatabase) supportsThreadedFontRendering_s(args ...interface{}) () {
  // supportsThreadedFontRendering()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase29supportsThreadedFontRenderingEv
    // invoke: bool supportsThreadedFontRendering()
    C.C_ZN13QFontDatabase29supportsThreadedFontRenderingEv()
  default:
    qtrt.ErrorResolve("QFontDatabase", "supportsThreadedFontRendering", args)
  }

}

// smoothSizes(const class QString &, const class QString &)
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
    // invoke: QList<int> smoothSizes(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QFontDatabase11smoothSizesERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "smoothSizes", args)
  }

}

// removeApplicationFont(int)
func (this *QFontDatabase) removeApplicationFont_s(args ...interface{}) () {
  // removeApplicationFont(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase21removeApplicationFontEi
    // invoke: bool removeApplicationFont(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QFontDatabase21removeApplicationFontEi(arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "removeApplicationFont", args)
  }

}

// hasFamily(const class QString &)
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
    // invoke: bool hasFamily(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QFontDatabase9hasFamilyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "hasFamily", args)
  }

}

// standardSizes()
func (this *QFontDatabase) standardSizes_s(args ...interface{}) () {
  // standardSizes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase13standardSizesEv
    // invoke: QList<int> standardSizes()
    C.C_ZN13QFontDatabase13standardSizesEv()
  default:
    qtrt.ErrorResolve("QFontDatabase", "standardSizes", args)
  }

}

// removeAllApplicationFonts()
func (this *QFontDatabase) removeAllApplicationFonts_s(args ...interface{}) () {
  // removeAllApplicationFonts()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontDatabase25removeAllApplicationFontsEv
    // invoke: bool removeAllApplicationFonts()
    C.C_ZN13QFontDatabase25removeAllApplicationFontsEv()
  default:
    qtrt.ErrorResolve("QFontDatabase", "removeAllApplicationFonts", args)
  }

}

// <= body block end

