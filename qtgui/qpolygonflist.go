package qtgui

// /usr/include/qt/QtGui/qpolygon.h
// #include <qpolygon.h>
// #include <QtGui>

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
}

//  keep block end

//  body block begin
type QPolygonFList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QPolygonFList) Operator_equal_0() *QPolygonFList {
	// QPolygonFList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QPolygonFList) Operator_equal_1() *QPolygonFList {
	// QPolygonFList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QPolygonFList) Swap_0() {
	// QPolygonFList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QPolygonFList) Operator_equal_equal_0() bool {
	// QPolygonFList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QPolygonFList) Operator_not_equal_0() bool {
	// QPolygonFList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QPolygonFList) Size_0() int {
	// QPolygonFList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QPolygonFList) Detach_0() {
	// QPolygonFList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QPolygonFList) DetachShared_0() {
	// QPolygonFList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QPolygonFList) IsDetached_0() bool {
	// QPolygonFList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QPolygonFList) SetSharable_0() {
	// QPolygonFList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QPolygonFList) IsSharedWith_0() bool {
	// QPolygonFList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QPolygonFList) IsEmpty_0() bool {
	// QPolygonFList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QPolygonFList) Clear_0() {
	// QPolygonFList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QPolygonFList) At_0() *QPolygonF {
	// QPolygonFList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// const T & operator[](int)
func (this *QPolygonFList) Operator_get_index_0() *QPolygonF {
	// QPolygonFList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// T & operator[](int)
func (this *QPolygonFList) Operator_get_index_1() *QPolygonF {
	// QPolygonFList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// void reserve(int)
func (this *QPolygonFList) Reserve_0() {
	// QPolygonFList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QPolygonFList) Append_0() {
	// QPolygonFList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QPolygonFList) Append_1() {
	// QPolygonFList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QPolygonFList) Prepend_0() {
	// QPolygonFList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QPolygonFList) Insert_0() {
	// QPolygonFList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QPolygonFList) Replace_0() {
	// QPolygonFList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QPolygonFList) RemoveAt_0() {
	// QPolygonFList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QPolygonFList) RemoveAll_0() int {
	// QPolygonFList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QPolygonFList) RemoveOne_0() bool {
	// QPolygonFList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QPolygonFList) TakeAt_0() *QPolygonF {
	// QPolygonFList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// T takeFirst()
func (this *QPolygonFList) TakeFirst_0() *QPolygonF {
	// QPolygonFList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// T takeLast()
func (this *QPolygonFList) TakeLast_0() *QPolygonF {
	// QPolygonFList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// void move(int, int)
func (this *QPolygonFList) Move_0() {
	// QPolygonFList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QPolygonFList) Swap_1() {
	// QPolygonFList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QPolygonFList) IndexOf_0() int {
	// QPolygonFList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QPolygonFList) LastIndexOf_0() int {
	// QPolygonFList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QPolygonFList) Contains_0() bool {
	// QPolygonFList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QPolygonFList) Count_0() int {
	// QPolygonFList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QPolygonFList) Begin_0() {
	// QPolygonFList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QPolygonFList) Begin_1() {
	// QPolygonFList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QPolygonFList) Cbegin_0() {
	// QPolygonFList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QPolygonFList) ConstBegin_0() {
	// QPolygonFList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QPolygonFList) End_0() {
	// QPolygonFList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QPolygonFList) End_1() {
	// QPolygonFList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QPolygonFList) Cend_0() {
	// QPolygonFList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QPolygonFList) ConstEnd_0() {
	// QPolygonFList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QPolygonFList) Rbegin_0() {
	// QPolygonFList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QPolygonFList) Rend_0() {
	// QPolygonFList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QPolygonFList) Rbegin_1() {
	// QPolygonFList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QPolygonFList) Rend_1() {
	// QPolygonFList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QPolygonFList) Crbegin_0() {
	// QPolygonFList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QPolygonFList) Crend_0() {
	// QPolygonFList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QPolygonFList) Insert_1() {
	// QPolygonFList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QPolygonFList) Erase_0() {
	// QPolygonFList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QPolygonFList) Erase_1() {
	// QPolygonFList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QPolygonFList) Count_1() int {
	// QPolygonFList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QPolygonFList) Length_0() int {
	// QPolygonFList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QPolygonFList) First_0() *QPolygonF {
	// QPolygonFList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// const T & constFirst()
func (this *QPolygonFList) ConstFirst_0() *QPolygonF {
	// QPolygonFList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// const T & first()
func (this *QPolygonFList) First_1() *QPolygonF {
	// QPolygonFList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// T & last()
func (this *QPolygonFList) Last_0() *QPolygonF {
	// QPolygonFList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// const T & last()
func (this *QPolygonFList) Last_1() *QPolygonF {
	// QPolygonFList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// const T & constLast()
func (this *QPolygonFList) ConstLast_0() *QPolygonF {
	// QPolygonFList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// void removeFirst()
func (this *QPolygonFList) RemoveFirst_0() {
	// QPolygonFList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QPolygonFList) RemoveLast_0() {
	// QPolygonFList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QPolygonFList) StartsWith_0() bool {
	// QPolygonFList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QPolygonFList) EndsWith_0() bool {
	// QPolygonFList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QPolygonFList) Mid_0() *QPolygonFList {
	// QPolygonFList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QPolygonFList) Value_0() *QPolygonF {
	// QPolygonFList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// T value(int, const T &)
func (this *QPolygonFList) Value_1() *QPolygonF {
	// QPolygonFList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// void push_back(const T &)
func (this *QPolygonFList) Push_back_0() {
	// QPolygonFList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QPolygonFList) Push_front_0() {
	// QPolygonFList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QPolygonFList) Front_0() *QPolygonF {
	// QPolygonFList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// const T & front()
func (this *QPolygonFList) Front_1() *QPolygonF {
	// QPolygonFList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// T & back()
func (this *QPolygonFList) Back_0() *QPolygonF {
	// QPolygonFList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// const T & back()
func (this *QPolygonFList) Back_1() *QPolygonF {
	// QPolygonFList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QPolygonF{}
}

// void pop_front()
func (this *QPolygonFList) Pop_front_0() {
	// QPolygonFList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QPolygonFList) Pop_back_0() {
	// QPolygonFList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QPolygonFList) Empty_0() bool {
	// QPolygonFList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QPolygonFList) Operator_add_equal_0() *QPolygonFList {
	// QPolygonFList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QPolygonFList) Operator_add_0() *QPolygonFList {
	// QPolygonFList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QPolygonFList) Operator_add_equal_1() *QPolygonFList {
	// QPolygonFList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QPolygonFList) Operator_left_shift_0() *QPolygonFList {
	// QPolygonFList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QPolygonFList) Operator_left_shift_1() *QPolygonFList {
	// QPolygonFList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QPolygonFList) ToVector_0() {
	// QPolygonFList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QPolygonFList) ToSet_0() {
	// QPolygonFList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QPolygonFList) FromVector_0() *QPolygonFList {
	// QPolygonFList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QPolygonFList) FromSet_0() *QPolygonFList {
	// QPolygonFList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QPolygonFList) FromStdList_0() *QPolygonFList {
	// QPolygonFList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QPolygonFList) ToStdList_0() {
	// QPolygonFList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QPolygonFList) Detach_helper_grow_0() {
	// QPolygonFList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QPolygonFList) Detach_helper_0() {
	// QPolygonFList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QPolygonFList) Detach_helper_1() {
	// QPolygonFList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QPolygonFList) Dealloc_0() {
	// QPolygonFList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QPolygonFList) Node_construct_0() {
	// QPolygonFList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QPolygonFList) Node_destruct_0() {
	// QPolygonFList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QPolygonFList) Node_copy_0() {
	// QPolygonFList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QPolygonFList) Node_destruct_1() {
	// QPolygonFList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QPolygonFList) IsValidIterator_0() bool {
	// QPolygonFList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QPolygonFList) Op_eq_impl_0() bool {
	// QPolygonFList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QPolygonFList) Op_eq_impl_1() bool {
	// QPolygonFList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QPolygonFList) Contains_impl_0() bool {
	// QPolygonFList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QPolygonFList) Contains_impl_1() bool {
	// QPolygonFList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QPolygonFList) Count_impl_0() int {
	// QPolygonFList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QPolygonFList) Count_impl_1() int {
	// QPolygonFList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QPolygonFList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
