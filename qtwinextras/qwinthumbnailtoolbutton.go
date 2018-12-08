package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h
// #include <qwinthumbnailtoolbutton.h>
// #include <Qtwinextras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"

//  ext block end

//  body block begin

/*

 */
type QWinThumbnailToolButton struct {
	*qtcore.QObject
}
type QWinThumbnailToolButton_ITF interface {
	qtcore.QObject_ITF
	QWinThumbnailToolButton_PTR() *QWinThumbnailToolButton
}

func (ptr *QWinThumbnailToolButton) QWinThumbnailToolButton_PTR() *QWinThumbnailToolButton { return ptr }

func (this *QWinThumbnailToolButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWinThumbnailToolButton) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWinThumbnailToolButtonFromPointer(cthis unsafe.Pointer) *QWinThumbnailToolButton {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWinThumbnailToolButton{bcthis0}
}
func (*QWinThumbnailToolButton) NewFromPointer(cthis unsafe.Pointer) *QWinThumbnailToolButton {
	return NewQWinThumbnailToolButtonFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWinThumbnailToolButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinThumbnailToolButton(QObject *)

/*
Constructs a QWinThumbnailToolButton with the specified parent.
*/
func (*QWinThumbnailToolButton) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QWinThumbnailToolButton {
	return NewQWinThumbnailToolButton(parent)
}
func NewQWinThumbnailToolButton(parent qtcore.QObject_ITF /*777 QObject **/) *QWinThumbnailToolButton {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButtonC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWinThumbnailToolButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWinThumbnailToolButton")
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinThumbnailToolButton(QObject *)

/*
Constructs a QWinThumbnailToolButton with the specified parent.
*/
func (*QWinThumbnailToolButton) NewForInheritp() *QWinThumbnailToolButton {
	return NewQWinThumbnailToolButtonp()
}
func NewQWinThumbnailToolButtonp() *QWinThumbnailToolButton {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButtonC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWinThumbnailToolButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWinThumbnailToolButton")
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinThumbnailToolButton()

/*

 */
func DeleteQWinThumbnailToolButton(this *QWinThumbnailToolButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*

 */
func (this *QWinThumbnailToolButton) SetToolTip(toolTip string) {
	var tmpArg0 = qtcore.NewQString5(toolTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toolTip() const

/*

 */
func (this *QWinThumbnailToolButton) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)

/*

 */
func (this *QWinThumbnailToolButton) SetIcon(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon() const

/*

 */
func (this *QWinThumbnailToolButton) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*

 */
func (this *QWinThumbnailToolButton) SetEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled() const

/*

 */
func (this *QWinThumbnailToolButton) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInteractive(bool)

/*

 */
func (this *QWinThumbnailToolButton) SetInteractive(interactive bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton14setInteractiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interactive)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInteractive() const

/*

 */
func (this *QWinThumbnailToolButton) IsInteractive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton13isInteractiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*

 */
func (this *QWinThumbnailToolButton) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible() const

/*

 */
func (this *QWinThumbnailToolButton) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDismissOnClick(bool)

/*

 */
func (this *QWinThumbnailToolButton) SetDismissOnClick(dismiss bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton17setDismissOnClickEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dismiss)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dismissOnClick() const

/*

 */
func (this *QWinThumbnailToolButton) DismissOnClick() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton14dismissOnClickEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlat(bool)

/*

 */
func (this *QWinThumbnailToolButton) SetFlat(flat bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton7setFlatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flat)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlat() const

/*

 */
func (this *QWinThumbnailToolButton) IsFlat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton6isFlatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void click()

/*
Performs a click. The clicked() signal is emitted as appropriate.

This function does nothing if the button is disabled or non-interactive.
*/
func (this *QWinThumbnailToolButton) Click() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton5clickEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked()

/*
This signal is emitted when the button is clicked.
*/
func (this *QWinThumbnailToolButton) Clicked() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton7clickedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void changed()

/*

 */
func (this *QWinThumbnailToolButton) Changed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton7changedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
	if false {
		qtwidgets.KeepMe()
	}
}

//  keep block end
