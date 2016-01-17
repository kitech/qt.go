package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
// src-file: /QtGui/qopenglcontext.h
// dst-file: /src/gui/qopenglcontext.go
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
  // proto:  QSurfaceFormat::OpenGLContextProfile QOpenGLVersionProfile::profile();
extern void _ZNK21QOpenGLVersionProfile7profileEv(void* qthis); // 4
  // proto:  void QOpenGLVersionProfile::QOpenGLVersionProfile(const QSurfaceFormat & format);
extern void _ZN21QOpenGLVersionProfileC2ERK14QSurfaceFormat(void* qthis, void* arg0); // 3
  // proto:  void QOpenGLVersionProfile::QOpenGLVersionProfile();
extern void _ZN21QOpenGLVersionProfileC2Ev(void* qthis); // 3
  // proto:  void QOpenGLVersionProfile::QOpenGLVersionProfile(const QOpenGLVersionProfile & other);
extern void _ZN21QOpenGLVersionProfileC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  bool QOpenGLVersionProfile::hasProfiles();
extern void _ZNK21QOpenGLVersionProfile11hasProfilesEv(void* qthis); // 4
  // proto:  bool QOpenGLVersionProfile::isValid();
extern void _ZNK21QOpenGLVersionProfile7isValidEv(void* qthis); // 4
  // proto:  void QOpenGLVersionProfile::~QOpenGLVersionProfile();
extern void _ZN21QOpenGLVersionProfileD2Ev(void* qthis); // 4
  // proto:  QPair<int, int> QOpenGLVersionProfile::version();
extern void _ZNK21QOpenGLVersionProfile7versionEv(void* qthis); // 4
  // proto:  bool QOpenGLVersionProfile::isLegacyVersion();
extern void _ZNK21QOpenGLVersionProfile15isLegacyVersionEv(void* qthis); // 4
  // proto:  void QOpenGLVersionProfile::setVersion(int majorVersion, int minorVersion);
extern void _ZN21QOpenGLVersionProfile10setVersionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QOpenGLContext::~QOpenGLContext();
extern void _ZN14QOpenGLContextD2Ev(void* qthis); // 4
  // proto: static void * QOpenGLContext::openGLModuleHandle();
extern void _ZN14QOpenGLContext18openGLModuleHandleEv(); // 4
  // proto:  QOpenGLContext * QOpenGLContext::shareContext();
extern void _ZNK14QOpenGLContext12shareContextEv(void* qthis); // 4
  // proto:  void QOpenGLContext::doneCurrent();
extern void _ZN14QOpenGLContext11doneCurrentEv(void* qthis); // 4
  // proto:  QSurface * QOpenGLContext::surface();
extern void _ZNK14QOpenGLContext7surfaceEv(void* qthis); // 4
  // proto:  void QOpenGLContext::setNativeHandle(const QVariant & handle);
extern void _ZN14QOpenGLContext15setNativeHandleERK8QVariant(void* qthis, void* arg0); // 4
  // proto: static bool QOpenGLContext::supportsThreadedOpenGL();
extern void _ZN14QOpenGLContext22supportsThreadedOpenGLEv(); // 4
  // proto:  QOpenGLFunctions * QOpenGLContext::functions();
extern void _ZNK14QOpenGLContext9functionsEv(void* qthis); // 4
  // proto: static bool QOpenGLContext::areSharing(QOpenGLContext * first, QOpenGLContext * second);
extern void _ZN14QOpenGLContext10areSharingEPS_S0_(void* arg0, void* arg1); // 4
  // proto:  void QOpenGLContext::setScreen(QScreen * screen);
extern void _ZN14QOpenGLContext9setScreenEP7QScreen(void* qthis, void* arg0); // 4
  // proto: static QOpenGLContext * QOpenGLContext::currentContext();
extern void _ZN14QOpenGLContext14currentContextEv(); // 4
  // proto:  void QOpenGLContext::setShareContext(QOpenGLContext * shareContext);
