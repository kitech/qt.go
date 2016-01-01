package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qdatawidgetmapper.h
// dst-file: /src/widgets/qdatawidgetmapper.go
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
// class sizeof(QDataWidgetMapper)=1
type QDataWidgetMapper struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _currentIndexChanged QDataWidgetMapper_currentIndexChanged_signal;
}


func NewQDataWidgetMapper(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "currentIndex", args)
 }

}


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
  case 1:
    // invoke: _ZN17QDataWidgetMapper10addMappingEP7QWidgeti
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "addMapping", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "rootIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setCurrentIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setModel", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedWidgetAt", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "removeMapping", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toFirst", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toPrevious", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setRootIndex", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "revert", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "clearMapping", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedSection", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "mappedPropertyName", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setItemDelegate", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "setCurrentModelIndex", args)
 }

}


func (this *QDataWidgetMapper) FreeQDataWidgetMapper(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "~QDataWidgetMapper", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toLast", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "model", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "itemDelegate", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "submit", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QDataWidgetMapper", "toNext", args)
 }

}

// <= body block end

