package qtwidgets

// /usr/include/qt/QtWidgets/qwizard.h
// #include <qwizard.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
func (this *QWizard) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QWizard) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(QPaintEvent *)
func (this *QWizard) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void done(int)
func (this *QWizard) InheritDone(f func(result int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "done", f)
}

// void initializePage(int)
func (this *QWizard) InheritInitializePage(f func(id int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initializePage", f)
}

// void cleanupPage(int)
func (this *QWizard) InheritCleanupPage(f func(id int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "cleanupPage", f)
}

/*

 */
type QWizard struct {
	*QDialog
}
type QWizard_ITF interface {
	QDialog_ITF
	QWizard_PTR() *QWizard
}

func (ptr *QWizard) QWizard_PTR() *QWizard { return ptr }

func (this *QWizard) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QWizard) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQWizardFromPointer(cthis unsafe.Pointer) *QWizard {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QWizard{bcthis0}
}
func (*QWizard) NewFromPointer(cthis unsafe.Pointer) *QWizard {
	return NewQWizardFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qwizard.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWizard) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWizard(QWidget *, Qt::WindowFlags)

/*
Constructs a wizard with the given parent and window flags.

See also parent() and windowFlags().
*/
func (*QWizard) NewForInherit(parent QWidget_ITF /*777 QWidget **/, flags int) *QWizard {
	return NewQWizard(parent, flags)
}
func NewQWizard(parent QWidget_ITF /*777 QWidget **/, flags int) *QWizard {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizardC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWizardFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWizard")
	return gothis
}

// /usr/include/qt/QtWidgets/qwizard.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWizard(QWidget *, Qt::WindowFlags)

/*
Constructs a wizard with the given parent and window flags.

See also parent() and windowFlags().
*/
func (*QWizard) NewForInherit__() *QWizard {
	return NewQWizard__()
}
func NewQWizard__() *QWizard {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizardC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWizardFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWizard")
	return gothis
}

// /usr/include/qt/QtWidgets/qwizard.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWizard(QWidget *, Qt::WindowFlags)

/*
Constructs a wizard with the given parent and window flags.

See also parent() and windowFlags().
*/
func (*QWizard) NewForInherit__1(parent QWidget_ITF /*777 QWidget **/) *QWizard {
	return NewQWizard__1(parent)
}
func NewQWizard__1(parent QWidget_ITF /*777 QWidget **/) *QWizard {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizardC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWizardFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWizard")
	return gothis
}

// /usr/include/qt/QtWidgets/qwizard.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWizard()

/*

 */
func DeleteQWizard(this *QWizard) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizardD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qwizard.h:126
// index:0
// Public Visibility=Default Availability=Available
// [4] int addPage(QWizardPage *)

