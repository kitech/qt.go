//  header block begin
// /usr/include/qt/QtWidgets/qabstractspinbox.h
// #include <qabstractspinbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 47
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
func NewQAbstractSpinBoxFromPointer(cthis unsafe.Pointer) *QAbstractSpinBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QAbstractSpinBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractSpinBox) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:73
// index:0
// Public
// void QAbstractSpinBox(class QWidget *)
func NewQAbstractSpinBox(parent unsafe.Pointer) *QAbstractSpinBox {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractSpinBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:74
// index:0
// Public virtual
// void ~QAbstractSpinBox()
func DeleteQAbstractSpinBox(*QAbstractSpinBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:83
// index:0
// Public
// QAbstractSpinBox::ButtonSymbols buttonSymbols()
func (this *QAbstractSpinBox) ButtonSymbols() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13buttonSymbolsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:84
// index:0
// Public
// void setButtonSymbols(enum QAbstractSpinBox::ButtonSymbols)
func (this *QAbstractSpinBox) SetButtonSymbols(bs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox16setButtonSymbolsENS_13ButtonSymbolsE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &bs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:89
// index:0
// Public
// void setCorrectionMode(enum QAbstractSpinBox::CorrectionMode)
func (this *QAbstractSpinBox) SetCorrectionMode(cm int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox17setCorrectionModeENS_14CorrectionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cm)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:90
// index:0
// Public
// QAbstractSpinBox::CorrectionMode correctionMode()
func (this *QAbstractSpinBox) CorrectionMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox14correctionModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:92
// index:0
// Public
// bool hasAcceptableInput()
func (this *QAbstractSpinBox) HasAcceptableInput() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox18hasAcceptableInputEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:93
// index:0
// Public
// QString text()
func (this *QAbstractSpinBox) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:95
// index:0
// Public
// QString specialValueText()
func (this *QAbstractSpinBox) SpecialValueText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16specialValueTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:96
// index:0
// Public
// void setSpecialValueText(const class QString &)
func (this *QAbstractSpinBox) SetSpecialValueText(txt *qtcore.QString) {
	var convArg0 = txt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setSpecialValueTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:98
// index:0
// Public
// bool wrapping()
func (this *QAbstractSpinBox) Wrapping() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8wrappingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:99
// index:0
// Public
// void setWrapping(_Bool)
func (this *QAbstractSpinBox) SetWrapping(w bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setWrappingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:101
// index:0
// Public
// void setReadOnly(_Bool)
func (this *QAbstractSpinBox) SetReadOnly(r bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setReadOnlyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:102
// index:0
// Public
// bool isReadOnly()
func (this *QAbstractSpinBox) IsReadOnly() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10isReadOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:104
// index:0
// Public
// void setKeyboardTracking(_Bool)
func (this *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setKeyboardTrackingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &kt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:105
// index:0
// Public
// bool keyboardTracking()
func (this *QAbstractSpinBox) KeyboardTracking() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16keyboardTrackingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:108
// index:0
// Public
// Qt::Alignment alignment()
func (this *QAbstractSpinBox) Alignment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:110
// index:0
// Public
// void setFrame(_Bool)
func (this *QAbstractSpinBox) SetFrame(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox8setFrameEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:111
// index:0
// Public
// bool hasFrame()
func (this *QAbstractSpinBox) HasFrame() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8hasFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:113
// index:0
// Public
// void setAccelerated(_Bool)
func (this *QAbstractSpinBox) SetAccelerated(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox14setAcceleratedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:114
// index:0
// Public
// bool isAccelerated()
func (this *QAbstractSpinBox) IsAccelerated() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13isAcceleratedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:116
// index:0
// Public
// void setGroupSeparatorShown(_Bool)
func (this *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox22setGroupSeparatorShownEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &shown)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:117
// index:0
// Public
// bool isGroupSeparatorShown()
func (this *QAbstractSpinBox) IsGroupSeparatorShown() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox21isGroupSeparatorShownEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:119
// index:0
// Public virtual
// QSize sizeHint()
func (this *QAbstractSpinBox) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:120
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QAbstractSpinBox) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:121
// index:0
// Public
// void interpretText()
func (this *QAbstractSpinBox) InterpretText() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13interpretTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:122
// index:0
// Public virtual
// bool event(class QEvent *)
func (this *QAbstractSpinBox) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:124
// index:0
// Public virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QAbstractSpinBox) InputMethodQuery(arg0 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:126
// index:0
// Public virtual
// QValidator::State validate(class QString &, int &)
func (this *QAbstractSpinBox) Validate(input *qtcore.QString, pos int) interface{} {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:127
// index:0
// Public virtual
// void fixup(class QString &)
func (this *QAbstractSpinBox) Fixup(input *qtcore.QString) {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox5fixupER7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:129
// index:0
// Public virtual
// void stepBy(int)
func (this *QAbstractSpinBox) StepBy(steps int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepByEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:131
// index:0
// Public
// void stepUp()
func (this *QAbstractSpinBox) StepUp() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepUpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:132
// index:0
// Public
// void stepDown()
func (this *QAbstractSpinBox) StepDown() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox8stepDownEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:133
// index:0
// Public
// void selectAll()
func (this *QAbstractSpinBox) SelectAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9selectAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:134
// index:0
// Public virtual
// void clear()
func (this *QAbstractSpinBox) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:136
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QAbstractSpinBox) ResizeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:137
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractSpinBox) KeyPressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:138
// index:0
// Protected virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QAbstractSpinBox) KeyReleaseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:140
// index:0
// Protected virtual
// void wheelEvent(class QWheelEvent *)
func (this *QAbstractSpinBox) WheelEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:142
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QAbstractSpinBox) FocusInEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:143
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractSpinBox) FocusOutEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:145
// index:0
// Protected virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QAbstractSpinBox) ContextMenuEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:147
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QAbstractSpinBox) ChangeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:148
// index:0
// Protected virtual
// void closeEvent(class QCloseEvent *)
func (this *QAbstractSpinBox) CloseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:149
// index:0
// Protected virtual
// void hideEvent(class QHideEvent *)
func (this *QAbstractSpinBox) HideEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:150
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) MousePressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:151
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) MouseReleaseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:152
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) MouseMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:153
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QAbstractSpinBox) TimerEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:154
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QAbstractSpinBox) PaintEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:155
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QAbstractSpinBox) ShowEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:156
// index:0
// Protected
// void initStyleOption(class QStyleOptionSpinBox *)
func (this *QAbstractSpinBox) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:158
// index:0
// Protected
// QLineEdit * lineEdit()
func (this *QAbstractSpinBox) LineEdit() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8lineEditEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:159
// index:0
// Protected
// void setLineEdit(class QLineEdit *)
func (this *QAbstractSpinBox) SetLineEdit(edit unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), edit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:161
// index:0
// Protected virtual
// QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QAbstractSpinBox) StepEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox11stepEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:163
// index:0
// Public
// void editingFinished()
func (this *QAbstractSpinBox) EditingFinished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15editingFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
