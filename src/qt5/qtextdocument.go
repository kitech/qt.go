package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qtextdocument.h
// dst-file: /src/gui/qtextdocument.go
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
  // proto:  void QTextDocument::setDefaultStyleSheet(const QString & sheet);
extern void _ZN13QTextDocument20setDefaultStyleSheetERK7QString(void* qthis, void* arg0);
  // proto:  QAbstractTextDocumentLayout * QTextDocument::documentLayout();
extern void _ZNK13QTextDocument14documentLayoutEv(void* qthis);
  // proto:  bool QTextDocument::isModified();
extern void _ZNK13QTextDocument10isModifiedEv(void* qthis);
  // proto:  int QTextDocument::revision();
extern void _ZNK13QTextDocument8revisionEv(void* qthis);
  // proto:  QSizeF QTextDocument::pageSize();
extern void _ZNK13QTextDocument8pageSizeEv(void* qthis);
  // proto:  void QTextDocument::redo(QTextCursor * cursor);
extern void _ZN13QTextDocument4redoEP11QTextCursor(void* qthis, void* arg0);
  // proto:  int QTextDocument::lineCount();
extern void _ZNK13QTextDocument9lineCountEv(void* qthis);
  // proto:  void QTextDocument::print(QPagedPaintDevice * printer);
extern void _ZNK13QTextDocument5printEP17QPagedPaintDevice(void* qthis, void* arg0);
  // proto:  QTextDocumentPrivate * QTextDocument::docHandle();
extern void _ZNK13QTextDocument9docHandleEv(void* qthis);
  // proto:  QString QTextDocument::toHtml(const QByteArray & encoding);
extern void _ZNK13QTextDocument6toHtmlERK10QByteArray(void* qthis, void* arg0);
  // proto:  int QTextDocument::availableUndoSteps();
extern void _ZNK13QTextDocument18availableUndoStepsEv(void* qthis);
  // proto:  void QTextDocument::setUndoRedoEnabled(bool enable);
extern void _ZN13QTextDocument18setUndoRedoEnabledEb(void* qthis, bool arg0);
  // proto:  void QTextDocument::undo(QTextCursor * cursor);
extern void _ZN13QTextDocument4undoEP11QTextCursor(void* qthis, void* arg0);
  // proto:  QString QTextDocument::toPlainText();
extern void _ZNK13QTextDocument11toPlainTextEv(void* qthis);
  // proto:  void QTextDocument::addResource(int type, const QUrl & name, const QVariant & resource);
extern void _ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  QSizeF QTextDocument::size();
extern void _ZNK13QTextDocument4sizeEv(void* qthis);
  // proto:  QTextObject * QTextDocument::object(int objectIndex);
extern void _ZNK13QTextDocument6objectEi(void* qthis, int arg0);
  // proto:  QTextDocument * QTextDocument::clone(QObject * parent);
extern void _ZNK13QTextDocument5cloneEP7QObject(void* qthis, void* arg0);
  // proto:  void QTextDocument::markContentsDirty(int from, int length);
extern void _ZN13QTextDocument17markContentsDirtyEii(void* qthis, int arg0, int arg1);
  // proto:  void QTextDocument::QTextDocument(QObject * parent);
extern void* dector_ZN13QTextDocumentC1EP7QObject(void* arg0);
extern void _ZN13QTextDocumentC1EP7QObject(void* qthis, void* arg0);
  // proto:  int QTextDocument::characterCount();
extern void _ZNK13QTextDocument14characterCountEv(void* qthis);
  // proto:  QTextFrame * QTextDocument::rootFrame();
extern void _ZNK13QTextDocument9rootFrameEv(void* qthis);
  // proto:  QTextBlock QTextDocument::firstBlock();
extern void _ZNK13QTextDocument10firstBlockEv(void* qthis);
  // proto:  int QTextDocument::blockCount();
extern void _ZNK13QTextDocument10blockCountEv(void* qthis);
  // proto:  qreal QTextDocument::idealWidth();
extern void _ZNK13QTextDocument10idealWidthEv(void* qthis);
  // proto:  void QTextDocument::adjustSize();
