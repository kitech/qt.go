package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
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
type QTextFrameList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QTextFrameList) Operator_equal_0() *QTextFrameList {
	// QTextFrameList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QTextFrameList) Operator_equal_1() *QTextFrameList {
	// QTextFrameList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QTextFrameList) Swap_0() {
	// QTextFrameList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QTextFrameList) Operator_equal_equal_0() bool {
	// QTextFrameList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QTextFrameList) Operator_not_equal_0() bool {
	// QTextFrameList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QTextFrameList) Size_0() int {
	// QTextFrameList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QTextFrameList) Detach_0() {
	// QTextFrameList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QTextFrameList) DetachShared_0() {
	// QTextFrameList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QTextFrameList) IsDetached_0() bool {
	// QTextFrameList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QTextFrameList) SetSharable_0() {
	// QTextFrameList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QTextFrameList) IsSharedWith_0() bool {
	// QTextFrameList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QTextFrameList) IsEmpty_0() bool {
	// QTextFrameList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QTextFrameList) Clear_0() {
	// QTextFrameList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QTextFrameList) At_0() *QTextFrame {
	// QTextFrameList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// const T & operator[](int)
func (this *QTextFrameList) Operator_get_index_0() *QTextFrame {
	// QTextFrameList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// T & operator[](int)
func (this *QTextFrameList) Operator_get_index_1() *QTextFrame {
	// QTextFrameList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// void reserve(int)
func (this *QTextFrameList) Reserve_0() {
	// QTextFrameList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QTextFrameList) Append_0() {
	// QTextFrameList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QTextFrameList) Append_1() {
	// QTextFrameList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QTextFrameList) Prepend_0() {
	// QTextFrameList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QTextFrameList) Insert_0() {
	// QTextFrameList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QTextFrameList) Replace_0() {
	// QTextFrameList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QTextFrameList) RemoveAt_0() {
	// QTextFrameList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QTextFrameList) RemoveAll_0() int {
	// QTextFrameList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QTextFrameList) RemoveOne_0() bool {
	// QTextFrameList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QTextFrameList) TakeAt_0() *QTextFrame {
	// QTextFrameList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// T takeFirst()
func (this *QTextFrameList) TakeFirst_0() *QTextFrame {
	// QTextFrameList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// T takeLast()
func (this *QTextFrameList) TakeLast_0() *QTextFrame {
	// QTextFrameList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// void move(int, int)
func (this *QTextFrameList) Move_0() {
	// QTextFrameList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QTextFrameList) Swap_1() {
	// QTextFrameList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QTextFrameList) IndexOf_0() int {
	// QTextFrameList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QTextFrameList) LastIndexOf_0() int {
	// QTextFrameList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QTextFrameList) Contains_0() bool {
	// QTextFrameList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QTextFrameList) Count_0() int {
	// QTextFrameList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QTextFrameList) Begin_0() {
	// QTextFrameList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QTextFrameList) Begin_1() {
	// QTextFrameList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QTextFrameList) Cbegin_0() {
	// QTextFrameList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QTextFrameList) ConstBegin_0() {
	// QTextFrameList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QTextFrameList) End_0() {
	// QTextFrameList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QTextFrameList) End_1() {
	// QTextFrameList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QTextFrameList) Cend_0() {
	// QTextFrameList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QTextFrameList) ConstEnd_0() {
	// QTextFrameList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QTextFrameList) Rbegin_0() {
	// QTextFrameList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QTextFrameList) Rend_0() {
	// QTextFrameList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QTextFrameList) Rbegin_1() {
	// QTextFrameList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QTextFrameList) Rend_1() {
	// QTextFrameList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QTextFrameList) Crbegin_0() {
	// QTextFrameList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QTextFrameList) Crend_0() {
	// QTextFrameList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QTextFrameList) Insert_1() {
	// QTextFrameList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QTextFrameList) Erase_0() {
	// QTextFrameList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QTextFrameList) Erase_1() {
	// QTextFrameList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QTextFrameList) Count_1() int {
	// QTextFrameList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QTextFrameList) Length_0() int {
	// QTextFrameList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QTextFrameList) First_0() *QTextFrame {
	// QTextFrameList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// const T & constFirst()
func (this *QTextFrameList) ConstFirst_0() *QTextFrame {
	// QTextFrameList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// const T & first()
func (this *QTextFrameList) First_1() *QTextFrame {
	// QTextFrameList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// T & last()
func (this *QTextFrameList) Last_0() *QTextFrame {
	// QTextFrameList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// const T & last()
func (this *QTextFrameList) Last_1() *QTextFrame {
	// QTextFrameList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// const T & constLast()
func (this *QTextFrameList) ConstLast_0() *QTextFrame {
	// QTextFrameList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// void removeFirst()
func (this *QTextFrameList) RemoveFirst_0() {
	// QTextFrameList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QTextFrameList) RemoveLast_0() {
	// QTextFrameList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QTextFrameList) StartsWith_0() bool {
	// QTextFrameList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QTextFrameList) EndsWith_0() bool {
	// QTextFrameList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QTextFrameList) Mid_0() *QTextFrameList {
	// QTextFrameList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QTextFrameList) Value_0() *QTextFrame {
	// QTextFrameList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// T value(int, const T &)
func (this *QTextFrameList) Value_1() *QTextFrame {
	// QTextFrameList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// void push_back(const T &)
func (this *QTextFrameList) Push_back_0() {
	// QTextFrameList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QTextFrameList) Push_front_0() {
	// QTextFrameList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QTextFrameList) Front_0() *QTextFrame {
	// QTextFrameList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// const T & front()
func (this *QTextFrameList) Front_1() *QTextFrame {
	// QTextFrameList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// T & back()
func (this *QTextFrameList) Back_0() *QTextFrame {
	// QTextFrameList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// const T & back()
func (this *QTextFrameList) Back_1() *QTextFrame {
	// QTextFrameList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextFrame{}
}

// void pop_front()
func (this *QTextFrameList) Pop_front_0() {
	// QTextFrameList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QTextFrameList) Pop_back_0() {
	// QTextFrameList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QTextFrameList) Empty_0() bool {
	// QTextFrameList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QTextFrameList) Operator_add_equal_0() *QTextFrameList {
	// QTextFrameList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QTextFrameList) Operator_add_0() *QTextFrameList {
	// QTextFrameList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QTextFrameList) Operator_add_equal_1() *QTextFrameList {
	// QTextFrameList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QTextFrameList) Operator_left_shift_0() *QTextFrameList {
	// QTextFrameList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QTextFrameList) Operator_left_shift_1() *QTextFrameList {
	// QTextFrameList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QTextFrameList) ToVector_0() {
	// QTextFrameList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QTextFrameList) ToSet_0() {
	// QTextFrameList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QTextFrameList) FromVector_0() *QTextFrameList {
	// QTextFrameList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QTextFrameList) FromSet_0() *QTextFrameList {
	// QTextFrameList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QTextFrameList) FromStdList_0() *QTextFrameList {
	// QTextFrameList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QTextFrameList) ToStdList_0() {
	// QTextFrameList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QTextFrameList) Detach_helper_grow_0() {
	// QTextFrameList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QTextFrameList) Detach_helper_0() {
	// QTextFrameList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QTextFrameList) Detach_helper_1() {
	// QTextFrameList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QTextFrameList) Dealloc_0() {
	// QTextFrameList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QTextFrameList) Node_construct_0() {
	// QTextFrameList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QTextFrameList) Node_destruct_0() {
	// QTextFrameList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QTextFrameList) Node_copy_0() {
	// QTextFrameList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QTextFrameList) Node_destruct_1() {
	// QTextFrameList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QTextFrameList) IsValidIterator_0() bool {
	// QTextFrameList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QTextFrameList) Op_eq_impl_0() bool {
	// QTextFrameList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QTextFrameList) Op_eq_impl_1() bool {
	// QTextFrameList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QTextFrameList) Contains_impl_0() bool {
	// QTextFrameList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QTextFrameList) Contains_impl_1() bool {
	// QTextFrameList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QTextFrameList) Count_impl_0() int {
	// QTextFrameList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QTextFrameList) Count_impl_1() int {
	// QTextFrameList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextFrameList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
