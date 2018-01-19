//  header block begin
// /usr/include/qt/QtGui/qtextcursor.h
// #include <qtextcursor.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QTextCursor struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextcursor.h:69
// index:0
// void QTextCursor()
func NewQTextCursor() *QTextCursor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextCursor{cthis}
}

// /usr/include/qt/QtGui/qtextcursor.h:70
// index:1
// void QTextCursor(class QTextDocument *)
func NewQTextCursor_1(document unsafe.Pointer) *QTextCursor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, document)
	gopp.ErrPrint(err, rv)
	return &QTextCursor{cthis}
}

// /usr/include/qt/QtGui/qtextcursor.h:71
// index:2
// void QTextCursor(class QTextDocumentPrivate *, int)
func NewQTextCursor_2(p unsafe.Pointer, pos int) *QTextCursor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorC2EP20QTextDocumentPrivatei", ffiqt.FFI_TYPE_VOID, cthis, p, &pos)
	gopp.ErrPrint(err, rv)
	return &QTextCursor{cthis}
}

// /usr/include/qt/QtGui/qtextcursor.h:72
// index:3
// void QTextCursor(class QTextCursorPrivate *)
func NewQTextCursor_3(d unsafe.Pointer) *QTextCursor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorC2EP18QTextCursorPrivate", ffiqt.FFI_TYPE_VOID, cthis, d)
	gopp.ErrPrint(err, rv)
	return &QTextCursor{cthis}
}

// /usr/include/qt/QtGui/qtextcursor.h:73
// index:4
// void QTextCursor(class QTextFrame *)
func NewQTextCursor_4(frame unsafe.Pointer) *QTextCursor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorC2EP10QTextFrame", ffiqt.FFI_TYPE_VOID, cthis, frame)
	gopp.ErrPrint(err, rv)
	return &QTextCursor{cthis}
}

// /usr/include/qt/QtGui/qtextcursor.h:74
// index:5
// void QTextCursor(const class QTextBlock &)
func NewQTextCursor_5(block unsafe.Pointer) *QTextCursor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorC2ERK10QTextBlock", ffiqt.FFI_TYPE_VOID, cthis, block)
	gopp.ErrPrint(err, rv)
	return &QTextCursor{cthis}
}

