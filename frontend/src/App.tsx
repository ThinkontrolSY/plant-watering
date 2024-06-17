import {
  Button,
  Card,
  CardContent,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  IconButton,
  Slider,
  Typography,
} from "@mui/material";
import SettingsIcon from "@mui/icons-material/Settings";
import "./App.css";
import {
  useChannelsQuery,
  useSetBaseTimeMutation,
  useWeatherQuery,
} from "./generated/graphql";
import { weatherIcon } from "./weather";
import { Channel } from "./channel";
import { useState } from "react";
function App() {
  const [open, setOpen] = useState(false);
  const { data: channelData } = useChannelsQuery();
  const { data: weatherData, refetch } = useWeatherQuery();
  const [setBasetime] = useSetBaseTimeMutation();

  return (
    <>
      <Dialog
        open={open}
        onClose={() => {
          setOpen(false);
        }}
        PaperProps={{
          component: "form",
          onSubmit: (event: React.FormEvent<HTMLFormElement>) => {
            event.preventDefault();
            const formData = new FormData(event.currentTarget);
            const formJson = Object.fromEntries((formData as any).entries());
            const baseTime = parseInt(formJson.baseTime);
            console.log(baseTime);
            if (baseTime > 0) {
              setBasetime({
                variables: {
                  baseTime,
                },
              })
                .then(() => {
                  refetch();
                  setOpen(false);
                })
                .catch((error) => {
                  console.error(error);
                });
            }
          },
        }}
      >
        <DialogTitle>设置基本浇水时长(秒)</DialogTitle>
        <DialogContent>
          <Slider
            style={{
              marginTop: 40,
            }}
            defaultValue={weatherData?.weather?.baseTime || 60}
            name="baseTime"
            aria-label="基本浇水时长(秒)"
            id="baseTime"
            min={0}
            max={120}
            step={1}
            valueLabelDisplay="on"
          />
        </DialogContent>
        <DialogActions>
          <Button
            onClick={() => {
              setOpen(false);
            }}
          >
            取消
          </Button>
          <Button type="submit">提交</Button>
        </DialogActions>
      </Dialog>
      <div className="container">
        <Card
          sx={{
            display: "flex",
            alignItems: "center",
            width: "100%",
          }}
        >
          <div
            style={{
              position: "relative",
              width: 160,
              height: 140,
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
            }}
          >
            <img
              style={{
                width: "100%",
              }}
              src={weatherIcon(weatherData?.weather?.weather || "")}
            />
          </div>
          <CardContent
            style={{
              flex: 1,
              position: "relative",
            }}
          >
            <Typography gutterBottom variant="h5" component="div">
              今天 {weatherData?.weather?.weather}
              <br />
              计划浇水
              {weatherData?.weather?.waterPlanSec || 0}秒
            </Typography>
            <Typography variant="body2" color="text.secondary">
              白天最高温度 {weatherData?.weather?.dayTemperature}°C
              <br />
              夜晚最低温度 {weatherData?.weather?.nightTemperature}°C
              <br />
              {weatherData?.weather?.windDirection}风{" "}
              {weatherData?.weather?.windPower}级
            </Typography>
            <IconButton
              color="primary"
              style={{
                position: "absolute",
                right: 20,
                bottom: 20,
              }}
              onClick={() => {
                setOpen(true);
              }}
            >
              <SettingsIcon />
            </IconButton>
          </CardContent>
        </Card>
        {channelData?.channels.map((channel) => (
          <Channel key={channel} channel={channel} />
        ))}
      </div>
    </>
  );
}

export default App;
