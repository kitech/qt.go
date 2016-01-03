package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
extern void _ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant(void* qthis, int32_t arg0, void* arg1, void* arg2);
  // proto:  QSizeF QTextDocument::size();
extern void _ZNK13QTextDocument4sizeEv(void* qthis);
  // proto:  QTextObject * QTextDocument::object(int objectIndex);
extern void _ZNK13QTextDocument6objectEi(void* qthis, int32_t arg0);
  // proto:  QTextDocument * QTextDocument::clone(QObject * parent);
extern void _ZNK13QTextDocument5cloneEP7QObject(void* qthis, void* arg0);
  // proto:  void QTextDocument::markContentsDirty(int from, int length);
extern void _ZN13QTextDocument17markContentsDirtyEii(void* qthis, int32_t arg0, int32_t arg1);
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
extern void _ZNK13QTextDocument7frameAtEi(void* qthis, int32_t arg0);
  // proto:  void QTextDocument::QTextDocument(const QTextDocument & );
extern void* dector_ZN13QTextDocumentC1ERKS_(void* arg0);
extern void _ZN13QTextDocumentC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTextDocument::setDefaultTextOption(const QTextOption & option);
extern void _ZN13QTextDocument20setDefaultTextOptionERK11QTextOption(void* qthis, void* arg0);
  // proto:  QFont QTextDocument::defaultFont();
extern void _ZNK13QTextDocument11defaultFontEv(void* qthis);
  // proto:  QTextBlock QTextDocument::findBlockByNumber(int blockNumber);
extern void _ZNK13QTextDocument17findBlockByNumberEi(void* qthis, int32_t arg0);
  // proto:  QTextOption QTextDocument::defaultTextOption();
extern void _ZNK13QTextDocument17defaultTextOptionEv(void* qthis);
  // proto:  QTextBlock QTextDocument::findBlock(int pos);
extern void _ZNK13QTextDocument9findBlockEi(void* qthis, int32_t arg0);
  // proto:  void QTextDocument::setBaseUrl(const QUrl & url);
extern void _ZN13QTextDocument10setBaseUrlERK4QUrl(void* qthis, void* arg0);
  // proto:  void QTextDocument::appendUndoItem(QAbstractUndoItem * );
extern void _ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem(void* qthis, void* arg0);
  // proto:  void QTextDocument::redo();
extern void _ZN13QTextDocument4redoEv(void* qthis);
  // proto:  void QTextDocument::drawContents(QPainter * painter, const QRectF & rect);
extern void _ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF(void* qthis, void* arg0, void* arg1);
  // proto:  QTextBlock QTextDocument::findBlockByLineNumber(int blockNumber);
extern void _ZNK13QTextDocument21findBlockByLineNumberEi(void* qthis, int32_t arg0);
  // proto:  void QTextDocument::undo();
extern void _ZN13QTextDocument4undoEv(void* qthis);
  // proto:  qreal QTextDocument::textWidth();
extern void _ZNK13QTextDocument9textWidthEv(void* qthis);
  // proto:  const QMetaObject * QTextDocument::metaObject();
extern void _ZNK13QTextDocument10metaObjectEv(void* qthis);
  // proto:  int QTextDocument::availableRedoSteps();
extern void _ZNK13QTextDocument18availableRedoStepsEv(void* qthis);
  // proto:  QChar QTextDocument::characterAt(int pos);
extern void _ZNK13QTextDocument11characterAtEi(void* qthis, int32_t arg0);
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
extern void _ZNK13QTextDocument8resourceEiRK4QUrl(void* qthis, int32_t arg0, void* arg1);
  // proto:  QTextBlock QTextDocument::begin();
extern void _ZNK13QTextDocument5beginEv(void* qthis);
  // proto:  void QTextDocument::setMaximumBlockCount(int maximum);
extern void _ZN13QTextDocument20setMaximumBlockCountEi(void* qthis, int32_t arg0);
  // proto:  void QTextDocument::setModified(bool m);
