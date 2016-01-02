package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto: static QTextDocumentFragment QTextDocumentFragment::fromHtml(const QString & html, const QTextDocument * resourceProvider);
extern void _ZN21QTextDocumentFragment8fromHtmlERK7QStringPK13QTextDocument(void* arg0, void* arg1);
  // proto:  void QTextDocumentFragment::QTextDocumentFragment(const QTextDocumentFragment & rhs);
extern void* dector_ZN21QTextDocumentFragmentC1ERKS_(void* arg0);
extern void _ZN21QTextDocumentFragmentC1ERKS_(void* qthis, void* arg0);
  // proto: static QTextDocumentFragment QTextDocumentFragment::fromPlainText(const QString & plainText);
extern void _ZN21QTextDocumentFragment13fromPlainTextERK7QString(void* arg0);
  // proto:  QString QTextDocumentFragment::toHtml(const QByteArray & encoding);
extern void _ZNK21QTextDocumentFragment6toHtmlERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QTextDocumentFragment::~QTextDocumentFragment();
extern void _ZN21QTextDocumentFragmentD0Ev(void* qthis);
  // proto: static QTextDocumentFragment QTextDocumentFragment::fromHtml(const QString & html);
extern void _ZN21QTextDocumentFragment8fromHtmlERK7QString(void* arg0);
  // proto:  void QTextDocumentFragment::QTextDocumentFragment();
extern void* dector_ZN21QTextDocumentFragmentC1Ev();
extern void _ZN21QTextDocumentFragmentC1Ev(void* qthis);
  // proto:  QString QTextDocumentFragment::toPlainText();
extern void _ZNK21QTextDocumentFragment11toPlainTextEv(void* qthis);
  // proto:  void QTextDocumentFragment::QTextDocumentFragment(const QTextCursor & range);
extern void* dector_ZN21QTextDocumentFragmentC1ERK11QTextCursor(void* arg0);
extern void _ZN21QTextDocumentFragmentC1ERK11QTextCursor(void* qthis, void* arg0);
  // proto:  void QTextDocumentFragment::QTextDocumentFragment(const QTextDocument * document);
extern void* dector_ZN21QTextDocumentFragmentC1EPK13QTextDocument(void* arg0);
extern void _ZN21QTextDocumentFragmentC1EPK13QTextDocument(void* qthis, void* arg0);
  // proto:  bool QTextDocumentFragment::isEmpty();
extern void _ZNK21QTextDocumentFragment7isEmptyEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto: static QTextDocumentFragment QTextDocumentFragment::fromHtml(const QString & html, const QTextDocument * resourceProvider);
func (this *QTextDocumentFragment) fromHtml_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextDocumentFragment", "fromHtml", args)
  }

}

  // proto:  void QTextDocumentFragment::QTextDocumentFragment(const QTextDocumentFragment & rhs);
func NewQTextDocumentFragment(args ...interface{}) QTextDocumentFragment {
  return QTextDocumentFragment{}
}

  // proto: static QTextDocumentFragment QTextDocumentFragment::fromPlainText(const QString & plainText);
func (this *QTextDocumentFragment) fromPlainText_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextDocumentFragment", "fromPlainText", args)
  }

}

  // proto:  QString QTextDocumentFragment::toHtml(const QByteArray & encoding);
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

  // proto:  void QTextDocumentFragment::~QTextDocumentFragment();
func (this *QTextDocumentFragment) FreeQTextDocumentFragment(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextDocumentFragment", "~QTextDocumentFragment", args)
  }

}

  // proto:  QString QTextDocumentFragment::toPlainText();
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

  // proto:  bool QTextDocumentFragment::isEmpty();
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

// <= body block end

