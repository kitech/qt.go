package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void C_ZNK9QSaveFile10metaObjectEv(void* qthis); // 4
  // proto:  void QSaveFile::setDirectWriteFallback(bool enabled);
extern void C_ZN9QSaveFile22setDirectWriteFallbackEb(void* qthis, bool arg0); // 4
  // proto:  void QSaveFile::~QSaveFile();
extern void C_ZN9QSaveFileD2Ev(void* qthis); // 4
  // proto:  void QSaveFile::cancelWriting();
extern void C_ZN9QSaveFile13cancelWritingEv(void* qthis); // 4
  // proto:  void QSaveFile::setFileName(const QString & name);
extern void C_ZN9QSaveFile11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QSaveFile::fileName();
extern void* C_ZNK9QSaveFile8fileNameEv(void* qthis); // 4
  // proto:  void QSaveFile::QSaveFile(const QString & name);
extern void* C_ZN9QSaveFileC2ERK7QString(void* arg0); // 3
  // proto:  void QSaveFile::QSaveFile(const QString & name, QObject * parent);
extern void* C_ZN9QSaveFileC2ERK7QStringP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QSaveFile::QSaveFile(QObject * parent);
extern void* C_ZN9QSaveFileC2EP7QObject(void* arg0); // 3
  // proto:  bool QSaveFile::directWriteFallback();
extern bool C_ZNK9QSaveFile19directWriteFallbackEv(void* qthis); // 4
  // proto:  bool QSaveFile::commit();
extern bool C_ZN9QSaveFile6commitEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QSaveFile) Metaobject(args ...interface{}) () {
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
    C.C_ZNK9QSaveFile10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSaveFile", "metaObject", args)
  }

  return
}

// setDirectWriteFallback(_Bool)
func (this *QSaveFile) Setdirectwritefallback(args ...interface{}) () {
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
    C.C_ZN9QSaveFile22setDirectWriteFallbackEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSaveFile", "setDirectWriteFallback", args)
  }

  return
}

// ~QSaveFile()
func (this *QSaveFile) Freeqsavefile(args ...interface{}) () {
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
    C.C_ZN9QSaveFileD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSaveFile", "~QSaveFile", args)
  }

  return
}

// cancelWriting()
func (this *QSaveFile) Cancelwriting(args ...interface{}) () {
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
    C.C_ZN9QSaveFile13cancelWritingEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSaveFile", "cancelWriting", args)
  }

  return
}

// setFileName(const class QString &)
func (this *QSaveFile) Setfilename(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSaveFile11setFileNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSaveFile", "setFileName", args)
  }

  return
}

// fileName()
func (this *QSaveFile) Filename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QSaveFile8fileNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSaveFile", "fileName", args)
  }

  return
}

// QSaveFile(const class QString &)
func NewQSaveFile(args ...interface{}) *QSaveFile {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QSaveFileC2ERK7QString(arg0)
    return &QSaveFile{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QSaveFileC1ERK7QStringP7QObject
    // invoke: void QSaveFile(const class QString &, class QObject *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QSaveFileC2ERK7QStringP7QObject(arg0, arg1)
    return &QSaveFile{Qclsinst:qthis}
  case 2:
    // invoke: _ZN9QSaveFileC1EP7QObject
    // invoke: void QSaveFile(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QSaveFileC2EP7QObject(arg0)
    return &QSaveFile{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSaveFile", "QSaveFile", args)
  }

  return nil // QSaveFile{Qclsinst:qthis}
}

// directWriteFallback()
func (this *QSaveFile) Directwritefallback(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QSaveFile19directWriteFallbackEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSaveFile", "directWriteFallback", args)
  }

  return
}

// commit()
func (this *QSaveFile) Commit(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN9QSaveFile6commitEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSaveFile", "commit", args)
  }

  return
}

// <= body block end

