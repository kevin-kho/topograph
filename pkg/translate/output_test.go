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

package translate

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/NVIDIA/topograph/pkg/topology"
)

const (
	testTreeConfig = `SwitchName=S1 Switches=S[2-3]
SwitchName=S2 Nodes=Node[201-202],Node205
SwitchName=S3 Nodes=Node[304-306]
`

	testBlockConfig = `BlockName=B1 Nodes=Node[104-106]
BlockName=B2 Nodes=Node[201-202],Node205
BlockSizes=3
`

	testBlockConfigDiffNumNodes = `BlockName=B1 Nodes=Node[104-106]
BlockName=B2 Nodes=Node[201-202],Node[205-206]
BlockSizes=2
`

	testBlockConfig2 = `BlockName=B3 Nodes=Node[301-303]
BlockName=B4 Nodes=Node[401-403]
BlockName=B1 Nodes=Node[104-106]
BlockName=B2 Nodes=Node[201-202],Node205
BlockSizes=3
`

	testBlockConfigDFS = `BlockName=B1 Nodes=Node202
BlockName=B2 Nodes=Node[104-105]
BlockName=B3 Nodes=Node205
BlockSizes=1
`

	shortNameExpectedResult = `# switch.3.1=hpcislandid-1
SwitchName=switch.3.1 Switches=switch.2.[1-2]
# switch.2.1=network-block-1
SwitchName=switch.2.1 Switches=switch.1.1
# switch.2.2=network-block-2
SwitchName=switch.2.2 Switches=switch.1.2
# switch.1.1=local-block-1
SwitchName=switch.1.1 Nodes=node-1
# switch.1.2=local-block-2
SwitchName=switch.1.2 Nodes=node-2
`
)

func TestToTreeTopology(t *testing.T) {
	v, _ := GetTreeTestSet(false)
	buf := &bytes.Buffer{}
	err := Write(buf, v)
	require.NoError(t, err)
	require.Equal(t, testTreeConfig, buf.String())
}

func TestToBlockTopology(t *testing.T) {
	v, _ := getBlockTestSet()
	buf := &bytes.Buffer{}
	err := Write(buf, v)
	require.NoError(t, err)
	require.Equal(t, testBlockConfig, buf.String())
}

func TestToBlockMultiIBTopology(t *testing.T) {
	v, _ := GetBlockWithMultiIBTestSet()
	buf := &bytes.Buffer{}
	err := Write(buf, v)
	require.NoError(t, err)
	switch buf.String() {
	case testBlockConfig2:
		// nop
	default:
		t.Errorf("unexpected result %s", buf.String())
	}
}

func TestToBlockIBTopology(t *testing.T) {
	v, _ := getBlockWithIBTestSet()
	buf := &bytes.Buffer{}
	err := Write(buf, v)
	require.NoError(t, err)
	switch buf.String() {
	case testBlockConfig:
		// nop
	default:
		t.Errorf("unexpected result %s", buf.String())
	}
}

func TestToBlockDiffNumNode(t *testing.T) {
	v, _ := getBlockWithDiffNumNodeTestSet()
	buf := &bytes.Buffer{}
	err := Write(buf, v)
	require.NoError(t, err)
	switch buf.String() {
	case testBlockConfigDiffNumNodes:
		// nop
	default:
		t.Errorf("unexpected result %s", buf.String())
	}
}

func TestToBlockDFSIBTopology(t *testing.T) {
	v, _ := getBlockWithDFSIBTestSet()
	buf := &bytes.Buffer{}
	err := Write(buf, v)
	require.NoError(t, err)
	switch buf.String() {
	case testBlockConfigDFS:
		// nop
	default:
		t.Errorf("unexpected result %s", buf.String())
	}
}

