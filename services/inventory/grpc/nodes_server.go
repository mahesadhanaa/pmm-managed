// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package grpc

import (
	"context"
	"fmt"

	"github.com/percona/pmm/api/inventorypb"

	"github.com/percona/pmm-managed/services/inventory"
)

type nodesServer struct {
	svc *inventory.NodesService
}

// NewNodesServer returns Inventory API handler for managing Nodes.
func NewNodesServer(svc *inventory.NodesService) inventorypb.NodesServer {
	return &nodesServer{svc}
}

// ListNodes returns a list of all Nodes.
func (s *nodesServer) ListNodes(ctx context.Context, req *inventorypb.ListNodesRequest) (*inventorypb.ListNodesResponse, error) {
	nodes, err := s.svc.List(ctx, req)
	if err != nil {
		return nil, err
	}

	res := new(inventorypb.ListNodesResponse)
	for _, node := range nodes {
		switch node := node.(type) {
		case *inventorypb.GenericNode:
			res.Generic = append(res.Generic, node)
		case *inventorypb.ContainerNode:
			res.Container = append(res.Container, node)
		case *inventorypb.RemoteNode:
			res.Remote = append(res.Remote, node)
		case *inventorypb.RemoteAmazonRDSNode:
			res.RemoteAmazonRds = append(res.RemoteAmazonRds, node)
		default:
			panic(fmt.Errorf("unhandled inventory Node type %T", node))
		}
	}
	return res, nil
}

// GetNode returns a single Node by ID.
func (s *nodesServer) GetNode(ctx context.Context, req *inventorypb.GetNodeRequest) (*inventorypb.GetNodeResponse, error) {
	node, err := s.svc.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	res := new(inventorypb.GetNodeResponse)
	switch node := node.(type) {
	case *inventorypb.GenericNode:
		res.Node = &inventorypb.GetNodeResponse_Generic{Generic: node}
	case *inventorypb.ContainerNode:
		res.Node = &inventorypb.GetNodeResponse_Container{Container: node}
	case *inventorypb.RemoteNode:
		res.Node = &inventorypb.GetNodeResponse_Remote{Remote: node}
	case *inventorypb.RemoteAmazonRDSNode:
		res.Node = &inventorypb.GetNodeResponse_RemoteAmazonRds{RemoteAmazonRds: node}
	default:
		panic(fmt.Errorf("unhandled inventory Node type %T", node))
	}
	return res, nil
}

// AddGenericNode adds Generic Node.
func (s *nodesServer) AddGenericNode(ctx context.Context, req *inventorypb.AddGenericNodeRequest) (*inventorypb.AddGenericNodeResponse, error) {
	node, err := s.svc.AddGenericNode(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddGenericNodeResponse{Generic: node}
	return res, nil
}

// AddContainerNode adds Container Node.
func (s *nodesServer) AddContainerNode(ctx context.Context, req *inventorypb.AddContainerNodeRequest) (*inventorypb.AddContainerNodeResponse, error) {
	node, err := s.svc.AddContainerNode(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddContainerNodeResponse{Container: node}
	return res, nil
}

// AddRemoteNode adds Remote Node.
func (s *nodesServer) AddRemoteNode(ctx context.Context, req *inventorypb.AddRemoteNodeRequest) (*inventorypb.AddRemoteNodeResponse, error) {
	node, err := s.svc.AddRemoteNode(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddRemoteNodeResponse{Remote: node}
	return res, nil
}

// AddRemoteAmazonRDSNode adds Amazon (AWS) RDS remote Node.
//nolint:lll
func (s *nodesServer) AddRemoteAmazonRDSNode(ctx context.Context, req *inventorypb.AddRemoteAmazonRDSNodeRequest) (*inventorypb.AddRemoteAmazonRDSNodeResponse, error) {
	node, err := s.svc.AddRemoteAmazonRDSNode(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddRemoteAmazonRDSNodeResponse{RemoteAmazonRds: node}
	return res, nil
}

// RemoveNode removes Node without any Agents and Services.
func (s *nodesServer) RemoveNode(ctx context.Context, req *inventorypb.RemoveNodeRequest) (*inventorypb.RemoveNodeResponse, error) {
	if err := s.svc.Remove(ctx, req.NodeId); err != nil {
		return nil, err
	}

	return new(inventorypb.RemoveNodeResponse), nil
}
