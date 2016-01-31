package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qwizard.h
// dst-file: /src/widgets/qwizard.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QWizardPage::setTitle(const QString & title);
extern void C_ZN11QWizardPage8setTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QWizardPage::isCommitPage();
extern void C_ZNK11QWizardPage12isCommitPageEv(void* qthis); // 4
  // proto:  void QWizardPage::cleanupPage();
extern void C_ZN11QWizardPage11cleanupPageEv(void* qthis); // 4
  // proto:  void QWizardPage::~QWizardPage();
extern void C_ZN11QWizardPageD2Ev(void* qthis); // 4
  // proto:  QString QWizardPage::subTitle();
extern void C_ZNK11QWizardPage8subTitleEv(void* qthis); // 4
  // proto:  QString QWizardPage::title();
extern void C_ZNK11QWizardPage5titleEv(void* qthis); // 4
  // proto:  bool QWizardPage::isComplete();
extern void C_ZNK11QWizardPage10isCompleteEv(void* qthis); // 4
  // proto:  bool QWizardPage::isFinalPage();
extern void C_ZNK11QWizardPage11isFinalPageEv(void* qthis); // 4
  // proto:  bool QWizardPage::validatePage();
extern void C_ZN11QWizardPage12validatePageEv(void* qthis); // 4
  // proto:  void QWizardPage::setSubTitle(const QString & subTitle);
extern void C_ZN11QWizardPage11setSubTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWizardPage::initializePage();
extern void C_ZN11QWizardPage14initializePageEv(void* qthis); // 4
  // proto:  const QMetaObject * QWizardPage::metaObject();
extern void C_ZNK11QWizardPage10metaObjectEv(void* qthis); // 4
  // proto:  void QWizardPage::QWizardPage(QWidget * parent);
extern void* C_ZN11QWizardPageC2EP7QWidget(void* arg0); // 3
  // proto:  void QWizardPage::setCommitPage(bool commitPage);
extern void C_ZN11QWizardPage13setCommitPageEb(void* qthis, bool arg0); // 4
  // proto:  int QWizardPage::nextId();
extern void C_ZNK11QWizardPage6nextIdEv(void* qthis); // 4
  // proto:  void QWizardPage::setFinalPage(bool finalPage);
extern void C_ZN11QWizardPage12setFinalPageEb(void* qthis, bool arg0); // 4
  // proto:  bool QWizard::validateCurrentPage();
extern void C_ZN7QWizard19validateCurrentPageEv(void* qthis); // 4
  // proto:  void QWizard::setSideWidget(QWidget * widget);
extern void C_ZN7QWizard13setSideWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QWizard::back();
extern void C_ZN7QWizard4backEv(void* qthis); // 4
  // proto:  void QWizard::setDefaultProperty(const char * className, const char * property, const char * changedSignal);
extern void C_ZN7QWizard18setDefaultPropertyEPKcS1_S1_(void* qthis, unsigned char* arg0, unsigned char* arg1, unsigned char* arg2); // 4
  // proto:  void QWizard::setVisible(bool visible);
extern void C_ZN7QWizard10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QWizard::setField(const QString & name, const QVariant & value);
extern void C_ZN7QWizard8setFieldERK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QList<int> QWizard::visitedPages();
extern void C_ZNK7QWizard12visitedPagesEv(void* qthis); // 4
  // proto:  QVariant QWizard::field(const QString & name);
extern void C_ZNK7QWizard5fieldERK7QString(void* qthis, void* arg0); // 4
  // proto:  Qt::TextFormat QWizard::titleFormat();
extern void C_ZNK7QWizard11titleFormatEv(void* qthis); // 4
  // proto:  void QWizard::next();
extern void C_ZN7QWizard4nextEv(void* qthis); // 4
  // proto:  QList<int> QWizard::pageIds();
extern void C_ZNK7QWizard7pageIdsEv(void* qthis); // 4
  // proto:  Qt::TextFormat QWizard::subTitleFormat();
extern void C_ZNK7QWizard14subTitleFormatEv(void* qthis); // 4
  // proto:  void QWizard::setPage(int id, QWizardPage * page);
