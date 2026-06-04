package main

// module:
// github.com/{user}/getjson

// folder structure:
// cmd/
//     main.go
// internal/
//     legacy/
//         legacy.go
//     getter/
//         get_json.go

// ---

// ALGO-

// main.go:
// FUNCTION main():
//     jsonData = `{"name": "Asha Rao","age":28,"active":true}`

//     value := setter.GetJson(jsonData, "name")
//     PRINT value

//     value := setter.GetJson(jsonData, "age")
//     PRINT value

// legacy.go:
// FUNC Legacy():
//     UNMARSHAL jsonData using encoding/json package
//     DEFINE struct for jsonData
//     PRINT value of fieldname
//     RETURN fieldname

// set_json.go:
// // Dont use any print statements inside the GetJson method
// FUNC GetJson(jsonData string, fieldName string) value any:
//     USE gjson package documentation to get value of a field
//     RETURN value

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////

// package main

// IMPORT GETJSON, SETJSON and JSON2YAML

// // CLI ALGO

// FUNCTION main():
//     // --- args ---
//     inputPath  = flag "--input"   (required)
//     outputPath = flag "--output"  (default "out.yaml")
//     PARSE flags

//     IF inputPath is empty:
//         PRINT "usage: jsontool --input <file> [--output <file>]"
//         EXIT 1

//     // --- load ---
//     jsonBytes, err = readFile(inputPath)
//     HANDLE err -> print + exit
//     jsonStr = string(jsonBytes)

//     fieldCount = gjson.Parse(jsonStr).Get("@this").... // or count keys, refer gjson docs
//     PRINT "Loaded", inputPath

//     // --- verify step --- If u r done with single field verification, u can optionally attempt to make this a loop, but not mandatory.
//     PRINT "── Verify field ──"
//     fieldName = prompt(reader, "Field to verify: ")
//     result = gjson.Get(jsonStr, fieldName)
//     IF NOT result.Exists():
//         PRINT "Field", fieldName, "not found"
//     ELSE:
//         PRINT "Current value:", result.String()
//         answer = prompt(reader, "Is this correct? (y/n): ")
//         IF answer == "n":
//             newVal = prompt(reader, "Enter new value: ")
//             jsonStr, err = setter.SetJson(jsonStr, fieldName, newVal) // reuse divakarans package
//             HANDLE err
//             PRINT "✓ Updated", fieldName

//     // --- update step ---
//     PRINT "── Update field ──"
//     fieldName = prompt(reader, "Field to update: ")
//     result = gjson.Get(jsonStr, fieldName)
//     IF NOT result.Exists():
//         PRINT "Field", fieldName, "not found"
//     ELSE:
//         PRINT "Current value:", result.String()
//         newVal = prompt(reader, "Enter new value: ")
//         jsonStr, err = setter.SetJson(jsonStr, fieldName, newVal)
//         HANDLE err
//         PRINT "✓ Updated", fieldName

//     // --- convert + write ---
//     PRINT "Converting to YAML..."
//     yamlStr, err = converter.Convert(jsonStr)   // reuse Harshitas package
//     HANDLE err
//     err = writeFile(outputPath, yamlStr)
//     HANDLE err
//     PRINT "✓ Wrote", outputPath
//     PRINT "Done."

// // helper used everywhere — the interactive prompt
// FUNCTION prompt(reader, label) RETURNS string:
//     PRINT label (no newline)
//     line = reader.ReadLine()
//     RETURN trimWhitespace(line)
