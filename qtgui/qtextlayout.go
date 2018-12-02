package qtgui

// /usr/include/qt/QtGui/qtextlayout.h
// #include <qtextlayout.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QTextLayout struct {
	*qtrt.CObject
}
type QTextLayout_ITF interface {
	QTextLayout_PTR() *QTextLayout
}

func (ptr *QTextLayout) QTextLayout_PTR() *QTextLayout { return ptr }

func (this *QTextLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextLayout) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextLayoutFromPointer(cthis unsafe.Pointer) *QTextLayout {
	return &QTextLayout{&qtrt.CObject{cthis}}
}
func (*QTextLayout) NewFromPointer(cthis unsafe.Pointer) *QTextLayout {
	return NewQTextLayoutFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextlayout.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout()

/*
Constructs an empty text layout.

See also setText().
*/
func (*QTextLayout) NewForInherit() *QTextLayout {
	return NewQTextLayout()
}
func NewQTextLayout() *QTextLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLayout)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:109
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout(const QString &)

/*
Constructs an empty text layout.

See also setText().
*/
func (*QTextLayout) NewForInherit_1(text string) *QTextLayout {
	return NewQTextLayout_1(text)
}
func NewQTextLayout_1(text string) *QTextLayout {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLayout)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:110
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout(const QString &, const QFont &, QPaintDevice *)

