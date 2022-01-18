package grpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/devnull-twitch/gameserver-manager/proto"
)

type GameServerManager interface {
	proto.GameserverManagerServer

	StopServer()
}

type server struct {
	proto.UnimplementedGameserverManagerServer

	cache          map[string]*proto.GetResponse
	processes      []*os.Process
	availablePorts []int64
}

func (s *server) GetGameserver(ctx context.Context, payload *proto.GetRequest) (*proto.GetResponse, error) {
	log.Printf("incoming request: %+v", payload)
	if _, hasZoneServer := s.cache[payload.GetZone()]; !hasZoneServer {
		startCms := exec.Command(
			"/home/devnull/Downloads/Godot_v3.4.2-stable_mono_x11_64/Godot_v3.4.2-stable_mono_x11.64",
			"--no-window",
			"--server",
			"--path",
			"/home/devnull/dev/2donlinerpg",
		)

		startCms.Start()

		s.processes = append(s.processes, startCms.Process)

		if len(s.availablePorts) <= 0 {
			return nil, fmt.Errorf("no available ports")
		}
		usedPort := s.availablePorts[0]
		s.availablePorts = s.availablePorts[1:]
		s.cache[payload.GetZone()] = &proto.GetResponse{
			GsIp:   "0.0.0.0",
			GsPort: usedPort,
		}

		go func() {
			startCms.Wait()
			s.availablePorts = append(s.availablePorts, s.cache[payload.GetZone()].GetGsPort())
		}()
	}

	return s.cache[payload.GetZone()], nil
}

func (s *server) StopServer() {
	for _, proc := range s.processes {
		proc.Kill()
	}
}

func GetServer() GameServerManager {
	return &server{
		processes:      []*os.Process{},
		cache:          map[string]*proto.GetResponse{},
		availablePorts: []int64{50123, 50124},
	}
}
