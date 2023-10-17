# Text Numbers
## Convert numbers into multiple languages 
#### Overview
A simple application which takes a single decimal number and converts it into its writen form in a given language.  
Its aim is to make the language configuration as simple as possible, making it simple to add additional languages as required.  
  
#### Usage
`tn 1234`  
Outputs: `one thousand two hundred and thirty four`  
(Same as `tn 1234 as english`)  

`tn 1234 as french`  
Outputs: `un mille deux cent et trente quatre`  

`tn 1234 as dutch`  
Outputs: `een duizend twee honderd en vier en dertig`

  

### Languages
French, Dutch and English are provided with English being the default.  
Languages are expressed as json files, located in the `languages` directory, which must reside in the same directory as the executable.  
Use the sample files as an example of the structure.  
  
##### Language File
Languagle files define a single language.  They are json files, with a filename which reflects the language.  
i.e. `english.json`, `french.json` etc and must be in the `languages` directory.
  
The document should be a single json object with the following properties:  
`title`:   The name of the language  
`separator`: An optional string to insert between labels and digits. e.g. "AND" = one hundred AND one  
`minus`: The label to give negative numbers   
`invert-digits`:  An optional bool flag which flips the final digits, adding the seperater.  
e.g. 23 in Dutch = drie en twintig  (3 and twenty)  
`names`: An array of named values  
`labels`: An Array of value labels  
e.g.
>{  
    "title": "English",  
    "seperator": "and",  
    "minus": "minus",  
    "names": [...],  
    "labels": [...],  
}

#### Named Values
Named values are direct mappings of a single value to a string.  
The `names` array should contain an array of objects containing the following properties:  
`value`: The Value it represents   
`name`: The string name to assign to that value  
e.g.  
>{ "value": 2, "name": "two"}  

Generally, in most languages, all values up to twenty have their own name.  



