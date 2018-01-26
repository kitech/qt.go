package qtwidgets

// /usr/include/qt/QtWidgets/qabstractspinbox.h
// #include <qabstractspinbox.h>
// #include <QtWidgets>

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
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QAbstractSpinBox) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQAbstractSpinBoxFromPointer(cthis unsafe.Pointer) *QAbstractSpinBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QAbstractSpinBox{bcthis0}
}
func (*QAbstractSpinBox) NewFromPointer(cthis unsafe.Pointer) *QAbstractSpinBox {
	return NewQAbstractSpinBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractSpinBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:73
// index:0
// Public
// void QAbstractSpinBox(class QWidget *)
func NewQAbstractSpinBox(parent *QWidget /*777 QWidget **/) *QAbstractSpinBox {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
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
func (this *QAbstractSpinBox) ButtonSymbols() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13buttonSymbolsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:84
// index:0
// Public
// void setButtonSymbols(enum QAbstractSpinBox::ButtonSymbols)
func (this *QAbstractSpinBox) SetButtonSymbols(bs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox16setButtonSymbolsENS_13ButtonSymbolsE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), bs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:89
// index:0
// Public
// void setCorrectionMode(enum QAbstractSpinBox::CorrectionMode)
func (this *QAbstractSpinBox) SetCorrectionMode(cm int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox17setCorrectionModeENS_14CorrectionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cm)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:90
// index:0
// Public
// QAbstractSpinBox::CorrectionMode correctionMode()
func (this *QAbstractSpinBox) CorrectionMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox14correctionModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:92
// index:0
// Public
// bool hasAcceptableInput()
func (this *QAbstractSpinBox) HasAcceptableInput() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox18hasAcceptableInputEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:93
// index:0
// Public
// QString text()
func (this *QAbstractSpinBox) Text() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox4textEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:95
// index:0
// Public
// QString specialValueText()
func (this *QAbstractSpinBox) SpecialValueText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16specialValueTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
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
func (this *QAbstractSpinBox) Wrapping() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8wrappingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:99
// index:0
// Public
// void setWrapping(_Bool)
func (this *QAbstractSpinBox) SetWrapping(w bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setWrappingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:101
// index:0
// Public
// void setReadOnly(_Bool)
func (this *QAbstractSpinBox) SetReadOnly(r bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setReadOnlyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:102
// index:0
// Public
// bool isReadOnly()
func (this *QAbstractSpinBox) IsReadOnly() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10isReadOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:104
// index:0
// Public
// void setKeyboardTracking(_Bool)
func (this *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setKeyboardTrackingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), kt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:105
// index:0
// Public
// bool keyboardTracking()
func (this *QAbstractSpinBox) KeyboardTracking() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16keyboardTrackingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:108
// index:0
// Public
// Qt::Alignment alignment()
func (this *QAbstractSpinBox) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:110
// index:0
// Public
// void setFrame(_Bool)
func (this *QAbstractSpinBox) SetFrame(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox8setFrameEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:111
// index:0
// Public
// bool hasFrame()
func (this *QAbstractSpinBox) HasFrame() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8hasFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:113
// index:0
// Public
// void setAccelerated(_Bool)
func (this *QAbstractSpinBox) SetAccelerated(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox14setAcceleratedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:114
// index:0
// Public
// bool isAccelerated()
func (this *QAbstractSpinBox) IsAccelerated() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13isAcceleratedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:116
// index:0
// Public
// void setGroupSeparatorShown(_Bool)
func (this *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox22setGroupSeparatorShownEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), shown)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:117
// index:0
// Public
// bool isGroupSeparatorShown()
func (this *QAbstractSpinBox) IsGroupSeparatorShown() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox21isGroupSeparatorShownEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:119
// index:0
// Public virtual
// QSize sizeHint()
func (this *QAbstractSpinBox) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:120
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QAbstractSpinBox) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
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
func (this *QAbstractSpinBox) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:124
// index:0
// Public virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QAbstractSpinBox) InputMethodQuery(arg0 int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:126
// index:0
// Public virtual
// QValidator::State validate(class QString &, int &)
func (this *QAbstractSpinBox) Validate(input *qtcore.QString, pos int) int {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepByEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), steps)
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
func (this *QAbstractSpinBox) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:137
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractSpinBox) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:138
// index:0
// Protected virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QAbstractSpinBox) KeyReleaseEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:140
// index:0
// Protected virtual
// void wheelEvent(class QWheelEvent *)
func (this *QAbstractSpinBox) WheelEvent(event *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:142
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QAbstractSpinBox) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:143
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractSpinBox) FocusOutEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:145
// index:0
// Protected virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QAbstractSpinBox) ContextMenuEvent(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:147
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QAbstractSpinBox) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:148
// index:0
// Protected virtual
// void closeEvent(class QCloseEvent *)
func (this *QAbstractSpinBox) CloseEvent(event *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:149
// index:0
// Protected virtual
// void hideEvent(class QHideEvent *)
func (this *QAbstractSpinBox) HideEvent(event *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:150
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:151
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:152
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:153
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QAbstractSpinBox) TimerEvent(event *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:154
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QAbstractSpinBox) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:155
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QAbstractSpinBox) ShowEvent(event *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:156
// index:0
// Protected
// void initStyleOption(class QStyleOptionSpinBox *)
func (this *QAbstractSpinBox) InitStyleOption(option *QStyleOptionSpinBox /*777 QStyleOptionSpinBox **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:158
// index:0
// Protected
// QLineEdit * lineEdit()
func (this *QAbstractSpinBox) LineEdit() *QLineEdit /*777 QLineEdit **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8lineEditEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:159
// index:0
// Protected
// void setLineEdit(class QLineEdit *)
func (this *QAbstractSpinBox) SetLineEdit(edit *QLineEdit /*777 QLineEdit **/) {
	var convArg0 = edit.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:161
// index:0
// Protected virtual
// QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QAbstractSpinBox) StepEnabled() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox11stepEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:163
// index:0
// Public
// void editingFinished()
func (this *QAbstractSpinBox) EditingFinished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15editingFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QAbstractSpinBox__StepEnabledFlag = int

const QAbstractSpinBox__StepNone QAbstractSpinBox__StepEnabledFlag = 0
const QAbstractSpinBox__StepUpEnabled QAbstractSpinBox__StepEnabledFlag = 1
const QAbstractSpinBox__StepDownEnabled QAbstractSpinBox__StepEnabledFlag = 2

type QAbstractSpinBox__ButtonSymbols = int

const QAbstractSpinBox__UpDownArrows QAbstractSpinBox__ButtonSymbols = 0
const QAbstractSpinBox__PlusMinus QAbstractSpinBox__ButtonSymbols = 1
const QAbstractSpinBox__NoButtons QAbstractSpinBox__ButtonSymbols = 2

type QAbstractSpinBox__CorrectionMode = int

const QAbstractSpinBox__CorrectToPreviousValue QAbstractSpinBox__CorrectionMode = 0
const QAbstractSpinBox__CorrectToNearestValue QAbstractSpinBox__CorrectionMode = 1

//  body block end
