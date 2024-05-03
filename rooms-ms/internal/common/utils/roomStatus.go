package utils

import "github.com/tesis/internal/common/constants"

func CheckRoomStatus(roomStatus string) bool {
	switch roomStatus {
	case constants.Active, constants.Deleted, constants.Inactive:
		return true
	default:
		return false
	}

}
