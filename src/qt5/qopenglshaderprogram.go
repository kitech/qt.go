package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qopenglshaderprogram.h
// dst-file: /src/gui/qopenglshaderprogram.go
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
// class sizeof(QOpenGLShader)=1
type QOpenGLShader struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLShaderProgram)=1
type QOpenGLShaderProgram struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQOpenGLShader(args ...interface{})() {
}


func (this *QOpenGLShader) isCompiled(args ...interface{}) () {
  // isCompiled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader10isCompiledEv
  default:
    qtrt.ErrorResolve("QOpenGLShader", "isCompiled", args)
 }

}


func (this *QOpenGLShader) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader10metaObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLShader", "metaObject", args)
 }

}


func (this *QOpenGLShader) log(args ...interface{}) () {
  // log()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader3logEv
  default:
    qtrt.ErrorResolve("QOpenGLShader", "log", args)
 }

}


func (this *QOpenGLShader) compileSourceCode(args ...interface{}) () {
  // compileSourceCode(const class QString &)
  // compileSourceCode(const class QByteArray &)
  // compileSourceCode(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLShader17compileSourceCodeERK7QString
  case 1:
    // invoke: _ZN13QOpenGLShader17compileSourceCodeERK10QByteArray
  case 2:
    // invoke: _ZN13QOpenGLShader17compileSourceCodeEPKc
  default:
    qtrt.ErrorResolve("QOpenGLShader", "compileSourceCode", args)
 }

}


func (this *QOpenGLShader) compileSourceFile(args ...interface{}) () {
  // compileSourceFile(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLShader17compileSourceFileERK7QString
  default:
    qtrt.ErrorResolve("QOpenGLShader", "compileSourceFile", args)
 }

}


func (this *QOpenGLShader) sourceCode(args ...interface{}) () {
  // sourceCode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader10sourceCodeEv
  default:
    qtrt.ErrorResolve("QOpenGLShader", "sourceCode", args)
 }

}


func (this *QOpenGLShader) FreeQOpenGLShader(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLShader", "~QOpenGLShader", args)
 }

}


func (this *QOpenGLShader) shaderId(args ...interface{}) () {
  // shaderId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader8shaderIdEv
  default:
    qtrt.ErrorResolve("QOpenGLShader", "shaderId", args)
 }

}


func (this *QOpenGLShaderProgram) isLinked(args ...interface{}) () {
  // isLinked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram8isLinkedEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "isLinked", args)
 }

}