extern void _ZN14QOpenGLContext15setShareContextEPS_(void* qthis, void* arg0); // 4
  // proto: static QOpenGLContext::OpenGLModuleType QOpenGLContext::openGLModuleType();
extern void _ZN14QOpenGLContext16openGLModuleTypeEv(); // 4
  // proto: static QOpenGLContext * QOpenGLContext::globalShareContext();
extern void _ZN14QOpenGLContext18globalShareContextEv(); // 4
  // proto:  bool QOpenGLContext::isOpenGLES();
extern void _ZNK14QOpenGLContext10isOpenGLESEv(void* qthis); // 4
  // proto:  QAbstractOpenGLFunctions * QOpenGLContext::versionFunctions(const QOpenGLVersionProfile & versionProfile);
extern void _ZNK14QOpenGLContext16versionFunctionsERK21QOpenGLVersionProfile(void* qthis, void* arg0); // 4
  // proto:  QScreen * QOpenGLContext::screen();
extern void _ZNK14QOpenGLContext6screenEv(void* qthis); // 4
  // proto:  void QOpenGLContext::swapBuffers(QSurface * surface);
extern void _ZN14QOpenGLContext11swapBuffersEP8QSurface(void* qthis, void* arg0); // 4
  // proto:  QPlatformOpenGLContext * QOpenGLContext::handle();
extern void _ZNK14QOpenGLContext6handleEv(void* qthis); // 4
  // proto:  bool QOpenGLContext::isValid();
extern void _ZNK14QOpenGLContext7isValidEv(void* qthis); // 4
  // proto:  bool QOpenGLContext::hasExtension(const QByteArray & extension);
extern void _ZNK14QOpenGLContext12hasExtensionERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  GLuint QOpenGLContext::defaultFramebufferObject();
extern void _ZNK14QOpenGLContext24defaultFramebufferObjectEv(void* qthis); // 4
  // proto:  QFunctionPointer QOpenGLContext::getProcAddress(const QByteArray & procName);
extern void _ZNK14QOpenGLContext14getProcAddressERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QOpenGLContextGroup * QOpenGLContext::shareGroup();
extern void _ZNK14QOpenGLContext10shareGroupEv(void* qthis); // 4
  // proto:  QPlatformOpenGLContext * QOpenGLContext::shareHandle();
extern void _ZNK14QOpenGLContext11shareHandleEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLContext::metaObject();
extern void _ZNK14QOpenGLContext10metaObjectEv(void* qthis); // 4
  // proto:  void QOpenGLContext::QOpenGLContext(QObject * parent);
extern void _ZN14QOpenGLContextC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QOpenGLContext::setFormat(const QSurfaceFormat & format);
extern void _ZN14QOpenGLContext9setFormatERK14QSurfaceFormat(void* qthis, void* arg0); // 4
  // proto:  bool QOpenGLContext::makeCurrent(QSurface * surface);
extern void _ZN14QOpenGLContext11makeCurrentEP8QSurface(void* qthis, void* arg0); // 4
  // proto:  QVariant QOpenGLContext::nativeHandle();
extern void _ZNK14QOpenGLContext12nativeHandleEv(void* qthis); // 4
  // proto:  QSet<QByteArray> QOpenGLContext::extensions();
extern void _ZNK14QOpenGLContext10extensionsEv(void* qthis); // 4
  // proto:  QSurfaceFormat QOpenGLContext::format();
extern void _ZNK14QOpenGLContext6formatEv(void* qthis); // 4
  // proto:  bool QOpenGLContext::create();
extern void _ZN14QOpenGLContext6createEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLContextGroup::metaObject();
extern void _ZNK19QOpenGLContextGroup10metaObjectEv(void* qthis); // 4
  // proto: static QOpenGLContextGroup * QOpenGLContextGroup::currentContextGroup();
extern void _ZN19QOpenGLContextGroup19currentContextGroupEv(); // 4
  // proto:  QList<QOpenGLContext *> QOpenGLContextGroup::shares();
extern void _ZNK19QOpenGLContextGroup6sharesEv(void* qthis); // 4
  // proto:  void QOpenGLContextGroup::~QOpenGLContextGroup();
