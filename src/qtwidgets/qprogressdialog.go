//  header block begin
// /usr/include/qt/QtWidgets/qprogressdialog.h
// #include <qprogressdialog.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:59
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QProgressDialog) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
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
	// 0: (, QLabel * label), (label)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setLabelEP6QLabel", ffiqt.FFI_TYPE_VOID, this.cthis, label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:78
// index:0
// void setCancelButton(class QPushButton *)
func (this *QProgressDialog) SetCancelButton(button unsafe.Pointer) {
	// 0: (, QPushButton * button), (button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog15setCancelButtonEP11QPushButton", ffiqt.FFI_TYPE_VOID, this.cthis, button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:79
// index:0
// void setBar(class QProgressBar *)
func (this *QProgressDialog) SetBar(bar unsafe.Pointer) {
	// 0: (, QProgressBar * bar), (bar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog6setBarEP12QProgressBar", ffiqt.FFI_TYPE_VOID, this.cthis, bar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:81
// index:0
// bool wasCanceled()
func (this *QProgressDialog) WasCanceled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog11wasCanceledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:83
// index:0
// int minimum()
func (this *QProgressDialog) Minimum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog7minimumEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:84
// index:0
// int maximum()
func (this *QProgressDialog) Maximum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog7maximumEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:86
// index:0
// int value()
func (this *QProgressDialog) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog5valueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:88
// index:0
// virtual
// QSize sizeHint()
func (this *QProgressDialog) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:90
// index:0
// QString labelText()
func (this *QProgressDialog) LabelText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9labelTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:91
// index:0
// int minimumDuration()
func (this *QProgressDialog) MinimumDuration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog15minimumDurationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:93
// index:0
// void setAutoReset(_Bool)
func (this *QProgressDialog) SetAutoReset(reset bool) {
	// 0: (, bool reset), (&reset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoResetEb", ffiqt.FFI_TYPE_VOID, this.cthis, &reset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:94
// index:0
// bool autoReset()
func (this *QProgressDialog) AutoReset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9autoResetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:95
// index:0
// void setAutoClose(_Bool)
func (this *QProgressDialog) SetAutoClose(close bool) {
	// 0: (, bool close), (&close)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoCloseEb", ffiqt.FFI_TYPE_VOID, this.cthis, &close)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:96
// index:0
// bool autoClose()
func (this *QProgressDialog) AutoClose() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9autoCloseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:99
// index:0
// void open(class QObject *, const char *)
func (this *QProgressDialog) Open(receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (, QObject * receiver, const char * member), (receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.cthis, receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:102
// index:0
// void cancel()
func (this *QProgressDialog) Cancel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog6cancelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:103
// index:0
// void reset()
func (this *QProgressDialog) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog5resetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:104
// index:0
// void setMaximum(int)
func (this *QProgressDialog) SetMaximum(maximum int) {
	// 0: (, int maximum), (&maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10setMaximumEi", ffiqt.FFI_TYPE_VOID, this.cthis, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:105
// index:0
// void setMinimum(int)
func (this *QProgressDialog) SetMinimum(minimum int) {
	// 0: (, int minimum), (&minimum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10setMinimumEi", ffiqt.FFI_TYPE_VOID, this.cthis, &minimum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:106
// index:0
// void setRange(int, int)
func (this *QProgressDialog) SetRange(minimum int, maximum int) {
	// 0: (, int minimum, int maximum), (&minimum, &maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setRangeEii", ffiqt.FFI_TYPE_VOID, this.cthis, &minimum, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:107
// index:0
// void setValue(int)
func (this *QProgressDialog) SetValue(progress int) {
	// 0: (, int progress), (&progress)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setValueEi", ffiqt.FFI_TYPE_VOID, this.cthis, &progress)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:108
// index:0
// void setLabelText(const class QString &)
func (this *QProgressDialog) SetLabelText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setLabelTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:109
// index:0
// void setCancelButtonText(const class QString &)
func (this *QProgressDialog) SetCancelButtonText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog19setCancelButtonTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:110
// index:0
// void setMinimumDuration(int)
func (this *QProgressDialog) SetMinimumDuration(ms int) {
	// 0: (, int ms), (&ms)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog18setMinimumDurationEi", ffiqt.FFI_TYPE_VOID, this.cthis, &ms)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:113
// index:0
// void canceled()
func (this *QProgressDialog) Canceled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8canceledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
