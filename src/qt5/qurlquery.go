package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QUrlQuery::QUrlQuery(const QString & queryString);
extern void* dector_ZN9QUrlQueryC1ERK7QString(void* arg0);
extern void _ZN9QUrlQueryC1ERK7QString(void* qthis, void* arg0);
  // proto:  void QUrlQuery::clear();
extern void _ZN9QUrlQuery5clearEv(void* qthis);
  // proto:  void QUrlQuery::setQuery(const QString & queryString);
extern void _ZN9QUrlQuery8setQueryERK7QString(void* qthis, void* arg0);
  // proto:  QChar QUrlQuery::queryValueDelimiter();
extern void _ZNK9QUrlQuery19queryValueDelimiterEv(void* qthis);
  // proto:  QChar QUrlQuery::queryPairDelimiter();
extern void _ZNK9QUrlQuery18queryPairDelimiterEv(void* qthis);
  // proto: static QChar QUrlQuery::defaultQueryValueDelimiter();
extern void demth_ZN9QUrlQuery26defaultQueryValueDelimiterEv();
  // proto:  void QUrlQuery::swap(QUrlQuery & other);
extern void demth_ZN9QUrlQuery4swapERS_(void* qthis, void* arg0);
  // proto:  bool QUrlQuery::isDetached();
extern void _ZNK9QUrlQuery10isDetachedEv(void* qthis);
  // proto:  void QUrlQuery::QUrlQuery();
extern void* dector_ZN9QUrlQueryC1Ev();
extern void _ZN9QUrlQueryC1Ev(void* qthis);
  // proto:  void QUrlQuery::setQueryDelimiters(QChar valueDelimiter, QChar pairDelimiter);
extern void _ZN9QUrlQuery18setQueryDelimitersE5QCharS0_(void* qthis, void* arg0, void* arg1);
  // proto:  void QUrlQuery::~QUrlQuery();
extern void _ZN9QUrlQueryD0Ev(void* qthis);
  // proto:  void QUrlQuery::removeAllQueryItems(const QString & key);
extern void _ZN9QUrlQuery19removeAllQueryItemsERK7QString(void* qthis, void* arg0);
  // proto:  bool QUrlQuery::isEmpty();
extern void _ZNK9QUrlQuery7isEmptyEv(void* qthis);
  // proto:  void QUrlQuery::removeQueryItem(const QString & key);
extern void _ZN9QUrlQuery15removeQueryItemERK7QString(void* qthis, void* arg0);
  // proto: static QChar QUrlQuery::defaultQueryPairDelimiter();
extern void demth_ZN9QUrlQuery25defaultQueryPairDelimiterEv();
  // proto:  void QUrlQuery::QUrlQuery(const QUrl & url);
extern void* dector_ZN9QUrlQueryC1ERK4QUrl(void* arg0);
extern void _ZN9QUrlQueryC1ERK4QUrl(void* qthis, void* arg0);
  // proto:  void QUrlQuery::addQueryItem(const QString & key, const QString & value);
extern void _ZN9QUrlQuery12addQueryItemERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QUrlQuery::QUrlQuery(const QUrlQuery & other);
extern void* dector_ZN9QUrlQueryC1ERKS_(void* arg0);
extern void _ZN9QUrlQueryC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QUrlQuery::hasQueryItem(const QString & key);
extern void _ZNK9QUrlQuery12hasQueryItemERK7QString(void* qthis, void* arg0);
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

  // proto:  void QUrlQuery::QUrlQuery(const QString & queryString);
func NewQUrlQuery(args ...interface{}) QUrlQuery {
  return QUrlQuery{}
}

  // proto:  void QUrlQuery::clear();
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
    // invoke: void clear()
    C._ZN9QUrlQuery5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrlQuery", "clear", args)
  }

}

  // proto:  void QUrlQuery::setQuery(const QString & queryString);
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
    // invoke: void setQuery(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QUrlQuery8setQueryERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrlQuery", "setQuery", args)
  }

}

  // proto:  QChar QUrlQuery::queryValueDelimiter();
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
    // invoke: QChar queryValueDelimiter()
    C._ZNK9QUrlQuery19queryValueDelimiterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrlQuery", "queryValueDelimiter", args)
  }

}

  // proto:  QChar QUrlQuery::queryPairDelimiter();
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
    // invoke: QChar queryPairDelimiter()
    C._ZNK9QUrlQuery18queryPairDelimiterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrlQuery", "queryPairDelimiter", args)
  }

}

  // proto: static QChar QUrlQuery::defaultQueryValueDelimiter();
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

  // proto:  void QUrlQuery::swap(QUrlQuery & other);
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
    // invoke: void swap(class QUrlQuery &)
    var arg0 = args[0].(QUrlQuery).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN9QUrlQuery4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrlQuery", "swap", args)
  }

}

  // proto:  bool QUrlQuery::isDetached();
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
    // invoke: bool isDetached()
    C._ZNK9QUrlQuery10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrlQuery", "isDetached", args)
  }

}

  // proto:  void QUrlQuery::setQueryDelimiters(QChar valueDelimiter, QChar pairDelimiter);
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
    // invoke: void setQueryDelimiters(class QChar, class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN9QUrlQuery18setQueryDelimitersE5QCharS0_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QUrlQuery", "setQueryDelimiters", args)
  }

}

  // proto:  void QUrlQuery::~QUrlQuery();
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

  // proto:  void QUrlQuery::removeAllQueryItems(const QString & key);
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
    // invoke: void removeAllQueryItems(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QUrlQuery19removeAllQueryItemsERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrlQuery", "removeAllQueryItems", args)
  }

}

  // proto:  bool QUrlQuery::isEmpty();
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
    // invoke: bool isEmpty()
    C._ZNK9QUrlQuery7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrlQuery", "isEmpty", args)
  }

}

  // proto:  void QUrlQuery::removeQueryItem(const QString & key);
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
    // invoke: void removeQueryItem(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QUrlQuery15removeQueryItemERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrlQuery", "removeQueryItem", args)
  }

}

  // proto: static QChar QUrlQuery::defaultQueryPairDelimiter();
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

  // proto:  void QUrlQuery::addQueryItem(const QString & key, const QString & value);
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
    // invoke: void addQueryItem(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN9QUrlQuery12addQueryItemERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QUrlQuery", "addQueryItem", args)
  }

}

  // proto:  bool QUrlQuery::hasQueryItem(const QString & key);
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
    // invoke: bool hasQueryItem(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QUrlQuery12hasQueryItemERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrlQuery", "hasQueryItem", args)
  }

}

// <= body block end

