package qtwidgets

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h
// #include <qdialogbuttonbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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

// void changeEvent(QEvent *)
func (this *QDialogButtonBox) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// bool event(QEvent *)
func (this *QDialogButtonBox) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QDialogButtonBox struct {
	*QWidget
}
type QDialogButtonBox_ITF interface {
	QWidget_ITF
	QDialogButtonBox_PTR() *QDialogButtonBox
}

func (ptr *QDialogButtonBox) QDialogButtonBox_PTR() *QDialogButtonBox { return ptr }

func (this *QDialogButtonBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QDialogButtonBox) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQDialogButtonBoxFromPointer(cthis unsafe.Pointer) *QDialogButtonBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDialogButtonBox{bcthis0}
}
func (*QDialogButtonBox) NewFromPointer(cthis unsafe.Pointer) *QDialogButtonBox {
	return NewQDialogButtonBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDialogButtonBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDialogButtonBox10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(QWidget *)

/*
Constructs an empty, horizontal button box with the given parent.

See also orientation and addButton().
*/
func (*QDialogButtonBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QDialogButtonBox {
	return NewQDialogButtonBox(parent)
}
func NewQDialogButtonBox(parent QWidget_ITF /*777 QWidget **/) *QDialogButtonBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDialogButtonBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(QWidget *)

/*
Constructs an empty, horizontal button box with the given parent.

See also orientation and addButton().
*/
func (*QDialogButtonBox) NewForInheritp() *QDialogButtonBox {
	return NewQDialogButtonBoxp()
}
func NewQDialogButtonBoxp() *QDialogButtonBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDialogButtonBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:121
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(Qt::Orientation, QWidget *)

/*
Constructs an empty, horizontal button box with the given parent.

See also orientation and addButton().
*/
func (*QDialogButtonBox) NewForInherit1(orientation int, parent QWidget_ITF /*777 QWidget **/) *QDialogButtonBox {
	return NewQDialogButtonBox1(orientation, parent)
}
func NewQDialogButtonBox1(orientation int, parent QWidget_ITF /*777 QWidget **/) *QDialogButtonBox {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2EN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, orientation, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDialogButtonBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:121
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(Qt::Orientation, QWidget *)

/*
Constructs an empty, horizontal button box with the given parent.

See also orientation and addButton().
*/
func (*QDialogButtonBox) NewForInherit1p(orientation int) *QDialogButtonBox {
	return NewQDialogButtonBox1p(orientation)
}
func NewQDialogButtonBox1p(orientation int) *QDialogButtonBox {
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2EN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, orientation, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDialogButtonBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:122
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(QDialogButtonBox::StandardButtons, QWidget *)

/*
Constructs an empty, horizontal button box with the given parent.

See also orientation and addButton().
*/
func (*QDialogButtonBox) NewForInherit2(buttons int, parent QWidget_ITF /*777 QWidget **/) *QDialogButtonBox {
	return NewQDialogButtonBox2(buttons, parent)
}
func NewQDialogButtonBox2(buttons int, parent QWidget_ITF /*777 QWidget **/) *QDialogButtonBox {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2E6QFlagsINS_14StandardButtonEEP7QWidget", qtrt.FFI_TYPE_POINTER, buttons, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDialogButtonBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:122
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(QDialogButtonBox::StandardButtons, QWidget *)

/*
Constructs an empty, horizontal button box with the given parent.

See also orientation and addButton().
*/
func (*QDialogButtonBox) NewForInherit2p(buttons int) *QDialogButtonBox {
	return NewQDialogButtonBox2p(buttons)
}
func NewQDialogButtonBox2p(buttons int) *QDialogButtonBox {
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2E6QFlagsINS_14StandardButtonEEP7QWidget", qtrt.FFI_TYPE_POINTER, buttons, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDialogButtonBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:123
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(QDialogButtonBox::StandardButtons, Qt::Orientation, QWidget *)

/*
Constructs an empty, horizontal button box with the given parent.

See also orientation and addButton().
*/
func (*QDialogButtonBox) NewForInherit3(buttons int, orientation int, parent QWidget_ITF /*777 QWidget **/) *QDialogButtonBox {
	return NewQDialogButtonBox3(buttons, orientation, parent)
}
func NewQDialogButtonBox3(buttons int, orientation int, parent QWidget_ITF /*777 QWidget **/) *QDialogButtonBox {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg2 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2E6QFlagsINS_14StandardButtonEEN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, buttons, orientation, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDialogButtonBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:123
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDialogButtonBox(QDialogButtonBox::StandardButtons, Qt::Orientation, QWidget *)

/*
Constructs an empty, horizontal button box with the given parent.

See also orientation and addButton().
*/
func (*QDialogButtonBox) NewForInherit3p(buttons int, orientation int) *QDialogButtonBox {
	return NewQDialogButtonBox3p(buttons, orientation)
}
func NewQDialogButtonBox3p(buttons int, orientation int) *QDialogButtonBox {
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBoxC2E6QFlagsINS_14StandardButtonEEN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, buttons, orientation, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialogButtonBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDialogButtonBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:125
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDialogButtonBox()

/*

 */
func DeleteQDialogButtonBox(this *QDialogButtonBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)

/*

 */
func (this *QDialogButtonBox) SetOrientation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation() const

/*

 */
func (this *QDialogButtonBox) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDialogButtonBox11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addButton(QAbstractButton *, QDialogButtonBox::ButtonRole)

/*
Adds the given button to the button box with the specified role. If the role is invalid, the button is not added.

If the button has already been added, it is removed and added again with the new role.

Note: The button box takes ownership of the button.

See also removeButton() and clear().
*/
func (this *QDialogButtonBox) AddButton(button QAbstractButton_ITF /*777 QAbstractButton **/, role int) {
	var convArg0 unsafe.Pointer
	if button != nil && button.QAbstractButton_PTR() != nil {
		convArg0 = button.QAbstractButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonEP15QAbstractButtonNS_10ButtonRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:131
// index:1
// Public Visibility=Default Availability=Available
// [8] QPushButton * addButton(const QString &, QDialogButtonBox::ButtonRole)

/*
Adds the given button to the button box with the specified role. If the role is invalid, the button is not added.

If the button has already been added, it is removed and added again with the new role.

Note: The button box takes ownership of the button.

See also removeButton() and clear().
*/
func (this *QDialogButtonBox) AddButton1(text string, role int) *QPushButton /*777 QPushButton **/ {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonERK7QStringNS_10ButtonRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:132
// index:2
// Public Visibility=Default Availability=Available
// [8] QPushButton * addButton(QDialogButtonBox::StandardButton)

/*
Adds the given button to the button box with the specified role. If the role is invalid, the button is not added.

If the button has already been added, it is removed and added again with the new role.

Note: The button box takes ownership of the button.

See also removeButton() and clear().
*/
func (this *QDialogButtonBox) AddButton2(button int) *QPushButton /*777 QPushButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox9addButtonENS_14StandardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), button)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeButton(QAbstractButton *)

/*
Removes button from the button box without deleting it and sets its parent to zero.

See also clear(), buttons(), and addButton().
*/
func (this *QDialogButtonBox) RemoveButton(button QAbstractButton_ITF /*777 QAbstractButton **/) {
	var convArg0 unsafe.Pointer
	if button != nil && button.QAbstractButton_PTR() != nil {
		convArg0 = button.QAbstractButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox12removeButtonEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the button box, deleting all buttons within it.

See also removeButton() and addButton().
*/
func (this *QDialogButtonBox) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] QDialogButtonBox::ButtonRole buttonRole(QAbstractButton *) const

/*
Returns the button role for the specified button. This function returns InvalidRole if button is 0 or has not been added to the button box.

See also buttons() and addButton().
*/
func (this *QDialogButtonBox) ButtonRole(button QAbstractButton_ITF /*777 QAbstractButton **/) int {
	var convArg0 unsafe.Pointer
	if button != nil && button.QAbstractButton_PTR() != nil {
		convArg0 = button.QAbstractButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDialogButtonBox10buttonRoleEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardButtons(QDialogButtonBox::StandardButtons)

/*

 */
func (this *QDialogButtonBox) SetStandardButtons(buttons int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox18setStandardButtonsE6QFlagsINS_14StandardButtonEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] QDialogButtonBox::StandardButtons standardButtons() const

/*

 */
func (this *QDialogButtonBox) StandardButtons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDialogButtonBox15standardButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:141
// index:0
// Public Visibility=Default Availability=Available
// [4] QDialogButtonBox::StandardButton standardButton(QAbstractButton *) const

/*
Returns the standard button enum value corresponding to the given button, or NoButton if the given button isn't a standard button.

See also button(), buttons(), and standardButtons().
*/
func (this *QDialogButtonBox) StandardButton(button QAbstractButton_ITF /*777 QAbstractButton **/) int {
	var convArg0 unsafe.Pointer
	if button != nil && button.QAbstractButton_PTR() != nil {
		convArg0 = button.QAbstractButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDialogButtonBox14standardButtonEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:142
// index:0
// Public Visibility=Default Availability=Available
// [8] QPushButton * button(QDialogButtonBox::StandardButton) const

/*
Returns the QPushButton corresponding to the standard button which, or 0 if the standard button doesn't exist in this button box.

See also standardButton(), standardButtons(), and buttons().
*/
func (this *QDialogButtonBox) Button(which int) *QPushButton /*777 QPushButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDialogButtonBox6buttonENS_14StandardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCenterButtons(bool)

/*

 */
func (this *QDialogButtonBox) SetCenterButtons(center bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox16setCenterButtonsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), center)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:145
// index:0
// Public Visibility=Default Availability=Available
// [1] bool centerButtons() const

/*

 */
func (this *QDialogButtonBox) CenterButtons() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QDialogButtonBox13centerButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked(QAbstractButton *)

/*
This signal is emitted when a button inside the button box is clicked. The specific button that was pressed is specified by button.

See also accepted(), rejected(), and helpRequested().
*/
func (this *QDialogButtonBox) Clicked(button QAbstractButton_ITF /*777 QAbstractButton **/) {
	var convArg0 unsafe.Pointer
	if button != nil && button.QAbstractButton_PTR() != nil {
		convArg0 = button.QAbstractButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox7clickedEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accepted()

/*
This signal is emitted when a button inside the button box is clicked, as long as it was defined with the AcceptRole or YesRole.

See also rejected(), clicked(), and helpRequested().
*/
func (this *QDialogButtonBox) Accepted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox8acceptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void helpRequested()

/*
This signal is emitted when a button inside the button box is clicked, as long as it was defined with the HelpRole.

See also accepted(), rejected(), and clicked().
*/
func (this *QDialogButtonBox) HelpRequested() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox13helpRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rejected()

/*
This signal is emitted when a button inside the button box is clicked, as long as it was defined with the RejectRole or NoRole.

See also accepted(), helpRequested(), and clicked().
*/
func (this *QDialogButtonBox) Rejected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox8rejectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:154
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QDialogButtonBox) ChangeEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialogbuttonbox.h:155
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QDialogButtonBox) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QDialogButtonBox5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum describes the roles that can be used to describe buttons in the button box. Combinations of these roles are as flags used to describe different aspects of their behavior.



See also StandardButton.

*/
type QDialogButtonBox__ButtonRole = int

//
const QDialogButtonBox__InvalidRole QDialogButtonBox__ButtonRole = -1

// Clicking the button causes the dialog to be accepted (e.g. OK).
const QDialogButtonBox__AcceptRole QDialogButtonBox__ButtonRole = 0

// Clicking the button causes the dialog to be rejected (e.g. Cancel).
const QDialogButtonBox__RejectRole QDialogButtonBox__ButtonRole = 1

// Clicking the button causes a destructive change (e.g. for Discarding Changes) and closes the dialog.
const QDialogButtonBox__DestructiveRole QDialogButtonBox__ButtonRole = 2

// Clicking the button causes changes to the elements within the dialog.
const QDialogButtonBox__ActionRole QDialogButtonBox__ButtonRole = 3

// The button can be clicked to request help.
const QDialogButtonBox__HelpRole QDialogButtonBox__ButtonRole = 4

// The button is a "Yes"-like button.
const QDialogButtonBox__YesRole QDialogButtonBox__ButtonRole = 5

// The button is a "No"-like button.
const QDialogButtonBox__NoRole QDialogButtonBox__ButtonRole = 6

// The button resets the dialog's fields to default values.
const QDialogButtonBox__ResetRole QDialogButtonBox__ButtonRole = 7

// The button applies current changes.
const QDialogButtonBox__ApplyRole QDialogButtonBox__ButtonRole = 8

//
const QDialogButtonBox__NRoles QDialogButtonBox__ButtonRole = 9

func (this *QDialogButtonBox) ButtonRoleItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QDialogButtonBox_ButtonRoleItemName(val int) string {
	var nilthis *QDialogButtonBox
	return nilthis.ButtonRoleItemName(val)
}

/*


 */
type QDialogButtonBox__StandardButton = int

//
const QDialogButtonBox__NoButton QDialogButtonBox__StandardButton = 0

//
const QDialogButtonBox__Ok QDialogButtonBox__StandardButton = 1024

//
const QDialogButtonBox__Save QDialogButtonBox__StandardButton = 2048

//
const QDialogButtonBox__SaveAll QDialogButtonBox__StandardButton = 4096

//
const QDialogButtonBox__Open QDialogButtonBox__StandardButton = 8192

//
const QDialogButtonBox__Yes QDialogButtonBox__StandardButton = 16384

//
const QDialogButtonBox__YesToAll QDialogButtonBox__StandardButton = 32768

//
const QDialogButtonBox__No QDialogButtonBox__StandardButton = 65536

//
const QDialogButtonBox__NoToAll QDialogButtonBox__StandardButton = 131072

//
const QDialogButtonBox__Abort QDialogButtonBox__StandardButton = 262144

//
const QDialogButtonBox__Retry QDialogButtonBox__StandardButton = 524288

//
const QDialogButtonBox__Ignore QDialogButtonBox__StandardButton = 1048576

//
const QDialogButtonBox__Close QDialogButtonBox__StandardButton = 2097152

//
const QDialogButtonBox__Cancel QDialogButtonBox__StandardButton = 4194304

//
const QDialogButtonBox__Discard QDialogButtonBox__StandardButton = 8388608

//
const QDialogButtonBox__Help QDialogButtonBox__StandardButton = 16777216

//
const QDialogButtonBox__Apply QDialogButtonBox__StandardButton = 33554432

//
const QDialogButtonBox__Reset QDialogButtonBox__StandardButton = 67108864

//
const QDialogButtonBox__RestoreDefaults QDialogButtonBox__StandardButton = 134217728

//
const QDialogButtonBox__FirstButton QDialogButtonBox__StandardButton = 1024

//
const QDialogButtonBox__LastButton QDialogButtonBox__StandardButton = 134217728

func (this *QDialogButtonBox) StandardButtonItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QDialogButtonBox_StandardButtonItemName(val int) string {
	var nilthis *QDialogButtonBox
	return nilthis.StandardButtonItemName(val)
}

/*
This enum describes the layout policy to be used when arranging the buttons contained in the button box.



The button layout is specified by the current style. However, on the X11 platform, it may be influenced by the desktop environment.

*/
type QDialogButtonBox__ButtonLayout = int

// Use a policy appropriate for applications on Windows.
const QDialogButtonBox__WinLayout QDialogButtonBox__ButtonLayout = 0

// Use a policy appropriate for applications on macOS.
const QDialogButtonBox__MacLayout QDialogButtonBox__ButtonLayout = 1

// Use a policy appropriate for applications on KDE.
const QDialogButtonBox__KdeLayout QDialogButtonBox__ButtonLayout = 2

// Use a policy appropriate for applications on GNOME.
const QDialogButtonBox__GnomeLayout QDialogButtonBox__ButtonLayout = 3

//
const QDialogButtonBox__AndroidLayout QDialogButtonBox__ButtonLayout = 5

func (this *QDialogButtonBox) ButtonLayoutItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QDialogButtonBox_ButtonLayoutItemName(val int) string {
	var nilthis *QDialogButtonBox
	return nilthis.ButtonLayoutItemName(val)
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
