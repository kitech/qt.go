package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern int32_t C_ZN13QFontDatabase26addApplicationFontFromDataERK10QByteArray(void* arg0); // 4
  // proto:  int QFontDatabase::weight(const QString & family, const QString & style);
extern int32_t C_ZNK13QFontDatabase6weightERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QFontDatabase::isFixedPitch(const QString & family, const QString & style);
extern bool C_ZNK13QFontDatabase12isFixedPitchERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QFont QFontDatabase::font(const QString & family, const QString & style, int pointSize);
extern void* C_ZNK13QFontDatabase4fontERK7QStringS2_i(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  QList<int> QFontDatabase::pointSizes(const QString & family, const QString & style);
extern void C_ZN13QFontDatabase10pointSizesERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QList<QFontDatabase::WritingSystem> QFontDatabase::writingSystems();
extern void C_ZNK13QFontDatabase14writingSystemsEv(void* qthis); // 4
  // proto:  QList<QFontDatabase::WritingSystem> QFontDatabase::writingSystems(const QString & family);
extern void C_ZNK13QFontDatabase14writingSystemsERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QFontDatabase::isPrivateFamily(const QString & family);
extern bool C_ZNK13QFontDatabase15isPrivateFamilyERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFontDatabase::QFontDatabase();
extern void* C_ZN13QFontDatabaseC2Ev(); // 3
  // proto:  QString QFontDatabase::styleString(const QFontInfo & fontInfo);
extern void* C_ZN13QFontDatabase11styleStringERK9QFontInfo(void* qthis, void* arg0); // 4
  // proto:  QString QFontDatabase::styleString(const QFont & font);
extern void* C_ZN13QFontDatabase11styleStringERK5QFont(void* qthis, void* arg0); // 4
  // proto:  bool QFontDatabase::italic(const QString & family, const QString & style);
extern bool C_ZNK13QFontDatabase6italicERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QStringList QFontDatabase::styles(const QString & family);
extern void C_ZNK13QFontDatabase6stylesERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QFontDatabase::isSmoothlyScalable(const QString & family, const QString & style);
extern bool C_ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QFontDatabase::bold(const QString & family, const QString & style);
extern bool C_ZNK13QFontDatabase4boldERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QFontDatabase::isScalable(const QString & family, const QString & style);
extern bool C_ZNK13QFontDatabase10isScalableERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto: static int QFontDatabase::addApplicationFont(const QString & fileName);
extern int32_t C_ZN13QFontDatabase18addApplicationFontERK7QString(void* arg0); // 4
  // proto:  bool QFontDatabase::isBitmapScalable(const QString & family, const QString & style);
extern bool C_ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto: static QStringList QFontDatabase::applicationFontFamilies(int id);
extern void C_ZN13QFontDatabase23applicationFontFamiliesEi(int32_t arg0); // 4
  // proto: static bool QFontDatabase::supportsThreadedFontRendering();
extern bool C_ZN13QFontDatabase29supportsThreadedFontRenderingEv(); // 4
  // proto:  QList<int> QFontDatabase::smoothSizes(const QString & family, const QString & style);
extern void C_ZN13QFontDatabase11smoothSizesERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto: static bool QFontDatabase::removeApplicationFont(int id);
extern bool C_ZN13QFontDatabase21removeApplicationFontEi(int32_t arg0); // 4
  // proto:  bool QFontDatabase::hasFamily(const QString & family);
extern bool C_ZNK13QFontDatabase9hasFamilyERK7QString(void* qthis, void* arg0); // 4
  // proto: static QList<int> QFontDatabase::standardSizes();
extern void C_ZN13QFontDatabase13standardSizesEv(); // 4
  // proto: static bool QFontDatabase::removeAllApplicationFonts();
extern bool C_ZN13QFontDatabase25removeAllApplicationFontsEv(); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// addApplicationFontFromData(const class QByteArray &)
func (this *QFontDatabase) Addapplicationfontfromdata_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QFontDatabase26addApplicationFontFromDataERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "addApplicationFontFromData", args)
  }

  return
}

// weight(const class QString &, const class QString &)
func (this *QFontDatabase) Weight(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QFontDatabase6weightERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "weight", args)
  }

  return
}

// isFixedPitch(const class QString &, const class QString &)
func (this *QFontDatabase) Isfixedpitch(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QFontDatabase12isFixedPitchERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "isFixedPitch", args)
  }

  return
}

