package qtwidgets

// /usr/include/qt/QtWidgets/qprogressdialog.h
// #include <qprogressdialog.h>
// #include <QtWidgets>

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QProgressDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:71
// index:0
// Public
// void QProgressDialog(class QWidget *, Qt::WindowFlags)
func NewQProgressDialog(parent *QWidget /*777 QWidget **/, flags int) *QProgressDialog {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQProgressDialogFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:72
// index:1
// Public
// void QProgressDialog(const class QString &, const class QString &, int, int, class QWidget *, Qt::WindowFlags)
func NewQProgressDialog_1(labelText *qtcore.QString, cancelButtonText *qtcore.QString, minimum int, maximum int, parent *QWidget /*777 QWidget **/, flags int) *QProgressDialog {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = labelText.GetCthis()
	var convArg1 = cancelButtonText.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialogC2ERK7QStringS2_iiP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, minimum, maximum, convArg4, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQProgressDialogFromPointer(cthis)
	return gothis
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
func (this *QProgressDialog) SetLabel(label *QLabel /*777 QLabel **/) {
	var convArg0 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setLabelEP6QLabel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:78
// index:0
// Public
// void setCancelButton(class QPushButton *)
func (this *QProgressDialog) SetCancelButton(button *QPushButton /*777 QPushButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog15setCancelButtonEP11QPushButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:79
// index:0
// Public
// void setBar(class QProgressBar *)
func (this *QProgressDialog) SetBar(bar *QProgressBar /*777 QProgressBar **/) {
	var convArg0 = bar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog6setBarEP12QProgressBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:81
// index:0
// Public
// bool wasCanceled()
func (this *QProgressDialog) WasCanceled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog11wasCanceledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:83
// index:0
// Public
// int minimum()
func (this *QProgressDialog) Minimum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog7minimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:84
// index:0
// Public
// int maximum()
func (this *QProgressDialog) Maximum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog7maximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:86
// index:0
// Public
// int value()
func (this *QProgressDialog) Value() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:88
// index:0
// Public virtual
// QSize sizeHint()
func (this *QProgressDialog) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:90
// index:0
// Public
// QString labelText()
func (this *QProgressDialog) LabelText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9labelTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:91
// index:0
// Public
// int minimumDuration()
func (this *QProgressDialog) MinimumDuration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog15minimumDurationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:93
// index:0
// Public
// void setAutoReset(_Bool)
func (this *QProgressDialog) SetAutoReset(reset bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoResetEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), reset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:94
// index:0
// Public
// bool autoReset()
func (this *QProgressDialog) AutoReset() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9autoResetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:95
// index:0
// Public
// void setAutoClose(_Bool)
func (this *QProgressDialog) SetAutoClose(close bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog12setAutoCloseEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), close)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:96
// index:0
// Public
// bool autoClose()
func (this *QProgressDialog) AutoClose() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QProgressDialog9autoCloseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:99
// index:0
// Public
// void open(class QObject *, const char *)
func (this *QProgressDialog) Open(receiver *qtcore.QObject /*777 QObject **/, member string) {
	var convArg0 = receiver.GetCthis()
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10setMaximumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:105
// index:0
// Public
// void setMinimum(int)
func (this *QProgressDialog) SetMinimum(minimum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10setMinimumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), minimum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:106
// index:0
// Public
// void setRange(int, int)
func (this *QProgressDialog) SetRange(minimum int, maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setRangeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), minimum, maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:107
// index:0
// Public
// void setValue(int)
func (this *QProgressDialog) SetValue(progress int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog8setValueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), progress)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog18setMinimumDurationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ms)
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
func (this *QProgressDialog) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:117
// index:0
// Protected virtual
// void closeEvent(class QCloseEvent *)
func (this *QProgressDialog) CloseEvent(event *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:118
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QProgressDialog) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qprogressdialog.h:119
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QProgressDialog) ShowEvent(event *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QProgressDialog9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
