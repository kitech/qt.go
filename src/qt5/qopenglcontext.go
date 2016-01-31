package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern void C_ZNK21QOpenGLVersionProfile7profileEv(void* qthis); // 4
  // proto:  void QOpenGLVersionProfile::QOpenGLVersionProfile(const QSurfaceFormat & format);
extern void* C_ZN21QOpenGLVersionProfileC2ERK14QSurfaceFormat(void* arg0); // 3
  // proto:  void QOpenGLVersionProfile::QOpenGLVersionProfile();
extern void* C_ZN21QOpenGLVersionProfileC2Ev(); // 3
  // proto:  void QOpenGLVersionProfile::QOpenGLVersionProfile(const QOpenGLVersionProfile & other);
extern void* C_ZN21QOpenGLVersionProfileC2ERKS_(void* arg0); // 3
  // proto:  bool QOpenGLVersionProfile::hasProfiles();
extern bool C_ZNK21QOpenGLVersionProfile11hasProfilesEv(void* qthis); // 4
  // proto:  bool QOpenGLVersionProfile::isValid();
extern bool C_ZNK21QOpenGLVersionProfile7isValidEv(void* qthis); // 4
  // proto:  void QOpenGLVersionProfile::~QOpenGLVersionProfile();
extern void C_ZN21QOpenGLVersionProfileD2Ev(void* qthis); // 4
  // proto:  QPair<int, int> QOpenGLVersionProfile::version();
extern void C_ZNK21QOpenGLVersionProfile7versionEv(void* qthis); // 4
  // proto:  bool QOpenGLVersionProfile::isLegacyVersion();
extern bool C_ZNK21QOpenGLVersionProfile15isLegacyVersionEv(void* qthis); // 4
  // proto:  void QOpenGLVersionProfile::setVersion(int majorVersion, int minorVersion);
extern void C_ZN21QOpenGLVersionProfile10setVersionEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QOpenGLContext::~QOpenGLContext();
extern void C_ZN14QOpenGLContextD2Ev(void* qthis); // 4
  // proto: static void * QOpenGLContext::openGLModuleHandle();
extern void* C_ZN14QOpenGLContext18openGLModuleHandleEv(); // 4
  // proto:  QOpenGLContext * QOpenGLContext::shareContext();
extern void C_ZNK14QOpenGLContext12shareContextEv(void* qthis); // 4
  // proto:  void QOpenGLContext::doneCurrent();
extern void C_ZN14QOpenGLContext11doneCurrentEv(void* qthis); // 4
  // proto:  QSurface * QOpenGLContext::surface();
extern void* C_ZNK14QOpenGLContext7surfaceEv(void* qthis); // 4
  // proto:  void QOpenGLContext::setNativeHandle(const QVariant & handle);
extern void C_ZN14QOpenGLContext15setNativeHandleERK8QVariant(void* qthis, void* arg0); // 4
  // proto: static bool QOpenGLContext::supportsThreadedOpenGL();
extern bool C_ZN14QOpenGLContext22supportsThreadedOpenGLEv(); // 4
  // proto:  QOpenGLFunctions * QOpenGLContext::functions();
extern void C_ZNK14QOpenGLContext9functionsEv(void* qthis); // 4
  // proto: static bool QOpenGLContext::areSharing(QOpenGLContext * first, QOpenGLContext * second);
extern bool C_ZN14QOpenGLContext10areSharingEPS_S0_(void* arg0, void* arg1); // 4
  // proto:  void QOpenGLContext::setScreen(QScreen * screen);
extern void C_ZN14QOpenGLContext9setScreenEP7QScreen(void* qthis, void* arg0); // 4
  // proto: static QOpenGLContext * QOpenGLContext::currentContext();
extern void C_ZN14QOpenGLContext14currentContextEv(); // 4
  // proto:  void QOpenGLContext::setShareContext(QOpenGLContext * shareContext);
extern void C_ZN14QOpenGLContext15setShareContextEPS_(void* qthis, void* arg0); // 4
  // proto: static QOpenGLContext::OpenGLModuleType QOpenGLContext::openGLModuleType();
extern void C_ZN14QOpenGLContext16openGLModuleTypeEv(); // 4
  // proto: static QOpenGLContext * QOpenGLContext::globalShareContext();
extern void C_ZN14QOpenGLContext18globalShareContextEv(); // 4
  // proto:  bool QOpenGLContext::isOpenGLES();
extern bool C_ZNK14QOpenGLContext10isOpenGLESEv(void* qthis); // 4
  // proto:  QAbstractOpenGLFunctions * QOpenGLContext::versionFunctions(const QOpenGLVersionProfile & versionProfile);
