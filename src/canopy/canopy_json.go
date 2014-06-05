package main

import (
    "canopy/datalayer"
    "encoding/json"
    "time"
)

type jsonDeviceClassItem struct {
    Category string `json:"category"`
    Datatype string `json:"datatype"`
    MinValue float64 `json:"min_value"`
    MaxValue float64 `json:"max_value"`
    Description string `json:"description"`
    ControlType string `json:"control_type"`
}

type jsonDevices struct {
    Devices []jsonDevicesItem `json:"devices"`
}

type jsonDevicesItem struct {
    DeviceId string `json:"device_id"`
    FriendlyName string `json:"friendly_name"`
    ClassItems map[string]jsonDeviceClassItem `json:"sddl_class"`
}

type jsonSample struct {
    Time string `json:"t"`
    Value float64 `json:"v"`
}
type jsonSamples struct {
    Samples []jsonSample `json:"samples"`
}

func devicesToJson(devices []*datalayer.CassandraDevice) (string, error) {
    var out jsonDevices

    for _, device := range devices {
        outDeviceClass := device.GetSDDLClass()

        out.Devices = append(
            out.Devices, jsonDevicesItem{
                device.GetId().String(), 
                device.GetFriendlyName(),
                outDeviceClass})
    }

    jsn, err := json.Marshal(out)
    if err != nil {
        return "", err
    }
    return string(jsn), nil
}

func samplesToJson(samples []datalayer.SensorSample) (string, error) {
    out := jsonSamples{[]jsonSample{}}
    for _, sample := range samples {
        out.Samples = append(out.Samples, jsonSample{
            sample.Timestamp.Format(time.RFC3339),
            sample.Value});
    }

    jsn, err := json.Marshal(out)
    if err != nil {
        return "", err
    }
    return string(jsn), nil
}
