package config

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Highway represents the highway configuration
type Highway struct {
	// ValidatorHost is the host of the validator
	ValidatorHost string `mapstructure:"validator_address" json:"validator_address" yaml:"validator_address"`

	// ValidatorRPC is the port of the validator for rpc
	ValidatorRPC int `mapstructure:"validator_rpc_port" json:"validator_rpc_port" yaml:"validator_rpc_port"`

	// ValidatorGRPC is the port of the validator for grpc
	ValidatorGRPC int `mapstructure:"validator_grpc_port" json:"validator_grpc_port" yaml:"validator_grpc_port"`

	// HighwayPort is the port that the gateway listens on
	HighwayPort int `json:"gateway_port" yaml:"gateway_port"`

	// HighwayHost is the host that the gateway listens on
	HighwayHost string `json:"host" yaml:"host"`

	// PostgresConnection is the connection string for the postgres database
	PostgresConnection string `json:"postgres_connection" yaml:"postgres_connection"`

	// RedisConnection is the connection string for the redis database
	RedisConnection string `json:"redis_connection" yaml:"redis_connection"`

	// SmtpHost is the host of the smtp server
	SmtpHost string `json:"smtp_host" yaml:"smtp_host"`

	// SmtpPort is the port of the smtp server
	SmtpPort int `json:"smtp_port" yaml:"smtp_port"`

	// SmtpUser is the user of the smtp server
	SmtpUser string `json:"smtp_user" yaml:"smtp_user"`

	// SmtpPassword is the password of the smtp server
	SmtpPassword string `json:"smtp_password" yaml:"smtp_password"`
}

// LoadConfig loads the configuration from the file
func LoadConfig() (*Highway, error) {
	viper.SetConfigName("highway")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Highway
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// NewHway returns a new HighwayOptions
func NewHway() *Highway {
	v := viper.New()
	v.SetEnvPrefix("HWAY")
	v.AutomaticEnv()
	conf := &Highway{
		HighwayPort: 8000,
		HighwayHost: "0.0.0.0",
	}
	return conf
}

func (o *Highway) ReadFlags(c *cobra.Command) error {
	host, err := c.Flags().GetString("hway-host")
	if err != nil {
		return err
	}
	o.HighwayHost = host

	port, err := c.Flags().GetInt("hway-port")
	if err != nil {
		return err
	}
	o.HighwayPort = port

	psql, err := c.Flags().GetString("hway-psql")
	if err != nil {
		return err
	}
	o.PostgresConnection = psql

	redis, err := c.Flags().GetString("hway-redis")
	if err != nil {
		return err
	}
	o.RedisConnection = redis
	valHost, err := c.Flags().GetString("val-host")
	if err != nil {
		return err
	}
	o.ValidatorHost = valHost

	valRpc, err := c.Flags().GetInt("val-rpc")
	if err != nil {
		return err
	}
	o.ValidatorRPC = valRpc

	valGrpc, err := c.Flags().GetInt("val-grpc")
	if err != nil {
		return err
	}
	o.ValidatorGRPC = valGrpc
	return nil
}

// ListenAddress returns the address that the gateway listens on
func (o *Highway) ListenAddress() string {
	return fmt.Sprintf("%s:%v", o.HighwayHost, o.HighwayPort)
}

// ValidatorGRPCAddress returns the address of the validator grpc
func (o *Highway) ValidatorGRPCAddress() string {
	return fmt.Sprintf("%s:%v", o.ValidatorHost, o.ValidatorGRPC)
}

// ValidatorRPCAddress returns the address of the validator rpc
func (o *Highway) ValidatorRPCAddress() string {
	return fmt.Sprintf("%s:%d", o.ValidatorHost, o.ValidatorRPC)
}

// PrintBanner prints the banner
func (o *Highway) PrintBanner() {
	pterm.DefaultHeader.Printf(persistentBanner(fmt.Sprintf("localhost:%d", o.HighwayPort)))
}

// Serve starts the highway server
func (o *Highway) Serve(e *echo.Echo) {
	o.PrintBanner()
	e.Logger.Fatal(e.Start(o.ListenAddress()))
}

// Validate validates the HighwayOptions
func (o *Highway) Validate() error {
	if o.HighwayPort < 0 {
		return fmt.Errorf("gateway port must be greater than 0")
	}
	if o.HighwayHost == "" {
		return fmt.Errorf("host must not be empty")
	}
	return nil
}

func persistentBanner(address string) string {
	return fmt.Sprintf(`
Sonr Highway
· Gateway: http://%s
· Node RPC: http://localhost:26657
`, address)
}
