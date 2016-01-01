package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qopenglfunctions.h
// dst-file: /src/gui/qopenglfunctions.go
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
// class sizeof(QOpenGLFunctionsPrivate)=1152
type QOpenGLFunctionsPrivate struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLFunctions)=8
type QOpenGLFunctions struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQOpenGLFunctionsPrivate(args ...interface{})() {
}


func (this *QOpenGLFunctions) glBindAttribLocation(args ...interface{}) () {
  // glBindAttribLocation(GLuint, GLuint, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions20glBindAttribLocationEjjPKc
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBindAttribLocation", args)
 }

}


func (this *QOpenGLFunctions) glGenFramebuffers(args ...interface{}) () {
  // glGenFramebuffers(GLsizei, GLuint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][1] = qtrt.Int32Ty(true) // "GLuint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions17glGenFramebuffersEiPj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGenFramebuffers", args)
 }

}


func (this *QOpenGLFunctions) glUniform3iv(args ...interface{}) () {
  // glUniform3iv(GLint, GLsizei, const GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.Int32Ty(true) // "const GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glUniform3ivEiiPKi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform3iv", args)
 }

}


func (this *QOpenGLFunctions) glVertexAttrib4fv(args ...interface{}) () {
  // glVertexAttrib4fv(GLuint, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions17glVertexAttrib4fvEjPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib4fv", args)
 }

}


func (this *QOpenGLFunctions) glIsBuffer(args ...interface{}) () {
  // glIsBuffer(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions10glIsBufferEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsBuffer", args)
 }

}


func (this *QOpenGLFunctions) glLineWidth(args ...interface{}) () {
  // glLineWidth(GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glLineWidthEf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glLineWidth", args)
 }

}


func (this *QOpenGLFunctions) glCompressedTexImage2D(args ...interface{}) () {
  // glCompressedTexImage2D(GLenum, GLint, GLenum, GLsizei, GLsizei, GLint, GLsizei, const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][5] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][6] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][7] = qtrt.VoidpTy() // "const void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions22glCompressedTexImage2DEjijiiiiPKv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCompressedTexImage2D", args)
 }

}


func (this *QOpenGLFunctions) glDepthRangef(args ...interface{}) () {
  // glDepthRangef(GLclampf, GLclampf)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "GLclampf"
  vtys[0][1] = qtrt.FloatTy(false) // "GLclampf"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glDepthRangefEff
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDepthRangef", args)
 }

}


func (this *QOpenGLFunctions) glVertexAttrib1fv(args ...interface{}) () {
  // glVertexAttrib1fv(GLuint, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions17glVertexAttrib1fvEjPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib1fv", args)
 }

}


func (this *QOpenGLFunctions) glTexParameteriv(args ...interface{}) () {
  // glTexParameteriv(GLenum, GLenum, const GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(true) // "const GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glTexParameterivEjjPKi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexParameteriv", args)
 }

}


func (this *QOpenGLFunctions) glTexSubImage2D(args ...interface{}) () {
  // glTexSubImage2D(GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][5] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][6] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][7] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][8] = qtrt.VoidpTy() // "const GLvoid *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glTexSubImage2DEjiiiiijjPKv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexSubImage2D", args)
 }

}


func (this *QOpenGLFunctions) glDeleteProgram(args ...interface{}) () {
  // glDeleteProgram(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glDeleteProgramEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteProgram", args)
 }

}


func (this *QOpenGLFunctions) glBlendEquationSeparate(args ...interface{}) () {
  // glBlendEquationSeparate(GLenum, GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions23glBlendEquationSeparateEjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBlendEquationSeparate", args)
 }

}


func (this *QOpenGLFunctions) glStencilMaskSeparate(args ...interface{}) () {
  // glStencilMaskSeparate(GLenum, GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions21glStencilMaskSeparateEjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilMaskSeparate", args)
 }

}


func (this *QOpenGLFunctions) glDrawArrays(args ...interface{}) () {
  // glDrawArrays(GLenum, GLint, GLsizei)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLsizei"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glDrawArraysEjii
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDrawArrays", args)
 }

}


func (this *QOpenGLFunctions) glFinish(args ...interface{}) () {
  // glFinish()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions8glFinishEv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glFinish", args)
 }

}


func (this *QOpenGLFunctions) glGetVertexAttribPointerv(args ...interface{}) () {
  // glGetVertexAttribPointerv(GLuint, GLenum, void **)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.VoidpTy() // "void **"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions25glGetVertexAttribPointervEjjPPv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetVertexAttribPointerv", args)
 }

}


