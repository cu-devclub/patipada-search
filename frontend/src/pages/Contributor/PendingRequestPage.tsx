import { Box, Grid } from "@chakra-ui/react";
import { PendingRequestTable } from "../../components/contributor/pending-request";
import { MessageToast } from "../../components";
import { PendingRequestHeader, Footer } from "../../components/layout";
import { useEffect, useState } from "react";
import { Request } from "../../models/request";
import { getRequestByParams } from "../../service/data";
import { getCookie } from "typescript-cookie";
import { ToastStatus } from "../../constant";
import { convertStatusWord, extractStringFromHTML } from "../../functions";

function PendingRequestPage() {
  const [data, setdata] = useState<Request[]>([]);
  const { addToast } = MessageToast();

  useEffect(() => {
    const getRequest = async () => {
      const username = getCookie("username");
      await getRequestByParams({ username: username })
        .then((res) => {
          const r = res.map((item) => {
            item.startTime = extractStringFromHTML(item.startTime);
            item.endTime = extractStringFromHTML(item.endTime);
            item.question = extractStringFromHTML(item.question);
            item.answer = extractStringFromHTML(item.answer);
            item.createdAt = new Date(item.createdAt).toLocaleString("th-TH");
            item.status = convertStatusWord(item.status);
            return item;
          });

          setdata(r);
          addToast({
            description: "ดึงข้อมูลสำเร็จ",
            status: ToastStatus.SUCCESS,
          });
        })
        .catch(() => {
          addToast({
            description: "เกิดข้อผิดพลาดขณะทำการดึงข้อมูล",
            status: ToastStatus.ERROR,
          });
        });
    };
    getRequest();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <Grid templateRows="auto 1fr auto" gap={4} w="full" minH="100svh">
      <PendingRequestHeader />
      <Box maxW="100%">
        <PendingRequestTable data={data} />
      </Box>
      <Footer />
    </Grid>
  );
}

export default PendingRequestPage;
