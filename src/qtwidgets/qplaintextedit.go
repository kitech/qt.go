//  header block begin
// /usr/include/qt/QtWidgets/qplaintextedit.h
// #include <qplaintextedit.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 128
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
	*QAbstractScrollArea
}

func (this *QPlainTextEdit) GetCthis() unsafe.Pointer {
	return this.QAbstractScrollArea.GetCthis()
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:66
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPlainTextEdit) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:94
// index:0
// void QPlainTextEdit(class QWidget *)
func NewQPlainTextEdit(parent unsafe.Pointer) *QPlainTextEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPlainTextEditFromPointer(cthis)
	return gothis
}
func NewQPlainTextEditFromPointer(cthis unsafe.Pointer) *QPlainTextEdit {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QPlainTextEdit{bcthis0}
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:95
// index:1
// void QPlainTextEdit(const class QString &, class QWidget *)
func NewQPlainTextEdit_1(text unsafe.Pointer, parent unsafe.Pointer) *QPlainTextEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEditC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPlainTextEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:270
// index:2
// void QPlainTextEdit(class QPlainTextEditPrivate &, class QWidget *)
func NewQPlainTextEdit_2(dd unsafe.Pointer, parent unsafe.Pointer) *QPlainTextEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEditC2ER21QPlainTextEditPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPlainTextEditFromPointer(cthis)
	return gothis
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
	// 0: (, document QTextDocument *), (document)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11setDocumentEP13QTextDocument", ffiqt.FFI_TYPE_VOID, this.GetCthis(), document)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:99
// index:0
// QTextDocument * document()
func (this *QPlainTextEdit) Document() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit8documentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:101
// index:0
// void setPlaceholderText(const class QString &)
func (this *QPlainTextEdit) SetPlaceholderText(placeholderText unsafe.Pointer) {
	// 0: (, placeholderText const QString &), (placeholderText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setPlaceholderTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), placeholderText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:102
// index:0
// QString placeholderText()
func (this *QPlainTextEdit) PlaceholderText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15placeholderTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:104
// index:0
// void setTextCursor(const class QTextCursor &)
func (this *QPlainTextEdit) SetTextCursor(cursor unsafe.Pointer) {
	// 0: (, cursor const QTextCursor &), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13setTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:105
// index:0
// QTextCursor textCursor()
func (this *QPlainTextEdit) TextCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10textCursorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:107
// index:0
// bool isReadOnly()
func (this *QPlainTextEdit) IsReadOnly() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10isReadOnlyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:108
// index:0
// void setReadOnly(_Bool)
func (this *QPlainTextEdit) SetReadOnly(ro bool) {
	// 0: (, ro bool), (&ro)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11setReadOnlyEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ro)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:110
// index:0
// void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QPlainTextEdit) SetTextInteractionFlags(flags int) {
	// 0: (, QFlags<Qt::TextInteractionFlag> flags), (&flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:111
// index:0
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QPlainTextEdit) TextInteractionFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit20textInteractionFlagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:113
// index:0
// void mergeCurrentCharFormat(const class QTextCharFormat &)
func (this *QPlainTextEdit) MergeCurrentCharFormat(modifier unsafe.Pointer) {
	// 0: (, modifier const QTextCharFormat &), (modifier)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), modifier)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:114
// index:0
// void setCurrentCharFormat(const class QTextCharFormat &)
func (this *QPlainTextEdit) SetCurrentCharFormat(format unsafe.Pointer) {
	// 0: (, format const QTextCharFormat &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:115
// index:0
// QTextCharFormat currentCharFormat()
func (this *QPlainTextEdit) CurrentCharFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17currentCharFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:117
// index:0
// bool tabChangesFocus()
func (this *QPlainTextEdit) TabChangesFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15tabChangesFocusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:118
// index:0
// void setTabChangesFocus(_Bool)
func (this *QPlainTextEdit) SetTabChangesFocus(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setTabChangesFocusEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:120
// index:0
// inline
// void setDocumentTitle(const class QString &)
func (this *QPlainTextEdit) SetDocumentTitle(title unsafe.Pointer) {
	// 0: (, title const QString &), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16setDocumentTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:122
// index:0
// inline
// QString documentTitle()
func (this *QPlainTextEdit) DocumentTitle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit13documentTitleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:125
// index:0
// inline
// bool isUndoRedoEnabled()
func (this *QPlainTextEdit) IsUndoRedoEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17isUndoRedoEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:127
// index:0
// inline
// void setUndoRedoEnabled(_Bool)
func (this *QPlainTextEdit) SetUndoRedoEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setUndoRedoEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:130
// index:0
// inline
// void setMaximumBlockCount(int)
func (this *QPlainTextEdit) SetMaximumBlockCount(maximum int) {
	// 0: (, maximum int), (&maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit20setMaximumBlockCountEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:132
// index:0
// inline
// int maximumBlockCount()
func (this *QPlainTextEdit) MaximumBlockCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17maximumBlockCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:136
// index:0
// QPlainTextEdit::LineWrapMode lineWrapMode()
func (this *QPlainTextEdit) LineWrapMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit12lineWrapModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:137
// index:0
// void setLineWrapMode(enum QPlainTextEdit::LineWrapMode)
func (this *QPlainTextEdit) SetLineWrapMode(mode int) {
	// 0: (, mode QPlainTextEdit::LineWrapMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15setLineWrapModeENS_12LineWrapModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:139
// index:0
// QTextOption::WrapMode wordWrapMode()
func (this *QPlainTextEdit) WordWrapMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit12wordWrapModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:140
// index:0
// void setWordWrapMode(class QTextOption::WrapMode)
func (this *QPlainTextEdit) SetWordWrapMode(policy int) {
	// 0: (, policy QTextOption::WrapMode), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15setWordWrapModeEN11QTextOption8WrapModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:142
// index:0
// void setBackgroundVisible(_Bool)
func (this *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit20setBackgroundVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:143
// index:0
// bool backgroundVisible()
func (this *QPlainTextEdit) BackgroundVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17backgroundVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:145
// index:0
// void setCenterOnScroll(_Bool)
func (this *QPlainTextEdit) SetCenterOnScroll(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit17setCenterOnScrollEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:146
// index:0
// bool centerOnScroll()
func (this *QPlainTextEdit) CenterOnScroll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit14centerOnScrollEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:148
// index:0
// bool find(const class QString &, class QTextDocument::FindFlags)
func (this *QPlainTextEdit) Find(exp unsafe.Pointer, options int) {
	// 0: (, exp const QString &, QFlags<QTextDocument::FindFlag> options), (exp, &options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4findERK7QString6QFlagsIN13QTextDocument8FindFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), exp, &options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:150
// index:1
// bool find(const class QRegExp &, class QTextDocument::FindFlags)
func (this *QPlainTextEdit) Find_1(exp unsafe.Pointer, options int) {
	// 1: (, exp const QRegExp &, QFlags<QTextDocument::FindFlag> options), (exp, &options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4findERK7QRegExp6QFlagsIN13QTextDocument8FindFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), exp, &options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:153
// index:0
// inline
// QString toPlainText()
func (this *QPlainTextEdit) ToPlainText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit11toPlainTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:156
// index:0
// void ensureCursorVisible()
func (this *QPlainTextEdit) EnsureCursorVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit19ensureCursorVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:158
// index:0
// virtual
// QVariant loadResource(int, const class QUrl &)
func (this *QPlainTextEdit) LoadResource(type_ int, name unsafe.Pointer) {
	// 0: (, type int, name const QUrl &), (&type_, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12loadResourceEiRK4QUrl", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &type_, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:160
// index:0
// QMenu * createStandardContextMenu()
func (this *QPlainTextEdit) CreateStandardContextMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit25createStandardContextMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:161
// index:1
// QMenu * createStandardContextMenu(const class QPoint &)
func (this *QPlainTextEdit) CreateStandardContextMenu_1(position unsafe.Pointer) {
	// 1: (, position const QPoint &), (position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:164
// index:0
// QTextCursor cursorForPosition(const class QPoint &)
func (this *QPlainTextEdit) CursorForPosition(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17cursorForPositionERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:165
// index:0
// QRect cursorRect(const class QTextCursor &)
func (this *QPlainTextEdit) CursorRect(cursor unsafe.Pointer) {
	// 0: (, cursor const QTextCursor &), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10cursorRectERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:166
// index:1
// QRect cursorRect()
func (this *QPlainTextEdit) CursorRect_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10cursorRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:168
// index:0
// QString anchorAt(const class QPoint &)
func (this *QPlainTextEdit) AnchorAt(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit8anchorAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:170
// index:0
// bool overwriteMode()
func (this *QPlainTextEdit) OverwriteMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit13overwriteModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:171
// index:0
// void setOverwriteMode(_Bool)
func (this *QPlainTextEdit) SetOverwriteMode(overwrite bool) {
	// 0: (, overwrite bool), (&overwrite)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16setOverwriteModeEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &overwrite)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:174
// index:0
// int tabStopWidth()
func (this *QPlainTextEdit) TabStopWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit12tabStopWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:175
// index:0
// void setTabStopWidth(int)
func (this *QPlainTextEdit) SetTabStopWidth(width int) {
	// 0: (, width int), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15setTabStopWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:178
// index:0
// qreal tabStopDistance()
func (this *QPlainTextEdit) TabStopDistance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15tabStopDistanceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:179
// index:0
// void setTabStopDistance(qreal)
func (this *QPlainTextEdit) SetTabStopDistance(distance float64) {
	// 0: (, distance qreal), (&distance)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setTabStopDistanceEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &distance)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:181
// index:0
// int cursorWidth()
func (this *QPlainTextEdit) CursorWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit11cursorWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:182
// index:0
// void setCursorWidth(int)
func (this *QPlainTextEdit) SetCursorWidth(width int) {
	// 0: (, width int), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit14setCursorWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:185
// index:0
// QList<QTextEdit::ExtraSelection> extraSelections()
func (this *QPlainTextEdit) ExtraSelections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15extraSelectionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:187
// index:0
// void moveCursor(class QTextCursor::MoveOperation, class QTextCursor::MoveMode)
func (this *QPlainTextEdit) MoveCursor(operation int, mode int) {
	// 0: (, operation QTextCursor::MoveOperation, mode QTextCursor::MoveMode), (&operation, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10moveCursorEN11QTextCursor13MoveOperationENS0_8MoveModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &operation, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:189
// index:0
// bool canPaste()
func (this *QPlainTextEdit) CanPaste() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit8canPasteEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:191
// index:0
// void print(class QPagedPaintDevice *)
func (this *QPlainTextEdit) Print(printer unsafe.Pointer) {
	// 0: (, printer QPagedPaintDevice *), (printer)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit5printEP17QPagedPaintDevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), printer)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:193
// index:0
// int blockCount()
func (this *QPlainTextEdit) BlockCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10blockCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:194
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QPlainTextEdit) InputMethodQuery(property int) {
	// 0: (, property Qt::InputMethodQuery), (&property)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &property)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:195
// index:1
// QVariant inputMethodQuery(Qt::InputMethodQuery, class QVariant)
func (this *QPlainTextEdit) InputMethodQuery_1(query int, argument unsafe.Pointer) {
	// 1: (, query Qt::InputMethodQuery, argument QVariant), (&query, argument)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &query, argument)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:199
// index:0
// void setPlainText(const class QString &)
func (this *QPlainTextEdit) SetPlainText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12setPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:202
// index:0
// void cut()
func (this *QPlainTextEdit) Cut() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit3cutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:203
// index:0
// void copy()
func (this *QPlainTextEdit) Copy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4copyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:204
// index:0
// void paste()
func (this *QPlainTextEdit) Paste() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit5pasteEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:207
// index:0
// void undo()
func (this *QPlainTextEdit) Undo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4undoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:208
// index:0
// void redo()
func (this *QPlainTextEdit) Redo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4redoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:210
// index:0
// void clear()
func (this *QPlainTextEdit) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:211
// index:0
// void selectAll()
func (this *QPlainTextEdit) SelectAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit9selectAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:213
// index:0
// void insertPlainText(const class QString &)
func (this *QPlainTextEdit) InsertPlainText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15insertPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:215
// index:0
// void appendPlainText(const class QString &)
func (this *QPlainTextEdit) AppendPlainText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15appendPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:216
// index:0
// void appendHtml(const class QString &)
func (this *QPlainTextEdit) AppendHtml(html unsafe.Pointer) {
	// 0: (, html const QString &), (html)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10appendHtmlERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), html)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:218
// index:0
// void centerCursor()
func (this *QPlainTextEdit) CenterCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12centerCursorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:220
// index:0
// void zoomIn(int)
func (this *QPlainTextEdit) ZoomIn(range_ int) {
	// 0: (, range int), (&range_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit6zoomInEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &range_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:221
// index:0
// void zoomOut(int)
func (this *QPlainTextEdit) ZoomOut(range_ int) {
	// 0: (, range int), (&range_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit7zoomOutEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &range_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:224
// index:0
// void textChanged()
func (this *QPlainTextEdit) TextChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11textChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:225
// index:0
// void undoAvailable(_Bool)
func (this *QPlainTextEdit) UndoAvailable(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13undoAvailableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:226
// index:0
// void redoAvailable(_Bool)
func (this *QPlainTextEdit) RedoAvailable(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13redoAvailableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:227
// index:0
// void copyAvailable(_Bool)
func (this *QPlainTextEdit) CopyAvailable(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13copyAvailableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:228
// index:0
// void selectionChanged()
func (this *QPlainTextEdit) SelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16selectionChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:229
// index:0
// void cursorPositionChanged()
func (this *QPlainTextEdit) CursorPositionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit21cursorPositionChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:231
// index:0
// void updateRequest(const class QRect &, int)
func (this *QPlainTextEdit) UpdateRequest(rect unsafe.Pointer, dy int) {
	// 0: (, rect const QRect &, dy int), (rect, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13updateRequestERK5QRecti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:232
// index:0
// void blockCountChanged(int)
func (this *QPlainTextEdit) BlockCountChanged(newBlockCount int) {
	// 0: (, newBlockCount int), (&newBlockCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit17blockCountChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &newBlockCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:233
// index:0
// void modificationChanged(_Bool)
func (this *QPlainTextEdit) ModificationChanged(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit19modificationChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:236
// index:0
// virtual
// bool event(class QEvent *)
func (this *QPlainTextEdit) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:237
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QPlainTextEdit) TimerEvent(e unsafe.Pointer) {
	// 0: (, e QTimerEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:238
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QPlainTextEdit) KeyPressEvent(e unsafe.Pointer) {
	// 0: (, e QKeyEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:239
// index:0
// virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QPlainTextEdit) KeyReleaseEvent(e unsafe.Pointer) {
	// 0: (, e QKeyEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:240
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QPlainTextEdit) ResizeEvent(e unsafe.Pointer) {
	// 0: (, e QResizeEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:241
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QPlainTextEdit) PaintEvent(e unsafe.Pointer) {
	// 0: (, e QPaintEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:242
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QPlainTextEdit) MousePressEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:243
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QPlainTextEdit) MouseMoveEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:244
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QPlainTextEdit) MouseReleaseEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:245
// index:0
// virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QPlainTextEdit) MouseDoubleClickEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:246
// index:0
// virtual
// bool focusNextPrevChild(_Bool)
func (this *QPlainTextEdit) FocusNextPrevChild(next bool) {
	// 0: (, next bool), (&next)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18focusNextPrevChildEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:248
// index:0
// virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QPlainTextEdit) ContextMenuEvent(e unsafe.Pointer) {
	// 0: (, e QContextMenuEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:251
// index:0
// virtual
// void dragEnterEvent(class QDragEnterEvent *)
func (this *QPlainTextEdit) DragEnterEvent(e unsafe.Pointer) {
	// 0: (, e QDragEnterEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:252
// index:0
// virtual
// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QPlainTextEdit) DragLeaveEvent(e unsafe.Pointer) {
	// 0: (, e QDragLeaveEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:253
// index:0
// virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QPlainTextEdit) DragMoveEvent(e unsafe.Pointer) {
	// 0: (, e QDragMoveEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:254
// index:0
// virtual
// void dropEvent(class QDropEvent *)
func (this *QPlainTextEdit) DropEvent(e unsafe.Pointer) {
	// 0: (, e QDropEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:256
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QPlainTextEdit) FocusInEvent(e unsafe.Pointer) {
	// 0: (, e QFocusEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:257
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QPlainTextEdit) FocusOutEvent(e unsafe.Pointer) {
	// 0: (, e QFocusEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:258
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QPlainTextEdit) ShowEvent(arg0 unsafe.Pointer) {
	// 0: (, QShowEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:259
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QPlainTextEdit) ChangeEvent(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:261
// index:0
// virtual
// void wheelEvent(class QWheelEvent *)
func (this *QPlainTextEdit) WheelEvent(e unsafe.Pointer) {
	// 0: (, e QWheelEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:264
// index:0
// virtual
// QMimeData * createMimeDataFromSelection()
func (this *QPlainTextEdit) CreateMimeDataFromSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit27createMimeDataFromSelectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:265
// index:0
// virtual
// bool canInsertFromMimeData(const class QMimeData *)
func (this *QPlainTextEdit) CanInsertFromMimeData(source unsafe.Pointer) {
	// 0: (, source const QMimeData *), (source)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit21canInsertFromMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_VOID, this.GetCthis(), source)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:266
// index:0
// virtual
// void insertFromMimeData(const class QMimeData *)
func (this *QPlainTextEdit) InsertFromMimeData(source unsafe.Pointer) {
	// 0: (, source const QMimeData *), (source)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18insertFromMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_VOID, this.GetCthis(), source)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:268
// index:0
// virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QPlainTextEdit) InputMethodEvent(arg0 unsafe.Pointer) {
	// 0: (, QInputMethodEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:272
// index:0
// virtual
// void scrollContentsBy(int, int)
func (this *QPlainTextEdit) ScrollContentsBy(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16scrollContentsByEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:273
// index:0
// virtual
// void doSetTextCursor(const class QTextCursor &)
func (this *QPlainTextEdit) DoSetTextCursor(cursor unsafe.Pointer) {
	// 0: (, cursor const QTextCursor &), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15doSetTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:275
// index:0
// QTextBlock firstVisibleBlock()
func (this *QPlainTextEdit) FirstVisibleBlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17firstVisibleBlockEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:276
// index:0
// QPointF contentOffset()
func (this *QPlainTextEdit) ContentOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit13contentOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:277
// index:0
// QRectF blockBoundingRect(const class QTextBlock &)
func (this *QPlainTextEdit) BlockBoundingRect(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17blockBoundingRectERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:278
// index:0
// QRectF blockBoundingGeometry(const class QTextBlock &)
func (this *QPlainTextEdit) BlockBoundingGeometry(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit21blockBoundingGeometryERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:279
// index:0
// QAbstractTextDocumentLayout::PaintContext getPaintContext()
func (this *QPlainTextEdit) GetPaintContext() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15getPaintContextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:281
// index:0
// void zoomInF(float)
func (this *QPlainTextEdit) ZoomInF(range_ float32) {
	// 0: (, range float), (&range_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit7zoomInFEf", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &range_)
	gopp.ErrPrint(err, rv)
}

//  body block end