extern void _ZN19QOpenGLContextGroupD2Ev(void* qthis); // 4
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

// class sizeof(QOpenGLVersionProfile)=8
type QOpenGLVersionProfile struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QOpenGLContext)=1
type QOpenGLContext struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToBeDestroyed QOpenGLContext_aboutToBeDestroyed_signal;
}

// class sizeof(QOpenGLContextGroup)=1
type QOpenGLContextGroup struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// profile()
func (this *QOpenGLVersionProfile) profile(args ...interface{}) () {
  // profile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QOpenGLVersionProfile7profileEv
    // invoke: QSurfaceFormat::OpenGLContextProfile profile()
    C._ZNK21QOpenGLVersionProfile7profileEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "profile", args)
  }

}

// QOpenGLVersionProfile(const class QSurfaceFormat &)
func NewQOpenGLVersionProfile(args ...interface{}) QOpenGLVersionProfile {
  // QOpenGLVersionProfile(const class QSurfaceFormat &)
  // QOpenGLVersionProfile()
  // QOpenGLVersionProfile(const class QOpenGLVersionProfile &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurfaceFormat{}) // "const QSurfaceFormat &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QOpenGLVersionProfile{}) // "const QOpenGLVersionProfile &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QOpenGLVersionProfileC1ERK14QSurfaceFormat
    // invoke: void QOpenGLVersionProfile(const class QSurfaceFormat &)
    var arg0 = args[0].(QSurfaceFormat).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN21QOpenGLVersionProfileC2ERK14QSurfaceFormat(qthis, arg0)
  case 1:
    // invoke: _ZN21QOpenGLVersionProfileC1Ev
    // invoke: void QOpenGLVersionProfile()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN21QOpenGLVersionProfileC2Ev(qthis)
  case 2:
    // invoke: _ZN21QOpenGLVersionProfileC1ERKS_
    // invoke: void QOpenGLVersionProfile(const class QOpenGLVersionProfile &)
    var arg0 = args[0].(QOpenGLVersionProfile).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN21QOpenGLVersionProfileC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "QOpenGLVersionProfile", args)
  }

  return QOpenGLVersionProfile{}
}

// hasProfiles()
func (this *QOpenGLVersionProfile) hasProfiles(args ...interface{}) () {
  // hasProfiles()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QOpenGLVersionProfile11hasProfilesEv
    // invoke: bool hasProfiles()
    C._ZNK21QOpenGLVersionProfile11hasProfilesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "hasProfiles", args)
  }

}

// isValid()
func (this *QOpenGLVersionProfile) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QOpenGLVersionProfile7isValidEv
    // invoke: bool isValid()
    C._ZNK21QOpenGLVersionProfile7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "isValid", args)
  }

}

// ~QOpenGLVersionProfile()
func (this *QOpenGLVersionProfile) FreeQOpenGLVersionProfile(args ...interface{}) () {
  // ~QOpenGLVersionProfile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QOpenGLVersionProfileD0Ev
    // invoke: void ~QOpenGLVersionProfile()
    C._ZN21QOpenGLVersionProfileD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "~QOpenGLVersionProfile", args)
  }

}

// version()
func (this *QOpenGLVersionProfile) version(args ...interface{}) () {
  // version()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QOpenGLVersionProfile7versionEv
    // invoke: QPair<int, int> version()
    C._ZNK21QOpenGLVersionProfile7versionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "version", args)
  }

}

// isLegacyVersion()
func (this *QOpenGLVersionProfile) isLegacyVersion(args ...interface{}) () {
  // isLegacyVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QOpenGLVersionProfile15isLegacyVersionEv
    // invoke: bool isLegacyVersion()
    C._ZNK21QOpenGLVersionProfile15isLegacyVersionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "isLegacyVersion", args)
  }

}

// setVersion(int, int)
func (this *QOpenGLVersionProfile) setVersion(args ...interface{}) () {
  // setVersion(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QOpenGLVersionProfile10setVersionEii
    // invoke: void setVersion(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN21QOpenGLVersionProfile10setVersionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "setVersion", args)
  }

}

