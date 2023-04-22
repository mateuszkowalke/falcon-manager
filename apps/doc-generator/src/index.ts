import express from "express";
import cookieParser from "cookie-parser";
import path from "path";
import { mainRouter } from "./router";
import { hbs } from "./hbs";

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(cookieParser());

app.use(express.static(path.join(__dirname, "..", "public")));

app.engine("hbs", hbs.engine);
app.set("view engine", "hbs");
app.set("views", path.join(__dirname, "..", "views"));

app.use(mainRouter);

app.listen(4000, () => {
  console.log("Doc generator listening on port 4000");
});
