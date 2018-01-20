//  header block begin
// /usr/include/qt/QtWidgets/qfontdialog.h
// #include <qfontdialog.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QFontDialog struct {
	*QDialog
}

func (this *QFontDialog) GetCthis() unsafe.Pointer {
	return this.QDialog.GetCthis()
}
func NewQFontDialogFromPointer(cthis unsafe.Pointer) *QFontDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QFontDialog{bcthis0}
}

// /usr/include/qt/QtWidgets/qfontdialog.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFontDialog) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFontDialog10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfontdialog.h:75
// index:0
// Public
// void QFontDialog(class QWidget *)
func NewQFontDialog(parent unsafe.Pointer) *QFontDialog {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialogC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontDialogFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qfontdialog.h:76
// index:1
// Public
// void QFontDialog(const class QFont &, class QWidget *)
func NewQFontDialog_1(initial *qtgui.QFont, parent unsafe.Pointer) *QFontDialog {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = initial.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialogC2ERK5QFontP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontDialogFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qfontdialog.h:77
// index:0
// Public virtual
// void ~QFontDialog()
func DeleteQFontDialog(*QFontDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:79
// index:0
// Public
// void setCurrentFont(const class QFont &)
func (this *QFontDialog) SetCurrentFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog14setCurrentFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:80
// index:0
// Public
// QFont currentFont()
func (this *QFontDialog) CurrentFont() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFontDialog11currentFontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfontdialog.h:82
// index:0
// Public
// QFont selectedFont()
func (this *QFontDialog) SelectedFont() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFontDialog12selectedFontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfontdialog.h:84
// index:0
// Public
// void setOption(enum QFontDialog::FontDialogOption, _Bool)
func (this *QFontDialog) SetOption(option int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog9setOptionENS_16FontDialogOptionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:85
// index:0
// Public
// bool testOption(enum QFontDialog::FontDialogOption)
func (this *QFontDialog) TestOption(option int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFontDialog10testOptionENS_16FontDialogOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfontdialog.h:87
// index:0
// Public
// QFontDialog::FontDialogOptions options()
func (this *QFontDialog) Options() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFontDialog7optionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfontdialog.h:90
// index:0
// Public
// void open(class QObject *, const char *)
func (this *QFontDialog) Open(receiver unsafe.Pointer, member string) {
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), receiver, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:92
// index:0
// Public virtual
// void setVisible(_Bool)
func (this *QFontDialog) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:94
// index:0
// Public static
// QFont getFont(_Bool *, class QWidget *)
func (this *QFontDialog) GetFont(ok unsafe.Pointer, parent unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbP7QWidget", ffiqt.FFI_TYPE_POINTER, ok, parent)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDialog_GetFont(ok unsafe.Pointer, parent unsafe.Pointer) {
	var nilthis *QFontDialog
	nilthis.GetFont(ok, parent)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:99
// index:0
// Public
// void currentFontChanged(const class QFont &)
func (this *QFontDialog) CurrentFontChanged(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog18currentFontChangedERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:100
// index:0
// Public
// void fontSelected(const class QFont &)
func (this *QFontDialog) FontSelected(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog12fontSelectedERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:103
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QFontDialog) ChangeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:104
// index:0
// Protected virtual
// void done(int)
func (this *QFontDialog) Done(result int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog4doneEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &result)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:105
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QFontDialog) EventFilter(object unsafe.Pointer, event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object, event)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