func (this *QOpenGLFunctions) glActiveTexture(args ...interface{}) () {
  // glActiveTexture(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glActiveTextureEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glActiveTexture", args)
 }

}


func (this *QOpenGLFunctions) glFrontFace(args ...interface{}) () {
  // glFrontFace(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glFrontFaceEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glFrontFace", args)
 }

}


func (this *QOpenGLFunctions) glGetTexParameterfv(args ...interface{}) () {
  // glGetTexParameterfv(GLenum, GLenum, GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.FloatTy(true) // "GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions19glGetTexParameterfvEjjPf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetTexParameterfv", args)
 }

}


func (this *QOpenGLFunctions) glPixelStorei(args ...interface{}) () {
  // glPixelStorei(GLenum, GLint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glPixelStoreiEji
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glPixelStorei", args)
 }

}


func (this *QOpenGLFunctions) glCullFace(args ...interface{}) () {
  // glCullFace(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions10glCullFaceEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCullFace", args)
 }

}


func (this *QOpenGLFunctions) glGetShaderiv(args ...interface{}) () {
  // glGetShaderiv(GLuint, GLenum, GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glGetShaderivEjjPi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetShaderiv", args)
 }

}


func (this *QOpenGLFunctions) glUniform4i(args ...interface{}) () {
  // glUniform4i(GLint, GLint, GLint, GLint, GLint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glUniform4iEiiiii
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform4i", args)
 }

}


func (this *QOpenGLFunctions) glReadPixels(args ...interface{}) () {
  // glReadPixels(GLint, GLint, GLsizei, GLsizei, GLenum, GLenum, GLvoid *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][5] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][6] = qtrt.VoidpTy() // "GLvoid *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glReadPixelsEiiiijjPv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glReadPixels", args)
 }

}


func (this *QOpenGLFunctions) glTexParameteri(args ...interface{}) () {
  // glTexParameteri(GLenum, GLenum, GLint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glTexParameteriEjji
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexParameteri", args)
 }

}


func (this *QOpenGLFunctions) glGetVertexAttribiv(args ...interface{}) () {
  // glGetVertexAttribiv(GLuint, GLenum, GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions19glGetVertexAttribivEjjPi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetVertexAttribiv", args)
 }

}


func (this *QOpenGLFunctions) glClearColor(args ...interface{}) () {
  // glClearColor(GLclampf, GLclampf, GLclampf, GLclampf)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "GLclampf"
  vtys[0][1] = qtrt.FloatTy(false) // "GLclampf"
  vtys[0][2] = qtrt.FloatTy(false) // "GLclampf"
  vtys[0][3] = qtrt.FloatTy(false) // "GLclampf"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glClearColorEffff
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glClearColor", args)
 }

}


func (this *QOpenGLFunctions) glClearDepthf(args ...interface{}) () {
  // glClearDepthf(GLclampf)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "GLclampf"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glClearDepthfEf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glClearDepthf", args)
 }

}


func (this *QOpenGLFunctions) glUniform2i(args ...interface{}) () {
  // glUniform2i(GLint, GLint, GLint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glUniform2iEiii
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform2i", args)
 }

}


func (this *QOpenGLFunctions) glGenerateMipmap(args ...interface{}) () {
  // glGenerateMipmap(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glGenerateMipmapEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGenerateMipmap", args)
 }

}


func (this *QOpenGLFunctions) glCompressedTexSubImage2D(args ...interface{}) () {
  // glCompressedTexSubImage2D(GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLsizei, const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][5] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][6] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][7] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][8] = qtrt.VoidpTy() // "const void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions25glCompressedTexSubImage2DEjiiiiijiPKv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCompressedTexSubImage2D", args)
 }

}


func (this *QOpenGLFunctions) glUniform3i(args ...interface{}) () {
  // glUniform3i(GLint, GLint, GLint, GLint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glUniform3iEiiii
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform3i", args)
 }

}


func (this *QOpenGLFunctions) glGenTextures(args ...interface{}) () {
  // glGenTextures(GLsizei, GLuint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][1] = qtrt.Int32Ty(true) // "GLuint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glGenTexturesEiPj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGenTextures", args)
 }

}


func (this *QOpenGLFunctions) glGetShaderPrecisionFormat(args ...interface{}) () {
  // glGetShaderPrecisionFormat(GLenum, GLenum, GLint *, GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLint *"
  vtys[0][3] = qtrt.Int32Ty(true) // "GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions26glGetShaderPrecisionFormatEjjPiS0_
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetShaderPrecisionFormat", args)
 }

}


func (this *QOpenGLFunctions) FreeQOpenGLFunctions(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "~QOpenGLFunctions", args)
 }

}


