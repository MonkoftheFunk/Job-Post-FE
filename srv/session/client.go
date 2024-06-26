package session

import (
	golaravelsession "Job-Post-FE/srv/session/laravel"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
	"math/big"
	"net/http"
	"net/url"
	"strings"
)

type Config struct {
	Config  redis.Options
	Key     string
	Cookie  string
	Prefix  string
	UserKey string
}

type Client struct {
	Redis  *redis.Client
	Config *Config
}

func NewClient(c *Config) *Client {
	return &Client{
		Redis:  redis.NewClient(&c.Config),
		Config: c,
	}
}

func (c *Client) GetLoggedIn(r *http.Request) (bool, error) {
	id, err := c.GetUserID(r)
	if err != nil {
		return false, err
	}

	return id.Cmp(big.NewInt(0)) == 0, nil
}

func (c *Client) GetUserID(r *http.Request) (*big.Int, error) {
	data, err := c.GetData(r)
	if err != nil {
		return big.NewInt(0), err
	}

	if data[c.Config.UserKey] == nil {
		return big.NewInt(0), nil
	}

	return big.NewInt(data[c.Config.UserKey].(int64)), nil
}

func (c *Client) GetData(r *http.Request) (php_serialize.PhpArray, error) {
	id, err := c.GetSessionID(r)
	if err != nil {
		return php_serialize.PhpArray{}, err
	}
	// no cookie/session yet, treat as guest
	if id == "" {
		return php_serialize.PhpArray{}, nil
	}
	return c.Session(id)
}

func (c *Client) GetSessionID(r *http.Request) (string, error) {
	if c.Config.Cookie == "" {
		panic("Session Cookie name required")
	}
	cookie, err := r.Cookie(c.Config.Cookie)
	if err != nil {
		// todo log, cookie probably doesn't exist
		return "", err
	}
	value, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		// todo log, cookie malformed
		return "", err
	}
	cookieValue, err := golaravelsession.GetSessionID(value, c.Config.Key)
	if err != nil {
		// todo log, cookie malformed
		return "", err
	}

	// remove cookie prefix from value
	parsed := strings.Split(cookieValue, "|")
	if len(parsed) != 2 {
		// todo log, value malformed
		return "", errors.New("cookie value missing |")
	}

	return parsed[1], nil
}

func (c *Client) Session(id string) (php_serialize.PhpArray, error) {
	id = c.Config.Prefix + ":" + id
	data, err := c.Redis.Get(context.Background(), id).Result()
	if err != nil {
		return php_serialize.PhpArray{}, err
	}
	return golaravelsession.ParseSessionData(data)
}
