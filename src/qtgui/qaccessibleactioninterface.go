//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QAccessibleActionInterface struct {
	*qtrt.CObject
}

// /usr/include/qt/QtGui/qaccessible.h:630
// index:0
// virtual
// void ~QAccessibleActionInterface()
func DeleteQAccessibleActionInterface(*QAccessibleActionInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:632
// index:0
// pure virtual
// QStringList actionNames()
func (this *QAccessibleActionInterface) ActionNames() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleActionInterface11actionNamesEv", ffiqt.FFI_TYPE_VOID, this.Cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:633
// index:0
// virtual
// QString localizedActionName(const class QString &)
func (this *QAccessibleActionInterface) LocalizedActionName(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleActionInterface19localizedActionNameERK7QString", ffiqt.FFI_TYPE_VOID, this.Cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:634
// index:0
// virtual
// QString localizedActionDescription(const class QString &)
func (this *QAccessibleActionInterface) LocalizedActionDescription(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleActionInterface26localizedActionDescriptionERK7QString", ffiqt.FFI_TYPE_VOID, this.Cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:635
// index:0
// pure virtual
// void doAction(const class QString &)
func (this *QAccessibleActionInterface) DoAction(actionName unsafe.Pointer) {
	// 0: (, actionName const QString &), (actionName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface8doActionERK7QString", ffiqt.FFI_TYPE_VOID, this.Cthis, actionName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:636
// index:0
// pure virtual
// QStringList keyBindingsForAction(const class QString &)
func (this *QAccessibleActionInterface) KeyBindingsForAction(actionName unsafe.Pointer) {
	// 0: (, actionName const QString &), (actionName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAccessibleActionInterface20keyBindingsForActionERK7QString", ffiqt.FFI_TYPE_VOID, this.Cthis, actionName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:638
// index:0
// static
// const QString & pressAction()
func (this *QAccessibleActionInterface) PressAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface11pressActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_PressAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.PressAction()
}

// /usr/include/qt/QtGui/qaccessible.h:639
// index:0
// static
// const QString & increaseAction()
func (this *QAccessibleActionInterface) IncreaseAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface14increaseActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_IncreaseAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.IncreaseAction()
}

// /usr/include/qt/QtGui/qaccessible.h:640
// index:0
// static
// const QString & decreaseAction()
func (this *QAccessibleActionInterface) DecreaseAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface14decreaseActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_DecreaseAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.DecreaseAction()
}

// /usr/include/qt/QtGui/qaccessible.h:641
// index:0
// static
// const QString & showMenuAction()
func (this *QAccessibleActionInterface) ShowMenuAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface14showMenuActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_ShowMenuAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.ShowMenuAction()
}

// /usr/include/qt/QtGui/qaccessible.h:642
// index:0
// static
// const QString & setFocusAction()
func (this *QAccessibleActionInterface) SetFocusAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface14setFocusActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_SetFocusAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.SetFocusAction()
}

// /usr/include/qt/QtGui/qaccessible.h:643
// index:0
// static
// const QString & toggleAction()
func (this *QAccessibleActionInterface) ToggleAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface12toggleActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_ToggleAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.ToggleAction()
}

// /usr/include/qt/QtGui/qaccessible.h:644
// index:0
// static
// QString scrollLeftAction()
func (this *QAccessibleActionInterface) ScrollLeftAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface16scrollLeftActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_ScrollLeftAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.ScrollLeftAction()
}

// /usr/include/qt/QtGui/qaccessible.h:645
// index:0
// static
// QString scrollRightAction()
func (this *QAccessibleActionInterface) ScrollRightAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface17scrollRightActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_ScrollRightAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.ScrollRightAction()
}

// /usr/include/qt/QtGui/qaccessible.h:646
// index:0
// static
// QString scrollUpAction()
func (this *QAccessibleActionInterface) ScrollUpAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface14scrollUpActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_ScrollUpAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.ScrollUpAction()
}

// /usr/include/qt/QtGui/qaccessible.h:647
// index:0
// static
// QString scrollDownAction()
func (this *QAccessibleActionInterface) ScrollDownAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface16scrollDownActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_ScrollDownAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.ScrollDownAction()
}

// /usr/include/qt/QtGui/qaccessible.h:648
// index:0
// static
// QString nextPageAction()
func (this *QAccessibleActionInterface) NextPageAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface14nextPageActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_NextPageAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.NextPageAction()
}

// /usr/include/qt/QtGui/qaccessible.h:649
// index:0
// static
// QString previousPageAction()
func (this *QAccessibleActionInterface) PreviousPageAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAccessibleActionInterface18previousPageActionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessibleActionInterface_PreviousPageAction() {
	// 0: (), ()
	var nilthis *QAccessibleActionInterface
	nilthis.PreviousPageAction()
}

//  body block end
