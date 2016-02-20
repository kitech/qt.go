package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QSizeF QTextDocument::pageSize();
extern void* C_ZNK13QTextDocument8pageSizeEv(void* qthis); // 4
  // proto:  int QTextDocument::availableRedoSteps();
extern int32_t C_ZNK13QTextDocument18availableRedoStepsEv(void* qthis); // 4
  // proto:  QString QTextDocument::defaultStyleSheet();
extern void* C_ZNK13QTextDocument17defaultStyleSheetEv(void* qthis); // 4
  // proto:  QVariant QTextDocument::resource(int type, const QUrl & name);
extern void* C_ZNK13QTextDocument8resourceEiRK4QUrl(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  qreal QTextDocument::indentWidth();
extern double C_ZNK13QTextDocument11indentWidthEv(void* qthis); // 4
  // proto:  const QMetaObject * QTextDocument::metaObject();
extern void C_ZNK13QTextDocument10metaObjectEv(void* qthis); // 4
  // proto:  QTextBlock QTextDocument::lastBlock();
extern void* C_ZNK13QTextDocument9lastBlockEv(void* qthis); // 4
  // proto:  void QTextDocument::setUndoRedoEnabled(bool enable);
extern void C_ZN13QTextDocument18setUndoRedoEnabledEb(void* qthis, bool arg0); // 4
  // proto:  bool QTextDocument::isUndoAvailable();
extern bool C_ZNK13QTextDocument15isUndoAvailableEv(void* qthis); // 4
  // proto:  void QTextDocument::setDefaultFont(const QFont & font);
extern void C_ZN13QTextDocument14setDefaultFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  bool QTextDocument::useDesignMetrics();
extern bool C_ZNK13QTextDocument16useDesignMetricsEv(void* qthis); // 4
  // proto:  qreal QTextDocument::textWidth();
extern double C_ZNK13QTextDocument9textWidthEv(void* qthis); // 4
  // proto:  int QTextDocument::characterCount();
extern int32_t C_ZNK13QTextDocument14characterCountEv(void* qthis); // 4
  // proto:  bool QTextDocument::isUndoRedoEnabled();
extern bool C_ZNK13QTextDocument17isUndoRedoEnabledEv(void* qthis); // 4
  // proto:  void QTextDocument::print(QPagedPaintDevice * printer);
extern void C_ZNK13QTextDocument5printEP17QPagedPaintDevice(void* qthis, void* arg0); // 4
  // proto:  QTextObject * QTextDocument::objectForFormat(const QTextFormat & );
extern void* C_ZNK13QTextDocument15objectForFormatERK11QTextFormat(void* qthis, void* arg0); // 4
  // proto:  QString QTextDocument::toPlainText();
extern void* C_ZNK13QTextDocument11toPlainTextEv(void* qthis); // 4
  // proto:  void QTextDocument::undo();
extern void C_ZN13QTextDocument4undoEv(void* qthis); // 4
  // proto:  void QTextDocument::undo(QTextCursor * cursor);
extern void C_ZN13QTextDocument4undoEP11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  void QTextDocument::drawContents(QPainter * painter, const QRectF & rect);
extern void C_ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QTextDocument::setModified(bool m);
extern void C_ZN13QTextDocument11setModifiedEb(void* qthis, bool arg0); // 4
  // proto:  int QTextDocument::blockCount();
extern int32_t C_ZNK13QTextDocument10blockCountEv(void* qthis); // 4
  // proto:  QFont QTextDocument::defaultFont();
extern void* C_ZNK13QTextDocument11defaultFontEv(void* qthis); // 4
  // proto:  bool QTextDocument::isRedoAvailable();
extern bool C_ZNK13QTextDocument15isRedoAvailableEv(void* qthis); // 4
  // proto:  QTextFrame * QTextDocument::frameAt(int pos);
extern void* C_ZNK13QTextDocument7frameAtEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTextDocument::pageCount();
extern int32_t C_ZNK13QTextDocument9pageCountEv(void* qthis); // 4
  // proto:  void QTextDocument::QTextDocument(QObject * parent);
extern void* C_ZN13QTextDocumentC2EP7QObject(void* arg0); // 3
  // proto:  void QTextDocument::QTextDocument(const QString & text, QObject * parent);
extern void* C_ZN13QTextDocumentC2ERK7QStringP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QTextDocument::~QTextDocument();
extern void C_ZN13QTextDocumentD2Ev(void* qthis); // 4
  // proto:  int QTextDocument::availableUndoSteps();
extern int32_t C_ZNK13QTextDocument18availableUndoStepsEv(void* qthis); // 4
  // proto:  int QTextDocument::maximumBlockCount();
extern int32_t C_ZNK13QTextDocument17maximumBlockCountEv(void* qthis); // 4
  // proto:  int QTextDocument::lineCount();
extern int32_t C_ZNK13QTextDocument9lineCountEv(void* qthis); // 4
  // proto:  QTextBlock QTextDocument::findBlock(int pos);
extern void* C_ZNK13QTextDocument9findBlockEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextDocument::redo(QTextCursor * cursor);
extern void C_ZN13QTextDocument4redoEP11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  void QTextDocument::redo();
extern void C_ZN13QTextDocument4redoEv(void* qthis); // 4
  // proto:  QSizeF QTextDocument::size();
extern void* C_ZNK13QTextDocument4sizeEv(void* qthis); // 4
  // proto:  bool QTextDocument::isModified();
extern bool C_ZNK13QTextDocument10isModifiedEv(void* qthis); // 4
  // proto:  QTextBlock QTextDocument::findBlockByLineNumber(int blockNumber);
extern void* C_ZNK13QTextDocument21findBlockByLineNumberEi(void* qthis, int32_t arg0); // 4
  // proto:  QUrl QTextDocument::baseUrl();
extern void* C_ZNK13QTextDocument7baseUrlEv(void* qthis); // 4
  // proto:  qreal QTextDocument::documentMargin();
extern double C_ZNK13QTextDocument14documentMarginEv(void* qthis); // 4
  // proto:  QChar QTextDocument::characterAt(int pos);
extern void* C_ZNK13QTextDocument11characterAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextDocument::setDocumentMargin(qreal margin);
extern void C_ZN13QTextDocument17setDocumentMarginEd(void* qthis, double arg0); // 4
  // proto:  QTextFrame * QTextDocument::rootFrame();
extern void* C_ZNK13QTextDocument9rootFrameEv(void* qthis); // 4
  // proto:  QAbstractTextDocumentLayout * QTextDocument::documentLayout();
extern void C_ZNK13QTextDocument14documentLayoutEv(void* qthis); // 4
  // proto:  Qt::CursorMoveStyle QTextDocument::defaultCursorMoveStyle();
extern void C_ZNK13QTextDocument22defaultCursorMoveStyleEv(void* qthis); // 4
  // proto:  void QTextDocument::setHtml(const QString & html);
extern void C_ZN13QTextDocument7setHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextDocument::setMaximumBlockCount(int maximum);
extern void C_ZN13QTextDocument20setMaximumBlockCountEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextDocument::addResource(int type, const QUrl & name, const QVariant & resource);
extern void C_ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  void QTextDocument::setDocumentLayout(QAbstractTextDocumentLayout * layout);
extern void C_ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout(void* qthis, void* arg0); // 4
  // proto:  QTextOption QTextDocument::defaultTextOption();
extern void* C_ZNK13QTextDocument17defaultTextOptionEv(void* qthis); // 4
  // proto:  QString QTextDocument::toHtml(const QByteArray & encoding);
extern void* C_ZNK13QTextDocument6toHtmlERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QTextBlock QTextDocument::firstBlock();
extern void* C_ZNK13QTextDocument10firstBlockEv(void* qthis); // 4
  // proto:  QTextBlock QTextDocument::findBlockByNumber(int blockNumber);
extern void* C_ZNK13QTextDocument17findBlockByNumberEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QTextDocument::idealWidth();
extern double C_ZNK13QTextDocument10idealWidthEv(void* qthis); // 4
  // proto:  QTextBlock QTextDocument::end();
extern void* C_ZNK13QTextDocument3endEv(void* qthis); // 4
  // proto:  void QTextDocument::setPageSize(const QSizeF & size);
extern void C_ZN13QTextDocument11setPageSizeERK6QSizeF(void* qthis, void* arg0); // 4
  // proto:  void QTextDocument::setDefaultStyleSheet(const QString & sheet);
extern void C_ZN13QTextDocument20setDefaultStyleSheetERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextDocument::appendUndoItem(QAbstractUndoItem * );
extern void C_ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem(void* qthis, void* arg0); // 4
  // proto:  QTextDocumentPrivate * QTextDocument::docHandle();
extern void C_ZNK13QTextDocument9docHandleEv(void* qthis); // 4
  // proto:  bool QTextDocument::isEmpty();
extern bool C_ZNK13QTextDocument7isEmptyEv(void* qthis); // 4
  // proto:  void QTextDocument::markContentsDirty(int from, int length);
extern void C_ZN13QTextDocument17markContentsDirtyEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QTextDocument::revision();
extern int32_t C_ZNK13QTextDocument8revisionEv(void* qthis); // 4
  // proto:  void QTextDocument::setBaseUrl(const QUrl & url);
extern void C_ZN13QTextDocument10setBaseUrlERK4QUrl(void* qthis, void* arg0); // 4
  // proto:  QTextBlock QTextDocument::begin();
extern void* C_ZNK13QTextDocument5beginEv(void* qthis); // 4
  // proto:  void QTextDocument::setPlainText(const QString & text);
extern void C_ZN13QTextDocument12setPlainTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QTextDocument * QTextDocument::clone(QObject * parent);
extern void* C_ZNK13QTextDocument5cloneEP7QObject(void* qthis, void* arg0); // 4
  // proto:  QTextObject * QTextDocument::object(int objectIndex);
extern void* C_ZNK13QTextDocument6objectEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextDocument::setTextWidth(qreal width);
extern void C_ZN13QTextDocument12setTextWidthEd(void* qthis, double arg0); // 4
  // proto:  void QTextDocument::adjustSize();
extern void C_ZN13QTextDocument10adjustSizeEv(void* qthis); // 4
  // proto:  QVector<QTextFormat> QTextDocument::allFormats();
extern void C_ZNK13QTextDocument10allFormatsEv(void* qthis); // 4
  // proto:  void QTextDocument::setDefaultTextOption(const QTextOption & option);
extern void C_ZN13QTextDocument20setDefaultTextOptionERK11QTextOption(void* qthis, void* arg0); // 4
  // proto:  void QTextDocument::setUseDesignMetrics(bool b);
extern void C_ZN13QTextDocument19setUseDesignMetricsEb(void* qthis, bool arg0); // 4
  // proto:  void QTextDocument::clear();
extern void C_ZN13QTextDocument5clearEv(void* qthis); // 4
  // proto:  void QTextDocument::setIndentWidth(qreal width);
extern void C_ZN13QTextDocument14setIndentWidthEd(void* qthis, double arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTextDocument)=1
type QTextDocument struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// pageSize()
func (this *QTextDocument) Pagesize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument8pageSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSizeF{}) // "QSizeF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "pageSize", args)
  }

  return
}

