package qtwidgets

// /usr/include/qt/QtWidgets/qgroupbox.h
// #include <qgroupbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 104
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

// bool event(QEvent *)
func (this *QGroupBox) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void childEvent(QChildEvent *)
func (this *QGroupBox) InheritChildEvent(f func(event *qtcore.QChildEvent /*777 QChildEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "childEvent", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QGroupBox) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QGroupBox) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QGroupBox) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void changeEvent(QEvent *)
func (this *QGroupBox) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QGroupBox) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QGroupBox) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QGroupBox) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void initStyleOption(QStyleOptionGroupBox *)
func (this *QGroupBox) InheritInitStyleOption(f func(option *QStyleOptionGroupBox /*777 QStyleOptionGroupBox **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QGroupBox struct {
	*QWidget
}
type QGroupBox_ITF interface {
	QWidget_ITF
	QGroupBox_PTR() *QGroupBox
}

func (ptr *QGroupBox) QGroupBox_PTR() *QGroupBox { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGroupBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGroupBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgroupbox.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(QWidget *)

/*
Constructs a group box widget with the given parent but with no title.
*/
func NewQGroupBox(parent QWidget_ITF /*777 QWidget **/) *QGroupBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGroupBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(QWidget *)

/*
Constructs a group box widget with the given parent but with no title.
*/
func NewQGroupBox__() *QGroupBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGroupBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(const QString &, QWidget *)

/*
Constructs a group box widget with the given parent but with no title.
*/
func NewQGroupBox_1(title string, parent QWidget_ITF /*777 QWidget **/) *QGroupBox {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBoxC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGroupBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGroupBox(const QString &, QWidget *)

/*
Constructs a group box widget with the given parent but with no title.
*/
func NewQGroupBox_1_(title string) *QGroupBox {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBoxC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGroupBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGroupBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qgroupbox.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGroupBox()

/*

 */
func DeleteQGroupBox(this *QGroupBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title() const

/*

 */
func (this *QGroupBox) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGroupBox5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qgroupbox.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)

/*

 */
func (this *QGroupBox) SetTitle(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox8setTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const

/*

 */
func (this *QGroupBox) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGroupBox9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(int)

/*

 */
func (this *QGroupBox) SetAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox12setAlignmentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QGroupBox) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGroupBox15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qgroupbox.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlat() const

/*

 */
func (this *QGroupBox) IsFlat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGroupBox6isFlatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlat(bool)

/*

 */
func (this *QGroupBox) SetFlat(flat bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox7setFlatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flat)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCheckable() const

/*

 */
func (this *QGroupBox) IsCheckable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGroupBox11isCheckableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCheckable(bool)

/*

 */
func (this *QGroupBox) SetCheckable(checkable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox12setCheckableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checkable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isChecked() const

/*

 */
func (this *QGroupBox) IsChecked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGroupBox9isCheckedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChecked(bool)

/*

 */
func (this *QGroupBox) SetChecked(checked bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox10setCheckedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(bool)

/*
This signal is emitted when the check box is activated (i.e., pressed down then released while the mouse cursor is inside the button), or when the shortcut key is typed. Notably, this signal is not emitted if you call setChecked().

If the check box is checked, checked is true; it is false if the check box is unchecked.

This function was introduced in  Qt 4.2.

See also checkable, toggled(), and checked.
*/
func (this *QGroupBox) Clicked(checked bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox7clickedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(bool)

/*
This signal is emitted when the check box is activated (i.e., pressed down then released while the mouse cursor is inside the button), or when the shortcut key is typed. Notably, this signal is not emitted if you call setChecked().

If the check box is checked, checked is true; it is false if the check box is unchecked.

This function was introduced in  Qt 4.2.

See also checkable, toggled(), and checked.
*/
func (this *QGroupBox) Clicked__() {
	// arg: 0, bool=Bool, =Invalid,
	checked := false
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox7clickedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toggled(bool)

/*
If the group box is checkable, this signal is emitted when the check box is toggled. on is true if the check box is checked; otherwise, it is false.

Note: Notifier signal for property checked.

See also checkable.
*/
func (this *QGroupBox) Toggled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox7toggledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:88
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QGroupBox) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgroupbox.h:89
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)

/*
Reimplemented from QObject::childEvent().
*/
func (this *QGroupBox) ChildEvent(event qtcore.QChildEvent_ITF /*777 QChildEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QChildEvent_PTR() != nil {
		convArg0 = event.QChildEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox10childEventEP11QChildEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QGroupBox) ResizeEvent(event qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QResizeEvent_PTR() != nil {
		convArg0 = event.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QGroupBox) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:92
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().
*/
func (this *QGroupBox) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:93
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QGroupBox) ChangeEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QGroupBox) MousePressEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QGroupBox) MouseMoveEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QGroupBox) MouseReleaseEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QGroupBox17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:97
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionGroupBox *) const

/*
Initialize option with the values from this QGroupBox. This method is useful for subclasses when they need a QStyleOptionGroupBox, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QGroupBox) InitStyleOption(option QStyleOptionGroupBox_ITF /*777 QStyleOptionGroupBox **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionGroupBox_PTR() != nil {
		convArg0 = option.QStyleOptionGroupBox_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QGroupBox15initStyleOptionEP20QStyleOptionGroupBox", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
