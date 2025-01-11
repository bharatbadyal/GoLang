const express = require("express");
const app = express();
const PORT = 3000;

app.use(
  "/user",
  (req, res, next) => {
    console.log("First Handler");
    // res.send("responded");
    next();
  },
  [
    (req, res, next) => {
      console.log("INside the second route handler");
      next();
    },
    (req, res) => {
      console.log(
        "Inside the third route handler with the use of next function"
      );
      res.send("Responding from 3rd handler");
    },
  ]
);

app.listen(PORT, () => {
  console.log(`Server is sucessfully listning on Port no. ${PORT}`);
});
