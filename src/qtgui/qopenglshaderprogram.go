package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtGui/qopenglshaderprogram.h
// dst-file: /src/gui/qopenglshaderprogram.go
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QOpenGLShader::isCompiled();
extern bool C_ZNK13QOpenGLShader10isCompiledEv(void* qthis); // 4
  // proto:  QString QOpenGLShader::log();
extern void* C_ZNK13QOpenGLShader3logEv(void* qthis); // 4
  // proto:  bool QOpenGLShader::compileSourceFile(const QString & fileName);
extern bool C_ZN13QOpenGLShader17compileSourceFileERK7QString(void* qthis, void* arg0); // 4
  // proto:  GLuint QOpenGLShader::shaderId();
extern int32_t C_ZNK13QOpenGLShader8shaderIdEv(void* qthis); // 4
  // proto:  bool QOpenGLShader::compileSourceCode(const QString & source);
extern bool C_ZN13QOpenGLShader17compileSourceCodeERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QOpenGLShader::compileSourceCode(const char * source);
extern bool C_ZN13QOpenGLShader17compileSourceCodeEPKc(void* qthis, void* arg0); // 4
  // proto:  bool QOpenGLShader::compileSourceCode(const QByteArray & source);
extern bool C_ZN13QOpenGLShader17compileSourceCodeERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QByteArray QOpenGLShader::sourceCode();
extern void* C_ZNK13QOpenGLShader10sourceCodeEv(void* qthis); // 4
  // proto:  QOpenGLShader::ShaderType QOpenGLShader::shaderType();
extern void C_ZNK13QOpenGLShader10shaderTypeEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLShader::metaObject();
extern void C_ZNK13QOpenGLShader10metaObjectEv(void* qthis); // 4
  // proto:  void QOpenGLShader::~QOpenGLShader();
extern void C_ZN13QOpenGLShaderD2Ev(void* qthis); // 4
  // proto:  void QOpenGLShaderProgram::setPatchVertexCount(int count);
extern void C_ZN20QOpenGLShaderProgram19setPatchVertexCountEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QOpenGLShaderProgram::addShader(QOpenGLShader * shader);
extern bool C_ZN20QOpenGLShaderProgram9addShaderEP13QOpenGLShader(void* qthis, void* arg0); // 4
  // proto:  QVector<float> QOpenGLShaderProgram::defaultOuterTessellationLevels();
extern void C_ZNK20QOpenGLShaderProgram30defaultOuterTessellationLevelsEv(void* qthis); // 4
  // proto:  int QOpenGLShaderProgram::attributeLocation(const char * name);
extern int32_t C_ZNK20QOpenGLShaderProgram17attributeLocationEPKc(void* qthis, void* arg0); // 4
  // proto:  int QOpenGLShaderProgram::attributeLocation(const QByteArray & name);
