---
slug: buildkit
title: BuildKit
authors: [kbbgl]
tags: [docker, buildkit, buildx, caching]
---

BuildKit is the build engine used by `docker build` (and `docker buildx build`). It’s more than “a Dockerfile builder”: it executes a **content-addressed DAG** of build steps, which enables aggressive caching and parallel execution.

## Key concepts

### LLB (Low-Level Build definition)

BuildKit’s internal “IR” is **LLB**, which represents build steps as a directed acyclic graph (DAG) of filesystem operations (run, copy, mount, etc.). Because the vertices are content-addressed, identical steps (same inputs) can be reused from cache.

### Frontends (bring your own syntax)

A **frontend** translates some build definition format into LLB. The default is the Dockerfile frontend (selected via a `# syntax=` directive), but frontends can be custom (YAML, JSON, HCL, etc.) as long as they emit valid LLB via the BuildKit Gateway API.

### Solver + cache

The **solver** executes the LLB graph. Caching happens at the operation level (not just linear layers), which is why independent branches can be executed in parallel and reused across builds.

## Outputs (BuildKit can build “more than images”)

BuildKit can export build results in several formats via `--output`, for example:

- **Image**: `type=image` (build/push an image)
- **Local directory**: `type=local,dest=./out` (write the final filesystem to a folder)
- **Tarball**: `type=tar,dest=./out.tar`
- **OCI tarball**: `type=oci,dest=./out-oci.tar`

Example (export build output to a local directory instead of producing an image):

```bash
docker buildx build --output type=local,dest=./out .
```

## Using a custom frontend (example pattern)

If a project provides a frontend image that can parse a non-Dockerfile spec, you can point BuildKit at it and still use `docker buildx build`. One common pattern is passing a build file (e.g. `spec.yml`) and choosing the frontend via a build arg:

```bash
docker buildx build \
  -f spec.yml \
  --build-arg BUILDKIT_SYNTAX=example/frontend-image \
  --output type=local,dest=./out \
  .
```

## Practical caching notes

BuildKit caches can be:

- **Local**: reused on the same machine/runner
- **Inline**: embedded in the image metadata
- **Remote/registry**: shareable across CI runners by pushing cache to a registry

## Related ecosystem

Several tools build on BuildKit/LLB (e.g. Earthly, Dagger, Depot). If you like “build pipelines as code” with strong caching, they’re worth exploring.

## Sources

- [BuildKit: Docker's Hidden Gem That Can Build Almost Anything (Tuan-Anh Tran, 2026-02-25)](https://tuananh.net/2026/02/25/buildkit-docker-hidden-gem/)
