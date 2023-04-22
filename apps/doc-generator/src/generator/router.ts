import { Router } from "express";
import { getGeneratorForm, postGeneratorForm } from "./controller";

const generatorRouter = Router();

generatorRouter.get("/", getGeneratorForm);
generatorRouter.post("/", postGeneratorForm);

export { generatorRouter };