extern void _ZN13QTextDocument10adjustSizeEv(void* qthis);
  // proto:  bool QTextDocument::isRedoAvailable();
extern void _ZNK13QTextDocument15isRedoAvailableEv(void* qthis);
  // proto:  QVector<QTextFormat> QTextDocument::allFormats();
extern void _ZNK13QTextDocument10allFormatsEv(void* qthis);
  // proto:  QString QTextDocument::defaultStyleSheet();
extern void _ZNK13QTextDocument17defaultStyleSheetEv(void* qthis);
  // proto:  QTextBlock QTextDocument::lastBlock();
extern void _ZNK13QTextDocument9lastBlockEv(void* qthis);
  // proto:  void QTextDocument::QTextDocument(const QString & text, QObject * parent);
extern void* dector_ZN13QTextDocumentC1ERK7QStringP7QObject(void* arg0, void* arg1);
extern void _ZN13QTextDocumentC1ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  bool QTextDocument::useDesignMetrics();
extern void _ZNK13QTextDocument16useDesignMetricsEv(void* qthis);
  // proto:  int QTextDocument::pageCount();
extern void _ZNK13QTextDocument9pageCountEv(void* qthis);
  // proto:  void QTextDocument::setTextWidth(qreal width);
extern void _ZN13QTextDocument12setTextWidthEd(void* qthis, double arg0);
  // proto:  void QTextDocument::setDocumentMargin(qreal margin);
extern void _ZN13QTextDocument17setDocumentMarginEd(void* qthis, double arg0);
  // proto:  bool QTextDocument::isUndoAvailable();
extern void _ZNK13QTextDocument15isUndoAvailableEv(void* qthis);
  // proto:  qreal QTextDocument::indentWidth();
extern void _ZNK13QTextDocument11indentWidthEv(void* qthis);
  // proto:  void QTextDocument::setUseDesignMetrics(bool b);
extern void _ZN13QTextDocument19setUseDesignMetricsEb(void* qthis, bool arg0);
  // proto:  void QTextDocument::setDocumentLayout(QAbstractTextDocumentLayout * layout);
extern void _ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout(void* qthis, void* arg0);
  // proto:  void QTextDocument::setIndentWidth(qreal width);
extern void _ZN13QTextDocument14setIndentWidthEd(void* qthis, double arg0);
  // proto:  QUrl QTextDocument::baseUrl();
extern void _ZNK13QTextDocument7baseUrlEv(void* qthis);
  // proto:  QTextFrame * QTextDocument::frameAt(int pos);
extern void _ZNK13QTextDocument7frameAtEi(void* qthis, int arg0);
  // proto:  void QTextDocument::QTextDocument(const QTextDocument & );
extern void* dector_ZN13QTextDocumentC1ERKS_(void* arg0);
extern void _ZN13QTextDocumentC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTextDocument::setDefaultTextOption(const QTextOption & option);
extern void _ZN13QTextDocument20setDefaultTextOptionERK11QTextOption(void* qthis, void* arg0);
  // proto:  QFont QTextDocument::defaultFont();
extern void _ZNK13QTextDocument11defaultFontEv(void* qthis);
  // proto:  QTextBlock QTextDocument::findBlockByNumber(int blockNumber);
extern void _ZNK13QTextDocument17findBlockByNumberEi(void* qthis, int arg0);
  // proto:  QTextOption QTextDocument::defaultTextOption();
extern void _ZNK13QTextDocument17defaultTextOptionEv(void* qthis);
  // proto:  QTextBlock QTextDocument::findBlock(int pos);
extern void _ZNK13QTextDocument9findBlockEi(void* qthis, int arg0);
  // proto:  void QTextDocument::setBaseUrl(const QUrl & url);
extern void _ZN13QTextDocument10setBaseUrlERK4QUrl(void* qthis, void* arg0);
  // proto:  void QTextDocument::appendUndoItem(QAbstractUndoItem * );
extern void _ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem(void* qthis, void* arg0);
  // proto:  void QTextDocument::redo();
extern void _ZN13QTextDocument4redoEv(void* qthis);
  // proto:  void QTextDocument::drawContents(QPainter * painter, const QRectF & rect);
