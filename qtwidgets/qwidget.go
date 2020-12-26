package qtwidgets

// /usr/include/qt/QtWidgets/qwidget.h
// #include <qwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QWidget struct {
	*qtcore.QObject
	*qtgui.QPaintDevice
}
type QWidget_ITF interface {
	qtcore.QObject_ITF
	qtgui.QPaintDevice_ITF
	QWidget_PTR() *QWidget
}

func (ptr *QWidget) QWidget_PTR() *QWidget { return ptr }

func (this *QWidget) GetCthis() Voidptr {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWidget) SetCthis(cthis Voidptr) {
	this.QObject = qtcore.QObjectFromptr(cthis)
	this.QPaintDevice = qtgui.QPaintDeviceFromptr(cthis)
}
func (this *QWidget) Addr() Voidptr {
	if this == nil {
		return nil
	} else {
		return this.QObject.Addr()
	}
}
func QWidgetFromptr(cthis Voidptr) *QWidget {
	bcthis0 := qtcore.QObjectFromptr(cthis)
	bcthis1 := qtgui.QPaintDeviceFromptr(cthis)
	return &QWidget{bcthis0, bcthis1}
}
func (*QWidget) Fromptr(cthis Voidptr) *QWidget {
	return QWidgetFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qwidget.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidget(QWidget *, Qt::WindowFlags)

/*
 */
func (*QWidget) NewForInherit(parent QWidget_ITF /*777 QWidget **/, f int) *QWidget {
	return NewQWidget(parent, f)
}
func NewQWidget(parent QWidget_ITF /*777 QWidget **/, f int) *QWidget {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(606334740, "_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
	gothis := QWidgetFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qwidget.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidget(QWidget *, Qt::WindowFlags)

/*
 */
func (*QWidget) NewForInheritp() *QWidget {
	return NewQWidgetp()
}
func NewQWidgetp() *QWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 Voidptr
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(606334740, "_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
	gothis := QWidgetFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qwidget.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidget(QWidget *, Qt::WindowFlags)

/*
 */
func (*QWidget) NewForInheritp1(parent QWidget_ITF /*777 QWidget **/) *QWidget {
	return NewQWidgetp1(parent)
}
func NewQWidgetp1(parent QWidget_ITF /*777 QWidget **/) *QWidget {
	var convArg0 Voidptr
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	f := 0
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc3(606334740, "_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&f))
	qtrt.ErrPrint2(err, rv)
	gothis := QWidgetFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qwidget.h:218
// index:0
// Public virtual Direct Visibility=Default Availability=Available
// [4] int devType() const

/*
 */
func (this *QWidget) DevType() int {
	rv, err := qtrt.Qtcc3(170809957, "_ZNK7QWidget7devTypeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:220
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] WId winId() const

/*
 */
func (this *QWidget) WinId() uint64 {
	rv, err := qtrt.Qtcc3(2263857736, "_ZNK7QWidget5winIdEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:222
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] WId internalWinId() const

/*
 */
func (this *QWidget) InternalWinId() uint64 {
	rv, err := qtrt.Qtcc3(546760316, "_ZNK7QWidget13internalWinIdEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:223
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] WId effectiveWinId() const

/*
 */
func (this *QWidget) EffectiveWinId() uint64 {
	rv, err := qtrt.Qtcc3(3257125328, "_ZNK7QWidget14effectiveWinIdEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:230
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isTopLevel() const

/*
 */
func (this *QWidget) IsTopLevel() bool {
	rv, err := qtrt.Qtcc3(434601045, "_ZNK7QWidget10isTopLevelEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:231
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isWindow() const

/*
 */
func (this *QWidget) IsWindow() bool {
	rv, err := qtrt.Qtcc3(612947221, "_ZNK7QWidget8isWindowEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:233
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isModal() const

/*
 */
func (this *QWidget) IsModal() bool {
	rv, err := qtrt.Qtcc3(4201175838, "_ZNK7QWidget7isModalEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:234
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] Qt::WindowModality windowModality() const

/*
 */
func (this *QWidget) WindowModality() int {
	rv, err := qtrt.Qtcc3(815097609, "_ZNK7QWidget14windowModalityEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:235
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setWindowModality(Qt::WindowModality)

/*
 */
func (this *QWidget) SetWindowModality(windowModality int) {
	rv, err := qtrt.Qtcc3(2668760869, "_ZN7QWidget17setWindowModalityEN2Qt14WindowModalityE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&windowModality))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:237
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isEnabled() const

/*
 */
func (this *QWidget) IsEnabled() bool {
	rv, err := qtrt.Qtcc3(1166952428, "_ZNK7QWidget9isEnabledEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:238
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isEnabledTo(const QWidget *) const

/*
 */
func (this *QWidget) IsEnabledTo(arg0 QWidget_ITF /*777 const QWidget **/) bool {
	var convArg0 Voidptr
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(3032138607, "_ZNK7QWidget11isEnabledToEPKS_", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:245
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*
 */
func (this *QWidget) SetEnabled(arg0 bool) {
	rv, err := qtrt.Qtcc3(1945868058, "_ZN7QWidget10setEnabledEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:246
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setDisabled(bool)

/*
 */
func (this *QWidget) SetDisabled(arg0 bool) {
	rv, err := qtrt.Qtcc3(4187610248, "_ZN7QWidget11setDisabledEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:247
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setWindowModified(bool)

/*
 */
func (this *QWidget) SetWindowModified(arg0 bool) {
	rv, err := qtrt.Qtcc3(800400522, "_ZN7QWidget17setWindowModifiedEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:256
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int x() const

/*
 */
func (this *QWidget) X() int {
	rv, err := qtrt.Qtcc3(2310490974, "_ZNK7QWidget1xEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:257
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int y() const

/*
 */
func (this *QWidget) Y() int {
	rv, err := qtrt.Qtcc3(2289377641, "_ZNK7QWidget1yEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:261
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int width() const

/*
 */
func (this *QWidget) Width() int {
	rv, err := qtrt.Qtcc3(3338959857, "_ZNK7QWidget5widthEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:262
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int height() const

/*
 */
func (this *QWidget) Height() int {
	rv, err := qtrt.Qtcc3(2964560984, "_ZNK7QWidget6heightEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:269
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int minimumWidth() const

/*
 */
func (this *QWidget) MinimumWidth() int {
	rv, err := qtrt.Qtcc3(828888399, "_ZNK7QWidget12minimumWidthEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:270
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int minimumHeight() const

/*
 */
func (this *QWidget) MinimumHeight() int {
	rv, err := qtrt.Qtcc3(192531497, "_ZNK7QWidget13minimumHeightEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:271
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int maximumWidth() const

/*
 */
func (this *QWidget) MaximumWidth() int {
	rv, err := qtrt.Qtcc3(2285981246, "_ZNK7QWidget12maximumWidthEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:272
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int maximumHeight() const

/*
 */
func (this *QWidget) MaximumHeight() int {
	rv, err := qtrt.Qtcc3(750955288, "_ZNK7QWidget13maximumHeightEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:274
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMinimumSize(int, int)

/*
 */
func (this *QWidget) SetMinimumSize(minw int, minh int) {
	rv, err := qtrt.Qtcc3(508498511, "_ZN7QWidget14setMinimumSizeEii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, this.Addr(), Voidptr(&minw), Voidptr(&minh))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:276
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaximumSize(int, int)

/*
 */
func (this *QWidget) SetMaximumSize(maxw int, maxh int) {
	rv, err := qtrt.Qtcc3(2808711486, "_ZN7QWidget14setMaximumSizeEii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, this.Addr(), Voidptr(&maxw), Voidptr(&maxh))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:277
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMinimumWidth(int)

/*
 */
func (this *QWidget) SetMinimumWidth(minw int) {
	rv, err := qtrt.Qtcc3(998451129, "_ZN7QWidget15setMinimumWidthEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&minw))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:278
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMinimumHeight(int)

/*
 */
func (this *QWidget) SetMinimumHeight(minh int) {
	rv, err := qtrt.Qtcc3(314665889, "_ZN7QWidget16setMinimumHeightEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&minh))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:279
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaximumWidth(int)

/*
 */
func (this *QWidget) SetMaximumWidth(maxw int) {
	rv, err := qtrt.Qtcc3(2191899848, "_ZN7QWidget15setMaximumWidthEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&maxw))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:280
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaximumHeight(int)

/*
 */
func (this *QWidget) SetMaximumHeight(maxh int) {
	rv, err := qtrt.Qtcc3(897191056, "_ZN7QWidget16setMaximumHeightEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&maxh))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:294
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setFixedSize(int, int)

/*
 */
func (this *QWidget) SetFixedSize(w int, h int) {
	rv, err := qtrt.Qtcc3(1882072225, "_ZN7QWidget12setFixedSizeEii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, this.Addr(), Voidptr(&w), Voidptr(&h))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:295
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setFixedWidth(int)

/*
 */
func (this *QWidget) SetFixedWidth(w int) {
	rv, err := qtrt.Qtcc3(349713959, "_ZN7QWidget13setFixedWidthEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&w))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:296
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setFixedHeight(int)

/*
 */
func (this *QWidget) SetFixedHeight(h int) {
	rv, err := qtrt.Qtcc3(2258465809, "_ZN7QWidget14setFixedHeightEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&h))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:332
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMouseTracking(bool)

/*
 */
func (this *QWidget) SetMouseTracking(enable bool) {
	rv, err := qtrt.Qtcc3(3977244840, "_ZN7QWidget16setMouseTrackingEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&enable))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:334
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool underMouse() const

/*
 */
func (this *QWidget) UnderMouse() bool {
	rv, err := qtrt.Qtcc3(4034935793, "_ZNK7QWidget10underMouseEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:365
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setWindowTitle(const QString &)

/*
 */
func (this *QWidget) SetWindowTitle(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(3223847508, "_ZN7QWidget14setWindowTitleERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:367
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStyleSheet(const QString &)

/*
 */
func (this *QWidget) SetStyleSheet(styleSheet string) {
	var tmpArg0 = qtcore.NewQString5(styleSheet)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(95364179, "_ZN7QWidget13setStyleSheetERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:371
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString styleSheet() const

/*
 */
func (this *QWidget) StyleSheet() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2046494658, "_ZNK7QWidget10styleSheetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:373
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString windowTitle() const

/*
 */
func (this *QWidget) WindowTitle() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2977210028, "_ZNK7QWidget11windowTitleEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:376
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setWindowIconText(const QString &)

/*
 */
func (this *QWidget) SetWindowIconText(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(4138903100, "_ZN7QWidget17setWindowIconTextERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:377
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString windowIconText() const

/*
 */
func (this *QWidget) WindowIconText() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(465509618, "_ZNK7QWidget14windowIconTextEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:378
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setWindowRole(const QString &)

/*
 */
func (this *QWidget) SetWindowRole(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(2541344425, "_ZN7QWidget13setWindowRoleERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:379
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString windowRole() const

/*
 */
func (this *QWidget) WindowRole() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(1153054232, "_ZNK7QWidget10windowRoleEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:380
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setWindowFilePath(const QString &)

/*
 */
func (this *QWidget) SetWindowFilePath(filePath string) {
	var tmpArg0 = qtcore.NewQString5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(2928768698, "_ZN7QWidget17setWindowFilePathERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:381
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString windowFilePath() const

/*
 */
func (this *QWidget) WindowFilePath() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(3270090626, "_ZNK7QWidget14windowFilePathEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:383
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setWindowOpacity(qreal)

/*
 */
func (this *QWidget) SetWindowOpacity(level float64) {
	rv, err := qtrt.Qtcc3(2415377882, "_ZN7QWidget16setWindowOpacityEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&level))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:384
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] qreal windowOpacity() const

/*
 */
func (this *QWidget) WindowOpacity() float64 {
	rv, err := qtrt.Qtcc3(3908986607, "_ZNK7QWidget13windowOpacityEv", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:386
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isWindowModified() const

/*
 */
func (this *QWidget) IsWindowModified() bool {
	rv, err := qtrt.Qtcc3(2356378208, "_ZNK7QWidget16isWindowModifiedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:388
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*
 */
func (this *QWidget) SetToolTip(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(2175679426, "_ZN7QWidget10setToolTipERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:389
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString toolTip() const

/*
 */
func (this *QWidget) ToolTip() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(602720826, "_ZNK7QWidget7toolTipEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwidget.h:390
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setToolTipDuration(int)

/*
 */
func (this *QWidget) SetToolTipDuration(msec int) {
	rv, err := qtrt.Qtcc3(1625368137, "_ZN7QWidget18setToolTipDurationEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&msec))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:391
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int toolTipDuration() const

/*
 */
func (this *QWidget) ToolTipDuration() int {
	rv, err := qtrt.Qtcc3(2630369908, "_ZNK7QWidget15toolTipDurationEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:420
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setFocus()

/*
 */
func (this *QWidget) SetFocus() {
	rv, err := qtrt.Qtcc3(1428955973, "_ZN7QWidget8setFocusEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:423
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isActiveWindow() const

/*
 */
func (this *QWidget) IsActiveWindow() bool {
	rv, err := qtrt.Qtcc3(3080343722, "_ZNK7QWidget14isActiveWindowEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:424
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void activateWindow()

/*
 */
func (this *QWidget) ActivateWindow() {
	rv, err := qtrt.Qtcc3(264887046, "_ZN7QWidget14activateWindowEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:425
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clearFocus()

/*
 */
func (this *QWidget) ClearFocus() {
	rv, err := qtrt.Qtcc3(1605758068, "_ZN7QWidget10clearFocusEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:463
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void update()

/*
 */
func (this *QWidget) Update() {
	rv, err := qtrt.Qtcc3(3733308912, "_ZN7QWidget6updateEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:467
// index:1
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void update(int, int, int, int)

/*
 */
func (this *QWidget) Update1(x int, y int, w int, h int) {
	rv, err := qtrt.Qtcc3(2935439564, "_ZN7QWidget6updateEiiii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_INT, this.Addr(), Voidptr(&x), Voidptr(&y), Voidptr(&w), Voidptr(&h))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:464
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void repaint()

/*
 */
func (this *QWidget) Repaint() {
	rv, err := qtrt.Qtcc3(1692267421, "_ZN7QWidget7repaintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:478
// index:0
// Public virtual Ignore Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
 */
func (this *QWidget) SetVisible(visible bool) {
	rv, err := qtrt.Qtcc3(3037660751, "_ZN7QWidget10setVisibleEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&visible))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:479
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setHidden(bool)

/*
 */
func (this *QWidget) SetHidden(hidden bool) {
	rv, err := qtrt.Qtcc3(1864783706, "_ZN7QWidget9setHiddenEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&hidden))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:480
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void show()

/*
 */
func (this *QWidget) Show() {
	rv, err := qtrt.Qtcc3(3435008533, "_ZN7QWidget4showEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:481
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void hide()

/*
 */
func (this *QWidget) Hide() {
	rv, err := qtrt.Qtcc3(1349269986, "_ZN7QWidget4hideEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:483
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void showMinimized()

/*
 */
func (this *QWidget) ShowMinimized() {
	rv, err := qtrt.Qtcc3(661499765, "_ZN7QWidget13showMinimizedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:484
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void showMaximized()

/*
 */
func (this *QWidget) ShowMaximized() {
	rv, err := qtrt.Qtcc3(3948991610, "_ZN7QWidget13showMaximizedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:485
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void showFullScreen()

/*
 */
func (this *QWidget) ShowFullScreen() {
	rv, err := qtrt.Qtcc3(1342832539, "_ZN7QWidget14showFullScreenEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:486
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void showNormal()

/*
 */
func (this *QWidget) ShowNormal() {
	rv, err := qtrt.Qtcc3(1806854372, "_ZN7QWidget10showNormalEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:488
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool close()

/*
 */
func (this *QWidget) Close() bool {
	rv, err := qtrt.Qtcc3(1976830908, "_ZN7QWidget5closeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:489
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void raise()

/*
 */
func (this *QWidget) Raise() {
	rv, err := qtrt.Qtcc3(2275637944, "_ZN7QWidget5raiseEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:490
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void lower()

/*
 */
func (this *QWidget) Lower() {
	rv, err := qtrt.Qtcc3(964913166, "_ZN7QWidget5lowerEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:496
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void resize(int, int)

/*
 */
func (this *QWidget) Resize(w int, h int) {
	rv, err := qtrt.Qtcc3(930448380, "_ZN7QWidget6resizeEii", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, this.Addr(), Voidptr(&w), Voidptr(&h))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:502
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void adjustSize()

/*
 */
func (this *QWidget) AdjustSize() {
	rv, err := qtrt.Qtcc3(2214362968, "_ZN7QWidget10adjustSizeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:503
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isVisible() const

/*
 */
func (this *QWidget) IsVisible() bool {
	rv, err := qtrt.Qtcc3(2205865657, "_ZNK7QWidget9isVisibleEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:504
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isVisibleTo(const QWidget *) const

/*
 */
func (this *QWidget) IsVisibleTo(arg0 QWidget_ITF /*777 const QWidget **/) bool {
	var convArg0 Voidptr
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1859694289, "_ZNK7QWidget11isVisibleToEPKS_", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:505
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool isHidden() const

/*
 */
func (this *QWidget) IsHidden() bool {
	rv, err := qtrt.Qtcc3(1556262255, "_ZNK7QWidget8isHiddenEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:507
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isMinimized() const

/*
 */
func (this *QWidget) IsMinimized() bool {
	rv, err := qtrt.Qtcc3(1936793402, "_ZNK7QWidget11isMinimizedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:508
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isMaximized() const

/*
 */
func (this *QWidget) IsMaximized() bool {
	rv, err := qtrt.Qtcc3(3212600373, "_ZNK7QWidget11isMaximizedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:509
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool isFullScreen() const

/*
 */
func (this *QWidget) IsFullScreen() bool {
	rv, err := qtrt.Qtcc3(2278664472, "_ZNK7QWidget12isFullScreenEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:537
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QLayout * layout() const

/*
 */
func (this *QWidget) Layout() *QLayout /*777 QLayout **/ {
	rv, err := qtrt.Qtcc3(1365471401, "_ZNK7QWidget6layoutEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QLayoutFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:538
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setLayout(QLayout *)

/*
 */
func (this *QWidget) SetLayout(arg0 QLayout_ITF /*777 QLayout **/) {
	var convArg0 Voidptr
	if arg0 != nil && arg0.QLayout_PTR() != nil {
		convArg0 = arg0.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1717599370, "_ZN7QWidget9setLayoutEP7QLayout", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:539
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void updateGeometry()

/*
 */
func (this *QWidget) UpdateGeometry() {
	rv, err := qtrt.Qtcc3(2348469980, "_ZN7QWidget14updateGeometryEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:572
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QWidget * parentWidget() const

/*
 */
func (this *QWidget) ParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.Qtcc3(2946385333, "_ZNK7QWidget12parentWidgetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QWidgetFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwidget.h:612
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void windowTitleChanged(const QString &)

/*
 */
func (this *QWidget) WindowTitleChanged(title string) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(3974145108, "_ZN7QWidget18windowTitleChangedERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

func DeleteQWidget(this *QWidget) {
	rv, err := qtrt.Qtcc1(0, "_ZN7QWidgetD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QWidget__RenderFlag = int

//
const QWidget__DrawWindowBackground QWidget__RenderFlag = 1

//
const QWidget__DrawChildren QWidget__RenderFlag = 2

//
const QWidget__IgnoreMask QWidget__RenderFlag = 4

func (this *QWidget) RenderFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWidget_RenderFlagItemName(val int) string {
	var nilthis *QWidget
	return nilthis.RenderFlagItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10169() {
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
