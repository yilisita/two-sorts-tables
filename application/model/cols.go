/*
 * @Author: Wen Jiajun
 * @Date: 2022-06-30 15:35:37
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-06-30 16:11:20
 * @FilePath: \application\model\cols.go
 * @Description:
 */
package model

var ProdCols = []string{
	"装机容量（千瓦）",
	"本月发电量（万千瓦时）",
	"本季累计发电量（万千瓦时）",
	"本月止累计发电量（万千瓦时）",
	"本月上网电量（万千瓦时）",
	"本季累计上网电量（万千瓦时）",
	"本月止累计上网电量（万千瓦时）",
	"本月综合厂用电量（万千瓦时）",
	"本季累计综合厂用电量（万千瓦时）",
	"本月止累计综合厂用电量（万千瓦时）",
	"本月自发自用电量（万千瓦时）",
	"本季累计自发自用电量（万千瓦时）",
	"本月止累计自发自用电量（万千瓦时）",
	"本月其他电量",
	"本月止累计其他电量",
	"本月购网电量",
	"本月止累计购网电量",
	"电厂个数",
}

// 用户个数	用户用电装接容量	用电量(万千瓦时)
// 		本月	上年同月	累计	上年累计

var EcwsCols = []string{
	"用户个数",
	"用户用电装接容量",
	"本月用电量（万千瓦时）",
	"上年同月用电量（万千瓦时）",
	"上年累计用电量（万千瓦时）",
}

var EcwsLabels = []string{
	"全社会用电总计",
	"全行业用电合计",
	"第一产业",
	"第二产业",
	"第三产业",
	"城乡居民生活用电合计",
	"城镇居民",
	"乡村居民",
	"全行业用电分类",
	"农、林、牧、渔业",
	"农业",
	"林业",
	"畜牧业",
	"渔业",
	"农、林、牧、渔专业及辅助性活动",
	"排灌",
	"工业",
	"采矿业",
	"煤炭开采和洗选业",
	"石油和天然气开采业",
	"黑色金属矿采选业",
	"有色金属矿采选业",
	"非金属矿采选业",
	"其他采矿活动",
	"制造业",
	"农副食品加工业",
	"食品制造业",
	"酒、饮料及精制茶制造业",
	"烟草制品业",
	"纺织业",
	"纺织服装、服饰业",
	"皮革、毛皮、羽毛及其制品和制鞋业",
	"木材加工和木、竹、藤、棕、草制品业",
	"家具制造业",
	"造纸和纸制品业",
	"印刷和记录媒介复制业",
	"文教、工美、体育和娱乐用品制造业",
	"体育用品制造",
	"石油、煤炭及其他燃料加工业",
	"煤化工",
	"化学原料和化学制品制造业",
	"氯碱",
	"电石",
	"黄磷",
	"肥料制造",
	"医药制造业",
	"中成药生产",
	"生物药品制品制造",
	"化学纤维制造业",
	"橡胶和塑料制品业",
	"橡胶制品业",
	"塑料制品业",
	"非金属矿物制品业",
	"水泥制造",
	"玻璃制造",
	"陶瓷制品制造",
	"碳化硅",
	"黑色金属冶炼和压延加工业",
	"钢铁",
	"铁合金冶炼",
	"有色金属冶炼和压延加工业",
	"铝冶炼",
	"铅锌冶炼",
	"稀有稀土金属冶炼",
	"金属制品业",
	"结构性金属制品制造",
	"通用设备制造业",
	"风能原动设备制造",
	"专用设备制造业",
	"医疗仪器设备及器械制造",
	"汽车制造业",
	"新能源车整车制造",
	"铁路.船舶.航空航天和其他运输设备制造业",
	"铁路运输设备制造",
	"城市轨道交通设备制造",
	"航空、航天器及设备制造",
	"电气机械和器材制造业",
	"光伏设备及元器件制造",
	"计算机、通信和其他电子设备制造业",
	"计算机制造",
	"通信设备制造",
	"仪器仪表制造业",
	"其他制造业",
	"废弃资源综合利用业",
	"金属制品、机械和设备修理业",
	"电力、热力、燃气及水生产和供应业",
	"电力、热力生产和供应业",
	"电厂生产全部耗用电量",
	"线路损失电量",
	"抽水蓄能抽水耗用电量",
	"燃气生产和供应业",
	"水的生产和供应业",
	"建筑业",
	"房屋建筑业",
	"土木工程建筑业",
	"建筑安装业",
	"建筑装饰、装修和其他建筑业",
	"交通运输、仓储和邮政业",
	"铁路运输业",
	"电气化铁路",
	"道路运输业",
	"城市公共交通运输",
	"水上运输业",
	"港口岸电",
	"航空运输业",
	"管道运输业",
	"多式联运和运输代理业",
	"装卸搬运和仓储业",
	"邮政业",
	"信息传输、软件和信息技术服务业",
	"电信、广播电视和卫星传输服务",
	"互联网和相关服务",
	"互联网数据服务",
	"软件和信息技术服务业",
	"批发和零售业",
	"充换电服务业",
	"住宿和餐饮业",
	"金融业",
	"房地产业",
	"租赁和商务服务业",
	"其中：租赁业",
	"公共服务及管理组织",
	"科学研究和技术服务业",
	"地质勘查",
	"科技推广和应用服务业",
	"水利、环境和公共设施管理业",
	"水利管理业",
	"公共照明",
	"居民服务、修理和其他服务业",
	"教育、文化、体育和娱乐业",
	"其中：教育",
	"卫生和社会工作",
	"公共管理和社会组织、国际组织",
	"补充指标",
	"开采专业及辅助性活动",
}