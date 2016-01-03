package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  bool QOpenGLVersionProfile::isLegacyVersion();
extern void _ZNK21QOpenGLVersionProfile15isLegacyVersionEv(void* qthis);
  // proto:  void QOpenGLVersionProfile::~QOpenGLVersionProfile();
extern void _ZN21QOpenGLVersionProfileD0Ev(void* qthis);
  // proto:  bool QOpenGLVersionProfile::hasProfiles();
extern void _ZNK21QOpenGLVersionProfile11hasProfilesEv(void* qthis);
  // proto:  void QOpenGLVersionProfile::QOpenGLVersionProfile(const QSurfaceFormat & format);
extern void* dector_ZN21QOpenGLVersionProfileC1ERK14QSurfaceFormat(void* arg0);
extern void _ZN21QOpenGLVersionProfileC1ERK14QSurfaceFormat(void* qthis, void* arg0);
  // proto:  void QOpenGLVersionProfile::QOpenGLVersionProfile(const QOpenGLVersionProfile & other);
extern void* dector_ZN21QOpenGLVersionProfileC1ERKS_(void* arg0);
extern void _ZN21QOpenGLVersionProfileC1ERKS_(void* qthis, void* arg0);
  // proto:  void QOpenGLVersionProfile::QOpenGLVersionProfile();
extern void* dector_ZN21QOpenGLVersionProfileC1Ev();
extern void _ZN21QOpenGLVersionProfileC1Ev(void* qthis);
  // proto:  QPair<int, int> QOpenGLVersionProfile::version();
extern void _ZNK21QOpenGLVersionProfile7versionEv(void* qthis);
  // proto:  void QOpenGLVersionProfile::setVersion(int majorVersion, int minorVersion);
extern void _ZN21QOpenGLVersionProfile10setVersionEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  bool QOpenGLVersionProfile::isValid();
extern void _ZNK21QOpenGLVersionProfile7isValidEv(void* qthis);
  // proto:  bool QOpenGLContext::isValid();
extern void _ZNK14QOpenGLContext7isValidEv(void* qthis);
  // proto:  void QOpenGLContext::setScreen(QScreen * screen);
extern void _ZN14QOpenGLContext9setScreenEP7QScreen(void* qthis, void* arg0);
  // proto:  void QOpenGLContext::QOpenGLContext(QObject * parent);
extern void* dector_ZN14QOpenGLContextC1EP7QObject(void* arg0);
extern void _ZN14QOpenGLContextC1EP7QObject(void* qthis, void* arg0);
  // proto:  QOpenGLFunctions * QOpenGLContext::functions();
extern void _ZNK14QOpenGLContext9functionsEv(void* qthis);
  // proto:  void QOpenGLContext::~QOpenGLContext();
extern void _ZN14QOpenGLContextD0Ev(void* qthis);
  // proto:  void QOpenGLContext::setFormat(const QSurfaceFormat & format);
extern void _ZN14QOpenGLContext9setFormatERK14QSurfaceFormat(void* qthis, void* arg0);
  // proto:  const QMetaObject * QOpenGLContext::metaObject();
extern void _ZNK14QOpenGLContext10metaObjectEv(void* qthis);
  // proto:  bool QOpenGLContext::hasExtension(const QByteArray & extension);
extern void _ZNK14QOpenGLContext12hasExtensionERK10QByteArray(void* qthis, void* arg0);
  // proto:  QSet<QByteArray> QOpenGLContext::extensions();
extern void _ZNK14QOpenGLContext10extensionsEv(void* qthis);
  // proto:  QSurface * QOpenGLContext::surface();
extern void _ZNK14QOpenGLContext7surfaceEv(void* qthis);
  // proto:  QAbstractOpenGLFunctions * QOpenGLContext::versionFunctions(const QOpenGLVersionProfile & versionProfile);
extern void _ZNK14QOpenGLContext16versionFunctionsERK21QOpenGLVersionProfile(void* qthis, void* arg0);
  // proto:  void QOpenGLContext::setShareContext(QOpenGLContext * shareContext);
extern void _ZN14QOpenGLContext15setShareContextEPS_(void* qthis, void* arg0);
  // proto: static bool QOpenGLContext::areSharing(QOpenGLContext * first, QOpenGLContext * second);
extern void _ZN14QOpenGLContext10areSharingEPS_S0_(void* arg0, void* arg1);
  // proto:  QScreen * QOpenGLContext::screen();
