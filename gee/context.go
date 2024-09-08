package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// reponse info
	StatusCode int
}

// newContext 创建一个新的Context实例，用于封装HTTP请求和响应。
// 它初始化了Context中的Writer、Req、Path和Method字段。
func newContext(w http.ResponseWriter, req *http.Request) *Context {
    return &Context{
        Writer: w,
        Req:    req,
        Path:   req.URL.Path,
        Method: req.Method,
    }
}

// PostForm 从请求的表单中获取指定键的值。
// 它返回请求中指定键的表单值。
func (c *Context) PostForm(key string) string {
    return c.Req.FormValue(key)
}

// Query 从请求的URL查询参数中获取指定键的值。
// 它返回请求URL中指定键的查询参数值。
func (c *Context) Query(key string) string {
    return c.Req.URL.Query().Get(key)
}

// Status 设置HTTP响应状态码。
// 它更新了Context中的StatusCode字段，并向客户端写入状态码。
func (c *Context) Status(code int) {
    c.StatusCode = code
    c.Writer.WriteHeader(code)
}

// SetHeader 向HTTP响应中添加或更新一个头信息。
// 它设置了响应头中指定键的值。
func (c *Context) SetHeader(key string, value string) {
    c.Writer.Header().Set(key, value)
}

// String 向客户端返回纯文本响应。
// 它设置响应头的Content-Type为"text/plain"，设置HTTP响应状态码，并将格式化后的字符串写入响应。
func (c *Context) String(code int, format string, values ...interface{}) {
    c.SetHeader("Content-Type", "text/plain")
    c.Status(code)
    c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 向客户端返回JSON格式的响应。
// 它设置响应头的Content-Type为"application/json"，设置HTTP响应状态码，并将给定的对象编码为JSON写入响应。
func (c *Context) JSON(code int, obj interface{}) {
    c.SetHeader("Content-Type", "application/json")
    c.Status(code)
    encoder := json.NewEncoder(c.Writer)
    if err := encoder.Encode(obj); err != nil {
        http.Error(c.Writer, err.Error(), 500)
    }
}

// Data 向客户端返回原始数据响应。
// 它设置HTTP响应状态码，并将给定的字节数据写入响应。
func (c *Context) Data(code int, data []byte) {
    c.Status(code)
    c.Writer.Write(data)
}

// HTML 向客户端返回HTML格式的响应。
// 它设置响应头的Content-Type为"text/html"，设置HTTP响应状态码，并将给定的HTML字符串写入响应。
func (c *Context) HTML(code int, html string) {
    c.SetHeader("Content-Type", "text/html")
    c.Status(code)
    c.Writer.Write([]byte(html))
}
