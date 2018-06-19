package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h
// #include <qwebenginepage.h>
// #include <QtWebEngineWidgets>

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

// QWebEnginePage * createWindow(QWebEnginePage::WebWindowType)
func (this *QWebEnginePage) InheritCreateWindow(f func(type_ int) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createWindow", f)
}

// QStringList chooseFiles(QWebEnginePage::FileSelectionMode, const QStringList &, const QStringList &)
func (this *QWebEnginePage) InheritChooseFiles(f func(mode int, oldFiles *qtcore.QStringList, acceptedMimeTypes *qtcore.QStringList) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "chooseFiles", f)
}

// void javaScriptAlert(const QUrl &, const QString &)
func (this *QWebEnginePage) InheritJavaScriptAlert(f func(securityOrigin *qtcore.QUrl, msg string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "javaScriptAlert", f)
}

// bool javaScriptConfirm(const QUrl &, const QString &)
func (this *QWebEnginePage) InheritJavaScriptConfirm(f func(securityOrigin *qtcore.QUrl, msg string) bool) {
	qtrt.SetAllInheritCallback(this, "javaScriptConfirm", f)
}

// bool javaScriptPrompt(const QUrl &, const QString &, const QString &, QString *)
func (this *QWebEnginePage) InheritJavaScriptPrompt(f func(securityOrigin *qtcore.QUrl, msg string, defaultValue string, result string) bool) {
	qtrt.SetAllInheritCallback(this, "javaScriptPrompt", f)
}

// void javaScriptConsoleMessage(QWebEnginePage::JavaScriptConsoleMessageLevel, const QString &, int, const QString &)
func (this *QWebEnginePage) InheritJavaScriptConsoleMessage(f func(level int, message string, lineNumber int, sourceID string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "javaScriptConsoleMessage", f)
}

// bool certificateError(const QWebEngineCertificateError &)
func (this *QWebEnginePage) InheritCertificateError(f func(certificateError *QWebEngineCertificateError) bool) {
	qtrt.SetAllInheritCallback(this, "certificateError", f)
}

// bool acceptNavigationRequest(const QUrl &, QWebEnginePage::NavigationType, bool)
func (this *QWebEnginePage) InheritAcceptNavigationRequest(f func(url *qtcore.QUrl, type_ int, isMainFrame bool) bool) {
	qtrt.SetAllInheritCallback(this, "acceptNavigationRequest", f)
}

/*

 */
type QWebEnginePage struct {
	*qtcore.QObject
}
type QWebEnginePage_ITF interface {
	qtcore.QObject_ITF
	QWebEnginePage_PTR() *QWebEnginePage
}

func (ptr *QWebEnginePage) QWebEnginePage_PTR() *QWebEnginePage { return ptr }

func (this *QWebEnginePage) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWebEnginePage) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWebEnginePageFromPointer(cthis unsafe.Pointer) *QWebEnginePage {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWebEnginePage{bcthis0}
}
func (*QWebEnginePage) NewFromPointer(cthis unsafe.Pointer) *QWebEnginePage {
	return NewQWebEnginePageFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWebEnginePage) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEnginePage(QObject *)

/*

 */
