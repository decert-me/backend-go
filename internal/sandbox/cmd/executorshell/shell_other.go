//go:build !linux

package main

import "backend-go/sandbox/pb"

func handleSizeChange(sendCh chan<- *pb.StreamRequest) {
}
