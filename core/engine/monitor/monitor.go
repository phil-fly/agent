package monitor

import "agent/core/engine/docker"

type MonitorEngineIf interface {
	EventRead()(error)
	OpenMonitor()(error)
	SetMmapFile(string)
	SetDockerKnow(*docker.DockerKnow)
	SetRuleEngine(interface{})
}

type MonitorEngine struct {
	evevtEngine MonitorEngineIf
}

func (c *MonitorEngine) SetMmapFile(file string) {
	c.evevtEngine.SetMmapFile(file)
}

func (c *MonitorEngine) SetDockerKnow(DockerKnow *docker.DockerKnow) {
	c.evevtEngine.SetDockerKnow(DockerKnow)
}

func (c *MonitorEngine) SetRuleEngine(RuleEngines interface{}) {
	c.evevtEngine.SetRuleEngine(RuleEngines)
}

func (c *MonitorEngine) MonitorOpen() error {
	return c.evevtEngine.OpenMonitor()
}

func (c *MonitorEngine) MonitorEventRead() error {
	return c.evevtEngine.EventRead()
}



func NewMonitorEngine(s MonitorEngineIf) *MonitorEngine {
	return &MonitorEngine{evevtEngine: s}
}