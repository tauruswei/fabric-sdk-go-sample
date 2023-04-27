package result

import "fmt"

/**
 * @Author: WeiBingtao/13156050650@163.com
 * @Version: 1.0
 * @Description:
 * @Date: 2021/9/8 下午11:02
 */
type CodeMsg struct {
	Code int
	Msg  string
}

func (codemsg CodeMsg) FillArgs(args ...string) CodeMsg {
	codemsg.Msg = fmt.Sprintf(codemsg.Msg, args)
	return codemsg
}

var (
	SUCCESS               = CodeMsg{Code: 200, Msg: "SUCCESS"}
	SERVER_ERROR          = CodeMsg{Code: 500100, Msg: "服务端异常: %s"}
	PARAMETER_VALID_ERROR = CodeMsg{Code: 500101, Msg: "参数校验异常: %s"}
	TOKEN_VALID_ERROR     = CodeMsg{Code: 500102, Msg: "token 验证失败: %s"}

	// 用户模块
	USER_EXIST_ERROR           = CodeMsg{Code: 500201, Msg: "用户已存在: %s"}
	USER_UPDATE_ERROR          = CodeMsg{Code: 500202, Msg: "用户更新失败: %s"}
	GEN_MNEMONICS_ERROR        = CodeMsg{Code: 500203, Msg: "生成助记词失败: %s"}
	GEN_RANDOMCODES_ERROR      = CodeMsg{Code: 500204, Msg: "生成验证码失败: %s"}
	LOGIN_ERROR                = CodeMsg{Code: 500205, Msg: "登录失败: %s"}
	LOGOUT_ERROR               = CodeMsg{Code: 500206, Msg: "退出登录失败: %s"}
	MODIFY_MOBILENUMBER_ERROR  = CodeMsg{Code: 500207, Msg: "修改手机号失败: %s"}
	MODIFY_PSSSWD_ERROR        = CodeMsg{Code: 500208, Msg: "修改密码失败: %s"}
	MODIFY_PERSONAL_DATA_ERROR = CodeMsg{Code: 500209, Msg: "修改个人资料失败: %s"}
	REGISTER_USER_ERROR        = CodeMsg{Code: 500210, Msg: "注册用户失败: %s"}
	SEARCH_USER_ERROR          = CodeMsg{Code: 500211, Msg: "查询用户失败: %s"}
	UNLOCK_USER_ERROR          = CodeMsg{Code: 500212, Msg: "解锁用户失败: %s"}
	USER_VERIFY_ERROR          = CodeMsg{Code: 500213, Msg: "用户不存在: %s"}

	// 文创作品交易平台
	FILE_UPLOAD_ERROR   = CodeMsg{Code: 500301, Msg: "上传文件失败: %s"}
	FILE_DOWNLOAD_ERROR = CodeMsg{Code: 500302, Msg: "下载文件失败: %s"}
	CREATE_NFT_ERROR    = CodeMsg{Code: 500303, Msg: "创建 NFT 失败: %s"}
	BUY_NFT_ERROR       = CodeMsg{Code: 500304, Msg: "购买 NFT 失败: %s"}
	GET_OSS_TOKEN_ERROR = CodeMsg{Code: 500305, Msg: "获取 oss token 失败: %s"}
	SIGN_CERT_ERROR     = CodeMsg{Code: 500306, Msg: "签发证书失败: %s"}
	RELEASE_SAVE_ERROR  = CodeMsg{Code: 500307, Msg: "release存库失败: %s"}

	//nft
	CONTRACT_LIST_ERROR                   = CodeMsg{Code: 500401, Msg: "获取合同列表失败：%s"}
	CREATE_MUSIC_CONTRACT_ERROR           = CodeMsg{Code: 500402, Msg: "创建音乐合同失败：%s"}
	CREATE_MUSIC_DELEGATED_CONTRACT_ERROR = CodeMsg{Code: 500403, Msg: "创建音乐授权合同失败：%s"}
	CREATE_MUSIC_DELEGATED_NFT_ERROR      = CodeMsg{Code: 500404, Msg: "创建音乐授权 NFT token 失败：%s"}
	CREATE_MUSIC_NFT_ERROR                = CodeMsg{Code: 500405, Msg: "创建音乐 NFT 失败：%s"}
	CREATE_NFT721_ERROR                   = CodeMsg{Code: 500406, Msg: "创建 NFT 721 失败：%s"}
	MODIFY_NFT721_ERROR                   = CodeMsg{Code: 500407, Msg: "修改 NFT 721 失败：%s"}
	NFT_LIST_ERROR                        = CodeMsg{Code: 500408, Msg: "获取 NFT 列表失败：%s"}
	QUERY_CONTRACT_DETAIL_ERROR           = CodeMsg{Code: 500409, Msg: "查询合同详情失败：%s"}
	QUERY_MUSIC_DELEGATED_TOKEN_ERROR     = CodeMsg{Code: 5004010, Msg: "查询音乐授权 NFT token 失败：%s"}
	QUERY_MUSIC_NFT_TOKEN_ERROR           = CodeMsg{Code: 5004011, Msg: "查询音乐 NFT token 失败：%s"}
	SELL_NFT721_ERROR                     = CodeMsg{Code: 5004012, Msg: "卖出 NFT 721 失败：%s"}

	// 订单
	CREATE_ORDER_ERROR = CodeMsg{Code: 500501, Msg: "创建订单失败：%s"}
	QUERY_ORDER_ERROR  = CodeMsg{Code: 500502, Msg: "查询订单失败：%s"}
	CANCEL_ORDER_ERROR = CodeMsg{Code: 500503, Msg: "取消订单失败：%s"}
	// 购物车
	CREATE_CART_ERROR = CodeMsg{Code: 500601, Msg: "创建购物车失败：%s"}
	QUERY_CART_ERROR  = CodeMsg{Code: 500602, Msg: "查询购物车失败：%s"}
	DELETE_CART_ERROR = CodeMsg{Code: 500603, Msg: "删除商品失败：%s"}
	//消息

	//
)
