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
// extern C begin: 1
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
func QPlainTextEditFromptr(cthis unsafe.Pointer) *QPlainTextEdit {
	bcthis0 := QAbstractScrollAreaFromptr(cthis)
	return &QPlainTextEdit{bcthis0}
}
func (*QPlainTextEdit) Fromptr(cthis unsafe.Pointer) *QPlainTextEdit {
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
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(396516486, "_ZN14QPlainTextEditC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
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
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(396516486, "_ZN14QPlainTextEditC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
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
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2780232515, "_ZN14QPlainTextEditC2ERK7QStringP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
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
	var convArg1 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2780232515, "_ZN14QPlainTextEditC2ERK7QStringP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QPlainTextEditFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

func DeleteQPlainTextEdit(this *QPlainTextEdit) {
	rv, err := qtrt.Qtcc1(0, "_ZN14QPlainTextEditD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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

func init_unused_10139() {
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
