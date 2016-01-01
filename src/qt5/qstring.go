package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qstring.h
// dst-file: /src/core/qstring.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QStringDataPtr)=8
type QStringDataPtr struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QString)=8
type QString struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QLatin1String)=16
type QLatin1String struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QCharRef)=16
type QCharRef struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QStringRef)=16
type QStringRef struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QString) toLongLong(args ...interface{}) () {
  // toLongLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString10toLongLongEPbi
  default:
    qtrt.ErrorResolve("QString", "toLongLong", args)
  }

}


func (this *QString) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6isNullEv
  default:
    qtrt.ErrorResolve("QString", "isNull", args)
  }

}


func (this *QString) append(args ...interface{}) () {
  // append(const class QChar *, int)
  // append(const class QByteArray &)
  // append(const class QStringRef &)
  // append(class QChar)
  // append(const class QString &)
  // append(const char *)
  // append(class QLatin1String)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(true) // "const char *"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6appendEPK5QChari
  case 1:
    // invoke: _ZN7QString6appendERK10QByteArray
  case 2:
    // invoke: _ZN7QString6appendERK10QStringRef
  case 3:
    // invoke: _ZN7QString6appendE5QChar
  case 4:
    // invoke: _ZN7QString6appendERKS_
  case 5:
    // invoke: _ZN7QString6appendEPKc
  case 6:
    // invoke: _ZN7QString6appendE13QLatin1String
  default:
    qtrt.ErrorResolve("QString", "append", args)
  }

}


func (this *QString) prepend(args ...interface{}) () {
  // prepend(class QLatin1String)
  // prepend(class QChar)
  // prepend(const char *)
  // prepend(const class QString &)
  // prepend(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString7prependE13QLatin1String
  case 1:
    // invoke: _ZN7QString7prependE5QChar
  case 2:
    // invoke: _ZN7QString7prependEPKc
  case 3:
    // invoke: _ZN7QString7prependERKS_
  case 4:
    // invoke: _ZN7QString7prependERK10QByteArray
  default:
    qtrt.ErrorResolve("QString", "prepend", args)
  }

}


func (this *QString) insert(args ...interface{}) () {
  // insert(int, class QChar)
  // insert(int, const class QChar *, int)
  // insert(int, const class QString &)
  // insert(int, class QLatin1String)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6insertEi5QChar
  case 1:
    // invoke: _ZN7QString6insertEiPK5QChari
  case 2:
    // invoke: _ZN7QString6insertEiRKS_
  case 3:
    // invoke: _ZN7QString6insertEi13QLatin1String
  default:
    qtrt.ErrorResolve("QString", "insert", args)
  }

}


func (this *QString) left(args ...interface{}) () {
  // left(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString4leftEi
  default:
    qtrt.ErrorResolve("QString", "left", args)
  }

}


func NewQString(args ...interface{}) QString {
  return QString{}
}


func (this *QString) lastIndexOf(args ...interface{}) () {
  // lastIndexOf(const class QRegularExpression &, int)
  // lastIndexOf(class QLatin1String, int, Qt::CaseSensitivity)
  // lastIndexOf(class QRegExp &, int)
  // lastIndexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
  // lastIndexOf(const class QRegExp &, int)
  // lastIndexOf(const class QStringRef &, int, Qt::CaseSensitivity)
  // lastIndexOf(const class QString &, int, Qt::CaseSensitivity)
  // lastIndexOf(class QChar, int, Qt::CaseSensitivity)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[6][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString11lastIndexOfERK18QRegularExpressioni
  case 1:
    // invoke: _ZNK7QString11lastIndexOfE13QLatin1StringiN2Qt15CaseSensitivityE
  case 2:
    // invoke: _ZNK7QString11lastIndexOfER7QRegExpi
  case 3:
    // invoke: _ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch
  case 4:
    // invoke: _ZNK7QString11lastIndexOfERK7QRegExpi
  case 5:
    // invoke: _ZNK7QString11lastIndexOfERK10QStringRefiN2Qt15CaseSensitivityE
  case 6:
    // invoke: _ZNK7QString11lastIndexOfERKS_iN2Qt15CaseSensitivityE
  case 7:
    // invoke: _ZNK7QString11lastIndexOfE5QChariN2Qt15CaseSensitivityE
  default:
    qtrt.ErrorResolve("QString", "lastIndexOf", args)
  }

}


func (this *QString) number_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "number", args)
  }

}


func (this *QString) resize(args ...interface{}) () {
  // resize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6resizeEi
  default:
    qtrt.ErrorResolve("QString", "resize", args)
  }

}


func (this *QString) push_front(args ...interface{}) () {
  // push_front(class QChar)
  // push_front(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString10push_frontE5QChar
  case 1:
    // invoke: _ZN7QString10push_frontERKS_
  default:
    qtrt.ErrorResolve("QString", "push_front", args)
  }

}


func (this *QString) toDouble(args ...interface{}) () {
  // toDouble(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8toDoubleEPb
  default:
    qtrt.ErrorResolve("QString", "toDouble", args)
  }

}


