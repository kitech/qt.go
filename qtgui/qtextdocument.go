package qtgui

// /usr/include/qt/QtGui/qtextdocument.h
// #include <qtextdocument.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QTextDocument) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQTextDocumentFromPointer(cthis unsafe.Pointer) *QTextDocument {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QTextDocument{bcthis0}
}
func (*QTextDocument) NewFromPointer(cthis unsafe.Pointer) *QTextDocument {
	return NewQTextDocumentFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextdocument.h:99
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTextDocument) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:119
// index:0
// Public
// void QTextDocument(QObject *)
func NewQTextDocument(parent *qtcore.QObject /*777 QObject **/) *QTextDocument {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocumentC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:120
// index:1
// Public
// void QTextDocument(const QString &, QObject *)
func NewQTextDocument_1(text *qtcore.QString, parent *qtcore.QObject /*777 QObject **/) *QTextDocument {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocumentC2ERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
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
// QTextDocument * clone(QObject *)
func (this *QTextDocument) Clone(parent *qtcore.QObject /*777 QObject **/) *QTextDocument /*777 QTextDocument **/ {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument5cloneEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:125
// index:0
// Public
// bool isEmpty()
func (this *QTextDocument) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
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
// void setUndoRedoEnabled(bool)
func (this *QTextDocument) SetUndoRedoEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument18setUndoRedoEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:129
// index:0
// Public
// bool isUndoRedoEnabled()
func (this *QTextDocument) IsUndoRedoEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17isUndoRedoEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:131
// index:0
// Public
// bool isUndoAvailable()
func (this *QTextDocument) IsUndoAvailable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15isUndoAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:132
// index:0
// Public
// bool isRedoAvailable()
func (this *QTextDocument) IsRedoAvailable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15isRedoAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:134
// index:0
// Public
// int availableUndoSteps()
func (this *QTextDocument) AvailableUndoSteps() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument18availableUndoStepsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextdocument.h:135
// index:0
// Public
// int availableRedoSteps()
func (this *QTextDocument) AvailableRedoSteps() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument18availableRedoStepsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextdocument.h:137
// index:0
// Public
// int revision()
func (this *QTextDocument) Revision() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument8revisionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextdocument.h:139
// index:0
// Public
// void setDocumentLayout(QAbstractTextDocumentLayout *)
func (this *QTextDocument) SetDocumentLayout(layout *QAbstractTextDocumentLayout /*777 QAbstractTextDocumentLayout **/) {
	var convArg0 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:140
// index:0
// Public
// QAbstractTextDocumentLayout * documentLayout()
func (this *QTextDocument) DocumentLayout() *QAbstractTextDocumentLayout /*777 QAbstractTextDocumentLayout **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument14documentLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractTextDocumentLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:146
// index:0
// Public
// void setMetaInformation(QTextDocument::MetaInformation, const QString &)
func (this *QTextDocument) SetMetaInformation(info int, arg1 *qtcore.QString) {
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument18setMetaInformationENS_15MetaInformationERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), info, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:147
// index:0
// Public
// QString metaInformation(QTextDocument::MetaInformation)
func (this *QTextDocument) MetaInformation(info int) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15metaInformationENS_15MetaInformationE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), info)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:150
// index:0
// Public
// QString toHtml(const QByteArray &)
func (this *QTextDocument) ToHtml(encoding *qtcore.QByteArray) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = encoding.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument6toHtmlERK10QByteArray", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:151
// index:0
// Public
// void setHtml(const QString &)
func (this *QTextDocument) SetHtml(html *qtcore.QString) {
	var convArg0 = html.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument7setHtmlERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:154
// index:0
// Public
// QString toRawText()
func (this *QTextDocument) ToRawText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9toRawTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:155
// index:0
// Public
// QString toPlainText()
func (this *QTextDocument) ToPlainText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11toPlainTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:156
// index:0
// Public
// void setPlainText(const QString &)
func (this *QTextDocument) SetPlainText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12setPlainTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:158
// index:0
// Public
// QChar characterAt(int)
func (this *QTextDocument) CharacterAt(pos int) *qtcore.QChar /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11characterAtEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:168
// index:0
// Public
// QTextCursor find(const QString &, int, QTextDocument::FindFlags)
func (this *QTextDocument) Find(subString *qtcore.QString, from int, options int) *QTextCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = subString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringi6QFlagsINS_8FindFlagEE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, from, options)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:169
// index:1
// Public
// QTextCursor find(const QString &, const QTextCursor &, QTextDocument::FindFlags)
func (this *QTextDocument) Find_1(subString *qtcore.QString, cursor *QTextCursor, options int) *QTextCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = subString.GetCthis()
	var convArg1 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringRK11QTextCursor6QFlagsINS_8FindFlagEE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1, options)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:172
// index:2
// Public
// QTextCursor find(const QRegExp &, int, QTextDocument::FindFlags)
func (this *QTextDocument) Find_2(expr *qtcore.QRegExp, from int, options int) *QTextCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = expr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpi6QFlagsINS_8FindFlagEE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, from, options)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:173
// index:3
// Public
// QTextCursor find(const QRegExp &, const QTextCursor &, QTextDocument::FindFlags)
func (this *QTextDocument) Find_3(expr *qtcore.QRegExp, cursor *QTextCursor, options int) *QTextCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = expr.GetCthis()
	var convArg1 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpRK11QTextCursor6QFlagsINS_8FindFlagEE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1, options)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:177
// index:4
// Public
// QTextCursor find(const QRegularExpression &, int, QTextDocument::FindFlags)
func (this *QTextDocument) Find_4(expr *qtcore.QRegularExpression, from int, options int) *QTextCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = expr.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressioni6QFlagsINS_8FindFlagEE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, from, options)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:178
// index:5
// Public
// QTextCursor find(const QRegularExpression &, const QTextCursor &, QTextDocument::FindFlags)
func (this *QTextDocument) Find_5(expr *qtcore.QRegularExpression, cursor *QTextCursor, options int) *QTextCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = expr.GetCthis()
	var convArg1 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressionRK11QTextCursor6QFlagsINS_8FindFlagEE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1, options)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:181
// index:0
// Public
// QTextFrame * frameAt(int)
func (this *QTextDocument) FrameAt(pos int) *QTextFrame /*777 QTextFrame **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument7frameAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:182
// index:0
// Public
// QTextFrame * rootFrame()
func (this *QTextDocument) RootFrame() *QTextFrame /*777 QTextFrame **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9rootFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:184
// index:0
// Public
// QTextObject * object(int)
func (this *QTextDocument) Object(objectIndex int) *QTextObject /*777 QTextObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument6objectEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), objectIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:185
// index:0
// Public
// QTextObject * objectForFormat(const QTextFormat &)
func (this *QTextDocument) ObjectForFormat(arg0 *QTextFormat) *QTextObject /*777 QTextObject **/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument15objectForFormatERK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:187
// index:0
// Public
// QTextBlock findBlock(int)
func (this *QTextDocument) FindBlock(pos int) *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9findBlockEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:188
// index:0
// Public
// QTextBlock findBlockByNumber(int)
func (this *QTextDocument) FindBlockByNumber(blockNumber int) *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17findBlockByNumberEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), blockNumber)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:189
// index:0
// Public
// QTextBlock findBlockByLineNumber(int)
func (this *QTextDocument) FindBlockByLineNumber(blockNumber int) *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument21findBlockByLineNumberEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), blockNumber)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:190
// index:0
// Public
// QTextBlock begin()
func (this *QTextDocument) Begin() *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument5beginEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:191
// index:0
// Public
// QTextBlock end()
func (this *QTextDocument) End() *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument3endEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:193
// index:0
// Public
// QTextBlock firstBlock()
func (this *QTextDocument) FirstBlock() *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10firstBlockEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:194
// index:0
// Public
// QTextBlock lastBlock()
func (this *QTextDocument) LastBlock() *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9lastBlockEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:196
// index:0
// Public
// void setPageSize(const QSizeF &)
func (this *QTextDocument) SetPageSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument11setPageSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:197
// index:0
// Public
// QSizeF pageSize()
func (this *QTextDocument) PageSize() *qtcore.QSizeF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument8pageSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:199
// index:0
// Public
// void setDefaultFont(const QFont &)
func (this *QTextDocument) SetDefaultFont(font *QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14setDefaultFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:200
// index:0
// Public
// QFont defaultFont()
func (this *QTextDocument) DefaultFont() *QFont /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11defaultFontEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:202
// index:0
// Public
// int pageCount()
func (this *QTextDocument) PageCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9pageCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextdocument.h:204
// index:0
// Public
// bool isModified()
func (this *QTextDocument) IsModified() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10isModifiedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:206
// index:0
// Public
// void print(QPagedPaintDevice *)
func (this *QTextDocument) Print(printer *QPagedPaintDevice /*777 QPagedPaintDevice **/) {
	var convArg0 = printer.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument5printEP17QPagedPaintDevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:216
// index:0
// Public
// QVariant resource(int, const QUrl &)
func (this *QTextDocument) Resource(type_ int, name *qtcore.QUrl) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument8resourceEiRK4QUrl", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), type_, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:217
// index:0
// Public
// void addResource(int, const QUrl &, const QVariant &)
func (this *QTextDocument) AddResource(type_ int, name *qtcore.QUrl, resource *qtcore.QVariant) {
	var convArg1 = name.GetCthis()
	var convArg2 = resource.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:221
// index:0
// Public
// void markContentsDirty(int, int)
func (this *QTextDocument) MarkContentsDirty(from int, length int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17markContentsDirtyEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), from, length)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:223
// index:0
// Public
// void setUseDesignMetrics(bool)
func (this *QTextDocument) SetUseDesignMetrics(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument19setUseDesignMetricsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:224
// index:0
// Public
// bool useDesignMetrics()
func (this *QTextDocument) UseDesignMetrics() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument16useDesignMetricsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:226
// index:0
// Public
// void drawContents(QPainter *, const QRectF &)
func (this *QTextDocument) DrawContents(painter *QPainter /*777 QPainter **/, rect *qtcore.QRectF) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:228
// index:0
// Public
// void setTextWidth(qreal)
func (this *QTextDocument) SetTextWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12setTextWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:229
// index:0
// Public
// qreal textWidth()
func (this *QTextDocument) TextWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9textWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextdocument.h:231
// index:0
// Public
// qreal idealWidth()
func (this *QTextDocument) IdealWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10idealWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextdocument.h:233
// index:0
// Public
// qreal indentWidth()
func (this *QTextDocument) IndentWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument11indentWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextdocument.h:234
// index:0
// Public
// void setIndentWidth(qreal)
func (this *QTextDocument) SetIndentWidth(width float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14setIndentWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:236
// index:0
// Public
// qreal documentMargin()
func (this *QTextDocument) DocumentMargin() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument14documentMarginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextdocument.h:237
// index:0
// Public
// void setDocumentMargin(qreal)
func (this *QTextDocument) SetDocumentMargin(margin float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17setDocumentMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), margin)
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
func (this *QTextDocument) Size() *qtcore.QSizeF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:242
// index:0
// Public
// int blockCount()
func (this *QTextDocument) BlockCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument10blockCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextdocument.h:243
// index:0
// Public
// int lineCount()
func (this *QTextDocument) LineCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument9lineCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextdocument.h:244
// index:0
// Public
// int characterCount()
func (this *QTextDocument) CharacterCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument14characterCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextdocument.h:247
// index:0
// Public
// void setDefaultStyleSheet(const QString &)
func (this *QTextDocument) SetDefaultStyleSheet(sheet *qtcore.QString) {
	var convArg0 = sheet.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument20setDefaultStyleSheetERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:248
// index:0
// Public
// QString defaultStyleSheet()
func (this *QTextDocument) DefaultStyleSheet() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17defaultStyleSheetEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:251
// index:0
// Public
// void undo(QTextCursor *)
func (this *QTextDocument) Undo(cursor *QTextCursor /*777 QTextCursor **/) {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument4undoEP11QTextCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
// void redo(QTextCursor *)
func (this *QTextDocument) Redo(cursor *QTextCursor /*777 QTextCursor **/) {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument4redoEP11QTextCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
// void clearUndoRedoStacks(QTextDocument::Stacks)
func (this *QTextDocument) ClearUndoRedoStacks(historyToClear int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument19clearUndoRedoStacksENS_6StacksE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), historyToClear)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:261
// index:0
// Public
// int maximumBlockCount()
func (this *QTextDocument) MaximumBlockCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17maximumBlockCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextdocument.h:262
// index:0
// Public
// void setMaximumBlockCount(int)
func (this *QTextDocument) SetMaximumBlockCount(maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument20setMaximumBlockCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:264
// index:0
// Public
// QTextOption defaultTextOption()
func (this *QTextDocument) DefaultTextOption() *QTextOption /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument17defaultTextOptionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:265
// index:0
// Public
// void setDefaultTextOption(const QTextOption &)
func (this *QTextDocument) SetDefaultTextOption(option *QTextOption) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument20setDefaultTextOptionERK11QTextOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:267
// index:0
// Public
// QUrl baseUrl()
func (this *QTextDocument) BaseUrl() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument7baseUrlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:268
// index:0
// Public
// void setBaseUrl(const QUrl &)
func (this *QTextDocument) SetBaseUrl(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument10setBaseUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:270
// index:0
// Public
// Qt::CursorMoveStyle defaultCursorMoveStyle()
func (this *QTextDocument) DefaultCursorMoveStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextDocument22defaultCursorMoveStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:271
// index:0
// Public
// void setDefaultCursorMoveStyle(Qt::CursorMoveStyle)
func (this *QTextDocument) SetDefaultCursorMoveStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument25setDefaultCursorMoveStyleEN2Qt15CursorMoveStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:274
// index:0
// Public
// void contentsChange(int, int, int)
func (this *QTextDocument) ContentsChange(from int, charsRemoved int, charsAdded int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14contentsChangeEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), from, charsRemoved, charsAdded)
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
// void undoAvailable(bool)
func (this *QTextDocument) UndoAvailable(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument13undoAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:277
// index:0
// Public
// void redoAvailable(bool)
func (this *QTextDocument) RedoAvailable(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument13redoAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
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
// void modificationChanged(bool)
func (this *QTextDocument) ModificationChanged(m bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument19modificationChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), m)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:280
// index:0
// Public
// void cursorPositionChanged(const QTextCursor &)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument17blockCountChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newBlockCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:282
// index:0
// Public
// void baseUrlChanged(const QUrl &)
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
// void appendUndoItem(QAbstractUndoItem *)
func (this *QTextDocument) AppendUndoItem(arg0 *QAbstractUndoItem /*777 QAbstractUndoItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:289
// index:0
// Public
// void setModified(bool)
func (this *QTextDocument) SetModified(m bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument11setModifiedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), m)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:292
// index:0
// Protected virtual
// QTextObject * createObject(const QTextFormat &)
func (this *QTextDocument) CreateObject(f *QTextFormat) *QTextObject /*777 QTextObject **/ {
	var convArg0 = f.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12createObjectERK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:293
// index:0
// Protected virtual
// QVariant loadResource(int, const QUrl &)
func (this *QTextDocument) LoadResource(type_ int, name *qtcore.QUrl) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextDocument12loadResourceEiRK4QUrl", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), type_, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QTextDocument__MetaInformation = int

const QTextDocument__DocumentTitle QTextDocument__MetaInformation = 0
const QTextDocument__DocumentUrl QTextDocument__MetaInformation = 1

type QTextDocument__FindFlag = int

const QTextDocument__FindBackward QTextDocument__FindFlag = 1
const QTextDocument__FindCaseSensitively QTextDocument__FindFlag = 2
const QTextDocument__FindWholeWords QTextDocument__FindFlag = 4

type QTextDocument__ResourceType = int

const QTextDocument__HtmlResource QTextDocument__ResourceType = 1
const QTextDocument__ImageResource QTextDocument__ResourceType = 2
const QTextDocument__StyleSheetResource QTextDocument__ResourceType = 3
const QTextDocument__UserResource QTextDocument__ResourceType = 100

type QTextDocument__Stacks = int

const QTextDocument__UndoStack QTextDocument__Stacks = 1
const QTextDocument__RedoStack QTextDocument__Stacks = 2
const QTextDocument__UndoAndRedoStacks QTextDocument__Stacks = 3

//  body block end
