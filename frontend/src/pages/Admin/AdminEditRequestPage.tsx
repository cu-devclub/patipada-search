import { useParams } from "react-router-dom";
import { getRequestByParams, insertRequest } from "../../service/data";
import { useEffect, useState } from "react";
import { EditRecordForm } from "../../components/contributor/edit-record";
import { Grid, Heading } from "@chakra-ui/react";
import { BaseHeader, Footer, MessageToast } from "../../components";
import { Request, mapRequestToInsertRequestModels } from "../../models/request";
import { ToastStatus } from "../../constant";
import { useNavigate } from "react-router-dom";
import { getCookie } from "typescript-cookie";
function AdminEditRequestPage() {
  const { requestID } = useParams();
  const [data, setData] = useState<Request>();
  const { addToast } = MessageToast();
  const navigate = useNavigate();

  useEffect(() => {
    // 1. Try getting data from request first
    // 2. If not found, try searching
    const getRecord = async (requestID: string) => {
      await getRequestByParams({ requestID: requestID })
        .then((res) => {
          setData(res[0]);
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
    if (requestID) getRecord(requestID);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [requestID]);

  const submit = async (data: Request) => {
    data.by = getCookie("username") || "";
    const insertRequestData = mapRequestToInsertRequestModels(data);
    await insertRequest(insertRequestData)
      .then(() => {
        addToast({
          description: "ส่งคำขอแก้ไขสำเร็จ",
          status: ToastStatus.SUCCESS,
        });
      })
      .catch(() => {
        addToast({
          description: "เกิดข้อผิดพลาดขณะทำการส่งคำขอแก้ไข",
          status: ToastStatus.ERROR,
        });
      })
      .finally(() => {
        navigate(-1);
      });
  };

  return (
    <Grid templateRows="auto 1fr auto" gap={4} w="full" minH="100svh">
      <BaseHeader>
        <Heading>ตรวจสอบคำขอแก้ไขเนื้อหา</Heading>
      </BaseHeader>
      {data && <EditRecordForm data={data} submit={submit} />}
      <Footer />
    </Grid>
  );
}

export default AdminEditRequestPage;