func (this *QString) arg(args ...interface{}) () {
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(long, int, int, class QChar)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(const class QString &, const class QString &, const class QString &)
  // arg(double, int, char, int, class QChar)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(const class QString &, const class QString &)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(qlonglong, int, int, class QChar)
  // arg(uint, int, int, class QChar)
  // arg(ushort, int, int, class QChar)
  // arg(short, int, int, class QChar)
  // arg(int, int, int, class QChar)
  // arg(class QChar, int, class QChar)
  // arg(char, int, class QChar)
  // arg(const class QString &, int, class QChar)
  // arg(qulonglong, int, int, class QChar)
  // arg(ulong, int, int, class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][6] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][7] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "long"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.DoubleTy(false) // "double"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.ByteTy(false) // "char"
  vtys[5][3] = qtrt.Int32Ty(false) // "int"
  vtys[5][4] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][6] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][6] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][7] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][8] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[10][1] = qtrt.Int32Ty(false) // "int"
  vtys[10][2] = qtrt.Int32Ty(false) // "int"
  vtys[10][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int32Ty(false) // "uint"
  vtys[11][1] = qtrt.Int32Ty(false) // "int"
  vtys[11][2] = qtrt.Int32Ty(false) // "int"
  vtys[11][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[12][1] = qtrt.Int32Ty(false) // "int"
  vtys[12][2] = qtrt.Int32Ty(false) // "int"
  vtys[12][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.Int16Ty(false) // "short"
  vtys[13][1] = qtrt.Int32Ty(false) // "int"
  vtys[13][2] = qtrt.Int32Ty(false) // "int"
  vtys[13][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.Int32Ty(false) // "int"
  vtys[14][1] = qtrt.Int32Ty(false) // "int"
  vtys[14][2] = qtrt.Int32Ty(false) // "int"
  vtys[14][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[15][1] = qtrt.Int32Ty(false) // "int"
  vtys[15][2] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.ByteTy(false) // "char"
  vtys[16][1] = qtrt.Int32Ty(false) // "int"
  vtys[16][2] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[17][1] = qtrt.Int32Ty(false) // "int"
  vtys[17][2] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[18] = make(map[int32]reflect.Type)
  vtys[18][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[18][1] = qtrt.Int32Ty(false) // "int"
  vtys[18][2] = qtrt.Int32Ty(false) // "int"
  vtys[18][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[19] = make(map[int32]reflect.Type)
  vtys[19][0] = qtrt.Int32Ty(false) // "ulong"
  vtys[19][1] = qtrt.Int32Ty(false) // "int"
  vtys[19][2] = qtrt.Int32Ty(false) // "int"
  vtys[19][3] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_
  case 1:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_
  case 2:
    // invoke: _ZNK7QString3argElii5QChar
  case 3:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_
  case 4:
    // invoke: _ZNK7QString3argERKS_S1_S1_
  case 5:
    // invoke: _ZNK7QString3argEdici5QChar
  case 6:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_
  case 7:
    // invoke: _ZNK7QString3argERKS_S1_
  case 8:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_
  case 9:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_
  case 10:
    // invoke: _ZNK7QString3argExii5QChar
  case 11:
    // invoke: _ZNK7QString3argEjii5QChar
  case 12:
    // invoke: _ZNK7QString3argEtii5QChar
  case 13:
    // invoke: _ZNK7QString3argEsii5QChar
  case 14:
    // invoke: _ZNK7QString3argEiii5QChar
  case 15:
    // invoke: _ZNK7QString3argE5QChariS0_
  case 16:
    // invoke: _ZNK7QString3argEci5QChar
  case 17:
    // invoke: _ZNK7QString3argERKS_i5QChar
  case 18:
    // invoke: _ZNK7QString3argEyii5QChar
  case 19:
    // invoke: _ZNK7QString3argEmii5QChar
  default:
    qtrt.ErrorResolve("QString", "arg", args)
  }

}


func (this *QString) rightRef(args ...interface{}) () {
  // rightRef(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8rightRefEi
  default:
    qtrt.ErrorResolve("QString", "rightRef", args)
  }

}


func (this *QString) setNum(args ...interface{}) () {
  // setNum(short, int)
  // setNum(qulonglong, int)
  // setNum(float, char, int)
  // setNum(qlonglong, int)
  // setNum(uint, int)
  // setNum(int, int)
  // setNum(long, int)
  // setNum(ulong, int)
  // setNum(ushort, int)
  // setNum(double, char, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(false) // "short"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(false) // "float"
  vtys[2][1] = qtrt.ByteTy(false) // "char"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "uint"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "long"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "ulong"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[8][1] = qtrt.Int32Ty(false) // "int"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.DoubleTy(false) // "double"
  vtys[9][1] = qtrt.ByteTy(false) // "char"
  vtys[9][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6setNumEsi
  case 1:
    // invoke: _ZN7QString6setNumEyi
  case 2:
    // invoke: _ZN7QString6setNumEfci
  case 3:
    // invoke: _ZN7QString6setNumExi
  case 4:
    // invoke: _ZN7QString6setNumEji
  case 5:
    // invoke: _ZN7QString6setNumEii
  case 6:
    // invoke: _ZN7QString6setNumEli
  case 7:
    // invoke: _ZN7QString6setNumEmi
  case 8:
    // invoke: _ZN7QString6setNumEti
  case 9:
    // invoke: _ZN7QString6setNumEdci
  default:
    qtrt.ErrorResolve("QString", "setNum", args)
  }

}


func (this *QString) toFloat(args ...interface{}) () {
  // toFloat(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7toFloatEPb
  default:
    qtrt.ErrorResolve("QString", "toFloat", args)
  }

}


func (this *QString) count(args ...interface{}) () {
  // count(const class QStringRef &, Qt::CaseSensitivity)
  // count(const class QRegularExpression &)
  // count()
  // count(const class QRegExp &)
  // count(const class QString &, Qt::CaseSensitivity)
  // count(class QChar, Qt::CaseSensitivity)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[5][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString5countERK10QStringRefN2Qt15CaseSensitivityE
  case 1:
    // invoke: _ZNK7QString5countERK18QRegularExpression
  case 2:
    // invoke: _ZNK7QString5countEv
  case 3:
    // invoke: _ZNK7QString5countERK7QRegExp
  case 4:
    // invoke: _ZNK7QString5countERKS_N2Qt15CaseSensitivityE
  case 5:
    // invoke: _ZNK7QString5countE5QCharN2Qt15CaseSensitivityE
  default:
    qtrt.ErrorResolve("QString", "count", args)
  }

}


func (this *QString) midRef(args ...interface{}) () {
  // midRef(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6midRefEii
  default:
    qtrt.ErrorResolve("QString", "midRef", args)
  }

}


func (this *QString) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6detachEv
  default:
    qtrt.ErrorResolve("QString", "detach", args)
  }

}


