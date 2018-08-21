package qtwidgets

// /usr/include/qt/QtWidgets/qcolordialog.h
// #include <qcolordialog.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
func (this *QColorDialog) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void done(int)
func (this *QColorDialog) InheritDone(f func(result int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "done", f)
}

/*

 */
type QColorDialog struct {
	*QDialog
}
type QColorDialog_ITF interface {
	QDialog_ITF
	QColorDialog_PTR() *QColorDialog
}

func (ptr *QColorDialog) QColorDialog_PTR() *QColorDialog { return ptr }

func (this *QColorDialog) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QColorDialog) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQColorDialogFromPointer(cthis unsafe.Pointer) *QColorDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QColorDialog{bcthis0}
}
func (*QColorDialog) NewFromPointer(cthis unsafe.Pointer) *QColorDialog {
	return NewQColorDialogFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QColorDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QColorDialog10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcolordialog.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QColorDialog(QWidget *)

/*
Constructs a color dialog with the given parent.

This function was introduced in  Qt 4.5.
*/
func NewQColorDialog(parent QWidget_ITF /*777 QWidget **/) *QColorDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialogC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQColorDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QColorDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qcolordialog.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QColorDialog(QWidget *)

/*
Constructs a color dialog with the given parent.

This function was introduced in  Qt 4.5.
*/
func NewQColorDialog__() *QColorDialog {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialogC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQColorDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QColorDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qcolordialog.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QColorDialog(const QColor &, QWidget *)

/*
Constructs a color dialog with the given parent.

This function was introduced in  Qt 4.5.
*/
func NewQColorDialog_1(initial qtgui.QColor_ITF, parent QWidget_ITF /*777 QWidget **/) *QColorDialog {
	var convArg0 unsafe.Pointer
	if initial != nil && initial.QColor_PTR() != nil {
		convArg0 = initial.QColor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialogC2ERK6QColorP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQColorDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QColorDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qcolordialog.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QColorDialog(const QColor &, QWidget *)

/*
Constructs a color dialog with the given parent.

This function was introduced in  Qt 4.5.
*/
func NewQColorDialog_1_(initial qtgui.QColor_ITF) *QColorDialog {
	var convArg0 unsafe.Pointer
	if initial != nil && initial.QColor_PTR() != nil {
		convArg0 = initial.QColor_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialogC2ERK6QColorP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQColorDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QColorDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qcolordialog.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QColorDialog()

/*

 */
func DeleteQColorDialog(this *QColorDialog) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialogD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentColor(const QColor &)

/*

 */
func (this *QColorDialog) SetCurrentColor(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog15setCurrentColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:76
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor currentColor() const

/*

 */
func (this *QColorDialog) CurrentColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QColorDialog12currentColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qcolordialog.h:78
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor selectedColor() const

/*
Returns the color that the user selected by clicking the OK or equivalent button.

Note: This color is not always the same as the color held by the currentColor property since the user can choose different colors before finally selecting the one to use.
*/
func (this *QColorDialog) SelectedColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QColorDialog13selectedColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qcolordialog.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QColorDialog::ColorDialogOption, bool)

/*
Sets the given option to be enabled if on is true; otherwise, clears the given option.

See also options and testOption().
*/
func (this *QColorDialog) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog9setOptionENS_17ColorDialogOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QColorDialog::ColorDialogOption, bool)

/*
Sets the given option to be enabled if on is true; otherwise, clears the given option.

See also options and testOption().
*/
func (this *QColorDialog) SetOption__(option int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog9setOptionENS_17ColorDialogOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(QColorDialog::ColorDialogOption) const

/*
Returns true if the given option is enabled; otherwise, returns false.

This function was introduced in  Qt 4.5.

See also options and setOption().
*/
func (this *QColorDialog) TestOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QColorDialog10testOptionENS_17ColorDialogOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcolordialog.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QColorDialog::ColorDialogOptions)

/*

 */
func (this *QColorDialog) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog10setOptionsE6QFlagsINS_17ColorDialogOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QColorDialog::ColorDialogOptions options() const

/*

 */
func (this *QColorDialog) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QColorDialog7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open(QObject *, const char *)

/*
This is an overloaded function.

Opens the dialog and connects its colorSelected() signal to the slot specified by receiver and member.

The signal will be disconnected from the slot when the dialog is closed.

This function was introduced in  Qt 4.5.
*/
func (this *QColorDialog) Open(receiver qtcore.QObject_ITF /*777 QObject **/, member string) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog4openEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
Reimplemented from QWidget::setVisible().

Changes the visibility of the dialog. If visible is true, the dialog is shown; otherwise, it is hidden.
*/
func (this *QColorDialog) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:90
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor getColor(const QColor &, QWidget *, const QString &, QColorDialog::ColorDialogOptions)

/*
Pops up a modal color dialog with the given window title (or "Select Color" if none is specified), lets the user choose a color, and returns that color. The color is initially set to initial. The dialog is a child of parent. It returns an invalid (see QColor::isValid()) color if the user cancels the dialog.

The options argument allows you to customize the dialog.

This function was introduced in  Qt 4.5.
*/
func (this *QColorDialog) GetColor(initial qtgui.QColor_ITF, parent QWidget_ITF /*777 QWidget **/, title string, options int) *qtgui.QColor /*123*/ {
	var convArg0 unsafe.Pointer
	if initial != nil && initial.QColor_PTR() != nil {
		convArg0 = initial.QColor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(title)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog8getColorERK6QColorP7QWidgetRK7QString6QFlagsINS_17ColorDialogOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}
func QColorDialog_GetColor(initial qtgui.QColor_ITF, parent QWidget_ITF /*777 QWidget **/, title string, options int) *qtgui.QColor /*123*/ {
	var nilthis *QColorDialog
	rv := nilthis.GetColor(initial, parent, title, options)
	return rv
}

// /usr/include/qt/QtWidgets/qcolordialog.h:90
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor getColor(const QColor &, QWidget *, const QString &, QColorDialog::ColorDialogOptions)

/*
Pops up a modal color dialog with the given window title (or "Select Color" if none is specified), lets the user choose a color, and returns that color. The color is initially set to initial. The dialog is a child of parent. It returns an invalid (see QColor::isValid()) color if the user cancels the dialog.

The options argument allows you to customize the dialog.

This function was introduced in  Qt 4.5.
*/
func (this *QColorDialog) GetColor__() *qtgui.QColor /*123*/ {
	// arg: 0, const QColor &=LValueReference, QColor=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, QColorDialog::ColorDialogOptions=Typedef, QColorDialog::ColorDialogOptions=Typedef, QFlags<QColorDialog::ColorDialogOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog8getColorERK6QColorP7QWidgetRK7QString6QFlagsINS_17ColorDialogOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qcolordialog.h:90
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor getColor(const QColor &, QWidget *, const QString &, QColorDialog::ColorDialogOptions)

/*
Pops up a modal color dialog with the given window title (or "Select Color" if none is specified), lets the user choose a color, and returns that color. The color is initially set to initial. The dialog is a child of parent. It returns an invalid (see QColor::isValid()) color if the user cancels the dialog.

The options argument allows you to customize the dialog.

This function was introduced in  Qt 4.5.
*/
func (this *QColorDialog) GetColor__1(initial qtgui.QColor_ITF) *qtgui.QColor /*123*/ {
	var convArg0 unsafe.Pointer
	if initial != nil && initial.QColor_PTR() != nil {
		convArg0 = initial.QColor_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, QColorDialog::ColorDialogOptions=Typedef, QColorDialog::ColorDialogOptions=Typedef, QFlags<QColorDialog::ColorDialogOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog8getColorERK6QColorP7QWidgetRK7QString6QFlagsINS_17ColorDialogOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qcolordialog.h:90
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor getColor(const QColor &, QWidget *, const QString &, QColorDialog::ColorDialogOptions)

/*
Pops up a modal color dialog with the given window title (or "Select Color" if none is specified), lets the user choose a color, and returns that color. The color is initially set to initial. The dialog is a child of parent. It returns an invalid (see QColor::isValid()) color if the user cancels the dialog.

The options argument allows you to customize the dialog.

This function was introduced in  Qt 4.5.
*/
func (this *QColorDialog) GetColor__2(initial qtgui.QColor_ITF, parent QWidget_ITF /*777 QWidget **/) *qtgui.QColor /*123*/ {
	var convArg0 unsafe.Pointer
	if initial != nil && initial.QColor_PTR() != nil {
		convArg0 = initial.QColor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, QColorDialog::ColorDialogOptions=Typedef, QColorDialog::ColorDialogOptions=Typedef, QFlags<QColorDialog::ColorDialogOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog8getColorERK6QColorP7QWidgetRK7QString6QFlagsINS_17ColorDialogOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qcolordialog.h:90
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor getColor(const QColor &, QWidget *, const QString &, QColorDialog::ColorDialogOptions)

/*
Pops up a modal color dialog with the given window title (or "Select Color" if none is specified), lets the user choose a color, and returns that color. The color is initially set to initial. The dialog is a child of parent. It returns an invalid (see QColor::isValid()) color if the user cancels the dialog.

The options argument allows you to customize the dialog.

This function was introduced in  Qt 4.5.
*/
func (this *QColorDialog) GetColor__3(initial qtgui.QColor_ITF, parent QWidget_ITF /*777 QWidget **/, title string) *qtgui.QColor /*123*/ {
	var convArg0 unsafe.Pointer
	if initial != nil && initial.QColor_PTR() != nil {
		convArg0 = initial.QColor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(title)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QColorDialog::ColorDialogOptions=Typedef, QColorDialog::ColorDialogOptions=Typedef, QFlags<QColorDialog::ColorDialogOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog8getColorERK6QColorP7QWidgetRK7QString6QFlagsINS_17ColorDialogOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qcolordialog.h:96
// index:0
// Public static Visibility=Default Availability=Available
// [4] QRgb getRgba(QRgb, bool *, QWidget *)

/*

 */
func (this *QColorDialog) GetRgba(rgba uint, ok *bool, parent QWidget_ITF /*777 QWidget **/) uint {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg2 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog7getRgbaEjPbP7QWidget", qtrt.FFI_TYPE_POINTER, rgba, ok, convArg2)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}
func QColorDialog_GetRgba(rgba uint, ok *bool, parent QWidget_ITF /*777 QWidget **/) uint {
	var nilthis *QColorDialog
	rv := nilthis.GetRgba(rgba, ok, parent)
	return rv
}

// /usr/include/qt/QtWidgets/qcolordialog.h:96
// index:0
// Public static Visibility=Default Availability=Available
// [4] QRgb getRgba(QRgb, bool *, QWidget *)

/*

 */
func (this *QColorDialog) GetRgba__() uint {
	// arg: 0, QRgb=Typedef, QRgb=Typedef, unsigned int, UInt
	rgba := uint(0xffffffff)
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog7getRgbaEjPbP7QWidget", qtrt.FFI_TYPE_POINTER, rgba, ok, convArg2)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtWidgets/qcolordialog.h:96
// index:0
// Public static Visibility=Default Availability=Available
// [4] QRgb getRgba(QRgb, bool *, QWidget *)

/*

 */
func (this *QColorDialog) GetRgba__1(rgba uint) uint {
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog7getRgbaEjPbP7QWidget", qtrt.FFI_TYPE_POINTER, rgba, ok, convArg2)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtWidgets/qcolordialog.h:96
// index:0
// Public static Visibility=Default Availability=Available
// [4] QRgb getRgba(QRgb, bool *, QWidget *)

/*

 */
func (this *QColorDialog) GetRgba__2(rgba uint, ok *bool) uint {
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog7getRgbaEjPbP7QWidget", qtrt.FFI_TYPE_POINTER, rgba, ok, convArg2)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtWidgets/qcolordialog.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [4] int customCount()

/*
Returns the number of custom colors supported by QColorDialog. All color dialogs share the same custom colors.
*/
func (this *QColorDialog) CustomCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog11customCountEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QColorDialog_CustomCount() int {
	var nilthis *QColorDialog
	rv := nilthis.CustomCount()
	return rv
}

// /usr/include/qt/QtWidgets/qcolordialog.h:99
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor customColor(int)

/*
Returns the custom color at the given index as a QColor value.

This function was introduced in  Qt 4.5.

See also setCustomColor().
*/
func (this *QColorDialog) CustomColor(index int) *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog11customColorEi", qtrt.FFI_TYPE_POINTER, index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}
func QColorDialog_CustomColor(index int) *qtgui.QColor /*123*/ {
	var nilthis *QColorDialog
	rv := nilthis.CustomColor(index)
	return rv
}

// /usr/include/qt/QtWidgets/qcolordialog.h:100
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setCustomColor(int, QColor)

/*
Sets the custom color at index to the QColor color value.

Note: This function does not apply to the Native Color Dialog on the macOS platform. If you still require this function, use the QColorDialog::DontUseNativeDialog option.

See also customColor().
*/
func (this *QColorDialog) SetCustomColor(index int, color qtgui.QColor_ITF /*123*/) {
	var convArg1 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg1 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog14setCustomColorEi6QColor", qtrt.FFI_TYPE_POINTER, index, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QColorDialog_SetCustomColor(index int, color qtgui.QColor_ITF /*123*/) {
	var nilthis *QColorDialog
	nilthis.SetCustomColor(index, color)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:101
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor standardColor(int)

/*
Returns the standard color at the given index as a QColor value.

This function was introduced in  Qt 5.0.

See also setStandardColor().
*/
func (this *QColorDialog) StandardColor(index int) *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog13standardColorEi", qtrt.FFI_TYPE_POINTER, index)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}
func QColorDialog_StandardColor(index int) *qtgui.QColor /*123*/ {
	var nilthis *QColorDialog
	rv := nilthis.StandardColor(index)
	return rv
}

// /usr/include/qt/QtWidgets/qcolordialog.h:102
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setStandardColor(int, QColor)

/*
Sets the standard color at index to the QColor color value.

Note: This function does not apply to the Native Color Dialog on the macOS platform. If you still require this function, use the QColorDialog::DontUseNativeDialog option.

See also standardColor().
*/
func (this *QColorDialog) SetStandardColor(index int, color qtgui.QColor_ITF /*123*/) {
	var convArg1 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg1 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog16setStandardColorEi6QColor", qtrt.FFI_TYPE_POINTER, index, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QColorDialog_SetStandardColor(index int, color qtgui.QColor_ITF /*123*/) {
	var nilthis *QColorDialog
	nilthis.SetStandardColor(index, color)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentColorChanged(const QColor &)

/*
This signal is emitted whenever the current color changes in the dialog. The current color is specified by color.

Note: Notifier signal for property currentColor.

See also color and colorSelected().
*/
func (this *QColorDialog) CurrentColorChanged(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog19currentColorChangedERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void colorSelected(const QColor &)

/*
This signal is emitted just after the user has clicked OK to select a color to use. The chosen color is specified by color.

See also color and currentColorChanged().
*/
func (this *QColorDialog) ColorSelected(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog13colorSelectedERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QColorDialog) ChangeEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:110
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void done(int)

/*
Reimplemented from QDialog::done().

Closes the dialog and sets its result code to result. If this dialog is shown with exec(), done() causes the local event loop to finish, and exec() to return result.

See also QDialog::done().
*/
func (this *QColorDialog) Done(result int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QColorDialog4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), result)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QColorDialog__ColorDialogOption = int

//
const QColorDialog__ShowAlphaChannel QColorDialog__ColorDialogOption = 1

//
const QColorDialog__NoButtons QColorDialog__ColorDialogOption = 2

//
const QColorDialog__DontUseNativeDialog QColorDialog__ColorDialogOption = 4

func (this *QColorDialog) ColorDialogOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QColorDialog_ColorDialogOptionItemName(val int) string {
	var nilthis *QColorDialog
	return nilthis.ColorDialogOptionItemName(val)
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
