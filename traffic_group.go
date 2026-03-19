package client

// TrafficGroup represents a traffic group.
type TrafficGroup struct {
	ID      string               `json:"id"`
	Name    string               `json:"name"`
	Targets []TrafficGroupTarget `json:"targets"`
}

// TrafficGroupTarget represents a target within a traffic group.
type TrafficGroupTarget struct {
	TargetType string `json:"target_type"`
	TargetID   string `json:"target_id"`
}

type trafficGroupList struct {
	TrafficGroups []*TrafficGroup `json:"traffic_groups"`
}
