package qtwidgets

// /usr/include/qt/QtWidgets/qactiongroup.h
// #include <qactiongroup.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 62
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

/*

 */
type QActionGroup struct {
	*qtcore.QObject
}
type QActionGroup_ITF interface {
	qtcore.QObject_ITF
	QActionGroup_PTR() *QActionGroup
}

func (ptr *QActionGroup) QActionGroup_PTR() *QActionGroup { return ptr }

func (this *QActionGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QActionGroup) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQActionGroupFromPointer(cthis unsafe.Pointer) *QActionGroup {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QActionGroup{bcthis0}
}
func (*QActionGroup) NewFromPointer(cthis unsafe.Pointer) *QActionGroup {
	return NewQActionGroupFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QActionGroup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionGroup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qactiongroup.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QActionGroup(QObject *)

/*
Constructs an action group for the parent object.

The action group is exclusive by default. Call setExclusive(false) to make the action group non-exclusive.
*/
func (*QActionGroup) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QActionGroup {
	return NewQActionGroup(parent)
}
func NewQActionGroup(parent qtcore.QObject_ITF /*777 QObject **/) *QActionGroup {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQActionGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QActionGroup")
	return gothis
}

// /usr/include/qt/QtWidgets/qactiongroup.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QActionGroup()

/*

 */
func DeleteQActionGroup(this *QActionGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(QAction *)

/*
Adds the action to this group, and returns it.

Normally an action is added to a group by creating it with the group as its parent, so this function is not usually used.

See also QAction::setActionGroup().
*/
func (this *QActionGroup) AddAction(a QAction_ITF /*777 QAction **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if a != nil && a.QAction_PTR() != nil {
		convArg0 = a.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup9addActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qactiongroup.h:67
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QString &)

/*
Adds the action to this group, and returns it.

Normally an action is added to a group by creating it with the group as its parent, so this function is not usually used.

See also QAction::setActionGroup().
*/
func (this *QActionGroup) AddAction1(text string) *QAction /*777 QAction **/ {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup9addActionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qactiongroup.h:68
// index:2
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, const QString &)

/*
Adds the action to this group, and returns it.

Normally an action is added to a group by creating it with the group as its parent, so this function is not usually used.

See also QAction::setActionGroup().
*/
func (this *QActionGroup) AddAction2(icon qtgui.QIcon_ITF, text string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup9addActionERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qactiongroup.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAction(QAction *)

/*
Removes the action from this group. The action will have no parent as a result.

See also QAction::setActionGroup().
*/
func (this *QActionGroup) RemoveAction(a QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if a != nil && a.QAction_PTR() != nil {
		convArg0 = a.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup12removeActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * checkedAction() const

/*
Returns the currently checked action in the group, or 0 if none are checked.
*/
func (this *QActionGroup) CheckedAction() *QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionGroup13checkedActionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qactiongroup.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExclusive() const

/*

 */
func (this *QActionGroup) IsExclusive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionGroup11isExclusiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qactiongroup.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled() const

/*

 */
func (this *QActionGroup) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionGroup9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qactiongroup.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible() const

/*

 */
func (this *QActionGroup) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QActionGroup9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qactiongroup.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*

 */
func (this *QActionGroup) SetEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setDisabled(bool)

/*
This is a convenience function for the enabled property, that is useful for signals--slots connections. If b is true the action group is disabled; otherwise it is enabled.
*/
func (this *QActionGroup) SetDisabled(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup11setDisabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*

 */
func (this *QActionGroup) SetVisible(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExclusive(bool)

/*

 */
func (this *QActionGroup) SetExclusive(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup12setExclusiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggered(QAction *)

/*
This signal is emitted when the given action in the action group is activated by the user; for example, when the user clicks a menu option, toolbar button, or presses an action's shortcut key combination.

Connect to this signal for command actions.

See also QAction::activate().
*/
func (this *QActionGroup) Triggered(arg0 QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QAction_PTR() != nil {
		convArg0 = arg0.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup9triggeredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hovered(QAction *)

/*
This signal is emitted when the given action in the action group is highlighted by the user; for example, when the user pauses with the cursor over a menu option, toolbar button, or presses an action's shortcut key combination.

See also QAction::activate().
*/
func (this *QActionGroup) Hovered(arg0 QAction_ITF /*777 QAction **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QAction_PTR() != nil {
		convArg0 = arg0.QAction_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QActionGroup7hoveredEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
