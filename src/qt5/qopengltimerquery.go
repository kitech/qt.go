package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtGui/qopengltimerquery.h
// dst-file: /src/gui/qopengltimerquery.go
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
  // proto:  void QOpenGLTimerQuery::begin();
extern void C_ZN17QOpenGLTimerQuery5beginEv(void* qthis); // 4
  // proto:  void QOpenGLTimerQuery::end();
extern void C_ZN17QOpenGLTimerQuery3endEv(void* qthis); // 4
  // proto:  GLuint QOpenGLTimerQuery::objectId();
extern int32_t C_ZNK17QOpenGLTimerQuery8objectIdEv(void* qthis); // 4
  // proto:  GLuint64 QOpenGLTimerQuery::waitForTimestamp();
extern int32_t C_ZNK17QOpenGLTimerQuery16waitForTimestampEv(void* qthis); // 4
  // proto:  bool QOpenGLTimerQuery::create();
extern bool C_ZN17QOpenGLTimerQuery6createEv(void* qthis); // 4
  // proto:  void QOpenGLTimerQuery::~QOpenGLTimerQuery();
extern void C_ZN17QOpenGLTimerQueryD2Ev(void* qthis); // 4
  // proto:  bool QOpenGLTimerQuery::isCreated();
extern bool C_ZNK17QOpenGLTimerQuery9isCreatedEv(void* qthis); // 4
  // proto:  GLuint64 QOpenGLTimerQuery::waitForResult();
extern int32_t C_ZNK17QOpenGLTimerQuery13waitForResultEv(void* qthis); // 4
  // proto:  void QOpenGLTimerQuery::recordTimestamp();
extern void C_ZN17QOpenGLTimerQuery15recordTimestampEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLTimerQuery::metaObject();
extern void C_ZNK17QOpenGLTimerQuery10metaObjectEv(void* qthis); // 4
  // proto:  void QOpenGLTimerQuery::destroy();
extern void C_ZN17QOpenGLTimerQuery7destroyEv(void* qthis); // 4
  // proto:  void QOpenGLTimerQuery::QOpenGLTimerQuery(QObject * parent);
extern void* C_ZN17QOpenGLTimerQueryC2EP7QObject(void* arg0); // 3
  // proto:  bool QOpenGLTimerQuery::isResultAvailable();
extern bool C_ZNK17QOpenGLTimerQuery17isResultAvailableEv(void* qthis); // 4
  // proto:  void QOpenGLTimeMonitor::reset();
extern void C_ZN18QOpenGLTimeMonitor5resetEv(void* qthis); // 4
  // proto:  int QOpenGLTimeMonitor::recordSample();
extern int32_t C_ZN18QOpenGLTimeMonitor12recordSampleEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLTimeMonitor::metaObject();
extern void C_ZNK18QOpenGLTimeMonitor10metaObjectEv(void* qthis); // 4
  // proto:  void QOpenGLTimeMonitor::~QOpenGLTimeMonitor();
extern void C_ZN18QOpenGLTimeMonitorD2Ev(void* qthis); // 4
  // proto:  int QOpenGLTimeMonitor::sampleCount();
extern int32_t C_ZNK18QOpenGLTimeMonitor11sampleCountEv(void* qthis); // 4
  // proto:  bool QOpenGLTimeMonitor::create();
extern bool C_ZN18QOpenGLTimeMonitor6createEv(void* qthis); // 4
  // proto:  void QOpenGLTimeMonitor::setSampleCount(int sampleCount);
extern void C_ZN18QOpenGLTimeMonitor14setSampleCountEi(void* qthis, int32_t arg0); // 4
  // proto:  QVector<GLuint> QOpenGLTimeMonitor::objectIds();
extern void C_ZNK18QOpenGLTimeMonitor9objectIdsEv(void* qthis); // 4
  // proto:  bool QOpenGLTimeMonitor::isCreated();
extern bool C_ZNK18QOpenGLTimeMonitor9isCreatedEv(void* qthis); // 4
  // proto:  QVector<GLuint64> QOpenGLTimeMonitor::waitForSamples();
extern void C_ZNK18QOpenGLTimeMonitor14waitForSamplesEv(void* qthis); // 4
  // proto:  bool QOpenGLTimeMonitor::isResultAvailable();
extern bool C_ZNK18QOpenGLTimeMonitor17isResultAvailableEv(void* qthis); // 4
  // proto:  QVector<GLuint64> QOpenGLTimeMonitor::waitForIntervals();
extern void C_ZNK18QOpenGLTimeMonitor16waitForIntervalsEv(void* qthis); // 4
  // proto:  void QOpenGLTimeMonitor::QOpenGLTimeMonitor(QObject * parent);
