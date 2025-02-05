package driver

import (
	"errors"
	"fmt"
	v1 "github.com/eolinker/apinto-dashboard/client/v1"
	"github.com/eolinker/apinto-dashboard/common"
	"github.com/eolinker/apinto-dashboard/enum"
	api_dto "github.com/eolinker/apinto-dashboard/modules/api/api-dto"
	api_entry "github.com/eolinker/apinto-dashboard/modules/api/api-entry"
	"net/textproto"
	"strings"
)

const apintoRestfulRegexp = "([0-9a-zA-Z-_]+)"

// formatProxyPath 格式化转发路径上，用于转发重写插件正则替换 比如 请求路径 /path/{A}/{B} 原转发路径：/path/{B}  格式化后 新转发路径： /path/$2
func formatProxyPath(requestPath, proxyPath string) string {
	restfulSet := make(map[string]string)
	newProxyPath := proxyPath
	rList := strings.Split(requestPath, "/")
	i := 1
	for _, param := range rList {
		if common.IsRestfulParam(param) {
			restfulSet[param] = fmt.Sprintf("$%d", i)
			i += 1
		}
	}

	for param, order := range restfulSet {
		newProxyPath = strings.ReplaceAll(newProxyPath, param, order)
	}
	return newProxyPath
}

func checkInput(input *api_dto.APIInfo) error {
	input.UUID = strings.TrimSpace(input.UUID)
	if input.UUID != "" {
		err := common.IsMatchString(common.UUIDExp, input.UUID)
		if err != nil {
			return err
		}
	}

	if input.GroupUUID == "" {
		return errors.New("group_uuid can't be nil. ")
	}

	if input.Method == nil {
		input.Method = []string{}
	}

	if input.Match == nil {
		input.Match = []*api_entry.MatchConf{}
	}

	if input.Header == nil {
		input.Header = []*api_entry.ProxyHeader{}
	}

	input.ServiceName = strings.TrimSpace(input.ServiceName)
	if input.ServiceName == "" {
		return errors.New("service_name can't be nil. ")
	}

	var err error
	//格式化,并且校验请求路径
	input.RequestPath = "/" + strings.TrimPrefix(strings.TrimSpace(input.RequestPath), "/")
	if input.RequestPath == "/" {
		return errors.New("request_path can't be / . ")
	}

	if input.RequestPath, err = common.CheckAndFormatPath(input.RequestPath); err != nil {
		return errors.New("request_path is illegal. ")
	}

	requestParamSet := make(map[string]struct{})
	for _, param := range strings.Split(input.RequestPath, "/") {
		if common.IsRestfulParam(param) {
			requestParamSet[param] = struct{}{}
		}
	}

	//格式化转发路径, 并校验转发路径， 若转发路径出现restful参数，则必须对得上请求路径出现过得restful参数
	input.ProxyPath = "/" + strings.TrimPrefix(strings.TrimSpace(input.ProxyPath), "/")
	if input.ProxyPath, err = common.CheckAndFormatPath(input.ProxyPath); err != nil {
		return errors.New("proxy_path is illegal. ")
	}
	for _, param := range strings.Split(input.ProxyPath, "/") {
		if common.IsRestfulParam(param) {
			_, exist := requestParamSet[param]
			if !exist {
				return fmt.Errorf("proxyPath is illegal. restful param %s is invalid. ", param)
			}
		}
	}

	if input.Timeout <= 0 {
		return errors.New("timeout must be greater than 0. ")
	}

	if input.Retry < 0 {
		return errors.New("retry can't be less than 0. ")
	}

	for i, m := range input.Match {
		input.Match[i].Key = strings.TrimSpace(m.Key)
		if input.Match[i].Key == "" {
			return errors.New("Match.Key can't be nil. ")
		}
		switch m.Position {
		case enum.MatchPositionHeader, enum.MatchPositionQuery, enum.MatchPositionCookie:
		default:
			return fmt.Errorf("position %s is illegal. ", m.Position)
		}

		switch m.MatchType {
		case enum.MatchTypeEqual, enum.MatchTypePrefix, enum.MatchTypeSuffix, enum.MatchTypeSubstr, enum.MatchTypeUneuqal, enum.MatchTypeRegexp, enum.MatchTypeRegexpG:
			input.Match[i].Pattern = strings.TrimSpace(m.Pattern)
			if input.Match[i].Pattern == "" {
				return fmt.Errorf("Match.Pattern can't be nil when MatchType is %s. ", m.MatchType)
			}
		case enum.MatchTypeNull, enum.MatchTypeExist, enum.MatchTypeUnexist, enum.MatchTypeAny:
		default:
			return fmt.Errorf("match_type %s is illegal. ", m.MatchType)
		}
	}

	for i, h := range input.Header {
		input.Header[i].Key = strings.TrimSpace(h.Key)
		if input.Header[i].Key == "" {
			return errors.New("Header.Key can't be nil. ")
		}
		switch h.OptType {
		case enum.HeaderOptTypeAdd:
			input.Header[i].Value = strings.TrimSpace(h.Value)
			if input.Header[i].Value == "" {
				return fmt.Errorf("Header.Value can't be nil when OptType is %s. ", h.OptType)
			}
		case enum.HeaderOptTypeDelete:
		default:
			return fmt.Errorf("opt_type %s is illegal. ", h.OptType)
		}
	}

	//设置前端显示所需要的请求路径， 用户输入什么就显示什么
	input.RequestPathLabel = input.RequestPath
	//校验请求路径是不是restful， 如果是则将request_path转换成统一的restful请求路径，如/path/{A} -> /path/{rest}
	if common.IsRestfulPath(input.RequestPath) {
		input.RequestPath = common.ReplaceRestfulPath(input.RequestPath, enum.RestfulLabel)
	}

	if len(input.Hosts) > 0 {
		for _, host := range input.Hosts {
			if !common.IsMatchDomainPort(host) && !common.IsMatchIpPort(host) {
				return fmt.Errorf("host %s is illegal. ", host)
			}
		}
	}

	return nil
}

