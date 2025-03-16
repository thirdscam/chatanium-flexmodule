# proto definitions and gRPC clients

## Summary

This directory contains the proto definitions for FlexModule.

## Description

This directory stores the written proto definitions, as well as the gRPC clients generated from them using `make buf`.
The generated gRPC clients are referenced in `/shared` and are used to create the FlexModule protocol wrappers.

We have also established some rules to improve the development experience and readability of the proto definitions.

### `hooks.proto`

This file defines the interface that allows the runtime to execute modules and trigger events within them.

### `helpers.proto`

This file defines the interface that allows modules to execute functions provided by the runtime (for example, `SendMessage()` in the case of Discord).
Therefore, this protocol is not simply unidirectional; it enables mutual function calls between the runtime and modules.
