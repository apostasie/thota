# Thota

> “garden” (borrowed from the Telugu tōṭa)
> 
> short for "take-home-take-away": a tool to turn Google map lists take-out into portable data
> that can be then imported into other maps applications.

Google take-out feature is insultingly bad for actually exporting your maps lists:
(broken) CSV files only contain the barebone names of places, alongside a Google Maps URL.

No geolocation, no address.

Typically, it looks like this:

```
Curio Espresso and Vintage Design Café,,https://www.google.com/maps/place/Curio+Espresso+and+Vintage+Design+Caf%C3%A9/data=!4m2!3m1!1s0x5ff8336e3fb4d789:0xb7961921c8dc217d,,
```

Needless to say, this is unusable for the purpose of doing anything with it without Google Maps, or for
importing it into other maps applications.

This project aims to fix that situtation by providing a tool able to export meaningful data from your Google Maps lists.

... along with adapters to ease importing this data into other maps applications (for now: WeGo).

## Usage

### Export from Google take-out

**MAKE SURE YOU REVIEW GOOGLE PRICING BEFORE DOING ANYTHING (READ BELOW).**

- export your Google Maps lists using the [Google Takeout](https://takeout.google.com/) feature
- get a Google Maps API key
- run the `takeout` command with the path to your exported data and your API key:

```bash
GOOGLE_API_KEY=XXXXXXXX go run ./cmd/thota takeout --source path-to-extracted-takeout-folder/Saved > takeaway.json
```

The resulting file `takeaway.json` will contain lists from the takeout and places with enough information.
Typically:

```json
{
    "name": "Curio Espresso and Vintage Design Café",
    "address": "Yasuechō, １−１３ Kanazawa-shi, １-13 安江町 金沢市 石川県 920-0854, Japan",
    "latitude": 36.5733357,
    "longitude": 136.6554211
}
```

### Generate WeGo Collections

- get a WeGo API key
- run the `wego` command with the path to your exported data and your API key:


```bash
WEGO_API_KEY=XXXXXXXX go run ./cmd/thota wego --source takeaway.json
```

This will generate `wego-X.json` files, each file being a WeGo collection containing the places from the takeout.

## Debugging

Use `LOG_LEVEL=debug|info|warn|error` to control log output.

## About cost

At the time of this writing:

https://developers.google.com/maps/billing-and-pricing/pricing#places-legacy-pricing

Thota uses Google Places API (Legacy), and does one "Place Details" request per location.

Google charges about 2 cents per request over 5000 requests.

If you do run Thota repeatedly for whatever reason, and you have hundreds (or thousands) of
places in your list, you will _very quickly_ incur a significant bill.

I personally find Google pricing absolutely outrageous (about 10 times what competitors charge).

WeGo pricing model is much more reasonable, and their free tier is generous enough, so, it is less likely
you will run in trouble on that part. Nevertheless, you should still review their pricing model and keep an eye on your usage.

## Disclaimer and caveats

This project is not affiliated with Google or WeGo in any way.
Stressing again: it is your responsibility to review the pricing of the services you use.

Also, there are many, many issues with the WeGo API.
Geocoding is not reliable, and will return wildly out results, or no results at all (while there is in their map).

Thota tries its best to handle these issues, but it is not perfect, and there is only so much that
can be done with the available APIs.

If you have issues, please open a ticket so that I can try to fix them.