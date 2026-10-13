<div align="center">
  <img src="public/logo/logo.gif" alt="Long-Toolkit CLI" width="820">
</div>

<p align="center"><big><big><strong>Long</strong>itudinal<strong>-Toolkit</strong> <em>CLI</em></big></big></p>

<div align="center">
  <img src="https://img.shields.io/static/v1?label=Go&message=1.25%2B&color=00ADD8&style=for-the-badge&logo=go&logoColor=white" alt="Go 1.25+">
  <img src="https://img.shields.io/static/v1?label=Bubble%20Tea&message=TUI&color=F97316&style=for-the-badge" alt="Bubble Tea TUI">
  <img src="https://img.shields.io/static/v1?label=Command%20Palette&message=ready&color=111827&style=for-the-badge" alt="Command Palette ready">
  <img src="https://img.shields.io/static/v1?label=Homebrew&message=available&color=2E7D32&style=for-the-badge&logo=homebrew&logoColor=white" alt="Homebrew available">
  <img src="https://img.shields.io/static/v1?label=Scoop%20Bucket&message=available&color=0EA5E9&style=for-the-badge" alt="Scoop Bucket available">
</div>

<div align="center">
  <a href="https://github.com/Long-Toolkit/long-toolkit">Primary end-user repository: long-toolkit</a> ·
  <a href="https://long-toolkit.readthedocs.io/en/latest/">Documentation</a> ·
  <a href="https://life-epi-psych.github.io">LEAP Group</a>
</div>

## <img src="public/icons/lucide/github.svg" width="32" alt="" /> About The Project

`Long-Toolkit` (short for Longitudinal-Toolkit) is a hub of toolkits for longitudinal data in Python, from data preparation and preprocessing to modelling and explainability. It is designed for social, medical and clinical researchers who work with repeated-measure data and need a stepping-stone path from raw cohort files to downstream modelling.

The main toolkit is the benchmarking toolkit (`benchmark_standard_ml`, `benchmark_longitudinal_ml` and `bench_standard_and_longitudinal_ml`). It ranks standard and longitudinal machine learning under one shared protocol to establish whether machine learning, and/or longitudinal machine learning, is worth it on a given longitudinal dataset compared with a statistical baseline. The explainability stage then explains a saved model, e.g., the best estimator of a benchmark, with SHAP values.

The initiative delivers two interconnected components. First, the [`long-toolkit`](https://github.com/Long-Toolkit/long-toolkit) Python package is the engine behind every tool, fully usable from scripts or notebooks through its application programming interface (API). Second, this repository ships the `long-toolkit` command, a fully interactive command-line interface (CLI), written in Go, for running and orchestrating the toolkit from start to finish without code.

Each stage (data preparation, data preprocessing, machine learning and explainability) opens straight onto its tools, which help researchers iterate quickly on their datasets. The toolset is built to grow through community contributions.

> [!CAUTION]
> This repository is the CLI component of the project. For the full end-to-end toolkit experience, workflows and primary documentation, head to [`long-toolkit`](https://github.com/Long-Toolkit/long-toolkit) and its [documentation](https://long-toolkit.readthedocs.io/en/latest/).

## <img src="public/icons/lucide/terminal.svg" width="32" alt="" /> Setup & Launch

### <img src="public/icons/lucide/tally-1.svg" width="24" alt="" /> Homebrew

```bash
brew tap Long-Toolkit/homebrew-tap
brew install long-toolkit
```

### <img src="public/icons/lucide/tally-2.svg" width="24" alt="" /> Scoop Bucket

```powershell
scoop bucket add long-toolkit https://github.com/Long-Toolkit/scoop-bucket
scoop install long-toolkit
```

### <img src="public/icons/lucide/tally-3.svg" width="24" alt="" /> From Source

```bash
git clone https://github.com/Long-Toolkit/CLI.git
cd CLI
make build
```

### <img src="public/icons/lucide/tally-4.svg" width="24" alt="" /> Launch It Anywhere

Homebrew and Scoop put `long-toolkit` on `PATH` automatically.

When building from source, run the installer target for the shell in use:

```bash
# bash
make install-bash

# fish
make install-fish
```

Then run it from any directory:

```bash
long-toolkit
```

The CLI drives the Python toolkit through a bridge. Run it from a project directory where the Python package is installed, for example after `uv add long-toolkit` (with `uv`, the state-of-the-art Python package manager).

## <img src="public/icons/lucide/command.svg" width="32" alt="" /> Novel Command Palette

We designed the CLI so that its UX stays smooth at every moment. The command palette is available from any screen and launches any tool instantly, without walking through each menu level by hand.

In practice, launch `long-toolkit`, open the palette with `:` or `Ctrl+P`, type a phrase like `benchmark longitudinal ml`, and press `Enter` to run it. Whilst typing, `Tab` auto-completes, `Ctrl+H` opens history, `Ctrl+L` clears input and `Esc` closes the palette.

## <img src="public/icons/lucide/goal.svg" width="32" alt="" /> What's Next?

Jump to the main repository and its documentation for workflows and end-user guidance:

- [`long-toolkit`](https://github.com/Long-Toolkit/long-toolkit)
- [Long-Toolkit documentation](https://long-toolkit.readthedocs.io/en/latest/)

## <img src="public/icons/lucide/fingerprint-pattern.svg" width="24" alt="" /> Security

Please do not share participant-level or restricted data in issues or pull requests.

Security policy and contact details:
- [`SECURITY.md`](SECURITY.md)

_Special thanks to [@charm.land](https://charm.land) for their amazing TUI framework!_
