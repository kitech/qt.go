package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qtranslator.h
// dst-file: /src/core/qtranslator.go
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
// class sizeof(QTranslator)=1
type QTranslator struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTranslator) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTranslator10metaObjectEv
  default:
    qtrt.ErrorResolve("QTranslator", "metaObject", args)
  }

}


func NewQTranslator(args ...interface{}) QTranslator {
  return QTranslator{}
}


func (this *QTranslator) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTranslator7isEmptyEv
  default:
    qtrt.ErrorResolve("QTranslator", "isEmpty", args)
  }

}


func (this *QTranslator) FreeQTranslator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTranslator", "~QTranslator", args)
  }

}


func (this *QTranslator) load(args ...interface{}) () {
  // load(const class QString &, const class QString &, const class QString &, const class QString &)
  // load(const uchar *, int, const class QString &)
  // load(const class QLocale &, const class QString &, const class QString &, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const uchar *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][4] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTranslator4loadERK7QStringS2_S2_S2_
  case 1:
    // invoke: _ZN11QTranslator4loadEPKhiRK7QString
  case 2:
    // invoke: _ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_
  default:
    qtrt.ErrorResolve("QTranslator", "load", args)
  }

}


func (this *QTranslator) translate(args ...interface{}) () {
  // translate(const char *, const char *, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTranslator9translateEPKcS1_S1_i
  default:
    qtrt.ErrorResolve("QTranslator", "translate", args)
  }

}

// <= body block end

