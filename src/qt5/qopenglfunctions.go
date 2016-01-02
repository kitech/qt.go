package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qopenglfunctions.h
// dst-file: /src/gui/qopenglfunctions.go
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
  // proto:  void QOpenGLFunctionsPrivate::QOpenGLFunctionsPrivate(QOpenGLContext * ctx);
extern void* dector_ZN23QOpenGLFunctionsPrivateC1EP14QOpenGLContext(void* arg0);
extern void _ZN23QOpenGLFunctionsPrivateC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto:  void QOpenGLFunctions::glBindAttribLocation(GLuint program, GLuint index, const char * name);
extern void _ZN16QOpenGLFunctions20glBindAttribLocationEjjPKc(void* qthis, unsigned int arg0, unsigned int arg1, char* arg2);
  // proto:  void QOpenGLFunctions::glGenFramebuffers(GLsizei n, GLuint * framebuffers);
extern void _ZN16QOpenGLFunctions17glGenFramebuffersEiPj(void* qthis, int arg0, unsigned int* arg1);
  // proto:  void QOpenGLFunctions::glUniform3iv(GLint location, GLsizei count, const GLint * v);
extern void _ZN16QOpenGLFunctions12glUniform3ivEiiPKi(void* qthis, int arg0, int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glVertexAttrib4fv(GLuint indx, const GLfloat * values);
extern void _ZN16QOpenGLFunctions17glVertexAttrib4fvEjPKf(void* qthis, unsigned int arg0, float* arg1);
  // proto:  GLboolean QOpenGLFunctions::glIsBuffer(GLuint buffer);
extern void _ZN16QOpenGLFunctions10glIsBufferEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glLineWidth(GLfloat width);
extern void _ZN16QOpenGLFunctions11glLineWidthEf(void* qthis, float arg0);
  // proto:  void QOpenGLFunctions::glCompressedTexImage2D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, const void * data);
extern void _ZN16QOpenGLFunctions22glCompressedTexImage2DEjijiiiiPKv(void* qthis, unsigned int arg0, int arg1, unsigned int arg2, int arg3, int arg4, int arg5, int arg6, void* arg7);
  // proto:  void QOpenGLFunctions::glDepthRangef(GLclampf zNear, GLclampf zFar);
extern void _ZN16QOpenGLFunctions13glDepthRangefEff(void* qthis, float arg0, float arg1);
  // proto:  void QOpenGLFunctions::glVertexAttrib1fv(GLuint indx, const GLfloat * values);
extern void _ZN16QOpenGLFunctions17glVertexAttrib1fvEjPKf(void* qthis, unsigned int arg0, float* arg1);
  // proto:  void QOpenGLFunctions::glTexParameteriv(GLenum target, GLenum pname, const GLint * params);
extern void _ZN16QOpenGLFunctions16glTexParameterivEjjPKi(void* qthis, unsigned int arg0, unsigned int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid * pixels);
extern void _ZN16QOpenGLFunctions15glTexSubImage2DEjiiiiijjPKv(void* qthis, unsigned int arg0, int arg1, int arg2, int arg3, int arg4, int arg5, unsigned int arg6, unsigned int arg7, void* arg8);
  // proto:  void QOpenGLFunctions::glDeleteProgram(GLuint program);
extern void _ZN16QOpenGLFunctions15glDeleteProgramEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glBlendEquationSeparate(GLenum modeRGB, GLenum modeAlpha);
extern void _ZN16QOpenGLFunctions23glBlendEquationSeparateEjj(void* qthis, unsigned int arg0, unsigned int arg1);
  // proto:  void QOpenGLFunctions::glStencilMaskSeparate(GLenum face, GLuint mask);
extern void _ZN16QOpenGLFunctions21glStencilMaskSeparateEjj(void* qthis, unsigned int arg0, unsigned int arg1);
  // proto:  void QOpenGLFunctions::glDrawArrays(GLenum mode, GLint first, GLsizei count);
extern void _ZN16QOpenGLFunctions12glDrawArraysEjii(void* qthis, unsigned int arg0, int arg1, int arg2);
  // proto:  void QOpenGLFunctions::glFinish();
extern void _ZN16QOpenGLFunctions8glFinishEv(void* qthis);
  // proto:  void QOpenGLFunctions::glGetVertexAttribPointerv(GLuint index, GLenum pname, void ** pointer);
extern void _ZN16QOpenGLFunctions25glGetVertexAttribPointervEjjPPv(void* qthis, unsigned int arg0, unsigned int arg1, void* arg2);
  // proto:  void QOpenGLFunctions::glActiveTexture(GLenum texture);
extern void _ZN16QOpenGLFunctions15glActiveTextureEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glFrontFace(GLenum mode);
extern void _ZN16QOpenGLFunctions11glFrontFaceEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glGetTexParameterfv(GLenum target, GLenum pname, GLfloat * params);
extern void _ZN16QOpenGLFunctions19glGetTexParameterfvEjjPf(void* qthis, unsigned int arg0, unsigned int arg1, float* arg2);
  // proto:  void QOpenGLFunctions::glPixelStorei(GLenum pname, GLint param);
extern void _ZN16QOpenGLFunctions13glPixelStoreiEji(void* qthis, unsigned int arg0, int arg1);
  // proto:  void QOpenGLFunctions::glCullFace(GLenum mode);
extern void _ZN16QOpenGLFunctions10glCullFaceEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glGetShaderiv(GLuint shader, GLenum pname, GLint * params);
extern void _ZN16QOpenGLFunctions13glGetShaderivEjjPi(void* qthis, unsigned int arg0, unsigned int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glUniform4i(GLint location, GLint x, GLint y, GLint z, GLint w);
extern void _ZN16QOpenGLFunctions11glUniform4iEiiiii(void* qthis, int arg0, int arg1, int arg2, int arg3, int arg4);
  // proto:  void QOpenGLFunctions::glReadPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid * pixels);
extern void _ZN16QOpenGLFunctions12glReadPixelsEiiiijjPv(void* qthis, int arg0, int arg1, int arg2, int arg3, unsigned int arg4, unsigned int arg5, void* arg6);
  // proto:  void QOpenGLFunctions::glTexParameteri(GLenum target, GLenum pname, GLint param);
extern void _ZN16QOpenGLFunctions15glTexParameteriEjji(void* qthis, unsigned int arg0, unsigned int arg1, int arg2);
  // proto:  void QOpenGLFunctions::glGetVertexAttribiv(GLuint index, GLenum pname, GLint * params);
extern void _ZN16QOpenGLFunctions19glGetVertexAttribivEjjPi(void* qthis, unsigned int arg0, unsigned int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glClearColor(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
extern void _ZN16QOpenGLFunctions12glClearColorEffff(void* qthis, float arg0, float arg1, float arg2, float arg3);
  // proto:  void QOpenGLFunctions::glClearDepthf(GLclampf depth);
extern void _ZN16QOpenGLFunctions13glClearDepthfEf(void* qthis, float arg0);
  // proto:  void QOpenGLFunctions::glUniform2i(GLint location, GLint x, GLint y);
extern void _ZN16QOpenGLFunctions11glUniform2iEiii(void* qthis, int arg0, int arg1, int arg2);
  // proto:  void QOpenGLFunctions::glGenerateMipmap(GLenum target);
extern void _ZN16QOpenGLFunctions16glGenerateMipmapEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glCompressedTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, const void * data);
extern void _ZN16QOpenGLFunctions25glCompressedTexSubImage2DEjiiiiijiPKv(void* qthis, unsigned int arg0, int arg1, int arg2, int arg3, int arg4, int arg5, unsigned int arg6, int arg7, void* arg8);
  // proto:  void QOpenGLFunctions::glUniform3i(GLint location, GLint x, GLint y, GLint z);
extern void _ZN16QOpenGLFunctions11glUniform3iEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QOpenGLFunctions::glGenTextures(GLsizei n, GLuint * textures);
extern void _ZN16QOpenGLFunctions13glGenTexturesEiPj(void* qthis, int arg0, unsigned int* arg1);
  // proto:  void QOpenGLFunctions::glGetShaderPrecisionFormat(GLenum shadertype, GLenum precisiontype, GLint * range, GLint * precision);
extern void _ZN16QOpenGLFunctions26glGetShaderPrecisionFormatEjjPiS0_(void* qthis, unsigned int arg0, unsigned int arg1, int* arg2, int* arg3);
  // proto:  void QOpenGLFunctions::~QOpenGLFunctions();
extern void _ZN16QOpenGLFunctionsD0Ev(void* qthis);
  // proto:  void QOpenGLFunctions::glUniform4fv(GLint location, GLsizei count, const GLfloat * v);
extern void _ZN16QOpenGLFunctions12glUniform4fvEiiPKf(void* qthis, int arg0, int arg1, float* arg2);
  // proto:  void QOpenGLFunctions::glGetProgramiv(GLuint program, GLenum pname, GLint * params);
extern void _ZN16QOpenGLFunctions14glGetProgramivEjjPi(void* qthis, unsigned int arg0, unsigned int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glVertexAttrib2fv(GLuint indx, const GLfloat * values);
extern void _ZN16QOpenGLFunctions17glVertexAttrib2fvEjPKf(void* qthis, unsigned int arg0, float* arg1);
  // proto:  void QOpenGLFunctions::glGetActiveAttrib(GLuint program, GLuint index, GLsizei bufsize, GLsizei * length, GLint * size, GLenum * type, char * name);
extern void _ZN16QOpenGLFunctions17glGetActiveAttribEjjiPiS0_PjPc(void* qthis, unsigned int arg0, unsigned int arg1, int arg2, int* arg3, int* arg4, unsigned int* arg5, char* arg6);
  // proto:  GLboolean QOpenGLFunctions::glIsRenderbuffer(GLuint renderbuffer);
extern void _ZN16QOpenGLFunctions16glIsRenderbufferEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glCopyTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height);
extern void _ZN16QOpenGLFunctions19glCopyTexSubImage2DEjiiiiiii(void* qthis, unsigned int arg0, int arg1, int arg2, int arg3, int arg4, int arg5, int arg6, int arg7);
  // proto:  void QOpenGLFunctions::glShaderSource(GLuint shader, GLsizei count, const char ** string, const GLint * length);
extern void _ZN16QOpenGLFunctions14glShaderSourceEjiPPKcPKi(void* qthis, unsigned int arg0, int arg1, char* arg2, int* arg3);
  // proto:  void QOpenGLFunctions::glGetVertexAttribfv(GLuint index, GLenum pname, GLfloat * params);
extern void _ZN16QOpenGLFunctions19glGetVertexAttribfvEjjPf(void* qthis, unsigned int arg0, unsigned int arg1, float* arg2);
  // proto:  void QOpenGLFunctions::glDepthFunc(GLenum func);
extern void _ZN16QOpenGLFunctions11glDepthFuncEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glTexImage2D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, const GLvoid * pixels);
extern void _ZN16QOpenGLFunctions12glTexImage2DEjiiiiijjPKv(void* qthis, unsigned int arg0, int arg1, int arg2, int arg3, int arg4, int arg5, unsigned int arg6, unsigned int arg7, void* arg8);
  // proto:  void QOpenGLFunctions::glDeleteFramebuffers(GLsizei n, const GLuint * framebuffers);
extern void _ZN16QOpenGLFunctions20glDeleteFramebuffersEiPKj(void* qthis, int arg0, unsigned int* arg1);
  // proto:  void QOpenGLFunctions::glHint(GLenum target, GLenum mode);
extern void _ZN16QOpenGLFunctions6glHintEjj(void* qthis, unsigned int arg0, unsigned int arg1);
  // proto:  GLint QOpenGLFunctions::glGetUniformLocation(GLuint program, const char * name);
extern void _ZN16QOpenGLFunctions20glGetUniformLocationEjPKc(void* qthis, unsigned int arg0, char* arg1);
  // proto:  GLboolean QOpenGLFunctions::glIsFramebuffer(GLuint framebuffer);
extern void _ZN16QOpenGLFunctions15glIsFramebufferEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glUniform1fv(GLint location, GLsizei count, const GLfloat * v);
extern void _ZN16QOpenGLFunctions12glUniform1fvEiiPKf(void* qthis, int arg0, int arg1, float* arg2);
  // proto:  const GLubyte * QOpenGLFunctions::glGetString(GLenum name);
extern void _ZN16QOpenGLFunctions11glGetStringEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glUniformMatrix2fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat * value);
extern void _ZN16QOpenGLFunctions18glUniformMatrix2fvEiihPKf(void* qthis, int arg0, int arg1, unsigned char arg2, float* arg3);
  // proto:  void QOpenGLFunctions::QOpenGLFunctions(QOpenGLContext * context);
extern void* dector_ZN16QOpenGLFunctionsC1EP14QOpenGLContext(void* arg0);
extern void _ZN16QOpenGLFunctionsC1EP14QOpenGLContext(void* qthis, void* arg0);
  // proto:  void QOpenGLFunctions::glUniformMatrix3fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat * value);
extern void _ZN16QOpenGLFunctions18glUniformMatrix3fvEiihPKf(void* qthis, int arg0, int arg1, unsigned char arg2, float* arg3);
  // proto:  void QOpenGLFunctions::glBindBuffer(GLenum target, GLuint buffer);
extern void _ZN16QOpenGLFunctions12glBindBufferEjj(void* qthis, unsigned int arg0, unsigned int arg1);
  // proto:  void QOpenGLFunctions::glUniform2f(GLint location, GLfloat x, GLfloat y);
extern void _ZN16QOpenGLFunctions11glUniform2fEiff(void* qthis, int arg0, float arg1, float arg2);
  // proto:  void QOpenGLFunctions::glUniform3fv(GLint location, GLsizei count, const GLfloat * v);
extern void _ZN16QOpenGLFunctions12glUniform3fvEiiPKf(void* qthis, int arg0, int arg1, float* arg2);
  // proto:  void QOpenGLFunctions::glUniform2fv(GLint location, GLsizei count, const GLfloat * v);
extern void _ZN16QOpenGLFunctions12glUniform2fvEiiPKf(void* qthis, int arg0, int arg1, float* arg2);
  // proto:  void QOpenGLFunctions::glGetRenderbufferParameteriv(GLenum target, GLenum pname, GLint * params);
extern void _ZN16QOpenGLFunctions28glGetRenderbufferParameterivEjjPi(void* qthis, unsigned int arg0, unsigned int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glGetBufferParameteriv(GLenum target, GLenum pname, GLint * params);
extern void _ZN16QOpenGLFunctions22glGetBufferParameterivEjjPi(void* qthis, unsigned int arg0, unsigned int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glUniform1iv(GLint location, GLsizei count, const GLint * v);
extern void _ZN16QOpenGLFunctions12glUniform1ivEiiPKi(void* qthis, int arg0, int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glBlendColor(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
extern void _ZN16QOpenGLFunctions12glBlendColorEffff(void* qthis, float arg0, float arg1, float arg2, float arg3);
  // proto:  void QOpenGLFunctions::glDrawElements(GLenum mode, GLsizei count, GLenum type, const GLvoid * indices);
extern void _ZN16QOpenGLFunctions14glDrawElementsEjijPKv(void* qthis, unsigned int arg0, int arg1, unsigned int arg2, void* arg3);
  // proto:  void QOpenGLFunctions::glBindFramebuffer(GLenum target, GLuint framebuffer);
extern void _ZN16QOpenGLFunctions17glBindFramebufferEjj(void* qthis, unsigned int arg0, unsigned int arg1);
  // proto:  GLboolean QOpenGLFunctions::glIsProgram(GLuint program);
extern void _ZN16QOpenGLFunctions11glIsProgramEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glBlendEquation(GLenum mode);
extern void _ZN16QOpenGLFunctions15glBlendEquationEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glShaderBinary(GLint n, const GLuint * shaders, GLenum binaryformat, const void * binary, GLint length);
extern void _ZN16QOpenGLFunctions14glShaderBinaryEiPKjjPKvi(void* qthis, int arg0, unsigned int* arg1, unsigned int arg2, void* arg3, int arg4);
  // proto:  void QOpenGLFunctions::glGetProgramInfoLog(GLuint program, GLsizei bufsize, GLsizei * length, char * infolog);
extern void _ZN16QOpenGLFunctions19glGetProgramInfoLogEjiPiPc(void* qthis, unsigned int arg0, int arg1, int* arg2, char* arg3);
  // proto:  void QOpenGLFunctions::glDeleteBuffers(GLsizei n, const GLuint * buffers);
extern void _ZN16QOpenGLFunctions15glDeleteBuffersEiPKj(void* qthis, int arg0, unsigned int* arg1);
  // proto:  void QOpenGLFunctions::glScissor(GLint x, GLint y, GLsizei width, GLsizei height);
extern void _ZN16QOpenGLFunctions9glScissorEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QOpenGLFunctions::glGenRenderbuffers(GLsizei n, GLuint * renderbuffers);
extern void _ZN16QOpenGLFunctions18glGenRenderbuffersEiPj(void* qthis, int arg0, unsigned int* arg1);
  // proto:  void QOpenGLFunctions::glVertexAttrib3f(GLuint indx, GLfloat x, GLfloat y, GLfloat z);
extern void _ZN16QOpenGLFunctions16glVertexAttrib3fEjfff(void* qthis, unsigned int arg0, float arg1, float arg2, float arg3);
  // proto:  GLuint QOpenGLFunctions::glCreateProgram();
extern void _ZN16QOpenGLFunctions15glCreateProgramEv(void* qthis);
  // proto:  void QOpenGLFunctions::glUniform4iv(GLint location, GLsizei count, const GLint * v);
extern void _ZN16QOpenGLFunctions12glUniform4ivEiiPKi(void* qthis, int arg0, int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glEnable(GLenum cap);
extern void _ZN16QOpenGLFunctions8glEnableEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glBindTexture(GLenum target, GLuint texture);
extern void _ZN16QOpenGLFunctions13glBindTextureEjj(void* qthis, unsigned int arg0, unsigned int arg1);
  // proto:  void QOpenGLFunctions::glTexParameterf(GLenum target, GLenum pname, GLfloat param);
extern void _ZN16QOpenGLFunctions15glTexParameterfEjjf(void* qthis, unsigned int arg0, unsigned int arg1, float arg2);
  // proto:  void QOpenGLFunctions::glViewport(GLint x, GLint y, GLsizei width, GLsizei height);
extern void _ZN16QOpenGLFunctions10glViewportEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QOpenGLFunctions::glSampleCoverage(GLclampf value, GLboolean invert);
extern void _ZN16QOpenGLFunctions16glSampleCoverageEfh(void* qthis, float arg0, unsigned char arg1);
  // proto:  void QOpenGLFunctions::glFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
extern void _ZN16QOpenGLFunctions22glFramebufferTexture2DEjjjji(void* qthis, unsigned int arg0, unsigned int arg1, unsigned int arg2, unsigned int arg3, int arg4);
  // proto:  void QOpenGLFunctions::glVertexAttribPointer(GLuint indx, GLint size, GLenum type, GLboolean normalized, GLsizei stride, const void * ptr);
extern void _ZN16QOpenGLFunctions21glVertexAttribPointerEjijhiPKv(void* qthis, unsigned int arg0, int arg1, unsigned int arg2, unsigned char arg3, int arg4, void* arg5);
  // proto:  void QOpenGLFunctions::glPolygonOffset(GLfloat factor, GLfloat units);
extern void _ZN16QOpenGLFunctions15glPolygonOffsetEff(void* qthis, float arg0, float arg1);
  // proto:  GLuint QOpenGLFunctions::glCreateShader(GLenum type);
extern void _ZN16QOpenGLFunctions14glCreateShaderEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glGetShaderSource(GLuint shader, GLsizei bufsize, GLsizei * length, char * source);
extern void _ZN16QOpenGLFunctions17glGetShaderSourceEjiPiPc(void* qthis, unsigned int arg0, int arg1, int* arg2, char* arg3);
  // proto:  GLboolean QOpenGLFunctions::glIsTexture(GLuint texture);
extern void _ZN16QOpenGLFunctions11glIsTextureEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glDeleteTextures(GLsizei n, const GLuint * textures);
extern void _ZN16QOpenGLFunctions16glDeleteTexturesEiPKj(void* qthis, int arg0, unsigned int* arg1);
  // proto:  void QOpenGLFunctions::glGetIntegerv(GLenum pname, GLint * params);
extern void _ZN16QOpenGLFunctions13glGetIntegervEjPi(void* qthis, unsigned int arg0, int* arg1);
  // proto:  void QOpenGLFunctions::glGetBooleanv(GLenum pname, GLboolean * params);
extern void _ZN16QOpenGLFunctions13glGetBooleanvEjPh(void* qthis, unsigned int arg0, unsigned char* arg1);
  // proto:  void QOpenGLFunctions::glGetFloatv(GLenum pname, GLfloat * params);
extern void _ZN16QOpenGLFunctions11glGetFloatvEjPf(void* qthis, unsigned int arg0, float* arg1);
  // proto:  void QOpenGLFunctions::glDeleteRenderbuffers(GLsizei n, const GLuint * renderbuffers);
extern void _ZN16QOpenGLFunctions21glDeleteRenderbuffersEiPKj(void* qthis, int arg0, unsigned int* arg1);
  // proto:  GLenum QOpenGLFunctions::glGetError();
extern void _ZN16QOpenGLFunctions10glGetErrorEv(void* qthis);
  // proto:  void QOpenGLFunctions::glDetachShader(GLuint program, GLuint shader);
extern void _ZN16QOpenGLFunctions14glDetachShaderEjj(void* qthis, unsigned int arg0, unsigned int arg1);
  // proto:  void QOpenGLFunctions::glVertexAttrib2f(GLuint indx, GLfloat x, GLfloat y);
extern void _ZN16QOpenGLFunctions16glVertexAttrib2fEjff(void* qthis, unsigned int arg0, float arg1, float arg2);
  // proto:  void QOpenGLFunctions::glVertexAttrib1f(GLuint indx, GLfloat x);
extern void _ZN16QOpenGLFunctions16glVertexAttrib1fEjf(void* qthis, unsigned int arg0, float arg1);
  // proto:  void QOpenGLFunctions::glGenBuffers(GLsizei n, GLuint * buffers);
extern void _ZN16QOpenGLFunctions12glGenBuffersEiPj(void* qthis, int arg0, unsigned int* arg1);
  // proto:  void QOpenGLFunctions::glClearStencil(GLint s);
extern void _ZN16QOpenGLFunctions14glClearStencilEi(void* qthis, int arg0);
  // proto:  void QOpenGLFunctions::glStencilMask(GLuint mask);
extern void _ZN16QOpenGLFunctions13glStencilMaskEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glGetShaderInfoLog(GLuint shader, GLsizei bufsize, GLsizei * length, char * infolog);
extern void _ZN16QOpenGLFunctions18glGetShaderInfoLogEjiPiPc(void* qthis, unsigned int arg0, int arg1, int* arg2, char* arg3);
  // proto:  void QOpenGLFunctions::glReleaseShaderCompiler();
extern void _ZN16QOpenGLFunctions23glReleaseShaderCompilerEv(void* qthis);
  // proto:  void QOpenGLFunctions::glDepthMask(GLboolean flag);
extern void _ZN16QOpenGLFunctions11glDepthMaskEh(void* qthis, unsigned char arg0);
  // proto:  void QOpenGLFunctions::glGetFramebufferAttachmentParameteriv(GLenum target, GLenum attachment, GLenum pname, GLint * params);
extern void _ZN16QOpenGLFunctions37glGetFramebufferAttachmentParameterivEjjjPi(void* qthis, unsigned int arg0, unsigned int arg1, unsigned int arg2, int* arg3);
  // proto:  void QOpenGLFunctions::glUniform1f(GLint location, GLfloat x);
extern void _ZN16QOpenGLFunctions11glUniform1fEif(void* qthis, int arg0, float arg1);
  // proto:  void QOpenGLFunctions::glGetAttachedShaders(GLuint program, GLsizei maxcount, GLsizei * count, GLuint * shaders);
extern void _ZN16QOpenGLFunctions20glGetAttachedShadersEjiPiPj(void* qthis, unsigned int arg0, int arg1, int* arg2, unsigned int* arg3);
  // proto:  void QOpenGLFunctions::glStencilOp(GLenum fail, GLenum zfail, GLenum zpass);
extern void _ZN16QOpenGLFunctions11glStencilOpEjjj(void* qthis, unsigned int arg0, unsigned int arg1, unsigned int arg2);
  // proto:  void QOpenGLFunctions::glStencilFunc(GLenum func, GLint ref, GLuint mask);
extern void _ZN16QOpenGLFunctions13glStencilFuncEjij(void* qthis, unsigned int arg0, int arg1, unsigned int arg2);
  // proto:  void QOpenGLFunctions::glAttachShader(GLuint program, GLuint shader);
extern void _ZN16QOpenGLFunctions14glAttachShaderEjj(void* qthis, unsigned int arg0, unsigned int arg1);
  // proto:  void QOpenGLFunctions::glDeleteShader(GLuint shader);
extern void _ZN16QOpenGLFunctions14glDeleteShaderEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glCompileShader(GLuint shader);
extern void _ZN16QOpenGLFunctions15glCompileShaderEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glEnableVertexAttribArray(GLuint index);
extern void _ZN16QOpenGLFunctions25glEnableVertexAttribArrayEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glFramebufferRenderbuffer(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer);
extern void _ZN16QOpenGLFunctions25glFramebufferRenderbufferEjjjj(void* qthis, unsigned int arg0, unsigned int arg1, unsigned int arg2, unsigned int arg3);
  // proto:  void QOpenGLFunctions::glColorMask(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
extern void _ZN16QOpenGLFunctions11glColorMaskEhhhh(void* qthis, unsigned char arg0, unsigned char arg1, unsigned char arg2, unsigned char arg3);
  // proto:  GLboolean QOpenGLFunctions::glIsEnabled(GLenum cap);
extern void _ZN16QOpenGLFunctions11glIsEnabledEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glBindRenderbuffer(GLenum target, GLuint renderbuffer);
extern void _ZN16QOpenGLFunctions18glBindRenderbufferEjj(void* qthis, unsigned int arg0, unsigned int arg1);
  // proto:  void QOpenGLFunctions::glVertexAttrib3fv(GLuint indx, const GLfloat * values);
extern void _ZN16QOpenGLFunctions17glVertexAttrib3fvEjPKf(void* qthis, unsigned int arg0, float* arg1);
  // proto:  void QOpenGLFunctions::glBlendFunc(GLenum sfactor, GLenum dfactor);
extern void _ZN16QOpenGLFunctions11glBlendFuncEjj(void* qthis, unsigned int arg0, unsigned int arg1);
  // proto:  void QOpenGLFunctions::glUniform3f(GLint location, GLfloat x, GLfloat y, GLfloat z);
extern void _ZN16QOpenGLFunctions11glUniform3fEifff(void* qthis, int arg0, float arg1, float arg2, float arg3);
  // proto:  void QOpenGLFunctions::glVertexAttrib4f(GLuint indx, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
extern void _ZN16QOpenGLFunctions16glVertexAttrib4fEjffff(void* qthis, unsigned int arg0, float arg1, float arg2, float arg3, float arg4);
  // proto:  GLint QOpenGLFunctions::glGetAttribLocation(GLuint program, const char * name);
extern void _ZN16QOpenGLFunctions19glGetAttribLocationEjPKc(void* qthis, unsigned int arg0, char* arg1);
  // proto:  void QOpenGLFunctions::glUniform2iv(GLint location, GLsizei count, const GLint * v);
extern void _ZN16QOpenGLFunctions12glUniform2ivEiiPKi(void* qthis, int arg0, int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glGetUniformiv(GLuint program, GLint location, GLint * params);
extern void _ZN16QOpenGLFunctions14glGetUniformivEjiPi(void* qthis, unsigned int arg0, int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glBufferSubData(GLenum target, qopengl_GLintptr offset, qopengl_GLsizeiptr size, const void * data);
extern void _ZN16QOpenGLFunctions15glBufferSubDataEjiiPKv(void* qthis, unsigned int arg0, int arg1, int arg2, void* arg3);
  // proto:  void QOpenGLFunctions::glUseProgram(GLuint program);
extern void _ZN16QOpenGLFunctions12glUseProgramEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glDisable(GLenum cap);
extern void _ZN16QOpenGLFunctions9glDisableEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glUniform4f(GLint location, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
extern void _ZN16QOpenGLFunctions11glUniform4fEiffff(void* qthis, int arg0, float arg1, float arg2, float arg3, float arg4);
  // proto:  void QOpenGLFunctions::glStencilFuncSeparate(GLenum face, GLenum func, GLint ref, GLuint mask);
extern void _ZN16QOpenGLFunctions21glStencilFuncSeparateEjjij(void* qthis, unsigned int arg0, unsigned int arg1, int arg2, unsigned int arg3);
  // proto:  void QOpenGLFunctions::glCopyTexImage2D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border);
extern void _ZN16QOpenGLFunctions16glCopyTexImage2DEjijiiiii(void* qthis, unsigned int arg0, int arg1, unsigned int arg2, int arg3, int arg4, int arg5, int arg6, int arg7);
  // proto:  void QOpenGLFunctions::glLinkProgram(GLuint program);
extern void _ZN16QOpenGLFunctions13glLinkProgramEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glBufferData(GLenum target, qopengl_GLsizeiptr size, const void * data, GLenum usage);
extern void _ZN16QOpenGLFunctions12glBufferDataEjiPKvj(void* qthis, unsigned int arg0, int arg1, void* arg2, unsigned int arg3);
  // proto:  void QOpenGLFunctions::glGetUniformfv(GLuint program, GLint location, GLfloat * params);
extern void _ZN16QOpenGLFunctions14glGetUniformfvEjiPf(void* qthis, unsigned int arg0, int arg1, float* arg2);
  // proto:  void QOpenGLFunctions::glRenderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height);
extern void _ZN16QOpenGLFunctions21glRenderbufferStorageEjjii(void* qthis, unsigned int arg0, unsigned int arg1, int arg2, int arg3);
  // proto:  GLboolean QOpenGLFunctions::glIsShader(GLuint shader);
extern void _ZN16QOpenGLFunctions10glIsShaderEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::initializeOpenGLFunctions();
extern void _ZN16QOpenGLFunctions25initializeOpenGLFunctionsEv(void* qthis);
  // proto:  void QOpenGLFunctions::glUniform1i(GLint location, GLint x);
extern void _ZN16QOpenGLFunctions11glUniform1iEii(void* qthis, int arg0, int arg1);
  // proto:  void QOpenGLFunctions::glBlendFuncSeparate(GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha);
extern void _ZN16QOpenGLFunctions19glBlendFuncSeparateEjjjj(void* qthis, unsigned int arg0, unsigned int arg1, unsigned int arg2, unsigned int arg3);
  // proto:  void QOpenGLFunctions::glTexParameterfv(GLenum target, GLenum pname, const GLfloat * params);
extern void _ZN16QOpenGLFunctions16glTexParameterfvEjjPKf(void* qthis, unsigned int arg0, unsigned int arg1, float* arg2);
  // proto:  void QOpenGLFunctions::glUniformMatrix4fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat * value);
extern void _ZN16QOpenGLFunctions18glUniformMatrix4fvEiihPKf(void* qthis, int arg0, int arg1, unsigned char arg2, float* arg3);
  // proto:  void QOpenGLFunctions::glValidateProgram(GLuint program);
extern void _ZN16QOpenGLFunctions17glValidateProgramEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::QOpenGLFunctions();
extern void* dector_ZN16QOpenGLFunctionsC1Ev();
extern void _ZN16QOpenGLFunctionsC1Ev(void* qthis);
  // proto:  void QOpenGLFunctions::glFlush();
extern void _ZN16QOpenGLFunctions7glFlushEv(void* qthis);
  // proto:  GLenum QOpenGLFunctions::glCheckFramebufferStatus(GLenum target);
extern void _ZN16QOpenGLFunctions24glCheckFramebufferStatusEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glStencilOpSeparate(GLenum face, GLenum fail, GLenum zfail, GLenum zpass);
extern void _ZN16QOpenGLFunctions19glStencilOpSeparateEjjjj(void* qthis, unsigned int arg0, unsigned int arg1, unsigned int arg2, unsigned int arg3);
  // proto:  void QOpenGLFunctions::glGetTexParameteriv(GLenum target, GLenum pname, GLint * params);
extern void _ZN16QOpenGLFunctions19glGetTexParameterivEjjPi(void* qthis, unsigned int arg0, unsigned int arg1, int* arg2);
  // proto:  void QOpenGLFunctions::glClear(GLbitfield mask);
extern void _ZN16QOpenGLFunctions7glClearEj(void* qthis, unsigned int arg0);
  // proto:  void QOpenGLFunctions::glGetActiveUniform(GLuint program, GLuint index, GLsizei bufsize, GLsizei * length, GLint * size, GLenum * type, char * name);
extern void _ZN16QOpenGLFunctions18glGetActiveUniformEjjiPiS0_PjPc(void* qthis, unsigned int arg0, unsigned int arg1, int arg2, int* arg3, int* arg4, unsigned int* arg5, char* arg6);
  // proto:  void QOpenGLFunctions::glDisableVertexAttribArray(GLuint index);
extern void _ZN16QOpenGLFunctions26glDisableVertexAttribArrayEj(void* qthis, unsigned int arg0);
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

  // proto:  void QOpenGLFunctionsPrivate::QOpenGLFunctionsPrivate(QOpenGLContext * ctx);
func NewQOpenGLFunctionsPrivate(args ...interface{}) QOpenGLFunctionsPrivate {
  return QOpenGLFunctionsPrivate{}
}

  // proto:  void QOpenGLFunctions::glBindAttribLocation(GLuint program, GLuint index, const char * name);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.CString(args[2].(string))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBindAttribLocation", args)
  }

}

  // proto:  void QOpenGLFunctions::glGenFramebuffers(GLsizei n, GLuint * framebuffers);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGenFramebuffers", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform3iv(GLint location, GLsizei count, const GLint * v);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform3iv", args)
  }

}

  // proto:  void QOpenGLFunctions::glVertexAttrib4fv(GLuint indx, const GLfloat * values);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib4fv", args)
  }

}

  // proto:  GLboolean QOpenGLFunctions::glIsBuffer(GLuint buffer);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsBuffer", args)
  }

}

  // proto:  void QOpenGLFunctions::glLineWidth(GLfloat width);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glLineWidth", args)
  }

}

  // proto:  void QOpenGLFunctions::glCompressedTexImage2D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, const void * data);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = args[7].(unsafe.Pointer)
    if false {fmt.Println(arg7)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCompressedTexImage2D", args)
  }

}

  // proto:  void QOpenGLFunctions::glDepthRangef(GLclampf zNear, GLclampf zFar);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDepthRangef", args)
  }

}

  // proto:  void QOpenGLFunctions::glVertexAttrib1fv(GLuint indx, const GLfloat * values);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib1fv", args)
  }

}

  // proto:  void QOpenGLFunctions::glTexParameteriv(GLenum target, GLenum pname, const GLint * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexParameteriv", args)
  }

}

  // proto:  void QOpenGLFunctions::glTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid * pixels);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(args[7].(int32))
    if false {fmt.Println(arg7)}
    var arg8 = args[8].(unsafe.Pointer)
    if false {fmt.Println(arg8)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexSubImage2D", args)
  }

}

  // proto:  void QOpenGLFunctions::glDeleteProgram(GLuint program);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteProgram", args)
  }

}

  // proto:  void QOpenGLFunctions::glBlendEquationSeparate(GLenum modeRGB, GLenum modeAlpha);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBlendEquationSeparate", args)
  }

}

  // proto:  void QOpenGLFunctions::glStencilMaskSeparate(GLenum face, GLuint mask);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilMaskSeparate", args)
  }

}

  // proto:  void QOpenGLFunctions::glDrawArrays(GLenum mode, GLint first, GLsizei count);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDrawArrays", args)
  }

}

  // proto:  void QOpenGLFunctions::glFinish();
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

  // proto:  void QOpenGLFunctions::glGetVertexAttribPointerv(GLuint index, GLenum pname, void ** pointer);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetVertexAttribPointerv", args)
  }

}

  // proto:  void QOpenGLFunctions::glActiveTexture(GLenum texture);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glActiveTexture", args)
  }

}

  // proto:  void QOpenGLFunctions::glFrontFace(GLenum mode);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glFrontFace", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetTexParameterfv(GLenum target, GLenum pname, GLfloat * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.float)(args[2].(*float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetTexParameterfv", args)
  }

}

  // proto:  void QOpenGLFunctions::glPixelStorei(GLenum pname, GLint param);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glPixelStorei", args)
  }

}

  // proto:  void QOpenGLFunctions::glCullFace(GLenum mode);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCullFace", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetShaderiv(GLuint shader, GLenum pname, GLint * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetShaderiv", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform4i(GLint location, GLint x, GLint y, GLint z, GLint w);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform4i", args)
  }

}

  // proto:  void QOpenGLFunctions::glReadPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid * pixels);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(unsafe.Pointer)
    if false {fmt.Println(arg6)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glReadPixels", args)
  }

}

  // proto:  void QOpenGLFunctions::glTexParameteri(GLenum target, GLenum pname, GLint param);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexParameteri", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetVertexAttribiv(GLuint index, GLenum pname, GLint * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetVertexAttribiv", args)
  }

}

  // proto:  void QOpenGLFunctions::glClearColor(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glClearColor", args)
  }

}

  // proto:  void QOpenGLFunctions::glClearDepthf(GLclampf depth);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glClearDepthf", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform2i(GLint location, GLint x, GLint y);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform2i", args)
  }

}

  // proto:  void QOpenGLFunctions::glGenerateMipmap(GLenum target);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGenerateMipmap", args)
  }

}

  // proto:  void QOpenGLFunctions::glCompressedTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, const void * data);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(args[7].(int32))
    if false {fmt.Println(arg7)}
    var arg8 = args[8].(unsafe.Pointer)
    if false {fmt.Println(arg8)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCompressedTexSubImage2D", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform3i(GLint location, GLint x, GLint y, GLint z);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform3i", args)
  }

}

  // proto:  void QOpenGLFunctions::glGenTextures(GLsizei n, GLuint * textures);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGenTextures", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetShaderPrecisionFormat(GLenum shadertype, GLenum precisiontype, GLint * range, GLint * precision);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetShaderPrecisionFormat", args)
  }

}

  // proto:  void QOpenGLFunctions::~QOpenGLFunctions();
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

  // proto:  void QOpenGLFunctions::glUniform4fv(GLint location, GLsizei count, const GLfloat * v);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.float)(args[2].(*float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform4fv", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetProgramiv(GLuint program, GLenum pname, GLint * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetProgramiv", args)
  }

}

  // proto:  void QOpenGLFunctions::glVertexAttrib2fv(GLuint indx, const GLfloat * values);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib2fv", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetActiveAttrib(GLuint program, GLuint index, GLsizei bufsize, GLsizei * length, GLint * size, GLenum * type, char * name);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var arg4 = (*C.int32_t)(args[4].(*int32))
    if false {fmt.Println(arg4)}
    var arg5 = (*C.int32_t)(args[5].(*int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.CString(args[6].(string))
    if false {fmt.Println(arg6)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetActiveAttrib", args)
  }

}

  // proto:  GLboolean QOpenGLFunctions::glIsRenderbuffer(GLuint renderbuffer);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsRenderbuffer", args)
  }

}

  // proto:  void QOpenGLFunctions::glCopyTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(args[7].(int32))
    if false {fmt.Println(arg7)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCopyTexSubImage2D", args)
  }

}

  // proto:  void QOpenGLFunctions::glShaderSource(GLuint shader, GLsizei count, const char ** string, const GLint * length);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.CString(args[2].(string))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glShaderSource", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetVertexAttribfv(GLuint index, GLenum pname, GLfloat * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.float)(args[2].(*float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetVertexAttribfv", args)
  }

}

  // proto:  void QOpenGLFunctions::glDepthFunc(GLenum func);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDepthFunc", args)
  }

}

  // proto:  void QOpenGLFunctions::glTexImage2D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, const GLvoid * pixels);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(args[7].(int32))
    if false {fmt.Println(arg7)}
    var arg8 = args[8].(unsafe.Pointer)
    if false {fmt.Println(arg8)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexImage2D", args)
  }

}

  // proto:  void QOpenGLFunctions::glDeleteFramebuffers(GLsizei n, const GLuint * framebuffers);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteFramebuffers", args)
  }

}

  // proto:  void QOpenGLFunctions::glHint(GLenum target, GLenum mode);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glHint", args)
  }

}

  // proto:  GLint QOpenGLFunctions::glGetUniformLocation(GLuint program, const char * name);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetUniformLocation", args)
  }

}

  // proto:  GLboolean QOpenGLFunctions::glIsFramebuffer(GLuint framebuffer);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsFramebuffer", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform1fv(GLint location, GLsizei count, const GLfloat * v);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.float)(args[2].(*float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform1fv", args)
  }

}

  // proto:  const GLubyte * QOpenGLFunctions::glGetString(GLenum name);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetString", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniformMatrix2fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat * value);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.char(args[2].(byte))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.float)(args[3].(*float32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniformMatrix2fv", args)
  }

}

  // proto:  void QOpenGLFunctions::QOpenGLFunctions(QOpenGLContext * context);