func TestToSlurmNameShortener(t *testing.T) {
	v := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{
			"hpcislandid-1": {
				ID:   "hpcislandid-1",
				Name: "switch.3.1",
				Vertices: map[string]*topology.Vertex{
					"network-block-1": {
						ID:   "network-block-1",
						Name: "switch.2.1",
						Vertices: map[string]*topology.Vertex{
							"local-block-1": {
								ID:   "local-block-1",
								Name: "switch.1.1",
								Vertices: map[string]*topology.Vertex{
									"node-1": {
										ID:   "node-1-id",
										Name: "node-1",
									},
								},
							},
						},
					},
					"network-block-2": {
						ID:   "network-block-2",
						Name: "switch.2.2",
						Vertices: map[string]*topology.Vertex{
							"local-block-2": {
								ID:   "local-block-2",
								Name: "switch.1.2",
								Vertices: map[string]*topology.Vertex{
									"node-2": {
										ID:   "node-2-id",
										Name: "node-2",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	root := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{topology.TopologyTree: v},
	}

	buf := &bytes.Buffer{}
	err := Write(buf, root)
	require.NoError(t, err)
	require.Equal(t, shortNameExpectedResult, buf.String())
}

func TestCompress(t *testing.T) {
	testCases := []struct {
		name          string
		input, output []string
	}{
		{
			name:   "Case 1: empty list",
			output: []string{},
		},
		{
			name:   "Case 2: ranges",
			input:  []string{"eos0507", "eos0509", "eos0482", "eos0483", "eos0508", "eos0484"},
			output: []string{"eos0[482-484]", "eos0[507-509]"},
		},
		{
			name:   "Case 3: singles",
			input:  []string{"eos0507", "eos0509", "eos0482"},
			output: []string{"eos0482", "eos0507", "eos0509"},
		},
		{
			name:   "Case 4: mix1",
			input:  []string{"eos0507", "eos0509", "abc", "eos0482", "eos0508"},
			output: []string{"abc", "eos0482", "eos0[507-509]"},
		},
		{
			name:   "Case 5: mix2",
			input:  []string{"eos0507", "eos0509", "abc", "eos0508", "eos0482"},
			output: []string{"abc", "eos0482", "eos0[507-509]"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.output, compress(tc.input))
		})
	}
}

func TestSplit(t *testing.T) {
	testCases := []struct {
		name                  string
		input, prefix, suffix string
	}{
		{
			name: "Case 1: empty string",
		},
		{
			name:   "Case 2: no digits",
			input:  "abc",
			prefix: "abc",
		},
		{
			name:   "Case 3: digits only",
			input:  "12345",
			suffix: "12345",
		},
		{
			name:   "Case 4: digits only, leading zeros",
			input:  "0012345",
			prefix: "00",
			suffix: "12345",
		},
		{
			name:   "Case 5: mix",
			input:  "abc1203045",
			prefix: "abc",
			suffix: "1203045",
		},
		{
			name:   "Case 6: mix, leading zeros",
			input:  "abc01203045",
			prefix: "abc0",
			suffix: "1203045",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			prefix, suffix := split(tc.input)
			require.Equal(t, tc.prefix, prefix)
			require.Equal(t, tc.suffix, suffix)
		})
	}
}

func getBlockWithIBTestSet() (*topology.Vertex, map[string]string) {
	//
	//     ibRoot1
	//        |
	//        S1
	//      /    \
	//    S2      S3
	//    |       |
	//   ---     ---
	//   I14\    I21\
	//   I15-B1  I22-B2
	//   I16/    I25/
	//   ---     ---
	//
	instance2node := map[string]string{
		"I14": "Node104", "I15": "Node105", "I16": "Node106",
		"I21": "Node201", "I22": "Node202", "I25": "Node205",
	}

	n14 := &topology.Vertex{ID: "I14", Name: "Node104"}
	n15 := &topology.Vertex{ID: "I15", Name: "Node105"}
	n16 := &topology.Vertex{ID: "I16", Name: "Node106"}

	n21 := &topology.Vertex{ID: "I21", Name: "Node201"}
	n22 := &topology.Vertex{ID: "I22", Name: "Node202"}
	n25 := &topology.Vertex{ID: "I25", Name: "Node205"}

	sw2 := &topology.Vertex{
		ID:       "S2",
		Vertices: map[string]*topology.Vertex{"I14": n14, "I15": n15, "I16": n16},
	}
	sw3 := &topology.Vertex{
		ID:       "S3",
		Vertices: map[string]*topology.Vertex{"I21": n21, "I22": n22, "I25": n25},
	}
	sw1 := &topology.Vertex{
		ID:       "S1",
		Vertices: map[string]*topology.Vertex{"S2": sw2, "S3": sw3},
	}
	treeRoot := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{"S1": sw1},
	}

	block1 := &topology.Vertex{
		ID:       "B1",
		Vertices: map[string]*topology.Vertex{"I14": n14, "I15": n15, "I16": n16},
	}
	block2 := &topology.Vertex{
		ID:       "B2",
		Vertices: map[string]*topology.Vertex{"I21": n21, "I22": n22, "I25": n25},
	}

	blockRoot := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{"B1": block1, "B2": block2},
	}

	root := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{topology.TopologyBlock: blockRoot, topology.TopologyTree: treeRoot},
		Metadata: map[string]string{
			topology.KeyPlugin:     topology.TopologyBlock,
			topology.KeyBlockSizes: "3",
		},
	}
	return root, instance2node
}

func getBlockWithDFSIBTestSet() (*topology.Vertex, map[string]string) {
	//
	//     		 ibRoot1
	//       /      |        \
	//   S1         S2         S3
	//   |          |          |
	//   S4        ---         S5
	//   |         I14\        |
	//  ---			  B2      ---
	//  I22-B1     I15/       I25-B3
	//  ---        ---        ---
	//
	instance2node := map[string]string{
		"I14": "Node104", "I15": "Node105",
		"I22": "Node202", "I25": "Node205",
	}

	n14 := &topology.Vertex{ID: "I14", Name: "Node104"}
	n15 := &topology.Vertex{ID: "I15", Name: "Node105"}

	n22 := &topology.Vertex{ID: "I22", Name: "Node202"}
	n25 := &topology.Vertex{ID: "I25", Name: "Node205"}

	sw2 := &topology.Vertex{
		ID:       "S2",
		Vertices: map[string]*topology.Vertex{"I14": n14, "I15": n15},
	}

	sw4 := &topology.Vertex{
		ID:       "S4",
		Vertices: map[string]*topology.Vertex{"I22": n22},
	}

	sw5 := &topology.Vertex{
		ID:       "S5",
		Vertices: map[string]*topology.Vertex{"I25": n25},
	}

	sw3 := &topology.Vertex{
		ID:       "S3",
		Vertices: map[string]*topology.Vertex{"S5": sw5},
	}
	sw1 := &topology.Vertex{
		ID:       "S1",
		Vertices: map[string]*topology.Vertex{"S4": sw4},
	}

	sw0 := &topology.Vertex{
		ID:       "S0",
		Vertices: map[string]*topology.Vertex{"S1": sw1, "S2": sw2, "S3": sw3},
	}

	treeRoot := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{"S0": sw0},
	}

	block2 := &topology.Vertex{
		ID:       "B2",
		Vertices: map[string]*topology.Vertex{"I14": n14, "I15": n15},
	}
	block1 := &topology.Vertex{
		ID:       "B1",
		Vertices: map[string]*topology.Vertex{"I22": n22},
	}

	block3 := &topology.Vertex{
		ID:       "B3",
		Vertices: map[string]*topology.Vertex{"I25": n25},
	}

	blockRoot := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{"B1": block1, "B2": block2, "B3": block3},
	}

	root := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{topology.TopologyBlock: blockRoot, topology.TopologyTree: treeRoot},
		Metadata: map[string]string{
			topology.KeyPlugin:     topology.TopologyBlock,
			topology.KeyBlockSizes: "1",
		},
	}
	return root, instance2node
}

