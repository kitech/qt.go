package qtwidgets

// /usr/include/qt/QtWidgets/qdialog.h
// #include <qdialog.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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

// void keyPressEvent(class QKeyEvent *)
func (this *QDialog) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void closeEvent(class QCloseEvent *)
func (this *QDialog) InheritCloseEvent(f func(arg0 *qtgui.QCloseEvent /*777 QCloseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QDialog) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QDialog) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QDialog) InheritContextMenuEvent(f func(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QDialog) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// void adjustPosition(class QWidget *)
func (this *QDialog) InheritAdjustPosition(f func(arg0 *QWidget /*777 QWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "adjustPosition", f)
}

type QDialog struct {
	*QWidget
}
type QDialog_ITF interface {
	QWidget_ITF
	QDialog_PTR() *QDialog
}

func (ptr *QDialog) QDialog_PTR() *QDialog { return ptr }

func (this *QDialog) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QDialog) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQDialogFromPointer(cthis unsafe.Pointer) *QDialog {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDialog{bcthis0}
}
func (*QDialog) NewFromPointer(cthis unsafe.Pointer) *QDialog {
	return NewQDialogFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdialog.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QDialog10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdialog.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDialog(QWidget *, Qt::WindowFlags)
func NewQDialog(parent QWidget_ITF /*777 QWidget **/, f int) *QDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qdialog.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDialog()
func DeleteQDialog(this *QDialog) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialogD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdialog.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] int result()
func (this *QDialog) Result() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QDialog6resultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdialog.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QDialog) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)
func (this *QDialog) SetOrientation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QDialog) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QDialog11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExtension(QWidget *)
func (this *QDialog) SetExtension(extension QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if extension != nil && extension.QWidget_PTR() != nil {
		convArg0 = extension.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog12setExtensionEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * extension()
func (this *QDialog) Extension() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QDialog9extensionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdialog.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QDialog) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QDialog8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qdialog.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QDialog) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QDialog15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qdialog.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeGripEnabled(_Bool)
func (this *QDialog) SetSizeGripEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog18setSizeGripEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSizeGripEnabled()
func (this *QDialog) IsSizeGripEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QDialog17isSizeGripEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdialog.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModal(_Bool)
func (this *QDialog) SetModal(modal bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog8setModalEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modal)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResult(int)
func (this *QDialog) SetResult(r int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog9setResultEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), r)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished(int)
func (this *QDialog) Finished(result int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog8finishedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), result)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accepted()
func (this *QDialog) Accepted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog8acceptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rejected()
func (this *QDialog) Rejected() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog8rejectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void open()
func (this *QDialog) Open() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog4openEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int exec()
func (this *QDialog) Exec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog4execEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdialog.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void done(int)
func (this *QDialog) Done(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void accept()
func (this *QDialog) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reject()
func (this *QDialog) Reject() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog6rejectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showExtension(_Bool)
func (this *QDialog) ShowExtension(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog13showExtensionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:104
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QDialog) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:105
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)
func (this *QDialog) CloseEvent(arg0 qtgui.QCloseEvent_ITF /*777 QCloseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QCloseEvent_PTR() != nil {
		convArg0 = arg0.QCloseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:106
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QDialog) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:107
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QDialog) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QDialog) ContextMenuEvent(arg0 qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QContextMenuEvent_PTR() != nil {
		convArg0 = arg0.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QDialog) EventFilter(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdialog.h:112
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void adjustPosition(QWidget *)
func (this *QDialog) AdjustPosition(arg0 QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QDialog14adjustPositionEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QDialog__DialogCode = int

const QDialog__Rejected QDialog__DialogCode = 0
const QDialog__Accepted QDialog__DialogCode = 1

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
