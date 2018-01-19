//  header block begin
// /usr/include/qt/QtWidgets/qwizard.h
// #include <qwizard.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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
type QWizardPage struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qwizard.h:213
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QWizardPage) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:218
// index:0
// void QWizardPage(class QWidget *)
func NewQWizardPage(parent unsafe.Pointer) *QWizardPage {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPageC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QWizardPage{cthis}
}

// /usr/include/qt/QtWidgets/qwizard.h:219
// index:0
// virtual
// void ~QWizardPage()
func DeleteQWizardPage(*QWizardPage) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPageD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:221
// index:0
// void setTitle(const class QString &)
func (this *QWizardPage) SetTitle(title unsafe.Pointer) {
	// 0: (, const QString & title), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage8setTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:222
// index:0
// QString title()
func (this *QWizardPage) Title() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage5titleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:223
// index:0
// void setSubTitle(const class QString &)
func (this *QWizardPage) SetSubTitle(subTitle unsafe.Pointer) {
	// 0: (, const QString & subTitle), (subTitle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage11setSubTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, subTitle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:224
// index:0
// QString subTitle()
func (this *QWizardPage) SubTitle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage8subTitleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:225
// index:0
// void setPixmap(class QWizard::WizardPixmap, const class QPixmap &)
func (this *QWizardPage) SetPixmap(which int, pixmap unsafe.Pointer) {
	// 0: (, QWizard::WizardPixmap which, const QPixmap & pixmap), (&which, pixmap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage9setPixmapEN7QWizard12WizardPixmapERK7QPixmap", ffiqt.FFI_TYPE_VOID, this.cthis, &which, pixmap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:226
// index:0
// QPixmap pixmap(class QWizard::WizardPixmap)
func (this *QWizardPage) Pixmap(which int) {
	// 0: (, QWizard::WizardPixmap which), (&which)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage6pixmapEN7QWizard12WizardPixmapE", ffiqt.FFI_TYPE_VOID, this.cthis, &which)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:227
// index:0
// void setFinalPage(_Bool)
func (this *QWizardPage) SetFinalPage(finalPage bool) {
	// 0: (, bool finalPage), (&finalPage)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage12setFinalPageEb", ffiqt.FFI_TYPE_VOID, this.cthis, &finalPage)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:228
// index:0
// bool isFinalPage()
func (this *QWizardPage) IsFinalPage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage11isFinalPageEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:229
// index:0
// void setCommitPage(_Bool)
func (this *QWizardPage) SetCommitPage(commitPage bool) {
	// 0: (, bool commitPage), (&commitPage)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage13setCommitPageEb", ffiqt.FFI_TYPE_VOID, this.cthis, &commitPage)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:230
// index:0
// bool isCommitPage()
func (this *QWizardPage) IsCommitPage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage12isCommitPageEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:231
// index:0
// void setButtonText(class QWizard::WizardButton, const class QString &)
func (this *QWizardPage) SetButtonText(which int, text unsafe.Pointer) {
	// 0: (, QWizard::WizardButton which, const QString & text), (&which, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage13setButtonTextEN7QWizard12WizardButtonERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &which, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:232
// index:0
// QString buttonText(class QWizard::WizardButton)
func (this *QWizardPage) ButtonText(which int) {
	// 0: (, QWizard::WizardButton which), (&which)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage10buttonTextEN7QWizard12WizardButtonE", ffiqt.FFI_TYPE_VOID, this.cthis, &which)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:234
// index:0
// virtual
// void initializePage()
func (this *QWizardPage) InitializePage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage14initializePageEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:235
// index:0
// virtual
// void cleanupPage()
func (this *QWizardPage) CleanupPage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage11cleanupPageEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:236
// index:0
// virtual
// bool validatePage()
func (this *QWizardPage) ValidatePage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage12validatePageEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:237
// index:0
// virtual
// bool isComplete()
func (this *QWizardPage) IsComplete() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage10isCompleteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:238
// index:0
// virtual
// int nextId()
func (this *QWizardPage) NextId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage6nextIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:241
// index:0
// void completeChanged()
func (this *QWizardPage) CompleteChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage15completeChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
