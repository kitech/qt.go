//  header block begin
// /usr/include/qt/QtWidgets/qlineedit.h
// #include <qlineedit.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 78
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
type QLineEdit struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qlineedit.h:65
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QLineEdit) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:93
// index:0
// void QLineEdit(class QWidget *)
func NewQLineEdit(parent unsafe.Pointer) *QLineEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QLineEdit{cthis}
}

// /usr/include/qt/QtWidgets/qlineedit.h:94
// index:1
// void QLineEdit(const class QString &, class QWidget *)
func NewQLineEdit_1(arg0 unsafe.Pointer, parent unsafe.Pointer) *QLineEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEditC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, arg0, parent)
	gopp.ErrPrint(err, rv)
	return &QLineEdit{cthis}
}

// /usr/include/qt/QtWidgets/qlineedit.h:95
// index:0
// virtual
// void ~QLineEdit()
func DeleteQLineEdit(*QLineEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:97
// index:0
// QString text()
func (this *QLineEdit) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:99
// index:0
// QString displayText()
func (this *QLineEdit) DisplayText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit11displayTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:101
// index:0
// QString placeholderText()
func (this *QLineEdit) PlaceholderText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15placeholderTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:102
// index:0
// void setPlaceholderText(const class QString &)
func (this *QLineEdit) SetPlaceholderText(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit18setPlaceholderTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:104
// index:0
// int maxLength()
func (this *QLineEdit) MaxLength() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit9maxLengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:105
// index:0
// void setMaxLength(int)
func (this *QLineEdit) SetMaxLength(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12setMaxLengthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:107
// index:0
// void setFrame(_Bool)
func (this *QLineEdit) SetFrame(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit8setFrameEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:108
// index:0
// bool hasFrame()
func (this *QLineEdit) HasFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit8hasFrameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:110
// index:0
// void setClearButtonEnabled(_Bool)
func (this *QLineEdit) SetClearButtonEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit21setClearButtonEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:111
// index:0
// bool isClearButtonEnabled()
func (this *QLineEdit) IsClearButtonEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit20isClearButtonEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:115
// index:0
// QLineEdit::EchoMode echoMode()
func (this *QLineEdit) EchoMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit8echoModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:116
// index:0
// void setEchoMode(enum QLineEdit::EchoMode)
func (this *QLineEdit) SetEchoMode(arg0 int) {
	// 0: (, QLineEdit::EchoMode arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit11setEchoModeENS_8EchoModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:118
// index:0
// bool isReadOnly()
func (this *QLineEdit) IsReadOnly() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit10isReadOnlyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:119
// index:0
// void setReadOnly(_Bool)
func (this *QLineEdit) SetReadOnly(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit11setReadOnlyEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:122
// index:0
// void setValidator(const class QValidator *)
func (this *QLineEdit) SetValidator(arg0 unsafe.Pointer) {
	// 0: (, const QValidator * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12setValidatorEPK10QValidator", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:123
// index:0
// const QValidator * validator()
func (this *QLineEdit) Validator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit9validatorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:127
// index:0
// void setCompleter(class QCompleter *)
func (this *QLineEdit) SetCompleter(completer unsafe.Pointer) {
	// 0: (, QCompleter * completer), (completer)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12setCompleterEP10QCompleter", ffiqt.FFI_TYPE_VOID, this.cthis, completer)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:128
// index:0
// QCompleter * completer()
func (this *QLineEdit) Completer() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit9completerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:131
// index:0
// virtual
// QSize sizeHint()
func (this *QLineEdit) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:132
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QLineEdit) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:134
// index:0
// int cursorPosition()
func (this *QLineEdit) CursorPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit14cursorPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:135
// index:0
// void setCursorPosition(int)
func (this *QLineEdit) SetCursorPosition(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit17setCursorPositionEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:136
// index:0
// int cursorPositionAt(const class QPoint &)
func (this *QLineEdit) CursorPositionAt(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit16cursorPositionAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:139
// index:0
// Qt::Alignment alignment()
func (this *QLineEdit) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit9alignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:141
// index:0
// void cursorForward(_Bool, int)
func (this *QLineEdit) CursorForward(mark bool, steps int) {
	// 0: (, bool mark, int steps), (&mark, &steps)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit13cursorForwardEbi", ffiqt.FFI_TYPE_VOID, this.cthis, &mark, &steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:142
// index:0
// void cursorBackward(_Bool, int)
func (this *QLineEdit) CursorBackward(mark bool, steps int) {
	// 0: (, bool mark, int steps), (&mark, &steps)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14cursorBackwardEbi", ffiqt.FFI_TYPE_VOID, this.cthis, &mark, &steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:143
// index:0
// void cursorWordForward(_Bool)
func (this *QLineEdit) CursorWordForward(mark bool) {
	// 0: (, bool mark), (&mark)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit17cursorWordForwardEb", ffiqt.FFI_TYPE_VOID, this.cthis, &mark)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:144
// index:0
// void cursorWordBackward(_Bool)
func (this *QLineEdit) CursorWordBackward(mark bool) {
	// 0: (, bool mark), (&mark)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit18cursorWordBackwardEb", ffiqt.FFI_TYPE_VOID, this.cthis, &mark)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:145
// index:0
// void backspace()
func (this *QLineEdit) Backspace() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit9backspaceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:146
// index:0
// void del()
func (this *QLineEdit) Del() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit3delEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:147
// index:0
// void home(_Bool)
func (this *QLineEdit) Home(mark bool) {
	// 0: (, bool mark), (&mark)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit4homeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &mark)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:148
// index:0
// void end(_Bool)
func (this *QLineEdit) End(mark bool) {
	// 0: (, bool mark), (&mark)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit3endEb", ffiqt.FFI_TYPE_VOID, this.cthis, &mark)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:150
// index:0
// bool isModified()
func (this *QLineEdit) IsModified() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit10isModifiedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:151
// index:0
// void setModified(_Bool)
func (this *QLineEdit) SetModified(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit11setModifiedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:153
// index:0
// void setSelection(int, int)
func (this *QLineEdit) SetSelection(arg0 int, arg1 int) {
	// 0: (, int arg0, int arg1), (&arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12setSelectionEii", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:154
// index:0
// bool hasSelectedText()
func (this *QLineEdit) HasSelectedText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15hasSelectedTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:155
// index:0
// QString selectedText()
func (this *QLineEdit) SelectedText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit12selectedTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:156
// index:0
// int selectionStart()
func (this *QLineEdit) SelectionStart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit14selectionStartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:157
// index:0
// int selectionEnd()
func (this *QLineEdit) SelectionEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit12selectionEndEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:158
// index:0
// int selectionLength()
func (this *QLineEdit) SelectionLength() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15selectionLengthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:160
// index:0
// bool isUndoAvailable()
func (this *QLineEdit) IsUndoAvailable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15isUndoAvailableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:161
// index:0
// bool isRedoAvailable()
func (this *QLineEdit) IsRedoAvailable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15isRedoAvailableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:163
// index:0
// void setDragEnabled(_Bool)
func (this *QLineEdit) SetDragEnabled(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14setDragEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:164
// index:0
// bool dragEnabled()
func (this *QLineEdit) DragEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit11dragEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:166
// index:0
// void setCursorMoveStyle(Qt::CursorMoveStyle)
func (this *QLineEdit) SetCursorMoveStyle(style int) {
	// 0: (, Qt::CursorMoveStyle style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit18setCursorMoveStyleEN2Qt15CursorMoveStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:167
// index:0
// Qt::CursorMoveStyle cursorMoveStyle()
func (this *QLineEdit) CursorMoveStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15cursorMoveStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:169
// index:0
// QString inputMask()
func (this *QLineEdit) InputMask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit9inputMaskEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:170
// index:0
// void setInputMask(const class QString &)
func (this *QLineEdit) SetInputMask(inputMask unsafe.Pointer) {
	// 0: (, const QString & inputMask), (inputMask)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12setInputMaskERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, inputMask)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:171
// index:0
// bool hasAcceptableInput()
func (this *QLineEdit) HasAcceptableInput() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit18hasAcceptableInputEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:173
// index:0
// void setTextMargins(int, int, int, int)
func (this *QLineEdit) SetTextMargins(left int, top int, right int, bottom int) {
	// 0: (, int left, int top, int right, int bottom), (&left, &top, &right, &bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14setTextMarginsEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:174
// index:1
// void setTextMargins(const class QMargins &)
func (this *QLineEdit) SetTextMargins_1(margins unsafe.Pointer) {
	// 1: (, const QMargins & margins), (margins)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14setTextMarginsERK8QMargins", ffiqt.FFI_TYPE_VOID, this.cthis, margins)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:175
// index:0
// void getTextMargins(int *, int *, int *, int *)
func (this *QLineEdit) GetTextMargins(left unsafe.Pointer, top unsafe.Pointer, right unsafe.Pointer, bottom unsafe.Pointer) {
	// 0: (, int * left, int * top, int * right, int * bottom), (left, top, right, bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit14getTextMarginsEPiS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:176
// index:0
// QMargins textMargins()
func (this *QLineEdit) TextMargins() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit11textMarginsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:180
// index:0
// void addAction(class QAction *, enum QLineEdit::ActionPosition)
func (this *QLineEdit) AddAction(action unsafe.Pointer, position int) {
	// 0: (, QAction * action, QLineEdit::ActionPosition position), (action, &position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit9addActionEP7QActionNS_14ActionPositionE", ffiqt.FFI_TYPE_VOID, this.cthis, action, &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:181
// index:1
// QAction * addAction(const class QIcon &, enum QLineEdit::ActionPosition)
func (this *QLineEdit) AddAction_1(icon unsafe.Pointer, position int) {
	// 1: (, const QIcon & icon, QLineEdit::ActionPosition position), (icon, &position)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit9addActionERK5QIconNS_14ActionPositionE", ffiqt.FFI_TYPE_VOID, this.cthis, icon, &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:185
// index:0
// void setText(const class QString &)
func (this *QLineEdit) SetText(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:186
// index:0
// void clear()
func (this *QLineEdit) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:187
// index:0
// void selectAll()
func (this *QLineEdit) SelectAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit9selectAllEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:188
// index:0
// void undo()
func (this *QLineEdit) Undo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit4undoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:189
// index:0
// void redo()
func (this *QLineEdit) Redo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit4redoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:191
// index:0
// void cut()
func (this *QLineEdit) Cut() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit3cutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:192
// index:0
// void copy()
func (this *QLineEdit) Copy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit4copyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:193
// index:0
// void paste()
func (this *QLineEdit) Paste() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit5pasteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:197
// index:0
// void deselect()
func (this *QLineEdit) Deselect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit8deselectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:200
// index:0
// QMenu * createStandardContextMenu()
func (this *QLineEdit) CreateStandardContextMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit25createStandardContextMenuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:204
// index:0
// void textChanged(const class QString &)
func (this *QLineEdit) TextChanged(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit11textChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:205
// index:0
// void textEdited(const class QString &)
func (this *QLineEdit) TextEdited(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit10textEditedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:206
// index:0
// void cursorPositionChanged(int, int)
func (this *QLineEdit) CursorPositionChanged(arg0 int, arg1 int) {
	// 0: (, int arg0, int arg1), (&arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit21cursorPositionChangedEii", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:207
// index:0
// void returnPressed()
func (this *QLineEdit) ReturnPressed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit13returnPressedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:208
// index:0
// void editingFinished()
func (this *QLineEdit) EditingFinished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit15editingFinishedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:209
// index:0
// void selectionChanged()
func (this *QLineEdit) SelectionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit16selectionChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:234
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QLineEdit) InputMethodQuery(arg0 int) {
	// 0: (, Qt::InputMethodQuery arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:235
// index:1
// QVariant inputMethodQuery(Qt::InputMethodQuery, class QVariant)
func (this *QLineEdit) InputMethodQuery_1(property int, argument unsafe.Pointer) {
	// 1: (, Qt::InputMethodQuery property, QVariant argument), (&property, argument)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit16inputMethodQueryEN2Qt16InputMethodQueryE8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &property, argument)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:236
// index:0
// virtual
// bool event(class QEvent *)
func (this *QLineEdit) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
