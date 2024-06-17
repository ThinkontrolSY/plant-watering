import React, { useState } from "react";
import { useWaterMutation, useWaterStatisticQuery } from "./generated/graphql";
import {
  Alert,
  Button,
  Card,
  CardActions,
  CardContent,
  Slider,
  Typography,
} from "@mui/material";

export const Channel: React.FC<{ channel: string }> = ({ channel }) => {
  const [pst, setPst] = useState(5);
  const { data, refetch } = useWaterStatisticQuery({
    variables: {
      channel: channel,
    },
    fetchPolicy: "network-only",
    pollInterval: 1000 * 60 * 5,
  });
  const [water] = useWaterMutation();
  const [error, setError] = useState<Error | null>(null);
  const [done, setDone] = useState(false);
  return (
    <>
      {error && (
        <Alert
          severity="error"
          onClose={() => {
            setError(null);
          }}
        >
          {error?.message}
        </Alert>
      )}
      {done && (
        <Alert
          severity="success"
          onClose={() => {
            setDone(false);
          }}
        >
          浇水成功
        </Alert>
      )}
      <Card>
        <CardContent>
          <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
            {channel}
          </Typography>
          <Typography variant="h6" component="div">
            已浇水{" "}
            {(data?.waterStatistic?.autoWatering || 0) +
              (data?.waterStatistic?.manualWatering || 0)}
            秒
          </Typography>
          <Typography sx={{ mb: 1.5 }} color="text.secondary">
            手动浇水时长(秒)
          </Typography>
          <Slider
            defaultValue={5}
            min={0}
            max={60}
            onChange={(_, value) => {
              if (typeof value === "number") {
                setPst(value);
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
              if (pst > 0) {
                water({
                  variables: {
                    input: {
                      channel: channel,
                      seconds: pst,
                    },
                  },
                })
                  .then((v) => {
                    setDone(v.data?.water || false);
                    refetch();
                  })
                  .catch((e) => {
                    console.error(e);
                    setError(e);
                    refetch();
                  });
              }
            }}
          >
            手动浇水
          </Button>
        </CardActions>
      </Card>
    </>
  );
};