extern void* C_ZN18QOpenGLTimeMonitorC2EP7QObject(void* arg0); // 3
  // proto:  void QOpenGLTimeMonitor::destroy();
extern void C_ZN18QOpenGLTimeMonitor7destroyEv(void* qthis); // 4
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

// class sizeof(QOpenGLTimerQuery)=1
type QOpenGLTimerQuery struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QOpenGLTimeMonitor)=1
type QOpenGLTimeMonitor struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// begin()
func (this *QOpenGLTimerQuery) Begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQuery5beginEv
    // invoke: void begin()
    C.C_ZN17QOpenGLTimerQuery5beginEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "begin", args)
  }

  return
}

// end()
func (this *QOpenGLTimerQuery) End(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQuery3endEv
    // invoke: void end()
    C.C_ZN17QOpenGLTimerQuery3endEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "end", args)
  }

  return
}

// objectId()
func (this *QOpenGLTimerQuery) Objectid(args ...interface{}) (ret interface{}) {
  // objectId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery8objectIdEv
    // invoke: GLuint objectId()
    var ret0 = C.C_ZNK17QOpenGLTimerQuery8objectIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "objectId", args)
  }

  return
}

// waitForTimestamp()
func (this *QOpenGLTimerQuery) Waitfortimestamp(args ...interface{}) (ret interface{}) {
  // waitForTimestamp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery16waitForTimestampEv
    // invoke: GLuint64 waitForTimestamp()
    var ret0 = C.C_ZNK17QOpenGLTimerQuery16waitForTimestampEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "waitForTimestamp", args)
  }

  return
}

// create()
func (this *QOpenGLTimerQuery) Create(args ...interface{}) (ret interface{}) {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQuery6createEv
    // invoke: bool create()
    var ret0 = C.C_ZN17QOpenGLTimerQuery6createEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "create", args)
  }

  return
}

// ~QOpenGLTimerQuery()
func (this *QOpenGLTimerQuery) Freeqopengltimerquery(args ...interface{}) () {
  // ~QOpenGLTimerQuery()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQueryD0Ev
    // invoke: void ~QOpenGLTimerQuery()
    C.C_ZN17QOpenGLTimerQueryD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "~QOpenGLTimerQuery", args)
  }

  return
}

// isCreated()
func (this *QOpenGLTimerQuery) Iscreated(args ...interface{}) (ret interface{}) {
  // isCreated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery9isCreatedEv
    // invoke: bool isCreated()
    var ret0 = C.C_ZNK17QOpenGLTimerQuery9isCreatedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "isCreated", args)
  }

  return
}

// waitForResult()
func (this *QOpenGLTimerQuery) Waitforresult(args ...interface{}) (ret interface{}) {
  // waitForResult()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery13waitForResultEv
    // invoke: GLuint64 waitForResult()
    var ret0 = C.C_ZNK17QOpenGLTimerQuery13waitForResultEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "waitForResult", args)
  }

  return
}

// recordTimestamp()
func (this *QOpenGLTimerQuery) Recordtimestamp(args ...interface{}) () {
  // recordTimestamp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQuery15recordTimestampEv
    // invoke: void recordTimestamp()
    C.C_ZN17QOpenGLTimerQuery15recordTimestampEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "recordTimestamp", args)
  }

  return
}

// metaObject()
func (this *QOpenGLTimerQuery) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK17QOpenGLTimerQuery10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "metaObject", args)
  }

  return
}

// destroy()
func (this *QOpenGLTimerQuery) Destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQuery7destroyEv
    // invoke: void destroy()
    C.C_ZN17QOpenGLTimerQuery7destroyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "destroy", args)
  }

  return
}

// QOpenGLTimerQuery(class QObject *)
func NewQOpenGLTimerQuery(args ...interface{}) *QOpenGLTimerQuery {
  // QOpenGLTimerQuery(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOpenGLTimerQueryC1EP7QObject
    // invoke: void QOpenGLTimerQuery(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QOpenGLTimerQueryC2EP7QObject(arg0)
    return &QOpenGLTimerQuery{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "QOpenGLTimerQuery", args)
  }

  return nil // QOpenGLTimerQuery{Qclsinst:qthis}
}

// isResultAvailable()
func (this *QOpenGLTimerQuery) Isresultavailable(args ...interface{}) (ret interface{}) {
  // isResultAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOpenGLTimerQuery17isResultAvailableEv
    // invoke: bool isResultAvailable()
    var ret0 = C.C_ZNK17QOpenGLTimerQuery17isResultAvailableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "isResultAvailable", args)
  }

  return
}

