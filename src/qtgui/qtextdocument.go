//  header block begin
// /usr/include/qt/QtGui/qtextdocument.h
// #include <qtextdocument.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
	*qtcore.QObject
}

func (this *QTextDocument) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQTextDocumentFromPointer(cthis unsafe.Pointer) *QTextDocument {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QTextDocument{bcthis0}
}

// /usr/include/qt/QtGui/qtextdocument.h:99
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTextDocument) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:119
// index:0
// Public
// void QTextDocument(class QObject *)
func NewQTextDocument(parent unsafe.Pointer) *QTextDocument {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocumentC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:120
// index:1
// Public
// void QTextDocument(const class QString &, class QObject *)
func NewQTextDocument_1(text *qtcore.QString, parent unsafe.Pointer) *QTextDocument {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocumentC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:121
// index:0
// Public virtual
// void ~QTextDocument()
func DeleteQTextDocument(*QTextDocument) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocumentD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:123
// index:0
// Public
// QTextDocument * clone(class QObject *)
func (this *QTextDocument) Clone(parent unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument5cloneEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:125
// index:0
// Public
// bool isEmpty()
func (this *QTextDocument) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:126
// index:0
// Public virtual
// void clear()
func (this *QTextDocument) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:128
// index:0
// Public
// void setUndoRedoEnabled(_Bool)
func (this *QTextDocument) SetUndoRedoEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument18setUndoRedoEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:129
// index:0
// Public
// bool isUndoRedoEnabled()
func (this *QTextDocument) IsUndoRedoEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17isUndoRedoEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:131
// index:0
// Public
// bool isUndoAvailable()
func (this *QTextDocument) IsUndoAvailable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15isUndoAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:132
// index:0
// Public
// bool isRedoAvailable()
func (this *QTextDocument) IsRedoAvailable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15isRedoAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:134
// index:0
// Public
// int availableUndoSteps()
func (this *QTextDocument) AvailableUndoSteps() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument18availableUndoStepsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:135
// index:0
// Public
// int availableRedoSteps()
func (this *QTextDocument) AvailableRedoSteps() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument18availableRedoStepsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:137
// index:0
// Public
// int revision()
func (this *QTextDocument) Revision() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument8revisionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:139
// index:0
// Public
// void setDocumentLayout(class QAbstractTextDocumentLayout *)
func (this *QTextDocument) SetDocumentLayout(layout unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), layout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:140
// index:0
// Public
// QAbstractTextDocumentLayout * documentLayout()
func (this *QTextDocument) DocumentLayout() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument14documentLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:146
// index:0
// Public
// void setMetaInformation(enum QTextDocument::MetaInformation, const class QString &)
func (this *QTextDocument) SetMetaInformation(info int, arg1 *qtcore.QString) {
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument18setMetaInformationENS_15MetaInformationERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &info, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:147
// index:0
// Public
// QString metaInformation(enum QTextDocument::MetaInformation)
func (this *QTextDocument) MetaInformation(info int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15metaInformationENS_15MetaInformationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &info)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:150
// index:0
// Public
// QString toHtml(const class QByteArray &)
func (this *QTextDocument) ToHtml(encoding *qtcore.QByteArray) interface{} {
	var convArg0 = encoding.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument6toHtmlERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:151
// index:0
// Public
// void setHtml(const class QString &)
func (this *QTextDocument) SetHtml(html *qtcore.QString) {
	var convArg0 = html.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument7setHtmlERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:154
// index:0
// Public
// QString toRawText()
func (this *QTextDocument) ToRawText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9toRawTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:155
// index:0
// Public
// QString toPlainText()
func (this *QTextDocument) ToPlainText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11toPlainTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:156
// index:0
// Public
// void setPlainText(const class QString &)
func (this *QTextDocument) SetPlainText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12setPlainTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:158
// index:0
// Public
// QChar characterAt(int)
func (this *QTextDocument) CharacterAt(pos int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11characterAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:181
// index:0
// Public
// QTextFrame * frameAt(int)
func (this *QTextDocument) FrameAt(pos int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument7frameAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:182
// index:0
// Public
// QTextFrame * rootFrame()
func (this *QTextDocument) RootFrame() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9rootFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:184
// index:0
// Public
// QTextObject * object(int)
func (this *QTextDocument) Object(objectIndex int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument6objectEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &objectIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:185
// index:0
// Public
// QTextObject * objectForFormat(const class QTextFormat &)
func (this *QTextDocument) ObjectForFormat(arg0 *QTextFormat) interface{} {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15objectForFormatERK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:187
// index:0
// Public
// QTextBlock findBlock(int)
func (this *QTextDocument) FindBlock(pos int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9findBlockEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:188
// index:0
// Public
// QTextBlock findBlockByNumber(int)
func (this *QTextDocument) FindBlockByNumber(blockNumber int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17findBlockByNumberEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &blockNumber)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:189
// index:0
// Public
// QTextBlock findBlockByLineNumber(int)
func (this *QTextDocument) FindBlockByLineNumber(blockNumber int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument21findBlockByLineNumberEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &blockNumber)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:190
// index:0
// Public
// QTextBlock begin()
func (this *QTextDocument) Begin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:191
// index:0
// Public
// QTextBlock end()
func (this *QTextDocument) End() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:193
// index:0
// Public
// QTextBlock firstBlock()
func (this *QTextDocument) FirstBlock() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10firstBlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:194
// index:0
// Public
// QTextBlock lastBlock()
func (this *QTextDocument) LastBlock() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9lastBlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:196
// index:0
// Public
// void setPageSize(const class QSizeF &)
func (this *QTextDocument) SetPageSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument11setPageSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:197
// index:0
// Public
// QSizeF pageSize()
func (this *QTextDocument) PageSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument8pageSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:199
// index:0
// Public
// void setDefaultFont(const class QFont &)
func (this *QTextDocument) SetDefaultFont(font *QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14setDefaultFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:200
// index:0
// Public
// QFont defaultFont()
func (this *QTextDocument) DefaultFont() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11defaultFontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:202
// index:0
// Public
// int pageCount()
func (this *QTextDocument) PageCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9pageCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:204
// index:0
// Public
// bool isModified()
func (this *QTextDocument) IsModified() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10isModifiedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:206
// index:0
// Public
// void print(class QPagedPaintDevice *)
func (this *QTextDocument) Print(printer unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument5printEP17QPagedPaintDevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), printer)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:216
// index:0
// Public
// QVariant resource(int, const class QUrl &)
func (this *QTextDocument) Resource(type_ int, name *qtcore.QUrl) interface{} {
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument8resourceEiRK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:217
// index:0
// Public
// void addResource(int, const class QUrl &, const class QVariant &)
func (this *QTextDocument) AddResource(type_ int, name *qtcore.QUrl, resource *qtcore.QVariant) {
	var convArg1 = name.GetCthis()
	var convArg2 = resource.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:219
// index:0
// Public
// QVector<QTextFormat> allFormats()
func (this *QTextDocument) AllFormats() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10allFormatsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:221
// index:0
// Public
// void markContentsDirty(int, int)
func (this *QTextDocument) MarkContentsDirty(from int, length int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17markContentsDirtyEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &from, &length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:223
// index:0
// Public
// void setUseDesignMetrics(_Bool)
func (this *QTextDocument) SetUseDesignMetrics(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument19setUseDesignMetricsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:224
// index:0
// Public
// bool useDesignMetrics()
func (this *QTextDocument) UseDesignMetrics() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument16useDesignMetricsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:226
// index:0
// Public
// void drawContents(class QPainter *, const class QRectF &)
func (this *QTextDocument) DrawContents(painter unsafe.Pointer, rect *qtcore.QRectF) {
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:228
// index:0
// Public
// void setTextWidth(qreal)
func (this *QTextDocument) SetTextWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12setTextWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:229
// index:0
// Public
// qreal textWidth()
func (this *QTextDocument) TextWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9textWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:231
// index:0
// Public
// qreal idealWidth()
func (this *QTextDocument) IdealWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10idealWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:233
// index:0
// Public
// qreal indentWidth()
func (this *QTextDocument) IndentWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11indentWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:234
// index:0
// Public
// void setIndentWidth(qreal)
func (this *QTextDocument) SetIndentWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14setIndentWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:236
// index:0
// Public
// qreal documentMargin()
func (this *QTextDocument) DocumentMargin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument14documentMarginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:237
// index:0
// Public
// void setDocumentMargin(qreal)
func (this *QTextDocument) SetDocumentMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17setDocumentMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:239
// index:0
// Public
// void adjustSize()
func (this *QTextDocument) AdjustSize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument10adjustSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:240
// index:0
// Public
// QSizeF size()
func (this *QTextDocument) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:242
// index:0
// Public
// int blockCount()
func (this *QTextDocument) BlockCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10blockCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:243
// index:0
// Public
// int lineCount()
func (this *QTextDocument) LineCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9lineCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:244
// index:0
// Public
// int characterCount()
func (this *QTextDocument) CharacterCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument14characterCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:247
// index:0
// Public
// void setDefaultStyleSheet(const class QString &)
func (this *QTextDocument) SetDefaultStyleSheet(sheet *qtcore.QString) {
	var convArg0 = sheet.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument20setDefaultStyleSheetERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:248
// index:0
// Public
// QString defaultStyleSheet()
func (this *QTextDocument) DefaultStyleSheet() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17defaultStyleSheetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:251
// index:0
// Public
// void undo(class QTextCursor *)
func (this *QTextDocument) Undo(cursor unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument4undoEP11QTextCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:286
// index:1
// Public
// void undo()
func (this *QTextDocument) Undo_1() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument4undoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:252
// index:0
// Public
// void redo(class QTextCursor *)
func (this *QTextDocument) Redo(cursor unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument4redoEP11QTextCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:287
// index:1
// Public
// void redo()
func (this *QTextDocument) Redo_1() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument4redoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:259
// index:0
// Public
// void clearUndoRedoStacks(enum QTextDocument::Stacks)
func (this *QTextDocument) ClearUndoRedoStacks(historyToClear int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument19clearUndoRedoStacksENS_6StacksE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &historyToClear)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:261
// index:0
// Public
// int maximumBlockCount()
func (this *QTextDocument) MaximumBlockCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17maximumBlockCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:262
// index:0
// Public
// void setMaximumBlockCount(int)
func (this *QTextDocument) SetMaximumBlockCount(maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument20setMaximumBlockCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:264
// index:0
// Public
// QTextOption defaultTextOption()
func (this *QTextDocument) DefaultTextOption() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17defaultTextOptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:265
// index:0
// Public
// void setDefaultTextOption(const class QTextOption &)
func (this *QTextDocument) SetDefaultTextOption(option *QTextOption) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument20setDefaultTextOptionERK11QTextOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:267
// index:0
// Public
// QUrl baseUrl()
func (this *QTextDocument) BaseUrl() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument7baseUrlEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:268
// index:0
// Public
// void setBaseUrl(const class QUrl &)
func (this *QTextDocument) SetBaseUrl(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument10setBaseUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:270
// index:0
// Public
// Qt::CursorMoveStyle defaultCursorMoveStyle()
func (this *QTextDocument) DefaultCursorMoveStyle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument22defaultCursorMoveStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:271
// index:0
// Public
// void setDefaultCursorMoveStyle(Qt::CursorMoveStyle)
func (this *QTextDocument) SetDefaultCursorMoveStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument25setDefaultCursorMoveStyleEN2Qt15CursorMoveStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:274
// index:0
// Public
// void contentsChange(int, int, int)
func (this *QTextDocument) ContentsChange(from int, charsRemoved int, charsAdded int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14contentsChangeEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &from, &charsRemoved, &charsAdded)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:275
// index:0
// Public
// void contentsChanged()
func (this *QTextDocument) ContentsChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument15contentsChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:276
// index:0
// Public
// void undoAvailable(_Bool)
func (this *QTextDocument) UndoAvailable(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument13undoAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:277
// index:0
// Public
// void redoAvailable(_Bool)
func (this *QTextDocument) RedoAvailable(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument13redoAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:278
// index:0
// Public
// void undoCommandAdded()
func (this *QTextDocument) UndoCommandAdded() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument16undoCommandAddedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:279
// index:0
// Public
// void modificationChanged(_Bool)
func (this *QTextDocument) ModificationChanged(m bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument19modificationChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &m)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:280
// index:0
// Public
// void cursorPositionChanged(const class QTextCursor &)
func (this *QTextDocument) CursorPositionChanged(cursor *QTextCursor) {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument21cursorPositionChangedERK11QTextCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:281
// index:0
// Public
// void blockCountChanged(int)
func (this *QTextDocument) BlockCountChanged(newBlockCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17blockCountChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &newBlockCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:282
// index:0
// Public
// void baseUrlChanged(const class QUrl &)
func (this *QTextDocument) BaseUrlChanged(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14baseUrlChangedERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:283
// index:0
// Public
// void documentLayoutChanged()
func (this *QTextDocument) DocumentLayoutChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument21documentLayoutChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:288
// index:0
// Public
// void appendUndoItem(class QAbstractUndoItem *)
func (this *QTextDocument) AppendUndoItem(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:289
// index:0
// Public
// void setModified(_Bool)
func (this *QTextDocument) SetModified(m bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument11setModifiedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &m)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:292
// index:0
// Protected virtual
// QTextObject * createObject(const class QTextFormat &)
func (this *QTextDocument) CreateObject(f *QTextFormat) interface{} {
	var convArg0 = f.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12createObjectERK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:293
// index:0
// Protected virtual
// QVariant loadResource(int, const class QUrl &)
func (this *QTextDocument) LoadResource(type_ int, name *qtcore.QUrl) interface{} {
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12loadResourceEiRK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocument.h:297
// index:0
// Public
// QTextDocumentPrivate * docHandle()
func (this *QTextDocument) DocHandle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9docHandleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
