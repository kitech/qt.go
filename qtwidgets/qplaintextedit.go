package qtwidgets

// /usr/include/qt/QtWidgets/qplaintextedit.h
// #include <qplaintextedit.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 126
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

// bool event(QEvent *)
func (this *QPlainTextEdit) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void timerEvent(QTimerEvent *)
func (this *QPlainTextEdit) InheritTimerEvent(f func(e *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QPlainTextEdit) InheritKeyPressEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QPlainTextEdit) InheritKeyReleaseEvent(f func(e *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QPlainTextEdit) InheritResizeEvent(f func(e *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QPlainTextEdit) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QPlainTextEdit) InheritMousePressEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QPlainTextEdit) InheritMouseMoveEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QPlainTextEdit) InheritMouseReleaseEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QPlainTextEdit) InheritMouseDoubleClickEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// bool focusNextPrevChild(bool)
func (this *QPlainTextEdit) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QPlainTextEdit) InheritContextMenuEvent(f func(e *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(QDragEnterEvent *)
func (this *QPlainTextEdit) InheritDragEnterEvent(f func(e *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QPlainTextEdit) InheritDragLeaveEvent(f func(e *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dragMoveEvent(QDragMoveEvent *)
func (this *QPlainTextEdit) InheritDragMoveEvent(f func(e *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dropEvent(QDropEvent *)
func (this *QPlainTextEdit) InheritDropEvent(f func(e *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QPlainTextEdit) InheritFocusInEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QPlainTextEdit) InheritFocusOutEvent(f func(e *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QPlainTextEdit) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void changeEvent(QEvent *)
func (this *QPlainTextEdit) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QPlainTextEdit) InheritWheelEvent(f func(e *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// QMimeData * createMimeDataFromSelection()
func (this *QPlainTextEdit) InheritCreateMimeDataFromSelection(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createMimeDataFromSelection", f)
}

// bool canInsertFromMimeData(const QMimeData *)
func (this *QPlainTextEdit) InheritCanInsertFromMimeData(f func(source *qtcore.QMimeData /*777 const QMimeData **/) bool) {
	qtrt.SetAllInheritCallback(this, "canInsertFromMimeData", f)
}

// void insertFromMimeData(const QMimeData *)
func (this *QPlainTextEdit) InheritInsertFromMimeData(f func(source *qtcore.QMimeData /*777 const QMimeData **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "insertFromMimeData", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QPlainTextEdit) InheritInputMethodEvent(f func(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void scrollContentsBy(int, int)
func (this *QPlainTextEdit) InheritScrollContentsBy(f func(dx int, dy int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// void doSetTextCursor(const QTextCursor &)
func (this *QPlainTextEdit) InheritDoSetTextCursor(f func(cursor *qtgui.QTextCursor) /*void*/) {
	qtrt.SetAllInheritCallback(this, "doSetTextCursor", f)
}

// QTextBlock firstVisibleBlock()
func (this *QPlainTextEdit) InheritFirstVisibleBlock(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "firstVisibleBlock", f)
}

// QPointF contentOffset()
func (this *QPlainTextEdit) InheritContentOffset(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "contentOffset", f)
}

// QRectF blockBoundingRect(const QTextBlock &)
func (this *QPlainTextEdit) InheritBlockBoundingRect(f func(block *qtgui.QTextBlock) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "blockBoundingRect", f)
}

// QRectF blockBoundingGeometry(const QTextBlock &)
func (this *QPlainTextEdit) InheritBlockBoundingGeometry(f func(block *qtgui.QTextBlock) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "blockBoundingGeometry", f)
}

// QAbstractTextDocumentLayout::PaintContext getPaintContext()
func (this *QPlainTextEdit) InheritGetPaintContext(f func() int) {
	qtrt.SetAllInheritCallback(this, "getPaintContext", f)
}

// void zoomInF(float)
func (this *QPlainTextEdit) InheritZoomInF(f func(range_ float32) /*void*/) {
	qtrt.SetAllInheritCallback(this, "zoomInF", f)
}

/*

 */
type QPlainTextEdit struct {
	*QAbstractScrollArea
}
type QPlainTextEdit_ITF interface {
	QAbstractScrollArea_ITF
	QPlainTextEdit_PTR() *QPlainTextEdit
}

func (ptr *QPlainTextEdit) QPlainTextEdit_PTR() *QPlainTextEdit { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QPlainTextEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(QWidget *)

/*
Constructs an empty QPlainTextEdit with parent parent.
*/
func NewQPlainTextEdit(parent QWidget_ITF /*777 QWidget **/) *QPlainTextEdit {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPlainTextEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(QWidget *)

/*
Constructs an empty QPlainTextEdit with parent parent.
*/
func NewQPlainTextEdit__() *QPlainTextEdit {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPlainTextEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:95
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(const QString &, QWidget *)

/*
Constructs an empty QPlainTextEdit with parent parent.
*/
func NewQPlainTextEdit_1(text string, parent QWidget_ITF /*777 QWidget **/) *QPlainTextEdit {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEditC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPlainTextEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:95
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(const QString &, QWidget *)

/*
Constructs an empty QPlainTextEdit with parent parent.
*/
func NewQPlainTextEdit_1_(text string) *QPlainTextEdit {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEditC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPlainTextEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPlainTextEdit()

/*

 */
func DeleteQPlainTextEdit(this *QPlainTextEdit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEditD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocument(QTextDocument *)

/*
Makes document the new document of the text editor.

The parent QObject of the provided document remains the owner of the object. If the current document is a child of the text editor, then it is deleted.

The document must have a document layout that inherits QPlainTextDocumentLayout (see QTextDocument::setDocumentLayout()).

See also document().
*/
func (this *QPlainTextEdit) SetDocument(document qtgui.QTextDocument_ITF /*777 QTextDocument **/) {
	var convArg0 unsafe.Pointer
	if document != nil && document.QTextDocument_PTR() != nil {
		convArg0 = document.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit11setDocumentEP13QTextDocument", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document() const

/*
Returns a pointer to the underlying document.

See also setDocument().
*/
func (this *QPlainTextEdit) Document() *qtgui.QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlaceholderText(const QString &)

/*

 */
func (this *QPlainTextEdit) SetPlaceholderText(placeholderText string) {
	var tmpArg0 = qtcore.NewQString_5(placeholderText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit18setPlaceholderTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QString placeholderText() const

/*

 */
func (this *QPlainTextEdit) PlaceholderText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit15placeholderTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextCursor(const QTextCursor &)

/*
Sets the visible cursor.

See also textCursor().
*/
func (this *QPlainTextEdit) SetTextCursor(cursor qtgui.QTextCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit13setTextCursorERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor textCursor() const

/*
Returns a copy of the QTextCursor that represents the currently visible cursor. Note that changes on the returned cursor do not affect QPlainTextEdit's cursor; use setTextCursor() to update the visible cursor.

See also setTextCursor().
*/
func (this *QPlainTextEdit) TextCursor() *qtgui.QTextCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit10textCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly() const

/*

 */
func (this *QPlainTextEdit) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(bool)

/*

 */
func (this *QPlainTextEdit) SetReadOnly(ro bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit11setReadOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ro)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextInteractionFlags(Qt::TextInteractionFlags)

/*

 */
func (this *QPlainTextEdit) SetTextInteractionFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextInteractionFlags textInteractionFlags() const

/*

 */
func (this *QPlainTextEdit) TextInteractionFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit20textInteractionFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mergeCurrentCharFormat(const QTextCharFormat &)

/*
Merges the properties specified in modifier into the current character format by calling QTextCursor::mergeCharFormat on the editor's cursor. If the editor has a selection then the properties of modifier are directly applied to the selection.

See also QTextCursor::mergeCharFormat().
*/
func (this *QPlainTextEdit) MergeCurrentCharFormat(modifier qtgui.QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if modifier != nil && modifier.QTextCharFormat_PTR() != nil {
		convArg0 = modifier.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentCharFormat(const QTextCharFormat &)

/*
Sets the char format that is be used when inserting new text to format by calling QTextCursor::setCharFormat() on the editor's cursor. If the editor has a selection then the char format is directly applied to the selection.

See also currentCharFormat().
*/
func (this *QPlainTextEdit) SetCurrentCharFormat(format qtgui.QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg0 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:115
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat currentCharFormat() const

/*
Returns the char format that is used when inserting new text.

See also setCurrentCharFormat().
*/
func (this *QPlainTextEdit) CurrentCharFormat() *qtgui.QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit17currentCharFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabChangesFocus() const

/*

 */
func (this *QPlainTextEdit) TabChangesFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit15tabChangesFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabChangesFocus(bool)

/*

 */
func (this *QPlainTextEdit) SetTabChangesFocus(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit18setTabChangesFocusEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setDocumentTitle(const QString &)

/*

 */
func (this *QPlainTextEdit) SetDocumentTitle(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit16setDocumentTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString documentTitle() const

/*

 */
func (this *QPlainTextEdit) DocumentTitle() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit13documentTitleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUndoRedoEnabled() const

/*

 */
func (this *QPlainTextEdit) IsUndoRedoEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit17isUndoRedoEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setUndoRedoEnabled(bool)

/*

 */
func (this *QPlainTextEdit) SetUndoRedoEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit18setUndoRedoEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setMaximumBlockCount(int)

/*

 */
func (this *QPlainTextEdit) SetMaximumBlockCount(maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit20setMaximumBlockCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int maximumBlockCount() const

/*

 */
func (this *QPlainTextEdit) MaximumBlockCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit17maximumBlockCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] QPlainTextEdit::LineWrapMode lineWrapMode() const

/*

 */
func (this *QPlainTextEdit) LineWrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit12lineWrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineWrapMode(QPlainTextEdit::LineWrapMode)

/*

 */
func (this *QPlainTextEdit) SetLineWrapMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit15setLineWrapModeENS_12LineWrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] QTextOption::WrapMode wordWrapMode() const

/*

 */
func (this *QPlainTextEdit) WordWrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit12wordWrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordWrapMode(QTextOption::WrapMode)

/*

 */
func (this *QPlainTextEdit) SetWordWrapMode(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit15setWordWrapModeEN11QTextOption8WrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundVisible(bool)

/*

 */
func (this *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit20setBackgroundVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool backgroundVisible() const

/*

 */
func (this *QPlainTextEdit) BackgroundVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit17backgroundVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCenterOnScroll(bool)

/*

 */
func (this *QPlainTextEdit) SetCenterOnScroll(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit17setCenterOnScrollEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:146
// index:0
// Public Visibility=Default Availability=Available
// [1] bool centerOnScroll() const

/*

 */
func (this *QPlainTextEdit) CenterOnScroll() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit14centerOnScrollEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool find(const QString &, QTextDocument::FindFlags)

/*
Finds the next occurrence of the string, exp, using the given options. Returns true if exp was found and changes the cursor to select the match; otherwise returns false.
*/
func (this *QPlainTextEdit) Find(exp string, options int) bool {
	var tmpArg0 = qtcore.NewQString_5(exp)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit4findERK7QString6QFlagsIN13QTextDocument8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool find(const QString &, QTextDocument::FindFlags)

/*
Finds the next occurrence of the string, exp, using the given options. Returns true if exp was found and changes the cursor to select the match; otherwise returns false.
*/
func (this *QPlainTextEdit) Find__(exp string) bool {
	var tmpArg0 = qtcore.NewQString_5(exp)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QTextDocument::FindFlags=Elaborated, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit4findERK7QString6QFlagsIN13QTextDocument8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:150
// index:1
// Public Visibility=Default Availability=Available
// [1] bool find(const QRegExp &, QTextDocument::FindFlags)

/*
Finds the next occurrence of the string, exp, using the given options. Returns true if exp was found and changes the cursor to select the match; otherwise returns false.
*/
func (this *QPlainTextEdit) Find_1(exp qtcore.QRegExp_ITF, options int) bool {
	var convArg0 unsafe.Pointer
	if exp != nil && exp.QRegExp_PTR() != nil {
		convArg0 = exp.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit4findERK7QRegExp6QFlagsIN13QTextDocument8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:150
// index:1
// Public Visibility=Default Availability=Available
// [1] bool find(const QRegExp &, QTextDocument::FindFlags)

/*
Finds the next occurrence of the string, exp, using the given options. Returns true if exp was found and changes the cursor to select the match; otherwise returns false.
*/
func (this *QPlainTextEdit) Find_1_(exp qtcore.QRegExp_ITF) bool {
	var convArg0 unsafe.Pointer
	if exp != nil && exp.QRegExp_PTR() != nil {
		convArg0 = exp.QRegExp_PTR().GetCthis()
	}
	// arg: 1, QTextDocument::FindFlags=Elaborated, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit4findERK7QRegExp6QFlagsIN13QTextDocument8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:153
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toPlainText() const

/*
Returns the text of the text edit as plain text.

Note: Getter function for property plainText.

See also QPlainTextEdit::setPlainText().
*/
func (this *QPlainTextEdit) ToPlainText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit11toPlainTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureCursorVisible()

/*
Ensures that the cursor is visible by scrolling the text edit if necessary.

See also centerCursor() and centerOnScroll.
*/
func (this *QPlainTextEdit) EnsureCursorVisible() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit19ensureCursorVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:158
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant loadResource(int, const QUrl &)

/*
Loads the resource specified by the given type and name.

This function is an extension of QTextDocument::loadResource().

See also QTextDocument::loadResource().
*/
func (this *QPlainTextEdit) LoadResource(type_ int, name qtcore.QUrl_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg1 = name.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit12loadResourceEiRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * createStandardContextMenu()

/*
This function creates the standard context menu which is shown when the user clicks on the text edit with the right mouse button. It is called from the default contextMenuEvent() handler. The popup menu's ownership is transferred to the caller.

We recommend that you use the createStandardContextMenu(QPoint) version instead which will enable the actions that are sensitive to where the user clicked.
*/
func (this *QPlainTextEdit) CreateStandardContextMenu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit25createStandardContextMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:161
// index:1
// Public Visibility=Default Availability=Available
// [8] QMenu * createStandardContextMenu(const QPoint &)

/*
This function creates the standard context menu which is shown when the user clicks on the text edit with the right mouse button. It is called from the default contextMenuEvent() handler. The popup menu's ownership is transferred to the caller.

We recommend that you use the createStandardContextMenu(QPoint) version instead which will enable the actions that are sensitive to where the user clicked.
*/
func (this *QPlainTextEdit) CreateStandardContextMenu_1(position qtcore.QPoint_ITF) *QMenu /*777 QMenu **/ {
	var convArg0 unsafe.Pointer
	if position != nil && position.QPoint_PTR() != nil {
		convArg0 = position.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:164
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor cursorForPosition(const QPoint &) const

/*
returns a QTextCursor at position pos (in viewport coordinates).
*/
func (this *QPlainTextEdit) CursorForPosition(pos qtcore.QPoint_ITF) *qtgui.QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit17cursorForPositionERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:165
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect cursorRect(const QTextCursor &) const

/*
returns a rectangle (in viewport coordinates) that includes the cursor.
*/
func (this *QPlainTextEdit) CursorRect(cursor qtgui.QTextCursor_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit10cursorRectERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:166
// index:1
// Public Visibility=Default Availability=Available
// [16] QRect cursorRect() const

/*
returns a rectangle (in viewport coordinates) that includes the cursor.
*/
func (this *QPlainTextEdit) CursorRect_1() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit10cursorRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QString anchorAt(const QPoint &) const

/*
Returns the reference of the anchor at position pos, or an empty string if no anchor exists at that point.

This function was introduced in  Qt 4.7.
*/
func (this *QPlainTextEdit) AnchorAt(pos qtcore.QPoint_ITF) string {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit8anchorAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool overwriteMode() const

/*

 */
func (this *QPlainTextEdit) OverwriteMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit13overwriteModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOverwriteMode(bool)

/*

 */
func (this *QPlainTextEdit) SetOverwriteMode(overwrite bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit16setOverwriteModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), overwrite)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:174
// index:0
// Public Visibility=Default Availability=Available
// [4] int tabStopWidth() const

/*

 */
func (this *QPlainTextEdit) TabStopWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit12tabStopWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabStopWidth(int)

/*

 */
func (this *QPlainTextEdit) SetTabStopWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit15setTabStopWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:178
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal tabStopDistance() const

/*

 */
func (this *QPlainTextEdit) TabStopDistance() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit15tabStopDistanceEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabStopDistance(qreal)

/*

 */
func (this *QPlainTextEdit) SetTabStopDistance(distance float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit18setTabStopDistanceEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), distance)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:181
// index:0
// Public Visibility=Default Availability=Available
// [4] int cursorWidth() const

/*

 */
func (this *QPlainTextEdit) CursorWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit11cursorWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorWidth(int)

/*

 */
func (this *QPlainTextEdit) SetCursorWidth(width int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit14setCursorWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveCursor(QTextCursor::MoveOperation, QTextCursor::MoveMode)

/*
Moves the cursor by performing the given operation.

If mode is QTextCursor::KeepAnchor, the cursor selects the text it moves over. This is the same effect that the user achieves when they hold down the Shift key and move the cursor with the cursor keys.

See also QTextCursor::movePosition().
*/
func (this *QPlainTextEdit) MoveCursor(operation int, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit10moveCursorEN11QTextCursor13MoveOperationENS0_8MoveModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), operation, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveCursor(QTextCursor::MoveOperation, QTextCursor::MoveMode)

/*
Moves the cursor by performing the given operation.

If mode is QTextCursor::KeepAnchor, the cursor selects the text it moves over. This is the same effect that the user achieves when they hold down the Shift key and move the cursor with the cursor keys.

See also QTextCursor::movePosition().
*/
func (this *QPlainTextEdit) MoveCursor__(operation int) {
	// arg: 1, QTextCursor::MoveMode=Elaborated, QTextCursor::MoveMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit10moveCursorEN11QTextCursor13MoveOperationENS0_8MoveModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), operation, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:189
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canPaste() const

/*
Returns whether text can be pasted from the clipboard into the textedit.
*/
func (this *QPlainTextEdit) CanPaste() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit8canPasteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void print(QPagedPaintDevice *) const

/*
Convenience function to print the text edit's document to the given printer. This is equivalent to calling the print method on the document directly except that this function also supports QPrinter::Selection as print range.

See also QTextDocument::print().
*/
func (this *QPlainTextEdit) Print(printer qtgui.QPagedPaintDevice_ITF /*777 QPagedPaintDevice **/) {
	var convArg0 unsafe.Pointer
	if printer != nil && printer.QPagedPaintDevice_PTR() != nil {
		convArg0 = printer.QPagedPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit5printEP17QPagedPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:193
// index:0
// Public Visibility=Default Availability=Available
// [4] int blockCount() const

/*

 */
func (this *QPlainTextEdit) BlockCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit10blockCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:194
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
Reimplemented from QWidget::inputMethodQuery().
*/
func (this *QPlainTextEdit) InputMethodQuery(property int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), property)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:195
// index:1
// Public Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery, QVariant) const

/*
Reimplemented from QWidget::inputMethodQuery().
*/
func (this *QPlainTextEdit) InputMethodQuery_1(query int, argument qtcore.QVariant_ITF /*123*/) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if argument != nil && argument.QVariant_PTR() != nil {
		convArg1 = argument.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit16inputMethodQueryEN2Qt16InputMethodQueryE8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlainText(const QString &)

/*
Changes the text of the text edit to the string text. Any previous text is removed.

text is interpreted as plain text.

Note that the undo/redo history is cleared by this function.

Note: Setter function for property plainText.

See also toPlainText().
*/
func (this *QPlainTextEdit) SetPlainText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit12setPlainTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cut()

/*
Copies the selected text to the clipboard and deletes it from the text edit.

If there is no selected text nothing happens.

See also copy() and paste().
*/
func (this *QPlainTextEdit) Cut() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit3cutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:203
// index:0
// Public Visibility=Default Availability=Available
// [-2] void copy()

/*
Copies any selected text to the clipboard.

See also copyAvailable().
*/
func (this *QPlainTextEdit) Copy() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit4copyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paste()

/*
Pastes the text from the clipboard into the text edit at the current cursor position.

If there is no text in the clipboard nothing happens.

To change the behavior of this function, i.e. to modify what QPlainTextEdit can paste and how it is being pasted, reimplement the virtual canInsertFromMimeData() and insertFromMimeData() functions.

See also cut() and copy().
*/
func (this *QPlainTextEdit) Paste() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit5pasteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undo()

/*
Undoes the last operation.

If there is no operation to undo, i.e. there is no undo step in the undo/redo history, nothing happens.

See also redo().
*/
func (this *QPlainTextEdit) Undo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit4undoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redo()

/*
Redoes the last operation.

If there is no operation to redo, i.e. there is no redo step in the undo/redo history, nothing happens.

See also undo().
*/
func (this *QPlainTextEdit) Redo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit4redoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:210
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Deletes all the text in the text edit.

Note that the undo/redo history is cleared by this function.

See also cut() and setPlainText().
*/
func (this *QPlainTextEdit) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectAll()

/*
Selects all text.

See also copy(), cut(), and textCursor().
*/
func (this *QPlainTextEdit) SelectAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit9selectAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertPlainText(const QString &)

/*
Convenience slot that inserts text at the current cursor position.

It is equivalent to


  edit->textCursor().insertText(text);
*/
func (this *QPlainTextEdit) InsertPlainText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit15insertPlainTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void appendPlainText(const QString &)

/*
Appends a new paragraph with text to the end of the text edit.

See also appendHtml().
*/
func (this *QPlainTextEdit) AppendPlainText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit15appendPlainTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:216
// index:0
// Public Visibility=Default Availability=Available
// [-2] void appendHtml(const QString &)

/*
Appends a new paragraph with html to the end of the text edit.

appendPlainText()
*/
func (this *QPlainTextEdit) AppendHtml(html string) {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit10appendHtmlERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void centerCursor()

/*
Scrolls the document in order to center the cursor vertically.

See also ensureCursorVisible() and centerOnScroll.
*/
func (this *QPlainTextEdit) CenterCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit12centerCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zoomIn(int)

/*
Zooms in on the text by making the base font size range points larger and recalculating all font sizes to be the new size. This does not change the size of any images.

See also zoomOut().
*/
func (this *QPlainTextEdit) ZoomIn(range_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit6zoomInEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zoomIn(int)

/*
Zooms in on the text by making the base font size range points larger and recalculating all font sizes to be the new size. This does not change the size of any images.

See also zoomOut().
*/
func (this *QPlainTextEdit) ZoomIn__() {
	// arg: 0, int=Int, =Invalid,
	range_ := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit6zoomInEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zoomOut(int)

/*
This is an overloaded function.

Zooms out on the text by making the base font size range points smaller and recalculating all font sizes to be the new size. This does not change the size of any images.

See also zoomIn().
*/
func (this *QPlainTextEdit) ZoomOut(range_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit7zoomOutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zoomOut(int)

/*
This is an overloaded function.

Zooms out on the text by making the base font size range points smaller and recalculating all font sizes to be the new size. This does not change the size of any images.

See also zoomIn().
*/
func (this *QPlainTextEdit) ZoomOut__() {
	// arg: 0, int=Int, =Invalid,
	range_ := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit7zoomOutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textChanged()

/*
This signal is emitted whenever the document's content changes; for example, when text is inserted or deleted, or when formatting is applied.

Note: Notifier signal for property plainText.
*/
func (this *QPlainTextEdit) TextChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit11textChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoAvailable(bool)

/*
This signal is emitted whenever undo operations become available (available is true) or unavailable (available is false).
*/
func (this *QPlainTextEdit) UndoAvailable(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit13undoAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redoAvailable(bool)

/*
This signal is emitted whenever redo operations become available (available is true) or unavailable (available is false).
*/
func (this *QPlainTextEdit) RedoAvailable(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit13redoAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void copyAvailable(bool)

/*
This signal is emitted when text is selected or de-selected in the text edit.

When text is selected this signal will be emitted with yes set to true. If no text has been selected or if the selected text is de-selected this signal is emitted with yes set to false.

If yes is true then copy() can be used to copy the selection to the clipboard. If yes is false then copy() does nothing.

See also selectionChanged().
*/
func (this *QPlainTextEdit) CopyAvailable(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit13copyAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()

/*
This signal is emitted whenever the selection changes.

See also copyAvailable().
*/
func (this *QPlainTextEdit) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:229
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorPositionChanged()

/*
This signal is emitted whenever the position of the cursor changed.
*/
func (this *QPlainTextEdit) CursorPositionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit21cursorPositionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:231
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateRequest(const QRect &, int)

/*
This signal is emitted when the text document needs an update of the specified rect. If the text is scrolled, rect will cover the entire viewport area. If the text is scrolled vertically, dy carries the amount of pixels the viewport was scrolled.

The purpose of the signal is to support extra widgets in plain text edit subclasses that e.g. show line numbers, breakpoints, or other extra information.
*/
func (this *QPlainTextEdit) UpdateRequest(rect qtcore.QRect_ITF, dy int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit13updateRequestERK5QRecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:232
// index:0
// Public Visibility=Default Availability=Available
// [-2] void blockCountChanged(int)

/*
This signal is emitted whenever the block count changes. The new block count is passed in newBlockCount.
*/
func (this *QPlainTextEdit) BlockCountChanged(newBlockCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit17blockCountChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newBlockCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void modificationChanged(bool)

/*
This signal is emitted whenever the content of the document changes in a way that affects the modification state. If changed is true, the document has been modified; otherwise it is false.

For example, calling setModified(false) on a document and then inserting text causes the signal to get emitted. If you undo that operation, causing the document to return to its original unmodified state, the signal will get emitted again.
*/
func (this *QPlainTextEdit) ModificationChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit19modificationChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:236
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*

 */
func (this *QPlainTextEdit) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:237
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*

 */
func (this *QPlainTextEdit) TimerEvent(e qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QTimerEvent_PTR() != nil {
		convArg0 = e.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:238
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QPlainTextEdit) KeyPressEvent(e qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QKeyEvent_PTR() != nil {
		convArg0 = e.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:239
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyReleaseEvent().
*/
func (this *QPlainTextEdit) KeyReleaseEvent(e qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QKeyEvent_PTR() != nil {
		convArg0 = e.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:240
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QPlainTextEdit) ResizeEvent(e qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QResizeEvent_PTR() != nil {
		convArg0 = e.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:241
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QPlainTextEdit) PaintEvent(e qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QPaintEvent_PTR() != nil {
		convArg0 = e.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:242
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QPlainTextEdit) MousePressEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:243
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QPlainTextEdit) MouseMoveEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:244
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QPlainTextEdit) MouseReleaseEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:245
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseDoubleClickEvent().
*/
func (this *QPlainTextEdit) MouseDoubleClickEvent(e qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QMouseEvent_PTR() != nil {
		convArg0 = e.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:246
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Reimplemented from QWidget::focusNextPrevChild().
*/
func (this *QPlainTextEdit) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:248
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
Reimplemented from QWidget::contextMenuEvent().

Shows the standard context menu created with createStandardContextMenu().

If you do not want the text edit to have a context menu, you can set its contextMenuPolicy to Qt::NoContextMenu. If you want to customize the context menu, reimplement this function. If you want to extend the standard context menu, reimplement this function, call createStandardContextMenu() and extend the menu returned.

Information about the event is passed in the event object.


  void MyQPlainTextEdit::contextMenuEvent(QContextMenuEvent *event)
  {
      QMenu *menu = createStandardContextMenu();
      menu->addAction(tr("My Menu Item"));
      //...
      menu->exec(event->globalPos());
      delete menu;
  }
*/
func (this *QPlainTextEdit) ContextMenuEvent(e qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QContextMenuEvent_PTR() != nil {
		convArg0 = e.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:251
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
Reimplemented from QWidget::dragEnterEvent().
*/
func (this *QPlainTextEdit) DragEnterEvent(e qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragEnterEvent_PTR() != nil {
		convArg0 = e.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:252
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
Reimplemented from QWidget::dragLeaveEvent().
*/
func (this *QPlainTextEdit) DragLeaveEvent(e qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragLeaveEvent_PTR() != nil {
		convArg0 = e.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:253
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QWidget::dragMoveEvent().
*/
func (this *QPlainTextEdit) DragMoveEvent(e qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragMoveEvent_PTR() != nil {
		convArg0 = e.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:254
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QWidget::dropEvent().
*/
func (this *QPlainTextEdit) DropEvent(e qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDropEvent_PTR() != nil {
		convArg0 = e.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:256
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().
*/
func (this *QPlainTextEdit) FocusInEvent(e qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QFocusEvent_PTR() != nil {
		convArg0 = e.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:257
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().
*/
func (this *QPlainTextEdit) FocusOutEvent(e qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QFocusEvent_PTR() != nil {
		convArg0 = e.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:258
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QPlainTextEdit) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:259
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QPlainTextEdit) ChangeEvent(e qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:261
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWidget::wheelEvent().
*/
func (this *QPlainTextEdit) WheelEvent(e qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QWheelEvent_PTR() != nil {
		convArg0 = e.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:264
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QMimeData * createMimeDataFromSelection() const

/*
This function returns a new MIME data object to represent the contents of the text edit's current selection. It is called when the selection needs to be encapsulated into a new QMimeData object; for example, when a drag and drop operation is started, or when data is copied to the clipboard.

If you reimplement this function, note that the ownership of the returned QMimeData object is passed to the caller. The selection can be retrieved by using the textCursor() function.
*/
func (this *QPlainTextEdit) CreateMimeDataFromSelection() *qtcore.QMimeData /*777 QMimeData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit27createMimeDataFromSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMimeDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:265
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool canInsertFromMimeData(const QMimeData *) const

/*
This function returns true if the contents of the MIME data object, specified by source, can be decoded and inserted into the document. It is called for example when during a drag operation the mouse enters this widget and it is necessary to determine whether it is possible to accept the drag.
*/
func (this *QPlainTextEdit) CanInsertFromMimeData(source qtcore.QMimeData_ITF /*777 const QMimeData **/) bool {
	var convArg0 unsafe.Pointer
	if source != nil && source.QMimeData_PTR() != nil {
		convArg0 = source.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit21canInsertFromMimeDataEPK9QMimeData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:266
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void insertFromMimeData(const QMimeData *)

/*
This function inserts the contents of the MIME data object, specified by source, into the text edit at the current cursor position. It is called whenever text is inserted as the result of a clipboard paste operation, or when the text edit accepts data from a drag and drop operation.
*/
func (this *QPlainTextEdit) InsertFromMimeData(source qtcore.QMimeData_ITF /*777 const QMimeData **/) {
	var convArg0 unsafe.Pointer
	if source != nil && source.QMimeData_PTR() != nil {
		convArg0 = source.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit18insertFromMimeDataEPK9QMimeData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:268
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
Reimplemented from QWidget::inputMethodEvent().
*/
func (this *QPlainTextEdit) InputMethodEvent(arg0 qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QInputMethodEvent_PTR() != nil {
		convArg0 = arg0.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:272
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)

/*
Reimplemented from QAbstractScrollArea::scrollContentsBy().
*/
func (this *QPlainTextEdit) ScrollContentsBy(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit16scrollContentsByEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:273
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void doSetTextCursor(const QTextCursor &)

/*

 */
func (this *QPlainTextEdit) DoSetTextCursor(cursor qtgui.QTextCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit15doSetTextCursorERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:275
// index:0
// Protected Visibility=Default Availability=Available
// [16] QTextBlock firstVisibleBlock() const

/*
Returns the first visible block.

See also blockBoundingRect().
*/
func (this *QPlainTextEdit) FirstVisibleBlock() *qtgui.QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit17firstVisibleBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:276
// index:0
// Protected Visibility=Default Availability=Available
// [16] QPointF contentOffset() const

/*
Returns the content's origin in viewport coordinates.

The origin of the content of a plain text edit is always the top left corner of the first visible text block. The content offset is different from (0,0) when the text has been scrolled horizontally, or when the first visible block has been scrolled partially off the screen, i.e. the visible text does not start with the first line of the first visible block, or when the first visible block is the very first block and the editor displays a margin.

See also firstVisibleBlock(), horizontalScrollBar(), and verticalScrollBar().
*/
func (this *QPlainTextEdit) ContentOffset() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit13contentOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:277
// index:0
// Protected Visibility=Default Availability=Available
// [32] QRectF blockBoundingRect(const QTextBlock &) const

/*
Returns the bounding rectangle of the text block in the block's own coordinates.

See also blockBoundingGeometry().
*/
func (this *QPlainTextEdit) BlockBoundingRect(block qtgui.QTextBlock_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit17blockBoundingRectERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:278
// index:0
// Protected Visibility=Default Availability=Available
// [32] QRectF blockBoundingGeometry(const QTextBlock &) const

/*
Returns the bounding rectangle of the text block in content coordinates. Translate the rectangle with the contentOffset() to get visual coordinates on the viewport.

See also firstVisibleBlock() and blockBoundingRect().
*/
func (this *QPlainTextEdit) BlockBoundingGeometry(block qtgui.QTextBlock_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit21blockBoundingGeometryERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:279
// index:0
// Protected Visibility=Default Availability=Available
// [64] QAbstractTextDocumentLayout::PaintContext getPaintContext() const

/*
Returns the paint context for the viewport(), useful only when reimplementing paintEvent().
*/
func (this *QPlainTextEdit) GetPaintContext() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QPlainTextEdit15getPaintContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:281
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void zoomInF(float)

/*

 */
func (this *QPlainTextEdit) ZoomInF(range_ float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QPlainTextEdit7zoomInFEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), range_)
	qtrt.ErrPrint(err, rv)
}

/*
ConstantValue
QPlainTextEdit::NoWrap0
QPlainTextEdit::WidgetWidth1

*/
type QPlainTextEdit__LineWrapMode = int

//
const QPlainTextEdit__NoWrap QPlainTextEdit__LineWrapMode = 0

//
const QPlainTextEdit__WidgetWidth QPlainTextEdit__LineWrapMode = 1

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
