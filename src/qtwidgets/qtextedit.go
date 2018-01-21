package qtwidgets

// /usr/include/qt/QtWidgets/qtextedit.h
// #include <qtextedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
	if this == nil {
		return nil
	} else {
		return this.QAbstractScrollArea.GetCthis()
	}
}
func NewQTextEditFromPointer(cthis unsafe.Pointer) *QTextEdit {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QTextEdit{bcthis0}
}

// /usr/include/qt/QtWidgets/qtextedit.h:63
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTextEdit) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:105
// index:0
// Public
// void QTextEdit(class QWidget *)
func NewQTextEdit(parent *QWidget /*444 QWidget **/) *QTextEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtextedit.h:106
// index:1
// Public
// void QTextEdit(const class QString &, class QWidget *)
func NewQTextEdit_1(text *qtcore.QString, parent *QWidget /*444 QWidget **/) *QTextEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEditC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtextedit.h:107
// index:0
// Public virtual
// void ~QTextEdit()
func DeleteQTextEdit(*QTextEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:109
// index:0
// Public
// void setDocument(class QTextDocument *)
func (this *QTextEdit) SetDocument(document *qtgui.QTextDocument /*444 QTextDocument **/) {
	var convArg0 = document.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit11setDocumentEP13QTextDocument", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:110
// index:0
// Public
// QTextDocument * document()
func (this *QTextEdit) Document() *qtgui.QTextDocument /*444 QTextDocument **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit8documentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:112
// index:0
// Public
// void setPlaceholderText(const class QString &)
func (this *QTextEdit) SetPlaceholderText(placeholderText *qtcore.QString) {
	var convArg0 = placeholderText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18setPlaceholderTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:113
// index:0
// Public
// QString placeholderText()
func (this *QTextEdit) PlaceholderText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit15placeholderTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:115
// index:0
// Public
// void setTextCursor(const class QTextCursor &)
func (this *QTextEdit) SetTextCursor(cursor *qtgui.QTextCursor) {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13setTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:116
// index:0
// Public
// QTextCursor textCursor()
func (this *QTextEdit) TextCursor() *qtgui.QTextCursor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10textCursorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:118
// index:0
// Public
// bool isReadOnly()
func (this *QTextEdit) IsReadOnly() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10isReadOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:119
// index:0
// Public
// void setReadOnly(_Bool)
func (this *QTextEdit) SetReadOnly(ro bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit11setReadOnlyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ro)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:122
// index:0
// Public
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QTextEdit) TextInteractionFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit20textInteractionFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:124
// index:0
// Public
// qreal fontPointSize()
func (this *QTextEdit) FontPointSize() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit13fontPointSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qtextedit.h:125
// index:0
// Public
// QString fontFamily()
func (this *QTextEdit) FontFamily() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10fontFamilyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:126
// index:0
// Public
// int fontWeight()
func (this *QTextEdit) FontWeight() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10fontWeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtextedit.h:127
// index:0
// Public
// bool fontUnderline()
func (this *QTextEdit) FontUnderline() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit13fontUnderlineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:128
// index:0
// Public
// bool fontItalic()
func (this *QTextEdit) FontItalic() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10fontItalicEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:129
// index:0
// Public
// QColor textColor()
func (this *QTextEdit) TextColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit9textColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:130
// index:0
// Public
// QColor textBackgroundColor()
func (this *QTextEdit) TextBackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit19textBackgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:131
// index:0
// Public
// QFont currentFont()
func (this *QTextEdit) CurrentFont() *qtgui.QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit11currentFontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:132
// index:0
// Public
// Qt::Alignment alignment()
func (this *QTextEdit) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:134
// index:0
// Public
// void mergeCurrentCharFormat(const class QTextCharFormat &)
func (this *QTextEdit) MergeCurrentCharFormat(modifier *qtgui.QTextCharFormat) {
	var convArg0 = modifier.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit22mergeCurrentCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:136
// index:0
// Public
// void setCurrentCharFormat(const class QTextCharFormat &)
func (this *QTextEdit) SetCurrentCharFormat(format *qtgui.QTextCharFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit20setCurrentCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:137
// index:0
// Public
// QTextCharFormat currentCharFormat()
func (this *QTextEdit) CurrentCharFormat() *qtgui.QTextCharFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit17currentCharFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:139
// index:0
// Public
// QTextEdit::AutoFormatting autoFormatting()
func (this *QTextEdit) AutoFormatting() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit14autoFormattingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:142
// index:0
// Public
// bool tabChangesFocus()
func (this *QTextEdit) TabChangesFocus() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit15tabChangesFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:143
// index:0
// Public
// void setTabChangesFocus(_Bool)
func (this *QTextEdit) SetTabChangesFocus(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18setTabChangesFocusEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:145
// index:0
// Public inline
// void setDocumentTitle(const class QString &)
func (this *QTextEdit) SetDocumentTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16setDocumentTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:147
// index:0
// Public inline
// QString documentTitle()
func (this *QTextEdit) DocumentTitle() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit13documentTitleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:150
// index:0
// Public inline
// bool isUndoRedoEnabled()
func (this *QTextEdit) IsUndoRedoEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit17isUndoRedoEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:152
// index:0
// Public inline
// void setUndoRedoEnabled(_Bool)
func (this *QTextEdit) SetUndoRedoEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18setUndoRedoEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:155
// index:0
// Public
// QTextEdit::LineWrapMode lineWrapMode()
func (this *QTextEdit) LineWrapMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit12lineWrapModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:156
// index:0
// Public
// void setLineWrapMode(enum QTextEdit::LineWrapMode)
func (this *QTextEdit) SetLineWrapMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15setLineWrapModeENS_12LineWrapModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:158
// index:0
// Public
// int lineWrapColumnOrWidth()
func (this *QTextEdit) LineWrapColumnOrWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit21lineWrapColumnOrWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtextedit.h:159
// index:0
// Public
// void setLineWrapColumnOrWidth(int)
func (this *QTextEdit) SetLineWrapColumnOrWidth(w int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit24setLineWrapColumnOrWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:161
// index:0
// Public
// QTextOption::WrapMode wordWrapMode()
func (this *QTextEdit) WordWrapMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit12wordWrapModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:162
// index:0
// Public
// void setWordWrapMode(class QTextOption::WrapMode)
func (this *QTextEdit) SetWordWrapMode(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15setWordWrapModeEN11QTextOption8WrapModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:169
// index:0
// Public
// QString toPlainText()
func (this *QTextEdit) ToPlainText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit11toPlainTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:171
// index:0
// Public
// QString toHtml()
func (this *QTextEdit) ToHtml() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit6toHtmlEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:174
// index:0
// Public
// void ensureCursorVisible()
func (this *QTextEdit) EnsureCursorVisible() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit19ensureCursorVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:176
// index:0
// Public virtual
// QVariant loadResource(int, const class QUrl &)
func (this *QTextEdit) LoadResource(type_ int, name *qtcore.QUrl) *qtcore.QVariant /*123*/ {
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit12loadResourceEiRK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:178
// index:0
// Public
// QMenu * createStandardContextMenu()
func (this *QTextEdit) CreateStandardContextMenu() *QMenu /*444 QMenu **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit25createStandardContextMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:179
// index:1
// Public
// QMenu * createStandardContextMenu(const class QPoint &)
func (this *QTextEdit) CreateStandardContextMenu_1(position *qtcore.QPoint) *QMenu /*444 QMenu **/ {
	var convArg0 = position.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit25createStandardContextMenuERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:182
// index:0
// Public
// QTextCursor cursorForPosition(const class QPoint &)
func (this *QTextEdit) CursorForPosition(pos *qtcore.QPoint) *qtgui.QTextCursor /*123*/ {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit17cursorForPositionERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:183
// index:0
// Public
// QRect cursorRect(const class QTextCursor &)
func (this *QTextEdit) CursorRect(cursor *qtgui.QTextCursor) *qtcore.QRect /*123*/ {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10cursorRectERK11QTextCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:184
// index:1
// Public
// QRect cursorRect()
func (this *QTextEdit) CursorRect_1() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit10cursorRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:186
// index:0
// Public
// QString anchorAt(const class QPoint &)
func (this *QTextEdit) AnchorAt(pos *qtcore.QPoint) *qtcore.QString /*123*/ {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit8anchorAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:188
// index:0
// Public
// bool overwriteMode()
func (this *QTextEdit) OverwriteMode() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit13overwriteModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:189
// index:0
// Public
// void setOverwriteMode(_Bool)
func (this *QTextEdit) SetOverwriteMode(overwrite bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16setOverwriteModeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &overwrite)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:192
// index:0
// Public
// int tabStopWidth()
func (this *QTextEdit) TabStopWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit12tabStopWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtextedit.h:193
// index:0
// Public
// void setTabStopWidth(int)
func (this *QTextEdit) SetTabStopWidth(width int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15setTabStopWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:196
// index:0
// Public
// qreal tabStopDistance()
func (this *QTextEdit) TabStopDistance() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit15tabStopDistanceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qtextedit.h:197
// index:0
// Public
// void setTabStopDistance(qreal)
func (this *QTextEdit) SetTabStopDistance(distance float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18setTabStopDistanceEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &distance)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:199
// index:0
// Public
// int cursorWidth()
func (this *QTextEdit) CursorWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit11cursorWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtextedit.h:200
// index:0
// Public
// void setCursorWidth(int)
func (this *QTextEdit) SetCursorWidth(width int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14setCursorWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:202
// index:0
// Public
// bool acceptRichText()
func (this *QTextEdit) AcceptRichText() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit14acceptRichTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:203
// index:0
// Public
// void setAcceptRichText(_Bool)
func (this *QTextEdit) SetAcceptRichText(accept bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit17setAcceptRichTextEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &accept)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:213
// index:0
// Public
// void moveCursor(class QTextCursor::MoveOperation, class QTextCursor::MoveMode)
func (this *QTextEdit) MoveCursor(operation int, mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit10moveCursorEN11QTextCursor13MoveOperationENS0_8MoveModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &operation, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:215
// index:0
// Public
// bool canPaste()
func (this *QTextEdit) CanPaste() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit8canPasteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:217
// index:0
// Public
// void print(class QPagedPaintDevice *)
func (this *QTextEdit) Print(printer *qtgui.QPagedPaintDevice /*444 QPagedPaintDevice **/) {
	var convArg0 = printer.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit5printEP17QPagedPaintDevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:219
// index:0
// Public virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QTextEdit) InputMethodQuery(property int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &property)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:220
// index:1
// Public
// QVariant inputMethodQuery(Qt::InputMethodQuery, class QVariant)
func (this *QTextEdit) InputMethodQuery_1(query int, argument *qtcore.QVariant /*123*/) *qtcore.QVariant /*123*/ {
	var convArg1 = argument.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &query, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:223
// index:0
// Public
// void setFontPointSize(qreal)
func (this *QTextEdit) SetFontPointSize(s float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16setFontPointSizeEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:224
// index:0
// Public
// void setFontFamily(const class QString &)
func (this *QTextEdit) SetFontFamily(fontFamily *qtcore.QString) {
	var convArg0 = fontFamily.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13setFontFamilyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:225
// index:0
// Public
// void setFontWeight(int)
func (this *QTextEdit) SetFontWeight(w int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13setFontWeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:226
// index:0
// Public
// void setFontUnderline(_Bool)
func (this *QTextEdit) SetFontUnderline(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16setFontUnderlineEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:227
// index:0
// Public
// void setFontItalic(_Bool)
func (this *QTextEdit) SetFontItalic(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13setFontItalicEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:228
// index:0
// Public
// void setTextColor(const class QColor &)
func (this *QTextEdit) SetTextColor(c *qtgui.QColor) {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit12setTextColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:229
// index:0
// Public
// void setTextBackgroundColor(const class QColor &)
func (this *QTextEdit) SetTextBackgroundColor(c *qtgui.QColor) {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit22setTextBackgroundColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:230
// index:0
// Public
// void setCurrentFont(const class QFont &)
func (this *QTextEdit) SetCurrentFont(f *qtgui.QFont) {
	var convArg0 = f.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14setCurrentFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:233
// index:0
// Public
// void setPlainText(const class QString &)
func (this *QTextEdit) SetPlainText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit12setPlainTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:235
// index:0
// Public
// void setHtml(const class QString &)
func (this *QTextEdit) SetHtml(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit7setHtmlERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:237
// index:0
// Public
// void setText(const class QString &)
func (this *QTextEdit) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:240
// index:0
// Public
// void cut()
func (this *QTextEdit) Cut() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit3cutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:241
// index:0
// Public
// void copy()
func (this *QTextEdit) Copy() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit4copyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:242
// index:0
// Public
// void paste()
func (this *QTextEdit) Paste() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit5pasteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:245
// index:0
// Public
// void undo()
func (this *QTextEdit) Undo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit4undoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:246
// index:0
// Public
// void redo()
func (this *QTextEdit) Redo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit4redoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:248
// index:0
// Public
// void clear()
func (this *QTextEdit) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:249
// index:0
// Public
// void selectAll()
func (this *QTextEdit) SelectAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit9selectAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:251
// index:0
// Public
// void insertPlainText(const class QString &)
func (this *QTextEdit) InsertPlainText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15insertPlainTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:253
// index:0
// Public
// void insertHtml(const class QString &)
func (this *QTextEdit) InsertHtml(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit10insertHtmlERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:258
// index:0
// Public
// void scrollToAnchor(const class QString &)
func (this *QTextEdit) ScrollToAnchor(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14scrollToAnchorERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:260
// index:0
// Public
// void zoomIn(int)
func (this *QTextEdit) ZoomIn(range_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit6zoomInEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &range_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:261
// index:0
// Public
// void zoomOut(int)
func (this *QTextEdit) ZoomOut(range_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit7zoomOutEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &range_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:264
// index:0
// Public
// void textChanged()
func (this *QTextEdit) TextChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit11textChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:265
// index:0
// Public
// void undoAvailable(_Bool)
func (this *QTextEdit) UndoAvailable(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13undoAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:266
// index:0
// Public
// void redoAvailable(_Bool)
func (this *QTextEdit) RedoAvailable(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13redoAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:267
// index:0
// Public
// void currentCharFormatChanged(const class QTextCharFormat &)
func (this *QTextEdit) CurrentCharFormatChanged(format *qtgui.QTextCharFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit24currentCharFormatChangedERK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:268
// index:0
// Public
// void copyAvailable(_Bool)
func (this *QTextEdit) CopyAvailable(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13copyAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:269
// index:0
// Public
// void selectionChanged()
func (this *QTextEdit) SelectionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16selectionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:270
// index:0
// Public
// void cursorPositionChanged()
func (this *QTextEdit) CursorPositionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit21cursorPositionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:273
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QTextEdit) Event(e *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:274
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QTextEdit) TimerEvent(e *qtcore.QTimerEvent /*444 QTimerEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:275
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QTextEdit) KeyPressEvent(e *qtgui.QKeyEvent /*444 QKeyEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:276
// index:0
// Protected virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QTextEdit) KeyReleaseEvent(e *qtgui.QKeyEvent /*444 QKeyEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:277
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QTextEdit) ResizeEvent(e *qtgui.QResizeEvent /*444 QResizeEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:278
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QTextEdit) PaintEvent(e *qtgui.QPaintEvent /*444 QPaintEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:279
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QTextEdit) MousePressEvent(e *qtgui.QMouseEvent /*444 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:280
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QTextEdit) MouseMoveEvent(e *qtgui.QMouseEvent /*444 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:281
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QTextEdit) MouseReleaseEvent(e *qtgui.QMouseEvent /*444 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:282
// index:0
// Protected virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QTextEdit) MouseDoubleClickEvent(e *qtgui.QMouseEvent /*444 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:283
// index:0
// Protected virtual
// bool focusNextPrevChild(_Bool)
func (this *QTextEdit) FocusNextPrevChild(next bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:285
// index:0
// Protected virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QTextEdit) ContextMenuEvent(e *qtgui.QContextMenuEvent /*444 QContextMenuEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:288
// index:0
// Protected virtual
// void dragEnterEvent(class QDragEnterEvent *)
func (this *QTextEdit) DragEnterEvent(e *qtgui.QDragEnterEvent /*444 QDragEnterEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:289
// index:0
// Protected virtual
// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QTextEdit) DragLeaveEvent(e *qtgui.QDragLeaveEvent /*444 QDragLeaveEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:290
// index:0
// Protected virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QTextEdit) DragMoveEvent(e *qtgui.QDragMoveEvent /*444 QDragMoveEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:291
// index:0
// Protected virtual
// void dropEvent(class QDropEvent *)
func (this *QTextEdit) DropEvent(e *qtgui.QDropEvent /*444 QDropEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:293
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QTextEdit) FocusInEvent(e *qtgui.QFocusEvent /*444 QFocusEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:294
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QTextEdit) FocusOutEvent(e *qtgui.QFocusEvent /*444 QFocusEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:295
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QTextEdit) ShowEvent(arg0 *qtgui.QShowEvent /*444 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:296
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QTextEdit) ChangeEvent(e *qtcore.QEvent /*444 QEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:298
// index:0
// Protected virtual
// void wheelEvent(class QWheelEvent *)
func (this *QTextEdit) WheelEvent(e *qtgui.QWheelEvent /*444 QWheelEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:301
// index:0
// Protected virtual
// QMimeData * createMimeDataFromSelection()
func (this *QTextEdit) CreateMimeDataFromSelection() *qtcore.QMimeData /*444 QMimeData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit27createMimeDataFromSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:302
// index:0
// Protected virtual
// bool canInsertFromMimeData(const class QMimeData *)
func (this *QTextEdit) CanInsertFromMimeData(source *qtcore.QMimeData /*444 const QMimeData **/) bool {
	var convArg0 = source.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTextEdit21canInsertFromMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:303
// index:0
// Protected virtual
// void insertFromMimeData(const class QMimeData *)
func (this *QTextEdit) InsertFromMimeData(source *qtcore.QMimeData /*444 const QMimeData **/) {
	var convArg0 = source.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit18insertFromMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:305
// index:0
// Protected virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QTextEdit) InputMethodEvent(arg0 *qtgui.QInputMethodEvent /*444 QInputMethodEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:309
// index:0
// Protected virtual
// void scrollContentsBy(int, int)
func (this *QTextEdit) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:310
// index:0
// Protected virtual
// void doSetTextCursor(const class QTextCursor &)
func (this *QTextEdit) DoSetTextCursor(cursor *qtgui.QTextCursor) {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit15doSetTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:312
// index:0
// Protected
// void zoomInF(float)
func (this *QTextEdit) ZoomInF(range_ float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTextEdit7zoomInFEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &range_)
	gopp.ErrPrint(err, rv)
}

//  body block end