func (this *QOpenGLFunctions) glUniform4fv(args ...interface{}) () {
  // glUniform4fv(GLint, GLsizei, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glUniform4fvEiiPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform4fv", args)
 }

}


func (this *QOpenGLFunctions) glGetProgramiv(args ...interface{}) () {
  // glGetProgramiv(GLuint, GLenum, GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glGetProgramivEjjPi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetProgramiv", args)
 }

}


func (this *QOpenGLFunctions) glVertexAttrib2fv(args ...interface{}) () {
  // glVertexAttrib2fv(GLuint, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions17glVertexAttrib2fvEjPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib2fv", args)
 }

}


func (this *QOpenGLFunctions) glGetActiveAttrib(args ...interface{}) () {
  // glGetActiveAttrib(GLuint, GLuint, GLsizei, GLsizei *, GLint *, GLenum *, char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][3] = qtrt.Int32Ty(true) // "GLsizei *"
  vtys[0][4] = qtrt.Int32Ty(true) // "GLint *"
  vtys[0][5] = qtrt.Int32Ty(true) // "GLenum *"
  vtys[0][6] = qtrt.ByteTy(true) // "char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions17glGetActiveAttribEjjiPiS0_PjPc
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetActiveAttrib", args)
 }

}


func (this *QOpenGLFunctions) glIsRenderbuffer(args ...interface{}) () {
  // glIsRenderbuffer(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glIsRenderbufferEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsRenderbuffer", args)
 }

}


func (this *QOpenGLFunctions) glCopyTexSubImage2D(args ...interface{}) () {
  // glCopyTexSubImage2D(GLenum, GLint, GLint, GLint, GLint, GLint, GLsizei, GLsizei)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][5] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][6] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][7] = qtrt.Int32Ty(false) // "GLsizei"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions19glCopyTexSubImage2DEjiiiiiii
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCopyTexSubImage2D", args)
 }

}


func (this *QOpenGLFunctions) glShaderSource(args ...interface{}) () {
  // glShaderSource(GLuint, GLsizei, const char **, const GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.ByteTy(true) // "const char **"
  vtys[0][3] = qtrt.Int32Ty(true) // "const GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glShaderSourceEjiPPKcPKi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glShaderSource", args)
 }

}


func (this *QOpenGLFunctions) glGetVertexAttribfv(args ...interface{}) () {
  // glGetVertexAttribfv(GLuint, GLenum, GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.FloatTy(true) // "GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions19glGetVertexAttribfvEjjPf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetVertexAttribfv", args)
 }

}


func (this *QOpenGLFunctions) glDepthFunc(args ...interface{}) () {
  // glDepthFunc(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glDepthFuncEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDepthFunc", args)
 }

}


func (this *QOpenGLFunctions) glTexImage2D(args ...interface{}) () {
  // glTexImage2D(GLenum, GLint, GLint, GLsizei, GLsizei, GLint, GLenum, GLenum, const GLvoid *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][5] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][6] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][7] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][8] = qtrt.VoidpTy() // "const GLvoid *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glTexImage2DEjiiiiijjPKv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexImage2D", args)
 }

}


func (this *QOpenGLFunctions) glDeleteFramebuffers(args ...interface{}) () {
  // glDeleteFramebuffers(GLsizei, const GLuint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][1] = qtrt.Int32Ty(true) // "const GLuint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions20glDeleteFramebuffersEiPKj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteFramebuffers", args)
 }

}


func (this *QOpenGLFunctions) glHint(args ...interface{}) () {
  // glHint(GLenum, GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions6glHintEjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glHint", args)
 }

}


func (this *QOpenGLFunctions) glGetUniformLocation(args ...interface{}) () {
  // glGetUniformLocation(GLuint, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions20glGetUniformLocationEjPKc
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetUniformLocation", args)
 }

}


func (this *QOpenGLFunctions) glIsFramebuffer(args ...interface{}) () {
  // glIsFramebuffer(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glIsFramebufferEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsFramebuffer", args)
 }

}


func (this *QOpenGLFunctions) glUniform1fv(args ...interface{}) () {
  // glUniform1fv(GLint, GLsizei, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glUniform1fvEiiPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform1fv", args)
 }

}


func (this *QOpenGLFunctions) glGetString(args ...interface{}) () {
  // glGetString(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glGetStringEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetString", args)
 }

}


func (this *QOpenGLFunctions) glUniformMatrix2fv(args ...interface{}) () {
  // glUniformMatrix2fv(GLint, GLsizei, GLboolean, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.ByteTy(false) // "GLboolean"
  vtys[0][3] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions18glUniformMatrix2fvEiihPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniformMatrix2fv", args)
 }

}


func NewQOpenGLFunctions(args ...interface{})() {
}


