package gsmanager

import (
	"errors"
	"fmt"
)

type (
	Getter func(community string, zone string) (int, error)

	portList     []int
	zoneMap      map[string]portList
	communityMap map[string]zoneMap
)

var mem = communityMap{}

var (
	ErrUnknownCommunity = errors.New("unknown")
	ErrUnknownZone      = errors.New("unknown")
)

func GetOrCreate(community string, zone string) (int, error) {
	if _, communityExists := mem[community]; !communityExists {
		return 0, fmt.Errorf("unable to fetch community %s: %w", community, ErrUnknownCommunity)
	}

	if _, zoneExists := mem[community][zone]; !zoneExists {
		// todo: check if valid zone => create zone and continue
		return 0, fmt.Errorf("unable to fetch zone %s: %w", zone, ErrUnknownZone)
	}

	for _, i := range mem[community][zone] {
		return i, nil
	}

	return 0, fmt.Errorf("not yet implemented")
}
