/*
 * Copyright (c) 2024, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2"

	pb "github.com/NVIDIA/topograph/pkg/protos"
	"github.com/NVIDIA/topograph/pkg/topology"
)

// follow example in pkg/toposim/testdata/toposim.yaml
func TestToGraph(t *testing.T) {
	klog.InitFlags(nil)
	if err := flag.Set("v", "5"); err != nil {
		t.Fatalf("Error setting verbosity: %v", err)
	}

	instances := []*pb.Instance{
		{
			Id:           "n10-1",
			InstanceType: "H100",
			NvlinkDomain: "nv1",
		},
		{
			Id:           "n10-2",
			InstanceType: "H100",
			NvlinkDomain: "nv1",
		},
		{
			Id:            "n11-1",
			InstanceType:  "H100",
			NetworkLayers: []string{"sw11", "sw21", "sw3"},
			NvlinkDomain:  "nv1",
		},
		{
			Id:            "n11-2",
			InstanceType:  "H100",
			NetworkLayers: []string{"sw11", "sw21", "sw3"},
			NvlinkDomain:  "nv1",
		},
		{
			Id:            "n12-1",
			InstanceType:  "H100",
			NetworkLayers: []string{"sw12", "sw21", "sw3"},
		},
		{
			Id:            "n12-2",
			InstanceType:  "H100",
			NetworkLayers: []string{"sw12", "sw21", "sw3"},
		},
		{
			Id:            "n13-1",
			InstanceType:  "H100",
			NetworkLayers: []string{"sw13", "sw22", "sw3"},
		},
		{
			Id:            "n13-2",
			InstanceType:  "H100",
			NetworkLayers: []string{"sw13", "sw22", "sw3"},
		},
		{
			Id:            "n14-1",
			InstanceType:  "H100",
			NetworkLayers: []string{"sw14", "sw22", "sw3"},
		},
		{
			Id:            "n14-2",
			InstanceType:  "H100",
			NetworkLayers: []string{"sw14", "sw22", "sw3"},
		},
		{
			Id:            "n15",
			InstanceType:  "H100",
			NetworkLayers: []string{"sw14", "sw22", "sw3"},
			NvlinkDomain:  "nv2",
		},
	}

	cis := []topology.ComputeInstances{
		{
			Instances: map[string]string{
				"n10-1": "N10-1",
				"n10-2": "N10-2",
				"n11-1": "N11-1",
				"n11-2": "N11-2",
				"n12-1": "N12-1",
				"n12-2": "N12-2",
				"n13-1": "N13-1",
				"n13-2": "N13-2",
				"n14-1": "N14-1",
				"n14-2": "N14-2",
				"cpu1":  "CPU1",
			},
		},
	}

	v101 := &topology.Vertex{Name: "N10-1", ID: "n10-1"}
	v102 := &topology.Vertex{Name: "N10-2", ID: "n10-2"}
	v111 := &topology.Vertex{Name: "N11-1", ID: "n11-1"}
	v112 := &topology.Vertex{Name: "N11-2", ID: "n11-2"}
	v121 := &topology.Vertex{Name: "N12-1", ID: "n12-1"}
	v122 := &topology.Vertex{Name: "N12-2", ID: "n12-2"}
	v131 := &topology.Vertex{Name: "N13-1", ID: "n13-1"}
	v132 := &topology.Vertex{Name: "N13-2", ID: "n13-2"}
	v141 := &topology.Vertex{Name: "N14-1", ID: "n14-1"}
	v142 := &topology.Vertex{Name: "N14-2", ID: "n14-2"}
	cpu1 := &topology.Vertex{Name: "CPU1", ID: "cpu1"}

	sw11 := &topology.Vertex{ID: "sw11", Vertices: map[string]*topology.Vertex{"n11-1": v111, "n11-2": v112}}
	sw12 := &topology.Vertex{ID: "sw12", Vertices: map[string]*topology.Vertex{"n12-1": v121, "n12-2": v122}}
	sw13 := &topology.Vertex{ID: "sw13", Vertices: map[string]*topology.Vertex{"n13-1": v131, "n13-2": v132}}
	sw14 := &topology.Vertex{ID: "sw14", Vertices: map[string]*topology.Vertex{"n14-1": v141, "n14-2": v142}}
	sw21 := &topology.Vertex{ID: "sw21", Vertices: map[string]*topology.Vertex{"sw11": sw11, "sw12": sw12}}
	sw22 := &topology.Vertex{ID: "sw22", Vertices: map[string]*topology.Vertex{"sw13": sw13, "sw14": sw14}}
	sw3 := &topology.Vertex{ID: "sw3", Vertices: map[string]*topology.Vertex{"sw21": sw21, "sw22": sw22}}

	nv1 := &topology.Vertex{ID: "nvlink-nv1", Vertices: map[string]*topology.Vertex{"n10-1": v101, "n10-2": v102, "n11-1": v111, "n11-2": v112}}

	extra := &topology.Vertex{ID: topology.NoTopology, Vertices: map[string]*topology.Vertex{"cpu1": cpu1}}
	treeRoot := &topology.Vertex{Vertices: map[string]*topology.Vertex{"sw3": sw3, topology.NoTopology: extra}}
	blockRoot := &topology.Vertex{Vertices: map[string]*topology.Vertex{"nvlink-nv1": nv1}}
	root := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{topology.TopologyBlock: blockRoot, topology.TopologyTree: treeRoot},
	}

	require.Equal(t, root, toGraph(&pb.TopologyResponse{Instances: instances}, cis, topology.TopologyBlock))
}

func TestGetTopologyFormat(t *testing.T) {
	testCases := []struct {
		name   string
		params map[string]any
		format string
	}{
		{
			name:   "Case 1: nil params",
			params: nil,
			format: topology.TopologyTree,
		},
		{
			name:   "Case 2: empty params",
			params: make(map[string]any),
			format: topology.TopologyTree,
		},
		{
			name:   "Case 3: missing key",
			params: map[string]any{"a": "b"},
			format: topology.TopologyTree,
		},
		{
			name:   "Case 4: block topology",
			params: map[string]any{topology.KeyPlugin: topology.TopologyBlock},
			format: topology.TopologyBlock,
		},
		{
			name:   "Case 5: tree topology",
			params: map[string]any{topology.KeyPlugin: topology.TopologyTree},
			format: topology.TopologyTree,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.format, getTopologyFormat(tc.params))
		})
	}
}
