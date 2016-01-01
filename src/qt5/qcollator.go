package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qcollator.h
// dst-file: /src/core/qcollator.go
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
// class sizeof(QCollator)=8
type QCollator struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QCollatorSortKey)=1
type QCollatorSortKey struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QCollator) numericMode(args ...interface{}) () {
  // numericMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator11numericModeEv
  default:
    qtrt.ErrorResolve("QCollator", "numericMode", args)
  }

}


func NewQCollator(args ...interface{}) QCollator {
  return QCollator{}
}


func (this *QCollator) setLocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator9setLocaleERK7QLocale
  default:
    qtrt.ErrorResolve("QCollator", "setLocale", args)
  }

}


func (this *QCollator) setNumericMode(args ...interface{}) () {
  // setNumericMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator14setNumericModeEb
  default:
    qtrt.ErrorResolve("QCollator", "setNumericMode", args)
  }

}


func (this *QCollator) compare(args ...interface{}) () {
  // compare(const class QStringRef &, const class QStringRef &)
  // compare(const class QChar *, int, const class QChar *, int)
  // compare(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator7compareERK10QStringRefS2_
  case 1:
    // invoke: _ZNK9QCollator7compareEPK5QChariS2_i
  case 2:
    // invoke: _ZNK9QCollator7compareERK7QStringS2_
  default:
    qtrt.ErrorResolve("QCollator", "compare", args)
  }

}


func (this *QCollator) sortKey(args ...interface{}) () {
  // sortKey(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator7sortKeyERK7QString
  default:
    qtrt.ErrorResolve("QCollator", "sortKey", args)
  }

}


func (this *QCollator) FreeQCollator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCollator", "~QCollator", args)
  }

}


func (this *QCollator) ignorePunctuation(args ...interface{}) () {
  // ignorePunctuation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator17ignorePunctuationEv
  default:
    qtrt.ErrorResolve("QCollator", "ignorePunctuation", args)
  }

}


func (this *QCollator) locale(args ...interface{}) () {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator6localeEv
  default:
    qtrt.ErrorResolve("QCollator", "locale", args)
  }

}


func (this *QCollator) swap(args ...interface{}) () {
  // swap(class QCollator &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCollator{}) // "QCollator &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator4swapERS_
  default:
    qtrt.ErrorResolve("QCollator", "swap", args)
  }

}


func (this *QCollator) setIgnorePunctuation(args ...interface{}) () {
  // setIgnorePunctuation(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator20setIgnorePunctuationEb
  default:
    qtrt.ErrorResolve("QCollator", "setIgnorePunctuation", args)
  }

}


func (this *QCollatorSortKey) FreeQCollatorSortKey(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "~QCollatorSortKey", args)
  }

}


func (this *QCollatorSortKey) swap(args ...interface{}) () {
  // swap(class QCollatorSortKey &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCollatorSortKey{}) // "QCollatorSortKey &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCollatorSortKey4swapERS_
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "swap", args)
  }

}


func (this *QCollatorSortKey) compare(args ...interface{}) () {
  // compare(const class QCollatorSortKey &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCollatorSortKey{}) // "const QCollatorSortKey &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QCollatorSortKey7compareERKS_
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "compare", args)
  }

}


func NewQCollatorSortKey(args ...interface{}) QCollatorSortKey {
  return QCollatorSortKey{}
}

// <= body block end