// /usr/include/qt/QtGui/qtextcursor.h:80
// index:0
// void ~QTextCursor()
func DeleteQTextCursor(*QTextCursor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:82
// index:0
// inline
// void swap(class QTextCursor &)
func (this *QTextCursor) Swap(other unsafe.Pointer) {
	// 0: (, QTextCursor & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:84
// index:0
// bool isNull()
func (this *QTextCursor) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor6isNullEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:91
// index:0
// void setPosition(int, enum QTextCursor::MoveMode)
func (this *QTextCursor) SetPosition(pos int, mode int) {
	// 0: (, int pos, QTextCursor::MoveMode mode), (&pos, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11setPositionEiNS_8MoveModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &pos, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:92
// index:0
// int position()
func (this *QTextCursor) Position() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor8positionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:93
// index:0
// int positionInBlock()
func (this *QTextCursor) PositionInBlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor15positionInBlockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:95
// index:0
// int anchor()
func (this *QTextCursor) Anchor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor6anchorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:97
// index:0
// void insertText(const class QString &)
func (this *QTextCursor) InsertText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10insertTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:98
// index:1
// void insertText(const class QString &, const class QTextCharFormat &)
func (this *QTextCursor) InsertText_1(text unsafe.Pointer, format unsafe.Pointer) {
	// 1: (, const QString & text, const QTextCharFormat & format), (text, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10insertTextERK7QStringRK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.cthis, text, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:131
// index:0
// bool movePosition(enum QTextCursor::MoveOperation, enum QTextCursor::MoveMode, int)
func (this *QTextCursor) MovePosition(op int, arg1 int, n int) {
	// 0: (, QTextCursor::MoveOperation op, QTextCursor::MoveMode arg1, int n), (&op, &arg1, &n)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor12movePositionENS_13MoveOperationENS_8MoveModeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &op, &arg1, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:133
// index:0
// bool visualNavigation()
func (this *QTextCursor) VisualNavigation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor16visualNavigationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:134
// index:0
// void setVisualNavigation(_Bool)
func (this *QTextCursor) SetVisualNavigation(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor19setVisualNavigationEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:136
// index:0
// void setVerticalMovementX(int)
func (this *QTextCursor) SetVerticalMovementX(x int) {
	// 0: (, int x), (&x)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor20setVerticalMovementXEi", ffiqt.FFI_TYPE_VOID, this.cthis, &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:137
// index:0
// int verticalMovementX()
func (this *QTextCursor) VerticalMovementX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor17verticalMovementXEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:139
// index:0
// void setKeepPositionOnInsert(_Bool)
func (this *QTextCursor) SetKeepPositionOnInsert(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor23setKeepPositionOnInsertEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:140
// index:0
// bool keepPositionOnInsert()
func (this *QTextCursor) KeepPositionOnInsert() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor20keepPositionOnInsertEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:142
// index:0
// void deleteChar()
func (this *QTextCursor) DeleteChar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10deleteCharEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:143
// index:0
// void deletePreviousChar()
func (this *QTextCursor) DeletePreviousChar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor18deletePreviousCharEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:151
// index:0
// void select(enum QTextCursor::SelectionType)
func (this *QTextCursor) Select(selection int) {
	// 0: (, QTextCursor::SelectionType selection), (&selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor6selectENS_13SelectionTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:153
// index:0
// bool hasSelection()
func (this *QTextCursor) HasSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12hasSelectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:154
// index:0
// bool hasComplexSelection()
func (this *QTextCursor) HasComplexSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor19hasComplexSelectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:155
// index:0
// void removeSelectedText()
func (this *QTextCursor) RemoveSelectedText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor18removeSelectedTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:156
// index:0
// void clearSelection()
func (this *QTextCursor) ClearSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor14clearSelectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:157
// index:0
// int selectionStart()
func (this *QTextCursor) SelectionStart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor14selectionStartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:158
// index:0
// int selectionEnd()
func (this *QTextCursor) SelectionEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12selectionEndEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:160
// index:0
// QString selectedText()
func (this *QTextCursor) SelectedText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12selectedTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:161
// index:0
// QTextDocumentFragment selection()
func (this *QTextCursor) Selection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor9selectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:162
// index:0
// void selectedTableCells(int *, int *, int *, int *)
func (this *QTextCursor) SelectedTableCells(firstRow unsafe.Pointer, numRows unsafe.Pointer, firstColumn unsafe.Pointer, numColumns unsafe.Pointer) {
	// 0: (, int * firstRow, int * numRows, int * firstColumn, int * numColumns), (firstRow, numRows, firstColumn, numColumns)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor18selectedTableCellsEPiS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, firstRow, numRows, firstColumn, numColumns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:164
// index:0
// QTextBlock block()
func (this *QTextCursor) Block() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor5blockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:166
// index:0
// QTextCharFormat charFormat()
func (this *QTextCursor) CharFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor10charFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:167
// index:0
// void setCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) SetCharFormat(format unsafe.Pointer) {
	// 0: (, const QTextCharFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor13setCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:168
// index:0
// void mergeCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) MergeCharFormat(modifier unsafe.Pointer) {
	// 0: (, const QTextCharFormat & modifier), (modifier)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor15mergeCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.cthis, modifier)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:170
// index:0
// QTextBlockFormat blockFormat()
func (this *QTextCursor) BlockFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor11blockFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:171
// index:0
// void setBlockFormat(const class QTextBlockFormat &)
func (this *QTextCursor) SetBlockFormat(format unsafe.Pointer) {
	// 0: (, const QTextBlockFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor14setBlockFormatERK16QTextBlockFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:172
// index:0
// void mergeBlockFormat(const class QTextBlockFormat &)
func (this *QTextCursor) MergeBlockFormat(modifier unsafe.Pointer) {
	// 0: (, const QTextBlockFormat & modifier), (modifier)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor16mergeBlockFormatERK16QTextBlockFormat", ffiqt.FFI_TYPE_VOID, this.cthis, modifier)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:174
// index:0
// QTextCharFormat blockCharFormat()
func (this *QTextCursor) BlockCharFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor15blockCharFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:175
// index:0
// void setBlockCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) SetBlockCharFormat(format unsafe.Pointer) {
	// 0: (, const QTextCharFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor18setBlockCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:176
// index:0
// void mergeBlockCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) MergeBlockCharFormat(modifier unsafe.Pointer) {
	// 0: (, const QTextCharFormat & modifier), (modifier)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor20mergeBlockCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.cthis, modifier)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:178
// index:0
// bool atBlockStart()
func (this *QTextCursor) AtBlockStart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12atBlockStartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:179
// index:0
// bool atBlockEnd()
func (this *QTextCursor) AtBlockEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor10atBlockEndEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:180
// index:0
// bool atStart()
func (this *QTextCursor) AtStart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor7atStartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:181
// index:0
// bool atEnd()
func (this *QTextCursor) AtEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor5atEndEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:183
// index:0
// void insertBlock()
func (this *QTextCursor) InsertBlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertBlockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:184
// index:1
// void insertBlock(const class QTextBlockFormat &)
func (this *QTextCursor) InsertBlock_1(format unsafe.Pointer) {
	// 1: (, const QTextBlockFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertBlockERK16QTextBlockFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:185
// index:2
// void insertBlock(const class QTextBlockFormat &, const class QTextCharFormat &)
func (this *QTextCursor) InsertBlock_2(format unsafe.Pointer, charFormat unsafe.Pointer) {
	// 2: (, const QTextBlockFormat & format, const QTextCharFormat & charFormat), (format, charFormat)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertBlockERK16QTextBlockFormatRK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format, charFormat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:187
// index:0
// QTextList * insertList(const class QTextListFormat &)
func (this *QTextCursor) InsertList(format unsafe.Pointer) {
	// 0: (, const QTextListFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10insertListERK15QTextListFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:188
// index:1
// QTextList * insertList(class QTextListFormat::Style)
func (this *QTextCursor) InsertList_1(style int) {
	// 1: (, QTextListFormat::Style style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10insertListEN15QTextListFormat5StyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:190
// index:0
// QTextList * createList(const class QTextListFormat &)
func (this *QTextCursor) CreateList(format unsafe.Pointer) {
	// 0: (, const QTextListFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10createListERK15QTextListFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:191
// index:1
// QTextList * createList(class QTextListFormat::Style)
func (this *QTextCursor) CreateList_1(style int) {
	// 1: (, QTextListFormat::Style style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10createListEN15QTextListFormat5StyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:192
// index:0
// QTextList * currentList()
func (this *QTextCursor) CurrentList() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor11currentListEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:194
// index:0
// QTextTable * insertTable(int, int, const class QTextTableFormat &)
func (this *QTextCursor) InsertTable(rows int, cols int, format unsafe.Pointer) {
	// 0: (, int rows, int cols, const QTextTableFormat & format), (&rows, &cols, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertTableEiiRK16QTextTableFormat", ffiqt.FFI_TYPE_VOID, this.cthis, &rows, &cols, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:195
// index:1
// QTextTable * insertTable(int, int)
func (this *QTextCursor) InsertTable_1(rows int, cols int) {
	// 1: (, int rows, int cols), (&rows, &cols)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertTableEii", ffiqt.FFI_TYPE_VOID, this.cthis, &rows, &cols)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:196
// index:0
// QTextTable * currentTable()
func (this *QTextCursor) CurrentTable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12currentTableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:198
// index:0
// QTextFrame * insertFrame(const class QTextFrameFormat &)
func (this *QTextCursor) InsertFrame(format unsafe.Pointer) {
	// 0: (, const QTextFrameFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertFrameERK16QTextFrameFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:199
// index:0
// QTextFrame * currentFrame()
func (this *QTextCursor) CurrentFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12currentFrameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:201
// index:0
// void insertFragment(const class QTextDocumentFragment &)
func (this *QTextCursor) InsertFragment(fragment unsafe.Pointer) {
	// 0: (, const QTextDocumentFragment & fragment), (fragment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor14insertFragmentERK21QTextDocumentFragment", ffiqt.FFI_TYPE_VOID, this.cthis, fragment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:204
// index:0
// void insertHtml(const class QString &)
func (this *QTextCursor) InsertHtml(html unsafe.Pointer) {
	// 0: (, const QString & html), (html)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10insertHtmlERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, html)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:207
// index:0
// void insertImage(const class QTextImageFormat &, class QTextFrameFormat::Position)
func (this *QTextCursor) InsertImage(format unsafe.Pointer, alignment int) {
	// 0: (, const QTextImageFormat & format, QTextFrameFormat::Position alignment), (format, &alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK16QTextImageFormatN16QTextFrameFormat8PositionE", ffiqt.FFI_TYPE_VOID, this.cthis, format, &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:208
// index:1
// void insertImage(const class QTextImageFormat &)
func (this *QTextCursor) InsertImage_1(format unsafe.Pointer) {
	// 1: (, const QTextImageFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK16QTextImageFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:209
// index:2
// void insertImage(const class QString &)
func (this *QTextCursor) InsertImage_2(name unsafe.Pointer) {
	// 2: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:210
// index:3
// void insertImage(const class QImage &, const class QString &)
func (this *QTextCursor) InsertImage_3(image unsafe.Pointer, name unsafe.Pointer) {
	// 3: (, const QImage & image, const QString & name), (image, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK6QImageRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, image, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:212
// index:0
// void beginEditBlock()
func (this *QTextCursor) BeginEditBlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor14beginEditBlockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:213
// index:0
// void joinPreviousEditBlock()
func (this *QTextCursor) JoinPreviousEditBlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor21joinPreviousEditBlockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:214
// index:0
// void endEditBlock()
func (this *QTextCursor) EndEditBlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor12endEditBlockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:223
// index:0
// bool isCopyOf(const class QTextCursor &)
func (this *QTextCursor) IsCopyOf(other unsafe.Pointer) {
	// 0: (, const QTextCursor & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor8isCopyOfERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:225
// index:0
// int blockNumber()
func (this *QTextCursor) BlockNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor11blockNumberEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:226
// index:0
// int columnNumber()
func (this *QTextCursor) ColumnNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12columnNumberEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:228
// index:0
// QTextDocument * document()
func (this *QTextCursor) Document() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor8documentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
