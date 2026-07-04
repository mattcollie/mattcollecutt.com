#!/usr/bin/env python3
"""Generate contour-line JSON for the site's hero map.

Two data sources:

  srtm   NASA SRTM elevation via the public AWS Terrain Tiles bucket
         (terrarium encoding). No key needed. ~30 m resolution.
         Attribution: "Elevation data: NASA SRTM, via AWS Terrain Tiles."

  linz   NZ Contours (Topo50), layer 50768 on the LINZ Data Service —
         the authoritative 20 m contours printed on Topo50 maps.
         Needs a free API key from https://data.linz.govt.nz (create an
         account, then Manage APIs > create key with "Full access to
         Data Service API"). Export LINZ_API_KEY before running.
         License: CC BY 4.0. Attribution: "Contours sourced from the
         LINZ Data Service and licensed for reuse under CC BY 4.0."

Examples:

  # the two current maps, from SRTM (what the site ships today)
  python3 generate.py --source srtm --lat -37.544 --lon 175.741 --interval 50 --name teAroha
  python3 generate.py --source srtm --lat -37.09 --lon 175.02 --interval 40 --name redHill

  # the same frames from LINZ Topo50 (20 m native; keep every 2nd/3rd line)
  LINZ_API_KEY=... python3 generate.py --source linz --lat -37.544 --lon 175.741 --interval 40 --name teAroha

Each run writes/updates one named map inside --out (default:
ui/portfolio/static/map/contours.json), preserving the other maps.
Output schema per map: {"w":1000, "h":<int>, "levels":[{"lv":m, "lines":[[[x,y],...],...]}]}
Coordinates are normalized so x spans 0..1000.
"""

import argparse
import io
import json
import math
import os
import struct
import sys
import urllib.request
import zlib

TILE_Z = 12
TILES_X = 3  # 3x2 slippy tiles ≈ 30 km x 20 km at this latitude
TILES_Y = 2


def tile_for(lat, lon, z=TILE_Z):
    n = 2 ** z
    x = (lon + 180) / 360 * n
    lr = math.radians(lat)
    y = (1 - math.log(math.tan(lr) + 1 / math.cos(lr)) / math.pi) / 2 * n
    return int(x), int(y)


def tile_bounds(x, y, z=TILE_Z):
    n = 2 ** z
    lon_w = x / n * 360 - 180
    lon_e = (x + 1) / n * 360 - 180
    lat_n = math.degrees(math.atan(math.sinh(math.pi * (1 - 2 * y / n))))
    lat_s = math.degrees(math.atan(math.sinh(math.pi * (1 - 2 * (y + 1) / n))))
    return lon_w, lat_s, lon_e, lat_n


def fetch(url):
    req = urllib.request.Request(url, headers={"User-Agent": "mattcollecutt.com contour generator"})
    with urllib.request.urlopen(req, timeout=60) as r:
        return r.read()


# --- minimal PNG decode (8-bit RGB/RGBA), no external deps -----------------

def decode_png_rgb(data):
    assert data[:8] == b"\x89PNG\r\n\x1a\n", "not a PNG"
    pos, w = 8, None
    idat = b""
    while pos < len(data):
        (ln,) = struct.unpack(">I", data[pos:pos + 4])
        typ = data[pos + 4:pos + 8]
        body = data[pos + 8:pos + 8 + ln]
        if typ == b"IHDR":
            w, h, bit, color = struct.unpack(">IIBB", body[:10])
            assert bit == 8 and color in (2, 6), "expected 8-bit RGB(A)"
            ch = 3 if color == 2 else 4
        elif typ == b"IDAT":
            idat += body
        pos += 12 + ln
    raw = zlib.decompress(idat)
    stride = w * ch
    out = bytearray(h * stride)
    prev = bytearray(stride)
    p = 0
    for row in range(h):
        f = raw[p]
        line = bytearray(raw[p + 1:p + 1 + stride])
        p += 1 + stride
        if f == 1:
            for i in range(ch, stride):
                line[i] = (line[i] + line[i - ch]) & 0xFF
        elif f == 2:
            for i in range(stride):
                line[i] = (line[i] + prev[i]) & 0xFF
        elif f == 3:
            for i in range(stride):
                a = line[i - ch] if i >= ch else 0
                line[i] = (line[i] + ((a + prev[i]) >> 1)) & 0xFF
        elif f == 4:
            for i in range(stride):
                a = line[i - ch] if i >= ch else 0
                b = prev[i]
                c = prev[i - ch] if i >= ch else 0
                pp = a + b - c
                pa, pb, pc = abs(pp - a), abs(pp - b), abs(pp - c)
                pr = a if (pa <= pb and pa <= pc) else (b if pb <= pc else c)
                line[i] = (line[i] + pr) & 0xFF
        out[row * stride:(row + 1) * stride] = line
        prev = line
    return w, h, ch, bytes(out)