func (this *QOpenGLFunctions) glUniformMatrix3fv(args ...interface{}) () {
  // glUniformMatrix3fv(GLint, GLsizei, GLboolean, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.ByteTy(false) // "GLboolean"
  vtys[0][3] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions18glUniformMatrix3fvEiihPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniformMatrix3fv", args)
 }

}


func (this *QOpenGLFunctions) glBindBuffer(args ...interface{}) () {
  // glBindBuffer(GLenum, GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glBindBufferEjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBindBuffer", args)
 }

}


func (this *QOpenGLFunctions) glUniform2f(args ...interface{}) () {
  // glUniform2f(GLint, GLfloat, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][2] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glUniform2fEiff
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform2f", args)
 }

}


func (this *QOpenGLFunctions) glUniform3fv(args ...interface{}) () {
  // glUniform3fv(GLint, GLsizei, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glUniform3fvEiiPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform3fv", args)
 }

}


func (this *QOpenGLFunctions) glUniform2fv(args ...interface{}) () {
  // glUniform2fv(GLint, GLsizei, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glUniform2fvEiiPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform2fv", args)
 }

}


func (this *QOpenGLFunctions) glGetRenderbufferParameteriv(args ...interface{}) () {
  // glGetRenderbufferParameteriv(GLenum, GLenum, GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions28glGetRenderbufferParameterivEjjPi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetRenderbufferParameteriv", args)
 }

}


func (this *QOpenGLFunctions) glGetBufferParameteriv(args ...interface{}) () {
  // glGetBufferParameteriv(GLenum, GLenum, GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions22glGetBufferParameterivEjjPi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetBufferParameteriv", args)
 }

}


func (this *QOpenGLFunctions) glUniform1iv(args ...interface{}) () {
  // glUniform1iv(GLint, GLsizei, const GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.Int32Ty(true) // "const GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glUniform1ivEiiPKi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform1iv", args)
 }

}


func (this *QOpenGLFunctions) glBlendColor(args ...interface{}) () {
  // glBlendColor(GLclampf, GLclampf, GLclampf, GLclampf)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "GLclampf"
  vtys[0][1] = qtrt.FloatTy(false) // "GLclampf"
  vtys[0][2] = qtrt.FloatTy(false) // "GLclampf"
  vtys[0][3] = qtrt.FloatTy(false) // "GLclampf"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glBlendColorEffff
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBlendColor", args)
 }

}


func (this *QOpenGLFunctions) glDrawElements(args ...interface{}) () {
  // glDrawElements(GLenum, GLsizei, GLenum, const GLvoid *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][3] = qtrt.VoidpTy() // "const GLvoid *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glDrawElementsEjijPKv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDrawElements", args)
 }

}


func (this *QOpenGLFunctions) glBindFramebuffer(args ...interface{}) () {
  // glBindFramebuffer(GLenum, GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions17glBindFramebufferEjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBindFramebuffer", args)
 }

}


func (this *QOpenGLFunctions) glIsProgram(args ...interface{}) () {
  // glIsProgram(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glIsProgramEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsProgram", args)
 }

}


func (this *QOpenGLFunctions) glBlendEquation(args ...interface{}) () {
  // glBlendEquation(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glBlendEquationEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBlendEquation", args)
 }

}


func (this *QOpenGLFunctions) glShaderBinary(args ...interface{}) () {
  // glShaderBinary(GLint, const GLuint *, GLenum, const void *, GLint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(true) // "const GLuint *"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][3] = qtrt.VoidpTy() // "const void *"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glShaderBinaryEiPKjjPKvi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glShaderBinary", args)
 }

}


func (this *QOpenGLFunctions) glGetProgramInfoLog(args ...interface{}) () {
  // glGetProgramInfoLog(GLuint, GLsizei, GLsizei *, char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLsizei *"
  vtys[0][3] = qtrt.ByteTy(true) // "char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions19glGetProgramInfoLogEjiPiPc
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetProgramInfoLog", args)
 }

}


func (this *QOpenGLFunctions) glDeleteBuffers(args ...interface{}) () {
  // glDeleteBuffers(GLsizei, const GLuint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][1] = qtrt.Int32Ty(true) // "const GLuint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glDeleteBuffersEiPKj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteBuffers", args)
 }

}


func (this *QOpenGLFunctions) glScissor(args ...interface{}) () {
  // glScissor(GLint, GLint, GLsizei, GLsizei)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLsizei"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions9glScissorEiiii
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glScissor", args)
 }

}


