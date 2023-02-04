package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json"
)

// From wttr.in or WorldWeatherOnline
var weatherCodesMap = map[string]string{
    "113": "Sunny",
    "116": "PartlyCloudy",
    "119": "Cloudy",
    "122": "VeryCloudy",
    "143": "Fog",
    "176": "LightShowers",
    "179": "LightSleetShowers",
    "182": "LightSleet",
    "185": "LightSleet",
    "200": "ThunderyShowers",
    "227": "LightSnow",
    "230": "HeavySnow",
    "248": "Fog",
    "260": "Fog",
    "263": "LightShowers",
    "266": "LightRain",
    "281": "LightSleet",
    "284": "LightSleet",
    "293": "LightRain",
    "296": "LightRain",
    "299": "HeavyShowers",
    "302": "HeavyRain",
    "305": "HeavyShowers",
    "308": "HeavyRain",
    "311": "LightSleet",
    "314": "LightSleet",
    "317": "LightSleet",
    "320": "LightSnow",
    "323": "LightSnowShowers",
    "326": "LightSnowShowers",
    "329": "HeavySnow",
    "332": "HeavySnow",
    "335": "HeavySnowShowers",
    "338": "HeavySnow",
    "350": "LightSleet",
    "353": "LightShowers",
    "356": "HeavyShowers",
    "359": "HeavyRain",
    "362": "LightSleetShowers",
    "365": "LightSleetShowers",
    "368": "LightSnowShowers",
    "371": "HeavySnowShowers",
    "374": "LightSleetShowers",
    "377": "LightSleet",
    "386": "ThunderyShowers",
    "389": "ThunderyHeavyRain",
    "392": "ThunderySnowShowers",
    "395": "HeavySnowShowers"}

// Data looks like
// {
//  current_condition: [
//    0: {
//      "temp_c": "18",
//      "weathercode": "395"
//       }
//    ]
// }
type WeatherData struct {
  Temp_C string
  WeatherCode string
}

type WeatherJSON struct {
  Current_condition []WeatherData
}

func getWeatherData(city string) (string, error) {
  
  url := fmt.Sprintf("https://wttr.in/%s?format=j1", city)
  
  resp, err := http.Get(url)
  if err != nil {return "", err}
  
  var dat WeatherJSON
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {return "", err}
  
  err = json.Unmarshal(body, &dat)
  if err != nil {return "", err}
  
  var weather = weatherCodesMap[dat.Current_condition[0].WeatherCode]
  var temperature = dat.Current_condition[0].Temp_C
  result := fmt.Sprintf("Today's weather is %s, temperature is: %s C°", weather, temperature)
  return result, nil
}
