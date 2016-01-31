package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZNK17QOpenGLTimerQuery8objectIdEv(void* qthis); // 4
  // proto:  GLuint64 QOpenGLTimerQuery::waitForTimestamp();
extern void C_ZNK17QOpenGLTimerQuery16waitForTimestampEv(void* qthis); // 4
  // proto:  bool QOpenGLTimerQuery::create();
extern void C_ZN17QOpenGLTimerQuery6createEv(void* qthis); // 4
  // proto:  void QOpenGLTimerQuery::~QOpenGLTimerQuery();
extern void C_ZN17QOpenGLTimerQueryD2Ev(void* qthis); // 4
  // proto:  bool QOpenGLTimerQuery::isCreated();
extern void C_ZNK17QOpenGLTimerQuery9isCreatedEv(void* qthis); // 4
  // proto:  GLuint64 QOpenGLTimerQuery::waitForResult();
extern void C_ZNK17QOpenGLTimerQuery13waitForResultEv(void* qthis); // 4
  // proto:  void QOpenGLTimerQuery::recordTimestamp();
extern void C_ZN17QOpenGLTimerQuery15recordTimestampEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLTimerQuery::metaObject();
extern void C_ZNK17QOpenGLTimerQuery10metaObjectEv(void* qthis); // 4
  // proto:  void QOpenGLTimerQuery::destroy();
extern void C_ZN17QOpenGLTimerQuery7destroyEv(void* qthis); // 4
  // proto:  void QOpenGLTimerQuery::QOpenGLTimerQuery(QObject * parent);
extern void C_ZN17QOpenGLTimerQueryC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  bool QOpenGLTimerQuery::isResultAvailable();
extern void C_ZNK17QOpenGLTimerQuery17isResultAvailableEv(void* qthis); // 4
  // proto:  void QOpenGLTimeMonitor::reset();
extern void C_ZN18QOpenGLTimeMonitor5resetEv(void* qthis); // 4
  // proto:  int QOpenGLTimeMonitor::recordSample();
extern void C_ZN18QOpenGLTimeMonitor12recordSampleEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLTimeMonitor::metaObject();
extern void C_ZNK18QOpenGLTimeMonitor10metaObjectEv(void* qthis); // 4
  // proto:  void QOpenGLTimeMonitor::~QOpenGLTimeMonitor();
extern void C_ZN18QOpenGLTimeMonitorD2Ev(void* qthis); // 4
  // proto:  int QOpenGLTimeMonitor::sampleCount();
extern void C_ZNK18QOpenGLTimeMonitor11sampleCountEv(void* qthis); // 4
  // proto:  bool QOpenGLTimeMonitor::create();
extern void C_ZN18QOpenGLTimeMonitor6createEv(void* qthis); // 4
  // proto:  void QOpenGLTimeMonitor::setSampleCount(int sampleCount);
extern void C_ZN18QOpenGLTimeMonitor14setSampleCountEi(void* qthis, int32_t arg0); // 4
  // proto:  QVector<GLuint> QOpenGLTimeMonitor::objectIds();
extern void C_ZNK18QOpenGLTimeMonitor9objectIdsEv(void* qthis); // 4
  // proto:  bool QOpenGLTimeMonitor::isCreated();
extern void C_ZNK18QOpenGLTimeMonitor9isCreatedEv(void* qthis); // 4
  // proto:  QVector<GLuint64> QOpenGLTimeMonitor::waitForSamples();
extern void C_ZNK18QOpenGLTimeMonitor14waitForSamplesEv(void* qthis); // 4
  // proto:  bool QOpenGLTimeMonitor::isResultAvailable();
extern void C_ZNK18QOpenGLTimeMonitor17isResultAvailableEv(void* qthis); // 4
  // proto:  QVector<GLuint64> QOpenGLTimeMonitor::waitForIntervals();
extern void C_ZNK18QOpenGLTimeMonitor16waitForIntervalsEv(void* qthis); // 4
  // proto:  void QOpenGLTimeMonitor::QOpenGLTimeMonitor(QObject * parent);
extern void C_ZN18QOpenGLTimeMonitorC2EP7QObject(void* qthis, void* arg0); // 3
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QOpenGLTimeMonitor)=1
type QOpenGLTimeMonitor struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// begin()
func (this *QOpenGLTimerQuery) begin(args ...interface{}) () {
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
    C.C_ZN17QOpenGLTimerQuery5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "begin", args)
  }

}

// end()
func (this *QOpenGLTimerQuery) end(args ...interface{}) () {
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
    C.C_ZN17QOpenGLTimerQuery3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "end", args)
  }

}

// objectId()
func (this *QOpenGLTimerQuery) objectId(args ...interface{}) () {
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
    C.C_ZNK17QOpenGLTimerQuery8objectIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "objectId", args)
  }

}

// waitForTimestamp()
func (this *QOpenGLTimerQuery) waitForTimestamp(args ...interface{}) () {
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
    C.C_ZNK17QOpenGLTimerQuery16waitForTimestampEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "waitForTimestamp", args)
  }

}