func (this *QString) fromStdWString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromStdWString", args)
  }

}


func (this *QString) push_back(args ...interface{}) () {
  // push_back(class QChar)
  // push_back(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString9push_backE5QChar
  case 1:
    // invoke: _ZN7QString9push_backERKS_
  default:
    qtrt.ErrorResolve("QString", "push_back", args)
  }

}


func (this *QString) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString4sizeEv
  default:
    qtrt.ErrorResolve("QString", "size", args)
  }

}


func (this *QString) replace(args ...interface{}) () {
  // replace(const class QString &, const class QString &, Qt::CaseSensitivity)
  // replace(int, int, const class QChar *, int)
  // replace(const class QString &, class QLatin1String, Qt::CaseSensitivity)
  // replace(class QChar, class QLatin1String, Qt::CaseSensitivity)
  // replace(class QLatin1String, const class QString &, Qt::CaseSensitivity)
  // replace(const class QRegExp &, const class QString &)
  // replace(class QChar, const class QString &, Qt::CaseSensitivity)
  // replace(int, int, const class QString &)
  // replace(const class QChar *, int, const class QChar *, int, Qt::CaseSensitivity)
  // replace(class QLatin1String, class QLatin1String, Qt::CaseSensitivity)
  // replace(class QChar, class QChar, Qt::CaseSensitivity)
  // replace(const class QRegularExpression &, const class QString &)
  // replace(int, int, class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[2][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[3][1] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[3][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[5][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[6][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[8][1] = qtrt.Int32Ty(false) // "int"
  vtys[8][2] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[8][3] = qtrt.Int32Ty(false) // "int"
  vtys[8][4] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[9][1] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[9][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[10][1] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[10][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[11][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.Int32Ty(false) // "int"
  vtys[12][1] = qtrt.Int32Ty(false) // "int"
  vtys[12][2] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString7replaceERKS_S1_N2Qt15CaseSensitivityE
  case 1:
    // invoke: _ZN7QString7replaceEiiPK5QChari
  case 2:
    // invoke: _ZN7QString7replaceERKS_13QLatin1StringN2Qt15CaseSensitivityE
  case 3:
    // invoke: _ZN7QString7replaceE5QChar13QLatin1StringN2Qt15CaseSensitivityE
  case 4:
    // invoke: _ZN7QString7replaceE13QLatin1StringRKS_N2Qt15CaseSensitivityE
  case 5:
    // invoke: _ZN7QString7replaceERK7QRegExpRKS_
  case 6:
    // invoke: _ZN7QString7replaceE5QCharRKS_N2Qt15CaseSensitivityE
  case 7:
    // invoke: _ZN7QString7replaceEiiRKS_
  case 8:
    // invoke: _ZN7QString7replaceEPK5QChariS2_iN2Qt15CaseSensitivityE
  case 9:
    // invoke: _ZN7QString7replaceE13QLatin1StringS0_N2Qt15CaseSensitivityE
  case 10:
    // invoke: _ZN7QString7replaceE5QCharS0_N2Qt15CaseSensitivityE
  case 11:
    // invoke: _ZN7QString7replaceERK18QRegularExpressionRKS_
  case 12:
    // invoke: _ZN7QString7replaceEii5QChar
  default:
    qtrt.ErrorResolve("QString", "replace", args)
  }

}


func (this *QString) fromRawData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromRawData", args)
  }

}


func (this *QString) setRawData(args ...interface{}) () {
  // setRawData(const class QChar *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString10setRawDataEPK5QChari
  default:
    qtrt.ErrorResolve("QString", "setRawData", args)
  }

}


func (this *QString) toULong(args ...interface{}) () {
  // toULong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7toULongEPbi
  default:
    qtrt.ErrorResolve("QString", "toULong", args)
  }

}


func (this *QString) chop(args ...interface{}) () {
  // chop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString4chopEi
  default:
    qtrt.ErrorResolve("QString", "chop", args)
  }

}


func (this *QString) fromUtf16_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromUtf16", args)
  }

}