// availableRedoSteps()
func (this *QTextDocument) Availableredosteps(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument18availableRedoStepsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "availableRedoSteps", args)
  }

  return
}

// defaultStyleSheet()
func (this *QTextDocument) Defaultstylesheet(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument17defaultStyleSheetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "defaultStyleSheet", args)
  }

  return
}

// resource(int, const class QUrl &)
func (this *QTextDocument) Resource(args ...interface{}) (ret interface{}) {
  // resource(int, const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument8resourceEiRK4QUrl
    // invoke: QVariant resource(int, const class QUrl &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QUrl).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QTextDocument8resourceEiRK4QUrl(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "resource", args)
  }

  return
}

// indentWidth()
func (this *QTextDocument) Indentwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument11indentWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "indentWidth", args)
  }

  return
}

// metaObject()
func (this *QTextDocument) Metaobject(args ...interface{}) () {
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
    C.C_ZNK13QTextDocument10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocument", "metaObject", args)
  }

  return
}

// lastBlock()
func (this *QTextDocument) Lastblock(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument9lastBlockEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "lastBlock", args)
  }

  return
}

// setUndoRedoEnabled(_Bool)
func (this *QTextDocument) Setundoredoenabled(args ...interface{}) () {
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
    C.C_ZN13QTextDocument18setUndoRedoEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setUndoRedoEnabled", args)
  }

  return
}

