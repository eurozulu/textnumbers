# Text Numbers

## Convert numbers into multiple languages

#### Overview

A simple application which takes a single decimal number and converts it into its writen form in a given language.  
Its aim is to make the language configuration as simple as possible, making it simple to add additional languages as
required.

#### Usage

`tn 1234`  
Outputs: `one thousand two hundred and thirty four`  
(Same as `tn 1234 as english`)

`tn 1234 as french`  
Outputs: `un mille deux cent et trente quatre`

`tn 1234 as dutch`  
Outputs: `een duizend twee honderd en vier en dertig`

Supports positive and negitive numbers upto 18446744073709551615 (Unsigned 64 bit)  
`as` and `in` can either be used to specify the language.

### Languages

French, Dutch and English are provided with English being the default.  
Languages are expressed as json files, located in the `languages` directory, which must reside in the same directory as
the executable.  
Use the sample files as an example of the structure.

##### Language File

A Language file defines a single language. They are json files, with a filename which reflects the language name.  
i.e. `english.json`, `french.json` etc and must be in the `languages` directory.

The document should be a single json object with the following properties:  
`title`:   The name of the language  
`minus`: The label to give negative numbers  
`names`: A list of value names of digits or groups of digits  
`separators`: An optional list of seperators to insert between specific value digits.  
e.g.
> {  
"title": "English",  
"minus": "minus",  
"separators": [{"value":100, "name": "and"}],  
"names": [...],  
> }

#### Names
Value Names are string names mapped to a value or group of values.  They are defined with the following properties:    
`name`: The string name to assign to that value  
`value`: The numeric value it represents   
e.g.
> { "value": 2, "name": "two"}

The `names` list should contain an array of such valueNames, each with a unique value.  
Names may act as a single representation of a value or as a group of values.  This is defined by the name value itself.  
If the name contains the `%v` in the string, the NameValue will act as a group label, otherwise it acts as a single value.  
The above example "two" would act as a simple value, mapping the value 2 to that string only.  
> { "value": 100, "name": "%v hundred"}

This example of a group name, will map any number of 100 or greater to this name.  
1200 would map to twelve hundred.  
In most languages, all numbers up to twenty and thirty to ninety have names.  Hundred and above generally use group naming.  

#### Separator

A separator is a string inserted between numbers based on the values which surround it.
An example, in English is the word 'and', which seperates numbers => 100 from numbers below 100.  
e.g. Three hundred AND twenty one.  
The separator is defined with the following properties:  
`name`: The seperator text  
`value`: The dividing boundary in which to insert the serperator    
`reverse`: An optional bool flag to reverse the seperated digits order.    
e.g.
> { "value": 100, "name": "and"}

The value specifies a boundary. If the digit being formatted is equal or above that boundary, and  
the digits following it are below that boundary, the seperator is inserted.  
e.g. If formatting the number '1234', when formatting the 1000 part, no seperator is triggered
as the following number (200) is also above the boundary.
When formatting the 200, the seperator is triggered as the following digits, '34' are below the boundary.
the `two hundred` will be tagged with the seperator `and`, before the '23' is formatted.
When formatting the 20 and 3, the seperator is no longer triggered as both are below the boundary.

The reverse flag is used with languages which reverse the order of certain digits when spoken.  
e.g. 23 in Dutch = drie en twintig  (3 and twenty)
> `{"value":10, "name": "en", "reverse":  true}`

When a reverse seperator is triggered, the current digits being formatted are reversed with the
following digits and the seperator name inserted between them.

Using a 'value' of 10, digits greater than ten, which are followed by digits less than 10  
have the 'name' "en" inserted between them. In this case, the digits are first reversed
so the lower value digit preceeds the "en" and the higher one follows it.
Reverse seperators will often be used along side non reverse seperators, where the non reverse will be a higher value,
so limiting which digits get reversed.  

