package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtCore/qurl.h
// dst-file: /src/core/qurl.go
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
  // proto:  bool QUrl::isLocalFile();
extern void _ZNK4QUrl11isLocalFileEv(void* qthis);
  // proto:  bool QUrl::isEmpty();
extern void _ZNK4QUrl7isEmptyEv(void* qthis);
  // proto:  void QUrl::setQuery(const QUrlQuery & query);
extern void _ZN4QUrl8setQueryERK9QUrlQuery(void* qthis, void* arg0);
  // proto: static QStringList QUrl::idnWhitelist();
extern void _ZN4QUrl12idnWhitelistEv();
  // proto:  void QUrl::~QUrl();
extern void _ZN4QUrlD0Ev(void* qthis);
  // proto:  void QUrl::setScheme(const QString & scheme);
extern void _ZN4QUrl9setSchemeERK7QString(void* qthis, void* arg0);
  // proto:  bool QUrl::isParentOf(const QUrl & url);
extern void _ZNK4QUrl10isParentOfERKS_(void* qthis, void* arg0);
  // proto:  QString QUrl::errorString();
extern void _ZNK4QUrl11errorStringEv(void* qthis);
  // proto:  int QUrl::port(int defaultPort);
extern void _ZNK4QUrl4portEi(void* qthis, int32_t arg0);
  // proto:  void QUrl::setPort(int port);
extern void _ZN4QUrl7setPortEi(void* qthis, int32_t arg0);
  // proto:  void QUrl::QUrl(const QUrl & copy);
extern void* dector_ZN4QUrlC1ERKS_(void* arg0);
extern void _ZN4QUrlC1ERKS_(void* qthis, void* arg0);
  // proto: static QString QUrl::fromAce(const QByteArray & );
extern void _ZN4QUrl7fromAceERK10QByteArray(void* arg0);
  // proto:  QUrl QUrl::resolved(const QUrl & relative);
extern void _ZNK4QUrl8resolvedERKS_(void* qthis, void* arg0);
  // proto:  QString QUrl::toLocalFile();
extern void _ZNK4QUrl11toLocalFileEv(void* qthis);
  // proto:  void QUrl::detach();
extern void _ZN4QUrl6detachEv(void* qthis);
  // proto:  bool QUrl::hasFragment();
extern void _ZNK4QUrl11hasFragmentEv(void* qthis);
  // proto: static QByteArray QUrl::toAce(const QString & );
extern void _ZN4QUrl5toAceERK7QString(void* arg0);
  // proto:  bool QUrl::hasQuery();
extern void _ZNK4QUrl8hasQueryEv(void* qthis);
  // proto: static QUrl QUrl::fromLocalFile(const QString & localfile);
extern void _ZN4QUrl13fromLocalFileERK7QString(void* arg0);
  // proto:  bool QUrl::isValid();
extern void _ZNK4QUrl7isValidEv(void* qthis);
  // proto:  void QUrl::QUrl();
extern void* dector_ZN4QUrlC1Ev();
extern void _ZN4QUrlC1Ev(void* qthis);
  // proto:  bool QUrl::isDetached();
extern void _ZNK4QUrl10isDetachedEv(void* qthis);
  // proto:  bool QUrl::isRelative();
extern void _ZNK4QUrl10isRelativeEv(void* qthis);
  // proto:  QString QUrl::scheme();
extern void _ZNK4QUrl6schemeEv(void* qthis);
  // proto: static QByteArray QUrl::toPercentEncoding(const QString & , const QByteArray & exclude, const QByteArray & include);
extern void _ZN4QUrl17toPercentEncodingERK7QStringRK10QByteArrayS5_(void* arg0, void* arg1, void* arg2);
  // proto: static void QUrl::setIdnWhitelist(const QStringList & );
extern void _ZN4QUrl15setIdnWhitelistERK11QStringList(void* arg0);
  // proto:  void QUrl::swap(QUrl & other);
extern void demth_ZN4QUrl4swapERS_(void* qthis, void* arg0);
  // proto: static QString QUrl::fromPercentEncoding(const QByteArray & );
extern void _ZN4QUrl19fromPercentEncodingERK10QByteArray(void* arg0);
  // proto: static QUrl QUrl::fromUserInput(const QString & userInput);
extern void _ZN4QUrl13fromUserInputERK7QString(void* arg0);
  // proto:  void QUrl::clear();
extern void _ZN4QUrl5clearEv(void* qthis);
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

// class sizeof(QUrl)=8
type QUrl struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  bool QUrl::isLocalFile();
func (this *QUrl) isLocalFile(args ...interface{}) () {
  // isLocalFile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl11isLocalFileEv
    // invoke: bool isLocalFile()
    C._ZNK4QUrl11isLocalFileEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "isLocalFile", args)
  }

}

  // proto:  bool QUrl::isEmpty();
func (this *QUrl) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK4QUrl7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "isEmpty", args)
  }

}

  // proto:  void QUrl::setQuery(const QUrlQuery & query);
func (this *QUrl) setQuery(args ...interface{}) () {
  // setQuery(const class QUrlQuery &)
  // setQuery(const class QString &, enum QUrl::ParsingMode)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrlQuery{}) // "const QUrlQuery &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "QUrl::ParsingMode"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl8setQueryERK9QUrlQuery
    // invoke: void setQuery(const class QUrlQuery &)
    var arg0 = args[0].(QUrlQuery).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN4QUrl8setQueryERK9QUrlQuery(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "setQuery", args)
  }

}

  // proto: static QStringList QUrl::idnWhitelist();
