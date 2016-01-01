package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qwizard.h
// dst-file: /src/widgets/qwizard.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QWizardPage)=1
type QWizardPage struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _completeChanged QWizardPage_completeChanged_signal;
}

// class sizeof(QWizard)=1
type QWizard struct {
  /*qbase*/ QDialog;
  qclsinst uint64 /* *mut c_void*/;
//  _pageRemoved QWizard_pageRemoved_signal;
//  _pageAdded QWizard_pageAdded_signal;
//  _helpRequested QWizard_helpRequested_signal;
//  _currentIdChanged QWizard_currentIdChanged_signal;
//  _customButtonClicked QWizard_customButtonClicked_signal;
}


func NewQWizardPage(args ...interface{}) QWizardPage {
  return QWizardPage{}
}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "metaObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "title", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "subTitle", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "isFinalPage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "validatePage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "nextId", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "cleanupPage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "isComplete", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "isCommitPage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "setFinalPage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "setSubTitle", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "setCommitPage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "initializePage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizardPage", "setTitle", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "setSideWidget", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "currentPage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "next", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "page", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "metaObject", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "setField", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "setPage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "restart", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "back", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "sizeHint", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "setDefaultProperty", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "setStartId", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "visitedPages", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "nextId", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "startId", args)
  }

}


func NewQWizard(args ...interface{}) QWizard {
  return QWizard{}
}


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
  default:
    qtrt.ErrorResolve("QWizard", "addPage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "removePage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "pageIds", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "currentId", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "setVisible", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "hasVisitedPage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "field", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "validateCurrentPage", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QWizard", "sideWidget", args)
  }

}

// <= body block end