extern void C_ZNK14QOpenGLContext16versionFunctionsERK21QOpenGLVersionProfile(void* qthis, void* arg0); // 4
  // proto:  QScreen * QOpenGLContext::screen();
extern void* C_ZNK14QOpenGLContext6screenEv(void* qthis); // 4
  // proto:  void QOpenGLContext::swapBuffers(QSurface * surface);
extern void C_ZN14QOpenGLContext11swapBuffersEP8QSurface(void* qthis, void* arg0); // 4
  // proto:  QPlatformOpenGLContext * QOpenGLContext::handle();
extern void C_ZNK14QOpenGLContext6handleEv(void* qthis); // 4
  // proto:  bool QOpenGLContext::isValid();
extern bool C_ZNK14QOpenGLContext7isValidEv(void* qthis); // 4
  // proto:  bool QOpenGLContext::hasExtension(const QByteArray & extension);
extern bool C_ZNK14QOpenGLContext12hasExtensionERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  GLuint QOpenGLContext::defaultFramebufferObject();
extern int32_t C_ZNK14QOpenGLContext24defaultFramebufferObjectEv(void* qthis); // 4
  // proto:  QFunctionPointer QOpenGLContext::getProcAddress(const QByteArray & procName);
extern void C_ZNK14QOpenGLContext14getProcAddressERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QOpenGLContextGroup * QOpenGLContext::shareGroup();
extern void C_ZNK14QOpenGLContext10shareGroupEv(void* qthis); // 4
  // proto:  QPlatformOpenGLContext * QOpenGLContext::shareHandle();
extern void C_ZNK14QOpenGLContext11shareHandleEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLContext::metaObject();
extern void C_ZNK14QOpenGLContext10metaObjectEv(void* qthis); // 4
  // proto:  void QOpenGLContext::QOpenGLContext(QObject * parent);
extern void* C_ZN14QOpenGLContextC2EP7QObject(void* arg0); // 3
  // proto:  void QOpenGLContext::setFormat(const QSurfaceFormat & format);
extern void C_ZN14QOpenGLContext9setFormatERK14QSurfaceFormat(void* qthis, void* arg0); // 4
  // proto:  bool QOpenGLContext::makeCurrent(QSurface * surface);
extern bool C_ZN14QOpenGLContext11makeCurrentEP8QSurface(void* qthis, void* arg0); // 4
  // proto:  QVariant QOpenGLContext::nativeHandle();
extern void* C_ZNK14QOpenGLContext12nativeHandleEv(void* qthis); // 4
  // proto:  QSet<QByteArray> QOpenGLContext::extensions();
extern void C_ZNK14QOpenGLContext10extensionsEv(void* qthis); // 4
  // proto:  QSurfaceFormat QOpenGLContext::format();
extern void* C_ZNK14QOpenGLContext6formatEv(void* qthis); // 4
  // proto:  bool QOpenGLContext::create();
extern bool C_ZN14QOpenGLContext6createEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLContextGroup::metaObject();
extern void C_ZNK19QOpenGLContextGroup10metaObjectEv(void* qthis); // 4
  // proto: static QOpenGLContextGroup * QOpenGLContextGroup::currentContextGroup();
extern void C_ZN19QOpenGLContextGroup19currentContextGroupEv(); // 4
  // proto:  QList<QOpenGLContext *> QOpenGLContextGroup::shares();
extern void C_ZNK19QOpenGLContextGroup6sharesEv(void* qthis); // 4
  // proto:  void QOpenGLContextGroup::~QOpenGLContextGroup();
extern void C_ZN19QOpenGLContextGroupD2Ev(void* qthis); // 4
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
func (this *QOpenGLVersionProfile) Profile(args ...interface{}) () {
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
    C.C_ZNK21QOpenGLVersionProfile7profileEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "profile", args)
  }

  return
}

// QOpenGLVersionProfile(const class QSurfaceFormat &)
func NewQOpenGLVersionProfile(args ...interface{}) *QOpenGLVersionProfile {
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
    qthis = C.C_ZN21QOpenGLVersionProfileC2ERK14QSurfaceFormat(arg0)
    return &QOpenGLVersionProfile{qclsinst:qthis}
  case 1:
    // invoke: _ZN21QOpenGLVersionProfileC1Ev
    // invoke: void QOpenGLVersionProfile()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QOpenGLVersionProfileC2Ev()
    return &QOpenGLVersionProfile{qclsinst:qthis}
  case 2:
    // invoke: _ZN21QOpenGLVersionProfileC1ERKS_
    // invoke: void QOpenGLVersionProfile(const class QOpenGLVersionProfile &)
    var arg0 = args[0].(QOpenGLVersionProfile).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QOpenGLVersionProfileC2ERKS_(arg0)
    return &QOpenGLVersionProfile{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "QOpenGLVersionProfile", args)
  }

  return nil // QOpenGLVersionProfile{qclsinst:qthis}
}

