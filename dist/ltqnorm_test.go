package dist

import(
	"testing"
)


func TestInvCumulativeNormalDistribution(t *testing.T) {
	delta := 0.00000001
	if x:= InvCumulativeNormalDistribution(0.00000000001); (x - -6.70602314994) > delta {
		t.Error(x)
	}
	if x:= InvCumulativeNormalDistribution(0.2); (x -  -0.841621232727) > delta {
		t.Error(x)
	}
	if x:= InvCumulativeNormalDistribution(0.5); x != 0 {
		t.Error(x)
	}
	if x:= InvCumulativeNormalDistribution(0.6); (x - 0.25334710286) > delta {
		t.Error(x)
	}
	if x:= InvCumulativeNormalDistribution(0.999999999999); (x - 7.03448690268) > delta {
		t.Error(x)
	}
}
