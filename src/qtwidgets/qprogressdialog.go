//  header block begin
// /usr/include/qt/QtWidgets/qprogressdialog.h
// #include <qprogressdialog.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
type QProgressDialog struct {
	*QDialog
}

func (this *QProgressDialog) GetCthis() unsafe.Pointer {
	return this.QDialog.GetCthis()
}
func NewQProgressDialogFromPointer(cthis unsafe.Pointer) *QProgressDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QProgressDialog{bcthis0}
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:59
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QProgressDialog) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:75
// index:0
// Public virtual
// void ~QProgressDialog()
func DeleteQProgressDialog(*QProgressDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:77
// index:0
// Public
// void setLabel(class QLabel *)
func (this *QProgressDialog) SetLabel(label unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setLabelEP6QLabel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:78
// index:0
// Public
// void setCancelButton(class QPushButton *)
func (this *QProgressDialog) SetCancelButton(button unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog15setCancelButtonEP11QPushButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:79
// index:0
// Public
// void setBar(class QProgressBar *)
func (this *QProgressDialog) SetBar(bar unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog6setBarEP12QProgressBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), bar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:81
// index:0
// Public
// bool wasCanceled()
func (this *QProgressDialog) WasCanceled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog11wasCanceledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:83
// index:0
// Public
// int minimum()
func (this *QProgressDialog) Minimum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog7minimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:84
// index:0
// Public
// int maximum()
func (this *QProgressDialog) Maximum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog7maximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:86
// index:0
// Public
// int value()
func (this *QProgressDialog) Value() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:88
// index:0
// Public virtual
// QSize sizeHint()
func (this *QProgressDialog) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:90
// index:0
// Public
// QString labelText()
func (this *QProgressDialog) LabelText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9labelTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:91
// index:0
// Public
// int minimumDuration()
func (this *QProgressDialog) MinimumDuration() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog15minimumDurationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:93
// index:0
// Public
// void setAutoReset(_Bool)
func (this *QProgressDialog) SetAutoReset(reset bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoResetEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &reset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:94
// index:0
// Public
// bool autoReset()
func (this *QProgressDialog) AutoReset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9autoResetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:95
// index:0
// Public
// void setAutoClose(_Bool)
func (this *QProgressDialog) SetAutoClose(close bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoCloseEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &close)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:96
// index:0
// Public
// bool autoClose()
func (this *QProgressDialog) AutoClose() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9autoCloseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:99
// index:0
// Public
// void open(class QObject *, const char *)
func (this *QProgressDialog) Open(receiver unsafe.Pointer, member string) {
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), receiver, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:102
// index:0
// Public
// void cancel()
func (this *QProgressDialog) Cancel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog6cancelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:103
// index:0
// Public
// void reset()
func (this *QProgressDialog) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:104
// index:0
// Public
// void setMaximum(int)
func (this *QProgressDialog) SetMaximum(maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10setMaximumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:105
// index:0
// Public
// void setMinimum(int)
func (this *QProgressDialog) SetMinimum(minimum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10setMinimumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &minimum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:106
// index:0
// Public
// void setRange(int, int)
func (this *QProgressDialog) SetRange(minimum int, maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setRangeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &minimum, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:107
// index:0
// Public
// void setValue(int)
func (this *QProgressDialog) SetValue(progress int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setValueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &progress)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:108
// index:0
// Public
// void setLabelText(const class QString &)
func (this *QProgressDialog) SetLabelText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setLabelTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:109
// index:0
// Public
// void setCancelButtonText(const class QString &)
func (this *QProgressDialog) SetCancelButtonText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog19setCancelButtonTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:110
// index:0
// Public
// void setMinimumDuration(int)
func (this *QProgressDialog) SetMinimumDuration(ms int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog18setMinimumDurationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ms)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:113
// index:0
// Public
// void canceled()
func (this *QProgressDialog) Canceled() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8canceledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:116
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QProgressDialog) ResizeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:117
// index:0
// Protected virtual
// void closeEvent(class QCloseEvent *)
func (this *QProgressDialog) CloseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:118
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QProgressDialog) ChangeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:119
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QProgressDialog) ShowEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:122
// index:0
// Protected
// void forceShow()
func (this *QProgressDialog) ForceShow() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog9forceShowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