// hasProfiles()
func (this *QOpenGLVersionProfile) Hasprofiles(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QOpenGLVersionProfile11hasProfilesEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "hasProfiles", args)
  }

  return
}

// isValid()
func (this *QOpenGLVersionProfile) Isvalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QOpenGLVersionProfile7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "isValid", args)
  }

  return
}

// ~QOpenGLVersionProfile()
func (this *QOpenGLVersionProfile) Freeqopenglversionprofile(args ...interface{}) () {
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
    C.C_ZN21QOpenGLVersionProfileD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "~QOpenGLVersionProfile", args)
  }

  return
}

// version()
func (this *QOpenGLVersionProfile) Version(args ...interface{}) () {
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
    C.C_ZNK21QOpenGLVersionProfile7versionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "version", args)
  }

  return
}

// isLegacyVersion()
func (this *QOpenGLVersionProfile) Islegacyversion(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK21QOpenGLVersionProfile15isLegacyVersionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "isLegacyVersion", args)
  }

  return
}

// setVersion(int, int)
func (this *QOpenGLVersionProfile) Setversion(args ...interface{}) () {
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
    C.C_ZN21QOpenGLVersionProfile10setVersionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "setVersion", args)
  }

  return
}

// ~QOpenGLContext()
func (this *QOpenGLContext) Freeqopenglcontext(args ...interface{}) () {
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
    C.C_ZN14QOpenGLContextD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "~QOpenGLContext", args)
  }

  return
}

// openGLModuleHandle()
func (this *QOpenGLContext) Openglmodulehandle_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN14QOpenGLContext18openGLModuleHandleEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.VoidpTy() // "void *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "openGLModuleHandle", args)
  }

  return
}

// shareContext()
func (this *QOpenGLContext) Sharecontext(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLContext12shareContextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "shareContext", args)
  }

  return
}

// doneCurrent()
func (this *QOpenGLContext) Donecurrent(args ...interface{}) () {
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
    C.C_ZN14QOpenGLContext11doneCurrentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "doneCurrent", args)
  }

  return
}

// surface()
func (this *QOpenGLContext) Surface(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLContext7surfaceEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSurface{}) // "QSurface *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "surface", args)
  }

  return
}

// setNativeHandle(const class QVariant &)
func (this *QOpenGLContext) Setnativehandle(args ...interface{}) () {
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
    C.C_ZN14QOpenGLContext15setNativeHandleERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setNativeHandle", args)
  }

  return
}

// supportsThreadedOpenGL()
func (this *QOpenGLContext) Supportsthreadedopengl_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN14QOpenGLContext22supportsThreadedOpenGLEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "supportsThreadedOpenGL", args)
  }

  return
}

// functions()
func (this *QOpenGLContext) Functions(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLContext9functionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "functions", args)
  }

  return
}

// areSharing(class QOpenGLContext *, class QOpenGLContext *)
func (this *QOpenGLContext) Aresharing_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN14QOpenGLContext10areSharingEPS_S0_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "areSharing", args)
  }

  return
}

// setScreen(class QScreen *)
func (this *QOpenGLContext) Setscreen(args ...interface{}) () {
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
    C.C_ZN14QOpenGLContext9setScreenEP7QScreen(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setScreen", args)
  }

  return
}

// currentContext()
func (this *QOpenGLContext) Currentcontext_S(args ...interface{}) () {
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
    C.C_ZN14QOpenGLContext14currentContextEv()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "currentContext", args)
  }

  return
}

// setShareContext(class QOpenGLContext *)
func (this *QOpenGLContext) Setsharecontext(args ...interface{}) () {
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
    C.C_ZN14QOpenGLContext15setShareContextEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setShareContext", args)
  }

  return
}

// openGLModuleType()
func (this *QOpenGLContext) Openglmoduletype_S(args ...interface{}) () {
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
    C.C_ZN14QOpenGLContext16openGLModuleTypeEv()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "openGLModuleType", args)
  }

  return
}