extern void _ZNK14QOpenGLContext6screenEv(void* qthis);
  // proto:  QVariant QOpenGLContext::nativeHandle();
extern void _ZNK14QOpenGLContext12nativeHandleEv(void* qthis);
  // proto:  bool QOpenGLContext::isOpenGLES();
extern void _ZNK14QOpenGLContext10isOpenGLESEv(void* qthis);
  // proto:  QPlatformOpenGLContext * QOpenGLContext::handle();
extern void _ZNK14QOpenGLContext6handleEv(void* qthis);
  // proto: static QOpenGLContext * QOpenGLContext::globalShareContext();
extern void _ZN14QOpenGLContext18globalShareContextEv();
  // proto:  bool QOpenGLContext::makeCurrent(QSurface * surface);
extern void _ZN14QOpenGLContext11makeCurrentEP8QSurface(void* qthis, void* arg0);
  // proto:  QPlatformOpenGLContext * QOpenGLContext::shareHandle();
extern void _ZNK14QOpenGLContext11shareHandleEv(void* qthis);
  // proto:  bool QOpenGLContext::create();
extern void _ZN14QOpenGLContext6createEv(void* qthis);
  // proto:  QOpenGLContext * QOpenGLContext::shareContext();
extern void _ZNK14QOpenGLContext12shareContextEv(void* qthis);
  // proto: static QOpenGLContext * QOpenGLContext::currentContext();
extern void _ZN14QOpenGLContext14currentContextEv();
  // proto:  GLuint QOpenGLContext::defaultFramebufferObject();
extern void _ZNK14QOpenGLContext24defaultFramebufferObjectEv(void* qthis);
  // proto: static bool QOpenGLContext::supportsThreadedOpenGL();
extern void _ZN14QOpenGLContext22supportsThreadedOpenGLEv();
  // proto:  void QOpenGLContext::doneCurrent();
extern void _ZN14QOpenGLContext11doneCurrentEv(void* qthis);
  // proto:  QOpenGLContextGroup * QOpenGLContext::shareGroup();
extern void _ZNK14QOpenGLContext10shareGroupEv(void* qthis);
  // proto:  QSurfaceFormat QOpenGLContext::format();
extern void _ZNK14QOpenGLContext6formatEv(void* qthis);
  // proto: static void * QOpenGLContext::openGLModuleHandle();
extern void _ZN14QOpenGLContext18openGLModuleHandleEv();
  // proto:  void QOpenGLContext::setNativeHandle(const QVariant & handle);
extern void _ZN14QOpenGLContext15setNativeHandleERK8QVariant(void* qthis, void* arg0);
  // proto:  QFunctionPointer QOpenGLContext::getProcAddress(const QByteArray & procName);
extern void _ZNK14QOpenGLContext14getProcAddressERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QOpenGLContext::swapBuffers(QSurface * surface);
extern void _ZN14QOpenGLContext11swapBuffersEP8QSurface(void* qthis, void* arg0);
  // proto:  const QMetaObject * QOpenGLContextGroup::metaObject();
extern void _ZNK19QOpenGLContextGroup10metaObjectEv(void* qthis);
  // proto: static QOpenGLContextGroup * QOpenGLContextGroup::currentContextGroup();
extern void _ZN19QOpenGLContextGroup19currentContextGroupEv();
  // proto:  void QOpenGLContextGroup::~QOpenGLContextGroup();
extern void _ZN19QOpenGLContextGroupD0Ev(void* qthis);
  // proto:  void QOpenGLContextGroup::QOpenGLContextGroup();
extern void* dector_ZN19QOpenGLContextGroupC1Ev();
extern void _ZN19QOpenGLContextGroupC1Ev(void* qthis);
  // proto:  QList<QOpenGLContext *> QOpenGLContextGroup::shares();
extern void _ZNK19QOpenGLContextGroup6sharesEv(void* qthis);
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

  // proto:  bool QOpenGLVersionProfile::isLegacyVersion();
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

  // proto:  void QOpenGLVersionProfile::~QOpenGLVersionProfile();
func (this *QOpenGLVersionProfile) FreeQOpenGLVersionProfile(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "~QOpenGLVersionProfile", args)
  }

}

  // proto:  bool QOpenGLVersionProfile::hasProfiles();
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

  // proto:  void QOpenGLVersionProfile::QOpenGLVersionProfile(const QSurfaceFormat & format);
