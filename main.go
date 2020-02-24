package main

import (
    "fmt"
    "os"
    "io"
    "io/ioutil"
    "encoding/json"
    "strings"
    "strconv"
)


//feel free to edit to match you json needs

type Data struct {
    Name          string `json:"Name"`
    Contact       string `json:"Contact"`
    Address       string `json:"Address"`
    City          string `json:"City"`
    State         string `json:"State"`
    Zip           int    `json:"Zip"`
    Email         string `json:"Email"`
    Phone         string `json:"Phone"`
    Size          string `json:"Size"`
    Notes         string `json:"Notes"`
    Description   string `json:"Description"`
    Accessibility string `json:"Accessibility"`
    Facebook      string `json:"Facebook"`
    Instagram     string `json:"Instagram"`
    Twitter       string `json:"Twitter"`
    Website       string `json:"Website"`
    Features      string `json:"Features"`
    Image         string `json:"image"`
}   

func main() {
    // Open jsonFile
    jsonFile, err := os.Open("./data.json")
    
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("JSON Opened") 
    defer jsonFile.Close()

    //convert to byte
    jsonByte, _ := ioutil.ReadAll(jsonFile) 

    var data []Data

    err = json.Unmarshal(jsonByte, &data)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(data) 
    for _, dataItem := range data {
        //metaData add double quote
        fmt.Println(dataItem.Name)
        title := "\"" + dataItem.Name + "\""
        description := dataItem.Description
        contact := "\"" + dataItem.Contact + "\""
        address := "\"" + dataItem.Address + "\""
        city := "\"" + dataItem.City + "\""
        state := "\"" + dataItem.State + "\""
        zip := dataItem.Zip
        email := "\"" + dataItem.Email + "\""
        phone := "\"" + dataItem.Phone + "\""
        size := "\"" + dataItem.Size + "\""
        notes := "\"" + dataItem.Notes + "\""
        accessibility := "\"" + dataItem.Accessibility + "\""
        website := "\"" + dataItem.Website + "\""

        zipstring := strconv.Itoa(zip)
        slug := strings.ToLower(strings.ReplaceAll(dataItem.Name, " ", "-"))


        //markdown file structure
        text := "---\n" +
                "title: "+ title + "\n" +
                "contact: "+ contact+ "\n" +
                "address: "+ address + "\n" +
                "city: "+ city + "\n" +
                "state: "+ state + "\n" +
                "zip: "+ zipstring + "\n" +
                "email: "+ email + "\n" +
                "phone: "+ phone + "\n" +
                "size: "+ size + "\n" +
                "website: "+ website + "\n" +
                "accessibility: "+ accessibility + "\n" +
                "notes: "+ notes + "\n" +
                "--- \n" +
                description

        if err := WriteStringToFile(slug +".md", text); err != nil {
          panic(err)

        }

    }
}

func WriteStringToFile(filepath, s string) error {
  fo, err := os.Create("./Content/" + filepath)
  if err != nil {
    return err
  }
  defer fo.Close()

  _, err = io.Copy(fo, strings.NewReader(s))
  if err != nil {
    return err
  }
  return nil
}

