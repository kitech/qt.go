//  header block begin
// /usr/include/qt/QtWidgets/qabstractspinbox.h
// #include <qabstractspinbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 48
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
type QAbstractSpinBox struct {
	*QWidget
}

func (this *QAbstractSpinBox) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractSpinBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:73
// index:0
// void QAbstractSpinBox(class QWidget *)
func NewQAbstractSpinBox(parent unsafe.Pointer) *QAbstractSpinBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractSpinBoxFromPointer(cthis)
	return gothis
}
func NewQAbstractSpinBoxFromPointer(cthis unsafe.Pointer) *QAbstractSpinBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QAbstractSpinBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:165
// index:1
// void QAbstractSpinBox(class QAbstractSpinBoxPrivate &, class QWidget *)
func NewQAbstractSpinBox_1(dd unsafe.Pointer, parent unsafe.Pointer) *QAbstractSpinBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBoxC2ER23QAbstractSpinBoxPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractSpinBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:74
// index:0
// virtual
// void ~QAbstractSpinBox()
func DeleteQAbstractSpinBox(*QAbstractSpinBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:83
// index:0
// QAbstractSpinBox::ButtonSymbols buttonSymbols()
func (this *QAbstractSpinBox) ButtonSymbols() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13buttonSymbolsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:84
// index:0
// void setButtonSymbols(enum QAbstractSpinBox::ButtonSymbols)
func (this *QAbstractSpinBox) SetButtonSymbols(bs int) {
	// 0: (, bs QAbstractSpinBox::ButtonSymbols), (&bs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox16setButtonSymbolsENS_13ButtonSymbolsE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &bs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:89
// index:0
// void setCorrectionMode(enum QAbstractSpinBox::CorrectionMode)
func (this *QAbstractSpinBox) SetCorrectionMode(cm int) {
	// 0: (, cm QAbstractSpinBox::CorrectionMode), (&cm)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox17setCorrectionModeENS_14CorrectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cm)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:90
// index:0
// QAbstractSpinBox::CorrectionMode correctionMode()
func (this *QAbstractSpinBox) CorrectionMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox14correctionModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:92
// index:0
// bool hasAcceptableInput()
func (this *QAbstractSpinBox) HasAcceptableInput() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox18hasAcceptableInputEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:93
// index:0
// QString text()
func (this *QAbstractSpinBox) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox4textEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:95
// index:0
// QString specialValueText()
func (this *QAbstractSpinBox) SpecialValueText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16specialValueTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:96
// index:0
// void setSpecialValueText(const class QString &)
func (this *QAbstractSpinBox) SetSpecialValueText(txt unsafe.Pointer) {
	// 0: (, txt const QString &), (txt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setSpecialValueTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), txt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:98
// index:0
// bool wrapping()
func (this *QAbstractSpinBox) Wrapping() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8wrappingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:99
// index:0
// void setWrapping(_Bool)
func (this *QAbstractSpinBox) SetWrapping(w bool) {
	// 0: (, w bool), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setWrappingEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:101
// index:0
// void setReadOnly(_Bool)
func (this *QAbstractSpinBox) SetReadOnly(r bool) {
	// 0: (, r bool), (&r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setReadOnlyEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:102
// index:0
// bool isReadOnly()
func (this *QAbstractSpinBox) IsReadOnly() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10isReadOnlyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:104
// index:0
// void setKeyboardTracking(_Bool)
func (this *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	// 0: (, kt bool), (&kt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setKeyboardTrackingEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &kt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:105
// index:0
// bool keyboardTracking()
func (this *QAbstractSpinBox) KeyboardTracking() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16keyboardTrackingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:107
// index:0
// void setAlignment(Qt::Alignment)
func (this *QAbstractSpinBox) SetAlignment(flag int) {
	// 0: (, QFlags<Qt::AlignmentFlag> flag), (&flag)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flag)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:108
// index:0
// Qt::Alignment alignment()
func (this *QAbstractSpinBox) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox9alignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:110
// index:0
// void setFrame(_Bool)
func (this *QAbstractSpinBox) SetFrame(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox8setFrameEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:111
// index:0
// bool hasFrame()
func (this *QAbstractSpinBox) HasFrame() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8hasFrameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:113
// index:0
// void setAccelerated(_Bool)
func (this *QAbstractSpinBox) SetAccelerated(on bool) {
	// 0: (, on bool), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox14setAcceleratedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:114
// index:0
// bool isAccelerated()
func (this *QAbstractSpinBox) IsAccelerated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13isAcceleratedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:116
// index:0
// void setGroupSeparatorShown(_Bool)
func (this *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	// 0: (, shown bool), (&shown)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox22setGroupSeparatorShownEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &shown)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:117
// index:0
// bool isGroupSeparatorShown()
func (this *QAbstractSpinBox) IsGroupSeparatorShown() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox21isGroupSeparatorShownEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:119
// index:0
// virtual
// QSize sizeHint()
func (this *QAbstractSpinBox) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:120
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QAbstractSpinBox) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:121
// index:0
// void interpretText()
func (this *QAbstractSpinBox) InterpretText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13interpretTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:122
// index:0
// virtual
// bool event(class QEvent *)
func (this *QAbstractSpinBox) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:124
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QAbstractSpinBox) InputMethodQuery(arg0 int) {
	// 0: (, Qt::InputMethodQuery arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:126
// index:0
// virtual
// QValidator::State validate(class QString &, int &)
func (this *QAbstractSpinBox) Validate(input unsafe.Pointer, pos int) {
	// 0: (, input QString &, pos int &), (input, &pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8validateER7QStringRi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), input, &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:127
// index:0
// virtual
// void fixup(class QString &)
func (this *QAbstractSpinBox) Fixup(input unsafe.Pointer) {
	// 0: (, input QString &), (input)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox5fixupER7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), input)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:129
// index:0
// virtual
// void stepBy(int)
func (this *QAbstractSpinBox) StepBy(steps int) {
	// 0: (, steps int), (&steps)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepByEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:131
// index:0
// void stepUp()
func (this *QAbstractSpinBox) StepUp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepUpEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:132
// index:0
// void stepDown()
func (this *QAbstractSpinBox) StepDown() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox8stepDownEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:133
// index:0
// void selectAll()
func (this *QAbstractSpinBox) SelectAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9selectAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:134
// index:0
// virtual
// void clear()
func (this *QAbstractSpinBox) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:136
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QAbstractSpinBox) ResizeEvent(event unsafe.Pointer) {
	// 0: (, event QResizeEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:137
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractSpinBox) KeyPressEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:138
// index:0
// virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QAbstractSpinBox) KeyReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:140
// index:0
// virtual
// void wheelEvent(class QWheelEvent *)
func (this *QAbstractSpinBox) WheelEvent(event unsafe.Pointer) {
	// 0: (, event QWheelEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:142
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QAbstractSpinBox) FocusInEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:143
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractSpinBox) FocusOutEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:145
// index:0
// virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QAbstractSpinBox) ContextMenuEvent(event unsafe.Pointer) {
	// 0: (, event QContextMenuEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:147
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QAbstractSpinBox) ChangeEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:148
// index:0
// virtual
// void closeEvent(class QCloseEvent *)
func (this *QAbstractSpinBox) CloseEvent(event unsafe.Pointer) {
	// 0: (, event QCloseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:149
// index:0
// virtual
// void hideEvent(class QHideEvent *)
func (this *QAbstractSpinBox) HideEvent(event unsafe.Pointer) {
	// 0: (, event QHideEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:150
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) MousePressEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:151
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) MouseReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:152
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) MouseMoveEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:153
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QAbstractSpinBox) TimerEvent(event unsafe.Pointer) {
	// 0: (, event QTimerEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:154
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QAbstractSpinBox) PaintEvent(event unsafe.Pointer) {
	// 0: (, event QPaintEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:155
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QAbstractSpinBox) ShowEvent(event unsafe.Pointer) {
	// 0: (, event QShowEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:156
// index:0
// void initStyleOption(class QStyleOptionSpinBox *)
func (this *QAbstractSpinBox) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOptionSpinBox *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:158
// index:0
// QLineEdit * lineEdit()
func (this *QAbstractSpinBox) LineEdit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8lineEditEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:159
// index:0
// void setLineEdit(class QLineEdit *)
func (this *QAbstractSpinBox) SetLineEdit(edit unsafe.Pointer) {
	// 0: (, edit QLineEdit *), (edit)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit", ffiqt.FFI_TYPE_VOID, this.GetCthis(), edit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:161
// index:0
// virtual
// QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QAbstractSpinBox) StepEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox11stepEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:163
// index:0
// void editingFinished()
func (this *QAbstractSpinBox) EditingFinished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15editingFinishedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
