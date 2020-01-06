package models

type RobotGirl struct {
}

type ReturnMessage struct {
	Display string
	Code string
	NextNodes []string
}

type WelcomeMessage struct {
	Display string
	Code string
}

var welcomeMessages []WelcomeMessage = []WelcomeMessage{
	WelcomeMessage{ Display: "查询快递", Code: "查询快递" },
	WelcomeMessage{ Display: "申请退货", Code: "申请退货" },
	WelcomeMessage{ Display: "新品上架", Code: "新品上架" },
	WelcomeMessage{ Display: "商品推荐", Code: "商品推荐" },
	WelcomeMessage{ Display: "如何获取优惠券", Code: "如何获取优惠券" },
	WelcomeMessage{ Display: "商品何时可以送达", Code: "商品何时可以送达" },
	WelcomeMessage{ Display: "商品何时可以发货", Code: "商品何时可以发货" },
}

var messageMap map[string]ReturnMessage = map[string]ReturnMessage{
	"查询快递": ReturnMessage{ Display: "请在链接<a>https://syliaftershipi1.aftership.com/</a>中输入您的快递单号。", Code: "查询快递完成" },
	"申请退货": ReturnMessage{ Display: "非常抱歉我们的商品没能让您满意。以下是我们退货政策: <b>1. 请在签收之日起7日内提交退货申请。2. 退货商品的外观良好，不影响二次销售。3. 请在提交退货申请14日内告诉我们您的退货单号。<b>请您的商品是否符合我们的退货政策？", Code: "确认申请退货" },
}

func (r *RobotGirl) ReplyOption() []WelcomeMessage {
	return welcomeMessages
}
