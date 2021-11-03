package lib

/*
错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三位表示服务模块代码
第四五位表示具体错误代码
*/

var (

    OK 				= ErrCode{0, "OK"}

    // 系统错误, 前缀为 100
    Err500 			= ErrCode{10001, "系统服务奔溃了"}
    ErrParam        = ErrCode{10002, "请求参数错误"}
	Err404			= ErrCode{10003, "404资源不存在"}
	ErrCfg			= ErrCode{10004, "系统配置出错"}

    // 数据库错误, 前缀为 201
    ErrDb       	= ErrCode{20100, "数据库错误"}
    ErrFill     	= ErrCode{20101, "从数据库填充 struct 时发生错误"}

    // 认证错误, 前缀是 202
	Err401			= ErrCode{20200, "401认证失败"}
	Err403			= ErrCode{20201, "403资源不可用"}
    ErrValidation   = ErrCode{20202, "验证失败"}
    ErrTokenInvalid = ErrCode{20203, "令牌无效"}
    ErrApi          = ErrCode{20204, "Api不可用"}

    // 用户错误, 前缀为 203
    ErrUser		    = ErrCode{20301, "用户没找到"}
    ErrPsw			= ErrCode{20302, "密码错误"}
    ErrUa           = ErrCode{20303, "错误码：20303"}

    
)