func (this *QString) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString10isDetachedEv
  default:
    qtrt.ErrorResolve("QString", "isDetached", args)
  }

}


func (this *QString) mid(args ...interface{}) () {
  // mid(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString3midEii
  default:
    qtrt.ErrorResolve("QString", "mid", args)
  }

}


func (this *QString) fromLocal8Bit_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromLocal8Bit", args)
  }

}


func (this *QString) swap(args ...interface{}) () {
  // swap(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString4swapERS_
  default:
    qtrt.ErrorResolve("QString", "swap", args)
  }

}


func (this *QString) fromUtf8_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromUtf8", args)
  }

}


func (this *QString) fromUcs4_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromUcs4", args)
  }

}


func (this *QString) leftJustified(args ...interface{}) () {
  // leftJustified(int, class QChar, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString13leftJustifiedEi5QCharb
  default:
    qtrt.ErrorResolve("QString", "leftJustified", args)
  }

}


func (this *QString) indexOf(args ...interface{}) () {
  // indexOf(const class QString &, int, Qt::CaseSensitivity)
  // indexOf(const class QRegExp &, int)
  // indexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
  // indexOf(const class QStringRef &, int, Qt::CaseSensitivity)
  // indexOf(const class QRegularExpression &, int)
  // indexOf(class QChar, int, Qt::CaseSensitivity)
  // indexOf(class QLatin1String, int, Qt::CaseSensitivity)
  // indexOf(class QRegExp &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[6][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7indexOfERKS_iN2Qt15CaseSensitivityE
  case 1:
    // invoke: _ZNK7QString7indexOfERK7QRegExpi
  case 2:
    // invoke: _ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch
  case 3:
    // invoke: _ZNK7QString7indexOfERK10QStringRefiN2Qt15CaseSensitivityE
  case 4:
    // invoke: _ZNK7QString7indexOfERK18QRegularExpressioni
  case 5:
    // invoke: _ZNK7QString7indexOfE5QChariN2Qt15CaseSensitivityE
  case 6:
    // invoke: _ZNK7QString7indexOfE13QLatin1StringiN2Qt15CaseSensitivityE
  case 7:
    // invoke: _ZNK7QString7indexOfER7QRegExpi
  default:
    qtrt.ErrorResolve("QString", "indexOf", args)
  }

}


func (this *QString) utf16(args ...interface{}) () {
  // utf16()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString5utf16Ev
  default:
    qtrt.ErrorResolve("QString", "utf16", args)
  }

}


func (this *QString) toInt(args ...interface{}) () {
  // toInt(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString5toIntEPbi
  default:
    qtrt.ErrorResolve("QString", "toInt", args)
  }

}


func (this *QString) data(args ...interface{}) () {
  // data()
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString4dataEv
  case 1:
    // invoke: _ZNK7QString4dataEv
  default:
    qtrt.ErrorResolve("QString", "data", args)
  }

}


func (this *QString) localeAwareCompare_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "localeAwareCompare", args)
  }

}


func (this *QString) repeated(args ...interface{}) () {
  // repeated(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8repeatedEi
  default:
    qtrt.ErrorResolve("QString", "repeated", args)
  }

}


func (this *QString) setUtf16(args ...interface{}) () {
  // setUtf16(const ushort *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(true) // "const ushort *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString8setUtf16EPKti
  default:
    qtrt.ErrorResolve("QString", "setUtf16", args)
  }

}


func (this *QString) fromStdU32String_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromStdU32String", args)
  }

}


func (this *QString) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString5clearEv
  default:
    qtrt.ErrorResolve("QString", "clear", args)
  }

}


func (this *QString) contains(args ...interface{}) () {
  // contains(class QLatin1String, Qt::CaseSensitivity)
  // contains(const class QRegExp &)
  // contains(const class QStringRef &, Qt::CaseSensitivity)
  // contains(const class QRegularExpression &, class QRegularExpressionMatch *)
  // contains(class QRegExp &)
  // contains(class QChar, Qt::CaseSensitivity)
  // contains(const class QRegularExpression &)
  // contains(const class QString &, Qt::CaseSensitivity)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[2][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[3][1] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[5][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8containsE13QLatin1StringN2Qt15CaseSensitivityE
  case 1:
    // invoke: _ZNK7QString8containsERK7QRegExp
  case 2:
    // invoke: _ZNK7QString8containsERK10QStringRefN2Qt15CaseSensitivityE
  case 3:
    // invoke: _ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch
  case 4:
    // invoke: _ZNK7QString8containsER7QRegExp
  case 5:
    // invoke: _ZNK7QString8containsE5QCharN2Qt15CaseSensitivityE
  case 6:
    // invoke: _ZNK7QString8containsERK18QRegularExpression
  case 7:
    // invoke: _ZNK7QString8containsERKS_N2Qt15CaseSensitivityE
  default:
    qtrt.ErrorResolve("QString", "contains", args)
  }

}


func (this *QString) isSharedWith(args ...interface{}) () {
  // isSharedWith(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString12isSharedWithERKS_
  default:
    qtrt.ErrorResolve("QString", "isSharedWith", args)
  }

}


func (this *QString) fromLatin1_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromLatin1", args)
  }

}


func (this *QString) FreeQString(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "~QString", args)
  }

}


