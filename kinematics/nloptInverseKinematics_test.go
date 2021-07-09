package kinematics

import (
	"context"
	"testing"

	pb "go.viam.com/core/proto/api/v1"
	"go.viam.com/core/utils"

	"github.com/edaniels/golog"
	"go.viam.com/test"

	"go.viam.com/core/arm"
)

func TestCreateNloptIKSolver(t *testing.T) {
	logger := golog.NewTestLogger(t)
	m, err := ParseJSONFile(utils.ResolveFile("robots/wx250s/wx250s_kinematics.json"))
	test.That(t, err, test.ShouldBeNil)
	ik := CreateNloptIKSolver(m, logger)

	pos := &pb.ArmPosition{X: 360, Z: 362}
	seed := arm.JointPositionsFromRadians([]float64{1, 1, 1, 1, 1, 0})

	_, err = ik.Solve(context.Background(), pos, seed)
	test.That(t, err, test.ShouldBeNil)

	pos = &pb.ArmPosition{X: -46, Y: -23, Z: 372, Theta: utils.RadToDeg(3.92), OX: -0.46, OY: 0.84, OZ: 0.28}

	seed = &pb.JointPositions{Degrees: []float64{49, 28, -101, 0, -73, 0}}

	_, err = ik.Solve(context.Background(), pos, seed)
	test.That(t, err, test.ShouldBeNil)
}