func NewQOpenGLFunctions(args ...interface{}) QOpenGLFunctions {
  return QOpenGLFunctions{}
}

  // proto:  void QOpenGLFunctions::glUniformMatrix3fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat * value);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.char(args[2].(byte))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.float)(args[3].(*float32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniformMatrix3fv", args)
  }

}

  // proto:  void QOpenGLFunctions::glBindBuffer(GLenum target, GLuint buffer);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBindBuffer", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform2f(GLint location, GLfloat x, GLfloat y);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform2f", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform3fv(GLint location, GLsizei count, const GLfloat * v);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.float)(args[2].(*float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform3fv", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform2fv(GLint location, GLsizei count, const GLfloat * v);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.float)(args[2].(*float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform2fv", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetRenderbufferParameteriv(GLenum target, GLenum pname, GLint * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetRenderbufferParameteriv", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetBufferParameteriv(GLenum target, GLenum pname, GLint * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetBufferParameteriv", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform1iv(GLint location, GLsizei count, const GLint * v);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform1iv", args)
  }

}

  // proto:  void QOpenGLFunctions::glBlendColor(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBlendColor", args)
  }

}

  // proto:  void QOpenGLFunctions::glDrawElements(GLenum mode, GLsizei count, GLenum type, const GLvoid * indices);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDrawElements", args)
  }

}

  // proto:  void QOpenGLFunctions::glBindFramebuffer(GLenum target, GLuint framebuffer);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBindFramebuffer", args)
  }

}

  // proto:  GLboolean QOpenGLFunctions::glIsProgram(GLuint program);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsProgram", args)
  }

}

  // proto:  void QOpenGLFunctions::glBlendEquation(GLenum mode);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBlendEquation", args)
  }

}

  // proto:  void QOpenGLFunctions::glShaderBinary(GLint n, const GLuint * shaders, GLenum binaryformat, const void * binary, GLint length);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glShaderBinary", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetProgramInfoLog(GLuint program, GLsizei bufsize, GLsizei * length, char * infolog);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.CString(args[3].(string))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetProgramInfoLog", args)
  }

}

  // proto:  void QOpenGLFunctions::glDeleteBuffers(GLsizei n, const GLuint * buffers);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteBuffers", args)
  }

}

  // proto:  void QOpenGLFunctions::glScissor(GLint x, GLint y, GLsizei width, GLsizei height);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glScissor", args)
  }

}

  // proto:  void QOpenGLFunctions::glGenRenderbuffers(GLsizei n, GLuint * renderbuffers);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGenRenderbuffers", args)
  }

}

  // proto:  void QOpenGLFunctions::glVertexAttrib3f(GLuint indx, GLfloat x, GLfloat y, GLfloat z);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib3f", args)
  }

}

  // proto:  GLuint QOpenGLFunctions::glCreateProgram();
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

  // proto:  void QOpenGLFunctions::glUniform4iv(GLint location, GLsizei count, const GLint * v);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform4iv", args)
  }

}

  // proto:  void QOpenGLFunctions::glEnable(GLenum cap);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glEnable", args)
  }

}

  // proto:  void QOpenGLFunctions::glBindTexture(GLenum target, GLuint texture);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBindTexture", args)
  }

}

  // proto:  void QOpenGLFunctions::glTexParameterf(GLenum target, GLenum pname, GLfloat param);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexParameterf", args)
  }

}

  // proto:  void QOpenGLFunctions::glViewport(GLint x, GLint y, GLsizei width, GLsizei height);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glViewport", args)
  }

}

  // proto:  void QOpenGLFunctions::glSampleCoverage(GLclampf value, GLboolean invert);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.char(args[1].(byte))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glSampleCoverage", args)
  }

}

  // proto:  void QOpenGLFunctions::glFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glFramebufferTexture2D", args)
  }

}

  // proto:  void QOpenGLFunctions::glVertexAttribPointer(GLuint indx, GLint size, GLenum type, GLboolean normalized, GLsizei stride, const void * ptr);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.char(args[3].(byte))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(unsafe.Pointer)
    if false {fmt.Println(arg5)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttribPointer", args)
  }

}

  // proto:  void QOpenGLFunctions::glPolygonOffset(GLfloat factor, GLfloat units);
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
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glPolygonOffset", args)
  }

}

  // proto:  GLuint QOpenGLFunctions::glCreateShader(GLenum type);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCreateShader", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetShaderSource(GLuint shader, GLsizei bufsize, GLsizei * length, char * source);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.CString(args[3].(string))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetShaderSource", args)
  }

}

  // proto:  GLboolean QOpenGLFunctions::glIsTexture(GLuint texture);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsTexture", args)
  }

}

  // proto:  void QOpenGLFunctions::glDeleteTextures(GLsizei n, const GLuint * textures);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteTextures", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetIntegerv(GLenum pname, GLint * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetIntegerv", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetBooleanv(GLenum pname, GLboolean * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetBooleanv", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetFloatv(GLenum pname, GLfloat * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetFloatv", args)
  }

}

  // proto:  void QOpenGLFunctions::glDeleteRenderbuffers(GLsizei n, const GLuint * renderbuffers);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteRenderbuffers", args)
  }

}

  // proto:  GLenum QOpenGLFunctions::glGetError();
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

  // proto:  void QOpenGLFunctions::glDetachShader(GLuint program, GLuint shader);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDetachShader", args)
  }

}

  // proto:  void QOpenGLFunctions::glVertexAttrib2f(GLuint indx, GLfloat x, GLfloat y);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib2f", args)
  }

}

  // proto:  void QOpenGLFunctions::glVertexAttrib1f(GLuint indx, GLfloat x);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib1f", args)
  }

}

  // proto:  void QOpenGLFunctions::glGenBuffers(GLsizei n, GLuint * buffers);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGenBuffers", args)
  }

}

  // proto:  void QOpenGLFunctions::glClearStencil(GLint s);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glClearStencil", args)
  }

}

  // proto:  void QOpenGLFunctions::glStencilMask(GLuint mask);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilMask", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetShaderInfoLog(GLuint shader, GLsizei bufsize, GLsizei * length, char * infolog);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.CString(args[3].(string))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetShaderInfoLog", args)
  }

}

  // proto:  void QOpenGLFunctions::glReleaseShaderCompiler();
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

  // proto:  void QOpenGLFunctions::glDepthMask(GLboolean flag);
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
    var arg0 = C.char(args[0].(byte))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDepthMask", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetFramebufferAttachmentParameteriv(GLenum target, GLenum attachment, GLenum pname, GLint * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetFramebufferAttachmentParameteriv", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform1f(GLint location, GLfloat x);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform1f", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetAttachedShaders(GLuint program, GLsizei maxcount, GLsizei * count, GLuint * shaders);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetAttachedShaders", args)
  }

}

  // proto:  void QOpenGLFunctions::glStencilOp(GLenum fail, GLenum zfail, GLenum zpass);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilOp", args)
  }

}

  // proto:  void QOpenGLFunctions::glStencilFunc(GLenum func, GLint ref, GLuint mask);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilFunc", args)
  }

}

  // proto:  void QOpenGLFunctions::glAttachShader(GLuint program, GLuint shader);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glAttachShader", args)
  }

}

  // proto:  void QOpenGLFunctions::glDeleteShader(GLuint shader);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDeleteShader", args)
  }

}

  // proto:  void QOpenGLFunctions::glCompileShader(GLuint shader);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCompileShader", args)
  }

}

  // proto:  void QOpenGLFunctions::glEnableVertexAttribArray(GLuint index);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glEnableVertexAttribArray", args)
  }

}

  // proto:  void QOpenGLFunctions::glFramebufferRenderbuffer(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glFramebufferRenderbuffer", args)
  }

}

  // proto:  void QOpenGLFunctions::glColorMask(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
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
    var arg0 = C.char(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.char(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.char(args[2].(byte))
    if false {fmt.Println(arg2)}
    var arg3 = C.char(args[3].(byte))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glColorMask", args)
  }

}

  // proto:  GLboolean QOpenGLFunctions::glIsEnabled(GLenum cap);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsEnabled", args)
  }

}

  // proto:  void QOpenGLFunctions::glBindRenderbuffer(GLenum target, GLuint renderbuffer);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBindRenderbuffer", args)
  }

}

  // proto:  void QOpenGLFunctions::glVertexAttrib3fv(GLuint indx, const GLfloat * values);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib3fv", args)
  }

}

  // proto:  void QOpenGLFunctions::glBlendFunc(GLenum sfactor, GLenum dfactor);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBlendFunc", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform3f(GLint location, GLfloat x, GLfloat y, GLfloat z);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform3f", args)
  }

}

  // proto:  void QOpenGLFunctions::glVertexAttrib4f(GLuint indx, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(args[4].(float32))
    if false {fmt.Println(arg4)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glVertexAttrib4f", args)
  }

}

  // proto:  GLint QOpenGLFunctions::glGetAttribLocation(GLuint program, const char * name);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetAttribLocation", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform2iv(GLint location, GLsizei count, const GLint * v);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform2iv", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetUniformiv(GLuint program, GLint location, GLint * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetUniformiv", args)
  }

}

  // proto:  void QOpenGLFunctions::glBufferSubData(GLenum target, qopengl_GLintptr offset, qopengl_GLsizeiptr size, const void * data);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(unsafe.Pointer)
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBufferSubData", args)
  }

}

  // proto:  void QOpenGLFunctions::glUseProgram(GLuint program);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUseProgram", args)
  }

}

  // proto:  void QOpenGLFunctions::glDisable(GLenum cap);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDisable", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniform4f(GLint location, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(args[4].(float32))
    if false {fmt.Println(arg4)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform4f", args)
  }

}

  // proto:  void QOpenGLFunctions::glStencilFuncSeparate(GLenum face, GLenum func, GLint ref, GLuint mask);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilFuncSeparate", args)
  }

}

  // proto:  void QOpenGLFunctions::glCopyTexImage2D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(args[7].(int32))
    if false {fmt.Println(arg7)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCopyTexImage2D", args)
  }

}

  // proto:  void QOpenGLFunctions::glLinkProgram(GLuint program);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glLinkProgram", args)
  }

}

  // proto:  void QOpenGLFunctions::glBufferData(GLenum target, qopengl_GLsizeiptr size, const void * data, GLenum usage);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBufferData", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetUniformfv(GLuint program, GLint location, GLfloat * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.float)(args[2].(*float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetUniformfv", args)
  }

}

  // proto:  void QOpenGLFunctions::glRenderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glRenderbufferStorage", args)
  }

}

  // proto:  GLboolean QOpenGLFunctions::glIsShader(GLuint shader);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glIsShader", args)
  }

}

  // proto:  void QOpenGLFunctions::initializeOpenGLFunctions();
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

  // proto:  void QOpenGLFunctions::glUniform1i(GLint location, GLint x);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniform1i", args)
  }

}

  // proto:  void QOpenGLFunctions::glBlendFuncSeparate(GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glBlendFuncSeparate", args)
  }

}

  // proto:  void QOpenGLFunctions::glTexParameterfv(GLenum target, GLenum pname, const GLfloat * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.float)(args[2].(*float32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glTexParameterfv", args)
  }

}

  // proto:  void QOpenGLFunctions::glUniformMatrix4fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat * value);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.char(args[2].(byte))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.float)(args[3].(*float32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glUniformMatrix4fv", args)
  }

}

  // proto:  void QOpenGLFunctions::glValidateProgram(GLuint program);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glValidateProgram", args)
  }

}

  // proto:  void QOpenGLFunctions::glFlush();
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

  // proto:  GLenum QOpenGLFunctions::glCheckFramebufferStatus(GLenum target);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glCheckFramebufferStatus", args)
  }

}

  // proto:  void QOpenGLFunctions::glStencilOpSeparate(GLenum face, GLenum fail, GLenum zfail, GLenum zpass);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glStencilOpSeparate", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetTexParameteriv(GLenum target, GLenum pname, GLint * params);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetTexParameteriv", args)
  }

}

  // proto:  void QOpenGLFunctions::glClear(GLbitfield mask);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glClear", args)
  }

}

  // proto:  void QOpenGLFunctions::glGetActiveUniform(GLuint program, GLuint index, GLsizei bufsize, GLsizei * length, GLint * size, GLenum * type, char * name);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var arg4 = (*C.int32_t)(args[4].(*int32))
    if false {fmt.Println(arg4)}
    var arg5 = (*C.int32_t)(args[5].(*int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.CString(args[6].(string))
    if false {fmt.Println(arg6)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glGetActiveUniform", args)
  }

}

  // proto:  void QOpenGLFunctions::glDisableVertexAttribArray(GLuint index);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLFunctions", "glDisableVertexAttribArray", args)
  }

}

// <= body block end

