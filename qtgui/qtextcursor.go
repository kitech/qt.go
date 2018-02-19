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
func (this *QTextCursor) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(int, enum QTextCursor::MoveMode)
func (this *QTextCursor) SetPosition(pos int, mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11setPositionEiNS_8MoveModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(int, enum QTextCursor::MoveMode)
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
func (this *QTextCursor) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int positionInBlock() const
func (this *QTextCursor) PositionInBlock() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor15positionInBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] int anchor() const
func (this *QTextCursor) Anchor() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor6anchorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertText(const QString &)
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
func (this *QTextCursor) MovePosition(op int, arg1 int, n int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor12movePositionENS_13MoveOperationENS_8MoveModeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), op, arg1, n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool movePosition(enum QTextCursor::MoveOperation, enum QTextCursor::MoveMode, int)
func (this *QTextCursor) MovePosition__(op int) bool {
	// arg: 1, QTextCursor::MoveMode=Enum, QTextCursor::MoveMode=Enum,
	arg1 := 0
	// arg: 2, int=Int, =Invalid,
	n := 1
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor12movePositionENS_13MoveOperationENS_8MoveModeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), op, arg1, n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool movePosition(enum QTextCursor::MoveOperation, enum QTextCursor::MoveMode, int)
func (this *QTextCursor) MovePosition__1(op int, arg1 int) bool {
	// arg: 2, int=Int, =Invalid,
	n := 1
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor12movePositionENS_13MoveOperationENS_8MoveModeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), op, arg1, n)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool visualNavigation() const
func (this *QTextCursor) VisualNavigation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor16visualNavigationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisualNavigation(_Bool)
func (this *QTextCursor) SetVisualNavigation(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor19setVisualNavigationEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalMovementX(int)
func (this *QTextCursor) SetVerticalMovementX(x int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor20setVerticalMovementXEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int verticalMovementX() const
func (this *QTextCursor) VerticalMovementX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor17verticalMovementXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeepPositionOnInsert(_Bool)
func (this *QTextCursor) SetKeepPositionOnInsert(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor23setKeepPositionOnInsertEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool keepPositionOnInsert() const
func (this *QTextCursor) KeepPositionOnInsert() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor20keepPositionOnInsertEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deleteChar()
func (this *QTextCursor) DeleteChar() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10deleteCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deletePreviousChar()
func (this *QTextCursor) DeletePreviousChar() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor18deletePreviousCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void select(enum QTextCursor::SelectionType)
func (this *QTextCursor) Select(selection int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor6selectENS_13SelectionTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), selection)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:153
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSelection() const
func (this *QTextCursor) HasSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12hasSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:154
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasComplexSelection() const
func (this *QTextCursor) HasComplexSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor19hasComplexSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeSelectedText()
func (this *QTextCursor) RemoveSelectedText() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor18removeSelectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearSelection()
func (this *QTextCursor) ClearSelection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor14clearSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionStart() const
func (this *QTextCursor) SelectionStart() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor14selectionStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:158
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionEnd() const
func (this *QTextCursor) SelectionEnd() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12selectionEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedText() const
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
func (this *QTextCursor) SelectedTableCells(firstRow unsafe.Pointer /*666*/, numRows unsafe.Pointer /*666*/, firstColumn unsafe.Pointer /*666*/, numColumns unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor18selectedTableCellsEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), firstRow, numRows, firstColumn, numColumns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:164
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock block() const
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
func (this *QTextCursor) AtBlockStart() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12atBlockStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:179
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atBlockEnd() const
func (this *QTextCursor) AtBlockEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor10atBlockEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:180
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atStart() const
func (this *QTextCursor) AtStart() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor7atStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:181
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atEnd() const
func (this *QTextCursor) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertBlock()
func (this *QTextCursor) InsertBlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:184
// index:1
// Public Visibility=Default Availability=Available
// [-2] void insertBlock(const QTextBlockFormat &)
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
func (this *QTextCursor) InsertList_1(style int) *QTextList /*777 QTextList **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10insertListEN15QTextListFormat5StyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:190
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextList * createList(const QTextListFormat &)
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
func (this *QTextCursor) CreateList_1(style int) *QTextList /*777 QTextList **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor10createListEN15QTextListFormat5StyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:192
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextList * currentList() const
func (this *QTextCursor) CurrentList() *QTextList /*777 QTextList **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor11currentListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:194
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextTable * insertTable(int, int, const QTextTableFormat &)
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
func (this *QTextCursor) InsertTable_1(rows int, cols int) *QTextTable /*777 QTextTable **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor11insertTableEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rows, cols)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextTableFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:196
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextTable * currentTable() const
func (this *QTextCursor) CurrentTable() *QTextTable /*777 QTextTable **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12currentTableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextTableFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:198
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextFrame * insertFrame(const QTextFrameFormat &)
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
func (this *QTextCursor) CurrentFrame() *QTextFrame /*777 QTextFrame **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12currentFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextcursor.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertFragment(const QTextDocumentFragment &)
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
func (this *QTextCursor) BeginEditBlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor14beginEditBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void joinPreviousEditBlock()
func (this *QTextCursor) JoinPreviousEditBlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor21joinPreviousEditBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endEditBlock()
func (this *QTextCursor) EndEditBlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextCursor12endEditBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:216
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QTextCursor &) const
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
func (this *QTextCursor) BlockNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor11blockNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:226
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnNumber() const
func (this *QTextCursor) ColumnNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor12columnNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextcursor.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document() const
func (this *QTextCursor) Document() *QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextCursor8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

type QTextCursor__MoveMode = int

const QTextCursor__MoveAnchor QTextCursor__MoveMode = 0
const QTextCursor__KeepAnchor QTextCursor__MoveMode = 1

type QTextCursor__MoveOperation = int

const QTextCursor__NoMove QTextCursor__MoveOperation = 0
const QTextCursor__Start QTextCursor__MoveOperation = 1
const QTextCursor__Up QTextCursor__MoveOperation = 2
const QTextCursor__StartOfLine QTextCursor__MoveOperation = 3
const QTextCursor__StartOfBlock QTextCursor__MoveOperation = 4
const QTextCursor__StartOfWord QTextCursor__MoveOperation = 5
const QTextCursor__PreviousBlock QTextCursor__MoveOperation = 6
const QTextCursor__PreviousCharacter QTextCursor__MoveOperation = 7
const QTextCursor__PreviousWord QTextCursor__MoveOperation = 8
const QTextCursor__Left QTextCursor__MoveOperation = 9
const QTextCursor__WordLeft QTextCursor__MoveOperation = 10
const QTextCursor__End QTextCursor__MoveOperation = 11
const QTextCursor__Down QTextCursor__MoveOperation = 12
const QTextCursor__EndOfLine QTextCursor__MoveOperation = 13
const QTextCursor__EndOfWord QTextCursor__MoveOperation = 14
const QTextCursor__EndOfBlock QTextCursor__MoveOperation = 15
const QTextCursor__NextBlock QTextCursor__MoveOperation = 16
const QTextCursor__NextCharacter QTextCursor__MoveOperation = 17
const QTextCursor__NextWord QTextCursor__MoveOperation = 18
const QTextCursor__Right QTextCursor__MoveOperation = 19
const QTextCursor__WordRight QTextCursor__MoveOperation = 20
const QTextCursor__NextCell QTextCursor__MoveOperation = 21
const QTextCursor__PreviousCell QTextCursor__MoveOperation = 22
const QTextCursor__NextRow QTextCursor__MoveOperation = 23
const QTextCursor__PreviousRow QTextCursor__MoveOperation = 24

type QTextCursor__SelectionType = int

const QTextCursor__WordUnderCursor QTextCursor__SelectionType = 0
const QTextCursor__LineUnderCursor QTextCursor__SelectionType = 1
const QTextCursor__BlockUnderCursor QTextCursor__SelectionType = 2
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
