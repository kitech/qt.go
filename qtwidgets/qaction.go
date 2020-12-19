package qtwidgets

// /usr/include/qt/QtWidgets/qaction.h
// #include <qaction.h>
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
// size 16
type QAction struct {
	*qtcore.QObject
}
type QAction_ITF interface {
	qtcore.QObject_ITF
	QAction_PTR() *QAction
}

func (ptr *QAction) QAction_PTR() *QAction { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QActionFromptr(cthis unsafe.Pointer) *QAction {
	bcthis0 := qtcore.QObjectFromptr(cthis)
	return &QAction{bcthis0}
}
func (*QAction) Fromptr(cthis unsafe.Pointer) *QAction {
	return QActionFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qaction.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAction(QObject *)

/*
 */
func (*QAction) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	return NewQAction(parent)
}
func NewQAction(parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc1(218008137, "_ZN7QActionC2EP7QObject", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QActionFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAction(QObject *)

/*
 */
func (*QAction) NewForInheritp() *QAction {
	return NewQActionp()
}
func NewQActionp() *QAction {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc1(218008137, "_ZN7QActionC2EP7QObject", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QActionFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAction(const QString &, QObject *)

/*
 */
func (*QAction) NewForInherit1(text string, parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	return NewQAction1(text, parent)
}
func NewQAction1(text string, parent qtcore.QObject_ITF /*777 QObject **/) *QAction {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc1(237064025, "_ZN7QActionC2ERK7QStringP7QObject", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QActionFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAction(const QString &, QObject *)

/*
 */
func (*QAction) NewForInherit1p(text string) *QAction {
	return NewQAction1p(text)
}
func NewQAction1p(text string) *QAction {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc1(237064025, "_ZN7QActionC2ERK7QStringP7QObject", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QActionFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

func DeleteQAction(this *QAction) {
	rv, err := qtrt.Qtcc1(0, "_ZN7QActionD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QAction__MenuRole = int

//
const QAction__NoRole QAction__MenuRole = 0

//
const QAction__TextHeuristicRole QAction__MenuRole = 1

//
const QAction__ApplicationSpecificRole QAction__MenuRole = 2

//
const QAction__AboutQtRole QAction__MenuRole = 3

//
const QAction__AboutRole QAction__MenuRole = 4

//
const QAction__PreferencesRole QAction__MenuRole = 5

//
const QAction__QuitRole QAction__MenuRole = 6

func (this *QAction) MenuRoleItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAction_MenuRoleItemName(val int) string {
	var nilthis *QAction
	return nilthis.MenuRoleItemName(val)
}

/*


 */
type QAction__Priority = int

//
const QAction__LowPriority QAction__Priority = 0

//
const QAction__NormalPriority QAction__Priority = 128

//
const QAction__HighPriority QAction__Priority = 256

func (this *QAction) PriorityItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAction_PriorityItemName(val int) string {
	var nilthis *QAction
	return nilthis.PriorityItemName(val)
}

/*


 */
type QAction__ActionEvent = int

//
const QAction__Trigger QAction__ActionEvent = 0

//
const QAction__Hover QAction__ActionEvent = 1

func (this *QAction) ActionEventItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAction_ActionEventItemName(val int) string {
	var nilthis *QAction
	return nilthis.ActionEventItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10095() {
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
