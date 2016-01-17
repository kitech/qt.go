package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZNK17QDataWidgetMapper11orientationEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::setItemDelegate(QAbstractItemDelegate * delegate);
extern void _ZN17QDataWidgetMapper15setItemDelegateEP21QAbstractItemDelegate(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::revert();
extern void _ZN17QDataWidgetMapper6revertEv(void* qthis); // 4
  // proto:  int QDataWidgetMapper::mappedSection(QWidget * widget);
extern void _ZNK17QDataWidgetMapper13mappedSectionEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::setCurrentIndex(int index);
extern void _ZN17QDataWidgetMapper15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  void QDataWidgetMapper::addMapping(QWidget * widget, int section, const QByteArray & propertyName);
extern void _ZN17QDataWidgetMapper10addMappingEP7QWidgetiRK10QByteArray(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QDataWidgetMapper::addMapping(QWidget * widget, int section);
extern void _ZN17QDataWidgetMapper10addMappingEP7QWidgeti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QDataWidgetMapper::toPrevious();
extern void _ZN17QDataWidgetMapper10toPreviousEv(void* qthis); // 4
  // proto:  bool QDataWidgetMapper::submit();
extern void _ZN17QDataWidgetMapper6submitEv(void* qthis); // 4
  // proto:  QModelIndex QDataWidgetMapper::rootIndex();
extern void _ZNK17QDataWidgetMapper9rootIndexEv(void* qthis); // 4
  // proto:  QByteArray QDataWidgetMapper::mappedPropertyName(QWidget * widget);
extern void _ZNK17QDataWidgetMapper18mappedPropertyNameEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::clearMapping();
extern void _ZN17QDataWidgetMapper12clearMappingEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::toLast();
extern void _ZN17QDataWidgetMapper6toLastEv(void* qthis); // 4
  // proto:  QDataWidgetMapper::SubmitPolicy QDataWidgetMapper::submitPolicy();
extern void _ZNK17QDataWidgetMapper12submitPolicyEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::QDataWidgetMapper(QObject * parent);
extern void _ZN17QDataWidgetMapperC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QDataWidgetMapper::setRootIndex(const QModelIndex & index);
extern void _ZN17QDataWidgetMapper12setRootIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::setModel(QAbstractItemModel * model);
extern void _ZN17QDataWidgetMapper8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::toFirst();
extern void _ZN17QDataWidgetMapper7toFirstEv(void* qthis); // 4
  // proto:  QAbstractItemDelegate * QDataWidgetMapper::itemDelegate();
extern void _ZNK17QDataWidgetMapper12itemDelegateEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::toNext();
extern void _ZN17QDataWidgetMapper6toNextEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::~QDataWidgetMapper();
extern void _ZN17QDataWidgetMapperD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QDataWidgetMapper::metaObject();
extern void _ZNK17QDataWidgetMapper10metaObjectEv(void* qthis); // 4
  // proto:  QWidget * QDataWidgetMapper::mappedWidgetAt(int section);
extern void _ZNK17QDataWidgetMapper14mappedWidgetAtEi(void* qthis, int32_t arg0); // 4
  // proto:  int QDataWidgetMapper::currentIndex();
extern void _ZNK17QDataWidgetMapper12currentIndexEv(void* qthis); // 4
  // proto:  QAbstractItemModel * QDataWidgetMapper::model();
extern void _ZNK17QDataWidgetMapper5modelEv(void* qthis); // 4
  // proto:  void QDataWidgetMapper::setCurrentModelIndex(const QModelIndex & index);
extern void _ZN17QDataWidgetMapper20setCurrentModelIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QDataWidgetMapper::removeMapping(QWidget * widget);
extern void _ZN17QDataWidgetMapper13removeMappingEP7QWidget(void* qthis, void* arg0); // 4
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
func (this *QDataWidgetMapper) orientation(args ...interface{}) () {
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
    C._ZNK17QDataWidgetMapper11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "orientation", args)
  }

}

// setItemDelegate(class QAbstractItemDelegate *)
func (this *QDataWidgetMapper) setItemDelegate(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper15setItemDelegateEP21QAbstractItemDelegate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setItemDelegate", args)
  }

}

// revert()
func (this *QDataWidgetMapper) revert(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper6revertEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "revert", args)
  }

}

