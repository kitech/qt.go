package qtwidgets

// /usr/include/qt/QtWidgets/qgroupbox.h
// #include <qgroupbox.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 104
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
// bool event(class QEvent *)
func (this *QGroupBox) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

// void childEvent(class QChildEvent *)
func (this *QGroupBox) InheritChildEvent(f func(event *qtcore.QChildEvent /*777 QChildEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "childEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QGroupBox) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QGroupBox) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QGroupBox) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QGroupBox) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "changeEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QGroupBox) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QGroupBox) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QGroupBox) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void initStyleOption(class QStyleOptionGroupBox *)
func (this *QGroupBox) InheritInitStyleOption(f func(option *QStyleOptionGroupBox /*777 QStyleOptionGroupBox **/)) {
	ffiqt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QGroupBox struct {
	*QWidget
}

func (this *QGroupBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QGroupBox) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQGroupBoxFromPointer(cthis unsafe.Pointer) *QGroupBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QGroupBox{bcthis0}
}
func (*QGroupBox) NewFromPointer(cthis unsafe.Pointer) *QGroupBox {
	return NewQGroupBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGroupBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgroupbox.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(QWidget *)
func NewQGroupBox(parent *QWidget /*777 QWidget **/) *QGroupBox {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBoxC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGroupBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(const QString &, QWidget *)
func NewQGroupBox_1(title *qtcore.QString, parent *QWidget /*777 QWidget **/) *QGroupBox {
	var convArg0 = title.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBoxC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQGroupBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGroupBox()
func DeleteQGroupBox(this *QGroupBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBoxD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title()
func (this *QGroupBox) Title() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox5titleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qgroupbox.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)
func (this *QGroupBox) SetTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox8setTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment()
func (this *QGroupBox) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(int)
func (this *QGroupBox) SetAlignment(alignment int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox12setAlignmentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QGroupBox) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qgroupbox.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlat()
func (this *QGroupBox) IsFlat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox6isFlatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlat(_Bool)
func (this *QGroupBox) SetFlat(flat bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox7setFlatEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCheckable()
func (this *QGroupBox) IsCheckable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox11isCheckableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCheckable(_Bool)
func (this *QGroupBox) SetCheckable(checkable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox12setCheckableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), checkable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isChecked()
func (this *QGroupBox) IsChecked() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox9isCheckedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChecked(_Bool)
func (this *QGroupBox) SetChecked(checked bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox10setCheckedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(_Bool)
func (this *QGroupBox) Clicked(checked bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox7clickedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggled(_Bool)
func (this *QGroupBox) Toggled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox7toggledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:88
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QGroupBox) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:89
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)
func (this *QGroupBox) ChildEvent(event *qtcore.QChildEvent /*777 QChildEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox10childEventEP11QChildEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QGroupBox) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QGroupBox) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:92
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QGroupBox) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:93
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QGroupBox) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QGroupBox) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QGroupBox) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QGroupBox) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:97
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionGroupBox *)
func (this *QGroupBox) InitStyleOption(option *QStyleOptionGroupBox /*777 QStyleOptionGroupBox **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox15initStyleOptionEP20QStyleOptionGroupBox", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
