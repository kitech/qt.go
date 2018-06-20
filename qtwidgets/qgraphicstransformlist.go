package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicstransform.h
// #include <qgraphicstransform.h>
// #include <QtWidgets>

//  header block end

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
import "github.com/kitech/qt.go/qtgui"

//  ext block end

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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end

//  body block begin
type QGraphicsTransformList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QGraphicsTransformList) Operator_equal_0() *QGraphicsTransformList {
	// QGraphicsTransformList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QGraphicsTransformList) Operator_equal_1() *QGraphicsTransformList {
	// QGraphicsTransformList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QGraphicsTransformList) Swap_0() {
	// QGraphicsTransformList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QGraphicsTransformList) Operator_equal_equal_0() bool {
	// QGraphicsTransformList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QGraphicsTransformList) Operator_not_equal_0() bool {
	// QGraphicsTransformList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QGraphicsTransformList) Size_0() int {
	// QGraphicsTransformList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QGraphicsTransformList) Detach_0() {
	// QGraphicsTransformList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QGraphicsTransformList) DetachShared_0() {
	// QGraphicsTransformList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QGraphicsTransformList) IsDetached_0() bool {
	// QGraphicsTransformList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QGraphicsTransformList) SetSharable_0() {
	// QGraphicsTransformList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QGraphicsTransformList) IsSharedWith_0() bool {
	// QGraphicsTransformList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QGraphicsTransformList) IsEmpty_0() bool {
	// QGraphicsTransformList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QGraphicsTransformList) Clear_0() {
	// QGraphicsTransformList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QGraphicsTransformList) At_0() *QGraphicsTransform {
	// QGraphicsTransformList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// const T & operator[](int)
func (this *QGraphicsTransformList) Operator_get_index_0() *QGraphicsTransform {
	// QGraphicsTransformList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// T & operator[](int)
func (this *QGraphicsTransformList) Operator_get_index_1() *QGraphicsTransform {
	// QGraphicsTransformList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// void reserve(int)
func (this *QGraphicsTransformList) Reserve_0() {
	// QGraphicsTransformList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QGraphicsTransformList) Append_0() {
	// QGraphicsTransformList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QGraphicsTransformList) Append_1() {
	// QGraphicsTransformList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QGraphicsTransformList) Prepend_0() {
	// QGraphicsTransformList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QGraphicsTransformList) Insert_0() {
	// QGraphicsTransformList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QGraphicsTransformList) Replace_0() {
	// QGraphicsTransformList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QGraphicsTransformList) RemoveAt_0() {
	// QGraphicsTransformList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QGraphicsTransformList) RemoveAll_0() int {
	// QGraphicsTransformList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QGraphicsTransformList) RemoveOne_0() bool {
	// QGraphicsTransformList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QGraphicsTransformList) TakeAt_0() *QGraphicsTransform {
	// QGraphicsTransformList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// T takeFirst()
func (this *QGraphicsTransformList) TakeFirst_0() *QGraphicsTransform {
	// QGraphicsTransformList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// T takeLast()
func (this *QGraphicsTransformList) TakeLast_0() *QGraphicsTransform {
	// QGraphicsTransformList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// void move(int, int)
func (this *QGraphicsTransformList) Move_0() {
	// QGraphicsTransformList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QGraphicsTransformList) Swap_1() {
	// QGraphicsTransformList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QGraphicsTransformList) IndexOf_0() int {
	// QGraphicsTransformList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QGraphicsTransformList) LastIndexOf_0() int {
	// QGraphicsTransformList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QGraphicsTransformList) Contains_0() bool {
	// QGraphicsTransformList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QGraphicsTransformList) Count_0() int {
	// QGraphicsTransformList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QGraphicsTransformList) Begin_0() {
	// QGraphicsTransformList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QGraphicsTransformList) Begin_1() {
	// QGraphicsTransformList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QGraphicsTransformList) Cbegin_0() {
	// QGraphicsTransformList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QGraphicsTransformList) ConstBegin_0() {
	// QGraphicsTransformList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QGraphicsTransformList) End_0() {
	// QGraphicsTransformList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QGraphicsTransformList) End_1() {
	// QGraphicsTransformList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QGraphicsTransformList) Cend_0() {
	// QGraphicsTransformList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QGraphicsTransformList) ConstEnd_0() {
	// QGraphicsTransformList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QGraphicsTransformList) Rbegin_0() {
	// QGraphicsTransformList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QGraphicsTransformList) Rend_0() {
	// QGraphicsTransformList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QGraphicsTransformList) Rbegin_1() {
	// QGraphicsTransformList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QGraphicsTransformList) Rend_1() {
	// QGraphicsTransformList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QGraphicsTransformList) Crbegin_0() {
	// QGraphicsTransformList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QGraphicsTransformList) Crend_0() {
	// QGraphicsTransformList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QGraphicsTransformList) Insert_1() {
	// QGraphicsTransformList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QGraphicsTransformList) Erase_0() {
	// QGraphicsTransformList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QGraphicsTransformList) Erase_1() {
	// QGraphicsTransformList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QGraphicsTransformList) Count_1() int {
	// QGraphicsTransformList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QGraphicsTransformList) Length_0() int {
	// QGraphicsTransformList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QGraphicsTransformList) First_0() *QGraphicsTransform {
	// QGraphicsTransformList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// const T & constFirst()
func (this *QGraphicsTransformList) ConstFirst_0() *QGraphicsTransform {
	// QGraphicsTransformList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// const T & first()
func (this *QGraphicsTransformList) First_1() *QGraphicsTransform {
	// QGraphicsTransformList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// T & last()
func (this *QGraphicsTransformList) Last_0() *QGraphicsTransform {
	// QGraphicsTransformList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// const T & last()
func (this *QGraphicsTransformList) Last_1() *QGraphicsTransform {
	// QGraphicsTransformList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// const T & constLast()
func (this *QGraphicsTransformList) ConstLast_0() *QGraphicsTransform {
	// QGraphicsTransformList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// void removeFirst()
func (this *QGraphicsTransformList) RemoveFirst_0() {
	// QGraphicsTransformList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QGraphicsTransformList) RemoveLast_0() {
	// QGraphicsTransformList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QGraphicsTransformList) StartsWith_0() bool {
	// QGraphicsTransformList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QGraphicsTransformList) EndsWith_0() bool {
	// QGraphicsTransformList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QGraphicsTransformList) Mid_0() *QGraphicsTransformList {
	// QGraphicsTransformList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QGraphicsTransformList) Value_0() *QGraphicsTransform {
	// QGraphicsTransformList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// T value(int, const T &)
func (this *QGraphicsTransformList) Value_1() *QGraphicsTransform {
	// QGraphicsTransformList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// void push_back(const T &)
func (this *QGraphicsTransformList) Push_back_0() {
	// QGraphicsTransformList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QGraphicsTransformList) Push_front_0() {
	// QGraphicsTransformList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QGraphicsTransformList) Front_0() *QGraphicsTransform {
	// QGraphicsTransformList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// const T & front()
func (this *QGraphicsTransformList) Front_1() *QGraphicsTransform {
	// QGraphicsTransformList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// T & back()
func (this *QGraphicsTransformList) Back_0() *QGraphicsTransform {
	// QGraphicsTransformList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// const T & back()
func (this *QGraphicsTransformList) Back_1() *QGraphicsTransform {
	// QGraphicsTransformList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGraphicsTransform{}
}

// void pop_front()
func (this *QGraphicsTransformList) Pop_front_0() {
	// QGraphicsTransformList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QGraphicsTransformList) Pop_back_0() {
	// QGraphicsTransformList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QGraphicsTransformList) Empty_0() bool {
	// QGraphicsTransformList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QGraphicsTransformList) Operator_add_equal_0() *QGraphicsTransformList {
	// QGraphicsTransformList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QGraphicsTransformList) Operator_add_0() *QGraphicsTransformList {
	// QGraphicsTransformList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QGraphicsTransformList) Operator_add_equal_1() *QGraphicsTransformList {
	// QGraphicsTransformList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QGraphicsTransformList) Operator_left_shift_0() *QGraphicsTransformList {
	// QGraphicsTransformList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QGraphicsTransformList) Operator_left_shift_1() *QGraphicsTransformList {
	// QGraphicsTransformList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QGraphicsTransformList) ToVector_0() {
	// QGraphicsTransformList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QGraphicsTransformList) ToSet_0() {
	// QGraphicsTransformList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QGraphicsTransformList) FromVector_0() *QGraphicsTransformList {
	// QGraphicsTransformList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QGraphicsTransformList) FromSet_0() *QGraphicsTransformList {
	// QGraphicsTransformList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QGraphicsTransformList) FromStdList_0() *QGraphicsTransformList {
	// QGraphicsTransformList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QGraphicsTransformList) ToStdList_0() {
	// QGraphicsTransformList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QGraphicsTransformList) Detach_helper_grow_0() {
	// QGraphicsTransformList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QGraphicsTransformList) Detach_helper_0() {
	// QGraphicsTransformList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QGraphicsTransformList) Detach_helper_1() {
	// QGraphicsTransformList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QGraphicsTransformList) Dealloc_0() {
	// QGraphicsTransformList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QGraphicsTransformList) Node_construct_0() {
	// QGraphicsTransformList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QGraphicsTransformList) Node_destruct_0() {
	// QGraphicsTransformList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QGraphicsTransformList) Node_copy_0() {
	// QGraphicsTransformList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QGraphicsTransformList) Node_destruct_1() {
	// QGraphicsTransformList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QGraphicsTransformList) IsValidIterator_0() bool {
	// QGraphicsTransformList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QGraphicsTransformList) Op_eq_impl_0() bool {
	// QGraphicsTransformList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QGraphicsTransformList) Op_eq_impl_1() bool {
	// QGraphicsTransformList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGraphicsTransformList) Contains_impl_0() bool {
	// QGraphicsTransformList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGraphicsTransformList) Contains_impl_1() bool {
	// QGraphicsTransformList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGraphicsTransformList) Count_impl_0() int {
	// QGraphicsTransformList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGraphicsTransformList) Count_impl_1() int {
	// QGraphicsTransformList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGraphicsTransformList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