// mappedSection(class QWidget *)
func (this *QDataWidgetMapper) mappedSection(args ...interface{}) () {
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
    C._ZNK17QDataWidgetMapper13mappedSectionEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedSection", args)
  }

}

// setCurrentIndex(int)
func (this *QDataWidgetMapper) setCurrentIndex(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setCurrentIndex", args)
  }

}

// addMapping(class QWidget *, int, const class QByteArray &)
func (this *QDataWidgetMapper) addMapping(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper10addMappingEP7QWidgetiRK10QByteArray(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN17QDataWidgetMapper10addMappingEP7QWidgeti
    // invoke: void addMapping(class QWidget *, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN17QDataWidgetMapper10addMappingEP7QWidgeti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "addMapping", args)
  }

}

// toPrevious()
func (this *QDataWidgetMapper) toPrevious(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper10toPreviousEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toPrevious", args)
  }

}

// submit()
func (this *QDataWidgetMapper) submit(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper6submitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "submit", args)
  }

}

// rootIndex()
func (this *QDataWidgetMapper) rootIndex(args ...interface{}) () {
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
    C._ZNK17QDataWidgetMapper9rootIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "rootIndex", args)
  }

}

// mappedPropertyName(class QWidget *)
func (this *QDataWidgetMapper) mappedPropertyName(args ...interface{}) () {
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
    C._ZNK17QDataWidgetMapper18mappedPropertyNameEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedPropertyName", args)
  }

}

// clearMapping()
func (this *QDataWidgetMapper) clearMapping(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper12clearMappingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "clearMapping", args)
  }

}

// toLast()
func (this *QDataWidgetMapper) toLast(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper6toLastEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toLast", args)
  }

}

// submitPolicy()
func (this *QDataWidgetMapper) submitPolicy(args ...interface{}) () {
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
    C._ZNK17QDataWidgetMapper12submitPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "submitPolicy", args)
  }

}

// QDataWidgetMapper(class QObject *)
func NewQDataWidgetMapper(args ...interface{}) QDataWidgetMapper {
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
    C._ZN17QDataWidgetMapperC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "QDataWidgetMapper", args)
  }

  return QDataWidgetMapper{}
}

// setRootIndex(const class QModelIndex &)
func (this *QDataWidgetMapper) setRootIndex(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper12setRootIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setRootIndex", args)
  }

}

// setModel(class QAbstractItemModel *)
func (this *QDataWidgetMapper) setModel(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setModel", args)
  }

}

// toFirst()
func (this *QDataWidgetMapper) toFirst(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper7toFirstEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toFirst", args)
  }

}

// itemDelegate()
func (this *QDataWidgetMapper) itemDelegate(args ...interface{}) () {
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
    C._ZNK17QDataWidgetMapper12itemDelegateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "itemDelegate", args)
  }

}

// toNext()
func (this *QDataWidgetMapper) toNext(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper6toNextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toNext", args)
  }

}

// ~QDataWidgetMapper()
func (this *QDataWidgetMapper) FreeQDataWidgetMapper(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapperD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "~QDataWidgetMapper", args)
  }

}

// metaObject()
func (this *QDataWidgetMapper) metaObject(args ...interface{}) () {
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
    C._ZNK17QDataWidgetMapper10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "metaObject", args)
  }

}

// mappedWidgetAt(int)
func (this *QDataWidgetMapper) mappedWidgetAt(args ...interface{}) () {
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
    C._ZNK17QDataWidgetMapper14mappedWidgetAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedWidgetAt", args)
  }

}

// currentIndex()
func (this *QDataWidgetMapper) currentIndex(args ...interface{}) () {
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
    C._ZNK17QDataWidgetMapper12currentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "currentIndex", args)
  }

}

// model()
func (this *QDataWidgetMapper) model(args ...interface{}) () {
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
    C._ZNK17QDataWidgetMapper5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "model", args)
  }

}

// setCurrentModelIndex(const class QModelIndex &)
func (this *QDataWidgetMapper) setCurrentModelIndex(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper20setCurrentModelIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setCurrentModelIndex", args)
  }

}

// removeMapping(class QWidget *)
func (this *QDataWidgetMapper) removeMapping(args ...interface{}) () {
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
    C._ZN17QDataWidgetMapper13removeMappingEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "removeMapping", args)
  }

}

// <= body block end

