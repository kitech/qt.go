package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h
// #include <qwebengineview.h>
// #include <QtWebEngineWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtprintsupport"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

//  ext block end

//  body block begin

// QWebEngineView * createWindow(QWebEnginePage::WebWindowType)
func (this *QWebEngineView) InheritCreateWindow(f func(type_ int) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createWindow", f)
}

// void contextMenuEvent(QContextMenuEvent *)
func (this *QWebEngineView) InheritContextMenuEvent(f func(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// bool event(QEvent *)
func (this *QWebEngineView) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void showEvent(QShowEvent *)
func (this *QWebEngineView) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QWebEngineView) InheritHideEvent(f func(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void dragEnterEvent(QDragEnterEvent *)
func (this *QWebEngineView) InheritDragEnterEvent(f func(e *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QWebEngineView) InheritDragLeaveEvent(f func(e *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dragMoveEvent(QDragMoveEvent *)
func (this *QWebEngineView) InheritDragMoveEvent(f func(e *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dropEvent(QDropEvent *)
func (this *QWebEngineView) InheritDropEvent(f func(e *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

/*

 */
type QWebEngineView struct {
	*qtwidgets.QWidget
}
type QWebEngineView_ITF interface {
	qtwidgets.QWidget_ITF
	QWebEngineView_PTR() *QWebEngineView
}

func (ptr *QWebEngineView) QWebEngineView_PTR() *QWebEngineView { return ptr }

func (this *QWebEngineView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QWebEngineView) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = qtwidgets.NewQWidgetFromPointer(cthis)
}
func NewQWebEngineViewFromPointer(cthis unsafe.Pointer) *QWebEngineView {
	bcthis0 := qtwidgets.NewQWidgetFromPointer(cthis)
	return &QWebEngineView{bcthis0}
}
func (*QWebEngineView) NewFromPointer(cthis unsafe.Pointer) *QWebEngineView {
	return NewQWebEngineViewFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWebEngineView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineView(QWidget *)

/*

 */
func NewQWebEngineView(parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QWebEngineView {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEngineView")
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineView(QWidget *)

/*

 */
func NewQWebEngineView__() *QWebEngineView {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineViewC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEngineView")
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWebEngineView()

/*

 */
func DeleteQWebEngineView(this *QWebEngineView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 56)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEnginePage * page() const

/*

 */
func (this *QWebEngineView) Page() *QWebEnginePage /*777 QWebEnginePage **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView4pageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEnginePageFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPage(QWebEnginePage *)

/*

 */
func (this *QWebEngineView) SetPage(page QWebEnginePage_ITF /*777 QWebEnginePage **/) {
	var convArg0 unsafe.Pointer
	if page != nil && page.QWebEnginePage_PTR() != nil {
		convArg0 = page.QWebEnginePage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView7setPageEP14QWebEnginePage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load(const QUrl &)

/*

 */
func (this *QWebEngineView) Load(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView4loadERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:76
// index:1
// Public Visibility=Default Availability=Available
// [-2] void load(const QWebEngineHttpRequest &)

/*

 */
func (this *QWebEngineView) Load_1(request qtwebenginecore.QWebEngineHttpRequest_ITF) {
	var convArg0 unsafe.Pointer
	if request != nil && request.QWebEngineHttpRequest_PTR() != nil {
		convArg0 = request.QWebEngineHttpRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView4loadERK21QWebEngineHttpRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &, const QUrl &)

/*

 */
func (this *QWebEngineView) SetHtml(html string, baseUrl qtcore.QUrl_ITF) {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if baseUrl != nil && baseUrl.QUrl_PTR() != nil {
		convArg1 = baseUrl.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView7setHtmlERK7QStringRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &, const QUrl &)

/*

 */
func (this *QWebEngineView) SetHtml__(html string) {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg1 = qtcore.NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView7setHtmlERK7QStringRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContent(const QByteArray &, const QString &, const QUrl &)

/*

 */
func (this *QWebEngineView) SetContent(data qtcore.QByteArray_ITF, mimeType string, baseUrl qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(mimeType)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if baseUrl != nil && baseUrl.QUrl_PTR() != nil {
		convArg2 = baseUrl.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView10setContentERK10QByteArrayRK7QStringRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContent(const QByteArray &, const QString &, const QUrl &)

/*

 */
func (this *QWebEngineView) SetContent__(data qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView10setContentERK10QByteArrayRK7QStringRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContent(const QByteArray &, const QString &, const QUrl &)

/*

 */
func (this *QWebEngineView) SetContent__1(data qtcore.QByteArray_ITF, mimeType string) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(mimeType)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView10setContentERK10QByteArrayRK7QStringRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineHistory * history() const

/*

 */
func (this *QWebEngineView) History() *QWebEngineHistory /*777 QWebEngineHistory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView7historyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineHistoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title() const

/*

 */
func (this *QWebEngineView) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QUrl &)

/*

 */
func (this *QWebEngineView) SetUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView6setUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*

 */
func (this *QWebEngineView) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl iconUrl() const

/*

 */
func (this *QWebEngineView) IconUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView7iconUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon() const

/*

 */
func (this *QWebEngineView) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSelection() const

/*

 */
func (this *QWebEngineView) HasSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView12hasSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedText() const

/*

 */
func (this *QWebEngineView) SelectedText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView12selectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * pageAction(QWebEnginePage::WebAction) const

/*

 */
func (this *QWebEngineView) PageAction(action int) *qtwidgets.QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView10pageActionEN14QWebEnginePage9WebActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
	return qtwidgets.NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggerPageAction(QWebEnginePage::WebAction, bool)

/*

 */
func (this *QWebEngineView) TriggerPageAction(action int, checked bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView17triggerPageActionEN14QWebEnginePage9WebActionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action, checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void triggerPageAction(QWebEnginePage::WebAction, bool)

/*

 */
func (this *QWebEngineView) TriggerPageAction__(action int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	checked := false
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView17triggerPageActionEN14QWebEnginePage9WebActionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action, checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal zoomFactor() const

/*

 */
func (this *QWebEngineView) ZoomFactor() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView10zoomFactorEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZoomFactor(qreal)

/*

 */
func (this *QWebEngineView) SetZoomFactor(factor float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView13setZoomFactorEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*

 */
func (this *QWebEngineView) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineSettings * settings() const

/*

 */
func (this *QWebEngineView) Settings() *QWebEngineSettings /*777 QWebEngineSettings **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEngineView8settingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*

 */
func (this *QWebEngineView) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void back()

/*

 */
func (this *QWebEngineView) Back() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void forward()

/*

 */
func (this *QWebEngineView) Forward() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView7forwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reload()

/*

 */
func (this *QWebEngineView) Reload() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView6reloadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadStarted()

/*

 */
func (this *QWebEngineView) LoadStarted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView11loadStartedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadProgress(int)

/*

 */
func (this *QWebEngineView) LoadProgress(progress int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView12loadProgressEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), progress)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadFinished(bool)

/*

 */
func (this *QWebEngineView) LoadFinished(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView12loadFinishedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void titleChanged(const QString &)

/*

 */
func (this *QWebEngineView) TitleChanged(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView12titleChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()

/*

 */
func (this *QWebEngineView) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void urlChanged(const QUrl &)

/*

 */
func (this *QWebEngineView) UrlChanged(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView10urlChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconUrlChanged(const QUrl &)

/*

 */
func (this *QWebEngineView) IconUrlChanged(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView14iconUrlChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconChanged(const QIcon &)

/*

 */
func (this *QWebEngineView) IconChanged(arg0 qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QIcon_PTR() != nil {
		convArg0 = arg0.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView11iconChangedERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void renderProcessTerminated(QWebEnginePage::RenderProcessTerminationStatus, int)

/*

 */
func (this *QWebEngineView) RenderProcessTerminated(terminationStatus int, exitCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView23renderProcessTerminatedEN14QWebEnginePage30RenderProcessTerminationStatusEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), terminationStatus, exitCode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:128
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QWebEngineView * createWindow(QWebEnginePage::WebWindowType)

/*

 */
func (this *QWebEngineView) CreateWindow(type_ int) *QWebEngineView /*777 QWebEngineView **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView12createWindowEN14QWebEnginePage13WebWindowTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:129
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)

/*

 */
func (this *QWebEngineView) ContextMenuEvent(arg0 qtgui.QContextMenuEvent_ITF /*777 QContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QContextMenuEvent_PTR() != nil {
		convArg0 = arg0.QContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:130
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*

 */
func (this *QWebEngineView) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:131
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*

 */
func (this *QWebEngineView) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:132
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*

 */
func (this *QWebEngineView) HideEvent(arg0 qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QHideEvent_PTR() != nil {
		convArg0 = arg0.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:133
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*

 */
func (this *QWebEngineView) DragEnterEvent(e qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragEnterEvent_PTR() != nil {
		convArg0 = e.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:134
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*

 */
func (this *QWebEngineView) DragLeaveEvent(e qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragLeaveEvent_PTR() != nil {
		convArg0 = e.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:135
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*

 */
func (this *QWebEngineView) DragMoveEvent(e qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDragMoveEvent_PTR() != nil {
		convArg0 = e.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebengineview.h:136
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*

 */
func (this *QWebEngineView) DropEvent(e qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QDropEvent_PTR() != nil {
		convArg0 = e.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEngineView9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtquickwidgets.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtwidgets.KeepMe()
	}
	if false {
		qtprintsupport.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
	if false {
		qtwebenginecore.KeepMe()
	}
}

//  keep block end
