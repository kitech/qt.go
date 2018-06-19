package qtwidgets

// /usr/include/qt/QtWidgets/qtextbrowser.h
// #include <qtextbrowser.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 77
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

// bool event(QEvent *)
func (this *QTextBrowser) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QTextBrowser) InheritKeyPressEvent(f func(ev *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QTextBrowser) InheritMouseMoveEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QTextBrowser) InheritMousePressEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QTextBrowser) InheritMouseReleaseEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QTextBrowser) InheritFocusOutEvent(f func(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// bool focusNextPrevChild(bool)
func (this *QTextBrowser) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// void paintEvent(QPaintEvent *)
func (this *QTextBrowser) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

/*

 */
type QTextBrowser struct {
	*QTextEdit
}
type QTextBrowser_ITF interface {
	QTextEdit_ITF
	QTextBrowser_PTR() *QTextBrowser
}

func (ptr *QTextBrowser) QTextBrowser_PTR() *QTextBrowser { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTextBrowser) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextBrowser(QWidget *)

/*
Constructs an empty QTextBrowser with parent parent.
*/
func NewQTextBrowser(parent QWidget_ITF /*777 QWidget **/) *QTextBrowser {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowserC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBrowserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextBrowser")
	return gothis
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextBrowser(QWidget *)

/*
Constructs an empty QTextBrowser with parent parent.
*/
func NewQTextBrowser__() *QTextBrowser {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowserC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBrowserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextBrowser")
	return gothis
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextBrowser()

/*

 */
func DeleteQTextBrowser(this *QTextBrowser) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowserD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl source() const

/*

 */
func (this *QTextBrowser) Source() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList searchPaths() const

/*

 */
func (this *QTextBrowser) SearchPaths() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser11searchPathsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSearchPaths(const QStringList &)

/*

 */
func (this *QTextBrowser) SetSearchPaths(paths qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if paths != nil && paths.QStringList_PTR() != nil {
		convArg0 = paths.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser14setSearchPathsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant loadResource(int, const QUrl &)

/*
Reimplemented from QTextEdit::loadResource().

This function is called when the document is loaded and for each image in the document. The type indicates the type of resource to be loaded. An invalid QVariant is returned if the resource cannot be loaded.

The default implementation ignores type and tries to locate the resources by interpreting name as a file name. If it is not an absolute path it tries to find the file in the paths of the searchPaths property and in the same directory as the current source. On success, the result is a QVariant that stores a QByteArray with the contents of the file.

If you reimplement this function, you can return other QVariant types. The table below shows which variant types are supported depending on the resource type:


 ResourceTypeQVariant::Type
QTextDocument::HtmlResourceQString or QByteArray
QTextDocument::ImageResourceQImage, QPixmap or QByteArray
QTextDocument::StyleSheetResourceQString or QByteArray
*/
func (this *QTextBrowser) LoadResource(type_ int, name qtcore.QUrl_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg1 = name.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser12loadResourceEiRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBackwardAvailable() const

/*
Returns true if the text browser can go backward in the document history using backward().

This function was introduced in  Qt 4.2.

See also backwardAvailable() and backward().
*/
func (this *QTextBrowser) IsBackwardAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser19isBackwardAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:77
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isForwardAvailable() const

/*
Returns true if the text browser can go forward in the document history using forward().

This function was introduced in  Qt 4.2.

See also forwardAvailable() and forward().
*/
func (this *QTextBrowser) IsForwardAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser18isForwardAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearHistory()

/*
Clears the history of visited documents and disables the forward and backward navigation.

This function was introduced in  Qt 4.2.

See also backward() and forward().
*/
func (this *QTextBrowser) ClearHistory() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser12clearHistoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString historyTitle(int) const

/*
Returns the documentTitle() of the HistoryItem.


 InputReturn
i < 0backward() history
i == 0current, see QTextBrowser::source()
i > 0forward() history



  backaction.setToolTip(browser.historyTitle(-1));
  forwardaction.setToolTip(browser.historyTitle(+1));



This function was introduced in  Qt 4.4.
*/
func (this *QTextBrowser) HistoryTitle(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser12historyTitleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl historyUrl(int) const

/*
Returns the url of the HistoryItem.


 InputReturn
i < 0backward() history
i == 0current, see QTextBrowser::source()
i > 0forward() history


This function was introduced in  Qt 4.4.
*/
func (this *QTextBrowser) HistoryUrl(arg0 int) *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser10historyUrlEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int backwardHistoryCount() const

/*
Returns the number of locations backward in the history.

This function was introduced in  Qt 4.4.
*/
func (this *QTextBrowser) BackwardHistoryCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser20backwardHistoryCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int forwardHistoryCount() const

/*
Returns the number of locations forward in the history.

This function was introduced in  Qt 4.4.
*/
func (this *QTextBrowser) ForwardHistoryCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser19forwardHistoryCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool openExternalLinks() const

/*

 */
func (this *QTextBrowser) OpenExternalLinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser17openExternalLinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpenExternalLinks(bool)

/*

 */
func (this *QTextBrowser) SetOpenExternalLinks(open bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser20setOpenExternalLinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), open)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool openLinks() const

/*

 */
func (this *QTextBrowser) OpenLinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser9openLinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpenLinks(bool)

/*

 */
func (this *QTextBrowser) SetOpenLinks(open bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser12setOpenLinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), open)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSource(const QUrl &)

/*

 */
func (this *QTextBrowser) SetSource(name qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg0 = name.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser9setSourceERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void backward()

/*
Changes the document displayed to the previous document in the list of documents built by navigating links. Does nothing if there is no previous document.

See also forward() and backwardAvailable().
*/
func (this *QTextBrowser) Backward() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser8backwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void forward()

/*
Changes the document displayed to the next document in the list of documents built by navigating links. Does nothing if there is no next document.

See also backward() and forwardAvailable().
*/
func (this *QTextBrowser) Forward() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser7forwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void home()

/*
Changes the document displayed to be the first document from the history.
*/
func (this *QTextBrowser) Home() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser4homeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reload()

/*
Reloads the current set source.
*/
func (this *QTextBrowser) Reload() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser6reloadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void backwardAvailable(bool)

/*
This signal is emitted when the availability of backward() changes. available is false when the user is at home(); otherwise it is true.
*/
func (this *QTextBrowser) BackwardAvailable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser17backwardAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void forwardAvailable(bool)

/*
This signal is emitted when the availability of forward() changes. available is true after the user navigates backward() and false when the user navigates or goes forward().
*/
func (this *QTextBrowser) ForwardAvailable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser16forwardAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void historyChanged()

/*
This signal is emitted when the history changes.

This function was introduced in  Qt 4.4.

See also historyTitle() and historyUrl().
*/
func (this *QTextBrowser) HistoryChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser14historyChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sourceChanged(const QUrl &)

/*
This signal is emitted when the source has changed, src being the new source.

Source changes happen both programmatically when calling setSource(), forward(), backword() or home() or when the user clicks on links or presses the equivalent key sequences.
*/
func (this *QTextBrowser) SourceChanged(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser13sourceChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void highlighted(const QUrl &)

/*
This signal is emitted when the user has selected but not activated an anchor in the document. The URL referred to by the anchor is passed in link.

Note: Signal highlighted is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(textBrowser, QOverload<const QUrl &>::of(&QTextBrowser::highlighted),
      [=](const QUrl &link){ \/* ... *\/ });
*/
func (this *QTextBrowser) Highlighted(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser11highlightedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void highlighted(const QString &)

/*
This signal is emitted when the user has selected but not activated an anchor in the document. The URL referred to by the anchor is passed in link.

Note: Signal highlighted is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(textBrowser, QOverload<const QUrl &>::of(&QTextBrowser::highlighted),
      [=](const QUrl &link){ \/* ... *\/ });
*/
func (this *QTextBrowser) Highlighted_1(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser11highlightedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void anchorClicked(const QUrl &)

/*
This signal is emitted when the user clicks an anchor. The URL referred to by the anchor is passed in link.

Note that the browser will automatically handle navigation to the location specified by link unless the openLinks property is set to false or you call setSource() in a slot connected. This mechanism is used to override the default navigation features of the browser.
*/
func (this *QTextBrowser) AnchorClicked(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser13anchorClickedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:107
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QTextBrowser) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:108
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().

The event ev is used to provide the following keyboard shortcuts:


 KeypressAction
Alt+Left Arrowbackward()
Alt+Right Arrowforward()
Alt+Up Arrowhome()
*/
func (this *QTextBrowser) KeyPressEvent(ev qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QKeyEvent_PTR() != nil {
		convArg0 = ev.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseMoveEvent().
*/
func (this *QTextBrowser) MouseMoveEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:110
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mousePressEvent().
*/
func (this *QTextBrowser) MousePressEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWidget::mouseReleaseEvent().
*/
func (this *QTextBrowser) MouseReleaseEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QWidget::focusOutEvent().
*/
func (this *QTextBrowser) FocusOutEvent(ev qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QFocusEvent_PTR() != nil {
		convArg0 = ev.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Reimplemented from QWidget::focusNextPrevChild().
*/
func (this *QTextBrowser) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QTextBrowser) PaintEvent(e qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QPaintEvent_PTR() != nil {
		convArg0 = e.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
