package model

type Tcpinfo struct {
	Id        int64  `xorm:"pk autoincr notnull" json:"id"`
	Uuid      string `xorm:"varchar(255)" json:"uuid"`
	Th_dport  int    `xorm:"int(5)" json:"th_dport"`
	Th_sport  int    `xorm:"int(5)" json:"th_sport"`
	Seq_num   int32  `xorm:"int(32)" json:"seq_num"`
	Ack_num   int32  `xorm:"int(32)" json:"ack_num"`
	Timestamp int64  `xorm:"bigint" json:"th_sport"`
	Ip_dst    string `xorm:"varchar(255)" json:"ip_dst"`
	Ip_src    string `xorm:"varchar(255)" json:"ip_src"`
	Body      string `xorm:"text" json:"body"`
}
