package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qdatawidgetmapper.h
// dst-file: /src/widgets/qdatawidgetmapper.go
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
  // proto:  Qt::Orientation QDataWidgetMapper::orientation();
extern void C_ZNK17QDataWidgetMapper11orientationEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::setItemDelegate(QAbstractItemDelegate * delegate);
extern void C_ZN17QDataWidgetMapper15setItemDelegateEP21QAbstractItemDelegate(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::revert();
extern void C_ZN17QDataWidgetMapper6revertEv(void* qthis); // 4
  // proto:  int QDataWidgetMapper::mappedSection(QWidget * widget);
extern int32_t C_ZNK17QDataWidgetMapper13mappedSectionEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::setCurrentIndex(int index);
extern void C_ZN17QDataWidgetMapper15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  void QDataWidgetMapper::addMapping(QWidget * widget, int section, const QByteArray & propertyName);
extern void C_ZN17QDataWidgetMapper10addMappingEP7QWidgetiRK10QByteArray(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QDataWidgetMapper::addMapping(QWidget * widget, int section);
extern void C_ZN17QDataWidgetMapper10addMappingEP7QWidgeti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QDataWidgetMapper::toPrevious();
extern void C_ZN17QDataWidgetMapper10toPreviousEv(void* qthis); // 4
  // proto:  bool QDataWidgetMapper::submit();
extern bool C_ZN17QDataWidgetMapper6submitEv(void* qthis); // 4
  // proto:  QModelIndex QDataWidgetMapper::rootIndex();
extern void* C_ZNK17QDataWidgetMapper9rootIndexEv(void* qthis); // 4
  // proto:  QByteArray QDataWidgetMapper::mappedPropertyName(QWidget * widget);
extern void* C_ZNK17QDataWidgetMapper18mappedPropertyNameEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::clearMapping();
extern void C_ZN17QDataWidgetMapper12clearMappingEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::toLast();
extern void C_ZN17QDataWidgetMapper6toLastEv(void* qthis); // 4
  // proto:  QDataWidgetMapper::SubmitPolicy QDataWidgetMapper::submitPolicy();
extern void C_ZNK17QDataWidgetMapper12submitPolicyEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::QDataWidgetMapper(QObject * parent);
extern void* C_ZN17QDataWidgetMapperC2EP7QObject(void* arg0); // 3
  // proto:  void QDataWidgetMapper::setRootIndex(const QModelIndex & index);
extern void C_ZN17QDataWidgetMapper12setRootIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::setModel(QAbstractItemModel * model);
extern void C_ZN17QDataWidgetMapper8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::toFirst();
extern void C_ZN17QDataWidgetMapper7toFirstEv(void* qthis); // 4
  // proto:  QAbstractItemDelegate * QDataWidgetMapper::itemDelegate();
extern void C_ZNK17QDataWidgetMapper12itemDelegateEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::toNext();
extern void C_ZN17QDataWidgetMapper6toNextEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::~QDataWidgetMapper();
extern void C_ZN17QDataWidgetMapperD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QDataWidgetMapper::metaObject();
extern void C_ZNK17QDataWidgetMapper10metaObjectEv(void* qthis); // 4
  // proto:  QWidget * QDataWidgetMapper::mappedWidgetAt(int section);
extern void* C_ZNK17QDataWidgetMapper14mappedWidgetAtEi(void* qthis, int32_t arg0); // 4
  // proto:  int QDataWidgetMapper::currentIndex();
extern int32_t C_ZNK17QDataWidgetMapper12currentIndexEv(void* qthis); // 4
  // proto:  QAbstractItemModel * QDataWidgetMapper::model();
extern void C_ZNK17QDataWidgetMapper5modelEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::setCurrentModelIndex(const QModelIndex & index);
extern void C_ZN17QDataWidgetMapper20setCurrentModelIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::removeMapping(QWidget * widget);
extern void C_ZN17QDataWidgetMapper13removeMappingEP7QWidget(void* qthis, void* arg0); // 4
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

// class sizeof(QDataWidgetMapper)=1
type QDataWidgetMapper struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _currentIndexChanged QDataWidgetMapper_currentIndexChanged_signal;
}

// orientation()
func (this *QDataWidgetMapper) Orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK17QDataWidgetMapper11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "orientation", args)
  }

  return
}