// isUndoAvailable()
func (this *QTextDocument) Isundoavailable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument15isUndoAvailableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "isUndoAvailable", args)
  }

  return
}

// setDefaultFont(const class QFont &)
func (this *QTextDocument) Setdefaultfont(args ...interface{}) () {
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
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument14setDefaultFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setDefaultFont", args)
  }

  return
}

// useDesignMetrics()
func (this *QTextDocument) Usedesignmetrics(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument16useDesignMetricsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "useDesignMetrics", args)
  }

  return
}

// textWidth()
func (this *QTextDocument) Textwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument9textWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "textWidth", args)
  }

  return
}

// characterCount()
func (this *QTextDocument) Charactercount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument14characterCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "characterCount", args)
  }

  return
}

// isUndoRedoEnabled()
func (this *QTextDocument) Isundoredoenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument17isUndoRedoEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "isUndoRedoEnabled", args)
  }

  return
}

// print(class QPagedPaintDevice *)
func (this *QTextDocument) Print(args ...interface{}) () {
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
    var arg0 = args[0].(*QPagedPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QTextDocument5printEP17QPagedPaintDevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "print", args)
  }

  return
}

// objectForFormat(const class QTextFormat &)
func (this *QTextDocument) Objectforformat(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTextFormat).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QTextDocument15objectForFormatERK11QTextFormat(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextObject{}) // "QTextObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "objectForFormat", args)
  }

  return
}

