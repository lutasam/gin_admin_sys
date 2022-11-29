package common

import "time"

const ISSUER = "LUTASAM"                                // jwt issuer
const PASSWORDSALT = "astaxie12798akljzmknm.ahkjkljl;k" // use only for password encryption
const OTHERSECRETSALT = "9871267812345mn812345xyz"      // user for other encryption
const EXPIRETIME = 86400000                             // jwt expiration time. 1 day's second
const ACTIVECODEEXPTIME = 300 * time.Second             // active code expiration time. 5 min
const ACTIVECODESUFFIX = "_active_code"
const DEFAULTAVATARURL = "http://baidu.com/test.png"
const DEFAULTSIGN = "This man doesn't say anything."
const DEFAULTNICKNAME = "SHAREUSER"
const MAXIMGSPACE = 1024 * 1024 * 1 // img upload should be less than 1 MB

const (
	STATUSOKCODE    = 200
	CLIENTERRORCODE = 400
	SERVERERRORCODE = 500
)

const (
	STATUSOKMSG    = "OK"
	CLIENTERRORMSG = "400 CLIENT ERROR"
	SERVERERRORMSG = "500 SERVER ERROR"
)