// setItemDelegate(class QAbstractItemDelegate *)
func (this *QDataWidgetMapper) Setitemdelegate(args ...interface{}) () {
  // setItemDelegate(class QAbstractItemDelegate *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemDelegate{}) // "QAbstractItemDelegate *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper15setItemDelegateEP21QAbstractItemDelegate
    // invoke: void setItemDelegate(class QAbstractItemDelegate *)
    var arg0 = args[0].(QAbstractItemDelegate).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QDataWidgetMapper15setItemDelegateEP21QAbstractItemDelegate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setItemDelegate", args)
  }

  return
}

// revert()
func (this *QDataWidgetMapper) Revert(args ...interface{}) () {
  // revert()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper6revertEv
    // invoke: void revert()
    C.C_ZN17QDataWidgetMapper6revertEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "revert", args)
  }

  return
}

// mappedSection(class QWidget *)
func (this *QDataWidgetMapper) Mappedsection(args ...interface{}) (ret interface{}) {
  // mappedSection(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper13mappedSectionEP7QWidget
    // invoke: int mappedSection(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QDataWidgetMapper13mappedSectionEP7QWidget(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedSection", args)
  }

  return
}

// setCurrentIndex(int)
func (this *QDataWidgetMapper) Setcurrentindex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN17QDataWidgetMapper15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setCurrentIndex", args)
  }

  return
}

// addMapping(class QWidget *, int, const class QByteArray &)
func (this *QDataWidgetMapper) Addmapping(args ...interface{}) () {
  // addMapping(class QWidget *, int, const class QByteArray &)
  // addMapping(class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper10addMappingEP7QWidgetiRK10QByteArray
    // invoke: void addMapping(class QWidget *, int, const class QByteArray &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QByteArray).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN17QDataWidgetMapper10addMappingEP7QWidgetiRK10QByteArray(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN17QDataWidgetMapper10addMappingEP7QWidgeti
    // invoke: void addMapping(class QWidget *, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN17QDataWidgetMapper10addMappingEP7QWidgeti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "addMapping", args)
  }

  return
}

// toPrevious()
func (this *QDataWidgetMapper) Toprevious(args ...interface{}) () {
  // toPrevious()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper10toPreviousEv
    // invoke: void toPrevious()
    C.C_ZN17QDataWidgetMapper10toPreviousEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toPrevious", args)
  }

  return
}

// submit()
func (this *QDataWidgetMapper) Submit(args ...interface{}) (ret interface{}) {
  // submit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper6submitEv
    // invoke: bool submit()
    var ret0 = C.C_ZN17QDataWidgetMapper6submitEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "submit", args)
  }

  return
}

// rootIndex()
func (this *QDataWidgetMapper) Rootindex(args ...interface{}) (ret interface{}) {
  // rootIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper9rootIndexEv
    // invoke: QModelIndex rootIndex()
    var ret0 = C.C_ZNK17QDataWidgetMapper9rootIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "rootIndex", args)
  }

  return
}

// mappedPropertyName(class QWidget *)
func (this *QDataWidgetMapper) Mappedpropertyname(args ...interface{}) (ret interface{}) {
  // mappedPropertyName(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper18mappedPropertyNameEP7QWidget
    // invoke: QByteArray mappedPropertyName(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QDataWidgetMapper18mappedPropertyNameEP7QWidget(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedPropertyName", args)
  }

  return
}

// clearMapping()
func (this *QDataWidgetMapper) Clearmapping(args ...interface{}) () {
  // clearMapping()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper12clearMappingEv
    // invoke: void clearMapping()
    C.C_ZN17QDataWidgetMapper12clearMappingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "clearMapping", args)
  }

  return
}

// toLast()
func (this *QDataWidgetMapper) Tolast(args ...interface{}) () {
  // toLast()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper6toLastEv
    // invoke: void toLast()
    C.C_ZN17QDataWidgetMapper6toLastEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toLast", args)
  }

  return
}