extern void _ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF(void* qthis, void* arg0, void* arg1);
  // proto:  QTextBlock QTextDocument::findBlockByLineNumber(int blockNumber);
extern void _ZNK13QTextDocument21findBlockByLineNumberEi(void* qthis, int arg0);
  // proto:  void QTextDocument::undo();
extern void _ZN13QTextDocument4undoEv(void* qthis);
  // proto:  qreal QTextDocument::textWidth();
extern void _ZNK13QTextDocument9textWidthEv(void* qthis);
  // proto:  const QMetaObject * QTextDocument::metaObject();
extern void _ZNK13QTextDocument10metaObjectEv(void* qthis);
  // proto:  int QTextDocument::availableRedoSteps();
extern void _ZNK13QTextDocument18availableRedoStepsEv(void* qthis);
  // proto:  QChar QTextDocument::characterAt(int pos);
extern void _ZNK13QTextDocument11characterAtEi(void* qthis, int arg0);
  // proto:  void QTextDocument::setDefaultFont(const QFont & font);
extern void _ZN13QTextDocument14setDefaultFontERK5QFont(void* qthis, void* arg0);
  // proto:  QTextObject * QTextDocument::objectForFormat(const QTextFormat & );
extern void _ZNK13QTextDocument15objectForFormatERK11QTextFormat(void* qthis, void* arg0);
  // proto:  bool QTextDocument::isEmpty();
extern void _ZNK13QTextDocument7isEmptyEv(void* qthis);
  // proto:  bool QTextDocument::isUndoRedoEnabled();
extern void _ZNK13QTextDocument17isUndoRedoEnabledEv(void* qthis);
  // proto:  void QTextDocument::~QTextDocument();
extern void _ZN13QTextDocumentD0Ev(void* qthis);
  // proto:  qreal QTextDocument::documentMargin();
extern void _ZNK13QTextDocument14documentMarginEv(void* qthis);
  // proto:  void QTextDocument::setPageSize(const QSizeF & size);
extern void _ZN13QTextDocument11setPageSizeERK6QSizeF(void* qthis, void* arg0);
  // proto:  void QTextDocument::setHtml(const QString & html);
extern void _ZN13QTextDocument7setHtmlERK7QString(void* qthis, void* arg0);
  // proto:  QTextBlock QTextDocument::end();
extern void _ZNK13QTextDocument3endEv(void* qthis);
  // proto:  int QTextDocument::maximumBlockCount();
extern void _ZNK13QTextDocument17maximumBlockCountEv(void* qthis);
  // proto:  void QTextDocument::setPlainText(const QString & text);
extern void _ZN13QTextDocument12setPlainTextERK7QString(void* qthis, void* arg0);
  // proto:  void QTextDocument::clear();
extern void _ZN13QTextDocument5clearEv(void* qthis);
  // proto:  QVariant QTextDocument::resource(int type, const QUrl & name);
extern void _ZNK13QTextDocument8resourceEiRK4QUrl(void* qthis, int arg0, void* arg1);
  // proto:  QTextBlock QTextDocument::begin();
extern void _ZNK13QTextDocument5beginEv(void* qthis);
  // proto:  void QTextDocument::setMaximumBlockCount(int maximum);
extern void _ZN13QTextDocument20setMaximumBlockCountEi(void* qthis, int arg0);
  // proto:  void QTextDocument::setModified(bool m);
extern void _ZN13QTextDocument11setModifiedEb(void* qthis, bool arg0);
  // proto:  void QAbstractUndoItem::undo();
extern void _ZN17QAbstractUndoItem4undoEv(void* qthis);
  // proto:  void QAbstractUndoItem::redo();
extern void _ZN17QAbstractUndoItem4redoEv(void* qthis);
  // proto:  void QAbstractUndoItem::~QAbstractUndoItem();
extern void _ZN17QAbstractUndoItemD0Ev(void* qthis);
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

