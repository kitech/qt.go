package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qurlquery.h
// dst-file: /src/core/qurlquery.go
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
// class sizeof(QUrlQuery)=1
type QUrlQuery struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQUrlQuery(args ...interface{}) QUrlQuery {
  return QUrlQuery{}
}


func (this *QUrlQuery) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQuery5clearEv
  default:
    qtrt.ErrorResolve("QUrlQuery", "clear", args)
  }

}


func (this *QUrlQuery) setQuery(args ...interface{}) () {
  // setQuery(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQuery8setQueryERK7QString
  default:
    qtrt.ErrorResolve("QUrlQuery", "setQuery", args)
  }

}


func (this *QUrlQuery) queryValueDelimiter(args ...interface{}) () {
  // queryValueDelimiter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUrlQuery19queryValueDelimiterEv
  default:
    qtrt.ErrorResolve("QUrlQuery", "queryValueDelimiter", args)
  }

}


func (this *QUrlQuery) queryPairDelimiter(args ...interface{}) () {
  // queryPairDelimiter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUrlQuery18queryPairDelimiterEv
  default:
    qtrt.ErrorResolve("QUrlQuery", "queryPairDelimiter", args)
  }

}


func (this *QUrlQuery) defaultQueryValueDelimiter_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrlQuery", "defaultQueryValueDelimiter", args)
  }

}


func (this *QUrlQuery) swap(args ...interface{}) () {
  // swap(class QUrlQuery &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrlQuery{}) // "QUrlQuery &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQuery4swapERS_
  default:
    qtrt.ErrorResolve("QUrlQuery", "swap", args)
  }

}


func (this *QUrlQuery) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUrlQuery10isDetachedEv
  default:
    qtrt.ErrorResolve("QUrlQuery", "isDetached", args)
  }

}


func (this *QUrlQuery) setQueryDelimiters(args ...interface{}) () {
  // setQueryDelimiters(class QChar, class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQuery18setQueryDelimitersE5QCharS0_
  default:
    qtrt.ErrorResolve("QUrlQuery", "setQueryDelimiters", args)
  }

}


func (this *QUrlQuery) FreeQUrlQuery(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrlQuery", "~QUrlQuery", args)
  }

}


func (this *QUrlQuery) removeAllQueryItems(args ...interface{}) () {
  // removeAllQueryItems(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQuery19removeAllQueryItemsERK7QString
  default:
    qtrt.ErrorResolve("QUrlQuery", "removeAllQueryItems", args)
  }

}


func (this *QUrlQuery) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUrlQuery7isEmptyEv
  default:
    qtrt.ErrorResolve("QUrlQuery", "isEmpty", args)
  }

}


func (this *QUrlQuery) removeQueryItem(args ...interface{}) () {
  // removeQueryItem(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQuery15removeQueryItemERK7QString
  default:
    qtrt.ErrorResolve("QUrlQuery", "removeQueryItem", args)
  }

}


func (this *QUrlQuery) defaultQueryPairDelimiter_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrlQuery", "defaultQueryPairDelimiter", args)
  }

}


func (this *QUrlQuery) addQueryItem(args ...interface{}) () {
  // addQueryItem(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQuery12addQueryItemERK7QStringS2_
  default:
    qtrt.ErrorResolve("QUrlQuery", "addQueryItem", args)
  }

}


func (this *QUrlQuery) hasQueryItem(args ...interface{}) () {
  // hasQueryItem(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUrlQuery12hasQueryItemERK7QString
  default:
    qtrt.ErrorResolve("QUrlQuery", "hasQueryItem", args)
  }

}

// <= body block end