// submitPolicy()
func (this *QDataWidgetMapper) Submitpolicy(args ...interface{}) () {
  // submitPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper12submitPolicyEv
    // invoke: QDataWidgetMapper::SubmitPolicy submitPolicy()
    C.C_ZNK17QDataWidgetMapper12submitPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "submitPolicy", args)
  }

  return
}

// QDataWidgetMapper(class QObject *)
func NewQDataWidgetMapper(args ...interface{}) *QDataWidgetMapper {
  // QDataWidgetMapper(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapperC1EP7QObject
    // invoke: void QDataWidgetMapper(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QDataWidgetMapperC2EP7QObject(arg0)
    return &QDataWidgetMapper{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "QDataWidgetMapper", args)
  }

  return nil // QDataWidgetMapper{qclsinst:qthis}
}

// setRootIndex(const class QModelIndex &)
func (this *QDataWidgetMapper) Setrootindex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper12setRootIndexERK11QModelIndex
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QDataWidgetMapper12setRootIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setRootIndex", args)
  }

  return
}

// setModel(class QAbstractItemModel *)
func (this *QDataWidgetMapper) Setmodel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QDataWidgetMapper8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setModel", args)
  }

  return
}

// toFirst()
func (this *QDataWidgetMapper) Tofirst(args ...interface{}) () {
  // toFirst()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper7toFirstEv
    // invoke: void toFirst()
    C.C_ZN17QDataWidgetMapper7toFirstEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toFirst", args)
  }

  return
}

// itemDelegate()
func (this *QDataWidgetMapper) Itemdelegate(args ...interface{}) () {
  // itemDelegate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper12itemDelegateEv
    // invoke: QAbstractItemDelegate * itemDelegate()
    C.C_ZNK17QDataWidgetMapper12itemDelegateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "itemDelegate", args)
  }

  return
}

// toNext()
func (this *QDataWidgetMapper) Tonext(args ...interface{}) () {
  // toNext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper6toNextEv
    // invoke: void toNext()
    C.C_ZN17QDataWidgetMapper6toNextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toNext", args)
  }

  return
}

// ~QDataWidgetMapper()
func (this *QDataWidgetMapper) Freeqdatawidgetmapper(args ...interface{}) () {
  // ~QDataWidgetMapper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapperD0Ev
    // invoke: void ~QDataWidgetMapper()
    C.C_ZN17QDataWidgetMapperD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "~QDataWidgetMapper", args)
  }

  return
}

// metaObject()
func (this *QDataWidgetMapper) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK17QDataWidgetMapper10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "metaObject", args)
  }

  return
}

// mappedWidgetAt(int)
func (this *QDataWidgetMapper) Mappedwidgetat(args ...interface{}) (ret interface{}) {
  // mappedWidgetAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper14mappedWidgetAtEi
    // invoke: QWidget * mappedWidgetAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QDataWidgetMapper14mappedWidgetAtEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedWidgetAt", args)
  }

  return
}

// currentIndex()
func (this *QDataWidgetMapper) Currentindex(args ...interface{}) (ret interface{}) {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper12currentIndexEv
    // invoke: int currentIndex()
    var ret0 = C.C_ZNK17QDataWidgetMapper12currentIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "currentIndex", args)
  }

  return
}

// model()
func (this *QDataWidgetMapper) Model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QDataWidgetMapper5modelEv
    // invoke: QAbstractItemModel * model()
    C.C_ZNK17QDataWidgetMapper5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "model", args)
  }

  return
}

// setCurrentModelIndex(const class QModelIndex &)
func (this *QDataWidgetMapper) Setcurrentmodelindex(args ...interface{}) () {
  // setCurrentModelIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper20setCurrentModelIndexERK11QModelIndex
    // invoke: void setCurrentModelIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QDataWidgetMapper20setCurrentModelIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setCurrentModelIndex", args)
  }

  return
}

// removeMapping(class QWidget *)
func (this *QDataWidgetMapper) Removemapping(args ...interface{}) () {
  // removeMapping(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QDataWidgetMapper13removeMappingEP7QWidget
    // invoke: void removeMapping(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QDataWidgetMapper13removeMappingEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "removeMapping", args)
  }

  return
}

// <= body block end