// class sizeof(QTextDocument)=1
type QTextDocument struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _cursorPositionChanged QTextDocument_cursorPositionChanged_signal;
//  _documentLayoutChanged QTextDocument_documentLayoutChanged_signal;
//  _undoCommandAdded QTextDocument_undoCommandAdded_signal;
//  _modificationChanged QTextDocument_modificationChanged_signal;
//  _redoAvailable QTextDocument_redoAvailable_signal;
//  _contentsChanged QTextDocument_contentsChanged_signal;
//  _baseUrlChanged QTextDocument_baseUrlChanged_signal;
//  _blockCountChanged QTextDocument_blockCountChanged_signal;
//  _undoAvailable QTextDocument_undoAvailable_signal;
//  _contentsChange QTextDocument_contentsChange_signal;
}

// class sizeof(QAbstractUndoItem)=8
type QAbstractUndoItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QTextDocument::setDefaultStyleSheet(const QString & sheet);
func (this *QTextDocument) setDefaultStyleSheet(args ...interface{}) () {
  // setDefaultStyleSheet(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument20setDefaultStyleSheetERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setDefaultStyleSheet", args)
  }

}

  // proto:  QAbstractTextDocumentLayout * QTextDocument::documentLayout();
func (this *QTextDocument) documentLayout(args ...interface{}) () {
  // documentLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument14documentLayoutEv
  default:
    qtrt.ErrorResolve("QTextDocument", "documentLayout", args)
  }

}

  // proto:  bool QTextDocument::isModified();
func (this *QTextDocument) isModified(args ...interface{}) () {
  // isModified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10isModifiedEv
  default:
    qtrt.ErrorResolve("QTextDocument", "isModified", args)
  }

}

  // proto:  int QTextDocument::revision();
func (this *QTextDocument) revision(args ...interface{}) () {
  // revision()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument8revisionEv
  default:
    qtrt.ErrorResolve("QTextDocument", "revision", args)
  }

}

  // proto:  QSizeF QTextDocument::pageSize();
func (this *QTextDocument) pageSize(args ...interface{}) () {
  // pageSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument8pageSizeEv
  default:
    qtrt.ErrorResolve("QTextDocument", "pageSize", args)
  }

}

  // proto:  void QTextDocument::redo(QTextCursor * cursor);
func (this *QTextDocument) redo(args ...interface{}) () {
  // redo(class QTextCursor *)
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "QTextCursor *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument4redoEP11QTextCursor
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN13QTextDocument4redoEv
  default:
    qtrt.ErrorResolve("QTextDocument", "redo", args)
  }

}

  // proto:  int QTextDocument::lineCount();
func (this *QTextDocument) lineCount(args ...interface{}) () {
  // lineCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9lineCountEv
  default:
    qtrt.ErrorResolve("QTextDocument", "lineCount", args)
  }

}

  // proto:  void QTextDocument::print(QPagedPaintDevice * printer);
func (this *QTextDocument) print(args ...interface{}) () {
  // print(class QPagedPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPagedPaintDevice{}) // "QPagedPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument5printEP17QPagedPaintDevice
    var arg0 = args[0].(QPagedPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "print", args)
  }

}

  // proto:  QTextDocumentPrivate * QTextDocument::docHandle();
func (this *QTextDocument) docHandle(args ...interface{}) () {
  // docHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9docHandleEv
  default:
    qtrt.ErrorResolve("QTextDocument", "docHandle", args)
  }

}

  // proto:  QString QTextDocument::toHtml(const QByteArray & encoding);
func (this *QTextDocument) toHtml(args ...interface{}) () {
  // toHtml(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument6toHtmlERK10QByteArray
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "toHtml", args)
  }

}

  // proto:  int QTextDocument::availableUndoSteps();
func (this *QTextDocument) availableUndoSteps(args ...interface{}) () {
  // availableUndoSteps()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument18availableUndoStepsEv
  default:
    qtrt.ErrorResolve("QTextDocument", "availableUndoSteps", args)
  }

}

  // proto:  void QTextDocument::setUndoRedoEnabled(bool enable);