/*
Constructs an empty text layout.

See also setText().
*/
func (*QTextLayout) NewForInherit_2(text string, font QFont_ITF, paintdevice QPaintDevice_ITF /*777 QPaintDevice **/) *QTextLayout {
	return NewQTextLayout_2(text, font, paintdevice)
}
func NewQTextLayout_2(text string, font QFont_ITF, paintdevice QPaintDevice_ITF /*777 QPaintDevice **/) *QTextLayout {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg1 = font.QFont_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if paintdevice != nil && paintdevice.QPaintDevice_PTR() != nil {
		convArg2 = paintdevice.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QStringRK5QFontP12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLayout)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:110
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout(const QString &, const QFont &, QPaintDevice *)

/*
Constructs an empty text layout.

See also setText().
*/
func (*QTextLayout) NewForInherit_2_(text string, font QFont_ITF) *QTextLayout {
	return NewQTextLayout_2_(text, font)
}
func NewQTextLayout_2_(text string, font QFont_ITF) *QTextLayout {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg1 = font.QFont_PTR().GetCthis()
	}
	// arg: 2, QPaintDevice *=Pointer, QPaintDevice=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK7QStringRK5QFontP12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLayout)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:111
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QTextLayout(const QTextBlock &)

/*
Constructs an empty text layout.

See also setText().
*/
func (*QTextLayout) NewForInherit_3(b QTextBlock_ITF) *QTextLayout {
	return NewQTextLayout_3(b)
}
func NewQTextLayout_3(b QTextBlock_ITF) *QTextLayout {
	var convArg0 unsafe.Pointer
	if b != nil && b.QTextBlock_PTR() != nil {
		convArg0 = b.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutC2ERK10QTextBlock", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLayout)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextLayout()

/*

 */
func DeleteQTextLayout(this *QTextLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextlayout.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)

/*
Sets the layout's font to the given font. The layout is invalidated and must be laid out again.

See also font().
*/
func (this *QTextLayout) SetFont(f QFont_ITF) {
	var convArg0 unsafe.Pointer
	if f != nil && f.QFont_PTR() != nil {
		convArg0 = f.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:115
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font() const

/*
Returns the current font that is used for the layout, or a default font if none is set.

See also setFont().
*/
func (this *QTextLayout) Font() *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawFont(const QRawFont &)

/*

 */
func (this *QTextLayout) SetRawFont(rawFont QRawFont_ITF) {
	var convArg0 unsafe.Pointer
	if rawFont != nil && rawFont.QRawFont_PTR() != nil {
		convArg0 = rawFont.QRawFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout10setRawFontERK8QRawFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*
Sets the layout's text to the given string. The layout is invalidated and must be laid out again.

Notice that when using this QTextLayout as part of a QTextDocument this method will have no effect.

See also text().
*/
func (this *QTextLayout) SetText(string string) {
	var tmpArg0 = qtcore.NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*
Returns the layout's text.

See also setText().
*/
func (this *QTextLayout) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextlayout.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextOption(const QTextOption &)

/*
Sets the text option structure that controls the layout process to the given option.

See also textOption().
*/
func (this *QTextLayout) SetTextOption(option QTextOption_ITF) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QTextOption_PTR() != nil {
		convArg0 = option.QTextOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout13setTextOptionERK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:125
// index:0
// Public Visibility=Default Availability=Available
// [32] const QTextOption & textOption() const

/*
Returns the current text option used to control the layout process.

See also setTextOption().
*/
func (this *QTextLayout) TextOption() *QTextOption {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout10textOptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextOption)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreeditArea(int, const QString &)

/*
Sets the position and text of the area in the layout that is processed before editing occurs. The layout is invalidated and must be laid out again.

See also preeditAreaPosition() and preeditAreaText().
*/
func (this *QTextLayout) SetPreeditArea(position int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout14setPreeditAreaEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] int preeditAreaPosition() const

/*
Returns the position of the area in the text layout that will be processed before editing occurs.

See also preeditAreaText().
*/
func (this *QTextLayout) PreeditAreaPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout19preeditAreaPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QString preeditAreaText() const

/*
Returns the text that is inserted in the layout before editing occurs.

See also preeditAreaPosition().
*/
func (this *QTextLayout) PreeditAreaText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout15preeditAreaTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextlayout.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearAdditionalFormats()

/*

 */
func (this *QTextLayout) ClearAdditionalFormats() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout22clearAdditionalFormatsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearFormats()

/*
Clears the list of additional formats supported by the text layout.

This function was introduced in  Qt 5.6.

See also formats() and setFormats().
*/
func (this *QTextLayout) ClearFormats() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout12clearFormatsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheEnabled(bool)

/*
Enables caching of the complete layout information if enable is true; otherwise disables layout caching. Usually QTextLayout throws most of the layouting information away after a call to endLayout() to reduce memory consumption. If you however want to draw the laid out text directly afterwards enabling caching might speed up drawing significantly.

See also cacheEnabled().
*/
func (this *QTextLayout) SetCacheEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout15setCacheEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool cacheEnabled() const

/*
Returns true if the complete layout information is cached; otherwise returns false.

See also setCacheEnabled().
*/
func (this *QTextLayout) CacheEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout12cacheEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorMoveStyle(Qt::CursorMoveStyle)

/*
Sets the visual cursor movement style to the given style. If the QTextLayout is backed by a document, you can ignore this and use the option in QTextDocument, this option is for widgets like QLineEdit or custom widgets without a QTextDocument. Default value is Qt::LogicalMoveStyle.

See also cursorMoveStyle().
*/
func (this *QTextLayout) SetCursorMoveStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout18setCursorMoveStyleEN2Qt15CursorMoveStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:154
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CursorMoveStyle cursorMoveStyle() const

/*
The cursor movement style of this QTextLayout. The default is Qt::LogicalMoveStyle.

See also setCursorMoveStyle().
*/
func (this *QTextLayout) CursorMoveStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout15cursorMoveStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginLayout()

/*
Begins the layout process.

Warning: This will invalidate the layout, so all existing QTextLine objects that refer to the previous contents should now be discarded.

See also endLayout().
*/
func (this *QTextLayout) BeginLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout11beginLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endLayout()

/*
Ends the layout process.

See also beginLayout().
*/
func (this *QTextLayout) EndLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout9endLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearLayout()

/*
Clears the line information in the layout. After having called this function, lineCount() returns 0.

Warning: This will invalidate the layout, so all existing QTextLine objects that refer to the previous contents should now be discarded.

This function was introduced in  Qt 4.4.
*/
func (this *QTextLayout) ClearLayout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout11clearLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:160
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLine createLine()

/*
Returns a new text line to be laid out if there is text to be inserted into the layout; otherwise returns an invalid text line.

The text layout creates a new line object that starts after the last line in the layout, or at the beginning if the layout is empty. The layout maintains an internal cursor, and each line is filled with text from the cursor position onwards when the QTextLine::setLineWidth() function is called.

Once QTextLine::setLineWidth() is called, a new line can be created and filled with text. Repeating this process will lay out the whole block of text contained in the QTextLayout. If there is no text left to be inserted into the layout, the QTextLine returned will not be valid (isValid() will return false).
*/
func (this *QTextLayout) CreateLine() *QTextLine /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout10createLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLine)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:162
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineCount() const

/*
Returns the number of lines in this text layout.

See also lineAt().
*/
func (this *QTextLayout) LineCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout9lineCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:163
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLine lineAt(int) const

/*
Returns the i-th line of text in this text layout.

See also lineCount() and lineForTextPosition().
*/
func (this *QTextLayout) LineAt(i int) *QTextLine /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout6lineAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLine)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:164
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLine lineForTextPosition(int) const

/*
Returns the line that contains the cursor position specified by pos.

See also isValidCursorPosition() and lineAt().
*/
func (this *QTextLayout) LineForTextPosition(pos int) *QTextLine /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout19lineForTextPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLine)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValidCursorPosition(int) const

/*
/ Returns true if position pos is a valid cursor position.

In a Unicode context some positions in the text are not valid cursor positions, because the position is inside a Unicode surrogate or a grapheme cluster.

A grapheme cluster is a sequence of two or more Unicode characters that form one indivisible entity on the screen. For example the latin character `Ä' can be represented in Unicode by two characters, `A' (0x41), and the combining diaresis (0x308). A text cursor can only validly be positioned before or after these two characters, never between them since that wouldn't make sense. In indic languages every syllable forms a grapheme cluster.
*/
func (this *QTextLayout) IsValidCursorPosition(pos int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout21isValidCursorPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:171
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextCursorPosition(int, QTextLayout::CursorMode) const

/*
Returns the next valid cursor position after oldPos that respects the given cursor mode. Returns value of oldPos, if oldPos is not a valid cursor position.

See also isValidCursorPosition() and previousCursorPosition().
*/
func (this *QTextLayout) NextCursorPosition(oldPos int, mode int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout18nextCursorPositionEiNS_10CursorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldPos, mode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:171
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextCursorPosition(int, QTextLayout::CursorMode) const

/*
Returns the next valid cursor position after oldPos that respects the given cursor mode. Returns value of oldPos, if oldPos is not a valid cursor position.

See also isValidCursorPosition() and previousCursorPosition().
*/
func (this *QTextLayout) NextCursorPosition__(oldPos int) int {
	// arg: 1, QTextLayout::CursorMode=Enum, QTextLayout::CursorMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout18nextCursorPositionEiNS_10CursorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldPos, mode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:172
// index:0
// Public Visibility=Default Availability=Available
// [4] int previousCursorPosition(int, QTextLayout::CursorMode) const

/*
Returns the first valid cursor position before oldPos that respects the given cursor mode. Returns value of oldPos, if oldPos is not a valid cursor position.

See also isValidCursorPosition() and nextCursorPosition().
*/
func (this *QTextLayout) PreviousCursorPosition(oldPos int, mode int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout22previousCursorPositionEiNS_10CursorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldPos, mode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:172
// index:0
// Public Visibility=Default Availability=Available
// [4] int previousCursorPosition(int, QTextLayout::CursorMode) const

/*
Returns the first valid cursor position before oldPos that respects the given cursor mode. Returns value of oldPos, if oldPos is not a valid cursor position.

See also isValidCursorPosition() and nextCursorPosition().
*/
func (this *QTextLayout) PreviousCursorPosition__(oldPos int) int {
	// arg: 1, QTextLayout::CursorMode=Enum, QTextLayout::CursorMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout22previousCursorPositionEiNS_10CursorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldPos, mode)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:173
// index:0
// Public Visibility=Default Availability=Available
// [4] int leftCursorPosition(int) const

/*
Returns the cursor position to the left of oldPos, next to it. It's dependent on the visual position of characters, after bi-directional reordering.

See also rightCursorPosition() and previousCursorPosition().
*/
func (this *QTextLayout) LeftCursorPosition(oldPos int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout18leftCursorPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldPos)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:174
// index:0
// Public Visibility=Default Availability=Available
// [4] int rightCursorPosition(int) const

/*
Returns the cursor position to the right of oldPos, next to it. It's dependent on the visual position of characters, after bi-directional reordering.

See also leftCursorPosition() and nextCursorPosition().
*/
func (this *QTextLayout) RightCursorPosition(oldPos int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout19rightCursorPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), oldPos)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawCursor(QPainter *, const QPointF &, int) const

/*
Draws a text cursor with the current pen and the specified width at the given position using the painter specified. The corresponding position within the text is specified by cursorPosition.
*/
func (this *QTextLayout) DrawCursor(p QPainter_ITF /*777 QPainter **/, pos qtcore.QPointF_ITF, cursorPosition int) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg1 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cursorPosition)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:179
// index:1
// Public Visibility=Default Availability=Available
// [-2] void drawCursor(QPainter *, const QPointF &, int, int) const

/*
Draws a text cursor with the current pen and the specified width at the given position using the painter specified. The corresponding position within the text is specified by cursorPosition.
*/
func (this *QTextLayout) DrawCursor_1(p QPainter_ITF /*777 QPainter **/, pos qtcore.QPointF_ITF, cursorPosition int, width int) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg1 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout10drawCursorEP8QPainterRK7QPointFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, cursorPosition, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:181
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF position() const

/*
The global position of the layout. This is independent of the bounding rectangle and of the layout process.

This function was introduced in  Qt 4.2.

See also setPosition().
*/
func (this *QTextLayout) Position() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPointF &)

/*
Moves the text layout to point p.

See also position().
*/
func (this *QTextLayout) SetPosition(p qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPointF_PTR() != nil {
		convArg0 = p.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout11setPositionERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:184
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect() const

/*
The smallest rectangle that contains all the lines in the layout.
*/
func (this *QTextLayout) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:186
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal minimumWidth() const

/*
The minimum width the layout needs. This is the width of the layout's smallest non-breakable substring.

Warning: This function only returns a valid value after the layout has been done.

See also maximumWidth().
*/
func (this *QTextLayout) MinimumWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout12minimumWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:187
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maximumWidth() const

/*
The maximum width the layout could expand to; this is essentially the width of the entire text.

Warning: This function only returns a valid value after the layout has been done.

See also minimumWidth().
*/
func (this *QTextLayout) MaximumWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLayout12maximumWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(int)

/*

 */
func (this *QTextLayout) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLayout8setFlagsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

/*
ConstantValue
QTextLayout::SkipCharacters0
QTextLayout::SkipWords1

*/
type QTextLayout__CursorMode = int

//
const QTextLayout__SkipCharacters QTextLayout__CursorMode = 0

//
const QTextLayout__SkipWords QTextLayout__CursorMode = 1

func (this *QTextLayout) CursorModeItemName(val int) string {
	switch val {
	case QTextLayout__SkipCharacters: // 0
		return "SkipCharacters"
	case QTextLayout__SkipWords: // 1
		return "SkipWords"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextLayout_CursorModeItemName(val int) string {
	var nilthis *QTextLayout
	return nilthis.CursorModeItemName(val)
}

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
