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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// void resizeEvent(class QResizeEvent *)
func (this *QAbstractSpinBox) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractSpinBox) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QAbstractSpinBox) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QAbstractSpinBox) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QAbstractSpinBox) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractSpinBox) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QAbstractSpinBox) InheritContextMenuEvent(f func(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QAbstractSpinBox) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "changeEvent", f)
}

// void closeEvent(class QCloseEvent *)
func (this *QAbstractSpinBox) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "closeEvent", f)
}

// void hideEvent(class QHideEvent *)
func (this *QAbstractSpinBox) InheritHideEvent(f func(event *qtgui.QHideEvent /*777 QHideEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "hideEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractSpinBox) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QAbstractSpinBox) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "timerEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QAbstractSpinBox) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QAbstractSpinBox) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "showEvent", f)
}

// void initStyleOption(class QStyleOptionSpinBox *)
func (this *QAbstractSpinBox) InheritInitStyleOption(f func(option *QStyleOptionSpinBox /*777 QStyleOptionSpinBox **/)) {
	ffiqt.SetAllInheritCallback(this, "initStyleOption", f)
}

// QLineEdit * lineEdit()
func (this *QAbstractSpinBox) InheritLineEdit(f func() unsafe.Pointer /*666*/) {
	ffiqt.SetAllInheritCallback(this, "lineEdit", f)
}

// void setLineEdit(class QLineEdit *)
func (this *QAbstractSpinBox) InheritSetLineEdit(f func(edit *QLineEdit /*777 QLineEdit **/)) {
	ffiqt.SetAllInheritCallback(this, "setLineEdit", f)
}

// QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QAbstractSpinBox) InheritStepEnabled(f func() int) {
	ffiqt.SetAllInheritCallback(this, "stepEnabled", f)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractSpinBox(QWidget *)
func NewQAbstractSpinBox(parent *QWidget /*777 QWidget **/) *QAbstractSpinBox {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBoxC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractSpinBox()
func DeleteQAbstractSpinBox(this *QAbstractSpinBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBoxD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSpinBox::ButtonSymbols buttonSymbols()
func (this *QAbstractSpinBox) ButtonSymbols() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13buttonSymbolsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtonSymbols(enum QAbstractSpinBox::ButtonSymbols)
func (this *QAbstractSpinBox) SetButtonSymbols(bs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox16setButtonSymbolsENS_13ButtonSymbolsE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), bs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCorrectionMode(enum QAbstractSpinBox::CorrectionMode)
func (this *QAbstractSpinBox) SetCorrectionMode(cm int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox17setCorrectionModeENS_14CorrectionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cm)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractSpinBox::CorrectionMode correctionMode()
func (this *QAbstractSpinBox) CorrectionMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox14correctionModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAcceptableInput()
func (this *QAbstractSpinBox) HasAcceptableInput() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox18hasAcceptableInputEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QAbstractSpinBox) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QString specialValueText()
func (this *QAbstractSpinBox) SpecialValueText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16specialValueTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpecialValueText(const QString &)
func (this *QAbstractSpinBox) SetSpecialValueText(txt *qtcore.QString) {
	var convArg0 = txt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setSpecialValueTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wrapping()
func (this *QAbstractSpinBox) Wrapping() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8wrappingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWrapping(_Bool)
func (this *QAbstractSpinBox) SetWrapping(w bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setWrappingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(_Bool)
func (this *QAbstractSpinBox) SetReadOnly(r bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setReadOnlyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly()
func (this *QAbstractSpinBox) IsReadOnly() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox10isReadOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeyboardTracking(_Bool)
func (this *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox19setKeyboardTrackingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), kt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool keyboardTracking()
func (this *QAbstractSpinBox) KeyboardTracking() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16keyboardTrackingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)
func (this *QAbstractSpinBox) SetAlignment(flag int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flag)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment()
func (this *QAbstractSpinBox) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrame(_Bool)
func (this *QAbstractSpinBox) SetFrame(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox8setFrameEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFrame()
func (this *QAbstractSpinBox) HasFrame() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8hasFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccelerated(_Bool)
func (this *QAbstractSpinBox) SetAccelerated(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox14setAcceleratedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAccelerated()
func (this *QAbstractSpinBox) IsAccelerated() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox13isAcceleratedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGroupSeparatorShown(_Bool)
func (this *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox22setGroupSeparatorShownEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), shown)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isGroupSeparatorShown()
func (this *QAbstractSpinBox) IsGroupSeparatorShown() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox21isGroupSeparatorShownEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QAbstractSpinBox) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:120
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QAbstractSpinBox) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void interpretText()
func (this *QAbstractSpinBox) InterpretText() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13interpretTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:122
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAbstractSpinBox) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QAbstractSpinBox) InputMethodQuery(arg0 int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &)
func (this *QAbstractSpinBox) Validate(input *qtcore.QString, pos int) int {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &)
func (this *QAbstractSpinBox) Fixup(input *qtcore.QString) {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox5fixupER7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:129
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void stepBy(int)
func (this *QAbstractSpinBox) StepBy(steps int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepByEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stepUp()
func (this *QAbstractSpinBox) StepUp() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox6stepUpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stepDown()
func (this *QAbstractSpinBox) StepDown() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox8stepDownEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectAll()
func (this *QAbstractSpinBox) SelectAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9selectAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:134
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clear()
func (this *QAbstractSpinBox) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:136
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QAbstractSpinBox) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:137
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QAbstractSpinBox) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:138
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QAbstractSpinBox) KeyReleaseEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:140
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QAbstractSpinBox) WheelEvent(event *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QAbstractSpinBox) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:143
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QAbstractSpinBox) FocusOutEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QAbstractSpinBox) ContextMenuEvent(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:147
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QAbstractSpinBox) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:148
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)
func (this *QAbstractSpinBox) CloseEvent(event *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:149
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QAbstractSpinBox) HideEvent(event *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:150
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QAbstractSpinBox) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:151
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QAbstractSpinBox) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:152
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QAbstractSpinBox) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:153
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QAbstractSpinBox) TimerEvent(event *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:154
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QAbstractSpinBox) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:155
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QAbstractSpinBox) ShowEvent(event *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:156
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionSpinBox *)
func (this *QAbstractSpinBox) InitStyleOption(option *QStyleOptionSpinBox /*777 QStyleOptionSpinBox **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:158
// index:0
// Protected Visibility=Default Availability=Available
// [8] QLineEdit * lineEdit()
func (this *QAbstractSpinBox) LineEdit() *QLineEdit /*777 QLineEdit **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox8lineEditEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:159
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setLineEdit(QLineEdit *)
func (this *QAbstractSpinBox) SetLineEdit(edit *QLineEdit /*777 QLineEdit **/) {
	var convArg0 = edit.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:161
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QAbstractSpinBox) StepEnabled() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAbstractSpinBox11stepEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractspinbox.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editingFinished()
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