// ~QOpenGLContext()
func (this *QOpenGLContext) FreeQOpenGLContext(args ...interface{}) () {
  // ~QOpenGLContext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContextD0Ev
    // invoke: void ~QOpenGLContext()
    C._ZN14QOpenGLContextD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "~QOpenGLContext", args)
  }

}

// openGLModuleHandle()
func (this *QOpenGLContext) openGLModuleHandle_s(args ...interface{}) () {
  // openGLModuleHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext18openGLModuleHandleEv
    // invoke: void * openGLModuleHandle()
    C._ZN14QOpenGLContext18openGLModuleHandleEv()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "openGLModuleHandle", args)
  }

}

// shareContext()
func (this *QOpenGLContext) shareContext(args ...interface{}) () {
  // shareContext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext12shareContextEv
    // invoke: QOpenGLContext * shareContext()
    C._ZNK14QOpenGLContext12shareContextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "shareContext", args)
  }

}

// doneCurrent()
func (this *QOpenGLContext) doneCurrent(args ...interface{}) () {
  // doneCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext11doneCurrentEv
    // invoke: void doneCurrent()
    C._ZN14QOpenGLContext11doneCurrentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "doneCurrent", args)
  }

}

// surface()
func (this *QOpenGLContext) surface(args ...interface{}) () {
  // surface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext7surfaceEv
    // invoke: QSurface * surface()
    C._ZNK14QOpenGLContext7surfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "surface", args)
  }

}

// setNativeHandle(const class QVariant &)
func (this *QOpenGLContext) setNativeHandle(args ...interface{}) () {
  // setNativeHandle(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext15setNativeHandleERK8QVariant
    // invoke: void setNativeHandle(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLContext15setNativeHandleERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setNativeHandle", args)
  }

}

// supportsThreadedOpenGL()
func (this *QOpenGLContext) supportsThreadedOpenGL_s(args ...interface{}) () {
  // supportsThreadedOpenGL()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext22supportsThreadedOpenGLEv
    // invoke: bool supportsThreadedOpenGL()
    C._ZN14QOpenGLContext22supportsThreadedOpenGLEv()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "supportsThreadedOpenGL", args)
  }

}

// functions()
func (this *QOpenGLContext) functions(args ...interface{}) () {
  // functions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext9functionsEv
    // invoke: QOpenGLFunctions * functions()
    C._ZNK14QOpenGLContext9functionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "functions", args)
  }

}

// areSharing(class QOpenGLContext *, class QOpenGLContext *)
func (this *QOpenGLContext) areSharing_s(args ...interface{}) () {
  // areSharing(class QOpenGLContext *, class QOpenGLContext *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLContext{}) // "QOpenGLContext *"
  vtys[0][1] = reflect.TypeOf(QOpenGLContext{}) // "QOpenGLContext *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext10areSharingEPS_S0_
    // invoke: bool areSharing(class QOpenGLContext *, class QOpenGLContext *)
    var arg0 = args[0].(QOpenGLContext).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QOpenGLContext).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN14QOpenGLContext10areSharingEPS_S0_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "areSharing", args)
  }

}

// setScreen(class QScreen *)
func (this *QOpenGLContext) setScreen(args ...interface{}) () {
  // setScreen(class QScreen *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScreen{}) // "QScreen *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext9setScreenEP7QScreen
    // invoke: void setScreen(class QScreen *)
    var arg0 = args[0].(QScreen).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLContext9setScreenEP7QScreen(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setScreen", args)
  }

}

// currentContext()
func (this *QOpenGLContext) currentContext_s(args ...interface{}) () {
  // currentContext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext14currentContextEv
    // invoke: QOpenGLContext * currentContext()
    C._ZN14QOpenGLContext14currentContextEv()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "currentContext", args)
  }

}

