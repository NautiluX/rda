# rda

Command-line tool utility and repository for tracking projects and their current stage of maturity

See [Rendered Registry](/rendered-registry) for a Markdown representation of the Registry.

## Build RDA CLI

```bash
git clone git@github.com:NautiluX/rda
go build
```

## Propose a new project

```bash
./rda add project -n <Project Name> -a <Your Name> -s <Sponsor Name> -d <Description> -r <Reference/Repository> -e <epic link>
```

Commit results and create PR against this repository.

## Promote a project

This command will update the stage as following:

```
sandbox -> incubation
incubation -> graduated.
```

```bash
./rda promote project -i <Project ID> -e <epic link>
```

Project ID is the ID assigned to the RDA (e.g. `RDA0001`).

Commit results and create PR against this repository.

## Index

* [RDA0001 RDA Repository](/rendered-registry/RDA0001%20RDA%20Repository.md) (sandbox)


