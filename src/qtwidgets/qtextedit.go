//  header block begin
// /usr/include/qt/QtWidgets/qtextedit.h
// #include <qtextedit.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QTextEdit struct {
	*QAbstractScrollArea
}

func (this *QTextEdit) GetCthis() unsafe.Pointer {
	return this.QAbstractScrollArea.GetCthis()
}

// /usr/include/qt/QtWidgets/qtextedit.h:63
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTextEdit) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:105
// index:0
// void QTextEdit(class QWidget *)
func NewQTextEdit(parent unsafe.Pointer) *QTextEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextEditFromPointer(cthis)
	return gothis
}
func NewQTextEditFromPointer(cthis unsafe.Pointer) *QTextEdit {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QTextEdit{bcthis0}
}

// /usr/include/qt/QtWidgets/qtextedit.h:106
// index:1
// void QTextEdit(const class QString &, class QWidget *)
func NewQTextEdit_1(text unsafe.Pointer, parent unsafe.Pointer) *QTextEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEditC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtextedit.h:307
// index:2
// void QTextEdit(class QTextEditPrivate &, class QWidget *)
func NewQTextEdit_2(dd unsafe.Pointer, parent unsafe.Pointer) *QTextEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEditC2ER16QTextEditPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtextedit.h:107
// index:0
// virtual
// void ~QTextEdit()
func DeleteQTextEdit(*QTextEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:109
// index:0
// void setDocument(class QTextDocument *)
func (this *QTextEdit) SetDocument(document unsafe.Pointer) {
	// 0: (, document QTextDocument *), (document)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit11setDocumentEP13QTextDocument", ffiqt.FFI_TYPE_VOID, this.GetCthis(), document)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:110
// index:0
// QTextDocument * document()
func (this *QTextEdit) Document() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit8documentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:112
// index:0
// void setPlaceholderText(const class QString &)
func (this *QTextEdit) SetPlaceholderText(placeholderText unsafe.Pointer) {
	// 0: (, placeholderText const QString &), (placeholderText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18setPlaceholderTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), placeholderText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:113
// index:0
// QString placeholderText()
func (this *QTextEdit) PlaceholderText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit15placeholderTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:115
// index:0
// void setTextCursor(const class QTextCursor &)
func (this *QTextEdit) SetTextCursor(cursor unsafe.Pointer) {
	// 0: (, cursor const QTextCursor &), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13setTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:116
// index:0
// QTextCursor textCursor()
func (this *QTextEdit) TextCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10textCursorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:118
// index:0
// bool isReadOnly()
func (this *QTextEdit) IsReadOnly() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10isReadOnlyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:119
// index:0
// void setReadOnly(_Bool)
func (this *QTextEdit) SetReadOnly(ro bool) {
	// 0: (, ro bool), (&ro)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit11setReadOnlyEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ro)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:121
// index:0
// void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QTextEdit) SetTextInteractionFlags(flags int) {
	// 0: (, QFlags<Qt::TextInteractionFlag> flags), (&flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:122
// index:0
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QTextEdit) TextInteractionFlags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit20textInteractionFlagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:124
// index:0
// qreal fontPointSize()
func (this *QTextEdit) FontPointSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit13fontPointSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:125
// index:0
// QString fontFamily()
func (this *QTextEdit) FontFamily() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10fontFamilyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:126
// index:0
// int fontWeight()
func (this *QTextEdit) FontWeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10fontWeightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:127
// index:0
// bool fontUnderline()
func (this *QTextEdit) FontUnderline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit13fontUnderlineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:128
// index:0
// bool fontItalic()
func (this *QTextEdit) FontItalic() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10fontItalicEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:129
// index:0
// QColor textColor()
func (this *QTextEdit) TextColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit9textColorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:130
// index:0
// QColor textBackgroundColor()
func (this *QTextEdit) TextBackgroundColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit19textBackgroundColorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:131
// index:0
// QFont currentFont()
func (this *QTextEdit) CurrentFont() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit11currentFontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:132
// index:0
// Qt::Alignment alignment()
func (this *QTextEdit) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit9alignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:134
// index:0
// void mergeCurrentCharFormat(const class QTextCharFormat &)
func (this *QTextEdit) MergeCurrentCharFormat(modifier unsafe.Pointer) {
	// 0: (, modifier const QTextCharFormat &), (modifier)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit22mergeCurrentCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), modifier)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:136
// index:0
// void setCurrentCharFormat(const class QTextCharFormat &)
func (this *QTextEdit) SetCurrentCharFormat(format unsafe.Pointer) {
	// 0: (, format const QTextCharFormat &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit20setCurrentCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:137
// index:0
// QTextCharFormat currentCharFormat()
func (this *QTextEdit) CurrentCharFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit17currentCharFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:139
// index:0
// QTextEdit::AutoFormatting autoFormatting()
func (this *QTextEdit) AutoFormatting() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit14autoFormattingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:140
// index:0
// void setAutoFormatting(QTextEdit::AutoFormatting)
func (this *QTextEdit) SetAutoFormatting(features int) {
	// 0: (, QFlags<QTextEdit::AutoFormattingFlag> features), (features)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit17setAutoFormattingE6QFlagsINS_18AutoFormattingFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), features)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:142
// index:0
// bool tabChangesFocus()
func (this *QTextEdit) TabChangesFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit15tabChangesFocusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:143
// index:0
// void setTabChangesFocus(_Bool)
func (this *QTextEdit) SetTabChangesFocus(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18setTabChangesFocusEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:145
// index:0
// inline
// void setDocumentTitle(const class QString &)
func (this *QTextEdit) SetDocumentTitle(title unsafe.Pointer) {
	// 0: (, title const QString &), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16setDocumentTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:147
// index:0
// inline
// QString documentTitle()
func (this *QTextEdit) DocumentTitle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit13documentTitleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:150
// index:0
// inline
// bool isUndoRedoEnabled()
func (this *QTextEdit) IsUndoRedoEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit17isUndoRedoEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:152
// index:0
// inline
// void setUndoRedoEnabled(_Bool)
func (this *QTextEdit) SetUndoRedoEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18setUndoRedoEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:155
// index:0
// QTextEdit::LineWrapMode lineWrapMode()
func (this *QTextEdit) LineWrapMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit12lineWrapModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:156
// index:0
// void setLineWrapMode(enum QTextEdit::LineWrapMode)
func (this *QTextEdit) SetLineWrapMode(mode int) {
	// 0: (, mode QTextEdit::LineWrapMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15setLineWrapModeENS_12LineWrapModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:158
// index:0
// int lineWrapColumnOrWidth()
func (this *QTextEdit) LineWrapColumnOrWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit21lineWrapColumnOrWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:159
// index:0
// void setLineWrapColumnOrWidth(int)
func (this *QTextEdit) SetLineWrapColumnOrWidth(w int) {
	// 0: (, w int), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit24setLineWrapColumnOrWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:161
// index:0
// QTextOption::WrapMode wordWrapMode()
func (this *QTextEdit) WordWrapMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit12wordWrapModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:162
// index:0
// void setWordWrapMode(class QTextOption::WrapMode)
func (this *QTextEdit) SetWordWrapMode(policy int) {
	// 0: (, policy QTextOption::WrapMode), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15setWordWrapModeEN11QTextOption8WrapModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:164
// index:0
// bool find(const class QString &, class QTextDocument::FindFlags)
func (this *QTextEdit) Find(exp unsafe.Pointer, options int) {
	// 0: (, exp const QString &, QFlags<QTextDocument::FindFlag> options), (exp, &options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit4findERK7QString6QFlagsIN13QTextDocument8FindFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), exp, &options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:166
// index:1
// bool find(const class QRegExp &, class QTextDocument::FindFlags)
func (this *QTextEdit) Find_1(exp unsafe.Pointer, options int) {
	// 1: (, exp const QRegExp &, QFlags<QTextDocument::FindFlag> options), (exp, &options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit4findERK7QRegExp6QFlagsIN13QTextDocument8FindFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), exp, &options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:169
// index:0
// QString toPlainText()
func (this *QTextEdit) ToPlainText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit11toPlainTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:171
// index:0
// QString toHtml()
func (this *QTextEdit) ToHtml() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit6toHtmlEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:174
// index:0
// void ensureCursorVisible()
func (this *QTextEdit) EnsureCursorVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit19ensureCursorVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:176
// index:0
// virtual
// QVariant loadResource(int, const class QUrl &)
func (this *QTextEdit) LoadResource(type_ int, name unsafe.Pointer) {
	// 0: (, type int, name const QUrl &), (&type_, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit12loadResourceEiRK4QUrl", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &type_, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:178
// index:0
// QMenu * createStandardContextMenu()
func (this *QTextEdit) CreateStandardContextMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit25createStandardContextMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:179
// index:1
// QMenu * createStandardContextMenu(const class QPoint &)
func (this *QTextEdit) CreateStandardContextMenu_1(position unsafe.Pointer) {
	// 1: (, position const QPoint &), (position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit25createStandardContextMenuERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:182
// index:0
// QTextCursor cursorForPosition(const class QPoint &)
func (this *QTextEdit) CursorForPosition(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit17cursorForPositionERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:183
// index:0
// QRect cursorRect(const class QTextCursor &)
func (this *QTextEdit) CursorRect(cursor unsafe.Pointer) {
	// 0: (, cursor const QTextCursor &), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10cursorRectERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:184
// index:1
// QRect cursorRect()
func (this *QTextEdit) CursorRect_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10cursorRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:186
// index:0
// QString anchorAt(const class QPoint &)
func (this *QTextEdit) AnchorAt(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit8anchorAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:188
// index:0
// bool overwriteMode()
func (this *QTextEdit) OverwriteMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit13overwriteModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:189
// index:0
// void setOverwriteMode(_Bool)
func (this *QTextEdit) SetOverwriteMode(overwrite bool) {
	// 0: (, overwrite bool), (&overwrite)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16setOverwriteModeEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &overwrite)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:192
// index:0
// int tabStopWidth()
func (this *QTextEdit) TabStopWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit12tabStopWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:193
// index:0
// void setTabStopWidth(int)
func (this *QTextEdit) SetTabStopWidth(width int) {
	// 0: (, width int), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15setTabStopWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:196
// index:0
// qreal tabStopDistance()
func (this *QTextEdit) TabStopDistance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit15tabStopDistanceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:197
// index:0
// void setTabStopDistance(qreal)
func (this *QTextEdit) SetTabStopDistance(distance float64) {
	// 0: (, distance qreal), (&distance)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18setTabStopDistanceEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &distance)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:199
// index:0
// int cursorWidth()
func (this *QTextEdit) CursorWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit11cursorWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:200
// index:0
// void setCursorWidth(int)
func (this *QTextEdit) SetCursorWidth(width int) {
	// 0: (, width int), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14setCursorWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:202
// index:0
// bool acceptRichText()
func (this *QTextEdit) AcceptRichText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit14acceptRichTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:203
// index:0
// void setAcceptRichText(_Bool)
func (this *QTextEdit) SetAcceptRichText(accept bool) {
	// 0: (, accept bool), (&accept)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit17setAcceptRichTextEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &accept)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:211
// index:0
// QList<QTextEdit::ExtraSelection> extraSelections()
func (this *QTextEdit) ExtraSelections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit15extraSelectionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:213
// index:0
// void moveCursor(class QTextCursor::MoveOperation, class QTextCursor::MoveMode)
func (this *QTextEdit) MoveCursor(operation int, mode int) {
	// 0: (, operation QTextCursor::MoveOperation, mode QTextCursor::MoveMode), (&operation, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit10moveCursorEN11QTextCursor13MoveOperationENS0_8MoveModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &operation, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:215
// index:0
// bool canPaste()
func (this *QTextEdit) CanPaste() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit8canPasteEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:217
// index:0
// void print(class QPagedPaintDevice *)
func (this *QTextEdit) Print(printer unsafe.Pointer) {
	// 0: (, printer QPagedPaintDevice *), (printer)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit5printEP17QPagedPaintDevice", ffiqt.FFI_TYPE_VOID, this.GetCthis(), printer)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:219
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QTextEdit) InputMethodQuery(property int) {
	// 0: (, property Qt::InputMethodQuery), (&property)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &property)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:220
// index:1
// QVariant inputMethodQuery(Qt::InputMethodQuery, class QVariant)
func (this *QTextEdit) InputMethodQuery_1(query int, argument unsafe.Pointer) {
	// 1: (, query Qt::InputMethodQuery, argument QVariant), (&query, argument)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &query, argument)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:223
// index:0
// void setFontPointSize(qreal)
func (this *QTextEdit) SetFontPointSize(s float64) {
	// 0: (, s qreal), (&s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16setFontPointSizeEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:224
// index:0
// void setFontFamily(const class QString &)
func (this *QTextEdit) SetFontFamily(fontFamily unsafe.Pointer) {
	// 0: (, fontFamily const QString &), (fontFamily)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13setFontFamilyERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fontFamily)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:225
// index:0
// void setFontWeight(int)
func (this *QTextEdit) SetFontWeight(w int) {
	// 0: (, w int), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13setFontWeightEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:226
// index:0
// void setFontUnderline(_Bool)
func (this *QTextEdit) SetFontUnderline(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16setFontUnderlineEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:227
// index:0
// void setFontItalic(_Bool)
func (this *QTextEdit) SetFontItalic(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13setFontItalicEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:228
// index:0
// void setTextColor(const class QColor &)
func (this *QTextEdit) SetTextColor(c unsafe.Pointer) {
	// 0: (, c const QColor &), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit12setTextColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:229
// index:0
// void setTextBackgroundColor(const class QColor &)
func (this *QTextEdit) SetTextBackgroundColor(c unsafe.Pointer) {
	// 0: (, c const QColor &), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit22setTextBackgroundColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:230
// index:0
// void setCurrentFont(const class QFont &)
func (this *QTextEdit) SetCurrentFont(f unsafe.Pointer) {
	// 0: (, f const QFont &), (f)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14setCurrentFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:231
// index:0
// void setAlignment(Qt::Alignment)
func (this *QTextEdit) SetAlignment(a int) {
	// 0: (, QFlags<Qt::AlignmentFlag> a), (&a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:233
// index:0
// void setPlainText(const class QString &)
func (this *QTextEdit) SetPlainText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit12setPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:235
// index:0
// void setHtml(const class QString &)
func (this *QTextEdit) SetHtml(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit7setHtmlERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:237
// index:0
// void setText(const class QString &)
func (this *QTextEdit) SetText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:240
// index:0
// void cut()
func (this *QTextEdit) Cut() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit3cutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:241
// index:0
// void copy()
func (this *QTextEdit) Copy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit4copyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:242
// index:0
// void paste()
func (this *QTextEdit) Paste() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit5pasteEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:245
// index:0
// void undo()
func (this *QTextEdit) Undo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit4undoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:246
// index:0
// void redo()
func (this *QTextEdit) Redo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit4redoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:248
// index:0
// void clear()
func (this *QTextEdit) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:249
// index:0
// void selectAll()
func (this *QTextEdit) SelectAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit9selectAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:251
// index:0
// void insertPlainText(const class QString &)
func (this *QTextEdit) InsertPlainText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15insertPlainTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:253
// index:0
// void insertHtml(const class QString &)
func (this *QTextEdit) InsertHtml(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit10insertHtmlERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:258
// index:0
// void scrollToAnchor(const class QString &)
func (this *QTextEdit) ScrollToAnchor(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14scrollToAnchorERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:260
// index:0
// void zoomIn(int)
func (this *QTextEdit) ZoomIn(range_ int) {
	// 0: (, range int), (&range_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit6zoomInEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &range_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:261
// index:0
// void zoomOut(int)
func (this *QTextEdit) ZoomOut(range_ int) {
	// 0: (, range int), (&range_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit7zoomOutEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &range_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:264
// index:0
// void textChanged()
func (this *QTextEdit) TextChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit11textChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:265
// index:0
// void undoAvailable(_Bool)
func (this *QTextEdit) UndoAvailable(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13undoAvailableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:266
// index:0
// void redoAvailable(_Bool)
func (this *QTextEdit) RedoAvailable(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13redoAvailableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:267
// index:0
// void currentCharFormatChanged(const class QTextCharFormat &)
func (this *QTextEdit) CurrentCharFormatChanged(format unsafe.Pointer) {
	// 0: (, format const QTextCharFormat &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit24currentCharFormatChangedERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:268
// index:0
// void copyAvailable(_Bool)
func (this *QTextEdit) CopyAvailable(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13copyAvailableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:269
// index:0
// void selectionChanged()
func (this *QTextEdit) SelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16selectionChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:270
// index:0
// void cursorPositionChanged()
func (this *QTextEdit) CursorPositionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit21cursorPositionChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:273
// index:0
// virtual
// bool event(class QEvent *)
func (this *QTextEdit) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:274
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QTextEdit) TimerEvent(e unsafe.Pointer) {
	// 0: (, e QTimerEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:275
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QTextEdit) KeyPressEvent(e unsafe.Pointer) {
	// 0: (, e QKeyEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:276
// index:0
// virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QTextEdit) KeyReleaseEvent(e unsafe.Pointer) {
	// 0: (, e QKeyEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:277
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QTextEdit) ResizeEvent(e unsafe.Pointer) {
	// 0: (, e QResizeEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:278
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QTextEdit) PaintEvent(e unsafe.Pointer) {
	// 0: (, e QPaintEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:279
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QTextEdit) MousePressEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:280
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QTextEdit) MouseMoveEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:281
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QTextEdit) MouseReleaseEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:282
// index:0
// virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QTextEdit) MouseDoubleClickEvent(e unsafe.Pointer) {
	// 0: (, e QMouseEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:283
// index:0
// virtual
// bool focusNextPrevChild(_Bool)
func (this *QTextEdit) FocusNextPrevChild(next bool) {
	// 0: (, next bool), (&next)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18focusNextPrevChildEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:285
// index:0
// virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QTextEdit) ContextMenuEvent(e unsafe.Pointer) {
	// 0: (, e QContextMenuEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:288
// index:0
// virtual
// void dragEnterEvent(class QDragEnterEvent *)
func (this *QTextEdit) DragEnterEvent(e unsafe.Pointer) {
	// 0: (, e QDragEnterEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:289
// index:0
// virtual
// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QTextEdit) DragLeaveEvent(e unsafe.Pointer) {
	// 0: (, e QDragLeaveEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:290
// index:0
// virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QTextEdit) DragMoveEvent(e unsafe.Pointer) {
	// 0: (, e QDragMoveEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:291
// index:0
// virtual
// void dropEvent(class QDropEvent *)
func (this *QTextEdit) DropEvent(e unsafe.Pointer) {
	// 0: (, e QDropEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:293
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QTextEdit) FocusInEvent(e unsafe.Pointer) {
	// 0: (, e QFocusEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:294
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QTextEdit) FocusOutEvent(e unsafe.Pointer) {
	// 0: (, e QFocusEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:295
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QTextEdit) ShowEvent(arg0 unsafe.Pointer) {
	// 0: (, QShowEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:296
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QTextEdit) ChangeEvent(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:298
// index:0
// virtual
// void wheelEvent(class QWheelEvent *)
func (this *QTextEdit) WheelEvent(e unsafe.Pointer) {
	// 0: (, e QWheelEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:301
// index:0
// virtual
// QMimeData * createMimeDataFromSelection()
func (this *QTextEdit) CreateMimeDataFromSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit27createMimeDataFromSelectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:302
// index:0
// virtual
// bool canInsertFromMimeData(const class QMimeData *)
func (this *QTextEdit) CanInsertFromMimeData(source unsafe.Pointer) {
	// 0: (, source const QMimeData *), (source)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit21canInsertFromMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_VOID, this.GetCthis(), source)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:303
// index:0
// virtual
// void insertFromMimeData(const class QMimeData *)
func (this *QTextEdit) InsertFromMimeData(source unsafe.Pointer) {
	// 0: (, source const QMimeData *), (source)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18insertFromMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_VOID, this.GetCthis(), source)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:305
// index:0
// virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QTextEdit) InputMethodEvent(arg0 unsafe.Pointer) {
	// 0: (, QInputMethodEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:309
// index:0
// virtual
// void scrollContentsBy(int, int)
func (this *QTextEdit) ScrollContentsBy(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16scrollContentsByEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:310
// index:0
// virtual
// void doSetTextCursor(const class QTextCursor &)
func (this *QTextEdit) DoSetTextCursor(cursor unsafe.Pointer) {
	// 0: (, cursor const QTextCursor &), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15doSetTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:312
// index:0
// void zoomInF(float)
func (this *QTextEdit) ZoomInF(range_ float32) {
	// 0: (, range float), (&range_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit7zoomInFEf", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &range_)
	gopp.ErrPrint(err, rv)
}

//  body block end
