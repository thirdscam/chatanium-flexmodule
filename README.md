# FlexModule
**Next Generation Module System for Chatanium**<br/>
We're completely stripping out Chatanium's existing Go Plugin-based module system and creating a new one here.

## Goals
![Diagram](https://github.com/user-attachments/assets/c9085d53-aabd-47ce-9b85-5b8254c94f08)

Once our PoC implementation for multi-provider (Backend) is finalized, this codebase will be adapted to Chatanium v2.

## Design

* **Module system integration based on gRPC (w. protobuf)**<br/>
Ultimately, Chatanium's goal is to deploy/develop/operate better chatbots. Therefore, the existing module system has some limitations (https://github.com/thirdscam/chatanium/issues/10), and fixing them is our top priority. To solve these problems, we will implement a module system based on gRPC, which will provide cross-platform compilation and better security.
