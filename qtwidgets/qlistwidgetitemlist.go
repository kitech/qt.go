package qtwidgets

// /usr/include/qt/QtWidgets/qlistwidget.h
// #include <qlistwidget.h>
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
type QListWidgetItemList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QListWidgetItemList) Operator_equal_0() *QListWidgetItemList {
	// QListWidgetItemList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QListWidgetItemList) Operator_equal_1() *QListWidgetItemList {
	// QListWidgetItemList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QListWidgetItemList) Swap_0() {
	// QListWidgetItemList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QListWidgetItemList) Operator_equal_equal_0() bool {
	// QListWidgetItemList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QListWidgetItemList) Operator_not_equal_0() bool {
	// QListWidgetItemList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QListWidgetItemList) Size_0() int {
	// QListWidgetItemList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QListWidgetItemList) Detach_0() {
	// QListWidgetItemList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QListWidgetItemList) DetachShared_0() {
	// QListWidgetItemList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QListWidgetItemList) IsDetached_0() bool {
	// QListWidgetItemList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QListWidgetItemList) SetSharable_0() {
	// QListWidgetItemList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QListWidgetItemList) IsSharedWith_0() bool {
	// QListWidgetItemList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QListWidgetItemList) IsEmpty_0() bool {
	// QListWidgetItemList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QListWidgetItemList) Clear_0() {
	// QListWidgetItemList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QListWidgetItemList) At_0() *QListWidgetItem {
	// QListWidgetItemList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// const T & operator[](int)
func (this *QListWidgetItemList) Operator_get_index_0() *QListWidgetItem {
	// QListWidgetItemList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// T & operator[](int)
func (this *QListWidgetItemList) Operator_get_index_1() *QListWidgetItem {
	// QListWidgetItemList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// void reserve(int)
func (this *QListWidgetItemList) Reserve_0() {
	// QListWidgetItemList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QListWidgetItemList) Append_0() {
	// QListWidgetItemList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QListWidgetItemList) Append_1() {
	// QListWidgetItemList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QListWidgetItemList) Prepend_0() {
	// QListWidgetItemList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QListWidgetItemList) Insert_0() {
	// QListWidgetItemList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QListWidgetItemList) Replace_0() {
	// QListWidgetItemList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QListWidgetItemList) RemoveAt_0() {
	// QListWidgetItemList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QListWidgetItemList) RemoveAll_0() int {
	// QListWidgetItemList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QListWidgetItemList) RemoveOne_0() bool {
	// QListWidgetItemList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QListWidgetItemList) TakeAt_0() *QListWidgetItem {
	// QListWidgetItemList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// T takeFirst()
func (this *QListWidgetItemList) TakeFirst_0() *QListWidgetItem {
	// QListWidgetItemList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// T takeLast()
func (this *QListWidgetItemList) TakeLast_0() *QListWidgetItem {
	// QListWidgetItemList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// void move(int, int)
func (this *QListWidgetItemList) Move_0() {
	// QListWidgetItemList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QListWidgetItemList) Swap_1() {
	// QListWidgetItemList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QListWidgetItemList) IndexOf_0() int {
	// QListWidgetItemList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QListWidgetItemList) LastIndexOf_0() int {
	// QListWidgetItemList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QListWidgetItemList) Contains_0() bool {
	// QListWidgetItemList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QListWidgetItemList) Count_0() int {
	// QListWidgetItemList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QListWidgetItemList) Begin_0() {
	// QListWidgetItemList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QListWidgetItemList) Begin_1() {
	// QListWidgetItemList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QListWidgetItemList) Cbegin_0() {
	// QListWidgetItemList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QListWidgetItemList) ConstBegin_0() {
	// QListWidgetItemList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QListWidgetItemList) End_0() {
	// QListWidgetItemList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QListWidgetItemList) End_1() {
	// QListWidgetItemList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QListWidgetItemList) Cend_0() {
	// QListWidgetItemList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QListWidgetItemList) ConstEnd_0() {
	// QListWidgetItemList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QListWidgetItemList) Rbegin_0() {
	// QListWidgetItemList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QListWidgetItemList) Rend_0() {
	// QListWidgetItemList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QListWidgetItemList) Rbegin_1() {
	// QListWidgetItemList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QListWidgetItemList) Rend_1() {
	// QListWidgetItemList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QListWidgetItemList) Crbegin_0() {
	// QListWidgetItemList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QListWidgetItemList) Crend_0() {
	// QListWidgetItemList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QListWidgetItemList) Insert_1() {
	// QListWidgetItemList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QListWidgetItemList) Erase_0() {
	// QListWidgetItemList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QListWidgetItemList) Erase_1() {
	// QListWidgetItemList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QListWidgetItemList) Count_1() int {
	// QListWidgetItemList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QListWidgetItemList) Length_0() int {
	// QListWidgetItemList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QListWidgetItemList) First_0() *QListWidgetItem {
	// QListWidgetItemList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// const T & constFirst()
func (this *QListWidgetItemList) ConstFirst_0() *QListWidgetItem {
	// QListWidgetItemList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// const T & first()
func (this *QListWidgetItemList) First_1() *QListWidgetItem {
	// QListWidgetItemList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// T & last()
func (this *QListWidgetItemList) Last_0() *QListWidgetItem {
	// QListWidgetItemList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// const T & last()
func (this *QListWidgetItemList) Last_1() *QListWidgetItem {
	// QListWidgetItemList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// const T & constLast()
func (this *QListWidgetItemList) ConstLast_0() *QListWidgetItem {
	// QListWidgetItemList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// void removeFirst()
func (this *QListWidgetItemList) RemoveFirst_0() {
	// QListWidgetItemList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QListWidgetItemList) RemoveLast_0() {
	// QListWidgetItemList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QListWidgetItemList) StartsWith_0() bool {
	// QListWidgetItemList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QListWidgetItemList) EndsWith_0() bool {
	// QListWidgetItemList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QListWidgetItemList) Mid_0() *QListWidgetItemList {
	// QListWidgetItemList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QListWidgetItemList) Value_0() *QListWidgetItem {
	// QListWidgetItemList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// T value(int, const T &)
func (this *QListWidgetItemList) Value_1() *QListWidgetItem {
	// QListWidgetItemList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// void push_back(const T &)
func (this *QListWidgetItemList) Push_back_0() {
	// QListWidgetItemList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QListWidgetItemList) Push_front_0() {
	// QListWidgetItemList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QListWidgetItemList) Front_0() *QListWidgetItem {
	// QListWidgetItemList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// const T & front()
func (this *QListWidgetItemList) Front_1() *QListWidgetItem {
	// QListWidgetItemList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// T & back()
func (this *QListWidgetItemList) Back_0() *QListWidgetItem {
	// QListWidgetItemList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// const T & back()
func (this *QListWidgetItemList) Back_1() *QListWidgetItem {
	// QListWidgetItemList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QListWidgetItem{}
}

// void pop_front()
func (this *QListWidgetItemList) Pop_front_0() {
	// QListWidgetItemList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QListWidgetItemList) Pop_back_0() {
	// QListWidgetItemList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QListWidgetItemList) Empty_0() bool {
	// QListWidgetItemList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QListWidgetItemList) Operator_add_equal_0() *QListWidgetItemList {
	// QListWidgetItemList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QListWidgetItemList) Operator_add_0() *QListWidgetItemList {
	// QListWidgetItemList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QListWidgetItemList) Operator_add_equal_1() *QListWidgetItemList {
	// QListWidgetItemList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QListWidgetItemList) Operator_left_shift_0() *QListWidgetItemList {
	// QListWidgetItemList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QListWidgetItemList) Operator_left_shift_1() *QListWidgetItemList {
	// QListWidgetItemList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QListWidgetItemList) ToVector_0() {
	// QListWidgetItemList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QListWidgetItemList) ToSet_0() {
	// QListWidgetItemList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QListWidgetItemList) FromVector_0() *QListWidgetItemList {
	// QListWidgetItemList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QListWidgetItemList) FromSet_0() *QListWidgetItemList {
	// QListWidgetItemList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QListWidgetItemList) FromStdList_0() *QListWidgetItemList {
	// QListWidgetItemList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QListWidgetItemList) ToStdList_0() {
	// QListWidgetItemList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QListWidgetItemList) Detach_helper_grow_0() {
	// QListWidgetItemList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QListWidgetItemList) Detach_helper_0() {
	// QListWidgetItemList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QListWidgetItemList) Detach_helper_1() {
	// QListWidgetItemList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QListWidgetItemList) Dealloc_0() {
	// QListWidgetItemList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QListWidgetItemList) Node_construct_0() {
	// QListWidgetItemList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QListWidgetItemList) Node_destruct_0() {
	// QListWidgetItemList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QListWidgetItemList) Node_copy_0() {
	// QListWidgetItemList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QListWidgetItemList) Node_destruct_1() {
	// QListWidgetItemList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QListWidgetItemList) IsValidIterator_0() bool {
	// QListWidgetItemList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QListWidgetItemList) Op_eq_impl_0() bool {
	// QListWidgetItemList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QListWidgetItemList) Op_eq_impl_1() bool {
	// QListWidgetItemList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QListWidgetItemList) Contains_impl_0() bool {
	// QListWidgetItemList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QListWidgetItemList) Contains_impl_1() bool {
	// QListWidgetItemList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QListWidgetItemList) Count_impl_0() int {
	// QListWidgetItemList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QListWidgetItemList) Count_impl_1() int {
	// QListWidgetItemList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QListWidgetItemList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
