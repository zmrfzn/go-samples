[![New Relic Experimental header](https://github.com/newrelic/opensource-website/raw/develop/src/images/categories/Experimental.png)](https://opensource.newrelic.com/oss-category/#new-relic-experimental)

# DevRel - Go Easy Instrumentation with New Relic

[New Relic Go easy instrumentation](https://docs.newrelic.com/docs/apm/agents/go-agent/installation/install-automation-new-relic-go) does most of the work for you by suggesting changes to your source code that instrument your application with the New Relic Go agent. This tool is currently in preview. 

## Prerequisites

Before you begin, make sure you have:

1. The Go-Samples repository cloned locally
2. Go-easy instrumentation tool installed:
    ```bash
    go install github.com/newrelic/go-easy-instrumentation@latest
    ```
3. Required environment variables set:
    ```bash
    export NEW_RELIC_LICENSE_KEY="your_license_key"
    export NEW_RELIC_APP_NAME="your_app_name"
    ```

## Instrumenting Sample Applications

1. **Hello World**

     From the root directory, run:
     
     ```bash
     go-easy-instrumentation instrument 1.hello-world -o 1.hello-world/greeting.diff
     ```
     
     The instrumentation changes will be saved in [1.hello-world](1.hello-world/).

2. **REST API**

    Run the following command from the go-easy directory:
    
    > **Note**: Ensure the go-easy repository is cloned under the go-samples directory
    
    ```bash
    go-easy-instrumentation instrument 2.rest-api --output 2.rest-api/rest-api.diff
    ```
    Check for the `rest-api.diff` file in the [2.rest-api](2.rest-api/) directory.

3. **Multi-endpoint REST API with Gin**

    Run the following command:
    
    ```bash
    go-easy-instrumentation instrument 3.gin -o 3.gin/gin-rest.diff
    ```
    Check for the `gin-rest.diff` file in the [3.gin](3.gin/) directory.

4. **gRPC Client & Server**

    This requires two separate commands, one for the client and another for the server.

    **Client**

    To instrument the gRPC client, run the following command:
    
    ```bash
    go-easy-instrumentation instrument 4.gRpc/client --output 4.gRpc/client/grpc-client.diff
    ```

    **Server**

    To instrument the gRPC server, run the following command:
    
    ```bash
    go-easy-instrumentation instrument 4.gRpc/server --output 4.gRpc/server/grpc-server.diff
    ```