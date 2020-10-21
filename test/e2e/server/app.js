const express = require("express");
const app = express();
const port = process.env.PORT;

app.get("/healthCheck", (req, res) => {
  console.log("healthChecked");
  res.send("Hello World!");
});

app.get("/", (req, res) => {
  res.send(`Proxy to port ${port}`);
});

app.listen(port, () => {
  console.log(`App listening at http://localhost:${port}`);
});
