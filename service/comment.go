package service

import "yizu/modules"

// CountHouseScore 计算房子的评分
func CountHouseScore(c *modules.HouseComment) float32 {
	total := c.BedRoomScore + c.KitchenScore + c.LivingRoomScore + c.LivingRoomScore + c.SurroundingScore
	t := float32(total) * 0.15
	t = float32(c.Score) * 0.25

	return t
}

// CheckScanPerm 检查浏览权限
func CheckScanPerm(sub, obj string) bool {

	return true
}
