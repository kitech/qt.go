package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QUrl::isDetached();
extern void C_ZNK4QUrl10isDetachedEv(void* qthis); // 4
  // proto:  void QUrl::~QUrl();
extern void C_ZN4QUrlD2Ev(void* qthis); // 4
  // proto:  bool QUrl::isRelative();
extern void C_ZNK4QUrl10isRelativeEv(void* qthis); // 4
  // proto:  bool QUrl::hasFragment();
extern void C_ZNK4QUrl11hasFragmentEv(void* qthis); // 4
  // proto:  void QUrl::setScheme(const QString & scheme);
extern void C_ZN4QUrl9setSchemeERK7QString(void* qthis, void* arg0); // 4
  // proto: static QUrl QUrl::fromUserInput(const QString & userInput);
extern void C_ZN4QUrl13fromUserInputERK7QString(void* arg0); // 4
  // proto: static QUrl QUrl::fromLocalFile(const QString & localfile);
extern void C_ZN4QUrl13fromLocalFileERK7QString(void* arg0); // 4
  // proto:  int QUrl::port(int defaultPort);
extern void C_ZNK4QUrl4portEi(void* qthis, int32_t arg0); // 4
  // proto: static QByteArray QUrl::toAce(const QString & );
extern void C_ZN4QUrl5toAceERK7QString(void* arg0); // 4
  // proto: static QString QUrl::fromAce(const QByteArray & );
extern void C_ZN4QUrl7fromAceERK10QByteArray(void* arg0); // 4
  // proto:  void QUrl::detach();
extern void C_ZN4QUrl6detachEv(void* qthis); // 4
  // proto:  void QUrl::QUrl();
extern void* C_ZN4QUrlC2Ev(); // 3
  // proto:  void QUrl::QUrl(const QUrl & copy);
extern void* C_ZN4QUrlC2ERKS_(void* arg0); // 3
  // proto:  bool QUrl::isEmpty();
extern void C_ZNK4QUrl7isEmptyEv(void* qthis); // 4
  // proto:  void QUrl::swap(QUrl & other);
extern void C_ZN4QUrl4swapERS_(void* qthis, void* arg0); // 2
  // proto: static QString QUrl::fromPercentEncoding(const QByteArray & );
extern void C_ZN4QUrl19fromPercentEncodingERK10QByteArray(void* arg0); // 4
  // proto:  bool QUrl::isParentOf(const QUrl & url);
extern void C_ZNK4QUrl10isParentOfERKS_(void* qthis, void* arg0); // 4
  // proto: static QStringList QUrl::idnWhitelist();
extern void C_ZN4QUrl12idnWhitelistEv(); // 4
  // proto:  QUrl QUrl::resolved(const QUrl & relative);
extern void C_ZNK4QUrl8resolvedERKS_(void* qthis, void* arg0); // 4
  // proto:  QString QUrl::errorString();
extern void C_ZNK4QUrl11errorStringEv(void* qthis); // 4
  // proto:  bool QUrl::isValid();
extern void C_ZNK4QUrl7isValidEv(void* qthis); // 4
  // proto:  bool QUrl::hasQuery();
extern void C_ZNK4QUrl8hasQueryEv(void* qthis); // 4
  // proto: static void QUrl::setIdnWhitelist(const QStringList & );
extern void C_ZN4QUrl15setIdnWhitelistERK11QStringList(void* arg0); // 4
  // proto:  QString QUrl::scheme();
extern void C_ZNK4QUrl6schemeEv(void* qthis); // 4
  // proto:  void QUrl::setQuery(const QUrlQuery & query);
extern void C_ZN4QUrl8setQueryERK9QUrlQuery(void* qthis, void* arg0); // 4
  // proto:  bool QUrl::isLocalFile();
extern void C_ZNK4QUrl11isLocalFileEv(void* qthis); // 4
  // proto: static QByteArray QUrl::toPercentEncoding(const QString & , const QByteArray & exclude, const QByteArray & include);
extern void C_ZN4QUrl17toPercentEncodingERK7QStringRK10QByteArrayS5_(void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QUrl::clear();
extern void C_ZN4QUrl5clearEv(void* qthis); // 4
  // proto:  void QUrl::setPort(int port);
extern void C_ZN4QUrl7setPortEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QUrl::toLocalFile();
extern void C_ZNK4QUrl11toLocalFileEv(void* qthis); // 4
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

// isDetached()
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
    var ret = C.C_ZNK4QUrl10isDetachedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "isDetached", args)
  }

}

// ~QUrl()
func (this *QUrl) FreeQUrl(args ...interface{}) () {
  // ~QUrl()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrlD0Ev
    // invoke: void ~QUrl()
    C.C_ZN4QUrlD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "~QUrl", args)
  }

}

// isRelative()
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
    var ret = C.C_ZNK4QUrl10isRelativeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "isRelative", args)
  }

}

// hasFragment()
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
    var ret = C.C_ZNK4QUrl11hasFragmentEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "hasFragment", args)
  }

}

// setScheme(const class QString &)
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
    C.C_ZN4QUrl9setSchemeERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "setScheme", args)
  }

}

// fromUserInput(const class QString &)
func (this *QUrl) fromUserInput_s(args ...interface{}) () {
  // fromUserInput(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl13fromUserInputERK7QString
    // invoke: QUrl fromUserInput(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QUrl13fromUserInputERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "fromUserInput", args)
  }

}