func getBlockTestSet() (*topology.Vertex, map[string]string) {
	//
	//	---        ---
	//   I14\      I21\
	//   I15-B1    I22-B2
	//   I16/      I25/
	//   ---       ---
	//
	instance2node := map[string]string{
		"I14": "Node104", "I15": "Node105", "I16": "Node106",
		"I21": "Node201", "I22": "Node202", "I25": "Node205",
	}

	n14 := &topology.Vertex{ID: "I14", Name: "Node104"}
	n15 := &topology.Vertex{ID: "I15", Name: "Node105"}
	n16 := &topology.Vertex{ID: "I16", Name: "Node106"}

	n21 := &topology.Vertex{ID: "I21", Name: "Node201"}
	n22 := &topology.Vertex{ID: "I22", Name: "Node202"}
	n25 := &topology.Vertex{ID: "I25", Name: "Node205"}

	block1 := &topology.Vertex{
		ID:       "B1",
		Vertices: map[string]*topology.Vertex{"I14": n14, "I15": n15, "I16": n16},
	}
	block2 := &topology.Vertex{
		ID:       "B2",
		Vertices: map[string]*topology.Vertex{"I21": n21, "I22": n22, "I25": n25},
	}

	blockRoot := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{"B1": block1, "B2": block2},
	}

	root := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{topology.TopologyBlock: blockRoot},
		Metadata: map[string]string{
			topology.KeyPlugin:     topology.TopologyBlock,
			topology.KeyBlockSizes: "3",
		},
	}
	return root, instance2node
}

