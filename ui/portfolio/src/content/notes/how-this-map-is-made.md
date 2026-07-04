---
title: How the map on this site is made
date: 2026-07-05
draft: true
---


> DRAFT SKELETON — the facts below are accurate; the voice is not yours yet.
> Rewrite every sentence you keep, delete what bores you, and add the
> personal parts only you know (why these two hills, what Te Aroha was like
> to grow up under, what the Hunua country is like now). The code samples and
> numbers can stay as-is.

The lines behind my name on the front page are real contour lines of two
real places: Te Aroha and the Kaimai Range, [where I grew up — your words],
and the Hunua Ranges above Papakura, [home ground these days — your words]. The site picks
one at random each visit.

> RAW MATERIAL — Matt's own words, transcribed July 2026, for the personal
> sections. Shape, don't replace:
>
> On Te Aroha: "We climbed it and typically try to every Christmas as well,
> but our big kitchen window/lounge looked up at the mountain too."
>
> On the Hunuas now: "A view, a place to hike/walk, it does also make us
> feel not-Auckland — we take a small walk around the corner and we see
> rolling hills/forests, cows, sheep, birds."

## The data

[Currently:] The elevation comes from NASA's Shuttle Radar Topography
Mission (SRTM) — radar flown on Endeavour in February 2000 that mapped the
Earth between 60°N and 56°S at roughly 30 m resolution. It's served today
as "terrarium" tiles from a public AWS bucket: PNG images where each
pixel's elevation in metres is encoded in its colour channels:

    elevation = R * 256 + G + B / 256 - 32768

[Soon / now:] The contours come from Toitū Te Whenua LINZ — the NZ
Contours (Topo50) layer on the LINZ Data Service, which are the actual
20 m contour lines printed on the Topo50 map series. They're CC BY 4.0,
which is why the credit line sits in the page footer. [A sentence about
why sourcing from LINZ mattered to you — NZ data for NZ hills.]

## From elevation to lines

For the SRTM path, the pipeline is about a hundred lines of Python with
no dependencies:

1. Fetch a 3 × 2 block of zoom-12 tiles centred on the hill (~30 × 20 km).
2. Decode the PNGs and stitch a 768 × 512 grid of elevations.
3. Smooth with a 3 × 3 box blur — SRTM is noisy at single-pixel scale.
4. Run marching squares at each contour level (50 m steps for Te Aroha,
   whose summit is 952 m; 40 m for the lower Hunua country) and chain the
   segments into polylines.
5. Decimate points, normalize coordinates to a 1000-wide frame, and write
   the lot to a ~120 KB JSON file that ships with the site.

The browser draws the polylines on a `<canvas>`, lowest elevation first —
so on load the plains appear before the summit does. That's the whole
animation. [A line about why you kept it that quiet, if you agree with
keeping it that quiet.]

## Why bother

[Yours entirely. The honest answer from the design process: the site had
contour lines as decoration first, and they felt cheap until they became
facts about real places. Whatever your version of that thought is, it
goes here.]

---
*Source for the generator: `tools/contours/generate.py` in the site repo.*
