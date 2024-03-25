package easyleetcodefirstpart

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	ex1 = "A man, a plan, a canal: Panama"
	ex2 = "race a car"
	ex3 = " "
)

func IsPalindrome(s string) bool {
	ts := s
	ts = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(ts, "")
	ts = strings.ToLower(ts)
	if ts == "" {
		return true
	}
	j := len(ts) - 1
	for i := 0; i < len(ts); i++ {

		fmt.Println(ts[i], ts[j])

		if ts[i] != ts[j] {
			return false
		}
		if i >= j {
			return true
		}
		j -= 1
	}
	fmt.Println(ts)
	return false
}

//func isPalindrome(s string) bool {
//	x := buildLowerCaseString(s)
//	for i,j :=0, len(x) - 1; i < j; i,j = i+1, j-1 {
//		if x[i] != x[j] {
//			return false
//		}
//	}
//	return true
//}
//
//func isUpper(c rune) bool {
//	return (c >= 'A' && c <= 'Z')
//}
//
//func toLowerCase(c rune) rune {
//	if isUpper(c) {
//		return c + 32
//	}
//	return c
//}
//
//func buildLowerCaseString(s string) string {
//	var sb strings.Builder
//	for _, c := range s {
//		if isAlphanumeric(c) {
//			sb.WriteRune(toLowerCase(c))
//		}
//	}
//	return sb.String()
//}
//
//
//func isAlphanumeric(c rune) bool {
//	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
//}
//Преобразование строки: Вы используете собственные функции для фильтрации алфавитно-цифровых символов и
//преобразования их в нижний регистр. Этот подход избегает использования регулярных выражений, что уменьшает
//накладные расходы на компиляцию и применение регулярного выражения к строке.

//Построение новой строки: Вы построили новую строку только из алфавитно-цифровых символов в нижнем регистре,
//используя strings.Builder, что является эффективным способом построения строк в Go, поскольку он минимизирует
//количество аллокаций памяти.

//Проверка на палиндром: Выполняется однократный проход по строке с двумя указателями, что является наиболее
//эффективным способом проверки на палиндром.
