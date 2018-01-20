//  header block begin
// /usr/include/qt/QtWidgets/qtextbrowser.h
// #include <qtextbrowser.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 81
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
	return this.QTextEdit.GetCthis()
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTextBrowser) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:66
// index:0
// void QTextBrowser(class QWidget *)
func NewQTextBrowser(parent unsafe.Pointer) *QTextBrowser {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowserC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBrowserFromPointer(cthis)
	return gothis
}
func NewQTextBrowserFromPointer(cthis unsafe.Pointer) *QTextBrowser {
	bcthis0 := NewQTextEditFromPointer(cthis)
	return &QTextBrowser{bcthis0}
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:67
// index:0
// virtual
// void ~QTextBrowser()
func DeleteQTextBrowser(*QTextBrowser) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowserD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:69
// index:0
// QUrl source()
func (this *QTextBrowser) Source() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser6sourceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:71
// index:0
// QStringList searchPaths()
func (this *QTextBrowser) SearchPaths() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser11searchPathsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:72
// index:0
// void setSearchPaths(const class QStringList &)
func (this *QTextBrowser) SetSearchPaths(paths unsafe.Pointer) {
	// 0: (, paths const QStringList &), (paths)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser14setSearchPathsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), paths)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:74
// index:0
// virtual
// QVariant loadResource(int, const class QUrl &)
func (this *QTextBrowser) LoadResource(type_ int, name unsafe.Pointer) {
	// 0: (, type int, name const QUrl &), (&type_, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser12loadResourceEiRK4QUrl", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &type_, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:76
// index:0
// bool isBackwardAvailable()
func (this *QTextBrowser) IsBackwardAvailable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser19isBackwardAvailableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:77
// index:0
// bool isForwardAvailable()
func (this *QTextBrowser) IsForwardAvailable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser18isForwardAvailableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:78
// index:0
// void clearHistory()
func (this *QTextBrowser) ClearHistory() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser12clearHistoryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:79
// index:0
// QString historyTitle(int)
func (this *QTextBrowser) HistoryTitle(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser12historyTitleEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:80
// index:0
// QUrl historyUrl(int)
func (this *QTextBrowser) HistoryUrl(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser10historyUrlEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:81
// index:0
// int backwardHistoryCount()
func (this *QTextBrowser) BackwardHistoryCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser20backwardHistoryCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:82
// index:0
// int forwardHistoryCount()
func (this *QTextBrowser) ForwardHistoryCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser19forwardHistoryCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:84
// index:0
// bool openExternalLinks()
func (this *QTextBrowser) OpenExternalLinks() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser17openExternalLinksEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:85
// index:0
// void setOpenExternalLinks(_Bool)
func (this *QTextBrowser) SetOpenExternalLinks(open bool) {
	// 0: (, open bool), (&open)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser20setOpenExternalLinksEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &open)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:87
// index:0
// bool openLinks()
func (this *QTextBrowser) OpenLinks() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextBrowser9openLinksEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:88
// index:0
// void setOpenLinks(_Bool)
func (this *QTextBrowser) SetOpenLinks(open bool) {
	// 0: (, open bool), (&open)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser12setOpenLinksEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &open)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:91
// index:0
// virtual
// void setSource(const class QUrl &)
func (this *QTextBrowser) SetSource(name unsafe.Pointer) {
	// 0: (, name const QUrl &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser9setSourceERK4QUrl", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:92
// index:0
// virtual
// void backward()
func (this *QTextBrowser) Backward() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser8backwardEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:93
// index:0
// virtual
// void forward()
func (this *QTextBrowser) Forward() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser7forwardEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:94
// index:0
// virtual
// void home()
func (this *QTextBrowser) Home() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser4homeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:95
// index:0
// virtual
// void reload()
func (this *QTextBrowser) Reload() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser6reloadEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:98
// index:0
// void backwardAvailable(_Bool)
func (this *QTextBrowser) BackwardAvailable(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser17backwardAvailableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:99
// index:0
// void forwardAvailable(_Bool)
func (this *QTextBrowser) ForwardAvailable(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser16forwardAvailableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:100
// index:0
// void historyChanged()
func (this *QTextBrowser) HistoryChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser14historyChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:101
// index:0
// void sourceChanged(const class QUrl &)
func (this *QTextBrowser) SourceChanged(arg0 unsafe.Pointer) {
	// 0: (, const QUrl & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser13sourceChangedERK4QUrl", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:102
// index:0
// void highlighted(const class QUrl &)
func (this *QTextBrowser) Highlighted(arg0 unsafe.Pointer) {
	// 0: (, const QUrl & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser11highlightedERK4QUrl", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:103
// index:1
// void highlighted(const class QString &)
func (this *QTextBrowser) Highlighted_1(arg0 unsafe.Pointer) {
	// 1: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser11highlightedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:104
// index:0
// void anchorClicked(const class QUrl &)
func (this *QTextBrowser) AnchorClicked(arg0 unsafe.Pointer) {
	// 0: (, const QUrl & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser13anchorClickedERK4QUrl", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:107
// index:0
// virtual
// bool event(class QEvent *)
func (this *QTextBrowser) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:108
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QTextBrowser) KeyPressEvent(ev unsafe.Pointer) {
	// 0: (, ev QKeyEvent *), (ev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:109
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QTextBrowser) MouseMoveEvent(ev unsafe.Pointer) {
	// 0: (, ev QMouseEvent *), (ev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:110
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QTextBrowser) MousePressEvent(ev unsafe.Pointer) {
	// 0: (, ev QMouseEvent *), (ev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:111
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QTextBrowser) MouseReleaseEvent(ev unsafe.Pointer) {
	// 0: (, ev QMouseEvent *), (ev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:112
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QTextBrowser) FocusOutEvent(ev unsafe.Pointer) {
	// 0: (, ev QFocusEvent *), (ev)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), ev)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:113
// index:0
// virtual
// bool focusNextPrevChild(_Bool)
func (this *QTextBrowser) FocusNextPrevChild(next bool) {
	// 0: (, next bool), (&next)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser18focusNextPrevChildEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:114
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QTextBrowser) PaintEvent(e unsafe.Pointer) {
	// 0: (, e QPaintEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextBrowser10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

//  body block end