extern void C_ZN7QWizard7setPageEiP11QWizardPage(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QWizardPage * QWizard::currentPage();
extern void C_ZNK7QWizard11currentPageEv(void* qthis); // 4
  // proto:  bool QWizard::hasVisitedPage(int id);
extern void C_ZNK7QWizard14hasVisitedPageEi(void* qthis, int32_t arg0); // 4
  // proto:  QWizard::WizardStyle QWizard::wizardStyle();
extern void C_ZNK7QWizard11wizardStyleEv(void* qthis); // 4
  // proto:  const QMetaObject * QWizard::metaObject();
extern void C_ZNK7QWizard10metaObjectEv(void* qthis); // 4
  // proto:  void QWizard::~QWizard();
extern void C_ZN7QWizardD2Ev(void* qthis); // 4
  // proto:  int QWizard::startId();
extern void C_ZNK7QWizard7startIdEv(void* qthis); // 4
  // proto:  void QWizard::restart();
extern void C_ZN7QWizard7restartEv(void* qthis); // 4
  // proto:  QWidget * QWizard::sideWidget();
extern void C_ZNK7QWizard10sideWidgetEv(void* qthis); // 4
  // proto:  QSize QWizard::sizeHint();
extern void C_ZNK7QWizard8sizeHintEv(void* qthis); // 4
  // proto:  int QWizard::currentId();
extern void C_ZNK7QWizard9currentIdEv(void* qthis); // 4
  // proto:  void QWizard::removePage(int id);
extern void C_ZN7QWizard10removePageEi(void* qthis, int32_t arg0); // 4
  // proto:  WizardOptions QWizard::options();
extern void C_ZNK7QWizard7optionsEv(void* qthis); // 4
  // proto:  void QWizard::setStartId(int id);
extern void C_ZN7QWizard10setStartIdEi(void* qthis, int32_t arg0); // 4
  // proto:  int QWizard::nextId();
extern void C_ZNK7QWizard6nextIdEv(void* qthis); // 4
  // proto:  QWizardPage * QWizard::page(int id);
extern void C_ZNK7QWizard4pageEi(void* qthis, int32_t arg0); // 4
  // proto:  int QWizard::addPage(QWizardPage * page);
extern void C_ZN7QWizard7addPageEP11QWizardPage(void* qthis, void* arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QWizardPage)=1
type QWizardPage struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _completeChanged QWizardPage_completeChanged_signal;
}

// class sizeof(QWizard)=1
type QWizard struct {
  /*qbase*/ QDialog;
  qclsinst unsafe.Pointer /* *C.void */;
//  _pageRemoved QWizard_pageRemoved_signal;
//  _pageAdded QWizard_pageAdded_signal;
//  _helpRequested QWizard_helpRequested_signal;
//  _currentIdChanged QWizard_currentIdChanged_signal;
//  _customButtonClicked QWizard_customButtonClicked_signal;
}

// setTitle(const class QString &)
func (this *QWizardPage) setTitle(args ...interface{}) () {
  // setTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWizardPage8setTitleERK7QString
    // invoke: void setTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QWizardPage8setTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizardPage", "setTitle", args)
  }

}

// isCommitPage()
func (this *QWizardPage) isCommitPage(args ...interface{}) () {
  // isCommitPage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWizardPage12isCommitPageEv
    // invoke: bool isCommitPage()
    var ret = C.C_ZNK11QWizardPage12isCommitPageEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizardPage", "isCommitPage", args)
  }

}

// cleanupPage()
func (this *QWizardPage) cleanupPage(args ...interface{}) () {
  // cleanupPage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWizardPage11cleanupPageEv
    // invoke: void cleanupPage()
    C.C_ZN11QWizardPage11cleanupPageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "cleanupPage", args)
  }

}

// ~QWizardPage()
func (this *QWizardPage) FreeQWizardPage(args ...interface{}) () {
  // ~QWizardPage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWizardPageD0Ev
    // invoke: void ~QWizardPage()
    C.C_ZN11QWizardPageD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "~QWizardPage", args)
  }

}

// subTitle()
func (this *QWizardPage) subTitle(args ...interface{}) () {
  // subTitle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWizardPage8subTitleEv
    // invoke: QString subTitle()
    var ret = C.C_ZNK11QWizardPage8subTitleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizardPage", "subTitle", args)
  }

}

// title()
func (this *QWizardPage) title(args ...interface{}) () {
  // title()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWizardPage5titleEv
    // invoke: QString title()
    var ret = C.C_ZNK11QWizardPage5titleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizardPage", "title", args)
  }

}

// isComplete()
func (this *QWizardPage) isComplete(args ...interface{}) () {
  // isComplete()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWizardPage10isCompleteEv
    // invoke: bool isComplete()
    var ret = C.C_ZNK11QWizardPage10isCompleteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizardPage", "isComplete", args)
  }

}

// isFinalPage()
func (this *QWizardPage) isFinalPage(args ...interface{}) () {
  // isFinalPage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWizardPage11isFinalPageEv
    // invoke: bool isFinalPage()
    var ret = C.C_ZNK11QWizardPage11isFinalPageEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizardPage", "isFinalPage", args)
  }

}

// validatePage()
func (this *QWizardPage) validatePage(args ...interface{}) () {
  // validatePage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWizardPage12validatePageEv
    // invoke: bool validatePage()
    var ret = C.C_ZN11QWizardPage12validatePageEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizardPage", "validatePage", args)
  }

}

