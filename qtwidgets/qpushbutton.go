package qtwidgets

// /usr/include/qt/QtWidgets/qpushbutton.h
// #include <qpushbutton.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 96
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
func (this *QPushButton) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(QPaintEvent *)
func (this *QPushButton) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QPushButton) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QPushButton) InheritFocusInEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QPushButton) InheritFocusOutEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void initStyleOption(QStyleOptionButton *)
func (this *QPushButton) InheritInitStyleOption(f func(option *QStyleOptionButton /*777 QStyleOptionButton **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

/*

 */
type QPushButton struct {
	*QAbstractButton
}
type QPushButton_ITF interface {
	QAbstractButton_ITF
	QPushButton_PTR() *QPushButton
}

func (ptr *QPushButton) QPushButton_PTR() *QPushButton { return ptr }

func (this *QPushButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractButton.GetCthis()
	}
}
func (this *QPushButton) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractButton = NewQAbstractButtonFromPointer(cthis)
}
func NewQPushButtonFromPointer(cthis unsafe.Pointer) *QPushButton {
	bcthis0 := NewQAbstractButtonFromPointer(cthis)
	return &QPushButton{bcthis0}
}
func (*QPushButton) NewFromPointer(cthis unsafe.Pointer) *QPushButton {
	return NewQPushButtonFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QPushButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qpushbutton.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(QWidget *)

/*
Constructs a push button with no text and a parent.
*/
func (*QPushButton) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QPushButton {
	return NewQPushButton(parent)
}
func NewQPushButton(parent QWidget_ITF /*777 QWidget **/) *QPushButton {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPushButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(QWidget *)

/*
Constructs a push button with no text and a parent.
*/
func (*QPushButton) NewForInheritp() *QPushButton {
	return NewQPushButtonp()
}
func NewQPushButtonp() *QPushButton {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPushButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(const QString &, QWidget *)

/*
Constructs a push button with no text and a parent.
*/
func (*QPushButton) NewForInherit1(text string, parent QWidget_ITF /*777 QWidget **/) *QPushButton {
	return NewQPushButton1(text, parent)
}
func NewQPushButton1(text string, parent QWidget_ITF /*777 QWidget **/) *QPushButton {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPushButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(const QString &, QWidget *)

/*
Constructs a push button with no text and a parent.
*/
func (*QPushButton) NewForInherit1p(text string) *QPushButton {
	return NewQPushButton1p(text)
}
func NewQPushButton1p(text string) *QPushButton {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPushButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:66
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(const QIcon &, const QString &, QWidget *)

/*
Constructs a push button with no text and a parent.
*/
func (*QPushButton) NewForInherit2(icon qtgui.QIcon_ITF, text string, parent QWidget_ITF /*777 QWidget **/) *QPushButton {
	return NewQPushButton2(icon, text, parent)
}
func NewQPushButton2(icon qtgui.QIcon_ITF, text string, parent QWidget_ITF /*777 QWidget **/) *QPushButton {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg2 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonC2ERK5QIconRK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPushButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:66
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(const QIcon &, const QString &, QWidget *)

/*
Constructs a push button with no text and a parent.
*/
func (*QPushButton) NewForInherit2p(icon qtgui.QIcon_ITF, text string) *QPushButton {
	return NewQPushButton2p(icon, text)
}
func NewQPushButton2p(icon qtgui.QIcon_ITF, text string) *QPushButton {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonC2ERK5QIconRK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPushButton")
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPushButton()

/*

 */
func DeleteQPushButton(this *QPushButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QPushButton) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QPushButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoDefault() const

/*

 */
func (this *QPushButton) AutoDefault() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton11autoDefaultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoDefault(bool)

/*

 */
func (this *QPushButton) SetAutoDefault(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton14setAutoDefaultEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDefault() const

/*

 */
func (this *QPushButton) IsDefault() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton9isDefaultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefault(bool)

/*

 */
func (this *QPushButton) SetDefault(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton10setDefaultEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenu(QMenu *)

/*
Associates the popup menu menu with this push button. This turns the button into a menu button, which in some styles will produce a small triangle to the right of the button's text.

Ownership of the menu is not transferred to the push button.



A push button with popup menus shown in the Fusion widget style.

See also menu().
*/
func (this *QPushButton) SetMenu(menu QMenu_ITF /*777 QMenu **/) {
	var convArg0 unsafe.Pointer
	if menu != nil && menu.QMenu_PTR() != nil {
		convArg0 = menu.QMenu_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton7setMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * menu() const

/*
Returns the button's associated popup menu or 0 if no popup menu has been set.

See also setMenu().
*/
func (this *QPushButton) Menu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton4menuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qpushbutton.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlat(bool)

/*

 */
func (this *QPushButton) SetFlat(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton7setFlatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlat() const

/*

 */
func (this *QPushButton) IsFlat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton6isFlatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMenu()

/*
Shows (pops up) the associated popup menu. If there is no such menu, this function does nothing. This function does not return until the popup menu has been closed by the user.
*/
func (this *QPushButton) ShowMenu() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton8showMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QPushButton) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:92
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QPushButton) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:93
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QPushButton) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusInEvent().
*/
func (this *QPushButton) FocusInEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().
*/
func (this *QPushButton) FocusOutEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:96
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionButton *) const

/*
Initialize option with the values from this QPushButton. This method is useful for subclasses when they need a QStyleOptionButton, but don't want to fill in all the information themselves.

See also QStyleOption::initFrom().
*/
func (this *QPushButton) InitStyleOption(option QStyleOptionButton_ITF /*777 QStyleOptionButton **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionButton_PTR() != nil {
		convArg0 = option.QStyleOptionButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton15initStyleOptionEP18QStyleOptionButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