func (this *QOpenGLFunctions) glGenRenderbuffers(args ...interface{}) () {
  // glGenRenderbuffers(GLsizei, GLuint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][1] = qtrt.Int32Ty(true) // "GLuint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions18glGenRenderbuffersEiPj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGenRenderbuffers", args)
 }

}


func (this *QOpenGLFunctions) glVertexAttrib3f(args ...interface{}) () {
  // glVertexAttrib3f(GLuint, GLfloat, GLfloat, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][3] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glVertexAttrib3fEjfff
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib3f", args)
 }

}


func (this *QOpenGLFunctions) glCreateProgram(args ...interface{}) () {
  // glCreateProgram()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glCreateProgramEv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCreateProgram", args)
 }

}


func (this *QOpenGLFunctions) glUniform4iv(args ...interface{}) () {
  // glUniform4iv(GLint, GLsizei, const GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.Int32Ty(true) // "const GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glUniform4ivEiiPKi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform4iv", args)
 }

}


func (this *QOpenGLFunctions) glEnable(args ...interface{}) () {
  // glEnable(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions8glEnableEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glEnable", args)
 }

}


func (this *QOpenGLFunctions) glBindTexture(args ...interface{}) () {
  // glBindTexture(GLenum, GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glBindTextureEjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBindTexture", args)
 }

}


func (this *QOpenGLFunctions) glTexParameterf(args ...interface{}) () {
  // glTexParameterf(GLenum, GLenum, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glTexParameterfEjjf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexParameterf", args)
 }

}


func (this *QOpenGLFunctions) glViewport(args ...interface{}) () {
  // glViewport(GLint, GLint, GLsizei, GLsizei)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLsizei"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions10glViewportEiiii
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glViewport", args)
 }

}


func (this *QOpenGLFunctions) glSampleCoverage(args ...interface{}) () {
  // glSampleCoverage(GLclampf, GLboolean)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "GLclampf"
  vtys[0][1] = qtrt.ByteTy(false) // "GLboolean"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glSampleCoverageEfh
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glSampleCoverage", args)
 }

}


func (this *QOpenGLFunctions) glFramebufferTexture2D(args ...interface{}) () {
  // glFramebufferTexture2D(GLenum, GLenum, GLenum, GLuint, GLint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions22glFramebufferTexture2DEjjjji
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glFramebufferTexture2D", args)
 }

}


func (this *QOpenGLFunctions) glVertexAttribPointer(args ...interface{}) () {
  // glVertexAttribPointer(GLuint, GLint, GLenum, GLboolean, GLsizei, const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][3] = qtrt.ByteTy(false) // "GLboolean"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][5] = qtrt.VoidpTy() // "const void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions21glVertexAttribPointerEjijhiPKv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttribPointer", args)
 }

}


func (this *QOpenGLFunctions) glPolygonOffset(args ...interface{}) () {
  // glPolygonOffset(GLfloat, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][1] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glPolygonOffsetEff
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glPolygonOffset", args)
 }

}


func (this *QOpenGLFunctions) glCreateShader(args ...interface{}) () {
  // glCreateShader(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glCreateShaderEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCreateShader", args)
 }

}


func (this *QOpenGLFunctions) glGetShaderSource(args ...interface{}) () {
  // glGetShaderSource(GLuint, GLsizei, GLsizei *, char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLsizei *"
  vtys[0][3] = qtrt.ByteTy(true) // "char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions17glGetShaderSourceEjiPiPc
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetShaderSource", args)
 }

}


func (this *QOpenGLFunctions) glIsTexture(args ...interface{}) () {
  // glIsTexture(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glIsTextureEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsTexture", args)
 }

}


func (this *QOpenGLFunctions) glDeleteTextures(args ...interface{}) () {
  // glDeleteTextures(GLsizei, const GLuint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][1] = qtrt.Int32Ty(true) // "const GLuint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glDeleteTexturesEiPKj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteTextures", args)
 }

}


func (this *QOpenGLFunctions) glGetIntegerv(args ...interface{}) () {
  // glGetIntegerv(GLenum, GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(true) // "GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glGetIntegervEjPi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetIntegerv", args)
 }

}


func (this *QOpenGLFunctions) glGetBooleanv(args ...interface{}) () {
  // glGetBooleanv(GLenum, GLboolean *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.ByteTy(true) // "GLboolean *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glGetBooleanvEjPh
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetBooleanv", args)
 }

}


func (this *QOpenGLFunctions) glGetFloatv(args ...interface{}) () {
  // glGetFloatv(GLenum, GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.FloatTy(true) // "GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glGetFloatvEjPf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetFloatv", args)
 }

}


