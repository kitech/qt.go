package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtGui/qtextdocumentfragment.h
// dst-file: /src/gui/qtextdocumentfragment.go
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
  // proto:  QString QTextDocumentFragment::toHtml(const QByteArray & encoding);
extern void _ZNK21QTextDocumentFragment6toHtmlERK10QByteArray(void* qthis, void* arg0); // 4
  // proto: static QTextDocumentFragment QTextDocumentFragment::fromPlainText(const QString & plainText);
extern void _ZN21QTextDocumentFragment13fromPlainTextERK7QString(void* arg0); // 4
  // proto:  void QTextDocumentFragment::~QTextDocumentFragment();
extern void _ZN21QTextDocumentFragmentD2Ev(void* qthis); // 4
  // proto:  bool QTextDocumentFragment::isEmpty();
extern void _ZNK21QTextDocumentFragment7isEmptyEv(void* qthis); // 4
  // proto:  QString QTextDocumentFragment::toPlainText();
extern void _ZNK21QTextDocumentFragment11toPlainTextEv(void* qthis); // 4
  // proto:  void QTextDocumentFragment::QTextDocumentFragment(const QTextDocument * document);
extern void _ZN21QTextDocumentFragmentC2EPK13QTextDocument(void* qthis, void* arg0); // 3
  // proto:  void QTextDocumentFragment::QTextDocumentFragment();
extern void _ZN21QTextDocumentFragmentC2Ev(void* qthis); // 3
  // proto:  void QTextDocumentFragment::QTextDocumentFragment(const QTextDocumentFragment & rhs);
extern void _ZN21QTextDocumentFragmentC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QTextDocumentFragment::QTextDocumentFragment(const QTextCursor & range);
extern void _ZN21QTextDocumentFragmentC2ERK11QTextCursor(void* qthis, void* arg0); // 3
  // proto: static QTextDocumentFragment QTextDocumentFragment::fromHtml(const QString & html);
extern void _ZN21QTextDocumentFragment8fromHtmlERK7QString(void* arg0); // 4
  // proto: static QTextDocumentFragment QTextDocumentFragment::fromHtml(const QString & html, const QTextDocument * resourceProvider);
extern void _ZN21QTextDocumentFragment8fromHtmlERK7QStringPK13QTextDocument(void* arg0, void* arg1); // 4
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

// class sizeof(QTextDocumentFragment)=8
type QTextDocumentFragment struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// toHtml(const class QByteArray &)
func (this *QTextDocumentFragment) toHtml(args ...interface{}) () {
  // toHtml(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QTextDocumentFragment6toHtmlERK10QByteArray
    // invoke: QString toHtml(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK21QTextDocumentFragment6toHtmlERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentFragment", "toHtml", args)
  }

}

// fromPlainText(const class QString &)
func (this *QTextDocumentFragment) fromPlainText_s(args ...interface{}) () {
  // fromPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QTextDocumentFragment13fromPlainTextERK7QString
    // invoke: QTextDocumentFragment fromPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN21QTextDocumentFragment13fromPlainTextERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentFragment", "fromPlainText", args)
  }

}

// ~QTextDocumentFragment()
func (this *QTextDocumentFragment) FreeQTextDocumentFragment(args ...interface{}) () {
  // ~QTextDocumentFragment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QTextDocumentFragmentD0Ev
    // invoke: void ~QTextDocumentFragment()
    C._ZN21QTextDocumentFragmentD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocumentFragment", "~QTextDocumentFragment", args)
  }

}

// isEmpty()
func (this *QTextDocumentFragment) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QTextDocumentFragment7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK21QTextDocumentFragment7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocumentFragment", "isEmpty", args)
  }

}

// toPlainText()
func (this *QTextDocumentFragment) toPlainText(args ...interface{}) () {
  // toPlainText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QTextDocumentFragment11toPlainTextEv
    // invoke: QString toPlainText()
    C._ZNK21QTextDocumentFragment11toPlainTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocumentFragment", "toPlainText", args)
  }

}

// QTextDocumentFragment(const class QTextDocument *)
func NewQTextDocumentFragment(args ...interface{}) QTextDocumentFragment {
  // QTextDocumentFragment(const class QTextDocument *)
  // QTextDocumentFragment()
  // QTextDocumentFragment(const class QTextDocumentFragment &)
  // QTextDocumentFragment(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "const QTextDocument *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QTextDocumentFragment{}) // "const QTextDocumentFragment &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QTextDocumentFragmentC1EPK13QTextDocument
    // invoke: void QTextDocumentFragment(const class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN21QTextDocumentFragmentC2EPK13QTextDocument(qthis, arg0)
  case 1:
    // invoke: _ZN21QTextDocumentFragmentC1Ev
    // invoke: void QTextDocumentFragment()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN21QTextDocumentFragmentC2Ev(qthis)
  case 2:
    // invoke: _ZN21QTextDocumentFragmentC1ERKS_
    // invoke: void QTextDocumentFragment(const class QTextDocumentFragment &)
    var arg0 = args[0].(QTextDocumentFragment).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN21QTextDocumentFragmentC2ERKS_(qthis, arg0)
  case 3:
    // invoke: _ZN21QTextDocumentFragmentC1ERK11QTextCursor
    // invoke: void QTextDocumentFragment(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN21QTextDocumentFragmentC2ERK11QTextCursor(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentFragment", "QTextDocumentFragment", args)
  }

  return QTextDocumentFragment{}
}

// fromHtml(const class QString &)
func (this *QTextDocumentFragment) fromHtml_s(args ...interface{}) () {
  // fromHtml(const class QString &)
  // fromHtml(const class QString &, const class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QTextDocument{}) // "const QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QTextDocumentFragment8fromHtmlERK7QString
    // invoke: QTextDocumentFragment fromHtml(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN21QTextDocumentFragment8fromHtmlERK7QString(arg0)
  case 1:
    // invoke: _ZN21QTextDocumentFragment8fromHtmlERK7QStringPK13QTextDocument
    // invoke: QTextDocumentFragment fromHtml(const class QString &, const class QTextDocument *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextDocument).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN21QTextDocumentFragment8fromHtmlERK7QStringPK13QTextDocument(arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextDocumentFragment", "fromHtml", args)
  }

}

// <= body block end

