package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtCore/qsavefile.h
// dst-file: /src/core/qsavefile.go
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
  // proto:  const QMetaObject * QSaveFile::metaObject();
extern void _ZNK9QSaveFile10metaObjectEv(void* qthis); // 4
  // proto:  void QSaveFile::setDirectWriteFallback(bool enabled);
extern void _ZN9QSaveFile22setDirectWriteFallbackEb(void* qthis, bool arg0); // 4
  // proto:  void QSaveFile::~QSaveFile();
extern void _ZN9QSaveFileD2Ev(void* qthis); // 4
  // proto:  void QSaveFile::cancelWriting();
extern void _ZN9QSaveFile13cancelWritingEv(void* qthis); // 4
  // proto:  void QSaveFile::setFileName(const QString & name);
extern void _ZN9QSaveFile11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QSaveFile::fileName();
extern void _ZNK9QSaveFile8fileNameEv(void* qthis); // 4
  // proto:  void QSaveFile::QSaveFile(const QString & name);
extern void _ZN9QSaveFileC2ERK7QString(void* qthis, void* arg0); // 3
  // proto:  void QSaveFile::QSaveFile(const QString & name, QObject * parent);
extern void _ZN9QSaveFileC2ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QSaveFile::QSaveFile(QObject * parent);
extern void _ZN9QSaveFileC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  bool QSaveFile::directWriteFallback();
extern void _ZNK9QSaveFile19directWriteFallbackEv(void* qthis); // 4
  // proto:  bool QSaveFile::commit();
extern void _ZN9QSaveFile6commitEv(void* qthis); // 4
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

// class sizeof(QSaveFile)=1
type QSaveFile struct {
  /*qbase*/ QFileDevice;
  qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QSaveFile) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSaveFile10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QSaveFile10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSaveFile", "metaObject", args)
  }

}

// setDirectWriteFallback(_Bool)
func (this *QSaveFile) setDirectWriteFallback(args ...interface{}) () {
  // setDirectWriteFallback(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSaveFile22setDirectWriteFallbackEb
    // invoke: void setDirectWriteFallback(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QSaveFile22setDirectWriteFallbackEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSaveFile", "setDirectWriteFallback", args)
  }

}

// ~QSaveFile()
func (this *QSaveFile) FreeQSaveFile(args ...interface{}) () {
  // ~QSaveFile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSaveFileD0Ev
    // invoke: void ~QSaveFile()
    C._ZN9QSaveFileD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSaveFile", "~QSaveFile", args)
  }

}

// cancelWriting()
func (this *QSaveFile) cancelWriting(args ...interface{}) () {
  // cancelWriting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSaveFile13cancelWritingEv
    // invoke: void cancelWriting()
    C._ZN9QSaveFile13cancelWritingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSaveFile", "cancelWriting", args)
  }

}

// setFileName(const class QString &)
func (this *QSaveFile) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSaveFile11setFileNameERK7QString
    // invoke: void setFileName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QSaveFile11setFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSaveFile", "setFileName", args)
  }

}

// fileName()
func (this *QSaveFile) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSaveFile8fileNameEv
    // invoke: QString fileName()
    C._ZNK9QSaveFile8fileNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSaveFile", "fileName", args)
  }

}

// QSaveFile(const class QString &)
func NewQSaveFile(args ...interface{}) QSaveFile {
  // QSaveFile(const class QString &)
  // QSaveFile(const class QString &, class QObject *)
  // QSaveFile(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSaveFileC1ERK7QString
    // invoke: void QSaveFile(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QSaveFileC2ERK7QString(qthis, arg0)
  case 1:
    // invoke: _ZN9QSaveFileC1ERK7QStringP7QObject
    // invoke: void QSaveFile(const class QString &, class QObject *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QSaveFileC2ERK7QStringP7QObject(qthis, arg0, arg1)
  case 2:
    // invoke: _ZN9QSaveFileC1EP7QObject
    // invoke: void QSaveFile(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QSaveFileC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QSaveFile", "QSaveFile", args)
  }

  return QSaveFile{}
}

// directWriteFallback()
func (this *QSaveFile) directWriteFallback(args ...interface{}) () {
  // directWriteFallback()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSaveFile19directWriteFallbackEv
    // invoke: bool directWriteFallback()
    C._ZNK9QSaveFile19directWriteFallbackEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSaveFile", "directWriteFallback", args)
  }

}

// commit()
func (this *QSaveFile) commit(args ...interface{}) () {
  // commit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSaveFile6commitEv
    // invoke: bool commit()
    C._ZN9QSaveFile6commitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSaveFile", "commit", args)
  }

}

// <= body block end

