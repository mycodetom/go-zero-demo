# go-zero-demo

#### Description
Introduction go-zero (https://go-zero.dev) provides a demo based on the official website's tutorials, specifically designed for educational purposes to learn how to use go-zero.

#### Software Architecture
Software Architecture Framework: go-zero (https://github.com/zeromicro/go-zero)

#### Installation

1.  Install etcd, MySQL, and Redis.
2.  Install Go 1.21.3. If a different version is used, update the version number in shorturl/go.mod accordingly.
3.  Run go mod tidy if any dependencies are missing.
4.  Modify the configuration files (located at 1. shorturl/api/etc/shorturl-api.yaml and 2. shorturl/rpc/transform/etc/transform.yaml) to replace the placeholders with your local addresses.

#### Instructions

1.  Usage Instructions Follow the tutorials on the official website and execute commands in the command line interface to verify functionality, as shown below
2.  Step 1: curl -i "http://localhost:8888/shorten?url=https://go-zero.dev"
3.  Step 2: curl -i "http://localhost:8888/expand?shorten=b0434f"

#### Note

1.  This demo is solely intended for educational purposes to learn how to use go-zero.

