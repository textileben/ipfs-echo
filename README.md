# ipfs-echo

Designed to run on an IPFS host that should have a connection to a remote IPFS peer.

The program will loop and perform these steps.
1. Create a random message
2. Addd the message to the remote IPFS instance (without pinning)
3. Query the local instance for the object CID
4. Sleep

During each step, metrics and log entries are created detailing the success or failure.

## Generated metrics

```
# HELP ipfsecho_attempts IPFS Echo attemps
# TYPE ipfsecho_attempts counter
ipfsecho_attempts{status="fail",type="local"} 6
ipfsecho_attempts{status="fail",type="remote"} 1
ipfsecho_attempts{status="success",type="local"} 1
ipfsecho_attempts{status="success",type="remote"} 7
# HELP ipfsecho_response_histogram IPFS Echo response histogram
# TYPE ipfsecho_response_histogram histogram
ipfsecho_response_histogram_bucket{status="failed",type="local",le="0.005"} 0
ipfsecho_response_histogram_bucket{status="failed",type="local",le="0.01"} 0
ipfsecho_response_histogram_bucket{status="failed",type="local",le="0.025"} 0
ipfsecho_response_histogram_bucket{status="failed",type="local",le="0.05"} 0
ipfsecho_response_histogram_bucket{status="failed",type="local",le="0.1"} 0
ipfsecho_response_histogram_bucket{status="failed",type="local",le="0.25"} 0
ipfsecho_response_histogram_bucket{status="failed",type="local",le="0.5"} 0
ipfsecho_response_histogram_bucket{status="failed",type="local",le="1"} 0
ipfsecho_response_histogram_bucket{status="failed",type="local",le="2.5"} 0
ipfsecho_response_histogram_bucket{status="failed",type="local",le="5"} 0
ipfsecho_response_histogram_bucket{status="failed",type="local",le="10"} 6
ipfsecho_response_histogram_bucket{status="failed",type="local",le="+Inf"} 6
ipfsecho_response_histogram_sum{status="failed",type="local"} 30.001771236
ipfsecho_response_histogram_count{status="failed",type="local"} 6
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="0.005"} 0
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="0.01"} 0
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="0.025"} 1
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="0.05"} 1
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="0.1"} 1
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="0.25"} 1
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="0.5"} 1
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="1"} 1
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="2.5"} 1
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="5"} 1
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="10"} 1
ipfsecho_response_histogram_bucket{status="failed",type="remote",le="+Inf"} 1
ipfsecho_response_histogram_sum{status="failed",type="remote"} 0.012331138
ipfsecho_response_histogram_count{status="failed",type="remote"} 1
ipfsecho_response_histogram_bucket{status="success",type="local",le="0.005"} 0
ipfsecho_response_histogram_bucket{status="success",type="local",le="0.01"} 0
ipfsecho_response_histogram_bucket{status="success",type="local",le="0.025"} 0
ipfsecho_response_histogram_bucket{status="success",type="local",le="0.05"} 1
ipfsecho_response_histogram_bucket{status="success",type="local",le="0.1"} 1
ipfsecho_response_histogram_bucket{status="success",type="local",le="0.25"} 1
ipfsecho_response_histogram_bucket{status="success",type="local",le="0.5"} 1
ipfsecho_response_histogram_bucket{status="success",type="local",le="1"} 1
ipfsecho_response_histogram_bucket{status="success",type="local",le="2.5"} 1
ipfsecho_response_histogram_bucket{status="success",type="local",le="5"} 1
ipfsecho_response_histogram_bucket{status="success",type="local",le="10"} 1
ipfsecho_response_histogram_bucket{status="success",type="local",le="+Inf"} 1
ipfsecho_response_histogram_sum{status="success",type="local"} 0.026202911
ipfsecho_response_histogram_count{status="success",type="local"} 1
ipfsecho_response_histogram_bucket{status="success",type="remote",le="0.005"} 0
ipfsecho_response_histogram_bucket{status="success",type="remote",le="0.01"} 0
ipfsecho_response_histogram_bucket{status="success",type="remote",le="0.025"} 5
ipfsecho_response_histogram_bucket{status="success",type="remote",le="0.05"} 7
ipfsecho_response_histogram_bucket{status="success",type="remote",le="0.1"} 7
ipfsecho_response_histogram_bucket{status="success",type="remote",le="0.25"} 7
ipfsecho_response_histogram_bucket{status="success",type="remote",le="0.5"} 7
ipfsecho_response_histogram_bucket{status="success",type="remote",le="1"} 7
ipfsecho_response_histogram_bucket{status="success",type="remote",le="2.5"} 7
ipfsecho_response_histogram_bucket{status="success",type="remote",le="5"} 7
ipfsecho_response_histogram_bucket{status="success",type="remote",le="10"} 7
ipfsecho_response_histogram_bucket{status="success",type="remote",le="+Inf"} 7
ipfsecho_response_histogram_sum{status="success",type="remote"} 0.126760107
ipfsecho_response_histogram_count{status="success",type="remote"} 7
```