// fromLocalFile(const class QString &)
func (this *QUrl) fromLocalFile_s(args ...interface{}) () {
  // fromLocalFile(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl13fromLocalFileERK7QString
    // invoke: QUrl fromLocalFile(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QUrl13fromLocalFileERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "fromLocalFile", args)
  }

}

// port(int)
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
    var ret = C.C_ZNK4QUrl4portEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "port", args)
  }

}

// toAce(const class QString &)
func (this *QUrl) toAce_s(args ...interface{}) () {
  // toAce(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl5toAceERK7QString
    // invoke: QByteArray toAce(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QUrl5toAceERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "toAce", args)
  }

}

// fromAce(const class QByteArray &)
func (this *QUrl) fromAce_s(args ...interface{}) () {
  // fromAce(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl7fromAceERK10QByteArray
    // invoke: QString fromAce(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QUrl7fromAceERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "fromAce", args)
  }

}

// detach()
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
    C.C_ZN4QUrl6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "detach", args)
  }

}

// QUrl()
func NewQUrl(args ...interface{}) *QUrl {
  // QUrl()
  // QUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrlC1Ev
    // invoke: void QUrl()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN4QUrlC2Ev()
    return &QUrl{qclsinst:qthis}
  case 1:
    // invoke: _ZN4QUrlC1ERKS_
    // invoke: void QUrl(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN4QUrlC2ERKS_(arg0)
    return &QUrl{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QUrl", "QUrl", args)
  }

  return nil // QUrl{qclsinst:qthis}
}

// isEmpty()
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
    var ret = C.C_ZNK4QUrl7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "isEmpty", args)
  }

}

// swap(class QUrl &)
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
    C.C_ZN4QUrl4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "swap", args)
  }

}

// fromPercentEncoding(const class QByteArray &)
func (this *QUrl) fromPercentEncoding_s(args ...interface{}) () {
  // fromPercentEncoding(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl19fromPercentEncodingERK10QByteArray
    // invoke: QString fromPercentEncoding(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN4QUrl19fromPercentEncodingERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "fromPercentEncoding", args)
  }

}

// isParentOf(const class QUrl &)
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
    var ret = C.C_ZNK4QUrl10isParentOfERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "isParentOf", args)
  }

}

// idnWhitelist()
func (this *QUrl) idnWhitelist_s(args ...interface{}) () {
  // idnWhitelist()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl12idnWhitelistEv
    // invoke: QStringList idnWhitelist()
    C.C_ZN4QUrl12idnWhitelistEv()
  default:
    qtrt.ErrorResolve("QUrl", "idnWhitelist", args)
  }

}

// resolved(const class QUrl &)
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
    var ret = C.C_ZNK4QUrl8resolvedERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "resolved", args)
  }

}

// errorString()
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
    var ret = C.C_ZNK4QUrl11errorStringEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "errorString", args)
  }

}

// isValid()
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
    var ret = C.C_ZNK4QUrl7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "isValid", args)
  }

}

// hasQuery()
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
    var ret = C.C_ZNK4QUrl8hasQueryEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "hasQuery", args)
  }

}

// setIdnWhitelist(const class QStringList &)
func (this *QUrl) setIdnWhitelist_s(args ...interface{}) () {
  // setIdnWhitelist(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl15setIdnWhitelistERK11QStringList
    // invoke: void setIdnWhitelist(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QUrl15setIdnWhitelistERK11QStringList(arg0)
  default:
    qtrt.ErrorResolve("QUrl", "setIdnWhitelist", args)
  }

}

// scheme()
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
    var ret = C.C_ZNK4QUrl6schemeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "scheme", args)
  }

}

// setQuery(const class QUrlQuery &)
func (this *QUrl) setQuery(args ...interface{}) () {
  // setQuery(const class QUrlQuery &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrlQuery{}) // "const QUrlQuery &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl8setQueryERK9QUrlQuery
    // invoke: void setQuery(const class QUrlQuery &)
    var arg0 = args[0].(QUrlQuery).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN4QUrl8setQueryERK9QUrlQuery(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "setQuery", args)
  }

}

// isLocalFile()
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
    var ret = C.C_ZNK4QUrl11isLocalFileEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "isLocalFile", args)
  }

}

// toPercentEncoding(const class QString &, const class QByteArray &, const class QByteArray &)
func (this *QUrl) toPercentEncoding_s(args ...interface{}) () {
  // toPercentEncoding(const class QString &, const class QByteArray &, const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][2] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN4QUrl17toPercentEncodingERK7QStringRK10QByteArrayS5_
    // invoke: QByteArray toPercentEncoding(const class QString &, const class QByteArray &, const class QByteArray &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QByteArray).qclsinst
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN4QUrl17toPercentEncodingERK7QStringRK10QByteArrayS5_(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "toPercentEncoding", args)
  }

}

// clear()
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
    C.C_ZN4QUrl5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUrl", "clear", args)
  }

}

// setPort(int)
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
    C.C_ZN4QUrl7setPortEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QUrl", "setPort", args)
  }

}

// toLocalFile()
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
    var ret = C.C_ZNK4QUrl11toLocalFileEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QUrl", "toLocalFile", args)
  }

}

// <= body block end

