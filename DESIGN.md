## [Top level MIME types](https://tools.ietf.org/html/rfc2046#section-2) don't need to be separately encoded

Commonality is derived by full MIME type (application/octet-stream) not by top level type (application).

## Two-level MIME type encoding based on popularity

An uncommon MIME type that gets promoted to common gets *removed* from the uncommon list (in order to save space, which is the point of having a common list). There could be a client that knows and can parse a MIME type, but doesn't yet know it was promoted to the common list. It can recognize the known type in the uncommon list at e.g. tag 2374, but doesn't recognize common type 42... Promotion is thus not backward compatible...

Do enum aliases help here? Reserve 1 to 127 for common types, alias types into that range when they are promoted? Same problem? Something else?

Send both the uncommon and common value for a while?

3 step promotion process:

1. Add new type to uncommon list
2. Decide to promote type based on usage
3. Choose and reserve tag number for type in common list (but don't add it)
4. Add enumvalueoption to type in uncommon list indicating reserved common list tag number
5. Wait
6. Add type to common list at the reserved tag number, LEAVE it in the uncommon list; servers can sendclients not receiving this last update may start to see unknown common values, they should look them up in the options of the uncommon values

Very old clients will still break. If a client sees an unknown MIME type with tag < 128, it should definitely update its definitions as this type has been popular enough for long enough to now be common.
