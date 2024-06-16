import {
  Card,
  CardActionArea,
  CardContent,
  CardMedia,
  Typography,
} from "@mui/material";
import "./App.css";
import { useChannelsQuery, useWeatherQuery } from "./generated/graphql";
import { weatherIcon } from "./weather";
import { Channel } from "./channel";
function App() {
  const { data: channelData } = useChannelsQuery();
  const { data: weatherData } = useWeatherQuery();

  return (
    <div className="container">
      <Card sx={{ maxWidth: 345 }}>
        <CardActionArea>
          <CardMedia
            component="img"
            height="140"
            image={weatherIcon(weatherData?.weather?.weather || "")}
            alt="green iguana"
          />
          <CardContent>
            <Typography gutterBottom variant="h5" component="div">
              今日天气 {weatherData?.weather?.weather}, 计划浇水{" "}
              {(weatherData?.weather?.waterPlanSec || 0) / 60} 分
            </Typography>
            <Typography variant="body2" color="text.secondary">
              白天最高温度 {weatherData?.weather?.dayTemperature}°C,
              夜晚最低温度 {weatherData?.weather?.nightTemperature}°C
              {weatherData?.weather?.windDirection}风{" "}
              {weatherData?.weather?.windPower}级
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
      {channelData?.channels.map((channel) => (
        <Channel key={channel} channel={channel} />
      ))}
    </div>
  );
}

export default App;
