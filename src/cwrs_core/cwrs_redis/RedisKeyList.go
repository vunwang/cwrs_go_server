package cwrs_redis

import "time"

/**
说明:
KEY开头: Redis Key的标识
PARAM开头: 某个标识的具体参数名称
*/

const KEY_SYS_USER_TOKEN = "sys_user_token:"   // 用户token
const KEY_SYS_USER_TOKEN_TIME = 24 * time.Hour // 用户token过期时间

const KEY_SYS_PARAM = "sys_param:" // 系统参数 如：平台logo、平台名称、大屏标题
