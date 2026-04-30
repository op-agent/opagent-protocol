# OpAgent Protocol TypeScript SDK

Minimal TypeScript SDK for writing OpAgent daemon agents.

```ts
import { OpServer, StdioTransport, textContent } from "@op-agent/opagent-protocol";

const server = new OpServer({ name: "demo", version: "0.1.0" });

server.addAgent({ name: "demo" }, async (req) => ({
  agentID: req.params.agentID,
  content: textContent("hello")
}));

await server.run(new StdioTransport());
```

The stdio transport is newline-delimited JSON-RPC 2.0 to match the Go SDK.
