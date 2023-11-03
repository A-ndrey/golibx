package dotlang

import (
	"errors"
	"fmt"
	"strings"
)

type ID string

func (id ID) Build(sb *strings.Builder) {
	if id == "" {
		return
	}

	sb.WriteString(string(id))
	sb.WriteByte(' ')
}

type Graph struct {
	Strict    bool
	DiGraph   bool
	ID        ID
	Nodes     []NodeStmt
	Edges     []EdgeStmt
	Attrs     []*AttrStmt
	SubGraphs []*SubGraph
}

func (g *Graph) Validate() error {
	if g == nil {
		return nil
	}

	for i, n := range g.Nodes {
		if err := n.Validate(); err != nil {
			return fmt.Errorf("graph -> [%d] %w", i, err)
		}
	}

	for i, e := range g.Edges {
		if err := e.Validate(); err != nil {
			return fmt.Errorf("graph -> [%d] %w", i, err)
		}
	}

	for i, a := range g.Attrs {
		if err := a.Validate(); err != nil {
			return fmt.Errorf("graph -> [%d] %w", i, err)
		}
	}

	for i, s := range g.SubGraphs {
		if err := s.Validate(); err != nil {
			return fmt.Errorf("graph -> [%d] %w", i, err)
		}
	}

	return nil
}

func (g *Graph) Build(sb *strings.Builder) {
	if g == nil {
		return
	}

	if g.Strict {
		sb.WriteString("strict")
		sb.WriteByte(' ')
	}
	if g.DiGraph {
		sb.WriteString("digraph")
		sb.WriteByte(' ')
	} else {
		sb.WriteString("graph")
		sb.WriteByte(' ')
	}
	g.ID.Build(sb)
	sb.WriteByte('{')
	sb.WriteByte(' ')
	for _, n := range g.Nodes {
		n.Build(sb)
	}
	for _, e := range g.Edges {
		e.Build(sb, g.DiGraph)
	}
	for _, a := range g.Attrs {
		a.Build(sb)
	}
	for _, s := range g.SubGraphs {
		s.Build(sb, g.DiGraph)
	}
	sb.WriteByte('}')
}

type SubGraph struct {
	ID        ID
	Nodes     []NodeStmt
	Edges     []EdgeStmt
	Attrs     []*AttrStmt
	SubGraphs []*SubGraph
}

func (sg *SubGraph) Validate() error {
	if sg == nil {
		return nil
	}

	for i, n := range sg.Nodes {
		if err := n.Validate(); err != nil {
			return fmt.Errorf("subgraph -> [%d] %w", i, err)
		}
	}

	for i, e := range sg.Edges {
		if err := e.Validate(); err != nil {
			return fmt.Errorf("subgraph -> [%d] %w", i, err)
		}
	}

	for i, a := range sg.Attrs {
		if err := a.Validate(); err != nil {
			return fmt.Errorf("subgraph -> [%d] %w", i, err)
		}
	}

	for i, s := range sg.SubGraphs {
		if err := s.Validate(); err != nil {
			return fmt.Errorf("subgraph -> [%d] %w", i, err)
		}
	}

	return nil
}

func (sg *SubGraph) Build(sb *strings.Builder, digraph bool) {
	if sg == nil {
		return
	}
	if sg.ID != "" {
		sb.WriteString("subgraph")
		sb.WriteByte(' ')
		sg.ID.Build(sb)
	}
	sb.WriteByte('{')
	sb.WriteByte(' ')
	for _, n := range sg.Nodes {
		n.Build(sb)
	}
	for _, e := range sg.Edges {
		e.Build(sb, digraph)
	}
	for _, a := range sg.Attrs {
		a.Build(sb)
	}
	for _, s := range sg.SubGraphs {
		s.Build(sb, digraph)
	}
	sb.WriteByte('}')
	sb.WriteByte(' ')
}

type NodeStmt struct {
	ID       NodeID
	AttrList AttrList
}

func (n NodeStmt) Validate() error {
	emptyID := NodeID{}
	if n.ID == emptyID {
		return errors.New("node -> ID must be specified")
	}

	if err := n.ID.Validate(); err != nil {
		return fmt.Errorf("node -> %w", err)
	}

	return nil
}

func (n NodeStmt) Build(sb *strings.Builder) {
	n.ID.Build(sb)
	n.AttrList.Build(sb)
}

type EdgeStmt struct {
	FromID   NodeID
	FromSG   *SubGraph
	EdgeRHS  *EdgeRHS
	AttrList AttrList
}

func (e EdgeStmt) Validate() error {
	emptyID := NodeID{}

	if e.FromID == emptyID && e.FromSG == nil {
		return errors.New("edge -> source must be specified")
	}

	if e.FromID != emptyID && e.FromSG != nil {
		return errors.New("edge -> only one source must be specified")
	}

	if e.FromID != emptyID {
		if err := e.FromID.Validate(); err != nil {
			return fmt.Errorf("edge -> %w", err)
		}
	}

	if e.FromSG != nil {
		if err := e.FromSG.Validate(); err != nil {
			return fmt.Errorf("edge -> %w", err)
		}
	}

	if err := e.EdgeRHS.Validate(); err != nil {
		return fmt.Errorf("edge -> %w", err)
	}

	return nil
}

