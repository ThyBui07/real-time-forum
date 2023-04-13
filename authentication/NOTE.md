# Cookies
type Cookie struct {
    Name  string
    Value string
    Path       string    - indicates a URL path that must exist in req URL to send Cookie Header
    Domain     string    - specifies which hosts are allowed to recieve the cookie
    Expires    time.Time - deletes at specified date
    RawExpires string    // for reading cookies only
    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int		- deletes cookie after specified amount of time in seconds
    Secure   bool		- sent to server on encrypted req, never unsecured (HTTP)
    HttpOnly bool		- not accessible by JavaScript, only sent to server
    SameSite SameSite 	- servers require that a cookie shouldn't be sent with cross-orign req
    Raw      string
    Unparsed []string   - raw text of unparsed attribute-value pairs
}