func (this *QOpenGLShaderProgram) setUniformValue(args ...interface{}) () {
  // setUniformValue(const char *, const class QVector3D &)
  // setUniformValue(int, const class QPoint &)
  // setUniformValue(const char *, GLfloat, GLfloat, GLfloat, GLfloat)
  // setUniformValue(int, const class QTransform &)
  // setUniformValue(int, const QMatrix2x3 &)
  // setUniformValue(const char *, const GLfloat (*)[2])
  // setUniformValue(const char *, const QMatrix3x3 &)
  // setUniformValue(int, const QMatrix3x2 &)
  // setUniformValue(const char *, const QMatrix4x3 &)
  // setUniformValue(const char *, const GLfloat (*)[3])
  // setUniformValue(int, const class QVector2D &)
  // setUniformValue(int, GLfloat)
  // setUniformValue(const char *, const QMatrix2x4 &)
  // setUniformValue(int, GLint)
  // setUniformValue(int, GLuint)
  // setUniformValue(int, const class QSize &)
  // setUniformValue(int, const QMatrix2x2 &)
  // setUniformValue(const char *, const class QPoint &)
  // setUniformValue(int, GLfloat, GLfloat)
  // setUniformValue(int, GLfloat, GLfloat, GLfloat)
  // setUniformValue(const char *, const class QPointF &)
  // setUniformValue(const char *, const QMatrix3x2 &)
  // setUniformValue(const char *, const QMatrix2x3 &)
  // setUniformValue(int, const class QColor &)
  // setUniformValue(const char *, const GLfloat (*)[4])
  // setUniformValue(const char *, const QMatrix2x2 &)
  // setUniformValue(const char *, const class QSize &)
  // setUniformValue(const char *, const QMatrix4x2 &)
  // setUniformValue(int, const GLfloat (*)[2])
  // setUniformValue(const char *, const class QMatrix4x4 &)
  // setUniformValue(int, const class QSizeF &)
  // setUniformValue(const char *, const QMatrix3x4 &)
  // setUniformValue(const char *, const class QVector4D &)
  // setUniformValue(const char *, GLfloat, GLfloat)
  // setUniformValue(int, const QMatrix3x4 &)
  // setUniformValue(int, const class QPointF &)
  // setUniformValue(int, const QMatrix4x2 &)
  // setUniformValue(int, const GLfloat (*)[4])
  // setUniformValue(int, const QMatrix3x3 &)
  // setUniformValue(const char *, const class QColor &)
  // setUniformValue(const char *, const class QTransform &)
  // setUniformValue(int, const class QVector3D &)
  // setUniformValue(const char *, GLfloat, GLfloat, GLfloat)
  // setUniformValue(int, const GLfloat (*)[3])
  // setUniformValue(const char *, GLuint)
  // setUniformValue(const char *, GLint)
  // setUniformValue(int, const QMatrix4x3 &)
  // setUniformValue(const char *, GLfloat)
  // setUniformValue(int, const class QMatrix4x4 &)
  // setUniformValue(int, const class QVector4D &)
  // setUniformValue(const char *, const class QSizeF &)
  // setUniformValue(int, GLfloat, GLfloat, GLfloat, GLfloat)
  // setUniformValue(int, const QMatrix2x4 &)
  // setUniformValue(const char *, const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[2][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[2][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[2][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  // vtys[4][1] = reflect.TypeOf(QGenericMatrix<2, 3, float>{}) // "const QMatrix2x3 &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(true) // "const char *"
  vtys[5][1] = qtrt.VoidpTy() // "const GLfloat [2][2]"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[6][1] = reflect.TypeOf(QGenericMatrix<3, 3, float>{}) // "const QMatrix3x3 &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  // vtys[7][1] = reflect.TypeOf(QGenericMatrix<3, 2, float>{}) // "const QMatrix3x2 &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[8][1] = reflect.TypeOf(QGenericMatrix<4, 3, float>{}) // "const QMatrix4x3 &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = qtrt.VoidpTy() // "const GLfloat [3][3]"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.Int32Ty(false) // "int"
  vtys[10][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int32Ty(false) // "int"
  vtys[11][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[12][1] = reflect.TypeOf(QGenericMatrix<2, 4, float>{}) // "const QMatrix2x4 &"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.Int32Ty(false) // "int"
  vtys[13][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.Int32Ty(false) // "int"
  vtys[14][1] = qtrt.Int32Ty(false) // "GLuint"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = qtrt.Int32Ty(false) // "int"
  vtys[15][1] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.Int32Ty(false) // "int"
  // vtys[16][1] = reflect.TypeOf(QGenericMatrix<2, 2, float>{}) // "const QMatrix2x2 &"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = qtrt.ByteTy(true) // "const char *"
  vtys[17][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[18] = make(map[int32]reflect.Type)
  vtys[18][0] = qtrt.Int32Ty(false) // "int"
  vtys[18][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[18][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[19] = make(map[int32]reflect.Type)
  vtys[19][0] = qtrt.Int32Ty(false) // "int"
  vtys[19][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[19][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[19][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[20] = make(map[int32]reflect.Type)
  vtys[20][0] = qtrt.ByteTy(true) // "const char *"
  vtys[20][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[21] = make(map[int32]reflect.Type)
  vtys[21][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[21][1] = reflect.TypeOf(QGenericMatrix<3, 2, float>{}) // "const QMatrix3x2 &"
  vtys[22] = make(map[int32]reflect.Type)
  vtys[22][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[22][1] = reflect.TypeOf(QGenericMatrix<2, 3, float>{}) // "const QMatrix2x3 &"
  vtys[23] = make(map[int32]reflect.Type)
  vtys[23][0] = qtrt.Int32Ty(false) // "int"
  vtys[23][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[24] = make(map[int32]reflect.Type)
  vtys[24][0] = qtrt.ByteTy(true) // "const char *"
  vtys[24][1] = qtrt.VoidpTy() // "const GLfloat [4][4]"
  vtys[25] = make(map[int32]reflect.Type)
  vtys[25][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[25][1] = reflect.TypeOf(QGenericMatrix<2, 2, float>{}) // "const QMatrix2x2 &"
  vtys[26] = make(map[int32]reflect.Type)
  vtys[26][0] = qtrt.ByteTy(true) // "const char *"
  vtys[26][1] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[27] = make(map[int32]reflect.Type)
  vtys[27][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[27][1] = reflect.TypeOf(QGenericMatrix<4, 2, float>{}) // "const QMatrix4x2 &"
  vtys[28] = make(map[int32]reflect.Type)
  vtys[28][0] = qtrt.Int32Ty(false) // "int"
  vtys[28][1] = qtrt.VoidpTy() // "const GLfloat [2][2]"
  vtys[29] = make(map[int32]reflect.Type)
  vtys[29][0] = qtrt.ByteTy(true) // "const char *"
  vtys[29][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[30] = make(map[int32]reflect.Type)
  vtys[30][0] = qtrt.Int32Ty(false) // "int"
  vtys[30][1] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"
  vtys[31] = make(map[int32]reflect.Type)
  vtys[31][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[31][1] = reflect.TypeOf(QGenericMatrix<3, 4, float>{}) // "const QMatrix3x4 &"
  vtys[32] = make(map[int32]reflect.Type)
  vtys[32][0] = qtrt.ByteTy(true) // "const char *"
  vtys[32][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[33] = make(map[int32]reflect.Type)
  vtys[33][0] = qtrt.ByteTy(true) // "const char *"
  vtys[33][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[33][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[34] = make(map[int32]reflect.Type)
  vtys[34][0] = qtrt.Int32Ty(false) // "int"
  // vtys[34][1] = reflect.TypeOf(QGenericMatrix<3, 4, float>{}) // "const QMatrix3x4 &"
  vtys[35] = make(map[int32]reflect.Type)
  vtys[35][0] = qtrt.Int32Ty(false) // "int"
  vtys[35][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[36] = make(map[int32]reflect.Type)
  vtys[36][0] = qtrt.Int32Ty(false) // "int"
  // vtys[36][1] = reflect.TypeOf(QGenericMatrix<4, 2, float>{}) // "const QMatrix4x2 &"
  vtys[37] = make(map[int32]reflect.Type)
  vtys[37][0] = qtrt.Int32Ty(false) // "int"
  vtys[37][1] = qtrt.VoidpTy() // "const GLfloat [4][4]"
  vtys[38] = make(map[int32]reflect.Type)
  vtys[38][0] = qtrt.Int32Ty(false) // "int"
  // vtys[38][1] = reflect.TypeOf(QGenericMatrix<3, 3, float>{}) // "const QMatrix3x3 &"
  vtys[39] = make(map[int32]reflect.Type)
  vtys[39][0] = qtrt.ByteTy(true) // "const char *"
  vtys[39][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[40] = make(map[int32]reflect.Type)
  vtys[40][0] = qtrt.ByteTy(true) // "const char *"
  vtys[40][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[41] = make(map[int32]reflect.Type)
  vtys[41][0] = qtrt.Int32Ty(false) // "int"
  vtys[41][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[42] = make(map[int32]reflect.Type)
  vtys[42][0] = qtrt.ByteTy(true) // "const char *"
  vtys[42][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[42][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[42][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[43] = make(map[int32]reflect.Type)
  vtys[43][0] = qtrt.Int32Ty(false) // "int"
  vtys[43][1] = qtrt.VoidpTy() // "const GLfloat [3][3]"
  vtys[44] = make(map[int32]reflect.Type)
  vtys[44][0] = qtrt.ByteTy(true) // "const char *"
  vtys[44][1] = qtrt.Int32Ty(false) // "GLuint"
  vtys[45] = make(map[int32]reflect.Type)
  vtys[45][0] = qtrt.ByteTy(true) // "const char *"
  vtys[45][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[46] = make(map[int32]reflect.Type)
  vtys[46][0] = qtrt.Int32Ty(false) // "int"
  // vtys[46][1] = reflect.TypeOf(QGenericMatrix<4, 3, float>{}) // "const QMatrix4x3 &"
  vtys[47] = make(map[int32]reflect.Type)
  vtys[47][0] = qtrt.ByteTy(true) // "const char *"
  vtys[47][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[48] = make(map[int32]reflect.Type)
  vtys[48][0] = qtrt.Int32Ty(false) // "int"
  vtys[48][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[49] = make(map[int32]reflect.Type)
  vtys[49][0] = qtrt.Int32Ty(false) // "int"
  vtys[49][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[50] = make(map[int32]reflect.Type)
  vtys[50][0] = qtrt.ByteTy(true) // "const char *"
  vtys[50][1] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"
  vtys[51] = make(map[int32]reflect.Type)
  vtys[51][0] = qtrt.Int32Ty(false) // "int"
  vtys[51][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[51][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[51][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[51][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[52] = make(map[int32]reflect.Type)
  vtys[52][0] = qtrt.Int32Ty(false) // "int"
  // vtys[52][1] = reflect.TypeOf(QGenericMatrix<2, 4, float>{}) // "const QMatrix2x4 &"
  vtys[53] = make(map[int32]reflect.Type)
  vtys[53][0] = qtrt.ByteTy(true) // "const char *"
  vtys[53][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector3D
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QPoint
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcffff
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK10QTransform
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi2ELi3EfE
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA2_Kf
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi3ELi3EfE
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi3ELi2EfE
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi4ELi3EfE
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA3_Kf
  case 10:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector2D
  case 11:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEif
  case 12:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi2ELi4EfE
  case 13:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEii
  case 14:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEij
  case 15:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK5QSize
  case 16:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi2ELi2EfE
  case 17:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QPoint
  case 18:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiff
  case 19:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEifff
  case 20:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK7QPointF
  case 21:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi3ELi2EfE
  case 22:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi2ELi3EfE
  case 23:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QColor
  case 24:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA4_Kf
  case 25:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi2ELi2EfE
  case 26:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK5QSize
  case 27:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi4ELi2EfE
  case 28:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiPA2_Kf
  case 29:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QMatrix4x4
  case 30:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QSizeF
  case 31:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi3ELi4EfE
  case 32:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector4D
  case 33:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcff
  case 34:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi3ELi4EfE
  case 35:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK7QPointF
  case 36:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi4ELi2EfE
  case 37:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiPA4_Kf
  case 38:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi3ELi3EfE
  case 39:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QColor
  case 40:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QTransform
  case 41:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector3D
  case 42:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcfff
  case 43:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiPA3_Kf
  case 44:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcj
  case 45:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKci
  case 46:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi4ELi3EfE
  case 47:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcf
  case 48:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK10QMatrix4x4
  case 49:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector4D
  case 50:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QSizeF
  case 51:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiffff
  case 52:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi2ELi4EfE
  case 53:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector2D
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setUniformValue", args)
 }

}


func (this *QOpenGLShaderProgram) setAttributeValue(args ...interface{}) () {
  // setAttributeValue(const char *, const class QVector3D &)
  // setAttributeValue(int, GLfloat, GLfloat)
  // setAttributeValue(const char *, const class QVector2D &)
  // setAttributeValue(int, const class QColor &)
  // setAttributeValue(int, GLfloat, GLfloat, GLfloat)
  // setAttributeValue(const char *, const class QColor &)
  // setAttributeValue(const char *, GLfloat, GLfloat, GLfloat, GLfloat)
  // setAttributeValue(const char *, const GLfloat *, int, int)
  // setAttributeValue(int, const class QVector2D &)
  // setAttributeValue(int, GLfloat, GLfloat, GLfloat, GLfloat)
  // setAttributeValue(const char *, const class QVector4D &)
  // setAttributeValue(const char *, GLfloat)
  // setAttributeValue(int, const class QVector3D &)
  // setAttributeValue(const char *, GLfloat, GLfloat)
  // setAttributeValue(int, const class QVector4D &)
  // setAttributeValue(int, const GLfloat *, int, int)
  // setAttributeValue(int, GLfloat)
  // setAttributeValue(const char *, GLfloat, GLfloat, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[1][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[4][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[4][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(true) // "const char *"
  vtys[5][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  vtys[6][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.ByteTy(true) // "const char *"
  vtys[7][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[7][3] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int32Ty(false) // "int"
  vtys[8][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.Int32Ty(false) // "int"
  vtys[9][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[9][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[9][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[9][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.ByteTy(true) // "const char *"
  vtys[10][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.ByteTy(true) // "const char *"
  vtys[11][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.Int32Ty(false) // "int"
  vtys[12][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.ByteTy(true) // "const char *"
  vtys[13][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[13][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.Int32Ty(false) // "int"
  vtys[14][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = qtrt.Int32Ty(false) // "int"
  vtys[15][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[15][2] = qtrt.Int32Ty(false) // "int"
  vtys[15][3] = qtrt.Int32Ty(false) // "int"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.Int32Ty(false) // "int"
  vtys[16][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = qtrt.ByteTy(true) // "const char *"
  vtys[17][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17][3] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector3D
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiff
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector2D
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK6QColor
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEifff
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK6QColor
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcffff
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcPKfii
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector2D
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiffff
  case 10:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector4D
  case 11:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcf
  case 12:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector3D
  case 13:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcff
  case 14:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector4D
  case 15:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiPKfii
  case 16:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEif
  case 17:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcfff
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setAttributeValue", args)
 }

}


func (this *QOpenGLShaderProgram) setUniformValueArray(args ...interface{}) () {
  // setUniformValueArray(const char *, const class QVector3D *, int)
  // setUniformValueArray(const char *, const QMatrix3x3 *, int)
  // setUniformValueArray(int, const QMatrix2x3 *, int)
  // setUniformValueArray(int, const class QVector3D *, int)
  // setUniformValueArray(const char *, const GLfloat *, int, int)
  // setUniformValueArray(int, const QMatrix2x4 *, int)
  // setUniformValueArray(const char *, const QMatrix3x4 *, int)
  // setUniformValueArray(int, const class QMatrix4x4 *, int)
  // setUniformValueArray(int, const QMatrix3x4 *, int)
  // setUniformValueArray(const char *, const class QMatrix4x4 *, int)
  // setUniformValueArray(const char *, const QMatrix2x3 *, int)
  // setUniformValueArray(int, const GLuint *, int)
  // setUniformValueArray(const char *, const QMatrix4x2 *, int)
  // setUniformValueArray(const char *, const GLint *, int)
  // setUniformValueArray(const char *, const GLuint *, int)
  // setUniformValueArray(int, const QMatrix3x2 *, int)
  // setUniformValueArray(int, const class QVector4D *, int)
  // setUniformValueArray(const char *, const class QVector4D *, int)
  // setUniformValueArray(const char *, const QMatrix2x4 *, int)
  // setUniformValueArray(int, const QMatrix4x2 *, int)
  // setUniformValueArray(const char *, const class QVector2D *, int)
  // setUniformValueArray(int, const GLint *, int)
  // setUniformValueArray(const char *, const QMatrix4x3 *, int)
  // setUniformValueArray(const char *, const QMatrix3x2 *, int)
  // setUniformValueArray(const char *, const QMatrix2x2 *, int)
  // setUniformValueArray(int, const QMatrix2x2 *, int)
  // setUniformValueArray(int, const class QVector2D *, int)
  // setUniformValueArray(int, const QMatrix4x3 *, int)
  // setUniformValueArray(int, const QMatrix3x3 *, int)
  // setUniformValueArray(int, const GLfloat *, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[1][1] = reflect.TypeOf(QGenericMatrix<3, 3, float>{}) // "const QMatrix3x3 *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  // vtys[2][1] = reflect.TypeOf(QGenericMatrix<2, 3, float>{}) // "const QMatrix2x3 *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(true) // "const char *"
  vtys[4][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[4][3] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  // vtys[5][1] = reflect.TypeOf(QGenericMatrix<2, 4, float>{}) // "const QMatrix2x4 *"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[6][1] = reflect.TypeOf(QGenericMatrix<3, 4, float>{}) // "const QMatrix3x4 *"
  vtys[6][2] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 *"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int32Ty(false) // "int"
  // vtys[8][1] = reflect.TypeOf(QGenericMatrix<3, 4, float>{}) // "const QMatrix3x4 *"
  vtys[8][2] = qtrt.Int32Ty(false) // "int"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 *"
  vtys[9][2] = qtrt.Int32Ty(false) // "int"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[10][1] = reflect.TypeOf(QGenericMatrix<2, 3, float>{}) // "const QMatrix2x3 *"
  vtys[10][2] = qtrt.Int32Ty(false) // "int"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int32Ty(false) // "int"
  vtys[11][1] = qtrt.Int32Ty(true) // "const GLuint *"
  vtys[11][2] = qtrt.Int32Ty(false) // "int"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[12][1] = reflect.TypeOf(QGenericMatrix<4, 2, float>{}) // "const QMatrix4x2 *"
  vtys[12][2] = qtrt.Int32Ty(false) // "int"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.ByteTy(true) // "const char *"
  vtys[13][1] = qtrt.Int32Ty(true) // "const GLint *"
  vtys[13][2] = qtrt.Int32Ty(false) // "int"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.ByteTy(true) // "const char *"
  vtys[14][1] = qtrt.Int32Ty(true) // "const GLuint *"
  vtys[14][2] = qtrt.Int32Ty(false) // "int"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = qtrt.Int32Ty(false) // "int"
  // vtys[15][1] = reflect.TypeOf(QGenericMatrix<3, 2, float>{}) // "const QMatrix3x2 *"
  vtys[15][2] = qtrt.Int32Ty(false) // "int"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.Int32Ty(false) // "int"
  vtys[16][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[16][2] = qtrt.Int32Ty(false) // "int"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = qtrt.ByteTy(true) // "const char *"
  vtys[17][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[17][2] = qtrt.Int32Ty(false) // "int"
  vtys[18] = make(map[int32]reflect.Type)
  vtys[18][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[18][1] = reflect.TypeOf(QGenericMatrix<2, 4, float>{}) // "const QMatrix2x4 *"
  vtys[18][2] = qtrt.Int32Ty(false) // "int"
  vtys[19] = make(map[int32]reflect.Type)
  vtys[19][0] = qtrt.Int32Ty(false) // "int"
  // vtys[19][1] = reflect.TypeOf(QGenericMatrix<4, 2, float>{}) // "const QMatrix4x2 *"
  vtys[19][2] = qtrt.Int32Ty(false) // "int"
  vtys[20] = make(map[int32]reflect.Type)
  vtys[20][0] = qtrt.ByteTy(true) // "const char *"
  vtys[20][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[20][2] = qtrt.Int32Ty(false) // "int"
  vtys[21] = make(map[int32]reflect.Type)
  vtys[21][0] = qtrt.Int32Ty(false) // "int"
  vtys[21][1] = qtrt.Int32Ty(true) // "const GLint *"
  vtys[21][2] = qtrt.Int32Ty(false) // "int"
  vtys[22] = make(map[int32]reflect.Type)
  vtys[22][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[22][1] = reflect.TypeOf(QGenericMatrix<4, 3, float>{}) // "const QMatrix4x3 *"
  vtys[22][2] = qtrt.Int32Ty(false) // "int"
  vtys[23] = make(map[int32]reflect.Type)
  vtys[23][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[23][1] = reflect.TypeOf(QGenericMatrix<3, 2, float>{}) // "const QMatrix3x2 *"
  vtys[23][2] = qtrt.Int32Ty(false) // "int"
  vtys[24] = make(map[int32]reflect.Type)
  vtys[24][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[24][1] = reflect.TypeOf(QGenericMatrix<2, 2, float>{}) // "const QMatrix2x2 *"
  vtys[24][2] = qtrt.Int32Ty(false) // "int"
  vtys[25] = make(map[int32]reflect.Type)
  vtys[25][0] = qtrt.Int32Ty(false) // "int"
  // vtys[25][1] = reflect.TypeOf(QGenericMatrix<2, 2, float>{}) // "const QMatrix2x2 *"
  vtys[25][2] = qtrt.Int32Ty(false) // "int"
  vtys[26] = make(map[int32]reflect.Type)
  vtys[26][0] = qtrt.Int32Ty(false) // "int"
  vtys[26][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[26][2] = qtrt.Int32Ty(false) // "int"
  vtys[27] = make(map[int32]reflect.Type)
  vtys[27][0] = qtrt.Int32Ty(false) // "int"
  // vtys[27][1] = reflect.TypeOf(QGenericMatrix<4, 3, float>{}) // "const QMatrix4x3 *"
  vtys[27][2] = qtrt.Int32Ty(false) // "int"
  vtys[28] = make(map[int32]reflect.Type)
  vtys[28][0] = qtrt.Int32Ty(false) // "int"
  // vtys[28][1] = reflect.TypeOf(QGenericMatrix<3, 3, float>{}) // "const QMatrix3x3 *"
  vtys[28][2] = qtrt.Int32Ty(false) // "int"
  vtys[29] = make(map[int32]reflect.Type)
  vtys[29][0] = qtrt.Int32Ty(false) // "int"
  vtys[29][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[29][2] = qtrt.Int32Ty(false) // "int"
  vtys[29][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector3Di
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi3ELi3EfEi
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi2ELi3EfEi
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector3Di
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKfii
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi2ELi4EfEi
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi3ELi4EfEi
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK10QMatrix4x4i
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi3ELi4EfEi
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK10QMatrix4x4i
  case 10:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi2ELi3EfEi
  case 11:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKji
  case 12:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi4ELi2EfEi
  case 13:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKii
  case 14:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKji
  case 15:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi3ELi2EfEi
  case 16:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector4Di
  case 17:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector4Di
  case 18:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi2ELi4EfEi
  case 19:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi4ELi2EfEi
  case 20:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector2Di
  case 21:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKii
  case 22:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi4ELi3EfEi
  case 23:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi3ELi2EfEi
  case 24:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi2ELi2EfEi
  case 25:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi2ELi2EfEi
  case 26:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector2Di
  case 27:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi4ELi3EfEi
  case 28:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi3ELi3EfEi
  case 29:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKfii
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setUniformValueArray", args)
 }

}


func (this *QOpenGLShaderProgram) setAttributeBuffer(args ...interface{}) () {
  // setAttributeBuffer(int, GLenum, int, int, int)
  // setAttributeBuffer(const char *, GLenum, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram18setAttributeBufferEijiii
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram18setAttributeBufferEPKcjiii
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setAttributeBuffer", args)
 }

}


func (this *QOpenGLShaderProgram) hasOpenGLShaderPrograms_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "hasOpenGLShaderPrograms", args)
 }

}


func (this *QOpenGLShaderProgram) setPatchVertexCount(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setPatchVertexCount", args)
 }

}


func (this *QOpenGLShaderProgram) setAttributeArray(args ...interface{}) () {
  // setAttributeArray(const char *, const GLfloat *, int, int)
  // setAttributeArray(int, GLenum, const void *, int, int)
  // setAttributeArray(int, const class QVector2D *, int)
  // setAttributeArray(const char *, const class QVector4D *, int)
  // setAttributeArray(int, const class QVector3D *, int)
  // setAttributeArray(int, const class QVector4D *, int)
  // setAttributeArray(const char *, GLenum, const void *, int, int)
  // setAttributeArray(int, const GLfloat *, int, int)
  // setAttributeArray(const char *, const class QVector3D *, int)
  // setAttributeArray(const char *, const class QVector2D *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[1][2] = qtrt.VoidpTy() // "const void *"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  vtys[6][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[6][2] = qtrt.VoidpTy() // "const void *"
  vtys[6][3] = qtrt.Int32Ty(false) // "int"
  vtys[6][4] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[7][3] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.ByteTy(true) // "const char *"
  vtys[8][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[8][2] = qtrt.Int32Ty(false) // "int"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[9][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPKfii
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEijPKvii
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector2Di
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector4Di
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector3Di
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector4Di
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcjPKvii
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPKfii
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector3Di
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector2Di
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setAttributeArray", args)
 }

}


func (this *QOpenGLShaderProgram) bindAttributeLocation(args ...interface{}) () {
  // bindAttributeLocation(const class QByteArray &, int)
  // bindAttributeLocation(const char *, int)
  // bindAttributeLocation(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram21bindAttributeLocationERK10QByteArrayi
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram21bindAttributeLocationEPKci
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram21bindAttributeLocationERK7QStringi
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "bindAttributeLocation", args)
 }

}


func (this *QOpenGLShaderProgram) bind(args ...interface{}) () {
  // bind()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram4bindEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "bind", args)
 }

}


func (this *QOpenGLShaderProgram) enableAttributeArray(args ...interface{}) () {
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
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram20enableAttributeArrayEPKc
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "enableAttributeArray", args)
 }

}


func (this *QOpenGLShaderProgram) addShader(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "addShader", args)
 }

}


func (this *QOpenGLShaderProgram) attributeLocation(args ...interface{}) () {
  // attributeLocation(const class QString &)
  // attributeLocation(const char *)
  // attributeLocation(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram17attributeLocationERK7QString
  case 1:
    // invoke: _ZNK20QOpenGLShaderProgram17attributeLocationEPKc
  case 2:
    // invoke: _ZNK20QOpenGLShaderProgram17attributeLocationERK10QByteArray
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "attributeLocation", args)
 }

}


func (this *QOpenGLShaderProgram) programId(args ...interface{}) () {
  // programId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram9programIdEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "programId", args)
 }

}


func (this *QOpenGLShaderProgram) disableAttributeArray(args ...interface{}) () {
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
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram21disableAttributeArrayEPKc
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "disableAttributeArray", args)
 }

}


func (this *QOpenGLShaderProgram) uniformLocation(args ...interface{}) () {
  // uniformLocation(const class QString &)
  // uniformLocation(const char *)
  // uniformLocation(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram15uniformLocationERK7QString
  case 1:
    // invoke: _ZNK20QOpenGLShaderProgram15uniformLocationEPKc
  case 2:
    // invoke: _ZNK20QOpenGLShaderProgram15uniformLocationERK10QByteArray
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "uniformLocation", args)
 }

}


func (this *QOpenGLShaderProgram) FreeQOpenGLShaderProgram(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "~QOpenGLShaderProgram", args)
 }

}


func (this *QOpenGLShaderProgram) defaultInnerTessellationLevels(args ...interface{}) () {
  // defaultInnerTessellationLevels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram30defaultInnerTessellationLevelsEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "defaultInnerTessellationLevels", args)
 }

}


func (this *QOpenGLShaderProgram) link(args ...interface{}) () {
  // link()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram4linkEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "link", args)
 }

}


func (this *QOpenGLShaderProgram) shaders(args ...interface{}) () {
  // shaders()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram7shadersEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "shaders", args)
 }

}


func (this *QOpenGLShaderProgram) defaultOuterTessellationLevels(args ...interface{}) () {
  // defaultOuterTessellationLevels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram30defaultOuterTessellationLevelsEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "defaultOuterTessellationLevels", args)
 }

}


func (this *QOpenGLShaderProgram) log(args ...interface{}) () {
  // log()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram3logEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "log", args)
 }

}


func (this *QOpenGLShaderProgram) patchVertexCount(args ...interface{}) () {
  // patchVertexCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram16patchVertexCountEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "patchVertexCount", args)
 }

}


func (this *QOpenGLShaderProgram) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram6createEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "create", args)
 }

}


func (this *QOpenGLShaderProgram) removeShader(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "removeShader", args)
 }

}


func NewQOpenGLShaderProgram(args ...interface{})() {
}


func (this *QOpenGLShaderProgram) removeAllShaders(args ...interface{}) () {
  // removeAllShaders()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram16removeAllShadersEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "removeAllShaders", args)
 }

}


func (this *QOpenGLShaderProgram) maxGeometryOutputVertices(args ...interface{}) () {
  // maxGeometryOutputVertices()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram25maxGeometryOutputVerticesEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "maxGeometryOutputVertices", args)
 }

}


func (this *QOpenGLShaderProgram) release(args ...interface{}) () {
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram7releaseEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "release", args)
 }

}


func (this *QOpenGLShaderProgram) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram10metaObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "metaObject", args)
 }

}

// <= body block end