func NewQWebEnginePage(parent qtcore.QObject_ITF /*777 QObject **/) *QWebEnginePage {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePageC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEnginePageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEnginePage")
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEnginePage(QObject *)

/*

 */
func NewQWebEnginePage__() *QWebEnginePage {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePageC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEnginePageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEnginePage")
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:223
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QWebEnginePage(QWebEngineProfile *, QObject *)

/*

 */
func NewQWebEnginePage_1(profile QWebEngineProfile_ITF /*777 QWebEngineProfile **/, parent qtcore.QObject_ITF /*777 QObject **/) *QWebEnginePage {
	var convArg0 unsafe.Pointer
	if profile != nil && profile.QWebEngineProfile_PTR() != nil {
		convArg0 = profile.QWebEngineProfile_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePageC2EP17QWebEngineProfileP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEnginePageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEnginePage")
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:223
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QWebEnginePage(QWebEngineProfile *, QObject *)

/*

 */
func NewQWebEnginePage_1_(profile QWebEngineProfile_ITF /*777 QWebEngineProfile **/) *QWebEnginePage {
	var convArg0 unsafe.Pointer
	if profile != nil && profile.QWebEngineProfile_PTR() != nil {
		convArg0 = profile.QWebEngineProfile_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePageC2EP17QWebEngineProfileP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEnginePageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWebEnginePage")
	return gothis
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:224
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWebEnginePage()

/*

 */
func DeleteQWebEnginePage(this *QWebEnginePage) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePageD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:225
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineHistory * history() const

/*

 */
func (this *QWebEnginePage) History() *QWebEngineHistory /*777 QWebEngineHistory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage7historyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineHistoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setView(QWidget *)

/*

 */
func (this *QWebEnginePage) SetView(view qtwidgets.QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if view != nil && view.QWidget_PTR() != nil {
		convArg0 = view.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage7setViewEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * view() const

/*

 */
func (this *QWebEnginePage) View() *qtwidgets.QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage4viewEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtwidgets.NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:230
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSelection() const

/*

 */
func (this *QWebEnginePage) HasSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage12hasSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:231
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedText() const

/*

 */
func (this *QWebEnginePage) SelectedText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage12selectedTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:233
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineProfile * profile() const

/*

 */
func (this *QWebEnginePage) Profile() *QWebEngineProfile /*777 QWebEngineProfile **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage7profileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineProfileFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:236
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * action(QWebEnginePage::WebAction) const

/*

 */
func (this *QWebEnginePage) Action(action int) *qtwidgets.QAction /*777 QAction **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage6actionENS_9WebActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action)
	qtrt.ErrPrint(err, rv)
	return qtwidgets.NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:238
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void triggerAction(QWebEnginePage::WebAction, bool)

/*

 */
func (this *QWebEnginePage) TriggerAction(action int, checked bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage13triggerActionENS_9WebActionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action, checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:238
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void triggerAction(QWebEnginePage::WebAction, bool)

/*

 */
func (this *QWebEnginePage) TriggerAction__(action int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	checked := false
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage13triggerActionENS_9WebActionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), action, checked)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void replaceMisspelledWord(const QString &)

/*

 */
func (this *QWebEnginePage) ReplaceMisspelledWord(replacement string) {
	var tmpArg0 = qtcore.NewQString_5(replacement)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage21replaceMisspelledWordERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:242
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*

 */
func (this *QWebEnginePage) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:249
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * createStandardContextMenu()

/*

 */
func (this *QWebEnginePage) CreateStandardContextMenu() *qtwidgets.QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage25createStandardContextMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtwidgets.NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:251
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFeaturePermission(const QUrl &, QWebEnginePage::Feature, QWebEnginePage::PermissionPolicy)

/*

 */
func (this *QWebEnginePage) SetFeaturePermission(securityOrigin qtcore.QUrl_ITF, feature int, policy int) {
	var convArg0 unsafe.Pointer
	if securityOrigin != nil && securityOrigin.QUrl_PTR() != nil {
		convArg0 = securityOrigin.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage20setFeaturePermissionERK4QUrlNS_7FeatureENS_16PermissionPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, feature, policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:253
// index:0
// Public Visibility=Default Availability=Available
// [-2] void load(const QUrl &)

/*

 */
func (this *QWebEnginePage) Load(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage4loadERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:254
// index:1
// Public Visibility=Default Availability=Available
// [-2] void load(const QWebEngineHttpRequest &)

/*

 */
func (this *QWebEnginePage) Load_1(request qtwebenginecore.QWebEngineHttpRequest_ITF) {
	var convArg0 unsafe.Pointer
	if request != nil && request.QWebEngineHttpRequest_PTR() != nil {
		convArg0 = request.QWebEngineHttpRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage4loadERK21QWebEngineHttpRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void download(const QUrl &, const QString &)

/*

 */
func (this *QWebEnginePage) Download(url qtcore.QUrl_ITF, filename string) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(filename)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage8downloadERK4QUrlRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void download(const QUrl &, const QString &)

/*

 */
func (this *QWebEnginePage) Download__(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage8downloadERK4QUrlRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &, const QUrl &)

/*

 */
func (this *QWebEnginePage) SetHtml(html string, baseUrl qtcore.QUrl_ITF) {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if baseUrl != nil && baseUrl.QUrl_PTR() != nil {
		convArg1 = baseUrl.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage7setHtmlERK7QStringRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &, const QUrl &)

/*

 */
func (this *QWebEnginePage) SetHtml__(html string) {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg1 = qtcore.NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage7setHtmlERK7QStringRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContent(const QByteArray &, const QString &, const QUrl &)

/*

 */
func (this *QWebEnginePage) SetContent(data qtcore.QByteArray_ITF, mimeType string, baseUrl qtcore.QUrl_ITF) {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage10setContentERK10QByteArrayRK7QStringRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContent(const QByteArray &, const QString &, const QUrl &)

/*

 */
func (this *QWebEnginePage) SetContent__(data qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage10setContentERK10QByteArrayRK7QStringRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContent(const QByteArray &, const QString &, const QUrl &)

/*

 */
func (this *QWebEnginePage) SetContent__1(data qtcore.QByteArray_ITF, mimeType string) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(mimeType)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage10setContentERK10QByteArrayRK7QStringRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:267
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title() const

/*

 */
func (this *QWebEnginePage) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QUrl &)

/*

 */
func (this *QWebEnginePage) SetUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage6setUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:269
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*

 */
func (this *QWebEnginePage) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:270
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl requestedUrl() const

/*

 */
func (this *QWebEnginePage) RequestedUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage12requestedUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:271
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl iconUrl() const

/*

 */
func (this *QWebEnginePage) IconUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage7iconUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:272
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon() const

/*

 */
func (this *QWebEnginePage) Icon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:274
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal zoomFactor() const

/*

 */
func (this *QWebEnginePage) ZoomFactor() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage10zoomFactorEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZoomFactor(qreal)

/*

 */
func (this *QWebEnginePage) SetZoomFactor(factor float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage13setZoomFactorEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:277
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scrollPosition() const

/*

 */
func (this *QWebEnginePage) ScrollPosition() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage14scrollPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:278
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF contentsSize() const

/*

 */
func (this *QWebEnginePage) ContentsSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage12contentsSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:280
// index:0
// Public Visibility=Default Availability=Available
// [-2] void runJavaScript(const QString &)

/*

 */
func (this *QWebEnginePage) RunJavaScript(scriptSource string) {
	var tmpArg0 = qtcore.NewQString_5(scriptSource)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage13runJavaScriptERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:281
// index:1
// Public Visibility=Default Availability=Available
// [-2] void runJavaScript(const QString &, quint32)

/*

 */
func (this *QWebEnginePage) RunJavaScript_1(scriptSource string, worldId uint) {
	var tmpArg0 = qtcore.NewQString_5(scriptSource)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage13runJavaScriptERK7QStringj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, worldId)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:289
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineScriptCollection & scripts()

/*

 */
func (this *QWebEnginePage) Scripts() *QWebEngineScriptCollection {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage7scriptsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineScriptCollectionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineScriptCollection)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:290
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineSettings * settings() const

/*

 */
func (this *QWebEnginePage) Settings() *QWebEngineSettings /*777 QWebEngineSettings **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage8settingsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:292
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebChannel * webChannel() const

/*

 */
func (this *QWebEnginePage) WebChannel() *qtwebchannel.QWebChannel /*777 QWebChannel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage10webChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtwebchannel.NewQWebChannelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWebChannel(QWebChannel *)

/*

 */
func (this *QWebEnginePage) SetWebChannel(arg0 qtwebchannel.QWebChannel_ITF /*777 QWebChannel **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWebChannel_PTR() != nil {
		convArg0 = arg0.QWebChannel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage13setWebChannelEP11QWebChannel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:294
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setWebChannel(QWebChannel *, uint)

/*

 */
func (this *QWebEnginePage) SetWebChannel_1(arg0 qtwebchannel.QWebChannel_ITF /*777 QWebChannel **/, worldId uint) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWebChannel_PTR() != nil {
		convArg0 = arg0.QWebChannel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage13setWebChannelEP11QWebChannelj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, worldId)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:295
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor backgroundColor() const

/*

 */
func (this *QWebEnginePage) BackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage15backgroundColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:296
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundColor(const QColor &)

/*

 */
func (this *QWebEnginePage) SetBackgroundColor(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage18setBackgroundColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:298
// index:0
// Public Visibility=Default Availability=Available
// [-2] void save(const QString &, QWebEngineDownloadItem::SavePageFormat) const

/*

 */
func (this *QWebEnginePage) Save(filePath string, format int) {
	var tmpArg0 = qtcore.NewQString_5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage4saveERK7QStringN22QWebEngineDownloadItem14SavePageFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:298
// index:0
// Public Visibility=Default Availability=Available
// [-2] void save(const QString &, QWebEngineDownloadItem::SavePageFormat) const

/*

 */
func (this *QWebEnginePage) Save__(filePath string) {
	var tmpArg0 = qtcore.NewQString_5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWebEngineDownloadItem::SavePageFormat=Elaborated, QWebEngineDownloadItem::SavePageFormat=Enum, , Invalid
	format := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage4saveERK7QStringN22QWebEngineDownloadItem14SavePageFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:301
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAudioMuted() const

/*

 */
func (this *QWebEnginePage) IsAudioMuted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage12isAudioMutedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:302
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAudioMuted(bool)

/*

 */
func (this *QWebEnginePage) SetAudioMuted(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage13setAudioMutedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:303
// index:0
// Public Visibility=Default Availability=Available
// [1] bool recentlyAudible() const

/*

 */
func (this *QWebEnginePage) RecentlyAudible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage15recentlyAudibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:305
// index:0
// Public Visibility=Default Availability=Available
// [-2] void printToPdf(const QString &, const QPageLayout &)

/*

 */
func (this *QWebEnginePage) PrintToPdf(filePath string, layout qtgui.QPageLayout_ITF) {
	var tmpArg0 = qtcore.NewQString_5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if layout != nil && layout.QPageLayout_PTR() != nil {
		convArg1 = layout.QPageLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage10printToPdfERK7QStringRK11QPageLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:305
// index:0
// Public Visibility=Default Availability=Available
// [-2] void printToPdf(const QString &, const QPageLayout &)

/*

 */
func (this *QWebEnginePage) PrintToPdf__(filePath string) {
	var tmpArg0 = qtcore.NewQString_5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QPageLayout &=LValueReference, QPageLayout=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage10printToPdfERK7QStringRK11QPageLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:318
// index:0
// Public Visibility=Default Availability=Available
// [8] const QWebEngineContextMenuData & contextMenuData() const

/*

 */
func (this *QWebEnginePage) ContextMenuData() *QWebEngineContextMenuData {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QWebEnginePage15contextMenuDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineContextMenuDataFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineContextMenuData)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:321
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadStarted()

/*

 */
func (this *QWebEnginePage) LoadStarted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage11loadStartedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:322
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadProgress(int)

/*

 */
func (this *QWebEnginePage) LoadProgress(progress int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage12loadProgressEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), progress)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:323
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadFinished(bool)

/*

 */
func (this *QWebEnginePage) LoadFinished(ok bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage12loadFinishedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ok)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:325
// index:0
// Public Visibility=Default Availability=Available
// [-2] void linkHovered(const QString &)

/*

 */
func (this *QWebEnginePage) LinkHovered(url string) {
	var tmpArg0 = qtcore.NewQString_5(url)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage11linkHoveredERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:326
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()

/*

 */
func (this *QWebEnginePage) SelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage16selectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:327
// index:0
// Public Visibility=Default Availability=Available
// [-2] void geometryChangeRequested(const QRect &)

/*

 */
func (this *QWebEnginePage) GeometryChangeRequested(geom qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if geom != nil && geom.QRect_PTR() != nil {
		convArg0 = geom.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage23geometryChangeRequestedERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:328
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowCloseRequested()

/*

 */
func (this *QWebEnginePage) WindowCloseRequested() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage20windowCloseRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:330
// index:0
// Public Visibility=Default Availability=Available
// [-2] void featurePermissionRequested(const QUrl &, QWebEnginePage::Feature)

/*

 */
func (this *QWebEnginePage) FeaturePermissionRequested(securityOrigin qtcore.QUrl_ITF, feature int) {
	var convArg0 unsafe.Pointer
	if securityOrigin != nil && securityOrigin.QUrl_PTR() != nil {
		convArg0 = securityOrigin.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage26featurePermissionRequestedERK4QUrlNS_7FeatureE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, feature)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:331
// index:0
// Public Visibility=Default Availability=Available
// [-2] void featurePermissionRequestCanceled(const QUrl &, QWebEnginePage::Feature)

/*

 */
func (this *QWebEnginePage) FeaturePermissionRequestCanceled(securityOrigin qtcore.QUrl_ITF, feature int) {
	var convArg0 unsafe.Pointer
	if securityOrigin != nil && securityOrigin.QUrl_PTR() != nil {
		convArg0 = securityOrigin.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage32featurePermissionRequestCanceledERK4QUrlNS_7FeatureE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, feature)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:332
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fullScreenRequested(QWebEngineFullScreenRequest)

/*

 */
func (this *QWebEnginePage) FullScreenRequested(fullScreenRequest QWebEngineFullScreenRequest_ITF /*123*/) {
	var convArg0 unsafe.Pointer
	if fullScreenRequest != nil && fullScreenRequest.QWebEngineFullScreenRequest_PTR() != nil {
		convArg0 = fullScreenRequest.QWebEngineFullScreenRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage19fullScreenRequestedE27QWebEngineFullScreenRequest", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:334
// index:0
// Public Visibility=Default Availability=Available
// [-2] void authenticationRequired(const QUrl &, QAuthenticator *)

/*

 */
func (this *QWebEnginePage) AuthenticationRequired(requestUrl qtcore.QUrl_ITF, authenticator qtnetwork.QAuthenticator_ITF /*777 QAuthenticator **/) {
	var convArg0 unsafe.Pointer
	if requestUrl != nil && requestUrl.QUrl_PTR() != nil {
		convArg0 = requestUrl.QUrl_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if authenticator != nil && authenticator.QAuthenticator_PTR() != nil {
		convArg1 = authenticator.QAuthenticator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage22authenticationRequiredERK4QUrlP14QAuthenticator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:335
// index:0
// Public Visibility=Default Availability=Available
// [-2] void proxyAuthenticationRequired(const QUrl &, QAuthenticator *, const QString &)

/*

 */
func (this *QWebEnginePage) ProxyAuthenticationRequired(requestUrl qtcore.QUrl_ITF, authenticator qtnetwork.QAuthenticator_ITF /*777 QAuthenticator **/, proxyHost string) {
	var convArg0 unsafe.Pointer
	if requestUrl != nil && requestUrl.QUrl_PTR() != nil {
		convArg0 = requestUrl.QUrl_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if authenticator != nil && authenticator.QAuthenticator_PTR() != nil {
		convArg1 = authenticator.QAuthenticator_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString_5(proxyHost)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage27proxyAuthenticationRequiredERK4QUrlP14QAuthenticatorRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:337
// index:0
// Public Visibility=Default Availability=Available
// [-2] void renderProcessTerminated(QWebEnginePage::RenderProcessTerminationStatus, int)

/*

 */
func (this *QWebEnginePage) RenderProcessTerminated(terminationStatus int, exitCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage23renderProcessTerminatedENS_30RenderProcessTerminationStatusEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), terminationStatus, exitCode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:340
// index:0
// Public Visibility=Default Availability=Available
// [-2] void titleChanged(const QString &)

/*

 */
func (this *QWebEnginePage) TitleChanged(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage12titleChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:341
// index:0
// Public Visibility=Default Availability=Available
// [-2] void urlChanged(const QUrl &)

/*

 */
func (this *QWebEnginePage) UrlChanged(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage10urlChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:342
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconUrlChanged(const QUrl &)

/*

 */
func (this *QWebEnginePage) IconUrlChanged(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage14iconUrlChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:343
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconChanged(const QIcon &)

/*

 */
func (this *QWebEnginePage) IconChanged(icon qtgui.QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if icon != nil && icon.QIcon_PTR() != nil {
		convArg0 = icon.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage11iconChangedERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:345
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollPositionChanged(const QPointF &)

/*

 */
func (this *QWebEnginePage) ScrollPositionChanged(position qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if position != nil && position.QPointF_PTR() != nil {
		convArg0 = position.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage21scrollPositionChangedERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:346
// index:0
// Public Visibility=Default Availability=Available
// [-2] void contentsSizeChanged(const QSizeF &)

/*

 */
func (this *QWebEnginePage) ContentsSizeChanged(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage19contentsSizeChangedERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:347
// index:0
// Public Visibility=Default Availability=Available
// [-2] void audioMutedChanged(bool)

/*

 */
func (this *QWebEnginePage) AudioMutedChanged(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage17audioMutedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:348
// index:0
// Public Visibility=Default Availability=Available
// [-2] void recentlyAudibleChanged(bool)

/*

 */
func (this *QWebEnginePage) RecentlyAudibleChanged(recentlyAudible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage22recentlyAudibleChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), recentlyAudible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:350
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pdfPrintingFinished(const QString &, bool)

/*

 */
func (this *QWebEnginePage) PdfPrintingFinished(filePath string, success bool) {
	var tmpArg0 = qtcore.NewQString_5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage19pdfPrintingFinishedERK7QStringb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, success)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:353
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QWebEnginePage * createWindow(QWebEnginePage::WebWindowType)

/*

 */
func (this *QWebEnginePage) CreateWindow(type_ int) *QWebEnginePage /*777 QWebEnginePage **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage12createWindowENS_13WebWindowTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEnginePageFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:354
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QStringList chooseFiles(QWebEnginePage::FileSelectionMode, const QStringList &, const QStringList &)

/*

 */
func (this *QWebEnginePage) ChooseFiles(mode int, oldFiles qtcore.QStringList_ITF, acceptedMimeTypes qtcore.QStringList_ITF) *qtcore.QStringList /*123*/ {
	var convArg1 unsafe.Pointer
	if oldFiles != nil && oldFiles.QStringList_PTR() != nil {
		convArg1 = oldFiles.QStringList_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if acceptedMimeTypes != nil && acceptedMimeTypes.QStringList_PTR() != nil {
		convArg2 = acceptedMimeTypes.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage11chooseFilesENS_17FileSelectionModeERK11QStringListS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:355
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void javaScriptAlert(const QUrl &, const QString &)

/*

 */
func (this *QWebEnginePage) JavaScriptAlert(securityOrigin qtcore.QUrl_ITF, msg string) {
	var convArg0 unsafe.Pointer
	if securityOrigin != nil && securityOrigin.QUrl_PTR() != nil {
		convArg0 = securityOrigin.QUrl_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(msg)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage15javaScriptAlertERK4QUrlRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:356
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool javaScriptConfirm(const QUrl &, const QString &)

/*

 */
func (this *QWebEnginePage) JavaScriptConfirm(securityOrigin qtcore.QUrl_ITF, msg string) bool {
	var convArg0 unsafe.Pointer
	if securityOrigin != nil && securityOrigin.QUrl_PTR() != nil {
		convArg0 = securityOrigin.QUrl_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(msg)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage17javaScriptConfirmERK4QUrlRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:357
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool javaScriptPrompt(const QUrl &, const QString &, const QString &, QString *)

/*

 */
func (this *QWebEnginePage) JavaScriptPrompt(securityOrigin qtcore.QUrl_ITF, msg string, defaultValue string, result string) bool {
	var convArg0 unsafe.Pointer
	if securityOrigin != nil && securityOrigin.QUrl_PTR() != nil {
		convArg0 = securityOrigin.QUrl_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(msg)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(defaultValue)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(result)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage16javaScriptPromptERK4QUrlRK7QStringS5_PS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:358
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void javaScriptConsoleMessage(QWebEnginePage::JavaScriptConsoleMessageLevel, const QString &, int, const QString &)

/*

 */
func (this *QWebEnginePage) JavaScriptConsoleMessage(level int, message string, lineNumber int, sourceID string) {
	var tmpArg1 = qtcore.NewQString_5(message)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(sourceID)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage24javaScriptConsoleMessageENS_29JavaScriptConsoleMessageLevelERK7QStringiS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), level, convArg1, lineNumber, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:359
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool certificateError(const QWebEngineCertificateError &)

/*

 */
func (this *QWebEnginePage) CertificateError(certificateError QWebEngineCertificateError_ITF) bool {
	var convArg0 unsafe.Pointer
	if certificateError != nil && certificateError.QWebEngineCertificateError_PTR() != nil {
		convArg0 = certificateError.QWebEngineCertificateError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage16certificateErrorERK26QWebEngineCertificateError", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:360
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool acceptNavigationRequest(const QUrl &, QWebEnginePage::NavigationType, bool)

/*

 */
func (this *QWebEnginePage) AcceptNavigationRequest(url qtcore.QUrl_ITF, type_ int, isMainFrame bool) bool {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QWebEnginePage23acceptNavigationRequestERK4QUrlNS_14NavigationTypeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, type_, isMainFrame)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QWebEnginePage__WebAction = int

//
const QWebEnginePage__NoWebAction QWebEnginePage__WebAction = -1

//
const QWebEnginePage__Back QWebEnginePage__WebAction = 0

//
const QWebEnginePage__Forward QWebEnginePage__WebAction = 1

//
const QWebEnginePage__Stop QWebEnginePage__WebAction = 2

//
const QWebEnginePage__Reload QWebEnginePage__WebAction = 3

//
const QWebEnginePage__Cut QWebEnginePage__WebAction = 4

//
const QWebEnginePage__Copy QWebEnginePage__WebAction = 5

//
const QWebEnginePage__Paste QWebEnginePage__WebAction = 6

//
const QWebEnginePage__Undo QWebEnginePage__WebAction = 7

//
const QWebEnginePage__Redo QWebEnginePage__WebAction = 8

//
const QWebEnginePage__SelectAll QWebEnginePage__WebAction = 9

//
const QWebEnginePage__ReloadAndBypassCache QWebEnginePage__WebAction = 10

//
const QWebEnginePage__PasteAndMatchStyle QWebEnginePage__WebAction = 11

//
const QWebEnginePage__OpenLinkInThisWindow QWebEnginePage__WebAction = 12

//
const QWebEnginePage__OpenLinkInNewWindow QWebEnginePage__WebAction = 13

//
const QWebEnginePage__OpenLinkInNewTab QWebEnginePage__WebAction = 14

//
const QWebEnginePage__CopyLinkToClipboard QWebEnginePage__WebAction = 15

//
const QWebEnginePage__DownloadLinkToDisk QWebEnginePage__WebAction = 16

//
const QWebEnginePage__CopyImageToClipboard QWebEnginePage__WebAction = 17

//
const QWebEnginePage__CopyImageUrlToClipboard QWebEnginePage__WebAction = 18

//
const QWebEnginePage__DownloadImageToDisk QWebEnginePage__WebAction = 19

//
const QWebEnginePage__CopyMediaUrlToClipboard QWebEnginePage__WebAction = 20

//
const QWebEnginePage__ToggleMediaControls QWebEnginePage__WebAction = 21

//
const QWebEnginePage__ToggleMediaLoop QWebEnginePage__WebAction = 22

//
const QWebEnginePage__ToggleMediaPlayPause QWebEnginePage__WebAction = 23

//
const QWebEnginePage__ToggleMediaMute QWebEnginePage__WebAction = 24

//
const QWebEnginePage__DownloadMediaToDisk QWebEnginePage__WebAction = 25

//
const QWebEnginePage__InspectElement QWebEnginePage__WebAction = 26

//
const QWebEnginePage__ExitFullScreen QWebEnginePage__WebAction = 27

//
const QWebEnginePage__RequestClose QWebEnginePage__WebAction = 28

//
const QWebEnginePage__Unselect QWebEnginePage__WebAction = 29

//
const QWebEnginePage__SavePage QWebEnginePage__WebAction = 30

//
const QWebEnginePage__OpenLinkInNewBackgroundTab QWebEnginePage__WebAction = 31

//
const QWebEnginePage__ViewSource QWebEnginePage__WebAction = 32

//
const QWebEnginePage__ToggleBold QWebEnginePage__WebAction = 33

//
const QWebEnginePage__ToggleItalic QWebEnginePage__WebAction = 34

//
const QWebEnginePage__ToggleUnderline QWebEnginePage__WebAction = 35

//
const QWebEnginePage__ToggleStrikethrough QWebEnginePage__WebAction = 36

//
const QWebEnginePage__AlignLeft QWebEnginePage__WebAction = 37

//
const QWebEnginePage__AlignCenter QWebEnginePage__WebAction = 38

//
const QWebEnginePage__AlignRight QWebEnginePage__WebAction = 39

//
const QWebEnginePage__AlignJustified QWebEnginePage__WebAction = 40

//
const QWebEnginePage__Indent QWebEnginePage__WebAction = 41

//
const QWebEnginePage__Outdent QWebEnginePage__WebAction = 42

//
const QWebEnginePage__InsertOrderedList QWebEnginePage__WebAction = 43

//
const QWebEnginePage__InsertUnorderedList QWebEnginePage__WebAction = 44

//
const QWebEnginePage__WebActionCount QWebEnginePage__WebAction = 45

/*


 */
type QWebEnginePage__FindFlag = int

//
const QWebEnginePage__FindBackward QWebEnginePage__FindFlag = 1

//
const QWebEnginePage__FindCaseSensitively QWebEnginePage__FindFlag = 2

/*


 */
type QWebEnginePage__WebWindowType = int

//
const QWebEnginePage__WebBrowserWindow QWebEnginePage__WebWindowType = 0

//
const QWebEnginePage__WebBrowserTab QWebEnginePage__WebWindowType = 1

//
const QWebEnginePage__WebDialog QWebEnginePage__WebWindowType = 2

//
const QWebEnginePage__WebBrowserBackgroundTab QWebEnginePage__WebWindowType = 3

/*


 */
type QWebEnginePage__PermissionPolicy = int

//
const QWebEnginePage__PermissionUnknown QWebEnginePage__PermissionPolicy = 0

//
const QWebEnginePage__PermissionGrantedByUser QWebEnginePage__PermissionPolicy = 1

//
const QWebEnginePage__PermissionDeniedByUser QWebEnginePage__PermissionPolicy = 2

/*


 */
type QWebEnginePage__NavigationType = int

//
const QWebEnginePage__NavigationTypeLinkClicked QWebEnginePage__NavigationType = 0

//
const QWebEnginePage__NavigationTypeTyped QWebEnginePage__NavigationType = 1

//
const QWebEnginePage__NavigationTypeFormSubmitted QWebEnginePage__NavigationType = 2

//
const QWebEnginePage__NavigationTypeBackForward QWebEnginePage__NavigationType = 3

//
const QWebEnginePage__NavigationTypeReload QWebEnginePage__NavigationType = 4

//
const QWebEnginePage__NavigationTypeOther QWebEnginePage__NavigationType = 5

/*


 */
type QWebEnginePage__Feature = int

//
const QWebEnginePage__Notifications QWebEnginePage__Feature = 0

//
const QWebEnginePage__Geolocation QWebEnginePage__Feature = 1

//
const QWebEnginePage__MediaAudioCapture QWebEnginePage__Feature = 2

//
const QWebEnginePage__MediaVideoCapture QWebEnginePage__Feature = 3

//
const QWebEnginePage__MediaAudioVideoCapture QWebEnginePage__Feature = 4

//
const QWebEnginePage__MouseLock QWebEnginePage__Feature = 5

//
const QWebEnginePage__DesktopVideoCapture QWebEnginePage__Feature = 6

//
const QWebEnginePage__DesktopAudioVideoCapture QWebEnginePage__Feature = 7

/*


 */
type QWebEnginePage__FileSelectionMode = int

//
const QWebEnginePage__FileSelectOpen QWebEnginePage__FileSelectionMode = 0

//
const QWebEnginePage__FileSelectOpenMultiple QWebEnginePage__FileSelectionMode = 1

/*


 */
type QWebEnginePage__JavaScriptConsoleMessageLevel = int

//
const QWebEnginePage__InfoMessageLevel QWebEnginePage__JavaScriptConsoleMessageLevel = 0

//
const QWebEnginePage__WarningMessageLevel QWebEnginePage__JavaScriptConsoleMessageLevel = 1

//
const QWebEnginePage__ErrorMessageLevel QWebEnginePage__JavaScriptConsoleMessageLevel = 2

/*


 */
type QWebEnginePage__RenderProcessTerminationStatus = int

//
const QWebEnginePage__NormalTerminationStatus QWebEnginePage__RenderProcessTerminationStatus = 0

//
const QWebEnginePage__AbnormalTerminationStatus QWebEnginePage__RenderProcessTerminationStatus = 1

//
const QWebEnginePage__CrashedTerminationStatus QWebEnginePage__RenderProcessTerminationStatus = 2

//
const QWebEnginePage__KilledTerminationStatus QWebEnginePage__RenderProcessTerminationStatus = 3

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
