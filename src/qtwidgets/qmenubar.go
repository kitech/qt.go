//  header block begin
// /usr/include/qt/QtWidgets/qmenubar.h
// #include <qmenubar.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 52
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
type QMenuBar struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qmenubar.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMenuBar) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:63
// index:0
// void QMenuBar(class QWidget *)
func NewQMenuBar(parent unsafe.Pointer) *QMenuBar {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QMenuBar{cthis}
}

// /usr/include/qt/QtWidgets/qmenubar.h:64
// index:0
// virtual
// void ~QMenuBar()
func DeleteQMenuBar(*QMenuBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:67
// index:0
// QAction * addAction(const class QString &)
func (this *QMenuBar) AddAction(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar9addActionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:68
// index:1
// QAction * addAction(const class QString &, const class QObject *, const char *)
func (this *QMenuBar) AddAction_1(text unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 1: (, const QString & text, const QObject * receiver, const char * member), (text, receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.cthis, text, receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:70
// index:0
// QAction * addMenu(class QMenu *)
func (this *QMenuBar) AddMenu(menu unsafe.Pointer) {
	// 0: (, QMenu * menu), (menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7addMenuEP5QMenu", ffiqt.FFI_TYPE_VOID, this.cthis, menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:71
// index:1
// QMenu * addMenu(const class QString &)
func (this *QMenuBar) AddMenu_1(title unsafe.Pointer) {
	// 1: (, const QString & title), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7addMenuERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:72
// index:2
// QMenu * addMenu(const class QIcon &, const class QString &)
func (this *QMenuBar) AddMenu_2(icon unsafe.Pointer, title unsafe.Pointer) {
	// 2: (, const QIcon & icon, const QString & title), (icon, title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7addMenuERK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, icon, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:75
// index:0
// QAction * addSeparator()
func (this *QMenuBar) AddSeparator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar12addSeparatorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:76
// index:0
// QAction * insertSeparator(class QAction *)
func (this *QMenuBar) InsertSeparator(before unsafe.Pointer) {
	// 0: (, QAction * before), (before)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15insertSeparatorEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, before)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:78
// index:0
// QAction * insertMenu(class QAction *, class QMenu *)
func (this *QMenuBar) InsertMenu(before unsafe.Pointer, menu unsafe.Pointer) {
	// 0: (, QAction * before, QMenu * menu), (before, menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10insertMenuEP7QActionP5QMenu", ffiqt.FFI_TYPE_VOID, this.cthis, before, menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:80
// index:0
// void clear()
func (this *QMenuBar) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:82
// index:0
// QAction * activeAction()
func (this *QMenuBar) ActiveAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar12activeActionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:83
// index:0
// void setActiveAction(class QAction *)
func (this *QMenuBar) SetActiveAction(action unsafe.Pointer) {
	// 0: (, QAction * action), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15setActiveActionEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:85
// index:0
// void setDefaultUp(_Bool)
func (this *QMenuBar) SetDefaultUp(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar12setDefaultUpEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:86
// index:0
// bool isDefaultUp()
func (this *QMenuBar) IsDefaultUp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar11isDefaultUpEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:88
// index:0
// virtual
// QSize sizeHint()
func (this *QMenuBar) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:89
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QMenuBar) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:90
// index:0
// virtual
// int heightForWidth(int)
func (this *QMenuBar) HeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:92
// index:0
// QRect actionGeometry(class QAction *)
func (this *QMenuBar) ActionGeometry(arg0 unsafe.Pointer) {
	// 0: (, QAction * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar14actionGeometryEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:93
// index:0
// QAction * actionAt(const class QPoint &)
func (this *QMenuBar) ActionAt(arg0 unsafe.Pointer) {
	// 0: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar8actionAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:95
// index:0
// void setCornerWidget(class QWidget *, Qt::Corner)
func (this *QMenuBar) SetCornerWidget(w unsafe.Pointer, corner int) {
	// 0: (, QWidget * w, Qt::Corner corner), (w, &corner)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15setCornerWidgetEP7QWidgetN2Qt6CornerE", ffiqt.FFI_TYPE_VOID, this.cthis, w, &corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:96
// index:0
// QWidget * cornerWidget(Qt::Corner)
func (this *QMenuBar) CornerWidget(corner int) {
	// 0: (, Qt::Corner corner), (&corner)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar12cornerWidgetEN2Qt6CornerE", ffiqt.FFI_TYPE_VOID, this.cthis, &corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:102
// index:0
// bool isNativeMenuBar()
func (this *QMenuBar) IsNativeMenuBar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMenuBar15isNativeMenuBarEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:103
// index:0
// void setNativeMenuBar(_Bool)
func (this *QMenuBar) SetNativeMenuBar(nativeMenuBar bool) {
	// 0: (, bool nativeMenuBar), (&nativeMenuBar)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar16setNativeMenuBarEb", ffiqt.FFI_TYPE_VOID, this.cthis, &nativeMenuBar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:104
// index:0
// QPlatformMenuBar * platformMenuBar()
func (this *QMenuBar) PlatformMenuBar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar15platformMenuBarEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:106
// index:0
// virtual
// void setVisible(_Bool)
func (this *QMenuBar) SetVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:109
// index:0
// void triggered(class QAction *)
func (this *QMenuBar) Triggered(action unsafe.Pointer) {
	// 0: (, QAction * action), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar9triggeredEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmenubar.h:110
// index:0
// void hovered(class QAction *)
func (this *QMenuBar) Hovered(action unsafe.Pointer) {
	// 0: (, QAction * action), (action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMenuBar7hoveredEP7QAction", ffiqt.FFI_TYPE_VOID, this.cthis, action)
	gopp.ErrPrint(err, rv)
}

//  body block end
