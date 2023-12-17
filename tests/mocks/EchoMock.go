package mocks

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

// CustomContext is a mock implementation of echo.Context
type CustomContext struct {
	mock.Mock
}

func (c *CustomContext) Request() *http.Request {
	args := c.Called()
	return args.Get(0).(*http.Request)
}

func (c *CustomContext) Response() *echo.Response {
	args := c.Called()
	return args.Get(0).(*echo.Response)
}

func (c *CustomContext) IsTLS() bool {
	args := c.Called()
	return args.Bool(0)
}

func (c *CustomContext) Scheme() string {
	args := c.Called()
	return args.String(0)
}

func (c *CustomContext) RealIP() string {
	args := c.Called()
	return args.String(0)
}

func (c *CustomContext) Path() string {
	args := c.Called()
	return args.String(0)
}

func (c *CustomContext) Param(name string) string {
	args := c.Called(name)
	return args.String(0)
}

func (c *CustomContext) ParamNames() []string {
	args := c.Called()
	return args.Get(0).([]string)
}

func (c *CustomContext) QueryParam(name string) string {
	args := c.Called(name)
	return args.String(0)
}

func (c *CustomContext) QueryParams() url.Values {
	args := c.Called()
	return args.Get(0).(url.Values)
}

func (c *CustomContext) QueryString() string {
	args := c.Called()
	return args.String(0)
}

func (c *CustomContext) FormValue(name string) string {
	args := c.Called(name)
	return args.String(0)
}

func (c *CustomContext) FormParams() (url.Values, error) {
	args := c.Called()
	return args.Get(0).(url.Values), args.Error(1)
}

func (c *CustomContext) FormFile(name string) (*multipart.FileHeader, error) {
	args := c.Called(name)
	return args.Get(0).(*multipart.FileHeader), args.Error(1)
}

func (c *CustomContext) MultipartForm() (*multipart.Form, error) {
	args := c.Called()
	return args.Get(0).(*multipart.Form), args.Error(1)
}

func (c *CustomContext) Cookie(name string) (*http.Cookie, error) {
	args := c.Called(name)
	return args.Get(0).(*http.Cookie), args.Error(1)
}

func (c *CustomContext) SetCookie(cookie *http.Cookie) {
	c.Called(cookie)
}

func (c *CustomContext) Cookies() []*http.Cookie {
	args := c.Called()
	return args.Get(0).([]*http.Cookie)
}

func (c *CustomContext) Bind(i interface{}) error {
	args := c.Called(i)
	return args.Error(0)
}

func (c *CustomContext) BindJSON(i interface{}) error {
	args := c.Called(i)
	return args.Error(0)
}

func (c *CustomContext) BindXML(i interface{}) error {
	args := c.Called(i)
	return args.Error(0)
}

func (c *CustomContext) BindQuery(i interface{}) error {
	args := c.Called(i)
	return args.Error(0)
}

func (c *CustomContext) BindParam(i interface{}) error {
	args := c.Called(i)
	return args.Error(0)
}

func (c *CustomContext) Validate(i interface{}) error {
	args := c.Called(i)
	return args.Error(0)
}

func (c *CustomContext) Render(code int, name string, data interface{}) error {
	args := c.Called(code, name, data)
	return args.Error(0)
}

func (c *CustomContext) HTML(code int, html string) error {
	args := c.Called(code, html)
	return args.Error(0)
}

func (c *CustomContext) HTMLBlob(code int, b []byte) error {
	args := c.Called(code, b)
	return args.Error(0)
}

func (c *CustomContext) String(code int, format string, values ...interface{}) error {
	args := c.Called(code, format, values)
	return args.Error(0)
}

func (c *CustomContext) JSON(code int, i interface{}) error {
	args := c.Called(code, i)
	return args.Error(0)
}

func (c *CustomContext) JSONPretty(code int, i interface{}, indent string) error {
	args := c.Called(code, i, indent)
	return args.Error(0)
}

func (c *CustomContext) JSONBlob(code int, b []byte) error {
	args := c.Called(code, b)
	return args.Error(0)
}

func (c *CustomContext) JSONP(code int, callback string, i interface{}) error {
	args := c.Called(code, callback, i)
	return args.Error(0)
}

func (c *CustomContext) JSONPBlob(code int, callback string, b []byte) error {
	args := c.Called(code, callback, b)
	return args.Error(0)
}

func (c *CustomContext) XML(code int, i interface{}) error {
	args := c.Called(code, i)
	return args.Error(0)
}

func (c *CustomContext) XMLPretty(code int, i interface{}, indent string) error {
	args := c.Called(code, i, indent)
	return args.Error(0)
}

func (c *CustomContext) XMLBlob(code int, b []byte) error {
	args := c.Called(code, b)
	return args.Error(0)
}

func (c *CustomContext) Blob(code int, contentType string, b []byte) error {
	args := c.Called(code, contentType, b)
	return args.Error(0)
}

func (c *CustomContext) Stream(code int, contentType string, r io.Reader) error {
	args := c.Called(code, contentType, r)
	return args.Error(0)
}

func (c *CustomContext) File(file string) error {
	args := c.Called(file)
	return args.Error(0)
}

func (c *CustomContext) Attachment(file string, name string) error {
	args := c.Called(file, name)
	return args.Error(0)
}

func (c *CustomContext) Inline(file string, name string) error {
	args := c.Called(file, name)
	return args.Error(0)
}

func (c *CustomContext) NoContent(code int) error {
	args := c.Called(code)
	return args.Error(0)
}

func (c *CustomContext) Redirect(code int, url string) error {
	args := c.Called(code, url)
	return args.Error(0)
}

func (c *CustomContext) Error(err error) {
	c.Called(err)
}

func (c *CustomContext) Handler() echo.HandlerFunc {
	args := c.Called()
	return args.Get(0).(echo.HandlerFunc)
}

func (c *CustomContext) SetHandler(h echo.HandlerFunc) {
	c.Called(h)
}

func (c *CustomContext) Logger() echo.Logger {
	args := c.Called()
	return args.Get(0).(echo.Logger)
}

func (c *CustomContext) Echo() *echo.Echo {
	args := c.Called()
	return args.Get(0).(*echo.Echo)
}

func (c *CustomContext) Reset(r *http.Request, w http.ResponseWriter) {
	c.Called(r, w)
}
