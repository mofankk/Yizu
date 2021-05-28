package service

import "yizu/modules"

func CountHouseScore(c *modules.HouseComment) float32 {
	total := c.BedRoomScore + c.KitchenScore + c.LivingRoomScore + c.LivingRoomScore + c.SurroundingScore
	t := float32(total) * 0.15
	t = float32(c.Score) * 0.25

	return t
}

