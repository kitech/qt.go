//  header block begin
// /usr/include/qt/QtWidgets/qwizard.h
// #include <qwizard.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qwizard.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QWizard) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:124
// index:0
// virtual
// void ~QWizard()
func DeleteQWizard(*QWizard) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizardD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:126
// index:0
// int addPage(class QWizardPage *)
func (this *QWizard) AddPage(page unsafe.Pointer) {
	// 0: (, QWizardPage * page), (page)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard7addPageEP11QWizardPage", ffiqt.FFI_TYPE_VOID, this.cthis, page)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:127
// index:0
// void setPage(int, class QWizardPage *)
func (this *QWizard) SetPage(id int, page unsafe.Pointer) {
	// 0: (, int id, QWizardPage * page), (&id, page)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard7setPageEiP11QWizardPage", ffiqt.FFI_TYPE_VOID, this.cthis, &id, page)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:128
// index:0
// void removePage(int)
func (this *QWizard) RemovePage(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard10removePageEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:129
// index:0
// QWizardPage * page(int)
func (this *QWizard) Page(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard4pageEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:130
// index:0
// bool hasVisitedPage(int)
func (this *QWizard) HasVisitedPage(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard14hasVisitedPageEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:131
// index:0
// QList<int> visitedPages()
func (this *QWizard) VisitedPages() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard12visitedPagesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:132
// index:0
// QList<int> pageIds()
func (this *QWizard) PageIds() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard7pageIdsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:133
// index:0
// void setStartId(int)
func (this *QWizard) SetStartId(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard10setStartIdEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:134
// index:0
// int startId()
func (this *QWizard) StartId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard7startIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:135
// index:0
// QWizardPage * currentPage()
func (this *QWizard) CurrentPage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard11currentPageEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:136
// index:0
// int currentId()
func (this *QWizard) CurrentId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard9currentIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:138
// index:0
// virtual
// bool validateCurrentPage()
func (this *QWizard) ValidateCurrentPage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard19validateCurrentPageEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:139
// index:0
// virtual
// int nextId()
func (this *QWizard) NextId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard6nextIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:141
// index:0
// void setField(const class QString &, const class QVariant &)
func (this *QWizard) SetField(name unsafe.Pointer, value unsafe.Pointer) {
	// 0: (, const QString & name, const QVariant & value), (name, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard8setFieldERK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, name, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:142
// index:0
// QVariant field(const class QString &)
func (this *QWizard) Field(name unsafe.Pointer) {
	// 0: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard5fieldERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:144
// index:0
// void setWizardStyle(enum QWizard::WizardStyle)
func (this *QWizard) SetWizardStyle(style int) {
	// 0: (, QWizard::WizardStyle style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard14setWizardStyleENS_11WizardStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:145
// index:0
// QWizard::WizardStyle wizardStyle()
func (this *QWizard) WizardStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard11wizardStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:147
// index:0
// void setOption(enum QWizard::WizardOption, _Bool)
func (this *QWizard) SetOption(option int, on bool) {
	// 0: (, QWizard::WizardOption option, bool on), (&option, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard9setOptionENS_12WizardOptionEb", ffiqt.FFI_TYPE_VOID, this.cthis, &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:148
// index:0
// bool testOption(enum QWizard::WizardOption)
func (this *QWizard) TestOption(option int) {
	// 0: (, QWizard::WizardOption option), (&option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard10testOptionENS_12WizardOptionE", ffiqt.FFI_TYPE_VOID, this.cthis, &option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:150
// index:0
// QWizard::WizardOptions options()
func (this *QWizard) Options() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard7optionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:152
// index:0
// void setButtonText(enum QWizard::WizardButton, const class QString &)
func (this *QWizard) SetButtonText(which int, text unsafe.Pointer) {
	// 0: (, QWizard::WizardButton which, const QString & text), (&which, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard13setButtonTextENS_12WizardButtonERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &which, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:153
// index:0
// QString buttonText(enum QWizard::WizardButton)
func (this *QWizard) ButtonText(which int) {
	// 0: (, QWizard::WizardButton which), (&which)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard10buttonTextENS_12WizardButtonE", ffiqt.FFI_TYPE_VOID, this.cthis, &which)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:155
// index:0
// void setButton(enum QWizard::WizardButton, class QAbstractButton *)
func (this *QWizard) SetButton(which int, button unsafe.Pointer) {
	// 0: (, QWizard::WizardButton which, QAbstractButton * button), (&which, button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard9setButtonENS_12WizardButtonEP15QAbstractButton", ffiqt.FFI_TYPE_VOID, this.cthis, &which, button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:156
// index:0
// QAbstractButton * button(enum QWizard::WizardButton)
func (this *QWizard) Button(which int) {
	// 0: (, QWizard::WizardButton which), (&which)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard6buttonENS_12WizardButtonE", ffiqt.FFI_TYPE_VOID, this.cthis, &which)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:158
// index:0
// void setTitleFormat(Qt::TextFormat)
func (this *QWizard) SetTitleFormat(format int) {
	// 0: (, Qt::TextFormat format), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard14setTitleFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_VOID, this.cthis, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:159
// index:0
// Qt::TextFormat titleFormat()
func (this *QWizard) TitleFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard11titleFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:160
// index:0
// void setSubTitleFormat(Qt::TextFormat)
func (this *QWizard) SetSubTitleFormat(format int) {
	// 0: (, Qt::TextFormat format), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard17setSubTitleFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_VOID, this.cthis, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:161
// index:0
// Qt::TextFormat subTitleFormat()
func (this *QWizard) SubTitleFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard14subTitleFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:162
// index:0
// void setPixmap(enum QWizard::WizardPixmap, const class QPixmap &)
func (this *QWizard) SetPixmap(which int, pixmap unsafe.Pointer) {
	// 0: (, QWizard::WizardPixmap which, const QPixmap & pixmap), (&which, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard9setPixmapENS_12WizardPixmapERK7QPixmap", ffiqt.FFI_TYPE_VOID, this.cthis, &which, pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:163
// index:0
// QPixmap pixmap(enum QWizard::WizardPixmap)
func (this *QWizard) Pixmap(which int) {
	// 0: (, QWizard::WizardPixmap which), (&which)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard6pixmapENS_12WizardPixmapE", ffiqt.FFI_TYPE_VOID, this.cthis, &which)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:165
// index:0
// void setSideWidget(class QWidget *)
func (this *QWizard) SetSideWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard13setSideWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:166
// index:0
// QWidget * sideWidget()
func (this *QWizard) SideWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard10sideWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:168
// index:0
// void setDefaultProperty(const char *, const char *, const char *)
func (this *QWizard) SetDefaultProperty(className unsafe.Pointer, property unsafe.Pointer, changedSignal unsafe.Pointer) {
	// 0: (, const char * className, const char * property, const char * changedSignal), (className, property, changedSignal)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard18setDefaultPropertyEPKcS1_S1_", ffiqt.FFI_TYPE_VOID, this.cthis, className, property, changedSignal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:171
// index:0
// virtual
// void setVisible(_Bool)
func (this *QWizard) SetVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:172
// index:0
// virtual
// QSize sizeHint()
func (this *QWizard) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWizard8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:175
// index:0
// void currentIdChanged(int)
func (this *QWizard) CurrentIdChanged(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard16currentIdChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:176
// index:0
// void helpRequested()
func (this *QWizard) HelpRequested() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard13helpRequestedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:177
// index:0
// void customButtonClicked(int)
func (this *QWizard) CustomButtonClicked(which int) {
	// 0: (, int which), (&which)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard19customButtonClickedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &which)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:178
// index:0
// void pageAdded(int)
func (this *QWizard) PageAdded(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard9pageAddedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:179
// index:0
// void pageRemoved(int)
func (this *QWizard) PageRemoved(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard11pageRemovedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:182
// index:0
// void back()
func (this *QWizard) Back() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard4backEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:183
// index:0
// void next()
func (this *QWizard) Next() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard4nextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:184
// index:0
// void restart()
func (this *QWizard) Restart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWizard7restartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
