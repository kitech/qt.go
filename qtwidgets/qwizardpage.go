package qtwidgets

// /usr/include/qt/QtWidgets/qwizard.h
// #include <qwizard.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 50
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	*QWidget
}

func (this *QWizardPage) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QWizardPage) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQWizardPageFromPointer(cthis unsafe.Pointer) *QWizardPage {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QWizardPage{bcthis0}
}
func (*QWizardPage) NewFromPointer(cthis unsafe.Pointer) *QWizardPage {
	return NewQWizardPageFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qwizard.h:213
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QWizardPage) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:218
// index:0
// Public
// void QWizardPage(class QWidget *)
func NewQWizardPage(parent *QWidget /*777 QWidget **/) *QWizardPage {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPageC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWizardPageFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qwizard.h:219
// index:0
// Public virtual
// void ~QWizardPage()
func DeleteQWizardPage(*QWizardPage) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPageD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:221
// index:0
// Public
// void setTitle(const class QString &)
func (this *QWizardPage) SetTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage8setTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:222
// index:0
// Public
// QString title()
func (this *QWizardPage) Title() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage5titleEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:223
// index:0
// Public
// void setSubTitle(const class QString &)
func (this *QWizardPage) SetSubTitle(subTitle *qtcore.QString) {
	var convArg0 = subTitle.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage11setSubTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:224
// index:0
// Public
// QString subTitle()
func (this *QWizardPage) SubTitle() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage8subTitleEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:225
// index:0
// Public
// void setPixmap(class QWizard::WizardPixmap, const class QPixmap &)
func (this *QWizardPage) SetPixmap(which int, pixmap *qtgui.QPixmap) {
	var convArg1 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage9setPixmapEN7QWizard12WizardPixmapERK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:226
// index:0
// Public
// QPixmap pixmap(class QWizard::WizardPixmap)
func (this *QWizardPage) Pixmap(which int) *qtgui.QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage6pixmapEN7QWizard12WizardPixmapE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), which)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:227
// index:0
// Public
// void setFinalPage(_Bool)
func (this *QWizardPage) SetFinalPage(finalPage bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage12setFinalPageEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), finalPage)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:228
// index:0
// Public
// bool isFinalPage()
func (this *QWizardPage) IsFinalPage() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage11isFinalPageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:229
// index:0
// Public
// void setCommitPage(_Bool)
func (this *QWizardPage) SetCommitPage(commitPage bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage13setCommitPageEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), commitPage)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:230
// index:0
// Public
// bool isCommitPage()
func (this *QWizardPage) IsCommitPage() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage12isCommitPageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:231
// index:0
// Public
// void setButtonText(class QWizard::WizardButton, const class QString &)
func (this *QWizardPage) SetButtonText(which int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage13setButtonTextEN7QWizard12WizardButtonERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:232
// index:0
// Public
// QString buttonText(class QWizard::WizardButton)
func (this *QWizardPage) ButtonText(which int) *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage10buttonTextEN7QWizard12WizardButtonE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), which)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:234
// index:0
// Public virtual
// void initializePage()
func (this *QWizardPage) InitializePage() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage14initializePageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:235
// index:0
// Public virtual
// void cleanupPage()
func (this *QWizardPage) CleanupPage() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage11cleanupPageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:236
// index:0
// Public virtual
// bool validatePage()
func (this *QWizardPage) ValidatePage() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage12validatePageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:237
// index:0
// Public virtual
// bool isComplete()
func (this *QWizardPage) IsComplete() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage10isCompleteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:238
// index:0
// Public virtual
// int nextId()
func (this *QWizardPage) NextId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage6nextIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwizard.h:241
// index:0
// Public
// void completeChanged()
func (this *QWizardPage) CompleteChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage15completeChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:244
// index:0
// Protected
// void setField(const class QString &, const class QVariant &)
func (this *QWizardPage) SetField(name *qtcore.QString, value *qtcore.QVariant) {
	var convArg0 = name.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage8setFieldERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:245
// index:0
// Protected
// QVariant field(const class QString &)
func (this *QWizardPage) Field(name *qtcore.QString) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage5fieldERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:246
// index:0
// Protected
// void registerField(const class QString &, class QWidget *, const char *, const char *)
func (this *QWizardPage) RegisterField(name *qtcore.QString, widget *QWidget /*777 QWidget **/, property string, changedSignal string) {
	var convArg0 = name.GetCthis()
	var convArg1 = widget.GetCthis()
	var convArg2 = qtrt.CString(property)
	defer qtrt.FreeMem(convArg2)
	var convArg3 = qtrt.CString(changedSignal)
	defer qtrt.FreeMem(convArg3)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWizardPage13registerFieldERK7QStringP7QWidgetPKcS6_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:248
// index:0
// Protected
// QWizard * wizard()
func (this *QWizardPage) Wizard() *QWizard /*777 QWizard **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWizardPage6wizardEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWizardFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
