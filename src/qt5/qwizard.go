package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QWizardPage::QWizardPage(const QWizardPage & );
extern void* dector_ZN11QWizardPageC1ERKS_(void* arg0);
extern void _ZN11QWizardPageC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QWizardPage::metaObject();
extern void _ZNK11QWizardPage10metaObjectEv(void* qthis);
  // proto:  QString QWizardPage::title();
extern void _ZNK11QWizardPage5titleEv(void* qthis);
  // proto:  QString QWizardPage::subTitle();
extern void _ZNK11QWizardPage8subTitleEv(void* qthis);
  // proto:  bool QWizardPage::isFinalPage();
extern void _ZNK11QWizardPage11isFinalPageEv(void* qthis);
  // proto:  bool QWizardPage::validatePage();
extern void _ZN11QWizardPage12validatePageEv(void* qthis);
  // proto:  int QWizardPage::nextId();
extern void _ZNK11QWizardPage6nextIdEv(void* qthis);
  // proto:  void QWizardPage::cleanupPage();
extern void _ZN11QWizardPage11cleanupPageEv(void* qthis);
  // proto:  bool QWizardPage::isComplete();
extern void _ZNK11QWizardPage10isCompleteEv(void* qthis);
  // proto:  bool QWizardPage::isCommitPage();
extern void _ZNK11QWizardPage12isCommitPageEv(void* qthis);
  // proto:  void QWizardPage::QWizardPage(QWidget * parent);
extern void* dector_ZN11QWizardPageC1EP7QWidget(void* arg0);
extern void _ZN11QWizardPageC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QWizardPage::setFinalPage(bool finalPage);
extern void _ZN11QWizardPage12setFinalPageEb(void* qthis, bool arg0);
  // proto:  void QWizardPage::setSubTitle(const QString & subTitle);
extern void _ZN11QWizardPage11setSubTitleERK7QString(void* qthis, void* arg0);
  // proto:  void QWizardPage::~QWizardPage();
extern void _ZN11QWizardPageD0Ev(void* qthis);
  // proto:  void QWizardPage::setCommitPage(bool commitPage);
extern void _ZN11QWizardPage13setCommitPageEb(void* qthis, bool arg0);
  // proto:  void QWizardPage::initializePage();
extern void _ZN11QWizardPage14initializePageEv(void* qthis);
  // proto:  void QWizardPage::setTitle(const QString & title);
extern void _ZN11QWizardPage8setTitleERK7QString(void* qthis, void* arg0);
  // proto:  void QWizard::setSideWidget(QWidget * widget);
extern void _ZN7QWizard13setSideWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  QWizardPage * QWizard::currentPage();
extern void _ZNK7QWizard11currentPageEv(void* qthis);
  // proto:  void QWizard::next();
extern void _ZN7QWizard4nextEv(void* qthis);
  // proto:  QWizardPage * QWizard::page(int id);
extern void _ZNK7QWizard4pageEi(void* qthis, int32_t arg0);
  // proto:  const QMetaObject * QWizard::metaObject();
extern void _ZNK7QWizard10metaObjectEv(void* qthis);
  // proto:  void QWizard::setField(const QString & name, const QVariant & value);
extern void _ZN7QWizard8setFieldERK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1);
  // proto:  void QWizard::setPage(int id, QWizardPage * page);
extern void _ZN7QWizard7setPageEiP11QWizardPage(void* qthis, int32_t arg0, void* arg1);
  // proto:  void QWizard::restart();
extern void _ZN7QWizard7restartEv(void* qthis);
  // proto:  void QWizard::back();
extern void _ZN7QWizard4backEv(void* qthis);
  // proto:  QSize QWizard::sizeHint();
extern void _ZNK7QWizard8sizeHintEv(void* qthis);
  // proto:  void QWizard::setDefaultProperty(const char * className, const char * property, const char * changedSignal);
extern void _ZN7QWizard18setDefaultPropertyEPKcS1_S1_(void* qthis, unsigned char* arg0, unsigned char* arg1, unsigned char* arg2);
  // proto:  void QWizard::setStartId(int id);
extern void _ZN7QWizard10setStartIdEi(void* qthis, int32_t arg0);
  // proto:  void QWizard::~QWizard();
