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

package gcp

import (
	"context"
	"fmt"
	"strings"
	"time"

	compute_v1 "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"cloud.google.com/go/compute/metadata"
	"google.golang.org/api/iterator"

	"github.com/NVIDIA/topograph/pkg/common"
)

type InstanceTopology struct {
	instances []*InstanceInfo
}

type InstanceInfo struct {
	clusterID string
	rackID    string
	name      string
}

func GenerateInstanceTopology(ctx context.Context, _ interface{}, instanceToNodeMap map[string]string) (*InstanceTopology, error) {
	zoneClient, err := compute_v1.NewZonesRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get zones client: %s", err.Error())
	}
	projectID, err := metadata.ProjectID()
	if err != nil {
		return nil, fmt.Errorf("unable to get project ID: %s", err.Error())
	}
	listZoneRequest := computepb.ListZonesRequest{Project: projectID}
	zones := make([]string, 0)

	timeNow := time.Now()
	res := zoneClient.List(ctx, &listZoneRequest)
	requestLatency.WithLabelValues("ListZones").Observe(time.Since(timeNow).Seconds())

	for {
		zone, err := res.Next()
		if err == iterator.Done {
			break
		}
		zones = append(zones, *zone.Name)
	}

	instanceClient, err := compute_v1.NewInstancesRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to instance client for zone: %s", err.Error())
	}

	instanceTopology := &InstanceTopology{instances: make([]*InstanceInfo, 0)}

	for _, zone := range zones {
		timeNow := time.Now()
		listInstanceRequest := computepb.ListInstancesRequest{Project: projectID, Zone: zone}
		requestLatency.WithLabelValues("ListInstances").Observe(time.Since(timeNow).Seconds())

		resInstance := instanceClient.List(ctx, &listInstanceRequest)
		for {
			instance, err := resInstance.Next()
			if err == iterator.Done {
				break
			}
			_, isNodeInCluster := instanceToNodeMap[*instance.Name]

			if instance.ResourceStatus == nil {
				resourceStatusNotFound.WithLabelValues(*instance.Name).Set(1)
				continue
			}
			resourceStatusNotFound.WithLabelValues(*instance.Name).Set(0)

			if instance.ResourceStatus.PhysicalHost == nil {
				physicalHostNotFound.WithLabelValues(*instance.Name).Set(1)
				continue
			}
			physicalHostNotFound.WithLabelValues(*instance.Name).Set(0)

			if isNodeInCluster {
				tokens := strings.Split(*instance.ResourceStatus.PhysicalHost, "/")
				physicalHostIDChunks.WithLabelValues(*instance.Name).Set(float64(getTokenCount(tokens)))
				instanceObj := &InstanceInfo{
					name:      *instance.Name,
					clusterID: tokens[1],
					rackID:    tokens[2],
				}
				instanceTopology.instances = append(instanceTopology.instances, instanceObj)
			}
		}
	}

	return instanceTopology, nil
}

func (cfg *InstanceTopology) toSLURM() (*common.Vertex, error) {
	forest := make(map[string]*common.Vertex)
	nodes := make(map[string]*common.Vertex)

	for _, c := range cfg.instances {
		instance := &common.Vertex{
			Name: c.name,
			ID:   c.name,
		}

		id2 := c.rackID
		sw2, ok := nodes[id2]
		if !ok {
			sw2 = &common.Vertex{
				ID:       id2,
				Vertices: make(map[string]*common.Vertex),
			}
			nodes[id2] = sw2
		}
		sw2.Vertices[instance.ID] = instance

		id1 := c.clusterID
		sw1, ok := nodes[id1]
		if !ok {
			sw1 = &common.Vertex{
				ID:       id1,
				Vertices: make(map[string]*common.Vertex),
			}
			nodes[id1] = sw1
			forest[id1] = sw1
		}
		sw1.Vertices[id2] = sw2
	}

	root := &common.Vertex{
		Vertices: make(map[string]*common.Vertex),
	}
	for name, node := range forest {
		root.Vertices[name] = node
	}

	return root, nil
}

func getTokenCount(tokens []string) int {
	c := 0
	for _, q := range tokens {
		if len(q) > 0 {
			c += 1
		}
	}
	return c
}