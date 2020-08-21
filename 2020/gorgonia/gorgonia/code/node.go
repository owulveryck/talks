package xvm

import (
	"context"
	"errors"

	"gorgonia.org/gorgonia"
)

// Doer is implementing the Do method of gorgonia's Op interface
type Doer interface {
	Do(...gorgonia.Value) (gorgonia.Value, error)
}

type node struct {
	// Some fields are omitted for clarity
	id             int64 // OMIT
	op             Doer
	output         gorgonia.Value //OMIT
	outputC        chan gorgonia.Value
	receivedValues int
	err            error
	inputValues    []gorgonia.Value //OMIT
	inputC         chan ioValue
}

// ioValue is a value with a position. as the infrastructure cannot guaranty the
// order of the input values, we use this structure carrying the position of the operator.
// this is mandatory for non commutative operations
type ioValue struct {
	pos int
	v   gorgonia.Value
}

type stateFn func(context.Context, *node) stateFn

// START_DEFAULT_STATE OMIT
func defaultState(_ context.Context, n *node) stateFn {
	n.receivedValues = 0
	n.err = nil
	if n.op == nil {
		return emitOutput
	}
	return receiveInput
}

// END_DEFAULT_STATE OMIT

// START_RECEIVEINPUT_STATE OMIT
func receiveInput(ctx context.Context, n *node) stateFn {
	// if inputC is nil, it is a variable or a constant, don't
	// wait for any input
	if n.inputC == nil {
		return computeFwd
	}
	select {
	case <-ctx.Done():
		n.err = ctx.Err()
		return nil
	case input := <-n.inputC:
		if input.pos >= len(n.inputValues) {
			n.err = errors.New("bad arity")
			return nil
		}
		n.receivedValues++
		n.inputValues[input.pos] = input.v
		if n.receivedValues < len(n.inputValues) {
			return receiveInput
		}
	}
	return computeFwd
}

// END_RECEIVEINPUT_STATE OMIT

// START_COMPUTEFWD_STATE OMIT
func computeFwd(_ context.Context, n *node) stateFn {
	v, err := n.op.Do(n.inputValues...)
	if err != nil {
		n.err = err
		return nil
	}
	n.output = v
	return emitOutput
}

// END_COMPUTEFWD_STATE OMIT

// START_EMITOUTPUT_STATE OMIT
func emitOutput(ctx context.Context, n *node) stateFn {
	if n == nil || n.outputC == nil {
		return nil
	}
	select {
	case <-ctx.Done():
		n.err = ctx.Err()
		return nil
	case n.outputC <- n.output:
	}
	return nil
}

// END_EMITOUTPUT_STATE OMIT

func computeBackward(_ context.Context, _ *node) stateFn {
	return nil
}

// START_COMPUTE OMIT
func (n *node) Compute(ctx context.Context) error {
	for state := defaultState; state != nil; {
		t := trace(ctx, nil, n, state) // OMIT
		state = state(ctx, n)
		trace(ctx, t, nil, nil) //OMIT
	}
	return n.err
}

// END_COMPUTE OMIT

func newOp(n *gorgonia.Node, hasOutputChan bool) *node {
	if n == nil {
		return nil
	}
	var outputC chan gorgonia.Value
	if hasOutputChan {
		outputC = make(chan gorgonia.Value, 0)

	}
	return &node{
		id:          n.ID(),
		op:          n.Op(),
		inputValues: make([]gorgonia.Value, n.Op().Arity()),
		inputC:      make(chan ioValue, 0),
		outputC:     outputC,
	}
}

func newInput(n *gorgonia.Node) *node {
	if n == nil {
		return nil
	}
	return &node{
		id:      n.ID(),
		output:  n.Value(),
		outputC: make(chan gorgonia.Value, 0),
	}
}
