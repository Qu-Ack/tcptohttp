package parser

type METHOD string
type HEADER string

const (
	OPTIONS METHOD = "OPTIONS"
	GET     METHOD = "GET"
	HEAD    METHOD = "HEAD"
	PUT     METHOD = "PUT"
	POST    METHOD = "POST"
	DELETE  METHOD = "DELETE"
	TRACE   METHOD = "TRACE"
	CONNECT METHOD = "CONNECT"

	ACCEPT_HEADER              HEADER = "Accept"
	ACCEPT_CHARSET_HEADER      HEADER = "Accept-Charset"
	ACCEPT_ENCODING_HEADER     HEADER = "Accept-Encoding"
	ACCEPT_LANGUAGE_HEADER     HEADER = "Accept-Language"
	AUTHORIAZATION_HEADER      HEADER = "Authorization"
	EXPECT_HEADER              HEADER = "Expect"
	FROM_HEADER                HEADER = "From"
	HOST_HEADER                HEADER = "Host"
	IF_MATCH_HEADER            HEADER = "If-Match"
	IF_MODIFIED_SINCE_HEADER   HEADER = "If-Modified-Since"
	IF_NONE_MATCH_HEADER       HEADER = "If-None-Match"
	IF_RANGE_HEADER            HEADER = "If-Range"
	IF_UNMODIFIED_SINCE_HEADER HEADER = "If-Unmodified-Since"
	MAX_FORWARDS_HEADER        HEADER = "Max-Forwards"
	PROXY_AUTHORIZATION_HEADER HEADER = "Proxy-Authorization"
	RANGE_HEADER               HEADER = "Range"
	REFERER_HEADER             HEADER = "Referer"
	TE_HEADER                  HEADER = "TE"
	USER_AGENT_HEADER          HEADER = "User-Agent"
	CONTENT_TYPE_HEADER        HEADER = "Content-Type"
)

type Request struct {
	Method          METHOD
	URI             string
	ProtocolVersion string

	Body []byte

	Headers map[HEADER]string
}