func (this *QString) remove(args ...interface{}) () {
  // remove(class QChar, Qt::CaseSensitivity)
  // remove(const class QRegularExpression &)
  // remove(const class QRegExp &)
  // remove(int, int)
  // remove(const class QString &, Qt::CaseSensitivity)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6removeE5QCharN2Qt15CaseSensitivityE
  case 1:
    // invoke: _ZN7QString6removeERK18QRegularExpression
  case 2:
    // invoke: _ZN7QString6removeERK7QRegExp
  case 3:
    // invoke: _ZN7QString6removeEii
  case 4:
    // invoke: _ZN7QString6removeERKS_N2Qt15CaseSensitivityE
  default:
    qtrt.ErrorResolve("QString", "remove", args)
  }

}


func (this *QString) cend(args ...interface{}) () {
  // cend()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString4cendEv
  default:
    qtrt.ErrorResolve("QString", "cend", args)
  }

}


func (this *QString) toHtmlEscaped(args ...interface{}) () {
  // toHtmlEscaped()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString13toHtmlEscapedEv
  default:
    qtrt.ErrorResolve("QString", "toHtmlEscaped", args)
  }

}


func (this *QString) toWCharArray(args ...interface{}) () {
  // toWCharArray(wchar_t *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.StringTy(false) // "wchar_t *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString12toWCharArrayEPw
  default:
    qtrt.ErrorResolve("QString", "toWCharArray", args)
  }

}


func (this *QString) cbegin(args ...interface{}) () {
  // cbegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6cbeginEv
  default:
    qtrt.ErrorResolve("QString", "cbegin", args)
  }

}


func (this *QString) fromStdString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromStdString", args)
  }

}


func (this *QString) fromWCharArray_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromWCharArray", args)
  }

}


func (this *QString) fill(args ...interface{}) () {
  // fill(class QChar, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString4fillE5QChari
  default:
    qtrt.ErrorResolve("QString", "fill", args)
  }

}


func (this *QString) constData(args ...interface{}) () {
  // constData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString9constDataEv
  default:
    qtrt.ErrorResolve("QString", "constData", args)
  }

}


func (this *QString) toLong(args ...interface{}) () {
  // toLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6toLongEPbi
  default:
    qtrt.ErrorResolve("QString", "toLong", args)
  }

}


func (this *QString) constEnd(args ...interface{}) () {
  // constEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8constEndEv
  default:
    qtrt.ErrorResolve("QString", "constEnd", args)
  }

}


func (this *QString) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6lengthEv
  default:
    qtrt.ErrorResolve("QString", "length", args)
  }

}


func (this *QString) leftRef(args ...interface{}) () {
  // leftRef(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7leftRefEi
  default:
    qtrt.ErrorResolve("QString", "leftRef", args)
  }

}


func (this *QString) isSimpleText(args ...interface{}) () {
  // isSimpleText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString12isSimpleTextEv
  default:
    qtrt.ErrorResolve("QString", "isSimpleText", args)
  }

}


func (this *QString) setUnicode(args ...interface{}) () {
  // setUnicode(const class QChar *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString10setUnicodeEPK5QChari
  default:
    qtrt.ErrorResolve("QString", "setUnicode", args)
  }

}


func (this *QString) constBegin(args ...interface{}) () {
  // constBegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString10constBeginEv
  default:
    qtrt.ErrorResolve("QString", "constBegin", args)
  }

}


func (this *QString) unicode(args ...interface{}) () {
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7unicodeEv
  default:
    qtrt.ErrorResolve("QString", "unicode", args)
  }

}


func (this *QString) at(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString2atEi
  default:
    qtrt.ErrorResolve("QString", "at", args)
  }

}


func (this *QString) begin(args ...interface{}) () {
  // begin()
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString5beginEv
  case 1:
    // invoke: _ZN7QString5beginEv
  default:
    qtrt.ErrorResolve("QString", "begin", args)
  }

}


func (this *QString) end(args ...interface{}) () {
  // end()
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString3endEv
  case 1:
    // invoke: _ZN7QString3endEv
  default:
    qtrt.ErrorResolve("QString", "end", args)
  }

}


func (this *QString) toUInt(args ...interface{}) () {
  // toUInt(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6toUIntEPbi
  default:
    qtrt.ErrorResolve("QString", "toUInt", args)
  }

}


func (this *QString) fromStdU16String_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromStdU16String", args)
  }

}


func (this *QString) toUShort(args ...interface{}) () {
  // toUShort(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8toUShortEPbi
  default:
    qtrt.ErrorResolve("QString", "toUShort", args)
  }

}


func (this *QString) toULongLong(args ...interface{}) () {
  // toULongLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString11toULongLongEPbi
  default:
    qtrt.ErrorResolve("QString", "toULongLong", args)
  }

}


func (this *QString) capacity(args ...interface{}) () {
  // capacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8capacityEv
  default:
    qtrt.ErrorResolve("QString", "capacity", args)
  }

}


func (this *QString) squeeze(args ...interface{}) () {
  // squeeze()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString7squeezeEv
  default:
    qtrt.ErrorResolve("QString", "squeeze", args)
  }

}


func (this *QString) truncate(args ...interface{}) () {
  // truncate(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString8truncateEi
  default:
    qtrt.ErrorResolve("QString", "truncate", args)
  }

}