def srtm_grid(lat, lon):
    """Stitch a 3x2 terrarium tile block centred near (lat, lon). Returns (grid, w, h)."""
    cx, cy = tile_for(lat, lon)
    xs = [cx - 1, cx, cx + 1]
    ys = [cy - 1, cy] if (cy % 2) else [cy, cy + 1]  # keep target row in frame
    W, H = 256 * TILES_X, 256 * TILES_Y
    grid = [[0.0] * W for _ in range(H)]
    for yi, ty in enumerate(ys):
        for xi, tx in enumerate(xs):
            url = f"https://s3.amazonaws.com/elevation-tiles-prod/terrarium/{TILE_Z}/{tx}/{ty}.png"
            w, h, ch, px = decode_png_rgb(fetch(url))
            for r in range(256):
                base = r * w * ch
                row = grid[yi * 256 + r]
                for c in range(256):
                    o = base + c * ch
                    row[xi * 256 + c] = px[o] * 256 + px[o + 1] + px[o + 2] / 256 - 32768
    return grid, W, H


def smooth(grid, w, h):
    out = [[0.0] * w for _ in range(h)]
    for y in range(h):
        for x in range(w):
            s = n = 0
            for dy in (-1, 0, 1):
                for dx in (-1, 0, 1):
                    yy, xx = y + dy, x + dx
                    if 0 <= yy < h and 0 <= xx < w:
                        s += grid[yy][xx]
                        n += 1
            out[y][x] = s / n
    return out


def marching_squares(grid, w, h, level):
    """Extract iso-lines at `level` as joined polylines (marching squares + segment chaining)."""
    segs = []
    def interp(a, b, va, vb):
        t = (level - va) / (vb - va) if vb != va else 0.5
        return (a[0] + (b[0] - a[0]) * t, a[1] + (b[1] - a[1]) * t)
    for y in range(h - 1):
        for x in range(w - 1):
            v = [grid[y][x], grid[y][x + 1], grid[y + 1][x + 1], grid[y + 1][x]]
            idx = sum(1 << i for i, val in enumerate(v) if val > level)
            if idx in (0, 15):
                continue
            c = [(x, y), (x + 1, y), (x + 1, y + 1), (x, y + 1)]
            e = {
                0: interp(c[0], c[1], v[0], v[1]), 1: interp(c[1], c[2], v[1], v[2]),
                2: interp(c[3], c[2], v[3], v[2]), 3: interp(c[0], c[3], v[0], v[3]),
            }
            table = {
                1: [(3, 0)], 2: [(0, 1)], 3: [(3, 1)], 4: [(1, 2)], 5: [(3, 0), (1, 2)],
                6: [(0, 2)], 7: [(3, 2)], 8: [(2, 3)], 9: [(2, 0)], 10: [(0, 1), (2, 3)],
                11: [(2, 1)], 12: [(1, 3)], 13: [(1, 0)], 14: [(0, 3)],
            }
            for a, b in table[idx]:
                segs.append((e[a], e[b]))
    # chain segments into polylines via greedy walk over shared endpoints
    from collections import defaultdict
    key = lambda p: (round(p[0] * 4), round(p[1] * 4))
    lines = []
    seg_used = [False] * len(segs)
    index = defaultdict(list)
    for i, s in enumerate(segs):
        index[key(s[0])].append(i)
        index[key(s[1])].append(i)
    for i, s in enumerate(segs):
        if seg_used[i]:
            continue
        seg_used[i] = True
        line = [s[0], s[1]]
        grew = True
        while grew:
            grew = False
            for j in index[key(line[-1])]:
                if seg_used[j]:
                    continue
                a, b = segs[j]
                if key(a) == key(line[-1]):
                    line.append(b); seg_used[j] = True; grew = True; break
                if key(b) == key(line[-1]):
                    line.append(a); seg_used[j] = True; grew = True; break
        lines.append(line)
    return lines


def build_srtm(lat, lon, interval):
    grid, w, h = srtm_grid(lat, lon)
    grid = smooth(grid, w, h)
    top = max(max(r) for r in grid)
    levels = []
    for lv in range(interval, int(top), interval):
        keep = []
        for ln in marching_squares(grid, w, h, lv):
            if len(ln) < 30:
                continue
            ln = ln[::4]
            keep.append([[int(round(px / w * 1000)), int(round(py / w * 1000))] for px, py in ln])
        levels.append({"lv": lv, "lines": keep})
    return {"w": 1000, "h": int(round(h / w * 1000)), "levels": levels}