func toApinto(name, desc string, disable bool, method []string, requestPath, requestPathLabel, proxyPath, serviceName string, timeout, retry int, hosts []string, match []*api_entry.MatchConf, header []*api_entry.ProxyHeader, templateUUID string) *v1.RouterConfig {

	rewriteHeaders := make(map[string]string)
	for _, ph := range header {
		//格式化header Key
		ph.Key = textproto.CanonicalMIMEHeaderKey(ph.Key)
		switch ph.OptType {
		case enum.HeaderOptTypeAdd:
			rewriteHeaders[ph.Key] = ph.Value
		case enum.HeaderOptTypeDelete:
			rewriteHeaders[ph.Key] = ""
		}
	}

	rewritePlugin := v1.PluginProxyRewriteV2Config{
		NotMatchErr: true,
		HostRewrite: false,
		Headers:     rewriteHeaders,
	}
	//若请求路径包含restful参数
	if common.IsRestfulPath(requestPath) {
		rewritePlugin.PathType = "regex" //正则替换

		//如果转发路径包含restful参数
		if common.IsRestfulPath(proxyPath) {
			proxyPath = formatProxyPath(requestPathLabel, proxyPath)
		}
		rewritePlugin.RegexPath = []*v1.RegexPath{
			{
				RegexPathMatch:   fmt.Sprintf("^%s$", common.ReplaceRestfulPath(requestPath, apintoRestfulRegexp)),
				RegexPathReplace: proxyPath,
			},
		}
		requestPath = fmt.Sprintf("~=^%s$", common.ReplaceRestfulPath(requestPath, apintoRestfulRegexp))
	} else {
		rewritePlugin.PathType = "prefix" //前缀替换
		rewritePlugin.PrefixPath = []*v1.PrefixPath{
			{
				PrefixPathMatch:   strings.TrimSuffix(requestPath, "*"),
				PrefixPathReplace: proxyPath,
			},
		}
	}

	rules := make([]*v1.RouterRule, 0, len(match))
	for _, m := range match {
		rule := &v1.RouterRule{
			Type:  m.Position,
			Name:  m.Key,
			Value: "",
		}

		if m.Position == enum.MatchPositionHeader {
			rule.Name = textproto.CanonicalMIMEHeaderKey(rule.Name)
		}

		switch m.MatchType {
		case enum.MatchTypeEqual:
			rule.Value = m.Pattern
		case enum.MatchTypePrefix:
			rule.Value = fmt.Sprintf("%s*", m.Pattern)
		case enum.MatchTypeSuffix:
			rule.Value = fmt.Sprintf("*%s", m.Pattern)
		case enum.MatchTypeSubstr:
			rule.Value = fmt.Sprintf("*%s*", m.Pattern)
		case enum.MatchTypeUneuqal:
			rule.Value = fmt.Sprintf("!=%s", m.Pattern)
		case enum.MatchTypeNull:
			rule.Value = "$"
		case enum.MatchTypeExist:
			rule.Value = "**"
		case enum.MatchTypeUnexist:
			rule.Value = "!"
		case enum.MatchTypeRegexp:
			rule.Value = fmt.Sprintf("~=%s", m.Pattern)
		case enum.MatchTypeRegexpG:
			rule.Value = fmt.Sprintf("~*=%s", m.Pattern)
		case enum.MatchTypeAny:
			rule.Value = "*"
		}

		rules = append(rules, rule)
	}

	templateID := ""
	if templateUUID != "" {
		templateID = fmt.Sprintf("%s@template", templateUUID)
	}

	if hosts == nil {
		hosts = []string{}
	}

	appendData := make(map[string]interface{}, 12)
	appendData["disable"] = disable
	appendData["listen"] = 0
	appendData["method"] = method
	appendData["host"] = hosts
	appendData["location"] = requestPath
	appendData["rules"] = rules
	appendData["service"] = serviceName
	appendData["template"] = templateID
	appendData["retry"] = retry
	appendData["time_out"] = timeout
	appendData["plugins"] = map[string]*v1.Plugin{
		"proxy_rewrite": { //插件名写死
			Disable: false,
			Config:  rewritePlugin,
		},
	}

	return &v1.RouterConfig{
		Name:        name,
		Description: desc,
		Append:      appendData,
	}
}
