import { LoaderFunctionArgs } from "@remix-run/node";

import { UserIconSampleImage } from "@/components/common/UserIcon/UserIconSampleImage";

export const loader = ({ params }: LoaderFunctionArgs) => {
  // const report = await getReport(params.id);
  // const pdf = await generateReportPDF(report);
  return new Response(UserIconSampleImage, {
    status: 200,
    headers: {
      // "Content-Type": "application/pdf",
      // "Content-Type": "image/png",
      // "Content-Transfer-Encoding": "base64",
    },
  });
};