// create()
func (this *QOpenGLTimerQuery) create(args ...interface{}) () {
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
    C.C_ZN17QOpenGLTimerQuery6createEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "create", args)
  }

}

// ~QOpenGLTimerQuery()
func (this *QOpenGLTimerQuery) FreeQOpenGLTimerQuery(args ...interface{}) () {
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
    C.C_ZN17QOpenGLTimerQueryD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "~QOpenGLTimerQuery", args)
  }

}

// isCreated()
func (this *QOpenGLTimerQuery) isCreated(args ...interface{}) () {
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
    C.C_ZNK17QOpenGLTimerQuery9isCreatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "isCreated", args)
  }

}

// waitForResult()
func (this *QOpenGLTimerQuery) waitForResult(args ...interface{}) () {
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
    C.C_ZNK17QOpenGLTimerQuery13waitForResultEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "waitForResult", args)
  }

}

// recordTimestamp()
func (this *QOpenGLTimerQuery) recordTimestamp(args ...interface{}) () {
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
    C.C_ZN17QOpenGLTimerQuery15recordTimestampEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "recordTimestamp", args)
  }

}

// metaObject()
func (this *QOpenGLTimerQuery) metaObject(args ...interface{}) () {
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
    C.C_ZNK17QOpenGLTimerQuery10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "metaObject", args)
  }

}

// destroy()
func (this *QOpenGLTimerQuery) destroy(args ...interface{}) () {
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
    C.C_ZN17QOpenGLTimerQuery7destroyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "destroy", args)
  }

}

// QOpenGLTimerQuery(class QObject *)
func NewQOpenGLTimerQuery(args ...interface{}) QOpenGLTimerQuery {
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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN17QOpenGLTimerQueryC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "QOpenGLTimerQuery", args)
  }

  return QOpenGLTimerQuery{}
}

// isResultAvailable()
func (this *QOpenGLTimerQuery) isResultAvailable(args ...interface{}) () {
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
    C.C_ZNK17QOpenGLTimerQuery17isResultAvailableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "isResultAvailable", args)
  }

}

// reset()
func (this *QOpenGLTimeMonitor) reset(args ...interface{}) () {
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
    C.C_ZN18QOpenGLTimeMonitor5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "reset", args)
  }

}

// recordSample()
func (this *QOpenGLTimeMonitor) recordSample(args ...interface{}) () {
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
    C.C_ZN18QOpenGLTimeMonitor12recordSampleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "recordSample", args)
  }

}

// metaObject()
func (this *QOpenGLTimeMonitor) metaObject(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLTimeMonitor10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "metaObject", args)
  }

}

// ~QOpenGLTimeMonitor()
func (this *QOpenGLTimeMonitor) FreeQOpenGLTimeMonitor(args ...interface{}) () {
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
    C.C_ZN18QOpenGLTimeMonitorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "~QOpenGLTimeMonitor", args)
  }

}

// sampleCount()
func (this *QOpenGLTimeMonitor) sampleCount(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLTimeMonitor11sampleCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "sampleCount", args)
  }

}

// create()
func (this *QOpenGLTimeMonitor) create(args ...interface{}) () {
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
    C.C_ZN18QOpenGLTimeMonitor6createEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "create", args)
  }

}

// setSampleCount(int)
func (this *QOpenGLTimeMonitor) setSampleCount(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QOpenGLTimeMonitor14setSampleCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "setSampleCount", args)
  }

}

// objectIds()
func (this *QOpenGLTimeMonitor) objectIds(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLTimeMonitor9objectIdsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "objectIds", args)
  }

}

// isCreated()
func (this *QOpenGLTimeMonitor) isCreated(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLTimeMonitor9isCreatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "isCreated", args)
  }

}

// waitForSamples()
func (this *QOpenGLTimeMonitor) waitForSamples(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLTimeMonitor14waitForSamplesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "waitForSamples", args)
  }

}

// isResultAvailable()
func (this *QOpenGLTimeMonitor) isResultAvailable(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLTimeMonitor17isResultAvailableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "isResultAvailable", args)
  }

}

// waitForIntervals()
func (this *QOpenGLTimeMonitor) waitForIntervals(args ...interface{}) () {
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
    C.C_ZNK18QOpenGLTimeMonitor16waitForIntervalsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "waitForIntervals", args)
  }

}

// QOpenGLTimeMonitor(class QObject *)
func NewQOpenGLTimeMonitor(args ...interface{}) QOpenGLTimeMonitor {
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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN18QOpenGLTimeMonitorC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "QOpenGLTimeMonitor", args)
  }

  return QOpenGLTimeMonitor{}
}

// destroy()
func (this *QOpenGLTimeMonitor) destroy(args ...interface{}) () {
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
    C.C_ZN18QOpenGLTimeMonitor7destroyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "destroy", args)
  }

}

// <= body block end

