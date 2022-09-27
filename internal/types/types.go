// Code generated by goctl. DO NOT EDIT.
package types

type BKCallBack struct {
	Type                string                `json:"type"`
	Scenario            string                `json:"scenario"`
	BkBizId             int64                 `json:"bk_biz_id"`
	BkBizName           string                `json:"bk_biz_name"`
	Event               BkEvent               `json:"event"`
	Strategy            BkStrategy            `json:"strategy"`
	LatestAnomalyRecord BkLatestAnomalyRecord `json:"latest_anomaly_record"`
	RelatedInfo         string                `json:"related_info"`
	Labels              []interface{}         `json:"labels,optional"`
	BotKey              string                `form:"key"`
	Secret              string                `form:"secret"`
	TimeLocal           int64                 `form:"timeLocal,default=1"`
}

type BkEvent struct {
	Id                   int64                  `json:"id"`
	EventId              string                 `json:"event_id"`
	BeginTime            string                 `json:"begin_time"`
	CreateTime           string                 `json:"create_time"`
	EndTime              string                 `json:"end_time"`
	Level                int64                  `json:"level"`
	LevelName            string                 `json:"level_name"`
	Dimensions           BkDimensions           `json:"dimensions,optional"`
	DimensionTranslation BkDimensionTranslation `json:"dimension_translation,optional"`
}

type BkDimensions struct {
	BkTargetCloudId string   `json:"bk_target_cloud_id"`
	BkTargetIp      string   `json:"bk_target_ip"`
	BkTopoNode      []string `json:"bk_topo_node"`
}

type BkDimensionTranslation struct {
	BkTargetCloudId BkTargetValue  `json:"bk_target_cloud_id"`
	BkTargetIp      BkTargetValue  `json:"bk_target_ip"`
	BktopoNode      BkTargetValues `json:"bk_topo_node"`
}

type BkTargetValue struct {
	Value        string `json:"value"`
	DisplayName  string `json:"display_name"`
	DisplayValue string `json:"display_value"`
}

type BkTargetValues struct {
	Value        []string          `json:"value"`
	DisplayName  string            `json:"display_name"`
	DisplayValue []BkDisplayValues `json:"display_value"`
}

type BkDisplayValues struct {
	ObjName  string `json:"bk_obj_name"`
	InstName string `json:"bk_inst_name"`
}

type BkStrategy struct {
	Id       int64        `json:"id"`
	Name     string       `json:"name"`
	Scenario string       `json:"scenario"`
	ItemList []BkItemList `json:"item_list"`
}

type BkItemList struct {
	MetricField     string `json:"metric_field"`
	MetricFieldName string `json:"metric_field_name"`
	DataSourceLabel string `json:"data_source_label"`
	DataSourceName  string `json:"data_source_name"`
	DataTypeLabel   string `json:"data_type_label"`
	DataTypeName    string `json:"data_type_name"`
	MetricId        string `json:"metric_id"`
}

type BkLatestAnomalyRecord struct {
	AnomalyId   int64         `json:"anomaly_id"`
	SourceTime  string        `json:"source_time"`
	CreateTime  string        `json:"create_time"`
	OriginAlarm BkOriginAlarm `json:"origin_alarm"`
}

type BkOriginAlarm struct {
	Data    BkData    `json:"data"`
	RecorId string    `json:"record_id,optional"`
	Anomaly BkAnomaly `json:"anomaly,optional"`
}

type BkData struct {
	Time       int64        `json:"time"`
	Value      float64      `json:"value"`
	Values     Bkvalues     `json:"values"`
	Dimensions BkDimensions `json:"dimensions,optional"`
}

type Bkvalues struct {
	A      float64 `json:"a,optional"`
	Result float64 `json:"_result_,optional"`
	Time   int64   `json:"time"`
}

type BkAnomaly struct {
	Field1 BkField1 `json:"1,optional"`
	Field2 BkField2 `json:"2,optional"`
	Field3 BkField3 `json:"3,optional"`
}

type BkField1 struct {
	AnomalyMessage string `json:"anomaly_message"`
	AnomalyTime    string `json:"anomaly_time"`
	AnomalyId      string `json:"anomaly_id"`
}

type BkField2 struct {
	AnomalyMessage string `json:"anomaly_message"`
	AnomalyTime    string `json:"anomaly_time"`
	AnomalyId      string `json:"anomaly_id"`
}

type BkField3 struct {
	AnomalyMessage string `json:"anomaly_message"`
	AnomalyTime    string `json:"anomaly_time"`
	AnomalyId      string `json:"anomaly_id"`
}