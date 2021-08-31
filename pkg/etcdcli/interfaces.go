package etcdcli

import (
	"context"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/etcdserverpb"
)

const (
	EtcdMemberStatusAvailable  = "EtcdMemberAvailable"
	EtcdMemberStatusNotStarted = "EtcdMemberNotStarted"
	EtcdMemberStatusUnhealthy  = "EtcdMemberUnhealthy"
	EtcdMemberStatusUnknown    = "EtcdMemberUnknown"
)

type EtcdClient interface {
	MemberAdder
	MemberHealth
	IsMemberHealthy
	MemberLister
	MemberRemover
	HealthyMemberLister
	UnhealthyMemberLister
	MemberStatusChecker

	GetMember(name string) (*etcdserverpb.Member, error)
	MemberUpdatePeerURL(id uint64, peerURL []string) error
}

type Status interface {
	Status(ctx context.Context, target string) (*clientv3.StatusResponse, error)
}

type MemberAdder interface {
	MemberAdd(peerURL string) error
}

type MemberHealth interface {
	MemberHealth() (memberHealth, error)
}

type IsMemberHealthy interface {
	IsMemberHealthy(member *etcdserverpb.Member) (bool, error)
}

type MemberRemover interface {
	MemberRemove(member string) error
}

type MemberLister interface {
	MemberList() ([]*etcdserverpb.Member, error)
}

type HealthyMemberLister interface {
	HealthyMembers(ctx context.Context) ([]*etcdserverpb.Member, error)
}

type UnhealthyMemberLister interface {
	UnhealthyMembers() ([]*etcdserverpb.Member, error)
}

type MemberStatusChecker interface {
	MemberStatus(member *etcdserverpb.Member) string
}
