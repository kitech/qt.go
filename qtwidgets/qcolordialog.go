package qtwidgets

// /usr/include/qt/QtWidgets/qcolordialog.h
// #include <qcolordialog.h>
// #include <QtWidgets>

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
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QColorDialog) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQColorDialogFromPointer(cthis unsafe.Pointer) *QColorDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QColorDialog{bcthis0}
}
func (*QColorDialog) NewFromPointer(cthis unsafe.Pointer) *QColorDialog {
	return NewQColorDialogFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QColorDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QColorDialog10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcolordialog.h:71
// index:0
// Public
// void QColorDialog(class QWidget *)
func NewQColorDialog(parent *QWidget /*777 QWidget **/) *QColorDialog {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialogC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorDialogFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcolordialog.h:72
// index:1
// Public
// void QColorDialog(const class QColor &, class QWidget *)
func NewQColorDialog_1(initial *qtgui.QColor, parent *QWidget /*777 QWidget **/) *QColorDialog {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = initial.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialogC2ERK6QColorP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorDialogFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcolordialog.h:73
// index:0
// Public virtual
// void ~QColorDialog()
func DeleteQColorDialog(*QColorDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:75
// index:0
// Public
// void setCurrentColor(const class QColor &)
func (this *QColorDialog) SetCurrentColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog15setCurrentColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:76
// index:0
// Public
// QColor currentColor()
func (this *QColorDialog) CurrentColor() *qtgui.QColor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QColorDialog12currentColorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcolordialog.h:78
// index:0
// Public
// QColor selectedColor()
func (this *QColorDialog) SelectedColor() *qtgui.QColor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QColorDialog13selectedColorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcolordialog.h:80
// index:0
// Public
// void setOption(enum QColorDialog::ColorDialogOption, _Bool)
func (this *QColorDialog) SetOption(option int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog9setOptionENS_17ColorDialogOptionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:81
// index:0
// Public
// bool testOption(enum QColorDialog::ColorDialogOption)
func (this *QColorDialog) TestOption(option int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QColorDialog10testOptionENS_17ColorDialogOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcolordialog.h:83
// index:0
// Public
// QColorDialog::ColorDialogOptions options()
func (this *QColorDialog) Options() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QColorDialog7optionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:86
// index:0
// Public
// void open(class QObject *, const char *)
func (this *QColorDialog) Open(receiver *qtcore.QObject /*777 QObject **/, member string) {
	var convArg0 = receiver.GetCthis()
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:88
// index:0
// Public virtual
// void setVisible(_Bool)
func (this *QColorDialog) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:96
// index:0
// Public static
// QRgb getRgba(QRgb, _Bool *, class QWidget *)
func (this *QColorDialog) GetRgba(rgba uint, ok unsafe.Pointer /*666*/, parent *QWidget /*777 QWidget **/) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog7getRgbaEjPbP7QWidget", ffiqt.FFI_TYPE_POINTER, rgba, ok, parent)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QColorDialog_GetRgba(rgba uint, ok unsafe.Pointer /*666*/, parent *QWidget /*777 QWidget **/) uint {
	var nilthis *QColorDialog
	rv := nilthis.GetRgba(rgba, ok, parent)
	return rv
}

// /usr/include/qt/QtWidgets/qcolordialog.h:98
// index:0
// Public static
// int customCount()
func (this *QColorDialog) CustomCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog11customCountEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QColorDialog_CustomCount() int {
	var nilthis *QColorDialog
	rv := nilthis.CustomCount()
	return rv
}

// /usr/include/qt/QtWidgets/qcolordialog.h:99
// index:0
// Public static
// QColor customColor(int)
func (this *QColorDialog) CustomColor(index int) *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog11customColorEi", ffiqt.FFI_TYPE_POINTER, index)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QColorDialog_CustomColor(index int) *qtgui.QColor /*123*/ {
	var nilthis *QColorDialog
	rv := nilthis.CustomColor(index)
	return rv
}

// /usr/include/qt/QtWidgets/qcolordialog.h:100
// index:0
// Public static
// void setCustomColor(int, class QColor)
func (this *QColorDialog) SetCustomColor(index int, color *qtgui.QColor /*123*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog14setCustomColorEi6QColor", ffiqt.FFI_TYPE_POINTER, index, color)
	gopp.ErrPrint(err, rv)
}
func QColorDialog_SetCustomColor(index int, color *qtgui.QColor /*123*/) {
	var nilthis *QColorDialog
	nilthis.SetCustomColor(index, color)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:101
// index:0
// Public static
// QColor standardColor(int)
func (this *QColorDialog) StandardColor(index int) *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog13standardColorEi", ffiqt.FFI_TYPE_POINTER, index)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QColorDialog_StandardColor(index int) *qtgui.QColor /*123*/ {
	var nilthis *QColorDialog
	rv := nilthis.StandardColor(index)
	return rv
}

// /usr/include/qt/QtWidgets/qcolordialog.h:102
// index:0
// Public static
// void setStandardColor(int, class QColor)
func (this *QColorDialog) SetStandardColor(index int, color *qtgui.QColor /*123*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog16setStandardColorEi6QColor", ffiqt.FFI_TYPE_POINTER, index, color)
	gopp.ErrPrint(err, rv)
}
func QColorDialog_SetStandardColor(index int, color *qtgui.QColor /*123*/) {
	var nilthis *QColorDialog
	nilthis.SetStandardColor(index, color)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:105
// index:0
// Public
// void currentColorChanged(const class QColor &)
func (this *QColorDialog) CurrentColorChanged(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog19currentColorChangedERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:106
// index:0
// Public
// void colorSelected(const class QColor &)
func (this *QColorDialog) ColorSelected(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog13colorSelectedERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:109
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QColorDialog) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolordialog.h:110
// index:0
// Protected virtual
// void done(int)
func (this *QColorDialog) Done(result int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QColorDialog4doneEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), result)
	gopp.ErrPrint(err, rv)
}

type QColorDialog__ColorDialogOption = int

const QColorDialog__ShowAlphaChannel QColorDialog__ColorDialogOption = 1
const QColorDialog__NoButtons QColorDialog__ColorDialogOption = 2
const QColorDialog__DontUseNativeDialog QColorDialog__ColorDialogOption = 4

//  body block end
