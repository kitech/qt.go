package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qopenglcontext.h
// dst-file: /src/gui/qopenglcontext.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QOpenGLVersionProfile)=8
type QOpenGLVersionProfile struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLContext)=1
type QOpenGLContext struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _aboutToBeDestroyed QOpenGLContext_aboutToBeDestroyed_signal;
}

// class sizeof(QOpenGLContextGroup)=1
type QOpenGLContextGroup struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


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
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "isLegacyVersion", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "hasProfiles", args)
 }

}


func NewQOpenGLVersionProfile(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "version", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "setVersion", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLVersionProfile", "isValid", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "isValid", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setScreen", args)
 }

}


func NewQOpenGLContext(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "functions", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setFormat", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "hasExtension", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "extensions", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "surface", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "versionFunctions", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setShareContext", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "screen", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "nativeHandle", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "isOpenGLES", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "handle", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "makeCurrent", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "shareHandle", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "create", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "shareContext", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "defaultFramebufferObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "doneCurrent", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "shareGroup", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "format", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "setNativeHandle", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "getProcAddress", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContext", "swapBuffers", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "metaObject", args)
 }

}


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


func NewQOpenGLContextGroup(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QOpenGLContextGroup", "shares", args)
 }

}

// <= body block end