func (this *QTextDocument) setUndoRedoEnabled(args ...interface{}) () {
  // setUndoRedoEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument18setUndoRedoEnabledEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setUndoRedoEnabled", args)
  }

}

  // proto:  void QTextDocument::undo(QTextCursor * cursor);
func (this *QTextDocument) undo(args ...interface{}) () {
  // undo(class QTextCursor *)
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "QTextCursor *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument4undoEP11QTextCursor
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN13QTextDocument4undoEv
  default:
    qtrt.ErrorResolve("QTextDocument", "undo", args)
  }

}

  // proto:  QString QTextDocument::toPlainText();
func (this *QTextDocument) toPlainText(args ...interface{}) () {
  // toPlainText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument11toPlainTextEv
  default:
    qtrt.ErrorResolve("QTextDocument", "toPlainText", args)
  }

}

  // proto:  void QTextDocument::addResource(int type, const QUrl & name, const QVariant & resource);
func (this *QTextDocument) addResource(args ...interface{}) () {
  // addResource(int, const class QUrl &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QUrl{}) // "const QUrl &"
  vtys[0][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUrl).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QVariant).qclsinst
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QTextDocument", "addResource", args)
  }

}

  // proto:  QSizeF QTextDocument::size();
func (this *QTextDocument) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument4sizeEv
  default:
    qtrt.ErrorResolve("QTextDocument", "size", args)
  }

}

  // proto:  QTextObject * QTextDocument::object(int objectIndex);
func (this *QTextDocument) object(args ...interface{}) () {
  // object(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument6objectEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "object", args)
  }

}

  // proto:  QTextDocument * QTextDocument::clone(QObject * parent);
func (this *QTextDocument) clone(args ...interface{}) () {
  // clone(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument5cloneEP7QObject
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "clone", args)
  }

}

  // proto:  void QTextDocument::markContentsDirty(int from, int length);
func (this *QTextDocument) markContentsDirty(args ...interface{}) () {
  // markContentsDirty(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument17markContentsDirtyEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTextDocument", "markContentsDirty", args)
  }

}

  // proto:  void QTextDocument::QTextDocument(QObject * parent);
func NewQTextDocument(args ...interface{}) QTextDocument {
  return QTextDocument{}
}

  // proto:  int QTextDocument::characterCount();
func (this *QTextDocument) characterCount(args ...interface{}) () {
  // characterCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument14characterCountEv
  default:
    qtrt.ErrorResolve("QTextDocument", "characterCount", args)
  }

}

  // proto:  QTextFrame * QTextDocument::rootFrame();
func (this *QTextDocument) rootFrame(args ...interface{}) () {
  // rootFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9rootFrameEv
  default:
    qtrt.ErrorResolve("QTextDocument", "rootFrame", args)
  }

}

  // proto:  QTextBlock QTextDocument::firstBlock();
func (this *QTextDocument) firstBlock(args ...interface{}) () {
  // firstBlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10firstBlockEv
  default:
    qtrt.ErrorResolve("QTextDocument", "firstBlock", args)
  }

}

  // proto:  int QTextDocument::blockCount();
func (this *QTextDocument) blockCount(args ...interface{}) () {
  // blockCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10blockCountEv
  default:
    qtrt.ErrorResolve("QTextDocument", "blockCount", args)
  }

}

  // proto:  qreal QTextDocument::idealWidth();
func (this *QTextDocument) idealWidth(args ...interface{}) () {
  // idealWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10idealWidthEv
  default:
    qtrt.ErrorResolve("QTextDocument", "idealWidth", args)
  }

}

  // proto:  void QTextDocument::adjustSize();
func (this *QTextDocument) adjustSize(args ...interface{}) () {
  // adjustSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument10adjustSizeEv
  default:
    qtrt.ErrorResolve("QTextDocument", "adjustSize", args)
  }

}

  // proto:  bool QTextDocument::isRedoAvailable();
func (this *QTextDocument) isRedoAvailable(args ...interface{}) () {
  // isRedoAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument15isRedoAvailableEv
  default:
    qtrt.ErrorResolve("QTextDocument", "isRedoAvailable", args)
  }

}

  // proto:  QVector<QTextFormat> QTextDocument::allFormats();
