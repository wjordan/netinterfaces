# NetInterfaces

This small module is a Linux performance fix for `net.Interfaces()`.
On systems with many network interfaces, listing IP addresses across all interfaces
returned by `net.Interfaces()` is inefficient because it requires a separate netlink
call for each interface, and each netlink call dumps all addresses across all interfaces.

This patch fixes that by using a single netlink call to collect addresses across all interfaces.

## How to use

Import this module, then use `NetInterfaces()` in place of `net.Interfaces()`:

```go
import . "github.com/wjordan/netinterfaces"

interfaces, _ := NetInterfaces()
for _, iface := range interfaces {
    addrs, _ := iface.Addrs()
}
```
