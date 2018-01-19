//  header block begin
// /usr/include/qt/QtWidgets/qactiongroup.h
// #include <qactiongroup.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 64
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
type QActionGroup struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qactiongroup.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QActionGroup) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:63
// index:0
// void QActionGroup(class QObject *)
func NewQActionGroup(parent unsafe.Pointer) *QActionGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QActionGroup{cthis}
}

// /usr/include/qt/QtWidgets/qactiongroup.h:64
// index:0
// virtual
// void ~QActionGroup()
func DeleteQActionGroup(*QActionGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:66
// index:0
// QAction * addAction(class QAction *)
func (this *QActionGroup) AddAction(a unsafe.Pointer) {
	// 0: (, QAction * a), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup9addActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:67
// index:1
// QAction * addAction(const class QString &)
func (this *QActionGroup) AddAction_1(text unsafe.Pointer) {
	// 1: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup9addActionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:68
// index:2
// QAction * addAction(const class QIcon &, const class QString &)
func (this *QActionGroup) AddAction_2(icon unsafe.Pointer, text unsafe.Pointer) {
	// 2: (, const QIcon & icon, const QString & text), (icon, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup9addActionERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, icon, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:69
// index:0
// void removeAction(class QAction *)
func (this *QActionGroup) RemoveAction(a unsafe.Pointer) {
	// 0: (, QAction * a), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup12removeActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:70
// index:0
// QList<QAction *> actions()
func (this *QActionGroup) Actions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup7actionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:72
// index:0
// QAction * checkedAction()
func (this *QActionGroup) CheckedAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup13checkedActionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:73
// index:0
// bool isExclusive()
func (this *QActionGroup) IsExclusive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup11isExclusiveEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:74
// index:0
// bool isEnabled()
func (this *QActionGroup) IsEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup9isEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:75
// index:0
// bool isVisible()
func (this *QActionGroup) IsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup9isVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:79
// index:0
// void setEnabled(_Bool)
func (this *QActionGroup) SetEnabled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup10setEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:80
// index:0
// inline
// void setDisabled(_Bool)
func (this *QActionGroup) SetDisabled(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup11setDisabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:81
// index:0
// void setVisible(_Bool)
func (this *QActionGroup) SetVisible(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:82
// index:0
// void setExclusive(_Bool)
func (this *QActionGroup) SetExclusive(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup12setExclusiveEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:85
// index:0
// void triggered(class QAction *)
func (this *QActionGroup) Triggered(arg0 unsafe.Pointer) {
	// 0: (, QAction * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup9triggeredEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:86
// index:0
// void hovered(class QAction *)
func (this *QActionGroup) Hovered(arg0 unsafe.Pointer) {
	// 0: (, QAction * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup7hoveredEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
