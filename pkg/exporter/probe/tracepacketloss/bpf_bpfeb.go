// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64be || armbe || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64
// +build arm64be armbe mips mips64 mips64p32 ppc64 s390 s390x sparc sparc64

package tracepacketloss

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfInspPlEventT struct {
	Target [20]int8
	Tuple  struct {
		Saddr struct {
			V4addr uint32
			_      [12]byte
		}
		Daddr struct {
			V4addr uint32
			_      [12]byte
		}
		Sport   uint16
		Dport   uint16
		L3Proto uint16
		L4Proto uint8
		Pad     uint8
	}
	SkbMeta struct {
		Netns    uint32
		Mark     uint32
		Ifindex  uint32
		Len      uint32
		Mtu      uint32
		SkState  uint32
		Protocol uint16
		Pad      uint16
	}
	Pid      uint32
	Cpu      uint32
	Location uint64
	StackId  int64
}

type bpfInspPlMetricT struct {
	Location uint64
	Netns    uint32
	Protocol uint8
	_        [3]byte
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	KfreeSkb *ebpf.ProgramSpec `ebpf:"kfree_skb"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	InspPlEvent  *ebpf.MapSpec `ebpf:"insp_pl_event"`
	InspPlMetric *ebpf.MapSpec `ebpf:"insp_pl_metric"`
	InspPlStack  *ebpf.MapSpec `ebpf:"insp_pl_stack"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	InspPlEvent  *ebpf.Map `ebpf:"insp_pl_event"`
	InspPlMetric *ebpf.Map `ebpf:"insp_pl_metric"`
	InspPlStack  *ebpf.Map `ebpf:"insp_pl_stack"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.InspPlEvent,
		m.InspPlMetric,
		m.InspPlStack,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	KfreeSkb *ebpf.Program `ebpf:"kfree_skb"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.KfreeSkb,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfeb.o
var _BpfBytes []byte
