package qtwidgets

// /usr/include/qt/QtWidgets/qtextbrowser.h
// #include <qtextbrowser.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 76
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
type QTextBrowser struct {
	*QTextEdit
}

func (this *QTextBrowser) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextEdit.GetCthis()
	}
}
func (this *QTextBrowser) SetCthis(cthis unsafe.Pointer) {
	this.QTextEdit = NewQTextEditFromPointer(cthis)
}
func NewQTextBrowserFromPointer(cthis unsafe.Pointer) *QTextBrowser {
	bcthis0 := NewQTextEditFromPointer(cthis)
	return &QTextBrowser{bcthis0}
}
func (*QTextBrowser) NewFromPointer(cthis unsafe.Pointer) *QTextBrowser {
	return NewQTextBrowserFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTextBrowser) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:66
// index:0
// Public
// void QTextBrowser(class QWidget *)
func NewQTextBrowser(parent *QWidget /*777 QWidget **/) *QTextBrowser {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowserC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBrowserFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:67
// index:0
// Public virtual
// void ~QTextBrowser()
func DeleteQTextBrowser(*QTextBrowser) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowserD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:69
// index:0
// Public
// QUrl source()
func (this *QTextBrowser) Source() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser6sourceEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:72
// index:0
// Public
// void setSearchPaths(const class QStringList &)
func (this *QTextBrowser) SetSearchPaths(paths *qtcore.QStringList) {
	var convArg0 = paths.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser14setSearchPathsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:74
// index:0
// Public virtual
// QVariant loadResource(int, const class QUrl &)
func (this *QTextBrowser) LoadResource(type_ int, name *qtcore.QUrl) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser12loadResourceEiRK4QUrl", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), type_, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:76
// index:0
// Public
// bool isBackwardAvailable()
func (this *QTextBrowser) IsBackwardAvailable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser19isBackwardAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:77
// index:0
// Public
// bool isForwardAvailable()
func (this *QTextBrowser) IsForwardAvailable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser18isForwardAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:78
// index:0
// Public
// void clearHistory()
func (this *QTextBrowser) ClearHistory() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser12clearHistoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:79
// index:0
// Public
// QString historyTitle(int)
func (this *QTextBrowser) HistoryTitle(arg0 int) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser12historyTitleEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:80
// index:0
// Public
// QUrl historyUrl(int)
func (this *QTextBrowser) HistoryUrl(arg0 int) *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser10historyUrlEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:81
// index:0
// Public
// int backwardHistoryCount()
func (this *QTextBrowser) BackwardHistoryCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser20backwardHistoryCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:82
// index:0
// Public
// int forwardHistoryCount()
func (this *QTextBrowser) ForwardHistoryCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser19forwardHistoryCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:84
// index:0
// Public
// bool openExternalLinks()
func (this *QTextBrowser) OpenExternalLinks() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser17openExternalLinksEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:85
// index:0
// Public
// void setOpenExternalLinks(_Bool)
func (this *QTextBrowser) SetOpenExternalLinks(open bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser20setOpenExternalLinksEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), open)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:87
// index:0
// Public
// bool openLinks()
func (this *QTextBrowser) OpenLinks() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser9openLinksEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:88
// index:0
// Public
// void setOpenLinks(_Bool)
func (this *QTextBrowser) SetOpenLinks(open bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser12setOpenLinksEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), open)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:91
// index:0
// Public virtual
// void setSource(const class QUrl &)
func (this *QTextBrowser) SetSource(name *qtcore.QUrl) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser9setSourceERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:92
// index:0
// Public virtual
// void backward()
func (this *QTextBrowser) Backward() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser8backwardEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:93
// index:0
// Public virtual
// void forward()
func (this *QTextBrowser) Forward() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser7forwardEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:94
// index:0
// Public virtual
// void home()
func (this *QTextBrowser) Home() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser4homeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:95
// index:0
// Public virtual
// void reload()
func (this *QTextBrowser) Reload() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser6reloadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:98
// index:0
// Public
// void backwardAvailable(_Bool)
func (this *QTextBrowser) BackwardAvailable(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser17backwardAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:99
// index:0
// Public
// void forwardAvailable(_Bool)
func (this *QTextBrowser) ForwardAvailable(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser16forwardAvailableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:100
// index:0
// Public
// void historyChanged()
func (this *QTextBrowser) HistoryChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser14historyChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:101
// index:0
// Public
// void sourceChanged(const class QUrl &)
func (this *QTextBrowser) SourceChanged(arg0 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser13sourceChangedERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:102
// index:0
// Public
// void highlighted(const class QUrl &)
func (this *QTextBrowser) Highlighted(arg0 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser11highlightedERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:103
// index:1
// Public
// void highlighted(const class QString &)
func (this *QTextBrowser) Highlighted_1(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser11highlightedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:104
// index:0
// Public
// void anchorClicked(const class QUrl &)
func (this *QTextBrowser) AnchorClicked(arg0 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser13anchorClickedERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:107
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QTextBrowser) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:108
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QTextBrowser) KeyPressEvent(ev *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:109
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QTextBrowser) MouseMoveEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:110
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QTextBrowser) MousePressEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:111
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QTextBrowser) MouseReleaseEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:112
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QTextBrowser) FocusOutEvent(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:113
// index:0
// Protected virtual
// bool focusNextPrevChild(_Bool)
func (this *QTextBrowser) FocusNextPrevChild(next bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:114
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QTextBrowser) PaintEvent(e *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