func (this *QOpenGLFunctions) glDeleteRenderbuffers(args ...interface{}) () {
  // glDeleteRenderbuffers(GLsizei, const GLuint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][1] = qtrt.Int32Ty(true) // "const GLuint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions21glDeleteRenderbuffersEiPKj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteRenderbuffers", args)
 }

}


func (this *QOpenGLFunctions) glGetError(args ...interface{}) () {
  // glGetError()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions10glGetErrorEv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetError", args)
 }

}


func (this *QOpenGLFunctions) glDetachShader(args ...interface{}) () {
  // glDetachShader(GLuint, GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glDetachShaderEjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDetachShader", args)
 }

}


func (this *QOpenGLFunctions) glVertexAttrib2f(args ...interface{}) () {
  // glVertexAttrib2f(GLuint, GLfloat, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][2] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glVertexAttrib2fEjff
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib2f", args)
 }

}


func (this *QOpenGLFunctions) glVertexAttrib1f(args ...interface{}) () {
  // glVertexAttrib1f(GLuint, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glVertexAttrib1fEjf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib1f", args)
 }

}


func (this *QOpenGLFunctions) glGenBuffers(args ...interface{}) () {
  // glGenBuffers(GLsizei, GLuint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][1] = qtrt.Int32Ty(true) // "GLuint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glGenBuffersEiPj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGenBuffers", args)
 }

}


func (this *QOpenGLFunctions) glClearStencil(args ...interface{}) () {
  // glClearStencil(GLint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glClearStencilEi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glClearStencil", args)
 }

}


func (this *QOpenGLFunctions) glStencilMask(args ...interface{}) () {
  // glStencilMask(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glStencilMaskEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilMask", args)
 }

}


func (this *QOpenGLFunctions) glGetShaderInfoLog(args ...interface{}) () {
  // glGetShaderInfoLog(GLuint, GLsizei, GLsizei *, char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLsizei *"
  vtys[0][3] = qtrt.ByteTy(true) // "char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions18glGetShaderInfoLogEjiPiPc
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetShaderInfoLog", args)
 }

}


func (this *QOpenGLFunctions) glReleaseShaderCompiler(args ...interface{}) () {
  // glReleaseShaderCompiler()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions23glReleaseShaderCompilerEv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glReleaseShaderCompiler", args)
 }

}


func (this *QOpenGLFunctions) glDepthMask(args ...interface{}) () {
  // glDepthMask(GLboolean)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "GLboolean"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glDepthMaskEh
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDepthMask", args)
 }

}


func (this *QOpenGLFunctions) glGetFramebufferAttachmentParameteriv(args ...interface{}) () {
  // glGetFramebufferAttachmentParameteriv(GLenum, GLenum, GLenum, GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][3] = qtrt.Int32Ty(true) // "GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions37glGetFramebufferAttachmentParameterivEjjjPi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetFramebufferAttachmentParameteriv", args)
 }

}


func (this *QOpenGLFunctions) glUniform1f(args ...interface{}) () {
  // glUniform1f(GLint, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glUniform1fEif
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform1f", args)
 }

}


func (this *QOpenGLFunctions) glGetAttachedShaders(args ...interface{}) () {
  // glGetAttachedShaders(GLuint, GLsizei, GLsizei *, GLuint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLsizei *"
  vtys[0][3] = qtrt.Int32Ty(true) // "GLuint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions20glGetAttachedShadersEjiPiPj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetAttachedShaders", args)
 }

}


func (this *QOpenGLFunctions) glStencilOp(args ...interface{}) () {
  // glStencilOp(GLenum, GLenum, GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glStencilOpEjjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilOp", args)
 }

}


func (this *QOpenGLFunctions) glStencilFunc(args ...interface{}) () {
  // glStencilFunc(GLenum, GLint, GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glStencilFuncEjij
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilFunc", args)
 }

}


func (this *QOpenGLFunctions) glAttachShader(args ...interface{}) () {
  // glAttachShader(GLuint, GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glAttachShaderEjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glAttachShader", args)
 }

}


func (this *QOpenGLFunctions) glDeleteShader(args ...interface{}) () {
  // glDeleteShader(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glDeleteShaderEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteShader", args)
 }

}


func (this *QOpenGLFunctions) glCompileShader(args ...interface{}) () {
  // glCompileShader(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glCompileShaderEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCompileShader", args)
 }

}


func (this *QOpenGLFunctions) glEnableVertexAttribArray(args ...interface{}) () {
  // glEnableVertexAttribArray(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions25glEnableVertexAttribArrayEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glEnableVertexAttribArray", args)
 }

}


func (this *QOpenGLFunctions) glFramebufferRenderbuffer(args ...interface{}) () {
  // glFramebufferRenderbuffer(GLenum, GLenum, GLenum, GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions25glFramebufferRenderbufferEjjjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glFramebufferRenderbuffer", args)
 }

}


