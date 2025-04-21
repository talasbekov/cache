package main

import (
	"fmt"
	"github.com/talasbekov/cache/cache"
)

func main() {
	c := cache.New()
	c.Set("userId", 58)

	userId := c.Get("userId")
	fmt.Println(userId)

	c.Delete("userId")
	userId = c.Get("userId")
	fmt.Println(userId)
}

//// Главная функция
//func main() {
//	// Число 342 (в обратном порядке: 2 -> 4 -> 3)
//	l1 := createList([]int{9, 7, 3})
//	// Число 465 (в обратном порядке: 5 -> 6 -> 4)
//	l2 := createList([]int{7, 8, 4})
//
//	result := addTwoNumbers(l1, l2)
//
//	fmt.Print(result)
//	//printList(result) // Должно вывести: 7 -> 0 -> 8 (807)
//}

//func main() {
//	dim := 100
//	s := make([]int, 0, dim)
//	// заполняем слайс числами
//	for i := 0; i < dim; i++ {
//		s = append(s, i+1)
//	}
//	// оставляем первые и последние 10 элементов
//	s = append(s[:10], s[dim-10:]...)
//	dim = len(s)
//	// разворачиваем слайс
//	for i := range s[:dim/2] {
//		s[i], s[dim-i-1] = s[dim-i-1], s[i]
//	}
//	fmt.Println(s)
//}

// Разернуть строку таким образом не получится, так как строка — неизменяемый тип данных.
// Строку можно преобразовать к слайсу байт ([]byte), но это опасно, если строка содержит Unicode-символы
// Лучше создать слайс рун из строки и развернуть его
// Например, так:

//// ListNode Определяем структуру
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//// Функция для сложения двух списков
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	dummyHead := &ListNode{}
//	current := dummyHead
//	carry := 0
//
//	for l1 != nil || l2 != nil || carry > 0 {
//		val1 := 0
//		if l1 != nil {
//			val1 = l1.Val
//			l1 = l1.Next
//		}
//
//		val2 := 0
//		if l2 != nil {
//			val2 = l2.Val
//			l2 = l2.Next
//		}
//
//		sum := val1 + val2 + carry
//		carry = sum / 10
//
//		current.Next = &ListNode{Val: sum % 10}
//		current = current.Next
//	}
//
//	return dummyHead.Next
//}

//// Вспомогательная функция для создания списка из слайса
//func createList(nums []int) *ListNode {
//	dummy := &ListNode{}
//	current := dummy
//	for _, n := range nums {
//		current.Next = &ListNode{Val: n}
//		current = current.Next
//	}
//	return dummy.Next
//}

// Вспомогательная функция для печати списка
//func printList(node *ListNode) {
//	for node != nil {
//		fmt.Print(node.Val)
//		if node.Next != nil {
//			fmt.Print(" -> ")
//		}
//		node = node.Next
//	}
//	fmt.Println()
//}

//func increment() func() int {
//	count := 56
//	return func() int {
//		count++
//		return count
//	}
//}
//
//func increment2() int {
//	count := 1
//	count++
//	return count
//}
