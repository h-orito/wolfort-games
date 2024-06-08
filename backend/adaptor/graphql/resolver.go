package graphql

import (
	"chat-role-play/application/usecase"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	playerUsecase        usecase.PlayerUsecase
	chinchiroRoomUsecase usecase.ChinchiroRoomUsecase
	chinchiroGameUsecase usecase.ChinchiroGameUsecase
	loaders              *Loaders
}

func NewResolver(
	playerUsecase usecase.PlayerUsecase,
	chinchiroRoomUsecase usecase.ChinchiroRoomUsecase,
	chinchiroGameUsecase usecase.ChinchiroGameUsecase,
	loaders *Loaders,
) Resolver {
	return Resolver{
		playerUsecase:        playerUsecase,
		chinchiroRoomUsecase: chinchiroRoomUsecase,
		chinchiroGameUsecase: chinchiroGameUsecase,
		loaders:              loaders,
	}
}

func idToUint32(id string) (uint32, error) {
	byte, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return 0, err
	}
	parts := strings.Split(string(byte), ":")
	if len(parts) == 2 {
		number, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(number), nil
	} else {
		return 0, fmt.Errorf("Invalid input format")
	}
}

func idsToUint32s(ids []string) ([]uint32, error) {
	var intids []uint32
	for _, id := range ids {
		intid, err := idToUint32(id)
		if err != nil {
			return nil, err
		}
		intids = append(intids, intid)
	}
	return intids, nil
}

func idToUint64(id string) (uint64, error) {
	byte, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return 0, err
	}
	parts := strings.Split(string(byte), ":")
	if len(parts) == 2 {
		number, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			return 0, err
		}
		return number, nil
	} else {
		return 0, fmt.Errorf("Invalid input format")
	}
}
