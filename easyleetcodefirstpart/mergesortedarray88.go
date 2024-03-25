package easyleetcodefirstpart

func Merge(nums1 []int, m int, nums2 []int, n int) {

	for m > 0 && n > 0 {

		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
		if m > 0 && nums2[n-1] >= nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	if n > 0 && m == 0 {
		copy(nums1[:n], nums2[:n])
	}
}

//func merge(nums1 []int, m int, nums2 []int, n int)  {
//	nums1 = append(nums1[:m], nums2...)
//	slices.Sort(nums1)
//}
// Но Использование append фактически создает новый слайс, а не использует предоставленное пространство.
