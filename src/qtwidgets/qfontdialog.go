//  header block begin
// /usr/include/qt/QtWidgets/qfontdialog.h
// #include <qfontdialog.h>
// #include <QtWidgets>
package qtwidgets

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

// /usr/include/qt/QtWidgets/qfontdialog.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFontDialog) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFontDialog10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:75
// index:0
// void QFontDialog(class QWidget *)
func NewQFontDialog(parent unsafe.Pointer) *QFontDialog {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialogC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontDialogFromPointer(cthis)
	return gothis
}
func NewQFontDialogFromPointer(cthis unsafe.Pointer) *QFontDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QFontDialog{bcthis0}
}

// /usr/include/qt/QtWidgets/qfontdialog.h:76
// index:1
// void QFontDialog(const class QFont &, class QWidget *)
func NewQFontDialog_1(initial unsafe.Pointer, parent unsafe.Pointer) *QFontDialog {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialogC2ERK5QFontP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, initial, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontDialogFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qfontdialog.h:77
// index:0
// virtual
// void ~QFontDialog()
func DeleteQFontDialog(*QFontDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:79
// index:0
// void setCurrentFont(const class QFont &)
func (this *QFontDialog) SetCurrentFont(font unsafe.Pointer) {
	// 0: (, font const QFont &), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog14setCurrentFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:80
// index:0
// QFont currentFont()
func (this *QFontDialog) CurrentFont() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFontDialog11currentFontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:82
// index:0
// QFont selectedFont()
func (this *QFontDialog) SelectedFont() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFontDialog12selectedFontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:84
// index:0
// void setOption(enum QFontDialog::FontDialogOption, _Bool)
func (this *QFontDialog) SetOption(option int, on bool) {
	// 0: (, option QFontDialog::FontDialogOption, on bool), (&option, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog9setOptionENS_16FontDialogOptionEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:85
// index:0
// bool testOption(enum QFontDialog::FontDialogOption)
func (this *QFontDialog) TestOption(option int) {
	// 0: (, option QFontDialog::FontDialogOption), (&option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFontDialog10testOptionENS_16FontDialogOptionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:86
// index:0
// void setOptions(QFontDialog::FontDialogOptions)
func (this *QFontDialog) SetOptions(options int) {
	// 0: (, QFlags<QFontDialog::FontDialogOption> options), (options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog10setOptionsE6QFlagsINS_16FontDialogOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:87
// index:0
// QFontDialog::FontDialogOptions options()
func (this *QFontDialog) Options() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFontDialog7optionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:90
// index:0
// void open(class QObject *, const char *)
func (this *QFontDialog) Open(receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (, receiver QObject *, member const char *), (receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:92
// index:0
// virtual
// void setVisible(_Bool)
func (this *QFontDialog) SetVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:94
// index:0
// static
// QFont getFont(_Bool *, class QWidget *)
func (this *QFontDialog) GetFont(ok unsafe.Pointer, parent unsafe.Pointer) {
	// 0: (ok bool *, parent QWidget *), (ok, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbP7QWidget", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDialog_GetFont(ok unsafe.Pointer, parent unsafe.Pointer) {
	// 0: (ok bool *, parent QWidget *), (ok, parent)
	var nilthis *QFontDialog
	nilthis.GetFont(ok, parent)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:95
// index:1
// static
// QFont getFont(_Bool *, const class QFont &, class QWidget *, const class QString &, QFontDialog::FontDialogOptions)
func (this *QFontDialog) GetFont_1(ok unsafe.Pointer, initial unsafe.Pointer, parent unsafe.Pointer, title unsafe.Pointer, options int) {
	// 1: (ok bool *, initial const QFont &, parent QWidget *, title const QString &, QFlags<QFontDialog::FontDialogOption> options), (ok, initial, parent, title, options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog7getFontEPbRK5QFontP7QWidgetRK7QString6QFlagsINS_16FontDialogOptionEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDialog_GetFont_1(ok unsafe.Pointer, initial unsafe.Pointer, parent unsafe.Pointer, title unsafe.Pointer, options int) {
	// 1: (ok bool *, initial const QFont &, parent QWidget *, title const QString &, QFlags<QFontDialog::FontDialogOption> options), (ok, initial, parent, title, options)
	var nilthis *QFontDialog
	nilthis.GetFont_1(ok, initial, parent, title, options)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:99
// index:0
// void currentFontChanged(const class QFont &)
func (this *QFontDialog) CurrentFontChanged(font unsafe.Pointer) {
	// 0: (, font const QFont &), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog18currentFontChangedERK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:100
// index:0
// void fontSelected(const class QFont &)
func (this *QFontDialog) FontSelected(font unsafe.Pointer) {
	// 0: (, font const QFont &), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog12fontSelectedERK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:103
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QFontDialog) ChangeEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:104
// index:0
// virtual
// void done(int)
func (this *QFontDialog) Done(result int) {
	// 0: (, result int), (&result)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog4doneEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &result)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontdialog.h:105
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QFontDialog) EventFilter(object unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, object QObject *, event QEvent *), (object, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFontDialog11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), object, event)
	gopp.ErrPrint(err, rv)
}

//  body block end
