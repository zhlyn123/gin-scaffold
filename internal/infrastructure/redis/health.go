package redis

import "context"

func (c *Client) Health(ctx context.Context) error {
	return c.DB.Ping(ctx).Err()
}