extern int32_t C_ZNK20QOpenGLShaderProgram17attributeLocationERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  int QOpenGLShaderProgram::attributeLocation(const QString & name);
extern int32_t C_ZNK20QOpenGLShaderProgram17attributeLocationERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, const QVector3D * values, int stride);
extern void C_ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector3Di(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeArray(int location, GLenum type, const void * values, int tupleSize, int stride);
extern void C_ZN20QOpenGLShaderProgram17setAttributeArrayEijPKvii(void* qthis, int32_t arg0, int32_t arg1, void* arg2, int32_t arg3, int32_t arg4); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeArray(int location, const QVector3D * values, int stride);
extern void C_ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector3Di(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeArray(int location, const QVector4D * values, int stride);
extern void C_ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector4Di(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeArray(int location, const QVector2D * values, int stride);
extern void C_ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector2Di(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, const QVector4D * values, int stride);
extern void C_ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector4Di(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeArray(int location, const GLfloat * values, int tupleSize, int stride);
extern void C_ZN20QOpenGLShaderProgram17setAttributeArrayEiPKfii(void* qthis, int32_t arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, const QVector2D * values, int stride);
extern void C_ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector2Di(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, GLenum type, const void * values, int tupleSize, int stride);
extern void C_ZN20QOpenGLShaderProgram17setAttributeArrayEPKcjPKvii(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3, int32_t arg4); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, const GLfloat * values, int tupleSize, int stride);
extern void C_ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPKfii(void* qthis, void* arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QOpenGLShaderProgram::removeShader(QOpenGLShader * shader);
extern void C_ZN20QOpenGLShaderProgram12removeShaderEP13QOpenGLShader(void* qthis, void* arg0); // 4
  // proto:  void QOpenGLShaderProgram::removeAllShaders();
extern void C_ZN20QOpenGLShaderProgram16removeAllShadersEv(void* qthis); // 4
  // proto:  int QOpenGLShaderProgram::patchVertexCount();
extern int32_t C_ZNK20QOpenGLShaderProgram16patchVertexCountEv(void* qthis); // 4
  // proto:  QString QOpenGLShaderProgram::log();
extern void* C_ZNK20QOpenGLShaderProgram3logEv(void* qthis); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QSize & size);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK5QSize(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QColor & color);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiRK6QColor(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const GLfloat [3][3] value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiPA3_Kf(void* qthis, int32_t arg0, void** arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QColor & color);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QColor(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QPoint & point);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiRK6QPoint(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLfloat x, GLfloat y);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiff(void* qthis, int32_t arg0, float arg1, float arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcffff(void* qthis, void* arg0, float arg1, float arg2, float arg3, float arg4); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QPointF & point);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK7QPointF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QTransform & value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiRK10QTransform(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLfloat x, GLfloat y, GLfloat z);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcfff(void* qthis, void* arg0, float arg1, float arg2, float arg3); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const GLfloat [2][2] value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiPA2_Kf(void* qthis, int32_t arg0, void** arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QMatrix4x4 & value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QMatrix4x4(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QVector3D & value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector3D(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLuint value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcj(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLint value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKci(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const GLfloat [2][2] value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcPA2_Kf(void* qthis, void* arg0, void** arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLfloat x, GLfloat y, GLfloat z);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEifff(void* qthis, int32_t arg0, float arg1, float arg2, float arg3); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLfloat value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcf(void* qthis, void* arg0, float arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QVector4D & value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector4D(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QPoint & point);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QPoint(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QSizeF & size);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiRK6QSizeF(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QSizeF & size);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QSizeF(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiffff(void* qthis, int32_t arg0, float arg1, float arg2, float arg3, float arg4); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const GLfloat [3][3] value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcPA3_Kf(void* qthis, void* arg0, void** arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QVector2D & value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector2D(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const GLfloat [4][4] value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcPA4_Kf(void* qthis, void* arg0, void** arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QVector4D & value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector4D(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLfloat x, GLfloat y);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcff(void* qthis, void* arg0, float arg1, float arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QTransform & value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QTransform(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QVector2D & value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector2D(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QPointF & point);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiRK7QPointF(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLfloat value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEif(void* qthis, int32_t arg0, float arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QMatrix4x4 & value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiRK10QMatrix4x4(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QVector3D & value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector3D(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const GLfloat [4][4] value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiPA4_Kf(void* qthis, int32_t arg0, void** arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLint value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLuint value);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEij(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QSize & size);
extern void C_ZN20QOpenGLShaderProgram15setUniformValueEiRK5QSize(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  bool QOpenGLShaderProgram::isLinked();
extern bool C_ZNK20QOpenGLShaderProgram8isLinkedEv(void* qthis); // 4
  // proto:  void QOpenGLShaderProgram::disableAttributeArray(int location);
extern void C_ZN20QOpenGLShaderProgram21disableAttributeArrayEi(void* qthis, int32_t arg0); // 4
  // proto:  void QOpenGLShaderProgram::disableAttributeArray(const char * name);
extern void C_ZN20QOpenGLShaderProgram21disableAttributeArrayEPKc(void* qthis, void* arg0); // 4
  // proto: static bool QOpenGLShaderProgram::hasOpenGLShaderPrograms(QOpenGLContext * context);
extern bool C_ZN20QOpenGLShaderProgram23hasOpenGLShaderProgramsEP14QOpenGLContext(void* arg0); // 4
  // proto:  void QOpenGLShaderProgram::bindAttributeLocation(const char * name, int location);
extern void C_ZN20QOpenGLShaderProgram21bindAttributeLocationEPKci(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QOpenGLShaderProgram::bindAttributeLocation(const QString & name, int location);
extern void C_ZN20QOpenGLShaderProgram21bindAttributeLocationERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QOpenGLShaderProgram::bindAttributeLocation(const QByteArray & name, int location);
extern void C_ZN20QOpenGLShaderProgram21bindAttributeLocationERK10QByteArrayi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  GLuint QOpenGLShaderProgram::programId();
extern int32_t C_ZNK20QOpenGLShaderProgram9programIdEv(void* qthis); // 4
  // proto:  bool QOpenGLShaderProgram::link();
extern bool C_ZN20QOpenGLShaderProgram4linkEv(void* qthis); // 4
  // proto:  void QOpenGLShaderProgram::enableAttributeArray(int location);
extern void C_ZN20QOpenGLShaderProgram20enableAttributeArrayEi(void* qthis, int32_t arg0); // 4
  // proto:  void QOpenGLShaderProgram::enableAttributeArray(const char * name);
extern void C_ZN20QOpenGLShaderProgram20enableAttributeArrayEPKc(void* qthis, void* arg0); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeBuffer(const char * name, GLenum type, int offset, int tupleSize, int stride);
extern void C_ZN20QOpenGLShaderProgram18setAttributeBufferEPKcjiii(void* qthis, void* arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeBuffer(int location, GLenum type, int offset, int tupleSize, int stride);
extern void C_ZN20QOpenGLShaderProgram18setAttributeBufferEijiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4); // 4
  // proto:  int QOpenGLShaderProgram::uniformLocation(const QString & name);
extern int32_t C_ZNK20QOpenGLShaderProgram15uniformLocationERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QOpenGLShaderProgram::uniformLocation(const char * name);
extern int32_t C_ZNK20QOpenGLShaderProgram15uniformLocationEPKc(void* qthis, void* arg0); // 4
  // proto:  int QOpenGLShaderProgram::uniformLocation(const QByteArray & name);
extern int32_t C_ZNK20QOpenGLShaderProgram15uniformLocationERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QOpenGLShaderProgram::~QOpenGLShaderProgram();
extern void C_ZN20QOpenGLShaderProgramD2Ev(void* qthis); // 4
  // proto:  void QOpenGLShaderProgram::QOpenGLShaderProgram(QObject * parent);
extern void* C_ZN20QOpenGLShaderProgramC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QOpenGLShaderProgram::metaObject();
extern void C_ZNK20QOpenGLShaderProgram10metaObjectEv(void* qthis); // 4
  // proto:  bool QOpenGLShaderProgram::bind();
extern bool C_ZN20QOpenGLShaderProgram4bindEv(void* qthis); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const QVector4D * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector4Di(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const GLfloat * values, int count, int tupleSize);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKfii(void* qthis, void* arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const GLuint * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKji(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const QVector3D * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector3Di(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const GLint * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKii(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const QMatrix4x4 * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK10QMatrix4x4i(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const GLuint * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKji(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const QVector2D * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector2Di(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const QMatrix4x4 * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK10QMatrix4x4i(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const GLfloat * values, int count, int tupleSize);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKfii(void* qthis, int32_t arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const QVector4D * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector4Di(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const QVector3D * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector3Di(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const GLint * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKii(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const QVector2D * values, int count);
extern void C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector2Di(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  int QOpenGLShaderProgram::maxGeometryOutputVertices();
extern int32_t C_ZNK20QOpenGLShaderProgram25maxGeometryOutputVerticesEv(void* qthis); // 4
  // proto:  QList<QOpenGLShader *> QOpenGLShaderProgram::shaders();
extern void C_ZNK20QOpenGLShaderProgram7shadersEv(void* qthis); // 4
  // proto:  QVector<float> QOpenGLShaderProgram::defaultInnerTessellationLevels();
extern void C_ZNK20QOpenGLShaderProgram30defaultInnerTessellationLevelsEv(void* qthis); // 4
  // proto:  void QOpenGLShaderProgram::release();
extern void C_ZN20QOpenGLShaderProgram7releaseEv(void* qthis); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, GLfloat x, GLfloat y);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEiff(void* qthis, int32_t arg0, float arg1, float arg2); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const QVector2D & value);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector2D(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, const QColor & value);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEiRK6QColor(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, GLfloat x, GLfloat y, GLfloat z);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEifff(void* qthis, int32_t arg0, float arg1, float arg2, float arg3); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, GLfloat x, GLfloat y, GLfloat z);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcfff(void* qthis, void* arg0, float arg1, float arg2, float arg3); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, const QVector4D & value);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector4D(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, GLfloat x, GLfloat y);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcff(void* qthis, void* arg0, float arg1, float arg2); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, const GLfloat * values, int columns, int rows);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEiPKfii(void* qthis, int32_t arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, const QVector3D & value);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector3D(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const GLfloat * values, int columns, int rows);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcPKfii(void* qthis, void* arg0, void* arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, GLfloat value);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcf(void* qthis, void* arg0, float arg1); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, const QVector2D & value);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector2D(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const QColor & value);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK6QColor(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const QVector3D & value);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector3D(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcffff(void* qthis, void* arg0, float arg1, float arg2, float arg3, float arg4); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const QVector4D & value);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector4D(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, GLfloat value);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEif(void* qthis, int32_t arg0, float arg1); // 4
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
extern void C_ZN20QOpenGLShaderProgram17setAttributeValueEiffff(void* qthis, int32_t arg0, float arg1, float arg2, float arg3, float arg4); // 4
  // proto:  bool QOpenGLShaderProgram::create();
extern bool C_ZN20QOpenGLShaderProgram6createEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QOpenGLShader)=1
type QOpenGLShader struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QOpenGLShaderProgram)=1
type QOpenGLShaderProgram struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// isCompiled()
func (this *QOpenGLShader) Iscompiled(args ...interface{}) (ret interface{}) {
  // isCompiled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader10isCompiledEv
    // invoke: bool isCompiled()
    var ret0 = C.C_ZNK13QOpenGLShader10isCompiledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShader", "isCompiled", args)
  }

  return
}

// log()
func (this *QOpenGLShader) Log(args ...interface{}) (ret interface{}) {
  // log()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader3logEv
    // invoke: QString log()
    var ret0 = C.C_ZNK13QOpenGLShader3logEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShader", "log", args)
  }

  return
}

// compileSourceFile(const class QString &)
func (this *QOpenGLShader) Compilesourcefile(args ...interface{}) (ret interface{}) {
  // compileSourceFile(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLShader17compileSourceFileERK7QString
    // invoke: bool compileSourceFile(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QOpenGLShader17compileSourceFileERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShader", "compileSourceFile", args)
  }

  return
}

// shaderId()
func (this *QOpenGLShader) Shaderid(args ...interface{}) (ret interface{}) {
  // shaderId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader8shaderIdEv
    // invoke: GLuint shaderId()
    var ret0 = C.C_ZNK13QOpenGLShader8shaderIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShader", "shaderId", args)
  }

  return
}

// compileSourceCode(const class QString &)
func (this *QOpenGLShader) Compilesourcecode(args ...interface{}) (ret interface{}) {
  // compileSourceCode(const class QString &)
  // compileSourceCode(const char *)
  // compileSourceCode(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLShader17compileSourceCodeERK7QString
    // invoke: bool compileSourceCode(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QOpenGLShader17compileSourceCodeERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN13QOpenGLShader17compileSourceCodeEPKc
    // invoke: bool compileSourceCode(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN13QOpenGLShader17compileSourceCodeEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN13QOpenGLShader17compileSourceCodeERK10QByteArray
    // invoke: bool compileSourceCode(const class QByteArray &)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QOpenGLShader17compileSourceCodeERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShader", "compileSourceCode", args)
  }

  return
}

// sourceCode()
func (this *QOpenGLShader) Sourcecode(args ...interface{}) (ret interface{}) {
  // sourceCode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader10sourceCodeEv
    // invoke: QByteArray sourceCode()
    var ret0 = C.C_ZNK13QOpenGLShader10sourceCodeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShader", "sourceCode", args)
  }

  return
}

// shaderType()
func (this *QOpenGLShader) Shadertype(args ...interface{}) () {
  // shaderType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader10shaderTypeEv
    // invoke: QOpenGLShader::ShaderType shaderType()
    C.C_ZNK13QOpenGLShader10shaderTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLShader", "shaderType", args)
  }

  return
}

// metaObject()
func (this *QOpenGLShader) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QOpenGLShader10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLShader", "metaObject", args)
  }

  return
}

// ~QOpenGLShader()
func (this *QOpenGLShader) Freeqopenglshader(args ...interface{}) () {
  // ~QOpenGLShader()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLShaderD0Ev
    // invoke: void ~QOpenGLShader()
    C.C_ZN13QOpenGLShaderD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLShader", "~QOpenGLShader", args)
  }

  return
}

// setPatchVertexCount(int)
func (this *QOpenGLShaderProgram) Setpatchvertexcount(args ...interface{}) () {
  // setPatchVertexCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram19setPatchVertexCountEi
    // invoke: void setPatchVertexCount(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN20QOpenGLShaderProgram19setPatchVertexCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setPatchVertexCount", args)
  }

  return
}

// addShader(class QOpenGLShader *)
func (this *QOpenGLShaderProgram) Addshader(args ...interface{}) (ret interface{}) {
  // addShader(class QOpenGLShader *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLShader{}) // "QOpenGLShader *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram9addShaderEP13QOpenGLShader
    // invoke: bool addShader(class QOpenGLShader *)
    var arg0 = args[0].(*QOpenGLShader).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN20QOpenGLShaderProgram9addShaderEP13QOpenGLShader(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "addShader", args)
  }

  return
}

// defaultOuterTessellationLevels()
func (this *QOpenGLShaderProgram) Defaultoutertessellationlevels(args ...interface{}) () {
  // defaultOuterTessellationLevels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram30defaultOuterTessellationLevelsEv
    // invoke: QVector<float> defaultOuterTessellationLevels()
    C.C_ZNK20QOpenGLShaderProgram30defaultOuterTessellationLevelsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "defaultOuterTessellationLevels", args)
  }

  return
}

// attributeLocation(const char *)
func (this *QOpenGLShaderProgram) Attributelocation(args ...interface{}) (ret interface{}) {
  // attributeLocation(const char *)
  // attributeLocation(const class QByteArray &)
  // attributeLocation(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram17attributeLocationEPKc
    // invoke: int attributeLocation(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK20QOpenGLShaderProgram17attributeLocationEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK20QOpenGLShaderProgram17attributeLocationERK10QByteArray
    // invoke: int attributeLocation(const class QByteArray &)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QOpenGLShaderProgram17attributeLocationERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK20QOpenGLShaderProgram17attributeLocationERK7QString
    // invoke: int attributeLocation(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QOpenGLShaderProgram17attributeLocationERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "attributeLocation", args)
  }

  return
}

// setAttributeArray(const char *, const class QVector3D *, int)
func (this *QOpenGLShaderProgram) Setattributearray(args ...interface{}) () {
  // setAttributeArray(const char *, const class QVector3D *, int)
  // setAttributeArray(int, GLenum, const void *, int, int)
  // setAttributeArray(int, const class QVector3D *, int)
  // setAttributeArray(int, const class QVector4D *, int)
  // setAttributeArray(int, const class QVector2D *, int)
  // setAttributeArray(const char *, const class QVector4D *, int)
  // setAttributeArray(int, const GLfloat *, int, int)
  // setAttributeArray(const char *, const class QVector2D *, int)
  // setAttributeArray(const char *, GLenum, const void *, int, int)
  // setAttributeArray(const char *, const GLfloat *, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[1][2] = qtrt.VoidpTy() // "const void *"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(true) // "const char *"
  vtys[5][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "int"
  vtys[6][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[6][2] = qtrt.Int32Ty(false) // "int"
  vtys[6][3] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.ByteTy(true) // "const char *"
  vtys[7][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.ByteTy(true) // "const char *"
  vtys[8][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[8][2] = qtrt.VoidpTy() // "const void *"
  vtys[8][3] = qtrt.Int32Ty(false) // "int"
  vtys[8][4] = qtrt.Int32Ty(false) // "int"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[9][2] = qtrt.Int32Ty(false) // "int"
  vtys[9][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector3Di
    // invoke: void setAttributeArray(const char *, const class QVector3D *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector3Di(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEijPKvii
    // invoke: void setAttributeArray(int, GLenum, const void *, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeArrayEijPKvii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector3Di
    // invoke: void setAttributeArray(int, const class QVector3D *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector3Di(this.Qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector4Di
    // invoke: void setAttributeArray(int, const class QVector4D *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector4D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector4Di(this.Qclsinst, arg0, arg1, arg2)
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector2Di
    // invoke: void setAttributeArray(int, const class QVector2D *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector2D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector2Di(this.Qclsinst, arg0, arg1, arg2)
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector4Di
    // invoke: void setAttributeArray(const char *, const class QVector4D *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[5][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector4D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector4Di(this.Qclsinst, arg0, arg1, arg2)
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPKfii
    // invoke: void setAttributeArray(int, const GLfloat *, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeArrayEiPKfii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector2Di
    // invoke: void setAttributeArray(const char *, const class QVector2D *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[7][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector2D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector2Di(this.Qclsinst, arg0, arg1, arg2)
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcjPKvii
    // invoke: void setAttributeArray(const char *, GLenum, const void *, int, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[8][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeArrayEPKcjPKvii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPKfii
    // invoke: void setAttributeArray(const char *, const GLfloat *, int, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[9][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPKfii(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setAttributeArray", args)
  }

  return
}

// removeShader(class QOpenGLShader *)
func (this *QOpenGLShaderProgram) Removeshader(args ...interface{}) () {
  // removeShader(class QOpenGLShader *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLShader{}) // "QOpenGLShader *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram12removeShaderEP13QOpenGLShader
    // invoke: void removeShader(class QOpenGLShader *)
    var arg0 = args[0].(*QOpenGLShader).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN20QOpenGLShaderProgram12removeShaderEP13QOpenGLShader(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "removeShader", args)
  }

  return
}

// removeAllShaders()
func (this *QOpenGLShaderProgram) Removeallshaders(args ...interface{}) () {
  // removeAllShaders()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram16removeAllShadersEv
    // invoke: void removeAllShaders()
    C.C_ZN20QOpenGLShaderProgram16removeAllShadersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "removeAllShaders", args)
  }

  return
}

// patchVertexCount()
func (this *QOpenGLShaderProgram) Patchvertexcount(args ...interface{}) (ret interface{}) {
  // patchVertexCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram16patchVertexCountEv
    // invoke: int patchVertexCount()
    var ret0 = C.C_ZNK20QOpenGLShaderProgram16patchVertexCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "patchVertexCount", args)
  }

  return
}

// log()
func (this *QOpenGLShaderProgram) Log(args ...interface{}) (ret interface{}) {
  // log()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram3logEv
    // invoke: QString log()
    var ret0 = C.C_ZNK20QOpenGLShaderProgram3logEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "log", args)
  }

  return
}

// setUniformValue(const char *, const class QSize &)
func (this *QOpenGLShaderProgram) Setuniformvalue(args ...interface{}) () {
  // setUniformValue(const char *, const class QSize &)
  // setUniformValue(int, const class QColor &)
  // setUniformValue(int, const GLfloat (*)[3])
  // setUniformValue(const char *, const class QColor &)
  // setUniformValue(int, const class QPoint &)
  // setUniformValue(int, GLfloat, GLfloat)
  // setUniformValue(const char *, GLfloat, GLfloat, GLfloat, GLfloat)
  // setUniformValue(const char *, const class QPointF &)
  // setUniformValue(int, const class QTransform &)
  // setUniformValue(const char *, GLfloat, GLfloat, GLfloat)
  // setUniformValue(int, const GLfloat (*)[2])
  // setUniformValue(const char *, const class QMatrix4x4 &)
  // setUniformValue(const char *, const class QVector3D &)
  // setUniformValue(const char *, GLuint)
  // setUniformValue(const char *, GLint)
  // setUniformValue(const char *, const GLfloat (*)[2])
  // setUniformValue(int, GLfloat, GLfloat, GLfloat)
  // setUniformValue(const char *, GLfloat)
  // setUniformValue(int, const class QVector4D &)
  // setUniformValue(const char *, const class QPoint &)
  // setUniformValue(int, const class QSizeF &)
  // setUniformValue(const char *, const class QSizeF &)
  // setUniformValue(int, GLfloat, GLfloat, GLfloat, GLfloat)
  // setUniformValue(const char *, const GLfloat (*)[3])
  // setUniformValue(int, const class QVector2D &)
  // setUniformValue(const char *, const GLfloat (*)[4])
  // setUniformValue(const char *, const class QVector4D &)
  // setUniformValue(const char *, GLfloat, GLfloat)
  // setUniformValue(const char *, const class QTransform &)
  // setUniformValue(const char *, const class QVector2D &)
  // setUniformValue(int, const class QPointF &)
  // setUniformValue(int, GLfloat)
  // setUniformValue(int, const class QMatrix4x4 &)
  // setUniformValue(int, const class QVector3D &)
  // setUniformValue(int, const GLfloat (*)[4])
  // setUniformValue(int, GLint)
  // setUniformValue(int, GLuint)
  // setUniformValue(int, const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.VoidpTy() // "const GLfloat [3][3]"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[5][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  vtys[6][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.ByteTy(true) // "const char *"
  vtys[7][1] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int32Ty(false) // "int"
  vtys[8][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[9][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[9][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.Int32Ty(false) // "int"
  vtys[10][1] = qtrt.VoidpTy() // "const GLfloat [2][2]"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.ByteTy(true) // "const char *"
  vtys[11][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.ByteTy(true) // "const char *"
  vtys[12][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.ByteTy(true) // "const char *"
  vtys[13][1] = qtrt.Int32Ty(false) // "GLuint"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.ByteTy(true) // "const char *"
  vtys[14][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = qtrt.ByteTy(true) // "const char *"
  vtys[15][1] = qtrt.VoidpTy() // "const GLfloat [2][2]"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.Int32Ty(false) // "int"
  vtys[16][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[16][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[16][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = qtrt.ByteTy(true) // "const char *"
  vtys[17][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[18] = make(map[int32]reflect.Type)
  vtys[18][0] = qtrt.Int32Ty(false) // "int"
  vtys[18][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[19] = make(map[int32]reflect.Type)
  vtys[19][0] = qtrt.ByteTy(true) // "const char *"
  vtys[19][1] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[20] = make(map[int32]reflect.Type)
  vtys[20][0] = qtrt.Int32Ty(false) // "int"
  vtys[20][1] = reflect.TypeOf(qtcore.QSizeF{}) // "const QSizeF &"
  vtys[21] = make(map[int32]reflect.Type)
  vtys[21][0] = qtrt.ByteTy(true) // "const char *"
  vtys[21][1] = reflect.TypeOf(qtcore.QSizeF{}) // "const QSizeF &"
  vtys[22] = make(map[int32]reflect.Type)
  vtys[22][0] = qtrt.Int32Ty(false) // "int"
  vtys[22][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[22][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[22][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[22][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[23] = make(map[int32]reflect.Type)
  vtys[23][0] = qtrt.ByteTy(true) // "const char *"
  vtys[23][1] = qtrt.VoidpTy() // "const GLfloat [3][3]"
  vtys[24] = make(map[int32]reflect.Type)
  vtys[24][0] = qtrt.Int32Ty(false) // "int"
  vtys[24][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[25] = make(map[int32]reflect.Type)
  vtys[25][0] = qtrt.ByteTy(true) // "const char *"
  vtys[25][1] = qtrt.VoidpTy() // "const GLfloat [4][4]"
  vtys[26] = make(map[int32]reflect.Type)
  vtys[26][0] = qtrt.ByteTy(true) // "const char *"
  vtys[26][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[27] = make(map[int32]reflect.Type)
  vtys[27][0] = qtrt.ByteTy(true) // "const char *"
  vtys[27][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[27][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[28] = make(map[int32]reflect.Type)
  vtys[28][0] = qtrt.ByteTy(true) // "const char *"
  vtys[28][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[29] = make(map[int32]reflect.Type)
  vtys[29][0] = qtrt.ByteTy(true) // "const char *"
  vtys[29][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[30] = make(map[int32]reflect.Type)
  vtys[30][0] = qtrt.Int32Ty(false) // "int"
  vtys[30][1] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  vtys[31] = make(map[int32]reflect.Type)
  vtys[31][0] = qtrt.Int32Ty(false) // "int"
  vtys[31][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[32] = make(map[int32]reflect.Type)
  vtys[32][0] = qtrt.Int32Ty(false) // "int"
  vtys[32][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[33] = make(map[int32]reflect.Type)
  vtys[33][0] = qtrt.Int32Ty(false) // "int"
  vtys[33][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[34] = make(map[int32]reflect.Type)
  vtys[34][0] = qtrt.Int32Ty(false) // "int"
  vtys[34][1] = qtrt.VoidpTy() // "const GLfloat [4][4]"
  vtys[35] = make(map[int32]reflect.Type)
  vtys[35][0] = qtrt.Int32Ty(false) // "int"
  vtys[35][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[36] = make(map[int32]reflect.Type)
  vtys[36][0] = qtrt.Int32Ty(false) // "int"
  vtys[36][1] = qtrt.Int32Ty(false) // "GLuint"
  vtys[37] = make(map[int32]reflect.Type)
  vtys[37][0] = qtrt.Int32Ty(false) // "int"
  vtys[37][1] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK5QSize
    // invoke: void setUniformValue(const char *, const class QSize &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK5QSize(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QColor
    // invoke: void setUniformValue(int, const class QColor &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QColor).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiRK6QColor(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiPA3_Kf
    // invoke: void setUniformValue(int, const GLfloat (*)[3])
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (**C.float)((unsafe.Pointer)(reflect.ValueOf(args[1].([][]float32)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiPA3_Kf(this.Qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QColor
    // invoke: void setUniformValue(const char *, const class QColor &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[3][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QColor).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QColor(this.Qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QPoint
    // invoke: void setUniformValue(int, const class QPoint &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiRK6QPoint(this.Qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiff
    // invoke: void setUniformValue(int, GLfloat, GLfloat)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiff(this.Qclsinst, arg0, arg1, arg2)
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcffff
    // invoke: void setUniformValue(const char *, GLfloat, GLfloat, GLfloat, GLfloat)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[6][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(qtrt.PrimConv(args[4], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg4)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcffff(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK7QPointF
    // invoke: void setUniformValue(const char *, const class QPointF &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[7][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK7QPointF(this.Qclsinst, arg0, arg1)
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK10QTransform
    // invoke: void setUniformValue(int, const class QTransform &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTransform).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiRK10QTransform(this.Qclsinst, arg0, arg1)
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcfff
    // invoke: void setUniformValue(const char *, GLfloat, GLfloat, GLfloat)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[9][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcfff(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 10:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiPA2_Kf
    // invoke: void setUniformValue(int, const GLfloat (*)[2])
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (**C.float)((unsafe.Pointer)(reflect.ValueOf(args[1].([][]float32)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiPA2_Kf(this.Qclsinst, arg0, arg1)
  case 11:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QMatrix4x4
    // invoke: void setUniformValue(const char *, const class QMatrix4x4 &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[11][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QMatrix4x4).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QMatrix4x4(this.Qclsinst, arg0, arg1)
  case 12:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector3D
    // invoke: void setUniformValue(const char *, const class QVector3D &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[12][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector3D(this.Qclsinst, arg0, arg1)
  case 13:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcj
    // invoke: void setUniformValue(const char *, GLuint)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[13][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcj(this.Qclsinst, arg0, arg1)
  case 14:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKci
    // invoke: void setUniformValue(const char *, GLint)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[14][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKci(this.Qclsinst, arg0, arg1)
  case 15:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA2_Kf
    // invoke: void setUniformValue(const char *, const GLfloat (*)[2])
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[15][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (**C.float)((unsafe.Pointer)(reflect.ValueOf(args[1].([][]float32)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcPA2_Kf(this.Qclsinst, arg0, arg1)
  case 16:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEifff
    // invoke: void setUniformValue(int, GLfloat, GLfloat, GLfloat)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEifff(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 17:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcf
    // invoke: void setUniformValue(const char *, GLfloat)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[17][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcf(this.Qclsinst, arg0, arg1)
  case 18:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector4D
    // invoke: void setUniformValue(int, const class QVector4D &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector4D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector4D(this.Qclsinst, arg0, arg1)
  case 19:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QPoint
    // invoke: void setUniformValue(const char *, const class QPoint &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[19][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QPoint(this.Qclsinst, arg0, arg1)
  case 20:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QSizeF
    // invoke: void setUniformValue(int, const class QSizeF &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QSizeF).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiRK6QSizeF(this.Qclsinst, arg0, arg1)
  case 21:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QSizeF
    // invoke: void setUniformValue(const char *, const class QSizeF &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[21][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*qtcore.QSizeF).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QSizeF(this.Qclsinst, arg0, arg1)
  case 22:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiffff
    // invoke: void setUniformValue(int, GLfloat, GLfloat, GLfloat, GLfloat)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(qtrt.PrimConv(args[4], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg4)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiffff(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 23:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA3_Kf
    // invoke: void setUniformValue(const char *, const GLfloat (*)[3])
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[23][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (**C.float)((unsafe.Pointer)(reflect.ValueOf(args[1].([][]float32)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcPA3_Kf(this.Qclsinst, arg0, arg1)
  case 24:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector2D
    // invoke: void setUniformValue(int, const class QVector2D &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector2D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector2D(this.Qclsinst, arg0, arg1)
  case 25:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA4_Kf
    // invoke: void setUniformValue(const char *, const GLfloat (*)[4])
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[25][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (**C.float)((unsafe.Pointer)(reflect.ValueOf(args[1].([][]float32)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcPA4_Kf(this.Qclsinst, arg0, arg1)
  case 26:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector4D
    // invoke: void setUniformValue(const char *, const class QVector4D &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[26][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector4D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector4D(this.Qclsinst, arg0, arg1)
  case 27:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcff
    // invoke: void setUniformValue(const char *, GLfloat, GLfloat)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[27][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcff(this.Qclsinst, arg0, arg1, arg2)
  case 28:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QTransform
    // invoke: void setUniformValue(const char *, const class QTransform &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[28][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QTransform).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QTransform(this.Qclsinst, arg0, arg1)
  case 29:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector2D
    // invoke: void setUniformValue(const char *, const class QVector2D &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[29][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector2D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector2D(this.Qclsinst, arg0, arg1)
  case 30:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK7QPointF
    // invoke: void setUniformValue(int, const class QPointF &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiRK7QPointF(this.Qclsinst, arg0, arg1)
  case 31:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEif
    // invoke: void setUniformValue(int, GLfloat)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEif(this.Qclsinst, arg0, arg1)
  case 32:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK10QMatrix4x4
    // invoke: void setUniformValue(int, const class QMatrix4x4 &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QMatrix4x4).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiRK10QMatrix4x4(this.Qclsinst, arg0, arg1)
  case 33:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector3D
    // invoke: void setUniformValue(int, const class QVector3D &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector3D(this.Qclsinst, arg0, arg1)
  case 34:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiPA4_Kf
    // invoke: void setUniformValue(int, const GLfloat (*)[4])
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (**C.float)((unsafe.Pointer)(reflect.ValueOf(args[1].([][]float32)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiPA4_Kf(this.Qclsinst, arg0, arg1)
  case 35:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEii
    // invoke: void setUniformValue(int, GLint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEii(this.Qclsinst, arg0, arg1)
  case 36:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEij
    // invoke: void setUniformValue(int, GLuint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEij(this.Qclsinst, arg0, arg1)
  case 37:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK5QSize
    // invoke: void setUniformValue(int, const class QSize &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram15setUniformValueEiRK5QSize(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setUniformValue", args)
  }

  return
}

// isLinked()
func (this *QOpenGLShaderProgram) Islinked(args ...interface{}) (ret interface{}) {
  // isLinked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram8isLinkedEv
    // invoke: bool isLinked()
    var ret0 = C.C_ZNK20QOpenGLShaderProgram8isLinkedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "isLinked", args)
  }

  return
}

// disableAttributeArray(int)
func (this *QOpenGLShaderProgram) Disableattributearray(args ...interface{}) () {
  // disableAttributeArray(int)
  // disableAttributeArray(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram21disableAttributeArrayEi
    // invoke: void disableAttributeArray(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN20QOpenGLShaderProgram21disableAttributeArrayEi(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram21disableAttributeArrayEPKc
    // invoke: void disableAttributeArray(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN20QOpenGLShaderProgram21disableAttributeArrayEPKc(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "disableAttributeArray", args)
  }

  return
}

// hasOpenGLShaderPrograms(class QOpenGLContext *)
func (this *QOpenGLShaderProgram) Hasopenglshaderprograms_S(args ...interface{}) (ret interface{}) {
  // hasOpenGLShaderPrograms(class QOpenGLContext *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLContext{}) // "QOpenGLContext *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram23hasOpenGLShaderProgramsEP14QOpenGLContext
    // invoke: bool hasOpenGLShaderPrograms(class QOpenGLContext *)
    var arg0 = args[0].(*QOpenGLContext).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN20QOpenGLShaderProgram23hasOpenGLShaderProgramsEP14QOpenGLContext(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "hasOpenGLShaderPrograms", args)
  }

  return
}

// bindAttributeLocation(const char *, int)
func (this *QOpenGLShaderProgram) Bindattributelocation(args ...interface{}) () {
  // bindAttributeLocation(const char *, int)
  // bindAttributeLocation(const class QString &, int)
  // bindAttributeLocation(const class QByteArray &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram21bindAttributeLocationEPKci
    // invoke: void bindAttributeLocation(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram21bindAttributeLocationEPKci(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram21bindAttributeLocationERK7QStringi
    // invoke: void bindAttributeLocation(const class QString &, int)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram21bindAttributeLocationERK7QStringi(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram21bindAttributeLocationERK10QByteArrayi
    // invoke: void bindAttributeLocation(const class QByteArray &, int)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram21bindAttributeLocationERK10QByteArrayi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "bindAttributeLocation", args)
  }

  return
}

// programId()
func (this *QOpenGLShaderProgram) Programid(args ...interface{}) (ret interface{}) {
  // programId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram9programIdEv
    // invoke: GLuint programId()
    var ret0 = C.C_ZNK20QOpenGLShaderProgram9programIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "programId", args)
  }

  return
}

// link()
func (this *QOpenGLShaderProgram) Link(args ...interface{}) (ret interface{}) {
  // link()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram4linkEv
    // invoke: bool link()
    var ret0 = C.C_ZN20QOpenGLShaderProgram4linkEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "link", args)
  }

  return
}

// enableAttributeArray(int)
func (this *QOpenGLShaderProgram) Enableattributearray(args ...interface{}) () {
  // enableAttributeArray(int)
  // enableAttributeArray(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram20enableAttributeArrayEi
    // invoke: void enableAttributeArray(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN20QOpenGLShaderProgram20enableAttributeArrayEi(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram20enableAttributeArrayEPKc
    // invoke: void enableAttributeArray(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN20QOpenGLShaderProgram20enableAttributeArrayEPKc(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "enableAttributeArray", args)
  }

  return
}

// setAttributeBuffer(const char *, GLenum, int, int, int)
func (this *QOpenGLShaderProgram) Setattributebuffer(args ...interface{}) () {
  // setAttributeBuffer(const char *, GLenum, int, int, int)
  // setAttributeBuffer(int, GLenum, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram18setAttributeBufferEPKcjiii
    // invoke: void setAttributeBuffer(const char *, GLenum, int, int, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN20QOpenGLShaderProgram18setAttributeBufferEPKcjiii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram18setAttributeBufferEijiii
    // invoke: void setAttributeBuffer(int, GLenum, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN20QOpenGLShaderProgram18setAttributeBufferEijiii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setAttributeBuffer", args)
  }

  return
}

// uniformLocation(const class QString &)
func (this *QOpenGLShaderProgram) Uniformlocation(args ...interface{}) (ret interface{}) {
  // uniformLocation(const class QString &)
  // uniformLocation(const char *)
  // uniformLocation(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram15uniformLocationERK7QString
    // invoke: int uniformLocation(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QOpenGLShaderProgram15uniformLocationERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK20QOpenGLShaderProgram15uniformLocationEPKc
    // invoke: int uniformLocation(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK20QOpenGLShaderProgram15uniformLocationEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK20QOpenGLShaderProgram15uniformLocationERK10QByteArray
    // invoke: int uniformLocation(const class QByteArray &)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK20QOpenGLShaderProgram15uniformLocationERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "uniformLocation", args)
  }

  return
}

// ~QOpenGLShaderProgram()
func (this *QOpenGLShaderProgram) Freeqopenglshaderprogram(args ...interface{}) () {
  // ~QOpenGLShaderProgram()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgramD0Ev
    // invoke: void ~QOpenGLShaderProgram()
    C.C_ZN20QOpenGLShaderProgramD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "~QOpenGLShaderProgram", args)
  }

  return
}

// QOpenGLShaderProgram(class QObject *)
func NewQOpenGLShaderProgram(args ...interface{}) *QOpenGLShaderProgram {
  // QOpenGLShaderProgram(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgramC1EP7QObject
    // invoke: void QOpenGLShaderProgram(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QOpenGLShaderProgramC2EP7QObject(arg0)
    return &QOpenGLShaderProgram{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "QOpenGLShaderProgram", args)
  }

  return nil // QOpenGLShaderProgram{Qclsinst:qthis}
}

// metaObject()
func (this *QOpenGLShaderProgram) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK20QOpenGLShaderProgram10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "metaObject", args)
  }

  return
}

// bind()
func (this *QOpenGLShaderProgram) Bind(args ...interface{}) (ret interface{}) {
  // bind()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram4bindEv
    // invoke: bool bind()
    var ret0 = C.C_ZN20QOpenGLShaderProgram4bindEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "bind", args)
  }

  return
}

// setUniformValueArray(int, const class QVector4D *, int)
func (this *QOpenGLShaderProgram) Setuniformvaluearray(args ...interface{}) () {
  // setUniformValueArray(int, const class QVector4D *, int)
  // setUniformValueArray(const char *, const GLfloat *, int, int)
  // setUniformValueArray(int, const GLuint *, int)
  // setUniformValueArray(int, const class QVector3D *, int)
  // setUniformValueArray(const char *, const GLint *, int)
  // setUniformValueArray(int, const class QMatrix4x4 *, int)
  // setUniformValueArray(const char *, const GLuint *, int)
  // setUniformValueArray(const char *, const class QVector2D *, int)
  // setUniformValueArray(const char *, const class QMatrix4x4 *, int)
  // setUniformValueArray(int, const GLfloat *, int, int)
  // setUniformValueArray(const char *, const class QVector4D *, int)
  // setUniformValueArray(const char *, const class QVector3D *, int)
  // setUniformValueArray(int, const GLint *, int)
  // setUniformValueArray(int, const class QVector2D *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(true) // "const GLuint *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(true) // "const char *"
  vtys[4][1] = qtrt.Int32Ty(true) // "const GLint *"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 *"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  vtys[6][1] = qtrt.Int32Ty(true) // "const GLuint *"
  vtys[6][2] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.ByteTy(true) // "const char *"
  vtys[7][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.ByteTy(true) // "const char *"
  vtys[8][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 *"
  vtys[8][2] = qtrt.Int32Ty(false) // "int"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.Int32Ty(false) // "int"
  vtys[9][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[9][2] = qtrt.Int32Ty(false) // "int"
  vtys[9][3] = qtrt.Int32Ty(false) // "int"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.ByteTy(true) // "const char *"
  vtys[10][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[10][2] = qtrt.Int32Ty(false) // "int"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.ByteTy(true) // "const char *"
  vtys[11][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[11][2] = qtrt.Int32Ty(false) // "int"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.Int32Ty(false) // "int"
  vtys[12][1] = qtrt.Int32Ty(true) // "const GLint *"
  vtys[12][2] = qtrt.Int32Ty(false) // "int"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.Int32Ty(false) // "int"
  vtys[13][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[13][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector4Di
    // invoke: void setUniformValueArray(int, const class QVector4D *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector4D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector4Di(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKfii
    // invoke: void setUniformValueArray(const char *, const GLfloat *, int, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKfii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKji
    // invoke: void setUniformValueArray(int, const GLuint *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKji(this.Qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector3Di
    // invoke: void setUniformValueArray(int, const class QVector3D *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector3Di(this.Qclsinst, arg0, arg1, arg2)
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKii
    // invoke: void setUniformValueArray(const char *, const GLint *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[4][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKii(this.Qclsinst, arg0, arg1, arg2)
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK10QMatrix4x4i
    // invoke: void setUniformValueArray(int, const class QMatrix4x4 *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QMatrix4x4).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK10QMatrix4x4i(this.Qclsinst, arg0, arg1, arg2)
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKji
    // invoke: void setUniformValueArray(const char *, const GLuint *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[6][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKji(this.Qclsinst, arg0, arg1, arg2)
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector2Di
    // invoke: void setUniformValueArray(const char *, const class QVector2D *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[7][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector2D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector2Di(this.Qclsinst, arg0, arg1, arg2)
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK10QMatrix4x4i
    // invoke: void setUniformValueArray(const char *, const class QMatrix4x4 *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[8][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QMatrix4x4).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK10QMatrix4x4i(this.Qclsinst, arg0, arg1, arg2)
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKfii
    // invoke: void setUniformValueArray(int, const GLfloat *, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKfii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 10:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector4Di
    // invoke: void setUniformValueArray(const char *, const class QVector4D *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[10][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector4D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector4Di(this.Qclsinst, arg0, arg1, arg2)
  case 11:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector3Di
    // invoke: void setUniformValueArray(const char *, const class QVector3D *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[11][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector3Di(this.Qclsinst, arg0, arg1, arg2)
  case 12:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKii
    // invoke: void setUniformValueArray(int, const GLint *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKii(this.Qclsinst, arg0, arg1, arg2)
  case 13:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector2Di
    // invoke: void setUniformValueArray(int, const class QVector2D *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector2D).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector2Di(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setUniformValueArray", args)
  }

  return
}

// maxGeometryOutputVertices()
func (this *QOpenGLShaderProgram) Maxgeometryoutputvertices(args ...interface{}) (ret interface{}) {
  // maxGeometryOutputVertices()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram25maxGeometryOutputVerticesEv
    // invoke: int maxGeometryOutputVertices()
    var ret0 = C.C_ZNK20QOpenGLShaderProgram25maxGeometryOutputVerticesEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "maxGeometryOutputVertices", args)
  }

  return
}

// shaders()
func (this *QOpenGLShaderProgram) Shaders(args ...interface{}) () {
  // shaders()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram7shadersEv
    // invoke: QList<QOpenGLShader *> shaders()
    C.C_ZNK20QOpenGLShaderProgram7shadersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "shaders", args)
  }

  return
}

// defaultInnerTessellationLevels()
func (this *QOpenGLShaderProgram) Defaultinnertessellationlevels(args ...interface{}) () {
  // defaultInnerTessellationLevels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram30defaultInnerTessellationLevelsEv
    // invoke: QVector<float> defaultInnerTessellationLevels()
    C.C_ZNK20QOpenGLShaderProgram30defaultInnerTessellationLevelsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "defaultInnerTessellationLevels", args)
  }

  return
}

// release()
func (this *QOpenGLShaderProgram) Release(args ...interface{}) () {
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram7releaseEv
    // invoke: void release()
    C.C_ZN20QOpenGLShaderProgram7releaseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "release", args)
  }

  return
}

// setAttributeValue(int, GLfloat, GLfloat)
func (this *QOpenGLShaderProgram) Setattributevalue(args ...interface{}) () {
  // setAttributeValue(int, GLfloat, GLfloat)
  // setAttributeValue(const char *, const class QVector2D &)
  // setAttributeValue(int, const class QColor &)
  // setAttributeValue(int, GLfloat, GLfloat, GLfloat)
  // setAttributeValue(const char *, GLfloat, GLfloat, GLfloat)
  // setAttributeValue(int, const class QVector4D &)
  // setAttributeValue(const char *, GLfloat, GLfloat)
  // setAttributeValue(int, const GLfloat *, int, int)
  // setAttributeValue(int, const class QVector3D &)
  // setAttributeValue(const char *, const GLfloat *, int, int)
  // setAttributeValue(const char *, GLfloat)
  // setAttributeValue(int, const class QVector2D &)
  // setAttributeValue(const char *, const class QColor &)
  // setAttributeValue(const char *, const class QVector3D &)
  // setAttributeValue(const char *, GLfloat, GLfloat, GLfloat, GLfloat)
  // setAttributeValue(const char *, const class QVector4D &)
  // setAttributeValue(int, GLfloat)
  // setAttributeValue(int, GLfloat, GLfloat, GLfloat, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[0][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[3][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[3][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(true) // "const char *"
  vtys[4][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[4][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[4][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  vtys[6][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[7][3] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int32Ty(false) // "int"
  vtys[8][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[9][2] = qtrt.Int32Ty(false) // "int"
  vtys[9][3] = qtrt.Int32Ty(false) // "int"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.ByteTy(true) // "const char *"
  vtys[10][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int32Ty(false) // "int"
  vtys[11][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.ByteTy(true) // "const char *"
  vtys[12][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.ByteTy(true) // "const char *"
  vtys[13][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.ByteTy(true) // "const char *"
  vtys[14][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[14][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[14][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[14][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = qtrt.ByteTy(true) // "const char *"
  vtys[15][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.Int32Ty(false) // "int"
  vtys[16][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = qtrt.Int32Ty(false) // "int"
  vtys[17][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17][4] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiff
    // invoke: void setAttributeValue(int, GLfloat, GLfloat)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEiff(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector2D
    // invoke: void setAttributeValue(const char *, const class QVector2D &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector2D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector2D(this.Qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK6QColor
    // invoke: void setAttributeValue(int, const class QColor &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QColor).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEiRK6QColor(this.Qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEifff
    // invoke: void setAttributeValue(int, GLfloat, GLfloat, GLfloat)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEifff(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcfff
    // invoke: void setAttributeValue(const char *, GLfloat, GLfloat, GLfloat)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[4][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcfff(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector4D
    // invoke: void setAttributeValue(int, const class QVector4D &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector4D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector4D(this.Qclsinst, arg0, arg1)
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcff
    // invoke: void setAttributeValue(const char *, GLfloat, GLfloat)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[6][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcff(this.Qclsinst, arg0, arg1, arg2)
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiPKfii
    // invoke: void setAttributeValue(int, const GLfloat *, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEiPKfii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector3D
    // invoke: void setAttributeValue(int, const class QVector3D &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector3D(this.Qclsinst, arg0, arg1)
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcPKfii
    // invoke: void setAttributeValue(const char *, const GLfloat *, int, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[9][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcPKfii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 10:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcf
    // invoke: void setAttributeValue(const char *, GLfloat)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[10][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcf(this.Qclsinst, arg0, arg1)
  case 11:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector2D
    // invoke: void setAttributeValue(int, const class QVector2D &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVector2D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector2D(this.Qclsinst, arg0, arg1)
  case 12:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK6QColor
    // invoke: void setAttributeValue(const char *, const class QColor &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[12][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QColor).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK6QColor(this.Qclsinst, arg0, arg1)
  case 13:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector3D
    // invoke: void setAttributeValue(const char *, const class QVector3D &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[13][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector3D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector3D(this.Qclsinst, arg0, arg1)
  case 14:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcffff
    // invoke: void setAttributeValue(const char *, GLfloat, GLfloat, GLfloat, GLfloat)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[14][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(qtrt.PrimConv(args[4], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg4)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcffff(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 15:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector4D
    // invoke: void setAttributeValue(const char *, const class QVector4D &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[15][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QVector4D).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector4D(this.Qclsinst, arg0, arg1)
  case 16:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEif
    // invoke: void setAttributeValue(int, GLfloat)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEif(this.Qclsinst, arg0, arg1)
  case 17:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiffff
    // invoke: void setAttributeValue(int, GLfloat, GLfloat, GLfloat, GLfloat)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(qtrt.PrimConv(args[1], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(qtrt.PrimConv(args[2], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(qtrt.PrimConv(args[3], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(qtrt.PrimConv(args[4], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg4)}
    C.C_ZN20QOpenGLShaderProgram17setAttributeValueEiffff(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setAttributeValue", args)
  }

  return
}

// create()
func (this *QOpenGLShaderProgram) Create(args ...interface{}) (ret interface{}) {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram6createEv
    // invoke: bool create()
    var ret0 = C.C_ZN20QOpenGLShaderProgram6createEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "create", args)
  }

  return
}

// <= body block end

