package qtwidgets

// /usr/include/qt/QtWidgets/qwizard.h
// #include <qwizard.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QWizard struct {
	*QDialog
}

func (this *QWizard) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func NewQWizardFromPointer(cthis unsafe.Pointer) *QWizard {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QWizard{bcthis0}
}

// /usr/include/qt/QtWidgets/qwizard.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QWizard) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:124
// index:0
// Public virtual
// void ~QWizard()
func DeleteQWizard(*QWizard) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizardD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:126
// index:0
// Public
// int addPage(class QWizardPage *)
func (this *QWizard) AddPage(page *QWizardPage /*444 QWizardPage **/) int {
	var convArg0 = page.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard7addPageEP11QWizardPage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwizard.h:127
// index:0
// Public
// void setPage(int, class QWizardPage *)
func (this *QWizard) SetPage(id int, page *QWizardPage /*444 QWizardPage **/) {
	var convArg1 = page.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard7setPageEiP11QWizardPage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:128
// index:0
// Public
// void removePage(int)
func (this *QWizard) RemovePage(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard10removePageEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:129
// index:0
// Public
// QWizardPage * page(int)
func (this *QWizard) Page(id int) *QWizardPage /*444 QWizardPage **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard4pageEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWizardPageFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:130
// index:0
// Public
// bool hasVisitedPage(int)
func (this *QWizard) HasVisitedPage(id int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard14hasVisitedPageEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:133
// index:0
// Public
// void setStartId(int)
func (this *QWizard) SetStartId(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard10setStartIdEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:134
// index:0
// Public
// int startId()
func (this *QWizard) StartId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard7startIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwizard.h:135
// index:0
// Public
// QWizardPage * currentPage()
func (this *QWizard) CurrentPage() *QWizardPage /*444 QWizardPage **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard11currentPageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWizardPageFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:136
// index:0
// Public
// int currentId()
func (this *QWizard) CurrentId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard9currentIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwizard.h:138
// index:0
// Public virtual
// bool validateCurrentPage()
func (this *QWizard) ValidateCurrentPage() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard19validateCurrentPageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:139
// index:0
// Public virtual
// int nextId()
func (this *QWizard) NextId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard6nextIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwizard.h:141
// index:0
// Public
// void setField(const class QString &, const class QVariant &)
func (this *QWizard) SetField(name *qtcore.QString, value *qtcore.QVariant) {
	var convArg0 = name.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard8setFieldERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:142
// index:0
// Public
// QVariant field(const class QString &)
func (this *QWizard) Field(name *qtcore.QString) *qtcore.QVariant /*123*/ {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard5fieldERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:144
// index:0
// Public
// void setWizardStyle(enum QWizard::WizardStyle)
func (this *QWizard) SetWizardStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard14setWizardStyleENS_11WizardStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:145
// index:0
// Public
// QWizard::WizardStyle wizardStyle()
func (this *QWizard) WizardStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard11wizardStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:147
// index:0
// Public
// void setOption(enum QWizard::WizardOption, _Bool)
func (this *QWizard) SetOption(option int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard9setOptionENS_12WizardOptionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:148
// index:0
// Public
// bool testOption(enum QWizard::WizardOption)
func (this *QWizard) TestOption(option int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard10testOptionENS_12WizardOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:150
// index:0
// Public
// QWizard::WizardOptions options()
func (this *QWizard) Options() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard7optionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:152
// index:0
// Public
// void setButtonText(enum QWizard::WizardButton, const class QString &)
func (this *QWizard) SetButtonText(which int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard13setButtonTextENS_12WizardButtonERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:153
// index:0
// Public
// QString buttonText(enum QWizard::WizardButton)
func (this *QWizard) ButtonText(which int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard10buttonTextENS_12WizardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:155
// index:0
// Public
// void setButton(enum QWizard::WizardButton, class QAbstractButton *)
func (this *QWizard) SetButton(which int, button *QAbstractButton /*444 QAbstractButton **/) {
	var convArg1 = button.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard9setButtonENS_12WizardButtonEP15QAbstractButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:156
// index:0
// Public
// QAbstractButton * button(enum QWizard::WizardButton)
func (this *QWizard) Button(which int) *QAbstractButton /*444 QAbstractButton **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard6buttonENS_12WizardButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractButtonFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:158
// index:0
// Public
// void setTitleFormat(Qt::TextFormat)
func (this *QWizard) SetTitleFormat(format int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard14setTitleFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:159
// index:0
// Public
// Qt::TextFormat titleFormat()
func (this *QWizard) TitleFormat() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard11titleFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:160
// index:0
// Public
// void setSubTitleFormat(Qt::TextFormat)
func (this *QWizard) SetSubTitleFormat(format int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard17setSubTitleFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:161
// index:0
// Public
// Qt::TextFormat subTitleFormat()
func (this *QWizard) SubTitleFormat() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard14subTitleFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:162
// index:0
// Public
// void setPixmap(enum QWizard::WizardPixmap, const class QPixmap &)
func (this *QWizard) SetPixmap(which int, pixmap *qtgui.QPixmap) {
	var convArg1 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard9setPixmapENS_12WizardPixmapERK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:163
// index:0
// Public
// QPixmap pixmap(enum QWizard::WizardPixmap)
func (this *QWizard) Pixmap(which int) *qtgui.QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard6pixmapENS_12WizardPixmapE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:165
// index:0
// Public
// void setSideWidget(class QWidget *)
func (this *QWizard) SetSideWidget(widget *QWidget /*444 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard13setSideWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:166
// index:0
// Public
// QWidget * sideWidget()
func (this *QWizard) SideWidget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard10sideWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:168
// index:0
// Public
// void setDefaultProperty(const char *, const char *, const char *)
func (this *QWizard) SetDefaultProperty(className string, property string, changedSignal string) {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(property)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(changedSignal)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard18setDefaultPropertyEPKcS1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:171
// index:0
// Public virtual
// void setVisible(_Bool)
func (this *QWizard) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:172
// index:0
// Public virtual
// QSize sizeHint()
func (this *QWizard) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:175
// index:0
// Public
// void currentIdChanged(int)
func (this *QWizard) CurrentIdChanged(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard16currentIdChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:176
// index:0
// Public
// void helpRequested()
func (this *QWizard) HelpRequested() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard13helpRequestedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:177
// index:0
// Public
// void customButtonClicked(int)
func (this *QWizard) CustomButtonClicked(which int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard19customButtonClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:178
// index:0
// Public
// void pageAdded(int)
func (this *QWizard) PageAdded(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard9pageAddedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:179
// index:0
// Public
// void pageRemoved(int)
func (this *QWizard) PageRemoved(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard11pageRemovedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:182
// index:0
// Public
// void back()
func (this *QWizard) Back() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard4backEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:183
// index:0
// Public
// void next()
func (this *QWizard) Next() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard4nextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:184
// index:0
// Public
// void restart()
func (this *QWizard) Restart() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard7restartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:187
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QWizard) Event(event *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:188
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QWizard) ResizeEvent(event *qtgui.QResizeEvent /*444 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:189
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QWizard) PaintEvent(event *qtgui.QPaintEvent /*444 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:193
// index:0
// Protected virtual
// void done(int)
func (this *QWizard) Done(result int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard4doneEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &result)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:194
// index:0
// Protected virtual
// void initializePage(int)
func (this *QWizard) InitializePage(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard14initializePageEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:195
// index:0
// Protected virtual
// void cleanupPage(int)
func (this *QWizard) CleanupPage(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard11cleanupPageEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
}

//  body block end
