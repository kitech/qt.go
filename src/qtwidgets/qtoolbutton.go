//  header block begin
// /usr/include/qt/QtWidgets/qtoolbutton.h
// #include <qtoolbutton.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QToolButton struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QToolButton) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:74
// index:0
// void QToolButton(class QWidget *)
func NewQToolButton(parent unsafe.Pointer) *QToolButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QToolButton{cthis}
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:75
// index:0
// virtual
// void ~QToolButton()
func DeleteQToolButton(*QToolButton) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButtonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:77
// index:0
// virtual
// QSize sizeHint()
func (this *QToolButton) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:78
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QToolButton) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:80
// index:0
// Qt::ToolButtonStyle toolButtonStyle()
func (this *QToolButton) ToolButtonStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton15toolButtonStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:82
// index:0
// Qt::ArrowType arrowType()
func (this *QToolButton) ArrowType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton9arrowTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:83
// index:0
// void setArrowType(Qt::ArrowType)
func (this *QToolButton) SetArrowType(type_ int) {
	// 0: (, Qt::ArrowType type), (&type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton12setArrowTypeEN2Qt9ArrowTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:86
// index:0
// void setMenu(class QMenu *)
func (this *QToolButton) SetMenu(menu unsafe.Pointer) {
	// 0: (, QMenu * menu), (menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton7setMenuEP5QMenu", ffiqt.FFI_TYPE_VOID, this.cthis, menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:87
// index:0
// QMenu * menu()
func (this *QToolButton) Menu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton4menuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:89
// index:0
// void setPopupMode(enum QToolButton::ToolButtonPopupMode)
func (this *QToolButton) SetPopupMode(mode int) {
	// 0: (, QToolButton::ToolButtonPopupMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton12setPopupModeENS_19ToolButtonPopupModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:90
// index:0
// QToolButton::ToolButtonPopupMode popupMode()
func (this *QToolButton) PopupMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton9popupModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:93
// index:0
// QAction * defaultAction()
func (this *QToolButton) DefaultAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton13defaultActionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:95
// index:0
// void setAutoRaise(_Bool)
func (this *QToolButton) SetAutoRaise(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton12setAutoRaiseEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:96
// index:0
// bool autoRaise()
func (this *QToolButton) AutoRaise() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QToolButton9autoRaiseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:100
// index:0
// void showMenu()
func (this *QToolButton) ShowMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton8showMenuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:102
// index:0
// void setToolButtonStyle(Qt::ToolButtonStyle)
func (this *QToolButton) SetToolButtonStyle(style int) {
	// 0: (, Qt::ToolButtonStyle style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton18setToolButtonStyleEN2Qt15ToolButtonStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:103
// index:0
// void setDefaultAction(class QAction *)
func (this *QToolButton) SetDefaultAction(arg0 unsafe.Pointer) {
	// 0: (, QAction * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton16setDefaultActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbutton.h:106
// index:0
// void triggered(class QAction *)
func (this *QToolButton) Triggered(arg0 unsafe.Pointer) {
	// 0: (, QAction * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QToolButton9triggeredEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
