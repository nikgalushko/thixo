# CIDR functions

## cidrhost
Returns an IP address for a given host number within IP network address prefix.

```
cidrhost "10.12.127.0/20" 16
```

## cidrnetmask
Converts an IPv4 address prefix with CIDR notation to a subnet mask address.

```
cidrnetmask "172.16.0.0/12"
```

## cidrsubnet 
Returns a subnet address within IP network address prefix.

```
cidrsubnet "172.16.0.0/12" 4 2
```

### cidrsubnets
Returns a list of sequential IP address within a CIDR prefix.

```
cidrsubnets "10.1.0.0/16" 4 4 8 4
```