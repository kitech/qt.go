package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern bool C_ZNK11QWizardPage12isCommitPageEv(void* qthis); // 4
  // proto:  void QWizardPage::cleanupPage();
extern void C_ZN11QWizardPage11cleanupPageEv(void* qthis); // 4
  // proto:  void QWizardPage::~QWizardPage();
extern void C_ZN11QWizardPageD2Ev(void* qthis); // 4
  // proto:  QString QWizardPage::subTitle();
extern void* C_ZNK11QWizardPage8subTitleEv(void* qthis); // 4
  // proto:  QString QWizardPage::title();
extern void* C_ZNK11QWizardPage5titleEv(void* qthis); // 4
  // proto:  bool QWizardPage::isComplete();
extern bool C_ZNK11QWizardPage10isCompleteEv(void* qthis); // 4
  // proto:  bool QWizardPage::isFinalPage();
extern bool C_ZNK11QWizardPage11isFinalPageEv(void* qthis); // 4
  // proto:  bool QWizardPage::validatePage();
extern bool C_ZN11QWizardPage12validatePageEv(void* qthis); // 4
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
extern int32_t C_ZNK11QWizardPage6nextIdEv(void* qthis); // 4
  // proto:  void QWizardPage::setFinalPage(bool finalPage);
extern void C_ZN11QWizardPage12setFinalPageEb(void* qthis, bool arg0); // 4
  // proto:  bool QWizard::validateCurrentPage();
extern bool C_ZN7QWizard19validateCurrentPageEv(void* qthis); // 4
  // proto:  void QWizard::setSideWidget(QWidget * widget);
extern void C_ZN7QWizard13setSideWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QWizard::back();
extern void C_ZN7QWizard4backEv(void* qthis); // 4
  // proto:  void QWizard::setDefaultProperty(const char * className, const char * property, const char * changedSignal);
extern void C_ZN7QWizard18setDefaultPropertyEPKcS1_S1_(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QWizard::setVisible(bool visible);
extern void C_ZN7QWizard10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QWizard::setField(const QString & name, const QVariant & value);
extern void C_ZN7QWizard8setFieldERK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QList<int> QWizard::visitedPages();
extern void C_ZNK7QWizard12visitedPagesEv(void* qthis); // 4
  // proto:  QVariant QWizard::field(const QString & name);
extern void* C_ZNK7QWizard5fieldERK7QString(void* qthis, void* arg0); // 4
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
extern void* C_ZNK7QWizard11currentPageEv(void* qthis); // 4
  // proto:  bool QWizard::hasVisitedPage(int id);
extern bool C_ZNK7QWizard14hasVisitedPageEi(void* qthis, int32_t arg0); // 4
  // proto:  QWizard::WizardStyle QWizard::wizardStyle();
extern void C_ZNK7QWizard11wizardStyleEv(void* qthis); // 4
  // proto:  const QMetaObject * QWizard::metaObject();
extern void C_ZNK7QWizard10metaObjectEv(void* qthis); // 4
  // proto:  void QWizard::~QWizard();
extern void C_ZN7QWizardD2Ev(void* qthis); // 4
  // proto:  int QWizard::startId();
extern int32_t C_ZNK7QWizard7startIdEv(void* qthis); // 4
  // proto:  void QWizard::restart();
extern void C_ZN7QWizard7restartEv(void* qthis); // 4
  // proto:  QWidget * QWizard::sideWidget();
extern void* C_ZNK7QWizard10sideWidgetEv(void* qthis); // 4
  // proto:  QSize QWizard::sizeHint();
extern void* C_ZNK7QWizard8sizeHintEv(void* qthis); // 4
  // proto:  int QWizard::currentId();
extern int32_t C_ZNK7QWizard9currentIdEv(void* qthis); // 4
  // proto:  void QWizard::removePage(int id);
extern void C_ZN7QWizard10removePageEi(void* qthis, int32_t arg0); // 4
  // proto:  WizardOptions QWizard::options();
extern void C_ZNK7QWizard7optionsEv(void* qthis); // 4
  // proto:  void QWizard::setStartId(int id);
extern void C_ZN7QWizard10setStartIdEi(void* qthis, int32_t arg0); // 4
  // proto:  int QWizard::nextId();
extern int32_t C_ZNK7QWizard6nextIdEv(void* qthis); // 4
  // proto:  QWizardPage * QWizard::page(int id);
extern void* C_ZNK7QWizard4pageEi(void* qthis, int32_t arg0); // 4
  // proto:  int QWizard::addPage(QWizardPage * page);
extern int32_t C_ZN7QWizard7addPageEP11QWizardPage(void* qthis, void* arg0); // 4
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
func (this *QWizardPage) Settitle(args ...interface{}) () {
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

  return
}

// isCommitPage()
func (this *QWizardPage) Iscommitpage(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWizardPage12isCommitPageEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizardPage", "isCommitPage", args)
  }

  return
}

// cleanupPage()
func (this *QWizardPage) Cleanuppage(args ...interface{}) () {
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

  return
}

// ~QWizardPage()
func (this *QWizardPage) Freeqwizardpage(args ...interface{}) () {
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

  return
}

