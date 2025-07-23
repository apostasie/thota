# WeGo

WeGo allows importing "collections" from a geoJSON file.

This package provides a simple way to create said geoJSON file from a takeaway.json file.

## Caveats

Unfortunately, WeGo seems to have many peculiarities and several issues regarding lists management.

1. Places, represented as features in the geoJSON format, seem to use a non-standard or
not well-supported format for their `properties` field, and the addition of custom keys (`title`).
2. It seems like only the desktop version of WeGo supports adding a note to a place.
3. It seems like only the mobile version of WeGo allows exporting lists to geoJSON, and it does not include the notes.
4. Deleting collections on Desktop does not delete them on mobile (maybe also vice versa?).
5. Overall, desktop and mobile lists do not appear compatible, with many issues if places have been added from one or the other
6. Import on desktop is completely broken.