func (this *QOpenGLFunctions) glColorMask(args ...interface{}) () {
  // glColorMask(GLboolean, GLboolean, GLboolean, GLboolean)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "GLboolean"
  vtys[0][1] = qtrt.ByteTy(false) // "GLboolean"
  vtys[0][2] = qtrt.ByteTy(false) // "GLboolean"
  vtys[0][3] = qtrt.ByteTy(false) // "GLboolean"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glColorMaskEhhhh
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glColorMask", args)
 }

}


func (this *QOpenGLFunctions) glIsEnabled(args ...interface{}) () {
  // glIsEnabled(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glIsEnabledEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsEnabled", args)
 }

}


func (this *QOpenGLFunctions) glBindRenderbuffer(args ...interface{}) () {
  // glBindRenderbuffer(GLenum, GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions18glBindRenderbufferEjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBindRenderbuffer", args)
 }

}


func (this *QOpenGLFunctions) glVertexAttrib3fv(args ...interface{}) () {
  // glVertexAttrib3fv(GLuint, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions17glVertexAttrib3fvEjPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib3fv", args)
 }

}


func (this *QOpenGLFunctions) glBlendFunc(args ...interface{}) () {
  // glBlendFunc(GLenum, GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glBlendFuncEjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBlendFunc", args)
 }

}


func (this *QOpenGLFunctions) glUniform3f(args ...interface{}) () {
  // glUniform3f(GLint, GLfloat, GLfloat, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][3] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glUniform3fEifff
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform3f", args)
 }

}


func (this *QOpenGLFunctions) glVertexAttrib4f(args ...interface{}) () {
  // glVertexAttrib4f(GLuint, GLfloat, GLfloat, GLfloat, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][4] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glVertexAttrib4fEjffff
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib4f", args)
 }

}


func (this *QOpenGLFunctions) glGetAttribLocation(args ...interface{}) () {
  // glGetAttribLocation(GLuint, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions19glGetAttribLocationEjPKc
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetAttribLocation", args)
 }

}


func (this *QOpenGLFunctions) glUniform2iv(args ...interface{}) () {
  // glUniform2iv(GLint, GLsizei, const GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.Int32Ty(true) // "const GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glUniform2ivEiiPKi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform2iv", args)
 }

}


func (this *QOpenGLFunctions) glGetUniformiv(args ...interface{}) () {
  // glGetUniformiv(GLuint, GLint, GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glGetUniformivEjiPi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetUniformiv", args)
 }

}


func (this *QOpenGLFunctions) glBufferSubData(args ...interface{}) () {
  // glBufferSubData(GLenum, qopengl_GLintptr, qopengl_GLsizeiptr, const void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "qopengl_GLintptr"
  vtys[0][2] = qtrt.Int32Ty(false) // "qopengl_GLsizeiptr"
  vtys[0][3] = qtrt.VoidpTy() // "const void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions15glBufferSubDataEjiiPKv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBufferSubData", args)
 }

}


func (this *QOpenGLFunctions) glUseProgram(args ...interface{}) () {
  // glUseProgram(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glUseProgramEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUseProgram", args)
 }

}


func (this *QOpenGLFunctions) glDisable(args ...interface{}) () {
  // glDisable(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions9glDisableEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDisable", args)
 }

}


func (this *QOpenGLFunctions) glUniform4f(args ...interface{}) () {
  // glUniform4f(GLint, GLfloat, GLfloat, GLfloat, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][4] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glUniform4fEiffff
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform4f", args)
 }

}


func (this *QOpenGLFunctions) glStencilFuncSeparate(args ...interface{}) () {
  // glStencilFuncSeparate(GLenum, GLenum, GLint, GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions21glStencilFuncSeparateEjjij
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilFuncSeparate", args)
 }

}


func (this *QOpenGLFunctions) glCopyTexImage2D(args ...interface{}) () {
  // glCopyTexImage2D(GLenum, GLint, GLenum, GLint, GLint, GLsizei, GLsizei, GLint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][4] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][5] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][6] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][7] = qtrt.Int32Ty(false) // "GLint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glCopyTexImage2DEjijiiiii
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCopyTexImage2D", args)
 }

}


func (this *QOpenGLFunctions) glLinkProgram(args ...interface{}) () {
  // glLinkProgram(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions13glLinkProgramEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glLinkProgram", args)
 }

}


func (this *QOpenGLFunctions) glBufferData(args ...interface{}) () {
  // glBufferData(GLenum, qopengl_GLsizeiptr, const void *, GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "qopengl_GLsizeiptr"
  vtys[0][2] = qtrt.VoidpTy() // "const void *"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions12glBufferDataEjiPKvj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBufferData", args)
 }

}