/*
Adds the given page to the wizard, and returns the page's ID.

The ID is guaranteed to be larger than any other ID in the QWizard so far.

See also setPage(), page(), and pageAdded().
*/
func (this *QWizard) AddPage(page QWizardPage_ITF /*777 QWizardPage **/) int {
	var convArg0 unsafe.Pointer
	if page != nil && page.QWizardPage_PTR() != nil {
		convArg0 = page.QWizardPage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard7addPageEP11QWizardPage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwizard.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPage(int, QWizardPage *)

/*
Adds the given page to the wizard with the given id.

Note: Adding a page may influence the value of the startId property in case it was not set explicitly.

See also addPage(), page(), and pageAdded().
*/
func (this *QWizard) SetPage(id int, page QWizardPage_ITF /*777 QWizardPage **/) {
	var convArg1 unsafe.Pointer
	if page != nil && page.QWizardPage_PTR() != nil {
		convArg1 = page.QWizardPage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard7setPageEiP11QWizardPage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removePage(int)

/*
Removes the page with the given id. cleanupPage() will be called if necessary.

Note: Removing a page may influence the value of the startId property.

This function was introduced in  Qt 4.5.

See also addPage(), setPage(), pageRemoved(), and startId().
*/
func (this *QWizard) RemovePage(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard10removePageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QWizardPage * page(int) const

/*
Returns the page with the given id, or 0 if there is no such page.

See also addPage() and setPage().
*/
func (this *QWizard) Page(id int) *QWizardPage /*777 QWizardPage **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard4pageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWizardPageFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasVisitedPage(int) const

/*
Returns true if the page history contains page id; otherwise, returns false.

Pressing Back marks the current page as "unvisited" again.

See also visitedPages().
*/
func (this *QWizard) HasVisitedPage(id int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard14hasVisitedPageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartId(int)

/*

 */
func (this *QWizard) SetStartId(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard10setStartIdEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int startId() const

/*

 */
func (this *QWizard) StartId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard7startIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwizard.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QWizardPage * currentPage() const

/*
Returns a pointer to the current page, or 0 if there is no current page (e.g., before the wizard is shown).

This is equivalent to calling page(currentId()).

See also page(), currentId(), and restart().
*/
func (this *QWizard) CurrentPage() *QWizardPage /*777 QWizardPage **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard11currentPageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWizardPageFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentId() const

/*

 */
func (this *QWizard) CurrentId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard9currentIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwizard.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool validateCurrentPage()

/*
This virtual function is called by QWizard when the user clicks Next or Finish to perform some last-minute validation. If it returns true, the next page is shown (or the wizard finishes); otherwise, the current page stays up.

The default implementation calls QWizardPage::validatePage() on the currentPage().

When possible, it is usually better style to disable the Next or Finish button (by specifying mandatory fields or by reimplementing QWizardPage::isComplete()) than to reimplement validateCurrentPage().

See also QWizardPage::validatePage() and currentPage().
*/
func (this *QWizard) ValidateCurrentPage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard19validateCurrentPageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int nextId() const

/*
This virtual function is called by QWizard to find out which page to show when the user clicks the Next button.

The return value is the ID of the next page, or -1 if no page follows.

The default implementation calls QWizardPage::nextId() on the currentPage().

By reimplementing this function, you can specify a dynamic page order.

See also QWizardPage::nextId() and currentPage().
*/
func (this *QWizard) NextId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard6nextIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwizard.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setField(const QString &, const QVariant &)

/*
Sets the value of the field called name to value.

This function can be used to set fields on any page of the wizard.

See also QWizardPage::registerField(), QWizardPage::setField(), and field().
*/
func (this *QWizard) SetField(name string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard8setFieldERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:142
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant field(const QString &) const

/*
Returns the value of the field called name.

This function can be used to access fields on any page of the wizard.

See also QWizardPage::registerField(), QWizardPage::field(), and setField().
*/
func (this *QWizard) Field(name string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard5fieldERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWizardStyle(QWizard::WizardStyle)

/*

 */
func (this *QWizard) SetWizardStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard14setWizardStyleENS_11WizardStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:145
// index:0
// Public Visibility=Default Availability=Available
// [4] QWizard::WizardStyle wizardStyle() const

/*

 */
func (this *QWizard) WizardStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard11wizardStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QWizard::WizardOption, bool)

/*
Sets the given option to be enabled if on is true; otherwise, clears the given option.

See also options, testOption(), and setWizardStyle().
*/
func (this *QWizard) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard9setOptionENS_12WizardOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QWizard::WizardOption, bool)

/*
Sets the given option to be enabled if on is true; otherwise, clears the given option.

See also options, testOption(), and setWizardStyle().
*/
func (this *QWizard) SetOption__(option int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard9setOptionENS_12WizardOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(QWizard::WizardOption) const

/*
Returns true if the given option is enabled; otherwise, returns false.

See also options, setOption(), and setWizardStyle().
*/
func (this *QWizard) TestOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard10testOptionENS_12WizardOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QWizard::WizardOptions)

/*

 */
func (this *QWizard) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard10setOptionsE6QFlagsINS_12WizardOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:150
// index:0
// Public Visibility=Default Availability=Available
// [4] QWizard::WizardOptions options() const

/*

 */
func (this *QWizard) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtonText(QWizard::WizardButton, const QString &)

/*
Sets the text on button which to be text.

By default, the text on buttons depends on the wizardStyle. For example, on macOS, the Next button is called Continue.

To add extra buttons to the wizard (e.g., a Print button), one way is to call setButtonText() with CustomButton1, CustomButton2, or CustomButton3 to set their text, and make the buttons visible using the HaveCustomButton1, HaveCustomButton2, and/or HaveCustomButton3 options.

Button texts may also be set on a per-page basis using QWizardPage::setButtonText().

See also buttonText(), setButton(), button(), setButtonLayout(), setOptions(), and QWizardPage::setButtonText().
*/
func (this *QWizard) SetButtonText(which int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard13setButtonTextENS_12WizardButtonERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] QString buttonText(QWizard::WizardButton) const

/*
Returns the text on button which.

If a text has ben set using setButtonText(), this text is returned.

By default, the text on buttons depends on the wizardStyle. For example, on macOS, the Next button is called Continue.

See also button(), setButton(), setButtonText(), QWizardPage::buttonText(), and QWizardPage::setButtonText().
*/
func (this *QWizard) ButtonText(which int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard10buttonTextENS_12WizardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwizard.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButton(QWizard::WizardButton, QAbstractButton *)

/*
Sets the button corresponding to role which to button.

To add extra buttons to the wizard (e.g., a Print button), one way is to call setButton() with CustomButton1 to CustomButton3, and make the buttons visible using the HaveCustomButton1 to HaveCustomButton3 options.

See also button(), setButtonText(), setButtonLayout(), and options.
*/
func (this *QWizard) SetButton(which int, button QAbstractButton_ITF /*777 QAbstractButton **/) {
	var convArg1 unsafe.Pointer
	if button != nil && button.QAbstractButton_PTR() != nil {
		convArg1 = button.QAbstractButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard9setButtonENS_12WizardButtonEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractButton * button(QWizard::WizardButton) const

/*
Returns the button corresponding to role which.

See also setButton() and setButtonText().
*/
func (this *QWizard) Button(which int) *QAbstractButton /*777 QAbstractButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard6buttonENS_12WizardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitleFormat(Qt::TextFormat)

/*

 */
func (this *QWizard) SetTitleFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard14setTitleFormatEN2Qt10TextFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:159
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextFormat titleFormat() const

/*

 */
func (this *QWizard) TitleFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard11titleFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSubTitleFormat(Qt::TextFormat)

/*

 */
func (this *QWizard) SetSubTitleFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard17setSubTitleFormatEN2Qt10TextFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:161
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextFormat subTitleFormat() const

/*

 */
func (this *QWizard) SubTitleFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard14subTitleFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(QWizard::WizardPixmap, const QPixmap &)

/*
Sets the pixmap for role which to pixmap.

The pixmaps are used by QWizard when displaying a page. Which pixmaps are actually used depend on the wizard style.

Pixmaps can also be set for a specific page using QWizardPage::setPixmap().

See also pixmap(), QWizardPage::setPixmap(), and Elements of a Wizard Page.
*/
func (this *QWizard) SetPixmap(which int, pixmap qtgui.QPixmap_ITF) {
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard9setPixmapENS_12WizardPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:163
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(QWizard::WizardPixmap) const

/*
Returns the pixmap set for role which.

By default, the only pixmap that is set is the BackgroundPixmap on macOS.

See also setPixmap(), QWizardPage::pixmap(), and Elements of a Wizard Page.
*/
func (this *QWizard) Pixmap(which int) *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard6pixmapENS_12WizardPixmapE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSideWidget(QWidget *)

/*
Sets the given widget to be shown on the left side of the wizard. For styles which use the WatermarkPixmap (ClassicStyle and ModernStyle) the side widget is displayed on top of the watermark, for other styles or when the watermark is not provided the side widget is displayed on the left side of the wizard.

Passing 0 shows no side widget.

When the widget is not 0 the wizard reparents it.

Any previous side widget is hidden.

You may call setSideWidget() with the same widget at different times.

All widgets set here will be deleted by the wizard when it is destroyed unless you separately reparent the widget after setting some other side widget (or 0).

By default, no side widget is present.

This function was introduced in  Qt 4.7.

See also sideWidget().
*/
func (this *QWizard) SetSideWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard13setSideWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * sideWidget() const

/*
Returns the widget on the left side of the wizard or 0.

By default, no side widget is present.

This function was introduced in  Qt 4.7.

See also setSideWidget().
*/
func (this *QWizard) SideWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard10sideWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultProperty(const char *, const char *, const char *)

/*
Sets the default property for className to be property, and the associated change signal to be changedSignal.

The default property is used when an instance of className (or of one of its subclasses) is passed to QWizardPage::registerField() and no property is specified.

QWizard knows the most common Qt widgets. For these (or their subclasses), you don't need to specify a property or a changedSignal. The table below lists these widgets:


 WidgetPropertyChange Notification Signal
QAbstractButtonbool checkedtoggled()
QAbstractSliderint valuevalueChanged()
QComboBoxint currentIndexcurrentIndexChanged()
QDateTimeEditQDateTime dateTimedateTimeChanged()
QLineEditQString texttextChanged()
QListWidgetint currentRowcurrentRowChanged()
QSpinBoxint valuevalueChanged()


See also QWizardPage::registerField().
*/
func (this *QWizard) SetDefaultProperty(className string, property string, changedSignal string) {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(property)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(changedSignal)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard18setDefaultPropertyEPKcS1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:171
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
Reimplemented from QWidget::setVisible().
*/
func (this *QWizard) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:172
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QWizard) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentIdChanged(int)

/*
This signal is emitted when the current page changes, with the new current id.

Note: Notifier signal for property currentId.

See also currentId() and currentPage().
*/
func (this *QWizard) CurrentIdChanged(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard16currentIdChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void helpRequested()

/*
This signal is emitted when the user clicks the Help button.

By default, no Help button is shown. Call setOption(HaveHelpButton, true) to have one.

Example:


  LicenseWizard::LicenseWizard(QWidget *parent)
      : QWizard(parent)
  {
      ...
      setOption(HaveHelpButton, true);
      connect(this, &QWizard::helpRequested, this, &LicenseWizard::showHelp);
      ...
  }

  void LicenseWizard::showHelp()
  {
      static QString lastHelpMessage;

      QString message;

      switch (currentId()) {
      case Page_Intro:
          message = tr("The decision you make here will affect which page you "
                       "get to see next.");
          break;
      ...
      default:
          message = tr("This help is likely not to be of any help.");
      }

      QMessageBox::information(this, tr("License Wizard Help"), message);

  }



See also customButtonClicked().
*/
func (this *QWizard) HelpRequested() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard13helpRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void customButtonClicked(int)

/*
This signal is emitted when the user clicks a custom button. which can be CustomButton1, CustomButton2, or CustomButton3.

By default, no custom button is shown. Call setOption() with HaveCustomButton1, HaveCustomButton2, or HaveCustomButton3 to have one, and use setButtonText() or setButton() to configure it.

See also helpRequested().
*/
func (this *QWizard) CustomButtonClicked(which int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard19customButtonClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pageAdded(int)

/*
This signal is emitted whenever a page is added to the wizard. The page's id is passed as parameter.

This function was introduced in  Qt 4.7.

See also addPage(), setPage(), and startId().
*/
func (this *QWizard) PageAdded(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard9pageAddedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pageRemoved(int)

/*
This signal is emitted whenever a page is removed from the wizard. The page's id is passed as parameter.

This function was introduced in  Qt 4.7.

See also removePage() and startId().
*/
func (this *QWizard) PageRemoved(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard11pageRemovedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void back()

/*
Goes back to the previous page.

This is equivalent to pressing the Back button.

See also next(), accept(), reject(), and restart().
*/
func (this *QWizard) Back() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void next()

/*
Advances to the next page.

This is equivalent to pressing the Next or Commit button.

See also nextId(), back(), accept(), reject(), and restart().
*/
func (this *QWizard) Next() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard4nextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void restart()

/*
Restarts the wizard at the start page. This function is called automatically when the wizard is shown.

See also startId().
*/
func (this *QWizard) Restart() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard7restartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:187
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QWizard) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:188
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QWizard) ResizeEvent(event qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QResizeEvent_PTR() != nil {
		convArg0 = event.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:189
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QWizard) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:193
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void done(int)

/*
Reimplemented from QDialog::done().
*/
func (this *QWizard) Done(result int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), result)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:194
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initializePage(int)

/*
This virtual function is called by QWizard to prepare page id just before it is shown either as a result of QWizard::restart() being called, or as a result of the user clicking Next. (However, if the QWizard::IndependentPages option is set, this function is only called the first time the page is shown.)

By reimplementing this function, you can ensure that the page's fields are properly initialized based on fields from previous pages.

The default implementation calls QWizardPage::initializePage() on page(id).

See also QWizardPage::initializePage() and cleanupPage().
*/
func (this *QWizard) InitializePage(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard14initializePageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:195
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void cleanupPage(int)

/*
This virtual function is called by QWizard to clean up page id just before the user leaves it by clicking Back (unless the QWizard::IndependentPages option is set).

The default implementation calls QWizardPage::cleanupPage() on page(id).

See also QWizardPage::cleanupPage() and initializePage().
*/
func (this *QWizard) CleanupPage(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard11cleanupPageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

/*
This enum specifies the buttons in a wizard.



The following value is only useful when calling setButtonLayout():



See also setButton(), setButtonText(), setButtonLayout(), and customButtonClicked().

*/
type QWizard__WizardButton = int

// The Back button (Go Back on macOS)
const QWizard__BackButton QWizard__WizardButton = 0

// The Next button (Continue on macOS)
const QWizard__NextButton QWizard__WizardButton = 1

// The Commit button
const QWizard__CommitButton QWizard__WizardButton = 2

// The Finish button (Done on macOS)
const QWizard__FinishButton QWizard__WizardButton = 3

// The Cancel button (see also NoCancelButton)
const QWizard__CancelButton QWizard__WizardButton = 4

// The Help button (see also HaveHelpButton)
const QWizard__HelpButton QWizard__WizardButton = 5

//
const QWizard__CustomButton1 QWizard__WizardButton = 6

//
const QWizard__CustomButton2 QWizard__WizardButton = 7

//
const QWizard__CustomButton3 QWizard__WizardButton = 8

// A horizontal stretch in the button layout
const QWizard__Stretch QWizard__WizardButton = 9

//
const QWizard__NoButton QWizard__WizardButton = -1

//
const QWizard__NStandardButtons QWizard__WizardButton = 6

//
const QWizard__NButtons QWizard__WizardButton = 9

func (this *QWizard) WizardButtonItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWizard_WizardButtonItemName(val int) string {
	var nilthis *QWizard
	return nilthis.WizardButtonItemName(val)
}

/*
This enum specifies the pixmaps that can be associated with a page.



See also setPixmap(), QWizardPage::setPixmap(), and Elements of a Wizard Page.

*/
type QWizard__WizardPixmap = int

// The tall pixmap on the left side of a ClassicStyle or ModernStyle page
const QWizard__WatermarkPixmap QWizard__WizardPixmap = 0

// The small pixmap on the right side of a ClassicStyle or ModernStyle page header
const QWizard__LogoPixmap QWizard__WizardPixmap = 1

// The pixmap that occupies the background of a ModernStyle page header
const QWizard__BannerPixmap QWizard__WizardPixmap = 2

// The pixmap that occupies the background of a MacStyle wizard
const QWizard__BackgroundPixmap QWizard__WizardPixmap = 3

//
const QWizard__NPixmaps QWizard__WizardPixmap = 4

func (this *QWizard) WizardPixmapItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWizard_WizardPixmapItemName(val int) string {
	var nilthis *QWizard
	return nilthis.WizardPixmapItemName(val)
}

/*
This enum specifies the different looks supported by QWizard.



See also setWizardStyle(), WizardOption, and Wizard Look and Feel.

*/
type QWizard__WizardStyle = int

// Classic Windows look
const QWizard__ClassicStyle QWizard__WizardStyle = 0

// Modern Windows look
const QWizard__ModernStyle QWizard__WizardStyle = 1

// macOS look
const QWizard__MacStyle QWizard__WizardStyle = 2

// Windows Aero look
const QWizard__AeroStyle QWizard__WizardStyle = 3

//
const QWizard__NStyles QWizard__WizardStyle = 4

func (this *QWizard) WizardStyleItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWizard_WizardStyleItemName(val int) string {
	var nilthis *QWizard
	return nilthis.WizardStyleItemName(val)
}

/*


 */
type QWizard__WizardOption = int

//
const QWizard__IndependentPages QWizard__WizardOption = 1

//
const QWizard__IgnoreSubTitles QWizard__WizardOption = 2

//
const QWizard__ExtendedWatermarkPixmap QWizard__WizardOption = 4

//
const QWizard__NoDefaultButton QWizard__WizardOption = 8

//
const QWizard__NoBackButtonOnStartPage QWizard__WizardOption = 16

//
const QWizard__NoBackButtonOnLastPage QWizard__WizardOption = 32

//
const QWizard__DisabledBackButtonOnLastPage QWizard__WizardOption = 64

//
const QWizard__HaveNextButtonOnLastPage QWizard__WizardOption = 128

//
const QWizard__HaveFinishButtonOnEarlyPages QWizard__WizardOption = 256

//
const QWizard__NoCancelButton QWizard__WizardOption = 512

//
const QWizard__CancelButtonOnLeft QWizard__WizardOption = 1024

//
const QWizard__HaveHelpButton QWizard__WizardOption = 2048

//
const QWizard__HelpButtonOnRight QWizard__WizardOption = 4096

//
const QWizard__HaveCustomButton1 QWizard__WizardOption = 8192

//
const QWizard__HaveCustomButton2 QWizard__WizardOption = 16384

//
const QWizard__HaveCustomButton3 QWizard__WizardOption = 32768

//
const QWizard__NoCancelButtonOnLastPage QWizard__WizardOption = 65536

func (this *QWizard) WizardOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWizard_WizardOptionItemName(val int) string {
	var nilthis *QWizard
	return nilthis.WizardOptionItemName(val)
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
