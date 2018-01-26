package qtwidgets

// /usr/include/qt/QtWidgets/qplaintextedit.h
// #include <qplaintextedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 124
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
	if this == nil {
		return nil
	} else {
		return this.QAbstractScrollArea.GetCthis()
	}
}
func (this *QPlainTextEdit) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractScrollArea = NewQAbstractScrollAreaFromPointer(cthis)
}
func NewQPlainTextEditFromPointer(cthis unsafe.Pointer) *QPlainTextEdit {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QPlainTextEdit{bcthis0}
}
func (*QPlainTextEdit) NewFromPointer(cthis unsafe.Pointer) *QPlainTextEdit {
	return NewQPlainTextEditFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:66
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPlainTextEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:94
// index:0
// Public
// void QPlainTextEdit(class QWidget *)
func NewQPlainTextEdit(parent *QWidget /*777 QWidget **/) *QPlainTextEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPlainTextEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:95
// index:1
// Public
// void QPlainTextEdit(const class QString &, class QWidget *)
func NewQPlainTextEdit_1(text *qtcore.QString, parent *QWidget /*777 QWidget **/) *QPlainTextEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEditC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQPlainTextEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:96
// index:0
// Public virtual
// void ~QPlainTextEdit()
func DeleteQPlainTextEdit(*QPlainTextEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:98
// index:0
// Public
// void setDocument(class QTextDocument *)
func (this *QPlainTextEdit) SetDocument(document *qtgui.QTextDocument /*777 QTextDocument **/) {
	var convArg0 = document.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11setDocumentEP13QTextDocument", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:99
// index:0
// Public
// QTextDocument * document()
func (this *QPlainTextEdit) Document() *qtgui.QTextDocument /*777 QTextDocument **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit8documentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:101
// index:0
// Public
// void setPlaceholderText(const class QString &)
func (this *QPlainTextEdit) SetPlaceholderText(placeholderText *qtcore.QString) {
	var convArg0 = placeholderText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setPlaceholderTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:102
// index:0
// Public
// QString placeholderText()
func (this *QPlainTextEdit) PlaceholderText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15placeholderTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:104
// index:0
// Public
// void setTextCursor(const class QTextCursor &)
func (this *QPlainTextEdit) SetTextCursor(cursor *qtgui.QTextCursor) {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13setTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:105
// index:0
// Public
// QTextCursor textCursor()
func (this *QPlainTextEdit) TextCursor() *qtgui.QTextCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10textCursorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:107
// index:0
// Public
// bool isReadOnly()
func (this *QPlainTextEdit) IsReadOnly() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10isReadOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:108
// index:0
// Public
// void setReadOnly(_Bool)
func (this *QPlainTextEdit) SetReadOnly(ro bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11setReadOnlyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ro)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:110
// index:0
// Public
// void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QPlainTextEdit) SetTextInteractionFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:111
// index:0
// Public
// Qt::TextInteractionFlags textInteractionFlags()
func (this *QPlainTextEdit) TextInteractionFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit20textInteractionFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:113
// index:0
// Public
// void mergeCurrentCharFormat(const class QTextCharFormat &)
func (this *QPlainTextEdit) MergeCurrentCharFormat(modifier *qtgui.QTextCharFormat) {
	var convArg0 = modifier.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:114
// index:0
// Public
// void setCurrentCharFormat(const class QTextCharFormat &)
func (this *QPlainTextEdit) SetCurrentCharFormat(format *qtgui.QTextCharFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:115
// index:0
// Public
// QTextCharFormat currentCharFormat()
func (this *QPlainTextEdit) CurrentCharFormat() *qtgui.QTextCharFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17currentCharFormatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:117
// index:0
// Public
// bool tabChangesFocus()
func (this *QPlainTextEdit) TabChangesFocus() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15tabChangesFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:118
// index:0
// Public
// void setTabChangesFocus(_Bool)
func (this *QPlainTextEdit) SetTabChangesFocus(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setTabChangesFocusEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:120
// index:0
// Public inline
// void setDocumentTitle(const class QString &)
func (this *QPlainTextEdit) SetDocumentTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16setDocumentTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:122
// index:0
// Public inline
// QString documentTitle()
func (this *QPlainTextEdit) DocumentTitle() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit13documentTitleEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:125
// index:0
// Public inline
// bool isUndoRedoEnabled()
func (this *QPlainTextEdit) IsUndoRedoEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17isUndoRedoEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:127
// index:0
// Public inline
// void setUndoRedoEnabled(_Bool)
func (this *QPlainTextEdit) SetUndoRedoEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setUndoRedoEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:130
// index:0
// Public inline
// void setMaximumBlockCount(int)
func (this *QPlainTextEdit) SetMaximumBlockCount(maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit20setMaximumBlockCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:132
// index:0
// Public inline
// int maximumBlockCount()
func (this *QPlainTextEdit) MaximumBlockCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17maximumBlockCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:136
// index:0
// Public
// QPlainTextEdit::LineWrapMode lineWrapMode()
func (this *QPlainTextEdit) LineWrapMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit12lineWrapModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:137
// index:0
// Public
// void setLineWrapMode(enum QPlainTextEdit::LineWrapMode)
func (this *QPlainTextEdit) SetLineWrapMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15setLineWrapModeENS_12LineWrapModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:139
// index:0
// Public
// QTextOption::WrapMode wordWrapMode()
func (this *QPlainTextEdit) WordWrapMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit12wordWrapModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:140
// index:0
// Public
// void setWordWrapMode(class QTextOption::WrapMode)
func (this *QPlainTextEdit) SetWordWrapMode(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15setWordWrapModeEN11QTextOption8WrapModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:142
// index:0
// Public
// void setBackgroundVisible(_Bool)
func (this *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit20setBackgroundVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:143
// index:0
// Public
// bool backgroundVisible()
func (this *QPlainTextEdit) BackgroundVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17backgroundVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:145
// index:0
// Public
// void setCenterOnScroll(_Bool)
func (this *QPlainTextEdit) SetCenterOnScroll(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit17setCenterOnScrollEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:146
// index:0
// Public
// bool centerOnScroll()
func (this *QPlainTextEdit) CenterOnScroll() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit14centerOnScrollEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:148
// index:0
// Public
// bool find(const class QString &, class QTextDocument::FindFlags)
func (this *QPlainTextEdit) Find(exp *qtcore.QString, options int) bool {
	var convArg0 = exp.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4findERK7QString6QFlagsIN13QTextDocument8FindFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:150
// index:1
// Public
// bool find(const class QRegExp &, class QTextDocument::FindFlags)
func (this *QPlainTextEdit) Find_1(exp *qtcore.QRegExp, options int) bool {
	var convArg0 = exp.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4findERK7QRegExp6QFlagsIN13QTextDocument8FindFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:153
// index:0
// Public inline
// QString toPlainText()
func (this *QPlainTextEdit) ToPlainText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit11toPlainTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:156
// index:0
// Public
// void ensureCursorVisible()
func (this *QPlainTextEdit) EnsureCursorVisible() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit19ensureCursorVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:158
// index:0
// Public virtual
// QVariant loadResource(int, const class QUrl &)
func (this *QPlainTextEdit) LoadResource(type_ int, name *qtcore.QUrl) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12loadResourceEiRK4QUrl", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), type_, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:160
// index:0
// Public
// QMenu * createStandardContextMenu()
func (this *QPlainTextEdit) CreateStandardContextMenu() *QMenu /*777 QMenu **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit25createStandardContextMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:161
// index:1
// Public
// QMenu * createStandardContextMenu(const class QPoint &)
func (this *QPlainTextEdit) CreateStandardContextMenu_1(position *qtcore.QPoint) *QMenu /*777 QMenu **/ {
	var convArg0 = position.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:164
// index:0
// Public
// QTextCursor cursorForPosition(const class QPoint &)
func (this *QPlainTextEdit) CursorForPosition(pos *qtcore.QPoint) *qtgui.QTextCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17cursorForPositionERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:165
// index:0
// Public
// QRect cursorRect(const class QTextCursor &)
func (this *QPlainTextEdit) CursorRect(cursor *qtgui.QTextCursor) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10cursorRectERK11QTextCursor", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:166
// index:1
// Public
// QRect cursorRect()
func (this *QPlainTextEdit) CursorRect_1() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10cursorRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:168
// index:0
// Public
// QString anchorAt(const class QPoint &)
func (this *QPlainTextEdit) AnchorAt(pos *qtcore.QPoint) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit8anchorAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:170
// index:0
// Public
// bool overwriteMode()
func (this *QPlainTextEdit) OverwriteMode() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit13overwriteModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:171
// index:0
// Public
// void setOverwriteMode(_Bool)
func (this *QPlainTextEdit) SetOverwriteMode(overwrite bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16setOverwriteModeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), overwrite)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:174
// index:0
// Public
// int tabStopWidth()
func (this *QPlainTextEdit) TabStopWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit12tabStopWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:175
// index:0
// Public
// void setTabStopWidth(int)
func (this *QPlainTextEdit) SetTabStopWidth(width int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15setTabStopWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:178
// index:0
// Public
// qreal tabStopDistance()
func (this *QPlainTextEdit) TabStopDistance() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15tabStopDistanceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:179
// index:0
// Public
// void setTabStopDistance(qreal)
func (this *QPlainTextEdit) SetTabStopDistance(distance float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18setTabStopDistanceEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), distance)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:181
// index:0
// Public
// int cursorWidth()
func (this *QPlainTextEdit) CursorWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit11cursorWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:182
// index:0
// Public
// void setCursorWidth(int)
func (this *QPlainTextEdit) SetCursorWidth(width int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit14setCursorWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:187
// index:0
// Public
// void moveCursor(class QTextCursor::MoveOperation, class QTextCursor::MoveMode)
func (this *QPlainTextEdit) MoveCursor(operation int, mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10moveCursorEN11QTextCursor13MoveOperationENS0_8MoveModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), operation, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:189
// index:0
// Public
// bool canPaste()
func (this *QPlainTextEdit) CanPaste() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit8canPasteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:191
// index:0
// Public
// void print(class QPagedPaintDevice *)
func (this *QPlainTextEdit) Print(printer *qtgui.QPagedPaintDevice /*777 QPagedPaintDevice **/) {
	var convArg0 = printer.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit5printEP17QPagedPaintDevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:193
// index:0
// Public
// int blockCount()
func (this *QPlainTextEdit) BlockCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit10blockCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:194
// index:0
// Public virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QPlainTextEdit) InputMethodQuery(property int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), property)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:195
// index:1
// Public
// QVariant inputMethodQuery(Qt::InputMethodQuery, class QVariant)
func (this *QPlainTextEdit) InputMethodQuery_1(query int, argument *qtcore.QVariant /*123*/) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = argument.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE8QVariant", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), query, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:199
// index:0
// Public
// void setPlainText(const class QString &)
func (this *QPlainTextEdit) SetPlainText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12setPlainTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:202
// index:0
// Public
// void cut()
func (this *QPlainTextEdit) Cut() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit3cutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:203
// index:0
// Public
// void copy()
func (this *QPlainTextEdit) Copy() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4copyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:204
// index:0
// Public
// void paste()
func (this *QPlainTextEdit) Paste() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit5pasteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:207
// index:0
// Public
// void undo()
func (this *QPlainTextEdit) Undo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4undoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:208
// index:0
// Public
// void redo()
func (this *QPlainTextEdit) Redo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit4redoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:210
// index:0
// Public
// void clear()
func (this *QPlainTextEdit) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:211
// index:0
// Public
// void selectAll()
func (this *QPlainTextEdit) SelectAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit9selectAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:213
// index:0
// Public
// void insertPlainText(const class QString &)
func (this *QPlainTextEdit) InsertPlainText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15insertPlainTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:215
// index:0
// Public
// void appendPlainText(const class QString &)
func (this *QPlainTextEdit) AppendPlainText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15appendPlainTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:216
// index:0
// Public
// void appendHtml(const class QString &)
func (this *QPlainTextEdit) AppendHtml(html *qtcore.QString) {
	var convArg0 = html.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10appendHtmlERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:218
// index:0
// Public
// void centerCursor()
func (this *QPlainTextEdit) CenterCursor() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12centerCursorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:220
// index:0
// Public
// void zoomIn(int)
func (this *QPlainTextEdit) ZoomIn(range_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit6zoomInEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:221
// index:0
// Public
// void zoomOut(int)
func (this *QPlainTextEdit) ZoomOut(range_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit7zoomOutEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:224
// index:0
// Public
// void textChanged()
func (this *QPlainTextEdit) TextChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11textChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:225
// index:0
// Public
// void undoAvailable(_Bool)
func (this *QPlainTextEdit) UndoAvailable(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13undoAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:226
// index:0
// Public
// void redoAvailable(_Bool)
func (this *QPlainTextEdit) RedoAvailable(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13redoAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:227
// index:0
// Public
// void copyAvailable(_Bool)
func (this *QPlainTextEdit) CopyAvailable(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13copyAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:228
// index:0
// Public
// void selectionChanged()
func (this *QPlainTextEdit) SelectionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16selectionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:229
// index:0
// Public
// void cursorPositionChanged()
func (this *QPlainTextEdit) CursorPositionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit21cursorPositionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:231
// index:0
// Public
// void updateRequest(const class QRect &, int)
func (this *QPlainTextEdit) UpdateRequest(rect *qtcore.QRect, dy int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13updateRequestERK5QRecti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:232
// index:0
// Public
// void blockCountChanged(int)
func (this *QPlainTextEdit) BlockCountChanged(newBlockCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit17blockCountChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newBlockCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:233
// index:0
// Public
// void modificationChanged(_Bool)
func (this *QPlainTextEdit) ModificationChanged(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit19modificationChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:236
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QPlainTextEdit) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:237
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QPlainTextEdit) TimerEvent(e *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:238
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QPlainTextEdit) KeyPressEvent(e *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:239
// index:0
// Protected virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QPlainTextEdit) KeyReleaseEvent(e *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:240
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QPlainTextEdit) ResizeEvent(e *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:241
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QPlainTextEdit) PaintEvent(e *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:242
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QPlainTextEdit) MousePressEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:243
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QPlainTextEdit) MouseMoveEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:244
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QPlainTextEdit) MouseReleaseEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:245
// index:0
// Protected virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QPlainTextEdit) MouseDoubleClickEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:246
// index:0
// Protected virtual
// bool focusNextPrevChild(_Bool)
func (this *QPlainTextEdit) FocusNextPrevChild(next bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:248
// index:0
// Protected virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QPlainTextEdit) ContextMenuEvent(e *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:251
// index:0
// Protected virtual
// void dragEnterEvent(class QDragEnterEvent *)
func (this *QPlainTextEdit) DragEnterEvent(e *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:252
// index:0
// Protected virtual
// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QPlainTextEdit) DragLeaveEvent(e *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:253
// index:0
// Protected virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QPlainTextEdit) DragMoveEvent(e *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:254
// index:0
// Protected virtual
// void dropEvent(class QDropEvent *)
func (this *QPlainTextEdit) DropEvent(e *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:256
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QPlainTextEdit) FocusInEvent(e *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:257
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QPlainTextEdit) FocusOutEvent(e *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:258
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QPlainTextEdit) ShowEvent(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:259
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QPlainTextEdit) ChangeEvent(e *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:261
// index:0
// Protected virtual
// void wheelEvent(class QWheelEvent *)
func (this *QPlainTextEdit) WheelEvent(e *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:264
// index:0
// Protected virtual
// QMimeData * createMimeDataFromSelection()
func (this *QPlainTextEdit) CreateMimeDataFromSelection() *qtcore.QMimeData /*777 QMimeData **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit27createMimeDataFromSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:265
// index:0
// Protected virtual
// bool canInsertFromMimeData(const class QMimeData *)
func (this *QPlainTextEdit) CanInsertFromMimeData(source *qtcore.QMimeData /*777 const QMimeData **/) bool {
	var convArg0 = source.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit21canInsertFromMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:266
// index:0
// Protected virtual
// void insertFromMimeData(const class QMimeData *)
func (this *QPlainTextEdit) InsertFromMimeData(source *qtcore.QMimeData /*777 const QMimeData **/) {
	var convArg0 = source.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit18insertFromMimeDataEPK9QMimeData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:268
// index:0
// Protected virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QPlainTextEdit) InputMethodEvent(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:272
// index:0
// Protected virtual
// void scrollContentsBy(int, int)
func (this *QPlainTextEdit) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:273
// index:0
// Protected virtual
// void doSetTextCursor(const class QTextCursor &)
func (this *QPlainTextEdit) DoSetTextCursor(cursor *qtgui.QTextCursor) {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit15doSetTextCursorERK11QTextCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:275
// index:0
// Protected
// QTextBlock firstVisibleBlock()
func (this *QPlainTextEdit) FirstVisibleBlock() *qtgui.QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17firstVisibleBlockEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:276
// index:0
// Protected
// QPointF contentOffset()
func (this *QPlainTextEdit) ContentOffset() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit13contentOffsetEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:277
// index:0
// Protected
// QRectF blockBoundingRect(const class QTextBlock &)
func (this *QPlainTextEdit) BlockBoundingRect(block *qtgui.QTextBlock) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit17blockBoundingRectERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:278
// index:0
// Protected
// QRectF blockBoundingGeometry(const class QTextBlock &)
func (this *QPlainTextEdit) BlockBoundingGeometry(block *qtgui.QTextBlock) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit21blockBoundingGeometryERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:279
// index:0
// Protected
// QAbstractTextDocumentLayout::PaintContext getPaintContext()
func (this *QPlainTextEdit) GetPaintContext() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QPlainTextEdit15getPaintContextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:281
// index:0
// Protected
// void zoomInF(float)
func (this *QPlainTextEdit) ZoomInF(range_ float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QPlainTextEdit7zoomInFEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	gopp.ErrPrint(err, rv)
}

type QPlainTextEdit__LineWrapMode = int

const QPlainTextEdit__NoWrap QPlainTextEdit__LineWrapMode = 0
const QPlainTextEdit__WidgetWidth QPlainTextEdit__LineWrapMode = 1

//  body block end
