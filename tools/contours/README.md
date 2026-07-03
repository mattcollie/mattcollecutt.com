# Contour map generator

Generates the contour-line JSON used by the site's hero map
(`ui/portfolio/src/lib/data/contours.json`). Pure Python 3 stdlib — no
dependencies to install.

The map currently holds two frames, both ~30 × 20 km at slippy-tile zoom 12:

| key       | place                                   | interval |
|-----------|------------------------------------------|----------|
| `teAroha` | Te Aroha & the Kaimai Range              | 50 m     |
| `redHill` | Red Hill, Papakura & the Hunua Ranges    | 40 m     |

## Regenerate from SRTM (no key needed)

```bash
python3 generate.py --source srtm --lat -37.544 --lon 175.741 --interval 50 --name teAroha
python3 generate.py --source srtm --lat -37.09  --lon 175.02  --interval 40 --name redHill
```

Data: NASA SRTM (~30 m) via the public AWS Terrain Tiles bucket
(terrarium encoding). Attribution used on the site:
"Elevation data: NASA SRTM, via AWS Terrain Tiles."

## Switch to LINZ Topo50 contours (authoritative)

1. Create a free account at https://data.linz.govt.nz
2. Manage APIs → create a key with "Full access to Data Service API"
3. Run:

```bash
LINZ_API_KEY=xxxx python3 generate.py --source linz --lat -37.544 --lon 175.741 --interval 40 --name teAroha
LINZ_API_KEY=xxxx python3 generate.py --source linz --lat -37.09  --lon 175.02  --interval 40 --name redHill
```

This queries **NZ Contours (Topo50)**, layer 50768, over WFS — the actual
20 m contour lines printed on Topo50 maps (`--interval 40` keeps every
second line). License is CC BY 4.0; the site must carry the attribution:

> Contours sourced from the LINZ Data Service and licensed for reuse
> under CC BY 4.0.

Note: the LINZ path was written against the documented WFS API but has not
been run yet (this sandbox cannot reach data.linz.govt.nz) — expect to
tweak the `elevation` property name or bbox axis order on first run.
Verify against a small bbox first.

Each map's `source` field in the JSON records which dataset it came from,
so the page caption can state it truthfully per map.