func getBlockWithDiffNumNodeTestSet() (*topology.Vertex, map[string]string) {
	//
	//     ibRoot1
	//        |
	//        S1
	//      /    \
	//    S2      S3
	//    |       |
	//   ---     ---
	//   I14\    I21\
	//   I15-B1  I22-B2
	//   I16/    I25  /
	//           I26 /
	//   ---     ---
	//
	instance2node := map[string]string{
		"I14": "Node104", "I15": "Node105", "I16": "Node106",
		"I21": "Node201", "I22": "Node202", "I25": "Node205", "I26": "Node206",
	}

	n14 := &topology.Vertex{ID: "I14", Name: "Node104"}
	n15 := &topology.Vertex{ID: "I15", Name: "Node105"}
	n16 := &topology.Vertex{ID: "I16", Name: "Node106"}

	n21 := &topology.Vertex{ID: "I21", Name: "Node201"}
	n22 := &topology.Vertex{ID: "I22", Name: "Node202"}
	n25 := &topology.Vertex{ID: "I25", Name: "Node205"}
	n26 := &topology.Vertex{ID: "I26", Name: "Node206"}

	sw2 := &topology.Vertex{
		ID:       "S2",
		Vertices: map[string]*topology.Vertex{"I14": n14, "I15": n15, "I16": n16},
	}
	sw3 := &topology.Vertex{

		ID:       "S3",
		Vertices: map[string]*topology.Vertex{"I21": n21, "I22": n22, "I25": n25, "I26": n26},
	}
	sw1 := &topology.Vertex{
		ID:       "S1",
		Vertices: map[string]*topology.Vertex{"S2": sw2, "S3": sw3},
	}
	treeRoot := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{"S1": sw1},
	}

	block1 := &topology.Vertex{
		ID:       "B1",
		Vertices: map[string]*topology.Vertex{"I14": n14, "I15": n15, "I16": n16},
	}
	block2 := &topology.Vertex{
		ID:       "B2",
		Vertices: map[string]*topology.Vertex{"I21": n21, "I22": n22, "I25": n25, "I26": n26},
	}

	blockRoot := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{"B1": block1, "B2": block2},
	}

	root := &topology.Vertex{
		Vertices: map[string]*topology.Vertex{topology.TopologyBlock: blockRoot, topology.TopologyTree: treeRoot},
		Metadata: map[string]string{
			topology.KeyPlugin: topology.TopologyBlock,
		},
	}
	return root, instance2node
}
