import React, { useState } from "react";
import { useWaterMutation, useWaterStatisticQuery } from "./generated/graphql";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  Slider,
  Typography,
} from "@mui/material";

export const Channel: React.FC<{ channel: string }> = ({ channel }) => {
  const [minute, setMinute] = useState(5);
  const { data } = useWaterStatisticQuery({
    variables: {
      channel: channel,
    },
    fetchPolicy: "network-only",
    pollInterval: 1000 * 60 * 5,
  });
  const [water] = useWaterMutation();
  return (
    <Card sx={{ maxWidth: 345 }}>
      <CardContent>
        <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
          {channel}
        </Typography>
        <Typography variant="h5" component="div">
          已自动浇水 {(data?.waterStatistic?.autoWatering || 0) / 60} 分钟
          已手动浇水 {(data?.waterStatistic?.manualWatering || 0) / 60} 分钟
        </Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          手动浇水时长(分)
        </Typography>
        <Slider
          defaultValue={5}
          min={0}
          max={10}
          onChange={(_, value) => {
            if (typeof value === "number") {
              setMinute(value);
            }
          }}
          step={1}
          valueLabelDisplay="on"
        />
      </CardContent>
      <CardActions>
        <Button
          size="small"
          onClick={() => {
            if (minute > 0) {
              water({
                variables: {
                  input: {
                    channel: channel,
                    seconds: minute * 60,
                  },
                },
              });
            }
          }}
        >
          手动浇水
        </Button>
      </CardActions>
    </Card>
  );
};