func (this *QTextDocument) allFormats(args ...interface{}) () {
  // allFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10allFormatsEv
  default:
    qtrt.ErrorResolve("QTextDocument", "allFormats", args)
  }

}

  // proto:  QString QTextDocument::defaultStyleSheet();
func (this *QTextDocument) defaultStyleSheet(args ...interface{}) () {
  // defaultStyleSheet()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument17defaultStyleSheetEv
  default:
    qtrt.ErrorResolve("QTextDocument", "defaultStyleSheet", args)
  }

}

  // proto:  QTextBlock QTextDocument::lastBlock();
func (this *QTextDocument) lastBlock(args ...interface{}) () {
  // lastBlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9lastBlockEv
  default:
    qtrt.ErrorResolve("QTextDocument", "lastBlock", args)
  }

}

  // proto:  bool QTextDocument::useDesignMetrics();
func (this *QTextDocument) useDesignMetrics(args ...interface{}) () {
  // useDesignMetrics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument16useDesignMetricsEv
  default:
    qtrt.ErrorResolve("QTextDocument", "useDesignMetrics", args)
  }

}

  // proto:  int QTextDocument::pageCount();
func (this *QTextDocument) pageCount(args ...interface{}) () {
  // pageCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9pageCountEv
  default:
    qtrt.ErrorResolve("QTextDocument", "pageCount", args)
  }

}

  // proto:  void QTextDocument::setTextWidth(qreal width);
func (this *QTextDocument) setTextWidth(args ...interface{}) () {
  // setTextWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument12setTextWidthEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setTextWidth", args)
  }

}

  // proto:  void QTextDocument::setDocumentMargin(qreal margin);
func (this *QTextDocument) setDocumentMargin(args ...interface{}) () {
  // setDocumentMargin(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument17setDocumentMarginEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setDocumentMargin", args)
  }

}

  // proto:  bool QTextDocument::isUndoAvailable();
func (this *QTextDocument) isUndoAvailable(args ...interface{}) () {
  // isUndoAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument15isUndoAvailableEv
  default:
    qtrt.ErrorResolve("QTextDocument", "isUndoAvailable", args)
  }

}

  // proto:  qreal QTextDocument::indentWidth();
func (this *QTextDocument) indentWidth(args ...interface{}) () {
  // indentWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument11indentWidthEv
  default:
    qtrt.ErrorResolve("QTextDocument", "indentWidth", args)
  }

}

  // proto:  void QTextDocument::setUseDesignMetrics(bool b);
func (this *QTextDocument) setUseDesignMetrics(args ...interface{}) () {
  // setUseDesignMetrics(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument19setUseDesignMetricsEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setUseDesignMetrics", args)
  }

}

  // proto:  void QTextDocument::setDocumentLayout(QAbstractTextDocumentLayout * layout);
func (this *QTextDocument) setDocumentLayout(args ...interface{}) () {
  // setDocumentLayout(class QAbstractTextDocumentLayout *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractTextDocumentLayout{}) // "QAbstractTextDocumentLayout *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout
    var arg0 = args[0].(QAbstractTextDocumentLayout).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setDocumentLayout", args)
  }

}

  // proto:  void QTextDocument::setIndentWidth(qreal width);
func (this *QTextDocument) setIndentWidth(args ...interface{}) () {
  // setIndentWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument14setIndentWidthEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setIndentWidth", args)
  }

}

  // proto:  QUrl QTextDocument::baseUrl();
func (this *QTextDocument) baseUrl(args ...interface{}) () {
  // baseUrl()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument7baseUrlEv
  default:
    qtrt.ErrorResolve("QTextDocument", "baseUrl", args)
  }

}

  // proto:  QTextFrame * QTextDocument::frameAt(int pos);
func (this *QTextDocument) frameAt(args ...interface{}) () {
  // frameAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument7frameAtEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "frameAt", args)
  }

}

  // proto:  void QTextDocument::setDefaultTextOption(const QTextOption & option);