// setSubTitle(const class QString &)
func (this *QWizardPage) setSubTitle(args ...interface{}) () {
  // setSubTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWizardPage11setSubTitleERK7QString
    // invoke: void setSubTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QWizardPage11setSubTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizardPage", "setSubTitle", args)
  }

}

// initializePage()
func (this *QWizardPage) initializePage(args ...interface{}) () {
  // initializePage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWizardPage14initializePageEv
    // invoke: void initializePage()
    C.C_ZN11QWizardPage14initializePageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "initializePage", args)
  }

}

// metaObject()
func (this *QWizardPage) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWizardPage10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QWizardPage10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "metaObject", args)
  }

}

// QWizardPage(class QWidget *)
func NewQWizardPage(args ...interface{}) *QWizardPage {
  // QWizardPage(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWizardPageC1EP7QWidget
    // invoke: void QWizardPage(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QWizardPageC2EP7QWidget(arg0)
    return &QWizardPage{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QWizardPage", "QWizardPage", args)
  }

  return nil // QWizardPage{qclsinst:qthis}
}

// setCommitPage(_Bool)
func (this *QWizardPage) setCommitPage(args ...interface{}) () {
  // setCommitPage(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWizardPage13setCommitPageEb
    // invoke: void setCommitPage(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QWizardPage13setCommitPageEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizardPage", "setCommitPage", args)
  }

}

// nextId()
func (this *QWizardPage) nextId(args ...interface{}) () {
  // nextId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWizardPage6nextIdEv
    // invoke: int nextId()
    var ret = C.C_ZNK11QWizardPage6nextIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizardPage", "nextId", args)
  }

}

// setFinalPage(_Bool)
func (this *QWizardPage) setFinalPage(args ...interface{}) () {
  // setFinalPage(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWizardPage12setFinalPageEb
    // invoke: void setFinalPage(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QWizardPage12setFinalPageEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizardPage", "setFinalPage", args)
  }

}

// validateCurrentPage()
func (this *QWizard) validateCurrentPage(args ...interface{}) () {
  // validateCurrentPage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard19validateCurrentPageEv
    // invoke: bool validateCurrentPage()
    var ret = C.C_ZN7QWizard19validateCurrentPageEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "validateCurrentPage", args)
  }

}

// setSideWidget(class QWidget *)
func (this *QWizard) setSideWidget(args ...interface{}) () {
  // setSideWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard13setSideWidgetEP7QWidget
    // invoke: void setSideWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWizard13setSideWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "setSideWidget", args)
  }

}

// back()
func (this *QWizard) back(args ...interface{}) () {
  // back()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard4backEv
    // invoke: void back()
    C.C_ZN7QWizard4backEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "back", args)
  }

}

// setDefaultProperty(const char *, const char *, const char *)
func (this *QWizard) setDefaultProperty(args ...interface{}) () {
  // setDefaultProperty(const char *, const char *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard18setDefaultPropertyEPKcS1_S1_
    // invoke: void setDefaultProperty(const char *, const char *, const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    C.C_ZN7QWizard18setDefaultPropertyEPKcS1_S1_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QWizard", "setDefaultProperty", args)
  }

}

// setVisible(_Bool)
func (this *QWizard) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWizard10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "setVisible", args)
  }

}

// setField(const class QString &, const class QVariant &)
func (this *QWizard) setField(args ...interface{}) () {
  // setField(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard8setFieldERK7QStringRK8QVariant
    // invoke: void setField(const class QString &, const class QVariant &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QWizard8setFieldERK7QStringRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWizard", "setField", args)
  }

}

// visitedPages()
func (this *QWizard) visitedPages(args ...interface{}) () {
  // visitedPages()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard12visitedPagesEv
    // invoke: QList<int> visitedPages()
    C.C_ZNK7QWizard12visitedPagesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "visitedPages", args)
  }

}

// field(const class QString &)
func (this *QWizard) field(args ...interface{}) () {
  // field(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard5fieldERK7QString
    // invoke: QVariant field(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QWizard5fieldERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "field", args)
  }

}

// titleFormat()
func (this *QWizard) titleFormat(args ...interface{}) () {
  // titleFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard11titleFormatEv
    // invoke: Qt::TextFormat titleFormat()
    C.C_ZNK7QWizard11titleFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "titleFormat", args)
  }

}

// next()
func (this *QWizard) next(args ...interface{}) () {
  // next()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard4nextEv
    // invoke: void next()
    C.C_ZN7QWizard4nextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "next", args)
  }

}

