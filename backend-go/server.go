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
// Hash 1584
// Hash 2231
// Hash 7507
// Hash 9666
// Hash 8803
// Hash 9387
// Hash 4979
// Hash 7928
// Hash 7751
// Hash 5719
// Hash 6452
// Hash 8812
// Hash 6718
// Hash 7979
// Hash 8989
// Hash 6934
// Hash 9571
// Hash 7035
// Hash 5343
// Hash 7984
// Hash 7246
// Hash 4888
// Hash 6925
// Hash 6767
// Hash 6189
// Hash 7733
// Hash 4075
// Hash 3283
// Hash 8026
// Hash 7298
// Hash 4632
// Hash 2384
// Hash 7729
// Hash 1311
// Hash 1332
// Hash 5427
// Hash 8388
// Hash 4386
// Hash 8365
// Hash 7898
// Hash 1711
// Hash 7176
// Hash 2693
// Hash 5024
// Hash 3250
// Hash 8648
// Hash 8673
// Hash 6693
// Hash 9944
// Hash 3315
// Hash 2359
// Hash 6697
// Hash 9060
// Hash 9838
// Hash 4643
// Hash 6098
// Hash 7772
// Hash 7648
// Hash 3699
// Hash 9998
// Hash 2140
// Hash 3692
// Hash 1814