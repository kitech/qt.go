package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h
// #include <qandroidjniobject.h>
// #include <QtAndroidExtras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QAndroidJniObject struct {
	*qtrt.CObject
}
type QAndroidJniObject_ITF interface {
	QAndroidJniObject_PTR() *QAndroidJniObject
}

func (ptr *QAndroidJniObject) QAndroidJniObject_PTR() *QAndroidJniObject { return ptr }

func (this *QAndroidJniObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidJniObject) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidJniObjectFromPointer(cthis unsafe.Pointer) *QAndroidJniObject {
	return &QAndroidJniObject{&qtrt.CObject{cthis}}
}
func (*QAndroidJniObject) NewFromPointer(cthis unsafe.Pointer) *QAndroidJniObject {
	return NewQAndroidJniObjectFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniObject()

/*
Constructs an invalid QAndroidJniObject.

See also isValid().
*/
func (*QAndroidJniObject) NewForInherit() *QAndroidJniObject {
	return NewQAndroidJniObject()
}
func NewQAndroidJniObject() *QAndroidJniObject {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObjectC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniObject)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:56
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniObject(const char *)

/*
Constructs an invalid QAndroidJniObject.

See also isValid().
*/
func (*QAndroidJniObject) NewForInherit_1(className string) *QAndroidJniObject {
	return NewQAndroidJniObject_1(className)
}
func NewQAndroidJniObject_1(className string) *QAndroidJniObject {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObjectC2EPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniObject)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAndroidJniObject()

/*

 */
func DeleteQAndroidJniObject(this *QAndroidJniObject) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObjectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [8] jobject object() const

/*
Returns the object held by the QAndroidJniObject as type T.


  QAndroidJniObject string = QAndroidJniObject::fromString("Hello, JNI");
  jstring jstring = string.object<jstring>();



Note: The returned object is still owned by the QAndroidJniObject. If you want to keep the object valid you should create a new QAndroidJniObject or make a new global reference to the object and free it yourself.


  void functionScope()
  {
      QString helloString("Hello");
      jstring myJString = 0;
      {
          QAndroidJniObject string = QAndroidJniObject::fromString(helloString);
          myJString = string.object<jstring>();
      }

     // Ops! myJString is no longer valid.
  }



Note: Since Qt 5.3 this function can be used without a template type, if the returned type is a jobject.


  jobject object = jniObject.object();
*/
func (this *QAndroidJniObject) Object() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAndroidJniObject6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:108
// index:0
// Public Visibility=Default Availability=Available
// [16] QAndroidJniObject getObjectField(const char *, const char *) const

/*
Retrieves the object of field fieldName.


  QAndroidJniObject field = jniObject.getObjectField<jstring>("FIELD_NAME");
*/
func (this *QAndroidJniObject) GetObjectField(fieldName string, sig string) *QAndroidJniObject /*123*/ {
	var convArg0 = qtrt.CString(fieldName)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(sig)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAndroidJniObject14getObjectFieldEPKcS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [16] QAndroidJniObject getStaticObjectField(const char *, const char *, const char *)

/*
Retrieves the object from the field fieldName on the class className.


  QAndroidJniObject jobj = QAndroidJniObject::getStaticObjectField<jstring>("class/with/Fields", "FIELD_NAME");
*/
func (this *QAndroidJniObject) GetStaticObjectField(className string, fieldName string, sig string) *QAndroidJniObject /*123*/ {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(fieldName)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(sig)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObject20getStaticObjectFieldEPKcS1_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}
func QAndroidJniObject_GetStaticObjectField(className string, fieldName string, sig string) *QAndroidJniObject /*123*/ {
	var nilthis *QAndroidJniObject
	rv := nilthis.GetStaticObjectField(className, fieldName, sig)
	return rv
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:158
// index:0
// Public static Visibility=Default Availability=Available
// [16] QAndroidJniObject fromString(const QString &)

/*
Creates a Java string from the QString string and returns a QAndroidJniObject holding that string.


  ...
  QString myQString = "QString";
  QAndroidJniObject myJavaString = QAndroidJniObject::fromString(myQString);
  ...



See also toString().
*/
func (this *QAndroidJniObject) FromString(string string) *QAndroidJniObject /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObject10fromStringERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}
func QAndroidJniObject_FromString(string string) *QAndroidJniObject /*123*/ {
	var nilthis *QAndroidJniObject
	rv := nilthis.FromString(string)
	return rv
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:159
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns a QString with a string representation of the java object. Calling this function on a Java String object is a convenient way of getting the actual string data.


  QAndroidJniObject string = ...; //  "Hello Java"
  QString qstring = string.toString(); // "Hello Java"



See also fromString().
*/
func (this *QAndroidJniObject) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAndroidJniObject8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:161
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isClassAvailable(const char *)

/*
Returns true if the Java class className is available.


  ...
  if (QAndroidJniObject::isClassAvailable("java/lang/String")) {
     ...
  }
  ...
*/
func (this *QAndroidJniObject) IsClassAvailable(className string) bool {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAndroidJniObject16isClassAvailableEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QAndroidJniObject_IsClassAvailable(className string) bool {
	var nilthis *QAndroidJniObject
	rv := nilthis.IsClassAvailable(className)
	return rv
}

// /usr/include/qt/QtAndroidExtras/qandroidjniobject.h:162
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this instance holds a valid Java object.


  ...
  QAndroidJniObject qjniObject;                        ==> isValid() == false
  QAndroidJniObject qjniObject(0)                      ==> isValid() == false
  QAndroidJniObject qjniObject("could/not/find/Class") ==> isValid() == false
  ...
*/
func (this *QAndroidJniObject) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QAndroidJniObject7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
