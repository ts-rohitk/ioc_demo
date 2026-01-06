package mapping

type IOCType string

const (
	IOCIPv4   IOCType = "ipv4"
	IOCIPv6   IOCType = "ipv6"
	IOCIPPort IOCType = "ip:port"
	IOCDomain IOCType = "domain"
	IOCURL    IOCType = "url"
	IOCMD5    IOCType = "md5"
	IOCSHA1   IOCType = "sha1"
	IOCSHA256 IOCType = "sha256"
)

type IOC struct {
	// ID           *primitive.ObjectID     `bson:"_id,omitempty" json:"id,omitempty"`
	UUID         string                  `bson:"uuid" json:"uuid"`
	HashCode     string                  `bson:"hashCode" json:"hashCode"`
	Type         IOCType                 `bson:"type" json:"type"`
	ValueRaw     string                  `bson:"rawValue" json:"rawValue"`
	ValueNorm    string                  `bson:"normalizedValue" json:"normalizedValue"`
	Title        *string                 `bson:"title" json:"title"`
	Key          string                  `bson:"key" json:"key"`
	FirstSeen    *int64                  `bson:"firstSeen" json:"firstSeen"`
	LastSeen     *int64                  `bson:"lastSeen"  json:"lastSeen"`
	CreatedAt    int64                   `bson:"createdAt" json:"createdAt"`
	UpdatedAt    *int64                  `bson:"updatedAt" json:"updatedAt"`
	ExpiresAt    *int64                  `bson:"expiresAt" json:"expiresAt"`
	ThreatType   *string                 `bson:"threatType" json:"threatType"`
	Tags         []string                `bson:"tags" json:"tags"`
	Confidence   *string                 `bson:"confidence" json:"confidence"`
	Malware      []*MalwareInfo          `bson:"malware" json:"malware"`
	Network      *Network                `bson:"network" json:"network"`
	ThreatActors []*ThreatActors         `bson:"threatActors" json:"threatActors"`
	Victims      []*Victim               `bson:"victims" json:"victims"`
	Sources      []*Source               `bson:"sources" json:"sources"`
	TTP          *[]string               `bson:"ttp" json:"ttp"`
	Meta         *map[string]interface{} `bson:"meta,omitempty" json:"meta,omitempty"`
}

type ThreatActors struct {
	UUID       string  `bson:"uuid" json:"uuid"`
	Name       string  `bson:"name" json:"name"`
	Source     *string `bson:"source" json:"source"`
	Confidence *string `bson:"confidence" json:"confidence"`
}

type Victim struct {
	Type  string   `bson:"type" json:"type"`
	Value []string `bson:"values" json:"values"`
}

type MalwareInfo struct {
	UUID         string   `bson:"uuid" json:"uuid"`
	Family       *string  `bson:"family" json:"family"`
	Aliases      []string `bson:"aliases" json:"aliases"`
	DisplayName  *string  `bson:"displayName"  json:"displayName"`
	PlatformHint *string  `bson:"platformHint" json:"platformHint"`
}

type Network struct {
	IP      *string `bson:"ip"      json:"ip"`
	Port    *int64  `bson:"port"    json:"port"`
	ASN     *string `bson:"asn"     json:"asn"`
	Country *string `bson:"country" json:"country"`
}

type Source struct {
	Name        string  `bson:"name" json:"name"`
	URL         *string `bson:"url"  json:"url"`
	CollectedAt int64   `bson:"collectedAt" json:"collectedAt"`
	Confidence  *string `bson:"confidence" json:"confidence"`
}
