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
// ignore GetCthis for 1 base
func QActionFromptr(cthis Voidptr) *QAction {
	bcthis0 := qtcore.QObjectFromptr(cthis)
	return &QAction{bcthis0}
}
func (*QAction) Fromptr(cthis Voidptr) *QAction {
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
	var convArg0 Voidptr
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(218008137, "_ZN7QActionC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
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
	var convArg0 Voidptr
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(218008137, "_ZN7QActionC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
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
	var convArg1 Voidptr
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(237064025, "_ZN7QActionC2ERK7QStringP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
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
	var convArg1 Voidptr
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(237064025, "_ZN7QActionC2ERK7QStringP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&convArg1))
	qtrt.ErrPrint2(err, rv)
	gothis := QActionFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAction")
	return gothis
}

// /usr/include/qt/QtWidgets/qaction.h:106
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*
 */
func (this *QAction) SetText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(3016118847, "_ZN7QAction7setTextERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:107
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString text() const

/*
 */
func (this *QAction) Text() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2317398157, "_ZNK7QAction4textEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:109
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setIconText(const QString &)

/*
 */
func (this *QAction) SetIconText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(4066431650, "_ZN7QAction11setIconTextERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:110
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString iconText() const

/*
 */
func (this *QAction) IconText() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(1536229387, "_ZNK7QAction8iconTextEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:112
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*
 */
func (this *QAction) SetToolTip(tip string) {
	var tmpArg0 = qtcore.NewQString5(tip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(1685067186, "_ZN7QAction10setToolTipERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:113
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString toolTip() const

/*
 */
func (this *QAction) ToolTip() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(3413091492, "_ZNK7QAction7toolTipEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:115
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)

/*
 */
func (this *QAction) SetStatusTip(statusTip string) {
	var tmpArg0 = qtcore.NewQString5(statusTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(4184243817, "_ZN7QAction12setStatusTipERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:116
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString statusTip() const

/*
 */
func (this *QAction) StatusTip() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(118887278, "_ZNK7QAction9statusTipEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qaction.h:150
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCheckable(bool)

/*
 */
func (this *QAction) SetCheckable(arg0 bool) {
	rv, err := qtrt.Qtcc3(1026958800, "_ZN7QAction12setCheckableEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:151
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isCheckable() const

/*
 */
func (this *QAction) IsCheckable() bool {
	rv, err := qtrt.Qtcc3(3197261506, "_ZNK7QAction11isCheckableEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:156
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isChecked() const

/*
 */
func (this *QAction) IsChecked() bool {
	rv, err := qtrt.Qtcc3(939532047, "_ZNK7QAction9isCheckedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:158
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isEnabled() const

/*
 */
func (this *QAction) IsEnabled() bool {
	rv, err := qtrt.Qtcc3(4605092, "_ZNK7QAction9isEnabledEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:160
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isVisible() const

/*
 */
func (this *QAction) IsVisible() bool {
	rv, err := qtrt.Qtcc3(3333608945, "_ZNK7QAction9isVisibleEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qaction.h:187
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void trigger()

/*
 */
func (this *QAction) Trigger() {
	rv, err := qtrt.Qtcc3(3248071529, "_ZN7QAction7triggerEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:188
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void hover()

/*
 */
func (this *QAction) Hover() {
	rv, err := qtrt.Qtcc3(606202509, "_ZN7QAction5hoverEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:189
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setChecked(bool)

/*
 */
func (this *QAction) SetChecked(arg0 bool) {
	rv, err := qtrt.Qtcc3(3927186894, "_ZN7QAction10setCheckedEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:190
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void toggle()

/*
 */
func (this *QAction) Toggle() {
	rv, err := qtrt.Qtcc3(612007024, "_ZN7QAction6toggleEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:191
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*
 */
func (this *QAction) SetEnabled(arg0 bool) {
	rv, err := qtrt.Qtcc3(3528607333, "_ZN7QAction10setEnabledEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:192
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setDisabled(bool)

/*
 */
func (this *QAction) SetDisabled(b bool) {
	rv, err := qtrt.Qtcc3(964827636, "_ZN7QAction11setDisabledEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&b))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:193
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
 */
func (this *QAction) SetVisible(arg0 bool) {
	rv, err := qtrt.Qtcc3(346478384, "_ZN7QAction10setVisibleEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:196
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void changed()

/*
 */
func (this *QAction) Changed() {
	rv, err := qtrt.Qtcc3(3892863327, "_ZN7QAction7changedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:197
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void triggered(bool)

/*
 */
func (this *QAction) Triggered(checked bool) {
	rv, err := qtrt.Qtcc3(3257549264, "_ZN7QAction9triggeredEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&checked))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:197
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void triggered(bool)

/*
 */
func (this *QAction) Triggeredp() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	checked := false
	rv, err := qtrt.Qtcc3(3257549264, "_ZN7QAction9triggeredEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&checked))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:198
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void hovered()

/*
 */
func (this *QAction) Hovered() {
	rv, err := qtrt.Qtcc3(2538730, "_ZN7QAction7hoveredEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qaction.h:199
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void toggled(bool)

/*
 */
func (this *QAction) Toggled(arg0 bool) {
	rv, err := qtrt.Qtcc3(62506592, "_ZN7QAction7toggledEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQAction(this *QAction) {
	rv, err := qtrt.Qtcc3(2632439783, "_ZN7QActionD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
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

func init_unused_10185() {
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
