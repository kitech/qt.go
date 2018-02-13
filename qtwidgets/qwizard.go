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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QWizard) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QWizard) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(class QPaintEvent *)
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
// [8] const QMetaObject * metaObject()
func (this *QWizard) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWizard(QWidget *, Qt::WindowFlags)
func NewQWizard(parent *QWidget /*777 QWidget **/, flags int) *QWizard {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizardC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWizardFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qwizard.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWizard()
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
func (this *QWizard) AddPage(page *QWizardPage /*777 QWizardPage **/) int {
	var convArg0 = page.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard7addPageEP11QWizardPage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwizard.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPage(int, QWizardPage *)
func (this *QWizard) SetPage(id int, page *QWizardPage /*777 QWizardPage **/) {
	var convArg1 = page.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard7setPageEiP11QWizardPage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removePage(int)
func (this *QWizard) RemovePage(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard10removePageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QWizardPage * page(int)
func (this *QWizard) Page(id int) *QWizardPage /*777 QWizardPage **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard4pageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWizardPageFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasVisitedPage(int)
func (this *QWizard) HasVisitedPage(id int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard14hasVisitedPageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartId(int)
func (this *QWizard) SetStartId(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard10setStartIdEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int startId()
func (this *QWizard) StartId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard7startIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwizard.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QWizardPage * currentPage()
func (this *QWizard) CurrentPage() *QWizardPage /*777 QWizardPage **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard11currentPageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWizardPageFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentId()
func (this *QWizard) CurrentId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard9currentIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwizard.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool validateCurrentPage()
func (this *QWizard) ValidateCurrentPage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard19validateCurrentPageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int nextId()
func (this *QWizard) NextId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard6nextIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwizard.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setField(const QString &, const QVariant &)
func (this *QWizard) SetField(name string, value *qtcore.QVariant) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard8setFieldERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:142
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant field(const QString &)
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
// [-2] void setWizardStyle(enum QWizard::WizardStyle)
func (this *QWizard) SetWizardStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard14setWizardStyleENS_11WizardStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:145
// index:0
// Public Visibility=Default Availability=Available
// [4] QWizard::WizardStyle wizardStyle()
func (this *QWizard) WizardStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard11wizardStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(enum QWizard::WizardOption, _Bool)
func (this *QWizard) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard9setOptionENS_12WizardOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(enum QWizard::WizardOption)
func (this *QWizard) TestOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard10testOptionENS_12WizardOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QWizard::WizardOptions)
func (this *QWizard) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard10setOptionsE6QFlagsINS_12WizardOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:150
// index:0
// Public Visibility=Default Availability=Available
// [4] QWizard::WizardOptions options()
func (this *QWizard) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtonText(enum QWizard::WizardButton, const QString &)
func (this *QWizard) SetButtonText(which int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard13setButtonTextENS_12WizardButtonERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] QString buttonText(enum QWizard::WizardButton)
func (this *QWizard) ButtonText(which int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard10buttonTextENS_12WizardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwizard.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButton(enum QWizard::WizardButton, QAbstractButton *)
func (this *QWizard) SetButton(which int, button *QAbstractButton /*777 QAbstractButton **/) {
	var convArg1 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard9setButtonENS_12WizardButtonEP15QAbstractButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractButton * button(enum QWizard::WizardButton)
func (this *QWizard) Button(which int) *QAbstractButton /*777 QAbstractButton **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard6buttonENS_12WizardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitleFormat(Qt::TextFormat)
func (this *QWizard) SetTitleFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard14setTitleFormatEN2Qt10TextFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:159
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextFormat titleFormat()
func (this *QWizard) TitleFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard11titleFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSubTitleFormat(Qt::TextFormat)
func (this *QWizard) SetSubTitleFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard17setSubTitleFormatEN2Qt10TextFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), format)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:161
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextFormat subTitleFormat()
func (this *QWizard) SubTitleFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard14subTitleFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(enum QWizard::WizardPixmap, const QPixmap &)
func (this *QWizard) SetPixmap(which int, pixmap *qtgui.QPixmap) {
	var convArg1 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard9setPixmapENS_12WizardPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:163
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(enum QWizard::WizardPixmap)
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
func (this *QWizard) SetSideWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard13setSideWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * sideWidget()
func (this *QWizard) SideWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWizard10sideWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultProperty(const char *, const char *, const char *)
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
// [-2] void setVisible(_Bool)
func (this *QWizard) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:172
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
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
func (this *QWizard) CurrentIdChanged(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard16currentIdChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void helpRequested()
func (this *QWizard) HelpRequested() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard13helpRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:177
// index:0
// Public Visibility=Default Availability=Available
// [-2] void customButtonClicked(int)
func (this *QWizard) CustomButtonClicked(which int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard19customButtonClickedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pageAdded(int)
func (this *QWizard) PageAdded(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard9pageAddedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pageRemoved(int)
func (this *QWizard) PageRemoved(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard11pageRemovedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void back()
func (this *QWizard) Back() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void next()
func (this *QWizard) Next() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard4nextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void restart()
func (this *QWizard) Restart() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard7restartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:187
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QWizard) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:188
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QWizard) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:189
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QWizard) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:193
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void done(int)
func (this *QWizard) Done(result int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), result)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:194
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initializePage(int)
func (this *QWizard) InitializePage(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard14initializePageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:195
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void cleanupPage(int)
func (this *QWizard) CleanupPage(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWizard11cleanupPageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

type QWizard__WizardButton = int

const QWizard__BackButton QWizard__WizardButton = 0
const QWizard__NextButton QWizard__WizardButton = 1
const QWizard__CommitButton QWizard__WizardButton = 2
const QWizard__FinishButton QWizard__WizardButton = 3
const QWizard__CancelButton QWizard__WizardButton = 4
const QWizard__HelpButton QWizard__WizardButton = 5
const QWizard__CustomButton1 QWizard__WizardButton = 6
const QWizard__CustomButton2 QWizard__WizardButton = 7
const QWizard__CustomButton3 QWizard__WizardButton = 8
const QWizard__Stretch QWizard__WizardButton = 9
const QWizard__NoButton QWizard__WizardButton = -1
const QWizard__NStandardButtons QWizard__WizardButton = 6
const QWizard__NButtons QWizard__WizardButton = 9

type QWizard__WizardPixmap = int

const QWizard__WatermarkPixmap QWizard__WizardPixmap = 0
const QWizard__LogoPixmap QWizard__WizardPixmap = 1
const QWizard__BannerPixmap QWizard__WizardPixmap = 2
const QWizard__BackgroundPixmap QWizard__WizardPixmap = 3
const QWizard__NPixmaps QWizard__WizardPixmap = 4

type QWizard__WizardStyle = int

const QWizard__ClassicStyle QWizard__WizardStyle = 0
const QWizard__ModernStyle QWizard__WizardStyle = 1
const QWizard__MacStyle QWizard__WizardStyle = 2
const QWizard__AeroStyle QWizard__WizardStyle = 3
const QWizard__NStyles QWizard__WizardStyle = 4

type QWizard__WizardOption = int

const QWizard__IndependentPages QWizard__WizardOption = 1
const QWizard__IgnoreSubTitles QWizard__WizardOption = 2
const QWizard__ExtendedWatermarkPixmap QWizard__WizardOption = 4
const QWizard__NoDefaultButton QWizard__WizardOption = 8
const QWizard__NoBackButtonOnStartPage QWizard__WizardOption = 16
const QWizard__NoBackButtonOnLastPage QWizard__WizardOption = 32
const QWizard__DisabledBackButtonOnLastPage QWizard__WizardOption = 64
const QWizard__HaveNextButtonOnLastPage QWizard__WizardOption = 128
const QWizard__HaveFinishButtonOnEarlyPages QWizard__WizardOption = 256
const QWizard__NoCancelButton QWizard__WizardOption = 512
const QWizard__CancelButtonOnLeft QWizard__WizardOption = 1024
const QWizard__HaveHelpButton QWizard__WizardOption = 2048
const QWizard__HelpButtonOnRight QWizard__WizardOption = 4096
const QWizard__HaveCustomButton1 QWizard__WizardOption = 8192
const QWizard__HaveCustomButton2 QWizard__WizardOption = 16384
const QWizard__HaveCustomButton3 QWizard__WizardOption = 32768
const QWizard__NoCancelButtonOnLastPage QWizard__WizardOption = 65536

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
