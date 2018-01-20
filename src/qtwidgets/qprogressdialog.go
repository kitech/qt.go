//  header block begin
// /usr/include/qt/QtWidgets/qprogressdialog.h
// #include <qprogressdialog.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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

// /usr/include/qt/QtWidgets/qprogressdialog.h:59
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QProgressDialog) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:71
// index:0
// void QProgressDialog(class QWidget *, Qt::WindowFlags)
func NewQProgressDialog(parent unsafe.Pointer, flags int) *QProgressDialog {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, parent, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQProgressDialogFromPointer(cthis)
	return gothis
}
func NewQProgressDialogFromPointer(cthis unsafe.Pointer) *QProgressDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QProgressDialog{bcthis0}
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:72
// index:1
// void QProgressDialog(const class QString &, const class QString &, int, int, class QWidget *, Qt::WindowFlags)
func NewQProgressDialog_1(labelText unsafe.Pointer, cancelButtonText unsafe.Pointer, minimum int, maximum int, parent unsafe.Pointer, flags int) *QProgressDialog {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialogC2ERK7QStringS2_iiP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, labelText, cancelButtonText, &minimum, &maximum, parent, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQProgressDialogFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:75
// index:0
// virtual
// void ~QProgressDialog()
func DeleteQProgressDialog(*QProgressDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:77
// index:0
// void setLabel(class QLabel *)
func (this *QProgressDialog) SetLabel(label unsafe.Pointer) {
	// 0: (, label QLabel *), (label)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setLabelEP6QLabel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:78
// index:0
// void setCancelButton(class QPushButton *)
func (this *QProgressDialog) SetCancelButton(button unsafe.Pointer) {
	// 0: (, button QPushButton *), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog15setCancelButtonEP11QPushButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:79
// index:0
// void setBar(class QProgressBar *)
func (this *QProgressDialog) SetBar(bar unsafe.Pointer) {
	// 0: (, bar QProgressBar *), (bar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog6setBarEP12QProgressBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), bar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:81
// index:0
// bool wasCanceled()
func (this *QProgressDialog) WasCanceled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog11wasCanceledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:83
// index:0
// int minimum()
func (this *QProgressDialog) Minimum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog7minimumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:84
// index:0
// int maximum()
func (this *QProgressDialog) Maximum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog7maximumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:86
// index:0
// int value()
func (this *QProgressDialog) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog5valueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:88
// index:0
// virtual
// QSize sizeHint()
func (this *QProgressDialog) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:90
// index:0
// QString labelText()
func (this *QProgressDialog) LabelText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9labelTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:91
// index:0
// int minimumDuration()
func (this *QProgressDialog) MinimumDuration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog15minimumDurationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:93
// index:0
// void setAutoReset(_Bool)
func (this *QProgressDialog) SetAutoReset(reset bool) {
	// 0: (, reset bool), (&reset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoResetEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &reset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:94
// index:0
// bool autoReset()
func (this *QProgressDialog) AutoReset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9autoResetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:95
// index:0
// void setAutoClose(_Bool)
func (this *QProgressDialog) SetAutoClose(close bool) {
	// 0: (, close bool), (&close)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoCloseEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &close)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:96
// index:0
// bool autoClose()
func (this *QProgressDialog) AutoClose() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9autoCloseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:99
// index:0
// void open(class QObject *, const char *)
func (this *QProgressDialog) Open(receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (, receiver QObject *, member const char *), (receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:102
// index:0
// void cancel()
func (this *QProgressDialog) Cancel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog6cancelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:103
// index:0
// void reset()
func (this *QProgressDialog) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog5resetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:104
// index:0
// void setMaximum(int)
func (this *QProgressDialog) SetMaximum(maximum int) {
	// 0: (, maximum int), (&maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10setMaximumEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:105
// index:0
// void setMinimum(int)
func (this *QProgressDialog) SetMinimum(minimum int) {
	// 0: (, minimum int), (&minimum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10setMinimumEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &minimum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:106
// index:0
// void setRange(int, int)
func (this *QProgressDialog) SetRange(minimum int, maximum int) {
	// 0: (, minimum int, maximum int), (&minimum, &maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setRangeEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &minimum, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:107
// index:0
// void setValue(int)
func (this *QProgressDialog) SetValue(progress int) {
	// 0: (, progress int), (&progress)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setValueEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &progress)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:108
// index:0
// void setLabelText(const class QString &)
func (this *QProgressDialog) SetLabelText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setLabelTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:109
// index:0
// void setCancelButtonText(const class QString &)
func (this *QProgressDialog) SetCancelButtonText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog19setCancelButtonTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:110
// index:0
// void setMinimumDuration(int)
func (this *QProgressDialog) SetMinimumDuration(ms int) {
	// 0: (, ms int), (&ms)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog18setMinimumDurationEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ms)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:113
// index:0
// void canceled()
func (this *QProgressDialog) Canceled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8canceledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:116
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QProgressDialog) ResizeEvent(event unsafe.Pointer) {
	// 0: (, event QResizeEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:117
// index:0
// virtual
// void closeEvent(class QCloseEvent *)
func (this *QProgressDialog) CloseEvent(event unsafe.Pointer) {
	// 0: (, event QCloseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:118
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QProgressDialog) ChangeEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:119
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QProgressDialog) ShowEvent(event unsafe.Pointer) {
	// 0: (, event QShowEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:122
// index:0
// void forceShow()
func (this *QProgressDialog) ForceShow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog9forceShowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
