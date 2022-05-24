/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"magma/orc8r/cloud/go/obsidian"
	swagger_protos "magma/orc8r/cloud/go/obsidian/swagger/protos"
	swagger "magma/orc8r/cloud/go/obsidian/swagger/servicers/protected"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/service"
	"magma/orc8r/cloud/go/services/analytics"
	analytics_protos "magma/orc8r/cloud/go/services/analytics/protos"
	analytics_servicers "magma/orc8r/cloud/go/services/analytics/servicers/protected"
	builder_protos "magma/orc8r/cloud/go/services/configurator/mconfig/protos"
	exporter_protos "magma/orc8r/cloud/go/services/metricsd/protos"
	"magma/orc8r/cloud/go/services/orchestrator"
	analytics_service "magma/orc8r/cloud/go/services/orchestrator/analytics"
	"magma/orc8r/cloud/go/services/orchestrator/obsidian/handlers"
	"magma/orc8r/cloud/go/services/orchestrator/servicers"
	protected_servicers "magma/orc8r/cloud/go/services/orchestrator/servicers/protected"
	indexer_protos "magma/orc8r/cloud/go/services/state/protos"
	streamer_protos "magma/orc8r/cloud/go/services/streamer/protos"
	"magma/orc8r/lib/go/service/config"
)

const (
	// Set max msg received to 50MB
	DefaultMaxGRPCMsgRecvSize = 50 * 1024 * 1024
)

func main() {
	srv, err := service.NewOrchestratorService(
		orc8r.ModuleName,
		orchestrator.ServiceName,
		grpc.MaxRecvMsgSize(DefaultMaxGRPCMsgRecvSize),
	)
	if err != nil {
		glog.Fatalf("Error creating orchestrator service %s", err)
	}

	obsidian.AttachHandlers(srv.EchoServer, handlers.GetObsidianHandlers())

	var exporterServicer exporter_protos.MetricsExporterServer

	var serviceConfig orchestrator.Config
	_, _, err = config.GetStructuredServiceConfig(orc8r.ModuleName, orchestrator.ServiceName, &serviceConfig)
	if err != nil {
		glog.Infof("err %v failed parsing the config file ", err)
		return
	}

	// workaround for https://github.com/spf13/viper/issues/324
	v := viper.NewWithOptions(viper.KeyDelimiter(":~~~:"))
	v.SetConfigName("orchestrator.yml")
	v.SetConfigType("yaml")
	v.AddConfigPath("/etc/magma/configs/orc8r/")
	v.AddConfigPath("/etc/magma/configs")
	v.AddConfigPath("/etc/magma")
	v.AddConfigPath("/var/opt/magma/configs")

	if err := v.ReadInConfig(); err != nil {
		glog.Errorf("err %v failed parsing the config file ", err)
		return
	}
	var viperConfig orchestrator.ViperConfig
	if err := v.Unmarshal(&viperConfig); err != nil {
		glog.Errorf("err %v failed unmarshalling the config file ", err)
		return
	}
	viperJson, _ := json.MarshalIndent(viperConfig, "", "\t")
	serviceJson, _ := json.MarshalIndent(serviceConfig, "", "\t")
	glog.Infof("VIPER: %s\n", viperJson)
	glog.Infof("OLD: %s\n", serviceJson)

	if serviceConfig.UseGRPCExporter {
		grpcAddress := serviceConfig.PrometheusGRPCPushAddress
		exporterServicer = protected_servicers.NewGRPCPushExporterServicer(grpcAddress)
	} else {
		exporterServicer = protected_servicers.NewPushExporterServicer(serviceConfig.PrometheusPushAddresses)
	}
	builder_protos.RegisterMconfigBuilderServer(srv.ProtectedGrpcServer, protected_servicers.NewBuilderServicer())
	exporter_protos.RegisterMetricsExporterServer(srv.ProtectedGrpcServer, exporterServicer)
	indexer_protos.RegisterIndexerServer(srv.ProtectedGrpcServer, protected_servicers.NewIndexerServicer())
	streamer_protos.RegisterStreamProviderServer(srv.ProtectedGrpcServer, servicers.NewProviderServicer())

	swagger_protos.RegisterSwaggerSpecServer(srv.ProtectedGrpcServer, swagger.NewSpecServicerFromFile(orchestrator.ServiceName))

	collectorServicer := analytics_servicers.NewCollectorServicer(
		&serviceConfig.Analytics,
		analytics.GetPrometheusClient(),
		analytics_service.GetAnalyticsCalculations(&serviceConfig.Analytics),
		nil,
	)
	analytics_protos.RegisterAnalyticsCollectorServer(srv.ProtectedGrpcServer, collectorServicer)

	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error while running service and echo server: %s", err)
	}
}
