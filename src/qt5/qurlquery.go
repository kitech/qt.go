package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qurlquery.h
// dst-file: /src/core/qurlquery.go
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
  // proto:  void QUrlQuery::addQueryItem(const QString & key, const QString & value);
extern void C_ZN9QUrlQuery12addQueryItemERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QChar QUrlQuery::queryValueDelimiter();
extern void* C_ZNK9QUrlQuery19queryValueDelimiterEv(void* qthis); // 4
  // proto:  bool QUrlQuery::hasQueryItem(const QString & key);
extern bool C_ZNK9QUrlQuery12hasQueryItemERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QUrlQuery::removeAllQueryItems(const QString & key);
extern void C_ZN9QUrlQuery19removeAllQueryItemsERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QUrlQuery::~QUrlQuery();
extern void C_ZN9QUrlQueryD2Ev(void* qthis); // 4
  // proto:  bool QUrlQuery::isEmpty();
extern bool C_ZNK9QUrlQuery7isEmptyEv(void* qthis); // 4
  // proto:  void QUrlQuery::swap(QUrlQuery & other);
extern void C_ZN9QUrlQuery4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QUrlQuery::isDetached();
extern bool C_ZNK9QUrlQuery10isDetachedEv(void* qthis); // 4
  // proto: static QChar QUrlQuery::defaultQueryPairDelimiter();
extern void* C_ZN9QUrlQuery25defaultQueryPairDelimiterEv(); // 2
  // proto:  void QUrlQuery::QUrlQuery();
extern void* C_ZN9QUrlQueryC2Ev(); // 3
  // proto:  void QUrlQuery::QUrlQuery(const QString & queryString);
extern void* C_ZN9QUrlQueryC2ERK7QString(void* arg0); // 3
  // proto:  void QUrlQuery::QUrlQuery(const QUrl & url);
extern void* C_ZN9QUrlQueryC2ERK4QUrl(void* arg0); // 3
  // proto:  void QUrlQuery::QUrlQuery(const QUrlQuery & other);
extern void* C_ZN9QUrlQueryC2ERKS_(void* arg0); // 3
  // proto:  void QUrlQuery::setQueryDelimiters(QChar valueDelimiter, QChar pairDelimiter);
extern void C_ZN9QUrlQuery18setQueryDelimitersE5QCharS0_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QUrlQuery::clear();
extern void C_ZN9QUrlQuery5clearEv(void* qthis); // 4
  // proto:  QChar QUrlQuery::queryPairDelimiter();
extern void* C_ZNK9QUrlQuery18queryPairDelimiterEv(void* qthis); // 4
  // proto:  void QUrlQuery::setQuery(const QString & queryString);
extern void C_ZN9QUrlQuery8setQueryERK7QString(void* qthis, void* arg0); // 4
  // proto: static QChar QUrlQuery::defaultQueryValueDelimiter();
extern void* C_ZN9QUrlQuery26defaultQueryValueDelimiterEv(); // 2
  // proto:  void QUrlQuery::removeQueryItem(const QString & key);
extern void C_ZN9QUrlQuery15removeQueryItemERK7QString(void* qthis, void* arg0); // 4
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

// class sizeof(QUrlQuery)=1
type QUrlQuery struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// addQueryItem(const class QString &, const class QString &)
func (this *QUrlQuery) Addqueryitem(args ...interface{}) () {
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
    // invoke: void addQueryItem(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QUrlQuery12addQueryItemERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QUrlQuery", "addQueryItem", args)
  }

  return
}

// queryValueDelimiter()
func (this *QUrlQuery) Queryvaluedelimiter(args ...interface{}) (ret interface{}) {
  // queryValueDelimiter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUrlQuery19queryValueDelimiterEv
    // invoke: QChar queryValueDelimiter()
    var ret0 = C.C_ZNK9QUrlQuery19queryValueDelimiterEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUrlQuery", "queryValueDelimiter", args)
  }

  return
}

// hasQueryItem(const class QString &)
func (this *QUrlQuery) Hasqueryitem(args ...interface{}) (ret interface{}) {
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
    // invoke: bool hasQueryItem(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QUrlQuery12hasQueryItemERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUrlQuery", "hasQueryItem", args)
  }

  return
}

// removeAllQueryItems(const class QString &)
func (this *QUrlQuery) Removeallqueryitems(args ...interface{}) () {
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
    // invoke: void removeAllQueryItems(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUrlQuery19removeAllQueryItemsERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrlQuery", "removeAllQueryItems", args)
  }

  return
}

// ~QUrlQuery()
func (this *QUrlQuery) Freequrlquery(args ...interface{}) () {
  // ~QUrlQuery()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQueryD0Ev
    // invoke: void ~QUrlQuery()
    C.C_ZN9QUrlQueryD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrlQuery", "~QUrlQuery", args)
  }

  return
}

