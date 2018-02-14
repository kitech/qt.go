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

// void resizeEvent(class QResizeEvent *)
func (this *QProgressDialog) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void closeEvent(class QCloseEvent *)
func (this *QProgressDialog) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QProgressDialog) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QProgressDialog) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void forceShow()
func (this *QProgressDialog) InheritForceShow(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "forceShow", f)
}

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
// [8] const QMetaObject * metaObject()
func (this *QProgressDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProgressDialog(QWidget *, Qt::WindowFlags)
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

// /usr/include/qt/QtWidgets/qprogressdialog.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QProgressDialog(const QString &, const QString &, int, int, QWidget *, Qt::WindowFlags)
func NewQProgressDialog_1(labelText string, cancelButtonText string, minimum int, maximum int, parent QWidget_ITF /*777 QWidget **/, flags int) *QProgressDialog {
	var tmpArg0 = qtcore.NewQString_5(labelText)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(cancelButtonText)
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

// /usr/include/qt/QtWidgets/qprogressdialog.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QProgressDialog()
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
// [1] bool wasCanceled()
func (this *QProgressDialog) WasCanceled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog11wasCanceledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimum()
func (this *QProgressDialog) Minimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog7minimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximum()
func (this *QProgressDialog) Maximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog7maximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int value()
func (this *QProgressDialog) Value() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
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
// [8] QString labelText()
func (this *QProgressDialog) LabelText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog9labelTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimumDuration()
func (this *QProgressDialog) MinimumDuration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog15minimumDurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoReset(_Bool)
func (this *QProgressDialog) SetAutoReset(reset bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoResetEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), reset)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoReset()
func (this *QProgressDialog) AutoReset() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog9autoResetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoClose(_Bool)
func (this *QProgressDialog) SetAutoClose(close bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoCloseEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), close)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoClose()
func (this *QProgressDialog) AutoClose() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QProgressDialog9autoCloseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open(QObject *, const char *)
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
func (this *QProgressDialog) Cancel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog6cancelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()
func (this *QProgressDialog) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximum(int)
func (this *QProgressDialog) SetMaximum(maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog10setMaximumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimum(int)
func (this *QProgressDialog) SetMinimum(minimum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog10setMinimumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRange(int, int)
func (this *QProgressDialog) SetRange(minimum int, maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog8setRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum, maximum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(int)
func (this *QProgressDialog) SetValue(progress int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog8setValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), progress)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLabelText(const QString &)
func (this *QProgressDialog) SetLabelText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog12setLabelTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCancelButtonText(const QString &)
func (this *QProgressDialog) SetCancelButtonText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog19setCancelButtonTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumDuration(int)
func (this *QProgressDialog) SetMinimumDuration(ms int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog18setMinimumDurationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ms)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canceled()
func (this *QProgressDialog) Canceled() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QProgressDialog8canceledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
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