// toPlainText()
func (this *QTextDocument) Toplaintext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument11toPlainTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "toPlainText", args)
  }

  return
}

// undo()
func (this *QTextDocument) Undo(args ...interface{}) () {
  // undo()
  // undo(class QTextCursor *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextCursor{}) // "QTextCursor *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument4undoEv
    // invoke: void undo()
    C.C_ZN13QTextDocument4undoEv(this.Qclsinst)
  case 1:
    // invoke: _ZN13QTextDocument4undoEP11QTextCursor
    // invoke: void undo(class QTextCursor *)
    var arg0 = args[0].(*QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument4undoEP11QTextCursor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "undo", args)
  }

  return
}

// drawContents(class QPainter *, const class QRectF &)
func (this *QTextDocument) Drawcontents(args ...interface{}) () {
  // drawContents(class QPainter *, const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(qtcore.QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF
    // invoke: void drawContents(class QPainter *, const class QRectF &)
    var arg0 = args[0].(*QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QRectF).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextDocument", "drawContents", args)
  }

  return
}

// setModified(_Bool)
func (this *QTextDocument) Setmodified(args ...interface{}) () {
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
    C.C_ZN13QTextDocument11setModifiedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setModified", args)
  }

  return
}

// blockCount()
func (this *QTextDocument) Blockcount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument10blockCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "blockCount", args)
  }

  return
}

// defaultFont()
func (this *QTextDocument) Defaultfont(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument11defaultFontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "defaultFont", args)
  }

  return
}

// isRedoAvailable()
func (this *QTextDocument) Isredoavailable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument15isRedoAvailableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "isRedoAvailable", args)
  }

  return
}

// frameAt(int)
func (this *QTextDocument) Frameat(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QTextDocument7frameAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextFrame{}) // "QTextFrame *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "frameAt", args)
  }

  return
}

// pageCount()
func (this *QTextDocument) Pagecount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument9pageCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "pageCount", args)
  }

  return
}

// QTextDocument(class QObject *)
func NewQTextDocument(args ...interface{}) *QTextDocument {
  // QTextDocument(class QObject *)
  // QTextDocument(const class QString &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocumentC1EP7QObject
    // invoke: void QTextDocument(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QTextDocumentC2EP7QObject(arg0)
    return &QTextDocument{Qclsinst:qthis}
  case 1:
    // invoke: _ZN13QTextDocumentC1ERK7QStringP7QObject
    // invoke: void QTextDocument(const class QString &, class QObject *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QTextDocumentC2ERK7QStringP7QObject(arg0, arg1)
    return &QTextDocument{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextDocument", "QTextDocument", args)
  }

  return nil // QTextDocument{Qclsinst:qthis}
}

// ~QTextDocument()
func (this *QTextDocument) Freeqtextdocument(args ...interface{}) () {
  // ~QTextDocument()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocumentD0Ev
    // invoke: void ~QTextDocument()
    C.C_ZN13QTextDocumentD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocument", "~QTextDocument", args)
  }

  return
}

// availableUndoSteps()
func (this *QTextDocument) Availableundosteps(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument18availableUndoStepsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "availableUndoSteps", args)
  }

  return
}

// maximumBlockCount()
func (this *QTextDocument) Maximumblockcount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument17maximumBlockCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "maximumBlockCount", args)
  }

  return
}