// font(const class QString &, const class QString &, int)
func (this *QFontDatabase) Font(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK13QFontDatabase4fontERK7QStringS2_i(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "font", args)
  }

  return
}

// pointSizes(const class QString &, const class QString &)
func (this *QFontDatabase) Pointsizes(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QFontDatabase10pointSizesERK7QStringS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "pointSizes", args)
  }

  return
}

// writingSystems()
func (this *QFontDatabase) Writingsystems(args ...interface{}) () {
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
    C.C_ZNK13QFontDatabase14writingSystemsEv(this.Qclsinst)
  case 1:
    // invoke: _ZNK13QFontDatabase14writingSystemsERK7QString
    // invoke: QList<QFontDatabase::WritingSystem> writingSystems(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QFontDatabase14writingSystemsERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "writingSystems", args)
  }

  return
}

// isPrivateFamily(const class QString &)
func (this *QFontDatabase) Isprivatefamily(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontDatabase15isPrivateFamilyERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "isPrivateFamily", args)
  }

  return
}

// QFontDatabase()
func NewQFontDatabase(args ...interface{}) *QFontDatabase {
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
    qthis = C.C_ZN13QFontDatabaseC2Ev()
    return &QFontDatabase{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFontDatabase", "QFontDatabase", args)
  }

  return nil // QFontDatabase{Qclsinst:qthis}
}

// styleString(const class QFontInfo &)
func (this *QFontDatabase) Stylestring(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QFontInfo).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QFontDatabase11styleStringERK9QFontInfo(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN13QFontDatabase11styleStringERK5QFont
    // invoke: QString styleString(const class QFont &)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QFontDatabase11styleStringERK5QFont(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "styleString", args)
  }

  return
}

// italic(const class QString &, const class QString &)
func (this *QFontDatabase) Italic(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QFontDatabase6italicERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "italic", args)
  }

  return
}

// styles(const class QString &)
func (this *QFontDatabase) Styles(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QFontDatabase6stylesERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "styles", args)
  }

  return
}

// isSmoothlyScalable(const class QString &, const class QString &)
func (this *QFontDatabase) Issmoothlyscalable(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "isSmoothlyScalable", args)
  }

  return
}

// bold(const class QString &, const class QString &)
func (this *QFontDatabase) Bold(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QFontDatabase4boldERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "bold", args)
  }

  return
}

// isScalable(const class QString &, const class QString &)
func (this *QFontDatabase) Isscalable(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QFontDatabase10isScalableERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "isScalable", args)
  }

  return
}

// addApplicationFont(const class QString &)
func (this *QFontDatabase) Addapplicationfont_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QFontDatabase18addApplicationFontERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "addApplicationFont", args)
  }

  return
}

// isBitmapScalable(const class QString &, const class QString &)
func (this *QFontDatabase) Isbitmapscalable(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "isBitmapScalable", args)
  }

  return
}

// applicationFontFamilies(int)
func (this *QFontDatabase) Applicationfontfamilies_S(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QFontDatabase23applicationFontFamiliesEi(arg0)
  default:
    qtrt.ErrorResolve("QFontDatabase", "applicationFontFamilies", args)
  }

  return
}

// supportsThreadedFontRendering()
func (this *QFontDatabase) Supportsthreadedfontrendering_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN13QFontDatabase29supportsThreadedFontRenderingEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "supportsThreadedFontRendering", args)
  }

  return
}

// smoothSizes(const class QString &, const class QString &)
func (this *QFontDatabase) Smoothsizes(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QFontDatabase11smoothSizesERK7QStringS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDatabase", "smoothSizes", args)
  }

  return
}

// removeApplicationFont(int)
func (this *QFontDatabase) Removeapplicationfont_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QFontDatabase21removeApplicationFontEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "removeApplicationFont", args)
  }

  return
}

// hasFamily(const class QString &)
func (this *QFontDatabase) Hasfamily(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontDatabase9hasFamilyERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "hasFamily", args)
  }

  return
}

// standardSizes()
func (this *QFontDatabase) Standardsizes_S(args ...interface{}) () {
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

  return
}

// removeAllApplicationFonts()
func (this *QFontDatabase) Removeallapplicationfonts_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN13QFontDatabase25removeAllApplicationFontsEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDatabase", "removeAllApplicationFonts", args)
  }

  return
}

// <= body block end