// setShareContext(class QOpenGLContext *)
func (this *QOpenGLContext) setShareContext(args ...interface{}) () {
  // setShareContext(class QOpenGLContext *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLContext{}) // "QOpenGLContext *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext15setShareContextEPS_
    // invoke: void setShareContext(class QOpenGLContext *)
    var arg0 = args[0].(QOpenGLContext).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLContext15setShareContextEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setShareContext", args)
  }

}

// openGLModuleType()
func (this *QOpenGLContext) openGLModuleType_s(args ...interface{}) () {
  // openGLModuleType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext16openGLModuleTypeEv
    // invoke: QOpenGLContext::OpenGLModuleType openGLModuleType()
    C._ZN14QOpenGLContext16openGLModuleTypeEv()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "openGLModuleType", args)
  }

}

// globalShareContext()
func (this *QOpenGLContext) globalShareContext_s(args ...interface{}) () {
  // globalShareContext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext18globalShareContextEv
    // invoke: QOpenGLContext * globalShareContext()
    C._ZN14QOpenGLContext18globalShareContextEv()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "globalShareContext", args)
  }

}

// isOpenGLES()
func (this *QOpenGLContext) isOpenGLES(args ...interface{}) () {
  // isOpenGLES()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext10isOpenGLESEv
    // invoke: bool isOpenGLES()
    C._ZNK14QOpenGLContext10isOpenGLESEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "isOpenGLES", args)
  }

}

// versionFunctions(const class QOpenGLVersionProfile &)
func (this *QOpenGLContext) versionFunctions(args ...interface{}) () {
  // versionFunctions(const class QOpenGLVersionProfile &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLVersionProfile{}) // "const QOpenGLVersionProfile &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext16versionFunctionsERK21QOpenGLVersionProfile
    // invoke: QAbstractOpenGLFunctions * versionFunctions(const class QOpenGLVersionProfile &)
    var arg0 = args[0].(QOpenGLVersionProfile).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QOpenGLContext16versionFunctionsERK21QOpenGLVersionProfile(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "versionFunctions", args)
  }

}

// screen()
func (this *QOpenGLContext) screen(args ...interface{}) () {
  // screen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext6screenEv
    // invoke: QScreen * screen()
    C._ZNK14QOpenGLContext6screenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "screen", args)
  }

}

// swapBuffers(class QSurface *)
func (this *QOpenGLContext) swapBuffers(args ...interface{}) () {
  // swapBuffers(class QSurface *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurface{}) // "QSurface *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext11swapBuffersEP8QSurface
    // invoke: void swapBuffers(class QSurface *)
    var arg0 = args[0].(QSurface).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLContext11swapBuffersEP8QSurface(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "swapBuffers", args)
  }

}

// handle()
func (this *QOpenGLContext) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext6handleEv
    // invoke: QPlatformOpenGLContext * handle()
    C._ZNK14QOpenGLContext6handleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "handle", args)
  }

}

// isValid()
func (this *QOpenGLContext) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext7isValidEv
    // invoke: bool isValid()
    C._ZNK14QOpenGLContext7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "isValid", args)
  }

}

// hasExtension(const class QByteArray &)
func (this *QOpenGLContext) hasExtension(args ...interface{}) () {
  // hasExtension(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext12hasExtensionERK10QByteArray
    // invoke: bool hasExtension(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QOpenGLContext12hasExtensionERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "hasExtension", args)
  }

}

// defaultFramebufferObject()
func (this *QOpenGLContext) defaultFramebufferObject(args ...interface{}) () {
  // defaultFramebufferObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext24defaultFramebufferObjectEv
    // invoke: GLuint defaultFramebufferObject()
    C._ZNK14QOpenGLContext24defaultFramebufferObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "defaultFramebufferObject", args)
  }

}

// getProcAddress(const class QByteArray &)
func (this *QOpenGLContext) getProcAddress(args ...interface{}) () {
  // getProcAddress(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext14getProcAddressERK10QByteArray
    // invoke: QFunctionPointer getProcAddress(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QOpenGLContext14getProcAddressERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "getProcAddress", args)
  }

}

