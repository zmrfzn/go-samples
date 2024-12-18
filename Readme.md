pre-req

- Clone Go-easy repo & tidy
- Setup directory for sample  apps


Add an ENV VAR for NEW_RELIC_LICENSE_KEY

Steps 

1. **hello-world**

    Run the below command from [go-easy-instrumentation](./go-easy-instrumentation/) directory.  
    
    > **Note**: This assumes the [go-easy-instrumentation](https://github.com/newrelic/go-easy-instrumentation) repo is cloned under go-samples directory
    
    ```bash
    go run . -path ../1.hello-world/ -name hello-world -diff ../1.hello-world/greetings.diff -agent nrApp
    ```
    Look for the `rest-api.diff` under the `2.rest-api` directory

2. **rest-api** 
    
    run the below command from go-easy directory.  
    
    Note: This assumes the go-easy repo is cloned under go-samples directory
    
    ```bash
    go run . -path ../2.rest-api/ -name rest-http -diff ../2.rest-api/rest-api.diff -agent nrApp
    ```
    Look for the `rest-api.diff` under the `2.rest-api` directory


3. **multi-endpoint rest API** 
    
    ```bash
    go run . -path ../3.multi-rest/ name multi-rest-http -diff ../3.multi-rest-http/multi-rest-http.diff -agent nrApp 
    ```

    Look for `multi-rest-http.diff` under the `3.multi-rest` directory

4. **gRPC Client & Server**

    This requires 2 different commands, one of Client and another for server

    **Client**

    To instrument gRpc client, run the following command 

    ```bash
    go run . -path ../4.gRpc/client -agent nrApp -name grcp-client -diff ../4.gRpc/client/client.diff
    ```

    To instrument gRpc server, run the following command. 

        ```bash
    go run . -path ../4.gRpc/server -agent nrApp -name grcp-server -diff ../4.gRpc/server/server.diff
    ```