import { Request, Response } from "express";
import { generateDocData, GenerateDocData } from "./dto/generateDocData.dto";
import { generateDocFormInput } from "./dto/generateDocFormInput.dto";
import { generatorService } from "./service";

export async function getGeneratorForm(req: Request, res: Response) {
  const formData = await generatorService.getFormData(req.user!.sessionCookie);
  res.render("generator/generator-form", {
    title: "Doc generator",
    user: req.user,
    formData,
  });
}

export async function postGeneratorForm(req: Request, res: Response) {
  const body = req.body;
  if (typeof body.falconsIds === "string") {
    body.falconsIds = [body.falconsIds];
  }

  const { officeId, falconsIds } = generateDocFormInput.parse(body);

  const rawDocData: GenerateDocData = await generatorService.getDocData(
    req.user!.sessionCookie,
    officeId,
    falconsIds
  );

  const docData = generateDocData.parse(rawDocData);

  const pdf = await generatorService.generateBredInCaptivity(docData);

  // TODO
  // generatorService.sendRawDoc(req.user!.sessionCookie, pdf);

  res.contentType("application/pdf");
  res.setHeader("Content-Disposition", 'attachment; filename="doc.pdf"');
  res.send(pdf);
}
