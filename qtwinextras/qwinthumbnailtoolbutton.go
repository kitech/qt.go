package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h
// #include <qwinthumbnailtoolbutton.h>
// #include <Qtwinextras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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

type QWinThumbnailToolButton struct {
	*qtcore.QObject
}

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
// [8] const QMetaObject * metaObject()
func (this *QWinThumbnailToolButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinThumbnailToolButton(QObject *)
func NewQWinThumbnailToolButton(parent *qtcore.QObject /*777 QObject **/) *QWinThumbnailToolButton {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButtonC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWinThumbnailToolButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinThumbnailToolButton()
func DeleteQWinThumbnailToolButton(this *QWinThumbnailToolButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)
func (this *QWinThumbnailToolButton) SetToolTip(toolTip string) {
	var tmpArg0 = qtcore.NewQString_5(toolTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toolTip()
func (this *QWinThumbnailToolButton) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QWinThumbnailToolButton) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QWinThumbnailToolButton) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QWinThumbnailToolButton) SetEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QWinThumbnailToolButton) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInteractive(_Bool)
func (this *QWinThumbnailToolButton) SetInteractive(interactive bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton14setInteractiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interactive)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInteractive()
func (this *QWinThumbnailToolButton) IsInteractive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton13isInteractiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QWinThumbnailToolButton) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QWinThumbnailToolButton) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDismissOnClick(_Bool)
func (this *QWinThumbnailToolButton) SetDismissOnClick(dismiss bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton17setDismissOnClickEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dismiss)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:78
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dismissOnClick()
func (this *QWinThumbnailToolButton) DismissOnClick() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton14dismissOnClickEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlat(_Bool)
func (this *QWinThumbnailToolButton) SetFlat(flat bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton7setFlatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlat()
func (this *QWinThumbnailToolButton) IsFlat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QWinThumbnailToolButton6isFlatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void click()
func (this *QWinThumbnailToolButton) Click() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton5clickEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clicked()
func (this *QWinThumbnailToolButton) Clicked() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton7clickedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbutton.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void changed()
func (this *QWinThumbnailToolButton) Changed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QWinThumbnailToolButton7changedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
