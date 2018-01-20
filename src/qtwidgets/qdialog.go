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
func NewQDialogFromPointer(cthis unsafe.Pointer) *QDialog {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDialog{bcthis0}
}

// /usr/include/qt/QtWidgets/qdialog.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDialog) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialog.h:64
// index:0
// Public virtual
// void ~QDialog()
func DeleteQDialog(*QDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:68
// index:0
// Public
// int result()
func (this *QDialog) Result() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog6resultEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialog.h:70
// index:0
// Public virtual
// void setVisible(_Bool)
func (this *QDialog) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:72
// index:0
// Public
// void setOrientation(Qt::Orientation)
func (this *QDialog) SetOrientation(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:73
// index:0
// Public
// Qt::Orientation orientation()
func (this *QDialog) Orientation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialog.h:75
// index:0
// Public
// void setExtension(class QWidget *)
func (this *QDialog) SetExtension(extension unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog12setExtensionEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:76
// index:0
// Public
// QWidget * extension()
func (this *QDialog) Extension() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog9extensionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialog.h:78
// index:0
// Public virtual
// QSize sizeHint()
func (this *QDialog) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialog.h:79
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QDialog) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialog.h:81
// index:0
// Public
// void setSizeGripEnabled(_Bool)
func (this *QDialog) SetSizeGripEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog18setSizeGripEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:82
// index:0
// Public
// bool isSizeGripEnabled()
func (this *QDialog) IsSizeGripEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog17isSizeGripEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialog.h:84
// index:0
// Public
// void setModal(_Bool)
func (this *QDialog) SetModal(modal bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8setModalEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &modal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:85
// index:0
// Public
// void setResult(int)
func (this *QDialog) SetResult(r int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog9setResultEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:88
// index:0
// Public
// void finished(int)
func (this *QDialog) Finished(result int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8finishedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &result)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:89
// index:0
// Public
// void accepted()
func (this *QDialog) Accepted() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8acceptedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:90
// index:0
// Public
// void rejected()
func (this *QDialog) Rejected() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8rejectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:93
// index:0
// Public virtual
// void open()
func (this *QDialog) Open() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog4openEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:94
// index:0
// Public virtual
// int exec()
func (this *QDialog) Exec() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog4execEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialog.h:95
// index:0
// Public virtual
// void done(int)
func (this *QDialog) Done(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog4doneEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:96
// index:0
// Public virtual
// void accept()
func (this *QDialog) Accept() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog6acceptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:97
// index:0
// Public virtual
// void reject()
func (this *QDialog) Reject() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog6rejectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:99
// index:0
// Public
// void showExtension(_Bool)
func (this *QDialog) ShowExtension(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog13showExtensionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:104
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QDialog) KeyPressEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:105
// index:0
// Protected virtual
// void closeEvent(class QCloseEvent *)
func (this *QDialog) CloseEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:106
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QDialog) ShowEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:107
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QDialog) ResizeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:109
// index:0
// Protected virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QDialog) ContextMenuEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:111
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QDialog) EventFilter(arg0 unsafe.Pointer, arg1 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qdialog.h:112
// index:0
// Protected
// void adjustPosition(class QWidget *)
func (this *QDialog) AdjustPosition(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog14adjustPositionEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