func (this *QString) localeAwareCompare(args ...interface{}) () {
  // localeAwareCompare(const class QString &, const class QString &)
  // localeAwareCompare(const class QString &, const class QStringRef &)
  // localeAwareCompare(const class QStringRef &)
  // localeAwareCompare(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString18localeAwareCompareERKS_S1_
  case 1:
    // invoke: _ZN7QString18localeAwareCompareERKS_RK10QStringRef
  case 2:
    // invoke: _ZNK7QString18localeAwareCompareERK10QStringRef
  case 3:
    // invoke: _ZNK7QString18localeAwareCompareERKS_
  default:
    qtrt.ErrorResolve("QString", "localeAwareCompare", args)
  }

}


func (this *QString) isRightToLeft(args ...interface{}) () {
  // isRightToLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString13isRightToLeftEv
  default:
    qtrt.ErrorResolve("QString", "isRightToLeft", args)
  }

}


func (this *QString) toUcs4(args ...interface{}) () {
  // toUcs4()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6toUcs4Ev
  default:
    qtrt.ErrorResolve("QString", "toUcs4", args)
  }

}


func (this *QString) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7isEmptyEv
  default:
    qtrt.ErrorResolve("QString", "isEmpty", args)
  }

}


func (this *QString) right(args ...interface{}) () {
  // right(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString5rightEi
  default:
    qtrt.ErrorResolve("QString", "right", args)
  }

}


func (this *QString) rightJustified(args ...interface{}) () {
  // rightJustified(int, class QChar, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString14rightJustifiedEi5QCharb
  default:
    qtrt.ErrorResolve("QString", "rightJustified", args)
  }

}


func (this *QString) reserve(args ...interface{}) () {
  // reserve(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString7reserveEi
  default:
    qtrt.ErrorResolve("QString", "reserve", args)
  }

}


func (this *QString) toShort(args ...interface{}) () {
  // toShort(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7toShortEPbi
  default:
    qtrt.ErrorResolve("QString", "toShort", args)
  }

}


func (this *QLatin1String) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QLatin1String4dataEv
  default:
    qtrt.ErrorResolve("QLatin1String", "data", args)
  }

}


func NewQLatin1String(args ...interface{}) QLatin1String {
  return QLatin1String{}
}


func (this *QLatin1String) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QLatin1String4sizeEv
  default:
    qtrt.ErrorResolve("QLatin1String", "size", args)
  }

}


func (this *QLatin1String) latin1(args ...interface{}) () {
  // latin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QLatin1String6latin1Ev
  default:
    qtrt.ErrorResolve("QLatin1String", "latin1", args)
  }

}


func (this *QCharRef) isLetterOrNumber(args ...interface{}) () {
  // isLetterOrNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QCharRef16isLetterOrNumberEv
  default:
    qtrt.ErrorResolve("QCharRef", "isLetterOrNumber", args)
  }

}


func (this *QCharRef) isDigit(args ...interface{}) () {
  // isDigit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isDigitEv
  default:
    qtrt.ErrorResolve("QCharRef", "isDigit", args)
  }

}


func (this *QCharRef) toLatin1(args ...interface{}) () {
  // toLatin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef8toLatin1Ev
  default:
    qtrt.ErrorResolve("QCharRef", "toLatin1", args)
  }

}


func (this *QCharRef) setCell(args ...interface{}) () {
  // setCell(uchar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "uchar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QCharRef7setCellEh
  default:
    qtrt.ErrorResolve("QCharRef", "setCell", args)
  }

}


func (this *QCharRef) isMark(args ...interface{}) () {
  // isMark()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef6isMarkEv
  default:
    qtrt.ErrorResolve("QCharRef", "isMark", args)
  }

}


func NewQCharRef(args ...interface{}) QCharRef {
  return QCharRef{}
}


func (this *QCharRef) digitValue(args ...interface{}) () {
  // digitValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef10digitValueEv
  default:
    qtrt.ErrorResolve("QCharRef", "digitValue", args)
  }

}


func (this *QCharRef) isLetter(args ...interface{}) () {
  // isLetter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef8isLetterEv
  default:
    qtrt.ErrorResolve("QCharRef", "isLetter", args)
  }

}


func (this *QCharRef) isNumber(args ...interface{}) () {
  // isNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef8isNumberEv
  default:
    qtrt.ErrorResolve("QCharRef", "isNumber", args)
  }

}


func (this *QCharRef) isPrint(args ...interface{}) () {
  // isPrint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isPrintEv
  default:
    qtrt.ErrorResolve("QCharRef", "isPrint", args)
  }

}


func (this *QCharRef) toLower(args ...interface{}) () {
  // toLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7toLowerEv
  default:
    qtrt.ErrorResolve("QCharRef", "toLower", args)
  }

}


func (this *QCharRef) setRow(args ...interface{}) () {
  // setRow(uchar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "uchar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QCharRef6setRowEh
  default:
    qtrt.ErrorResolve("QCharRef", "setRow", args)
  }

}


func (this *QCharRef) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef6isNullEv
  default:
    qtrt.ErrorResolve("QCharRef", "isNull", args)
  }

}


func (this *QCharRef) toTitleCase(args ...interface{}) () {
  // toTitleCase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef11toTitleCaseEv
  default:
    qtrt.ErrorResolve("QCharRef", "toTitleCase", args)
  }

}