// lineCount()
func (this *QTextDocument) Linecount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument9lineCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "lineCount", args)
  }

  return
}

// findBlock(int)
func (this *QTextDocument) Findblock(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QTextDocument9findBlockEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "findBlock", args)
  }

  return
}

// redo(class QTextCursor *)
func (this *QTextDocument) Redo(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument4redoEP11QTextCursor(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN13QTextDocument4redoEv
    // invoke: void redo()
    C.C_ZN13QTextDocument4redoEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocument", "redo", args)
  }

  return
}

// size()
func (this *QTextDocument) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSizeF{}) // "QSizeF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "size", args)
  }

  return
}

// isModified()
func (this *QTextDocument) Ismodified(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument10isModifiedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "isModified", args)
  }

  return
}

// findBlockByLineNumber(int)
func (this *QTextDocument) Findblockbylinenumber(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QTextDocument21findBlockByLineNumberEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "findBlockByLineNumber", args)
  }

  return
}

// baseUrl()
func (this *QTextDocument) Baseurl(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument7baseUrlEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QUrl{}) // "QUrl"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "baseUrl", args)
  }

  return
}

// documentMargin()
func (this *QTextDocument) Documentmargin(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument14documentMarginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "documentMargin", args)
  }

  return
}

// characterAt(int)
func (this *QTextDocument) Characterat(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QTextDocument11characterAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "characterAt", args)
  }

  return
}

// setDocumentMargin(qreal)
func (this *QTextDocument) Setdocumentmargin(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument17setDocumentMarginEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setDocumentMargin", args)
  }

  return
}

// rootFrame()
func (this *QTextDocument) Rootframe(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument9rootFrameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextFrame{}) // "QTextFrame *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "rootFrame", args)
  }

  return
}

// documentLayout()
func (this *QTextDocument) Documentlayout(args ...interface{}) () {
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
    C.C_ZNK13QTextDocument14documentLayoutEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocument", "documentLayout", args)
  }

  return
}

// defaultCursorMoveStyle()
func (this *QTextDocument) Defaultcursormovestyle(args ...interface{}) () {
  // defaultCursorMoveStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument22defaultCursorMoveStyleEv
    // invoke: Qt::CursorMoveStyle defaultCursorMoveStyle()
    C.C_ZNK13QTextDocument22defaultCursorMoveStyleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocument", "defaultCursorMoveStyle", args)
  }

  return
}

// setHtml(const class QString &)
func (this *QTextDocument) Sethtml(args ...interface{}) () {
  // setHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument7setHtmlERK7QString
    // invoke: void setHtml(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument7setHtmlERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setHtml", args)
  }

  return
}

// setMaximumBlockCount(int)
func (this *QTextDocument) Setmaximumblockcount(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument20setMaximumBlockCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setMaximumBlockCount", args)
  }

  return
}

// addResource(int, const class QUrl &, const class QVariant &)
func (this *QTextDocument) Addresource(args ...interface{}) () {
  // addResource(int, const class QUrl &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QUrl{}) // "const QUrl &"
  vtys[0][2] = reflect.TypeOf(qtcore.QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant
    // invoke: void addResource(int, const class QUrl &, const class QVariant &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QUrl).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QVariant).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTextDocument", "addResource", args)
  }

  return
}

// setDocumentLayout(class QAbstractTextDocumentLayout *)
func (this *QTextDocument) Setdocumentlayout(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractTextDocumentLayout).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setDocumentLayout", args)
  }

  return
}

// defaultTextOption()
func (this *QTextDocument) Defaulttextoption(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument17defaultTextOptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextOption{}) // "QTextOption"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "defaultTextOption", args)
  }

  return
}

// toHtml(const class QByteArray &)
func (this *QTextDocument) Tohtml(args ...interface{}) (ret interface{}) {
  // toHtml(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument6toHtmlERK10QByteArray
    // invoke: QString toHtml(const class QByteArray &)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QTextDocument6toHtmlERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "toHtml", args)
  }

  return
}

// firstBlock()
func (this *QTextDocument) Firstblock(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument10firstBlockEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "firstBlock", args)
  }

  return
}

// findBlockByNumber(int)
func (this *QTextDocument) Findblockbynumber(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QTextDocument17findBlockByNumberEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "findBlockByNumber", args)
  }

  return
}