// reset()
func (this *QOpenGLTimeMonitor) Reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitor5resetEv
    // invoke: void reset()
    C.C_ZN18QOpenGLTimeMonitor5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "reset", args)
  }

  return
}

// recordSample()
func (this *QOpenGLTimeMonitor) Recordsample(args ...interface{}) (ret interface{}) {
  // recordSample()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitor12recordSampleEv
    // invoke: int recordSample()
    var ret0 = C.C_ZN18QOpenGLTimeMonitor12recordSampleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "recordSample", args)
  }

  return
}

// metaObject()
func (this *QOpenGLTimeMonitor) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QOpenGLTimeMonitor10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "metaObject", args)
  }

  return
}

// ~QOpenGLTimeMonitor()
func (this *QOpenGLTimeMonitor) Freeqopengltimemonitor(args ...interface{}) () {
  // ~QOpenGLTimeMonitor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitorD0Ev
    // invoke: void ~QOpenGLTimeMonitor()
    C.C_ZN18QOpenGLTimeMonitorD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "~QOpenGLTimeMonitor", args)
  }

  return
}

// sampleCount()
func (this *QOpenGLTimeMonitor) Samplecount(args ...interface{}) (ret interface{}) {
  // sampleCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor11sampleCountEv
    // invoke: int sampleCount()
    var ret0 = C.C_ZNK18QOpenGLTimeMonitor11sampleCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "sampleCount", args)
  }

  return
}

// create()
func (this *QOpenGLTimeMonitor) Create(args ...interface{}) (ret interface{}) {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitor6createEv
    // invoke: bool create()
    var ret0 = C.C_ZN18QOpenGLTimeMonitor6createEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "create", args)
  }

  return
}

// setSampleCount(int)
func (this *QOpenGLTimeMonitor) Setsamplecount(args ...interface{}) () {
  // setSampleCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitor14setSampleCountEi
    // invoke: void setSampleCount(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLTimeMonitor14setSampleCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "setSampleCount", args)
  }

  return
}

// objectIds()
func (this *QOpenGLTimeMonitor) Objectids(args ...interface{}) () {
  // objectIds()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor9objectIdsEv
    // invoke: QVector<GLuint> objectIds()
    C.C_ZNK18QOpenGLTimeMonitor9objectIdsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "objectIds", args)
  }

  return
}

// isCreated()
func (this *QOpenGLTimeMonitor) Iscreated(args ...interface{}) (ret interface{}) {
  // isCreated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor9isCreatedEv
    // invoke: bool isCreated()
    var ret0 = C.C_ZNK18QOpenGLTimeMonitor9isCreatedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "isCreated", args)
  }

  return
}

// waitForSamples()
func (this *QOpenGLTimeMonitor) Waitforsamples(args ...interface{}) () {
  // waitForSamples()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor14waitForSamplesEv
    // invoke: QVector<GLuint64> waitForSamples()
    C.C_ZNK18QOpenGLTimeMonitor14waitForSamplesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "waitForSamples", args)
  }

  return
}

// isResultAvailable()
func (this *QOpenGLTimeMonitor) Isresultavailable(args ...interface{}) (ret interface{}) {
  // isResultAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor17isResultAvailableEv
    // invoke: bool isResultAvailable()
    var ret0 = C.C_ZNK18QOpenGLTimeMonitor17isResultAvailableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "isResultAvailable", args)
  }

  return
}

// waitForIntervals()
func (this *QOpenGLTimeMonitor) Waitforintervals(args ...interface{}) () {
  // waitForIntervals()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLTimeMonitor16waitForIntervalsEv
    // invoke: QVector<GLuint64> waitForIntervals()
    C.C_ZNK18QOpenGLTimeMonitor16waitForIntervalsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "waitForIntervals", args)
  }

  return
}

// QOpenGLTimeMonitor(class QObject *)
func NewQOpenGLTimeMonitor(args ...interface{}) *QOpenGLTimeMonitor {
  // QOpenGLTimeMonitor(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitorC1EP7QObject
    // invoke: void QOpenGLTimeMonitor(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QOpenGLTimeMonitorC2EP7QObject(arg0)
    return &QOpenGLTimeMonitor{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "QOpenGLTimeMonitor", args)
  }

  return nil // QOpenGLTimeMonitor{Qclsinst:qthis}
}

// destroy()
func (this *QOpenGLTimeMonitor) Destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLTimeMonitor7destroyEv
    // invoke: void destroy()
    C.C_ZN18QOpenGLTimeMonitor7destroyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "destroy", args)
  }

  return
}

// <= body block end

