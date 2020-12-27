// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qplaintextedit.h
// #include <qplaintextedit.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

/*
 */
// size 48
type QPlainTextEdit struct {
	*QAbstractScrollArea
}
type QPlainTextEdit_ITF interface {
	QAbstractScrollArea_ITF
	QPlainTextEdit_PTR() *QPlainTextEdit
}

func (ptr *QPlainTextEdit) QPlainTextEdit_PTR() *QPlainTextEdit { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QPlainTextEditFromptr(cthis Voidptr) *QPlainTextEdit {
	bcthis0 := QAbstractScrollAreaFromptr(cthis)
	return &QPlainTextEdit{bcthis0}
}
func (*QPlainTextEdit) Fromptr(cthis Voidptr) *QPlainTextEdit {
	return QPlainTextEditFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(QWidget *)

/*
 */
func (*QPlainTextEdit) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QPlainTextEdit {
	return NewQPlainTextEdit(parent)
}
func NewQPlainTextEdit(parent QWidget_ITF /*777 QWidget **/) *QPlainTextEdit {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(396516486, "_ZN14QPlainTextEditC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QPlainTextEditFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(QWidget *)

/*
 */
func (*QPlainTextEdit) NewForInheritp() *QPlainTextEdit {
	return NewQPlainTextEditp()
}
func NewQPlainTextEditp() *QPlainTextEdit {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(396516486, "_ZN14QPlainTextEditC2EP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QPlainTextEditFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(const QString &, QWidget *)

/*
 */
func (*QPlainTextEdit) NewForInherit1(text string, parent QWidget_ITF /*777 QWidget **/) *QPlainTextEdit {
	return NewQPlainTextEdit1(text, parent)
}
func NewQPlainTextEdit1(text string, parent QWidget_ITF /*777 QWidget **/) *QPlainTextEdit {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(2780232515, "_ZN14QPlainTextEditC2ERK7QStringP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	gothis := QPlainTextEditFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(const QString &, QWidget *)

/*
 */
func (*QPlainTextEdit) NewForInherit1p(text string) *QPlainTextEdit {
	return NewQPlainTextEdit1p(text)
}
func NewQPlainTextEdit1p(text string) *QPlainTextEdit {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 Voidptr
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(2780232515, "_ZN14QPlainTextEditC2ERK7QStringP7QWidget", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	gothis := QPlainTextEditFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:102
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setPlaceholderText(const QString &)

/*
 */
func (this *QPlainTextEdit) SetPlaceholderText(placeholderText string) {
	var tmpArg0 = qtcore.NewQString5(placeholderText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(3191527676, "_ZN14QPlainTextEdit18setPlaceholderTextERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:103
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString placeholderText() const

/*
 */
func (this *QPlainTextEdit) PlaceholderText() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(388573047, "_ZNK14QPlainTextEdit15placeholderTextEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(rv.Ptr())
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:109
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setReadOnly(bool)

/*
 */
func (this *QPlainTextEdit) SetReadOnly(ro bool) {
	rv, err := qtrt.Qtcc3(3324855670, "_ZN14QPlainTextEdit11setReadOnlyEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&ro))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextInteractionFlags(Qt::TextInteractionFlags)

/*
 */
func (this *QPlainTextEdit) SetTextInteractionFlags(flags int) {
	rv, err := qtrt.Qtcc3(400835228, "_ZN14QPlainTextEdit23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&flags))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:143
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setBackgroundVisible(bool)

/*
 */
func (this *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	rv, err := qtrt.Qtcc3(957632610, "_ZN14QPlainTextEdit20setBackgroundVisibleEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&visible))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:157
// index:0
// Public inline Indirect Visibility=Default Availability=Available
// [8] QString toPlainText() const

/*
 */
func (this *QPlainTextEdit) ToPlainText() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(4142553692, "_ZNK14QPlainTextEdit11toPlainTextEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(rv.Ptr())
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:160
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void ensureCursorVisible()

/*
 */
func (this *QPlainTextEdit) EnsureCursorVisible() {
	rv, err := qtrt.Qtcc3(3070872959, "_ZN14QPlainTextEdit19ensureCursorVisibleEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:203
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setPlainText(const QString &)

/*
 */
func (this *QPlainTextEdit) SetPlainText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(2634248857, "_ZN14QPlainTextEdit12setPlainTextERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:228
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void textChanged()

/*
 */
func (this *QPlainTextEdit) TextChanged() {
	rv, err := qtrt.Qtcc3(4250366701, "_ZN14QPlainTextEdit11textChangedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
}

func DeleteQPlainTextEdit(this *QPlainTextEdit) {
	rv, err := qtrt.Qtcc3(2127496755, "_ZN14QPlainTextEditD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QPlainTextEdit__LineWrapMode = int

//
const QPlainTextEdit__NoWrap QPlainTextEdit__LineWrapMode = 0

//
const QPlainTextEdit__WidgetWidth QPlainTextEdit__LineWrapMode = 1

func (this *QPlainTextEdit) LineWrapModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QPlainTextEdit_LineWrapModeItemName(val int) string {
	var nilthis *QPlainTextEdit
	return nilthis.LineWrapModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10239() {
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
