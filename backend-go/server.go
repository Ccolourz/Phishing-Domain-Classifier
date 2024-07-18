package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 1484
// Hash 7249
// Hash 4413
// Hash 9249
// Hash 9122
// Hash 8489
// Hash 1479
// Hash 7880
// Hash 3769
// Hash 2811
// Hash 9956
// Hash 3968
// Hash 7608
// Hash 1477
// Hash 7324
// Hash 9870
// Hash 5082
// Hash 9561
// Hash 3993
// Hash 3756
// Hash 1668
// Hash 6288
// Hash 1280
// Hash 4813
// Hash 2488
// Hash 5247
// Hash 8324
// Hash 1115
// Hash 2624
// Hash 5644
// Hash 4816
// Hash 4294
// Hash 2734
// Hash 7300
// Hash 2693
// Hash 9984
// Hash 8330
// Hash 8478
// Hash 3535
// Hash 5496
// Hash 1496
// Hash 5409
// Hash 9046
// Hash 8132
// Hash 1796