// idealWidth()
func (this *QTextDocument) Idealwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument10idealWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "idealWidth", args)
  }

  return
}

// end()
func (this *QTextDocument) End(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument3endEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "end", args)
  }

  return
}

// setPageSize(const class QSizeF &)
func (this *QTextDocument) Setpagesize(args ...interface{}) () {
  // setPageSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument11setPageSizeERK6QSizeF
    // invoke: void setPageSize(const class QSizeF &)
    var arg0 = args[0].(*qtcore.QSizeF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument11setPageSizeERK6QSizeF(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setPageSize", args)
  }

  return
}

// setDefaultStyleSheet(const class QString &)
func (this *QTextDocument) Setdefaultstylesheet(args ...interface{}) () {
  // setDefaultStyleSheet(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument20setDefaultStyleSheetERK7QString
    // invoke: void setDefaultStyleSheet(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument20setDefaultStyleSheetERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setDefaultStyleSheet", args)
  }

  return
}

// appendUndoItem(class QAbstractUndoItem *)
func (this *QTextDocument) Appendundoitem(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractUndoItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "appendUndoItem", args)
  }

  return
}

// docHandle()
func (this *QTextDocument) Dochandle(args ...interface{}) () {
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
    C.C_ZNK13QTextDocument9docHandleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocument", "docHandle", args)
  }

  return
}

// isEmpty()
func (this *QTextDocument) Isempty(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "isEmpty", args)
  }

  return
}

// markContentsDirty(int, int)
func (this *QTextDocument) Markcontentsdirty(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN13QTextDocument17markContentsDirtyEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextDocument", "markContentsDirty", args)
  }

  return
}

// revision()
func (this *QTextDocument) Revision(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument8revisionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "revision", args)
  }

  return
}

// setBaseUrl(const class QUrl &)
func (this *QTextDocument) Setbaseurl(args ...interface{}) () {
  // setBaseUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument10setBaseUrlERK4QUrl
    // invoke: void setBaseUrl(const class QUrl &)
    var arg0 = args[0].(*qtcore.QUrl).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument10setBaseUrlERK4QUrl(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setBaseUrl", args)
  }

  return
}

// begin()
func (this *QTextDocument) Begin(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QTextDocument5beginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "begin", args)
  }

  return
}

// setPlainText(const class QString &)
func (this *QTextDocument) Setplaintext(args ...interface{}) () {
  // setPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTextDocument12setPlainTextERK7QString
    // invoke: void setPlainText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument12setPlainTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setPlainText", args)
  }

  return
}

// clone(class QObject *)
func (this *QTextDocument) Clone(args ...interface{}) (ret interface{}) {
  // clone(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTextDocument5cloneEP7QObject
    // invoke: QTextDocument * clone(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QTextDocument5cloneEP7QObject(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "clone", args)
  }

  return
}

// object(int)
func (this *QTextDocument) Object(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QTextDocument6objectEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextObject{}) // "QTextObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocument", "object", args)
  }

  return
}

// setTextWidth(qreal)
func (this *QTextDocument) Settextwidth(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument12setTextWidthEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setTextWidth", args)
  }

  return
}

// adjustSize()
func (this *QTextDocument) Adjustsize(args ...interface{}) () {
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
    C.C_ZN13QTextDocument10adjustSizeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocument", "adjustSize", args)
  }

  return
}

// allFormats()
func (this *QTextDocument) Allformats(args ...interface{}) () {
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
    C.C_ZNK13QTextDocument10allFormatsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocument", "allFormats", args)
  }

  return
}

// setDefaultTextOption(const class QTextOption &)
func (this *QTextDocument) Setdefaulttextoption(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextOption).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument20setDefaultTextOptionERK11QTextOption(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setDefaultTextOption", args)
  }

  return
}

// setUseDesignMetrics(_Bool)
func (this *QTextDocument) Setusedesignmetrics(args ...interface{}) () {
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
    C.C_ZN13QTextDocument19setUseDesignMetricsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setUseDesignMetrics", args)
  }

  return
}

// clear()
func (this *QTextDocument) Clear(args ...interface{}) () {
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
    C.C_ZN13QTextDocument5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocument", "clear", args)
  }

  return
}

// setIndentWidth(qreal)
func (this *QTextDocument) Setindentwidth(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN13QTextDocument14setIndentWidthEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocument", "setIndentWidth", args)
  }

  return
}

// <= body block end

