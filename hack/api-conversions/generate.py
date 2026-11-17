#!/usr/bin/env python3

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

"""Derives the API conversion data of resources that got a new API version.

For every Terraform resource in the given list, the two versions of its CRD
are compared: the deprecated one and the storage one. The output says which
fields are renamed and which fields only one of the versions has, which is
what the conversions between the two versions are built from, see
config/registry_versions.go.

Run it from the repository root after `make generate`:

  hack/api-conversions/generate.py package/crds apis/cluster \\
      config/v5-breaking-apis.txt hack/api-conversions/v5-rename-candidates.tsv \\
      > config/v5-api-conversions.json
"""

import glob
import json
import os
import re
import sys

import yaml

SECTIONS = {
    "spec.forProvider": ("spec", "forProvider"),
    "spec.initProvider": ("spec", "initProvider"),
    "status.atProvider": ("status", "atProvider"),
}


def camel(segment):
    words = segment.split("_")
    return words[0] + "".join(w[:1].upper() + w[1:] for w in words[1:])


def shape(schema):
    """Returns the kind of value a schema node holds, ignoring its fields."""
    t = schema.get("type")
    if t == "array":
        return "array<" + shape(schema.get("items", {})) + ">"
    if t == "object":
        return "object" if "properties" in schema else "map"
    return t or "any"


def compare(prev, cur, path, result):
    """Records the top-most fields only one of the two schema nodes has."""
    pp, cp = prev.get("properties", {}), cur.get("properties", {})
    for name in sorted(set(pp) | set(cp)):
        p = ".".join(path + [name]).replace(".[*]", "[*]")
        if name not in cp:
            result["previousOnly"].append(p)
        elif name not in pp:
            result["currentOnly"].append(p)
        elif shape(pp[name]) != shape(cp[name]):
            # the identity conversion cannot copy such a field, and no
            # conversion registered from this data handles it
            sys.exit(f"{p}: {shape(pp[name])} became {shape(cp[name])}")
        elif shape(pp[name]) == "object":
            compare(pp[name], cp[name], path + [name], result)
        elif shape(pp[name]) == "array<object>":
            compare(pp[name]["items"], cp[name]["items"], path + [name, "[*]"], result)


def normalize(schema):
    """Strips what does not affect the stored value from a schema node."""
    s = {k: v for k, v in schema.items()
         if k not in ("description", "required", "x-kubernetes-validations")}
    if "properties" in s:
        s["properties"] = {k: normalize(v) for k, v in s["properties"].items()}
    if "items" in s:
        s["items"] = normalize(s["items"])
    return s


def node(schema, path):
    """Returns the schema node at path, or None if a list is on the way."""
    for name in path.split("."):
        if schema.get("type") != "object" or name not in schema.get("properties", {}):
            return None
        schema = schema["properties"][name]
    return schema


def conversions(crd, candidates):
    versions = {v["name"]: v["schema"]["openAPIV3Schema"]["properties"]
                for v in crd["spec"]["versions"]}
    prev = [v["name"] for v in crd["spec"]["versions"] if v.get("deprecated")]
    cur = [v["name"] for v in crd["spec"]["versions"] if v["storage"]]
    if len(prev) != 1 or len(cur) != 1:
        sys.exit(f'{crd["metadata"]["name"]}: expected one deprecated and one storage version')
    out = {"renames": [], "previousOnly": [], "currentOnly": []}
    for section, (top, sub) in SECTIONS.items():
        p = versions[prev[0]][top]["properties"][sub]
        c = versions[cur[0]][top]["properties"][sub]
        result = {"previousOnly": [], "currentOnly": []}
        compare(p, c, [], result)
        for old, new in candidates:
            if old not in result["previousOnly"] or new not in result["currentOnly"]:
                continue
            if normalize(node(p, old)) != normalize(node(c, new)):
                continue
            out["renames"].append([f"{section}.{old}", f"{section}.{new}"])
            result["previousOnly"].remove(old)
            result["currentOnly"].remove(new)
        for k in result:
            out[k] += [f"{section}.{f}" for f in result[k]]
    return out


def main(crd_dir, apis_dir, list_path, candidates_path):
    kinds = {}
    for f in glob.glob(os.path.join(apis_dir, "*", "*", "zz_*_terraformed.go")):
        with open(f) as s:
            src = s.read()
        kind = re.search(r"func \(mg \*(\w+)\) GetTerraformResourceType", src).group(1)
        name = re.search(r'return "(\w+)"', src).group(1)
        kinds[name] = (f.split(os.sep)[-3], kind)

    candidates = {}
    with open(candidates_path) as s:
        for line in s:
            if line.startswith("#") or not line.strip():
                continue
            name, old, new = line.rstrip("\n").split("\t")
            candidates.setdefault(name, []).append(
                (".".join(camel(x) for x in old.split(".")), ".".join(camel(x) for x in new.split("."))))

    crds = {}
    for f in glob.glob(os.path.join(crd_dir, "*.yaml")):
        with open(f) as s:
            crd = yaml.safe_load(s)
        group = crd["spec"]["group"].split(".")
        scope = "namespaced" if "m" in group[1:-2] else "cluster"
        crds[(scope, group[0], crd["spec"]["names"]["kind"])] = crd

    out = {}
    with open(list_path) as s:
        names = [line.strip() for line in s if line.strip()]
    for name in sorted(names):
        group, kind = kinds[name]
        out[name] = conversions(crds[("cluster", group, kind)], candidates.get(name, []))
        if conversions(crds[("namespaced", group, kind)], candidates.get(name, [])) != out[name]:
            sys.exit(f"{name}: the cluster and the namespaced CRDs differ")
    json.dump(out, sys.stdout, indent=1)
    sys.stdout.write("\n")


if __name__ == "__main__":
    if len(sys.argv) != 5:
        sys.exit("usage: generate.py <CRD dir> <cluster apis dir> <resource list> <rename candidates>")
    main(*sys.argv[1:])
