//  header block begin
// /usr/include/qt/QtWidgets/qfontcombobox.h
// #include <qfontcombobox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QFontComboBox struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFontComboBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontComboBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:61
// index:0
// void QFontComboBox(class QWidget *)
func NewQFontComboBox(parent unsafe.Pointer) *QFontComboBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QFontComboBox{cthis}
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:62
// index:0
// virtual
// void ~QFontComboBox()
func DeleteQFontComboBox(*QFontComboBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:64
// index:0
// void setWritingSystem(class QFontDatabase::WritingSystem)
func (this *QFontComboBox) SetWritingSystem(arg0 int) {
	// 0: (, QFontDatabase::WritingSystem arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBox16setWritingSystemEN13QFontDatabase13WritingSystemE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:65
// index:0
// QFontDatabase::WritingSystem writingSystem()
func (this *QFontComboBox) WritingSystem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontComboBox13writingSystemEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:78
// index:0
// QFontComboBox::FontFilters fontFilters()
func (this *QFontComboBox) FontFilters() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontComboBox11fontFiltersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:80
// index:0
// QFont currentFont()
func (this *QFontComboBox) CurrentFont() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontComboBox11currentFontEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:81
// index:0
// virtual
// QSize sizeHint()
func (this *QFontComboBox) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontComboBox8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:84
// index:0
// void setCurrentFont(const class QFont &)
func (this *QFontComboBox) SetCurrentFont(f unsafe.Pointer) {
	// 0: (, const QFont & f), (f)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBox14setCurrentFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:87
// index:0
// void currentFontChanged(const class QFont &)
func (this *QFontComboBox) CurrentFontChanged(f unsafe.Pointer) {
	// 0: (, const QFont & f), (f)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBox18currentFontChangedERK5QFont", ffiqt.FFI_TYPE_VOID, this.cthis, f)
	gopp.ErrPrint(err, rv)
}

//  body block end
