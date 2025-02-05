package api

import (
	"context"

	"github.com/eolinker/apinto-dashboard/client/v1"
	"github.com/eolinker/apinto-dashboard/driver"
	"github.com/eolinker/apinto-dashboard/modules/api/api-dto"
	api_entry "github.com/eolinker/apinto-dashboard/modules/api/api-entry"
	apimodel "github.com/eolinker/apinto-dashboard/modules/api/model"
	"github.com/eolinker/apinto-dashboard/modules/base/frontend-model"
	"github.com/eolinker/apinto-dashboard/modules/group/group-model"
	"github.com/eolinker/apinto-dashboard/modules/strategy/strategy-model"
)

type IAPIService interface {
	GetAPIList(ctx context.Context, namespaceID int, groupUUID, searchName string, searchSources []string, pageNum, pageSize int) ([]*apimodel.APIListItem, int, error)
	GetAPICountByGroupUUID(ctx context.Context, namespaceID int, groupUUID string) int64
	GetAPIVersionInfo(ctx context.Context, namespaceID int, uuid string) (*apimodel.APIVersionInfo, error)
	GetAPIInfo(ctx context.Context, namespaceID int, uuid string) (*apimodel.APIInfo, error)
	GetAPIInfoById(ctx context.Context, id int) (*apimodel.APIInfo, error)
	GetAPIInfoByGroupUUID(ctx context.Context, namespaceID int, groupUUID string) ([]*apimodel.APIInfo, error)
	GetAPIInfoByUUIDS(ctx context.Context, namespaceID int, uuids []string) ([]*apimodel.APIInfo, error)
	GetAPIInfoByPath(ctx context.Context, namespaceID int, path string) ([]*apimodel.APIInfo, error)
	GetAPIListItemByUUIDS(ctx context.Context, namespaceID int, uuids []string) ([]*apimodel.APIListItem, error)

	GetAPIInfoAll(ctx context.Context, namespaceID int) ([]*apimodel.APIInfo, error)

	GetAPIListItemAll(ctx context.Context, namespaceID int) ([]*apimodel.APIListItem, error)
	GetAPIsForSync(ctx context.Context, namespaceID int) ([]*apimodel.APIVersionInfo, error)

	CreateAPI(ctx context.Context, namespaceID int, operator int, input *api_dto.APIInfo) (string, int, error)
	UpdateAPI(ctx context.Context, namespaceID int, operator int, input *api_dto.APIInfo) error
	DeleteAPI(ctx context.Context, namespaceId, operator int, uuid string) error
	GetGroups(ctx context.Context, namespaceId int, parentUuid, queryName string) (*group_model.CommonGroupRoot, []*group_model.CommonGroupApi, error)

	BatchOnline(ctx context.Context, namespaceId int, operator int, onlineToken string) ([]*apimodel.BatchListItem, error)
	BatchOffline(ctx context.Context, namespaceId int, operator int, apiUUIDs, clusterNames []string) ([]*apimodel.BatchListItem, error)
	BatchOnlineCheck(ctx context.Context, namespaceId int, operator int, apiUUIDs, clusterNames []string) ([]*apimodel.BatchOnlineCheckListItem, string, error)

	OnlineInfo(ctx context.Context, namespaceId int, uuid string) (*apimodel.APIVersionInfo, []*apimodel.APIOnlineListItem, error)
	OnlineList(ctx context.Context, namespaceId int, uuid string) ([]*apimodel.APIOnlineListItem, error)
	OnlineAPI(ctx context.Context, namespaceId, operator int, uuid string, clusterNames []string) ([]*frontend_model.Router, error)
	OfflineAPI(ctx context.Context, namespaceId, operator int, uuid string, clusterNames []string) ([]*apimodel.BatchListItem, error)

	GetSource(ctx context.Context) ([]*apimodel.SourceListItem, error)
	GetImportCheckList(ctx context.Context, namespaceId int, fileData []byte, groupID, serviceName, requestPrefix string) ([]*apimodel.ImportAPIListItem, string, error)
	ImportAPI(ctx context.Context, namespaceId, operator int, input *api_dto.ImportAPIInfos) error

	GetAPIListByName(ctx context.Context, namespaceId int, name string) ([]*group_model.CommonGroupApi, error)
	GetAPIListByServiceName(ctx context.Context, namespaceId int, serviceName []string) ([]*apimodel.APIInfo, error)
	GetLatestAPIVersion(ctx context.Context, apiId int) (*api_entry.APIVersion, error)
	IsAPIOnline(ctx context.Context, clusterName, clusterAddr string, apiUUID int) bool
	GetAPIDriver(driverName string) IAPIDriver
	GetAPINameByID(ctx context.Context, apiID int) (string, error)
	GetAPIRemoteOptions(ctx context.Context, namespaceId, pageNum, pageSize int, keyword, groupUuid string) ([]any, int, error)
	GetAPIRemoteByUUIDS(ctx context.Context, namespace int, uuids []string) ([]*strategy_model.RemoteApis, error)
	APICount(ctx context.Context, namespaceId int) (int64, error)
	APIOnlineCount(ctx context.Context, namespaceId int) (int64, error)
}

type IAPIDriverManager interface {
	driver.IDriverManager[IAPIDriver]
	List() []*APIDriverInfo
}

type APIDriverInfo struct {
	Name string
}

type IAPIDriver interface {
	CheckInput(input *api_dto.APIInfo) error
	ToApinto(name, desc string, disable bool, method []string, requestPath, requestPathLabel, proxyPath, serviceName string, timeout, retry int, hosts []string, match []*api_entry.MatchConf, header []*api_entry.ProxyHeader, templateUUID string) *v1.RouterConfig
}
