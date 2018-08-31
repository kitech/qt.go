package qtwidgets

// /usr/include/qt/QtWidgets/qlineedit.h
// #include <qlineedit.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 109
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

// void mousePressEvent(QMouseEvent *)
func (this *QLineEdit) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QLineEdit) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QLineEdit) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QLineEdit) InheritMouseDoubleClickEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QLineEdit) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QLineEdit) InheritFocusInEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QLineEdit) InheritFocusOutEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QLineEdit) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void dragEnterEvent(QDragEnterEvent *)
func (this *QLineEdit) InheritDragEnterEvent(f func(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(QDragMoveEvent *)
func (this *QLineEdit) InheritDragMoveEvent(f func(e *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QLineEdit) InheritDragLeaveEvent(f func(e *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(QDropEvent *)
func (this *QLineEdit) InheritDropEvent(f func(arg0 *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void changeEvent(QEvent *)
func (this *QLineEdit) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QLineEdit) InheritContextMenuEvent(f func(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QLineEdit) InheritInputMethodEvent(f func(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void initStyleOption(QStyleOptionFrame *)
func (this *QLineEdit) InheritInitStyleOption(f func(option *QStyleOptionFrame /*777 QStyleOptionFrame **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

// QRect cursorRect()
func (this *QLineEdit) InheritCursorRect(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "cursorRect", f)
}

/*

 */
type QLineEdit struct {
	*QWidget
}
type QLineEdit_ITF interface {
	QWidget_ITF
	QLineEdit_PTR() *QLineEdit
}

func (ptr *QLineEdit) QLineEdit_PTR() *QLineEdit { return ptr }

func (this *QLineEdit) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QLineEdit) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQLineEditFromPointer(cthis unsafe.Pointer) *QLineEdit {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QLineEdit{bcthis0}
}
func (*QLineEdit) NewFromPointer(cthis unsafe.Pointer) *QLineEdit {
	return NewQLineEditFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlineedit.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QLineEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlineedit.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLineEdit(QWidget *)

/*
Constructs a line edit with no text.

The maximum text length is set to 32767 characters.

The parent argument is sent to the QWidget constructor.

See also setText() and setMaxLength().
*/
func (*QLineEdit) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QLineEdit {
	return NewQLineEdit(parent)
}
func NewQLineEdit(parent QWidget_ITF /*777 QWidget **/) *QLineEdit {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLineEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qlineedit.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLineEdit(QWidget *)

/*
Constructs a line edit with no text.

The maximum text length is set to 32767 characters.

The parent argument is sent to the QWidget constructor.

See also setText() and setMaxLength().
*/
func (*QLineEdit) NewForInherit__() *QLineEdit {
	return NewQLineEdit__()
}
func NewQLineEdit__() *QLineEdit {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLineEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qlineedit.h:94
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLineEdit(const QString &, QWidget *)

/*
Constructs a line edit with no text.

The maximum text length is set to 32767 characters.

The parent argument is sent to the QWidget constructor.

See also setText() and setMaxLength().
*/
func (*QLineEdit) NewForInherit_1(arg0 string, parent QWidget_ITF /*777 QWidget **/) *QLineEdit {
	return NewQLineEdit_1(arg0, parent)
}
func NewQLineEdit_1(arg0 string, parent QWidget_ITF /*777 QWidget **/) *QLineEdit {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEditC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLineEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qlineedit.h:94
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLineEdit(const QString &, QWidget *)

/*
Constructs a line edit with no text.

The maximum text length is set to 32767 characters.

The parent argument is sent to the QWidget constructor.

See also setText() and setMaxLength().
*/
func (*QLineEdit) NewForInherit_1_(arg0 string) *QLineEdit {
	return NewQLineEdit_1_(arg0)
}
func NewQLineEdit_1_(arg0 string) *QLineEdit {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEditC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLineEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qlineedit.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLineEdit()

/*

 */
func DeleteQLineEdit(this *QLineEdit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEditD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlineedit.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*

 */
func (this *QLineEdit) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlineedit.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QString displayText() const

/*

 */
func (this *QLineEdit) DisplayText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit11displayTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlineedit.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QString placeholderText() const

/*

 */
func (this *QLineEdit) PlaceholderText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit15placeholderTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlineedit.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlaceholderText(const QString &)

/*

 */
func (this *QLineEdit) SetPlaceholderText(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit18setPlaceholderTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxLength() const

/*

 */
func (this *QLineEdit) MaxLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit9maxLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxLength(int)

/*

 */
func (this *QLineEdit) SetMaxLength(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit12setMaxLengthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrame(bool)

/*

 */
func (this *QLineEdit) SetFrame(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit8setFrameEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFrame() const

/*

 */
func (this *QLineEdit) HasFrame() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit8hasFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClearButtonEnabled(bool)

/*

 */
func (this *QLineEdit) SetClearButtonEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit21setClearButtonEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isClearButtonEnabled() const

/*

 */
func (this *QLineEdit) IsClearButtonEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit20isClearButtonEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QLineEdit::EchoMode echoMode() const

/*

 */
func (this *QLineEdit) EchoMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit8echoModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEchoMode(QLineEdit::EchoMode)

/*

 */
func (this *QLineEdit) SetEchoMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit11setEchoModeENS_8EchoModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly() const

/*

 */
func (this *QLineEdit) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(bool)

/*

 */
func (this *QLineEdit) SetReadOnly(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit11setReadOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValidator(const QValidator *)

/*
Sets this line edit to only accept input that the validator, v, will accept. This allows you to place any arbitrary constraints on the text which may be entered.

If v == 0, setValidator() removes the current input validator. The initial setting is to have no input validator (i.e. any input is accepted up to maxLength()).

See also validator(), hasAcceptableInput(), QIntValidator, QDoubleValidator, and QRegExpValidator.
*/
func (this *QLineEdit) SetValidator(arg0 qtgui.QValidator_ITF /*777 const QValidator **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QValidator_PTR() != nil {
		convArg0 = arg0.QValidator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit12setValidatorEPK10QValidator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] const QValidator * validator() const

/*
Returns a pointer to the current input validator, or 0 if no validator has been set.

See also setValidator().
*/
func (this *QLineEdit) Validator() *qtgui.QValidator /*777 const QValidator **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit9validatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQValidatorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlineedit.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompleter(QCompleter *)

/*
Sets this line edit to provide auto completions from the completer, c. The completion mode is set using QCompleter::setCompletionMode().

To use a QCompleter with a QValidator or QLineEdit::inputMask, you need to ensure that the model provided to QCompleter contains valid entries. You can use the QSortFilterProxyModel to ensure that the QCompleter's model contains only valid entries.

If c == 0, setCompleter() removes the current completer, effectively disabling auto completion.

This function was introduced in  Qt 4.2.

See also completer() and QCompleter.
*/
func (this *QLineEdit) SetCompleter(completer QCompleter_ITF /*777 QCompleter **/) {
	var convArg0 unsafe.Pointer
	if completer != nil && completer.QCompleter_PTR() != nil {
		convArg0 = completer.QCompleter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit12setCompleterEP10QCompleter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QCompleter * completer() const

/*
Returns the current QCompleter that provides completions.

This function was introduced in  Qt 4.2.

See also setCompleter().
*/
func (this *QLineEdit) Completer() *QCompleter /*777 QCompleter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit9completerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlineedit.h:131
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().

Returns a recommended size for the widget.

The width returned, in pixels, is usually enough for about 15 to 20 characters.
*/
func (this *QLineEdit) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:132
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().

Returns a minimum size for the line edit.

The width returned is enough for at least one character.
*/
func (this *QLineEdit) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int cursorPosition() const

/*

 */
func (this *QLineEdit) CursorPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit14cursorPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorPosition(int)

/*

 */
func (this *QLineEdit) SetCursorPosition(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit17setCursorPositionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] int cursorPositionAt(const QPoint &)

/*
Returns the cursor position under the point pos.
*/
func (this *QLineEdit) CursorPositionAt(pos qtcore.QPoint_ITF) int {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit16cursorPositionAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)

/*

 */
func (this *QLineEdit) SetAlignment(flag int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const

/*

 */
func (this *QLineEdit) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorForward(bool, int)

/*
Moves the cursor forward steps characters. If mark is true each character moved over is added to the selection; if mark is false the selection is cleared.

See also cursorBackward().
*/
func (this *QLineEdit) CursorForward(mark bool, steps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit13cursorForwardEbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mark, steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorForward(bool, int)

/*
Moves the cursor forward steps characters. If mark is true each character moved over is added to the selection; if mark is false the selection is cleared.

See also cursorBackward().
*/
func (this *QLineEdit) CursorForward__(mark bool) {
	// arg: 1, int=Int, =Invalid, , Invalid
	steps := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit13cursorForwardEbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mark, steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorBackward(bool, int)

/*
Moves the cursor back steps characters. If mark is true each character moved over is added to the selection; if mark is false the selection is cleared.

See also cursorForward().
*/
func (this *QLineEdit) CursorBackward(mark bool, steps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit14cursorBackwardEbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mark, steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorBackward(bool, int)

/*
Moves the cursor back steps characters. If mark is true each character moved over is added to the selection; if mark is false the selection is cleared.

See also cursorForward().
*/
func (this *QLineEdit) CursorBackward__(mark bool) {
	// arg: 1, int=Int, =Invalid, , Invalid
	steps := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit14cursorBackwardEbi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mark, steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorWordForward(bool)

/*
Moves the cursor one word forward. If mark is true, the word is also selected.

See also cursorWordBackward().
*/
func (this *QLineEdit) CursorWordForward(mark bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit17cursorWordForwardEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mark)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorWordBackward(bool)

/*
Moves the cursor one word backward. If mark is true, the word is also selected.

See also cursorWordForward().
*/
func (this *QLineEdit) CursorWordBackward(mark bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit18cursorWordBackwardEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mark)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void backspace()

/*
If no text is selected, deletes the character to the left of the text cursor and moves the cursor one position to the left. If any text is selected, the cursor is moved to the beginning of the selected text and the selected text is deleted.

See also del().
*/
func (this *QLineEdit) Backspace() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit9backspaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void del()

/*
If no text is selected, deletes the character to the right of the text cursor. If any text is selected, the cursor is moved to the beginning of the selected text and the selected text is deleted.

See also backspace().
*/
func (this *QLineEdit) Del() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit3delEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void home(bool)

/*
Moves the text cursor to the beginning of the line unless it is already there. If mark is true, text is selected towards the first position; otherwise, any selected text is unselected if the cursor is moved.

See also end().
*/
func (this *QLineEdit) Home(mark bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit4homeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mark)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void end(bool)

/*
Moves the text cursor to the end of the line unless it is already there. If mark is true, text is selected towards the last position; otherwise, any selected text is unselected if the cursor is moved.

See also home().
*/
func (this *QLineEdit) End(mark bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit3endEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mark)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:150
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isModified() const

/*

 */
func (this *QLineEdit) IsModified() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit10isModifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModified(bool)

/*

 */
func (this *QLineEdit) SetModified(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit11setModifiedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelection(int, int)

/*
Selects text from position start and for length characters. Negative lengths are allowed.

See also deselect(), selectAll(), and selectedText().
*/
func (this *QLineEdit) SetSelection(arg0 int, arg1 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit12setSelectionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:154
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSelectedText() const

/*

 */
func (this *QLineEdit) HasSelectedText() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit15hasSelectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedText() const

/*

 */
func (this *QLineEdit) SelectedText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit12selectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlineedit.h:156
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionStart() const

/*
Returns the index of the first selected character in the line edit or -1 if no text is selected.

See also selectedText(), selectionEnd(), and selectionLength().
*/
func (this *QLineEdit) SelectionStart() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit14selectionStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionEnd() const

/*
Returns the index of the character directly after the selection in the line edit or -1 if no text is selected.

This function was introduced in  Qt 5.10.

See also selectedText(), selectionStart(), and selectionLength().
*/
func (this *QLineEdit) SelectionEnd() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit12selectionEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:158
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionLength() const

/*
Returns the length of the selection.

This function was introduced in  Qt 5.10.

See also selectedText(), selectionStart(), and selectionEnd().
*/
func (this *QLineEdit) SelectionLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit15selectionLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUndoAvailable() const

/*

 */
func (this *QLineEdit) IsUndoAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit15isUndoAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:161
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRedoAvailable() const

/*

 */
func (this *QLineEdit) IsRedoAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit15isRedoAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragEnabled(bool)

/*

 */
func (this *QLineEdit) SetDragEnabled(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit14setDragEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dragEnabled() const

/*

 */
func (this *QLineEdit) DragEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit11dragEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorMoveStyle(Qt::CursorMoveStyle)

/*

 */
func (this *QLineEdit) SetCursorMoveStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit18setCursorMoveStyleEN2Qt15CursorMoveStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CursorMoveStyle cursorMoveStyle() const

/*

 */
func (this *QLineEdit) CursorMoveStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit15cursorMoveStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QString inputMask() const

/*

 */
func (this *QLineEdit) InputMask() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit9inputMaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qlineedit.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputMask(const QString &)

/*

 */
func (this *QLineEdit) SetInputMask(inputMask string) {
	var tmpArg0 = qtcore.NewQString_5(inputMask)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit12setInputMaskERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:171
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAcceptableInput() const

/*

 */
func (this *QLineEdit) HasAcceptableInput() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit18hasAcceptableInputEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextMargins(int, int, int, int)

/*
Sets the margins around the text inside the frame to have the sizes left, top, right, and bottom.

See also getTextMargins().

This function was introduced in  Qt 4.5.

See also textMargins().
*/
func (this *QLineEdit) SetTextMargins(left int, top int, right int, bottom int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit14setTextMarginsEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:174
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setTextMargins(const QMargins &)

/*
Sets the margins around the text inside the frame to have the sizes left, top, right, and bottom.

See also getTextMargins().

This function was introduced in  Qt 4.5.

See also textMargins().
*/
func (this *QLineEdit) SetTextMargins_1(margins qtcore.QMargins_ITF) {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg0 = margins.QMargins_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit14setTextMarginsERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getTextMargins(int *, int *, int *, int *) const

/*
Returns the widget's text margins for left, top, right, and bottom.

This function was introduced in  Qt 4.5.

See also setTextMargins().
*/
func (this *QLineEdit) GetTextMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit14getTextMarginsEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:176
// index:0
// Public Visibility=Default Availability=Available
// [16] QMargins textMargins() const

/*
Returns the widget's text margins.

This function was introduced in  Qt 4.6.

See also setTextMargins().
*/
func (this *QLineEdit) TextMargins() *qtcore.QMargins /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit11textMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAction(QAction *, QLineEdit::ActionPosition)

/*
This is an overloaded function.

Adds the action to the list of actions at the position.

This function was introduced in  Qt 5.2.
*/
func (this *QLineEdit) AddAction(action QAction_ITF /*777 QAction **/, position int) {
	var convArg0 unsafe.Pointer
	if action != nil && action.QAction_PTR() != nil {
		convArg0 = action.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit9addActionEP7QActionNS_14ActionPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:181
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, QLineEdit::ActionPosition)

/*
This is an overloaded function.

Adds the action to the list of actions at the position.

This function was introduced in  Qt 5.2.
*/
func (this *QLineEdit) AddAction_1(icon qtgui.QIcon_ITF, position int) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit9addActionERK5QIconNS_14ActionPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, position)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlineedit.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*

 */
func (this *QLineEdit) SetText(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the contents of the line edit.

See also setText() and insert().
*/
func (this *QLineEdit) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectAll()

/*
Selects all the text (i.e. highlights it) and moves the cursor to the end. This is useful when a default value has been inserted because if the user types before clicking on the widget, the selected text will be deleted.

See also setSelection() and deselect().
*/
func (this *QLineEdit) SelectAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit9selectAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undo()

/*
Undoes the last operation if undo is available. Deselects any current selection, and updates the selection start to the current cursor position.
*/
func (this *QLineEdit) Undo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit4undoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redo()

/*
Redoes the last operation if redo is available.
*/
func (this *QLineEdit) Redo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit4redoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cut()

/*
Copies the selected text to the clipboard and deletes it, if there is any, and if echoMode() is Normal.

If the current validator disallows deleting the selected text, cut() will copy without deleting.

See also copy(), paste(), and setValidator().
*/
func (this *QLineEdit) Cut() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit3cutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void copy() const

/*
Copies the selected text to the clipboard, if there is any, and if echoMode() is Normal.

See also cut() and paste().
*/
func (this *QLineEdit) Copy() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit4copyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paste()

/*
Inserts the clipboard's text at the cursor position, deleting any selected text, providing the line edit is not read-only.

If the end result would not be acceptable to the current validator, nothing happens.

See also copy() and cut().
*/
func (this *QLineEdit) Paste() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit5pasteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deselect()

/*
Deselects any selected text.

See also setSelection() and selectAll().
*/
func (this *QLineEdit) Deselect() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit8deselectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * createStandardContextMenu()

/*
This function creates the standard context menu which is shown when the user clicks on the line edit with the right mouse button. It is called from the default contextMenuEvent() handler. The popup menu's ownership is transferred to the caller.
*/
func (this *QLineEdit) CreateStandardContextMenu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit25createStandardContextMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlineedit.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textChanged(const QString &)

/*
This signal is emitted whenever the text changes. The text argument is the new text.

Unlike textEdited(), this signal is also emitted when the text is changed programmatically, for example, by calling setText().

Note: Notifier signal for property text.
*/
func (this *QLineEdit) TextChanged(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit11textChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textEdited(const QString &)

/*
This signal is emitted whenever the text is edited. The text argument is the new text.

Unlike textChanged(), this signal is not emitted when the text is changed programmatically, for example, by calling setText().
*/
func (this *QLineEdit) TextEdited(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit10textEditedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorPositionChanged(int, int)

/*
This signal is emitted whenever the cursor moves. The previous position is given by old, and the new position by new.

See also setCursorPosition() and cursorPosition().
*/
func (this *QLineEdit) CursorPositionChanged(arg0 int, arg1 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit21cursorPositionChangedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void returnPressed()

/*
This signal is emitted when the Return or Enter key is pressed. Note that if there is a validator() or inputMask() set on the line edit, the returnPressed() signal will only be emitted if the input follows the inputMask() and the validator() returns QValidator::Acceptable.
*/
func (this *QLineEdit) ReturnPressed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit13returnPressedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editingFinished()

/*
This signal is emitted when the Return or Enter key is pressed or the line edit loses focus. Note that if there is a validator() or inputMask() set on the line edit and enter/return is pressed, the editingFinished() signal will only be emitted if the input follows the inputMask() and the validator() returns QValidator::Acceptable.
*/
func (this *QLineEdit) EditingFinished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit15editingFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:209
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()

/*
This signal is emitted whenever the selection changes.

See also hasSelectedText() and selectedText().
*/
func (this *QLineEdit) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:212
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QLineEdit) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:213
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QLineEdit) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:214
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QLineEdit) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:215
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseDoubleClickEvent().
*/
func (this *QLineEdit) MouseDoubleClickEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:216
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().

Converts the given key press event into a line edit action.

If Return or Enter is pressed and the current text is valid (or can be made valid by the validator), the signal returnPressed() is emitted.

The default key bindings are listed in the class's detailed description.
*/
func (this *QLineEdit) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:217
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().
*/
func (this *QLineEdit) FocusInEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().
*/
func (this *QLineEdit) FocusOutEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QLineEdit) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
Reimplemented from QWidget::dragEnterEvent().
*/
func (this *QLineEdit) DragEnterEvent(arg0 qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragEnterEvent_PTR() != nil {
		convArg0 = arg0.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
Reimplemented from QWidget::dragMoveEvent().
*/
func (this *QLineEdit) DragMoveEvent(e qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragMoveEvent_PTR() != nil {
		convArg0 = e.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:223
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
Reimplemented from QWidget::dragLeaveEvent().
*/
func (this *QLineEdit) DragLeaveEvent(e qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragLeaveEvent_PTR() != nil {
		convArg0 = e.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:224
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
Reimplemented from QWidget::dropEvent().
*/
func (this *QLineEdit) DropEvent(arg0 qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDropEvent_PTR() != nil {
		convArg0 = arg0.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:226
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QLineEdit) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:228
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*
Reimplemented from QWidget::contextMenuEvent().

Shows the standard context menu created with createStandardContextMenu().

If you do not want the line edit to have a context menu, you can set its contextMenuPolicy to Qt::NoContextMenu. If you want to customize the context menu, reimplement this function. If you want to extend the standard context menu, reimplement this function, call createStandardContextMenu() and extend the menu returned.


  void LineEdit::contextMenuEvent(QContextMenuEvent *event)
  {
      QMenu *menu = createStandardContextMenu();
      menu->addAction(tr("My Menu Item"));
      //...
      menu->exec(event->globalPos());
      delete menu;
  }



The event parameter is used to obtain the position where the mouse cursor was when the event was generated.

See also setContextMenuPolicy().
*/
func (this *QLineEdit) ContextMenuEvent(arg0 qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QContextMenuEvent_PTR() != nil {
		convArg0 = arg0.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:231
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
Reimplemented from QWidget::inputMethodEvent().
*/
func (this *QLineEdit) InputMethodEvent(arg0 qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QInputMethodEvent_PTR() != nil {
		convArg0 = arg0.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:232
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionFrame *) const

/*
Initialize option with the values from this QLineEdit. This method is useful for subclasses when they need a QStyleOptionFrame, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QLineEdit) InitStyleOption(option QStyleOptionFrame_ITF /*777 QStyleOptionFrame **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionFrame_PTR() != nil {
		convArg0 = option.QStyleOptionFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit15initStyleOptionEP17QStyleOptionFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:234
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
Reimplemented from QWidget::inputMethodQuery().
*/
func (this *QLineEdit) InputMethodQuery(arg0 int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:235
// index:1
// Public Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery, QVariant) const

/*
Reimplemented from QWidget::inputMethodQuery().
*/
func (this *QLineEdit) InputMethodQuery_1(property int, argument qtcore.QVariant_ITF /*123*/) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if argument != nil && argument.QVariant_PTR() != nil {
		convArg1 = argument.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit16inputMethodQueryEN2Qt16InputMethodQueryE8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), property, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:236
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QLineEdit) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLineEdit5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:238
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect cursorRect() const

/*
Returns a rectangle that includes the lineedit cursor.

This function was introduced in  Qt 4.4.
*/
func (this *QLineEdit) CursorRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLineEdit10cursorRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

/*
This enum type describes how a line edit should display the action widgets to be added.



This enum was introduced or modified in  Qt 5.2.

See also addAction(), removeAction(), and QWidget::layoutDirection.

*/
type QLineEdit__ActionPosition = int

// The widget is displayed to the left of the text when using layout direction Qt::LeftToRight or to the right when using Qt::RightToLeft, respectively.
const QLineEdit__LeadingPosition QLineEdit__ActionPosition = 0

// The widget is displayed to the right of the text when using layout direction Qt::LeftToRight or to the left when using Qt::RightToLeft, respectively.
const QLineEdit__TrailingPosition QLineEdit__ActionPosition = 1

func (this *QLineEdit) ActionPositionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QLineEdit_ActionPositionItemName(val int) string {
	var nilthis *QLineEdit
	return nilthis.ActionPositionItemName(val)
}

/*
This enum type describes how a line edit should display its contents.



See also setEchoMode() and echoMode().

*/
type QLineEdit__EchoMode = int

// Display characters as they are entered. This is the default.
const QLineEdit__Normal QLineEdit__EchoMode = 0

// Do not display anything. This may be appropriate for passwords where even the length of the password should be kept secret.
const QLineEdit__NoEcho QLineEdit__EchoMode = 1

// Display platform-dependent password mask characters instead of the characters actually entered.
const QLineEdit__Password QLineEdit__EchoMode = 2

// Display characters as they are entered while editing otherwise display characters as with Password.
const QLineEdit__PasswordEchoOnEdit QLineEdit__EchoMode = 3

func (this *QLineEdit) EchoModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QLineEdit_EchoModeItemName(val int) string {
	var nilthis *QLineEdit
	return nilthis.EchoModeItemName(val)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