extern void _ZN13QTextDocument11setModifiedEb(void* qthis, bool arg0);
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
  qclsinst unsafe.Pointer /* *C.void */;
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
  qclsinst unsafe.Pointer /* *C.void */;
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
    // invoke: void setDefaultStyleSheet(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument20setDefaultStyleSheetERK7QString(this.qclsinst, arg0)
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
    // invoke: QAbstractTextDocumentLayout * documentLayout()
    C._ZNK13QTextDocument14documentLayoutEv(this.qclsinst)
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
    // invoke: bool isModified()
    C._ZNK13QTextDocument10isModifiedEv(this.qclsinst)
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
    // invoke: int revision()
    C._ZNK13QTextDocument8revisionEv(this.qclsinst)
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
    // invoke: QSizeF pageSize()
    C._ZNK13QTextDocument8pageSizeEv(this.qclsinst)
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
    // invoke: void redo(class QTextCursor *)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument4redoEP11QTextCursor(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN13QTextDocument4redoEv
    // invoke: void redo()
    C._ZN13QTextDocument4redoEv(this.qclsinst)
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
    // invoke: int lineCount()
    C._ZNK13QTextDocument9lineCountEv(this.qclsinst)
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
    // invoke: void print(class QPagedPaintDevice *)
    var arg0 = args[0].(QPagedPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QTextDocument5printEP17QPagedPaintDevice(this.qclsinst, arg0)
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
    // invoke: QTextDocumentPrivate * docHandle()
    C._ZNK13QTextDocument9docHandleEv(this.qclsinst)
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
    // invoke: QString toHtml(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QTextDocument6toHtmlERK10QByteArray(this.qclsinst, arg0)
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
    // invoke: int availableUndoSteps()
    C._ZNK13QTextDocument18availableUndoStepsEv(this.qclsinst)
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
    // invoke: void setUndoRedoEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument18setUndoRedoEnabledEb(this.qclsinst, arg0)
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
    // invoke: void undo(class QTextCursor *)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument4undoEP11QTextCursor(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN13QTextDocument4undoEv
    // invoke: void undo()
    C._ZN13QTextDocument4undoEv(this.qclsinst)
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
    // invoke: QString toPlainText()
    C._ZNK13QTextDocument11toPlainTextEv(this.qclsinst)
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
    // invoke: void addResource(int, const class QUrl &, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUrl).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QVariant).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant(this.qclsinst, arg0, arg1, arg2)
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
    // invoke: QSizeF size()
    C._ZNK13QTextDocument4sizeEv(this.qclsinst)
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
    // invoke: QTextObject * object(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QTextDocument6objectEi(this.qclsinst, arg0)
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
    // invoke: QTextDocument * clone(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QTextDocument5cloneEP7QObject(this.qclsinst, arg0)
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
    // invoke: void markContentsDirty(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN13QTextDocument17markContentsDirtyEii(this.qclsinst, arg0, arg1)
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
    // invoke: int characterCount()
    C._ZNK13QTextDocument14characterCountEv(this.qclsinst)
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
    // invoke: QTextFrame * rootFrame()
    C._ZNK13QTextDocument9rootFrameEv(this.qclsinst)
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
    // invoke: QTextBlock firstBlock()
    C._ZNK13QTextDocument10firstBlockEv(this.qclsinst)
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
    // invoke: int blockCount()
    C._ZNK13QTextDocument10blockCountEv(this.qclsinst)
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
    // invoke: qreal idealWidth()
    C._ZNK13QTextDocument10idealWidthEv(this.qclsinst)
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
    // invoke: void adjustSize()
    C._ZN13QTextDocument10adjustSizeEv(this.qclsinst)
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
    // invoke: bool isRedoAvailable()
    C._ZNK13QTextDocument15isRedoAvailableEv(this.qclsinst)
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
    // invoke: QVector<QTextFormat> allFormats()
    C._ZNK13QTextDocument10allFormatsEv(this.qclsinst)
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
    // invoke: QString defaultStyleSheet()
    C._ZNK13QTextDocument17defaultStyleSheetEv(this.qclsinst)
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
    // invoke: QTextBlock lastBlock()
    C._ZNK13QTextDocument9lastBlockEv(this.qclsinst)
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
    // invoke: bool useDesignMetrics()
    C._ZNK13QTextDocument16useDesignMetricsEv(this.qclsinst)
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
    // invoke: int pageCount()
    C._ZNK13QTextDocument9pageCountEv(this.qclsinst)
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
    // invoke: void setTextWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument12setTextWidthEd(this.qclsinst, arg0)
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
    // invoke: void setDocumentMargin(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument17setDocumentMarginEd(this.qclsinst, arg0)
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
    // invoke: bool isUndoAvailable()
    C._ZNK13QTextDocument15isUndoAvailableEv(this.qclsinst)
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
    // invoke: qreal indentWidth()
    C._ZNK13QTextDocument11indentWidthEv(this.qclsinst)
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
    // invoke: void setUseDesignMetrics(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument19setUseDesignMetricsEb(this.qclsinst, arg0)
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
    // invoke: void setDocumentLayout(class QAbstractTextDocumentLayout *)
    var arg0 = args[0].(QAbstractTextDocumentLayout).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout(this.qclsinst, arg0)
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
    // invoke: void setIndentWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument14setIndentWidthEd(this.qclsinst, arg0)
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
    // invoke: QUrl baseUrl()
    C._ZNK13QTextDocument7baseUrlEv(this.qclsinst)
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
    // invoke: QTextFrame * frameAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QTextDocument7frameAtEi(this.qclsinst, arg0)
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
    // invoke: void setDefaultTextOption(const class QTextOption &)
    var arg0 = args[0].(QTextOption).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument20setDefaultTextOptionERK11QTextOption(this.qclsinst, arg0)
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
    // invoke: QFont defaultFont()
    C._ZNK13QTextDocument11defaultFontEv(this.qclsinst)
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
    // invoke: QTextBlock findBlockByNumber(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QTextDocument17findBlockByNumberEi(this.qclsinst, arg0)
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
    // invoke: QTextOption defaultTextOption()
    C._ZNK13QTextDocument17defaultTextOptionEv(this.qclsinst)
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
    // invoke: QTextBlock findBlock(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QTextDocument9findBlockEi(this.qclsinst, arg0)
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
    // invoke: void setBaseUrl(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument10setBaseUrlERK4QUrl(this.qclsinst, arg0)
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
    // invoke: void appendUndoItem(class QAbstractUndoItem *)
    var arg0 = args[0].(QAbstractUndoItem).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem(this.qclsinst, arg0)
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
    // invoke: void drawContents(class QPainter *, const class QRectF &)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRectF).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF(this.qclsinst, arg0, arg1)
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
    // invoke: QTextBlock findBlockByLineNumber(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QTextDocument21findBlockByLineNumberEi(this.qclsinst, arg0)
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
    // invoke: qreal textWidth()
    C._ZNK13QTextDocument9textWidthEv(this.qclsinst)
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QTextDocument10metaObjectEv(this.qclsinst)
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
    // invoke: int availableRedoSteps()
    C._ZNK13QTextDocument18availableRedoStepsEv(this.qclsinst)
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
    // invoke: QChar characterAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QTextDocument11characterAtEi(this.qclsinst, arg0)
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
    // invoke: void setDefaultFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument14setDefaultFontERK5QFont(this.qclsinst, arg0)
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
    // invoke: QTextObject * objectForFormat(const class QTextFormat &)
    var arg0 = args[0].(QTextFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QTextDocument15objectForFormatERK11QTextFormat(this.qclsinst, arg0)
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
    // invoke: bool isEmpty()
    C._ZNK13QTextDocument7isEmptyEv(this.qclsinst)
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
    // invoke: bool isUndoRedoEnabled()
    C._ZNK13QTextDocument17isUndoRedoEnabledEv(this.qclsinst)
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
    // invoke: qreal documentMargin()
    C._ZNK13QTextDocument14documentMarginEv(this.qclsinst)
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
    // invoke: void setPageSize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument11setPageSizeERK6QSizeF(this.qclsinst, arg0)
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
    // invoke: void setHtml(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument7setHtmlERK7QString(this.qclsinst, arg0)
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
    // invoke: QTextBlock end()
    C._ZNK13QTextDocument3endEv(this.qclsinst)
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
    // invoke: int maximumBlockCount()
    C._ZNK13QTextDocument17maximumBlockCountEv(this.qclsinst)
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
    // invoke: void setPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument12setPlainTextERK7QString(this.qclsinst, arg0)
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
    // invoke: void clear()
    C._ZN13QTextDocument5clearEv(this.qclsinst)
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
    // invoke: QVariant resource(int, const class QUrl &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUrl).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK13QTextDocument8resourceEiRK4QUrl(this.qclsinst, arg0, arg1)
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
    // invoke: QTextBlock begin()
    C._ZNK13QTextDocument5beginEv(this.qclsinst)
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
    // invoke: void setMaximumBlockCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument20setMaximumBlockCountEi(this.qclsinst, arg0)
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
    // invoke: void setModified(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN13QTextDocument11setModifiedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setModified", args)
  }

}

// <= body block end

