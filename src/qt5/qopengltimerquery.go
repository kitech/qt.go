package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  void QOpenGLTimerQuery::QOpenGLTimerQuery(const QOpenGLTimerQuery & );
extern void* dector_ZN17QOpenGLTimerQueryC1ERKS_(void* arg0);
extern void _ZN17QOpenGLTimerQueryC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QOpenGLTimerQuery::create();
extern void _ZN17QOpenGLTimerQuery6createEv(void* qthis);
  // proto:  bool QOpenGLTimerQuery::isCreated();
extern void _ZNK17QOpenGLTimerQuery9isCreatedEv(void* qthis);
  // proto:  void QOpenGLTimerQuery::end();
extern void _ZN17QOpenGLTimerQuery3endEv(void* qthis);
  // proto:  void QOpenGLTimerQuery::~QOpenGLTimerQuery();
extern void _ZN17QOpenGLTimerQueryD0Ev(void* qthis);
  // proto:  void QOpenGLTimerQuery::begin();
extern void _ZN17QOpenGLTimerQuery5beginEv(void* qthis);
  // proto:  void QOpenGLTimerQuery::QOpenGLTimerQuery(QObject * parent);
extern void* dector_ZN17QOpenGLTimerQueryC1EP7QObject(void* arg0);
extern void _ZN17QOpenGLTimerQueryC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QOpenGLTimerQuery::destroy();
extern void _ZN17QOpenGLTimerQuery7destroyEv(void* qthis);
  // proto:  GLuint64 QOpenGLTimerQuery::waitForResult();
extern void _ZNK17QOpenGLTimerQuery13waitForResultEv(void* qthis);
  // proto:  GLuint QOpenGLTimerQuery::objectId();
extern void _ZNK17QOpenGLTimerQuery8objectIdEv(void* qthis);
  // proto:  GLuint64 QOpenGLTimerQuery::waitForTimestamp();
extern void _ZNK17QOpenGLTimerQuery16waitForTimestampEv(void* qthis);
  // proto:  const QMetaObject * QOpenGLTimerQuery::metaObject();
extern void _ZNK17QOpenGLTimerQuery10metaObjectEv(void* qthis);
  // proto:  void QOpenGLTimerQuery::recordTimestamp();
extern void _ZN17QOpenGLTimerQuery15recordTimestampEv(void* qthis);
  // proto:  bool QOpenGLTimerQuery::isResultAvailable();
extern void _ZNK17QOpenGLTimerQuery17isResultAvailableEv(void* qthis);
  // proto:  void QOpenGLTimeMonitor::setSampleCount(int sampleCount);
extern void _ZN18QOpenGLTimeMonitor14setSampleCountEi(void* qthis, int32_t arg0);
  // proto:  int QOpenGLTimeMonitor::sampleCount();
extern void _ZNK18QOpenGLTimeMonitor11sampleCountEv(void* qthis);
  // proto:  void QOpenGLTimeMonitor::destroy();
extern void _ZN18QOpenGLTimeMonitor7destroyEv(void* qthis);
  // proto:  void QOpenGLTimeMonitor::QOpenGLTimeMonitor(const QOpenGLTimeMonitor & );
extern void* dector_ZN18QOpenGLTimeMonitorC1ERKS_(void* arg0);
extern void _ZN18QOpenGLTimeMonitorC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QOpenGLTimeMonitor::create();
extern void _ZN18QOpenGLTimeMonitor6createEv(void* qthis);
  // proto:  void QOpenGLTimeMonitor::~QOpenGLTimeMonitor();
extern void _ZN18QOpenGLTimeMonitorD0Ev(void* qthis);
  // proto:  bool QOpenGLTimeMonitor::isResultAvailable();
extern void _ZNK18QOpenGLTimeMonitor17isResultAvailableEv(void* qthis);
  // proto:  QVector<GLuint64> QOpenGLTimeMonitor::waitForIntervals();
extern void _ZNK18QOpenGLTimeMonitor16waitForIntervalsEv(void* qthis);
  // proto:  QVector<GLuint> QOpenGLTimeMonitor::objectIds();
extern void _ZNK18QOpenGLTimeMonitor9objectIdsEv(void* qthis);
  // proto:  int QOpenGLTimeMonitor::recordSample();
extern void _ZN18QOpenGLTimeMonitor12recordSampleEv(void* qthis);
  // proto:  void QOpenGLTimeMonitor::reset();
extern void _ZN18QOpenGLTimeMonitor5resetEv(void* qthis);
  // proto:  void QOpenGLTimeMonitor::QOpenGLTimeMonitor(QObject * parent);
extern void* dector_ZN18QOpenGLTimeMonitorC1EP7QObject(void* arg0);
extern void _ZN18QOpenGLTimeMonitorC1EP7QObject(void* qthis, void* arg0);
  // proto:  QVector<GLuint64> QOpenGLTimeMonitor::waitForSamples();
extern void _ZNK18QOpenGLTimeMonitor14waitForSamplesEv(void* qthis);
  // proto:  bool QOpenGLTimeMonitor::isCreated();
extern void _ZNK18QOpenGLTimeMonitor9isCreatedEv(void* qthis);
  // proto:  const QMetaObject * QOpenGLTimeMonitor::metaObject();
