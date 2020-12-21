package qtwidgets

// /usr/include/qt/QtWidgets/qsystemtrayicon.h
// #include <qsystemtrayicon.h>
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
// size 16
type QSystemTrayIcon struct {
	*qtcore.QObject
}
type QSystemTrayIcon_ITF interface {
	qtcore.QObject_ITF
	QSystemTrayIcon_PTR() *QSystemTrayIcon
}

func (ptr *QSystemTrayIcon) QSystemTrayIcon_PTR() *QSystemTrayIcon { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QSystemTrayIconFromptr(cthis Voidptr) *QSystemTrayIcon {
	bcthis0 := qtcore.QObjectFromptr(cthis)
	return &QSystemTrayIcon{bcthis0}
}
func (*QSystemTrayIcon) Fromptr(cthis Voidptr) *QSystemTrayIcon {
	return QSystemTrayIconFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemTrayIcon(QObject *)

/*
 */
func (*QSystemTrayIcon) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QSystemTrayIcon {
	return NewQSystemTrayIcon(parent)
}
func NewQSystemTrayIcon(parent qtcore.QObject_ITF /*777 QObject **/) *QSystemTrayIcon {
	var convArg0 Voidptr
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(3284962045, "_ZN15QSystemTrayIconC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QSystemTrayIconFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QSystemTrayIcon")
	return gothis
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSystemTrayIcon(QObject *)

/*
 */
func (*QSystemTrayIcon) NewForInheritp() *QSystemTrayIcon {
	return NewQSystemTrayIconp()
}
func NewQSystemTrayIconp() *QSystemTrayIcon {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(3284962045, "_ZN15QSystemTrayIconC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	gothis := QSystemTrayIconFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QSystemTrayIcon")
	return gothis
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:86
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QIcon icon() const

/*
 */
func (this *QSystemTrayIcon) Icon() *qtgui.QIcon /*123*/ {
	sretobj := qtrt.Malloc(8) // QIcon
	rv, err := qtrt.Qtcc3(2914842153, "_ZNK15QSystemTrayIcon4iconEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtgui.QIconFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:87
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)

/*
 */
func (this *QSystemTrayIcon) SetIcon(icon qtgui.QIcon_ITF) {
	var convArg0 Voidptr
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1753229533, "_ZN15QSystemTrayIcon7setIconERK5QIcon", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:89
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString toolTip() const

/*
 */
func (this *QSystemTrayIcon) ToolTip() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(3511699919, "_ZNK15QSystemTrayIcon7toolTipEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:90
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*
 */
func (this *QSystemTrayIcon) SetToolTip(tip string) {
	var tmpArg0 = qtcore.NewQString5(tip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(677619068, "_ZN15QSystemTrayIcon10setToolTipERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:92
// index:0
// Public static Extend Visibility=Default Availability=Available
// [1] bool isSystemTrayAvailable()

/*
 */
func (this *QSystemTrayIcon) IsSystemTrayAvailable() bool {
	rv, err := qtrt.Qtcc3(2178589267, "_ZN15QSystemTrayIcon21isSystemTrayAvailableEv", qtrt.FFITO_POINTER)
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}
func QSystemTrayIcon_IsSystemTrayAvailable() bool {
	var nilthis *QSystemTrayIcon
	rv := nilthis.IsSystemTrayAvailable()
	return rv
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:93
// index:0
// Public static Extend Visibility=Default Availability=Available
// [1] bool supportsMessages()

/*
 */
func (this *QSystemTrayIcon) SupportsMessages() bool {
	rv, err := qtrt.Qtcc3(85488849, "_ZN15QSystemTrayIcon16supportsMessagesEv", qtrt.FFITO_POINTER)
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}
func QSystemTrayIcon_SupportsMessages() bool {
	var nilthis *QSystemTrayIcon
	rv := nilthis.SupportsMessages()
	return rv
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:98
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isVisible() const

/*
 */
func (this *QSystemTrayIcon) IsVisible() bool {
	rv, err := qtrt.Qtcc3(2650858343, "_ZNK15QSystemTrayIcon9isVisibleEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:101
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
 */
func (this *QSystemTrayIcon) SetVisible(visible bool) {
	rv, err := qtrt.Qtcc3(2484040638, "_ZN15QSystemTrayIcon10setVisibleEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&visible))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:102
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void show()

/*
 */
func (this *QSystemTrayIcon) Show() {
	rv, err := qtrt.Qtcc3(1600631068, "_ZN15QSystemTrayIcon4showEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:103
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void hide()

/*
 */
func (this *QSystemTrayIcon) Hide() {
	rv, err := qtrt.Qtcc3(3283458795, "_ZN15QSystemTrayIcon4hideEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:104
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, const QIcon &, int)

/*
 */
func (this *QSystemTrayIcon) ShowMessage(title string, msg string, icon qtgui.QIcon_ITF, msecs int) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(msg)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 Voidptr
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg2 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1080679435, "_ZN15QSystemTrayIcon11showMessageERK7QStringS2_RK5QIconi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&convArg2), Voidptr(&msecs))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:104
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, const QIcon &, int)

/*
 */
func (this *QSystemTrayIcon) ShowMessagep(title string, msg string, icon qtgui.QIcon_ITF) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(msg)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 Voidptr
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg2 = icon.QIcon_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid, , Invalid
	msecs := int(10000)
	rv, err := qtrt.Qtcc3(1080679435, "_ZN15QSystemTrayIcon11showMessageERK7QStringS2_RK5QIconi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&convArg2), Voidptr(&msecs))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:105
// index:1
// Public Ignore Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, QSystemTrayIcon::MessageIcon, int)

/*
 */
func (this *QSystemTrayIcon) ShowMessage1(title string, msg string, icon int, msecs int) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(msg)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.Qtcc3(1630116740, "_ZN15QSystemTrayIcon11showMessageERK7QStringS2_NS_11MessageIconEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&icon), Voidptr(&msecs))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:105
// index:1
// Public Ignore Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, QSystemTrayIcon::MessageIcon, int)

/*
 */
func (this *QSystemTrayIcon) ShowMessage1p(title string, msg string) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(msg)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QSystemTrayIcon::MessageIcon=Elaborated, QSystemTrayIcon::MessageIcon=Enum, , Invalid
	icon := 0
	// arg: 3, int=Int, =Invalid, , Invalid
	msecs := int(10000)
	rv, err := qtrt.Qtcc3(1630116740, "_ZN15QSystemTrayIcon11showMessageERK7QStringS2_NS_11MessageIconEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&icon), Voidptr(&msecs))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:105
// index:1
// Public Ignore Visibility=Default Availability=Available
// [-2] void showMessage(const QString &, const QString &, QSystemTrayIcon::MessageIcon, int)

/*
 */
func (this *QSystemTrayIcon) ShowMessage1p1(title string, msg string, icon int) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString5(msg)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 3, int=Int, =Invalid, , Invalid
	msecs := int(10000)
	rv, err := qtrt.Qtcc3(1630116740, "_ZN15QSystemTrayIcon11showMessageERK7QStringS2_NS_11MessageIconEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&icon), Voidptr(&msecs))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:109
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void activated(QSystemTrayIcon::ActivationReason)

/*
 */
func (this *QSystemTrayIcon) Activated(reason int) {
	rv, err := qtrt.Qtcc3(4062607574, "_ZN15QSystemTrayIcon9activatedENS_16ActivationReasonE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&reason))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qsystemtrayicon.h:110
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void messageClicked()

/*
 */
func (this *QSystemTrayIcon) MessageClicked() {
	rv, err := qtrt.Qtcc3(2131465345, "_ZN15QSystemTrayIcon14messageClickedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

func DeleteQSystemTrayIcon(this *QSystemTrayIcon) {
	rv, err := qtrt.Qtcc1(0, "_ZN15QSystemTrayIconD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QSystemTrayIcon__ActivationReason = int

//
const QSystemTrayIcon__Unknown QSystemTrayIcon__ActivationReason = 0

//
const QSystemTrayIcon__Context QSystemTrayIcon__ActivationReason = 1

//
const QSystemTrayIcon__DoubleClick QSystemTrayIcon__ActivationReason = 2

//
const QSystemTrayIcon__Trigger QSystemTrayIcon__ActivationReason = 3

//
const QSystemTrayIcon__MiddleClick QSystemTrayIcon__ActivationReason = 4

func (this *QSystemTrayIcon) ActivationReasonItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSystemTrayIcon_ActivationReasonItemName(val int) string {
	var nilthis *QSystemTrayIcon
	return nilthis.ActivationReasonItemName(val)
}

/*


 */
type QSystemTrayIcon__MessageIcon = int

//
const QSystemTrayIcon__NoIcon QSystemTrayIcon__MessageIcon = 0

//
const QSystemTrayIcon__Information QSystemTrayIcon__MessageIcon = 1

//
const QSystemTrayIcon__Warning QSystemTrayIcon__MessageIcon = 2

//
const QSystemTrayIcon__Critical QSystemTrayIcon__MessageIcon = 3

func (this *QSystemTrayIcon) MessageIconItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSystemTrayIcon_MessageIconItemName(val int) string {
	var nilthis *QSystemTrayIcon
	return nilthis.MessageIconItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10157() {
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
