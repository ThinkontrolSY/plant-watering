import "./App.css";
import { useChannelsQuery, useWaterMutation } from "./generated/graphql";
import Button from "@mui/material/Button";
function App() {
  const { data, error } = useChannelsQuery();

  const [water] = useWaterMutation();

  if (error) {
    return <span>Error: {error.message}</span>;
  }

  return (
    <>
      {data?.channels.map((channel) => (
        <span key={channel}>{channel}</span>
      ))}
      <Button
        onClick={() =>
          water({
            variables: {
              input: {
                channel: "Water",
                seconds: 5,
              },
            },
          })
        }
      >
        Water
      </Button>
    </>
  );
}

export default App;
