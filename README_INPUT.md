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
./rda add project -n <Project Name> -a <Your Name> -s <Sponsor Name> -d <Description> -r <Reference/Repository>
```

Commit results and create PR against this repository.

## Promote a project

This command will update the stage as following:

```
sandbox -> incubating
incubating -> graduated.
```

```bash
./rda promote project -i <Project ID> 
```

Project ID is the ID assigned to the RDA (e.g. `RDA0001`).

Commit results and create PR against this repository.

## Index
