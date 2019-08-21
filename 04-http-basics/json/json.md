
# JSON Encoding / Decoding

## JSON Mapping with struct tags

```
type Book struct {
    Title string `json:"book_title"`
    PageCount int `json:"pages,string"`
    ISBN string `json:"-"`
    Authors []Name `json:"auths,omniempty"`
    Publisher string `json:",omniempty"`
    PublishDate time.Time `json:"pub_date"`
}
```

Maps the PageCount struct field to the JSON object key, "pages", and outputs the value as a string instead of a number.

The dash causes the ISBN field to be skipped during encoding and decoding.

Maps the Authors field to the JSON object key, "auths". The annotation, omniempty,
causes the field to be omitted if its value is nil.


## Marshal and Unmarshal
When the encoder encounters a value whose type
implements json.Marshaler, it delegates serialization of the value to the method
MarshalJSON defined in the Marshaller interface