package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qitemeditorfactory.h
// dst-file: /src/widgets/qitemeditorfactory.go
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
  // proto:  QByteArray QItemEditorCreatorBase::valuePropertyName();
extern void _ZNK22QItemEditorCreatorBase17valuePropertyNameEv(void* qthis);
  // proto:  QWidget * QItemEditorCreatorBase::createWidget(QWidget * parent);
extern void _ZNK22QItemEditorCreatorBase12createWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QItemEditorCreatorBase::~QItemEditorCreatorBase();
extern void _ZN22QItemEditorCreatorBaseD0Ev(void* qthis);
  // proto:  void QItemEditorFactory::QItemEditorFactory();
extern void* dector_ZN18QItemEditorFactoryC1Ev();
extern void demth_ZN18QItemEditorFactoryC1Ev(void* qthis);
  // proto:  QByteArray QItemEditorFactory::valuePropertyName(int userType);
extern void _ZNK18QItemEditorFactory17valuePropertyNameEi(void* qthis, int arg0);
  // proto: static const QItemEditorFactory * QItemEditorFactory::defaultFactory();
extern void _ZN18QItemEditorFactory14defaultFactoryEv();
  // proto:  void QItemEditorFactory::~QItemEditorFactory();
extern void _ZN18QItemEditorFactoryD0Ev(void* qthis);
  // proto:  void QItemEditorFactory::registerEditor(int userType, QItemEditorCreatorBase * creator);
extern void _ZN18QItemEditorFactory14registerEditorEiP22QItemEditorCreatorBase(void* qthis, int arg0, void* arg1);
  // proto: static void QItemEditorFactory::setDefaultFactory(QItemEditorFactory * factory);
extern void _ZN18QItemEditorFactory17setDefaultFactoryEPS_(void* arg0);
  // proto:  QWidget * QItemEditorFactory::createEditor(int userType, QWidget * parent);
extern void _ZNK18QItemEditorFactory12createEditorEiP7QWidget(void* qthis, int arg0, void* arg1);
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

// class sizeof(QItemEditorCreatorBase)=8
type QItemEditorCreatorBase struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QItemEditorFactory)=1
type QItemEditorFactory struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QByteArray QItemEditorCreatorBase::valuePropertyName();
func (this *QItemEditorCreatorBase) valuePropertyName(args ...interface{}) () {
  // valuePropertyName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QItemEditorCreatorBase17valuePropertyNameEv
    // invoke: QByteArray valuePropertyName()
    C._ZNK22QItemEditorCreatorBase17valuePropertyNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QItemEditorCreatorBase", "valuePropertyName", args)
  }

}

  // proto:  QWidget * QItemEditorCreatorBase::createWidget(QWidget * parent);
func (this *QItemEditorCreatorBase) createWidget(args ...interface{}) () {
  // createWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QItemEditorCreatorBase12createWidgetEP7QWidget
    // invoke: QWidget * createWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK22QItemEditorCreatorBase12createWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemEditorCreatorBase", "createWidget", args)
  }

}

  // proto:  void QItemEditorCreatorBase::~QItemEditorCreatorBase();
func (this *QItemEditorCreatorBase) FreeQItemEditorCreatorBase(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemEditorCreatorBase", "~QItemEditorCreatorBase", args)
  }

}

  // proto:  void QItemEditorFactory::QItemEditorFactory();
func NewQItemEditorFactory(args ...interface{}) QItemEditorFactory {
  return QItemEditorFactory{}
}

  // proto:  QByteArray QItemEditorFactory::valuePropertyName(int userType);
func (this *QItemEditorFactory) valuePropertyName(args ...interface{}) () {
  // valuePropertyName(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QItemEditorFactory17valuePropertyNameEi
    // invoke: QByteArray valuePropertyName(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK18QItemEditorFactory17valuePropertyNameEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "valuePropertyName", args)
  }

}

  // proto: static const QItemEditorFactory * QItemEditorFactory::defaultFactory();
func (this *QItemEditorFactory) defaultFactory_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "defaultFactory", args)
  }

}

  // proto:  void QItemEditorFactory::~QItemEditorFactory();
func (this *QItemEditorFactory) FreeQItemEditorFactory(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "~QItemEditorFactory", args)
  }

}

  // proto:  void QItemEditorFactory::registerEditor(int userType, QItemEditorCreatorBase * creator);
func (this *QItemEditorFactory) registerEditor(args ...interface{}) () {
  // registerEditor(int, class QItemEditorCreatorBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QItemEditorCreatorBase{}) // "QItemEditorCreatorBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QItemEditorFactory14registerEditorEiP22QItemEditorCreatorBase
    // invoke: void registerEditor(int, class QItemEditorCreatorBase *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QItemEditorCreatorBase).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN18QItemEditorFactory14registerEditorEiP22QItemEditorCreatorBase(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "registerEditor", args)
  }

}

  // proto: static void QItemEditorFactory::setDefaultFactory(QItemEditorFactory * factory);
func (this *QItemEditorFactory) setDefaultFactory_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "setDefaultFactory", args)
  }

}

  // proto:  QWidget * QItemEditorFactory::createEditor(int userType, QWidget * parent);
func (this *QItemEditorFactory) createEditor(args ...interface{}) () {
  // createEditor(int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QItemEditorFactory12createEditorEiP7QWidget
    // invoke: QWidget * createEditor(int, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK18QItemEditorFactory12createEditorEiP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "createEditor", args)
  }

}

// <= body block end