// isEmpty()
func (this *QUrlQuery) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUrlQuery7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK9QUrlQuery7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUrlQuery", "isEmpty", args)
  }

  return
}

// swap(class QUrlQuery &)
func (this *QUrlQuery) Swap(args ...interface{}) () {
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
    // invoke: void swap(class QUrlQuery &)
    var arg0 = args[0].(QUrlQuery).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUrlQuery4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrlQuery", "swap", args)
  }

  return
}

// isDetached()
func (this *QUrlQuery) Isdetached(args ...interface{}) (ret interface{}) {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUrlQuery10isDetachedEv
    // invoke: bool isDetached()
    var ret0 = C.C_ZNK9QUrlQuery10isDetachedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUrlQuery", "isDetached", args)
  }

  return
}

// defaultQueryPairDelimiter()
func (this *QUrlQuery) Defaultquerypairdelimiter_S(args ...interface{}) (ret interface{}) {
  // defaultQueryPairDelimiter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQuery25defaultQueryPairDelimiterEv
    // invoke: QChar defaultQueryPairDelimiter()
    var ret0 = C.C_ZN9QUrlQuery25defaultQueryPairDelimiterEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUrlQuery", "defaultQueryPairDelimiter", args)
  }

  return
}

// QUrlQuery()
func NewQUrlQuery(args ...interface{}) *QUrlQuery {
  // QUrlQuery()
  // QUrlQuery(const class QString &)
  // QUrlQuery(const class QUrl &)
  // QUrlQuery(const class QUrlQuery &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QUrlQuery{}) // "const QUrlQuery &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQueryC1Ev
    // invoke: void QUrlQuery()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QUrlQueryC2Ev()
    return &QUrlQuery{qclsinst:qthis}
  case 1:
    // invoke: _ZN9QUrlQueryC1ERK7QString
    // invoke: void QUrlQuery(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QUrlQueryC2ERK7QString(arg0)
    return &QUrlQuery{qclsinst:qthis}
  case 2:
    // invoke: _ZN9QUrlQueryC1ERK4QUrl
    // invoke: void QUrlQuery(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QUrlQueryC2ERK4QUrl(arg0)
    return &QUrlQuery{qclsinst:qthis}
  case 3:
    // invoke: _ZN9QUrlQueryC1ERKS_
    // invoke: void QUrlQuery(const class QUrlQuery &)
    var arg0 = args[0].(QUrlQuery).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QUrlQueryC2ERKS_(arg0)
    return &QUrlQuery{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QUrlQuery", "QUrlQuery", args)
  }

  return nil // QUrlQuery{qclsinst:qthis}
}

// setQueryDelimiters(class QChar, class QChar)
func (this *QUrlQuery) Setquerydelimiters(args ...interface{}) () {
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
    // invoke: void setQueryDelimiters(class QChar, class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QUrlQuery18setQueryDelimitersE5QCharS0_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QUrlQuery", "setQueryDelimiters", args)
  }

  return
}

// clear()
func (this *QUrlQuery) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQuery5clearEv
    // invoke: void clear()
    C.C_ZN9QUrlQuery5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrlQuery", "clear", args)
  }

  return
}

// queryPairDelimiter()
func (this *QUrlQuery) Querypairdelimiter(args ...interface{}) (ret interface{}) {
  // queryPairDelimiter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QUrlQuery18queryPairDelimiterEv
    // invoke: QChar queryPairDelimiter()
    var ret0 = C.C_ZNK9QUrlQuery18queryPairDelimiterEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUrlQuery", "queryPairDelimiter", args)
  }

  return
}

// setQuery(const class QString &)
func (this *QUrlQuery) Setquery(args ...interface{}) () {
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
    // invoke: void setQuery(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUrlQuery8setQueryERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrlQuery", "setQuery", args)
  }

  return
}

// defaultQueryValueDelimiter()
func (this *QUrlQuery) Defaultqueryvaluedelimiter_S(args ...interface{}) (ret interface{}) {
  // defaultQueryValueDelimiter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QUrlQuery26defaultQueryValueDelimiterEv
    // invoke: QChar defaultQueryValueDelimiter()
    var ret0 = C.C_ZN9QUrlQuery26defaultQueryValueDelimiterEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QUrlQuery", "defaultQueryValueDelimiter", args)
  }

  return
}

// removeQueryItem(const class QString &)
func (this *QUrlQuery) Removequeryitem(args ...interface{}) () {
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
    // invoke: void removeQueryItem(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QUrlQuery15removeQueryItemERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrlQuery", "removeQueryItem", args)
  }

  return
}

// <= body block end