func (this *QCharRef) hasMirrored(args ...interface{}) () {
  // hasMirrored()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef11hasMirroredEv
  default:
    qtrt.ErrorResolve("QCharRef", "hasMirrored", args)
  }

}


func (this *QCharRef) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef3rowEv
  default:
    qtrt.ErrorResolve("QCharRef", "row", args)
  }

}


func (this *QCharRef) unicode(args ...interface{}) () {
  // unicode()
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QCharRef7unicodeEv
  case 1:
    // invoke: _ZNK8QCharRef7unicodeEv
  default:
    qtrt.ErrorResolve("QCharRef", "unicode", args)
  }

}


func (this *QCharRef) isTitleCase(args ...interface{}) () {
  // isTitleCase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef11isTitleCaseEv
  default:
    qtrt.ErrorResolve("QCharRef", "isTitleCase", args)
  }

}


func (this *QCharRef) isUpper(args ...interface{}) () {
  // isUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isUpperEv
  default:
    qtrt.ErrorResolve("QCharRef", "isUpper", args)
  }

}


func (this *QCharRef) cell(args ...interface{}) () {
  // cell()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef4cellEv
  default:
    qtrt.ErrorResolve("QCharRef", "cell", args)
  }

}


func (this *QCharRef) decomposition(args ...interface{}) () {
  // decomposition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef13decompositionEv
  default:
    qtrt.ErrorResolve("QCharRef", "decomposition", args)
  }

}


func (this *QCharRef) combiningClass(args ...interface{}) () {
  // combiningClass()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef14combiningClassEv
  default:
    qtrt.ErrorResolve("QCharRef", "combiningClass", args)
  }

}


func (this *QCharRef) mirroredChar(args ...interface{}) () {
  // mirroredChar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef12mirroredCharEv
  default:
    qtrt.ErrorResolve("QCharRef", "mirroredChar", args)
  }

}


func (this *QCharRef) isSpace(args ...interface{}) () {
  // isSpace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isSpaceEv
  default:
    qtrt.ErrorResolve("QCharRef", "isSpace", args)
  }

}


func (this *QCharRef) isPunct(args ...interface{}) () {
  // isPunct()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isPunctEv
  default:
    qtrt.ErrorResolve("QCharRef", "isPunct", args)
  }

}


func (this *QCharRef) toUpper(args ...interface{}) () {
  // toUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7toUpperEv
  default:
    qtrt.ErrorResolve("QCharRef", "toUpper", args)
  }

}


func (this *QCharRef) isLower(args ...interface{}) () {
  // isLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isLowerEv
  default:
    qtrt.ErrorResolve("QCharRef", "isLower", args)
  }

}


func (this *QStringRef) toShort(args ...interface{}) () {
  // toShort(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef7toShortEPbi
  default:
    qtrt.ErrorResolve("QStringRef", "toShort", args)
  }

}


func NewQStringRef(args ...interface{}) QStringRef {
  return QStringRef{}
}


func (this *QStringRef) toULongLong(args ...interface{}) () {
  // toULongLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef11toULongLongEPbi
  default:
    qtrt.ErrorResolve("QStringRef", "toULongLong", args)
  }

}


func (this *QStringRef) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStringRef5clearEv
  default:
    qtrt.ErrorResolve("QStringRef", "clear", args)
  }

}


func (this *QStringRef) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8positionEv
  default:
    qtrt.ErrorResolve("QStringRef", "position", args)
  }

}


func (this *QStringRef) toLong(args ...interface{}) () {
  // toLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6toLongEPbi
  default:
    qtrt.ErrorResolve("QStringRef", "toLong", args)
  }

}


func (this *QStringRef) cbegin(args ...interface{}) () {
  // cbegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6cbeginEv
  default:
    qtrt.ErrorResolve("QStringRef", "cbegin", args)
  }

}


func (this *QStringRef) toUShort(args ...interface{}) () {
  // toUShort(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8toUShortEPbi
  default:
    qtrt.ErrorResolve("QStringRef", "toUShort", args)
  }

}


func (this *QStringRef) toUInt(args ...interface{}) () {
  // toUInt(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6toUIntEPbi
  default:
    qtrt.ErrorResolve("QStringRef", "toUInt", args)
  }

}


func (this *QStringRef) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef7isEmptyEv
  default:
    qtrt.ErrorResolve("QStringRef", "isEmpty", args)
  }

}


func (this *QStringRef) localeAwareCompare(args ...interface{}) () {
  // localeAwareCompare(const class QString &)
  // localeAwareCompare(const class QStringRef &)
  // localeAwareCompare(const class QStringRef &, const class QString &)
  // localeAwareCompare(const class QStringRef &, const class QStringRef &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[3][1] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef18localeAwareCompareERK7QString
  case 1:
    // invoke: _ZNK10QStringRef18localeAwareCompareERKS_
  case 2:
    // invoke: _ZN10QStringRef18localeAwareCompareERKS_RK7QString
  case 3:
    // invoke: _ZN10QStringRef18localeAwareCompareERKS_S1_
  default:
    qtrt.ErrorResolve("QStringRef", "localeAwareCompare", args)
  }

}


func (this *QStringRef) toUtf8(args ...interface{}) () {
  // toUtf8()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6toUtf8Ev
  default:
    qtrt.ErrorResolve("QStringRef", "toUtf8", args)
  }

}


