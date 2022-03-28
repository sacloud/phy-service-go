# phy-service-go

[![Go Reference](https://pkg.go.dev/badge/github.com/sacloud/phy-service-go.svg)](https://pkg.go.dev/github.com/sacloud/phy-service-go)
[![Tests](https://github.com/sacloud/phy-service-go/workflows/Tests/badge.svg)](https://github.com/sacloud/phy-service-go/actions/workflows/tests.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sacloud/phy-service-go)](https://goreportcard.com/report/github.com/sacloud/phy-service-go)


さくらの専用サーバPHY向け高レベルAPIライブラリ  

## 概要

さくらの専用サーバPHYのAPIをラップし、CRUD+L+Action操作を統一的な手順で行えるインターフェースを提供します。  

インターフェースの例:
```go
// サーバ操作の例
func (s *Service) Find(req *FindRequest) ([]*Server, error)
func (s *Service) FindWithContext(ctx context.Context, req *FindRequest) ([]*Server, error)

func (s *Service) Install(req *InstallRequest) error
func (s *Service) InstallWithContext(ctx context.Context, req *InstallRequest) error

func (s *Service) Power(req *PowerRequest) error
func (s *Service) PowerWithContext(ctx context.Context, req *PowerRequest) error

func (s *Service) Read(req *ReadRequest) (*Server, error)
func (s *Service) ReadWithContext(ctx context.Context, req *ReadRequest) (*Server, error)
```

以下のリソースに対応しています。

```console
.
├── dedicated-subnet
├── private-network
├── server
│   ├── port
│   └── port-channel
└── service
```

## License

`sacloud/phy-service-go` Copyright (C) 2022 [The sacloud/phy-service-go Authors](AUTHORS).

This project is published under [Apache 2.0 License](LICENSE.txt).
