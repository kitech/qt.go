package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QItemEditorCreatorBase::~QItemEditorCreatorBase();
extern void C_ZN22QItemEditorCreatorBaseD2Ev(void* qthis); // 4
  // proto: static void QItemEditorFactory::setDefaultFactory(QItemEditorFactory * factory);
extern void C_ZN18QItemEditorFactory17setDefaultFactoryEPS_(void* arg0); // 4
  // proto:  void QItemEditorFactory::~QItemEditorFactory();
extern void C_ZN18QItemEditorFactoryD2Ev(void* qthis); // 4
  // proto:  QWidget * QItemEditorFactory::createEditor(int userType, QWidget * parent);
extern void* C_ZNK18QItemEditorFactory12createEditorEiP7QWidget(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QItemEditorFactory::QItemEditorFactory();
extern void* C_ZN18QItemEditorFactoryC2Ev(); // 1
  // proto: static const QItemEditorFactory * QItemEditorFactory::defaultFactory();
extern void* C_ZN18QItemEditorFactory14defaultFactoryEv(); // 4
  // proto:  void QItemEditorFactory::registerEditor(int userType, QItemEditorCreatorBase * creator);
extern void C_ZN18QItemEditorFactory14registerEditorEiP22QItemEditorCreatorBase(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QByteArray QItemEditorFactory::valuePropertyName(int userType);
extern void* C_ZNK18QItemEditorFactory17valuePropertyNameEi(void* qthis, int32_t arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QItemEditorCreatorBase)=8
type QItemEditorCreatorBase struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QItemEditorFactory)=1
type QItemEditorFactory struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QItemEditorCreatorBase()
func (this *QItemEditorCreatorBase) Free(args ...interface{}) () {
  // ~QItemEditorCreatorBase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QItemEditorCreatorBaseD0Ev
    // invoke: void ~QItemEditorCreatorBase()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN22QItemEditorCreatorBaseD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QItemEditorCreatorBase", "~QItemEditorCreatorBase", args)
  }

  return
}

// setDefaultFactory(class QItemEditorFactory *)
func (this *QItemEditorFactory) SetDefaultFactory_s(args ...interface{}) () {
  // setDefaultFactory(class QItemEditorFactory *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemEditorFactory{}) // "QItemEditorFactory *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QItemEditorFactory17setDefaultFactoryEPS_
    // invoke: void setDefaultFactory(class QItemEditorFactory *)
    var arg0 = args[0].(*QItemEditorFactory).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QItemEditorFactory17setDefaultFactoryEPS_(arg0)
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "setDefaultFactory", args)
  }

  return
}

// ~QItemEditorFactory()
func (this *QItemEditorFactory) Free(args ...interface{}) () {
  // ~QItemEditorFactory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QItemEditorFactoryD0Ev
    // invoke: void ~QItemEditorFactory()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN18QItemEditorFactoryD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "~QItemEditorFactory", args)
  }

  return
}

// createEditor(int, class QWidget *)
func (this *QItemEditorFactory) CreateEditor(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK18QItemEditorFactory12createEditorEiP7QWidget(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "createEditor", args)
  }

  return
}

// QItemEditorFactory()
func GcfreeQItemEditorFactory(this *QItemEditorFactory) {
  qtrt.UniverseFree(this)
}
func NewQItemEditorFactory(args ...interface{}) *QItemEditorFactory {
  // QItemEditorFactory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QItemEditorFactoryC1Ev
    // invoke: void QItemEditorFactory()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QItemEditorFactoryC2Ev()
    this := &QItemEditorFactory{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQItemEditorFactory)
    return this
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "QItemEditorFactory", args)
  }

  return nil // QItemEditorFactory{Qclsinst:qthis}
}

// defaultFactory()
func (this *QItemEditorFactory) DefaultFactory_s(args ...interface{}) (ret interface{}) {
  // defaultFactory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QItemEditorFactory14defaultFactoryEv
    // invoke: const QItemEditorFactory * defaultFactory()
    var ret0 = C.C_ZN18QItemEditorFactory14defaultFactoryEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QItemEditorFactory{}) // "const QItemEditorFactory *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "defaultFactory", args)
  }

  return
}

// registerEditor(int, class QItemEditorCreatorBase *)
func (this *QItemEditorFactory) RegisterEditor(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QItemEditorCreatorBase).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN18QItemEditorFactory14registerEditorEiP22QItemEditorCreatorBase(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "registerEditor", args)
  }

  return
}

// valuePropertyName(int)
func (this *QItemEditorFactory) ValuePropertyName(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QItemEditorFactory17valuePropertyNameEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "valuePropertyName", args)
  }

  return
}

// <= body block end

