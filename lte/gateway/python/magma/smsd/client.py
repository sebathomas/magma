
from opentelemetry import trace
from opentelemetry.exporter.jaeger.thrift import JaegerExporter
from opentelemetry.instrumentation.grpc import GrpcInstrumentorServer, GrpcInstrumentorClient
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor, SimpleSpanProcessor
from magma.common.service_registry import ServiceRegistry
from lte.protos.subscriberdb_pb2_grpc import SubscriberDBStub
from orc8r.protos.common_pb2 import Void


# Connect to Jaeger
trace.set_tracer_provider(
    TracerProvider(
        resource=Resource.create({SERVICE_NAME: "fake-client"})
    )
)
jaeger_exporter = JaegerExporter(
    agent_host_name="192.168.60.1",  # Host IP address
    agent_port=6831,
)
trace.get_tracer_provider().add_span_processor(
    SimpleSpanProcessor(jaeger_exporter)
)
#GrpcInstrumentorServer().instrument()
instrumentor = GrpcInstrumentorClient().instrument()
#tracer = trace.get_tracer(__name__)


def main():
    print("Hallo")

    chan = ServiceRegistry.get_rpc_channel(
        'subscriberdb',
        ServiceRegistry.LOCAL,
    )
    stub = SubscriberDBStub(chan)
    stub.ListSubscribers(Void())
    print("Tschüss")



if __name__ == "__main__":
    main()
