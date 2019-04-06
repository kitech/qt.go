package qtgui

// /usr/include/qt/QtGui/qvalidator.h
// #include <qvalidator.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QRegularExpressionValidator struct {
	*QValidator
}
type QRegularExpressionValidator_ITF interface {
	QValidator_ITF
	QRegularExpressionValidator_PTR() *QRegularExpressionValidator
}

func (ptr *QRegularExpressionValidator) QRegularExpressionValidator_PTR() *QRegularExpressionValidator {
	return ptr
}

func (this *QRegularExpressionValidator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QValidator.GetCthis()
	}
}
func (this *QRegularExpressionValidator) SetCthis(cthis unsafe.Pointer) {
	this.QValidator = NewQValidatorFromPointer(cthis)
}
func NewQRegularExpressionValidatorFromPointer(cthis unsafe.Pointer) *QRegularExpressionValidator {
	bcthis0 := NewQValidatorFromPointer(cthis)
	return &QRegularExpressionValidator{bcthis0}
}
func (*QRegularExpressionValidator) NewFromPointer(cthis unsafe.Pointer) *QRegularExpressionValidator {
	return NewQRegularExpressionValidatorFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_10977() {
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
