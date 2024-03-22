package session

import (
	golaravelsession "Job-Post-FE/srv/session/laravel"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
	"net/http"
)

type Config struct {
	Config redis.Options
	Key    string
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

func (c *Client) Get(r *http.Request) (php_serialize.PhpArray, error) {
	id, err := c.SessionID(r)
	if err != nil {
		return php_serialize.PhpArray{}, err
	}
	return c.Session(id)
}

func (c *Client) SessionID(r *http.Request) (string, error) {
	cookie, err := r.Cookie("appointment")
	if err != nil {
		return "", err
	}
	return golaravelsession.GetSessionID(cookie.Value, c.Config.Key)
}

func (c *Client) Session(id string) (php_serialize.PhpArray, error) {
	data, err := c.Redis.Get(context.Background(), id).Result()
	if err != nil {
		return php_serialize.PhpArray{}, err
	}
	return golaravelsession.ParseSessionData(data)
}