func NewQOpenGLVersionProfile(args ...interface{}) QOpenGLVersionProfile {
  return QOpenGLVersionProfile{}
}

  // proto:  QPair<int, int> QOpenGLVersionProfile::version();
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

  // proto:  void QOpenGLVersionProfile::setVersion(int majorVersion, int minorVersion);
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

  // proto:  bool QOpenGLVersionProfile::isValid();
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

  // proto:  bool QOpenGLContext::isValid();
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

  // proto:  void QOpenGLContext::setScreen(QScreen * screen);
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

  // proto:  void QOpenGLContext::QOpenGLContext(QObject * parent);
func NewQOpenGLContext(args ...interface{}) QOpenGLContext {
  return QOpenGLContext{}
}

  // proto:  QOpenGLFunctions * QOpenGLContext::functions();
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

  // proto:  void QOpenGLContext::~QOpenGLContext();
func (this *QOpenGLContext) FreeQOpenGLContext(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLContext", "~QOpenGLContext", args)
  }

}

  // proto:  void QOpenGLContext::setFormat(const QSurfaceFormat & format);
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

  // proto:  const QMetaObject * QOpenGLContext::metaObject();
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

  // proto:  bool QOpenGLContext::hasExtension(const QByteArray & extension);
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

  // proto:  QSet<QByteArray> QOpenGLContext::extensions();
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

  // proto:  QSurface * QOpenGLContext::surface();
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

  // proto:  QAbstractOpenGLFunctions * QOpenGLContext::versionFunctions(const QOpenGLVersionProfile & versionProfile);
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

  // proto:  void QOpenGLContext::setShareContext(QOpenGLContext * shareContext);
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

  // proto: static bool QOpenGLContext::areSharing(QOpenGLContext * first, QOpenGLContext * second);
func (this *QOpenGLContext) areSharing_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLContext", "areSharing", args)
  }

}

  // proto:  QScreen * QOpenGLContext::screen();
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

  // proto:  QVariant QOpenGLContext::nativeHandle();
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

  // proto:  bool QOpenGLContext::isOpenGLES();
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

  // proto:  QPlatformOpenGLContext * QOpenGLContext::handle();
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

  // proto: static QOpenGLContext * QOpenGLContext::globalShareContext();
func (this *QOpenGLContext) globalShareContext_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLContext", "globalShareContext", args)
  }

}

  // proto:  bool QOpenGLContext::makeCurrent(QSurface * surface);
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

  // proto:  QPlatformOpenGLContext * QOpenGLContext::shareHandle();
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

  // proto:  bool QOpenGLContext::create();
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

  // proto:  QOpenGLContext * QOpenGLContext::shareContext();
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

  // proto: static QOpenGLContext * QOpenGLContext::currentContext();
func (this *QOpenGLContext) currentContext_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLContext", "currentContext", args)
  }

}

  // proto:  GLuint QOpenGLContext::defaultFramebufferObject();
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

  // proto: static bool QOpenGLContext::supportsThreadedOpenGL();
func (this *QOpenGLContext) supportsThreadedOpenGL_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLContext", "supportsThreadedOpenGL", args)
  }

}

  // proto:  void QOpenGLContext::doneCurrent();
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

  // proto:  QOpenGLContextGroup * QOpenGLContext::shareGroup();
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

  // proto:  QSurfaceFormat QOpenGLContext::format();
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

  // proto: static void * QOpenGLContext::openGLModuleHandle();
func (this *QOpenGLContext) openGLModuleHandle_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLContext", "openGLModuleHandle", args)
  }

}

  // proto:  void QOpenGLContext::setNativeHandle(const QVariant & handle);
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

  // proto:  QFunctionPointer QOpenGLContext::getProcAddress(const QByteArray & procName);
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

  // proto:  void QOpenGLContext::swapBuffers(QSurface * surface);
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

  // proto:  const QMetaObject * QOpenGLContextGroup::metaObject();
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

  // proto: static QOpenGLContextGroup * QOpenGLContextGroup::currentContextGroup();
func (this *QOpenGLContextGroup) currentContextGroup_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "currentContextGroup", args)
  }

}

  // proto:  void QOpenGLContextGroup::~QOpenGLContextGroup();
func (this *QOpenGLContextGroup) FreeQOpenGLContextGroup(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "~QOpenGLContextGroup", args)
  }

}

  // proto:  void QOpenGLContextGroup::QOpenGLContextGroup();
func NewQOpenGLContextGroup(args ...interface{}) QOpenGLContextGroup {
  return QOpenGLContextGroup{}
}

  // proto:  QList<QOpenGLContext *> QOpenGLContextGroup::shares();
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

// <= body block end