func (this *QTextDocument) setDefaultTextOption(args ...interface{}) () {
  // setDefaultTextOption(const class QTextOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument20setDefaultTextOptionERK11QTextOption
    var arg0 = args[0].(QTextOption).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setDefaultTextOption", args)
  }

}

  // proto:  QFont QTextDocument::defaultFont();
func (this *QTextDocument) defaultFont(args ...interface{}) () {
  // defaultFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument11defaultFontEv
  default:
    qtrt.ErrorResolve("QTextDocument", "defaultFont", args)
  }

}

  // proto:  QTextBlock QTextDocument::findBlockByNumber(int blockNumber);
func (this *QTextDocument) findBlockByNumber(args ...interface{}) () {
  // findBlockByNumber(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument17findBlockByNumberEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "findBlockByNumber", args)
  }

}

  // proto:  QTextOption QTextDocument::defaultTextOption();
func (this *QTextDocument) defaultTextOption(args ...interface{}) () {
  // defaultTextOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument17defaultTextOptionEv
  default:
    qtrt.ErrorResolve("QTextDocument", "defaultTextOption", args)
  }

}

  // proto:  QTextBlock QTextDocument::findBlock(int pos);
func (this *QTextDocument) findBlock(args ...interface{}) () {
  // findBlock(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9findBlockEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "findBlock", args)
  }

}

  // proto:  void QTextDocument::setBaseUrl(const QUrl & url);
func (this *QTextDocument) setBaseUrl(args ...interface{}) () {
  // setBaseUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument10setBaseUrlERK4QUrl
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setBaseUrl", args)
  }

}

  // proto:  void QTextDocument::appendUndoItem(QAbstractUndoItem * );
func (this *QTextDocument) appendUndoItem(args ...interface{}) () {
  // appendUndoItem(class QAbstractUndoItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractUndoItem{}) // "QAbstractUndoItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem
    var arg0 = args[0].(QAbstractUndoItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "appendUndoItem", args)
  }

}

  // proto:  void QTextDocument::drawContents(QPainter * painter, const QRectF & rect);
func (this *QTextDocument) drawContents(args ...interface{}) () {
  // drawContents(class QPainter *, const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRectF).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTextDocument", "drawContents", args)
  }

}

  // proto:  QTextBlock QTextDocument::findBlockByLineNumber(int blockNumber);
func (this *QTextDocument) findBlockByLineNumber(args ...interface{}) () {
  // findBlockByLineNumber(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument21findBlockByLineNumberEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "findBlockByLineNumber", args)
  }

}

  // proto:  qreal QTextDocument::textWidth();
func (this *QTextDocument) textWidth(args ...interface{}) () {
  // textWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument9textWidthEv
  default:
    qtrt.ErrorResolve("QTextDocument", "textWidth", args)
  }

}

  // proto:  const QMetaObject * QTextDocument::metaObject();
func (this *QTextDocument) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument10metaObjectEv
  default:
    qtrt.ErrorResolve("QTextDocument", "metaObject", args)
  }

}

  // proto:  int QTextDocument::availableRedoSteps();
func (this *QTextDocument) availableRedoSteps(args ...interface{}) () {
  // availableRedoSteps()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument18availableRedoStepsEv
  default:
    qtrt.ErrorResolve("QTextDocument", "availableRedoSteps", args)
  }

}

  // proto:  QChar QTextDocument::characterAt(int pos);
func (this *QTextDocument) characterAt(args ...interface{}) () {
  // characterAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument11characterAtEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "characterAt", args)
  }

}

  // proto:  void QTextDocument::setDefaultFont(const QFont & font);
func (this *QTextDocument) setDefaultFont(args ...interface{}) () {
  // setDefaultFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument14setDefaultFontERK5QFont
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setDefaultFont", args)
  }

}

  // proto:  QTextObject * QTextDocument::objectForFormat(const QTextFormat & );
func (this *QTextDocument) objectForFormat(args ...interface{}) () {
  // objectForFormat(const class QTextFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFormat{}) // "const QTextFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument15objectForFormatERK11QTextFormat
    var arg0 = args[0].(QTextFormat).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "objectForFormat", args)
  }

}

  // proto:  bool QTextDocument::isEmpty();