func (this *QUrl) idnWhitelist_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrl", "idnWhitelist", args)
  }

}

  // proto:  void QUrl::~QUrl();
func (this *QUrl) FreeQUrl(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrl", "~QUrl", args)
  }

}

  // proto:  void QUrl::setScheme(const QString & scheme);
func (this *QUrl) setScheme(args ...interface{}) () {
  // setScheme(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl9setSchemeERK7QString
    // invoke: void setScheme(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN4QUrl9setSchemeERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "setScheme", args)
  }

}

  // proto:  bool QUrl::isParentOf(const QUrl & url);
func (this *QUrl) isParentOf(args ...interface{}) () {
  // isParentOf(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl10isParentOfERKS_
    // invoke: bool isParentOf(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK4QUrl10isParentOfERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "isParentOf", args)
  }

}

  // proto:  QString QUrl::errorString();
func (this *QUrl) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl11errorStringEv
    // invoke: QString errorString()
    C._ZNK4QUrl11errorStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "errorString", args)
  }

}

  // proto:  int QUrl::port(int defaultPort);
func (this *QUrl) port(args ...interface{}) () {
  // port(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl4portEi
    // invoke: int port(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK4QUrl4portEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "port", args)
  }

}

  // proto:  void QUrl::setPort(int port);
func (this *QUrl) setPort(args ...interface{}) () {
  // setPort(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl7setPortEi
    // invoke: void setPort(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN4QUrl7setPortEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "setPort", args)
  }

}

  // proto:  void QUrl::QUrl(const QUrl & copy);
func NewQUrl(args ...interface{}) QUrl {
  return QUrl{}
}

  // proto: static QString QUrl::fromAce(const QByteArray & );
func (this *QUrl) fromAce_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrl", "fromAce", args)
  }

}

  // proto:  QUrl QUrl::resolved(const QUrl & relative);
func (this *QUrl) resolved(args ...interface{}) () {
  // resolved(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl8resolvedERKS_
    // invoke: QUrl resolved(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK4QUrl8resolvedERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "resolved", args)
  }

}

  // proto:  QString QUrl::toLocalFile();
func (this *QUrl) toLocalFile(args ...interface{}) () {
  // toLocalFile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl11toLocalFileEv
    // invoke: QString toLocalFile()
    C._ZNK4QUrl11toLocalFileEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "toLocalFile", args)
  }

}

  // proto:  void QUrl::detach();
func (this *QUrl) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl6detachEv
    // invoke: void detach()
    C._ZN4QUrl6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "detach", args)
  }

}

  // proto:  bool QUrl::hasFragment();
func (this *QUrl) hasFragment(args ...interface{}) () {
  // hasFragment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl11hasFragmentEv
    // invoke: bool hasFragment()
    C._ZNK4QUrl11hasFragmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "hasFragment", args)
  }

}

  // proto: static QByteArray QUrl::toAce(const QString & );
func (this *QUrl) toAce_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrl", "toAce", args)
  }

}

  // proto:  bool QUrl::hasQuery();
func (this *QUrl) hasQuery(args ...interface{}) () {
  // hasQuery()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl8hasQueryEv
    // invoke: bool hasQuery()
    C._ZNK4QUrl8hasQueryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "hasQuery", args)
  }

}

  // proto: static QUrl QUrl::fromLocalFile(const QString & localfile);
func (this *QUrl) fromLocalFile_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrl", "fromLocalFile", args)
  }

}

  // proto:  bool QUrl::isValid();
func (this *QUrl) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl7isValidEv
    // invoke: bool isValid()
    C._ZNK4QUrl7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "isValid", args)
  }

}

  // proto:  bool QUrl::isDetached();
func (this *QUrl) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl10isDetachedEv
    // invoke: bool isDetached()
    C._ZNK4QUrl10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "isDetached", args)
  }

}

  // proto:  bool QUrl::isRelative();
func (this *QUrl) isRelative(args ...interface{}) () {
  // isRelative()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl10isRelativeEv
    // invoke: bool isRelative()
    C._ZNK4QUrl10isRelativeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "isRelative", args)
  }

}

  // proto:  QString QUrl::scheme();
func (this *QUrl) scheme(args ...interface{}) () {
  // scheme()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK4QUrl6schemeEv
    // invoke: QString scheme()
    C._ZNK4QUrl6schemeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "scheme", args)
  }

}

  // proto: static QByteArray QUrl::toPercentEncoding(const QString & , const QByteArray & exclude, const QByteArray & include);
func (this *QUrl) toPercentEncoding_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrl", "toPercentEncoding", args)
  }

}

  // proto: static void QUrl::setIdnWhitelist(const QStringList & );
func (this *QUrl) setIdnWhitelist_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrl", "setIdnWhitelist", args)
  }

}

  // proto:  void QUrl::swap(QUrl & other);
func (this *QUrl) swap(args ...interface{}) () {
  // swap(class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl4swapERS_
    // invoke: void swap(class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN4QUrl4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "swap", args)
  }

}

  // proto: static QString QUrl::fromPercentEncoding(const QByteArray & );
func (this *QUrl) fromPercentEncoding_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrl", "fromPercentEncoding", args)
  }

}

  // proto: static QUrl QUrl::fromUserInput(const QString & userInput);
func (this *QUrl) fromUserInput_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUrl", "fromUserInput", args)
  }

}

  // proto:  void QUrl::clear();
func (this *QUrl) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl5clearEv
    // invoke: void clear()
    C._ZN4QUrl5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "clear", args)
  }

}

// <= body block end

