package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qsignalmapper.h
// dst-file: /src/core/qsignalmapper.go
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
  // proto:  void QSignalMapper::removeMappings(QObject * sender);
extern void _ZN13QSignalMapper14removeMappingsEP7QObject(void* qthis, void* arg0);
  // proto:  void QSignalMapper::map(QObject * sender);
extern void _ZN13QSignalMapper3mapEP7QObject(void* qthis, void* arg0);
  // proto:  void QSignalMapper::QSignalMapper(const QSignalMapper & );
extern void* dector_ZN13QSignalMapperC1ERKS_(void* arg0);
extern void _ZN13QSignalMapperC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QSignalMapper::metaObject();
extern void _ZNK13QSignalMapper10metaObjectEv(void* qthis);
  // proto:  void QSignalMapper::setMapping(QObject * sender, QObject * object);
extern void _ZN13QSignalMapper10setMappingEP7QObjectS1_(void* qthis, void* arg0, void* arg1);
  // proto:  QObject * QSignalMapper::mapping(int id);
extern void _ZNK13QSignalMapper7mappingEi(void* qthis, int arg0);
  // proto:  void QSignalMapper::QSignalMapper(QObject * parent);
extern void* dector_ZN13QSignalMapperC1EP7QObject(void* arg0);
extern void _ZN13QSignalMapperC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QSignalMapper::~QSignalMapper();
extern void _ZN13QSignalMapperD0Ev(void* qthis);
  // proto:  void QSignalMapper::setMapping(QObject * sender, int id);
extern void _ZN13QSignalMapper10setMappingEP7QObjecti(void* qthis, void* arg0, int arg1);
  // proto:  QObject * QSignalMapper::mapping(const QString & text);
extern void _ZNK13QSignalMapper7mappingERK7QString(void* qthis, void* arg0);
  // proto:  void QSignalMapper::map();
extern void _ZN13QSignalMapper3mapEv(void* qthis);
  // proto:  QObject * QSignalMapper::mapping(QObject * object);
extern void _ZNK13QSignalMapper7mappingEP7QObject(void* qthis, void* arg0);
  // proto:  void QSignalMapper::setMapping(QObject * sender, const QString & text);
extern void _ZN13QSignalMapper10setMappingEP7QObjectRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QObject * QSignalMapper::mapping(QWidget * widget);
extern void _ZNK13QSignalMapper7mappingEP7QWidget(void* qthis, void* arg0);
  // proto:  void QSignalMapper::setMapping(QObject * sender, QWidget * widget);
extern void _ZN13QSignalMapper10setMappingEP7QObjectP7QWidget(void* qthis, void* arg0, void* arg1);
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

// class sizeof(QSignalMapper)=1
type QSignalMapper struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _mapped QSignalMapper_mapped_signal;
}

  // proto:  void QSignalMapper::removeMappings(QObject * sender);
func (this *QSignalMapper) removeMappings(args ...interface{}) () {
  // removeMappings(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSignalMapper14removeMappingsEP7QObject
  default:
    qtrt.ErrorResolve("QSignalMapper", "removeMappings", args)
  }

}

  // proto:  void QSignalMapper::map(QObject * sender);
func (this *QSignalMapper) map_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSignalMapper", "map", args)
  }

}

  // proto:  void QSignalMapper::QSignalMapper(const QSignalMapper & );
func NewQSignalMapper(args ...interface{}) QSignalMapper {
  return QSignalMapper{}
}

  // proto:  const QMetaObject * QSignalMapper::metaObject();
func (this *QSignalMapper) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSignalMapper10metaObjectEv
  default:
    qtrt.ErrorResolve("QSignalMapper", "metaObject", args)
  }

}

  // proto:  void QSignalMapper::setMapping(QObject * sender, QObject * object);
func (this *QSignalMapper) setMapping(args ...interface{}) () {
  // setMapping(class QObject *, class QObject *)
  // setMapping(class QObject *, int)
  // setMapping(class QObject *, const class QString &)
  // setMapping(class QObject *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[3][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjectS1_
  case 1:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjecti
  case 2:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjectRK7QString
  case 3:
    // invoke: _ZN13QSignalMapper10setMappingEP7QObjectP7QWidget
  default:
    qtrt.ErrorResolve("QSignalMapper", "setMapping", args)
  }

}

  // proto:  QObject * QSignalMapper::mapping(int id);
func (this *QSignalMapper) mapping(args ...interface{}) () {
  // mapping(int)
  // mapping(const class QString &)
  // mapping(class QObject *)
  // mapping(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSignalMapper7mappingEi
  case 1:
    // invoke: _ZNK13QSignalMapper7mappingERK7QString
  case 2:
    // invoke: _ZNK13QSignalMapper7mappingEP7QObject
  case 3:
    // invoke: _ZNK13QSignalMapper7mappingEP7QWidget
  default:
    qtrt.ErrorResolve("QSignalMapper", "mapping", args)
  }

}

  // proto:  void QSignalMapper::~QSignalMapper();
func (this *QSignalMapper) FreeQSignalMapper(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSignalMapper", "~QSignalMapper", args)
  }

}

// <= body block end

