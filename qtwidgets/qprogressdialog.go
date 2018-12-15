// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qprogressdialog.h
// #include <qprogressdialog.h>
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

// void resizeEvent(QResizeEvent *)
func (this *QProgressDialog) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void closeEvent(QCloseEvent *)
func (this *QProgressDialog) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void changeEvent(QEvent *)
func (this *QProgressDialog) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QProgressDialog) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void forceShow()
func (this *QProgressDialog) InheritForceShow(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "forceShow", f)
}

/*

 */
type QProgressDialog struct {
	*QDialog
}
type QProgressDialog_ITF interface {
	QDialog_ITF
	QProgressDialog_PTR() *QProgressDialog
}

func (ptr *QProgressDialog) QProgressDialog_PTR() *QProgressDialog { return ptr }

func (this *QProgressDialog) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QProgressDialog) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQProgressDialogFromPointer(cthis unsafe.Pointer) *QProgressDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QProgressDialog{bcthis0}
}
func (*QProgressDialog) NewFromPointer(cthis unsafe.Pointer) *QProgressDialog {
	return NewQProgressDialogFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QProgressDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProgressDialog(QWidget *, Qt::WindowFlags)

/*
Constructs a progress dialog.

Default settings:


The label text is empty.
The cancel button text is (translated) "Cancel".
minimum is 0;
maximum is 100


The parent argument is dialog's parent widget. The widget flags, f, are passed to the QDialog::QDialog() constructor.

See also setLabelText(), setCancelButtonText(), setCancelButton(), setMinimum(), and setMaximum().
*/
func (*QProgressDialog) NewForInherit(parent QWidget_ITF /*777 QWidget **/, flags int) *QProgressDialog {
	return NewQProgressDialog(parent, flags)
}
func NewQProgressDialog(parent QWidget_ITF /*777 QWidget **/, flags int) *QProgressDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProgressDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProgressDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProgressDialog(QWidget *, Qt::WindowFlags)

/*
Constructs a progress dialog.

Default settings:


The label text is empty.
The cancel button text is (translated) "Cancel".
minimum is 0;
maximum is 100


The parent argument is dialog's parent widget. The widget flags, f, are passed to the QDialog::QDialog() constructor.

See also setLabelText(), setCancelButtonText(), setCancelButton(), setMinimum(), and setMaximum().
*/
func (*QProgressDialog) NewForInheritp() *QProgressDialog {
	return NewQProgressDialogp()
}
func NewQProgressDialogp() *QProgressDialog {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProgressDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProgressDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProgressDialog(QWidget *, Qt::WindowFlags)

/*
Constructs a progress dialog.

Default settings:


The label text is empty.
The cancel button text is (translated) "Cancel".
minimum is 0;
maximum is 100


The parent argument is dialog's parent widget. The widget flags, f, are passed to the QDialog::QDialog() constructor.

See also setLabelText(), setCancelButtonText(), setCancelButton(), setMinimum(), and setMaximum().
*/
func (*QProgressDialog) NewForInheritp1(parent QWidget_ITF /*777 QWidget **/) *QProgressDialog {
	return NewQProgressDialogp1(parent)
}
func NewQProgressDialogp1(parent QWidget_ITF /*777 QWidget **/) *QProgressDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProgressDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProgressDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QProgressDialog(const QString &, const QString &, int, int, QWidget *, Qt::WindowFlags)

/*
Constructs a progress dialog.

Default settings:


The label text is empty.
The cancel button text is (translated) "Cancel".
minimum is 0;
maximum is 100


The parent argument is dialog's parent widget. The widget flags, f, are passed to the QDialog::QDialog() constructor.

See also setLabelText(), setCancelButtonText(), setCancelButton(), setMinimum(), and setMaximum().
*/
func (*QProgressDialog) NewForInherit1(labelText string, cancelButtonText string, minimum int, maximum int, parent QWidget_ITF /*777 QWidget **/, flags int) *QProgressDialog {
	return NewQProgressDialog1(labelText, cancelButtonText, minimum, maximum, parent, flags)
}
func NewQProgressDialog1(labelText string, cancelButtonText string, minimum int, maximum int, parent QWidget_ITF /*777 QWidget **/, flags int) *QProgressDialog {
	var tmpArg0 = qtcore.NewQString5(labelText)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(cancelButtonText)
	var convArg1 = tmpArg1.GetCthis()
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg4 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialogC2ERK7QStringS2_iiP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, minimum, maximum, convArg4, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProgressDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProgressDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QProgressDialog(const QString &, const QString &, int, int, QWidget *, Qt::WindowFlags)

/*
Constructs a progress dialog.

Default settings:


The label text is empty.
The cancel button text is (translated) "Cancel".
minimum is 0;
maximum is 100


The parent argument is dialog's parent widget. The widget flags, f, are passed to the QDialog::QDialog() constructor.

See also setLabelText(), setCancelButtonText(), setCancelButton(), setMinimum(), and setMaximum().
*/
func (*QProgressDialog) NewForInherit1p(labelText string, cancelButtonText string, minimum int, maximum int) *QProgressDialog {
	return NewQProgressDialog1p(labelText, cancelButtonText, minimum, maximum)
}
func NewQProgressDialog1p(labelText string, cancelButtonText string, minimum int, maximum int) *QProgressDialog {
	var tmpArg0 = qtcore.NewQString5(labelText)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(cancelButtonText)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 4, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialogC2ERK7QStringS2_iiP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, minimum, maximum, convArg4, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProgressDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProgressDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QProgressDialog(const QString &, const QString &, int, int, QWidget *, Qt::WindowFlags)

/*
Constructs a progress dialog.

Default settings:


The label text is empty.
The cancel button text is (translated) "Cancel".
minimum is 0;
maximum is 100


The parent argument is dialog's parent widget. The widget flags, f, are passed to the QDialog::QDialog() constructor.

See also setLabelText(), setCancelButtonText(), setCancelButton(), setMinimum(), and setMaximum().
*/
func (*QProgressDialog) NewForInherit1p1(labelText string, cancelButtonText string, minimum int, maximum int, parent QWidget_ITF /*777 QWidget **/) *QProgressDialog {
	return NewQProgressDialog1p1(labelText, cancelButtonText, minimum, maximum, parent)
}
func NewQProgressDialog1p1(labelText string, cancelButtonText string, minimum int, maximum int, parent QWidget_ITF /*777 QWidget **/) *QProgressDialog {
	var tmpArg0 = qtcore.NewQString5(labelText)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(cancelButtonText)
	var convArg1 = tmpArg1.GetCthis()
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg4 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 5, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialogC2ERK7QStringS2_iiP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, minimum, maximum, convArg4, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProgressDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProgressDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QProgressDialog()

/*

 */
func DeleteQProgressDialog(this *QProgressDialog) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialogD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLabel(QLabel *)

/*
Sets the label to label. The progress dialog resizes to fit. The label becomes owned by the progress dialog and will be deleted when necessary, so do not pass the address of an object on the stack.

See also setLabelText().
*/
func (this *QProgressDialog) SetLabel(label QLabel_ITF /*777 QLabel **/) {
	var convArg0 unsafe.Pointer
	if label != nil && label.QLabel_PTR() != nil {
		convArg0 = label.QLabel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog8setLabelEP6QLabel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCancelButton(QPushButton *)

/*
Sets the cancel button to the push button, cancelButton. The progress dialog takes ownership of this button which will be deleted when necessary, so do not pass the address of an object that is on the stack, i.e. use new() to create the button. If 0 is passed then no cancel button will be shown.

See also setCancelButtonText().
*/
func (this *QProgressDialog) SetCancelButton(button QPushButton_ITF /*777 QPushButton **/) {
	var convArg0 unsafe.Pointer
	if button != nil && button.QPushButton_PTR() != nil {
		convArg0 = button.QPushButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog15setCancelButtonEP11QPushButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBar(QProgressBar *)

/*
Sets the progress bar widget to bar. The progress dialog resizes to fit. The progress dialog takes ownership of the progress bar which will be deleted when necessary, so do not use a progress bar allocated on the stack.
*/
func (this *QProgressDialog) SetBar(bar QProgressBar_ITF /*777 QProgressBar **/) {
	var convArg0 unsafe.Pointer
	if bar != nil && bar.QProgressBar_PTR() != nil {
		convArg0 = bar.QProgressBar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog6setBarEP12QProgressBar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wasCanceled() const

/*

 */
func (this *QProgressDialog) WasCanceled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog11wasCanceledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimum() const

/*

 */
func (this *QProgressDialog) Minimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog7minimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximum() const

/*

 */
func (this *QProgressDialog) Maximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog7maximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int value() const

/*

 */
func (this *QProgressDialog) Value() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QDialog::sizeHint().

Returns a size that fits the contents of the progress dialog. The progress dialog resizes itself as required, so you should not need to call this yourself.
*/
func (this *QProgressDialog) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QString labelText() const

/*

 */
func (this *QProgressDialog) LabelText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog9labelTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimumDuration() const

/*

 */
func (this *QProgressDialog) MinimumDuration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog15minimumDurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoReset(bool)

/*

 */
func (this *QProgressDialog) SetAutoReset(reset bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoResetEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), reset)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoReset() const

/*

 */
func (this *QProgressDialog) AutoReset() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog9autoResetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoClose(bool)

/*

 */
func (this *QProgressDialog) SetAutoClose(close bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoCloseEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), close)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoClose() const

/*

 */
func (this *QProgressDialog) AutoClose() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog9autoCloseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open(QObject *, const char *)

/*
Opens the dialog and connects its canceled() signal to the slot specified by receiver and member.

The signal will be disconnected from the slot when the dialog is closed.

This function was introduced in  Qt 4.5.
*/
func (this *QProgressDialog) Open(receiver qtcore.QObject_ITF /*777 QObject **/, member string) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog4openEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancel()

/*
Resets the progress dialog. wasCanceled() becomes true until the progress dialog is reset. The progress dialog becomes hidden.
*/
func (this *QProgressDialog) Cancel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog6cancelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()

/*
Resets the progress dialog. The progress dialog becomes hidden if autoClose() is true.

See also setAutoClose() and setAutoReset().
*/
func (this *QProgressDialog) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximum(int)

/*

 */
func (this *QProgressDialog) SetMaximum(maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog10setMaximumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimum(int)

/*

 */
func (this *QProgressDialog) SetMinimum(minimum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog10setMinimumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRange(int, int)

/*
Sets the progress dialog's minimum and maximum values to minimum and maximum, respectively.

If maximum is smaller than minimum, minimum becomes the only legal value.

If the current value falls outside the new range, the progress dialog is reset with reset().

See also minimum and maximum.
*/
func (this *QProgressDialog) SetRange(minimum int, maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog8setRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum, maximum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(int)

/*

 */
func (this *QProgressDialog) SetValue(progress int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog8setValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), progress)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLabelText(const QString &)

/*

 */
func (this *QProgressDialog) SetLabelText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog12setLabelTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCancelButtonText(const QString &)

/*
Sets the cancel button's text to cancelButtonText. If the text is set to QString() then it will cause the cancel button to be hidden and deleted.

See also setCancelButton().
*/
func (this *QProgressDialog) SetCancelButtonText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog19setCancelButtonTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumDuration(int)

/*

 */
func (this *QProgressDialog) SetMinimumDuration(ms int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog18setMinimumDurationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ms)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canceled()

/*
This signal is emitted when the cancel button is clicked. It is connected to the cancel() slot by default.

See also wasCanceled().
*/
func (this *QProgressDialog) Canceled() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog8canceledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QDialog::resizeEvent().
*/
func (this *QProgressDialog) ResizeEvent(event qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QResizeEvent_PTR() != nil {
		convArg0 = event.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)

/*
Reimplemented from QDialog::closeEvent().
*/
func (this *QProgressDialog) CloseEvent(event qtgui.QCloseEvent_ITF /*777 QCloseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QCloseEvent_PTR() != nil {
		convArg0 = event.QCloseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:118
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QProgressDialog) ChangeEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QDialog::showEvent().
*/
func (this *QProgressDialog) ShowEvent(event qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QShowEvent_PTR() != nil {
		convArg0 = event.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:122
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void forceShow()

/*
Shows the dialog if it is still hidden after the algorithm has been started and minimumDuration milliseconds have passed.

See also setMinimumDuration().
*/
func (this *QProgressDialog) ForceShow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog9forceShowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
