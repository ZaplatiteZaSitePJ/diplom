package jwt

type Config struct {
	AccessSecret  string `toml:"access_secret"`
	RefreshSecret string `toml:"refresh_secret"`
	AccessTTL     int    `toml:"access_ttl"`
	RefreshTTL    int    `toml:"refresh_ttl"`
}