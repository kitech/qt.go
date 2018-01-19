//  header block begin
// /usr/include/qt/QtGui/qtextdocument.h
// #include <qtextdocument.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTextDocument struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextdocument.h:99
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTextDocument) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:119
// index:0
// void QTextDocument(class QObject *)
func NewQTextDocument(parent unsafe.Pointer) *QTextDocument {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocumentC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QTextDocument{cthis}
}

// /usr/include/qt/QtGui/qtextdocument.h:120
// index:1
// void QTextDocument(const class QString &, class QObject *)
func NewQTextDocument_1(text unsafe.Pointer, parent unsafe.Pointer) *QTextDocument {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocumentC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	return &QTextDocument{cthis}
}

// /usr/include/qt/QtGui/qtextdocument.h:121
// index:0
// virtual
// void ~QTextDocument()
func DeleteQTextDocument(*QTextDocument) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocumentD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:123
// index:0
// QTextDocument * clone(class QObject *)
func (this *QTextDocument) Clone(parent unsafe.Pointer) {
	// 0: (, QObject * parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument5cloneEP7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:125
// index:0
// bool isEmpty()
func (this *QTextDocument) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:126
// index:0
// virtual
// void clear()
func (this *QTextDocument) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:128
// index:0
// void setUndoRedoEnabled(_Bool)
func (this *QTextDocument) SetUndoRedoEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument18setUndoRedoEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:129
// index:0
// bool isUndoRedoEnabled()
func (this *QTextDocument) IsUndoRedoEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17isUndoRedoEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:131
// index:0
// bool isUndoAvailable()
func (this *QTextDocument) IsUndoAvailable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15isUndoAvailableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:132
// index:0
// bool isRedoAvailable()
func (this *QTextDocument) IsRedoAvailable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15isRedoAvailableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:134
// index:0
// int availableUndoSteps()
func (this *QTextDocument) AvailableUndoSteps() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument18availableUndoStepsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:135
// index:0
// int availableRedoSteps()
func (this *QTextDocument) AvailableRedoSteps() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument18availableRedoStepsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:137
// index:0
// int revision()
func (this *QTextDocument) Revision() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument8revisionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:139
// index:0
// void setDocumentLayout(class QAbstractTextDocumentLayout *)
func (this *QTextDocument) SetDocumentLayout(layout unsafe.Pointer) {
	// 0: (, QAbstractTextDocumentLayout * layout), (layout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout", ffiqt.FFI_TYPE_VOID, this.cthis, layout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:140
// index:0
// QAbstractTextDocumentLayout * documentLayout()
func (this *QTextDocument) DocumentLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument14documentLayoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:146
// index:0
// void setMetaInformation(enum QTextDocument::MetaInformation, const class QString &)
func (this *QTextDocument) SetMetaInformation(info int, arg1 unsafe.Pointer) {
	// 0: (, QTextDocument::MetaInformation info, const QString & arg1), (&info, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument18setMetaInformationENS_15MetaInformationERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &info, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:147
// index:0
// QString metaInformation(enum QTextDocument::MetaInformation)
func (this *QTextDocument) MetaInformation(info int) {
	// 0: (, QTextDocument::MetaInformation info), (&info)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15metaInformationENS_15MetaInformationE", ffiqt.FFI_TYPE_VOID, this.cthis, &info)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:150
// index:0
// QString toHtml(const class QByteArray &)
func (this *QTextDocument) ToHtml(encoding unsafe.Pointer) {
	// 0: (, const QByteArray & encoding), (encoding)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument6toHtmlERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, encoding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:151
// index:0
// void setHtml(const class QString &)
func (this *QTextDocument) SetHtml(html unsafe.Pointer) {
	// 0: (, const QString & html), (html)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument7setHtmlERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, html)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:154
// index:0
// QString toRawText()
func (this *QTextDocument) ToRawText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9toRawTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:155
// index:0
// QString toPlainText()
func (this *QTextDocument) ToPlainText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11toPlainTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:156
// index:0
// void setPlainText(const class QString &)
func (this *QTextDocument) SetPlainText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12setPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:158
// index:0
// QChar characterAt(int)
func (this *QTextDocument) CharacterAt(pos int) {
	// 0: (, int pos), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11characterAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:181
// index:0
// QTextFrame * frameAt(int)
func (this *QTextDocument) FrameAt(pos int) {
	// 0: (, int pos), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument7frameAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:182
// index:0
// QTextFrame * rootFrame()
func (this *QTextDocument) RootFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9rootFrameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:184
// index:0
// QTextObject * object(int)
func (this *QTextDocument) Object(objectIndex int) {
	// 0: (, int objectIndex), (&objectIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument6objectEi", ffiqt.FFI_TYPE_VOID, this.cthis, &objectIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:185
// index:0
// QTextObject * objectForFormat(const class QTextFormat &)
func (this *QTextDocument) ObjectForFormat(arg0 unsafe.Pointer) {
	// 0: (, const QTextFormat & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15objectForFormatERK11QTextFormat", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:187
// index:0
// QTextBlock findBlock(int)
func (this *QTextDocument) FindBlock(pos int) {
	// 0: (, int pos), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9findBlockEi", ffiqt.FFI_TYPE_VOID, this.cthis, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:188
// index:0
// QTextBlock findBlockByNumber(int)
func (this *QTextDocument) FindBlockByNumber(blockNumber int) {
	// 0: (, int blockNumber), (&blockNumber)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17findBlockByNumberEi", ffiqt.FFI_TYPE_VOID, this.cthis, &blockNumber)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:189
// index:0
// QTextBlock findBlockByLineNumber(int)
func (this *QTextDocument) FindBlockByLineNumber(blockNumber int) {
	// 0: (, int blockNumber), (&blockNumber)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument21findBlockByLineNumberEi", ffiqt.FFI_TYPE_VOID, this.cthis, &blockNumber)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:190
// index:0
// QTextBlock begin()
func (this *QTextDocument) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument5beginEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:191
// index:0
// QTextBlock end()
func (this *QTextDocument) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument3endEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:193
// index:0
// QTextBlock firstBlock()
func (this *QTextDocument) FirstBlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10firstBlockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:194
// index:0
// QTextBlock lastBlock()
func (this *QTextDocument) LastBlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9lastBlockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:196
// index:0
// void setPageSize(const class QSizeF &)
func (this *QTextDocument) SetPageSize(size unsafe.Pointer) {
	// 0: (, const QSizeF & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument11setPageSizeERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:197
// index:0
// QSizeF pageSize()
func (this *QTextDocument) PageSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument8pageSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:199
// index:0
// void setDefaultFont(const class QFont &)
func (this *QTextDocument) SetDefaultFont(font unsafe.Pointer) {
	// 0: (, const QFont & font), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14setDefaultFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:200
// index:0
// QFont defaultFont()
func (this *QTextDocument) DefaultFont() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11defaultFontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:202
// index:0
// int pageCount()
func (this *QTextDocument) PageCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9pageCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:204
// index:0
// bool isModified()
func (this *QTextDocument) IsModified() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10isModifiedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:206
// index:0
// void print(class QPagedPaintDevice *)
func (this *QTextDocument) Print(printer unsafe.Pointer) {
	// 0: (, QPagedPaintDevice * printer), (printer)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument5printEP17QPagedPaintDevice", ffiqt.FFI_TYPE_VOID, this.cthis, printer)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:216
// index:0
// QVariant resource(int, const class QUrl &)
func (this *QTextDocument) Resource(type_ int, name unsafe.Pointer) {
	// 0: (, int type, const QUrl & name), (&type_, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument8resourceEiRK4QUrl", ffiqt.FFI_TYPE_VOID, this.cthis, &type_, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:217
// index:0
// void addResource(int, const class QUrl &, const class QVariant &)
func (this *QTextDocument) AddResource(type_ int, name unsafe.Pointer, resource unsafe.Pointer) {
	// 0: (, int type, const QUrl & name, const QVariant & resource), (&type_, name, resource)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &type_, name, resource)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:219
// index:0
// QVector<QTextFormat> allFormats()
func (this *QTextDocument) AllFormats() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10allFormatsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:221
// index:0
// void markContentsDirty(int, int)
func (this *QTextDocument) MarkContentsDirty(from int, length int) {
	// 0: (, int from, int length), (&from, &length)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17markContentsDirtyEii", ffiqt.FFI_TYPE_VOID, this.cthis, &from, &length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:223
// index:0
// void setUseDesignMetrics(_Bool)
func (this *QTextDocument) SetUseDesignMetrics(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument19setUseDesignMetricsEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:224
// index:0
// bool useDesignMetrics()
func (this *QTextDocument) UseDesignMetrics() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument16useDesignMetricsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:226
// index:0
// void drawContents(class QPainter *, const class QRectF &)
func (this *QTextDocument) DrawContents(painter unsafe.Pointer, rect unsafe.Pointer) {
	// 0: (, QPainter * painter, const QRectF & rect), (painter, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, painter, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:228
// index:0
// void setTextWidth(qreal)
func (this *QTextDocument) SetTextWidth(width float64) {
	// 0: (, qreal width), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12setTextWidthEd", ffiqt.FFI_TYPE_VOID, this.cthis, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:229
// index:0
// qreal textWidth()
func (this *QTextDocument) TextWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9textWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:231
// index:0
// qreal idealWidth()
func (this *QTextDocument) IdealWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10idealWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:233
// index:0
// qreal indentWidth()
func (this *QTextDocument) IndentWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11indentWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:234
// index:0
// void setIndentWidth(qreal)
func (this *QTextDocument) SetIndentWidth(width float64) {
	// 0: (, qreal width), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14setIndentWidthEd", ffiqt.FFI_TYPE_VOID, this.cthis, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:236
// index:0
// qreal documentMargin()
func (this *QTextDocument) DocumentMargin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument14documentMarginEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:237
// index:0
// void setDocumentMargin(qreal)
func (this *QTextDocument) SetDocumentMargin(margin float64) {
	// 0: (, qreal margin), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17setDocumentMarginEd", ffiqt.FFI_TYPE_VOID, this.cthis, &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:239
// index:0
// void adjustSize()
func (this *QTextDocument) AdjustSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument10adjustSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:240
// index:0
// QSizeF size()
func (this *QTextDocument) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:242
// index:0
// int blockCount()
func (this *QTextDocument) BlockCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10blockCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:243
// index:0
// int lineCount()
func (this *QTextDocument) LineCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9lineCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:244
// index:0
// int characterCount()
func (this *QTextDocument) CharacterCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument14characterCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:247
// index:0
// void setDefaultStyleSheet(const class QString &)
func (this *QTextDocument) SetDefaultStyleSheet(sheet unsafe.Pointer) {
	// 0: (, const QString & sheet), (sheet)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument20setDefaultStyleSheetERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, sheet)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:248
// index:0
// QString defaultStyleSheet()
func (this *QTextDocument) DefaultStyleSheet() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17defaultStyleSheetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:251
// index:0
// void undo(class QTextCursor *)
func (this *QTextDocument) Undo(cursor unsafe.Pointer) {
	// 0: (, QTextCursor * cursor), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument4undoEP11QTextCursor", ffiqt.FFI_TYPE_VOID, this.cthis, cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:286
// index:1
// void undo()
func (this *QTextDocument) Undo_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument4undoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:252
// index:0
// void redo(class QTextCursor *)
func (this *QTextDocument) Redo(cursor unsafe.Pointer) {
	// 0: (, QTextCursor * cursor), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument4redoEP11QTextCursor", ffiqt.FFI_TYPE_VOID, this.cthis, cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:287
// index:1
// void redo()
func (this *QTextDocument) Redo_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument4redoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:259
// index:0
// void clearUndoRedoStacks(enum QTextDocument::Stacks)
func (this *QTextDocument) ClearUndoRedoStacks(historyToClear int) {
	// 0: (, QTextDocument::Stacks historyToClear), (&historyToClear)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument19clearUndoRedoStacksENS_6StacksE", ffiqt.FFI_TYPE_VOID, this.cthis, &historyToClear)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:261
// index:0
// int maximumBlockCount()
func (this *QTextDocument) MaximumBlockCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17maximumBlockCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:262
// index:0
// void setMaximumBlockCount(int)
func (this *QTextDocument) SetMaximumBlockCount(maximum int) {
	// 0: (, int maximum), (&maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument20setMaximumBlockCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:264
// index:0
// QTextOption defaultTextOption()
func (this *QTextDocument) DefaultTextOption() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17defaultTextOptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:265
// index:0
// void setDefaultTextOption(const class QTextOption &)
func (this *QTextDocument) SetDefaultTextOption(option unsafe.Pointer) {
	// 0: (, const QTextOption & option), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument20setDefaultTextOptionERK11QTextOption", ffiqt.FFI_TYPE_VOID, this.cthis, option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:267
// index:0
// QUrl baseUrl()
func (this *QTextDocument) BaseUrl() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument7baseUrlEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:268
// index:0
// void setBaseUrl(const class QUrl &)
func (this *QTextDocument) SetBaseUrl(url unsafe.Pointer) {
	// 0: (, const QUrl & url), (url)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument10setBaseUrlERK4QUrl", ffiqt.FFI_TYPE_VOID, this.cthis, url)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:270
// index:0
// Qt::CursorMoveStyle defaultCursorMoveStyle()
func (this *QTextDocument) DefaultCursorMoveStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument22defaultCursorMoveStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:271
// index:0
// void setDefaultCursorMoveStyle(Qt::CursorMoveStyle)
func (this *QTextDocument) SetDefaultCursorMoveStyle(style int) {
	// 0: (, Qt::CursorMoveStyle style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument25setDefaultCursorMoveStyleEN2Qt15CursorMoveStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:274
// index:0
// void contentsChange(int, int, int)
func (this *QTextDocument) ContentsChange(from int, charsRemoved int, charsAdded int) {
	// 0: (, int from, int charsRemoved, int charsAdded), (&from, &charsRemoved, &charsAdded)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14contentsChangeEiii", ffiqt.FFI_TYPE_VOID, this.cthis, &from, &charsRemoved, &charsAdded)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:275
// index:0
// void contentsChanged()
func (this *QTextDocument) ContentsChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument15contentsChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:276
// index:0
// void undoAvailable(_Bool)
func (this *QTextDocument) UndoAvailable(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument13undoAvailableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:277
// index:0
// void redoAvailable(_Bool)
func (this *QTextDocument) RedoAvailable(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument13redoAvailableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:278
// index:0
// void undoCommandAdded()
func (this *QTextDocument) UndoCommandAdded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument16undoCommandAddedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:279
// index:0
// void modificationChanged(_Bool)
func (this *QTextDocument) ModificationChanged(m bool) {
	// 0: (, bool m), (&m)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument19modificationChangedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &m)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:280
// index:0
// void cursorPositionChanged(const class QTextCursor &)
func (this *QTextDocument) CursorPositionChanged(cursor unsafe.Pointer) {
	// 0: (, const QTextCursor & cursor), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument21cursorPositionChangedERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.cthis, cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:281
// index:0
// void blockCountChanged(int)
func (this *QTextDocument) BlockCountChanged(newBlockCount int) {
	// 0: (, int newBlockCount), (&newBlockCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17blockCountChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &newBlockCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:282
// index:0
// void baseUrlChanged(const class QUrl &)
func (this *QTextDocument) BaseUrlChanged(url unsafe.Pointer) {
	// 0: (, const QUrl & url), (url)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14baseUrlChangedERK4QUrl", ffiqt.FFI_TYPE_VOID, this.cthis, url)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:283
// index:0
// void documentLayoutChanged()
func (this *QTextDocument) DocumentLayoutChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument21documentLayoutChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:288
// index:0
// void appendUndoItem(class QAbstractUndoItem *)
func (this *QTextDocument) AppendUndoItem(arg0 unsafe.Pointer) {
	// 0: (, QAbstractUndoItem * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:289
// index:0
// void setModified(_Bool)
func (this *QTextDocument) SetModified(m bool) {
	// 0: (, bool m), (&m)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument11setModifiedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &m)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:297
// index:0
// QTextDocumentPrivate * docHandle()
func (this *QTextDocument) DocHandle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9docHandleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
