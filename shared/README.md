# protocol wrappers

## Summary

This directory contains gRPC wrappers (written in Go) for FlexModule.

## Description

Each subdirectory contains a Go gRPC wrapper library implemented for a specific backend.  These libraries are used by both the runtime and modules to establish a common protocol for communication.

## Adding Support for Other Languages

Currently, only Go is supported, as the Proof of Concept (PoC) implementation is not yet complete.

Of course, you can generate gRPC clients in your language of choice using the proto definitions located in `/proto`.  By implementing wrappers for each backend (similar to those in the subdirectories), you can re-implement the FlexModule protocol in other languages.  (Naturally, you can exclude functions used exclusively by the runtime.)

### "I've created a FlexModule protocol wrapper in an unsupported language!  Please merge it into the repository."

While we appreciate the contribution, the protocol is expected to undergo numerous changes and significant evolution in the future.

Therefore, we are not adding support for new languages at this time.  We may be able to accept contributions after this project is merged into Chatanium v2 (which will take some time!).
