//  header block begin
// /usr/include/qt/QtWidgets/qfontcombobox.h
// #include <qfontcombobox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
	*QComboBox
}

func (this *QFontComboBox) GetCthis() unsafe.Pointer {
	return this.QComboBox.GetCthis()
}
func NewQFontComboBoxFromPointer(cthis unsafe.Pointer) *QFontComboBox {
	bcthis0 := NewQComboBoxFromPointer(cthis)
	return &QFontComboBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFontComboBox) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontComboBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:61
// index:0
// Public
// void QFontComboBox(class QWidget *)
func NewQFontComboBox(parent unsafe.Pointer) *QFontComboBox {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontComboBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:62
// index:0
// Public virtual
// void ~QFontComboBox()
func DeleteQFontComboBox(*QFontComboBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:64
// index:0
// Public
// void setWritingSystem(class QFontDatabase::WritingSystem)
func (this *QFontComboBox) SetWritingSystem(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBox16setWritingSystemEN13QFontDatabase13WritingSystemE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:65
// index:0
// Public
// QFontDatabase::WritingSystem writingSystem()
func (this *QFontComboBox) WritingSystem() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontComboBox13writingSystemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:78
// index:0
// Public
// QFontComboBox::FontFilters fontFilters()
func (this *QFontComboBox) FontFilters() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontComboBox11fontFiltersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:80
// index:0
// Public
// QFont currentFont()
func (this *QFontComboBox) CurrentFont() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontComboBox11currentFontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:81
// index:0
// Public virtual
// QSize sizeHint()
func (this *QFontComboBox) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontComboBox8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:84
// index:0
// Public
// void setCurrentFont(const class QFont &)
func (this *QFontComboBox) SetCurrentFont(f *qtgui.QFont) {
	var convArg0 = f.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBox14setCurrentFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:87
// index:0
// Public
// void currentFontChanged(const class QFont &)
func (this *QFontComboBox) CurrentFontChanged(f *qtgui.QFont) {
	var convArg0 = f.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBox18currentFontChangedERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfontcombobox.h:90
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QFontComboBox) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontComboBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
