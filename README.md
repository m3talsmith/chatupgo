# Chatupgo

A stateless terminal based peer to peer chat application using GRPC and Go.

## Development

### Built and Tested Using

- [Go 1.22.1](https://go.dev)
- [Protocol Buffers 25.2](https://grpc.io/docs/protoc-installation/)
- [GNU Make 4.4.1](https://www.gnu.org/software/make/)
- [CMake 3.28.3](https://cmake.org/)
- [Git 2.43.2](https://git-scm.com/)

### Build

The dependencies can be installed using: `make deps`

To build the project run: `make build`
All binaries will be built then into the `target` directory.

To install run: `make install`

To uninstall run: `make uninstall`

There are a list of other tasks in the `Makefile` for those interested in the details.
