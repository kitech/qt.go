// +build !minimal

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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void resizeEvent(QResizeEvent *)
func (this *QAbstractSpinBox) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QAbstractSpinBox) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QAbstractSpinBox) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QAbstractSpinBox) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QAbstractSpinBox) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QAbstractSpinBox) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QAbstractSpinBox) InheritContextMenuEvent(f func(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void changeEvent(QEvent *)
func (this *QAbstractSpinBox) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void closeEvent(QCloseEvent *)
func (this *QAbstractSpinBox) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QAbstractSpinBox) InheritHideEvent(f func(event *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QAbstractSpinBox) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QAbstractSpinBox) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QAbstractSpinBox) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QAbstractSpinBox) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QAbstractSpinBox) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QAbstractSpinBox) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void initStyleOption(QStyleOptionSpinBox *)
func (this *QAbstractSpinBox) InheritInitStyleOption(f func(option *QStyleOptionSpinBox /*777 QStyleOptionSpinBox **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

// QLineEdit * lineEdit()
func (this *QAbstractSpinBox) InheritLineEdit(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "lineEdit", f)
}

// void setLineEdit(QLineEdit *)
func (this *QAbstractSpinBox) InheritSetLineEdit(f func(edit *QLineEdit /*777 QLineEdit **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setLineEdit", f)
}

// QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QAbstractSpinBox) InheritStepEnabled(f func() int) {
	qtrt.SetAllInheritCallback(this, "stepEnabled", f)
}

/*

 */
type QAbstractSpinBox struct {
	*QWidget
}
type QAbstractSpinBox_ITF interface {
	QWidget_ITF
	QAbstractSpinBox_PTR() *QAbstractSpinBox
}

func (ptr *QAbstractSpinBox) QAbstractSpinBox_PTR() *QAbstractSpinBox { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractSpinBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractSpinBox(QWidget *)

/*
Constructs an abstract spinbox with the given parent with default wrapping, and alignment properties.
*/
func (*QAbstractSpinBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QAbstractSpinBox {
	return NewQAbstractSpinBox(parent)
}
func NewQAbstractSpinBox(parent QWidget_ITF /*777 QWidget **/) *QAbstractSpinBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractSpinBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractSpinBox(QWidget *)

/*
Constructs an abstract spinbox with the given parent with default wrapping, and alignment properties.
*/
func (*QAbstractSpinBox) NewForInheritp() *QAbstractSpinBox {
	return NewQAbstractSpinBoxp()
}
func NewQAbstractSpinBoxp() *QAbstractSpinBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractSpinBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractSpinBox()

/*

 */
func DeleteQAbstractSpinBox(this *QAbstractSpinBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSpinBox::ButtonSymbols buttonSymbols() const

/*

 */
func (this *QAbstractSpinBox) ButtonSymbols() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13buttonSymbolsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtonSymbols(QAbstractSpinBox::ButtonSymbols)

/*

 */
func (this *QAbstractSpinBox) SetButtonSymbols(bs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox16setButtonSymbolsENS_13ButtonSymbolsE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCorrectionMode(QAbstractSpinBox::CorrectionMode)

/*

 */
func (this *QAbstractSpinBox) SetCorrectionMode(cm int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox17setCorrectionModeENS_14CorrectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cm)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSpinBox::CorrectionMode correctionMode() const

/*

 */
func (this *QAbstractSpinBox) CorrectionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox14correctionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAcceptableInput() const

/*

 */
func (this *QAbstractSpinBox) HasAcceptableInput() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox18hasAcceptableInputEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*

 */
func (this *QAbstractSpinBox) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QString specialValueText() const

/*

 */
func (this *QAbstractSpinBox) SpecialValueText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16specialValueTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpecialValueText(const QString &)

/*

 */
func (this *QAbstractSpinBox) SetSpecialValueText(txt string) {
	var tmpArg0 = qtcore.NewQString5(txt)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setSpecialValueTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wrapping() const

/*

 */
func (this *QAbstractSpinBox) Wrapping() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8wrappingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWrapping(bool)

/*

 */
func (this *QAbstractSpinBox) SetWrapping(w bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setWrappingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(bool)

/*

 */
func (this *QAbstractSpinBox) SetReadOnly(r bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setReadOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), r)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly() const

/*

 */
func (this *QAbstractSpinBox) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeyboardTracking(bool)

/*

 */
func (this *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setKeyboardTrackingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), kt)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool keyboardTracking() const

/*

 */
func (this *QAbstractSpinBox) KeyboardTracking() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16keyboardTrackingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)

/*

 */
func (this *QAbstractSpinBox) SetAlignment(flag int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const

/*

 */
func (this *QAbstractSpinBox) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrame(bool)

/*

 */
func (this *QAbstractSpinBox) SetFrame(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox8setFrameEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFrame() const

/*

 */
func (this *QAbstractSpinBox) HasFrame() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8hasFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccelerated(bool)

/*

 */
func (this *QAbstractSpinBox) SetAccelerated(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox14setAcceleratedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAccelerated() const

/*

 */
func (this *QAbstractSpinBox) IsAccelerated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13isAcceleratedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGroupSeparatorShown(bool)

/*

 */
func (this *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox22setGroupSeparatorShownEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), shown)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isGroupSeparatorShown() const

/*

 */
func (this *QAbstractSpinBox) IsGroupSeparatorShown() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox21isGroupSeparatorShownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
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
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
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

/*
This function interprets the text of the spin box. If the value has changed since last interpretation it will emit signals.
*/
func (this *QAbstractSpinBox) InterpretText() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox13interpretTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:122
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QWidget::event().
*/
func (this *QAbstractSpinBox) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
Reimplemented from QWidget::inputMethodQuery().
*/
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
// [4] QValidator::State validate(QString &, int &) const

/*
This virtual function is called by the QAbstractSpinBox to determine whether input is valid. The pos parameter indicates the position in the string. Reimplemented in the various subclasses.
*/
func (this *QAbstractSpinBox) Validate(input string, pos int) int {
	var tmpArg0 = qtcore.NewQString5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &) const

/*
This virtual function is called by the QAbstractSpinBox if the input is not validated to QValidator::Acceptable when Return is pressed or interpretText() is called. It will try to change the text so it is valid. Reimplemented in the various subclasses.
*/
func (this *QAbstractSpinBox) Fixup(input string) {
	var tmpArg0 = qtcore.NewQString5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox5fixupER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:129
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void stepBy(int)

/*
Virtual function that is called whenever the user triggers a step. The steps parameter indicates how many steps were taken. For example, pressing Qt::Key_Down will trigger a call to stepBy(-1), whereas pressing Qt::Key_PageUp will trigger a call to stepBy(10).

If you subclass QAbstractSpinBox you must reimplement this function. Note that this function is called even if the resulting value will be outside the bounds of minimum and maximum. It's this function's job to handle these situations.

See also stepUp(), stepDown(), and keyPressEvent().
*/
func (this *QAbstractSpinBox) StepBy(steps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepByEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stepUp()

/*
Steps up by one linestep Calling this slot is analogous to calling stepBy(1);

See also stepBy() and stepDown().
*/
func (this *QAbstractSpinBox) StepUp() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stepDown()

/*
Steps down by one linestep Calling this slot is analogous to calling stepBy(-1);

See also stepBy() and stepUp().
*/
func (this *QAbstractSpinBox) StepDown() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox8stepDownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectAll()

/*
Selects all the text in the spinbox except the prefix and suffix.
*/
func (this *QAbstractSpinBox) SelectAll() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox9selectAllEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:141
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the lineedit of all text but prefix and suffix.
*/
func (this *QAbstractSpinBox) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:143
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QAbstractSpinBox) ResizeEvent(event qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QResizeEvent_PTR() != nil {
		convArg0 = event.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:144
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().

This function handles keyboard input.

The following keys are handled specifically:


 Enter/ReturnThis will reinterpret the text and emit a signal even if the value has not changed since last time a signal was emitted.
UpThis will invoke stepBy(1)
DownThis will invoke stepBy(-1)
Page upThis will invoke stepBy(10)
Page downThis will invoke stepBy(-10)


See also stepBy().
*/
func (this *QAbstractSpinBox) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyReleaseEvent().
*/
func (this *QAbstractSpinBox) KeyReleaseEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:149
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().
*/
func (this *QAbstractSpinBox) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:150
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().
*/
func (this *QAbstractSpinBox) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:154
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QAbstractSpinBox) ChangeEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:155
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)

/*
Reimplemented from QWidget::closeEvent().
*/
func (this *QAbstractSpinBox) CloseEvent(event qtgui.QCloseEvent_ITF /*777 QCloseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QCloseEvent_PTR() != nil {
		convArg0 = event.QCloseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:156
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
Reimplemented from QWidget::hideEvent().
*/
func (this *QAbstractSpinBox) HideEvent(event qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QHideEvent_PTR() != nil {
		convArg0 = event.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:157
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QAbstractSpinBox) MousePressEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:158
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QAbstractSpinBox) MouseReleaseEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:159
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QAbstractSpinBox) MouseMoveEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:160
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QAbstractSpinBox) TimerEvent(event qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QTimerEvent_PTR() != nil {
		convArg0 = event.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:161
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QAbstractSpinBox) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:162
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QWidget::showEvent().
*/
func (this *QAbstractSpinBox) ShowEvent(event qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QShowEvent_PTR() != nil {
		convArg0 = event.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:163
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionSpinBox *) const

/*
Initialize option with the values from this QSpinBox. This method is useful for subclasses when they need a QStyleOptionSpinBox, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QAbstractSpinBox) InitStyleOption(option QStyleOptionSpinBox_ITF /*777 QStyleOptionSpinBox **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionSpinBox_PTR() != nil {
		convArg0 = option.QStyleOptionSpinBox_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:165
// index:0
// Protected Visibility=Default Availability=Available
// [8] QLineEdit * lineEdit() const

/*
This function returns a pointer to the line edit of the spin box.

See also setLineEdit().
*/
func (this *QAbstractSpinBox) LineEdit() *QLineEdit /*777 QLineEdit **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8lineEditEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:166
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setLineEdit(QLineEdit *)

/*
Sets the line edit of the spinbox to be lineEdit instead of the current line edit widget. lineEdit can not be 0.

QAbstractSpinBox takes ownership of the new lineEdit

If QLineEdit::validator() for the lineEdit returns 0, the internal validator of the spinbox will be set on the line edit.

See also lineEdit().
*/
func (this *QAbstractSpinBox) SetLineEdit(edit QLineEdit_ITF /*777 QLineEdit **/) {
	var convArg0 unsafe.Pointer
	if edit != nil && edit.QLineEdit_PTR() != nil {
		convArg0 = edit.QLineEdit_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:168
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QAbstractSpinBox::StepEnabled stepEnabled() const

/*
Virtual function that determines whether stepping up and down is legal at any given time.

The up arrow will be painted as disabled unless (stepEnabled() & StepUpEnabled) != 0.

The default implementation will return (StepUpEnabled| StepDownEnabled) if wrapping is turned on. Else it will return StepDownEnabled if value is > minimum() or'ed with StepUpEnabled if value < maximum().

If you subclass QAbstractSpinBox you will need to reimplement this function.

See also QSpinBox::minimum(), QSpinBox::maximum(), and wrapping().
*/
func (this *QAbstractSpinBox) StepEnabled() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAbstractSpinBox11stepEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editingFinished()

/*
This signal is emitted editing is finished. This happens when the spinbox loses focus and when enter is pressed.
*/
func (this *QAbstractSpinBox) EditingFinished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAbstractSpinBox15editingFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QAbstractSpinBox__StepEnabledFlag = int

//
const QAbstractSpinBox__StepNone QAbstractSpinBox__StepEnabledFlag = 0

//
const QAbstractSpinBox__StepUpEnabled QAbstractSpinBox__StepEnabledFlag = 1

//
const QAbstractSpinBox__StepDownEnabled QAbstractSpinBox__StepEnabledFlag = 2

func (this *QAbstractSpinBox) StepEnabledFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSpinBox_StepEnabledFlagItemName(val int) string {
	var nilthis *QAbstractSpinBox
	return nilthis.StepEnabledFlagItemName(val)
}

/*
This enum type describes the symbols that can be displayed on the buttons in a spin box.





See also QAbstractSpinBox::buttonSymbols.

*/
type QAbstractSpinBox__ButtonSymbols = int

// Little arrows in the classic style.
const QAbstractSpinBox__UpDownArrows QAbstractSpinBox__ButtonSymbols = 0

// + and - symbols.
const QAbstractSpinBox__PlusMinus QAbstractSpinBox__ButtonSymbols = 1

// Don't display buttons.
const QAbstractSpinBox__NoButtons QAbstractSpinBox__ButtonSymbols = 2

func (this *QAbstractSpinBox) ButtonSymbolsItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSpinBox_ButtonSymbolsItemName(val int) string {
	var nilthis *QAbstractSpinBox
	return nilthis.ButtonSymbolsItemName(val)
}

/*
This enum type describes the mode the spinbox will use to correct an Intermediate value if editing finishes.



See also correctionMode.

*/
type QAbstractSpinBox__CorrectionMode = int

// The spinbox will revert to the last valid value.
const QAbstractSpinBox__CorrectToPreviousValue QAbstractSpinBox__CorrectionMode = 0

// The spinbox will revert to the nearest valid value.
const QAbstractSpinBox__CorrectToNearestValue QAbstractSpinBox__CorrectionMode = 1

func (this *QAbstractSpinBox) CorrectionModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSpinBox_CorrectionModeItemName(val int) string {
	var nilthis *QAbstractSpinBox
	return nilthis.CorrectionModeItemName(val)
}

/*
ConstantValue
QAbstractSpinBox::DefaultStepType0
QAbstractSpinBox::AdaptiveDecimalStepType1

*/
type QAbstractSpinBox__StepType = int

//
const QAbstractSpinBox__DefaultStepType QAbstractSpinBox__StepType = 0

//
const QAbstractSpinBox__AdaptiveDecimalStepType QAbstractSpinBox__StepType = 1

func (this *QAbstractSpinBox) StepTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSpinBox_StepTypeItemName(val int) string {
	var nilthis *QAbstractSpinBox
	return nilthis.StepTypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10987() {
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
