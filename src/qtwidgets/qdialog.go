//  header block begin
// /usr/include/qt/QtWidgets/qdialog.h
// #include <qdialog.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QDialog struct {
	*QWidget
}

func (this *QDialog) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qdialog.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDialog) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:63
// index:0
// void QDialog(class QWidget *, Qt::WindowFlags)
func NewQDialog(parent unsafe.Pointer, f int) *QDialog {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, parent, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogFromPointer(cthis)
	return gothis
}
func NewQDialogFromPointer(cthis unsafe.Pointer) *QDialog {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDialog{bcthis0}
}

// /usr/include/qt/QtWidgets/qdialog.h:102
// index:1
// void QDialog(class QDialogPrivate &, class QWidget *, Qt::WindowFlags)
func NewQDialog_1(arg0 unsafe.Pointer, parent unsafe.Pointer, f int) *QDialog {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialogC2ER14QDialogPrivateP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, arg0, parent, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdialog.h:64
// index:0
// virtual
// void ~QDialog()
func DeleteQDialog(*QDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:68
// index:0
// int result()
func (this *QDialog) Result() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog6resultEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:70
// index:0
// virtual
// void setVisible(_Bool)
func (this *QDialog) SetVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:72
// index:0
// void setOrientation(Qt::Orientation)
func (this *QDialog) SetOrientation(orientation int) {
	// 0: (, orientation Qt::Orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:73
// index:0
// Qt::Orientation orientation()
func (this *QDialog) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:75
// index:0
// void setExtension(class QWidget *)
func (this *QDialog) SetExtension(extension unsafe.Pointer) {
	// 0: (, extension QWidget *), (extension)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog12setExtensionEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), extension)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:76
// index:0
// QWidget * extension()
func (this *QDialog) Extension() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog9extensionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:78
// index:0
// virtual
// QSize sizeHint()
func (this *QDialog) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:79
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QDialog) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:81
// index:0
// void setSizeGripEnabled(_Bool)
func (this *QDialog) SetSizeGripEnabled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog18setSizeGripEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:82
// index:0
// bool isSizeGripEnabled()
func (this *QDialog) IsSizeGripEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog17isSizeGripEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:84
// index:0
// void setModal(_Bool)
func (this *QDialog) SetModal(modal bool) {
	// 0: (, modal bool), (&modal)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8setModalEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &modal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:85
// index:0
// void setResult(int)
func (this *QDialog) SetResult(r int) {
	// 0: (, r int), (&r)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog9setResultEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:88
// index:0
// void finished(int)
func (this *QDialog) Finished(result int) {
	// 0: (, result int), (&result)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8finishedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &result)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:89
// index:0
// void accepted()
func (this *QDialog) Accepted() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8acceptedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:90
// index:0
// void rejected()
func (this *QDialog) Rejected() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8rejectedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:93
// index:0
// virtual
// void open()
func (this *QDialog) Open() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog4openEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:94
// index:0
// virtual
// int exec()
func (this *QDialog) Exec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog4execEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:95
// index:0
// virtual
// void done(int)
func (this *QDialog) Done(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog4doneEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:96
// index:0
// virtual
// void accept()
func (this *QDialog) Accept() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog6acceptEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:97
// index:0
// virtual
// void reject()
func (this *QDialog) Reject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog6rejectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:99
// index:0
// void showExtension(_Bool)
func (this *QDialog) ShowExtension(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog13showExtensionEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:104
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QDialog) KeyPressEvent(arg0 unsafe.Pointer) {
	// 0: (, QKeyEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:105
// index:0
// virtual
// void closeEvent(class QCloseEvent *)
func (this *QDialog) CloseEvent(arg0 unsafe.Pointer) {
	// 0: (, QCloseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:106
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QDialog) ShowEvent(arg0 unsafe.Pointer) {
	// 0: (, QShowEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:107
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QDialog) ResizeEvent(arg0 unsafe.Pointer) {
	// 0: (, QResizeEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:109
// index:0
// virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QDialog) ContextMenuEvent(arg0 unsafe.Pointer) {
	// 0: (, QContextMenuEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:111
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QDialog) EventFilter(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, QObject * arg0, QEvent * arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:112
// index:0
// void adjustPosition(class QWidget *)
func (this *QDialog) AdjustPosition(arg0 unsafe.Pointer) {
	// 0: (, QWidget * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog14adjustPositionEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