func (this *QStringRef) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef4sizeEv
  default:
    qtrt.ErrorResolve("QStringRef", "size", args)
  }

}


func (this *QStringRef) constData(args ...interface{}) () {
  // constData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef9constDataEv
  default:
    qtrt.ErrorResolve("QStringRef", "constData", args)
  }

}


func (this *QStringRef) left(args ...interface{}) () {
  // left(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef4leftEi
  default:
    qtrt.ErrorResolve("QStringRef", "left", args)
  }

}


func (this *QStringRef) toUcs4(args ...interface{}) () {
  // toUcs4()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6toUcs4Ev
  default:
    qtrt.ErrorResolve("QStringRef", "toUcs4", args)
  }

}


func (this *QStringRef) count(args ...interface{}) () {
  // count(const class QStringRef &, Qt::CaseSensitivity)
  // count(const class QString &, Qt::CaseSensitivity)
  // count()
  // count(class QChar, Qt::CaseSensitivity)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[3][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef5countERKS_N2Qt15CaseSensitivityE
  case 1:
    // invoke: _ZNK10QStringRef5countERK7QStringN2Qt15CaseSensitivityE
  case 2:
    // invoke: _ZNK10QStringRef5countEv
  case 3:
    // invoke: _ZNK10QStringRef5countE5QCharN2Qt15CaseSensitivityE
  default:
    qtrt.ErrorResolve("QStringRef", "count", args)
  }

}


func (this *QStringRef) right(args ...interface{}) () {
  // right(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef5rightEi
  default:
    qtrt.ErrorResolve("QStringRef", "right", args)
  }

}


func (this *QStringRef) at(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef2atEi
  default:
    qtrt.ErrorResolve("QStringRef", "at", args)
  }

}


func (this *QStringRef) toDouble(args ...interface{}) () {
  // toDouble(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8toDoubleEPb
  default:
    qtrt.ErrorResolve("QStringRef", "toDouble", args)
  }

}


func (this *QStringRef) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6isNullEv
  default:
    qtrt.ErrorResolve("QStringRef", "isNull", args)
  }

}


func (this *QStringRef) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef4dataEv
  default:
    qtrt.ErrorResolve("QStringRef", "data", args)
  }

}


func (this *QStringRef) toLongLong(args ...interface{}) () {
  // toLongLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef10toLongLongEPbi
  default:
    qtrt.ErrorResolve("QStringRef", "toLongLong", args)
  }

}


func (this *QStringRef) toLatin1(args ...interface{}) () {
  // toLatin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8toLatin1Ev
  default:
    qtrt.ErrorResolve("QStringRef", "toLatin1", args)
  }

}


func (this *QStringRef) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef5beginEv
  default:
    qtrt.ErrorResolve("QStringRef", "begin", args)
  }

}


func (this *QStringRef) unicode(args ...interface{}) () {
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef7unicodeEv
  default:
    qtrt.ErrorResolve("QStringRef", "unicode", args)
  }

}


func (this *QStringRef) mid(args ...interface{}) () {
  // mid(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef3midEii
  default:
    qtrt.ErrorResolve("QStringRef", "mid", args)
  }

}


func (this *QStringRef) toFloat(args ...interface{}) () {
  // toFloat(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef7toFloatEPb
  default:
    qtrt.ErrorResolve("QStringRef", "toFloat", args)
  }

}


func (this *QStringRef) string(args ...interface{}) () {
  // string()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6stringEv
  default:
    qtrt.ErrorResolve("QStringRef", "string", args)
  }

}


func (this *QStringRef) toString(args ...interface{}) () {
  // toString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8toStringEv
  default:
    qtrt.ErrorResolve("QStringRef", "toString", args)
  }

}


func (this *QStringRef) trimmed(args ...interface{}) () {
  // trimmed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef7trimmedEv
  default:
    qtrt.ErrorResolve("QStringRef", "trimmed", args)
  }

}


func (this *QStringRef) toInt(args ...interface{}) () {
  // toInt(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef5toIntEPbi
  default:
    qtrt.ErrorResolve("QStringRef", "toInt", args)
  }

}


func (this *QStringRef) cend(args ...interface{}) () {
  // cend()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef4cendEv
  default:
    qtrt.ErrorResolve("QStringRef", "cend", args)
  }

}


func (this *QStringRef) appendTo(args ...interface{}) () {
  // appendTo(class QString *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8appendToEP7QString
  default:
    qtrt.ErrorResolve("QStringRef", "appendTo", args)
  }

}


func (this *QStringRef) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6lengthEv
  default:
    qtrt.ErrorResolve("QStringRef", "length", args)
  }

}


func (this *QStringRef) FreeQStringRef(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStringRef", "~QStringRef", args)
  }

}


func (this *QStringRef) toLocal8Bit(args ...interface{}) () {
  // toLocal8Bit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef11toLocal8BitEv
  default:
    qtrt.ErrorResolve("QStringRef", "toLocal8Bit", args)
  }

}


func (this *QStringRef) toULong(args ...interface{}) () {
  // toULong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef7toULongEPbi
  default:
    qtrt.ErrorResolve("QStringRef", "toULong", args)
  }

}


func (this *QStringRef) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef3endEv
  default:
    qtrt.ErrorResolve("QStringRef", "end", args)
  }

}

// <= body block end