extern void _ZN7QWizardD0Ev(void* qthis);
  // proto:  QList<int> QWizard::visitedPages();
extern void _ZNK7QWizard12visitedPagesEv(void* qthis);
  // proto:  int QWizard::nextId();
extern void _ZNK7QWizard6nextIdEv(void* qthis);
  // proto:  int QWizard::startId();
extern void _ZNK7QWizard7startIdEv(void* qthis);
  // proto:  void QWizard::QWizard(const QWizard & );
extern void* dector_ZN7QWizardC1ERKS_(void* arg0);
extern void _ZN7QWizardC1ERKS_(void* qthis, void* arg0);
  // proto:  int QWizard::addPage(QWizardPage * page);
extern void _ZN7QWizard7addPageEP11QWizardPage(void* qthis, void* arg0);
  // proto:  void QWizard::removePage(int id);
extern void _ZN7QWizard10removePageEi(void* qthis, int32_t arg0);
  // proto:  QList<int> QWizard::pageIds();
extern void _ZNK7QWizard7pageIdsEv(void* qthis);
  // proto:  int QWizard::currentId();
extern void _ZNK7QWizard9currentIdEv(void* qthis);
  // proto:  void QWizard::setVisible(bool visible);
extern void _ZN7QWizard10setVisibleEb(void* qthis, bool arg0);
  // proto:  bool QWizard::hasVisitedPage(int id);
extern void _ZNK7QWizard14hasVisitedPageEi(void* qthis, int32_t arg0);
  // proto:  QVariant QWizard::field(const QString & name);
extern void _ZNK7QWizard5fieldERK7QString(void* qthis, void* arg0);
  // proto:  bool QWizard::validateCurrentPage();
extern void _ZN7QWizard19validateCurrentPageEv(void* qthis);
  // proto:  QWidget * QWizard::sideWidget();
extern void _ZNK7QWizard10sideWidgetEv(void* qthis);
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

  // proto:  void QWizardPage::QWizardPage(const QWizardPage & );
func NewQWizardPage(args ...interface{}) QWizardPage {
  return QWizardPage{}
}

  // proto:  const QMetaObject * QWizardPage::metaObject();
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
    C._ZNK11QWizardPage10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "metaObject", args)
  }

}

  // proto:  QString QWizardPage::title();
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
    C._ZNK11QWizardPage5titleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "title", args)
  }

}

  // proto:  QString QWizardPage::subTitle();
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
    C._ZNK11QWizardPage8subTitleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "subTitle", args)
  }

}

  // proto:  bool QWizardPage::isFinalPage();
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
    C._ZNK11QWizardPage11isFinalPageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "isFinalPage", args)
  }

}

  // proto:  bool QWizardPage::validatePage();
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
    C._ZN11QWizardPage12validatePageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "validatePage", args)
  }

}

  // proto:  int QWizardPage::nextId();
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
    C._ZNK11QWizardPage6nextIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "nextId", args)
  }

}

  // proto:  void QWizardPage::cleanupPage();
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
    C._ZN11QWizardPage11cleanupPageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "cleanupPage", args)
  }

}

  // proto:  bool QWizardPage::isComplete();
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
    C._ZNK11QWizardPage10isCompleteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "isComplete", args)
  }

}

  // proto:  bool QWizardPage::isCommitPage();
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
    C._ZNK11QWizardPage12isCommitPageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "isCommitPage", args)
  }

}

  // proto:  void QWizardPage::setFinalPage(bool finalPage);
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
    C._ZN11QWizardPage12setFinalPageEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizardPage", "setFinalPage", args)
  }

}

  // proto:  void QWizardPage::setSubTitle(const QString & subTitle);
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
    C._ZN11QWizardPage11setSubTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizardPage", "setSubTitle", args)
  }

}

  // proto:  void QWizardPage::~QWizardPage();
