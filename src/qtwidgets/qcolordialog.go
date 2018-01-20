//  header block begin
// /usr/include/qt/QtWidgets/qcolordialog.h
// #include <qcolordialog.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
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
type QColorDialog struct {
	*QDialog
}

func (this *QColorDialog) GetCthis() unsafe.Pointer {
	return this.QDialog.GetCthis()
}

// /usr/include/qt/QtWidgets/qcolordialog.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QColorDialog) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QColorDialog10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:71
// index:0
// void QColorDialog(class QWidget *)
func NewQColorDialog(parent unsafe.Pointer) *QColorDialog {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialogC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorDialogFromPointer(cthis)
	return gothis
}
func NewQColorDialogFromPointer(cthis unsafe.Pointer) *QColorDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QColorDialog{bcthis0}
}

// /usr/include/qt/QtWidgets/qcolordialog.h:72
// index:1
// void QColorDialog(const class QColor &, class QWidget *)
func NewQColorDialog_1(initial unsafe.Pointer, parent unsafe.Pointer) *QColorDialog {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialogC2ERK6QColorP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, initial, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorDialogFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcolordialog.h:73
// index:0
// virtual
// void ~QColorDialog()
func DeleteQColorDialog(*QColorDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:75
// index:0
// void setCurrentColor(const class QColor &)
func (this *QColorDialog) SetCurrentColor(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog15setCurrentColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:76
// index:0
// QColor currentColor()
func (this *QColorDialog) CurrentColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QColorDialog12currentColorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:78
// index:0
// QColor selectedColor()
func (this *QColorDialog) SelectedColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QColorDialog13selectedColorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:80
// index:0
// void setOption(enum QColorDialog::ColorDialogOption, _Bool)
func (this *QColorDialog) SetOption(option int, on bool) {
	// 0: (, option QColorDialog::ColorDialogOption, on bool), (&option, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog9setOptionENS_17ColorDialogOptionEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:81
// index:0
// bool testOption(enum QColorDialog::ColorDialogOption)
func (this *QColorDialog) TestOption(option int) {
	// 0: (, option QColorDialog::ColorDialogOption), (&option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QColorDialog10testOptionENS_17ColorDialogOptionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:82
// index:0
// void setOptions(QColorDialog::ColorDialogOptions)
func (this *QColorDialog) SetOptions(options int) {
	// 0: (, QFlags<QColorDialog::ColorDialogOption> options), (options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog10setOptionsE6QFlagsINS_17ColorDialogOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:83
// index:0
// QColorDialog::ColorDialogOptions options()
func (this *QColorDialog) Options() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QColorDialog7optionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:86
// index:0
// void open(class QObject *, const char *)
func (this *QColorDialog) Open(receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (, receiver QObject *, member const char *), (receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:88
// index:0
// virtual
// void setVisible(_Bool)
func (this *QColorDialog) SetVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:90
// index:0
// static
// QColor getColor(const class QColor &, class QWidget *, const class QString &, QColorDialog::ColorDialogOptions)
func (this *QColorDialog) GetColor(initial unsafe.Pointer, parent unsafe.Pointer, title unsafe.Pointer, options int) {
	// 0: (initial const QColor &, parent QWidget *, title const QString &, QFlags<QColorDialog::ColorDialogOption> options), (initial, parent, title, options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog8getColorERK6QColorP7QWidgetRK7QString6QFlagsINS_17ColorDialogOptionEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColorDialog_GetColor(initial unsafe.Pointer, parent unsafe.Pointer, title unsafe.Pointer, options int) {
	// 0: (initial const QColor &, parent QWidget *, title const QString &, QFlags<QColorDialog::ColorDialogOption> options), (initial, parent, title, options)
	var nilthis *QColorDialog
	nilthis.GetColor(initial, parent, title, options)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:96
// index:0
// static
// QRgb getRgba(QRgb, _Bool *, class QWidget *)
func (this *QColorDialog) GetRgba(rgba uint, ok unsafe.Pointer, parent unsafe.Pointer) {
	// 0: (rgba QRgb, ok bool *, parent QWidget *), (rgba, ok, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog7getRgbaEjPbP7QWidget", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColorDialog_GetRgba(rgba uint, ok unsafe.Pointer, parent unsafe.Pointer) {
	// 0: (rgba QRgb, ok bool *, parent QWidget *), (rgba, ok, parent)
	var nilthis *QColorDialog
	nilthis.GetRgba(rgba, ok, parent)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:98
// index:0
// static
// int customCount()
func (this *QColorDialog) CustomCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog11customCountEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColorDialog_CustomCount() {
	// 0: (), ()
	var nilthis *QColorDialog
	nilthis.CustomCount()
}

// /usr/include/qt/QtWidgets/qcolordialog.h:99
// index:0
// static
// QColor customColor(int)
func (this *QColorDialog) CustomColor(index int) {
	// 0: (index int), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog11customColorEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColorDialog_CustomColor(index int) {
	// 0: (index int), (index)
	var nilthis *QColorDialog
	nilthis.CustomColor(index)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:100
// index:0
// static
// void setCustomColor(int, class QColor)
func (this *QColorDialog) SetCustomColor(index int, color unsafe.Pointer) {
	// 0: (index int, color QColor), (index, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog14setCustomColorEi6QColor", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColorDialog_SetCustomColor(index int, color unsafe.Pointer) {
	// 0: (index int, color QColor), (index, color)
	var nilthis *QColorDialog
	nilthis.SetCustomColor(index, color)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:101
// index:0
// static
// QColor standardColor(int)
func (this *QColorDialog) StandardColor(index int) {
	// 0: (index int), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog13standardColorEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColorDialog_StandardColor(index int) {
	// 0: (index int), (index)
	var nilthis *QColorDialog
	nilthis.StandardColor(index)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:102
// index:0
// static
// void setStandardColor(int, class QColor)
func (this *QColorDialog) SetStandardColor(index int, color unsafe.Pointer) {
	// 0: (index int, color QColor), (index, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog16setStandardColorEi6QColor", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColorDialog_SetStandardColor(index int, color unsafe.Pointer) {
	// 0: (index int, color QColor), (index, color)
	var nilthis *QColorDialog
	nilthis.SetStandardColor(index, color)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:105
// index:0
// void currentColorChanged(const class QColor &)
func (this *QColorDialog) CurrentColorChanged(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog19currentColorChangedERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:106
// index:0
// void colorSelected(const class QColor &)
func (this *QColorDialog) ColorSelected(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog13colorSelectedERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:109
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QColorDialog) ChangeEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:110
// index:0
// virtual
// void done(int)
func (this *QColorDialog) Done(result int) {
	// 0: (, result int), (&result)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog4doneEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &result)
	gopp.ErrPrint(err, rv)
}

//  body block end
