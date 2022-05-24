/*
 Copyright 2020 The Magma Authors.

 This source code is licensed under the BSD-style license found in the
 LICENSE file in the root directory of this source tree.

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package orchestrator

import (
	"magma/orc8r/cloud/go/services/analytics/calculations"
)

// Config represents the configuration provided to lte service
type Config struct {
	UseGRPCExporter           bool                         `yaml:"useGRPCExporter"`
	PrometheusGRPCPushAddress string                       `yaml:"prometheusGRPCPushAddress"`
	PrometheusPushAddresses   []string                     `yaml:"prometheusPushAddresses"`
	Analytics                 calculations.AnalyticsConfig `yaml:"analytics"`
}

type ViperConfig struct {
	UseGRPCExporter           bool                 `mapstructure:"useGRPCExporter"`
	PrometheusGRPCPushAddress string               `mapstructure:"prometheusGRPCPushAddress"`
	PrometheusPushAddresses   []string             `mapstructure:"prometheusPushAddresses"`
	Analytics                 ViperAnalyticsConfig `mapstructure:"analytics"`
}

// AnalyticsConfig represents the configuration provided to the components
// implementing analytics collector service.
type ViperAnalyticsConfig struct {
	// MinUserThreshold sets the value below which aggregated user metrics
	// shouldn't be exported.
	MinUserThreshold int                          `mapstructure:"minUserThreshold"`
	Metrics          map[string]ViperMetricConfig `mapstructure:"metrics"`
}

type ViperMetricConfig struct {
	Register                bool              `mapstructure:"register"`
	Export                  bool              `mapstructure:"export"`
	Expr                    string            `mapstructure:"expr"`
	Labels                  map[string]string `mapstructure:"labels"`
	EnforceMinUserThreshold bool              `mapstructure:"enforceMinUserThreshold"`
	LogConfig               *ViperLogConfig   `mapstructure:"logConfig"`
}

type ViperLogConfig struct {
	// key value pair to perform Term search on elastic
	Tags map[string]string `mapstructure:"tags"`

	// Custom fields to match the query on
	Fields []string `mapstructure:"fields"`

	// Query specifies the actual query string
	Query string `mapstructure:"query"`
}