func (this *QWizardPage) FreeQWizardPage(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWizardPage", "~QWizardPage", args)
  }

}

  // proto:  void QWizardPage::setCommitPage(bool commitPage);
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
    C._ZN11QWizardPage13setCommitPageEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizardPage", "setCommitPage", args)
  }

}

  // proto:  void QWizardPage::initializePage();
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
    C._ZN11QWizardPage14initializePageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizardPage", "initializePage", args)
  }

}

  // proto:  void QWizardPage::setTitle(const QString & title);
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
    C._ZN11QWizardPage8setTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizardPage", "setTitle", args)
  }

}

  // proto:  void QWizard::setSideWidget(QWidget * widget);
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
    C._ZN7QWizard13setSideWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "setSideWidget", args)
  }

}

  // proto:  QWizardPage * QWizard::currentPage();
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
    C._ZNK7QWizard11currentPageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "currentPage", args)
  }

}

  // proto:  void QWizard::next();
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
    C._ZN7QWizard4nextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "next", args)
  }

}

  // proto:  QWizardPage * QWizard::page(int id);
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
    C._ZNK7QWizard4pageEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "page", args)
  }

}

  // proto:  const QMetaObject * QWizard::metaObject();
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
    C._ZNK7QWizard10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "metaObject", args)
  }

}

  // proto:  void QWizard::setField(const QString & name, const QVariant & value);
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
    C._ZN7QWizard8setFieldERK7QStringRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWizard", "setField", args)
  }

}

  // proto:  void QWizard::setPage(int id, QWizardPage * page);
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
    C._ZN7QWizard7setPageEiP11QWizardPage(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWizard", "setPage", args)
  }

}

  // proto:  void QWizard::restart();
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
    C._ZN7QWizard7restartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "restart", args)
  }

}

  // proto:  void QWizard::back();
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
    C._ZN7QWizard4backEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "back", args)
  }

}

  // proto:  QSize QWizard::sizeHint();
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
    C._ZNK7QWizard8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "sizeHint", args)
  }

}

  // proto:  void QWizard::setDefaultProperty(const char * className, const char * property, const char * changedSignal);
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg2)}
    C._ZN7QWizard18setDefaultPropertyEPKcS1_S1_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QWizard", "setDefaultProperty", args)
  }

}

  // proto:  void QWizard::setStartId(int id);
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
    C._ZN7QWizard10setStartIdEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "setStartId", args)
  }

}

  // proto:  void QWizard::~QWizard();
func (this *QWizard) FreeQWizard(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWizard", "~QWizard", args)
  }

}

  // proto:  QList<int> QWizard::visitedPages();
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
    C._ZNK7QWizard12visitedPagesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "visitedPages", args)
  }

}

  // proto:  int QWizard::nextId();
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
    C._ZNK7QWizard6nextIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "nextId", args)
  }

}

  // proto:  int QWizard::startId();
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
    C._ZNK7QWizard7startIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "startId", args)
  }

}

  // proto:  void QWizard::QWizard(const QWizard & );
func NewQWizard(args ...interface{}) QWizard {
  return QWizard{}
}

  // proto:  int QWizard::addPage(QWizardPage * page);
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
    C._ZN7QWizard7addPageEP11QWizardPage(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "addPage", args)
  }

}

  // proto:  void QWizard::removePage(int id);
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
    C._ZN7QWizard10removePageEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "removePage", args)
  }

}

  // proto:  QList<int> QWizard::pageIds();
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
    C._ZNK7QWizard7pageIdsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "pageIds", args)
  }

}

  // proto:  int QWizard::currentId();
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
    C._ZNK7QWizard9currentIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "currentId", args)
  }

}

  // proto:  void QWizard::setVisible(bool visible);
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
    C._ZN7QWizard10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "setVisible", args)
  }

}

  // proto:  bool QWizard::hasVisitedPage(int id);
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
    C._ZNK7QWizard14hasVisitedPageEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "hasVisitedPage", args)
  }

}

  // proto:  QVariant QWizard::field(const QString & name);
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
    C._ZNK7QWizard5fieldERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWizard", "field", args)
  }

}

  // proto:  bool QWizard::validateCurrentPage();
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
    C._ZN7QWizard19validateCurrentPageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "validateCurrentPage", args)
  }

}

  // proto:  QWidget * QWizard::sideWidget();
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
    C._ZNK7QWizard10sideWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWizard", "sideWidget", args)
  }

}

// <= body block end

