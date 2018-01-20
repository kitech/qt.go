//  header block begin
// /usr/include/qt/QtWidgets/qactiongroup.h
// #include <qactiongroup.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 65
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
	*qtcore.QObject
}

func (this *QActionGroup) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQActionGroupFromPointer(cthis unsafe.Pointer) *QActionGroup {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QActionGroup{bcthis0}
}

// /usr/include/qt/QtWidgets/qactiongroup.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QActionGroup) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qactiongroup.h:63
// index:0
// Public
// void QActionGroup(class QObject *)
func NewQActionGroup(parent unsafe.Pointer) *QActionGroup {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQActionGroupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qactiongroup.h:64
// index:0
// Public virtual
// void ~QActionGroup()
func DeleteQActionGroup(*QActionGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:66
// index:0
// Public
// QAction * addAction(class QAction *)
func (this *QActionGroup) AddAction(a unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup9addActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qactiongroup.h:67
// index:1
// Public
// QAction * addAction(const class QString &)
func (this *QActionGroup) AddAction_1(text *qtcore.QString) interface{} {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup9addActionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qactiongroup.h:68
// index:2
// Public
// QAction * addAction(const class QIcon &, const class QString &)
func (this *QActionGroup) AddAction_2(icon *qtgui.QIcon, text *qtcore.QString) interface{} {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup9addActionERK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qactiongroup.h:69
// index:0
// Public
// void removeAction(class QAction *)
func (this *QActionGroup) RemoveAction(a unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup12removeActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:70
// index:0
// Public
// QList<QAction *> actions()
func (this *QActionGroup) Actions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup7actionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qactiongroup.h:72
// index:0
// Public
// QAction * checkedAction()
func (this *QActionGroup) CheckedAction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup13checkedActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qactiongroup.h:73
// index:0
// Public
// bool isExclusive()
func (this *QActionGroup) IsExclusive() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup11isExclusiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qactiongroup.h:74
// index:0
// Public
// bool isEnabled()
func (this *QActionGroup) IsEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qactiongroup.h:75
// index:0
// Public
// bool isVisible()
func (this *QActionGroup) IsVisible() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QActionGroup9isVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qactiongroup.h:79
// index:0
// Public
// void setEnabled(_Bool)
func (this *QActionGroup) SetEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:80
// index:0
// Public inline
// void setDisabled(_Bool)
func (this *QActionGroup) SetDisabled(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup11setDisabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:81
// index:0
// Public
// void setVisible(_Bool)
func (this *QActionGroup) SetVisible(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:82
// index:0
// Public
// void setExclusive(_Bool)
func (this *QActionGroup) SetExclusive(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup12setExclusiveEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:85
// index:0
// Public
// void triggered(class QAction *)
func (this *QActionGroup) Triggered(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup9triggeredEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qactiongroup.h:86
// index:0
// Public
// void hovered(class QAction *)
func (this *QActionGroup) Hovered(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QActionGroup7hoveredEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
