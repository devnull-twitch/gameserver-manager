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
	log.Printf("incoming request for zone %s", payload.GetZone())

	switch payload.GetZone() {
	case "overworld":
	case "otherworld":
	default:
		return nil, fmt.Errorf("invalid zone. got %s", payload.GetZone())
	}

	if _, hasZoneServer := s.cache[payload.GetZone()]; !hasZoneServer {
		if len(s.availablePorts) <= 0 {
			return nil, fmt.Errorf("no available ports")
		}
		usedPort := s.availablePorts[0]

		startCms := exec.Command(
			"/home/devnull/Downloads/Godot_v3.4.2-stable_mono_x11_64/Godot_v3.4.2-stable_mono_x11.64",
			"--no-window",
			"--path",
			"/home/devnull/dev/2donlinerpg",
			fmt.Sprintf("%s.tscn", payload.GetZone()),
			"--server",
			fmt.Sprintf("%d", usedPort),
		)

		logFileRef, err := os.Create(fmt.Sprintf("%s.log.txt", payload.GetZone()))
		if err != nil {
			panic(err)
		}
		reader, err := startCms.StdoutPipe()
		if err != nil {
			panic(err)
		}
		startCms.Start()

		s.processes = append(s.processes, startCms.Process)

		s.availablePorts = s.availablePorts[1:]
		s.cache[payload.GetZone()] = &proto.GetResponse{
			GsIp:   "0.0.0.0",
			GsPort: usedPort,
		}

		go func() {
			goon := true
			for goon {
				buffer := make([]byte, 1024)
				_, err := reader.Read(buffer)
				if err != nil {
					log.Printf("unable to read from game server stdout: %s", err)
					goon = false
				}

				logFileRef.Write(buffer)
			}

			if err := startCms.Wait(); err != nil {
				log.Printf("wait error for game server: %s", err)
			}

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
