package qtwidgets

// /usr/include/qt/QtWidgets/qfontdialog.h
// #include <qfontdialog.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
func (this *QFontDialog) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void done(int)
func (this *QFontDialog) InheritDone(f func(result int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "done", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QFontDialog) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

/*

 */
type QFontDialog struct {
	*QDialog
}
type QFontDialog_ITF interface {
	QDialog_ITF
	QFontDialog_PTR() *QFontDialog
}

func (ptr *QFontDialog) QFontDialog_PTR() *QFontDialog { return ptr }

func (this *QFontDialog) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QFontDialog) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQFontDialogFromPointer(cthis unsafe.Pointer) *QFontDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QFontDialog{bcthis0}
}
func (*QFontDialog) NewFromPointer(cthis unsafe.Pointer) *QFontDialog {
	return NewQFontDialogFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QFontDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFontDialog10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfontdialog.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontDialog(QWidget *)

/*
Constructs a standard font dialog.

Use setCurrentFont() to set the initial font attributes.

The parent parameter is passed to the QDialog constructor.

This function was introduced in  Qt 4.5.

See also getFont().
*/
func (*QFontDialog) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QFontDialog {
	return NewQFontDialog(parent)
}
func NewQFontDialog(parent QWidget_ITF /*777 QWidget **/) *QFontDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialogC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFontDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qfontdialog.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontDialog(QWidget *)

/*
Constructs a standard font dialog.

Use setCurrentFont() to set the initial font attributes.

The parent parameter is passed to the QDialog constructor.

This function was introduced in  Qt 4.5.

See also getFont().
*/
func (*QFontDialog) NewForInherit__() *QFontDialog {
	return NewQFontDialog__()
}
func NewQFontDialog__() *QFontDialog {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialogC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFontDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qfontdialog.h:76
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFontDialog(const QFont &, QWidget *)

/*
Constructs a standard font dialog.

Use setCurrentFont() to set the initial font attributes.

The parent parameter is passed to the QDialog constructor.

This function was introduced in  Qt 4.5.

See also getFont().
*/
func (*QFontDialog) NewForInherit_1(initial qtgui.QFont_ITF, parent QWidget_ITF /*777 QWidget **/) *QFontDialog {
	return NewQFontDialog_1(initial, parent)
}
func NewQFontDialog_1(initial qtgui.QFont_ITF, parent QWidget_ITF /*777 QWidget **/) *QFontDialog {
	var convArg0 unsafe.Pointer
	if initial != nil && initial.QFont_PTR() != nil {
		convArg0 = initial.QFont_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialogC2ERK5QFontP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFontDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qfontdialog.h:76
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFontDialog(const QFont &, QWidget *)

/*
Constructs a standard font dialog.

Use setCurrentFont() to set the initial font attributes.

The parent parameter is passed to the QDialog constructor.

This function was introduced in  Qt 4.5.

See also getFont().
*/
func (*QFontDialog) NewForInherit_1_(initial qtgui.QFont_ITF) *QFontDialog {
	return NewQFontDialog_1_(initial)
}
func NewQFontDialog_1_(initial qtgui.QFont_ITF) *QFontDialog {
	var convArg0 unsafe.Pointer
	if initial != nil && initial.QFont_PTR() != nil {
		convArg0 = initial.QFont_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialogC2ERK5QFontP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFontDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qfontdialog.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFontDialog()

/*

 */
func DeleteQFontDialog(this *QFontDialog) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialogD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentFont(const QFont &)

/*
Sets the font highlighted in the QFontDialog to the given font.

This function was introduced in  Qt 4.5.

Note: Setter function for property currentFont.

See also currentFont() and selectedFont().
*/
func (this *QFontDialog) SetCurrentFont(font qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog14setCurrentFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:80
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont currentFont() const

/*
Returns the current font.

This function was introduced in  Qt 4.5.

Note: Getter function for property currentFont.

See also setCurrentFont() and selectedFont().
*/
func (this *QFontDialog) CurrentFont() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFontDialog11currentFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontdialog.h:82
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont selectedFont() const

/*
Returns the font that the user selected by clicking the OK or equivalent button.

Note: This font is not always the same as the font held by the currentFont property since the user can choose different fonts before finally selecting the one to use.
*/
func (this *QFontDialog) SelectedFont() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFontDialog12selectedFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontdialog.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QFontDialog::FontDialogOption, bool)

/*
Sets the given option to be enabled if on is true; otherwise, clears the given option.

See also options and testOption().
*/
func (this *QFontDialog) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog9setOptionENS_16FontDialogOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QFontDialog::FontDialogOption, bool)

/*
Sets the given option to be enabled if on is true; otherwise, clears the given option.

See also options and testOption().
*/
func (this *QFontDialog) SetOption__(option int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog9setOptionENS_16FontDialogOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(QFontDialog::FontDialogOption) const

/*
Returns true if the given option is enabled; otherwise, returns false.

See also options and setOption().
*/
func (this *QFontDialog) TestOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFontDialog10testOptionENS_16FontDialogOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfontdialog.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QFontDialog::FontDialogOptions)

/*

 */
func (this *QFontDialog) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog10setOptionsE6QFlagsINS_16FontDialogOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] QFontDialog::FontDialogOptions options() const

/*

 */
func (this *QFontDialog) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFontDialog7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open(QObject *, const char *)

/*
This is an overloaded function.

Opens the dialog and connects its fontSelected() signal to the slot specified by receiver and member.

The signal will be disconnected from the slot when the dialog is closed.

This function was introduced in  Qt 4.5.
*/
func (this *QFontDialog) Open(receiver qtcore.QObject_ITF /*777 QObject **/, member string) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog4openEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
Reimplemented from QWidget::setVisible().
*/
func (this *QFontDialog) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [16] QFont getFont(bool *, QWidget *)

/*
Executes a modal font dialog and returns a font.

If the user clicks OK, the selected font is returned. If the user clicks Cancel, the initial font is returned.

The dialog is constructed with the given parent and the options specified in options. title is shown as the window title of the dialog and initial is the initially selected font. If the ok parameter is not-null, the value it refers to is set to true if the user clicks OK, and set to false if the user clicks Cancel.

Examples:


  bool ok;
  QFont font = QFontDialog::getFont(&ok, QFont("Times", 12), this);
  if (ok) {
      // font is set to the font the user selected
  } else {
      // the user canceled the dialog; font is set to the initial
      // value, in this case Times, 12.
  }



The dialog can also be used to set a widget's font directly:


  myWidget.setFont(QFontDialog::getFont(0, myWidget.font()));



In this example, if the user clicks OK the font they chose will be used, and if they click Cancel the original font is used.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFontDialog constructors.
*/
func (this *QFontDialog) GetFont(ok *bool, parent QWidget_ITF /*777 QWidget **/) *qtgui.QFont /*123*/ {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbP7QWidget", qtrt.FFI_TYPE_POINTER, ok, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QFontDialog_GetFont(ok *bool, parent QWidget_ITF /*777 QWidget **/) *qtgui.QFont /*123*/ {
	var nilthis *QFontDialog
	rv := nilthis.GetFont(ok, parent)
	return rv
}

// /usr/include/qt/QtWidgets/qfontdialog.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [16] QFont getFont(bool *, QWidget *)

/*
Executes a modal font dialog and returns a font.

If the user clicks OK, the selected font is returned. If the user clicks Cancel, the initial font is returned.

The dialog is constructed with the given parent and the options specified in options. title is shown as the window title of the dialog and initial is the initially selected font. If the ok parameter is not-null, the value it refers to is set to true if the user clicks OK, and set to false if the user clicks Cancel.

Examples:


  bool ok;
  QFont font = QFontDialog::getFont(&ok, QFont("Times", 12), this);
  if (ok) {
      // font is set to the font the user selected
  } else {
      // the user canceled the dialog; font is set to the initial
      // value, in this case Times, 12.
  }



The dialog can also be used to set a widget's font directly:


  myWidget.setFont(QFontDialog::getFont(0, myWidget.font()));



In this example, if the user clicks OK the font they chose will be used, and if they click Cancel the original font is used.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFontDialog constructors.
*/
func (this *QFontDialog) GetFont__(ok *bool) *qtgui.QFont /*123*/ {
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbP7QWidget", qtrt.FFI_TYPE_POINTER, ok, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontdialog.h:95
// index:1
// Public static Visibility=Default Availability=Available
// [16] QFont getFont(bool *, const QFont &, QWidget *, const QString &, QFontDialog::FontDialogOptions)

/*
Executes a modal font dialog and returns a font.

If the user clicks OK, the selected font is returned. If the user clicks Cancel, the initial font is returned.

The dialog is constructed with the given parent and the options specified in options. title is shown as the window title of the dialog and initial is the initially selected font. If the ok parameter is not-null, the value it refers to is set to true if the user clicks OK, and set to false if the user clicks Cancel.

Examples:


  bool ok;
  QFont font = QFontDialog::getFont(&ok, QFont("Times", 12), this);
  if (ok) {
      // font is set to the font the user selected
  } else {
      // the user canceled the dialog; font is set to the initial
      // value, in this case Times, 12.
  }



The dialog can also be used to set a widget's font directly:


  myWidget.setFont(QFontDialog::getFont(0, myWidget.font()));



In this example, if the user clicks OK the font they chose will be used, and if they click Cancel the original font is used.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFontDialog constructors.
*/
func (this *QFontDialog) GetFont_1(ok *bool, initial qtgui.QFont_ITF, parent QWidget_ITF /*777 QWidget **/, title string, options int) *qtgui.QFont /*123*/ {
	var convArg1 unsafe.Pointer
	if initial != nil && initial.QFont_PTR() != nil {
		convArg1 = initial.QFont_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg2 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString_5(title)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbRK5QFontP7QWidgetRK7QString6QFlagsINS_16FontDialogOptionEE", qtrt.FFI_TYPE_POINTER, ok, convArg1, convArg2, convArg3, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QFontDialog_GetFont_1(ok *bool, initial qtgui.QFont_ITF, parent QWidget_ITF /*777 QWidget **/, title string, options int) *qtgui.QFont /*123*/ {
	var nilthis *QFontDialog
	rv := nilthis.GetFont_1(ok, initial, parent, title, options)
	return rv
}

// /usr/include/qt/QtWidgets/qfontdialog.h:95
// index:1
// Public static Visibility=Default Availability=Available
// [16] QFont getFont(bool *, const QFont &, QWidget *, const QString &, QFontDialog::FontDialogOptions)

/*
Executes a modal font dialog and returns a font.

If the user clicks OK, the selected font is returned. If the user clicks Cancel, the initial font is returned.

The dialog is constructed with the given parent and the options specified in options. title is shown as the window title of the dialog and initial is the initially selected font. If the ok parameter is not-null, the value it refers to is set to true if the user clicks OK, and set to false if the user clicks Cancel.

Examples:


  bool ok;
  QFont font = QFontDialog::getFont(&ok, QFont("Times", 12), this);
  if (ok) {
      // font is set to the font the user selected
  } else {
      // the user canceled the dialog; font is set to the initial
      // value, in this case Times, 12.
  }



The dialog can also be used to set a widget's font directly:


  myWidget.setFont(QFontDialog::getFont(0, myWidget.font()));



In this example, if the user clicks OK the font they chose will be used, and if they click Cancel the original font is used.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFontDialog constructors.
*/
func (this *QFontDialog) GetFont_1_(ok *bool, initial qtgui.QFont_ITF) *qtgui.QFont /*123*/ {
	var convArg1 unsafe.Pointer
	if initial != nil && initial.QFont_PTR() != nil {
		convArg1 = initial.QFont_PTR().GetCthis()
	}
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QFontDialog::FontDialogOptions=Typedef, QFontDialog::FontDialogOptions=Typedef, QFlags<QFontDialog::FontDialogOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbRK5QFontP7QWidgetRK7QString6QFlagsINS_16FontDialogOptionEE", qtrt.FFI_TYPE_POINTER, ok, convArg1, convArg2, convArg3, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontdialog.h:95
// index:1
// Public static Visibility=Default Availability=Available
// [16] QFont getFont(bool *, const QFont &, QWidget *, const QString &, QFontDialog::FontDialogOptions)

/*
Executes a modal font dialog and returns a font.

If the user clicks OK, the selected font is returned. If the user clicks Cancel, the initial font is returned.

The dialog is constructed with the given parent and the options specified in options. title is shown as the window title of the dialog and initial is the initially selected font. If the ok parameter is not-null, the value it refers to is set to true if the user clicks OK, and set to false if the user clicks Cancel.

Examples:


  bool ok;
  QFont font = QFontDialog::getFont(&ok, QFont("Times", 12), this);
  if (ok) {
      // font is set to the font the user selected
  } else {
      // the user canceled the dialog; font is set to the initial
      // value, in this case Times, 12.
  }



The dialog can also be used to set a widget's font directly:


  myWidget.setFont(QFontDialog::getFont(0, myWidget.font()));



In this example, if the user clicks OK the font they chose will be used, and if they click Cancel the original font is used.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFontDialog constructors.
*/
func (this *QFontDialog) GetFont_1_1(ok *bool, initial qtgui.QFont_ITF, parent QWidget_ITF /*777 QWidget **/) *qtgui.QFont /*123*/ {
	var convArg1 unsafe.Pointer
	if initial != nil && initial.QFont_PTR() != nil {
		convArg1 = initial.QFont_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg2 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QFontDialog::FontDialogOptions=Typedef, QFontDialog::FontDialogOptions=Typedef, QFlags<QFontDialog::FontDialogOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbRK5QFontP7QWidgetRK7QString6QFlagsINS_16FontDialogOptionEE", qtrt.FFI_TYPE_POINTER, ok, convArg1, convArg2, convArg3, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontdialog.h:95
// index:1
// Public static Visibility=Default Availability=Available
// [16] QFont getFont(bool *, const QFont &, QWidget *, const QString &, QFontDialog::FontDialogOptions)

/*
Executes a modal font dialog and returns a font.

If the user clicks OK, the selected font is returned. If the user clicks Cancel, the initial font is returned.

The dialog is constructed with the given parent and the options specified in options. title is shown as the window title of the dialog and initial is the initially selected font. If the ok parameter is not-null, the value it refers to is set to true if the user clicks OK, and set to false if the user clicks Cancel.

Examples:


  bool ok;
  QFont font = QFontDialog::getFont(&ok, QFont("Times", 12), this);
  if (ok) {
      // font is set to the font the user selected
  } else {
      // the user canceled the dialog; font is set to the initial
      // value, in this case Times, 12.
  }



The dialog can also be used to set a widget's font directly:


  myWidget.setFont(QFontDialog::getFont(0, myWidget.font()));



In this example, if the user clicks OK the font they chose will be used, and if they click Cancel the original font is used.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFontDialog constructors.
*/
func (this *QFontDialog) GetFont_1_2(ok *bool, initial qtgui.QFont_ITF, parent QWidget_ITF /*777 QWidget **/, title string) *qtgui.QFont /*123*/ {
	var convArg1 unsafe.Pointer
	if initial != nil && initial.QFont_PTR() != nil {
		convArg1 = initial.QFont_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg2 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString_5(title)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 4, QFontDialog::FontDialogOptions=Typedef, QFontDialog::FontDialogOptions=Typedef, QFlags<QFontDialog::FontDialogOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbRK5QFontP7QWidgetRK7QString6QFlagsINS_16FontDialogOptionEE", qtrt.FFI_TYPE_POINTER, ok, convArg1, convArg2, convArg3, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontdialog.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentFontChanged(const QFont &)

/*
This signal is emitted when the current font is changed. The new font is specified in font.

The signal is emitted while a user is selecting a font. Ultimately, the chosen font may differ from the font currently selected.

This function was introduced in  Qt 4.5.

Note: Notifier signal for property currentFont.

See also currentFont, fontSelected(), and selectedFont().
*/
func (this *QFontDialog) CurrentFontChanged(font qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog18currentFontChangedERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fontSelected(const QFont &)

/*
This signal is emitted when a font has been selected. The selected font is specified in font.

The signal is only emitted when a user has chosen the final font to be used. It is not emitted while the user is changing the current font in the font dialog.

This function was introduced in  Qt 4.5.

See also selectedFont(), currentFontChanged(), and currentFont.
*/
func (this *QFontDialog) FontSelected(font qtgui.QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog12fontSelectedERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:103
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QFontDialog) ChangeEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:104
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void done(int)

/*
Reimplemented from QDialog::done().

Closes the dialog and sets its result code to result. If this dialog is shown with exec(), done() causes the local event loop to finish, and exec() to return result.

See also QDialog::done().
*/
func (this *QFontDialog) Done(result int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), result)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:105
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*

 */
func (this *QFontDialog) EventFilter(object qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QFontDialog__FontDialogOption = int

//
const QFontDialog__NoButtons QFontDialog__FontDialogOption = 1

//
const QFontDialog__DontUseNativeDialog QFontDialog__FontDialogOption = 2

//
const QFontDialog__ScalableFonts QFontDialog__FontDialogOption = 4

//
const QFontDialog__NonScalableFonts QFontDialog__FontDialogOption = 8

//
const QFontDialog__MonospacedFonts QFontDialog__FontDialogOption = 16

//
const QFontDialog__ProportionalFonts QFontDialog__FontDialogOption = 32

func (this *QFontDialog) FontDialogOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFontDialog_FontDialogOptionItemName(val int) string {
	var nilthis *QFontDialog
	return nilthis.FontDialogOptionItemName(val)
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
