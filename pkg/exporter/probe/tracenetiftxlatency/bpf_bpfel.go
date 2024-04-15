// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package tracenetif

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfInspNftxlatEventT struct {
	Target [20]int8
	Type   uint32
	Tuple  struct {
		Saddr   struct{ V6addr [16]uint8 }
		Daddr   struct{ V6addr [16]uint8 }
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
	Pid     uint32
	Cpu     uint32
	_       [4]byte
	Latency uint64
	StackId int64
}

type bpfInspNftxlatMetricT struct {
	Netns  uint32
	Bucket uint32
	Action uint32
	Cpu    uint32
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
	NetDevQueue     *ebpf.ProgramSpec `ebpf:"net_dev_queue"`
	NetDevStartXmit *ebpf.ProgramSpec `ebpf:"net_dev_start_xmit"`
	NetDevXmit      *ebpf.ProgramSpec `ebpf:"net_dev_xmit"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	InspSklatEvent  *ebpf.MapSpec `ebpf:"insp_sklat_event"`
	InspSklatMetric *ebpf.MapSpec `ebpf:"insp_sklat_metric"`
	InspTxq         *ebpf.MapSpec `ebpf:"insp_txq"`
	InspTxs         *ebpf.MapSpec `ebpf:"insp_txs"`
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
	InspSklatEvent  *ebpf.Map `ebpf:"insp_sklat_event"`
	InspSklatMetric *ebpf.Map `ebpf:"insp_sklat_metric"`
	InspTxq         *ebpf.Map `ebpf:"insp_txq"`
	InspTxs         *ebpf.Map `ebpf:"insp_txs"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.InspSklatEvent,
		m.InspSklatMetric,
		m.InspTxq,
		m.InspTxs,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	NetDevQueue     *ebpf.Program `ebpf:"net_dev_queue"`
	NetDevStartXmit *ebpf.Program `ebpf:"net_dev_start_xmit"`
	NetDevXmit      *ebpf.Program `ebpf:"net_dev_xmit"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.NetDevQueue,
		p.NetDevStartXmit,
		p.NetDevXmit,
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
//go:embed bpf_bpfel.o
var _BpfBytes []byte
