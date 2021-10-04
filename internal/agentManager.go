package internal

import (
	"errors"
	"sync"
)

type AgentManager struct {
	lock   sync.Mutex
	agents map[string][]string
	Wdogs  map[string]*Watchdog
}

func NewAgentManager() *AgentManager {
	m := make(map[string][]string)
	d := make(map[string]*Watchdog)
	return &AgentManager{
		agents: m,
		Wdogs:  d,
	}
}

func (m *AgentManager) Get(agentID string) ([]string, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if l, ok := m.agents[agentID]; ok {
		res := make([]string, 0, len(l))
		res = append(res, l...)
		return res, nil
	} else {
		return nil, errors.New("unknown agent id")
	}
}

func (m *AgentManager) Set(agentID string, addrs []string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.agents[agentID] = addrs
}