def build_linz(lat, lon, interval):
    key = os.environ.get("LINZ_API_KEY")
    if not key:
        sys.exit("Set LINZ_API_KEY (free key from https://data.linz.govt.nz — Manage APIs).")
    cx, cy = tile_for(lat, lon)
    xs = [cx - 1, cx, cx + 1]
    ys = [cy - 1, cy] if (cy % 2) else [cy, cy + 1]
    w_lon, s_lat, _, _ = tile_bounds(xs[0], ys[1])
    _, _, e_lon, n_lat = tile_bounds(xs[2], ys[0])
    url = (
        f"https://data.linz.govt.nz/services;key={key}/wfs"
        f"?service=WFS&version=2.0.0&request=GetFeature"
        f"&typeNames=data.linz.govt.nz:layer-50768&outputFormat=json"
        # NB: LDS GeoServer treats a plain EPSG:4326 bbox as lon,lat (x,y) order;
        # lat,lon order silently returns an empty FeatureCollection with HTTP 200.
        f"&srsName=EPSG:4326&bbox={w_lon},{s_lat},{e_lon},{n_lat},EPSG:4326"
    )
    fc = json.loads(fetch(url))
    # project lon/lat into the same normalized frame as the SRTM path (web mercator)
    def merc_y(lat_):
        r = math.radians(lat_)
        return math.log(math.tan(math.pi / 4 + r / 2))
    x0, x1 = math.radians(w_lon), math.radians(e_lon)
    y0, y1 = merc_y(n_lat), merc_y(s_lat)
    MARGIN = 50  # px beyond the 0..1000 frame; WFS returns whole lines, unclipped
    by_level = {}
    for f in fc.get("features", []):
        lv = int(f["properties"].get("elevation", 0))
        if lv <= 0 or lv % interval:
            continue
        geoms = f["geometry"]
        coords_list = geoms["coordinates"] if geoms["type"] == "MultiLineString" else [geoms["coordinates"]]
        for coords in coords_list:
            if len(coords) < 8:
                continue
            pts = coords[::3]
            line = []
            for lon_, lat_ in pts:
                nx = (math.radians(lon_) - x0) / (x1 - x0)
                ny = (merc_y(lat_) - y0) / (y1 - y0)
                x, y = int(round(nx * 1000)), int(round(ny * 1000 * 2 / 3))
                if -MARGIN <= x <= 1000 + MARGIN and -MARGIN <= y <= 667 + MARGIN:
                    line.append([x, y])
                elif len(line) >= 8:  # left the frame: flush this visible segment
                    by_level.setdefault(lv, []).append(line)
                    line = []
                else:
                    line = []
            if len(line) >= 8:
                by_level.setdefault(lv, []).append(line)
    levels = [{"lv": lv, "lines": by_level[lv]} for lv in sorted(by_level)]
    return {"w": 1000, "h": 667, "levels": levels}


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--source", choices=["srtm", "linz"], default="srtm")
    ap.add_argument("--lat", type=float, required=True)
    ap.add_argument("--lon", type=float, required=True)
    ap.add_argument("--interval", type=int, default=50)
    ap.add_argument("--name", required=True, help="map key in the output JSON, e.g. teAroha")
    ap.add_argument("--out", default=os.path.join(os.path.dirname(__file__), "../../ui/portfolio/static/map/contours.json"))
    args = ap.parse_args()

    out_path = os.path.abspath(args.out)
    data = {}
    if os.path.exists(out_path):
        data = json.load(open(out_path))
    build = build_srtm if args.source == "srtm" else build_linz
    result = build(args.lat, args.lon, args.interval)
    if not any(l["lines"] for l in result["levels"]):
        sys.exit(
            f"Refusing to write: {args.source} returned no contour lines for "
            f"'{args.name}' — existing data in {out_path} left untouched."
        )
    data[args.name] = result
    data[args.name]["source"] = (
        "NASA SRTM via AWS Terrain Tiles" if args.source == "srtm"
        else "LINZ Data Service, NZ Contours (Topo50), CC BY 4.0"
    )
    os.makedirs(os.path.dirname(out_path), exist_ok=True)
    json.dump(data, open(out_path, "w"), separators=(",", ":"))
    n = sum(len(l["lines"]) for l in data[args.name]["levels"])
    print(f"{args.name}: {len(data[args.name]['levels'])} levels, {n} polylines -> {out_path}")


if __name__ == "__main__":
    main()
