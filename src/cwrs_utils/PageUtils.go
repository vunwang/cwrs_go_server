package cwrs_utils

func CalcOffset(pageNum int, pageSize int) int {
	return (pageNum - 1) * pageSize
}