func (e EdgeStmt) Build(sb *strings.Builder, digraph bool) {
	e.FromID.Build(sb)
	e.FromSG.Build(sb, digraph)
	e.EdgeRHS.Build(sb, digraph)
	e.AttrList.Build(sb)
}

type EdgeRHS struct {
	ToID    NodeID
	ToSG    *SubGraph
	EdgeRHS *EdgeRHS
}

func (erhs *EdgeRHS) Validate() error {
	if erhs == nil {
		return nil
	}

	emptyID := NodeID{}

	if erhs.ToID == emptyID && erhs.ToSG == nil {
		return errors.New("edge_rhs -> destination must be specified")
	}

	if erhs.ToID != emptyID && erhs.ToSG != nil {
		return errors.New("edge_rhs -> only one destination must be specified")
	}

	if erhs.ToID != emptyID {
		if err := erhs.ToID.Validate(); err != nil {
			return fmt.Errorf("edge_rhs -> %w", err)
		}
	}

	if erhs.ToSG != nil {
		if err := erhs.ToSG.Validate(); err != nil {
			return fmt.Errorf("edge_rhs -> %w", err)
		}
	}

	if err := erhs.EdgeRHS.Validate(); err != nil {
		return fmt.Errorf("edge_rhs -> %w", err)
	}

	return nil
}

func (erhs *EdgeRHS) Build(sb *strings.Builder, digraph bool) {
	if erhs == nil {
		return
	}

	if digraph {
		sb.WriteString("->")
	} else {
		sb.WriteString("--")
	}
	sb.WriteByte(' ')

	erhs.ToID.Build(sb)
	erhs.ToSG.Build(sb, digraph)
	erhs.EdgeRHS.Build(sb, digraph)
}

type AttrStmt struct {
	Graph    bool
	Node     bool
	Edge     bool
	AttrList AttrList
}

func (a *AttrStmt) Validate() error {
	if a == nil {
		return nil
	}

	if !(a.Graph || a.Node || a.Edge) {
		return errors.New("attr_stmt -> attr_stmt type must be specified")
	}

	if a.Graph && a.Node || a.Graph && a.Edge || a.Node && a.Edge {
		return errors.New("attr_stmt -> only one attr_stmt type must be specified")
	}

	return nil
}

func (a *AttrStmt) Build(sb *strings.Builder) {
	if a.Graph {
		sb.WriteString("graph ")
	} else if a.Node {
		sb.WriteString("node ")
	} else if a.Edge {
		sb.WriteString("edge ")
	}
	a.AttrList.Build(sb)
}

type AttrList map[ID]ID

func (al AttrList) Build(sb *strings.Builder) {
	if len(al) == 0 {
		return
	}
	sb.WriteByte('[')
	for k, v := range al {
		sb.WriteByte(' ')
		k.Build(sb)
		sb.WriteByte('=')
		sb.WriteByte(' ')
		v.Build(sb)
	}
	sb.WriteByte(' ')
	sb.WriteByte(']')
	sb.WriteByte(' ')
}

type NodeID struct {
	ID   ID
	Port Port
}

func (nid NodeID) Validate() error {
	if nid.ID == "" {
		return errors.New("nodeID -> ID must be specified")
	}
	if err := nid.Port.Validate(); err != nil {
		if nid.ID == "" {
			return fmt.Errorf("nodeID -> %w", err)
		} else {
			return fmt.Errorf("nodeID %s -> %w", nid.ID, err)
		}
	}

	return nil
}

func (nid NodeID) Build(sb *strings.Builder) {
	empty := NodeID{}
	if nid == empty {
		return
	}

	nid.ID.Build(sb)
	nid.Port.Build(sb)
}

type Port struct {
	ID        ID
	CompassPt CompassPt
}

func (p Port) Validate() error {
	if err := p.CompassPt.Validate(); err != nil {
		if p.ID == "" {
			return fmt.Errorf("port -> %w", err)
		}
		return fmt.Errorf("port %s -> %w", p.ID, err)
	}

	return nil
}

func (p Port) Build(sb *strings.Builder) {
	if p.ID != "" {
		sb.WriteByte(':')
		p.ID.Build(sb)
	}
	if p.CompassPt != CompassDefault {
		sb.WriteByte(':')
		p.CompassPt.Build(sb)
	}
}

type CompassPt uint8

const (
	CompassDefault CompassPt = iota
	CompassN
	CompassNE
	CompassE
	CompassSE
	CompassS
	CompassSW
	CompassW
	CompassNW
	CompassC
)

func (cpt CompassPt) Validate() error {
	if cpt > CompassC {
		return errors.New("unknown CompassPt")
	}

	return nil
}

func (cpt CompassPt) Build(sb *strings.Builder) {
	switch cpt {
	case CompassN:
		sb.WriteString("n")
	case CompassNE:
		sb.WriteString("ne")
	case CompassE:
		sb.WriteString("e")
	case CompassSE:
		sb.WriteString("se")
	case CompassS:
		sb.WriteString("s")
	case CompassSW:
		sb.WriteString("sw")
	case CompassW:
		sb.WriteString("w")
	case CompassNW:
		sb.WriteString("nw")
	case CompassC:
		sb.WriteString("c")
	default:
		return
	}

	sb.WriteByte(' ')
}