// globalShareContext()
func (this *QOpenGLContext) Globalsharecontext_S(args ...interface{}) () {
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
    C.C_ZN14QOpenGLContext18globalShareContextEv()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "globalShareContext", args)
  }

  return
}

// isOpenGLES()
func (this *QOpenGLContext) Isopengles(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLContext10isOpenGLESEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "isOpenGLES", args)
  }

  return
}

// versionFunctions(const class QOpenGLVersionProfile &)
func (this *QOpenGLContext) Versionfunctions(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLContext16versionFunctionsERK21QOpenGLVersionProfile(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "versionFunctions", args)
  }

  return
}

// screen()
func (this *QOpenGLContext) Screen(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLContext6screenEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QScreen{}) // "QScreen *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "screen", args)
  }

  return
}

// swapBuffers(class QSurface *)
func (this *QOpenGLContext) Swapbuffers(args ...interface{}) () {
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
    C.C_ZN14QOpenGLContext11swapBuffersEP8QSurface(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "swapBuffers", args)
  }

  return
}

// handle()
func (this *QOpenGLContext) Handle(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLContext6handleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "handle", args)
  }

  return
}

// isValid()
func (this *QOpenGLContext) Isvalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLContext7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "isValid", args)
  }

  return
}

// hasExtension(const class QByteArray &)
func (this *QOpenGLContext) Hasextension(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLContext12hasExtensionERK10QByteArray(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "hasExtension", args)
  }

  return
}

// defaultFramebufferObject()
func (this *QOpenGLContext) Defaultframebufferobject(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLContext24defaultFramebufferObjectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "defaultFramebufferObject", args)
  }

  return
}

// getProcAddress(const class QByteArray &)
func (this *QOpenGLContext) Getprocaddress(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLContext14getProcAddressERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "getProcAddress", args)
  }

  return
}

// shareGroup()
func (this *QOpenGLContext) Sharegroup(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLContext10shareGroupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "shareGroup", args)
  }

  return
}

// shareHandle()
func (this *QOpenGLContext) Sharehandle(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLContext11shareHandleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "shareHandle", args)
  }

  return
}

// metaObject()
func (this *QOpenGLContext) Metaobject(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLContext10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "metaObject", args)
  }

  return
}

// QOpenGLContext(class QObject *)
func NewQOpenGLContext(args ...interface{}) *QOpenGLContext {
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
    qthis = C.C_ZN14QOpenGLContextC2EP7QObject(arg0)
    return &QOpenGLContext{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLContext", "QOpenGLContext", args)
  }

  return nil // QOpenGLContext{qclsinst:qthis}
}

// setFormat(const class QSurfaceFormat &)
func (this *QOpenGLContext) Setformat(args ...interface{}) () {
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
    C.C_ZN14QOpenGLContext9setFormatERK14QSurfaceFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setFormat", args)
  }

  return
}

// makeCurrent(class QSurface *)
func (this *QOpenGLContext) Makecurrent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN14QOpenGLContext11makeCurrentEP8QSurface(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "makeCurrent", args)
  }

  return
}

// nativeHandle()
func (this *QOpenGLContext) Nativehandle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLContext12nativeHandleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "nativeHandle", args)
  }

  return
}

// extensions()
func (this *QOpenGLContext) Extensions(args ...interface{}) () {
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
    C.C_ZNK14QOpenGLContext10extensionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContext", "extensions", args)
  }

  return
}

// format()
func (this *QOpenGLContext) Format(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK14QOpenGLContext6formatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSurfaceFormat{}) // "QSurfaceFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "format", args)
  }

  return
}

// create()
func (this *QOpenGLContext) Create(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN14QOpenGLContext6createEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOpenGLContext", "create", args)
  }

  return
}

// metaObject()
func (this *QOpenGLContextGroup) Metaobject(args ...interface{}) () {
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
    C.C_ZNK19QOpenGLContextGroup10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "metaObject", args)
  }

  return
}

// currentContextGroup()
func (this *QOpenGLContextGroup) Currentcontextgroup_S(args ...interface{}) () {
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
    C.C_ZN19QOpenGLContextGroup19currentContextGroupEv()
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "currentContextGroup", args)
  }

  return
}

// shares()
func (this *QOpenGLContextGroup) Shares(args ...interface{}) () {
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
    C.C_ZNK19QOpenGLContextGroup6sharesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "shares", args)
  }

  return
}

// ~QOpenGLContextGroup()
func (this *QOpenGLContextGroup) Freeqopenglcontextgroup(args ...interface{}) () {
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
    C.C_ZN19QOpenGLContextGroupD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "~QOpenGLContextGroup", args)
  }

  return
}

// <= body block end

