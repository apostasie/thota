# Away

"Away" defines an ad-hoc interchange format to store lists of places information in a simple, human-readable format.
It is agnostic, and does not depend on any particular service.

Example:

```json
  {
    "name": "Bars",
    "description": "The bars I like to go to",
    "places": [
      {
        "name": "Betty Wine Bar \u0026 Bottle Shop",
        "address": "1103 T St, Sacramento, CA 95811, USA",
        "latitude": 38.569021,
        "longitude": -121.4965609
      },
      {
        "name": "Ro Sham Beaux",
        "address": "2413 J St, Sacramento, CA 95816, USA",
        "latitude": 38.5752552,
        "longitude": -121.4739301,
        "note": "Bla bla bla"
      }
    ]
  }
```

Note this format is meant to evolve, is not a panacea, and is mostly meant to be used internally.

If you are looking for a more robust format (also much more complex), consider using [GeoJSON](https://geojson.org/),
which is a widely accepted standard for representing geographic data.