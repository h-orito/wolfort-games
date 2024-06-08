package gqlmodel

import (
	"chat-role-play/domain/model"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

func (q *PageableQuery) MapToPagingQuery() *model.PagingQuery {
	if q == nil {
		return nil
	}
	return &model.PagingQuery{
		PageSize:   q.PageSize,
		PageNumber: q.PageNumber,
		Desc:       q.IsDesc,
		Latest:     q.IsLatest,
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
