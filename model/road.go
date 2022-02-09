package model

type Road struct {
	ID            int
	Name          string
	From          int
	To            int
	Through       []int
	SpeedLimit    int
	Length        int
	BiDirectional bool
}
