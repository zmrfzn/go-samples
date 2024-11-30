pre-req

- Clone Go-easy repo & tidy
- Setup directory for sample  apps


Add an ENV VAR for NEW_RELIC_LICENSE_KEY

Steps 

1. **rest-api** 
    
    run the below command from go-easy directory.  
    
    Note: This assumes the go-easy repo is cloned under sample-apps directory
    
    ```powershell
    go run . -path ..\2.rest-api\ -name rest-http -diff ..\2.rest-api\rest-api.diff -agent nrApp
    ```
    Look for the `rest-api.diff` under the `2.rest-api` directory


2. **gin** 
    
    ```powershell
    go run . -path ..\3.gin\ name rest-http -diff ..\3.rest-api\gin.diff -agent nrApp -debug true
    ```

    Look for `gin.diff` under the `3.gin` directory