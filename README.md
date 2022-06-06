# Weather Operator - Built with Operator SDK

## What are we talking about?
```bash
$ kubectl apply -f config/samples/london.yaml
$ kubectl logs pod/weather-report-london
Weather report: london

                Overcast
       .--.     +15(14) °C
    .-(    ).   → 15 km/h
   (___.__)__)  10 km
                0.0 mm
                                                       ┌─────────────┐
┌──────────────────────────────┬───────────────────────┤  Mon 06 Jun ├───────────────────────┬──────────────────────────────┐
│            Morning           │             Noon      └──────┬──────┘     Evening           │             Night            │
├──────────────────────────────┼──────────────────────────────┼──────────────────────────────┼──────────────────────────────┤
│      .-.      Light drizzle  │      .-.      Light drizzle  │  _`/"".-.     Light rain sho…│  _`/"".-.     Patchy rain po…│
│     (   ).    +13(12) °C     │     (   ).    +15(14) °C     │   ,\_(   ).   17 °C          │   ,\_(   ).   +14(15) °C     │
│    (___(__)   → 13-17 km/h   │    (___(__)   → 10-13 km/h   │    /(___(__)  → 6-7 km/h     │    /(___(__)  ↙ 3-5 km/h     │
│     ‘ ‘ ‘ ‘   2 km           │     ‘ ‘ ‘ ‘   2 km           │      ‘ ‘ ‘ ‘  10 km          │      ‘ ‘ ‘ ‘  10 km          │
│    ‘ ‘ ‘ ‘    0.2 mm | 80%   │    ‘ ‘ ‘ ‘    0.2 mm | 71%   │     ‘ ‘ ‘ ‘   0.5 mm | 88%   │     ‘ ‘ ‘ ‘   0.1 mm | 70%   │
└──────────────────────────────┴──────────────────────────────┴──────────────────────────────┴──────────────────────────────┘
Location: London [51.509648,-0.099076]

Follow @igor_chubin for wttr.in updates

```

## Deploy Operator
Deploy Operator
```bash
$ make deploy
```

## Create a WeatherService
create yaml file `config/samples/london.yaml`
```yaml
apiVersion: weatherservice.iqbal.com/v1alpha1
kind: WeatherService
metadata:
  name: london
spec:
  city: london # city name
  days: 2 # how many days weather report you want
```
apply yaml file
```bash
$ kubectl -n weather-operator-system apply -f config/samples/london.yaml
```

get logs from the pod
```bash
$ kubectl -n weather-operator-system logs pod/weather-report-london
Weather report: london

                Overcast
       .--.     +15(14) °C
    .-(    ).   → 15 km/h
   (___.__)__)  10 km
                0.0 mm
                                                       ┌─────────────┐
┌──────────────────────────────┬───────────────────────┤  Mon 06 Jun ├───────────────────────┬──────────────────────────────┐
│            Morning           │             Noon      └──────┬──────┘     Evening           │             Night            │
├──────────────────────────────┼──────────────────────────────┼──────────────────────────────┼──────────────────────────────┤
│      .-.      Light drizzle  │      .-.      Light drizzle  │  _`/"".-.     Light rain sho…│  _`/"".-.     Patchy rain po…│
│     (   ).    +13(12) °C     │     (   ).    +15(14) °C     │   ,\_(   ).   17 °C          │   ,\_(   ).   +14(15) °C     │
│    (___(__)   → 13-17 km/h   │    (___(__)   → 10-13 km/h   │    /(___(__)  → 6-7 km/h     │    /(___(__)  ↙ 3-5 km/h     │
│     ‘ ‘ ‘ ‘   2 km           │     ‘ ‘ ‘ ‘   2 km           │      ‘ ‘ ‘ ‘  10 km          │      ‘ ‘ ‘ ‘  10 km          │
│    ‘ ‘ ‘ ‘    0.2 mm | 80%   │    ‘ ‘ ‘ ‘    0.2 mm | 71%   │     ‘ ‘ ‘ ‘   0.5 mm | 88%   │     ‘ ‘ ‘ ‘   0.1 mm | 70%   │
└──────────────────────────────┴──────────────────────────────┴──────────────────────────────┴──────────────────────────────┘
Location: London [51.509648,-0.099076]

Follow @igor_chubin for wttr.in updates
```

## Automatically create WeatherService and get logs
Use below shell script to automatically create everything
```bash
./get-weather-report ${CITY} ${DAYS}
```
Example:
```bash
$ ./get-weather-report dehradun 1
[INFO] updating ./config/samples/dehradun.yaml
[INFO] Using namespace: weather-operator-system
weatherservice.weatherservice.iqbal.com/dehradun created
[INFO] Waiting for weather-report-dehradun pod to be ready...
[INFO] weather-report-dehradun pod is ready
Weather report: dehradun

      \   /     Clear
       .-.      +33(31) °C
    ― (   ) ―   ↘ 5 km/h
       `-’      10 km
      /   \     0.0 mm
                                                       ┌─────────────┐
┌──────────────────────────────┬───────────────────────┤  Mon 06 Jun ├───────────────────────┬──────────────────────────────┐
│            Morning           │             Noon      └──────┬──────┘     Evening           │             Night            │
├──────────────────────────────┼──────────────────────────────┼──────────────────────────────┼──────────────────────────────┤
│     \   /     Sunny          │     \   /     Sunny          │     \   /     Sunny          │     \   /     Clear          │
│      .-.      +35(33) °C     │      .-.      +43(45) °C     │      .-.      +38(37) °C     │      .-.      +30(28) °C     │
│   ― (   ) ―   ↑ 7-8 km/h     │   ― (   ) ―   ↗ 14-16 km/h   │   ― (   ) ―   ↗ 10-11 km/h   │   ― (   ) ―   ↙ 9-19 km/h    │
│      `-’      10 km          │      `-’      10 km          │      `-’      10 km          │      `-’      10 km          │
│     /   \     0.0 mm | 0%    │     /   \     0.0 mm | 0%    │     /   \     0.0 mm | 0%    │     /   \     0.0 mm | 0%    │
└──────────────────────────────┴──────────────────────────────┴──────────────────────────────┴──────────────────────────────┘
Location: Dehradun, Dehra Dūn, Uttarakhand, 248001, India [30.3255646,78.0436813]

Follow @igor_chubin for wttr.in updates
```

## Cleanup
```bash
$ make undeploy
```
