package qtgui

// /usr/include/qt/QtGui/qtextcursor.h
// #include <qtextcursor.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QTextCursor struct {
	*qtrt.CObject
}
type QTextCursor_ITF interface {
	QTextCursor_PTR() *QTextCursor
}

func (ptr *QTextCursor) QTextCursor_PTR() *QTextCursor { return ptr }

func (this *QTextCursor) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextCursor) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextCursorFromPointer(cthis unsafe.Pointer) *QTextCursor {
	return &QTextCursor{&qtrt.CObject{cthis}}
}
func (*QTextCursor) NewFromPointer(cthis unsafe.Pointer) *QTextCursor {
	return NewQTextCursorFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextcursor.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextCursor()

/*
Constructs a null cursor.
*/
func NewQTextCursor() *QTextCursor {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursorC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextCursor)
	return gothis
}

// /usr/include/qt/QtGui/qtextcursor.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextCursor(QTextDocument *)

/*
Constructs a null cursor.
*/
func NewQTextCursor_1(document QTextDocument_ITF /*777 QTextDocument **/) *QTextCursor {
	var convArg0 unsafe.Pointer
	if document != nil && document.QTextDocument_PTR() != nil {
		convArg0 = document.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursorC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextCursor)
	return gothis
}

// /usr/include/qt/QtGui/qtextcursor.h:73
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextCursor(QTextFrame *)

