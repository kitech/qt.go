//  header block begin
// /usr/include/qt/QtWidgets/qgroupbox.h
// #include <qgroupbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 80
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
type QGroupBox struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgroupbox.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGroupBox) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:62
// index:0
// void QGroupBox(class QWidget *)
func NewQGroupBox(parent unsafe.Pointer) *QGroupBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGroupBox{cthis}
}

// /usr/include/qt/QtWidgets/qgroupbox.h:63
// index:1
// void QGroupBox(const class QString &, class QWidget *)
func NewQGroupBox_1(title unsafe.Pointer, parent unsafe.Pointer) *QGroupBox {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBoxC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, title, parent)
	gopp.ErrPrint(err, rv)
	return &QGroupBox{cthis}
}

// /usr/include/qt/QtWidgets/qgroupbox.h:64
// index:0
// virtual
// void ~QGroupBox()
func DeleteQGroupBox(*QGroupBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:66
// index:0
// QString title()
func (this *QGroupBox) Title() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox5titleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:67
// index:0
// void setTitle(const class QString &)
func (this *QGroupBox) SetTitle(title unsafe.Pointer) {
	// 0: (, const QString & title), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox8setTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:69
// index:0
// Qt::Alignment alignment()
func (this *QGroupBox) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox9alignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:70
// index:0
// void setAlignment(int)
func (this *QGroupBox) SetAlignment(alignment int) {
	// 0: (, int alignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox12setAlignmentEi", ffiqt.FFI_TYPE_VOID, this.cthis, &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:72
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QGroupBox) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:74
// index:0
// bool isFlat()
func (this *QGroupBox) IsFlat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox6isFlatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:75
// index:0
// void setFlat(_Bool)
func (this *QGroupBox) SetFlat(flat bool) {
	// 0: (, bool flat), (&flat)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox7setFlatEb", ffiqt.FFI_TYPE_VOID, this.cthis, &flat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:76
// index:0
// bool isCheckable()
func (this *QGroupBox) IsCheckable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox11isCheckableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:77
// index:0
// void setCheckable(_Bool)
func (this *QGroupBox) SetCheckable(checkable bool) {
	// 0: (, bool checkable), (&checkable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox12setCheckableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &checkable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:78
// index:0
// bool isChecked()
func (this *QGroupBox) IsChecked() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QGroupBox9isCheckedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:81
// index:0
// void setChecked(_Bool)
func (this *QGroupBox) SetChecked(checked bool) {
	// 0: (, bool checked), (&checked)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox10setCheckedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:84
// index:0
// void clicked(_Bool)
func (this *QGroupBox) Clicked(checked bool) {
	// 0: (, bool checked), (&checked)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox7clickedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &checked)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgroupbox.h:85
// index:0
// void toggled(_Bool)
func (this *QGroupBox) Toggled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QGroupBox7toggledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
