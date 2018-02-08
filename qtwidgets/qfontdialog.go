package qtwidgets

// /usr/include/qt/QtWidgets/qfontdialog.h
// #include <qfontdialog.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
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
// void changeEvent(class QEvent *)
func (this *QFontDialog) InheritChangeEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void done(int)
func (this *QFontDialog) InheritDone(f func(result int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "done", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QFontDialog) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

type QFontDialog struct {
	*QDialog
}

func (this *QFontDialog) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QFontDialog) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQFontDialogFromPointer(cthis unsafe.Pointer) *QFontDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QFontDialog{bcthis0}
}
func (*QFontDialog) NewFromPointer(cthis unsafe.Pointer) *QFontDialog {
	return NewQFontDialogFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFontDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFontDialog10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfontdialog.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontDialog(QWidget *)
func NewQFontDialog(parent *QWidget /*777 QWidget **/) *QFontDialog {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialogC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qfontdialog.h:76
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFontDialog(const QFont &, QWidget *)
func NewQFontDialog_1(initial *qtgui.QFont, parent *QWidget /*777 QWidget **/) *QFontDialog {
	var convArg0 = initial.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialogC2ERK5QFontP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qfontdialog.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFontDialog()
func DeleteQFontDialog(this *QFontDialog) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialogD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentFont(const QFont &)
func (this *QFontDialog) SetCurrentFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog14setCurrentFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:80
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont currentFont()
func (this *QFontDialog) CurrentFont() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFontDialog11currentFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontdialog.h:82
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont selectedFont()
func (this *QFontDialog) SelectedFont() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFontDialog12selectedFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qfontdialog.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(enum QFontDialog::FontDialogOption, _Bool)
func (this *QFontDialog) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog9setOptionENS_16FontDialogOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(enum QFontDialog::FontDialogOption)
func (this *QFontDialog) TestOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFontDialog10testOptionENS_16FontDialogOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfontdialog.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QFontDialog::FontDialogOptions)
func (this *QFontDialog) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog10setOptionsE6QFlagsINS_16FontDialogOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] QFontDialog::FontDialogOptions options()
func (this *QFontDialog) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFontDialog7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open(QObject *, const char *)
func (this *QFontDialog) Open(receiver *qtcore.QObject /*777 QObject **/, member string) {
	var convArg0 = receiver.GetCthis()
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog4openEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QFontDialog) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [16] QFont getFont(_Bool *, QWidget *)
func (this *QFontDialog) GetFont(ok unsafe.Pointer /*666*/, parent *QWidget /*777 QWidget **/) *qtgui.QFont /*123*/ {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbP7QWidget", qtrt.FFI_TYPE_POINTER, &ok, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QFontDialog_GetFont(ok unsafe.Pointer /*666*/, parent *QWidget /*777 QWidget **/) *qtgui.QFont /*123*/ {
	var nilthis *QFontDialog
	rv := nilthis.GetFont(ok, parent)
	return rv
}

// /usr/include/qt/QtWidgets/qfontdialog.h:95
// index:1
// Public static Visibility=Default Availability=Available
// [16] QFont getFont(_Bool *, const QFont &, QWidget *, const QString &, QFontDialog::FontDialogOptions)
func (this *QFontDialog) GetFont_1(ok unsafe.Pointer /*666*/, initial *qtgui.QFont, parent *QWidget /*777 QWidget **/, title string, options int) *qtgui.QFont /*123*/ {
	var convArg1 = initial.GetCthis()
	var convArg2 = parent.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(title)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbRK5QFontP7QWidgetRK7QString6QFlagsINS_16FontDialogOptionEE", qtrt.FFI_TYPE_POINTER, &ok, convArg1, convArg2, convArg3, options)
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QFontDialog_GetFont_1(ok unsafe.Pointer /*666*/, initial *qtgui.QFont, parent *QWidget /*777 QWidget **/, title string, options int) *qtgui.QFont /*123*/ {
	var nilthis *QFontDialog
	rv := nilthis.GetFont_1(ok, initial, parent, title, options)
	return rv
}

// /usr/include/qt/QtWidgets/qfontdialog.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentFontChanged(const QFont &)
func (this *QFontDialog) CurrentFontChanged(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog18currentFontChangedERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fontSelected(const QFont &)
func (this *QFontDialog) FontSelected(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog12fontSelectedERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:103
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QFontDialog) ChangeEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:104
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void done(int)
func (this *QFontDialog) Done(result int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), result)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:105
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QFontDialog) EventFilter(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = object.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFontDialog11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

type QFontDialog__FontDialogOption = int

const QFontDialog__NoButtons QFontDialog__FontDialogOption = 1
const QFontDialog__DontUseNativeDialog QFontDialog__FontDialogOption = 2
const QFontDialog__ScalableFonts QFontDialog__FontDialogOption = 4
const QFontDialog__NonScalableFonts QFontDialog__FontDialogOption = 8
const QFontDialog__MonospacedFonts QFontDialog__FontDialogOption = 16
const QFontDialog__ProportionalFonts QFontDialog__FontDialogOption = 32

//  body block end