/*
Constructs a null cursor.
*/
func NewQTextCursor_2(frame QTextFrame_ITF /*777 QTextFrame **/) *QTextCursor {
	var convArg0 unsafe.Pointer
	if frame != nil && frame.QTextFrame_PTR() != nil {
		convArg0 = frame.QTextFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursorC2EP10QTextFrame", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextCursor)
	return gothis
}

// /usr/include/qt/QtGui/qtextcursor.h:74
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTextCursor(const QTextBlock &)

/*
Constructs a null cursor.
*/
func NewQTextCursor_3(block QTextBlock_ITF) *QTextCursor {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursorC2ERK10QTextBlock", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextCursor)
	return gothis
}

// /usr/include/qt/QtGui/qtextcursor.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QTextCursor & operator=(QTextCursor &&)

/*

 */
func (this *QTextCursor) Operator_equal(other unsafe.Pointer /*333*/) *QTextCursor {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursoraSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:79
// index:1
// Public Visibility=Default Availability=Available
// [8] QTextCursor & operator=(const QTextCursor &)

/*

 */
func (this *QTextCursor) Operator_equal_1(other QTextCursor_ITF) *QTextCursor {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextCursor_PTR() != nil {
		convArg0 = other.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursoraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextCursor()

/*

 */
func DeleteQTextCursor(this *QTextCursor) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextcursor.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QTextCursor &)

/*
Swaps this text cursor instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QTextCursor) Swap(other QTextCursor_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextCursor_PTR() != nil {
		convArg0 = other.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if the cursor is null; otherwise returns false. A null cursor is created by the default constructor.
*/
func (this *QTextCursor) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(int, enum QTextCursor::MoveMode)

/*
Moves the cursor to the absolute position in the document specified by pos using a MoveMode specified by m. The cursor is positioned between characters.

See also position(), movePosition(), and anchor().
*/
func (this *QTextCursor) SetPosition(pos int, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11setPositionEiNS_8MoveModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(int, enum QTextCursor::MoveMode)

/*
Moves the cursor to the absolute position in the document specified by pos using a MoveMode specified by m. The cursor is positioned between characters.

See also position(), movePosition(), and anchor().
*/
func (this *QTextCursor) SetPosition__(pos int) {
	// arg: 1, QTextCursor::MoveMode=Enum, QTextCursor::MoveMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11setPositionEiNS_8MoveModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int position() const

/*
Returns the absolute position of the cursor within the document. The cursor is positioned between characters.

See also setPosition(), movePosition(), anchor(), and positionInBlock().
*/
func (this *QTextCursor) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int positionInBlock() const

/*
Returns the relative position of the cursor within the block. The cursor is positioned between characters.

This is equivalent to position() - block().position().

This function was introduced in  Qt 4.7.

See also position().
*/
func (this *QTextCursor) PositionInBlock() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor15positionInBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] int anchor() const

/*
Returns the anchor position; this is the same as position() unless there is a selection in which case position() marks one end of the selection and anchor() marks the other end. Just like the cursor position, the anchor position is between characters.

See also position(), setPosition(), movePosition(), selectionStart(), and selectionEnd().
*/
func (this *QTextCursor) Anchor() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor6anchorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertText(const QString &)

/*
Inserts text at the current position, using the current character format.

If there is a selection, the selection is deleted and replaced by text, for example:


  cursor.clearSelection();
  cursor.movePosition(QTextCursor::NextWord, QTextCursor::KeepAnchor);
  cursor.insertText("Hello World");



This clears any existing selection, selects the word at the cursor (i.e. from position() forward), and replaces the selection with the phrase "Hello World".

Any ASCII linefeed characters (\n) in the inserted text are transformed into unicode block separators, corresponding to insertBlock() calls.

See also charFormat() and hasSelection().
*/
func (this *QTextCursor) InsertText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10insertTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:98
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertText(const QString &, const QTextCharFormat &)

/*
Inserts text at the current position, using the current character format.

If there is a selection, the selection is deleted and replaced by text, for example:


  cursor.clearSelection();
  cursor.movePosition(QTextCursor::NextWord, QTextCursor::KeepAnchor);
  cursor.insertText("Hello World");



This clears any existing selection, selects the word at the cursor (i.e. from position() forward), and replaces the selection with the phrase "Hello World".

Any ASCII linefeed characters (\n) in the inserted text are transformed into unicode block separators, corresponding to insertBlock() calls.

See also charFormat() and hasSelection().
*/
func (this *QTextCursor) InsertText_1(text string, format QTextCharFormat_ITF) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg1 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10insertTextERK7QStringRK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool movePosition(enum QTextCursor::MoveOperation, enum QTextCursor::MoveMode, int)

/*
Moves the cursor by performing the given operation n times, using the specified mode, and returns true if all operations were completed successfully; otherwise returns false.

For example, if this function is repeatedly used to seek to the end of the next word, it will eventually fail when the end of the document is reached.

By default, the move operation is performed once (n = 1).

If mode is KeepAnchor, the cursor selects the text it moves over. This is the same effect that the user achieves when they hold down the Shift key and move the cursor with the cursor keys.

See also setVisualNavigation().
*/
func (this *QTextCursor) MovePosition(op int, arg1 int, n int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor12movePositionENS_13MoveOperationENS_8MoveModeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), op, arg1, n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool movePosition(enum QTextCursor::MoveOperation, enum QTextCursor::MoveMode, int)

/*
Moves the cursor by performing the given operation n times, using the specified mode, and returns true if all operations were completed successfully; otherwise returns false.

For example, if this function is repeatedly used to seek to the end of the next word, it will eventually fail when the end of the document is reached.

By default, the move operation is performed once (n = 1).

If mode is KeepAnchor, the cursor selects the text it moves over. This is the same effect that the user achieves when they hold down the Shift key and move the cursor with the cursor keys.

See also setVisualNavigation().
*/
func (this *QTextCursor) MovePosition__(op int) bool {
	// arg: 1, QTextCursor::MoveMode=Enum, QTextCursor::MoveMode=Enum,
	arg1 := 0
	// arg: 2, int=Int, =Invalid,
	n := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor12movePositionENS_13MoveOperationENS_8MoveModeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), op, arg1, n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool movePosition(enum QTextCursor::MoveOperation, enum QTextCursor::MoveMode, int)

/*
Moves the cursor by performing the given operation n times, using the specified mode, and returns true if all operations were completed successfully; otherwise returns false.

For example, if this function is repeatedly used to seek to the end of the next word, it will eventually fail when the end of the document is reached.

By default, the move operation is performed once (n = 1).

If mode is KeepAnchor, the cursor selects the text it moves over. This is the same effect that the user achieves when they hold down the Shift key and move the cursor with the cursor keys.

See also setVisualNavigation().
*/
func (this *QTextCursor) MovePosition__1(op int, arg1 int) bool {
	// arg: 2, int=Int, =Invalid,
	n := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor12movePositionENS_13MoveOperationENS_8MoveModeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), op, arg1, n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool visualNavigation() const

/*
Returns true if the cursor does visual navigation; otherwise returns false.

Visual navigation means skipping over hidden text paragraphs. The default is false.

This function was introduced in  Qt 4.4.

See also setVisualNavigation() and movePosition().
*/
func (this *QTextCursor) VisualNavigation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor16visualNavigationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisualNavigation(_Bool)

/*
Sets visual navigation to b.

Visual navigation means skipping over hidden text paragraphs. The default is false.

This function was introduced in  Qt 4.4.

See also visualNavigation() and movePosition().
*/
func (this *QTextCursor) SetVisualNavigation(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor19setVisualNavigationEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalMovementX(int)

/*
Sets the visual x position for vertical cursor movements to x.

The vertical movement x position is cleared automatically when the cursor moves horizontally, and kept unchanged when the cursor moves vertically. The mechanism allows the cursor to move up and down on a visually straight line with proportional fonts, and to gently "jump" over short lines.

A value of -1 indicates no predefined x position. It will then be set automatically the next time the cursor moves up or down.

This function was introduced in  Qt 4.7.

See also verticalMovementX().
*/
func (this *QTextCursor) SetVerticalMovementX(x int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor20setVerticalMovementXEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int verticalMovementX() const

/*
Returns the visual x position for vertical cursor movements.

A value of -1 indicates no predefined x position. It will then be set automatically the next time the cursor moves up or down.

This function was introduced in  Qt 4.7.

See also setVerticalMovementX().
*/
func (this *QTextCursor) VerticalMovementX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor17verticalMovementXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeepPositionOnInsert(_Bool)

/*
Defines whether the cursor should keep its current position when text gets inserted at the current position of the cursor.

If b is true, the cursor keeps its current position when text gets inserted at the positing of the cursor. If b is false, the cursor moves along with the inserted text.

The default is false.

Note that a cursor always moves when text is inserted before the current position of the cursor, and it always keeps its position when text is inserted after the current position of the cursor.

This function was introduced in  Qt 4.7.

See also keepPositionOnInsert().
*/
func (this *QTextCursor) SetKeepPositionOnInsert(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor23setKeepPositionOnInsertEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool keepPositionOnInsert() const

/*
Returns whether the cursor should keep its current position when text gets inserted at the position of the cursor.

The default is false;

This function was introduced in  Qt 4.7.

See also setKeepPositionOnInsert().
*/
func (this *QTextCursor) KeepPositionOnInsert() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor20keepPositionOnInsertEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deleteChar()

/*
If there is no selected text, deletes the character at the current cursor position; otherwise deletes the selected text.

See also deletePreviousChar(), hasSelection(), and clearSelection().
*/
func (this *QTextCursor) DeleteChar() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10deleteCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deletePreviousChar()

/*
If there is no selected text, deletes the character before the current cursor position; otherwise deletes the selected text.

See also deleteChar(), hasSelection(), and clearSelection().
*/
func (this *QTextCursor) DeletePreviousChar() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor18deletePreviousCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void select(enum QTextCursor::SelectionType)

/*
Selects text in the document according to the given selection.
*/
func (this *QTextCursor) Select(selection int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor6selectENS_13SelectionTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), selection)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:153
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSelection() const

/*
Returns true if the cursor contains a selection; otherwise returns false.
*/
func (this *QTextCursor) HasSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12hasSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:154
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasComplexSelection() const

/*
Returns true if the cursor contains a selection that is not simply a range from selectionStart() to selectionEnd(); otherwise returns false.

Complex selections are ones that span at least two cells in a table; their extent is specified by selectedTableCells().
*/
func (this *QTextCursor) HasComplexSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor19hasComplexSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeSelectedText()

/*
If there is a selection, its content is deleted; otherwise does nothing.

See also hasSelection().
*/
func (this *QTextCursor) RemoveSelectedText() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor18removeSelectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearSelection()

/*
Clears the current selection by setting the anchor to the cursor position.

Note that it does not delete the text of the selection.

See also removeSelectedText() and hasSelection().
*/
func (this *QTextCursor) ClearSelection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor14clearSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionStart() const

/*
Returns the start of the selection or position() if the cursor doesn't have a selection.

See also selectionEnd(), position(), and anchor().
*/
func (this *QTextCursor) SelectionStart() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor14selectionStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:158
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionEnd() const

/*
Returns the end of the selection or position() if the cursor doesn't have a selection.

See also selectionStart(), position(), and anchor().
*/
func (this *QTextCursor) SelectionEnd() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12selectionEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedText() const

/*
Returns the current selection's text (which may be empty). This only returns the text, with no rich text formatting information. If you want a document fragment (i.e. formatted rich text) use selection() instead.

Note: If the selection obtained from an editor spans a line break, the text will contain a Unicode U+2029 paragraph separator character instead of a newline \n character. Use QString::replace() to replace these characters with newlines.
*/
func (this *QTextCursor) SelectedText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12selectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextcursor.h:161
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocumentFragment selection() const

/*
Returns the current selection (which may be empty) with all its formatting information. If you just want the selected text (i.e. plain text) use selectedText() instead.

Note: Unlike QTextDocumentFragment::toPlainText(), selectedText() may include special unicode characters such as QChar::ParagraphSeparator.

See also QTextDocumentFragment::toPlainText().
*/
func (this *QTextCursor) Selection() *QTextDocumentFragment /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor9selectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextDocumentFragment)
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectedTableCells(int *, int *, int *, int *) const

/*
If the selection spans over table cells, firstRow is populated with the number of the first row in the selection, firstColumn with the number of the first column in the selection, and numRows and numColumns with the number of rows and columns in the selection. If the selection does not span any table cells the results are harmless but undefined.
*/
func (this *QTextCursor) SelectedTableCells(firstRow unsafe.Pointer /*666*/, numRows unsafe.Pointer /*666*/, firstColumn unsafe.Pointer /*666*/, numColumns unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor18selectedTableCellsEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), firstRow, numRows, firstColumn, numColumns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:164
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock block() const

/*
Returns the block that contains the cursor.
*/
func (this *QTextCursor) Block() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor5blockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:166
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat charFormat() const

/*
Returns the format of the character immediately before the cursor position(). If the cursor is positioned at the beginning of a text block that is not empty then the format of the character immediately after the cursor is returned.

See also setCharFormat(), insertText(), and blockFormat().
*/
func (this *QTextCursor) CharFormat() *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor10charFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCharFormat(const QTextCharFormat &)

/*
Sets the cursor's current character format to the given format. If the cursor has a selection, the given format is applied to the current selection.

See also charFormat(), hasSelection(), and mergeCharFormat().
*/
func (this *QTextCursor) SetCharFormat(format QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg0 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor13setCharFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mergeCharFormat(const QTextCharFormat &)

/*
Merges the cursor's current character format with the properties described by format modifier. If the cursor has a selection, this function applies all the properties set in modifier to all the character formats that are part of the selection.

See also hasSelection() and setCharFormat().
*/
func (this *QTextCursor) MergeCharFormat(modifier QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if modifier != nil && modifier.QTextCharFormat_PTR() != nil {
		convArg0 = modifier.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor15mergeCharFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:170
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlockFormat blockFormat() const

/*
Returns the block format of the block the cursor is in.

See also setBlockFormat() and charFormat().
*/
func (this *QTextCursor) BlockFormat() *QTextBlockFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor11blockFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlockFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:171
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBlockFormat(const QTextBlockFormat &)

/*
Sets the block format of the current block (or all blocks that are contained in the selection) to format.

See also blockFormat() and mergeBlockFormat().
*/
func (this *QTextCursor) SetBlockFormat(format QTextBlockFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextBlockFormat_PTR() != nil {
		convArg0 = format.QTextBlockFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor14setBlockFormatERK16QTextBlockFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mergeBlockFormat(const QTextBlockFormat &)

/*
Modifies the block format of the current block (or all blocks that are contained in the selection) with the block format specified by modifier.

See also setBlockFormat() and blockFormat().
*/
func (this *QTextCursor) MergeBlockFormat(modifier QTextBlockFormat_ITF) {
	var convArg0 unsafe.Pointer
	if modifier != nil && modifier.QTextBlockFormat_PTR() != nil {
		convArg0 = modifier.QTextBlockFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor16mergeBlockFormatERK16QTextBlockFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:174
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat blockCharFormat() const

/*
Returns the block character format of the block the cursor is in.

The block char format is the format used when inserting text at the beginning of an empty block.

See also setBlockCharFormat().
*/
func (this *QTextCursor) BlockCharFormat() *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor15blockCharFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBlockCharFormat(const QTextCharFormat &)

/*
Sets the block char format of the current block (or all blocks that are contained in the selection) to format.

See also blockCharFormat().
*/
func (this *QTextCursor) SetBlockCharFormat(format QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg0 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor18setBlockCharFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mergeBlockCharFormat(const QTextCharFormat &)

/*
Modifies the block char format of the current block (or all blocks that are contained in the selection) with the block format specified by modifier.

See also setBlockCharFormat().
*/
func (this *QTextCursor) MergeBlockCharFormat(modifier QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if modifier != nil && modifier.QTextCharFormat_PTR() != nil {
		convArg0 = modifier.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor20mergeBlockCharFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:178
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atBlockStart() const

/*
Returns true if the cursor is at the start of a block; otherwise returns false.

See also atBlockEnd() and atStart().
*/
func (this *QTextCursor) AtBlockStart() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12atBlockStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:179
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atBlockEnd() const

/*
Returns true if the cursor is at the end of a block; otherwise returns false.

See also atBlockStart() and atEnd().
*/
func (this *QTextCursor) AtBlockEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor10atBlockEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:180
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atStart() const

/*
Returns true if the cursor is at the start of the document; otherwise returns false.

See also atBlockStart() and atEnd().
*/
func (this *QTextCursor) AtStart() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor7atStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atEnd() const

/*
Returns true if the cursor is at the end of the document; otherwise returns false.

This function was introduced in  Qt 4.6.

See also atStart() and atBlockEnd().
*/
func (this *QTextCursor) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertBlock()

/*
Inserts a new empty block at the cursor position() with the current blockFormat() and charFormat().

See also setBlockFormat().
*/
func (this *QTextCursor) InsertBlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:184
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertBlock(const QTextBlockFormat &)

/*
Inserts a new empty block at the cursor position() with the current blockFormat() and charFormat().

See also setBlockFormat().
*/
func (this *QTextCursor) InsertBlock_1(format QTextBlockFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextBlockFormat_PTR() != nil {
		convArg0 = format.QTextBlockFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertBlockERK16QTextBlockFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:185
// index:2
// Public Visibility=Default Availability=Available
// [-2] void insertBlock(const QTextBlockFormat &, const QTextCharFormat &)

/*
Inserts a new empty block at the cursor position() with the current blockFormat() and charFormat().

See also setBlockFormat().
*/
func (this *QTextCursor) InsertBlock_2(format QTextBlockFormat_ITF, charFormat QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextBlockFormat_PTR() != nil {
		convArg0 = format.QTextBlockFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if charFormat != nil && charFormat.QTextCharFormat_PTR() != nil {
		convArg1 = charFormat.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertBlockERK16QTextBlockFormatRK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:187
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextList * insertList(const QTextListFormat &)

/*
Inserts a new block at the current position and makes it the first list item of a newly created list with the given format. Returns the created list.

See also currentList(), createList(), and insertBlock().
*/
func (this *QTextCursor) InsertList(format QTextListFormat_ITF) *QTextList /*777 QTextList **/ {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextListFormat_PTR() != nil {
		convArg0 = format.QTextListFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10insertListERK15QTextListFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:188
// index:1
// Public Visibility=Default Availability=Available
// [8] QTextList * insertList(QTextListFormat::Style)

/*
Inserts a new block at the current position and makes it the first list item of a newly created list with the given format. Returns the created list.

See also currentList(), createList(), and insertBlock().
*/
func (this *QTextCursor) InsertList_1(style int) *QTextList /*777 QTextList **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10insertListEN15QTextListFormat5StyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:190
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextList * createList(const QTextListFormat &)

/*
Creates and returns a new list with the given format, and makes the current paragraph the cursor is in the first list item.

See also insertList() and currentList().
*/
func (this *QTextCursor) CreateList(format QTextListFormat_ITF) *QTextList /*777 QTextList **/ {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextListFormat_PTR() != nil {
		convArg0 = format.QTextListFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10createListERK15QTextListFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:191
// index:1
// Public Visibility=Default Availability=Available
// [8] QTextList * createList(QTextListFormat::Style)

/*
Creates and returns a new list with the given format, and makes the current paragraph the cursor is in the first list item.

See also insertList() and currentList().
*/
func (this *QTextCursor) CreateList_1(style int) *QTextList /*777 QTextList **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10createListEN15QTextListFormat5StyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:192
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextList * currentList() const

/*
Returns the current list if the cursor position() is inside a block that is part of a list; otherwise returns 0.

See also insertList() and createList().
*/
func (this *QTextCursor) CurrentList() *QTextList /*777 QTextList **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor11currentListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:194
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextTable * insertTable(int, int, const QTextTableFormat &)

/*
Creates a new table with the given number of rows and columns in the specified format, inserts it at the current cursor position() in the document, and returns the table object. The cursor is moved to the beginning of the first cell.

There must be at least one row and one column in the table.

See also currentTable().
*/
func (this *QTextCursor) InsertTable(rows int, cols int, format QTextTableFormat_ITF) *QTextTable /*777 QTextTable **/ {
	var convArg2 unsafe.Pointer
	if format != nil && format.QTextTableFormat_PTR() != nil {
		convArg2 = format.QTextTableFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertTableEiiRK16QTextTableFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rows, cols, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextTableFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:195
// index:1
// Public Visibility=Default Availability=Available
// [8] QTextTable * insertTable(int, int)

/*
Creates a new table with the given number of rows and columns in the specified format, inserts it at the current cursor position() in the document, and returns the table object. The cursor is moved to the beginning of the first cell.

There must be at least one row and one column in the table.

See also currentTable().
*/
func (this *QTextCursor) InsertTable_1(rows int, cols int) *QTextTable /*777 QTextTable **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertTableEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rows, cols)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextTableFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:196
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextTable * currentTable() const

/*
Returns a pointer to the current table if the cursor position() is inside a block that is part of a table; otherwise returns 0.

See also insertTable().
*/
func (this *QTextCursor) CurrentTable() *QTextTable /*777 QTextTable **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12currentTableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextTableFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:198
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextFrame * insertFrame(const QTextFrameFormat &)

/*
Inserts a frame with the given format at the current cursor position(), moves the cursor position() inside the frame, and returns the frame.

If the cursor holds a selection, the whole selection is moved inside the frame.

See also hasSelection().
*/
func (this *QTextCursor) InsertFrame(format QTextFrameFormat_ITF) *QTextFrame /*777 QTextFrame **/ {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextFrameFormat_PTR() != nil {
		convArg0 = format.QTextFrameFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertFrameERK16QTextFrameFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:199
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextFrame * currentFrame() const

/*
Returns a pointer to the current frame. Returns 0 if the cursor is invalid.

See also insertFrame().
*/
func (this *QTextCursor) CurrentFrame() *QTextFrame /*777 QTextFrame **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12currentFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertFragment(const QTextDocumentFragment &)

/*
Inserts the text fragment at the current position().
*/
func (this *QTextCursor) InsertFragment(fragment QTextDocumentFragment_ITF) {
	var convArg0 unsafe.Pointer
	if fragment != nil && fragment.QTextDocumentFragment_PTR() != nil {
		convArg0 = fragment.QTextDocumentFragment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor14insertFragmentERK21QTextDocumentFragment", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertHtml(const QString &)

/*
Inserts the text html at the current position(). The text is interpreted as HTML.

Note: When using this function with a style sheet, the style sheet will only apply to the current block in the document. In order to apply a style sheet throughout a document, use QTextDocument::setDefaultStyleSheet() instead.

This function was introduced in  Qt 4.2.
*/
func (this *QTextCursor) InsertHtml(html string) {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10insertHtmlERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertImage(const QTextImageFormat &, QTextFrameFormat::Position)

/*
Inserts the image defined by format at the current position().
*/
func (this *QTextCursor) InsertImage(format QTextImageFormat_ITF, alignment int) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextImageFormat_PTR() != nil {
		convArg0 = format.QTextImageFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK16QTextImageFormatN16QTextFrameFormat8PositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:208
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertImage(const QTextImageFormat &)

/*
Inserts the image defined by format at the current position().
*/
func (this *QTextCursor) InsertImage_1(format QTextImageFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextImageFormat_PTR() != nil {
		convArg0 = format.QTextImageFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK16QTextImageFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:209
// index:2
// Public Visibility=Default Availability=Available
// [-2] void insertImage(const QString &)

/*
Inserts the image defined by format at the current position().
*/
func (this *QTextCursor) InsertImage_2(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:210
// index:3
// Public Visibility=Default Availability=Available
// [-2] void insertImage(const QImage &, const QString &)

/*
Inserts the image defined by format at the current position().
*/
func (this *QTextCursor) InsertImage_3(image QImage_ITF, name string) {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK6QImageRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:210
// index:3
// Public Visibility=Default Availability=Available
// [-2] void insertImage(const QImage &, const QString &)

/*
Inserts the image defined by format at the current position().
*/
func (this *QTextCursor) InsertImage_3_(image QImage_ITF) {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record,
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK6QImageRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:212
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginEditBlock()

/*
Indicates the start of a block of editing operations on the document that should appear as a single operation from an undo/redo point of view.

For example:


  QTextCursor cursor(textDocument);
  cursor.beginEditBlock();
  cursor.insertText("Hello");
  cursor.insertText("World");
  cursor.endEditBlock();

  textDocument->undo();



The call to undo() will cause both insertions to be undone, causing both "World" and "Hello" to be removed.

It is possible to nest calls to beginEditBlock and endEditBlock. The top-most pair will determine the scope of the undo/redo operation.

See also endEditBlock().
*/
func (this *QTextCursor) BeginEditBlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor14beginEditBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void joinPreviousEditBlock()

/*
Like beginEditBlock() indicates the start of a block of editing operations that should appear as a single operation for undo/redo. However unlike beginEditBlock() it does not start a new block but reverses the previous call to endEditBlock() and therefore makes following operations part of the previous edit block created.

For example:


  QTextCursor cursor(textDocument);
  cursor.beginEditBlock();
  cursor.insertText("Hello");
  cursor.insertText("World");
  cursor.endEditBlock();

  ...

  cursor.joinPreviousEditBlock();
  cursor.insertText("Hey");
  cursor.endEditBlock();

  textDocument->undo();



The call to undo() will cause all three insertions to be undone.

See also beginEditBlock() and endEditBlock().
*/
func (this *QTextCursor) JoinPreviousEditBlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor21joinPreviousEditBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endEditBlock()

/*
Indicates the end of a block of editing operations on the document that should appear as a single operation from an undo/redo point of view.

See also beginEditBlock().
*/
func (this *QTextCursor) EndEditBlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor12endEditBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:216
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QTextCursor &) const

/*

 */
func (this *QTextCursor) Operator_not_equal(rhs QTextCursor_ITF) bool {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextCursor_PTR() != nil {
		convArg0 = rhs.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursorneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:217
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<(const QTextCursor &) const

/*

 */
func (this *QTextCursor) Operator_less_than(rhs QTextCursor_ITF) bool {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextCursor_PTR() != nil {
		convArg0 = rhs.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursorltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:218
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<=(const QTextCursor &) const

/*

 */
func (this *QTextCursor) Operator_less_than_equal(rhs QTextCursor_ITF) bool {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextCursor_PTR() != nil {
		convArg0 = rhs.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursorleERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:219
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QTextCursor &) const

/*

 */
func (this *QTextCursor) Operator_equal_equal(rhs QTextCursor_ITF) bool {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextCursor_PTR() != nil {
		convArg0 = rhs.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursoreqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:220
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator>=(const QTextCursor &) const

/*

 */
func (this *QTextCursor) Operator_greater_than_equal(rhs QTextCursor_ITF) bool {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextCursor_PTR() != nil {
		convArg0 = rhs.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursorgeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:221
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator>(const QTextCursor &) const

/*

 */
func (this *QTextCursor) Operator_greater_than(rhs QTextCursor_ITF) bool {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextCursor_PTR() != nil {
		convArg0 = rhs.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursorgtERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:223
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCopyOf(const QTextCursor &) const

/*
Returns true if this cursor and other are copies of each other, i.e. one of them was created as a copy of the other and neither has moved since. This is much stricter than equality.

See also operator=() and operator==().
*/
func (this *QTextCursor) IsCopyOf(other QTextCursor_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextCursor_PTR() != nil {
		convArg0 = other.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor8isCopyOfERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:225
// index:0
// Public Visibility=Default Availability=Available
// [4] int blockNumber() const

/*
Returns the number of the block the cursor is in, or 0 if the cursor is invalid.

Note that this function only makes sense in documents without complex objects such as tables or frames.

This function was introduced in  Qt 4.2.
*/
func (this *QTextCursor) BlockNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor11blockNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:226
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnNumber() const

/*
Returns the position of the cursor within its containing line.

Note that this is the column number relative to a wrapped line, not relative to the block (i.e. the paragraph).

You probably want to call positionInBlock() instead.

This function was introduced in  Qt 4.2.

See also positionInBlock().
*/
func (this *QTextCursor) ColumnNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12columnNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document() const

/*
Returns the document this cursor is associated with.

This function was introduced in  Qt 4.5.
*/
func (this *QTextCursor) Document() *QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

/*


If the anchor() is kept where it is and the position() is moved, the text in between will be selected.

*/
type QTextCursor__MoveMode = int

// Moves the anchor to the same position as the cursor itself.
const QTextCursor__MoveAnchor QTextCursor__MoveMode = 0

// Keeps the anchor where it is.
const QTextCursor__KeepAnchor QTextCursor__MoveMode = 1

/*


See also movePosition().

*/
type QTextCursor__MoveOperation = int

// Keep the cursor where it is
const QTextCursor__NoMove QTextCursor__MoveOperation = 0

// Move to the start of the document.
const QTextCursor__Start QTextCursor__MoveOperation = 1

// Move up one line.
const QTextCursor__Up QTextCursor__MoveOperation = 2

// Move to the start of the current line.
const QTextCursor__StartOfLine QTextCursor__MoveOperation = 3

// Move to the start of the current block.
const QTextCursor__StartOfBlock QTextCursor__MoveOperation = 4

// Move to the start of the current word.
const QTextCursor__StartOfWord QTextCursor__MoveOperation = 5

// Move to the start of the previous block.
const QTextCursor__PreviousBlock QTextCursor__MoveOperation = 6

// Move to the previous character.
const QTextCursor__PreviousCharacter QTextCursor__MoveOperation = 7

// Move to the beginning of the previous word.
const QTextCursor__PreviousWord QTextCursor__MoveOperation = 8

// Move left one character.
const QTextCursor__Left QTextCursor__MoveOperation = 9

//
const QTextCursor__WordLeft QTextCursor__MoveOperation = 10

//
const QTextCursor__End QTextCursor__MoveOperation = 11

//
const QTextCursor__Down QTextCursor__MoveOperation = 12

//
const QTextCursor__EndOfLine QTextCursor__MoveOperation = 13

//
const QTextCursor__EndOfWord QTextCursor__MoveOperation = 14

//
const QTextCursor__EndOfBlock QTextCursor__MoveOperation = 15

//
const QTextCursor__NextBlock QTextCursor__MoveOperation = 16

//
const QTextCursor__NextCharacter QTextCursor__MoveOperation = 17

//
const QTextCursor__NextWord QTextCursor__MoveOperation = 18

//
const QTextCursor__Right QTextCursor__MoveOperation = 19

//
const QTextCursor__WordRight QTextCursor__MoveOperation = 20

//
const QTextCursor__NextCell QTextCursor__MoveOperation = 21

//
const QTextCursor__PreviousCell QTextCursor__MoveOperation = 22

//
const QTextCursor__NextRow QTextCursor__MoveOperation = 23

//
const QTextCursor__PreviousRow QTextCursor__MoveOperation = 24

/*
This enum describes the types of selection that can be applied with the select() function.


*/
type QTextCursor__SelectionType = int

// Selects the word under the cursor. If the cursor is not positioned within a string of selectable characters, no text is selected.
const QTextCursor__WordUnderCursor QTextCursor__SelectionType = 0

// Selects the line of text under the cursor.
const QTextCursor__LineUnderCursor QTextCursor__SelectionType = 1

// Selects the block of text under the cursor.
const QTextCursor__BlockUnderCursor QTextCursor__SelectionType = 2

// Selects the entire document.
const QTextCursor__Document QTextCursor__SelectionType = 3

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
}

//  keep block end