func (this *QTextDocument) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument7isEmptyEv
  default:
    qtrt.ErrorResolve("QTextDocument", "isEmpty", args)
  }

}

  // proto:  bool QTextDocument::isUndoRedoEnabled();
func (this *QTextDocument) isUndoRedoEnabled(args ...interface{}) () {
  // isUndoRedoEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument17isUndoRedoEnabledEv
  default:
    qtrt.ErrorResolve("QTextDocument", "isUndoRedoEnabled", args)
  }

}

  // proto:  void QTextDocument::~QTextDocument();
func (this *QTextDocument) FreeQTextDocument(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextDocument", "~QTextDocument", args)
  }

}

  // proto:  qreal QTextDocument::documentMargin();
func (this *QTextDocument) documentMargin(args ...interface{}) () {
  // documentMargin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument14documentMarginEv
  default:
    qtrt.ErrorResolve("QTextDocument", "documentMargin", args)
  }

}

  // proto:  void QTextDocument::setPageSize(const QSizeF & size);
func (this *QTextDocument) setPageSize(args ...interface{}) () {
  // setPageSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument11setPageSizeERK6QSizeF
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setPageSize", args)
  }

}

  // proto:  void QTextDocument::setHtml(const QString & html);
func (this *QTextDocument) setHtml(args ...interface{}) () {
  // setHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument7setHtmlERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setHtml", args)
  }

}

  // proto:  QTextBlock QTextDocument::end();
func (this *QTextDocument) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument3endEv
  default:
    qtrt.ErrorResolve("QTextDocument", "end", args)
  }

}

  // proto:  int QTextDocument::maximumBlockCount();
func (this *QTextDocument) maximumBlockCount(args ...interface{}) () {
  // maximumBlockCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument17maximumBlockCountEv
  default:
    qtrt.ErrorResolve("QTextDocument", "maximumBlockCount", args)
  }

}

  // proto:  void QTextDocument::setPlainText(const QString & text);
func (this *QTextDocument) setPlainText(args ...interface{}) () {
  // setPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument12setPlainTextERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setPlainText", args)
  }

}

  // proto:  void QTextDocument::clear();
func (this *QTextDocument) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument5clearEv
  default:
    qtrt.ErrorResolve("QTextDocument", "clear", args)
  }

}

  // proto:  QVariant QTextDocument::resource(int type, const QUrl & name);
func (this *QTextDocument) resource(args ...interface{}) () {
  // resource(int, const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument8resourceEiRK4QUrl
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUrl).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTextDocument", "resource", args)
  }

}

  // proto:  QTextBlock QTextDocument::begin();
func (this *QTextDocument) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument5beginEv
  default:
    qtrt.ErrorResolve("QTextDocument", "begin", args)
  }

}

  // proto:  void QTextDocument::setMaximumBlockCount(int maximum);
func (this *QTextDocument) setMaximumBlockCount(args ...interface{}) () {
  // setMaximumBlockCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument20setMaximumBlockCountEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setMaximumBlockCount", args)
  }

}

  // proto:  void QTextDocument::setModified(bool m);
func (this *QTextDocument) setModified(args ...interface{}) () {
  // setModified(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument11setModifiedEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextDocument", "setModified", args)
  }

}

  // proto:  void QAbstractUndoItem::undo();
func (this *QAbstractUndoItem) undo(args ...interface{}) () {
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractUndoItem4undoEv
  default:
    qtrt.ErrorResolve("QAbstractUndoItem", "undo", args)
  }

}

  // proto:  void QAbstractUndoItem::redo();
func (this *QAbstractUndoItem) redo(args ...interface{}) () {
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QAbstractUndoItem4redoEv
  default:
    qtrt.ErrorResolve("QAbstractUndoItem", "redo", args)
  }

}

  // proto:  void QAbstractUndoItem::~QAbstractUndoItem();
func (this *QAbstractUndoItem) FreeQAbstractUndoItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractUndoItem", "~QAbstractUndoItem", args)
  }

}

// <= body block end

