package qtwidgets

// /usr/include/qt/QtWidgets/qtabbar.h
// #include <qtabbar.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// QSize tabSizeHint(int)
func (this *QTabBar) InheritTabSizeHint(f func(index int) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "tabSizeHint", f)
}

// QSize minimumTabSizeHint(int)
func (this *QTabBar) InheritMinimumTabSizeHint(f func(index int) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "minimumTabSizeHint", f)
}

// void tabInserted(int)
func (this *QTabBar) InheritTabInserted(f func(index int)) {
	ffiqt.SetAllInheritCallback(this, "tabInserted", f)
}

// void tabRemoved(int)
func (this *QTabBar) InheritTabRemoved(f func(index int)) {
	ffiqt.SetAllInheritCallback(this, "tabRemoved", f)
}

// void tabLayoutChange()
func (this *QTabBar) InheritTabLayoutChange(f func()) {
	ffiqt.SetAllInheritCallback(this, "tabLayoutChange", f)
}

// bool event(class QEvent *)
func (this *QTabBar) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QTabBar) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QTabBar) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(class QHideEvent *)
func (this *QTabBar) InheritHideEvent(f func(arg0 *qtgui.QHideEvent /*777 QHideEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "hideEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QTabBar) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QTabBar) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QTabBar) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QTabBar) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QTabBar) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QTabBar) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QTabBar) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "changeEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QTabBar) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "timerEvent", f)
}

// void initStyleOption(class QStyleOptionTab *, int)
func (this *QTabBar) InheritInitStyleOption(f func(option *QStyleOptionTab /*777 QStyleOptionTab **/, tabIndex int)) {
	ffiqt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QTabBar struct {
	*QWidget
}

func (this *QTabBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QTabBar) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQTabBarFromPointer(cthis unsafe.Pointer) *QTabBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QTabBar{bcthis0}
}
func (*QTabBar) NewFromPointer(cthis unsafe.Pointer) *QTabBar {
	return NewQTabBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtabbar.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTabBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabBar(QWidget *)
func NewQTabBar(parent *QWidget /*777 QWidget **/) *QTabBar {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBarC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTabBarFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qtabbar.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTabBar()
func DeleteQTabBar(this *QTabBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBarD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtabbar.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabBar::Shape shape()
func (this *QTabBar) Shape() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShape(enum QTabBar::Shape)
func (this *QTabBar) SetShape(shape int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar8setShapeENS_5ShapeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), shape)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int addTab(const QString &)
func (this *QTabBar) AddTab(text *qtcore.QString) int {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar6addTabERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:97
// index:1
// Public Visibility=Default Availability=Available
// [4] int addTab(const QIcon &, const QString &)
func (this *QTabBar) AddTab_1(icon *qtgui.QIcon, text *qtcore.QString) int {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar6addTabERK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertTab(int, const QString &)
func (this *QTabBar) InsertTab(index int, text *qtcore.QString) int {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar9insertTabEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:100
// index:1
// Public Visibility=Default Availability=Available
// [4] int insertTab(int, const QIcon &, const QString &)
func (this *QTabBar) InsertTab_1(index int, icon *qtgui.QIcon, text *qtcore.QString) int {
	var convArg1 = icon.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar9insertTabEiRK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeTab(int)
func (this *QTabBar) RemoveTab(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar9removeTabEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveTab(int, int)
func (this *QTabBar) MoveTab(from int, to int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar7moveTabEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), from, to)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTabEnabled(int)
func (this *QTabBar) IsTabEnabled(index int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12isTabEnabledEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabEnabled(int, _Bool)
func (this *QTabBar) SetTabEnabled(index int, arg1 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar13setTabEnabledEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabText(int)
func (this *QTabBar) TabText(index int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar7tabTextEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabText(int, const QString &)
func (this *QTabBar) SetTabText(index int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10setTabTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:111
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor tabTextColor(int)
func (this *QTabBar) TabTextColor(index int) *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12tabTextColorEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabTextColor(int, const QColor &)
func (this *QTabBar) SetTabTextColor(index int, color *qtgui.QColor) {
	var convArg1 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15setTabTextColorEiRK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon tabIcon(int)
func (this *QTabBar) TabIcon(index int) *qtgui.QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar7tabIconEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabIcon(int, const QIcon &)
func (this *QTabBar) SetTabIcon(index int, icon *qtgui.QIcon) {
	var convArg1 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10setTabIconEiRK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextElideMode elideMode()
func (this *QTabBar) ElideMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar9elideModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setElideMode(Qt::TextElideMode)
func (this *QTabBar) SetElideMode(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar12setElideModeEN2Qt13TextElideModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabToolTip(int, const QString &)
func (this *QTabBar) SetTabToolTip(index int, tip *qtcore.QString) {
	var convArg1 = tip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar13setTabToolTipEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabToolTip(int)
func (this *QTabBar) TabToolTip(index int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar10tabToolTipEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabWhatsThis(int, const QString &)
func (this *QTabBar) SetTabWhatsThis(index int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15setTabWhatsThisEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabWhatsThis(int)
func (this *QTabBar) TabWhatsThis(index int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12tabWhatsThisEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabData(int, const QVariant &)
func (this *QTabBar) SetTabData(index int, data *qtcore.QVariant) {
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10setTabDataEiRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:131
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant tabData(int)
func (this *QTabBar) TabData(index int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar7tabDataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:133
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect tabRect(int)
func (this *QTabBar) TabRect(index int) *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar7tabRectEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int tabAt(const QPoint &)
func (this *QTabBar) TabAt(pos *qtcore.QPoint) int {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar5tabAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex()
func (this *QTabBar) CurrentIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QTabBar) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabbar.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QTabBar) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:140
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QTabBar) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDrawBase(_Bool)
func (this *QTabBar) SetDrawBase(drawTheBase bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar11setDrawBaseEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), drawTheBase)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool drawBase()
func (this *QTabBar) DrawBase() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar8drawBaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize()
func (this *QTabBar) IconSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar8iconSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)
func (this *QTabBar) SetIconSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar11setIconSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool usesScrollButtons()
func (this *QTabBar) UsesScrollButtons() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar17usesScrollButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUsesScrollButtons(_Bool)
func (this *QTabBar) SetUsesScrollButtons(useButtons bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar20setUsesScrollButtonsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), useButtons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabsClosable()
func (this *QTabBar) TabsClosable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12tabsClosableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabsClosable(_Bool)
func (this *QTabBar) SetTabsClosable(closable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15setTabsClosableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), closable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabButton(int, enum QTabBar::ButtonPosition, QWidget *)
func (this *QTabBar) SetTabButton(index int, position int, widget *QWidget /*777 QWidget **/) {
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar12setTabButtonEiNS_14ButtonPositionEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, position, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * tabButton(int, enum QTabBar::ButtonPosition)
func (this *QTabBar) TabButton(index int, position int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar9tabButtonEiNS_14ButtonPositionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, position)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabBar::SelectionBehavior selectionBehaviorOnRemove()
func (this *QTabBar) SelectionBehaviorOnRemove() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar25selectionBehaviorOnRemoveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionBehaviorOnRemove(enum QTabBar::SelectionBehavior)
func (this *QTabBar) SetSelectionBehaviorOnRemove(behavior int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar28setSelectionBehaviorOnRemoveENS_17SelectionBehaviorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), behavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool expanding()
func (this *QTabBar) Expanding() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar9expandingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpanding(_Bool)
func (this *QTabBar) SetExpanding(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar12setExpandingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMovable()
func (this *QTabBar) IsMovable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar9isMovableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovable(_Bool)
func (this *QTabBar) SetMovable(movable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10setMovableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool documentMode()
func (this *QTabBar) DocumentMode() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar12documentModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMode(_Bool)
func (this *QTabBar) SetDocumentMode(set bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15setDocumentModeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), set)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:169
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoHide()
func (this *QTabBar) AutoHide() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar8autoHideEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoHide(_Bool)
func (this *QTabBar) SetAutoHide(hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar11setAutoHideEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:172
// index:0
// Public Visibility=Default Availability=Available
// [1] bool changeCurrentOnDrag()
func (this *QTabBar) ChangeCurrentOnDrag() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar19changeCurrentOnDragEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChangeCurrentOnDrag(_Bool)
func (this *QTabBar) SetChangeCurrentOnDrag(change bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar22setChangeCurrentOnDragEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), change)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:176
// index:0
// Public Visibility=Default Availability=Available
// [8] QString accessibleTabName(int)
func (this *QTabBar) AccessibleTabName(index int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar17accessibleTabNameEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccessibleTabName(int, const QString &)
func (this *QTabBar) SetAccessibleTabName(index int, name *qtcore.QString) {
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar20setAccessibleTabNameEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)
func (this *QTabBar) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(int)
func (this *QTabBar) CurrentChanged(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar14currentChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabCloseRequested(int)
func (this *QTabBar) TabCloseRequested(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar17tabCloseRequestedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabMoved(int, int)
func (this *QTabBar) TabMoved(from int, to int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar8tabMovedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), from, to)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabBarClicked(int)
func (this *QTabBar) TabBarClicked(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar13tabBarClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabBarDoubleClicked(int)
func (this *QTabBar) TabBarDoubleClicked(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar19tabBarDoubleClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:191
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize tabSizeHint(int)
func (this *QTabBar) TabSizeHint(index int) *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar11tabSizeHintEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:192
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize minimumTabSizeHint(int)
func (this *QTabBar) MinimumTabSizeHint(index int) *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar18minimumTabSizeHintEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabbar.h:193
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabInserted(int)
func (this *QTabBar) TabInserted(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar11tabInsertedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:194
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabRemoved(int)
func (this *QTabBar) TabRemoved(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10tabRemovedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:195
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabLayoutChange()
func (this *QTabBar) TabLayoutChange() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15tabLayoutChangeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:197
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QTabBar) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabbar.h:198
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QTabBar) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:199
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QTabBar) ShowEvent(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:200
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QTabBar) HideEvent(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:201
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QTabBar) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:202
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QTabBar) MousePressEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:203
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QTabBar) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:204
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QTabBar) MouseReleaseEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:206
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QTabBar) WheelEvent(event *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QTabBar) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:209
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QTabBar) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:210
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QTabBar) TimerEvent(event *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QTabBar10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabbar.h:211
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionTab *, int)
func (this *QTabBar) InitStyleOption(option *QStyleOptionTab /*777 QStyleOptionTab **/, tabIndex int) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QTabBar15initStyleOptionEP15QStyleOptionTabi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, tabIndex)
	gopp.ErrPrint(err, rv)
}

type QTabBar__Shape = int

const QTabBar__RoundedNorth QTabBar__Shape = 0
const QTabBar__RoundedSouth QTabBar__Shape = 1
const QTabBar__RoundedWest QTabBar__Shape = 2
const QTabBar__RoundedEast QTabBar__Shape = 3
const QTabBar__TriangularNorth QTabBar__Shape = 4
const QTabBar__TriangularSouth QTabBar__Shape = 5
const QTabBar__TriangularWest QTabBar__Shape = 6
const QTabBar__TriangularEast QTabBar__Shape = 7

type QTabBar__ButtonPosition = int

const QTabBar__LeftSide QTabBar__ButtonPosition = 0
const QTabBar__RightSide QTabBar__ButtonPosition = 1

type QTabBar__SelectionBehavior = int

const QTabBar__SelectLeftTab QTabBar__SelectionBehavior = 0
const QTabBar__SelectRightTab QTabBar__SelectionBehavior = 1
const QTabBar__SelectPreviousTab QTabBar__SelectionBehavior = 2

//  body block end
