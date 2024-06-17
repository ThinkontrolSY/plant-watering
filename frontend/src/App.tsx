import { Card, CardContent, Typography } from "@mui/material";
import "./App.css";
import { useChannelsQuery, useWeatherQuery } from "./generated/graphql";
import { weatherIcon } from "./weather";
import { Channel } from "./channel";
function App() {
  const { data: channelData } = useChannelsQuery();
  const { data: weatherData } = useWeatherQuery();

  return (
    <div className="container">
      <Card
        sx={{
          display: "flex",
          alignItems: "center",
          width: "100%",
        }}
      >
        <img
          style={{
            width: "140px",
          }}
          src={weatherIcon(weatherData?.weather?.weather || "")}
        />
        <CardContent>
          <Typography gutterBottom variant="h5" component="div">
            今天 {weatherData?.weather?.weather}
            <br />
            计划浇水
            {(weatherData?.weather?.waterPlanSec || 0) / 60}分
          </Typography>
          <Typography variant="body2" color="text.secondary">
            白天最高温度 {weatherData?.weather?.dayTemperature}°C
            <br />
            夜晚最低温度 {weatherData?.weather?.nightTemperature}°C
            <br />
            {weatherData?.weather?.windDirection}风{" "}
            {weatherData?.weather?.windPower}级
          </Typography>
        </CardContent>
      </Card>
      {channelData?.channels.map((channel) => (
        <Channel key={channel} channel={channel} />
      ))}
    </div>
  );
}

export default App;
