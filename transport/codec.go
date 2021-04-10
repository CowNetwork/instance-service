package transport

import (
	instancev1 "github.com/cownetwork/instance-controller/api/v1"
	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
)

func instanceToProto(instance *instancev1.Instance) *instanceapiv1.Instance {
	proto := &instanceapiv1.Instance{
		Id:    "", // TODO
		Name:  instance.Name,
		Ip:    instance.Status.IP,
		State: toAPIState(instance.Status.State),
		Metadata: &instanceapiv1.Metadata{
			State: string(instance.Status.Metadata.State),
		},
	}
	players := make([]*instanceapiv1.Player, 0)
	for _, p := range instance.Status.Metadata.Players {
		players = append(players, toAPIPlayer(p))
	}
	proto.Metadata.Players = players
	return proto
}

func toAPIPlayer(player instancev1.InstancePlayer) *instanceapiv1.Player {
	return &instanceapiv1.Player{
		Id:       player.ID,
		Metadata: string(player.Metadata),
	}
}

func toAPIState(state instancev1.InstanceState) instanceapiv1.Instance_State {
	switch state {
	case instancev1.StateInitializing:
		return instanceapiv1.Instance_STATE_INITIALIZING
		break
	case instancev1.StateRunning:
		return instanceapiv1.Instance_STATE_RUNNING
		break
	case instancev1.StateEnding:
		return instanceapiv1.Instance_STATE_ENDING
		break
	}
	return instanceapiv1.Instance_STATE_UNKNOWN
}
