//go:build linux

// Package command provides functions to decode MySQL command packets.
package command

import (
	"context"
	"fmt"

	"go.keploy.io/server/v2/pkg/models/mysql"
)

// COM_QUERY: https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_com_query.html

func DecodeQuery(_ context.Context, data []byte) (*mysql.QueryPacket, error) {
	if len(data) < 2 {
		return nil, fmt.Errorf("query packet too short")
	}

	packet := &mysql.QueryPacket{
		Command: data[0],
		Query:   string(data[1:]),
	}

	return packet, nil
}