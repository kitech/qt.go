package qtwidgets

// /usr/include/qt/QtWidgets/qtextedit.h
// #include <qtextedit.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QTextEdit) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QTextEdit) InheritTimerEvent(f func(e *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QTextEdit) InheritKeyPressEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QTextEdit) InheritKeyReleaseEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QTextEdit) InheritResizeEvent(f func(e *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QTextEdit) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QTextEdit) InheritMousePressEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QTextEdit) InheritMouseMoveEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QTextEdit) InheritMouseReleaseEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QTextEdit) InheritMouseDoubleClickEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// bool focusNextPrevChild(_Bool)
func (this *QTextEdit) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QTextEdit) InheritContextMenuEvent(f func(e *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(class QDragEnterEvent *)
func (this *QTextEdit) InheritDragEnterEvent(f func(e *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QTextEdit) InheritDragLeaveEvent(f func(e *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dragMoveEvent(class QDragMoveEvent *)
func (this *QTextEdit) InheritDragMoveEvent(f func(e *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dropEvent(class QDropEvent *)
func (this *QTextEdit) InheritDropEvent(f func(e *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QTextEdit) InheritFocusInEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QTextEdit) InheritFocusOutEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QTextEdit) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QTextEdit) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QTextEdit) InheritWheelEvent(f func(e *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// QMimeData * createMimeDataFromSelection()
func (this *QTextEdit) InheritCreateMimeDataFromSelection(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createMimeDataFromSelection", f)
}

// bool canInsertFromMimeData(const class QMimeData *)
func (this *QTextEdit) InheritCanInsertFromMimeData(f func(source *qtcore.QMimeData /*777 const QMimeData **/) bool) {
	qtrt.SetAllInheritCallback(this, "canInsertFromMimeData", f)
}

// void insertFromMimeData(const class QMimeData *)
func (this *QTextEdit) InheritInsertFromMimeData(f func(source *qtcore.QMimeData /*777 const QMimeData **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "insertFromMimeData", f)
}

// void inputMethodEvent(class QInputMethodEvent *)
func (this *QTextEdit) InheritInputMethodEvent(f func(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void scrollContentsBy(int, int)
func (this *QTextEdit) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// void doSetTextCursor(const class QTextCursor &)
func (this *QTextEdit) InheritDoSetTextCursor(f func(cursor *qtgui.QTextCursor) /*void*/) {
	qtrt.SetAllInheritCallback(this, "doSetTextCursor", f)
}

// void zoomInF(float)
func (this *QTextEdit) InheritZoomInF(f func(range_ float32) /*void*/) {
	qtrt.SetAllInheritCallback(this, "zoomInF", f)
}

type QTextEdit struct {
	*QAbstractScrollArea
}
type QTextEdit_ITF interface {
	QAbstractScrollArea_ITF
	QTextEdit_PTR() *QTextEdit
}

func (ptr *QTextEdit) QTextEdit_PTR() *QTextEdit { return ptr }

func (this *QTextEdit) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractScrollArea.GetCthis()
	}
}
func (this *QTextEdit) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractScrollArea = NewQAbstractScrollAreaFromPointer(cthis)
}
func NewQTextEditFromPointer(cthis unsafe.Pointer) *QTextEdit {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QTextEdit{bcthis0}
}
func (*QTextEdit) NewFromPointer(cthis unsafe.Pointer) *QTextEdit {
	return NewQTextEditFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtextedit.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QTextEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtextedit.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextEdit(QWidget *)
func NewQTextEdit(parent QWidget_ITF /*777 QWidget **/) *QTextEdit {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qtextedit.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextEdit(QWidget *)
func NewQTextEdit__() *QTextEdit {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qtextedit.h:106
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextEdit(const QString &, QWidget *)
func NewQTextEdit_1(text string, parent QWidget_ITF /*777 QWidget **/) *QTextEdit {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEditC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qtextedit.h:106
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextEdit(const QString &, QWidget *)
func NewQTextEdit_1_(text string) *QTextEdit {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEditC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qtextedit.h:107
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextEdit()
func DeleteQTextEdit(this *QTextEdit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEditD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtextedit.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocument(QTextDocument *)
func (this *QTextEdit) SetDocument(document qtgui.QTextDocument_ITF /*777 QTextDocument **/) {
	var convArg0 unsafe.Pointer
	if document != nil && document.QTextDocument_PTR() != nil {
		convArg0 = document.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit11setDocumentEP13QTextDocument", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document() const
func (this *QTextEdit) Document() *qtgui.QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtextedit.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlaceholderText(const QString &)
func (this *QTextEdit) SetPlaceholderText(placeholderText string) {
	var tmpArg0 = qtcore.NewQString_5(placeholderText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit18setPlaceholderTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QString placeholderText() const
func (this *QTextEdit) PlaceholderText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit15placeholderTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtextedit.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextCursor(const QTextCursor &)
func (this *QTextEdit) SetTextCursor(cursor qtgui.QTextCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit13setTextCursorERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor textCursor() const
func (this *QTextEdit) TextCursor() *qtgui.QTextCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit10textCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly() const
func (this *QTextEdit) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(_Bool)
func (this *QTextEdit) SetReadOnly(ro bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit11setReadOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ro)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextInteractionFlags(Qt::TextInteractionFlags)
func (this *QTextEdit) SetTextInteractionFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:122
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextInteractionFlags textInteractionFlags() const
func (this *QTextEdit) TextInteractionFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit20textInteractionFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal fontPointSize() const
func (this *QTextEdit) FontPointSize() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit13fontPointSizeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qtextedit.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fontFamily() const
func (this *QTextEdit) FontFamily() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit10fontFamilyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtextedit.h:126
// index:0
// Public Visibility=Default Availability=Available
// [4] int fontWeight() const
func (this *QTextEdit) FontWeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit10fontWeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtextedit.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool fontUnderline() const
func (this *QTextEdit) FontUnderline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit13fontUnderlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:128
// index:0
// Public Visibility=Default Availability=Available
// [1] bool fontItalic() const
func (this *QTextEdit) FontItalic() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit10fontItalicEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:129
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor textColor() const
func (this *QTextEdit) TextColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit9textColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:130
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor textBackgroundColor() const
func (this *QTextEdit) TextBackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit19textBackgroundColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:131
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont currentFont() const
func (this *QTextEdit) CurrentFont() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit11currentFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const
func (this *QTextEdit) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mergeCurrentCharFormat(const QTextCharFormat &)
func (this *QTextEdit) MergeCurrentCharFormat(modifier qtgui.QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if modifier != nil && modifier.QTextCharFormat_PTR() != nil {
		convArg0 = modifier.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit22mergeCurrentCharFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentCharFormat(const QTextCharFormat &)
func (this *QTextEdit) SetCurrentCharFormat(format qtgui.QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg0 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit20setCurrentCharFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:137
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat currentCharFormat() const
func (this *QTextEdit) CurrentCharFormat() *qtgui.QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit17currentCharFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextEdit::AutoFormatting autoFormatting() const
func (this *QTextEdit) AutoFormatting() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit14autoFormattingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoFormatting(QTextEdit::AutoFormatting)
func (this *QTextEdit) SetAutoFormatting(features int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit17setAutoFormattingE6QFlagsINS_18AutoFormattingFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), features)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:142
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabChangesFocus() const
func (this *QTextEdit) TabChangesFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit15tabChangesFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabChangesFocus(_Bool)
func (this *QTextEdit) SetTabChangesFocus(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit18setTabChangesFocusEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:145
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setDocumentTitle(const QString &)
func (this *QTextEdit) SetDocumentTitle(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit16setDocumentTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:147
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString documentTitle() const
func (this *QTextEdit) DocumentTitle() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit13documentTitleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtextedit.h:150
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUndoRedoEnabled() const
func (this *QTextEdit) IsUndoRedoEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit17isUndoRedoEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setUndoRedoEnabled(_Bool)
func (this *QTextEdit) SetUndoRedoEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit18setUndoRedoEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:155
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextEdit::LineWrapMode lineWrapMode() const
func (this *QTextEdit) LineWrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit12lineWrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineWrapMode(enum QTextEdit::LineWrapMode)
func (this *QTextEdit) SetLineWrapMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit15setLineWrapModeENS_12LineWrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:158
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineWrapColumnOrWidth() const
func (this *QTextEdit) LineWrapColumnOrWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit21lineWrapColumnOrWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtextedit.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineWrapColumnOrWidth(int)
func (this *QTextEdit) SetLineWrapColumnOrWidth(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit24setLineWrapColumnOrWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:161
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextOption::WrapMode wordWrapMode() const
func (this *QTextEdit) WordWrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit12wordWrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordWrapMode(QTextOption::WrapMode)
func (this *QTextEdit) SetWordWrapMode(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit15setWordWrapModeEN11QTextOption8WrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool find(const QString &, QTextDocument::FindFlags)
func (this *QTextEdit) Find(exp string, options int) bool {
	var tmpArg0 = qtcore.NewQString_5(exp)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit4findERK7QString6QFlagsIN13QTextDocument8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool find(const QString &, QTextDocument::FindFlags)
func (this *QTextEdit) Find__(exp string) bool {
	var tmpArg0 = qtcore.NewQString_5(exp)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QTextDocument::FindFlags=Elaborated, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit4findERK7QString6QFlagsIN13QTextDocument8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:166
// index:1
// Public Visibility=Default Availability=Available
// [1] bool find(const QRegExp &, QTextDocument::FindFlags)
func (this *QTextEdit) Find_1(exp qtcore.QRegExp_ITF, options int) bool {
	var convArg0 unsafe.Pointer
	if exp != nil && exp.QRegExp_PTR() != nil {
		convArg0 = exp.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit4findERK7QRegExp6QFlagsIN13QTextDocument8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:166
// index:1
// Public Visibility=Default Availability=Available
// [1] bool find(const QRegExp &, QTextDocument::FindFlags)
func (this *QTextEdit) Find_1_(exp qtcore.QRegExp_ITF) bool {
	var convArg0 unsafe.Pointer
	if exp != nil && exp.QRegExp_PTR() != nil {
		convArg0 = exp.QRegExp_PTR().GetCthis()
	}
	// arg: 1, QTextDocument::FindFlags=Elaborated, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit4findERK7QRegExp6QFlagsIN13QTextDocument8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toPlainText() const
func (this *QTextEdit) ToPlainText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit11toPlainTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtextedit.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toHtml() const
func (this *QTextEdit) ToHtml() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit6toHtmlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtextedit.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureCursorVisible()
func (this *QTextEdit) EnsureCursorVisible() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit19ensureCursorVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:176
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant loadResource(int, const QUrl &)
func (this *QTextEdit) LoadResource(type_ int, name qtcore.QUrl_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg1 = name.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit12loadResourceEiRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:178
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * createStandardContextMenu()
func (this *QTextEdit) CreateStandardContextMenu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit25createStandardContextMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtextedit.h:179
// index:1
// Public Visibility=Default Availability=Available
// [8] QMenu * createStandardContextMenu(const QPoint &)
func (this *QTextEdit) CreateStandardContextMenu_1(position qtcore.QPoint_ITF) *QMenu /*777 QMenu **/ {
	var convArg0 unsafe.Pointer
	if position != nil && position.QPoint_PTR() != nil {
		convArg0 = position.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit25createStandardContextMenuERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtextedit.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor cursorForPosition(const QPoint &) const
func (this *QTextEdit) CursorForPosition(pos qtcore.QPoint_ITF) *qtgui.QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit17cursorForPositionERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:183
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect cursorRect(const QTextCursor &) const
func (this *QTextEdit) CursorRect(cursor qtgui.QTextCursor_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit10cursorRectERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:184
// index:1
// Public Visibility=Default Availability=Available
// [16] QRect cursorRect() const
func (this *QTextEdit) CursorRect_1() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit10cursorRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:186
// index:0
// Public Visibility=Default Availability=Available
// [8] QString anchorAt(const QPoint &) const
func (this *QTextEdit) AnchorAt(pos qtcore.QPoint_ITF) string {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit8anchorAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtextedit.h:188
// index:0
// Public Visibility=Default Availability=Available
// [1] bool overwriteMode() const
func (this *QTextEdit) OverwriteMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit13overwriteModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOverwriteMode(_Bool)
func (this *QTextEdit) SetOverwriteMode(overwrite bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit16setOverwriteModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), overwrite)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:192
// index:0
// Public Visibility=Default Availability=Available
// [4] int tabStopWidth() const
func (this *QTextEdit) TabStopWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit12tabStopWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtextedit.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabStopWidth(int)
func (this *QTextEdit) SetTabStopWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit15setTabStopWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:196
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal tabStopDistance() const
func (this *QTextEdit) TabStopDistance() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit15tabStopDistanceEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qtextedit.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabStopDistance(qreal)
func (this *QTextEdit) SetTabStopDistance(distance float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit18setTabStopDistanceEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), distance)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:199
// index:0
// Public Visibility=Default Availability=Available
// [4] int cursorWidth() const
func (this *QTextEdit) CursorWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit11cursorWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtextedit.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorWidth(int)
func (this *QTextEdit) SetCursorWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit14setCursorWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:202
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptRichText() const
func (this *QTextEdit) AcceptRichText() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit14acceptRichTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:203
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptRichText(_Bool)
func (this *QTextEdit) SetAcceptRichText(accept bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit17setAcceptRichTextEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), accept)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveCursor(QTextCursor::MoveOperation, QTextCursor::MoveMode)
func (this *QTextEdit) MoveCursor(operation int, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit10moveCursorEN11QTextCursor13MoveOperationENS0_8MoveModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), operation, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveCursor(QTextCursor::MoveOperation, QTextCursor::MoveMode)
func (this *QTextEdit) MoveCursor__(operation int) {
	// arg: 1, QTextCursor::MoveMode=Elaborated, QTextCursor::MoveMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit10moveCursorEN11QTextCursor13MoveOperationENS0_8MoveModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), operation, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:215
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canPaste() const
func (this *QTextEdit) CanPaste() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit8canPasteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void print(QPagedPaintDevice *) const
func (this *QTextEdit) Print(printer qtgui.QPagedPaintDevice_ITF /*777 QPagedPaintDevice **/) {
	var convArg0 unsafe.Pointer
	if printer != nil && printer.QPagedPaintDevice_PTR() != nil {
		convArg0 = printer.QPagedPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit5printEP17QPagedPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:219
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const
func (this *QTextEdit) InputMethodQuery(property int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), property)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:220
// index:1
// Public Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery, QVariant) const
func (this *QTextEdit) InputMethodQuery_1(query int, argument qtcore.QVariant_ITF /*123*/) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if argument != nil && argument.QVariant_PTR() != nil {
		convArg1 = argument.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextedit.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFontPointSize(qreal)
func (this *QTextEdit) SetFontPointSize(s float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit16setFontPointSizeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), s)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFontFamily(const QString &)
func (this *QTextEdit) SetFontFamily(fontFamily string) {
	var tmpArg0 = qtcore.NewQString_5(fontFamily)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit13setFontFamilyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFontWeight(int)
func (this *QTextEdit) SetFontWeight(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit13setFontWeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFontUnderline(_Bool)
func (this *QTextEdit) SetFontUnderline(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit16setFontUnderlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFontItalic(_Bool)
func (this *QTextEdit) SetFontItalic(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit13setFontItalicEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextColor(const QColor &)
func (this *QTextEdit) SetTextColor(c qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if c != nil && c.QColor_PTR() != nil {
		convArg0 = c.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit12setTextColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:229
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextBackgroundColor(const QColor &)
func (this *QTextEdit) SetTextBackgroundColor(c qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if c != nil && c.QColor_PTR() != nil {
		convArg0 = c.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit22setTextBackgroundColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentFont(const QFont &)
func (this *QTextEdit) SetCurrentFont(f qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if f != nil && f.QFont_PTR() != nil {
		convArg0 = f.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit14setCurrentFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:231
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)
func (this *QTextEdit) SetAlignment(a int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlainText(const QString &)
func (this *QTextEdit) SetPlainText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit12setPlainTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:235
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &)
func (this *QTextEdit) SetHtml(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit7setHtmlERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QTextEdit) SetText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cut()
func (this *QTextEdit) Cut() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit3cutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void copy()
func (this *QTextEdit) Copy() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit4copyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paste()
func (this *QTextEdit) Paste() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit5pasteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undo()
func (this *QTextEdit) Undo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit4undoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:246
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redo()
func (this *QTextEdit) Redo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit4redoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:248
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QTextEdit) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:249
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectAll()
func (this *QTextEdit) SelectAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit9selectAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:251
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertPlainText(const QString &)
func (this *QTextEdit) InsertPlainText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit15insertPlainTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:253
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertHtml(const QString &)
func (this *QTextEdit) InsertHtml(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit10insertHtmlERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToAnchor(const QString &)
func (this *QTextEdit) ScrollToAnchor(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit14scrollToAnchorERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zoomIn(int)
func (this *QTextEdit) ZoomIn(range_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit6zoomInEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:260
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zoomIn(int)
func (this *QTextEdit) ZoomIn__() {
	// arg: 0, int=Int, =Invalid,
	range_ := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit6zoomInEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zoomOut(int)
func (this *QTextEdit) ZoomOut(range_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit7zoomOutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zoomOut(int)
func (this *QTextEdit) ZoomOut__() {
	// arg: 0, int=Int, =Invalid,
	range_ := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit7zoomOutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textChanged()
func (this *QTextEdit) TextChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit11textChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoAvailable(_Bool)
func (this *QTextEdit) UndoAvailable(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit13undoAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:266
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redoAvailable(_Bool)
func (this *QTextEdit) RedoAvailable(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit13redoAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentCharFormatChanged(const QTextCharFormat &)
func (this *QTextEdit) CurrentCharFormatChanged(format qtgui.QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg0 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit24currentCharFormatChangedERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void copyAvailable(_Bool)
func (this *QTextEdit) CopyAvailable(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit13copyAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:269
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()
func (this *QTextEdit) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:270
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorPositionChanged()
func (this *QTextEdit) CursorPositionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit21cursorPositionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:273
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QTextEdit) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:274
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QTextEdit) TimerEvent(e qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QTimerEvent_PTR() != nil {
		convArg0 = e.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:275
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QTextEdit) KeyPressEvent(e qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QKeyEvent_PTR() != nil {
		convArg0 = e.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:276
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QTextEdit) KeyReleaseEvent(e qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QKeyEvent_PTR() != nil {
		convArg0 = e.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:277
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QTextEdit) ResizeEvent(e qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QResizeEvent_PTR() != nil {
		convArg0 = e.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:278
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QTextEdit) PaintEvent(e qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QPaintEvent_PTR() != nil {
		convArg0 = e.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:279
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QTextEdit) MousePressEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:280
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QTextEdit) MouseMoveEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:281
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QTextEdit) MouseReleaseEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:282
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QTextEdit) MouseDoubleClickEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:283
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(_Bool)
func (this *QTextEdit) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:285
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QTextEdit) ContextMenuEvent(e qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QContextMenuEvent_PTR() != nil {
		convArg0 = e.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:288
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)
func (this *QTextEdit) DragEnterEvent(e qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragEnterEvent_PTR() != nil {
		convArg0 = e.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:289
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)
func (this *QTextEdit) DragLeaveEvent(e qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragLeaveEvent_PTR() != nil {
		convArg0 = e.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:290
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)
func (this *QTextEdit) DragMoveEvent(e qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragMoveEvent_PTR() != nil {
		convArg0 = e.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:291
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)
func (this *QTextEdit) DropEvent(e qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDropEvent_PTR() != nil {
		convArg0 = e.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:293
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QTextEdit) FocusInEvent(e qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QFocusEvent_PTR() != nil {
		convArg0 = e.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:294
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QTextEdit) FocusOutEvent(e qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QFocusEvent_PTR() != nil {
		convArg0 = e.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:295
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QTextEdit) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:296
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QTextEdit) ChangeEvent(e qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:298
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QTextEdit) WheelEvent(e qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QWheelEvent_PTR() != nil {
		convArg0 = e.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:301
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QMimeData * createMimeDataFromSelection() const
func (this *QTextEdit) CreateMimeDataFromSelection() *qtcore.QMimeData /*777 QMimeData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit27createMimeDataFromSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtextedit.h:302
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool canInsertFromMimeData(const QMimeData *) const
func (this *QTextEdit) CanInsertFromMimeData(source qtcore.QMimeData_ITF /*777 const QMimeData **/) bool {
	var convArg0 unsafe.Pointer
	if source != nil && source.QMimeData_PTR() != nil {
		convArg0 = source.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextEdit21canInsertFromMimeDataEPK9QMimeData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextedit.h:303
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void insertFromMimeData(const QMimeData *)
func (this *QTextEdit) InsertFromMimeData(source qtcore.QMimeData_ITF /*777 const QMimeData **/) {
	var convArg0 unsafe.Pointer
	if source != nil && source.QMimeData_PTR() != nil {
		convArg0 = source.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit18insertFromMimeDataEPK9QMimeData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:305
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)
func (this *QTextEdit) InputMethodEvent(arg0 qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QInputMethodEvent_PTR() != nil {
		convArg0 = arg0.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:309
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)
func (this *QTextEdit) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:310
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void doSetTextCursor(const QTextCursor &)
func (this *QTextEdit) DoSetTextCursor(cursor qtgui.QTextCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit15doSetTextCursorERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextedit.h:312
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void zoomInF(float)
func (this *QTextEdit) ZoomInF(range_ float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextEdit7zoomInFEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	qtrt.ErrPrint(err, rv)
}

type QTextEdit__LineWrapMode = int

const QTextEdit__NoWrap QTextEdit__LineWrapMode = 0
const QTextEdit__WidgetWidth QTextEdit__LineWrapMode = 1
const QTextEdit__FixedPixelWidth QTextEdit__LineWrapMode = 2
const QTextEdit__FixedColumnWidth QTextEdit__LineWrapMode = 3

type QTextEdit__AutoFormattingFlag = int

const QTextEdit__AutoNone QTextEdit__AutoFormattingFlag = 0
const QTextEdit__AutoBulletList QTextEdit__AutoFormattingFlag = 1
const QTextEdit__AutoAll QTextEdit__AutoFormattingFlag = -1

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
