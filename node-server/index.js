const express = require("express");

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (_, res) => {
  res.status(200).json({ message: "Welcome to node server" });
});

app.post("/post", (req, res) => {
  const body = req.body;

  res.status(200).json(body);
});

app.post("/postform", (req, res) => {
  const body = req.body;

  res.status(200).json(body);
});

app.listen("8000", () => {
  console.log("Server is listening on port 8000...");
});
