//go:build !linux && !darwin

package grpcexecutor

import (
	"os"

	"backend-go/sandbox/pb"
)

func setWinsize(f *os.File, i *pb.StreamRequest_ExecResize) error {
	return nil
}
