package HAProxyGoStat

import "time"

// HAProxyColumnNames is a map of the names HAProxy gives to the fields in the HAProxyStat struct
var HAProxyColumnNames = map[string]string{
	"pxname":         "ProxyName",
	"svname":         "ServerName",
	"qcur":           "CurrentQueue",
	"qmax":           "MaxQueue",
	"scur":           "CurrentSessions",
	"smax":           "MaxSessions",
	"slim":           "SessionLimit",
	"stot":           "TotalSessions",
	"bin":            "BytesIn",
	"bout":           "BytesOut",
	"dreq":           "DeniedRequests",
	"dresp":          "DeniedResponses",
	"ereq":           "ErrorRequests",
	"econ":           "ErrorConnections",
	"eresp":          "ErrorResponses",
	"wretr":          "ServerRetry",
	"wredis":         "ServerRedispatch",
	"status":         "Status",
	"weight":         "Weight",
	"act":            "ActiveServers",
	"bck":            "BackupServers",
	"chkfail":        "FailedChecks",
	"chkdown":        "DownTransitions",
	"lastchg":        "LastTransition",
	"downtime":       "Downtime",
	"qlimit":         "QueueLimit",
	"pid":            "PID",
	"iid":            "IID",
	"sid":            "ServerID",
	"throttle":       "Throttle",
	"lbtot":          "LoadBalanceTotal",
	"tracked":        "Tracked",
	"type":           "Type",
	"rate":           "Rate",
	"rate_lim":       "RateLimit",
	"rate_max":       "RateMax",
	"check_status":   "CheckStatus",
	"check_code":     "CheckCode",
	"check_duration": "CheckDuration",
	"hrsp_1xx":       "HTTP1xx",
	"hrsp_2xx":       "HTTP2xx",
	"hrsp_3xx":       "HTTP3xx",
	"hrsp_4xx":       "HTTP4xx",
	"hrsp_5xx":       "HTTP5xx",
	"hrsp_other":     "HTTPOther",
	"hanafail":       "HANAFail",
	"req_rate":       "HTTPRequestRate",
	"req_rate_max":   "HTTPRequestRateMax",
	"req_tot":        "HTTPRequestsTotal",
	"cli_abrt":       "ClientAbort",
	"srv_abrt":       "ServerAbort",
	"comp_in":        "CompressorIn",
	"comp_out":       "CompressorOut",
	"comp_byp":       "CompressorBypass",
	"comp_resp":      "CompressorResponse",
	"lastsess":       "LastSession",
	"last_chk":       "LastCheck",
	"last_agt":       "LastAgent",
	"qtime":          "QueueTime",
	"ctime":          "ConnectTime",
	"rtime":          "ResponseTime",
	"ttime":          "TotalTime",
	"agent_status":   "AgentStatus",
	"agent_code":     "AgentCode",
	"agent_duration": "AgentDuration",
	"check_desc":     "CheckDescription",
	"agent_desc":     "AgentDescription",
	"check_rise":     "CheckRise",
	"check_fall":     "CheckFall",
	"check_health":   "CheckHealth",
	"agent_rise":     "AgentRise",
	"agent_fall":     "AgentFall",
	"agent_health":   "AgentHealth",
	"address":        "Address",
	"cookie":         "Cookie",
	"mode":           "Mode",
	"algo":           "Algorithm",
	"conn_rate":      "ConnectionRate",
	"conn_rate_max":  "ConnectionRateMax",
	"conn_tot":       "ConnectionTotal",
	"intercepted":    "Intercepted",
	"dcon":           "DeniedConnections",
	"dses":           "DeniedSessions",
}

type HAProxyStat struct {
	Timestamp		   time.Time
	ProxyName          string // proxy name
	ServerName         string // server name or FRONTEND/BACKEND
	CurrentQueue       int
	MaxQueue           int
	CurrentSessions    int
	MaxSessions        int
	SessionLimit       int
	TotalSessions      int
	BytesIn            int
	BytesOut           int
	DeniedRequests     int
	DeniedResponses    int
	ErrorRequests      int
	ErrorConnections   int
	ErrorResponses     int
	ServerRetry        int
	ServerRedispatch   int
	Status             string
	Weight             int
	ActiveServers      int
	BackupServers      int
	FailedChecks       int
	DownTransitions    int
	LastTransition     int    //number of seconds since last transition
	Downtime           int    // total downtime (in seconds)
	QueueLimit         int    // configured maxqueue for server
	PID                int    //process ID
	IID                int    //unique proxy id
	ServerID           int    //server ID (unique within proxy)
	Throttle           int    // current throttle percentage
	LoadBalanceTotal   int    // total times server was selected for session
	Tracked            int    // id of proxy/server if tracking is enabled
	Type               string // comes in as 0/1/2/3, Frontend/Backend/Server/Listener
	Rate               int    // number of sessions in last second
	RateLimit          int    // configured limit on new sessions per second
	RateMax            int    // max number of new sessions per second
	CheckStatus        string
	CheckCode          int
	CheckDuration      int // ms to perform last health check
	HTTP1xx            int
	HTTP2xx            int
	HTTP3xx            int
	HTTP4xx            int
	HTTP5xx            int
	HTTPOther          int
	HANAFail           string
	HTTPRequestRate    int // HTTP requests in the past second
	HTTPRequestRateMax int // maximum of HTTPRequestRate
	HTTPRequestsTotal  int
	ClientAbort        int // number of data transfers aborted by client
	ServerAbort        int // number of data transfers aborted by server
	CompressorIn       int // bytes fed to compressor
	CompressorOut      int // bytes out of compressor
	CompressorBypass   int // bytes that bypassed compressor
	CompressorResponse int // number of responses compressed
	LastSession        int // seconds since last assigned session
	LastCheck          string
	LastAgent          string
	QueueTime          int // average queue time in ms over last 1024 requests
	ConnectTime        int // average connect time in ms over last 1024 requests
	ResponseTime       int // average response time in ms over last 1024 requests
	TotalTime          int // average total time in ms over last 1024 requests
	AgentStatus        string
	AgentCode          int
	AgentDuration      int
	CheckDescription   string
	AgentDescription   string
	CheckRise          int
	CheckFall          int
	CheckHealth        int
	AgentRise          int
	AgentFall          int
	AgentHealth        int
	Address            string // address:port or "unix"
	Cookie             string // server's cookie value or backend's cookie name
	Mode               string // tcp, http, health, unknown
	Algorithm          string // load balancing algorithm
	ConnectionRate     int    // connections in the last second
	ConnectionRateMax  int
	ConnectionTotal    int
	Intercepted        int // cumulative number of intercepted requests
	DeniedConnections  int
	DeniedSessions     int
}
