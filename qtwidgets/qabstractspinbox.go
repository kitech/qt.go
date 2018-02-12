package qtwidgets

// /usr/include/qt/QtWidgets/qabstractspinbox.h
// #include <qabstractspinbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 47
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void resizeEvent(class QResizeEvent *)
func (this *QAbstractSpinBox) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractSpinBox) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QAbstractSpinBox) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QAbstractSpinBox) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QAbstractSpinBox) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractSpinBox) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QAbstractSpinBox) InheritContextMenuEvent(f func(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QAbstractSpinBox) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void closeEvent(class QCloseEvent *)
func (this *QAbstractSpinBox) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void hideEvent(class QHideEvent *)
func (this *QAbstractSpinBox) InheritHideEvent(f func(event *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QAbstractSpinBox) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QAbstractSpinBox) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QAbstractSpinBox) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void initStyleOption(class QStyleOptionSpinBox *)
func (this *QAbstractSpinBox) InheritInitStyleOption(f func(option *QStyleOptionSpinBox /*777 QStyleOptionSpinBox **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

// QLineEdit * lineEdit()
func (this *QAbstractSpinBox) InheritLineEdit(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "lineEdit", f)
}

// void setLineEdit(class QLineEdit *)
func (this *QAbstractSpinBox) InheritSetLineEdit(f func(edit *QLineEdit /*777 QLineEdit **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setLineEdit", f)
}

// QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QAbstractSpinBox) InheritStepEnabled(f func() int) {
	qtrt.SetAllInheritCallback(this, "stepEnabled", f)
}

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAbstractSpinBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractSpinBox(QWidget *)
func NewQAbstractSpinBox(parent *QWidget /*777 QWidget **/) *QAbstractSpinBox {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractSpinBox()
func DeleteQAbstractSpinBox(this *QAbstractSpinBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSpinBox::ButtonSymbols buttonSymbols()
func (this *QAbstractSpinBox) ButtonSymbols() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13buttonSymbolsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtonSymbols(enum QAbstractSpinBox::ButtonSymbols)
func (this *QAbstractSpinBox) SetButtonSymbols(bs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox16setButtonSymbolsENS_13ButtonSymbolsE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCorrectionMode(enum QAbstractSpinBox::CorrectionMode)
func (this *QAbstractSpinBox) SetCorrectionMode(cm int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox17setCorrectionModeENS_14CorrectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cm)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSpinBox::CorrectionMode correctionMode()
func (this *QAbstractSpinBox) CorrectionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox14correctionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAcceptableInput()
func (this *QAbstractSpinBox) HasAcceptableInput() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox18hasAcceptableInputEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QAbstractSpinBox) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QString specialValueText()
func (this *QAbstractSpinBox) SpecialValueText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16specialValueTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpecialValueText(const QString &)
func (this *QAbstractSpinBox) SetSpecialValueText(txt string) {
	var tmpArg0 = qtcore.NewQString_5(txt)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setSpecialValueTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wrapping()
func (this *QAbstractSpinBox) Wrapping() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8wrappingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWrapping(_Bool)
func (this *QAbstractSpinBox) SetWrapping(w bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setWrappingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(_Bool)
func (this *QAbstractSpinBox) SetReadOnly(r bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setReadOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), r)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly()
func (this *QAbstractSpinBox) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeyboardTracking(_Bool)
func (this *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setKeyboardTrackingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), kt)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool keyboardTracking()
func (this *QAbstractSpinBox) KeyboardTracking() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16keyboardTrackingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)
func (this *QAbstractSpinBox) SetAlignment(flag int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment()
func (this *QAbstractSpinBox) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrame(_Bool)
func (this *QAbstractSpinBox) SetFrame(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox8setFrameEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFrame()
func (this *QAbstractSpinBox) HasFrame() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8hasFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccelerated(_Bool)
func (this *QAbstractSpinBox) SetAccelerated(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox14setAcceleratedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAccelerated()
func (this *QAbstractSpinBox) IsAccelerated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13isAcceleratedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGroupSeparatorShown(_Bool)
func (this *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox22setGroupSeparatorShownEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), shown)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isGroupSeparatorShown()
func (this *QAbstractSpinBox) IsGroupSeparatorShown() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox21isGroupSeparatorShownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QAbstractSpinBox) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:120
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QAbstractSpinBox) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void interpretText()
func (this *QAbstractSpinBox) InterpretText() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox13interpretTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:122
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAbstractSpinBox) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QAbstractSpinBox) InputMethodQuery(arg0 int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &)
func (this *QAbstractSpinBox) Validate(input string, pos int) int {
	var tmpArg0 = qtcore.NewQString_5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &)
func (this *QAbstractSpinBox) Fixup(input string) {
	var tmpArg0 = qtcore.NewQString_5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox5fixupER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:129
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void stepBy(int)
func (this *QAbstractSpinBox) StepBy(steps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepByEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stepUp()
func (this *QAbstractSpinBox) StepUp() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stepDown()
func (this *QAbstractSpinBox) StepDown() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox8stepDownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectAll()
func (this *QAbstractSpinBox) SelectAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox9selectAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:134
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clear()
func (this *QAbstractSpinBox) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:136
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QAbstractSpinBox) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:137
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QAbstractSpinBox) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:138
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QAbstractSpinBox) KeyReleaseEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:140
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QAbstractSpinBox) WheelEvent(event *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QAbstractSpinBox) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:143
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QAbstractSpinBox) FocusOutEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QAbstractSpinBox) ContextMenuEvent(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:147
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QAbstractSpinBox) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:148
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)
func (this *QAbstractSpinBox) CloseEvent(event *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:149
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QAbstractSpinBox) HideEvent(event *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:150
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QAbstractSpinBox) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:151
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QAbstractSpinBox) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:152
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QAbstractSpinBox) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:153
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QAbstractSpinBox) TimerEvent(event *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:154
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QAbstractSpinBox) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:155
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QAbstractSpinBox) ShowEvent(event *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:156
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionSpinBox *)
func (this *QAbstractSpinBox) InitStyleOption(option *QStyleOptionSpinBox /*777 QStyleOptionSpinBox **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:158
// index:0
// Protected Visibility=Default Availability=Available
// [8] QLineEdit * lineEdit()
func (this *QAbstractSpinBox) LineEdit() *QLineEdit /*777 QLineEdit **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8lineEditEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:159
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setLineEdit(QLineEdit *)
func (this *QAbstractSpinBox) SetLineEdit(edit *QLineEdit /*777 QLineEdit **/) {
	var convArg0 = edit.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:161
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QAbstractSpinBox) StepEnabled() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox11stepEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editingFinished()
func (this *QAbstractSpinBox) EditingFinished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox15editingFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
