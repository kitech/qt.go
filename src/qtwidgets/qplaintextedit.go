//  header block begin
// /usr/include/qt/QtWidgets/qplaintextedit.h
// #include <qplaintextedit.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 94
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QPlainTextEdit struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:66
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPlainTextEdit) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:94
// index:0
// void QPlainTextEdit(class QWidget *)
func NewQPlainTextEdit(parent unsafe.Pointer) *QPlainTextEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QPlainTextEdit{cthis}
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:95
// index:1
// void QPlainTextEdit(const class QString &, class QWidget *)
func NewQPlainTextEdit_1(text unsafe.Pointer, parent unsafe.Pointer) *QPlainTextEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEditC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	return &QPlainTextEdit{cthis}
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:96
// index:0
// virtual
// void ~QPlainTextEdit()
func DeleteQPlainTextEdit(*QPlainTextEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:98
// index:0
// void setDocument(class QTextDocument *)
func (this *QPlainTextEdit) SetDocument(document unsafe.Pointer) {
	// 0: (, QTextDocument * document), (document)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11setDocumentEP13QTextDocument", ffiqt.FFI_TYPE_VOID, this.cthis, document)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:99
// index:0
// QTextDocument * document()
func (this *QPlainTextEdit) Document() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit8documentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:101
// index:0
// void setPlaceholderText(const class QString &)
func (this *QPlainTextEdit) SetPlaceholderText(placeholderText unsafe.Pointer) {
	// 0: (, const QString & placeholderText), (placeholderText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setPlaceholderTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, placeholderText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:102
// index:0
// QString placeholderText()
func (this *QPlainTextEdit) PlaceholderText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15placeholderTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:104
// index:0
// void setTextCursor(const class QTextCursor &)
func (this *QPlainTextEdit) SetTextCursor(cursor unsafe.Pointer) {
	// 0: (, const QTextCursor & cursor), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13setTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.cthis, cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:105
// index:0
// QTextCursor textCursor()
func (this *QPlainTextEdit) TextCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10textCursorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:107
// index:0
// bool isReadOnly()
func (this *QPlainTextEdit) IsReadOnly() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10isReadOnlyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:108
// index:0
// void setReadOnly(_Bool)
func (this *QPlainTextEdit) SetReadOnly(ro bool) {
	// 0: (, bool ro), (&ro)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11setReadOnlyEb", ffiqt.FFI_TYPE_VOID, this.cthis, &ro)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:111
// index:0
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QPlainTextEdit) TextInteractionFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit20textInteractionFlagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:113
// index:0
// void mergeCurrentCharFormat(const class QTextCharFormat &)
func (this *QPlainTextEdit) MergeCurrentCharFormat(modifier unsafe.Pointer) {
	// 0: (, const QTextCharFormat & modifier), (modifier)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.cthis, modifier)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:114
// index:0
// void setCurrentCharFormat(const class QTextCharFormat &)
func (this *QPlainTextEdit) SetCurrentCharFormat(format unsafe.Pointer) {
	// 0: (, const QTextCharFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:115
// index:0
// QTextCharFormat currentCharFormat()
func (this *QPlainTextEdit) CurrentCharFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17currentCharFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:117
// index:0
// bool tabChangesFocus()
func (this *QPlainTextEdit) TabChangesFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15tabChangesFocusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:118
// index:0
// void setTabChangesFocus(_Bool)
func (this *QPlainTextEdit) SetTabChangesFocus(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setTabChangesFocusEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:120
// index:0
// inline
// void setDocumentTitle(const class QString &)
func (this *QPlainTextEdit) SetDocumentTitle(title unsafe.Pointer) {
	// 0: (, const QString & title), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16setDocumentTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:122
// index:0
// inline
// QString documentTitle()
func (this *QPlainTextEdit) DocumentTitle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit13documentTitleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:125
// index:0
// inline
// bool isUndoRedoEnabled()
func (this *QPlainTextEdit) IsUndoRedoEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17isUndoRedoEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:127
// index:0
// inline
// void setUndoRedoEnabled(_Bool)
func (this *QPlainTextEdit) SetUndoRedoEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setUndoRedoEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:130
// index:0
// inline
// void setMaximumBlockCount(int)
func (this *QPlainTextEdit) SetMaximumBlockCount(maximum int) {
	// 0: (, int maximum), (&maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit20setMaximumBlockCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:132
// index:0
// inline
// int maximumBlockCount()
func (this *QPlainTextEdit) MaximumBlockCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17maximumBlockCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:136
// index:0
// QPlainTextEdit::LineWrapMode lineWrapMode()
func (this *QPlainTextEdit) LineWrapMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit12lineWrapModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:137
// index:0
// void setLineWrapMode(enum QPlainTextEdit::LineWrapMode)
func (this *QPlainTextEdit) SetLineWrapMode(mode int) {
	// 0: (, QPlainTextEdit::LineWrapMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15setLineWrapModeENS_12LineWrapModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:139
// index:0
// QTextOption::WrapMode wordWrapMode()
func (this *QPlainTextEdit) WordWrapMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit12wordWrapModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:140
// index:0
// void setWordWrapMode(class QTextOption::WrapMode)
func (this *QPlainTextEdit) SetWordWrapMode(policy int) {
	// 0: (, QTextOption::WrapMode policy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15setWordWrapModeEN11QTextOption8WrapModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:142
// index:0
// void setBackgroundVisible(_Bool)
func (this *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit20setBackgroundVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:143
// index:0
// bool backgroundVisible()
func (this *QPlainTextEdit) BackgroundVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17backgroundVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:145
// index:0
// void setCenterOnScroll(_Bool)
func (this *QPlainTextEdit) SetCenterOnScroll(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit17setCenterOnScrollEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:146
// index:0
// bool centerOnScroll()
func (this *QPlainTextEdit) CenterOnScroll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit14centerOnScrollEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:153
// index:0
// inline
// QString toPlainText()
func (this *QPlainTextEdit) ToPlainText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit11toPlainTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:156
// index:0
// void ensureCursorVisible()
func (this *QPlainTextEdit) EnsureCursorVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit19ensureCursorVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:158
// index:0
// virtual
// QVariant loadResource(int, const class QUrl &)
func (this *QPlainTextEdit) LoadResource(type_ int, name unsafe.Pointer) {
	// 0: (, int type, const QUrl & name), (&type_, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12loadResourceEiRK4QUrl", ffiqt.FFI_TYPE_VOID, this.cthis, &type_, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:160
// index:0
// QMenu * createStandardContextMenu()
func (this *QPlainTextEdit) CreateStandardContextMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit25createStandardContextMenuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:161
// index:1
// QMenu * createStandardContextMenu(const class QPoint &)
func (this *QPlainTextEdit) CreateStandardContextMenu_1(position unsafe.Pointer) {
	// 1: (, const QPoint & position), (position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:164
// index:0
// QTextCursor cursorForPosition(const class QPoint &)
func (this *QPlainTextEdit) CursorForPosition(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17cursorForPositionERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:165
// index:0
// QRect cursorRect(const class QTextCursor &)
func (this *QPlainTextEdit) CursorRect(cursor unsafe.Pointer) {
	// 0: (, const QTextCursor & cursor), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10cursorRectERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.cthis, cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:166
// index:1
// QRect cursorRect()
func (this *QPlainTextEdit) CursorRect_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10cursorRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:168
// index:0
// QString anchorAt(const class QPoint &)
func (this *QPlainTextEdit) AnchorAt(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit8anchorAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:170
// index:0
// bool overwriteMode()
func (this *QPlainTextEdit) OverwriteMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit13overwriteModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:171
// index:0
// void setOverwriteMode(_Bool)
func (this *QPlainTextEdit) SetOverwriteMode(overwrite bool) {
	// 0: (, bool overwrite), (&overwrite)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16setOverwriteModeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &overwrite)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:174
// index:0
// int tabStopWidth()
func (this *QPlainTextEdit) TabStopWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit12tabStopWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:175
// index:0
// void setTabStopWidth(int)
func (this *QPlainTextEdit) SetTabStopWidth(width int) {
	// 0: (, int width), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15setTabStopWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:178
// index:0
// qreal tabStopDistance()
func (this *QPlainTextEdit) TabStopDistance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15tabStopDistanceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:179
// index:0
// void setTabStopDistance(qreal)
func (this *QPlainTextEdit) SetTabStopDistance(distance float64) {
	// 0: (, qreal distance), (&distance)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setTabStopDistanceEd", ffiqt.FFI_TYPE_VOID, this.cthis, &distance)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:181
// index:0
// int cursorWidth()
func (this *QPlainTextEdit) CursorWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit11cursorWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:182
// index:0
// void setCursorWidth(int)
func (this *QPlainTextEdit) SetCursorWidth(width int) {
	// 0: (, int width), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit14setCursorWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:185
// index:0
// QList<QTextEdit::ExtraSelection> extraSelections()
func (this *QPlainTextEdit) ExtraSelections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15extraSelectionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:187
// index:0
// void moveCursor(class QTextCursor::MoveOperation, class QTextCursor::MoveMode)
func (this *QPlainTextEdit) MoveCursor(operation int, mode int) {
	// 0: (, QTextCursor::MoveOperation operation, QTextCursor::MoveMode mode), (&operation, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10moveCursorEN11QTextCursor13MoveOperationENS0_8MoveModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &operation, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:189
// index:0
// bool canPaste()
func (this *QPlainTextEdit) CanPaste() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit8canPasteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:191
// index:0
// void print(class QPagedPaintDevice *)
func (this *QPlainTextEdit) Print(printer unsafe.Pointer) {
	// 0: (, QPagedPaintDevice * printer), (printer)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit5printEP17QPagedPaintDevice", ffiqt.FFI_TYPE_VOID, this.cthis, printer)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:193
// index:0
// int blockCount()
func (this *QPlainTextEdit) BlockCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10blockCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:194
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QPlainTextEdit) InputMethodQuery(property int) {
	// 0: (, Qt::InputMethodQuery property), (&property)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.cthis, &property)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:195
// index:1
// QVariant inputMethodQuery(Qt::InputMethodQuery, class QVariant)
func (this *QPlainTextEdit) InputMethodQuery_1(query int, argument unsafe.Pointer) {
	// 1: (, Qt::InputMethodQuery query, QVariant argument), (&query, argument)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &query, argument)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:199
// index:0
// void setPlainText(const class QString &)
func (this *QPlainTextEdit) SetPlainText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12setPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:202
// index:0
// void cut()
func (this *QPlainTextEdit) Cut() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit3cutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:203
// index:0
// void copy()
func (this *QPlainTextEdit) Copy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4copyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:204
// index:0
// void paste()
func (this *QPlainTextEdit) Paste() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit5pasteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:207
// index:0
// void undo()
func (this *QPlainTextEdit) Undo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4undoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:208
// index:0
// void redo()
func (this *QPlainTextEdit) Redo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4redoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:210
// index:0
// void clear()
func (this *QPlainTextEdit) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:211
// index:0
// void selectAll()
func (this *QPlainTextEdit) SelectAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit9selectAllEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:213
// index:0
// void insertPlainText(const class QString &)
func (this *QPlainTextEdit) InsertPlainText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15insertPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:215
// index:0
// void appendPlainText(const class QString &)
func (this *QPlainTextEdit) AppendPlainText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15appendPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:216
// index:0
// void appendHtml(const class QString &)
func (this *QPlainTextEdit) AppendHtml(html unsafe.Pointer) {
	// 0: (, const QString & html), (html)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10appendHtmlERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, html)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:218
// index:0
// void centerCursor()
func (this *QPlainTextEdit) CenterCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12centerCursorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:220
// index:0
// void zoomIn(int)
func (this *QPlainTextEdit) ZoomIn(range_ int) {
	// 0: (, int range), (&range_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit6zoomInEi", ffiqt.FFI_TYPE_VOID, this.cthis, &range_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:221
// index:0
// void zoomOut(int)
func (this *QPlainTextEdit) ZoomOut(range_ int) {
	// 0: (, int range), (&range_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit7zoomOutEi", ffiqt.FFI_TYPE_VOID, this.cthis, &range_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:224
// index:0
// void textChanged()
func (this *QPlainTextEdit) TextChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11textChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:225
// index:0
// void undoAvailable(_Bool)
func (this *QPlainTextEdit) UndoAvailable(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13undoAvailableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:226
// index:0
// void redoAvailable(_Bool)
func (this *QPlainTextEdit) RedoAvailable(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13redoAvailableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:227
// index:0
// void copyAvailable(_Bool)
func (this *QPlainTextEdit) CopyAvailable(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13copyAvailableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:228
// index:0
// void selectionChanged()
func (this *QPlainTextEdit) SelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16selectionChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:229
// index:0
// void cursorPositionChanged()
func (this *QPlainTextEdit) CursorPositionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit21cursorPositionChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:231
// index:0
// void updateRequest(const class QRect &, int)
func (this *QPlainTextEdit) UpdateRequest(rect unsafe.Pointer, dy int) {
	// 0: (, const QRect & rect, int dy), (rect, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13updateRequestERK5QRecti", ffiqt.FFI_TYPE_VOID, this.cthis, rect, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:232
// index:0
// void blockCountChanged(int)
func (this *QPlainTextEdit) BlockCountChanged(newBlockCount int) {
	// 0: (, int newBlockCount), (&newBlockCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit17blockCountChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &newBlockCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:233
// index:0
// void modificationChanged(_Bool)
func (this *QPlainTextEdit) ModificationChanged(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit19modificationChangedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