// pageIds()
func (this *QWizard) pageIds(args ...interface{}) () {
  // pageIds()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard7pageIdsEv
    // invoke: QList<int> pageIds()
    C.C_ZNK7QWizard7pageIdsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "pageIds", args)
  }

}

// subTitleFormat()
func (this *QWizard) subTitleFormat(args ...interface{}) () {
  // subTitleFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard14subTitleFormatEv
    // invoke: Qt::TextFormat subTitleFormat()
    C.C_ZNK7QWizard14subTitleFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "subTitleFormat", args)
  }

}

// setPage(int, class QWizardPage *)
func (this *QWizard) setPage(args ...interface{}) () {
  // setPage(int, class QWizardPage *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWizardPage{}) // "QWizardPage *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard7setPageEiP11QWizardPage
    // invoke: void setPage(int, class QWizardPage *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWizardPage).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QWizard7setPageEiP11QWizardPage(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWizard", "setPage", args)
  }

}

// currentPage()
func (this *QWizard) currentPage(args ...interface{}) () {
  // currentPage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard11currentPageEv
    // invoke: QWizardPage * currentPage()
    var ret = C.C_ZNK7QWizard11currentPageEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "currentPage", args)
  }

}

// hasVisitedPage(int)
func (this *QWizard) hasVisitedPage(args ...interface{}) () {
  // hasVisitedPage(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard14hasVisitedPageEi
    // invoke: bool hasVisitedPage(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QWizard14hasVisitedPageEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "hasVisitedPage", args)
  }

}

// wizardStyle()
func (this *QWizard) wizardStyle(args ...interface{}) () {
  // wizardStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard11wizardStyleEv
    // invoke: QWizard::WizardStyle wizardStyle()
    C.C_ZNK7QWizard11wizardStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "wizardStyle", args)
  }

}

// metaObject()
func (this *QWizard) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QWizard10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "metaObject", args)
  }

}

// ~QWizard()
func (this *QWizard) FreeQWizard(args ...interface{}) () {
  // ~QWizard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizardD0Ev
    // invoke: void ~QWizard()
    C.C_ZN7QWizardD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "~QWizard", args)
  }

}

// startId()
func (this *QWizard) startId(args ...interface{}) () {
  // startId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard7startIdEv
    // invoke: int startId()
    var ret = C.C_ZNK7QWizard7startIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "startId", args)
  }

}

// restart()
func (this *QWizard) restart(args ...interface{}) () {
  // restart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard7restartEv
    // invoke: void restart()
    C.C_ZN7QWizard7restartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "restart", args)
  }

}

// sideWidget()
func (this *QWizard) sideWidget(args ...interface{}) () {
  // sideWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard10sideWidgetEv
    // invoke: QWidget * sideWidget()
    var ret = C.C_ZNK7QWizard10sideWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "sideWidget", args)
  }

}

// sizeHint()
func (this *QWizard) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard8sizeHintEv
    // invoke: QSize sizeHint()
    var ret = C.C_ZNK7QWizard8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "sizeHint", args)
  }

}

// currentId()
func (this *QWizard) currentId(args ...interface{}) () {
  // currentId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard9currentIdEv
    // invoke: int currentId()
    var ret = C.C_ZNK7QWizard9currentIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "currentId", args)
  }

}

// removePage(int)
func (this *QWizard) removePage(args ...interface{}) () {
  // removePage(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard10removePageEi
    // invoke: void removePage(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWizard10removePageEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "removePage", args)
  }

}

// options()
func (this *QWizard) options(args ...interface{}) () {
  // options()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard7optionsEv
    // invoke: WizardOptions options()
    C.C_ZNK7QWizard7optionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "options", args)
  }

}

// setStartId(int)
func (this *QWizard) setStartId(args ...interface{}) () {
  // setStartId(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard10setStartIdEi
    // invoke: void setStartId(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWizard10setStartIdEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "setStartId", args)
  }

}

// nextId()
func (this *QWizard) nextId(args ...interface{}) () {
  // nextId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard6nextIdEv
    // invoke: int nextId()
    var ret = C.C_ZNK7QWizard6nextIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "nextId", args)
  }

}

// page(int)
func (this *QWizard) page(args ...interface{}) () {
  // page(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWizard4pageEi
    // invoke: QWizardPage * page(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK7QWizard4pageEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "page", args)
  }

}

// addPage(class QWizardPage *)
func (this *QWizard) addPage(args ...interface{}) () {
  // addPage(class QWizardPage *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWizardPage{}) // "QWizardPage *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWizard7addPageEP11QWizardPage
    // invoke: int addPage(class QWizardPage *)
    var arg0 = args[0].(QWizardPage).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN7QWizard7addPageEP11QWizardPage(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QWizard", "addPage", args)
  }

}

// <= body block end