// subTitle()
func (this *QWizardPage) Subtitle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWizardPage8subTitleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizardPage", "subTitle", args)
  }

  return
}

// title()
func (this *QWizardPage) Title(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWizardPage5titleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizardPage", "title", args)
  }

  return
}

// isComplete()
func (this *QWizardPage) Iscomplete(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWizardPage10isCompleteEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizardPage", "isComplete", args)
  }

  return
}

// isFinalPage()
func (this *QWizardPage) Isfinalpage(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWizardPage11isFinalPageEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizardPage", "isFinalPage", args)
  }

  return
}

// validatePage()
func (this *QWizardPage) Validatepage(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QWizardPage12validatePageEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizardPage", "validatePage", args)
  }

  return
}

// setSubTitle(const class QString &)
func (this *QWizardPage) Setsubtitle(args ...interface{}) () {
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

  return
}

// initializePage()
func (this *QWizardPage) Initializepage(args ...interface{}) () {
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

  return
}

// metaObject()
func (this *QWizardPage) Metaobject(args ...interface{}) () {
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

  return
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
func (this *QWizardPage) Setcommitpage(args ...interface{}) () {
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

  return
}

// nextId()
func (this *QWizardPage) Nextid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QWizardPage6nextIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizardPage", "nextId", args)
  }

  return
}

// setFinalPage(_Bool)
func (this *QWizardPage) Setfinalpage(args ...interface{}) () {
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

  return
}

// validateCurrentPage()
func (this *QWizard) Validatecurrentpage(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN7QWizard19validateCurrentPageEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "validateCurrentPage", args)
  }

  return
}

// setSideWidget(class QWidget *)
func (this *QWizard) Setsidewidget(args ...interface{}) () {
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

  return
}

// back()
func (this *QWizard) Back(args ...interface{}) () {
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

  return
}

// setDefaultProperty(const char *, const char *, const char *)
func (this *QWizard) Setdefaultproperty(args ...interface{}) () {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[0][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    C.C_ZN7QWizard18setDefaultPropertyEPKcS1_S1_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QWizard", "setDefaultProperty", args)
  }

  return
}

// setVisible(_Bool)
func (this *QWizard) Setvisible(args ...interface{}) () {
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

  return
}

// setField(const class QString &, const class QVariant &)
func (this *QWizard) Setfield(args ...interface{}) () {
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

  return
}

// visitedPages()
func (this *QWizard) Visitedpages(args ...interface{}) () {
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

  return
}

// field(const class QString &)
func (this *QWizard) Field(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWizard5fieldERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "field", args)
  }

  return
}

// titleFormat()
func (this *QWizard) Titleformat(args ...interface{}) () {
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

  return
}

// next()
func (this *QWizard) Next(args ...interface{}) () {
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

  return
}

// pageIds()
func (this *QWizard) Pageids(args ...interface{}) () {
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

  return
}

// subTitleFormat()
func (this *QWizard) Subtitleformat(args ...interface{}) () {
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

  return
}

// setPage(int, class QWizardPage *)
func (this *QWizard) Setpage(args ...interface{}) () {
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

  return
}

// currentPage()
func (this *QWizard) Currentpage(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWizard11currentPageEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWizardPage{}) // "QWizardPage *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "currentPage", args)
  }

  return
}

// hasVisitedPage(int)
func (this *QWizard) Hasvisitedpage(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWizard14hasVisitedPageEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "hasVisitedPage", args)
  }

  return
}

// wizardStyle()
func (this *QWizard) Wizardstyle(args ...interface{}) () {
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

  return
}

// metaObject()
func (this *QWizard) Metaobject(args ...interface{}) () {
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

  return
}

// ~QWizard()
func (this *QWizard) Freeqwizard(args ...interface{}) () {
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

  return
}

// startId()
func (this *QWizard) Startid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWizard7startIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "startId", args)
  }

  return
}

// restart()
func (this *QWizard) Restart(args ...interface{}) () {
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

  return
}

// sideWidget()
func (this *QWizard) Sidewidget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWizard10sideWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "sideWidget", args)
  }

  return
}

// sizeHint()
func (this *QWizard) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWizard8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "sizeHint", args)
  }

  return
}

// currentId()
func (this *QWizard) Currentid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWizard9currentIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "currentId", args)
  }

  return
}

// removePage(int)
func (this *QWizard) Removepage(args ...interface{}) () {
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

  return
}

// options()
func (this *QWizard) Options(args ...interface{}) () {
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

  return
}

// setStartId(int)
func (this *QWizard) Setstartid(args ...interface{}) () {
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

  return
}

// nextId()
func (this *QWizard) Nextid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWizard6nextIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "nextId", args)
  }

  return
}

// page(int)
func (this *QWizard) Page(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QWizard4pageEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWizardPage{}) // "QWizardPage *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "page", args)
  }

  return
}

// addPage(class QWizardPage *)
func (this *QWizard) Addpage(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN7QWizard7addPageEP11QWizardPage(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWizard", "addPage", args)
  }

  return
}

// <= body block end

