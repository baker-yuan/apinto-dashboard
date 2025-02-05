package cluster

import (
	"context"

	v1 "github.com/eolinker/apinto-dashboard/client/v1"
	"github.com/eolinker/apinto-dashboard/driver"
	"github.com/eolinker/apinto-dashboard/modules/cluster/cluster-dto"
	"github.com/eolinker/apinto-dashboard/modules/cluster/cluster-entry"
	"github.com/eolinker/apinto-dashboard/modules/cluster/cluster-model"
)

const (
	ProviderName = "cluster"
)

type IApintoClient interface {
	GetClient(ctx context.Context, clusterId int) (v1.IClient, error)
}
type IClusterCertificateService interface {
	Insert(ctx context.Context, operator, namespaceId int, clusterName, key, pem string) error
	Update(ctx context.Context, operator, namespaceId, certificateId int, clusterName, key, pem string) error
	Info(ctx context.Context, namespaceId, certificateId int, clusterName string) (*cluster_model.Certificate, error)
	QueryList(ctx context.Context, namespaceId int, clusterName string) ([]*cluster_model.ClusterCertificate, error)
	DeleteById(ctx context.Context, namespaceId int, clusterName string, id int) error
}

type IClusterService interface {
	GetAllCluster(ctx context.Context) ([]*cluster_model.Cluster, error)
	Count(ctx context.Context) (int, error)
	SimpleCluster(ctx context.Context, namespaceId int) ([]*cluster_model.ClusterSimple, error)
	CheckByNamespaceByName(ctx context.Context, namespaceId int, name string) (int, error)
	GetByClusterId(ctx context.Context, clusterId int) (*cluster_model.Cluster, error)
	GetByNamespaceByName(ctx context.Context, namespaceId int, name string) (*cluster_model.Cluster, error)
	GetByNamespaceId(ctx context.Context, namespaceId int) ([]*cluster_model.Cluster, error)
	GetByNames(ctx context.Context, namespaceId int, names []string) ([]*cluster_model.Cluster, error)
	GetByUUIDs(ctx context.Context, namespaceId int, uuids []string) ([]*cluster_model.Cluster, error)
	Insert(ctx context.Context, namespaceId, userId int, clusterInput *cluster_dto.ClusterInput) error
	QueryByNamespaceId(ctx context.Context, namespaceId int, clusterName string) (*cluster_model.Cluster, error)
	QueryListByNamespaceId(ctx context.Context, namespaceId int) ([]*cluster_model.Cluster, error)
	QueryListByClusterNames(ctx context.Context, namespaceId int, clusterNames []string) ([]*cluster_model.Cluster, error)

	DeleteByNamespaceIdByName(ctx context.Context, namespaceId, userId int, name string) error
	Update(ctx context.Context, namespaceId, userId int, name string, clusterInput *cluster_dto.ClusterInput) error
	UpdateAddr(ctx context.Context, userId, clusterId int, addr, uuid string) error
	ClusterCount(ctx context.Context, namespaceId int) (int64, error)
}

type IClusterNodeService interface {
	QueryList(ctx context.Context, namespaceId int, clusterName string) ([]*cluster_model.ClusterNode, bool, error)
	List(ctx context.Context, namespaceId int, clusterName string) ([]*cluster_model.Node, error)
	QueryByClusterId(ctx context.Context, id int) ([]*cluster_model.Node, error)
	QueryAllCluster(ctx context.Context) ([]*cluster_model.Node, error)
	QueryAdminAddrByClusterId(ctx context.Context, id int) ([]string, error)
	Reset(ctx context.Context, namespaceId, userId int, clusterName, clusterAddr, source string) error
	Update(ctx context.Context, namespaceId int, clusterName string) error
	Delete(ctx context.Context, namespaceId int, clusterId int) error
	NodeRepeatContrast(ctx context.Context, namespaceId, clusterId int, newList []*cluster_model.Node) error
	Insert(ctx context.Context, nodes []*cluster_model.Node) error
	GetNodesByUrl(addr string) ([]*cluster_model.Node, error)
	GetClusterInfo(addr string) (*v1.ClusterInfo, error)
}

// ICLConfigDriverManager 集群配置驱动管理器
type ICLConfigDriverManager interface {
	driver.IDriverManager[ICLConfigDriver]
	List() []*driver.DriverInfo
}

const (
	CLConfigRedis    = "redis"
	CLConfigInfluxV2 = "influxdbv2"
)

type ICLConfigDriver interface {
	CheckInput(config []byte) error
	ToApinto(name string, config []byte) interface{}
	FormatOut(operator string, config *cluster_entry.ClusterConfig) interface{}
	InitConfig(config []byte) error
}
