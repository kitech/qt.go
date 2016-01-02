package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QSaveFile::cancelWriting();
extern void _ZN9QSaveFile13cancelWritingEv(void* qthis);
  // proto:  void QSaveFile::QSaveFile(QObject * parent);
extern void* dector_ZN9QSaveFileC1EP7QObject(void* arg0);
extern void _ZN9QSaveFileC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QSaveFile::QSaveFile(const QSaveFile & );
extern void* dector_ZN9QSaveFileC1ERKS_(void* arg0);
extern void _ZN9QSaveFileC1ERKS_(void* qthis, void* arg0);
  // proto:  void QSaveFile::QSaveFile(const QString & name, QObject * parent);
extern void* dector_ZN9QSaveFileC1ERK7QStringP7QObject(void* arg0, void* arg1);
extern void _ZN9QSaveFileC1ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  QString QSaveFile::fileName();
extern void _ZNK9QSaveFile8fileNameEv(void* qthis);
  // proto:  void QSaveFile::QSaveFile(const QString & name);
extern void* dector_ZN9QSaveFileC1ERK7QString(void* arg0);
extern void _ZN9QSaveFileC1ERK7QString(void* qthis, void* arg0);
  // proto:  const QMetaObject * QSaveFile::metaObject();
extern void _ZNK9QSaveFile10metaObjectEv(void* qthis);
  // proto:  bool QSaveFile::commit();
extern void _ZN9QSaveFile6commitEv(void* qthis);
  // proto:  void QSaveFile::~QSaveFile();
extern void _ZN9QSaveFileD0Ev(void* qthis);
  // proto:  void QSaveFile::setFileName(const QString & name);
extern void _ZN9QSaveFile11setFileNameERK7QString(void* qthis, void* arg0);
  // proto:  bool QSaveFile::directWriteFallback();
extern void _ZNK9QSaveFile19directWriteFallbackEv(void* qthis);
  // proto:  void QSaveFile::setDirectWriteFallback(bool enabled);
extern void _ZN9QSaveFile22setDirectWriteFallbackEb(void* qthis, bool arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QSaveFile::cancelWriting();
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

  // proto:  void QSaveFile::QSaveFile(QObject * parent);
func NewQSaveFile(args ...interface{}) QSaveFile {
  return QSaveFile{}
}

  // proto:  QString QSaveFile::fileName();
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

  // proto:  const QMetaObject * QSaveFile::metaObject();
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

  // proto:  bool QSaveFile::commit();
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

  // proto:  void QSaveFile::~QSaveFile();
func (this *QSaveFile) FreeQSaveFile(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSaveFile", "~QSaveFile", args)
  }

}

  // proto:  void QSaveFile::setFileName(const QString & name);
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

  // proto:  bool QSaveFile::directWriteFallback();
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

  // proto:  void QSaveFile::setDirectWriteFallback(bool enabled);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN9QSaveFile22setDirectWriteFallbackEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSaveFile", "setDirectWriteFallback", args)
  }

}

// <= body block end

