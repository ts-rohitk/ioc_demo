package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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

/*

   {
       "id": "1689288", x reference == null ? https://threatfox.abuse.ch/ioc/{id} : reference_url
       "ioc": "vlxx.bz", ioc
       "threat_type": "botnet_cc", command & control
       "threat_type_desc": "Indicator that identifies a botnet command&control server (C&C)",
       "ioc_type": "domain", Type (required)
       "ioc_type_desc": "Domain that is used for botnet Command&control (C&C)",
       "malware": "win.quasar_rat", malware_display_name
       "malware_printable": "Quasar RAT", malware_family
       "malware_alias": "CinaRAT,QuasarRAT,Yggdrasil",
       "malware_malpedia": "https://malpedia.caad.fkie.fraunhofer.de/details/win.quasar_rat",
       "confidence_level": 100, 0-25 limited , 25-50 : moderate , 50-75: elevated , 75-100: high
       "first_seen": "2026-01-01 04:07:49 UTC", -> in64 milliseconds
       "last_seen": null, { null ? first_seen : last_seen }
       "reference": "https://www.virustotal.com/gui/file/ecfc9a5ee4b9ea2c241ce89f30b276d832270fdc826af3c683ac29d832db5231/detection",
       "reporter": "Lucas_x",
       "tags": [
           "quasar",
           "RAT"
       ]
   },

   -> malware_family is null , if  not f"{threat_type} (server/infrastructure/)used by(verb) {malware_family} "
	| Domain "vlxx.bz" used for botnet command & control by {malware_name}
	-> union tags
   	-> earliest first_seen
   	-> latest last seen
	- if new malware family , save !
	- UpdatedAt if updating

	[x] if incoming ioc has a malware & existing ioc has no specific malware or a different malware family , lookup a malware
*/

type IOC struct {
	ID           *primitive.ObjectID     `bson:"_id,omitempty" json:"id,omitempty"`
	UUID         string                  `bson:"uuid" json:"uuid"`
	Type         IOCType                 `bson:"type" json:"type"`
	ValueRaw     string                  `bson:"rawValue" json:"rawValue"`               // ioc raw data
	ValueNorm    string                  `bson:"normalizedValue" json:"normalizedValue"` // small letters , cleanup
	Title        *string                 `bson:"title" json:"title"`                     // `${ioc} + ${threat_type} + ${type} + ${malware_family_name}`
	Key          string                  `bson:"key" json:"key"`                         // required f"{Type}|{valueNorm}"
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
	Name        string  `bson:"name" json:"name"`               // source_name (ThreatFox)
	URL         *string `bson:"url"  json:"url"`                // reference_url
	CollectedAt int64   `bson:"collectedAt" json:"collectedAt"` // current timestamp in milliseconds
	Confidence  *string `bson:"confidence" json:"confidence"`   // same as confidence from incomming data
}