func (this *QOpenGLFunctions) glGetUniformfv(args ...interface{}) () {
  // glGetUniformfv(GLuint, GLint, GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][2] = qtrt.FloatTy(true) // "GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions14glGetUniformfvEjiPf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetUniformfv", args)
 }

}


func (this *QOpenGLFunctions) glRenderbufferStorage(args ...interface{}) () {
  // glRenderbufferStorage(GLenum, GLenum, GLsizei, GLsizei)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLsizei"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions21glRenderbufferStorageEjjii
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glRenderbufferStorage", args)
 }

}


func (this *QOpenGLFunctions) glIsShader(args ...interface{}) () {
  // glIsShader(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions10glIsShaderEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsShader", args)
 }

}


func (this *QOpenGLFunctions) initializeOpenGLFunctions(args ...interface{}) () {
  // initializeOpenGLFunctions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions25initializeOpenGLFunctionsEv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "initializeOpenGLFunctions", args)
 }

}


func (this *QOpenGLFunctions) glUniform1i(args ...interface{}) () {
  // glUniform1i(GLint, GLint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions11glUniform1iEii
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform1i", args)
 }

}


func (this *QOpenGLFunctions) glBlendFuncSeparate(args ...interface{}) () {
  // glBlendFuncSeparate(GLenum, GLenum, GLenum, GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions19glBlendFuncSeparateEjjjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBlendFuncSeparate", args)
 }

}


func (this *QOpenGLFunctions) glTexParameterfv(args ...interface{}) () {
  // glTexParameterfv(GLenum, GLenum, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions16glTexParameterfvEjjPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexParameterfv", args)
 }

}


func (this *QOpenGLFunctions) glUniformMatrix4fv(args ...interface{}) () {
  // glUniformMatrix4fv(GLint, GLsizei, GLboolean, const GLfloat *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][2] = qtrt.ByteTy(false) // "GLboolean"
  vtys[0][3] = qtrt.FloatTy(true) // "const GLfloat *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions18glUniformMatrix4fvEiihPKf
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniformMatrix4fv", args)
 }

}


func (this *QOpenGLFunctions) glValidateProgram(args ...interface{}) () {
  // glValidateProgram(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions17glValidateProgramEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glValidateProgram", args)
 }

}


func (this *QOpenGLFunctions) glFlush(args ...interface{}) () {
  // glFlush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions7glFlushEv
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glFlush", args)
 }

}


func (this *QOpenGLFunctions) glCheckFramebufferStatus(args ...interface{}) () {
  // glCheckFramebufferStatus(GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions24glCheckFramebufferStatusEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCheckFramebufferStatus", args)
 }

}


func (this *QOpenGLFunctions) glStencilOpSeparate(args ...interface{}) () {
  // glStencilOpSeparate(GLenum, GLenum, GLenum, GLenum)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][3] = qtrt.Int32Ty(false) // "GLenum"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions19glStencilOpSeparateEjjjj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilOpSeparate", args)
 }

}


func (this *QOpenGLFunctions) glGetTexParameteriv(args ...interface{}) () {
  // glGetTexParameteriv(GLenum, GLenum, GLint *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(true) // "GLint *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions19glGetTexParameterivEjjPi
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetTexParameteriv", args)
 }

}


func (this *QOpenGLFunctions) glClear(args ...interface{}) () {
  // glClear(GLbitfield)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLbitfield"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions7glClearEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glClear", args)
 }

}


func (this *QOpenGLFunctions) glGetActiveUniform(args ...interface{}) () {
  // glGetActiveUniform(GLuint, GLuint, GLsizei, GLsizei *, GLint *, GLenum *, char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLuint"
  vtys[0][2] = qtrt.Int32Ty(false) // "GLsizei"
  vtys[0][3] = qtrt.Int32Ty(true) // "GLsizei *"
  vtys[0][4] = qtrt.Int32Ty(true) // "GLint *"
  vtys[0][5] = qtrt.Int32Ty(true) // "GLenum *"
  vtys[0][6] = qtrt.ByteTy(true) // "char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions18glGetActiveUniformEjjiPiS0_PjPc
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetActiveUniform", args)
 }

}


func (this *QOpenGLFunctions) glDisableVertexAttribArray(args ...interface{}) () {
  // glDisableVertexAttribArray(GLuint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "GLuint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN16QOpenGLFunctions26glDisableVertexAttribArrayEj
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDisableVertexAttribArray", args)
 }

}

// <= body block end