// shareGroup()
func (this *QOpenGLContext) shareGroup(args ...interface{}) () {
  // shareGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext10shareGroupEv
    // invoke: QOpenGLContextGroup * shareGroup()
    C._ZNK14QOpenGLContext10shareGroupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "shareGroup", args)
  }

}

// shareHandle()
func (this *QOpenGLContext) shareHandle(args ...interface{}) () {
  // shareHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext11shareHandleEv
    // invoke: QPlatformOpenGLContext * shareHandle()
    C._ZNK14QOpenGLContext11shareHandleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "shareHandle", args)
  }

}

// metaObject()
func (this *QOpenGLContext) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK14QOpenGLContext10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "metaObject", args)
  }

}

// QOpenGLContext(class QObject *)
func NewQOpenGLContext(args ...interface{}) QOpenGLContext {
  // QOpenGLContext(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContextC1EP7QObject
    // invoke: void QOpenGLContext(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN14QOpenGLContextC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "QOpenGLContext", args)
  }

  return QOpenGLContext{}
}

// setFormat(const class QSurfaceFormat &)
func (this *QOpenGLContext) setFormat(args ...interface{}) () {
  // setFormat(const class QSurfaceFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurfaceFormat{}) // "const QSurfaceFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext9setFormatERK14QSurfaceFormat
    // invoke: void setFormat(const class QSurfaceFormat &)
    var arg0 = args[0].(QSurfaceFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLContext9setFormatERK14QSurfaceFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setFormat", args)
  }

}

// makeCurrent(class QSurface *)
func (this *QOpenGLContext) makeCurrent(args ...interface{}) () {
  // makeCurrent(class QSurface *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurface{}) // "QSurface *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext11makeCurrentEP8QSurface
    // invoke: bool makeCurrent(class QSurface *)
    var arg0 = args[0].(QSurface).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QOpenGLContext11makeCurrentEP8QSurface(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "makeCurrent", args)
  }

}

// nativeHandle()
func (this *QOpenGLContext) nativeHandle(args ...interface{}) () {
  // nativeHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext12nativeHandleEv
    // invoke: QVariant nativeHandle()
    C._ZNK14QOpenGLContext12nativeHandleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "nativeHandle", args)
  }

}

// extensions()
func (this *QOpenGLContext) extensions(args ...interface{}) () {
  // extensions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext10extensionsEv
    // invoke: QSet<QByteArray> extensions()
    C._ZNK14QOpenGLContext10extensionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "extensions", args)
  }

}

// format()
func (this *QOpenGLContext) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QOpenGLContext6formatEv
    // invoke: QSurfaceFormat format()
    C._ZNK14QOpenGLContext6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "format", args)
  }

}

// create()
func (this *QOpenGLContext) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QOpenGLContext6createEv
    // invoke: bool create()
    C._ZN14QOpenGLContext6createEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "create", args)
  }

}

// metaObject()
func (this *QOpenGLContextGroup) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QOpenGLContextGroup10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK19QOpenGLContextGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "metaObject", args)
  }

}

// currentContextGroup()
func (this *QOpenGLContextGroup) currentContextGroup_s(args ...interface{}) () {
  // currentContextGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QOpenGLContextGroup19currentContextGroupEv
    // invoke: QOpenGLContextGroup * currentContextGroup()
    C._ZN19QOpenGLContextGroup19currentContextGroupEv()
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "currentContextGroup", args)
  }

}

// shares()
func (this *QOpenGLContextGroup) shares(args ...interface{}) () {
  // shares()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QOpenGLContextGroup6sharesEv
    // invoke: QList<QOpenGLContext *> shares()
    C._ZNK19QOpenGLContextGroup6sharesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "shares", args)
  }

}

// ~QOpenGLContextGroup()
func (this *QOpenGLContextGroup) FreeQOpenGLContextGroup(args ...interface{}) () {
  // ~QOpenGLContextGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QOpenGLContextGroupD0Ev
    // invoke: void ~QOpenGLContextGroup()
    C._ZN19QOpenGLContextGroupD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "~QOpenGLContextGroup", args)
  }

}

// <= body block end

