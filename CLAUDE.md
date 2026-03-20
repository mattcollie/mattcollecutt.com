# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Personal portfolio website built with SvelteKit 5, Tailwind CSS v4, and shadcn-svelte. Deployed via Cloudflare Pages.

## Development Commands

```bash
# Run dev server
cd ui/portfolio && npm run dev

# Build for production
cd ui/portfolio && npm run build

# Preview production build
cd ui/portfolio && npm run preview
```

## Architecture

All source code lives under `ui/portfolio/`.

- **src/routes/** — SvelteKit file-based routing. `+page.svelte` (home), `+layout.svelte` (root layout with sidebar), `api/health/` and `api/ping/` (JSON endpoints)
- **src/lib/components/** — Reusable Svelte components: `Sidebar.svelte`, `SidebarLink.svelte`, `SidebarLinkGroup.svelte`, `Banner.svelte`
- **src/lib/components/ui/** — shadcn-svelte primitives (`button`, `sheet`)
- **src/lib/utils.ts** — `cn()` helper for Tailwind class merging
- **src/app.css** — Tailwind v4 theme config with custom color palette and shadcn CSS variable overrides
- **src/app.html** — HTML shell with meta tags, Google Fonts, PostHog analytics
- **static/** — Favicon, SVG icons, robots.txt

## Key Patterns

- SvelteKit 5 with Svelte runes (`$state`, `$derived`, `$props`, `$bindable`)
- Tailwind CSS v4 with `@theme` directive for custom tokens
- `adapter-cloudflare` for Cloudflare Pages deployment
- Mobile navigation uses Sheet component (slide-out drawer), desktop uses fixed sidebar

## Deployment

- Deployed via Cloudflare Pages (connected to repo, auto-deploys on push)
- Build command: `npm run build` with root directory `ui/portfolio`
