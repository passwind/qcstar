package problems

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	res := findMedianSortedArrays(nums1, nums2)
	if res != 2.0 {
		t.Errorf("error : should be 2.0, in fact %f", res)
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	res = findMedianSortedArrays(nums1, nums2)
	if res != 2.5 {
		t.Errorf("error : should be 2.5, in fact %f", res)
	}

	nums1 = []int{}
	nums2 = []int{0}
	res = findMedianSortedArrays(nums1, nums2)
	if res != 0 {
		t.Errorf("error : should be 0, in fact %f", res)
	}

	nums1 = []int{-100, -1}
	nums2 = []int{0}
	res = findMedianSortedArrays(nums1, nums2)
	if res != -1 {
		t.Errorf("error : should be -1, in fact %f", res)
	}

	nums1 = []int{}
	nums2 = []int{0, 4}
	res = findMedianSortedArrays(nums1, nums2)
	if res != 2 {
		t.Errorf("error : should be 2, in fact %f", res)
	}

	nums1 = []int{0, 4, 6}
	nums2 = []int{}
	res = findMedianSortedArrays(nums1, nums2)
	if res != 4 {
		t.Errorf("error : should be 4, in fact %f", res)
	}
}
