package param

type RequestTcp struct {
	Uuid      string `json:"uuid"`
	Th_dport  int    `json:"th_dport"`
	Th_sport  int    `json:"th_sport"`
	Seq_num   int32  `json:"seq_num"`
	Ack_num   int32  `json:"ack_num"`
	Timestamp int64  `json:"th_sport"`
	Ip_dst    string `json:"ip_dst"`
	Ip_src    string `json:"ip_src"`
	Body      string `json:"body"`
}