extern void _ZNK18QOpenGLTimeMonitor10metaObjectEv(void* qthis);
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

  // proto:  void QOpenGLTimerQuery::QOpenGLTimerQuery(const QOpenGLTimerQuery & );
func NewQOpenGLTimerQuery(args ...interface{}) QOpenGLTimerQuery {
  return QOpenGLTimerQuery{}
}

  // proto:  bool QOpenGLTimerQuery::create();
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
    C._ZN17QOpenGLTimerQuery6createEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "create", args)
  }

}

  // proto:  bool QOpenGLTimerQuery::isCreated();
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
    C._ZNK17QOpenGLTimerQuery9isCreatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "isCreated", args)
  }

}

  // proto:  void QOpenGLTimerQuery::end();
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
    C._ZN17QOpenGLTimerQuery3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "end", args)
  }

}

  // proto:  void QOpenGLTimerQuery::~QOpenGLTimerQuery();
func (this *QOpenGLTimerQuery) FreeQOpenGLTimerQuery(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "~QOpenGLTimerQuery", args)
  }

}

  // proto:  void QOpenGLTimerQuery::begin();
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
    C._ZN17QOpenGLTimerQuery5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "begin", args)
  }

}

  // proto:  void QOpenGLTimerQuery::destroy();
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
    C._ZN17QOpenGLTimerQuery7destroyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "destroy", args)
  }

}

  // proto:  GLuint64 QOpenGLTimerQuery::waitForResult();
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
    C._ZNK17QOpenGLTimerQuery13waitForResultEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "waitForResult", args)
  }

}

  // proto:  GLuint QOpenGLTimerQuery::objectId();
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
    C._ZNK17QOpenGLTimerQuery8objectIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "objectId", args)
  }

}

  // proto:  GLuint64 QOpenGLTimerQuery::waitForTimestamp();
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
    C._ZNK17QOpenGLTimerQuery16waitForTimestampEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "waitForTimestamp", args)
  }

}

  // proto:  const QMetaObject * QOpenGLTimerQuery::metaObject();
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
    C._ZNK17QOpenGLTimerQuery10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "metaObject", args)
  }

}

  // proto:  void QOpenGLTimerQuery::recordTimestamp();
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
    C._ZN17QOpenGLTimerQuery15recordTimestampEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "recordTimestamp", args)
  }

}

  // proto:  bool QOpenGLTimerQuery::isResultAvailable();
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
    C._ZNK17QOpenGLTimerQuery17isResultAvailableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimerQuery", "isResultAvailable", args)
  }

}

  // proto:  void QOpenGLTimeMonitor::setSampleCount(int sampleCount);
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
    C._ZN18QOpenGLTimeMonitor14setSampleCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "setSampleCount", args)
  }

}

  // proto:  int QOpenGLTimeMonitor::sampleCount();
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
    C._ZNK18QOpenGLTimeMonitor11sampleCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "sampleCount", args)
  }

}

  // proto:  void QOpenGLTimeMonitor::destroy();
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
    C._ZN18QOpenGLTimeMonitor7destroyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "destroy", args)
  }

}

  // proto:  void QOpenGLTimeMonitor::QOpenGLTimeMonitor(const QOpenGLTimeMonitor & );
func NewQOpenGLTimeMonitor(args ...interface{}) QOpenGLTimeMonitor {
  return QOpenGLTimeMonitor{}
}

  // proto:  bool QOpenGLTimeMonitor::create();
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
    C._ZN18QOpenGLTimeMonitor6createEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "create", args)
  }

}

  // proto:  void QOpenGLTimeMonitor::~QOpenGLTimeMonitor();
func (this *QOpenGLTimeMonitor) FreeQOpenGLTimeMonitor(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "~QOpenGLTimeMonitor", args)
  }

}

  // proto:  bool QOpenGLTimeMonitor::isResultAvailable();
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
    C._ZNK18QOpenGLTimeMonitor17isResultAvailableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "isResultAvailable", args)
  }

}

  // proto:  QVector<GLuint64> QOpenGLTimeMonitor::waitForIntervals();
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
    C._ZNK18QOpenGLTimeMonitor16waitForIntervalsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "waitForIntervals", args)
  }

}

  // proto:  QVector<GLuint> QOpenGLTimeMonitor::objectIds();
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
    C._ZNK18QOpenGLTimeMonitor9objectIdsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "objectIds", args)
  }

}

  // proto:  int QOpenGLTimeMonitor::recordSample();
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
    C._ZN18QOpenGLTimeMonitor12recordSampleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "recordSample", args)
  }

}

  // proto:  void QOpenGLTimeMonitor::reset();
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
    C._ZN18QOpenGLTimeMonitor5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "reset", args)
  }

}

  // proto:  QVector<GLuint64> QOpenGLTimeMonitor::waitForSamples();
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
    C._ZNK18QOpenGLTimeMonitor14waitForSamplesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "waitForSamples", args)
  }

}

  // proto:  bool QOpenGLTimeMonitor::isCreated();
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
    C._ZNK18QOpenGLTimeMonitor9isCreatedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "isCreated", args)
  }

}

  // proto:  const QMetaObject * QOpenGLTimeMonitor::metaObject();
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
    C._ZNK18QOpenGLTimeMonitor10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLTimeMonitor", "metaObject", args)
  }

}

// <= body block end

